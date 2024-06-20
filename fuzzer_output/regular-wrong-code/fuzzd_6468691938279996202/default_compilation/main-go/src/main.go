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
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true, false, false, false, false)))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(984))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("iwtniohu"), _dafny.UnicodeSeqOfUtf8Bytes("aa")))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('l')
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-104)), _dafny.IntOfInt64(964))).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('f'))).Cardinality()), _dafny.IntOfInt64(57))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rjys")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(907)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-984), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bsuch")).Cardinality()))).Cardinality()), false)
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality())))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-992))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(970)
	}))).Cardinality())))).Cardinality())).Cardinality()), _dafny.IntOfInt64(782)), (_dafny.CodePoint('u')))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("frnc"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-34))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	})))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('j')
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(false, true))).Union((_dafny.MultiSetOf(false, true)).Intersection(_dafny.MultiSetOf(true, false)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(false), _dafny.SetOf(false, true, false, true, false)), _dafny.SeqOf(_dafny.SetOf(true), _dafny.SetOf(!(!(false))), _dafny.SetOf(false))), (Companion_D34_.Create_DC84_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Set {
		return _dafny.SetOf(true, true)
	})))).Dtor_cf145())
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(751), (_dafny.MultiSetOf(false)).Cardinality())).Union(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(935), true), true)).Cardinality()))).Cardinality()), (_dafny.MultiSetOf(false)).Cardinality()))).Difference(_dafny.SetOf(_dafny.IntOfInt64(82), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(316), _dafny.IntOfInt64(-425))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(316)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(-425)) < 0) {
				_coll0.Add((_4_v0).Times((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(833), _dafny.CodePoint('b'))).Cardinality())), false)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(874))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("awte"), _dafny.UnicodeSeqOfUtf8Bytes("skgl"))
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("yr")
	}
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) D1 {
	var _source0 D1 = Companion_D1_.Create_DC4_(false, !(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jee")).Cardinality()))
	_ = _source0
	if _source0.Is_DC3() {
		var _5___mcc_h0 _dafny.Sequence = _source0.Get_().(D1_DC3).Cf5
		_ = _5___mcc_h0
		var _6_cf5 _dafny.Sequence = _5___mcc_h0
		_ = _6_cf5
		return Companion_D1_.Create_DC2_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(126))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}((func(_7_cf5 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_8_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_7_cf5).Cardinality())
			}
		})(_6_cf5))))
	} else if _source0.Is_DC4() {
		var _9___mcc_h1 bool = _source0.Get_().(D1_DC4).Cf6
		_ = _9___mcc_h1
		var _10___mcc_h2 bool = _source0.Get_().(D1_DC4).Cf7
		_ = _10___mcc_h2
		var _11___mcc_h3 _dafny.Int = _source0.Get_().(D1_DC4).Cf8
		_ = _11___mcc_h3
		var _12_cf8 _dafny.Int = _11___mcc_h3
		_ = _12_cf8
		var _13_cf7 bool = _10___mcc_h2
		_ = _13_cf7
		var _14_cf6 bool = _9___mcc_h1
		_ = _14_cf6
		return Companion_D1_.Create_DC2_(_dafny.SeqOf(_dafny.IntOfInt64(246)))
	} else if _source0.Is_DC2() {
		var _15___mcc_h4 _dafny.Sequence = _source0.Get_().(D1_DC2).Cf4
		_ = _15___mcc_h4
		var _16_cf4 _dafny.Sequence = _15___mcc_h4
		_ = _16_cf4
		return Companion_D1_.Create_DC2_(_16_cf4)
	} else {
		var _17___mcc_h5 D1 = _source0.Get_().(D1_DC5).Cf9
		_ = _17___mcc_h5
		var _18_cf9 D1 = _17___mcc_h5
		_ = _18_cf9
		return Companion_D1_.Create_DC2_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gkgim")).Cardinality()), _dafny.IntOfInt64(-329)))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fihalg"), _dafny.UnicodeSeqOfUtf8Bytes("cileu")), _dafny.UnicodeSeqOfUtf8Bytes("urjn"))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D0 {
	if false {
		return Companion_D0_.Create_DC0_(!(true))
	} else {
		return Companion_D0_.Create_DC0_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Intersection((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Intersection(func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _19_v0 _dafny.Map
			_19_v0 = interface{}(_compr_1).(_dafny.Map)
			if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Contains(_19_v0) {
				_coll1.Add(_19_v0)
			}
		}
		return _coll1.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(510))).Intersection((_dafny.SetOf(_dafny.IntOfInt64(239))).Union(_dafny.SetOf(_dafny.IntOfInt64(179))))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.CodePoint {
	var _source1 _dafny.Map = (func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("alriycn")).Cardinality()), _dafny.IntOfInt64(463))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(896))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))).Cardinality()), _dafny.IntOfInt64(486))
	})()
	_ = _source1
	var _21___mcc_h0 _dafny.Map = _source1
	_ = _21___mcc_h0
	var _22_cf45 _dafny.Map = _21___mcc_h0
	_ = _22_cf45
	return _dafny.CodePoint('c')
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 bool, p2 _dafny.Set, p3 _dafny.Set, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(_dafny.CodePoint('o')))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(_dafny.CodePoint('v'))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(_dafny.CodePoint('x'), _dafny.CodePoint('h'), _dafny.CodePoint('o'), _dafny.CodePoint('s'))))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true)
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	return (Companion_D6_.Create_DC18_(_dafny.CodePoint('l'), _dafny.IntOfInt64(486), (_dafny.SetOf(false)).Cardinality())).Dtor_cf34()
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	var _source2 D6 = Companion_D6_.Create_DC18_(_dafny.CodePoint('p'), (_dafny.SetOf(_dafny.IntOfInt64(-530))).Cardinality(), (_dafny.MultiSetOf(false)).Cardinality())
	_ = _source2
	if _source2.Is_DC17() {
		var _23___mcc_h0 bool = _source2.Get_().(D6_DC17).Cf29
		_ = _23___mcc_h0
		var _24___mcc_h1 bool = _source2.Get_().(D6_DC17).Cf30
		_ = _24___mcc_h1
		var _25___mcc_h2 _dafny.Int = _source2.Get_().(D6_DC17).Cf31
		_ = _25___mcc_h2
		var _26___mcc_h3 _dafny.Int = _source2.Get_().(D6_DC17).Cf32
		_ = _26___mcc_h3
		var _27___mcc_h4 _dafny.Int = _source2.Get_().(D6_DC17).Cf33
		_ = _27___mcc_h4
		var _28_cf33 _dafny.Int = _27___mcc_h4
		_ = _28_cf33
		var _29_cf32 _dafny.Int = _26___mcc_h3
		_ = _29_cf32
		var _30_cf31 _dafny.Int = _25___mcc_h2
		_ = _30_cf31
		var _31_cf30 bool = _24___mcc_h1
		_ = _31_cf30
		var _32_cf29 bool = _23___mcc_h0
		_ = _32_cf29
		return _30_cf31
	} else if _source2.Is_DC18() {
		var _33___mcc_h5 _dafny.CodePoint = _source2.Get_().(D6_DC18).Cf34
		_ = _33___mcc_h5
		var _34___mcc_h6 _dafny.Int = _source2.Get_().(D6_DC18).Cf35
		_ = _34___mcc_h6
		var _35___mcc_h7 _dafny.Int = _source2.Get_().(D6_DC18).Cf36
		_ = _35___mcc_h7
		var _36_cf36 _dafny.Int = _35___mcc_h7
		_ = _36_cf36
		var _37_cf35 _dafny.Int = _34___mcc_h6
		_ = _37_cf35
		var _38_cf34 _dafny.CodePoint = _33___mcc_h5
		_ = _38_cf34
		return _37_cf35
	} else if _source2.Is_DC19() {
		var _39___mcc_h8 bool = _source2.Get_().(D6_DC19).Cf37
		_ = _39___mcc_h8
		var _40___mcc_h9 _dafny.CodePoint = _source2.Get_().(D6_DC19).Cf38
		_ = _40___mcc_h9
		var _41_cf38 _dafny.CodePoint = _40___mcc_h9
		_ = _41_cf38
		var _42_cf37 bool = _39___mcc_h8
		_ = _42_cf37
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-815), _dafny.IntOfInt64(685))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _43_v0 _dafny.Int
				_43_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(-815)).Cmp(_43_v0) <= 0) && ((_43_v0).Cmp(_dafny.IntOfInt64(685)) < 0) {
					_coll2.Add(Companion_Default___.SafeModuloInt(_43_v0, _dafny.IntOfInt64(371)), (Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fv")).Cardinality()), _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(847), _41_cf38)).Cardinality()), _42_cf37)).Dtor_cf18())
				}
			}
			return _coll2.ToMap()
		}())).Cardinality()), _dafny.IntOfInt64(618))
	} else if _source2.Is_DC16() {
		var _44___mcc_h10 _dafny.Map = _source2.Get_().(D6_DC16).Cf28
		_ = _44___mcc_h10
		var _45_cf28 _dafny.Map = _44___mcc_h10
		_ = _45_cf28
		return (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("dgejc"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.UnicodeSeqOfUtf8Bytes("vben")))).Cardinality())
	} else {
		var _46___mcc_h11 D6 = _source2.Get_().(D6_DC20).Cf39
		_ = _46___mcc_h11
		var _47_cf39 D6 = _46___mcc_h11
		_ = _47_cf39
		return _dafny.IntOfInt64(408)
	}
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((func() bool {
		if true {
			return false
		}
		return false
	})())
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, globalState *GlobalState) D3 {
	var _source3 D14 = Companion_D14_.Create_DC34_(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer6 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_48_i0 _dafny.Int) D4 {
			return Companion_D4_.Create_DC11_(_dafny.IntOfInt64(-429), _dafny.SetOf(_dafny.IntOfInt64(694), _dafny.IntOfInt64(859)), true)
		})), _dafny.CodePoint('q'))).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _49_v0 _dafny.Sequence
			_49_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer7 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_48_i0 _dafny.Int) D4 {
				return Companion_D4_.Create_DC11_(_dafny.IntOfInt64(-429), _dafny.SetOf(_dafny.IntOfInt64(694), _dafny.IntOfInt64(859)), true)
			})), _dafny.CodePoint('q'))).Contains(_49_v0) {
				_coll3.Add(_49_v0, _dafny.IntOfInt64(-622))
			}
		}
		return _coll3.ToMap()
	}())
	_ = _source3
	if _source3.Is_DC35() {
		var _50___mcc_h0 _dafny.Int = _source3.Get_().(D14_DC35).Cf62
		_ = _50___mcc_h0
		var _51___mcc_h1 bool = _source3.Get_().(D14_DC35).Cf63
		_ = _51___mcc_h1
		var _52___mcc_h2 _dafny.CodePoint = _source3.Get_().(D14_DC35).Cf64
		_ = _52___mcc_h2
		var _53_cf64 _dafny.CodePoint = _52___mcc_h2
		_ = _53_cf64
		var _54_cf63 bool = _51___mcc_h1
		_ = _54_cf63
		var _55_cf62 _dafny.Int = _50___mcc_h0
		_ = _55_cf62
		return Companion_D3_.Create_DC8_(_54_cf63, _54_cf63)
	} else if _source3.Is_DC36() {
		var _56___mcc_h3 bool = _source3.Get_().(D14_DC36).Cf65
		_ = _56___mcc_h3
		var _57___mcc_h4 _dafny.Sequence = _source3.Get_().(D14_DC36).Cf66
		_ = _57___mcc_h4
		var _58___mcc_h5 _dafny.Sequence = _source3.Get_().(D14_DC36).Cf67
		_ = _58___mcc_h5
		var _59_cf67 _dafny.Sequence = _58___mcc_h5
		_ = _59_cf67
		var _60_cf66 _dafny.Sequence = _57___mcc_h4
		_ = _60_cf66
		var _61_cf65 bool = _56___mcc_h3
		_ = _61_cf65
		return Companion_D3_.Create_DC8_(!(_61_cf65), _61_cf65)
	} else {
		var _62___mcc_h6 _dafny.Map = _source3.Get_().(D14_DC34).Cf61
		_ = _62___mcc_h6
		var _63_cf61 _dafny.Map = _62___mcc_h6
		_ = _63_cf61
		return Companion_D3_.Create_DC8_(!(false), !(true))
	}
}
func (_static *CompanionStruct_Default___) Fm36(globalState *GlobalState) _dafny.CodePoint {
	var _source4 D6 = Companion_D6_.Create_DC20_(Companion_D6_.Create_DC18_(_dafny.CodePoint('m'), _dafny.IntOfInt64(779), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(555))).Cardinality()))
	_ = _source4
	if _source4.Is_DC17() {
		var _64___mcc_h0 bool = _source4.Get_().(D6_DC17).Cf29
		_ = _64___mcc_h0
		var _65___mcc_h1 bool = _source4.Get_().(D6_DC17).Cf30
		_ = _65___mcc_h1
		var _66___mcc_h2 _dafny.Int = _source4.Get_().(D6_DC17).Cf31
		_ = _66___mcc_h2
		var _67___mcc_h3 _dafny.Int = _source4.Get_().(D6_DC17).Cf32
		_ = _67___mcc_h3
		var _68___mcc_h4 _dafny.Int = _source4.Get_().(D6_DC17).Cf33
		_ = _68___mcc_h4
		var _69_cf33 _dafny.Int = _68___mcc_h4
		_ = _69_cf33
		var _70_cf32 _dafny.Int = _67___mcc_h3
		_ = _70_cf32
		var _71_cf31 _dafny.Int = _66___mcc_h2
		_ = _71_cf31
		var _72_cf30 bool = _65___mcc_h1
		_ = _72_cf30
		var _73_cf29 bool = _64___mcc_h0
		_ = _73_cf29
		return _dafny.CodePoint('o')
	} else if _source4.Is_DC18() {
		var _74___mcc_h5 _dafny.CodePoint = _source4.Get_().(D6_DC18).Cf34
		_ = _74___mcc_h5
		var _75___mcc_h6 _dafny.Int = _source4.Get_().(D6_DC18).Cf35
		_ = _75___mcc_h6
		var _76___mcc_h7 _dafny.Int = _source4.Get_().(D6_DC18).Cf36
		_ = _76___mcc_h7
		var _77_cf36 _dafny.Int = _76___mcc_h7
		_ = _77_cf36
		var _78_cf35 _dafny.Int = _75___mcc_h6
		_ = _78_cf35
		var _79_cf34 _dafny.CodePoint = _74___mcc_h5
		_ = _79_cf34
		return _dafny.CodePoint('x')
	} else if _source4.Is_DC19() {
		var _80___mcc_h8 bool = _source4.Get_().(D6_DC19).Cf37
		_ = _80___mcc_h8
		var _81___mcc_h9 _dafny.CodePoint = _source4.Get_().(D6_DC19).Cf38
		_ = _81___mcc_h9
		var _82_cf38 _dafny.CodePoint = _81___mcc_h9
		_ = _82_cf38
		var _83_cf37 bool = _80___mcc_h8
		_ = _83_cf37
		return _dafny.CodePoint('g')
	} else if _source4.Is_DC16() {
		var _84___mcc_h10 _dafny.Map = _source4.Get_().(D6_DC16).Cf28
		_ = _84___mcc_h10
		var _85_cf28 _dafny.Map = _84___mcc_h10
		_ = _85_cf28
		return _dafny.CodePoint('s')
	} else {
		var _86___mcc_h11 D6 = _source4.Get_().(D6_DC20).Cf39
		_ = _86___mcc_h11
		var _87_cf39 D6 = _86___mcc_h11
		_ = _87_cf39
		return _dafny.CodePoint('q')
	}
}
func (_static *CompanionStruct_Default___) Fm37(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(972))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_88_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(929)
			}))
		}
		return _dafny.SeqOf(_dafny.IntOfInt64(342), _dafny.IntOfInt64(642))
	})(), _dafny.SeqOf(_dafny.IntOfInt64(208)))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, false, !(false)), _dafny.SeqOf(false, false, false))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(true, false, true, !(true), true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(335), !(true))).Cardinality()), (func() _dafny.Int {
		if false {
			return (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(-919))).Cardinality())
		}
		return _dafny.IntOfInt64(247)
	})())
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Sequence, p1 D5, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(188), _dafny.IntOfInt64(-896)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(894), _dafny.IntOfInt64(511))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(698))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_89_i0 _dafny.Int) _dafny.Map {
		return func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(631))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(575))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_90_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			}))).Cardinality()))).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _91_v0 _dafny.Int
				_91_v0 = interface{}(_compr_4).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(631))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(575))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_90_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()))).Contains(_91_v0) {
					_coll4.Add((_91_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Cardinality()), _dafny.IntOfInt64(696))
				}
			}
			return _coll4.ToMap()
		}()
	})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(587), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cldmyuqk")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(179)))))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if true {
		return _dafny.CodePoint('y')
	} else {
		return _dafny.CodePoint('b')
	}
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.MultiSetOf((Companion_D17_.Create_DC44_(_dafny.IntOfInt64(357), _dafny.SeqOf(false), _dafny.CodePoint('t'), true, _dafny.IntOfInt64(983))).Dtor_cf84())).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _92_v0 _dafny.Int
			_92_v0 = interface{}(_compr_5).(_dafny.Int)
			if (_dafny.MultiSetOf((Companion_D17_.Create_DC44_(_dafny.IntOfInt64(357), _dafny.SeqOf(false), _dafny.CodePoint('t'), true, _dafny.IntOfInt64(983))).Dtor_cf84())).Contains(_92_v0) {
				_coll5.Add((_92_v0).Plus(_dafny.IntOfInt64(-17)), !(true))
			}
		}
		return _coll5.ToMap()
	}()).Merge((Companion_D35_.Create_DC87_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((Companion_D15_.Create_DC38_(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_93_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), true, !(true), _dafny.CodePoint('i'))).Dtor_cf70()).Cardinality()), true))).Dtor_cf149())
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.UnicodeSeqOfUtf8Bytes("wgb")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lkwgqmac"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_94_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	}))))
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfInt64(347), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ye")).Cardinality())))).Intersection((_dafny.MultiSetOf(_dafny.IntOfInt64(-658))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(42))))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, globalState *GlobalState) _dafny.Sequence {
	if !(true) {
		return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bvpnsv"), _dafny.UnicodeSeqOfUtf8Bytes("f"))
	} else {
		return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(420))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_95_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})), _dafny.UnicodeSeqOfUtf8Bytes("ch"))
	}
}
func (_static *CompanionStruct_Default___) Fm48(p0 bool, globalState *GlobalState) _dafny.Map {
	if (false) || (true) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true))
	}
}
func (_static *CompanionStruct_Default___) Fm49(p0 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC17_(true, (_dafny.MultiSetOf(_dafny.IntOfInt64(-731))).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_96_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), (func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('r')))).Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _97_v1 _dafny.Sequence
				_97_v1 = interface{}(_compr_7).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('r'))), _97_v1) {
					_coll7.Add(_97_v1, false)
				}
			}
			return _coll7.ToMap()
		}()).Cardinality())).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _98_v0 _dafny.Sequence
			_98_v0 = interface{}(_compr_6).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}(func(_96_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), (func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('r')))).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _97_v1 _dafny.Sequence
					_97_v1 = interface{}(_compr_8).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('r'))), _97_v1) {
						_coll8.Add(_97_v1, false)
					}
				}
				return _coll8.ToMap()
			}()).Cardinality())).Contains(_98_v0) {
				_coll6.Add(_98_v0, _dafny.IntOfInt64(-740))
			}
		}
		return _coll6.ToMap()
	}()).Cardinality()))), _dafny.IntOfInt64(330), (_dafny.IntOfInt64(3)).Times(_dafny.IntOfInt64(285)), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-341), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sn")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm51(globalState *GlobalState) D14 {
	if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(106))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_99_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-181)
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf((_dafny.CodePoint('d')))).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(516))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_100_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})))).Cardinality())).Cardinality()))) {
		return Companion_D14_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-130))).Uint32(), func(coer19 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_101_i2 _dafny.Int) D4 {
			return Companion_D4_.Create_DC11_(_dafny.IntOfInt64(-556), _dafny.SetOf(_dafny.IntOfInt64(253)), true)
		})), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_102_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		}))).Cardinality())))
	} else {
		return Companion_D14_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D4_.Create_DC11_(_dafny.IntOfInt64(636), func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(712), _dafny.IntOfInt64(679))); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _103_v0 _dafny.Int
				_103_v0 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(712)).Cmp(_103_v0) <= 0) && ((_103_v0).Cmp(_dafny.IntOfInt64(679)) < 0) {
					_coll9.Add((_103_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqsvf")).Cardinality())))
				}
			}
			return _coll9.ToSet()
		}(), true)), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true, true, !(true))).Cardinality()))).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, globalState *GlobalState) D8 {
	if !(((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(419), _dafny.IntOfInt64(579))); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _104_v2 _dafny.Int
					_104_v2 = interface{}(_compr_12).(_dafny.Int)
					if ((_dafny.IntOfInt64(419)).Cmp(_104_v2) <= 0) && ((_104_v2).Cmp(_dafny.IntOfInt64(579)) < 0) {
						_coll12.Add(Companion_Default___.SafeDivisionInt(_104_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfqi")).Cardinality())), _dafny.IntOfInt64(971))
					}
				}
				return _coll12.ToMap()
			}()).Keys().Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _105_v1 _dafny.Int
				_105_v1 = interface{}(_compr_11).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll13 = _dafny.NewMapBuilder()
					_ = _coll13
					for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(419), _dafny.IntOfInt64(579))); ; {
						_compr_13, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _104_v2 _dafny.Int
						_104_v2 = interface{}(_compr_13).(_dafny.Int)
						if ((_dafny.IntOfInt64(419)).Cmp(_104_v2) <= 0) && ((_104_v2).Cmp(_dafny.IntOfInt64(579)) < 0) {
							_coll13.Add(Companion_Default___.SafeDivisionInt(_104_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfqi")).Cardinality())), _dafny.IntOfInt64(971))
						}
					}
					return _coll13.ToMap()
				}()).Contains(_105_v1) {
					_coll11.Add(Companion_Default___.SafeModuloInt(_105_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(689))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg21 _dafny.Int) interface{} {
							return coer21(arg21)
						}
					}(func(_106_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					}))).Cardinality())), true)
				}
			}
			return _coll11.ToMap()
		}()).Keys().Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _107_v0 _dafny.Int
			_107_v0 = interface{}(_compr_10).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate((func() _dafny.Map {
					var _coll15 = _dafny.NewMapBuilder()
					_ = _coll15
					for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(419), _dafny.IntOfInt64(579))); ; {
						_compr_15, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _104_v2 _dafny.Int
						_104_v2 = interface{}(_compr_15).(_dafny.Int)
						if ((_dafny.IntOfInt64(419)).Cmp(_104_v2) <= 0) && ((_104_v2).Cmp(_dafny.IntOfInt64(579)) < 0) {
							_coll15.Add(Companion_Default___.SafeDivisionInt(_104_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfqi")).Cardinality())), _dafny.IntOfInt64(971))
						}
					}
					return _coll15.ToMap()
				}()).Keys().Elements()); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _105_v1 _dafny.Int
					_105_v1 = interface{}(_compr_14).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll16 = _dafny.NewMapBuilder()
						_ = _coll16
						for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(419), _dafny.IntOfInt64(579))); ; {
							_compr_16, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _104_v2 _dafny.Int
							_104_v2 = interface{}(_compr_16).(_dafny.Int)
							if ((_dafny.IntOfInt64(419)).Cmp(_104_v2) <= 0) && ((_104_v2).Cmp(_dafny.IntOfInt64(579)) < 0) {
								_coll16.Add(Companion_Default___.SafeDivisionInt(_104_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfqi")).Cardinality())), _dafny.IntOfInt64(971))
							}
						}
						return _coll16.ToMap()
					}()).Contains(_105_v1) {
						_coll14.Add(Companion_Default___.SafeModuloInt(_105_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(689))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg22 _dafny.Int) interface{} {
								return coer22(arg22)
							}
						}(func(_106_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('f')
						}))).Cardinality())), true)
					}
				}
				return _coll14.ToMap()
			}()).Contains(_107_v0) {
				_coll10.Add(Companion_Default___.SafeModuloInt(_107_v0, (_dafny.MultiSetOf(!(false))).Cardinality()), _dafny.IntOfInt64(-334))
			}
		}
		return _coll10.ToMap()
	}()).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())) < 0) {
		return Companion_D8_.Create_DC23_(_dafny.SeqOf(true))
	} else {
		return Companion_D8_.Create_DC23_(_dafny.SeqOf(!(!(true))))
	}
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if !(!(false)) {
			return !(true)
		}
		return false
	})(), (func() D4 {
		if false {
			return Companion_D4_.Create_DC11_(_dafny.IntOfInt64(859), func() _dafny.Set {
				var _coll17 = _dafny.NewBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(845), _dafny.IntOfInt64(623))); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _108_v0 _dafny.Int
					_108_v0 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(845)).Cmp(_108_v0) <= 0) && ((_108_v0).Cmp(_dafny.IntOfInt64(623)) < 0) {
						_coll17.Add(Companion_Default___.SafeDivisionInt(_108_v0, _dafny.IntOfInt64(-722)))
					}
				}
				return _coll17.ToSet()
			}(), false)
		}
		return Companion_D4_.Create_DC11_((_dafny.SetOf(!(false), true, !(!(!(false))))).Cardinality(), func() _dafny.Set {
			var _coll18 = _dafny.NewBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-473), _dafny.IntOfInt64(423))); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _109_v1 _dafny.Int
				_109_v1 = interface{}(_compr_18).(_dafny.Int)
				if ((_dafny.IntOfInt64(-473)).Cmp(_109_v1) <= 0) && ((_109_v1).Cmp(_dafny.IntOfInt64(423)) < 0) {
					_coll18.Add((_109_v1).Plus(_dafny.IntOfInt64(470)))
				}
			}
			return _coll18.ToSet()
		}(), !(true))
	})())
}
func (_static *CompanionStruct_Default___) Fm54(globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll19 = _dafny.NewBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_110_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})))).Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _111_v0 _dafny.CodePoint
			_111_v0 = interface{}(_compr_19).(_dafny.CodePoint)
			if (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}(func(_110_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})))).Contains(_111_v0) {
				_coll19.Add(_111_v0)
			}
		}
		return _coll19.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC20_((func() D6 {
		if true {
			return Companion_D6_.Create_DC19_(!(false), _dafny.CodePoint('i'))
		}
		return Companion_D6_.Create_DC20_(Companion_D6_.Create_DC20_(Companion_D6_.Create_DC19_(false, _dafny.CodePoint('q'))))
	})())
}
func (_static *CompanionStruct_Default___) Fm56(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-87))).Uint32(), func(coer25 func(_dafny.Int) D14) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_112_i0 _dafny.Int) D14 {
		return Companion_D14_.Create_DC36_(false, _dafny.UnicodeSeqOfUtf8Bytes("mpbldi"), _dafny.UnicodeSeqOfUtf8Bytes("dccddbi"))
	})), _dafny.SeqOf(Companion_D14_.Create_DC36_(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_113_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.UnicodeSeqOfUtf8Bytes("xdxoru")), Companion_D14_.Create_DC36_(true, _dafny.UnicodeSeqOfUtf8Bytes("igbltr"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg27 _dafny.Int) interface{} {
			return coer27(arg27)
		}
	}(func(_114_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))))), _dafny.SeqOf(Companion_D14_.Create_DC36_(!(!(true)), _dafny.UnicodeSeqOfUtf8Bytes("d"), _dafny.UnicodeSeqOfUtf8Bytes("kwqsarg"))))
}
func (_static *CompanionStruct_Default___) Fm57(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.CodePoint('g'))).Difference(_dafny.MultiSetOf(_dafny.CodePoint('s')))).Difference((_dafny.MultiSetOf(_dafny.CodePoint('k'), _dafny.CodePoint('q'), _dafny.CodePoint('i'), _dafny.CodePoint('m'))).Union(_dafny.MultiSetOf(_dafny.CodePoint('i'), _dafny.CodePoint('r'))))
}
func (_static *CompanionStruct_Default___) Fm58(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Union(_dafny.SetOf((_dafny.MultiSetOf(false)).Cardinality())), func() _dafny.Set {
		var _coll20 = _dafny.NewBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(630), _dafny.IntOfInt64(790))); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _115_v0 _dafny.Int
			_115_v0 = interface{}(_compr_20).(_dafny.Int)
			if ((_dafny.IntOfInt64(630)).Cmp(_115_v0) <= 0) && ((_115_v0).Cmp(_dafny.IntOfInt64(790)) < 0) {
				_coll20.Add((_115_v0).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("mnhgh"))).Cardinality()))
			}
		}
		return _coll20.ToSet()
	}(), _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.SetOf(false, false)).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm59(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.SeqOf(!(true), true, !(true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), (Companion_D17_.Create_DC44_((Companion_D34_.Create_DC86_(true, _dafny.IntOfInt64(906))).Dtor_cf148(), _dafny.SeqOf(true), _dafny.CodePoint('p'), !(true), (_dafny.SetOf(false)).Cardinality())).Dtor_cf85()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), _dafny.SeqOf(!(!(false)), true)))
}
func (_static *CompanionStruct_Default___) Fm60(p0 D28, p1 bool, globalState *GlobalState) D15 {
	var _source5 _dafny.Sequence = _dafny.SeqOf(Companion_D4_.Create_DC11_((_dafny.MultiSetOf(true)).Cardinality(), _dafny.SetOf(_dafny.IntOfInt64(943), _dafny.IntOfInt64(-813), _dafny.IntOfInt64(-314)), !(false)))
	_ = _source5
	var _116___mcc_h0 _dafny.Sequence = _source5
	_ = _116___mcc_h0
	var _117_cf60 _dafny.Sequence = _116___mcc_h0
	_ = _117_cf60
	return Companion_D15_.Create_DC39_(true, true)
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	var _source6 D34 = Companion_D34_.Create_DC85_(_dafny.CodePoint('p'))
	_ = _source6
	if _source6.Is_DC85() {
		var _118___mcc_h0 _dafny.CodePoint = _source6.Get_().(D34_DC85).Cf146
		_ = _118___mcc_h0
		var _119_cf146 _dafny.CodePoint = _118___mcc_h0
		_ = _119_cf146
		return _dafny.MultiSetOf(_dafny.IntOfInt64(857))
	} else if _source6.Is_DC86() {
		var _120___mcc_h1 bool = _source6.Get_().(D34_DC86).Cf147
		_ = _120___mcc_h1
		var _121___mcc_h2 _dafny.Int = _source6.Get_().(D34_DC86).Cf148
		_ = _121___mcc_h2
		var _122_cf148 _dafny.Int = _121___mcc_h2
		_ = _122_cf148
		var _123_cf147 bool = _120___mcc_h1
		_ = _123_cf147
		if _123_cf147 {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(_122_cf148, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_124_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("ey")
			}))).Cardinality())))
		} else {
			return _dafny.MultiSetOf(_122_cf148, _122_cf148)
		}
	} else {
		var _125___mcc_h3 _dafny.Sequence = _source6.Get_().(D34_DC84).Cf145
		_ = _125___mcc_h3
		var _126_cf145 _dafny.Sequence = _125___mcc_h3
		_ = _126_cf145
		return _dafny.MultiSetOf(_dafny.IntOfInt64(-923), (_dafny.SetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qgxceec")).Cardinality()), _dafny.IntOfInt64(738))
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _127_v0 _dafny.Array
	_ = _127_v0
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = func(_128_i0 _dafny.Int) bool {
			return false
		}
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
	_127_v0 = _nw0
	var _129_v1 bool
	_ = _129_v1
	_129_v1 = true
	var _130_v2 _dafny.Map
	_ = _130_v2
	_130_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(4), _129_v1)
	var _131_globalState *GlobalState
	_ = _131_globalState
	var _nw1 *GlobalState = New_GlobalState_()
	_ = _nw1
	_nw1.Ctor__(false, false, _127_v0, true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v2, !(_129_v1)), _dafny.IntOfInt64(-523), false, _dafny.IntOfInt64(-26), false)
	_131_globalState = _nw1
	var _132_v3 *C1
	_ = _132_v3
	var _nw2 *C1 = New_C1_()
	_ = _nw2
	_nw2.Ctor__(_129_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}(func(_133_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})))
	_132_v3 = _nw2
	var _134_v4 _dafny.Array
	_ = _134_v4
	var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
	_ = _nw3
	_134_v4 = _nw3
	var _135_v5 D17
	_ = _135_v5
	_135_v5 = Companion_D17_.Create_DC43_(_134_v4)
	var _136_v6 _dafny.MultiSet
	_ = _136_v6
	_136_v6 = _dafny.MultiSetOf(_135_v5)
	var _137_v7 _dafny.Map
	_ = _137_v7
	_137_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_136_v6, _dafny.IntOfUint32((_dafny.SeqOf(_129_v1)).Cardinality()))
	var _138_v8 _dafny.Int
	_ = _138_v8
	_138_v8 = _dafny.IntOfInt64(925)
	if (_132_v3).Fm32((func() _dafny.Int {
		if (_137_v7).Contains((_136_v6).Update(_135_v5, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdlthb")).Cardinality())))) {
			return (_137_v7).Get((_136_v6).Update(_135_v5, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdlthb")).Cardinality())))).(_dafny.Int)
		}
		return _138_v8
	})(), (func() _dafny.CodePoint {
		if _129_v1 {
			return _dafny.CodePoint('j')
		}
		return _dafny.CodePoint('l')
	})(), _131_globalState) {
		var _139_v9 *C0
		_ = _139_v9
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__()
		_139_v9 = _nw4
		var _140_v10 _dafny.CodePoint
		_ = _140_v10
		_140_v10 = _dafny.CodePoint('v')
		_138_v8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("pc"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pc")).Cardinality()))).Uint32(), _140_v10)).Cardinality())
		(_131_globalState).F8 = !((_138_v8).Cmp(Companion_Default___.SafeDivisionInt(_138_v8, _dafny.IntOfInt64(67))) != 0)
		var _141_v11 _dafny.Map
		_ = _141_v11
		_141_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
			if (_132_v3).F12() {
				return (_132_v3).F12()
			}
			return (_132_v3).F12()
		})()), _138_v8)
		var _142_v12 _dafny.Sequence
		_ = _142_v12
		_142_v12 = _dafny.SeqOf(_129_v1, (_132_v3).F12())
		var _143_v13 _dafny.Set
		_ = _143_v13
		_143_v13 = _dafny.SetOf(_138_v8, _dafny.IntOfInt64(521))
		_141_v11 = (_141_v11).Update((_132_v3).F12(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v10, _142_v12)).Merge(Companion_Default___.Fm59((_143_v13).Cardinality(), _138_v8, _131_globalState))).Cardinality())
		var _144_v14 D12
		_ = _144_v14
		_144_v14 = Companion_D12_.Create_DC31_((_132_v3).F12(), (_132_v3).F13(), (_132_v3).F13())
		var _145_v15 _dafny.Map
		_ = _145_v15
		_145_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, (_132_v3).F13())
		var _146_v16 D28
		_ = _146_v16
		_146_v16 = Companion_D28_.Create_DC67_(Companion_Default___.Fm58(_131_globalState))
		var _147_v17 _dafny.Array
		_ = _147_v17
		var _nwElement0_0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13())
		_ = _nwElement0_0
		var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(20))
		_ = _nw5
		(_nw5).ArraySet1(_nwElement0_0, 0)
		(_nw5).ArraySet1((_132_v3).F13(), 1)
		(_nw5).ArraySet1((_132_v3).F13(), 2)
		(_nw5).ArraySet1((_132_v3).F13(), 3)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13()), 4)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13()), 5)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_144_v14).Dtor_cf56()), 6)
		(_nw5).ArraySet1((_132_v3).F13(), 7)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13()), 8)
		(_nw5).ArraySet1(Companion_Default___.Fm22(_138_v8, (_132_v3).F12(), false, _138_v8, _131_globalState), 9)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _140_v10), _dafny.UnicodeSeqOfUtf8Bytes("bmidbbbbs")), 10)
		(_nw5).ArraySet1((_132_v3).F13(), 11)
		(_nw5).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_145_v15).Contains((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v10, (Companion_Default___.Fm60(_146_v16, (_132_v3).F12(), _131_globalState)).Dtor_cf75())).Cardinality())) {
				return (_145_v15).Get((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v10, (Companion_Default___.Fm60(_146_v16, (_132_v3).F12(), _131_globalState)).Dtor_cf75())).Cardinality())).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("etfudx")
		})(), (_132_v3).F13()), 12)
		(_nw5).ArraySet1((_132_v3).F13(), 13)
		(_nw5).ArraySet1((_132_v3).F13(), 14)
		(_nw5).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_148_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_149_i2 _dafny.Int) _dafny.CodePoint {
				return _148_v10
			}
		})(_140_v10))), 15)
		(_nw5).ArraySet1(Companion_Default___.Fm13(_138_v8, _131_globalState), 16)
		(_nw5).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ieoynisy"), 17)
		(_nw5).ArraySet1((_132_v3).F13(), 18)
		(_nw5).ArraySet1(Companion_Default___.Fm45((_132_v3).F12(), _138_v8, (_132_v3).F12(), _131_globalState), 19)
		_147_v17 = _nw5
		_147_v17 = _147_v17
	} else {
		var _150_v18 *C0
		_ = _150_v18
		var _nw6 *C0 = New_C0_()
		_ = _nw6
		_nw6.Ctor__()
		_150_v18 = _nw6
		var _151_v19 _dafny.Map
		_ = _151_v19
		_151_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _138_v8)
		var _152_v20 _dafny.Map
		_ = _152_v20
		_152_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_132_v3).F12(), _129_v1)
		_151_v19 = (_151_v19).Update(_129_v1, Companion_Default___.SafeDivisionInt(_138_v8, (_dafny.Zero).Minus((_152_v20).Cardinality())))
		if (_132_v3).F12() {
			_138_v8 = (_138_v8).Minus(Companion_Default___.SafeModuloInt(_138_v8, Companion_Default___.Fm33((_132_v3).F12(), (_132_v3).F12(), _138_v8, (_132_v3).F12(), _131_globalState)))
			var _153_v21 _dafny.MultiSet
			_ = _153_v21
			_153_v21 = _dafny.MultiSetOf((_132_v3).F12())
			var _154_v22 _dafny.Map
			_ = _154_v22
			_154_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if _129_v1 {
					return (_132_v3).F12()
				}
				return _129_v1
			})(), _153_v21)
			_154_v22 = (_154_v22).Update((_138_v8).Cmp(_138_v8) < 0, _153_v21)
			var _155_v23 _dafny.Sequence
			_ = _155_v23
			_155_v23 = _dafny.UnicodeSeqOfUtf8Bytes("ynqtib")
			var _156_v24 _dafny.Set
			_ = _156_v24
			_156_v24 = _dafny.SetOf(!((_132_v3).F12()))
			var _157_v25 _dafny.Map
			_ = _157_v25
			_157_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm44((_132_v3).F12(), _131_globalState), true)
			var _rhs0 _dafny.Int = ((_156_v24).Difference(Companion_Default___.Fm34(_138_v8, (_132_v3).F12(), _129_v1, (func() bool {
				if (_157_v25).Contains(_130_v2) {
					return (_157_v25).Get(_130_v2).(bool)
				}
				return (_132_v3).F12()
			})(), _131_globalState))).Cardinality()
			_ = _rhs0
			var _rhs1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_132_v3).F12() {
					return (_132_v3).F13()
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("i")
			})(), _dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13()))
			_ = _rhs1
			_138_v8 = _rhs0
			_155_v23 = _rhs1
			var _158_v26 *C5
			_ = _158_v26
			var _nw7 *C5 = New_C5_()
			_ = _nw7
			_nw7.Ctor__(_152_v20, (_132_v3).F12())
			_158_v26 = _nw7
			_158_v26 = _158_v26
			var _159_v27 _dafny.CodePoint
			_ = _159_v27
			_159_v27 = _dafny.CodePoint('f')
			_155_v23 = _dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _159_v27)
		} else {
			_138_v8 = _dafny.IntOfInt64(-455)
			var _160_v28 T0
			_ = _160_v28
			var _nw8 *C3 = New_C3_()
			_ = _nw8
			_nw8.Ctor__()
			_160_v28 = _nw8
			var _161_v29 _dafny.CodePoint
			_ = _161_v29
			_161_v29 = _dafny.CodePoint('o')
			var _rhs2 bool = !(((_138_v8).Plus(_138_v8)).Cmp(Companion_Default___.SafeModuloInt(_138_v8, _138_v8)) != 0)
			_ = _rhs2
			var _rhs3 bool = (func() bool {
				if !((_132_v3).Fm32(_138_v8, _161_v29, _131_globalState)) {
					return _129_v1
				}
				return _129_v1
			})()
			_ = _rhs3
			var _rhs4 T0 = _160_v28
			_ = _rhs4
			var _rhs5 _dafny.Int = Companion_Default___.SafeDivisionInt(_138_v8, _138_v8)
			_ = _rhs5
			var _lhs0 *GlobalState = _131_globalState
			_ = _lhs0
			_lhs0.F8 = _rhs2
			_129_v1 = _rhs3
			_160_v28 = _rhs4
			_138_v8 = _rhs5
			_138_v8 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_138_v8, _138_v8))
			_138_v8 = Companion_Default___.Fm33((_132_v3).F12(), (_132_v3).F12(), Companion_Default___.Fm33((_132_v3).F12(), _129_v1, _dafny.IntOfInt64(590), (_132_v3).F12(), _131_globalState), _129_v1, _131_globalState)
			_138_v8 = _138_v8
		}
		var _162_v30 _dafny.Array
		_ = _162_v30
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
		_ = _nw9
		_162_v30 = _nw9
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_162_v30), 0))
		_ = _index0
		(_162_v30).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mlch"), (_132_v3).F13())).Cardinality()), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_162_v30), 0))
		_ = _index1
		(_162_v30).ArraySet1((_138_v8).Minus(Companion_Default___.Fm33(_129_v1, (_132_v3).F12(), (_dafny.Zero).Minus(_138_v8), !(true), _131_globalState)), (_index1).Int())
		var _163_v32 _dafny.Array
		_ = _163_v32
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_1
		var _nw10 _dafny.Array
		_ = _nw10
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw10 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Map = (func(_164_v3 *C1) func(_dafny.Int) _dafny.Map {
				return func(_165_i3 _dafny.Int) _dafny.Map {
					return func() _dafny.Map {
						var _coll21 = _dafny.NewMapBuilder()
						_ = _coll21
						for _iter21 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}(func(_166_i4 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(610))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg32 _dafny.Int) interface{} {
									return coer32(arg32)
								}
							}(func(_167_i5 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('s')
							}))
						}))).Elements()); ; {
							_compr_21, _ok21 := _iter21()
							if !_ok21 {
								break
							}
							var _168_v31 _dafny.Sequence
							_168_v31 = interface{}(_compr_21).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg33 _dafny.Int) interface{} {
									return coer33(arg33)
								}
							}(func(_166_i4 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(610))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg34 _dafny.Int) interface{} {
										return coer34(arg34)
									}
								}(func(_167_i5 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('s')
								}))
							})), _168_v31) {
								_coll21.Add(_168_v31, _dafny.IntOfUint32(((_164_v3).F13()).Cardinality()))
							}
						}
						return _coll21.ToMap()
					}()
				}
			})(_132_v3)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw10 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw10).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw10).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_163_v32 = _nw10
		var _169_v33 _dafny.Map
		_ = _169_v33
		_169_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_132_v3).F13(), (_162_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_162_v30), 0))).Int()).(_dafny.Int))
		var _170_v34 _dafny.CodePoint
		_ = _170_v34
		_170_v34 = _dafny.CodePoint('p')
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_163_v32), 0))
		_ = _index2
		(_163_v32).ArraySet1(((_169_v33).Update(_dafny.SeqOf(_170_v34, _170_v34), (_dafny.MultiSetOf(_138_v8)).Cardinality())).Update((_132_v3).F13(), _138_v8), (_index2).Int())
		var _171_v35 _dafny.Map
		_ = _171_v35
		_171_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v30, _138_v8)
		var _172_v36 _dafny.MultiSet
		_ = _172_v36
		_172_v36 = _dafny.MultiSetOf((_132_v3).F12())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_163_v32), 0))
		_ = _index3
		var _rhs6 _dafny.Int = (func() _dafny.Int {
			if (_171_v35).Equals(_171_v35) {
				return _138_v8
			}
			return (_dafny.MultiSetOf((_172_v36).Cardinality())).Cardinality()
		})()
		_ = _rhs6
		var _rhs7 _dafny.Map = ((_169_v33).Update((_132_v3).F13(), (_162_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_162_v30), 0))).Int()).(_dafny.Int))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_132_v3).F13(), (_162_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_162_v30), 0))).Int()).(_dafny.Int))).Merge((_169_v33).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}((func(_173_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_174_i6 _dafny.Int) _dafny.CodePoint {
				return _173_v34
			}
		})(_170_v34))), (_162_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_162_v30), 0))).Int()).(_dafny.Int))))
		_ = _rhs7
		var _rhs8 _dafny.CodePoint = _170_v34
		_ = _rhs8
		var _lhs1 _dafny.Array = _163_v32
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_163_v32), 0))
		_ = _lhs2
		_138_v8 = _rhs6
		(_lhs1).ArraySet1(_rhs7, (_lhs2).Int())
		_170_v34 = _rhs8
	}
	var _175_v37 _dafny.CodePoint
	_ = _175_v37
	_175_v37 = _dafny.CodePoint('q')
	_175_v37 = _175_v37
	var _176_v38 _dafny.Map
	_ = _176_v38
	_176_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_132_v3).F12(), _129_v1)
	var _177_v39 T1
	_ = _177_v39
	var _nw11 *C6 = New_C6_()
	_ = _nw11
	_nw11.Ctor__(_176_v38, true)
	_177_v39 = _nw11
	var _178_v40 D22
	_ = _178_v40
	_178_v40 = Companion_D22_.Create_DC56_(_177_v39)
	var _179_v41 _dafny.Array
	_ = _179_v41
	var _nwElement0_1 T1 = _177_v39
	_ = _nwElement0_1
	var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
	_ = _nw12
	(_nw12).ArraySet1(_nwElement0_1, 0)
	(_nw12).ArraySet1(_177_v39, 1)
	(_nw12).ArraySet1(_177_v39, 2)
	(_nw12).ArraySet1(_177_v39, 3)
	(_nw12).ArraySet1(_177_v39, 4)
	(_nw12).ArraySet1(_177_v39, 5)
	(_nw12).ArraySet1(_177_v39, 6)
	(_nw12).ArraySet1(_177_v39, 7)
	(_nw12).ArraySet1(_177_v39, 8)
	(_nw12).ArraySet1(_177_v39, 9)
	(_nw12).ArraySet1(_177_v39, 10)
	(_nw12).ArraySet1(_177_v39, 11)
	(_nw12).ArraySet1(_177_v39, 12)
	(_nw12).ArraySet1(_177_v39, 13)
	(_nw12).ArraySet1(_177_v39, 14)
	(_nw12).ArraySet1(_177_v39, 15)
	(_nw12).ArraySet1(_177_v39, 16)
	(_nw12).ArraySet1(_177_v39, 17)
	(_nw12).ArraySet1(_177_v39, 18)
	(_nw12).ArraySet1(_177_v39, 19)
	(_nw12).ArraySet1((_178_v40).Dtor_cf104(), 20)
	(_nw12).ArraySet1(_177_v39, 21)
	(_nw12).ArraySet1(_177_v39, 22)
	(_nw12).ArraySet1(_177_v39, 23)
	(_nw12).ArraySet1(_177_v39, 24)
	(_nw12).ArraySet1((func() T1 {
		if (_177_v39).F10() {
			return _177_v39
		}
		return _177_v39
	})(), 25)
	_179_v41 = _nw12
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_179_v41), 0))
	_ = _index4
	(_179_v41).ArraySet1(_177_v39, (_index4).Int())
	var _180_v42 _dafny.Sequence
	_ = _180_v42
	_180_v42 = _dafny.SeqOf(_177_v39, _177_v39, _177_v39)
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_179_v41), 0))
	_ = _index5
	(_179_v41).ArraySet1((_180_v42).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_138_v8, _138_v8), _dafny.IntOfUint32((_180_v42).Cardinality()))).Uint32()).(T1), (_index5).Int())
	_138_v8 = _dafny.IntOfInt64(-556)
	var _181_v43 *C4
	_ = _181_v43
	var _nw13 *C4 = New_C4_()
	_ = _nw13
	_nw13.Ctor__((_177_v39).F9(), (_177_v39).F10())
	_181_v43 = _nw13
	var _182_v44 D30
	_ = _182_v44
	_182_v44 = Companion_D30_.Create_DC73_(_181_v43)
	_181_v43 = (func() *C4 {
		if (_132_v3).F12() {
			return _181_v43
		}
		return (_182_v44).Dtor_cf127()
	})()
	var _183_v45 _dafny.Sequence
	_ = _183_v45
	_183_v45 = _dafny.SeqOf(_138_v8)
	var _184_v46 D1
	_ = _184_v46
	_184_v46 = Companion_D1_.Create_DC2_(_183_v45)
	var _pat_let_tv0 = _132_v3
	_ = _pat_let_tv0
	var _pat_let_tv1 = _138_v8
	_ = _pat_let_tv1
	var _pat_let_tv2 = _138_v8
	_ = _pat_let_tv2
	_138_v8 = func(_source7 D1) _dafny.Int {
		if _source7.Is_DC3() {
			var _185___mcc_h0 _dafny.Sequence = _source7.Get_().(D1_DC3).Cf5
			_ = _185___mcc_h0
			var _186_cf5 _dafny.Sequence = _185___mcc_h0
			_ = _186_cf5
			return _dafny.IntOfInt64(138)
		} else if _source7.Is_DC4() {
			var _187___mcc_h1 bool = _source7.Get_().(D1_DC4).Cf6
			_ = _187___mcc_h1
			var _188___mcc_h2 bool = _source7.Get_().(D1_DC4).Cf7
			_ = _188___mcc_h2
			var _189___mcc_h3 _dafny.Int = _source7.Get_().(D1_DC4).Cf8
			_ = _189___mcc_h3
			var _190_cf8 _dafny.Int = _189___mcc_h3
			_ = _190_cf8
			var _191_cf7 bool = _188___mcc_h2
			_ = _191_cf7
			var _192_cf6 bool = _187___mcc_h1
			_ = _192_cf6
			return _dafny.IntOfUint32((_dafny.SeqOf(_191_cf7, !(_191_cf7), (_pat_let_tv0).F12())).Cardinality())
		} else if _source7.Is_DC2() {
			var _193___mcc_h4 _dafny.Sequence = _source7.Get_().(D1_DC2).Cf4
			_ = _193___mcc_h4
			var _194_cf4 _dafny.Sequence = _193___mcc_h4
			_ = _194_cf4
			return _dafny.IntOfInt64(558)
		} else {
			var _195___mcc_h5 D1 = _source7.Get_().(D1_DC5).Cf9
			_ = _195___mcc_h5
			var _196_cf9 D1 = _195___mcc_h5
			_ = _196_cf9
			return Companion_Default___.SafeModuloInt(_pat_let_tv1, _pat_let_tv2)
		}
	}(_184_v46)
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
	_ = _index6
	(_127_v0).ArraySet1((_132_v3).Fm32(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), _175_v37, _131_globalState), (_index6).Int())
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
	_ = _index7
	(_127_v0).ArraySet1((_132_v3).F12(), (_index7).Int())
	if !(_129_v1) || ((_177_v39).F10()) {
		if (_177_v39).F10() {
			var _197_v47 _dafny.Array
			_ = _197_v47
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw14
			_197_v47 = _nw14
			var _198_v48 _dafny.Map
			_ = _198_v48
			_198_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_197_v47, _138_v8)
			var _199_v49 _dafny.Map
			_ = _199_v49
			_199_v49 = _198_v48
			var _200_v50 _dafny.Set
			_ = _200_v50
			_200_v50 = _dafny.SetOf(_199_v49)
			var _201_v51 _dafny.Sequence
			_ = _201_v51
			_201_v51 = _dafny.SeqOf(_200_v50)
			var _202_v52 _dafny.Array
			_ = _202_v52
			var _nwElement0_2 _dafny.Sequence = _dafny.SeqOf(_200_v50)
			_ = _nwElement0_2
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_2, 0)
			(_nw15).ArraySet1(_201_v51, 1)
			_202_v52 = _nw15
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_202_v52), 0))
			_ = _index8
			(_202_v52).ArraySet1(_201_v51, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_197_v47), 0))
			_ = _index9
			(_197_v47).ArraySet1(_138_v8, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_202_v52), 0))
			_ = _index10
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_197_v47), 0))
			_ = _index11
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index12
			var _rhs9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_201_v51, _201_v51)
			_ = _rhs9
			var _rhs10 _dafny.Int = (func() _dafny.Int {
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _175_v37)).Contains(_138_v8) {
					return _138_v8
				}
				return _138_v8
			})()
			_ = _rhs10
			var _rhs11 _dafny.Int = Companion_Default___.SafeModuloInt((_183_v45).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _138_v8)).Cardinality(), _dafny.IntOfUint32((_183_v45).Cardinality()))).Uint32()).(_dafny.Int), (Companion_D8_.Create_DC24_(true, _dafny.SetOf(_197_v47, _197_v47, _197_v47), _138_v8)).Dtor_cf44())
			_ = _rhs11
			var _rhs12 bool = _129_v1
			_ = _rhs12
			var _lhs3 _dafny.Array = _202_v52
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_202_v52), 0))
			_ = _lhs4
			var _lhs5 _dafny.Array = _197_v47
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_197_v47), 0))
			_ = _lhs6
			var _lhs7 _dafny.Array = _127_v0
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _lhs8
			(_lhs3).ArraySet1(_rhs9, (_lhs4).Int())
			(_lhs5).ArraySet1(_rhs10, (_lhs6).Int())
			_138_v8 = _rhs11
			(_lhs7).ArraySet1(_rhs12, (_lhs8).Int())
			var _203_v53 _dafny.Array
			_ = _203_v53
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_2
			var _nw16 _dafny.Array
			_ = _nw16
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw16 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Sequence = (func(_204_v3 *C1) func(_dafny.Int) _dafny.Sequence {
					return func(_205_i7 _dafny.Int) _dafny.Sequence {
						return (_204_v3).F13()
					}
				})(_132_v3)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw16 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw16).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw16).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_203_v53 = _nw16
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_203_v53), 0))
			_ = _index13
			(_203_v53).ArraySet1((_132_v3).F13(), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_203_v53), 0))
			_ = _index14
			(_203_v53).ArraySet1((_132_v3).F13(), (_index14).Int())
			var _206_v54 D30
			_ = _206_v54
			_206_v54 = Companion_D30_.Create_DC74_((_132_v3).F13())
			var _207_v55 D30
			_ = _207_v55
			_207_v55 = Companion_D30_.Create_DC76_(Companion_D30_.Create_DC76_(_206_v54))
			var _pat_let_tv3 = _206_v54
			_ = _pat_let_tv3
			var _208_v56 _dafny.Array
			_ = _208_v56
			var _nwElement0_3 D30 = _207_v55
			_ = _nwElement0_3
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(15))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_3, 0)
			(_nw17).ArraySet1(_207_v55, 1)
			(_nw17).ArraySet1(_207_v55, 2)
			(_nw17).ArraySet1(_207_v55, 3)
			(_nw17).ArraySet1(_207_v55, 4)
			(_nw17).ArraySet1(_207_v55, 5)
			(_nw17).ArraySet1(_207_v55, 6)
			(_nw17).ArraySet1(_207_v55, 7)
			(_nw17).ArraySet1(_207_v55, 8)
			(_nw17).ArraySet1(_207_v55, 9)
			(_nw17).ArraySet1(_207_v55, 10)
			(_nw17).ArraySet1(func(_pat_let0_0 D30) D30 {
				return func(_209_dt__update__tmp_h0 D30) D30 {
					return func(_pat_let1_0 D30) D30 {
						return func(_210_dt__update_hcf130_h0 D30) D30 {
							return Companion_D30_.Create_DC76_(_210_dt__update_hcf130_h0)
						}(_pat_let1_0)
					}(_pat_let_tv3)
				}(_pat_let0_0)
			}(_207_v55), 11)
			(_nw17).ArraySet1(_207_v55, 12)
			(_nw17).ArraySet1(_207_v55, 13)
			(_nw17).ArraySet1(_207_v55, 14)
			_208_v56 = _nw17
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_208_v56), 0))
			_ = _index15
			(_208_v56).ArraySet1(_207_v55, (_index15).Int())
			var _211_v57 _dafny.Sequence
			_ = _211_v57
			_211_v57 = _dafny.SeqOf((_177_v39).F10())
			var _212_v58 _dafny.MultiSet
			_ = _212_v58
			_212_v58 = _dafny.MultiSetOf(!((_132_v3).F12()), !((_177_v39).Fm0(_211_v57, (_177_v39).F10(), _138_v8, false, _131_globalState)), (_132_v3).F12())
			var _213_v59 *C2
			_ = _213_v59
			var _nw18 *C2 = New_C2_()
			_ = _nw18
			_nw18.Ctor__((_132_v3).F12(), _127_v0, (_177_v39).F9(), (_177_v39).F10())
			_213_v59 = _nw18
			var _214_v60 _dafny.Map
			_ = _214_v60
			_214_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(955), _213_v59)
			var _215_v61 _dafny.MultiSet
			_ = _215_v61
			_215_v61 = _dafny.MultiSetOf((_214_v60).Cardinality(), (Companion_Default___.Fm15(true, (_177_v39).F10(), _138_v8, _175_v37, _131_globalState)).Cardinality(), (_197_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_197_v47), 0))).Int()).(_dafny.Int))
			var _216_v62 _dafny.Sequence
			_ = _216_v62
			_216_v62 = _dafny.SeqOf(_132_v3)
			var _217_v63 _dafny.Map
			_ = _217_v63
			_217_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_215_v61).Intersection(Companion_Default___.Fm46((_177_v39).F10(), (_132_v3).F12(), (_132_v3).F12(), _131_globalState)), _216_v62)
			var _218_v64 _dafny.MultiSet
			_ = _218_v64
			_218_v64 = _215_v61
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_208_v56), 0))
			_ = _index16
			var _rhs13 bool = !(_129_v1) || ((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool))
			_ = _rhs13
			var _rhs14 D30 = _207_v55
			_ = _rhs14
			var _rhs15 bool = (_132_v3).F12()
			_ = _rhs15
			var _rhs16 _dafny.MultiSet = _212_v58
			_ = _rhs16
			var _rhs17 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_218_v64), _dafny.SeqOf(_132_v3))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v61, _216_v62))
			_ = _rhs17
			var _lhs9 *GlobalState = _131_globalState
			_ = _lhs9
			var _lhs10 _dafny.Array = _208_v56
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_208_v56), 0))
			_ = _lhs11
			_lhs9.F8 = _rhs13
			(_lhs10).ArraySet1(_rhs14, (_lhs11).Int())
			_129_v1 = _rhs15
			_212_v58 = _rhs16
			_217_v63 = _rhs17
			var _219_v65 _dafny.Map
			_ = _219_v65
			_219_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(865), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_203_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_203_v53), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32(((_203_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_203_v53), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _175_v37)).Cardinality()))
			var _220_v66 _dafny.Map
			_ = _220_v66
			_220_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), (func() _dafny.Int {
				if (_219_v65).Contains(_138_v8) {
					return (_219_v65).Get(_138_v8).(_dafny.Int)
				}
				return _138_v8
			})())
			var _221_v67 _dafny.Sequence
			_ = _221_v67
			_221_v67 = _dafny.SeqOf((_220_v66).Merge(_219_v65))
			_221_v67 = _221_v67
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_197_v47), 0))
			_ = _index17
			(_197_v47).ArraySet1((_197_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_197_v47), 0))).Int()).(_dafny.Int), (_index17).Int())
		} else {
			var _222_v68 _dafny.Map
			_ = _222_v68
			_222_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v3, _127_v0)
			var _223_v69 _dafny.Array
			_ = _223_v69
			var _nwElement0_4 _dafny.Array = _127_v0
			_ = _nwElement0_4
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_4, 0)
			(_nw19).ArraySet1(_127_v0, 1)
			(_nw19).ArraySet1(_127_v0, 2)
			(_nw19).ArraySet1(_127_v0, 3)
			(_nw19).ArraySet1(_127_v0, 4)
			(_nw19).ArraySet1(_127_v0, 5)
			(_nw19).ArraySet1(_127_v0, 6)
			(_nw19).ArraySet1(_127_v0, 7)
			(_nw19).ArraySet1(_127_v0, 8)
			(_nw19).ArraySet1(_127_v0, 9)
			(_nw19).ArraySet1(_127_v0, 10)
			(_nw19).ArraySet1(_127_v0, 11)
			(_nw19).ArraySet1(_127_v0, 12)
			(_nw19).ArraySet1(_127_v0, 13)
			(_nw19).ArraySet1(_127_v0, 14)
			(_nw19).ArraySet1(_127_v0, 15)
			(_nw19).ArraySet1((func() _dafny.Array {
				if (_222_v68).Contains(_132_v3) {
					return (_222_v68).Get(_132_v3).(_dafny.Array)
				}
				return _127_v0
			})(), 16)
			(_nw19).ArraySet1(_127_v0, 17)
			(_nw19).ArraySet1(_127_v0, 18)
			(_nw19).ArraySet1(_127_v0, 19)
			(_nw19).ArraySet1(_127_v0, 20)
			(_nw19).ArraySet1(_127_v0, 21)
			(_nw19).ArraySet1((func() _dafny.Array {
				if (_177_v39).F10() {
					return _127_v0
				}
				return _127_v0
			})(), 22)
			(_nw19).ArraySet1(_127_v0, 23)
			(_nw19).ArraySet1(_127_v0, 24)
			(_nw19).ArraySet1(_127_v0, 25)
			(_nw19).ArraySet1(_127_v0, 26)
			_223_v69 = _nw19
			var _224_v70 D32
			_ = _224_v70
			_224_v70 = Companion_D32_.Create_DC78_(_223_v69)
			_223_v69 = (_224_v70).Dtor_cf132()
			(_131_globalState).F8 = true
			_130_v2 = (_130_v2).Update(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), (_132_v3).F12())
			var _225_v71 _dafny.Array
			_ = _225_v71
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw20
			_225_v71 = _nw20
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_225_v71), 0))
			_ = _index18
			(_225_v71).ArraySet1((_dafny.Zero).Minus(_138_v8), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_225_v71), 0))
			_ = _index19
			(_225_v71).ArraySet1(_dafny.IntOfInt64(299), (_index19).Int())
			var _226_v72 _dafny.MultiSet
			_ = _226_v72
			_226_v72 = _dafny.MultiSetOf((_225_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_225_v71), 0))).Int()).(_dafny.Int), (_225_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_225_v71), 0))).Int()).(_dafny.Int), (_225_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_225_v71), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(178))
			var _227_v73 _dafny.Set
			_ = _227_v73
			_227_v73 = _dafny.SetOf(_138_v8, ((_226_v72).Update(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), Companion_Default___.Abs((_225_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_225_v71), 0))).Int()).(_dafny.Int)))).Cardinality())
			_227_v73 = (_227_v73).Union(_227_v73)
		}
		var _228_v74 _dafny.Sequence
		_ = _228_v74
		var _229_v75 bool
		_ = _229_v75
		var _230_v76 _dafny.Sequence
		_ = _230_v76
		var _231_v77 _dafny.Sequence
		_ = _231_v77
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 _dafny.Sequence
		_ = _out2
		var _out3 _dafny.Sequence
		_ = _out3
		_out0, _out1, _out2, _out3 = (_181_v43).M0(_131_globalState)
		_228_v74 = _out0
		_229_v75 = _out1
		_230_v76 = _out2
		_231_v77 = _out3
		var _232_v78 _dafny.Sequence
		_ = _232_v78
		var _233_v79 bool
		_ = _233_v79
		var _234_v80 _dafny.Sequence
		_ = _234_v80
		var _235_v81 _dafny.Sequence
		_ = _235_v81
		var _out4 _dafny.Sequence
		_ = _out4
		var _out5 bool
		_ = _out5
		var _out6 _dafny.Sequence
		_ = _out6
		var _out7 _dafny.Sequence
		_ = _out7
		_out4, _out5, _out6, _out7 = (_181_v43).M0(_131_globalState)
		_232_v78 = _out4
		_233_v79 = _out5
		_234_v80 = _out6
		_235_v81 = _out7
		(_181_v43).M5(_131_globalState)
		var _236_v82 _dafny.Set
		_ = _236_v82
		_236_v82 = _dafny.SetOf(Companion_Default___.Fm37(_131_globalState))
		_236_v82 = _236_v82
	} else {
		(_131_globalState).F8 = _129_v1
		var _237_v83 _dafny.Array
		_ = _237_v83
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_3
		var _nw21 _dafny.Array
		_ = _nw21
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw21 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.CodePoint = func(_238_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw21 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw21).ArraySet1CodePoint(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw21).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_237_v83 = _nw21
		var _239_v84 _dafny.Array
		_ = _239_v84
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw22
		_239_v84 = _nw22
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))
		_ = _index20
		(_239_v84).ArraySet1(_138_v8, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _index21
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))
		_ = _index22
		var _rhs18 _dafny.Int = (_138_v8).Minus(_138_v8)
		_ = _rhs18
		var _rhs19 bool = (_177_v39).F10()
		_ = _rhs19
		var _rhs20 _dafny.Array = _237_v83
		_ = _rhs20
		var _rhs21 _dafny.Int = _138_v8
		_ = _rhs21
		var _rhs22 _dafny.Int = _138_v8
		_ = _rhs22
		var _lhs12 _dafny.Array = _127_v0
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _lhs13
		var _lhs14 _dafny.Array = _239_v84
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))
		_ = _lhs15
		_138_v8 = _rhs18
		(_lhs12).ArraySet1(_rhs19, (_lhs13).Int())
		_237_v83 = _rhs20
		_138_v8 = _rhs21
		(_lhs14).ArraySet1(_rhs22, (_lhs15).Int())
		var _240_v85 _dafny.Sequence
		_ = _240_v85
		_240_v85 = _dafny.SeqOf((Companion_D3_.Create_DC7_(_239_v84)).Dtor_cf11(), _239_v84)
		var _241_v86 _dafny.Set
		_ = _241_v86
		_241_v86 = _dafny.SetOf((_240_v85).Select((Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_240_v85).Cardinality()))).Uint32()).(_dafny.Array), _239_v84, _239_v84)
		if !((_241_v86).IsDisjointFrom(_241_v86)) {
			_129_v1 = (func() bool {
				if ((_177_v39).F9()).Contains((_177_v39).F10()) {
					return ((_177_v39).F9()).Get((_177_v39).F10()).(bool)
				}
				return ((_dafny.Zero).Minus((_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nlwuecg")).Cardinality())) > 0
			})()
			_138_v8 = (_dafny.Zero).Minus(_138_v8)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))
			_ = _index23
			(_239_v84).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(19))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_242_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_243_i9 _dafny.Int) _dafny.CodePoint {
					return _242_v37
				}
			})(_175_v37)))).Cardinality()), (_index23).Int())
			var _244_v87 _dafny.Map
			_ = _244_v87
			_244_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_129_v1), _138_v8)
			var _245_v88 _dafny.Set
			_ = _245_v88
			_245_v88 = _dafny.SetOf((_132_v3).F13())
			_138_v8 = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt((_244_v87).Cardinality(), (_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int))).Plus((_245_v88).Cardinality()))
			_138_v8 = Companion_Default___.SafeDivisionInt(_138_v8, (_138_v8).Times(_138_v8))
		} else {
			var _246_v89 _dafny.Map
			_ = _246_v89
			_246_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _183_v45)
			var _247_v91 _dafny.MultiSet
			_ = _247_v91
			_247_v91 = _dafny.MultiSetOf((_132_v3).F12(), (_177_v39).F10(), (_132_v3).F12())
			var _248_v92 _dafny.Set
			_ = _248_v92
			_248_v92 = _dafny.SetOf((_247_v91).Cardinality(), _138_v8)
			var _249_v96 _dafny.MultiSet
			_ = _249_v96
			_249_v96 = _dafny.MultiSetOf((_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int))
			var _250_v98 _dafny.Array
			_ = _250_v98
			var _nwElement0_5 _dafny.Map = _246_v89
			_ = _nwElement0_5
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(16))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_5, 0)
			(_nw23).ArraySet1((func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate((_248_v92).Elements()); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _251_v90 _dafny.Int
					_251_v90 = interface{}(_compr_22).(_dafny.Int)
					if (_248_v92).Contains(_251_v90) {
						_coll22.Add((_251_v90).Minus(_138_v8), _183_v45)
					}
				}
				return _coll22.ToMap()
			}()).Merge(_246_v89), 1)
			(_nw23).ArraySet1(_246_v89, 2)
			(_nw23).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(380), _183_v45), 3)
			(_nw23).ArraySet1(_246_v89, 4)
			(_nw23).ArraySet1((_246_v89).Merge((Companion_D33_.Create_DC81_(_246_v89)).Dtor_cf139()), 5)
			(_nw23).ArraySet1((func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter23 := _dafny.Iterate((_183_v45).Elements()); ; {
					_compr_23, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _252_v93 _dafny.Int
					_252_v93 = interface{}(_compr_23).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_183_v45, _252_v93) {
						_coll23.Add(Companion_Default___.SafeModuloInt(_252_v93, (_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int)), _183_v45)
					}
				}
				return _coll23.ToMap()
			}()).Merge(_246_v89), 6)
			(_nw23).ArraySet1(func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-33), _dafny.IntOfInt64(-393))); ; {
					_compr_24, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _253_v94 _dafny.Int
					_253_v94 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(-33)).Cmp(_253_v94) <= 0) && ((_253_v94).Cmp(_dafny.IntOfInt64(-393)) < 0) {
						_coll24.Add((_253_v94).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg37 _dafny.Int) interface{} {
								return coer37(arg37)
							}
						}((func(_254_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_255_i10 _dafny.Int) _dafny.CodePoint {
								return _254_v37
							}
						})(_175_v37)))).Cardinality())), _183_v45)
					}
				}
				return _coll24.ToMap()
			}(), 7)
			(_nw23).ArraySet1(_246_v89, 8)
			(_nw23).ArraySet1(_246_v89, 9)
			(_nw23).ArraySet1(func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter25 := _dafny.Iterate((_249_v96).Elements()); ; {
					_compr_25, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _256_v95 _dafny.Int
					_256_v95 = interface{}(_compr_25).(_dafny.Int)
					if (_249_v96).Contains(_256_v95) {
						_coll25.Add((_256_v95).Times((_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int)), _dafny.SeqOf(_138_v8, (_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll25.ToMap()
			}(), 10)
			(_nw23).ArraySet1(_246_v89, 11)
			(_nw23).ArraySet1(_246_v89, 12)
			(_nw23).ArraySet1(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter26 := _dafny.Iterate((_249_v96).Elements()); ; {
					_compr_26, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _257_v97 _dafny.Int
					_257_v97 = interface{}(_compr_26).(_dafny.Int)
					if (_249_v96).Contains(_257_v97) {
						_coll26.Add(Companion_Default___.SafeModuloInt(_257_v97, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvnf")).Cardinality())), _183_v45)
					}
				}
				return _coll26.ToMap()
			}(), 13)
			(_nw23).ArraySet1(_246_v89, 14)
			(_nw23).ArraySet1(_246_v89, 15)
			_250_v98 = _nw23
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_250_v98), 0))
			_ = _index24
			(_250_v98).ArraySet1(_246_v89, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_250_v98), 0))
			_ = _index25
			(_250_v98).ArraySet1((func() _dafny.Map {
				if false {
					return (_246_v89).Merge(_246_v89)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _183_v45)).Update((_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Update(_183_v45, (Companion_Default___.SafeIndex((_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_183_v45).Cardinality()))).Uint32(), _dafny.IntOfUint32((_180_v42).Cardinality())))
			})(), (_index25).Int())
			var _rhs23 _dafny.Int = (_dafny.Zero).Minus(_138_v8)
			_ = _rhs23
			var _rhs24 bool = !((_132_v3).F12())
			_ = _rhs24
			var _rhs25 _dafny.Array = _239_v84
			_ = _rhs25
			var _lhs16 *GlobalState = _131_globalState
			_ = _lhs16
			_138_v8 = _rhs23
			_lhs16.F8 = _rhs24
			_239_v84 = _rhs25
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index26
			(_127_v0).ArraySet1((_177_v39).F10(), (_index26).Int())
			(_131_globalState).F8 = (_177_v39).F10()
			var _258_v99 _dafny.Set
			_ = _258_v99
			_258_v99 = _dafny.SetOf((_177_v39).F10())
			var _259_v100 *C7
			_ = _259_v100
			var _nw24 *C7 = New_C7_()
			_ = _nw24
			_nw24.Ctor__(_258_v99, (_177_v39).F9(), (_132_v3).Fm32((_239_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_239_v84), 0))).Int()).(_dafny.Int), _175_v37, _131_globalState))
			_259_v100 = _nw24
		}
		var _260_v101 _dafny.Sequence
		_ = _260_v101
		_260_v101 = _dafny.SeqOf((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool))
		var _261_v102 _dafny.Sequence
		_ = _261_v102
		_261_v102 = _dafny.SeqOf((_132_v3).F13())
		_138_v8 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_260_v101).Cardinality()), _138_v8)).Plus(_dafny.IntOfUint32((_261_v102).Cardinality()))
		(_131_globalState).F8 = (_132_v3).F12()
	}
	if (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool) {
		_138_v8 = (_138_v8).Minus(_dafny.IntOfInt64(-444))
		_175_v37 = _175_v37
		var _262_v103 _dafny.Sequence
		_ = _262_v103
		_262_v103 = _dafny.UnicodeSeqOfUtf8Bytes("elcmh")
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _index27
		var _rhs26 bool = !((_dafny.IntOfInt64(645)).Cmp(_138_v8) < 0)
		_ = _rhs26
		var _rhs27 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13())
		_ = _rhs27
		var _lhs17 _dafny.Array = _127_v0
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _lhs18
		(_lhs17).ArraySet1(_rhs26, (_lhs18).Int())
		_262_v103 = _rhs27
		(_181_v43).M5(_131_globalState)
		var _263_v104 _dafny.Sequence
		_ = _263_v104
		_263_v104 = _dafny.SeqOf(_130_v2)
		_263_v104 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg38 _dafny.Int) interface{} {
				return coer38(arg38)
			}
		}((func(_264_v2 _dafny.Map) func(_dafny.Int) _dafny.Map {
			return func(_265_i11 _dafny.Int) _dafny.Map {
				return (_264_v2).Merge(_264_v2)
			}
		})(_130_v2)))
	} else {
		_138_v8 = (_138_v8).Minus(_138_v8)
		var _266_v105 *C0
		_ = _266_v105
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__()
		_266_v105 = _nw25
		var _267_v106 _dafny.Sequence
		_ = _267_v106
		_267_v106 = _dafny.SeqOf(_127_v0)
		var _268_v107 _dafny.Array
		_ = _268_v107
		var _nwElement0_6 _dafny.Array = _127_v0
		_ = _nwElement0_6
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(26))
		_ = _nw26
		(_nw26).ArraySet1(_nwElement0_6, 0)
		(_nw26).ArraySet1(_127_v0, 1)
		(_nw26).ArraySet1(_127_v0, 2)
		(_nw26).ArraySet1(_127_v0, 3)
		(_nw26).ArraySet1(_127_v0, 4)
		(_nw26).ArraySet1(_127_v0, 5)
		(_nw26).ArraySet1(_127_v0, 6)
		(_nw26).ArraySet1(_127_v0, 7)
		(_nw26).ArraySet1(_127_v0, 8)
		(_nw26).ArraySet1(_127_v0, 9)
		(_nw26).ArraySet1(_127_v0, 10)
		(_nw26).ArraySet1(_127_v0, 11)
		(_nw26).ArraySet1(_127_v0, 12)
		(_nw26).ArraySet1(_127_v0, 13)
		(_nw26).ArraySet1(_127_v0, 14)
		(_nw26).ArraySet1((_267_v106).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_138_v8), _dafny.IntOfUint32((_267_v106).Cardinality()))).Uint32()).(_dafny.Array), 15)
		(_nw26).ArraySet1(_127_v0, 16)
		(_nw26).ArraySet1(_127_v0, 17)
		(_nw26).ArraySet1(_127_v0, 18)
		(_nw26).ArraySet1(_127_v0, 19)
		(_nw26).ArraySet1(_127_v0, 20)
		(_nw26).ArraySet1(_127_v0, 21)
		(_nw26).ArraySet1(_127_v0, 22)
		(_nw26).ArraySet1(_127_v0, 23)
		(_nw26).ArraySet1(_127_v0, 24)
		(_nw26).ArraySet1(_127_v0, 25)
		_268_v107 = _nw26
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_268_v107), 0))
		_ = _index28
		(_268_v107).ArraySet1(_127_v0, (_index28).Int())
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_268_v107), 0))
		_ = _index29
		(_268_v107).ArraySet1(_127_v0, (_index29).Int())
		var _269_v108 T0
		_ = _269_v108
		var _nw27 *C1 = New_C1_()
		_ = _nw27
		_nw27.Ctor__((_132_v3).F12(), (_132_v3).F13())
		_269_v108 = _nw27
		_269_v108 = _269_v108
		var _270_v109 _dafny.Array
		_ = _270_v109
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw28
		_270_v109 = _nw28
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_270_v109), 0))
		_ = _index30
		(_270_v109).ArraySet1(_267_v106, (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_270_v109), 0))
		_ = _index31
		(_270_v109).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.ArrayCastTo((_268_v107).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_268_v107), 0))).Int()))), _dafny.Companion_Sequence_.Concatenate(_267_v106, _267_v106)), (_index31).Int())
	}
	(_131_globalState).F8 = !((_177_v39).F10())
	var _271_v110 _dafny.MultiSet
	_ = _271_v110
	_271_v110 = _dafny.MultiSetOf((_132_v3).F12(), (_132_v3).F12())
	var _272_v111 T0
	_ = _272_v111
	var _nw29 *C1 = New_C1_()
	_ = _nw29
	_nw29.Ctor__((_dafny.IntOfInt64(948)).Cmp((_271_v110).Cardinality()) != 0, _dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13()))
	_272_v111 = _nw29
	var _273_v112 *C5
	_ = _273_v112
	var _nw30 *C5 = New_C5_()
	_ = _nw30
	_nw30.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v1, true), (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool))
	_273_v112 = _nw30
	var _274_v113 _dafny.Map
	_ = _274_v113
	_274_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_273_v112, _138_v8)
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
	_ = _index32
	var _rhs28 bool = (_132_v3).Fm32(_138_v8, _175_v37, _131_globalState)
	_ = _rhs28
	var _rhs29 T0 = _272_v111
	_ = _rhs29
	var _rhs30 bool = _129_v1
	_ = _rhs30
	var _rhs31 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_138_v8, (func() _dafny.Int {
		if (_274_v113).Contains(_273_v112) {
			return (_274_v113).Get(_273_v112).(_dafny.Int)
		}
		return _138_v8
	})()), _138_v8)
	_ = _rhs31
	var _lhs19 *GlobalState = _131_globalState
	_ = _lhs19
	var _lhs20 _dafny.Array = _127_v0
	_ = _lhs20
	var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
	_ = _lhs21
	_lhs19.F8 = _rhs28
	_272_v111 = _rhs29
	(_lhs20).ArraySet1(_rhs30, (_lhs21).Int())
	_138_v8 = _rhs31
	if (_177_v39).F10() {
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _index33
		(_127_v0).ArraySet1((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool), (_index33).Int())
		var _275_v114 _dafny.Map
		_ = _275_v114
		_275_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v1, _dafny.MultiSetFromSeq(_183_v45))
		var _276_v115 _dafny.MultiSet
		_ = _276_v115
		_276_v115 = _dafny.MultiSetOf(_138_v8, _dafny.IntOfUint32((_dafny.SeqOf((_177_v39).F10(), (_132_v3).F12())).Cardinality()))
		var _277_v116 _dafny.Sequence
		_ = _277_v116
		_277_v116 = _dafny.SeqOf(_276_v115, _276_v115, (_276_v115).Update(_138_v8, Companion_Default___.Abs(_138_v8)), _276_v115, _276_v115)
		(_131_globalState).F8 = ((_275_v114).Update((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool), (_277_v116).Select((Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_277_v116).Cardinality()))).Uint32()).(_dafny.MultiSet))).Contains((func() bool {
			if (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool) {
				return (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)
			}
			return true
		})())
		var _278_v117 _dafny.Array
		_ = _278_v117
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
		_ = _nw31
		_278_v117 = _nw31
		var _279_v118 D32
		_ = _279_v118
		_279_v118 = Companion_D32_.Create_DC78_(_278_v117)
		var _source8 D32 = _279_v118
		_ = _source8
		if _source8.Is_DC79() {
			var _280___mcc_h6 bool = _source8.Get_().(D32_DC79).Cf133
			_ = _280___mcc_h6
			var _281___mcc_h7 _dafny.Sequence = _source8.Get_().(D32_DC79).Cf134
			_ = _281___mcc_h7
			var _282___mcc_h8 _dafny.Sequence = _source8.Get_().(D32_DC79).Cf135
			_ = _282___mcc_h8
			var _283_cf135 _dafny.Sequence = _282___mcc_h8
			_ = _283_cf135
			var _284_cf134 _dafny.Sequence = _281___mcc_h7
			_ = _284_cf134
			var _285_cf133 bool = _280___mcc_h6
			_ = _285_cf133
			var _286_v119 _dafny.Sequence
			_ = _286_v119
			_286_v119 = _dafny.SeqOf(_129_v1)
			var _287_v120 _dafny.Set
			_ = _287_v120
			_287_v120 = _dafny.SetOf((_132_v3).F12(), (_177_v39).Fm0(_286_v119, (_132_v3).F12(), _138_v8, (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool), _131_globalState))
			_287_v120 = (Companion_Default___.Fm34(_138_v8, false, (_177_v39).F10(), true, _131_globalState)).Union(_287_v120)
			var _288_v121 D0
			_ = _288_v121
			_288_v121 = Companion_D0_.Create_DC1_(_138_v8, (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool), _138_v8)
			var _289_v122 _dafny.Sequence
			_ = _289_v122
			_289_v122 = _dafny.SeqOf(_272_v111, _272_v111)
			var _290_v123 _dafny.Sequence
			_ = _290_v123
			_290_v123 = _dafny.SeqOf(_289_v122)
			var _291_v124 _dafny.Map
			_ = _291_v124
			_291_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _175_v37)
			var _292_v125 _dafny.Array
			_ = _292_v125
			var _nwElement0_7 _dafny.Int = _138_v8
			_ = _nwElement0_7
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(18))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_7, 0)
			(_nw32).ArraySet1(_138_v8, 1)
			(_nw32).ArraySet1(_138_v8, 2)
			(_nw32).ArraySet1(_138_v8, 3)
			(_nw32).ArraySet1(_dafny.IntOfUint32((_290_v123).Cardinality()), 4)
			(_nw32).ArraySet1(_dafny.IntOfUint32((_183_v45).Cardinality()), 5)
			(_nw32).ArraySet1((_287_v120).Cardinality(), 6)
			(_nw32).ArraySet1(_138_v8, 7)
			(_nw32).ArraySet1(_138_v8, 8)
			(_nw32).ArraySet1(_dafny.IntOfUint32((_286_v119).Cardinality()), 9)
			(_nw32).ArraySet1(_138_v8, 10)
			(_nw32).ArraySet1((_dafny.Zero).Minus(_138_v8), 11)
			(_nw32).ArraySet1(_138_v8, 12)
			(_nw32).ArraySet1((_291_v124).Cardinality(), 13)
			(_nw32).ArraySet1((_dafny.Zero).Minus(_138_v8), 14)
			(_nw32).ArraySet1(_138_v8, 15)
			(_nw32).ArraySet1(_138_v8, 16)
			(_nw32).ArraySet1(_138_v8, 17)
			_292_v125 = _nw32
			var _293_v126 _dafny.Set
			_ = _293_v126
			_293_v126 = _dafny.SetOf(_292_v125, _292_v125, _292_v125, _292_v125, _292_v125)
			var _294_v127 D8
			_ = _294_v127
			_294_v127 = Companion_D8_.Create_DC24_((_288_v121).Dtor_cf2(), (_293_v126).Intersection(_293_v126), (_273_v112).Fm1((_177_v39).F9(), _131_globalState))
			_294_v127 = _294_v127
			_175_v37 = _175_v37
			var _295_v128 _dafny.Set
			_ = _295_v128
			_295_v128 = _dafny.SetOf(_dafny.IntOfInt64(921), _138_v8)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index34
			(_127_v0).ArraySet1((_295_v128).IsSubsetOf((_295_v128).Intersection(_295_v128)), (_index34).Int())
		} else if _source8.Is_DC80() {
			var _296___mcc_h9 _dafny.Int = _source8.Get_().(D32_DC80).Cf136
			_ = _296___mcc_h9
			var _297___mcc_h10 _dafny.Int = _source8.Get_().(D32_DC80).Cf137
			_ = _297___mcc_h10
			var _298___mcc_h11 _dafny.Int = _source8.Get_().(D32_DC80).Cf138
			_ = _298___mcc_h11
			var _299_cf138 _dafny.Int = _298___mcc_h11
			_ = _299_cf138
			var _300_cf137 _dafny.Int = _297___mcc_h10
			_ = _300_cf137
			var _301_cf136 _dafny.Int = _296___mcc_h9
			_ = _301_cf136
			(_131_globalState).F8 = _129_v1
			_278_v117 = _278_v117
			_301_cf136 = Companion_Default___.SafeDivisionInt(_299_cf138, _300_cf137)
			var _302_v129 _dafny.Set
			_ = _302_v129
			_302_v129 = _dafny.SetOf((_132_v3).F12(), (_177_v39).F10())
			var _303_v130 _dafny.Map
			_ = _303_v130
			_303_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v129).Cardinality(), (_132_v3).F13())
			var _304_v131 _dafny.Sequence
			_ = _304_v131
			_304_v131 = _dafny.SeqOf(_130_v2, (Companion_Default___.Fm11((_276_v115).Cardinality(), _dafny.IntOfUint32((_183_v45).Cardinality()), (_177_v39).F10(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_177_v39).F10(), Companion_T1_.CastTo_((_179_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_179_v41), 0))).Int())))).Update((_132_v3).F12(), Companion_T1_.CastTo_((_179_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_179_v41), 0))).Int())))).Cardinality(), _131_globalState)).Update(_300_cf137, _129_v1))
			var _305_v132 D16
			_ = _305_v132
			_305_v132 = Companion_D16_.Create_DC40_(_304_v131)
			var _pat_let_tv4 = _304_v131
			_ = _pat_let_tv4
			var _306_v133 _dafny.Map
			_ = _306_v133
			_306_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
				if (_303_v130).Contains((_dafny.Zero).Minus(_300_cf137)) {
					return (_303_v130).Get((_dafny.Zero).Minus(_300_cf137)).(_dafny.Sequence)
				}
				return (_132_v3).F13()
			})(), _dafny.CodePoint('c')), func(_pat_let2_0 D16) D16 {
				return func(_307_dt__update__tmp_h1 D16) D16 {
					return func(_pat_let3_0 _dafny.Sequence) D16 {
						return func(_308_dt__update_hcf76_h0 _dafny.Sequence) D16 {
							return Companion_D16_.Create_DC40_(_308_dt__update_hcf76_h0)
						}(_pat_let3_0)
					}(_pat_let_tv4)
				}(_pat_let2_0)
			}(_305_v132))
			_306_v133 = (_306_v133).Update(_dafny.Companion_Sequence_.IsPrefixOf(_183_v45, _dafny.SeqOf(_dafny.IntOfInt64(153))), _305_v132)
		} else {
			var _309___mcc_h12 _dafny.Array = _source8.Get_().(D32_DC78).Cf132
			_ = _309___mcc_h12
			var _310_cf132 _dafny.Array = _309___mcc_h12
			_ = _310_cf132
			_175_v37 = _175_v37
			var _311_v134 _dafny.Sequence
			_ = _311_v134
			_311_v134 = _dafny.SeqOf((_177_v39).F10())
			(_131_globalState).F8 = (_311_v134).Select((Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_311_v134).Cardinality()))).Uint32()).(bool)
			var _312_v135 _dafny.Sequence
			_ = _312_v135
			var _313_v136 bool
			_ = _313_v136
			var _314_v137 _dafny.Sequence
			_ = _314_v137
			var _315_v138 _dafny.Sequence
			_ = _315_v138
			var _out8 _dafny.Sequence
			_ = _out8
			var _out9 bool
			_ = _out9
			var _out10 _dafny.Sequence
			_ = _out10
			var _out11 _dafny.Sequence
			_ = _out11
			_out8, _out9, _out10, _out11 = (_273_v112).M0(_131_globalState)
			_312_v135 = _out8
			_313_v136 = _out9
			_314_v137 = _out10
			_315_v138 = _out11
			_315_v138 = (_132_v3).F13()
		}
		var _316_v139 D25
		_ = _316_v139
		_316_v139 = Companion_D25_.Create_DC63_((_132_v3).F13(), _138_v8)
		var _317_v140 _dafny.Set
		_ = _317_v140
		_317_v140 = _dafny.SetOf((_316_v139).Dtor_cf113())
		var _318_v141 _dafny.Array
		_ = _318_v141
		var _nwElement0_8 _dafny.Int = _138_v8
		_ = _nwElement0_8
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(25))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_8, 0)
		(_nw33).ArraySet1(_138_v8, 1)
		(_nw33).ArraySet1(_138_v8, 2)
		(_nw33).ArraySet1(_138_v8, 3)
		(_nw33).ArraySet1(_138_v8, 4)
		(_nw33).ArraySet1(_138_v8, 5)
		(_nw33).ArraySet1(_138_v8, 6)
		(_nw33).ArraySet1(_138_v8, 7)
		(_nw33).ArraySet1(_138_v8, 8)
		(_nw33).ArraySet1(_138_v8, 9)
		(_nw33).ArraySet1(_138_v8, 10)
		(_nw33).ArraySet1(_138_v8, 11)
		(_nw33).ArraySet1(_138_v8, 12)
		(_nw33).ArraySet1(_138_v8, 13)
		(_nw33).ArraySet1(_138_v8, 14)
		(_nw33).ArraySet1((_317_v140).Cardinality(), 15)
		(_nw33).ArraySet1(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), 16)
		(_nw33).ArraySet1(_138_v8, 17)
		(_nw33).ArraySet1(_138_v8, 18)
		(_nw33).ArraySet1(_138_v8, 19)
		(_nw33).ArraySet1(_138_v8, 20)
		(_nw33).ArraySet1(Companion_Default___.Fm33(_129_v1, (_177_v39).F10(), _138_v8, (_177_v39).F10(), _131_globalState), 21)
		(_nw33).ArraySet1(_138_v8, 22)
		(_nw33).ArraySet1(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), 23)
		(_nw33).ArraySet1(_138_v8, 24)
		_318_v141 = _nw33
		var _319_v142 _dafny.Set
		_ = _319_v142
		_319_v142 = _dafny.SetOf(_318_v141)
		var _320_v143 D8
		_ = _320_v143
		_320_v143 = Companion_D8_.Create_DC24_((_177_v39).F10(), _319_v142, _138_v8)
		var _321_v144 D16
		_ = _321_v144
		_321_v144 = Companion_D16_.Create_DC42_(_129_v1, (_320_v143).Dtor_cf44(), (_273_v112).Fm19(_138_v8, _138_v8, _131_globalState))
		var _source9 D16 = _321_v144
		_ = _source9
		if _source9.Is_DC41() {
			var _322___mcc_h13 bool = _source9.Get_().(D16_DC41).Cf77
			_ = _322___mcc_h13
			var _323___mcc_h14 _dafny.Int = _source9.Get_().(D16_DC41).Cf78
			_ = _323___mcc_h14
			var _324___mcc_h15 bool = _source9.Get_().(D16_DC41).Cf79
			_ = _324___mcc_h15
			var _325_cf79 bool = _324___mcc_h15
			_ = _325_cf79
			var _326_cf78 _dafny.Int = _323___mcc_h14
			_ = _326_cf78
			var _327_cf77 bool = _322___mcc_h13
			_ = _327_cf77
			_129_v1 = ((((_176_v38).Update((_177_v39).F10(), (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool))).Merge((_177_v39).F9())).Cardinality()).Cmp(((_dafny.Zero).Minus(_138_v8)).Plus(_326_cf78)) != 0
			(_181_v43).M5(_131_globalState)
			var _328_v145 D21
			_ = _328_v145
			_328_v145 = Companion_D21_.Create_DC53_(_272_v111)
			var _pat_let_tv5 = _272_v111
			_ = _pat_let_tv5
			_328_v145 = func(_pat_let4_0 D21) D21 {
				return func(_329_dt__update__tmp_h2 D21) D21 {
					return func(_pat_let5_0 T0) D21 {
						return func(_330_dt__update_hcf99_h0 T0) D21 {
							return Companion_D21_.Create_DC53_(_330_dt__update_hcf99_h0)
						}(_pat_let5_0)
					}(_pat_let_tv5)
				}(_pat_let4_0)
			}(_328_v145)
			(_131_globalState).F8 = !_dafny.Companion_Sequence_.Contains((_132_v3).F13(), _175_v37)
		} else if _source9.Is_DC42() {
			var _331___mcc_h16 bool = _source9.Get_().(D16_DC42).Cf80
			_ = _331___mcc_h16
			var _332___mcc_h17 _dafny.Int = _source9.Get_().(D16_DC42).Cf81
			_ = _332___mcc_h17
			var _333___mcc_h18 _dafny.Int = _source9.Get_().(D16_DC42).Cf82
			_ = _333___mcc_h18
			var _334_cf82 _dafny.Int = _333___mcc_h18
			_ = _334_cf82
			var _335_cf81 _dafny.Int = _332___mcc_h17
			_ = _335_cf81
			var _336_cf80 bool = _331___mcc_h16
			_ = _336_cf80
			var _337_v146 _dafny.Map
			_ = _337_v146
			_337_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_132_v3).F12()), ((_dafny.Zero).Minus(_335_cf81)).Minus(_334_cf82))
			var _rhs32 bool = (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)
			_ = _rhs32
			var _rhs33 _dafny.Map = _337_v146
			_ = _rhs33
			var _rhs34 bool = !(!((_132_v3).F12()) || (!((_132_v3).F12()))) || (!((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)) || ((_177_v39).F10()))
			_ = _rhs34
			var _lhs22 *GlobalState = _131_globalState
			_ = _lhs22
			var _lhs23 *GlobalState = _131_globalState
			_ = _lhs23
			_lhs22.F8 = _rhs32
			_337_v146 = _rhs33
			_lhs23.F8 = _rhs34
			_138_v8 = _334_cf82
			var _338_v147 _dafny.Set
			_ = _338_v147
			_338_v147 = _dafny.SetOf(_127_v0)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index35
			var _rhs35 bool = (_177_v39).F10()
			_ = _rhs35
			var _rhs36 _dafny.Int = (_334_cf82).Plus((_181_v43).Fm1((_177_v39).F9(), _131_globalState))
			_ = _rhs36
			var _rhs37 _dafny.Set = _338_v147
			_ = _rhs37
			var _lhs24 _dafny.Array = _127_v0
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _lhs25
			(_lhs24).ArraySet1(_rhs35, (_lhs25).Int())
			_335_cf81 = _rhs36
			_338_v147 = _rhs37
			var _339_v148 _dafny.Sequence
			_ = _339_v148
			_339_v148 = _dafny.UnicodeSeqOfUtf8Bytes("kspfey")
			_339_v148 = _dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(784))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}((func(_340_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_341_i12 _dafny.Int) _dafny.CodePoint {
					return _340_v37
				}
			})(_175_v37))))
		} else {
			var _342___mcc_h19 _dafny.Sequence = _source9.Get_().(D16_DC40).Cf76
			_ = _342___mcc_h19
			var _343_cf76 _dafny.Sequence = _342___mcc_h19
			_ = _343_cf76
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index36
			(_127_v0).ArraySet1((_177_v39).F10(), (_index36).Int())
			var _344_v149 _dafny.Sequence
			_ = _344_v149
			var _345_v150 bool
			_ = _345_v150
			var _346_v151 _dafny.Sequence
			_ = _346_v151
			var _347_v152 _dafny.Sequence
			_ = _347_v152
			var _out12 _dafny.Sequence
			_ = _out12
			var _out13 bool
			_ = _out13
			var _out14 _dafny.Sequence
			_ = _out14
			var _out15 _dafny.Sequence
			_ = _out15
			_out12, _out13, _out14, _out15 = (_181_v43).M0(_131_globalState)
			_344_v149 = _out12
			_345_v150 = _out13
			_346_v151 = _out14
			_347_v152 = _out15
			var _348_v153 _dafny.Sequence
			_ = _348_v153
			var _349_v154 bool
			_ = _349_v154
			var _350_v155 _dafny.Sequence
			_ = _350_v155
			var _351_v156 _dafny.Sequence
			_ = _351_v156
			var _out16 _dafny.Sequence
			_ = _out16
			var _out17 bool
			_ = _out17
			var _out18 _dafny.Sequence
			_ = _out18
			var _out19 _dafny.Sequence
			_ = _out19
			_out16, _out17, _out18, _out19 = (_132_v3).M0(_131_globalState)
			_348_v153 = _out16
			_349_v154 = _out17
			_350_v155 = _out18
			_351_v156 = _out19
			(_131_globalState).F8 = _345_v150
		}
		var _352_v157 _dafny.Map
		_ = _352_v157
		_352_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _138_v8)
		_138_v8 = Companion_Default___.SafeModuloInt(_138_v8, Companion_Default___.SafeModuloInt((_352_v157).Cardinality(), _138_v8))
	} else {
		var _353_v158 _dafny.Sequence
		_ = _353_v158
		var _354_v159 bool
		_ = _354_v159
		var _355_v160 _dafny.Sequence
		_ = _355_v160
		var _356_v161 _dafny.Sequence
		_ = _356_v161
		var _out20 _dafny.Sequence
		_ = _out20
		var _out21 bool
		_ = _out21
		var _out22 _dafny.Sequence
		_ = _out22
		var _out23 _dafny.Sequence
		_ = _out23
		_out20, _out21, _out22, _out23 = (_273_v112).M0(_131_globalState)
		_353_v158 = _out20
		_354_v159 = _out21
		_355_v160 = _out22
		_356_v161 = _out23
		var _357_v162 *C4
		_ = _357_v162
		var _nw34 *C4 = New_C4_()
		_ = _nw34
		_nw34.Ctor__(_176_v38, ((_132_v3).F12()) || ((_177_v39).F10()))
		_357_v162 = _nw34
		var _358_v163 _dafny.Map
		_ = _358_v163
		_358_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_272_v111, _129_v1)
		var _359_v164 _dafny.Map
		_ = _359_v164
		_359_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v163, _138_v8)
		var _360_v165 D32
		_ = _360_v165
		_360_v165 = Companion_D32_.Create_DC80_(_138_v8, _138_v8, _dafny.IntOfInt64(665))
		var _361_v166 _dafny.MultiSet
		_ = _361_v166
		_361_v166 = _dafny.MultiSetOf(_dafny.IntOfInt64(879))
		var _362_v167 _dafny.Array
		_ = _362_v167
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
		_ = _nw35
		_362_v167 = _nw35
		var _363_v168 _dafny.Map
		_ = _363_v168
		_363_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_362_v167, _356_v161)
		var _364_v169 _dafny.Array
		_ = _364_v169
		var _nwElement0_9 _dafny.Sequence = Companion_Default___.Fm22(_138_v8, _129_v1, (_177_v39).F10(), _138_v8, _131_globalState)
		_ = _nwElement0_9
		var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(23))
		_ = _nw36
		(_nw36).ArraySet1(_nwElement0_9, 0)
		(_nw36).ArraySet1(_353_v158, 1)
		(_nw36).ArraySet1(_356_v161, 2)
		(_nw36).ArraySet1(_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _175_v37), 3)
		(_nw36).ArraySet1((_132_v3).F13(), 4)
		(_nw36).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yieoik"), 5)
		(_nw36).ArraySet1(_356_v161, 6)
		(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), _353_v158), 7)
		(_nw36).ArraySet1(_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex((_359_v164).Cardinality(), _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _175_v37), 8)
		(_nw36).ArraySet1(Companion_Default___.Fm13((_360_v165).Dtor_cf138(), _131_globalState), 9)
		(_nw36).ArraySet1(_356_v161, 10)
		(_nw36).ArraySet1(_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _dafny.CodePoint('p')), 11)
		(_nw36).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(476))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}((func(_365_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_366_i13 _dafny.Int) _dafny.CodePoint {
				return _365_v37
			}
		})(_175_v37))), 12)
		(_nw36).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("a"), 13)
		(_nw36).ArraySet1(Companion_Default___.Fm45((_132_v3).F12(), (func() _dafny.Int {
			if (_361_v166).Contains(_138_v8) {
				return (_361_v166).Multiplicity(_138_v8)
			}
			return _138_v8
		})(), (_132_v3).F12(), _131_globalState), 14)
		(_nw36).ArraySet1(_353_v158, 15)
		(_nw36).ArraySet1(_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_353_v158).Cardinality()), _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _dafny.CodePoint('p')), 16)
		(_nw36).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("iybtk"), 17)
		(_nw36).ArraySet1(Companion_Default___.Fm45(_354_v159, Companion_Default___.Fm33(_129_v1, _354_v159, (_dafny.Zero).Minus(_138_v8), (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool), _131_globalState), !(!(!((_132_v3).F12()))), _131_globalState), 18)
		(_nw36).ArraySet1((func() _dafny.Sequence {
			if (_177_v39).F10() {
				return _dafny.UnicodeSeqOfUtf8Bytes("wgwipapdd")
			}
			return (_132_v3).F13()
		})(), 19)
		(_nw36).ArraySet1(_353_v158, 20)
		(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_353_v158, _353_v158), 21)
		(_nw36).ArraySet1((func() _dafny.Sequence {
			if (_363_v168).Contains(_362_v167) {
				return (_363_v168).Get(_362_v167).(_dafny.Sequence)
			}
			return (_132_v3).F13()
		})(), 22)
		_364_v169 = _nw36
		var _367_v170 _dafny.Map
		_ = _367_v170
		_367_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _364_v169)
		var _rhs38 _dafny.Int = (_357_v162).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_177_v39).F10(), (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)), _131_globalState)
		_ = _rhs38
		var _rhs39 _dafny.Sequence = (_132_v3).F13()
		_ = _rhs39
		var _rhs40 _dafny.Array = (func() _dafny.Array {
			if (_367_v170).Contains(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality())) {
				return (_367_v170).Get(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality())).(_dafny.Array)
			}
			return _364_v169
		})()
		_ = _rhs40
		var _rhs41 bool = _129_v1
		_ = _rhs41
		_138_v8 = _rhs38
		_356_v161 = _rhs39
		_364_v169 = _rhs40
		_354_v159 = _rhs41
		var _368_v171 _dafny.Map
		_ = _368_v171
		_368_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_132_v3).F12(), _127_v0)
		_368_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_177_v39).F10()) || (_129_v1), _127_v0)
		_129_v1 = (_177_v39).F10()
	}
	var _369_v172 _dafny.Sequence
	_ = _369_v172
	_369_v172 = _dafny.SeqOf((_177_v39).F10(), (func() bool {
		if _129_v1 {
			return (_177_v39).F10()
		}
		return !(_129_v1)
	})(), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(534))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg41 _dafny.Int) interface{} {
			return coer41(arg41)
		}
	}((func(_370_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_371_i14 _dafny.Int) _dafny.CodePoint {
			return _370_v37
		}
	})(_175_v37))), (_132_v3).F13()))
	if (_369_v172).Select((Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_369_v172).Cardinality()))).Uint32()).(bool) {
		var _372_v173 _dafny.Map
		_ = _372_v173
		var _373_v174 bool
		_ = _373_v174
		var _out24 _dafny.Map
		_ = _out24
		var _out25 bool
		_ = _out25
		_out24, _out25 = (_272_v111).M1(_175_v37, _131_globalState)
		_372_v173 = _out24
		_373_v174 = _out25
		(_181_v43).M5(_131_globalState)
		var _374_v175 _dafny.Sequence
		_ = _374_v175
		var _375_v176 bool
		_ = _375_v176
		var _376_v177 _dafny.Sequence
		_ = _376_v177
		var _377_v178 _dafny.Sequence
		_ = _377_v178
		var _out26 _dafny.Sequence
		_ = _out26
		var _out27 bool
		_ = _out27
		var _out28 _dafny.Sequence
		_ = _out28
		var _out29 _dafny.Sequence
		_ = _out29
		_out26, _out27, _out28, _out29 = (_273_v112).M0(_131_globalState)
		_374_v175 = _out26
		_375_v176 = _out27
		_376_v177 = _out28
		_377_v178 = _out29
		var _378_v179 *C0
		_ = _378_v179
		var _nw37 *C0 = New_C0_()
		_ = _nw37
		_nw37.Ctor__()
		_378_v179 = _nw37
		var _379_v181 D33
		_ = _379_v181
		_379_v181 = Companion_D33_.Create_DC81_(func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(547), _dafny.IntOfInt64(232))); ; {
				_compr_27, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _380_v180 _dafny.Int
				_380_v180 = interface{}(_compr_27).(_dafny.Int)
				if ((_dafny.IntOfInt64(547)).Cmp(_380_v180) <= 0) && ((_380_v180).Cmp(_dafny.IntOfInt64(232)) < 0) {
					_coll27.Add((_380_v180).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}((func(_381_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_382_i15 _dafny.Int) _dafny.CodePoint {
							return _381_v37
						}
					})(_175_v37)))).Cardinality())), _183_v45)
				}
			}
			return _coll27.ToMap()
		}())
		var _383_v182 D33
		_ = _383_v182
		_383_v182 = Companion_D33_.Create_DC83_(_379_v181)
		var _384_v183 D33
		_ = _384_v183
		_384_v183 = Companion_D33_.Create_DC83_(_379_v181)
		_384_v183 = _384_v183
	} else {
		var _385_v184 _dafny.Sequence
		_ = _385_v184
		_385_v184 = _dafny.UnicodeSeqOfUtf8Bytes("jytbikv")
		_385_v184 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _175_v37), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_132_v3).F13(), (Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))).Uint32(), _175_v37)).Cardinality()))).Uint32(), _175_v37)
		var _386_v185 _dafny.MultiSet
		_ = _386_v185
		_386_v185 = _dafny.MultiSetOf(_138_v8)
		_386_v185 = (Companion_Default___.Fm61(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(986), _138_v8), _131_globalState))
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _index37
		(_127_v0).ArraySet1(!((_132_v3).F12()) || (!(((_132_v3).F12()) == ((_369_v172).Select((Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_369_v172).Cardinality()))).Uint32()).(bool)))), (_index37).Int())
		_138_v8 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_385_v184, _385_v184)).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-69))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg43 _dafny.Int) interface{} {
				return coer43(arg43)
			}
		}((func(_387_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_388_i16 _dafny.Int) _dafny.CodePoint {
				return _387_v37
			}
		})(_175_v37)))).Cardinality()))
		var _389_v186 D14
		_ = _389_v186
		_389_v186 = Companion_D14_.Create_DC36_(_129_v1, _385_v184, (_132_v3).F13())
		var _source10 D14 = _389_v186
		_ = _source10
		if _source10.Is_DC35() {
			var _390___mcc_h20 _dafny.Int = _source10.Get_().(D14_DC35).Cf62
			_ = _390___mcc_h20
			var _391___mcc_h21 bool = _source10.Get_().(D14_DC35).Cf63
			_ = _391___mcc_h21
			var _392___mcc_h22 _dafny.CodePoint = _source10.Get_().(D14_DC35).Cf64
			_ = _392___mcc_h22
			var _393_cf64 _dafny.CodePoint = _392___mcc_h22
			_ = _393_cf64
			var _394_cf63 bool = _391___mcc_h21
			_ = _394_cf63
			var _395_cf62 _dafny.Int = _390___mcc_h20
			_ = _395_cf62
			var _396_v187 _dafny.Map
			_ = _396_v187
			_396_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _395_cf62)
			_395_cf62 = (func() _dafny.Int {
				if (_396_v187).Contains((Companion_D5_.Create_DC14_(_dafny.IntOfUint32((_385_v184).Cardinality()), (_132_v3).F12())).Dtor_cf25()) {
					return (_396_v187).Get((Companion_D5_.Create_DC14_(_dafny.IntOfUint32((_385_v184).Cardinality()), (_132_v3).F12())).Dtor_cf25()).(_dafny.Int)
				}
				return _395_cf62
			})()
			_395_cf62 = _138_v8
			var _397_v188 _dafny.Array
			_ = _397_v188
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_4
			var _nw38 _dafny.Array
			_ = _nw38
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw38 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Int = (func(_398_v8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_399_i17 _dafny.Int) _dafny.Int {
						return (_399_i17).Minus(_398_v8)
					}
				})(_138_v8)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw38 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw38).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw38).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_397_v188 = _nw38
			var _400_v189 _dafny.Set
			_ = _400_v189
			_400_v189 = _dafny.SetOf(_138_v8)
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_397_v188), 0))
			_ = _index38
			(_397_v188).ArraySet1((_400_v189).Cardinality(), (_index38).Int())
			var _401_v190 _dafny.Map
			_ = _401_v190
			_401_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_cf62, _183_v45)
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index39
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_397_v188), 0))
			_ = _index40
			var _rhs42 bool = (_132_v3).F12()
			_ = _rhs42
			var _rhs43 bool = !(_386_v185).Contains(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()))
			_ = _rhs43
			var _rhs44 _dafny.Array = _127_v0
			_ = _rhs44
			var _rhs45 _dafny.Int = (_dafny.Zero).Minus(((_401_v190).Merge(func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(492), _dafny.IntOfInt64(183))); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _402_v191 _dafny.Int
					_402_v191 = interface{}(_compr_28).(_dafny.Int)
					if ((_dafny.IntOfInt64(492)).Cmp(_402_v191) <= 0) && ((_402_v191).Cmp(_dafny.IntOfInt64(183)) < 0) {
						_coll28.Add(Companion_Default___.SafeDivisionInt(_402_v191, _dafny.IntOfInt64(-771)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(996))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}((func(_403_cf62 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_404_i18 _dafny.Int) _dafny.Int {
								return _403_cf62
							}
						})(_395_cf62))))
					}
				}
				return _coll28.ToMap()
			}())).Cardinality())
			_ = _rhs45
			var _lhs26 _dafny.Array = _127_v0
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _lhs27
			var _lhs28 *GlobalState = _131_globalState
			_ = _lhs28
			var _lhs29 _dafny.Array = _397_v188
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_397_v188), 0))
			_ = _lhs30
			(_lhs26).ArraySet1(_rhs42, (_lhs27).Int())
			_lhs28.F8 = _rhs43
			_127_v0 = _rhs44
			(_lhs29).ArraySet1(_rhs45, (_lhs30).Int())
			_395_cf62 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_395_cf62, _dafny.IntOfInt64(270)), _395_cf62)
		} else if _source10.Is_DC36() {
			var _405___mcc_h23 bool = _source10.Get_().(D14_DC36).Cf65
			_ = _405___mcc_h23
			var _406___mcc_h24 _dafny.Sequence = _source10.Get_().(D14_DC36).Cf66
			_ = _406___mcc_h24
			var _407___mcc_h25 _dafny.Sequence = _source10.Get_().(D14_DC36).Cf67
			_ = _407___mcc_h25
			var _408_cf67 _dafny.Sequence = _407___mcc_h25
			_ = _408_cf67
			var _409_cf66 _dafny.Sequence = _406___mcc_h24
			_ = _409_cf66
			var _410_cf65 bool = _405___mcc_h23
			_ = _410_cf65
			var _411_v192 _dafny.Map
			_ = _411_v192
			var _412_v193 bool
			_ = _412_v193
			var _out30 _dafny.Map
			_ = _out30
			var _out31 bool
			_ = _out31
			_out30, _out31 = (_272_v111).M1(_175_v37, _131_globalState)
			_411_v192 = _out30
			_412_v193 = _out31
			var _413_v194 _dafny.Map
			_ = _413_v194
			var _414_v195 bool
			_ = _414_v195
			var _out32 _dafny.Map
			_ = _out32
			var _out33 bool
			_ = _out33
			_out32, _out33 = (_177_v39).M1(_175_v37, _131_globalState)
			_413_v194 = _out32
			_414_v195 = _out33
			_410_cf65 = _dafny.Companion_Sequence_.Contains(_409_cf66, _175_v37)
			_369_v172 = (func() _dafny.Sequence {
				if (_177_v39).F10() {
					return _369_v172
				}
				return _369_v172
			})()
		} else {
			var _415___mcc_h26 _dafny.Map = _source10.Get_().(D14_DC34).Cf61
			_ = _415___mcc_h26
			var _416_cf61 _dafny.Map = _415___mcc_h26
			_ = _416_cf61
			var _417_v196 _dafny.Map
			_ = _417_v196
			_417_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v8, _dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13()))
			_417_v196 = (_417_v196).Update(_138_v8, _385_v184)
			var _418_v197 _dafny.Map
			_ = _418_v197
			_418_v197 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v37, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v1, _138_v8)).Cardinality())
			var _419_v198 _dafny.Set
			_ = _419_v198
			_419_v198 = _dafny.SetOf(_138_v8, (_dafny.Zero).Minus(_138_v8), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_132_v3).F12()), (Companion_Default___.SafeIndex((_418_v197).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_132_v3).F12())).Cardinality()))).Uint32(), !((_132_v3).F12()))).Cardinality()))
			(_131_globalState).F8 = (_419_v198).IsDisjointFrom(_419_v198)
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index41
			(_127_v0).ArraySet1((_138_v8).Cmp((_419_v198).Cardinality()) < 0, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _index42
			var _rhs46 bool = (_132_v3).F12()
			_ = _rhs46
			var _rhs47 _dafny.Array = _127_v0
			_ = _rhs47
			var _lhs31 _dafny.Array = _127_v0
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
			_ = _lhs32
			(_lhs31).ArraySet1(_rhs46, (_lhs32).Int())
			_127_v0 = _rhs47
		}
	}
	var _source11 D16 = Companion_D16_.Create_DC41_(!(_129_v1) || ((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)), (_dafny.Zero).Minus(_138_v8), (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool))
	_ = _source11
	if _source11.Is_DC41() {
		var _420___mcc_h27 bool = _source11.Get_().(D16_DC41).Cf77
		_ = _420___mcc_h27
		var _421___mcc_h28 _dafny.Int = _source11.Get_().(D16_DC41).Cf78
		_ = _421___mcc_h28
		var _422___mcc_h29 bool = _source11.Get_().(D16_DC41).Cf79
		_ = _422___mcc_h29
		var _423_cf79 bool = _422___mcc_h29
		_ = _423_cf79
		var _424_cf78 _dafny.Int = _421___mcc_h28
		_ = _424_cf78
		var _425_cf77 bool = _420___mcc_h27
		_ = _425_cf77
		(_131_globalState).F8 = true
		var _426_v199 _dafny.Sequence
		_ = _426_v199
		_426_v199 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v1, _dafny.IntOfUint32(((_132_v3).F13()).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_177_v39).F10(), _138_v8))
		var _427_v200 _dafny.Map
		_ = _427_v200
		_427_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v1, _424_cf78)
		_138_v8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_426_v199, _dafny.SeqOf(_427_v200)), _dafny.Companion_Sequence_.Concatenate(_426_v199, _dafny.SeqOf(_427_v200)))).Cardinality())
		_138_v8 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_424_cf78))
		var _428_v201 _dafny.Set
		_ = _428_v201
		_428_v201 = _dafny.SetOf(_424_cf78, _dafny.IntOfUint32((_369_v172).Cardinality()))
		var _429_v202 _dafny.Set
		_ = _429_v202
		_429_v202 = _dafny.SetOf((_428_v201).Cardinality(), _138_v8, _424_cf78, _138_v8, _424_cf78)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _index43
		var _rhs48 bool = (_429_v202).IsSubsetOf(_428_v201)
		_ = _rhs48
		var _rhs49 _dafny.Int = _424_cf78
		_ = _rhs49
		var _rhs50 bool = _425_cf77
		_ = _rhs50
		var _lhs33 *GlobalState = _131_globalState
		_ = _lhs33
		var _lhs34 _dafny.Array = _127_v0
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _lhs35
		_lhs33.F8 = _rhs48
		_424_cf78 = _rhs49
		(_lhs34).ArraySet1(_rhs50, (_lhs35).Int())
	} else if _source11.Is_DC42() {
		var _430___mcc_h30 bool = _source11.Get_().(D16_DC42).Cf80
		_ = _430___mcc_h30
		var _431___mcc_h31 _dafny.Int = _source11.Get_().(D16_DC42).Cf81
		_ = _431___mcc_h31
		var _432___mcc_h32 _dafny.Int = _source11.Get_().(D16_DC42).Cf82
		_ = _432___mcc_h32
		var _433_cf82 _dafny.Int = _432___mcc_h32
		_ = _433_cf82
		var _434_cf81 _dafny.Int = _431___mcc_h31
		_ = _434_cf81
		var _435_cf80 bool = _430___mcc_h30
		_ = _435_cf80
		var _436_v203 _dafny.Array
		_ = _436_v203
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw39
		_436_v203 = _nw39
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_436_v203), 0))
		_ = _index44
		(_436_v203).ArraySet1(_434_cf81, (_index44).Int())
		var _437_v204 _dafny.Array
		_ = _437_v204
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(12))
		_ = _nw40
		_437_v204 = _nw40
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_436_v203), 0))
		_ = _index45
		var _rhs51 _dafny.Int = (_dafny.Zero).Minus(_433_cf82)
		_ = _rhs51
		var _rhs52 _dafny.Array = _437_v204
		_ = _rhs52
		var _lhs36 _dafny.Array = _436_v203
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_436_v203), 0))
		_ = _lhs37
		(_lhs36).ArraySet1(_rhs51, (_lhs37).Int())
		_437_v204 = _rhs52
		var _438_v205 _dafny.Sequence
		_ = _438_v205
		var _439_v206 bool
		_ = _439_v206
		var _440_v207 _dafny.Sequence
		_ = _440_v207
		var _441_v208 _dafny.Sequence
		_ = _441_v208
		var _out34 _dafny.Sequence
		_ = _out34
		var _out35 bool
		_ = _out35
		var _out36 _dafny.Sequence
		_ = _out36
		var _out37 _dafny.Sequence
		_ = _out37
		_out34, _out35, _out36, _out37 = (_177_v39).M0(_131_globalState)
		_438_v205 = _out34
		_439_v206 = _out35
		_440_v207 = _out36
		_441_v208 = _out37
		var _442_v209 _dafny.Set
		_ = _442_v209
		_442_v209 = _dafny.SetOf((_177_v39).F10(), (_132_v3).F12())
		var _443_v210 *C7
		_ = _443_v210
		var _nw41 *C7 = New_C7_()
		_ = _nw41
		_nw41.Ctor__((_442_v209).Union(Companion_Default___.Fm34(_434_cf81, true, (_132_v3).F12(), (_177_v39).F10(), _131_globalState)), _176_v38, (func() bool {
			if _439_v206 {
				return _439_v206
			}
			return (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)
		})())
		_443_v210 = _nw41
		var _444_v211 _dafny.Array
		_ = _444_v211
		var _nw42 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
		_ = _nw42
		_444_v211 = _nw42
		var _445_v212 *C6
		_ = _445_v212
		var _nw43 *C6 = New_C6_()
		_ = _nw43
		_nw43.Ctor__(_176_v38, false)
		_445_v212 = _nw43
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_444_v211), 0))
		_ = _index46
		(_444_v211).ArraySet1(_445_v212, (_index46).Int())
		var _446_v213 _dafny.Array
		_ = _446_v213
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(24))
		_ = _nw44
		_446_v213 = _nw44
		var _447_v214 _dafny.Sequence
		_ = _447_v214
		_447_v214 = _dafny.SeqOf(_446_v213)
		var _448_v215 D15
		_ = _448_v215
		_448_v215 = Companion_D15_.Create_DC37_((_447_v214).Select((Companion_Default___.SafeIndex(_138_v8, _dafny.IntOfUint32((_447_v214).Cardinality()))).Uint32()).(_dafny.Array))
		var _449_v216 _dafny.Sequence
		_ = _449_v216
		_449_v216 = _dafny.SeqOf(_dafny.SetOf(_138_v8, _138_v8))
		var _450_v217 _dafny.Set
		_ = _450_v217
		_450_v217 = _dafny.SetOf((_dafny.SetOf(true, (_132_v3).F12())).Cardinality(), _138_v8)
		var _451_v218 _dafny.Set
		_ = _451_v218
		_451_v218 = _dafny.SetOf(_438_v205, (_132_v3).F13(), (_132_v3).F13())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_444_v211), 0))
		_ = _index47
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_436_v203), 0))
		_ = _index48
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _index49
		var _rhs53 *C7 = _443_v210
		_ = _rhs53
		var _rhs54 *C6 = _445_v212
		_ = _rhs54
		var _rhs55 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_449_v216, _dafny.SeqOf(_450_v217))).Cardinality()))
		_ = _rhs55
		var _rhs56 D15 = Companion_D15_.Create_DC37_(_446_v213)
		_ = _rhs56
		var _rhs57 bool = !((_451_v218).Contains(_dafny.Companion_Sequence_.Concatenate((_132_v3).F13(), (_132_v3).F13())))
		_ = _rhs57
		var _lhs38 _dafny.Array = _444_v211
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_444_v211), 0))
		_ = _lhs39
		var _lhs40 _dafny.Array = _436_v203
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(552), _dafny.ArrayLen((_436_v203), 0))
		_ = _lhs41
		var _lhs42 _dafny.Array = _127_v0
		_ = _lhs42
		var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))
		_ = _lhs43
		_443_v210 = _rhs53
		(_lhs38).ArraySet1(_rhs54, (_lhs39).Int())
		(_lhs40).ArraySet1(_rhs55, (_lhs41).Int())
		_448_v215 = _rhs56
		(_lhs42).ArraySet1(_rhs57, (_lhs43).Int())
		_435_cf80 = (_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool)
	} else {
		var _452___mcc_h33 _dafny.Sequence = _source11.Get_().(D16_DC40).Cf76
		_ = _452___mcc_h33
		var _453_cf76 _dafny.Sequence = _452___mcc_h33
		_ = _453_cf76
		_183_v45 = _183_v45
		_138_v8 = _138_v8
		_127_v0 = _127_v0
		var _454_v219 _dafny.Map
		_ = _454_v219
		_454_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_132_v3).F13()).Cardinality()), _369_v172)
		_138_v8 = Companion_Default___.Fm33((_127_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_127_v0), 0))).Int()).(bool), (_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_454_v219).Contains(_dafny.IntOfInt64(-795)) {
				return (_454_v219).Get(_dafny.IntOfInt64(-795)).(_dafny.Sequence)
			}
			return (func() _dafny.Sequence {
				if (_454_v219).Contains(_138_v8) {
					return (_454_v219).Get(_138_v8).(_dafny.Sequence)
				}
				return _369_v172
			})()
		})()).Cardinality())).Cmp(_138_v8) <= 0, (_138_v8).Times(_138_v8), (_132_v3).F12(), _131_globalState)
	}
	var _455_v220 _dafny.Sequence
	_ = _455_v220
	var _456_v221 bool
	_ = _456_v221
	var _457_v222 _dafny.Sequence
	_ = _457_v222
	var _458_v223 _dafny.Sequence
	_ = _458_v223
	var _out38 _dafny.Sequence
	_ = _out38
	var _out39 bool
	_ = _out39
	var _out40 _dafny.Sequence
	_ = _out40
	var _out41 _dafny.Sequence
	_ = _out41
	_out38, _out39, _out40, _out41 = (_177_v39).M0(_131_globalState)
	_455_v220 = _out38
	_456_v221 = _out39
	_457_v222 = _out40
	_458_v223 = _out41
	_dafny.Print((_127_v0).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v0).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_129_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(4), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_131_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState.F4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(4), true), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_132_v3).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_132_v3).F13(), _dafny.SeqOf(_dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'), _dafny.CodePoint('n'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v7).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_138_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_v37)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v38).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_177_v39).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_177_v39).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_178_v40).Dtor_cf104()).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_178_v40).Dtor_cf104()).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.Zero).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.Zero).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.One).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.One).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((Companion_T1_.CastTo_((_179_v41).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_180_v42).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_182_v44).Dtor_cf127()).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_v44).Dtor_cf127()).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_183_v45, _dafny.SeqOf(_dafny.IntOfInt64(-556))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_184_v46).Dtor_cf4(), _dafny.SeqOf(_dafny.IntOfInt64(-556))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_271_v110).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v113).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_369_v172, _dafny.SeqOf(true, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_455_v220.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_456_v221)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_457_v222, _dafny.SeqOf(_dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_458_v223.VerbatimString(false))
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
	Cf2 bool
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 bool, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
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

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() bool {
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
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

type D1_DC3 struct {
	Cf5 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf6 bool
	Cf7 bool
	Cf8 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 bool, Cf7 bool, Cf8 _dafny.Int) D1 {
	return D1{D1_DC4{Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC5 struct {
	Cf9 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 D1) D1 {
	return D1{D1_DC5{Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.EmptySeq)
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf9() D1 {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf5.VerbatimString(true) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf10 _dafny.CodePoint
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.CodePoint) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D2) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10 == data2.Cf10
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
	Cf13 bool
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 bool, Cf13 bool) D3 {
	return D3{D3_DC8{Cf12, Cf13}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf11 _dafny.Array
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Array) D3 {
	return D3{D3_DC7{Cf11}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC9 struct {
	Cf14 D3
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf14 D3) D3 {
	return D3{D3_DC9{Cf14}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(false, false)
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) Dtor_cf11() _dafny.Array {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) Dtor_cf14() D3 {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ")"
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
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11 == data2.Cf11
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
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

type D4_DC11 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.Set
	Cf18 bool
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf16 _dafny.Int, Cf17 _dafny.Set, Cf18 bool) D4 {
	return D4{D4_DC11{Cf16, Cf17, Cf18}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf19 bool
	Cf20 _dafny.Int
	Cf21 _dafny.CodePoint
	Cf22 _dafny.Int
	Cf23 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf19 bool, Cf20 _dafny.Int, Cf21 _dafny.CodePoint, Cf22 _dafny.Int, Cf23 _dafny.Int) D4 {
	return D4{D4_DC12{Cf19, Cf20, Cf21, Cf22, Cf23}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC10 struct {
	Cf15 _dafny.Map
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf15 _dafny.Map) D4 {
	return D4{D4_DC10{Cf15}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.Zero, _dafny.EmptySet, false)
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Set {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) Dtor_cf18() bool {
	return _this.Get_().(D4_DC11).Cf18
}

func (_this D4) Dtor_cf19() bool {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf21() _dafny.CodePoint {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf23
}

func (_this D4) Dtor_cf15() _dafny.Map {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Equals(data2.Cf17) && data1.Cf18 == data2.Cf18
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf15.Equals(data2.Cf15)
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

type D5_DC14 struct {
	Cf25 _dafny.Int
	Cf26 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf25 _dafny.Int, Cf26 bool) D5 {
	return D5{D5_DC14{Cf25, Cf26}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf24 _dafny.Array
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf24 _dafny.Array) D5 {
	return D5{D5_DC13{Cf24}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC15 struct {
	Cf27 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf27 D5) D5 {
	return D5{D5_DC15{Cf27}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC14_(_dafny.Zero, false)
}

func (_this D5) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) Dtor_cf26() bool {
	return _this.Get_().(D5_DC14).Cf26
}

func (_this D5) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf24
}

func (_this D5) Dtor_cf27() D5 {
	return _this.Get_().(D5_DC15).Cf27
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26 == data2.Cf26
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	Cf29 bool
	Cf30 bool
	Cf31 _dafny.Int
	Cf32 _dafny.Int
	Cf33 _dafny.Int
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf29 bool, Cf30 bool, Cf31 _dafny.Int, Cf32 _dafny.Int, Cf33 _dafny.Int) D6 {
	return D6{D6_DC17{Cf29, Cf30, Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf34 _dafny.CodePoint
	Cf35 _dafny.Int
	Cf36 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf34 _dafny.CodePoint, Cf35 _dafny.Int, Cf36 _dafny.Int) D6 {
	return D6{D6_DC18{Cf34, Cf35, Cf36}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC19 struct {
	Cf37 bool
	Cf38 _dafny.CodePoint
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf37 bool, Cf38 _dafny.CodePoint) D6 {
	return D6{D6_DC19{Cf37, Cf38}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC16 struct {
	Cf28 _dafny.Map
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf28 _dafny.Map) D6 {
	return D6{D6_DC16{Cf28}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC20 struct {
	Cf39 D6
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf39 D6) D6 {
	return D6{D6_DC20{Cf39}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false, false, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC17).Cf29
}

func (_this D6) Dtor_cf30() bool {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf32
}

func (_this D6) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf33
}

func (_this D6) Dtor_cf34() _dafny.CodePoint {
	return _this.Get_().(D6_DC18).Cf34
}

func (_this D6) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf35
}

func (_this D6) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf36
}

func (_this D6) Dtor_cf37() bool {
	return _this.Get_().(D6_DC19).Cf37
}

func (_this D6) Dtor_cf38() _dafny.CodePoint {
	return _this.Get_().(D6_DC19).Cf38
}

func (_this D6) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D6_DC16).Cf28
}

func (_this D6) Dtor_cf39() D6 {
	return _this.Get_().(D6_DC20).Cf39
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Cmp(data2.Cf36) == 0
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf39.Equals(data2.Cf39)
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

type D7_DC22 struct {
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_() D7 {
	return D7{D7_DC22{}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC21 struct {
	Cf40 _dafny.Set
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf40 _dafny.Set) D7 {
	return D7{D7_DC21{Cf40}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC22_()
}

func (_this D7) Dtor_cf40() _dafny.Set {
	return _this.Get_().(D7_DC21).Cf40
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC22:
		{
			return "D7.DC22"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC22:
		{
			_, ok := other.Get_().(D7_DC22)
			return ok
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf40.Equals(data2.Cf40)
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
	Cf42 bool
	Cf43 _dafny.Set
	Cf44 _dafny.Int
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf42 bool, Cf43 _dafny.Set, Cf44 _dafny.Int) D8 {
	return D8{D8_DC24{Cf42, Cf43, Cf44}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC23 struct {
	Cf41 _dafny.Sequence
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf41 _dafny.Sequence) D8 {
	return D8{D8_DC23{Cf41}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC24_(false, _dafny.EmptySet, _dafny.Zero)
}

func (_this D8) Dtor_cf42() bool {
	return _this.Get_().(D8_DC24).Cf42
}

func (_this D8) Dtor_cf43() _dafny.Set {
	return _this.Get_().(D8_DC24).Cf43
}

func (_this D8) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf44
}

func (_this D8) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D8_DC23).Cf41
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf42 == data2.Cf42 && data1.Cf43.Equals(data2.Cf43) && data1.Cf44.Cmp(data2.Cf44) == 0
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf41.Equals(data2.Cf41)
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
	Cf45 _dafny.Map
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf45 _dafny.Map) D9 {
	return D9{D9_DC25{Cf45}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D9) Dtor_cf45() _dafny.Map {
	return _this.Get_().(D9_DC25).Cf45
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf45) + ")"
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
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D10_DC26 struct {
	Cf46 _dafny.Map
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf46 _dafny.Map) D10 {
	return D10{D10_DC26{Cf46}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D10) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D10_DC26).Cf46
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf46) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf46.Equals(data2.Cf46)
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

type D11_DC28 struct {
	Cf48 _dafny.Int
	Cf49 bool
	Cf50 _dafny.Map
	Cf51 bool
	Cf52 _dafny.Sequence
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf48 _dafny.Int, Cf49 bool, Cf50 _dafny.Map, Cf51 bool, Cf52 _dafny.Sequence) D11 {
	return D11{D11_DC28{Cf48, Cf49, Cf50, Cf51, Cf52}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC27 struct {
	Cf47 _dafny.MultiSet
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf47 _dafny.MultiSet) D11 {
	return D11{D11_DC27{Cf47}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC29 struct {
	Cf53 D11
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf53 D11) D11 {
	return D11{D11_DC29{Cf53}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC28_(_dafny.Zero, false, _dafny.EmptyMap, false, _dafny.EmptySeq)
}

func (_this D11) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D11_DC28).Cf48
}

func (_this D11) Dtor_cf49() bool {
	return _this.Get_().(D11_DC28).Cf49
}

func (_this D11) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D11_DC28).Cf50
}

func (_this D11) Dtor_cf51() bool {
	return _this.Get_().(D11_DC28).Cf51
}

func (_this D11) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D11_DC28).Cf52
}

func (_this D11) Dtor_cf47() _dafny.MultiSet {
	return _this.Get_().(D11_DC27).Cf47
}

func (_this D11) Dtor_cf53() D11 {
	return _this.Get_().(D11_DC29).Cf53
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + data.Cf52.VerbatimString(true) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf53) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf48.Cmp(data2.Cf48) == 0 && data1.Cf49 == data2.Cf49 && data1.Cf50.Equals(data2.Cf50) && data1.Cf51 == data2.Cf51 && data1.Cf52.Equals(data2.Cf52)
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf47.Equals(data2.Cf47)
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf53.Equals(data2.Cf53)
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

type D12_DC31 struct {
	Cf55 bool
	Cf56 _dafny.Sequence
	Cf57 _dafny.Sequence
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf55 bool, Cf56 _dafny.Sequence, Cf57 _dafny.Sequence) D12 {
	return D12{D12_DC31{Cf55, Cf56, Cf57}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC32 struct {
	Cf58 _dafny.Array
	Cf59 _dafny.Sequence
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf58 _dafny.Array, Cf59 _dafny.Sequence) D12 {
	return D12{D12_DC32{Cf58, Cf59}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC30 struct {
	Cf54 _dafny.Array
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf54 _dafny.Array) D12 {
	return D12{D12_DC30{Cf54}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC31_(false, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D12) Dtor_cf55() bool {
	return _this.Get_().(D12_DC31).Cf55
}

func (_this D12) Dtor_cf56() _dafny.Sequence {
	return _this.Get_().(D12_DC31).Cf56
}

func (_this D12) Dtor_cf57() _dafny.Sequence {
	return _this.Get_().(D12_DC31).Cf57
}

func (_this D12) Dtor_cf58() _dafny.Array {
	return _this.Get_().(D12_DC32).Cf58
}

func (_this D12) Dtor_cf59() _dafny.Sequence {
	return _this.Get_().(D12_DC32).Cf59
}

func (_this D12) Dtor_cf54() _dafny.Array {
	return _this.Get_().(D12_DC30).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf55) + ", " + data.Cf56.VerbatimString(true) + ", " + data.Cf57.VerbatimString(true) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf55 == data2.Cf55 && data1.Cf56.Equals(data2.Cf56) && data1.Cf57.Equals(data2.Cf57)
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59.Equals(data2.Cf59)
		}
	case D12_DC30:
		{
			data2, ok := other.Get_().(D12_DC30)
			return ok && data1.Cf54 == data2.Cf54
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

type D13_DC33 struct {
	Cf60 _dafny.Sequence
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf60 _dafny.Sequence) D13 {
	return D13{D13_DC33{Cf60}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D13) Dtor_cf60() _dafny.Sequence {
	return _this.Get_().(D13_DC33).Cf60
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf60.Equals(data2.Cf60)
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

type D14_DC35 struct {
	Cf62 _dafny.Int
	Cf63 bool
	Cf64 _dafny.CodePoint
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf62 _dafny.Int, Cf63 bool, Cf64 _dafny.CodePoint) D14 {
	return D14{D14_DC35{Cf62, Cf63, Cf64}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC36 struct {
	Cf65 bool
	Cf66 _dafny.Sequence
	Cf67 _dafny.Sequence
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf65 bool, Cf66 _dafny.Sequence, Cf67 _dafny.Sequence) D14 {
	return D14{D14_DC36{Cf65, Cf66, Cf67}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC34 struct {
	Cf61 _dafny.Map
}

func (D14_DC34) isD14() {}

func (CompanionStruct_D14_) Create_DC34_(Cf61 _dafny.Map) D14 {
	return D14{D14_DC34{Cf61}}
}

func (_this D14) Is_DC34() bool {
	_, ok := _this.Get_().(D14_DC34)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC35_(_dafny.Zero, false, _dafny.CodePoint('D'))
}

func (_this D14) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D14_DC35).Cf62
}

func (_this D14) Dtor_cf63() bool {
	return _this.Get_().(D14_DC35).Cf63
}

func (_this D14) Dtor_cf64() _dafny.CodePoint {
	return _this.Get_().(D14_DC35).Cf64
}

func (_this D14) Dtor_cf65() bool {
	return _this.Get_().(D14_DC36).Cf65
}

func (_this D14) Dtor_cf66() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf66
}

func (_this D14) Dtor_cf67() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf67
}

func (_this D14) Dtor_cf61() _dafny.Map {
	return _this.Get_().(D14_DC34).Cf61
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf65) + ", " + data.Cf66.VerbatimString(true) + ", " + data.Cf67.VerbatimString(true) + ")"
		}
	case D14_DC34:
		{
			return "D14.DC34" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63 == data2.Cf63 && data1.Cf64 == data2.Cf64
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf65 == data2.Cf65 && data1.Cf66.Equals(data2.Cf66) && data1.Cf67.Equals(data2.Cf67)
		}
	case D14_DC34:
		{
			data2, ok := other.Get_().(D14_DC34)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D15_DC38 struct {
	Cf69 bool
	Cf70 _dafny.Sequence
	Cf71 bool
	Cf72 bool
	Cf73 _dafny.CodePoint
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf69 bool, Cf70 _dafny.Sequence, Cf71 bool, Cf72 bool, Cf73 _dafny.CodePoint) D15 {
	return D15{D15_DC38{Cf69, Cf70, Cf71, Cf72, Cf73}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC39 struct {
	Cf74 bool
	Cf75 bool
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf74 bool, Cf75 bool) D15 {
	return D15{D15_DC39{Cf74, Cf75}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

type D15_DC37 struct {
	Cf68 _dafny.Array
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf68 _dafny.Array) D15 {
	return D15{D15_DC37{Cf68}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC38_(false, _dafny.EmptySeq, false, false, _dafny.CodePoint('D'))
}

func (_this D15) Dtor_cf69() bool {
	return _this.Get_().(D15_DC38).Cf69
}

func (_this D15) Dtor_cf70() _dafny.Sequence {
	return _this.Get_().(D15_DC38).Cf70
}

func (_this D15) Dtor_cf71() bool {
	return _this.Get_().(D15_DC38).Cf71
}

func (_this D15) Dtor_cf72() bool {
	return _this.Get_().(D15_DC38).Cf72
}

func (_this D15) Dtor_cf73() _dafny.CodePoint {
	return _this.Get_().(D15_DC38).Cf73
}

func (_this D15) Dtor_cf74() bool {
	return _this.Get_().(D15_DC39).Cf74
}

func (_this D15) Dtor_cf75() bool {
	return _this.Get_().(D15_DC39).Cf75
}

func (_this D15) Dtor_cf68() _dafny.Array {
	return _this.Get_().(D15_DC37).Cf68
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf69) + ", " + data.Cf70.VerbatimString(true) + ", " + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ")"
		}
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf69 == data2.Cf69 && data1.Cf70.Equals(data2.Cf70) && data1.Cf71 == data2.Cf71 && data1.Cf72 == data2.Cf72 && data1.Cf73 == data2.Cf73
		}
	case D15_DC39:
		{
			data2, ok := other.Get_().(D15_DC39)
			return ok && data1.Cf74 == data2.Cf74 && data1.Cf75 == data2.Cf75
		}
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf68 == data2.Cf68
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

type D16_DC41 struct {
	Cf77 bool
	Cf78 _dafny.Int
	Cf79 bool
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf77 bool, Cf78 _dafny.Int, Cf79 bool) D16 {
	return D16{D16_DC41{Cf77, Cf78, Cf79}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

type D16_DC42 struct {
	Cf80 bool
	Cf81 _dafny.Int
	Cf82 _dafny.Int
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf80 bool, Cf81 _dafny.Int, Cf82 _dafny.Int) D16 {
	return D16{D16_DC42{Cf80, Cf81, Cf82}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC40 struct {
	Cf76 _dafny.Sequence
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf76 _dafny.Sequence) D16 {
	return D16{D16_DC40{Cf76}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC41_(false, _dafny.Zero, false)
}

func (_this D16) Dtor_cf77() bool {
	return _this.Get_().(D16_DC41).Cf77
}

func (_this D16) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D16_DC41).Cf78
}

func (_this D16) Dtor_cf79() bool {
	return _this.Get_().(D16_DC41).Cf79
}

func (_this D16) Dtor_cf80() bool {
	return _this.Get_().(D16_DC42).Cf80
}

func (_this D16) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf81
}

func (_this D16) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf82
}

func (_this D16) Dtor_cf76() _dafny.Sequence {
	return _this.Get_().(D16_DC40).Cf76
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf76) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf77 == data2.Cf77 && data1.Cf78.Cmp(data2.Cf78) == 0 && data1.Cf79 == data2.Cf79
		}
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf80 == data2.Cf80 && data1.Cf81.Cmp(data2.Cf81) == 0 && data1.Cf82.Cmp(data2.Cf82) == 0
		}
	case D16_DC40:
		{
			data2, ok := other.Get_().(D16_DC40)
			return ok && data1.Cf76.Equals(data2.Cf76)
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

type D17_DC44 struct {
	Cf84 _dafny.Int
	Cf85 _dafny.Sequence
	Cf86 _dafny.CodePoint
	Cf87 bool
	Cf88 _dafny.Int
}

func (D17_DC44) isD17() {}

func (CompanionStruct_D17_) Create_DC44_(Cf84 _dafny.Int, Cf85 _dafny.Sequence, Cf86 _dafny.CodePoint, Cf87 bool, Cf88 _dafny.Int) D17 {
	return D17{D17_DC44{Cf84, Cf85, Cf86, Cf87, Cf88}}
}

func (_this D17) Is_DC44() bool {
	_, ok := _this.Get_().(D17_DC44)
	return ok
}

type D17_DC45 struct {
}

func (D17_DC45) isD17() {}

func (CompanionStruct_D17_) Create_DC45_() D17 {
	return D17{D17_DC45{}}
}

func (_this D17) Is_DC45() bool {
	_, ok := _this.Get_().(D17_DC45)
	return ok
}

type D17_DC43 struct {
	Cf83 _dafny.Array
}

func (D17_DC43) isD17() {}

func (CompanionStruct_D17_) Create_DC43_(Cf83 _dafny.Array) D17 {
	return D17{D17_DC43{Cf83}}
}

func (_this D17) Is_DC43() bool {
	_, ok := _this.Get_().(D17_DC43)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC44_(_dafny.Zero, _dafny.EmptySeq, _dafny.CodePoint('D'), false, _dafny.Zero)
}

func (_this D17) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D17_DC44).Cf84
}

func (_this D17) Dtor_cf85() _dafny.Sequence {
	return _this.Get_().(D17_DC44).Cf85
}

func (_this D17) Dtor_cf86() _dafny.CodePoint {
	return _this.Get_().(D17_DC44).Cf86
}

func (_this D17) Dtor_cf87() bool {
	return _this.Get_().(D17_DC44).Cf87
}

func (_this D17) Dtor_cf88() _dafny.Int {
	return _this.Get_().(D17_DC44).Cf88
}

func (_this D17) Dtor_cf83() _dafny.Array {
	return _this.Get_().(D17_DC43).Cf83
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC44:
		{
			return "D17.DC44" + "(" + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ")"
		}
	case D17_DC45:
		{
			return "D17.DC45"
		}
	case D17_DC43:
		{
			return "D17.DC43" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC44:
		{
			data2, ok := other.Get_().(D17_DC44)
			return ok && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85.Equals(data2.Cf85) && data1.Cf86 == data2.Cf86 && data1.Cf87 == data2.Cf87 && data1.Cf88.Cmp(data2.Cf88) == 0
		}
	case D17_DC45:
		{
			_, ok := other.Get_().(D17_DC45)
			return ok
		}
	case D17_DC43:
		{
			data2, ok := other.Get_().(D17_DC43)
			return ok && data1.Cf83 == data2.Cf83
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

type D18_DC47 struct {
	Cf90 bool
	Cf91 D6
}

func (D18_DC47) isD18() {}

func (CompanionStruct_D18_) Create_DC47_(Cf90 bool, Cf91 D6) D18 {
	return D18{D18_DC47{Cf90, Cf91}}
}

func (_this D18) Is_DC47() bool {
	_, ok := _this.Get_().(D18_DC47)
	return ok
}

type D18_DC46 struct {
	Cf89 _dafny.Map
}

func (D18_DC46) isD18() {}

func (CompanionStruct_D18_) Create_DC46_(Cf89 _dafny.Map) D18 {
	return D18{D18_DC46{Cf89}}
}

func (_this D18) Is_DC46() bool {
	_, ok := _this.Get_().(D18_DC46)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC47_(false, Companion_D6_.Default())
}

func (_this D18) Dtor_cf90() bool {
	return _this.Get_().(D18_DC47).Cf90
}

func (_this D18) Dtor_cf91() D6 {
	return _this.Get_().(D18_DC47).Cf91
}

func (_this D18) Dtor_cf89() _dafny.Map {
	return _this.Get_().(D18_DC46).Cf89
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC47:
		{
			return "D18.DC47" + "(" + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ")"
		}
	case D18_DC46:
		{
			return "D18.DC46" + "(" + _dafny.String(data.Cf89) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC47:
		{
			data2, ok := other.Get_().(D18_DC47)
			return ok && data1.Cf90 == data2.Cf90 && data1.Cf91.Equals(data2.Cf91)
		}
	case D18_DC46:
		{
			data2, ok := other.Get_().(D18_DC46)
			return ok && data1.Cf89.Equals(data2.Cf89)
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

type D19_DC49 struct {
	Cf93 bool
}

func (D19_DC49) isD19() {}

func (CompanionStruct_D19_) Create_DC49_(Cf93 bool) D19 {
	return D19{D19_DC49{Cf93}}
}

func (_this D19) Is_DC49() bool {
	_, ok := _this.Get_().(D19_DC49)
	return ok
}

type D19_DC50 struct {
	Cf94 bool
	Cf95 _dafny.Int
	Cf96 _dafny.Array
}

func (D19_DC50) isD19() {}

func (CompanionStruct_D19_) Create_DC50_(Cf94 bool, Cf95 _dafny.Int, Cf96 _dafny.Array) D19 {
	return D19{D19_DC50{Cf94, Cf95, Cf96}}
}

func (_this D19) Is_DC50() bool {
	_, ok := _this.Get_().(D19_DC50)
	return ok
}

type D19_DC48 struct {
	Cf92 _dafny.Array
}

func (D19_DC48) isD19() {}

func (CompanionStruct_D19_) Create_DC48_(Cf92 _dafny.Array) D19 {
	return D19{D19_DC48{Cf92}}
}

func (_this D19) Is_DC48() bool {
	_, ok := _this.Get_().(D19_DC48)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC49_(false)
}

func (_this D19) Dtor_cf93() bool {
	return _this.Get_().(D19_DC49).Cf93
}

func (_this D19) Dtor_cf94() bool {
	return _this.Get_().(D19_DC50).Cf94
}

func (_this D19) Dtor_cf95() _dafny.Int {
	return _this.Get_().(D19_DC50).Cf95
}

func (_this D19) Dtor_cf96() _dafny.Array {
	return _this.Get_().(D19_DC50).Cf96
}

func (_this D19) Dtor_cf92() _dafny.Array {
	return _this.Get_().(D19_DC48).Cf92
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC49:
		{
			return "D19.DC49" + "(" + _dafny.String(data.Cf93) + ")"
		}
	case D19_DC50:
		{
			return "D19.DC50" + "(" + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ")"
		}
	case D19_DC48:
		{
			return "D19.DC48" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC49:
		{
			data2, ok := other.Get_().(D19_DC49)
			return ok && data1.Cf93 == data2.Cf93
		}
	case D19_DC50:
		{
			data2, ok := other.Get_().(D19_DC50)
			return ok && data1.Cf94 == data2.Cf94 && data1.Cf95.Cmp(data2.Cf95) == 0 && data1.Cf96 == data2.Cf96
		}
	case D19_DC48:
		{
			data2, ok := other.Get_().(D19_DC48)
			return ok && data1.Cf92 == data2.Cf92
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

type D20_DC52 struct {
	Cf98 _dafny.Int
}

func (D20_DC52) isD20() {}

func (CompanionStruct_D20_) Create_DC52_(Cf98 _dafny.Int) D20 {
	return D20{D20_DC52{Cf98}}
}

func (_this D20) Is_DC52() bool {
	_, ok := _this.Get_().(D20_DC52)
	return ok
}

type D20_DC51 struct {
	Cf97 _dafny.Array
}

func (D20_DC51) isD20() {}

func (CompanionStruct_D20_) Create_DC51_(Cf97 _dafny.Array) D20 {
	return D20{D20_DC51{Cf97}}
}

func (_this D20) Is_DC51() bool {
	_, ok := _this.Get_().(D20_DC51)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC52_(_dafny.Zero)
}

func (_this D20) Dtor_cf98() _dafny.Int {
	return _this.Get_().(D20_DC52).Cf98
}

func (_this D20) Dtor_cf97() _dafny.Array {
	return _this.Get_().(D20_DC51).Cf97
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC52:
		{
			return "D20.DC52" + "(" + _dafny.String(data.Cf98) + ")"
		}
	case D20_DC51:
		{
			return "D20.DC51" + "(" + _dafny.String(data.Cf97) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC52:
		{
			data2, ok := other.Get_().(D20_DC52)
			return ok && data1.Cf98.Cmp(data2.Cf98) == 0
		}
	case D20_DC51:
		{
			data2, ok := other.Get_().(D20_DC51)
			return ok && data1.Cf97 == data2.Cf97
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

type D21_DC54 struct {
	Cf100 _dafny.CodePoint
	Cf101 _dafny.Int
	Cf102 _dafny.Sequence
}

func (D21_DC54) isD21() {}

func (CompanionStruct_D21_) Create_DC54_(Cf100 _dafny.CodePoint, Cf101 _dafny.Int, Cf102 _dafny.Sequence) D21 {
	return D21{D21_DC54{Cf100, Cf101, Cf102}}
}

func (_this D21) Is_DC54() bool {
	_, ok := _this.Get_().(D21_DC54)
	return ok
}

type D21_DC53 struct {
	Cf99 T0
}

func (D21_DC53) isD21() {}

func (CompanionStruct_D21_) Create_DC53_(Cf99 T0) D21 {
	return D21{D21_DC53{Cf99}}
}

func (_this D21) Is_DC53() bool {
	_, ok := _this.Get_().(D21_DC53)
	return ok
}

type D21_DC55 struct {
	Cf103 D21
}

func (D21_DC55) isD21() {}

func (CompanionStruct_D21_) Create_DC55_(Cf103 D21) D21 {
	return D21{D21_DC55{Cf103}}
}

func (_this D21) Is_DC55() bool {
	_, ok := _this.Get_().(D21_DC55)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC54_(_dafny.CodePoint('D'), _dafny.Zero, _dafny.EmptySeq)
}

func (_this D21) Dtor_cf100() _dafny.CodePoint {
	return _this.Get_().(D21_DC54).Cf100
}

func (_this D21) Dtor_cf101() _dafny.Int {
	return _this.Get_().(D21_DC54).Cf101
}

func (_this D21) Dtor_cf102() _dafny.Sequence {
	return _this.Get_().(D21_DC54).Cf102
}

func (_this D21) Dtor_cf99() T0 {
	return _this.Get_().(D21_DC53).Cf99
}

func (_this D21) Dtor_cf103() D21 {
	return _this.Get_().(D21_DC55).Cf103
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC54:
		{
			return "D21.DC54" + "(" + _dafny.String(data.Cf100) + ", " + _dafny.String(data.Cf101) + ", " + data.Cf102.VerbatimString(true) + ")"
		}
	case D21_DC53:
		{
			return "D21.DC53" + "(" + _dafny.String(data.Cf99) + ")"
		}
	case D21_DC55:
		{
			return "D21.DC55" + "(" + _dafny.String(data.Cf103) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC54:
		{
			data2, ok := other.Get_().(D21_DC54)
			return ok && data1.Cf100 == data2.Cf100 && data1.Cf101.Cmp(data2.Cf101) == 0 && data1.Cf102.Equals(data2.Cf102)
		}
	case D21_DC53:
		{
			data2, ok := other.Get_().(D21_DC53)
			return ok && _dafny.AreEqual(data1.Cf99, data2.Cf99)
		}
	case D21_DC55:
		{
			data2, ok := other.Get_().(D21_DC55)
			return ok && data1.Cf103.Equals(data2.Cf103)
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
}

func (D22_DC57) isD22() {}

func (CompanionStruct_D22_) Create_DC57_() D22 {
	return D22{D22_DC57{}}
}

func (_this D22) Is_DC57() bool {
	_, ok := _this.Get_().(D22_DC57)
	return ok
}

type D22_DC56 struct {
	Cf104 T1
}

func (D22_DC56) isD22() {}

func (CompanionStruct_D22_) Create_DC56_(Cf104 T1) D22 {
	return D22{D22_DC56{Cf104}}
}

func (_this D22) Is_DC56() bool {
	_, ok := _this.Get_().(D22_DC56)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC57_()
}

func (_this D22) Dtor_cf104() T1 {
	return _this.Get_().(D22_DC56).Cf104
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC57:
		{
			return "D22.DC57"
		}
	case D22_DC56:
		{
			return "D22.DC56" + "(" + _dafny.String(data.Cf104) + ")"
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
			_, ok := other.Get_().(D22_DC57)
			return ok
		}
	case D22_DC56:
		{
			data2, ok := other.Get_().(D22_DC56)
			return ok && _dafny.AreEqual(data1.Cf104, data2.Cf104)
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

type D23_DC59 struct {
	Cf106 D6
	Cf107 bool
	Cf108 _dafny.Int
	Cf109 _dafny.Array
}

func (D23_DC59) isD23() {}

func (CompanionStruct_D23_) Create_DC59_(Cf106 D6, Cf107 bool, Cf108 _dafny.Int, Cf109 _dafny.Array) D23 {
	return D23{D23_DC59{Cf106, Cf107, Cf108, Cf109}}
}

func (_this D23) Is_DC59() bool {
	_, ok := _this.Get_().(D23_DC59)
	return ok
}

type D23_DC58 struct {
	Cf105 _dafny.Array
}

func (D23_DC58) isD23() {}

func (CompanionStruct_D23_) Create_DC58_(Cf105 _dafny.Array) D23 {
	return D23{D23_DC58{Cf105}}
}

func (_this D23) Is_DC58() bool {
	_, ok := _this.Get_().(D23_DC58)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC59_(Companion_D6_.Default(), false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D23) Dtor_cf106() D6 {
	return _this.Get_().(D23_DC59).Cf106
}

func (_this D23) Dtor_cf107() bool {
	return _this.Get_().(D23_DC59).Cf107
}

func (_this D23) Dtor_cf108() _dafny.Int {
	return _this.Get_().(D23_DC59).Cf108
}

func (_this D23) Dtor_cf109() _dafny.Array {
	return _this.Get_().(D23_DC59).Cf109
}

func (_this D23) Dtor_cf105() _dafny.Array {
	return _this.Get_().(D23_DC58).Cf105
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC59:
		{
			return "D23.DC59" + "(" + _dafny.String(data.Cf106) + ", " + _dafny.String(data.Cf107) + ", " + _dafny.String(data.Cf108) + ", " + _dafny.String(data.Cf109) + ")"
		}
	case D23_DC58:
		{
			return "D23.DC58" + "(" + _dafny.String(data.Cf105) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC59:
		{
			data2, ok := other.Get_().(D23_DC59)
			return ok && data1.Cf106.Equals(data2.Cf106) && data1.Cf107 == data2.Cf107 && data1.Cf108.Cmp(data2.Cf108) == 0 && data1.Cf109 == data2.Cf109
		}
	case D23_DC58:
		{
			data2, ok := other.Get_().(D23_DC58)
			return ok && data1.Cf105 == data2.Cf105
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

type D24_DC60 struct {
	Cf110 _dafny.Map
}

func (D24_DC60) isD24() {}

func (CompanionStruct_D24_) Create_DC60_(Cf110 _dafny.Map) D24 {
	return D24{D24_DC60{Cf110}}
}

func (_this D24) Is_DC60() bool {
	_, ok := _this.Get_().(D24_DC60)
	return ok
}

func (CompanionStruct_D24_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D24) Dtor_cf110() _dafny.Map {
	return _this.Get_().(D24_DC60).Cf110
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC60:
		{
			return "D24.DC60" + "(" + _dafny.String(data.Cf110) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC60:
		{
			data2, ok := other.Get_().(D24_DC60)
			return ok && data1.Cf110.Equals(data2.Cf110)
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

type D25_DC62 struct {
}

func (D25_DC62) isD25() {}

func (CompanionStruct_D25_) Create_DC62_() D25 {
	return D25{D25_DC62{}}
}

func (_this D25) Is_DC62() bool {
	_, ok := _this.Get_().(D25_DC62)
	return ok
}

type D25_DC63 struct {
	Cf112 _dafny.Sequence
	Cf113 _dafny.Int
}

func (D25_DC63) isD25() {}

func (CompanionStruct_D25_) Create_DC63_(Cf112 _dafny.Sequence, Cf113 _dafny.Int) D25 {
	return D25{D25_DC63{Cf112, Cf113}}
}

func (_this D25) Is_DC63() bool {
	_, ok := _this.Get_().(D25_DC63)
	return ok
}

type D25_DC61 struct {
	Cf111 _dafny.Array
}

func (D25_DC61) isD25() {}

func (CompanionStruct_D25_) Create_DC61_(Cf111 _dafny.Array) D25 {
	return D25{D25_DC61{Cf111}}
}

func (_this D25) Is_DC61() bool {
	_, ok := _this.Get_().(D25_DC61)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC62_()
}

func (_this D25) Dtor_cf112() _dafny.Sequence {
	return _this.Get_().(D25_DC63).Cf112
}

func (_this D25) Dtor_cf113() _dafny.Int {
	return _this.Get_().(D25_DC63).Cf113
}

func (_this D25) Dtor_cf111() _dafny.Array {
	return _this.Get_().(D25_DC61).Cf111
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC62:
		{
			return "D25.DC62"
		}
	case D25_DC63:
		{
			return "D25.DC63" + "(" + data.Cf112.VerbatimString(true) + ", " + _dafny.String(data.Cf113) + ")"
		}
	case D25_DC61:
		{
			return "D25.DC61" + "(" + _dafny.String(data.Cf111) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC62:
		{
			_, ok := other.Get_().(D25_DC62)
			return ok
		}
	case D25_DC63:
		{
			data2, ok := other.Get_().(D25_DC63)
			return ok && data1.Cf112.Equals(data2.Cf112) && data1.Cf113.Cmp(data2.Cf113) == 0
		}
	case D25_DC61:
		{
			data2, ok := other.Get_().(D25_DC61)
			return ok && data1.Cf111 == data2.Cf111
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

type D26_DC65 struct {
	Cf115 bool
	Cf116 _dafny.Array
	Cf117 bool
}

func (D26_DC65) isD26() {}

func (CompanionStruct_D26_) Create_DC65_(Cf115 bool, Cf116 _dafny.Array, Cf117 bool) D26 {
	return D26{D26_DC65{Cf115, Cf116, Cf117}}
}

func (_this D26) Is_DC65() bool {
	_, ok := _this.Get_().(D26_DC65)
	return ok
}

type D26_DC64 struct {
	Cf114 *C1
}

func (D26_DC64) isD26() {}

func (CompanionStruct_D26_) Create_DC64_(Cf114 *C1) D26 {
	return D26{D26_DC64{Cf114}}
}

func (_this D26) Is_DC64() bool {
	_, ok := _this.Get_().(D26_DC64)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC65_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D26) Dtor_cf115() bool {
	return _this.Get_().(D26_DC65).Cf115
}

func (_this D26) Dtor_cf116() _dafny.Array {
	return _this.Get_().(D26_DC65).Cf116
}

func (_this D26) Dtor_cf117() bool {
	return _this.Get_().(D26_DC65).Cf117
}

func (_this D26) Dtor_cf114() *C1 {
	return _this.Get_().(D26_DC64).Cf114
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC65:
		{
			return "D26.DC65" + "(" + _dafny.String(data.Cf115) + ", " + _dafny.String(data.Cf116) + ", " + _dafny.String(data.Cf117) + ")"
		}
	case D26_DC64:
		{
			return "D26.DC64" + "(" + _dafny.String(data.Cf114) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC65:
		{
			data2, ok := other.Get_().(D26_DC65)
			return ok && data1.Cf115 == data2.Cf115 && data1.Cf116 == data2.Cf116 && data1.Cf117 == data2.Cf117
		}
	case D26_DC64:
		{
			data2, ok := other.Get_().(D26_DC64)
			return ok && data1.Cf114 == data2.Cf114
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

// Definition of datatype D27
type D27 struct {
	Data_D27_
}

func (_this D27) Get_() Data_D27_ {
	return _this.Data_D27_
}

type Data_D27_ interface {
	isD27()
}

type CompanionStruct_D27_ struct {
}

var Companion_D27_ = CompanionStruct_D27_{}

type D27_DC66 struct {
	Cf118 _dafny.MultiSet
}

func (D27_DC66) isD27() {}

func (CompanionStruct_D27_) Create_DC66_(Cf118 _dafny.MultiSet) D27 {
	return D27{D27_DC66{Cf118}}
}

func (_this D27) Is_DC66() bool {
	_, ok := _this.Get_().(D27_DC66)
	return ok
}

func (CompanionStruct_D27_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D27) Dtor_cf118() _dafny.MultiSet {
	return _this.Get_().(D27_DC66).Cf118
}

func (_this D27) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D27_DC66:
		{
			return "D27.DC66" + "(" + _dafny.String(data.Cf118) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D27) Equals(other D27) bool {
	switch data1 := _this.Get_().(type) {
	case D27_DC66:
		{
			data2, ok := other.Get_().(D27_DC66)
			return ok && data1.Cf118.Equals(data2.Cf118)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D27) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D27)
	return ok && _this.Equals(typed)
}

func Type_D27_() _dafny.TypeDescriptor {
	return type_D27_{}
}

type type_D27_ struct {
}

func (_this type_D27_) Default() interface{} {
	return Companion_D27_.Default()
}

func (_this type_D27_) String() string {
	return "main.D27"
}
func (_this D27) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D27{}

// End of datatype D27

// Definition of datatype D28
type D28 struct {
	Data_D28_
}

func (_this D28) Get_() Data_D28_ {
	return _this.Data_D28_
}

type Data_D28_ interface {
	isD28()
}

type CompanionStruct_D28_ struct {
}

var Companion_D28_ = CompanionStruct_D28_{}

type D28_DC68 struct {
	Cf120 _dafny.Int
	Cf121 bool
	Cf122 _dafny.Int
}

func (D28_DC68) isD28() {}

func (CompanionStruct_D28_) Create_DC68_(Cf120 _dafny.Int, Cf121 bool, Cf122 _dafny.Int) D28 {
	return D28{D28_DC68{Cf120, Cf121, Cf122}}
}

func (_this D28) Is_DC68() bool {
	_, ok := _this.Get_().(D28_DC68)
	return ok
}

type D28_DC69 struct {
}

func (D28_DC69) isD28() {}

func (CompanionStruct_D28_) Create_DC69_() D28 {
	return D28{D28_DC69{}}
}

func (_this D28) Is_DC69() bool {
	_, ok := _this.Get_().(D28_DC69)
	return ok
}

type D28_DC70 struct {
	Cf123 bool
	Cf124 bool
}

func (D28_DC70) isD28() {}

func (CompanionStruct_D28_) Create_DC70_(Cf123 bool, Cf124 bool) D28 {
	return D28{D28_DC70{Cf123, Cf124}}
}

func (_this D28) Is_DC70() bool {
	_, ok := _this.Get_().(D28_DC70)
	return ok
}

type D28_DC67 struct {
	Cf119 _dafny.Sequence
}

func (D28_DC67) isD28() {}

func (CompanionStruct_D28_) Create_DC67_(Cf119 _dafny.Sequence) D28 {
	return D28{D28_DC67{Cf119}}
}

func (_this D28) Is_DC67() bool {
	_, ok := _this.Get_().(D28_DC67)
	return ok
}

type D28_DC71 struct {
	Cf125 D28
}

func (D28_DC71) isD28() {}

func (CompanionStruct_D28_) Create_DC71_(Cf125 D28) D28 {
	return D28{D28_DC71{Cf125}}
}

func (_this D28) Is_DC71() bool {
	_, ok := _this.Get_().(D28_DC71)
	return ok
}

func (CompanionStruct_D28_) Default() D28 {
	return Companion_D28_.Create_DC68_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D28) Dtor_cf120() _dafny.Int {
	return _this.Get_().(D28_DC68).Cf120
}

func (_this D28) Dtor_cf121() bool {
	return _this.Get_().(D28_DC68).Cf121
}

func (_this D28) Dtor_cf122() _dafny.Int {
	return _this.Get_().(D28_DC68).Cf122
}

func (_this D28) Dtor_cf123() bool {
	return _this.Get_().(D28_DC70).Cf123
}

func (_this D28) Dtor_cf124() bool {
	return _this.Get_().(D28_DC70).Cf124
}

func (_this D28) Dtor_cf119() _dafny.Sequence {
	return _this.Get_().(D28_DC67).Cf119
}

func (_this D28) Dtor_cf125() D28 {
	return _this.Get_().(D28_DC71).Cf125
}

func (_this D28) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D28_DC68:
		{
			return "D28.DC68" + "(" + _dafny.String(data.Cf120) + ", " + _dafny.String(data.Cf121) + ", " + _dafny.String(data.Cf122) + ")"
		}
	case D28_DC69:
		{
			return "D28.DC69"
		}
	case D28_DC70:
		{
			return "D28.DC70" + "(" + _dafny.String(data.Cf123) + ", " + _dafny.String(data.Cf124) + ")"
		}
	case D28_DC67:
		{
			return "D28.DC67" + "(" + _dafny.String(data.Cf119) + ")"
		}
	case D28_DC71:
		{
			return "D28.DC71" + "(" + _dafny.String(data.Cf125) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D28) Equals(other D28) bool {
	switch data1 := _this.Get_().(type) {
	case D28_DC68:
		{
			data2, ok := other.Get_().(D28_DC68)
			return ok && data1.Cf120.Cmp(data2.Cf120) == 0 && data1.Cf121 == data2.Cf121 && data1.Cf122.Cmp(data2.Cf122) == 0
		}
	case D28_DC69:
		{
			_, ok := other.Get_().(D28_DC69)
			return ok
		}
	case D28_DC70:
		{
			data2, ok := other.Get_().(D28_DC70)
			return ok && data1.Cf123 == data2.Cf123 && data1.Cf124 == data2.Cf124
		}
	case D28_DC67:
		{
			data2, ok := other.Get_().(D28_DC67)
			return ok && data1.Cf119.Equals(data2.Cf119)
		}
	case D28_DC71:
		{
			data2, ok := other.Get_().(D28_DC71)
			return ok && data1.Cf125.Equals(data2.Cf125)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D28) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D28)
	return ok && _this.Equals(typed)
}

func Type_D28_() _dafny.TypeDescriptor {
	return type_D28_{}
}

type type_D28_ struct {
}

func (_this type_D28_) Default() interface{} {
	return Companion_D28_.Default()
}

func (_this type_D28_) String() string {
	return "main.D28"
}
func (_this D28) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D28{}

// End of datatype D28

// Definition of datatype D29
type D29 struct {
	Data_D29_
}

func (_this D29) Get_() Data_D29_ {
	return _this.Data_D29_
}

type Data_D29_ interface {
	isD29()
}

type CompanionStruct_D29_ struct {
}

var Companion_D29_ = CompanionStruct_D29_{}

type D29_DC72 struct {
	Cf126 _dafny.Map
}

func (D29_DC72) isD29() {}

func (CompanionStruct_D29_) Create_DC72_(Cf126 _dafny.Map) D29 {
	return D29{D29_DC72{Cf126}}
}

func (_this D29) Is_DC72() bool {
	_, ok := _this.Get_().(D29_DC72)
	return ok
}

func (CompanionStruct_D29_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D29) Dtor_cf126() _dafny.Map {
	return _this.Get_().(D29_DC72).Cf126
}

func (_this D29) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D29_DC72:
		{
			return "D29.DC72" + "(" + _dafny.String(data.Cf126) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D29) Equals(other D29) bool {
	switch data1 := _this.Get_().(type) {
	case D29_DC72:
		{
			data2, ok := other.Get_().(D29_DC72)
			return ok && data1.Cf126.Equals(data2.Cf126)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D29) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D29)
	return ok && _this.Equals(typed)
}

func Type_D29_() _dafny.TypeDescriptor {
	return type_D29_{}
}

type type_D29_ struct {
}

func (_this type_D29_) Default() interface{} {
	return Companion_D29_.Default()
}

func (_this type_D29_) String() string {
	return "main.D29"
}
func (_this D29) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D29{}

// End of datatype D29

// Definition of datatype D30
type D30 struct {
	Data_D30_
}

func (_this D30) Get_() Data_D30_ {
	return _this.Data_D30_
}

type Data_D30_ interface {
	isD30()
}

type CompanionStruct_D30_ struct {
}

var Companion_D30_ = CompanionStruct_D30_{}

type D30_DC74 struct {
	Cf128 _dafny.Sequence
}

func (D30_DC74) isD30() {}

func (CompanionStruct_D30_) Create_DC74_(Cf128 _dafny.Sequence) D30 {
	return D30{D30_DC74{Cf128}}
}

func (_this D30) Is_DC74() bool {
	_, ok := _this.Get_().(D30_DC74)
	return ok
}

type D30_DC75 struct {
	Cf129 bool
}

func (D30_DC75) isD30() {}

func (CompanionStruct_D30_) Create_DC75_(Cf129 bool) D30 {
	return D30{D30_DC75{Cf129}}
}

func (_this D30) Is_DC75() bool {
	_, ok := _this.Get_().(D30_DC75)
	return ok
}

type D30_DC73 struct {
	Cf127 *C4
}

func (D30_DC73) isD30() {}

func (CompanionStruct_D30_) Create_DC73_(Cf127 *C4) D30 {
	return D30{D30_DC73{Cf127}}
}

func (_this D30) Is_DC73() bool {
	_, ok := _this.Get_().(D30_DC73)
	return ok
}

type D30_DC76 struct {
	Cf130 D30
}

func (D30_DC76) isD30() {}

func (CompanionStruct_D30_) Create_DC76_(Cf130 D30) D30 {
	return D30{D30_DC76{Cf130}}
}

func (_this D30) Is_DC76() bool {
	_, ok := _this.Get_().(D30_DC76)
	return ok
}

func (CompanionStruct_D30_) Default() D30 {
	return Companion_D30_.Create_DC74_(_dafny.EmptySeq)
}

func (_this D30) Dtor_cf128() _dafny.Sequence {
	return _this.Get_().(D30_DC74).Cf128
}

func (_this D30) Dtor_cf129() bool {
	return _this.Get_().(D30_DC75).Cf129
}

func (_this D30) Dtor_cf127() *C4 {
	return _this.Get_().(D30_DC73).Cf127
}

func (_this D30) Dtor_cf130() D30 {
	return _this.Get_().(D30_DC76).Cf130
}

func (_this D30) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D30_DC74:
		{
			return "D30.DC74" + "(" + data.Cf128.VerbatimString(true) + ")"
		}
	case D30_DC75:
		{
			return "D30.DC75" + "(" + _dafny.String(data.Cf129) + ")"
		}
	case D30_DC73:
		{
			return "D30.DC73" + "(" + _dafny.String(data.Cf127) + ")"
		}
	case D30_DC76:
		{
			return "D30.DC76" + "(" + _dafny.String(data.Cf130) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D30) Equals(other D30) bool {
	switch data1 := _this.Get_().(type) {
	case D30_DC74:
		{
			data2, ok := other.Get_().(D30_DC74)
			return ok && data1.Cf128.Equals(data2.Cf128)
		}
	case D30_DC75:
		{
			data2, ok := other.Get_().(D30_DC75)
			return ok && data1.Cf129 == data2.Cf129
		}
	case D30_DC73:
		{
			data2, ok := other.Get_().(D30_DC73)
			return ok && data1.Cf127 == data2.Cf127
		}
	case D30_DC76:
		{
			data2, ok := other.Get_().(D30_DC76)
			return ok && data1.Cf130.Equals(data2.Cf130)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D30) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D30)
	return ok && _this.Equals(typed)
}

func Type_D30_() _dafny.TypeDescriptor {
	return type_D30_{}
}

type type_D30_ struct {
}

func (_this type_D30_) Default() interface{} {
	return Companion_D30_.Default()
}

func (_this type_D30_) String() string {
	return "main.D30"
}
func (_this D30) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D30{}

// End of datatype D30

// Definition of datatype D31
type D31 struct {
	Data_D31_
}

func (_this D31) Get_() Data_D31_ {
	return _this.Data_D31_
}

type Data_D31_ interface {
	isD31()
}

type CompanionStruct_D31_ struct {
}

var Companion_D31_ = CompanionStruct_D31_{}

type D31_DC77 struct {
	Cf131 _dafny.MultiSet
}

func (D31_DC77) isD31() {}

func (CompanionStruct_D31_) Create_DC77_(Cf131 _dafny.MultiSet) D31 {
	return D31{D31_DC77{Cf131}}
}

func (_this D31) Is_DC77() bool {
	_, ok := _this.Get_().(D31_DC77)
	return ok
}

func (CompanionStruct_D31_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D31) Dtor_cf131() _dafny.MultiSet {
	return _this.Get_().(D31_DC77).Cf131
}

func (_this D31) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D31_DC77:
		{
			return "D31.DC77" + "(" + _dafny.String(data.Cf131) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D31) Equals(other D31) bool {
	switch data1 := _this.Get_().(type) {
	case D31_DC77:
		{
			data2, ok := other.Get_().(D31_DC77)
			return ok && data1.Cf131.Equals(data2.Cf131)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D31) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D31)
	return ok && _this.Equals(typed)
}

func Type_D31_() _dafny.TypeDescriptor {
	return type_D31_{}
}

type type_D31_ struct {
}

func (_this type_D31_) Default() interface{} {
	return Companion_D31_.Default()
}

func (_this type_D31_) String() string {
	return "main.D31"
}
func (_this D31) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D31{}

// End of datatype D31

// Definition of datatype D32
type D32 struct {
	Data_D32_
}

func (_this D32) Get_() Data_D32_ {
	return _this.Data_D32_
}

type Data_D32_ interface {
	isD32()
}

type CompanionStruct_D32_ struct {
}

var Companion_D32_ = CompanionStruct_D32_{}

type D32_DC79 struct {
	Cf133 bool
	Cf134 _dafny.Sequence
	Cf135 _dafny.Sequence
}

func (D32_DC79) isD32() {}

func (CompanionStruct_D32_) Create_DC79_(Cf133 bool, Cf134 _dafny.Sequence, Cf135 _dafny.Sequence) D32 {
	return D32{D32_DC79{Cf133, Cf134, Cf135}}
}

func (_this D32) Is_DC79() bool {
	_, ok := _this.Get_().(D32_DC79)
	return ok
}

type D32_DC80 struct {
	Cf136 _dafny.Int
	Cf137 _dafny.Int
	Cf138 _dafny.Int
}

func (D32_DC80) isD32() {}

func (CompanionStruct_D32_) Create_DC80_(Cf136 _dafny.Int, Cf137 _dafny.Int, Cf138 _dafny.Int) D32 {
	return D32{D32_DC80{Cf136, Cf137, Cf138}}
}

func (_this D32) Is_DC80() bool {
	_, ok := _this.Get_().(D32_DC80)
	return ok
}

type D32_DC78 struct {
	Cf132 _dafny.Array
}

func (D32_DC78) isD32() {}

func (CompanionStruct_D32_) Create_DC78_(Cf132 _dafny.Array) D32 {
	return D32{D32_DC78{Cf132}}
}

func (_this D32) Is_DC78() bool {
	_, ok := _this.Get_().(D32_DC78)
	return ok
}

func (CompanionStruct_D32_) Default() D32 {
	return Companion_D32_.Create_DC79_(false, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D32) Dtor_cf133() bool {
	return _this.Get_().(D32_DC79).Cf133
}

func (_this D32) Dtor_cf134() _dafny.Sequence {
	return _this.Get_().(D32_DC79).Cf134
}

func (_this D32) Dtor_cf135() _dafny.Sequence {
	return _this.Get_().(D32_DC79).Cf135
}

func (_this D32) Dtor_cf136() _dafny.Int {
	return _this.Get_().(D32_DC80).Cf136
}

func (_this D32) Dtor_cf137() _dafny.Int {
	return _this.Get_().(D32_DC80).Cf137
}

func (_this D32) Dtor_cf138() _dafny.Int {
	return _this.Get_().(D32_DC80).Cf138
}

func (_this D32) Dtor_cf132() _dafny.Array {
	return _this.Get_().(D32_DC78).Cf132
}

func (_this D32) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D32_DC79:
		{
			return "D32.DC79" + "(" + _dafny.String(data.Cf133) + ", " + data.Cf134.VerbatimString(true) + ", " + _dafny.String(data.Cf135) + ")"
		}
	case D32_DC80:
		{
			return "D32.DC80" + "(" + _dafny.String(data.Cf136) + ", " + _dafny.String(data.Cf137) + ", " + _dafny.String(data.Cf138) + ")"
		}
	case D32_DC78:
		{
			return "D32.DC78" + "(" + _dafny.String(data.Cf132) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D32) Equals(other D32) bool {
	switch data1 := _this.Get_().(type) {
	case D32_DC79:
		{
			data2, ok := other.Get_().(D32_DC79)
			return ok && data1.Cf133 == data2.Cf133 && data1.Cf134.Equals(data2.Cf134) && data1.Cf135.Equals(data2.Cf135)
		}
	case D32_DC80:
		{
			data2, ok := other.Get_().(D32_DC80)
			return ok && data1.Cf136.Cmp(data2.Cf136) == 0 && data1.Cf137.Cmp(data2.Cf137) == 0 && data1.Cf138.Cmp(data2.Cf138) == 0
		}
	case D32_DC78:
		{
			data2, ok := other.Get_().(D32_DC78)
			return ok && data1.Cf132 == data2.Cf132
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D32) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D32)
	return ok && _this.Equals(typed)
}

func Type_D32_() _dafny.TypeDescriptor {
	return type_D32_{}
}

type type_D32_ struct {
}

func (_this type_D32_) Default() interface{} {
	return Companion_D32_.Default()
}

func (_this type_D32_) String() string {
	return "main.D32"
}
func (_this D32) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D32{}

// End of datatype D32

// Definition of datatype D33
type D33 struct {
	Data_D33_
}

func (_this D33) Get_() Data_D33_ {
	return _this.Data_D33_
}

type Data_D33_ interface {
	isD33()
}

type CompanionStruct_D33_ struct {
}

var Companion_D33_ = CompanionStruct_D33_{}

type D33_DC82 struct {
	Cf140 bool
	Cf141 bool
	Cf142 T0
	Cf143 _dafny.Int
}

func (D33_DC82) isD33() {}

func (CompanionStruct_D33_) Create_DC82_(Cf140 bool, Cf141 bool, Cf142 T0, Cf143 _dafny.Int) D33 {
	return D33{D33_DC82{Cf140, Cf141, Cf142, Cf143}}
}

func (_this D33) Is_DC82() bool {
	_, ok := _this.Get_().(D33_DC82)
	return ok
}

type D33_DC81 struct {
	Cf139 _dafny.Map
}

func (D33_DC81) isD33() {}

func (CompanionStruct_D33_) Create_DC81_(Cf139 _dafny.Map) D33 {
	return D33{D33_DC81{Cf139}}
}

func (_this D33) Is_DC81() bool {
	_, ok := _this.Get_().(D33_DC81)
	return ok
}

type D33_DC83 struct {
	Cf144 D33
}

func (D33_DC83) isD33() {}

func (CompanionStruct_D33_) Create_DC83_(Cf144 D33) D33 {
	return D33{D33_DC83{Cf144}}
}

func (_this D33) Is_DC83() bool {
	_, ok := _this.Get_().(D33_DC83)
	return ok
}

func (CompanionStruct_D33_) Default() D33 {
	return Companion_D33_.Create_DC82_(false, false, (T0)(nil), _dafny.Zero)
}

func (_this D33) Dtor_cf140() bool {
	return _this.Get_().(D33_DC82).Cf140
}

func (_this D33) Dtor_cf141() bool {
	return _this.Get_().(D33_DC82).Cf141
}

func (_this D33) Dtor_cf142() T0 {
	return _this.Get_().(D33_DC82).Cf142
}

func (_this D33) Dtor_cf143() _dafny.Int {
	return _this.Get_().(D33_DC82).Cf143
}

func (_this D33) Dtor_cf139() _dafny.Map {
	return _this.Get_().(D33_DC81).Cf139
}

func (_this D33) Dtor_cf144() D33 {
	return _this.Get_().(D33_DC83).Cf144
}

func (_this D33) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D33_DC82:
		{
			return "D33.DC82" + "(" + _dafny.String(data.Cf140) + ", " + _dafny.String(data.Cf141) + ", " + _dafny.String(data.Cf142) + ", " + _dafny.String(data.Cf143) + ")"
		}
	case D33_DC81:
		{
			return "D33.DC81" + "(" + _dafny.String(data.Cf139) + ")"
		}
	case D33_DC83:
		{
			return "D33.DC83" + "(" + _dafny.String(data.Cf144) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D33) Equals(other D33) bool {
	switch data1 := _this.Get_().(type) {
	case D33_DC82:
		{
			data2, ok := other.Get_().(D33_DC82)
			return ok && data1.Cf140 == data2.Cf140 && data1.Cf141 == data2.Cf141 && _dafny.AreEqual(data1.Cf142, data2.Cf142) && data1.Cf143.Cmp(data2.Cf143) == 0
		}
	case D33_DC81:
		{
			data2, ok := other.Get_().(D33_DC81)
			return ok && data1.Cf139.Equals(data2.Cf139)
		}
	case D33_DC83:
		{
			data2, ok := other.Get_().(D33_DC83)
			return ok && data1.Cf144.Equals(data2.Cf144)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D33) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D33)
	return ok && _this.Equals(typed)
}

func Type_D33_() _dafny.TypeDescriptor {
	return type_D33_{}
}

type type_D33_ struct {
}

func (_this type_D33_) Default() interface{} {
	return Companion_D33_.Default()
}

func (_this type_D33_) String() string {
	return "main.D33"
}
func (_this D33) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D33{}

// End of datatype D33

// Definition of datatype D34
type D34 struct {
	Data_D34_
}

func (_this D34) Get_() Data_D34_ {
	return _this.Data_D34_
}

type Data_D34_ interface {
	isD34()
}

type CompanionStruct_D34_ struct {
}

var Companion_D34_ = CompanionStruct_D34_{}

type D34_DC85 struct {
	Cf146 _dafny.CodePoint
}

func (D34_DC85) isD34() {}

func (CompanionStruct_D34_) Create_DC85_(Cf146 _dafny.CodePoint) D34 {
	return D34{D34_DC85{Cf146}}
}

func (_this D34) Is_DC85() bool {
	_, ok := _this.Get_().(D34_DC85)
	return ok
}

type D34_DC86 struct {
	Cf147 bool
	Cf148 _dafny.Int
}

func (D34_DC86) isD34() {}

func (CompanionStruct_D34_) Create_DC86_(Cf147 bool, Cf148 _dafny.Int) D34 {
	return D34{D34_DC86{Cf147, Cf148}}
}

func (_this D34) Is_DC86() bool {
	_, ok := _this.Get_().(D34_DC86)
	return ok
}

type D34_DC84 struct {
	Cf145 _dafny.Sequence
}

func (D34_DC84) isD34() {}

func (CompanionStruct_D34_) Create_DC84_(Cf145 _dafny.Sequence) D34 {
	return D34{D34_DC84{Cf145}}
}

func (_this D34) Is_DC84() bool {
	_, ok := _this.Get_().(D34_DC84)
	return ok
}

func (CompanionStruct_D34_) Default() D34 {
	return Companion_D34_.Create_DC85_(_dafny.CodePoint('D'))
}

func (_this D34) Dtor_cf146() _dafny.CodePoint {
	return _this.Get_().(D34_DC85).Cf146
}

func (_this D34) Dtor_cf147() bool {
	return _this.Get_().(D34_DC86).Cf147
}

func (_this D34) Dtor_cf148() _dafny.Int {
	return _this.Get_().(D34_DC86).Cf148
}

func (_this D34) Dtor_cf145() _dafny.Sequence {
	return _this.Get_().(D34_DC84).Cf145
}

func (_this D34) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D34_DC85:
		{
			return "D34.DC85" + "(" + _dafny.String(data.Cf146) + ")"
		}
	case D34_DC86:
		{
			return "D34.DC86" + "(" + _dafny.String(data.Cf147) + ", " + _dafny.String(data.Cf148) + ")"
		}
	case D34_DC84:
		{
			return "D34.DC84" + "(" + _dafny.String(data.Cf145) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D34) Equals(other D34) bool {
	switch data1 := _this.Get_().(type) {
	case D34_DC85:
		{
			data2, ok := other.Get_().(D34_DC85)
			return ok && data1.Cf146 == data2.Cf146
		}
	case D34_DC86:
		{
			data2, ok := other.Get_().(D34_DC86)
			return ok && data1.Cf147 == data2.Cf147 && data1.Cf148.Cmp(data2.Cf148) == 0
		}
	case D34_DC84:
		{
			data2, ok := other.Get_().(D34_DC84)
			return ok && data1.Cf145.Equals(data2.Cf145)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D34) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D34)
	return ok && _this.Equals(typed)
}

func Type_D34_() _dafny.TypeDescriptor {
	return type_D34_{}
}

type type_D34_ struct {
}

func (_this type_D34_) Default() interface{} {
	return Companion_D34_.Default()
}

func (_this type_D34_) String() string {
	return "main.D34"
}
func (_this D34) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D34{}

// End of datatype D34

// Definition of datatype D35
type D35 struct {
	Data_D35_
}

func (_this D35) Get_() Data_D35_ {
	return _this.Data_D35_
}

type Data_D35_ interface {
	isD35()
}

type CompanionStruct_D35_ struct {
}

var Companion_D35_ = CompanionStruct_D35_{}

type D35_DC88 struct {
}

func (D35_DC88) isD35() {}

func (CompanionStruct_D35_) Create_DC88_() D35 {
	return D35{D35_DC88{}}
}

func (_this D35) Is_DC88() bool {
	_, ok := _this.Get_().(D35_DC88)
	return ok
}

type D35_DC89 struct {
	Cf150 _dafny.Int
	Cf151 _dafny.Int
	Cf152 _dafny.Int
}

func (D35_DC89) isD35() {}

func (CompanionStruct_D35_) Create_DC89_(Cf150 _dafny.Int, Cf151 _dafny.Int, Cf152 _dafny.Int) D35 {
	return D35{D35_DC89{Cf150, Cf151, Cf152}}
}

func (_this D35) Is_DC89() bool {
	_, ok := _this.Get_().(D35_DC89)
	return ok
}

type D35_DC87 struct {
	Cf149 _dafny.Map
}

func (D35_DC87) isD35() {}

func (CompanionStruct_D35_) Create_DC87_(Cf149 _dafny.Map) D35 {
	return D35{D35_DC87{Cf149}}
}

func (_this D35) Is_DC87() bool {
	_, ok := _this.Get_().(D35_DC87)
	return ok
}

func (CompanionStruct_D35_) Default() D35 {
	return Companion_D35_.Create_DC88_()
}

func (_this D35) Dtor_cf150() _dafny.Int {
	return _this.Get_().(D35_DC89).Cf150
}

func (_this D35) Dtor_cf151() _dafny.Int {
	return _this.Get_().(D35_DC89).Cf151
}

func (_this D35) Dtor_cf152() _dafny.Int {
	return _this.Get_().(D35_DC89).Cf152
}

func (_this D35) Dtor_cf149() _dafny.Map {
	return _this.Get_().(D35_DC87).Cf149
}

func (_this D35) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D35_DC88:
		{
			return "D35.DC88"
		}
	case D35_DC89:
		{
			return "D35.DC89" + "(" + _dafny.String(data.Cf150) + ", " + _dafny.String(data.Cf151) + ", " + _dafny.String(data.Cf152) + ")"
		}
	case D35_DC87:
		{
			return "D35.DC87" + "(" + _dafny.String(data.Cf149) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D35) Equals(other D35) bool {
	switch data1 := _this.Get_().(type) {
	case D35_DC88:
		{
			_, ok := other.Get_().(D35_DC88)
			return ok
		}
	case D35_DC89:
		{
			data2, ok := other.Get_().(D35_DC89)
			return ok && data1.Cf150.Cmp(data2.Cf150) == 0 && data1.Cf151.Cmp(data2.Cf151) == 0 && data1.Cf152.Cmp(data2.Cf152) == 0
		}
	case D35_DC87:
		{
			data2, ok := other.Get_().(D35_DC87)
			return ok && data1.Cf149.Equals(data2.Cf149)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D35) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D35)
	return ok && _this.Equals(typed)
}

func Type_D35_() _dafny.TypeDescriptor {
	return type_D35_{}
}

type type_D35_ struct {
}

func (_this type_D35_) Default() interface{} {
	return Companion_D35_.Default()
}

func (_this type_D35_) String() string {
	return "main.D35"
}
func (_this D35) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D35{}

// End of datatype D35

// Definition of trait T0
type T0 interface {
	String() string
	M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence)
	M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool)
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
	M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence)
	M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool)
	Fm0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool
	Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int
	F9() _dafny.Map
	F10() bool
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
	F4  _dafny.Map
	F8  bool
	_f0 bool
	_f1 bool
	_f2 _dafny.Array
	_f3 bool
	_f5 _dafny.Int
	_f6 bool
	_f7 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.EmptyMap
	_this.F8 = false
	_this._f0 = false
	_this._f1 = false
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = false
	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f7 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 _dafny.Array, f3 bool, f4 _dafny.Map, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
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
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
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
func (_this *C0) Fm5(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return ((_dafny.IntOfInt64(-952)).Minus(_dafny.IntOfInt64(834))).Cmp(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("aiohhcmxm"))).Cardinality(), _dafny.IntOfInt64(-442))) >= 0
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f12 bool
	_f13 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f12 = false
	_this._f13 = _dafny.EmptySeq
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

func (_this *C1) Ctor__(f12 bool, f13 _dafny.Sequence) {
	{
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C1) Fm32(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (_this).F12()
	}
}
func (_this *C1) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _459_v0 _dafny.Array
		_ = _459_v0
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
		_ = _nw45
		_459_v0 = _nw45
		var _460_v1 _dafny.Int
		_ = _460_v1
		_460_v1 = _dafny.IntOfInt64(261)
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))
		_ = _index50
		(_459_v0).ArraySet1(Companion_Default___.SafeDivisionInt(_460_v1, _460_v1), (_index50).Int())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))
		_ = _index51
		(_459_v0).ArraySet1(_460_v1, (_index51).Int())
		(globalState).F8 = (_this).F12()
		var _461_v2 _dafny.Map
		_ = _461_v2
		_461_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_460_v1, (_this).F12())
		_461_v2 = (_461_v2).Update((Companion_Default___.Fm33((_this).F12(), (_this).F12(), _460_v1, true, globalState)).Times((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)), ((_this).F12()) == (!((_this).F12())))
		var _462_v3 _dafny.CodePoint
		_ = _462_v3
		_462_v3 = _dafny.CodePoint('m')
		if (_this).Fm32((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int), _462_v3, globalState) {
			r0 = _dafny.UnicodeSeqOfUtf8Bytes("tsoaubs")
			var _463_v4 *C0
			_ = _463_v4
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__()
			_463_v4 = _nw46
			if ((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)).Cmp((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)) != 0 {
				(globalState).F8 = !((_this).F12())
				var _464_v5 _dafny.Array
				_ = _464_v5
				var _nw47 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw47
				_464_v5 = _nw47
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_464_v5), 0))
				_ = _index52
				(_464_v5).ArraySet1((func() bool {
					if (_461_v2).Contains((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)) {
						return (_461_v2).Get((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)).(bool)
					}
					return false
				})(), (_index52).Int())
				var _465_v6 _dafny.MultiSet
				_ = _465_v6
				_465_v6 = _dafny.MultiSetOf(_dafny.SeqOf((_this).F12()))
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_464_v5), 0))
				_ = _index53
				(_464_v5).ArraySet1((_465_v6).IsProperSubsetOf(_465_v6), (_index53).Int())
				var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw48
				_459_v0 = _nw48
				(globalState).F8 = (_464_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_464_v5), 0))).Int()).(bool)
				var _466_v7 _dafny.Array
				_ = _466_v7
				var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
				_ = _nw49
				_466_v7 = _nw49
				_466_v7 = _466_v7
			} else {
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))
				_ = _index54
				(_459_v0).ArraySet1(_dafny.IntOfInt64(606), (_index54).Int())
				var _467_v8 *C0
				_ = _467_v8
				var _nw50 *C0 = New_C0_()
				_ = _nw50
				_nw50.Ctor__()
				_467_v8 = _nw50
				var _468_v9 _dafny.MultiSet
				_ = _468_v9
				_468_v9 = _dafny.MultiSetOf(!((_this).F12()))
				_460_v1 = Companion_Default___.Fm33((_this).F12(), !((_467_v8).Fm5((_this).F12(), _460_v1, (_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int), _468_v9, globalState)), Companion_Default___.SafeModuloInt(_460_v1, _460_v1), (_this).F12(), globalState)
				var _469_v10 _dafny.Map
				_ = _469_v10
				_469_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v3, (_this).F13())
				var _470_v11 _dafny.Map
				_ = _470_v11
				_470_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int), _460_v1)
				var _rhs58 _dafny.Map = (_469_v10).Merge((_469_v10).Update(((_this).F13()).Select((Companion_Default___.SafeIndex((_470_v11).Cardinality(), _dafny.IntOfUint32(((_this).F13()).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F13()))
				_ = _rhs58
				var _rhs59 _dafny.Int = (_460_v1).Plus(_460_v1)
				_ = _rhs59
				var _rhs60 bool = (_this).F12()
				_ = _rhs60
				var _lhs44 *GlobalState = globalState
				_ = _lhs44
				_469_v10 = _rhs58
				_460_v1 = _rhs59
				_lhs44.F8 = _rhs60
				var _471_v12 _dafny.Array
				_ = _471_v12
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_5
				var _nw51 _dafny.Array
				_ = _nw51
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw51 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) bool = func(_472_i0 _dafny.Int) bool {
						return (_this).F12()
					}
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw51 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw51).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw51).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_471_v12 = _nw51
				var _473_v13 _dafny.MultiSet
				_ = _473_v13
				_473_v13 = _dafny.MultiSetOf(_471_v12)
				_473_v13 = ((_473_v13).Union(_473_v13)).Intersection(_473_v13)
			}
			var _474_v14 _dafny.MultiSet
			_ = _474_v14
			_474_v14 = _dafny.MultiSetOf(_dafny.IntOfInt64(367), Companion_Default___.Fm33((_this).F12(), (_this).F12(), (_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int), (_this).F12(), globalState))
			var _475_v15 _dafny.Array
			_ = _475_v15
			var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
			_ = _nw52
			_475_v15 = _nw52
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_475_v15), 0))
			_ = _index55
			(_475_v15).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F13(), _dafny.UnicodeSeqOfUtf8Bytes("ejfjtu")), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_475_v15), 0))
			_ = _index56
			var _rhs61 _dafny.MultiSet = ((_474_v14).Intersection(_dafny.MultiSetOf(_460_v1, _460_v1))).Intersection(_474_v14)
			_ = _rhs61
			var _rhs62 _dafny.Int = (_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs62
			var _rhs63 _dafny.Sequence = (func() _dafny.Sequence {
				if true {
					return (_this).F13()
				}
				return (func() _dafny.Sequence {
					if (_this).F12() {
						return (_this).F13()
					}
					return (_this).F13()
				})()
			})()
			_ = _rhs63
			var _lhs45 _dafny.Array = _475_v15
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_475_v15), 0))
			_ = _lhs46
			_474_v14 = _rhs61
			_460_v1 = _rhs62
			(_lhs45).ArraySet1(_rhs63, (_lhs46).Int())
			r1 = false
		} else {
			var _476_v16 D3
			_ = _476_v16
			_476_v16 = Companion_D3_.Create_DC7_(_459_v0)
			_476_v16 = (func() D3 {
				if !((_this).F12()) || ((_this).F12()) {
					return (func() D3 {
						if (_this).F12() {
							return Companion_D3_.Create_DC7_(_459_v0)
						}
						return _476_v16
					})()
				}
				return _476_v16
			})()
			var _477_v17 _dafny.Set
			_ = _477_v17
			_477_v17 = _dafny.SetOf(_460_v1, Companion_Default___.Fm33((_this).F12(), (_this).F12(), (_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int), (_this).F12(), globalState))
			var _478_v18 _dafny.Map
			_ = _478_v18
			_478_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_477_v17).IsDisjointFrom(_477_v17), (_460_v1).Minus((_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int)))
			_478_v18 = (_478_v18).Update((_this).F12(), (_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int))
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))
			_ = _index57
			(_459_v0).ArraySet1((_dafny.Zero).Minus(_460_v1), (_index57).Int())
			var _479_v19 _dafny.Array
			_ = _479_v19
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw53
			_479_v19 = _nw53
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_479_v19), 0))
			_ = _index58
			(_479_v19).ArraySet1((_this).F12(), (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_479_v19), 0))
			_ = _index59
			(_479_v19).ArraySet1((_this).F12(), (_index59).Int())
			var _480_v20 _dafny.Array
			_ = _480_v20
			var _nwElement0_10 _dafny.Array = _459_v0
			_ = _nwElement0_10
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(4))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_10, 0)
			(_nw54).ArraySet1(_459_v0, 1)
			(_nw54).ArraySet1(_459_v0, 2)
			(_nw54).ArraySet1(_459_v0, 3)
			_480_v20 = _nw54
			var _nwElement0_11 _dafny.Array = (_476_v16).Dtor_cf11()
			_ = _nwElement0_11
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(28))
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_11, 0)
			(_nw55).ArraySet1(_459_v0, 1)
			(_nw55).ArraySet1(_459_v0, 2)
			(_nw55).ArraySet1(_459_v0, 3)
			(_nw55).ArraySet1(_459_v0, 4)
			(_nw55).ArraySet1(_459_v0, 5)
			(_nw55).ArraySet1(_459_v0, 6)
			(_nw55).ArraySet1(_459_v0, 7)
			(_nw55).ArraySet1(_459_v0, 8)
			(_nw55).ArraySet1(_459_v0, 9)
			(_nw55).ArraySet1(_459_v0, 10)
			(_nw55).ArraySet1(_459_v0, 11)
			(_nw55).ArraySet1(_459_v0, 12)
			(_nw55).ArraySet1(_459_v0, 13)
			(_nw55).ArraySet1((func() _dafny.Array {
				if (_479_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_479_v19), 0))).Int()).(bool) {
					return _459_v0
				}
				return _459_v0
			})(), 14)
			(_nw55).ArraySet1(_459_v0, 15)
			(_nw55).ArraySet1(_459_v0, 16)
			(_nw55).ArraySet1(_459_v0, 17)
			(_nw55).ArraySet1((func() _dafny.Array {
				if (_479_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_479_v19), 0))).Int()).(bool) {
					return _459_v0
				}
				return _459_v0
			})(), 18)
			(_nw55).ArraySet1(_459_v0, 19)
			(_nw55).ArraySet1((func() _dafny.Array {
				if (_479_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_479_v19), 0))).Int()).(bool) {
					return _459_v0
				}
				return _459_v0
			})(), 20)
			(_nw55).ArraySet1(_459_v0, 21)
			(_nw55).ArraySet1(_459_v0, 22)
			(_nw55).ArraySet1(_459_v0, 23)
			(_nw55).ArraySet1(_459_v0, 24)
			(_nw55).ArraySet1(_459_v0, 25)
			(_nw55).ArraySet1(_459_v0, 26)
			(_nw55).ArraySet1((Companion_D3_.Create_DC7_(_459_v0)).Dtor_cf11(), 27)
			_480_v20 = _nw55
		}
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))
		_ = _index60
		(_459_v0).ArraySet1(_460_v1, (_index60).Int())
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))
		_ = _index61
		(_459_v0).ArraySet1(_460_v1, (_index61).Int())
		r0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F13(), _dafny.Companion_Sequence_.Concatenate((_this).F13(), (_this).F13())), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-935), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F13(), _dafny.Companion_Sequence_.Concatenate((_this).F13(), (_this).F13()))).Cardinality()))).Uint32(), _462_v3)
		r1 = (func() bool {
			if (_this).F12() {
				return (_this).F12()
			}
			return (_this).F12()
		})()
		var _481_v21 _dafny.Set
		_ = _481_v21
		_481_v21 = _dafny.SetOf((_this).Fm32(_460_v1, _462_v3, globalState), (_this).F12(), false)
		var _482_v22 _dafny.Sequence
		_ = _482_v22
		_482_v22 = _dafny.SeqOf(_481_v21)
		var _483_v23 _dafny.Map
		_ = _483_v23
		_483_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_459_v0, (_459_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_459_v0), 0))).Int()).(_dafny.Int))
		r2 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_481_v21, (_482_v22).Select((Companion_Default___.SafeIndex(_460_v1, _dafny.IntOfUint32((_482_v22).Cardinality()))).Uint32()).(_dafny.Set)), _dafny.SeqOf(_481_v21, _481_v21, Companion_Default___.Fm34((_483_v23).Cardinality(), (_this).F12(), (_this).F12(), (_this).F12(), globalState), _dafny.SetOf(false)))
		r3 = _dafny.Companion_Sequence_.Concatenate((_this).F13(), (_this).F13())
		return r0, r1, r2, r3
	}
}
func (_this *C1) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _484_v0 _dafny.Int
		_ = _484_v0
		_484_v0 = _dafny.IntOfInt64(776)
		var _485_v1 _dafny.Map
		_ = _485_v1
		_485_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_v0, (_this).F12())
		var _486_v2 D0
		_ = _486_v2
		_486_v2 = Companion_D0_.Create_DC1_(_484_v0, (_this).Fm32((_485_v1).Cardinality(), p0, globalState), (_dafny.Zero).Minus(_484_v0))
		var _487_v4 _dafny.Map
		_ = _487_v4
		_487_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_484_v0).Cmp((_486_v2).Dtor_cf3()) == 0, func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(7), _dafny.IntOfInt64(695))); ; {
				_compr_29, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _488_v3 _dafny.Int
				_488_v3 = interface{}(_compr_29).(_dafny.Int)
				if ((_dafny.IntOfInt64(7)).Cmp(_488_v3) <= 0) && ((_488_v3).Cmp(_dafny.IntOfInt64(695)) < 0) {
					_coll29.Add((_488_v3).Times(_484_v0), _484_v0)
				}
			}
			return _coll29.ToMap()
		}())
		var _489_v5 _dafny.Map
		_ = _489_v5
		_489_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(620))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg45 _dafny.Int) interface{} {
				return coer45(arg45)
			}
		}(func(_490_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(674)
		}))).Cardinality()), _dafny.IntOfInt64(54))
		_487_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm32(_484_v0, p0, globalState)), _489_v5)
		var _491_v6 *C0
		_ = _491_v6
		var _nw56 *C0 = New_C0_()
		_ = _nw56
		_nw56.Ctor__()
		_491_v6 = _nw56
		var _492_v7 _dafny.CodePoint
		_ = _492_v7
		_492_v7 = _dafny.CodePoint('i')
		var _pat_let_tv6 = p0
		_ = _pat_let_tv6
		var _pat_let_tv7 = p0
		_ = _pat_let_tv7
		var _pat_let_tv8 = p0
		_ = _pat_let_tv8
		_492_v7 = func(_source12 D3) _dafny.CodePoint {
			if _source12.Is_DC8() {
				var _493___mcc_h0 bool = _source12.Get_().(D3_DC8).Cf12
				_ = _493___mcc_h0
				var _494___mcc_h1 bool = _source12.Get_().(D3_DC8).Cf13
				_ = _494___mcc_h1
				var _495_cf13 bool = _494___mcc_h1
				_ = _495_cf13
				var _496_cf12 bool = _493___mcc_h0
				_ = _496_cf12
				if _495_cf13 {
					return _dafny.CodePoint('o')
				} else {
					return _pat_let_tv6
				}
			} else if _source12.Is_DC7() {
				var _497___mcc_h2 _dafny.Array = _source12.Get_().(D3_DC7).Cf11
				_ = _497___mcc_h2
				var _498_cf11 _dafny.Array = _497___mcc_h2
				_ = _498_cf11
				return _pat_let_tv7
			} else {
				var _499___mcc_h3 D3 = _source12.Get_().(D3_DC9).Cf14
				_ = _499___mcc_h3
				var _500_cf14 D3 = _499___mcc_h3
				_ = _500_cf14
				return (_pat_let_tv8)
			}
		}(Companion_Default___.Fm35((_this).F12(), globalState))
		var _501_v8 _dafny.MultiSet
		_ = _501_v8
		_501_v8 = _dafny.MultiSetOf((_this).F12())
		var _502_v9 D1
		_ = _502_v9
		_502_v9 = Companion_D1_.Create_DC4_((_491_v6).Fm5((_this).F12(), _484_v0, _484_v0, _501_v8, globalState), (_this).F12(), (func() _dafny.Int {
			if (_this).F12() {
				return (_501_v8).Cardinality()
			}
			return _484_v0
		})())
		var _source13 D1 = _502_v9
		_ = _source13
		if _source13.Is_DC3() {
			var _503___mcc_h4 _dafny.Sequence = _source13.Get_().(D1_DC3).Cf5
			_ = _503___mcc_h4
			var _504_cf5 _dafny.Sequence = _503___mcc_h4
			_ = _504_cf5
			var _505_v10 _dafny.MultiSet
			_ = _505_v10
			_505_v10 = _dafny.MultiSetOf(_484_v0)
			var _506_v11 _dafny.Map
			_ = _506_v11
			_506_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _505_v10)
			var _507_v12 _dafny.Sequence
			_ = _507_v12
			_507_v12 = _dafny.SeqOf(_484_v0)
			_506_v11 = (_506_v11).Update(!((_this).F12()) || (true), (_dafny.MultiSetFromSeq(_507_v12)).Difference(_505_v10))
			(globalState).F8 = (_dafny.IntOfInt64(924)).Cmp((_501_v8).Cardinality()) != 0
			var _508_v13 _dafny.Array
			_ = _508_v13
			var _nwElement0_12 bool = (_this).F12()
			_ = _nwElement0_12
			var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(16))
			_ = _nw57
			(_nw57).ArraySet1(_nwElement0_12, 0)
			(_nw57).ArraySet1((_this).F12(), 1)
			(_nw57).ArraySet1((_this).F12(), 2)
			(_nw57).ArraySet1((_this).F12(), 3)
			(_nw57).ArraySet1((_this).F12(), 4)
			(_nw57).ArraySet1((_dafny.IntOfUint32(((_this).F13()).Cardinality())).Cmp(_484_v0) < 0, 5)
			(_nw57).ArraySet1((_this).F12(), 6)
			(_nw57).ArraySet1(false, 7)
			(_nw57).ArraySet1((func() bool {
				if (_this).F12() {
					return (_this).F12()
				}
				return (_this).F12()
			})(), 8)
			(_nw57).ArraySet1((_this).F12(), 9)
			(_nw57).ArraySet1((_this).F12(), 10)
			(_nw57).ArraySet1((_this).F12(), 11)
			(_nw57).ArraySet1((_this).F12(), 12)
			(_nw57).ArraySet1(!((func() bool {
				if (_this).F12() {
					return (_this).F12()
				}
				return (_this).F12()
			})()), 13)
			(_nw57).ArraySet1(!(!((_491_v6).Fm5((_this).F12(), _484_v0, _484_v0, _501_v8, globalState)) || ((_this).F12())), 14)
			(_nw57).ArraySet1(((_this).F12()) || ((_this).F12()), 15)
			_508_v13 = _nw57
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_508_v13), 0))
			_ = _index62
			(_508_v13).ArraySet1((_this).F12(), (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_508_v13), 0))
			_ = _index63
			(_508_v13).ArraySet1(!((_this).F12()), (_index63).Int())
			var _509_v14 _dafny.Sequence
			_ = _509_v14
			_509_v14 = _dafny.SeqOf((_508_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_508_v13), 0))).Int()).(bool), (_this).F12())
			_492_v7 = (Companion_D4_.Create_DC12_(true, _dafny.IntOfUint32((_509_v14).Cardinality()), Companion_Default___.Fm36(globalState), _dafny.IntOfInt64(739), _484_v0)).Dtor_cf21()
		} else if _source13.Is_DC4() {
			var _510___mcc_h5 bool = _source13.Get_().(D1_DC4).Cf6
			_ = _510___mcc_h5
			var _511___mcc_h6 bool = _source13.Get_().(D1_DC4).Cf7
			_ = _511___mcc_h6
			var _512___mcc_h7 _dafny.Int = _source13.Get_().(D1_DC4).Cf8
			_ = _512___mcc_h7
			var _513_cf8 _dafny.Int = _512___mcc_h7
			_ = _513_cf8
			var _514_cf7 bool = _511___mcc_h6
			_ = _514_cf7
			var _515_cf6 bool = _510___mcc_h5
			_ = _515_cf6
			_513_cf8 = _dafny.IntOfInt64(-389)
			var _516_v15 _dafny.Set
			_ = _516_v15
			_516_v15 = _dafny.SetOf(true, _515_cf6)
			_514_cf7 = (_516_v15).Contains(_515_cf6)
			_491_v6 = _491_v6
			var _517_v16 D0
			_ = _517_v16
			_517_v16 = Companion_D0_.Create_DC0_((_this).F12())
			var _518_v17 _dafny.Sequence
			_ = _518_v17
			_518_v17 = _dafny.SeqOf(_514_cf7, (_517_v16).Dtor_cf0())
			if (_515_cf6) || ((_dafny.MultiSetFromSeq(_518_v17)).IsSubsetOf(_501_v8)) {
				r1 = ((_516_v15).Intersection(_dafny.SetOf(_514_cf7))).IsSubsetOf(_516_v15)
				(globalState).F8 = _515_cf6
				var _519_v18 _dafny.Sequence
				_ = _519_v18
				_519_v18 = _dafny.SeqOf(_484_v0, _484_v0)
				_484_v0 = (_513_cf8).Minus(_dafny.IntOfUint32((_519_v18).Cardinality()))
				var _520_v19 *C0
				_ = _520_v19
				var _nw58 *C0 = New_C0_()
				_ = _nw58
				_nw58.Ctor__()
				_520_v19 = _nw58
				var _rhs64 bool = _514_cf7
				_ = _rhs64
				var _rhs65 bool = _514_cf7
				_ = _rhs65
				_515_cf6 = _rhs64
				r1 = _rhs65
			} else {
				var _521_v20 _dafny.Sequence
				_ = _521_v20
				_521_v20 = _dafny.UnicodeSeqOfUtf8Bytes("xrh")
				_521_v20 = _521_v20
				_513_cf8 = ((_dafny.Zero).Minus(_513_cf8)).Minus((func() _dafny.Int {
					if (_489_v5).Contains(_484_v0) {
						return (_489_v5).Get(_484_v0).(_dafny.Int)
					}
					return _513_cf8
				})())
				var _522_v21 _dafny.Array
				_ = _522_v21
				var _nwElement0_13 *C0 = _491_v6
				_ = _nwElement0_13
				var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(3))
				_ = _nw59
				(_nw59).ArraySet1(_nwElement0_13, 0)
				(_nw59).ArraySet1(_491_v6, 1)
				(_nw59).ArraySet1(_491_v6, 2)
				_522_v21 = _nw59
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_522_v21), 0))
				_ = _index64
				(_522_v21).ArraySet1(_491_v6, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_522_v21), 0))
				_ = _index65
				(_522_v21).ArraySet1(_491_v6, (_index65).Int())
				(globalState).F8 = _515_cf6
				var _523_v22 _dafny.Sequence
				_ = _523_v22
				_523_v22 = _dafny.SeqOf(_513_cf8, (Companion_Default___.Fm33(_515_cf6, (_this).F12(), _513_cf8, (_this).F12(), globalState)).Minus(_513_cf8), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_521_v20).Cardinality()), _484_v0))
				_523_v22 = Companion_Default___.Fm37(globalState)
			}
		} else if _source13.Is_DC2() {
			var _524___mcc_h8 _dafny.Sequence = _source13.Get_().(D1_DC2).Cf4
			_ = _524___mcc_h8
			var _525_cf4 _dafny.Sequence = _524___mcc_h8
			_ = _525_cf4
			_492_v7 = _492_v7
			(globalState).F8 = ((_this).F12()) == ((_this).F12())
			(globalState).F8 = (_this).F12()
			if (_this).F12() {
				var _526_v23 _dafny.Sequence
				_ = _526_v23
				_526_v23 = _dafny.SeqOf(true, (_this).F12())
				var _527_v24 _dafny.Map
				_ = _527_v24
				_527_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F12()), (_this).F12())
				var _528_v25 _dafny.Sequence
				_ = _528_v25
				_528_v25 = _dafny.SeqOf(_526_v23)
				var _529_v26 _dafny.Array
				_ = _529_v26
				var _nwElement0_14 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23)
				_ = _nwElement0_14
				var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(18))
				_ = _nw60
				(_nw60).ArraySet1(_nwElement0_14, 0)
				(_nw60).ArraySet1(Companion_Default___.Fm38((func() bool {
					if (_527_v24).Contains((_this).F12()) {
						return (_527_v24).Get((_this).F12()).(bool)
					}
					return (_this).F12()
				})(), globalState), 1)
				(_nw60).ArraySet1(_526_v23, 2)
				(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_526_v23, (_528_v25).Select((Companion_Default___.SafeIndex(_484_v0, _dafny.IntOfUint32((_528_v25).Cardinality()))).Uint32()).(_dafny.Sequence)), 3)
				(_nw60).ArraySet1((func() _dafny.Sequence {
					if (_this).F12() {
						return _526_v23
					}
					return _526_v23
				})(), 4)
				(_nw60).ArraySet1(_526_v23, 5)
				(_nw60).ArraySet1(_dafny.SeqOf((_this).F12(), !((_this).F12())), 6)
				(_nw60).ArraySet1(_526_v23, 7)
				(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23), 8)
				(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23), 9)
				(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23), 10)
				(_nw60).ArraySet1(_526_v23, 11)
				(_nw60).ArraySet1(Companion_Default___.Fm38((_this).F12(), globalState), 12)
				(_nw60).ArraySet1(_526_v23, 13)
				(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23), 14)
				(_nw60).ArraySet1(_526_v23, 15)
				(_nw60).ArraySet1(_dafny.SeqOf((_this).F12(), !((_491_v6).Fm5((_this).F12(), _484_v0, _484_v0, _501_v8, globalState)), (_this).F12()), 16)
				(_nw60).ArraySet1(_526_v23, 17)
				_529_v26 = _nw60
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_529_v26), 0))
				_ = _index66
				(_529_v26).ArraySet1(_dafny.Companion_Sequence_.Update(_526_v23, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_484_v0), _dafny.IntOfUint32((_526_v23).Cardinality()))).Uint32(), (_this).F12()), (_index66).Int())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_529_v26), 0))
				_ = _index67
				(_529_v26).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23), (Companion_Default___.SafeIndex(_484_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_526_v23, _526_v23)).Cardinality()))).Uint32(), (_this).F12()), (_index67).Int())
				(globalState).F8 = ((_529_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_529_v26), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_484_v0, _dafny.IntOfUint32(((_529_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_529_v26), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(bool)
				(globalState).F8 = (_this).F12()
				var _530_v27 *C0
				_ = _530_v27
				var _nw61 *C0 = New_C0_()
				_ = _nw61
				_nw61.Ctor__()
				_530_v27 = _nw61
				_484_v0 = _484_v0
			} else {
				_492_v7 = _dafny.CodePoint('m')
				_484_v0 = (func() _dafny.Int {
					if (_484_v0).Cmp(_484_v0) == 0 {
						return _484_v0
					}
					return _484_v0
				})()
				var _531_v28 _dafny.Array
				_ = _531_v28
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_6
				var _nw62 _dafny.Array
				_ = _nw62
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw62 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Int = (func(_532_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_533_i1 _dafny.Int) _dafny.Int {
							return (_533_i1).Plus(_532_v0)
						}
					})(_484_v0)
					_ = _init6
					var _element0_6 = _init6(_dafny.Zero)
					_ = _element0_6
					_nw62 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
					(_nw62).ArraySet1(_element0_6, 0)
					var _nativeLen0_6 = (_len0_6).Int()
					_ = _nativeLen0_6
					for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
						(_nw62).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
					}
				}
				_531_v28 = _nw62
				var _534_v29 _dafny.Map
				_ = _534_v29
				_534_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), _531_v28)
				var _535_v30 _dafny.Array
				_ = _535_v30
				var _nwElement0_15 _dafny.Array = _531_v28
				_ = _nwElement0_15
				var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(16))
				_ = _nw63
				(_nw63).ArraySet1(_nwElement0_15, 0)
				(_nw63).ArraySet1(_531_v28, 1)
				(_nw63).ArraySet1(_531_v28, 2)
				(_nw63).ArraySet1(_531_v28, 3)
				(_nw63).ArraySet1(_531_v28, 4)
				(_nw63).ArraySet1(_531_v28, 5)
				(_nw63).ArraySet1(_531_v28, 6)
				(_nw63).ArraySet1(_531_v28, 7)
				(_nw63).ArraySet1(_531_v28, 8)
				(_nw63).ArraySet1(_531_v28, 9)
				(_nw63).ArraySet1(_531_v28, 10)
				(_nw63).ArraySet1(_531_v28, 11)
				(_nw63).ArraySet1(_531_v28, 12)
				(_nw63).ArraySet1((func() _dafny.Array {
					if (_534_v29).Contains((_this).F12()) {
						return (_534_v29).Get((_this).F12()).(_dafny.Array)
					}
					return _531_v28
				})(), 13)
				(_nw63).ArraySet1(_531_v28, 14)
				(_nw63).ArraySet1(_531_v28, 15)
				_535_v30 = _nw63
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_535_v30), 0))
				_ = _index68
				(_535_v30).ArraySet1(_531_v28, (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_535_v30), 0))
				_ = _index69
				(_535_v30).ArraySet1(_531_v28, (_index69).Int())
				var _536_v31 _dafny.Sequence
				_ = _536_v31
				_536_v31 = _dafny.SeqOf((_this).F12(), !((_this).F12()))
				var _537_v32 _dafny.Map
				_ = _537_v32
				_537_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_this).F12())
				var _538_v33 D6
				_ = _538_v33
				_538_v33 = Companion_D6_.Create_DC16_(_537_v32)
				var _539_v34 _dafny.Map
				_ = _539_v34
				_539_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_536_v31).Select((Companion_Default___.SafeIndex(((_538_v33).Dtor_cf28()).Cardinality(), _dafny.IntOfUint32((_536_v31).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(-724))
				_539_v34 = (_539_v34).Update(!((_this).F12()), _484_v0)
				(globalState).F8 = (_484_v0).Cmp(_484_v0) < 0
			}
		} else {
			var _540___mcc_h9 D1 = _source13.Get_().(D1_DC5).Cf9
			_ = _540___mcc_h9
			var _541_cf9 D1 = _540___mcc_h9
			_ = _541_cf9
			_484_v0 = _484_v0
			r1 = false
			var _542_v35 D7
			_ = _542_v35
			_542_v35 = Companion_D7_.Create_DC21_(_dafny.SetOf((_this).F12()))
			var _543_v36 _dafny.Set
			_ = _543_v36
			_543_v36 = _dafny.SetOf((_this).F12(), (_this).F12(), (_this).F12())
			var _544_v37 _dafny.MultiSet
			_ = _544_v37
			_544_v37 = _dafny.MultiSetOf(_484_v0, _dafny.IntOfInt64(-245), (((_542_v35).Dtor_cf40()).Union(_543_v36)).Cardinality(), _484_v0, _484_v0)
			_544_v37 = (_544_v37).Difference(_544_v37)
			if (_484_v0).Cmp(_484_v0) != 0 {
				_485_v1 = (_485_v1).Merge((func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter30 := _dafny.Iterate((_544_v37).Elements()); ; {
						_compr_30, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _545_v38 _dafny.Int
						_545_v38 = interface{}(_compr_30).(_dafny.Int)
						if (_544_v37).Contains(_545_v38) {
							_coll30.Add(Companion_Default___.SafeModuloInt(_545_v38, _dafny.IntOfUint32(((_this).F13()).Cardinality())), (_this).F12())
						}
					}
					return _coll30.ToMap()
				}()).Update(_484_v0, (_this).F12()))
				var _546_v39 _dafny.Array
				_ = _546_v39
				var _nwElement0_16 _dafny.CodePoint = _492_v7
				_ = _nwElement0_16
				var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(26))
				_ = _nw64
				(_nw64).ArraySet1CodePoint(_nwElement0_16, 0)
				(_nw64).ArraySet1CodePoint(p0, 1)
				(_nw64).ArraySet1CodePoint(_492_v7, 2)
				(_nw64).ArraySet1CodePoint(_492_v7, 3)
				(_nw64).ArraySet1CodePoint(p0, 4)
				(_nw64).ArraySet1CodePoint(p0, 5)
				(_nw64).ArraySet1CodePoint(p0, 6)
				(_nw64).ArraySet1CodePoint(_492_v7, 7)
				(_nw64).ArraySet1CodePoint(_492_v7, 8)
				(_nw64).ArraySet1CodePoint(p0, 9)
				(_nw64).ArraySet1CodePoint(_492_v7, 10)
				(_nw64).ArraySet1CodePoint(p0, 11)
				(_nw64).ArraySet1CodePoint(_dafny.CodePoint('v'), 12)
				(_nw64).ArraySet1CodePoint(_dafny.CodePoint('m'), 13)
				(_nw64).ArraySet1CodePoint(_492_v7, 14)
				(_nw64).ArraySet1CodePoint(p0, 15)
				(_nw64).ArraySet1CodePoint(p0, 16)
				(_nw64).ArraySet1CodePoint(_492_v7, 17)
				(_nw64).ArraySet1CodePoint(p0, 18)
				(_nw64).ArraySet1CodePoint(p0, 19)
				(_nw64).ArraySet1CodePoint(p0, 20)
				(_nw64).ArraySet1CodePoint(p0, 21)
				(_nw64).ArraySet1CodePoint(_492_v7, 22)
				(_nw64).ArraySet1CodePoint(p0, 23)
				(_nw64).ArraySet1CodePoint(_492_v7, 24)
				(_nw64).ArraySet1CodePoint(_492_v7, 25)
				_546_v39 = _nw64
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_546_v39), 0))
				_ = _index70
				(_546_v39).ArraySet1CodePoint(_492_v7, (_index70).Int())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_546_v39), 0))
				_ = _index71
				(_546_v39).ArraySet1CodePoint(_dafny.CodePoint('b'), (_index71).Int())
				var _547_v40 _dafny.Sequence
				_ = _547_v40
				_547_v40 = _dafny.SeqOf(_dafny.IntOfInt64(-832))
				var _rhs66 bool = (_this).F12()
				_ = _rhs66
				var _rhs67 _dafny.Int = ((_543_v36).Cardinality()).Times(Companion_Default___.SafeModuloInt(_484_v0, _dafny.IntOfUint32((_547_v40).Cardinality())))
				_ = _rhs67
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				_lhs47.F8 = _rhs66
				_484_v0 = _rhs67
				var _548_v41 _dafny.Sequence
				_ = _548_v41
				_548_v41 = _dafny.UnicodeSeqOfUtf8Bytes("rax")
				_548_v41 = (_this).F13()
				_485_v1 = (_485_v1).Update(_484_v0, true)
			} else {
				var _549_v42 *C0
				_ = _549_v42
				var _nw65 *C0 = New_C0_()
				_ = _nw65
				_nw65.Ctor__()
				_549_v42 = _nw65
				var _550_v43 _dafny.Sequence
				_ = _550_v43
				_550_v43 = _dafny.SeqOf((_this).F12())
				_550_v43 = _550_v43
				var _551_v44 _dafny.Sequence
				_ = _551_v44
				_551_v44 = _dafny.UnicodeSeqOfUtf8Bytes("feeckpr")
				_551_v44 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(9))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_552_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_553_i2 _dafny.Int) _dafny.CodePoint {
						return _552_p0
					}
				})(p0)))
				var _554_v45 _dafny.Map
				_ = _554_v45
				_554_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_this).F12())
				var _555_v46 _dafny.Set
				_ = _555_v46
				_555_v46 = _dafny.SetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-682), Companion_Default___.Fm33((_this).F12(), (_this).F12(), (_554_v45).Cardinality(), (Companion_D1_.Create_DC4_((_this).F12(), (_this).F12(), _484_v0)).Dtor_cf6(), globalState)))
				_555_v46 = func() _dafny.Set {
					var _coll31 = _dafny.NewBuilder()
					_ = _coll31
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(825), _dafny.IntOfInt64(781))); ; {
						_compr_31, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _556_v47 _dafny.Int
						_556_v47 = interface{}(_compr_31).(_dafny.Int)
						if ((_dafny.IntOfInt64(825)).Cmp(_556_v47) <= 0) && ((_556_v47).Cmp(_dafny.IntOfInt64(781)) < 0) {
							_coll31.Add((_556_v47).Times(_dafny.IntOfUint32((_550_v43).Cardinality())))
						}
					}
					return _coll31.ToSet()
				}()
				var _557_v48 _dafny.Array
				_ = _557_v48
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw66
				_557_v48 = _nw66
				_557_v48 = _557_v48
			}
		}
		var _558_v49 *C0
		_ = _558_v49
		var _nw67 *C0 = New_C0_()
		_ = _nw67
		_nw67.Ctor__()
		_558_v49 = _nw67
		if (_this).F12() {
			if (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xplooehy")).Cardinality()), _484_v0)).Cmp(_484_v0) < 0 {
				var _559_v50 D3
				_ = _559_v50
				_559_v50 = Companion_D3_.Create_DC8_((_this).F12(), (_this).F12())
				_559_v50 = (func() D3 {
					if (_this).F12() {
						return _559_v50
					}
					return func(_pat_let6_0 D3) D3 {
						return func(_560_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let7_0 bool) D3 {
								return func(_561_dt__update_hcf12_h0 bool) D3 {
									return Companion_D3_.Create_DC8_(_561_dt__update_hcf12_h0, (_560_dt__update__tmp_h0).Dtor_cf13())
								}(_pat_let7_0)
							}((_this).F12())
						}(_pat_let6_0)
					}(_559_v50)
				})()
				var _562_v51 _dafny.Sequence
				_ = _562_v51
				_562_v51 = _dafny.SeqOf(_485_v1)
				var _563_v52 _dafny.Map
				_ = _563_v52
				_563_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_562_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F13()).Cardinality()), _dafny.IntOfUint32((_562_v51).Cardinality()))).Uint32()).(_dafny.Map)).Update(_dafny.IntOfInt64(-24), (_this).F12()), _484_v0)
				_484_v0 = ((_563_v52).Update(_485_v1, (func() _dafny.Int {
					if (_this).F12() {
						return _484_v0
					}
					return _dafny.IntOfInt64(-270)
				})())).Cardinality()
				(globalState).F8 = (_491_v6).Fm5((_this).F12(), (_dafny.Zero).Minus(_484_v0), Companion_Default___.SafeDivisionInt(_484_v0, _484_v0), _501_v8, globalState)
				var _564_v53 *C0
				_ = _564_v53
				var _nw68 *C0 = New_C0_()
				_ = _nw68
				_nw68.Ctor__()
				_564_v53 = _nw68
				var _565_v55 _dafny.Set
				_ = _565_v55
				_565_v55 = _dafny.SetOf(Companion_Default___.Fm33((_this).F12(), (_this).F12(), _dafny.IntOfInt64(953), (_this).F12(), globalState), _484_v0, _484_v0, _484_v0, _484_v0)
				var _566_v56 _dafny.MultiSet
				_ = _566_v56
				_566_v56 = _dafny.MultiSetOf(func() _dafny.Set {
					var _coll32 = _dafny.NewBuilder()
					_ = _coll32
					for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(954), _dafny.IntOfInt64(800))); ; {
						_compr_32, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _567_v54 _dafny.Int
						_567_v54 = interface{}(_compr_32).(_dafny.Int)
						if ((_dafny.IntOfInt64(954)).Cmp(_567_v54) <= 0) && ((_567_v54).Cmp(_dafny.IntOfInt64(800)) < 0) {
							_coll32.Add((_567_v54).Plus((func() _dafny.Int {
								if (_489_v5).Contains(_484_v0) {
									return (_489_v5).Get(_484_v0).(_dafny.Int)
								}
								return _484_v0
							})()))
						}
					}
					return _coll32.ToSet()
				}(), _565_v55, _565_v55)
				_566_v56 = (_566_v56).Union(_566_v56)
			} else {
				var _568_v57 _dafny.Array
				_ = _568_v57
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_7
				var _nw69 _dafny.Array
				_ = _nw69
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw69 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = func(_569_i3 _dafny.Int) bool {
						return true
					}
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw69 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw69).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw69).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_568_v57 = _nw69
				var _570_v58 _dafny.Array
				_ = _570_v58
				var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw70
				_570_v58 = _nw70
				var _571_v59 D8
				_ = _571_v59
				_571_v59 = Companion_D8_.Create_DC23_(_dafny.SeqOf((_this).F12(), (_this).F12()))
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))
				_ = _index72
				(_570_v58).ArraySet1((_dafny.MultiSetFromSeq((_571_v59).Dtor_cf41())).Cardinality(), (_index72).Int())
				var _572_v60 _dafny.Array
				_ = _572_v60
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
				_ = _nw71
				_572_v60 = _nw71
				var _573_v61 _dafny.Map
				_ = _573_v61
				_573_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_572_v60, (_484_v0).Times(Companion_Default___.Fm33((_this).F12(), (_this).Fm32(_484_v0, _492_v7, globalState), _484_v0, true, globalState)))
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))
				_ = _index73
				var _rhs68 _dafny.Array = _568_v57
				_ = _rhs68
				var _rhs69 bool = (_this).F12()
				_ = _rhs69
				var _rhs70 _dafny.Int = (func() _dafny.Int {
					if (_573_v61).Contains(_572_v60) {
						return (_573_v61).Get(_572_v60).(_dafny.Int)
					}
					return _484_v0
				})()
				_ = _rhs70
				var _lhs48 _dafny.Array = _570_v58
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))
				_ = _lhs49
				_568_v57 = _rhs68
				r1 = _rhs69
				(_lhs48).ArraySet1(_rhs70, (_lhs49).Int())
				var _574_v62 _dafny.Set
				_ = _574_v62
				_574_v62 = _dafny.SetOf(_570_v58, _570_v58)
				var _575_v63 D8
				_ = _575_v63
				_575_v63 = Companion_D8_.Create_DC24_((_this).F12(), _574_v62, (_dafny.IntOfInt64(-247)).Times((_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int)))
				_575_v63 = _575_v63
				var _576_v65 _dafny.MultiSet
				_ = _576_v65
				_576_v65 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("igpmkqt")).Cardinality()), (_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int))
				var _577_v66 _dafny.Sequence
				_ = _577_v66
				_577_v66 = _dafny.SeqOf(Companion_Default___.Fm39((_this).F12(), (_dafny.Zero).Minus((_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int)), true, globalState), func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter33 := _dafny.Iterate((_576_v65).Elements()); ; {
						_compr_33, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _578_v64 _dafny.Int
						_578_v64 = interface{}(_compr_33).(_dafny.Int)
						if (_576_v65).Contains(_578_v64) {
							_coll33.Add((_578_v64).Times((_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int)), (_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll33.ToMap()
				}(), _489_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(90)))
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_568_v57), 0))
				_ = _index74
				(_568_v57).ArraySet1((false) || (true), (_index74).Int())
				var _579_v67 _dafny.Sequence
				_ = _579_v67
				_579_v67 = _dafny.SeqOf(_484_v0)
				var _580_v68 D5
				_ = _580_v68
				_580_v68 = Companion_D5_.Create_DC14_((_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int), (_this).F12())
				var _581_v69 _dafny.Set
				_ = _581_v69
				_581_v69 = _dafny.SetOf((_this).F12(), (_this).F12())
				var _582_v70 _dafny.Sequence
				_ = _582_v70
				_582_v70 = _dafny.SeqOf((_558_v49).Fm5((_this).F12(), (_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int), (_581_v69).Cardinality(), _501_v8, globalState))
				var _583_v71 _dafny.Sequence
				_ = _583_v71
				_583_v71 = _dafny.SeqOf(_582_v70)
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_568_v57), 0))
				_ = _index75
				var _rhs71 _dafny.Int = (_570_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))).Int()).(_dafny.Int)
				_ = _rhs71
				var _rhs72 _dafny.Sequence = Companion_Default___.Fm40(_579_v67, _580_v68, _dafny.IntOfUint32(((_this).F13()).Cardinality()), globalState)
				_ = _rhs72
				var _rhs73 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_582_v70, _582_v70), (_583_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-736), _dafny.IntOfUint32((_583_v71).Cardinality()))).Uint32()).(_dafny.Sequence))
				_ = _rhs73
				var _lhs50 _dafny.Array = _568_v57
				_ = _lhs50
				var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_568_v57), 0))
				_ = _lhs51
				_484_v0 = _rhs71
				_577_v66 = _rhs72
				(_lhs50).ArraySet1(_rhs73, (_lhs51).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen((_570_v58), 0))
				_ = _index76
				(_570_v58).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(135), _484_v0), (_index76).Int())
				var _584_v72 *C0
				_ = _584_v72
				var _nw72 *C0 = New_C0_()
				_ = _nw72
				_nw72.Ctor__()
				_584_v72 = _nw72
			}
			var _585_v73 _dafny.Array
			_ = _585_v73
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_8
			var _nw73 _dafny.Array
			_ = _nw73
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw73 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) bool = func(_586_i4 _dafny.Int) bool {
					return (_this).F12()
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw73 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw73).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw73).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_585_v73 = _nw73
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_585_v73), 0))
			_ = _index77
			(_585_v73).ArraySet1((_this).F12(), (_index77).Int())
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_585_v73), 0))
			_ = _index78
			(_585_v73).ArraySet1((_this).F12(), (_index78).Int())
			var _587_v74 *C0
			_ = _587_v74
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__()
			_587_v74 = _nw74
			var _588_v75 _dafny.Sequence
			_ = _588_v75
			_588_v75 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F13(), (Companion_Default___.SafeIndex(_484_v0, _dafny.IntOfUint32(((_this).F13()).Cardinality()))).Uint32(), p0)).Cardinality()), _484_v0)
			var _589_v76 _dafny.Map
			_ = _589_v76
			_589_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_v0, _588_v75)
			var _590_v77 _dafny.Map
			_ = _590_v77
			_590_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_588_v75, (func() _dafny.Sequence {
				if (_589_v76).Contains(_484_v0) {
					return (_589_v76).Get(_484_v0).(_dafny.Sequence)
				}
				return _588_v75
			})())).Cardinality()))
			_590_v77 = (_590_v77).Update((_this).F13(), (_484_v0).Plus(_dafny.IntOfUint32((_588_v75).Cardinality())))
			_484_v0 = (_dafny.Zero).Minus((_dafny.IntOfInt64(144)).Plus(_484_v0))
		} else {
			var _591_v78 _dafny.Sequence
			_ = _591_v78
			_591_v78 = _dafny.SeqOf((_this).F12(), (_this).F12())
			r1 = ((_591_v78).Select((Companion_Default___.SafeIndex(_484_v0, _dafny.IntOfUint32((_591_v78).Cardinality()))).Uint32()).(bool)) || (!_dafny.Companion_Sequence_.Contains((_this).F13(), p0))
			var _592_v79 D6
			_ = _592_v79
			_592_v79 = Companion_D6_.Create_DC18_(p0, _484_v0, (_484_v0).Minus(_484_v0))
			_592_v79 = _592_v79
			r1 = (func() bool {
				if (_485_v1).Contains((func() _dafny.Int {
					if (_489_v5).Contains(_484_v0) {
						return (_489_v5).Get(_484_v0).(_dafny.Int)
					}
					return _dafny.IntOfInt64(142)
				})()) {
					return (_485_v1).Get((func() _dafny.Int {
						if (_489_v5).Contains(_484_v0) {
							return (_489_v5).Get(_484_v0).(_dafny.Int)
						}
						return _dafny.IntOfInt64(142)
					})()).(bool)
				}
				return (_this).F12()
			})()
			var _593_v80 *C0
			_ = _593_v80
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__()
			_593_v80 = _nw75
			_492_v7 = _492_v7
		}
		var _594_v81 _dafny.Array
		_ = _594_v81
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_9
		var _nw76 _dafny.Array
		_ = _nw76
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw76 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Int = (func(_595_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_596_i5 _dafny.Int) _dafny.Int {
					return (_596_i5).Times(_595_v0)
				}
			})(_484_v0)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw76 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw76).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw76).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_594_v81 = _nw76
		var _597_v82 _dafny.Map
		_ = _597_v82
		_597_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_594_v81, _dafny.UnicodeSeqOfUtf8Bytes("lhau"))
		r0 = _597_v82
		r1 = _dafny.Companion_Sequence_.Contains((_this).F13(), (func() _dafny.CodePoint {
			if (_this).F12() {
				return p0
			}
			return p0
		})())
		return r0, r1
	}
}
func (_this *C1) F12() bool {
	{
		return _this._f12
	}
}
func (_this *C1) F13() _dafny.Sequence {
	{
		return _this._f13
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f9  _dafny.Map
	_f10 bool
	F14  bool
	_f15 _dafny.Array
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f9 = _dafny.EmptyMap
	_this._f10 = false
	_this.F14 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F9() _dafny.Map {
	return _this._f9
}
func (_this *C2) F10() bool {
	return _this._f10
}
func (_this *C2) Ctor__(f14 bool, f15 _dafny.Array, f9 _dafny.Map, f10 bool) {
	{
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C2) Fm0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Map {
			var _coll34 = _dafny.NewMapBuilder()
			_ = _coll34
			for _iter34 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnq")).Cardinality()), (_this).F9())).Keys().Elements()); ; {
				_compr_34, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _598_v0 _dafny.Int
				_598_v0 = interface{}(_compr_34).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnq")).Cardinality()), (_this).F9())).Contains(_598_v0) {
					_coll34.Add((_598_v0).Times(_dafny.IntOfInt64(688)), (Companion_D0_.Create_DC1_((_dafny.SetOf(_this.F14, (_this).F10())).Cardinality(), _this.F14, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(289))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}(func(_599_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('m')
					}))).Cardinality()))).Cardinality()))).Dtor_cf3())
				}
			}
			return _coll34.ToMap()
		}()).Cardinality()).Cmp(_dafny.IntOfInt64(301)) != 0
	}
}
func (_this *C2) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bnhgpf")).Cardinality()), Companion_D6_.Create_DC17_((_this).F10(), _this.F14, _dafny.IntOfInt64(55), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xy")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uyxnksw")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(255), _this.F14)).Cardinality()))).Cardinality()).Cmp(_dafny.IntOfInt64(-987)) != 0 {
			return ((_dafny.Zero).Minus(_dafny.IntOfInt64(-395))).Minus((_dafny.SetOf(_this.F14)).Cardinality())
		} else {
			return ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(963), _this.F14)).Cardinality())).Plus(_dafny.IntOfInt64(-169))
		}
	}
}
func (_this *C2) Fm41(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf(_this.F14, _this.F14, (_dafny.MultiSetOf(_dafny.IntOfInt64(196), _dafny.IntOfInt64(419), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kepfnihg")).Cardinality()), (func() _dafny.Map {
			var _coll35 = _dafny.NewMapBuilder()
			_ = _coll35
			for _iter35 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(411))).Elements()); ; {
				_compr_35, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _600_v0 _dafny.Int
				_600_v0 = interface{}(_compr_35).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(411))).Contains(_600_v0) {
					_coll35.Add((_600_v0).Minus(_dafny.IntOfInt64(396)), _dafny.IntOfInt64(585))
				}
			}
			return _coll35.ToMap()
		}()).Cardinality())).Cardinality())).Equals(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xeverqrey")).Cardinality()))), !((_dafny.IntOfInt64(289)).Cmp(_dafny.IntOfInt64(944)) >= 0))
	}
}
func (_this *C2) Fm42(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(24)
	}
}
func (_this *C2) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _601_v0 _dafny.MultiSet
		_ = _601_v0
		_601_v0 = _dafny.MultiSetOf(_this.F14)
		var _602_v1 _dafny.Sequence
		_ = _602_v1
		_602_v1 = _dafny.SeqOf(_this.F14)
		var _603_v2 _dafny.Int
		_ = _603_v2
		_603_v2 = _dafny.IntOfInt64(689)
		var _604_v3 _dafny.Map
		_ = _604_v3
		_604_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_601_v0).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_602_v1, (Companion_Default___.SafeIndex((_this).Fm42(_603_v2, _603_v2, _603_v2, globalState), _dafny.IntOfUint32((_602_v1).Cardinality()))).Uint32(), (_this).F10()))), _603_v2)
		var _605_v4 _dafny.MultiSet
		_ = _605_v4
		_605_v4 = _dafny.MultiSetOf(_603_v2)
		_604_v3 = (_604_v3).Update((func() bool {
			if false {
				return true
			}
			return _this.F14
		})(), (_dafny.Zero).Minus((func() _dafny.Int {
			if (_605_v4).Contains((_dafny.Zero).Minus(_603_v2)) {
				return (_605_v4).Multiplicity((_dafny.Zero).Minus(_603_v2))
			}
			return _603_v2
		})()))
		(_this).F14 = (_this).F10()
		var _606_v5 _dafny.Array
		_ = _606_v5
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_10
		var _nw77 _dafny.Array
		_ = _nw77
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw77 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.CodePoint = func(_607_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw77 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw77).ArraySet1CodePoint(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw77).ArraySet1CodePoint(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_606_v5 = _nw77
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))
		_ = _index79
		(_606_v5).ArraySet1CodePoint(_dafny.CodePoint('d'), (_index79).Int())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
		_ = _index80
		((_this).F15()).ArraySet1((_this).Fm0(_602_v1, (_this).F10(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-427))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}((func(_608_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_609_i1 _dafny.Int) _dafny.Int {
				return _608_v2
			}
		})(_603_v2)))).Cardinality()), _this.F14, globalState), (_index80).Int())
		var _610_v7 _dafny.Sequence
		_ = _610_v7
		_610_v7 = _dafny.SeqOf(_603_v2)
		var _611_v8 D6
		_ = _611_v8
		_611_v8 = Companion_D6_.Create_DC17_(_this.F14, (_this).F10(), _dafny.IntOfInt64(-155), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter36 := _dafny.Iterate((_610_v7).Elements()); ; {
				_compr_36, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _612_v6 _dafny.Int
				_612_v6 = interface{}(_compr_36).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_610_v7, _612_v6) {
					_coll36.Add(Companion_Default___.SafeDivisionInt(_612_v6, _603_v2), _this.F14)
				}
			}
			return _coll36.ToMap()
		}()).Cardinality()))), (_this).Fm1((_this).F9(), globalState))
		var _613_v9 D5
		_ = _613_v9
		_613_v9 = Companion_D5_.Create_DC13_((_this).F15())
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))
		_ = _index81
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
		_ = _index82
		var _rhs74 _dafny.Int = (_611_v8).Dtor_cf33()
		_ = _rhs74
		var _rhs75 _dafny.CodePoint = Companion_Default___.Fm43((_dafny.IntOfUint32((_dafny.SeqOf(_613_v9)).Cardinality())).Minus(_dafny.IntOfInt64(170)), globalState)
		_ = _rhs75
		var _rhs76 _dafny.Map = _604_v3
		_ = _rhs76
		var _rhs77 bool = !((_this).F10()) || ((_this).F10())
		_ = _rhs77
		var _rhs78 _dafny.Int = _603_v2
		_ = _rhs78
		var _lhs52 _dafny.Array = _606_v5
		_ = _lhs52
		var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))
		_ = _lhs53
		var _lhs54 _dafny.Array = (_this).F15()
		_ = _lhs54
		var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
		_ = _lhs55
		_603_v2 = _rhs74
		(_lhs52).ArraySet1CodePoint(_rhs75, (_lhs53).Int())
		_604_v3 = _rhs76
		(_lhs54).ArraySet1(_rhs77, (_lhs55).Int())
		_603_v2 = _rhs78
		var _614_i2 _dafny.Int
		_ = _614_i2
		_614_i2 = _dafny.Zero
		{
			for !(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)) {
				{
					if (_614_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_614_i2 = (_614_i2).Plus(_dafny.One)
					(globalState).F8 = (_603_v2).Cmp(_603_v2) >= 0
					r1 = ((_this).F10()) || ((_this).F10())
					_603_v2 = _603_v2
					var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
					_ = _index83
					((_this).F15()).ArraySet1((false) && (_this.F14), (_index83).Int())
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _615_v10 _dafny.Map
		_ = _615_v10
		_615_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_603_v2, ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))
		_615_v10 = Companion_Default___.Fm44(false, globalState)
		if ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool) {
			var _616_v11 _dafny.Sequence
			_ = _616_v11
			_616_v11 = _dafny.UnicodeSeqOfUtf8Bytes("wem")
			var _617_v12 _dafny.Sequence
			_ = _617_v12
			_617_v12 = _dafny.SeqOf(_616_v11)
			var _618_v13 *C1
			_ = _618_v13
			var _nw78 *C1 = New_C1_()
			_ = _nw78
			_nw78.Ctor__(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), (_617_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_616_v11).Cardinality()), _dafny.IntOfUint32((_617_v12).Cardinality()))).Uint32()).(_dafny.Sequence))
			_618_v13 = _nw78
			var _619_v14 _dafny.Array
			_ = _619_v14
			var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
			_ = _nw79
			_619_v14 = _nw79
			var _620_v15 _dafny.Map
			_ = _620_v15
			_620_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _602_v1)
			var _621_v16 _dafny.Array
			_ = _621_v16
			var _nwElement0_17 _dafny.Sequence = (_618_v13).F13()
			_ = _nwElement0_17
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(22))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_17, 0)
			(_nw80).ArraySet1(_616_v11, 1)
			(_nw80).ArraySet1(Companion_Default___.Fm45(!(_this.F14), _603_v2, (_618_v13).F12(), globalState), 2)
			(_nw80).ArraySet1(_616_v11, 3)
			(_nw80).ArraySet1(Companion_Default___.Fm45((_602_v1).Select((Companion_Default___.SafeIndex(_603_v2, _dafny.IntOfUint32((_602_v1).Cardinality()))).Uint32()).(bool), _603_v2, _this.F14, globalState), 4)
			(_nw80).ArraySet1((_618_v13).F13(), 5)
			(_nw80).ArraySet1(_616_v11, 6)
			(_nw80).ArraySet1(_616_v11, 7)
			(_nw80).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("cmnxwuk"), (Companion_Default___.SafeIndex(_603_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cmnxwuk")).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int())), 8)
			(_nw80).ArraySet1(_616_v11, 9)
			(_nw80).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(486))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_622_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_623_i3 _dafny.Int) _dafny.CodePoint {
					return (_622_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_622_v5), 0))).Int())
				}
			})(_606_v5))), 10)
			(_nw80).ArraySet1((_618_v13).F13(), 11)
			(_nw80).ArraySet1((_618_v13).F13(), 12)
			(_nw80).ArraySet1(_dafny.Companion_Sequence_.Update((_618_v13).F13(), (Companion_Default___.SafeIndex((_620_v15).Cardinality(), _dafny.IntOfUint32(((_618_v13).F13()).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int())), 13)
			(_nw80).ArraySet1((_618_v13).F13(), 14)
			(_nw80).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(202))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_624_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_625_i4 _dafny.Int) _dafny.CodePoint {
					return (_624_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_624_v5), 0))).Int())
				}
			})(_606_v5))), 15)
			(_nw80).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_626_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_627_i5 _dafny.Int) _dafny.CodePoint {
					return (_626_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_626_v5), 0))).Int())
				}
			})(_606_v5))), 16)
			(_nw80).ArraySet1(_616_v11, 17)
			(_nw80).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_616_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_618_v13).F13()).Cardinality()), _dafny.IntOfUint32((_616_v11).Cardinality()))).Uint32(), _dafny.CodePoint('o')), (Companion_Default___.SafeIndex(_603_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_616_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_618_v13).F13()).Cardinality()), _dafny.IntOfUint32((_616_v11).Cardinality()))).Uint32(), _dafny.CodePoint('o'))).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int())), (Companion_Default___.SafeIndex(_603_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_616_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_618_v13).F13()).Cardinality()), _dafny.IntOfUint32((_616_v11).Cardinality()))).Uint32(), _dafny.CodePoint('o')), (Companion_Default___.SafeIndex(_603_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_616_v11, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_618_v13).F13()).Cardinality()), _dafny.IntOfUint32((_616_v11).Cardinality()))).Uint32(), _dafny.CodePoint('o'))).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()))).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int())), 18)
			(_nw80).ArraySet1(_616_v11, 19)
			(_nw80).ArraySet1((_618_v13).F13(), 20)
			(_nw80).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}(func(_628_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), 21)
			_621_v16 = _nw80
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_619_v14), 0))
			_ = _index84
			(_619_v14).ArraySet1(_621_v16, (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_619_v14), 0))
			_ = _index85
			(_619_v14).ArraySet1(_621_v16, (_index85).Int())
			(globalState).F8 = (_this).Fm0(_dafny.Companion_Sequence_.Update(_602_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_602_v1).Cardinality()), _dafny.IntOfUint32((_602_v1).Cardinality()))).Uint32(), _this.F14), (_618_v13).F12(), _603_v2, _dafny.Companion_Sequence_.Equal(_602_v1, _602_v1), globalState)
			var _629_v17 *C1
			_ = _629_v17
			var _nw81 *C1 = New_C1_()
			_ = _nw81
			_nw81.Ctor__(!(_this.F14), (_618_v13).F13())
			_629_v17 = _nw81
			var _630_v18 D6
			_ = _630_v18
			_630_v18 = Companion_D6_.Create_DC18_((_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()), _603_v2, (_603_v2).Plus(_603_v2))
			var _rhs79 D6 = _630_v18
			_ = _rhs79
			var _rhs80 bool = !(_dafny.SetOf(_630_v18)).Contains(Companion_D6_.Create_DC18_(((_629_v17).F13()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_610_v7).Cardinality()), _dafny.IntOfUint32(((_629_v17).F13()).Cardinality()))).Uint32()).(_dafny.CodePoint), _603_v2, _603_v2))
			_ = _rhs80
			var _lhs56 *C2 = _this
			_ = _lhs56
			_630_v18 = _rhs79
			_lhs56.F14 = _rhs80
		} else {
			var _source14 D1 = Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("n")))
			_ = _source14
			if _source14.Is_DC3() {
				var _631___mcc_h0 _dafny.Sequence = _source14.Get_().(D1_DC3).Cf5
				_ = _631___mcc_h0
				var _632_cf5 _dafny.Sequence = _631___mcc_h0
				_ = _632_cf5
				var _rhs81 bool = (_this).F10()
				_ = _rhs81
				var _rhs82 bool = ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)
				_ = _rhs82
				var _rhs83 _dafny.Int = _603_v2
				_ = _rhs83
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				var _lhs58 *C2 = _this
				_ = _lhs58
				_lhs57.F8 = _rhs81
				_lhs58.F14 = _rhs82
				_603_v2 = _rhs83
				var _633_v19 *C1
				_ = _633_v19
				var _nw82 *C1 = New_C1_()
				_ = _nw82
				_nw82.Ctor__(_this.F14, _dafny.Companion_Sequence_.Update(_632_cf5, (Companion_Default___.SafeIndex(_603_v2, _dafny.IntOfUint32((_632_cf5).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int())))
				_633_v19 = _nw82
				var _634_v20 _dafny.Set
				_ = _634_v20
				_634_v20 = _dafny.SetOf(_603_v2, _603_v2)
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index86
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index87
				var _rhs84 bool = _this.F14
				_ = _rhs84
				var _rhs85 bool = (_634_v20).IsProperSubsetOf(_634_v20)
				_ = _rhs85
				var _lhs59 _dafny.Array = (_this).F15()
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _lhs60
				var _lhs61 _dafny.Array = (_this).F15()
				_ = _lhs61
				var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _lhs62
				(_lhs59).ArraySet1(_rhs84, (_lhs60).Int())
				(_lhs61).ArraySet1(_rhs85, (_lhs62).Int())
				var _635_v21 _dafny.Array
				_ = _635_v21
				var _nwElement0_18 bool = _this.F14
				_ = _nwElement0_18
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(23))
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_18, 0)
				(_nw83).ArraySet1((_this).F10(), 1)
				(_nw83).ArraySet1((_633_v19).F12(), 2)
				(_nw83).ArraySet1((_633_v19).F12(), 3)
				(_nw83).ArraySet1(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), 4)
				(_nw83).ArraySet1((_633_v19).F12(), 5)
				(_nw83).ArraySet1((_633_v19).Fm32(_603_v2, _dafny.CodePoint('i'), globalState), 6)
				(_nw83).ArraySet1(true, 7)
				(_nw83).ArraySet1(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), 8)
				(_nw83).ArraySet1(false, 9)
				(_nw83).ArraySet1(_this.F14, 10)
				(_nw83).ArraySet1(!((_633_v19).F12()), 11)
				(_nw83).ArraySet1(!((_this).F10()), 12)
				(_nw83).ArraySet1(!(true), 13)
				(_nw83).ArraySet1(_this.F14, 14)
				(_nw83).ArraySet1(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), 15)
				(_nw83).ArraySet1(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), 16)
				(_nw83).ArraySet1((_633_v19).F12(), 17)
				(_nw83).ArraySet1(_this.F14, 18)
				(_nw83).ArraySet1(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), 19)
				(_nw83).ArraySet1((_633_v19).F12(), 20)
				(_nw83).ArraySet1(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), 21)
				(_nw83).ArraySet1((_this).F10(), 22)
				_635_v21 = _nw83
				var _636_v22 _dafny.Map
				_ = _636_v22
				_636_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_633_v19).F12()), _635_v21)
				_636_v22 = ((func() _dafny.Map {
					if ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool) {
						return _636_v22
					}
					return _636_v22
				})()).Merge((_636_v22))
			} else if _source14.Is_DC4() {
				var _637___mcc_h1 bool = _source14.Get_().(D1_DC4).Cf6
				_ = _637___mcc_h1
				var _638___mcc_h2 bool = _source14.Get_().(D1_DC4).Cf7
				_ = _638___mcc_h2
				var _639___mcc_h3 _dafny.Int = _source14.Get_().(D1_DC4).Cf8
				_ = _639___mcc_h3
				var _640_cf8 _dafny.Int = _639___mcc_h3
				_ = _640_cf8
				var _641_cf7 bool = _638___mcc_h2
				_ = _641_cf7
				var _642_cf6 bool = _637___mcc_h1
				_ = _642_cf6
				var _643_v23 D4
				_ = _643_v23
				_643_v23 = Companion_D4_.Create_DC11_((_this).Fm1((_this).F9(), globalState), _dafny.SetOf(_640_cf8, _603_v2), _641_cf7)
				_640_cf8 = (_603_v2).Times((_643_v23).Dtor_cf16())
				var _rhs86 _dafny.Int = _dafny.IntOfInt64(-571)
				_ = _rhs86
				var _rhs87 bool = _641_cf7
				_ = _rhs87
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				_640_cf8 = _rhs86
				_lhs63.F8 = _rhs87
				var _644_v24 _dafny.Map
				_ = _644_v24
				_644_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_603_v2, Companion_Default___.SafeModuloInt((_dafny.SetOf(_640_cf8, _640_cf8)).Cardinality(), _640_cf8))
				_644_v24 = (_644_v24).Update((_this).Fm42(_603_v2, (_dafny.Zero).Minus(_dafny.IntOfInt64(-380)), _603_v2, globalState), Companion_Default___.SafeModuloInt(_640_cf8, _640_cf8))
				var _645_v25 _dafny.Array
				_ = _645_v25
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw84
				_645_v25 = _nw84
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_645_v25), 0))
				_ = _index88
				(_645_v25).ArraySet1(_640_cf8, (_index88).Int())
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_645_v25), 0))
				_ = _index89
				(_645_v25).ArraySet1((_this).Fm1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(_602_v1, !(_642_cf6), _603_v2, _641_cf7, globalState), _this.F14)).Update(true, _642_cf6), globalState), (_index89).Int())
			} else if _source14.Is_DC2() {
				var _646___mcc_h4 _dafny.Sequence = _source14.Get_().(D1_DC2).Cf4
				_ = _646___mcc_h4
				var _647_cf4 _dafny.Sequence = _646___mcc_h4
				_ = _647_cf4
				(globalState).F8 = (_603_v2).Cmp(_603_v2) >= 0
				var _648_v26 *C0
				_ = _648_v26
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__()
				_648_v26 = _nw85
				var _649_v27 _dafny.Array
				_ = _649_v27
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_11
				var _nw86 _dafny.Array
				_ = _nw86
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw86 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Int = (func(_650_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_651_i7 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_651_i7, (_dafny.Zero).Minus(_650_v2))
						}
					})(_603_v2)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw86 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw86).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw86).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_649_v27 = _nw86
				var _652_v28 _dafny.Map
				_ = _652_v28
				_652_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm1((_this).F9(), globalState), _dafny.IntOfInt64(-190))
				var _653_v29 _dafny.Sequence
				_ = _653_v29
				_653_v29 = _dafny.UnicodeSeqOfUtf8Bytes("rnvidwis")
				var _654_v30 _dafny.Map
				_ = _654_v30
				_654_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_653_v29, _603_v2)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_649_v27), 0))
				_ = _index90
				(_649_v27).ArraySet1(((func() _dafny.Int {
					if (_652_v28).Contains(_dafny.IntOfUint32((_653_v29).Cardinality())) {
						return (_652_v28).Get(_dafny.IntOfUint32((_653_v29).Cardinality())).(_dafny.Int)
					}
					return _603_v2
				})()).Minus((func() _dafny.Int {
					if (_654_v30).Contains(_653_v29) {
						return (_654_v30).Get(_653_v29).(_dafny.Int)
					}
					return _603_v2
				})()), (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_649_v27), 0))
				_ = _index91
				(_649_v27).ArraySet1(_603_v2, (_index91).Int())
				var _655_v31 _dafny.Map
				_ = _655_v31
				_655_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_603_v2).Cmp(_603_v2) != 0, ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))
				_655_v31 = (_655_v31).Update(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))
			} else {
				var _656___mcc_h5 D1 = _source14.Get_().(D1_DC5).Cf9
				_ = _656___mcc_h5
				var _657_cf9 D1 = _656___mcc_h5
				_ = _657_cf9
				_601_v0 = (_601_v0).Difference(_601_v0)
				var _658_v32 D0
				_ = _658_v32
				_658_v32 = Companion_D0_.Create_DC0_(_this.F14)
				(_this).F14 = ((func() D0 {
					if (_this).F10() {
						return _658_v32
					}
					return _658_v32
				})()).Dtor_cf0()
				var _659_v33 _dafny.Array
				_ = _659_v33
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw87
				_659_v33 = _nw87
				var _660_v34 _dafny.Set
				_ = _660_v34
				_660_v34 = _dafny.SetOf(_659_v33)
				var _661_v35 D8
				_ = _661_v35
				_661_v35 = Companion_D8_.Create_DC24_((_this).F10(), _660_v34, _603_v2)
				var _662_v36 _dafny.Sequence
				_ = _662_v36
				_662_v36 = _dafny.SeqOf(_661_v35)
				var _663_v37 _dafny.MultiSet
				_ = _663_v37
				_663_v37 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_662_v36, _dafny.Companion_Sequence_.Update(_662_v36, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.IntOfUint32((_662_v36).Cardinality()))).Uint32(), _661_v35)))
				_663_v37 = (_663_v37).Intersection(_663_v37)
				var _664_v38 bool
				_ = _664_v38
				var _665_v39 _dafny.Map
				_ = _665_v39
				var _666_v40 D7
				_ = _666_v40
				var _out42 bool
				_ = _out42
				var _out43 _dafny.Map
				_ = _out43
				var _out44 D7
				_ = _out44
				_out42, _out43, _out44 = (_this).M7(_dafny.IntOfInt64(511), _603_v2, globalState)
				_664_v38 = _out42
				_665_v39 = _out43
				_666_v40 = _out44
			}
			var _667_v41 _dafny.Array
			_ = _667_v41
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw88
			_667_v41 = _nw88
			var _668_v42 _dafny.Array
			_ = _668_v42
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_12
			var _nw89 _dafny.Array
			_ = _nw89
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw89 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) bool = func(_669_i8 _dafny.Int) bool {
					return ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw89 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw89).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw89).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_668_v42 = _nw89
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_667_v41), 0))
			_ = _index92
			(_667_v41).ArraySet1(_668_v42, (_index92).Int())
			var _670_v43 _dafny.Sequence
			_ = _670_v43
			_670_v43 = _dafny.UnicodeSeqOfUtf8Bytes("lueejueq")
			var _671_v44 _dafny.Array
			_ = _671_v44
			var _nwElement0_19 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm45(_this.F14, (_601_v0).Cardinality(), ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), globalState), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_603_v2), _dafny.IntOfUint32((Companion_Default___.Fm45(_this.F14, (_601_v0).Cardinality(), ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), globalState)).Cardinality()))).Uint32(), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()))
			_ = _nwElement0_19
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(8))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_19, 0)
			(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(690))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_672_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_673_i9 _dafny.Int) _dafny.CodePoint {
					return (_672_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_672_v5), 0))).Int())
				}
			})(_606_v5))), 1)
			(_nw90).ArraySet1(_670_v43, 2)
			(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_670_v43, _670_v43), 3)
			(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(584))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}((func(_674_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_675_i10 _dafny.Int) _dafny.CodePoint {
					return (_674_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_674_v5), 0))).Int())
				}
			})(_606_v5))), _670_v43), 4)
			(_nw90).ArraySet1(_670_v43, 5)
			(_nw90).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cafhh"), 6)
			(_nw90).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wgxv"), _670_v43), 7)
			_671_v44 = _nw90
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_671_v44), 0))
			_ = _index93
			(_671_v44).ArraySet1(_670_v43, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_667_v41), 0))
			_ = _index94
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_671_v44), 0))
			_ = _index95
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _index96
			var _rhs88 _dafny.Array = (func() _dafny.Array {
				if ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool) {
					return _668_v42
				}
				return _668_v42
			})()
			_ = _rhs88
			var _rhs89 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("j"), Companion_Default___.Fm45((_this).Fm0(_602_v1, _this.F14, _603_v2, false, globalState), _603_v2, _this.F14, globalState))
			_ = _rhs89
			var _rhs90 bool = (_this).F10()
			_ = _rhs90
			var _rhs91 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_670_v43).Cardinality()))).Plus((_603_v2).Plus(_603_v2))
			_ = _rhs91
			var _lhs64 _dafny.Array = _667_v41
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_667_v41), 0))
			_ = _lhs65
			var _lhs66 _dafny.Array = _671_v44
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_671_v44), 0))
			_ = _lhs67
			var _lhs68 _dafny.Array = (_this).F15()
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _lhs69
			(_lhs64).ArraySet1(_rhs88, (_lhs65).Int())
			(_lhs66).ArraySet1(_rhs89, (_lhs67).Int())
			(_lhs68).ArraySet1(_rhs90, (_lhs69).Int())
			_603_v2 = _rhs91
			if (func() bool {
				if (_615_v10).Contains(_603_v2) {
					return (_615_v10).Get(_603_v2).(bool)
				}
				return (_603_v2).Cmp(_603_v2) == 0
			})() {
				_613_v9 = _613_v9
				var _676_v45 D1
				_ = _676_v45
				_676_v45 = Companion_D1_.Create_DC3_(_670_v43)
				var _pat_let_tv9 = _670_v43
				_ = _pat_let_tv9
				_676_v45 = func(_pat_let8_0 D1) D1 {
					return func(_677_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let9_0 _dafny.Sequence) D1 {
							return func(_678_dt__update_hcf5_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_678_dt__update_hcf5_h0)
							}(_pat_let9_0)
						}(_pat_let_tv9)
					}(_pat_let8_0)
				}(Companion_D1_.Create_DC3_(_670_v43))
				var _679_v46 _dafny.Array
				_ = _679_v46
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw91
				_679_v46 = _nw91
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_679_v46), 0))
				_ = _index97
				(_679_v46).ArraySet1(_603_v2, (_index97).Int())
				var _680_v47 _dafny.Array
				_ = _680_v47
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_13
				var _nw92 _dafny.Array
				_ = _nw92
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw92 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) D1 = (func(_681_v2 _dafny.Int) func(_dafny.Int) D1 {
						return func(_682_i11 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_(_this.F14, (_this).F10(), _681_v2)
						}
					})(_603_v2)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw92 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw92).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw92).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_680_v47 = _nw92
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_680_v47), 0))
				_ = _index98
				(_680_v47).ArraySet1(Companion_D1_.Create_DC4_(!(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)), _this.F14, _603_v2), (_index98).Int())
				var _683_v48 _dafny.Sequence
				_ = _683_v48
				_683_v48 = _dafny.SeqOf(_668_v42, _dafny.ArrayCastTo((_667_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_667_v41), 0))).Int())))
				var _684_v49 D1
				_ = _684_v49
				_684_v49 = Companion_D1_.Create_DC4_(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool), (((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)) == ((func() bool {
					if ((_this).F9()).Contains(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)) {
						return ((_this).F9()).Get(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)).(bool)
					}
					return false
				})()), _603_v2)
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_667_v41), 0))
				_ = _index99
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_679_v46), 0))
				_ = _index100
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_680_v47), 0))
				_ = _index101
				var _rhs92 _dafny.Int = (_this).Fm1((_this).F9(), globalState)
				_ = _rhs92
				var _rhs93 _dafny.Array = (_683_v48).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()), (_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()), _dafny.CodePoint('r'))).Cardinality(), _dafny.IntOfUint32((_683_v48).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs93
				var _rhs94 _dafny.Int = _603_v2
				_ = _rhs94
				var _rhs95 D1 = _684_v49
				_ = _rhs95
				var _rhs96 _dafny.Int = _dafny.IntOfUint32((_602_v1).Cardinality())
				_ = _rhs96
				var _lhs70 _dafny.Array = _667_v41
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_667_v41), 0))
				_ = _lhs71
				var _lhs72 _dafny.Array = _679_v46
				_ = _lhs72
				var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_679_v46), 0))
				_ = _lhs73
				var _lhs74 _dafny.Array = _680_v47
				_ = _lhs74
				var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(636), _dafny.ArrayLen((_680_v47), 0))
				_ = _lhs75
				_603_v2 = _rhs92
				(_lhs70).ArraySet1(_rhs93, (_lhs71).Int())
				(_lhs72).ArraySet1(_rhs94, (_lhs73).Int())
				(_lhs74).ArraySet1(_rhs95, (_lhs75).Int())
				_603_v2 = _rhs96
				_603_v2 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_602_v1).Cardinality()), ((_dafny.Zero).Minus((_this).Fm42((_679_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_679_v46), 0))).Int()).(_dafny.Int), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32(((_671_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_671_v44), 0))).Int()).(_dafny.Sequence)).Cardinality())), (_679_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_679_v46), 0))).Int()).(_dafny.Int))).Cardinality(), _dafny.IntOfInt64(-430), globalState))).Plus(_603_v2))
				var _685_v50 bool
				_ = _685_v50
				var _686_v51 _dafny.Map
				_ = _686_v51
				var _687_v52 D7
				_ = _687_v52
				var _out45 bool
				_ = _out45
				var _out46 _dafny.Map
				_ = _out46
				var _out47 D7
				_ = _out47
				_out45, _out46, _out47 = (_this).M7(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ou")).Cardinality()), _dafny.IntOfUint32(((_671_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_671_v44), 0))).Int()).(_dafny.Sequence)).Cardinality())), _603_v2, globalState)
				_685_v50 = _out45
				_686_v51 = _out46
				_687_v52 = _out47
			} else {
				var _688_v53 _dafny.Set
				_ = _688_v53
				_688_v53 = _dafny.SetOf((_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()))
				var _689_v54 _dafny.Map
				_ = _689_v54
				_689_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_603_v2, _603_v2)
				var _690_v55 D1
				_ = _690_v55
				_690_v55 = Companion_D1_.Create_DC4_((_this).Fm0(_dafny.SeqOf(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)), (_this).F10(), _dafny.IntOfInt64(445), (_this).F10(), globalState), _this.F14, _603_v2)
				var _691_v56 D5
				_ = _691_v56
				_691_v56 = Companion_D5_.Create_DC14_(_603_v2, _this.F14)
				var _692_v57 _dafny.Array
				_ = _692_v57
				var _nwElement0_20 _dafny.Int = _603_v2
				_ = _nwElement0_20
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(12))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_20, 0)
				(_nw93).ArraySet1(_dafny.IntOfUint32((_610_v7).Cardinality()), 1)
				(_nw93).ArraySet1((_603_v2).Times(_603_v2), 2)
				(_nw93).ArraySet1(_603_v2, 3)
				(_nw93).ArraySet1(Companion_Default___.SafeModuloInt((_689_v54).Cardinality(), _603_v2), 4)
				(_nw93).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if _this.F14 {
						return _603_v2
					}
					return _603_v2
				})()), 5)
				(_nw93).ArraySet1(_603_v2, 6)
				(_nw93).ArraySet1(_603_v2, 7)
				(_nw93).ArraySet1((_690_v55).Dtor_cf8(), 8)
				(_nw93).ArraySet1(_603_v2, 9)
				(_nw93).ArraySet1((_691_v56).Dtor_cf25(), 10)
				(_nw93).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(364))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_693_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_694_i12 _dafny.Int) _dafny.CodePoint {
						return (_693_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_693_v5), 0))).Int())
					}
				})(_606_v5)))).Cardinality()), 11)
				_692_v57 = _nw93
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))
				_ = _index102
				(_692_v57).ArraySet1(_603_v2, (_index102).Int())
				var _695_v58 _dafny.Map
				_ = _695_v58
				_695_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))
				var _696_v59 _dafny.Map
				_ = _696_v59
				_696_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))).Cardinality(), _695_v58)
				var _697_v61 _dafny.MultiSet
				_ = _697_v61
				_697_v61 = _dafny.MultiSetOf((_this).F9())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))
				_ = _index103
				var _rhs97 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_603_v2), Companion_Default___.SafeDivisionInt(((func() _dafny.Map {
					if (_696_v59).Contains(_603_v2) {
						return (_696_v59).Get(_603_v2).(_dafny.Map)
					}
					return func() _dafny.Map {
						var _coll37 = _dafny.NewMapBuilder()
						_ = _coll37
						for _iter37 := _dafny.Iterate((_697_v61).Elements()); ; {
							_compr_37, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _698_v60 _dafny.Map
							_698_v60 = interface{}(_compr_37).(_dafny.Map)
							if (_697_v61).Contains(_698_v60) {
								_coll37.Add(_698_v60, ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))
							}
						}
						return _coll37.ToMap()
					}()
				})()).Cardinality(), _603_v2))
				_ = _rhs97
				var _rhs98 _dafny.Set = ((_688_v53).Difference(_688_v53)).Union((_dafny.SetOf((_606_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_606_v5), 0))).Int()))).Difference(func() _dafny.Set {
					var _coll38 = _dafny.NewBuilder()
					_ = _coll38
					for _iter38 := _dafny.Iterate((_670_v43).Elements()); ; {
						_compr_38, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _699_v62 _dafny.CodePoint
						_699_v62 = interface{}(_compr_38).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_670_v43, _699_v62) {
							_coll38.Add(_699_v62)
						}
					}
					return _coll38.ToSet()
				}()))
				_ = _rhs98
				var _rhs99 _dafny.Int = _603_v2
				_ = _rhs99
				var _rhs100 _dafny.Int = (_dafny.Zero).Minus((_601_v0).Cardinality())
				_ = _rhs100
				var _lhs76 _dafny.Array = _692_v57
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))
				_ = _lhs77
				_603_v2 = _rhs97
				_688_v53 = _rhs98
				_603_v2 = _rhs99
				(_lhs76).ArraySet1(_rhs100, (_lhs77).Int())
				r0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(67))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_700_v5 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_701_i13 _dafny.Int) _dafny.CodePoint {
						return (_700_v5).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_700_v5), 0))).Int())
					}
				})(_606_v5)))
				var _702_v63 D8
				_ = _702_v63
				_702_v63 = Companion_D8_.Create_DC24_(!(((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool)), _dafny.SetOf(_692_v57), (_692_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))).Int()).(_dafny.Int))
				var _703_v64 _dafny.Sequence
				_ = _703_v64
				_703_v64 = _dafny.SeqOf(_702_v63, _702_v63, _702_v63)
				var _704_v65 _dafny.Set
				_ = _704_v65
				_704_v65 = _dafny.SetOf(_692_v57)
				var _705_v66 _dafny.Map
				_ = _705_v66
				_705_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_692_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))).Int()).(_dafny.Int), _703_v64)
				var _706_v67 _dafny.Array
				_ = _706_v67
				var _nwElement0_21 _dafny.Sequence = _703_v64
				_ = _nwElement0_21
				var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(21))
				_ = _nw94
				(_nw94).ArraySet1(_nwElement0_21, 0)
				(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_703_v64, _dafny.SeqOf(_702_v63)), 1)
				(_nw94).ArraySet1(_703_v64, 2)
				(_nw94).ArraySet1(_703_v64, 3)
				(_nw94).ArraySet1(_703_v64, 4)
				(_nw94).ArraySet1(_703_v64, 5)
				(_nw94).ArraySet1(_dafny.SeqOf(_702_v63, _702_v63, _702_v63, Companion_D8_.Create_DC24_(_this.F14, _704_v65, (_692_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))).Int()).(_dafny.Int)), _702_v63), 6)
				(_nw94).ArraySet1(_dafny.SeqOf(_702_v63, _702_v63), 7)
				(_nw94).ArraySet1((func() _dafny.Sequence {
					if true {
						return _703_v64
					}
					return _703_v64
				})(), 8)
				(_nw94).ArraySet1(_703_v64, 9)
				(_nw94).ArraySet1((func() _dafny.Sequence {
					if ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool) {
						return _dafny.SeqOf(_702_v63, _702_v63, _702_v63)
					}
					return _703_v64
				})(), 10)
				(_nw94).ArraySet1(_dafny.Companion_Sequence_.Update(_703_v64, (Companion_Default___.SafeIndex((_692_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_692_v57), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_703_v64).Cardinality()))).Uint32(), _702_v63), 11)
				(_nw94).ArraySet1(_703_v64, 12)
				(_nw94).ArraySet1(_703_v64, 13)
				(_nw94).ArraySet1(_703_v64, 14)
				(_nw94).ArraySet1(_dafny.SeqOf(_702_v63, _702_v63, _702_v63), 15)
				(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_703_v64, _dafny.SeqOf(_702_v63, _702_v63)), 16)
				(_nw94).ArraySet1(_703_v64, 17)
				(_nw94).ArraySet1(_703_v64, 18)
				(_nw94).ArraySet1((func() _dafny.Sequence {
					if (_705_v66).Contains(_603_v2) {
						return (_705_v66).Get(_603_v2).(_dafny.Sequence)
					}
					return _703_v64
				})(), 19)
				(_nw94).ArraySet1(_703_v64, 20)
				_706_v67 = _nw94
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_706_v67), 0))
				_ = _index104
				(_706_v67).ArraySet1(_703_v64, (_index104).Int())
				var _707_v68 _dafny.Map
				_ = _707_v68
				_707_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_601_v0, _dafny.SeqOf(_702_v63))
				var _pat_let_tv10 = _602_v1
				_ = _pat_let_tv10
				var _pat_let_tv11 = _605_v4
				_ = _pat_let_tv11
				var _pat_let_tv12 = globalState
				_ = _pat_let_tv12
				var _pat_let_tv13 = _602_v1
				_ = _pat_let_tv13
				var _pat_let_tv14 = _605_v4
				_ = _pat_let_tv14
				var _pat_let_tv15 = globalState
				_ = _pat_let_tv15
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_706_v67), 0))
				_ = _index105
				(_706_v67).ArraySet1((func() _dafny.Sequence {
					if (_707_v68).Contains((_601_v0).Update((func(_pat_let10_0 D1) D1 {
						return func(_708_dt__update__tmp_h1 D1) D1 {
							return func(_pat_let11_0 bool) D1 {
								return func(_709_dt__update_hcf7_h0 bool) D1 {
									return func(_pat_let12_0 bool) D1 {
										return func(_710_dt__update_hcf6_h0 bool) D1 {
											return Companion_D1_.Create_DC4_(_710_dt__update_hcf6_h0, _709_dt__update_hcf7_h0, (_708_dt__update__tmp_h1).Dtor_cf8())
										}(_pat_let12_0)
									}((_this).Fm0(_pat_let_tv10, (_this).F10(), (_pat_let_tv11).Cardinality(), _this.F14, _pat_let_tv12))
								}(_pat_let11_0)
							}((_this).F10())
						}(_pat_let10_0)
					}(_690_v55)).Dtor_cf6(), Companion_Default___.Abs((_this).Fm1((_this).F9(), globalState)))) {
						return (_707_v68).Get((_601_v0).Update((func(_pat_let13_0 D1) D1 {
							return func(_711_dt__update__tmp_h2 D1) D1 {
								return func(_pat_let14_0 bool) D1 {
									return func(_712_dt__update_hcf7_h1 bool) D1 {
										return func(_pat_let15_0 bool) D1 {
											return func(_713_dt__update_hcf6_h1 bool) D1 {
												return Companion_D1_.Create_DC4_(_713_dt__update_hcf6_h1, _712_dt__update_hcf7_h1, (_711_dt__update__tmp_h2).Dtor_cf8())
											}(_pat_let15_0)
										}((_this).Fm0(_pat_let_tv13, (_this).F10(), (_pat_let_tv14).Cardinality(), _this.F14, _pat_let_tv15))
									}(_pat_let14_0)
								}((_this).F10())
							}(_pat_let13_0)
						}(_690_v55)).Dtor_cf6(), Companion_Default___.Abs((_this).Fm1((_this).F9(), globalState)))).(_dafny.Sequence)
					}
					return _703_v64
				})(), (_index105).Int())
				var _714_v69 _dafny.Array
				_ = _714_v69
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(7))
				_ = _nw95
				_714_v69 = _nw95
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_714_v69), 0))
				_ = _index106
				(_714_v69).ArraySet1(_702_v63, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_714_v69), 0))
				_ = _index107
				(_714_v69).ArraySet1(_702_v63, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
				_ = _index108
				((_this).F15()).ArraySet1(!(_dafny.Companion_Sequence_.Contains(_602_v1, ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool))), (_index108).Int())
			}
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _index109
			var _rhs101 bool = (_this).F10()
			_ = _rhs101
			var _rhs102 _dafny.Int = (_dafny.IntOfInt64(176)).Minus(((_this).F9()).Cardinality())
			_ = _rhs102
			var _lhs78 _dafny.Array = (_this).F15()
			_ = _lhs78
			var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _lhs79
			(_lhs78).ArraySet1(_rhs101, (_lhs79).Int())
			_603_v2 = _rhs102
			_603_v2 = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_603_v2, _603_v2)).Minus((func() _dafny.Int {
				if ((_this).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_this).F15()), 0))).Int()).(bool) {
					return _603_v2
				}
				return _603_v2
			})()))
		}
		var _715_v70 _dafny.Sequence
		_ = _715_v70
		_715_v70 = _dafny.UnicodeSeqOfUtf8Bytes("ysxcsgfq")
		r0 = _715_v70
		r1 = false
		var _716_v71 _dafny.Set
		_ = _716_v71
		_716_v71 = _dafny.SetOf(_this.F14)
		var _717_v72 _dafny.Sequence
		_ = _717_v72
		_717_v72 = _dafny.SeqOf((_716_v71).Union(_716_v71))
		r2 = _717_v72
		r3 = _dafny.Companion_Sequence_.Update(_715_v70, (Companion_Default___.SafeIndex(((_716_v71).Cardinality()).Times(_603_v2), _dafny.IntOfUint32((_715_v70).Cardinality()))).Uint32(), _dafny.CodePoint('v'))
		return r0, r1, r2, r3
	}
}
func (_this *C2) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _718_v0 _dafny.Int
		_ = _718_v0
		_718_v0 = _dafny.IntOfInt64(125)
		_718_v0 = _dafny.IntOfInt64(857)
		if (_this.F14) && (!(_this.F14) || (_this.F14)) {
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _index110
			((_this).F15()).ArraySet1((_dafny.IntOfInt64(910)).Cmp(_718_v0) > 0, (_index110).Int())
			var _719_v1 _dafny.Sequence
			_ = _719_v1
			_719_v1 = _dafny.SeqOf((true) == ((_this).F10()), false, false)
			var _720_v2 _dafny.MultiSet
			_ = _720_v2
			_720_v2 = _dafny.MultiSetOf(_718_v0)
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _index111
			var _rhs103 bool = true
			_ = _rhs103
			var _rhs104 _dafny.Sequence = _719_v1
			_ = _rhs104
			var _rhs105 bool = (_this).F10()
			_ = _rhs105
			var _rhs106 _dafny.Int = (_this).Fm42(_718_v0, ((_dafny.Zero).Minus(_718_v0)).Plus(_718_v0), (_720_v2).Cardinality(), globalState)
			_ = _rhs106
			var _lhs80 _dafny.Array = (_this).F15()
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen(((_this).F15()), 0))
			_ = _lhs81
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			(_lhs80).ArraySet1(_rhs103, (_lhs81).Int())
			_719_v1 = _rhs104
			_lhs82.F8 = _rhs105
			_718_v0 = _rhs106
			_718_v0 = (_718_v0).Minus(_718_v0)
			var _721_v3 _dafny.Sequence
			_ = _721_v3
			_721_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ll")
			var _722_v4 _dafny.Array
			_ = _722_v4
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_14
			var _nw96 _dafny.Array
			_ = _nw96
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw96 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = (func(_723_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_724_i0 _dafny.Int) _dafny.Int {
						return (_724_i0).Times(_723_v0)
					}
				})(_718_v0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw96 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw96).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw96).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_722_v4 = _nw96
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_722_v4), 0))
			_ = _index112
			(_722_v4).ArraySet1((_dafny.Zero).Minus(_718_v0), (_index112).Int())
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_722_v4), 0))
			_ = _index113
			var _rhs107 _dafny.Int = _718_v0
			_ = _rhs107
			var _rhs108 bool = _this.F14
			_ = _rhs108
			var _rhs109 bool = (_this).F10()
			_ = _rhs109
			var _rhs110 _dafny.Sequence = _721_v3
			_ = _rhs110
			var _rhs111 _dafny.Int = ((_718_v0).Times(_dafny.IntOfInt64(840))).Times(Companion_Default___.SafeModuloInt(_718_v0, _718_v0))
			_ = _rhs111
			var _lhs83 *C2 = _this
			_ = _lhs83
			var _lhs84 *GlobalState = globalState
			_ = _lhs84
			var _lhs85 _dafny.Array = _722_v4
			_ = _lhs85
			var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_722_v4), 0))
			_ = _lhs86
			_718_v0 = _rhs107
			_lhs83.F14 = _rhs108
			_lhs84.F8 = _rhs109
			_721_v3 = _rhs110
			(_lhs85).ArraySet1(_rhs111, (_lhs86).Int())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_722_v4), 0))
			_ = _index114
			(_722_v4).ArraySet1(((_722_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_722_v4), 0))).Int()).(_dafny.Int)).Minus(_718_v0), (_index114).Int())
			var _725_v5 *C0
			_ = _725_v5
			var _nw97 *C0 = New_C0_()
			_ = _nw97
			_nw97.Ctor__()
			_725_v5 = _nw97
		} else {
			var _726_v6 _dafny.Array
			_ = _726_v6
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_15
			var _nw98 _dafny.Array
			_ = _nw98
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw98 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = (func(_727_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_728_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_728_i1, (Companion_D5_.Create_DC14_(_727_v0, (_this).F10())).Dtor_cf25())
					}
				})(_718_v0)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw98 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw98).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw98).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_726_v6 = _nw98
			_726_v6 = _726_v6
			var _729_v7 _dafny.Sequence
			_ = _729_v7
			_729_v7 = _dafny.SeqOf(_718_v0, _718_v0, _dafny.IntOfInt64(858))
			var _730_v8 _dafny.Map
			_ = _730_v8
			_730_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, Companion_Default___.SafeDivisionInt(_718_v0, _dafny.IntOfUint32((_729_v7).Cardinality())))
			_730_v8 = (_730_v8).Update(true, (_dafny.Zero).Minus((_this).Fm1((_this).F9(), globalState)))
			var _731_v9 _dafny.Array
			_ = _731_v9
			var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw99
			_731_v9 = _nw99
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))
			_ = _index115
			(_731_v9).ArraySet1(_726_v6, (_index115).Int())
			var _732_v10 D3
			_ = _732_v10
			_732_v10 = Companion_D3_.Create_DC7_(_726_v6)
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))
			_ = _index116
			(_731_v9).ArraySet1((_732_v10).Dtor_cf11(), (_index116).Int())
			var _733_v11 _dafny.Sequence
			_ = _733_v11
			_733_v11 = _dafny.SeqOf(true)
			var _734_v12 D6
			_ = _734_v12
			_734_v12 = Companion_D6_.Create_DC18_(p0, _dafny.IntOfUint32((_733_v11).Cardinality()), _718_v0)
			var _735_v13 D6
			_ = _735_v13
			_735_v13 = Companion_D6_.Create_DC20_(Companion_D6_.Create_DC20_(_734_v12))
			var _736_v14 _dafny.Set
			_ = _736_v14
			_736_v14 = _dafny.SetOf(func(_pat_let16_0 D6) D6 {
				return func(_737_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let17_0 D6) D6 {
						return func(_738_dt__update_hcf39_h0 D6) D6 {
							return Companion_D6_.Create_DC20_(_738_dt__update_hcf39_h0)
						}(_pat_let17_0)
					}(Companion_D6_.Create_DC16_((_this).F9()))
				}(_pat_let16_0)
			}(_735_v13))
			var _739_v15 _dafny.Sequence
			_ = _739_v15
			_739_v15 = _dafny.SeqOf(_736_v14)
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))
			_ = _arr0
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))), 0))
			_ = _index117
			(_arr0).ArraySet1(_718_v0, (_index117).Int())
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))
			_ = _arr1
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))), 0))
			_ = _index118
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))
			_ = _index119
			var _rhs112 _dafny.Int = ((_dafny.Zero).Minus((_718_v0).Times(_718_v0))).Times(_718_v0)
			_ = _rhs112
			var _rhs113 _dafny.Sequence = _739_v15
			_ = _rhs113
			var _rhs114 _dafny.Int = (_dafny.IntOfInt64(195)).Times(_718_v0)
			_ = _rhs114
			var _rhs115 _dafny.Array = (_732_v10).Dtor_cf11()
			_ = _rhs115
			var _lhs87 _dafny.Array = _dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))
			_ = _lhs87
			var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))), 0))
			_ = _lhs88
			var _lhs89 _dafny.Array = _731_v9
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))
			_ = _lhs90
			_718_v0 = _rhs112
			_739_v15 = _rhs113
			(_lhs87).ArraySet1(_rhs114, (_lhs88).Int())
			(_lhs89).ArraySet1(_rhs115, (_lhs90).Int())
			(globalState).F8 = (_718_v0).Cmp((func() _dafny.Int {
				if (_this).F10() {
					return (_dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_dafny.ArrayCastTo((_731_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_731_v9), 0))).Int()))), 0))).Int()).(_dafny.Int)
				}
				return (Companion_Default___.Fm46((_this).F10(), (_this).F10(), (_this).F10(), globalState)).Cardinality()
			})()) < 0
		}
		var _740_v16 _dafny.Array
		_ = _740_v16
		var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(18))
		_ = _nw100
		_740_v16 = _nw100
		var _741_v17 _dafny.Set
		_ = _741_v17
		_741_v17 = _dafny.SetOf((_this).F10())
		var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_740_v16), 0))
		_ = _index120
		(_740_v16).ArraySet1(_741_v17, (_index120).Int())
		var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_740_v16), 0))
		_ = _index121
		(_740_v16).ArraySet1(_741_v17, (_index121).Int())
		var _742_v18 _dafny.Array
		_ = _742_v18
		var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
		_ = _nw101
		_742_v18 = _nw101
		for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_742_v18), 0))); ; {
			_guard_loop_0, _ok39 := _iter39()
			if !_ok39 {
				break
			}
			var _743_i2 _dafny.Int
			_743_i2 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_743_i2).Sign() != -1) && ((_743_i2).Cmp(_dafny.ArrayLen((_742_v18), 0)) < 0)) {
				(_742_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_718_v0)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("upbj")).Cardinality()), _718_v0), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kmewrgjgd")).Cardinality()), _718_v0, _dafny.IntOfInt64(-56))), (_743_i2).Int())
			}
		}
		r1 = _this.F14
		var _744_v19 _dafny.Map
		_ = _744_v19
		_744_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _718_v0)
		var _745_v20 _dafny.CodePoint
		_ = _745_v20
		_745_v20 = p0
		var _746_v21 _dafny.Map
		_ = _746_v21
		_746_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_745_v20, _744_v19)
		_744_v19 = (func() _dafny.Map {
			if (_746_v21).Contains(p0) {
				return (_746_v21).Get(p0).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _dafny.IntOfInt64(39))
		})()
		var _747_v22 _dafny.Array
		_ = _747_v22
		var _nwElement0_22 _dafny.Int = _718_v0
		_ = _nwElement0_22
		var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(7))
		_ = _nw102
		(_nw102).ArraySet1(_nwElement0_22, 0)
		(_nw102).ArraySet1(_718_v0, 1)
		(_nw102).ArraySet1(_718_v0, 2)
		(_nw102).ArraySet1(_dafny.IntOfInt64(-454), 3)
		(_nw102).ArraySet1(_718_v0, 4)
		(_nw102).ArraySet1(_718_v0, 5)
		(_nw102).ArraySet1(_718_v0, 6)
		_747_v22 = _nw102
		var _748_v23 _dafny.Sequence
		_ = _748_v23
		_748_v23 = _dafny.UnicodeSeqOfUtf8Bytes("by")
		var _749_v24 _dafny.Map
		_ = _749_v24
		_749_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_747_v22, _748_v23)
		r0 = (_749_v24).Update(_747_v22, _748_v23)
		var _750_v25 D1
		_ = _750_v25
		_750_v25 = Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(626))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}((func(_751_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_752_i4 _dafny.Int) _dafny.CodePoint {
				return _751_p0
			}
		})(p0))))
		r1 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(322))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}((func(_753_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_754_i3 _dafny.Int) _dafny.CodePoint {
				return _753_p0
			}
		})(p0))), (_750_v25).Dtor_cf5())
		return r0, r1
	}
}
func (_this *C2) M7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Map, D7) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 D7 = Companion_D7_.Default()
		_ = r2
		r0 = (_this).F10()
		var _755_v0 *C0
		_ = _755_v0
		var _nw103 *C0 = New_C0_()
		_ = _nw103
		_nw103.Ctor__()
		_755_v0 = _nw103
		var _756_i0 _dafny.Int
		_ = _756_i0
		_756_i0 = _dafny.Zero
		{
			for !(_this.F14) {
				{
					if (_756_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_756_i0 = (_756_i0).Plus(_dafny.One)
					var _757_v1 _dafny.Int
					_ = _757_v1
					_757_v1 = _dafny.IntOfInt64(891)
					_757_v1 = p0
					var _758_v2 _dafny.CodePoint
					_ = _758_v2
					_758_v2 = _dafny.CodePoint('o')
					var _rhs116 _dafny.CodePoint = _758_v2
					_ = _rhs116
					var _rhs117 _dafny.Int = (_757_v1).Times(p1)
					_ = _rhs117
					var _rhs118 _dafny.Int = p0
					_ = _rhs118
					_758_v2 = _rhs116
					_757_v1 = _rhs117
					_757_v1 = _rhs118
					_757_v1 = (_dafny.Zero).Minus(_757_v1)
					_757_v1 = (_dafny.Zero).Minus((_this).Fm42(p0, p0, p0, globalState))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _759_v3 _dafny.Sequence
		_ = _759_v3
		_759_v3 = _dafny.SeqOf(p0)
		var _760_v4 _dafny.Map
		_ = _760_v4
		_760_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_759_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.IntOfUint32((_759_v3).Cardinality()))).Uint32()).(_dafny.Int))
		var _761_v5 _dafny.Map
		_ = _761_v5
		_761_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, _760_v4)
		var _762_v6 _dafny.Map
		_ = _762_v6
		_762_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14, p0)
		var _763_v7 _dafny.MultiSet
		_ = _763_v7
		_763_v7 = _dafny.MultiSetOf(p1)
		var _764_v8 _dafny.Map
		_ = _764_v8
		_764_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_763_v7).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nixblma")).Cardinality())))
		_761_v5 = (_761_v5).Update((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0))).Contains(_762_v6), _764_v8)
		r0 = (((_this).F10()) || (_this.F14)) == (_this.F14)
		var _765_v9 _dafny.Array
		_ = _765_v9
		var _nwElement0_23 bool = _this.F14
		_ = _nwElement0_23
		var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
		_ = _nw104
		(_nw104).ArraySet1(_nwElement0_23, 0)
		(_nw104).ArraySet1((func() bool {
			if (_this).F10() {
				return (_this).F10()
			}
			return _this.F14
		})(), 1)
		(_nw104).ArraySet1((_this).F10(), 2)
		(_nw104).ArraySet1(_this.F14, 3)
		(_nw104).ArraySet1(_this.F14, 4)
		(_nw104).ArraySet1(_this.F14, 5)
		(_nw104).ArraySet1((func() bool {
			if (_this).F10() {
				return (_this).F10()
			}
			return (_this).F10()
		})(), 6)
		(_nw104).ArraySet1((_this).F10(), 7)
		(_nw104).ArraySet1((_this).F10(), 8)
		(_nw104).ArraySet1((_dafny.IntOfInt64(708)).Cmp(_dafny.IntOfInt64(-800)) > 0, 9)
		(_nw104).ArraySet1(!((_this).F10()) || (_this.F14), 10)
		(_nw104).ArraySet1(!((_this).F10()) || (_this.F14), 11)
		(_nw104).ArraySet1(_this.F14, 12)
		_765_v9 = _nw104
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_765_v9), 0))); ; {
			_guard_loop_1, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _766_i1 _dafny.Int
			_766_i1 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_766_i1).Sign() != -1) && ((_766_i1).Cmp(_dafny.ArrayLen((_765_v9), 0)) < 0)) {
				(_765_v9).ArraySet1(!(((func() _dafny.Set {
					var _coll39 = _dafny.NewBuilder()
					_ = _coll39
					for _iter41 := _dafny.Iterate((_dafny.SetOf(p1)).Elements()); ; {
						_compr_39, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _767_v10 _dafny.Int
						_767_v10 = interface{}(_compr_39).(_dafny.Int)
						if (_dafny.SetOf(p1)).Contains(_767_v10) {
							_coll39.Add((_767_v10).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(true))), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg59 _dafny.Int) interface{} {
									return coer59(arg59)
								}
							}(func(_768_i2 _dafny.Int) _dafny.Int {
								return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality()
							}))).Cardinality()))))))).Cardinality()))
						}
					}
					return _coll39.ToSet()
				}()).Cardinality()).Cmp(p0) == 0) || ((_this).F10()), (_766_i1).Int())
			}
		}
		r0 = _this.F14
		r1 = _762_v6
		var _769_v11 D7
		_ = _769_v11
		_769_v11 = Companion_D7_.Create_DC22_()
		r2 = _769_v11
		return r0, r1, r2
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
func (_this *C3) Fm27(p0 bool, p1 bool, p2 D3, globalState *GlobalState) bool {
	{
		return (!(false)) && ((_dafny.IntOfInt64(-159)).Cmp(_dafny.IntOfInt64(821)) == 0)
	}
}
func (_this *C3) Fm28(p0 bool, p1 D1, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (func(_source15 D1) _dafny.Map {
			if _source15.Is_DC3() {
				var _770___mcc_h0 _dafny.Sequence = _source15.Get_().(D1_DC3).Cf5
				_ = _770___mcc_h0
				var _771_cf5 _dafny.Sequence = _770___mcc_h0
				_ = _771_cf5
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false))
			} else if _source15.Is_DC4() {
				var _772___mcc_h1 bool = _source15.Get_().(D1_DC4).Cf6
				_ = _772___mcc_h1
				var _773___mcc_h2 bool = _source15.Get_().(D1_DC4).Cf7
				_ = _773___mcc_h2
				var _774___mcc_h3 _dafny.Int = _source15.Get_().(D1_DC4).Cf8
				_ = _774___mcc_h3
				var _775_cf8 _dafny.Int = _774___mcc_h3
				_ = _775_cf8
				var _776_cf7 bool = _773___mcc_h2
				_ = _776_cf7
				var _777_cf6 bool = _772___mcc_h1
				_ = _777_cf6
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), !(_776_cf7))
			} else if _source15.Is_DC2() {
				var _778___mcc_h4 _dafny.Sequence = _source15.Get_().(D1_DC2).Cf4
				_ = _778___mcc_h4
				var _779_cf4 _dafny.Sequence = _778___mcc_h4
				_ = _779_cf4
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), true))
			} else {
				var _780___mcc_h5 D1 = _source15.Get_().(D1_DC5).Cf9
				_ = _780___mcc_h5
				var _781_cf9 D1 = _780___mcc_h5
				_ = _781_cf9
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false))
			}
		}(Companion_D1_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(725))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg60 _dafny.Int) interface{} {
				return coer60(arg60)
			}
		}(func(_782_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))))).Cardinality()
	}
}
func (_this *C3) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _783_v0 bool
		_ = _783_v0
		_783_v0 = true
		r1 = !(_783_v0) || (_783_v0)
		var _784_v1 _dafny.Array
		_ = _784_v1
		var _nw105 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw105
		_784_v1 = _nw105
		var _785_v2 _dafny.Map
		_ = _785_v2
		_785_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_784_v1, _783_v0)
		_785_v2 = (_785_v2).Update(_784_v1, (_this).Fm27(_783_v0, _783_v0, Companion_D3_.Create_DC8_(_783_v0, _783_v0), globalState))
		var _786_v3 _dafny.Int
		_ = _786_v3
		_786_v3 = _dafny.IntOfInt64(838)
		var _787_v4 _dafny.CodePoint
		_ = _787_v4
		_787_v4 = _dafny.CodePoint('q')
		var _788_v5 _dafny.Sequence
		_ = _788_v5
		_788_v5 = _dafny.UnicodeSeqOfUtf8Bytes("gkun")
		var _789_v6 D1
		_ = _789_v6
		_789_v6 = Companion_D1_.Create_DC3_(_788_v5)
		var _790_v7 _dafny.Map
		_ = _790_v7
		_790_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v6, _786_v3)
		var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
		_ = _index122
		(_784_v1).ArraySet1(((_790_v7).Update(_789_v6, _786_v3)).Cardinality(), (_index122).Int())
		var _791_v8 D3
		_ = _791_v8
		_791_v8 = Companion_D3_.Create_DC8_(_783_v0, _783_v0)
		var _792_v9 _dafny.MultiSet
		_ = _792_v9
		_792_v9 = _dafny.MultiSetOf((_this).Fm27(_783_v0, _783_v0, _791_v8, globalState))
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
		_ = _index123
		var _rhs119 _dafny.Int = ((_792_v9).Difference(_dafny.MultiSetOf(true))).Cardinality()
		_ = _rhs119
		var _rhs120 bool = _783_v0
		_ = _rhs120
		var _rhs121 _dafny.CodePoint = (_788_v5).Select((Companion_Default___.SafeIndex(_786_v3, _dafny.IntOfUint32((_788_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
		_ = _rhs121
		var _rhs122 _dafny.Int = (_786_v3).Minus(_786_v3)
		_ = _rhs122
		var _lhs91 *GlobalState = globalState
		_ = _lhs91
		var _lhs92 _dafny.Array = _784_v1
		_ = _lhs92
		var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
		_ = _lhs93
		_786_v3 = _rhs119
		_lhs91.F8 = _rhs120
		_787_v4 = _rhs121
		(_lhs92).ArraySet1(_rhs122, (_lhs93).Int())
		var _793_v10 _dafny.Array
		_ = _793_v10
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_16
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) bool = (func(_794_v0 bool) func(_dafny.Int) bool {
				return func(_795_i0 _dafny.Int) bool {
					return _794_v0
				}
			})(_783_v0)
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
		_793_v10 = _nw106
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))
		_ = _index124
		(_793_v10).ArraySet1(_783_v0, (_index124).Int())
		var _796_v11 _dafny.MultiSet
		_ = _796_v11
		_796_v11 = _dafny.MultiSetOf((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int), (_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int))
		var _797_v12 _dafny.Sequence
		_ = _797_v12
		_797_v12 = _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_788_v5).Cardinality()), _786_v3), _796_v11)
		var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))
		_ = _index125
		(_793_v10).ArraySet1(!((_797_v12).Select((Companion_Default___.SafeIndex(_786_v3, _dafny.IntOfUint32((_797_v12).Cardinality()))).Uint32()).(_dafny.MultiSet)).Contains(_786_v3), (_index125).Int())
		var _hi0 _dafny.Int = (_786_v3).Plus(_786_v3)
		_ = _hi0
		for _798_i1 := ((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int)).Minus((func() _dafny.Set {
			var _coll40 = _dafny.NewBuilder()
			_ = _coll40
			for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-59), _dafny.IntOfInt64(873))); ; {
				_compr_40, _ok42 := _iter42()
				if !_ok42 {
					break
				}
				var _818_v13 _dafny.Int
				_818_v13 = interface{}(_compr_40).(_dafny.Int)
				if ((_dafny.IntOfInt64(-59)).Cmp(_818_v13) <= 0) && ((_818_v13).Cmp(_dafny.IntOfInt64(873)) < 0) {
					_coll40.Add((_818_v13).Times((_dafny.Zero).Minus((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int))))
				}
			}
			return _coll40.ToSet()
		}()).Cardinality()); _798_i1.Cmp(_hi0) < 0; _798_i1 = _798_i1.Plus(_dafny.One) {
			var _799_v14 D1
			_ = _799_v14
			_799_v14 = Companion_D1_.Create_DC4_(_783_v0, _783_v0, _798_i1)
			if (_799_v14).Dtor_cf7() {
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
				_ = _index126
				(_784_v1).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_798_i1)), (_index126).Int())
				var _800_v15 _dafny.Array
				_ = _800_v15
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
				_ = _nw107
				_800_v15 = _nw107
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw108
				_800_v15 = _nw108
				var _801_v16 _dafny.Sequence
				_ = _801_v16
				_801_v16 = _dafny.SeqOf(false)
				var _rhs123 _dafny.Sequence = _801_v16
				_ = _rhs123
				var _rhs124 bool = _783_v0
				_ = _rhs124
				var _rhs125 _dafny.Int = (_dafny.Zero).Minus(((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int)).Plus((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int)))
				_ = _rhs125
				var _rhs126 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_788_v5, _788_v5)
				_ = _rhs126
				_801_v16 = _rhs123
				_783_v0 = _rhs124
				_786_v3 = _rhs125
				_788_v5 = _rhs126
				var _802_v17 *C0
				_ = _802_v17
				var _nw109 *C0 = New_C0_()
				_ = _nw109
				_nw109.Ctor__()
				_802_v17 = _nw109
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))
				_ = _index127
				(_793_v10).ArraySet1((_this).Fm27((_798_i1).Cmp((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int)) <= 0, false, _791_v8, globalState), (_index127).Int())
			} else {
				var _803_v18 _dafny.Array
				_ = _803_v18
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_17
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) D1 = (func(_804_v14 D1) func(_dafny.Int) D1 {
						return func(_805_i2 _dafny.Int) D1 {
							return _804_v14
						}
					})(_799_v14)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw110 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw110).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw110).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_803_v18 = _nw110
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_803_v18), 0))
				_ = _index128
				(_803_v18).ArraySet1(_799_v14, (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_803_v18), 0))
				_ = _index129
				(_803_v18).ArraySet1(_799_v14, (_index129).Int())
				var _806_v19 *C0
				_ = _806_v19
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__()
				_806_v19 = _nw111
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))
				_ = _index130
				(_793_v10).ArraySet1(_783_v0, (_index130).Int())
				var _807_v20 _dafny.MultiSet
				_ = _807_v20
				_807_v20 = _dafny.MultiSetOf(_787_v4)
				var _808_v21 _dafny.Map
				_ = _808_v21
				_808_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v0, _807_v20)
				var _809_v22 _dafny.Set
				_ = _809_v22
				_809_v22 = _dafny.SetOf(_798_i1)
				var _810_v23 _dafny.Set
				_ = _810_v23
				_810_v23 = _dafny.SetOf((_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool), (_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool))
				var _811_v24 _dafny.Sequence
				_ = _811_v24
				_811_v24 = _dafny.SeqOf(true, false)
				var _812_v25 _dafny.Sequence
				_ = _812_v25
				_812_v25 = _dafny.SeqOf((_811_v24).Select((Companion_Default___.SafeIndex(_798_i1, _dafny.IntOfUint32((_811_v24).Cardinality()))).Uint32()).(bool))
				var _813_v26 _dafny.Array
				_ = _813_v26
				var _nwElement0_24 _dafny.Map = (_808_v21).Merge(_808_v21)
				_ = _nwElement0_24
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(19))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_24, 0)
				(_nw112).ArraySet1(_808_v21, 1)
				(_nw112).ArraySet1(_808_v21, 2)
				(_nw112).ArraySet1((_808_v21).Merge(_808_v21), 3)
				(_nw112).ArraySet1(_808_v21, 4)
				(_nw112).ArraySet1(_808_v21, 5)
				(_nw112).ArraySet1(_808_v21, 6)
				(_nw112).ArraySet1(_808_v21, 7)
				(_nw112).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v0, _807_v20), 8)
				(_nw112).ArraySet1(Companion_Default___.Fm29(_783_v0, true, Companion_Default___.Fm30(globalState), _809_v22, globalState), 9)
				(_nw112).ArraySet1(Companion_Default___.Fm29(_783_v0, _783_v0, _810_v23, _809_v22, globalState), 10)
				(_nw112).ArraySet1((_808_v21).Update(_783_v0, _807_v20), 11)
				(_nw112).ArraySet1((func() _dafny.Map {
					if (_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_812_v25).Select((Companion_Default___.SafeIndex(_798_i1, _dafny.IntOfUint32((_812_v25).Cardinality()))).Uint32()).(bool)), _807_v20)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v0, _807_v20)
				})(), 12)
				(_nw112).ArraySet1((Companion_Default___.Fm29((_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool), _783_v0, _810_v23, _809_v22, globalState)).Merge(_808_v21), 13)
				(_nw112).ArraySet1((_808_v21).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool), _807_v20)), 14)
				(_nw112).ArraySet1(_808_v21, 15)
				(_nw112).ArraySet1((Companion_D4_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_783_v0, _807_v20))).Dtor_cf15(), 16)
				(_nw112).ArraySet1(_808_v21, 17)
				(_nw112).ArraySet1(_808_v21, 18)
				_813_v26 = _nw112
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_813_v26), 0))
				_ = _index131
				(_813_v26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.MultiSetOf(_787_v4)), (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_813_v26), 0))
				_ = _index132
				(_813_v26).ArraySet1(_808_v21, (_index132).Int())
				_787_v4 = _787_v4
			}
			var _814_v27 *C0
			_ = _814_v27
			var _nw113 *C0 = New_C0_()
			_ = _nw113
			_nw113.Ctor__()
			_814_v27 = _nw113
			var _815_v28 *C0
			_ = _815_v28
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__()
			_815_v28 = _nw114
			var _816_v29 _dafny.Sequence
			_ = _816_v29
			_816_v29 = _dafny.SeqOf((_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool))
			var _817_v30 _dafny.Map
			_ = _817_v30
			_817_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-256), _816_v29)
			_817_v30 = (_817_v30).Update((_796_v11).Cardinality(), _816_v29)
		}
		var _hi1 _dafny.Int = (_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int)
		_ = _hi1
		for _819_i3 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(838))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg63 _dafny.Int) interface{} {
				return coer63(arg63)
			}
		}((func(_828_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_829_i4 _dafny.Int) _dafny.CodePoint {
				return _828_v4
			}
		})(_787_v4)))).Cardinality())); _819_i3.Cmp(_hi1) < 0; _819_i3 = _819_i3.Plus(_dafny.One) {
			_792_v9 = _792_v9
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))
			_ = _index133
			(_793_v10).ArraySet1((_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool), (_index133).Int())
			(globalState).F8 = ((_dafny.Zero).Minus(_819_i3)).Cmp((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int)) >= 0
			var _820_v31 _dafny.Set
			_ = _820_v31
			_820_v31 = _dafny.SetOf(_786_v3, _819_i3)
			var _821_v32 D4
			_ = _821_v32
			_821_v32 = Companion_D4_.Create_DC12_(!(_820_v31).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-450))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}((func(_822_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_823_i5 _dafny.Int) _dafny.CodePoint {
					return _822_v4
				}
			})(_787_v4)))).Cardinality())), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(303), _819_i3), _787_v4, _786_v3, _dafny.IntOfInt64(754))
			var _824_v33 _dafny.Map
			_ = _824_v33
			_824_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(338))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}((func(_825_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_826_i6 _dafny.Int) _dafny.CodePoint {
					return _825_v4
				}
			})(_787_v4))), false)
			var _827_v34 D1
			_ = _827_v34
			_827_v34 = Companion_D1_.Create_DC4_(_783_v0, _783_v0, (_dafny.Zero).Minus((_824_v33).Cardinality()))
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
			_ = _index134
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
			_ = _index135
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
			_ = _index136
			var _rhs127 D4 = _821_v32
			_ = _rhs127
			var _rhs128 _dafny.Int = (_819_i3).Times((_this).Fm28((_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool), _827_v34, _786_v3, _dafny.IntOfInt64(-108), globalState))
			_ = _rhs128
			var _rhs129 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(261), ((_this).Fm28(_783_v0, _827_v34, (_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int), _819_i3, globalState)).Minus(_819_i3))
			_ = _rhs129
			var _rhs130 _dafny.Int = _786_v3
			_ = _rhs130
			var _lhs94 _dafny.Array = _784_v1
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
			_ = _lhs95
			var _lhs96 _dafny.Array = _784_v1
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
			_ = _lhs97
			var _lhs98 _dafny.Array = _784_v1
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))
			_ = _lhs99
			_821_v32 = _rhs127
			(_lhs94).ArraySet1(_rhs128, (_lhs95).Int())
			(_lhs96).ArraySet1(_rhs129, (_lhs97).Int())
			(_lhs98).ArraySet1(_rhs130, (_lhs99).Int())
		}
		r0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_788_v5, _788_v5), (Companion_Default___.SafeIndex((_784_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_784_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_788_v5, _788_v5)).Cardinality()))).Uint32(), _787_v4)
		r1 = _783_v0
		var _830_v35 _dafny.Set
		_ = _830_v35
		_830_v35 = _dafny.SetOf(false, (_793_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_793_v10), 0))).Int()).(bool))
		r2 = _dafny.SeqOf(_830_v35)
		r3 = _dafny.Companion_Sequence_.Concatenate(_788_v5, _dafny.UnicodeSeqOfUtf8Bytes("letpy"))
		return r0, r1, r2, r3
	}
}
func (_this *C3) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _831_v0 *C0
		_ = _831_v0
		var _nw115 *C0 = New_C0_()
		_ = _nw115
		_nw115.Ctor__()
		_831_v0 = _nw115
		var _832_v1 bool
		_ = _832_v1
		_832_v1 = true
		var _833_v2 _dafny.Int
		_ = _833_v2
		_833_v2 = _dafny.IntOfInt64(248)
		var _834_v3 _dafny.Map
		_ = _834_v3
		_834_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_833_v2, _832_v1)
		var _835_v4 D4
		_ = _835_v4
		_835_v4 = Companion_D4_.Create_DC12_(_832_v1, _833_v2, Companion_Default___.Fm31(_832_v1, _834_v3, globalState), _833_v2, _833_v2)
		var _pat_let_tv16 = _832_v1
		_ = _pat_let_tv16
		var _pat_let_tv17 = _832_v1
		_ = _pat_let_tv17
		if func(_source16 D4) bool {
			if _source16.Is_DC11() {
				var _836___mcc_h0 _dafny.Int = _source16.Get_().(D4_DC11).Cf16
				_ = _836___mcc_h0
				var _837___mcc_h1 _dafny.Set = _source16.Get_().(D4_DC11).Cf17
				_ = _837___mcc_h1
				var _838___mcc_h2 bool = _source16.Get_().(D4_DC11).Cf18
				_ = _838___mcc_h2
				var _839_cf18 bool = _838___mcc_h2
				_ = _839_cf18
				var _840_cf17 _dafny.Set = _837___mcc_h1
				_ = _840_cf17
				var _841_cf16 _dafny.Int = _836___mcc_h0
				_ = _841_cf16
				return _pat_let_tv16
			} else if _source16.Is_DC12() {
				var _842___mcc_h3 bool = _source16.Get_().(D4_DC12).Cf19
				_ = _842___mcc_h3
				var _843___mcc_h4 _dafny.Int = _source16.Get_().(D4_DC12).Cf20
				_ = _843___mcc_h4
				var _844___mcc_h5 _dafny.CodePoint = _source16.Get_().(D4_DC12).Cf21
				_ = _844___mcc_h5
				var _845___mcc_h6 _dafny.Int = _source16.Get_().(D4_DC12).Cf22
				_ = _845___mcc_h6
				var _846___mcc_h7 _dafny.Int = _source16.Get_().(D4_DC12).Cf23
				_ = _846___mcc_h7
				var _847_cf23 _dafny.Int = _846___mcc_h7
				_ = _847_cf23
				var _848_cf22 _dafny.Int = _845___mcc_h6
				_ = _848_cf22
				var _849_cf21 _dafny.CodePoint = _844___mcc_h5
				_ = _849_cf21
				var _850_cf20 _dafny.Int = _843___mcc_h4
				_ = _850_cf20
				var _851_cf19 bool = _842___mcc_h3
				_ = _851_cf19
				return (_848_cf22).Cmp(_850_cf20) < 0
			} else {
				var _852___mcc_h8 _dafny.Map = _source16.Get_().(D4_DC10).Cf15
				_ = _852___mcc_h8
				var _853_cf15 _dafny.Map = _852___mcc_h8
				_ = _853_cf15
				return _pat_let_tv17
			}
		}(_835_v4) {
			var _854_v6 _dafny.Array
			_ = _854_v6
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_18
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = (func(_855_v1 bool, _856_v2 _dafny.Int) func(_dafny.Int) bool {
					return func(_857_i0 _dafny.Int) bool {
						return (Companion_D4_.Create_DC11_(_dafny.IntOfInt64(-957), func() _dafny.Set {
							var _coll41 = _dafny.NewBuilder()
							_ = _coll41
							for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-371), _dafny.IntOfInt64(330))); ; {
								_compr_41, _ok43 := _iter43()
								if !_ok43 {
									break
								}
								var _858_v5 _dafny.Int
								_858_v5 = interface{}(_compr_41).(_dafny.Int)
								if ((_dafny.IntOfInt64(-371)).Cmp(_858_v5) <= 0) && ((_858_v5).Cmp(_dafny.IntOfInt64(330)) < 0) {
									_coll41.Add((_858_v5).Plus(_856_v2))
								}
							}
							return _coll41.ToSet()
						}(), _855_v1)).Dtor_cf18()
					}
				})(_832_v1, _833_v2)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw116 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw116).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw116).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_854_v6 = _nw116
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
			_ = _index137
			(_854_v6).ArraySet1(_832_v1, (_index137).Int())
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
			_ = _index138
			(_854_v6).ArraySet1(_832_v1, (_index138).Int())
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
			_ = _index139
			(_854_v6).ArraySet1(_832_v1, (_index139).Int())
			var _859_v7 *C0
			_ = _859_v7
			var _nw117 *C0 = New_C0_()
			_ = _nw117
			_nw117.Ctor__()
			_859_v7 = _nw117
			var _860_v8 _dafny.Sequence
			_ = _860_v8
			_860_v8 = _dafny.SeqOf(_832_v1)
			if ((_860_v8).Select((Companion_Default___.SafeIndex(_833_v2, _dafny.IntOfUint32((_860_v8).Cardinality()))).Uint32()).(bool)) || ((func() bool {
				if (_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool) {
					return _832_v1
				}
				return (_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool)
			})()) {
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
				_ = _index140
				(_854_v6).ArraySet1(!(_832_v1), (_index140).Int())
				var _861_v9 _dafny.Sequence
				_ = _861_v9
				_861_v9 = _dafny.UnicodeSeqOfUtf8Bytes("ewyay")
				_861_v9 = _dafny.Companion_Sequence_.Concatenate(_861_v9, _dafny.Companion_Sequence_.Concatenate(_861_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-365))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_862_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_863_i1 _dafny.Int) _dafny.CodePoint {
						return _862_p0
					}
				})(p0)))))
				var _864_v10 _dafny.Map
				_ = _864_v10
				_864_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_v6, _832_v1)
				var _865_v11 D5
				_ = _865_v11
				_865_v11 = Companion_D5_.Create_DC13_(_854_v6)
				_864_v10 = (_864_v10).Update((_865_v11).Dtor_cf24(), (_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool))
				var _866_v12 _dafny.MultiSet
				_ = _866_v12
				_866_v12 = _dafny.MultiSetOf(false, _832_v1, true, (_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), ((_dafny.Zero).Minus(_833_v2)).Cmp(_833_v2) == 0)
				_866_v12 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_860_v8, _860_v8))).Intersection(_866_v12)
				var _867_v13 _dafny.Map
				_ = _867_v13
				_867_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_833_v2, _831_v0)
				var _868_v14 _dafny.Map
				_ = _868_v14
				_868_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_833_v2, (_867_v13).Cardinality())
				var _869_v15 D5
				_ = _869_v15
				_869_v15 = Companion_D5_.Create_DC14_(_833_v2, (_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool))
				var _870_v16 _dafny.Map
				_ = _870_v16
				_870_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), _dafny.SetOf(_859_v7, _831_v0)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_869_v15).Dtor_cf25(), _833_v2))
				var _871_v17 _dafny.Set
				_ = _871_v17
				_871_v17 = _dafny.SetOf(_831_v0)
				var _872_v18 _dafny.Map
				_ = _872_v18
				_872_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), _871_v17)
				var _873_v19 D1
				_ = _873_v19
				_873_v19 = Companion_D1_.Create_DC4_(_832_v1, (_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), _833_v2)
				var _874_v20 _dafny.MultiSet
				_ = _874_v20
				_874_v20 = _dafny.MultiSetOf(_833_v2, _833_v2)
				var _875_v21 _dafny.Array
				_ = _875_v21
				var _nwElement0_25 _dafny.Map = _868_v14
				_ = _nwElement0_25
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(13))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_25, 0)
				(_nw118).ArraySet1((func() _dafny.Map {
					if (_870_v16).Contains(_872_v18) {
						return (_870_v16).Get(_872_v18).(_dafny.Map)
					}
					return _868_v14
				})(), 1)
				(_nw118).ArraySet1(_868_v14, 2)
				(_nw118).ArraySet1(((_868_v14).Update(_833_v2, _833_v2)).Update(_833_v2, _833_v2), 3)
				(_nw118).ArraySet1(_868_v14, 4)
				(_nw118).ArraySet1(_868_v14, 5)
				(_nw118).ArraySet1((_868_v14).Update((_this).Fm28((_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), _873_v19, (_874_v20).Cardinality(), _833_v2, globalState), _833_v2), 6)
				(_nw118).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_833_v2, _833_v2), 7)
				(_nw118).ArraySet1((_868_v14).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(136), _dafny.IntOfInt64(457))), 8)
				(_nw118).ArraySet1(_868_v14, 9)
				(_nw118).ArraySet1(_868_v14, 10)
				(_nw118).ArraySet1(_868_v14, 11)
				(_nw118).ArraySet1(_868_v14, 12)
				_875_v21 = _nw118
				_875_v21 = _875_v21
			} else {
				var _876_v22 _dafny.Sequence
				_ = _876_v22
				_876_v22 = _dafny.UnicodeSeqOfUtf8Bytes("nhxyjlt")
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
				_ = _index141
				(_854_v6).ArraySet1((Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_876_v22).Cardinality()), _833_v2)).Cmp((_dafny.SetOf(_833_v2)).Cardinality()) < 0, (_index141).Int())
				var _877_v23 _dafny.Sequence
				_ = _877_v23
				_877_v23 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("hpdclvpfa"))
				_876_v22 = _dafny.Companion_Sequence_.Concatenate(_876_v22, (_877_v23).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_833_v2), _dafny.IntOfUint32((_877_v23).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _878_v24 _dafny.Set
				_ = _878_v24
				_878_v24 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_879_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_880_i2 _dafny.Int) _dafny.CodePoint {
						return _879_p0
					}
				})(p0))))
				var _881_v25 _dafny.Map
				_ = _881_v25
				_881_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_860_v8, false)
				var _882_v26 _dafny.Map
				_ = _882_v26
				_882_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_833_v2, p0)
				var _883_v27 _dafny.MultiSet
				_ = _883_v27
				_883_v27 = _dafny.MultiSetOf(_833_v2, _833_v2, _833_v2, (_882_v26).Cardinality())
				var _884_v28 _dafny.MultiSet
				_ = _884_v28
				_884_v28 = _dafny.MultiSetOf(false, false)
				var _885_v29 _dafny.Map
				_ = _885_v29
				_885_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), !(_832_v1))
				var _886_v30 _dafny.Sequence
				_ = _886_v30
				_886_v30 = _dafny.SeqOf(_833_v2)
				var _887_v31 _dafny.Array
				_ = _887_v31
				var _nwElement0_26 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_833_v2)), _833_v2)
				_ = _nwElement0_26
				var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(24))
				_ = _nw119
				(_nw119).ArraySet1(_nwElement0_26, 0)
				(_nw119).ArraySet1(Companion_Default___.SafeDivisionInt(_833_v2, _833_v2), 1)
				(_nw119).ArraySet1(_dafny.IntOfInt64(705), 2)
				(_nw119).ArraySet1(_833_v2, 3)
				(_nw119).ArraySet1(_dafny.IntOfInt64(559), 4)
				(_nw119).ArraySet1(_833_v2, 5)
				(_nw119).ArraySet1(_833_v2, 6)
				(_nw119).ArraySet1((_833_v2).Minus(_833_v2), 7)
				(_nw119).ArraySet1((_878_v24).Cardinality(), 8)
				(_nw119).ArraySet1(_833_v2, 9)
				(_nw119).ArraySet1(((_881_v25).Merge(_881_v25)).Cardinality(), 10)
				(_nw119).ArraySet1((func() _dafny.Int {
					if (_883_v27).Contains((_884_v28).Cardinality()) {
						return (_883_v27).Multiplicity((_884_v28).Cardinality())
					}
					return _833_v2
				})(), 11)
				(_nw119).ArraySet1(_833_v2, 12)
				(_nw119).ArraySet1(((_dafny.Zero).Minus(_833_v2)).Times(_833_v2), 13)
				(_nw119).ArraySet1(_dafny.IntOfInt64(779), 14)
				(_nw119).ArraySet1(_dafny.IntOfUint32((_860_v8).Cardinality()), 15)
				(_nw119).ArraySet1((_dafny.Zero).Minus(((_885_v29).Cardinality()).Plus((_886_v30).Select((Companion_Default___.SafeIndex(_833_v2, _dafny.IntOfUint32((_886_v30).Cardinality()))).Uint32()).(_dafny.Int))), 16)
				(_nw119).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qryuduwux")).Cardinality()), 17)
				(_nw119).ArraySet1(_833_v2, 18)
				(_nw119).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_886_v30).Cardinality())), 19)
				(_nw119).ArraySet1(_833_v2, 20)
				(_nw119).ArraySet1(_833_v2, 21)
				(_nw119).ArraySet1(_833_v2, 22)
				(_nw119).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("loerye")).Cardinality()), _833_v2), 23)
				_887_v31 = _nw119
				var _888_v32 _dafny.Sequence
				_ = _888_v32
				_888_v32 = _dafny.SeqOf(_887_v31, _887_v31, _887_v31)
				_887_v31 = (_888_v32).Select((Companion_Default___.SafeIndex((_833_v2).Minus(_833_v2), _dafny.IntOfUint32((_888_v32).Cardinality()))).Uint32()).(_dafny.Array)
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_887_v31), 0))
				_ = _index142
				(_887_v31).ArraySet1(_833_v2, (_index142).Int())
				var _889_v33 _dafny.Set
				_ = _889_v33
				_889_v33 = _dafny.SetOf((_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool), false)
				var _890_v34 _dafny.Map
				_ = _890_v34
				_890_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_833_v2, _889_v33)
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_887_v31), 0))
				_ = _index143
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
				_ = _index144
				var _rhs131 _dafny.Sequence = _876_v22
				_ = _rhs131
				var _rhs132 _dafny.Int = (_833_v2).Times((func() _dafny.Int {
					if _832_v1 {
						return _833_v2
					}
					return _dafny.IntOfInt64(865)
				})())
				_ = _rhs132
				var _rhs133 bool = !((_854_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))).Int()).(bool))
				_ = _rhs133
				var _rhs134 _dafny.Map = _890_v34
				_ = _rhs134
				var _rhs135 _dafny.Int = _833_v2
				_ = _rhs135
				var _lhs100 _dafny.Array = _887_v31
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_887_v31), 0))
				_ = _lhs101
				var _lhs102 _dafny.Array = _854_v6
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_854_v6), 0))
				_ = _lhs103
				_876_v22 = _rhs131
				(_lhs100).ArraySet1(_rhs132, (_lhs101).Int())
				(_lhs102).ArraySet1(_rhs133, (_lhs103).Int())
				_890_v34 = _rhs134
				_833_v2 = _rhs135
				var _891_v35 _dafny.Map
				_ = _891_v35
				_891_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _833_v2)
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_887_v31), 0))
				_ = _index145
				(_887_v31).ArraySet1((func() _dafny.Int {
					if (_891_v35).Contains(Companion_Default___.Fm31(!(true), _834_v3, globalState)) {
						return (_891_v35).Get(Companion_Default___.Fm31(!(true), _834_v3, globalState)).(_dafny.Int)
					}
					return _dafny.IntOfInt64(111)
				})(), (_index145).Int())
			}
			var _892_v36 *C0
			_ = _892_v36
			var _nw120 *C0 = New_C0_()
			_ = _nw120
			_nw120.Ctor__()
			_892_v36 = _nw120
		} else {
			(globalState).F8 = _832_v1
			(globalState).F8 = _832_v1
			var _893_v37 T0
			_ = _893_v37
			var _nw121 *C1 = New_C1_()
			_ = _nw121
			_nw121.Ctor__(_832_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(173))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}(func(_894_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})))
			_893_v37 = _nw121
			var _895_v38 _dafny.Array
			_ = _895_v38
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_19
			var _nw122 _dafny.Array
			_ = _nw122
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw122 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) bool = (func(_896_v1 bool) func(_dafny.Int) bool {
					return func(_897_i4 _dafny.Int) bool {
						return _896_v1
					}
				})(_832_v1)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw122 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw122).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw122).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_895_v38 = _nw122
			var _898_v39 _dafny.Map
			_ = _898_v39
			_898_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _832_v1)
			var _899_v40 T1
			_ = _899_v40
			var _nw123 *C2 = New_C2_()
			_ = _nw123
			_nw123.Ctor__(_832_v1, _895_v38, _898_v39, _832_v1)
			_899_v40 = _nw123
			var _900_v41 _dafny.Map
			_ = _900_v41
			_900_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_893_v37, _899_v40)
			_900_v41 = (_900_v41).Update(_893_v37, _899_v40)
			_833_v2 = (_833_v2).Plus(_dafny.IntOfInt64(752))
			var _901_v42 _dafny.Map
			_ = _901_v42
			_901_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v38, _833_v2)
			_901_v42 = (_901_v42).Update(_895_v38, (func() _dafny.Int {
				if _832_v1 {
					return _833_v2
				}
				return _833_v2
			})())
		}
		var _rhs136 bool = !(_832_v1)
		_ = _rhs136
		var _rhs137 _dafny.Int = (_dafny.Zero).Minus(_833_v2)
		_ = _rhs137
		r1 = _rhs136
		_833_v2 = _rhs137
		var _902_v43 _dafny.Set
		_ = _902_v43
		_902_v43 = _dafny.SetOf(_833_v2, _833_v2)
		var _903_v44 _dafny.Sequence
		_ = _903_v44
		_903_v44 = _dafny.SeqOf((_dafny.IntOfInt64(725)).Cmp((_902_v43).Cardinality()) >= 0, (_833_v2).Cmp(_833_v2) < 0)
		r1 = (_903_v44).Select((Companion_Default___.SafeIndex(_833_v2, _dafny.IntOfUint32((_903_v44).Cardinality()))).Uint32()).(bool)
		var _904_v45 _dafny.Array
		_ = _904_v45
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_20
		var _nw124 _dafny.Array
		_ = _nw124
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw124 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = (func(_905_v2 _dafny.Int, _906_v44 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_907_i6 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_907_i6, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_905_v2, _dafny.IntOfUint32((_906_v44).Cardinality()))).Cardinality()))
				}
			})(_833_v2, _903_v44)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw124 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw124).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw124).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_904_v45 = _nw124
		for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_904_v45), 0))); ; {
			_guard_loop_2, _ok44 := _iter44()
			if !_ok44 {
				break
			}
			var _908_i5 _dafny.Int
			_908_i5 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_908_i5).Sign() != -1) && ((_908_i5).Cmp(_dafny.ArrayLen((_904_v45), 0)) < 0)) {
				(_904_v45).ArraySet1(Companion_Default___.SafeModuloInt(_908_i5, _833_v2), (_908_i5).Int())
			}
		}
		(globalState).F8 = _832_v1
		var _909_v46 _dafny.Sequence
		_ = _909_v46
		_909_v46 = _dafny.UnicodeSeqOfUtf8Bytes("tgohvumtv")
		r0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if _832_v1 {
				return _904_v45
			}
			return _904_v45
		})(), _dafny.Companion_Sequence_.Concatenate(_909_v46, _dafny.UnicodeSeqOfUtf8Bytes("kaeltdutj")))
		r1 = (func() bool {
			if _832_v1 {
				return _832_v1
			}
			return _832_v1
		})()
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f9  _dafny.Map
	_f10 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f9 = _dafny.EmptyMap
	_this._f10 = false
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

func (_this *C4) F9() _dafny.Map {
	return _this._f9
}
func (_this *C4) F10() bool {
	return _this._f10
}
func (_this *C4) Ctor__(f9 _dafny.Map, f10 bool) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C4) Fm0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !((_this).F10()) || ((_this).F10())
	}
}
func (_this *C4) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-969), (_dafny.IntOfInt64(-383)).Times(_dafny.IntOfInt64(184)))
	}
}
func (_this *C4) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _910_v0 _dafny.Sequence
		_ = _910_v0
		_910_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qlknnmx")
		var _911_v1 D1
		_ = _911_v1
		_911_v1 = Companion_D1_.Create_DC3_(_910_v0)
		var _912_v2 _dafny.Sequence
		_ = _912_v2
		_912_v2 = _dafny.SeqOf((_this).F10())
		var _913_v3 _dafny.Int
		_ = _913_v3
		_913_v3 = _dafny.IntOfInt64(30)
		var _pat_let_tv18 = _912_v2
		_ = _pat_let_tv18
		var _pat_let_tv19 = globalState
		_ = _pat_let_tv19
		var _pat_let_tv20 = _913_v3
		_ = _pat_let_tv20
		var _pat_let_tv21 = globalState
		_ = _pat_let_tv21
		var _914_v4 _dafny.Array
		_ = _914_v4
		var _nwElement0_27 D1 = _911_v1
		_ = _nwElement0_27
		var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(19))
		_ = _nw125
		(_nw125).ArraySet1(_nwElement0_27, 0)
		(_nw125).ArraySet1(_911_v1, 1)
		(_nw125).ArraySet1(_911_v1, 2)
		(_nw125).ArraySet1(_911_v1, 3)
		(_nw125).ArraySet1(_911_v1, 4)
		(_nw125).ArraySet1(_911_v1, 5)
		(_nw125).ArraySet1(_911_v1, 6)
		(_nw125).ArraySet1(_911_v1, 7)
		(_nw125).ArraySet1(_911_v1, 8)
		(_nw125).ArraySet1(_911_v1, 9)
		(_nw125).ArraySet1(func(_pat_let18_0 D1) D1 {
			return func(_915_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let19_0 _dafny.Sequence) D1 {
					return func(_918_dt__update_hcf5_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC3_(_918_dt__update_hcf5_h0)
					}(_pat_let19_0)
				}(Companion_Default___.Fm22(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}(func(_916_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))).Cardinality()), (_this).F10(), (_this).Fm0(_pat_let_tv18, false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-736))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}(func(_917_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality()), (_this).F10(), _pat_let_tv19), _pat_let_tv20, _pat_let_tv21))
			}(_pat_let18_0)
		}(_911_v1), 10)
		(_nw125).ArraySet1(_911_v1, 11)
		(_nw125).ArraySet1(Companion_D1_.Create_DC3_(_910_v0), 12)
		(_nw125).ArraySet1(_911_v1, 13)
		(_nw125).ArraySet1(_911_v1, 14)
		(_nw125).ArraySet1(_911_v1, 15)
		(_nw125).ArraySet1(_911_v1, 16)
		(_nw125).ArraySet1(Companion_D1_.Create_DC3_(_910_v0), 17)
		(_nw125).ArraySet1(_911_v1, 18)
		_914_v4 = _nw125
		var _919_v5 _dafny.Map
		_ = _919_v5
		_919_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_913_v3, _914_v4)
		_914_v4 = (func() _dafny.Array {
			if (_this).F10() {
				return (func() _dafny.Array {
					if (_919_v5).Contains(_dafny.IntOfInt64(67)) {
						return (_919_v5).Get(_dafny.IntOfInt64(67)).(_dafny.Array)
					}
					return _914_v4
				})()
			}
			return _914_v4
		})()
		_910_v0 = _910_v0
		r1 = (_this).F10()
		(globalState).F8 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("bvcipvb"), _dafny.UnicodeSeqOfUtf8Bytes("ruf"))
		var _920_v6 *C0
		_ = _920_v6
		var _nw126 *C0 = New_C0_()
		_ = _nw126
		_nw126.Ctor__()
		_920_v6 = _nw126
		_920_v6 = _920_v6
		_913_v3 = _913_v3
		var _921_v7 _dafny.MultiSet
		_ = _921_v7
		_921_v7 = _dafny.MultiSetOf((_this).F10())
		var _922_v8 _dafny.Map
		_ = _922_v8
		_922_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_921_v7).Cardinality(), _dafny.IntOfUint32((_910_v0).Cardinality()))
		r0 = _dafny.Companion_Sequence_.Update(_910_v0, (Companion_Default___.SafeIndex(_913_v3, _dafny.IntOfUint32((_910_v0).Cardinality()))).Uint32(), (_910_v0).Select((Companion_Default___.SafeIndex((_922_v8).Cardinality(), _dafny.IntOfUint32((_910_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
		r1 = (_this).F10()
		r2 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer69 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_923_i2 _dafny.Int) _dafny.Set {
			return _dafny.SetOf((_this).F10(), (_this).F10(), !(true))
		}))
		r3 = _910_v0
		return r0, r1, r2, r3
	}
}
func (_this *C4) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _924_i0 _dafny.Int
		_ = _924_i0
		_924_i0 = _dafny.Zero
		{
			for !(((func() bool {
				if (_this).F10() {
					return (_this).F10()
				}
				return (_this).F10()
			})()) || ((func() bool {
				if (_this).F10() {
					return !((_this).F10())
				}
				return (_this).F10()
			})())) {
				{
					if (_924_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_924_i0 = (_924_i0).Plus(_dafny.One)
					var _925_v0 _dafny.Sequence
					_ = _925_v0
					_925_v0 = _dafny.SeqOf((_this).F10())
					var _926_v1 _dafny.Int
					_ = _926_v1
					_926_v1 = _dafny.IntOfInt64(169)
					var _927_v2 _dafny.Map
					_ = _927_v2
					_927_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_926_v1, !((_this).F10()))
					var _928_v3 _dafny.Set
					_ = _928_v3
					_928_v3 = _dafny.SetOf((_this).Fm0(_925_v0, (_this).F10(), (_927_v2).Cardinality(), (_this).F10(), globalState))
					var _929_v4 _dafny.Map
					_ = _929_v4
					_929_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if (_this).F10() {
							return false
						}
						return !((_this).F10())
					})(), _928_v3)
					var _930_v5 _dafny.Array
					_ = _930_v5
					var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
					_ = _nw127
					_930_v5 = _nw127
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_930_v5), 0))
					_ = _index146
					(_930_v5).ArraySet1(_926_v1, (_index146).Int())
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_930_v5), 0))
					_ = _index147
					var _rhs138 _dafny.Map = (_929_v4).Merge(_929_v4)
					_ = _rhs138
					var _rhs139 _dafny.Int = (_dafny.Zero).Minus(_926_v1)
					_ = _rhs139
					var _lhs104 _dafny.Array = _930_v5
					_ = _lhs104
					var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_930_v5), 0))
					_ = _lhs105
					_929_v4 = _rhs138
					(_lhs104).ArraySet1(_rhs139, (_lhs105).Int())
					(globalState).F8 = (_this).F10()
					var _931_v6 _dafny.MultiSet
					_ = _931_v6
					_931_v6 = _dafny.MultiSetOf(false)
					var _932_v7 _dafny.Map
					_ = _932_v7
					_932_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_926_v1)).Cardinality(), _931_v6), _926_v1)
					var _933_v8 _dafny.Map
					_ = _933_v8
					_933_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_930_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_930_v5), 0))).Int()).(_dafny.Int), _931_v6)
					_932_v7 = (_932_v7).Update((_933_v8).Merge(_933_v8), _926_v1)
					var _934_v10 _dafny.Set
					_ = _934_v10
					_934_v10 = _dafny.SetOf(p0)
					_926_v1 = ((_930_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_930_v5), 0))).Int()).(_dafny.Int)).Minus(((_this).Fm1((_this).F9(), globalState)).Plus(((func() _dafny.Map {
						var _coll42 = _dafny.NewMapBuilder()
						_ = _coll42
						for _iter45 := _dafny.Iterate((_934_v10).Elements()); ; {
							_compr_42, _ok45 := _iter45()
							if !_ok45 {
								break
							}
							var _935_v9 _dafny.CodePoint
							_935_v9 = interface{}(_compr_42).(_dafny.CodePoint)
							if (_934_v10).Contains(_935_v9) {
								_coll42.Add(_935_v9, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(186))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg70 _dafny.Int) interface{} {
										return coer70(arg70)
									}
								}((func(_936_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_937_i1 _dafny.Int) _dafny.CodePoint {
										return _936_p0
									}
								})(p0)))).Cardinality()))
							}
						}
						return _coll42.ToMap()
					}()).Update(p0, (_930_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_930_v5), 0))).Int()).(_dafny.Int))).Cardinality()))
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _938_v11 _dafny.Sequence
		_ = _938_v11
		_938_v11 = _dafny.SeqOf(_dafny.SetOf((_this).F10(), true))
		var _939_v12 _dafny.Int
		_ = _939_v12
		_939_v12 = _dafny.IntOfInt64(-285)
		var _940_v13 _dafny.Set
		_ = _940_v13
		_940_v13 = _dafny.SetOf((_this).F10(), (_this).F10())
		if !((_938_v11).Select((Companion_Default___.SafeIndex(_939_v12, _dafny.IntOfUint32((_938_v11).Cardinality()))).Uint32()).(_dafny.Set)).Equals(_940_v13) {
			(globalState).F8 = (_this).F10()
			r1 = (_this).F10()
			var _941_v14 D0
			_ = _941_v14
			_941_v14 = Companion_D0_.Create_DC0_((_this).F10())
			var _942_v15 _dafny.Map
			_ = _942_v15
			_942_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D0 {
				if (_this).F10() {
					return _941_v14
				}
				return _941_v14
			})(), !(!((_this).F10()) || ((_this).F10())))
			_942_v15 = (_942_v15).Update(Companion_Default___.Fm23(globalState), (_this).F10())
			var _943_v16 _dafny.Sequence
			_ = _943_v16
			_943_v16 = _dafny.UnicodeSeqOfUtf8Bytes("hlsgfllkl")
			if ((_dafny.IntOfUint32((_943_v16).Cardinality())).Cmp(_939_v12) != 0) == ((_this).F10()) {
				var _944_v17 _dafny.Map
				_ = _944_v17
				_944_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _939_v12)
				_944_v17 = (_944_v17).Update(p0, _939_v12)
				_939_v12 = _939_v12
				var _945_v18 _dafny.Array
				_ = _945_v18
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw128
				_945_v18 = _nw128
				_945_v18 = _945_v18
				var _946_v19 _dafny.Array
				_ = _946_v19
				var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
				_ = _nw129
				_946_v19 = _nw129
				var _947_v20 _dafny.Array
				_ = _947_v20
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_21
				var _nw130 _dafny.Array
				_ = _nw130
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw130 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = func(_948_i2 _dafny.Int) _dafny.Int {
						return (_948_i2).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F10())).Cardinality()))
					}
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw130 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw130).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw130).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_947_v20 = _nw130
				var _949_v21 D3
				_ = _949_v21
				_949_v21 = Companion_D3_.Create_DC7_(_947_v20)
				var _950_v22 _dafny.Map
				_ = _950_v22
				_950_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_949_v21).Dtor_cf11(), (_this).Fm1((_this).F9(), globalState))
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_946_v19), 0))
				_ = _index148
				(_946_v19).ArraySet1(_950_v22, (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_946_v19), 0))
				_ = _index149
				(_946_v19).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_947_v20, (_this).Fm1((_this).F9(), globalState))).Merge(_950_v22), (_index149).Int())
				var _951_v23 _dafny.Sequence
				_ = _951_v23
				_951_v23 = _dafny.SeqOf(_dafny.IntOfInt64(762))
				var _952_v24 _dafny.Sequence
				_ = _952_v24
				_952_v24 = _dafny.SeqOf((_this).F10())
				_951_v23 = (func() _dafny.Sequence {
					if (_this).Fm0(_952_v24, (_this).F10(), _dafny.IntOfUint32((_dafny.SeqOf(_939_v12, _939_v12)).Cardinality()), !((_this).F10()), globalState) {
						return _951_v23
					}
					return _951_v23
				})()
			} else {
				var _953_v25 _dafny.Array
				_ = _953_v25
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw131
				_953_v25 = _nw131
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))
				_ = _index150
				(_953_v25).ArraySet1(!((_this).F10()) || ((_this).F10()), (_index150).Int())
				var _954_v26 _dafny.Sequence
				_ = _954_v26
				_954_v26 = _dafny.SeqOf((_this).F10())
				var _955_v27 _dafny.Map
				_ = _955_v27
				_955_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_954_v26, _954_v26), (_dafny.SetOf(_939_v12, _939_v12)).Cardinality())
				var _956_v28 _dafny.Set
				_ = _956_v28
				_956_v28 = _dafny.SetOf((_this).F9())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))
				_ = _index151
				var _rhs140 _dafny.Array = _953_v25
				_ = _rhs140
				var _rhs141 bool = (func() bool {
					if (_this).F10() {
						return (_this).Fm0(_954_v26, false, _939_v12, (_this).F10(), globalState)
					}
					return (_this).F10()
				})()
				_ = _rhs141
				var _rhs142 _dafny.Int = Companion_Default___.SafeModuloInt(_939_v12, _939_v12)
				_ = _rhs142
				var _rhs143 bool = (Companion_Default___.Fm24((_this).F10(), globalState)).IsDisjointFrom(_956_v28)
				_ = _rhs143
				var _rhs144 _dafny.Map = _955_v27
				_ = _rhs144
				var _lhs106 _dafny.Array = _953_v25
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))
				_ = _lhs107
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				_953_v25 = _rhs140
				(_lhs106).ArraySet1(_rhs141, (_lhs107).Int())
				_939_v12 = _rhs142
				_lhs108.F8 = _rhs143
				_955_v27 = _rhs144
				var _957_v29 *C0
				_ = _957_v29
				var _nw132 *C0 = New_C0_()
				_ = _nw132
				_nw132.Ctor__()
				_957_v29 = _nw132
				var _958_v30 _dafny.CodePoint
				_ = _958_v30
				_958_v30 = _dafny.CodePoint('h')
				var _959_v32 _dafny.Map
				_ = _959_v32
				_959_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_953_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))).Int()).(bool), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm25((_dafny.Zero).Minus(_939_v12), _dafny.IntOfUint32((_943_v16).Cardinality()), globalState)).Cardinality(), _939_v12)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v12, _939_v12)).Update(_dafny.IntOfUint32((_943_v16).Cardinality()), (func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter46 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v12, (_dafny.Zero).Minus(_939_v12))).Keys().Elements()); ; {
						_compr_43, _ok46 := _iter46()
						if !_ok46 {
							break
						}
						var _960_v31 _dafny.Int
						_960_v31 = interface{}(_compr_43).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v12, (_dafny.Zero).Minus(_939_v12))).Contains(_960_v31) {
							_coll43.Add(Companion_Default___.SafeModuloInt(_960_v31, _939_v12), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v12, _939_v12)).Cardinality())
						}
					}
					return _coll43.ToMap()
				}()).Cardinality())))
				var _961_v33 _dafny.Array
				_ = _961_v33
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw133
				_961_v33 = _nw133
				var _962_v34 _dafny.Map
				_ = _962_v34
				_962_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(323), _939_v12)
				var _rhs145 _dafny.CodePoint = _958_v30
				_ = _rhs145
				var _rhs146 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_953_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))).Int()).(bool), (_962_v34).Merge(_962_v34))
				_ = _rhs146
				var _rhs147 _dafny.Array = _961_v33
				_ = _rhs147
				_958_v30 = _rhs145
				_959_v32 = _rhs146
				_961_v33 = _rhs147
				var _963_v35 _dafny.Map
				_ = _963_v35
				_963_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_958_v30, (_this).F10())
				var _964_v36 _dafny.Map
				_ = _964_v36
				_964_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(_954_v26, false, _939_v12, (_this).Fm0(_954_v26, (_this).F10(), _939_v12, (_this).F10(), globalState), globalState), (Companion_D1_.Create_DC4_((_this).F10(), (func() bool {
					if (_963_v35).Contains(_958_v30) {
						return (_963_v35).Get(_958_v30).(bool)
					}
					return (_953_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))).Int()).(bool)
				})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pojq")).Cardinality())))).Dtor_cf7())
				_964_v36 = (_964_v36).Update((_this).F10(), true)
				var _965_v37 _dafny.Sequence
				_ = _965_v37
				_965_v37 = _dafny.SeqOf(_943_v16)
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(728), _dafny.ArrayLen((_953_v25), 0))
				_ = _index152
				(_953_v25).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_965_v37, (Companion_Default___.SafeIndex(_939_v12, _dafny.IntOfUint32((_965_v37).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-464))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_966_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_967_i3 _dafny.Int) _dafny.CodePoint {
						return _966_v30
					}
				})(_958_v30)))), _965_v37), (_index152).Int())
			}
			var _968_v38 _dafny.MultiSet
			_ = _968_v38
			_968_v38 = _dafny.MultiSetOf((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_939_v12, _dafny.IntOfInt64(391))))
			_968_v38 = (_968_v38).Union(_968_v38)
		} else {
			var _969_v39 *C0
			_ = _969_v39
			var _nw134 *C0 = New_C0_()
			_ = _nw134
			_nw134.Ctor__()
			_969_v39 = _nw134
			var _970_v40 _dafny.Sequence
			_ = _970_v40
			_970_v40 = _dafny.SeqOf(_969_v39, _969_v39)
			var _971_v41 _dafny.Sequence
			_ = _971_v41
			_971_v41 = _dafny.SeqOf((_this).F10(), (_this).F10())
			var _rhs148 _dafny.Int = (func() _dafny.Int {
				if (_this).F10() {
					return _dafny.IntOfUint32((_971_v41).Cardinality())
				}
				return _939_v12
			})()
			_ = _rhs148
			var _rhs149 bool = (_this).F10()
			_ = _rhs149
			var _rhs150 bool = (_this).F10()
			_ = _rhs150
			var _rhs151 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_970_v40, _970_v40), _970_v40)
			_ = _rhs151
			var _lhs109 *GlobalState = globalState
			_ = _lhs109
			var _lhs110 *GlobalState = globalState
			_ = _lhs110
			_939_v12 = _rhs148
			_lhs109.F8 = _rhs149
			_lhs110.F8 = _rhs150
			_970_v40 = _rhs151
			var _972_v42 _dafny.Array
			_ = _972_v42
			var _nw135 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw135
			_972_v42 = _nw135
			var _973_v43 _dafny.Sequence
			_ = _973_v43
			_973_v43 = _dafny.UnicodeSeqOfUtf8Bytes("coavp")
			var _974_v44 _dafny.MultiSet
			_ = _974_v44
			_974_v44 = _dafny.MultiSetOf(_dafny.IntOfUint32((_973_v43).Cardinality()))
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))
			_ = _index153
			(_972_v42).ArraySet1((_974_v44).Equals(_974_v44), (_index153).Int())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))
			_ = _index154
			var _rhs152 bool = (_this).F10()
			_ = _rhs152
			var _rhs153 bool = (_this).F10()
			_ = _rhs153
			var _rhs154 bool = (_this).F10()
			_ = _rhs154
			var _lhs111 *GlobalState = globalState
			_ = _lhs111
			var _lhs112 _dafny.Array = _972_v42
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))
			_ = _lhs113
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			_lhs111.F8 = _rhs152
			(_lhs112).ArraySet1(_rhs153, (_lhs113).Int())
			_lhs114.F8 = _rhs154
			_939_v12 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_972_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))).Int()).(bool) {
					return _939_v12
				}
				return _939_v12
			})(), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(323), _939_v12)))
			var _975_v45 _dafny.Array
			_ = _975_v45
			var _nwElement0_28 _dafny.CodePoint = p0
			_ = _nwElement0_28
			var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(12))
			_ = _nw136
			(_nw136).ArraySet1CodePoint(_nwElement0_28, 0)
			(_nw136).ArraySet1CodePoint(p0, 1)
			(_nw136).ArraySet1CodePoint(p0, 2)
			(_nw136).ArraySet1CodePoint(p0, 3)
			(_nw136).ArraySet1CodePoint(p0, 4)
			(_nw136).ArraySet1CodePoint(p0, 5)
			(_nw136).ArraySet1CodePoint(p0, 6)
			(_nw136).ArraySet1CodePoint(p0, 7)
			(_nw136).ArraySet1CodePoint(_dafny.CodePoint('c'), 8)
			(_nw136).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_972_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))).Int()).(bool) {
					return Companion_Default___.Fm26(p0, globalState)
				}
				return p0
			})(), 9)
			(_nw136).ArraySet1CodePoint(p0, 10)
			(_nw136).ArraySet1CodePoint(p0, 11)
			_975_v45 = _nw136
			_975_v45 = (func() _dafny.Array {
				if !((_this).F10()) {
					return _975_v45
				}
				return _975_v45
			})()
			var _976_v46 _dafny.Map
			_ = _976_v46
			_976_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.UnicodeSeqOfUtf8Bytes("kxm"))
			var _977_v47 _dafny.Sequence
			_ = _977_v47
			_977_v47 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_976_v46).Contains((_972_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))).Int()).(bool)) {
					return (_976_v46).Get((_972_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_972_v42), 0))).Int()).(bool)).(_dafny.Sequence)
				}
				return _973_v43
			})(), _973_v43)
			(globalState).F8 = !_dafny.Companion_Sequence_.Contains((_977_v47).Select((Companion_Default___.SafeIndex(_939_v12, _dafny.IntOfUint32((_977_v47).Cardinality()))).Uint32()).(_dafny.Sequence), p0)
		}
		var _978_v48 _dafny.Sequence
		_ = _978_v48
		_978_v48 = _dafny.SeqOf(_939_v12, (_939_v12).Minus(_939_v12), _939_v12)
		_939_v12 = (_978_v48).Select((Companion_Default___.SafeIndex(_939_v12, _dafny.IntOfUint32((_978_v48).Cardinality()))).Uint32()).(_dafny.Int)
		r1 = (_this).F10()
		var _979_v49 *C0
		_ = _979_v49
		var _nw137 *C0 = New_C0_()
		_ = _nw137
		_nw137.Ctor__()
		_979_v49 = _nw137
		var _980_v50 _dafny.Sequence
		_ = _980_v50
		_980_v50 = _dafny.UnicodeSeqOfUtf8Bytes("f")
		var _rhs155 bool = !(_dafny.Companion_Sequence_.Equal(_980_v50, _980_v50))
		_ = _rhs155
		var _rhs156 bool = (_this).F10()
		_ = _rhs156
		r1 = _rhs155
		r1 = _rhs156
		var _981_v51 _dafny.Array
		_ = _981_v51
		var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw138
		_981_v51 = _nw138
		var _982_v52 _dafny.Map
		_ = _982_v52
		_982_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_981_v51, _dafny.UnicodeSeqOfUtf8Bytes("xxghgxtt"))
		r0 = (_982_v52).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_981_v51, _980_v50))
		r1 = (_this).F10()
		return r0, r1
	}
}
func (_this *C4) M5(globalState *GlobalState) {
	{
		var _983_v0 _dafny.Sequence
		_ = _983_v0
		_983_v0 = _dafny.SeqOf((_this).F10(), (_this).F10())
		(globalState).F8 = _dafny.Companion_Sequence_.Contains(_983_v0, !((_this).F10()))
		(globalState).F8 = (_this).F10()
		var _984_i0 _dafny.Int
		_ = _984_i0
		_984_i0 = _dafny.Zero
		{
			for (_this).F10() {
				{
					if (_984_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_984_i0 = (_984_i0).Plus(_dafny.One)
					if (_this).F10() {
						var _985_v1 _dafny.Int
						_ = _985_v1
						_985_v1 = _dafny.IntOfInt64(762)
						_985_v1 = ((_dafny.Zero).Minus(_985_v1)).Minus(_985_v1)
						var _986_v2 _dafny.Sequence
						_ = _986_v2
						_986_v2 = _dafny.UnicodeSeqOfUtf8Bytes("qqk")
						var _987_v3 _dafny.Array
						_ = _987_v3
						var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
						_ = _nw139
						_987_v3 = _nw139
						var _988_v4 _dafny.Map
						_ = _988_v4
						_988_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("nsdcjxlyo"), _987_v3)
						var _989_v5 _dafny.Sequence
						_ = _989_v5
						var _990_v6 _dafny.Int
						_ = _990_v6
						var _991_v7 _dafny.Sequence
						_ = _991_v7
						var _992_v8 _dafny.Int
						_ = _992_v8
						var _out48 _dafny.Sequence
						_ = _out48
						var _out49 _dafny.Int
						_ = _out49
						var _out50 _dafny.Sequence
						_ = _out50
						var _out51 _dafny.Int
						_ = _out51
						_out48, _out49, _out50, _out51 = (_this).M6(_dafny.IntOfUint32((_986_v2).Cardinality()), (func() _dafny.Array {
							if (_988_v4).Contains(_986_v2) {
								return (_988_v4).Get(_986_v2).(_dafny.Array)
							}
							return _987_v3
						})(), (_985_v1).Times(_985_v1), globalState)
						_989_v5 = _out48
						_990_v6 = _out49
						_991_v7 = _out50
						_992_v8 = _out51
						var _993_v9 _dafny.CodePoint
						_ = _993_v9
						_993_v9 = _dafny.CodePoint('f')
						_993_v9 = _993_v9
						_986_v2 = _dafny.SeqOf(_993_v9)
						var _994_v10 T0
						_ = _994_v10
						var _nw140 *C3 = New_C3_()
						_ = _nw140
						_nw140.Ctor__()
						_994_v10 = _nw140
						_994_v10 = _994_v10
					} else {
						var _995_v11 _dafny.Int
						_ = _995_v11
						_995_v11 = _dafny.IntOfInt64(153)
						_995_v11 = _995_v11
						var _996_v12 _dafny.Array
						_ = _996_v12
						var _len0_22 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_22
						var _nw141 _dafny.Array
						_ = _nw141
						if _len0_22.Cmp(_dafny.Zero) == 0 {
							_nw141 = _dafny.NewArray(_len0_22)
						} else {
							var _init22 func(_dafny.Int) bool = func(_997_i1 _dafny.Int) bool {
								return !((_this).F10())
							}
							_ = _init22
							var _element0_22 = _init22(_dafny.Zero)
							_ = _element0_22
							_nw141 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
							(_nw141).ArraySet1(_element0_22, 0)
							var _nativeLen0_22 = (_len0_22).Int()
							_ = _nativeLen0_22
							for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
								(_nw141).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
							}
						}
						_996_v12 = _nw141
						var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_996_v12), 0))
						_ = _index155
						(_996_v12).ArraySet1((_this).F10(), (_index155).Int())
						var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_996_v12), 0))
						_ = _index156
						(_996_v12).ArraySet1(!((_983_v0).Select((Companion_Default___.SafeIndex(_995_v11, _dafny.IntOfUint32((_983_v0).Cardinality()))).Uint32()).(bool)), (_index156).Int())
						var _998_v13 _dafny.Sequence
						_ = _998_v13
						_998_v13 = _dafny.UnicodeSeqOfUtf8Bytes("xum")
						_998_v13 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-636))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg72 _dafny.Int) interface{} {
								return coer72(arg72)
							}
						}((func(_999_v12 _dafny.Array, _1000_v11 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
							return func(_1001_i2 _dafny.Int) _dafny.CodePoint {
								return (Companion_D4_.Create_DC12_((_999_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_999_v12), 0))).Int()).(bool), _dafny.IntOfInt64(357), _dafny.CodePoint('d'), (_dafny.Zero).Minus(_1000_v11), _1000_v11)).Dtor_cf21()
							}
						})(_996_v12, _995_v11)))
						var _1002_v14 *C1
						_ = _1002_v14
						var _nw142 *C1 = New_C1_()
						_ = _nw142
						_nw142.Ctor__(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}(func(_1003_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('a')
						})), _dafny.UnicodeSeqOfUtf8Bytes("sixbvx")), _998_v13)
						_1002_v14 = _nw142
						var _1004_v15 _dafny.Sequence
						_ = _1004_v15
						_1004_v15 = _dafny.SeqOf(_995_v11, _995_v11, _995_v11, _dafny.IntOfInt64(369), (_dafny.SetOf(_995_v11)).Cardinality())
						var _1005_v16 _dafny.Sequence
						_ = _1005_v16
						_1005_v16 = _dafny.SeqOf(_1004_v15)
						var _1006_v17 _dafny.Sequence
						_ = _1006_v17
						_1006_v17 = _dafny.SeqOf((_1005_v16).Select((Companion_Default___.SafeIndex(_995_v11, _dafny.IntOfUint32((_1005_v16).Cardinality()))).Uint32()).(_dafny.Sequence))
						var _1007_v18 _dafny.Sequence
						_ = _1007_v18
						_1007_v18 = _dafny.SeqOf((_dafny.IntOfUint32(((_1005_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1002_v14).F13()).Cardinality()), _dafny.IntOfUint32((_1005_v16).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Minus(_dafny.IntOfInt64(650)))
						_1004_v15 = _1007_v18
					}
					var _1008_v19 _dafny.Int
					_ = _1008_v19
					_1008_v19 = _dafny.IntOfInt64(745)
					var _1009_v20 _dafny.Set
					_ = _1009_v20
					_1009_v20 = _dafny.SetOf(true)
					_1008_v19 = (func() _dafny.Int {
						if (_1009_v20).IsProperSubsetOf(_1009_v20) {
							return Companion_Default___.SafeDivisionInt(_1008_v19, _dafny.IntOfInt64(426))
						}
						return _1008_v19
					})()
					(globalState).F8 = (_this).F10()
					var _1010_v21 _dafny.Array
					_ = _1010_v21
					var _nw143 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
					_ = _nw143
					_1010_v21 = _nw143
					var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1010_v21), 0))
					_ = _index157
					(_1010_v21).ArraySet1((_this).F10(), (_index157).Int())
					var _1011_v22 _dafny.Sequence
					_ = _1011_v22
					_1011_v22 = _dafny.SeqOf(_1008_v19)
					var _1012_v23 _dafny.Map
					_ = _1012_v23
					_1012_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_dafny.Zero).Minus((_this).Fm1((_this).F9(), globalState)))
					var _1013_v24 _dafny.Set
					_ = _1013_v24
					_1013_v24 = _dafny.SetOf(_1012_v23)
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1010_v21), 0))
					_ = _index158
					(_1010_v21).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1011_v22, _dafny.SeqOf((func() _dafny.Int {
						if (_1012_v23).Contains((_this).F10()) {
							return (_1012_v23).Get((_this).F10()).(_dafny.Int)
						}
						return _dafny.IntOfInt64(-128)
					})(), _dafny.IntOfInt64(411), (_1013_v24).Cardinality(), _1008_v19)), (_index158).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _1014_i4 _dafny.Int
		_ = _1014_i4
		_1014_i4 = _dafny.Zero
		{
			for (_this).F10() {
				{
					if (_1014_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1014_i4 = (_1014_i4).Plus(_dafny.One)
					if ((_this).F10()) == (!((_this).F10())) {
						var _1015_v25 _dafny.Sequence
						_ = _1015_v25
						_1015_v25 = _dafny.UnicodeSeqOfUtf8Bytes("yuctjfpul")
						var _1016_v26 *C1
						_ = _1016_v26
						var _nw144 *C1 = New_C1_()
						_ = _nw144
						_nw144.Ctor__((_this).F10(), _1015_v25)
						_1016_v26 = _nw144
						var _1017_v27 _dafny.Int
						_ = _1017_v27
						_1017_v27 = _dafny.IntOfInt64(729)
						var _1018_v28 _dafny.Map
						_ = _1018_v28
						_1018_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1017_v27, (_this).F9())
						var _1019_v29 _dafny.Map
						_ = _1019_v29
						_1019_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v28).Merge(_1018_v28), (_1017_v27).Plus(_dafny.IntOfUint32(((_1016_v26).F13()).Cardinality())))
						var _1020_v30 _dafny.Map
						_ = _1020_v30
						_1020_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1016_v26).F12(), _1017_v27)
						_1019_v29 = (_1019_v29).Update(_1018_v28, Companion_Default___.SafeModuloInt(_1017_v27, (func() _dafny.Int {
							if (_1020_v30).Contains((_this).F10()) {
								return (_1020_v30).Get((_this).F10()).(_dafny.Int)
							}
							return _1017_v27
						})()))
						var _1021_v31 _dafny.Sequence
						_ = _1021_v31
						_1021_v31 = _dafny.SeqOf((_1016_v26).F13(), (_1016_v26).F13(), (_1016_v26).F13())
						_1015_v25 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("w"), _dafny.Companion_Sequence_.Concatenate((_1021_v31).Select((Companion_Default___.SafeIndex(_1017_v27, _dafny.IntOfUint32((_1021_v31).Cardinality()))).Uint32()).(_dafny.Sequence), _1015_v25))
						_1017_v27 = (((_this).F9()).Merge((_this).F9())).Cardinality()
						var _1022_v32 _dafny.MultiSet
						_ = _1022_v32
						_1022_v32 = _dafny.MultiSetOf(_1017_v27)
						_1020_v30 = (_1020_v30).Update((_1022_v32).IsProperSubsetOf(_1022_v32), _1017_v27)
					} else {
						var _1023_v33 _dafny.Sequence
						_ = _1023_v33
						_1023_v33 = _dafny.UnicodeSeqOfUtf8Bytes("cckevci")
						var _1024_v34 _dafny.Int
						_ = _1024_v34
						_1024_v34 = _dafny.IntOfInt64(419)
						var _1025_v35 _dafny.Sequence
						_ = _1025_v35
						_1025_v35 = _dafny.SeqOf(_1024_v34)
						var _1026_v36 _dafny.Array
						_ = _1026_v36
						var _len0_23 _dafny.Int = _dafny.IntOfInt64(29)
						_ = _len0_23
						var _nw145 _dafny.Array
						_ = _nw145
						if _len0_23.Cmp(_dafny.Zero) == 0 {
							_nw145 = _dafny.NewArray(_len0_23)
						} else {
							var _init23 func(_dafny.Int) _dafny.CodePoint = func(_1027_i5 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('n')
							}
							_ = _init23
							var _element0_23 = _init23(_dafny.Zero)
							_ = _element0_23
							_nw145 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
							(_nw145).ArraySet1CodePoint(_element0_23, 0)
							var _nativeLen0_23 = (_len0_23).Int()
							_ = _nativeLen0_23
							for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
								(_nw145).ArraySet1CodePoint(_init23(_dafny.IntOf(_i0_23)), _i0_23)
							}
						}
						_1026_v36 = _nw145
						var _1028_v37 _dafny.MultiSet
						_ = _1028_v37
						_1028_v37 = _dafny.MultiSetOf(_1024_v34)
						var _rhs157 _dafny.Sequence = _1023_v33
						_ = _rhs157
						var _rhs158 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1025_v35, _dafny.SeqOf(_1024_v34, _1024_v34)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg74 _dafny.Int) interface{} {
								return coer74(arg74)
							}
						}((func(_1029_v34 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1030_i6 _dafny.Int) _dafny.Int {
								return _1029_v34
							}
						})(_1024_v34))))
						_ = _rhs158
						var _rhs159 _dafny.Array = _1026_v36
						_ = _rhs159
						var _rhs160 _dafny.Int = Companion_Default___.SafeDivisionInt(_1024_v34, _1024_v34)
						_ = _rhs160
						var _rhs161 _dafny.Int = Companion_Default___.SafeDivisionInt(_1024_v34, ((_1028_v37).Intersection(_1028_v37)).Cardinality())
						_ = _rhs161
						_1023_v33 = _rhs157
						_1025_v35 = _rhs158
						_1026_v36 = _rhs159
						_1024_v34 = _rhs160
						_1024_v34 = _rhs161
						(globalState).F8 = (_this).F10()
						var _1031_v38 _dafny.MultiSet
						_ = _1031_v38
						_1031_v38 = _dafny.MultiSetOf(!((_this).F10()))
						var _1032_v39 _dafny.Sequence
						_ = _1032_v39
						_1032_v39 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_983_v0), _1031_v38)
						var _1033_v40 _dafny.Map
						_ = _1033_v40
						_1033_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.CodePoint('g'))
						var _1034_v41 _dafny.MultiSet
						_ = _1034_v41
						_1034_v41 = _dafny.MultiSetOf(Companion_Default___.Fm22(_dafny.IntOfUint32((_dafny.SeqOf(_1024_v34)).Cardinality()), (_this).F10(), (_this).Fm0(_983_v0, (_this).F10(), (_this).Fm1((_this).F9(), globalState), (_this).F10(), globalState), _1024_v34, globalState))
						var _1035_v42 _dafny.Array
						_ = _1035_v42
						var _nwElement0_29 bool = (_this).F10()
						_ = _nwElement0_29
						var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(5))
						_ = _nw146
						(_nw146).ArraySet1(_nwElement0_29, 0)
						(_nw146).ArraySet1((_this).F10(), 1)
						(_nw146).ArraySet1((_this).F10(), 2)
						(_nw146).ArraySet1(!(false), 3)
						(_nw146).ArraySet1((_this).F10(), 4)
						_1035_v42 = _nw146
						var _1036_v43 _dafny.Sequence
						_ = _1036_v43
						_1036_v43 = _dafny.SeqOf(_1035_v42)
						var _1037_v44 _dafny.Array
						_ = _1037_v44
						var _nwElement0_30 _dafny.Int = _1024_v34
						_ = _nwElement0_30
						var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(25))
						_ = _nw147
						(_nw147).ArraySet1(_nwElement0_30, 0)
						(_nw147).ArraySet1((_1024_v34).Times(_dafny.IntOfUint32((_1023_v33).Cardinality())), 1)
						(_nw147).ArraySet1(_1024_v34, 2)
						(_nw147).ArraySet1(_1024_v34, 3)
						(_nw147).ArraySet1(((_1032_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.IntOfUint32((_1032_v39).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), 4)
						(_nw147).ArraySet1(_1024_v34, 5)
						(_nw147).ArraySet1(_1024_v34, 6)
						(_nw147).ArraySet1(_1024_v34, 7)
						(_nw147).ArraySet1((_1024_v34).Minus(_dafny.IntOfInt64(915)), 8)
						(_nw147).ArraySet1(_1024_v34, 9)
						(_nw147).ArraySet1(_1024_v34, 10)
						(_nw147).ArraySet1((func() _dafny.Int {
							if (_this).F10() {
								return (_1033_v40).Cardinality()
							}
							return _1024_v34
						})(), 11)
						(_nw147).ArraySet1(Companion_Default___.SafeModuloInt(_1024_v34, _1024_v34), 12)
						(_nw147).ArraySet1((_dafny.Zero).Minus(_1024_v34), 13)
						(_nw147).ArraySet1((func() _dafny.Int {
							if (_this).F10() {
								return (_dafny.Zero).Minus((_dafny.Zero).Minus(_1024_v34))
							}
							return _1024_v34
						})(), 14)
						(_nw147).ArraySet1((_1034_v41).Cardinality(), 15)
						(_nw147).ArraySet1((_1024_v34).Times(_dafny.IntOfInt64(-334)), 16)
						(_nw147).ArraySet1(_1024_v34, 17)
						(_nw147).ArraySet1(_1024_v34, 18)
						(_nw147).ArraySet1((func() _dafny.Int {
							if (_1028_v37).Contains((_dafny.Zero).Minus(_1024_v34)) {
								return (_1028_v37).Multiplicity((_dafny.Zero).Minus(_1024_v34))
							}
							return _1024_v34
						})(), 19)
						(_nw147).ArraySet1((_1024_v34).Times(_1024_v34), 20)
						(_nw147).ArraySet1((_dafny.Zero).Minus(_1024_v34), 21)
						(_nw147).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_this).F10() {
								return _1036_v43
							}
							return _1036_v43
						})()).Cardinality()), 22)
						(_nw147).ArraySet1(_1024_v34, 23)
						(_nw147).ArraySet1(_dafny.IntOfInt64(236), 24)
						_1037_v44 = _nw147
						var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1037_v44), 0))
						_ = _index159
						(_1037_v44).ArraySet1(Companion_Default___.SafeModuloInt(_1024_v34, _1024_v34), (_index159).Int())
						var _1038_v45 _dafny.CodePoint
						_ = _1038_v45
						_1038_v45 = _dafny.CodePoint('u')
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1037_v44), 0))
						_ = _index160
						(_1037_v44).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1023_v33, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.IntOfUint32((_1023_v33).Cardinality()))).Uint32(), _1038_v45)).Cardinality()))).Minus((_dafny.Zero).Minus(_1024_v34)), (_index160).Int())
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1035_v42), 0))
						_ = _index161
						(_1035_v42).ArraySet1((_this).F10(), (_index161).Int())
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1035_v42), 0))
						_ = _index162
						var _rhs162 bool = (Companion_Default___.Fm46((_this).F10(), (_983_v0).Select((Companion_Default___.SafeIndex(_1024_v34, _dafny.IntOfUint32((_983_v0).Cardinality()))).Uint32()).(bool), !((_this).F10()), globalState)).IsSubsetOf(_1028_v37)
						_ = _rhs162
						var _rhs163 bool = (_this).F10()
						_ = _rhs163
						var _lhs115 _dafny.Array = _1035_v42
						_ = _lhs115
						var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1035_v42), 0))
						_ = _lhs116
						var _lhs117 *GlobalState = globalState
						_ = _lhs117
						(_lhs115).ArraySet1(_rhs162, (_lhs116).Int())
						_lhs117.F8 = _rhs163
						var _1039_v46 _dafny.Map
						_ = _1039_v46
						_1039_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1024_v34, _1024_v34)
						_1039_v46 = (_1039_v46).Update(_dafny.IntOfInt64(154), (_1037_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1037_v44), 0))).Int()).(_dafny.Int))
					}
					var _1040_v47 _dafny.Array
					_ = _1040_v47
					var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
					_ = _nw148
					_1040_v47 = _nw148
					var _1041_v48 _dafny.Int
					_ = _1041_v48
					_1041_v48 = _dafny.IntOfInt64(915)
					var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1040_v47), 0))
					_ = _index163
					(_1040_v47).ArraySet1(_1041_v48, (_index163).Int())
					var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1040_v47), 0))
					_ = _index164
					(_1040_v47).ArraySet1(Companion_Default___.SafeDivisionInt((_1041_v48).Minus(_1041_v48), _1041_v48), (_index164).Int())
					var _1042_v49 _dafny.Array
					_ = _1042_v49
					var _nw149 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
					_ = _nw149
					_1042_v49 = _nw149
					var _1043_v50 _dafny.Array
					_ = _1043_v50
					var _nw150 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
					_ = _nw150
					_1043_v50 = _nw150
					var _1044_v51 T1
					_ = _1044_v51
					var _nw151 *C2 = New_C2_()
					_ = _nw151
					_nw151.Ctor__(!((_this).F10()), _1043_v50, (_this).F9(), (_this).F10())
					_1044_v51 = _nw151
					var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1042_v49), 0))
					_ = _index165
					(_1042_v49).ArraySet1(_1044_v51, (_index165).Int())
					var _1045_v52 _dafny.Map
					_ = _1045_v52
					_1045_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-785))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg75 _dafny.Int) interface{} {
							return coer75(arg75)
						}
					}(func(_1046_i7 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("krxj")).Cardinality())
					})), _1044_v51)
					var _1047_v54 _dafny.Set
					_ = _1047_v54
					_1047_v54 = _dafny.SetOf((_1044_v51).F10())
					var _1048_v55 _dafny.Sequence
					_ = _1048_v55
					_1048_v55 = _dafny.SeqOf((func() _dafny.Set {
						var _coll44 = _dafny.NewBuilder()
						_ = _coll44
						for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(525), _dafny.IntOfInt64(392))); ; {
							_compr_44, _ok47 := _iter47()
							if !_ok47 {
								break
							}
							var _1049_v53 _dafny.Int
							_1049_v53 = interface{}(_compr_44).(_dafny.Int)
							if ((_dafny.IntOfInt64(525)).Cmp(_1049_v53) <= 0) && ((_1049_v53).Cmp(_dafny.IntOfInt64(392)) < 0) {
								_coll44.Add((_1049_v53).Times(_1041_v48))
							}
						}
						return _coll44.ToSet()
					}()).Cardinality(), (_1047_v54).Cardinality())
					var _1050_v56 _dafny.Sequence
					_ = _1050_v56
					_1050_v56 = _dafny.SeqOf(_1048_v55)
					var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(479), _dafny.ArrayLen((_1042_v49), 0))
					_ = _index166
					(_1042_v49).ArraySet1((func() T1 {
						if (_1045_v52).Contains((_1050_v56).Select((Companion_Default___.SafeIndex((_1040_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1040_v47), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1050_v56).Cardinality()))).Uint32()).(_dafny.Sequence)) {
							return (_1045_v52).Get((_1050_v56).Select((Companion_Default___.SafeIndex((_1040_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1040_v47), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1050_v56).Cardinality()))).Uint32()).(_dafny.Sequence)).(T1)
						}
						return _1044_v51
					})(), (_index166).Int())
					var _1051_v57 _dafny.Array
					_ = _1051_v57
					var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
					_ = _nw152
					_1051_v57 = _nw152
					var _1052_v58 *C3
					_ = _1052_v58
					var _nw153 *C3 = New_C3_()
					_ = _nw153
					_nw153.Ctor__()
					_1052_v58 = _nw153
					var _1053_v59 _dafny.Sequence
					_ = _1053_v59
					_1053_v59 = _dafny.SeqOf(_1052_v58)
					var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1051_v57), 0))
					_ = _index167
					(_1051_v57).ArraySet1(_1053_v59, (_index167).Int())
					var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1051_v57), 0))
					_ = _index168
					(_1051_v57).ArraySet1(_1053_v59, (_index168).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		if (_this).F10() {
			var _1054_v60 _dafny.Sequence
			_ = _1054_v60
			_1054_v60 = _dafny.UnicodeSeqOfUtf8Bytes("tndwyncs")
			var _1055_v61 _dafny.MultiSet
			_ = _1055_v61
			_1055_v61 = _dafny.MultiSetOf(_1054_v60)
			var _1056_v62 _dafny.Int
			_ = _1056_v62
			_1056_v62 = _dafny.IntOfInt64(469)
			var _1057_v63 D11
			_ = _1057_v63
			_1057_v63 = Companion_D11_.Create_DC27_((_1055_v61).Update(_1054_v60, Companion_Default___.Abs(_1056_v62)))
			if (_dafny.MultiSetOf(_1054_v60)).IsSubsetOf((_1057_v63).Dtor_cf47()) {
				_1056_v62 = _1056_v62
				var _1058_v64 _dafny.Array
				_ = _1058_v64
				var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw154
				_1058_v64 = _nw154
				var _1059_v65 _dafny.Set
				_ = _1059_v65
				_1059_v65 = _dafny.SetOf((_this).F10(), false)
				var _1060_v66 D8
				_ = _1060_v66
				_1060_v66 = Companion_D8_.Create_DC24_(!((_this).F10()), _dafny.SetOf(_1058_v64), (_1059_v65).Cardinality())
				var _pat_let_tv22 = globalState
				_ = _pat_let_tv22
				var _1061_v67 _dafny.Array
				_ = _1061_v67
				var _nwElement0_31 D8 = _1060_v66
				_ = _nwElement0_31
				var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(8))
				_ = _nw155
				(_nw155).ArraySet1(_nwElement0_31, 0)
				(_nw155).ArraySet1(_1060_v66, 1)
				(_nw155).ArraySet1(_1060_v66, 2)
				(_nw155).ArraySet1(_1060_v66, 3)
				(_nw155).ArraySet1(_1060_v66, 4)
				(_nw155).ArraySet1(_1060_v66, 5)
				(_nw155).ArraySet1(_1060_v66, 6)
				(_nw155).ArraySet1(func(_pat_let20_0 D8) D8 {
					return func(_1062_dt__update__tmp_h0 D8) D8 {
						return func(_pat_let21_0 _dafny.Int) D8 {
							return func(_1063_dt__update_hcf44_h0 _dafny.Int) D8 {
								return Companion_D8_.Create_DC24_((_1062_dt__update__tmp_h0).Dtor_cf42(), (_1062_dt__update__tmp_h0).Dtor_cf43(), _1063_dt__update_hcf44_h0)
							}(_pat_let21_0)
						}((_this).Fm1((_this).F9(), _pat_let_tv22))
					}(_pat_let20_0)
				}(_1060_v66), 7)
				_1061_v67 = _nw155
				var _1064_v68 _dafny.Sequence
				_ = _1064_v68
				_1064_v68 = _dafny.SeqOf(_1061_v67, _1061_v67)
				var _1065_v69 D12
				_ = _1065_v69
				_1065_v69 = Companion_D12_.Create_DC30_(_1061_v67)
				var _1066_v70 _dafny.Array
				_ = _1066_v70
				var _nwElement0_32 _dafny.Array = _1061_v67
				_ = _nwElement0_32
				var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(26))
				_ = _nw156
				(_nw156).ArraySet1(_nwElement0_32, 0)
				(_nw156).ArraySet1(_1061_v67, 1)
				(_nw156).ArraySet1(_1061_v67, 2)
				(_nw156).ArraySet1(_1061_v67, 3)
				(_nw156).ArraySet1(_1061_v67, 4)
				(_nw156).ArraySet1(_1061_v67, 5)
				(_nw156).ArraySet1(_1061_v67, 6)
				(_nw156).ArraySet1(_1061_v67, 7)
				(_nw156).ArraySet1(_1061_v67, 8)
				(_nw156).ArraySet1((_1064_v68).Select((Companion_Default___.SafeIndex(_1056_v62, _dafny.IntOfUint32((_1064_v68).Cardinality()))).Uint32()).(_dafny.Array), 9)
				(_nw156).ArraySet1((func() _dafny.Array {
					if (_this).F10() {
						return _1061_v67
					}
					return _1061_v67
				})(), 10)
				(_nw156).ArraySet1((func() _dafny.Array {
					if false {
						return _1061_v67
					}
					return _1061_v67
				})(), 11)
				(_nw156).ArraySet1(_1061_v67, 12)
				(_nw156).ArraySet1((_1065_v69).Dtor_cf54(), 13)
				(_nw156).ArraySet1((func() _dafny.Array {
					if (_this).F10() {
						return _1061_v67
					}
					return _1061_v67
				})(), 14)
				(_nw156).ArraySet1(_1061_v67, 15)
				(_nw156).ArraySet1(_1061_v67, 16)
				(_nw156).ArraySet1(_1061_v67, 17)
				(_nw156).ArraySet1(_1061_v67, 18)
				(_nw156).ArraySet1(_1061_v67, 19)
				(_nw156).ArraySet1(_1061_v67, 20)
				(_nw156).ArraySet1(_1061_v67, 21)
				(_nw156).ArraySet1((func() _dafny.Array {
					if (_this).F10() {
						return _1061_v67
					}
					return _1061_v67
				})(), 22)
				(_nw156).ArraySet1(_1061_v67, 23)
				(_nw156).ArraySet1(_1061_v67, 24)
				(_nw156).ArraySet1(_1061_v67, 25)
				_1066_v70 = _nw156
				_1066_v70 = _1066_v70
				(globalState).F8 = (_this).F10()
				var _1067_v71 _dafny.Array
				_ = _1067_v71
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw157
				_1067_v71 = _nw157
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1067_v71), 0))
				_ = _index169
				(_1067_v71).ArraySet1((_1056_v62).Cmp(_1056_v62) <= 0, (_index169).Int())
				var _1068_v72 _dafny.Map
				_ = _1068_v72
				_1068_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v60, (_this).F10())
				var _1069_v73 D1
				_ = _1069_v73
				_1069_v73 = Companion_D1_.Create_DC4_((_this).F10(), (_this).F10(), _1056_v62)
				var _1070_v74 _dafny.Map
				_ = _1070_v74
				_1070_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1056_v62, (_1069_v73).Dtor_cf8())
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1067_v71), 0))
				_ = _index170
				var _rhs164 bool = (func() bool {
					if (_1068_v72).Contains(_1054_v60) {
						return (_1068_v72).Get(_1054_v60).(bool)
					}
					return (_this).F10()
				})()
				_ = _rhs164
				var _rhs165 bool = !((_this).F10())
				_ = _rhs165
				var _rhs166 bool = (_this).F10()
				_ = _rhs166
				var _rhs167 bool = !((((_1070_v74).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1056_v62, _1056_v62))).Cardinality()).Cmp(_1056_v62) >= 0)
				_ = _rhs167
				var _lhs118 *GlobalState = globalState
				_ = _lhs118
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				var _lhs120 _dafny.Array = _1067_v71
				_ = _lhs120
				var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1067_v71), 0))
				_ = _lhs121
				var _lhs122 *GlobalState = globalState
				_ = _lhs122
				_lhs118.F8 = _rhs164
				_lhs119.F8 = _rhs165
				(_lhs120).ArraySet1(_rhs166, (_lhs121).Int())
				_lhs122.F8 = _rhs167
				var _1071_v75 _dafny.MultiSet
				_ = _1071_v75
				_1071_v75 = _dafny.MultiSetOf((_1067_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1067_v71), 0))).Int()).(bool), (_this).F10())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_1058_v64), 0))
				_ = _index171
				(_1058_v64).ArraySet1((func() _dafny.Int {
					if (_1071_v75).Contains((_this).Fm0(_983_v0, (_this).F10(), _1056_v62, (_this).F10(), globalState)) {
						return (_1071_v75).Multiplicity((_this).Fm0(_983_v0, (_this).F10(), _1056_v62, (_this).F10(), globalState))
					}
					return _dafny.IntOfInt64(819)
				})(), (_index171).Int())
				var _1072_v76 _dafny.CodePoint
				_ = _1072_v76
				_1072_v76 = _dafny.CodePoint('a')
				var _1073_v77 _dafny.Map
				_ = _1073_v77
				_1073_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1058_v64, _1056_v62)
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_1058_v64), 0))
				_ = _index172
				(_1058_v64).ArraySet1(Companion_Default___.SafeModuloInt((Companion_D6_.Create_DC18_(_1072_v76, _dafny.IntOfUint32((_1054_v60).Cardinality()), _dafny.IntOfInt64(142))).Dtor_cf36(), ((_1073_v77).Merge(_1073_v77)).Cardinality()), (_index172).Int())
			} else {
				var _1074_v78 _dafny.Array
				_ = _1074_v78
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_24
				var _nw158 _dafny.Array
				_ = _nw158
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw158 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) bool = func(_1075_i8 _dafny.Int) bool {
						return (_this).F10()
					}
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw158 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw158).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw158).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1074_v78 = _nw158
				var _1076_v79 T1
				_ = _1076_v79
				var _nw159 *C2 = New_C2_()
				_ = _nw159
				_nw159.Ctor__((_this).F10(), (Companion_D5_.Create_DC13_(_1074_v78)).Dtor_cf24(), (_this).F9(), (_this).F10())
				_1076_v79 = _nw159
				var _nw160 *C2 = New_C2_()
				_ = _nw160
				_nw160.Ctor__((_this).Fm0(_983_v0, (_1076_v79).F10(), _1056_v62, true, globalState), _1074_v78, (_1076_v79).F9(), (_1076_v79).F10())
				_1076_v79 = _nw160
				var _1077_v80 _dafny.Array
				_ = _1077_v80
				var _nwElement0_33 _dafny.Sequence = _983_v0
				_ = _nwElement0_33
				var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(7))
				_ = _nw161
				(_nw161).ArraySet1(_nwElement0_33, 0)
				(_nw161).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_983_v0, _dafny.SeqOf((_this).F10(), (_1076_v79).F10())), 1)
				(_nw161).ArraySet1(_983_v0, 2)
				(_nw161).ArraySet1(_dafny.Companion_Sequence_.Update(_983_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfInt64(-31)), _dafny.IntOfUint32((_983_v0).Cardinality()))).Uint32(), !((_1076_v79).F10())), 3)
				(_nw161).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_983_v0, _983_v0), 4)
				(_nw161).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_983_v0, _983_v0), 5)
				(_nw161).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_983_v0, _983_v0), 6)
				_1077_v80 = _nw161
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1077_v80), 0))
				_ = _index173
				(_1077_v80).ArraySet1(_983_v0, (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1077_v80), 0))
				_ = _index174
				var _rhs168 _dafny.Sequence = _983_v0
				_ = _rhs168
				var _rhs169 _dafny.Sequence = _1054_v60
				_ = _rhs169
				var _rhs170 bool = (func() bool {
					if (_this).F10() {
						return false
					}
					return (_1076_v79).F10()
				})()
				_ = _rhs170
				var _lhs123 _dafny.Array = _1077_v80
				_ = _lhs123
				var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1077_v80), 0))
				_ = _lhs124
				var _lhs125 *GlobalState = globalState
				_ = _lhs125
				(_lhs123).ArraySet1(_rhs168, (_lhs124).Int())
				_1054_v60 = _rhs169
				_lhs125.F8 = _rhs170
				var _1078_v81 _dafny.Sequence
				_ = _1078_v81
				_1078_v81 = _dafny.SeqOf((_1077_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1077_v80), 0))).Int()).(_dafny.Sequence))
				var _1079_v82 *C1
				_ = _1079_v82
				var _nw162 *C1 = New_C1_()
				_ = _nw162
				_nw162.Ctor__(!((_this).F10()), _1054_v60)
				_1079_v82 = _nw162
				var _1080_v83 _dafny.CodePoint
				_ = _1080_v83
				_1080_v83 = _dafny.CodePoint('b')
				var _rhs171 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_983_v0, (_1078_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1079_v82, _1079_v82)).Cardinality()), _dafny.IntOfUint32((_1078_v81).Cardinality()))).Uint32()).(_dafny.Sequence)), _dafny.Companion_Sequence_.Concatenate((_1077_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1077_v80), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm38((_1079_v82).F12(), globalState)))
				_ = _rhs171
				var _rhs172 bool = !((_this).Fm0((_1077_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1077_v80), 0))).Int()).(_dafny.Sequence), (_1076_v79).F10(), _1056_v62, (_1076_v79).F10(), globalState))
				_ = _rhs172
				var _rhs173 bool = (_1079_v82).Fm32(_1056_v62, _1080_v83, globalState)
				_ = _rhs173
				var _lhs126 *GlobalState = globalState
				_ = _lhs126
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 *GlobalState = globalState
				_ = _lhs128
				_lhs126.F8 = _rhs171
				_lhs127.F8 = _rhs172
				_lhs128.F8 = _rhs173
				var _1081_v84 _dafny.Array
				_ = _1081_v84
				var _nwElement0_34 _dafny.CodePoint = _1080_v83
				_ = _nwElement0_34
				var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(13))
				_ = _nw163
				(_nw163).ArraySet1CodePoint(_nwElement0_34, 0)
				(_nw163).ArraySet1CodePoint(_1080_v83, 1)
				(_nw163).ArraySet1CodePoint(_1080_v83, 2)
				(_nw163).ArraySet1CodePoint(_1080_v83, 3)
				(_nw163).ArraySet1CodePoint(_1080_v83, 4)
				(_nw163).ArraySet1CodePoint(_1080_v83, 5)
				(_nw163).ArraySet1CodePoint(_1080_v83, 6)
				(_nw163).ArraySet1CodePoint(_1080_v83, 7)
				(_nw163).ArraySet1CodePoint(_1080_v83, 8)
				(_nw163).ArraySet1CodePoint(_1080_v83, 9)
				(_nw163).ArraySet1CodePoint(_1080_v83, 10)
				(_nw163).ArraySet1CodePoint(_1080_v83, 11)
				(_nw163).ArraySet1CodePoint(_1080_v83, 12)
				_1081_v84 = _nw163
				var _1082_v85 _dafny.Array
				_ = _1082_v85
				var _nwElement0_35 _dafny.Array = _1081_v84
				_ = _nwElement0_35
				var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(5))
				_ = _nw164
				(_nw164).ArraySet1(_nwElement0_35, 0)
				(_nw164).ArraySet1(_1081_v84, 1)
				(_nw164).ArraySet1(_1081_v84, 2)
				(_nw164).ArraySet1(_1081_v84, 3)
				(_nw164).ArraySet1(_1081_v84, 4)
				_1082_v85 = _nw164
				var _1083_v86 _dafny.Map
				_ = _1083_v86
				_1083_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1082_v85, (func() bool {
					if (_1079_v82).F12() {
						return (_1076_v79).F10()
					}
					return (_1079_v82).F12()
				})())
				var _1084_v87 _dafny.MultiSet
				_ = _1084_v87
				_1084_v87 = _dafny.MultiSetOf(_1080_v83)
				_1083_v86 = (_1083_v86).Update(_1082_v85, (_1084_v87).IsSubsetOf(_1084_v87))
				var _1085_v88 D5
				_ = _1085_v88
				_1085_v88 = Companion_D5_.Create_DC13_(_1074_v78)
				var _1086_v89 _dafny.Map
				_ = _1086_v89
				_1086_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1079_v82).F12(), (_1085_v88).Dtor_cf24())
				var _1087_v90 _dafny.Map
				_ = _1087_v90
				_1087_v90 = _1086_v89
				_1087_v90 = _1087_v90
			}
			var _1088_v91 *C3
			_ = _1088_v91
			var _nw165 *C3 = New_C3_()
			_ = _nw165
			_nw165.Ctor__()
			_1088_v91 = _nw165
			var _1089_v92 _dafny.MultiSet
			_ = _1089_v92
			_1089_v92 = _dafny.MultiSetOf(_1088_v91, _1088_v91)
			var _1090_v93 _dafny.MultiSet
			_ = _1090_v93
			_1090_v93 = _dafny.MultiSetOf((_this).F10(), (_this).F10(), (_this).F10())
			var _1091_v94 _dafny.Map
			_ = _1091_v94
			_1091_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1090_v93)
			var _rhs174 _dafny.MultiSet = _1089_v92
			_ = _rhs174
			var _rhs175 _dafny.Map = _1091_v94
			_ = _rhs175
			_1089_v92 = _rhs174
			_1091_v94 = _rhs175
			(globalState).F8 = (_this).F10()
			var _1092_v95 _dafny.Sequence
			_ = _1092_v95
			var _1093_v96 bool
			_ = _1093_v96
			var _1094_v97 _dafny.Sequence
			_ = _1094_v97
			var _1095_v98 _dafny.Sequence
			_ = _1095_v98
			var _out52 _dafny.Sequence
			_ = _out52
			var _out53 bool
			_ = _out53
			var _out54 _dafny.Sequence
			_ = _out54
			var _out55 _dafny.Sequence
			_ = _out55
			_out52, _out53, _out54, _out55 = (_1088_v91).M0(globalState)
			_1092_v95 = _out52
			_1093_v96 = _out53
			_1094_v97 = _out54
			_1095_v98 = _out55
			_1056_v62 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm45(_1093_v96, _1056_v62, _1093_v96, globalState), _dafny.Companion_Sequence_.Concatenate(_1054_v60, _1095_v98))).Cardinality())
		} else {
			var _1096_v99 _dafny.Int
			_ = _1096_v99
			_1096_v99 = _dafny.IntOfInt64(295)
			(globalState).F8 = ((func() _dafny.Int {
				if (_this).F10() {
					return _1096_v99
				}
				return (_dafny.Zero).Minus(_1096_v99)
			})()).Cmp((_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10()), globalState)) >= 0
			var _1097_v100 *C3
			_ = _1097_v100
			var _nw166 *C3 = New_C3_()
			_ = _nw166
			_nw166.Ctor__()
			_1097_v100 = _nw166
			var _1098_v101 _dafny.Sequence
			_ = _1098_v101
			_1098_v101 = _dafny.SeqOf(_1097_v100)
			var _1099_v102 _dafny.Map
			_ = _1099_v102
			_1099_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1097_v100)
			var _1100_v103 _dafny.Array
			_ = _1100_v103
			var _nwElement0_36 *C3 = _1097_v100
			_ = _nwElement0_36
			var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(15))
			_ = _nw167
			(_nw167).ArraySet1(_nwElement0_36, 0)
			(_nw167).ArraySet1(_1097_v100, 1)
			(_nw167).ArraySet1((_1098_v101).Select((Companion_Default___.SafeIndex(_1096_v99, _dafny.IntOfUint32((_1098_v101).Cardinality()))).Uint32()).(*C3), 2)
			(_nw167).ArraySet1(_1097_v100, 3)
			(_nw167).ArraySet1(_1097_v100, 4)
			(_nw167).ArraySet1(_1097_v100, 5)
			(_nw167).ArraySet1(_1097_v100, 6)
			(_nw167).ArraySet1((func() *C3 {
				if (_this).F10() {
					return _1097_v100
				}
				return _1097_v100
			})(), 7)
			(_nw167).ArraySet1(_1097_v100, 8)
			(_nw167).ArraySet1(_1097_v100, 9)
			(_nw167).ArraySet1((func() *C3 {
				if (_1099_v102).Contains((_this).F10()) {
					return (_1099_v102).Get((_this).F10()).(*C3)
				}
				return _1097_v100
			})(), 10)
			(_nw167).ArraySet1(_1097_v100, 11)
			(_nw167).ArraySet1(_1097_v100, 12)
			(_nw167).ArraySet1(_1097_v100, 13)
			(_nw167).ArraySet1(_1097_v100, 14)
			_1100_v103 = _nw167
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1100_v103), 0))
			_ = _index175
			(_1100_v103).ArraySet1(_1097_v100, (_index175).Int())
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1100_v103), 0))
			_ = _index176
			(_1100_v103).ArraySet1(_1097_v100, (_index176).Int())
			var _1101_v104 _dafny.Sequence
			_ = _1101_v104
			_1101_v104 = _dafny.UnicodeSeqOfUtf8Bytes("smnhcyhf")
			(globalState).F8 = _dafny.Companion_Sequence_.IsPrefixOf(_1101_v104, _1101_v104)
			var _1102_v105 *C0
			_ = _1102_v105
			var _nw168 *C0 = New_C0_()
			_ = _nw168
			_nw168.Ctor__()
			_1102_v105 = _nw168
			var _1103_v106 _dafny.CodePoint
			_ = _1103_v106
			_1103_v106 = _dafny.CodePoint('s')
			_1103_v106 = _1103_v106
		}
		var _1104_v107 _dafny.Int
		_ = _1104_v107
		_1104_v107 = _dafny.IntOfInt64(521)
		_1104_v107 = (_dafny.Zero).Minus(_1104_v107)
	}
}
func (_this *C4) M6(p0 _dafny.Int, p1 _dafny.Array, p2 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, _dafny.Sequence, _dafny.Int) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1105_v0 _dafny.Map
		_ = _1105_v0
		_1105_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg76 _dafny.Int) interface{} {
				return coer76(arg76)
			}
		}(func(_1106_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), ((_dafny.Zero).Minus(p2)).Plus(p0))
		var _1107_v1 _dafny.Sequence
		_ = _1107_v1
		_1107_v1 = _dafny.UnicodeSeqOfUtf8Bytes("veytmuod")
		_1105_v0 = (_1105_v0).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kxkg"), _dafny.Companion_Sequence_.Update(_1107_v1, (Companion_Default___.SafeIndex((_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F10()), (_this).F10()), globalState), _dafny.IntOfUint32((_1107_v1).Cardinality()))).Uint32(), _dafny.CodePoint('u'))), p2)
		if !((func() bool {
			if (_this).F10() {
				return (_this).F10()
			}
			return (_this).F10()
		})()) {
			(globalState).F8 = (_this).F10()
			var _1108_v2 D0
			_ = _1108_v2
			_1108_v2 = Companion_D0_.Create_DC0_(false)
			var _1109_v3 _dafny.Sequence
			_ = _1109_v3
			_1109_v3 = _dafny.SeqOf((_this).F10(), (_this).F10())
			var _1110_v4 D8
			_ = _1110_v4
			_1110_v4 = Companion_D8_.Create_DC23_(_1109_v3)
			var _1111_v5 _dafny.Map
			_ = _1111_v5
			_1111_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1108_v2, _1110_v4)
			r3 = ((_1111_v5).Update((func() D0 {
				if true {
					return _1108_v2
				}
				return _1108_v2
			})(), _1110_v4)).Cardinality()
			var _1112_v6 _dafny.CodePoint
			_ = _1112_v6
			_1112_v6 = _dafny.CodePoint('s')
			_1112_v6 = _1112_v6
			var _1113_v7 _dafny.Map
			_ = _1113_v7
			_1113_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F10()), p2)
			_1113_v7 = (_1113_v7).Update((_this).F10(), _dafny.IntOfUint32((Companion_Default___.Fm47((_this).F10(), globalState)).Cardinality()))
			var _1114_v8 D6
			_ = _1114_v8
			_1114_v8 = Companion_D6_.Create_DC16_((_this).F9())
			_1114_v8 = Companion_D6_.Create_DC16_(Companion_Default___.Fm48((_this).F10(), globalState))
		} else {
			var _1115_v9 *C1
			_ = _1115_v9
			var _nw169 *C1 = New_C1_()
			_ = _nw169
			_nw169.Ctor__((_this).F10(), _1107_v1)
			_1115_v9 = _nw169
			var _1116_v10 _dafny.Sequence
			_ = _1116_v10
			_1116_v10 = _dafny.SeqOf(!(true), (_1115_v9).F12())
			var _1117_v11 D6
			_ = _1117_v11
			_1117_v11 = Companion_D6_.Create_DC17_((_1115_v9).F12(), (_this).F10(), _dafny.IntOfUint32(((_1115_v9).F13()).Cardinality()), _dafny.IntOfUint32(((_1115_v9).F13()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(p2, _dafny.IntOfUint32((Companion_Default___.Fm45((_1115_v9).F12(), p2, (_1115_v9).F12(), globalState)).Cardinality()), _dafny.IntOfUint32((_1116_v10).Cardinality()), p0, p0)).Cardinality()))
			var _1118_v12 D6
			_ = _1118_v12
			_1118_v12 = Companion_D6_.Create_DC20_(_1117_v11)
			var _source17 D6 = _1118_v12
			_ = _source17
			if _source17.Is_DC17() {
				var _1119___mcc_h0 bool = _source17.Get_().(D6_DC17).Cf29
				_ = _1119___mcc_h0
				var _1120___mcc_h1 bool = _source17.Get_().(D6_DC17).Cf30
				_ = _1120___mcc_h1
				var _1121___mcc_h2 _dafny.Int = _source17.Get_().(D6_DC17).Cf31
				_ = _1121___mcc_h2
				var _1122___mcc_h3 _dafny.Int = _source17.Get_().(D6_DC17).Cf32
				_ = _1122___mcc_h3
				var _1123___mcc_h4 _dafny.Int = _source17.Get_().(D6_DC17).Cf33
				_ = _1123___mcc_h4
				var _1124_cf33 _dafny.Int = _1123___mcc_h4
				_ = _1124_cf33
				var _1125_cf32 _dafny.Int = _1122___mcc_h3
				_ = _1125_cf32
				var _1126_cf31 _dafny.Int = _1121___mcc_h2
				_ = _1126_cf31
				var _1127_cf30 bool = _1120___mcc_h1
				_ = _1127_cf30
				var _1128_cf29 bool = _1119___mcc_h0
				_ = _1128_cf29
				var _1129_v13 _dafny.Array
				_ = _1129_v13
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_25
				var _nw170 _dafny.Array
				_ = _nw170
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw170 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) bool = (func(_1130_cf29 bool) func(_dafny.Int) bool {
						return func(_1131_i1 _dafny.Int) bool {
							return _1130_cf29
						}
					})(_1128_cf29)
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw170 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw170).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw170).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1129_v13 = _nw170
				var _1132_v14 _dafny.Map
				_ = _1132_v14
				_1132_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1127_cf30), _1129_v13)
				var _1133_v15 _dafny.Map
				_ = _1133_v15
				_1133_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1124_cf33, _1128_cf29)
				_1132_v14 = (_1132_v14).Update((func() bool {
					if (_1133_v15).Contains((_dafny.Zero).Minus(_1126_cf31)) {
						return (_1133_v15).Get((_dafny.Zero).Minus(_1126_cf31)).(bool)
					}
					return _1127_cf30
				})(), _1129_v13)
				(globalState).F8 = (_1115_v9).F12()
				var _1134_v16 _dafny.Map
				_ = _1134_v16
				_1134_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_1115_v9).F13()).Cardinality()), Companion_Default___.Fm33((_1115_v9).F12(), _1128_cf29, _dafny.IntOfInt64(-327), _1128_cf29, globalState))
				var _1135_v17 _dafny.Sequence
				_ = _1135_v17
				_1135_v17 = _dafny.SeqOf(_1134_v16, _1134_v16, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1107_v1).Cardinality())))
				var _1136_v18 _dafny.Set
				_ = _1136_v18
				_1136_v18 = _dafny.SetOf((_1115_v9).F12(), false, (_this).Fm0(_1116_v10, _1128_cf29, _1126_cf31, (_1115_v9).F12(), globalState), !((_1115_v9).F12()), true)
				var _1137_v19 _dafny.Map
				_ = _1137_v19
				_1137_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_v18, _1125_cf32)
				var _1138_v20 _dafny.CodePoint
				_ = _1138_v20
				_1138_v20 = _dafny.CodePoint('a')
				var _1139_v21 D4
				_ = _1139_v21
				_1139_v21 = Companion_D4_.Create_DC12_(_1127_cf30, _dafny.IntOfUint32(((_1115_v9).F13()).Cardinality()), _1138_v20, (_dafny.Zero).Minus(_1126_cf31), _1124_cf33)
				var _pat_let_tv23 = _1128_cf29
				_ = _pat_let_tv23
				var _pat_let_tv24 = p2
				_ = _pat_let_tv24
				_1127_cf30 = ((((_1135_v17).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1135_v17).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()).Times((_1137_v19).Cardinality())).Cmp((func(_pat_let22_0 D4) D4 {
					return func(_1140_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let23_0 bool) D4 {
							return func(_1141_dt__update_hcf19_h0 bool) D4 {
								return func(_pat_let24_0 _dafny.Int) D4 {
									return func(_1142_dt__update_hcf23_h0 _dafny.Int) D4 {
										return Companion_D4_.Create_DC12_(_1141_dt__update_hcf19_h0, (_1140_dt__update__tmp_h0).Dtor_cf20(), (_1140_dt__update__tmp_h0).Dtor_cf21(), (_1140_dt__update__tmp_h0).Dtor_cf22(), _1142_dt__update_hcf23_h0)
									}(_pat_let24_0)
								}(_pat_let_tv24)
							}(_pat_let23_0)
						}(_pat_let_tv23)
					}(_pat_let22_0)
				}(_1139_v21)).Dtor_cf20()) == 0
				var _1143_v22 _dafny.Map
				_ = _1143_v22
				_1143_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (Companion_Default___.Fm33((_1115_v9).F12(), (_1115_v9).F12(), _1126_cf31, (_1115_v9).F12(), globalState)).Cmp(_1125_cf32) <= 0)
				_1143_v22 = _1143_v22
			} else if _source17.Is_DC18() {
				var _1144___mcc_h5 _dafny.CodePoint = _source17.Get_().(D6_DC18).Cf34
				_ = _1144___mcc_h5
				var _1145___mcc_h6 _dafny.Int = _source17.Get_().(D6_DC18).Cf35
				_ = _1145___mcc_h6
				var _1146___mcc_h7 _dafny.Int = _source17.Get_().(D6_DC18).Cf36
				_ = _1146___mcc_h7
				var _1147_cf36 _dafny.Int = _1146___mcc_h7
				_ = _1147_cf36
				var _1148_cf35 _dafny.Int = _1145___mcc_h6
				_ = _1148_cf35
				var _1149_cf34 _dafny.CodePoint = _1144___mcc_h5
				_ = _1149_cf34
				var _1150_v23 _dafny.Array
				_ = _1150_v23
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw171
				_1150_v23 = _nw171
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
				_ = _nw172
				_1150_v23 = _nw172
				var _1151_v25 _dafny.Array
				_ = _1151_v25
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_26
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Set = (func(_1152_v9 *C1, _1153_v10 _dafny.Sequence, _1154_p0 _dafny.Int, _1155_cf36 _dafny.Int, _1156_cf35 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_1157_i2 _dafny.Int) _dafny.Set {
							return (func() _dafny.Set {
								if (_1152_v9).F12() {
									return func() _dafny.Set {
										var _coll45 = _dafny.NewBuilder()
										_ = _coll45
										for _iter48 := _dafny.Iterate((_dafny.MultiSetOf(_1154_p0, _dafny.IntOfInt64(604), _1155_cf36, _dafny.IntOfInt64(916), _1156_cf35)).Elements()); ; {
											_compr_45, _ok48 := _iter48()
											if !_ok48 {
												break
											}
											var _1158_v24 _dafny.Int
											_1158_v24 = interface{}(_compr_45).(_dafny.Int)
											if (_dafny.MultiSetOf(_1154_p0, _dafny.IntOfInt64(604), _1155_cf36, _dafny.IntOfInt64(916), _1156_cf35)).Contains(_1158_v24) {
												_coll45.Add((_1158_v24).Minus(_dafny.IntOfUint32(((Companion_D8_.Create_DC23_(_dafny.SeqOf(false))).Dtor_cf41()).Cardinality())))
											}
										}
										return _coll45.ToSet()
									}()
								}
								return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_1153_v10), !((_1152_v9).F12()))).Cardinality())
							})()
						}
					})(_1115_v9, _1116_v10, p0, _1147_cf36, _1148_cf35)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw173 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw173).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw173).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1151_v25 = _nw173
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
				_ = _nw174
				_1151_v25 = _nw174
				(globalState).F8 = (p0).Cmp((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(331))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1159_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1160_i3 _dafny.Int) _dafny.Int {
						return _1159_p0
					}
				})(p0)))).Cardinality())).Minus(p2)) < 0
				_1107_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm22(_1148_cf35, (_1115_v9).F12(), (_this).F10(), p0, globalState), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.IntOfUint32((Companion_Default___.Fm22(_1148_cf35, (_1115_v9).F12(), (_this).F10(), p0, globalState)).Cardinality()))).Uint32(), _1149_cf34), (_1115_v9).F13()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1161_cf34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1162_i4 _dafny.Int) _dafny.CodePoint {
						return _1161_cf34
					}
				})(_1149_cf34))))
			} else if _source17.Is_DC19() {
				var _1163___mcc_h8 bool = _source17.Get_().(D6_DC19).Cf37
				_ = _1163___mcc_h8
				var _1164___mcc_h9 _dafny.CodePoint = _source17.Get_().(D6_DC19).Cf38
				_ = _1164___mcc_h9
				var _1165_cf38 _dafny.CodePoint = _1164___mcc_h9
				_ = _1165_cf38
				var _1166_cf37 bool = _1163___mcc_h8
				_ = _1166_cf37
				r1 = p2
				var _1167_v26 _dafny.Sequence
				_ = _1167_v26
				_1167_v26 = _dafny.SeqOf(p2)
				(globalState).F8 = ((p0).Cmp((_1167_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1116_v10).Cardinality()), _dafny.IntOfUint32((_1167_v26).Cardinality()))).Uint32()).(_dafny.Int)) != 0) || ((_1115_v9).Fm32((_this).Fm1(((_this).F9()).Update(_1166_cf37, (_1115_v9).F12()), globalState), _1165_cf38, globalState))
				var _1168_v27 *C1
				_ = _1168_v27
				var _nw175 *C1 = New_C1_()
				_ = _nw175
				_nw175.Ctor__((_this).F10(), _dafny.Companion_Sequence_.Update((_1115_v9).F13(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_1115_v9).F13()).Cardinality()))).Uint32(), _1165_cf38))
				_1168_v27 = _nw175
				(globalState).F8 = (p0).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-486))) == 0
			} else if _source17.Is_DC16() {
				var _1169___mcc_h10 _dafny.Map = _source17.Get_().(D6_DC16).Cf28
				_ = _1169___mcc_h10
				var _1170_cf28 _dafny.Map = _1169___mcc_h10
				_ = _1170_cf28
				var _1171_v28 D7
				_ = _1171_v28
				_1171_v28 = Companion_D7_.Create_DC22_()
				var _1172_v29 _dafny.Array
				_ = _1172_v29
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw176
				_1172_v29 = _nw176
				var _1173_v30 _dafny.Set
				_ = _1173_v30
				_1173_v30 = _dafny.SetOf((_this).F10(), (_1115_v9).F12())
				var _1174_v31 _dafny.Map
				_ = _1174_v31
				_1174_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F10())
				var _1175_v32 _dafny.Map
				_ = _1175_v32
				_1175_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), (_1115_v9).F12())
				var _1176_v33 _dafny.Map
				_ = _1176_v33
				_1176_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1175_v32, _1173_v30)
				var _1177_v34 _dafny.Sequence
				_ = _1177_v34
				_1177_v34 = _dafny.SeqOf(p1, p1)
				var _rhs176 D7 = _1171_v28
				_ = _rhs176
				var _rhs177 _dafny.Int = ((p2).Times(p2)).Minus(p2)
				_ = _rhs177
				var _rhs178 _dafny.Array = _1172_v29
				_ = _rhs178
				var _rhs179 _dafny.Set = (Companion_Default___.Fm34((_1174_v31).Cardinality(), (_this).F10(), (_1115_v9).F12(), (_1115_v9).F12(), globalState)).Union((func() _dafny.Set {
					if (_1176_v33).Contains(_1175_v32) {
						return (_1176_v33).Get(_1175_v32).(_dafny.Set)
					}
					return _1173_v30
				})())
				_ = _rhs179
				var _rhs180 _dafny.Int = _dafny.IntOfUint32((_1177_v34).Cardinality())
				_ = _rhs180
				_1171_v28 = _rhs176
				r1 = _rhs177
				_1172_v29 = _rhs178
				_1173_v30 = _rhs179
				r3 = _rhs180
				var _1178_v35 _dafny.Sequence
				_ = _1178_v35
				_1178_v35 = _dafny.SeqOf(p2)
				var _1179_v36 D1
				_ = _1179_v36
				_1179_v36 = Companion_D1_.Create_DC2_(_1178_v35)
				var _1180_v37 _dafny.Set
				_ = _1180_v37
				_1180_v37 = _dafny.SetOf(_1179_v36)
				var _1181_v38 _dafny.Map
				_ = _1181_v38
				_1181_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf((_1115_v9).F13(), (_1115_v9).F13()), _1180_v37)
				var _1182_v39 _dafny.CodePoint
				_ = _1182_v39
				_1182_v39 = _dafny.CodePoint('q')
				var _1183_v40 _dafny.Map
				_ = _1183_v40
				_1183_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1182_v39, _1182_v39)
				var _rhs181 bool = !(_1183_v40).Contains(_1182_v39)
				_ = _rhs181
				var _rhs182 _dafny.Map = (_1181_v38).Merge(_1181_v38)
				_ = _rhs182
				var _rhs183 bool = (_1173_v30).IsSubsetOf(_1173_v30)
				_ = _rhs183
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				_lhs129.F8 = _rhs181
				_1181_v38 = _rhs182
				_lhs130.F8 = _rhs183
				var _1184_v41 _dafny.Map
				_ = _1184_v41
				_1184_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus((_1174_v31).Cardinality())).Cmp(p0) != 0, (_dafny.Zero).Minus((p2).Plus(p0)))
				_1184_v41 = _1184_v41
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1172_v29), 0))
				_ = _index177
				(_1172_v29).ArraySet1((_1115_v9).F12(), (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((p1), 0))
				_ = _index178
				(p1).ArraySet1(p0, (_index178).Int())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1172_v29), 0))
				_ = _index179
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((p1), 0))
				_ = _index180
				var _rhs184 bool = (p2).Cmp(p0) < 0
				_ = _rhs184
				var _rhs185 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).Fm1(_1170_cf28, globalState)))
				_ = _rhs185
				var _rhs186 _dafny.Int = p2
				_ = _rhs186
				var _rhs187 _dafny.Int = p0
				_ = _rhs187
				var _rhs188 bool = (_1115_v9).F12()
				_ = _rhs188
				var _lhs131 _dafny.Array = _1172_v29
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1172_v29), 0))
				_ = _lhs132
				var _lhs133 _dafny.Array = p1
				_ = _lhs133
				var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((p1), 0))
				_ = _lhs134
				var _lhs135 *GlobalState = globalState
				_ = _lhs135
				(_lhs131).ArraySet1(_rhs184, (_lhs132).Int())
				r3 = _rhs185
				r3 = _rhs186
				(_lhs133).ArraySet1(_rhs187, (_lhs134).Int())
				_lhs135.F8 = _rhs188
			} else {
				var _1185___mcc_h11 D6 = _source17.Get_().(D6_DC20).Cf39
				_ = _1185___mcc_h11
				var _1186_cf39 D6 = _1185___mcc_h11
				_ = _1186_cf39
				var _1187_v42 _dafny.Set
				_ = _1187_v42
				_1187_v42 = _dafny.SetOf((_1115_v9).F12(), (_this).F10())
				var _1188_v43 D3
				_ = _1188_v43
				_1188_v43 = Companion_D3_.Create_DC8_((_1187_v42).IsSubsetOf(_1187_v42), (_this).F10())
				_1188_v43 = Companion_Default___.Fm35(false, globalState)
				var _1189_v44 _dafny.Map
				_ = _1189_v44
				_1189_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1115_v9).F12(), ((_this).F9()).Cardinality())
				var _1190_v45 _dafny.Set
				_ = _1190_v45
				_1190_v45 = _dafny.SetOf((_1189_v44).Cardinality())
				var _1191_v46 D4
				_ = _1191_v46
				_1191_v46 = Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_1116_v10).Cardinality()), _1190_v45, !((_1115_v9).F12()))
				var _1192_v47 _dafny.Map
				_ = _1192_v47
				_1192_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(309), p2)
				var _1193_v48 _dafny.Map
				_ = _1193_v48
				_1193_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1192_v47).Cardinality())
				var _pat_let_tv25 = _1190_v45
				_ = _pat_let_tv25
				var _1194_v50 _dafny.Sequence
				_ = _1194_v50
				_1194_v50 = _dafny.SeqOf(_1191_v46, func(_pat_let25_0 D4) D4 {
					return func(_1199_dt__update__tmp_h1 D4) D4 {
						return func(_pat_let26_0 _dafny.Set) D4 {
							return func(_1200_dt__update_hcf17_h0 _dafny.Set) D4 {
								return Companion_D4_.Create_DC11_((_1199_dt__update__tmp_h1).Dtor_cf16(), _1200_dt__update_hcf17_h0, (_1199_dt__update__tmp_h1).Dtor_cf18())
							}(_pat_let26_0)
						}(_pat_let_tv25)
					}(_pat_let25_0)
				}(Companion_D4_.Create_DC11_((_1192_v47).Cardinality(), func() _dafny.Set {
					var _coll46 = _dafny.NewBuilder()
					_ = _coll46
					for _iter49 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-534))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg79 _dafny.Int) interface{} {
							return coer79(arg79)
						}
					}((func(_1195_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1196_i5 _dafny.Int) _dafny.Int {
							return _1195_p0
						}
					})(p0)))).Elements()); ; {
						_compr_46, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1197_v49 _dafny.Int
						_1197_v49 = interface{}(_compr_46).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-534))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg80 _dafny.Int) interface{} {
								return coer80(arg80)
							}
						}((func(_1198_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1196_i5 _dafny.Int) _dafny.Int {
								return _1198_p0
							}
						})(p0))), _1197_v49) {
							_coll46.Add(Companion_Default___.SafeModuloInt(_1197_v49, _dafny.IntOfInt64(-700)))
						}
					}
					return _coll46.ToSet()
				}(), (_1115_v9).F12())), _1191_v46)
				var _1201_v51 _dafny.Sequence
				_ = _1201_v51
				_1201_v51 = _1194_v50
				var _1202_v52 _dafny.Map
				_ = _1202_v52
				_1202_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1194_v50, _dafny.IntOfInt64(-803))
				var _rhs189 bool = !((Companion_D14_.Create_DC34_(_1202_v52)).Dtor_cf61()).Contains((_1201_v51))
				_ = _rhs189
				var _rhs190 _dafny.Int = p2
				_ = _rhs190
				var _lhs136 *GlobalState = globalState
				_ = _lhs136
				_lhs136.F8 = _rhs189
				r1 = _rhs190
				_1107_v1 = _dafny.Companion_Sequence_.Concatenate((_1115_v9).F13(), _dafny.Companion_Sequence_.Concatenate(_1107_v1, (_1115_v9).F13()))
				var _1203_v53 _dafny.Map
				_ = _1203_v53
				_1203_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1201_v51, _1190_v45)
				_1189_v44 = (_1189_v44).Update((p0).Cmp((_1203_v53).Cardinality()) <= 0, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1107_v1).Cardinality()), p0))
			}
			var _1204_v54 *C3
			_ = _1204_v54
			var _nw177 *C3 = New_C3_()
			_ = _nw177
			_nw177.Ctor__()
			_1204_v54 = _nw177
			var _1205_v55 _dafny.Array
			_ = _1205_v55
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_27
			var _nw178 _dafny.Array
			_ = _nw178
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw178 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = (func(_1206_v10 _dafny.Sequence, _1207_p2 _dafny.Int) func(_dafny.Int) bool {
					return func(_1208_i6 _dafny.Int) bool {
						return (_dafny.IntOfUint32((_1206_v10).Cardinality())).Cmp(_1207_p2) < 0
					}
				})(_1116_v10, p2)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw178 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw178).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw178).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1205_v55 = _nw178
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_1205_v55), 0))
			_ = _index181
			(_1205_v55).ArraySet1((_1115_v9).F12(), (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_1205_v55), 0))
			_ = _index182
			(_1205_v55).ArraySet1((_1116_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1116_v10).Cardinality()))).Uint32()).(bool), (_index182).Int())
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((p1), 0))
			_ = _index183
			(p1).ArraySet1((_dafny.IntOfInt64(207)).Plus(p0), (_index183).Int())
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((p1), 0))
			_ = _index184
			(p1).ArraySet1(p2, (_index184).Int())
		}
		r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm45(true, p2, true, globalState), _dafny.Companion_Sequence_.Concatenate(_1107_v1, _1107_v1))).Cardinality())
		var _1209_v56 _dafny.Array
		_ = _1209_v56
		var _nw179 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw179
		_1209_v56 = _nw179
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1209_v56), 0))); ; {
			_guard_loop_3, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _1210_i7 _dafny.Int
			_1210_i7 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1210_i7).Sign() != -1) && ((_1210_i7).Cmp(_dafny.ArrayLen((_1209_v56), 0)) < 0)) {
				(_1209_v56).ArraySet1(!((_this).F10()), (_1210_i7).Int())
			}
		}
		var _1211_v57 _dafny.Map
		_ = _1211_v57
		_1211_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F10())
		var _1212_v58 _dafny.Sequence
		_ = _1212_v58
		_1212_v58 = _dafny.SeqOf((_this).F10())
		(globalState).F8 = (func() bool {
			if (_1211_v57).Contains(p0) {
				return (_1211_v57).Get(p0).(bool)
			}
			return _dafny.Companion_Sequence_.Contains(_1212_v58, (_this).F10())
		})()
		var _1213_v59 *C2
		_ = _1213_v59
		var _nw180 *C2 = New_C2_()
		_ = _nw180
		_nw180.Ctor__((_this).F10(), _1209_v56, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!((_this).F10())))), (_this).F10())
		_1213_v59 = _nw180
		r0 = _dafny.SeqOf(_dafny.IntOfInt64(831))
		var _1214_v60 _dafny.Sequence
		_ = _1214_v60
		_1214_v60 = _dafny.SeqOf(p0, p0)
		var _1215_v61 _dafny.Map
		_ = _1215_v61
		_1215_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
		var _1216_v62 _dafny.Map
		_ = _1216_v62
		_1216_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.Companion_Sequence_.Concatenate(_1214_v60, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.Zero).Minus(p0), (func() _dafny.Int {
			if (_1215_v61).Contains(_dafny.IntOfUint32((_1214_v60).Cardinality())) {
				return (_1215_v61).Get(_dafny.IntOfUint32((_1214_v60).Cardinality())).(_dafny.Int)
			}
			return p2
		})(), p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vyqvooyef")).Cardinality())), (Companion_Default___.SafeIndex((_1213_v59).Fm42(p0, p2, p2, globalState), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p0), (func() _dafny.Int {
			if (_1215_v61).Contains(_dafny.IntOfUint32((_1214_v60).Cardinality())) {
				return (_1215_v61).Get(_dafny.IntOfUint32((_1214_v60).Cardinality())).(_dafny.Int)
			}
			return p2
		})(), p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vyqvooyef")).Cardinality()))).Cardinality()))).Uint32(), p0)))
		r1 = (_1216_v62).Cardinality()
		r2 = Companion_Default___.Fm37(globalState)
		r3 = p2
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f9  _dafny.Map
	_f10 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f9 = _dafny.EmptyMap
	_this._f10 = false
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

func (_this *C5) F9() _dafny.Map {
	return _this._f9
}
func (_this *C5) F10() bool {
	return _this._f10
}
func (_this *C5) Ctor__(f9 _dafny.Map, f10 bool) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C5) Fm0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return !((Companion_D0_.Create_DC1_((_dafny.SetOf((_dafny.MultiSetOf((_this).F10())).Cardinality())).Cardinality(), (_this).F10(), _dafny.IntOfInt64(-312))).Dtor_cf2())
	}
}
func (_this *C5) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(138), _dafny.IntOfUint32((_dafny.SeqOf((_this).F10(), (_this).F10())).Cardinality())))
	}
}
func (_this *C5) Fm18(p0 _dafny.Set, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus((_dafny.IntOfInt64(-298)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(133), _dafny.UnicodeSeqOfUtf8Bytes("n"))).Cardinality()))).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-447), (_this).F10())).Cardinality(), _dafny.IntOfInt64(686)))) < 0
	}
}
func (_this *C5) Fm19(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((func() _dafny.Map {
			if (_this).F10() {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC2_(_dafny.SeqOf(_dafny.IntOfInt64(725)))), Companion_D1_.Create_DC5_(Companion_D1_.Create_DC4_((_this).F10(), (_this).F10(), _dafny.IntOfInt64(103)))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("pppcgeiju"))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(322))).Uint32(), func(coer81 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
					return func(arg81 _dafny.Int) interface{} {
						return coer81(arg81)
					}
				}(func(_1217_i0 _dafny.Int) D1 {
					return Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("wjqjefffv")))
				})), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("ojakub")))))
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("vlvw")))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("i"))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("oa"))), Companion_D1_.Create_DC5_(Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("lualcqxw")))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), Companion_D1_.Create_DC3_(_dafny.UnicodeSeqOfUtf8Bytes("jlo")))))
		})()).Cardinality()
	}
}
func (_this *C5) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _1218_v0 _dafny.Sequence
		_ = _1218_v0
		_1218_v0 = _dafny.UnicodeSeqOfUtf8Bytes("avx")
		var _1219_v1 _dafny.Int
		_ = _1219_v1
		_1219_v1 = _dafny.IntOfInt64(649)
		var _hi2 _dafny.Int = Companion_Default___.SafeModuloInt(_1219_v1, _1219_v1)
		_ = _hi2
		for _1220_i0 := _dafny.IntOfUint32((_1218_v0).Cardinality()); _1220_i0.Cmp(_hi2) < 0; _1220_i0 = _1220_i0.Plus(_dafny.One) {
			var _1221_v2 _dafny.Array
			_ = _1221_v2
			var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw181
			_1221_v2 = _nw181
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1221_v2), 0))
			_ = _index185
			(_1221_v2).ArraySet1(_1220_i0, (_index185).Int())
			var _1222_v3 _dafny.Sequence
			_ = _1222_v3
			_1222_v3 = _dafny.SeqOf(_1218_v0, Companion_Default___.Fm20(_1220_i0, (_this).F10(), (_this).F10(), (_this).Fm19(_1219_v1, _1220_i0, globalState), globalState))
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1221_v2), 0))
			_ = _index186
			(_1221_v2).ArraySet1(_dafny.IntOfUint32(((_1222_v3).Select((Companion_Default___.SafeIndex(_1220_i0, _dafny.IntOfUint32((_1222_v3).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_index186).Int())
			_1219_v1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.IntOfInt64(524)).Minus((_1221_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1221_v2), 0))).Int()).(_dafny.Int)), _1219_v1))
			var _1223_v4 _dafny.Array
			_ = _1223_v4
			var _1224_v5 _dafny.Sequence
			_ = _1224_v5
			var _1225_v6 _dafny.Set
			_ = _1225_v6
			var _out56 _dafny.Array
			_ = _out56
			var _out57 _dafny.Sequence
			_ = _out57
			var _out58 _dafny.Set
			_ = _out58
			_out56, _out57, _out58 = (_this).M4((_this).F10(), (_this).F10(), globalState)
			_1223_v4 = _out56
			_1224_v5 = _out57
			_1225_v6 = _out58
			if (_this).F10() {
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1223_v4), 0))
				_ = _index187
				(_1223_v4).ArraySet1(true, (_index187).Int())
				var _1226_v7 _dafny.Set
				_ = _1226_v7
				_1226_v7 = _dafny.SetOf((_this).F10())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1223_v4), 0))
				_ = _index188
				(_1223_v4).ArraySet1((_1226_v7).IsSubsetOf(_1226_v7), (_index188).Int())
				var _1227_v8 _dafny.Map
				_ = _1227_v8
				_1227_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1220_i0, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_1219_v1)).Cardinality()), _1219_v1, _1220_i0, _1219_v1))
				var _1228_v9 _dafny.Sequence
				_ = _1228_v9
				_1228_v9 = _dafny.SeqOf(_1220_i0)
				_1227_v8 = (_1227_v8).Update(_1220_i0, _1228_v9)
				_1218_v0 = _1218_v0
				r0 = _1218_v0
				r1 = true
			} else {
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1223_v4), 0))
				_ = _index189
				(_1223_v4).ArraySet1((_this).F10(), (_index189).Int())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1223_v4), 0))
				_ = _index190
				(_1223_v4).ArraySet1((_this).F10(), (_index190).Int())
				var _1229_v10 _dafny.Map
				_ = _1229_v10
				_1229_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm21(globalState), (_1220_i0).Plus(_dafny.IntOfUint32((_1218_v0).Cardinality())))
				var _1230_v11 _dafny.Sequence
				_ = _1230_v11
				_1230_v11 = _dafny.SeqOf(_dafny.IntOfInt64(284), _dafny.IntOfUint32((_1218_v0).Cardinality()))
				var _1231_v12 D1
				_ = _1231_v12
				_1231_v12 = Companion_D1_.Create_DC2_(_1230_v11)
				_1229_v10 = (_1229_v10).Update(_1231_v12, _1219_v1)
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1221_v2), 0))
				_ = _index191
				(_1221_v2).ArraySet1((_1230_v11).Select((Companion_Default___.SafeIndex((_1221_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_1221_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1230_v11).Cardinality()))).Uint32()).(_dafny.Int), (_index191).Int())
				r0 = _dafny.Companion_Sequence_.Concatenate(_1218_v0, _1218_v0)
				(globalState).F8 = (_1223_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_1223_v4), 0))).Int()).(bool)
			}
		}
		_1219_v1 = _1219_v1
		var _1232_v13 _dafny.Array
		_ = _1232_v13
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_28
		var _nw182 _dafny.Array
		_ = _nw182
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw182 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) bool = func(_1233_i1 _dafny.Int) bool {
				return !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_this).F10(), (_this).F10(), (_this).F10()), (_this).F10())
			}
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw182 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw182).ArraySet1(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw182).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_1232_v13 = _nw182
		var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))
		_ = _index192
		(_1232_v13).ArraySet1((func() bool {
			if ((_this).F9()).Contains((_this).F10()) {
				return ((_this).F9()).Get((_this).F10()).(bool)
			}
			return true
		})(), (_index192).Int())
		var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))
		_ = _index193
		(_1232_v13).ArraySet1(true, (_index193).Int())
		var _1234_v14 _dafny.Set
		_ = _1234_v14
		_1234_v14 = _dafny.SetOf(_1218_v0, _1218_v0, _1218_v0)
		var _1235_v15 _dafny.Map
		_ = _1235_v15
		_1235_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1234_v14).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1218_v0).Cardinality())))
		var _1236_v16 T1
		_ = _1236_v16
		var _nw183 *C2 = New_C2_()
		_ = _nw183
		_nw183.Ctor__((_this).F10(), _1232_v13, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool)), (_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool))
		_1236_v16 = _nw183
		var _1237_v17 _dafny.Map
		_ = _1237_v17
		_1237_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v15, _1236_v16)
		var _1238_v19 _dafny.Map
		_ = _1238_v19
		_1238_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1219_v1, _1236_v16)
		_1237_v17 = (_1237_v17).Update(func() _dafny.Map {
			var _coll47 = _dafny.NewMapBuilder()
			_ = _coll47
			for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(218), _dafny.IntOfInt64(-929))); ; {
				_compr_47, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _1239_v18 _dafny.Int
				_1239_v18 = interface{}(_compr_47).(_dafny.Int)
				if ((_dafny.IntOfInt64(218)).Cmp(_1239_v18) <= 0) && ((_1239_v18).Cmp(_dafny.IntOfInt64(-929)) < 0) {
					_coll47.Add((_1239_v18).Minus(_dafny.IntOfInt64(413)), _1219_v1)
				}
			}
			return _coll47.ToMap()
		}(), (func() T1 {
			if (_1238_v19).Contains(_1219_v1) {
				return (_1238_v19).Get(_1219_v1).(T1)
			}
			return _1236_v16
		})())
		var _1240_v20 *C1
		_ = _1240_v20
		var _nw184 *C1 = New_C1_()
		_ = _nw184
		_nw184.Ctor__((_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool), _1218_v0)
		_1240_v20 = _nw184
		var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))
		_ = _index194
		var _rhs191 bool = !((_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool))
		_ = _rhs191
		var _rhs192 _dafny.Int = (func() _dafny.Int {
			if (_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool) {
				return _1219_v1
			}
			return _1219_v1
		})()
		_ = _rhs192
		var _rhs193 _dafny.Int = (_dafny.Zero).Minus(_1219_v1)
		_ = _rhs193
		var _rhs194 *C1 = _1240_v20
		_ = _rhs194
		var _rhs195 bool = (_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool)
		_ = _rhs195
		var _lhs137 _dafny.Array = _1232_v13
		_ = _lhs137
		var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))
		_ = _lhs138
		var _lhs139 *GlobalState = globalState
		_ = _lhs139
		(_lhs137).ArraySet1(_rhs191, (_lhs138).Int())
		_1219_v1 = _rhs192
		_1219_v1 = _rhs193
		_1240_v20 = _rhs194
		_lhs139.F8 = _rhs195
		var _1241_i2 _dafny.Int
		_ = _1241_i2
		_1241_i2 = _dafny.Zero
		{
			for (_1240_v20).F12() {
				{
					if (_1241_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1241_i2 = (_1241_i2).Plus(_dafny.One)
					var _1242_v21 _dafny.MultiSet
					_ = _1242_v21
					_1242_v21 = _dafny.MultiSetOf(_1218_v0, _1218_v0)
					var _1243_v22 D11
					_ = _1243_v22
					_1243_v22 = Companion_D11_.Create_DC27_(_1242_v21)
					var _1244_v23 _dafny.Map
					_ = _1244_v23
					_1244_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1243_v22)
					_1244_v23 = _1244_v23
					var _1245_v24 _dafny.Map
					_ = _1245_v24
					_1245_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1240_v20).F12(), (_1240_v20).F12())
					_1245_v24 = (func() _dafny.Map {
						if (_1240_v20).F12() {
							return (_1236_v16).F9()
						}
						return (_1236_v16).F9()
					})()
					if (_1240_v20).F12() {
						var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))
						_ = _index195
						(_1232_v13).ArraySet1(((_dafny.Zero).Minus(_1219_v1)).Cmp((_1219_v1).Plus(_1219_v1)) < 0, (_index195).Int())
						var _1246_v25 _dafny.Sequence
						_ = _1246_v25
						_1246_v25 = _dafny.SeqOf(_1232_v13, _1232_v13, _1232_v13, _1232_v13)
						_1246_v25 = _1246_v25
						(globalState).F8 = false
						var _1247_v26 *C4
						_ = _1247_v26
						var _nw185 *C4 = New_C4_()
						_ = _nw185
						_nw185.Ctor__((_1236_v16).F9(), (_1236_v16).F10())
						_1247_v26 = _nw185
						var _1248_v27 _dafny.Sequence
						_ = _1248_v27
						_1248_v27 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(113))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg82 _dafny.Int) interface{} {
								return coer82(arg82)
							}
						}(func(_1249_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('v')
						}))).Cardinality()))
						var _1250_v28 _dafny.Set
						_ = _1250_v28
						_1250_v28 = _dafny.SetOf(_dafny.IntOfUint32((_1248_v27).Cardinality()))
						var _1251_v29 D6
						_ = _1251_v29
						_1251_v29 = Companion_D6_.Create_DC17_((_1236_v16).F10(), true, _1219_v1, _1219_v1, (_1235_v15).Cardinality())
						var _1252_v30 D4
						_ = _1252_v30
						_1252_v30 = Companion_D4_.Create_DC11_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1247_v26, _1219_v1)).Cardinality(), _1250_v28, !((_1251_v29).Dtor_cf30()))
						var _1253_v31 _dafny.Sequence
						_ = _1253_v31
						_1253_v31 = _dafny.SeqOf(_1252_v30, _1252_v30, Companion_D4_.Create_DC11_(((_1236_v16).F9()).Cardinality(), _1250_v28, (_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool)))
						var _1254_v32 _dafny.Sequence
						_ = _1254_v32
						_1254_v32 = _1253_v31
						_1254_v32 = _1254_v32
						var _1255_v33 _dafny.Map
						_ = _1255_v33
						_1255_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1219_v1, _1232_v13)
						_1255_v33 = (_1255_v33).Update(_1219_v1, (func() _dafny.Array {
							if (_1255_v33).Contains(_1219_v1) {
								return (_1255_v33).Get(_1219_v1).(_dafny.Array)
							}
							return _1232_v13
						})())
					} else {
						var _1256_v34 _dafny.Array
						_ = _1256_v34
						var _nw186 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(27))
						_ = _nw186
						_1256_v34 = _nw186
						var _1257_v35 D1
						_ = _1257_v35
						_1257_v35 = Companion_D1_.Create_DC4_((_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool), (_1236_v16).F10(), _1219_v1)
						var _1258_v36 _dafny.Map
						_ = _1258_v36
						_1258_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1256_v34, (_dafny.Zero).Minus((_1257_v35).Dtor_cf8()))
						var _1259_v37 D15
						_ = _1259_v37
						_1259_v37 = Companion_D15_.Create_DC37_(_1256_v34)
						_1258_v36 = (_1258_v36).Update((_1259_v37).Dtor_cf68(), _dafny.IntOfInt64(207))
						var _1260_v38 _dafny.CodePoint
						_ = _1260_v38
						_1260_v38 = _dafny.CodePoint('l')
						var _1261_v39 _dafny.Map
						_ = _1261_v39
						var _1262_v40 bool
						_ = _1262_v40
						var _out59 _dafny.Map
						_ = _out59
						var _out60 bool
						_ = _out60
						_out59, _out60 = (_1236_v16).M1(_1260_v38, globalState)
						_1261_v39 = _out59
						_1262_v40 = _out60
						var _1263_v41 _dafny.Sequence
						_ = _1263_v41
						_1263_v41 = _dafny.SeqOf(_1262_v40)
						var _1264_v42 _dafny.MultiSet
						_ = _1264_v42
						_1264_v42 = _dafny.MultiSetOf(_1219_v1)
						var _1265_v43 _dafny.Sequence
						_ = _1265_v43
						_1265_v43 = _dafny.SeqOf(_1264_v42)
						var _1266_v44 _dafny.Map
						_ = _1266_v44
						_1266_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool), _1236_v16)
						var _1267_v45 _dafny.Array
						_ = _1267_v45
						var _nwElement0_37 _dafny.Int = _1219_v1
						_ = _nwElement0_37
						var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(23))
						_ = _nw187
						(_nw187).ArraySet1(_nwElement0_37, 0)
						(_nw187).ArraySet1(_1219_v1, 1)
						(_nw187).ArraySet1(_1219_v1, 2)
						(_nw187).ArraySet1((_dafny.Zero).Minus((_this).Fm19(_1219_v1, _1219_v1, globalState)), 3)
						(_nw187).ArraySet1(_1219_v1, 4)
						(_nw187).ArraySet1(_1219_v1, 5)
						(_nw187).ArraySet1(_1219_v1, 6)
						(_nw187).ArraySet1(_1219_v1, 7)
						(_nw187).ArraySet1(_1219_v1, 8)
						(_nw187).ArraySet1(_1219_v1, 9)
						(_nw187).ArraySet1(((_1265_v43).Select((Companion_Default___.SafeIndex(_1219_v1, _dafny.IntOfUint32((_1265_v43).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), 10)
						(_nw187).ArraySet1(_dafny.IntOfInt64(-273), 11)
						(_nw187).ArraySet1(_1219_v1, 12)
						(_nw187).ArraySet1((_1266_v44).Cardinality(), 13)
						(_nw187).ArraySet1(_1219_v1, 14)
						(_nw187).ArraySet1(_1219_v1, 15)
						(_nw187).ArraySet1(_1219_v1, 16)
						(_nw187).ArraySet1(_dafny.IntOfInt64(-283), 17)
						(_nw187).ArraySet1(_1219_v1, 18)
						(_nw187).ArraySet1((_1235_v15).Cardinality(), 19)
						(_nw187).ArraySet1(_1219_v1, 20)
						(_nw187).ArraySet1(_1219_v1, 21)
						(_nw187).ArraySet1(_dafny.IntOfInt64(485), 22)
						_1267_v45 = _nw187
						var _1268_v46 _dafny.Map
						_ = _1268_v46
						_1268_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1267_v45, (_1236_v16).F10())
						var _1269_v47 _dafny.Map
						_ = _1269_v47
						_1269_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1260_v38, _1268_v46)
						(globalState).F8 = (_1236_v16).Fm0(_dafny.Companion_Sequence_.Concatenate(_1263_v41, _1263_v41), (Companion_Default___.Fm49(!((_1236_v16).F10()), globalState)).Dtor_cf29(), ((func() _dafny.Map {
							if (_1269_v47).Contains(_1260_v38) {
								return (_1269_v47).Get(_1260_v38).(_dafny.Map)
							}
							return _1268_v46
						})()).Cardinality(), (_1236_v16).Fm0(_1263_v41, (_1236_v16).F10(), _1219_v1, (_this).F10(), globalState), globalState)
						var _1270_v48 _dafny.MultiSet
						_ = _1270_v48
						_1270_v48 = _dafny.MultiSetOf((_1236_v16).F10())
						var _1271_v49 _dafny.Map
						_ = _1271_v49
						_1271_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1270_v48, _1232_v13)
						(globalState).F8 = ((_1271_v49).Cardinality()).Cmp(_1219_v1) != 0
						(globalState).F8 = false
					}
					var _1272_v50 _dafny.Array
					_ = _1272_v50
					var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
					_ = _nw188
					_1272_v50 = _nw188
					var _1273_v51 _dafny.Sequence
					_ = _1273_v51
					_1273_v51 = _dafny.SeqOf(_1272_v50, _1272_v50)
					var _1274_v52 D12
					_ = _1274_v52
					_1274_v52 = Companion_D12_.Create_DC32_(_1272_v50, _1273_v51)
					var _1275_v53 _dafny.Map
					_ = _1275_v53
					_1275_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1219_v1), _1274_v52)
					var _1276_v54 _dafny.Map
					_ = _1276_v54
					_1276_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1275_v53, _1219_v1)
					_1235_v15 = (_1235_v15).Update(((_1276_v54).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1275_v53, _1219_v1))).Cardinality(), _dafny.IntOfUint32((_1218_v0).Cardinality()))
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		r0 = _1218_v0
		r1 = (_1236_v16).F10()
		var _1277_v55 _dafny.Set
		_ = _1277_v55
		_1277_v55 = _dafny.SetOf((_this).F10(), (_1240_v20).F12(), (Companion_D0_.Create_DC0_((_1236_v16).F10())).Dtor_cf0(), !((_this).F10()))
		var _1278_v56 _dafny.Sequence
		_ = _1278_v56
		_1278_v56 = _dafny.SeqOf(_1277_v55)
		var _1279_v57 _dafny.Sequence
		_ = _1279_v57
		_1279_v57 = _dafny.SeqOf((_this).F10())
		r2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1278_v56, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(956))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg83 _dafny.Int) interface{} {
				return coer83(arg83)
			}
		}((func(_1280_v55 _dafny.Set) func(_dafny.Int) _dafny.Set {
			return func(_1281_i4 _dafny.Int) _dafny.Set {
				return _1280_v55
			}
		})(_1277_v55)))), _dafny.Companion_Sequence_.Update(_1278_v56, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.IntOfUint32((_1278_v56).Cardinality()))).Uint32(), _dafny.SetOf((_1236_v16).Fm0(_1279_v57, (_1236_v16).F10(), _1219_v1, (_1232_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1232_v13), 0))).Int()).(bool), globalState))))
		r3 = Companion_Default___.Fm22(_1219_v1, (_this).F10(), (_this).F10(), _dafny.IntOfUint32(((_1240_v20).F13()).Cardinality()), globalState)
		return r0, r1, r2, r3
	}
}
func (_this *C5) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _1282_v0 _dafny.Int
		_ = _1282_v0
		_1282_v0 = _dafny.IntOfInt64(579)
		var _1283_i0 _dafny.Int
		_ = _1283_i0
		_1283_i0 = _dafny.Zero
		{
			for (_1282_v0).Cmp(_1282_v0) > 0 {
				{
					if (_1283_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1283_i0 = (_1283_i0).Plus(_dafny.One)
					var _1284_v1 _dafny.Sequence
					_ = _1284_v1
					_1284_v1 = _dafny.UnicodeSeqOfUtf8Bytes("nllkwrs")
					var _rhs196 bool = (_this).F10()
					_ = _rhs196
					var _rhs197 _dafny.Int = _dafny.IntOfInt64(994)
					_ = _rhs197
					var _rhs198 _dafny.Sequence = (func() _dafny.Sequence {
						if (_this).F10() {
							return _dafny.UnicodeSeqOfUtf8Bytes("b")
						}
						return _dafny.Companion_Sequence_.Concatenate(_1284_v1, _dafny.UnicodeSeqOfUtf8Bytes("aogmc"))
					})()
					_ = _rhs198
					var _lhs140 *GlobalState = globalState
					_ = _lhs140
					_lhs140.F8 = _rhs196
					_1282_v0 = _rhs197
					_1284_v1 = _rhs198
					var _1285_v2 _dafny.Array
					_ = _1285_v2
					var _len0_29 _dafny.Int = _dafny.IntOfInt64(6)
					_ = _len0_29
					var _nw189 _dafny.Array
					_ = _nw189
					if _len0_29.Cmp(_dafny.Zero) == 0 {
						_nw189 = _dafny.NewArray(_len0_29)
					} else {
						var _init29 func(_dafny.Int) _dafny.CodePoint = (func(_1286_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1287_i1 _dafny.Int) _dafny.CodePoint {
								return _1286_p0
							}
						})(p0)
						_ = _init29
						var _element0_29 = _init29(_dafny.Zero)
						_ = _element0_29
						_nw189 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
						(_nw189).ArraySet1CodePoint(_element0_29, 0)
						var _nativeLen0_29 = (_len0_29).Int()
						_ = _nativeLen0_29
						for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
							(_nw189).ArraySet1CodePoint(_init29(_dafny.IntOf(_i0_29)), _i0_29)
						}
					}
					_1285_v2 = _nw189
					var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1285_v2), 0))
					_ = _index196
					(_1285_v2).ArraySet1CodePoint(p0, (_index196).Int())
					var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1285_v2), 0))
					_ = _index197
					(_1285_v2).ArraySet1CodePoint(p0, (_index197).Int())
					var _1288_v3 _dafny.Array
					_ = _1288_v3
					var _len0_30 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_30
					var _nw190 _dafny.Array
					_ = _nw190
					if _len0_30.Cmp(_dafny.Zero) == 0 {
						_nw190 = _dafny.NewArray(_len0_30)
					} else {
						var _init30 func(_dafny.Int) bool = func(_1289_i2 _dafny.Int) bool {
							return (func() bool {
								if ((_this).F9()).Contains((_this).F10()) {
									return ((_this).F9()).Get((_this).F10()).(bool)
								}
								return (_this).F10()
							})()
						}
						_ = _init30
						var _element0_30 = _init30(_dafny.Zero)
						_ = _element0_30
						_nw190 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
						(_nw190).ArraySet1(_element0_30, 0)
						var _nativeLen0_30 = (_len0_30).Int()
						_ = _nativeLen0_30
						for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
							(_nw190).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
						}
					}
					_1288_v3 = _nw190
					var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1288_v3), 0))
					_ = _index198
					(_1288_v3).ArraySet1((_this).F10(), (_index198).Int())
					var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1288_v3), 0))
					_ = _index199
					(_1288_v3).ArraySet1((_this).F10(), (_index199).Int())
					var _1290_v4 _dafny.Map
					_ = _1290_v4
					_1290_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-414), (_this).F10())
					_1290_v4 = (_1290_v4).Merge(Companion_Default___.Fm44((_this).F10(), globalState))
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1291_v5 _dafny.MultiSet
		_ = _1291_v5
		_1291_v5 = _dafny.MultiSetOf((_this).F10())
		_1282_v0 = (_dafny.Zero).Minus((((_1291_v5).Update(!(true), Companion_Default___.Abs(_1282_v0))).Difference(_dafny.MultiSetOf((_this).F10()))).Cardinality())
		(globalState).F8 = (_this).F10()
		var _1292_v6 _dafny.Sequence
		_ = _1292_v6
		_1292_v6 = _dafny.SeqOf(_1282_v0, _1282_v0)
		if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1292_v6, _dafny.SeqOf(_1282_v0)), _dafny.Companion_Sequence_.Concatenate(_1292_v6, _1292_v6)) {
			_1282_v0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1292_v6, _1292_v6), _1292_v6)).Cardinality())
			var _1293_v8 _dafny.Sequence
			_ = _1293_v8
			_1293_v8 = _dafny.UnicodeSeqOfUtf8Bytes("edykaowmo")
			var _1294_v9 _dafny.Array
			_ = _1294_v9
			var _nwElement0_38 _dafny.Int = _1282_v0
			_ = _nwElement0_38
			var _nw191 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(15))
			_ = _nw191
			(_nw191).ArraySet1(_nwElement0_38, 0)
			(_nw191).ArraySet1(_1282_v0, 1)
			(_nw191).ArraySet1(_1282_v0, 2)
			(_nw191).ArraySet1(_1282_v0, 3)
			(_nw191).ArraySet1(_1282_v0, 4)
			(_nw191).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_v0, (_this).F10())).Cardinality(), 5)
			(_nw191).ArraySet1(_1282_v0, 6)
			(_nw191).ArraySet1(_dafny.IntOfInt64(-569), 7)
			(_nw191).ArraySet1(_1282_v0, 8)
			(_nw191).ArraySet1((func() _dafny.Set {
				var _coll48 = _dafny.NewBuilder()
				_ = _coll48
				for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(98), _dafny.IntOfInt64(915))); ; {
					_compr_48, _ok52 := _iter52()
					if !_ok52 {
						break
					}
					var _1295_v7 _dafny.Int
					_1295_v7 = interface{}(_compr_48).(_dafny.Int)
					if ((_dafny.IntOfInt64(98)).Cmp(_1295_v7) <= 0) && ((_1295_v7).Cmp(_dafny.IntOfInt64(915)) < 0) {
						_coll48.Add(Companion_Default___.SafeModuloInt(_1295_v7, _dafny.IntOfUint32((_dafny.SeqOf((_this).F10())).Cardinality())))
					}
				}
				return _coll48.ToSet()
			}()).Cardinality(), 9)
			(_nw191).ArraySet1(_1282_v0, 10)
			(_nw191).ArraySet1(_dafny.IntOfUint32((_1293_v8).Cardinality()), 11)
			(_nw191).ArraySet1(_1282_v0, 12)
			(_nw191).ArraySet1(_1282_v0, 13)
			(_nw191).ArraySet1(_1282_v0, 14)
			_1294_v9 = _nw191
			var _1296_v10 _dafny.MultiSet
			_ = _1296_v10
			_1296_v10 = _dafny.MultiSetOf(_1294_v9)
			var _1297_v11 _dafny.Map
			_ = _1297_v11
			_1297_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1296_v10).Difference(_1296_v10), Companion_Default___.Fm33((_this).F10(), (_this).F10(), (_dafny.Zero).Minus(_1282_v0), !((_this).F10()), globalState))
			_1297_v11 = (_1297_v11).Update(_1296_v10, _1282_v0)
			var _1298_v12 _dafny.Array
			_ = _1298_v12
			var _nwElement0_39 bool = (_this).F10()
			_ = _nwElement0_39
			var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(4))
			_ = _nw192
			(_nw192).ArraySet1(_nwElement0_39, 0)
			(_nw192).ArraySet1((_this).F10(), 1)
			(_nw192).ArraySet1((_this).F10(), 2)
			(_nw192).ArraySet1((_this).F10(), 3)
			_1298_v12 = _nw192
			var _1299_v13 *C2
			_ = _1299_v13
			var _nw193 *C2 = New_C2_()
			_ = _nw193
			_nw193.Ctor__((_1282_v0).Cmp(_1282_v0) >= 0, _1298_v12, (_this).F9(), (_1282_v0).Cmp(_1282_v0) > 0)
			_1299_v13 = _nw193
			var _1300_v14 _dafny.Set
			_ = _1300_v14
			_1300_v14 = _dafny.SetOf((_this).F10())
			(_1299_v13).F14 = !((_1300_v14).IsSubsetOf(_1300_v14)) || ((_1282_v0).Cmp(_dafny.IntOfInt64(640)) <= 0)
			_1294_v9 = _1294_v9
		} else {
			_1282_v0 = _1282_v0
			var _1301_v15 _dafny.Array
			_ = _1301_v15
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_31
			var _nw194 _dafny.Array
			_ = _nw194
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw194 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Int = func(_1302_i3 _dafny.Int) _dafny.Int {
					return (_1302_i3).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uneixc")).Cardinality()))
				}
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw194 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw194).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw194).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1301_v15 = _nw194
			var _1303_v16 _dafny.Map
			_ = _1303_v16
			_1303_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1301_v15)
			_1303_v16 = (_1303_v16).Update((_this).F10(), _1301_v15)
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_1301_v15), 0))
			_ = _index200
			(_1301_v15).ArraySet1(_1282_v0, (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_1301_v15), 0))
			_ = _index201
			(_1301_v15).ArraySet1(_1282_v0, (_index201).Int())
			var _1304_v17 *C0
			_ = _1304_v17
			var _nw195 *C0 = New_C0_()
			_ = _nw195
			_nw195.Ctor__()
			_1304_v17 = _nw195
			r1 = (((_1301_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_1301_v15), 0))).Int()).(_dafny.Int)).Times(_1282_v0)).Cmp(_1282_v0) == 0
		}
		var _1305_v18 _dafny.Array
		_ = _1305_v18
		var _len0_32 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_32
		var _nw196 _dafny.Array
		_ = _nw196
		if _len0_32.Cmp(_dafny.Zero) == 0 {
			_nw196 = _dafny.NewArray(_len0_32)
		} else {
			var _init32 func(_dafny.Int) _dafny.Int = func(_1306_i5 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_1306_i5, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kradmkld")).Cardinality()))
			}
			_ = _init32
			var _element0_32 = _init32(_dafny.Zero)
			_ = _element0_32
			_nw196 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
			(_nw196).ArraySet1(_element0_32, 0)
			var _nativeLen0_32 = (_len0_32).Int()
			_ = _nativeLen0_32
			for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
				(_nw196).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
			}
		}
		_1305_v18 = _nw196
		for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1305_v18), 0))); ; {
			_guard_loop_4, _ok53 := _iter53()
			if !_ok53 {
				break
			}
			var _1307_i4 _dafny.Int
			_1307_i4 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1307_i4).Sign() != -1) && ((_1307_i4).Cmp(_dafny.ArrayLen((_1305_v18), 0)) < 0)) {
				(_1305_v18).ArraySet1(Companion_Default___.SafeDivisionInt(_1307_i4, _1282_v0), (_1307_i4).Int())
			}
		}
		if (_this).F10() {
			if (_this).F10() {
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index202
				(_1305_v18).ArraySet1((_1282_v0).Times(_1282_v0), (_index202).Int())
				var _1308_v19 _dafny.Sequence
				_ = _1308_v19
				_1308_v19 = _dafny.SeqOf((_this).F10())
				var _1309_v20 _dafny.Set
				_ = _1309_v20
				_1309_v20 = _dafny.SetOf(!((_this).F10()))
				var _1310_v21 _dafny.Sequence
				_ = _1310_v21
				_1310_v21 = _dafny.UnicodeSeqOfUtf8Bytes("cgysxnmd")
				var _1311_v22 _dafny.Array
				_ = _1311_v22
				var _nwElement0_40 bool = true
				_ = _nwElement0_40
				var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(29))
				_ = _nw197
				(_nw197).ArraySet1(_nwElement0_40, 0)
				(_nw197).ArraySet1((_this).Fm18(_1309_v20, _1310_v21, globalState), 1)
				(_nw197).ArraySet1((func() bool {
					if ((_this).F9()).Contains((_this).F10()) {
						return ((_this).F9()).Get((_this).F10()).(bool)
					}
					return (_this).F10()
				})(), 2)
				(_nw197).ArraySet1(true, 3)
				(_nw197).ArraySet1(false, 4)
				(_nw197).ArraySet1(true, 5)
				(_nw197).ArraySet1(true, 6)
				(_nw197).ArraySet1((_this).F10(), 7)
				(_nw197).ArraySet1((_this).F10(), 8)
				(_nw197).ArraySet1((_this).F10(), 9)
				(_nw197).ArraySet1((_this).F10(), 10)
				(_nw197).ArraySet1((_this).F10(), 11)
				(_nw197).ArraySet1((_this).F10(), 12)
				(_nw197).ArraySet1((_this).F10(), 13)
				(_nw197).ArraySet1(false, 14)
				(_nw197).ArraySet1((_this).F10(), 15)
				(_nw197).ArraySet1((_this).F10(), 16)
				(_nw197).ArraySet1((_this).F10(), 17)
				(_nw197).ArraySet1(false, 18)
				(_nw197).ArraySet1((_this).F10(), 19)
				(_nw197).ArraySet1((_this).F10(), 20)
				(_nw197).ArraySet1((_this).F10(), 21)
				(_nw197).ArraySet1((_this).F10(), 22)
				(_nw197).ArraySet1((_this).F10(), 23)
				(_nw197).ArraySet1((_this).F10(), 24)
				(_nw197).ArraySet1((func() bool {
					if ((_this).F9()).Contains((_this).F10()) {
						return ((_this).F9()).Get((_this).F10()).(bool)
					}
					return (_this).F10()
				})(), 25)
				(_nw197).ArraySet1((_this).F10(), 26)
				(_nw197).ArraySet1(true, 27)
				(_nw197).ArraySet1(true, 28)
				_1311_v22 = _nw197
				var _1312_v23 _dafny.Map
				_ = _1312_v23
				_1312_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1292_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.IntOfUint32((_1292_v6).Cardinality()))).Uint32()).(_dafny.Int))
				var _1313_v24 T1
				_ = _1313_v24
				var _nw198 *C2 = New_C2_()
				_ = _nw198
				_nw198.Ctor__((Companion_D1_.Create_DC4_((_this).Fm0(_1308_v19, true, _1282_v0, (_this).F10(), globalState), !((_this).F10()), _dafny.IntOfUint32((_1292_v6).Cardinality()))).Dtor_cf6(), _1311_v22, (_this).F9(), (_this).Fm0(_1308_v19, (_this).F10(), (func() _dafny.Int {
					if (_1312_v23).Contains(false) {
						return (_1312_v23).Get(false).(_dafny.Int)
					}
					return (_1292_v6).Select((Companion_Default___.SafeIndex(_1282_v0, _dafny.IntOfUint32((_1292_v6).Cardinality()))).Uint32()).(_dafny.Int)
				})(), (_this).F10(), globalState))
				_1313_v24 = _nw198
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index203
				var _rhs199 _dafny.Int = (_1282_v0).Plus(_1282_v0)
				_ = _rhs199
				var _rhs200 T1 = (func() T1 {
					if ((_this).F10()) == ((_this).F10()) {
						return _1313_v24
					}
					return _1313_v24
				})()
				_ = _rhs200
				var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(7))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1314_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1315_i6 _dafny.Int) _dafny.CodePoint {
						return _1314_p0
					}
				})(p0))), _1310_v21)
				_ = _rhs201
				var _lhs141 _dafny.Array = _1305_v18
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _lhs142
				(_lhs141).ArraySet1(_rhs199, (_lhs142).Int())
				_1313_v24 = _rhs200
				_1310_v21 = _rhs201
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index204
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index205
				var _rhs202 _dafny.Int = _1282_v0
				_ = _rhs202
				var _rhs203 bool = (_this).F10()
				_ = _rhs203
				var _rhs204 _dafny.Int = Companion_Default___.SafeDivisionInt(_1282_v0, (_dafny.Zero).Minus((_this).Fm19(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(347))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}((func(_1316_v0 _dafny.Int, _1317_v24 T1, _1318_v18 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_1319_i7 _dafny.Int) _dafny.Int {
						return (func() _dafny.Map {
							var _coll49 = _dafny.NewMapBuilder()
							_ = _coll49
							for _iter54 := _dafny.Iterate((func() _dafny.Map {
								var _coll50 = _dafny.NewMapBuilder()
								_ = _coll50
								for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(586), _dafny.IntOfInt64(962))); ; {
									_compr_50, _ok55 := _iter55()
									if !_ok55 {
										break
									}
									var _1320_v26 _dafny.Int
									_1320_v26 = interface{}(_compr_50).(_dafny.Int)
									if ((_dafny.IntOfInt64(586)).Cmp(_1320_v26) <= 0) && ((_1320_v26).Cmp(_dafny.IntOfInt64(962)) < 0) {
										_coll50.Add((_1320_v26).Minus((_1318_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1318_v18), 0))).Int()).(_dafny.Int)), false)
									}
								}
								return _coll50.ToMap()
							}()).Keys().Elements()); ; {
								_compr_49, _ok54 := _iter54()
								if !_ok54 {
									break
								}
								var _1321_v25 _dafny.Int
								_1321_v25 = interface{}(_compr_49).(_dafny.Int)
								if (func() _dafny.Map {
									var _coll51 = _dafny.NewMapBuilder()
									_ = _coll51
									for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(586), _dafny.IntOfInt64(962))); ; {
										_compr_51, _ok56 := _iter56()
										if !_ok56 {
											break
										}
										var _1320_v26 _dafny.Int
										_1320_v26 = interface{}(_compr_51).(_dafny.Int)
										if ((_dafny.IntOfInt64(586)).Cmp(_1320_v26) <= 0) && ((_1320_v26).Cmp(_dafny.IntOfInt64(962)) < 0) {
											_coll51.Add((_1320_v26).Minus((_1318_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1318_v18), 0))).Int()).(_dafny.Int)), false)
										}
									}
									return _coll51.ToMap()
								}()).Contains(_1321_v25) {
									_coll49.Add(Companion_Default___.SafeModuloInt(_1321_v25, _1316_v0), (_1317_v24).F10())
								}
							}
							return _coll49.ToMap()
						}()).Cardinality()
					}
				})(_1282_v0, _1313_v24, _1305_v18)))).Cardinality()), (_1305_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))).Int()).(_dafny.Int), globalState)))
				_ = _rhs204
				var _rhs205 _dafny.Sequence = _1310_v21
				_ = _rhs205
				var _rhs206 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("cxs")
				_ = _rhs206
				var _lhs143 _dafny.Array = _1305_v18
				_ = _lhs143
				var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _lhs144
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				var _lhs146 _dafny.Array = _1305_v18
				_ = _lhs146
				var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_1305_v18), 0))
				_ = _lhs147
				(_lhs143).ArraySet1(_rhs202, (_lhs144).Int())
				_lhs145.F8 = _rhs203
				(_lhs146).ArraySet1(_rhs204, (_lhs147).Int())
				_1310_v21 = _rhs205
				_1310_v21 = _rhs206
				var _1322_v27 *C3
				_ = _1322_v27
				var _nw199 *C3 = New_C3_()
				_ = _nw199
				_nw199.Ctor__()
				_1322_v27 = _nw199
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1311_v22), 0))
				_ = _index206
				(_1311_v22).ArraySet1(!((_1313_v24).F10()) || ((_1313_v24).F10()), (_index206).Int())
				var _1323_v28 D14
				_ = _1323_v28
				_1323_v28 = Companion_D14_.Create_DC35_(_1282_v0, (_1313_v24).F10(), p0)
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1311_v22), 0))
				_ = _index207
				(_1311_v22).ArraySet1((_1323_v28).Dtor_cf63(), (_index207).Int())
				var _1324_v29 *C4
				_ = _1324_v29
				var _nw200 *C4 = New_C4_()
				_ = _nw200
				_nw200.Ctor__((_this).F9(), (_1282_v0).Cmp(_1282_v0) > 0)
				_1324_v29 = _nw200
			} else {
				var _1325_v30 _dafny.Map
				_ = _1325_v30
				_1325_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), true)
				var _1326_v31 _dafny.MultiSet
				_ = _1326_v31
				_1326_v31 = _dafny.MultiSetOf(_1282_v0)
				var _1327_v32 _dafny.Set
				_ = _1327_v32
				_1327_v32 = _dafny.SetOf(_1282_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1291_v5).Cardinality(), _1282_v0)).Cardinality())
				_1325_v30 = (_1325_v30).Update((_this).F10(), (_dafny.MultiSetOf((_1327_v32).Cardinality(), _1282_v0)).IsProperSubsetOf(_1326_v31))
				var _1328_v33 _dafny.Sequence
				_ = _1328_v33
				_1328_v33 = _dafny.SeqOf((_this).F10(), (_this).F10())
				var _1329_v34 _dafny.Map
				_ = _1329_v34
				_1329_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1328_v33, _dafny.IntOfInt64(953))
				var _1330_v35 _dafny.Map
				_ = _1330_v35
				_1330_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1329_v34).Contains(_1328_v33) {
						return (_1329_v34).Get(_1328_v33).(_dafny.Int)
					}
					return _1282_v0
				})(), (_this).F10())
				var _1331_v36 _dafny.Set
				_ = _1331_v36
				_1331_v36 = _dafny.SetOf((Companion_D14_.Create_DC35_((_1330_v35).Cardinality(), (_this).F10(), p0)).Dtor_cf63())
				(globalState).F8 = !(!((_this).Fm18(_1331_v36, _dafny.UnicodeSeqOfUtf8Bytes("m"), globalState)))
				var _1332_v37 _dafny.Sequence
				_ = _1332_v37
				_1332_v37 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				_1332_v37 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-22))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}(func(_1333_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})), _1332_v37)
				var _1334_v38 _dafny.Array
				_ = _1334_v38
				var _nw201 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw201
				_1334_v38 = _nw201
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_1334_v38), 0))
				_ = _index208
				(_1334_v38).ArraySet1((_this).Fm18(_1331_v36, _1332_v37, globalState), (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_1334_v38), 0))
				_ = _index209
				(_1334_v38).ArraySet1((_this).F10(), (_index209).Int())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_1334_v38), 0))
				_ = _index210
				(_1334_v38).ArraySet1((Companion_Default___.Fm48((_this).F10(), globalState)).Contains((_this).Fm18(_1331_v36, _1332_v37, globalState)), (_index210).Int())
			}
			(globalState).F8 = (_this).F10()
			var _1335_v39 _dafny.Array
			_ = _1335_v39
			var _nwElement0_41 _dafny.Array = (func() _dafny.Array {
				if (_this).F10() {
					return _1305_v18
				}
				return _1305_v18
			})()
			_ = _nwElement0_41
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(5))
			_ = _nw202
			(_nw202).ArraySet1(_nwElement0_41, 0)
			(_nw202).ArraySet1(_1305_v18, 1)
			(_nw202).ArraySet1(_1305_v18, 2)
			(_nw202).ArraySet1(_1305_v18, 3)
			(_nw202).ArraySet1(_1305_v18, 4)
			_1335_v39 = _nw202
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1335_v39), 0))
			_ = _index211
			(_1335_v39).ArraySet1(_1305_v18, (_index211).Int())
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1335_v39), 0))
			_ = _index212
			(_1335_v39).ArraySet1(_1305_v18, (_index212).Int())
			var _1336_v40 _dafny.Array
			_ = _1336_v40
			var _nw203 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw203
			_1336_v40 = _nw203
			var _1337_v41 _dafny.Set
			_ = _1337_v41
			_1337_v41 = _dafny.SetOf((_this).F10())
			var _1338_v42 _dafny.Sequence
			_ = _1338_v42
			_1338_v42 = _dafny.UnicodeSeqOfUtf8Bytes("ksbootweo")
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))
			_ = _index213
			(_1336_v40).ArraySet1(((_this).Fm18(_1337_v41, _1338_v42, globalState)) == ((_this).F10()), (_index213).Int())
			var _1339_v43 _dafny.Sequence
			_ = _1339_v43
			_1339_v43 = _dafny.SeqOf(false, (_this).F10())
			var _1340_v44 _dafny.MultiSet
			_ = _1340_v44
			_1340_v44 = _dafny.MultiSetOf(_1282_v0, _dafny.IntOfInt64(55), _1282_v0)
			var _1341_v45 _dafny.Map
			_ = _1341_v45
			_1341_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_1282_v0, _1282_v0)).Cardinality(), p0)
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))
			_ = _index214
			(_1336_v40).ArraySet1((_dafny.IntOfUint32((_1338_v42).Cardinality())).Cmp((Companion_D6_.Create_DC17_((_this).Fm0(_1339_v43, (_this).F10(), (_1340_v44).Cardinality(), (_this).F10(), globalState), true, (_1341_v45).Cardinality(), _1282_v0, _1282_v0)).Dtor_cf31()) <= 0, (_index214).Int())
			if (Companion_D0_.Create_DC1_(_1282_v0, (func() bool {
				if ((_this).F9()).Contains((_1336_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))).Int()).(bool)) {
					return ((_this).F9()).Get((_1336_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))).Int()).(bool)).(bool)
				}
				return (_1336_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))).Int()).(bool)
			})(), _1282_v0)).Dtor_cf2() {
				(globalState).F8 = !((_1336_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))).Int()).(bool))
				var _1342_v46 _dafny.Array
				_ = _1342_v46
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_33
				var _nw204 _dafny.Array
				_ = _nw204
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw204 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.MultiSet = (func(_1343_v40 _dafny.Array, _1344_v5 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_1345_i9 _dafny.Int) _dafny.MultiSet {
							return (func() _dafny.MultiSet {
								if (_1343_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1343_v40), 0))).Int()).(bool) {
									return _1344_v5
								}
								return _dafny.MultiSetOf((_1343_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1343_v40), 0))).Int()).(bool), (_1343_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1343_v40), 0))).Int()).(bool))
							})()
						}
					})(_1336_v40, _1291_v5)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw204 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw204).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw204).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1342_v46 = _nw204
				var _1346_v47 _dafny.Array
				_ = _1346_v47
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_34
				var _nw205 _dafny.Array
				_ = _nw205
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw205 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1347_v43 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1348_i10 _dafny.Int) _dafny.Sequence {
							return _1347_v43
						}
					})(_1339_v43)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw205 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw205).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw205).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1346_v47 = _nw205
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1346_v47), 0))
				_ = _index215
				(_1346_v47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1339_v43, _dafny.Companion_Sequence_.Update(_1339_v43, (Companion_Default___.SafeIndex(_1282_v0, _dafny.IntOfUint32((_1339_v43).Cardinality()))).Uint32(), (_1336_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))).Int()).(bool))), (_index215).Int())
				var _1349_v48 _dafny.Array
				_ = _1349_v48
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw206
				_1349_v48 = _nw206
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1349_v48), 0))
				_ = _index216
				(_1349_v48).ArraySet1(_1336_v40, (_index216).Int())
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1346_v47), 0))
				_ = _index217
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1349_v48), 0))
				_ = _index218
				var _rhs207 _dafny.Array = _1342_v46
				_ = _rhs207
				var _rhs208 _dafny.Sequence = _1339_v43
				_ = _rhs208
				var _rhs209 _dafny.Array = _1336_v40
				_ = _rhs209
				var _rhs210 bool = ((_this).F10()) || ((_this).F10())
				_ = _rhs210
				var _rhs211 _dafny.Int = _1282_v0
				_ = _rhs211
				var _lhs148 _dafny.Array = _1346_v47
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1346_v47), 0))
				_ = _lhs149
				var _lhs150 _dafny.Array = _1349_v48
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_1349_v48), 0))
				_ = _lhs151
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				_1342_v46 = _rhs207
				(_lhs148).ArraySet1(_rhs208, (_lhs149).Int())
				(_lhs150).ArraySet1(_rhs209, (_lhs151).Int())
				_lhs152.F8 = _rhs210
				_1282_v0 = _rhs211
				_1282_v0 = (_dafny.IntOfInt64(345)).Times(_1282_v0)
				var _1350_v49 _dafny.Set
				_ = _1350_v49
				_1350_v49 = _dafny.SetOf(_1282_v0, _1282_v0)
				var _1351_v50 D4
				_ = _1351_v50
				_1351_v50 = Companion_D4_.Create_DC11_(_1282_v0, _1350_v49, (_1336_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1336_v40), 0))).Int()).(bool))
				var _1352_v51 _dafny.Sequence
				_ = _1352_v51
				_1352_v51 = _dafny.SeqOf(_1351_v50, _1351_v50, _1351_v50)
				var _1353_v52 _dafny.Sequence
				_ = _1353_v52
				_1353_v52 = _dafny.Companion_Sequence_.Concatenate(_1352_v51, _1352_v51)
				_1353_v52 = _dafny.SeqOf(_1351_v50, _1351_v50, _1351_v50, _1351_v50, _1351_v50)
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1335_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1335_v39), 0))).Int()))
				_ = _arr2
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_dafny.ArrayCastTo((_1335_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1335_v39), 0))).Int()))), 0))
				_ = _index219
				(_arr2).ArraySet1(Companion_Default___.SafeModuloInt(_1282_v0, _dafny.IntOfUint32((_1338_v42).Cardinality())), (_index219).Int())
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1335_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1335_v39), 0))).Int()))
				_ = _arr3
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_dafny.ArrayCastTo((_1335_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1335_v39), 0))).Int()))), 0))
				_ = _index220
				(_arr3).ArraySet1(((((_this).F9()).Cardinality()).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_1338_v42).Cardinality())))).Plus(_1282_v0), (_index220).Int())
			} else {
				var _1354_v53 _dafny.Sequence
				_ = _1354_v53
				_1354_v53 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1292_v6, (Companion_Default___.SafeIndex(_1282_v0, _dafny.IntOfUint32((_1292_v6).Cardinality()))).Uint32(), _1282_v0))
				var _1355_v54 *C2
				_ = _1355_v54
				var _nw207 *C2 = New_C2_()
				_ = _nw207
				_nw207.Ctor__(_dafny.Companion_Sequence_.Contains(_1354_v53, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1356_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1357_i11 _dafny.Int) _dafny.Int {
						return _1356_v0
					}
				})(_1282_v0)))), _1336_v40, ((_this).F9()).Merge((_this).F9()), (_this).F10())
				_1355_v54 = _nw207
				var _rhs212 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1282_v0))
				_ = _rhs212
				var _rhs213 _dafny.Int = _1282_v0
				_ = _rhs213
				var _rhs214 _dafny.Int = _1282_v0
				_ = _rhs214
				var _rhs215 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("tm"), _1338_v42)
				_ = _rhs215
				var _rhs216 bool = false
				_ = _rhs216
				var _lhs153 *GlobalState = globalState
				_ = _lhs153
				var _lhs154 *C2 = _1355_v54
				_ = _lhs154
				_1282_v0 = _rhs212
				_1282_v0 = _rhs213
				_1282_v0 = _rhs214
				_lhs153.F8 = _rhs215
				_lhs154.F14 = _rhs216
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index221
				(_1305_v18).ArraySet1(_1282_v0, (_index221).Int())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index222
				(_1305_v18).ArraySet1(_1282_v0, (_index222).Int())
				var _1358_v55 _dafny.Sequence
				_ = _1358_v55
				_1358_v55 = _dafny.SeqOf(Companion_Default___.Fm50((_dafny.Zero).Minus((_dafny.MultiSetOf((_1305_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1305_v18), 0))).Int()).(_dafny.Int))).Cardinality()), (_1305_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_1305_v18), 0))).Int()).(_dafny.Int), globalState))
				_1358_v55 = _1358_v55
				(globalState).F8 = _1355_v54.F14
			}
		} else {
			var _1359_v56 _dafny.Sequence
			_ = _1359_v56
			_1359_v56 = _dafny.UnicodeSeqOfUtf8Bytes("ty")
			var _1360_v57 _dafny.Sequence
			_ = _1360_v57
			_1360_v57 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yof"), _dafny.UnicodeSeqOfUtf8Bytes("uhu")), _dafny.Companion_Sequence_.Concatenate(_1359_v56, _1359_v56), _1359_v56, _dafny.Companion_Sequence_.Concatenate(_1359_v56, _1359_v56), _1359_v56)
			_1360_v57 = _dafny.SeqOf(_1359_v56, _dafny.UnicodeSeqOfUtf8Bytes("y"))
			var _1361_v58 D14
			_ = _1361_v58
			_1361_v58 = Companion_D14_.Create_DC36_(false, _1359_v56, _1359_v56)
			var _1362_v59 _dafny.MultiSet
			_ = _1362_v59
			_1362_v59 = _dafny.MultiSetOf(_1361_v58)
			_1282_v0 = ((_1362_v59).Update(_1361_v58, Companion_Default___.Abs(_1282_v0))).Cardinality()
			var _1363_v60 D1
			_ = _1363_v60
			_1363_v60 = Companion_D1_.Create_DC3_(_1359_v56)
			var _1364_v61 D1
			_ = _1364_v61
			_1364_v61 = Companion_D1_.Create_DC5_(_1363_v60)
			var _source18 D1 = _1364_v61
			_ = _source18
			if _source18.Is_DC3() {
				var _1365___mcc_h0 _dafny.Sequence = _source18.Get_().(D1_DC3).Cf5
				_ = _1365___mcc_h0
				var _1366_cf5 _dafny.Sequence = _1365___mcc_h0
				_ = _1366_cf5
				var _1367_v62 _dafny.Sequence
				_ = _1367_v62
				_1367_v62 = _dafny.SeqOf(_1305_v18)
				var _1368_v63 _dafny.Set
				_ = _1368_v63
				_1368_v63 = _dafny.SetOf(_1282_v0, _1282_v0)
				var _1369_v64 D4
				_ = _1369_v64
				_1369_v64 = Companion_D4_.Create_DC11_(_1282_v0, _1368_v63, (_this).F10())
				var _1370_v65 _dafny.Map
				_ = _1370_v65
				_1370_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1367_v62, ((_1369_v64).Dtor_cf16()).Minus(_1282_v0))
				_1370_v65 = (_1370_v65).Update(_1367_v62, (_dafny.Zero).Minus(_1282_v0))
				var _1371_v66 _dafny.Map
				_ = _1371_v66
				_1371_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf((_this).F10(), (_this).F10()))
				var _1372_v67 _dafny.Sequence
				_ = _1372_v67
				_1372_v67 = _dafny.SeqOf(false)
				_1371_v66 = (_1371_v66).Update((_this).F10(), _1372_v67)
				_1282_v0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1372_v67).Cardinality()), (_this).Fm1((_this).F9(), globalState))
				var _1373_v68 _dafny.Array
				_ = _1373_v68
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_35
				var _nw208 _dafny.Array
				_ = _nw208
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw208 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.CodePoint = func(_1374_i12 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('r')
					}
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw208 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw208).ArraySet1CodePoint(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw208).ArraySet1CodePoint(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1373_v68 = _nw208
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_1373_v68), 0))
				_ = _index223
				(_1373_v68).ArraySet1CodePoint(p0, (_index223).Int())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_1373_v68), 0))
				_ = _index224
				(_1373_v68).ArraySet1CodePoint(p0, (_index224).Int())
			} else if _source18.Is_DC4() {
				var _1375___mcc_h1 bool = _source18.Get_().(D1_DC4).Cf6
				_ = _1375___mcc_h1
				var _1376___mcc_h2 bool = _source18.Get_().(D1_DC4).Cf7
				_ = _1376___mcc_h2
				var _1377___mcc_h3 _dafny.Int = _source18.Get_().(D1_DC4).Cf8
				_ = _1377___mcc_h3
				var _1378_cf8 _dafny.Int = _1377___mcc_h3
				_ = _1378_cf8
				var _1379_cf7 bool = _1376___mcc_h2
				_ = _1379_cf7
				var _1380_cf6 bool = _1375___mcc_h1
				_ = _1380_cf6
				var _1381_v69 _dafny.Set
				_ = _1381_v69
				_1381_v69 = _dafny.SetOf(_dafny.IntOfInt64(976))
				var _1382_v70 D4
				_ = _1382_v70
				_1382_v70 = Companion_D4_.Create_DC11_(_1282_v0, _1381_v69, (_this).F10())
				var _1383_v71 _dafny.Sequence
				_ = _1383_v71
				_1383_v71 = _dafny.SeqOf(_1382_v70)
				var _1384_v72 _dafny.Map
				_ = _1384_v72
				_1384_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1383_v71, _1282_v0)
				var _1385_v73 _dafny.Sequence
				_ = _1385_v73
				_1385_v73 = _dafny.SeqOf(_1384_v72, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1383_v71, (_dafny.Zero).Minus(_1282_v0)), _1384_v72, _1384_v72)
				var _1386_v74 D14
				_ = _1386_v74
				_1386_v74 = Companion_D14_.Create_DC34_((_1385_v73).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.IntOfUint32((_1385_v73).Cardinality()))).Uint32()).(_dafny.Map))
				_1386_v74 = Companion_Default___.Fm51(globalState)
				var _1387_v76 _dafny.Array
				_ = _1387_v76
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_36
				var _nw209 _dafny.Array
				_ = _nw209
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw209 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Set = (func(_1388_cf7 bool, _1389_v0 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_1390_i13 _dafny.Int) _dafny.Set {
							return func() _dafny.Set {
								var _coll52 = _dafny.NewBuilder()
								_ = _coll52
								for _iter57 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1388_cf7), _1389_v0)).Keys().Elements()); ; {
									_compr_52, _ok57 := _iter57()
									if !_ok57 {
										break
									}
									var _1391_v75 _dafny.Sequence
									_1391_v75 = interface{}(_compr_52).(_dafny.Sequence)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1388_cf7), _1389_v0)).Contains(_1391_v75) {
										_coll52.Add(_1391_v75)
									}
								}
								return _coll52.ToSet()
							}()
						}
					})(_1379_cf7, _1282_v0)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw209 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw209).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw209).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1387_v76 = _nw209
				var _1392_v77 _dafny.Sequence
				_ = _1392_v77
				_1392_v77 = _dafny.SeqOf(_1379_cf7)
				var _1393_v78 _dafny.Set
				_ = _1393_v78
				_1393_v78 = _dafny.SetOf(_1392_v77)
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1387_v76), 0))
				_ = _index225
				(_1387_v76).ArraySet1(_1393_v78, (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1387_v76), 0))
				_ = _index226
				(_1387_v76).ArraySet1(_1393_v78, (_index226).Int())
				var _1394_v79 _dafny.Sequence
				_ = _1394_v79
				_1394_v79 = _dafny.SeqOf(_1305_v18, _1305_v18, _1305_v18, _1305_v18, _1305_v18)
				_1305_v18 = (_1394_v79).Select((Companion_Default___.SafeIndex((_this).Fm1(((_this).F9()).Update(false, !((func() bool {
					if ((_this).F9()).Contains(false) {
						return ((_this).F9()).Get(false).(bool)
					}
					return true
				})())), globalState), _dafny.IntOfUint32((_1394_v79).Cardinality()))).Uint32()).(_dafny.Array)
				(globalState).F8 = _1379_cf7
			} else if _source18.Is_DC2() {
				var _1395___mcc_h4 _dafny.Sequence = _source18.Get_().(D1_DC2).Cf4
				_ = _1395___mcc_h4
				var _1396_cf4 _dafny.Sequence = _1395___mcc_h4
				_ = _1396_cf4
				var _rhs217 bool = (_this).F10()
				_ = _rhs217
				var _rhs218 bool = (_this).F10()
				_ = _rhs218
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				_lhs155.F8 = _rhs217
				r1 = _rhs218
				var _1397_v80 *C0
				_ = _1397_v80
				var _nw210 *C0 = New_C0_()
				_ = _nw210
				_nw210.Ctor__()
				_1397_v80 = _nw210
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index227
				(_1305_v18).ArraySet1((_dafny.Zero).Minus((_this).Fm19(_1282_v0, (_dafny.Zero).Minus(_1282_v0), globalState)), (_index227).Int())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_1305_v18), 0))
				_ = _index228
				(_1305_v18).ArraySet1(_dafny.IntOfInt64(-366), (_index228).Int())
				var _1398_v81 _dafny.CodePoint
				_ = _1398_v81
				_1398_v81 = _dafny.CodePoint('o')
				_1398_v81 = _1398_v81
			} else {
				var _1399___mcc_h5 D1 = _source18.Get_().(D1_DC5).Cf9
				_ = _1399___mcc_h5
				var _1400_cf9 D1 = _1399___mcc_h5
				_ = _1400_cf9
				var _1401_v82 _dafny.Set
				_ = _1401_v82
				_1401_v82 = _dafny.SetOf(_1305_v18)
				_1401_v82 = _1401_v82
				(globalState).F8 = (_this).F10()
				_1282_v0 = _1282_v0
				var _1402_v83 _dafny.Array
				_ = _1402_v83
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw211
				_1402_v83 = _nw211
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1402_v83), 0))
				_ = _index229
				(_1402_v83).ArraySet1(_1359_v56, (_index229).Int())
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1402_v83), 0))
				_ = _index230
				(_1402_v83).ArraySet1(_1359_v56, (_index230).Int())
			}
			r1 = (_this).F10()
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_1305_v18), 0))
			_ = _index231
			(_1305_v18).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1292_v6).Cardinality()), _1282_v0), (_index231).Int())
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_1305_v18), 0))
			_ = _index232
			(_1305_v18).ArraySet1(_dafny.IntOfInt64(106), (_index232).Int())
		}
		var _1403_v84 _dafny.Sequence
		_ = _1403_v84
		_1403_v84 = _dafny.UnicodeSeqOfUtf8Bytes("dh")
		var _1404_v85 _dafny.Map
		_ = _1404_v85
		_1404_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1305_v18, _dafny.Companion_Sequence_.Concatenate(_1403_v84, _1403_v84))
		r0 = _1404_v85
		r1 = (_this).F10()
		return r0, r1
	}
}
func (_this *C5) M4(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Array, _dafny.Sequence, _dafny.Set) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var _1405_v0 _dafny.Int
		_ = _1405_v0
		_1405_v0 = _dafny.IntOfInt64(49)
		var _hi3 _dafny.Int = _1405_v0
		_ = _hi3
		for _1406_i0 := Companion_Default___.SafeDivisionInt(_1405_v0, _dafny.IntOfInt64(146)); _1406_i0.Cmp(_hi3) < 0; _1406_i0 = _1406_i0.Plus(_dafny.One) {
			var _1407_v1 _dafny.Sequence
			_ = _1407_v1
			_1407_v1 = _dafny.UnicodeSeqOfUtf8Bytes("itssvdb")
			_1407_v1 = _dafny.UnicodeSeqOfUtf8Bytes("mfa")
			var _source19 D6 = Companion_D6_.Create_DC17_(p1, (_this).F10(), _1406_i0, _1405_v0, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(576), _1406_i0))
			_ = _source19
			if _source19.Is_DC17() {
				var _1408___mcc_h0 bool = _source19.Get_().(D6_DC17).Cf29
				_ = _1408___mcc_h0
				var _1409___mcc_h1 bool = _source19.Get_().(D6_DC17).Cf30
				_ = _1409___mcc_h1
				var _1410___mcc_h2 _dafny.Int = _source19.Get_().(D6_DC17).Cf31
				_ = _1410___mcc_h2
				var _1411___mcc_h3 _dafny.Int = _source19.Get_().(D6_DC17).Cf32
				_ = _1411___mcc_h3
				var _1412___mcc_h4 _dafny.Int = _source19.Get_().(D6_DC17).Cf33
				_ = _1412___mcc_h4
				var _1413_cf33 _dafny.Int = _1412___mcc_h4
				_ = _1413_cf33
				var _1414_cf32 _dafny.Int = _1411___mcc_h3
				_ = _1414_cf32
				var _1415_cf31 _dafny.Int = _1410___mcc_h2
				_ = _1415_cf31
				var _1416_cf30 bool = _1409___mcc_h1
				_ = _1416_cf30
				var _1417_cf29 bool = _1408___mcc_h0
				_ = _1417_cf29
				_1415_cf31 = (_1406_i0).Plus(_1405_v0)
				var _1418_v2 _dafny.CodePoint
				_ = _1418_v2
				_1418_v2 = _dafny.CodePoint('t')
				_1418_v2 = _1418_v2
				var _1419_v3 _dafny.Sequence
				_ = _1419_v3
				_1419_v3 = _dafny.SeqOf(_1416_cf30)
				var _1420_v4 D8
				_ = _1420_v4
				_1420_v4 = Companion_D8_.Create_DC23_(_dafny.SeqOf((_this).F10(), p0, (_this).Fm0(_1419_v3, p1, _dafny.IntOfInt64(313), false, globalState)))
				_1420_v4 = _1420_v4
				var _1421_v5 _dafny.Set
				_ = _1421_v5
				_1421_v5 = _dafny.SetOf((_this).F10())
				var _1422_v6 _dafny.Array
				_ = _1422_v6
				var _nwElement0_42 _dafny.Int = (_1421_v5).Cardinality()
				_ = _nwElement0_42
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(4))
				_ = _nw212
				(_nw212).ArraySet1(_nwElement0_42, 0)
				(_nw212).ArraySet1(_dafny.IntOfInt64(-973), 1)
				(_nw212).ArraySet1(_1414_cf32, 2)
				(_nw212).ArraySet1(_1415_cf31, 3)
				_1422_v6 = _nw212
				var _1423_v7 _dafny.Set
				_ = _1423_v7
				_1423_v7 = _dafny.SetOf(_1422_v6, _1422_v6)
				_1423_v7 = _1423_v7
			} else if _source19.Is_DC18() {
				var _1424___mcc_h5 _dafny.CodePoint = _source19.Get_().(D6_DC18).Cf34
				_ = _1424___mcc_h5
				var _1425___mcc_h6 _dafny.Int = _source19.Get_().(D6_DC18).Cf35
				_ = _1425___mcc_h6
				var _1426___mcc_h7 _dafny.Int = _source19.Get_().(D6_DC18).Cf36
				_ = _1426___mcc_h7
				var _1427_cf36 _dafny.Int = _1426___mcc_h7
				_ = _1427_cf36
				var _1428_cf35 _dafny.Int = _1425___mcc_h6
				_ = _1428_cf35
				var _1429_cf34 _dafny.CodePoint = _1424___mcc_h5
				_ = _1429_cf34
				var _1430_v8 D14
				_ = _1430_v8
				_1430_v8 = Companion_D14_.Create_DC35_(_1427_cf36, !(true), _1429_cf34)
				_1427_cf36 = ((_1430_v8).Dtor_cf62()).Minus(_1428_cf35)
				var _1431_v9 _dafny.Array
				_ = _1431_v9
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_37
				var _nw213 _dafny.Array
				_ = _nw213
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw213 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) bool = (func(_1432_p1 bool) func(_dafny.Int) bool {
						return func(_1433_i1 _dafny.Int) bool {
							return _1432_p1
						}
					})(p1)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw213 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw213).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw213).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1431_v9 = _nw213
				var _1434_v10 *C2
				_ = _1434_v10
				var _nw214 *C2 = New_C2_()
				_ = _nw214
				_nw214.Ctor__(p1, _1431_v9, (_this).F9(), p0)
				_1434_v10 = _nw214
				var _1435_v11 _dafny.Map
				_ = _1435_v11
				var _1436_v12 bool
				_ = _1436_v12
				var _out61 _dafny.Map
				_ = _out61
				var _out62 bool
				_ = _out62
				_out61, _out62 = (_1434_v10).M1(_1429_cf34, globalState)
				_1435_v11 = _out61
				_1436_v12 = _out62
				var _1437_v13 D0
				_ = _1437_v13
				_1437_v13 = Companion_D0_.Create_DC1_(_1406_i0, !(!(p0)), _1428_cf35)
				var _rhs219 _dafny.Int = ((_1427_cf36).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1406_i0)).Cardinality()))).Times(_1406_i0)
				_ = _rhs219
				var _rhs220 bool = (_1437_v13).Dtor_cf2()
				_ = _rhs220
				var _rhs221 bool = p1
				_ = _rhs221
				var _rhs222 bool = (_this).F10()
				_ = _rhs222
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				_1427_cf36 = _rhs219
				_lhs156.F8 = _rhs220
				_lhs157.F8 = _rhs221
				_lhs158.F8 = _rhs222
			} else if _source19.Is_DC19() {
				var _1438___mcc_h8 bool = _source19.Get_().(D6_DC19).Cf37
				_ = _1438___mcc_h8
				var _1439___mcc_h9 _dafny.CodePoint = _source19.Get_().(D6_DC19).Cf38
				_ = _1439___mcc_h9
				var _1440_cf38 _dafny.CodePoint = _1439___mcc_h9
				_ = _1440_cf38
				var _1441_cf37 bool = _1438___mcc_h8
				_ = _1441_cf37
				_1405_v0 = _1406_i0
				var _1442_v14 _dafny.Sequence
				_ = _1442_v14
				_1442_v14 = _dafny.SeqOf((_this).F10(), !((_this).F10()))
				var _1443_v15 _dafny.Set
				_ = _1443_v15
				_1443_v15 = _dafny.SetOf(_1442_v14)
				var _1444_v16 _dafny.MultiSet
				_ = _1444_v16
				_1444_v16 = _dafny.MultiSetOf(p1)
				var _1445_v17 _dafny.Sequence
				_ = _1445_v17
				_1445_v17 = _dafny.SeqOf(_1406_i0)
				var _1446_v18 _dafny.Set
				_ = _1446_v18
				_1446_v18 = _dafny.SetOf((func() _dafny.Int {
					if (_1444_v16).Contains(_1441_cf37) {
						return (_1444_v16).Multiplicity(_1441_cf37)
					}
					return _dafny.IntOfUint32((_1445_v17).Cardinality())
				})())
				_1441_cf37 = (_1446_v18).IsProperSubsetOf(_dafny.SetOf((_1443_v15).Cardinality()))
				var _1447_v19 D8
				_ = _1447_v19
				_1447_v19 = Companion_D8_.Create_DC23_(_1442_v14)
				var _pat_let_tv26 = _1442_v14
				_ = _pat_let_tv26
				var _pat_let_tv27 = _1442_v14
				_ = _pat_let_tv27
				var _1448_v20 _dafny.Array
				_ = _1448_v20
				var _nwElement0_43 D8 = _1447_v19
				_ = _nwElement0_43
				var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(28))
				_ = _nw215
				(_nw215).ArraySet1(_nwElement0_43, 0)
				(_nw215).ArraySet1(Companion_Default___.Fm52(!((_this).F10()), globalState), 1)
				(_nw215).ArraySet1(_1447_v19, 2)
				(_nw215).ArraySet1(Companion_D8_.Create_DC23_(_1442_v14), 3)
				(_nw215).ArraySet1(_1447_v19, 4)
				(_nw215).ArraySet1(_1447_v19, 5)
				(_nw215).ArraySet1(_1447_v19, 6)
				(_nw215).ArraySet1(_1447_v19, 7)
				(_nw215).ArraySet1(func(_pat_let27_0 D8) D8 {
					return func(_1449_dt__update__tmp_h0 D8) D8 {
						return func(_pat_let28_0 _dafny.Sequence) D8 {
							return func(_1450_dt__update_hcf41_h0 _dafny.Sequence) D8 {
								return Companion_D8_.Create_DC23_(_1450_dt__update_hcf41_h0)
							}(_pat_let28_0)
						}(_pat_let_tv26)
					}(_pat_let27_0)
				}(_1447_v19), 8)
				(_nw215).ArraySet1(Companion_Default___.Fm52(_1441_cf37, globalState), 9)
				(_nw215).ArraySet1(_1447_v19, 10)
				(_nw215).ArraySet1(_1447_v19, 11)
				(_nw215).ArraySet1(_1447_v19, 12)
				(_nw215).ArraySet1(_1447_v19, 13)
				(_nw215).ArraySet1(_1447_v19, 14)
				(_nw215).ArraySet1(_1447_v19, 15)
				(_nw215).ArraySet1(_1447_v19, 16)
				(_nw215).ArraySet1(_1447_v19, 17)
				(_nw215).ArraySet1(_1447_v19, 18)
				(_nw215).ArraySet1(_1447_v19, 19)
				(_nw215).ArraySet1(_1447_v19, 20)
				(_nw215).ArraySet1(func(_pat_let29_0 D8) D8 {
					return func(_1451_dt__update__tmp_h1 D8) D8 {
						return func(_pat_let30_0 _dafny.Sequence) D8 {
							return func(_1452_dt__update_hcf41_h1 _dafny.Sequence) D8 {
								return Companion_D8_.Create_DC23_(_1452_dt__update_hcf41_h1)
							}(_pat_let30_0)
						}(_pat_let_tv27)
					}(_pat_let29_0)
				}(_1447_v19), 21)
				(_nw215).ArraySet1(Companion_D8_.Create_DC23_(_1442_v14), 22)
				(_nw215).ArraySet1(_1447_v19, 23)
				(_nw215).ArraySet1(Companion_D8_.Create_DC23_(_1442_v14), 24)
				(_nw215).ArraySet1(Companion_D8_.Create_DC23_(_1442_v14), 25)
				(_nw215).ArraySet1(_1447_v19, 26)
				(_nw215).ArraySet1(_1447_v19, 27)
				_1448_v20 = _nw215
				var _1453_v21 _dafny.Sequence
				_ = _1453_v21
				_1453_v21 = _dafny.SeqOf(_1448_v20)
				var _rhs223 _dafny.Sequence = _1453_v21
				_ = _rhs223
				var _rhs224 bool = p1
				_ = _rhs224
				var _rhs225 _dafny.Int = _1406_i0
				_ = _rhs225
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				_1453_v21 = _rhs223
				_lhs159.F8 = _rhs224
				_1405_v0 = _rhs225
				var _1454_v22 _dafny.Map
				_ = _1454_v22
				_1454_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1440_cf38, _1440_cf38)
				var _1455_v23 D11
				_ = _1455_v23
				_1455_v23 = Companion_D11_.Create_DC28_(_1405_v0, _1441_cf37, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1446_v18).Cardinality(), _1454_v22), (_this).F10(), _dafny.Companion_Sequence_.Update(_1407_v1, (Companion_Default___.SafeIndex(_1406_i0, _dafny.IntOfUint32((_1407_v1).Cardinality()))).Uint32(), _1440_cf38))
				(globalState).F8 = !(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1407_v1).Cardinality()), _1440_cf38)).Cardinality()).Cmp(_1405_v0) > 0) || ((_1455_v23).Dtor_cf51())
			} else if _source19.Is_DC16() {
				var _1456___mcc_h10 _dafny.Map = _source19.Get_().(D6_DC16).Cf28
				_ = _1456___mcc_h10
				var _1457_cf28 _dafny.Map = _1456___mcc_h10
				_ = _1457_cf28
				var _1458_v24 *C3
				_ = _1458_v24
				var _nw216 *C3 = New_C3_()
				_ = _nw216
				_nw216.Ctor__()
				_1458_v24 = _nw216
				var _1459_v25 _dafny.CodePoint
				_ = _1459_v25
				_1459_v25 = _dafny.CodePoint('x')
				var _1460_v26 _dafny.MultiSet
				_ = _1460_v26
				_1460_v26 = _dafny.MultiSetOf(_1406_i0)
				var _rhs226 bool = (_this).F10()
				_ = _rhs226
				var _rhs227 *C3 = (func() *C3 {
					if (_1460_v26).IsSubsetOf(_dafny.MultiSetOf(_1405_v0, _1405_v0)) {
						return _1458_v24
					}
					return _1458_v24
				})()
				_ = _rhs227
				var _rhs228 bool = p1
				_ = _rhs228
				var _rhs229 _dafny.CodePoint = _1459_v25
				_ = _rhs229
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				_lhs160.F8 = _rhs226
				_1458_v24 = _rhs227
				_lhs161.F8 = _rhs228
				_1459_v25 = _rhs229
				(globalState).F8 = !((_dafny.IntOfUint32((_1407_v1).Cardinality())).Cmp(_1405_v0) > 0) || (p1)
				var _1461_v27 D7
				_ = _1461_v27
				_1461_v27 = Companion_D7_.Create_DC21_(_dafny.SetOf(p1, false))
				var _1462_v28 _dafny.Set
				_ = _1462_v28
				_1462_v28 = _dafny.SetOf((_this).F10(), (_this).F10())
				_1461_v27 = Companion_D7_.Create_DC21_(_1462_v28)
				var _1463_v29 _dafny.Sequence
				_ = _1463_v29
				_1463_v29 = _dafny.SeqOf(_1460_v26, _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-567))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}((func(_1464_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1465_i2 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_1464_v0)
					}
				})(_1405_v0)))))
				var _1466_v30 _dafny.MultiSet
				_ = _1466_v30
				_1466_v30 = _dafny.MultiSetOf(p1, p0)
				var _1467_v31 _dafny.Set
				_ = _1467_v31
				_1467_v31 = _dafny.SetOf(_dafny.IntOfInt64(628), (_1466_v30).Cardinality())
				var _1468_v32 D4
				_ = _1468_v32
				_1468_v32 = Companion_D4_.Create_DC11_(_dafny.IntOfInt64(914), _1467_v31, (_this).F10())
				var _1469_v33 _dafny.Map
				_ = _1469_v33
				_1469_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, p1)
				var _pat_let_tv28 = _1469_v33
				_ = _pat_let_tv28
				var _pat_let_tv29 = _1469_v33
				_ = _pat_let_tv29
				var _pat_let_tv30 = _1468_v32
				_ = _pat_let_tv30
				_1405_v0 = (func(_pat_let31_0 D1) D1 {
					return func(_1470_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let32_0 bool) D1 {
							return func(_1471_dt__update_hcf7_h0 bool) D1 {
								return func(_pat_let33_0 _dafny.Int) D1 {
									return func(_1472_dt__update_hcf8_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC4_((_1470_dt__update__tmp_h2).Dtor_cf6(), _1471_dt__update_hcf7_h0, _1472_dt__update_hcf8_h0)
									}(_pat_let33_0)
								}((_pat_let_tv30).Dtor_cf16())
							}(_pat_let32_0)
						}((func() bool {
							if (_pat_let_tv28).Contains(_1406_i0) {
								return (_pat_let_tv29).Get(_1406_i0).(bool)
							}
							return (_this).F10()
						})())
					}(_pat_let31_0)
				}(Companion_D1_.Create_DC4_(p1, p1, _dafny.IntOfUint32((_1463_v29).Cardinality())))).Dtor_cf8()
			} else {
				var _1473___mcc_h11 D6 = _source19.Get_().(D6_DC20).Cf39
				_ = _1473___mcc_h11
				var _1474_cf39 D6 = _1473___mcc_h11
				_ = _1474_cf39
				var _1475_v34 _dafny.Sequence
				_ = _1475_v34
				_1475_v34 = _dafny.SeqOf(p1)
				(globalState).F8 = ((_this).Fm0(_1475_v34, (_1475_v34).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("axewyqwhx")).Cardinality()), _dafny.IntOfUint32((_1475_v34).Cardinality()))).Uint32()).(bool), _1405_v0, (_this).F10(), globalState)) && (p0)
				var _1476_v35 _dafny.Sequence
				_ = _1476_v35
				_1476_v35 = _dafny.SeqOf(_1405_v0, _1406_i0)
				var _1477_v36 _dafny.MultiSet
				_ = _1477_v36
				_1477_v36 = _dafny.MultiSetOf(p1, (_this).F10(), (_this).F10())
				var _1478_v37 _dafny.Map
				_ = _1478_v37
				_1478_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1477_v36).Cardinality(), _1406_i0)
				var _1479_v38 _dafny.Array
				_ = _1479_v38
				var _nwElement0_44 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1406_i0))
				_ = _nwElement0_44
				var _nw217 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(12))
				_ = _nw217
				(_nw217).ArraySet1(_nwElement0_44, 0)
				(_nw217).ArraySet1((_this).Fm1((_this).F9(), globalState), 1)
				(_nw217).ArraySet1((_dafny.IntOfUint32((_1475_v34).Cardinality())).Times((_dafny.MultiSetFromSeq(_1476_v35)).Cardinality()), 2)
				(_nw217).ArraySet1((_dafny.Zero).Minus((_this).Fm19(_1405_v0, _1406_i0, globalState)), 3)
				(_nw217).ArraySet1(_1406_i0, 4)
				(_nw217).ArraySet1(_1405_v0, 5)
				(_nw217).ArraySet1(((_1478_v37).Merge(_1478_v37)).Cardinality(), 6)
				(_nw217).ArraySet1(_1405_v0, 7)
				(_nw217).ArraySet1(_1405_v0, 8)
				(_nw217).ArraySet1(_1406_i0, 9)
				(_nw217).ArraySet1(_dafny.IntOfInt64(-769), 10)
				(_nw217).ArraySet1((_dafny.IntOfInt64(197)).Times(_1406_i0), 11)
				_1479_v38 = _nw217
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1479_v38), 0))
				_ = _index233
				(_1479_v38).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(543), ((_1478_v37).Update(_1405_v0, _1406_i0)).Cardinality()), (_index233).Int())
				var _1480_v39 D5
				_ = _1480_v39
				_1480_v39 = Companion_D5_.Create_DC14_(_1405_v0, p1)
				var _1481_v40 _dafny.Sequence
				_ = _1481_v40
				_1481_v40 = _dafny.SeqOf(_1480_v39)
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1479_v38), 0))
				_ = _index234
				var _rhs230 bool = p0
				_ = _rhs230
				var _rhs231 bool = (_this).F10()
				_ = _rhs231
				var _rhs232 _dafny.Int = Companion_Default___.Fm33(p0, _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer89 func(_dafny.Int) D5) func(_dafny.Int) interface{} {
					return func(arg89 _dafny.Int) interface{} {
						return coer89(arg89)
					}
				}((func(_1482_i0 _dafny.Int, _1483_p0 bool) func(_dafny.Int) D5 {
					return func(_1484_i3 _dafny.Int) D5 {
						return Companion_D5_.Create_DC14_(_1482_i0, _1483_p0)
					}
				})(_1406_i0, p0))), _1481_v40), ((func() _dafny.Int {
					if (_1477_v36).Contains((_this).F10()) {
						return (_1477_v36).Multiplicity((_this).F10())
					}
					return _1405_v0
				})()).Times((_dafny.Zero).Minus(_1406_i0)), (_this).Fm0(_1475_v34, p0, _1406_i0, (_this).F10(), globalState), globalState)
				_ = _rhs232
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				var _lhs163 *GlobalState = globalState
				_ = _lhs163
				var _lhs164 _dafny.Array = _1479_v38
				_ = _lhs164
				var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1479_v38), 0))
				_ = _lhs165
				_lhs162.F8 = _rhs230
				_lhs163.F8 = _rhs231
				(_lhs164).ArraySet1(_rhs232, (_lhs165).Int())
				var _1485_v41 _dafny.CodePoint
				_ = _1485_v41
				_1485_v41 = _dafny.CodePoint('o')
				_1485_v41 = _1485_v41
				var _1486_v42 _dafny.Array
				_ = _1486_v42
				var _nwElement0_45 _dafny.Sequence = _1407_v1
				_ = _nwElement0_45
				var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(11))
				_ = _nw218
				(_nw218).ArraySet1(_nwElement0_45, 0)
				(_nw218).ArraySet1(_dafny.SeqOf(_1485_v41, _1485_v41, _1485_v41), 1)
				(_nw218).ArraySet1(_1407_v1, 2)
				(_nw218).ArraySet1(_dafny.SeqOf(_1485_v41), 3)
				(_nw218).ArraySet1(_1407_v1, 4)
				(_nw218).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1407_v1, _1407_v1), 5)
				(_nw218).ArraySet1(_1407_v1, 6)
				(_nw218).ArraySet1(Companion_Default___.Fm20(_dafny.IntOfInt64(753), p1, !(false), (_1479_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1479_v38), 0))).Int()).(_dafny.Int), globalState), 7)
				(_nw218).ArraySet1(_1407_v1, 8)
				(_nw218).ArraySet1(_dafny.SeqOf(_1485_v41), 9)
				(_nw218).ArraySet1(_1407_v1, 10)
				_1486_v42 = _nw218
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1486_v42), 0))
				_ = _index235
				(_1486_v42).ArraySet1(_1407_v1, (_index235).Int())
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1486_v42), 0))
				_ = _index236
				(_1486_v42).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1407_v1, _1407_v1), (_index236).Int())
			}
			if false {
				var _1487_v43 _dafny.Map
				_ = _1487_v43
				_1487_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1406_i0, (_this).F10())
				var _1488_v45 _dafny.Sequence
				_ = _1488_v45
				_1488_v45 = _dafny.SeqOf(_1487_v43, _1487_v43, (_1487_v43).Merge(func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(185), _dafny.IntOfInt64(321))); ; {
						_compr_53, _ok58 := _iter58()
						if !_ok58 {
							break
						}
						var _1489_v44 _dafny.Int
						_1489_v44 = interface{}(_compr_53).(_dafny.Int)
						if ((_dafny.IntOfInt64(185)).Cmp(_1489_v44) <= 0) && ((_1489_v44).Cmp(_dafny.IntOfInt64(321)) < 0) {
							_coll53.Add((_1489_v44).Times(_dafny.IntOfInt64(887)), p0)
						}
					}
					return _coll53.ToMap()
				}()), (Companion_Default___.Fm44(p1, globalState)).Merge(Companion_Default___.Fm44(p1, globalState)), _1487_v43)
				var _1490_v46 D16
				_ = _1490_v46
				_1490_v46 = Companion_D16_.Create_DC40_(_1488_v45)
				_1488_v45 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1487_v43), (_1490_v46).Dtor_cf76()), _dafny.Companion_Sequence_.Concatenate(_1488_v45, _1488_v45)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1406_i0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1487_v43), (_1490_v46).Dtor_cf76()), _dafny.Companion_Sequence_.Concatenate(_1488_v45, _1488_v45))).Cardinality()))).Uint32(), _1487_v43)
				var _1491_v47 _dafny.Sequence
				_ = _1491_v47
				_1491_v47 = _dafny.SeqOf((_this).F10(), p0, p1)
				(globalState).F8 = !((_this).Fm0(_1491_v47, (p1) == ((_this).F10()), Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(_dafny.IntOfInt64(-494))).Cardinality(), _1405_v0), p0, globalState))
				var _1492_v48 _dafny.Array
				_ = _1492_v48
				var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw219
				_1492_v48 = _nw219
				var _1493_v49 _dafny.Map
				_ = _1493_v49
				_1493_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _1492_v48)
				var _1494_v50 _dafny.CodePoint
				_ = _1494_v50
				_1494_v50 = _dafny.CodePoint('u')
				var _1495_v51 *C3
				_ = _1495_v51
				var _nw220 *C3 = New_C3_()
				_ = _nw220
				_nw220.Ctor__()
				_1495_v51 = _nw220
				var _1496_v52 _dafny.Map
				_ = _1496_v52
				_1496_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if (_1493_v49).Contains(_1494_v50) {
						return (_1493_v49).Get(_1494_v50).(_dafny.Array)
					}
					return _1492_v48
				})(), _1495_v51)
				_1496_v52 = ((_1496_v52).Merge(_1496_v52)).Merge(_1496_v52)
				_1494_v50 = _1494_v50
				_1405_v0 = _1406_i0
			} else {
				var _1497_v53 D1
				_ = _1497_v53
				_1497_v53 = Companion_D1_.Create_DC3_(_1407_v1)
				var _1498_v54 _dafny.Sequence
				_ = _1498_v54
				_1498_v54 = _dafny.SeqOf((_1405_v0).Plus(_1405_v0), _dafny.IntOfUint32(((_1497_v53).Dtor_cf5()).Cardinality()))
				_1498_v54 = _1498_v54
				_1405_v0 = (_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_1407_v1).Cardinality()))).Plus(_1406_i0))
				var _1499_v55 _dafny.Array
				_ = _1499_v55
				var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw221
				_1499_v55 = _nw221
				var _1500_v56 _dafny.Sequence
				_ = _1500_v56
				_1500_v56 = _dafny.SeqOf(p1, p0)
				var _1501_v57 _dafny.MultiSet
				_ = _1501_v57
				_1501_v57 = _dafny.MultiSetOf(_1500_v56, _1500_v56)
				var _1502_v58 _dafny.Map
				_ = _1502_v58
				_1502_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ayxnih")).Cardinality()), _1406_i0)
				var _rhs233 _dafny.Array = _1499_v55
				_ = _rhs233
				var _rhs234 _dafny.Int = (func() _dafny.Int {
					if (_1405_v0).Cmp(_1406_i0) >= 0 {
						return Companion_Default___.SafeDivisionInt(_1405_v0, (_1501_v57).Cardinality())
					}
					return (func() _dafny.Int {
						if (_1502_v58).Contains((_dafny.Zero).Minus(_1405_v0)) {
							return (_1502_v58).Get((_dafny.Zero).Minus(_1405_v0)).(_dafny.Int)
						}
						return _1406_i0
					})()
				})()
				_ = _rhs234
				var _rhs235 _dafny.Int = (_dafny.Zero).Minus(_1406_i0)
				_ = _rhs235
				_1499_v55 = _rhs233
				_1405_v0 = _rhs234
				_1405_v0 = _rhs235
				var _1503_v59 _dafny.Set
				_ = _1503_v59
				_1503_v59 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_1498_v54, _1498_v54))
				_1405_v0 = (_1503_v59).Cardinality()
				var _1504_v60 _dafny.MultiSet
				_ = _1504_v60
				_1504_v60 = _dafny.MultiSetOf(false, true, false, p1, p1)
				var _1505_v61 _dafny.Set
				_ = _1505_v61
				_1505_v61 = _dafny.SetOf(_1499_v55, _1499_v55, _1499_v55, _1499_v55, _1499_v55)
				var _1506_v62 D8
				_ = _1506_v62
				_1506_v62 = Companion_D8_.Create_DC24_((_1504_v60).IsProperSubsetOf(_1504_v60), _1505_v61, _1405_v0)
				var _rhs236 D8 = _1506_v62
				_ = _rhs236
				var _rhs237 bool = p0
				_ = _rhs237
				var _rhs238 bool = p0
				_ = _rhs238
				var _rhs239 _dafny.Int = (_1405_v0).Plus(_1405_v0)
				_ = _rhs239
				var _rhs240 bool = (_this).F10()
				_ = _rhs240
				var _lhs166 *GlobalState = globalState
				_ = _lhs166
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				_1506_v62 = _rhs236
				_lhs166.F8 = _rhs237
				_lhs167.F8 = _rhs238
				_1405_v0 = _rhs239
				_lhs168.F8 = _rhs240
			}
			var _1507_v63 _dafny.Array
			_ = _1507_v63
			var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
			_ = _nw222
			_1507_v63 = _nw222
			var _1508_v64 _dafny.Map
			_ = _1508_v64
			_1508_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.IntOfUint32((_1407_v1).Cardinality()))
			var _1509_v65 _dafny.Sequence
			_ = _1509_v65
			_1509_v65 = _dafny.SeqOf(_1406_i0, _1406_i0)
			var _1510_v66 _dafny.Map
			_ = _1510_v66
			_1510_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1508_v64, _1509_v65)
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_1507_v63), 0))
			_ = _index237
			(_1507_v63).ArraySet1((_1510_v66).Merge(_1510_v66), (_index237).Int())
			var _1511_v67 _dafny.Sequence
			_ = _1511_v67
			_1511_v67 = _dafny.SeqOf(_1510_v66, _1510_v66, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1508_v64, _1509_v65))
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_1507_v63), 0))
			_ = _index238
			(_1507_v63).ArraySet1(((_1510_v66).Merge((_1511_v67).Select((Companion_Default___.SafeIndex(_1405_v0, _dafny.IntOfUint32((_1511_v67).Cardinality()))).Uint32()).(_dafny.Map))).Merge(_1510_v66), (_index238).Int())
		}
		(globalState).F8 = p1
		var _1512_v68 _dafny.MultiSet
		_ = _1512_v68
		_1512_v68 = _dafny.MultiSetOf((_1405_v0).Minus(_1405_v0), (_this).Fm1((_this).F9(), globalState), _dafny.IntOfInt64(208), _1405_v0)
		_1512_v68 = (_1512_v68).Update(Companion_Default___.SafeDivisionInt(_1405_v0, _1405_v0), Companion_Default___.Abs(_1405_v0))
		(globalState).F8 = p0
		var _1513_v69 D5
		_ = _1513_v69
		_1513_v69 = Companion_D5_.Create_DC14_(_dafny.IntOfInt64(75), p1)
		var _1514_v70 _dafny.Array
		_ = _1514_v70
		var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
		_ = _nw223
		_1514_v70 = _nw223
		var _1515_v71 _dafny.Set
		_ = _1515_v71
		_1515_v71 = _dafny.SetOf(_1514_v70)
		var _1516_v72 _dafny.Sequence
		_ = _1516_v72
		_1516_v72 = _dafny.SeqOf(_1405_v0)
		var _1517_v73 _dafny.Sequence
		_ = _1517_v73
		_1517_v73 = _dafny.UnicodeSeqOfUtf8Bytes("jeudkbik")
		var _1518_v74 _dafny.Map
		_ = _1518_v74
		_1518_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1405_v0, p0)
		var _1519_v75 _dafny.CodePoint
		_ = _1519_v75
		_1519_v75 = _dafny.CodePoint('e')
		var _1520_v76 _dafny.Array
		_ = _1520_v76
		var _nwElement0_46 _dafny.Int = _1405_v0
		_ = _nwElement0_46
		var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(27))
		_ = _nw224
		(_nw224).ArraySet1(_nwElement0_46, 0)
		(_nw224).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1513_v69).Dtor_cf25()), (_1515_v71).Cardinality())).Cardinality(), 1)
		(_nw224).ArraySet1(_1405_v0, 2)
		(_nw224).ArraySet1(_1405_v0, 3)
		(_nw224).ArraySet1(_dafny.IntOfInt64(-813), 4)
		(_nw224).ArraySet1(_1405_v0, 5)
		(_nw224).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1516_v72, _1516_v72)).Cardinality())), 6)
		(_nw224).ArraySet1(_1405_v0, 7)
		(_nw224).ArraySet1(_dafny.IntOfUint32((_1517_v73).Cardinality()), 8)
		(_nw224).ArraySet1((_1518_v74).Cardinality(), 9)
		(_nw224).ArraySet1((_dafny.Zero).Minus(_1405_v0), 10)
		(_nw224).ArraySet1((_1405_v0).Times(_dafny.IntOfInt64(184)), 11)
		(_nw224).ArraySet1(_1405_v0, 12)
		(_nw224).ArraySet1(_1405_v0, 13)
		(_nw224).ArraySet1(_1405_v0, 14)
		(_nw224).ArraySet1(_1405_v0, 15)
		(_nw224).ArraySet1((_1405_v0).Minus(_1405_v0), 16)
		(_nw224).ArraySet1((_1516_v72).Select((Companion_Default___.SafeIndex(_1405_v0, _dafny.IntOfUint32((_1516_v72).Cardinality()))).Uint32()).(_dafny.Int), 17)
		(_nw224).ArraySet1(_1405_v0, 18)
		(_nw224).ArraySet1(Companion_Default___.SafeModuloInt(_1405_v0, _1405_v0), 19)
		(_nw224).ArraySet1(_dafny.IntOfInt64(822), 20)
		(_nw224).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1405_v0, _1405_v0)), 21)
		(_nw224).ArraySet1(_dafny.IntOfUint32((_1517_v73).Cardinality()), 22)
		(_nw224).ArraySet1(_dafny.IntOfInt64(104), 23)
		(_nw224).ArraySet1(_1405_v0, 24)
		(_nw224).ArraySet1((_1405_v0).Minus(_dafny.IntOfInt64(464)), 25)
		(_nw224).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1519_v75, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nbqhra"), (Companion_Default___.SafeIndex(_1405_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nbqhra")).Cardinality()))).Uint32(), _1519_v75))).Cardinality(), 26)
		_1520_v76 = _nw224
		var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1514_v70), 0))
		_ = _index239
		(_1514_v70).ArraySet1(_1405_v0, (_index239).Int())
		var _1521_v77 D6
		_ = _1521_v77
		_1521_v77 = Companion_D6_.Create_DC18_(_dafny.CodePoint('m'), _1405_v0, _1405_v0)
		var _pat_let_tv31 = p1
		_ = _pat_let_tv31
		var _pat_let_tv32 = _1405_v0
		_ = _pat_let_tv32
		var _pat_let_tv33 = _1405_v0
		_ = _pat_let_tv33
		var _pat_let_tv34 = _1517_v73
		_ = _pat_let_tv34
		var _pat_let_tv35 = _1517_v73
		_ = _pat_let_tv35
		var _pat_let_tv36 = _1517_v73
		_ = _pat_let_tv36
		var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1514_v70), 0))
		_ = _index240
		(_1514_v70).ArraySet1(func(_source20 D6) _dafny.Int {
			if _source20.Is_DC17() {
				var _1522___mcc_h12 bool = _source20.Get_().(D6_DC17).Cf29
				_ = _1522___mcc_h12
				var _1523___mcc_h13 bool = _source20.Get_().(D6_DC17).Cf30
				_ = _1523___mcc_h13
				var _1524___mcc_h14 _dafny.Int = _source20.Get_().(D6_DC17).Cf31
				_ = _1524___mcc_h14
				var _1525___mcc_h15 _dafny.Int = _source20.Get_().(D6_DC17).Cf32
				_ = _1525___mcc_h15
				var _1526___mcc_h16 _dafny.Int = _source20.Get_().(D6_DC17).Cf33
				_ = _1526___mcc_h16
				var _1527_cf33 _dafny.Int = _1526___mcc_h16
				_ = _1527_cf33
				var _1528_cf32 _dafny.Int = _1525___mcc_h15
				_ = _1528_cf32
				var _1529_cf31 _dafny.Int = _1524___mcc_h14
				_ = _1529_cf31
				var _1530_cf30 bool = _1523___mcc_h13
				_ = _1530_cf30
				var _1531_cf29 bool = _1522___mcc_h12
				_ = _1531_cf29
				return ((_dafny.SetOf(_1530_cf30, _pat_let_tv31)).Intersection(_dafny.SetOf(_1531_cf29))).Cardinality()
			} else if _source20.Is_DC18() {
				var _1532___mcc_h17 _dafny.CodePoint = _source20.Get_().(D6_DC18).Cf34
				_ = _1532___mcc_h17
				var _1533___mcc_h18 _dafny.Int = _source20.Get_().(D6_DC18).Cf35
				_ = _1533___mcc_h18
				var _1534___mcc_h19 _dafny.Int = _source20.Get_().(D6_DC18).Cf36
				_ = _1534___mcc_h19
				var _1535_cf36 _dafny.Int = _1534___mcc_h19
				_ = _1535_cf36
				var _1536_cf35 _dafny.Int = _1533___mcc_h18
				_ = _1536_cf35
				var _1537_cf34 _dafny.CodePoint = _1532___mcc_h17
				_ = _1537_cf34
				return _pat_let_tv32
			} else if _source20.Is_DC19() {
				var _1538___mcc_h20 bool = _source20.Get_().(D6_DC19).Cf37
				_ = _1538___mcc_h20
				var _1539___mcc_h21 _dafny.CodePoint = _source20.Get_().(D6_DC19).Cf38
				_ = _1539___mcc_h21
				var _1540_cf38 _dafny.CodePoint = _1539___mcc_h21
				_ = _1540_cf38
				var _1541_cf37 bool = _1538___mcc_h20
				_ = _1541_cf37
				return _pat_let_tv33
			} else if _source20.Is_DC16() {
				var _1542___mcc_h22 _dafny.Map = _source20.Get_().(D6_DC16).Cf28
				_ = _1542___mcc_h22
				var _1543_cf28 _dafny.Map = _1542___mcc_h22
				_ = _1543_cf28
				return (_dafny.IntOfUint32((_pat_let_tv34).Cardinality())).Minus((_dafny.Zero).Minus((_1543_cf28).Cardinality()))
			} else {
				var _1544___mcc_h23 D6 = _source20.Get_().(D6_DC20).Cf39
				_ = _1544___mcc_h23
				var _1545_cf39 D6 = _1544___mcc_h23
				_ = _1545_cf39
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv35, _pat_let_tv36)).Cardinality()))
			}
		}(_1521_v77), (_index240).Int())
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1520_v76), 0))); ; {
			_guard_loop_5, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _1546_i4 _dafny.Int
			_1546_i4 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1546_i4).Sign() != -1) && ((_1546_i4).Cmp(_dafny.ArrayLen((_1520_v76), 0)) < 0)) {
				(_1520_v76).ArraySet1((_1546_i4).Times(_dafny.IntOfInt64(703)), (_1546_i4).Int())
			}
		}
		var _1547_v78 _dafny.Array
		_ = _1547_v78
		var _nw225 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw225
		_1547_v78 = _nw225
		r0 = _1547_v78
		var _1548_v79 D0
		_ = _1548_v79
		_1548_v79 = Companion_D0_.Create_DC0_(p1)
		var _1549_v80 _dafny.Sequence
		_ = _1549_v80
		_1549_v80 = _dafny.SeqOf(_1548_v79, _1548_v79, _1548_v79, Companion_D0_.Create_DC0_(p0))
		r1 = _1549_v80
		var _1550_v82 D4
		_ = _1550_v82
		_1550_v82 = Companion_D4_.Create_DC11_(_dafny.IntOfInt64(919), func() _dafny.Set {
			var _coll54 = _dafny.NewBuilder()
			_ = _coll54
			for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(742), _dafny.IntOfInt64(-453))); ; {
				_compr_54, _ok60 := _iter60()
				if !_ok60 {
					break
				}
				var _1551_v81 _dafny.Int
				_1551_v81 = interface{}(_compr_54).(_dafny.Int)
				if ((_dafny.IntOfInt64(742)).Cmp(_1551_v81) <= 0) && ((_1551_v81).Cmp(_dafny.IntOfInt64(-453)) < 0) {
					_coll54.Add((_1551_v81).Minus((_1514_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_1514_v70), 0))).Int()).(_dafny.Int)))
				}
			}
			return _coll54.ToSet()
		}(), true)
		r2 = ((func() D4 {
			if false {
				return Companion_D4_.Create_DC11_(_1405_v0, _dafny.SetOf(_dafny.IntOfUint32((_1517_v73).Cardinality())), p1)
			}
			return _1550_v82
		})()).Dtor_cf17()
		return r0, r1, r2
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f9  _dafny.Map
	_f10 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f9 = _dafny.EmptyMap
	_this._f10 = false
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

func (_this *C6) F9() _dafny.Map {
	return _this._f9
}
func (_this *C6) F10() bool {
	return _this._f10
}
func (_this *C6) Ctor__(f9 _dafny.Map, f10 bool) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C6) Fm0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return (_this).F10()
	}
}
func (_this *C6) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(93)
	}
}
func (_this *C6) Fm8(globalState *GlobalState) _dafny.Int {
	{
		var _source21 D1 = Companion_D1_.Create_DC4_((_this).F10(), (_this).F10(), _dafny.IntOfInt64(-971))
		_ = _source21
		if _source21.Is_DC3() {
			var _1552___mcc_h0 _dafny.Sequence = _source21.Get_().(D1_DC3).Cf5
			_ = _1552___mcc_h0
			var _1553_cf5 _dafny.Sequence = _1552___mcc_h0
			_ = _1553_cf5
			return (_dafny.IntOfUint32((_1553_cf5).Cardinality())).Minus(_dafny.IntOfInt64(-673))
		} else if _source21.Is_DC4() {
			var _1554___mcc_h1 bool = _source21.Get_().(D1_DC4).Cf6
			_ = _1554___mcc_h1
			var _1555___mcc_h2 bool = _source21.Get_().(D1_DC4).Cf7
			_ = _1555___mcc_h2
			var _1556___mcc_h3 _dafny.Int = _source21.Get_().(D1_DC4).Cf8
			_ = _1556___mcc_h3
			var _1557_cf8 _dafny.Int = _1556___mcc_h3
			_ = _1557_cf8
			var _1558_cf7 bool = _1555___mcc_h2
			_ = _1558_cf7
			var _1559_cf6 bool = _1554___mcc_h1
			_ = _1559_cf6
			return _1557_cf8
		} else if _source21.Is_DC2() {
			var _1560___mcc_h4 _dafny.Sequence = _source21.Get_().(D1_DC2).Cf4
			_ = _1560___mcc_h4
			var _1561_cf4 _dafny.Sequence = _1560___mcc_h4
			_ = _1561_cf4
			return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(862))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg90 _dafny.Int) interface{} {
					return coer90(arg90)
				}
			}(func(_1562_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf((_this).F10(), (Companion_D0_.Create_DC1_(_dafny.IntOfInt64(576), (_this).F10(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qfvjs")).Cardinality()))).Dtor_cf2()), _dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F10())))).Cardinality())
			}))).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F10(), true, (_this).F10(), (_this).F10(), (_this).F10())).Cardinality()))
		} else {
			var _1563___mcc_h5 D1 = _source21.Get_().(D1_DC5).Cf9
			_ = _1563___mcc_h5
			var _1564_cf9 D1 = _1563___mcc_h5
			_ = _1564_cf9
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(753))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg91 _dafny.Int) interface{} {
					return coer91(arg91)
				}
			}(func(_1565_i1 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_dafny.IntOfInt64(462), _dafny.IntOfInt64(3))
			}))).Cardinality()))
		}
	}
}
func (_this *C6) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) bool {
	{
		return !((_this).F10())
	}
}
func (_this *C6) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		if !(!((_this).F10()) || (true)) {
			var _1566_v0 _dafny.Array
			_ = _1566_v0
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_38
			var _nw226 _dafny.Array
			_ = _nw226
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw226 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) _dafny.Int = func(_1567_i0 _dafny.Int) _dafny.Int {
					return (_1567_i0).Minus(((_this).F9()).Cardinality())
				}
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw226 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw226).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw226).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1566_v0 = _nw226
			var _1568_v1 _dafny.Set
			_ = _1568_v1
			_1568_v1 = _dafny.SetOf((_this).F10(), (_this).F10(), (_this).F10())
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1566_v0), 0))
			_ = _index241
			(_1566_v0).ArraySet1(((_1568_v1).Cardinality()).Minus(((_this).F9()).Cardinality()), (_index241).Int())
			var _1569_v2 _dafny.Int
			_ = _1569_v2
			_1569_v2 = _dafny.IntOfInt64(-316)
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1566_v0), 0))
			_ = _index242
			(_1566_v0).ArraySet1((_1569_v2).Times(_1569_v2), (_index242).Int())
			var _1570_v4 _dafny.Array
			_ = _1570_v4
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_39
			var _nw227 _dafny.Array
			_ = _nw227
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw227 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Set = (func(_1571_v0 _dafny.Array) func(_dafny.Int) _dafny.Set {
					return func(_1572_i1 _dafny.Int) _dafny.Set {
						return (func() _dafny.Set {
							var _coll55 = _dafny.NewBuilder()
							_ = _coll55
							for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(249), _dafny.IntOfInt64(61))); ; {
								_compr_55, _ok61 := _iter61()
								if !_ok61 {
									break
								}
								var _1573_v3 _dafny.Int
								_1573_v3 = interface{}(_compr_55).(_dafny.Int)
								if ((_dafny.IntOfInt64(249)).Cmp(_1573_v3) <= 0) && ((_1573_v3).Cmp(_dafny.IntOfInt64(61)) < 0) {
									_coll55.Add((_1573_v3).Times((_1571_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1571_v0), 0))).Int()).(_dafny.Int)))
								}
							}
							return _coll55.ToSet()
						}()).Difference(_dafny.SetOf((_1571_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1571_v0), 0))).Int()).(_dafny.Int)))
					}
				})(_1566_v0)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw227 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw227).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw227).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1570_v4 = _nw227
			_1570_v4 = _1570_v4
			var _1574_v5 _dafny.Array
			_ = _1574_v5
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_40
			var _nw228 _dafny.Array
			_ = _nw228
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw228 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) D0 = func(_1575_i2 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_((_this).F10())
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw228 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw228).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw228).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1574_v5 = _nw228
			var _1576_v6 D0
			_ = _1576_v6
			_1576_v6 = Companion_D0_.Create_DC0_((_this).F10())
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1574_v5), 0))
			_ = _index243
			(_1574_v5).ArraySet1(func(_pat_let34_0 D0) D0 {
				return func(_1577_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let35_0 bool) D0 {
						return func(_1578_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_1578_dt__update_hcf0_h0)
						}(_pat_let35_0)
					}(!(!((_this).F10())))
				}(_pat_let34_0)
			}(_1576_v6), (_index243).Int())
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_1574_v5), 0))
			_ = _index244
			(_1574_v5).ArraySet1(func(_pat_let36_0 D0) D0 {
				return func(_1579_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let37_0 bool) D0 {
						return func(_1580_dt__update_hcf0_h1 bool) D0 {
							return Companion_D0_.Create_DC0_(_1580_dt__update_hcf0_h1)
						}(_pat_let37_0)
					}((_this).F10())
				}(_pat_let36_0)
			}(_1576_v6), (_index244).Int())
			var _1581_v7 *C0
			_ = _1581_v7
			var _nw229 *C0 = New_C0_()
			_ = _nw229
			_nw229.Ctor__()
			_1581_v7 = _nw229
			var _1582_v8 _dafny.Sequence
			_ = _1582_v8
			_1582_v8 = _dafny.UnicodeSeqOfUtf8Bytes("qbbn")
			var _1583_v9 _dafny.Map
			_ = _1583_v9
			_1583_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1566_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1566_v0), 0))).Int()).(_dafny.Int), _1569_v2)
			var _1584_v10 _dafny.MultiSet
			_ = _1584_v10
			_1584_v10 = _dafny.MultiSetOf((_1566_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1566_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(312))
			(globalState).F8 = (Companion_Default___.Fm10((_this).F10(), _1569_v2, _1582_v8, _1583_v9, globalState)).IsDisjointFrom(_1584_v10)
		} else {
			var _1585_v11 *C0
			_ = _1585_v11
			var _nw230 *C0 = New_C0_()
			_ = _nw230
			_nw230.Ctor__()
			_1585_v11 = _nw230
			var _1586_v12 _dafny.Array
			_ = _1586_v12
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_41
			var _nw231 _dafny.Array
			_ = _nw231
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw231 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) bool = func(_1587_i3 _dafny.Int) bool {
					return (_this).F10()
				}
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw231 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw231).ArraySet1(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw231).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1586_v12 = _nw231
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))
			_ = _index245
			(_1586_v12).ArraySet1(((_this).F10()) == ((_this).F10()), (_index245).Int())
			var _1588_v13 _dafny.Int
			_ = _1588_v13
			_1588_v13 = _dafny.IntOfInt64(274)
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))
			_ = _index246
			(_1586_v12).ArraySet1(((_dafny.Zero).Minus(_1588_v13)).Cmp(_1588_v13) > 0, (_index246).Int())
			var _1589_v14 D1
			_ = _1589_v14
			_1589_v14 = Companion_D1_.Create_DC4_((_this).F10(), false, (_this).Fm1((_this).F9(), globalState))
			var _rhs241 D1 = _1589_v14
			_ = _rhs241
			var _rhs242 bool = (_this).F10()
			_ = _rhs242
			var _rhs243 _dafny.Int = _1588_v13
			_ = _rhs243
			var _rhs244 bool = false
			_ = _rhs244
			var _lhs169 *GlobalState = globalState
			_ = _lhs169
			var _lhs170 *GlobalState = globalState
			_ = _lhs170
			_1589_v14 = _rhs241
			_lhs169.F8 = _rhs242
			_1588_v13 = _rhs243
			_lhs170.F8 = _rhs244
			var _1590_v15 _dafny.Map
			_ = _1590_v15
			_1590_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, (_this).F9())
			if ((func() _dafny.Map {
				if (_1590_v15).Contains(_1588_v13) {
					return (_1590_v15).Get(_1588_v13).(_dafny.Map)
				}
				return (_this).F9()
			})()).Equals((_this).F9()) {
				var _1591_v16 _dafny.Map
				_ = _1591_v16
				_1591_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, _1588_v13)
				var _1592_v17 _dafny.MultiSet
				_ = _1592_v17
				_1592_v17 = _dafny.MultiSetOf((_this).F10(), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))
				_1588_v13 = (_1588_v13).Times((func() _dafny.Int {
					if (_1591_v16).Contains((_1592_v17).Cardinality()) {
						return (_1591_v16).Get((_1592_v17).Cardinality()).(_dafny.Int)
					}
					return _1588_v13
				})())
				var _1593_v18 _dafny.Map
				_ = _1593_v18
				_1593_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))
				var _1594_v19 _dafny.MultiSet
				_ = _1594_v19
				_1594_v19 = _dafny.MultiSetOf(_1588_v13)
				var _1595_v20 _dafny.Set
				_ = _1595_v20
				_1595_v20 = _dafny.SetOf(_1594_v19)
				var _1596_v22 _dafny.Sequence
				_ = _1596_v22
				_1596_v22 = _dafny.SeqOf(_1588_v13, _1588_v13)
				var _1597_v23 _dafny.Array
				_ = _1597_v23
				var _nwElement0_47 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))).Merge(_1593_v18)
				_ = _nwElement0_47
				var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(21))
				_ = _nw232
				(_nw232).ArraySet1(_nwElement0_47, 0)
				(_nw232).ArraySet1(_1593_v18, 1)
				(_nw232).ArraySet1((Companion_Default___.Fm11((_1595_v20).Cardinality(), _1588_v13, false, _1588_v13, globalState)).Merge(_1593_v18), 2)
				(_nw232).ArraySet1((_1593_v18).Merge(_1593_v18), 3)
				(_nw232).ArraySet1(_1593_v18, 4)
				(_nw232).ArraySet1(_1593_v18, 5)
				(_nw232).ArraySet1(_1593_v18, 6)
				(_nw232).ArraySet1(_1593_v18, 7)
				(_nw232).ArraySet1(_1593_v18, 8)
				(_nw232).ArraySet1(_1593_v18, 9)
				(_nw232).ArraySet1(func() _dafny.Map {
					var _coll56 = _dafny.NewMapBuilder()
					_ = _coll56
					for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.IntOfInt64(445))); ; {
						_compr_56, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1598_v21 _dafny.Int
						_1598_v21 = interface{}(_compr_56).(_dafny.Int)
						if ((_1598_v21).Sign() != -1) && ((_1598_v21).Cmp(_dafny.IntOfInt64(445)) < 0) {
							_coll56.Add((_1598_v21).Times(_dafny.IntOfInt64(351)), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))
						}
					}
					return _coll56.ToMap()
				}(), 10)
				(_nw232).ArraySet1(Companion_Default___.Fm11((_1596_v22).Select((Companion_Default___.SafeIndex((_1591_v16).Cardinality(), _dafny.IntOfUint32((_1596_v22).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(439), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool), _1588_v13, globalState), 11)
				(_nw232).ArraySet1(_1593_v18, 12)
				(_nw232).ArraySet1(_1593_v18, 13)
				(_nw232).ArraySet1((_1593_v18).Merge(_1593_v18), 14)
				(_nw232).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool)), 15)
				(_nw232).ArraySet1(_1593_v18, 16)
				(_nw232).ArraySet1(_1593_v18, 17)
				(_nw232).ArraySet1(Companion_Default___.Fm11(_dafny.IntOfInt64(-251), _1588_v13, (_this).F10(), _1588_v13, globalState), 18)
				(_nw232).ArraySet1(_1593_v18, 19)
				(_nw232).ArraySet1(_1593_v18, 20)
				_1597_v23 = _nw232
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1597_v23), 0))
				_ = _index247
				(_1597_v23).ArraySet1((func() _dafny.Map {
					if (_this).F10() {
						return _1593_v18
					}
					return _1593_v18
				})(), (_index247).Int())
				var _1599_v24 _dafny.Map
				_ = _1599_v24
				_1599_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1593_v18).Merge(_1593_v18))
				var _1600_v25 _dafny.Sequence
				_ = _1600_v25
				_1600_v25 = _dafny.UnicodeSeqOfUtf8Bytes("tnx")
				var _1601_v26 _dafny.Set
				_ = _1601_v26
				_1601_v26 = _dafny.SetOf(_1600_v25)
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1597_v23), 0))
				_ = _index248
				var _rhs245 _dafny.Int = _dafny.IntOfInt64(-464)
				_ = _rhs245
				var _rhs246 bool = (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool)
				_ = _rhs246
				var _rhs247 _dafny.Map = (func() _dafny.Map {
					if (_1599_v24).Contains(!((_1601_v26).IsSubsetOf(_1601_v26))) {
						return (_1599_v24).Get(!((_1601_v26).IsSubsetOf(_1601_v26))).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))
				})()
				_ = _rhs247
				var _rhs248 bool = !_dafny.Companion_Sequence_.Equal(_1600_v25, _dafny.UnicodeSeqOfUtf8Bytes("xiesenboe"))
				_ = _rhs248
				var _lhs171 _dafny.Array = _1597_v23
				_ = _lhs171
				var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1597_v23), 0))
				_ = _lhs172
				_1588_v13 = _rhs245
				r1 = _rhs246
				(_lhs171).ArraySet1(_rhs247, (_lhs172).Int())
				r1 = _rhs248
				var _1602_v27 *C0
				_ = _1602_v27
				var _nw233 *C0 = New_C0_()
				_ = _nw233
				_nw233.Ctor__()
				_1602_v27 = _nw233
				var _1603_v28 *C0
				_ = _1603_v28
				var _nw234 *C0 = New_C0_()
				_ = _nw234
				_nw234.Ctor__()
				_1603_v28 = _nw234
				(globalState).F8 = (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool)
			} else {
				var _1604_v29 _dafny.Map
				_ = _1604_v29
				_1604_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm13(_1588_v13, globalState)).Cardinality()), Companion_Default___.Fm14(globalState))
				var _1605_v30 _dafny.CodePoint
				_ = _1605_v30
				_1605_v30 = _dafny.CodePoint('q')
				var _1606_v33 _dafny.Sequence
				_ = _1606_v33
				_1606_v33 = _dafny.UnicodeSeqOfUtf8Bytes("qxxtkeihl")
				var _1607_v34 _dafny.Map
				_ = _1607_v34
				_1607_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rvcpxn"), _1588_v13)
				var _1608_v35 _dafny.Map
				_ = _1608_v35
				_1608_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1607_v34).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, _1605_v30))
				var _1609_v36 _dafny.Sequence
				_ = _1609_v36
				_1609_v36 = _dafny.SeqOf(_1588_v13, _1588_v13)
				var _1610_v37 _dafny.Array
				_ = _1610_v37
				var _nwElement0_48 _dafny.Map = (Companion_Default___.Fm12((_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool), _1588_v13, (_dafny.Zero).Minus(_1588_v13), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool), globalState)).Merge(_1604_v29)
				_ = _nwElement0_48
				var _nw235 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(19))
				_ = _nw235
				(_nw235).ArraySet1(_nwElement0_48, 0)
				(_nw235).ArraySet1(_1604_v29, 1)
				(_nw235).ArraySet1((_1604_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, _1605_v30)), 2)
				(_nw235).ArraySet1(_1604_v29, 3)
				(_nw235).ArraySet1((_1604_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, _1605_v30)), 4)
				(_nw235).ArraySet1(_1604_v29, 5)
				(_nw235).ArraySet1(func() _dafny.Map {
					var _coll57 = _dafny.NewMapBuilder()
					_ = _coll57
					for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(123), _dafny.IntOfInt64(-659))); ; {
						_compr_57, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1611_v31 _dafny.Int
						_1611_v31 = interface{}(_compr_57).(_dafny.Int)
						if ((_dafny.IntOfInt64(123)).Cmp(_1611_v31) <= 0) && ((_1611_v31).Cmp(_dafny.IntOfInt64(-659)) < 0) {
							_coll57.Add((_1611_v31).Plus(_1588_v13), _1605_v30)
						}
					}
					return _coll57.ToMap()
				}(), 6)
				(_nw235).ArraySet1((_1604_v29).Merge(_1604_v29), 7)
				(_nw235).ArraySet1(func() _dafny.Map {
					var _coll58 = _dafny.NewMapBuilder()
					_ = _coll58
					for _iter64 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg92 _dafny.Int) interface{} {
							return coer92(arg92)
						}
					}(func(_1612_i4 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(360)
					}))).Elements()); ; {
						_compr_58, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _1613_v32 _dafny.Int
						_1613_v32 = interface{}(_compr_58).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg93 _dafny.Int) interface{} {
								return coer93(arg93)
							}
						}(func(_1612_i4 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(360)
						})), _1613_v32) {
							_coll58.Add((_1613_v32).Times((_dafny.MultiSetOf((_this).F10())).Cardinality()), _1605_v30)
						}
					}
					return _coll58.ToMap()
				}(), 8)
				(_nw235).ArraySet1(_1604_v29, 9)
				(_nw235).ArraySet1((Companion_Default___.Fm12(!((_this).F10()), _1588_v13, _dafny.IntOfUint32((_1606_v33).Cardinality()), (_this).F10(), globalState)).Merge(_1604_v29), 10)
				(_nw235).ArraySet1(_1604_v29, 11)
				(_nw235).ArraySet1((_1604_v29).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, _1605_v30)), 12)
				(_nw235).ArraySet1(((_1604_v29).Update(_1588_v13, _1605_v30)).Merge((func() _dafny.Map {
					if (_1608_v35).Contains(_1588_v13) {
						return (_1608_v35).Get(_1588_v13).(_dafny.Map)
					}
					return Companion_Default___.Fm12(true, _1588_v13, _1588_v13, (_this).F10(), globalState)
				})()), 13)
				(_nw235).ArraySet1(Companion_Default___.Fm12((_this).F10(), _dafny.IntOfUint32((_1606_v33).Cardinality()), _dafny.IntOfUint32((_1609_v36).Cardinality()), (_this).F10(), globalState), 14)
				(_nw235).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1588_v13, _1605_v30), 15)
				(_nw235).ArraySet1(_1604_v29, 16)
				(_nw235).ArraySet1(_1604_v29, 17)
				(_nw235).ArraySet1(_1604_v29, 18)
				_1610_v37 = _nw235
				_1610_v37 = _1610_v37
				_1606_v33 = _1606_v33
				var _1614_v38 _dafny.Array
				_ = _1614_v38
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_42
				var _nw236 _dafny.Array
				_ = _nw236
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw236 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Int = (func(_1615_v13 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1616_i5 _dafny.Int) _dafny.Int {
							return (_1616_i5).Minus(_1615_v13)
						}
					})(_1588_v13)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw236 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw236).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw236).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1614_v38 = _nw236
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1614_v38), 0))
				_ = _index249
				(_1614_v38).ArraySet1(_dafny.IntOfInt64(940), (_index249).Int())
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1614_v38), 0))
				_ = _index250
				(_1614_v38).ArraySet1((_dafny.IntOfInt64(607)).Minus(_1588_v13), (_index250).Int())
				(globalState).F8 = (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool)
				var _1617_v39 _dafny.Map
				_ = _1617_v39
				_1617_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(696), (_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))
				var _1618_v40 _dafny.Map
				_ = _1618_v40
				_1618_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), Companion_Default___.Fm13(_1588_v13, globalState))
				var _rhs249 bool = !(((_this).F10()) || ((_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))) || ((_this).Fm9((_1614_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1614_v38), 0))).Int()).(_dafny.Int), _1588_v13, _1588_v13, _1618_v40, globalState))
				_ = _rhs249
				var _rhs250 _dafny.Map = (func() _dafny.Map {
					if (_this).F10() {
						return _1617_v39
					}
					return _1617_v39
				})()
				_ = _rhs250
				var _lhs173 *GlobalState = globalState
				_ = _lhs173
				_lhs173.F8 = _rhs249
				_1617_v39 = _rhs250
			}
			(globalState).F8 = !(!((_1586_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1586_v12), 0))).Int()).(bool))) || ((_this).F10())
		}
		var _1619_v41 _dafny.Sequence
		_ = _1619_v41
		_1619_v41 = _dafny.UnicodeSeqOfUtf8Bytes("fuy")
		(globalState).F8 = _dafny.Companion_Sequence_.IsPrefixOf(_1619_v41, _dafny.UnicodeSeqOfUtf8Bytes("d"))
		var _1620_v42 _dafny.Int
		_ = _1620_v42
		_1620_v42 = _dafny.IntOfInt64(-36)
		var _1621_v43 _dafny.MultiSet
		_ = _1621_v43
		_1621_v43 = _dafny.MultiSetOf(_1620_v42, _dafny.IntOfInt64(94))
		var _1622_v44 _dafny.MultiSet
		_ = _1622_v44
		_1622_v44 = _dafny.MultiSetOf(_1621_v43, _1621_v43)
		var _1623_i6 _dafny.Int
		_ = _1623_i6
		_1623_i6 = _dafny.Zero
		{
			for (_1622_v44).IsSubsetOf(_1622_v44) {
				{
					if (_1623_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1623_i6 = (_1623_i6).Plus(_dafny.One)
					var _1624_v45 *C0
					_ = _1624_v45
					var _nw237 *C0 = New_C0_()
					_ = _nw237
					_nw237.Ctor__()
					_1624_v45 = _nw237
					_1620_v42 = _1620_v42
					var _1625_v46 _dafny.MultiSet
					_ = _1625_v46
					_1625_v46 = _dafny.MultiSetOf((_this).F10(), !(false), (_this).F10(), (_this).F10(), (_this).F10())
					var _1626_v47 _dafny.Set
					_ = _1626_v47
					_1626_v47 = _dafny.SetOf((_this).F10())
					var _1627_v48 _dafny.Sequence
					_ = _1627_v48
					_1627_v48 = _dafny.SeqOf((_1626_v47).Cardinality())
					var _1628_v49 _dafny.Map
					_ = _1628_v49
					_1628_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1620_v42, (_1627_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.IntOfUint32((_1627_v48).Cardinality()))).Uint32()).(_dafny.Int))
					var _1629_v50 _dafny.Map
					_ = _1629_v50
					_1629_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1625_v46).Contains((_this).F10()), (_dafny.IntOfUint32((_1619_v41).Cardinality())).Times((func() _dafny.Int {
						if (_1628_v49).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_1620_v42, _dafny.IntOfInt64(296))).Cardinality())) {
							return (_1628_v49).Get(_dafny.IntOfUint32((_dafny.SeqOf(_1620_v42, _dafny.IntOfInt64(296))).Cardinality())).(_dafny.Int)
						}
						return _1620_v42
					})()))
					var _1630_v51 _dafny.Set
					_ = _1630_v51
					_1630_v51 = _dafny.SetOf(_dafny.IntOfInt64(863))
					_1620_v42 = (func() _dafny.Int {
						if (_1629_v50).Contains((_this).F10()) {
							return (_1629_v50).Get((_this).F10()).(_dafny.Int)
						}
						return (_1620_v42).Times((_1630_v51).Cardinality())
					})()
					var _1631_v52 _dafny.Map
					_ = _1631_v52
					_1631_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1619_v41)
					var _1632_v53 _dafny.Array
					_ = _1632_v53
					var _nwElement0_49 bool = (_this).F10()
					_ = _nwElement0_49
					var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(25))
					_ = _nw238
					(_nw238).ArraySet1(_nwElement0_49, 0)
					(_nw238).ArraySet1((_this).F10(), 1)
					(_nw238).ArraySet1((_this).F10(), 2)
					(_nw238).ArraySet1(true, 3)
					(_nw238).ArraySet1((_this).F10(), 4)
					(_nw238).ArraySet1((_this).F10(), 5)
					(_nw238).ArraySet1((_this).F10(), 6)
					(_nw238).ArraySet1(!((_this).F10()), 7)
					(_nw238).ArraySet1((_this).F10(), 8)
					(_nw238).ArraySet1((_this).F10(), 9)
					(_nw238).ArraySet1(true, 10)
					(_nw238).ArraySet1((_this).F10(), 11)
					(_nw238).ArraySet1((_this).F10(), 12)
					(_nw238).ArraySet1((_this).F10(), 13)
					(_nw238).ArraySet1((_this).F10(), 14)
					(_nw238).ArraySet1((_this).F10(), 15)
					(_nw238).ArraySet1((_this).F10(), 16)
					(_nw238).ArraySet1(!((_this).F10()), 17)
					(_nw238).ArraySet1((_this).F10(), 18)
					(_nw238).ArraySet1((_this).Fm9((_this).Fm1((_this).F9(), globalState), _1620_v42, _1620_v42, _1631_v52, globalState), 19)
					(_nw238).ArraySet1((_this).F10(), 20)
					(_nw238).ArraySet1((_this).F10(), 21)
					(_nw238).ArraySet1(!((_this).F10()), 22)
					(_nw238).ArraySet1((_this).F10(), 23)
					(_nw238).ArraySet1((_this).F10(), 24)
					_1632_v53 = _nw238
					var _1633_v54 _dafny.Sequence
					_ = _1633_v54
					_1633_v54 = _dafny.SeqOf(_1632_v53, _1632_v53)
					var _1634_v55 _dafny.MultiSet
					_ = _1634_v55
					_1634_v55 = _dafny.MultiSetOf((_1633_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1627_v48).Cardinality()), _dafny.IntOfUint32((_1633_v54).Cardinality()))).Uint32()).(_dafny.Array), (func() _dafny.Array {
						if (_this).F10() {
							return _1632_v53
						}
						return (_1633_v54).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1620_v42, _1628_v49)).Cardinality(), _dafny.IntOfUint32((_1633_v54).Cardinality()))).Uint32()).(_dafny.Array)
					})(), _1632_v53, _1632_v53, _1632_v53)
					var _rhs251 _dafny.MultiSet = _1634_v55
					_ = _rhs251
					var _rhs252 bool = (_1620_v42).Cmp(_1620_v42) != 0
					_ = _rhs252
					var _lhs174 *GlobalState = globalState
					_ = _lhs174
					_1634_v55 = _rhs251
					_lhs174.F8 = _rhs252
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1635_v56 _dafny.Array
		_ = _1635_v56
		var _nw239 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw239
		_1635_v56 = _nw239
		for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1635_v56), 0))); ; {
			_guard_loop_6, _ok65 := _iter65()
			if !_ok65 {
				break
			}
			var _1636_i7 _dafny.Int
			_1636_i7 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1636_i7).Sign() != -1) && ((_1636_i7).Cmp(_dafny.ArrayLen((_1635_v56), 0)) < 0)) {
				(_1635_v56).ArraySet1((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_this).F10())), _1620_v42)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F10()), _1620_v42))).Cardinality()).Cmp(_dafny.IntOfInt64(939)) <= 0, (_1636_i7).Int())
			}
		}
		var _1637_v57 _dafny.Set
		_ = _1637_v57
		_1637_v57 = _dafny.SetOf(_1620_v42)
		var _1638_v58 _dafny.MultiSet
		_ = _1638_v58
		_1638_v58 = _dafny.MultiSetOf(!((_this).F10()), (_this).F10())
		var _1639_v59 _dafny.Sequence
		_ = _1639_v59
		_1639_v59 = _dafny.SeqOf(_1638_v58, _1638_v58, _1638_v58)
		var _1640_v60 _dafny.Map
		_ = _1640_v60
		_1640_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1637_v57, _1639_v59)
		_1640_v60 = (_1640_v60).Update((func() _dafny.Set {
			if (_this).F10() {
				return _1637_v57
			}
			return func() _dafny.Set {
				var _coll59 = _dafny.NewBuilder()
				_ = _coll59
				for _iter66 := _dafny.Iterate((_dafny.SeqOf(_1620_v42)).Elements()); ; {
					_compr_59, _ok66 := _iter66()
					if !_ok66 {
						break
					}
					var _1641_v61 _dafny.Int
					_1641_v61 = interface{}(_compr_59).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1620_v42), _1641_v61) {
						_coll59.Add(Companion_Default___.SafeDivisionInt(_1641_v61, _dafny.IntOfUint32((_dafny.SeqOf(true, true, false, false)).Cardinality())))
					}
				}
				return _coll59.ToSet()
			}()
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1638_v58), _1639_v59))
		var _1642_v62 D0
		_ = _1642_v62
		_1642_v62 = Companion_D0_.Create_DC1_(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(765), _1620_v42), (_this).F10(), _1620_v42)
		var _source22 D0 = _1642_v62
		_ = _source22
		if _source22.Is_DC1() {
			var _1643___mcc_h0 _dafny.Int = _source22.Get_().(D0_DC1).Cf1
			_ = _1643___mcc_h0
			var _1644___mcc_h1 bool = _source22.Get_().(D0_DC1).Cf2
			_ = _1644___mcc_h1
			var _1645___mcc_h2 _dafny.Int = _source22.Get_().(D0_DC1).Cf3
			_ = _1645___mcc_h2
			var _1646_cf3 _dafny.Int = _1645___mcc_h2
			_ = _1646_cf3
			var _1647_cf2 bool = _1644___mcc_h1
			_ = _1647_cf2
			var _1648_cf1 _dafny.Int = _1643___mcc_h0
			_ = _1648_cf1
			var _1649_v63 _dafny.Array
			_ = _1649_v63
			var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw240
			_1649_v63 = _nw240
			var _1650_v64 D1
			_ = _1650_v64
			_1650_v64 = Companion_D1_.Create_DC4_(false, _1647_cf2, _dafny.IntOfUint32((_1619_v41).Cardinality()))
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_1649_v63), 0))
			_ = _index251
			(_1649_v63).ArraySet1((_1650_v64).Dtor_cf8(), (_index251).Int())
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_1649_v63), 0))
			_ = _index252
			var _rhs253 _dafny.Int = _1646_cf3
			_ = _rhs253
			var _rhs254 bool = (_this).F10()
			_ = _rhs254
			var _rhs255 _dafny.Int = _1620_v42
			_ = _rhs255
			var _lhs175 _dafny.Array = _1649_v63
			_ = _lhs175
			var _lhs176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_1649_v63), 0))
			_ = _lhs176
			_1620_v42 = _rhs253
			_1647_cf2 = _rhs254
			(_lhs175).ArraySet1(_rhs255, (_lhs176).Int())
			var _1651_v65 _dafny.Sequence
			_ = _1651_v65
			_1651_v65 = _dafny.SeqOf(_1619_v41)
			var _1652_v66 D1
			_ = _1652_v66
			_1652_v66 = Companion_D1_.Create_DC3_((_1651_v65).Select((Companion_Default___.SafeIndex(_1620_v42, _dafny.IntOfUint32((_1651_v65).Cardinality()))).Uint32()).(_dafny.Sequence))
			_1652_v66 = _1652_v66
			_1635_v56 = _1635_v56
			_1647_cf2 = _1647_cf2
		} else {
			var _1653___mcc_h3 bool = _source22.Get_().(D0_DC0).Cf0
			_ = _1653___mcc_h3
			var _1654_cf0 bool = _1653___mcc_h3
			_ = _1654_cf0
			var _1655_v67 _dafny.Array
			_ = _1655_v67
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_43
			var _nw241 _dafny.Array
			_ = _nw241
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw241 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.MultiSet = (func(_1656_v58 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_1657_i8 _dafny.Int) _dafny.MultiSet {
						return (_1656_v58).Intersection(_1656_v58)
					}
				})(_1638_v58)
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw241 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw241).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw241).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1655_v67 = _nw241
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_1655_v67), 0))
			_ = _index253
			(_1655_v67).ArraySet1((_1638_v58).Union(_1638_v58), (_index253).Int())
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_1655_v67), 0))
			_ = _index254
			(_1655_v67).ArraySet1(Companion_Default___.Fm15((_this).F10(), (_this).F10(), (_1620_v42).Times((func() _dafny.Map {
				var _coll60 = _dafny.NewMapBuilder()
				_ = _coll60
				for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-192), _dafny.IntOfInt64(379))); ; {
					_compr_60, _ok67 := _iter67()
					if !_ok67 {
						break
					}
					var _1658_v68 _dafny.Int
					_1658_v68 = interface{}(_compr_60).(_dafny.Int)
					if ((_dafny.IntOfInt64(-192)).Cmp(_1658_v68) <= 0) && ((_1658_v68).Cmp(_dafny.IntOfInt64(379)) < 0) {
						_coll60.Add((_1658_v68).Plus(_1620_v42), _1654_cf0)
					}
				}
				return _coll60.ToMap()
			}()).Cardinality()), _dafny.CodePoint('k'), globalState), (_index254).Int())
			var _1659_v69 *C0
			_ = _1659_v69
			var _nw242 *C0 = New_C0_()
			_ = _nw242
			_nw242.Ctor__()
			_1659_v69 = _nw242
			var _1660_v70 _dafny.Array
			_ = _1660_v70
			var _len0_44 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_44
			var _nw243 _dafny.Array
			_ = _nw243
			if _len0_44.Cmp(_dafny.Zero) == 0 {
				_nw243 = _dafny.NewArray(_len0_44)
			} else {
				var _init44 func(_dafny.Int) D0 = (func(_1661_cf0 bool, _1662_v62 D0) func(_dafny.Int) D0 {
					return func(_1663_i9 _dafny.Int) D0 {
						return (func() D0 {
							if _1661_cf0 {
								return _1662_v62
							}
							return _1662_v62
						})()
					}
				})(_1654_cf0, _1642_v62)
				_ = _init44
				var _element0_44 = _init44(_dafny.Zero)
				_ = _element0_44
				_nw243 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
				(_nw243).ArraySet1(_element0_44, 0)
				var _nativeLen0_44 = (_len0_44).Int()
				_ = _nativeLen0_44
				for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
					(_nw243).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
				}
			}
			_1660_v70 = _nw243
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1660_v70), 0))
			_ = _index255
			(_1660_v70).ArraySet1(_1642_v62, (_index255).Int())
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1660_v70), 0))
			_ = _index256
			(_1660_v70).ArraySet1(_1642_v62, (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1635_v56), 0))
			_ = _index257
			(_1635_v56).ArraySet1((_this).F10(), (_index257).Int())
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1635_v56), 0))
			_ = _index258
			(_1635_v56).ArraySet1(true, (_index258).Int())
		}
		r0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1619_v41, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}(func(_1664_i10 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))), (Companion_Default___.SafeIndex((func() _dafny.Int {
			if true {
				return _1620_v42
			}
			return (_dafny.Zero).Minus(_1620_v42)
		})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1619_v41, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg95 _dafny.Int) interface{} {
				return coer95(arg95)
			}
		}(func(_1665_i10 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})))).Cardinality()))).Uint32(), (_1619_v41).Select((Companion_Default___.SafeIndex(_1620_v42, _dafny.IntOfUint32((_1619_v41).Cardinality()))).Uint32()).(_dafny.CodePoint))
		r1 = (_this).F10()
		var _1666_v71 _dafny.Set
		_ = _1666_v71
		_1666_v71 = _dafny.SetOf((_this).F10())
		var _1667_v72 _dafny.Sequence
		_ = _1667_v72
		_1667_v72 = _dafny.SeqOf(_1666_v71)
		var _1668_v73 _dafny.Map
		_ = _1668_v73
		_1668_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1667_v72)
		var _1669_v74 _dafny.Sequence
		_ = _1669_v74
		_1669_v74 = _dafny.SeqOf(_1620_v42)
		r2 = (func() _dafny.Sequence {
			if (_1668_v73).Contains((_this).F10()) {
				return (_1668_v73).Get((_this).F10()).(_dafny.Sequence)
			}
			return Companion_Default___.Fm16(_dafny.IntOfUint32((_1669_v74).Cardinality()), _1620_v42, _dafny.IntOfInt64(913), globalState)
		})()
		r3 = _1619_v41
		return r0, r1, r2, r3
	}
}
func (_this *C6) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _1670_v0 *C0
		_ = _1670_v0
		var _nw244 *C0 = New_C0_()
		_ = _nw244
		_nw244.Ctor__()
		_1670_v0 = _nw244
		var _1671_v1 _dafny.Int
		_ = _1671_v1
		_1671_v1 = _dafny.IntOfInt64(699)
		var _1672_v2 D1
		_ = _1672_v2
		_1672_v2 = Companion_D1_.Create_DC4_((_this).F10(), (_this).F10(), _1671_v1)
		r1 = func(_source23 D1) bool {
			if _source23.Is_DC3() {
				var _1673___mcc_h0 _dafny.Sequence = _source23.Get_().(D1_DC3).Cf5
				_ = _1673___mcc_h0
				var _1674_cf5 _dafny.Sequence = _1673___mcc_h0
				_ = _1674_cf5
				return (_this).F10()
			} else if _source23.Is_DC4() {
				var _1675___mcc_h1 bool = _source23.Get_().(D1_DC4).Cf6
				_ = _1675___mcc_h1
				var _1676___mcc_h2 bool = _source23.Get_().(D1_DC4).Cf7
				_ = _1676___mcc_h2
				var _1677___mcc_h3 _dafny.Int = _source23.Get_().(D1_DC4).Cf8
				_ = _1677___mcc_h3
				var _1678_cf8 _dafny.Int = _1677___mcc_h3
				_ = _1678_cf8
				var _1679_cf7 bool = _1676___mcc_h2
				_ = _1679_cf7
				var _1680_cf6 bool = _1675___mcc_h1
				_ = _1680_cf6
				return (_dafny.MultiSetOf(_1680_cf6, (_this).F10(), (_this).F10(), true)).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)))
			} else if _source23.Is_DC2() {
				var _1681___mcc_h4 _dafny.Sequence = _source23.Get_().(D1_DC2).Cf4
				_ = _1681___mcc_h4
				var _1682_cf4 _dafny.Sequence = _1681___mcc_h4
				_ = _1682_cf4
				return (_this).F10()
			} else {
				var _1683___mcc_h5 D1 = _source23.Get_().(D1_DC5).Cf9
				_ = _1683___mcc_h5
				var _1684_cf9 D1 = _1683___mcc_h5
				_ = _1684_cf9
				return (_this).F10()
			}
		}(Companion_D1_.Create_DC5_(_1672_v2))
		var _1685_v3 _dafny.MultiSet
		_ = _1685_v3
		_1685_v3 = _dafny.MultiSetOf(_1671_v1, _1671_v1)
		var _1686_v4 _dafny.Sequence
		_ = _1686_v4
		_1686_v4 = _dafny.SeqOf(_1671_v1)
		var _1687_v5 D1
		_ = _1687_v5
		_1687_v5 = Companion_D1_.Create_DC4_((_this).F10(), (_this).F10(), _1671_v1)
		var _1688_v6 _dafny.Sequence
		_ = _1688_v6
		_1688_v6 = _dafny.UnicodeSeqOfUtf8Bytes("qfpc")
		var _1689_v7 _dafny.Set
		_ = _1689_v7
		_1689_v7 = _dafny.SetOf((_this).F10())
		var _1690_v8 _dafny.Map
		_ = _1690_v8
		_1690_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v1, (_this).F10())
		var _1691_v9 _dafny.Array
		_ = _1691_v9
		var _nwElement0_50 _dafny.Int = _dafny.IntOfInt64(-961)
		_ = _nwElement0_50
		var _nw245 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(19))
		_ = _nw245
		(_nw245).ArraySet1(_nwElement0_50, 0)
		(_nw245).ArraySet1(_1671_v1, 1)
		(_nw245).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(810), _1671_v1), 2)
		(_nw245).ArraySet1(((_1685_v3).Difference(_1685_v3)).Cardinality(), 3)
		(_nw245).ArraySet1(_dafny.IntOfInt64(139), 4)
		(_nw245).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1686_v4, _1686_v4)).Cardinality()), 5)
		(_nw245).ArraySet1((_dafny.IntOfInt64(-855)).Plus((_1687_v5).Dtor_cf8()), 6)
		(_nw245).ArraySet1(_dafny.IntOfUint32((_1688_v6).Cardinality()), 7)
		(_nw245).ArraySet1(_1671_v1, 8)
		(_nw245).ArraySet1((_1689_v7).Cardinality(), 9)
		(_nw245).ArraySet1((_dafny.IntOfInt64(-393)).Minus(_1671_v1), 10)
		(_nw245).ArraySet1(_1671_v1, 11)
		(_nw245).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1671_v1, _1671_v1)), 12)
		(_nw245).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(869), _1671_v1), 13)
		(_nw245).ArraySet1(_1671_v1, 14)
		(_nw245).ArraySet1(_1671_v1, 15)
		(_nw245).ArraySet1(_1671_v1, 16)
		(_nw245).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1688_v6, (Companion_Default___.SafeIndex((_1690_v8).Cardinality(), _dafny.IntOfUint32((_1688_v6).Cardinality()))).Uint32(), p0), _1688_v6)).Cardinality()), 17)
		(_nw245).ArraySet1(_1671_v1, 18)
		_1691_v9 = _nw245
		for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1691_v9), 0))); ; {
			_guard_loop_7, _ok68 := _iter68()
			if !_ok68 {
				break
			}
			var _1692_i0 _dafny.Int
			_1692_i0 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_1692_i0).Sign() != -1) && ((_1692_i0).Cmp(_dafny.ArrayLen((_1691_v9), 0)) < 0)) {
				(_1691_v9).ArraySet1((_1692_i0).Plus(_1671_v1), (_1692_i0).Int())
			}
		}
		var _hi4 _dafny.Int = (_dafny.Zero).Minus(_1671_v1)
		_ = _hi4
		for _1693_i1 := (_1671_v1).Minus(_1671_v1); _1693_i1.Cmp(_hi4) < 0; _1693_i1 = _1693_i1.Plus(_dafny.One) {
			(globalState).F8 = true
			var _1694_v10 *C0
			_ = _1694_v10
			var _nw246 *C0 = New_C0_()
			_ = _nw246
			_nw246.Ctor__()
			_1694_v10 = _nw246
			var _1695_v11 _dafny.Map
			_ = _1695_v11
			_1695_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1691_v9, _1688_v6)
			var _1696_v12 _dafny.Array
			_ = _1696_v12
			var _nwElement0_51 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1688_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(877))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg96 _dafny.Int) interface{} {
					return coer96(arg96)
				}
			}(func(_1697_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})))
			_ = _nwElement0_51
			var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(25))
			_ = _nw247
			(_nw247).ArraySet1(_nwElement0_51, 0)
			(_nw247).ArraySet1(_1688_v6, 1)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("amb"), _1688_v6), 2)
			(_nw247).ArraySet1(_1688_v6, 3)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xmfr"), _dafny.Companion_Sequence_.Update(_1688_v6, (Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1688_v6).Cardinality()))).Uint32(), p0)), 4)
			(_nw247).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bk"), 5)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _dafny.UnicodeSeqOfUtf8Bytes("bac")), (Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1688_v6, _dafny.UnicodeSeqOfUtf8Bytes("bac"))).Cardinality()))).Uint32(), p0), 6)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _1688_v6), 7)
			(_nw247).ArraySet1(_1688_v6, 8)
			(_nw247).ArraySet1(_1688_v6, 9)
			(_nw247).ArraySet1(_1688_v6, 10)
			(_nw247).ArraySet1(_1688_v6, 11)
			(_nw247).ArraySet1(_1688_v6, 12)
			(_nw247).ArraySet1(_1688_v6, 13)
			(_nw247).ArraySet1(_1688_v6, 14)
			(_nw247).ArraySet1((func() _dafny.Sequence {
				if (_1695_v11).Contains(_1691_v9) {
					return (_1695_v11).Get(_1691_v9).(_dafny.Sequence)
				}
				return _1688_v6
			})(), 15)
			(_nw247).ArraySet1(Companion_Default___.Fm13(_dafny.IntOfInt64(592), globalState), 16)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _1688_v6), 17)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _1688_v6), 18)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg97 _dafny.Int) interface{} {
					return coer97(arg97)
				}
			}((func(_1698_v6 _dafny.Sequence, _1699_v1 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_1700_i3 _dafny.Int) _dafny.CodePoint {
					return (_1698_v6).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_1699_v1, _1699_v1)).Cardinality(), _dafny.IntOfUint32((_1698_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_1688_v6, _1671_v1)))), 19)
			(_nw247).ArraySet1(_1688_v6, 20)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _1688_v6), 21)
			(_nw247).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1688_v6, _dafny.UnicodeSeqOfUtf8Bytes("om")), 22)
			(_nw247).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lsvc"), 23)
			(_nw247).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ldunu"), 24)
			_1696_v12 = _nw247
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1696_v12), 0))
			_ = _index259
			(_1696_v12).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg98 _dafny.Int) interface{} {
					return coer98(arg98)
				}
			}((func(_1701_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1702_i4 _dafny.Int) _dafny.CodePoint {
					return _1701_p0
				}
			})(p0))), (_index259).Int())
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1691_v9), 0))
			_ = _index260
			(_1691_v9).ArraySet1(_1693_i1, (_index260).Int())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1696_v12), 0))
			_ = _index261
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1691_v9), 0))
			_ = _index262
			var _rhs256 _dafny.Int = _1671_v1
			_ = _rhs256
			var _rhs257 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1688_v6, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gpkthj"), _1688_v6), (Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gpkthj"), _1688_v6)).Cardinality()))).Uint32(), _dafny.CodePoint('i')))
			_ = _rhs257
			var _rhs258 bool = (_this).F10()
			_ = _rhs258
			var _rhs259 _dafny.Int = (_dafny.IntOfInt64(-553)).Times((_1693_i1).Times(_dafny.IntOfInt64(583)))
			_ = _rhs259
			var _lhs177 _dafny.Array = _1696_v12
			_ = _lhs177
			var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1696_v12), 0))
			_ = _lhs178
			var _lhs179 _dafny.Array = _1691_v9
			_ = _lhs179
			var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_1691_v9), 0))
			_ = _lhs180
			_1671_v1 = _rhs256
			(_lhs177).ArraySet1(_rhs257, (_lhs178).Int())
			r1 = _rhs258
			(_lhs179).ArraySet1(_rhs259, (_lhs180).Int())
			var _1703_v13 *C0
			_ = _1703_v13
			var _nw248 *C0 = New_C0_()
			_ = _nw248
			_nw248.Ctor__()
			_1703_v13 = _nw248
		}
		var _1704_v14 _dafny.CodePoint
		_ = _1704_v14
		_1704_v14 = _dafny.CodePoint('j')
		_1704_v14 = _1704_v14
		var _1705_v15 D1
		_ = _1705_v15
		_1705_v15 = Companion_D1_.Create_DC2_(_dafny.Companion_Sequence_.Update(_1686_v4, (Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1686_v4).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_1671_v1)))
		var _source24 D1 = _1705_v15
		_ = _source24
		if _source24.Is_DC3() {
			var _1706___mcc_h6 _dafny.Sequence = _source24.Get_().(D1_DC3).Cf5
			_ = _1706___mcc_h6
			var _1707_cf5 _dafny.Sequence = _1706___mcc_h6
			_ = _1707_cf5
			var _1708_v16 _dafny.Map
			_ = _1708_v16
			_1708_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_this).F10())
			var _1709_v17 _dafny.Sequence
			_ = _1709_v17
			_1709_v17 = _dafny.SeqOf((_this).F10())
			_1686_v4 = (func() _dafny.Sequence {
				if (func() bool {
					if ((_this).F9()).Contains((_this).F10()) {
						return ((_this).F9()).Get((_this).F10()).(bool)
					}
					return !((func() bool {
						if (_1708_v16).Contains(_1704_v14) {
							return (_1708_v16).Get(_1704_v14).(bool)
						}
						return (_1709_v17).Select((Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1709_v17).Cardinality()))).Uint32()).(bool)
					})())
				})() {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(149))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg99 _dafny.Int) interface{} {
							return coer99(arg99)
						}
					}((func(_1710_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1711_i5 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_1710_v1)
						}
					})(_1671_v1))), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F10(), (_this).F10())).Cardinality())))
				}
				return _1686_v4
			})()
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_1691_v9), 0))
			_ = _index263
			(_1691_v9).ArraySet1(_1671_v1, (_index263).Int())
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_1691_v9), 0))
			_ = _index264
			(_1691_v9).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1671_v1)), (_index264).Int())
			_1690_v8 = _1690_v8
			var _1712_v18 _dafny.Array
			_ = _1712_v18
			var _nwElement0_52 bool = (_this).F10()
			_ = _nwElement0_52
			var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(4))
			_ = _nw249
			(_nw249).ArraySet1(_nwElement0_52, 0)
			(_nw249).ArraySet1((_this).F10(), 1)
			(_nw249).ArraySet1((func() bool {
				if (_this).F10() {
					return true
				}
				return (_this).F10()
			})(), 2)
			(_nw249).ArraySet1((_dafny.IntOfInt64(-255)).Cmp(_dafny.IntOfInt64(528)) <= 0, 3)
			_1712_v18 = _nw249
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_1712_v18), 0))
			_ = _index265
			(_1712_v18).ArraySet1(true, (_index265).Int())
			var _1713_v19 _dafny.Set
			_ = _1713_v19
			_1713_v19 = _dafny.SetOf(_1671_v1)
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_1712_v18), 0))
			_ = _index266
			(_1712_v18).ArraySet1((Companion_Default___.Fm17(true, globalState)).IsSubsetOf((_1713_v19).Difference(_1713_v19)), (_index266).Int())
		} else if _source24.Is_DC4() {
			var _1714___mcc_h7 bool = _source24.Get_().(D1_DC4).Cf6
			_ = _1714___mcc_h7
			var _1715___mcc_h8 bool = _source24.Get_().(D1_DC4).Cf7
			_ = _1715___mcc_h8
			var _1716___mcc_h9 _dafny.Int = _source24.Get_().(D1_DC4).Cf8
			_ = _1716___mcc_h9
			var _1717_cf8 _dafny.Int = _1716___mcc_h9
			_ = _1717_cf8
			var _1718_cf7 bool = _1715___mcc_h8
			_ = _1718_cf7
			var _1719_cf6 bool = _1714___mcc_h7
			_ = _1719_cf6
			var _1720_v20 D1
			_ = _1720_v20
			_1720_v20 = Companion_D1_.Create_DC5_(_1672_v2)
			var _1721_v21 _dafny.Map
			_ = _1721_v21
			_1721_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1720_v20, Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_1671_v1), _1718_cf7, _1671_v1))
			var _1722_v22 D0
			_ = _1722_v22
			_1722_v22 = Companion_D0_.Create_DC0_(_1718_cf7)
			var _1723_v23 _dafny.Sequence
			_ = _1723_v23
			_1723_v23 = _dafny.SeqOf(_1722_v22, _1722_v22)
			var _1724_v24 _dafny.Map
			_ = _1724_v24
			_1724_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1718_cf7, _dafny.IntOfInt64(-583))
			var _1725_v25 _dafny.Sequence
			_ = _1725_v25
			_1725_v25 = _dafny.SeqOf((_this).F10())
			var _rhs260 bool = (_1671_v1).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1723_v23, _1723_v23)).Cardinality())) > 0
			_ = _rhs260
			var _rhs261 _dafny.Map = _1721_v21
			_ = _rhs261
			var _rhs262 _dafny.Int = (func() _dafny.Int {
				if _1719_cf6 {
					return _1717_cf8
				}
				return (func() _dafny.Int {
					if (_1724_v24).Contains((_1725_v25).Select((Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1725_v25).Cardinality()))).Uint32()).(bool)) {
						return (_1724_v24).Get((_1725_v25).Select((Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1725_v25).Cardinality()))).Uint32()).(bool)).(_dafny.Int)
					}
					return (_1689_v7).Cardinality()
				})()
			})()
			_ = _rhs262
			_1718_cf7 = _rhs260
			_1721_v21 = _rhs261
			_1717_cf8 = _rhs262
			var _1726_v26 *C0
			_ = _1726_v26
			var _nw250 *C0 = New_C0_()
			_ = _nw250
			_nw250.Ctor__()
			_1726_v26 = _nw250
			var _1727_v27 _dafny.Sequence
			_ = _1727_v27
			_1727_v27 = _dafny.SeqOf(_1691_v9)
			_1717_cf8 = (_dafny.IntOfUint32((_1727_v27).Cardinality())).Plus(_1671_v1)
			var _1728_v28 _dafny.CodePoint
			_ = _1728_v28
			_1728_v28 = _1704_v14
			var _1729_v29 D1
			_ = _1729_v29
			_1729_v29 = Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Update(_1688_v6, (Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1688_v6).Cardinality()))).Uint32(), (_1728_v28)))
			var _source25 D1 = _1729_v29
			_ = _source25
			if _source25.Is_DC3() {
				var _1730___mcc_h12 _dafny.Sequence = _source25.Get_().(D1_DC3).Cf5
				_ = _1730___mcc_h12
				var _1731_cf5 _dafny.Sequence = _1730___mcc_h12
				_ = _1731_cf5
				var _1732_v30 *C0
				_ = _1732_v30
				var _nw251 *C0 = New_C0_()
				_ = _nw251
				_nw251.Ctor__()
				_1732_v30 = _nw251
				_1728_v28 = _1704_v14
				var _1733_v31 _dafny.Array
				_ = _1733_v31
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_45
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) bool = func(_1734_i6 _dafny.Int) bool {
						return (_this).F10()
					}
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw252 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw252).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw252).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1733_v31 = _nw252
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1733_v31), 0))
				_ = _index267
				(_1733_v31).ArraySet1(_1718_cf7, (_index267).Int())
				var _1735_v32 _dafny.Map
				_ = _1735_v32
				_1735_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1719_cf6, _1691_v9)
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1733_v31), 0))
				_ = _index268
				var _rhs263 bool = _1719_cf6
				_ = _rhs263
				var _rhs264 _dafny.Array = (func() _dafny.Array {
					if (_1735_v32).Contains(false) {
						return (_1735_v32).Get(false).(_dafny.Array)
					}
					return (_1727_v27).Select((Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1727_v27).Cardinality()))).Uint32()).(_dafny.Array)
				})()
				_ = _rhs264
				var _rhs265 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fnhqhubuo"), _1688_v6)).Cardinality())
				_ = _rhs265
				var _rhs266 _dafny.CodePoint = Companion_Default___.Fm14(globalState)
				_ = _rhs266
				var _lhs181 _dafny.Array = _1733_v31
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(138), _dafny.ArrayLen((_1733_v31), 0))
				_ = _lhs182
				(_lhs181).ArraySet1(_rhs263, (_lhs182).Int())
				_1691_v9 = _rhs264
				_1717_cf8 = _rhs265
				_1704_v14 = _rhs266
				_1690_v8 = _1690_v8
			} else if _source25.Is_DC4() {
				var _1736___mcc_h13 bool = _source25.Get_().(D1_DC4).Cf6
				_ = _1736___mcc_h13
				var _1737___mcc_h14 bool = _source25.Get_().(D1_DC4).Cf7
				_ = _1737___mcc_h14
				var _1738___mcc_h15 _dafny.Int = _source25.Get_().(D1_DC4).Cf8
				_ = _1738___mcc_h15
				var _1739_cf8 _dafny.Int = _1738___mcc_h15
				_ = _1739_cf8
				var _1740_cf7 bool = _1737___mcc_h14
				_ = _1740_cf7
				var _1741_cf6 bool = _1736___mcc_h13
				_ = _1741_cf6
				_1690_v8 = (_1690_v8).Update(Companion_Default___.SafeModuloInt(_1739_cf8, _dafny.IntOfInt64(472)), _1741_cf6)
				var _1742_v33 _dafny.Set
				_ = _1742_v33
				_1742_v33 = _dafny.SetOf((_dafny.SetOf(_1724_v24)).Cardinality())
				var _1743_v43 _dafny.Map
				_ = _1743_v43
				_1743_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1740_cf7, _1670_v0)
				var _1744_v44 _dafny.Sequence
				_ = _1744_v44
				_1744_v44 = _dafny.SeqOf(func() _dafny.Set {
					var _coll61 = _dafny.NewBuilder()
					_ = _coll61
					for _iter69 := _dafny.Iterate((_1686_v4).Elements()); ; {
						_compr_61, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _1745_v41 _dafny.Int
						_1745_v41 = interface{}(_compr_61).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1686_v4, _1745_v41) {
							_coll61.Add((_1745_v41).Minus(_dafny.IntOfInt64(754)))
						}
					}
					return _coll61.ToSet()
				}(), _1742_v33, func() _dafny.Set {
					var _coll62 = _dafny.NewBuilder()
					_ = _coll62
					for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(487), _dafny.IntOfInt64(259))); ; {
						_compr_62, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _1746_v42 _dafny.Int
						_1746_v42 = interface{}(_compr_62).(_dafny.Int)
						if ((_dafny.IntOfInt64(487)).Cmp(_1746_v42) <= 0) && ((_1746_v42).Cmp(_dafny.IntOfInt64(259)) < 0) {
							_coll62.Add((_1746_v42).Minus(_dafny.IntOfUint32((_1688_v6).Cardinality())))
						}
					}
					return _coll62.ToSet()
				}(), _1742_v33, _dafny.SetOf(_1739_cf8, _1739_cf8, _dafny.IntOfInt64(-918), _1671_v1, (_1743_v43).Cardinality()))
				var _1747_v45 _dafny.Array
				_ = _1747_v45
				var _nwElement0_53 _dafny.Set = _1742_v33
				_ = _nwElement0_53
				var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(13))
				_ = _nw253
				(_nw253).ArraySet1(_nwElement0_53, 0)
				(_nw253).ArraySet1(_1742_v33, 1)
				(_nw253).ArraySet1(_1742_v33, 2)
				(_nw253).ArraySet1(func() _dafny.Set {
					var _coll63 = _dafny.NewBuilder()
					_ = _coll63
					for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(600), _dafny.IntOfInt64(-636))); ; {
						_compr_63, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _1748_v34 _dafny.Int
						_1748_v34 = interface{}(_compr_63).(_dafny.Int)
						if ((_dafny.IntOfInt64(600)).Cmp(_1748_v34) <= 0) && ((_1748_v34).Cmp(_dafny.IntOfInt64(-636)) < 0) {
							_coll63.Add(Companion_Default___.SafeModuloInt(_1748_v34, (func() _dafny.Map {
								var _coll64 = _dafny.NewMapBuilder()
								_ = _coll64
								for _iter72 := _dafny.Iterate((func() _dafny.Map {
									var _coll65 = _dafny.NewMapBuilder()
									_ = _coll65
									for _iter73 := _dafny.Iterate((_dafny.MultiSetOf(_1688_v6)).Elements()); ; {
										_compr_65, _ok73 := _iter73()
										if !_ok73 {
											break
										}
										var _1749_v36 _dafny.Sequence
										_1749_v36 = interface{}(_compr_65).(_dafny.Sequence)
										if (_dafny.MultiSetOf(_1688_v6)).Contains(_1749_v36) {
											_coll65.Add(_1749_v36, _dafny.IntOfUint32((_1688_v6).Cardinality()))
										}
									}
									return _coll65.ToMap()
								}()).Keys().Elements()); ; {
									_compr_64, _ok72 := _iter72()
									if !_ok72 {
										break
									}
									var _1750_v35 _dafny.Sequence
									_1750_v35 = interface{}(_compr_64).(_dafny.Sequence)
									if (func() _dafny.Map {
										var _coll66 = _dafny.NewMapBuilder()
										_ = _coll66
										for _iter74 := _dafny.Iterate((_dafny.MultiSetOf(_1688_v6)).Elements()); ; {
											_compr_66, _ok74 := _iter74()
											if !_ok74 {
												break
											}
											var _1749_v36 _dafny.Sequence
											_1749_v36 = interface{}(_compr_66).(_dafny.Sequence)
											if (_dafny.MultiSetOf(_1688_v6)).Contains(_1749_v36) {
												_coll66.Add(_1749_v36, _dafny.IntOfUint32((_1688_v6).Cardinality()))
											}
										}
										return _coll66.ToMap()
									}()).Contains(_1750_v35) {
										_coll64.Add(_1750_v35, _1741_cf6)
									}
								}
								return _coll64.ToMap()
							}()).Cardinality()))
						}
					}
					return _coll63.ToSet()
				}(), 3)
				(_nw253).ArraySet1(func() _dafny.Set {
					var _coll67 = _dafny.NewBuilder()
					_ = _coll67
					for _iter75 := _dafny.Iterate((_1685_v3).Elements()); ; {
						_compr_67, _ok75 := _iter75()
						if !_ok75 {
							break
						}
						var _1751_v37 _dafny.Int
						_1751_v37 = interface{}(_compr_67).(_dafny.Int)
						if (_1685_v3).Contains(_1751_v37) {
							_coll67.Add((_1751_v37).Times(_dafny.IntOfInt64(485)))
						}
					}
					return _coll67.ToSet()
				}(), 4)
				(_nw253).ArraySet1((func() _dafny.Set {
					if _1740_cf7 {
						return _1742_v33
					}
					return _1742_v33
				})(), 5)
				(_nw253).ArraySet1(_1742_v33, 6)
				(_nw253).ArraySet1(_1742_v33, 7)
				(_nw253).ArraySet1(_1742_v33, 8)
				(_nw253).ArraySet1(func() _dafny.Set {
					var _coll68 = _dafny.NewBuilder()
					_ = _coll68
					for _iter76 := _dafny.Iterate((_1686_v4).Elements()); ; {
						_compr_68, _ok76 := _iter76()
						if !_ok76 {
							break
						}
						var _1752_v38 _dafny.Int
						_1752_v38 = interface{}(_compr_68).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1686_v4, _1752_v38) {
							_coll68.Add(Companion_Default___.SafeDivisionInt(_1752_v38, (func() _dafny.Map {
								var _coll69 = _dafny.NewMapBuilder()
								_ = _coll69
								for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(305), _dafny.IntOfInt64(277))); ; {
									_compr_69, _ok77 := _iter77()
									if !_ok77 {
										break
									}
									var _1753_v39 _dafny.Int
									_1753_v39 = interface{}(_compr_69).(_dafny.Int)
									if ((_dafny.IntOfInt64(305)).Cmp(_1753_v39) <= 0) && ((_1753_v39).Cmp(_dafny.IntOfInt64(277)) < 0) {
										_coll69.Add(Companion_Default___.SafeModuloInt(_1753_v39, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(330))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
											return func(arg100 _dafny.Int) interface{} {
												return coer100(arg100)
											}
										}(func(_1755_i7 _dafny.Int) _dafny.Int {
											return _dafny.IntOfInt64(809)
										}))).Cardinality()))), func() _dafny.Set {
											var _coll70 = _dafny.NewBuilder()
											_ = _coll70
											for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(390), _dafny.IntOfInt64(962))); ; {
												_compr_70, _ok78 := _iter78()
												if !_ok78 {
													break
												}
												var _1754_v40 _dafny.Int
												_1754_v40 = interface{}(_compr_70).(_dafny.Int)
												if ((_dafny.IntOfInt64(390)).Cmp(_1754_v40) <= 0) && ((_1754_v40).Cmp(_dafny.IntOfInt64(962)) < 0) {
													_coll70.Add((_1754_v40).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())))
												}
											}
											return _coll70.ToSet()
										}())
									}
								}
								return _coll69.ToMap()
							}()).Cardinality()))
						}
					}
					return _coll68.ToSet()
				}(), 9)
				(_nw253).ArraySet1((_1744_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.IntOfUint32((_1744_v44).Cardinality()))).Uint32()).(_dafny.Set), 10)
				(_nw253).ArraySet1(_1742_v33, 11)
				(_nw253).ArraySet1(Companion_Default___.Fm17((_this).F10(), globalState), 12)
				_1747_v45 = _nw253
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1747_v45), 0))
				_ = _index269
				(_1747_v45).ArraySet1(_1742_v33, (_index269).Int())
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1747_v45), 0))
				_ = _index270
				(_1747_v45).ArraySet1(_1742_v33, (_index270).Int())
				var _1756_v46 _dafny.Map
				_ = _1756_v46
				_1756_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-407), _1739_cf8)
				var _1757_v47 T1
				_ = _1757_v47
				var _nw254 *C4 = New_C4_()
				_ = _nw254
				_nw254.Ctor__((_this).F9(), _1741_cf6)
				_1757_v47 = _nw254
				var _1758_v48 _dafny.Sequence
				_ = _1758_v48
				_1758_v48 = _dafny.SeqOf(_1757_v47)
				var _1759_v49 _dafny.Map
				_ = _1759_v49
				_1759_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_cf6, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1758_v48, (Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1758_v48).Cardinality()))).Uint32(), _1757_v47)))
				_1756_v46 = (_1756_v46).Update((_1671_v1).Plus(_1739_cf8), ((_dafny.Zero).Minus(_1739_cf8)).Times((_1759_v49).Cardinality()))
				var _1760_v50 D6
				_ = _1760_v50
				_1760_v50 = Companion_D6_.Create_DC17_(_1741_cf6, (_this).F10(), _1671_v1, _1739_cf8, (_1685_v3).Cardinality())
				var _1761_v51 _dafny.Array
				_ = _1761_v51
				var _nwElement0_54 bool = (_1757_v47).F10()
				_ = _nwElement0_54
				var _nw255 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(2))
				_ = _nw255
				(_nw255).ArraySet1(_nwElement0_54, 0)
				(_nw255).ArraySet1(!(_1741_cf6), 1)
				_1761_v51 = _nw255
				var _1762_v52 _dafny.Map
				_ = _1762_v52
				_1762_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v1, _1761_v51)
				var _1763_v53 _dafny.Map
				_ = _1763_v53
				_1763_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1762_v52, _1741_cf6)
				var _1764_v54 _dafny.Array
				_ = _1764_v54
				var _nwElement0_55 bool = _1741_cf6
				_ = _nwElement0_55
				var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(21))
				_ = _nw256
				(_nw256).ArraySet1(_nwElement0_55, 0)
				(_nw256).ArraySet1((_this).F10(), 1)
				(_nw256).ArraySet1(_1741_cf6, 2)
				(_nw256).ArraySet1(!(_1740_cf7), 3)
				(_nw256).ArraySet1((_1760_v50).Dtor_cf29(), 4)
				(_nw256).ArraySet1((_this).F10(), 5)
				(_nw256).ArraySet1(true, 6)
				(_nw256).ArraySet1((_this).F10(), 7)
				(_nw256).ArraySet1(false, 8)
				(_nw256).ArraySet1(_1741_cf6, 9)
				(_nw256).ArraySet1((func() bool {
					if (_1763_v53).Contains(_1762_v52) {
						return (_1763_v53).Get(_1762_v52).(bool)
					}
					return true
				})(), 10)
				(_nw256).ArraySet1((_this).F10(), 11)
				(_nw256).ArraySet1((_1757_v47).F10(), 12)
				(_nw256).ArraySet1(false, 13)
				(_nw256).ArraySet1(false, 14)
				(_nw256).ArraySet1(_1740_cf7, 15)
				(_nw256).ArraySet1((_this).F10(), 16)
				(_nw256).ArraySet1(true, 17)
				(_nw256).ArraySet1(_1741_cf6, 18)
				(_nw256).ArraySet1((_1757_v47).F10(), 19)
				(_nw256).ArraySet1(_1741_cf6, 20)
				_1764_v54 = _nw256
				var _1765_v55 _dafny.Map
				_ = _1765_v55
				_1765_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_cf6, _1764_v54)
				var _1766_v56 _dafny.Map
				_ = _1766_v56
				_1766_v56 = _1765_v55
				var _1767_v57 _dafny.Array
				_ = _1767_v57
				var _nwElement0_56 _dafny.Map = _1766_v56
				_ = _nwElement0_56
				var _nw257 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(21))
				_ = _nw257
				(_nw257).ArraySet1(_nwElement0_56, 0)
				(_nw257).ArraySet1(_1766_v56, 1)
				(_nw257).ArraySet1(_1766_v56, 2)
				(_nw257).ArraySet1(_1766_v56, 3)
				(_nw257).ArraySet1(_1765_v55, 4)
				(_nw257).ArraySet1(_1766_v56, 5)
				(_nw257).ArraySet1(_1766_v56, 6)
				(_nw257).ArraySet1(_1766_v56, 7)
				(_nw257).ArraySet1(_1766_v56, 8)
				(_nw257).ArraySet1(_1766_v56, 9)
				(_nw257).ArraySet1(_1766_v56, 10)
				(_nw257).ArraySet1(_1766_v56, 11)
				(_nw257).ArraySet1(_1766_v56, 12)
				(_nw257).ArraySet1(_1766_v56, 13)
				(_nw257).ArraySet1(_1766_v56, 14)
				(_nw257).ArraySet1(_1766_v56, 15)
				(_nw257).ArraySet1(_1766_v56, 16)
				(_nw257).ArraySet1(_1766_v56, 17)
				(_nw257).ArraySet1(_1766_v56, 18)
				(_nw257).ArraySet1(_1766_v56, 19)
				(_nw257).ArraySet1((func() _dafny.Map {
					if _1740_cf7 {
						return _1766_v56
					}
					return _1766_v56
				})(), 20)
				_1767_v57 = _nw257
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1767_v57), 0))
				_ = _index271
				(_1767_v57).ArraySet1(_1766_v56, (_index271).Int())
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1767_v57), 0))
				_ = _index272
				(_1767_v57).ArraySet1(_1766_v56, (_index272).Int())
			} else if _source25.Is_DC2() {
				var _1768___mcc_h16 _dafny.Sequence = _source25.Get_().(D1_DC2).Cf4
				_ = _1768___mcc_h16
				var _1769_cf4 _dafny.Sequence = _1768___mcc_h16
				_ = _1769_cf4
				_1671_v1 = _1717_cf8
				(globalState).F8 = _1719_cf6
				var _1770_v58 _dafny.Array
				_ = _1770_v58
				var _nw258 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
				_ = _nw258
				_1770_v58 = _nw258
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1770_v58), 0))
				_ = _index273
				(_1770_v58).ArraySet1(_1726_v26, (_index273).Int())
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1770_v58), 0))
				_ = _index274
				(_1770_v58).ArraySet1(_1726_v26, (_index274).Int())
				var _1771_v59 _dafny.Map
				_ = _1771_v59
				_1771_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v1, Companion_Default___.Fm13(_1671_v1, globalState))
				_1771_v59 = (_1771_v59).Update(_1717_cf8, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}(func(_1772_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				})), _1688_v6))
			} else {
				var _1773___mcc_h17 D1 = _source25.Get_().(D1_DC5).Cf9
				_ = _1773___mcc_h17
				var _1774_cf9 D1 = _1773___mcc_h17
				_ = _1774_cf9
				var _1775_v60 D3
				_ = _1775_v60
				_1775_v60 = Companion_D3_.Create_DC7_(_1691_v9)
				var _pat_let_tv37 = _1691_v9
				_ = _pat_let_tv37
				var _1776_v61 _dafny.Array
				_ = _1776_v61
				var _nwElement0_57 D3 = _1775_v60
				_ = _nwElement0_57
				var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(8))
				_ = _nw259
				(_nw259).ArraySet1(_nwElement0_57, 0)
				(_nw259).ArraySet1(func(_pat_let38_0 D3) D3 {
					return func(_1777_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let39_0 _dafny.Array) D3 {
							return func(_1778_dt__update_hcf11_h0 _dafny.Array) D3 {
								return Companion_D3_.Create_DC7_(_1778_dt__update_hcf11_h0)
							}(_pat_let39_0)
						}(_pat_let_tv37)
					}(_pat_let38_0)
				}(_1775_v60), 1)
				(_nw259).ArraySet1(_1775_v60, 2)
				(_nw259).ArraySet1(_1775_v60, 3)
				(_nw259).ArraySet1(_1775_v60, 4)
				(_nw259).ArraySet1(_1775_v60, 5)
				(_nw259).ArraySet1(_1775_v60, 6)
				(_nw259).ArraySet1(_1775_v60, 7)
				_1776_v61 = _nw259
				var _1779_v62 _dafny.Map
				_ = _1779_v62
				_1779_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1776_v61, false)
				r1 = (func() bool {
					if _1718_cf7 {
						return _1719_cf6
					}
					return !((_1779_v62).Equals(_1779_v62))
				})()
				var _1780_v63 _dafny.Array
				_ = _1780_v63
				var _nwElement0_58 bool = _1719_cf6
				_ = _nwElement0_58
				var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.One)
				_ = _nw260
				(_nw260).ArraySet1(_nwElement0_58, 0)
				_1780_v63 = _nw260
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1780_v63), 0))
				_ = _index275
				(_1780_v63).ArraySet1(!((_this).Fm0((Companion_D8_.Create_DC23_(_dafny.SeqOf(_1719_cf6))).Dtor_cf41(), false, _1671_v1, _1718_cf7, globalState)), (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1780_v63), 0))
				_ = _index276
				(_1780_v63).ArraySet1(!(_1718_cf7), (_index276).Int())
				var _1781_v65 _dafny.Map
				_ = _1781_v65
				_1781_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1690_v8, func() _dafny.Map {
					var _coll71 = _dafny.NewMapBuilder()
					_ = _coll71
					for _iter79 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-892), _dafny.IntOfInt64(990))); ; {
						_compr_71, _ok79 := _iter79()
						if !_ok79 {
							break
						}
						var _1782_v64 _dafny.Int
						_1782_v64 = interface{}(_compr_71).(_dafny.Int)
						if ((_dafny.IntOfInt64(-892)).Cmp(_1782_v64) <= 0) && ((_1782_v64).Cmp(_dafny.IntOfInt64(990)) < 0) {
							_coll71.Add(Companion_Default___.SafeDivisionInt(_1782_v64, _dafny.IntOfInt64(547)), (Companion_D14_.Create_DC35_(_1671_v1, _1719_cf6, p0)).Dtor_cf62())
						}
					}
					return _coll71.ToMap()
				}())
				_1671_v1 = (_1781_v65).Cardinality()
				_1725_v25 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1780_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1780_v63), 0))).Int()).(bool), (_this).F10()), _1725_v25), _1725_v25)
			}
		} else if _source24.Is_DC2() {
			var _1783___mcc_h10 _dafny.Sequence = _source24.Get_().(D1_DC2).Cf4
			_ = _1783___mcc_h10
			var _1784_cf4 _dafny.Sequence = _1783___mcc_h10
			_ = _1784_cf4
			_1671_v1 = _1671_v1
			var _rhs267 bool = ((_dafny.Zero).Minus(_1671_v1)).Cmp((_dafny.Zero).Minus((_this).Fm1((_this).F9(), globalState))) < 0
			_ = _rhs267
			var _rhs268 bool = true
			_ = _rhs268
			var _rhs269 _dafny.Array = _1691_v9
			_ = _rhs269
			var _rhs270 bool = (_1671_v1).Cmp((_1690_v8).Cardinality()) != 0
			_ = _rhs270
			var _lhs183 *GlobalState = globalState
			_ = _lhs183
			var _lhs184 *GlobalState = globalState
			_ = _lhs184
			var _lhs185 *GlobalState = globalState
			_ = _lhs185
			_lhs183.F8 = _rhs267
			_lhs184.F8 = _rhs268
			_1691_v9 = _rhs269
			_lhs185.F8 = _rhs270
			var _1785_v66 _dafny.Sequence
			_ = _1785_v66
			_1785_v66 = _dafny.SeqOf((_this).F10(), false)
			if (_1785_v66).Select((Companion_Default___.SafeIndex(_1671_v1, _dafny.IntOfUint32((_1785_v66).Cardinality()))).Uint32()).(bool) {
				var _1786_v67 _dafny.Map
				_ = _1786_v67
				_1786_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1690_v8).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_1691_v9, _1691_v9)).Cardinality())) {
						return (_1690_v8).Get(_dafny.IntOfUint32((_dafny.SeqOf(_1691_v9, _1691_v9)).Cardinality())).(bool)
					}
					return (_this).F10()
				})(), _1671_v1)
				var _1787_v68 _dafny.Map
				_ = _1787_v68
				_1787_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1786_v67, _1785_v66)
				var _rhs271 bool = !(!(_1787_v68).Equals(_1787_v68)) || (false)
				_ = _rhs271
				var _rhs272 _dafny.Map = (_1690_v8).Merge((_1690_v8).Merge(_1690_v8))
				_ = _rhs272
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				_lhs186.F8 = _rhs271
				_1690_v8 = _rhs272
				var _1788_v69 _dafny.Array
				_ = _1788_v69
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_46
				var _nw261 _dafny.Array
				_ = _nw261
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw261 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.CodePoint = (func(_1789_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1790_i9 _dafny.Int) _dafny.CodePoint {
							return _1789_p0
						}
					})(p0)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw261 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw261).ArraySet1CodePoint(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw261).ArraySet1CodePoint(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1788_v69 = _nw261
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1788_v69), 0))
				_ = _index277
				(_1788_v69).ArraySet1CodePoint(_1704_v14, (_index277).Int())
				var _1791_v70 *C3
				_ = _1791_v70
				var _nw262 *C3 = New_C3_()
				_ = _nw262
				_nw262.Ctor__()
				_1791_v70 = _nw262
				var _1792_v71 _dafny.Map
				_ = _1792_v71
				_1792_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v1, _1791_v70)
				var _1793_v72 *C5
				_ = _1793_v72
				var _nw263 *C5 = New_C5_()
				_ = _nw263
				_nw263.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10()), (_this).F10())
				_1793_v72 = _nw263
				var _1794_v73 _dafny.Map
				_ = _1794_v73
				_1794_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1793_v72, _1671_v1)
				var _1795_v74 _dafny.Map
				_ = _1795_v74
				_1795_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D6_.Create_DC17_(false, (_this).F10(), _1671_v1, _dafny.IntOfInt64(-929), (_1792_v71).Cardinality())).Dtor_cf32(), (_1794_v73).Cardinality())
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1788_v69), 0))
				_ = _index278
				var _rhs273 bool = (_1685_v3).IsProperSubsetOf(Companion_Default___.Fm10((_this).F10(), (_dafny.Zero).Minus(_1671_v1), _1688_v6, _1795_v74, globalState))
				_ = _rhs273
				var _rhs274 _dafny.CodePoint = p0
				_ = _rhs274
				var _lhs187 *GlobalState = globalState
				_ = _lhs187
				var _lhs188 _dafny.Array = _1788_v69
				_ = _lhs188
				var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1788_v69), 0))
				_ = _lhs189
				_lhs187.F8 = _rhs273
				(_lhs188).ArraySet1CodePoint(_rhs274, (_lhs189).Int())
				var _1796_v75 _dafny.Map
				_ = _1796_v75
				_1796_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1784_cf4)
				_1671_v1 = (((_1796_v75).Merge(_1796_v75)).Merge((_1796_v75).Merge(_1796_v75))).Cardinality()
				var _1797_v76 *C3
				_ = _1797_v76
				var _nw264 *C3 = New_C3_()
				_ = _nw264
				_nw264.Ctor__()
				_1797_v76 = _nw264
				var _1798_v77 _dafny.MultiSet
				_ = _1798_v77
				_1798_v77 = _dafny.MultiSetOf((_this).F10(), false)
				var _1799_v78 D6
				_ = _1799_v78
				_1799_v78 = Companion_D6_.Create_DC17_((_this).F10(), (_this).F10(), (func() _dafny.Int {
					if (_1798_v77).Contains((_this).F10()) {
						return (_1798_v77).Multiplicity((_this).F10())
					}
					return _dafny.IntOfInt64(667)
				})(), _1671_v1, _1671_v1)
				var _1800_v79 D3
				_ = _1800_v79
				_1800_v79 = Companion_D3_.Create_DC8_((_this).F10(), (_this).F10())
				var _pat_let_tv38 = _1797_v76
				_ = _pat_let_tv38
				var _pat_let_tv39 = _1800_v79
				_ = _pat_let_tv39
				var _pat_let_tv40 = globalState
				_ = _pat_let_tv40
				var _pat_let_tv41 = globalState
				_ = _pat_let_tv41
				var _pat_let_tv42 = _1671_v1
				_ = _pat_let_tv42
				var _1801_v80 _dafny.Sequence
				_ = _1801_v80
				_1801_v80 = _dafny.SeqOf(_1786_v67, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func(_pat_let40_0 D6) D6 {
					return func(_1802_dt__update__tmp_h2 D6) D6 {
						return func(_pat_let41_0 _dafny.Int) D6 {
							return func(_1803_dt__update_hcf33_h0 _dafny.Int) D6 {
								return func(_pat_let42_0 _dafny.Int) D6 {
									return func(_1804_dt__update_hcf32_h0 _dafny.Int) D6 {
										return Companion_D6_.Create_DC17_((_1802_dt__update__tmp_h2).Dtor_cf29(), (_1802_dt__update__tmp_h2).Dtor_cf30(), (_1802_dt__update__tmp_h2).Dtor_cf31(), _1804_dt__update_hcf32_h0, _1803_dt__update_hcf33_h0)
									}(_pat_let42_0)
								}(_pat_let_tv42)
							}(_pat_let41_0)
						}((Companion_Default___.Fm46((_this).F10(), (_pat_let_tv38).Fm27((_this).F10(), (_this).F10(), _pat_let_tv39, _pat_let_tv40), (_this).F10(), _pat_let_tv41)).Cardinality())
					}(_pat_let40_0)
				}(_1799_v78)).Dtor_cf31()), _1786_v67)
				_1801_v80 = _dafny.SeqOf((_1786_v67).Merge(_1786_v67))
			} else {
				var _1805_v81 *C0
				_ = _1805_v81
				var _nw265 *C0 = New_C0_()
				_ = _nw265
				_nw265.Ctor__()
				_1805_v81 = _nw265
				var _1806_v82 _dafny.Array
				_ = _1806_v82
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_47
				var _nw266 _dafny.Array
				_ = _nw266
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw266 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Sequence = (func(_1807_cf4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1808_i10 _dafny.Int) _dafny.Sequence {
							return _1807_cf4
						}
					})(_1784_cf4)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw266 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw266).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw266).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1806_v82 = _nw266
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_1806_v82), 0))
				_ = _index279
				(_1806_v82).ArraySet1(_dafny.SeqOf(_1671_v1), (_index279).Int())
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_1806_v82), 0))
				_ = _index280
				(_1806_v82).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1784_cf4, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.IntOfUint32((_1784_cf4).Cardinality()))).Uint32(), _1671_v1), Companion_Default___.Fm37(globalState)), (_index280).Int())
				_1704_v14 = _1704_v14
				var _1809_v83 _dafny.Array
				_ = _1809_v83
				var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw267
				_1809_v83 = _nw267
				var _1810_v84 D17
				_ = _1810_v84
				_1810_v84 = Companion_D17_.Create_DC43_(_1809_v83)
				var _1811_v85 _dafny.Array
				_ = _1811_v85
				var _nwElement0_59 _dafny.Array = _1809_v83
				_ = _nwElement0_59
				var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(16))
				_ = _nw268
				(_nw268).ArraySet1(_nwElement0_59, 0)
				(_nw268).ArraySet1(_1809_v83, 1)
				(_nw268).ArraySet1(_1809_v83, 2)
				(_nw268).ArraySet1(_1809_v83, 3)
				(_nw268).ArraySet1(_1809_v83, 4)
				(_nw268).ArraySet1((_1810_v84).Dtor_cf83(), 5)
				(_nw268).ArraySet1(_1809_v83, 6)
				(_nw268).ArraySet1(_1809_v83, 7)
				(_nw268).ArraySet1(_1809_v83, 8)
				(_nw268).ArraySet1(_1809_v83, 9)
				(_nw268).ArraySet1(_1809_v83, 10)
				(_nw268).ArraySet1(_1809_v83, 11)
				(_nw268).ArraySet1(_1809_v83, 12)
				(_nw268).ArraySet1(_1809_v83, 13)
				(_nw268).ArraySet1((_1810_v84).Dtor_cf83(), 14)
				(_nw268).ArraySet1(_1809_v83, 15)
				_1811_v85 = _nw268
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1811_v85), 0))
				_ = _index281
				(_1811_v85).ArraySet1(_1809_v83, (_index281).Int())
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_1811_v85), 0))
				_ = _index282
				(_1811_v85).ArraySet1(_1809_v83, (_index282).Int())
				var _1812_v86 _dafny.Array
				_ = _1812_v86
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_48
				var _nw269 _dafny.Array
				_ = _nw269
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw269 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Set = (func(_1813_v6 _dafny.Sequence, _1814_v1 _dafny.Int) func(_dafny.Int) _dafny.Set {
						return func(_1815_i11 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_dafny.IntOfUint32((_1813_v6).Cardinality()), _dafny.IntOfInt64(684), _1814_v1)
						}
					})(_1688_v6, _1671_v1)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw269 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw269).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw269).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1812_v86 = _nw269
				var _1816_v87 _dafny.Set
				_ = _1816_v87
				_1816_v87 = _dafny.SetOf(_1671_v1, _1671_v1)
				var _1817_v88 D4
				_ = _1817_v88
				_1817_v88 = Companion_D4_.Create_DC11_(_1671_v1, _1816_v87, (_this).F10())
				var _1818_v89 _dafny.Set
				_ = _1818_v89
				_1818_v89 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1817_v88).Dtor_cf16(), _dafny.CodePoint('p'))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjnevp")).Cardinality()))
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1812_v86), 0))
				_ = _index283
				(_1812_v86).ArraySet1((_1816_v87).Union(_1816_v87), (_index283).Int())
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1812_v86), 0))
				_ = _index284
				var _rhs275 bool = (func() bool {
					if (_1690_v8).Contains(_1671_v1) {
						return (_1690_v8).Get(_1671_v1).(bool)
					}
					return _dafny.Companion_Sequence_.IsProperPrefixOf(_1688_v6, _1688_v6)
				})()
				_ = _rhs275
				var _rhs276 _dafny.Set = Companion_Default___.Fm25(_dafny.IntOfUint32((_1688_v6).Cardinality()), _1671_v1, globalState)
				_ = _rhs276
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				var _lhs191 _dafny.Array = _1812_v86
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1812_v86), 0))
				_ = _lhs192
				_lhs190.F8 = _rhs275
				(_lhs191).ArraySet1(_rhs276, (_lhs192).Int())
			}
			var _1819_v90 _dafny.Set
			_ = _1819_v90
			_1819_v90 = _dafny.SetOf(_1671_v1)
			_1819_v90 = _1819_v90
		} else {
			var _1820___mcc_h11 D1 = _source24.Get_().(D1_DC5).Cf9
			_ = _1820___mcc_h11
			var _1821_cf9 D1 = _1820___mcc_h11
			_ = _1821_cf9
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
			_ = _index285
			(_1691_v9).ArraySet1(_1671_v1, (_index285).Int())
			var _1822_v91 _dafny.Set
			_ = _1822_v91
			_1822_v91 = _dafny.SetOf(_1688_v6, _dafny.UnicodeSeqOfUtf8Bytes("yp"), _1688_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-425))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}((func(_1823_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1824_i12 _dafny.Int) _dafny.CodePoint {
					return _1823_v14
				}
			})(_1704_v14))), _1688_v6)
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
			_ = _index286
			var _rhs277 _dafny.Int = ((_1822_v91).Cardinality()).Minus(_1671_v1)
			_ = _rhs277
			var _rhs278 _dafny.Int = (_1671_v1).Plus(_1671_v1)
			_ = _rhs278
			var _lhs193 _dafny.Array = _1691_v9
			_ = _lhs193
			var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
			_ = _lhs194
			_1671_v1 = _rhs277
			(_lhs193).ArraySet1(_rhs278, (_lhs194).Int())
			if (_this).F10() {
				var _1825_v92 D8
				_ = _1825_v92
				_1825_v92 = Companion_D8_.Create_DC23_(_dafny.SeqOf((_this).F10(), (_this).F10()))
				var _1826_v93 _dafny.Array
				_ = _1826_v93
				var _nw270 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw270
				_1826_v93 = _nw270
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))
				_ = _index287
				(_1826_v93).ArraySet1((_this).F10(), (_index287).Int())
				var _1827_v94 _dafny.Sequence
				_ = _1827_v94
				_1827_v94 = _dafny.SeqOf((_this).F10(), (_1671_v1).Cmp(_dafny.IntOfInt64(91)) == 0)
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))
				_ = _index288
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _index289
				var _rhs279 D8 = _1825_v92
				_ = _rhs279
				var _rhs280 bool = (_1671_v1).Cmp(Companion_Default___.SafeDivisionInt(_1671_v1, _1671_v1)) < 0
				_ = _rhs280
				var _rhs281 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(983))).Uint32(), func(coer103 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg103 _dafny.Int) interface{} {
						return coer103(arg103)
					}
				}((func(_1828_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1829_i13 _dafny.Int) _dafny.Int {
						return _1828_v1
					}
				})(_1671_v1))), _1671_v1)
				_ = _rhs281
				var _rhs282 _dafny.Int = _1671_v1
				_ = _rhs282
				var _rhs283 bool = (_1827_v94).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1688_v6).Cardinality()), _dafny.IntOfUint32((_1827_v94).Cardinality()))).Uint32()).(bool)
				_ = _rhs283
				var _lhs195 _dafny.Array = _1826_v93
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))
				_ = _lhs196
				var _lhs197 *GlobalState = globalState
				_ = _lhs197
				var _lhs198 _dafny.Array = _1691_v9
				_ = _lhs198
				var _lhs199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _lhs199
				var _lhs200 *GlobalState = globalState
				_ = _lhs200
				_1825_v92 = _rhs279
				(_lhs195).ArraySet1(_rhs280, (_lhs196).Int())
				_lhs197.F8 = _rhs281
				(_lhs198).ArraySet1(_rhs282, (_lhs199).Int())
				_lhs200.F8 = _rhs283
				var _1830_v95 D14
				_ = _1830_v95
				_1830_v95 = Companion_D14_.Create_DC35_(_1671_v1, (_this).F10(), _1704_v14)
				_1704_v14 = (_1830_v95).Dtor_cf64()
				(globalState).F8 = !_dafny.Companion_Sequence_.Equal(_1688_v6, _dafny.Companion_Sequence_.Concatenate(_1688_v6, _1688_v6))
				var _1831_v96 *C3
				_ = _1831_v96
				var _nw271 *C3 = New_C3_()
				_ = _nw271
				_nw271.Ctor__()
				_1831_v96 = _nw271
				var _1832_v97 _dafny.Map
				_ = _1832_v97
				_1832_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm53(_1671_v1, globalState), _1831_v96)
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))
				_ = _index290
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _index291
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _index292
				var _rhs284 bool = !((_this).F10()) || ((_this).F10())
				_ = _rhs284
				var _rhs285 bool = (Companion_Default___.SafeModuloInt((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int), _1671_v1)).Cmp((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int)) < 0
				_ = _rhs285
				var _rhs286 bool = !(_1832_v97).Equals(_1832_v97)
				_ = _rhs286
				var _rhs287 _dafny.Int = ((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(471))
				_ = _rhs287
				var _rhs288 _dafny.Int = (_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int)
				_ = _rhs288
				var _lhs201 _dafny.Array = _1826_v93
				_ = _lhs201
				var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))
				_ = _lhs202
				var _lhs203 *GlobalState = globalState
				_ = _lhs203
				var _lhs204 *GlobalState = globalState
				_ = _lhs204
				var _lhs205 _dafny.Array = _1691_v9
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _lhs206
				var _lhs207 _dafny.Array = _1691_v9
				_ = _lhs207
				var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _lhs208
				(_lhs201).ArraySet1(_rhs284, (_lhs202).Int())
				_lhs203.F8 = _rhs285
				_lhs204.F8 = _rhs286
				(_lhs205).ArraySet1(_rhs287, (_lhs206).Int())
				(_lhs207).ArraySet1(_rhs288, (_lhs208).Int())
				var _1833_v98 _dafny.Map
				_ = _1833_v98
				_1833_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v1, (_this).F9())
				var _1834_v99 _dafny.Map
				_ = _1834_v99
				_1834_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1686_v4)
				var _1835_v100 *C2
				_ = _1835_v100
				var _nw272 *C2 = New_C2_()
				_ = _nw272
				_nw272.Ctor__((_1826_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))).Int()).(bool), _1826_v93, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1826_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))).Int()).(bool))).Merge((func() _dafny.Map {
					if (_1833_v98).Contains(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_1834_v99).Contains(!((_1826_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))).Int()).(bool))) {
							return (_1834_v99).Get(!((_1826_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))).Int()).(bool))).(_dafny.Sequence)
						}
						return _1686_v4
					})()).Cardinality())) {
						return (_1833_v98).Get(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_1834_v99).Contains(!((_1826_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))).Int()).(bool))) {
								return (_1834_v99).Get(!((_1826_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_1826_v93), 0))).Int()).(bool))).(_dafny.Sequence)
							}
							return _1686_v4
						})()).Cardinality())).(_dafny.Map)
					}
					return (_this).F9()
				})()), (_this).F10())
				_1835_v100 = _nw272
			} else {
				var _1836_v101 *C3
				_ = _1836_v101
				var _nw273 *C3 = New_C3_()
				_ = _nw273
				_nw273.Ctor__()
				_1836_v101 = _nw273
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _index293
				(_1691_v9).ArraySet1(((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int)).Minus((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int)), (_index293).Int())
				_1685_v3 = (_1685_v3).Union((_1685_v3).Union(_1685_v3))
				var _1837_v102 *C4
				_ = _1837_v102
				var _nw274 *C4 = New_C4_()
				_ = _nw274
				_nw274.Ctor__((_this).F9(), (_1671_v1).Cmp((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int)) > 0)
				_1837_v102 = _nw274
				_1837_v102 = _1837_v102
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))
				_ = _index294
				(_1691_v9).ArraySet1((_1691_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1691_v9), 0))).Int()).(_dafny.Int), (_index294).Int())
			}
			var _1838_v103 _dafny.Array
			_ = _1838_v103
			var _nw275 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw275
			_1838_v103 = _nw275
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1838_v103), 0))
			_ = _index295
			(_1838_v103).ArraySet1((_this).F10(), (_index295).Int())
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1838_v103), 0))
			_ = _index296
			(_1838_v103).ArraySet1((_this).F10(), (_index296).Int())
			var _1839_v104 _dafny.Array
			_ = _1839_v104
			var _len0_49 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_49
			var _nw276 _dafny.Array
			_ = _nw276
			if _len0_49.Cmp(_dafny.Zero) == 0 {
				_nw276 = _dafny.NewArray(_len0_49)
			} else {
				var _init49 func(_dafny.Int) _dafny.Int = (func(_1840_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1841_i14 _dafny.Int) _dafny.Int {
						return (_1841_i14).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1840_v6).Cardinality())))
					}
				})(_1688_v6)
				_ = _init49
				var _element0_49 = _init49(_dafny.Zero)
				_ = _element0_49
				_nw276 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
				(_nw276).ArraySet1(_element0_49, 0)
				var _nativeLen0_49 = (_len0_49).Int()
				_ = _nativeLen0_49
				for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
					(_nw276).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
				}
			}
			_1839_v104 = _nw276
			var _1842_v105 _dafny.Set
			_ = _1842_v105
			_1842_v105 = _dafny.SetOf(_1839_v104, _1839_v104)
			var _1843_v106 _dafny.Map
			_ = _1843_v106
			_1843_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1842_v105, (_1838_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1838_v103), 0))).Int()).(bool))
			_1843_v106 = (_1843_v106).Update((func() _dafny.Set {
				if (_1838_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1838_v103), 0))).Int()).(bool) {
					return _1842_v105
				}
				return _1842_v105
			})(), !((_1838_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(734), _dafny.ArrayLen((_1838_v103), 0))).Int()).(bool)))
		}
		var _1844_v107 _dafny.Map
		_ = _1844_v107
		_1844_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1691_v9, Companion_Default___.Fm20(_1671_v1, (_this).F10(), (_this).F10(), _1671_v1, globalState))
		r0 = _1844_v107
		r1 = !((func() bool {
			if (_this).F10() {
				return false
			}
			return (_this).F10()
		})())
		return r0, r1
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f9  _dafny.Map
	_f10 bool
	_f11 _dafny.Set
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f9 = _dafny.EmptyMap
	_this._f10 = false
	_this._f11 = _dafny.EmptySet
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

func (_this *C7) F9() _dafny.Map {
	return _this._f9
}
func (_this *C7) F10() bool {
	return _this._f10
}
func (_this *C7) Ctor__(f11 _dafny.Set, f9 _dafny.Map, f10 bool) {
	{
		(_this)._f11 = f11
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *C7) Fm0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	{
		return (_this).F10()
	}
}
func (_this *C7) Fm1(p0 _dafny.Map, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F10())).Cardinality()))
	}
}
func (_this *C7) Fm2(globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(((_this).F11()).Cardinality()), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(73))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-305)), (_dafny.Zero).Minus(_dafny.IntOfInt64(-142)))).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((Companion_D0_.Create_DC0_(!(!((_this).F10())))).Dtor_cf0())).Cardinality())) <= 0
	}
}
func (_this *C7) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32(((Companion_D1_.Create_DC2_(_dafny.SeqOf((((_this).F9()).Update(true, (_this).F10())).Cardinality()))).Dtor_cf4()).Cardinality())
	}
}
func (_this *C7) M0(globalState *GlobalState) (_dafny.Sequence, bool, _dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _1845_v0 _dafny.Int
		_ = _1845_v0
		_1845_v0 = _dafny.IntOfInt64(233)
		var _1846_v1 _dafny.Set
		_ = _1846_v1
		_1846_v1 = _dafny.SetOf(_1845_v0, _1845_v0)
		var _1847_v2 _dafny.Sequence
		_ = _1847_v2
		_1847_v2 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F10()), !((_this).F10())))
		var _hi5 _dafny.Int = (func() _dafny.Int {
			if (_this).F10() {
				return _dafny.IntOfUint32((_1847_v2).Cardinality())
			}
			return (_dafny.Zero).Minus(_1845_v0)
		})()
		_ = _hi5
		for _1848_i0 := (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1846_v1).Cardinality(), _dafny.IntOfInt64(-93))); _1848_i0.Cmp(_hi5) < 0; _1848_i0 = _1848_i0.Plus(_dafny.One) {
			var _1849_v3 _dafny.Sequence
			_ = _1849_v3
			_1849_v3 = _dafny.UnicodeSeqOfUtf8Bytes("bt")
			var _1850_v4 D0
			_ = _1850_v4
			_1850_v4 = Companion_D0_.Create_DC1_(_1845_v0, (_this).F10(), (_dafny.IntOfUint32((_1849_v3).Cardinality())).Plus(_1845_v0))
			var _1851_v6 _dafny.Map
			_ = _1851_v6
			_1851_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1848_i0)
			var _1852_v7 _dafny.Array
			_ = _1852_v7
			var _nwElement0_60 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Map {
				var _coll72 = _dafny.NewMapBuilder()
				_ = _coll72
				for _iter80 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(23), _dafny.IntOfInt64(718))); ; {
					_compr_72, _ok80 := _iter80()
					if !_ok80 {
						break
					}
					var _1853_v5 _dafny.Int
					_1853_v5 = interface{}(_compr_72).(_dafny.Int)
					if ((_dafny.IntOfInt64(23)).Cmp(_1853_v5) <= 0) && ((_1853_v5).Cmp(_dafny.IntOfInt64(718)) < 0) {
						_coll72.Add((_1853_v5).Plus(_1845_v0), _1848_i0)
					}
				}
				return _coll72.ToMap()
			}()).Cardinality()))
			_ = _nwElement0_60
			var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(6))
			_ = _nw277
			(_nw277).ArraySet1(_nwElement0_60, 0)
			(_nw277).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_1849_v3).Cardinality()))).Minus((func() _dafny.Int {
				if (_1851_v6).Contains((_this).F10()) {
					return (_1851_v6).Get((_this).F10()).(_dafny.Int)
				}
				return _1845_v0
			})()), 1)
			(_nw277).ArraySet1(_dafny.IntOfInt64(169), 2)
			(_nw277).ArraySet1((_dafny.Zero).Minus(_1845_v0), 3)
			(_nw277).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1849_v3, _dafny.UnicodeSeqOfUtf8Bytes("mootujt"))).Cardinality()), 4)
			(_nw277).ArraySet1(_1845_v0, 5)
			_1852_v7 = _nw277
			var _1854_v8 _dafny.Map
			_ = _1854_v8
			_1854_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(729))).Uint32(), func(coer104 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}((func(_1855_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1856_i1 _dafny.Int) _dafny.Int {
					return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_1855_v3, _dafny.UnicodeSeqOfUtf8Bytes("ywcvn")))).Cardinality()
				}
			})(_1849_v3))), (_this).F10())
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
			_ = _index297
			(_1852_v7).ArraySet1(((_1854_v8).Merge(_1854_v8)).Cardinality(), (_index297).Int())
			var _1857_v9 _dafny.Array
			_ = _1857_v9
			var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
			_ = _nw278
			_1857_v9 = _nw278
			var _1858_v10 _dafny.Sequence
			_ = _1858_v10
			_1858_v10 = _dafny.SeqOf((_this).F10(), (_this).F10())
			var _1859_v11 _dafny.Map
			_ = _1859_v11
			_1859_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1857_v9, _dafny.Companion_Sequence_.Update(_1858_v10, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1845_v0), _dafny.IntOfUint32((_1858_v10).Cardinality()))).Uint32(), (_this).F10()))
			var _1860_v12 _dafny.CodePoint
			_ = _1860_v12
			_1860_v12 = _dafny.CodePoint('c')
			var _1861_v13 _dafny.Sequence
			_ = _1861_v13
			_1861_v13 = _dafny.SeqOf((_1859_v11).Merge((_1859_v11).Update(_1857_v9, Companion_Default___.Fm4(_1860_v12, globalState))))
			var _1862_v14 _dafny.Sequence
			_ = _1862_v14
			_1862_v14 = _dafny.SeqOf(_dafny.SeqOf(_1848_i0))
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
			_ = _index298
			var _rhs289 bool = true
			_ = _rhs289
			var _rhs290 D0 = _1850_v4
			_ = _rhs290
			var _rhs291 _dafny.Int = ((_1861_v13).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1862_v14).Cardinality())), _dafny.IntOfUint32((_1861_v13).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()
			_ = _rhs291
			var _lhs209 *GlobalState = globalState
			_ = _lhs209
			var _lhs210 _dafny.Array = _1852_v7
			_ = _lhs210
			var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
			_ = _lhs211
			_lhs209.F8 = _rhs289
			_1850_v4 = _rhs290
			(_lhs210).ArraySet1(_rhs291, (_lhs211).Int())
			var _1863_v15 _dafny.Sequence
			_ = _1863_v15
			_1863_v15 = _dafny.SeqOf((_1852_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))).Int()).(_dafny.Int))
			var _1864_v16 D1
			_ = _1864_v16
			_1864_v16 = Companion_D1_.Create_DC2_(_1863_v15)
			var _source26 D1 = Companion_D1_.Create_DC5_(_1864_v16)
			_ = _source26
			if _source26.Is_DC3() {
				var _1865___mcc_h0 _dafny.Sequence = _source26.Get_().(D1_DC3).Cf5
				_ = _1865___mcc_h0
				var _1866_cf5 _dafny.Sequence = _1865___mcc_h0
				_ = _1866_cf5
				(globalState).F8 = (_1846_v1).IsSubsetOf((_1846_v1).Union(_1846_v1))
				(globalState).F8 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1845_v0), (_dafny.Zero).Minus((_1863_v15).Select((Companion_Default___.SafeIndex(_1845_v0, _dafny.IntOfUint32((_1863_v15).Cardinality()))).Uint32()).(_dafny.Int)))).Cmp(_1845_v0) <= 0
				var _1867_v17 *C0
				_ = _1867_v17
				var _nw279 *C0 = New_C0_()
				_ = _nw279
				_nw279.Ctor__()
				_1867_v17 = _nw279
				var _1868_v18 *C0
				_ = _1868_v18
				var _nw280 *C0 = New_C0_()
				_ = _nw280
				_nw280.Ctor__()
				_1868_v18 = _nw280
			} else if _source26.Is_DC4() {
				var _1869___mcc_h1 bool = _source26.Get_().(D1_DC4).Cf6
				_ = _1869___mcc_h1
				var _1870___mcc_h2 bool = _source26.Get_().(D1_DC4).Cf7
				_ = _1870___mcc_h2
				var _1871___mcc_h3 _dafny.Int = _source26.Get_().(D1_DC4).Cf8
				_ = _1871___mcc_h3
				var _1872_cf8 _dafny.Int = _1871___mcc_h3
				_ = _1872_cf8
				var _1873_cf7 bool = _1870___mcc_h2
				_ = _1873_cf7
				var _1874_cf6 bool = _1869___mcc_h1
				_ = _1874_cf6
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _index299
				(_1852_v7).ArraySet1(_1845_v0, (_index299).Int())
				var _1875_v19 _dafny.Array
				_ = _1875_v19
				var _len0_50 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_50
				var _nw281 _dafny.Array
				_ = _nw281
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw281 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) _dafny.Sequence = (func(_1876_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1877_i2 _dafny.Int) _dafny.Sequence {
							return _1876_v3
						}
					})(_1849_v3)
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw281 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw281).ArraySet1(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw281).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_1875_v19 = _nw281
				var _1878_v20 D1
				_ = _1878_v20
				_1878_v20 = Companion_D1_.Create_DC3_(_1849_v3)
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1875_v19), 0))
				_ = _index300
				(_1875_v19).ArraySet1((_1878_v20).Dtor_cf5(), (_index300).Int())
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1875_v19), 0))
				_ = _index301
				(_1875_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if !((_this).F10()) {
						return _1849_v3
					}
					return _1849_v3
				})(), _1849_v3), (_index301).Int())
				(globalState).F8 = _1873_cf7
				var _1879_v21 _dafny.Map
				_ = _1879_v21
				_1879_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1845_v0, _1845_v0)
				var _1880_v22 D0
				_ = _1880_v22
				_1880_v22 = Companion_D0_.Create_DC0_(!(_1874_cf6))
				var _1881_v23 _dafny.Map
				_ = _1881_v23
				_1881_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1848_i0).Minus((_1879_v21).Cardinality())), _1880_v22)
				_1881_v23 = (_1881_v23).Update(_dafny.IntOfInt64(308), func(_pat_let43_0 D0) D0 {
					return func(_1882_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let44_0 bool) D0 {
							return func(_1883_dt__update_hcf0_h0 bool) D0 {
								return Companion_D0_.Create_DC0_(_1883_dt__update_hcf0_h0)
							}(_pat_let44_0)
						}((_this).F10())
					}(_pat_let43_0)
				}(_1880_v22))
			} else if _source26.Is_DC2() {
				var _1884___mcc_h4 _dafny.Sequence = _source26.Get_().(D1_DC2).Cf4
				_ = _1884___mcc_h4
				var _1885_cf4 _dafny.Sequence = _1884___mcc_h4
				_ = _1885_cf4
				var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _index302
				(_1852_v7).ArraySet1(_1845_v0, (_index302).Int())
				var _1886_v24 _dafny.MultiSet
				_ = _1886_v24
				_1886_v24 = _dafny.MultiSetOf(_dafny.IntOfInt64(58))
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _index303
				var _rhs292 bool = (_1858_v10).Select((Companion_Default___.SafeIndex(_1845_v0, _dafny.IntOfUint32((_1858_v10).Cardinality()))).Uint32()).(bool)
				_ = _rhs292
				var _rhs293 _dafny.Int = _1848_i0
				_ = _rhs293
				var _rhs294 bool = !((_this).F10()) || ((_dafny.IntOfUint32((_1858_v10).Cardinality())).Cmp((_1886_v24).Cardinality()) < 0)
				_ = _rhs294
				var _rhs295 _dafny.Int = (_1845_v0).Times(_1845_v0)
				_ = _rhs295
				var _lhs212 _dafny.Array = _1852_v7
				_ = _lhs212
				var _lhs213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _lhs213
				var _lhs214 *GlobalState = globalState
				_ = _lhs214
				r1 = _rhs292
				(_lhs212).ArraySet1(_rhs293, (_lhs213).Int())
				_lhs214.F8 = _rhs294
				_1845_v0 = _rhs295
				var _1887_v25 D1
				_ = _1887_v25
				_1887_v25 = Companion_D1_.Create_DC4_((_this).Fm2(globalState), (_1846_v1).Equals(_1846_v1), _1848_i0)
				_1887_v25 = (func() D1 {
					if (_this).F10() {
						return _1887_v25
					}
					return _1887_v25
				})()
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1857_v9), 0))
				_ = _index304
				(_1857_v9).ArraySet1CodePoint(_1860_v12, (_index304).Int())
				var _1888_v26 _dafny.Map
				_ = _1888_v26
				_1888_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F10())
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _index305
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1857_v9), 0))
				_ = _index306
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _index307
				var _rhs296 _dafny.Sequence = Companion_Default___.Fm6(_1845_v0, _1845_v0, _1888_v26, globalState)
				_ = _rhs296
				var _rhs297 _dafny.Int = Companion_Default___.SafeDivisionInt((_1852_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))).Int()).(_dafny.Int), ((_dafny.Zero).Minus((_1852_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))).Int()).(_dafny.Int))).Times(_dafny.IntOfInt64(678)))
				_ = _rhs297
				var _rhs298 _dafny.Int = _1848_i0
				_ = _rhs298
				var _rhs299 _dafny.CodePoint = Companion_Default___.Fm7(_1845_v0, globalState)
				_ = _rhs299
				var _rhs300 _dafny.Int = (_dafny.Zero).Minus((_1852_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))).Int()).(_dafny.Int))
				_ = _rhs300
				var _lhs215 _dafny.Array = _1852_v7
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _lhs216
				var _lhs217 _dafny.Array = _1857_v9
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1857_v9), 0))
				_ = _lhs218
				var _lhs219 _dafny.Array = _1852_v7
				_ = _lhs219
				var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))
				_ = _lhs220
				r0 = _rhs296
				_1845_v0 = _rhs297
				(_lhs215).ArraySet1(_rhs298, (_lhs216).Int())
				(_lhs217).ArraySet1CodePoint(_rhs299, (_lhs218).Int())
				(_lhs219).ArraySet1(_rhs300, (_lhs220).Int())
			} else {
				var _1889___mcc_h5 D1 = _source26.Get_().(D1_DC5).Cf9
				_ = _1889___mcc_h5
				var _1890_cf9 D1 = _1889___mcc_h5
				_ = _1890_cf9
				var _1891_v27 D1
				_ = _1891_v27
				_1891_v27 = Companion_D1_.Create_DC5_(_1864_v16)
				_1891_v27 = _1891_v27
				var _1892_v28 _dafny.MultiSet
				_ = _1892_v28
				_1892_v28 = _dafny.MultiSetOf((_this).F10())
				var _1893_v29 _dafny.Array
				_ = _1893_v29
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_51
				var _nw282 _dafny.Array
				_ = _nw282
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw282 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) bool = func(_1894_i3 _dafny.Int) bool {
						return (_this).F10()
					}
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw282 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw282).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw282).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1893_v29 = _nw282
				var _1895_v30 T1
				_ = _1895_v30
				var _nw283 *C2 = New_C2_()
				_ = _nw283
				_nw283.Ctor__((_this).F10(), _1893_v29, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10()), false)
				_1895_v30 = _nw283
				var _1896_v31 _dafny.MultiSet
				_ = _1896_v31
				_1896_v31 = _dafny.MultiSetOf(_1895_v30)
				var _1897_v32 _dafny.Map
				_ = _1897_v32
				_1897_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1896_v31).Update(_1895_v30, Companion_Default___.Abs(_dafny.IntOfUint32((_1858_v10).Cardinality()))), _1863_v15)
				var _rhs301 _dafny.Int = (((_1897_v32).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1896_v31, _1863_v15)).Update(_dafny.MultiSetOf(_1895_v30, _1895_v30), _1863_v15))).Cardinality()).Plus((_dafny.Zero).Minus((_1852_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1852_v7), 0))).Int()).(_dafny.Int)))
				_ = _rhs301
				var _rhs302 _dafny.MultiSet = (_1892_v28).Intersection(_1892_v28)
				_ = _rhs302
				var _rhs303 _dafny.Array = _1852_v7
				_ = _rhs303
				var _rhs304 bool = !((_1895_v30).F10())
				_ = _rhs304
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				_1845_v0 = _rhs301
				_1892_v28 = _rhs302
				_1852_v7 = _rhs303
				_lhs221.F8 = _rhs304
				var _1898_v33 _dafny.Array
				_ = _1898_v33
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw284
				_1898_v33 = _nw284
				var _1899_v34 _dafny.Map
				_ = _1899_v34
				_1899_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1895_v30).F10(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), !((_1895_v30).F10())))
				var _1900_v35 _dafny.Map
				_ = _1900_v35
				_1900_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F10())
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1898_v33), 0))
				_ = _index308
				(_1898_v33).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6(_1845_v0, _1848_i0, (func() _dafny.Map {
					if (_1899_v34).Contains(true) {
						return (_1899_v34).Get(true).(_dafny.Map)
					}
					return _1900_v35
				})(), globalState), _1849_v3), (_index308).Int())
				var _1901_v36 _dafny.Array
				_ = _1901_v36
				var _nw285 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
				_ = _nw285
				_1901_v36 = _nw285
				var _1902_v37 _dafny.Map
				_ = _1902_v37
				_1902_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1860_v12, _1848_i0)
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_1901_v36), 0))
				_ = _index309
				(_1901_v36).ArraySet1(_1902_v37, (_index309).Int())
				var _1903_v38 D3
				_ = _1903_v38
				_1903_v38 = Companion_D3_.Create_DC7_(_1852_v7)
				var _1904_v39 D3
				_ = _1904_v39
				_1904_v39 = Companion_D3_.Create_DC9_(_1903_v38)
				var _1905_v41 _dafny.Map
				_ = _1905_v41
				_1905_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1860_v12, _1860_v12)
				var _1906_v42 D18
				_ = _1906_v42
				_1906_v42 = Companion_D18_.Create_DC46_(_1905_v41)
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1898_v33), 0))
				_ = _index310
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_1901_v36), 0))
				_ = _index311
				var _rhs305 _dafny.Sequence = _1849_v3
				_ = _rhs305
				var _rhs306 _dafny.Map = func() _dafny.Map {
					var _coll73 = _dafny.NewMapBuilder()
					_ = _coll73
					for _iter81 := _dafny.Iterate(((_1906_v42).Dtor_cf89()).Keys().Elements()); ; {
						_compr_73, _ok81 := _iter81()
						if !_ok81 {
							break
						}
						var _1907_v40 _dafny.CodePoint
						_1907_v40 = interface{}(_compr_73).(_dafny.CodePoint)
						if ((_1906_v42).Dtor_cf89()).Contains(_1907_v40) {
							_coll73.Add(_1907_v40, (_1848_i0).Plus(_dafny.IntOfInt64(51)))
						}
					}
					return _coll73.ToMap()
				}()
				_ = _rhs306
				var _rhs307 D3 = _1904_v39
				_ = _rhs307
				var _rhs308 bool = (_this).F10()
				_ = _rhs308
				var _lhs222 _dafny.Array = _1898_v33
				_ = _lhs222
				var _lhs223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1898_v33), 0))
				_ = _lhs223
				var _lhs224 _dafny.Array = _1901_v36
				_ = _lhs224
				var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_1901_v36), 0))
				_ = _lhs225
				(_lhs222).ArraySet1(_rhs305, (_lhs223).Int())
				(_lhs224).ArraySet1(_rhs306, (_lhs225).Int())
				_1904_v39 = _rhs307
				r1 = _rhs308
				var _1908_v43 _dafny.CodePoint
				_ = _1908_v43
				_1908_v43 = _1860_v12
				var _1909_v44 _dafny.Map
				_ = _1909_v44
				_1909_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if (_this).F10() {
						return _1908_v43
					}
					return _1908_v43
				})(), (_1898_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1898_v33), 0))).Int()).(_dafny.Sequence))
				_1909_v44 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1908_v43, (_1898_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1898_v33), 0))).Int()).(_dafny.Sequence))).Merge(_1909_v44)
			}
			var _1910_v45 _dafny.Array
			_ = _1910_v45
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_52
			var _nw286 _dafny.Array
			_ = _nw286
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw286 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Sequence = (func(_1911_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1912_i4 _dafny.Int) _dafny.Sequence {
						return _1911_v10
					}
				})(_1858_v10)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw286 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw286).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw286).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1910_v45 = _nw286
			var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_1910_v45), 0))
			_ = _index312
			(_1910_v45).ArraySet1(_1858_v10, (_index312).Int())
			var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_1910_v45), 0))
			_ = _index313
			var _rhs309 _dafny.Int = Companion_Default___.SafeModuloInt(_1845_v0, Companion_Default___.Fm33((_this).F10(), (_this).F10(), _1848_i0, (_this).F10(), globalState))
			_ = _rhs309
			var _rhs310 _dafny.CodePoint = _dafny.CodePoint('w')
			_ = _rhs310
			var _rhs311 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1858_v10, _dafny.SeqOf((_this).Fm2(globalState), (_this).F10(), (_this).F10()))
			_ = _rhs311
			var _lhs226 _dafny.Array = _1910_v45
			_ = _lhs226
			var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_1910_v45), 0))
			_ = _lhs227
			_1845_v0 = _rhs309
			_1860_v12 = _rhs310
			(_lhs226).ArraySet1(_rhs311, (_lhs227).Int())
			var _1913_v46 *C0
			_ = _1913_v46
			var _nw287 *C0 = New_C0_()
			_ = _nw287
			_nw287.Ctor__()
			_1913_v46 = _nw287
		}
		var _1914_v47 T0
		_ = _1914_v47
		var _nw288 *C3 = New_C3_()
		_ = _nw288
		_nw288.Ctor__()
		_1914_v47 = _nw288
		var _1915_v48 _dafny.MultiSet
		_ = _1915_v48
		_1915_v48 = _dafny.MultiSetOf(_1914_v47, _1914_v47, _1914_v47)
		var _1916_v49 _dafny.Map
		_ = _1916_v49
		_1916_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1845_v0, _1915_v48)
		var _hi6 _dafny.Int = (_1916_v49).Cardinality()
		_ = _hi6
		for _1917_i5 := Companion_Default___.SafeDivisionInt(_1845_v0, _1845_v0); _1917_i5.Cmp(_hi6) < 0; _1917_i5 = _1917_i5.Plus(_dafny.One) {
			r1 = (_this).F10()
			if ((_this).F10()) == ((_this).F10()) {
				var _1918_v50 _dafny.Array
				_ = _1918_v50
				var _nw289 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw289
				_1918_v50 = _nw289
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1918_v50), 0))
				_ = _index314
				(_1918_v50).ArraySet1(_1917_i5, (_index314).Int())
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1918_v50), 0))
				_ = _index315
				(_1918_v50).ArraySet1(_1917_i5, (_index315).Int())
				(globalState).F8 = (_this).Fm2(globalState)
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1918_v50), 0))
				_ = _index316
				(_1918_v50).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if !(!((_this).F10()) || ((_this).F10())) {
						return _1917_i5
					}
					return (_1918_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1918_v50), 0))).Int()).(_dafny.Int)
				})()), (_index316).Int())
				var _1919_v51 _dafny.Sequence
				_ = _1919_v51
				_1919_v51 = _dafny.SeqOf((_1918_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1918_v50), 0))).Int()).(_dafny.Int))
				var _1920_v52 _dafny.Sequence
				_ = _1920_v52
				_1920_v52 = _dafny.UnicodeSeqOfUtf8Bytes("p")
				_1919_v51 = _dafny.SeqOf((_1845_v0).Times(_dafny.IntOfUint32((_1920_v52).Cardinality())))
				(globalState).F8 = (_this).F10()
			} else {
				var _1921_v53 _dafny.CodePoint
				_ = _1921_v53
				_1921_v53 = _dafny.CodePoint('x')
				var _1922_v54 D6
				_ = _1922_v54
				_1922_v54 = Companion_D6_.Create_DC19_((_this).F10(), _1921_v53)
				var _pat_let_tv43 = _1921_v53
				_ = _pat_let_tv43
				(globalState).F8 = (func(_pat_let45_0 D6) D6 {
					return func(_1923_dt__update__tmp_h1 D6) D6 {
						return func(_pat_let46_0 _dafny.CodePoint) D6 {
							return func(_1924_dt__update_hcf38_h0 _dafny.CodePoint) D6 {
								return Companion_D6_.Create_DC19_((_1923_dt__update__tmp_h1).Dtor_cf37(), _1924_dt__update_hcf38_h0)
							}(_pat_let46_0)
						}(_pat_let_tv43)
					}(_pat_let45_0)
				}(_1922_v54)).Dtor_cf37()
				var _1925_v55 *C3
				_ = _1925_v55
				var _nw290 *C3 = New_C3_()
				_ = _nw290
				_nw290.Ctor__()
				_1925_v55 = _nw290
				var _1926_v56 *C0
				_ = _1926_v56
				var _nw291 *C0 = New_C0_()
				_ = _nw291
				_nw291.Ctor__()
				_1926_v56 = _nw291
				_1845_v0 = Companion_Default___.SafeDivisionInt(_1845_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_1845_v0)))
				var _1927_v57 _dafny.Map
				_ = _1927_v57
				_1927_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1845_v0, _1917_i5)
				var _1928_v58 _dafny.Map
				_ = _1928_v58
				_1928_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1927_v57)
				var _1929_v59 _dafny.Sequence
				_ = _1929_v59
				_1929_v59 = _dafny.UnicodeSeqOfUtf8Bytes("whn")
				var _1930_v60 _dafny.Map
				_ = _1930_v60
				_1930_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1845_v0, _dafny.IntOfUint32((_1929_v59).Cardinality()))
				_1928_v58 = (_1928_v58).Update((_this).F10(), _1930_v60)
			}
			if (_this).F10() {
				r1 = (_this).F10()
				var _1931_v61 _dafny.Array
				_ = _1931_v61
				var _len0_53 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_53
				var _nw292 _dafny.Array
				_ = _nw292
				if _len0_53.Cmp(_dafny.Zero) == 0 {
					_nw292 = _dafny.NewArray(_len0_53)
				} else {
					var _init53 func(_dafny.Int) _dafny.Int = func(_1932_i6 _dafny.Int) _dafny.Int {
						return (_1932_i6).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-602))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg105 _dafny.Int) interface{} {
								return coer105(arg105)
							}
						}(func(_1933_i7 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						}))).Cardinality()))
					}
					_ = _init53
					var _element0_53 = _init53(_dafny.Zero)
					_ = _element0_53
					_nw292 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
					(_nw292).ArraySet1(_element0_53, 0)
					var _nativeLen0_53 = (_len0_53).Int()
					_ = _nativeLen0_53
					for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
						(_nw292).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
					}
				}
				_1931_v61 = _nw292
				var _1934_v62 _dafny.Sequence
				_ = _1934_v62
				_1934_v62 = _dafny.SeqOf(_1931_v61)
				var _1935_v63 D12
				_ = _1935_v63
				_1935_v63 = Companion_D12_.Create_DC32_(_1931_v61, _dafny.Companion_Sequence_.Update(_1934_v62, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.IntOfUint32((_1934_v62).Cardinality()))).Uint32(), _1931_v61))
				_1935_v63 = _1935_v63
				var _1936_v64 _dafny.MultiSet
				_ = _1936_v64
				_1936_v64 = _dafny.MultiSetOf((_this).F10())
				var _1937_v65 _dafny.MultiSet
				_ = _1937_v65
				_1937_v65 = _dafny.MultiSetOf(((_this).F9()).Cardinality())
				var _1938_v66 _dafny.CodePoint
				_ = _1938_v66
				_1938_v66 = _dafny.CodePoint('a')
				var _1939_v67 _dafny.MultiSet
				_ = _1939_v67
				_1939_v67 = _dafny.MultiSetOf(_1938_v66)
				var _1940_v68 _dafny.Sequence
				_ = _1940_v68
				_1940_v68 = _dafny.UnicodeSeqOfUtf8Bytes("dokjnmwo")
				var _1941_v69 _dafny.Map
				_ = _1941_v69
				_1941_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1917_i5, _1940_v68)
				var _1942_v70 _dafny.Sequence
				_ = _1942_v70
				_1942_v70 = _dafny.SeqOf(_1940_v68)
				var _1943_v71 _dafny.Array
				_ = _1943_v71
				var _nwElement0_61 bool = (_this).F10()
				_ = _nwElement0_61
				var _nw293 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(20))
				_ = _nw293
				(_nw293).ArraySet1(_nwElement0_61, 0)
				(_nw293).ArraySet1((_this).F10(), 1)
				(_nw293).ArraySet1((_1917_i5).Cmp(_1845_v0) <= 0, 2)
				(_nw293).ArraySet1((_this).F10(), 3)
				(_nw293).ArraySet1(!((_this).F10()) || ((_this).F10()), 4)
				(_nw293).ArraySet1((_this).F10(), 5)
				(_nw293).ArraySet1((_1917_i5).Cmp((func() _dafny.Int {
					if (_1936_v64).Contains(true) {
						return (_1936_v64).Multiplicity(true)
					}
					return _1845_v0
				})()) >= 0, 6)
				(_nw293).ArraySet1((_this).F10(), 7)
				(_nw293).ArraySet1((_this).F10(), 8)
				(_nw293).ArraySet1((_this).F10(), 9)
				(_nw293).ArraySet1(!((_1937_v65).IsProperSubsetOf(_1937_v65)), 10)
				(_nw293).ArraySet1((_1939_v67).IsProperSubsetOf(_1939_v67), 11)
				(_nw293).ArraySet1(!_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
					if (_1941_v69).Contains(_1917_i5) {
						return (_1941_v69).Get(_1917_i5).(_dafny.Sequence)
					}
					return (_1942_v70).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1845_v0), _dafny.IntOfUint32((_1942_v70).Cardinality()))).Uint32()).(_dafny.Sequence)
				})(), _1938_v66), 12)
				(_nw293).ArraySet1((_this).F10(), 13)
				(_nw293).ArraySet1((_this).F10(), 14)
				(_nw293).ArraySet1(!(true), 15)
				(_nw293).ArraySet1(!((_this).F10()), 16)
				(_nw293).ArraySet1((_this).F10(), 17)
				(_nw293).ArraySet1(true, 18)
				(_nw293).ArraySet1((func() bool {
					if false {
						return (_this).F10()
					}
					return (_this).F10()
				})(), 19)
				_1943_v71 = _nw293
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1943_v71), 0))
				_ = _index317
				(_1943_v71).ArraySet1((_this).F10(), (_index317).Int())
				var _1944_v72 D12
				_ = _1944_v72
				_1944_v72 = Companion_D12_.Create_DC31_((_this).F10(), _1940_v68, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg106 _dafny.Int) interface{} {
						return coer106(arg106)
					}
				}((func(_1945_v66 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1946_i8 _dafny.Int) _dafny.CodePoint {
						return _1945_v66
					}
				})(_1938_v66))))
				var _pat_let_tv44 = _1940_v68
				_ = _pat_let_tv44
				var _pat_let_tv45 = _1845_v0
				_ = _pat_let_tv45
				var _pat_let_tv46 = _1938_v66
				_ = _pat_let_tv46
				var _1947_v73 T1
				_ = _1947_v73
				var _nw294 *C2 = New_C2_()
				_ = _nw294
				_nw294.Ctor__(true, _1943_v71, (_this).F9(), (func(_pat_let47_0 D12) D12 {
					return func(_1948_dt__update__tmp_h2 D12) D12 {
						return func(_pat_let48_0 _dafny.Sequence) D12 {
							return func(_1949_dt__update_hcf57_h0 _dafny.Sequence) D12 {
								return func(_pat_let49_0 _dafny.Sequence) D12 {
									return func(_1950_dt__update_hcf56_h0 _dafny.Sequence) D12 {
										return Companion_D12_.Create_DC31_((_1948_dt__update__tmp_h2).Dtor_cf55(), _1950_dt__update_hcf56_h0, _1949_dt__update_hcf57_h0)
									}(_pat_let49_0)
								}(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("cguqrqn"), (Companion_Default___.SafeIndex(_pat_let_tv45, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cguqrqn")).Cardinality()))).Uint32(), _pat_let_tv46))
							}(_pat_let48_0)
						}(_pat_let_tv44)
					}(_pat_let47_0)
				}(_1944_v72)).Dtor_cf55())
				_1947_v73 = _nw294
				var _1951_v74 D16
				_ = _1951_v74
				_1951_v74 = Companion_D16_.Create_DC41_((_1947_v73).F10(), _dafny.IntOfInt64(563), (_this).F10())
				var _1952_v75 _dafny.Sequence
				_ = _1952_v75
				_1952_v75 = _dafny.SeqOf(_dafny.IntOfInt64(576), (_1951_v74).Dtor_cf78())
				var _1953_v76 _dafny.Sequence
				_ = _1953_v76
				_1953_v76 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1952_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(194))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg107 _dafny.Int) interface{} {
						return coer107(arg107)
					}
				}(func(_1954_i9 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality())
				})))))
				var _1955_v77 _dafny.MultiSet
				_ = _1955_v77
				_1955_v77 = _dafny.MultiSetOf(_1952_v75, _1952_v75)
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1943_v71), 0))
				_ = _index318
				var _rhs312 bool = (_1955_v77).IsProperSubsetOf((_1953_v76).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1937_v65).Contains(_1917_i5) {
						return (_1937_v65).Multiplicity(_1917_i5)
					}
					return _1845_v0
				})(), _dafny.IntOfUint32((_1953_v76).Cardinality()))).Uint32()).(_dafny.MultiSet))
				_ = _rhs312
				var _rhs313 T1 = _1947_v73
				_ = _rhs313
				var _lhs228 _dafny.Array = _1943_v71
				_ = _lhs228
				var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1943_v71), 0))
				_ = _lhs229
				(_lhs228).ArraySet1(_rhs312, (_lhs229).Int())
				_1947_v73 = _rhs313
				var _1956_v78 *C2
				_ = _1956_v78
				var _nw295 *C2 = New_C2_()
				_ = _nw295
				_nw295.Ctor__((_1943_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1943_v71), 0))).Int()).(bool), _1943_v71, (_1947_v73).F9(), (_1943_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1943_v71), 0))).Int()).(bool))
				_1956_v78 = _nw295
				var _1957_v79 _dafny.Array
				_ = _1957_v79
				var _nwElement0_62 *C2 = _1956_v78
				_ = _nwElement0_62
				var _nw296 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(18))
				_ = _nw296
				(_nw296).ArraySet1(_nwElement0_62, 0)
				(_nw296).ArraySet1(_1956_v78, 1)
				(_nw296).ArraySet1(_1956_v78, 2)
				(_nw296).ArraySet1(_1956_v78, 3)
				(_nw296).ArraySet1(_1956_v78, 4)
				(_nw296).ArraySet1(_1956_v78, 5)
				(_nw296).ArraySet1(_1956_v78, 6)
				(_nw296).ArraySet1(_1956_v78, 7)
				(_nw296).ArraySet1((func() *C2 {
					if _1956_v78.F14 {
						return _1956_v78
					}
					return _1956_v78
				})(), 8)
				(_nw296).ArraySet1(_1956_v78, 9)
				(_nw296).ArraySet1(_1956_v78, 10)
				(_nw296).ArraySet1(_1956_v78, 11)
				(_nw296).ArraySet1(_1956_v78, 12)
				(_nw296).ArraySet1(_1956_v78, 13)
				(_nw296).ArraySet1(_1956_v78, 14)
				(_nw296).ArraySet1(_1956_v78, 15)
				(_nw296).ArraySet1(_1956_v78, 16)
				(_nw296).ArraySet1(_1956_v78, 17)
				_1957_v79 = _nw296
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1957_v79), 0))
				_ = _index319
				(_1957_v79).ArraySet1(_1956_v78, (_index319).Int())
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1957_v79), 0))
				_ = _index320
				(_1957_v79).ArraySet1(_1956_v78, (_index320).Int())
				_1936_v64 = (_1936_v64).Union((_dafny.MultiSetOf(_1956_v78.F14, (_1947_v73).F10())).Intersection(_1936_v64))
			} else {
				r1 = (_this).F10()
				(globalState).F8 = !((_this).F10())
				var _1958_v80 _dafny.Sequence
				_ = _1958_v80
				_1958_v80 = _dafny.SeqOf(_dafny.IntOfInt64(623))
				var _1959_v81 _dafny.Sequence
				_ = _1959_v81
				_1959_v81 = _dafny.SeqOf(_1917_i5, _dafny.IntOfUint32((_1958_v80).Cardinality()), _dafny.IntOfInt64(227))
				var _1960_v82 _dafny.Array
				_ = _1960_v82
				var _nwElement0_63 bool = (_this).F10()
				_ = _nwElement0_63
				var _nw297 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(8))
				_ = _nw297
				(_nw297).ArraySet1(_nwElement0_63, 0)
				(_nw297).ArraySet1((_this).F10(), 1)
				(_nw297).ArraySet1((_1845_v0).Cmp(_1845_v0) == 0, 2)
				(_nw297).ArraySet1((_this).F10(), 3)
				(_nw297).ArraySet1((_this).F10(), 4)
				(_nw297).ArraySet1((_dafny.IntOfUint32((_1958_v80).Cardinality())).Cmp(((_this).F11()).Cardinality()) > 0, 5)
				(_nw297).ArraySet1((_this).F10(), 6)
				(_nw297).ArraySet1((_this).F10(), 7)
				_1960_v82 = _nw297
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1960_v82), 0))
				_ = _index321
				(_1960_v82).ArraySet1((_this).F10(), (_index321).Int())
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1960_v82), 0))
				_ = _index322
				(_1960_v82).ArraySet1(((_this).F10()) || ((_this).F10()), (_index322).Int())
				var _rhs314 _dafny.Int = _1917_i5
				_ = _rhs314
				var _rhs315 _dafny.Array = _1960_v82
				_ = _rhs315
				var _rhs316 bool = (_1960_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1960_v82), 0))).Int()).(bool)
				_ = _rhs316
				var _rhs317 _dafny.Array = _1960_v82
				_ = _rhs317
				var _lhs230 *GlobalState = globalState
				_ = _lhs230
				_1845_v0 = _rhs314
				_1960_v82 = _rhs315
				_lhs230.F8 = _rhs316
				_1960_v82 = _rhs317
				_1845_v0 = (Companion_Default___.Fm54(globalState)).Cardinality()
			}
			var _1961_v83 _dafny.CodePoint
			_ = _1961_v83
			_1961_v83 = _dafny.CodePoint('f')
			var _1962_v84 _dafny.Sequence
			_ = _1962_v84
			_1962_v84 = _dafny.SeqOf((_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10()), globalState))
			var _1963_v85 _dafny.Array
			_ = _1963_v85
			var _nwElement0_64 _dafny.Int = (_1962_v84).Select((Companion_Default___.SafeIndex(_1845_v0, _dafny.IntOfUint32((_1962_v84).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _nwElement0_64
			var _nw298 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(4))
			_ = _nw298
			(_nw298).ArraySet1(_nwElement0_64, 0)
			(_nw298).ArraySet1(_1917_i5, 1)
			(_nw298).ArraySet1(_1845_v0, 2)
			(_nw298).ArraySet1(_1845_v0, 3)
			_1963_v85 = _nw298
			var _1964_v86 _dafny.Map
			_ = _1964_v86
			_1964_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1963_v85, (_this).F10())
			var _1965_v87 _dafny.Sequence
			_ = _1965_v87
			_1965_v87 = _dafny.SeqOf(_1964_v86)
			var _1966_v88 _dafny.Map
			_ = _1966_v88
			_1966_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1961_v83, (_1965_v87).Select((Companion_Default___.SafeIndex(_1845_v0, _dafny.IntOfUint32((_1965_v87).Cardinality()))).Uint32()).(_dafny.Map))
			var _1967_v89 _dafny.Map
			_ = _1967_v89
			_1967_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1963_v85)
			var _1968_v90 _dafny.Sequence
			_ = _1968_v90
			_1968_v90 = _dafny.SeqOf(!((_this).F10()), (_this).F10(), !((_this).F10()), true, (_this).F10())
			var _1969_v91 _dafny.Sequence
			_ = _1969_v91
			_1969_v91 = _dafny.UnicodeSeqOfUtf8Bytes("uqmwx")
			var _1970_v92 _dafny.Map
			_ = _1970_v92
			_1970_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1917_i5, (_1964_v86).Update((func() _dafny.Array {
				if (_1967_v89).Contains((_this).F10()) {
					return (_1967_v89).Get((_this).F10()).(_dafny.Array)
				}
				return _1963_v85
			})(), (_1968_v90).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1969_v91).Cardinality()), _dafny.IntOfUint32((_1968_v90).Cardinality()))).Uint32()).(bool)))
			var _1971_v93 *C5
			_ = _1971_v93
			var _nw299 *C5 = New_C5_()
			_ = _nw299
			_nw299.Ctor__((_this).F9(), (_this).F10())
			_1971_v93 = _nw299
			var _1972_v94 *C6
			_ = _1972_v94
			var _nw300 *C6 = New_C6_()
			_ = _nw300
			_nw300.Ctor__((_this).F9(), true)
			_1972_v94 = _nw300
			var _1973_v95 _dafny.Map
			_ = _1973_v95
			_1973_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1971_v93, _1972_v94)
			var _1974_v96 D8
			_ = _1974_v96
			_1974_v96 = Companion_D8_.Create_DC24_((_this).F10(), _dafny.SetOf(_1963_v85, _1963_v85), (_1973_v95).Cardinality())
			if (((func() _dafny.Map {
				if (_1966_v88).Contains(_1961_v83) {
					return (_1966_v88).Get(_1961_v83).(_dafny.Map)
				}
				return _1964_v86
			})()).Merge((func() _dafny.Map {
				if (_1970_v92).Contains((_1974_v96).Dtor_cf44()) {
					return (_1970_v92).Get((_1974_v96).Dtor_cf44()).(_dafny.Map)
				}
				return _1964_v86
			})())).Equals((_1964_v86).Merge(_1964_v86)) {
				(globalState).F8 = (_this).F10()
				var _1975_v97 _dafny.Array
				_ = _1975_v97
				var _nw301 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw301
				_1975_v97 = _nw301
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1975_v97), 0))
				_ = _index323
				(_1975_v97).ArraySet1(!((_this).F10()), (_index323).Int())
				var _1976_v98 _dafny.Map
				_ = _1976_v98
				_1976_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1917_i5), (_this).F10())
				var _1977_v99 D16
				_ = _1977_v99
				_1977_v99 = Companion_D16_.Create_DC42_((_this).F10(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("im")).Cardinality()), _1917_i5)
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1975_v97), 0))
				_ = _index324
				(_1975_v97).ArraySet1(((func() D16 {
					if (func() bool {
						if (_1976_v98).Contains(_1846_v1) {
							return (_1976_v98).Get(_1846_v1).(bool)
						}
						return (_this).F10()
					})() {
						return _1977_v99
					}
					return _1977_v99
				})()).Dtor_cf80(), (_index324).Int())
				var _1978_v101 *C2
				_ = _1978_v101
				var _nw302 *C2 = New_C2_()
				_ = _nw302
				_nw302.Ctor__((_1975_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1975_v97), 0))).Int()).(bool), _1975_v97, (_this).F9(), (_1968_v90).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).Fm3((func() _dafny.Set {
					var _coll74 = _dafny.NewBuilder()
					_ = _coll74
					for _iter82 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(47), _dafny.IntOfInt64(633))); ; {
						_compr_74, _ok82 := _iter82()
						if !_ok82 {
							break
						}
						var _1979_v100 _dafny.Int
						_1979_v100 = interface{}(_compr_74).(_dafny.Int)
						if ((_dafny.IntOfInt64(47)).Cmp(_1979_v100) <= 0) && ((_1979_v100).Cmp(_dafny.IntOfInt64(633)) < 0) {
							_coll74.Add((_1979_v100).Plus(_1845_v0))
						}
					}
					return _coll74.ToSet()
				}()).Cardinality(), _dafny.IntOfInt64(113), _1961_v83, _1845_v0, globalState)), _dafny.IntOfUint32((_1968_v90).Cardinality()))).Uint32()).(bool))
				_1978_v101 = _nw302
				var _1980_v102 _dafny.Map
				_ = _1980_v102
				_1980_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1975_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1975_v97), 0))).Int()).(bool), _1978_v101)
				_1980_v102 = _1980_v102
				_1845_v0 = _1917_i5
				_1845_v0 = Companion_Default___.SafeDivisionInt(_1917_i5, _1845_v0)
			} else {
				var _1981_v103 _dafny.MultiSet
				_ = _1981_v103
				_1981_v103 = _dafny.MultiSetOf((_this).F10(), true, (_this).F10(), ((_this).F10()) && ((_this).F10()))
				_1845_v0 = (func() _dafny.Int {
					if (_1981_v103).Contains((_this).F10()) {
						return (_1981_v103).Multiplicity((_this).F10())
					}
					return _1845_v0
				})()
				_1963_v85 = _1963_v85
				var _1982_v104 _dafny.Array
				_ = _1982_v104
				var _nw303 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
				_ = _nw303
				_1982_v104 = _nw303
				var _1983_v105 D19
				_ = _1983_v105
				_1983_v105 = Companion_D19_.Create_DC48_(_1982_v104)
				var _1984_v106 _dafny.Array
				_ = _1984_v106
				var _nwElement0_65 _dafny.Array = _1982_v104
				_ = _nwElement0_65
				var _nw304 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(6))
				_ = _nw304
				(_nw304).ArraySet1(_nwElement0_65, 0)
				(_nw304).ArraySet1(_1982_v104, 1)
				(_nw304).ArraySet1(_1982_v104, 2)
				(_nw304).ArraySet1(_1982_v104, 3)
				(_nw304).ArraySet1(_1982_v104, 4)
				(_nw304).ArraySet1((_1983_v105).Dtor_cf92(), 5)
				_1984_v106 = _nw304
				_1984_v106 = _1984_v106
				var _1985_v107 *C1
				_ = _1985_v107
				var _nw305 *C1 = New_C1_()
				_ = _nw305
				_nw305.Ctor__((_this).F10(), _1969_v91)
				_1985_v107 = _nw305
				var _rhs318 *C1 = _1985_v107
				_ = _rhs318
				var _rhs319 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1962_v84, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(313))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg108 _dafny.Int) interface{} {
						return coer108(arg108)
					}
				}((func(_1986_v84 _dafny.Sequence, _1987_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1988_i10 _dafny.Int) _dafny.Int {
						return (_1986_v84).Select((Companion_Default___.SafeIndex(_1987_v0, _dafny.IntOfUint32((_1986_v84).Cardinality()))).Uint32()).(_dafny.Int)
					}
				})(_1962_v84, _1845_v0))))).Cardinality())
				_ = _rhs319
				var _rhs320 _dafny.Int = (_this).Fm1((_this).F9(), globalState)
				_ = _rhs320
				var _rhs321 _dafny.Sequence = _dafny.SeqOf(_dafny.IntOfInt64(13))
				_ = _rhs321
				_1985_v107 = _rhs318
				_1845_v0 = _rhs319
				_1845_v0 = _rhs320
				_1962_v84 = _rhs321
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1963_v85), 0))
				_ = _index325
				(_1963_v85).ArraySet1((_1917_i5).Minus(_1845_v0), (_index325).Int())
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_1963_v85), 0))
				_ = _index326
				(_1963_v85).ArraySet1((_dafny.Zero).Minus(_1845_v0), (_index326).Int())
			}
		}
		var _1989_v108 _dafny.Sequence
		_ = _1989_v108
		_1989_v108 = _dafny.SeqOf((_this).F10())
		var _1990_v109 _dafny.Sequence
		_ = _1990_v109
		_1990_v109 = _dafny.SeqOf(((_this).Fm0(_1989_v108, (_this).F10(), _1845_v0, (_this).F10(), globalState)) || ((_this).F10()))
		var _1991_v110 D15
		_ = _1991_v110
		_1991_v110 = Companion_D15_.Create_DC39_(!((_this).F10()), true)
		var _pat_let_tv47 = _1989_v108
		_ = _pat_let_tv47
		var _pat_let_tv48 = _1990_v109
		_ = _pat_let_tv48
		var _pat_let_tv49 = _1990_v109
		_ = _pat_let_tv49
		_1990_v109 = func(_source27 D15) _dafny.Sequence {
			if _source27.Is_DC38() {
				var _1992___mcc_h6 bool = _source27.Get_().(D15_DC38).Cf69
				_ = _1992___mcc_h6
				var _1993___mcc_h7 _dafny.Sequence = _source27.Get_().(D15_DC38).Cf70
				_ = _1993___mcc_h7
				var _1994___mcc_h8 bool = _source27.Get_().(D15_DC38).Cf71
				_ = _1994___mcc_h8
				var _1995___mcc_h9 bool = _source27.Get_().(D15_DC38).Cf72
				_ = _1995___mcc_h9
				var _1996___mcc_h10 _dafny.CodePoint = _source27.Get_().(D15_DC38).Cf73
				_ = _1996___mcc_h10
				var _1997_cf73 _dafny.CodePoint = _1996___mcc_h10
				_ = _1997_cf73
				var _1998_cf72 bool = _1995___mcc_h9
				_ = _1998_cf72
				var _1999_cf71 bool = _1994___mcc_h8
				_ = _1999_cf71
				var _2000_cf70 _dafny.Sequence = _1993___mcc_h7
				_ = _2000_cf70
				var _2001_cf69 bool = _1992___mcc_h6
				_ = _2001_cf69
				return (Companion_D8_.Create_DC23_(_pat_let_tv47)).Dtor_cf41()
			} else if _source27.Is_DC39() {
				var _2002___mcc_h11 bool = _source27.Get_().(D15_DC39).Cf74
				_ = _2002___mcc_h11
				var _2003___mcc_h12 bool = _source27.Get_().(D15_DC39).Cf75
				_ = _2003___mcc_h12
				var _2004_cf75 bool = _2003___mcc_h12
				_ = _2004_cf75
				var _2005_cf74 bool = _2002___mcc_h11
				_ = _2005_cf74
				return _pat_let_tv48
			} else {
				var _2006___mcc_h13 _dafny.Array = _source27.Get_().(D15_DC37).Cf68
				_ = _2006___mcc_h13
				var _2007_cf68 _dafny.Array = _2006___mcc_h13
				_ = _2007_cf68
				return _pat_let_tv49
			}
		}((func() D15 {
			if (_this).F10() {
				return _1991_v110
			}
			return Companion_D15_.Create_DC39_((_this).F10(), (_this).F10())
		})())
		var _2008_v111 _dafny.Array
		_ = _2008_v111
		var _nw306 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw306
		_2008_v111 = _nw306
		for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2008_v111), 0))); ; {
			_guard_loop_8, _ok83 := _iter83()
			if !_ok83 {
				break
			}
			var _2009_i11 _dafny.Int
			_2009_i11 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_2009_i11).Sign() != -1) && ((_2009_i11).Cmp(_dafny.ArrayLen((_2008_v111), 0)) < 0)) {
				(_2008_v111).ArraySet1(((_this).F10()) || ((_this).F10()), (_2009_i11).Int())
			}
		}
		var _2010_i12 _dafny.Int
		_ = _2010_i12
		_2010_i12 = _dafny.Zero
		{
			for (_this).F10() {
				{
					if (_2010_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_2010_i12 = (_2010_i12).Plus(_dafny.One)
					var _2011_v112 *C2
					_ = _2011_v112
					var _nw307 *C2 = New_C2_()
					_ = _nw307
					_nw307.Ctor__((_this).F10(), _2008_v111, (_this).F9(), (_this).F10())
					_2011_v112 = _nw307
					var _2012_v113 _dafny.CodePoint
					_ = _2012_v113
					_2012_v113 = _dafny.CodePoint('q')
					_1845_v0 = ((_dafny.Zero).Minus((_1845_v0).Times(_1845_v0))).Plus((func() _dafny.Int {
						if (_this).F10() {
							return (_this).Fm3(_dafny.IntOfUint32((_dafny.SeqOf(_2011_v112)).Cardinality()), ((_this).F11()).Cardinality(), _2012_v113, _1845_v0, globalState)
						}
						return _1845_v0
					})())
					var _2013_v114 _dafny.Array
					_ = _2013_v114
					var _nw308 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
					_ = _nw308
					_2013_v114 = _nw308
					var _2014_v115 _dafny.Sequence
					_ = _2014_v115
					_2014_v115 = _dafny.UnicodeSeqOfUtf8Bytes("fie")
					var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2013_v114), 0))
					_ = _index327
					(_2013_v114).ArraySet1(_2014_v115, (_index327).Int())
					var _2015_v116 D14
					_ = _2015_v116
					_2015_v116 = Companion_D14_.Create_DC36_((_this).F10(), _2014_v115, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(425))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg109 _dafny.Int) interface{} {
							return coer109(arg109)
						}
					}((func(_2016_v113 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2017_i13 _dafny.Int) _dafny.CodePoint {
							return _2016_v113
						}
					})(_2012_v113))))
					var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2013_v114), 0))
					_ = _index328
					(_2013_v114).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2014_v115, (func() _dafny.Sequence {
						if true {
							return Companion_Default___.Fm20(_1845_v0, _2011_v112.F14, !((_this).F10()), _1845_v0, globalState)
						}
						return (_2015_v116).Dtor_cf66()
					})()), (_index328).Int())
					_1845_v0 = _1845_v0
					if (_this).F10() {
						_1845_v0 = _dafny.IntOfInt64(609)
						var _2018_v117 *C4
						_ = _2018_v117
						var _nw309 *C4 = New_C4_()
						_ = _nw309
						_nw309.Ctor__((_this).F9(), true)
						_2018_v117 = _nw309
						var _2019_v118 _dafny.Map
						_ = _2019_v118
						_2019_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2011_v112.F14, _2018_v117)
						var _2020_v119 _dafny.MultiSet
						_ = _2020_v119
						_2020_v119 = _dafny.MultiSetOf(_1845_v0, (_2019_v118).Cardinality())
						var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen(((_2011_v112).F15()), 0))
						_ = _index329
						((_2011_v112).F15()).ArraySet1((_dafny.MultiSetOf(_1845_v0, _1845_v0)).IsProperSubsetOf(_2020_v119), (_index329).Int())
						var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen(((_2011_v112).F15()), 0))
						_ = _index330
						((_2011_v112).F15()).ArraySet1(true, (_index330).Int())
						var _2021_v120 _dafny.Array
						_ = _2021_v120
						var _len0_54 _dafny.Int = _dafny.IntOfInt64(22)
						_ = _len0_54
						var _nw310 _dafny.Array
						_ = _nw310
						if _len0_54.Cmp(_dafny.Zero) == 0 {
							_nw310 = _dafny.NewArray(_len0_54)
						} else {
							var _init54 func(_dafny.Int) _dafny.Int = (func(_2022_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_2023_i14 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeDivisionInt(_2023_i14, _2022_v0)
								}
							})(_1845_v0)
							_ = _init54
							var _element0_54 = _init54(_dafny.Zero)
							_ = _element0_54
							_nw310 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
							(_nw310).ArraySet1(_element0_54, 0)
							var _nativeLen0_54 = (_len0_54).Int()
							_ = _nativeLen0_54
							for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
								(_nw310).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
							}
						}
						_2021_v120 = _nw310
						var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2021_v120), 0))
						_ = _index331
						(_2021_v120).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}((func(_2024_v113 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2025_i15 _dafny.Int) _dafny.CodePoint {
								return _2024_v113
							}
						})(_2012_v113)))).Cardinality())), (_index331).Int())
						var _2026_v121 _dafny.Array
						_ = _2026_v121
						var _len0_55 _dafny.Int = _dafny.IntOfInt64(3)
						_ = _len0_55
						var _nw311 _dafny.Array
						_ = _nw311
						if _len0_55.Cmp(_dafny.Zero) == 0 {
							_nw311 = _dafny.NewArray(_len0_55)
						} else {
							var _init55 func(_dafny.Int) _dafny.Map = (func(_2027_v0 _dafny.Int, _2028_v1 _dafny.Set, _2029_v112 *C2, _2030_v113 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
								return func(_2031_i16 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D4_.Create_DC11_(_2027_v0, _2028_v1, _2029_v112.F14), Companion_D4_.Create_DC11_(_2027_v0, _2028_v1, ((_2029_v112).F15()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen(((_2029_v112).F15()), 0))).Int()).(bool)), Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(847))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg111 _dafny.Int) interface{} {
											return coer111(arg111)
										}
									}((func(_2032_v113 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
										return func(_2033_i17 _dafny.Int) _dafny.CodePoint {
											return _2032_v113
										}
									})(_2030_v113)))).Cardinality()), _2028_v1, true)), _2027_v0))
								}
							})(_1845_v0, _1846_v1, _2011_v112, _2012_v113)
							_ = _init55
							var _element0_55 = _init55(_dafny.Zero)
							_ = _element0_55
							_nw311 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
							(_nw311).ArraySet1(_element0_55, 0)
							var _nativeLen0_55 = (_len0_55).Int()
							_ = _nativeLen0_55
							for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
								(_nw311).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
							}
						}
						_2026_v121 = _nw311
						var _2034_v122 _dafny.Sequence
						_ = _2034_v122
						_2034_v122 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-656))).Uint32(), func(coer112 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
							return func(arg112 _dafny.Int) interface{} {
								return coer112(arg112)
							}
						}((func(_2035_v0 _dafny.Int, _2036_v1 _dafny.Set, _2037_v112 *C2) func(_dafny.Int) D4 {
							return func(_2038_i18 _dafny.Int) D4 {
								return Companion_D4_.Create_DC11_(_2035_v0, _2036_v1, _2037_v112.F14)
							}
						})(_1845_v0, _1846_v1, _2011_v112)))
						var _2039_v123 _dafny.Map
						_ = _2039_v123
						_2039_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2034_v122, _1845_v0)
						var _2040_v124 _dafny.Map
						_ = _2040_v124
						_2040_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1845_v0, _2039_v123)
						var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2026_v121), 0))
						_ = _index332
						(_2026_v121).ArraySet1(_2040_v124, (_index332).Int())
						var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2021_v120), 0))
						_ = _index333
						var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2026_v121), 0))
						_ = _index334
						var _rhs322 _dafny.Int = (_1845_v0).Plus(_1845_v0)
						_ = _rhs322
						var _rhs323 _dafny.CodePoint = _2012_v113
						_ = _rhs323
						var _rhs324 _dafny.Map = (_2040_v124).Merge(_2040_v124)
						_ = _rhs324
						var _lhs231 _dafny.Array = _2021_v120
						_ = _lhs231
						var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2021_v120), 0))
						_ = _lhs232
						var _lhs233 _dafny.Array = _2026_v121
						_ = _lhs233
						var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2026_v121), 0))
						_ = _lhs234
						(_lhs231).ArraySet1(_rhs322, (_lhs232).Int())
						_2012_v113 = _rhs323
						(_lhs233).ArraySet1(_rhs324, (_lhs234).Int())
						var _2041_v125 _dafny.Map
						_ = _2041_v125
						_2041_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(313), _1845_v0)
						var _2042_v126 _dafny.Sequence
						_ = _2042_v126
						_2042_v126 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Int {
							if (_2041_v125).Contains((_2021_v120).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2021_v120), 0))).Int()).(_dafny.Int)) {
								return (_2041_v125).Get((_2021_v120).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_2021_v120), 0))).Int()).(_dafny.Int)).(_dafny.Int)
							}
							return _1845_v0
						})(), _1845_v0)).Cardinality()))
						var _2043_v127 _dafny.Map
						_ = _2043_v127
						_2043_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2042_v126).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1989_v108).Cardinality()), _dafny.IntOfUint32((_2042_v126).Cardinality()))).Uint32()).(_dafny.Int), true)
						(_2011_v112).F14 = !((_2043_v127).Merge(_2043_v127)).Equals(_2043_v127)
						var _2044_v128 D6
						_ = _2044_v128
						_2044_v128 = Companion_D6_.Create_DC19_((_this).F10(), _2012_v113)
						var _2045_v129 _dafny.MultiSet
						_ = _2045_v129
						_2045_v129 = _dafny.MultiSetOf(_2044_v128, _2044_v128, _2044_v128)
						var _2046_v130 _dafny.Sequence
						_ = _2046_v130
						_2046_v130 = _dafny.SeqOf(_2045_v129)
						var _rhs325 _dafny.Sequence = (_2013_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2013_v114), 0))).Int()).(_dafny.Sequence)
						_ = _rhs325
						var _rhs326 _dafny.MultiSet = (_2046_v130).Select((Companion_Default___.SafeIndex(_1845_v0, _dafny.IntOfUint32((_2046_v130).Cardinality()))).Uint32()).(_dafny.MultiSet)
						_ = _rhs326
						_2014_v115 = _rhs325
						_2045_v129 = _rhs326
					} else {
						var _2047_v131 _dafny.Array
						_ = _2047_v131
						var _nw312 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(23))
						_ = _nw312
						_2047_v131 = _nw312
						var _2048_v132 _dafny.Set
						_ = _2048_v132
						_2048_v132 = _dafny.SetOf(_2047_v131, _2047_v131, _2047_v131, _2047_v131)
						_2048_v132 = _2048_v132
						var _2049_v133 *C5
						_ = _2049_v133
						var _nw313 *C5 = New_C5_()
						_ = _nw313
						_nw313.Ctor__(Companion_Default___.Fm48(_2011_v112.F14, globalState), _2011_v112.F14)
						_2049_v133 = _nw313
						var _2050_v134 _dafny.Array
						_ = _2050_v134
						var _nw314 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
						_ = _nw314
						_2050_v134 = _nw314
						_2050_v134 = _2050_v134
						var _2051_v135 _dafny.Set
						_ = _2051_v135
						_2051_v135 = _dafny.SetOf((_2013_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2013_v114), 0))).Int()).(_dafny.Sequence))
						_2051_v135 = (_2051_v135).Difference(_2051_v135)
						var _2052_v136 _dafny.Array
						_ = _2052_v136
						var _nw315 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
						_ = _nw315
						_2052_v136 = _nw315
						var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_2052_v136), 0))
						_ = _index335
						(_2052_v136).ArraySet1CodePoint(_dafny.CodePoint('a'), (_index335).Int())
						var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_2052_v136), 0))
						_ = _index336
						(_2052_v136).ArraySet1CodePoint(_2012_v113, (_index336).Int())
					}
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		r1 = (_this).F10()
		var _2053_v137 _dafny.Sequence
		_ = _2053_v137
		_2053_v137 = _dafny.UnicodeSeqOfUtf8Bytes("y")
		r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2053_v137, _2053_v137), _2053_v137)
		var _2054_v138 _dafny.Sequence
		_ = _2054_v138
		_2054_v138 = _dafny.SeqOf(_1845_v0)
		var _2055_v139 _dafny.CodePoint
		_ = _2055_v139
		_2055_v139 = _dafny.CodePoint('f')
		var _2056_v140 _dafny.Map
		_ = _2056_v140
		_2056_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1845_v0, (_this).F10())
		var _2057_v141 D14
		_ = _2057_v141
		_2057_v141 = Companion_D14_.Create_DC35_((_2056_v140).Cardinality(), (_this).F10(), _2055_v139)
		r1 = (((_dafny.Zero).Minus(_1845_v0)).Plus((_this).Fm3(_dafny.IntOfUint32((_2054_v138).Cardinality()), _dafny.IntOfUint32((_1990_v109).Cardinality()), _2055_v139, _1845_v0, globalState))).Cmp((_dafny.Zero).Minus((_2057_v141).Dtor_cf62())) != 0
		var _2058_v142 D7
		_ = _2058_v142
		_2058_v142 = Companion_D7_.Create_DC21_((_this).F11())
		var _2059_v143 _dafny.Sequence
		_ = _2059_v143
		_2059_v143 = _dafny.SeqOf(_dafny.SetOf((_this).F10(), (_this).F10(), (_this).F10()), (_this).F11(), (func(_pat_let50_0 D7) D7 {
			return func(_2060_dt__update__tmp_h3 D7) D7 {
				return func(_pat_let51_0 _dafny.Set) D7 {
					return func(_2061_dt__update_hcf40_h0 _dafny.Set) D7 {
						return Companion_D7_.Create_DC21_(_2061_dt__update_hcf40_h0)
					}(_pat_let51_0)
				}(_dafny.SetOf((_this).F10()))
			}(_pat_let50_0)
		}(_2058_v142)).Dtor_cf40())
		var _2062_v144 _dafny.Sequence
		_ = _2062_v144
		_2062_v144 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg113 _dafny.Int) interface{} {
				return coer113(arg113)
			}
		}(func(_2063_i19 _dafny.Int) _dafny.Set {
			return (_this).F11()
		})), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2059_v143, _2059_v143), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F9())).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2059_v143, _2059_v143)).Cardinality()))).Uint32(), (_this).F11()), _2059_v143, _dafny.SeqOf((_this).F11(), (_this).F11()), _dafny.SeqOf((_this).F11(), _dafny.SetOf((_this).F10()), (_this).F11(), (_this).F11()))
		r2 = (_2062_v144).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_1845_v0, _dafny.IntOfUint32((_2054_v138).Cardinality())), _dafny.IntOfUint32((_2062_v144).Cardinality()))).Uint32()).(_dafny.Sequence)
		r3 = _dafny.UnicodeSeqOfUtf8Bytes("xjnxgcyay")
		return r0, r1, r2, r3
	}
}
func (_this *C7) M1(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Map, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var _2064_v0 _dafny.Array
		_ = _2064_v0
		var _nw316 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
		_ = _nw316
		_2064_v0 = _nw316
		for _iter84 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2064_v0), 0))); ; {
			_guard_loop_9, _ok84 := _iter84()
			if !_ok84 {
				break
			}
			var _2065_i0 _dafny.Int
			_2065_i0 = interface{}(_guard_loop_9).(_dafny.Int)
			if (true) && (((_2065_i0).Sign() != -1) && ((_2065_i0).Cmp(_dafny.ArrayLen((_2064_v0), 0)) < 0)) {
				(_2064_v0).ArraySet1CodePoint(p0, (_2065_i0).Int())
			}
		}
		if (_this).F10() {
			var _2066_v1 _dafny.Array
			_ = _2066_v1
			var _nw317 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw317
			_2066_v1 = _nw317
			_2066_v1 = _2066_v1
			var _2067_v2 _dafny.CodePoint
			_ = _2067_v2
			_2067_v2 = _dafny.CodePoint('y')
			_2067_v2 = p0
			r1 = (func() bool {
				if (_this).F10() {
					return (_this).F10()
				}
				return true
			})()
			var _2068_v3 _dafny.Array
			_ = _2068_v3
			var _nw318 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw318
			_2068_v3 = _nw318
			_2068_v3 = _2068_v3
			var _2069_v4 _dafny.Int
			_ = _2069_v4
			_2069_v4 = _dafny.IntOfInt64(128)
			_2069_v4 = (func() _dafny.Int {
				if (_2069_v4).Cmp(_2069_v4) != 0 {
					return _dafny.IntOfInt64(662)
				}
				return _2069_v4
			})()
		} else {
			var _2070_v5 _dafny.Int
			_ = _2070_v5
			_2070_v5 = _dafny.IntOfInt64(817)
			_2070_v5 = _2070_v5
			if (_this).F10() {
				var _2071_v6 _dafny.Array
				_ = _2071_v6
				var _len0_56 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_56
				var _nw319 _dafny.Array
				_ = _nw319
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw319 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) bool = func(_2072_i1 _dafny.Int) bool {
						return (_this).F10()
					}
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw319 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw319).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw319).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2071_v6 = _nw319
				var _2073_v7 _dafny.Array
				_ = _2073_v7
				var _nwElement0_66 _dafny.Array = _2071_v6
				_ = _nwElement0_66
				var _nw320 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(2))
				_ = _nw320
				(_nw320).ArraySet1(_nwElement0_66, 0)
				(_nw320).ArraySet1(_2071_v6, 1)
				_2073_v7 = _nw320
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2071_v6), 0))
				_ = _index337
				(_2071_v6).ArraySet1((_this).F10(), (_index337).Int())
				var _2074_v8 _dafny.Sequence
				_ = _2074_v8
				_2074_v8 = _dafny.UnicodeSeqOfUtf8Bytes("gfnvwqdiy")
				var _2075_v9 D14
				_ = _2075_v9
				_2075_v9 = Companion_D14_.Create_DC36_(!((_this).F10()), _2074_v8, _2074_v8)
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2071_v6), 0))
				_ = _index338
				var _rhs327 _dafny.Array = _2073_v7
				_ = _rhs327
				var _rhs328 bool = (_2075_v9).Dtor_cf65()
				_ = _rhs328
				var _rhs329 bool = !((_this).F10())
				_ = _rhs329
				var _rhs330 _dafny.Int = _dafny.IntOfInt64(-820)
				_ = _rhs330
				var _lhs235 _dafny.Array = _2071_v6
				_ = _lhs235
				var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2071_v6), 0))
				_ = _lhs236
				var _lhs237 *GlobalState = globalState
				_ = _lhs237
				_2073_v7 = _rhs327
				(_lhs235).ArraySet1(_rhs328, (_lhs236).Int())
				_lhs237.F8 = _rhs329
				_2070_v5 = _rhs330
				_2070_v5 = _dafny.IntOfInt64(557)
				r1 = (_this).F10()
				var _2076_v10 _dafny.Sequence
				_ = _2076_v10
				_2076_v10 = _dafny.SeqOf((_2071_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_2071_v6), 0))).Int()).(bool))
				(globalState).F8 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2070_v5, _dafny.Companion_Sequence_.Update(_2076_v10, (Companion_Default___.SafeIndex(_2070_v5, _dafny.IntOfUint32((_2076_v10).Cardinality()))).Uint32(), (_this).Fm2(globalState)))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm1((_this).F9(), globalState), _dafny.Companion_Sequence_.Update(_2076_v10, (Companion_Default___.SafeIndex(_2070_v5, _dafny.IntOfUint32((_2076_v10).Cardinality()))).Uint32(), (_this).F10())))
				(globalState).F8 = (Companion_D15_.Create_DC39_((_this).F10(), true)).Dtor_cf74()
			} else {
				_2070_v5 = _2070_v5
				(globalState).F8 = !((_this).F10())
				(globalState).F8 = (_this).F10()
				var _2077_v11 _dafny.Array
				_ = _2077_v11
				var _nw321 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw321
				_2077_v11 = _nw321
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_2077_v11), 0))
				_ = _index339
				(_2077_v11).ArraySet1((_dafny.Zero).Minus((_2070_v5).Minus(_2070_v5)), (_index339).Int())
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_2077_v11), 0))
				_ = _index340
				(_2077_v11).ArraySet1(_dafny.IntOfInt64(224), (_index340).Int())
				var _2078_v12 _dafny.Sequence
				_ = _2078_v12
				_2078_v12 = _dafny.UnicodeSeqOfUtf8Bytes("uppopviv")
				(globalState).F8 = !((Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2078_v12).Cardinality()), (_dafny.SetOf((_this).F10(), !((_this).F10()))).Cardinality())).Cmp(_dafny.IntOfUint32((_2078_v12).Cardinality())) < 0)
			}
			_2070_v5 = Companion_Default___.SafeModuloInt(_2070_v5, _2070_v5)
			var _rhs331 _dafny.Int = _2070_v5
			_ = _rhs331
			var _rhs332 _dafny.Int = _2070_v5
			_ = _rhs332
			var _rhs333 bool = (_this).F10()
			_ = _rhs333
			var _lhs238 *GlobalState = globalState
			_ = _lhs238
			_2070_v5 = _rhs331
			_2070_v5 = _rhs332
			_lhs238.F8 = _rhs333
			var _2079_v13 _dafny.Sequence
			_ = _2079_v13
			_2079_v13 = _dafny.SeqOf((_this).F10(), true)
			if !(!(!(_dafny.Companion_Sequence_.IsPrefixOf(_2079_v13, _dafny.SeqOf((_this).F10()))))) || ((_this).F10()) {
				var _2080_v14 _dafny.Array
				_ = _2080_v14
				var _len0_57 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_57
				var _nw322 _dafny.Array
				_ = _nw322
				if _len0_57.Cmp(_dafny.Zero) == 0 {
					_nw322 = _dafny.NewArray(_len0_57)
				} else {
					var _init57 func(_dafny.Int) _dafny.Int = func(_2081_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_2081_i2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eqk")).Cardinality()))
					}
					_ = _init57
					var _element0_57 = _init57(_dafny.Zero)
					_ = _element0_57
					_nw322 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
					(_nw322).ArraySet1(_element0_57, 0)
					var _nativeLen0_57 = (_len0_57).Int()
					_ = _nativeLen0_57
					for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
						(_nw322).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
					}
				}
				_2080_v14 = _nw322
				_2080_v14 = _2080_v14
				var _2082_v15 D0
				_ = _2082_v15
				_2082_v15 = Companion_D0_.Create_DC0_((_this).F10())
				var _2083_v16 _dafny.Map
				_ = _2083_v16
				_2083_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2082_v15).Dtor_cf0()), _2070_v5)
				var _2084_v17 _dafny.Map
				_ = _2084_v17
				_2084_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2070_v5, true)
				_2083_v16 = (_2083_v16).Update(!((func() bool {
					if (_2084_v17).Contains(_2070_v5) {
						return (_2084_v17).Get(_2070_v5).(bool)
					}
					return (_this).F10()
				})()), Companion_Default___.Fm33((_this).F10(), (_2079_v13).Select((Companion_Default___.SafeIndex(_2070_v5, _dafny.IntOfUint32((_2079_v13).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(704), !((_this).F10()), globalState))
				var _2085_v18 *C3
				_ = _2085_v18
				var _nw323 *C3 = New_C3_()
				_ = _nw323
				_nw323.Ctor__()
				_2085_v18 = _nw323
				var _2086_v19 *C4
				_ = _2086_v19
				var _nw324 *C4 = New_C4_()
				_ = _nw324
				_nw324.Ctor__((_this).F9(), (_this).F10())
				_2086_v19 = _nw324
				var _2087_v20 _dafny.Map
				_ = _2087_v20
				_2087_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2086_v19)
				var _2088_v21 _dafny.Map
				_ = _2088_v21
				_2088_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2085_v18, _2087_v20)
				_2088_v21 = (_2088_v21).Update(_2085_v18, _2087_v20)
				var _2089_v22 _dafny.CodePoint
				_ = _2089_v22
				_2089_v22 = _dafny.CodePoint('r')
				_2089_v22 = _2089_v22
				var _2090_v23 _dafny.Sequence
				_ = _2090_v23
				_2090_v23 = _dafny.UnicodeSeqOfUtf8Bytes("sta")
				var _2091_v24 _dafny.Sequence
				_ = _2091_v24
				_2091_v24 = _dafny.SeqOf(_2090_v23)
				var _2092_v25 _dafny.CodePoint
				_ = _2092_v25
				_2092_v25 = _2089_v22
				var _2093_v26 D15
				_ = _2093_v26
				_2093_v26 = Companion_D15_.Create_DC38_(!((_this).F10()), (_2091_v24).Select((Companion_Default___.SafeIndex(_2070_v5, _dafny.IntOfUint32((_2091_v24).Cardinality()))).Uint32()).(_dafny.Sequence), true, (_this).F10(), _2092_v25)
				var _2094_v27 _dafny.Sequence
				_ = _2094_v27
				_2094_v27 = _dafny.SeqOf(_2084_v17, _2084_v17, _2084_v17, _2084_v17)
				var _2095_v28 _dafny.Array
				_ = _2095_v28
				var _nw325 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw325
				_2095_v28 = _nw325
				var _2096_v29 _dafny.Map
				_ = _2096_v29
				_2096_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2095_v28)
				var _2097_v30 *C2
				_ = _2097_v30
				var _nw326 *C2 = New_C2_()
				_ = _nw326
				_nw326.Ctor__((_this).F10(), (func() _dafny.Array {
					if (_2096_v29).Contains((_this).F10()) {
						return (_2096_v29).Get((_this).F10()).(_dafny.Array)
					}
					return _2095_v28
				})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10()), true)
				_2097_v30 = _nw326
				var _2098_v31 _dafny.MultiSet
				_ = _2098_v31
				_2098_v31 = _dafny.MultiSetOf(_2097_v30)
				var _2099_v32 D4
				_ = _2099_v32
				_2099_v32 = Companion_D4_.Create_DC12_((_this).F10(), _2070_v5, Companion_Default___.Fm31((_this).F10(), (_2094_v27).Select((Companion_Default___.SafeIndex((_2098_v31).Cardinality(), _dafny.IntOfUint32((_2094_v27).Cardinality()))).Uint32()).(_dafny.Map), globalState), (_dafny.Zero).Minus(_2070_v5), _dafny.IntOfUint32((_2090_v23).Cardinality()))
				var _2100_v33 _dafny.Map
				_ = _2100_v33
				_2100_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2093_v26, _2099_v32)
				var _2101_v34 D6
				_ = _2101_v34
				_2101_v34 = Companion_D6_.Create_DC18_(p0, _2070_v5, _2070_v5)
				var _2102_v35 D6
				_ = _2102_v35
				_2102_v35 = Companion_D6_.Create_DC20_(_2101_v34)
				var _2103_v36 _dafny.Sequence
				_ = _2103_v36
				_2103_v36 = _dafny.SeqOf(Companion_Default___.Fm55((_this).F10(), globalState))
				var _2104_v37 _dafny.Array
				_ = _2104_v37
				var _nwElement0_67 bool = (_this).F10()
				_ = _nwElement0_67
				var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(22))
				_ = _nw327
				(_nw327).ArraySet1(_nwElement0_67, 0)
				(_nw327).ArraySet1((_this).F10(), 1)
				(_nw327).ArraySet1((_this).F10(), 2)
				(_nw327).ArraySet1(!((_this).F10()), 3)
				(_nw327).ArraySet1((_this).F10(), 4)
				(_nw327).ArraySet1(!((_this).F10()), 5)
				(_nw327).ArraySet1((_this).F10(), 6)
				(_nw327).ArraySet1((_this).F10(), 7)
				(_nw327).ArraySet1((_this).F10(), 8)
				(_nw327).ArraySet1(!(_2100_v33).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2093_v26, _2099_v32)), 9)
				(_nw327).ArraySet1((_this).F10(), 10)
				(_nw327).ArraySet1((_this).F10(), 11)
				(_nw327).ArraySet1((_2070_v5).Cmp(((_this).F9()).Cardinality()) > 0, 12)
				(_nw327).ArraySet1(!((_2097_v30.F14) || ((_this).F10())), 13)
				(_nw327).ArraySet1((_this).F10(), 14)
				(_nw327).ArraySet1((_2097_v30.F14) || (_2097_v30.F14), 15)
				(_nw327).ArraySet1(!_dafny.Companion_Sequence_.Contains(_2103_v36, Companion_D6_.Create_DC20_(_2102_v35)), 16)
				(_nw327).ArraySet1((_this).F10(), 17)
				(_nw327).ArraySet1(!(!((_dafny.MultiSetFromSeq(_2079_v13)).Update(_2097_v30.F14, Companion_Default___.Abs(_2070_v5))).Contains(false)), 18)
				(_nw327).ArraySet1(_2097_v30.F14, 19)
				(_nw327).ArraySet1((_this).F10(), 20)
				(_nw327).ArraySet1(_2097_v30.F14, 21)
				_2104_v37 = _nw327
				var _2105_v38 _dafny.MultiSet
				_ = _2105_v38
				_2105_v38 = _dafny.MultiSetOf(_2070_v5)
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_2104_v37), 0))
				_ = _index341
				(_2104_v37).ArraySet1((_2105_v38).IsDisjointFrom(_2105_v38), (_index341).Int())
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_2104_v37), 0))
				_ = _index342
				(_2104_v37).ArraySet1((_this).Fm0(Companion_Default___.Fm4(_dafny.CodePoint('o'), globalState), _2097_v30.F14, _2070_v5, (_this).F10(), globalState), (_index342).Int())
			} else {
				var _2106_v39 _dafny.Array
				_ = _2106_v39
				var _nw328 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw328
				_2106_v39 = _nw328
				var _2107_v40 _dafny.Array
				_ = _2107_v40
				var _len0_58 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_58
				var _nw329 _dafny.Array
				_ = _nw329
				if _len0_58.Cmp(_dafny.Zero) == 0 {
					_nw329 = _dafny.NewArray(_len0_58)
				} else {
					var _init58 func(_dafny.Int) _dafny.Int = (func(_2108_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2109_i3 _dafny.Int) _dafny.Int {
							return (_2109_i3).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2108_v5, (_this).F10())).Cardinality())
						}
					})(_2070_v5)
					_ = _init58
					var _element0_58 = _init58(_dafny.Zero)
					_ = _element0_58
					_nw329 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
					(_nw329).ArraySet1(_element0_58, 0)
					var _nativeLen0_58 = (_len0_58).Int()
					_ = _nativeLen0_58
					for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
						(_nw329).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
					}
				}
				_2107_v40 = _nw329
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))
				_ = _index343
				(_2106_v39).ArraySet1(_2107_v40, (_index343).Int())
				var _2110_v41 _dafny.Map
				_ = _2110_v41
				_2110_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2107_v40)
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))
				_ = _index344
				(_2106_v39).ArraySet1((func() _dafny.Array {
					if (_2110_v41).Contains((_this).F10()) {
						return (_2110_v41).Get((_this).F10()).(_dafny.Array)
					}
					return _2107_v40
				})(), (_index344).Int())
				var _2111_v42 D6
				_ = _2111_v42
				_2111_v42 = Companion_D6_.Create_DC18_(p0, _2070_v5, ((_this).F9()).Cardinality())
				var _arr4 _dafny.Array = _dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))
				_ = _arr4
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))), 0))
				_ = _index345
				(_arr4).ArraySet1((_2111_v42).Dtor_cf35(), (_index345).Int())
				var _2112_v43 *C4
				_ = _2112_v43
				var _nw330 *C4 = New_C4_()
				_ = _nw330
				_nw330.Ctor__((_this).F9(), (_this).F10())
				_2112_v43 = _nw330
				var _2113_v44 _dafny.Map
				_ = _2113_v44
				_2113_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2070_v5, _2112_v43)
				var _2114_v45 _dafny.Array
				_ = _2114_v45
				var _nwElement0_68 *C4 = _2112_v43
				_ = _nwElement0_68
				var _nw331 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(16))
				_ = _nw331
				(_nw331).ArraySet1(_nwElement0_68, 0)
				(_nw331).ArraySet1(_2112_v43, 1)
				(_nw331).ArraySet1(_2112_v43, 2)
				(_nw331).ArraySet1(_2112_v43, 3)
				(_nw331).ArraySet1(_2112_v43, 4)
				(_nw331).ArraySet1(_2112_v43, 5)
				(_nw331).ArraySet1(_2112_v43, 6)
				(_nw331).ArraySet1(_2112_v43, 7)
				(_nw331).ArraySet1(_2112_v43, 8)
				(_nw331).ArraySet1((func() *C4 {
					if (_2113_v44).Contains(_2070_v5) {
						return (_2113_v44).Get(_2070_v5).(*C4)
					}
					return _2112_v43
				})(), 9)
				(_nw331).ArraySet1(_2112_v43, 10)
				(_nw331).ArraySet1(_2112_v43, 11)
				(_nw331).ArraySet1(_2112_v43, 12)
				(_nw331).ArraySet1(_2112_v43, 13)
				(_nw331).ArraySet1(_2112_v43, 14)
				(_nw331).ArraySet1(_2112_v43, 15)
				_2114_v45 = _nw331
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_2114_v45), 0))
				_ = _index346
				(_2114_v45).ArraySet1(_2112_v43, (_index346).Int())
				var _arr5 _dafny.Array = _dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))
				_ = _arr5
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))), 0))
				_ = _index347
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_2114_v45), 0))
				_ = _index348
				var _rhs334 bool = (_this).F10()
				_ = _rhs334
				var _rhs335 _dafny.Int = _2070_v5
				_ = _rhs335
				var _rhs336 *C4 = _2112_v43
				_ = _rhs336
				var _lhs239 *GlobalState = globalState
				_ = _lhs239
				var _lhs240 _dafny.Array = _dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))
				_ = _lhs240
				var _lhs241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))), 0))
				_ = _lhs241
				var _lhs242 _dafny.Array = _2114_v45
				_ = _lhs242
				var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_2114_v45), 0))
				_ = _lhs243
				_lhs239.F8 = _rhs334
				(_lhs240).ArraySet1(_rhs335, (_lhs241).Int())
				(_lhs242).ArraySet1(_rhs336, (_lhs243).Int())
				var _2115_v46 _dafny.Array
				_ = _2115_v46
				var _len0_59 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_59
				var _nw332 _dafny.Array
				_ = _nw332
				if _len0_59.Cmp(_dafny.Zero) == 0 {
					_nw332 = _dafny.NewArray(_len0_59)
				} else {
					var _init59 func(_dafny.Int) bool = func(_2116_i4 _dafny.Int) bool {
						return !((_this).F10())
					}
					_ = _init59
					var _element0_59 = _init59(_dafny.Zero)
					_ = _element0_59
					_nw332 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
					(_nw332).ArraySet1(_element0_59, 0)
					var _nativeLen0_59 = (_len0_59).Int()
					_ = _nativeLen0_59
					for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
						(_nw332).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
					}
				}
				_2115_v46 = _nw332
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2115_v46), 0))
				_ = _index349
				(_2115_v46).ArraySet1((_this).F10(), (_index349).Int())
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2115_v46), 0))
				_ = _index350
				(_2115_v46).ArraySet1((_dafny.IntOfInt64(210)).Cmp((_dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_dafny.ArrayCastTo((_2106_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2106_v39), 0))).Int()))), 0))).Int()).(_dafny.Int)) <= 0, (_index350).Int())
				r1 = (_2115_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_2115_v46), 0))).Int()).(bool)
				var _len0_60 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_60
				var _nw333 _dafny.Array
				_ = _nw333
				if _len0_60.Cmp(_dafny.Zero) == 0 {
					_nw333 = _dafny.NewArray(_len0_60)
				} else {
					var _init60 func(_dafny.Int) bool = (func(_2117_v39 _dafny.Array, _2118_v5 _dafny.Int) func(_dafny.Int) bool {
						return func(_2119_i5 _dafny.Int) bool {
							return !(func() _dafny.Map {
								var _coll75 = _dafny.NewMapBuilder()
								_ = _coll75
								for _iter85 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(239))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg114 _dafny.Int) interface{} {
										return coer114(arg114)
									}
								}((func(_2120_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
									return func(_2121_i6 _dafny.Int) _dafny.Int {
										return _2120_v5
									}
								})(_2118_v5)))).Elements()); ; {
									_compr_75, _ok85 := _iter85()
									if !_ok85 {
										break
									}
									var _2122_v47 _dafny.Int
									_2122_v47 = interface{}(_compr_75).(_dafny.Int)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(239))).Uint32(), func(coer115 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
										return func(arg115 _dafny.Int) interface{} {
											return coer115(arg115)
										}
									}((func(_2123_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
										return func(_2121_i6 _dafny.Int) _dafny.Int {
											return _2123_v5
										}
									})(_2118_v5))), _2122_v47) {
										_coll75.Add(Companion_Default___.SafeModuloInt(_2122_v47, (_dafny.ArrayCastTo((_2117_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2117_v39), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_dafny.ArrayCastTo((_2117_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2117_v39), 0))).Int()))), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cxlqhjcb")).Cardinality()))
									}
								}
								return _coll75.ToMap()
							}()).Contains(_dafny.IntOfInt64(671))
						}
					})(_2106_v39, _2070_v5)
					_ = _init60
					var _element0_60 = _init60(_dafny.Zero)
					_ = _element0_60
					_nw333 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
					(_nw333).ArraySet1(_element0_60, 0)
					var _nativeLen0_60 = (_len0_60).Int()
					_ = _nativeLen0_60
					for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
						(_nw333).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
					}
				}
				_2115_v46 = _nw333
			}
		}
		var _2124_v48 _dafny.Int
		_ = _2124_v48
		_2124_v48 = _dafny.IntOfInt64(-908)
		var _rhs337 _dafny.Int = _2124_v48
		_ = _rhs337
		var _rhs338 bool = (!((_this).Fm2(globalState)) || ((_this).F10())) || ((_this).F10())
		_ = _rhs338
		_2124_v48 = _rhs337
		r1 = _rhs338
		var _2125_v49 _dafny.Sequence
		_ = _2125_v49
		_2125_v49 = _dafny.SeqOf((_this).F10())
		if (_2125_v49).Select((Companion_Default___.SafeIndex(_2124_v48, _dafny.IntOfUint32((_2125_v49).Cardinality()))).Uint32()).(bool) {
			var _2126_v50 _dafny.Array
			_ = _2126_v50
			var _len0_61 _dafny.Int = _dafny.One
			_ = _len0_61
			var _nw334 _dafny.Array
			_ = _nw334
			if _len0_61.Cmp(_dafny.Zero) == 0 {
				_nw334 = _dafny.NewArray(_len0_61)
			} else {
				var _init61 func(_dafny.Int) bool = func(_2127_i7 _dafny.Int) bool {
					return (_this).F10()
				}
				_ = _init61
				var _element0_61 = _init61(_dafny.Zero)
				_ = _element0_61
				_nw334 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
				(_nw334).ArraySet1(_element0_61, 0)
				var _nativeLen0_61 = (_len0_61).Int()
				_ = _nativeLen0_61
				for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
					(_nw334).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
				}
			}
			_2126_v50 = _nw334
			var _2128_v51 _dafny.MultiSet
			_ = _2128_v51
			_2128_v51 = _dafny.MultiSetOf((_this).F10())
			var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))
			_ = _index351
			(_2126_v50).ArraySet1((_2128_v51).IsSubsetOf(_dafny.MultiSetFromSeq(_2125_v49)), (_index351).Int())
			var _2129_v52 _dafny.Array
			_ = _2129_v52
			var _len0_62 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_62
			var _nw335 _dafny.Array
			_ = _nw335
			if _len0_62.Cmp(_dafny.Zero) == 0 {
				_nw335 = _dafny.NewArray(_len0_62)
			} else {
				var _init62 func(_dafny.Int) _dafny.Map = (func(_2130_v49 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_2131_i8 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2130_v49)
					}
				})(_2125_v49)
				_ = _init62
				var _element0_62 = _init62(_dafny.Zero)
				_ = _element0_62
				_nw335 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
				(_nw335).ArraySet1(_element0_62, 0)
				var _nativeLen0_62 = (_len0_62).Int()
				_ = _nativeLen0_62
				for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
					(_nw335).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
				}
			}
			_2129_v52 = _nw335
			var _2132_v53 _dafny.Map
			_ = _2132_v53
			_2132_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2125_v49)
			var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_2129_v52), 0))
			_ = _index352
			(_2129_v52).ArraySet1((_2132_v53).Merge(_2132_v53), (_index352).Int())
			var _2133_v54 _dafny.Map
			_ = _2133_v54
			_2133_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_dafny.Zero).Minus(_2124_v48))
			var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))
			_ = _index353
			var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_2129_v52), 0))
			_ = _index354
			var _rhs339 bool = !((_this).F10()) || ((_this).F10())
			_ = _rhs339
			var _rhs340 bool = ((_this).F10()) == ((_2124_v48).Cmp((_2133_v54).Cardinality()) > 0)
			_ = _rhs340
			var _rhs341 _dafny.Map = ((_2132_v53).Merge((_2132_v53).Update(!((_this).Fm2(globalState)), _2125_v49))).Merge((_2132_v53).Update((_this).F10(), _2125_v49))
			_ = _rhs341
			var _rhs342 _dafny.Int = (_2124_v48).Minus(_2124_v48)
			_ = _rhs342
			var _lhs244 _dafny.Array = _2126_v50
			_ = _lhs244
			var _lhs245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))
			_ = _lhs245
			var _lhs246 _dafny.Array = _2129_v52
			_ = _lhs246
			var _lhs247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_2129_v52), 0))
			_ = _lhs247
			(_lhs244).ArraySet1(_rhs339, (_lhs245).Int())
			r1 = _rhs340
			(_lhs246).ArraySet1(_rhs341, (_lhs247).Int())
			_2124_v48 = _rhs342
			var _2134_v55 D6
			_ = _2134_v55
			_2134_v55 = Companion_D6_.Create_DC16_((_this).F9())
			var _2135_v56 _dafny.Array
			_ = _2135_v56
			var _nwElement0_69 _dafny.Map = (_this).F9()
			_ = _nwElement0_69
			var _nw336 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(15))
			_ = _nw336
			(_nw336).ArraySet1(_nwElement0_69, 0)
			(_nw336).ArraySet1(((_this).F9()).Merge((_this).F9()), 1)
			(_nw336).ArraySet1((Companion_Default___.Fm48((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool), globalState)).Update((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool), (_this).F10()), 2)
			(_nw336).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool), (_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool)), 3)
			(_nw336).ArraySet1(((_this).F9()).Update((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool), (_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool)), 4)
			(_nw336).ArraySet1((_this).F9(), 5)
			(_nw336).ArraySet1(((_this).F9()).Merge((_this).F9()), 6)
			(_nw336).ArraySet1((_2134_v55).Dtor_cf28(), 7)
			(_nw336).ArraySet1((_this).F9(), 8)
			(_nw336).ArraySet1(((_this).F9()).Update((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool), (_this).F10()), 9)
			(_nw336).ArraySet1((_this).F9(), 10)
			(_nw336).ArraySet1((_this).F9(), 11)
			(_nw336).ArraySet1(((_this).F9()).Merge((_this).F9()), 12)
			(_nw336).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool)), (_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool)), 13)
			(_nw336).ArraySet1((_this).F9(), 14)
			_2135_v56 = _nw336
			var _2136_v57 _dafny.Sequence
			_ = _2136_v57
			_2136_v57 = _dafny.SeqOf((_this).F9())
			var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_2135_v56), 0))
			_ = _index355
			(_2135_v56).ArraySet1(((_2136_v57).Select((Companion_Default___.SafeIndex(_2124_v48, _dafny.IntOfUint32((_2136_v57).Cardinality()))).Uint32()).(_dafny.Map)).Merge((_this).F9()), (_index355).Int())
			var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_2135_v56), 0))
			_ = _index356
			(_2135_v56).ArraySet1((_this).F9(), (_index356).Int())
			var _2137_v58 D15
			_ = _2137_v58
			_2137_v58 = Companion_D15_.Create_DC39_(!((_this).F10()) || (!((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool))), (_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool))
			var _source28 D15 = _2137_v58
			_ = _source28
			if _source28.Is_DC38() {
				var _2138___mcc_h0 bool = _source28.Get_().(D15_DC38).Cf69
				_ = _2138___mcc_h0
				var _2139___mcc_h1 _dafny.Sequence = _source28.Get_().(D15_DC38).Cf70
				_ = _2139___mcc_h1
				var _2140___mcc_h2 bool = _source28.Get_().(D15_DC38).Cf71
				_ = _2140___mcc_h2
				var _2141___mcc_h3 bool = _source28.Get_().(D15_DC38).Cf72
				_ = _2141___mcc_h3
				var _2142___mcc_h4 _dafny.CodePoint = _source28.Get_().(D15_DC38).Cf73
				_ = _2142___mcc_h4
				var _2143_cf73 _dafny.CodePoint = _2142___mcc_h4
				_ = _2143_cf73
				var _2144_cf72 bool = _2141___mcc_h3
				_ = _2144_cf72
				var _2145_cf71 bool = _2140___mcc_h2
				_ = _2145_cf71
				var _2146_cf70 _dafny.Sequence = _2139___mcc_h1
				_ = _2146_cf70
				var _2147_cf69 bool = _2138___mcc_h0
				_ = _2147_cf69
				var _2148_v59 _dafny.Map
				_ = _2148_v59
				_2148_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2126_v50, _2144_cf72)
				_2148_v59 = ((_2148_v59).Merge((_2148_v59).Update(_2126_v50, !((_this).F10())))).Merge(_2148_v59)
				_2145_cf71 = _2144_cf72
				var _2149_v60 *C0
				_ = _2149_v60
				var _nw337 *C0 = New_C0_()
				_ = _nw337
				_nw337.Ctor__()
				_2149_v60 = _nw337
				_2124_v48 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_2124_v48), _2124_v48)
			} else if _source28.Is_DC39() {
				var _2150___mcc_h5 bool = _source28.Get_().(D15_DC39).Cf74
				_ = _2150___mcc_h5
				var _2151___mcc_h6 bool = _source28.Get_().(D15_DC39).Cf75
				_ = _2151___mcc_h6
				var _2152_cf75 bool = _2151___mcc_h6
				_ = _2152_cf75
				var _2153_cf74 bool = _2150___mcc_h5
				_ = _2153_cf74
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))
				_ = _index357
				(_2126_v50).ArraySet1((_this).F10(), (_index357).Int())
				var _2154_v61 _dafny.Sequence
				_ = _2154_v61
				_2154_v61 = _dafny.UnicodeSeqOfUtf8Bytes("dsijjtbga")
				var _2155_v62 *C1
				_ = _2155_v62
				var _nw338 *C1 = New_C1_()
				_ = _nw338
				_nw338.Ctor__((_this).Fm2(globalState), _2154_v61)
				_2155_v62 = _nw338
				var _2156_v63 _dafny.Array
				_ = _2156_v63
				var _nw339 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw339
				_2156_v63 = _nw339
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_2156_v63), 0))
				_ = _index358
				(_2156_v63).ArraySet1(_2154_v61, (_index358).Int())
				var _2157_v64 _dafny.Map
				_ = _2157_v64
				_2157_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2155_v62).F13())
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_2156_v63), 0))
				_ = _index359
				(_2156_v63).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_2157_v64).Contains(p0) {
						return (_2157_v64).Get(p0).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-651))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg116 _dafny.Int) interface{} {
							return coer116(arg116)
						}
					}((func(_2158_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_2159_i9 _dafny.Int) _dafny.CodePoint {
							return _2158_p0
						}
					})(p0)))
				})(), _2154_v61), (_2155_v62).F13()), (_index359).Int())
				var _2160_v65 _dafny.Sequence
				_ = _2160_v65
				var _2161_v66 bool
				_ = _2161_v66
				var _2162_v67 _dafny.Sequence
				_ = _2162_v67
				var _2163_v68 _dafny.Sequence
				_ = _2163_v68
				var _out63 _dafny.Sequence
				_ = _out63
				var _out64 bool
				_ = _out64
				var _out65 _dafny.Sequence
				_ = _out65
				var _out66 _dafny.Sequence
				_ = _out66
				_out63, _out64, _out65, _out66 = (_2155_v62).M0(globalState)
				_2160_v65 = _out63
				_2161_v66 = _out64
				_2162_v67 = _out65
				_2163_v68 = _out66
			} else {
				var _2164___mcc_h7 _dafny.Array = _source28.Get_().(D15_DC37).Cf68
				_ = _2164___mcc_h7
				var _2165_cf68 _dafny.Array = _2164___mcc_h7
				_ = _2165_cf68
				var _2166_v69 _dafny.CodePoint
				_ = _2166_v69
				_2166_v69 = _dafny.CodePoint('r')
				var _2167_v70 _dafny.Map
				_ = _2167_v70
				_2167_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2124_v48, (_this).F10())
				_2166_v69 = Companion_Default___.Fm31((_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool), _2167_v70, globalState)
				var _2168_v71 _dafny.Set
				_ = _2168_v71
				_2168_v71 = _dafny.SetOf(_2124_v48, _2124_v48)
				var _2169_v72 _dafny.Map
				_ = _2169_v72
				_2169_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2168_v71, _2064_v0)
				_2124_v48 = ((_2169_v72).Merge(_2169_v72)).Cardinality()
				r1 = (_2126_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2126_v50), 0))).Int()).(bool)
				var _2170_v73 _dafny.Array
				_ = _2170_v73
				var _nw340 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw340
				_2170_v73 = _nw340
				var _2171_v74 _dafny.Map
				_ = _2171_v74
				_2171_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2124_v48, _2170_v73)
				_2171_v74 = (_2171_v74).Update((_2124_v48).Times(_2124_v48), _2170_v73)
			}
			_2124_v48 = _2124_v48
			var _2172_v75 _dafny.Array
			_ = _2172_v75
			var _len0_63 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_63
			var _nw341 _dafny.Array
			_ = _nw341
			if _len0_63.Cmp(_dafny.Zero) == 0 {
				_nw341 = _dafny.NewArray(_len0_63)
			} else {
				var _init63 func(_dafny.Int) _dafny.Int = (func(_2173_v48 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2174_i10 _dafny.Int) _dafny.Int {
						return (_2174_i10).Plus(_2173_v48)
					}
				})(_2124_v48)
				_ = _init63
				var _element0_63 = _init63(_dafny.Zero)
				_ = _element0_63
				_nw341 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
				(_nw341).ArraySet1(_element0_63, 0)
				var _nativeLen0_63 = (_len0_63).Int()
				_ = _nativeLen0_63
				for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
					(_nw341).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
				}
			}
			_2172_v75 = _nw341
			var _2175_v76 _dafny.Map
			_ = _2175_v76
			_2175_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2126_v50, _2172_v75)
			var _2176_v77 _dafny.Map
			_ = _2176_v77
			_2176_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2172_v75, (_2175_v76).Merge(_2175_v76))
			_2176_v77 = (_2176_v77).Update(_2172_v75, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2126_v50, _2172_v75))
		} else {
			var _2177_v78 D0
			_ = _2177_v78
			_2177_v78 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(598), (_this).F10(), (_dafny.Zero).Minus(_2124_v48))
			var _2178_v79 _dafny.Map
			_ = _2178_v79
			_2178_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2124_v48, (_this).F11())
			var _2179_v80 _dafny.Array
			_ = _2179_v80
			var _nwElement0_70 bool = (_2177_v78).Dtor_cf2()
			_ = _nwElement0_70
			var _nw342 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(22))
			_ = _nw342
			(_nw342).ArraySet1(_nwElement0_70, 0)
			(_nw342).ArraySet1((_this).F10(), 1)
			(_nw342).ArraySet1((_this).F10(), 2)
			(_nw342).ArraySet1((_this).F10(), 3)
			(_nw342).ArraySet1((_this).F10(), 4)
			(_nw342).ArraySet1((_this).F10(), 5)
			(_nw342).ArraySet1((_this).F10(), 6)
			(_nw342).ArraySet1((_this).F10(), 7)
			(_nw342).ArraySet1((_this).F10(), 8)
			(_nw342).ArraySet1((_this).F10(), 9)
			(_nw342).ArraySet1(!((_this).F10()), 10)
			(_nw342).ArraySet1((_this).Fm2(globalState), 11)
			(_nw342).ArraySet1((_this).F10(), 12)
			(_nw342).ArraySet1(!(!((_this).F10())), 13)
			(_nw342).ArraySet1((_this).F10(), 14)
			(_nw342).ArraySet1((_this).F10(), 15)
			(_nw342).ArraySet1(!((_this).F10()), 16)
			(_nw342).ArraySet1((_this).F10(), 17)
			(_nw342).ArraySet1(((_this).F11()).IsDisjointFrom((func() _dafny.Set {
				if (_2178_v79).Contains(_2124_v48) {
					return (_2178_v79).Get(_2124_v48).(_dafny.Set)
				}
				return (_this).F11()
			})()), 18)
			(_nw342).ArraySet1(false, 19)
			(_nw342).ArraySet1((_this).F10(), 20)
			(_nw342).ArraySet1((_this).F10(), 21)
			_2179_v80 = _nw342
			var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))
			_ = _index360
			(_2179_v80).ArraySet1((_this).F10(), (_index360).Int())
			var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))
			_ = _index361
			(_2179_v80).ArraySet1(((_this).F10()) || (false), (_index361).Int())
			var _2180_v81 _dafny.CodePoint
			_ = _2180_v81
			_2180_v81 = _dafny.CodePoint('x')
			_2180_v81 = p0
			var _2181_v82 _dafny.Sequence
			_ = _2181_v82
			_2181_v82 = _dafny.UnicodeSeqOfUtf8Bytes("tf")
			if _dafny.Companion_Sequence_.IsPrefixOf(_2181_v82, _2181_v82) {
				var _2182_v83 _dafny.Array
				_ = _2182_v83
				var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
				_ = _nw343
				_2182_v83 = _nw343
				var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_2182_v83), 0))
				_ = _index362
				(_2182_v83).ArraySet1(_2179_v80, (_index362).Int())
				var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.ArrayLen((_2182_v83), 0))
				_ = _index363
				(_2182_v83).ArraySet1(_2179_v80, (_index363).Int())
				var _2183_v84 *C3
				_ = _2183_v84
				var _nw344 *C3 = New_C3_()
				_ = _nw344
				_nw344.Ctor__()
				_2183_v84 = _nw344
				_2124_v48 = Companion_Default___.SafeDivisionInt(_2124_v48, _2124_v48)
				var _2184_v85 _dafny.MultiSet
				_ = _2184_v85
				_2184_v85 = _dafny.MultiSetOf(_dafny.IntOfInt64(997), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2124_v48)).Cardinality())
				var _2185_v86 _dafny.Map
				_ = _2185_v86
				_2185_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2184_v85).Update(_2124_v48, Companion_Default___.Abs(_2124_v48)), _2125_v49)
				var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))
				_ = _index364
				(_2179_v80).ArraySet1(((_2124_v48).Cmp(_2124_v48) <= 0) || (!(!(_dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
					if (_2185_v86).Contains(_2184_v85) {
						return (_2185_v86).Get(_2184_v85).(_dafny.Sequence)
					}
					return _2125_v49
				})(), _2125_v49)))), (_index364).Int())
				var _2186_v87 _dafny.Sequence
				_ = _2186_v87
				_2186_v87 = _dafny.SeqOf(_2124_v48, _2124_v48)
				var _2187_v88 _dafny.Array
				_ = _2187_v88
				var _nwElement0_71 _dafny.MultiSet = _2184_v85
				_ = _nwElement0_71
				var _nw345 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(15))
				_ = _nw345
				(_nw345).ArraySet1(_nwElement0_71, 0)
				(_nw345).ArraySet1(_2184_v85, 1)
				(_nw345).ArraySet1(_2184_v85, 2)
				(_nw345).ArraySet1(_dafny.MultiSetFromSeq(_2186_v87), 3)
				(_nw345).ArraySet1(Companion_Default___.Fm46((_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), (_this).F10(), (_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), globalState), 4)
				(_nw345).ArraySet1(_2184_v85, 5)
				(_nw345).ArraySet1(_2184_v85, 6)
				(_nw345).ArraySet1(_2184_v85, 7)
				(_nw345).ArraySet1(_2184_v85, 8)
				(_nw345).ArraySet1(_dafny.MultiSetOf(_2124_v48), 9)
				(_nw345).ArraySet1(_2184_v85, 10)
				(_nw345).ArraySet1(_2184_v85, 11)
				(_nw345).ArraySet1(_dafny.MultiSetOf(_2124_v48, _2124_v48), 12)
				(_nw345).ArraySet1(_2184_v85, 13)
				(_nw345).ArraySet1(_dafny.MultiSetFromSeq(_2186_v87), 14)
				_2187_v88 = _nw345
				var _2188_v89 _dafny.Map
				_ = _2188_v89
				_2188_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2187_v88)
				var _2189_v90 D20
				_ = _2189_v90
				_2189_v90 = Companion_D20_.Create_DC51_(_2187_v88)
				var _2190_v91 _dafny.Array
				_ = _2190_v91
				var _nwElement0_72 _dafny.Array = (func() _dafny.Array {
					if (_2188_v89).Contains((_this).F10()) {
						return (_2188_v89).Get((_this).F10()).(_dafny.Array)
					}
					return _2187_v88
				})()
				_ = _nwElement0_72
				var _nw346 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(14))
				_ = _nw346
				(_nw346).ArraySet1(_nwElement0_72, 0)
				(_nw346).ArraySet1((_2189_v90).Dtor_cf97(), 1)
				(_nw346).ArraySet1(_2187_v88, 2)
				(_nw346).ArraySet1((func() _dafny.Array {
					if (_2188_v89).Contains(true) {
						return (_2188_v89).Get(true).(_dafny.Array)
					}
					return _2187_v88
				})(), 3)
				(_nw346).ArraySet1(_2187_v88, 4)
				(_nw346).ArraySet1(_2187_v88, 5)
				(_nw346).ArraySet1(_2187_v88, 6)
				(_nw346).ArraySet1(_2187_v88, 7)
				(_nw346).ArraySet1(_2187_v88, 8)
				(_nw346).ArraySet1(_2187_v88, 9)
				(_nw346).ArraySet1(_2187_v88, 10)
				(_nw346).ArraySet1(_2187_v88, 11)
				(_nw346).ArraySet1(_2187_v88, 12)
				(_nw346).ArraySet1((func() _dafny.Array {
					if (_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool) {
						return _2187_v88
					}
					return _2187_v88
				})(), 13)
				_2190_v91 = _nw346
				var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_2190_v91), 0))
				_ = _index365
				(_2190_v91).ArraySet1(_2187_v88, (_index365).Int())
				var _2191_v92 _dafny.Sequence
				_ = _2191_v92
				_2191_v92 = _dafny.SeqOf(_2187_v88)
				var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(45), _dafny.ArrayLen((_2190_v91), 0))
				_ = _index366
				(_2190_v91).ArraySet1((_2191_v92).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(84)).Minus(_dafny.IntOfUint32((_2125_v49).Cardinality())), _dafny.IntOfUint32((_2191_v92).Cardinality()))).Uint32()).(_dafny.Array), (_index366).Int())
			} else {
				var _2192_v93 _dafny.Map
				_ = _2192_v93
				_2192_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), true)
				_2192_v93 = ((_this).F9()).Merge((Companion_D6_.Create_DC16_((_this).F9())).Dtor_cf28())
				var _2193_v94 _dafny.Sequence
				_ = _2193_v94
				_2193_v94 = _dafny.SeqOf(_2124_v48)
				var _2194_v95 _dafny.Array
				_ = _2194_v95
				var _len0_64 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_64
				var _nw347 _dafny.Array
				_ = _nw347
				if _len0_64.Cmp(_dafny.Zero) == 0 {
					_nw347 = _dafny.NewArray(_len0_64)
				} else {
					var _init64 func(_dafny.Int) _dafny.Int = (func(_2195_v48 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2196_i11 _dafny.Int) _dafny.Int {
							return (_2196_i11).Minus(_2195_v48)
						}
					})(_2124_v48)
					_ = _init64
					var _element0_64 = _init64(_dafny.Zero)
					_ = _element0_64
					_nw347 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
					(_nw347).ArraySet1(_element0_64, 0)
					var _nativeLen0_64 = (_len0_64).Int()
					_ = _nativeLen0_64
					for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
						(_nw347).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
					}
				}
				_2194_v95 = _nw347
				var _2197_v96 _dafny.Map
				_ = _2197_v96
				_2197_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2124_v48, _2124_v48)
				var _2198_v97 _dafny.Map
				_ = _2198_v97
				_2198_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2194_v95, (func() _dafny.Int {
					if (_2197_v96).Contains(_2124_v48) {
						return (_2197_v96).Get(_2124_v48).(_dafny.Int)
					}
					return _2124_v48
				})())
				var _2199_v98 _dafny.Map
				_ = _2199_v98
				_2199_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D4_.Create_DC12_(true, _2124_v48, _2180_v81, _2124_v48, _2124_v48), _2124_v48)
				var _2200_v99 D4
				_ = _2200_v99
				_2200_v99 = Companion_D4_.Create_DC12_((_this).F10(), _2124_v48, _2180_v81, _2124_v48, _2124_v48)
				var _2201_v100 _dafny.Sequence
				_ = _2201_v100
				_2201_v100 = _dafny.SeqOf(_2193_v94, _2193_v94)
				var _2202_v101 _dafny.Array
				_ = _2202_v101
				var _nwElement0_73 _dafny.Int = (_dafny.SetOf(Companion_Default___.Fm17((_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), globalState))).Cardinality()
				_ = _nwElement0_73
				var _nw348 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(6))
				_ = _nw348
				(_nw348).ArraySet1(_nwElement0_73, 0)
				(_nw348).ArraySet1((_2193_v94).Select((Companion_Default___.SafeIndex(_2124_v48, _dafny.IntOfUint32((_2193_v94).Cardinality()))).Uint32()).(_dafny.Int), 1)
				(_nw348).ArraySet1(((_2198_v97).Cardinality()).Times((func() _dafny.Int {
					if (_2199_v98).Contains(_2200_v99) {
						return (_2199_v98).Get(_2200_v99).(_dafny.Int)
					}
					return _2124_v48
				})()), 2)
				(_nw348).ArraySet1(_2124_v48, 3)
				(_nw348).ArraySet1(_dafny.IntOfInt64(-343), 4)
				(_nw348).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(327))).Uint32(), func(coer117 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg117 _dafny.Int) interface{} {
						return coer117(arg117)
					}
				}((func(_2203_v94 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2204_i12 _dafny.Int) _dafny.Sequence {
						return _2203_v94
					}
				})(_2193_v94))), _2201_v100)).Cardinality()), 5)
				_2202_v101 = _nw348
				var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2194_v95), 0))
				_ = _index367
				(_2194_v95).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2181_v82, _2181_v82)).Cardinality()), (_index367).Int())
				var _2205_v102 _dafny.MultiSet
				_ = _2205_v102
				_2205_v102 = _dafny.MultiSetOf((_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool))
				var _2206_v103 _dafny.MultiSet
				_ = _2206_v103
				_2206_v103 = _dafny.MultiSetOf(_2124_v48, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2181_v82, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_2124_v48), _dafny.IntOfUint32((_2181_v82).Cardinality()))).Uint32(), _2180_v81)).Cardinality()), _2124_v48)
				var _2207_v104 _dafny.Sequence
				_ = _2207_v104
				_2207_v104 = _dafny.SeqOf(_2206_v103, _2206_v103)
				var _2208_v105 _dafny.Sequence
				_ = _2208_v105
				_2208_v105 = _dafny.SeqOf(_2206_v103, _2206_v103, (_2207_v104).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2181_v82).Cardinality()), _dafny.IntOfUint32((_2207_v104).Cardinality()))).Uint32()).(_dafny.MultiSet), Companion_Default___.Fm46((_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), true, (_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), globalState))
				var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))
				_ = _index368
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2194_v95), 0))
				_ = _index369
				var _rhs343 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2181_v82, _dafny.Companion_Sequence_.Concatenate(_2181_v82, _2181_v82))
				_ = _rhs343
				var _rhs344 bool = true
				_ = _rhs344
				var _rhs345 _dafny.Int = (func() _dafny.Int {
					if (_this).F10() {
						return (func() _dafny.Int {
							if (_2197_v96).Contains(_dafny.IntOfInt64(799)) {
								return (_2197_v96).Get(_dafny.IntOfInt64(799)).(_dafny.Int)
							}
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2193_v94, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.IntOfUint32((_2193_v94).Cardinality()))).Uint32(), (_dafny.Zero).Minus((_2205_v102).Cardinality()))).Cardinality()))
						})()
					}
					return Companion_Default___.SafeDivisionInt(_2124_v48, (_this).Fm3(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fscb")).Cardinality()), _2124_v48, _2180_v81, _dafny.IntOfUint32((_2208_v105).Cardinality()), globalState))
				})()
				_ = _rhs345
				var _rhs346 _dafny.CodePoint = p0
				_ = _rhs346
				var _rhs347 _dafny.Int = (func() _dafny.Int {
					if (_this).F10() {
						return _2124_v48
					}
					return _dafny.IntOfInt64(81)
				})()
				_ = _rhs347
				var _lhs248 _dafny.Array = _2179_v80
				_ = _lhs248
				var _lhs249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))
				_ = _lhs249
				var _lhs250 _dafny.Array = _2194_v95
				_ = _lhs250
				var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2194_v95), 0))
				_ = _lhs251
				_2181_v82 = _rhs343
				(_lhs248).ArraySet1(_rhs344, (_lhs249).Int())
				_2124_v48 = _rhs345
				_2180_v81 = _rhs346
				(_lhs250).ArraySet1(_rhs347, (_lhs251).Int())
				var _2209_v106 D5
				_ = _2209_v106
				_2209_v106 = Companion_D5_.Create_DC13_(_2179_v80)
				var _2210_v107 _dafny.Set
				_ = _2210_v107
				_2210_v107 = _dafny.SetOf(_2209_v106, _2209_v106, _2209_v106)
				var _2211_v108 _dafny.Map
				_ = _2211_v108
				_2211_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2180_v81, _dafny.SetOf(_2209_v106))
				var _2212_v109 _dafny.CodePoint
				_ = _2212_v109
				_2212_v109 = _2180_v81
				_2210_v107 = (func() _dafny.Set {
					if (_2211_v108).Contains(_2212_v109) {
						return (_2211_v108).Get(_2212_v109).(_dafny.Set)
					}
					return _2210_v107
				})()
				_2124_v48 = ((_2194_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_2194_v95), 0))).Int()).(_dafny.Int)).Times(_2124_v48)
				var _2213_v110 _dafny.Map
				_ = _2213_v110
				_2213_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2124_v48, _2202_v101)
				_2202_v101 = (func() _dafny.Array {
					if (_2213_v110).Contains(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(570), _2124_v48)) {
						return (_2213_v110).Get(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(570), _2124_v48)).(_dafny.Array)
					}
					return _2194_v95
				})()
			}
			var _2214_v111 _dafny.Map
			_ = _2214_v111
			_2214_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2179_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))).Int()).(bool), _2124_v48)
			_2214_v111 = (_2214_v111).Update((_this).F10(), _2124_v48)
			var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_2179_v80), 0))
			_ = _index370
			(_2179_v80).ArraySet1((_this).F10(), (_index370).Int())
		}
		var _2215_v112 _dafny.Int
		_ = _2215_v112
		var _2216_v113 _dafny.Sequence
		_ = _2216_v113
		var _out67 _dafny.Int
		_ = _out67
		var _out68 _dafny.Sequence
		_ = _out68
		_out67, _out68 = (_this).M2(globalState)
		_2215_v112 = _out67
		_2216_v113 = _out68
		var _2217_v114 _dafny.CodePoint
		_ = _2217_v114
		_2217_v114 = _dafny.CodePoint('j')
		_2217_v114 = _2217_v114
		var _2218_v115 _dafny.Array
		_ = _2218_v115
		var _nw349 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw349
		_2218_v115 = _nw349
		var _2219_v116 _dafny.Map
		_ = _2219_v116
		_2219_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2218_v115, _dafny.UnicodeSeqOfUtf8Bytes("gwjllvov"))
		r0 = _2219_v116
		r1 = ((_this).F10()) || (true)
		return r0, r1
	}
}
func (_this *C7) M2(globalState *GlobalState) (_dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _2220_v0 _dafny.Int
		_ = _2220_v0
		_2220_v0 = _dafny.IntOfInt64(-474)
		var _2221_v1 _dafny.MultiSet
		_ = _2221_v1
		_2221_v1 = _dafny.MultiSetOf(_2220_v0)
		var _2222_v2 _dafny.MultiSet
		_ = _2222_v2
		_2222_v2 = _dafny.MultiSetOf((_2221_v1).Cardinality())
		(globalState).F8 = (_2222_v2).IsProperSubsetOf((_dafny.MultiSetOf(_2220_v0)).Intersection(_2221_v1))
		if (_this).F10() {
			var _2223_v3 *C5
			_ = _2223_v3
			var _nw350 *C5 = New_C5_()
			_ = _nw350
			_nw350.Ctor__(((_this).F9()).Update(false, (_this).F10()), !((_this).F10()))
			_2223_v3 = _nw350
			var _2224_v4 _dafny.Sequence
			_ = _2224_v4
			_2224_v4 = _dafny.UnicodeSeqOfUtf8Bytes("r")
			var _2225_v5 T0
			_ = _2225_v5
			var _nw351 *C1 = New_C1_()
			_ = _nw351
			_nw351.Ctor__((_this).F10(), _2224_v4)
			_2225_v5 = _nw351
			var _2226_v6 D21
			_ = _2226_v6
			_2226_v6 = Companion_D21_.Create_DC53_(_2225_v5)
			var _2227_v7 _dafny.Array
			_ = _2227_v7
			var _nwElement0_74 T0 = _2225_v5
			_ = _nwElement0_74
			var _nw352 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(12))
			_ = _nw352
			(_nw352).ArraySet1(_nwElement0_74, 0)
			(_nw352).ArraySet1(_2225_v5, 1)
			(_nw352).ArraySet1(_2225_v5, 2)
			(_nw352).ArraySet1(_2225_v5, 3)
			(_nw352).ArraySet1((_2226_v6).Dtor_cf99(), 4)
			(_nw352).ArraySet1(_2225_v5, 5)
			(_nw352).ArraySet1(_2225_v5, 6)
			(_nw352).ArraySet1(_2225_v5, 7)
			(_nw352).ArraySet1(_2225_v5, 8)
			(_nw352).ArraySet1(_2225_v5, 9)
			(_nw352).ArraySet1(_2225_v5, 10)
			(_nw352).ArraySet1(_2225_v5, 11)
			_2227_v7 = _nw352
			var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_2227_v7), 0))
			_ = _index371
			(_2227_v7).ArraySet1(_2225_v5, (_index371).Int())
			var _2228_v8 _dafny.CodePoint
			_ = _2228_v8
			_2228_v8 = _dafny.CodePoint('m')
			var _2229_v9 D6
			_ = _2229_v9
			_2229_v9 = Companion_D6_.Create_DC19_(true, _2228_v8)
			var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_2227_v7), 0))
			_ = _index372
			var _rhs348 T0 = (func() T0 {
				if (_2220_v0).Cmp(_2220_v0) >= 0 {
					return _2225_v5
				}
				return _2225_v5
			})()
			_ = _rhs348
			var _rhs349 bool = (!((_this).F10())) == (false)
			_ = _rhs349
			var _rhs350 T0 = _2225_v5
			_ = _rhs350
			var _rhs351 D6 = _2229_v9
			_ = _rhs351
			var _lhs252 *GlobalState = globalState
			_ = _lhs252
			var _lhs253 _dafny.Array = _2227_v7
			_ = _lhs253
			var _lhs254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_2227_v7), 0))
			_ = _lhs254
			_2225_v5 = _rhs348
			_lhs252.F8 = _rhs349
			(_lhs253).ArraySet1(_rhs350, (_lhs254).Int())
			_2229_v9 = _rhs351
			var _2230_v10 _dafny.Array
			_ = _2230_v10
			var _len0_65 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_65
			var _nw353 _dafny.Array
			_ = _nw353
			if _len0_65.Cmp(_dafny.Zero) == 0 {
				_nw353 = _dafny.NewArray(_len0_65)
			} else {
				var _init65 func(_dafny.Int) _dafny.Int = (func(_2231_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2232_i0 _dafny.Int) _dafny.Int {
						return (_2232_i0).Times(_2231_v0)
					}
				})(_2220_v0)
				_ = _init65
				var _element0_65 = _init65(_dafny.Zero)
				_ = _element0_65
				_nw353 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
				(_nw353).ArraySet1(_element0_65, 0)
				var _nativeLen0_65 = (_len0_65).Int()
				_ = _nativeLen0_65
				for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
					(_nw353).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
				}
			}
			_2230_v10 = _nw353
			var _2233_v11 _dafny.MultiSet
			_ = _2233_v11
			_2233_v11 = _dafny.MultiSetOf((_2223_v3).Fm18((_this).F11(), _dafny.UnicodeSeqOfUtf8Bytes("boc"), globalState), true)
			var _2234_v12 _dafny.Map
			_ = _2234_v12
			_2234_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ejnrlofe")).Cardinality()), (_2233_v11).Cardinality())
			var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_2230_v10), 0))
			_ = _index373
			(_2230_v10).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_this).F10(), (_this).F10(), (_this).F10())).Cardinality()), (_2234_v12).Cardinality()), (_index373).Int())
			var _2235_v13 T1
			_ = _2235_v13
			var _nw354 *C4 = New_C4_()
			_ = _nw354
			_nw354.Ctor__((_this).F9(), (_this).F10())
			_2235_v13 = _nw354
			var _2236_v14 _dafny.Array
			_ = _2236_v14
			var _nwElement0_75 T1 = _2235_v13
			_ = _nwElement0_75
			var _nw355 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(4))
			_ = _nw355
			(_nw355).ArraySet1(_nwElement0_75, 0)
			(_nw355).ArraySet1((Companion_D22_.Create_DC56_(_2235_v13)).Dtor_cf104(), 1)
			(_nw355).ArraySet1((func() T1 {
				if (_this).F10() {
					return _2235_v13
				}
				return _2235_v13
			})(), 2)
			(_nw355).ArraySet1(_2235_v13, 3)
			_2236_v14 = _nw355
			var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_2236_v14), 0))
			_ = _index374
			(_2236_v14).ArraySet1(_2235_v13, (_index374).Int())
			var _2237_v15 _dafny.Sequence
			_ = _2237_v15
			_2237_v15 = _dafny.SeqOf((_this).F10())
			var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_2230_v10), 0))
			_ = _index375
			var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_2236_v14), 0))
			_ = _index376
			var _rhs352 _dafny.Int = ((_2223_v3).Fm1((_this).F9(), globalState)).Times(_dafny.IntOfInt64(122))
			_ = _rhs352
			var _rhs353 _dafny.Int = _2220_v0
			_ = _rhs353
			var _rhs354 _dafny.Int = (_this).Fm1(((_2235_v13).F9()).Update((_2235_v13).F10(), (_2235_v13).F10()), globalState)
			_ = _rhs354
			var _rhs355 bool = (_dafny.IntOfInt64(604)).Cmp(_dafny.IntOfUint32((_2237_v15).Cardinality())) == 0
			_ = _rhs355
			var _rhs356 T1 = _2235_v13
			_ = _rhs356
			var _lhs255 _dafny.Array = _2230_v10
			_ = _lhs255
			var _lhs256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_2230_v10), 0))
			_ = _lhs256
			var _lhs257 *GlobalState = globalState
			_ = _lhs257
			var _lhs258 _dafny.Array = _2236_v14
			_ = _lhs258
			var _lhs259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(884), _dafny.ArrayLen((_2236_v14), 0))
			_ = _lhs259
			(_lhs255).ArraySet1(_rhs352, (_lhs256).Int())
			_2220_v0 = _rhs353
			r0 = _rhs354
			_lhs257.F8 = _rhs355
			(_lhs258).ArraySet1(_rhs356, (_lhs259).Int())
			(globalState).F8 = false
			_2224_v4 = _dafny.Companion_Sequence_.Update(_2224_v4, (Companion_Default___.SafeIndex((_2230_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_2230_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2224_v4).Cardinality()))).Uint32(), _2228_v8)
		} else {
			var _2238_v16 _dafny.Array
			_ = _2238_v16
			var _len0_66 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_66
			var _nw356 _dafny.Array
			_ = _nw356
			if _len0_66.Cmp(_dafny.Zero) == 0 {
				_nw356 = _dafny.NewArray(_len0_66)
			} else {
				var _init66 func(_dafny.Int) _dafny.Int = (func(_2239_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2240_i1 _dafny.Int) _dafny.Int {
						return (_2240_i1).Times(_2239_v0)
					}
				})(_2220_v0)
				_ = _init66
				var _element0_66 = _init66(_dafny.Zero)
				_ = _element0_66
				_nw356 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
				(_nw356).ArraySet1(_element0_66, 0)
				var _nativeLen0_66 = (_len0_66).Int()
				_ = _nativeLen0_66
				for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
					(_nw356).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
				}
			}
			_2238_v16 = _nw356
			var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _index377
			(_2238_v16).ArraySet1(_2220_v0, (_index377).Int())
			var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _index378
			(_2238_v16).ArraySet1((_dafny.Zero).Minus(_2220_v0), (_index378).Int())
			var _2241_v17 _dafny.Sequence
			_ = _2241_v17
			_2241_v17 = _dafny.UnicodeSeqOfUtf8Bytes("f")
			var _2242_v18 D14
			_ = _2242_v18
			_2242_v18 = Companion_D14_.Create_DC36_((_this).F10(), _2241_v17, _2241_v17)
			var _2243_v19 _dafny.Sequence
			_ = _2243_v19
			_2243_v19 = _dafny.SeqOf(_2242_v18)
			var _2244_v20 _dafny.CodePoint
			_ = _2244_v20
			_2244_v20 = _dafny.CodePoint('e')
			var _2245_v21 _dafny.Array
			_ = _2245_v21
			var _nwElement0_76 _dafny.Sequence = _dafny.SeqOf(_2242_v18, _2242_v18, _2242_v18)
			_ = _nwElement0_76
			var _nw357 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(17))
			_ = _nw357
			(_nw357).ArraySet1(_nwElement0_76, 0)
			(_nw357).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm56((_this).F10(), globalState), Companion_Default___.Fm56(true, globalState)), 1)
			(_nw357).ArraySet1(_2243_v19, 2)
			(_nw357).ArraySet1(_2243_v19, 3)
			(_nw357).ArraySet1(_dafny.SeqOf(_2242_v18, Companion_D14_.Create_DC36_((_this).F10(), _2241_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg118 _dafny.Int) interface{} {
					return coer118(arg118)
				}
			}(func(_2246_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			})))), 4)
			(_nw357).ArraySet1(_dafny.SeqOf(_2242_v18), 5)
			(_nw357).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2242_v18), _2243_v19), 6)
			(_nw357).ArraySet1(_2243_v19, 7)
			(_nw357).ArraySet1(_2243_v19, 8)
			(_nw357).ArraySet1(_2243_v19, 9)
			(_nw357).ArraySet1(_2243_v19, 10)
			(_nw357).ArraySet1(_2243_v19, 11)
			(_nw357).ArraySet1(_2243_v19, 12)
			(_nw357).ArraySet1(Companion_Default___.Fm56((_this).F10(), globalState), 13)
			(_nw357).ArraySet1(_2243_v19, 14)
			(_nw357).ArraySet1(_2243_v19, 15)
			(_nw357).ArraySet1(_dafny.SeqOf(Companion_D14_.Create_DC36_((_this).F10(), _dafny.Companion_Sequence_.Update(_2241_v17, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lmmm")).Cardinality()), _dafny.IntOfUint32((_2241_v17).Cardinality()))).Uint32(), _2244_v20), _2241_v17)), 16)
			_2245_v21 = _nw357
			var _pat_let_tv50 = _2245_v21
			_ = _pat_let_tv50
			var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _index379
			var _rhs357 _dafny.Int = (_2238_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))).Int()).(_dafny.Int)
			_ = _rhs357
			var _rhs358 _dafny.Array = (func() _dafny.Array {
				if (_this).F10() {
					return (func(_pat_let52_0 D23) D23 {
						return func(_2247_dt__update__tmp_h0 D23) D23 {
							return func(_pat_let53_0 _dafny.Array) D23 {
								return func(_2248_dt__update_hcf105_h0 _dafny.Array) D23 {
									return Companion_D23_.Create_DC58_(_2248_dt__update_hcf105_h0)
								}(_pat_let53_0)
							}(_pat_let_tv50)
						}(_pat_let52_0)
					}(Companion_D23_.Create_DC58_(_2245_v21))).Dtor_cf105()
				}
				return _2245_v21
			})()
			_ = _rhs358
			var _lhs260 _dafny.Array = _2238_v16
			_ = _lhs260
			var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _lhs261
			(_lhs260).ArraySet1(_rhs357, (_lhs261).Int())
			_2245_v21 = _rhs358
			var _2249_v22 _dafny.Array
			_ = _2249_v22
			var _nw358 _dafny.Array = _dafny.NewArrayWithValue(Companion_D15_.Default(), _dafny.IntOfInt64(6))
			_ = _nw358
			_2249_v22 = _nw358
			var _2250_v23 D19
			_ = _2250_v23
			_2250_v23 = Companion_D19_.Create_DC49_(false)
			var _2251_v24 _dafny.Sequence
			_ = _2251_v24
			_2251_v24 = _dafny.SeqOf((_this).F10(), (_2250_v23).Dtor_cf93(), (_this).F10(), (_this).F10(), (_this).F10())
			var _2252_v25 D15
			_ = _2252_v25
			_2252_v25 = Companion_D15_.Create_DC39_((_this).F10(), (_2251_v24).Select((Companion_Default___.SafeIndex((_2238_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2251_v24).Cardinality()))).Uint32()).(bool))
			var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2249_v22), 0))
			_ = _index380
			(_2249_v22).ArraySet1(_2252_v25, (_index380).Int())
			var _2253_v26 _dafny.MultiSet
			_ = _2253_v26
			_2253_v26 = _dafny.MultiSetOf((_this).F10(), (_this).F10(), (_this).F10(), (_this).F10())
			var _pat_let_tv51 = _2253_v26
			_ = _pat_let_tv51
			var _pat_let_tv52 = _2251_v24
			_ = _pat_let_tv52
			var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2249_v22), 0))
			_ = _index381
			var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _index382
			var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _index383
			var _rhs359 D15 = func(_pat_let54_0 D15) D15 {
				return func(_2254_dt__update__tmp_h1 D15) D15 {
					return func(_pat_let55_0 bool) D15 {
						return func(_2255_dt__update_hcf74_h0 bool) D15 {
							return Companion_D15_.Create_DC39_(_2255_dt__update_hcf74_h0, (_2254_dt__update__tmp_h1).Dtor_cf75())
						}(_pat_let55_0)
					}((_pat_let_tv51).IsDisjointFrom(_dafny.MultiSetFromSeq(_pat_let_tv52)))
				}(_pat_let54_0)
			}(_2252_v25)
			_ = _rhs359
			var _rhs360 bool = (_this).F10()
			_ = _rhs360
			var _rhs361 _dafny.Int = ((_this).F9()).Cardinality()
			_ = _rhs361
			var _rhs362 _dafny.Int = _2220_v0
			_ = _rhs362
			var _lhs262 _dafny.Array = _2249_v22
			_ = _lhs262
			var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(13), _dafny.ArrayLen((_2249_v22), 0))
			_ = _lhs263
			var _lhs264 *GlobalState = globalState
			_ = _lhs264
			var _lhs265 _dafny.Array = _2238_v16
			_ = _lhs265
			var _lhs266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _lhs266
			var _lhs267 _dafny.Array = _2238_v16
			_ = _lhs267
			var _lhs268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.ArrayLen((_2238_v16), 0))
			_ = _lhs268
			(_lhs262).ArraySet1(_rhs359, (_lhs263).Int())
			_lhs264.F8 = _rhs360
			(_lhs265).ArraySet1(_rhs361, (_lhs266).Int())
			(_lhs267).ArraySet1(_rhs362, (_lhs268).Int())
			var _2256_v27 *C5
			_ = _2256_v27
			var _nw359 *C5 = New_C5_()
			_ = _nw359
			_nw359.Ctor__((_this).F9(), (_this).F10())
			_2256_v27 = _nw359
			var _2257_v28 _dafny.Array
			_ = _2257_v28
			var _nw360 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw360
			_2257_v28 = _nw360
			var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_2257_v28), 0))
			_ = _index384
			(_2257_v28).ArraySet1((_this).F10(), (_index384).Int())
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_2257_v28), 0))
			_ = _index385
			(_2257_v28).ArraySet1((_this).F10(), (_index385).Int())
		}
		var _2258_v29 _dafny.Sequence
		_ = _2258_v29
		_2258_v29 = _dafny.UnicodeSeqOfUtf8Bytes("glcese")
		if !(!(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _dafny.UnicodeSeqOfUtf8Bytes("nhghrw")), _2258_v29)))) {
			var _2259_v30 _dafny.Array
			_ = _2259_v30
			var _nw361 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw361
			_2259_v30 = _nw361
			var _2260_v31 _dafny.MultiSet
			_ = _2260_v31
			_2260_v31 = _dafny.MultiSetOf((_this).F10())
			var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))
			_ = _index386
			(_2259_v30).ArraySet1((_2260_v31).IsProperSubsetOf(_2260_v31), (_index386).Int())
			var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))
			_ = _index387
			var _rhs363 bool = (_this).F10()
			_ = _rhs363
			var _rhs364 _dafny.Int = _2220_v0
			_ = _rhs364
			var _rhs365 _dafny.Int = Companion_Default___.SafeDivisionInt(_2220_v0, ((_this).F11()).Cardinality())
			_ = _rhs365
			var _lhs269 _dafny.Array = _2259_v30
			_ = _lhs269
			var _lhs270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))
			_ = _lhs270
			(_lhs269).ArraySet1(_rhs363, (_lhs270).Int())
			r0 = _rhs364
			r0 = _rhs365
			var _2261_v32 _dafny.MultiSet
			_ = _2261_v32
			_2261_v32 = _dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kiarxiha")).Cardinality())))
			(globalState).F8 = !((_2261_v32).IsSubsetOf(_2261_v32))
			r0 = _2220_v0
			if (_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool) {
				var _2262_v33 _dafny.Array
				_ = _2262_v33
				var _nw362 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw362
				_2262_v33 = _nw362
				var _2263_v34 _dafny.Sequence
				_ = _2263_v34
				_2263_v34 = _dafny.SeqOf(_2220_v0, _2220_v0, _2220_v0, _dafny.IntOfInt64(93))
				var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))
				_ = _index388
				(_2262_v33).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2263_v34, _2263_v34), (Companion_Default___.SafeIndex((_2263_v34).Select((Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2263_v34).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2263_v34, _2263_v34)).Cardinality()))).Uint32(), _2220_v0)).Cardinality()), (_index388).Int())
				var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))
				_ = _index389
				(_2262_v33).ArraySet1((_2263_v34).Select((Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2263_v34).Cardinality()))).Uint32()).(_dafny.Int), (_index389).Int())
				var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))
				_ = _index390
				(_2262_v33).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2263_v34, _2263_v34), _2263_v34)).Cardinality()), (_index390).Int())
				var _2264_v35 _dafny.Set
				_ = _2264_v35
				_2264_v35 = _dafny.SetOf((_2262_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(815))
				var _2265_v36 _dafny.Map
				_ = _2265_v36
				_2265_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2264_v35).Cardinality(), (_2262_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))).Int()).(_dafny.Int))
				var _2266_v37 _dafny.Sequence
				_ = _2266_v37
				_2266_v37 = _dafny.SeqOf((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool))
				var _2267_v38 _dafny.CodePoint
				_ = _2267_v38
				_2267_v38 = _dafny.CodePoint('v')
				var _2268_v40 _dafny.Sequence
				_ = _2268_v40
				_2268_v40 = _dafny.SeqOf(_2264_v35)
				(globalState).F8 = !((_2265_v36).Update((_dafny.Zero).Minus((Companion_Default___.Fm15((_this).F10(), !((_2266_v37).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-274), _dafny.IntOfUint32((_2266_v37).Cardinality()))).Uint32()).(bool)), (_2264_v35).Cardinality(), _2267_v38, globalState)).Cardinality()), _2220_v0)).Equals(func() _dafny.Map {
					var _coll76 = _dafny.NewMapBuilder()
					_ = _coll76
					for _iter86 := _dafny.Iterate(((_2268_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ykixtfsa")).Cardinality()), _dafny.IntOfUint32((_2268_v40).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
						_compr_76, _ok86 := _iter86()
						if !_ok86 {
							break
						}
						var _2269_v39 _dafny.Int
						_2269_v39 = interface{}(_compr_76).(_dafny.Int)
						if ((_2268_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ykixtfsa")).Cardinality()), _dafny.IntOfUint32((_2268_v40).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_2269_v39) {
							_coll76.Add((_2269_v39).Plus((_2262_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))).Int()).(_dafny.Int)), _2220_v0)
						}
					}
					return _coll76.ToMap()
				}())
				var _2270_v41 *C4
				_ = _2270_v41
				var _nw363 *C4 = New_C4_()
				_ = _nw363
				_nw363.Ctor__((_this).F9(), (_this).F10())
				_2270_v41 = _nw363
				var _2271_v42 _dafny.Sequence
				_ = _2271_v42
				_2271_v42 = _dafny.SeqOf(_2270_v41)
				var _2272_v43 _dafny.Map
				_ = _2272_v43
				_2272_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_2265_v36).Contains(_2220_v0) {
						return (_2265_v36).Get(_2220_v0).(_dafny.Int)
					}
					return (_2262_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_2262_v33), 0))).Int()).(_dafny.Int)
				})(), (_2271_v42).Select((Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2271_v42).Cardinality()))).Uint32()).(*C4))
				_2272_v43 = _2272_v43
				var _2273_v44 *C2
				_ = _2273_v44
				var _nw364 *C2 = New_C2_()
				_ = _nw364
				_nw364.Ctor__((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool), _2259_v30, (_this).F9(), (_this).Fm0(_2266_v37, (_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool), _2220_v0, (_2270_v41).Fm0(_2266_v37, (_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool), _dafny.IntOfInt64(939), false, globalState), globalState))
				_2273_v44 = _nw364
				_2273_v44 = _2273_v44
			} else {
				var _2274_v45 _dafny.Map
				_ = _2274_v45
				_2274_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_this).F11()).IsProperSubsetOf((_this).F11())), (_dafny.IntOfInt64(237)).Cmp(_2220_v0) == 0)
				_2274_v45 = _2274_v45
				_2220_v0 = _2220_v0
				var _2275_v46 _dafny.Map
				_ = _2275_v46
				_2275_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2258_v29)
				var _2276_v47 *C1
				_ = _2276_v47
				var _nw365 *C1 = New_C1_()
				_ = _nw365
				_nw365.Ctor__((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_2275_v46).Contains((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool)) {
						return (_2275_v46).Get((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool)).(_dafny.Sequence)
					}
					return _2258_v29
				})(), _2258_v29))
				_2276_v47 = _nw365
				var _2277_v48 _dafny.Array
				_ = _2277_v48
				var _nw366 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw366
				_2277_v48 = _nw366
				var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_2277_v48), 0))
				_ = _index391
				(_2277_v48).ArraySet1(_2220_v0, (_index391).Int())
				var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_2277_v48), 0))
				_ = _index392
				(_2277_v48).ArraySet1(_2220_v0, (_index392).Int())
				var _2278_v49 *C5
				_ = _2278_v49
				var _nw367 *C5 = New_C5_()
				_ = _nw367
				_nw367.Ctor__((_2274_v45).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool))), ((_2277_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_2277_v48), 0))).Int()).(_dafny.Int)).Cmp((_2277_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_2277_v48), 0))).Int()).(_dafny.Int)) == 0)
				_2278_v49 = _nw367
			}
			var _2279_v50 _dafny.Array
			_ = _2279_v50
			var _nwElement0_77 _dafny.Int = Companion_Default___.Fm33((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool), (_this).F10(), _2220_v0, (_this).F10(), globalState)
			_ = _nwElement0_77
			var _nw368 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(10))
			_ = _nw368
			(_nw368).ArraySet1(_nwElement0_77, 0)
			(_nw368).ArraySet1((_2220_v0).Times(_2220_v0), 1)
			(_nw368).ArraySet1(_2220_v0, 2)
			(_nw368).ArraySet1(_2220_v0, 3)
			(_nw368).ArraySet1(_2220_v0, 4)
			(_nw368).ArraySet1((_dafny.Zero).Minus(_2220_v0), 5)
			(_nw368).ArraySet1(Companion_Default___.SafeModuloInt(_2220_v0, _2220_v0), 6)
			(_nw368).ArraySet1(_2220_v0, 7)
			(_nw368).ArraySet1(_2220_v0, 8)
			(_nw368).ArraySet1(Companion_Default___.Fm33((_2259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_2259_v30), 0))).Int()).(bool), (_this).F10(), _2220_v0, true, globalState), 9)
			_2279_v50 = _nw368
			_2279_v50 = _2279_v50
		} else {
			var _2280_v51 _dafny.Map
			_ = _2280_v51
			_2280_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()) == ((_this).F10()), !((_this).F10()))
			_2280_v51 = (_2280_v51).Update((_this).F10(), (_this).F10())
			var _2281_v52 _dafny.MultiSet
			_ = _2281_v52
			var _2282_v53 bool
			_ = _2282_v53
			var _2283_v54 _dafny.Int
			_ = _2283_v54
			var _2284_v55 _dafny.Sequence
			_ = _2284_v55
			var _out69 _dafny.MultiSet
			_ = _out69
			var _out70 bool
			_ = _out70
			var _out71 _dafny.Int
			_ = _out71
			var _out72 _dafny.Sequence
			_ = _out72
			_out69, _out70, _out71, _out72 = (_this).M3(_dafny.IntOfInt64(272), globalState)
			_2281_v52 = _out69
			_2282_v53 = _out70
			_2283_v54 = _out71
			_2284_v55 = _out72
			(globalState).F8 = (_this).F10()
			var _2285_v56 _dafny.Map
			_ = _2285_v56
			_2285_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_2283_v54), (_this).F10()), _2220_v0)
			var _2286_v57 _dafny.Sequence
			_ = _2286_v57
			_2286_v57 = _dafny.SeqOf((_2285_v56).Cardinality())
			var _2287_v58 _dafny.Map
			_ = _2287_v58
			_2287_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2282_v53, Companion_D1_.Create_DC2_(_2286_v57))
			var _2288_v59 _dafny.CodePoint
			_ = _2288_v59
			_2288_v59 = _dafny.CodePoint('c')
			_2284_v55 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29), (Companion_Default___.SafeIndex((_2287_v58).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29)).Cardinality()))).Uint32(), (Companion_D14_.Create_DC35_(_2220_v0, _2282_v53, _2288_v59)).Dtor_cf64()), (Companion_Default___.SafeIndex(_2283_v54, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29), (Companion_Default___.SafeIndex((_2287_v58).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29)).Cardinality()))).Uint32(), (Companion_D14_.Create_DC35_(_2220_v0, _2282_v53, _2288_v59)).Dtor_cf64())).Cardinality()))).Uint32(), _dafny.CodePoint('a'))
			var _2289_v60 _dafny.Array
			_ = _2289_v60
			var _len0_67 _dafny.Int = _dafny.One
			_ = _len0_67
			var _nw369 _dafny.Array
			_ = _nw369
			if _len0_67.Cmp(_dafny.Zero) == 0 {
				_nw369 = _dafny.NewArray(_len0_67)
			} else {
				var _init67 func(_dafny.Int) bool = func(_2290_i3 _dafny.Int) bool {
					return (_this).F10()
				}
				_ = _init67
				var _element0_67 = _init67(_dafny.Zero)
				_ = _element0_67
				_nw369 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
				(_nw369).ArraySet1(_element0_67, 0)
				var _nativeLen0_67 = (_len0_67).Int()
				_ = _nativeLen0_67
				for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
					(_nw369).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
				}
			}
			_2289_v60 = _nw369
			var _2291_v61 _dafny.Sequence
			_ = _2291_v61
			_2291_v61 = _dafny.SeqOf(false, _2282_v53)
			var _2292_v62 *C2
			_ = _2292_v62
			var _nw370 *C2 = New_C2_()
			_ = _nw370
			_nw370.Ctor__(false, _2289_v60, (_this).F9(), ((_this).Fm2(globalState)) || ((_2291_v61).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_2283_v54), _dafny.IntOfUint32((_2291_v61).Cardinality()))).Uint32()).(bool)))
			_2292_v62 = _nw370
		}
		var _2293_v63 _dafny.Map
		_ = _2293_v63
		_2293_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2222_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(761))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg119 _dafny.Int) interface{} {
				return coer119(arg119)
			}
		}(func(_2294_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()))
		var _hi7 _dafny.Int = (_2293_v63).Cardinality()
		_ = _hi7
		for _2295_i4 := _2220_v0; _2295_i4.Cmp(_hi7) < 0; _2295_i4 = _2295_i4.Plus(_dafny.One) {
			(globalState).F8 = (_2220_v0).Cmp(_2295_i4) != 0
			r0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2258_v29).Cardinality()), _dafny.IntOfUint32((_2258_v29).Cardinality()))).Times(_2295_i4)
			r1 = _dafny.Companion_Sequence_.Concatenate(_2258_v29, _dafny.UnicodeSeqOfUtf8Bytes("cemlbkmst"))
			var _2296_v64 _dafny.CodePoint
			_ = _2296_v64
			_2296_v64 = _dafny.CodePoint('b')
			var _2297_v65 D1
			_ = _2297_v65
			_2297_v65 = Companion_D1_.Create_DC3_(_2258_v29)
			var _2298_v66 _dafny.Array
			_ = _2298_v66
			var _nwElement0_78 _dafny.Sequence = _2258_v29
			_ = _nwElement0_78
			var _nw371 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(19))
			_ = _nw371
			(_nw371).ArraySet1(_nwElement0_78, 0)
			(_nw371).ArraySet1(_2258_v29, 1)
			(_nw371).ArraySet1(_2258_v29, 2)
			(_nw371).ArraySet1(_2258_v29, 3)
			(_nw371).ArraySet1(_2258_v29, 4)
			(_nw371).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2258_v29, (Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2258_v29).Cardinality()))).Uint32(), _dafny.CodePoint('q')), (Companion_Default___.SafeIndex(_2295_i4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2258_v29, (Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2258_v29).Cardinality()))).Uint32(), _dafny.CodePoint('q'))).Cardinality()))).Uint32(), _2296_v64), 5)
			(_nw371).ArraySet1(_dafny.SeqOf(_dafny.CodePoint('t')), 6)
			(_nw371).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29), 7)
			(_nw371).ArraySet1(_2258_v29, 8)
			(_nw371).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29), 9)
			(_nw371).ArraySet1(Companion_Default___.Fm13(_2220_v0, globalState), 10)
			(_nw371).ArraySet1((_2297_v65).Dtor_cf5(), 11)
			(_nw371).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29), 12)
			(_nw371).ArraySet1(_2258_v29, 13)
			(_nw371).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _dafny.SeqOf(_2296_v64)), 14)
			(_nw371).ArraySet1(_dafny.SeqOf(_2296_v64), 15)
			(_nw371).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(524))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg120 _dafny.Int) interface{} {
					return coer120(arg120)
				}
			}(func(_2299_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), 16)
			(_nw371).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29), 17)
			(_nw371).ArraySet1(_2258_v29, 18)
			_2298_v66 = _nw371
			var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_2298_v66), 0))
			_ = _index393
			(_2298_v66).ArraySet1(_2258_v29, (_index393).Int())
			var _2300_v67 T1
			_ = _2300_v67
			var _nw372 *C6 = New_C6_()
			_ = _nw372
			_nw372.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), false), !((_this).F10()))
			_2300_v67 = _nw372
			var _2301_v68 _dafny.Set
			_ = _2301_v68
			_2301_v68 = _dafny.SetOf(_2295_i4, _2295_i4)
			var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_2298_v66), 0))
			_ = _index394
			var _rhs366 _dafny.MultiSet = _dafny.MultiSetOf(_2220_v0, _dafny.IntOfUint32((_dafny.SeqOf(!((_this).F10()), (_this).F10())).Cardinality()), (func() _dafny.Int {
				if (_this).F10() {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2300_v67, !((_this).F10()))).Cardinality()
				}
				return _2295_i4
			})(), _2295_i4, (_2301_v68).Cardinality())
			_ = _rhs366
			var _rhs367 _dafny.Int = Companion_Default___.SafeDivisionInt(_2295_i4, (_dafny.Zero).Minus(_2295_i4))
			_ = _rhs367
			var _rhs368 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_2258_v29, (Companion_Default___.SafeIndex(_2295_i4, _dafny.IntOfUint32((_2258_v29).Cardinality()))).Uint32(), _dafny.CodePoint('x'))
			_ = _rhs368
			var _rhs369 bool = (_this).F10()
			_ = _rhs369
			var _rhs370 _dafny.CodePoint = _2296_v64
			_ = _rhs370
			var _lhs271 _dafny.Array = _2298_v66
			_ = _lhs271
			var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_2298_v66), 0))
			_ = _lhs272
			var _lhs273 *GlobalState = globalState
			_ = _lhs273
			_2222_v2 = _rhs366
			_2220_v0 = _rhs367
			(_lhs271).ArraySet1(_rhs368, (_lhs272).Int())
			_lhs273.F8 = _rhs369
			_2296_v64 = _rhs370
		}
		var _2302_v69 _dafny.Sequence
		_ = _2302_v69
		_2302_v69 = _dafny.SeqOf((_this).F10())
		var _2303_v70 _dafny.Array
		_ = _2303_v70
		var _nwElement0_79 bool = true
		_ = _nwElement0_79
		var _nw373 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(13))
		_ = _nw373
		(_nw373).ArraySet1(_nwElement0_79, 0)
		(_nw373).ArraySet1((_this).F10(), 1)
		(_nw373).ArraySet1((_this).F10(), 2)
		(_nw373).ArraySet1((_this).F10(), 3)
		(_nw373).ArraySet1((_this).F10(), 4)
		(_nw373).ArraySet1((_this).F10(), 5)
		(_nw373).ArraySet1((_this).Fm0(_2302_v69, (_this).F10(), (Companion_Default___.Fm25(_2220_v0, _dafny.IntOfInt64(435), globalState)).Cardinality(), (_this).F10(), globalState), 6)
		(_nw373).ArraySet1((_this).F10(), 7)
		(_nw373).ArraySet1((_this).F10(), 8)
		(_nw373).ArraySet1((_this).F10(), 9)
		(_nw373).ArraySet1((_this).F10(), 10)
		(_nw373).ArraySet1((_this).F10(), 11)
		(_nw373).ArraySet1((_this).F10(), 12)
		_2303_v70 = _nw373
		var _2304_v71 _dafny.Sequence
		_ = _2304_v71
		_2304_v71 = _dafny.SeqOf(_2303_v70)
		var _2305_v72 _dafny.Map
		_ = _2305_v72
		_2305_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_2220_v0)), (_this).F10())
		var _2306_v73 D5
		_ = _2306_v73
		_2306_v73 = Companion_D5_.Create_DC13_((_2304_v71).Select((Companion_Default___.SafeIndex((_2305_v72).Cardinality(), _dafny.IntOfUint32((_2304_v71).Cardinality()))).Uint32()).(_dafny.Array))
		var _source29 D5 = _2306_v73
		_ = _source29
		if _source29.Is_DC14() {
			var _2307___mcc_h0 _dafny.Int = _source29.Get_().(D5_DC14).Cf25
			_ = _2307___mcc_h0
			var _2308___mcc_h1 bool = _source29.Get_().(D5_DC14).Cf26
			_ = _2308___mcc_h1
			var _2309_cf26 bool = _2308___mcc_h1
			_ = _2309_cf26
			var _2310_cf25 _dafny.Int = _2307___mcc_h0
			_ = _2310_cf25
			var _2311_v74 _dafny.Sequence
			_ = _2311_v74
			_2311_v74 = _dafny.SeqOf((_this).F11(), (_this).F11(), (_this).F11(), (_this).F11())
			var _2312_v75 _dafny.Array
			_ = _2312_v75
			var _nwElement0_80 _dafny.Set = Companion_Default___.Fm30(globalState)
			_ = _nwElement0_80
			var _nw374 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(17))
			_ = _nw374
			(_nw374).ArraySet1(_nwElement0_80, 0)
			(_nw374).ArraySet1((_dafny.SetOf(_2309_cf26, (_this).Fm0(_2302_v69, _2309_cf26, _dafny.IntOfInt64(497), (_this).F10(), globalState), _2309_cf26, _2309_cf26, _2309_cf26)).Intersection(_dafny.SetOf((_this).F10(), false, _2309_cf26)), 1)
			(_nw374).ArraySet1(_dafny.SetOf(_2309_cf26, _2309_cf26, _2309_cf26), 2)
			(_nw374).ArraySet1(((_2311_v74).Select((Companion_Default___.SafeIndex(_2310_cf25, _dafny.IntOfUint32((_2311_v74).Cardinality()))).Uint32()).(_dafny.Set)).Difference((_this).F11()), 3)
			(_nw374).ArraySet1(((_this).F11()).Union((_this).F11()), 4)
			(_nw374).ArraySet1((_this).F11(), 5)
			(_nw374).ArraySet1((_this).F11(), 6)
			(_nw374).ArraySet1((_this).F11(), 7)
			(_nw374).ArraySet1((_this).F11(), 8)
			(_nw374).ArraySet1((_this).F11(), 9)
			(_nw374).ArraySet1((_this).F11(), 10)
			(_nw374).ArraySet1((_this).F11(), 11)
			(_nw374).ArraySet1(_dafny.SetOf((_this).F10()), 12)
			(_nw374).ArraySet1((_this).F11(), 13)
			(_nw374).ArraySet1(((_this).F11()).Difference((_this).F11()), 14)
			(_nw374).ArraySet1((_this).F11(), 15)
			(_nw374).ArraySet1((_this).F11(), 16)
			_2312_v75 = _nw374
			_2312_v75 = _2312_v75
			var _2313_v76 *C5
			_ = _2313_v76
			var _nw375 *C5 = New_C5_()
			_ = _nw375
			_nw375.Ctor__((_this).F9(), _2309_cf26)
			_2313_v76 = _nw375
			var _2314_v77 _dafny.Array
			_ = _2314_v77
			var _nw376 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw376
			_2314_v77 = _nw376
			var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_2314_v77), 0))
			_ = _index395
			(_2314_v77).ArraySet1(_2220_v0, (_index395).Int())
			var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_2314_v77), 0))
			_ = _index396
			(_2314_v77).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2258_v29).Cardinality()), _2310_cf25), (_index396).Int())
			_2310_cf25 = (_dafny.Zero).Minus((_2220_v0).Minus(_2220_v0))
		} else if _source29.Is_DC13() {
			var _2315___mcc_h2 _dafny.Array = _source29.Get_().(D5_DC13).Cf24
			_ = _2315___mcc_h2
			var _2316_cf24 _dafny.Array = _2315___mcc_h2
			_ = _2316_cf24
			var _2317_v78 _dafny.CodePoint
			_ = _2317_v78
			_2317_v78 = _dafny.CodePoint('c')
			var _2318_v79 _dafny.Map
			_ = _2318_v79
			_2318_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm36(globalState), _2220_v0)
			var _2319_v80 _dafny.Set
			_ = _2319_v80
			_2319_v80 = _dafny.SetOf((_this).Fm3((_dafny.Zero).Minus(_2220_v0), _2220_v0, _2317_v78, (_2318_v79).Cardinality(), globalState))
			var _2320_v81 _dafny.Sequence
			_ = _2320_v81
			_2320_v81 = _dafny.SeqOf((func() _dafny.Set {
				if (_this).F10() {
					return _dafny.SetOf(_2220_v0, _2220_v0)
				}
				return _dafny.SetOf(_2220_v0, _2220_v0)
			})(), _dafny.SetOf(_dafny.IntOfInt64(803), _2220_v0), _2319_v80)
			_2320_v81 = _dafny.Companion_Sequence_.Concatenate(_2320_v81, _2320_v81)
			_2220_v0 = _2220_v0
			if (_this).Fm2(globalState) {
				r0 = _2220_v0
				var _2321_v82 _dafny.Array
				_ = _2321_v82
				var _nw377 _dafny.Array = _dafny.NewArrayWithValue(Companion_D15_.Default(), _dafny.IntOfInt64(5))
				_ = _nw377
				_2321_v82 = _nw377
				var _2322_v83 _dafny.Map
				_ = _2322_v83
				_2322_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(_2220_v0, _2220_v0, _2317_v78, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2317_v78, !((_this).F10()))).Cardinality(), globalState), _2321_v82)
				var _2323_v84 _dafny.Array
				_ = _2323_v84
				var _nwElement0_81 _dafny.Map = (_2322_v83).Merge(_2322_v83)
				_ = _nwElement0_81
				var _nw378 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(6))
				_ = _nw378
				(_nw378).ArraySet1(_nwElement0_81, 0)
				(_nw378).ArraySet1((_2322_v83).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2220_v0, _2321_v82)), 1)
				(_nw378).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2220_v0, _2321_v82)).Merge(_2322_v83), 2)
				(_nw378).ArraySet1((_2322_v83).Merge(_2322_v83), 3)
				(_nw378).ArraySet1(_2322_v83, 4)
				(_nw378).ArraySet1(_2322_v83, 5)
				_2323_v84 = _nw378
				var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen((_2323_v84), 0))
				_ = _index397
				(_2323_v84).ArraySet1(_2322_v83, (_index397).Int())
				var _2324_v85 _dafny.Array
				_ = _2324_v85
				var _nw379 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw379
				_2324_v85 = _nw379
				var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen((_2323_v84), 0))
				_ = _index398
				var _rhs371 _dafny.Map = (func() _dafny.Map {
					if (_this).F10() {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2220_v0, _2321_v82)
					}
					return _2322_v83
				})()
				_ = _rhs371
				var _rhs372 _dafny.CodePoint = _2317_v78
				_ = _rhs372
				var _rhs373 _dafny.Array = _2324_v85
				_ = _rhs373
				var _rhs374 _dafny.Int = _2220_v0
				_ = _rhs374
				var _lhs274 _dafny.Array = _2323_v84
				_ = _lhs274
				var _lhs275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen((_2323_v84), 0))
				_ = _lhs275
				(_lhs274).ArraySet1(_rhs371, (_lhs275).Int())
				_2317_v78 = _rhs372
				_2324_v85 = _rhs373
				_2220_v0 = _rhs374
				var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_2316_cf24), 0))
				_ = _index399
				(_2316_cf24).ArraySet1(false, (_index399).Int())
				var _2325_v86 _dafny.Map
				_ = _2325_v86
				_2325_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2324_v85, true)
				var _2326_v87 _dafny.Map
				_ = _2326_v87
				_2326_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(427))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg121 _dafny.Int) interface{} {
						return coer121(arg121)
					}
				}((func(_2327_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2328_i7 _dafny.Int) _dafny.CodePoint {
						return _2327_v78
					}
				})(_2317_v78))), (func() bool {
					if (_2325_v86).Contains(_2324_v85) {
						return (_2325_v86).Get(_2324_v85).(bool)
					}
					return (_this).F10()
				})())
				var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_2316_cf24), 0))
				_ = _index400
				var _rhs375 bool = (func() bool {
					if (_2326_v87).Contains(_dafny.Companion_Sequence_.Update((Companion_D1_.Create_DC3_(_2258_v29)).Dtor_cf5(), (Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32(((Companion_D1_.Create_DC3_(_2258_v29)).Dtor_cf5()).Cardinality()))).Uint32(), _dafny.CodePoint('d'))) {
						return (_2326_v87).Get(_dafny.Companion_Sequence_.Update((Companion_D1_.Create_DC3_(_2258_v29)).Dtor_cf5(), (Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32(((Companion_D1_.Create_DC3_(_2258_v29)).Dtor_cf5()).Cardinality()))).Uint32(), _dafny.CodePoint('d'))).(bool)
					}
					return (_this).F10()
				})()
				_ = _rhs375
				var _rhs376 bool = !((!((_this).F10())) || ((_this).F10()))
				_ = _rhs376
				var _rhs377 _dafny.Int = _2220_v0
				_ = _rhs377
				var _rhs378 bool = (_this).F10()
				_ = _rhs378
				var _lhs276 _dafny.Array = _2316_cf24
				_ = _lhs276
				var _lhs277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_2316_cf24), 0))
				_ = _lhs277
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				var _lhs279 *GlobalState = globalState
				_ = _lhs279
				(_lhs276).ArraySet1(_rhs375, (_lhs277).Int())
				_lhs278.F8 = _rhs376
				r0 = _rhs377
				_lhs279.F8 = _rhs378
				var _2329_v88 *C6
				_ = _2329_v88
				var _nw380 *C6 = New_C6_()
				_ = _nw380
				_nw380.Ctor__((_this).F9(), (_this).F10())
				_2329_v88 = _nw380
				var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_2324_v85), 0))
				_ = _index401
				(_2324_v85).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_2220_v0, _2220_v0)), (_index401).Int())
				var _2330_v89 _dafny.Map
				_ = _2330_v89
				_2330_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2258_v29)).Cardinality())
				var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_2324_v85), 0))
				_ = _index402
				(_2324_v85).ArraySet1((func() _dafny.Int {
					if (_2330_v89).Contains((_this).F10()) {
						return (_2330_v89).Get((_this).F10()).(_dafny.Int)
					}
					return _2220_v0
				})(), (_index402).Int())
			} else {
				r1 = (func() _dafny.Sequence {
					if (_this).F10() {
						return _2258_v29
					}
					return _dafny.Companion_Sequence_.Concatenate(_2258_v29, _2258_v29)
				})()
				var _2331_v90 D22
				_ = _2331_v90
				_2331_v90 = Companion_D22_.Create_DC57_()
				_2331_v90 = _2331_v90
				r0 = (_dafny.Zero).Minus(((func() _dafny.Int {
					if (_this).F10() {
						return _2220_v0
					}
					return _2220_v0
				})()).Plus(_dafny.IntOfUint32((_2258_v29).Cardinality())))
				var _2332_v91 *C1
				_ = _2332_v91
				var _nw381 *C1 = New_C1_()
				_ = _nw381
				_nw381.Ctor__((_this).F10(), _dafny.UnicodeSeqOfUtf8Bytes("v"))
				_2332_v91 = _nw381
				var _2333_v92 _dafny.Sequence
				_ = _2333_v92
				_2333_v92 = _dafny.SeqOf(_2332_v91, _2332_v91)
				var _2334_v93 _dafny.Set
				_ = _2334_v93
				_2334_v93 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2333_v92, (Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2333_v92).Cardinality()))).Uint32(), _2332_v91), _2333_v92), _dafny.Companion_Sequence_.Concatenate(_2333_v92, _dafny.Companion_Sequence_.Update(_2333_v92, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2258_v29).Cardinality()), _dafny.IntOfUint32((_2333_v92).Cardinality()))).Uint32(), _2332_v91)))
				_2334_v93 = _2334_v93
				_2220_v0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2320_v81, _2320_v81)).Cardinality())))
			}
			(globalState).F8 = (_this).F10()
		} else {
			var _2335___mcc_h3 D5 = _source29.Get_().(D5_DC15).Cf27
			_ = _2335___mcc_h3
			var _2336_cf27 D5 = _2335___mcc_h3
			_ = _2336_cf27
			var _2337_v94 _dafny.Sequence
			_ = _2337_v94
			_2337_v94 = _dafny.SeqOf((_dafny.Zero).Minus(_2220_v0), _2220_v0)
			_2337_v94 = _2337_v94
			var _2338_v95 *C1
			_ = _2338_v95
			var _nw382 *C1 = New_C1_()
			_ = _nw382
			_nw382.Ctor__((_this).F10(), _2258_v29)
			_2338_v95 = _nw382
			_2338_v95 = (func() *C1 {
				if (_this).Fm2(globalState) {
					return _2338_v95
				}
				return _2338_v95
			})()
			var _2339_v96 _dafny.CodePoint
			_ = _2339_v96
			_2339_v96 = _dafny.CodePoint('i')
			var _2340_v97 _dafny.Map
			_ = _2340_v97
			_2340_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm36(globalState), _2339_v96)
			(globalState).F8 = ((_2340_v97).Cardinality()).Cmp(_2220_v0) != 0
			_2220_v0 = (_2220_v0).Minus((_dafny.IntOfInt64(33)).Minus(_2220_v0))
		}
		var _2341_v98 *C2
		_ = _2341_v98
		var _nw383 *C2 = New_C2_()
		_ = _nw383
		_nw383.Ctor__((_2302_v69).Select((Companion_Default___.SafeIndex(_2220_v0, _dafny.IntOfUint32((_2302_v69).Cardinality()))).Uint32()).(bool), _2303_v70, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10())).Merge((_this).F9()), (_this).F10())
		_2341_v98 = _nw383
		r0 = (Companion_Default___.Fm46(false, (_2222_v2).IsDisjointFrom(_2222_v2), _2341_v98.F14, globalState)).Cardinality()
		var _2342_v99 _dafny.CodePoint
		_ = _2342_v99
		_2342_v99 = _dafny.CodePoint('g')
		r1 = _dafny.Companion_Sequence_.Update(_2258_v29, (Companion_Default___.SafeIndex(((_this).F11()).Cardinality(), _dafny.IntOfUint32((_2258_v29).Cardinality()))).Uint32(), _2342_v99)
		return r0, r1
	}
}
func (_this *C7) M3(p0 _dafny.Int, globalState *GlobalState) (_dafny.MultiSet, bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _2343_v0 _dafny.CodePoint
		_ = _2343_v0
		_2343_v0 = _dafny.CodePoint('x')
		var _2344_v1 D6
		_ = _2344_v1
		_2344_v1 = Companion_D6_.Create_DC18_(_2343_v0, p0, p0)
		var _2345_v2 D18
		_ = _2345_v2
		_2345_v2 = Companion_D18_.Create_DC47_((_this).F10(), _2344_v1)
		var _2346_v3 _dafny.Map
		_ = _2346_v3
		_2346_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(958))), _2345_v2)
		var _2347_v4 _dafny.Map
		_ = _2347_v4
		_2347_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), p0)
		var _2348_v5 _dafny.MultiSet
		_ = _2348_v5
		_2348_v5 = _dafny.MultiSetOf((_this).F10(), (_this).F10(), (_this).F10(), (_this).F10())
		var _2349_v6 _dafny.MultiSet
		_ = _2349_v6
		_2349_v6 = _dafny.MultiSetOf((_2348_v5).Cardinality())
		var _2350_v7 _dafny.Map
		_ = _2350_v7
		_2350_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2347_v4).Cardinality(), (_2349_v6).Cardinality())
		var _rhs379 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2350_v7).Cardinality(), _2345_v2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2345_v2))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2345_v2)).Merge(_2346_v3))
		_ = _rhs379
		var _rhs380 _dafny.Int = p0
		_ = _rhs380
		_2346_v3 = _rhs379
		r2 = _rhs380
		var _2351_i0 _dafny.Int
		_ = _2351_i0
		_2351_i0 = _dafny.Zero
		{
			for (_this).F10() {
				{
					if (_2351_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_2351_i0 = (_2351_i0).Plus(_dafny.One)
					var _2352_v8 _dafny.Array
					_ = _2352_v8
					var _len0_68 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_68
					var _nw384 _dafny.Array
					_ = _nw384
					if _len0_68.Cmp(_dafny.Zero) == 0 {
						_nw384 = _dafny.NewArray(_len0_68)
					} else {
						var _init68 func(_dafny.Int) _dafny.Sequence = func(_2353_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.UnicodeSeqOfUtf8Bytes("dvupvsxfh")
						}
						_ = _init68
						var _element0_68 = _init68(_dafny.Zero)
						_ = _element0_68
						_nw384 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
						(_nw384).ArraySet1(_element0_68, 0)
						var _nativeLen0_68 = (_len0_68).Int()
						_ = _nativeLen0_68
						for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
							(_nw384).ArraySet1(_init68(_dafny.IntOf(_i0_68)), _i0_68)
						}
					}
					_2352_v8 = _nw384
					var _2354_v9 _dafny.Map
					_ = _2354_v9
					_2354_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2352_v8)
					_2354_v9 = (_2354_v9).Update((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, p0)), (func() _dafny.Array {
						if (_this).F10() {
							return _2352_v8
						}
						return _2352_v8
					})())
					r2 = (_dafny.IntOfInt64(-261)).Minus(p0)
					r1 = (_this).F10()
					var _2355_v10 *C5
					_ = _2355_v10
					var _nw385 *C5 = New_C5_()
					_ = _nw385
					_nw385.Ctor__((_this).F9(), false)
					_2355_v10 = _nw385
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		r2 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_this).F10() {
				return p0
			}
			return p0
		})(), Companion_Default___.SafeDivisionInt(p0, p0))
		if (_this).F10() {
			r2 = (_this).Fm1((_this).F9(), globalState)
			var _2356_v11 _dafny.Array
			_ = _2356_v11
			var _nw386 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw386
			_2356_v11 = _nw386
			var _2357_v12 _dafny.Array
			_ = _2357_v12
			var _nw387 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw387
			_2357_v12 = _nw387
			var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_2357_v12), 0))
			_ = _index403
			(_2357_v12).ArraySet1(_dafny.SeqOf((_this).F10(), (_this).F10()), (_index403).Int())
			var _2358_v13 *C4
			_ = _2358_v13
			var _nw388 *C4 = New_C4_()
			_ = _nw388
			_nw388.Ctor__(((_this).F9()).Update((_this).F10(), (_this).F10()), (_this).F10())
			_2358_v13 = _nw388
			var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_2356_v11), 0))
			_ = _index404
			(_2356_v11).ArraySet1(p0, (_index404).Int())
			var _2359_v14 _dafny.MultiSet
			_ = _2359_v14
			_2359_v14 = _dafny.MultiSetOf(_2343_v0)
			var _2360_v15 _dafny.Sequence
			_ = _2360_v15
			_2360_v15 = _dafny.SeqOf((p0).Cmp((_dafny.Zero).Minus(p0)) >= 0, (_2359_v14).IsDisjointFrom(_2359_v14))
			var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_2357_v12), 0))
			_ = _index405
			var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_2356_v11), 0))
			_ = _index406
			var _rhs381 _dafny.Array = _2356_v11
			_ = _rhs381
			var _rhs382 _dafny.Sequence = _2360_v15
			_ = _rhs382
			var _rhs383 *C4 = _2358_v13
			_ = _rhs383
			var _rhs384 _dafny.Int = ((func() _dafny.Int {
				if !(false) {
					return p0
				}
				return p0
			})()).Times((func() _dafny.Int {
				if (_2349_v6).Contains(p0) {
					return (_2349_v6).Multiplicity(p0)
				}
				return p0
			})())
			_ = _rhs384
			var _lhs280 _dafny.Array = _2357_v12
			_ = _lhs280
			var _lhs281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_2357_v12), 0))
			_ = _lhs281
			var _lhs282 _dafny.Array = _2356_v11
			_ = _lhs282
			var _lhs283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_2356_v11), 0))
			_ = _lhs283
			_2356_v11 = _rhs381
			(_lhs280).ArraySet1(_rhs382, (_lhs281).Int())
			_2358_v13 = _rhs383
			(_lhs282).ArraySet1(_rhs384, (_lhs283).Int())
			var _2361_v16 _dafny.Map
			_ = _2361_v16
			_2361_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(155)).Cmp((_dafny.Zero).Minus((_2356_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_2356_v11), 0))).Int()).(_dafny.Int))) <= 0, _2343_v0)
			_2343_v0 = (func() _dafny.CodePoint {
				if (_2361_v16).Contains(false) {
					return (_2361_v16).Get(false).(_dafny.CodePoint)
				}
				return _2343_v0
			})()
			var _2362_v17 _dafny.Map
			_ = _2362_v17
			_2362_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_this).F10() {
					return (_this).F10()
				}
				return true
			})(), _2348_v5)
			_2362_v17 = (_2362_v17).Update((_this).F10(), _2348_v5)
			var _2363_v18 _dafny.Array
			_ = _2363_v18
			var _len0_69 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_69
			var _nw389 _dafny.Array
			_ = _nw389
			if _len0_69.Cmp(_dafny.Zero) == 0 {
				_nw389 = _dafny.NewArray(_len0_69)
			} else {
				var _init69 func(_dafny.Int) bool = func(_2364_i2 _dafny.Int) bool {
					return (_this).F10()
				}
				_ = _init69
				var _element0_69 = _init69(_dafny.Zero)
				_ = _element0_69
				_nw389 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
				(_nw389).ArraySet1(_element0_69, 0)
				var _nativeLen0_69 = (_len0_69).Int()
				_ = _nativeLen0_69
				for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
					(_nw389).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
				}
			}
			_2363_v18 = _nw389
			_2363_v18 = (func() _dafny.Array {
				if true {
					return _2363_v18
				}
				return _2363_v18
			})()
		} else {
			var _2365_v19 D19
			_ = _2365_v19
			_2365_v19 = Companion_D19_.Create_DC49_(true)
			var _2366_v20 D1
			_ = _2366_v20
			_2366_v20 = Companion_D1_.Create_DC4_((_2365_v19).Dtor_cf93(), (_this).F10(), p0)
			var _2367_v21 _dafny.Sequence
			_ = _2367_v21
			_2367_v21 = _dafny.SeqOf((_this).F10())
			var _2368_v22 _dafny.Map
			_ = _2368_v22
			_2368_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2367_v21)
			var _2369_v23 _dafny.Set
			_ = _2369_v23
			_2369_v23 = _dafny.SetOf(_dafny.CodePoint('d'))
			var _2370_v24 _dafny.Map
			_ = _2370_v24
			_2370_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2366_v20).Dtor_cf8()).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_2368_v22).Contains(p0) {
					return (_2368_v22).Get(p0).(_dafny.Sequence)
				}
				return _2367_v21
			})()).Cardinality())), (_2369_v23).IsSubsetOf(_2369_v23))
			_2370_v24 = (_2370_v24).Update(p0, (_this).F10())
			var _2371_v25 _dafny.Array
			_ = _2371_v25
			var _nw390 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw390
			_2371_v25 = _nw390
			var _2372_v26 D6
			_ = _2372_v26
			_2372_v26 = Companion_D6_.Create_DC16_((_this).F9())
			var _2373_v27 _dafny.Set
			_ = _2373_v27
			_2373_v27 = _dafny.SetOf(_dafny.IntOfInt64(-117))
			var _2374_v28 D23
			_ = _2374_v28
			_2374_v28 = Companion_D23_.Create_DC59_(_2372_v26, (_this).F10(), (_2373_v27).Cardinality(), (Companion_D5_.Create_DC13_((Companion_D5_.Create_DC13_(_2371_v25)).Dtor_cf24())).Dtor_cf24())
			var _2375_v29 _dafny.Array
			_ = _2375_v29
			var _nwElement0_82 _dafny.Array = _2371_v25
			_ = _nwElement0_82
			var _nw391 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(5))
			_ = _nw391
			(_nw391).ArraySet1(_nwElement0_82, 0)
			(_nw391).ArraySet1(_2371_v25, 1)
			(_nw391).ArraySet1(_2371_v25, 2)
			(_nw391).ArraySet1((_2374_v28).Dtor_cf109(), 3)
			(_nw391).ArraySet1(_2371_v25, 4)
			_2375_v29 = _nw391
			_2375_v29 = _2375_v29
			var _2376_v30 _dafny.Sequence
			_ = _2376_v30
			_2376_v30 = _dafny.SeqOf(p0)
			var _2377_v31 _dafny.Array
			_ = _2377_v31
			var _nwElement0_83 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2376_v30, _2376_v30)
			_ = _nwElement0_83
			var _nw392 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(2))
			_ = _nw392
			(_nw392).ArraySet1(_nwElement0_83, 0)
			(_nw392).ArraySet1(_2376_v30, 1)
			_2377_v31 = _nw392
			var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_2377_v31), 0))
			_ = _index407
			(_2377_v31).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2376_v30, _dafny.SeqOf(p0)), (_index407).Int())
			var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_2377_v31), 0))
			_ = _index408
			(_2377_v31).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer122 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg122 _dafny.Int) interface{} {
					return coer122(arg122)
				}
			}(func(_2378_i3 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-744)
			})), _2376_v30), (_index408).Int())
			_2370_v24 = (_2370_v24).Update((_dafny.Zero).Minus(p0), (_this).F10())
			_2367_v21 = _dafny.Companion_Sequence_.Update(_2367_v21, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2367_v21).Cardinality()))).Uint32(), (_this).F10())
		}
		var _hi8 _dafny.Int = _dafny.IntOfInt64(643)
		_ = _hi8
		for _2379_i4 := p0; _2379_i4.Cmp(_hi8) < 0; _2379_i4 = _2379_i4.Plus(_dafny.One) {
			var _2380_v32 _dafny.Map
			_ = _2380_v32
			_2380_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _2379_i4)
			var _2381_v33 _dafny.Sequence
			_ = _2381_v33
			_2381_v33 = _dafny.SeqOf((_2380_v32).Cardinality())
			var _source30 D16 = Companion_D16_.Create_DC42_((_this).F10(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), !(false))).Cardinality(), (func() _dafny.Int {
				if (_this).F10() {
					return (_dafny.MultiSetFromSeq(_2381_v33)).Cardinality()
				}
				return _2379_i4
			})())
			_ = _source30
			if _source30.Is_DC41() {
				var _2382___mcc_h0 bool = _source30.Get_().(D16_DC41).Cf77
				_ = _2382___mcc_h0
				var _2383___mcc_h1 _dafny.Int = _source30.Get_().(D16_DC41).Cf78
				_ = _2383___mcc_h1
				var _2384___mcc_h2 bool = _source30.Get_().(D16_DC41).Cf79
				_ = _2384___mcc_h2
				var _2385_cf79 bool = _2384___mcc_h2
				_ = _2385_cf79
				var _2386_cf78 _dafny.Int = _2383___mcc_h1
				_ = _2386_cf78
				var _2387_cf77 bool = _2382___mcc_h0
				_ = _2387_cf77
				var _2388_v34 _dafny.Array
				_ = _2388_v34
				var _nw393 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw393
				_2388_v34 = _nw393
				var _2389_v35 T1
				_ = _2389_v35
				var _nw394 *C2 = New_C2_()
				_ = _nw394
				_nw394.Ctor__(_2385_cf79, _2388_v34, (_this).F9(), _2387_cf77)
				_2389_v35 = _nw394
				var _2390_v36 _dafny.Sequence
				_ = _2390_v36
				_2390_v36 = _dafny.SeqOf(_2389_v35, _2389_v35)
				_2390_v36 = _2390_v36
				var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2388_v34), 0))
				_ = _index409
				(_2388_v34).ArraySet1(_2387_cf77, (_index409).Int())
				var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2388_v34), 0))
				_ = _index410
				(_2388_v34).ArraySet1(!((_this).F10()), (_index410).Int())
				var _2391_v37 _dafny.Array
				_ = _2391_v37
				var _nwElement0_84 _dafny.CodePoint = _2343_v0
				_ = _nwElement0_84
				var _nw395 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(11))
				_ = _nw395
				(_nw395).ArraySet1CodePoint(_nwElement0_84, 0)
				(_nw395).ArraySet1CodePoint(_2343_v0, 1)
				(_nw395).ArraySet1CodePoint(_2343_v0, 2)
				(_nw395).ArraySet1CodePoint(_2343_v0, 3)
				(_nw395).ArraySet1CodePoint(_2343_v0, 4)
				(_nw395).ArraySet1CodePoint(_2343_v0, 5)
				(_nw395).ArraySet1CodePoint(_2343_v0, 6)
				(_nw395).ArraySet1CodePoint(_2343_v0, 7)
				(_nw395).ArraySet1CodePoint(_2343_v0, 8)
				(_nw395).ArraySet1CodePoint(_2343_v0, 9)
				(_nw395).ArraySet1CodePoint(_dafny.CodePoint('y'), 10)
				_2391_v37 = _nw395
				var _2392_v38 D25
				_ = _2392_v38
				_2392_v38 = Companion_D25_.Create_DC61_(_2391_v37)
				var _2393_v39 _dafny.Sequence
				_ = _2393_v39
				_2393_v39 = _dafny.SeqOf(_2391_v37, _2391_v37, _2391_v37, _2391_v37)
				var _2394_v40 _dafny.Map
				_ = _2394_v40
				_2394_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2388_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.ArrayLen((_2388_v34), 0))).Int()).(bool), _2380_v32)
				var _2395_v41 _dafny.Array
				_ = _2395_v41
				var _nwElement0_85 _dafny.Array = _2391_v37
				_ = _nwElement0_85
				var _nw396 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(13))
				_ = _nw396
				(_nw396).ArraySet1(_nwElement0_85, 0)
				(_nw396).ArraySet1((_2392_v38).Dtor_cf111(), 1)
				(_nw396).ArraySet1(_2391_v37, 2)
				(_nw396).ArraySet1(_2391_v37, 3)
				(_nw396).ArraySet1(_2391_v37, 4)
				(_nw396).ArraySet1((_2393_v39).Select((Companion_Default___.SafeIndex((_2394_v40).Cardinality(), _dafny.IntOfUint32((_2393_v39).Cardinality()))).Uint32()).(_dafny.Array), 5)
				(_nw396).ArraySet1(_2391_v37, 6)
				(_nw396).ArraySet1(_2391_v37, 7)
				(_nw396).ArraySet1(_2391_v37, 8)
				(_nw396).ArraySet1(_2391_v37, 9)
				(_nw396).ArraySet1(_2391_v37, 10)
				(_nw396).ArraySet1(_2391_v37, 11)
				(_nw396).ArraySet1((func() _dafny.Array {
					if (_2389_v35).F10() {
						return _2391_v37
					}
					return _2391_v37
				})(), 12)
				_2395_v41 = _nw396
				var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_2395_v41), 0))
				_ = _index411
				(_2395_v41).ArraySet1((func() _dafny.Array {
					if _2385_cf79 {
						return _2391_v37
					}
					return _2391_v37
				})(), (_index411).Int())
				var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_2395_v41), 0))
				_ = _index412
				var _rhs385 _dafny.Int = _2379_i4
				_ = _rhs385
				var _rhs386 _dafny.Array = _2391_v37
				_ = _rhs386
				var _lhs284 _dafny.Array = _2395_v41
				_ = _lhs284
				var _lhs285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_2395_v41), 0))
				_ = _lhs285
				r2 = _rhs385
				(_lhs284).ArraySet1(_rhs386, (_lhs285).Int())
				var _2396_v42 *C4
				_ = _2396_v42
				var _nw397 *C4 = New_C4_()
				_ = _nw397
				_nw397.Ctor__((_2389_v35).F9(), (_this).F10())
				_2396_v42 = _nw397
			} else if _source30.Is_DC42() {
				var _2397___mcc_h3 bool = _source30.Get_().(D16_DC42).Cf80
				_ = _2397___mcc_h3
				var _2398___mcc_h4 _dafny.Int = _source30.Get_().(D16_DC42).Cf81
				_ = _2398___mcc_h4
				var _2399___mcc_h5 _dafny.Int = _source30.Get_().(D16_DC42).Cf82
				_ = _2399___mcc_h5
				var _2400_cf82 _dafny.Int = _2399___mcc_h5
				_ = _2400_cf82
				var _2401_cf81 _dafny.Int = _2398___mcc_h4
				_ = _2401_cf81
				var _2402_cf80 bool = _2397___mcc_h3
				_ = _2402_cf80
				var _2403_v43 *C1
				_ = _2403_v43
				var _nw398 *C1 = New_C1_()
				_ = _nw398
				_nw398.Ctor__(_2402_cf80, _dafny.UnicodeSeqOfUtf8Bytes("knxgb"))
				_2403_v43 = _nw398
				var _2404_v44 D26
				_ = _2404_v44
				_2404_v44 = Companion_D26_.Create_DC64_(_2403_v43)
				var _2405_v45 _dafny.Set
				_ = _2405_v45
				_2405_v45 = _dafny.SetOf(_2403_v43, _2403_v43, _2403_v43, _2403_v43, (_2404_v44).Dtor_cf114())
				var _2406_v47 _dafny.Array
				_ = _2406_v47
				var _nwElement0_86 bool = ((_this).F10()) && ((_this).F10())
				_ = _nwElement0_86
				var _nw399 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(25))
				_ = _nw399
				(_nw399).ArraySet1(_nwElement0_86, 0)
				(_nw399).ArraySet1(_2402_cf80, 1)
				(_nw399).ArraySet1(_2402_cf80, 2)
				(_nw399).ArraySet1(!(_2402_cf80) || (_2402_cf80), 3)
				(_nw399).ArraySet1(true, 4)
				(_nw399).ArraySet1((_2405_v45).IsDisjointFrom(_dafny.SetOf(_2403_v43)), 5)
				(_nw399).ArraySet1(false, 6)
				(_nw399).ArraySet1((_this).F10(), 7)
				(_nw399).ArraySet1((_2403_v43).F12(), 8)
				(_nw399).ArraySet1(true, 9)
				(_nw399).ArraySet1((_this).F10(), 10)
				(_nw399).ArraySet1((func() _dafny.Set {
					var _coll77 = _dafny.NewBuilder()
					_ = _coll77
					for _iter87 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(707), _dafny.IntOfInt64(683))); ; {
						_compr_77, _ok87 := _iter87()
						if !_ok87 {
							break
						}
						var _2407_v46 _dafny.Int
						_2407_v46 = interface{}(_compr_77).(_dafny.Int)
						if ((_dafny.IntOfInt64(707)).Cmp(_2407_v46) <= 0) && ((_2407_v46).Cmp(_dafny.IntOfInt64(683)) < 0) {
							_coll77.Add(Companion_Default___.SafeModuloInt(_2407_v46, _2379_i4))
						}
					}
					return _coll77.ToSet()
				}()).IsSubsetOf(Companion_Default___.Fm25(_2401_cf81, _2400_cf82, globalState)), 11)
				(_nw399).ArraySet1((_2400_cf82).Cmp(p0) <= 0, 12)
				(_nw399).ArraySet1(_2402_cf80, 13)
				(_nw399).ArraySet1((_2403_v43).F12(), 14)
				(_nw399).ArraySet1(!(_2402_cf80) || (_2402_cf80), 15)
				(_nw399).ArraySet1((_2401_cf81).Cmp((_this).Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2403_v43).F12(), _2402_cf80), globalState)) == 0, 16)
				(_nw399).ArraySet1((_this).F10(), 17)
				(_nw399).ArraySet1(_2402_cf80, 18)
				(_nw399).ArraySet1((_2403_v43).F12(), 19)
				(_nw399).ArraySet1(_2402_cf80, 20)
				(_nw399).ArraySet1((func() bool {
					if _2402_cf80 {
						return (_this).F10()
					}
					return !((_this).F10())
				})(), 21)
				(_nw399).ArraySet1(_2402_cf80, 22)
				(_nw399).ArraySet1((_this).F10(), 23)
				(_nw399).ArraySet1((_2402_cf80) || ((_this).F10()), 24)
				_2406_v47 = _nw399
				_2406_v47 = _2406_v47
				(globalState).F8 = _2402_cf80
				var _2408_v48 _dafny.Map
				_ = _2408_v48
				_2408_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if !((_2403_v43).Fm32(_2379_i4, _2343_v0, globalState)) {
						return _dafny.IntOfInt64(630)
					}
					return _2401_cf81
				})(), _2406_v47)
				_2406_v47 = (func() _dafny.Array {
					if (_2408_v48).Contains(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(524), p0)) {
						return (_2408_v48).Get(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(524), p0)).(_dafny.Array)
					}
					return _2406_v47
				})()
				var _2409_v49 _dafny.Sequence
				_ = _2409_v49
				_2409_v49 = _dafny.SeqOf((_this).F10(), _2402_cf80)
				var _2410_v50 _dafny.Map
				_ = _2410_v50
				_2410_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('c'), (_this).F10())
				var _2411_v51 *C4
				_ = _2411_v51
				var _nw400 *C4 = New_C4_()
				_ = _nw400
				_nw400.Ctor__(Companion_Default___.Fm48((_this).Fm0(_2409_v49, (_2403_v43).F12(), (_2410_v50).Cardinality(), (_this).F10(), globalState), globalState), _2402_cf80)
				_2411_v51 = _nw400
				var _2412_v52 _dafny.Map
				_ = _2412_v52
				_2412_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2403_v43).F13(), _dafny.IntOfUint32((_2409_v49).Cardinality()))
				var _2413_v53 T0
				_ = _2413_v53
				var _nw401 *C1 = New_C1_()
				_ = _nw401
				_nw401.Ctor__((_dafny.SetOf(_2401_cf81, (_2381_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2409_v49).Cardinality()), _dafny.IntOfUint32((_2381_v33).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_2412_v52).Contains((_2403_v43).F13()) {
						return (_2412_v52).Get((_2403_v43).F13()).(_dafny.Int)
					}
					return p0
				})()))))).Contains(_2401_cf81), (_2403_v43).F13())
				_2413_v53 = _nw401
				var _2414_v54 _dafny.Map
				_ = _2414_v54
				_2414_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2403_v43).F12(), _2406_v47)
				var _2415_v55 _dafny.Map
				_ = _2415_v55
				_2415_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_2403_v43).F13()).Cardinality()), (_2403_v43).F13())
				var _rhs387 *C4 = _2411_v51
				_ = _rhs387
				var _rhs388 T0 = _2413_v53
				_ = _rhs388
				var _rhs389 _dafny.MultiSet = ((_dafny.MultiSetOf(_2401_cf81, _2379_i4, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_2415_v55).Contains((_2381_v33).Select((Companion_Default___.SafeIndex(_2400_cf82, _dafny.IntOfUint32((_2381_v33).Cardinality()))).Uint32()).(_dafny.Int)) {
						return (_2415_v55).Get((_2381_v33).Select((Companion_Default___.SafeIndex(_2400_cf82, _dafny.IntOfUint32((_2381_v33).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Sequence)
					}
					return (_2403_v43).F13()
				})()).Cardinality()), _2379_i4)).Union(_dafny.MultiSetOf(_2379_i4))).Intersection(Companion_Default___.Fm46(_2402_cf80, (_2403_v43).F12(), !((_2403_v43).F12()), globalState))
				_ = _rhs389
				var _rhs390 _dafny.Int = ((func() _dafny.Int {
					if (_2411_v51).Fm0(_2409_v49, (_2403_v43).F12(), _2379_i4, (_2403_v43).F12(), globalState) {
						return _2379_i4
					}
					return _2400_cf82
				})()).Times(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfUint32(((_2403_v43).F13()).Cardinality())))
				_ = _rhs390
				var _rhs391 _dafny.Map = _2414_v54
				_ = _rhs391
				_2411_v51 = _rhs387
				_2413_v53 = _rhs388
				_2349_v6 = _rhs389
				_2401_cf81 = _rhs390
				_2414_v54 = _rhs391
			} else {
				var _2416___mcc_h6 _dafny.Sequence = _source30.Get_().(D16_DC40).Cf76
				_ = _2416___mcc_h6
				var _2417_cf76 _dafny.Sequence = _2416___mcc_h6
				_ = _2417_cf76
				var _2418_v56 *C0
				_ = _2418_v56
				var _nw402 *C0 = New_C0_()
				_ = _nw402
				_nw402.Ctor__()
				_2418_v56 = _nw402
				_2348_v5 = _2348_v5
				var _2419_v57 *C0
				_ = _2419_v57
				var _nw403 *C0 = New_C0_()
				_ = _nw403
				_nw403.Ctor__()
				_2419_v57 = _nw403
				r2 = p0
			}
			_2349_v6 = (_2349_v6).Difference(_2349_v6)
			if !((_this).F10()) {
				var _2420_v58 _dafny.MultiSet
				_ = _2420_v58
				_2420_v58 = _dafny.MultiSetOf(_2343_v0, _2343_v0)
				var _2421_v59 _dafny.Sequence
				_ = _2421_v59
				_2421_v59 = _dafny.SeqOf(_2343_v0, _2343_v0)
				var _2422_v60 _dafny.Map
				_ = _2422_v60
				_2422_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2420_v58)
				var _2423_v61 _dafny.Sequence
				_ = _2423_v61
				_2423_v61 = _dafny.SeqOf(_2420_v58)
				var _2424_v62 _dafny.Array
				_ = _2424_v62
				var _nw404 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw404
				_2424_v62 = _nw404
				var _2425_v63 _dafny.MultiSet
				_ = _2425_v63
				_2425_v63 = _dafny.MultiSetOf(_2424_v62)
				var _2426_v64 _dafny.MultiSet
				_ = _2426_v64
				_2426_v64 = _2420_v58
				var _2427_v65 _dafny.Array
				_ = _2427_v65
				var _nwElement0_87 _dafny.MultiSet = (_dafny.MultiSetOf(_2343_v0, _2343_v0, _2343_v0)).Difference(_2420_v58)
				_ = _nwElement0_87
				var _nw405 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(28))
				_ = _nw405
				(_nw405).ArraySet1(_nwElement0_87, 0)
				(_nw405).ArraySet1(_dafny.MultiSetFromSeq(_2421_v59), 1)
				(_nw405).ArraySet1(_2420_v58, 2)
				(_nw405).ArraySet1((func() _dafny.MultiSet {
					if (_2422_v60).Contains(p0) {
						return (_2422_v60).Get(p0).(_dafny.MultiSet)
					}
					return (func() _dafny.MultiSet {
						if (_2422_v60).Contains(_2379_i4) {
							return (_2422_v60).Get(_2379_i4).(_dafny.MultiSet)
						}
						return _2420_v58
					})()
				})(), 3)
				(_nw405).ArraySet1(_dafny.MultiSetOf(_2343_v0, _2343_v0), 4)
				(_nw405).ArraySet1(_2420_v58, 5)
				(_nw405).ArraySet1(_2420_v58, 6)
				(_nw405).ArraySet1(_2420_v58, 7)
				(_nw405).ArraySet1((_2420_v58).Update(_dafny.CodePoint('i'), Companion_Default___.Abs(_2379_i4)), 8)
				(_nw405).ArraySet1((_2423_v61).Select((Companion_Default___.SafeIndex((_2425_v63).Cardinality(), _dafny.IntOfUint32((_2423_v61).Cardinality()))).Uint32()).(_dafny.MultiSet), 9)
				(_nw405).ArraySet1((_2420_v58).Update(_2343_v0, Companion_Default___.Abs(_dafny.IntOfInt64(274))), 10)
				(_nw405).ArraySet1(_2420_v58, 11)
				(_nw405).ArraySet1(_2420_v58, 12)
				(_nw405).ArraySet1((_2426_v64), 13)
				(_nw405).ArraySet1((_2420_v58).Difference(_dafny.MultiSetFromSeq(_2421_v59)), 14)
				(_nw405).ArraySet1(_2420_v58, 15)
				(_nw405).ArraySet1((_2420_v58).Union(_dafny.MultiSetFromSeq(_2421_v59)), 16)
				(_nw405).ArraySet1(_2420_v58, 17)
				(_nw405).ArraySet1(_2420_v58, 18)
				(_nw405).ArraySet1(_2420_v58, 19)
				(_nw405).ArraySet1(_2420_v58, 20)
				(_nw405).ArraySet1(_2420_v58, 21)
				(_nw405).ArraySet1(_2420_v58, 22)
				(_nw405).ArraySet1(_2420_v58, 23)
				(_nw405).ArraySet1(_2420_v58, 24)
				(_nw405).ArraySet1((_2420_v58).Union(_2420_v58), 25)
				(_nw405).ArraySet1(_2420_v58, 26)
				(_nw405).ArraySet1(_2420_v58, 27)
				_2427_v65 = _nw405
				var _2428_v66 _dafny.Sequence
				_ = _2428_v66
				_2428_v66 = _dafny.SeqOf((_this).F10())
				var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_2427_v65), 0))
				_ = _index413
				(_2427_v65).ArraySet1((_2420_v58).Intersection(Companion_Default___.Fm57(_dafny.IntOfUint32((_2428_v66).Cardinality()), globalState)), (_index413).Int())
				var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_2427_v65), 0))
				_ = _index414
				(_2427_v65).ArraySet1(((_2420_v58).Intersection(_2420_v58)).Difference((func() _dafny.MultiSet {
					if (_this).F10() {
						return (_2423_v61).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2423_v61).Cardinality()))).Uint32()).(_dafny.MultiSet)
					}
					return _2420_v58
				})()), (_index414).Int())
				var _2429_v67 D6
				_ = _2429_v67
				_2429_v67 = Companion_D6_.Create_DC16_((_this).F9())
				var _2430_v68 *C5
				_ = _2430_v68
				var _nw406 *C5 = New_C5_()
				_ = _nw406
				_nw406.Ctor__((_2429_v67).Dtor_cf28(), (_this).F10())
				_2430_v68 = _nw406
				var _2431_v69 _dafny.Array
				_ = _2431_v69
				var _nw407 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw407
				_2431_v69 = _nw407
				var _2432_v70 _dafny.Map
				_ = _2432_v70
				_2432_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2421_v59)
				var _2433_v71 _dafny.Map
				_ = _2433_v71
				_2433_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2431_v69, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(733))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}((func(_2434_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2435_i5 _dafny.Int) _dafny.CodePoint {
						return _2434_v0
					}
				})(_2343_v0))), (func() _dafny.Sequence {
					if (_2432_v70).Contains(p0) {
						return (_2432_v70).Get(p0).(_dafny.Sequence)
					}
					return _2421_v59
				})()))
				var _rhs392 _dafny.Sequence = (func() _dafny.Sequence {
					if (_2433_v71).Contains(_2431_v69) {
						return (_2433_v71).Get(_2431_v69).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("sxeiw")
				})()
				_ = _rhs392
				var _rhs393 *C5 = _2430_v68
				_ = _rhs393
				_2421_v59 = _rhs392
				_2430_v68 = _rhs393
				_2343_v0 = _2343_v0
				var _nw408 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw408
				_2424_v62 = _nw408
				r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2343_v0), _dafny.SeqOf(_2343_v0, Companion_Default___.Fm36(globalState), _2343_v0)), _dafny.Companion_Sequence_.Concatenate(_2421_v59, _2421_v59))).Cardinality())
			} else {
				var _2436_v72 _dafny.Sequence
				_ = _2436_v72
				_2436_v72 = _dafny.UnicodeSeqOfUtf8Bytes("p")
				var _2437_v73 *C1
				_ = _2437_v73
				var _nw409 *C1 = New_C1_()
				_ = _nw409
				_nw409.Ctor__((_this).F10(), _2436_v72)
				_2437_v73 = _nw409
				var _2438_v74 D26
				_ = _2438_v74
				_2438_v74 = Companion_D26_.Create_DC64_(_2437_v73)
				var _2439_v75 *C3
				_ = _2439_v75
				var _nw410 *C3 = New_C3_()
				_ = _nw410
				_nw410.Ctor__()
				_2439_v75 = _nw410
				var _pat_let_tv53 = _2437_v73
				_ = _pat_let_tv53
				var _rhs394 _dafny.Int = _2379_i4
				_ = _rhs394
				var _rhs395 D26 = func(_pat_let56_0 D26) D26 {
					return func(_2440_dt__update__tmp_h0 D26) D26 {
						return func(_pat_let57_0 *C1) D26 {
							return func(_2441_dt__update_hcf114_h0 *C1) D26 {
								return Companion_D26_.Create_DC64_(_2441_dt__update_hcf114_h0)
							}(_pat_let57_0)
						}(_pat_let_tv53)
					}(_pat_let56_0)
				}(Companion_D26_.Create_DC64_(_2437_v73))
				_ = _rhs395
				var _rhs396 bool = (_2437_v73).F12()
				_ = _rhs396
				var _rhs397 *C3 = (func() *C3 {
					if ((_this).F10()) == ((_2437_v73).F12()) {
						return _2439_v75
					}
					return _2439_v75
				})()
				_ = _rhs397
				var _lhs286 *GlobalState = globalState
				_ = _lhs286
				r2 = _rhs394
				_2438_v74 = _rhs395
				_lhs286.F8 = _rhs396
				_2439_v75 = _rhs397
				r2 = p0
				var _2442_v76 *C0
				_ = _2442_v76
				var _nw411 *C0 = New_C0_()
				_ = _nw411
				_nw411.Ctor__()
				_2442_v76 = _nw411
				r2 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_2379_i4, _dafny.IntOfUint32(((_2437_v73).F13()).Cardinality())))
				var _2443_v77 _dafny.Array
				_ = _2443_v77
				var _len0_70 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_70
				var _nw412 _dafny.Array
				_ = _nw412
				if _len0_70.Cmp(_dafny.Zero) == 0 {
					_nw412 = _dafny.NewArray(_len0_70)
				} else {
					var _init70 func(_dafny.Int) _dafny.Int = (func(_2444_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2445_i6 _dafny.Int) _dafny.Int {
							return (_2445_i6).Times(_2444_p0)
						}
					})(p0)
					_ = _init70
					var _element0_70 = _init70(_dafny.Zero)
					_ = _element0_70
					_nw412 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
					(_nw412).ArraySet1(_element0_70, 0)
					var _nativeLen0_70 = (_len0_70).Int()
					_ = _nativeLen0_70
					for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
						(_nw412).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
					}
				}
				_2443_v77 = _nw412
				_2443_v77 = _2443_v77
			}
			var _2446_v79 _dafny.Set
			_ = _2446_v79
			_2446_v79 = _dafny.SetOf(p0)
			var _2447_v80 _dafny.Map
			_ = _2447_v80
			_2447_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_D28_.Create_DC67_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.SetOf(p0), func() _dafny.Set {
				var _coll78 = _dafny.NewBuilder()
				_ = _coll78
				for _iter88 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-767), _dafny.IntOfInt64(994))); ; {
					_compr_78, _ok88 := _iter88()
					if !_ok88 {
						break
					}
					var _2448_v78 _dafny.Int
					_2448_v78 = interface{}(_compr_78).(_dafny.Int)
					if ((_dafny.IntOfInt64(-767)).Cmp(_2448_v78) <= 0) && ((_2448_v78).Cmp(_dafny.IntOfInt64(994)) < 0) {
						_coll78.Add(Companion_Default___.SafeModuloInt(_2448_v78, _2379_i4))
					}
				}
				return _coll78.ToSet()
			}(), _2446_v79, _2446_v79, _2446_v79), (Companion_Default___.SafeIndex(_2379_i4, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf(p0), func() _dafny.Set {
				var _coll79 = _dafny.NewBuilder()
				_ = _coll79
				for _iter89 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-767), _dafny.IntOfInt64(994))); ; {
					_compr_79, _ok89 := _iter89()
					if !_ok89 {
						break
					}
					var _2449_v78 _dafny.Int
					_2449_v78 = interface{}(_compr_79).(_dafny.Int)
					if ((_dafny.IntOfInt64(-767)).Cmp(_2449_v78) <= 0) && ((_2449_v78).Cmp(_dafny.IntOfInt64(994)) < 0) {
						_coll79.Add(Companion_Default___.SafeModuloInt(_2449_v78, _2379_i4))
					}
				}
				return _coll79.ToSet()
			}(), _2446_v79, _2446_v79, _2446_v79)).Cardinality()))).Uint32(), _2446_v79))).Dtor_cf119())
			var _2450_v81 _dafny.Sequence
			_ = _2450_v81
			_2450_v81 = _dafny.SeqOf(_2446_v79)
			if !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm58(globalState), (func() _dafny.Sequence {
				if (_2447_v80).Contains(_dafny.IntOfInt64(-615)) {
					return (_2447_v80).Get(_dafny.IntOfInt64(-615)).(_dafny.Sequence)
				}
				return _2450_v81
			})()) {
				(globalState).F8 = (_this).F10()
				var _2451_v82 _dafny.Sequence
				_ = _2451_v82
				_2451_v82 = _dafny.SeqOf((_this).F11(), (_this).F11(), (_this).F11())
				var _rhs398 bool = true
				_ = _rhs398
				var _rhs399 _dafny.Int = (_dafny.IntOfUint32((_2451_v82).Cardinality())).Minus(_2379_i4)
				_ = _rhs399
				var _rhs400 bool = (_this).F10()
				_ = _rhs400
				var _rhs401 _dafny.Int = _2379_i4
				_ = _rhs401
				var _lhs287 *GlobalState = globalState
				_ = _lhs287
				r1 = _rhs398
				r2 = _rhs399
				_lhs287.F8 = _rhs400
				r2 = _rhs401
				var _2452_v83 _dafny.Sequence
				_ = _2452_v83
				_2452_v83 = _dafny.SeqOf((_this).F10())
				(globalState).F8 = (p0).Cmp(_dafny.IntOfUint32((_2452_v83).Cardinality())) <= 0
				r2 = p0
				var _2453_v84 T0
				_ = _2453_v84
				var _nw413 *C1 = New_C1_()
				_ = _nw413
				_nw413.Ctor__(!((_this).F10()), _dafny.UnicodeSeqOfUtf8Bytes("ukpsmds"))
				_2453_v84 = _nw413
				var _2454_v85 _dafny.Map
				_ = _2454_v85
				_2454_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2453_v84)
				_2454_v85 = (_2454_v85).Update((_dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if (_this).F10() {
						return _2452_v83
					}
					return _2452_v83
				})())).Cardinality(), _2453_v84)
			} else {
				r2 = (p0).Times((_this).Fm1((_this).F9(), globalState))
				var _2455_v86 _dafny.Sequence
				_ = _2455_v86
				_2455_v86 = _dafny.SeqOf((_this).F10())
				var _2456_v87 *C0
				_ = _2456_v87
				var _nw414 *C0 = New_C0_()
				_ = _nw414
				_nw414.Ctor__()
				_2456_v87 = _nw414
				var _2457_v88 _dafny.MultiSet
				_ = _2457_v88
				_2457_v88 = _dafny.MultiSetOf(_2456_v87)
				var _2458_v89 _dafny.Array
				_ = _2458_v89
				var _nwElement0_88 bool = !((_this).F10())
				_ = _nwElement0_88
				var _nw415 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(22))
				_ = _nw415
				(_nw415).ArraySet1(_nwElement0_88, 0)
				(_nw415).ArraySet1((_this).F10(), 1)
				(_nw415).ArraySet1((_this).F10(), 2)
				(_nw415).ArraySet1((_this).F10(), 3)
				(_nw415).ArraySet1((_this).F10(), 4)
				(_nw415).ArraySet1((_this).F10(), 5)
				(_nw415).ArraySet1(!(!((_this).F10())) || ((_this).F10()), 6)
				(_nw415).ArraySet1((_2379_i4).Cmp(_dafny.IntOfUint32((_2455_v86).Cardinality())) == 0, 7)
				(_nw415).ArraySet1((_this).F10(), 8)
				(_nw415).ArraySet1(!((_this).F10()), 9)
				(_nw415).ArraySet1((_this).F10(), 10)
				(_nw415).ArraySet1((_2457_v88).IsProperSubsetOf(_2457_v88), 11)
				(_nw415).ArraySet1((_this).F10(), 12)
				(_nw415).ArraySet1(!((_this).F10()), 13)
				(_nw415).ArraySet1((_this).F10(), 14)
				(_nw415).ArraySet1(!((_this).F10()), 15)
				(_nw415).ArraySet1((_this).F10(), 16)
				(_nw415).ArraySet1((_this).F10(), 17)
				(_nw415).ArraySet1((_this).F10(), 18)
				(_nw415).ArraySet1((_2455_v86).Select((Companion_Default___.SafeIndex(_2379_i4, _dafny.IntOfUint32((_2455_v86).Cardinality()))).Uint32()).(bool), 19)
				(_nw415).ArraySet1((_this).F10(), 20)
				(_nw415).ArraySet1((_this).F10(), 21)
				_2458_v89 = _nw415
				var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))
				_ = _index415
				(_2458_v89).ArraySet1((_this).F10(), (_index415).Int())
				var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))
				_ = _index416
				(_2458_v89).ArraySet1((_this).F10(), (_index416).Int())
				(globalState).F8 = ((_2458_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))).Int()).(bool)) == ((_2458_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))).Int()).(bool))
				var _2459_v90 D7
				_ = _2459_v90
				_2459_v90 = Companion_D7_.Create_DC21_((_this).F11())
				var _2460_v91 D7
				_ = _2460_v91
				_2460_v91 = Companion_D7_.Create_DC21_((_2459_v90).Dtor_cf40())
				var _2461_v92 _dafny.Map
				_ = _2461_v92
				_2461_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_2458_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))).Int()).(bool) {
						return (_this).F10()
					}
					return (_this).F10()
				})(), ((_2460_v91).Dtor_cf40()).IsProperSubsetOf((_this).F11()))
				_2461_v92 = _2461_v92
				var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))
				_ = _index417
				(_2458_v89).ArraySet1(!((_2458_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_2458_v89), 0))).Int()).(bool)) || ((_this).F10()), (_index417).Int())
			}
		}
		var _2462_v93 _dafny.Sequence
		_ = _2462_v93
		_2462_v93 = _dafny.SeqOf(p0, _dafny.IntOfInt64(-457), p0)
		var _2463_v94 _dafny.Array
		_ = _2463_v94
		var _nwElement0_89 _dafny.Int = _dafny.IntOfInt64(864)
		_ = _nwElement0_89
		var _nw416 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(10))
		_ = _nw416
		(_nw416).ArraySet1(_nwElement0_89, 0)
		(_nw416).ArraySet1(p0, 1)
		(_nw416).ArraySet1(p0, 2)
		(_nw416).ArraySet1((_2462_v93).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.Fm33(!((_this).F10()), (_this).F10(), p0, (_this).F10(), globalState)), _dafny.IntOfUint32((_2462_v93).Cardinality()))).Uint32()).(_dafny.Int), 3)
		(_nw416).ArraySet1(p0, 4)
		(_nw416).ArraySet1(p0, 5)
		(_nw416).ArraySet1(p0, 6)
		(_nw416).ArraySet1(p0, 7)
		(_nw416).ArraySet1(p0, 8)
		(_nw416).ArraySet1(p0, 9)
		_2463_v94 = _nw416
		var _2464_v95 _dafny.MultiSet
		_ = _2464_v95
		_2464_v95 = _dafny.MultiSetOf(_2463_v94, _2463_v94, _2463_v94)
		var _hi9 _dafny.Int = (p0).Minus(p0)
		_ = _hi9
		for _2465_i7 := (_dafny.Zero).Minus((_2464_v95).Cardinality()); _2465_i7.Cmp(_hi9) < 0; _2465_i7 = _2465_i7.Plus(_dafny.One) {
			var _2466_v96 *C5
			_ = _2466_v96
			var _nw417 *C5 = New_C5_()
			_ = _nw417
			_nw417.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10())).Update((_this).F10(), !((_this).F10())), !((_this).F10()))
			_2466_v96 = _nw417
			_2466_v96 = _2466_v96
			var _2467_v97 _dafny.Sequence
			_ = _2467_v97
			_2467_v97 = _dafny.UnicodeSeqOfUtf8Bytes("ax")
			var _2468_v98 _dafny.Map
			_ = _2468_v98
			_2468_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2467_v97).Cardinality()), (_this).F10())
			var _rhs402 bool = (_2466_v96).Fm18((_this).F11(), _dafny.Companion_Sequence_.Concatenate(_2467_v97, _dafny.Companion_Sequence_.Update(_2467_v97, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2467_v97).Cardinality()))).Uint32(), Companion_Default___.Fm31((_this).F10(), _2468_v98, globalState))), globalState)
			_ = _rhs402
			var _rhs403 _dafny.Int = _2465_i7
			_ = _rhs403
			var _rhs404 bool = (_this).F10()
			_ = _rhs404
			var _lhs288 *GlobalState = globalState
			_ = _lhs288
			r1 = _rhs402
			r2 = _rhs403
			_lhs288.F8 = _rhs404
			var _index418 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_2463_v94), 0))
			_ = _index418
			(_2463_v94).ArraySet1(_2465_i7, (_index418).Int())
			var _2469_v99 _dafny.Map
			_ = _2469_v99
			_2469_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_2462_v93).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2462_v93).Cardinality()))).Uint32()).(_dafny.Int))
			var _index419 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_2463_v94), 0))
			_ = _index419
			(_2463_v94).ArraySet1((_2465_i7).Minus(((_2469_v99).Cardinality()).Minus((_dafny.Zero).Minus(_2465_i7))), (_index419).Int())
			var _2470_v100 *C0
			_ = _2470_v100
			var _nw418 *C0 = New_C0_()
			_ = _nw418
			_nw418.Ctor__()
			_2470_v100 = _nw418
		}
		var _2471_v101 _dafny.Sequence
		_ = _2471_v101
		_2471_v101 = _dafny.UnicodeSeqOfUtf8Bytes("gbx")
		r0 = (Companion_Default___.Fm10((_this).F10(), _dafny.IntOfUint32((_2462_v93).Cardinality()), _2471_v101, _2350_v7, globalState)).Difference(_2349_v6)
		r1 = true
		var _2472_v102 _dafny.Sequence
		_ = _2472_v102
		_2472_v102 = _dafny.SeqOf((_this).F10())
		r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((func() bool {
			if (_this).F10() {
				return !((_this).F10())
			}
			return false
		})()), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf((func() bool {
			if (_this).F10() {
				return !((_this).F10())
			}
			return false
		})())).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Contains(_2472_v102, !((_this).F10())))).Cardinality())
		r3 = _2471_v101
		return r0, r1, r2, r3
	}
}
func (_this *C7) F11() _dafny.Set {
	{
		return _this._f11
	}
}

// End of class C7
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
