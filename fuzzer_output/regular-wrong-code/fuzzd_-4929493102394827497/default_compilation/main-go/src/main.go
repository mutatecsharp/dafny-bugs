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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	if (false) && (false) {
		return !(true) || (!(true))
	} else {
		return !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(402))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})), _dafny.UnicodeSeqOfUtf8Bytes("pdxgcrjxl"))
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(!(!(false)))).Intersection((_dafny.MultiSetOf(!(false), true)).Intersection(_dafny.MultiSetOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	return ((func() _dafny.Int {
		if false {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(240), _dafny.IntOfInt64(-231), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rljhwsp")).Cardinality()), _dafny.IntOfInt64(206))).Cardinality())
		}
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xe")).Cardinality()), (func() _dafny.Set {
			var _coll0 = _dafny.NewBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(223))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _1_v0 _dafny.Int
				_1_v0 = interface{}(_compr_0).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(223)), _1_v0) {
					_coll0.Add((_1_v0).Plus(_dafny.IntOfInt64(741)))
				}
			}
			return _coll0.ToSet()
		}()).Cardinality())).Cardinality())).Cardinality()
	})()).Plus((_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('d'))).Cardinality())).Times(_dafny.IntOfInt64(-387)))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.UnicodeSeqOfUtf8Bytes("cs")
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("huytibwmr")
	}
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.CodePoint {
	var _source0 D2 = Companion_D2_.Create_DC6_(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC5_(_dafny.IntOfInt64(-139), true, true, true)))
	_ = _source0
	if _source0.Is_DC4() {
		var _2___mcc_h0 _dafny.Int = _source0.Get_().(D2_DC4).Cf4
		_ = _2___mcc_h0
		var _3___mcc_h1 bool = _source0.Get_().(D2_DC4).Cf5
		_ = _3___mcc_h1
		var _4_cf5 bool = _3___mcc_h1
		_ = _4_cf5
		var _5_cf4 _dafny.Int = _2___mcc_h0
		_ = _5_cf4
		return _dafny.CodePoint('t')
	} else if _source0.Is_DC5() {
		var _6___mcc_h2 _dafny.Int = _source0.Get_().(D2_DC5).Cf6
		_ = _6___mcc_h2
		var _7___mcc_h3 bool = _source0.Get_().(D2_DC5).Cf7
		_ = _7___mcc_h3
		var _8___mcc_h4 bool = _source0.Get_().(D2_DC5).Cf8
		_ = _8___mcc_h4
		var _9___mcc_h5 bool = _source0.Get_().(D2_DC5).Cf9
		_ = _9___mcc_h5
		var _10_cf9 bool = _9___mcc_h5
		_ = _10_cf9
		var _11_cf8 bool = _8___mcc_h4
		_ = _11_cf8
		var _12_cf7 bool = _7___mcc_h3
		_ = _12_cf7
		var _13_cf6 _dafny.Int = _6___mcc_h2
		_ = _13_cf6
		return _dafny.CodePoint('v')
	} else if _source0.Is_DC3() {
		var _14___mcc_h6 _dafny.Sequence = _source0.Get_().(D2_DC3).Cf3
		_ = _14___mcc_h6
		var _15_cf3 _dafny.Sequence = _14___mcc_h6
		_ = _15_cf3
		if false {
			return _dafny.CodePoint('e')
		} else {
			return _dafny.CodePoint('n')
		}
	} else {
		var _16___mcc_h7 D2 = _source0.Get_().(D2_DC6).Cf10
		_ = _16___mcc_h7
		var _17_cf10 D2 = _16___mcc_h7
		_ = _17_cf10
		return _dafny.CodePoint('o')
	}
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC4_((_dafny.IntOfInt64(855)).Plus(_dafny.IntOfInt64(195)), (_dafny.SetOf(_dafny.IntOfInt64(188))).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(-694), _dafny.IntOfInt64(297))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.Zero).Minus((_dafny.IntOfInt64(354)).Minus(_dafny.IntOfInt64(954))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(!(!(false))))), _dafny.IntOfInt64(430))).Merge((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false, !(true)), _dafny.SeqOf(false), _dafny.SeqOf(true))).Cardinality()))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('a')
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(!(true)), false, false, (_dafny.SetOf(_dafny.IntOfInt64(360))).IsSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(35))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ckp"), (func() _dafny.Sequence {
		if !(false) {
			return _dafny.UnicodeSeqOfUtf8Bytes("dy")
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(360))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_18_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))
	})())
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kimkmjliq")).Cardinality())))).Cardinality(), _dafny.CodePoint('x'))).Merge(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-691), _dafny.IntOfInt64(373))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_1).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-691), _dafny.IntOfInt64(373)), _19_v0) {
				_coll1.Add((_19_v0).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(310))).Cardinality())), _dafny.CodePoint('x'))
			}
		}
		return _coll1.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(Companion_D3_.Create_DC9_(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.CodePoint('c'), _dafny.IntOfInt64(512)))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _20_v1 D3
			_20_v1 = interface{}(_compr_2).(D3)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D3_.Create_DC9_(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.CodePoint('c'), _dafny.IntOfInt64(512))), _20_v1) {
				_coll2.Add(_20_v1, false)
			}
		}
		return _coll2.ToMap()
	}()).Cardinality(), _dafny.CodePoint('d')))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(Companion_D2_.Create_DC5_((_dafny.MultiSetOf(_dafny.CodePoint('n'))).Cardinality(), !(!(true)), !(false), false))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(640)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-152))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_21_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("moh")).Cardinality())
	}))).Cardinality()), _dafny.IntOfInt64(532)))).Cardinality()), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(798), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hnvcppjf")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	if !(false) {
		return _dafny.SetOf(_dafny.CodePoint('p'))
	} else {
		return _dafny.SetOf(_dafny.CodePoint('h'), _dafny.CodePoint('l'))
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC21_(_dafny.IntOfInt64(-405), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf((_dafny.SetOf(!(!(false)))).Cardinality())).Union(_dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bdxocc")).Cardinality())))))).Union(_dafny.SetOf(_dafny.IntOfInt64(449)))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) D3 {
	if false {
		return Companion_D3_.Create_DC7_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC22_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(693), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Cardinality(), _dafny.IntOfInt64(42))).Cardinality(), (_dafny.SetOf(_dafny.CodePoint('p'))).Cardinality(), false)).Dtor_cf40(), true))
	} else {
		return Companion_D3_.Create_DC7_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(319), true))
	}
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), (func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(385))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_22_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-738))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_23_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
			}))).Cardinality())
		}))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _24_v0 _dafny.Int
			_24_v0 = interface{}(_compr_3).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(385))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_22_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-738))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_23_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
				}))).Cardinality())
			})), _24_v0) {
				_coll3.Add((_24_v0).Plus(_dafny.IntOfInt64(686)))
			}
		}
		return _coll3.ToSet()
	}()).Cardinality())).Cardinality(), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-473), true))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Map {
	var _source1 D3 = Companion_D3_.Create_DC8_(false)
	_ = _source1
	if _source1.Is_DC8() {
		var _25___mcc_h0 bool = _source1.Get_().(D3_DC8).Cf12
		_ = _25___mcc_h0
		var _26_cf12 bool = _25___mcc_h0
		_ = _26_cf12
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_cf12, !(!(_26_cf12)))
	} else if _source1.Is_DC9() {
		var _27___mcc_h1 _dafny.Int = _source1.Get_().(D3_DC9).Cf13
		_ = _27___mcc_h1
		var _28___mcc_h2 _dafny.CodePoint = _source1.Get_().(D3_DC9).Cf14
		_ = _28___mcc_h2
		var _29___mcc_h3 _dafny.Int = _source1.Get_().(D3_DC9).Cf15
		_ = _29___mcc_h3
		var _30_cf15 _dafny.Int = _29___mcc_h3
		_ = _30_cf15
		var _31_cf14 _dafny.CodePoint = _28___mcc_h2
		_ = _31_cf14
		var _32_cf13 _dafny.Int = _27___mcc_h1
		_ = _32_cf13
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
	} else if _source1.Is_DC7() {
		var _33___mcc_h4 _dafny.Map = _source1.Get_().(D3_DC7).Cf11
		_ = _33___mcc_h4
		var _34_cf11 _dafny.Map = _33___mcc_h4
		_ = _34_cf11
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), !(!(false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	} else {
		var _35___mcc_h5 D3 = _source1.Get_().(D3_DC10).Cf16
		_ = _35___mcc_h5
		var _36_cf16 D3 = _35___mcc_h5
		_ = _36_cf16
		return (Companion_D7_.Create_DC21_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Dtor_cf39()
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Union(_dafny.SetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(192), _dafny.IntOfInt64(255))).Cardinality()))), func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(235), _dafny.IntOfInt64(-647))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _37_v0 _dafny.Int
			_37_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(235)).Cmp(_37_v0) <= 0) && ((_37_v0).Cmp(_dafny.IntOfInt64(-647)) < 0) {
				_coll4.Add((_37_v0).Minus(_dafny.IntOfInt64(54)), true)
			}
		}
		return _coll4.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _38_v0 _dafny.Int
	_ = _38_v0
	_38_v0 = _dafny.IntOfInt64(142)
	var _39_v1 _dafny.Sequence
	_ = _39_v1
	_39_v1 = _dafny.SeqOf(_38_v0)
	var _40_v2 _dafny.Sequence
	_ = _40_v2
	_40_v2 = _dafny.UnicodeSeqOfUtf8Bytes("rfkqvu")
	var _41_v3 _dafny.Set
	_ = _41_v3
	_41_v3 = _dafny.SetOf(_38_v0, (_39_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_40_v2).Cardinality()), _dafny.IntOfUint32((_39_v1).Cardinality()))).Uint32()).(_dafny.Int))
	var _hi0 _dafny.Int = Companion_Default___.SafeDivisionInt(_38_v0, (_41_v3).Cardinality())
	_ = _hi0
	for _42_i0 := _dafny.IntOfInt64(927); _42_i0.Cmp(_hi0) < 0; _42_i0 = _42_i0.Plus(_dafny.One) {
		var _43_v4 bool
		_ = _43_v4
		_43_v4 = false
		(globalState).F2 = _43_v4
		var _44_v5 D4
		_ = _44_v5
		_44_v5 = Companion_D4_.Create_DC11_(_40_v2)
		_43_v4 = _dafny.Companion_Sequence_.IsProperPrefixOf(_40_v2, (_44_v5).Dtor_cf17())
		var _45_v6 *C0
		_ = _45_v6
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__()
		_45_v6 = _nw0
		var _46_v7 _dafny.CodePoint
		_ = _46_v7
		_46_v7 = _dafny.CodePoint('m')
		var _47_v8 _dafny.Set
		_ = _47_v8
		_47_v8 = _dafny.SetOf(_46_v7, _46_v7, _46_v7, _46_v7, _46_v7)
		var _48_v9 _dafny.MultiSet
		_ = _48_v9
		_48_v9 = _dafny.MultiSetOf(_47_v8)
		var _rhs0 _dafny.Int = _42_i0
		_ = _rhs0
		var _rhs1 _dafny.Int = Companion_Default___.SafeDivisionInt(_38_v0, _42_i0)
		_ = _rhs1
		var _rhs2 _dafny.MultiSet = (_48_v9).Intersection(_48_v9)
		_ = _rhs2
		var _rhs3 _dafny.Int = _38_v0
		_ = _rhs3
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		_lhs0.F18 = _rhs0
		_lhs1.F21 = _rhs1
		_48_v9 = _rhs2
		_lhs2.F5 = _rhs3
	}
	if false {
		var _49_v10 bool
		_ = _49_v10
		_49_v10 = false
		(globalState).F2 = !(_dafny.Companion_Sequence_.IsPrefixOf(_40_v2, (func() _dafny.Sequence {
			if _49_v10 {
				return _40_v2
			}
			return _40_v2
		})()))
		var _50_v11 _dafny.Set
		_ = _50_v11
		_50_v11 = _dafny.SetOf(_49_v10, _49_v10, _49_v10)
		if !(Companion_Default___.Fm0((_dafny.MultiSetOf((_50_v11).Cardinality(), _38_v0)).Cardinality(), Companion_Default___.SafeDivisionInt(_38_v0, _dafny.IntOfInt64(775)), _49_v10, globalState)) {
			var _51_v12 _dafny.CodePoint
			_ = _51_v12
			_51_v12 = _dafny.CodePoint('u')
			var _52_v13 _dafny.Array
			_ = _52_v13
			var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
			_ = _nw1
			_52_v13 = _nw1
			var _53_v14 *C2
			_ = _53_v14
			var _nw2 *C2 = New_C2_()
			_ = _nw2
			_nw2.Ctor__(_51_v12, _52_v13, true)
			_53_v14 = _nw2
			var _54_v15 _dafny.Array
			_ = _54_v15
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw3
			_54_v15 = _nw3
			var _55_v16 D4
			_ = _55_v16
			_55_v16 = Companion_D4_.Create_DC12_(true, _54_v15, _38_v0)
			(globalState).F2 = (_55_v16).Dtor_cf18()
			var _56_v17 _dafny.Sequence
			_ = _56_v17
			_56_v17 = _dafny.SeqOf(_49_v10)
			(globalState).F2 = (_dafny.IntOfUint32((_56_v17).Cardinality())).Cmp(_38_v0) >= 0
			_54_v15 = _54_v15
			var _57_v18 *C0
			_ = _57_v18
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__()
			_57_v18 = _nw4
		} else {
			(globalState).F2 = _49_v10
			var _58_v19 _dafny.Array
			_ = _58_v19
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_0
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = (func(_59_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_60_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_60_i1, _59_v0)
					}
				})(_38_v0)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw5 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw5).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw5).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_58_v19 = _nw5
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_58_v19), 0))
			_ = _index0
			(_58_v19).ArraySet1((func() _dafny.Int {
				if _49_v10 {
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_39_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ec")).Cardinality()), _dafny.IntOfUint32((_39_v1).Cardinality()))).Uint32(), _38_v0)).Cardinality())
				}
				return (_41_v3).Cardinality()
			})(), (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_58_v19), 0))
			_ = _index1
			(_58_v19).ArraySet1((_39_v1).Select((Companion_Default___.SafeIndex(_38_v0, _dafny.IntOfUint32((_39_v1).Cardinality()))).Uint32()).(_dafny.Int), (_index1).Int())
			_40_v2 = _40_v2
			(globalState).F21 = (_58_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_58_v19), 0))).Int()).(_dafny.Int)
			var _61_v20 _dafny.Array
			_ = _61_v20
			var _nwElement0_0 bool = true
			_ = _nwElement0_0
			var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(7))
			_ = _nw6
			(_nw6).ArraySet1(_nwElement0_0, 0)
			(_nw6).ArraySet1(_49_v10, 1)
			(_nw6).ArraySet1(!(_49_v10), 2)
			(_nw6).ArraySet1(_49_v10, 3)
			(_nw6).ArraySet1(_49_v10, 4)
			(_nw6).ArraySet1(_49_v10, 5)
			(_nw6).ArraySet1(_49_v10, 6)
			_61_v20 = _nw6
			var _62_v21 _dafny.MultiSet
			_ = _62_v21
			_62_v21 = _dafny.MultiSetOf(_61_v20)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_58_v19), 0))
			_ = _index2
			(_58_v19).ArraySet1((_62_v21).Cardinality(), (_index2).Int())
		}
		(globalState).F5 = Companion_Default___.SafeModuloInt(_38_v0, _dafny.IntOfInt64(860))
		var _63_v22 *C0
		_ = _63_v22
		var _nw7 *C0 = New_C0_()
		_ = _nw7
		_nw7.Ctor__()
		_63_v22 = _nw7
		var _64_v23 _dafny.Sequence
		_ = _64_v23
		_64_v23 = _dafny.SeqOf(!(_49_v10), _49_v10)
		var _65_v24 _dafny.Array
		_ = _65_v24
		var _nwElement0_1 _dafny.Int = _38_v0
		_ = _nwElement0_1
		var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(12))
		_ = _nw8
		(_nw8).ArraySet1(_nwElement0_1, 0)
		(_nw8).ArraySet1(_38_v0, 1)
		(_nw8).ArraySet1(_dafny.IntOfUint32((_64_v23).Cardinality()), 2)
		(_nw8).ArraySet1(_38_v0, 3)
		(_nw8).ArraySet1(((_dafny.Zero).Minus(_38_v0)).Plus((_dafny.Zero).Minus(_38_v0)), 4)
		(_nw8).ArraySet1(_38_v0, 5)
		(_nw8).ArraySet1(_38_v0, 6)
		(_nw8).ArraySet1(_38_v0, 7)
		(_nw8).ArraySet1(_38_v0, 8)
		(_nw8).ArraySet1(_38_v0, 9)
		(_nw8).ArraySet1((_dafny.Zero).Minus((_38_v0).Times(_dafny.IntOfInt64(250))), 10)
		(_nw8).ArraySet1(_38_v0, 11)
		_65_v24 = _nw8
		var _66_v25 _dafny.Map
		_ = _66_v25
		_66_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v0, _49_v10)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_65_v24), 0))
		_ = _index3
		(_65_v24).ArraySet1(Companion_Default___.SafeDivisionInt(_38_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v0, _66_v25)).Cardinality()), (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_65_v24), 0))
		_ = _index4
		(_65_v24).ArraySet1((_63_v22).Fm5(globalState), (_index4).Int())
	} else {
		var _67_v26 _dafny.CodePoint
		_ = _67_v26
		_67_v26 = _dafny.CodePoint('p')
		var _68_v27 _dafny.MultiSet
		_ = _68_v27
		_68_v27 = _dafny.MultiSetOf(_67_v26)
		var _69_v28 _dafny.Array
		_ = _69_v28
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
		_ = _nw9
		_69_v28 = _nw9
		var _70_v29 bool
		_ = _70_v29
		_70_v29 = false
		var _71_v30 D2
		_ = _71_v30
		_71_v30 = Companion_D2_.Create_DC4_(_38_v0, _70_v29)
		var _pat_let_tv0 = _70_v29
		_ = _pat_let_tv0
		var _72_v31 _dafny.Array
		_ = _72_v31
		var _nwElement0_2 D2 = _71_v30
		_ = _nwElement0_2
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(26))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_2, 0)
		(_nw10).ArraySet1(_71_v30, 1)
		(_nw10).ArraySet1(_71_v30, 2)
		(_nw10).ArraySet1(_71_v30, 3)
		(_nw10).ArraySet1(_71_v30, 4)
		(_nw10).ArraySet1(_71_v30, 5)
		(_nw10).ArraySet1(_71_v30, 6)
		(_nw10).ArraySet1(Companion_D2_.Create_DC4_((_dafny.Zero).Minus(_38_v0), _70_v29), 7)
		(_nw10).ArraySet1(_71_v30, 8)
		(_nw10).ArraySet1(_71_v30, 9)
		(_nw10).ArraySet1(_71_v30, 10)
		(_nw10).ArraySet1(Companion_D2_.Create_DC4_(_38_v0, _70_v29), 11)
		(_nw10).ArraySet1(_71_v30, 12)
		(_nw10).ArraySet1(_71_v30, 13)
		(_nw10).ArraySet1(_71_v30, 14)
		(_nw10).ArraySet1(_71_v30, 15)
		(_nw10).ArraySet1(func(_pat_let0_0 D2) D2 {
			return func(_73_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let1_0 bool) D2 {
					return func(_74_dt__update_hcf5_h0 bool) D2 {
						return Companion_D2_.Create_DC4_((_73_dt__update__tmp_h0).Dtor_cf4(), _74_dt__update_hcf5_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(Companion_D2_.Create_DC4_(_38_v0, _70_v29)), 16)
		(_nw10).ArraySet1(_71_v30, 17)
		(_nw10).ArraySet1(_71_v30, 18)
		(_nw10).ArraySet1(_71_v30, 19)
		(_nw10).ArraySet1(_71_v30, 20)
		(_nw10).ArraySet1(Companion_D2_.Create_DC4_(_dafny.IntOfUint32((_40_v2).Cardinality()), _70_v29), 21)
		(_nw10).ArraySet1(_71_v30, 22)
		(_nw10).ArraySet1(_71_v30, 23)
		(_nw10).ArraySet1(Companion_D2_.Create_DC4_((_dafny.Zero).Minus(_38_v0), false), 24)
		(_nw10).ArraySet1(_71_v30, 25)
		_72_v31 = _nw10
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_69_v28), 0))
		_ = _index5
		(_69_v28).ArraySet1(_72_v31, (_index5).Int())
		var _75_v32 _dafny.Set
		_ = _75_v32
		_75_v32 = _dafny.SetOf(_70_v29, _70_v29)
		var _76_v33 _dafny.Sequence
		_ = _76_v33
		_76_v33 = _dafny.SeqOf(_72_v31, _72_v31, (func() _dafny.Array {
			if _70_v29 {
				return _72_v31
			}
			return _72_v31
		})())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_69_v28), 0))
		_ = _index6
		var _rhs4 _dafny.MultiSet = ((_68_v27).Update(_67_v26, Companion_Default___.Abs((_75_v32).Cardinality()))).Difference(_68_v27)
		_ = _rhs4
		var _rhs5 _dafny.Array = (_76_v33).Select((Companion_Default___.SafeIndex(_38_v0, _dafny.IntOfUint32((_76_v33).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs5
		var _rhs6 _dafny.Int = (_dafny.Zero).Minus(_38_v0)
		_ = _rhs6
		var _rhs7 _dafny.Set = _41_v3
		_ = _rhs7
		var _rhs8 _dafny.Int = Companion_Default___.Fm2((_38_v0).Cmp((_dafny.MultiSetFromSeq(_39_v1)).Cardinality()) == 0, (_dafny.Zero).Minus(Companion_Default___.Fm2(_70_v29, (_dafny.Zero).Minus(_38_v0), _70_v29, _67_v26, globalState)), (_38_v0).Cmp(_38_v0) >= 0, _dafny.CodePoint('d'), globalState)
		_ = _rhs8
		var _lhs3 _dafny.Array = _69_v28
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_69_v28), 0))
		_ = _lhs4
		var _lhs5 *GlobalState = globalState
		_ = _lhs5
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		_68_v27 = _rhs4
		(_lhs3).ArraySet1(_rhs5, (_lhs4).Int())
		_lhs5.F18 = _rhs6
		_41_v3 = _rhs7
		_lhs6.F5 = _rhs8
		(globalState).F2 = _70_v29
		var _77_v34 _dafny.Array
		_ = _77_v34
		var _nw11 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
		_ = _nw11
		_77_v34 = _nw11
		var _78_v35 _dafny.Array
		_ = _78_v35
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_1
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.MultiSet = (func(_79_v0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
				return func(_80_i2 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_79_v0, _79_v0)
				}
			})(_38_v0)
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
		_78_v35 = _nw12
		var _81_v36 *C2
		_ = _81_v36
		var _nw13 *C2 = New_C2_()
		_ = _nw13
		_nw13.Ctor__(_67_v26, _78_v35, _70_v29)
		_81_v36 = _nw13
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_77_v34), 0))
		_ = _index7
		(_77_v34).ArraySet1(_81_v36, (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_77_v34), 0))
		_ = _index8
		(_77_v34).ArraySet1(_81_v36, (_index8).Int())
		var _82_v37 _dafny.Array
		_ = _82_v37
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_2
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_83_v2 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_84_i3 _dafny.Int) bool {
					return !_dafny.Companion_Sequence_.Contains(_83_v2, _dafny.CodePoint('n'))
				}
			})(_40_v2)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw14 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw14).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw14).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_82_v37 = _nw14
		var _85_v38 _dafny.Sequence
		_ = _85_v38
		_85_v38 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-407))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_86_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_87_i4 _dafny.Int) _dafny.Int {
				return _86_v0
			}
		})(_38_v0))))
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_82_v37), 0))
		_ = _index9
		(_82_v37).ArraySet1((_dafny.IntOfUint32(((_85_v38).Select((Companion_Default___.SafeIndex(_38_v0, _dafny.IntOfUint32((_85_v38).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Cmp(_38_v0) != 0, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(744), _dafny.ArrayLen((_82_v37), 0))
		_ = _index10
		(_82_v37).ArraySet1((func() bool {
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_39_v1, _dafny.SeqOf(_38_v0, _dafny.IntOfUint32((_39_v1).Cardinality()))) {
				return _70_v29
			}
			return _70_v29
		})(), (_index10).Int())
		(globalState).F19 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-27))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_88_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_89_i5 _dafny.Int) _dafny.Int {
				return _88_v0
			}
		})(_38_v0)))
	}
	var _90_v39 bool
	_ = _90_v39
	_90_v39 = false
	var _91_v40 _dafny.CodePoint
	_ = _91_v40
	_91_v40 = _dafny.CodePoint('p')
	var _92_v41 T0
	_ = _92_v41
	var _nw15 *C1 = New_C1_()
	_ = _nw15
	_nw15.Ctor__(_40_v2, _38_v0, _90_v39)
	_92_v41 = _nw15
	var _93_v42 _dafny.Sequence
	_ = _93_v42
	_93_v42 = _dafny.SeqOf(_92_v41)
	var _94_v43 _dafny.Sequence
	_ = _94_v43
	_94_v43 = _dafny.SeqOf((_93_v42).Select((Companion_Default___.SafeIndex(_38_v0, _dafny.IntOfUint32((_93_v42).Cardinality()))).Uint32()).(T0))
	var _95_v44 _dafny.Map
	_ = _95_v44
	_95_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_90_v39, _38_v0, _90_v39, _91_v40, globalState), _dafny.IntOfUint32((_94_v43).Cardinality()))
	(globalState).F2 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_40_v2).Cardinality()), _38_v0)), (_95_v44).Merge(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_39_v1).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _96_v45 _dafny.Int
			_96_v45 = interface{}(_compr_5).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_39_v1, _96_v45) {
				_coll5.Add(Companion_Default___.SafeDivisionInt(_96_v45, _38_v0), _dafny.IntOfInt64(2))
			}
		}
		return _coll5.ToMap()
	}()))
	var _97_v46 _dafny.Array
	_ = _97_v46
	var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
	_ = _nw16
	_97_v46 = _nw16
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_97_v46), 0))); ; {
		_guard_loop_0, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _98_i6 _dafny.Int
		_98_i6 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_98_i6).Sign() != -1) && ((_98_i6).Cmp(_dafny.ArrayLen((_97_v46), 0)) < 0)) {
			(_97_v46).ArraySet1(_dafny.MultiSetOf((Companion_D2_.Create_DC5_(_38_v0, (_92_v41).F25(), (_92_v41).F25(), (_92_v41).F25())).Dtor_cf6(), _dafny.IntOfInt64(843)), (_98_i6).Int())
		}
	}
	var _99_v47 _dafny.MultiSet
	_ = _99_v47
	_99_v47 = _dafny.MultiSetOf(_38_v0)
	var _100_v49 _dafny.Map
	_ = _100_v49
	_100_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v0, Companion_Default___.Fm0(_dafny.IntOfInt64(-145), (_dafny.Zero).Minus(_38_v0), true, globalState))
	var _101_v50 _dafny.Set
	_ = _101_v50
	_101_v50 = _dafny.SetOf((func() bool {
		if (_100_v49).Contains(_38_v0) {
			return (_100_v49).Get(_38_v0).(bool)
		}
		return _90_v39
	})(), _90_v39)
	var _102_v51 _dafny.Map
	_ = _102_v51
	_102_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_92_v41).F25(), _91_v40)
	var _103_v52 _dafny.Map
	_ = _103_v52
	_103_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v0, _102_v51)
	var _104_v53 _dafny.Map
	_ = _104_v53
	_104_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_92_v41).F25(), _90_v39)
	var _105_v54 D7
	_ = _105_v54
	_105_v54 = Companion_D7_.Create_DC21_(_38_v0, Companion_Default___.Fm27(globalState))
	var _106_v55 *C3
	_ = _106_v55
	var _nw17 *C3 = New_C3_()
	_ = _nw17
	_nw17.Ctor__(_91_v40, _97_v46, (_92_v41).F25())
	_106_v55 = _nw17
	var _107_v56 _dafny.Array
	_ = _107_v56
	var _nwElement0_3 _dafny.Int = _dafny.IntOfUint32((_39_v1).Cardinality())
	_ = _nwElement0_3
	var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(28))
	_ = _nw18
	(_nw18).ArraySet1(_nwElement0_3, 0)
	(_nw18).ArraySet1((_dafny.Zero).Minus(_38_v0), 1)
	(_nw18).ArraySet1(_38_v0, 2)
	(_nw18).ArraySet1(_38_v0, 3)
	(_nw18).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm2(Companion_Default___.Fm0(_38_v0, _38_v0, _90_v39, globalState), (_99_v47).Cardinality(), _90_v39, _91_v40, globalState), _38_v0), 4)
	(_nw18).ArraySet1(_38_v0, 5)
	(_nw18).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(616))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}((func(_108_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_109_i7 _dafny.Int) _dafny.CodePoint {
			return _108_v40
		}
	})(_91_v40)))).Cardinality()), 6)
	(_nw18).ArraySet1(_38_v0, 7)
	(_nw18).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_38_v0).Times(_38_v0))), 8)
	(_nw18).ArraySet1(_38_v0, 9)
	(_nw18).ArraySet1(_38_v0, 10)
	(_nw18).ArraySet1(((func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(117), _dafny.IntOfInt64(380))); ; {
			_compr_6, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _110_v48 _dafny.Int
			_110_v48 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(117)).Cmp(_110_v48) <= 0) && ((_110_v48).Cmp(_dafny.IntOfInt64(380)) < 0) {
				_coll6.Add(Companion_Default___.SafeDivisionInt(_110_v48, _38_v0))
			}
		}
		return _coll6.ToSet()
	}()).Cardinality()).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_38_v0))), 11)
	(_nw18).ArraySet1(((_101_v50).Cardinality()).Plus(_dafny.IntOfInt64(4)), 12)
	(_nw18).ArraySet1(_38_v0, 13)
	(_nw18).ArraySet1(_38_v0, 14)
	(_nw18).ArraySet1(_dafny.IntOfInt64(111), 15)
	(_nw18).ArraySet1(((_103_v52).Cardinality()).Times(_38_v0), 16)
	(_nw18).ArraySet1(_38_v0, 17)
	(_nw18).ArraySet1(Companion_Default___.Fm2(_90_v39, _dafny.IntOfInt64(191), _90_v39, _91_v40, globalState), 18)
	(_nw18).ArraySet1(_38_v0, 19)
	(_nw18).ArraySet1(_38_v0, 20)
	(_nw18).ArraySet1(_38_v0, 21)
	(_nw18).ArraySet1((((_104_v53).Update(_90_v39, (_92_v41).F25())).Merge((_105_v54).Dtor_cf39())).Cardinality(), 22)
	(_nw18).ArraySet1(_38_v0, 23)
	(_nw18).ArraySet1(Companion_Default___.SafeModuloInt(_38_v0, (_dafny.Zero).Minus(_38_v0)), 24)
	(_nw18).ArraySet1(_dafny.IntOfInt64(77), 25)
	(_nw18).ArraySet1(_38_v0, 26)
	(_nw18).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v55, _38_v0)).Cardinality(), 27)
	_107_v56 = _nw18
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))
	_ = _index11
	(_107_v56).ArraySet1((_dafny.Zero).Minus((_38_v0).Minus(_dafny.IntOfUint32((_40_v2).Cardinality()))), (_index11).Int())
	var _111_v57 *C1
	_ = _111_v57
	var _nw19 *C1 = New_C1_()
	_ = _nw19
	_nw19.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("s"), _dafny.IntOfUint32((_40_v2).Cardinality()), (_92_v41).F25())
	_111_v57 = _nw19
	var _112_v58 D1
	_ = _112_v58
	_112_v58 = Companion_D1_.Create_DC1_(_90_v39)
	var _113_v59 _dafny.Map
	_ = _113_v59
	_113_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v57, _112_v58)
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))
	_ = _index12
	(_107_v56).ArraySet1((_dafny.Zero).Minus((((_113_v59).Update(_111_v57, _112_v58)).Merge(_113_v59)).Cardinality()), (_index12).Int())
	if !(true) || (_90_v39) {
		var _114_v60 _dafny.Map
		_ = _114_v60
		_114_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v57.F29, _dafny.SetOf(_90_v39, (_92_v41).F25(), _90_v39, (_92_v41).F25()))
		var _115_v61 _dafny.Sequence
		_ = _115_v61
		_115_v61 = _dafny.SeqOf((_92_v41).F25())
		var _116_v62 _dafny.Map
		_ = _116_v62
		_116_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
			if (_114_v60).Contains(_111_v57.F29) {
				return (_114_v60).Get(_111_v57.F29).(_dafny.Set)
			}
			return _101_v50
		})(), _115_v61)
		var _117_v63 _dafny.Sequence
		_ = _117_v63
		_117_v63 = _dafny.SeqOf(_116_v62)
		var _118_v64 _dafny.Sequence
		_ = _118_v64
		_118_v64 = _dafny.SeqOf(_99_v47)
		var _119_v66 _dafny.Set
		_ = _119_v66
		_119_v66 = _dafny.SetOf(_99_v47)
		var _120_v67 _dafny.Array
		_ = _120_v67
		var _nw20 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
		_ = _nw20
		_120_v67 = _nw20
		var _121_v68 D4
		_ = _121_v68
		_121_v68 = Companion_D4_.Create_DC13_(_120_v67, _90_v39, (_92_v41).F25())
		var _122_v69 _dafny.Map
		_ = _122_v69
		_122_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(38), _dafny.IntOfUint32((_39_v1).Cardinality())), (_92_v41).F25())
		var _123_v70 _dafny.Array
		_ = _123_v70
		var _nwElement0_4 bool = _90_v39
		_ = _nwElement0_4
		var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(29))
		_ = _nw21
		(_nw21).ArraySet1(_nwElement0_4, 0)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_40_v2, _40_v2), 1)
		(_nw21).ArraySet1((_dafny.SetOf((_92_v41).F25(), (_92_v41).F25())).IsSubsetOf(_101_v50), 2)
		(_nw21).ArraySet1((_92_v41).F25(), 3)
		(_nw21).ArraySet1(true, 4)
		(_nw21).ArraySet1(false, 5)
		(_nw21).ArraySet1((_92_v41).F25(), 6)
		(_nw21).ArraySet1(((_117_v63).Select((Companion_Default___.SafeIndex(_111_v57.F29, _dafny.IntOfUint32((_117_v63).Cardinality()))).Uint32()).(_dafny.Map)).Contains(_101_v50), 7)
		(_nw21).ArraySet1((_92_v41).F25(), 8)
		(_nw21).ArraySet1(_90_v39, 9)
		(_nw21).ArraySet1((_92_v41).F25(), 10)
		(_nw21).ArraySet1((_92_v41).F25(), 11)
		(_nw21).ArraySet1((_92_v41).F25(), 12)
		(_nw21).ArraySet1(!(_95_v44).Contains((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)), 13)
		(_nw21).ArraySet1(true, 14)
		(_nw21).ArraySet1((_92_v41).F25(), 15)
		(_nw21).ArraySet1(_90_v39, 16)
		(_nw21).ArraySet1((_92_v41).F25(), 17)
		(_nw21).ArraySet1((_92_v41).F25(), 18)
		(_nw21).ArraySet1((_92_v41).F25(), 19)
		(_nw21).ArraySet1((_92_v41).F25(), 20)
		(_nw21).ArraySet1((_92_v41).F25(), 21)
		(_nw21).ArraySet1((func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter8 := _dafny.Iterate((_118_v64).Elements()); ; {
				_compr_7, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _124_v65 _dafny.MultiSet
				_124_v65 = interface{}(_compr_7).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_118_v64, _124_v65) {
					_coll7.Add(_124_v65)
				}
			}
			return _coll7.ToSet()
		}()).IsDisjointFrom(_119_v66), 22)
		(_nw21).ArraySet1((_121_v68).Dtor_cf22(), 23)
		(_nw21).ArraySet1((_dafny.MultiSetFromSeq(_39_v1)).IsSubsetOf(_99_v47), 24)
		(_nw21).ArraySet1(!(((_122_v69).Cardinality()).Cmp(_dafny.IntOfInt64(91)) <= 0), 25)
		(_nw21).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_111_v57.F28, _40_v2), 26)
		(_nw21).ArraySet1((_111_v57.F29).Cmp((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)) > 0, 27)
		(_nw21).ArraySet1((_38_v0).Cmp((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)) >= 0, 28)
		_123_v70 = _nw21
		var _125_v71 _dafny.Map
		_ = _125_v71
		_125_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_99_v47).Intersection(_99_v47), _106_v55)
		var _rhs9 _dafny.Array = _123_v70
		_ = _rhs9
		var _rhs10 *C3 = (func() *C3 {
			if (_125_v71).Contains(_99_v47) {
				return (_125_v71).Get(_99_v47).(*C3)
			}
			return _106_v55
		})()
		_ = _rhs10
		_123_v70 = _rhs9
		_106_v55 = _rhs10
		var _126_v72 D5
		_ = _126_v72
		_126_v72 = Companion_D5_.Create_DC15_((_92_v41).F25(), _123_v70, _111_v57.F28, _90_v39)
		var _127_v73 D5
		_ = _127_v73
		_127_v73 = Companion_D5_.Create_DC16_(_126_v72)
		var _128_v74 _dafny.Map
		_ = _128_v74
		_128_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_38_v0).Cmp(_111_v57.F29) < 0, Companion_D5_.Create_DC16_(_126_v72))
		var _129_v75 D5
		_ = _129_v75
		_129_v75 = Companion_D5_.Create_DC16_(_126_v72)
		_128_v74 = (_128_v74).Update(_90_v39, _129_v75)
		var _130_v76 T1
		_ = _130_v76
		var _nw22 *C3 = New_C3_()
		_ = _nw22
		_nw22.Ctor__(_91_v40, _97_v46, (_92_v41).F25())
		_130_v76 = _nw22
		var _131_v77 _dafny.Sequence
		_ = _131_v77
		_131_v77 = _dafny.SeqOf(_130_v76)
		_131_v77 = _dafny.Companion_Sequence_.Concatenate(_131_v77, _dafny.Companion_Sequence_.Concatenate(_131_v77, _131_v77))
		var _132_v78 *C0
		_ = _132_v78
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__()
		_132_v78 = _nw23
		(globalState).F18 = _111_v57.F29
	} else {
		(globalState).F22 = (_dafny.Zero).Minus(_111_v57.F29)
		var _133_v79 _dafny.Map
		_ = _133_v79
		_133_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-974), _92_v41)
		if (_133_v79).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-283), _92_v41)) {
			var _134_v80 _dafny.Array
			_ = _134_v80
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw24
			_134_v80 = _nw24
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_134_v80), 0))
			_ = _index13
			(_134_v80).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_111_v57.F28, _40_v2), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_134_v80), 0))
			_ = _index14
			(_134_v80).ArraySet1(_90_v39, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_134_v80), 0))
			_ = _index15
			var _rhs11 bool = !_dafny.Companion_Sequence_.Equal(_111_v57.F28, _111_v57.F28)
			_ = _rhs11
			var _rhs12 *C3 = _106_v55
			_ = _rhs12
			var _lhs7 _dafny.Array = _134_v80
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_134_v80), 0))
			_ = _lhs8
			(_lhs7).ArraySet1(_rhs11, (_lhs8).Int())
			_106_v55 = _rhs12
			(globalState).F2 = (((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(739))).Cmp(_111_v57.F29) == 0
			_99_v47 = (_dafny.MultiSetOf(_111_v57.F29)).Intersection((_99_v47).Difference((_99_v47).Update((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)))))
			(globalState).F2 = (_92_v41).F25()
		} else {
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))
			_ = _index16
			var _rhs13 _dafny.Int = (Companion_Default___.SafeModuloInt((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int), _111_v57.F29)).Times(_111_v57.F29)
			_ = _rhs13
			var _rhs14 _dafny.Int = (_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)
			_ = _rhs14
			var _lhs9 _dafny.Array = _107_v56
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			(_lhs9).ArraySet1(_rhs13, (_lhs10).Int())
			_lhs11.F18 = _rhs14
			(globalState).F2 = !((_92_v41).F25())
			var _135_v81 *C2
			_ = _135_v81
			var _nw25 *C2 = New_C2_()
			_ = _nw25
			_nw25.Ctor__(_91_v40, _97_v46, !((_92_v41).F25()))
			_135_v81 = _nw25
			var _136_v82 _dafny.Sequence
			_ = _136_v82
			_136_v82 = _dafny.SeqOf(_135_v81)
			var _137_v83 _dafny.Map
			_ = _137_v83
			_137_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm18(_dafny.IntOfUint32((_136_v82).Cardinality()), _111_v57.F29, globalState), _91_v40)
			var _138_v84 T1
			_ = _138_v84
			var _nw26 *C3 = New_C3_()
			_ = _nw26
			_nw26.Ctor__((func() _dafny.CodePoint {
				if (_137_v83).Contains(_111_v57.F28) {
					return (_137_v83).Get(_111_v57.F28).(_dafny.CodePoint)
				}
				return _91_v40
			})(), _97_v46, !(true))
			_138_v84 = _nw26
			var _139_v85 _dafny.Map
			_ = _139_v85
			_139_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v56, _138_v84)
			var _140_v86 _dafny.Map
			_ = _140_v86
			_140_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() T1 {
				if (_139_v85).Contains(_107_v56) {
					return (_139_v85).Get(_107_v56).(T1)
				}
				return _138_v84
			})())
			_138_v84 = (func() T1 {
				if (_140_v86).Contains(true) {
					return (_140_v86).Get(true).(T1)
				}
				return _138_v84
			})()
			var _141_v87 _dafny.Array
			_ = _141_v87
			var _nw27 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw27
			_141_v87 = _nw27
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_141_v87), 0))
			_ = _index17
			(_141_v87).ArraySet1((_92_v41).F25(), (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_141_v87), 0))
			_ = _index18
			(_141_v87).ArraySet1((_138_v84).F25(), (_index18).Int())
			var _142_v88 _dafny.Map
			_ = _142_v88
			_142_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v3, _100_v49)
			var _143_v90 _dafny.Map
			_ = _143_v90
			_143_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_138_v84).F26(), _111_v57.F29)
			var _144_v91 _dafny.MultiSet
			_ = _144_v91
			_144_v91 = _dafny.MultiSetOf((_92_v41).F25(), (_141_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_141_v87), 0))).Int()).(bool), (_92_v41).F25(), (_92_v41).F25(), (_92_v41).F25())
			_142_v88 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll8 = _dafny.NewBuilder()
				_ = _coll8
				for _iter9 := _dafny.Iterate((_41_v3).Elements()); ; {
					_compr_8, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _145_v89 _dafny.Int
					_145_v89 = interface{}(_compr_8).(_dafny.Int)
					if (_41_v3).Contains(_145_v89) {
						_coll8.Add((_145_v89).Plus((_dafny.MultiSetOf(_dafny.IntOfInt64(313))).Cardinality()))
					}
				}
				return _coll8.ToSet()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-325), (_141_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_141_v87), 0))).Int()).(bool)))).Merge((Companion_Default___.Fm28((_106_v55).Fm3((_141_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_141_v87), 0))).Int()).(bool), _dafny.IntOfInt64(819), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v57.F29, (_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int))).Cardinality(), _143_v90, globalState), (func() _dafny.Int {
				if (_144_v91).Contains((_92_v41).F25()) {
					return (_144_v91).Multiplicity((_92_v41).F25())
				}
				return _111_v57.F29
			})(), _dafny.IntOfInt64(342), globalState)).Update(_41_v3, _100_v49))
		}
		(globalState).F2 = (_92_v41).F25()
		(globalState).F22 = _38_v0
		var _146_v92 _dafny.Sequence
		_ = _146_v92
		_146_v92 = _dafny.SeqOf((_92_v41).F25())
		r0 = Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_146_v92).Cardinality())).Plus((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)), _111_v57.F29)
	}
	var _147_v93 _dafny.MultiSet
	_ = _147_v93
	_147_v93 = _dafny.MultiSetOf(_90_v39)
	r0 = (((_107_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_107_v56), 0))).Int()).(_dafny.Int)).Minus(((_147_v93).Update((_92_v41).F25(), Companion_Default___.Abs(_dafny.IntOfInt64(629)))).Cardinality())).Times((_41_v3).Cardinality())
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _148_v0 _dafny.Sequence
	_ = _148_v0
	_148_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qyvk")
	var _149_v1 _dafny.Array
	_ = _149_v1
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_3
	var _nw28 _dafny.Array
	_ = _nw28
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw28 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Set = (func(_150_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
			return func(_151_i0 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_dafny.IntOfUint32((_150_v0).Cardinality()), _dafny.IntOfUint32((_150_v0).Cardinality()))
			}
		})(_148_v0)
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
	_149_v1 = _nw28
	var _152_v2 _dafny.Array
	_ = _152_v2
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
	_ = _nw29
	_152_v2 = _nw29
	var _153_v3 bool
	_ = _153_v3
	_153_v3 = true
	var _154_v4 _dafny.Set
	_ = _154_v4
	_154_v4 = _dafny.SetOf(_153_v3)
	var _155_v5 _dafny.Sequence
	_ = _155_v5
	_155_v5 = _dafny.SeqOf((_154_v4).Cardinality())
	var _156_v6 _dafny.Array
	_ = _156_v6
	var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
	_ = _nw30
	_156_v6 = _nw30
	var _157_globalState *GlobalState
	_ = _157_globalState
	var _nw31 *GlobalState = New_GlobalState_()
	_ = _nw31
	_nw31.Ctor__(_dafny.IntOfInt64(460), _dafny.IntOfInt64(238), true, _148_v0, _149_v1, _dafny.IntOfInt64(698), false, _dafny.IntOfInt64(939), false, _152_v2, _dafny.IntOfInt64(659), false, _dafny.CodePoint('w'), false, false, _dafny.IntOfInt64(856), _dafny.IntOfInt64(966), false, _dafny.IntOfInt64(806), _155_v5, _dafny.IntOfInt64(962), _dafny.IntOfInt64(561), _dafny.IntOfInt64(182), _156_v6, _dafny.CodePoint('w'))
	_157_globalState = _nw31
	var _158_v7 _dafny.Int
	_ = _158_v7
	var _out0 _dafny.Int
	_ = _out0
	_out0 = Companion_Default___.M0(_157_globalState)
	_158_v7 = _out0
	var _159_v8 _dafny.MultiSet
	_ = _159_v8
	_159_v8 = _dafny.MultiSetOf(_148_v0)
	var _160_v9 _dafny.MultiSet
	_ = _160_v9
	_160_v9 = _dafny.MultiSetOf(_153_v3, !(!(_159_v8).Equals(_159_v8)), !(!(_153_v3)) || (Companion_Default___.Fm0(_158_v7, _158_v7, _153_v3, _157_globalState)), _153_v3, _153_v3)
	var _161_v10 _dafny.Sequence
	_ = _161_v10
	_161_v10 = _dafny.SeqOf(_153_v3, _153_v3)
	_160_v9 = Companion_Default___.Fm1(_158_v7, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_161_v10, _161_v10)).Cardinality()), _157_globalState)
	var _rhs15 _dafny.Int = (_158_v7).Times(_158_v7)
	_ = _rhs15
	var _rhs16 bool = !(_153_v3)
	_ = _rhs16
	var _rhs17 _dafny.Array = _156_v6
	_ = _rhs17
	var _lhs12 *GlobalState = _157_globalState
	_ = _lhs12
	var _lhs13 *GlobalState = _157_globalState
	_ = _lhs13
	_lhs12.F22 = _rhs15
	_lhs13.F2 = _rhs16
	_156_v6 = _rhs17
	var _162_v11 _dafny.CodePoint
	_ = _162_v11
	_162_v11 = _dafny.CodePoint('k')
	_162_v11 = _162_v11
	var _163_v12 _dafny.Map
	_ = _163_v12
	_163_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_153_v3, _158_v7, _153_v3, _162_v11, _157_globalState), _153_v3)
	var _164_v13 _dafny.Sequence
	_ = _164_v13
	_164_v13 = _dafny.SeqOf(_160_v9)
	var _165_v14 _dafny.Array
	_ = _165_v14
	var _nwElement0_5 _dafny.MultiSet = (_dafny.MultiSetOf((func() bool {
		if (_163_v12).Contains(_dafny.IntOfInt64(926)) {
			return (_163_v12).Get(_dafny.IntOfInt64(926)).(bool)
		}
		return _153_v3
	})())).Union(_160_v9)
	_ = _nwElement0_5
	var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(21))
	_ = _nw32
	(_nw32).ArraySet1(_nwElement0_5, 0)
	(_nw32).ArraySet1(_160_v9, 1)
	(_nw32).ArraySet1(_160_v9, 2)
	(_nw32).ArraySet1(_160_v9, 3)
	(_nw32).ArraySet1(Companion_Default___.Fm1(_158_v7, _158_v7, _157_globalState), 4)
	(_nw32).ArraySet1(_dafny.MultiSetOf(true), 5)
	(_nw32).ArraySet1(_160_v9, 6)
	(_nw32).ArraySet1((_164_v13).Select((Companion_Default___.SafeIndex(_158_v7, _dafny.IntOfUint32((_164_v13).Cardinality()))).Uint32()).(_dafny.MultiSet), 7)
	(_nw32).ArraySet1(Companion_Default___.Fm1(_158_v7, _158_v7, _157_globalState), 8)
	(_nw32).ArraySet1(_160_v9, 9)
	(_nw32).ArraySet1((_160_v9).Intersection(Companion_Default___.Fm1(_158_v7, _158_v7, _157_globalState)), 10)
	(_nw32).ArraySet1(_160_v9, 11)
	(_nw32).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfInt64(-362), _158_v7, _157_globalState), 12)
	(_nw32).ArraySet1(_dafny.MultiSetFromSeq(_161_v10), 13)
	(_nw32).ArraySet1(_160_v9, 14)
	(_nw32).ArraySet1(_160_v9, 15)
	(_nw32).ArraySet1((_dafny.MultiSetOf(_153_v3)).Intersection(_160_v9), 16)
	(_nw32).ArraySet1(_160_v9, 17)
	(_nw32).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_153_v3), (Companion_Default___.SafeIndex(_158_v7, _dafny.IntOfUint32((_dafny.SeqOf(_153_v3)).Cardinality()))).Uint32(), _153_v3), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_153_v3), (Companion_Default___.SafeIndex(_158_v7, _dafny.IntOfUint32((_dafny.SeqOf(_153_v3)).Cardinality()))).Uint32(), _153_v3)).Cardinality()))).Uint32(), _153_v3)), 18)
	(_nw32).ArraySet1(_160_v9, 19)
	(_nw32).ArraySet1((_160_v9).Intersection(_160_v9), 20)
	_165_v14 = _nw32
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_165_v14), 0))
	_ = _index19
	(_165_v14).ArraySet1(_160_v9, (_index19).Int())
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))
	_ = _index20
	(_156_v6).ArraySet1(true, (_index20).Int())
	var _166_v15 _dafny.Map
	_ = _166_v15
	_166_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_Default___.Fm2(_153_v3, _158_v7, !(_153_v3), _162_v11, _157_globalState)).Cmp(_dafny.IntOfUint32((_148_v0).Cardinality())) > 0), _dafny.MultiSetFromSeq(_161_v10))
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_165_v14), 0))
	_ = _index21
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))
	_ = _index22
	var _rhs18 _dafny.Int = _158_v7
	_ = _rhs18
	var _rhs19 bool = false
	_ = _rhs19
	var _rhs20 _dafny.MultiSet = (func() _dafny.MultiSet {
		if (_166_v15).Contains(_153_v3) {
			return (_166_v15).Get(_153_v3).(_dafny.MultiSet)
		}
		return _160_v9
	})()
	_ = _rhs20
	var _rhs21 bool = (_158_v7).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm2(_153_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_148_v0, (Companion_Default___.SafeIndex(_158_v7, _dafny.IntOfUint32((_148_v0).Cardinality()))).Uint32(), _162_v11)).Cardinality()), _153_v3, _162_v11, _157_globalState))).Cardinality())) >= 0
	_ = _rhs21
	var _rhs22 _dafny.Int = _158_v7
	_ = _rhs22
	var _lhs14 _dafny.Array = _165_v14
	_ = _lhs14
	var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_165_v14), 0))
	_ = _lhs15
	var _lhs16 _dafny.Array = _156_v6
	_ = _lhs16
	var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))
	_ = _lhs17
	_158_v7 = _rhs18
	_153_v3 = _rhs19
	(_lhs14).ArraySet1(_rhs20, (_lhs15).Int())
	(_lhs16).ArraySet1(_rhs21, (_lhs17).Int())
	_158_v7 = _rhs22
	var _167_v16 _dafny.Int
	_ = _167_v16
	var _out1 _dafny.Int
	_ = _out1
	_out1 = Companion_Default___.M0(_157_globalState)
	_167_v16 = _out1
	var _168_v17 _dafny.Map
	_ = _168_v17
	_168_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool))
	var _169_v18 _dafny.MultiSet
	_ = _169_v18
	_169_v18 = _dafny.MultiSetOf((_168_v17).Cardinality())
	var _170_v19 _dafny.Set
	_ = _170_v19
	_170_v19 = _dafny.SetOf(_167_v16, _158_v7)
	var _171_v20 _dafny.Array
	_ = _171_v20
	var _nwElement0_6 _dafny.Int = _167_v16
	_ = _nwElement0_6
	var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(19))
	_ = _nw33
	(_nw33).ArraySet1(_nwElement0_6, 0)
	(_nw33).ArraySet1(_167_v16, 1)
	(_nw33).ArraySet1(_167_v16, 2)
	(_nw33).ArraySet1(_167_v16, 3)
	(_nw33).ArraySet1(_dafny.IntOfUint32((_148_v0).Cardinality()), 4)
	(_nw33).ArraySet1(_167_v16, 5)
	(_nw33).ArraySet1(_dafny.IntOfInt64(757), 6)
	(_nw33).ArraySet1(_167_v16, 7)
	(_nw33).ArraySet1(_167_v16, 8)
	(_nw33).ArraySet1(Companion_Default___.Fm2((_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), _158_v7, _153_v3, _162_v11, _157_globalState), 9)
	(_nw33).ArraySet1((_169_v18).Cardinality(), 10)
	(_nw33).ArraySet1(_167_v16, 11)
	(_nw33).ArraySet1((_170_v19).Cardinality(), 12)
	(_nw33).ArraySet1(_158_v7, 13)
	(_nw33).ArraySet1(Companion_Default___.Fm2(_153_v3, _167_v16, _153_v3, _dafny.CodePoint('w'), _157_globalState), 14)
	(_nw33).ArraySet1(_167_v16, 15)
	(_nw33).ArraySet1((_dafny.Zero).Minus(_158_v7), 16)
	(_nw33).ArraySet1(_158_v7, 17)
	(_nw33).ArraySet1(Companion_Default___.Fm2((_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), (_154_v4).Cardinality(), _153_v3, _162_v11, _157_globalState), 18)
	_171_v20 = _nw33
	var _hi1 _dafny.Int = _167_v16
	_ = _hi1
	for _172_i1 := ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v5, _171_v20)).Update(_155_v5, _171_v20)).Cardinality(); _172_i1.Cmp(_hi1) < 0; _172_i1 = _172_i1.Plus(_dafny.One) {
		var _173_v21 _dafny.Array
		_ = _173_v21
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
		_ = _nw34
		_173_v21 = _nw34
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_173_v21), 0))
		_ = _index23
		(_173_v21).ArraySet1(_163_v12, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_165_v14), 0))
		_ = _index24
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_173_v21), 0))
		_ = _index25
		var _rhs23 _dafny.MultiSet = _dafny.MultiSetOf((_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool))
		_ = _rhs23
		var _rhs24 _dafny.Map = _163_v12
		_ = _rhs24
		var _lhs18 _dafny.Array = _165_v14
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_165_v14), 0))
		_ = _lhs19
		var _lhs20 _dafny.Array = _173_v21
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_173_v21), 0))
		_ = _lhs21
		(_lhs18).ArraySet1(_rhs23, (_lhs19).Int())
		(_lhs20).ArraySet1(_rhs24, (_lhs21).Int())
		_148_v0 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool) {
				return _148_v0
			}
			return _148_v0
		})(), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_172_i1, _167_v16), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool) {
				return _148_v0
			}
			return _148_v0
		})()).Cardinality()))).Uint32(), _dafny.CodePoint('n'))
		var _174_v22 _dafny.Int
		_ = _174_v22
		var _out2 _dafny.Int
		_ = _out2
		_out2 = Companion_Default___.M0(_157_globalState)
		_174_v22 = _out2
		var _175_v23 *C0
		_ = _175_v23
		var _nw35 *C0 = New_C0_()
		_ = _nw35
		_nw35.Ctor__()
		_175_v23 = _nw35
	}
	var _176_v24 _dafny.Int
	_ = _176_v24
	var _out3 _dafny.Int
	_ = _out3
	_out3 = Companion_Default___.M0(_157_globalState)
	_176_v24 = _out3
	_176_v24 = _167_v16
	(_157_globalState).F2 = (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool)
	var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_171_v20), 0))
	_ = _index26
	(_171_v20).ArraySet1(_158_v7, (_index26).Int())
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_171_v20), 0))
	_ = _index27
	(_171_v20).ArraySet1(_dafny.IntOfInt64(967), (_index27).Int())
	(_157_globalState).F5 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm2((func() bool {
		if (_168_v17).Contains(true) {
			return (_168_v17).Get(true).(bool)
		}
		return !(_153_v3)
	})(), _dafny.IntOfInt64(317), (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), _162_v11, _157_globalState), (_dafny.Zero).Minus(((func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(670), _dafny.IntOfInt64(941))); ; {
			_compr_9, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _177_v25 _dafny.Int
			_177_v25 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(670)).Cmp(_177_v25) <= 0) && ((_177_v25).Cmp(_dafny.IntOfInt64(941)) < 0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_177_v25, _dafny.IntOfInt64(913)), _153_v3)
			}
		}
		return _coll9.ToMap()
	}()).Merge(Companion_Default___.Fm26((_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), (func() _dafny.Int {
		if (_160_v9).Contains(_153_v3) {
			return (_160_v9).Multiplicity(_153_v3)
		}
		return _167_v16
	})(), _dafny.IntOfUint32((_148_v0).Cardinality()), _157_globalState))).Cardinality()))
	var _178_v26 _dafny.Array
	_ = _178_v26
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_4
	var _nw36 _dafny.Array
	_ = _nw36
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw36 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.MultiSet = (func(_179_v18 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
			return func(_180_i2 _dafny.Int) _dafny.MultiSet {
				return _179_v18
			}
		})(_169_v18)
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
	_178_v26 = _nw36
	var _181_v27 *C3
	_ = _181_v27
	var _nw37 *C3 = New_C3_()
	_ = _nw37
	_nw37.Ctor__(_162_v11, _178_v26, ((_171_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_171_v20), 0))).Int()).(_dafny.Int)).Cmp(_158_v7) == 0)
	_181_v27 = _nw37
	var _182_v28 D7
	_ = _182_v28
	_182_v28 = Companion_D7_.Create_DC23_(_176_v24, _dafny.IntOfInt64(294), _153_v3, _158_v7)
	var _183_v29 D7
	_ = _183_v29
	_183_v29 = Companion_D7_.Create_DC24_(_182_v28)
	var _184_i3 _dafny.Int
	_ = _184_i3
	_184_i3 = _dafny.Zero
	{
		var _pat_let_tv1 = _153_v3
		_ = _pat_let_tv1
		var _pat_let_tv2 = _153_v3
		_ = _pat_let_tv2
		var _pat_let_tv3 = _162_v11
		_ = _pat_let_tv3
		var _pat_let_tv4 = _156_v6
		_ = _pat_let_tv4
		var _pat_let_tv5 = _156_v6
		_ = _pat_let_tv5
		var _pat_let_tv6 = _148_v0
		_ = _pat_let_tv6
		for func(_source2 D7) bool {
			if _source2.Is_DC21() {
				var _189___mcc_h0 _dafny.Int = _source2.Get_().(D7_DC21).Cf38
				_ = _189___mcc_h0
				var _190___mcc_h1 _dafny.Map = _source2.Get_().(D7_DC21).Cf39
				_ = _190___mcc_h1
				var _191_cf39 _dafny.Map = _190___mcc_h1
				_ = _191_cf39
				var _192_cf38 _dafny.Int = _189___mcc_h0
				_ = _192_cf38
				return _pat_let_tv1
			} else if _source2.Is_DC22() {
				var _193___mcc_h2 _dafny.Int = _source2.Get_().(D7_DC22).Cf40
				_ = _193___mcc_h2
				var _194___mcc_h3 _dafny.Int = _source2.Get_().(D7_DC22).Cf41
				_ = _194___mcc_h3
				var _195___mcc_h4 bool = _source2.Get_().(D7_DC22).Cf42
				_ = _195___mcc_h4
				var _196_cf42 bool = _195___mcc_h4
				_ = _196_cf42
				var _197_cf41 _dafny.Int = _194___mcc_h3
				_ = _197_cf41
				var _198_cf40 _dafny.Int = _193___mcc_h2
				_ = _198_cf40
				return _pat_let_tv2
			} else if _source2.Is_DC23() {
				var _199___mcc_h5 _dafny.Int = _source2.Get_().(D7_DC23).Cf43
				_ = _199___mcc_h5
				var _200___mcc_h6 _dafny.Int = _source2.Get_().(D7_DC23).Cf44
				_ = _200___mcc_h6
				var _201___mcc_h7 bool = _source2.Get_().(D7_DC23).Cf45
				_ = _201___mcc_h7
				var _202___mcc_h8 _dafny.Int = _source2.Get_().(D7_DC23).Cf46
				_ = _202___mcc_h8
				var _203_cf46 _dafny.Int = _202___mcc_h8
				_ = _203_cf46
				var _204_cf45 bool = _201___mcc_h7
				_ = _204_cf45
				var _205_cf44 _dafny.Int = _200___mcc_h6
				_ = _205_cf44
				var _206_cf43 _dafny.Int = _199___mcc_h5
				_ = _206_cf43
				return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("tdneecvnd"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-460))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}((func(_207_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_208_i4 _dafny.Int) _dafny.CodePoint {
						return _207_v11
					}
				})(_pat_let_tv3))))
			} else if _source2.Is_DC20() {
				var _209___mcc_h9 _dafny.Sequence = _source2.Get_().(D7_DC20).Cf37
				_ = _209___mcc_h9
				var _210_cf37 _dafny.Sequence = _209___mcc_h9
				_ = _210_cf37
				return (_pat_let_tv5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_pat_let_tv4), 0))).Int()).(bool)
			} else {
				var _211___mcc_h10 D7 = _source2.Get_().(D7_DC24).Cf47
				_ = _211___mcc_h10
				var _212_cf47 D7 = _211___mcc_h10
				_ = _212_cf47
				return _dafny.Companion_Sequence_.IsPrefixOf(_pat_let_tv6, _dafny.UnicodeSeqOfUtf8Bytes("qarhghwh"))
			}
		}(_183_v29) {
			{
				if (_184_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_184_i3 = (_184_i3).Plus(_dafny.One)
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))
				_ = _index28
				(_156_v6).ArraySet1(_153_v3, (_index28).Int())
				var _185_v30 _dafny.Int
				_ = _185_v30
				var _out4 _dafny.Int
				_ = _out4
				_out4 = Companion_Default___.M0(_157_globalState)
				_185_v30 = _out4
				_148_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mukobu"), _148_v0)
				var _186_v31 _dafny.Set
				_ = _186_v31
				_186_v31 = _dafny.SetOf(_162_v11)
				(_181_v27).M2((func() _dafny.Set {
					if (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool) {
						return _186_v31
					}
					return _dafny.SetOf(_162_v11, _162_v11, _dafny.CodePoint('o'), _162_v11, _162_v11)
				})(), _153_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-712))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_187_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_188_i5 _dafny.Int) _dafny.Sequence {
						return _187_v0
					}
				})(_148_v0)))).Cardinality()), _157_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_156_v6), 0))); ; {
		_guard_loop_1, _ok11 := _iter11()
		if !_ok11 {
			break
		}
		var _213_i6 _dafny.Int
		_213_i6 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_213_i6).Sign() != -1) && ((_213_i6).Cmp(_dafny.ArrayLen((_156_v6), 0)) < 0)) {
			(_156_v6).ArraySet1(_153_v3, (_213_i6).Int())
		}
	}
	var _214_v32 _dafny.Sequence
	_ = _214_v32
	_214_v32 = _dafny.SeqOf(_156_v6)
	var _215_v33 D3
	_ = _215_v33
	_215_v33 = Companion_D3_.Create_DC9_(_dafny.IntOfUint32((_214_v32).Cardinality()), _162_v11, _dafny.IntOfInt64(326))
	var _216_v34 _dafny.Map
	_ = _216_v34
	_216_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_215_v33).Dtor_cf14(), _167_v16)
	if (_181_v27).Fm3(false, _176_v24, Companion_Default___.Fm2(_153_v3, _167_v16, (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), _162_v11, _157_globalState), (_216_v34).Merge(_216_v34), _157_globalState) {
		(_157_globalState).F5 = (_171_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_171_v20), 0))).Int()).(_dafny.Int)
		(_157_globalState).F22 = (func() _dafny.Int {
			if (_169_v18).Contains(_176_v24) {
				return (_169_v18).Multiplicity(_176_v24)
			}
			return _176_v24
		})()
		_214_v32 = _214_v32
		var _217_v35 _dafny.Array
		_ = _217_v35
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(24))
		_ = _nw38
		_217_v35 = _nw38
		var _218_v36 _dafny.Map
		_ = _218_v36
		_218_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v35, _dafny.Companion_Sequence_.Concatenate(_161_v10, _161_v10))
		var _rhs25 _dafny.Int = _167_v16
		_ = _rhs25
		var _rhs26 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v35, _dafny.Companion_Sequence_.Concatenate(_161_v10, _dafny.SeqOf(_153_v3)))
		_ = _rhs26
		var _rhs27 _dafny.MultiSet = (_169_v18).Intersection((_169_v18).Difference(_dafny.MultiSetFromSeq(_155_v5)))
		_ = _rhs27
		_167_v16 = _rhs25
		_218_v36 = _rhs26
		_169_v18 = _rhs27
		(_157_globalState).F2 = (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool)
	} else {
		(_157_globalState).F2 = !(_169_v18).Equals(_169_v18)
		(_157_globalState).F2 = !((_154_v4).IsDisjointFrom(_154_v4))
		_153_v3 = (_181_v27).Fm3(_153_v3, _dafny.IntOfInt64(971), (_171_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_171_v20), 0))).Int()).(_dafny.Int), _216_v34, _157_globalState)
		var _219_v37 D8
		_ = _219_v37
		_219_v37 = Companion_D8_.Create_DC26_(_171_v20, (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), (_148_v0).Select((Companion_Default___.SafeIndex((_171_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_171_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_148_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), _153_v3)
		var _220_v38 _dafny.Sequence
		_ = _220_v38
		_220_v38 = _dafny.SeqOf(_219_v37)
		_220_v38 = (func() _dafny.Sequence {
			if _153_v3 {
				return _220_v38
			}
			return (func() _dafny.Sequence {
				if (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool) {
					return _220_v38
				}
				return _220_v38
			})()
		})()
		var _221_v39 _dafny.Map
		_ = _221_v39
		_221_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_158_v7, _161_v10)
		_221_v39 = (_221_v39).Update(_176_v24, _dafny.SeqOf(Companion_Default___.Fm0(_158_v7, _dafny.IntOfInt64(512), (_156_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_156_v6), 0))).Int()).(bool), _157_globalState)))
	}
	_dafny.Print(_148_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_149_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_153_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v4).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_155_v5, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v6).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F3().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_157_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F12())
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
	_dafny.Print(_157_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_157_globalState.F19, _dafny.SeqOf(_dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142), _dafny.IntOfInt64(142))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_157_globalState).F23()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v8).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qyvk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v9).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_161_v10, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_v11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-386), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_164_v13, _dafny.SeqOf(_dafny.MultiSetOf())))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_165_v14).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_166_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_167_v16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v18).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v19).Equals(_dafny.SetOf(_dafny.IntOfInt64(-599))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v20).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_176_v24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v26).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v28).Dtor_cf43())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v28).Dtor_cf44())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v28).Dtor_cf45())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v28).Dtor_cf46())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_183_v29).Dtor_cf47()).Dtor_cf43())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_183_v29).Dtor_cf47()).Dtor_cf44())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_183_v29).Dtor_cf47()).Dtor_cf45())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_183_v29).Dtor_cf47()).Dtor_cf46())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_184_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_214_v32).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_215_v33).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_215_v33).Dtor_cf14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_215_v33).Dtor_cf15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_216_v34).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _dafny.IntOfInt64(-599))))
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
	Cf1 bool
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 bool) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(false)
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf1() bool {
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
			return ok && data1.Cf2 == data2.Cf2
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1 == data2.Cf1
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
	Cf5 bool
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Int, Cf5 bool) D2 {
	return D2{D2_DC4{Cf4, Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC5 struct {
	Cf6 _dafny.Int
	Cf7 bool
	Cf8 bool
	Cf9 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 _dafny.Int, Cf7 bool, Cf8 bool, Cf9 bool) D2 {
	return D2{D2_DC5{Cf6, Cf7, Cf8, Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
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

type D2_DC6 struct {
	Cf10 D2
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 D2) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.Zero, false)
}

func (_this D2) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) Dtor_cf5() bool {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D2_DC3).Cf3
}

func (_this D2) Dtor_cf10() D2 {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10.Equals(data2.Cf10)
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
	Cf12 bool
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 bool) D3 {
	return D3{D3_DC8{Cf12}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf13 _dafny.Int
	Cf14 _dafny.CodePoint
	Cf15 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf13 _dafny.Int, Cf14 _dafny.CodePoint, Cf15 _dafny.Int) D3 {
	return D3{D3_DC9{Cf13, Cf14, Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC7 struct {
	Cf11 _dafny.Map
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Map) D3 {
	return D3{D3_DC7{Cf11}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC10 struct {
	Cf16 D3
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 D3) D3 {
	return D3{D3_DC10{Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(false)
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf13
}

func (_this D3) Dtor_cf14() _dafny.CodePoint {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) Dtor_cf16() D3 {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf12 == data2.Cf12
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
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
	Cf18 bool
	Cf19 _dafny.Array
	Cf20 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf18 bool, Cf19 _dafny.Array, Cf20 _dafny.Int) D4 {
	return D4{D4_DC12{Cf18, Cf19, Cf20}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf21 _dafny.Array
	Cf22 bool
	Cf23 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf21 _dafny.Array, Cf22 bool, Cf23 bool) D4 {
	return D4{D4_DC13{Cf21, Cf22, Cf23}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC11 struct {
	Cf17 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D4) Dtor_cf18() bool {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() bool {
	return _this.Get_().(D4_DC13).Cf23
}

func (_this D4) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + data.Cf17.VerbatimString(true) + ")"
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
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23
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

type D5_DC15 struct {
	Cf25 bool
	Cf26 _dafny.Array
	Cf27 _dafny.Sequence
	Cf28 bool
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf25 bool, Cf26 _dafny.Array, Cf27 _dafny.Sequence, Cf28 bool) D5 {
	return D5{D5_DC15{Cf25, Cf26, Cf27, Cf28}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf24 _dafny.Map
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf24 _dafny.Map) D5 {
	return D5{D5_DC14{Cf24}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC16 struct {
	Cf29 D5
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf29 D5) D5 {
	return D5{D5_DC16{Cf29}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq, false)
}

func (_this D5) Dtor_cf25() bool {
	return _this.Get_().(D5_DC15).Cf25
}

func (_this D5) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf26
}

func (_this D5) Dtor_cf27() _dafny.Sequence {
	return _this.Get_().(D5_DC15).Cf27
}

func (_this D5) Dtor_cf28() bool {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D5_DC14).Cf24
}

func (_this D5) Dtor_cf29() D5 {
	return _this.Get_().(D5_DC16).Cf29
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + data.Cf27.VerbatimString(true) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27.Equals(data2.Cf27) && data1.Cf28 == data2.Cf28
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf24.Equals(data2.Cf24)
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
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

type D6_DC18 struct {
	Cf31 _dafny.Int
	Cf32 bool
	Cf33 bool
	Cf34 bool
	Cf35 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf31 _dafny.Int, Cf32 bool, Cf33 bool, Cf34 bool, Cf35 _dafny.Int) D6 {
	return D6{D6_DC18{Cf31, Cf32, Cf33, Cf34, Cf35}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC17 struct {
	Cf30 _dafny.Map
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf30 _dafny.Map) D6 {
	return D6{D6_DC17{Cf30}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC19 struct {
	Cf36 D6
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf36 D6) D6 {
	return D6{D6_DC19{Cf36}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_(_dafny.Zero, false, false, false, _dafny.Zero)
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC18).Cf32
}

func (_this D6) Dtor_cf33() bool {
	return _this.Get_().(D6_DC18).Cf33
}

func (_this D6) Dtor_cf34() bool {
	return _this.Get_().(D6_DC18).Cf34
}

func (_this D6) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf35
}

func (_this D6) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) Dtor_cf36() D6 {
	return _this.Get_().(D6_DC19).Cf36
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf30.Equals(data2.Cf30)
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf36.Equals(data2.Cf36)
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
	Cf38 _dafny.Int
	Cf39 _dafny.Map
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf38 _dafny.Int, Cf39 _dafny.Map) D7 {
	return D7{D7_DC21{Cf38, Cf39}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC22 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.Int
	Cf42 bool
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf40 _dafny.Int, Cf41 _dafny.Int, Cf42 bool) D7 {
	return D7{D7_DC22{Cf40, Cf41, Cf42}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC23 struct {
	Cf43 _dafny.Int
	Cf44 _dafny.Int
	Cf45 bool
	Cf46 _dafny.Int
}

func (D7_DC23) isD7() {}

func (CompanionStruct_D7_) Create_DC23_(Cf43 _dafny.Int, Cf44 _dafny.Int, Cf45 bool, Cf46 _dafny.Int) D7 {
	return D7{D7_DC23{Cf43, Cf44, Cf45, Cf46}}
}

func (_this D7) Is_DC23() bool {
	_, ok := _this.Get_().(D7_DC23)
	return ok
}

type D7_DC20 struct {
	Cf37 _dafny.Sequence
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf37 _dafny.Sequence) D7 {
	return D7{D7_DC20{Cf37}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC24 struct {
	Cf47 D7
}

func (D7_DC24) isD7() {}

func (CompanionStruct_D7_) Create_DC24_(Cf47 D7) D7 {
	return D7{D7_DC24{Cf47}}
}

func (_this D7) Is_DC24() bool {
	_, ok := _this.Get_().(D7_DC24)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC21_(_dafny.Zero, _dafny.EmptyMap)
}

func (_this D7) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D7_DC21).Cf39
}

func (_this D7) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf40
}

func (_this D7) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf41
}

func (_this D7) Dtor_cf42() bool {
	return _this.Get_().(D7_DC22).Cf42
}

func (_this D7) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf43
}

func (_this D7) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf44
}

func (_this D7) Dtor_cf45() bool {
	return _this.Get_().(D7_DC23).Cf45
}

func (_this D7) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf46
}

func (_this D7) Dtor_cf37() _dafny.Sequence {
	return _this.Get_().(D7_DC20).Cf37
}

func (_this D7) Dtor_cf47() D7 {
	return _this.Get_().(D7_DC24).Cf47
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D7_DC23:
		{
			return "D7.DC23" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D7_DC24:
		{
			return "D7.DC24" + "(" + _dafny.String(data.Cf47) + ")"
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
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39.Equals(data2.Cf39)
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42 == data2.Cf42
		}
	case D7_DC23:
		{
			data2, ok := other.Get_().(D7_DC23)
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45 == data2.Cf45 && data1.Cf46.Cmp(data2.Cf46) == 0
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D7_DC24:
		{
			data2, ok := other.Get_().(D7_DC24)
			return ok && data1.Cf47.Equals(data2.Cf47)
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
	Cf49 _dafny.Array
	Cf50 bool
	Cf51 _dafny.CodePoint
	Cf52 bool
	Cf53 bool
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf49 _dafny.Array, Cf50 bool, Cf51 _dafny.CodePoint, Cf52 bool, Cf53 bool) D8 {
	return D8{D8_DC26{Cf49, Cf50, Cf51, Cf52, Cf53}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC25 struct {
	Cf48 _dafny.MultiSet
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf48 _dafny.MultiSet) D8 {
	return D8{D8_DC25{Cf48}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC26_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.CodePoint('D'), false, false)
}

func (_this D8) Dtor_cf49() _dafny.Array {
	return _this.Get_().(D8_DC26).Cf49
}

func (_this D8) Dtor_cf50() bool {
	return _this.Get_().(D8_DC26).Cf50
}

func (_this D8) Dtor_cf51() _dafny.CodePoint {
	return _this.Get_().(D8_DC26).Cf51
}

func (_this D8) Dtor_cf52() bool {
	return _this.Get_().(D8_DC26).Cf52
}

func (_this D8) Dtor_cf53() bool {
	return _this.Get_().(D8_DC26).Cf53
}

func (_this D8) Dtor_cf48() _dafny.MultiSet {
	return _this.Get_().(D8_DC25).Cf48
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf48) + ")"
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
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50 && data1.Cf51 == data2.Cf51 && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf48.Equals(data2.Cf48)
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
	M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int)
	F25() bool
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
	M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int)
	F25() bool
	Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool
	M2(p0 _dafny.Set, p1 bool, p2 _dafny.Int, globalState *GlobalState)
	F26() _dafny.CodePoint
	F27() _dafny.Array
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
	F2   bool
	F5   _dafny.Int
	F18  _dafny.Int
	F19  _dafny.Sequence
	F21  _dafny.Int
	F22  _dafny.Int
	_f0  _dafny.Int
	_f1  _dafny.Int
	_f3  _dafny.Sequence
	_f4  _dafny.Array
	_f6  bool
	_f7  _dafny.Int
	_f8  bool
	_f9  _dafny.Array
	_f10 _dafny.Int
	_f11 bool
	_f12 _dafny.CodePoint
	_f13 bool
	_f14 bool
	_f15 _dafny.Int
	_f16 _dafny.Int
	_f17 bool
	_f20 _dafny.Int
	_f23 _dafny.Array
	_f24 _dafny.CodePoint
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this.F5 = _dafny.Zero
	_this.F18 = _dafny.Zero
	_this.F19 = _dafny.EmptySeq
	_this.F21 = _dafny.Zero
	_this.F22 = _dafny.Zero
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f3 = _dafny.EmptySeq
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f8 = false
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = _dafny.CodePoint('D')
	_this._f13 = false
	_this._f14 = false
	_this._f15 = _dafny.Zero
	_this._f16 = _dafny.Zero
	_this._f17 = false
	_this._f20 = _dafny.Zero
	_this._f23 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f24 = _dafny.CodePoint('D')
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 bool, f3 _dafny.Sequence, f4 _dafny.Array, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 bool, f9 _dafny.Array, f10 _dafny.Int, f11 bool, f12 _dafny.CodePoint, f13 bool, f14 bool, f15 _dafny.Int, f16 _dafny.Int, f17 bool, f18 _dafny.Int, f19 _dafny.Sequence, f20 _dafny.Int, f21 _dafny.Int, f22 _dafny.Int, f23 _dafny.Array, f24 _dafny.CodePoint) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this).F18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this).F21 = f21
		(_this).F22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
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
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Array {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.CodePoint {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F20() _dafny.Int {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F23() _dafny.Array {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() _dafny.CodePoint {
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__() {
	{
	}
}
func (_this *C0) Fm5(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((func() _dafny.Int {
			if true {
				return _dafny.IntOfInt64(-144)
			}
			return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(610), true)).Cardinality())
		})()).Plus(_dafny.IntOfInt64(-131)))
	}
}
func (_this *C0) Fm6(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("roa"), _dafny.UnicodeSeqOfUtf8Bytes("w"))).Cardinality()))
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f25 bool
	F28  _dafny.Sequence
	F29  _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f25 = false
	_this.F28 = _dafny.EmptySeq
	_this.F29 = _dafny.Zero
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

func (_this *C1) F25() bool {
	return _this._f25
}
func (_this *C1) Ctor__(f28 _dafny.Sequence, f29 _dafny.Int, f25 bool) {
	{
		(_this).F28 = f28
		(_this).F29 = f29
		(_this)._f25 = f25
	}
}
func (_this *C1) Fm14(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(((_dafny.MultiSetOf((_this).F25(), (_this).F25(), (_this).F25(), false)).Union(_dafny.MultiSetOf((_this).F25()))).Cardinality()), _this.F29)
	}
}
func (_this *C1) Fm15(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_222_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F29), _dafny.SeqOf((func() _dafny.Set {
				var _coll10 = _dafny.NewBuilder()
				_ = _coll10
				for _iter12 := _dafny.Iterate((_dafny.MultiSetOf(_this.F29)).Elements()); ; {
					_compr_10, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _223_v0 _dafny.Int
					_223_v0 = interface{}(_compr_10).(_dafny.Int)
					if (_dafny.MultiSetOf(_this.F29)).Contains(_223_v0) {
						_coll10.Add(Companion_Default___.SafeModuloInt(_223_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-186), (func() _dafny.Map {
							var _coll11 = _dafny.NewMapBuilder()
							_ = _coll11
							for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.CodePoint('j'))).Keys().Elements()); ; {
								_compr_11, _ok13 := _iter13()
								if !_ok13 {
									break
								}
								var _224_v1 _dafny.Sequence
								_224_v1 = interface{}(_compr_11).(_dafny.Sequence)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.CodePoint('j'))).Contains(_224_v1) {
									_coll11.Add(_224_v1, _dafny.IntOfInt64(873))
								}
							}
							return _coll11.ToMap()
						}()).Cardinality())).Cardinality()))
					}
				}
				return _coll10.ToSet()
			}()).Cardinality()))).Cardinality())
		}))
	}
}
func (_this *C1) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _225_v0 D3
		_ = _225_v0
		_225_v0 = Companion_D3_.Create_DC7_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, (_this).F25()))
		var _226_v2 _dafny.Map
		_ = _226_v2
		_226_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F25())
		var _227_v3 _dafny.Map
		_ = _227_v3
		_227_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), false)
		var _228_v4 _dafny.Array
		_ = _228_v4
		var _nwElement0_7 bool = (_this).F25()
		_ = _nwElement0_7
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(14))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_7, 0)
		(_nw39).ArraySet1((_this).F25(), 1)
		(_nw39).ArraySet1((_this).F25(), 2)
		(_nw39).ArraySet1((_this).F25(), 3)
		(_nw39).ArraySet1((_this).F25(), 4)
		(_nw39).ArraySet1(false, 5)
		(_nw39).ArraySet1((_this).F25(), 6)
		(_nw39).ArraySet1((_this).F25(), 7)
		(_nw39).ArraySet1((_this).F25(), 8)
		(_nw39).ArraySet1((p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), 9)
		(_nw39).ArraySet1((func() bool {
			if (_227_v3).Contains((_this).F25()) {
				return (_227_v3).Get((_this).F25()).(bool)
			}
			return (_this).F25()
		})(), 10)
		(_nw39).ArraySet1((_this).F25(), 11)
		(_nw39).ArraySet1((_this).F25(), 12)
		(_nw39).ArraySet1((_this).F25(), 13)
		_228_v4 = _nw39
		var _229_v5 _dafny.Map
		_ = _229_v5
		_229_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_228_v4, _226_v2)
		var _230_v7 _dafny.Map
		_ = _230_v7
		_230_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(270), p1)
		var _231_v8 _dafny.Sequence
		_ = _231_v8
		_231_v8 = _dafny.SeqOf(func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter14 := _dafny.Iterate((_230_v7).Keys().Elements()); ; {
				_compr_12, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _232_v6 _dafny.Int
				_232_v6 = interface{}(_compr_12).(_dafny.Int)
				if (_230_v7).Contains(_232_v6) {
					_coll12.Add(Companion_Default___.SafeDivisionInt(_232_v6, _this.F29), (_this).F25())
				}
			}
			return _coll12.ToMap()
		}())
		var _233_v9 _dafny.Array
		_ = _233_v9
		var _nwElement0_8 _dafny.Map = (_225_v0).Dtor_cf11()
		_ = _nwElement0_8
		var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(21))
		_ = _nw40
		(_nw40).ArraySet1(_nwElement0_8, 0)
		(_nw40).ArraySet1(func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(148), _dafny.IntOfInt64(557))); ; {
				_compr_13, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _234_v1 _dafny.Int
				_234_v1 = interface{}(_compr_13).(_dafny.Int)
				if ((_dafny.IntOfInt64(148)).Cmp(_234_v1) <= 0) && ((_234_v1).Cmp(_dafny.IntOfInt64(557)) < 0) {
					_coll13.Add((_234_v1).Minus(_this.F29), (_this).F25())
				}
			}
			return _coll13.ToMap()
		}(), 1)
		(_nw40).ArraySet1((_226_v2).Merge((func() _dafny.Map {
			if (_229_v5).Contains(_228_v4) {
				return (_229_v5).Get(_228_v4).(_dafny.Map)
			}
			return _226_v2
		})()), 2)
		(_nw40).ArraySet1(_226_v2, 3)
		(_nw40).ArraySet1(_226_v2, 4)
		(_nw40).ArraySet1((func() _dafny.Map {
			if true {
				return _226_v2
			}
			return _226_v2
		})(), 5)
		(_nw40).ArraySet1((_231_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-530), _dafny.IntOfUint32((_231_v8).Cardinality()))).Uint32()).(_dafny.Map), 6)
		(_nw40).ArraySet1(_226_v2, 7)
		(_nw40).ArraySet1(_226_v2, 8)
		(_nw40).ArraySet1(_226_v2, 9)
		(_nw40).ArraySet1(_226_v2, 10)
		(_nw40).ArraySet1(_226_v2, 11)
		(_nw40).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_227_v3).Cardinality(), (_this).F25()), 12)
		(_nw40).ArraySet1(_226_v2, 13)
		(_nw40).ArraySet1(_226_v2, 14)
		(_nw40).ArraySet1((_226_v2).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, (_this).F25())), 15)
		(_nw40).ArraySet1(_226_v2, 16)
		(_nw40).ArraySet1((_226_v2).Merge(_226_v2), 17)
		(_nw40).ArraySet1(_226_v2, 18)
		(_nw40).ArraySet1(_226_v2, 19)
		(_nw40).ArraySet1(_226_v2, 20)
		_233_v9 = _nw40
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_233_v9), 0))); ; {
			_guard_loop_2, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _235_i0 _dafny.Int
			_235_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_235_i0).Sign() != -1) && ((_235_i0).Cmp(_dafny.ArrayLen((_233_v9), 0)) < 0)) {
				(_233_v9).ArraySet1(_226_v2, (_235_i0).Int())
			}
		}
		if (_this).F25() {
			var _236_v10 _dafny.Sequence
			_ = _236_v10
			_236_v10 = _dafny.SeqOf((_dafny.IntOfUint32((_this.F28).Cardinality())).Plus(_this.F29))
			var _237_v11 _dafny.CodePoint
			_ = _237_v11
			_237_v11 = _dafny.CodePoint('f')
			var _238_v12 _dafny.MultiSet
			_ = _238_v12
			_238_v12 = _dafny.MultiSetOf(_237_v11, _237_v11)
			(_this).F29 = (_236_v10).Select((Companion_Default___.SafeIndex((_238_v12).Cardinality(), _dafny.IntOfUint32((_236_v10).Cardinality()))).Uint32()).(_dafny.Int)
			(_this).F28 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_239_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_240_i1 _dafny.Int) _dafny.CodePoint {
					return _239_v11
				}
			})(_237_v11))), _dafny.SeqOf(_237_v11)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(p1, _this.F29), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_241_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_242_i1 _dafny.Int) _dafny.CodePoint {
					return _241_v11
				}
			})(_237_v11))), _dafny.SeqOf(_237_v11))).Cardinality()))).Uint32(), _dafny.CodePoint('c'))
			(globalState).F2 = (p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool)
			var _243_v13 *C0
			_ = _243_v13
			var _nw41 *C0 = New_C0_()
			_ = _nw41
			_nw41.Ctor__()
			_243_v13 = _nw41
			_237_v11 = Companion_Default___.Fm16(_this.F29, globalState)
		} else {
			(_this).F28 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(530))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_244_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))
			var _245_v14 _dafny.CodePoint
			_ = _245_v14
			_245_v14 = _dafny.CodePoint('y')
			var _246_v15 _dafny.Set
			_ = _246_v15
			_246_v15 = _dafny.SetOf(_this.F29, (_230_v7).Cardinality())
			var _247_v16 _dafny.Sequence
			_ = _247_v16
			_247_v16 = _dafny.SeqOf(_246_v15)
			var _248_v17 _dafny.Sequence
			_ = _248_v17
			_248_v17 = _dafny.SeqOf(Companion_Default___.Fm2((_this).F25(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality()), (func() bool {
				if (_226_v2).Contains(_dafny.IntOfUint32((_this.F28).Cardinality())) {
					return (_226_v2).Get(_dafny.IntOfUint32((_this.F28).Cardinality())).(bool)
				}
				return (_this).F25()
			})(), _245_v14, globalState), ((_247_v16).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_247_v16).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _this.F29)
			(globalState).F19 = _248_v17
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_228_v4), 0))
			_ = _index29
			(_228_v4).ArraySet1((_this).F25(), (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_228_v4), 0))
			_ = _index30
			(_228_v4).ArraySet1((_this).F25(), (_index30).Int())
			var _249_v18 *C0
			_ = _249_v18
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__()
			_249_v18 = _nw42
			var _250_v19 D3
			_ = _250_v19
			_250_v19 = Companion_D3_.Create_DC9_(p1, _245_v14, _this.F29)
			var _251_v20 _dafny.Map
			_ = _251_v20
			_251_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_250_v19, (_this).F25())
			_251_v20 = (_251_v20).Update(_250_v19, (_228_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_228_v4), 0))).Int()).(bool))
		}
		var _252_v21 _dafny.Array
		_ = _252_v21
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_5
		var _nw43 _dafny.Array
		_ = _nw43
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw43 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.Int = func(_253_i3 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_253_i3, (_dafny.MultiSetOf((_this).F25(), (_this).F25())).Cardinality())
			}
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw43 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw43).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw43).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_252_v21 = _nw43
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_252_v21), 0))
		_ = _index31
		(_252_v21).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(218), p1), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_252_v21), 0))
		_ = _index32
		(_252_v21).ArraySet1((p1).Plus(_dafny.IntOfInt64(-798)), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_252_v21), 0))
		_ = _index33
		(_252_v21).ArraySet1(_this.F29, (_index33).Int())
		var _254_v22 _dafny.Array
		_ = _254_v22
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw44
		_254_v22 = _nw44
		_254_v22 = _254_v22
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_228_v4), 0))); ; {
			_guard_loop_3, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _255_i4 _dafny.Int
			_255_i4 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_255_i4).Sign() != -1) && ((_255_i4).Cmp(_dafny.ArrayLen((_228_v4), 0)) < 0)) {
				(_228_v4).ArraySet1((_this).F25(), (_255_i4).Int())
			}
		}
		r0 = (_this).F25()
		r1 = _this.F29
		return r0, r1
	}
}
func (_this *C1) M3(globalState *GlobalState) {
	{
		var _256_v0 _dafny.Array
		_ = _256_v0
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_6
		var _nw45 _dafny.Array
		_ = _nw45
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw45 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = func(_257_i0 _dafny.Int) _dafny.Int {
				return (_257_i0).Minus(_this.F29)
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw45 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw45).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw45).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_256_v0 = _nw45
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
		_ = _index34
		(_256_v0).ArraySet1(_this.F29, (_index34).Int())
		var _258_v1 _dafny.Sequence
		_ = _258_v1
		_258_v1 = _dafny.SeqOf((_this).F25(), (_this).F25())
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
		_ = _index35
		(_256_v0).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F29, (_dafny.IntOfUint32((_258_v1).Cardinality())).Plus(_this.F29)), (_index35).Int())
		var _259_v2 *C0
		_ = _259_v2
		var _nw46 *C0 = New_C0_()
		_ = _nw46
		_nw46.Ctor__()
		_259_v2 = _nw46
		var _260_v3 _dafny.MultiSet
		_ = _260_v3
		_260_v3 = _dafny.MultiSetOf(_this.F28)
		var _261_v4 _dafny.Map
		_ = _261_v4
		_261_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_this).F25())
		var _262_v5 _dafny.Array
		_ = _262_v5
		var _nwElement0_9 bool = (_this).F25()
		_ = _nwElement0_9
		var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(19))
		_ = _nw47
		(_nw47).ArraySet1(_nwElement0_9, 0)
		(_nw47).ArraySet1((_this).F25(), 1)
		(_nw47).ArraySet1((_dafny.MultiSetOf(_this.F28, _this.F28)).IsSubsetOf(_260_v3), 2)
		(_nw47).ArraySet1((_this.F29).Cmp(_dafny.IntOfInt64(43)) > 0, 3)
		(_nw47).ArraySet1(Companion_Default___.Fm0(_this.F29, _this.F29, (_this).F25(), globalState), 4)
		(_nw47).ArraySet1(true, 5)
		(_nw47).ArraySet1(Companion_Default___.Fm0((_dafny.Zero).Minus((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_258_v1).Cardinality()), !((_this).F25()), globalState), 6)
		(_nw47).ArraySet1((_this).F25(), 7)
		(_nw47).ArraySet1((_this).F25(), 8)
		(_nw47).ArraySet1((_this).F25(), 9)
		(_nw47).ArraySet1(true, 10)
		(_nw47).ArraySet1(_dafny.Companion_Sequence_.Equal(_this.F28, _this.F28), 11)
		(_nw47).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((_this).F25()), _258_v1), 12)
		(_nw47).ArraySet1((_258_v1).Select((Companion_Default___.SafeIndex((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_258_v1).Cardinality()))).Uint32()).(bool), 13)
		(_nw47).ArraySet1((_this).F25(), 14)
		(_nw47).ArraySet1((_this).F25(), 15)
		(_nw47).ArraySet1((func() bool {
			if (_261_v4).Contains(_dafny.IntOfInt64(299)) {
				return (_261_v4).Get(_dafny.IntOfInt64(299)).(bool)
			}
			return false
		})(), 16)
		(_nw47).ArraySet1((_this).F25(), 17)
		(_nw47).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfInt64(-373), _dafny.IntOfUint32((_this.F28).Cardinality()), !((_this).F25()), globalState), 18)
		_262_v5 = _nw47
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))
		_ = _index36
		(_262_v5).ArraySet1(Companion_Default___.Fm0(_this.F29, _this.F29, (_this).F25(), globalState), (_index36).Int())
		var _263_v6 _dafny.CodePoint
		_ = _263_v6
		_263_v6 = _dafny.CodePoint('i')
		var _264_v7 _dafny.Map
		_ = _264_v7
		_264_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _263_v6)
		var _265_v8 _dafny.Sequence
		_ = _265_v8
		_265_v8 = _dafny.SeqOf((_264_v7).Cardinality())
		var _266_v9 D3
		_ = _266_v9
		_266_v9 = Companion_D3_.Create_DC9_(_this.F29, _263_v6, _dafny.IntOfUint32((_265_v8).Cardinality()))
		var _pat_let_tv7 = _256_v0
		_ = _pat_let_tv7
		var _pat_let_tv8 = _256_v0
		_ = _pat_let_tv8
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))
		_ = _index37
		(_262_v5).ArraySet1(func(_source3 D3) bool {
			if _source3.Is_DC8() {
				var _267___mcc_h0 bool = _source3.Get_().(D3_DC8).Cf12
				_ = _267___mcc_h0
				var _268_cf12 bool = _267___mcc_h0
				_ = _268_cf12
				return (_this).F25()
			} else if _source3.Is_DC9() {
				var _269___mcc_h1 _dafny.Int = _source3.Get_().(D3_DC9).Cf13
				_ = _269___mcc_h1
				var _270___mcc_h2 _dafny.CodePoint = _source3.Get_().(D3_DC9).Cf14
				_ = _270___mcc_h2
				var _271___mcc_h3 _dafny.Int = _source3.Get_().(D3_DC9).Cf15
				_ = _271___mcc_h3
				var _272_cf15 _dafny.Int = _271___mcc_h3
				_ = _272_cf15
				var _273_cf14 _dafny.CodePoint = _270___mcc_h2
				_ = _273_cf14
				var _274_cf13 _dafny.Int = _269___mcc_h1
				_ = _274_cf13
				return ((_pat_let_tv8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_pat_let_tv7), 0))).Int()).(_dafny.Int)).Cmp(_272_cf15) <= 0
			} else if _source3.Is_DC7() {
				var _275___mcc_h4 _dafny.Map = _source3.Get_().(D3_DC7).Cf11
				_ = _275___mcc_h4
				var _276_cf11 _dafny.Map = _275___mcc_h4
				_ = _276_cf11
				return ((_this).F25()) && ((_this).F25())
			} else {
				var _277___mcc_h5 D3 = _source3.Get_().(D3_DC10).Cf16
				_ = _277___mcc_h5
				var _278_cf16 D3 = _277___mcc_h5
				_ = _278_cf16
				return (_this).F25()
			}
		}(_266_v9), (_index37).Int())
		(globalState).F21 = _dafny.IntOfUint32((Companion_Default___.Fm17((_this).F25(), ((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(298)), globalState)).Cardinality())
		var _279_v10 _dafny.Map
		_ = _279_v10
		_279_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(778))
		var _280_v11 _dafny.Sequence
		_ = _280_v11
		_280_v11 = _dafny.SeqOf(_279_v10, ((_279_v10).Update((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int))).Update(true, _dafny.IntOfInt64(48)))
		var _hi2 _dafny.Int = _dafny.IntOfUint32((_280_v11).Cardinality())
		_ = _hi2
		for _281_i1 := (_259_v2).Fm6(_265_v8, (_this).F25(), globalState); _281_i1.Cmp(_hi2) < 0; _281_i1 = _281_i1.Plus(_dafny.One) {
			var _282_v12 _dafny.Sequence
			_ = _282_v12
			_282_v12 = _dafny.SeqOf(_259_v2, _259_v2, _259_v2, _259_v2)
			var _283_v13 _dafny.Array
			_ = _283_v13
			var _nwElement0_10 *C0 = _259_v2
			_ = _nwElement0_10
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(17))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_10, 0)
			(_nw48).ArraySet1((func() *C0 {
				if (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool) {
					return _259_v2
				}
				return _259_v2
			})(), 1)
			(_nw48).ArraySet1(_259_v2, 2)
			(_nw48).ArraySet1(_259_v2, 3)
			(_nw48).ArraySet1(_259_v2, 4)
			(_nw48).ArraySet1(_259_v2, 5)
			(_nw48).ArraySet1(_259_v2, 6)
			(_nw48).ArraySet1(_259_v2, 7)
			(_nw48).ArraySet1(_259_v2, 8)
			(_nw48).ArraySet1(_259_v2, 9)
			(_nw48).ArraySet1(_259_v2, 10)
			(_nw48).ArraySet1(_259_v2, 11)
			(_nw48).ArraySet1(_259_v2, 12)
			(_nw48).ArraySet1((_282_v12).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_282_v12).Cardinality()))).Uint32()).(*C0), 13)
			(_nw48).ArraySet1(_259_v2, 14)
			(_nw48).ArraySet1(_259_v2, 15)
			(_nw48).ArraySet1(_259_v2, 16)
			_283_v13 = _nw48
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_283_v13), 0))
			_ = _index38
			(_283_v13).ArraySet1((_282_v12).Select((Companion_Default___.SafeIndex((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_282_v12).Cardinality()))).Uint32()).(*C0), (_index38).Int())
			var _284_v14 _dafny.Array
			_ = _284_v14
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw49
			_284_v14 = _nw49
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_284_v14), 0))
			_ = _index39
			(_284_v14).ArraySet1(_262_v5, (_index39).Int())
			var _285_v15 D1
			_ = _285_v15
			_285_v15 = Companion_D1_.Create_DC1_((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool))
			var _286_v16 _dafny.Sequence
			_ = _286_v16
			_286_v16 = _dafny.SeqOf(_285_v15)
			var _287_v17 D4
			_ = _287_v17
			_287_v17 = Companion_D4_.Create_DC11_(Companion_Default___.Fm18(_dafny.IntOfUint32((_286_v16).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool), false, (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool), (_this).F25(), (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool))).Cardinality()), globalState))
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_283_v13), 0))
			_ = _index40
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_284_v14), 0))
			_ = _index41
			var _rhs28 _dafny.Sequence = (_287_v17).Dtor_cf17()
			_ = _rhs28
			var _rhs29 _dafny.Map = (_279_v10).Merge(_279_v10)
			_ = _rhs29
			var _rhs30 *C0 = _259_v2
			_ = _rhs30
			var _rhs31 _dafny.Array = _262_v5
			_ = _rhs31
			var _lhs22 *C1 = _this
			_ = _lhs22
			var _lhs23 _dafny.Array = _283_v13
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_283_v13), 0))
			_ = _lhs24
			var _lhs25 _dafny.Array = _284_v14
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_284_v14), 0))
			_ = _lhs26
			_lhs22.F28 = _rhs28
			_279_v10 = _rhs29
			(_lhs23).ArraySet1(_rhs30, (_lhs24).Int())
			(_lhs25).ArraySet1(_rhs31, (_lhs26).Int())
			var _288_v18 _dafny.Array
			_ = _288_v18
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_7
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) D3 = func(_289_i2 _dafny.Int) D3 {
					return Companion_D3_.Create_DC8_((_this).F25())
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw50 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw50).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw50).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_288_v18 = _nw50
			var _290_v19 D3
			_ = _290_v19
			_290_v19 = Companion_D3_.Create_DC8_((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool))
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_288_v18), 0))
			_ = _index42
			(_288_v18).ArraySet1(_290_v19, (_index42).Int())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_288_v18), 0))
			_ = _index43
			var _rhs32 _dafny.Int = (_dafny.IntOfUint32((_this.F28).Cardinality())).Times(_this.F29)
			_ = _rhs32
			var _rhs33 _dafny.Array = (func() _dafny.Array {
				if (_this).F25() {
					return _256_v0
				}
				return (func() _dafny.Array {
					if (_this).F25() {
						return _256_v0
					}
					return _256_v0
				})()
			})()
			_ = _rhs33
			var _rhs34 _dafny.CodePoint = _263_v6
			_ = _rhs34
			var _rhs35 D3 = _290_v19
			_ = _rhs35
			var _lhs27 *C1 = _this
			_ = _lhs27
			var _lhs28 _dafny.Array = _288_v18
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_288_v18), 0))
			_ = _lhs29
			_lhs27.F29 = _rhs32
			_256_v0 = _rhs33
			_263_v6 = _rhs34
			(_lhs28).ArraySet1(_rhs35, (_lhs29).Int())
			(globalState).F2 = true
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
			_ = _index44
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
			_ = _index45
			var _rhs36 bool = true
			_ = _rhs36
			var _rhs37 _dafny.Int = (_281_i1).Times(_281_i1)
			_ = _rhs37
			var _rhs38 D3 = _266_v9
			_ = _rhs38
			var _rhs39 _dafny.Int = (_281_i1).Plus((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int))
			_ = _rhs39
			var _rhs40 _dafny.Int = (func() _dafny.Int {
				if (_this).F25() {
					return _this.F29
				}
				return (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)
			})()
			_ = _rhs40
			var _lhs30 *GlobalState = globalState
			_ = _lhs30
			var _lhs31 _dafny.Array = _256_v0
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _256_v0
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
			_ = _lhs34
			var _lhs35 *GlobalState = globalState
			_ = _lhs35
			_lhs30.F2 = _rhs36
			(_lhs31).ArraySet1(_rhs37, (_lhs32).Int())
			_266_v9 = _rhs38
			(_lhs33).ArraySet1(_rhs39, (_lhs34).Int())
			_lhs35.F5 = _rhs40
		}
		if Companion_Default___.Fm0((_dafny.Zero).Minus((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_this).F25(), globalState) {
			_263_v6 = _dafny.CodePoint('w')
			if (func() bool {
				if (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool) {
					return (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool)
				}
				return false
			})() {
				var _291_v20 _dafny.MultiSet
				_ = _291_v20
				_291_v20 = _dafny.MultiSetOf((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool), true)
				var _292_v21 _dafny.Map
				_ = _292_v21
				_292_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool), _291_v20)
				_292_v21 = _292_v21
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
				_ = _index46
				(_256_v0).ArraySet1((_dafny.Zero).Minus(_this.F29), (_index46).Int())
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
				_ = _index47
				(_256_v0).ArraySet1(_dafny.IntOfInt64(-356), (_index47).Int())
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))
				_ = _index48
				(_262_v5).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_this.F28, (Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32(), _263_v6), _this.F28), (_index48).Int())
				(globalState).F2 = (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool)
			} else {
				var _293_v22 _dafny.Map
				_ = _293_v22
				_293_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool), globalState), (func(_pat_let2_0 D4) D4 {
					return func(_294_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let3_0 _dafny.Sequence) D4 {
							return func(_295_dt__update_hcf17_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC11_(_295_dt__update_hcf17_h0)
							}(_pat_let3_0)
						}(_this.F28)
					}(_pat_let2_0)
				}(Companion_D4_.Create_DC11_(_this.F28))).Dtor_cf17())
				var _296_v23 _dafny.Sequence
				_ = _296_v23
				_296_v23 = _dafny.SeqOf(_this.F28, _dafny.Companion_Sequence_.Update(_this.F28, (Companion_Default___.SafeIndex((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32(), _263_v6), _this.F28, (func() _dafny.Sequence {
					if (_293_v22).Contains((_this).F25()) {
						return (_293_v22).Get((_this).F25()).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("idiijp")
				})())
				var _297_v24 _dafny.Array
				_ = _297_v24
				var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
				_ = _nw51
				_297_v24 = _nw51
				var _298_v25 _dafny.Map
				_ = _298_v25
				_298_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_296_v23).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_296_v23).Cardinality()))).Uint32()).(_dafny.Sequence), _297_v24)
				_298_v25 = (_298_v25).Update(_this.F28, _297_v24)
				(globalState).F18 = (_dafny.IntOfInt64(-494)).Times((func() _dafny.Int {
					if !(false) {
						return (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)
					}
					return (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)
				})())
				(globalState).F2 = ((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)).Cmp(_this.F29) <= 0
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))
				_ = _index49
				(_256_v0).ArraySet1((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_index49).Int())
				var _299_v26 *C0
				_ = _299_v26
				var _nw52 *C0 = New_C0_()
				_ = _nw52
				_nw52.Ctor__()
				_299_v26 = _nw52
			}
			(globalState).F2 = ((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool)) || ((_this).F25())
			(_this).F28 = _this.F28
			(globalState).F5 = _this.F29
		} else {
			var _300_v27 D4
			_ = _300_v27
			_300_v27 = Companion_D4_.Create_DC11_(_dafny.UnicodeSeqOfUtf8Bytes("qptks"))
			var _301_v28 _dafny.Map
			_ = _301_v28
			_301_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _300_v27)
			_301_v28 = (_301_v28).Update((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _300_v27)
			var _302_v29 _dafny.Set
			_ = _302_v29
			_302_v29 = _dafny.SetOf(((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)).Times((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), Companion_Default___.SafeDivisionInt(_this.F29, (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)))
			_302_v29 = (_302_v29).Union(_302_v29)
			var _303_v30 D2
			_ = _303_v30
			_303_v30 = Companion_D2_.Create_DC4_(_this.F29, (_this).F25())
			var _source4 D2 = _303_v30
			_ = _source4
			if _source4.Is_DC4() {
				var _304___mcc_h6 _dafny.Int = _source4.Get_().(D2_DC4).Cf4
				_ = _304___mcc_h6
				var _305___mcc_h7 bool = _source4.Get_().(D2_DC4).Cf5
				_ = _305___mcc_h7
				var _306_cf5 bool = _305___mcc_h7
				_ = _306_cf5
				var _307_cf4 _dafny.Int = _304___mcc_h6
				_ = _307_cf4
				var _308_v31 _dafny.Map
				_ = _308_v31
				_308_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-706))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}((func(_309_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_310_i3 _dafny.Int) _dafny.Int {
						return (_309_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_309_v0), 0))).Int()).(_dafny.Int)
					}
				})(_256_v0))), _265_v8), _306_cf5)
				_308_v31 = (_308_v31).Update(_265_v8, _306_cf5)
				var _311_v32 _dafny.Map
				_ = _311_v32
				_311_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (func() _dafny.MultiSet {
					if true {
						return _260_v3
					}
					return _260_v3
				})())
				_311_v32 = (_311_v32).Update(_dafny.IntOfUint32((_265_v8).Cardinality()), _260_v3)
				var _312_v33 _dafny.Set
				_ = _312_v33
				_312_v33 = _dafny.SetOf(true, (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool))
				var _313_v34 _dafny.Map
				_ = _313_v34
				_313_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_312_v33).Cardinality())
				var _314_v35 _dafny.Set
				_ = _314_v35
				_314_v35 = _dafny.SetOf(_dafny.CodePoint('t'))
				var _315_v36 _dafny.MultiSet
				_ = _315_v36
				_315_v36 = _dafny.MultiSetOf((_314_v35).Cardinality(), _dafny.IntOfInt64(305))
				var _316_v37 _dafny.Map
				_ = _316_v37
				_316_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
					if (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool) {
						return _302_v29
					}
					return _302_v29
				})(), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm18((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if (_313_v34).Contains(_this.F29) {
						return (_313_v34).Get(_this.F29).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_315_v36).Contains((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)) {
							return (_315_v36).Multiplicity((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int))
						}
						return _307_cf4
					})()
				})(), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-110))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}((func(_317_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_318_i4 _dafny.Int) _dafny.CodePoint {
						return _317_v6
					}
				})(_263_v6)))))
				_316_v37 = _316_v37
				var _rhs41 _dafny.Int = _307_cf4
				_ = _rhs41
				var _rhs42 bool = (_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool)
				_ = _rhs42
				var _rhs43 _dafny.Int = (_dafny.Zero).Minus(_this.F29)
				_ = _rhs43
				var _rhs44 _dafny.Int = _this.F29
				_ = _rhs44
				var _rhs45 _dafny.Array = _262_v5
				_ = _rhs45
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 *GlobalState = globalState
				_ = _lhs37
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				_lhs36.F5 = _rhs41
				_306_cf5 = _rhs42
				_lhs37.F18 = _rhs43
				_lhs38.F18 = _rhs44
				_262_v5 = _rhs45
			} else if _source4.Is_DC5() {
				var _319___mcc_h8 _dafny.Int = _source4.Get_().(D2_DC5).Cf6
				_ = _319___mcc_h8
				var _320___mcc_h9 bool = _source4.Get_().(D2_DC5).Cf7
				_ = _320___mcc_h9
				var _321___mcc_h10 bool = _source4.Get_().(D2_DC5).Cf8
				_ = _321___mcc_h10
				var _322___mcc_h11 bool = _source4.Get_().(D2_DC5).Cf9
				_ = _322___mcc_h11
				var _323_cf9 bool = _322___mcc_h11
				_ = _323_cf9
				var _324_cf8 bool = _321___mcc_h10
				_ = _324_cf8
				var _325_cf7 bool = _320___mcc_h9
				_ = _325_cf7
				var _326_cf6 _dafny.Int = _319___mcc_h8
				_ = _326_cf6
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))
				_ = _index50
				(_262_v5).ArraySet1((_326_cf6).Cmp(_dafny.IntOfInt64(275)) > 0, (_index50).Int())
				(globalState).F22 = (Companion_Default___.Fm1((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _326_cf6, globalState)).Cardinality()
				var _327_v38 _dafny.MultiSet
				_ = _327_v38
				_327_v38 = _dafny.MultiSetOf(_this.F29, _this.F29)
				var _328_v39 _dafny.Int
				_ = _328_v39
				_328_v39 = (_327_v38).Cardinality()
				var _329_v42 _dafny.Sequence
				_ = _329_v42
				_329_v42 = _dafny.SeqOf(_264_v7)
				var _330_v43 _dafny.Array
				_ = _330_v43
				var _nwElement0_11 _dafny.Map = _264_v7
				_ = _nwElement0_11
				var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(26))
				_ = _nw53
				(_nw53).ArraySet1(_nwElement0_11, 0)
				(_nw53).ArraySet1((Companion_Default___.Fm19(_dafny.IntOfUint32((_this.F28).Cardinality()), true, globalState)).Merge(_264_v7), 1)
				(_nw53).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_258_v1).Cardinality()), _dafny.CodePoint('i')), 2)
				(_nw53).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, (_this.F28).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32()).(_dafny.CodePoint)), 3)
				(_nw53).ArraySet1((_264_v7).Update((_328_v39), _dafny.CodePoint('e')), 4)
				(_nw53).ArraySet1((_264_v7).Merge(_264_v7), 5)
				(_nw53).ArraySet1(_264_v7, 6)
				(_nw53).ArraySet1(_264_v7, 7)
				(_nw53).ArraySet1(_264_v7, 8)
				(_nw53).ArraySet1((_264_v7).Update(_this.F29, _263_v6), 9)
				(_nw53).ArraySet1(Companion_Default___.Fm19(_this.F29, _325_cf7, globalState), 10)
				(_nw53).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_this.F28).Select((Companion_Default___.SafeIndex((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_this.F28).Cardinality()))).Uint32()).(_dafny.CodePoint)), 11)
				(_nw53).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _263_v6), 12)
				(_nw53).ArraySet1((_264_v7).Merge(_264_v7), 13)
				(_nw53).ArraySet1((_264_v7).Merge(_264_v7), 14)
				(_nw53).ArraySet1((_264_v7).Merge(_264_v7), 15)
				(_nw53).ArraySet1(_264_v7, 16)
				(_nw53).ArraySet1(((Companion_D5_.Create_DC14_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _263_v6))).Dtor_cf24()).Merge(_264_v7), 17)
				(_nw53).ArraySet1(func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter18 := _dafny.Iterate((_265_v8).Elements()); ; {
						_compr_14, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _331_v40 _dafny.Int
						_331_v40 = interface{}(_compr_14).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_265_v8, _331_v40) {
							_coll14.Add((_331_v40).Times(_this.F29), _263_v6)
						}
					}
					return _coll14.ToMap()
				}(), 18)
				(_nw53).ArraySet1(_264_v7, 19)
				(_nw53).ArraySet1(_264_v7, 20)
				(_nw53).ArraySet1((func() _dafny.Map {
					var _coll15 = _dafny.NewMapBuilder()
					_ = _coll15
					for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(297), _dafny.IntOfInt64(675))); ; {
						_compr_15, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _332_v41 _dafny.Int
						_332_v41 = interface{}(_compr_15).(_dafny.Int)
						if ((_dafny.IntOfInt64(297)).Cmp(_332_v41) <= 0) && ((_332_v41).Cmp(_dafny.IntOfInt64(675)) < 0) {
							_coll15.Add(Companion_Default___.SafeDivisionInt(_332_v41, _this.F29), _dafny.CodePoint('s'))
						}
					}
					return _coll15.ToMap()
				}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(402), _263_v6)), 21)
				(_nw53).ArraySet1(_264_v7, 22)
				(_nw53).ArraySet1((_329_v42).Select((Companion_Default___.SafeIndex((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_329_v42).Cardinality()))).Uint32()).(_dafny.Map), 23)
				(_nw53).ArraySet1(_264_v7, 24)
				(_nw53).ArraySet1(_264_v7, 25)
				_330_v43 = _nw53
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_330_v43), 0))
				_ = _index51
				(_330_v43).ArraySet1(_264_v7, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_330_v43), 0))
				_ = _index52
				(_330_v43).ArraySet1(((_264_v7).Merge(_264_v7)).Merge((_329_v42).Select((Companion_Default___.SafeIndex(_326_cf6, _dafny.IntOfUint32((_329_v42).Cardinality()))).Uint32()).(_dafny.Map)), (_index52).Int())
				(globalState).F22 = (Companion_Default___.SafeModuloInt((_dafny.SetOf(_326_cf6)).Cardinality(), _this.F29)).Times((_dafny.IntOfInt64(331)).Plus(_this.F29))
			} else if _source4.Is_DC3() {
				var _333___mcc_h12 _dafny.Sequence = _source4.Get_().(D2_DC3).Cf3
				_ = _333___mcc_h12
				var _334_cf3 _dafny.Sequence = _333___mcc_h12
				_ = _334_cf3
				(globalState).F18 = _this.F29
				(_this).F28 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ccueaym"), _this.F28), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(248))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_335_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_336_i5 _dafny.Int) _dafny.CodePoint {
						return _335_v6
					}
				})(_263_v6))))
				var _337_v44 D2
				_ = _337_v44
				_337_v44 = Companion_D2_.Create_DC4_(_this.F29, (_this).F25())
				var _338_v45 D2
				_ = _338_v45
				_338_v45 = Companion_D2_.Create_DC6_(_337_v44)
				_338_v45 = Companion_Default___.Fm20((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), globalState)
				(globalState).F18 = _this.F29
			} else {
				var _339___mcc_h13 D2 = _source4.Get_().(D2_DC6).Cf10
				_ = _339___mcc_h13
				var _340_cf10 D2 = _339___mcc_h13
				_ = _340_cf10
				var _341_v46 _dafny.MultiSet
				_ = _341_v46
				_341_v46 = _dafny.MultiSetOf((_262_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_262_v5), 0))).Int()).(bool))
				(globalState).F2 = Companion_Default___.Fm0((_341_v46).Cardinality(), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_this).F25(), globalState)
				_256_v0 = (func() _dafny.Array {
					if !((_this).F25()) {
						return _256_v0
					}
					return _256_v0
				})()
				(_this).F29 = (_259_v2).Fm6(_dafny.Companion_Sequence_.Concatenate(_265_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}(func(_342_i6 _dafny.Int) _dafny.Int {
					return (_dafny.MultiSetOf(_dafny.IntOfInt64(882))).Cardinality()
				}))), (_this).F25(), globalState)
				_265_v8 = (_this).Fm15(globalState)
			}
			(globalState).F2 = (_this).F25()
			var _343_v47 _dafny.Map
			_ = _343_v47
			_343_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_258_v1).Cardinality()), _this.F29)
			var _344_v50 D6
			_ = _344_v50
			_344_v50 = Companion_D6_.Create_DC17_(Companion_Default___.Fm21(globalState))
			var _345_v52 _dafny.Sequence
			_ = _345_v52
			_345_v52 = _dafny.SeqOf(Companion_Default___.Fm21(globalState), _343_v47, _343_v47)
			var _346_v53 _dafny.Array
			_ = _346_v53
			var _nwElement0_12 _dafny.Map = (_343_v47).Merge(_343_v47)
			_ = _nwElement0_12
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(24))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_12, 0)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29, _dafny.IntOfInt64(159)), 1)
			(_nw54).ArraySet1(((func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(155), _dafny.IntOfInt64(340))); ; {
					_compr_16, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _347_v48 _dafny.Int
					_347_v48 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(155)).Cmp(_347_v48) <= 0) && ((_347_v48).Cmp(_dafny.IntOfInt64(340)) < 0) {
						_coll16.Add((_347_v48).Plus((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_258_v1).Select((Companion_Default___.SafeIndex(_this.F29, _dafny.IntOfUint32((_258_v1).Cardinality()))).Uint32()).(bool), true)).Cardinality())
					}
				}
				return _coll16.ToMap()
			}()).Update(_this.F29, _this.F29)).Merge(_343_v47), 2)
			(_nw54).ArraySet1((func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(222), _dafny.IntOfInt64(365))); ; {
					_compr_17, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _348_v49 _dafny.Int
					_348_v49 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(222)).Cmp(_348_v49) <= 0) && ((_348_v49).Cmp(_dafny.IntOfInt64(365)) < 0) {
						_coll17.Add(Companion_Default___.SafeModuloInt(_348_v49, (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int))
					}
				}
				return _coll17.ToMap()
			}()).Merge(_343_v47), 3)
			(_nw54).ArraySet1(_343_v47, 4)
			(_nw54).ArraySet1(Companion_Default___.Fm21(globalState), 5)
			(_nw54).ArraySet1((_343_v47).Update((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_259_v2).Fm5(globalState)), 6)
			(_nw54).ArraySet1((_343_v47).Update((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), 7)
			(_nw54).ArraySet1((_343_v47).Update(_dafny.IntOfUint32((Companion_Default___.Fm18((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), (_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int), globalState)).Cardinality()), _this.F29), 8)
			(_nw54).ArraySet1(_343_v47, 9)
			(_nw54).ArraySet1(_343_v47, 10)
			(_nw54).ArraySet1(_343_v47, 11)
			(_nw54).ArraySet1(_343_v47, 12)
			(_nw54).ArraySet1(((_344_v50).Dtor_cf30()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F29), (_dafny.Zero).Minus(_this.F29))), 13)
			(_nw54).ArraySet1(_343_v47, 14)
			(_nw54).ArraySet1(_343_v47, 15)
			(_nw54).ArraySet1(_343_v47, 16)
			(_nw54).ArraySet1(func() _dafny.Map {
				var _coll18 = _dafny.NewMapBuilder()
				_ = _coll18
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(463), _dafny.IntOfInt64(475))); ; {
					_compr_18, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _349_v51 _dafny.Int
					_349_v51 = interface{}(_compr_18).(_dafny.Int)
					if ((_dafny.IntOfInt64(463)).Cmp(_349_v51) <= 0) && ((_349_v51).Cmp(_dafny.IntOfInt64(475)) < 0) {
						_coll18.Add((_349_v51).Times((_256_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_256_v0), 0))).Int()).(_dafny.Int)), _dafny.IntOfInt64(685))
					}
				}
				return _coll18.ToMap()
			}(), 17)
			(_nw54).ArraySet1((_345_v52).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F29), _dafny.IntOfUint32((_345_v52).Cardinality()))).Uint32()).(_dafny.Map), 18)
			(_nw54).ArraySet1((_343_v47).Merge(_343_v47), 19)
			(_nw54).ArraySet1(_343_v47, 20)
			(_nw54).ArraySet1(_343_v47, 21)
			(_nw54).ArraySet1(_343_v47, 22)
			(_nw54).ArraySet1(_343_v47, 23)
			_346_v53 = _nw54
			_346_v53 = _346_v53
		}
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f26 _dafny.CodePoint
	_f27 _dafny.Array
	_f25 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f26 = _dafny.CodePoint('D')
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f25 = false
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

func (_this *C2) F26() _dafny.CodePoint {
	return _this._f26
}
func (_this *C2) F27() _dafny.Array {
	return _this._f27
}
func (_this *C2) F25() bool {
	return _this._f25
}
func (_this *C2) Ctor__(f26 _dafny.CodePoint, f27 _dafny.Array, f25 bool) {
	{
		(_this)._f26 = f26
		(_this)._f27 = f27
		(_this)._f25 = f25
	}
}
func (_this *C2) Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return (_dafny.SetOf(_dafny.IntOfInt64(195))).IsDisjointFrom(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _dafny.IntOfInt64(602)))).Cardinality()))
	}
}
func (_this *C2) Fm8(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("w"), _dafny.UnicodeSeqOfUtf8Bytes("goffav")), _dafny.UnicodeSeqOfUtf8Bytes("atkuqwah"))
	}
}
func (_this *C2) Fm9(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(681), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), true)).Cardinality()))
	}
}
func (_this *C2) M2(p0 _dafny.Set, p1 bool, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _350_v0 _dafny.Map
		_ = _350_v0
		_350_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
		(globalState).F5 = (_350_v0).Cardinality()
		var _351_v1 _dafny.Sequence
		_ = _351_v1
		_351_v1 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _352_v2 _dafny.Sequence
		_ = _352_v2
		_352_v2 = _dafny.SeqOf(_dafny.IntOfUint32((_351_v1).Cardinality()))
		var _hi3 _dafny.Int = p2
		_ = _hi3
		for _353_i0 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2)).Cardinality()); _353_i0.Cmp(_hi3) < 0; _353_i0 = _353_i0.Plus(_dafny.One) {
			(globalState).F2 = !(true)
			var _354_v3 _dafny.Set
			_ = _354_v3
			_354_v3 = _dafny.SetOf(p1)
			var _355_v4 D2
			_ = _355_v4
			_355_v4 = Companion_D2_.Create_DC4_((_354_v3).Cardinality(), p1)
			var _356_v5 D2
			_ = _356_v5
			_356_v5 = Companion_D2_.Create_DC6_(_355_v4)
			var _source5 D2 = _356_v5
			_ = _source5
			if _source5.Is_DC4() {
				var _357___mcc_h0 _dafny.Int = _source5.Get_().(D2_DC4).Cf4
				_ = _357___mcc_h0
				var _358___mcc_h1 bool = _source5.Get_().(D2_DC4).Cf5
				_ = _358___mcc_h1
				var _359_cf5 bool = _358___mcc_h1
				_ = _359_cf5
				var _360_cf4 _dafny.Int = _357___mcc_h0
				_ = _360_cf4
				var _361_v6 _dafny.Array
				_ = _361_v6
				var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw55
				_361_v6 = _nw55
				var _362_v7 _dafny.Map
				_ = _362_v7
				_362_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_359_cf5, _361_v6)
				_362_v7 = (_362_v7).Update(_359_cf5, _361_v6)
				var _363_v8 _dafny.Map
				_ = _363_v8
				_363_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_359_cf5, (_this).F25())
				var _364_v9 _dafny.Map
				_ = _364_v9
				_364_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_v8, p1)
				_364_v9 = _364_v9
				_359_cf5 = false
				_363_v8 = _363_v8
			} else if _source5.Is_DC5() {
				var _365___mcc_h2 _dafny.Int = _source5.Get_().(D2_DC5).Cf6
				_ = _365___mcc_h2
				var _366___mcc_h3 bool = _source5.Get_().(D2_DC5).Cf7
				_ = _366___mcc_h3
				var _367___mcc_h4 bool = _source5.Get_().(D2_DC5).Cf8
				_ = _367___mcc_h4
				var _368___mcc_h5 bool = _source5.Get_().(D2_DC5).Cf9
				_ = _368___mcc_h5
				var _369_cf9 bool = _368___mcc_h5
				_ = _369_cf9
				var _370_cf8 bool = _367___mcc_h4
				_ = _370_cf8
				var _371_cf7 bool = _366___mcc_h3
				_ = _371_cf7
				var _372_cf6 _dafny.Int = _365___mcc_h2
				_ = _372_cf6
				(globalState).F2 = (func() bool {
					if _dafny.Companion_Sequence_.Contains(_351_v1, Companion_Default___.Fm10(globalState)) {
						return _369_cf9
					}
					return (_this).F25()
				})()
				var _373_v10 _dafny.Int
				_ = _373_v10
				_373_v10 = (p2).Times(p2)
				_373_v10 = _373_v10
				(globalState).F19 = _dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2)
				var _374_v11 _dafny.Array
				_ = _374_v11
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_8
				var _nw56 _dafny.Array
				_ = _nw56
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw56 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) bool = func(_375_i1 _dafny.Int) bool {
						return (_this).F25()
					}
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw56 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw56).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw56).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_374_v11 = _nw56
				var _376_v12 _dafny.Sequence
				_ = _376_v12
				_376_v12 = _dafny.SeqOf(p1, _371_cf7)
				var _377_v13 _dafny.Map
				_ = _377_v13
				_377_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_351_v1).Cardinality()))
				var _378_v14 _dafny.MultiSet
				_ = _378_v14
				_378_v14 = _dafny.MultiSetOf(_dafny.IntOfInt64(-568), _372_cf6, _353_i0, (_377_v13).Cardinality())
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_374_v11), 0))
				_ = _index53
				(_374_v11).ArraySet1((!((_376_v12).Select((Companion_Default___.SafeIndex((func() _dafny.Set {
					var _coll19 = _dafny.NewBuilder()
					_ = _coll19
					for _iter23 := _dafny.Iterate((_378_v14).Elements()); ; {
						_compr_19, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _379_v15 _dafny.Int
						_379_v15 = interface{}(_compr_19).(_dafny.Int)
						if (_378_v14).Contains(_379_v15) {
							_coll19.Add(Companion_Default___.SafeModuloInt(_379_v15, (_dafny.MultiSetOf((func() _dafny.Set {
								var _coll20 = _dafny.NewBuilder()
								_ = _coll20
								for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(308), _dafny.IntOfInt64(305))); ; {
									_compr_20, _ok24 := _iter24()
									if !_ok24 {
										break
									}
									var _380_v16 _dafny.Int
									_380_v16 = interface{}(_compr_20).(_dafny.Int)
									if ((_dafny.IntOfInt64(308)).Cmp(_380_v16) <= 0) && ((_380_v16).Cmp(_dafny.IntOfInt64(305)) < 0) {
										_coll20.Add(Companion_Default___.SafeModuloInt(_380_v16, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-982))).Cardinality()))).Cardinality()))
									}
								}
								return _coll20.ToSet()
							}()).Cardinality())).Cardinality()))
						}
					}
					return _coll19.ToSet()
				}()).Cardinality(), _dafny.IntOfUint32((_376_v12).Cardinality()))).Uint32()).(bool))) == ((_this).F25()), (_index53).Int())
				var _381_v17 _dafny.Array
				_ = _381_v17
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_9
				var _nw57 _dafny.Array
				_ = _nw57
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw57 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Set = (func(_382_v3 _dafny.Set) func(_dafny.Int) _dafny.Set {
						return func(_383_i2 _dafny.Int) _dafny.Set {
							return _382_v3
						}
					})(_354_v3)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw57 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw57).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw57).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_381_v17 = _nw57
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_374_v11), 0))
				_ = _index54
				var _rhs46 bool = p1
				_ = _rhs46
				var _rhs47 bool = true
				_ = _rhs47
				var _rhs48 _dafny.Array = _381_v17
				_ = _rhs48
				var _lhs39 _dafny.Array = _374_v11
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_374_v11), 0))
				_ = _lhs40
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				(_lhs39).ArraySet1(_rhs46, (_lhs40).Int())
				_lhs41.F2 = _rhs47
				_381_v17 = _rhs48
			} else if _source5.Is_DC3() {
				var _384___mcc_h6 _dafny.Sequence = _source5.Get_().(D2_DC3).Cf3
				_ = _384___mcc_h6
				var _385_cf3 _dafny.Sequence = _384___mcc_h6
				_ = _385_cf3
				var _386_v18 *C0
				_ = _386_v18
				var _nw58 *C0 = New_C0_()
				_ = _nw58
				_nw58.Ctor__()
				_386_v18 = _nw58
				_386_v18 = (func() *C0 {
					if p1 {
						return _386_v18
					}
					return _386_v18
				})()
				(globalState).F22 = (func() _dafny.Int {
					if p1 {
						return _353_i0
					}
					return _353_i0
				})()
				var _387_v19 _dafny.Int
				_ = _387_v19
				var _out5 _dafny.Int
				_ = _out5
				_out5 = Companion_Default___.M0(globalState)
				_387_v19 = _out5
				(globalState).F2 = !((_this).F25())
			} else {
				var _388___mcc_h7 D2 = _source5.Get_().(D2_DC6).Cf10
				_ = _388___mcc_h7
				var _389_cf10 D2 = _388___mcc_h7
				_ = _389_cf10
				var _390_v20 _dafny.Map
				_ = _390_v20
				_390_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(685))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}((func(_391_v0 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_392_i3 _dafny.Int) _dafny.Map {
						return _391_v0
					}
				})(_350_v0))), (func() bool {
					if p1 {
						return (_this).F25()
					}
					return (_this).F25()
				})())
				var _393_v21 _dafny.Sequence
				_ = _393_v21
				_393_v21 = _dafny.SeqOf(_350_v0, _350_v0)
				var _394_v22 D2
				_ = _394_v22
				_394_v22 = Companion_D2_.Create_DC5_(_353_i0, p1, (_this).F25(), (_this).F25())
				_390_v20 = (_390_v20).Update(_393_v21, (_394_v22).Dtor_cf9())
				(globalState).F22 = _353_i0
				(globalState).F2 = p1
				var _395_v23 _dafny.Array
				_ = _395_v23
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_10
				var _nw59 _dafny.Array
				_ = _nw59
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw59 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = (func(_396_v3 _dafny.Set) func(_dafny.Int) bool {
						return func(_397_i4 _dafny.Int) bool {
							return (_dafny.IntOfInt64(444)).Cmp((_396_v3).Cardinality()) > 0
						}
					})(_354_v3)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw59 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw59).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw59).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_395_v23 = _nw59
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_395_v23), 0))
				_ = _index55
				(_395_v23).ArraySet1((_this).F25(), (_index55).Int())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_395_v23), 0))
				_ = _index56
				(_395_v23).ArraySet1((_this).F25(), (_index56).Int())
			}
			var _398_v24 _dafny.CodePoint
			_ = _398_v24
			_398_v24 = _dafny.CodePoint('y')
			_398_v24 = Companion_Default___.Fm10(globalState)
			var _399_v25 _dafny.Array
			_ = _399_v25
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw60
			_399_v25 = _nw60
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw61
			_399_v25 = _nw61
		}
		var _400_i5 _dafny.Int
		_ = _400_i5
		_400_i5 = _dafny.Zero
		{
			for (_this).F25() {
				{
					if (_400_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_400_i5 = (_400_i5).Plus(_dafny.One)
					var _401_v26 D2
					_ = _401_v26
					_401_v26 = Companion_D2_.Create_DC4_(p2, p1)
					var _402_v27 _dafny.Sequence
					_ = _402_v27
					_402_v27 = _dafny.SeqOf(_401_v26, Companion_Default___.Fm11(globalState), _401_v26)
					var _403_v28 _dafny.Map
					_ = _403_v28
					_403_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), p2)
					var _404_v29 _dafny.Array
					_ = _404_v29
					var _nwElement0_13 bool = p1
					_ = _nwElement0_13
					var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(20))
					_ = _nw62
					(_nw62).ArraySet1(_nwElement0_13, 0)
					(_nw62).ArraySet1((_this).F25(), 1)
					(_nw62).ArraySet1(p1, 2)
					(_nw62).ArraySet1((_this).F25(), 3)
					(_nw62).ArraySet1(!(p1), 4)
					(_nw62).ArraySet1(p1, 5)
					(_nw62).ArraySet1((_this).F25(), 6)
					(_nw62).ArraySet1(!(p1) || ((_this).F25()), 7)
					(_nw62).ArraySet1(!(p1) || (p1), 8)
					(_nw62).ArraySet1((_this).F25(), 9)
					(_nw62).ArraySet1((_this).Fm3((_this).F25(), p2, _dafny.IntOfUint32((_351_v1).Cardinality()), _403_v28, globalState), 10)
					(_nw62).ArraySet1(!(false), 11)
					(_nw62).ArraySet1((_this).Fm3(p1, p2, p2, _403_v28, globalState), 12)
					(_nw62).ArraySet1((_this).F25(), 13)
					(_nw62).ArraySet1((_this).F25(), 14)
					(_nw62).ArraySet1(p1, 15)
					(_nw62).ArraySet1(!(true), 16)
					(_nw62).ArraySet1(!((_this).F25()), 17)
					(_nw62).ArraySet1((_this).F25(), 18)
					(_nw62).ArraySet1(true, 19)
					_404_v29 = _nw62
					var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))
					_ = _index57
					(_404_v29).ArraySet1(p1, (_index57).Int())
					var _405_v30 _dafny.Map
					_ = _405_v30
					_405_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
					var _406_v31 _dafny.Set
					_ = _406_v31
					_406_v31 = _dafny.SetOf(p1)
					var _407_v32 _dafny.MultiSet
					_ = _407_v32
					_407_v32 = _dafny.MultiSetOf(p2, p2, Companion_Default___.SafeModuloInt((_406_v31).Cardinality(), _dafny.IntOfInt64(502)), Companion_Default___.SafeDivisionInt(p2, p2))
					var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))
					_ = _index58
					var _rhs49 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_402_v27, _dafny.Companion_Sequence_.Concatenate(_402_v27, _402_v27))
					_ = _rhs49
					var _rhs50 _dafny.Int = _dafny.IntOfInt64(387)
					_ = _rhs50
					var _rhs51 _dafny.Int = (_407_v32).Cardinality()
					_ = _rhs51
					var _rhs52 bool = (func() bool {
						if true {
							return (_this).F25()
						}
						return false
					})()
					_ = _rhs52
					var _rhs53 _dafny.Map = ((_405_v30).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p2)).Merge(_405_v30))
					_ = _rhs53
					var _lhs42 *GlobalState = globalState
					_ = _lhs42
					var _lhs43 *GlobalState = globalState
					_ = _lhs43
					var _lhs44 _dafny.Array = _404_v29
					_ = _lhs44
					var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))
					_ = _lhs45
					_402_v27 = _rhs49
					_lhs42.F22 = _rhs50
					_lhs43.F18 = _rhs51
					(_lhs44).ArraySet1(_rhs52, (_lhs45).Int())
					_405_v30 = _rhs53
					var _408_v33 _dafny.Array
					_ = _408_v33
					var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
					_ = _nw63
					_408_v33 = _nw63
					_408_v33 = _408_v33
					var _rhs54 _dafny.Array = (func() _dafny.Array {
						if p1 {
							return _404_v29
						}
						return _404_v29
					})()
					_ = _rhs54
					var _rhs55 bool = !(Companion_Default___.Fm0(Companion_Default___.Fm2((_404_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))).Int()).(bool), _dafny.IntOfInt64(994), !(p1), (_this).F26(), globalState), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-823), p2), _dafny.Companion_Sequence_.IsPrefixOf(_351_v1, _351_v1), globalState))
					_ = _rhs55
					var _lhs46 *GlobalState = globalState
					_ = _lhs46
					_404_v29 = _rhs54
					_lhs46.F2 = _rhs55
					if p1 {
						(globalState).F2 = (_this).F25()
						var _409_v34 _dafny.Array
						_ = _409_v34
						var _len0_11 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_11
						var _nw64 _dafny.Array
						_ = _nw64
						if _len0_11.Cmp(_dafny.Zero) == 0 {
							_nw64 = _dafny.NewArray(_len0_11)
						} else {
							var _init11 func(_dafny.Int) _dafny.Int = func(_410_i6 _dafny.Int) _dafny.Int {
								return (_410_i6).Times(_dafny.IntOfInt64(-782))
							}
							_ = _init11
							var _element0_11 = _init11(_dafny.Zero)
							_ = _element0_11
							_nw64 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
							(_nw64).ArraySet1(_element0_11, 0)
							var _nativeLen0_11 = (_len0_11).Int()
							_ = _nativeLen0_11
							for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
								(_nw64).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
							}
						}
						_409_v34 = _nw64
						var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_409_v34), 0))
						_ = _index59
						(_409_v34).ArraySet1(p2, (_index59).Int())
						var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_409_v34), 0))
						_ = _index60
						(_409_v34).ArraySet1((Companion_Default___.SafeModuloInt(p2, p2)).Times(p2), (_index60).Int())
						var _411_v35 _dafny.Sequence
						_ = _411_v35
						_411_v35 = _dafny.SeqOf((_404_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))).Int()).(bool))
						var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_409_v34), 0))
						_ = _index61
						(_409_v34).ArraySet1(((_409_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_409_v34), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, p1, (_404_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))).Int()).(bool)), _dafny.SeqOf((_411_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_411_v35).Cardinality()), _dafny.IntOfUint32((_411_v35).Cardinality()))).Uint32()).(bool)))).Cardinality())))), (_index61).Int())
						var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))
						_ = _index62
						(_404_v29).ArraySet1(false, (_index62).Int())
						var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))
						_ = _index63
						(_404_v29).ArraySet1((_this).F25(), (_index63).Int())
					} else {
						(globalState).F2 = (_dafny.IntOfInt64(-36)).Cmp(p2) <= 0
						var _412_v36 _dafny.Array
						_ = _412_v36
						var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
						_ = _nw65
						_412_v36 = _nw65
						var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_412_v36), 0))
						_ = _index64
						(_412_v36).ArraySet1(p2, (_index64).Int())
						var _413_v37 _dafny.Sequence
						_ = _413_v37
						_413_v37 = _dafny.SeqOf(_404_v29, _404_v29)
						var _414_v38 _dafny.MultiSet
						_ = _414_v38
						_414_v38 = _dafny.MultiSetOf((_this).F25())
						var _415_v39 _dafny.MultiSet
						_ = _415_v39
						_415_v39 = _dafny.MultiSetOf(_404_v29, (func() _dafny.Array {
							if (_this).F25() {
								return _404_v29
							}
							return _404_v29
						})(), _404_v29, (_413_v37).Select((Companion_Default___.SafeIndex((_414_v38).Cardinality(), _dafny.IntOfUint32((_413_v37).Cardinality()))).Uint32()).(_dafny.Array))
						var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_412_v36), 0))
						_ = _index65
						var _rhs56 _dafny.Int = (_415_v39).Cardinality()
						_ = _rhs56
						var _rhs57 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(811), _dafny.IntOfInt64(345))).Minus(Companion_Default___.SafeDivisionInt((_414_v38).Cardinality(), _dafny.IntOfUint32((_351_v1).Cardinality())))
						_ = _rhs57
						var _rhs58 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(884), Companion_Default___.SafeModuloInt(p2, p2))
						_ = _rhs58
						var _lhs47 *GlobalState = globalState
						_ = _lhs47
						var _lhs48 *GlobalState = globalState
						_ = _lhs48
						var _lhs49 _dafny.Array = _412_v36
						_ = _lhs49
						var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_412_v36), 0))
						_ = _lhs50
						_lhs47.F18 = _rhs56
						_lhs48.F5 = _rhs57
						(_lhs49).ArraySet1(_rhs58, (_lhs50).Int())
						(globalState).F22 = p2
						(globalState).F18 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p2), p2))
						var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_404_v29), 0))
						_ = _index66
						(_404_v29).ArraySet1((_this).F25(), (_index66).Int())
					}
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _416_v40 _dafny.Array
		_ = _416_v40
		var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw66
		_416_v40 = _nw66
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))
		_ = _index67
		(_416_v40).ArraySet1(!((func() bool {
			if (_350_v0).Contains(p2) {
				return (_350_v0).Get(p2).(bool)
			}
			return !(false)
		})()), (_index67).Int())
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))
		_ = _index68
		(_416_v40).ArraySet1((_this).F25(), (_index68).Int())
		var _417_v41 _dafny.Array
		_ = _417_v41
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_12
		var _nw67 _dafny.Array
		_ = _nw67
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw67 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Map = (func(_418_p1 bool, _419_p2 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_420_i7 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_418_p1, _dafny.MultiSetOf(_419_p2, _419_p2))
				}
			})(p1, p2)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw67 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw67).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw67).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_417_v41 = _nw67
		var _421_v42 _dafny.MultiSet
		_ = _421_v42
		_421_v42 = _dafny.MultiSetOf((_this).F25())
		var _422_v43 _dafny.MultiSet
		_ = _422_v43
		_422_v43 = _dafny.MultiSetOf((_421_v42).Cardinality())
		var _423_v44 _dafny.Map
		_ = _423_v44
		_423_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _422_v43)
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_417_v41), 0))
		_ = _index69
		(_417_v41).ArraySet1(_423_v44, (_index69).Int())
		var _424_v45 D2
		_ = _424_v45
		_424_v45 = Companion_D2_.Create_DC4_(p2, !((_this).F25()))
		var _425_v46 *C0
		_ = _425_v46
		var _nw68 *C0 = New_C0_()
		_ = _nw68
		_nw68.Ctor__()
		_425_v46 = _nw68
		var _426_v47 _dafny.Sequence
		_ = _426_v47
		_426_v47 = _dafny.SeqOf(_425_v46, _425_v46)
		var _427_v48 _dafny.Set
		_ = _427_v48
		_427_v48 = _dafny.SetOf((_426_v47).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_426_v47).Cardinality()))).Uint32()).(*C0))
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_417_v41), 0))
		_ = _index70
		var _rhs59 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p2, p2, (_427_v48).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2))
		_ = _rhs59
		var _rhs60 _dafny.Map = (_423_v44).Merge(_423_v44)
		_ = _rhs60
		var _rhs61 D2 = Companion_D2_.Create_DC4_((p2).Minus(_dafny.IntOfInt64(-845)), p1)
		_ = _rhs61
		var _lhs51 _dafny.Array = _417_v41
		_ = _lhs51
		var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen((_417_v41), 0))
		_ = _lhs52
		_352_v2 = _rhs59
		(_lhs51).ArraySet1(_rhs60, (_lhs52).Int())
		_424_v45 = _rhs61
		var _428_v49 _dafny.Sequence
		_ = _428_v49
		_428_v49 = _dafny.SeqOf(_352_v2)
		var _hi4 _dafny.Int = (_421_v42).Cardinality()
		_ = _hi4
		for _429_i8 := _dafny.IntOfUint32((_428_v49).Cardinality()); _429_i8.Cmp(_hi4) < 0; _429_i8 = _429_i8.Plus(_dafny.One) {
			_351_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_351_v1, _351_v1), _dafny.Companion_Sequence_.Concatenate(_351_v1, _351_v1))
			(globalState).F5 = _429_i8
			var _430_v50 _dafny.Map
			_ = _430_v50
			_430_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_416_v40, _429_i8)
			var _431_v51 _dafny.Set
			_ = _431_v51
			_431_v51 = _dafny.SetOf((_430_v50).Cardinality())
			if !(_431_v51).Contains(_429_i8) {
				var _432_v52 D1
				_ = _432_v52
				_432_v52 = Companion_D1_.Create_DC2_(p1)
				var _rhs62 _dafny.Array = _416_v40
				_ = _rhs62
				var _rhs63 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(572)).Times(_dafny.IntOfInt64(896)), _429_i8))
				_ = _rhs63
				var _rhs64 bool = (_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool)
				_ = _rhs64
				var _rhs65 D1 = _432_v52
				_ = _rhs65
				var _rhs66 _dafny.Int = (_425_v46).Fm5(globalState)
				_ = _rhs66
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				var _lhs55 *GlobalState = globalState
				_ = _lhs55
				_416_v40 = _rhs62
				_lhs53.F5 = _rhs63
				_lhs54.F2 = _rhs64
				_432_v52 = _rhs65
				_lhs55.F22 = _rhs66
				(globalState).F18 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_352_v2, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_352_v2).Cardinality()))).Uint32(), _dafny.IntOfUint32((_352_v2).Cardinality()))).Cardinality())
				var _433_v53 _dafny.Map
				_ = _433_v53
				_433_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_416_v40, (_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool))
				var _434_v54 _dafny.Array
				_ = _434_v54
				var _nwElement0_14 _dafny.Map = _433_v53
				_ = _nwElement0_14
				var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(3))
				_ = _nw69
				(_nw69).ArraySet1(_nwElement0_14, 0)
				(_nw69).ArraySet1((_433_v53).Update(_416_v40, (_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool)), 1)
				(_nw69).ArraySet1(_433_v53, 2)
				_434_v54 = _nw69
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_434_v54), 0))
				_ = _index71
				(_434_v54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_416_v40, true), (_index71).Int())
				var _435_v55 _dafny.Map
				_ = _435_v55
				_435_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool), _433_v53)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_434_v54), 0))
				_ = _index72
				(_434_v54).ArraySet1(((func() _dafny.Map {
					if (_435_v55).Contains((_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool)) {
						return (_435_v55).Get((_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool)).(_dafny.Map)
					}
					return _433_v53
				})()).Merge(_433_v53), (_index72).Int())
				var _436_v56 _dafny.Map
				_ = _436_v56
				_436_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v43, (p2).Plus(_dafny.IntOfUint32((_351_v1).Cardinality())))
				_436_v56 = (func() _dafny.Map {
					if p1 {
						return _436_v56
					}
					return _436_v56
				})()
				var _437_v57 *C0
				_ = _437_v57
				var _nw70 *C0 = New_C0_()
				_ = _nw70
				_nw70.Ctor__()
				_437_v57 = _nw70
			} else {
				(globalState).F21 = _dafny.IntOfUint32((_351_v1).Cardinality())
				(globalState).F5 = _dafny.IntOfInt64(-32)
				var _438_v58 _dafny.Array
				_ = _438_v58
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_13
				var _nw71 _dafny.Array
				_ = _nw71
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw71 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.MultiSet = (func(_439_v1 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
						return func(_440_i9 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_439_v1, _439_v1)).Union(_dafny.MultiSetOf(_439_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(441))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg21 _dafny.Int) interface{} {
									return coer21(arg21)
								}
							}(func(_441_i10 _dafny.Int) _dafny.CodePoint {
								return (_this).F26()
							}))))
						}
					})(_351_v1)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw71 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw71).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw71).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_438_v58 = _nw71
				var _442_v59 _dafny.MultiSet
				_ = _442_v59
				_442_v59 = _dafny.MultiSetOf(_351_v1, _351_v1)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_438_v58), 0))
				_ = _index73
				(_438_v58).ArraySet1(_442_v59, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_438_v58), 0))
				_ = _index74
				(_438_v58).ArraySet1((_442_v59).Update(_351_v1, Companion_Default___.Abs(p2)), (_index74).Int())
				var _443_v60 _dafny.Array
				_ = _443_v60
				var _nwElement0_15 _dafny.Sequence = _352_v2
				_ = _nwElement0_15
				var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(23))
				_ = _nw72
				(_nw72).ArraySet1(_nwElement0_15, 0)
				(_nw72).ArraySet1(_352_v2, 1)
				(_nw72).ArraySet1(_352_v2, 2)
				(_nw72).ArraySet1(_352_v2, 3)
				(_nw72).ArraySet1(_352_v2, 4)
				(_nw72).ArraySet1(_352_v2, 5)
				(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2), 6)
				(_nw72).ArraySet1((func() _dafny.Sequence {
					if Companion_Default___.Fm0(_429_i8, _429_i8, p1, globalState) {
						return _352_v2
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-463))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg22 _dafny.Int) interface{} {
							return coer22(arg22)
						}
					}((func(_444_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_445_i11 _dafny.Int) _dafny.Int {
							return _444_p2
						}
					})(p2)))
				})(), 7)
				(_nw72).ArraySet1((_428_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_351_v1).Cardinality()), _dafny.IntOfUint32((_428_v49).Cardinality()))).Uint32()).(_dafny.Sequence), 8)
				(_nw72).ArraySet1(Companion_Default___.Fm12(p2, (_this).F25(), (_350_v0).Update(_429_i8, (_416_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))).Int()).(bool)), p2, globalState), 9)
				(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2), 10)
				(_nw72).ArraySet1(_352_v2, 11)
				(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_429_i8, p2), _352_v2), 12)
				(_nw72).ArraySet1(_352_v2, 13)
				(_nw72).ArraySet1(_dafny.SeqOf((_dafny.MultiSetFromSeq(_428_v49)).Cardinality(), p2, (_dafny.Zero).Minus(p2)), 14)
				(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p2), _352_v2), 15)
				(_nw72).ArraySet1(_352_v2, 16)
				(_nw72).ArraySet1(_dafny.SeqOf((Companion_Default___.Fm13(p1, p2, globalState)).Cardinality(), _dafny.IntOfInt64(100), p2, p2), 17)
				(_nw72).ArraySet1(_352_v2, 18)
				(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_352_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-853))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}((func(_446_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_447_i12 _dafny.Int) _dafny.Int {
						return _446_i8
					}
				})(_429_i8)))), 19)
				(_nw72).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(237))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_448_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_449_i13 _dafny.Int) _dafny.Int {
						return _448_p2
					}
				})(p2))), 20)
				(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2), 21)
				(_nw72).ArraySet1(Companion_Default___.Fm12(p2, !((_this).F25()), _350_v0, p2, globalState), 22)
				_443_v60 = _nw72
				var _450_v61 _dafny.Map
				_ = _450_v61
				_450_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf(p2))
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_443_v60), 0))
				_ = _index75
				(_443_v60).ArraySet1((func() _dafny.Sequence {
					if (_this).F25() {
						return (func() _dafny.Sequence {
							if (_450_v61).Contains(_429_i8) {
								return (_450_v61).Get(_429_i8).(_dafny.Sequence)
							}
							return _352_v2
						})()
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}((func(_451_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_452_i14 _dafny.Int) _dafny.Int {
							return _451_i8
						}
					})(_429_i8)))
				})(), (_index75).Int())
				var _453_v62 _dafny.Map
				_ = _453_v62
				_453_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_421_v42).Cardinality(), _dafny.IntOfInt64(669))
				var _454_v63 _dafny.Sequence
				_ = _454_v63
				_454_v63 = _dafny.SeqOf(_453_v62, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(41), p2))
				var _455_v64 _dafny.Map
				_ = _455_v64
				_455_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _429_i8)
				var _456_v65 D2
				_ = _456_v65
				_456_v65 = Companion_D2_.Create_DC5_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(Companion_D2_.Create_DC3_(_454_v63)), p2)).Cardinality(), (_this).Fm3(p1, _429_i8, p2, _455_v64, globalState), p1, (_this).F25())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_443_v60), 0))
				_ = _index76
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))
				_ = _index77
				var _rhs67 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_352_v2, _352_v2), _dafny.SeqOf(_dafny.IntOfUint32((_351_v1).Cardinality()), _429_i8))
				_ = _rhs67
				var _rhs68 _dafny.Int = Companion_Default___.Fm2(p1, ((_456_v65).Dtor_cf6()).Minus(p2), !((_this).F25()), (_this).F26(), globalState)
				_ = _rhs68
				var _rhs69 _dafny.Int = _429_i8
				_ = _rhs69
				var _rhs70 bool = (func() bool {
					if (_this).F25() {
						return (_this).F25()
					}
					return p1
				})()
				_ = _rhs70
				var _lhs56 _dafny.Array = _443_v60
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_443_v60), 0))
				_ = _lhs57
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				var _lhs59 *GlobalState = globalState
				_ = _lhs59
				var _lhs60 _dafny.Array = _416_v40
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))
				_ = _lhs61
				(_lhs56).ArraySet1(_rhs67, (_lhs57).Int())
				_lhs58.F5 = _rhs68
				_lhs59.F22 = _rhs69
				(_lhs60).ArraySet1(_rhs70, (_lhs61).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_416_v40), 0))
				_ = _index78
				(_416_v40).ArraySet1((p1) == ((_this).F25()), (_index78).Int())
			}
			var _457_v66 _dafny.Int
			_ = _457_v66
			var _out6 _dafny.Int
			_ = _out6
			_out6 = Companion_Default___.M0(globalState)
			_457_v66 = _out6
		}
	}
}
func (_this *C2) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		if !(!(false)) {
			(globalState).F21 = p1
			var _458_v0 *C0
			_ = _458_v0
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__()
			_458_v0 = _nw73
			var _459_v1 _dafny.Int
			_ = _459_v1
			var _out7 _dafny.Int
			_ = _out7
			_out7 = Companion_Default___.M0(globalState)
			_459_v1 = _out7
			var _rhs71 bool = (_this).F25()
			_ = _rhs71
			var _rhs72 _dafny.Int = Companion_Default___.SafeModuloInt((_458_v0).Fm5(globalState), Companion_Default___.SafeModuloInt(p1, _459_v1))
			_ = _rhs72
			var _lhs62 *GlobalState = globalState
			_ = _lhs62
			var _lhs63 *GlobalState = globalState
			_ = _lhs63
			_lhs62.F2 = _rhs71
			_lhs63.F5 = _rhs72
			if (_this).F25() {
				var _460_v2 _dafny.Map
				_ = _460_v2
				_460_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _461_v3 _dafny.Sequence
				_ = _461_v3
				_461_v3 = _dafny.SeqOf(_460_v2, _460_v2)
				var _462_v4 D2
				_ = _462_v4
				_462_v4 = Companion_D2_.Create_DC3_(_461_v3)
				var _463_v5 _dafny.Map
				_ = _463_v5
				_463_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v4, p1)
				_463_v5 = ((_463_v5).Merge(_463_v5)).Merge(_463_v5)
				var _464_v6 _dafny.Map
				_ = _464_v6
				_464_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _459_v1)
				var _465_v7 _dafny.Array
				_ = _465_v7
				var _nwElement0_16 bool = true
				_ = _nwElement0_16
				var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(16))
				_ = _nw74
				(_nw74).ArraySet1(_nwElement0_16, 0)
				(_nw74).ArraySet1(true, 1)
				(_nw74).ArraySet1(false, 2)
				(_nw74).ArraySet1(!((_this).F25()), 3)
				(_nw74).ArraySet1((_this).F25(), 4)
				(_nw74).ArraySet1((_this).F25(), 5)
				(_nw74).ArraySet1((_this).Fm3((_this).F25(), p1, p1, _464_v6, globalState), 6)
				(_nw74).ArraySet1((_this).F25(), 7)
				(_nw74).ArraySet1((_this).F25(), 8)
				(_nw74).ArraySet1((_this).F25(), 9)
				(_nw74).ArraySet1((_this).F25(), 10)
				(_nw74).ArraySet1((_this).F25(), 11)
				(_nw74).ArraySet1(!((_this).F25()), 12)
				(_nw74).ArraySet1(false, 13)
				(_nw74).ArraySet1((_this).F25(), 14)
				(_nw74).ArraySet1((_this).F25(), 15)
				_465_v7 = _nw74
				var _466_v8 _dafny.MultiSet
				_ = _466_v8
				_466_v8 = _dafny.MultiSetOf(_465_v7, _465_v7, _465_v7, _465_v7, _465_v7)
				var _467_v9 _dafny.Map
				_ = _467_v9
				_467_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _466_v8)
				r0 = !((func() _dafny.MultiSet {
					if (_467_v9).Contains(!((_this).F25())) {
						return (_467_v9).Get(!((_this).F25())).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf(_465_v7, _465_v7, _465_v7, _465_v7, _465_v7)
				})()).Contains(_465_v7)
				(globalState).F22 = Companion_Default___.Fm2((_this).F25(), p1, (_this).F25(), (_this).F26(), globalState)
				var _468_v10 _dafny.Map
				_ = _468_v10
				_468_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p1)
				var _469_v11 _dafny.Set
				_ = _469_v11
				_469_v11 = _dafny.SetOf(Companion_Default___.Fm0(p1, p1, (_this).F25(), globalState))
				var _470_v12 _dafny.Sequence
				_ = _470_v12
				_470_v12 = _dafny.SeqOf(_459_v1)
				var _471_v13 _dafny.Int
				_ = _471_v13
				_471_v13 = p1
				var _472_v14 _dafny.Array
				_ = _472_v14
				var _nwElement0_17 _dafny.Int = _459_v1
				_ = _nwElement0_17
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(26))
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_17, 0)
				(_nw75).ArraySet1(((func() _dafny.Int {
					if (_468_v10).Contains((_this).F25()) {
						return (_468_v10).Get((_this).F25()).(_dafny.Int)
					}
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality()))
				})()).Plus(_459_v1), 1)
				(_nw75).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 2)
				(_nw75).ArraySet1((_458_v0).Fm5(globalState), 3)
				(_nw75).ArraySet1(_dafny.IntOfInt64(187), 4)
				(_nw75).ArraySet1(_459_v1, 5)
				(_nw75).ArraySet1(_dafny.IntOfInt64(184), 6)
				(_nw75).ArraySet1(_459_v1, 7)
				(_nw75).ArraySet1(_459_v1, 8)
				(_nw75).ArraySet1(Companion_Default___.SafeDivisionInt(_459_v1, p1), 9)
				(_nw75).ArraySet1((func() _dafny.Int {
					if (_this).F25() {
						return _459_v1
					}
					return _459_v1
				})(), 10)
				(_nw75).ArraySet1((p1).Plus(_dafny.IntOfUint32((p2).Cardinality())), 11)
				(_nw75).ArraySet1((_469_v11).Cardinality(), 12)
				(_nw75).ArraySet1(_dafny.IntOfInt64(940), 13)
				(_nw75).ArraySet1(p1, 14)
				(_nw75).ArraySet1(p1, 15)
				(_nw75).ArraySet1(_459_v1, 16)
				(_nw75).ArraySet1((_458_v0).Fm5(globalState), 17)
				(_nw75).ArraySet1(_459_v1, 18)
				(_nw75).ArraySet1(_459_v1, 19)
				(_nw75).ArraySet1(_459_v1, 20)
				(_nw75).ArraySet1(_dafny.IntOfInt64(676), 21)
				(_nw75).ArraySet1((_470_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-626), _dafny.IntOfUint32((_470_v12).Cardinality()))).Uint32()).(_dafny.Int), 22)
				(_nw75).ArraySet1(_dafny.IntOfInt64(419), 23)
				(_nw75).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm12(_459_v1, (_this).F25(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_459_v1, (_this).F25()), p1, globalState)).Cardinality()), 24)
				(_nw75).ArraySet1((_471_v13), 25)
				_472_v14 = _nw75
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_472_v14), 0))
				_ = _index79
				(_472_v14).ArraySet1(_459_v1, (_index79).Int())
				var _473_v15 _dafny.MultiSet
				_ = _473_v15
				_473_v15 = _dafny.MultiSetOf((_this).Fm3((_this).F25(), p1, _459_v1, _464_v6, globalState))
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_472_v14), 0))
				_ = _index80
				(_472_v14).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_473_v15).Contains((_this).F25()) {
						return (_473_v15).Multiplicity((_this).F25())
					}
					return p1
				})(), (_459_v1).Minus((_468_v10).Cardinality())), (_index80).Int())
				var _474_v16 *C0
				_ = _474_v16
				var _nw76 *C0 = New_C0_()
				_ = _nw76
				_nw76.Ctor__()
				_474_v16 = _nw76
			} else {
				(globalState).F2 = (_459_v1).Cmp(p1) < 0
				var _475_v17 _dafny.Array
				_ = _475_v17
				var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw77
				_475_v17 = _nw77
				_475_v17 = _475_v17
				var _476_v18 _dafny.Set
				_ = _476_v18
				_476_v18 = _dafny.SetOf(p1, p1)
				var _477_v19 _dafny.MultiSet
				_ = _477_v19
				_477_v19 = _dafny.MultiSetOf((_476_v18).Cardinality())
				var _478_v20 _dafny.Sequence
				_ = _478_v20
				_478_v20 = _dafny.SeqOf((_this).F25(), (_this).F25(), (p1).Cmp((_477_v19).Cardinality()) <= 0)
				_478_v20 = _478_v20
				var _479_v21 _dafny.Array
				_ = _479_v21
				var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw78
				_479_v21 = _nw78
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_479_v21), 0))
				_ = _index81
				(_479_v21).ArraySet1(_475_v17, (_index81).Int())
				var _480_v22 D7
				_ = _480_v22
				_480_v22 = Companion_D7_.Create_DC20_(_dafny.SeqOf((_this).F25(), (_this).F25(), (_this).F25(), !((_this).F25())))
				var _481_v23 T0
				_ = _481_v23
				var _nw79 *C1 = New_C1_()
				_ = _nw79
				_nw79.Ctor__(_dafny.Companion_Sequence_.Concatenate(p0, p0), _dafny.IntOfUint32(((_480_v22).Dtor_cf37()).Cardinality()), false)
				_481_v23 = _nw79
				var _482_v24 _dafny.Sequence
				_ = _482_v24
				_482_v24 = _dafny.SeqOf(_458_v0, _458_v0)
				var _483_v25 _dafny.Array
				_ = _483_v25
				var _nwElement0_18 _dafny.Sequence = _482_v24
				_ = _nwElement0_18
				var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(20))
				_ = _nw80
				(_nw80).ArraySet1(_nwElement0_18, 0)
				(_nw80).ArraySet1(_482_v24, 1)
				(_nw80).ArraySet1(_dafny.Companion_Sequence_.Update(_482_v24, (Companion_Default___.SafeIndex(_459_v1, _dafny.IntOfUint32((_482_v24).Cardinality()))).Uint32(), _458_v0), 2)
				(_nw80).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_458_v0, (_482_v24).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_482_v24).Cardinality()))).Uint32()).(*C0)), _482_v24), 3)
				(_nw80).ArraySet1(_482_v24, 4)
				(_nw80).ArraySet1(_482_v24, 5)
				(_nw80).ArraySet1(_482_v24, 6)
				(_nw80).ArraySet1(_482_v24, 7)
				(_nw80).ArraySet1(_482_v24, 8)
				(_nw80).ArraySet1(_dafny.SeqOf(_458_v0), 9)
				(_nw80).ArraySet1(_dafny.SeqOf(_458_v0, _458_v0), 10)
				(_nw80).ArraySet1(_482_v24, 11)
				(_nw80).ArraySet1(_482_v24, 12)
				(_nw80).ArraySet1(_482_v24, 13)
				(_nw80).ArraySet1(_482_v24, 14)
				(_nw80).ArraySet1(_dafny.SeqOf(_458_v0, _458_v0), 15)
				(_nw80).ArraySet1(_482_v24, 16)
				(_nw80).ArraySet1(_dafny.SeqOf(_458_v0), 17)
				(_nw80).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_482_v24, _dafny.SeqOf(_458_v0)), 18)
				(_nw80).ArraySet1(_482_v24, 19)
				_483_v25 = _nw80
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_483_v25), 0))
				_ = _index82
				(_483_v25).ArraySet1(_dafny.SeqOf(_458_v0, _458_v0), (_index82).Int())
				var _484_v26 _dafny.Set
				_ = _484_v26
				_484_v26 = _dafny.SetOf(_dafny.CodePoint('d'))
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_479_v21), 0))
				_ = _index83
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_483_v25), 0))
				_ = _index84
				var _rhs73 _dafny.Array = _475_v17
				_ = _rhs73
				var _rhs74 T0 = _481_v23
				_ = _rhs74
				var _rhs75 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_482_v24, (Companion_Default___.SafeIndex((_484_v26).Cardinality(), _dafny.IntOfUint32((_482_v24).Cardinality()))).Uint32(), _458_v0)
				_ = _rhs75
				var _rhs76 bool = (_481_v23).F25()
				_ = _rhs76
				var _lhs64 _dafny.Array = _479_v21
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_479_v21), 0))
				_ = _lhs65
				var _lhs66 _dafny.Array = _483_v25
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_483_v25), 0))
				_ = _lhs67
				var _lhs68 *GlobalState = globalState
				_ = _lhs68
				(_lhs64).ArraySet1(_rhs73, (_lhs65).Int())
				_481_v23 = _rhs74
				(_lhs66).ArraySet1(_rhs75, (_lhs67).Int())
				_lhs68.F2 = _rhs76
				var _485_v27 _dafny.Array
				_ = _485_v27
				var _nwElement0_19 bool = (p1).Cmp(_459_v1) >= 0
				_ = _nwElement0_19
				var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(11))
				_ = _nw81
				(_nw81).ArraySet1(_nwElement0_19, 0)
				(_nw81).ArraySet1((_481_v23).F25(), 1)
				(_nw81).ArraySet1((Companion_Default___.Fm22(_459_v1, _459_v1, p1, (_this).F25(), globalState)).IsSubsetOf(_484_v26), 2)
				(_nw81).ArraySet1(!((_481_v23).F25()), 3)
				(_nw81).ArraySet1((_this).F25(), 4)
				(_nw81).ArraySet1(!((_this).F25()), 5)
				(_nw81).ArraySet1((_481_v23).F25(), 6)
				(_nw81).ArraySet1((p2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_459_v1), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), 7)
				(_nw81).ArraySet1((_481_v23).F25(), 8)
				(_nw81).ArraySet1(((_this).F25()) == ((_this).F25()), 9)
				(_nw81).ArraySet1((_this).F25(), 10)
				_485_v27 = _nw81
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_485_v27), 0))
				_ = _index85
				(_485_v27).ArraySet1((_481_v23).F25(), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_485_v27), 0))
				_ = _index86
				(_485_v27).ArraySet1((_this).F25(), (_index86).Int())
			}
		} else {
			(globalState).F5 = (func() _dafny.Int {
				if (_this).F25() {
					return p1
				}
				return (_dafny.MultiSetOf(!((_this).F25()))).Cardinality()
			})()
			var _486_v28 *C1
			_ = _486_v28
			var _nw82 *C1 = New_C1_()
			_ = _nw82
			_nw82.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("g"), p1, (_this).F25())
			_486_v28 = _nw82
			var _487_v29 _dafny.Array
			_ = _487_v29
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_14
			var _nw83 _dafny.Array
			_ = _nw83
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw83 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = func(_488_i0 _dafny.Int) _dafny.Int {
					return (_488_i0).Minus(_dafny.IntOfInt64(586))
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw83 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw83).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw83).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_487_v29 = _nw83
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_487_v29), 0))
			_ = _index87
			(_487_v29).ArraySet1(_486_v28.F29, (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_487_v29), 0))
			_ = _index88
			(_487_v29).ArraySet1(p1, (_index88).Int())
			var _489_v30 _dafny.Map
			_ = _489_v30
			_489_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F25()), (_this).F25())
			var _490_v31 D7
			_ = _490_v31
			_490_v31 = Companion_D7_.Create_DC21_((_487_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_487_v29), 0))).Int()).(_dafny.Int), (_489_v30).Merge((_489_v30).Update(!(true), (_this).F25())))
			_490_v31 = Companion_Default___.Fm23(!(!((_this).F25())), globalState)
			var _491_v32 *C0
			_ = _491_v32
			var _nw84 *C0 = New_C0_()
			_ = _nw84
			_nw84.Ctor__()
			_491_v32 = _nw84
		}
		var _492_v33 _dafny.Int
		_ = _492_v33
		_492_v33 = p1
		var _source6 D7 = Companion_D7_.Create_DC22_(p1, _492_v33, (_this).F25())
		_ = _source6
		if _source6.Is_DC21() {
			var _493___mcc_h0 _dafny.Int = _source6.Get_().(D7_DC21).Cf38
			_ = _493___mcc_h0
			var _494___mcc_h1 _dafny.Map = _source6.Get_().(D7_DC21).Cf39
			_ = _494___mcc_h1
			var _495_cf39 _dafny.Map = _494___mcc_h1
			_ = _495_cf39
			var _496_cf38 _dafny.Int = _493___mcc_h0
			_ = _496_cf38
			var _497_v34 _dafny.Map
			_ = _497_v34
			_497_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), (_this).F26())
			var _498_v35 D5
			_ = _498_v35
			_498_v35 = Companion_D5_.Create_DC14_(_497_v34)
			var _499_v36 D5
			_ = _499_v36
			_499_v36 = Companion_D5_.Create_DC16_(_498_v35)
			var _500_v37 _dafny.Set
			_ = _500_v37
			_500_v37 = _dafny.SetOf((_this).F25(), (_this).F25())
			_499_v36 = (func() D5 {
				if !((_500_v37).IsProperSubsetOf(_500_v37)) {
					return _499_v36
				}
				return _499_v36
			})()
			(globalState).F21 = (p1).Plus(p1)
			var _501_v38 _dafny.Array
			_ = _501_v38
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_15
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) bool = func(_502_i1 _dafny.Int) bool {
					return (_this).F25()
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw85 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw85).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw85).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_501_v38 = _nw85
			var _503_v39 _dafny.Set
			_ = _503_v39
			_503_v39 = _dafny.SetOf(p1)
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_501_v38), 0))
			_ = _index89
			(_501_v38).ArraySet1((_503_v39).IsProperSubsetOf(_503_v39), (_index89).Int())
			var _504_v40 _dafny.Map
			_ = _504_v40
			_504_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_cf38, (_this).F25())
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_501_v38), 0))
			_ = _index90
			(_501_v38).ArraySet1(!(((func() bool {
				if (_504_v40).Contains(_496_cf38) {
					return (_504_v40).Get(_496_cf38).(bool)
				}
				return (_this).F25()
			})()) || (false)) || ((_this).F25()), (_index90).Int())
			var _505_v41 _dafny.Sequence
			_ = _505_v41
			_505_v41 = _dafny.SeqOf(_dafny.IntOfUint32((p2).Cardinality()), p1)
			var _506_v42 _dafny.Map
			_ = _506_v42
			_506_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), !((_501_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_501_v38), 0))).Int()).(bool)))
			var _507_v43 _dafny.Array
			_ = _507_v43
			var _nwElement0_20 _dafny.Int = (func() _dafny.Int {
				if (_this).F25() {
					return p1
				}
				return p1
			})()
			_ = _nwElement0_20
			var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(12))
			_ = _nw86
			(_nw86).ArraySet1(_nwElement0_20, 0)
			(_nw86).ArraySet1(_496_cf38, 1)
			(_nw86).ArraySet1(_496_cf38, 2)
			(_nw86).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.SetOf((_501_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_501_v38), 0))).Int()).(bool), (p2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_496_cf38), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool))).Cardinality(), p1), 3)
			(_nw86).ArraySet1(_496_cf38, 4)
			(_nw86).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(97), _496_cf38), 5)
			(_nw86).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F25() {
					return (_505_v41).Select((Companion_Default___.SafeIndex(_496_cf38, _dafny.IntOfUint32((_505_v41).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return p1
			})()), 6)
			(_nw86).ArraySet1(_496_cf38, 7)
			(_nw86).ArraySet1(p1, 8)
			(_nw86).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(((_506_v42).Merge(_506_v42)).Cardinality())), 9)
			(_nw86).ArraySet1((Companion_D3_.Create_DC9_(p1, (_this).F26(), _dafny.IntOfUint32((p0).Cardinality()))).Dtor_cf15(), 10)
			(_nw86).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), 11)
			_507_v43 = _nw86
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_507_v43), 0))
			_ = _index91
			(_507_v43).ArraySet1((_505_v41).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_505_v41).Cardinality()))).Uint32()).(_dafny.Int), (_index91).Int())
			var _508_v44 _dafny.MultiSet
			_ = _508_v44
			_508_v44 = _dafny.MultiSetOf((_501_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_501_v38), 0))).Int()).(bool), (_this).F25())
			var _509_v45 _dafny.Map
			_ = _509_v45
			_509_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_508_v44).Cardinality())
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_507_v43), 0))
			_ = _index92
			var _rhs77 _dafny.Int = Companion_Default___.SafeModuloInt((_509_v45).Cardinality(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-466), _496_cf38))
			_ = _rhs77
			var _rhs78 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p1, _496_cf38), (_this).F25())
			_ = _rhs78
			var _lhs69 _dafny.Array = _507_v43
			_ = _lhs69
			var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_507_v43), 0))
			_ = _lhs70
			(_lhs69).ArraySet1(_rhs77, (_lhs70).Int())
			_504_v40 = _rhs78
		} else if _source6.Is_DC22() {
			var _510___mcc_h2 _dafny.Int = _source6.Get_().(D7_DC22).Cf40
			_ = _510___mcc_h2
			var _511___mcc_h3 _dafny.Int = _source6.Get_().(D7_DC22).Cf41
			_ = _511___mcc_h3
			var _512___mcc_h4 bool = _source6.Get_().(D7_DC22).Cf42
			_ = _512___mcc_h4
			var _513_cf42 bool = _512___mcc_h4
			_ = _513_cf42
			var _514_cf41 _dafny.Int = _511___mcc_h3
			_ = _514_cf41
			var _515_cf40 _dafny.Int = _510___mcc_h2
			_ = _515_cf40
			var _516_v46 *C0
			_ = _516_v46
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__()
			_516_v46 = _nw87
			var _517_v47 _dafny.Int
			_ = _517_v47
			var _out8 _dafny.Int
			_ = _out8
			_out8 = Companion_Default___.M0(globalState)
			_517_v47 = _out8
			var _518_v48 _dafny.Array
			_ = _518_v48
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_16
			var _nw88 _dafny.Array
			_ = _nw88
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw88 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Sequence = (func(_519_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_520_i2 _dafny.Int) _dafny.Sequence {
						return _519_p0
					}
				})(p0)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw88 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw88).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw88).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_518_v48 = _nw88
			var _521_v49 _dafny.Map
			_ = _521_v49
			_521_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _518_v48)
			var _522_v50 _dafny.Map
			_ = _522_v50
			_522_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_513_cf42, (_this).F26())
			var _523_v51 _dafny.Sequence
			_ = _523_v51
			_523_v51 = _dafny.SeqOf(_522_v50, _522_v50)
			_521_v49 = (_521_v49).Update((_dafny.IntOfInt64(-937)).Cmp(_dafny.IntOfUint32((_523_v51).Cardinality())) == 0, _518_v48)
			var _524_v52 D3
			_ = _524_v52
			_524_v52 = Companion_D3_.Create_DC8_(!((_this).F25()))
			(globalState).F2 = !((_524_v52).Dtor_cf12()) || ((p1).Cmp(_517_v47) != 0)
		} else if _source6.Is_DC23() {
			var _525___mcc_h5 _dafny.Int = _source6.Get_().(D7_DC23).Cf43
			_ = _525___mcc_h5
			var _526___mcc_h6 _dafny.Int = _source6.Get_().(D7_DC23).Cf44
			_ = _526___mcc_h6
			var _527___mcc_h7 bool = _source6.Get_().(D7_DC23).Cf45
			_ = _527___mcc_h7
			var _528___mcc_h8 _dafny.Int = _source6.Get_().(D7_DC23).Cf46
			_ = _528___mcc_h8
			var _529_cf46 _dafny.Int = _528___mcc_h8
			_ = _529_cf46
			var _530_cf45 bool = _527___mcc_h7
			_ = _530_cf45
			var _531_cf44 _dafny.Int = _526___mcc_h6
			_ = _531_cf44
			var _532_cf43 _dafny.Int = _525___mcc_h5
			_ = _532_cf43
			var _533_v53 _dafny.MultiSet
			_ = _533_v53
			_533_v53 = _dafny.MultiSetOf((_this).F25())
			_530_cf45 = (_533_v53).IsProperSubsetOf(_533_v53)
			if !(_530_cf45) {
				(globalState).F18 = (func() _dafny.Int {
					if _530_cf45 {
						return _531_cf44
					}
					return (_dafny.Zero).Minus(p1)
				})()
				var _534_v54 D2
				_ = _534_v54
				_534_v54 = Companion_D2_.Create_DC5_(_dafny.IntOfInt64(-739), (_this).F25(), _530_cf45, Companion_Default___.Fm0(_531_cf44, _532_cf43, true, globalState))
				var _535_v55 _dafny.Map
				_ = _535_v55
				_535_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_dafny.Zero).Minus(_531_cf44))
				var _536_v56 _dafny.Map
				_ = _536_v56
				_536_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_530_cf45, (_this).F25())
				var _pat_let_tv9 = _530_cf45
				_ = _pat_let_tv9
				var _537_v57 _dafny.Array
				_ = _537_v57
				var _nwElement0_21 D2 = _534_v54
				_ = _nwElement0_21
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(18))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_21, 0)
				(_nw89).ArraySet1(_534_v54, 1)
				(_nw89).ArraySet1(_534_v54, 2)
				(_nw89).ArraySet1(Companion_D2_.Create_DC5_((_535_v55).Cardinality(), (_this).F25(), (func() bool {
					if (_536_v56).Contains((_this).F25()) {
						return (_536_v56).Get((_this).F25()).(bool)
					}
					return (_this).F25()
				})(), _530_cf45), 3)
				(_nw89).ArraySet1(_534_v54, 4)
				(_nw89).ArraySet1(func(_pat_let4_0 D2) D2 {
					return func(_538_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let5_0 bool) D2 {
							return func(_539_dt__update_hcf7_h0 bool) D2 {
								return func(_pat_let6_0 bool) D2 {
									return func(_540_dt__update_hcf8_h0 bool) D2 {
										return Companion_D2_.Create_DC5_((_538_dt__update__tmp_h0).Dtor_cf6(), _539_dt__update_hcf7_h0, _540_dt__update_hcf8_h0, (_538_dt__update__tmp_h0).Dtor_cf9())
									}(_pat_let6_0)
								}((_this).F25())
							}(_pat_let5_0)
						}(_pat_let_tv9)
					}(_pat_let4_0)
				}(_534_v54), 5)
				(_nw89).ArraySet1(_534_v54, 6)
				(_nw89).ArraySet1(_534_v54, 7)
				(_nw89).ArraySet1(_534_v54, 8)
				(_nw89).ArraySet1(_534_v54, 9)
				(_nw89).ArraySet1(_534_v54, 10)
				(_nw89).ArraySet1(_534_v54, 11)
				(_nw89).ArraySet1(_534_v54, 12)
				(_nw89).ArraySet1(_534_v54, 13)
				(_nw89).ArraySet1(_534_v54, 14)
				(_nw89).ArraySet1(_534_v54, 15)
				(_nw89).ArraySet1(_534_v54, 16)
				(_nw89).ArraySet1(_534_v54, 17)
				_537_v57 = _nw89
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_537_v57), 0))
				_ = _index93
				(_537_v57).ArraySet1(_534_v54, (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_537_v57), 0))
				_ = _index94
				(_537_v57).ArraySet1(_534_v54, (_index94).Int())
				r0 = (_this).F25()
				var _541_v58 _dafny.Array
				_ = _541_v58
				var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(7))
				_ = _nw90
				_541_v58 = _nw90
				var _542_v59 *C0
				_ = _542_v59
				var _nw91 *C0 = New_C0_()
				_ = _nw91
				_nw91.Ctor__()
				_542_v59 = _nw91
				var _543_v60 _dafny.Set
				_ = _543_v60
				_543_v60 = _dafny.SetOf(_542_v59, _542_v59)
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_541_v58), 0))
				_ = _index95
				(_541_v58).ArraySet1(_543_v60, (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(351), _dafny.ArrayLen((_541_v58), 0))
				_ = _index96
				(_541_v58).ArraySet1(_dafny.SetOf(_542_v59), (_index96).Int())
				var _544_v61 *C0
				_ = _544_v61
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__()
				_544_v61 = _nw92
			} else {
				var _545_v62 _dafny.Sequence
				_ = _545_v62
				_545_v62 = _dafny.SeqOf(_529_cf46, _529_cf46)
				var _546_v63 _dafny.Map
				_ = _546_v63
				_546_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_529_cf46, Companion_Default___.SafeModuloInt((_545_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_545_v62).Cardinality()))).Uint32()).(_dafny.Int), p1))
				_546_v63 = (_546_v63).Update((_529_cf46).Minus(Companion_Default___.Fm2(false, _531_cf44, _530_cf45, (_this).F26(), globalState)), _529_cf46)
				var _547_v64 *C1
				_ = _547_v64
				var _nw93 *C1 = New_C1_()
				_ = _nw93
				_nw93.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("awhlibqi"), (_529_cf46).Minus(p1), _530_cf45)
				_547_v64 = _nw93
				_530_cf45 = _530_cf45
				(globalState).F22 = (Companion_Default___.SafeDivisionInt(_529_cf46, _dafny.IntOfInt64(160))).Times(_547_v64.F29)
				var _548_v66 _dafny.Map
				_ = _548_v66
				_548_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_cf44, (_this).F25())
				var _549_v67 D6
				_ = _549_v67
				_549_v67 = Companion_D6_.Create_DC17_(func() _dafny.Map {
					var _coll21 = _dafny.NewMapBuilder()
					_ = _coll21
					for _iter25 := _dafny.Iterate((_548_v66).Keys().Elements()); ; {
						_compr_21, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _550_v65 _dafny.Int
						_550_v65 = interface{}(_compr_21).(_dafny.Int)
						if (_548_v66).Contains(_550_v65) {
							_coll21.Add((_550_v65).Minus(_531_cf44), _531_cf44)
						}
					}
					return _coll21.ToMap()
				}())
				var _551_v68 D6
				_ = _551_v68
				_551_v68 = Companion_D6_.Create_DC19_(_549_v67)
				_551_v68 = _551_v68
			}
			var _552_v69 _dafny.Array
			_ = _552_v69
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_17
			var _nw94 _dafny.Array
			_ = _nw94
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw94 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.CodePoint = func(_553_i3 _dafny.Int) _dafny.CodePoint {
					return (_this).F26()
				}
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw94 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw94).ArraySet1CodePoint(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw94).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_552_v69 = _nw94
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_552_v69), 0))
			_ = _index97
			(_552_v69).ArraySet1CodePoint((_this).F26(), (_index97).Int())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_552_v69), 0))
			_ = _index98
			(_552_v69).ArraySet1CodePoint((_this).F26(), (_index98).Int())
			var _554_v70 _dafny.Set
			_ = _554_v70
			_554_v70 = _dafny.SetOf(!(_530_cf45))
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_552_v69), 0))
			_ = _index99
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_552_v69), 0))
			_ = _index100
			var _rhs79 _dafny.Int = (_dafny.IntOfInt64(169)).Plus(_532_cf43)
			_ = _rhs79
			var _rhs80 _dafny.CodePoint = (_this).F26()
			_ = _rhs80
			var _rhs81 _dafny.CodePoint = _dafny.CodePoint('n')
			_ = _rhs81
			var _rhs82 _dafny.Int = ((_554_v70).Intersection(_554_v70)).Cardinality()
			_ = _rhs82
			var _rhs83 _dafny.Int = p1
			_ = _rhs83
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			var _lhs72 _dafny.Array = _552_v69
			_ = _lhs72
			var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_552_v69), 0))
			_ = _lhs73
			var _lhs74 _dafny.Array = _552_v69
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_552_v69), 0))
			_ = _lhs75
			_lhs71.F22 = _rhs79
			(_lhs72).ArraySet1CodePoint(_rhs80, (_lhs73).Int())
			(_lhs74).ArraySet1CodePoint(_rhs81, (_lhs75).Int())
			r1 = _rhs82
			_531_cf44 = _rhs83
			(globalState).F18 = (p1).Plus(_529_cf46)
		} else if _source6.Is_DC20() {
			var _555___mcc_h9 _dafny.Sequence = _source6.Get_().(D7_DC20).Cf37
			_ = _555___mcc_h9
			var _556_cf37 _dafny.Sequence = _555___mcc_h9
			_ = _556_cf37
			(globalState).F2 = (_this).F25()
			if !(true) {
				var _557_v71 _dafny.MultiSet
				_ = _557_v71
				_557_v71 = _dafny.MultiSetOf(p1, p1, _dafny.IntOfInt64(988), p1, p1)
				(globalState).F2 = Companion_Default___.Fm0((_dafny.IntOfInt64(610)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(763))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_558_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_559_i4 _dafny.Int) _dafny.Int {
						return _558_p1
					}
				})(p1)))).Cardinality())), Companion_Default___.Fm2((_this).F25(), (_dafny.Zero).Minus(((_557_v71).Update(_dafny.IntOfInt64(909), Companion_Default___.Abs(p1))).Cardinality()), !((_this).F25()), (_this).F26(), globalState), !(false), globalState)
				(globalState).F21 = p1
				(globalState).F2 = (_this).F25()
				var _560_v72 _dafny.Array
				_ = _560_v72
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_18
				var _nw95 _dafny.Array
				_ = _nw95
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw95 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Map = func(_561_i5 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25()))
					}
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw95 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw95).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw95).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_560_v72 = _nw95
				_560_v72 = _560_v72
				var _562_v73 _dafny.Array
				_ = _562_v73
				var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw96
				_562_v73 = _nw96
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_562_v73), 0))
				_ = _index101
				(_562_v73).ArraySet1(p1, (_index101).Int())
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_562_v73), 0))
				_ = _index102
				(_562_v73).ArraySet1(p1, (_index102).Int())
			} else {
				var _563_v74 _dafny.Array
				_ = _563_v74
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_19
				var _nw97 _dafny.Array
				_ = _nw97
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw97 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = func(_564_i6 _dafny.Int) bool {
						return !((_this).F25()) || ((_this).F25())
					}
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw97 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw97).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw97).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_563_v74 = _nw97
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))
				_ = _index103
				(_563_v74).ArraySet1((_this).F25(), (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))
				_ = _index104
				(_563_v74).ArraySet1((_this).F25(), (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))
				_ = _index105
				(_563_v74).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("hqccefcs"), Companion_Default___.Fm16(p1, globalState)), (_index105).Int())
				var _565_v75 _dafny.Sequence
				_ = _565_v75
				_565_v75 = _dafny.SeqOf(p2, p2, _556_cf37, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F25()), p2), p2)
				var _566_v76 D7
				_ = _566_v76
				_566_v76 = Companion_D7_.Create_DC20_(_556_cf37)
				var _567_v77 _dafny.Map
				_ = _567_v77
				_567_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_563_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))).Int()).(bool))
				var _rhs84 bool = (_563_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))).Int()).(bool)
				_ = _rhs84
				var _rhs85 _dafny.Int = (_dafny.IntOfUint32((p0).Cardinality())).Minus((_dafny.Zero).Minus((p1).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _566_v76)).Cardinality())).Cardinality()))))
				_ = _rhs85
				var _rhs86 _dafny.Int = (func() _dafny.Int {
					if (_this).F25() {
						return _dafny.IntOfUint32((p0).Cardinality())
					}
					return p1
				})()
				_ = _rhs86
				var _rhs87 _dafny.Sequence = _565_v75
				_ = _rhs87
				var _rhs88 _dafny.Int = (_567_v77).Cardinality()
				_ = _rhs88
				var _lhs76 *GlobalState = globalState
				_ = _lhs76
				var _lhs77 *GlobalState = globalState
				_ = _lhs77
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				_lhs76.F2 = _rhs84
				r1 = _rhs85
				_lhs77.F22 = _rhs86
				_565_v75 = _rhs87
				_lhs78.F22 = _rhs88
				_556_cf37 = p2
				var _568_v78 _dafny.Map
				_ = _568_v78
				_568_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), p1)
				var _569_v79 _dafny.MultiSet
				_ = _569_v79
				_569_v79 = _dafny.MultiSetOf((_563_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))).Int()).(bool), (_563_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))).Int()).(bool), (_563_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))).Int()).(bool))
				var _570_v80 _dafny.MultiSet
				_ = _570_v80
				_570_v80 = _dafny.MultiSetOf((_568_v78).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_563_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_563_v74), 0))).Int()).(bool), (_569_v79).Cardinality())))
				_570_v80 = (_570_v80).Union(_570_v80)
			}
			var _571_v81 _dafny.Sequence
			_ = _571_v81
			_571_v81 = _dafny.SeqOf(p1)
			var _572_v82 _dafny.MultiSet
			_ = _572_v82
			_572_v82 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())))
			(globalState).F2 = (_572_v82).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_571_v81, _571_v81)))
			var _573_v83 _dafny.Map
			_ = _573_v83
			_573_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(612), (_this).F26())
			var _574_v84 D5
			_ = _574_v84
			_574_v84 = Companion_D5_.Create_DC16_(Companion_D5_.Create_DC14_(_573_v83))
			var _575_v85 _dafny.MultiSet
			_ = _575_v85
			_575_v85 = _dafny.MultiSetOf(p2, _556_cf37, _dafny.SeqOf((_this).F25()))
			var _rhs89 D5 = (func() D5 {
				if (_this).F25() {
					return _574_v84
				}
				return _574_v84
			})()
			_ = _rhs89
			var _rhs90 bool = (p1).Cmp(((_575_v85).Update(p2, Companion_Default___.Abs(p1))).Cardinality()) <= 0
			_ = _rhs90
			var _rhs91 _dafny.Int = _492_v33
			_ = _rhs91
			_574_v84 = _rhs89
			r0 = _rhs90
			_492_v33 = _rhs91
		} else {
			var _576___mcc_h10 D7 = _source6.Get_().(D7_DC24).Cf47
			_ = _576___mcc_h10
			var _577_cf47 D7 = _576___mcc_h10
			_ = _577_cf47
			var _578_v86 D2
			_ = _578_v86
			_578_v86 = Companion_D2_.Create_DC5_(p1, (_this).F25(), !((_this).F25()), (_this).F25())
			var _579_v87 D2
			_ = _579_v87
			_579_v87 = Companion_D2_.Create_DC6_(_578_v86)
			var _580_v88 D2
			_ = _580_v88
			_580_v88 = Companion_D2_.Create_DC6_(_579_v87)
			var _pat_let_tv10 = _579_v87
			_ = _pat_let_tv10
			_580_v88 = func(_pat_let7_0 D2) D2 {
				return func(_581_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let8_0 D2) D2 {
						return func(_582_dt__update_hcf10_h0 D2) D2 {
							return Companion_D2_.Create_DC6_(_582_dt__update_hcf10_h0)
						}(_pat_let8_0)
					}(_pat_let_tv10)
				}(_pat_let7_0)
			}(_580_v88)
			var _583_v89 _dafny.Array
			_ = _583_v89
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_20
			var _nw98 _dafny.Array
			_ = _nw98
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw98 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Int = func(_584_i7 _dafny.Int) _dafny.Int {
					return (_584_i7).Times(_dafny.IntOfInt64(741))
				}
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw98 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw98).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw98).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_583_v89 = _nw98
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_583_v89), 0))
			_ = _index106
			(_583_v89).ArraySet1(p1, (_index106).Int())
			var _585_v90 D7
			_ = _585_v90
			_585_v90 = Companion_D7_.Create_DC22_(_dafny.IntOfUint32((p0).Cardinality()), _492_v33, (_this).F25())
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_583_v89), 0))
			_ = _index107
			(_583_v89).ArraySet1(Companion_Default___.SafeDivisionInt((_585_v90).Dtor_cf40(), p1), (_index107).Int())
			var _586_v91 _dafny.MultiSet
			_ = _586_v91
			_586_v91 = _dafny.MultiSetOf(p1, p1, p1, p1)
			var _587_v92 _dafny.Sequence
			_ = _587_v92
			_587_v92 = _dafny.SeqOf(p1, (Companion_Default___.Fm24((_this).F25(), p1, p1, globalState)).Cardinality())
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_583_v89), 0))
			_ = _index108
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_583_v89), 0))
			_ = _index109
			var _rhs92 _dafny.Int = p1
			_ = _rhs92
			var _rhs93 _dafny.Int = p1
			_ = _rhs93
			var _rhs94 bool = (_dafny.MultiSetOf(_dafny.IntOfUint32((_587_v92).Cardinality()), p1, p1)).IsProperSubsetOf(_586_v91)
			_ = _rhs94
			var _lhs79 _dafny.Array = _583_v89
			_ = _lhs79
			var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_583_v89), 0))
			_ = _lhs80
			var _lhs81 _dafny.Array = _583_v89
			_ = _lhs81
			var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_583_v89), 0))
			_ = _lhs82
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			(_lhs79).ArraySet1(_rhs92, (_lhs80).Int())
			(_lhs81).ArraySet1(_rhs93, (_lhs82).Int())
			_lhs83.F2 = _rhs94
			(globalState).F2 = (_this).F25()
			var _588_v93 _dafny.Set
			_ = _588_v93
			_588_v93 = _dafny.SetOf(!((_this).F25()))
			var _589_v94 _dafny.Map
			_ = _589_v94
			_589_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_588_v93, (_this).F25())
			_589_v94 = (_589_v94).Update(_588_v93, (p1).Cmp(p1) != 0)
		}
		var _590_v95 _dafny.Sequence
		_ = _590_v95
		_590_v95 = _dafny.SeqOf(p1, p1, p1)
		var _hi5 _dafny.Int = _dafny.IntOfUint32((_590_v95).Cardinality())
		_ = _hi5
		for _591_i8 := p1; _591_i8.Cmp(_hi5) < 0; _591_i8 = _591_i8.Plus(_dafny.One) {
			var _592_v96 *C0
			_ = _592_v96
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__()
			_592_v96 = _nw99
			var _593_v97 _dafny.Map
			_ = _593_v97
			_593_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), true)
			var _594_v98 _dafny.Map
			_ = _594_v98
			_594_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if (_593_v97).Contains(false) {
					return (_593_v97).Get(false).(bool)
				}
				return (_this).F25()
			})()), (_this).F25())
			_594_v98 = (_594_v98).Update(true, (_this).F25())
			var _595_v99 _dafny.Sequence
			_ = _595_v99
			_595_v99 = _dafny.SeqOf(_590_v95)
			(globalState).F19 = _dafny.Companion_Sequence_.Concatenate(_590_v95, (_595_v99).Select((Companion_Default___.SafeIndex(_591_i8, _dafny.IntOfUint32((_595_v99).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _596_v100 _dafny.Array
			_ = _596_v100
			var _nw100 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw100
			_596_v100 = _nw100
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_596_v100), 0))
			_ = _index110
			(_596_v100).ArraySet1((p1).Cmp(p1) == 0, (_index110).Int())
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_596_v100), 0))
			_ = _index111
			(_596_v100).ArraySet1(true, (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_596_v100), 0))
			_ = _index112
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_596_v100), 0))
			_ = _index113
			var _rhs95 bool = (_this).F25()
			_ = _rhs95
			var _rhs96 bool = (_this).F25()
			_ = _rhs96
			var _rhs97 bool = (p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool)
			_ = _rhs97
			var _rhs98 _dafny.Int = p1
			_ = _rhs98
			var _rhs99 _dafny.Int = _591_i8
			_ = _rhs99
			var _lhs84 _dafny.Array = _596_v100
			_ = _lhs84
			var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(269), _dafny.ArrayLen((_596_v100), 0))
			_ = _lhs85
			var _lhs86 _dafny.Array = _596_v100
			_ = _lhs86
			var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_596_v100), 0))
			_ = _lhs87
			var _lhs88 *GlobalState = globalState
			_ = _lhs88
			var _lhs89 *GlobalState = globalState
			_ = _lhs89
			r0 = _rhs95
			(_lhs84).ArraySet1(_rhs96, (_lhs85).Int())
			(_lhs86).ArraySet1(_rhs97, (_lhs87).Int())
			_lhs88.F22 = _rhs98
			_lhs89.F22 = _rhs99
		}
		var _597_v101 _dafny.MultiSet
		_ = _597_v101
		_597_v101 = _dafny.MultiSetOf(p1)
		var _598_v102 _dafny.Map
		_ = _598_v102
		_598_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
		var _599_v103 D7
		_ = _599_v103
		_599_v103 = Companion_D7_.Create_DC21_((_dafny.Zero).Minus(p1), _598_v102)
		var _600_v104 _dafny.Array
		_ = _600_v104
		var _nwElement0_22 bool = (_this).F25()
		_ = _nwElement0_22
		var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(24))
		_ = _nw101
		(_nw101).ArraySet1(_nwElement0_22, 0)
		(_nw101).ArraySet1((p1).Cmp(_dafny.IntOfInt64(-887)) != 0, 1)
		(_nw101).ArraySet1((_597_v101).IsSubsetOf((_597_v101).Update(p1, Companion_Default___.Abs((_599_v103).Dtor_cf38()))), 2)
		(_nw101).ArraySet1(!((_this).F25()), 3)
		(_nw101).ArraySet1(((_dafny.Zero).Minus(p1)).Cmp(p1) < 0, 4)
		(_nw101).ArraySet1((_this).F25(), 5)
		(_nw101).ArraySet1((!((p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool))) && ((_this).F25()), 6)
		(_nw101).ArraySet1((_this).F25(), 7)
		(_nw101).ArraySet1((_this).F25(), 8)
		(_nw101).ArraySet1((_this).F25(), 9)
		(_nw101).ArraySet1((_this).F25(), 10)
		(_nw101).ArraySet1((_this).F25(), 11)
		(_nw101).ArraySet1((_this).F25(), 12)
		(_nw101).ArraySet1((p1).Cmp(p1) == 0, 13)
		(_nw101).ArraySet1(!_dafny.Companion_Sequence_.Contains(p0, (_this).F26()), 14)
		(_nw101).ArraySet1((_597_v101).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-374))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_601_i9 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(254)
		})))), 15)
		(_nw101).ArraySet1((_this).F25(), 16)
		(_nw101).ArraySet1((p2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), 17)
		(_nw101).ArraySet1((_this).F25(), 18)
		(_nw101).ArraySet1(false, 19)
		(_nw101).ArraySet1((func() bool {
			if (_this).F25() {
				return (_this).F25()
			}
			return (_this).F25()
		})(), 20)
		(_nw101).ArraySet1((_this).F25(), 21)
		(_nw101).ArraySet1((_this).F25(), 22)
		(_nw101).ArraySet1((_this).F25(), 23)
		_600_v104 = _nw101
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))
		_ = _index114
		(_600_v104).ArraySet1((_this).F25(), (_index114).Int())
		var _602_v105 _dafny.Set
		_ = _602_v105
		_602_v105 = _dafny.SetOf((_this).F26(), (_this).F26(), Companion_Default___.Fm16(p1, globalState))
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))
		_ = _index115
		(_600_v104).ArraySet1((Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-284), p1)).Cmp(((_602_v105).Difference(_602_v105)).Cardinality()) != 0, (_index115).Int())
		var _603_i10 _dafny.Int
		_ = _603_i10
		_603_i10 = _dafny.Zero
		{
			for (_this).F25() {
				{
					if (_603_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_603_i10 = (_603_i10).Plus(_dafny.One)
					var _604_v106 _dafny.Array
					_ = _604_v106
					var _len0_21 _dafny.Int = _dafny.IntOfInt64(19)
					_ = _len0_21
					var _nw102 _dafny.Array
					_ = _nw102
					if _len0_21.Cmp(_dafny.Zero) == 0 {
						_nw102 = _dafny.NewArray(_len0_21)
					} else {
						var _init21 func(_dafny.Int) _dafny.Int = (func(_605_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_606_i11 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_606_i11, _605_p1)
							}
						})(p1)
						_ = _init21
						var _element0_21 = _init21(_dafny.Zero)
						_ = _element0_21
						_nw102 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
						(_nw102).ArraySet1(_element0_21, 0)
						var _nativeLen0_21 = (_len0_21).Int()
						_ = _nativeLen0_21
						for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
							(_nw102).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
						}
					}
					_604_v106 = _nw102
					var _607_v107 _dafny.Set
					_ = _607_v107
					_607_v107 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(505))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}((func(_608_p0 _dafny.Sequence, _609_v105 _dafny.Set) func(_dafny.Int) _dafny.Int {
						return func(_610_i12 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_608_p0, (_609_v105).Cardinality())).Cardinality(), (_this).F25())).Cardinality()
						}
					})(p0, _602_v105))))
					var _611_v108 D7
					_ = _611_v108
					_611_v108 = Companion_D7_.Create_DC23_(_dafny.IntOfInt64(934), p1, (_600_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))).Int()).(bool), p1)
					var _612_v109 _dafny.Map
					_ = _612_v109
					_612_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_590_v95, _611_v108)
					var _613_v111 _dafny.Sequence
					_ = _613_v111
					_613_v111 = _dafny.SeqOf(_600_v104)
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))
					_ = _index116
					var _rhs100 bool = (func() _dafny.Set {
						var _coll22 = _dafny.NewBuilder()
						_ = _coll22
						for _iter26 := _dafny.Iterate((_612_v109).Keys().Elements()); ; {
							_compr_22, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _614_v110 _dafny.Sequence
							_614_v110 = interface{}(_compr_22).(_dafny.Sequence)
							if (_612_v109).Contains(_614_v110) {
								_coll22.Add(_614_v110)
							}
						}
						return _coll22.ToSet()
					}()).IsSubsetOf(_607_v107)
					_ = _rhs100
					var _rhs101 bool = (_600_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))).Int()).(bool)
					_ = _rhs101
					var _rhs102 _dafny.Array = (_613_v111).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_613_v111).Cardinality()))).Uint32()).(_dafny.Array)
					_ = _rhs102
					var _rhs103 bool = (_this).F25()
					_ = _rhs103
					var _rhs104 _dafny.Array = _604_v106
					_ = _rhs104
					var _lhs90 _dafny.Array = _600_v104
					_ = _lhs90
					var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))
					_ = _lhs91
					(_lhs90).ArraySet1(_rhs100, (_lhs91).Int())
					r0 = _rhs101
					_600_v104 = _rhs102
					r0 = _rhs103
					_604_v106 = _rhs104
					var _615_v112 _dafny.Map
					_ = _615_v112
					_615_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F25())
					var _616_v113 _dafny.Set
					_ = _616_v113
					_616_v113 = _dafny.SetOf(Companion_Default___.Fm2((_this).F25(), _dafny.IntOfInt64(-283), (_600_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))).Int()).(bool), (_this).F26(), globalState))
					var _617_v114 _dafny.Array
					_ = _617_v114
					var _nwElement0_23 _dafny.Set = _dafny.SetOf(p1, Companion_Default___.Fm2((_600_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))).Int()).(bool), _dafny.IntOfUint32((p2).Cardinality()), (_600_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))).Int()).(bool), _dafny.CodePoint('m'), globalState), p1)
					_ = _nwElement0_23
					var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(6))
					_ = _nw103
					(_nw103).ArraySet1(_nwElement0_23, 0)
					(_nw103).ArraySet1(Companion_Default___.Fm24((_600_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))).Int()).(bool), (_615_v112).Cardinality(), p1, globalState), 1)
					(_nw103).ArraySet1(_616_v113, 2)
					(_nw103).ArraySet1((_616_v113).Union(_dafny.SetOf(p1, (_dafny.Zero).Minus(p1))), 3)
					(_nw103).ArraySet1(_616_v113, 4)
					(_nw103).ArraySet1((_616_v113).Union(_dafny.SetOf(Companion_Default___.Fm2((_this).F25(), p1, (_this).F25(), (_this).F26(), globalState), _dafny.IntOfUint32((p2).Cardinality()))), 5)
					_617_v114 = _nw103
					_617_v114 = _617_v114
					(globalState).F18 = (func() _dafny.Int {
						if true {
							return ((_dafny.Zero).Minus(p1)).Minus(p1)
						}
						return (p1).Minus(_dafny.IntOfInt64(-111))
					})()
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_600_v104), 0))
					_ = _index117
					(_600_v104).ArraySet1(true, (_index117).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _618_v115 D3
		_ = _618_v115
		_618_v115 = Companion_D3_.Create_DC9_(p1, (_this).F26(), p1)
		var _619_v116 _dafny.MultiSet
		_ = _619_v116
		_619_v116 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_590_v95).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_618_v115).Dtor_cf14()), p0)
		r0 = ((_619_v116).Update(_dafny.UnicodeSeqOfUtf8Bytes("v"), Companion_Default___.Abs(_dafny.IntOfInt64(764)))).IsProperSubsetOf(_619_v116)
		r0 = (_this).F25()
		r1 = p1
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f26 _dafny.CodePoint
	_f27 _dafny.Array
	_f25 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f26 = _dafny.CodePoint('D')
	_this._f27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f25 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F26() _dafny.CodePoint {
	return _this._f26
}
func (_this *C3) F27() _dafny.Array {
	return _this._f27
}
func (_this *C3) F25() bool {
	return _this._f25
}
func (_this *C3) Ctor__(f26 _dafny.CodePoint, f27 _dafny.Array, f25 bool) {
	{
		(_this)._f26 = f26
		(_this)._f27 = f27
		(_this)._f25 = f25
	}
}
func (_this *C3) Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return (_this).F25()
	}
}
func (_this *C3) Fm4(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus(_dafny.IntOfInt64(-836))).Times(_dafny.IntOfUint32((_dafny.SeqOf((_this).F25(), (_this).F25())).Cardinality())), _dafny.IntOfInt64(314))
	}
}
func (_this *C3) M2(p0 _dafny.Set, p1 bool, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _620_v0 _dafny.Array
		_ = _620_v0
		var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
		_ = _nw104
		_620_v0 = _nw104
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))
		_ = _index118
		(_620_v0).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), (_index118).Int())
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))
		_ = _index119
		(_620_v0).ArraySet1(Companion_Default___.SafeDivisionInt(p2, Companion_Default___.SafeDivisionInt(Companion_Default___.Fm2((_this).F25(), Companion_Default___.Fm2((_this).F25(), p2, false, _dafny.CodePoint('e'), globalState), p1, (_this).F26(), globalState), (func() _dafny.Map {
			var _coll23 = _dafny.NewMapBuilder()
			_ = _coll23
			for _iter27 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_621_p0 _dafny.Set) func(_dafny.Int) _dafny.Sequence {
				return func(_622_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_621_p0).Cardinality())
				}
			})(p0)))).Elements()); ; {
				_compr_23, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _623_v1 _dafny.Sequence
				_623_v1 = interface{}(_compr_23).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_624_p0 _dafny.Set) func(_dafny.Int) _dafny.Sequence {
					return func(_622_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((_624_p0).Cardinality())
					}
				})(p0))), _623_v1) {
					_coll23.Add(_623_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cela")).Cardinality()))
				}
			}
			return _coll23.ToMap()
		}()).Cardinality())), (_index119).Int())
		var _625_v2 _dafny.Sequence
		_ = _625_v2
		_625_v2 = _dafny.UnicodeSeqOfUtf8Bytes("pgjv")
		_625_v2 = _625_v2
		var _626_v3 *C0
		_ = _626_v3
		var _nw105 *C0 = New_C0_()
		_ = _nw105
		_nw105.Ctor__()
		_626_v3 = _nw105
		var _627_v4 *C0
		_ = _627_v4
		var _nw106 *C0 = New_C0_()
		_ = _nw106
		_nw106.Ctor__()
		_627_v4 = _nw106
		(globalState).F2 = ((_this).F25()) == (p1)
		if p1 {
			(globalState).F18 = Companion_Default___.SafeDivisionInt(p2, p2)
			var _628_v5 _dafny.Set
			_ = _628_v5
			_628_v5 = _dafny.SetOf((_this).F25())
			(globalState).F5 = (_628_v5).Cardinality()
			(globalState).F22 = ((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int)).Plus((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int))
			if p1 {
				var _629_v6 _dafny.Array
				_ = _629_v6
				var _nw107 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
				_ = _nw107
				_629_v6 = _nw107
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_629_v6), 0))
				_ = _index120
				(_629_v6).ArraySet1(_626_v3, (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_629_v6), 0))
				_ = _index121
				(_629_v6).ArraySet1(_626_v3, (_index121).Int())
				var _630_v7 _dafny.Array
				_ = _630_v7
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
				_ = _nw108
				_630_v7 = _nw108
				var _631_v8 _dafny.Sequence
				_ = _631_v8
				_631_v8 = _dafny.SeqOf(_625_v2)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_630_v7), 0))
				_ = _index122
				(_630_v7).ArraySet1(_631_v8, (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_630_v7), 0))
				_ = _index123
				(_630_v7).ArraySet1(_631_v8, (_index123).Int())
				(globalState).F2 = !(p1)
				_625_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_625_v2, _625_v2), _625_v2)
				_625_v2 = _dafny.Companion_Sequence_.Concatenate(_625_v2, _dafny.UnicodeSeqOfUtf8Bytes("rr"))
			} else {
				(globalState).F21 = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int), p2)).Minus(_dafny.IntOfInt64(495)))
				(globalState).F18 = p2
				var _632_v9 _dafny.Map
				_ = _632_v9
				_632_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int), p2)
				_632_v9 = (_632_v9).Update(p2, (p2).Minus(_dafny.IntOfInt64(319)))
				var _633_v10 _dafny.Array
				_ = _633_v10
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw109
				_633_v10 = _nw109
				var _634_v11 _dafny.Array
				_ = _634_v11
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_22
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_635_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_636_i1 _dafny.Int) _dafny.Sequence {
							return _635_v2
						}
					})(_625_v2)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw110 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw110).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw110).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_634_v11 = _nw110
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_633_v10), 0))
				_ = _index124
				(_633_v10).ArraySet1(_634_v11, (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_633_v10), 0))
				_ = _index125
				(_633_v10).ArraySet1(_634_v11, (_index125).Int())
				var _637_v12 _dafny.Array
				_ = _637_v12
				var _nw111 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
				_ = _nw111
				_637_v12 = _nw111
				var _638_v13 _dafny.Sequence
				_ = _638_v13
				_638_v13 = _dafny.SeqOf(_637_v12)
				_637_v12 = (_638_v13).Select((Companion_Default___.SafeIndex((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_638_v13).Cardinality()))).Uint32()).(_dafny.Array)
			}
			var _639_v14 _dafny.Int
			_ = _639_v14
			_639_v14 = p2
			var _640_v15 _dafny.Sequence
			_ = _640_v15
			_640_v15 = _dafny.SeqOf(p2, (_639_v14))
			var _641_v16 _dafny.Map
			_ = _641_v16
			_641_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-107), p1)
			var _642_v17 _dafny.Map
			_ = _642_v17
			_642_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_640_v15).Cardinality())).Cmp((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int)) != 0, _641_v16)
			_642_v17 = (_642_v17).Update(true, _641_v16)
		} else {
			var _643_v18 _dafny.Map
			_ = _643_v18
			_643_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
			(globalState).F2 = Companion_Default___.Fm0((_dafny.Zero).Minus((_643_v18).Cardinality()), (_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int), (_this).F25(), globalState)
			(globalState).F18 = (_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int)
			(globalState).F2 = (_this).F25()
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))
			_ = _index126
			(_620_v0).ArraySet1(p2, (_index126).Int())
			var _644_v19 _dafny.Map
			_ = _644_v19
			_644_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F25())
			var _645_v20 _dafny.Map
			_ = _645_v20
			_645_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _646_v21 _dafny.Set
			_ = _646_v21
			_646_v21 = _dafny.SetOf(p1, p1, (func() bool {
				if (_645_v20).Contains(!(!((_this).F25()))) {
					return (_645_v20).Get(!(!((_this).F25()))).(bool)
				}
				return p1
			})(), false, p1)
			var _647_v22 D1
			_ = _647_v22
			_647_v22 = Companion_D1_.Create_DC2_(p1)
			var _648_v23 D1
			_ = _648_v23
			_648_v23 = Companion_D1_.Create_DC1_((_this).F25())
			var _649_v24 _dafny.Sequence
			_ = _649_v24
			_649_v24 = _dafny.SeqOf((_648_v23).Dtor_cf1())
			var _650_v25 _dafny.Set
			_ = _650_v25
			_650_v25 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_649_v24, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_649_v24).Cardinality()))).Uint32(), p1))
			var _651_v26 _dafny.Array
			_ = _651_v26
			var _nwElement0_24 bool = p1
			_ = _nwElement0_24
			var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(16))
			_ = _nw112
			(_nw112).ArraySet1(_nwElement0_24, 0)
			(_nw112).ArraySet1((_this).F25(), 1)
			(_nw112).ArraySet1((_this).F25(), 2)
			(_nw112).ArraySet1(!((_this).F25()), 3)
			(_nw112).ArraySet1(((_dafny.Zero).Minus((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int))).Cmp((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int)) <= 0, 4)
			(_nw112).ArraySet1(false, 5)
			(_nw112).ArraySet1((p2).Cmp(p2) == 0, 6)
			(_nw112).ArraySet1((func() bool {
				if (_644_v19).Contains((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int)) {
					return (_644_v19).Get((_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int)).(bool)
				}
				return (_this).F25()
			})(), 7)
			(_nw112).ArraySet1((_this).Fm3((_this).F25(), (_646_v21).Cardinality(), Companion_Default___.Fm2((_this).F25(), p2, (_this).Fm3(p1, (_620_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(994), _dafny.ArrayLen((_620_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-711), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), p2), globalState), (_this).F26(), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfInt64(926)), globalState), 8)
			(_nw112).ArraySet1((_this).F25(), 9)
			(_nw112).ArraySet1((_this).F25(), 10)
			(_nw112).ArraySet1((_647_v22).Dtor_cf2(), 11)
			(_nw112).ArraySet1((_this).F25(), 12)
			(_nw112).ArraySet1(p1, 13)
			(_nw112).ArraySet1((_650_v25).IsProperSubsetOf(_650_v25), 14)
			(_nw112).ArraySet1(p1, 15)
			_651_v26 = _nw112
			_651_v26 = _651_v26
		}
	}
}
func (_this *C3) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _652_v0 _dafny.MultiSet
		_ = _652_v0
		_652_v0 = _dafny.MultiSetOf(p1)
		var _653_v1 _dafny.MultiSet
		_ = _653_v1
		_653_v1 = _dafny.MultiSetOf(p1, _dafny.IntOfInt64(149), (_652_v0).Cardinality(), p1)
		var _654_i0 _dafny.Int
		_ = _654_i0
		_654_i0 = _dafny.Zero
		{
			for ((_652_v0).Union(_652_v0)).IsProperSubsetOf(_653_v1) {
				{
					if (_654_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_654_i0 = (_654_i0).Plus(_dafny.One)
					var _655_v2 _dafny.Array
					_ = _655_v2
					var _nwElement0_25 _dafny.Sequence = p0
					_ = _nwElement0_25
					var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(14))
					_ = _nw113
					(_nw113).ArraySet1(_nwElement0_25, 0)
					(_nw113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(768))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg31 _dafny.Int) interface{} {
							return coer31(arg31)
						}
					}(func(_656_i1 _dafny.Int) _dafny.CodePoint {
						return (_this).F26()
					}))), 1)
					(_nw113).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-96))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}(func(_657_i2 _dafny.Int) _dafny.CodePoint {
						return (_this).F26()
					})), 2)
					(_nw113).ArraySet1(p0, 3)
					(_nw113).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_this).F26()), 4)
					(_nw113).ArraySet1(p0, 5)
					(_nw113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.UnicodeSeqOfUtf8Bytes("dhnjxv")), 6)
					(_nw113).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(469))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg33 _dafny.Int) interface{} {
							return coer33(arg33)
						}
					}(func(_658_i3 _dafny.Int) _dafny.CodePoint {
						return (_this).F26()
					})), 7)
					(_nw113).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nyxcrvbu"), 8)
					(_nw113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg34 _dafny.Int) interface{} {
							return coer34(arg34)
						}
					}(func(_659_i4 _dafny.Int) _dafny.CodePoint {
						return (_this).F26()
					})), _dafny.UnicodeSeqOfUtf8Bytes("imrgbtikj")), 9)
					(_nw113).ArraySet1(p0, 10)
					(_nw113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dewtel"), _dafny.UnicodeSeqOfUtf8Bytes("fcxwj")), 11)
					(_nw113).ArraySet1((func() _dafny.Sequence {
						if (_this).F25() {
							return p0
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(304))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}(func(_660_i5 _dafny.Int) _dafny.CodePoint {
							return (_this).F26()
						}))
					})(), 12)
					(_nw113).ArraySet1(p0, 13)
					_655_v2 = _nw113
					var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))
					_ = _index127
					(_655_v2).ArraySet1(p0, (_index127).Int())
					var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))
					_ = _index128
					(_655_v2).ArraySet1(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_this).F26()), (_index128).Int())
					var _661_v3 _dafny.Map
					_ = _661_v3
					_661_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence)).Cardinality()), p1)
					var _662_v4 _dafny.Sequence
					_ = _662_v4
					_662_v4 = _dafny.SeqOf(_661_v3)
					var _663_v5 D2
					_ = _663_v5
					_663_v5 = Companion_D2_.Create_DC3_(_662_v4)
					var _664_v6 _dafny.MultiSet
					_ = _664_v6
					_664_v6 = _dafny.MultiSetOf((_this).F25())
					var _665_v7 _dafny.Map
					_ = _665_v7
					_665_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_664_v6, _662_v4)
					(globalState).F21 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_663_v5).Dtor_cf3(), (func() _dafny.Sequence {
						if (_665_v7).Contains(_664_v6) {
							return (_665_v7).Get(_664_v6).(_dafny.Sequence)
						}
						return _dafny.SeqOf(_661_v3)
					})())).Cardinality())
					var _666_v8 _dafny.Sequence
					_ = _666_v8
					_666_v8 = _dafny.SeqOf(p0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence), (_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence)), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence), (_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence))).Cardinality()))).Uint32(), (_this).F26()), (_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7((_this).F26(), globalState), _dafny.UnicodeSeqOfUtf8Bytes("h")), (_655_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))).Int()).(_dafny.Sequence))
					var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_655_v2), 0))
					_ = _index129
					(_655_v2).ArraySet1((_666_v8).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_666_v8).Cardinality()))).Uint32()).(_dafny.Sequence), (_index129).Int())
					var _667_v9 _dafny.Map
					_ = _667_v9
					_667_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), Companion_Default___.Fm2((_this).F25(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(887))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg36 _dafny.Int) interface{} {
							return coer36(arg36)
						}
					}(func(_668_i6 _dafny.Int) _dafny.CodePoint {
						return (_this).F26()
					}))).Cardinality()), (_this).F25(), _dafny.CodePoint('g'), globalState))
					var _669_v10 _dafny.Array
					_ = _669_v10
					var _len0_23 _dafny.Int = _dafny.IntOfInt64(4)
					_ = _len0_23
					var _nw114 _dafny.Array
					_ = _nw114
					if _len0_23.Cmp(_dafny.Zero) == 0 {
						_nw114 = _dafny.NewArray(_len0_23)
					} else {
						var _init23 func(_dafny.Int) bool = func(_670_i7 _dafny.Int) bool {
							return true
						}
						_ = _init23
						var _element0_23 = _init23(_dafny.Zero)
						_ = _element0_23
						_nw114 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
						(_nw114).ArraySet1(_element0_23, 0)
						var _nativeLen0_23 = (_len0_23).Int()
						_ = _nativeLen0_23
						for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
							(_nw114).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
						}
					}
					_669_v10 = _nw114
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_669_v10), 0))
					_ = _index130
					(_669_v10).ArraySet1((_this).F25(), (_index130).Int())
					var _671_v11 D1
					_ = _671_v11
					_671_v11 = Companion_D1_.Create_DC2_((_this).F25())
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_669_v10), 0))
					_ = _index131
					(_669_v10).ArraySet1((_671_v11).Dtor_cf2(), (_index131).Int())
					var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_669_v10), 0))
					_ = _index132
					var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_669_v10), 0))
					_ = _index133
					var _rhs105 _dafny.Map = _667_v9
					_ = _rhs105
					var _rhs106 bool = (_this).F25()
					_ = _rhs106
					var _rhs107 bool = Companion_Default___.Fm0((p1).Plus(p1), (p1).Minus(p1), (_this).F25(), globalState)
					_ = _rhs107
					var _rhs108 _dafny.Int = ((_dafny.Zero).Minus((_dafny.Zero).Minus(p1))).Minus(p1)
					_ = _rhs108
					var _rhs109 bool = true
					_ = _rhs109
					var _lhs92 *GlobalState = globalState
					_ = _lhs92
					var _lhs93 _dafny.Array = _669_v10
					_ = _lhs93
					var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_669_v10), 0))
					_ = _lhs94
					var _lhs95 _dafny.Array = _669_v10
					_ = _lhs95
					var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_669_v10), 0))
					_ = _lhs96
					_667_v9 = _rhs105
					_lhs92.F2 = _rhs106
					(_lhs93).ArraySet1(_rhs107, (_lhs94).Int())
					r1 = _rhs108
					(_lhs95).ArraySet1(_rhs109, (_lhs96).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _672_v12 _dafny.Array
		_ = _672_v12
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_24
		var _nw115 _dafny.Array
		_ = _nw115
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw115 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) _dafny.Int = (func(_673_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_674_i8 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_674_i8, _673_p1)
				}
			})(p1)
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw115 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw115).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw115).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_672_v12 = _nw115
		var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
		_ = _index134
		(_672_v12).ArraySet1(_dafny.IntOfInt64(249), (_index134).Int())
		var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
		_ = _index135
		(_672_v12).ArraySet1(p1, (_index135).Int())
		if (_this).F25() {
			r0 = (_this).F25()
			var _675_v13 _dafny.CodePoint
			_ = _675_v13
			_675_v13 = _dafny.CodePoint('w')
			var _676_v14 _dafny.Set
			_ = _676_v14
			_676_v14 = _dafny.SetOf(p0, _dafny.Companion_Sequence_.Concatenate(p0, p0), p0)
			var _677_v15 _dafny.Sequence
			_ = _677_v15
			_677_v15 = _dafny.SeqOf((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int))
			var _678_v16 _dafny.Int
			_ = _678_v16
			_678_v16 = _dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality())
			var _679_v17 _dafny.Map
			_ = _679_v17
			_679_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _678_v16)
			var _rhs110 bool = !(Companion_Default___.Fm0((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int))), (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), (_this).F25(), globalState))
			_ = _rhs110
			var _rhs111 _dafny.CodePoint = (_this).F26()
			_ = _rhs111
			var _rhs112 bool = ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Cmp((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)) == 0
			_ = _rhs112
			var _rhs113 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_677_v15).Cardinality()), Companion_Default___.SafeModuloInt((_677_v15).Select((Companion_Default___.SafeIndex((_679_v17).Cardinality(), _dafny.IntOfUint32((_677_v15).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_this).F26())).Cardinality())))
			_ = _rhs113
			var _rhs114 _dafny.Set = _676_v14
			_ = _rhs114
			var _lhs97 *GlobalState = globalState
			_ = _lhs97
			var _lhs98 *GlobalState = globalState
			_ = _lhs98
			_lhs97.F2 = _rhs110
			_675_v13 = _rhs111
			r0 = _rhs112
			_lhs98.F21 = _rhs113
			_676_v14 = _rhs114
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
			_ = _index136
			(_672_v12).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(609), (_this).F25())).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))), (_index136).Int())
			var _680_v18 _dafny.Array
			_ = _680_v18
			var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw116
			_680_v18 = _nw116
			var _681_v19 _dafny.Array
			_ = _681_v19
			var _nwElement0_26 _dafny.Array = _680_v18
			_ = _nwElement0_26
			var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(18))
			_ = _nw117
			(_nw117).ArraySet1(_nwElement0_26, 0)
			(_nw117).ArraySet1(_680_v18, 1)
			(_nw117).ArraySet1(_680_v18, 2)
			(_nw117).ArraySet1(_680_v18, 3)
			(_nw117).ArraySet1(_680_v18, 4)
			(_nw117).ArraySet1(_680_v18, 5)
			(_nw117).ArraySet1(_680_v18, 6)
			(_nw117).ArraySet1(_680_v18, 7)
			(_nw117).ArraySet1(_680_v18, 8)
			(_nw117).ArraySet1(_680_v18, 9)
			(_nw117).ArraySet1(_680_v18, 10)
			(_nw117).ArraySet1(_680_v18, 11)
			(_nw117).ArraySet1(_680_v18, 12)
			(_nw117).ArraySet1(_680_v18, 13)
			(_nw117).ArraySet1(_680_v18, 14)
			(_nw117).ArraySet1(_680_v18, 15)
			(_nw117).ArraySet1(_680_v18, 16)
			(_nw117).ArraySet1((func() _dafny.Array {
				if (_this).F25() {
					return _680_v18
				}
				return _680_v18
			})(), 17)
			_681_v19 = _nw117
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_681_v19), 0))
			_ = _index137
			(_681_v19).ArraySet1(_680_v18, (_index137).Int())
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_681_v19), 0))
			_ = _index138
			(_681_v19).ArraySet1(_680_v18, (_index138).Int())
			(globalState).F5 = (p1).Plus(p1)
		} else {
			var _682_v20 _dafny.Set
			_ = _682_v20
			_682_v20 = _dafny.SetOf(_dafny.IntOfInt64(586), p1)
			var _683_v21 _dafny.Sequence
			_ = _683_v21
			_683_v21 = _dafny.SeqOf((_682_v20).Cardinality())
			if (func() bool {
				if !((_dafny.IntOfUint32((_683_v21).Cardinality())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((p2).Cardinality()))) < 0) {
					return ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Cmp((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)) != 0
				}
				return !((_this).F25()) || (false)
			})() {
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
				_ = _index139
				(_672_v12).ArraySet1((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), (_index139).Int())
				r0 = ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(496)) >= 0
				(globalState).F18 = _dafny.IntOfUint32((p0).Cardinality())
				(globalState).F2 = false
				var _684_v22 _dafny.CodePoint
				_ = _684_v22
				_684_v22 = _dafny.CodePoint('t')
				var _rhs115 _dafny.CodePoint = _684_v22
				_ = _rhs115
				var _rhs116 bool = (_this).F25()
				_ = _rhs116
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				_684_v22 = _rhs115
				_lhs99.F2 = _rhs116
			} else {
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
				_ = _index140
				(_672_v12).ArraySet1(p1, (_index140).Int())
				(globalState).F2 = (_this).F25()
				var _685_v23 T1
				_ = _685_v23
				var _nw118 *C2 = New_C2_()
				_ = _nw118
				_nw118.Ctor__((_this).F26(), (_this).F27(), Companion_Default___.Fm0((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), (_this).F25(), globalState))
				_685_v23 = _nw118
				var _686_v24 _dafny.Array
				_ = _686_v24
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_25
				var _nw119 _dafny.Array
				_ = _nw119
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw119 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) bool = func(_687_i9 _dafny.Int) bool {
						return (_this).F25()
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw119 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw119).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw119).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_686_v24 = _nw119
				(globalState).F5 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_685_v23, _686_v24)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_685_v23, _686_v24))).Cardinality()
				var _688_v25 _dafny.MultiSet
				_ = _688_v25
				_688_v25 = _dafny.MultiSetOf(true, (_this).F25())
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
				_ = _index141
				(_672_v12).ArraySet1((func() _dafny.Int {
					if (_685_v23).F25() {
						return ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Plus((func() _dafny.Int {
							if (_688_v25).Contains((_this).F25()) {
								return (_688_v25).Multiplicity((_this).F25())
							}
							return (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)
						})())
					}
					return p1
				})(), (_index141).Int())
				var _689_v26 _dafny.Array
				_ = _689_v26
				var _nw120 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
				_ = _nw120
				_689_v26 = _nw120
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_689_v26), 0))
				_ = _index142
				(_689_v26).ArraySet1(_685_v23, (_index142).Int())
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_689_v26), 0))
				_ = _index143
				var _rhs117 bool = (_this).F25()
				_ = _rhs117
				var _rhs118 T1 = _685_v23
				_ = _rhs118
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				var _lhs101 _dafny.Array = _689_v26
				_ = _lhs101
				var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_689_v26), 0))
				_ = _lhs102
				_lhs100.F2 = _rhs117
				(_lhs101).ArraySet1(_rhs118, (_lhs102).Int())
			}
			var _690_v27 _dafny.Set
			_ = _690_v27
			_690_v27 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}(func(_691_i10 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F25())).Cardinality()
			})), _dafny.SeqOf((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)), _683_v21, _683_v21)
			if (_690_v27).IsDisjointFrom(_690_v27) {
				(globalState).F5 = p1
				(globalState).F18 = ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Plus((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int))
				var _692_v28 _dafny.Map
				_ = _692_v28
				_692_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F25())
				var _693_v29 D3
				_ = _693_v29
				_693_v29 = Companion_D3_.Create_DC7_(_692_v28)
				var _694_v30 _dafny.Array
				_ = _694_v30
				var _nwElement0_27 D3 = Companion_Default___.Fm25(globalState)
				_ = _nwElement0_27
				var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(7))
				_ = _nw121
				(_nw121).ArraySet1(_nwElement0_27, 0)
				(_nw121).ArraySet1(Companion_Default___.Fm25(globalState), 1)
				(_nw121).ArraySet1(Companion_D3_.Create_DC7_(_692_v28), 2)
				(_nw121).ArraySet1(Companion_Default___.Fm25(globalState), 3)
				(_nw121).ArraySet1(_693_v29, 4)
				(_nw121).ArraySet1(_693_v29, 5)
				(_nw121).ArraySet1(_693_v29, 6)
				_694_v30 = _nw121
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_694_v30), 0))
				_ = _index144
				(_694_v30).ArraySet1(_693_v29, (_index144).Int())
				var _pat_let_tv11 = _692_v28
				_ = _pat_let_tv11
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_694_v30), 0))
				_ = _index145
				(_694_v30).ArraySet1(func(_pat_let9_0 D3) D3 {
					return func(_695_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let10_0 _dafny.Map) D3 {
							return func(_696_dt__update_hcf11_h0 _dafny.Map) D3 {
								return Companion_D3_.Create_DC7_(_696_dt__update_hcf11_h0)
							}(_pat_let10_0)
						}(_pat_let_tv11)
					}(_pat_let9_0)
				}(_693_v29), (_index145).Int())
				var _697_v31 _dafny.Map
				_ = _697_v31
				_697_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_683_v21, p0)
				_697_v31 = (_697_v31).Update(_683_v21, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yjmlwg"), _dafny.UnicodeSeqOfUtf8Bytes("x")))
				var _698_v32 *C2
				_ = _698_v32
				var _nw122 *C2 = New_C2_()
				_ = _nw122
				_nw122.Ctor__((_this).F26(), (_this).F27(), (_this).F25())
				_698_v32 = _nw122
			} else {
				(globalState).F5 = _dafny.IntOfInt64(807)
				(globalState).F5 = p1
				var _699_v33 _dafny.Set
				_ = _699_v33
				_699_v33 = _dafny.SetOf((_this).F25())
				(globalState).F18 = (_699_v33).Cardinality()
				var _700_v34 T1
				_ = _700_v34
				var _nw123 *C2 = New_C2_()
				_ = _nw123
				_nw123.Ctor__((_this).F26(), (_this).F27(), (_this).F25())
				_700_v34 = _nw123
				var _rhs119 bool = (_this).F25()
				_ = _rhs119
				var _rhs120 T1 = (func() T1 {
					if (_this).F25() {
						return _700_v34
					}
					return _700_v34
				})()
				_ = _rhs120
				var _lhs103 *GlobalState = globalState
				_ = _lhs103
				_lhs103.F2 = _rhs119
				_700_v34 = _rhs120
				(globalState).F2 = (_this).F25()
			}
			(globalState).F2 = (_this).F25()
			var _701_v35 _dafny.Array
			_ = _701_v35
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_26
			var _nw124 _dafny.Array
			_ = _nw124
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw124 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) bool = func(_702_i11 _dafny.Int) bool {
					return false
				}
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw124 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw124).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw124).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_701_v35 = _nw124
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))
			_ = _index146
			(_701_v35).ArraySet1(false, (_index146).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))
			_ = _index147
			(_701_v35).ArraySet1(((_this).F25()) || (!((_this).F25())), (_index147).Int())
			var _703_v36 D5
			_ = _703_v36
			_703_v36 = Companion_D5_.Create_DC15_((_this).F25(), _701_v35, p0, false)
			var _source7 D5 = _703_v36
			_ = _source7
			if _source7.Is_DC15() {
				var _704___mcc_h0 bool = _source7.Get_().(D5_DC15).Cf25
				_ = _704___mcc_h0
				var _705___mcc_h1 _dafny.Array = _source7.Get_().(D5_DC15).Cf26
				_ = _705___mcc_h1
				var _706___mcc_h2 _dafny.Sequence = _source7.Get_().(D5_DC15).Cf27
				_ = _706___mcc_h2
				var _707___mcc_h3 bool = _source7.Get_().(D5_DC15).Cf28
				_ = _707___mcc_h3
				var _708_cf28 bool = _707___mcc_h3
				_ = _708_cf28
				var _709_cf27 _dafny.Sequence = _706___mcc_h2
				_ = _709_cf27
				var _710_cf26 _dafny.Array = _705___mcc_h1
				_ = _710_cf26
				var _711_cf25 bool = _704___mcc_h0
				_ = _711_cf25
				(globalState).F18 = p1
				(globalState).F22 = Companion_Default___.Fm2((_701_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))).Int()).(bool), (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), (_this).F25(), _dafny.CodePoint('f'), globalState)
				var _712_v37 *C2
				_ = _712_v37
				var _nw125 *C2 = New_C2_()
				_ = _nw125
				_nw125.Ctor__((_this).F26(), (_this).F27(), _711_cf25)
				_712_v37 = _nw125
				var _713_v38 D3
				_ = _713_v38
				_713_v38 = Companion_D3_.Create_DC9_(p1, _dafny.CodePoint('r'), p1)
				var _714_v39 D3
				_ = _714_v39
				_714_v39 = Companion_D3_.Create_DC10_(_713_v38)
				var _715_v40 _dafny.Map
				_ = _715_v40
				_715_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F25())
				var _pat_let_tv12 = _715_v40
				_ = _pat_let_tv12
				_714_v39 = func(_pat_let11_0 D3) D3 {
					return func(_716_dt__update__tmp_h2 D3) D3 {
						return func(_pat_let12_0 D3) D3 {
							return func(_717_dt__update_hcf16_h0 D3) D3 {
								return Companion_D3_.Create_DC10_(_717_dt__update_hcf16_h0)
							}(_pat_let12_0)
						}(Companion_D3_.Create_DC7_(_pat_let_tv12))
					}(_pat_let11_0)
				}(_714_v39)
			} else if _source7.Is_DC14() {
				var _718___mcc_h4 _dafny.Map = _source7.Get_().(D5_DC14).Cf24
				_ = _718___mcc_h4
				var _719_cf24 _dafny.Map = _718___mcc_h4
				_ = _719_cf24
				var _720_v41 D2
				_ = _720_v41
				_720_v41 = Companion_D2_.Create_DC5_(_dafny.IntOfInt64(204), true, (_701_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))).Int()).(bool), (_this).F25())
				var _721_v42 _dafny.Map
				_ = _721_v42
				_721_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int), _701_v35)
				var _722_v43 _dafny.Map
				_ = _722_v43
				_722_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_720_v41).Dtor_cf6(), _721_v42)
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))
				_ = _index148
				(_672_v12).ArraySet1((func() _dafny.Int {
					if (_701_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))).Int()).(bool) {
						return (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)
					}
					return ((func() _dafny.Map {
						if (_722_v43).Contains(p1) {
							return (_722_v43).Get(p1).(_dafny.Map)
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(171), _701_v35)
					})()).Cardinality()
				})(), (_index148).Int())
				(globalState).F18 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(335)).Plus((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)), ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)))))
				var _723_v44 *C2
				_ = _723_v44
				var _nw126 *C2 = New_C2_()
				_ = _nw126
				_nw126.Ctor__((_this).F26(), (_this).F27(), (_this).F25())
				_723_v44 = _nw126
				var _724_v45 _dafny.Map
				_ = _724_v45
				_724_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_v44, _dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(44))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}(func(_725_i12 _dafny.Int) _dafny.CodePoint {
					return (_this).F26()
				}))))
				_724_v45 = _724_v45
				var _726_v46 _dafny.Sequence
				_ = _726_v46
				_726_v46 = _dafny.UnicodeSeqOfUtf8Bytes("if")
				_726_v46 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.UnicodeSeqOfUtf8Bytes("xdalhwjvr")), _726_v46)
			} else {
				var _727___mcc_h5 D5 = _source7.Get_().(D5_DC16).Cf29
				_ = _727___mcc_h5
				var _728_cf29 D5 = _727___mcc_h5
				_ = _728_cf29
				(globalState).F2 = ((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Cmp((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)) != 0
				(globalState).F5 = (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)
				var _729_v47 _dafny.Map
				_ = _729_v47
				_729_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_701_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))).Int()).(bool), p0)
				(globalState).F5 = (_dafny.Zero).Minus((_683_v21).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_729_v47).Cardinality()), _dafny.IntOfUint32((_683_v21).Cardinality()))).Uint32()).(_dafny.Int))
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_701_v35), 0))
				_ = _index149
				(_701_v35).ArraySet1(!((_652_v0).Update(p1, Companion_Default___.Abs((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)))).Contains(p1), (_index149).Int())
			}
		}
		r1 = (((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)).Times((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int))).Plus(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((p0).Cardinality())))
		r1 = (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)
		r0 = (_this).F25()
		var _730_v49 D8
		_ = _730_v49
		_730_v49 = Companion_D8_.Create_DC25_(_653_v1)
		var _731_v50 D6
		_ = _731_v50
		_731_v50 = Companion_D6_.Create_DC17_(func() _dafny.Map {
			var _coll24 = _dafny.NewMapBuilder()
			_ = _coll24
			for _iter28 := _dafny.Iterate((((_730_v49).Dtor_cf48()).Update(p1, Companion_Default___.Abs((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)))).Elements()); ; {
				_compr_24, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _732_v48 _dafny.Int
				_732_v48 = interface{}(_compr_24).(_dafny.Int)
				if (((_730_v49).Dtor_cf48()).Update(p1, Companion_Default___.Abs((_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)))).Contains(_732_v48) {
					_coll24.Add(Companion_Default___.SafeModuloInt(_732_v48, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg39 _dafny.Int) interface{} {
							return coer39(arg39)
						}
					}(func(_733_i13 _dafny.Int) _dafny.CodePoint {
						return (_this).F26()
					}))).Cardinality())), p1)
				}
			}
			return _coll24.ToMap()
		}())
		r0 = func(_source8 D6) bool {
			if _source8.Is_DC18() {
				var _734___mcc_h6 _dafny.Int = _source8.Get_().(D6_DC18).Cf31
				_ = _734___mcc_h6
				var _735___mcc_h7 bool = _source8.Get_().(D6_DC18).Cf32
				_ = _735___mcc_h7
				var _736___mcc_h8 bool = _source8.Get_().(D6_DC18).Cf33
				_ = _736___mcc_h8
				var _737___mcc_h9 bool = _source8.Get_().(D6_DC18).Cf34
				_ = _737___mcc_h9
				var _738___mcc_h10 _dafny.Int = _source8.Get_().(D6_DC18).Cf35
				_ = _738___mcc_h10
				var _739_cf35 _dafny.Int = _738___mcc_h10
				_ = _739_cf35
				var _740_cf34 bool = _737___mcc_h9
				_ = _740_cf34
				var _741_cf33 bool = _736___mcc_h8
				_ = _741_cf33
				var _742_cf32 bool = _735___mcc_h7
				_ = _742_cf32
				var _743_cf31 _dafny.Int = _734___mcc_h6
				_ = _743_cf31
				return (_741_cf33) || (_740_cf34)
			} else if _source8.Is_DC17() {
				var _744___mcc_h11 _dafny.Map = _source8.Get_().(D6_DC17).Cf30
				_ = _744___mcc_h11
				var _745_cf30 _dafny.Map = _744___mcc_h11
				_ = _745_cf30
				return true
			} else {
				var _746___mcc_h12 D6 = _source8.Get_().(D6_DC19).Cf36
				_ = _746___mcc_h12
				var _747_cf36 D6 = _746___mcc_h12
				_ = _747_cf36
				return (_this).F25()
			}
		}(_731_v50)
		r1 = (_672_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_672_v12), 0))).Int()).(_dafny.Int)
		return r0, r1
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
