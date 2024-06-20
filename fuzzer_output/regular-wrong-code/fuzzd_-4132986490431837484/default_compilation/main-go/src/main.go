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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	return !(!(!(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})), _dafny.UnicodeSeqOfUtf8Bytes("uhvt")))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(false))).Cardinality()), _dafny.IntOfInt64(242)), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-242)))), (func() _dafny.Sequence {
		if true {
			return _dafny.SeqOf(_dafny.IntOfInt64(-647), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), _dafny.IntOfInt64(578), _dafny.IntOfInt64(227))
		}
		return _dafny.SeqOf((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-573))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})))).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_2_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(395), _dafny.IntOfInt64(-527))).Cardinality())
		}))).Cardinality()))))
	})())
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Set, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	if false {
		return (_dafny.SetOf(_dafny.SetOf(!(false)))).Intersection(_dafny.SetOf(_dafny.SetOf((Companion_D10_.Create_DC29_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfInt64(-791))).Cardinality()), false)).Dtor_cf57()), _dafny.SetOf(false), _dafny.SetOf(true, true), _dafny.SetOf(false), _dafny.SetOf(true)))
	} else {
		return _dafny.SetOf(_dafny.SetOf(false, false), _dafny.SetOf(false))
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if true {
			return _dafny.SetOf(true, true, true)
		}
		return _dafny.SetOf(false)
	})()).Difference((_dafny.SetOf(true, true, true, false)).Intersection(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(207), _dafny.IntOfInt64(-76)))), (_dafny.MultiSetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(678), _dafny.SeqOf(true))).Cardinality())).Cardinality(), _dafny.IntOfInt64(590)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tx")).Cardinality())), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-348)
	}))))).IsSubsetOf(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(732)))), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_4_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()
	}))).Cardinality()), false, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) D3 {
	if (!(false)) || (!(!(false))) {
		if false {
			return Companion_D3_.Create_DC8_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('u'), _dafny.CodePoint('h'))), !(true), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-603), false)).Cardinality()))
		} else {
			return Companion_D3_.Create_DC8_(_dafny.MultiSetOf(_dafny.CodePoint('o')), false, _dafny.IntOfInt64(251))
		}
	} else {
		return Companion_D3_.Create_DC8_(_dafny.MultiSetOf(_dafny.CodePoint('v')), !(false), (_dafny.MultiSetOf(false)).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 D0, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('u')
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D0 {
	var _source0 D4 = Companion_D4_.Create_DC10_(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-492))))
	_ = _source0
	if _source0.Is_DC11() {
		var _5___mcc_h0 _dafny.Int = _source0.Get_().(D4_DC11).Cf23
		_ = _5___mcc_h0
		var _6___mcc_h1 bool = _source0.Get_().(D4_DC11).Cf24
		_ = _6___mcc_h1
		var _7___mcc_h2 _dafny.Int = _source0.Get_().(D4_DC11).Cf25
		_ = _7___mcc_h2
		var _8_cf25 _dafny.Int = _7___mcc_h2
		_ = _8_cf25
		var _9_cf24 bool = _6___mcc_h1
		_ = _9_cf24
		var _10_cf23 _dafny.Int = _5___mcc_h0
		_ = _10_cf23
		return Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_9_cf24))
	} else {
		var _11___mcc_h3 _dafny.MultiSet = _source0.Get_().(D4_DC10).Cf22
		_ = _11___mcc_h3
		var _12_cf22 _dafny.MultiSet = _11___mcc_h3
		_ = _12_cf22
		return Companion_D0_.Create_DC1_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 bool, globalState *GlobalState) D0 {
	var _source1 D21 = Companion_D21_.Create_DC52_(_dafny.CodePoint('k'))
	_ = _source1
	if _source1.Is_DC52() {
		var _13___mcc_h0 _dafny.CodePoint = _source1.Get_().(D21_DC52).Cf89
		_ = _13___mcc_h0
		var _14_cf89 _dafny.CodePoint = _13___mcc_h0
		_ = _14_cf89
		return Companion_D0_.Create_DC0_(true)
	} else if _source1.Is_DC51() {
		var _15___mcc_h1 _dafny.Map = _source1.Get_().(D21_DC51).Cf88
		_ = _15___mcc_h1
		var _16_cf88 _dafny.Map = _15___mcc_h1
		_ = _16_cf88
		return Companion_D0_.Create_DC0_(false)
	} else {
		var _17___mcc_h2 D21 = _source1.Get_().(D21_DC53).Cf90
		_ = _17___mcc_h2
		var _18_cf90 D21 = _17___mcc_h2
		_ = _18_cf90
		if !(false) {
			return Companion_D0_.Create_DC0_(!(false))
		} else {
			return Companion_D0_.Create_DC0_(true)
		}
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.One).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-669))) != 0 {
		return _dafny.CodePoint('h')
	} else {
		return _dafny.CodePoint('x')
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D4_.Create_DC10_(_dafny.MultiSetOf(_dafny.IntOfInt64(656))), Companion_D4_.Create_DC10_(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-781))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_19_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()), _dafny.IntOfInt64(413))))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	var _source2 D20 = Companion_D20_.Create_DC49_(false, _dafny.CodePoint('w'))
	_ = _source2
	if _source2.Is_DC48() {
		var _20___mcc_h0 _dafny.Int = _source2.Get_().(D20_DC48).Cf82
		_ = _20___mcc_h0
		var _21___mcc_h1 _dafny.Array = _source2.Get_().(D20_DC48).Cf83
		_ = _21___mcc_h1
		var _22___mcc_h2 bool = _source2.Get_().(D20_DC48).Cf84
		_ = _22___mcc_h2
		var _23_cf84 bool = _22___mcc_h2
		_ = _23_cf84
		var _24_cf83 _dafny.Array = _21___mcc_h1
		_ = _24_cf83
		var _25_cf82 _dafny.Int = _20___mcc_h0
		_ = _25_cf82
		return (_dafny.IntOfInt64(470)).Minus(_dafny.IntOfInt64(613))
	} else if _source2.Is_DC49() {
		var _26___mcc_h3 bool = _source2.Get_().(D20_DC49).Cf85
		_ = _26___mcc_h3
		var _27___mcc_h4 _dafny.CodePoint = _source2.Get_().(D20_DC49).Cf86
		_ = _27___mcc_h4
		var _28_cf86 _dafny.CodePoint = _27___mcc_h4
		_ = _28_cf86
		var _29_cf85 bool = _26___mcc_h3
		_ = _29_cf85
		return Companion_Default___.SafeDivisionInt((_dafny.SetOf((_dafny.MultiSetOf(_29_cf85)).Cardinality(), _dafny.IntOfInt64(-440))).Cardinality(), _dafny.IntOfInt64(609))
	} else if _source2.Is_DC47() {
		var _30___mcc_h5 _dafny.Array = _source2.Get_().(D20_DC47).Cf81
		_ = _30___mcc_h5
		var _31_cf81 _dafny.Array = _30___mcc_h5
		_ = _31_cf81
		return _dafny.IntOfInt64(132)
	} else {
		var _32___mcc_h6 D20 = _source2.Get_().(D20_DC50).Cf87
		_ = _32___mcc_h6
		var _33_cf87 D20 = _32___mcc_h6
		_ = _33_cf87
		return _dafny.IntOfInt64(626)
	}
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(_dafny.CodePoint('u'), _dafny.CodePoint('u'))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-198))).Cardinality())).Cardinality(), _dafny.IntOfInt64(496))).Union((_dafny.MultiSetOf((_dafny.SetOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(550), _dafny.IntOfInt64(-304))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _34_v0 _dafny.Int
			_34_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(550)).Cmp(_34_v0) <= 0) && ((_34_v0).Cmp(_dafny.IntOfInt64(-304)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_34_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-980))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_35_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				}))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(165), _dafny.IntOfInt64(557))).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Cardinality())).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(297), (_dafny.Zero).Minus((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(68), true)).Cardinality())).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(true))).Merge((Companion_D26_.Create_DC62_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false, true, false, !(true))))).Dtor_cf100())
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), !(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), false))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _36_v0 _dafny.Map
			_36_v0 = interface{}(_compr_1).(_dafny.Map)
			if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), !(true)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), false))).Contains(_36_v0) {
				_coll1.Add(_36_v0, _dafny.MultiSetOf(_dafny.CodePoint('p')))
			}
		}
		return _coll1.ToMap()
	}()).Merge((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(484))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_37_i0 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(727), false)
		}))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _38_v1 _dafny.Map
			_38_v1 = interface{}(_compr_2).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(484))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_37_i0 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(727), false)
			})), _38_v1) {
				_coll2.Add(_38_v1, _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('o'))))
			}
		}
		return _coll2.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(107), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-709), (Companion_D10_.Create_DC29_(_dafny.IntOfInt64(433), true)).Dtor_cf57()))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _39_v2 _dafny.Map
			_39_v2 = interface{}(_compr_3).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(107), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-709), (Companion_D10_.Create_DC29_(_dafny.IntOfInt64(433), true)).Dtor_cf57())), _39_v2) {
				_coll3.Add(_39_v2, _dafny.MultiSetOf(_dafny.CodePoint('a'), _dafny.CodePoint('r')))
			}
		}
		return _coll3.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(323), Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(_dafny.IntOfInt64(826))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cvfvdpjmg")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true, false, true, true)).Difference(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source3 D8 = Companion_D8_.Create_DC22_((_dafny.Zero).Minus(_dafny.IntOfInt64(-494)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(772))).Cardinality()))
	_ = _source3
	if _source3.Is_DC22() {
		var _40___mcc_h0 _dafny.Int = _source3.Get_().(D8_DC22).Cf41
		_ = _40___mcc_h0
		var _41___mcc_h1 _dafny.Int = _source3.Get_().(D8_DC22).Cf42
		_ = _41___mcc_h1
		var _42_cf42 _dafny.Int = _41___mcc_h1
		_ = _42_cf42
		var _43_cf41 _dafny.Int = _40___mcc_h0
		_ = _43_cf41
		return _dafny.CodePoint('w')
	} else if _source3.Is_DC23() {
		var _44___mcc_h2 bool = _source3.Get_().(D8_DC23).Cf43
		_ = _44___mcc_h2
		var _45___mcc_h3 bool = _source3.Get_().(D8_DC23).Cf44
		_ = _45___mcc_h3
		var _46___mcc_h4 bool = _source3.Get_().(D8_DC23).Cf45
		_ = _46___mcc_h4
		var _47_cf45 bool = _46___mcc_h4
		_ = _47_cf45
		var _48_cf44 bool = _45___mcc_h3
		_ = _48_cf44
		var _49_cf43 bool = _44___mcc_h2
		_ = _49_cf43
		if true {
			return _dafny.CodePoint('n')
		} else {
			return _dafny.CodePoint('j')
		}
	} else if _source3.Is_DC24() {
		var _50___mcc_h5 bool = _source3.Get_().(D8_DC24).Cf46
		_ = _50___mcc_h5
		var _51___mcc_h6 _dafny.Int = _source3.Get_().(D8_DC24).Cf47
		_ = _51___mcc_h6
		var _52___mcc_h7 _dafny.Sequence = _source3.Get_().(D8_DC24).Cf48
		_ = _52___mcc_h7
		var _53___mcc_h8 bool = _source3.Get_().(D8_DC24).Cf49
		_ = _53___mcc_h8
		var _54___mcc_h9 _dafny.Sequence = _source3.Get_().(D8_DC24).Cf50
		_ = _54___mcc_h9
		var _55_cf50 _dafny.Sequence = _54___mcc_h9
		_ = _55_cf50
		var _56_cf49 bool = _53___mcc_h8
		_ = _56_cf49
		var _57_cf48 _dafny.Sequence = _52___mcc_h7
		_ = _57_cf48
		var _58_cf47 _dafny.Int = _51___mcc_h6
		_ = _58_cf47
		var _59_cf46 bool = _50___mcc_h5
		_ = _59_cf46
		return _dafny.CodePoint('r')
	} else if _source3.Is_DC21() {
		var _60___mcc_h10 _dafny.Map = _source3.Get_().(D8_DC21).Cf40
		_ = _60___mcc_h10
		var _61_cf40 _dafny.Map = _60___mcc_h10
		_ = _61_cf40
		if true {
			return _dafny.CodePoint('p')
		} else {
			return _dafny.CodePoint('a')
		}
	} else {
		var _62___mcc_h11 D8 = _source3.Get_().(D8_DC25).Cf51
		_ = _62___mcc_h11
		var _63_cf51 D8 = _62___mcc_h11
		_ = _63_cf51
		return _dafny.CodePoint('l')
	}
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, globalState *GlobalState) _dafny.Map {
	var _source4 D8 = Companion_D8_.Create_DC24_(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-792))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_64_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))).Cardinality()), _dafny.SeqOf(true), false, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), _dafny.IntOfInt64(905))).Cardinality()))
	_ = _source4
	if _source4.Is_DC22() {
		var _65___mcc_h0 _dafny.Int = _source4.Get_().(D8_DC22).Cf41
		_ = _65___mcc_h0
		var _66___mcc_h1 _dafny.Int = _source4.Get_().(D8_DC22).Cf42
		_ = _66___mcc_h1
		var _67_cf42 _dafny.Int = _66___mcc_h1
		_ = _67_cf42
		var _68_cf41 _dafny.Int = _65___mcc_h0
		_ = _68_cf41
		return ((Companion_D8_.Create_DC21_(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer10 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_69_i1 _dafny.Int) D0 {
				return Companion_D0_.Create_DC1_(_dafny.MultiSetOf(true))
			}))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _70_v0 D0
				_70_v0 = interface{}(_compr_4).(D0)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(682))).Uint32(), func(coer11 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_69_i1 _dafny.Int) D0 {
					return Companion_D0_.Create_DC1_(_dafny.MultiSetOf(true))
				})), _70_v0) {
					_coll4.Add(_70_v0, false)
				}
			}
			return _coll4.ToMap()
		}())).Dtor_cf40()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))), false))
	} else if _source4.Is_DC23() {
		var _71___mcc_h2 bool = _source4.Get_().(D8_DC23).Cf43
		_ = _71___mcc_h2
		var _72___mcc_h3 bool = _source4.Get_().(D8_DC23).Cf44
		_ = _72___mcc_h3
		var _73___mcc_h4 bool = _source4.Get_().(D8_DC23).Cf45
		_ = _73___mcc_h4
		var _74_cf45 bool = _73___mcc_h4
		_ = _74_cf45
		var _75_cf44 bool = _72___mcc_h3
		_ = _75_cf44
		var _76_cf43 bool = _71___mcc_h2
		_ = _76_cf43
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(false)), false)
	} else if _source4.Is_DC24() {
		var _77___mcc_h5 bool = _source4.Get_().(D8_DC24).Cf46
		_ = _77___mcc_h5
		var _78___mcc_h6 _dafny.Int = _source4.Get_().(D8_DC24).Cf47
		_ = _78___mcc_h6
		var _79___mcc_h7 _dafny.Sequence = _source4.Get_().(D8_DC24).Cf48
		_ = _79___mcc_h7
		var _80___mcc_h8 bool = _source4.Get_().(D8_DC24).Cf49
		_ = _80___mcc_h8
		var _81___mcc_h9 _dafny.Sequence = _source4.Get_().(D8_DC24).Cf50
		_ = _81___mcc_h9
		var _82_cf50 _dafny.Sequence = _81___mcc_h9
		_ = _82_cf50
		var _83_cf49 bool = _80___mcc_h8
		_ = _83_cf49
		var _84_cf48 _dafny.Sequence = _79___mcc_h7
		_ = _84_cf48
		var _85_cf47 _dafny.Int = _78___mcc_h6
		_ = _85_cf47
		var _86_cf46 bool = _77___mcc_h5
		_ = _86_cf46
		if _86_cf46 {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_83_cf49, _83_cf49)), _86_cf46)
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetFromSeq(_84_cf48)), _83_cf49)
		}
	} else if _source4.Is_DC21() {
		var _87___mcc_h10 _dafny.Map = _source4.Get_().(D8_DC21).Cf40
		_ = _87___mcc_h10
		var _88_cf40 _dafny.Map = _87___mcc_h10
		_ = _88_cf40
		return _88_cf40
	} else {
		var _89___mcc_h11 D8 = _source4.Get_().(D8_DC25).Cf51
		_ = _89___mcc_h11
		var _90_cf51 D8 = _89___mcc_h11
		_ = _90_cf51
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(true)), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(false)), true))
	}
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xsk")).Cardinality()))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _91_v0 _dafny.Int
			_91_v0 = interface{}(_compr_5).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xsk")).Cardinality())), _91_v0) {
				_coll5.Add(Companion_Default___.SafeDivisionInt(_91_v0, _dafny.IntOfInt64(-244)), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-729))))
			}
		}
		return _coll5.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(312), _dafny.IntOfInt64(669)))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC7_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-880)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-879)))), _dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('u'), _dafny.CodePoint('l')))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 D8, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(true)), false, true, false))).Union(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm41(p0 bool, p1 _dafny.CodePoint, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('d')
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(988), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(59))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_92_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(202)
	}))).Cardinality()))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.SetOf(true, !(false), true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_93_i1 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-690))
	}))))
}
func (_static *CompanionStruct_Default___) Fm43(globalState *GlobalState) _dafny.MultiSet {
	var _source5 D8 = Companion_D8_.Create_DC25_(Companion_D8_.Create_DC24_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gbqytxrv"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Cardinality())).Cardinality(), _dafny.SeqOf(true), !(true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_94_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-972)
	}))))
	_ = _source5
	if _source5.Is_DC22() {
		var _95___mcc_h0 _dafny.Int = _source5.Get_().(D8_DC22).Cf41
		_ = _95___mcc_h0
		var _96___mcc_h1 _dafny.Int = _source5.Get_().(D8_DC22).Cf42
		_ = _96___mcc_h1
		var _97_cf42 _dafny.Int = _96___mcc_h1
		_ = _97_cf42
		var _98_cf41 _dafny.Int = _95___mcc_h0
		_ = _98_cf41
		return _dafny.MultiSetOf(_dafny.IntOfInt64(-766), _dafny.IntOfInt64(342), _98_cf41, _97_cf42)
	} else if _source5.Is_DC23() {
		var _99___mcc_h2 bool = _source5.Get_().(D8_DC23).Cf43
		_ = _99___mcc_h2
		var _100___mcc_h3 bool = _source5.Get_().(D8_DC23).Cf44
		_ = _100___mcc_h3
		var _101___mcc_h4 bool = _source5.Get_().(D8_DC23).Cf45
		_ = _101___mcc_h4
		var _102_cf45 bool = _101___mcc_h4
		_ = _102_cf45
		var _103_cf44 bool = _100___mcc_h3
		_ = _103_cf44
		var _104_cf43 bool = _99___mcc_h2
		_ = _104_cf43
		return (_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-246)))))).Union(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(427))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_105_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality()))))
	} else if _source5.Is_DC24() {
		var _106___mcc_h5 bool = _source5.Get_().(D8_DC24).Cf46
		_ = _106___mcc_h5
		var _107___mcc_h6 _dafny.Int = _source5.Get_().(D8_DC24).Cf47
		_ = _107___mcc_h6
		var _108___mcc_h7 _dafny.Sequence = _source5.Get_().(D8_DC24).Cf48
		_ = _108___mcc_h7
		var _109___mcc_h8 bool = _source5.Get_().(D8_DC24).Cf49
		_ = _109___mcc_h8
		var _110___mcc_h9 _dafny.Sequence = _source5.Get_().(D8_DC24).Cf50
		_ = _110___mcc_h9
		var _111_cf50 _dafny.Sequence = _110___mcc_h9
		_ = _111_cf50
		var _112_cf49 bool = _109___mcc_h8
		_ = _112_cf49
		var _113_cf48 _dafny.Sequence = _108___mcc_h7
		_ = _113_cf48
		var _114_cf47 _dafny.Int = _107___mcc_h6
		_ = _114_cf47
		var _115_cf46 bool = _106___mcc_h5
		_ = _115_cf46
		return _dafny.MultiSetOf(_114_cf47, _114_cf47, _114_cf47, _114_cf47)
	} else if _source5.Is_DC21() {
		var _116___mcc_h10 _dafny.Map = _source5.Get_().(D8_DC21).Cf40
		_ = _116___mcc_h10
		var _117_cf40 _dafny.Map = _116___mcc_h10
		_ = _117_cf40
		return (_dafny.MultiSetOf(_dafny.IntOfInt64(832), (_dafny.SetOf(!(false))).Cardinality(), _dafny.IntOfInt64(789), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fuonpos")).Cardinality()), false)).Cardinality())).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(351))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_118_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality())))
	} else {
		var _119___mcc_h11 D8 = _source5.Get_().(D8_DC25).Cf51
		_ = _119___mcc_h11
		var _120_cf51 D8 = _119___mcc_h11
		_ = _120_cf51
		return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("schrvre")).Cardinality()), _dafny.IntOfInt64(887))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yrovpuj")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))))
	}
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality())) == 0)
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((Companion_D16_.Create_DC40_(_dafny.UnicodeSeqOfUtf8Bytes("llk"), _dafny.IntOfInt64(894))).Dtor_cf73(), _dafny.UnicodeSeqOfUtf8Bytes("hf")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(401))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_121_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})), _dafny.UnicodeSeqOfUtf8Bytes("lacfjew")))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, p1 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(668), false)).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-920))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_122_i0 _dafny.Int) _dafny.CodePoint {
		return (Companion_D21_.Create_DC52_(_dafny.CodePoint('j'))).Dtor_cf89()
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(!((Companion_D10_.Create_DC31_(false, !(!(true)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bbpvswxo")).Cardinality()), _dafny.IntOfInt64(340))).Cardinality()), (_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(481), (func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-128), _dafny.IntOfInt64(591))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _123_v0 _dafny.Int
			_123_v0 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(-128)).Cmp(_123_v0) <= 0) && ((_123_v0).Cmp(_dafny.IntOfInt64(591)) < 0) {
				_coll6.Add((_123_v0).Times(_dafny.IntOfInt64(-966)), true)
			}
		}
		return _coll6.ToMap()
	}()).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pnyamwon")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kd")).Cardinality()))).Cardinality()))).Cardinality())).Dtor_cf61()), false))).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(396))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_124_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))))
}
func (_static *CompanionStruct_Default___) Fm48(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(937))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_125_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(5)
	})), _dafny.SeqOf(_dafny.IntOfInt64(753)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(371))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_126_i1 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-478))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_127_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))).Cardinality()))
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(111)), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(916), _dafny.IntOfInt64(995))).Cardinality()))), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_128_i3 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(908)
	}))).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(989)), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(886))))).Cardinality()))), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(382)), _dafny.SeqOf(_dafny.IntOfInt64(-56), _dafny.IntOfInt64(776)), _dafny.SeqOf((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfInt64(36)))) {
		return _dafny.CodePoint('d')
	} else {
		return _dafny.CodePoint('j')
	}
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.MultiSet, p1 _dafny.CodePoint, p2 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(!(true)))).Cardinality()))
		}
		return _dafny.SeqOf(_dafny.IntOfInt64(220), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yxr")).Cardinality())), true)).Cardinality())
	})())).Union((_dafny.MultiSetOf(_dafny.IntOfInt64(18), _dafny.IntOfInt64(-111))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality()), _dafny.IntOfInt64(478), _dafny.IntOfInt64(708), _dafny.IntOfInt64(-3))))
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(112), _dafny.IntOfInt64(890))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _129_v0 _dafny.Int
			_129_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(112)).Cmp(_129_v0) <= 0) && ((_129_v0).Cmp(_dafny.IntOfInt64(890)) < 0) {
				_coll7.Add((_129_v0).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qxkxsgkqe")).Cardinality()), (_dafny.SetOf(false)).Cardinality())).Cardinality()), true)
			}
		}
		return _coll7.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jpyifw")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm51(p0 bool, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	if !(((_dafny.Zero).Minus(_dafny.IntOfInt64(-184))).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(284), false)).Cardinality()) != 0) {
		return _dafny.SeqOf(_dafny.SeqOf(true, true, true, !((Companion_D26_.Create_DC63_(_dafny.IntOfInt64(-626), _dafny.IntOfInt64(673), true)).Dtor_cf103())))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_130_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(true)
		})), _dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(!(false)), _dafny.SeqOf(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm52(p0 D8, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqOf(false))).Intersection(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(true, true))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _131_v0 _dafny.Sequence
			_131_v0 = interface{}(_compr_8).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(true, true)), _131_v0) {
				_coll8.Add(_131_v0)
			}
		}
		return _coll8.ToSet()
	}())).Difference((_dafny.SetOf(_dafny.SeqOf(false, false), _dafny.SeqOf(!(false)))).Difference(_dafny.SetOf(_dafny.SeqOf(true), _dafny.SeqOf(!(false), false), _dafny.SeqOf(true, false))))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) D10 {
	return Companion_D10_.Create_DC31_(!(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("cvutgcy"), _dafny.UnicodeSeqOfUtf8Bytes("g"))), !(false) || (true), ((_dafny.MultiSetOf(false)).Cardinality()).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(415))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_132_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(746))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_133_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vqw")).Cardinality())
	})), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pqg")).Cardinality())), _dafny.SeqOf((_dafny.SetOf(false)).Cardinality()))).Difference(func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(276))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_134_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(272), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(639))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_135_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()))).Cardinality())).Cardinality()))).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _136_v0 _dafny.Sequence
			_136_v0 = interface{}(_compr_9).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(276))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}(func(_134_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(272), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(639))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_135_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Cardinality()))).Cardinality())).Cardinality()))).Contains(_136_v0) {
				_coll9.Add(_136_v0)
			}
		}
		return _coll9.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm55(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	var _source6 D10 = Companion_D10_.Create_DC29_(_dafny.IntOfInt64(100), true)
	_ = _source6
	if _source6.Is_DC29() {
		var _137___mcc_h0 _dafny.Int = _source6.Get_().(D10_DC29).Cf56
		_ = _137___mcc_h0
		var _138___mcc_h1 bool = _source6.Get_().(D10_DC29).Cf57
		_ = _138___mcc_h1
		var _139_cf57 bool = _138___mcc_h1
		_ = _139_cf57
		var _140_cf56 _dafny.Int = _137___mcc_h0
		_ = _140_cf56
		return _dafny.SetOf(false)
	} else if _source6.Is_DC30() {
		var _141___mcc_h2 _dafny.CodePoint = _source6.Get_().(D10_DC30).Cf58
		_ = _141___mcc_h2
		var _142___mcc_h3 _dafny.MultiSet = _source6.Get_().(D10_DC30).Cf59
		_ = _142___mcc_h3
		var _143_cf59 _dafny.MultiSet = _142___mcc_h3
		_ = _143_cf59
		var _144_cf58 _dafny.CodePoint = _141___mcc_h2
		_ = _144_cf58
		return _dafny.SetOf(false)
	} else if _source6.Is_DC31() {
		var _145___mcc_h4 bool = _source6.Get_().(D10_DC31).Cf60
		_ = _145___mcc_h4
		var _146___mcc_h5 bool = _source6.Get_().(D10_DC31).Cf61
		_ = _146___mcc_h5
		var _147___mcc_h6 _dafny.Int = _source6.Get_().(D10_DC31).Cf62
		_ = _147___mcc_h6
		var _148_cf62 _dafny.Int = _147___mcc_h6
		_ = _148_cf62
		var _149_cf61 bool = _146___mcc_h5
		_ = _149_cf61
		var _150_cf60 bool = _145___mcc_h4
		_ = _150_cf60
		if !(_150_cf60) {
			return _dafny.SetOf(_149_cf61)
		} else {
			return _dafny.SetOf(_149_cf61)
		}
	} else if _source6.Is_DC28() {
		var _151___mcc_h7 *C1 = _source6.Get_().(D10_DC28).Cf55
		_ = _151___mcc_h7
		var _152_cf55 *C1 = _151___mcc_h7
		_ = _152_cf55
		return _dafny.SetOf((_152_cf55).F43(), (_152_cf55).F43(), (_152_cf55).F43(), !((_152_cf55).F43()), (_152_cf55).F43())
	} else {
		var _153___mcc_h8 D10 = _source6.Get_().(D10_DC32).Cf63
		_ = _153___mcc_h8
		var _154_cf63 D10 = _153___mcc_h8
		_ = _154_cf63
		return _dafny.SetOf(false, !(true))
	}
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(875), _dafny.IntOfInt64(-348)), (func() bool {
		if !(true) {
			return true
		}
		return !(!(!(!(true))))
	})(), _dafny.IntOfInt64(-364), !((_dafny.IntOfInt64(55)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(232), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality())).Cardinality()) > 0), _dafny.IntOfInt64(667))
}
func (_static *CompanionStruct_Default___) Fm57(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(833), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bcbb")).Cardinality()), false))
}
func (_static *CompanionStruct_Default___) Fm58(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC37_(Companion_Default___.SafeModuloInt((Companion_D26_.Create_DC63_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('q'))).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _155_v0 _dafny.CodePoint
			_155_v0 = interface{}(_compr_10).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('q')), _155_v0) {
				_coll10.Add(_155_v0)
			}
		}
		return _coll10.ToSet()
	}()).Cardinality())).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(70))).Cardinality(), true)).Dtor_cf101(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(323))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg31 _dafny.Int) interface{} {
			return coer31(arg31)
		}
	}(func(_156_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality())), true, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(723), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-464), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), (Companion_D8_.Create_DC22_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-737))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Dtor_cf41())).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm59(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}(func(_157_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
		}))).Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _158_v0 _dafny.Int
			_158_v0 = interface{}(_compr_11).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_157_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
			})), _158_v0) {
				_coll11.Add(Companion_Default___.SafeModuloInt(_158_v0, (_dafny.SetOf(false)).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(326)))
			}
		}
		return _coll11.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-574), _dafny.SetOf(_dafny.IntOfInt64(88))))).Merge(func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekgojohii")).Cardinality()), _dafny.IntOfInt64(374), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), _dafny.IntOfInt64(185), _dafny.IntOfInt64(958))).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(976))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_159_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(248), _dafny.IntOfInt64(565))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _160_v1 _dafny.Int
			_160_v1 = interface{}(_compr_12).(_dafny.Int)
			if (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekgojohii")).Cardinality()), _dafny.IntOfInt64(374), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), _dafny.IntOfInt64(185), _dafny.IntOfInt64(958))).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(976))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_159_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			}))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(248), _dafny.IntOfInt64(565))).Contains(_160_v1) {
				_coll12.Add((_160_v1).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vrr")).Cardinality())), func() _dafny.Set {
					var _coll13 = _dafny.NewBuilder()
					_ = _coll13
					for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(901), _dafny.IntOfInt64(58))); ; {
						_compr_13, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _161_v2 _dafny.Int
						_161_v2 = interface{}(_compr_13).(_dafny.Int)
						if ((_dafny.IntOfInt64(901)).Cmp(_161_v2) <= 0) && ((_161_v2).Cmp(_dafny.IntOfInt64(58)) < 0) {
							_coll13.Add((_161_v2).Plus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(789))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg36 _dafny.Int) interface{} {
									return coer36(arg36)
								}
							}(func(_162_i2 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('d')
							}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true))).Cardinality(), _dafny.IntOfInt64(-470))).Cardinality())).Cardinality()))
						}
					}
					return _coll13.ToSet()
				}())
			}
		}
		return _coll12.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm60(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(110))).Cardinality())).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vcr")).Cardinality()), (_dafny.SetOf((Companion_D22_.Create_DC55_((_dafny.SetOf(_dafny.CodePoint('p'), _dafny.CodePoint('q'), _dafny.CodePoint('l'), _dafny.CodePoint('l'))).Cardinality(), _dafny.IntOfInt64(372), true)).Dtor_cf93())).Cardinality(), _dafny.IntOfInt64(-613)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(342))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg37 _dafny.Int) interface{} {
			return coer37(arg37)
		}
	}(func(_163_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(543)
	}))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(746))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg38 _dafny.Int) interface{} {
			return coer38(arg38)
		}
	}(func(_164_i1 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-801))
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg39 _dafny.Int) interface{} {
			return coer39(arg39)
		}
	}(func(_165_i2 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vgo")).Cardinality())
	})))))
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) D8 {
	var _source7 D16 = Companion_D16_.Create_DC41_(Companion_D16_.Create_DC39_(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(522))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg40 _dafny.Int) interface{} {
			return coer40(arg40)
		}
	}(func(_166_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(707)
	})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), false)).Cardinality(), _dafny.IntOfInt64(395), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-512))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg41 _dafny.Int) interface{} {
			return coer41(arg41)
		}
	}(func(_167_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gkgia")).Cardinality()), _dafny.IntOfInt64(806)), _dafny.SeqOf(_dafny.IntOfInt64(124)))))
	_ = _source7
	if _source7.Is_DC40() {
		var _168___mcc_h0 _dafny.Sequence = _source7.Get_().(D16_DC40).Cf73
		_ = _168___mcc_h0
		var _169___mcc_h1 _dafny.Int = _source7.Get_().(D16_DC40).Cf74
		_ = _169___mcc_h1
		var _170_cf74 _dafny.Int = _169___mcc_h1
		_ = _170_cf74
		var _171_cf73 _dafny.Sequence = _168___mcc_h0
		_ = _171_cf73
		return Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(false)), false))
	} else if _source7.Is_DC39() {
		var _172___mcc_h2 _dafny.Set = _source7.Get_().(D16_DC39).Cf72
		_ = _172___mcc_h2
		var _173_cf72 _dafny.Set = _172___mcc_h2
		_ = _173_cf72
		return Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(!(false), !(false))), !(true)))
	} else {
		var _174___mcc_h3 D16 = _source7.Get_().(D16_DC41).Cf75
		_ = _174___mcc_h3
		var _175_cf75 D16 = _174___mcc_h3
		_ = _175_cf75
		return Companion_D8_.Create_DC21_(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(!(false))), _dafny.IntOfInt64(618))).Keys().Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _176_v0 D0
				_176_v0 = interface{}(_compr_14).(D0)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(!(false))), _dafny.IntOfInt64(618))).Contains(_176_v0) {
					_coll14.Add(_176_v0, true)
				}
			}
			return _coll14.ToMap()
		}())
	}
}
func (_static *CompanionStruct_Default___) Fm62(p0 _dafny.CodePoint, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D21 {
	return Companion_D21_.Create_DC51_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(285), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(902))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm63(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg42 _dafny.Int) interface{} {
			return coer42(arg42)
		}
	}(func(_177_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality())), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(385))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg43 _dafny.Int) interface{} {
			return coer43(arg43)
		}
	}(func(_178_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("eix")
	}))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm64(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(279), (Companion_D16_.Create_DC40_(_dafny.UnicodeSeqOfUtf8Bytes("stc"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Dtor_cf73())).Keys().Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _179_v0 _dafny.Int
			_179_v0 = interface{}(_compr_15).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(279), (Companion_D16_.Create_DC40_(_dafny.UnicodeSeqOfUtf8Bytes("stc"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Dtor_cf73())).Contains(_179_v0) {
				_coll15.Add((_179_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), Companion_D0_.Create_DC1_(_dafny.MultiSetOf(false, true)))
			}
		}
		return _coll15.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm65(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if !(true) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D8_.Create_DC22_(_dafny.IntOfInt64(-188), _dafny.Zero))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), Companion_D8_.Create_DC22_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality()), _dafny.IntOfInt64(398)))
	})()).Merge((Companion_D29_.Create_DC67_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D8_.Create_DC22_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), _dafny.IntOfInt64(792))))).Dtor_cf107())
}
func (_static *CompanionStruct_Default___) Fm66(globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC15_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fflctru")).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(137))))
}
func (_static *CompanionStruct_Default___) Fm67(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(true))), false))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)))
}
func (_static *CompanionStruct_Default___) Fm68(p0 bool, p1 bool, p2 bool, globalState *GlobalState) D20 {
	return Companion_D20_.Create_DC49_(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(!(false)), _dafny.SeqOf(false, false, true, false)), _dafny.CodePoint('q'))
}
func (_static *CompanionStruct_Default___) Fm69(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D16_.Create_DC39_(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(814)), _dafny.SeqOf(_dafny.IntOfInt64(920)))), _dafny.UnicodeSeqOfUtf8Bytes("wrtq"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D16_.Create_DC39_(_dafny.SetOf((Companion_D8_.Create_DC24_(false, (func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-784), _dafny.IntOfInt64(969))); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _180_v0 _dafny.Int
			_180_v0 = interface{}(_compr_16).(_dafny.Int)
			if ((_dafny.IntOfInt64(-784)).Cmp(_180_v0) <= 0) && ((_180_v0).Cmp(_dafny.IntOfInt64(969)) < 0) {
				_coll16.Add((_180_v0).Minus((_dafny.SetOf(_dafny.IntOfInt64(689), _dafny.IntOfInt64(252))).Cardinality()), false)
			}
		}
		return _coll16.ToMap()
	}()).Cardinality(), _dafny.SeqOf(false), !(!(true)), _dafny.SeqOf((_dafny.SetOf(false)).Cardinality()))).Dtor_cf50(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg44 _dafny.Int) interface{} {
			return coer44(arg44)
		}
	}(func(_181_i0 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("neweldmjf")).Cardinality()), _dafny.IntOfInt64(783))).Cardinality()
	}))).Cardinality())))), _dafny.UnicodeSeqOfUtf8Bytes("dbhwx")))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var _182_v0 _dafny.Map
	_ = _182_v0
	_182_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(555), Companion_Default___.SafeModuloInt(p2, p2))
	_182_v0 = _182_v0
	var _183_v1 _dafny.Map
	_ = _183_v1
	_183_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
	var _184_v2 _dafny.CodePoint
	_ = _184_v2
	_184_v2 = _dafny.CodePoint('y')
	var _185_v3 D20
	_ = _185_v3
	_185_v3 = Companion_D20_.Create_DC49_(p3, _184_v2)
	var _186_v4 _dafny.MultiSet
	_ = _186_v4
	_186_v4 = _dafny.MultiSetOf(_185_v3)
	var _187_v5 _dafny.MultiSet
	_ = _187_v5
	_187_v5 = _dafny.MultiSetOf(p1)
	var _188_v6 _dafny.Set
	_ = _188_v6
	_188_v6 = _dafny.SetOf((func() _dafny.Int {
		if (_187_v5).Contains(p3) {
			return (_187_v5).Multiplicity(p3)
		}
		return p2
	})())
	var _hi0 _dafny.Int = _dafny.IntOfInt64(312)
	_ = _hi0
	for _189_i0 := ((_dafny.Zero).Minus((_183_v1).Cardinality())).Times(Companion_Default___.Fm46((func() _dafny.Int {
		if (_186_v4).Contains(Companion_Default___.Fm68(p3, p1, p1, globalState)) {
			return (_186_v4).Multiplicity(Companion_Default___.Fm68(p3, p1, p1, globalState))
		}
		return p2
	})(), (_188_v6).Cardinality(), p1, globalState)); _189_i0.Cmp(_hi0) < 0; _189_i0 = _189_i0.Plus(_dafny.One) {
		var _190_v7 _dafny.Array
		_ = _190_v7
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) D2 = (func(_191_p2 _dafny.Int, _192_p3 bool, _193_p1 bool) func(_dafny.Int) D2 {
				return func(_194_i1 _dafny.Int) D2 {
					return Companion_D2_.Create_DC6_((_dafny.Zero).Minus(_191_p2), _192_p3, _dafny.IntOfInt64(597), _193_p1, _191_p2)
				}
			})(p2, p3, p1)
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
		_190_v7 = _nw0
		var _195_v8 _dafny.MultiSet
		_ = _195_v8
		_195_v8 = _dafny.MultiSetOf(_190_v7, _190_v7)
		var _196_v9 _dafny.Sequence
		_ = _196_v9
		_196_v9 = _dafny.UnicodeSeqOfUtf8Bytes("kuxvyifu")
		(globalState).F26 = (_dafny.MultiSetOf(_190_v7)).IsDisjointFrom((_195_v8).Update(_190_v7, Companion_Default___.Abs(_dafny.IntOfUint32((_196_v9).Cardinality()))))
		var _197_v10 _dafny.MultiSet
		_ = _197_v10
		_197_v10 = _dafny.MultiSetOf(_189_i0)
		var _198_v11 _dafny.Array
		_ = _198_v11
		var _nwElement0_0 _dafny.Int = _dafny.IntOfInt64(537)
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(16))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1(_189_i0, 1)
		(_nw1).ArraySet1((_dafny.SetOf(p2)).Cardinality(), 2)
		(_nw1).ArraySet1((p2).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_196_v9).Cardinality()))), 3)
		(_nw1).ArraySet1((func() _dafny.Int {
			if p3 {
				return _189_i0
			}
			return (_197_v10).Cardinality()
		})(), 4)
		(_nw1).ArraySet1(p2, 5)
		(_nw1).ArraySet1(p2, 6)
		(_nw1).ArraySet1(p2, 7)
		(_nw1).ArraySet1(Companion_Default___.Fm39(p1, _197_v10, p3, globalState), 8)
		(_nw1).ArraySet1(_dafny.IntOfUint32((_196_v9).Cardinality()), 9)
		(_nw1).ArraySet1(p2, 10)
		(_nw1).ArraySet1(_189_i0, 11)
		(_nw1).ArraySet1(p2, 12)
		(_nw1).ArraySet1((p2).Times(p2), 13)
		(_nw1).ArraySet1(((_182_v0).Merge(_182_v0)).Cardinality(), 14)
		(_nw1).ArraySet1(((p0).Select((Companion_Default___.SafeIndex(_189_i0, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)).Times(p2), 15)
		_198_v11 = _nw1
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))
		_ = _index0
		(_198_v11).ArraySet1(_189_i0, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))
		_ = _index1
		(_198_v11).ArraySet1(_189_i0, (_index1).Int())
		if p1 {
			(globalState).F13 = p2
			var _199_v12 _dafny.Map
			_ = _199_v12
			_199_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _184_v2)
			var _200_v13 _dafny.Map
			_ = _200_v13
			_200_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
			var _201_v14 _dafny.Map
			_ = _201_v14
			_201_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_199_v12).Cardinality(), _200_v13)
			var _202_v15 _dafny.Sequence
			_ = _202_v15
			_202_v15 = _dafny.SeqOf(p1)
			var _203_v16 _dafny.Sequence
			_ = _203_v16
			_203_v16 = _dafny.SeqOf(_202_v15)
			var _204_v17 *C9
			_ = _204_v17
			var _nw2 *C9 = New_C9_()
			_ = _nw2
			_nw2.Ctor__((_187_v5).Cardinality(), _201_v14, _184_v2, (_dafny.IntOfUint32((_196_v9).Cardinality())).Times((_200_v13).Cardinality()), p1, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3), _202_v15), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3), _202_v15)).Cardinality()))).Uint32(), p3), _dafny.Companion_Sequence_.Concatenate(_203_v16, _203_v16), _189_i0, ((_198_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))).Int()).(_dafny.Int)).Cmp(p2) > 0, p3, p3)
			_204_v17 = _nw2
			(globalState).F13 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_202_v15).Cardinality()), ((_204_v17).F38()).Times(_189_i0))
			(globalState).F26 = p3
			var _205_v18 _dafny.Array
			_ = _205_v18
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_1
			var _nw3 _dafny.Array
			_ = _nw3
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw3 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_206_p1 bool) func(_dafny.Int) bool {
					return func(_207_i2 _dafny.Int) bool {
						return _206_p1
					}
				})(p1)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw3 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw3).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw3).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_205_v18 = _nw3
			var _208_v19 _dafny.Map
			_ = _208_v19
			_208_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_196_v9, _205_v18)
			_208_v19 = (_208_v19).Update(_dafny.Companion_Sequence_.Concatenate(_196_v9, _dafny.Companion_Sequence_.Update(_196_v9, (Companion_Default___.SafeIndex(_189_i0, _dafny.IntOfUint32((_196_v9).Cardinality()))).Uint32(), _184_v2)), _205_v18)
		} else {
			var _209_v20 _dafny.Sequence
			_ = _209_v20
			_209_v20 = _dafny.SeqOf(p3, p3)
			var _210_v21 _dafny.Sequence
			_ = _210_v21
			_210_v21 = _dafny.SeqOf(_209_v20)
			var _211_v22 _dafny.Array
			_ = _211_v22
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
			_ = _nw4
			_211_v22 = _nw4
			var _212_v23 *C0
			_ = _212_v23
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(_211_v22)
			_212_v23 = _nw5
			var _213_v24 _dafny.Sequence
			_ = _213_v24
			_213_v24 = _dafny.SeqOf(_212_v23)
			var _214_v25 *C6
			_ = _214_v25
			var _nw6 *C6 = New_C6_()
			_ = _nw6
			_nw6.Ctor__(_184_v2, (func() _dafny.Int {
				if (_182_v0).Contains(_dafny.IntOfInt64(220)) {
					return (_182_v0).Get(_dafny.IntOfInt64(220)).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("f")).Cardinality())
			})(), p3, _dafny.SeqOf(p1), _210_v21, _189_i0, p1, !_dafny.Companion_Sequence_.Equal(_213_v24, _213_v24), p3)
			_214_v25 = _nw6
			var _215_v26 _dafny.Array
			_ = _215_v26
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
			_ = _nw7
			_215_v26 = _nw7
			var _216_v27 _dafny.Array
			_ = _216_v27
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw8
			_216_v27 = _nw8
			var _217_v28 _dafny.Map
			_ = _217_v28
			_217_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v27, _dafny.MultiSetOf(p3, (func() bool {
				if (_183_v1).Contains(_189_i0) {
					return (_183_v1).Get(_189_i0).(bool)
				}
				return false
			})()))
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_215_v26), 0))
			_ = _index2
			(_215_v26).ArraySet1(_217_v28, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_215_v26), 0))
			_ = _index3
			var _rhs0 _dafny.Map = _217_v28
			_ = _rhs0
			var _rhs1 _dafny.Int = _189_i0
			_ = _rhs1
			var _lhs0 _dafny.Array = _215_v26
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_215_v26), 0))
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			_lhs2.F13 = _rhs1
			_197_v10 = (_197_v10).Intersection((Companion_D4_.Create_DC10_(_197_v10)).Dtor_cf22())
			var _218_v29 _dafny.Set
			_ = _218_v29
			_218_v29 = _dafny.SetOf(_dafny.SeqOf(_189_i0), p0, _dafny.SeqOf((_198_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))).Int()).(_dafny.Int), p2), p0)
			var _219_v30 D16
			_ = _219_v30
			_219_v30 = Companion_D16_.Create_DC39_(_218_v29)
			var _220_v31 _dafny.Map
			_ = _220_v31
			_220_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v30, _196_v9)
			var _221_v32 D0
			_ = _221_v32
			_221_v32 = Companion_D0_.Create_DC0_(p3)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))
			_ = _index4
			var _rhs2 _dafny.Map = Companion_Default___.Fm69(p1, p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}((func(_222_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_223_i3 _dafny.Int) _dafny.CodePoint {
					return _222_v2
				}
			})(_184_v2)))).Cardinality()), !((_221_v32).Dtor_cf0()), globalState)
			_ = _rhs2
			var _rhs3 _dafny.Int = (_dafny.Zero).Minus(p2)
			_ = _rhs3
			var _lhs3 _dafny.Array = _198_v11
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))
			_ = _lhs4
			_220_v31 = _rhs2
			(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
			var _224_v33 T2
			_ = _224_v33
			var _nw9 *C8 = New_C8_()
			_ = _nw9
			_nw9.Ctor__(p3, _184_v2, _dafny.IntOfInt64(689), p1, _209_v20, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(418))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}((func(_225_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_226_i4 _dafny.Int) _dafny.Sequence {
					return _225_v20
				}
			})(_209_v20))), p2, p3, true, p3)
			_224_v33 = _nw9
			var _227_v34 _dafny.Sequence
			_ = _227_v34
			_227_v34 = _dafny.SeqOf(_224_v33)
			var _228_v35 _dafny.Sequence
			_ = _228_v35
			_228_v35 = _dafny.SeqOf(p0)
			var _229_v36 T3
			_ = _229_v36
			var _nw10 *C5 = New_C5_()
			_ = _nw10
			_nw10.Ctor__(p1, _dafny.IntOfUint32((_227_v34).Cardinality()), _dafny.IntOfUint32((_228_v35).Cardinality()), (_214_v25).Fm6(globalState), _209_v20, _210_v21, (_dafny.Zero).Minus(_dafny.IntOfUint32((_196_v9).Cardinality())), _224_v33.F28(), true, true)
			_229_v36 = _nw10
			var _230_v37 _dafny.Array
			_ = _230_v37
			var _nwElement0_1 T3 = _229_v36
			_ = _nwElement0_1
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(3))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_1, 0)
			(_nw11).ArraySet1(_229_v36, 1)
			(_nw11).ArraySet1(_229_v36, 2)
			_230_v37 = _nw11
			var _231_v38 _dafny.MultiSet
			_ = _231_v38
			_231_v38 = _dafny.MultiSetOf((func() _dafny.Array {
				if _224_v33.F28() {
					return _198_v11
				}
				return _198_v11
			})(), _198_v11, _198_v11, _198_v11, (func() _dafny.Array {
				if p3 {
					return _198_v11
				}
				return _198_v11
			})())
			var _rhs4 _dafny.Int = (_231_v38).Cardinality()
			_ = _rhs4
			var _rhs5 _dafny.Int = _189_i0
			_ = _rhs5
			var _rhs6 _dafny.Array = _230_v37
			_ = _rhs6
			var _rhs7 bool = (_229_v36).Fm6(globalState)
			_ = _rhs7
			var _rhs8 bool = !(((_198_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(656), _dafny.ArrayLen((_198_v11), 0))).Int()).(_dafny.Int)).Cmp(_189_i0) <= 0) || ((_209_v20).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_224_v33.F29())).Cardinality(), _dafny.IntOfUint32((_209_v20).Cardinality()))).Uint32()).(bool))
			_ = _rhs8
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			r0 = _rhs4
			r0 = _rhs5
			_230_v37 = _rhs6
			_lhs5.F26 = _rhs7
			_lhs6.F12 = _rhs8
		}
		var _232_v39 _dafny.Map
		_ = _232_v39
		_232_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
		var _233_v40 _dafny.Sequence
		_ = _233_v40
		_233_v40 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(p1, _dafny.IntOfUint32((_196_v9).Cardinality()), globalState), p2)).Update(p1, p2), (_232_v39).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _189_i0)), (_232_v39).Merge(_232_v39), _232_v39)
		_233_v40 = _233_v40
	}
	var _234_v41 _dafny.Set
	_ = _234_v41
	_234_v41 = _dafny.SetOf(_dafny.CodePoint('j'))
	var _235_v42 _dafny.Sequence
	_ = _235_v42
	_235_v42 = _dafny.SeqOf(p3, p1, true)
	var _236_v43 _dafny.MultiSet
	_ = _236_v43
	_236_v43 = _dafny.MultiSetOf(_234_v41)
	var _237_v44 _dafny.Sequence
	_ = _237_v44
	_237_v44 = _dafny.SeqOf(_235_v42)
	var _238_v45 *C10
	_ = _238_v45
	var _nw12 *C10 = New_C10_()
	_ = _nw12
	_nw12.Ctor__(_184_v2, p2, p3, _235_v42, _237_v44, p2, (_235_v42).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_235_v42).Cardinality()))).Uint32()).(bool), p1, p3)
	_238_v45 = _nw12
	var _239_v46 T3
	_ = _239_v46
	var _nw13 *C4 = New_C4_()
	_ = _nw13
	_nw13.Ctor__(p2, false, _dafny.SeqOf(!(p3), p3), _237_v44, p2, p1, p1, p3)
	_239_v46 = _nw13
	var _240_v47 _dafny.Map
	_ = _240_v47
	_240_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_238_v45, _239_v46)
	var _241_v48 _dafny.Sequence
	_ = _241_v48
	_241_v48 = _dafny.SeqOf(_240_v47, _240_v47, _240_v47, _240_v47, _240_v47)
	var _242_v49 _dafny.MultiSet
	_ = _242_v49
	_242_v49 = _dafny.MultiSetOf((_241_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_241_v48).Cardinality()))).Uint32()).(_dafny.Map), _240_v47)
	var _243_v50 _dafny.Array
	_ = _243_v50
	var _nwElement0_2 bool = p1
	_ = _nwElement0_2
	var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(23))
	_ = _nw14
	(_nw14).ArraySet1(_nwElement0_2, 0)
	(_nw14).ArraySet1((func() bool {
		if p1 {
			return true
		}
		return p3
	})(), 1)
	(_nw14).ArraySet1(!(p1) || (p3), 2)
	(_nw14).ArraySet1((_183_v1).Contains(Companion_Default___.Fm39(!(p1), _dafny.MultiSetOf(p2, p2, p2, (_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(574))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg47 _dafny.Int) interface{} {
			return coer47(arg47)
		}
	}((func(_244_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_245_i5 _dafny.Int) _dafny.CodePoint {
			return _244_v2
		}
	})(_184_v2)))).Cardinality())), p1, globalState)), 3)
	(_nw14).ArraySet1(p1, 4)
	(_nw14).ArraySet1(p3, 5)
	(_nw14).ArraySet1(p3, 6)
	(_nw14).ArraySet1(p3, 7)
	(_nw14).ArraySet1(p3, 8)
	(_nw14).ArraySet1((p3) && (p1), 9)
	(_nw14).ArraySet1((p3) == (p3), 10)
	(_nw14).ArraySet1(p3, 11)
	(_nw14).ArraySet1(p3, 12)
	(_nw14).ArraySet1(true, 13)
	(_nw14).ArraySet1(p1, 14)
	(_nw14).ArraySet1((_234_v41).Contains(_184_v2), 15)
	(_nw14).ArraySet1(p1, 16)
	(_nw14).ArraySet1((func() bool {
		if (_183_v1).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_235_v42).Cardinality()))) {
			return (_183_v1).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_235_v42).Cardinality()))).(bool)
		}
		return p1
	})(), 17)
	(_nw14).ArraySet1((p2).Cmp(p2) < 0, 18)
	(_nw14).ArraySet1(!((_236_v43).IsDisjointFrom(_236_v43)), 19)
	(_nw14).ArraySet1(p3, 20)
	(_nw14).ArraySet1(p1, 21)
	(_nw14).ArraySet1((_242_v49).Equals(_242_v49), 22)
	_243_v50 = _nw14
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_243_v50), 0))
	_ = _index5
	(_243_v50).ArraySet1(_239_v46.F29(), (_index5).Int())
	var _246_v51 _dafny.Sequence
	_ = _246_v51
	_246_v51 = _dafny.UnicodeSeqOfUtf8Bytes("ujnosv")
	var _247_v52 _dafny.MultiSet
	_ = _247_v52
	_247_v52 = _dafny.MultiSetOf(_246_v51)
	var _248_v53 _dafny.Sequence
	_ = _248_v53
	_248_v53 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg48 _dafny.Int) interface{} {
			return coer48(arg48)
		}
	}((func(_249_v51 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
		return func(_250_i6 _dafny.Int) _dafny.Sequence {
			return _249_v51
		}
	})(_246_v51))))
	var _251_v54 _dafny.MultiSet
	_ = _251_v54
	_251_v54 = _dafny.MultiSetOf(p2)
	var _252_v55 T2
	_ = _252_v55
	var _nw15 *C6 = New_C6_()
	_ = _nw15
	_nw15.Ctor__(_184_v2, (_239_v46).F30(), p3, _235_v42, (_239_v46).F33(), p2, _239_v46.F35(), _239_v46.F35(), p3)
	_252_v55 = _nw15
	var _253_v56 _dafny.Map
	_ = _253_v56
	_253_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_246_v51).Cardinality()), _252_v55)
	var _254_v57 _dafny.Sequence
	_ = _254_v57
	_254_v57 = _dafny.SeqOf((func() T2 {
		if (_253_v56).Contains(_dafny.IntOfInt64(330)) {
			return (_253_v56).Get(_dafny.IntOfInt64(330)).(T2)
		}
		return _252_v55
	})())
	var _255_v58 _dafny.Map
	_ = _255_v58
	_255_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v2, _dafny.IntOfUint32((_254_v57).Cardinality()))
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_243_v50), 0))
	_ = _index6
	var _rhs9 bool = (_239_v46).Fm6(globalState)
	_ = _rhs9
	var _rhs10 _dafny.MultiSet = ((_247_v52).Intersection(_247_v52)).Union(_dafny.MultiSetFromSeq((_248_v53).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
		if (_251_v54).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if (_255_v58).Contains(_184_v2) {
				return (_255_v58).Get(_184_v2).(_dafny.Int)
			}
			return (_239_v46).F30()
		})()))) {
			return (_251_v54).Multiplicity((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_255_v58).Contains(_184_v2) {
					return (_255_v58).Get(_184_v2).(_dafny.Int)
				}
				return (_239_v46).F30()
			})())))
		}
		return _dafny.IntOfUint32((_246_v51).Cardinality())
	})(), _dafny.IntOfUint32((_248_v53).Cardinality()))).Uint32()).(_dafny.Sequence)))
	_ = _rhs10
	var _lhs7 _dafny.Array = _243_v50
	_ = _lhs7
	var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_243_v50), 0))
	_ = _lhs8
	(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
	_247_v52 = _rhs10
	r0 = (_239_v46).F30()
	var _256_v59 _dafny.Array
	_ = _256_v59
	var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
	_ = _nw16
	_256_v59 = _nw16
	_256_v59 = _256_v59
	var _hi1 _dafny.Int = (_239_v46).F30()
	_ = _hi1
	for _257_i7 := (_252_v55).F30(); _257_i7.Cmp(_hi1) < 0; _257_i7 = _257_i7.Plus(_dafny.One) {
		(_239_v46).F31_set_(false)
		(globalState).F13 = ((_239_v46).F30()).Minus((_239_v46).F34())
		var _258_v60 D0
		_ = _258_v60
		_258_v60 = Companion_D0_.Create_DC1_(_187_v5)
		var _259_v61 *C10
		_ = _259_v61
		var _nw17 *C10 = New_C10_()
		_ = _nw17
		_nw17.Ctor__(Companion_Default___.Fm14(_dafny.IntOfInt64(386), ((_252_v55).F32()).Select((Companion_Default___.SafeIndex((_239_v46).F34(), _dafny.IntOfUint32(((_252_v55).F32()).Cardinality()))).Uint32()).(bool), (_252_v55).F30(), _258_v60, globalState), _dafny.IntOfInt64(-339), p1, _dafny.Companion_Sequence_.Concatenate((_239_v46).F32(), _235_v42), _dafny.Companion_Sequence_.Concatenate((_239_v46).F33(), _237_v44), (_239_v46).F30(), (_243_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_243_v50), 0))).Int()).(bool), _239_v46.F31(), (func() bool {
			if true {
				return (_238_v45).Fm7(_251_v54, (_dafny.Zero).Minus(_dafny.IntOfInt64(-158)), globalState)
			}
			return _252_v55.F29()
		})())
		_259_v61 = _nw17
		var _260_v62 T1
		_ = _260_v62
		var _nw18 *C2 = New_C2_()
		_ = _nw18
		_nw18.Ctor__((_239_v46).F34(), (_252_v55).F30(), _252_v55.F29(), _252_v55.F29(), false)
		_260_v62 = _nw18
		var _261_v63 _dafny.Sequence
		_ = _261_v63
		_261_v63 = _dafny.SeqOf(_260_v62, _260_v62)
		var _262_v64 _dafny.Sequence
		_ = _262_v64
		_262_v64 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_260_v62), _261_v63)
		_262_v64 = _262_v64
	}
	r0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p2, _dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_243_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(311), _dafny.ArrayLen((_243_v50), 0))).Int()).(bool) {
			return _246_v51
		}
		return _246_v51
	})()).Cardinality())))
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _263_v0 _dafny.Array
	_ = _263_v0
	var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
	_ = _nw19
	_263_v0 = _nw19
	var _264_v1 _dafny.Int
	_ = _264_v1
	_264_v1 = _dafny.IntOfInt64(20)
	var _265_v2 _dafny.Sequence
	_ = _265_v2
	_265_v2 = _dafny.SeqOf(_264_v1)
	var _266_v3 _dafny.Map
	_ = _266_v3
	_266_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_263_v0, _265_v2)
	var _267_v4 bool
	_ = _267_v4
	_267_v4 = true
	var _268_v5 _dafny.Map
	_ = _268_v5
	_268_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
		if (_266_v3).Contains(_263_v0) {
			return (_266_v3).Get(_263_v0).(_dafny.Sequence)
		}
		return _265_v2
	})(), _267_v4)
	var _269_v6 _dafny.Sequence
	_ = _269_v6
	_269_v6 = _dafny.UnicodeSeqOfUtf8Bytes("rhhbc")
	var _270_v7 _dafny.Sequence
	_ = _270_v7
	_270_v7 = _dafny.SeqOf(_269_v6)
	var _271_v8 _dafny.Array
	_ = _271_v8
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
	_ = _nw20
	_271_v8 = _nw20
	var _272_v9 _dafny.Array
	_ = _272_v9
	var _len0_2 _dafny.Int = _dafny.One
	_ = _len0_2
	var _nw21 _dafny.Array
	_ = _nw21
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw21 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) bool = (func(_273_v4 bool) func(_dafny.Int) bool {
			return func(_274_i0 _dafny.Int) bool {
				return _273_v4
			}
		})(_267_v4)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw21 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw21).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw21).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_272_v9 = _nw21
	var _275_globalState *GlobalState
	_ = _275_globalState
	var _nw22 *GlobalState = New_GlobalState_()
	_ = _nw22
	_nw22.Ctor__(_dafny.IntOfInt64(324), false, true, false, _dafny.IntOfInt64(862), (_268_v5).Update(_265_v2, _267_v4), _dafny.IntOfInt64(38), _269_v6, (_270_v7).Select((Companion_Default___.SafeIndex(_264_v1, _dafny.IntOfUint32((_270_v7).Cardinality()))).Uint32()).(_dafny.Sequence), true, _dafny.IntOfInt64(167), _dafny.IntOfInt64(-961), true, _dafny.IntOfInt64(457), true, _271_v8, _dafny.IntOfInt64(745), _dafny.CodePoint('d'), _272_v9, true, true, _272_v9, _263_v0, true, _dafny.CodePoint('y'), _265_v2, false, _dafny.IntOfInt64(426))
	_275_globalState = _nw22
	_264_v1 = _dafny.IntOfInt64(569)
	if _267_v4 {
		var _276_v10 _dafny.MultiSet
		_ = _276_v10
		_276_v10 = _dafny.MultiSetOf(_264_v1)
		var _277_v11 _dafny.Map
		_ = _277_v11
		_277_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, _264_v1)
		var _278_v12 _dafny.Sequence
		_ = _278_v12
		_278_v12 = _dafny.SeqOf(_267_v4)
		var _279_v13 _dafny.Map
		_ = _279_v13
		_279_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_264_v1), _dafny.IntOfUint32((_278_v12).Cardinality()))
		_264_v1 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_276_v10).Contains((func() _dafny.Int {
				if (_277_v11).Contains(_267_v4) {
					return (_277_v11).Get(_267_v4).(_dafny.Int)
				}
				return _264_v1
			})()) {
				return (_276_v10).Multiplicity((func() _dafny.Int {
					if (_277_v11).Contains(_267_v4) {
						return (_277_v11).Get(_267_v4).(_dafny.Int)
					}
					return _264_v1
				})())
			}
			return _264_v1
		})(), (func() _dafny.Int {
			if (_279_v13).Contains(_264_v1) {
				return (_279_v13).Get(_264_v1).(_dafny.Int)
			}
			return _264_v1
		})())
		var _280_v14 _dafny.Int
		_ = _280_v14
		var _out0 _dafny.Int
		_ = _out0
		_out0 = Companion_Default___.M0(_265_v2, _267_v4, _264_v1, !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, _263_v0)).Update(true, _263_v0)).Contains(_267_v4), _275_globalState)
		_280_v14 = _out0
		(_275_globalState).F1 = _dafny.Companion_Sequence_.Equal(_265_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(940))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}((func(_281_v14 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_282_i1 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_281_v14)
			}
		})(_280_v14))))
		var _283_v15 _dafny.Set
		_ = _283_v15
		_283_v15 = _dafny.SetOf(_267_v4, _267_v4)
		var _284_v16 _dafny.Map
		_ = _284_v16
		_284_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_272_v9, (_283_v15).Contains(_267_v4))
		_284_v16 = _284_v16
		var _285_v17 _dafny.MultiSet
		_ = _285_v17
		_285_v17 = _dafny.MultiSetOf(_267_v4, _267_v4)
		var _286_v18 _dafny.CodePoint
		_ = _286_v18
		_286_v18 = _dafny.CodePoint('b')
		var _287_v19 _dafny.Map
		_ = _287_v19
		_287_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, _286_v18)
		var _288_v20 _dafny.Array
		_ = _288_v20
		var _nwElement0_3 _dafny.Sequence = _278_v12
		_ = _nwElement0_3
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(11))
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_3, 0)
		(_nw23).ArraySet1(_278_v12, 1)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_278_v12, Companion_Default___.Fm0(_280_v14, _267_v4, _267_v4, (_285_v17).Cardinality(), _275_globalState)), 2)
		(_nw23).ArraySet1(_dafny.SeqOf(_267_v4), 3)
		(_nw23).ArraySet1((func() _dafny.Sequence {
			if _267_v4 {
				return _278_v12
			}
			return _278_v12
		})(), 4)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Update(_278_v12, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_287_v19).Cardinality()), _dafny.IntOfUint32((_278_v12).Cardinality()))).Uint32(), _267_v4), 5)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_267_v4, _267_v4), _278_v12), 6)
		(_nw23).ArraySet1(_278_v12, 7)
		(_nw23).ArraySet1(Companion_Default___.Fm0(_264_v1, _267_v4, !(_267_v4), (_285_v17).Cardinality(), _275_globalState), 8)
		(_nw23).ArraySet1(_278_v12, 9)
		(_nw23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_278_v12, _278_v12), 10)
		_288_v20 = _nw23
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_288_v20), 0))
		_ = _index7
		(_288_v20).ArraySet1(_278_v12, (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_288_v20), 0))
		_ = _index8
		(_288_v20).ArraySet1(_278_v12, (_index8).Int())
	} else {
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_272_v9), 0))
		_ = _index9
		(_272_v9).ArraySet1(_267_v4, (_index9).Int())
		var _289_v21 _dafny.MultiSet
		_ = _289_v21
		_289_v21 = _dafny.MultiSetOf(!(_267_v4))
		var _290_v22 _dafny.Map
		_ = _290_v22
		_290_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, !(_289_v21).Contains(_267_v4))
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_272_v9), 0))
		_ = _index10
		var _rhs11 bool = (func() bool {
			if (_290_v22).Contains(!(_267_v4) || (_267_v4)) {
				return (_290_v22).Get(!(_267_v4) || (_267_v4)).(bool)
			}
			return _267_v4
		})()
		_ = _rhs11
		var _rhs12 bool = !(Companion_Default___.Fm1((func() bool {
			if _267_v4 {
				return false
			}
			return _267_v4
		})(), _264_v1, _275_globalState))
		_ = _rhs12
		var _lhs9 _dafny.Array = _272_v9
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_272_v9), 0))
		_ = _lhs10
		var _lhs11 *GlobalState = _275_globalState
		_ = _lhs11
		(_lhs9).ArraySet1(_rhs11, (_lhs10).Int())
		_lhs11.F26 = _rhs12
		var _291_v23 _dafny.MultiSet
		_ = _291_v23
		_291_v23 = _dafny.MultiSetOf(_264_v1)
		var _292_v24 _dafny.Array
		_ = _292_v24
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_3
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_293_i2 _dafny.Int) bool {
				return false
			}
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
		_292_v24 = _nw24
		var _294_v25 _dafny.Set
		_ = _294_v25
		_294_v25 = _dafny.SetOf(_292_v24)
		var _295_v26 _dafny.Map
		_ = _295_v26
		_295_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool)), _292_v24)
		var _296_v27 _dafny.Map
		_ = _296_v27
		_296_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v23, (_294_v25).IsDisjointFrom(_dafny.SetOf(_292_v24, (func() _dafny.Array {
			if (_295_v26).Contains(_267_v4) {
				return (_295_v26).Get(_267_v4).(_dafny.Array)
			}
			return _292_v24
		})())))
		(_275_globalState).F1 = (func() bool {
			if (_296_v27).Contains(_291_v23) {
				return (_296_v27).Get(_291_v23).(bool)
			}
			return _267_v4
		})()
		var _297_v28 _dafny.Sequence
		_ = _297_v28
		_297_v28 = _dafny.SeqOf((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool))
		var _298_v29 _dafny.Int
		_ = _298_v29
		var _out1 _dafny.Int
		_ = _out1
		_out1 = Companion_Default___.M0(_265_v2, (_297_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.IntOfUint32((_297_v28).Cardinality()))).Uint32()).(bool), _264_v1, !(_267_v4), _275_globalState)
		_298_v29 = _out1
		_297_v28 = _297_v28
		var _299_v30 _dafny.Set
		_ = _299_v30
		_299_v30 = _dafny.SetOf(_dafny.IntOfUint32((_297_v28).Cardinality()))
		var _300_v31 _dafny.Set
		_ = _300_v31
		_300_v31 = _dafny.SetOf(_297_v28, _297_v28, _297_v28)
		var _301_v33 _dafny.MultiSet
		_ = _301_v33
		_301_v33 = _dafny.MultiSetOf(func() _dafny.Set {
			var _coll17 = _dafny.NewBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_300_v31).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _302_v32 _dafny.Sequence
				_302_v32 = interface{}(_compr_17).(_dafny.Sequence)
				if (_300_v31).Contains(_302_v32) {
					_coll17.Add(_302_v32)
				}
			}
			return _coll17.ToSet()
		}(), _300_v31)
		(_275_globalState).F26 = !(((_299_v30).IsDisjointFrom(_299_v30)) == ((_298_v29).Cmp((func() _dafny.Int {
			if (_301_v33).Contains(_300_v31) {
				return (_301_v33).Multiplicity(_300_v31)
			}
			return _298_v29
		})()) > 0))
	}
	var _hi2 _dafny.Int = _264_v1
	_ = _hi2
	for _303_i3 := _264_v1; _303_i3.Cmp(_hi2) < 0; _303_i3 = _303_i3.Plus(_dafny.One) {
		(_275_globalState).F1 = _267_v4
		var _304_v34 _dafny.MultiSet
		_ = _304_v34
		_304_v34 = _dafny.MultiSetOf(_303_i3)
		var _305_v35 *C7
		_ = _305_v35
		var _nw25 *C7 = New_C7_()
		_ = _nw25
		_nw25.Ctor__(_303_i3, _304_v34, (_264_v1).Times(_dafny.IntOfInt64(568)), !(_267_v4) || (_267_v4), _267_v4, _267_v4)
		_305_v35 = _nw25
		(_275_globalState).F12 = (_303_i3).Cmp((_305_v35).F41()) >= 0
		(_275_globalState).F13 = _303_i3
	}
	var _rhs13 _dafny.Int = _264_v1
	_ = _rhs13
	var _rhs14 _dafny.Int = _dafny.IntOfInt64(417)
	_ = _rhs14
	var _lhs12 *GlobalState = _275_globalState
	_ = _lhs12
	_lhs12.F13 = _rhs13
	_264_v1 = _rhs14
	var _306_v36 *C7
	_ = _306_v36
	var _nw26 *C7 = New_C7_()
	_ = _nw26
	_nw26.Ctor__(_264_v1, _dafny.MultiSetFromSeq(_265_v2), _dafny.IntOfInt64(780), _267_v4, false, _267_v4)
	_306_v36 = _nw26
	var _307_v37 _dafny.Map
	_ = _307_v37
	_307_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_272_v9, _306_v36)
	var _308_v38 _dafny.Map
	_ = _308_v38
	_308_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_306_v36).F41(), (_307_v37).Update(_272_v9, _306_v36))
	_307_v37 = (_307_v37).Merge((func() _dafny.Map {
		if (_308_v38).Contains(Companion_Default___.Fm46(_dafny.IntOfInt64(811), _264_v1, _267_v4, _275_globalState)) {
			return (_308_v38).Get(Companion_Default___.Fm46(_dafny.IntOfInt64(811), _264_v1, _267_v4, _275_globalState)).(_dafny.Map)
		}
		return _307_v37
	})())
	var _309_v39 _dafny.Map
	_ = _309_v39
	_309_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_265_v2).Select((Companion_Default___.SafeIndex(_264_v1, _dafny.IntOfUint32((_265_v2).Cardinality()))).Uint32()).(_dafny.Int), ((_306_v36).F41()).Minus(_264_v1))
	var _310_v40 _dafny.CodePoint
	_ = _310_v40
	_310_v40 = _dafny.CodePoint('m')
	var _311_v41 _dafny.Set
	_ = _311_v41
	_311_v41 = _dafny.SetOf(_310_v40)
	_309_v39 = (_309_v39).Update(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_264_v1), _264_v1), ((_311_v41).Union(_311_v41)).Cardinality())
	var _312_v42 _dafny.Array
	_ = _312_v42
	var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
	_ = _nw27
	_312_v42 = _nw27
	var _313_v43 _dafny.Array
	_ = _313_v43
	_313_v43 = _312_v42
	_312_v42 = (_313_v43)
	var _hi3 _dafny.Int = (_306_v36).F41()
	_ = _hi3
	for _314_i4 := _264_v1; _314_i4.Cmp(_hi3) < 0; _314_i4 = _314_i4.Plus(_dafny.One) {
		if _267_v4 {
			var _315_v44 D22
			_ = _315_v44
			_315_v44 = Companion_D22_.Create_DC55_(_dafny.IntOfInt64(263), _264_v1, _267_v4)
			var _316_v45 D3
			_ = _316_v45
			_316_v45 = Companion_D3_.Create_DC9_((_315_v44).Dtor_cf94(), _267_v4, _272_v9, _267_v4)
			var _317_v46 _dafny.Array
			_ = _317_v46
			var _318_v47 _dafny.Int
			_ = _318_v47
			var _out2 _dafny.Array
			_ = _out2
			var _out3 _dafny.Int
			_ = _out3
			_out2, _out3 = (_306_v36).M8((_316_v45).Dtor_cf20(), _275_globalState)
			_317_v46 = _out2
			_318_v47 = _out3
			_264_v1 = _314_i4
			(_275_globalState).F12 = _267_v4
			var _319_v48 _dafny.Array
			_ = _319_v48
			var _320_v49 _dafny.Int
			_ = _320_v49
			var _out4 _dafny.Array
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out4, _out5 = (_306_v36).M8(_272_v9, _275_globalState)
			_319_v48 = _out4
			_320_v49 = _out5
			(_275_globalState).F26 = _267_v4
		} else {
			(_275_globalState).F1 = (_dafny.IntOfInt64(934)).Cmp(_314_i4) <= 0
			var _321_v50 _dafny.Array
			_ = _321_v50
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(13))
			_ = _nw28
			_321_v50 = _nw28
			var _322_v51 _dafny.Set
			_ = _322_v51
			_322_v51 = _dafny.SetOf(_267_v4)
			var _323_v52 _dafny.Sequence
			_ = _323_v52
			_323_v52 = _dafny.SeqOf(_267_v4)
			var _324_v53 _dafny.Sequence
			_ = _324_v53
			_324_v53 = _dafny.SeqOf(_323_v52)
			var _325_v54 D14
			_ = _325_v54
			_325_v54 = Companion_D14_.Create_DC37_((_306_v36).F41(), _267_v4, (_306_v36).F41())
			var _326_v55 _dafny.MultiSet
			_ = _326_v55
			_326_v55 = _dafny.MultiSetOf(true)
			var _327_v56 _dafny.Map
			_ = _327_v56
			_327_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, !(false))
			var _nwElement0_4 _dafny.Set = _322_v51
			_ = _nwElement0_4
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
			_ = _nw29
			(_nw29).ArraySet1(_nwElement0_4, 0)
			(_nw29).ArraySet1(_322_v51, 1)
			(_nw29).ArraySet1(_322_v51, 2)
			(_nw29).ArraySet1((_322_v51).Difference(_dafny.SetOf(_267_v4, _267_v4)), 3)
			(_nw29).ArraySet1((_322_v51).Union(_322_v51), 4)
			(_nw29).ArraySet1((_322_v51).Difference(Companion_Default___.Fm11(_324_v53, !(_267_v4), _267_v4, (_306_v36).F41(), _275_globalState)), 5)
			(_nw29).ArraySet1(_322_v51, 6)
			(_nw29).ArraySet1(_dafny.SetOf(false), 7)
			(_nw29).ArraySet1(_322_v51, 8)
			(_nw29).ArraySet1(_322_v51, 9)
			(_nw29).ArraySet1((_322_v51).Difference(_322_v51), 10)
			(_nw29).ArraySet1(_dafny.SetOf(_267_v4, _267_v4, _267_v4), 11)
			(_nw29).ArraySet1((_322_v51).Difference(_322_v51), 12)
			(_nw29).ArraySet1(_dafny.SetOf(_267_v4, _267_v4, _267_v4), 13)
			(_nw29).ArraySet1(_322_v51, 14)
			(_nw29).ArraySet1(_322_v51, 15)
			(_nw29).ArraySet1(_322_v51, 16)
			(_nw29).ArraySet1((_322_v51), 17)
			(_nw29).ArraySet1(Companion_Default___.Fm45((_325_v54).Dtor_cf70(), _326_v55, (func() bool {
				if (_327_v56).Contains((_306_v36).Fm19(_275_globalState)) {
					return (_327_v56).Get((_306_v36).Fm19(_275_globalState)).(bool)
				}
				return true
			})(), _314_i4, _275_globalState), 18)
			(_nw29).ArraySet1(_322_v51, 19)
			(_nw29).ArraySet1(_322_v51, 20)
			(_nw29).ArraySet1((_322_v51).Intersection(_322_v51), 21)
			(_nw29).ArraySet1(_322_v51, 22)
			(_nw29).ArraySet1(_dafny.SetOf((_323_v52).Select((Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_323_v52).Cardinality()))).Uint32()).(bool), _267_v4), 23)
			(_nw29).ArraySet1((_322_v51).Intersection(_322_v51), 24)
			(_nw29).ArraySet1(_322_v51, 25)
			(_nw29).ArraySet1(_322_v51, 26)
			(_nw29).ArraySet1((Companion_Default___.Fm45(_dafny.IntOfUint32((_269_v6).Cardinality()), _326_v55, (_306_v36).Fm19(_275_globalState), _dafny.IntOfInt64(492), _275_globalState)).Union(_322_v51), 27)
			_321_v50 = _nw29
			var _328_v57 D1
			_ = _328_v57
			_328_v57 = Companion_D1_.Create_DC2_(_272_v9)
			var _pat_let_tv0 = _272_v9
			_ = _pat_let_tv0
			_328_v57 = func(_pat_let0_0 D1) D1 {
				return func(_329_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let1_0 _dafny.Array) D1 {
						return func(_330_dt__update_hcf2_h0 _dafny.Array) D1 {
							return Companion_D1_.Create_DC2_(_330_dt__update_hcf2_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(Companion_D1_.Create_DC2_(_272_v9))
			var _331_v58 _dafny.Array
			_ = _331_v58
			var _332_v59 _dafny.Int
			_ = _332_v59
			var _out6 _dafny.Array
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out6, _out7 = (_306_v36).M8(_272_v9, _275_globalState)
			_331_v58 = _out6
			_332_v59 = _out7
			var _333_v60 _dafny.Array
			_ = _333_v60
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_4
			var _nw30 _dafny.Array
			_ = _nw30
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw30 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.CodePoint = (func(_334_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_335_i5 _dafny.Int) _dafny.CodePoint {
						return _334_v40
					}
				})(_310_v40)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw30 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw30).ArraySet1CodePoint(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw30).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_333_v60 = _nw30
			var _336_v61 *C0
			_ = _336_v61
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(_333_v60)
			_336_v61 = _nw31
			_336_v61 = _336_v61
		}
		var _337_v62 *C2
		_ = _337_v62
		var _nw32 *C2 = New_C2_()
		_ = _nw32
		_nw32.Ctor__((_dafny.IntOfUint32((_dafny.SeqOf(_269_v6)).Cardinality())).Times(_dafny.IntOfUint32((_269_v6).Cardinality())), _264_v1, _267_v4, _267_v4, _267_v4)
		_337_v62 = _nw32
		var _338_v63 _dafny.Array
		_ = _338_v63
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_5
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_339_v4 bool) func(_dafny.Int) _dafny.MultiSet {
				return func(_340_i6 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_339_v4)
				}
			})(_267_v4)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw33 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw33).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw33).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_338_v63 = _nw33
		var _341_v64 _dafny.MultiSet
		_ = _341_v64
		_341_v64 = _dafny.MultiSetOf(_267_v4)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_338_v63), 0))
		_ = _index11
		(_338_v63).ArraySet1((_341_v64).Union(_341_v64), (_index11).Int())
		var _342_v65 _dafny.Sequence
		_ = _342_v65
		_342_v65 = _dafny.SeqOf(_267_v4, _267_v4)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_338_v63), 0))
		_ = _index12
		(_338_v63).ArraySet1(((func() _dafny.MultiSet {
			if (_342_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(292))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_343_v4 bool) func(_dafny.Int) _dafny.Map {
				return func(_344_i7 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v4, _343_v4)
				}
			})(_267_v4)))).Cardinality()), _dafny.IntOfUint32((_342_v65).Cardinality()))).Uint32()).(bool) {
				return _dafny.MultiSetOf(_267_v4, _267_v4)
			}
			return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_342_v65, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_269_v6).Cardinality()), _dafny.IntOfUint32((_342_v65).Cardinality()))).Uint32(), _267_v4), (Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_342_v65, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_269_v6).Cardinality()), _dafny.IntOfUint32((_342_v65).Cardinality()))).Uint32(), _267_v4)).Cardinality()))).Uint32(), _267_v4))
		})()).Intersection(_341_v64), (_index12).Int())
		(_275_globalState).F13 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_269_v6).Cardinality()))
	}
	var _345_v66 _dafny.Sequence
	_ = _345_v66
	_345_v66 = _dafny.SeqOf(_267_v4)
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
	_ = _index13
	(_263_v0).ArraySet1((_dafny.IntOfUint32((_345_v66).Cardinality())).Times((_306_v36).F41()), (_index13).Int())
	var _346_v67 D8
	_ = _346_v67
	_346_v67 = Companion_D8_.Create_DC24_(false, _dafny.IntOfUint32((_269_v6).Cardinality()), _345_v66, true, _265_v2)
	var _347_v68 _dafny.MultiSet
	_ = _347_v68
	_347_v68 = _dafny.MultiSetOf(_310_v40)
	var _348_v69 D3
	_ = _348_v69
	_348_v69 = Companion_D3_.Create_DC8_(_347_v68, _267_v4, (_306_v36).F41())
	var _349_v70 _dafny.Set
	_ = _349_v70
	_349_v70 = _dafny.SetOf(false, _267_v4, _267_v4, (_346_v67).Dtor_cf49(), (_348_v69).Dtor_cf16())
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
	_ = _index14
	(_263_v0).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_306_v36).F41(), (_349_v70).Cardinality()), _264_v1), (_index14).Int())
	(_275_globalState).F1 = (func() bool {
		if _267_v4 {
			return !(false)
		}
		return _267_v4
	})()
	var _350_i8 _dafny.Int
	_ = _350_i8
	_350_i8 = _dafny.Zero
	{
		for _267_v4 {
			{
				if (_350_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_350_i8 = (_350_i8).Plus(_dafny.One)
				var _351_v71 _dafny.CodePoint
				_ = _351_v71
				var _352_v72 _dafny.Sequence
				_ = _352_v72
				var _353_v73 _dafny.Int
				_ = _353_v73
				var _354_v74 _dafny.Sequence
				_ = _354_v74
				var _out8 _dafny.CodePoint
				_ = _out8
				var _out9 _dafny.Sequence
				_ = _out9
				var _out10 _dafny.Int
				_ = _out10
				var _out11 _dafny.Sequence
				_ = _out11
				_out8, _out9, _out10, _out11 = (_306_v36).M7(false, _272_v9, (_349_v70).IsSubsetOf(_dafny.SetOf((_306_v36).Fm19(_275_globalState))), _275_globalState)
				_351_v71 = _out8
				_352_v72 = _out9
				_353_v73 = _out10
				_354_v74 = _out11
				(_275_globalState).F1 = Companion_Default___.Fm1((_345_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-315), _dafny.IntOfUint32((_345_v66).Cardinality()))).Uint32()).(bool), (_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _275_globalState)
				_353_v73 = (_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int)
				var _355_v75 _dafny.Array
				_ = _355_v75
				var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
				_ = _nw34
				_355_v75 = _nw34
				var _356_v76 D20
				_ = _356_v76
				_356_v76 = Companion_D20_.Create_DC50_(Companion_D20_.Create_DC47_(_355_v75))
				var _357_v77 _dafny.Map
				_ = _357_v77
				_357_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, _356_v76)
				var _358_v78 D20
				_ = _358_v78
				_358_v78 = Companion_D20_.Create_DC50_((func() D20 {
					if (_357_v77).Contains(_267_v4) {
						return (_357_v77).Get(_267_v4).(D20)
					}
					return _356_v76
				})())
				var _source8 D20 = _358_v78
				_ = _source8
				if _source8.Is_DC48() {
					var _359___mcc_h0 _dafny.Int = _source8.Get_().(D20_DC48).Cf82
					_ = _359___mcc_h0
					var _360___mcc_h1 _dafny.Array = _source8.Get_().(D20_DC48).Cf83
					_ = _360___mcc_h1
					var _361___mcc_h2 bool = _source8.Get_().(D20_DC48).Cf84
					_ = _361___mcc_h2
					var _362_cf84 bool = _361___mcc_h2
					_ = _362_cf84
					var _363_cf83 _dafny.Array = _360___mcc_h1
					_ = _363_cf83
					var _364_cf82 _dafny.Int = _359___mcc_h0
					_ = _364_cf82
					var _365_v79 _dafny.Sequence
					_ = _365_v79
					_365_v79 = _dafny.SeqOf((Companion_D1_.Create_DC2_(_272_v9)).Dtor_cf2(), _272_v9, _272_v9, _272_v9)
					(_275_globalState).F21 = (_365_v79).Select((Companion_Default___.SafeIndex(_364_cf82, _dafny.IntOfUint32((_365_v79).Cardinality()))).Uint32()).(_dafny.Array)
					_270_v7 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_270_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.IntOfUint32((_270_v7).Cardinality()))).Uint32(), (_306_v36).Fm3(_267_v4, false, _275_globalState)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm39(_267_v4, (_306_v36).F42(), _267_v4, _275_globalState), _dafny.IntOfInt64(-975)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_270_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.IntOfUint32((_270_v7).Cardinality()))).Uint32(), (_306_v36).Fm3(_267_v4, false, _275_globalState))).Cardinality()))).Uint32(), _354_v74), (Companion_Default___.SafeIndex((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_270_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.IntOfUint32((_270_v7).Cardinality()))).Uint32(), (_306_v36).Fm3(_267_v4, false, _275_globalState)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm39(_267_v4, (_306_v36).F42(), _267_v4, _275_globalState), _dafny.IntOfInt64(-975)), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_270_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.IntOfUint32((_270_v7).Cardinality()))).Uint32(), (_306_v36).Fm3(_267_v4, false, _275_globalState))).Cardinality()))).Uint32(), _354_v74)).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("ydtxhh"))
					_362_cf84 = _362_cf84
					_309_v39 = (_309_v39).Update(((_306_v36).F41()).Plus((_306_v36).F41()), (_306_v36).F41())
				} else if _source8.Is_DC49() {
					var _366___mcc_h3 bool = _source8.Get_().(D20_DC49).Cf85
					_ = _366___mcc_h3
					var _367___mcc_h4 _dafny.CodePoint = _source8.Get_().(D20_DC49).Cf86
					_ = _367___mcc_h4
					var _368_cf86 _dafny.CodePoint = _367___mcc_h4
					_ = _368_cf86
					var _369_cf85 bool = _366___mcc_h3
					_ = _369_cf85
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_272_v9), 0))
					_ = _index15
					(_272_v9).ArraySet1((_306_v36).Fm19(_275_globalState), (_index15).Int())
					var _370_v81 _dafny.Map
					_ = _370_v81
					_370_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_cf86, _351_v71)
					var _371_v82 *C2
					_ = _371_v82
					var _nw35 *C2 = New_C2_()
					_ = _nw35
					_nw35.Ctor__((func() _dafny.Int {
						if ((_306_v36).F42()).Contains(Companion_Default___.Fm46((_dafny.Zero).Minus((func() _dafny.Map {
							var _coll18 = _dafny.NewMapBuilder()
							_ = _coll18
							for _iter18 := _dafny.Iterate((_370_v81).Keys().Elements()); ; {
								_compr_18, _ok18 := _iter18()
								if !_ok18 {
									break
								}
								var _372_v80 _dafny.CodePoint
								_372_v80 = interface{}(_compr_18).(_dafny.CodePoint)
								if (_370_v81).Contains(_372_v80) {
									_coll18.Add(_372_v80, _267_v4)
								}
							}
							return _coll18.ToMap()
						}()).Cardinality()), (_306_v36).F41(), _369_cf85, _275_globalState)) {
							return ((_306_v36).F42()).Multiplicity(Companion_Default___.Fm46((_dafny.Zero).Minus((func() _dafny.Map {
								var _coll19 = _dafny.NewMapBuilder()
								_ = _coll19
								for _iter19 := _dafny.Iterate((_370_v81).Keys().Elements()); ; {
									_compr_19, _ok19 := _iter19()
									if !_ok19 {
										break
									}
									var _373_v80 _dafny.CodePoint
									_373_v80 = interface{}(_compr_19).(_dafny.CodePoint)
									if (_370_v81).Contains(_373_v80) {
										_coll19.Add(_373_v80, _267_v4)
									}
								}
								return _coll19.ToMap()
							}()).Cardinality()), (_306_v36).F41(), _369_cf85, _275_globalState))
						}
						return (func() _dafny.Int {
							if ((_306_v36).F42()).Contains((_306_v36).F41()) {
								return ((_306_v36).F42()).Multiplicity((_306_v36).F41())
							}
							return (_306_v36).F41()
						})()
					})(), _dafny.IntOfUint32((_354_v74).Cardinality()), _267_v4, !((_345_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.IntOfUint32((_345_v66).Cardinality()))).Uint32()).(bool)), _267_v4)
					_371_v82 = _nw35
					var _374_v83 _dafny.Map
					_ = _374_v83
					_374_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_265_v2).Cardinality()), _371_v82)
					var _375_v84 _dafny.Array
					_ = _375_v84
					var _nwElement0_5 *C2 = _371_v82
					_ = _nwElement0_5
					var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(23))
					_ = _nw36
					(_nw36).ArraySet1(_nwElement0_5, 0)
					(_nw36).ArraySet1(_371_v82, 1)
					(_nw36).ArraySet1((func() *C2 {
						if (_374_v83).Contains((_306_v36).F41()) {
							return (_374_v83).Get((_306_v36).F41()).(*C2)
						}
						return _371_v82
					})(), 2)
					(_nw36).ArraySet1(_371_v82, 3)
					(_nw36).ArraySet1(_371_v82, 4)
					(_nw36).ArraySet1(_371_v82, 5)
					(_nw36).ArraySet1(_371_v82, 6)
					(_nw36).ArraySet1(_371_v82, 7)
					(_nw36).ArraySet1(_371_v82, 8)
					(_nw36).ArraySet1(_371_v82, 9)
					(_nw36).ArraySet1(_371_v82, 10)
					(_nw36).ArraySet1(_371_v82, 11)
					(_nw36).ArraySet1((func() *C2 {
						if (_374_v83).Contains(_dafny.IntOfInt64(22)) {
							return (_374_v83).Get(_dafny.IntOfInt64(22)).(*C2)
						}
						return _371_v82
					})(), 12)
					(_nw36).ArraySet1(_371_v82, 13)
					(_nw36).ArraySet1(_371_v82, 14)
					(_nw36).ArraySet1(_371_v82, 15)
					(_nw36).ArraySet1(_371_v82, 16)
					(_nw36).ArraySet1(_371_v82, 17)
					(_nw36).ArraySet1(_371_v82, 18)
					(_nw36).ArraySet1(_371_v82, 19)
					(_nw36).ArraySet1(_371_v82, 20)
					(_nw36).ArraySet1(_371_v82, 21)
					(_nw36).ArraySet1(_371_v82, 22)
					_375_v84 = _nw36
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_375_v84), 0))
					_ = _index16
					(_375_v84).ArraySet1(_371_v82, (_index16).Int())
					var _376_v85 D20
					_ = _376_v85
					_376_v85 = Companion_D20_.Create_DC49_(_267_v4, _310_v40)
					var _377_v86 D5
					_ = _377_v86
					_377_v86 = Companion_D5_.Create_DC13_(_369_cf85, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_369_cf85), (_371_v82).F46()), _310_v40)
					var _378_v87 _dafny.Array
					_ = _378_v87
					var _nwElement0_6 _dafny.CodePoint = _351_v71
					_ = _nwElement0_6
					var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(29))
					_ = _nw37
					(_nw37).ArraySet1CodePoint(_nwElement0_6, 0)
					(_nw37).ArraySet1CodePoint(_310_v40, 1)
					(_nw37).ArraySet1CodePoint(Companion_Default___.Fm41(_369_cf85, _368_cf86, _267_v4, (_306_v36).F41(), _275_globalState), 2)
					(_nw37).ArraySet1CodePoint(_310_v40, 3)
					(_nw37).ArraySet1CodePoint(_351_v71, 4)
					(_nw37).ArraySet1CodePoint(_351_v71, 5)
					(_nw37).ArraySet1CodePoint(_351_v71, 6)
					(_nw37).ArraySet1CodePoint(_dafny.CodePoint('t'), 7)
					(_nw37).ArraySet1CodePoint(_310_v40, 8)
					(_nw37).ArraySet1CodePoint(_351_v71, 9)
					(_nw37).ArraySet1CodePoint(_368_cf86, 10)
					(_nw37).ArraySet1CodePoint((_376_v85).Dtor_cf86(), 11)
					(_nw37).ArraySet1CodePoint(_351_v71, 12)
					(_nw37).ArraySet1CodePoint(_310_v40, 13)
					(_nw37).ArraySet1CodePoint(_dafny.CodePoint('x'), 14)
					(_nw37).ArraySet1CodePoint(_dafny.CodePoint('t'), 15)
					(_nw37).ArraySet1CodePoint((_377_v86).Dtor_cf29(), 16)
					(_nw37).ArraySet1CodePoint(_351_v71, 17)
					(_nw37).ArraySet1CodePoint(_351_v71, 18)
					(_nw37).ArraySet1CodePoint(_351_v71, 19)
					(_nw37).ArraySet1CodePoint(_dafny.CodePoint('a'), 20)
					(_nw37).ArraySet1CodePoint(_351_v71, 21)
					(_nw37).ArraySet1CodePoint(_310_v40, 22)
					(_nw37).ArraySet1CodePoint(_368_cf86, 23)
					(_nw37).ArraySet1CodePoint(_310_v40, 24)
					(_nw37).ArraySet1CodePoint(_310_v40, 25)
					(_nw37).ArraySet1CodePoint(_368_cf86, 26)
					(_nw37).ArraySet1CodePoint(_368_cf86, 27)
					(_nw37).ArraySet1CodePoint(Companion_Default___.Fm48(_369_cf85, _369_cf85, _264_v1, _275_globalState), 28)
					_378_v87 = _nw37
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_378_v87), 0))
					_ = _index17
					(_378_v87).ArraySet1CodePoint(_dafny.CodePoint('m'), (_index17).Int())
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_272_v9), 0))
					_ = _index18
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_375_v84), 0))
					_ = _index19
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_378_v87), 0))
					_ = _index20
					var _rhs15 bool = (_371_v82).Fm30(_369_cf85, _275_globalState)
					_ = _rhs15
					var _rhs16 *C2 = _371_v82
					_ = _rhs16
					var _rhs17 _dafny.Int = Companion_Default___.SafeModuloInt((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(908))
					_ = _rhs17
					var _rhs18 bool = !(_267_v4)
					_ = _rhs18
					var _rhs19 _dafny.CodePoint = _368_cf86
					_ = _rhs19
					var _lhs13 _dafny.Array = _272_v9
					_ = _lhs13
					var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_272_v9), 0))
					_ = _lhs14
					var _lhs15 _dafny.Array = _375_v84
					_ = _lhs15
					var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_375_v84), 0))
					_ = _lhs16
					var _lhs17 *GlobalState = _275_globalState
					_ = _lhs17
					var _lhs18 _dafny.Array = _378_v87
					_ = _lhs18
					var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_378_v87), 0))
					_ = _lhs19
					(_lhs13).ArraySet1(_rhs15, (_lhs14).Int())
					(_lhs15).ArraySet1(_rhs16, (_lhs16).Int())
					_353_v73 = _rhs17
					_lhs17.F1 = _rhs18
					(_lhs18).ArraySet1CodePoint(_rhs19, (_lhs19).Int())
					(_275_globalState).F26 = (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool)
					_306_v36 = _306_v36
					_368_cf86 = (_378_v87).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_378_v87), 0))).Int())
				} else if _source8.Is_DC47() {
					var _379___mcc_h5 _dafny.Array = _source8.Get_().(D20_DC47).Cf81
					_ = _379___mcc_h5
					var _380_cf81 _dafny.Array = _379___mcc_h5
					_ = _380_cf81
					var _381_v88 _dafny.Array
					_ = _381_v88
					var _382_v89 _dafny.Int
					_ = _382_v89
					var _out12 _dafny.Array
					_ = _out12
					var _out13 _dafny.Int
					_ = _out13
					_out12, _out13 = (_306_v36).M8(_272_v9, _275_globalState)
					_381_v88 = _out12
					_382_v89 = _out13
					var _383_v90 _dafny.Map
					_ = _383_v90
					_383_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, _267_v4)
					var _384_v91 _dafny.Map
					_ = _384_v91
					_384_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_306_v36).F41(), _383_v90)
					var _385_v92 _dafny.Set
					_ = _385_v92
					_385_v92 = _dafny.SetOf(_263_v0)
					var _386_v93 *C9
					_ = _386_v93
					var _nw38 *C9 = New_C9_()
					_ = _nw38
					_nw38.Ctor__(_382_v89, _384_v91, _310_v40, _382_v89, !(_267_v4), _dafny.Companion_Sequence_.Update(_345_v66, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.IntOfUint32((_345_v66).Cardinality()))).Uint32(), (_306_v36).Fm19(_275_globalState)), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_345_v66, _345_v66), (Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_dafny.SeqOf(_345_v66, _345_v66)).Cardinality()))).Uint32(), _345_v66), (_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _267_v4, (_385_v92).IsProperSubsetOf(_385_v92), _267_v4)
					_386_v93 = _nw38
					_386_v93 = _386_v93
					(_275_globalState).F1 = !(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg51 _dafny.Int) interface{} {
							return coer51(arg51)
						}
					}((func(_387_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_388_i9 _dafny.Int) _dafny.CodePoint {
							return _387_v40
						}
					})(_310_v40))), (Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(851))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg52 _dafny.Int) interface{} {
							return coer52(arg52)
						}
					}((func(_389_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_390_i9 _dafny.Int) _dafny.CodePoint {
							return _389_v40
						}
					})(_310_v40)))).Cardinality()))).Uint32(), _351_v71), _dafny.Companion_Sequence_.Concatenate(_354_v74, _354_v74))))
					var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
					_ = _index21
					(_263_v0).ArraySet1(((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(265))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}((func(_391_v71 _dafny.CodePoint, _392_v36 *C7, _393_v89 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_394_i10 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqOf(Companion_D21_.Create_DC53_(Companion_D21_.Create_DC53_(Companion_D21_.Create_DC52_(_391_v71))), Companion_D21_.Create_DC53_(Companion_D21_.Create_DC51_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_392_v36).F41(), _dafny.SetOf((_392_v36).F41(), _393_v89)))))).Cardinality())
						}
					})(_351_v71, _306_v36, _382_v89)))).Cardinality())).Minus(_dafny.IntOfUint32((_269_v6).Cardinality()))).Times((_306_v36).F41()), (_index21).Int())
				} else {
					var _395___mcc_h6 D20 = _source8.Get_().(D20_DC50).Cf87
					_ = _395___mcc_h6
					var _396_cf87 D20 = _395___mcc_h6
					_ = _396_cf87
					var _397_v94 _dafny.Array
					_ = _397_v94
					var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
					_ = _nw39
					_397_v94 = _nw39
					var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_397_v94), 0))
					_ = _index22
					(_397_v94).ArraySet1(_272_v9, (_index22).Int())
					var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_397_v94), 0))
					_ = _index23
					(_397_v94).ArraySet1(_272_v9, (_index23).Int())
					var _398_v95 _dafny.Set
					_ = _398_v95
					_398_v95 = _dafny.SetOf(_306_v36, _306_v36, _306_v36)
					var _399_v96 _dafny.Set
					_ = _399_v96
					_399_v96 = _dafny.SetOf(_398_v95)
					var _400_v97 _dafny.Sequence
					_ = _400_v97
					_400_v97 = _dafny.SeqOf(_399_v96)
					var _401_v98 D24
					_ = _401_v98
					_401_v98 = Companion_D24_.Create_DC57_(_dafny.SetOf(_398_v95, _398_v95))
					var _402_v99 _dafny.Array
					_ = _402_v99
					var _nwElement0_7 _dafny.Set = _399_v96
					_ = _nwElement0_7
					var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(15))
					_ = _nw40
					(_nw40).ArraySet1(_nwElement0_7, 0)
					(_nw40).ArraySet1(_dafny.SetOf(_398_v95, _398_v95), 1)
					(_nw40).ArraySet1(_399_v96, 2)
					(_nw40).ArraySet1((_399_v96).Union(_dafny.SetOf(_398_v95)), 3)
					(_nw40).ArraySet1(_399_v96, 4)
					(_nw40).ArraySet1(_399_v96, 5)
					(_nw40).ArraySet1((_399_v96).Difference((_400_v97).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_354_v74).Cardinality()), _dafny.IntOfUint32((_400_v97).Cardinality()))).Uint32()).(_dafny.Set)), 6)
					(_nw40).ArraySet1(_399_v96, 7)
					(_nw40).ArraySet1(_399_v96, 8)
					(_nw40).ArraySet1(_399_v96, 9)
					(_nw40).ArraySet1(_dafny.SetOf(_398_v95, _398_v95), 10)
					(_nw40).ArraySet1(_399_v96, 11)
					(_nw40).ArraySet1(_399_v96, 12)
					(_nw40).ArraySet1((_401_v98).Dtor_cf96(), 13)
					(_nw40).ArraySet1(_399_v96, 14)
					_402_v99 = _nw40
					var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_402_v99), 0))
					_ = _index24
					(_402_v99).ArraySet1(_399_v96, (_index24).Int())
					var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_402_v99), 0))
					_ = _index25
					(_402_v99).ArraySet1(_399_v96, (_index25).Int())
					var _403_v100 D3
					_ = _403_v100
					_403_v100 = Companion_D3_.Create_DC9_(_267_v4, _267_v4, _dafny.ArrayCastTo((_397_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_397_v94), 0))).Int())), _267_v4)
					var _pat_let_tv1 = _267_v4
					_ = _pat_let_tv1
					var _pat_let_tv2 = _267_v4
					_ = _pat_let_tv2
					var _404_v101 _dafny.Map
					_ = _404_v101
					_404_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let2_0 D3) D3 {
						return func(_405_dt__update__tmp_h1 D3) D3 {
							return func(_pat_let3_0 bool) D3 {
								return func(_406_dt__update_hcf21_h0 bool) D3 {
									return func(_pat_let4_0 bool) D3 {
										return func(_407_dt__update_hcf18_h0 bool) D3 {
											return Companion_D3_.Create_DC9_(_407_dt__update_hcf18_h0, (_405_dt__update__tmp_h1).Dtor_cf19(), (_405_dt__update__tmp_h1).Dtor_cf20(), _406_dt__update_hcf21_h0)
										}(_pat_let4_0)
									}(_pat_let_tv2)
								}(_pat_let3_0)
							}(_pat_let_tv1)
						}(_pat_let2_0)
					}(_403_v100), (_306_v36).F41())
					var _408_v102 _dafny.Map
					_ = _408_v102
					_408_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_306_v36).F41(), _310_v40)
					var _409_v103 D2
					_ = _409_v103
					_409_v103 = Companion_D2_.Create_DC6_((_306_v36).F41(), _267_v4, Companion_Default___.Fm46((_306_v36).F41(), (_404_v101).Cardinality(), true, _275_globalState), _267_v4, (_408_v102).Cardinality())
					_264_v1 = ((_409_v103).Dtor_cf13()).Minus((_dafny.IntOfUint32((_354_v74).Cardinality())).Times(_264_v1))
					var _410_v104 _dafny.CodePoint
					_ = _410_v104
					var _411_v105 _dafny.Sequence
					_ = _411_v105
					var _412_v106 _dafny.Int
					_ = _412_v106
					var _413_v107 _dafny.Sequence
					_ = _413_v107
					var _out14 _dafny.CodePoint
					_ = _out14
					var _out15 _dafny.Sequence
					_ = _out15
					var _out16 _dafny.Int
					_ = _out16
					var _out17 _dafny.Sequence
					_ = _out17
					_out14, _out15, _out16, _out17 = (_306_v36).M7(_dafny.Companion_Sequence_.Contains(_345_v66, _267_v4), _dafny.ArrayCastTo((_397_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_397_v94), 0))).Int())), Companion_Default___.Fm1(true, (Companion_Default___.Fm33(_267_v4, _275_globalState)).Cardinality(), _275_globalState), _275_globalState)
					_410_v104 = _out14
					_411_v105 = _out15
					_412_v106 = _out16
					_413_v107 = _out17
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _414_v108 _dafny.Sequence
	_ = _414_v108
	_414_v108 = _dafny.SeqOf(_345_v66)
	var _415_v109 *C4
	_ = _415_v109
	var _nw41 *C4 = New_C4_()
	_ = _nw41
	_nw41.Ctor__((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _267_v4, _345_v66, _414_v108, ((_306_v36).F41()).Plus((_349_v70).Cardinality()), true, _267_v4, _267_v4)
	_415_v109 = _nw41
	_270_v7 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_270_v7, (Companion_Default___.SafeIndex(_264_v1, _dafny.IntOfUint32((_270_v7).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg54 _dafny.Int) interface{} {
			return coer54(arg54)
		}
	}((func(_416_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_417_i11 _dafny.Int) _dafny.CodePoint {
			return _416_v40
		}
	})(_310_v40)))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ssisbjr")))
	var _418_v110 _dafny.MultiSet
	_ = _418_v110
	_418_v110 = _dafny.MultiSetOf(!(true))
	if (_418_v110).IsProperSubsetOf(_418_v110) {
		var _419_v111 _dafny.Array
		_ = _419_v111
		var _nw42 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
		_ = _nw42
		_419_v111 = _nw42
		var _420_v112 T0
		_ = _420_v112
		var _nw43 *C1 = New_C1_()
		_ = _nw43
		_nw43.Ctor__(false, _267_v4, _267_v4)
		_420_v112 = _nw43
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_419_v111), 0))
		_ = _index26
		(_419_v111).ArraySet1(_420_v112, (_index26).Int())
		var _421_v114 _dafny.Array
		_ = _421_v114
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_6
		var _nw44 _dafny.Array
		_ = _nw44
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw44 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Set = (func(_422_v1 _dafny.Int, _423_v36 *C7) func(_dafny.Int) _dafny.Set {
				return func(_424_i12 _dafny.Int) _dafny.Set {
					return (_dafny.SetOf(_dafny.IntOfInt64(646), (_dafny.Zero).Minus(_422_v1))).Union(func() _dafny.Set {
						var _coll20 = _dafny.NewBuilder()
						_ = _coll20
						for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(718), _dafny.IntOfInt64(593))); ; {
							_compr_20, _ok20 := _iter20()
							if !_ok20 {
								break
							}
							var _425_v113 _dafny.Int
							_425_v113 = interface{}(_compr_20).(_dafny.Int)
							if ((_dafny.IntOfInt64(718)).Cmp(_425_v113) <= 0) && ((_425_v113).Cmp(_dafny.IntOfInt64(593)) < 0) {
								_coll20.Add((_425_v113).Plus((_423_v36).F41()))
							}
						}
						return _coll20.ToSet()
					}())
				}
			})(_264_v1, _306_v36)
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
		_421_v114 = _nw44
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_419_v111), 0))
		_ = _index27
		var _rhs20 T0 = _420_v112
		_ = _rhs20
		var _rhs21 _dafny.Array = _421_v114
		_ = _rhs21
		var _rhs22 _dafny.MultiSet = _418_v110
		_ = _rhs22
		var _lhs20 _dafny.Array = _419_v111
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(55), _dafny.ArrayLen((_419_v111), 0))
		_ = _lhs21
		(_lhs20).ArraySet1(_rhs20, (_lhs21).Int())
		_421_v114 = _rhs21
		_418_v110 = _rhs22
		var _426_v115 _dafny.CodePoint
		_ = _426_v115
		var _427_v116 _dafny.Sequence
		_ = _427_v116
		var _428_v117 _dafny.Int
		_ = _428_v117
		var _429_v118 _dafny.Sequence
		_ = _429_v118
		var _out18 _dafny.CodePoint
		_ = _out18
		var _out19 _dafny.Sequence
		_ = _out19
		var _out20 _dafny.Int
		_ = _out20
		var _out21 _dafny.Sequence
		_ = _out21
		_out18, _out19, _out20, _out21 = (_306_v36).M7(false, _272_v9, _420_v112.F29(), _275_globalState)
		_426_v115 = _out18
		_427_v116 = _out19
		_428_v117 = _out20
		_429_v118 = _out21
		var _430_v119 _dafny.Map
		_ = _430_v119
		_430_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v4, (_306_v36).F41())
		var _431_v120 T3
		_ = _431_v120
		var _nw45 *C4 = New_C4_()
		_ = _nw45
		_nw45.Ctor__(_dafny.IntOfUint32((_429_v118).Cardinality()), _420_v112.F28(), _345_v66, _414_v108, (func() _dafny.Int {
			if (_430_v119).Contains(_267_v4) {
				return (_430_v119).Get(_267_v4).(_dafny.Int)
			}
			return (_306_v36).F41()
		})(), _267_v4, (_415_v109).Fm6(_275_globalState), _420_v112.F28())
		_431_v120 = _nw45
		var _432_v121 _dafny.MultiSet
		_ = _432_v121
		_432_v121 = _dafny.MultiSetOf(_431_v120)
		var _433_v122 D10
		_ = _433_v122
		_433_v122 = Companion_D10_.Create_DC30_(_310_v40, _432_v121)
		var _source9 D10 = (func() D10 {
			if _267_v4 {
				return _433_v122
			}
			return _433_v122
		})()
		_ = _source9
		if _source9.Is_DC29() {
			var _434___mcc_h7 _dafny.Int = _source9.Get_().(D10_DC29).Cf56
			_ = _434___mcc_h7
			var _435___mcc_h8 bool = _source9.Get_().(D10_DC29).Cf57
			_ = _435___mcc_h8
			var _436_cf57 bool = _435___mcc_h8
			_ = _436_cf57
			var _437_cf56 _dafny.Int = _434___mcc_h7
			_ = _437_cf56
			var _438_v123 _dafny.CodePoint
			_ = _438_v123
			var _439_v124 _dafny.Sequence
			_ = _439_v124
			var _440_v125 _dafny.Int
			_ = _440_v125
			var _441_v126 _dafny.Sequence
			_ = _441_v126
			var _out22 _dafny.CodePoint
			_ = _out22
			var _out23 _dafny.Sequence
			_ = _out23
			var _out24 _dafny.Int
			_ = _out24
			var _out25 _dafny.Sequence
			_ = _out25
			_out22, _out23, _out24, _out25 = (_306_v36).M7(_267_v4, _272_v9, _267_v4, _275_globalState)
			_438_v123 = _out22
			_439_v124 = _out23
			_440_v125 = _out24
			_441_v126 = _out25
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _index28
			(_263_v0).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_431_v120).F30())).Cardinality())), (_index28).Int())
			var _442_v127 *C7
			_ = _442_v127
			var _nw46 *C7 = New_C7_()
			_ = _nw46
			_nw46.Ctor__((_265_v2).Select((Companion_Default___.SafeIndex(_440_v125, _dafny.IntOfUint32((_265_v2).Cardinality()))).Uint32()).(_dafny.Int), ((_306_v36).F42()).Union((_306_v36).F42()), (_dafny.IntOfInt64(-421)).Minus(_dafny.IntOfUint32((_269_v6).Cardinality())), _431_v120.F35(), _436_cf57, _431_v120.F31())
			_442_v127 = _nw46
			var _rhs23 _dafny.Int = (_306_v36).F41()
			_ = _rhs23
			var _rhs24 _dafny.Sequence = (_431_v120).Fm3(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_431_v120.F28()), _345_v66), _420_v112.F28(), _275_globalState)
			_ = _rhs24
			var _rhs25 bool = !(_420_v112.F28()) || (!(_420_v112.F29()) || (_267_v4))
			_ = _rhs25
			var _rhs26 bool = ((_306_v36).F42()).IsProperSubsetOf((_442_v127).F42())
			_ = _rhs26
			var _lhs22 *GlobalState = _275_globalState
			_ = _lhs22
			var _lhs23 T3 = _431_v120
			_ = _lhs23
			_264_v1 = _rhs23
			_441_v126 = _rhs24
			_lhs22.F12 = _rhs25
			_lhs23.F35_set_(_rhs26)
		} else if _source9.Is_DC30() {
			var _443___mcc_h9 _dafny.CodePoint = _source9.Get_().(D10_DC30).Cf58
			_ = _443___mcc_h9
			var _444___mcc_h10 _dafny.MultiSet = _source9.Get_().(D10_DC30).Cf59
			_ = _444___mcc_h10
			var _445_cf59 _dafny.MultiSet = _444___mcc_h10
			_ = _445_cf59
			var _446_cf58 _dafny.CodePoint = _443___mcc_h9
			_ = _446_cf58
			var _rhs27 _dafny.Sequence = _429_v118
			_ = _rhs27
			var _rhs28 _dafny.Array = _272_v9
			_ = _rhs28
			var _lhs24 *GlobalState = _275_globalState
			_ = _lhs24
			_429_v118 = _rhs27
			_lhs24.F21 = _rhs28
			var _447_v128 _dafny.Set
			_ = _447_v128
			_447_v128 = _dafny.SetOf((_306_v36).F41())
			var _rhs29 bool = _431_v120.F31()
			_ = _rhs29
			var _rhs30 bool = (_447_v128).IsSubsetOf(_447_v128)
			_ = _rhs30
			var _lhs25 *GlobalState = _275_globalState
			_ = _lhs25
			var _lhs26 *GlobalState = _275_globalState
			_ = _lhs26
			_lhs25.F26 = _rhs29
			_lhs26.F26 = _rhs30
			var _448_v129 _dafny.Sequence
			_ = _448_v129
			_448_v129 = _dafny.SeqOf(_265_v2)
			_430_v119 = (_430_v119).Update(_431_v120.F35(), _dafny.IntOfUint32((_448_v129).Cardinality()))
			var _449_v130 _dafny.Sequence
			_ = _449_v130
			_449_v130 = _dafny.SeqOf(Companion_Default___.Fm66(_275_globalState))
			var _450_v131 T2
			_ = _450_v131
			var _nw47 *C5 = New_C5_()
			_ = _nw47
			_nw47.Ctor__(_431_v120.F31(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(535))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}((func(_451_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_452_i13 _dafny.Int) _dafny.CodePoint {
					return _451_v40
				}
			})(_310_v40)))).Cardinality()), (_dafny.Zero).Minus((_306_v36).F41()), _431_v120.F35(), (_431_v120).F32(), _dafny.SeqOf((_431_v120).F32(), (_431_v120).F32(), (_431_v120).F32(), _dafny.SeqOf(_431_v120.F29(), _431_v120.F29(), false, _267_v4, _431_v120.F31()), (_431_v120).F32()), _dafny.IntOfUint32((_449_v130).Cardinality()), true, _420_v112.F29(), _431_v120.F35())
			_450_v131 = _nw47
			var _453_v132 _dafny.MultiSet
			_ = _453_v132
			_453_v132 = _dafny.MultiSetOf(_450_v131)
			_453_v132 = (_dafny.MultiSetOf(_450_v131)).Difference(_453_v132)
		} else if _source9.Is_DC31() {
			var _454___mcc_h11 bool = _source9.Get_().(D10_DC31).Cf60
			_ = _454___mcc_h11
			var _455___mcc_h12 bool = _source9.Get_().(D10_DC31).Cf61
			_ = _455___mcc_h12
			var _456___mcc_h13 _dafny.Int = _source9.Get_().(D10_DC31).Cf62
			_ = _456___mcc_h13
			var _457_cf62 _dafny.Int = _456___mcc_h13
			_ = _457_cf62
			var _458_cf61 bool = _455___mcc_h12
			_ = _458_cf61
			var _459_cf60 bool = _454___mcc_h11
			_ = _459_cf60
			var _460_v133 *C2
			_ = _460_v133
			var _nw48 *C2 = New_C2_()
			_ = _nw48
			_nw48.Ctor__(_264_v1, _dafny.IntOfUint32(((_431_v120).F32()).Cardinality()), _431_v120.F28(), _431_v120.F29(), _431_v120.F29())
			_460_v133 = _nw48
			var _461_v134 D9
			_ = _461_v134
			_461_v134 = Companion_D9_.Create_DC26_(_460_v133)
			var _462_v135 _dafny.Array
			_ = _462_v135
			var _nw49 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
			_ = _nw49
			_462_v135 = _nw49
			var _463_v136 _dafny.Array
			_ = _463_v136
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw50
			_463_v136 = _nw50
			var _464_v137 *C3
			_ = _464_v137
			var _nw51 *C3 = New_C3_()
			_ = _nw51
			_nw51.Ctor__(_dafny.IntOfInt64(-956), _463_v136, _310_v40, _dafny.IntOfInt64(-341), _431_v120.F31(), _345_v66, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_465_v66 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_466_i14 _dafny.Int) _dafny.Sequence {
					return _465_v66
				}
			})(_345_v66))), _dafny.IntOfInt64(775), _431_v120.F28(), _459_cf60, _459_cf60)
			_464_v137 = _nw51
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_462_v135), 0))
			_ = _index29
			(_462_v135).ArraySet1(_464_v137, (_index29).Int())
			var _467_v138 _dafny.Array
			_ = _467_v138
			var _nwElement0_8 _dafny.Sequence = _429_v118
			_ = _nwElement0_8
			var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.One)
			_ = _nw52
			(_nw52).ArraySet1(_nwElement0_8, 0)
			_467_v138 = _nw52
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_467_v138), 0))
			_ = _index30
			(_467_v138).ArraySet1(_429_v118, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_462_v135), 0))
			_ = _index31
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_467_v138), 0))
			_ = _index32
			var _rhs31 _dafny.Int = (_265_v2).Select((Companion_Default___.SafeIndex((_431_v120).F34(), _dafny.IntOfUint32((_265_v2).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs31
			var _rhs32 D9 = Companion_D9_.Create_DC26_(_460_v133)
			_ = _rhs32
			var _rhs33 *C3 = _464_v137
			_ = _rhs33
			var _rhs34 _dafny.Sequence = _429_v118
			_ = _rhs34
			var _rhs35 _dafny.Int = (_431_v120).F34()
			_ = _rhs35
			var _lhs27 *GlobalState = _275_globalState
			_ = _lhs27
			var _lhs28 _dafny.Array = _462_v135
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_462_v135), 0))
			_ = _lhs29
			var _lhs30 _dafny.Array = _467_v138
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_467_v138), 0))
			_ = _lhs31
			_lhs27.F13 = _rhs31
			_461_v134 = _rhs32
			(_lhs28).ArraySet1(_rhs33, (_lhs29).Int())
			(_lhs30).ArraySet1(_rhs34, (_lhs31).Int())
			_264_v1 = _rhs35
			var _468_v139 bool
			_ = _468_v139
			var _469_v140 _dafny.Int
			_ = _469_v140
			var _470_v141 _dafny.Sequence
			_ = _470_v141
			var _out26 bool
			_ = _out26
			var _out27 _dafny.Int
			_ = _out27
			var _out28 _dafny.Sequence
			_ = _out28
			_out26, _out27, _out28 = (_431_v120).M1(_275_globalState)
			_468_v139 = _out26
			_469_v140 = _out27
			_470_v141 = _out28
			_469_v140 = _264_v1
			var _471_v142 D8
			_ = _471_v142
			_471_v142 = Companion_D8_.Create_DC24_(_420_v112.F28(), _469_v140, _dafny.SeqOf(_459_cf60), _267_v4, _265_v2)
			var _472_v143 D8
			_ = _472_v143
			_472_v143 = Companion_D8_.Create_DC25_(_471_v142)
			var _473_v144 D8
			_ = _473_v144
			_473_v144 = Companion_D8_.Create_DC25_(_472_v143)
			var _474_v145 _dafny.Map
			_ = _474_v145
			_474_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _426_v115)
			var _rhs36 D8 = _473_v144
			_ = _rhs36
			var _rhs37 bool = (_474_v145).Contains(_431_v120.F28())
			_ = _rhs37
			var _lhs32 T3 = _431_v120
			_ = _lhs32
			_473_v144 = _rhs36
			_lhs32.F29_set_(_rhs37)
		} else if _source9.Is_DC28() {
			var _475___mcc_h14 *C1 = _source9.Get_().(D10_DC28).Cf55
			_ = _475___mcc_h14
			var _476_cf55 *C1 = _475___mcc_h14
			_ = _476_cf55
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_272_v9), 0))
			_ = _index33
			(_272_v9).ArraySet1((_428_v117).Cmp((_418_v110).Cardinality()) >= 0, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_272_v9), 0))
			_ = _index34
			(_272_v9).ArraySet1(!_dafny.Companion_Sequence_.Equal((_431_v120).F32(), (_431_v120).F32()), (_index34).Int())
			var _477_v146 _dafny.Array
			_ = _477_v146
			var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(4))
			_ = _nw53
			_477_v146 = _nw53
			var _478_v147 *C0
			_ = _478_v147
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__(_477_v146)
			_478_v147 = _nw54
			var _479_v148 _dafny.Map
			_ = _479_v148
			_479_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_306_v36).F41()).Plus((_306_v36).F41()), (_415_v109).Fm6(_275_globalState))
			var _480_v149 _dafny.Sequence
			_ = _480_v149
			_480_v149 = _dafny.SeqOf((_306_v36).F42())
			var _481_v150 D5
			_ = _481_v150
			_481_v150 = Companion_D5_.Create_DC13_(((_306_v36).F42()).IsProperSubsetOf((_480_v149).Select((Companion_Default___.SafeIndex((_431_v120).F34(), _dafny.IntOfUint32((_480_v149).Cardinality()))).Uint32()).(_dafny.MultiSet)), _430_v119, _dafny.CodePoint('e'))
			var _482_v151 _dafny.Map
			_ = _482_v151
			_482_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_431_v120).F30(), (_431_v120).F30()), _478_v147)
			var _rhs38 *C0 = (func() *C0 {
				if (_482_v151).Contains((func() _dafny.Int {
					if _431_v120.F31() {
						return (_306_v36).F41()
					}
					return (_431_v120).F30()
				})()) {
					return (_482_v151).Get((func() _dafny.Int {
						if _431_v120.F31() {
							return (_306_v36).F41()
						}
						return (_431_v120).F30()
					})()).(*C0)
				}
				return _478_v147
			})()
			_ = _rhs38
			var _rhs39 _dafny.Map = (_479_v148).Update(_428_v117, (_420_v112.F28()) && (_431_v120.F29()))
			_ = _rhs39
			var _rhs40 bool = !((_dafny.SetOf(_431_v120.F35(), (_476_cf55).F43(), _431_v120.F29(), _420_v112.F28())).IsProperSubsetOf((_349_v70).Difference(Companion_Default___.Fm45((_306_v36).F41(), _418_v110, _420_v112.F29(), _264_v1, _275_globalState))))
			_ = _rhs40
			var _rhs41 _dafny.Int = (func() _dafny.Int {
				if ((_306_v36).Fm20((_431_v120).F30(), (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), _275_globalState)).Cmp((_306_v36).F41()) == 0 {
					return ((_306_v36).F41()).Minus(_dafny.IntOfUint32((_429_v118).Cardinality()))
				}
				return _264_v1
			})()
			_ = _rhs41
			var _rhs42 D5 = _481_v150
			_ = _rhs42
			var _lhs33 *GlobalState = _275_globalState
			_ = _lhs33
			var _lhs34 *GlobalState = _275_globalState
			_ = _lhs34
			_478_v147 = _rhs38
			_479_v148 = _rhs39
			_lhs33.F26 = _rhs40
			_lhs34.F13 = _rhs41
			_481_v150 = _rhs42
			var _483_v152 _dafny.Array
			_ = _483_v152
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw55
			_483_v152 = _nw55
			var _484_v153 _dafny.Map
			_ = _484_v153
			_484_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_483_v152, (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool))
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_272_v9), 0))
			_ = _index35
			(_272_v9).ArraySet1(((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_306_v36).F41(), _264_v1))).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(433), (_484_v153).Cardinality())) >= 0, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_271_v8), 0))
			_ = _index36
			(_271_v8).ArraySet1(_263_v0, (_index36).Int())
			var _485_v154 _dafny.Sequence
			_ = _485_v154
			_485_v154 = _dafny.SeqOf(_263_v0, _263_v0, _263_v0)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_271_v8), 0))
			_ = _index37
			(_271_v8).ArraySet1((_485_v154).Select((Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_485_v154).Cardinality()))).Uint32()).(_dafny.Array), (_index37).Int())
		} else {
			var _486___mcc_h15 D10 = _source9.Get_().(D10_DC32).Cf63
			_ = _486___mcc_h15
			var _487_cf63 D10 = _486___mcc_h15
			_ = _487_cf63
			_418_v110 = _dafny.MultiSetOf((_dafny.SetOf(_420_v112.F29())).IsSubsetOf(_349_v70), true)
			(_275_globalState).F13 = _dafny.IntOfInt64(481)
			var _488_v155 _dafny.Int
			_ = _488_v155
			var _out29 _dafny.Int
			_ = _out29
			_out29 = Companion_Default___.M0(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(555))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_489_v120 T3) func(_dafny.Int) _dafny.Int {
				return func(_490_i15 _dafny.Int) _dafny.Int {
					return (_489_v120).F34()
				}
			})(_431_v120))), !(_420_v112.F28()), (_dafny.MultiSetOf(_420_v112.F29(), _431_v120.F35())).Cardinality(), _431_v120.F29(), _275_globalState)
			_488_v155 = _out29
			var _491_v156 _dafny.Set
			_ = _491_v156
			_491_v156 = _dafny.SetOf((_431_v120).F30(), _428_v117, (_431_v120).F30(), _264_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v112.F28(), _415_v109)).Cardinality())
			var _492_v157 T2
			_ = _492_v157
			var _nw56 *C6 = New_C6_()
			_ = _nw56
			_nw56.Ctor__(_426_v115, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_269_v6, _429_v118)).Cardinality()), (_dafny.SetOf((_306_v36).F41(), _dafny.IntOfInt64(313))).IsSubsetOf(_491_v156), _345_v66, (_431_v120).F33(), _428_v117, _431_v120.F35(), ((_431_v120).F32()).Select((Companion_Default___.SafeIndex(_428_v117, _dafny.IntOfUint32(((_431_v120).F32()).Cardinality()))).Uint32()).(bool), _420_v112.F28())
			_492_v157 = _nw56
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _index38
			var _rhs43 _dafny.Int = _488_v155
			_ = _rhs43
			var _rhs44 bool = false
			_ = _rhs44
			var _rhs45 T2 = _492_v157
			_ = _rhs45
			var _rhs46 bool = _492_v157.F29()
			_ = _rhs46
			var _lhs35 _dafny.Array = _263_v0
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _lhs36
			var _lhs37 *GlobalState = _275_globalState
			_ = _lhs37
			var _lhs38 T2 = _492_v157
			_ = _lhs38
			(_lhs35).ArraySet1(_rhs43, (_lhs36).Int())
			_lhs37.F12 = _rhs44
			_492_v157 = _rhs45
			_lhs38.F31_set_(_rhs46)
		}
		if false {
			var _493_v158 _dafny.Array
			_ = _493_v158
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
			_ = _nw57
			_493_v158 = _nw57
			var _494_v159 *C0
			_ = _494_v159
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__(_493_v158)
			_494_v159 = _nw58
			var _495_v160 _dafny.Sequence
			_ = _495_v160
			_495_v160 = _dafny.SeqOf(_494_v159, _494_v159)
			var _rhs47 _dafny.Int = (_dafny.Zero).Minus(((_306_v36).Fm20(_dafny.IntOfUint32((_495_v160).Cardinality()), _431_v120.F28(), _275_globalState)).Minus((_306_v36).F41()))
			_ = _rhs47
			var _rhs48 _dafny.Int = (_dafny.Zero).Minus((_431_v120).F34())
			_ = _rhs48
			var _lhs39 *GlobalState = _275_globalState
			_ = _lhs39
			_lhs39.F13 = _rhs47
			_264_v1 = _rhs48
			var _496_v161 D1
			_ = _496_v161
			_496_v161 = Companion_D1_.Create_DC3_(_431_v120.F31(), _428_v117, _494_v159, Companion_Default___.Fm32(_420_v112.F29(), _275_globalState))
			(_275_globalState).F13 = (_496_v161).Dtor_cf4()
			var _497_v162 _dafny.Set
			_ = _497_v162
			_497_v162 = _dafny.SetOf(_265_v2, _dafny.SeqOf(_264_v1, (_431_v120).F34()))
			var _498_v163 D16
			_ = _498_v163
			_498_v163 = Companion_D16_.Create_DC39_((_497_v162).Difference(_497_v162))
			_498_v163 = Companion_D16_.Create_DC39_(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(405))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}((func(_499_v120 T3) func(_dafny.Int) _dafny.Int {
				return func(_500_i16 _dafny.Int) _dafny.Int {
					return (_499_v120).F34()
				}
			})(_431_v120)))))
			(_431_v120).F31_set_((_431_v120).Fm6(_275_globalState))
			var _501_v164 _dafny.Array
			_ = _501_v164
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw59
			_501_v164 = _nw59
			var _502_v165 *C3
			_ = _502_v165
			var _nw60 *C3 = New_C3_()
			_ = _nw60
			_nw60.Ctor__(_dafny.IntOfUint32((_265_v2).Cardinality()), _501_v164, _dafny.CodePoint('d'), (_dafny.Zero).Minus((_306_v36).F41()), _431_v120.F28(), (_431_v120).F32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(18))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}((func(_503_v120 T3) func(_dafny.Int) _dafny.Sequence {
				return func(_504_i17 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_503_v120.F28(), _503_v120.F31())
				}
			})(_431_v120))), (_dafny.IntOfUint32(((_431_v120).F32()).Cardinality())).Plus((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int)), _420_v112.F28(), !(_431_v120.F29()), (_264_v1).Cmp((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int)) != 0)
			_502_v165 = _nw60
		} else {
			var _505_v166 *C10
			_ = _505_v166
			var _nw61 *C10 = New_C10_()
			_ = _nw61
			_nw61.Ctor__((func() _dafny.CodePoint {
				if (_306_v36).Fm19(_275_globalState) {
					return _310_v40
				}
				return (_429_v118).Select((Companion_Default___.SafeIndex((_431_v120).F34(), _dafny.IntOfUint32((_429_v118).Cardinality()))).Uint32()).(_dafny.CodePoint)
			})(), (_431_v120).F30(), _420_v112.F28(), (_431_v120).F32(), (_431_v120).F33(), (_431_v120).F34(), _420_v112.F29(), _420_v112.F29(), ((_306_v36).F42()).Equals(Companion_Default___.Fm25(_275_globalState)))
			_505_v166 = _nw61
			var _rhs49 bool = (((_306_v36).F42()).Intersection((_306_v36).F42())).IsProperSubsetOf((_306_v36).F42())
			_ = _rhs49
			var _rhs50 bool = (func() bool {
				if _431_v120.F29() {
					return _431_v120.F31()
				}
				return Companion_Default___.Fm1(_267_v4, _dafny.IntOfInt64(419), _275_globalState)
			})()
			_ = _rhs50
			var _rhs51 bool = ((_306_v36).F41()).Cmp((_306_v36).F41()) >= 0
			_ = _rhs51
			var _rhs52 *C10 = _505_v166
			_ = _rhs52
			var _lhs40 *GlobalState = _275_globalState
			_ = _lhs40
			var _lhs41 *GlobalState = _275_globalState
			_ = _lhs41
			var _lhs42 T0 = _420_v112
			_ = _lhs42
			_lhs40.F1 = _rhs49
			_lhs41.F12 = _rhs50
			_lhs42.F28_set_(_rhs51)
			_505_v166 = _rhs52
			var _506_v167 *C8
			_ = _506_v167
			var _nw62 *C8 = New_C8_()
			_ = _nw62
			_nw62.Ctor__(_431_v120.F29(), Companion_Default___.Fm48(_431_v120.F28(), _431_v120.F29(), (_431_v120).F30(), _275_globalState), ((_306_v36).F41()).Plus((_306_v36).F41()), _420_v112.F28(), _345_v66, _dafny.Companion_Sequence_.Concatenate(_414_v108, (_431_v120).F33()), (_306_v36).F41(), _420_v112.F28(), (_311_v41).IsProperSubsetOf(_311_v41), false)
			_506_v167 = _nw62
			var _507_v168 T4
			_ = _507_v168
			var _nw63 *C8 = New_C8_()
			_ = _nw63
			_nw63.Ctor__(_420_v112.F28(), _426_v115, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v120.F29(), _431_v120.F31())).Cardinality(), _431_v120.F29(), (_431_v120).F32(), _dafny.Companion_Sequence_.Concatenate(_414_v108, (_431_v120).F33()), (_306_v36).F41(), _431_v120.F28(), ((_431_v120).F30()).Cmp((_306_v36).F41()) >= 0, _431_v120.F31())
			_507_v168 = _nw63
			_507_v168 = _507_v168
			var _508_v169 _dafny.Map
			_ = _508_v169
			_508_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v120.F35(), (_507_v168).Fm7((_306_v36).F42(), (_431_v120).F34(), _275_globalState))
			var _509_v170 _dafny.Map
			_ = _509_v170
			_509_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v120.F28(), (_508_v169).Merge(_508_v169))
			var _510_v171 _dafny.Array
			_ = _510_v171
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw64
			_510_v171 = _nw64
			var _511_v172 _dafny.Array
			_ = _511_v172
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_7
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = (func(_512_v118 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_513_i18 _dafny.Int) _dafny.Sequence {
						return _512_v118
					}
				})(_429_v118)
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
			_511_v172 = _nw65
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_511_v172), 0))
			_ = _index39
			(_511_v172).ArraySet1((_431_v120).Fm3(_431_v120.F35(), !(_431_v120.F31()), _275_globalState), (_index39).Int())
			var _514_v173 _dafny.Map
			_ = _514_v173
			_514_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v39, true)
			var _515_v175 _dafny.Array
			_ = _515_v175
			var _nwElement0_9 _dafny.CodePoint = _310_v40
			_ = _nwElement0_9
			var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(15))
			_ = _nw66
			(_nw66).ArraySet1CodePoint(_nwElement0_9, 0)
			(_nw66).ArraySet1CodePoint(_426_v115, 1)
			(_nw66).ArraySet1CodePoint(_426_v115, 2)
			(_nw66).ArraySet1CodePoint(_310_v40, 3)
			(_nw66).ArraySet1CodePoint(_507_v168.F36(), 4)
			(_nw66).ArraySet1CodePoint(_310_v40, 5)
			(_nw66).ArraySet1CodePoint(_426_v115, 6)
			(_nw66).ArraySet1CodePoint(_426_v115, 7)
			(_nw66).ArraySet1CodePoint(_507_v168.F36(), 8)
			(_nw66).ArraySet1CodePoint(_507_v168.F36(), 9)
			(_nw66).ArraySet1CodePoint(_507_v168.F36(), 10)
			(_nw66).ArraySet1CodePoint(_507_v168.F36(), 11)
			(_nw66).ArraySet1CodePoint(_426_v115, 12)
			(_nw66).ArraySet1CodePoint(_310_v40, 13)
			(_nw66).ArraySet1CodePoint(_310_v40, 14)
			_515_v175 = _nw66
			var _516_v176 *C0
			_ = _516_v176
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__(_515_v175)
			_516_v176 = _nw67
			var _517_v177 _dafny.Map
			_ = _517_v177
			_517_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm39(_431_v120.F35(), (_306_v36).F42(), _431_v120.F28(), _275_globalState), _516_v176)
			var _518_v178 _dafny.Map
			_ = _518_v178
			_518_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_517_v177).Cardinality(), _431_v120.F29())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_511_v172), 0))
			_ = _index40
			var _rhs53 _dafny.Map = Companion_Default___.Fm67(_275_globalState)
			_ = _rhs53
			var _rhs54 bool = !((func() bool {
				if _431_v120.F31() {
					return (func() bool {
						if (_514_v173).Contains(func() _dafny.Map {
							var _coll21 = _dafny.NewMapBuilder()
							_ = _coll21
							for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(783), _dafny.IntOfInt64(265))); ; {
								_compr_21, _ok21 := _iter21()
								if !_ok21 {
									break
								}
								var _519_v174 _dafny.Int
								_519_v174 = interface{}(_compr_21).(_dafny.Int)
								if ((_dafny.IntOfInt64(783)).Cmp(_519_v174) <= 0) && ((_519_v174).Cmp(_dafny.IntOfInt64(265)) < 0) {
									_coll21.Add(Companion_Default___.SafeDivisionInt(_519_v174, (_507_v168).F34()), (_507_v168).F30())
								}
							}
							return _coll21.ToMap()
						}()) {
							return (_514_v173).Get(func() _dafny.Map {
								var _coll22 = _dafny.NewMapBuilder()
								_ = _coll22
								for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(783), _dafny.IntOfInt64(265))); ; {
									_compr_22, _ok22 := _iter22()
									if !_ok22 {
										break
									}
									var _520_v174 _dafny.Int
									_520_v174 = interface{}(_compr_22).(_dafny.Int)
									if ((_dafny.IntOfInt64(783)).Cmp(_520_v174) <= 0) && ((_520_v174).Cmp(_dafny.IntOfInt64(265)) < 0) {
										_coll22.Add(Companion_Default___.SafeDivisionInt(_520_v174, (_507_v168).F34()), (_507_v168).F30())
									}
								}
								return _coll22.ToMap()
							}()).(bool)
						}
						return _420_v112.F29()
					})()
				}
				return (func() bool {
					if (_518_v178).Contains((_507_v168).F30()) {
						return (_518_v178).Get((_507_v168).F30()).(bool)
					}
					return _420_v112.F29()
				})()
			})())
			_ = _rhs54
			var _rhs55 _dafny.Array = _510_v171
			_ = _rhs55
			var _rhs56 _dafny.Sequence = _429_v118
			_ = _rhs56
			var _rhs57 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if !(_420_v112.F28()) {
					return _429_v118
				}
				return _429_v118
			})(), (func() _dafny.Sequence {
				if _507_v168.F29() {
					return _dafny.UnicodeSeqOfUtf8Bytes("lqm")
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("xb")
			})())
			_ = _rhs57
			var _lhs43 T3 = _431_v120
			_ = _lhs43
			var _lhs44 _dafny.Array = _511_v172
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_511_v172), 0))
			_ = _lhs45
			_509_v170 = _rhs53
			_lhs43.F29_set_(_rhs54)
			_510_v171 = _rhs55
			_429_v118 = _rhs56
			(_lhs44).ArraySet1(_rhs57, (_lhs45).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_272_v9), 0))
			_ = _index41
			(_272_v9).ArraySet1(false, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_272_v9), 0))
			_ = _index42
			(_272_v9).ArraySet1(((_431_v120).F32()).Select((Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32(((_431_v120).F32()).Cardinality()))).Uint32()).(bool), (_index42).Int())
		}
		var _521_v179 *C1
		_ = _521_v179
		var _nw68 *C1 = New_C1_()
		_ = _nw68
		_nw68.Ctor__(_431_v120.F35(), _431_v120.F29(), _420_v112.F29())
		_521_v179 = _nw68
		var _522_v180 D10
		_ = _522_v180
		_522_v180 = Companion_D10_.Create_DC28_(_521_v179)
		var _pat_let_tv3 = _521_v179
		_ = _pat_let_tv3
		var _523_v181 _dafny.Array
		_ = _523_v181
		var _nwElement0_10 D10 = _522_v180
		_ = _nwElement0_10
		var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(26))
		_ = _nw69
		(_nw69).ArraySet1(_nwElement0_10, 0)
		(_nw69).ArraySet1(Companion_D10_.Create_DC28_(_521_v179), 1)
		(_nw69).ArraySet1(_522_v180, 2)
		(_nw69).ArraySet1(Companion_D10_.Create_DC28_(_521_v179), 3)
		(_nw69).ArraySet1(Companion_D10_.Create_DC28_(_521_v179), 4)
		(_nw69).ArraySet1(_522_v180, 5)
		(_nw69).ArraySet1(_522_v180, 6)
		(_nw69).ArraySet1(_522_v180, 7)
		(_nw69).ArraySet1(_522_v180, 8)
		(_nw69).ArraySet1(_522_v180, 9)
		(_nw69).ArraySet1(func(_pat_let5_0 D10) D10 {
			return func(_524_dt__update__tmp_h2 D10) D10 {
				return func(_pat_let6_0 *C1) D10 {
					return func(_525_dt__update_hcf55_h0 *C1) D10 {
						return Companion_D10_.Create_DC28_(_525_dt__update_hcf55_h0)
					}(_pat_let6_0)
				}(_pat_let_tv3)
			}(_pat_let5_0)
		}(_522_v180), 10)
		(_nw69).ArraySet1(_522_v180, 11)
		(_nw69).ArraySet1(_522_v180, 12)
		(_nw69).ArraySet1(_522_v180, 13)
		(_nw69).ArraySet1(_522_v180, 14)
		(_nw69).ArraySet1(Companion_D10_.Create_DC28_((_522_v180).Dtor_cf55()), 15)
		(_nw69).ArraySet1(_522_v180, 16)
		(_nw69).ArraySet1(_522_v180, 17)
		(_nw69).ArraySet1(_522_v180, 18)
		(_nw69).ArraySet1(_522_v180, 19)
		(_nw69).ArraySet1(_522_v180, 20)
		(_nw69).ArraySet1(_522_v180, 21)
		(_nw69).ArraySet1(_522_v180, 22)
		(_nw69).ArraySet1(_522_v180, 23)
		(_nw69).ArraySet1(Companion_D10_.Create_DC28_(_521_v179), 24)
		(_nw69).ArraySet1(_522_v180, 25)
		_523_v181 = _nw69
		var _526_v182 D20
		_ = _526_v182
		_526_v182 = Companion_D20_.Create_DC48_((_431_v120).F30(), _523_v181, _431_v120.F29())
		_526_v182 = _526_v182
	} else {
		var _527_v183 D3
		_ = _527_v183
		_527_v183 = Companion_D3_.Create_DC9_(false, !(_267_v4), _272_v9, _267_v4)
		var _528_v184 _dafny.Array
		_ = _528_v184
		var _529_v185 _dafny.Int
		_ = _529_v185
		var _out30 _dafny.Array
		_ = _out30
		var _out31 _dafny.Int
		_ = _out31
		_out30, _out31 = (_306_v36).M8((_527_v183).Dtor_cf20(), _275_globalState)
		_528_v184 = _out30
		_529_v185 = _out31
		var _530_v186 _dafny.Array
		_ = _530_v186
		var _nw70 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
		_ = _nw70
		_530_v186 = _nw70
		var _531_v187 T3
		_ = _531_v187
		var _nw71 *C5 = New_C5_()
		_ = _nw71
		_nw71.Ctor__(!(true), _264_v1, _529_v185, _267_v4, _345_v66, _414_v108, (_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _267_v4, _267_v4, _267_v4)
		_531_v187 = _nw71
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_530_v186), 0))
		_ = _index43
		(_530_v186).ArraySet1(_531_v187, (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_530_v186), 0))
		_ = _index44
		(_530_v186).ArraySet1(_531_v187, (_index44).Int())
		var _532_v188 _dafny.Sequence
		_ = _532_v188
		_532_v188 = _dafny.SeqOf(_415_v109, _415_v109, _415_v109, _415_v109, _415_v109)
		var _533_v189 D17
		_ = _533_v189
		_533_v189 = Companion_D17_.Create_DC42_(_dafny.Companion_Sequence_.Concatenate(_532_v188, _532_v188))
		var _534_v190 _dafny.Map
		_ = _534_v190
		_534_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v187.F31(), _532_v188)
		_533_v189 = Companion_D17_.Create_DC42_((func() _dafny.Sequence {
			if (_534_v190).Contains((_531_v187).Fm6(_275_globalState)) {
				return (_534_v190).Get((_531_v187).Fm6(_275_globalState)).(_dafny.Sequence)
			}
			return _532_v188
		})())
		var _535_v191 _dafny.Array
		_ = _535_v191
		var _536_v192 _dafny.Int
		_ = _536_v192
		var _out32 _dafny.Array
		_ = _out32
		var _out33 _dafny.Int
		_ = _out33
		_out32, _out33 = (_306_v36).M8(_272_v9, _275_globalState)
		_535_v191 = _out32
		_536_v192 = _out33
		var _537_v193 _dafny.Array
		_ = _537_v193
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_8
		var _nw72 _dafny.Array
		_ = _nw72
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw72 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.CodePoint = (func(_538_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_539_i19 _dafny.Int) _dafny.CodePoint {
					return _538_v40
				}
			})(_310_v40)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw72 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw72).ArraySet1CodePoint(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw72).ArraySet1CodePoint(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_537_v193 = _nw72
		var _540_v194 *C0
		_ = _540_v194
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(_537_v193)
		_540_v194 = _nw73
		var _541_v195 D1
		_ = _541_v195
		_541_v195 = Companion_D1_.Create_DC3_(_531_v187.F28(), (_dafny.Zero).Minus(_529_v185), _540_v194, (_265_v2).Select((Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_265_v2).Cardinality()))).Uint32()).(_dafny.Int))
		var _542_v196 *C1
		_ = _542_v196
		var _nw74 *C1 = New_C1_()
		_ = _nw74
		_nw74.Ctor__(_531_v187.F28(), _531_v187.F28(), !((_541_v195).Dtor_cf3()))
		_542_v196 = _nw74
	}
	var _543_v197 _dafny.Sequence
	_ = _543_v197
	_543_v197 = _dafny.SeqOf(_348_v69)
	if (_dafny.MultiSetOf(_543_v197)).Equals(_dafny.MultiSetOf(_543_v197, _543_v197, _543_v197)) {
		_309_v39 = (_309_v39).Update((_306_v36).F41(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fy")).Cardinality()))
		var _544_v198 T1
		_ = _544_v198
		var _nw75 *C7 = New_C7_()
		_ = _nw75
		_nw75.Ctor__((func() _dafny.Int {
			if ((_306_v36).F42()).Contains(_dafny.IntOfInt64(-183)) {
				return ((_306_v36).F42()).Multiplicity(_dafny.IntOfInt64(-183))
			}
			return _264_v1
		})(), (_306_v36).F42(), (_349_v70).Cardinality(), _267_v4, _267_v4, _267_v4)
		_544_v198 = _nw75
		var _545_v199 _dafny.Sequence
		_ = _545_v199
		_545_v199 = _dafny.SeqOf(_544_v198, _544_v198, _544_v198, _544_v198)
		var _546_v200 _dafny.Map
		_ = _546_v200
		_546_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v109, _545_v199)
		var _547_v201 _dafny.Set
		_ = _547_v201
		_547_v201 = _dafny.SetOf(_dafny.IntOfUint32((_269_v6).Cardinality()))
		var _548_v202 _dafny.Sequence
		_ = _548_v202
		_548_v202 = _dafny.SeqOf(_545_v199, _545_v199)
		var _549_v203 _dafny.Sequence
		_ = _549_v203
		_549_v203 = _545_v199
		var _550_v204 _dafny.Array
		_ = _550_v204
		var _nwElement0_11 _dafny.Sequence = (func() _dafny.Sequence {
			if _267_v4 {
				return _545_v199
			}
			return _545_v199
		})()
		_ = _nwElement0_11
		var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(27))
		_ = _nw76
		(_nw76).ArraySet1(_nwElement0_11, 0)
		(_nw76).ArraySet1(_545_v199, 1)
		(_nw76).ArraySet1(_545_v199, 2)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 3)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 4)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 5)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _dafny.SeqOf(_544_v198)), 6)
		(_nw76).ArraySet1(_545_v199, 7)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 8)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_545_v199, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.IntOfUint32((_545_v199).Cardinality()))).Uint32(), _544_v198), _545_v199), 9)
		(_nw76).ArraySet1(_545_v199, 10)
		(_nw76).ArraySet1(_545_v199, 11)
		(_nw76).ArraySet1(_dafny.SeqOf(_544_v198), 12)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Update(_545_v199, (Companion_Default___.SafeIndex((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_545_v199).Cardinality()))).Uint32(), _544_v198), 13)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 14)
		(_nw76).ArraySet1(_545_v199, 15)
		(_nw76).ArraySet1(_545_v199, 16)
		(_nw76).ArraySet1(_545_v199, 17)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, (func() _dafny.Sequence {
			if (_546_v200).Contains(_415_v109) {
				return (_546_v200).Get(_415_v109).(_dafny.Sequence)
			}
			return _545_v199
		})()), 18)
		(_nw76).ArraySet1(_dafny.SeqOf(_544_v198, _544_v198, _544_v198, _544_v198, _544_v198), 19)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_545_v199, (Companion_Default___.SafeIndex((_547_v201).Cardinality(), _dafny.IntOfUint32((_545_v199).Cardinality()))).Uint32(), _544_v198), (Companion_Default___.SafeIndex((_306_v36).F41(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_545_v199, (Companion_Default___.SafeIndex((_547_v201).Cardinality(), _dafny.IntOfUint32((_545_v199).Cardinality()))).Uint32(), _544_v198)).Cardinality()))).Uint32(), _544_v198)), 20)
		(_nw76).ArraySet1((_548_v202).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if ((_306_v36).F42()).Contains(_dafny.IntOfInt64(576)) {
				return ((_306_v36).F42()).Multiplicity(_dafny.IntOfInt64(576))
			}
			return _dafny.IntOfInt64(-151)
		})(), _dafny.IntOfUint32((_548_v202).Cardinality()))).Uint32()).(_dafny.Sequence), 21)
		(_nw76).ArraySet1((_549_v203), 22)
		(_nw76).ArraySet1(_dafny.SeqOf(_544_v198), 23)
		(_nw76).ArraySet1(_545_v199, 24)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 25)
		(_nw76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _545_v199), 26)
		_550_v204 = _nw76
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_550_v204), 0))
		_ = _index45
		(_550_v204).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_545_v199, _dafny.SeqOf(_544_v198)), (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_550_v204), 0))
		_ = _index46
		(_550_v204).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if _544_v198.F28() {
				return _545_v199
			}
			return _545_v199
		})(), _545_v199), (_index46).Int())
		var _551_v205 _dafny.CodePoint
		_ = _551_v205
		var _552_v206 _dafny.Sequence
		_ = _552_v206
		var _553_v207 _dafny.Int
		_ = _553_v207
		var _554_v208 _dafny.Sequence
		_ = _554_v208
		var _out34 _dafny.CodePoint
		_ = _out34
		var _out35 _dafny.Sequence
		_ = _out35
		var _out36 _dafny.Int
		_ = _out36
		var _out37 _dafny.Sequence
		_ = _out37
		_out34, _out35, _out36, _out37 = (_306_v36).M7(_544_v198.F31(), _272_v9, _544_v198.F29(), _275_globalState)
		_551_v205 = _out34
		_552_v206 = _out35
		_553_v207 = _out36
		_554_v208 = _out37
		var _555_v209 *C1
		_ = _555_v209
		var _nw77 *C1 = New_C1_()
		_ = _nw77
		_nw77.Ctor__(!((_306_v36).Fm19(_275_globalState)), false, _544_v198.F28())
		_555_v209 = _nw77
		var _556_v210 D10
		_ = _556_v210
		_556_v210 = Companion_D10_.Create_DC28_(_555_v209)
		var _pat_let_tv4 = _555_v209
		_ = _pat_let_tv4
		var _557_v211 _dafny.Array
		_ = _557_v211
		var _nwElement0_12 D10 = _556_v210
		_ = _nwElement0_12
		var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(17))
		_ = _nw78
		(_nw78).ArraySet1(_nwElement0_12, 0)
		(_nw78).ArraySet1(_556_v210, 1)
		(_nw78).ArraySet1(_556_v210, 2)
		(_nw78).ArraySet1(_556_v210, 3)
		(_nw78).ArraySet1(_556_v210, 4)
		(_nw78).ArraySet1(_556_v210, 5)
		(_nw78).ArraySet1(_556_v210, 6)
		(_nw78).ArraySet1(_556_v210, 7)
		(_nw78).ArraySet1(func(_pat_let7_0 D10) D10 {
			return func(_558_dt__update__tmp_h3 D10) D10 {
				return func(_pat_let8_0 *C1) D10 {
					return func(_559_dt__update_hcf55_h1 *C1) D10 {
						return Companion_D10_.Create_DC28_(_559_dt__update_hcf55_h1)
					}(_pat_let8_0)
				}(_pat_let_tv4)
			}(_pat_let7_0)
		}(_556_v210), 8)
		(_nw78).ArraySet1(_556_v210, 9)
		(_nw78).ArraySet1(_556_v210, 10)
		(_nw78).ArraySet1(_556_v210, 11)
		(_nw78).ArraySet1(_556_v210, 12)
		(_nw78).ArraySet1(_556_v210, 13)
		(_nw78).ArraySet1(_556_v210, 14)
		(_nw78).ArraySet1(_556_v210, 15)
		(_nw78).ArraySet1(_556_v210, 16)
		_557_v211 = _nw78
		var _560_v212 D20
		_ = _560_v212
		_560_v212 = Companion_D20_.Create_DC50_(Companion_D20_.Create_DC48_((_306_v36).F41(), _557_v211, (_555_v209).F43()))
		var _source10 D20 = _560_v212
		_ = _source10
		if _source10.Is_DC48() {
			var _561___mcc_h16 _dafny.Int = _source10.Get_().(D20_DC48).Cf82
			_ = _561___mcc_h16
			var _562___mcc_h17 _dafny.Array = _source10.Get_().(D20_DC48).Cf83
			_ = _562___mcc_h17
			var _563___mcc_h18 bool = _source10.Get_().(D20_DC48).Cf84
			_ = _563___mcc_h18
			var _564_cf84 bool = _563___mcc_h18
			_ = _564_cf84
			var _565_cf83 _dafny.Array = _562___mcc_h17
			_ = _565_cf83
			var _566_cf82 _dafny.Int = _561___mcc_h16
			_ = _566_cf82
			var _567_v213 bool
			_ = _567_v213
			var _568_v214 _dafny.Int
			_ = _568_v214
			var _569_v215 _dafny.Sequence
			_ = _569_v215
			var _out38 bool
			_ = _out38
			var _out39 _dafny.Int
			_ = _out39
			var _out40 _dafny.Sequence
			_ = _out40
			_out38, _out39, _out40 = (_415_v109).M1(_275_globalState)
			_567_v213 = _out38
			_568_v214 = _out39
			_569_v215 = _out40
			_568_v214 = (Companion_Default___.SafeModuloInt((_306_v36).F41(), _dafny.IntOfInt64(50))).Minus(_dafny.IntOfUint32((_345_v66).Cardinality()))
			var _570_v216 _dafny.Sequence
			_ = _570_v216
			_570_v216 = _dafny.SeqOf(_349_v70)
			_349_v70 = (_570_v216).Select((Companion_Default___.SafeIndex(_553_v207, _dafny.IntOfUint32((_570_v216).Cardinality()))).Uint32()).(_dafny.Set)
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _index47
			(_263_v0).ArraySet1((_306_v36).F41(), (_index47).Int())
		} else if _source10.Is_DC49() {
			var _571___mcc_h19 bool = _source10.Get_().(D20_DC49).Cf85
			_ = _571___mcc_h19
			var _572___mcc_h20 _dafny.CodePoint = _source10.Get_().(D20_DC49).Cf86
			_ = _572___mcc_h20
			var _573_cf86 _dafny.CodePoint = _572___mcc_h20
			_ = _573_cf86
			var _574_cf85 bool = _571___mcc_h19
			_ = _574_cf85
			(_544_v198).F31_set_((_555_v209).F43())
			var _575_v217 bool
			_ = _575_v217
			var _576_v218 _dafny.Int
			_ = _576_v218
			var _577_v219 _dafny.Sequence
			_ = _577_v219
			var _out41 bool
			_ = _out41
			var _out42 _dafny.Int
			_ = _out42
			var _out43 _dafny.Sequence
			_ = _out43
			_out41, _out42, _out43 = (_415_v109).M1(_275_globalState)
			_575_v217 = _out41
			_576_v218 = _out42
			_577_v219 = _out43
			(_544_v198).F29_set_(_dafny.Companion_Sequence_.Equal(_269_v6, (func() _dafny.Sequence {
				if _267_v4 {
					return _577_v219
				}
				return _269_v6
			})()))
			var _578_v220 D4
			_ = _578_v220
			_578_v220 = Companion_D4_.Create_DC11_((_544_v198).F30(), _574_cf85, _dafny.IntOfInt64(854))
			var _579_v221 _dafny.Map
			_ = _579_v221
			_579_v221 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_551_v205, (_555_v209).F43())
			var _580_v222 *C10
			_ = _580_v222
			var _nw79 *C10 = New_C10_()
			_ = _nw79
			_nw79.Ctor__(_551_v205, (_578_v220).Dtor_cf25(), (func() bool {
				if (_555_v209).Fm22(_576_v218, Companion_Default___.Fm24(false, (_306_v36).F41(), _275_globalState), _544_v198.F29(), _275_globalState) {
					return _544_v198.F29()
				}
				return (_555_v209).F43()
			})(), _345_v66, _414_v108, _264_v1, _544_v198.F31(), _544_v198.F31(), (func() bool {
				if (_579_v221).Contains(_573_cf86) {
					return (_579_v221).Get(_573_cf86).(bool)
				}
				return _544_v198.F28()
			})())
			_580_v222 = _nw79
		} else if _source10.Is_DC47() {
			var _581___mcc_h21 _dafny.Array = _source10.Get_().(D20_DC47).Cf81
			_ = _581___mcc_h21
			var _582_cf81 _dafny.Array = _581___mcc_h21
			_ = _582_cf81
			var _583_v223 T3
			_ = _583_v223
			var _nw80 *C4 = New_C4_()
			_ = _nw80
			_nw80.Ctor__(_264_v1, _544_v198.F31(), _345_v66, _414_v108, (_306_v36).F41(), (_555_v209).F43(), _544_v198.F31(), _544_v198.F29())
			_583_v223 = _nw80
			var _584_v224 _dafny.Sequence
			_ = _584_v224
			_584_v224 = _dafny.SeqOf(_583_v223)
			var _585_v225 *C10
			_ = _585_v225
			var _nw81 *C10 = New_C10_()
			_ = _nw81
			_nw81.Ctor__(_310_v40, _dafny.IntOfUint32((_584_v224).Cardinality()), _583_v223.F31(), (_583_v223).F32(), _414_v108, _553_v207, _583_v223.F28(), _267_v4, _583_v223.F29())
			_585_v225 = _nw81
			var _586_v226 _dafny.Array
			_ = _586_v226
			var _nwElement0_13 *C10 = _585_v225
			_ = _nwElement0_13
			var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(20))
			_ = _nw82
			(_nw82).ArraySet1(_nwElement0_13, 0)
			(_nw82).ArraySet1(_585_v225, 1)
			(_nw82).ArraySet1(_585_v225, 2)
			(_nw82).ArraySet1(_585_v225, 3)
			(_nw82).ArraySet1(_585_v225, 4)
			(_nw82).ArraySet1(_585_v225, 5)
			(_nw82).ArraySet1(_585_v225, 6)
			(_nw82).ArraySet1(_585_v225, 7)
			(_nw82).ArraySet1(_585_v225, 8)
			(_nw82).ArraySet1(_585_v225, 9)
			(_nw82).ArraySet1(_585_v225, 10)
			(_nw82).ArraySet1(_585_v225, 11)
			(_nw82).ArraySet1(_585_v225, 12)
			(_nw82).ArraySet1(_585_v225, 13)
			(_nw82).ArraySet1(_585_v225, 14)
			(_nw82).ArraySet1(_585_v225, 15)
			(_nw82).ArraySet1(_585_v225, 16)
			(_nw82).ArraySet1(_585_v225, 17)
			(_nw82).ArraySet1(_585_v225, 18)
			(_nw82).ArraySet1(_585_v225, 19)
			_586_v226 = _nw82
			_586_v226 = _586_v226
			(_275_globalState).F13 = (_583_v223).F30()
			var _587_v227 _dafny.Map
			_ = _587_v227
			_587_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_583_v223.F28(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(834))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}((func(_588_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_589_i20 _dafny.Int) _dafny.CodePoint {
					return _588_v40
				}
			})(_310_v40)))).Cardinality()))
			var _590_v228 _dafny.Set
			_ = _590_v228
			_590_v228 = _dafny.SetOf(_544_v198)
			var _591_v229 _dafny.Map
			_ = _591_v229
			_591_v229 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-925), Companion_Default___.Fm41(_583_v223.F31(), _310_v40, _544_v198.F28(), (_590_v228).Cardinality(), _275_globalState))
			var _592_v230 T2
			_ = _592_v230
			var _nw83 *C6 = New_C6_()
			_ = _nw83
			_nw83.Ctor__(_310_v40, _dafny.IntOfInt64(925), _583_v223.F35(), _dafny.SeqOf(_583_v223.F31()), (_583_v223).F33(), (func() _dafny.Int {
				if (_587_v227).Contains(false) {
					return (_587_v227).Get(false).(_dafny.Int)
				}
				return (_591_v229).Cardinality()
			})(), _583_v223.F31(), false, _544_v198.F31())
			_592_v230 = _nw83
			var _593_v231 _dafny.Map
			_ = _593_v231
			_593_v231 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v230, (_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int))
			var _594_v232 _dafny.Sequence
			_ = _594_v232
			_594_v232 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v230, ((_306_v36).F42()).Cardinality()), _593_v231)
			var _595_v233 *C6
			_ = _595_v233
			var _nw84 *C6 = New_C6_()
			_ = _nw84
			_nw84.Ctor__(_310_v40, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(723), (_587_v227).Cardinality()), (_547_v201).IsSubsetOf(_547_v201), _345_v66, Companion_Default___.Fm51(_544_v198.F31(), _310_v40, _275_globalState), _264_v1, _544_v198.F28(), _544_v198.F29(), _dafny.Companion_Sequence_.Contains(_594_v232, _593_v231))
			_595_v233 = _nw84
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_272_v9), 0))
			_ = _index48
			(_272_v9).ArraySet1(!(_583_v223.F31()), (_index48).Int())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(365), _dafny.ArrayLen((_272_v9), 0))
			_ = _index49
			(_272_v9).ArraySet1(_592_v230.F31(), (_index49).Int())
		} else {
			var _596___mcc_h22 D20 = _source10.Get_().(D20_DC50).Cf87
			_ = _596___mcc_h22
			var _597_cf87 D20 = _596___mcc_h22
			_ = _597_cf87
			var _598_v234 _dafny.Array
			_ = _598_v234
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_9
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = func(_599_i21 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("vtf")
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw85 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw85).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw85).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_598_v234 = _nw85
			_598_v234 = _598_v234
			(_275_globalState).F13 = (_553_v207).Plus((func() _dafny.Int {
				if _544_v198.F31() {
					return (_dafny.Zero).Minus((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int))
				}
				return _553_v207
			})())
			_553_v207 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_544_v198).F30()), (_306_v36).F41()))
			var _600_v235 _dafny.Map
			_ = _600_v235
			_600_v235 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_547_v201, _269_v6)
			(_275_globalState).F13 = Companion_Default___.Fm32(!((_600_v235).Equals(_600_v235)), _275_globalState)
		}
		var _601_v236 _dafny.Array
		_ = _601_v236
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_10
		var _nw86 _dafny.Array
		_ = _nw86
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw86 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Map = (func(_602_v198 T1) func(_dafny.Int) _dafny.Map {
				return func(_603_i22 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_602_v198.F31(), !(_602_v198.F29()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_602_v198.F28(), !(false)))
				}
			})(_544_v198)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw86 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw86).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw86).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_601_v236 = _nw86
		var _604_v237 D22
		_ = _604_v237
		_604_v237 = Companion_D22_.Create_DC55_((_544_v198).F30(), _264_v1, (_555_v209).F43())
		var _605_v238 _dafny.Map
		_ = _605_v238
		_605_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v198.F29(), (_604_v237).Dtor_cf94())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_601_v236), 0))
		_ = _index50
		(_601_v236).ArraySet1(_605_v238, (_index50).Int())
		var _606_v239 T4
		_ = _606_v239
		var _nw87 *C8 = New_C8_()
		_ = _nw87
		_nw87.Ctor__(false, _551_v205, _dafny.IntOfInt64(694), !((_555_v209).F43()), _dafny.Companion_Sequence_.Update(_345_v66, (Companion_Default___.SafeIndex(_264_v1, _dafny.IntOfUint32((_345_v66).Cardinality()))).Uint32(), (_555_v209).F43()), _dafny.SeqOf(_345_v66), (_306_v36).F41(), (_555_v209).F43(), _544_v198.F31(), false)
		_606_v239 = _nw87
		var _607_v240 _dafny.Map
		_ = _607_v240
		_607_v240 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v198.F31(), _606_v239)
		var _608_v241 _dafny.Array
		_ = _608_v241
		var _nwElement0_14 _dafny.Map = (_607_v240).Update(_606_v239.F31(), _606_v239)
		_ = _nwElement0_14
		var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(6))
		_ = _nw88
		(_nw88).ArraySet1(_nwElement0_14, 0)
		(_nw88).ArraySet1((_607_v240).Merge(_607_v240), 1)
		(_nw88).ArraySet1((func() _dafny.Map {
			if _544_v198.F28() {
				return _607_v240
			}
			return _607_v240
		})(), 2)
		(_nw88).ArraySet1(_607_v240, 3)
		(_nw88).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_v239.F31(), _606_v239), 4)
		(_nw88).ArraySet1(_607_v240, 5)
		_608_v241 = _nw88
		var _609_v242 _dafny.Array
		_ = _609_v242
		var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw89
		_609_v242 = _nw89
		var _610_v243 _dafny.Map
		_ = _610_v243
		_610_v243 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v242, !(!(_606_v239.F28()) || (_606_v239.F35())))
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_601_v236), 0))
		_ = _index51
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
		_ = _index52
		var _rhs58 _dafny.Map = _605_v238
		_ = _rhs58
		var _rhs59 _dafny.Array = _608_v241
		_ = _rhs59
		var _rhs60 bool = ((_606_v239).F34()).Cmp((func() _dafny.Int {
			if _267_v4 {
				return (_547_v201).Cardinality()
			}
			return _dafny.IntOfInt64(797)
		})()) > 0
		_ = _rhs60
		var _rhs61 _dafny.Int = ((_306_v36).F41()).Plus((_263_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))).Int()).(_dafny.Int))
		_ = _rhs61
		var _rhs62 bool = (func() bool {
			if (_610_v243).Contains(_609_v242) {
				return (_610_v243).Get(_609_v242).(bool)
			}
			return _606_v239.F29()
		})()
		_ = _rhs62
		var _lhs46 _dafny.Array = _601_v236
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_601_v236), 0))
		_ = _lhs47
		var _lhs48 *GlobalState = _275_globalState
		_ = _lhs48
		var _lhs49 _dafny.Array = _263_v0
		_ = _lhs49
		var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
		_ = _lhs50
		var _lhs51 *GlobalState = _275_globalState
		_ = _lhs51
		(_lhs46).ArraySet1(_rhs58, (_lhs47).Int())
		_608_v241 = _rhs59
		_lhs48.F1 = _rhs60
		(_lhs49).ArraySet1(_rhs61, (_lhs50).Int())
		_lhs51.F12 = _rhs62
	} else {
		var _rhs63 _dafny.Array = _263_v0
		_ = _rhs63
		var _rhs64 bool = (_306_v36).Fm19(_275_globalState)
		_ = _rhs64
		var _lhs52 *GlobalState = _275_globalState
		_ = _lhs52
		_263_v0 = _rhs63
		_lhs52.F1 = _rhs64
		var _611_v244 _dafny.Set
		_ = _611_v244
		_611_v244 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), (_306_v36).F41())).Cardinality(), _dafny.IntOfUint32((_269_v6).Cardinality()))
		(_275_globalState).F1 = !(((_611_v244).Cardinality()).Cmp((_306_v36).F41()) < 0)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
		_ = _index53
		(_263_v0).ArraySet1(((_306_v36).F41()).Plus((_306_v36).F41()), (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))
		_ = _index54
		(_272_v9).ArraySet1(_267_v4, (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))
		_ = _index55
		(_272_v9).ArraySet1(!((func() bool {
			if _267_v4 {
				return !(_267_v4) || (_267_v4)
			}
			return _267_v4
		})()), (_index55).Int())
		if (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool) {
			(_275_globalState).F1 = (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _index56
			(_263_v0).ArraySet1(_dafny.IntOfInt64(501), (_index56).Int())
			var _612_v245 *C1
			_ = _612_v245
			var _nw90 *C1 = New_C1_()
			_ = _nw90
			_nw90.Ctor__((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), _267_v4)
			_612_v245 = _nw90
			_264_v1 = (_306_v36).F41()
			var _613_v246 D7
			_ = _613_v246
			_613_v246 = Companion_D7_.Create_DC20_(Companion_D7_.Create_DC19_((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), (_306_v36).F41(), _dafny.IntOfInt64(-249)))
			var _614_v247 _dafny.Map
			_ = _614_v247
			_614_v247 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_613_v246, _345_v66)
			var _615_v248 D7
			_ = _615_v248
			_615_v248 = Companion_D7_.Create_DC19_((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), (_306_v36).F41(), (_306_v36).F41())
			_614_v247 = (_614_v247).Update(Companion_D7_.Create_DC20_(_615_v248), _dafny.SeqOf(_267_v4))
		} else {
			var _616_v249 _dafny.Array
			_ = _616_v249
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_11
			var _nw91 _dafny.Array
			_ = _nw91
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw91 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_617_v9 _dafny.Array) func(_dafny.Int) bool {
					return func(_618_i23 _dafny.Int) bool {
						return (_617_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_617_v9), 0))).Int()).(bool)
					}
				})(_272_v9)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw91 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw91).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw91).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_616_v249 = _nw91
			var _619_v250 _dafny.CodePoint
			_ = _619_v250
			var _620_v251 _dafny.Sequence
			_ = _620_v251
			var _621_v252 _dafny.Int
			_ = _621_v252
			var _622_v253 _dafny.Sequence
			_ = _622_v253
			var _out44 _dafny.CodePoint
			_ = _out44
			var _out45 _dafny.Sequence
			_ = _out45
			var _out46 _dafny.Int
			_ = _out46
			var _out47 _dafny.Sequence
			_ = _out47
			_out44, _out45, _out46, _out47 = (_306_v36).M7(_267_v4, _616_v249, (_306_v36).Fm19(_275_globalState), _275_globalState)
			_619_v250 = _out44
			_620_v251 = _out45
			_621_v252 = _out46
			_622_v253 = _out47
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _index57
			var _rhs65 _dafny.Int = (_306_v36).F41()
			_ = _rhs65
			var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_269_v6, _622_v253)
			_ = _rhs66
			var _rhs67 _dafny.Sequence = _269_v6
			_ = _rhs67
			var _rhs68 _dafny.Int = Companion_Default___.SafeDivisionInt((_306_v36).F41(), Companion_Default___.Fm32((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), _275_globalState))
			_ = _rhs68
			var _lhs53 _dafny.Array = _263_v0
			_ = _lhs53
			var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
			_ = _lhs54
			(_lhs53).ArraySet1(_rhs65, (_lhs54).Int())
			_269_v6 = _rhs66
			_622_v253 = _rhs67
			_621_v252 = _rhs68
			var _623_v254 *C5
			_ = _623_v254
			var _nw92 *C5 = New_C5_()
			_ = _nw92
			_nw92.Ctor__((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), Companion_Default___.SafeModuloInt(Companion_Default___.Fm24(false, (_306_v36).F41(), _275_globalState), _621_v252), _264_v1, (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), _345_v66, _dafny.Companion_Sequence_.Concatenate(_414_v108, _414_v108), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool))).Cardinality()), !(!((_267_v4) || (_267_v4))), (_264_v1).Cmp((_306_v36).F41()) == 0, !((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool)) || (false))
			_623_v254 = _nw92
			var _624_v255 _dafny.Array
			_ = _624_v255
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
			_ = _nw93
			_624_v255 = _nw93
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_624_v255), 0))
			_ = _index58
			(_624_v255).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _267_v4 {
					return _310_v40
				}
				return _310_v40
			})(), (_index58).Int())
			var _625_v256 _dafny.Map
			_ = _625_v256
			_625_v256 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_415_v109, (_306_v36).F41())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.ArrayLen((_624_v255), 0))
			_ = _index59
			(_624_v255).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_dafny.IntOfInt64(99)).Cmp((_623_v254).F48()) > 0 {
					return _619_v250
				}
				return Companion_Default___.Fm48((_272_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_272_v9), 0))).Int()).(bool), true, (_625_v256).Cardinality(), _275_globalState)
			})(), (_index59).Int())
			(_275_globalState).F13 = _dafny.IntOfInt64(-574)
		}
	}
	var _626_v257 _dafny.MultiSet
	_ = _626_v257
	_626_v257 = _dafny.MultiSetOf(_272_v9, _272_v9)
	var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_263_v0), 0))
	_ = _index60
	(_263_v0).ArraySet1((Companion_Default___.SafeModuloInt((_626_v257).Cardinality(), (_306_v36).F41())).Plus(_dafny.IntOfInt64(329)), (_index60).Int())
	_dafny.Print((_263_v0).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_264_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_265_v2, _dafny.SeqOf(_dafny.IntOfInt64(20))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_266_v3).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_267_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_268_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(20)), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_269_v6.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_270_v7, _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m')), _dafny.UnicodeSeqOfUtf8Bytes("ssisbjr"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_272_v9).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_275_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_275_globalState).F5()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(20)), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F7().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_275_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_275_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_275_globalState).F18()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState.F21).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_275_globalState).F22()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_275_globalState).F25(), _dafny.SeqOf(_dafny.IntOfInt64(20))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_275_globalState.F26)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_globalState).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v36).F41())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_306_v36).F42()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(20))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_307_v37).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_308_v38).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_v39).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(20), _dafny.Zero).UpdateUnsafe(_dafny.Zero, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_310_v40)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_311_v41).Equals(_dafny.SetOf(_dafny.CodePoint('m'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_345_v66, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_346_v67).Dtor_cf46())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_346_v67).Dtor_cf47())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_346_v67).Dtor_cf48(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_346_v67).Dtor_cf49())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_346_v67).Dtor_cf50(), _dafny.SeqOf(_dafny.IntOfInt64(20))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_347_v68).Equals(_dafny.MultiSetOf(_dafny.CodePoint('m'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_348_v69).Dtor_cf15()).Equals(_dafny.MultiSetOf(_dafny.CodePoint('m'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_348_v69).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_348_v69).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v70).Equals(_dafny.SetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_350_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_414_v108, _dafny.SeqOf(_dafny.SeqOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_418_v110).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_543_v197, _dafny.SeqOf(Companion_D3_.Create_DC8_(_dafny.MultiSetOf(_dafny.CodePoint('m')), true, _dafny.IntOfInt64(417)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_626_v257).Cardinality())
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
	Cf1 _dafny.MultiSet
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.MultiSet) D0 {
	return D0{D0_DC1{Cf1}}
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
	return Companion_D0_.Create_DC1_(_dafny.EmptyMultiSet)
}

func (_this D0) Dtor_cf1() _dafny.MultiSet {
	return _this.Get_().(D0_DC1).Cf1
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf1.Equals(data2.Cf1)
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
	Cf3 bool
	Cf4 _dafny.Int
	Cf5 *C0
	Cf6 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 bool, Cf4 _dafny.Int, Cf5 *C0, Cf6 _dafny.Int) D1 {
	return D1{D1_DC3{Cf3, Cf4, Cf5, Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf2 _dafny.Array
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Array) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC4 struct {
	Cf7 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 D1) D1 {
	return D1{D1_DC4{Cf7}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false, _dafny.Zero, (*C0)(nil), _dafny.Zero)
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() *C0 {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf7() D1 {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2 == data2.Cf2
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf9  _dafny.Int
	Cf10 bool
	Cf11 _dafny.Int
	Cf12 bool
	Cf13 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf9 _dafny.Int, Cf10 bool, Cf11 _dafny.Int, Cf12 bool, Cf13 _dafny.Int) D2 {
	return D2{D2_DC6{Cf9, Cf10, Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf8 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf8 _dafny.Array) D2 {
	return D2{D2_DC5{Cf8}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.Zero, false, _dafny.Zero, false, _dafny.Zero)
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12 && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf8 == data2.Cf8
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
	Cf15 _dafny.MultiSet
	Cf16 bool
	Cf17 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf15 _dafny.MultiSet, Cf16 bool, Cf17 _dafny.Int) D3 {
	return D3{D3_DC8{Cf15, Cf16, Cf17}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf18 bool
	Cf19 bool
	Cf20 _dafny.Array
	Cf21 bool
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf18 bool, Cf19 bool, Cf20 _dafny.Array, Cf21 bool) D3 {
	return D3{D3_DC9{Cf18, Cf19, Cf20, Cf21}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC7 struct {
	Cf14 _dafny.Sequence
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf14 _dafny.Sequence) D3 {
	return D3{D3_DC7{Cf14}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.EmptyMultiSet, false, _dafny.Zero)
}

func (_this D3) Dtor_cf15() _dafny.MultiSet {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf17
}

func (_this D3) Dtor_cf18() bool {
	return _this.Get_().(D3_DC9).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC9).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D3_DC9).Cf20
}

func (_this D3) Dtor_cf21() bool {
	return _this.Get_().(D3_DC9).Cf21
}

func (_this D3) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D3_DC7).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf15.Equals(data2.Cf15) && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
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
	Cf23 _dafny.Int
	Cf24 bool
	Cf25 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf23 _dafny.Int, Cf24 bool, Cf25 _dafny.Int) D4 {
	return D4{D4_DC11{Cf23, Cf24, Cf25}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf22 _dafny.MultiSet
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf22 _dafny.MultiSet) D4 {
	return D4{D4_DC10{Cf22}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf23
}

func (_this D4) Dtor_cf24() bool {
	return _this.Get_().(D4_DC11).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf25
}

func (_this D4) Dtor_cf22() _dafny.MultiSet {
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
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

type D5_DC13 struct {
	Cf27 bool
	Cf28 _dafny.Map
	Cf29 _dafny.CodePoint
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf27 bool, Cf28 _dafny.Map, Cf29 _dafny.CodePoint) D5 {
	return D5{D5_DC13{Cf27, Cf28, Cf29}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf30 bool
	Cf31 _dafny.Int
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf30 bool, Cf31 _dafny.Int) D5 {
	return D5{D5_DC14{Cf30, Cf31}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf26 _dafny.Sequence
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf26 _dafny.Sequence) D5 {
	return D5{D5_DC12{Cf26}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(false, _dafny.EmptyMap, _dafny.CodePoint('D'))
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC13).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf28
}

func (_this D5) Dtor_cf29() _dafny.CodePoint {
	return _this.Get_().(D5_DC13).Cf29
}

func (_this D5) Dtor_cf30() bool {
	return _this.Get_().(D5_DC14).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf31
}

func (_this D5) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf26
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + data.Cf26.VerbatimString(true) + ")"
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
			return ok && data1.Cf27 == data2.Cf27 && data1.Cf28.Equals(data2.Cf28) && data1.Cf29 == data2.Cf29
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D6_DC16 struct {
	Cf33 _dafny.Array
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf33 _dafny.Array) D6 {
	return D6{D6_DC16{Cf33}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf32 _dafny.Sequence
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf32 _dafny.Sequence) D6 {
	return D6{D6_DC15{Cf32}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC17 struct {
	Cf34 D6
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf34 D6) D6 {
	return D6{D6_DC17{Cf34}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D6) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D6_DC16).Cf33
}

func (_this D6) Dtor_cf32() _dafny.Sequence {
	return _this.Get_().(D6_DC15).Cf32
}

func (_this D6) Dtor_cf34() D6 {
	return _this.Get_().(D6_DC17).Cf34
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf34) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf33 == data2.Cf33
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf32.Equals(data2.Cf32)
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
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

type D7_DC19 struct {
	Cf36 bool
	Cf37 _dafny.Int
	Cf38 _dafny.Int
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf36 bool, Cf37 _dafny.Int, Cf38 _dafny.Int) D7 {
	return D7{D7_DC19{Cf36, Cf37, Cf38}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC18 struct {
	Cf35 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf35 _dafny.Map) D7 {
	return D7{D7_DC18{Cf35}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC20 struct {
	Cf39 D7
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf39 D7) D7 {
	return D7{D7_DC20{Cf39}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D7) Dtor_cf36() bool {
	return _this.Get_().(D7_DC19).Cf36
}

func (_this D7) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf37
}

func (_this D7) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf38
}

func (_this D7) Dtor_cf35() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf35
}

func (_this D7) Dtor_cf39() D7 {
	return _this.Get_().(D7_DC20).Cf39
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf41 _dafny.Int
	Cf42 _dafny.Int
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf41 _dafny.Int, Cf42 _dafny.Int) D8 {
	return D8{D8_DC22{Cf41, Cf42}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC23 struct {
	Cf43 bool
	Cf44 bool
	Cf45 bool
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf43 bool, Cf44 bool, Cf45 bool) D8 {
	return D8{D8_DC23{Cf43, Cf44, Cf45}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf46 bool
	Cf47 _dafny.Int
	Cf48 _dafny.Sequence
	Cf49 bool
	Cf50 _dafny.Sequence
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf46 bool, Cf47 _dafny.Int, Cf48 _dafny.Sequence, Cf49 bool, Cf50 _dafny.Sequence) D8 {
	return D8{D8_DC24{Cf46, Cf47, Cf48, Cf49, Cf50}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC21 struct {
	Cf40 _dafny.Map
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf40 _dafny.Map) D8 {
	return D8{D8_DC21{Cf40}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC25 struct {
	Cf51 D8
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf51 D8) D8 {
	return D8{D8_DC25{Cf51}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC22_(_dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D8_DC22).Cf41
}

func (_this D8) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D8_DC22).Cf42
}

func (_this D8) Dtor_cf43() bool {
	return _this.Get_().(D8_DC23).Cf43
}

func (_this D8) Dtor_cf44() bool {
	return _this.Get_().(D8_DC23).Cf44
}

func (_this D8) Dtor_cf45() bool {
	return _this.Get_().(D8_DC23).Cf45
}

func (_this D8) Dtor_cf46() bool {
	return _this.Get_().(D8_DC24).Cf46
}

func (_this D8) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf47
}

func (_this D8) Dtor_cf48() _dafny.Sequence {
	return _this.Get_().(D8_DC24).Cf48
}

func (_this D8) Dtor_cf49() bool {
	return _this.Get_().(D8_DC24).Cf49
}

func (_this D8) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D8_DC24).Cf50
}

func (_this D8) Dtor_cf40() _dafny.Map {
	return _this.Get_().(D8_DC21).Cf40
}

func (_this D8) Dtor_cf51() D8 {
	return _this.Get_().(D8_DC25).Cf51
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf51) + ")"
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
			return ok && data1.Cf41.Cmp(data2.Cf41) == 0 && data1.Cf42.Cmp(data2.Cf42) == 0
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf43 == data2.Cf43 && data1.Cf44 == data2.Cf44 && data1.Cf45 == data2.Cf45
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Equals(data2.Cf48) && data1.Cf49 == data2.Cf49 && data1.Cf50.Equals(data2.Cf50)
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf40.Equals(data2.Cf40)
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D9_DC27 struct {
	Cf53 *C0
	Cf54 _dafny.Sequence
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf53 *C0, Cf54 _dafny.Sequence) D9 {
	return D9{D9_DC27{Cf53, Cf54}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC26 struct {
	Cf52 *C2
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf52 *C2) D9 {
	return D9{D9_DC26{Cf52}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC27_((*C0)(nil), _dafny.EmptySeq)
}

func (_this D9) Dtor_cf53() *C0 {
	return _this.Get_().(D9_DC27).Cf53
}

func (_this D9) Dtor_cf54() _dafny.Sequence {
	return _this.Get_().(D9_DC27).Cf54
}

func (_this D9) Dtor_cf52() *C2 {
	return _this.Get_().(D9_DC26).Cf52
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf53) + ", " + data.Cf54.VerbatimString(true) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf52) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf53 == data2.Cf53 && data1.Cf54.Equals(data2.Cf54)
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf52 == data2.Cf52
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
	Cf56 _dafny.Int
	Cf57 bool
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf56 _dafny.Int, Cf57 bool) D10 {
	return D10{D10_DC29{Cf56, Cf57}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC30 struct {
	Cf58 _dafny.CodePoint
	Cf59 _dafny.MultiSet
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf58 _dafny.CodePoint, Cf59 _dafny.MultiSet) D10 {
	return D10{D10_DC30{Cf58, Cf59}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

type D10_DC31 struct {
	Cf60 bool
	Cf61 bool
	Cf62 _dafny.Int
}

func (D10_DC31) isD10() {}

func (CompanionStruct_D10_) Create_DC31_(Cf60 bool, Cf61 bool, Cf62 _dafny.Int) D10 {
	return D10{D10_DC31{Cf60, Cf61, Cf62}}
}

func (_this D10) Is_DC31() bool {
	_, ok := _this.Get_().(D10_DC31)
	return ok
}

type D10_DC28 struct {
	Cf55 *C1
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf55 *C1) D10 {
	return D10{D10_DC28{Cf55}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC32 struct {
	Cf63 D10
}

func (D10_DC32) isD10() {}

func (CompanionStruct_D10_) Create_DC32_(Cf63 D10) D10 {
	return D10{D10_DC32{Cf63}}
}

func (_this D10) Is_DC32() bool {
	_, ok := _this.Get_().(D10_DC32)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC29_(_dafny.Zero, false)
}

func (_this D10) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D10_DC29).Cf56
}

func (_this D10) Dtor_cf57() bool {
	return _this.Get_().(D10_DC29).Cf57
}

func (_this D10) Dtor_cf58() _dafny.CodePoint {
	return _this.Get_().(D10_DC30).Cf58
}

func (_this D10) Dtor_cf59() _dafny.MultiSet {
	return _this.Get_().(D10_DC30).Cf59
}

func (_this D10) Dtor_cf60() bool {
	return _this.Get_().(D10_DC31).Cf60
}

func (_this D10) Dtor_cf61() bool {
	return _this.Get_().(D10_DC31).Cf61
}

func (_this D10) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D10_DC31).Cf62
}

func (_this D10) Dtor_cf55() *C1 {
	return _this.Get_().(D10_DC28).Cf55
}

func (_this D10) Dtor_cf63() D10 {
	return _this.Get_().(D10_DC32).Cf63
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ")"
		}
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D10_DC31:
		{
			return "D10.DC31" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D10_DC32:
		{
			return "D10.DC32" + "(" + _dafny.String(data.Cf63) + ")"
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
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0 && data1.Cf57 == data2.Cf57
		}
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59.Equals(data2.Cf59)
		}
	case D10_DC31:
		{
			data2, ok := other.Get_().(D10_DC31)
			return ok && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61 && data1.Cf62.Cmp(data2.Cf62) == 0
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf55 == data2.Cf55
		}
	case D10_DC32:
		{
			data2, ok := other.Get_().(D10_DC32)
			return ok && data1.Cf63.Equals(data2.Cf63)
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

type D11_DC33 struct {
	Cf64 _dafny.Set
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf64 _dafny.Set) D11 {
	return D11{D11_DC33{Cf64}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D11) Dtor_cf64() _dafny.Set {
	return _this.Get_().(D11_DC33).Cf64
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf64.Equals(data2.Cf64)
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

type D12_DC34 struct {
	Cf65 _dafny.Set
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf65 _dafny.Set) D12 {
	return D12{D12_DC34{Cf65}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D12) Dtor_cf65() _dafny.Set {
	return _this.Get_().(D12_DC34).Cf65
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D13_DC35 struct {
	Cf66 _dafny.Map
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf66 _dafny.Map) D13 {
	return D13{D13_DC35{Cf66}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D13) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D13_DC35).Cf66
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D14_DC37 struct {
	Cf68 _dafny.Int
	Cf69 bool
	Cf70 _dafny.Int
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf68 _dafny.Int, Cf69 bool, Cf70 _dafny.Int) D14 {
	return D14{D14_DC37{Cf68, Cf69, Cf70}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

type D14_DC36 struct {
	Cf67 _dafny.Sequence
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf67 _dafny.Sequence) D14 {
	return D14{D14_DC36{Cf67}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC37_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D14) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D14_DC37).Cf68
}

func (_this D14) Dtor_cf69() bool {
	return _this.Get_().(D14_DC37).Cf69
}

func (_this D14) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D14_DC37).Cf70
}

func (_this D14) Dtor_cf67() _dafny.Sequence {
	return _this.Get_().(D14_DC36).Cf67
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69 == data2.Cf69 && data1.Cf70.Cmp(data2.Cf70) == 0
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf67.Equals(data2.Cf67)
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
	Cf71 _dafny.Array
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf71 _dafny.Array) D15 {
	return D15{D15_DC38{Cf71}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

func (CompanionStruct_D15_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D15) Dtor_cf71() _dafny.Array {
	return _this.Get_().(D15_DC38).Cf71
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf71) + ")"
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
			return ok && data1.Cf71 == data2.Cf71
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

type D16_DC40 struct {
	Cf73 _dafny.Sequence
	Cf74 _dafny.Int
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf73 _dafny.Sequence, Cf74 _dafny.Int) D16 {
	return D16{D16_DC40{Cf73, Cf74}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

type D16_DC39 struct {
	Cf72 _dafny.Set
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf72 _dafny.Set) D16 {
	return D16{D16_DC39{Cf72}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

type D16_DC41 struct {
	Cf75 D16
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf75 D16) D16 {
	return D16{D16_DC41{Cf75}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC40_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D16) Dtor_cf73() _dafny.Sequence {
	return _this.Get_().(D16_DC40).Cf73
}

func (_this D16) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D16_DC40).Cf74
}

func (_this D16) Dtor_cf72() _dafny.Set {
	return _this.Get_().(D16_DC39).Cf72
}

func (_this D16) Dtor_cf75() D16 {
	return _this.Get_().(D16_DC41).Cf75
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC40:
		{
			return "D16.DC40" + "(" + data.Cf73.VerbatimString(true) + ", " + _dafny.String(data.Cf74) + ")"
		}
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf72) + ")"
		}
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC40:
		{
			data2, ok := other.Get_().(D16_DC40)
			return ok && data1.Cf73.Equals(data2.Cf73) && data1.Cf74.Cmp(data2.Cf74) == 0
		}
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf72.Equals(data2.Cf72)
		}
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf75.Equals(data2.Cf75)
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

type D17_DC43 struct {
}

func (D17_DC43) isD17() {}

func (CompanionStruct_D17_) Create_DC43_() D17 {
	return D17{D17_DC43{}}
}

func (_this D17) Is_DC43() bool {
	_, ok := _this.Get_().(D17_DC43)
	return ok
}

type D17_DC44 struct {
	Cf77 _dafny.MultiSet
	Cf78 _dafny.Int
}

func (D17_DC44) isD17() {}

func (CompanionStruct_D17_) Create_DC44_(Cf77 _dafny.MultiSet, Cf78 _dafny.Int) D17 {
	return D17{D17_DC44{Cf77, Cf78}}
}

func (_this D17) Is_DC44() bool {
	_, ok := _this.Get_().(D17_DC44)
	return ok
}

type D17_DC42 struct {
	Cf76 _dafny.Sequence
}

func (D17_DC42) isD17() {}

func (CompanionStruct_D17_) Create_DC42_(Cf76 _dafny.Sequence) D17 {
	return D17{D17_DC42{Cf76}}
}

func (_this D17) Is_DC42() bool {
	_, ok := _this.Get_().(D17_DC42)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC43_()
}

func (_this D17) Dtor_cf77() _dafny.MultiSet {
	return _this.Get_().(D17_DC44).Cf77
}

func (_this D17) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D17_DC44).Cf78
}

func (_this D17) Dtor_cf76() _dafny.Sequence {
	return _this.Get_().(D17_DC42).Cf76
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC43:
		{
			return "D17.DC43"
		}
	case D17_DC44:
		{
			return "D17.DC44" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ")"
		}
	case D17_DC42:
		{
			return "D17.DC42" + "(" + _dafny.String(data.Cf76) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC43:
		{
			_, ok := other.Get_().(D17_DC43)
			return ok
		}
	case D17_DC44:
		{
			data2, ok := other.Get_().(D17_DC44)
			return ok && data1.Cf77.Equals(data2.Cf77) && data1.Cf78.Cmp(data2.Cf78) == 0
		}
	case D17_DC42:
		{
			data2, ok := other.Get_().(D17_DC42)
			return ok && data1.Cf76.Equals(data2.Cf76)
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

type D18_DC45 struct {
	Cf79 _dafny.Sequence
}

func (D18_DC45) isD18() {}

func (CompanionStruct_D18_) Create_DC45_(Cf79 _dafny.Sequence) D18 {
	return D18{D18_DC45{Cf79}}
}

func (_this D18) Is_DC45() bool {
	_, ok := _this.Get_().(D18_DC45)
	return ok
}

func (CompanionStruct_D18_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D18) Dtor_cf79() _dafny.Sequence {
	return _this.Get_().(D18_DC45).Cf79
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC45:
		{
			return "D18.DC45" + "(" + _dafny.String(data.Cf79) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC45:
		{
			data2, ok := other.Get_().(D18_DC45)
			return ok && data1.Cf79.Equals(data2.Cf79)
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
	Cf80 _dafny.MultiSet
}

func (D19_DC46) isD19() {}

func (CompanionStruct_D19_) Create_DC46_(Cf80 _dafny.MultiSet) D19 {
	return D19{D19_DC46{Cf80}}
}

func (_this D19) Is_DC46() bool {
	_, ok := _this.Get_().(D19_DC46)
	return ok
}

func (CompanionStruct_D19_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D19) Dtor_cf80() _dafny.MultiSet {
	return _this.Get_().(D19_DC46).Cf80
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC46:
		{
			return "D19.DC46" + "(" + _dafny.String(data.Cf80) + ")"
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
			return ok && data1.Cf80.Equals(data2.Cf80)
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

type D20_DC48 struct {
	Cf82 _dafny.Int
	Cf83 _dafny.Array
	Cf84 bool
}

func (D20_DC48) isD20() {}

func (CompanionStruct_D20_) Create_DC48_(Cf82 _dafny.Int, Cf83 _dafny.Array, Cf84 bool) D20 {
	return D20{D20_DC48{Cf82, Cf83, Cf84}}
}

func (_this D20) Is_DC48() bool {
	_, ok := _this.Get_().(D20_DC48)
	return ok
}

type D20_DC49 struct {
	Cf85 bool
	Cf86 _dafny.CodePoint
}

func (D20_DC49) isD20() {}

func (CompanionStruct_D20_) Create_DC49_(Cf85 bool, Cf86 _dafny.CodePoint) D20 {
	return D20{D20_DC49{Cf85, Cf86}}
}

func (_this D20) Is_DC49() bool {
	_, ok := _this.Get_().(D20_DC49)
	return ok
}

type D20_DC47 struct {
	Cf81 _dafny.Array
}

func (D20_DC47) isD20() {}

func (CompanionStruct_D20_) Create_DC47_(Cf81 _dafny.Array) D20 {
	return D20{D20_DC47{Cf81}}
}

func (_this D20) Is_DC47() bool {
	_, ok := _this.Get_().(D20_DC47)
	return ok
}

type D20_DC50 struct {
	Cf87 D20
}

func (D20_DC50) isD20() {}

func (CompanionStruct_D20_) Create_DC50_(Cf87 D20) D20 {
	return D20{D20_DC50{Cf87}}
}

func (_this D20) Is_DC50() bool {
	_, ok := _this.Get_().(D20_DC50)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC48_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D20) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D20_DC48).Cf82
}

func (_this D20) Dtor_cf83() _dafny.Array {
	return _this.Get_().(D20_DC48).Cf83
}

func (_this D20) Dtor_cf84() bool {
	return _this.Get_().(D20_DC48).Cf84
}

func (_this D20) Dtor_cf85() bool {
	return _this.Get_().(D20_DC49).Cf85
}

func (_this D20) Dtor_cf86() _dafny.CodePoint {
	return _this.Get_().(D20_DC49).Cf86
}

func (_this D20) Dtor_cf81() _dafny.Array {
	return _this.Get_().(D20_DC47).Cf81
}

func (_this D20) Dtor_cf87() D20 {
	return _this.Get_().(D20_DC50).Cf87
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC48:
		{
			return "D20.DC48" + "(" + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ")"
		}
	case D20_DC49:
		{
			return "D20.DC49" + "(" + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ")"
		}
	case D20_DC47:
		{
			return "D20.DC47" + "(" + _dafny.String(data.Cf81) + ")"
		}
	case D20_DC50:
		{
			return "D20.DC50" + "(" + _dafny.String(data.Cf87) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC48:
		{
			data2, ok := other.Get_().(D20_DC48)
			return ok && data1.Cf82.Cmp(data2.Cf82) == 0 && data1.Cf83 == data2.Cf83 && data1.Cf84 == data2.Cf84
		}
	case D20_DC49:
		{
			data2, ok := other.Get_().(D20_DC49)
			return ok && data1.Cf85 == data2.Cf85 && data1.Cf86 == data2.Cf86
		}
	case D20_DC47:
		{
			data2, ok := other.Get_().(D20_DC47)
			return ok && data1.Cf81 == data2.Cf81
		}
	case D20_DC50:
		{
			data2, ok := other.Get_().(D20_DC50)
			return ok && data1.Cf87.Equals(data2.Cf87)
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

type D21_DC52 struct {
	Cf89 _dafny.CodePoint
}

func (D21_DC52) isD21() {}

func (CompanionStruct_D21_) Create_DC52_(Cf89 _dafny.CodePoint) D21 {
	return D21{D21_DC52{Cf89}}
}

func (_this D21) Is_DC52() bool {
	_, ok := _this.Get_().(D21_DC52)
	return ok
}

type D21_DC51 struct {
	Cf88 _dafny.Map
}

func (D21_DC51) isD21() {}

func (CompanionStruct_D21_) Create_DC51_(Cf88 _dafny.Map) D21 {
	return D21{D21_DC51{Cf88}}
}

func (_this D21) Is_DC51() bool {
	_, ok := _this.Get_().(D21_DC51)
	return ok
}

type D21_DC53 struct {
	Cf90 D21
}

func (D21_DC53) isD21() {}

func (CompanionStruct_D21_) Create_DC53_(Cf90 D21) D21 {
	return D21{D21_DC53{Cf90}}
}

func (_this D21) Is_DC53() bool {
	_, ok := _this.Get_().(D21_DC53)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC52_(_dafny.CodePoint('D'))
}

func (_this D21) Dtor_cf89() _dafny.CodePoint {
	return _this.Get_().(D21_DC52).Cf89
}

func (_this D21) Dtor_cf88() _dafny.Map {
	return _this.Get_().(D21_DC51).Cf88
}

func (_this D21) Dtor_cf90() D21 {
	return _this.Get_().(D21_DC53).Cf90
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC52:
		{
			return "D21.DC52" + "(" + _dafny.String(data.Cf89) + ")"
		}
	case D21_DC51:
		{
			return "D21.DC51" + "(" + _dafny.String(data.Cf88) + ")"
		}
	case D21_DC53:
		{
			return "D21.DC53" + "(" + _dafny.String(data.Cf90) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC52:
		{
			data2, ok := other.Get_().(D21_DC52)
			return ok && data1.Cf89 == data2.Cf89
		}
	case D21_DC51:
		{
			data2, ok := other.Get_().(D21_DC51)
			return ok && data1.Cf88.Equals(data2.Cf88)
		}
	case D21_DC53:
		{
			data2, ok := other.Get_().(D21_DC53)
			return ok && data1.Cf90.Equals(data2.Cf90)
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

type D22_DC55 struct {
	Cf92 _dafny.Int
	Cf93 _dafny.Int
	Cf94 bool
}

func (D22_DC55) isD22() {}

func (CompanionStruct_D22_) Create_DC55_(Cf92 _dafny.Int, Cf93 _dafny.Int, Cf94 bool) D22 {
	return D22{D22_DC55{Cf92, Cf93, Cf94}}
}

func (_this D22) Is_DC55() bool {
	_, ok := _this.Get_().(D22_DC55)
	return ok
}

type D22_DC54 struct {
	Cf91 _dafny.Array
}

func (D22_DC54) isD22() {}

func (CompanionStruct_D22_) Create_DC54_(Cf91 _dafny.Array) D22 {
	return D22{D22_DC54{Cf91}}
}

func (_this D22) Is_DC54() bool {
	_, ok := _this.Get_().(D22_DC54)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC55_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D22) Dtor_cf92() _dafny.Int {
	return _this.Get_().(D22_DC55).Cf92
}

func (_this D22) Dtor_cf93() _dafny.Int {
	return _this.Get_().(D22_DC55).Cf93
}

func (_this D22) Dtor_cf94() bool {
	return _this.Get_().(D22_DC55).Cf94
}

func (_this D22) Dtor_cf91() _dafny.Array {
	return _this.Get_().(D22_DC54).Cf91
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC55:
		{
			return "D22.DC55" + "(" + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ", " + _dafny.String(data.Cf94) + ")"
		}
	case D22_DC54:
		{
			return "D22.DC54" + "(" + _dafny.String(data.Cf91) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC55:
		{
			data2, ok := other.Get_().(D22_DC55)
			return ok && data1.Cf92.Cmp(data2.Cf92) == 0 && data1.Cf93.Cmp(data2.Cf93) == 0 && data1.Cf94 == data2.Cf94
		}
	case D22_DC54:
		{
			data2, ok := other.Get_().(D22_DC54)
			return ok && data1.Cf91 == data2.Cf91
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

type D23_DC56 struct {
	Cf95 _dafny.Array
}

func (D23_DC56) isD23() {}

func (CompanionStruct_D23_) Create_DC56_(Cf95 _dafny.Array) D23 {
	return D23{D23_DC56{Cf95}}
}

func (_this D23) Is_DC56() bool {
	_, ok := _this.Get_().(D23_DC56)
	return ok
}

func (CompanionStruct_D23_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D23) Dtor_cf95() _dafny.Array {
	return _this.Get_().(D23_DC56).Cf95
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC56:
		{
			return "D23.DC56" + "(" + _dafny.String(data.Cf95) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC56:
		{
			data2, ok := other.Get_().(D23_DC56)
			return ok && data1.Cf95 == data2.Cf95
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

type D24_DC58 struct {
}

func (D24_DC58) isD24() {}

func (CompanionStruct_D24_) Create_DC58_() D24 {
	return D24{D24_DC58{}}
}

func (_this D24) Is_DC58() bool {
	_, ok := _this.Get_().(D24_DC58)
	return ok
}

type D24_DC59 struct {
	Cf97 bool
}

func (D24_DC59) isD24() {}

func (CompanionStruct_D24_) Create_DC59_(Cf97 bool) D24 {
	return D24{D24_DC59{Cf97}}
}

func (_this D24) Is_DC59() bool {
	_, ok := _this.Get_().(D24_DC59)
	return ok
}

type D24_DC57 struct {
	Cf96 _dafny.Set
}

func (D24_DC57) isD24() {}

func (CompanionStruct_D24_) Create_DC57_(Cf96 _dafny.Set) D24 {
	return D24{D24_DC57{Cf96}}
}

func (_this D24) Is_DC57() bool {
	_, ok := _this.Get_().(D24_DC57)
	return ok
}

type D24_DC60 struct {
	Cf98 D24
}

func (D24_DC60) isD24() {}

func (CompanionStruct_D24_) Create_DC60_(Cf98 D24) D24 {
	return D24{D24_DC60{Cf98}}
}

func (_this D24) Is_DC60() bool {
	_, ok := _this.Get_().(D24_DC60)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC58_()
}

func (_this D24) Dtor_cf97() bool {
	return _this.Get_().(D24_DC59).Cf97
}

func (_this D24) Dtor_cf96() _dafny.Set {
	return _this.Get_().(D24_DC57).Cf96
}

func (_this D24) Dtor_cf98() D24 {
	return _this.Get_().(D24_DC60).Cf98
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC58:
		{
			return "D24.DC58"
		}
	case D24_DC59:
		{
			return "D24.DC59" + "(" + _dafny.String(data.Cf97) + ")"
		}
	case D24_DC57:
		{
			return "D24.DC57" + "(" + _dafny.String(data.Cf96) + ")"
		}
	case D24_DC60:
		{
			return "D24.DC60" + "(" + _dafny.String(data.Cf98) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC58:
		{
			_, ok := other.Get_().(D24_DC58)
			return ok
		}
	case D24_DC59:
		{
			data2, ok := other.Get_().(D24_DC59)
			return ok && data1.Cf97 == data2.Cf97
		}
	case D24_DC57:
		{
			data2, ok := other.Get_().(D24_DC57)
			return ok && data1.Cf96.Equals(data2.Cf96)
		}
	case D24_DC60:
		{
			data2, ok := other.Get_().(D24_DC60)
			return ok && data1.Cf98.Equals(data2.Cf98)
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

type D25_DC61 struct {
	Cf99 _dafny.Sequence
}

func (D25_DC61) isD25() {}

func (CompanionStruct_D25_) Create_DC61_(Cf99 _dafny.Sequence) D25 {
	return D25{D25_DC61{Cf99}}
}

func (_this D25) Is_DC61() bool {
	_, ok := _this.Get_().(D25_DC61)
	return ok
}

func (CompanionStruct_D25_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D25) Dtor_cf99() _dafny.Sequence {
	return _this.Get_().(D25_DC61).Cf99
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC61:
		{
			return "D25.DC61" + "(" + _dafny.String(data.Cf99) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC61:
		{
			data2, ok := other.Get_().(D25_DC61)
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

type D26_DC63 struct {
	Cf101 _dafny.Int
	Cf102 _dafny.Int
	Cf103 bool
}

func (D26_DC63) isD26() {}

func (CompanionStruct_D26_) Create_DC63_(Cf101 _dafny.Int, Cf102 _dafny.Int, Cf103 bool) D26 {
	return D26{D26_DC63{Cf101, Cf102, Cf103}}
}

func (_this D26) Is_DC63() bool {
	_, ok := _this.Get_().(D26_DC63)
	return ok
}

type D26_DC62 struct {
	Cf100 _dafny.Map
}

func (D26_DC62) isD26() {}

func (CompanionStruct_D26_) Create_DC62_(Cf100 _dafny.Map) D26 {
	return D26{D26_DC62{Cf100}}
}

func (_this D26) Is_DC62() bool {
	_, ok := _this.Get_().(D26_DC62)
	return ok
}

type D26_DC64 struct {
	Cf104 D26
}

func (D26_DC64) isD26() {}

func (CompanionStruct_D26_) Create_DC64_(Cf104 D26) D26 {
	return D26{D26_DC64{Cf104}}
}

func (_this D26) Is_DC64() bool {
	_, ok := _this.Get_().(D26_DC64)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC63_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D26) Dtor_cf101() _dafny.Int {
	return _this.Get_().(D26_DC63).Cf101
}

func (_this D26) Dtor_cf102() _dafny.Int {
	return _this.Get_().(D26_DC63).Cf102
}

func (_this D26) Dtor_cf103() bool {
	return _this.Get_().(D26_DC63).Cf103
}

func (_this D26) Dtor_cf100() _dafny.Map {
	return _this.Get_().(D26_DC62).Cf100
}

func (_this D26) Dtor_cf104() D26 {
	return _this.Get_().(D26_DC64).Cf104
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC63:
		{
			return "D26.DC63" + "(" + _dafny.String(data.Cf101) + ", " + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D26_DC62:
		{
			return "D26.DC62" + "(" + _dafny.String(data.Cf100) + ")"
		}
	case D26_DC64:
		{
			return "D26.DC64" + "(" + _dafny.String(data.Cf104) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC63:
		{
			data2, ok := other.Get_().(D26_DC63)
			return ok && data1.Cf101.Cmp(data2.Cf101) == 0 && data1.Cf102.Cmp(data2.Cf102) == 0 && data1.Cf103 == data2.Cf103
		}
	case D26_DC62:
		{
			data2, ok := other.Get_().(D26_DC62)
			return ok && data1.Cf100.Equals(data2.Cf100)
		}
	case D26_DC64:
		{
			data2, ok := other.Get_().(D26_DC64)
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

type D27_DC65 struct {
	Cf105 _dafny.Map
}

func (D27_DC65) isD27() {}

func (CompanionStruct_D27_) Create_DC65_(Cf105 _dafny.Map) D27 {
	return D27{D27_DC65{Cf105}}
}

func (_this D27) Is_DC65() bool {
	_, ok := _this.Get_().(D27_DC65)
	return ok
}

func (CompanionStruct_D27_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D27) Dtor_cf105() _dafny.Map {
	return _this.Get_().(D27_DC65).Cf105
}

func (_this D27) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D27_DC65:
		{
			return "D27.DC65" + "(" + _dafny.String(data.Cf105) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D27) Equals(other D27) bool {
	switch data1 := _this.Get_().(type) {
	case D27_DC65:
		{
			data2, ok := other.Get_().(D27_DC65)
			return ok && data1.Cf105.Equals(data2.Cf105)
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

type D28_DC66 struct {
	Cf106 _dafny.Map
}

func (D28_DC66) isD28() {}

func (CompanionStruct_D28_) Create_DC66_(Cf106 _dafny.Map) D28 {
	return D28{D28_DC66{Cf106}}
}

func (_this D28) Is_DC66() bool {
	_, ok := _this.Get_().(D28_DC66)
	return ok
}

func (CompanionStruct_D28_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D28) Dtor_cf106() _dafny.Map {
	return _this.Get_().(D28_DC66).Cf106
}

func (_this D28) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D28_DC66:
		{
			return "D28.DC66" + "(" + _dafny.String(data.Cf106) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D28) Equals(other D28) bool {
	switch data1 := _this.Get_().(type) {
	case D28_DC66:
		{
			data2, ok := other.Get_().(D28_DC66)
			return ok && data1.Cf106.Equals(data2.Cf106)
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

type D29_DC68 struct {
	Cf108 _dafny.Sequence
	Cf109 _dafny.Int
	Cf110 bool
	Cf111 bool
}

func (D29_DC68) isD29() {}

func (CompanionStruct_D29_) Create_DC68_(Cf108 _dafny.Sequence, Cf109 _dafny.Int, Cf110 bool, Cf111 bool) D29 {
	return D29{D29_DC68{Cf108, Cf109, Cf110, Cf111}}
}

func (_this D29) Is_DC68() bool {
	_, ok := _this.Get_().(D29_DC68)
	return ok
}

type D29_DC67 struct {
	Cf107 _dafny.Map
}

func (D29_DC67) isD29() {}

func (CompanionStruct_D29_) Create_DC67_(Cf107 _dafny.Map) D29 {
	return D29{D29_DC67{Cf107}}
}

func (_this D29) Is_DC67() bool {
	_, ok := _this.Get_().(D29_DC67)
	return ok
}

type D29_DC69 struct {
	Cf112 D29
}

func (D29_DC69) isD29() {}

func (CompanionStruct_D29_) Create_DC69_(Cf112 D29) D29 {
	return D29{D29_DC69{Cf112}}
}

func (_this D29) Is_DC69() bool {
	_, ok := _this.Get_().(D29_DC69)
	return ok
}

func (CompanionStruct_D29_) Default() D29 {
	return Companion_D29_.Create_DC68_(_dafny.EmptySeq, _dafny.Zero, false, false)
}

func (_this D29) Dtor_cf108() _dafny.Sequence {
	return _this.Get_().(D29_DC68).Cf108
}

func (_this D29) Dtor_cf109() _dafny.Int {
	return _this.Get_().(D29_DC68).Cf109
}

func (_this D29) Dtor_cf110() bool {
	return _this.Get_().(D29_DC68).Cf110
}

func (_this D29) Dtor_cf111() bool {
	return _this.Get_().(D29_DC68).Cf111
}

func (_this D29) Dtor_cf107() _dafny.Map {
	return _this.Get_().(D29_DC67).Cf107
}

func (_this D29) Dtor_cf112() D29 {
	return _this.Get_().(D29_DC69).Cf112
}

func (_this D29) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D29_DC68:
		{
			return "D29.DC68" + "(" + _dafny.String(data.Cf108) + ", " + _dafny.String(data.Cf109) + ", " + _dafny.String(data.Cf110) + ", " + _dafny.String(data.Cf111) + ")"
		}
	case D29_DC67:
		{
			return "D29.DC67" + "(" + _dafny.String(data.Cf107) + ")"
		}
	case D29_DC69:
		{
			return "D29.DC69" + "(" + _dafny.String(data.Cf112) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D29) Equals(other D29) bool {
	switch data1 := _this.Get_().(type) {
	case D29_DC68:
		{
			data2, ok := other.Get_().(D29_DC68)
			return ok && data1.Cf108.Equals(data2.Cf108) && data1.Cf109.Cmp(data2.Cf109) == 0 && data1.Cf110 == data2.Cf110 && data1.Cf111 == data2.Cf111
		}
	case D29_DC67:
		{
			data2, ok := other.Get_().(D29_DC67)
			return ok && data1.Cf107.Equals(data2.Cf107)
		}
	case D29_DC69:
		{
			data2, ok := other.Get_().(D29_DC69)
			return ok && data1.Cf112.Equals(data2.Cf112)
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

// Definition of trait T0
type T0 interface {
	String() string
	F28() bool
	F28_set_(value bool)
	F29() bool
	F29_set_(value bool)
	Fm2(globalState *GlobalState) _dafny.Map
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
	F28() bool
	F28_set_(value bool)
	F29() bool
	F29_set_(value bool)
	Fm2(globalState *GlobalState) _dafny.Map
	F31() bool
	F31_set_(value bool)
	Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence
	F30() _dafny.Int
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
	F28() bool
	F28_set_(value bool)
	F29() bool
	F29_set_(value bool)
	F31() bool
	F31_set_(value bool)
	Fm2(globalState *GlobalState) _dafny.Map
	Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence
	F30() _dafny.Int
	Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set
	Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence)
	F32() _dafny.Sequence
	F33() _dafny.Sequence
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

// Definition of trait T3
type T3 interface {
	String() string
	F28() bool
	F28_set_(value bool)
	F29() bool
	F29_set_(value bool)
	F31() bool
	F31_set_(value bool)
	Fm2(globalState *GlobalState) _dafny.Map
	Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence
	Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set
	Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence)
	F32() _dafny.Sequence
	F33() _dafny.Sequence
	F30() _dafny.Int
	F35() bool
	F35_set_(value bool)
	Fm6(globalState *GlobalState) bool
	F34() _dafny.Int
}
type CompanionStruct_T3_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T3_ = CompanionStruct_T3_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T3_) CastTo_(x interface{}) T3 {
	var t T3
	t, _ = x.(T3)
	return t
}

// End of trait T3

// Definition of trait T4
type T4 interface {
	String() string
	F28() bool
	F28_set_(value bool)
	F29() bool
	F29_set_(value bool)
	F31() bool
	F31_set_(value bool)
	F35() bool
	F35_set_(value bool)
	Fm2(globalState *GlobalState) _dafny.Map
	Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence
	Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set
	Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	Fm6(globalState *GlobalState) bool
	M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence)
	F34() _dafny.Int
	F32() _dafny.Sequence
	F33() _dafny.Sequence
	F30() _dafny.Int
	F36() _dafny.CodePoint
	F36_set_(value _dafny.CodePoint)
	Fm7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool
}
type CompanionStruct_T4_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T4_ = CompanionStruct_T4_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T4_) CastTo_(x interface{}) T4 {
	var t T4
	t, _ = x.(T4)
	return t
}

// End of trait T4

// Definition of class GlobalState
type GlobalState struct {
	F1   bool
	F12  bool
	F13  _dafny.Int
	F21  _dafny.Array
	F26  bool
	_f0  _dafny.Int
	_f2  bool
	_f3  bool
	_f4  _dafny.Int
	_f5  _dafny.Map
	_f6  _dafny.Int
	_f7  _dafny.Sequence
	_f8  _dafny.Sequence
	_f9  bool
	_f10 _dafny.Int
	_f11 _dafny.Int
	_f14 bool
	_f15 _dafny.Array
	_f16 _dafny.Int
	_f17 _dafny.CodePoint
	_f18 _dafny.Array
	_f19 bool
	_f20 bool
	_f22 _dafny.Array
	_f23 bool
	_f24 _dafny.CodePoint
	_f25 _dafny.Sequence
	_f27 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = false
	_this.F12 = false
	_this.F13 = _dafny.Zero
	_this.F21 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F26 = false
	_this._f0 = _dafny.Zero
	_this._f2 = false
	_this._f3 = false
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.EmptyMap
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.EmptySeq
	_this._f8 = _dafny.EmptySeq
	_this._f9 = false
	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.Zero
	_this._f14 = false
	_this._f15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.CodePoint('D')
	_this._f18 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f19 = false
	_this._f20 = false
	_this._f22 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f23 = false
	_this._f24 = _dafny.CodePoint('D')
	_this._f25 = _dafny.EmptySeq
	_this._f27 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 bool, f3 bool, f4 _dafny.Int, f5 _dafny.Map, f6 _dafny.Int, f7 _dafny.Sequence, f8 _dafny.Sequence, f9 bool, f10 _dafny.Int, f11 _dafny.Int, f12 bool, f13 _dafny.Int, f14 bool, f15 _dafny.Array, f16 _dafny.Int, f17 _dafny.CodePoint, f18 _dafny.Array, f19 bool, f20 bool, f21 _dafny.Array, f22 _dafny.Array, f23 bool, f24 _dafny.CodePoint, f25 _dafny.Sequence, f26 bool, f27 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this).F21 = f21
		(_this)._f22 = f22
		(_this)._f23 = f23
		(_this)._f24 = f24
		(_this)._f25 = f25
		(_this).F26 = f26
		(_this)._f27 = f27
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Map {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Sequence {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Array {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.CodePoint {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Array {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() bool {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() bool {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F22() _dafny.Array {
	{
		return _this._f22
	}
}
func (_this *GlobalState) F23() bool {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F24() _dafny.CodePoint {
	{
		return _this._f24
	}
}
func (_this *GlobalState) F25() _dafny.Sequence {
	{
		return _this._f25
	}
}
func (_this *GlobalState) F27() _dafny.Int {
	{
		return _this._f27
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F37 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F37 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f37 _dafny.Array) {
	{
		(_this).F37 = f37
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f28 bool
	_f29 bool
	_f43 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f28 = false
	_this._f29 = false
	_this._f43 = false
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

func (_this *C1) F28() bool {
	return _this._f28
}
func (_this *C1) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C1) F29() bool {
	return _this._f29
}
func (_this *C1) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C1) Ctor__(f43 bool, f28 bool, f29 bool) {
	{
		(_this)._f43 = f43
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C1) Fm2(globalState *GlobalState) _dafny.Map {
	{
		var _source11 D3 = Companion_D3_.Create_DC8_(_dafny.MultiSetOf(_dafny.CodePoint('s'), _dafny.CodePoint('e'), _dafny.CodePoint('y'), _dafny.CodePoint('y'), _dafny.CodePoint('a')), false, _dafny.IntOfInt64(611))
		_ = _source11
		if _source11.Is_DC8() {
			var _627___mcc_h0 _dafny.MultiSet = _source11.Get_().(D3_DC8).Cf15
			_ = _627___mcc_h0
			var _628___mcc_h1 bool = _source11.Get_().(D3_DC8).Cf16
			_ = _628___mcc_h1
			var _629___mcc_h2 _dafny.Int = _source11.Get_().(D3_DC8).Cf17
			_ = _629___mcc_h2
			var _630_cf17 _dafny.Int = _629___mcc_h2
			_ = _630_cf17
			var _631_cf16 bool = _628___mcc_h1
			_ = _631_cf16
			var _632_cf15 _dafny.MultiSet = _627___mcc_h0
			_ = _632_cf15
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F43(), _630_cf17)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F43(), (_dafny.Zero).Minus(_630_cf17)))
		} else if _source11.Is_DC9() {
			var _633___mcc_h3 bool = _source11.Get_().(D3_DC9).Cf18
			_ = _633___mcc_h3
			var _634___mcc_h4 bool = _source11.Get_().(D3_DC9).Cf19
			_ = _634___mcc_h4
			var _635___mcc_h5 _dafny.Array = _source11.Get_().(D3_DC9).Cf20
			_ = _635___mcc_h5
			var _636___mcc_h6 bool = _source11.Get_().(D3_DC9).Cf21
			_ = _636___mcc_h6
			var _637_cf21 bool = _636___mcc_h6
			_ = _637_cf21
			var _638_cf20 _dafny.Array = _635___mcc_h5
			_ = _638_cf20
			var _639_cf19 bool = _634___mcc_h4
			_ = _639_cf19
			var _640_cf18 bool = _633___mcc_h3
			_ = _640_cf18
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F43(), _dafny.IntOfInt64(513))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(615)))
		} else {
			var _641___mcc_h7 _dafny.Sequence = _source11.Get_().(D3_DC7).Cf14
			_ = _641___mcc_h7
			var _642_cf14 _dafny.Sequence = _641___mcc_h7
			_ = _642_cf14
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC6_(_dafny.IntOfInt64(643), _this.F28(), _dafny.IntOfInt64(167), (_this).F43(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-627))).Cardinality()))).Dtor_cf12(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('h'), _dafny.CodePoint('m'))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F43(), _dafny.IntOfInt64(342)))
		}
	}
}
func (_this *C1) Fm22(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C1) M9(globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, _dafny.CodePoint, bool) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r2
		var r3 bool = false
		_ = r3
		var _643_v0 _dafny.Int
		_ = _643_v0
		_643_v0 = _dafny.IntOfInt64(690)
		(_this).F28_set_(!(Companion_Default___.Fm1(_this.F28(), _643_v0, globalState)))
		var _644_v1 _dafny.Map
		_ = _644_v1
		_644_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_643_v0, _643_v0)
		var _645_v2 _dafny.Map
		_ = _645_v2
		_645_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (func() _dafny.Int {
			if (_644_v1).Contains(_dafny.IntOfInt64(193)) {
				return (_644_v1).Get(_dafny.IntOfInt64(193)).(_dafny.Int)
			}
			return (_dafny.SetOf(_643_v0)).Cardinality()
		})())
		var _646_v3 _dafny.MultiSet
		_ = _646_v3
		_646_v3 = _dafny.MultiSetOf((_645_v2).Cardinality())
		r3 = (_646_v3).Equals((_646_v3).Union(_646_v3))
		_644_v1 = (_644_v1).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_643_v0, _dafny.IntOfInt64(854)))
		var _647_v4 _dafny.Map
		_ = _647_v4
		_647_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(410), _this.F28())
		var _648_v5 _dafny.Sequence
		_ = _648_v5
		_648_v5 = _dafny.SeqOf(_643_v0)
		_643_v0 = ((_dafny.Zero).Minus(((_647_v4).Update(_643_v0, _this.F28())).Cardinality())).Times((_dafny.IntOfInt64(545)).Minus(_dafny.IntOfUint32((_648_v5).Cardinality())))
		var _649_v6 _dafny.Sequence
		_ = _649_v6
		_649_v6 = _dafny.UnicodeSeqOfUtf8Bytes("avexbc")
		var _650_v7 _dafny.MultiSet
		_ = _650_v7
		_650_v7 = _dafny.MultiSetOf((_this).F43(), _this.F29(), ((_this).F43()) == ((_this).F43()), Companion_Default___.Fm1((_this).F43(), _643_v0, globalState))
		var _651_v8 _dafny.Set
		_ = _651_v8
		_651_v8 = _dafny.SetOf(false)
		var _652_v9 _dafny.Sequence
		_ = _652_v9
		_652_v9 = _dafny.SeqOf(_645_v2, _645_v2)
		var _653_v10 D3
		_ = _653_v10
		_653_v10 = Companion_D3_.Create_DC7_(_652_v9)
		var _654_v11 _dafny.Sequence
		_ = _654_v11
		_654_v11 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_643_v0, _643_v0))
		var _655_v12 _dafny.Set
		_ = _655_v12
		_655_v12 = _dafny.SetOf(_dafny.IntOfInt64(297))
		var _pat_let_tv5 = _650_v7
		_ = _pat_let_tv5
		var _pat_let_tv6 = _650_v7
		_ = _pat_let_tv6
		var _pat_let_tv7 = _650_v7
		_ = _pat_let_tv7
		var _pat_let_tv8 = _650_v7
		_ = _pat_let_tv8
		var _rhs69 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_649_v6, _649_v6)
		_ = _rhs69
		var _rhs70 _dafny.CodePoint = (_649_v6).Select((Companion_Default___.SafeIndex((_651_v8).Cardinality(), _dafny.IntOfUint32((_649_v6).Cardinality()))).Uint32()).(_dafny.CodePoint)
		_ = _rhs70
		var _rhs71 _dafny.MultiSet = func(_source12 D3) _dafny.MultiSet {
			if _source12.Is_DC8() {
				var _656___mcc_h0 _dafny.MultiSet = _source12.Get_().(D3_DC8).Cf15
				_ = _656___mcc_h0
				var _657___mcc_h1 bool = _source12.Get_().(D3_DC8).Cf16
				_ = _657___mcc_h1
				var _658___mcc_h2 _dafny.Int = _source12.Get_().(D3_DC8).Cf17
				_ = _658___mcc_h2
				var _659_cf17 _dafny.Int = _658___mcc_h2
				_ = _659_cf17
				var _660_cf16 bool = _657___mcc_h1
				_ = _660_cf16
				var _661_cf15 _dafny.MultiSet = _656___mcc_h0
				_ = _661_cf15
				return (_pat_let_tv5).Intersection((Companion_D0_.Create_DC1_(_pat_let_tv6)).Dtor_cf1())
			} else if _source12.Is_DC9() {
				var _662___mcc_h3 bool = _source12.Get_().(D3_DC9).Cf18
				_ = _662___mcc_h3
				var _663___mcc_h4 bool = _source12.Get_().(D3_DC9).Cf19
				_ = _663___mcc_h4
				var _664___mcc_h5 _dafny.Array = _source12.Get_().(D3_DC9).Cf20
				_ = _664___mcc_h5
				var _665___mcc_h6 bool = _source12.Get_().(D3_DC9).Cf21
				_ = _665___mcc_h6
				var _666_cf21 bool = _665___mcc_h6
				_ = _666_cf21
				var _667_cf20 _dafny.Array = _664___mcc_h5
				_ = _667_cf20
				var _668_cf19 bool = _663___mcc_h4
				_ = _668_cf19
				var _669_cf18 bool = _662___mcc_h3
				_ = _669_cf18
				return _pat_let_tv7
			} else {
				var _670___mcc_h7 _dafny.Sequence = _source12.Get_().(D3_DC7).Cf14
				_ = _670___mcc_h7
				var _671_cf14 _dafny.Sequence = _670___mcc_h7
				_ = _671_cf14
				return _pat_let_tv8
			}
		}(_653_v10)
		_ = _rhs71
		var _rhs72 _dafny.Map = ((_644_v1).Merge(_644_v1)).Merge(((_654_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_649_v6).Cardinality()), _dafny.IntOfUint32((_654_v11).Cardinality()))).Uint32()).(_dafny.Map)).Merge((_654_v11).Select((Companion_Default___.SafeIndex(_643_v0, _dafny.IntOfUint32((_654_v11).Cardinality()))).Uint32()).(_dafny.Map)))
		_ = _rhs72
		var _rhs73 bool = (_this).Fm22(_dafny.IntOfUint32((_649_v6).Cardinality()), _643_v0, (_655_v12).IsSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_648_v5).Cardinality()))), globalState)
		_ = _rhs73
		var _lhs55 *GlobalState = globalState
		_ = _lhs55
		_649_v6 = _rhs69
		r2 = _rhs70
		_650_v7 = _rhs71
		_644_v1 = _rhs72
		_lhs55.F12 = _rhs73
		if !(Companion_Default___.Fm1((_this).F43(), (_643_v0).Times(_643_v0), globalState)) {
			var _672_v13 _dafny.CodePoint
			_ = _672_v13
			_672_v13 = _dafny.CodePoint('t')
			var _673_v14 _dafny.Map
			_ = _673_v14
			_673_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_v13, _648_v5)
			var _674_v15 _dafny.Array
			_ = _674_v15
			var _nwElement0_15 _dafny.Sequence = _648_v5
			_ = _nwElement0_15
			var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(24))
			_ = _nw94
			(_nw94).ArraySet1(_nwElement0_15, 0)
			(_nw94).ArraySet1(_648_v5, 1)
			(_nw94).ArraySet1(Companion_Default___.Fm23(_this.F28(), Companion_Default___.Fm24(_this.F29(), _dafny.IntOfInt64(489), globalState), globalState), 2)
			(_nw94).ArraySet1(_648_v5, 3)
			(_nw94).ArraySet1(_dafny.Companion_Sequence_.Update(_648_v5, (Companion_Default___.SafeIndex(_643_v0, _dafny.IntOfUint32((_648_v5).Cardinality()))).Uint32(), Companion_Default___.Fm24(false, Companion_Default___.Fm24(false, _643_v0, globalState), globalState)), 4)
			(_nw94).ArraySet1(_648_v5, 5)
			(_nw94).ArraySet1(_648_v5, 6)
			(_nw94).ArraySet1(_648_v5, 7)
			(_nw94).ArraySet1(_648_v5, 8)
			(_nw94).ArraySet1(_dafny.SeqOf(_643_v0), 9)
			(_nw94).ArraySet1(_648_v5, 10)
			(_nw94).ArraySet1(_648_v5, 11)
			(_nw94).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(376))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}((func(_675_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_676_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_675_v6).Cardinality())
				}
			})(_649_v6))), 12)
			(_nw94).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(955))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}((func(_677_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_678_i1 _dafny.Int) _dafny.Int {
					return _677_v0
				}
			})(_643_v0))), 13)
			(_nw94).ArraySet1(Companion_Default___.Fm23((_this).F43(), _dafny.IntOfUint32((_649_v6).Cardinality()), globalState), 14)
			(_nw94).ArraySet1(_648_v5, 15)
			(_nw94).ArraySet1(_648_v5, 16)
			(_nw94).ArraySet1((func() _dafny.Sequence {
				if (_673_v14).Contains(_672_v13) {
					return (_673_v14).Get(_672_v13).(_dafny.Sequence)
				}
				return _648_v5
			})(), 17)
			(_nw94).ArraySet1(Companion_Default___.Fm23(_this.F29(), _643_v0, globalState), 18)
			(_nw94).ArraySet1(_648_v5, 19)
			(_nw94).ArraySet1(_648_v5, 20)
			(_nw94).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_643_v0, _dafny.IntOfInt64(698), _dafny.IntOfInt64(327)), (Companion_Default___.SafeIndex((_648_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.IntOfUint32((_648_v5).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(_643_v0, _dafny.IntOfInt64(698), _dafny.IntOfInt64(327))).Cardinality()))).Uint32(), (_646_v3).Cardinality()), 21)
			(_nw94).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_648_v5, _648_v5), 22)
			(_nw94).ArraySet1((func() _dafny.Sequence {
				if _this.F28() {
					return _648_v5
				}
				return _648_v5
			})(), 23)
			_674_v15 = _nw94
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_674_v15), 0))
			_ = _index61
			(_674_v15).ArraySet1(_648_v5, (_index61).Int())
			var _679_v16 D6
			_ = _679_v16
			_679_v16 = Companion_D6_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-544))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_680_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_681_i2 _dafny.Int) _dafny.Int {
					return _680_v0
				}
			})(_643_v0))))
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(956), _dafny.ArrayLen((_674_v15), 0))
			_ = _index62
			(_674_v15).ArraySet1((_679_v16).Dtor_cf32(), (_index62).Int())
			var _682_v17 D4
			_ = _682_v17
			_682_v17 = Companion_D4_.Create_DC10_((_646_v3).Union(_646_v3))
			var _source13 D4 = _682_v17
			_ = _source13
			if _source13.Is_DC11() {
				var _683___mcc_h8 _dafny.Int = _source13.Get_().(D4_DC11).Cf23
				_ = _683___mcc_h8
				var _684___mcc_h9 bool = _source13.Get_().(D4_DC11).Cf24
				_ = _684___mcc_h9
				var _685___mcc_h10 _dafny.Int = _source13.Get_().(D4_DC11).Cf25
				_ = _685___mcc_h10
				var _686_cf25 _dafny.Int = _685___mcc_h10
				_ = _686_cf25
				var _687_cf24 bool = _684___mcc_h9
				_ = _687_cf24
				var _688_cf23 _dafny.Int = _683___mcc_h8
				_ = _688_cf23
				var _689_v18 _dafny.MultiSet
				_ = _689_v18
				_689_v18 = _dafny.MultiSetOf(_672_v13)
				var _690_v19 _dafny.Sequence
				_ = _690_v19
				_690_v19 = _dafny.SeqOf((_this).F43(), _this.F28())
				(globalState).F12 = (_646_v3).IsDisjointFrom((_646_v3).Update((Companion_D3_.Create_DC8_(_689_v18, (_690_v19).Select((Companion_Default___.SafeIndex(_686_cf25, _dafny.IntOfUint32((_690_v19).Cardinality()))).Uint32()).(bool), _688_cf23)).Dtor_cf17(), Companion_Default___.Abs(_dafny.IntOfUint32((_649_v6).Cardinality()))))
				var _691_v20 _dafny.Array
				_ = _691_v20
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw95
				_691_v20 = _nw95
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_691_v20), 0))
				_ = _index63
				(_691_v20).ArraySet1(Companion_Default___.Fm24(_687_cf24, _dafny.IntOfInt64(818), globalState), (_index63).Int())
				var _692_v21 _dafny.Sequence
				_ = _692_v21
				_692_v21 = _dafny.SeqOf(_649_v6, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_649_v6, _649_v6), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_643_v0, _690_v19)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_649_v6, _649_v6)).Cardinality()))).Uint32(), _672_v13))
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_691_v20), 0))
				_ = _index64
				var _rhs74 _dafny.Int = _dafny.IntOfUint32(((_692_v21).Select((Companion_Default___.SafeIndex((_686_cf25).Minus(Companion_Default___.Fm24((_this).F43(), Companion_Default___.Fm24(_this.F28(), _688_cf23, globalState), globalState)), _dafny.IntOfUint32((_692_v21).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
				_ = _rhs74
				var _rhs75 _dafny.Sequence = _690_v19
				_ = _rhs75
				var _lhs56 _dafny.Array = _691_v20
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_691_v20), 0))
				_ = _lhs57
				(_lhs56).ArraySet1(_rhs74, (_lhs57).Int())
				_690_v19 = _rhs75
				(globalState).F12 = _this.F29()
				_686_cf25 = ((_691_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_691_v20), 0))).Int()).(_dafny.Int)).Minus((_643_v0).Plus(_686_cf25))
			} else {
				var _693___mcc_h11 _dafny.MultiSet = _source13.Get_().(D4_DC10).Cf22
				_ = _693___mcc_h11
				var _694_cf22 _dafny.MultiSet = _693___mcc_h11
				_ = _694_cf22
				var _695_v22 _dafny.Array
				_ = _695_v22
				var _nw96 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw96
				_695_v22 = _nw96
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_695_v22), 0))
				_ = _index65
				(_695_v22).ArraySet1(_this.F29(), (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_695_v22), 0))
				_ = _index66
				(_695_v22).ArraySet1((_this).F43(), (_index66).Int())
				(globalState).F12 = _this.F29()
				r1 = (_643_v0).Plus((_643_v0).Minus((_dafny.SetOf((func() _dafny.Int {
					if (_644_v1).Contains(_643_v0) {
						return (_644_v1).Get(_643_v0).(_dafny.Int)
					}
					return _643_v0
				})())).Cardinality()))
				(globalState).F13 = _dafny.IntOfUint32((_649_v6).Cardinality())
			}
			_649_v6 = _649_v6
			(globalState).F1 = ((_this).F43()) == (true)
			(globalState).F13 = (_643_v0).Minus(_643_v0)
		} else {
			if _this.F29() {
				(globalState).F26 = (_this).F43()
				var _696_v23 _dafny.Array
				_ = _696_v23
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw97
				_696_v23 = _nw97
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_696_v23), 0))
				_ = _index67
				(_696_v23).ArraySet1(Companion_Default___.Fm1(_this.F28(), _643_v0, globalState), (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_696_v23), 0))
				_ = _index68
				(_696_v23).ArraySet1((func() bool {
					if _this.F28() {
						return _this.F29()
					}
					return (_this).F43()
				})(), (_index68).Int())
				var _697_v24 _dafny.MultiSet
				_ = _697_v24
				_697_v24 = _dafny.MultiSetOf(_648_v5)
				(globalState).F1 = (_697_v24).Equals(_697_v24)
				r1 = _643_v0
				r1 = _643_v0
			} else {
				var _698_v25 _dafny.Array
				_ = _698_v25
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw98
				_698_v25 = _nw98
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))
				_ = _index69
				(_698_v25).ArraySet1((_643_v0).Minus(_643_v0), (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))
				_ = _index70
				(_698_v25).ArraySet1(Companion_Default___.SafeModuloInt(_643_v0, _dafny.IntOfInt64(15)), (_index70).Int())
				r1 = _643_v0
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))
				_ = _index71
				(_698_v25).ArraySet1((func() _dafny.Int {
					if _this.F29() {
						return (_698_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))).Int()).(_dafny.Int)
					}
					return ((_698_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))).Int()).(_dafny.Int)).Times((_698_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))).Int()).(_dafny.Int))
				})(), (_index71).Int())
				var _699_v26 _dafny.Map
				_ = _699_v26
				_699_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_698_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))).Int()).(_dafny.Int), (_698_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))).Int()).(_dafny.Int))), _649_v6)
				_699_v26 = (_699_v26).Update((_698_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_698_v25), 0))).Int()).(_dafny.Int), _649_v6)
				var _700_v27 _dafny.Sequence
				_ = _700_v27
				_700_v27 = _dafny.SeqOf((_this).F43())
				var _701_v28 _dafny.Map
				_ = _701_v28
				_701_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_643_v0).Cmp(_dafny.IntOfUint32((_649_v6).Cardinality())) > 0, (_700_v27).Select((Companion_Default___.SafeIndex((_650_v7).Cardinality(), _dafny.IntOfUint32((_700_v27).Cardinality()))).Uint32()).(bool))
				_701_v28 = (_701_v28).Update(_this.F29(), _this.F28())
			}
			r2 = _dafny.CodePoint('h')
			var _702_v29 _dafny.Array
			_ = _702_v29
			var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw99
			_702_v29 = _nw99
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_702_v29), 0))
			_ = _index72
			(_702_v29).ArraySet1(_this.F28(), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_702_v29), 0))
			_ = _index73
			(_702_v29).ArraySet1(!((_643_v0).Cmp(_643_v0) >= 0), (_index73).Int())
			r3 = (func() bool {
				if !(_this.F28()) || ((_this).F43()) {
					return !((_this).Fm22(_643_v0, _643_v0, _this.F28(), globalState))
				}
				return _this.F29()
			})()
			var _703_v30 _dafny.Array
			_ = _703_v30
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_12
			var _nw100 _dafny.Array
			_ = _nw100
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw100 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Int = (func(_704_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_705_i3 _dafny.Int) _dafny.Int {
						return (_705_i3).Times(_704_v0)
					}
				})(_643_v0)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw100 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw100).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw100).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_703_v30 = _nw100
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_703_v30), 0))
			_ = _index74
			(_703_v30).ArraySet1((_643_v0).Times(_643_v0), (_index74).Int())
			var _706_v31 _dafny.Map
			_ = _706_v31
			_706_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F28())
			var _707_v32 _dafny.Set
			_ = _707_v32
			_707_v32 = _dafny.SetOf(_706_v31, _706_v31)
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_703_v30), 0))
			_ = _index75
			(_703_v30).ArraySet1((_707_v32).Cardinality(), (_index75).Int())
		}
		var _708_v33 _dafny.CodePoint
		_ = _708_v33
		_708_v33 = _dafny.CodePoint('r')
		r0 = _708_v33
		r1 = (_643_v0).Minus(_643_v0)
		r2 = _708_v33
		r3 = func(_source14 D6) bool {
			if _source14.Is_DC16() {
				var _709___mcc_h12 _dafny.Array = _source14.Get_().(D6_DC16).Cf33
				_ = _709___mcc_h12
				var _710_cf33 _dafny.Array = _709___mcc_h12
				_ = _710_cf33
				return (_this).F43()
			} else if _source14.Is_DC15() {
				var _711___mcc_h13 _dafny.Sequence = _source14.Get_().(D6_DC15).Cf32
				_ = _711___mcc_h13
				var _712_cf32 _dafny.Sequence = _711___mcc_h13
				_ = _712_cf32
				return !(_this.F28())
			} else {
				var _713___mcc_h14 D6 = _source14.Get_().(D6_DC17).Cf34
				_ = _713___mcc_h14
				var _714_cf34 D6 = _713___mcc_h14
				_ = _714_cf34
				return _this.F29()
			}
		}(Companion_D6_.Create_DC15_(_648_v5))
		return r0, r1, r2, r3
	}
}
func (_this *C1) M10(globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _715_v0 _dafny.Sequence
		_ = _715_v0
		_715_v0 = _dafny.UnicodeSeqOfUtf8Bytes("cjeniah")
		var _716_v1 _dafny.Int
		_ = _716_v1
		_716_v1 = _dafny.IntOfInt64(229)
		(globalState).F13 = (_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_715_v0).Cardinality()))).Times(_716_v1))
		var _717_v2 _dafny.Array
		_ = _717_v2
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_13
		var _nw101 _dafny.Array
		_ = _nw101
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw101 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Int = (func(_718_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_719_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_719_i0, _718_v1)
				}
			})(_716_v1)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw101 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw101).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw101).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_717_v2 = _nw101
		var _720_v3 _dafny.Set
		_ = _720_v3
		_720_v3 = _dafny.SetOf(_717_v2)
		var _721_v4 _dafny.Array
		_ = _721_v4
		var _nw102 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw102
		_721_v4 = _nw102
		var _722_v5 D2
		_ = _722_v5
		_722_v5 = Companion_D2_.Create_DC5_(_721_v4)
		var _723_v6 _dafny.Map
		_ = _723_v6
		_723_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_717_v2)).IsSubsetOf(_720_v3), (func() D2 {
			if _this.F29() {
				return _722_v5
			}
			return Companion_D2_.Create_DC5_(_721_v4)
		})())
		var _rhs76 _dafny.Int = (_dafny.Zero).Minus((_716_v1).Minus(Companion_Default___.SafeModuloInt(_716_v1, _dafny.IntOfInt64(390))))
		_ = _rhs76
		var _rhs77 bool = false
		_ = _rhs77
		var _rhs78 _dafny.Map = _723_v6
		_ = _rhs78
		var _lhs58 *GlobalState = globalState
		_ = _lhs58
		_716_v1 = _rhs76
		_lhs58.F12 = _rhs77
		_723_v6 = _rhs78
		var _724_i1 _dafny.Int
		_ = _724_i1
		_724_i1 = _dafny.Zero
		{
			for !(_this.F28()) {
				{
					if (_724_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_724_i1 = (_724_i1).Plus(_dafny.One)
					var _725_v7 _dafny.Array
					_ = _725_v7
					var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
					_ = _nw103
					_725_v7 = _nw103
					(globalState).F21 = _725_v7
					var _726_v8 _dafny.Array
					_ = _726_v8
					var _len0_14 _dafny.Int = _dafny.IntOfInt64(9)
					_ = _len0_14
					var _nw104 _dafny.Array
					_ = _nw104
					if _len0_14.Cmp(_dafny.Zero) == 0 {
						_nw104 = _dafny.NewArray(_len0_14)
					} else {
						var _init14 func(_dafny.Int) _dafny.CodePoint = func(_727_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						}
						_ = _init14
						var _element0_14 = _init14(_dafny.Zero)
						_ = _element0_14
						_nw104 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
						(_nw104).ArraySet1CodePoint(_element0_14, 0)
						var _nativeLen0_14 = (_len0_14).Int()
						_ = _nativeLen0_14
						for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
							(_nw104).ArraySet1CodePoint(_init14(_dafny.IntOf(_i0_14)), _i0_14)
						}
					}
					_726_v8 = _nw104
					var _728_v9 *C0
					_ = _728_v9
					var _nw105 *C0 = New_C0_()
					_ = _nw105
					_nw105.Ctor__(_726_v8)
					_728_v9 = _nw105
					var _729_v10 _dafny.Sequence
					_ = _729_v10
					_729_v10 = _dafny.SeqOf(_dafny.IntOfInt64(111), _716_v1, _716_v1)
					var _730_v11 _dafny.Sequence
					_ = _730_v11
					_730_v11 = _dafny.SeqOf(_this.F29())
					var _731_v12 _dafny.Int
					_ = _731_v12
					var _out48 _dafny.Int
					_ = _out48
					_out48 = Companion_Default___.M0(_729_v10, (_730_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg64 _dafny.Int) interface{} {
							return coer64(arg64)
						}
					}(func(_732_i3 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(758)
					}))).Cardinality()), _dafny.IntOfUint32((_730_v11).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus(_716_v1), _this.F29(), globalState)
					_731_v12 = _out48
					var _733_v13 _dafny.MultiSet
					_ = _733_v13
					_733_v13 = _dafny.MultiSetOf(_731_v12, _731_v12)
					var _734_v14 _dafny.Set
					_ = _734_v14
					_734_v14 = _dafny.SetOf(_733_v13)
					var _735_v15 _dafny.Map
					_ = _735_v15
					_735_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
						if _this.F28() {
							return _734_v14
						}
						return _734_v14
					})(), _731_v12)
					_735_v15 = (_735_v15).Update((_dafny.SetOf(Companion_Default___.Fm25(globalState))).Difference(_734_v14), _dafny.IntOfUint32((_729_v10).Cardinality()))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _736_v16 _dafny.Array
		_ = _736_v16
		var _nw106 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw106
		_736_v16 = _nw106
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_736_v16), 0))
		_ = _index76
		(_736_v16).ArraySet1((_this).F43(), (_index76).Int())
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_736_v16), 0))
		_ = _index77
		(_736_v16).ArraySet1(_this.F29(), (_index77).Int())
		var _737_v17 _dafny.Array
		_ = _737_v17
		var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
		_ = _nw107
		_737_v17 = _nw107
		var _738_v18 *C0
		_ = _738_v18
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__(_737_v17)
		_738_v18 = _nw108
		var _739_v19 _dafny.Set
		_ = _739_v19
		_739_v19 = _dafny.SetOf(_716_v1, _716_v1)
		_716_v1 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(442), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v19, _716_v1)).Cardinality())
		var _740_v20 _dafny.Map
		_ = _740_v20
		_740_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_736_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_736_v16), 0))).Int()).(bool)), _dafny.SeqOf(_this.F28(), !(!(!(_this.F29()))), _this.F29(), _this.F28(), _this.F29()))
		r0 = (_740_v20).Merge((_740_v20).Merge(Companion_Default___.Fm26(_this.F29(), globalState)))
		return r0
	}
}
func (_this *C1) F43() bool {
	{
		return _this._f43
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f30 _dafny.Int
	_f46 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f30 = _dafny.Zero
	_this._f46 = _dafny.Zero
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

func (_this *C2) F28() bool {
	return _this._f28
}
func (_this *C2) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C2) F29() bool {
	return _this._f29
}
func (_this *C2) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C2) F31() bool {
	return _this._f31
}
func (_this *C2) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C2) F30() _dafny.Int {
	return _this._f30
}
func (_this *C2) Ctor__(f46 _dafny.Int, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f46 = f46
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C2) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return (Companion_D5_.Create_DC12_(_dafny.UnicodeSeqOfUtf8Bytes("xlpme"))).Dtor_cf26()
	}
}
func (_this *C2) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			if _this.F28() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F30())
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F30())
		})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _dafny.IntOfInt64(-74)))
	}
}
func (_this *C2) Fm29(globalState *GlobalState) _dafny.Int {
	{
		var _source15 D5 = (func() D5 {
			if _this.F29() {
				return Companion_D5_.Create_DC13_(_this.F29(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), (_this).F30()), _dafny.CodePoint('k'))
			}
			return Companion_D5_.Create_DC13_(_this.F31(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F46()), _dafny.CodePoint('m'))
		})()
		_ = _source15
		if _source15.Is_DC13() {
			var _741___mcc_h0 bool = _source15.Get_().(D5_DC13).Cf27
			_ = _741___mcc_h0
			var _742___mcc_h1 _dafny.Map = _source15.Get_().(D5_DC13).Cf28
			_ = _742___mcc_h1
			var _743___mcc_h2 _dafny.CodePoint = _source15.Get_().(D5_DC13).Cf29
			_ = _743___mcc_h2
			var _744_cf29 _dafny.CodePoint = _743___mcc_h2
			_ = _744_cf29
			var _745_cf28 _dafny.Map = _742___mcc_h1
			_ = _745_cf28
			var _746_cf27 bool = _741___mcc_h0
			_ = _746_cf27
			return (_this).F46()
		} else if _source15.Is_DC14() {
			var _747___mcc_h3 bool = _source15.Get_().(D5_DC14).Cf30
			_ = _747___mcc_h3
			var _748___mcc_h4 _dafny.Int = _source15.Get_().(D5_DC14).Cf31
			_ = _748___mcc_h4
			var _749_cf31 _dafny.Int = _748___mcc_h4
			_ = _749_cf31
			var _750_cf30 bool = _747___mcc_h3
			_ = _750_cf30
			return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_749_cf31))).Cardinality()
		} else {
			var _751___mcc_h5 _dafny.Sequence = _source15.Get_().(D5_DC12).Cf26
			_ = _751___mcc_h5
			var _752_cf26 _dafny.Sequence = _751___mcc_h5
			_ = _752_cf26
			return (_this).F46()
		}
	}
}
func (_this *C2) Fm30(p0 bool, globalState *GlobalState) bool {
	{
		return _this.F28()
	}
}
func (_this *C2) M13(p0 _dafny.Sequence, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _753_v0 _dafny.Sequence
		_ = _753_v0
		_753_v0 = _dafny.SeqOf((_dafny.Zero).Minus(p3))
		var _754_i0 _dafny.Int
		_ = _754_i0
		_754_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsProperPrefixOf(_753_v0, _753_v0) {
				{
					if (_754_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_754_i0 = (_754_i0).Plus(_dafny.One)
					var _755_v1 _dafny.Array
					_ = _755_v1
					var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
					_ = _nw109
					_755_v1 = _nw109
					var _756_v2 _dafny.Map
					_ = _756_v2
					_756_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_755_v1)).Cardinality()))
					(globalState).F13 = (func() _dafny.Int {
						if p2 {
							return p3
						}
						return (func() _dafny.Int {
							if (_756_v2).Contains((_this).F46()) {
								return (_756_v2).Get((_this).F46()).(_dafny.Int)
							}
							return (_this).F30()
						})()
					})()
					(globalState).F13 = (_this).F30()
					if _this.F31() {
						(globalState).F12 = p2
						var _757_v3 _dafny.Sequence
						_ = _757_v3
						_757_v3 = _dafny.UnicodeSeqOfUtf8Bytes("lgyuevq")
						_757_v3 = (_this).Fm3(_this.F28(), (p3).Cmp(p3) <= 0, globalState)
						(globalState).F13 = (func() _dafny.Int {
							if (p1).Contains(_this.F28()) {
								return (p1).Get(_this.F28()).(_dafny.Int)
							}
							return (func() _dafny.Int {
								if !(true) {
									return (_this).F46()
								}
								return _dafny.IntOfUint32((_757_v3).Cardinality())
							})()
						})()
						_753_v0 = _dafny.SeqOf((_dafny.IntOfInt64(-583)).Minus((_this).F46()))
						(globalState).F13 = (_this).F46()
					} else {
						(globalState).F13 = (_this).F46()
						(globalState).F13 = (_this).F46()
						var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_755_v1), 0))
						_ = _index78
						(_755_v1).ArraySet1(_this.F31(), (_index78).Int())
						var _758_v4 _dafny.Sequence
						_ = _758_v4
						_758_v4 = _dafny.UnicodeSeqOfUtf8Bytes("d")
						var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_755_v1), 0))
						_ = _index79
						(_755_v1).ArraySet1((_dafny.IntOfUint32((_758_v4).Cardinality())).Cmp((_this).F46()) < 0, (_index79).Int())
						var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_755_v1), 0))
						_ = _index80
						(_755_v1).ArraySet1((_755_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_755_v1), 0))).Int()).(bool), (_index80).Int())
						var _759_v5 _dafny.Array
						_ = _759_v5
						var _nw110 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
						_ = _nw110
						_759_v5 = _nw110
						var _760_v6 _dafny.Map
						_ = _760_v6
						_760_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_759_v5, _dafny.Companion_Sequence_.Concatenate(_758_v4, _758_v4))
						_760_v6 = (_760_v6).Update(_759_v5, _758_v4)
					}
					var _761_v7 _dafny.Array
					_ = _761_v7
					var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(27))
					_ = _nw111
					_761_v7 = _nw111
					var _762_v8 _dafny.Set
					_ = _762_v8
					_762_v8 = _dafny.SetOf(_this.F31(), _this.F28())
					var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_761_v7), 0))
					_ = _index81
					(_761_v7).ArraySet1(_762_v8, (_index81).Int())
					var _763_v9 _dafny.Map
					_ = _763_v9
					_763_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F29())
					var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_761_v7), 0))
					_ = _index82
					var _rhs79 _dafny.Set = (_dafny.SetOf(_this.F28(), _this.F29())).Union(_762_v8)
					_ = _rhs79
					var _rhs80 _dafny.Int = (func() _dafny.Int {
						if (true) == (_this.F31()) {
							return (_this).F46()
						}
						return (_this).F46()
					})()
					_ = _rhs80
					var _rhs81 _dafny.Int = ((_this).F46()).Plus(((_763_v9).Cardinality()).Times((_this).F30()))
					_ = _rhs81
					var _lhs59 _dafny.Array = _761_v7
					_ = _lhs59
					var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_761_v7), 0))
					_ = _lhs60
					var _lhs61 *GlobalState = globalState
					_ = _lhs61
					var _lhs62 *GlobalState = globalState
					_ = _lhs62
					(_lhs59).ArraySet1(_rhs79, (_lhs60).Int())
					_lhs61.F13 = _rhs80
					_lhs62.F13 = _rhs81
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _764_v10 _dafny.Array
		_ = _764_v10
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_15
		var _nw112 _dafny.Array
		_ = _nw112
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw112 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.Set = (func(_765_p3 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_766_i2 _dafny.Int) _dafny.Set {
					return (_dafny.SetOf((_this).F46(), _dafny.IntOfInt64(-550))).Union(_dafny.SetOf(_765_p3))
				}
			})(p3)
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw112 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw112).ArraySet1(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw112).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_764_v10 = _nw112
		for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_764_v10), 0))); ; {
			_guard_loop_0, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _767_i1 _dafny.Int
			_767_i1 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_767_i1).Sign() != -1) && ((_767_i1).Cmp(_dafny.ArrayLen((_764_v10), 0)) < 0)) {
				(_764_v10).ArraySet1(((_dafny.SetOf((_this).F30(), (_this).F30())).Difference(_dafny.SetOf((_753_v0).Select((Companion_Default___.SafeIndex((_this).F46(), _dafny.IntOfUint32((_753_v0).Cardinality()))).Uint32()).(_dafny.Int)))).Intersection((_dafny.SetOf(p3)).Union(_dafny.SetOf(_dafny.IntOfInt64(950)))), (_767_i1).Int())
			}
		}
		var _768_v11 *C1
		_ = _768_v11
		var _nw113 *C1 = New_C1_()
		_ = _nw113
		_nw113.Ctor__(_this.F29(), _this.F29(), !(true) || (!(_this.F28())))
		_768_v11 = _nw113
		var _769_v12 _dafny.Array
		_ = _769_v12
		var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
		_ = _nw114
		_769_v12 = _nw114
		var _770_v13 *C0
		_ = _770_v13
		var _nw115 *C0 = New_C0_()
		_ = _nw115
		_nw115.Ctor__(_769_v12)
		_770_v13 = _nw115
		var _771_v14 _dafny.Sequence
		_ = _771_v14
		_771_v14 = _dafny.SeqOf(true)
		var _772_v15 _dafny.Sequence
		_ = _772_v15
		_772_v15 = _dafny.UnicodeSeqOfUtf8Bytes("xuygtgj")
		_771_v14 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_772_v15).Cardinality()), (_this).F30()), _this.F29(), _this.F29(), _dafny.IntOfInt64(808), globalState), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_772_v15).Cardinality()), (_this).F30()), _this.F29(), _this.F29(), _dafny.IntOfInt64(808), globalState)).Cardinality()))).Uint32(), true)
		var _773_v16 _dafny.Map
		_ = _773_v16
		_773_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(p2, (p1).Cardinality(), globalState), _770_v13.F37)
		var _774_v17 *C0
		_ = _774_v17
		var _nw116 *C0 = New_C0_()
		_ = _nw116
		_nw116.Ctor__((func() _dafny.Array {
			if (_773_v16).Contains(p2) {
				return (_773_v16).Get(p2).(_dafny.Array)
			}
			return _770_v13.F37
		})())
		_774_v17 = _nw116
		r0 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(61))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg65 _dafny.Int) interface{} {
				return coer65(arg65)
			}
		}(func(_775_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('u')
		})), _772_v15)).Cardinality())).Cmp((_753_v0).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_753_v0).Cardinality()))).Uint32()).(_dafny.Int)) <= 0
		return r0
	}
}
func (_this *C2) F46() _dafny.Int {
	{
		return _this._f46
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f36 _dafny.CodePoint
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
	_f44 _dafny.Int
	_f45 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f36 = _dafny.CodePoint('D')
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
	_this._f44 = _dafny.Zero
	_this._f45 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C3{}
var _ T3 = &C3{}
var _ T2 = &C3{}
var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F28() bool {
	return _this._f28
}
func (_this *C3) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C3) F29() bool {
	return _this._f29
}
func (_this *C3) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C3) F31() bool {
	return _this._f31
}
func (_this *C3) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C3) F35() bool {
	return _this._f35
}
func (_this *C3) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C3) F36() _dafny.CodePoint {
	return _this._f36
}
func (_this *C3) F36_set_(value _dafny.CodePoint) {
	_this._f36 = value
}
func (_this *C3) F34() _dafny.Int {
	return _this._f34
}
func (_this *C3) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C3) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C3) F30() _dafny.Int {
	return _this._f30
}
func (_this *C3) Ctor__(f44 _dafny.Int, f45 _dafny.Array, f36 _dafny.CodePoint, f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f44 = f44
		(_this)._f45 = f45
		(_this)._f36 = f36
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C3) Fm7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vsrndyx"), _dafny.UnicodeSeqOfUtf8Bytes("jsqkph")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mfxuo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(673))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}(func(_776_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F36()
		}))))
	}
}
func (_this *C3) Fm6(globalState *GlobalState) bool {
	{
		return _this.F29()
	}
}
func (_this *C3) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf((func() _dafny.Int {
			if _this.F35() {
				return (_this).F34()
			}
			return (_dafny.Zero).Minus((_this).F34())
		})(), (_this).F34(), ((_this).F34()).Minus((_this).F30()), (_this).F30(), Companion_Default___.SafeDivisionInt((_this).F30(), (_this).F30()))
	}
}
func (_this *C3) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32()), (_this).F32())
	}
}
func (_this *C3) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		var _source16 D5 = (func() D5 {
			if _this.F28() {
				return Companion_D5_.Create_DC14_(_this.F35(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F44(), (_this).F44())).Cardinality())
			}
			return Companion_D5_.Create_DC14_(_this.F29(), (_this).F44())
		})()
		_ = _source16
		if _source16.Is_DC13() {
			var _777___mcc_h0 bool = _source16.Get_().(D5_DC13).Cf27
			_ = _777___mcc_h0
			var _778___mcc_h1 _dafny.Map = _source16.Get_().(D5_DC13).Cf28
			_ = _778___mcc_h1
			var _779___mcc_h2 _dafny.CodePoint = _source16.Get_().(D5_DC13).Cf29
			_ = _779___mcc_h2
			var _780_cf29 _dafny.CodePoint = _779___mcc_h2
			_ = _780_cf29
			var _781_cf28 _dafny.Map = _778___mcc_h1
			_ = _781_cf28
			var _782_cf27 bool = _777___mcc_h0
			_ = _782_cf27
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pcm"), _dafny.UnicodeSeqOfUtf8Bytes("cca"))
		} else if _source16.Is_DC14() {
			var _783___mcc_h3 bool = _source16.Get_().(D5_DC14).Cf30
			_ = _783___mcc_h3
			var _784___mcc_h4 _dafny.Int = _source16.Get_().(D5_DC14).Cf31
			_ = _784___mcc_h4
			var _785_cf31 _dafny.Int = _784___mcc_h4
			_ = _785_cf31
			var _786_cf30 bool = _783___mcc_h3
			_ = _786_cf30
			return _dafny.UnicodeSeqOfUtf8Bytes("auph")
		} else {
			var _787___mcc_h5 _dafny.Sequence = _source16.Get_().(D5_DC12).Cf26
			_ = _787___mcc_h5
			var _788_cf26 _dafny.Sequence = _787___mcc_h5
			_ = _788_cf26
			return _788_cf26
		}
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _dafny.IntOfInt64(230))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F44()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), (_this).F34()))
	}
}
func (_this *C3) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		(globalState).F26 = ((_this).F44()).Cmp(_dafny.IntOfInt64(66)) >= 0
		var _789_v0 _dafny.Array
		_ = _789_v0
		var _nwElement0_16 _dafny.CodePoint = _this.F36()
		_ = _nwElement0_16
		var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(4))
		_ = _nw117
		(_nw117).ArraySet1CodePoint(_nwElement0_16, 0)
		(_nw117).ArraySet1CodePoint(_this.F36(), 1)
		(_nw117).ArraySet1CodePoint(_this.F36(), 2)
		(_nw117).ArraySet1CodePoint(_dafny.CodePoint('o'), 3)
		_789_v0 = _nw117
		var _790_v1 _dafny.Set
		_ = _790_v1
		_790_v1 = _dafny.SetOf(_789_v0, _789_v0)
		if ((_790_v1).Difference(_790_v1)).IsProperSubsetOf(_790_v1) {
			var _791_v2 _dafny.Sequence
			_ = _791_v2
			_791_v2 = _dafny.UnicodeSeqOfUtf8Bytes("mmi")
			var _792_v3 _dafny.Map
			_ = _792_v3
			_792_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_791_v2, (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_791_v2).Cardinality()))).Uint32(), (_791_v2).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_791_v2).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality()))
			var _793_v4 _dafny.MultiSet
			_ = _793_v4
			_793_v4 = _dafny.MultiSetOf((_this).F44())
			var _794_v5 T1
			_ = _794_v5
			var _nw118 *C2 = New_C2_()
			_ = _nw118
			_nw118.Ctor__((Companion_Default___.Fm31(Companion_Default___.Fm32(!(_this.F28()), globalState), _792_v3, (func(_pat_let9_0 D5) D5 {
				return func(_795_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let10_0 _dafny.Int) D5 {
						return func(_796_dt__update_hcf31_h0 _dafny.Int) D5 {
							return Companion_D5_.Create_DC14_((_795_dt__update__tmp_h0).Dtor_cf30(), _796_dt__update_hcf31_h0)
						}(_pat_let10_0)
					}(_dafny.IntOfInt64(-18))
				}(_pat_let9_0)
			}(Companion_D5_.Create_DC14_(_this.F35(), (_this).F44()))).Dtor_cf30(), ((_this).F32()).Select((Companion_Default___.SafeIndex((_793_v4).Cardinality(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), globalState)).Cardinality(), _dafny.IntOfInt64(454), ((_this).F34()).Cmp((_this).F30()) == 0, _this.F35(), ((_this).F32()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool))
			_794_v5 = _nw118
			_794_v5 = (func() T1 {
				if _this.F31() {
					return _794_v5
				}
				return _794_v5
			})()
			var _797_v6 _dafny.Map
			_ = _797_v6
			_797_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_794_v5.F31(), (_this).F44())
			var _798_v7 D5
			_ = _798_v7
			_798_v7 = Companion_D5_.Create_DC13_(_794_v5.F31(), _797_v6, _this.F36())
			var _source17 D5 = _798_v7
			_ = _source17
			if _source17.Is_DC13() {
				var _799___mcc_h0 bool = _source17.Get_().(D5_DC13).Cf27
				_ = _799___mcc_h0
				var _800___mcc_h1 _dafny.Map = _source17.Get_().(D5_DC13).Cf28
				_ = _800___mcc_h1
				var _801___mcc_h2 _dafny.CodePoint = _source17.Get_().(D5_DC13).Cf29
				_ = _801___mcc_h2
				var _802_cf29 _dafny.CodePoint = _801___mcc_h2
				_ = _802_cf29
				var _803_cf28 _dafny.Map = _800___mcc_h1
				_ = _803_cf28
				var _804_cf27 bool = _799___mcc_h0
				_ = _804_cf27
				(globalState).F13 = Companion_Default___.SafeModuloInt((_dafny.MultiSetOf((_this).F34())).Cardinality(), (_this).F34())
				var _805_v8 _dafny.Sequence
				_ = _805_v8
				_805_v8 = _dafny.SeqOf((_this).F34(), (_794_v5).F30())
				var _806_v9 _dafny.Array
				_ = _806_v9
				var _nwElement0_17 _dafny.Int = Companion_Default___.Fm32(_this.F28(), globalState)
				_ = _nwElement0_17
				var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(10))
				_ = _nw119
				(_nw119).ArraySet1(_nwElement0_17, 0)
				(_nw119).ArraySet1((_805_v8).Select((Companion_Default___.SafeIndex((_794_v5).F30(), _dafny.IntOfUint32((_805_v8).Cardinality()))).Uint32()).(_dafny.Int), 1)
				(_nw119).ArraySet1(((_794_v5).F30()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(238))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_807_cf29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_808_i0 _dafny.Int) _dafny.CodePoint {
						return _807_cf29
					}
				})(_802_cf29)))).Cardinality())), 2)
				(_nw119).ArraySet1(_dafny.IntOfUint32((_805_v8).Cardinality()), 3)
				(_nw119).ArraySet1((_dafny.Zero).Minus((_this).F30()), 4)
				(_nw119).ArraySet1((_794_v5).F30(), 5)
				(_nw119).ArraySet1((_this).F34(), 6)
				(_nw119).ArraySet1(_dafny.IntOfInt64(-91), 7)
				(_nw119).ArraySet1(Companion_Default___.SafeDivisionInt((_794_v5).F30(), (_794_v5).F30()), 8)
				(_nw119).ArraySet1((_this).F30(), 9)
				_806_v9 = _nw119
				var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw120
				_806_v9 = _nw120
				var _809_v10 *C0
				_ = _809_v10
				var _nw121 *C0 = New_C0_()
				_ = _nw121
				_nw121.Ctor__(_789_v0)
				_809_v10 = _nw121
				var _810_v11 D1
				_ = _810_v11
				_810_v11 = Companion_D1_.Create_DC3_(_794_v5.F29(), (_this).F34(), _809_v10, (_this).F44())
				var _811_v12 D1
				_ = _811_v12
				_811_v12 = Companion_D1_.Create_DC4_(_810_v11)
				var _812_v13 _dafny.MultiSet
				_ = _812_v13
				_812_v13 = _dafny.MultiSetOf(_811_v12, _811_v12)
				var _813_v14 _dafny.Sequence
				_ = _813_v14
				_813_v14 = _dafny.SeqOf(_811_v12)
				var _pat_let_tv9 = _810_v11
				_ = _pat_let_tv9
				var _814_v15 _dafny.Array
				_ = _814_v15
				var _nwElement0_18 _dafny.MultiSet = _dafny.MultiSetOf(_811_v12, func(_pat_let11_0 D1) D1 {
					return func(_815_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let12_0 D1) D1 {
							return func(_816_dt__update_hcf7_h0 D1) D1 {
								return Companion_D1_.Create_DC4_(_816_dt__update_hcf7_h0)
							}(_pat_let12_0)
						}(_pat_let_tv9)
					}(_pat_let11_0)
				}(_811_v12), _811_v12, Companion_D1_.Create_DC4_(_810_v11), _811_v12)
				_ = _nwElement0_18
				var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(14))
				_ = _nw122
				(_nw122).ArraySet1(_nwElement0_18, 0)
				(_nw122).ArraySet1((_812_v13).Union(_812_v13), 1)
				(_nw122).ArraySet1(_812_v13, 2)
				(_nw122).ArraySet1(_812_v13, 3)
				(_nw122).ArraySet1(_812_v13, 4)
				(_nw122).ArraySet1((_812_v13).Update(Companion_D1_.Create_DC4_(_810_v11), Companion_Default___.Abs(_dafny.IntOfUint32((_791_v2).Cardinality()))), 5)
				(_nw122).ArraySet1(_812_v13, 6)
				(_nw122).ArraySet1(_812_v13, 7)
				(_nw122).ArraySet1((_812_v13).Intersection(_812_v13), 8)
				(_nw122).ArraySet1(_dafny.MultiSetOf(Companion_D1_.Create_DC4_(_810_v11)), 9)
				(_nw122).ArraySet1(_812_v13, 10)
				(_nw122).ArraySet1(_812_v13, 11)
				(_nw122).ArraySet1((func() _dafny.MultiSet {
					if _794_v5.F29() {
						return _dafny.MultiSetFromSeq(_813_v14)
					}
					return _812_v13
				})(), 12)
				(_nw122).ArraySet1(_812_v13, 13)
				_814_v15 = _nw122
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_814_v15), 0))
				_ = _index83
				(_814_v15).ArraySet1(_dafny.MultiSetOf(_811_v12), (_index83).Int())
				var _817_v16 _dafny.Sequence
				_ = _817_v16
				_817_v16 = _dafny.SeqOf(_813_v14, _813_v14)
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_814_v15), 0))
				_ = _index84
				(_814_v15).ArraySet1(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if _794_v5.F31() {
						return _813_v14
					}
					return (_817_v16).Select((Companion_Default___.SafeIndex((_794_v5).F30(), _dafny.IntOfUint32((_817_v16).Cardinality()))).Uint32()).(_dafny.Sequence)
				})()), (_index84).Int())
				(globalState).F12 = _794_v5.F29()
			} else if _source17.Is_DC14() {
				var _818___mcc_h3 bool = _source17.Get_().(D5_DC14).Cf30
				_ = _818___mcc_h3
				var _819___mcc_h4 _dafny.Int = _source17.Get_().(D5_DC14).Cf31
				_ = _819___mcc_h4
				var _820_cf31 _dafny.Int = _819___mcc_h4
				_ = _820_cf31
				var _821_cf30 bool = _818___mcc_h3
				_ = _821_cf30
				(_794_v5).F31_set_(!(_794_v5.F31()))
				(globalState).F1 = !(_this.F28())
				var _822_v17 _dafny.Array
				_ = _822_v17
				var _nw123 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw123
				_822_v17 = _nw123
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_822_v17), 0))
				_ = _index85
				(_822_v17).ArraySet1(_794_v5.F31(), (_index85).Int())
				var _823_v18 _dafny.Map
				_ = _823_v18
				_823_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(242), false)
				var _824_v19 D2
				_ = _824_v19
				_824_v19 = Companion_D2_.Create_DC6_((_this).F34(), _794_v5.F28(), (_823_v18).Cardinality(), _this.F31(), (_794_v5).F30())
				var _825_v20 _dafny.Sequence
				_ = _825_v20
				_825_v20 = _dafny.SeqOf(_824_v19, _824_v19)
				var _826_v21 D2
				_ = _826_v21
				_826_v21 = Companion_D2_.Create_DC6_(Companion_Default___.SafeModuloInt(Companion_Default___.Fm32(_this.F31(), globalState), (_this).F30()), !_dafny.Companion_Sequence_.Contains(_825_v20, _824_v19), (_794_v5).F30(), _794_v5.F29(), (_this).F44())
				var _827_v22 D5
				_ = _827_v22
				_827_v22 = Companion_D5_.Create_DC14_(_794_v5.F31(), (_this).F30())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_822_v17), 0))
				_ = _index86
				var _rhs82 bool = _this.F31()
				_ = _rhs82
				var _rhs83 D2 = (func() D2 {
					if true {
						return _826_v21
					}
					return _824_v19
				})()
				_ = _rhs83
				var _rhs84 _dafny.Int = (_dafny.IntOfInt64(227)).Times((_this).F34())
				_ = _rhs84
				var _rhs85 bool = Companion_Default___.Fm1((_827_v22).Dtor_cf30(), ((_this).F30()).Plus((_this).F30()), globalState)
				_ = _rhs85
				var _rhs86 bool = (func() bool {
					if (func() bool {
						if _this.F35() {
							return _this.F28()
						}
						return _794_v5.F31()
					})() {
						return _794_v5.F31()
					}
					return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_821_cf30), (_this).F32())
				})()
				_ = _rhs86
				var _lhs63 _dafny.Array = _822_v17
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_822_v17), 0))
				_ = _lhs64
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				var _lhs66 T1 = _794_v5
				_ = _lhs66
				(_lhs63).ArraySet1(_rhs82, (_lhs64).Int())
				_824_v19 = _rhs83
				_820_cf31 = _rhs84
				_lhs65.F12 = _rhs85
				_lhs66.F28_set_(_rhs86)
				(globalState).F13 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if ((_this).F32()).Select((Companion_Default___.SafeIndex(_820_cf31, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool) {
						return _dafny.IntOfInt64(-11)
					}
					return (_794_v5).F30()
				})(), (_this).F30())
			} else {
				var _828___mcc_h5 _dafny.Sequence = _source17.Get_().(D5_DC12).Cf26
				_ = _828___mcc_h5
				var _829_cf26 _dafny.Sequence = _828___mcc_h5
				_ = _829_cf26
				var _830_v23 _dafny.Map
				_ = _830_v23
				_830_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_792_v3, _794_v5.F28())
				var _831_v24 _dafny.Set
				_ = _831_v24
				var _832_v25 bool
				_ = _832_v25
				var _out49 _dafny.Set
				_ = _out49
				var _out50 bool
				_ = _out50
				_out49, _out50 = (_this).M12((_this).F30(), _830_v23, globalState)
				_831_v24 = _out49
				_832_v25 = _out50
				(globalState).F13 = ((_this).F34()).Minus((_794_v5).F30())
				var _833_v26 _dafny.Array
				_ = _833_v26
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw124
				_833_v26 = _nw124
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_833_v26), 0))
				_ = _index87
				(_833_v26).ArraySet1((_797_v6).Cardinality(), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_833_v26), 0))
				_ = _index88
				var _rhs87 _dafny.Int = Companion_Default___.Fm32(_832_v25, globalState)
				_ = _rhs87
				var _rhs88 bool = _794_v5.F29()
				_ = _rhs88
				var _lhs67 _dafny.Array = _833_v26
				_ = _lhs67
				var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_833_v26), 0))
				_ = _lhs68
				var _lhs69 T1 = _794_v5
				_ = _lhs69
				(_lhs67).ArraySet1(_rhs87, (_lhs68).Int())
				_lhs69.F31_set_(_rhs88)
				var _834_v27 *C1
				_ = _834_v27
				var _nw125 *C1 = New_C1_()
				_ = _nw125
				_nw125.Ctor__(_this.F31(), _794_v5.F29(), _794_v5.F28())
				_834_v27 = _nw125
				var _835_v28 _dafny.MultiSet
				_ = _835_v28
				_835_v28 = _dafny.MultiSetOf(_834_v27)
				(globalState).F1 = ((_835_v28).Update(_834_v27, Companion_Default___.Abs(_dafny.IntOfInt64(895)))).IsDisjointFrom((_835_v28).Update(_834_v27, Companion_Default___.Abs((_this).F34())))
			}
			(globalState).F1 = _794_v5.F31()
			var _836_v29 *C0
			_ = _836_v29
			var _nw126 *C0 = New_C0_()
			_ = _nw126
			_nw126.Ctor__(_789_v0)
			_836_v29 = _nw126
			(_794_v5).F28_set_(true)
		} else {
			var _837_v30 _dafny.Map
			_ = _837_v30
			_837_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _this.F28())
			var _838_v31 D0
			_ = _838_v31
			_838_v31 = Companion_D0_.Create_DC1_(_dafny.MultiSetOf(!((func() bool {
				if (_837_v30).Contains(false) {
					return (_837_v30).Get(false).(bool)
				}
				return !(_this.F28())
			})())))
			_838_v31 = _838_v31
			(globalState).F13 = (_this).F30()
			(_this).F31_set_(((_this).F30()).Cmp((_this).F34()) >= 0)
			var _839_v32 _dafny.Set
			_ = _839_v32
			_839_v32 = _dafny.SetOf(_this.F28(), _this.F35())
			var _840_v33 _dafny.Sequence
			_ = _840_v33
			_840_v33 = _dafny.SeqOf((_839_v32).Cardinality())
			var _841_v34 _dafny.Map
			_ = _841_v34
			_841_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-187), _dafny.IntOfUint32((_840_v33).Cardinality()))
			var _842_v35 _dafny.Sequence
			_ = _842_v35
			_842_v35 = _dafny.SeqOf((_841_v34).Update((_this).F44(), (_this).F34()))
			var _rhs89 _dafny.Sequence = _842_v35
			_ = _rhs89
			var _rhs90 bool = _this.F28()
			_ = _rhs90
			var _rhs91 bool = _this.F29()
			_ = _rhs91
			var _lhs70 *GlobalState = globalState
			_ = _lhs70
			var _lhs71 *GlobalState = globalState
			_ = _lhs71
			_842_v35 = _rhs89
			_lhs70.F26 = _rhs90
			_lhs71.F1 = _rhs91
			var _843_v36 *C0
			_ = _843_v36
			var _nw127 *C0 = New_C0_()
			_ = _nw127
			_nw127.Ctor__(_789_v0)
			_843_v36 = _nw127
		}
		var _844_v37 _dafny.MultiSet
		_ = _844_v37
		_844_v37 = _dafny.MultiSetOf(_this.F31(), _this.F31())
		var _hi4 _dafny.Int = ((_this).F44()).Minus((func() _dafny.Int {
			if (_844_v37).Contains(_this.F31()) {
				return (_844_v37).Multiplicity(_this.F31())
			}
			return (_this).F44()
		})())
		_ = _hi4
		for _845_i1 := (_this).F30(); _845_i1.Cmp(_hi4) < 0; _845_i1 = _845_i1.Plus(_dafny.One) {
			(globalState).F12 = _this.F28()
			var _846_v38 D2
			_ = _846_v38
			_846_v38 = Companion_D2_.Create_DC6_(_dafny.IntOfInt64(51), _this.F28(), (_this).F30(), _this.F35(), _dafny.IntOfInt64(290))
			var _847_v39 _dafny.MultiSet
			_ = _847_v39
			_847_v39 = _dafny.MultiSetOf((_this).F34())
			var _848_v40 _dafny.Sequence
			_ = _848_v40
			_848_v40 = _dafny.UnicodeSeqOfUtf8Bytes("ksmswsjps")
			var _849_v41 _dafny.Set
			_ = _849_v41
			_849_v41 = _dafny.SetOf(_848_v40, _848_v40, _848_v40, _848_v40, _848_v40)
			var _850_v42 _dafny.Sequence
			_ = _850_v42
			_850_v42 = _dafny.SeqOf((_this).F34(), (_this).F44())
			var _851_v43 _dafny.Sequence
			_ = _851_v43
			_851_v43 = _dafny.SeqOf(_dafny.IntOfUint32((_850_v42).Cardinality()))
			var _852_v44 _dafny.Set
			_ = _852_v44
			_852_v44 = _dafny.SetOf(_dafny.IntOfUint32((_850_v42).Cardinality()))
			var _853_v45 _dafny.Set
			_ = _853_v45
			_853_v45 = _dafny.SetOf(_this.F28())
			var _854_v46 _dafny.Map
			_ = _854_v46
			_854_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F44(), _this.F28())
			var _855_v47 _dafny.Array
			_ = _855_v47
			var _nwElement0_19 bool = !((_846_v38).Dtor_cf10())
			_ = _nwElement0_19
			var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(25))
			_ = _nw128
			(_nw128).ArraySet1(_nwElement0_19, 0)
			(_nw128).ArraySet1((_this).Fm7(_847_v39, _dafny.IntOfInt64(125), globalState), 1)
			(_nw128).ArraySet1((_849_v41).IsSubsetOf(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("aa"), _848_v40)), 2)
			(_nw128).ArraySet1(_this.F29(), 3)
			(_nw128).ArraySet1(_this.F31(), 4)
			(_nw128).ArraySet1(_this.F29(), 5)
			(_nw128).ArraySet1(((_this).F34()).Cmp(_845_i1) < 0, 6)
			(_nw128).ArraySet1(_this.F31(), 7)
			(_nw128).ArraySet1(true, 8)
			(_nw128).ArraySet1(false, 9)
			(_nw128).ArraySet1(_this.F35(), 10)
			(_nw128).ArraySet1(_this.F35(), 11)
			(_nw128).ArraySet1(_this.F35(), 12)
			(_nw128).ArraySet1(_this.F35(), 13)
			(_nw128).ArraySet1(!(_this.F29()), 14)
			(_nw128).ArraySet1(_this.F28(), 15)
			(_nw128).ArraySet1((func() bool {
				if _this.F31() {
					return false
				}
				return _this.F31()
			})(), 16)
			(_nw128).ArraySet1((_852_v44).Equals(_852_v44), 17)
			(_nw128).ArraySet1(!(_this.F31()), 18)
			(_nw128).ArraySet1(_this.F31(), 19)
			(_nw128).ArraySet1((_853_v45).IsProperSubsetOf(_853_v45), 20)
			(_nw128).ArraySet1((_854_v46).Contains(_dafny.IntOfInt64(135)), 21)
			(_nw128).ArraySet1(_this.F29(), 22)
			(_nw128).ArraySet1(!(_dafny.MultiSetOf(_dafny.IntOfUint32((_848_v40).Cardinality()), (_this).F34())).Contains(_dafny.IntOfUint32((_850_v42).Cardinality())), 23)
			(_nw128).ArraySet1(false, 24)
			_855_v47 = _nw128
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_855_v47), 0))
			_ = _index89
			(_855_v47).ArraySet1((Companion_Default___.Fm32(_this.F29(), globalState)).Cmp(_845_i1) == 0, (_index89).Int())
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(277), _dafny.ArrayLen((_855_v47), 0))
			_ = _index90
			(_855_v47).ArraySet1((_this).Fm7(_847_v39, Companion_Default___.SafeDivisionInt((_853_v45).Cardinality(), (_this).F44()), globalState), (_index90).Int())
			var _856_v48 _dafny.MultiSet
			_ = _856_v48
			_856_v48 = _dafny.MultiSetOf(_this.F36(), _dafny.CodePoint('o'))
			var _857_v49 _dafny.Int
			_ = _857_v49
			var _out51 _dafny.Int
			_ = _out51
			_out51 = Companion_Default___.M0(_851_v43, !(_dafny.Companion_Sequence_.Contains(_848_v40, _this.F36())), (_856_v48).Cardinality(), !(_this.F28()), globalState)
			_857_v49 = _out51
			var _858_v50 _dafny.Map
			_ = _858_v50
			_858_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), _789_v0)
			var _859_v51 D0
			_ = _859_v51
			_859_v51 = Companion_D0_.Create_DC1_(Companion_Default___.Fm33(_this.F29(), globalState))
			_858_v50 = (_858_v50).Update((_859_v51).Dtor_cf1(), _789_v0)
		}
		var _860_v52 _dafny.Sequence
		_ = _860_v52
		_860_v52 = _dafny.SeqOf(_844_v37, _844_v37, _844_v37, _844_v37)
		var _861_v53 *C0
		_ = _861_v53
		var _nw129 *C0 = New_C0_()
		_ = _nw129
		_nw129.Ctor__(_789_v0)
		_861_v53 = _nw129
		var _862_v54 _dafny.Sequence
		_ = _862_v54
		_862_v54 = _dafny.UnicodeSeqOfUtf8Bytes("jrneeb")
		var _863_v55 D1
		_ = _863_v55
		_863_v55 = Companion_D1_.Create_DC3_(_this.F35(), (_this).F34(), _861_v53, _dafny.IntOfUint32((_862_v54).Cardinality()))
		var _864_v56 _dafny.MultiSet
		_ = _864_v56
		_864_v56 = _dafny.MultiSetOf((_844_v37).Cardinality(), (_this).F34())
		var _865_v57 _dafny.Array
		_ = _865_v57
		var _nwElement0_20 _dafny.Int = (_this).F34()
		_ = _nwElement0_20
		var _nw130 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(26))
		_ = _nw130
		(_nw130).ArraySet1(_nwElement0_20, 0)
		(_nw130).ArraySet1((_this).F30(), 1)
		(_nw130).ArraySet1(((_860_v52).Select((Companion_Default___.SafeIndex((_863_v55).Dtor_cf4(), _dafny.IntOfUint32((_860_v52).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), 2)
		(_nw130).ArraySet1((_this).F34(), 3)
		(_nw130).ArraySet1(((_this).F44()).Plus(_dafny.IntOfUint32((_862_v54).Cardinality())), 4)
		(_nw130).ArraySet1(((_this).F44()).Times((_this).F34()), 5)
		(_nw130).ArraySet1((_this).F30(), 6)
		(_nw130).ArraySet1(_dafny.IntOfInt64(-578), 7)
		(_nw130).ArraySet1((_this).F44(), 8)
		(_nw130).ArraySet1((_this).F30(), 9)
		(_nw130).ArraySet1((_this).F34(), 10)
		(_nw130).ArraySet1((_this).F30(), 11)
		(_nw130).ArraySet1((_this).F44(), 12)
		(_nw130).ArraySet1((_this).F34(), 13)
		(_nw130).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.IntOfInt64(337))).Cardinality()).Times((_this).F30()), 14)
		(_nw130).ArraySet1(((_this).F30()).Plus((_this).F44()), 15)
		(_nw130).ArraySet1((_dafny.Zero).Minus((_this).F34()), 16)
		(_nw130).ArraySet1((_this).F34(), 17)
		(_nw130).ArraySet1(((_this).F34()).Minus((_this).F30()), 18)
		(_nw130).ArraySet1((_this).F34(), 19)
		(_nw130).ArraySet1((func() _dafny.Int {
			if _this.F28() {
				return (_dafny.Zero).Minus((_this).F34())
			}
			return (_this).F44()
		})(), 20)
		(_nw130).ArraySet1((func() _dafny.Int {
			if (_864_v56).Contains((_dafny.Zero).Minus((_this).F44())) {
				return (_864_v56).Multiplicity((_dafny.Zero).Minus((_this).F44()))
			}
			return _dafny.IntOfUint32((_dafny.SeqOf(_this.F36())).Cardinality())
		})(), 21)
		(_nw130).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F44(), _dafny.IntOfUint32((_862_v54).Cardinality())), 22)
		(_nw130).ArraySet1((_this).F44(), 23)
		(_nw130).ArraySet1((_this).F44(), 24)
		(_nw130).ArraySet1((func() _dafny.Int {
			if (_844_v37).Contains(!(_this.F35())) {
				return (_844_v37).Multiplicity(!(_this.F35()))
			}
			return (_this).F44()
		})(), 25)
		_865_v57 = _nw130
		_865_v57 = _865_v57
		if (_this).Fm6(globalState) {
			if _this.F35() {
				var _866_v58 _dafny.Array
				_ = _866_v58
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw131
				_866_v58 = _nw131
				var _rhs92 bool = _this.F31()
				_ = _rhs92
				var _rhs93 _dafny.Array = _866_v58
				_ = _rhs93
				var _rhs94 _dafny.Array = _865_v57
				_ = _rhs94
				var _rhs95 bool = _this.F28()
				_ = _rhs95
				var _rhs96 bool = false
				_ = _rhs96
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				var _lhs73 *GlobalState = globalState
				_ = _lhs73
				var _lhs74 *C3 = _this
				_ = _lhs74
				r0 = _rhs92
				_lhs72.F21 = _rhs93
				_865_v57 = _rhs94
				_lhs73.F12 = _rhs95
				_lhs74.F28_set_(_rhs96)
				var _867_v59 _dafny.Sequence
				_ = _867_v59
				_867_v59 = _dafny.SeqOf(_865_v57, _865_v57)
				var _rhs97 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("d")
				_ = _rhs97
				var _rhs98 _dafny.Array = _866_v58
				_ = _rhs98
				var _rhs99 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_this).Fm3(!(!_dafny.Companion_Sequence_.Contains(_867_v59, _865_v57)), _this.F29(), globalState), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_864_v56).Cardinality(), (_this).F30()), _dafny.IntOfUint32(((_this).Fm3(!(!_dafny.Companion_Sequence_.Contains(_867_v59, _865_v57)), _this.F29(), globalState)).Cardinality()))).Uint32(), (_862_v54).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_860_v52).Cardinality())), _dafny.IntOfUint32((_862_v54).Cardinality()))).Uint32()).(_dafny.CodePoint))
				_ = _rhs99
				var _lhs75 *GlobalState = globalState
				_ = _lhs75
				_862_v54 = _rhs97
				_lhs75.F21 = _rhs98
				_862_v54 = _rhs99
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_865_v57), 0))
				_ = _index91
				(_865_v57).ArraySet1((_this).F30(), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_865_v57), 0))
				_ = _index92
				(_865_v57).ArraySet1(((_this).F34()).Minus((_this).F30()), (_index92).Int())
				var _868_v60 _dafny.Sequence
				_ = _868_v60
				_868_v60 = _dafny.SeqOf((_dafny.Zero).Minus((_865_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_865_v57), 0))).Int()).(_dafny.Int)))
				(globalState).F12 = ((_dafny.Zero).Minus((_dafny.IntOfUint32((_868_v60).Cardinality())).Times((_865_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_865_v57), 0))).Int()).(_dafny.Int)))).Cmp((_this).F34()) > 0
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_866_v58), 0))
				_ = _index93
				(_866_v58).ArraySet1(!(false), (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_866_v58), 0))
				_ = _index94
				(_866_v58).ArraySet1(((_this).F44()).Cmp((_this).F30()) < 0, (_index94).Int())
			} else {
				r2 = _862_v54
				var _869_v61 _dafny.Map
				_ = _869_v61
				_869_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _this.F36())
				var _870_v62 _dafny.Map
				_ = _870_v62
				_870_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if (_869_v61).Contains(_this.F35()) {
						return (_869_v61).Get(_this.F35()).(_dafny.CodePoint)
					}
					return _this.F36()
				})(), _dafny.IntOfUint32((_862_v54).Cardinality()))
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_865_v57), 0))
				_ = _index95
				(_865_v57).ArraySet1(Companion_Default___.SafeDivisionInt((_870_v62).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), (_this).F34())).Cardinality()), (_index95).Int())
				var _871_v63 D2
				_ = _871_v63
				_871_v63 = Companion_D2_.Create_DC6_((_this).F34(), _this.F31(), (_this).F34(), _this.F29(), (_this).F30())
				var _872_v64 _dafny.MultiSet
				_ = _872_v64
				_872_v64 = _dafny.MultiSetOf(_this.F36(), _this.F36(), _this.F36(), _this.F36())
				var _873_v65 D3
				_ = _873_v65
				_873_v65 = Companion_D3_.Create_DC8_(_872_v64, _this.F28(), (_this).F44())
				var _874_v66 _dafny.Array
				_ = _874_v66
				var _nwElement0_21 bool = (_871_v63).Dtor_cf10()
				_ = _nwElement0_21
				var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(11))
				_ = _nw132
				(_nw132).ArraySet1(_nwElement0_21, 0)
				(_nw132).ArraySet1(_this.F31(), 1)
				(_nw132).ArraySet1(!(_844_v37).Equals(_dafny.MultiSetOf(_this.F31())), 2)
				(_nw132).ArraySet1(false, 3)
				(_nw132).ArraySet1(((_this).F34()).Cmp(((_873_v65).Dtor_cf15()).Cardinality()) != 0, 4)
				(_nw132).ArraySet1(_this.F35(), 5)
				(_nw132).ArraySet1(_this.F29(), 6)
				(_nw132).ArraySet1(_this.F35(), 7)
				(_nw132).ArraySet1(_this.F29(), 8)
				(_nw132).ArraySet1((_dafny.IntOfUint32((_862_v54).Cardinality())).Cmp((_this).F30()) <= 0, 9)
				(_nw132).ArraySet1(_this.F28(), 10)
				_874_v66 = _nw132
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_874_v66), 0))
				_ = _index96
				(_874_v66).ArraySet1((_this.F35()) || (_this.F29()), (_index96).Int())
				var _875_v67 _dafny.Set
				_ = _875_v67
				_875_v67 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}(func(_876_i2 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}))).Cardinality()))
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_865_v57), 0))
				_ = _index97
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_874_v66), 0))
				_ = _index98
				var _rhs100 _dafny.Int = (_this).F44()
				_ = _rhs100
				var _rhs101 bool = (func() _dafny.Set {
					var _coll23 = _dafny.NewBuilder()
					_ = _coll23
					for _iter24 := _dafny.Iterate((_875_v67).Elements()); ; {
						_compr_23, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _877_v68 _dafny.Int
						_877_v68 = interface{}(_compr_23).(_dafny.Int)
						if (_875_v67).Contains(_877_v68) {
							_coll23.Add(Companion_Default___.SafeModuloInt(_877_v68, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aohlb")).Cardinality())))
						}
					}
					return _coll23.ToSet()
				}()).IsSubsetOf(_dafny.SetOf((_this).F30()))
				_ = _rhs101
				var _lhs76 _dafny.Array = _865_v57
				_ = _lhs76
				var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_865_v57), 0))
				_ = _lhs77
				var _lhs78 _dafny.Array = _874_v66
				_ = _lhs78
				var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_874_v66), 0))
				_ = _lhs79
				(_lhs76).ArraySet1(_rhs100, (_lhs77).Int())
				(_lhs78).ArraySet1(_rhs101, (_lhs79).Int())
				(globalState).F13 = ((_871_v63).Dtor_cf11()).Times((_865_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_865_v57), 0))).Int()).(_dafny.Int))
				var _878_v69 _dafny.Map
				_ = _878_v69
				_878_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_874_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_874_v66), 0))).Int()).(bool), (_865_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_865_v57), 0))).Int()).(_dafny.Int))
				var _879_v70 _dafny.MultiSet
				_ = _879_v70
				_879_v70 = _dafny.MultiSetOf(_878_v69, _878_v69, (func() _dafny.Map {
					if _this.F35() {
						return _878_v69
					}
					return _878_v69
				})())
				_879_v70 = _879_v70
				(globalState).F12 = (_874_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_874_v66), 0))).Int()).(bool)
			}
			var _880_v71 _dafny.Array
			_ = _880_v71
			var _len0_16 _dafny.Int = _dafny.One
			_ = _len0_16
			var _nw133 _dafny.Array
			_ = _nw133
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw133 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) bool = func(_881_i3 _dafny.Int) bool {
					return (func() bool {
						if _this.F29() {
							return _this.F28()
						}
						return _this.F35()
					})()
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw133 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw133).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw133).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_880_v71 = _nw133
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_880_v71), 0))
			_ = _index99
			(_880_v71).ArraySet1(_this.F29(), (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_880_v71), 0))
			_ = _index100
			(_880_v71).ArraySet1(!_dafny.Companion_Sequence_.Contains((_this).F32(), !(_this.F28()) || (_this.F28())), (_index100).Int())
			(globalState).F13 = Companion_Default___.Fm32(true, globalState)
			var _882_v72 _dafny.Map
			_ = _882_v72
			_882_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _this.F31())
			(_this).F36_set_(Companion_Default___.Fm34(((_882_v72).Update((_this).F44(), _this.F28())).Cardinality(), (_this).F34(), (_this).F44(), _this.F31(), globalState))
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_865_v57), 0))
			_ = _index101
			(_865_v57).ArraySet1((_this).F30(), (_index101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_865_v57), 0))
			_ = _index102
			(_865_v57).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm32((_880_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(727), _dafny.ArrayLen((_880_v71), 0))).Int()).(bool), globalState), (_this).F30()), (_index102).Int())
		} else {
			(globalState).F13 = (Companion_Default___.Fm32(_this.F28(), globalState)).Times((_this).F30())
			r0 = _this.F28()
			var _883_v73 _dafny.Sequence
			_ = _883_v73
			_883_v73 = _dafny.SeqOf((_this).F34(), _dafny.IntOfInt64(307), Companion_Default___.SafeDivisionInt((_this).F34(), _dafny.IntOfInt64(381)), (_this).F44(), (_this).F30())
			var _884_v74 _dafny.MultiSet
			_ = _884_v74
			_884_v74 = _dafny.MultiSetOf(_883_v73)
			var _rhs102 _dafny.Sequence = _dafny.SeqOf((func() _dafny.Int {
				if _this.F31() {
					return (_884_v74).Cardinality()
				}
				return (_this).F44()
			})(), (_dafny.Zero).Minus((_this).F30()), (_this).F44(), (_this).F34(), Companion_Default___.SafeModuloInt((_this).F34(), _dafny.IntOfInt64(-445)))
			_ = _rhs102
			var _rhs103 bool = !(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_883_v73, _883_v73))).Contains((_this).F30())
			_ = _rhs103
			var _rhs104 bool = !(((_this).F44()).Cmp(((_this).F30()).Plus((_this).F34())) < 0)
			_ = _rhs104
			var _lhs80 *C3 = _this
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			_883_v73 = _rhs102
			_lhs80.F29_set_(_rhs103)
			_lhs81.F26 = _rhs104
			var _885_v75 _dafny.Array
			_ = _885_v75
			var _nwElement0_22 bool = true
			_ = _nwElement0_22
			var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(10))
			_ = _nw134
			(_nw134).ArraySet1(_nwElement0_22, 0)
			(_nw134).ArraySet1(_this.F28(), 1)
			(_nw134).ArraySet1((_this).Fm7(_dafny.MultiSetFromSeq(_883_v73), (_this).F30(), globalState), 2)
			(_nw134).ArraySet1(!(_this.F29()), 3)
			(_nw134).ArraySet1(true, 4)
			(_nw134).ArraySet1(!(_this.F29()), 5)
			(_nw134).ArraySet1(_this.F28(), 6)
			(_nw134).ArraySet1(_this.F31(), 7)
			(_nw134).ArraySet1(_this.F28(), 8)
			(_nw134).ArraySet1(_this.F35(), 9)
			_885_v75 = _nw134
			var _886_v76 _dafny.Map
			_ = _886_v76
			_886_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F44(), _885_v75)
			var _887_v77 _dafny.Map
			_ = _887_v77
			_887_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F44(), _862_v54)
			_886_v76 = (_886_v76).Update(((_this).F34()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex((_887_v77).Cardinality(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), _this.F29()), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex((_887_v77).Cardinality(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), _this.F29())).Cardinality()))).Uint32(), _this.F31())).Cardinality())), _885_v75)
			(globalState).F13 = (_this).F30()
		}
		r2 = _862_v54
		var _888_v78 _dafny.Map
		_ = _888_v78
		_888_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_864_v56).Equals(_864_v56), _this.F28())
		var _889_v79 _dafny.Map
		_ = _889_v79
		_889_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F30())
		r0 = (func() bool {
			if (_888_v78).Contains((_889_v79).Contains((_this).F30())) {
				return (_888_v78).Get((_889_v79).Contains((_this).F30())).(bool)
			}
			return _this.F28()
		})()
		var _890_v80 _dafny.Map
		_ = _890_v80
		_890_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), !(_this.F31()))
		var _891_v81 _dafny.Sequence
		_ = _891_v81
		_891_v81 = _dafny.SeqOf((_890_v80).Cardinality())
		var _892_v82 D6
		_ = _892_v82
		_892_v82 = Companion_D6_.Create_DC15_(_891_v81)
		var _893_v83 _dafny.Map
		_ = _893_v83
		_893_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_892_v82, (_this).F34())
		var _894_v84 _dafny.Sequence
		_ = _894_v84
		_894_v84 = _dafny.SeqOf((_this).F44(), (func() _dafny.Int {
			if (_893_v83).Contains(_892_v82) {
				return (_893_v83).Get(_892_v82).(_dafny.Int)
			}
			return _dafny.IntOfInt64(601)
		})())
		r1 = (_894_v84).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_894_v84).Cardinality()))).Uint32()).(_dafny.Int)
		r2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(926))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_895_i4 _dafny.Int) _dafny.CodePoint {
			return _this.F36()
		})), (Companion_Default___.SafeIndex((_this).F44(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(926))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}(func(_896_i4 _dafny.Int) _dafny.CodePoint {
			return _this.F36()
		}))).Cardinality()))).Uint32(), _this.F36()), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_862_v54, _862_v54), (Companion_Default___.SafeIndex((_this).F44(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_862_v54, _862_v54)).Cardinality()))).Uint32(), _this.F36()))
		return r0, r1, r2
	}
}
func (_this *C3) M12(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) (_dafny.Set, bool) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 bool = false
		_ = r1
		var _897_v0 _dafny.Array
		_ = _897_v0
		var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
		_ = _nw135
		_897_v0 = _nw135
		for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_897_v0), 0))); ; {
			_guard_loop_1, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _898_i0 _dafny.Int
			_898_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_898_i0).Sign() != -1) && ((_898_i0).Cmp(_dafny.ArrayLen((_897_v0), 0)) < 0)) {
				(_897_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pftnlhv"), _dafny.UnicodeSeqOfUtf8Bytes("mekvfibbt")), (_898_i0).Int())
			}
		}
		var _899_v1 _dafny.Sequence
		_ = _899_v1
		_899_v1 = _dafny.SeqOf((_this).F34())
		var _900_v2 _dafny.Int
		_ = _900_v2
		var _out52 _dafny.Int
		_ = _out52
		_out52 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_this).F30()), (_this).F44()), _899_v1), _this.F35(), (_this).F44(), _this.F31(), globalState)
		_900_v2 = _out52
		var _901_v3 D3
		_ = _901_v3
		_901_v3 = Companion_D3_.Create_DC8_(_dafny.MultiSetOf(_this.F36()), _this.F28(), p0)
		if (_this.F29()) || ((_901_v3).Dtor_cf16()) {
			var _902_v4 _dafny.Array
			_ = _902_v4
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_17
			var _nw136 _dafny.Array
			_ = _nw136
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw136 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.CodePoint = func(_903_i1 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw136 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw136).ArraySet1CodePoint(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw136).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_902_v4 = _nw136
			var _904_v5 *C0
			_ = _904_v5
			var _nw137 *C0 = New_C0_()
			_ = _nw137
			_nw137.Ctor__(_902_v4)
			_904_v5 = _nw137
			var _905_v6 *C0
			_ = _905_v6
			var _nw138 *C0 = New_C0_()
			_ = _nw138
			_nw138.Ctor__(_904_v5.F37)
			_905_v6 = _nw138
			var _906_v7 _dafny.Sequence
			_ = _906_v7
			_906_v7 = _dafny.UnicodeSeqOfUtf8Bytes("eaganir")
			(globalState).F13 = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_906_v7).Cardinality()))).Plus((_this).F34())
			var _907_v8 _dafny.Set
			_ = _907_v8
			_907_v8 = _dafny.SetOf(_902_v4)
			var _908_v9 _dafny.Int
			_ = _908_v9
			var _out53 _dafny.Int
			_ = _out53
			_out53 = Companion_Default___.M0(_dafny.SeqOf((_this).F44()), !((_907_v8).IsProperSubsetOf(_907_v8)), p0, _this.F28(), globalState)
			_908_v9 = _out53
			var _909_v10 _dafny.Map
			_ = _909_v10
			_909_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F35())
			var _910_v11 *C2
			_ = _910_v11
			var _nw139 *C2 = New_C2_()
			_ = _nw139
			_nw139.Ctor__((_this).F30(), (_dafny.Zero).Minus((_909_v10).Cardinality()), _this.F35(), (_this.F31()) || (_this.F29()), _this.F28())
			_910_v11 = _nw139
		} else {
			var _911_v12 _dafny.Array
			_ = _911_v12
			var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw140
			_911_v12 = _nw140
			var _912_v13 _dafny.Array
			_ = _912_v13
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_18
			var _nw141 _dafny.Array
			_ = _nw141
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw141 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = func(_913_i2 _dafny.Int) bool {
					return _this.F29()
				}
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw141 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw141).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw141).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_912_v13 = _nw141
			var _nwElement0_23 _dafny.Array = _912_v13
			_ = _nwElement0_23
			var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
			_ = _nw142
			(_nw142).ArraySet1(_nwElement0_23, 0)
			(_nw142).ArraySet1(_912_v13, 1)
			(_nw142).ArraySet1(_912_v13, 2)
			(_nw142).ArraySet1(_912_v13, 3)
			(_nw142).ArraySet1(_912_v13, 4)
			(_nw142).ArraySet1(_912_v13, 5)
			(_nw142).ArraySet1(_912_v13, 6)
			(_nw142).ArraySet1(_912_v13, 7)
			(_nw142).ArraySet1(_912_v13, 8)
			(_nw142).ArraySet1(_912_v13, 9)
			(_nw142).ArraySet1(_912_v13, 10)
			(_nw142).ArraySet1(_912_v13, 11)
			(_nw142).ArraySet1(_912_v13, 12)
			_911_v12 = _nw142
			var _914_v14 _dafny.Sequence
			_ = _914_v14
			_914_v14 = _dafny.SeqOf(_this.F28(), _this.F35())
			var _915_v15 _dafny.MultiSet
			_ = _915_v15
			_915_v15 = _dafny.MultiSetOf((_this).F34())
			var _916_v16 _dafny.Array
			_ = _916_v16
			var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw143
			_916_v16 = _nw143
			var _917_v17 _dafny.Map
			_ = _917_v17
			_917_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_916_v16, _912_v13)
			var _918_v18 _dafny.Sequence
			_ = _918_v18
			_918_v18 = _dafny.UnicodeSeqOfUtf8Bytes("l")
			var _919_v19 _dafny.MultiSet
			_ = _919_v19
			_919_v19 = _dafny.MultiSetOf(_918_v18)
			var _rhs105 _dafny.Int = ((_dafny.MultiSetOf(_918_v18)).Intersection((_919_v19).Update(_918_v18, Companion_Default___.Abs((_this).F34())))).Cardinality()
			_ = _rhs105
			var _rhs106 bool = _this.F28()
			_ = _rhs106
			var _rhs107 _dafny.Sequence = (_this).F32()
			_ = _rhs107
			var _rhs108 _dafny.MultiSet = (_915_v15).Update(_900_v2, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqOf(p0, (_this).F30(), (_899_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_899_v1).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality())))
			_ = _rhs108
			var _rhs109 _dafny.Map = _917_v17
			_ = _rhs109
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			_900_v2 = _rhs105
			_lhs82.F1 = _rhs106
			_914_v14 = _rhs107
			_915_v15 = _rhs108
			_917_v17 = _rhs109
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_912_v13), 0))
			_ = _index103
			(_912_v13).ArraySet1(!((((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F44(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)) == (_this.F29())), (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_912_v13), 0))
			_ = _index104
			var _rhs110 bool = _this.F35()
			_ = _rhs110
			var _rhs111 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_914_v14, _dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32()))
			_ = _rhs111
			var _rhs112 _dafny.Int = _900_v2
			_ = _rhs112
			var _rhs113 _dafny.CodePoint = (_918_v18).Select((Companion_Default___.SafeIndex(_900_v2, _dafny.IntOfUint32((_918_v18).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs113
			var _rhs114 bool = _this.F31()
			_ = _rhs114
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			var _lhs84 *C3 = _this
			_ = _lhs84
			var _lhs85 _dafny.Array = _912_v13
			_ = _lhs85
			var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_912_v13), 0))
			_ = _lhs86
			_lhs83.F26 = _rhs110
			_914_v14 = _rhs111
			_900_v2 = _rhs112
			_lhs84.F36_set_(_rhs113)
			(_lhs85).ArraySet1(_rhs114, (_lhs86).Int())
			var _920_v20 *C1
			_ = _920_v20
			var _nw144 *C1 = New_C1_()
			_ = _nw144
			_nw144.Ctor__(_this.F35(), ((_this).F32()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), _this.F31())
			_920_v20 = _nw144
			(globalState).F1 = _this.F28()
		}
		var _921_i3 _dafny.Int
		_ = _921_i3
		_921_i3 = _dafny.Zero
		{
			for _this.F29() {
				{
					if (_921_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_921_i3 = (_921_i3).Plus(_dafny.One)
					var _922_v21 _dafny.Sequence
					_ = _922_v21
					_922_v21 = _dafny.SeqOf(_this.F28())
					var _rhs115 _dafny.Sequence = (func() _dafny.Sequence {
						if _this.F28() {
							return (_this).F32()
						}
						return _922_v21
					})()
					_ = _rhs115
					var _rhs116 _dafny.Int = _900_v2
					_ = _rhs116
					var _rhs117 bool = !(_this.F31()) || (_this.F28())
					_ = _rhs117
					var _lhs87 *GlobalState = globalState
					_ = _lhs87
					_922_v21 = _rhs115
					_900_v2 = _rhs116
					_lhs87.F12 = _rhs117
					var _923_v22 _dafny.Array
					_ = _923_v22
					var _len0_19 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_19
					var _nw145 _dafny.Array
					_ = _nw145
					if _len0_19.Cmp(_dafny.Zero) == 0 {
						_nw145 = _dafny.NewArray(_len0_19)
					} else {
						var _init19 func(_dafny.Int) bool = (func(_924_p0 _dafny.Int) func(_dafny.Int) bool {
							return func(_925_i4 _dafny.Int) bool {
								return (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _this.F28()))).Cardinality())).IsSubsetOf(_dafny.MultiSetOf(_924_p0))
							}
						})(p0)
						_ = _init19
						var _element0_19 = _init19(_dafny.Zero)
						_ = _element0_19
						_nw145 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
						(_nw145).ArraySet1(_element0_19, 0)
						var _nativeLen0_19 = (_len0_19).Int()
						_ = _nativeLen0_19
						for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
							(_nw145).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
						}
					}
					_923_v22 = _nw145
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_923_v22), 0))
					_ = _index105
					(_923_v22).ArraySet1(_this.F29(), (_index105).Int())
					var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_923_v22), 0))
					_ = _index106
					(_923_v22).ArraySet1(_this.F28(), (_index106).Int())
					var _926_v23 _dafny.MultiSet
					_ = _926_v23
					_926_v23 = _dafny.MultiSetOf((_this).F34(), (_this).F34(), Companion_Default___.Fm32(!((_923_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_923_v22), 0))).Int()).(bool)), globalState), (_this).F44())
					r1 = (_this).Fm7(_926_v23, (_this).F44(), globalState)
					var _927_v24 _dafny.Array
					_ = _927_v24
					var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(4))
					_ = _nw146
					_927_v24 = _nw146
					var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_927_v24), 0))
					_ = _index107
					(_927_v24).ArraySet1CodePoint(_this.F36(), (_index107).Int())
					var _928_v25 D5
					_ = _928_v25
					_928_v25 = Companion_D5_.Create_DC13_(_this.F29(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_923_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_923_v22), 0))).Int()).(bool), (_this).F30()), _dafny.CodePoint('y'))
					var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_927_v24), 0))
					_ = _index108
					(_927_v24).ArraySet1CodePoint((_928_v25).Dtor_cf29(), (_index108).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _929_v26 *C2
		_ = _929_v26
		var _nw147 *C2 = New_C2_()
		_ = _nw147
		_nw147.Ctor__((_this).F34(), (_dafny.Zero).Minus(_900_v2), !(_this.F35()), !(_this.F31()), _this.F31())
		_929_v26 = _nw147
		var _930_v27 _dafny.Sequence
		_ = _930_v27
		_930_v27 = _dafny.SeqOf(_this.F36(), _this.F36())
		var _931_v28 _dafny.Map
		_ = _931_v28
		_931_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _dafny.Companion_Sequence_.IsPrefixOf(_930_v27, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(377))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg71 _dafny.Int) interface{} {
				return coer71(arg71)
			}
		}(func(_932_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))))
		var _933_v29 _dafny.Array
		_ = _933_v29
		var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
		_ = _nw148
		_933_v29 = _nw148
		var _934_v30 _dafny.MultiSet
		_ = _934_v30
		_934_v30 = _dafny.MultiSetOf(_this.F36())
		var _935_v31 _dafny.Sequence
		_ = _935_v31
		_935_v31 = _dafny.SeqOf(_933_v29)
		var _rhs118 _dafny.Int = (func() _dafny.Int {
			if (_934_v30).Contains(_this.F36()) {
				return (_934_v30).Multiplicity(_this.F36())
			}
			return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _this.F35())).Cardinality()).Times(p0)
		})()
		_ = _rhs118
		var _rhs119 _dafny.Map = _931_v28
		_ = _rhs119
		var _rhs120 _dafny.Array = (_935_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F32()).Cardinality()), _dafny.IntOfUint32((_935_v31).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs120
		var _lhs88 *GlobalState = globalState
		_ = _lhs88
		_lhs88.F13 = _rhs118
		_931_v28 = _rhs119
		_933_v29 = _rhs120
		var _936_v32 _dafny.Set
		_ = _936_v32
		_936_v32 = _dafny.SetOf(_900_v2)
		r0 = (_936_v32).Union((_936_v32).Union(_936_v32))
		r1 = _this.F28()
		return r0, r1
	}
}
func (_this *C3) F44() _dafny.Int {
	{
		return _this._f44
	}
}
func (_this *C3) F45() _dafny.Array {
	{
		return _this._f45
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T3 = &C4{}
var _ T2 = &C4{}
var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F28() bool {
	return _this._f28
}
func (_this *C4) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C4) F29() bool {
	return _this._f29
}
func (_this *C4) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C4) F31() bool {
	return _this._f31
}
func (_this *C4) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C4) F35() bool {
	return _this._f35
}
func (_this *C4) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C4) F34() _dafny.Int {
	return _this._f34
}
func (_this *C4) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C4) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C4) F30() _dafny.Int {
	return _this._f30
}
func (_this *C4) Ctor__(f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C4) Fm6(globalState *GlobalState) bool {
	{
		return !(_this.F31())
	}
}
func (_this *C4) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return ((func() _dafny.Set {
			var _coll24 = _dafny.NewBuilder()
			_ = _coll24
			for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(460), _dafny.IntOfInt64(658))); ; {
				_compr_24, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _937_v0 _dafny.Int
				_937_v0 = interface{}(_compr_24).(_dafny.Int)
				if ((_dafny.IntOfInt64(460)).Cmp(_937_v0) <= 0) && ((_937_v0).Cmp(_dafny.IntOfInt64(658)) < 0) {
					_coll24.Add(Companion_Default___.SafeModuloInt(_937_v0, (_this).F34()))
				}
			}
			return _coll24.ToSet()
		}()).Difference(_dafny.SetOf(_dafny.IntOfInt64(972)))).Union((_dafny.SetOf((_this).F34())).Union(_dafny.SetOf((_dafny.SetOf(_this.F28())).Cardinality())))
	}
}
func (_this *C4) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return (_this).F32()
	}
}
func (_this *C4) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		if (_dafny.MultiSetOf((_this).F30())).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(845))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg72 _dafny.Int) interface{} {
				return coer72(arg72)
			}
		}(func(_938_i0 _dafny.Int) _dafny.Int {
			return (_this).F30()
		})))) {
			return _dafny.UnicodeSeqOfUtf8Bytes("fe")
		} else {
			return _dafny.UnicodeSeqOfUtf8Bytes("sjcbfbtrq")
		}
	}
}
func (_this *C4) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(!(!(((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool))))).IsSubsetOf(_dafny.SetOf(_this.F31())), (_this).F34())
	}
}
func (_this *C4) Fm44(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeModuloInt((_this).F30(), (_dafny.Zero).Minus((_this).F30()))).Cmp((_this).F30()) > 0
	}
}
func (_this *C4) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		(_this).F35_set_(_this.F31())
		var _939_i0 _dafny.Int
		_ = _939_i0
		_939_i0 = _dafny.Zero
		{
			for (_this).Fm6(globalState) {
				{
					if (_939_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_939_i0 = (_939_i0).Plus(_dafny.One)
					var _940_v0 _dafny.MultiSet
					_ = _940_v0
					_940_v0 = _dafny.MultiSetOf(_this.F28())
					var _941_v1 _dafny.Set
					_ = _941_v1
					_941_v1 = _dafny.SetOf(_this.F31())
					var _942_v2 _dafny.Array
					_ = _942_v2
					var _nwElement0_24 bool = _this.F31()
					_ = _nwElement0_24
					var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(20))
					_ = _nw149
					(_nw149).ArraySet1(_nwElement0_24, 0)
					(_nw149).ArraySet1(_this.F28(), 1)
					(_nw149).ArraySet1(_this.F28(), 2)
					(_nw149).ArraySet1(_this.F29(), 3)
					(_nw149).ArraySet1(_this.F28(), 4)
					(_nw149).ArraySet1(_this.F29(), 5)
					(_nw149).ArraySet1(_this.F28(), 6)
					(_nw149).ArraySet1(_this.F28(), 7)
					(_nw149).ArraySet1(_this.F28(), 8)
					(_nw149).ArraySet1(false, 9)
					(_nw149).ArraySet1(_this.F29(), 10)
					(_nw149).ArraySet1(_this.F29(), 11)
					(_nw149).ArraySet1(_this.F35(), 12)
					(_nw149).ArraySet1(_this.F29(), 13)
					(_nw149).ArraySet1(_this.F29(), 14)
					(_nw149).ArraySet1(true, 15)
					(_nw149).ArraySet1(true, 16)
					(_nw149).ArraySet1((_this).Fm6(globalState), 17)
					(_nw149).ArraySet1(_this.F35(), 18)
					(_nw149).ArraySet1((_this).Fm6(globalState), 19)
					_942_v2 = _nw149
					var _943_v3 _dafny.Map
					_ = _943_v3
					_943_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_942_v2, _941_v1)
					var _944_v4 _dafny.Sequence
					_ = _944_v4
					_944_v4 = _dafny.SeqOf(_941_v1)
					var _945_v5 _dafny.Array
					_ = _945_v5
					var _nwElement0_25 _dafny.Set = Companion_Default___.Fm45((_this).F30(), _940_v0, _this.F35(), _dafny.IntOfUint32(((_this).F32()).Cardinality()), globalState)
					_ = _nwElement0_25
					var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(23))
					_ = _nw150
					(_nw150).ArraySet1(_nwElement0_25, 0)
					(_nw150).ArraySet1((_941_v1).Difference(_dafny.SetOf(_this.F29(), false, _this.F29(), _this.F35(), _this.F31())), 1)
					(_nw150).ArraySet1(_941_v1, 2)
					(_nw150).ArraySet1((func() _dafny.Set {
						if (_943_v3).Contains(_942_v2) {
							return (_943_v3).Get(_942_v2).(_dafny.Set)
						}
						return _941_v1
					})(), 3)
					(_nw150).ArraySet1((_941_v1).Difference(_941_v1), 4)
					(_nw150).ArraySet1(_941_v1, 5)
					(_nw150).ArraySet1(_941_v1, 6)
					(_nw150).ArraySet1(_941_v1, 7)
					(_nw150).ArraySet1(_941_v1, 8)
					(_nw150).ArraySet1((_dafny.SetOf(_this.F35())).Difference(Companion_Default___.Fm45((_dafny.Zero).Minus((_this).F30()), _940_v0, _this.F35(), (_this).F34(), globalState)), 9)
					(_nw150).ArraySet1((_941_v1).Intersection(_941_v1), 10)
					(_nw150).ArraySet1((_944_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.IntOfUint32((_944_v4).Cardinality()))).Uint32()).(_dafny.Set), 11)
					(_nw150).ArraySet1(_941_v1, 12)
					(_nw150).ArraySet1(_941_v1, 13)
					(_nw150).ArraySet1(_941_v1, 14)
					(_nw150).ArraySet1(_941_v1, 15)
					(_nw150).ArraySet1(_941_v1, 16)
					(_nw150).ArraySet1(_941_v1, 17)
					(_nw150).ArraySet1(Companion_Default___.Fm45((_this).F30(), _940_v0, _this.F35(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm6(globalState), (_this).F34())).Cardinality(), globalState), 18)
					(_nw150).ArraySet1(_dafny.SetOf(true), 19)
					(_nw150).ArraySet1(_941_v1, 20)
					(_nw150).ArraySet1(_941_v1, 21)
					(_nw150).ArraySet1(_941_v1, 22)
					_945_v5 = _nw150
					_945_v5 = _945_v5
					var _946_v6 _dafny.MultiSet
					_ = _946_v6
					_946_v6 = _dafny.MultiSetOf((_this).F30(), _dafny.IntOfInt64(355))
					if ((_this).F32()).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm46((_this).F34(), (_946_v6).Cardinality(), _this.F28(), globalState), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool) {
						var _947_v7 _dafny.Set
						_ = _947_v7
						_947_v7 = _dafny.SetOf((_this).F34(), (_this).F34())
						r1 = Companion_Default___.SafeDivisionInt(((_947_v7).Difference(_947_v7)).Cardinality(), (_946_v6).Cardinality())
						var _948_v8 _dafny.Map
						_ = _948_v8
						_948_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F29())
						var _949_v9 _dafny.Array
						_ = _949_v9
						var _nwElement0_26 _dafny.Int = _dafny.IntOfUint32(((_this).F32()).Cardinality())
						_ = _nwElement0_26
						var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(12))
						_ = _nw151
						(_nw151).ArraySet1(_nwElement0_26, 0)
						(_nw151).ArraySet1((_this).F30(), 1)
						(_nw151).ArraySet1((_this).F30(), 2)
						(_nw151).ArraySet1(((_this).F34()).Times((_this).F30()), 3)
						(_nw151).ArraySet1((_948_v8).Cardinality(), 4)
						(_nw151).ArraySet1((_this).F30(), 5)
						(_nw151).ArraySet1((_this).F30(), 6)
						(_nw151).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F30(), (_this).F30()), 7)
						(_nw151).ArraySet1((_this).F34(), 8)
						(_nw151).ArraySet1((func() _dafny.Int {
							if _this.F29() {
								return (_this).F34()
							}
							return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("khil")).Cardinality())
						})(), 9)
						(_nw151).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F34(), (_this).F34())), 10)
						(_nw151).ArraySet1((_this).F34(), 11)
						_949_v9 = _nw151
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_949_v9), 0))
						_ = _index109
						(_949_v9).ArraySet1(((_this).F30()).Times((_dafny.Zero).Minus((_this).F30())), (_index109).Int())
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_949_v9), 0))
						_ = _index110
						(_949_v9).ArraySet1(_dafny.IntOfInt64(802), (_index110).Int())
						(globalState).F13 = ((_947_v7).Union(_947_v7)).Cardinality()
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_949_v9), 0))
						_ = _index111
						(_949_v9).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_949_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_949_v9), 0))).Int()).(_dafny.Int))), (_index111).Int())
						(globalState).F1 = ((_this).F30()).Cmp((_949_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_949_v9), 0))).Int()).(_dafny.Int)) < 0
					} else {
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_942_v2), 0))
						_ = _index112
						(_942_v2).ArraySet1(((_this).F30()).Cmp((_this).F30()) == 0, (_index112).Int())
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_942_v2), 0))
						_ = _index113
						(_942_v2).ArraySet1(_this.F29(), (_index113).Int())
						(globalState).F13 = ((_946_v6).Difference(_946_v6)).Cardinality()
						(globalState).F13 = (_this).F30()
						r1 = _dafny.IntOfUint32(((_this).F32()).Cardinality())
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_942_v2), 0))
						_ = _index114
						(_942_v2).ArraySet1(_this.F28(), (_index114).Int())
					}
					var _950_v10 _dafny.CodePoint
					_ = _950_v10
					_950_v10 = _dafny.CodePoint('p')
					_950_v10 = _950_v10
					var _951_v11 _dafny.Array
					_ = _951_v11
					var _nwElement0_27 _dafny.Array = _942_v2
					_ = _nwElement0_27
					var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(18))
					_ = _nw152
					(_nw152).ArraySet1(_nwElement0_27, 0)
					(_nw152).ArraySet1(_942_v2, 1)
					(_nw152).ArraySet1(_942_v2, 2)
					(_nw152).ArraySet1(_942_v2, 3)
					(_nw152).ArraySet1(_942_v2, 4)
					(_nw152).ArraySet1(_942_v2, 5)
					(_nw152).ArraySet1(_942_v2, 6)
					(_nw152).ArraySet1(_942_v2, 7)
					(_nw152).ArraySet1(_942_v2, 8)
					(_nw152).ArraySet1(_942_v2, 9)
					(_nw152).ArraySet1(_942_v2, 10)
					(_nw152).ArraySet1(_942_v2, 11)
					(_nw152).ArraySet1(_942_v2, 12)
					(_nw152).ArraySet1(_942_v2, 13)
					(_nw152).ArraySet1(_942_v2, 14)
					(_nw152).ArraySet1(_942_v2, 15)
					(_nw152).ArraySet1(_942_v2, 16)
					(_nw152).ArraySet1(_942_v2, 17)
					_951_v11 = _nw152
					var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_951_v11), 0))
					_ = _index115
					(_951_v11).ArraySet1(_942_v2, (_index115).Int())
					var _952_v12 _dafny.MultiSet
					_ = _952_v12
					_952_v12 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32()), _dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32()), (_this).F32(), _dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).Fm5((_this).F30(), (_this).F34(), !(_this.F35()), (_this).F30(), globalState)))
					var _953_v13 _dafny.Set
					_ = _953_v13
					_953_v13 = _dafny.SetOf((_this).F34())
					var _954_v14 _dafny.Set
					_ = _954_v14
					_954_v14 = _953_v13
					var _955_v15 _dafny.Sequence
					_ = _955_v15
					_955_v15 = _dafny.SeqOf((_954_v14), _953_v13, _953_v13)
					var _956_v16 _dafny.Map
					_ = _956_v16
					_956_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_953_v13).IsSubsetOf((_955_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.IntOfUint32((_955_v15).Cardinality()))).Uint32()).(_dafny.Set)))
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_951_v11), 0))
					_ = _index116
					var _rhs121 _dafny.Array = _942_v2
					_ = _rhs121
					var _rhs122 _dafny.MultiSet = (_952_v12).Difference((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F32(), (_this).F32()), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F32(), (_this).F32())).Cardinality()))).Uint32(), (_this).F32()))).Union(_952_v12))
					_ = _rhs122
					var _rhs123 _dafny.Set = _953_v13
					_ = _rhs123
					var _rhs124 _dafny.Map = (_956_v16).Merge(_956_v16)
					_ = _rhs124
					var _lhs89 _dafny.Array = _951_v11
					_ = _lhs89
					var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_951_v11), 0))
					_ = _lhs90
					(_lhs89).ArraySet1(_rhs121, (_lhs90).Int())
					_952_v12 = _rhs122
					_953_v13 = _rhs123
					_956_v16 = _rhs124
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r1 = (_this).F30()
		if _this.F35() {
			var _957_v17 _dafny.Array
			_ = _957_v17
			var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
			_ = _nw153
			_957_v17 = _nw153
			var _958_v18 _dafny.CodePoint
			_ = _958_v18
			_958_v18 = _dafny.CodePoint('p')
			var _959_v19 _dafny.Map
			_ = _959_v19
			_959_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(768), _958_v18)
			var _960_v20 _dafny.Map
			_ = _960_v20
			_960_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.UnicodeSeqOfUtf8Bytes("qefsgg"))
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_957_v17), 0))
			_ = _index117
			(_957_v17).ArraySet1((Companion_Default___.Fm47(_this.F29(), _959_v19, globalState)).Merge(_960_v20), (_index117).Int())
			var _961_v21 _dafny.Array
			_ = _961_v21
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw154
			_961_v21 = _nw154
			var _962_v22 _dafny.Map
			_ = _962_v22
			_962_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_961_v21, _this.F28())
			var _963_v23 _dafny.Sequence
			_ = _963_v23
			_963_v23 = _dafny.SeqOf((_this).F34())
			var _964_v24 D8
			_ = _964_v24
			_964_v24 = Companion_D8_.Create_DC24_(_this.F28(), (_this).F34(), (_this).F32(), _this.F31(), _963_v23)
			var _965_v25 _dafny.Array
			_ = _965_v25
			var _nwElement0_28 bool = ((_962_v22).Cardinality()).Cmp((_this).F34()) > 0
			_ = _nwElement0_28
			var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(9))
			_ = _nw155
			(_nw155).ArraySet1(_nwElement0_28, 0)
			(_nw155).ArraySet1(_this.F31(), 1)
			(_nw155).ArraySet1((func() bool {
				if _this.F31() {
					return true
				}
				return Companion_Default___.Fm1(_this.F31(), (_this).F30(), globalState)
			})(), 2)
			(_nw155).ArraySet1((_dafny.IntOfInt64(711)).Cmp(_dafny.IntOfInt64(238)) == 0, 3)
			(_nw155).ArraySet1(!_dafny.Companion_Sequence_.Contains((_964_v24).Dtor_cf50(), (_963_v23).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_963_v23).Cardinality()))).Uint32()).(_dafny.Int)), 4)
			(_nw155).ArraySet1(true, 5)
			(_nw155).ArraySet1(_this.F29(), 6)
			(_nw155).ArraySet1(_this.F31(), 7)
			(_nw155).ArraySet1(!_dafny.Companion_Sequence_.Equal((_this).F32(), _dafny.SeqOf(_this.F31(), _this.F31())), 8)
			_965_v25 = _nw155
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_957_v17), 0))
			_ = _index118
			var _rhs125 bool = _this.F31()
			_ = _rhs125
			var _rhs126 _dafny.Int = ((_this).F30()).Plus(_dafny.IntOfInt64(845))
			_ = _rhs126
			var _rhs127 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yowl")).Cardinality())).Plus((_this).F34()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(526))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg73 _dafny.Int) interface{} {
					return coer73(arg73)
				}
			}((func(_966_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_967_i1 _dafny.Int) _dafny.CodePoint {
					return _966_v18
				}
			})(_958_v18))))
			_ = _rhs127
			var _rhs128 _dafny.Array = _965_v25
			_ = _rhs128
			var _rhs129 bool = true
			_ = _rhs129
			var _lhs91 *GlobalState = globalState
			_ = _lhs91
			var _lhs92 _dafny.Array = _957_v17
			_ = _lhs92
			var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_957_v17), 0))
			_ = _lhs93
			var _lhs94 *GlobalState = globalState
			_ = _lhs94
			var _lhs95 *GlobalState = globalState
			_ = _lhs95
			_lhs91.F26 = _rhs125
			r1 = _rhs126
			(_lhs92).ArraySet1(_rhs127, (_lhs93).Int())
			_lhs94.F21 = _rhs128
			_lhs95.F12 = _rhs129
			(globalState).F13 = ((_this).F34()).Plus((_this).F34())
			var _968_v26 _dafny.Sequence
			_ = _968_v26
			_968_v26 = _dafny.UnicodeSeqOfUtf8Bytes("uiiih")
			(_this).F31_set_(((_this).F32()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_968_v26, _968_v26)).Cardinality()))), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool))
			var _969_v27 _dafny.Map
			_ = _969_v27
			_969_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _this.F29())
			var _970_v28 _dafny.Array
			_ = _970_v28
			var _nwElement0_29 _dafny.Int = (_dafny.SetOf(_dafny.IntOfInt64(931))).Cardinality()
			_ = _nwElement0_29
			var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(5))
			_ = _nw156
			(_nw156).ArraySet1(_nwElement0_29, 0)
			(_nw156).ArraySet1((_dafny.Zero).Minus((_this).F34()), 1)
			(_nw156).ArraySet1(((_969_v27).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-557), _this.F35()))).Cardinality(), 2)
			(_nw156).ArraySet1(Companion_Default___.SafeModuloInt((_this).F30(), (_this).F34()), 3)
			(_nw156).ArraySet1((_this).F30(), 4)
			_970_v28 = _nw156
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_970_v28), 0))
			_ = _index119
			(_970_v28).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F28()), (_this).F34())).Cardinality(), (_index119).Int())
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_970_v28), 0))
			_ = _index120
			(_970_v28).ArraySet1((_this).F30(), (_index120).Int())
			var _971_v29 _dafny.Array
			_ = _971_v29
			var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw157
			_971_v29 = _nw157
			var _972_v30 _dafny.MultiSet
			_ = _972_v30
			_972_v30 = _dafny.MultiSetOf(_dafny.IntOfUint32((_963_v23).Cardinality()))
			var _973_v31 _dafny.Set
			_ = _973_v31
			_973_v31 = _dafny.SetOf((_this).F34(), (_972_v30).Cardinality())
			var _974_v32 D1
			_ = _974_v32
			_974_v32 = Companion_D1_.Create_DC2_(_961_v21)
			var _975_v33 D1
			_ = _975_v33
			_975_v33 = Companion_D1_.Create_DC4_(_974_v32)
			var _976_v34 D1
			_ = _976_v34
			_976_v34 = Companion_D1_.Create_DC4_(_974_v32)
			var _977_v35 _dafny.Map
			_ = _977_v35
			_977_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _976_v34)
			var _978_v36 _dafny.Map
			_ = _978_v36
			_978_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_973_v31).Cardinality(), _977_v35)
			var _979_v37 _dafny.Map
			_ = _979_v37
			_979_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_dafny.Zero).Minus((_978_v36).Cardinality()))
			var _980_v38 _dafny.Array
			_ = _980_v38
			var _nwElement0_30 _dafny.CodePoint = _dafny.CodePoint('t')
			_ = _nwElement0_30
			var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(27))
			_ = _nw158
			(_nw158).ArraySet1CodePoint(_nwElement0_30, 0)
			(_nw158).ArraySet1CodePoint(_958_v18, 1)
			(_nw158).ArraySet1CodePoint(_958_v18, 2)
			(_nw158).ArraySet1CodePoint(_958_v18, 3)
			(_nw158).ArraySet1CodePoint(_958_v18, 4)
			(_nw158).ArraySet1CodePoint(_958_v18, 5)
			(_nw158).ArraySet1CodePoint(_958_v18, 6)
			(_nw158).ArraySet1CodePoint(_958_v18, 7)
			(_nw158).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_959_v19).Contains(_dafny.IntOfInt64(125)) {
					return (_959_v19).Get(_dafny.IntOfInt64(125)).(_dafny.CodePoint)
				}
				return _958_v18
			})(), 8)
			(_nw158).ArraySet1CodePoint(_958_v18, 9)
			(_nw158).ArraySet1CodePoint(_dafny.CodePoint('o'), 10)
			(_nw158).ArraySet1CodePoint(_dafny.CodePoint('b'), 11)
			(_nw158).ArraySet1CodePoint(_958_v18, 12)
			(_nw158).ArraySet1CodePoint(_958_v18, 13)
			(_nw158).ArraySet1CodePoint(_958_v18, 14)
			(_nw158).ArraySet1CodePoint(_958_v18, 15)
			(_nw158).ArraySet1CodePoint(_958_v18, 16)
			(_nw158).ArraySet1CodePoint(_958_v18, 17)
			(_nw158).ArraySet1CodePoint(_958_v18, 18)
			(_nw158).ArraySet1CodePoint(_958_v18, 19)
			(_nw158).ArraySet1CodePoint(Companion_Default___.Fm48(true, Companion_Default___.Fm1(_this.F28(), (_this).F34(), globalState), _dafny.IntOfUint32(((_this).F32()).Cardinality()), globalState), 20)
			(_nw158).ArraySet1CodePoint(_958_v18, 21)
			(_nw158).ArraySet1CodePoint(_958_v18, 22)
			(_nw158).ArraySet1CodePoint(_958_v18, 23)
			(_nw158).ArraySet1CodePoint(_dafny.CodePoint('l'), 24)
			(_nw158).ArraySet1CodePoint(_958_v18, 25)
			(_nw158).ArraySet1CodePoint(_958_v18, 26)
			_980_v38 = _nw158
			var _981_v39 *C0
			_ = _981_v39
			var _nw159 *C0 = New_C0_()
			_ = _nw159
			_nw159.Ctor__(_980_v38)
			_981_v39 = _nw159
			var _982_v40 D1
			_ = _982_v40
			_982_v40 = Companion_D1_.Create_DC3_(_this.F35(), (_this).F30(), _981_v39, _dafny.IntOfUint32(((_this).F32()).Cardinality()))
			var _983_v41 *C3
			_ = _983_v41
			var _nw160 *C3 = New_C3_()
			_ = _nw160
			_nw160.Ctor__((_this).F34(), _971_v29, _958_v18, (_this).F34(), !(_this.F35()), _dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), _this.F29()), (_this).F33(), (func() _dafny.Int {
				if (_979_v37).Contains(_this.F29()) {
					return (_979_v37).Get(_this.F29()).(_dafny.Int)
				}
				return (_this).F30()
			})(), (_982_v40).Dtor_cf3(), !(_this.F29()), _this.F31())
			_983_v41 = _nw160
		} else {
			var _984_v42 _dafny.Set
			_ = _984_v42
			_984_v42 = _dafny.SetOf(_this.F31())
			var _985_v43 _dafny.Set
			_ = _985_v43
			_985_v43 = _984_v42
			var _986_v44 _dafny.CodePoint
			_ = _986_v44
			_986_v44 = _dafny.CodePoint('k')
			var _987_v45 _dafny.Sequence
			_ = _987_v45
			_987_v45 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F34()))
			var _988_v46 D4
			_ = _988_v46
			_988_v46 = Companion_D4_.Create_DC10_(Companion_Default___.Fm49(_dafny.MultiSetOf(_dafny.IntOfInt64(977), (_985_v43).Cardinality()), _986_v44, _dafny.SeqOf(_dafny.SeqOf((_dafny.Zero).Minus((_this).F34())), _987_v45, _987_v45), globalState))
			_988_v46 = _988_v46
			var _989_v47 _dafny.Array
			_ = _989_v47
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw161
			_989_v47 = _nw161
			(globalState).F21 = _989_v47
			(globalState).F1 = _this.F35()
			var _990_v48 _dafny.Sequence
			_ = _990_v48
			_990_v48 = _dafny.UnicodeSeqOfUtf8Bytes("hngyxcapn")
			var _991_v49 _dafny.Array
			_ = _991_v49
			var _nwElement0_31 _dafny.CodePoint = _986_v44
			_ = _nwElement0_31
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(3))
			_ = _nw162
			(_nw162).ArraySet1CodePoint(_nwElement0_31, 0)
			(_nw162).ArraySet1CodePoint(_986_v44, 1)
			(_nw162).ArraySet1CodePoint((_990_v48).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_990_v48).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
			_991_v49 = _nw162
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_991_v49), 0))
			_ = _index121
			(_991_v49).ArraySet1CodePoint(_986_v44, (_index121).Int())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(724), _dafny.ArrayLen((_991_v49), 0))
			_ = _index122
			(_991_v49).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_this).Fm6(globalState) {
					return _986_v44
				}
				return _986_v44
			})(), (_index122).Int())
			var _992_v50 _dafny.Map
			_ = _992_v50
			_992_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_989_v47, !(true))
			(_this).F28_set_(((_992_v50).Cardinality()).Cmp((_dafny.Zero).Minus((_this).F34())) != 0)
		}
		var _993_v51 _dafny.Set
		_ = _993_v51
		_993_v51 = _dafny.SetOf((_this).F30())
		var _994_v52 _dafny.MultiSet
		_ = _994_v52
		_994_v52 = _dafny.MultiSetOf((_this).F34(), (_this).F30())
		var _995_v53 _dafny.Map
		_ = _995_v53
		_995_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))
		var _rhs130 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F30(), (_this).F30())
		_ = _rhs130
		var _rhs131 _dafny.Int = Companion_Default___.SafeModuloInt(((_994_v52).Intersection(_994_v52)).Cardinality(), (func() _dafny.Int {
			if _this.F29() {
				return (_995_v53).Cardinality()
			}
			return (_this).F30()
		})())
		_ = _rhs131
		var _rhs132 _dafny.Set = (_993_v51).Intersection(_993_v51)
		_ = _rhs132
		var _lhs96 *GlobalState = globalState
		_ = _lhs96
		r1 = _rhs130
		_lhs96.F13 = _rhs131
		_993_v51 = _rhs132
		var _996_v54 _dafny.Sequence
		_ = _996_v54
		_996_v54 = _dafny.UnicodeSeqOfUtf8Bytes("cmjcaq")
		var _997_v55 _dafny.Set
		_ = _997_v55
		_997_v55 = _dafny.SetOf(_996_v54)
		var _998_v56 _dafny.CodePoint
		_ = _998_v56
		_998_v56 = _dafny.CodePoint('h')
		(globalState).F13 = (((_997_v55).Intersection(_dafny.SetOf(_996_v54, _996_v54, _dafny.Companion_Sequence_.Update(_996_v54, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F34()), _dafny.IntOfUint32((_996_v54).Cardinality()))).Uint32(), _998_v56)))).Difference((_997_v55).Union(_997_v55))).Cardinality()
		var _999_v57 _dafny.Map
		_ = _999_v57
		_999_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F31())
		r0 = !((func() bool {
			if (_999_v57).Contains(((_this).F30()).Plus((_this).F34())) {
				return (_999_v57).Get(((_this).F30()).Plus((_this).F34())).(bool)
			}
			return _this.F35()
		})())
		r1 = (_this).F30()
		var _1000_v58 D5
		_ = _1000_v58
		_1000_v58 = Companion_D5_.Create_DC12_(_dafny.UnicodeSeqOfUtf8Bytes("rpyunkfoa"))
		r2 = _dafny.Companion_Sequence_.Concatenate(_996_v54, (_1000_v58).Dtor_cf26())
		return r0, r1, r2
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
	_f47 bool
	_f48 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
	_this._f47 = false
	_this._f48 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T3 = &C5{}
var _ T2 = &C5{}
var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F28() bool {
	return _this._f28
}
func (_this *C5) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C5) F29() bool {
	return _this._f29
}
func (_this *C5) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C5) F31() bool {
	return _this._f31
}
func (_this *C5) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C5) F35() bool {
	return _this._f35
}
func (_this *C5) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C5) F34() _dafny.Int {
	return _this._f34
}
func (_this *C5) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C5) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C5) F30() _dafny.Int {
	return _this._f30
}
func (_this *C5) Ctor__(f47 bool, f48 _dafny.Int, f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f47 = f47
		(_this)._f48 = f48
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C5) Fm6(globalState *GlobalState) bool {
	{
		return ((_this).F30()).Cmp((_this).F34()) > 0
	}
}
func (_this *C5) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return (func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter27 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-735))).Elements()); ; {
				_compr_25, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _1001_v0 _dafny.Int
				_1001_v0 = interface{}(_compr_25).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-735)), _1001_v0) {
					_coll25.Add((_1001_v0).Minus(_dafny.IntOfInt64(599)))
				}
			}
			return _coll25.ToSet()
		}()).Union((_dafny.SetOf((_this).F48())).Difference(_dafny.SetOf(_dafny.IntOfInt64(255))))
	}
}
func (_this *C5) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return (_this).F32()
	}
}
func (_this *C5) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(741))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg74 _dafny.Int) interface{} {
				return coer74(arg74)
			}
		}(func(_1002_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), _dafny.UnicodeSeqOfUtf8Bytes("bfcfykdh"))
	}
}
func (_this *C5) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F48())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F48()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F30()))
	}
}
func (_this *C5) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		r1 = (_this).F48()
		var _1003_v0 _dafny.CodePoint
		_ = _1003_v0
		_1003_v0 = _dafny.CodePoint('a')
		var _1004_v1 _dafny.MultiSet
		_ = _1004_v1
		_1004_v1 = _dafny.MultiSetOf(_1003_v0, _1003_v0)
		var _1005_v2 _dafny.MultiSet
		_ = _1005_v2
		_1005_v2 = _dafny.MultiSetOf((_this).F48())
		var _1006_v3 _dafny.Map
		_ = _1006_v3
		_1006_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1004_v1, Companion_Default___.Fm39((_this).F47(), _1005_v2, _this.F31(), globalState))
		var _1007_v4 _dafny.Map
		_ = _1007_v4
		_1007_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1006_v3, (_this).F47())
		_1007_v4 = (_1007_v4).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1004_v1, (_this).F34()), !((_this).Fm6(globalState)))
		var _1008_v5 _dafny.Sequence
		_ = _1008_v5
		_1008_v5 = _dafny.SeqOf((_this).F48())
		var _1009_v6 _dafny.Map
		_ = _1009_v6
		_1009_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F28()), _dafny.Companion_Sequence_.IsPrefixOf(_1008_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg75 _dafny.Int) interface{} {
				return coer75(arg75)
			}
		}(func(_1010_i0 _dafny.Int) _dafny.Int {
			return (_this).F48()
		}))))
		var _1011_v7 _dafny.Sequence
		_ = _1011_v7
		_1011_v7 = _dafny.SeqOf(((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool))
		var _1012_v8 _dafny.Array
		_ = _1012_v8
		var _nw163 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(13))
		_ = _nw163
		_1012_v8 = _nw163
		var _1013_v9 D4
		_ = _1013_v9
		_1013_v9 = Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(248))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg76 _dafny.Int) interface{} {
				return coer76(arg76)
			}
		}((func(_1014_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1015_i1 _dafny.Int) _dafny.CodePoint {
				return _1014_v0
			}
		})(_1003_v0)))).Cardinality()), true, (_this).F30())
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1012_v8), 0))
		_ = _index123
		(_1012_v8).ArraySet1(_1013_v9, (_index123).Int())
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1012_v8), 0))
		_ = _index124
		var _rhs133 _dafny.Map = _1009_v6
		_ = _rhs133
		var _rhs134 _dafny.Sequence = _1011_v7
		_ = _rhs134
		var _rhs135 D4 = Companion_D4_.Create_DC11_((_this).F30(), ((_this).F34()).Cmp((_this).F48()) >= 0, _dafny.IntOfInt64(-352))
		_ = _rhs135
		var _rhs136 bool = !(_this.F35()) || (((_this).F34()).Cmp((_this).F30()) == 0)
		_ = _rhs136
		var _lhs97 _dafny.Array = _1012_v8
		_ = _lhs97
		var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1012_v8), 0))
		_ = _lhs98
		var _lhs99 *C5 = _this
		_ = _lhs99
		_1009_v6 = _rhs133
		_1011_v7 = _rhs134
		(_lhs97).ArraySet1(_rhs135, (_lhs98).Int())
		_lhs99.F35_set_(_rhs136)
		var _1016_v10 _dafny.Map
		_ = _1016_v10
		_1016_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(941))
		var _1017_v11 _dafny.MultiSet
		_ = _1017_v11
		_1017_v11 = _dafny.MultiSetOf(_this.F29())
		var _1018_v12 _dafny.Map
		_ = _1018_v12
		_1018_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1003_v0)
		var _1019_v13 _dafny.Array
		_ = _1019_v13
		var _nwElement0_32 _dafny.Int = (_this).F30()
		_ = _nwElement0_32
		var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(29))
		_ = _nw164
		(_nw164).ArraySet1(_nwElement0_32, 0)
		(_nw164).ArraySet1((_this).F30(), 1)
		(_nw164).ArraySet1((_this).F48(), 2)
		(_nw164).ArraySet1((_1008_v5).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_1005_v2).Contains((_this).F34()) {
				return (_1005_v2).Multiplicity((_this).F34())
			}
			return (_this).F30()
		})(), _dafny.IntOfUint32((_1008_v5).Cardinality()))).Uint32()).(_dafny.Int), 3)
		(_nw164).ArraySet1((_this).F30(), 4)
		(_nw164).ArraySet1((_this).F30(), 5)
		(_nw164).ArraySet1((_this).F34(), 6)
		(_nw164).ArraySet1((_this).F34(), 7)
		(_nw164).ArraySet1(_dafny.IntOfInt64(381), 8)
		(_nw164).ArraySet1((_this).F30(), 9)
		(_nw164).ArraySet1((_dafny.Zero).Minus((_this).F34()), 10)
		(_nw164).ArraySet1((_this).F48(), 11)
		(_nw164).ArraySet1((_this).F48(), 12)
		(_nw164).ArraySet1((_this).F48(), 13)
		(_nw164).ArraySet1(_dafny.IntOfUint32((_1008_v5).Cardinality()), 14)
		(_nw164).ArraySet1(_dafny.IntOfInt64(803), 15)
		(_nw164).ArraySet1((_this).F34(), 16)
		(_nw164).ArraySet1((_this).F34(), 17)
		(_nw164).ArraySet1((_this).F48(), 18)
		(_nw164).ArraySet1((_dafny.MultiSetOf(_this.F29(), false, Companion_Default___.Fm1(_this.F28(), (func() _dafny.Int {
			if (_1016_v10).Contains(_this.F28()) {
				return (_1016_v10).Get(_this.F28()).(_dafny.Int)
			}
			return (_this).F48()
		})(), globalState), _this.F35(), true)).Cardinality(), 19)
		(_nw164).ArraySet1((_this).F48(), 20)
		(_nw164).ArraySet1((_1017_v11).Cardinality(), 21)
		(_nw164).ArraySet1((_this).F34(), 22)
		(_nw164).ArraySet1((_this).F48(), 23)
		(_nw164).ArraySet1((_this).F48(), 24)
		(_nw164).ArraySet1((_this).F34(), 25)
		(_nw164).ArraySet1((_1018_v12).Cardinality(), 26)
		(_nw164).ArraySet1((_this).F34(), 27)
		(_nw164).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(656))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg77 _dafny.Int) interface{} {
				return coer77(arg77)
			}
		}((func(_1020_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1021_i2 _dafny.Int) _dafny.CodePoint {
				return _1020_v0
			}
		})(_1003_v0)))).Cardinality()), 28)
		_1019_v13 = _nw164
		var _1022_v14 _dafny.Set
		_ = _1022_v14
		_1022_v14 = _dafny.SetOf(_1019_v13, _1019_v13, _1019_v13, _1019_v13, _1019_v13)
		_1022_v14 = _1022_v14
		var _1023_v15 D8
		_ = _1023_v15
		_1023_v15 = Companion_D8_.Create_DC22_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("djojhst")).Cardinality()), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(948), (_this).F48()))
		var _source18 D8 = _1023_v15
		_ = _source18
		if _source18.Is_DC22() {
			var _1024___mcc_h0 _dafny.Int = _source18.Get_().(D8_DC22).Cf41
			_ = _1024___mcc_h0
			var _1025___mcc_h1 _dafny.Int = _source18.Get_().(D8_DC22).Cf42
			_ = _1025___mcc_h1
			var _1026_cf42 _dafny.Int = _1025___mcc_h1
			_ = _1026_cf42
			var _1027_cf41 _dafny.Int = _1024___mcc_h0
			_ = _1027_cf41
			var _rhs137 bool = _this.F31()
			_ = _rhs137
			var _rhs138 _dafny.Sequence = (_this).Fm3(_this.F28(), _this.F29(), globalState)
			_ = _rhs138
			var _lhs100 *GlobalState = globalState
			_ = _lhs100
			_lhs100.F1 = _rhs137
			r2 = _rhs138
			(globalState).F13 = _1026_cf42
			if !(_this.F28()) {
				var _1028_v16 _dafny.Set
				_ = _1028_v16
				_1028_v16 = _dafny.SetOf(_1017_v11, _1017_v11)
				var _1029_v17 _dafny.Map
				_ = _1029_v17
				_1029_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1017_v11, (_this).F48())
				var _1030_v19 D8
				_ = _1030_v19
				_1030_v19 = Companion_D8_.Create_DC24_(true, (_1009_v6).Cardinality(), (_this).F32(), _this.F35(), _1008_v5)
				var _1031_v20 D8
				_ = _1031_v20
				_1031_v20 = Companion_D8_.Create_DC25_(_1030_v19)
				var _1032_v21 _dafny.Array
				_ = _1032_v21
				var _nwElement0_33 _dafny.Set = _1028_v16
				_ = _nwElement0_33
				var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(8))
				_ = _nw165
				(_nw165).ArraySet1(_nwElement0_33, 0)
				(_nw165).ArraySet1(func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter28 := _dafny.Iterate((_1029_v17).Keys().Elements()); ; {
						_compr_26, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _1033_v18 _dafny.MultiSet
						_1033_v18 = interface{}(_compr_26).(_dafny.MultiSet)
						if (_1029_v17).Contains(_1033_v18) {
							_coll26.Add(_1033_v18)
						}
					}
					return _coll26.ToSet()
				}(), 1)
				(_nw165).ArraySet1((_1028_v16).Intersection(Companion_Default___.Fm40(_1027_cf41, _this.F29(), _1031_v20, globalState)), 2)
				(_nw165).ArraySet1(_1028_v16, 3)
				(_nw165).ArraySet1(_1028_v16, 4)
				(_nw165).ArraySet1(_1028_v16, 5)
				(_nw165).ArraySet1((_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F31(), _this.F35(), _this.F35())))).Intersection(_1028_v16), 6)
				(_nw165).ArraySet1(_dafny.SetOf(_1017_v11, _dafny.MultiSetOf(false, _this.F28()), _1017_v11, _1017_v11), 7)
				_1032_v21 = _nw165
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1032_v21), 0))
				_ = _index125
				(_1032_v21).ArraySet1(_1028_v16, (_index125).Int())
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1032_v21), 0))
				_ = _index126
				(_1032_v21).ArraySet1(_1028_v16, (_index126).Int())
				(globalState).F13 = Companion_Default___.Fm39(_this.F35(), _1005_v2, (!(_this.F35())) || (_this.F29()), globalState)
				var _1034_v22 _dafny.Array
				_ = _1034_v22
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw166
				_1034_v22 = _nw166
				var _1035_v23 _dafny.Array
				_ = _1035_v23
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_20
				var _nw167 _dafny.Array
				_ = _nw167
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw167 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) bool = func(_1036_i3 _dafny.Int) bool {
						return _this.F31()
					}
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw167 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw167).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw167).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_1035_v23 = _nw167
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1034_v22), 0))
				_ = _index127
				(_1034_v22).ArraySet1(_1035_v23, (_index127).Int())
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1019_v13), 0))
				_ = _index128
				(_1019_v13).ArraySet1(_1027_cf41, (_index128).Int())
				var _1037_v24 _dafny.Array
				_ = _1037_v24
				var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
				_ = _nw168
				_1037_v24 = _nw168
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1034_v22), 0))
				_ = _index129
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1019_v13), 0))
				_ = _index130
				var _rhs139 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F47())).Update(_this.F35(), _this.F35())).Update(_this.F35(), (_this).F47())
				_ = _rhs139
				var _rhs140 _dafny.Array = _1035_v23
				_ = _rhs140
				var _rhs141 _dafny.Int = _1027_cf41
				_ = _rhs141
				var _rhs142 _dafny.Array = _1037_v24
				_ = _rhs142
				var _rhs143 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32())
				_ = _rhs143
				var _lhs101 _dafny.Array = _1034_v22
				_ = _lhs101
				var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_1034_v22), 0))
				_ = _lhs102
				var _lhs103 _dafny.Array = _1019_v13
				_ = _lhs103
				var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1019_v13), 0))
				_ = _lhs104
				_1009_v6 = _rhs139
				(_lhs101).ArraySet1(_rhs140, (_lhs102).Int())
				(_lhs103).ArraySet1(_rhs141, (_lhs104).Int())
				_1037_v24 = _rhs142
				_1011_v7 = _rhs143
				var _1038_v25 _dafny.Array
				_ = _1038_v25
				var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
				_ = _nw169
				_1038_v25 = _nw169
				var _1039_v26 _dafny.Map
				_ = _1039_v26
				_1039_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(183), _this.F35())
				var _1040_v27 _dafny.Sequence
				_ = _1040_v27
				_1040_v27 = _dafny.SeqOf(_1003_v0)
				var _1041_v28 *C1
				_ = _1041_v28
				var _nw170 *C1 = New_C1_()
				_ = _nw170
				_nw170.Ctor__(_this.F29(), _this.F31(), _this.F35())
				_1041_v28 = _nw170
				var _1042_v29 _dafny.Set
				_ = _1042_v29
				_1042_v29 = _dafny.SetOf((_this).F48(), _dafny.IntOfInt64(696), _1027_cf41, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1041_v28, Companion_Default___.Fm41(_this.F31(), _1003_v0, _this.F28(), (_this).F48(), globalState))).Cardinality())
				var _1043_v30 _dafny.Array
				_ = _1043_v30
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_21
				var _nw171 _dafny.Array
				_ = _nw171
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw171 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.CodePoint = (func(_1044_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1045_i4 _dafny.Int) _dafny.CodePoint {
							return _1044_v0
						}
					})(_1003_v0)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw171 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw171).ArraySet1CodePoint(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw171).ArraySet1CodePoint(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_1043_v30 = _nw171
				var _1046_v31 *C0
				_ = _1046_v31
				var _nw172 *C0 = New_C0_()
				_ = _nw172
				_nw172.Ctor__(_1043_v30)
				_1046_v31 = _nw172
				var _1047_v32 _dafny.Map
				_ = _1047_v32
				_1047_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC27_(_1046_v31, _dafny.UnicodeSeqOfUtf8Bytes("dxwqwtkoe")), _1035_v23)
				var _nwElement0_34 _dafny.Int = _1027_cf41
				_ = _nwElement0_34
				var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(22))
				_ = _nw173
				(_nw173).ArraySet1(_nwElement0_34, 0)
				(_nw173).ArraySet1(_1026_cf42, 1)
				(_nw173).ArraySet1(((_this).F30()).Plus(_1026_cf42), 2)
				(_nw173).ArraySet1(Companion_Default___.Fm39((func() bool {
					if (_1039_v26).Contains((_1039_v26).Cardinality()) {
						return (_1039_v26).Get((_1039_v26).Cardinality()).(bool)
					}
					return _this.F28()
				})(), _1005_v2, _this.F31(), globalState), 3)
				(_nw173).ArraySet1(Companion_Default___.SafeModuloInt(_1026_cf42, _dafny.IntOfUint32((_1040_v27).Cardinality())), 4)
				(_nw173).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1008_v5, _dafny.SeqOf((_1019_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1019_v13), 0))).Int()).(_dafny.Int)))).Cardinality()), 5)
				(_nw173).ArraySet1((_this).F30(), 6)
				(_nw173).ArraySet1((Companion_Default___.Fm39(false, _1005_v2, _this.F29(), globalState)).Times((_this).F48()), 7)
				(_nw173).ArraySet1((_this).F30(), 8)
				(_nw173).ArraySet1(_dafny.IntOfInt64(-634), 9)
				(_nw173).ArraySet1(_1026_cf42, 10)
				(_nw173).ArraySet1(_1027_cf41, 11)
				(_nw173).ArraySet1((_this).F34(), 12)
				(_nw173).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F30(), ((_this).Fm4(Companion_Default___.Fm1(_this.F35(), (_1019_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1019_v13), 0))).Int()).(_dafny.Int), globalState), _1042_v29, _1026_cf42, _this.F28(), globalState)).Cardinality()), 13)
				(_nw173).ArraySet1(Companion_Default___.Fm39(_this.F31(), _1005_v2, !((_1041_v28).F43()), globalState), 14)
				(_nw173).ArraySet1((_this).F30(), 15)
				(_nw173).ArraySet1(((_1047_v32).Merge(_1047_v32)).Cardinality(), 16)
				(_nw173).ArraySet1((_this).F48(), 17)
				(_nw173).ArraySet1((_this).F30(), 18)
				(_nw173).ArraySet1((_this).F34(), 19)
				(_nw173).ArraySet1((_1019_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1019_v13), 0))).Int()).(_dafny.Int), 20)
				(_nw173).ArraySet1((_this).F34(), 21)
				_1038_v25 = _nw173
				_1026_cf42 = Companion_Default___.Fm39(_this.F29(), _dafny.MultiSetFromSeq(Companion_Default___.Fm42(_1040_v27, (_this).F48(), _dafny.UnicodeSeqOfUtf8Bytes("gkp"), globalState)), (_1041_v28).F43(), globalState)
			} else {
				(_this).F35_set_(((_this).F30()).Cmp(Companion_Default___.Fm39((_this).Fm6(globalState), _1005_v2, (_this).F47(), globalState)) == 0)
				_1026_cf42 = _dafny.IntOfInt64(441)
				r1 = (func() _dafny.Int {
					if !(_this.F28()) {
						return (_dafny.Zero).Minus(((func() _dafny.Int {
							if (_1017_v11).Contains(!(_this.F29())) {
								return (_1017_v11).Multiplicity(!(_this.F29()))
							}
							return _dafny.IntOfUint32((_1008_v5).Cardinality())
						})()).Plus(_dafny.IntOfInt64(545)))
					}
					return (_this).F30()
				})()
				var _1048_v33 _dafny.Array
				_ = _1048_v33
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
				_ = _nw174
				_1048_v33 = _nw174
				var _1049_v34 _dafny.Set
				_ = _1049_v34
				_1049_v34 = _dafny.SetOf((_this).F48())
				var _1050_v35 _dafny.Map
				_ = _1050_v35
				_1050_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm41(_this.F35(), _1003_v0, _this.F29(), (_this).F34(), globalState), (_1005_v2).Update((_dafny.Zero).Minus((_this).F30()), Companion_Default___.Abs((_1049_v34).Cardinality())))
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1048_v33), 0))
				_ = _index131
				(_1048_v33).ArraySet1((func() _dafny.MultiSet {
					if (_1050_v35).Contains(Companion_Default___.Fm41(false, _1003_v0, _this.F28(), Companion_Default___.Fm39(false, _1005_v2, _this.F35(), globalState), globalState)) {
						return (_1050_v35).Get(Companion_Default___.Fm41(false, _1003_v0, _this.F28(), Companion_Default___.Fm39(false, _1005_v2, _this.F35(), globalState), globalState)).(_dafny.MultiSet)
					}
					return _1005_v2
				})(), (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_1048_v33), 0))
				_ = _index132
				(_1048_v33).ArraySet1(Companion_Default___.Fm43(globalState), (_index132).Int())
				var _1051_v36 _dafny.Array
				_ = _1051_v36
				var _nwElement0_35 bool = _this.F35()
				_ = _nwElement0_35
				var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.One)
				_ = _nw175
				(_nw175).ArraySet1(_nwElement0_35, 0)
				_1051_v36 = _nw175
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1051_v36), 0))
				_ = _index133
				(_1051_v36).ArraySet1((_this).F47(), (_index133).Int())
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1051_v36), 0))
				_ = _index134
				(_1051_v36).ArraySet1(((_this).F30()).Cmp((_this).F48()) == 0, (_index134).Int())
			}
			var _1052_v37 _dafny.Map
			_ = _1052_v37
			_1052_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_this).F30())
			_1052_v37 = (_1052_v37).Update(_1027_cf41, (_this).F34())
		} else if _source18.Is_DC23() {
			var _1053___mcc_h2 bool = _source18.Get_().(D8_DC23).Cf43
			_ = _1053___mcc_h2
			var _1054___mcc_h3 bool = _source18.Get_().(D8_DC23).Cf44
			_ = _1054___mcc_h3
			var _1055___mcc_h4 bool = _source18.Get_().(D8_DC23).Cf45
			_ = _1055___mcc_h4
			var _1056_cf45 bool = _1055___mcc_h4
			_ = _1056_cf45
			var _1057_cf44 bool = _1054___mcc_h3
			_ = _1057_cf44
			var _1058_cf43 bool = _1053___mcc_h2
			_ = _1058_cf43
			_1003_v0 = _1003_v0
			var _1059_v38 _dafny.Sequence
			_ = _1059_v38
			_1059_v38 = _dafny.UnicodeSeqOfUtf8Bytes("nujp")
			(globalState).F1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_1059_v38, (Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32((_1059_v38).Cardinality()))).Uint32(), _1003_v0), _1059_v38)
			(_this).F35_set_(false)
			(globalState).F13 = (_this).F48()
		} else if _source18.Is_DC24() {
			var _1060___mcc_h5 bool = _source18.Get_().(D8_DC24).Cf46
			_ = _1060___mcc_h5
			var _1061___mcc_h6 _dafny.Int = _source18.Get_().(D8_DC24).Cf47
			_ = _1061___mcc_h6
			var _1062___mcc_h7 _dafny.Sequence = _source18.Get_().(D8_DC24).Cf48
			_ = _1062___mcc_h7
			var _1063___mcc_h8 bool = _source18.Get_().(D8_DC24).Cf49
			_ = _1063___mcc_h8
			var _1064___mcc_h9 _dafny.Sequence = _source18.Get_().(D8_DC24).Cf50
			_ = _1064___mcc_h9
			var _1065_cf50 _dafny.Sequence = _1064___mcc_h9
			_ = _1065_cf50
			var _1066_cf49 bool = _1063___mcc_h8
			_ = _1066_cf49
			var _1067_cf48 _dafny.Sequence = _1062___mcc_h7
			_ = _1067_cf48
			var _1068_cf47 _dafny.Int = _1061___mcc_h6
			_ = _1068_cf47
			var _1069_cf46 bool = _1060___mcc_h5
			_ = _1069_cf46
			_1005_v2 = (_1005_v2).Difference(_1005_v2)
			_1003_v0 = _dafny.CodePoint('i')
			var _1070_v39 _dafny.Sequence
			_ = _1070_v39
			_1070_v39 = _dafny.UnicodeSeqOfUtf8Bytes("vvj")
			var _1071_v40 _dafny.Array
			_ = _1071_v40
			var _nwElement0_36 _dafny.Sequence = _1070_v39
			_ = _nwElement0_36
			var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(5))
			_ = _nw176
			(_nw176).ArraySet1(_nwElement0_36, 0)
			(_nw176).ArraySet1(_1070_v39, 1)
			(_nw176).ArraySet1(_1070_v39, 2)
			(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1070_v39, _1070_v39), 3)
			(_nw176).ArraySet1(_dafny.Companion_Sequence_.Update(_1070_v39, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F34()), _dafny.IntOfUint32((_1070_v39).Cardinality()))).Uint32(), _1003_v0), 4)
			_1071_v40 = _nw176
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1071_v40), 0))
			_ = _index135
			(_1071_v40).ArraySet1(_1070_v39, (_index135).Int())
			var _1072_v41 _dafny.Sequence
			_ = _1072_v41
			_1072_v41 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(849))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg78 _dafny.Int) interface{} {
					return coer78(arg78)
				}
			}((func(_1073_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1074_i5 _dafny.Int) _dafny.CodePoint {
					return _1073_v0
				}
			})(_1003_v0))))
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1071_v40), 0))
			_ = _index136
			(_1071_v40).ArraySet1((_1072_v41).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1072_v41).Cardinality()))).Uint32()).(_dafny.Sequence), (_index136).Int())
			var _1075_v42 _dafny.Sequence
			_ = _1075_v42
			_1075_v42 = _dafny.SeqOf(_1008_v5, _1065_cf50, _1065_cf50)
			var _1076_v43 _dafny.Map
			_ = _1076_v43
			_1076_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1075_v42).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1068_cf47), _dafny.IntOfUint32((_1075_v42).Cardinality()))).Uint32()).(_dafny.Sequence), _1066_cf49)
			var _1077_v44 D6
			_ = _1077_v44
			_1077_v44 = Companion_D6_.Create_DC15_(_1008_v5)
			_1076_v43 = ((_1076_v43).Merge(_1076_v43)).Update((_1077_v44).Dtor_cf32(), _1066_cf49)
		} else if _source18.Is_DC21() {
			var _1078___mcc_h10 _dafny.Map = _source18.Get_().(D8_DC21).Cf40
			_ = _1078___mcc_h10
			var _1079_cf40 _dafny.Map = _1078___mcc_h10
			_ = _1079_cf40
			var _1080_v45 *C2
			_ = _1080_v45
			var _nw177 *C2 = New_C2_()
			_ = _nw177
			_nw177.Ctor__((_this).F48(), (_this).F30(), _this.F31(), _this.F28(), _this.F35())
			_1080_v45 = _nw177
			var _1081_v46 _dafny.Array
			_ = _1081_v46
			var _nwElement0_37 bool = (_this).F47()
			_ = _nwElement0_37
			var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(8))
			_ = _nw178
			(_nw178).ArraySet1(_nwElement0_37, 0)
			(_nw178).ArraySet1(_this.F35(), 1)
			(_nw178).ArraySet1(_this.F29(), 2)
			(_nw178).ArraySet1(true, 3)
			(_nw178).ArraySet1(_this.F35(), 4)
			(_nw178).ArraySet1(_this.F35(), 5)
			(_nw178).ArraySet1(_this.F28(), 6)
			(_nw178).ArraySet1(_this.F35(), 7)
			_1081_v46 = _nw178
			var _1082_v47 D3
			_ = _1082_v47
			_1082_v47 = Companion_D3_.Create_DC9_((_1080_v45).Fm30(_this.F35(), globalState), (_this).Fm6(globalState), _1081_v46, (_1011_v7).Select((Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32((_1011_v7).Cardinality()))).Uint32()).(bool))
			var _1083_v48 D6
			_ = _1083_v48
			_1083_v48 = Companion_D6_.Create_DC15_(_1008_v5)
			var _1084_v49 _dafny.Sequence
			_ = _1084_v49
			_1084_v49 = _dafny.SeqOf(_1008_v5)
			var _1085_v50 D8
			_ = _1085_v50
			_1085_v50 = Companion_D8_.Create_DC24_(_this.F35(), (_this).F34(), (_this).F32(), _this.F28(), (_1084_v49).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1084_v49).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _pat_let_tv10 = _1085_v50
			_ = _pat_let_tv10
			var _1086_v51 bool
			_ = _1086_v51
			var _1087_v52 _dafny.CodePoint
			_ = _1087_v52
			var _out54 bool
			_ = _out54
			var _out55 _dafny.CodePoint
			_ = _out55
			_out54, _out55 = (_this).M14(_1003_v0, (_1082_v47).Dtor_cf20(), func(_pat_let13_0 D6) D6 {
				return func(_1088_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let14_0 _dafny.Sequence) D6 {
						return func(_1089_dt__update_hcf32_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC15_(_1089_dt__update_hcf32_h0)
						}(_pat_let14_0)
					}((_pat_let_tv10).Dtor_cf50())
				}(_pat_let13_0)
			}(_1083_v48), globalState)
			_1086_v51 = _out54
			_1087_v52 = _out55
			(_this).F29_set_(_this.F29())
			(globalState).F13 = (_this).F34()
		} else {
			var _1090___mcc_h11 D8 = _source18.Get_().(D8_DC25).Cf51
			_ = _1090___mcc_h11
			var _1091_cf51 D8 = _1090___mcc_h11
			_ = _1091_cf51
			var _1092_v53 _dafny.Map
			_ = _1092_v53
			_1092_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F34())
			var _1093_v54 _dafny.Sequence
			_ = _1093_v54
			_1093_v54 = _dafny.UnicodeSeqOfUtf8Bytes("e")
			var _1094_v55 _dafny.Sequence
			_ = _1094_v55
			_1094_v55 = _dafny.SeqOf(_1092_v53)
			_1092_v53 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1093_v54).Cardinality()), _dafny.IntOfUint32((_1093_v54).Cardinality()))).Merge((_1094_v55).Select((Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32((_1094_v55).Cardinality()))).Uint32()).(_dafny.Map))
			(globalState).F13 = _dafny.IntOfInt64(342)
			var _source19 D8 = _1023_v15
			_ = _source19
			if _source19.Is_DC22() {
				var _1095___mcc_h12 _dafny.Int = _source19.Get_().(D8_DC22).Cf41
				_ = _1095___mcc_h12
				var _1096___mcc_h13 _dafny.Int = _source19.Get_().(D8_DC22).Cf42
				_ = _1096___mcc_h13
				var _1097_cf42 _dafny.Int = _1096___mcc_h13
				_ = _1097_cf42
				var _1098_cf41 _dafny.Int = _1095___mcc_h12
				_ = _1098_cf41
				var _1099_v56 _dafny.Map
				_ = _1099_v56
				_1099_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F30()).Minus(_1097_cf42), _this.F28())
				_1099_v56 = (_1099_v56).Update((_dafny.IntOfInt64(991)).Times((_this).F34()), ((_this).F34()).Cmp(_1097_cf42) == 0)
				var _1100_v58 _dafny.Map
				_ = _1100_v58
				_1100_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _1099_v56)
				var _1101_v59 _dafny.Array
				_ = _1101_v59
				var _nwElement0_38 _dafny.Map = func() _dafny.Map {
					var _coll27 = _dafny.NewMapBuilder()
					_ = _coll27
					for _iter29 := _dafny.Iterate((_1008_v5).Elements()); ; {
						_compr_27, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _1102_v57 _dafny.Int
						_1102_v57 = interface{}(_compr_27).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1008_v5, _1102_v57) {
							_coll27.Add(Companion_Default___.SafeDivisionInt(_1102_v57, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_dafny.SetOf(_1008_v5)).Cardinality())).Cardinality()), _this.F31())
						}
					}
					return _coll27.ToMap()
				}()
				_ = _nwElement0_38
				var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(10))
				_ = _nw179
				(_nw179).ArraySet1(_nwElement0_38, 0)
				(_nw179).ArraySet1(_1099_v56, 1)
				(_nw179).ArraySet1(_1099_v56, 2)
				(_nw179).ArraySet1(_1099_v56, 3)
				(_nw179).ArraySet1((func() _dafny.Map {
					if (_1100_v58).Contains((_this).F47()) {
						return (_1100_v58).Get((_this).F47()).(_dafny.Map)
					}
					return _1099_v56
				})(), 4)
				(_nw179).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _this.F35()), 5)
				(_nw179).ArraySet1(_1099_v56, 6)
				(_nw179).ArraySet1(_1099_v56, 7)
				(_nw179).ArraySet1(_1099_v56, 8)
				(_nw179).ArraySet1(_1099_v56, 9)
				_1101_v59 = _nw179
				var _1103_v60 _dafny.Array
				_ = _1103_v60
				var _nwElement0_39 _dafny.Array = _1101_v59
				_ = _nwElement0_39
				var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(9))
				_ = _nw180
				(_nw180).ArraySet1(_nwElement0_39, 0)
				(_nw180).ArraySet1(_1101_v59, 1)
				(_nw180).ArraySet1(_1101_v59, 2)
				(_nw180).ArraySet1(_1101_v59, 3)
				(_nw180).ArraySet1(_1101_v59, 4)
				(_nw180).ArraySet1(_1101_v59, 5)
				(_nw180).ArraySet1(_1101_v59, 6)
				(_nw180).ArraySet1((func() _dafny.Array {
					if _this.F35() {
						return _1101_v59
					}
					return _1101_v59
				})(), 7)
				(_nw180).ArraySet1(_1101_v59, 8)
				_1103_v60 = _nw180
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_1103_v60), 0))
				_ = _index137
				(_1103_v60).ArraySet1(_1101_v59, (_index137).Int())
				var _1104_v61 _dafny.Set
				_ = _1104_v61
				_1104_v61 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _dafny.IntOfInt64(-133))).Cardinality(), (_1009_v6).Cardinality(), (_this).F48(), _1098_cf41)
				var _1105_v63 _dafny.Array
				_ = _1105_v63
				var _nwElement0_40 bool = _this.F28()
				_ = _nwElement0_40
				var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(17))
				_ = _nw181
				(_nw181).ArraySet1(_nwElement0_40, 0)
				(_nw181).ArraySet1((func() _dafny.Set {
					var _coll28 = _dafny.NewBuilder()
					_ = _coll28
					for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(57), _dafny.IntOfInt64(761))); ; {
						_compr_28, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _1106_v62 _dafny.Int
						_1106_v62 = interface{}(_compr_28).(_dafny.Int)
						if ((_dafny.IntOfInt64(57)).Cmp(_1106_v62) <= 0) && ((_1106_v62).Cmp(_dafny.IntOfInt64(761)) < 0) {
							_coll28.Add((_1106_v62).Times((_this).F30()))
						}
					}
					return _coll28.ToSet()
				}()).IsProperSubsetOf(_1104_v61), 1)
				(_nw181).ArraySet1((!((_this).F47())) || ((_this).F47()), 2)
				(_nw181).ArraySet1((Companion_Default___.Fm1(_this.F35(), _1098_cf41, globalState)) || (_this.F28()), 3)
				(_nw181).ArraySet1(false, 4)
				(_nw181).ArraySet1(_this.F28(), 5)
				(_nw181).ArraySet1(_this.F31(), 6)
				(_nw181).ArraySet1((_this).F47(), 7)
				(_nw181).ArraySet1(false, 8)
				(_nw181).ArraySet1((_this).F47(), 9)
				(_nw181).ArraySet1(_this.F35(), 10)
				(_nw181).ArraySet1((_this).Fm6(globalState), 11)
				(_nw181).ArraySet1(_this.F28(), 12)
				(_nw181).ArraySet1(!((_dafny.MultiSetFromSeq(_1011_v7)).Update(!(_this.F31()), Companion_Default___.Abs(_1097_cf42))).Contains(_this.F29()), 13)
				(_nw181).ArraySet1(_this.F29(), 14)
				(_nw181).ArraySet1(_this.F35(), 15)
				(_nw181).ArraySet1(_this.F31(), 16)
				_1105_v63 = _nw181
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1105_v63), 0))
				_ = _index138
				(_1105_v63).ArraySet1(_this.F29(), (_index138).Int())
				var _1107_v64 _dafny.Sequence
				_ = _1107_v64
				_1107_v64 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(889))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}((func(_1108_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1109_i6 _dafny.Int) _dafny.CodePoint {
						return _1108_v0
					}
				})(_1003_v0))))
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_1103_v60), 0))
				_ = _index139
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1105_v63), 0))
				_ = _index140
				var _rhs144 _dafny.Int = _dafny.IntOfInt64(788)
				_ = _rhs144
				var _rhs145 _dafny.Array = _1101_v59
				_ = _rhs145
				var _rhs146 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1107_v64).Select((Companion_Default___.SafeIndex(_1098_cf41, _dafny.IntOfUint32((_1107_v64).Cardinality()))).Uint32()).(_dafny.Sequence), _1093_v54)).Cardinality())).Cmp((_1017_v11).Cardinality()) <= 0
				_ = _rhs146
				var _lhs105 *GlobalState = globalState
				_ = _lhs105
				var _lhs106 _dafny.Array = _1103_v60
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_1103_v60), 0))
				_ = _lhs107
				var _lhs108 _dafny.Array = _1105_v63
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_1105_v63), 0))
				_ = _lhs109
				_lhs105.F13 = _rhs144
				(_lhs106).ArraySet1(_rhs145, (_lhs107).Int())
				(_lhs108).ArraySet1(_rhs146, (_lhs109).Int())
				(globalState).F13 = (_this).F34()
				var _1110_v65 _dafny.Array
				_ = _1110_v65
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_22
				var _nw182 _dafny.Array
				_ = _nw182
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw182 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.MultiSet = (func(_1111_v11 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_1112_i7 _dafny.Int) _dafny.MultiSet {
							return _1111_v11
						}
					})(_1017_v11)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw182 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw182).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw182).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_1110_v65 = _nw182
				var _1113_v66 _dafny.Map
				_ = _1113_v66
				_1113_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1110_v65, _1098_cf41)
				_1113_v66 = (_1113_v66).Update(_1110_v65, (_this).F34())
			} else if _source19.Is_DC23() {
				var _1114___mcc_h14 bool = _source19.Get_().(D8_DC23).Cf43
				_ = _1114___mcc_h14
				var _1115___mcc_h15 bool = _source19.Get_().(D8_DC23).Cf44
				_ = _1115___mcc_h15
				var _1116___mcc_h16 bool = _source19.Get_().(D8_DC23).Cf45
				_ = _1116___mcc_h16
				var _1117_cf45 bool = _1116___mcc_h16
				_ = _1117_cf45
				var _1118_cf44 bool = _1115___mcc_h15
				_ = _1118_cf44
				var _1119_cf43 bool = _1114___mcc_h14
				_ = _1119_cf43
				(globalState).F13 = (_this).F34()
				var _1120_v67 _dafny.Sequence
				_ = _1120_v67
				_1120_v67 = _dafny.SeqOf(_1093_v54, _dafny.Companion_Sequence_.Concatenate(_1093_v54, _1093_v54))
				_1120_v67 = _dafny.Companion_Sequence_.Concatenate(_1120_v67, _1120_v67)
				var _1121_v68 _dafny.Set
				_ = _1121_v68
				_1121_v68 = _dafny.SetOf(_1117_cf45, (_this).F47())
				var _1122_v69 *C1
				_ = _1122_v69
				var _nw183 *C1 = New_C1_()
				_ = _nw183
				_nw183.Ctor__(((_1121_v68).Cardinality()).Cmp((_this).F30()) == 0, _this.F28(), _this.F31())
				_1122_v69 = _nw183
				_1019_v13 = _1019_v13
			} else if _source19.Is_DC24() {
				var _1123___mcc_h17 bool = _source19.Get_().(D8_DC24).Cf46
				_ = _1123___mcc_h17
				var _1124___mcc_h18 _dafny.Int = _source19.Get_().(D8_DC24).Cf47
				_ = _1124___mcc_h18
				var _1125___mcc_h19 _dafny.Sequence = _source19.Get_().(D8_DC24).Cf48
				_ = _1125___mcc_h19
				var _1126___mcc_h20 bool = _source19.Get_().(D8_DC24).Cf49
				_ = _1126___mcc_h20
				var _1127___mcc_h21 _dafny.Sequence = _source19.Get_().(D8_DC24).Cf50
				_ = _1127___mcc_h21
				var _1128_cf50 _dafny.Sequence = _1127___mcc_h21
				_ = _1128_cf50
				var _1129_cf49 bool = _1126___mcc_h20
				_ = _1129_cf49
				var _1130_cf48 _dafny.Sequence = _1125___mcc_h19
				_ = _1130_cf48
				var _1131_cf47 _dafny.Int = _1124___mcc_h18
				_ = _1131_cf47
				var _1132_cf46 bool = _1123___mcc_h17
				_ = _1132_cf46
				(globalState).F13 = (Companion_Default___.SafeModuloInt((_this).F48(), (_this).F34())).Minus((_this).F34())
				var _1133_v70 D3
				_ = _1133_v70
				_1133_v70 = Companion_D3_.Create_DC8_(_dafny.MultiSetOf(_1003_v0), (_this).F47(), _dafny.IntOfInt64(735))
				var _pat_let_tv11 = _1129_cf49
				_ = _pat_let_tv11
				_1133_v70 = func(_pat_let15_0 D3) D3 {
					return func(_1134_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let16_0 bool) D3 {
							return func(_1135_dt__update_hcf16_h0 bool) D3 {
								return Companion_D3_.Create_DC8_((_1134_dt__update__tmp_h1).Dtor_cf15(), _1135_dt__update_hcf16_h0, (_1134_dt__update__tmp_h1).Dtor_cf17())
							}(_pat_let16_0)
						}(_pat_let_tv11)
					}(_pat_let15_0)
				}(_1133_v70)
				var _1136_v71 _dafny.Array
				_ = _1136_v71
				var _len0_23 _dafny.Int = _dafny.One
				_ = _len0_23
				var _nw184 _dafny.Array
				_ = _nw184
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw184 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Map = func(_1137_i8 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F31())
					}
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw184 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw184).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw184).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1136_v71 = _nw184
				var _1138_v72 _dafny.Map
				_ = _1138_v72
				_1138_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _1129_cf49)
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_1136_v71), 0))
				_ = _index141
				(_1136_v71).ArraySet1((_1138_v72).Merge(_1138_v72), (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_1136_v71), 0))
				_ = _index142
				(_1136_v71).ArraySet1((func() _dafny.Map {
					if (func() bool {
						if _this.F29() {
							return true
						}
						return ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)
					})() {
						return _1138_v72
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F47())
				})(), (_index142).Int())
				var _1139_v74 _dafny.Set
				_ = _1139_v74
				_1139_v74 = _dafny.SetOf((_dafny.Zero).Minus((func() _dafny.Set {
					var _coll29 = _dafny.NewBuilder()
					_ = _coll29
					for _iter31 := _dafny.Iterate(((_1136_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_1136_v71), 0))).Int()).(_dafny.Map)).Keys().Elements()); ; {
						_compr_29, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _1140_v73 _dafny.Int
						_1140_v73 = interface{}(_compr_29).(_dafny.Int)
						if ((_1136_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_1136_v71), 0))).Int()).(_dafny.Map)).Contains(_1140_v73) {
							_coll29.Add((_1140_v73).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hmcvtul")).Cardinality())))
						}
					}
					return _coll29.ToSet()
				}()).Cardinality()), (_this).F30())
				var _1141_v76 _dafny.Sequence
				_ = _1141_v76
				_1141_v76 = _dafny.SeqOf(_1139_v74)
				var _1142_v77 _dafny.Array
				_ = _1142_v77
				var _nwElement0_41 _dafny.Set = _1139_v74
				_ = _nwElement0_41
				var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(13))
				_ = _nw185
				(_nw185).ArraySet1(_nwElement0_41, 0)
				(_nw185).ArraySet1((_1139_v74).Difference(_dafny.SetOf((_this).F34(), (_this).F34())), 1)
				(_nw185).ArraySet1(_dafny.SetOf((_dafny.Zero).Minus((_this).F30()), (_this).F30()), 2)
				(_nw185).ArraySet1(func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(718), _dafny.IntOfInt64(-585))); ; {
						_compr_30, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _1143_v75 _dafny.Int
						_1143_v75 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(718)).Cmp(_1143_v75) <= 0) && ((_1143_v75).Cmp(_dafny.IntOfInt64(-585)) < 0) {
							_coll30.Add(Companion_Default___.SafeModuloInt(_1143_v75, _dafny.IntOfInt64(470)))
						}
					}
					return _coll30.ToSet()
				}(), 3)
				(_nw185).ArraySet1(_1139_v74, 4)
				(_nw185).ArraySet1((_1139_v74).Difference(_1139_v74), 5)
				(_nw185).ArraySet1((_1139_v74).Difference(_1139_v74), 6)
				(_nw185).ArraySet1((_1141_v76).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1128_cf50).Cardinality()), _dafny.IntOfUint32((_1141_v76).Cardinality()))).Uint32()).(_dafny.Set), 7)
				(_nw185).ArraySet1(_1139_v74, 8)
				(_nw185).ArraySet1(_1139_v74, 9)
				(_nw185).ArraySet1(_1139_v74, 10)
				(_nw185).ArraySet1(_1139_v74, 11)
				(_nw185).ArraySet1(_1139_v74, 12)
				_1142_v77 = _nw185
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1142_v77), 0))
				_ = _index143
				(_1142_v77).ArraySet1(_dafny.SetOf(_1131_cf47, (_this).F30()), (_index143).Int())
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1142_v77), 0))
				_ = _index144
				(_1142_v77).ArraySet1((_this).Fm4(_1129_cf49, func() _dafny.Set {
					var _coll31 = _dafny.NewBuilder()
					_ = _coll31
					for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-346), _dafny.IntOfInt64(278))); ; {
						_compr_31, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _1144_v78 _dafny.Int
						_1144_v78 = interface{}(_compr_31).(_dafny.Int)
						if ((_dafny.IntOfInt64(-346)).Cmp(_1144_v78) <= 0) && ((_1144_v78).Cmp(_dafny.IntOfInt64(278)) < 0) {
							_coll31.Add(Companion_Default___.SafeDivisionInt(_1144_v78, _1131_cf47))
						}
					}
					return _coll31.ToSet()
				}(), (_this).F34(), _1132_cf46, globalState), (_index144).Int())
			} else if _source19.Is_DC21() {
				var _1145___mcc_h22 _dafny.Map = _source19.Get_().(D8_DC21).Cf40
				_ = _1145___mcc_h22
				var _1146_cf40 _dafny.Map = _1145___mcc_h22
				_ = _1146_cf40
				var _1147_v79 _dafny.Map
				_ = _1147_v79
				_1147_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_1005_v2).Union(_dafny.MultiSetOf((_this).F30())))
				var _rhs147 _dafny.MultiSet = (func() _dafny.MultiSet {
					if (_1147_v79).Contains((_this.F28()) == (_this.F29())) {
						return (_1147_v79).Get((_this.F28()) == (_this.F29())).(_dafny.MultiSet)
					}
					return (_1005_v2).Update((_this).F48(), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gonoudi")).Cardinality())))
				})()
				_ = _rhs147
				var _rhs148 bool = (_this).F47()
				_ = _rhs148
				var _rhs149 _dafny.Int = ((_this).F48()).Minus((_this).F30())
				_ = _rhs149
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				_1005_v2 = _rhs147
				_lhs110.F12 = _rhs148
				_lhs111.F13 = _rhs149
				var _1148_v80 _dafny.Array
				_ = _1148_v80
				var _nwElement0_42 bool = false
				_ = _nwElement0_42
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(6))
				_ = _nw186
				(_nw186).ArraySet1(_nwElement0_42, 0)
				(_nw186).ArraySet1(((_this).F47()) == (_this.F29()), 1)
				(_nw186).ArraySet1(((_this).F30()).Cmp((_this).F48()) > 0, 2)
				(_nw186).ArraySet1(!(_this.F28()), 3)
				(_nw186).ArraySet1(true, 4)
				(_nw186).ArraySet1(_this.F28(), 5)
				_1148_v80 = _nw186
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1148_v80), 0))
				_ = _index145
				(_1148_v80).ArraySet1((func() bool {
					if _this.F31() {
						return _this.F31()
					}
					return _this.F29()
				})(), (_index145).Int())
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1148_v80), 0))
				_ = _index146
				(_1148_v80).ArraySet1(false, (_index146).Int())
				var _1149_v81 _dafny.MultiSet
				_ = _1149_v81
				_1149_v81 = _dafny.MultiSetOf((_1012_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_1012_v8), 0))).Int()).(D4))
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1148_v80), 0))
				_ = _index147
				var _rhs150 _dafny.Array = _1148_v80
				_ = _rhs150
				var _rhs151 bool = !((((func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter34 := _dafny.Iterate((_1008_v5).Elements()); ; {
						_compr_32, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _1150_v82 _dafny.Int
						_1150_v82 = interface{}(_compr_32).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1008_v5, _1150_v82) {
							_coll32.Add(Companion_Default___.SafeModuloInt(_1150_v82, (_this).F30()), _1093_v54)
						}
					}
					return _coll32.ToMap()
				}()).Cardinality()).Minus(_dafny.IntOfUint32((_1093_v54).Cardinality()))).Cmp(_dafny.IntOfUint32((_1011_v7).Cardinality())) >= 0)
				_ = _rhs151
				var _rhs152 _dafny.MultiSet = _dafny.MultiSetOf(Companion_D4_.Create_DC11_((_this).F30(), (_1148_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1148_v80), 0))).Int()).(bool), (_this).F30()))
				_ = _rhs152
				var _lhs112 _dafny.Array = _1148_v80
				_ = _lhs112
				var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1148_v80), 0))
				_ = _lhs113
				_1148_v80 = _rhs150
				(_lhs112).ArraySet1(_rhs151, (_lhs113).Int())
				_1149_v81 = _rhs152
				(_this).F29_set_(false)
			} else {
				var _1151___mcc_h23 D8 = _source19.Get_().(D8_DC25).Cf51
				_ = _1151___mcc_h23
				var _1152_cf51 D8 = _1151___mcc_h23
				_ = _1152_cf51
				(_this).F35_set_(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1011_v7, (_this).F32()), _dafny.Companion_Sequence_.Concatenate((_this).Fm5((_this).F48(), (_this).F48(), (_this).F47(), (_this).F48(), globalState), (_this).F32())))
				var _1153_v83 *C1
				_ = _1153_v83
				var _nw187 *C1 = New_C1_()
				_ = _nw187
				_nw187.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_1093_v54, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(314))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}(func(_1154_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))), ((_this).F34()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tsm")).Cardinality())) == 0, _this.F31())
				_1153_v83 = _nw187
				(globalState).F26 = (_1153_v83).F43()
				var _1155_v84 _dafny.Array
				_ = _1155_v84
				var _nw188 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw188
				_1155_v84 = _nw188
				(globalState).F21 = _1155_v84
			}
			var _rhs153 _dafny.Int = Companion_Default___.Fm39(_this.F28(), _1005_v2, !(_this.F31()), globalState)
			_ = _rhs153
			var _rhs154 bool = ((_this).F30()).Cmp((_this).F48()) < 0
			_ = _rhs154
			var _rhs155 bool = (_dafny.IntOfInt64(358)).Cmp((func() _dafny.Int {
				if (_1017_v11).Contains(_this.F28()) {
					return (_1017_v11).Multiplicity(_this.F28())
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1008_v5, (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1008_v5).Cardinality()))).Uint32(), (_this).F48())).Cardinality())
			})()) == 0
			_ = _rhs155
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			var _lhs115 *C5 = _this
			_ = _lhs115
			var _lhs116 *GlobalState = globalState
			_ = _lhs116
			_lhs114.F13 = _rhs153
			_lhs115.F29_set_(_rhs154)
			_lhs116.F12 = _rhs155
		}
		var _hi5 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-182))
		_ = _hi5
		for _1156_i10 := ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _1005_v2)).Cardinality()).Minus((_dafny.Zero).Minus((_this).F48())); _1156_i10.Cmp(_hi5) < 0; _1156_i10 = _1156_i10.Plus(_dafny.One) {
			(globalState).F12 = Companion_Default___.Fm1(_this.F35(), (_1005_v2).Cardinality(), globalState)
			(globalState).F12 = ((_dafny.IntOfInt64(169)).Cmp(_1156_i10) != 0) && (_this.F35())
			(globalState).F13 = ((_this).F34()).Times((_this).F48())
			var _1157_v85 _dafny.Map
			_ = _1157_v85
			_1157_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _1017_v11)
			_1157_v85 = (_1157_v85).Update(_this.F28(), _dafny.MultiSetOf(_this.F31()))
		}
		var _1158_v86 _dafny.Sequence
		_ = _1158_v86
		_1158_v86 = _dafny.UnicodeSeqOfUtf8Bytes("jjtojxk")
		r0 = (_dafny.IntOfUint32((_1158_v86).Cardinality())).Cmp((_this).F34()) > 0
		r1 = (func() _dafny.Int {
			if !(_this.F28()) {
				return ((_this).F34()).Plus((_this).F30())
			}
			return (_dafny.Zero).Minus((_this).F34())
		})()
		r2 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("l"), _1158_v86), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("l"), _1158_v86)).Cardinality()))).Uint32(), _1003_v0)
		return r0, r1, r2
	}
}
func (_this *C5) M14(p0 _dafny.CodePoint, p1 _dafny.Array, p2 D6, globalState *GlobalState) (bool, _dafny.CodePoint) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r1
		var _hi6 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bxgmww")).Cardinality())
		_ = _hi6
		for _1159_i0 := (_this).F30(); _1159_i0.Cmp(_hi6) < 0; _1159_i0 = _1159_i0.Plus(_dafny.One) {
			(globalState).F12 = _this.F28()
			(globalState).F13 = Companion_Default___.SafeDivisionInt(_1159_i0, (_1159_i0).Minus(_1159_i0))
			var _1160_v0 _dafny.Map
			_ = _1160_v0
			_1160_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(73), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(137))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}((func(_1161_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1162_i1 _dafny.Int) _dafny.CodePoint {
					return _1161_p0
				}
			})(p0)))).Cardinality()))
			var _1163_v1 T3
			_ = _1163_v1
			var _nw189 *C4 = New_C4_()
			_ = _nw189
			_nw189.Ctor__((_this).F30(), _this.F29(), (_this).F32(), (_this).F33(), (_1160_v0).Cardinality(), _this.F35(), Companion_Default___.Fm1(false, (_this).F48(), globalState), (_this).F47())
			_1163_v1 = _nw189
			var _1164_v2 _dafny.Map
			_ = _1164_v2
			_1164_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1159_i0, _1163_v1)
			var _1165_v3 D10
			_ = _1165_v3
			_1165_v3 = Companion_D10_.Create_DC30_(Companion_Default___.Fm41(_this.F31(), p0, _this.F31(), (_this).F34(), globalState), _dafny.MultiSetOf(_1163_v1, _1163_v1, (func() T3 {
				if (_1164_v2).Contains(_dafny.IntOfInt64(764)) {
					return (_1164_v2).Get(_dafny.IntOfInt64(764)).(T3)
				}
				return _1163_v1
			})(), _1163_v1))
			var _1166_v4 _dafny.MultiSet
			_ = _1166_v4
			_1166_v4 = _dafny.MultiSetOf(_1163_v1, _1163_v1, _1163_v1, _1163_v1, _1163_v1)
			var _pat_let_tv12 = _1166_v4
			_ = _pat_let_tv12
			var _source20 D10 = func(_pat_let17_0 D10) D10 {
				return func(_1167_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let18_0 _dafny.MultiSet) D10 {
						return func(_1168_dt__update_hcf59_h0 _dafny.MultiSet) D10 {
							return Companion_D10_.Create_DC30_((_1167_dt__update__tmp_h0).Dtor_cf58(), _1168_dt__update_hcf59_h0)
						}(_pat_let18_0)
					}(_pat_let_tv12)
				}(_pat_let17_0)
			}(_1165_v3)
			_ = _source20
			if _source20.Is_DC29() {
				var _1169___mcc_h0 _dafny.Int = _source20.Get_().(D10_DC29).Cf56
				_ = _1169___mcc_h0
				var _1170___mcc_h1 bool = _source20.Get_().(D10_DC29).Cf57
				_ = _1170___mcc_h1
				var _1171_cf57 bool = _1170___mcc_h1
				_ = _1171_cf57
				var _1172_cf56 _dafny.Int = _1169___mcc_h0
				_ = _1172_cf56
				(globalState).F13 = (_1172_cf56).Minus(((_this).F34()).Times((_this).F34()))
				var _1173_v5 _dafny.Set
				_ = _1173_v5
				_1173_v5 = _dafny.SetOf(p0, p0, p0)
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((p1), 0))
				_ = _index148
				(p1).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1163_v1.F31(), _1173_v5)).Contains(_this.F31()), (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((p1), 0))
				_ = _index149
				(p1).ArraySet1(_1163_v1.F29(), (_index149).Int())
				var _1174_v6 _dafny.Sequence
				_ = _1174_v6
				_1174_v6 = _dafny.UnicodeSeqOfUtf8Bytes("xji")
				var _1175_v7 _dafny.Map
				_ = _1175_v7
				_1175_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1174_v6, (_this).F30())
				_1175_v7 = _1175_v7
				var _1176_v8 _dafny.Array
				_ = _1176_v8
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_24
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Sequence = (func(_1177_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_1178_i2 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(315))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg82 _dafny.Int) interface{} {
									return coer82(arg82)
								}
							}((func(_1179_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1180_i3 _dafny.Int) _dafny.CodePoint {
									return _1179_p0
								}
							})(_1177_p0)))
						}
					})(p0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw190 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw190).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw190).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1176_v8 = _nw190
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1176_v8), 0))
				_ = _index150
				(_1176_v8).ArraySet1(_1174_v6, (_index150).Int())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_1176_v8), 0))
				_ = _index151
				(_1176_v8).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(822))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1181_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1182_i4 _dafny.Int) _dafny.CodePoint {
						return _1181_p0
					}
				})(p0))), (_index151).Int())
			} else if _source20.Is_DC30() {
				var _1183___mcc_h2 _dafny.CodePoint = _source20.Get_().(D10_DC30).Cf58
				_ = _1183___mcc_h2
				var _1184___mcc_h3 _dafny.MultiSet = _source20.Get_().(D10_DC30).Cf59
				_ = _1184___mcc_h3
				var _1185_cf59 _dafny.MultiSet = _1184___mcc_h3
				_ = _1185_cf59
				var _1186_cf58 _dafny.CodePoint = _1183___mcc_h2
				_ = _1186_cf58
				(globalState).F13 = (_this).F48()
				var _1187_v9 _dafny.Array
				_ = _1187_v9
				var _nwElement0_43 _dafny.Int = (_1163_v1).F34()
				_ = _nwElement0_43
				var _nw191 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(4))
				_ = _nw191
				(_nw191).ArraySet1(_nwElement0_43, 0)
				(_nw191).ArraySet1((_1163_v1).F34(), 1)
				(_nw191).ArraySet1((_this).F34(), 2)
				(_nw191).ArraySet1((_1163_v1).F34(), 3)
				_1187_v9 = _nw191
				var _1188_v10 _dafny.Map
				_ = _1188_v10
				_1188_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1187_v9, _1163_v1.F28())
				var _1189_v11 D7
				_ = _1189_v11
				_1189_v11 = Companion_D7_.Create_DC19_(_1163_v1.F31(), (_dafny.Zero).Minus((_1163_v1).F34()), (_1188_v10).Cardinality())
				(globalState).F13 = (_1189_v11).Dtor_cf38()
				var _1190_v12 _dafny.Sequence
				_ = _1190_v12
				_1190_v12 = _dafny.UnicodeSeqOfUtf8Bytes("dplmqrvf")
				var _1191_v13 _dafny.Map
				_ = _1191_v13
				_1191_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F48(), _1190_v12)
				var _1192_v14 _dafny.Set
				_ = _1192_v14
				_1192_v14 = _dafny.SetOf(_this.F31())
				var _1193_v16 _dafny.Map
				_ = _1193_v16
				_1193_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F28())
				var _1194_v17 _dafny.Sequence
				_ = _1194_v17
				_1194_v17 = _dafny.SeqOf(_dafny.IntOfUint32(((_1163_v1).F32()).Cardinality()), (_1193_v16).Cardinality())
				var _1195_v18 _dafny.Array
				_ = _1195_v18
				var _nwElement0_44 _dafny.Sequence = (_1163_v1).F32()
				_ = _nwElement0_44
				var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(26))
				_ = _nw192
				(_nw192).ArraySet1(_nwElement0_44, 0)
				(_nw192).ArraySet1((_1163_v1).F32(), 1)
				(_nw192).ArraySet1((_1163_v1).F32(), 2)
				(_nw192).ArraySet1((_1163_v1).F32(), 3)
				(_nw192).ArraySet1((_1163_v1).F32(), 4)
				(_nw192).ArraySet1((_1163_v1).Fm5(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1191_v13).Contains((_this).F48()) {
						return (_1191_v13).Get((_this).F48()).(_dafny.Sequence)
					}
					return _1190_v12
				})()).Cardinality()), (_this).F34(), (_this).F47(), (_1163_v1).F30(), globalState), 5)
				(_nw192).ArraySet1((_1163_v1).F32(), 6)
				(_nw192).ArraySet1((_1163_v1).F32(), 7)
				(_nw192).ArraySet1((_this).F32(), 8)
				(_nw192).ArraySet1(_dafny.SeqOf(_1163_v1.F28(), _1163_v1.F29()), 9)
				(_nw192).ArraySet1((_1163_v1).F32(), 10)
				(_nw192).ArraySet1((_1163_v1).F32(), 11)
				(_nw192).ArraySet1((_1163_v1).F32(), 12)
				(_nw192).ArraySet1((_1163_v1).F32(), 13)
				(_nw192).ArraySet1((_1163_v1).F32(), 14)
				(_nw192).ArraySet1((_1163_v1).Fm5((_1192_v14).Cardinality(), (_this).F30(), _1163_v1.F29(), (func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter35 := _dafny.Iterate((_1194_v17).Elements()); ; {
						_compr_33, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _1196_v15 _dafny.Int
						_1196_v15 = interface{}(_compr_33).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1194_v17, _1196_v15) {
							_coll33.Add(Companion_Default___.SafeModuloInt(_1196_v15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())).Cardinality()), _1163_v1.F28())
						}
					}
					return _coll33.ToMap()
				}()).Cardinality(), globalState), 15)
				(_nw192).ArraySet1(_dafny.SeqOf(_this.F29()), 16)
				(_nw192).ArraySet1((_this).F32(), 17)
				(_nw192).ArraySet1(_dafny.SeqOf(_1163_v1.F35()), 18)
				(_nw192).ArraySet1((_1163_v1).F32(), 19)
				(_nw192).ArraySet1((_1163_v1).F32(), 20)
				(_nw192).ArraySet1((_this).F32(), 21)
				(_nw192).ArraySet1((_1163_v1).F32(), 22)
				(_nw192).ArraySet1((_1163_v1).F32(), 23)
				(_nw192).ArraySet1(_dafny.SeqOf((_this).F47(), _1163_v1.F35(), _this.F28()), 24)
				(_nw192).ArraySet1((_1163_v1).F32(), 25)
				_1195_v18 = _nw192
				var _1197_v19 T4
				_ = _1197_v19
				var _nw193 *C3 = New_C3_()
				_ = _nw193
				_nw193.Ctor__(_dafny.IntOfInt64(893), _1195_v18, p0, (Companion_Default___.Fm46(_dafny.IntOfUint32((_1190_v12).Cardinality()), (_this).F34(), _this.F31(), globalState)).Times((_1163_v1).F34()), !(!(((_this).F32()).Select((Companion_Default___.SafeIndex((_1163_v1).F30(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), (_1163_v1).F32()), _dafny.Companion_Sequence_.Concatenate((_1163_v1).F33(), (_this).F33()), (_1163_v1).F30(), !((func() bool {
					if _1163_v1.F35() {
						return _1163_v1.F31()
					}
					return _this.F35()
				})()), (func() bool {
					if _1163_v1.F31() {
						return true
					}
					return false
				})(), true)
				_1197_v19 = _nw193
				_1197_v19 = _1197_v19
				var _1198_v20 _dafny.Map
				_ = _1198_v20
				_1198_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}((func(_1199_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1200_i5 _dafny.Int) _dafny.CodePoint {
						return _1199_p0
					}
				})(p0))), _1192_v14)
				_1198_v20 = (_1198_v20).Update(_dafny.Companion_Sequence_.Concatenate(_1190_v12, _dafny.Companion_Sequence_.Update((_this).Fm3((_this).F47(), _1163_v1.F28(), globalState), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1160_v0).Contains((_1160_v0).Cardinality()) {
						return (_1160_v0).Get((_1160_v0).Cardinality()).(_dafny.Int)
					}
					return _dafny.IntOfUint32(((_this).F32()).Cardinality())
				})(), _dafny.IntOfUint32(((_this).Fm3((_this).F47(), _1163_v1.F28(), globalState)).Cardinality()))).Uint32(), _1197_v19.F36())), _1192_v14)
			} else if _source20.Is_DC31() {
				var _1201___mcc_h4 bool = _source20.Get_().(D10_DC31).Cf60
				_ = _1201___mcc_h4
				var _1202___mcc_h5 bool = _source20.Get_().(D10_DC31).Cf61
				_ = _1202___mcc_h5
				var _1203___mcc_h6 _dafny.Int = _source20.Get_().(D10_DC31).Cf62
				_ = _1203___mcc_h6
				var _1204_cf62 _dafny.Int = _1203___mcc_h6
				_ = _1204_cf62
				var _1205_cf61 bool = _1202___mcc_h5
				_ = _1205_cf61
				var _1206_cf60 bool = _1201___mcc_h4
				_ = _1206_cf60
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((p1), 0))
				_ = _index152
				(p1).ArraySet1(((_1163_v1).F32()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.IntOfUint32(((_1163_v1).F32()).Cardinality()))).Uint32()).(bool), (_index152).Int())
				var _1207_v21 D7
				_ = _1207_v21
				_1207_v21 = Companion_D7_.Create_DC19_((_this).F47(), (_1163_v1).F34(), (_1163_v1).F34())
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((p1), 0))
				_ = _index153
				(p1).ArraySet1((_1207_v21).Dtor_cf36(), (_index153).Int())
				var _1208_v22 _dafny.Array
				_ = _1208_v22
				var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw194
				_1208_v22 = _nw194
				_1208_v22 = _1208_v22
				var _1209_v23 *C1
				_ = _1209_v23
				var _nw195 *C1 = New_C1_()
				_ = _nw195
				_nw195.Ctor__(false, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((p1), 0))).Int()).(bool), false)
				_1209_v23 = _nw195
				var _1210_v24 D10
				_ = _1210_v24
				_1210_v24 = Companion_D10_.Create_DC28_(_1209_v23)
				var _1211_v25 _dafny.MultiSet
				_ = _1211_v25
				_1211_v25 = _dafny.MultiSetOf(_1210_v24)
				var _1212_v26 _dafny.Set
				_ = _1212_v26
				_1212_v26 = _dafny.SetOf((_1211_v25).Difference(_1211_v25))
				_1212_v26 = _1212_v26
				(globalState).F13 = _1159_i0
			} else if _source20.Is_DC28() {
				var _1213___mcc_h7 *C1 = _source20.Get_().(D10_DC28).Cf55
				_ = _1213___mcc_h7
				var _1214_cf55 *C1 = _1213___mcc_h7
				_ = _1214_cf55
				r1 = p0
				var _1215_v27 _dafny.Array
				_ = _1215_v27
				var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw196
				_1215_v27 = _nw196
				var _1216_v28 _dafny.Map
				_ = _1216_v28
				_1216_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1215_v27, (_this.F28()) || (_this.F31()))
				_1216_v28 = (_1216_v28)
				var _nw197 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw197
				_1215_v27 = _nw197
				(globalState).F12 = !((_1214_cf55).F43())
			} else {
				var _1217___mcc_h8 D10 = _source20.Get_().(D10_DC32).Cf63
				_ = _1217___mcc_h8
				var _1218_cf63 D10 = _1217___mcc_h8
				_ = _1218_cf63
				var _1219_v29 _dafny.Map
				_ = _1219_v29
				_1219_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _this.F28())
				var _1220_v30 _dafny.MultiSet
				_ = _1220_v30
				_1220_v30 = _dafny.MultiSetOf(((_1163_v1).F34()).Cmp((_this).F48()) >= 0, (func() bool {
					if (_1219_v29).Contains(_1163_v1.F28()) {
						return (_1219_v29).Get(_1163_v1.F28()).(bool)
					}
					return _1163_v1.F28()
				})(), (func() bool {
					if _1163_v1.F35() {
						return _this.F28()
					}
					return _1163_v1.F28()
				})(), _1163_v1.F28())
				(globalState).F13 = (_1220_v30).Cardinality()
				(_1163_v1).F35_set_(_1163_v1.F35())
				r1 = p0
				(globalState).F13 = Companion_Default___.Fm46((_1163_v1).F30(), (func() _dafny.Int {
					if _this.F31() {
						return (_this).F48()
					}
					return (_1163_v1).F30()
				})(), _1163_v1.F28(), globalState)
			}
			var _1221_v31 _dafny.MultiSet
			_ = _1221_v31
			_1221_v31 = _dafny.MultiSetOf(false)
			if (func() bool {
				if (_1221_v31).IsSubsetOf(_dafny.MultiSetOf(_this.F35(), (_this).F47())) {
					return _dafny.Companion_Sequence_.Equal((_this).F32(), (_this).F32())
				}
				return true
			})() {
				(globalState).F13 = (_this).F30()
				var _1222_v32 _dafny.Sequence
				_ = _1222_v32
				_1222_v32 = _dafny.SeqOf((_1163_v1).F34(), _1159_i0, (_1163_v1).F34())
				var _1223_v33 D8
				_ = _1223_v33
				_1223_v33 = Companion_D8_.Create_DC24_(_1163_v1.F35(), (_1163_v1).F30(), (_this).F32(), _1163_v1.F31(), _dafny.Companion_Sequence_.Update(_1222_v32, (Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32((_1222_v32).Cardinality()))).Uint32(), (_this).F48()))
				var _1224_v34 D0
				_ = _1224_v34
				_1224_v34 = Companion_D0_.Create_DC0_(_1163_v1.F31())
				var _1225_v35 *C2
				_ = _1225_v35
				var _nw198 *C2 = New_C2_()
				_ = _nw198
				_nw198.Ctor__((_1160_v0).Cardinality(), (_1223_v33).Dtor_cf47(), (_1224_v34).Dtor_cf0(), !(_1163_v1.F29()), _this.F29())
				_1225_v35 = _nw198
				var _1226_v36 _dafny.Map
				_ = _1226_v36
				_1226_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1163_v1.F29()) == (_this.F35())), _1163_v1.F35())
				_1226_v36 = (_1226_v36).Update((_this).F47(), (_this).F47())
				var _1227_v37 _dafny.Array
				_ = _1227_v37
				var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
				_ = _nw199
				_1227_v37 = _nw199
				var _1228_v38 _dafny.Map
				_ = _1228_v38
				_1228_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F34())
				var _1229_v39 _dafny.Set
				_ = _1229_v39
				_1229_v39 = _dafny.SetOf((_1163_v1).F30(), (_1163_v1).F34(), (_1228_v38).Cardinality(), (_this).F48(), (_1225_v35).F46())
				var _1230_v41 _dafny.Map
				_ = _1230_v41
				_1230_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(127), _dafny.IntOfInt64(328))); ; {
						_compr_34, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _1231_v40 _dafny.Int
						_1231_v40 = interface{}(_compr_34).(_dafny.Int)
						if ((_dafny.IntOfInt64(127)).Cmp(_1231_v40) <= 0) && ((_1231_v40).Cmp(_dafny.IntOfInt64(328)) < 0) {
							_coll34.Add(Companion_Default___.SafeModuloInt(_1231_v40, (_1163_v1).F34()))
						}
					}
					return _coll34.ToSet()
				}(), p1)
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1227_v37), 0))
				_ = _index154
				(_1227_v37).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1229_v39, p1)).Merge(_1230_v41), (_index154).Int())
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1227_v37), 0))
				_ = _index155
				(_1227_v37).ArraySet1(_1230_v41, (_index155).Int())
				var _1232_v42 *C1
				_ = _1232_v42
				var _nw200 *C1 = New_C1_()
				_ = _nw200
				_nw200.Ctor__(true, (func() bool {
					if (_1226_v36).Contains(_this.F35()) {
						return (_1226_v36).Get(_this.F35()).(bool)
					}
					return _this.F29()
				})(), _1163_v1.F35())
				_1232_v42 = _nw200
				var _1233_v43 _dafny.Map
				_ = _1233_v43
				_1233_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1163_v1).F34(), Companion_D10_.Create_DC28_(_1232_v42))
				var _1234_v44 _dafny.Sequence
				_ = _1234_v44
				_1234_v44 = _dafny.UnicodeSeqOfUtf8Bytes("mo")
				(_1163_v1).F29_set_(((_1233_v43).Cardinality()).Cmp(_dafny.IntOfUint32((_1234_v44).Cardinality())) <= 0)
			} else {
				(globalState).F26 = !(_1163_v1.F29())
				var _1235_v45 _dafny.Sequence
				_ = _1235_v45
				_1235_v45 = _dafny.UnicodeSeqOfUtf8Bytes("sqrfmscn")
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((p1), 0))
				_ = _index156
				(p1).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1235_v45, _1235_v45), (_index156).Int())
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((p1), 0))
				_ = _index157
				(p1).ArraySet1(_1163_v1.F31(), (_index157).Int())
				var _1236_v46 _dafny.Map
				_ = _1236_v46
				_1236_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1221_v31, p0)
				_1236_v46 = (_1236_v46).Update((_1221_v31).Update(!(_1163_v1.F31()), Companion_Default___.Abs((_1163_v1).F30())), (_1235_v45).Select((Companion_Default___.SafeIndex((_1163_v1).F34(), _dafny.IntOfUint32((_1235_v45).Cardinality()))).Uint32()).(_dafny.CodePoint))
				(globalState).F26 = _dafny.Companion_Sequence_.Contains((_this).F32(), Companion_Default___.Fm1(_1163_v1.F28(), (_1163_v1).F34(), globalState))
				(_this).F35_set_(_1163_v1.F28())
			}
		}
		if !(((_this).F32()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-172), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)) {
			r1 = p0
			(globalState).F13 = (_dafny.IntOfInt64(560)).Times((_this).F34())
			var _1237_v47 _dafny.Map
			_ = _1237_v47
			_1237_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F32())
			var _1238_v48 D3
			_ = _1238_v48
			_1238_v48 = Companion_D3_.Create_DC9_(_this.F35(), (_this).F47(), p1, _this.F28())
			var _1239_v49 _dafny.Array
			_ = _1239_v49
			var _nwElement0_45 _dafny.Sequence = (_this).F32()
			_ = _nwElement0_45
			var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(21))
			_ = _nw201
			(_nw201).ArraySet1(_nwElement0_45, 0)
			(_nw201).ArraySet1((_this).F32(), 1)
			(_nw201).ArraySet1((_this).F32(), 2)
			(_nw201).ArraySet1((_this).F32(), 3)
			(_nw201).ArraySet1((_this).F32(), 4)
			(_nw201).ArraySet1((func() _dafny.Sequence {
				if (_1237_v47).Contains((_1238_v48).Dtor_cf21()) {
					return (_1237_v47).Get((_1238_v48).Dtor_cf21()).(_dafny.Sequence)
				}
				return (_this).F32()
			})(), 5)
			(_nw201).ArraySet1((_this).F32(), 6)
			(_nw201).ArraySet1((_this).F32(), 7)
			(_nw201).ArraySet1(_dafny.SeqOf(false, (_this).F47(), _this.F35(), !(true)), 8)
			(_nw201).ArraySet1((_this).F32(), 9)
			(_nw201).ArraySet1((_this).F32(), 10)
			(_nw201).ArraySet1((_this).F32(), 11)
			(_nw201).ArraySet1((_this).F32(), 12)
			(_nw201).ArraySet1(_dafny.SeqOf(true, _this.F31()), 13)
			(_nw201).ArraySet1((_this).F32(), 14)
			(_nw201).ArraySet1((_this).F32(), 15)
			(_nw201).ArraySet1((_this).F32(), 16)
			(_nw201).ArraySet1((_this).F32(), 17)
			(_nw201).ArraySet1(_dafny.SeqOf(true, _this.F28(), _this.F28()), 18)
			(_nw201).ArraySet1((_this).F32(), 19)
			(_nw201).ArraySet1(_dafny.SeqOf((_this).F47()), 20)
			_1239_v49 = _nw201
			var _1240_v50 *C3
			_ = _1240_v50
			var _nw202 *C3 = New_C3_()
			_ = _nw202
			_nw202.Ctor__(Companion_Default___.SafeDivisionInt((_this).F30(), (_this).F30()), _1239_v49, p0, ((_this).F48()).Times((_this).F30()), ((_this).F47()) == (_this.F28()), (_this).F32(), (_this).F33(), (_this).F30(), false, (_this).F47(), _this.F35())
			_1240_v50 = _nw202
			(globalState).F12 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg85 _dafny.Int) interface{} {
					return coer85(arg85)
				}
			}(func(_1241_i6 _dafny.Int) _dafny.Int {
				return (_this).F30()
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(685))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}(func(_1242_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oaqupmrc")).Cardinality())
			})))
			var _1243_v51 _dafny.Sequence
			_ = _1243_v51
			_1243_v51 = _dafny.SeqOf((_1240_v50).F44())
			var _1244_v52 _dafny.MultiSet
			_ = _1244_v52
			_1244_v52 = _dafny.MultiSetOf(_1243_v51, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(618))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg87 _dafny.Int) interface{} {
					return coer87(arg87)
				}
			}((func(_1245_v50 *C3) func(_dafny.Int) _dafny.Int {
				return func(_1246_i8 _dafny.Int) _dafny.Int {
					return (_1245_v50).F44()
				}
			})(_1240_v50))), (p2).Dtor_cf32())
			var _1247_v53 _dafny.Map
			_ = _1247_v53
			_1247_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _1244_v52)
			var _1248_v54 _dafny.Set
			_ = _1248_v54
			_1248_v54 = _dafny.SetOf((_this).F34(), _dafny.IntOfUint32((_1243_v51).Cardinality()), (_this).F34())
			_1247_v53 = (_1247_v53).Update((_dafny.SetOf((_this).F34())).IsDisjointFrom(_1248_v54), _1244_v52)
		} else {
			var _1249_v55 _dafny.Map
			_ = _1249_v55
			_1249_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), ((_dafny.Zero).Minus((_this).F48())).Minus((_this).F30()))
			var _1250_v56 _dafny.Sequence
			_ = _1250_v56
			_1250_v56 = _dafny.UnicodeSeqOfUtf8Bytes("axjmgmig")
			_1249_v55 = (_1249_v55).Update(_dafny.Companion_Sequence_.IsPrefixOf(_1250_v56, _1250_v56), (_this).F30())
			var _1251_v57 _dafny.Map
			_ = _1251_v57
			_1251_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F48(), p0)
			_1251_v57 = (_1251_v57).Update((_this).F30(), p0)
			var _1252_v58 _dafny.Sequence
			_ = _1252_v58
			_1252_v58 = _dafny.SeqOf((_this).F48(), (_this).F48())
			var _1253_v59 _dafny.Map
			_ = _1253_v59
			_1253_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.SetOf(_this.F35())).Cardinality()).Minus((_this).F34()), _1252_v58)
			var _1254_v60 _dafny.Map
			_ = _1254_v60
			_1254_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F35())
			_1253_v59 = (_1253_v59).Update((_1254_v60).Cardinality(), _1252_v58)
			(globalState).F1 = !(_this.F35())
			var _1255_v61 _dafny.Sequence
			_ = _1255_v61
			_1255_v61 = _dafny.SeqOf(_1250_v56, _dafny.UnicodeSeqOfUtf8Bytes("yumismnte"))
			var _1256_v63 _dafny.Map
			_ = _1256_v63
			_1256_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.IntOfInt64(692))
			var _1257_v64 _dafny.MultiSet
			_ = _1257_v64
			_1257_v64 = _dafny.MultiSetOf((_this).F47())
			var _1258_v65 _dafny.Array
			_ = _1258_v65
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_25
			var _nw203 _dafny.Array
			_ = _nw203
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw203 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = func(_1259_i9 _dafny.Int) _dafny.Int {
					return (_1259_i9).Plus((Companion_D7_.Create_DC19_(_this.F35(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(146))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg88 _dafny.Int) interface{} {
							return coer88(arg88)
						}
					}(func(_1260_i10 _dafny.Int) _dafny.Int {
						return (_this).F30()
					}))).Cardinality()), (_this).F34())).Dtor_cf37())
				}
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw203 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw203).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw203).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1258_v65 = _nw203
			var _1261_v66 D0
			_ = _1261_v66
			_1261_v66 = Companion_D0_.Create_DC0_(_this.F35())
			var _1262_v67 _dafny.Sequence
			_ = _1262_v67
			_1262_v67 = _dafny.SeqOf(_1261_v66)
			var _1263_v68 _dafny.Set
			_ = _1263_v68
			_1263_v68 = _dafny.SetOf((_this).F48())
			var _1264_v69 D2
			_ = _1264_v69
			_1264_v69 = Companion_D2_.Create_DC6_((_dafny.Zero).Minus((_this).F30()), _this.F35(), (_this).F34(), (_this).F47(), (_this).F34())
			var _1265_v70 _dafny.MultiSet
			_ = _1265_v70
			_1265_v70 = _dafny.MultiSetOf((_this).F48(), (_this).F34())
			var _1266_v71 _dafny.Array
			_ = _1266_v71
			var _nwElement0_46 _dafny.Int = (((_this).Fm4(true, _dafny.SetOf((_this).F48(), (_this).F48()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1255_v61).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kuntapvl")).Cardinality()), _dafny.IntOfUint32((_1255_v61).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32(((_1255_v61).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kuntapvl")).Cardinality()), _dafny.IntOfUint32((_1255_v61).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.CodePoint('j'))).Cardinality()), _this.F28(), globalState)).Cardinality()).Plus((func() _dafny.Set {
				var _coll35 = _dafny.NewBuilder()
				_ = _coll35
				for _iter37 := _dafny.Iterate((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wa"))).Elements()); ; {
					_compr_35, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _1267_v62 _dafny.Sequence
					_1267_v62 = interface{}(_compr_35).(_dafny.Sequence)
					if (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wa"))).Contains(_1267_v62) {
						_coll35.Add(_1267_v62)
					}
				}
				return _coll35.ToSet()
			}()).Cardinality())
			_ = _nwElement0_46
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(29))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_46, 0)
			(_nw204).ArraySet1((func() _dafny.Int {
				if (_1256_v63).Contains((_this).F48()) {
					return (_1256_v63).Get((_this).F48()).(_dafny.Int)
				}
				return (_this).F30()
			})(), 1)
			(_nw204).ArraySet1(((_this).F48()).Minus((_this).F30()), 2)
			(_nw204).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_1252_v58).Cardinality()))).Times((func() _dafny.Int {
				if (_1257_v64).Contains((Companion_D10_.Create_DC31_(_this.F29(), (_this).F47(), (_this).F48())).Dtor_cf61()) {
					return (_1257_v64).Multiplicity((Companion_D10_.Create_DC31_(_this.F29(), (_this).F47(), (_this).F48())).Dtor_cf61())
				}
				return (_this).F34()
			})()), 3)
			(_nw204).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F34(), (_this).F30())), 4)
			(_nw204).ArraySet1((_this).F30(), 5)
			(_nw204).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F30(), (_this).F34()), 6)
			(_nw204).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1258_v65, _1258_v65, _1258_v65, _1258_v65)).Cardinality()), 7)
			(_nw204).ArraySet1(_dafny.IntOfUint32((_1262_v67).Cardinality()), 8)
			(_nw204).ArraySet1(Companion_Default___.SafeModuloInt((_this).F30(), (_this).F48()), 9)
			(_nw204).ArraySet1((_this).F48(), 10)
			(_nw204).ArraySet1(((_this).F34()).Minus((_this).F34()), 11)
			(_nw204).ArraySet1((_1263_v68).Cardinality(), 12)
			(_nw204).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F48()))).Cardinality())), (_this).F30()), 13)
			(_nw204).ArraySet1(((_1264_v69).Dtor_cf11()).Times((_dafny.Zero).Minus((_this).F30())), 14)
			(_nw204).ArraySet1((_this).F34(), 15)
			(_nw204).ArraySet1((_this).F34(), 16)
			(_nw204).ArraySet1((_this).F48(), 17)
			(_nw204).ArraySet1((_this).F30(), 18)
			(_nw204).ArraySet1((_this).F30(), 19)
			(_nw204).ArraySet1((_this).F30(), 20)
			(_nw204).ArraySet1((_this).F30(), 21)
			(_nw204).ArraySet1((_dafny.Zero).Minus((_this).F34()), 22)
			(_nw204).ArraySet1((_this).F48(), 23)
			(_nw204).ArraySet1((_1263_v68).Cardinality(), 24)
			(_nw204).ArraySet1(_dafny.IntOfInt64(-628), 25)
			(_nw204).ArraySet1((func() _dafny.Int {
				if (_1265_v70).Contains((_this).F30()) {
					return (_1265_v70).Multiplicity((_this).F30())
				}
				return (_this).F34()
			})(), 26)
			(_nw204).ArraySet1(((_this).F48()).Times(Companion_Default___.Fm39((_this).F47(), _1265_v70, _this.F31(), globalState)), 27)
			(_nw204).ArraySet1((_1249_v55).Cardinality(), 28)
			_1266_v71 = _nw204
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1266_v71), 0))
			_ = _index158
			(_1266_v71).ArraySet1((_this).F48(), (_index158).Int())
			var _1268_v72 _dafny.Map
			_ = _1268_v72
			_1268_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _1258_v65)
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1266_v71), 0))
			_ = _index159
			var _rhs156 _dafny.Int = _dafny.IntOfInt64(808)
			_ = _rhs156
			var _rhs157 _dafny.Int = ((_1268_v72).Cardinality()).Times(Companion_Default___.Fm39((_this).F47(), _1265_v70, _this.F31(), globalState))
			_ = _rhs157
			var _rhs158 _dafny.Int = (_dafny.Zero).Minus((_this).F34())
			_ = _rhs158
			var _lhs117 _dafny.Array = _1266_v71
			_ = _lhs117
			var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1266_v71), 0))
			_ = _lhs118
			var _lhs119 *GlobalState = globalState
			_ = _lhs119
			var _lhs120 *GlobalState = globalState
			_ = _lhs120
			(_lhs117).ArraySet1(_rhs156, (_lhs118).Int())
			_lhs119.F13 = _rhs157
			_lhs120.F13 = _rhs158
		}
		var _1269_v73 D6
		_ = _1269_v73
		_1269_v73 = Companion_D6_.Create_DC16_(p1)
		var _1270_v74 D6
		_ = _1270_v74
		_1270_v74 = Companion_D6_.Create_DC17_(_1269_v73)
		_1270_v74 = _1270_v74
		(_this).F28_set_(func(_source21 D6) bool {
			if _source21.Is_DC16() {
				var _1271___mcc_h9 _dafny.Array = _source21.Get_().(D6_DC16).Cf33
				_ = _1271___mcc_h9
				var _1272_cf33 _dafny.Array = _1271___mcc_h9
				_ = _1272_cf33
				return (_this).F47()
			} else if _source21.Is_DC15() {
				var _1273___mcc_h10 _dafny.Sequence = _source21.Get_().(D6_DC15).Cf32
				_ = _1273___mcc_h10
				var _1274_cf32 _dafny.Sequence = _1273___mcc_h10
				_ = _1274_cf32
				return false
			} else {
				var _1275___mcc_h11 D6 = _source21.Get_().(D6_DC17).Cf34
				_ = _1275___mcc_h11
				var _1276_cf34 D6 = _1275___mcc_h11
				_ = _1276_cf34
				return (_this).F47()
			}
		}(p2))
		var _1277_v75 _dafny.Map
		_ = _1277_v75
		_1277_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F48(), _this.F35())
		_1277_v75 = _1277_v75
		var _1278_v76 _dafny.Set
		_ = _1278_v76
		_1278_v76 = _dafny.SetOf(p1)
		var _1279_v77 _dafny.Map
		_ = _1279_v77
		_1279_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_this).F34())
		var _1280_v78 _dafny.Map
		_ = _1280_v78
		_1280_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1278_v76).Contains(p1), _1279_v77)
		_1280_v78 = (Companion_Default___.Fm50(_1277_v75, globalState)).Merge(_1280_v78)
		var _1281_v79 _dafny.Set
		_ = _1281_v79
		_1281_v79 = _dafny.SetOf((_this).F47(), true)
		var _1282_v80 _dafny.Map
		_ = _1282_v80
		_1282_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_this.F31(), (_this).F30(), globalState), (_this).F30())
		var _1283_v81 _dafny.MultiSet
		_ = _1283_v81
		_1283_v81 = _dafny.MultiSetOf((func() _dafny.Int {
			if (_1282_v80).Contains((_this).F47()) {
				return (_1282_v80).Get((_this).F47()).(_dafny.Int)
			}
			return (_this).F30()
		})(), (_this).F30(), (_this).F34(), (_dafny.Zero).Minus((_this).F30()), (_this).F30())
		r0 = ((_1281_v79).Cardinality()).Cmp((_1283_v81).Cardinality()) == 0
		var _1284_v82 _dafny.Sequence
		_ = _1284_v82
		_1284_v82 = _dafny.SeqOf(p0, p0)
		r1 = (_1284_v82).Select((Companion_Default___.SafeIndex((_this).F48(), _dafny.IntOfUint32((_1284_v82).Cardinality()))).Uint32()).(_dafny.CodePoint)
		return r0, r1
	}
}
func (_this *C5) F47() bool {
	{
		return _this._f47
	}
}
func (_this *C5) F48() _dafny.Int {
	{
		return _this._f48
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f36 _dafny.CodePoint
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f36 = _dafny.CodePoint('D')
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T2_.TraitID_, Companion_T3_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C6{}
var _ T2 = &C6{}
var _ T3 = &C6{}
var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F28() bool {
	return _this._f28
}
func (_this *C6) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C6) F29() bool {
	return _this._f29
}
func (_this *C6) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C6) F31() bool {
	return _this._f31
}
func (_this *C6) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C6) F35() bool {
	return _this._f35
}
func (_this *C6) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C6) F36() _dafny.CodePoint {
	return _this._f36
}
func (_this *C6) F36_set_(value _dafny.CodePoint) {
	_this._f36 = value
}
func (_this *C6) F34() _dafny.Int {
	return _this._f34
}
func (_this *C6) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C6) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C6) F30() _dafny.Int {
	return _this._f30
}
func (_this *C6) Ctor__(f36 _dafny.CodePoint, f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f36 = f36
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C6) Fm7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _this.F31()
	}
}
func (_this *C6) Fm6(globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.CodePoint('y')), _dafny.SeqOf(_dafny.CodePoint('v')))
	}
}
func (_this *C6) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return ((func() _dafny.Set {
			var _coll36 = _dafny.NewBuilder()
			_ = _coll36
			for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(924), _dafny.IntOfInt64(550))); ; {
				_compr_36, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _1285_v0 _dafny.Int
				_1285_v0 = interface{}(_compr_36).(_dafny.Int)
				if ((_dafny.IntOfInt64(924)).Cmp(_1285_v0) <= 0) && ((_1285_v0).Cmp(_dafny.IntOfInt64(550)) < 0) {
					_coll36.Add((_1285_v0).Times(_dafny.IntOfUint32(((_this).F32()).Cardinality())))
				}
			}
			return _coll36.ToSet()
		}()).Difference(_dafny.SetOf(_dafny.IntOfInt64(-281)))).Intersection(_dafny.SetOf(_dafny.IntOfInt64(543), _dafny.IntOfUint32((_dafny.SeqOf((_this).F34(), (_this).F34(), (_this).F30())).Cardinality()), (_this).F30(), (_this).F30()))
	}
}
func (_this *C6) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return ((_this).F33()).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if false {
				return (_this).F30()
			}
			return (_this).F30()
		})(), _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence)
	}
}
func (_this *C6) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("crmrunqsx")
	}
}
func (_this *C6) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())
	}
}
func (_this *C6) Fm27(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		if (func() bool {
			if _this.F29() {
				return false
			}
			return _this.F29()
		})() {
			return (_dafny.Zero).Minus((_this).F30())
		} else {
			return (_this).F34()
		}
	}
}
func (_this *C6) Fm28(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-675)
	}
}
func (_this *C6) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _hi7 _dafny.Int = (_this).F34()
		_ = _hi7
		for _1286_i0 := (_this).F30(); _1286_i0.Cmp(_hi7) < 0; _1286_i0 = _1286_i0.Plus(_dafny.One) {
			var _1287_v0 _dafny.Array
			_ = _1287_v0
			var _nw205 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw205
			_1287_v0 = _nw205
			var _1288_v1 _dafny.Sequence
			_ = _1288_v1
			_1288_v1 = _dafny.SeqOf(_1287_v0)
			var _1289_v2 _dafny.Map
			_ = _1289_v2
			_1289_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1286_i0).Times(_1286_i0), (_1288_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(380))).Cardinality()), _dafny.IntOfUint32((_1288_v1).Cardinality()))).Uint32()).(_dafny.Array))
			var _1290_v3 _dafny.Map
			_ = _1290_v3
			_1290_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _1287_v0)
			_1289_v2 = (_1289_v2).Update(Companion_Default___.SafeDivisionInt(_1286_i0, _dafny.IntOfInt64(-558)), (func() _dafny.Array {
				if (_1290_v3).Contains(_this.F28()) {
					return (_1290_v3).Get(_this.F28()).(_dafny.Array)
				}
				return _1287_v0
			})())
			var _1291_v4 _dafny.Sequence
			_ = _1291_v4
			_1291_v4 = _dafny.UnicodeSeqOfUtf8Bytes("xku")
			var _1292_v5 _dafny.Map
			_ = _1292_v5
			_1292_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-301), _1291_v4)
			if (_1286_i0).Cmp((_1292_v5).Cardinality()) != 0 {
				var _1293_v6 _dafny.Sequence
				_ = _1293_v6
				_1293_v6 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(476), _1286_i0))
				var _1294_v7 _dafny.Map
				_ = _1294_v7
				_1294_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1293_v6).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1293_v6).Cardinality()))).Uint32()).(_dafny.Set), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F32(), (_this).F34())).Cardinality()))
				var _1295_v8 _dafny.Set
				_ = _1295_v8
				_1295_v8 = _dafny.SetOf(_dafny.IntOfInt64(12), _1286_i0, _1286_i0)
				var _1296_v9 _dafny.Array
				_ = _1296_v9
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
				_ = _nw206
				_1296_v9 = _nw206
				var _1297_v10 _dafny.Map
				_ = _1297_v10
				_1297_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _this.F28())
				var _1298_v11 T4
				_ = _1298_v11
				var _nw207 *C3 = New_C3_()
				_ = _nw207
				_nw207.Ctor__(_1286_i0, _1296_v9, (_1291_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F34())).Cardinality()), _dafny.IntOfUint32((_1291_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F34(), _this.F29(), (_this).F32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(658))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg89 _dafny.Int) interface{} {
						return coer89(arg89)
					}
				}(func(_1299_i1 _dafny.Int) _dafny.Sequence {
					return (_this).F32()
				})), _dafny.IntOfUint32((_dafny.SeqOf((_1297_v10).Cardinality(), _dafny.IntOfInt64(387))).Cardinality()), _this.F29(), _this.F31(), _this.F35())
				_1298_v11 = _nw207
				var _1300_v12 _dafny.Sequence
				_ = _1300_v12
				_1300_v12 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yor")).Cardinality()), (_this).F30(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _1298_v11)).Cardinality())
				_1294_v7 = (_1294_v7).Update((func() _dafny.Set {
					if _this.F29() {
						return _1295_v8
					}
					return _1295_v8
				})(), _dafny.Companion_Sequence_.Concatenate(_1300_v12, _1300_v12))
				var _1301_v13 _dafny.Sequence
				_ = _1301_v13
				_1301_v13 = _dafny.SeqOf(_1296_v9, _1296_v9)
				var _1302_v14 _dafny.Map
				_ = _1302_v14
				_1302_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _dafny.CodePoint('i'))
				var _1303_v15 *C3
				_ = _1303_v15
				var _nw208 *C3 = New_C3_()
				_ = _nw208
				_nw208.Ctor__(_1286_i0, (_1301_v13).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1301_v13).Cardinality()))).Uint32()).(_dafny.Array), (func() _dafny.CodePoint {
					if (_1302_v14).Contains(_1298_v11.F36()) {
						return (_1302_v14).Get(_1298_v11.F36()).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('u')
				})(), _dafny.IntOfInt64(35), _this.F28(), _dafny.SeqOf(_1298_v11.F29(), _1298_v11.F28()), (_this).F33(), _1286_i0, (!(_1298_v11.F31())) || (_this.F28()), true, _this.F31())
				_1303_v15 = _nw208
				var _1304_v16 _dafny.Array
				_ = _1304_v16
				var _nw209 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
				_ = _nw209
				_1304_v16 = _nw209
				var _1305_v17 _dafny.MultiSet
				_ = _1305_v17
				_1305_v17 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("twqhca"))
				var _1306_v18 _dafny.Map
				_ = _1306_v18
				_1306_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1305_v17).Contains(_1291_v4)), _1304_v16)
				var _rhs159 _dafny.Int = ((_this).F30()).Times((func() _dafny.Int {
					if _1298_v11.F29() {
						return (_1298_v11).F34()
					}
					return (_1303_v15).F44()
				})())
				_ = _rhs159
				var _rhs160 _dafny.Array = (func() _dafny.Array {
					if (_1306_v18).Contains(!(_1298_v11.F28())) {
						return (_1306_v18).Get(!(_1298_v11.F28())).(_dafny.Array)
					}
					return _1304_v16
				})()
				_ = _rhs160
				var _rhs161 bool = !(!(!(_dafny.Companion_Sequence_.Contains(_1291_v4, _1298_v11.F36()))))
				_ = _rhs161
				var _lhs121 *GlobalState = globalState
				_ = _lhs121
				r1 = _rhs159
				_1304_v16 = _rhs160
				_lhs121.F26 = _rhs161
				var _1307_v19 *C2
				_ = _1307_v19
				var _nw210 *C2 = New_C2_()
				_ = _nw210
				_nw210.Ctor__((_this).F30(), (_this).F30(), _1298_v11.F29(), _1298_v11.F31(), _1298_v11.F35())
				_1307_v19 = _nw210
				var _1308_v20 _dafny.Map
				_ = _1308_v20
				_1308_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1298_v11.F31(), _1291_v4)
				var _1309_v21 _dafny.Map
				_ = _1309_v21
				_1309_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1307_v19, _dafny.IntOfUint32((_1291_v4).Cardinality()))).Cardinality()).Plus((_1307_v19).F46()), (func() _dafny.Map {
					if _this.F31() {
						return _1308_v20
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _1291_v4)
				})())
				var _1310_v22 D5
				_ = _1310_v22
				_1310_v22 = Companion_D5_.Create_DC12_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-852))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg90 _dafny.Int) interface{} {
						return coer90(arg90)
					}
				}((func(_1311_v11 T4) func(_dafny.Int) _dafny.CodePoint {
					return func(_1312_i2 _dafny.Int) _dafny.CodePoint {
						return _1311_v11.F36()
					}
				})(_1298_v11))))
				_1309_v21 = (_1309_v21).Update((_1303_v15).F44(), (func() _dafny.Map {
					if !(_this.F31()) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1298_v11.F28(), (_1310_v22).Dtor_cf26())
					}
					return _1308_v20
				})())
				var _rhs162 bool = _this.F31()
				_ = _rhs162
				var _rhs163 _dafny.Int = _1286_i0
				_ = _rhs163
				var _rhs164 _dafny.Int = (_dafny.IntOfInt64(308)).Times((_this).F34())
				_ = _rhs164
				var _lhs122 *C6 = _this
				_ = _lhs122
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				var _lhs124 *GlobalState = globalState
				_ = _lhs124
				_lhs122.F35_set_(_rhs162)
				_lhs123.F13 = _rhs163
				_lhs124.F13 = _rhs164
			} else {
				var _1313_v23 _dafny.Array
				_ = _1313_v23
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw211
				_1313_v23 = _nw211
				var _1314_v24 *C3
				_ = _1314_v24
				var _nw212 *C3 = New_C3_()
				_ = _nw212
				_nw212.Ctor__(_dafny.IntOfInt64(-200), _1313_v23, _dafny.CodePoint('s'), (_this).Fm28((_this).F30(), true, globalState), _this.F35(), (_this).F32(), (_this).F33(), _1286_i0, _dafny.Companion_Sequence_.IsPrefixOf((_this).F32(), (_this).F32()), _this.F29(), ((_this).F34()).Cmp(_dafny.IntOfUint32((_1291_v4).Cardinality())) == 0)
				_1314_v24 = _nw212
				_1291_v4 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}(func(_1315_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}))
				var _1316_v25 _dafny.Map
				_ = _1316_v25
				_1316_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v24, _dafny.IntOfUint32(((_this).F32()).Cardinality()))
				var _1317_v26 _dafny.Sequence
				_ = _1317_v26
				_1317_v26 = _dafny.SeqOf(_1316_v25, _1316_v25, _1316_v25, _1316_v25)
				var _1318_v27 _dafny.Map
				_ = _1318_v27
				_1318_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())
				r1 = (_dafny.Zero).Minus((func() _dafny.Int {
					if _dafny.Companion_Sequence_.IsPrefixOf(_1317_v26, _1317_v26) {
						return (_1318_v27).Cardinality()
					}
					return _1286_i0
				})())
				var _1319_v28 D5
				_ = _1319_v28
				_1319_v28 = Companion_D5_.Create_DC14_(_this.F35(), (_this).F30())
				var _1320_v29 _dafny.Array
				_ = _1320_v29
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_26
				var _nw213 _dafny.Array
				_ = _nw213
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw213 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.CodePoint = func(_1321_i4 _dafny.Int) _dafny.CodePoint {
						return _this.F36()
					}
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw213 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw213).ArraySet1CodePoint(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw213).ArraySet1CodePoint(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1320_v29 = _nw213
				var _1322_v30 *C0
				_ = _1322_v30
				var _nw214 *C0 = New_C0_()
				_ = _nw214
				_nw214.Ctor__((func() _dafny.Array {
					if (_1319_v28).Dtor_cf30() {
						return _1320_v29
					}
					return _1320_v29
				})())
				_1322_v30 = _nw214
				(globalState).F13 = (_1314_v24).F44()
			}
			var _1323_v31 _dafny.Map
			_ = _1323_v31
			_1323_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _this.F29())
			var _1324_v32 _dafny.MultiSet
			_ = _1324_v32
			_1324_v32 = _dafny.MultiSetOf(_dafny.IntOfInt64(605), (_this).F34(), (_this).F30(), (_1323_v31).Cardinality())
			var _1325_v33 _dafny.Sequence
			_ = _1325_v33
			_1325_v33 = _dafny.SeqOf(true, (_1324_v32).IsDisjointFrom(_1324_v32))
			_1325_v33 = _dafny.Companion_Sequence_.Concatenate((_this).F32(), _dafny.SeqOf(false, !(_this.F35())))
			var _1326_v34 _dafny.Set
			_ = _1326_v34
			_1326_v34 = _dafny.SetOf((_this).F30())
			_1326_v34 = ((_1326_v34).Difference(func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter39 := _dafny.Iterate((_1326_v34).Elements()); ; {
					_compr_37, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _1327_v35 _dafny.Int
					_1327_v35 = interface{}(_compr_37).(_dafny.Int)
					if (_1326_v34).Contains(_1327_v35) {
						_coll37.Add(Companion_Default___.SafeModuloInt(_1327_v35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(-246))).Cardinality()))
					}
				}
				return _coll37.ToSet()
			}())).Difference(_1326_v34)
		}
		var _1328_v36 _dafny.Array
		_ = _1328_v36
		var _nw215 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw215
		_1328_v36 = _nw215
		var _1329_v37 _dafny.Map
		_ = _1329_v37
		_1329_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1328_v36, (_this).F30())
		var _1330_v38 D7
		_ = _1330_v38
		_1330_v38 = Companion_D7_.Create_DC18_((_1329_v37).Update(_1328_v36, (_this).F30()))
		var _pat_let_tv13 = _1329_v37
		_ = _pat_let_tv13
		var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))
		_ = _index160
		(_1328_v36).ArraySet1(((func(_pat_let19_0 D7) D7 {
			return func(_1331_dt__update__tmp_h0 D7) D7 {
				return func(_pat_let20_0 _dafny.Map) D7 {
					return func(_1332_dt__update_hcf35_h0 _dafny.Map) D7 {
						return Companion_D7_.Create_DC18_(_1332_dt__update_hcf35_h0)
					}(_pat_let20_0)
				}(_pat_let_tv13)
			}(_pat_let19_0)
		}(_1330_v38)).Dtor_cf35()).Cardinality(), (_index160).Int())
		var _1333_v39 _dafny.MultiSet
		_ = _1333_v39
		_1333_v39 = _dafny.MultiSetOf(true)
		var _1334_v40 D0
		_ = _1334_v40
		_1334_v40 = Companion_D0_.Create_DC1_(_1333_v39)
		var _1335_v41 _dafny.Map
		_ = _1335_v41
		_1335_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v40, _this.F29())
		var _1336_v42 D8
		_ = _1336_v42
		_1336_v42 = Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v40, _this.F35()))
		var _1337_v44 _dafny.Set
		_ = _1337_v44
		_1337_v44 = _dafny.SetOf(_1334_v40)
		var _1338_v46 _dafny.Sequence
		_ = _1338_v46
		_1338_v46 = _dafny.SeqOf((Companion_D8_.Create_DC21_(_1335_v41)).Dtor_cf40())
		var _pat_let_tv14 = _1333_v39
		_ = _pat_let_tv14
		var _1339_v47 _dafny.Array
		_ = _1339_v47
		var _nwElement0_47 _dafny.Map = _1335_v41
		_ = _nwElement0_47
		var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(25))
		_ = _nw216
		(_nw216).ArraySet1(_nwElement0_47, 0)
		(_nw216).ArraySet1((func() _dafny.Map {
			if _this.F28() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v40, _this.F28())
			}
			return _1335_v41
		})(), 1)
		(_nw216).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v40, !(!(_this.F35())))).Merge(_1335_v41), 2)
		(_nw216).ArraySet1(_1335_v41, 3)
		(_nw216).ArraySet1(((_1336_v42).Dtor_cf40()).Merge(_1335_v41), 4)
		(_nw216).ArraySet1(_1335_v41, 5)
		(_nw216).ArraySet1(_1335_v41, 6)
		(_nw216).ArraySet1(_1335_v41, 7)
		(_nw216).ArraySet1(_1335_v41, 8)
		(_nw216).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_1333_v39), Companion_Default___.Fm1(_this.F31(), (_this).F34(), globalState)), 9)
		(_nw216).ArraySet1(_1335_v41, 10)
		(_nw216).ArraySet1(Companion_Default___.Fm35(_this.F35(), globalState), 11)
		(_nw216).ArraySet1(_1335_v41, 12)
		(_nw216).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm36(_dafny.IntOfInt64(893), (_this).F34(), _this.F31(), globalState), !(_this.F35()))).Merge(_1335_v41), 13)
		(_nw216).ArraySet1((func() _dafny.Map {
			if _this.F28() {
				return _1335_v41
			}
			return func() _dafny.Map {
				var _coll38 = _dafny.NewMapBuilder()
				_ = _coll38
				for _iter40 := _dafny.Iterate((_1337_v44).Elements()); ; {
					_compr_38, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _1340_v43 D0
					_1340_v43 = interface{}(_compr_38).(D0)
					if (_1337_v44).Contains(_1340_v43) {
						_coll38.Add(_1340_v43, _this.F35())
					}
				}
				return _coll38.ToMap()
			}()
		})(), 14)
		(_nw216).ArraySet1((_1335_v41).Merge(func() _dafny.Map {
			var _coll39 = _dafny.NewMapBuilder()
			_ = _coll39
			for _iter41 := _dafny.Iterate((_1335_v41).Keys().Elements()); ; {
				_compr_39, _ok41 := _iter41()
				if !_ok41 {
					break
				}
				var _1341_v45 D0
				_1341_v45 = interface{}(_compr_39).(D0)
				if (_1335_v41).Contains(_1341_v45) {
					_coll39.Add(_1341_v45, true)
				}
			}
			return _coll39.ToMap()
		}()), 15)
		(_nw216).ArraySet1((_1338_v46).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F34()), _dafny.IntOfUint32((_1338_v46).Cardinality()))).Uint32()).(_dafny.Map), 16)
		(_nw216).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let21_0 D0) D0 {
			return func(_1342_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let22_0 _dafny.MultiSet) D0 {
					return func(_1343_dt__update_hcf1_h0 _dafny.MultiSet) D0 {
						return Companion_D0_.Create_DC1_(_1343_dt__update_hcf1_h0)
					}(_pat_let22_0)
				}(_pat_let_tv14)
			}(_pat_let21_0)
		}(Companion_D0_.Create_DC1_(_1333_v39)), _this.F28()), 17)
		(_nw216).ArraySet1(_1335_v41, 18)
		(_nw216).ArraySet1(_1335_v41, 19)
		(_nw216).ArraySet1(_1335_v41, 20)
		(_nw216).ArraySet1(_1335_v41, 21)
		(_nw216).ArraySet1(Companion_Default___.Fm35(_this.F35(), globalState), 22)
		(_nw216).ArraySet1(_1335_v41, 23)
		(_nw216).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1334_v40, _this.F28()), 24)
		_1339_v47 = _nw216
		var _1344_v48 _dafny.Array
		_ = _1344_v48
		var _nw217 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw217
		_1344_v48 = _nw217
		var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1344_v48), 0))
		_ = _index161
		(_1344_v48).ArraySet1(_this.F35(), (_index161).Int())
		var _1345_v49 _dafny.Array
		_ = _1345_v49
		var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
		_ = _nw218
		_1345_v49 = _nw218
		var _1346_v50 _dafny.Sequence
		_ = _1346_v50
		_1346_v50 = _dafny.UnicodeSeqOfUtf8Bytes("ehackpxcv")
		var _1347_v51 _dafny.Map
		_ = _1347_v51
		_1347_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(781))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg92 _dafny.Int) interface{} {
				return coer92(arg92)
			}
		}(func(_1348_i5 _dafny.Int) _dafny.CodePoint {
			return _this.F36()
		})), _this.F29())
		var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1345_v49), 0))
		_ = _index162
		(_1345_v49).ArraySet1CodePoint(Companion_Default___.Fm34(_dafny.IntOfUint32((_1346_v50).Cardinality()), (func() _dafny.Int {
			if (_1333_v39).Contains(_this.F35()) {
				return (_1333_v39).Multiplicity(_this.F35())
			}
			return (Companion_Default___.Fm37(_this.F36(), _this.F31(), globalState)).Cardinality()
		})(), (_this).F30(), (func() bool {
			if (_1347_v51).Contains(_dafny.UnicodeSeqOfUtf8Bytes("exfveuas")) {
				return (_1347_v51).Get(_dafny.UnicodeSeqOfUtf8Bytes("exfveuas")).(bool)
			}
			return true
		})(), globalState), (_index162).Int())
		var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))
		_ = _index163
		var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1344_v48), 0))
		_ = _index164
		var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1345_v49), 0))
		_ = _index165
		var _rhs165 _dafny.Int = ((_this).Fm28((_this).F34(), _this.F35(), globalState)).Plus((_this).Fm28((_this).F30(), _this.F29(), globalState))
		_ = _rhs165
		var _rhs166 _dafny.Array = _1339_v47
		_ = _rhs166
		var _rhs167 bool = !(true)
		_ = _rhs167
		var _rhs168 bool = (_this.F29()) || (_this.F28())
		_ = _rhs168
		var _rhs169 _dafny.CodePoint = _dafny.CodePoint('u')
		_ = _rhs169
		var _lhs125 _dafny.Array = _1328_v36
		_ = _lhs125
		var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))
		_ = _lhs126
		var _lhs127 *GlobalState = globalState
		_ = _lhs127
		var _lhs128 _dafny.Array = _1344_v48
		_ = _lhs128
		var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1344_v48), 0))
		_ = _lhs129
		var _lhs130 _dafny.Array = _1345_v49
		_ = _lhs130
		var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1345_v49), 0))
		_ = _lhs131
		(_lhs125).ArraySet1(_rhs165, (_lhs126).Int())
		_1339_v47 = _rhs166
		_lhs127.F26 = _rhs167
		(_lhs128).ArraySet1(_rhs168, (_lhs129).Int())
		(_lhs130).ArraySet1CodePoint(_rhs169, (_lhs131).Int())
		var _1349_v52 _dafny.MultiSet
		_ = _1349_v52
		_1349_v52 = _dafny.MultiSetOf((_this).Fm27((_1344_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1344_v48), 0))).Int()).(bool), (_1328_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))).Int()).(_dafny.Int), (_this).F30(), globalState))
		var _1350_v53 _dafny.Sequence
		_ = _1350_v53
		_1350_v53 = _dafny.SeqOf((_1328_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))).Int()).(_dafny.Int), (_1328_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))).Int()).(_dafny.Int), (_this).F30(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-627), (_this).F34())).Cardinality()), (_dafny.Zero).Minus((_this).F30()))
		var _1351_v54 *C2
		_ = _1351_v54
		var _nw219 *C2 = New_C2_()
		_ = _nw219
		_nw219.Ctor__((_1328_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
			if (_1349_v52).Contains((_1350_v53).Select((Companion_Default___.SafeIndex((_1328_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1350_v53).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_1349_v52).Multiplicity((_1350_v53).Select((Companion_Default___.SafeIndex((_1328_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_1328_v36), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1350_v53).Cardinality()))).Uint32()).(_dafny.Int))
			}
			return (_this).F34()
		})(), _this.F28(), _this.F35(), !(_this.F29()))
		_1351_v54 = _nw219
		var _1352_v55 _dafny.Map
		_ = _1352_v55
		_1352_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(799))).Uint32(), func(coer93 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}(func(_1353_i6 _dafny.Int) _dafny.Int {
			return (_this).F30()
		})), _dafny.IntOfUint32((_1350_v53).Cardinality()))
		var _rhs170 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_dafny.IntOfUint32(((_this).F32()).Cardinality())).Cmp((_this).F30()) >= 0 {
				return _1346_v50
			}
			return _1346_v50
		})(), (Companion_Default___.SafeIndex(((_1351_v54).F46()).Minus((_dafny.Zero).Minus((_1352_v55).Cardinality())), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_dafny.IntOfUint32(((_this).F32()).Cardinality())).Cmp((_this).F30()) >= 0 {
				return _1346_v50
			}
			return _1346_v50
		})()).Cardinality()))).Uint32(), (_1345_v49).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1345_v49), 0))).Int()))).Cardinality())
		_ = _rhs170
		var _rhs171 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1346_v50, _1346_v50), _dafny.SeqOf(_this.F36()))
		_ = _rhs171
		var _rhs172 _dafny.Int = (func() _dafny.Int {
			if _this.F35() {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1346_v50, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(178))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}((func(_1354_v49 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_1355_i7 _dafny.Int) _dafny.CodePoint {
						return (_1354_v49).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1354_v49), 0))).Int())
					}
				})(_1345_v49))), (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(178))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}((func(_1356_v49 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_1357_i7 _dafny.Int) _dafny.CodePoint {
						return (_1356_v49).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1356_v49), 0))).Int())
					}
				})(_1345_v49)))).Cardinality()))).Uint32(), (_1345_v49).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_1345_v49), 0))).Int())))).Cardinality())
			}
			return (_this).F34()
		})()
		_ = _rhs172
		var _rhs173 bool = !(_this.F35()) || (_this.F28())
		_ = _rhs173
		var _lhs132 *GlobalState = globalState
		_ = _lhs132
		var _lhs133 *GlobalState = globalState
		_ = _lhs133
		r1 = _rhs170
		_1346_v50 = _rhs171
		_lhs132.F13 = _rhs172
		_lhs133.F12 = _rhs173
		var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1344_v48), 0))
		_ = _index166
		(_1344_v48).ArraySet1(_this.F29(), (_index166).Int())
		var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1344_v48), 0))
		_ = _index167
		(_1344_v48).ArraySet1((_1351_v54).Fm30(false, globalState), (_index167).Int())
		r0 = _this.F31()
		r1 = (_this).F34()
		r2 = _1346_v50
		return r0, r1, r2
	}
}
func (_this *C6) M11(p0 bool, globalState *GlobalState) (bool, bool, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1358_v0 D8
		_ = _1358_v0
		_1358_v0 = Companion_D8_.Create_DC22_((_this).F34(), (_this).F34())
		var _1359_v1 D8
		_ = _1359_v1
		_1359_v1 = Companion_D8_.Create_DC25_(Companion_D8_.Create_DC25_(_1358_v0))
		var _source22 D8 = _1359_v1
		_ = _source22
		if _source22.Is_DC22() {
			var _1360___mcc_h0 _dafny.Int = _source22.Get_().(D8_DC22).Cf41
			_ = _1360___mcc_h0
			var _1361___mcc_h1 _dafny.Int = _source22.Get_().(D8_DC22).Cf42
			_ = _1361___mcc_h1
			var _1362_cf42 _dafny.Int = _1361___mcc_h1
			_ = _1362_cf42
			var _1363_cf41 _dafny.Int = _1360___mcc_h0
			_ = _1363_cf41
			var _1364_v2 _dafny.Array
			_ = _1364_v2
			var _nw220 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw220
			_1364_v2 = _nw220
			var _1365_v3 _dafny.Array
			_ = _1365_v3
			var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw221
			_1365_v3 = _nw221
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_1364_v2), 0))
			_ = _index168
			(_1364_v2).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _1365_v3), (_index168).Int())
			var _1366_v4 _dafny.Map
			_ = _1366_v4
			_1366_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _1365_v3)
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_1364_v2), 0))
			_ = _index169
			(_1364_v2).ArraySet1((_1366_v4).Merge(_1366_v4), (_index169).Int())
			var _1367_v5 _dafny.Array
			_ = _1367_v5
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_27
			var _nw222 _dafny.Array
			_ = _nw222
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw222 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = func(_1368_i0 _dafny.Int) bool {
					return false
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw222 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw222).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw222).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1367_v5 = _nw222
			var _1369_v6 D1
			_ = _1369_v6
			_1369_v6 = Companion_D1_.Create_DC2_(_1367_v5)
			var _1370_v7 D1
			_ = _1370_v7
			_1370_v7 = Companion_D1_.Create_DC4_(_1369_v6)
			var _source23 D1 = _1370_v7
			_ = _source23
			if _source23.Is_DC3() {
				var _1371___mcc_h12 bool = _source23.Get_().(D1_DC3).Cf3
				_ = _1371___mcc_h12
				var _1372___mcc_h13 _dafny.Int = _source23.Get_().(D1_DC3).Cf4
				_ = _1372___mcc_h13
				var _1373___mcc_h14 *C0 = _source23.Get_().(D1_DC3).Cf5
				_ = _1373___mcc_h14
				var _1374___mcc_h15 _dafny.Int = _source23.Get_().(D1_DC3).Cf6
				_ = _1374___mcc_h15
				var _1375_cf6 _dafny.Int = _1374___mcc_h15
				_ = _1375_cf6
				var _1376_cf5 *C0 = _1373___mcc_h14
				_ = _1376_cf5
				var _1377_cf4 _dafny.Int = _1372___mcc_h13
				_ = _1377_cf4
				var _1378_cf3 bool = _1371___mcc_h12
				_ = _1378_cf3
				var _1379_v8 _dafny.Sequence
				_ = _1379_v8
				_1379_v8 = _dafny.UnicodeSeqOfUtf8Bytes("tyollv")
				var _1380_v9 _dafny.Sequence
				_ = _1380_v9
				_1380_v9 = _dafny.SeqOf(_dafny.IntOfUint32((_1379_v8).Cardinality()))
				var _1381_v10 _dafny.Int
				_ = _1381_v10
				var _out56 _dafny.Int
				_ = _out56
				_out56 = Companion_Default___.M0(_1380_v9, !(_this.F29()), _1363_cf41, p0, globalState)
				_1381_v10 = _out56
				var _1382_v11 _dafny.Array
				_ = _1382_v11
				var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw223
				_1382_v11 = _nw223
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1382_v11), 0))
				_ = _index170
				(_1382_v11).ArraySet1((func() _dafny.Array {
					if _this.F31() {
						return _1367_v5
					}
					return _1367_v5
				})(), (_index170).Int())
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1365_v3), 0))
				_ = _index171
				(_1365_v3).ArraySet1(_1377_cf4, (_index171).Int())
				var _1383_v12 _dafny.MultiSet
				_ = _1383_v12
				_1383_v12 = _dafny.MultiSetOf(!(_this.F35()))
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1382_v11), 0))
				_ = _index172
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1365_v3), 0))
				_ = _index173
				var _rhs174 _dafny.Array = _1367_v5
				_ = _rhs174
				var _rhs175 _dafny.Int = (func() _dafny.Int {
					if (_1383_v12).Contains(_this.F31()) {
						return (_1383_v12).Multiplicity(_this.F31())
					}
					return (_this).F30()
				})()
				_ = _rhs175
				var _rhs176 _dafny.Int = _1375_cf6
				_ = _rhs176
				var _rhs177 D8 = _1359_v1
				_ = _rhs177
				var _lhs134 _dafny.Array = _1382_v11
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1382_v11), 0))
				_ = _lhs135
				var _lhs136 _dafny.Array = _1365_v3
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1365_v3), 0))
				_ = _lhs137
				(_lhs134).ArraySet1(_rhs174, (_lhs135).Int())
				r2 = _rhs175
				(_lhs136).ArraySet1(_rhs176, (_lhs137).Int())
				_1359_v1 = _rhs177
				_1379_v8 = _1379_v8
				r2 = _1363_cf41
			} else if _source23.Is_DC2() {
				var _1384___mcc_h16 _dafny.Array = _source23.Get_().(D1_DC2).Cf2
				_ = _1384___mcc_h16
				var _1385_cf2 _dafny.Array = _1384___mcc_h16
				_ = _1385_cf2
				var _1386_v13 _dafny.Array
				_ = _1386_v13
				var _nw224 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
				_ = _nw224
				_1386_v13 = _nw224
				var _1387_v14 _dafny.MultiSet
				_ = _1387_v14
				_1387_v14 = _dafny.MultiSetOf(_1362_cf42)
				var _1388_v15 _dafny.Set
				_ = _1388_v15
				_1388_v15 = _dafny.SetOf(_1387_v14)
				var _1389_v16 *C2
				_ = _1389_v16
				var _nw225 *C2 = New_C2_()
				_ = _nw225
				_nw225.Ctor__(_dafny.IntOfInt64(643), (_1388_v15).Cardinality(), _this.F28(), p0, !(_this.F28()))
				_1389_v16 = _nw225
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1386_v13), 0))
				_ = _index174
				(_1386_v13).ArraySet1(_1389_v16, (_index174).Int())
				var _1390_v17 D9
				_ = _1390_v17
				_1390_v17 = Companion_D9_.Create_DC26_(_1389_v16)
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_1386_v13), 0))
				_ = _index175
				(_1386_v13).ArraySet1((_1390_v17).Dtor_cf52(), (_index175).Int())
				var _1391_v18 _dafny.MultiSet
				_ = _1391_v18
				_1391_v18 = _dafny.MultiSetOf((_this).F32(), (_this).F32())
				(globalState).F13 = ((_this).F30()).Minus((_1391_v18).Cardinality())
				var _1392_v19 _dafny.Set
				_ = _1392_v19
				_1392_v19 = _dafny.SetOf((_this).F30())
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1365_v3), 0))
				_ = _index176
				(_1365_v3).ArraySet1((_1392_v19).Cardinality(), (_index176).Int())
				var _1393_v20 _dafny.Map
				_ = _1393_v20
				_1393_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), (_1389_v16).F46())
				var _1394_v21 _dafny.Map
				_ = _1394_v21
				_1394_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1365_v3, (_1393_v20).Cardinality())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1365_v3), 0))
				_ = _index177
				(_1365_v3).ArraySet1((func() _dafny.Int {
					if (_1394_v21).Contains(_1365_v3) {
						return (_1394_v21).Get(_1365_v3).(_dafny.Int)
					}
					return (_this).Fm27(_this.F29(), _1363_cf41, _dafny.IntOfInt64(736), globalState)
				})(), (_index177).Int())
				var _1395_v22 _dafny.Map
				_ = _1395_v22
				_1395_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _this.F31())
				var _1396_v23 _dafny.Sequence
				_ = _1396_v23
				_1396_v23 = _dafny.SeqOf(_1362_cf42)
				_1395_v22 = (_1395_v22).Update(!((func() bool {
					if false {
						return _this.F31()
					}
					return _this.F31()
				})()), _dafny.Companion_Sequence_.Equal(_1396_v23, _1396_v23))
			} else {
				var _1397___mcc_h17 D1 = _source23.Get_().(D1_DC4).Cf7
				_ = _1397___mcc_h17
				var _1398_cf7 D1 = _1397___mcc_h17
				_ = _1398_cf7
				var _1399_v24 _dafny.Map
				_ = _1399_v24
				_1399_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1367_v5, _this.F36())
				var _1400_v25 _dafny.Map
				_ = _1400_v25
				_1400_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1399_v24)
				var _1401_v26 _dafny.Array
				_ = _1401_v26
				var _nwElement0_48 _dafny.Map = _1399_v24
				_ = _nwElement0_48
				var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(15))
				_ = _nw226
				(_nw226).ArraySet1(_nwElement0_48, 0)
				(_nw226).ArraySet1(_1399_v24, 1)
				(_nw226).ArraySet1(_1399_v24, 2)
				(_nw226).ArraySet1(_1399_v24, 3)
				(_nw226).ArraySet1((_1399_v24).Merge(_1399_v24), 4)
				(_nw226).ArraySet1((_1399_v24).Merge(_1399_v24), 5)
				(_nw226).ArraySet1(_1399_v24, 6)
				(_nw226).ArraySet1(_1399_v24, 7)
				(_nw226).ArraySet1(((_1399_v24).Update(_1367_v5, _this.F36())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1367_v5, _this.F36())), 8)
				(_nw226).ArraySet1(_1399_v24, 9)
				(_nw226).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1367_v5, _this.F36()), 10)
				(_nw226).ArraySet1((func() _dafny.Map {
					if (_1400_v25).Contains(_this.F31()) {
						return (_1400_v25).Get(_this.F31()).(_dafny.Map)
					}
					return _1399_v24
				})(), 11)
				(_nw226).ArraySet1((_1399_v24).Merge(_1399_v24), 12)
				(_nw226).ArraySet1(_1399_v24, 13)
				(_nw226).ArraySet1(_1399_v24, 14)
				_1401_v26 = _nw226
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1401_v26), 0))
				_ = _index178
				(_1401_v26).ArraySet1(_1399_v24, (_index178).Int())
				var _1402_v27 D1
				_ = _1402_v27
				_1402_v27 = Companion_D1_.Create_DC2_(_1367_v5)
				var _1403_v28 *C1
				_ = _1403_v28
				var _nw227 *C1 = New_C1_()
				_ = _nw227
				_nw227.Ctor__(_this.F35(), _this.F29(), false)
				_1403_v28 = _nw227
				var _1404_v29 _dafny.Map
				_ = _1404_v29
				_1404_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1403_v28, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("auuk")).Cardinality()))
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1401_v26), 0))
				_ = _index179
				var _rhs178 _dafny.Map = _1399_v24
				_ = _rhs178
				var _rhs179 bool = false
				_ = _rhs179
				var _rhs180 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(781))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}((func(_1405_cf41 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
					return func(_1406_i1 _dafny.Int) _dafny.CodePoint {
						return (Companion_D5_.Create_DC13_(_this.F28(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1405_cf41), _this.F36())).Dtor_cf29()
					}
				})(_1363_cf41)))).Cardinality()), ((_this).F34()).Plus((func() _dafny.Int {
					if (_1404_v29).Contains(_1403_v28) {
						return (_1404_v29).Get(_1403_v28).(_dafny.Int)
					}
					return (_this).Fm28((_dafny.MultiSetOf(_this.F35())).Cardinality(), _this.F31(), globalState)
				})()))))
				_ = _rhs180
				var _rhs181 _dafny.Int = _1362_cf42
				_ = _rhs181
				var _rhs182 D1 = _1402_v27
				_ = _rhs182
				var _lhs138 _dafny.Array = _1401_v26
				_ = _lhs138
				var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1401_v26), 0))
				_ = _lhs139
				var _lhs140 *C6 = _this
				_ = _lhs140
				var _lhs141 *GlobalState = globalState
				_ = _lhs141
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				(_lhs138).ArraySet1(_rhs178, (_lhs139).Int())
				_lhs140.F29_set_(_rhs179)
				_lhs141.F13 = _rhs180
				_lhs142.F13 = _rhs181
				_1402_v27 = _rhs182
				var _1407_v30 _dafny.Array
				_ = _1407_v30
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_28
				var _nw228 _dafny.Array
				_ = _nw228
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw228 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Sequence = func(_1408_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jhgvtevm"), _dafny.UnicodeSeqOfUtf8Bytes("chiwr"))
					}
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw228 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw228).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw228).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1407_v30 = _nw228
				var _1409_v31 _dafny.Sequence
				_ = _1409_v31
				_1409_v31 = _dafny.UnicodeSeqOfUtf8Bytes("yguw")
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_1407_v30), 0))
				_ = _index180
				(_1407_v30).ArraySet1((func() _dafny.Sequence {
					if p0 {
						return _1409_v31
					}
					return _1409_v31
				})(), (_index180).Int())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_1407_v30), 0))
				_ = _index181
				(_1407_v30).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1409_v31, _dafny.UnicodeSeqOfUtf8Bytes("xhyw")), (_index181).Int())
				_1363_cf41 = _1363_cf41
				_1363_cf41 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(554))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}(func(_1410_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				})), _1409_v31)).Cardinality())
			}
			var _1411_v33 _dafny.MultiSet
			_ = _1411_v33
			_1411_v33 = _dafny.MultiSetOf(_1363_cf41)
			var _1412_v34 D4
			_ = _1412_v34
			_1412_v34 = Companion_D4_.Create_DC10_(_1411_v33)
			var _1413_v35 _dafny.Map
			_ = _1413_v35
			_1413_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1412_v34).Dtor_cf22(), true)
			var _1414_v36 _dafny.Sequence
			_ = _1414_v36
			_1414_v36 = _dafny.SeqOf((func() _dafny.Map {
				var _coll40 = _dafny.NewMapBuilder()
				_ = _coll40
				for _iter42 := _dafny.Iterate((_1413_v35).Keys().Elements()); ; {
					_compr_40, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _1415_v32 _dafny.MultiSet
					_1415_v32 = interface{}(_compr_40).(_dafny.MultiSet)
					if (_1413_v35).Contains(_1415_v32) {
						_coll40.Add(_1415_v32, _this.F29())
					}
				}
				return _coll40.ToMap()
			}()).Cardinality(), _1363_cf41)
			var _1416_v37 _dafny.Set
			_ = _1416_v37
			_1416_v37 = _dafny.SetOf((_this).F34())
			(globalState).F13 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1414_v36, (Companion_Default___.SafeIndex((_1416_v37).Cardinality(), _dafny.IntOfUint32((_1414_v36).Cardinality()))).Uint32(), ((_this).F30()).Plus(_1363_cf41))).Cardinality())
			(globalState).F13 = (_dafny.Zero).Minus((_this).F30())
		} else if _source22.Is_DC23() {
			var _1417___mcc_h2 bool = _source22.Get_().(D8_DC23).Cf43
			_ = _1417___mcc_h2
			var _1418___mcc_h3 bool = _source22.Get_().(D8_DC23).Cf44
			_ = _1418___mcc_h3
			var _1419___mcc_h4 bool = _source22.Get_().(D8_DC23).Cf45
			_ = _1419___mcc_h4
			var _1420_cf45 bool = _1419___mcc_h4
			_ = _1420_cf45
			var _1421_cf44 bool = _1418___mcc_h3
			_ = _1421_cf44
			var _1422_cf43 bool = _1417___mcc_h2
			_ = _1422_cf43
			var _1423_v38 _dafny.Array
			_ = _1423_v38
			var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
			_ = _nw229
			_1423_v38 = _nw229
			var _1424_v39 *C0
			_ = _1424_v39
			var _nw230 *C0 = New_C0_()
			_ = _nw230
			_nw230.Ctor__(_1423_v38)
			_1424_v39 = _nw230
			var _1425_v40 _dafny.Map
			_ = _1425_v40
			_1425_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _1424_v39)
			var _1426_v41 _dafny.Sequence
			_ = _1426_v41
			_1426_v41 = _dafny.SeqOf(_dafny.IntOfInt64(835))
			var _1427_v42 _dafny.Sequence
			_ = _1427_v42
			_1427_v42 = _dafny.SeqOf(_dafny.IntOfInt64(161), _dafny.IntOfUint32((_1426_v41).Cardinality()))
			var _1428_v43 _dafny.Set
			_ = _1428_v43
			_1428_v43 = _dafny.SetOf(_1427_v42)
			var _1429_v44 _dafny.Map
			_ = _1429_v44
			_1429_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1428_v43).Cardinality(), _1421_cf44)
			var _1430_v45 _dafny.Set
			_ = _1430_v45
			_1430_v45 = _dafny.SetOf(true)
			var _1431_v46 _dafny.Map
			_ = _1431_v46
			_1431_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1429_v44, _1430_v45)
			var _1432_v47 _dafny.Sequence
			_ = _1432_v47
			_1432_v47 = _dafny.SeqOf((func() _dafny.Set {
				if (_1431_v46).Contains(_1429_v44) {
					return (_1431_v46).Get(_1429_v44).(_dafny.Set)
				}
				return _1430_v45
			})(), _dafny.SetOf(_1421_cf44, _1420_cf45), _1430_v45)
			var _1433_v48 _dafny.Array
			_ = _1433_v48
			var _nwElement0_49 *C0 = _1424_v39
			_ = _nwElement0_49
			var _nw231 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(28))
			_ = _nw231
			(_nw231).ArraySet1(_nwElement0_49, 0)
			(_nw231).ArraySet1(_1424_v39, 1)
			(_nw231).ArraySet1(_1424_v39, 2)
			(_nw231).ArraySet1(_1424_v39, 3)
			(_nw231).ArraySet1(_1424_v39, 4)
			(_nw231).ArraySet1(_1424_v39, 5)
			(_nw231).ArraySet1(_1424_v39, 6)
			(_nw231).ArraySet1(_1424_v39, 7)
			(_nw231).ArraySet1(_1424_v39, 8)
			(_nw231).ArraySet1(_1424_v39, 9)
			(_nw231).ArraySet1(_1424_v39, 10)
			(_nw231).ArraySet1(_1424_v39, 11)
			(_nw231).ArraySet1(_1424_v39, 12)
			(_nw231).ArraySet1(_1424_v39, 13)
			(_nw231).ArraySet1(_1424_v39, 14)
			(_nw231).ArraySet1(_1424_v39, 15)
			(_nw231).ArraySet1(_1424_v39, 16)
			(_nw231).ArraySet1(_1424_v39, 17)
			(_nw231).ArraySet1(_1424_v39, 18)
			(_nw231).ArraySet1((func() *C0 {
				if (_1425_v40).Contains(_dafny.IntOfUint32((_1432_v47).Cardinality())) {
					return (_1425_v40).Get(_dafny.IntOfUint32((_1432_v47).Cardinality())).(*C0)
				}
				return _1424_v39
			})(), 19)
			(_nw231).ArraySet1(_1424_v39, 20)
			(_nw231).ArraySet1(_1424_v39, 21)
			(_nw231).ArraySet1(_1424_v39, 22)
			(_nw231).ArraySet1(_1424_v39, 23)
			(_nw231).ArraySet1(_1424_v39, 24)
			(_nw231).ArraySet1(_1424_v39, 25)
			(_nw231).ArraySet1(_1424_v39, 26)
			(_nw231).ArraySet1(_1424_v39, 27)
			_1433_v48 = _nw231
			_1433_v48 = _1433_v48
			r2 = (_dafny.IntOfUint32(((_this).F32()).Cardinality())).Minus((_this).F30())
			(globalState).F13 = _dafny.IntOfInt64(262)
			(globalState).F12 = _1422_cf43
		} else if _source22.Is_DC24() {
			var _1434___mcc_h5 bool = _source22.Get_().(D8_DC24).Cf46
			_ = _1434___mcc_h5
			var _1435___mcc_h6 _dafny.Int = _source22.Get_().(D8_DC24).Cf47
			_ = _1435___mcc_h6
			var _1436___mcc_h7 _dafny.Sequence = _source22.Get_().(D8_DC24).Cf48
			_ = _1436___mcc_h7
			var _1437___mcc_h8 bool = _source22.Get_().(D8_DC24).Cf49
			_ = _1437___mcc_h8
			var _1438___mcc_h9 _dafny.Sequence = _source22.Get_().(D8_DC24).Cf50
			_ = _1438___mcc_h9
			var _1439_cf50 _dafny.Sequence = _1438___mcc_h9
			_ = _1439_cf50
			var _1440_cf49 bool = _1437___mcc_h8
			_ = _1440_cf49
			var _1441_cf48 _dafny.Sequence = _1436___mcc_h7
			_ = _1441_cf48
			var _1442_cf47 _dafny.Int = _1435___mcc_h6
			_ = _1442_cf47
			var _1443_cf46 bool = _1434___mcc_h5
			_ = _1443_cf46
			var _1444_v49 _dafny.Sequence
			_ = _1444_v49
			_1444_v49 = _dafny.UnicodeSeqOfUtf8Bytes("c")
			var _1445_v50 _dafny.Map
			_ = _1445_v50
			_1445_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F34()), _this.F36())
			var _1446_v51 _dafny.Array
			_ = _1446_v51
			var _nwElement0_50 _dafny.CodePoint = _this.F36()
			_ = _nwElement0_50
			var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(16))
			_ = _nw232
			(_nw232).ArraySet1CodePoint(_nwElement0_50, 0)
			(_nw232).ArraySet1CodePoint(_this.F36(), 1)
			(_nw232).ArraySet1CodePoint(_this.F36(), 2)
			(_nw232).ArraySet1CodePoint(_this.F36(), 3)
			(_nw232).ArraySet1CodePoint(_dafny.CodePoint('l'), 4)
			(_nw232).ArraySet1CodePoint(_this.F36(), 5)
			(_nw232).ArraySet1CodePoint(_this.F36(), 6)
			(_nw232).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1445_v50).Contains(_dafny.IntOfUint32((_1444_v49).Cardinality())) {
					return (_1445_v50).Get(_dafny.IntOfUint32((_1444_v49).Cardinality())).(_dafny.CodePoint)
				}
				return _this.F36()
			})(), 7)
			(_nw232).ArraySet1CodePoint(_dafny.CodePoint('c'), 8)
			(_nw232).ArraySet1CodePoint(_this.F36(), 9)
			(_nw232).ArraySet1CodePoint(_this.F36(), 10)
			(_nw232).ArraySet1CodePoint(_this.F36(), 11)
			(_nw232).ArraySet1CodePoint(_this.F36(), 12)
			(_nw232).ArraySet1CodePoint(_this.F36(), 13)
			(_nw232).ArraySet1CodePoint(_this.F36(), 14)
			(_nw232).ArraySet1CodePoint(_dafny.CodePoint('f'), 15)
			_1446_v51 = _nw232
			var _1447_v52 *C0
			_ = _1447_v52
			var _nw233 *C0 = New_C0_()
			_ = _nw233
			_nw233.Ctor__(_1446_v51)
			_1447_v52 = _nw233
			var _source24 D1 = Companion_D1_.Create_DC3_(_this.F35(), (_1439_cf50).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1444_v49).Cardinality()), _dafny.IntOfUint32((_1439_cf50).Cardinality()))).Uint32()).(_dafny.Int), _1447_v52, Companion_Default___.SafeModuloInt(_1442_cf47, _1442_cf47))
			_ = _source24
			if _source24.Is_DC3() {
				var _1448___mcc_h18 bool = _source24.Get_().(D1_DC3).Cf3
				_ = _1448___mcc_h18
				var _1449___mcc_h19 _dafny.Int = _source24.Get_().(D1_DC3).Cf4
				_ = _1449___mcc_h19
				var _1450___mcc_h20 *C0 = _source24.Get_().(D1_DC3).Cf5
				_ = _1450___mcc_h20
				var _1451___mcc_h21 _dafny.Int = _source24.Get_().(D1_DC3).Cf6
				_ = _1451___mcc_h21
				var _1452_cf6 _dafny.Int = _1451___mcc_h21
				_ = _1452_cf6
				var _1453_cf5 *C0 = _1450___mcc_h20
				_ = _1453_cf5
				var _1454_cf4 _dafny.Int = _1449___mcc_h19
				_ = _1454_cf4
				var _1455_cf3 bool = _1448___mcc_h18
				_ = _1455_cf3
				var _1456_v53 *C0
				_ = _1456_v53
				var _nw234 *C0 = New_C0_()
				_ = _nw234
				_nw234.Ctor__(_1447_v52.F37)
				_1456_v53 = _nw234
				var _1457_v54 T0
				_ = _1457_v54
				var _nw235 *C1 = New_C1_()
				_ = _nw235
				_nw235.Ctor__(_1440_cf49, false, true)
				_1457_v54 = _nw235
				var _1458_v55 _dafny.Map
				_ = _1458_v55
				_1458_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v54, _this.F36())
				_1454_cf4 = (Companion_Default___.SafeModuloInt((_1458_v55).Cardinality(), _1452_cf6)).Minus(_1454_cf4)
				var _1459_v56 _dafny.Array
				_ = _1459_v56
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_29
				var _nw236 _dafny.Array
				_ = _nw236
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw236 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Int = func(_1460_i4 _dafny.Int) _dafny.Int {
						return (_1460_i4).Minus((_this).F34())
					}
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw236 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw236).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw236).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1459_v56 = _nw236
				var _1461_v57 _dafny.Map
				_ = _1461_v57
				_1461_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _this.F36())
				var _1462_v58 _dafny.Map
				_ = _1462_v58
				_1462_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1459_v56, (_1461_v57).Cardinality())
				var _1463_v59 _dafny.MultiSet
				_ = _1463_v59
				_1463_v59 = _dafny.MultiSetOf(_1457_v54.F28())
				var _1464_v60 *C2
				_ = _1464_v60
				var _nw237 *C2 = New_C2_()
				_ = _nw237
				_nw237.Ctor__((_this).F30(), (_1462_v58).Cardinality(), (_1463_v59).IsDisjointFrom(_dafny.MultiSetOf(_1457_v54.F28(), _this.F28())), !(!(_1457_v54.F28())), !(_1457_v54.F28()))
				_1464_v60 = _nw237
				r3 = (_this).Fm27(((_this).F32()).Select((Companion_Default___.SafeIndex(_1452_cf6, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), (_dafny.IntOfUint32((_1439_cf50).Cardinality())).Times(_1454_cf4), _1454_cf4, globalState)
			} else if _source24.Is_DC2() {
				var _1465___mcc_h22 _dafny.Array = _source24.Get_().(D1_DC2).Cf2
				_ = _1465___mcc_h22
				var _1466_cf2 _dafny.Array = _1465___mcc_h22
				_ = _1466_cf2
				(_this).F36_set_(_this.F36())
				var _1467_v61 *C0
				_ = _1467_v61
				var _nw238 *C0 = New_C0_()
				_ = _nw238
				_nw238.Ctor__(_1447_v52.F37)
				_1467_v61 = _nw238
				r1 = !((_1441_cf48).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1441_cf48).Cardinality()))).Uint32()).(bool))
				(globalState).F13 = (_this).F30()
			} else {
				var _1468___mcc_h23 D1 = _source24.Get_().(D1_DC4).Cf7
				_ = _1468___mcc_h23
				var _1469_cf7 D1 = _1468___mcc_h23
				_ = _1469_cf7
				var _1470_v62 *C0
				_ = _1470_v62
				var _nw239 *C0 = New_C0_()
				_ = _nw239
				_nw239.Ctor__(_1447_v52.F37)
				_1470_v62 = _nw239
				r2 = Companion_Default___.SafeModuloInt((_this).F34(), _1442_cf47)
				var _1471_v63 _dafny.Map
				_ = _1471_v63
				_1471_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), !(_1443_cf46))
				var _1472_v64 _dafny.Map
				_ = _1472_v64
				_1472_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_1471_v63).Cardinality())
				var _1473_v65 _dafny.Array
				_ = _1473_v65
				var _nwElement0_51 _dafny.Sequence = _1441_cf48
				_ = _nwElement0_51
				var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(23))
				_ = _nw240
				(_nw240).ArraySet1(_nwElement0_51, 0)
				(_nw240).ArraySet1((_this).F32(), 1)
				(_nw240).ArraySet1(_1441_cf48, 2)
				(_nw240).ArraySet1((_this).F32(), 3)
				(_nw240).ArraySet1(_1441_cf48, 4)
				(_nw240).ArraySet1((_this).F32(), 5)
				(_nw240).ArraySet1(_1441_cf48, 6)
				(_nw240).ArraySet1(Companion_Default___.Fm0(_1442_cf47, _this.F35(), p0, (_1472_v64).Cardinality(), globalState), 7)
				(_nw240).ArraySet1((_this).F32(), 8)
				(_nw240).ArraySet1((_this).F32(), 9)
				(_nw240).ArraySet1(_1441_cf48, 10)
				(_nw240).ArraySet1((_this).F32(), 11)
				(_nw240).ArraySet1(_dafny.SeqOf(_this.F35(), _1440_cf49, _1443_cf46, _this.F31(), _this.F29()), 12)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), p0), (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), p0)).Cardinality()))).Uint32(), _this.F35()), 13)
				(_nw240).ArraySet1(_1441_cf48, 14)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), !(_1443_cf46)), 15)
				(_nw240).ArraySet1((_this).F32(), 16)
				(_nw240).ArraySet1(_dafny.SeqOf(p0), 17)
				(_nw240).ArraySet1((_this).F32(), 18)
				(_nw240).ArraySet1((_this).F32(), 19)
				(_nw240).ArraySet1(_dafny.SeqOf(_this.F35()), 20)
				(_nw240).ArraySet1(_1441_cf48, 21)
				(_nw240).ArraySet1(_dafny.SeqOf(p0, _1443_cf46, true, _this.F28(), false), 22)
				_1473_v65 = _nw240
				var _1474_v66 _dafny.Set
				_ = _1474_v66
				_1474_v66 = _dafny.SetOf(_this.F31(), _1440_cf49)
				var _1475_v67 D2
				_ = _1475_v67
				_1475_v67 = Companion_D2_.Create_DC6_(_1442_cf47, false, (_this).F30(), _1443_cf46, _1442_cf47)
				var _1476_v68 *C3
				_ = _1476_v68
				var _nw241 *C3 = New_C3_()
				_ = _nw241
				_nw241.Ctor__((_this).F34(), _1473_v65, _this.F36(), (_this).F34(), (Companion_Default___.Fm1(_this.F35(), (Companion_Default___.Fm38(false, (_1474_v66).Cardinality(), globalState)).Cardinality(), globalState)) && (_1440_cf49), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(((_this).F33()).Select((Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_this.F36())).Cardinality()), _dafny.IntOfUint32((((_this).F33()).Select((Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), p0), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(((_this).F33()).Select((Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_this.F36())).Cardinality()), _dafny.IntOfUint32((((_this).F33()).Select((Companion_Default___.SafeIndex(_1442_cf47, _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), p0)).Cardinality()))).Uint32(), _this.F35()), (_this).F33(), (_this).F34(), _this.F35(), (_1475_v67).Dtor_cf10(), p0)
				_1476_v68 = _nw241
				r3 = (_1442_cf47).Times((_1476_v68).F44())
			}
			if _1443_cf46 {
				var _1477_v69 _dafny.Map
				_ = _1477_v69
				_1477_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), !(_1440_cf49))
				var _1478_v70 _dafny.Set
				_ = _1478_v70
				_1478_v70 = _dafny.SetOf(_1477_v69)
				var _1479_v71 D7
				_ = _1479_v71
				_1479_v71 = Companion_D7_.Create_DC19_(_this.F29(), (_dafny.Zero).Minus(_1442_cf47), (_this).F30())
				_1478_v70 = (func() _dafny.Set {
					if (_1479_v71).Dtor_cf36() {
						return (_1478_v70).Difference(_1478_v70)
					}
					return _1478_v70
				})()
				var _1480_v72 _dafny.Map
				_ = _1480_v72
				_1480_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1442_cf47, _dafny.SetOf((_this).F34()))
				var _1481_v73 _dafny.Set
				_ = _1481_v73
				_1481_v73 = _dafny.SetOf(_this.F35(), _this.F31())
				var _1482_v74 _dafny.Array
				_ = _1482_v74
				var _nwElement0_52 _dafny.Int = (_this).F34()
				_ = _nwElement0_52
				var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(20))
				_ = _nw242
				(_nw242).ArraySet1(_nwElement0_52, 0)
				(_nw242).ArraySet1((_this).F34(), 1)
				(_nw242).ArraySet1((_dafny.Zero).Minus(_1442_cf47), 2)
				(_nw242).ArraySet1((_this).F30(), 3)
				(_nw242).ArraySet1((_1480_v72).Cardinality(), 4)
				(_nw242).ArraySet1((_this).F30(), 5)
				(_nw242).ArraySet1((_this).F34(), 6)
				(_nw242).ArraySet1(_1442_cf47, 7)
				(_nw242).ArraySet1((_1481_v73).Cardinality(), 8)
				(_nw242).ArraySet1(_1442_cf47, 9)
				(_nw242).ArraySet1((_this).F34(), 10)
				(_nw242).ArraySet1((_dafny.Zero).Minus((_this).F30()), 11)
				(_nw242).ArraySet1((_this).F30(), 12)
				(_nw242).ArraySet1((_this).F30(), 13)
				(_nw242).ArraySet1(_dafny.IntOfInt64(300), 14)
				(_nw242).ArraySet1(_1442_cf47, 15)
				(_nw242).ArraySet1((_this).F30(), 16)
				(_nw242).ArraySet1((_this).F34(), 17)
				(_nw242).ArraySet1((_this).F34(), 18)
				(_nw242).ArraySet1(_dafny.IntOfInt64(162), 19)
				_1482_v74 = _nw242
				var _1483_v75 _dafny.Sequence
				_ = _1483_v75
				_1483_v75 = _dafny.SeqOf(_1482_v74)
				var _1484_v76 _dafny.Map
				_ = _1484_v76
				_1484_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).F34(), (_this).F30()), (_this).Fm27(_this.F31(), _dafny.IntOfUint32((_1483_v75).Cardinality()), (_this).Fm28((_this).F34(), (_this).Fm7(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F34())), _1442_cf47, globalState), globalState), globalState))
				_1484_v76 = (_1484_v76).Update((_dafny.Zero).Minus(_1442_cf47), Companion_Default___.SafeModuloInt(_1442_cf47, _1442_cf47))
				r3 = (func() _dafny.Int {
					if _this.F35() {
						return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1444_v49).Cardinality()), (_this).F34())
					}
					return (_dafny.IntOfInt64(636)).Minus((_this).F30())
				})()
				var _1485_v77 _dafny.Array
				_ = _1485_v77
				var _nw243 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
				_ = _nw243
				_1485_v77 = _nw243
				var _1486_v78 _dafny.MultiSet
				_ = _1486_v78
				_1486_v78 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1442_cf47))
				var _1487_v79 *C1
				_ = _1487_v79
				var _nw244 *C1 = New_C1_()
				_ = _nw244
				_nw244.Ctor__(_this.F31(), _this.F35(), (_this).Fm7(_1486_v78, (_this).F34(), globalState))
				_1487_v79 = _nw244
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1485_v77), 0))
				_ = _index182
				(_1485_v77).ArraySet1(_1487_v79, (_index182).Int())
				var _1488_v80 _dafny.Array
				_ = _1488_v80
				var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
				_ = _nw245
				_1488_v80 = _nw245
				var _1489_v81 _dafny.Array
				_ = _1489_v81
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_30
				var _nw246 _dafny.Array
				_ = _nw246
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw246 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) bool = func(_1490_i5 _dafny.Int) bool {
						return _this.F29()
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw246 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw246).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw246).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1489_v81 = _nw246
				var _1491_v82 _dafny.Map
				_ = _1491_v82
				_1491_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1484_v76).Contains((_this).F30()) {
						return (_1484_v76).Get((_this).F30()).(_dafny.Int)
					}
					return (_this).F30()
				})(), _1489_v81)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1488_v80), 0))
				_ = _index183
				(_1488_v80).ArraySet1(_1491_v82, (_index183).Int())
				var _1492_v83 _dafny.Map
				_ = _1492_v83
				_1492_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1487_v79)
				var _1493_v84 _dafny.Map
				_ = _1493_v84
				_1493_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _1491_v82)
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1485_v77), 0))
				_ = _index184
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1488_v80), 0))
				_ = _index185
				var _rhs183 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F34(), _1442_cf47)
				_ = _rhs183
				var _rhs184 *C1 = (func() *C1 {
					if (_1492_v83).Contains(_this.F29()) {
						return (_1492_v83).Get(_this.F29()).(*C1)
					}
					return (Companion_D10_.Create_DC28_(_1487_v79)).Dtor_cf55()
				})()
				_ = _rhs184
				var _rhs185 _dafny.Map = (((func() _dafny.Map {
					if (_1493_v84).Contains((_1487_v79).F43()) {
						return (_1493_v84).Get((_1487_v79).F43()).(_dafny.Map)
					}
					return _1491_v82
				})()).Merge(_1491_v82)).Merge(_1491_v82)
				_ = _rhs185
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 _dafny.Array = _1485_v77
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1485_v77), 0))
				_ = _lhs145
				var _lhs146 _dafny.Array = _1488_v80
				_ = _lhs146
				var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_1488_v80), 0))
				_ = _lhs147
				_lhs143.F13 = _rhs183
				(_lhs144).ArraySet1(_rhs184, (_lhs145).Int())
				(_lhs146).ArraySet1(_rhs185, (_lhs147).Int())
				var _1494_v85 _dafny.Array
				_ = _1494_v85
				var _nwElement0_53 _dafny.Array = _1489_v81
				_ = _nwElement0_53
				var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(11))
				_ = _nw247
				(_nw247).ArraySet1(_nwElement0_53, 0)
				(_nw247).ArraySet1(_1489_v81, 1)
				(_nw247).ArraySet1(_1489_v81, 2)
				(_nw247).ArraySet1(_1489_v81, 3)
				(_nw247).ArraySet1(_1489_v81, 4)
				(_nw247).ArraySet1(_1489_v81, 5)
				(_nw247).ArraySet1(_1489_v81, 6)
				(_nw247).ArraySet1(_1489_v81, 7)
				(_nw247).ArraySet1(_1489_v81, 8)
				(_nw247).ArraySet1((func() _dafny.Array {
					if _this.F31() {
						return _1489_v81
					}
					return _1489_v81
				})(), 9)
				(_nw247).ArraySet1(_1489_v81, 10)
				_1494_v85 = _nw247
				var _1495_v86 _dafny.Sequence
				_ = _1495_v86
				_1495_v86 = _dafny.SeqOf(_1489_v81, _1489_v81)
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1494_v85), 0))
				_ = _index186
				(_1494_v85).ArraySet1((_1495_v86).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1444_v49).Cardinality())), _dafny.IntOfUint32((_1495_v86).Cardinality()))).Uint32()).(_dafny.Array), (_index186).Int())
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1494_v85), 0))
				_ = _index187
				(_1494_v85).ArraySet1(_1489_v81, (_index187).Int())
			} else {
				(globalState).F1 = _this.F28()
				(globalState).F1 = Companion_Default___.Fm1(false, (_this).Fm28((_this).F30(), _1440_cf49, globalState), globalState)
				var _1496_v87 _dafny.Map
				_ = _1496_v87
				_1496_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_cf46, !(_this.F31()))
				_1496_v87 = (_1496_v87).Update(_this.F29(), ((_this).F34()).Cmp((_this).F34()) >= 0)
				var _1497_v88 *C0
				_ = _1497_v88
				var _nw248 *C0 = New_C0_()
				_ = _nw248
				_nw248.Ctor__(_1447_v52.F37)
				_1497_v88 = _nw248
				(globalState).F26 = (_this.F31()) && (_this.F31())
			}
			var _1498_v89 D10
			_ = _1498_v89
			_1498_v89 = Companion_D10_.Create_DC29_(_1442_cf47, Companion_Default___.Fm1(_this.F29(), _1442_cf47, globalState))
			(_this).F28_set_((_1498_v89).Dtor_cf57())
			(globalState).F1 = !(_this.F35()) || (_1440_cf49)
		} else if _source22.Is_DC21() {
			var _1499___mcc_h10 _dafny.Map = _source22.Get_().(D8_DC21).Cf40
			_ = _1499___mcc_h10
			var _1500_cf40 _dafny.Map = _1499___mcc_h10
			_ = _1500_cf40
			var _1501_v90 _dafny.Sequence
			_ = _1501_v90
			_1501_v90 = _dafny.UnicodeSeqOfUtf8Bytes("ggytjvvm")
			var _1502_v91 _dafny.Array
			_ = _1502_v91
			var _nwElement0_54 _dafny.Sequence = _1501_v90
			_ = _nwElement0_54
			var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(29))
			_ = _nw249
			(_nw249).ArraySet1(_nwElement0_54, 0)
			(_nw249).ArraySet1(_1501_v90, 1)
			(_nw249).ArraySet1((func() _dafny.Sequence {
				if false {
					return _1501_v90
				}
				return _1501_v90
			})(), 2)
			(_nw249).ArraySet1(_1501_v90, 3)
			(_nw249).ArraySet1(_1501_v90, 4)
			(_nw249).ArraySet1(_1501_v90, 5)
			(_nw249).ArraySet1(_1501_v90, 6)
			(_nw249).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1501_v90, _1501_v90), 7)
			(_nw249).ArraySet1(_1501_v90, 8)
			(_nw249).ArraySet1(_1501_v90, 9)
			(_nw249).ArraySet1(_dafny.Companion_Sequence_.Update(_1501_v90, (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1501_v90).Cardinality()))).Uint32(), _this.F36()), 10)
			(_nw249).ArraySet1(_1501_v90, 11)
			(_nw249).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nqoj"), 12)
			(_nw249).ArraySet1(_1501_v90, 13)
			(_nw249).ArraySet1(_1501_v90, 14)
			(_nw249).ArraySet1(_1501_v90, 15)
			(_nw249).ArraySet1(_1501_v90, 16)
			(_nw249).ArraySet1(_1501_v90, 17)
			(_nw249).ArraySet1(_1501_v90, 18)
			(_nw249).ArraySet1(_1501_v90, 19)
			(_nw249).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1501_v90, _1501_v90), 20)
			(_nw249).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-167))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg98 _dafny.Int) interface{} {
					return coer98(arg98)
				}
			}(func(_1503_i6 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			})), 21)
			(_nw249).ArraySet1(_1501_v90, 22)
			(_nw249).ArraySet1((_this).Fm3(_this.F29(), p0, globalState), 23)
			(_nw249).ArraySet1(_1501_v90, 24)
			(_nw249).ArraySet1(_1501_v90, 25)
			(_nw249).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lskefdq"), 26)
			(_nw249).ArraySet1(_1501_v90, 27)
			(_nw249).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1501_v90, _dafny.UnicodeSeqOfUtf8Bytes("k")), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F34()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1501_v90, _dafny.UnicodeSeqOfUtf8Bytes("k"))).Cardinality()))).Uint32(), _this.F36()), 28)
			_1502_v91 = _nw249
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))
			_ = _index188
			(_1502_v91).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1501_v90, _1501_v90), (_index188).Int())
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))
			_ = _index189
			var _rhs186 _dafny.Sequence = _1501_v90
			_ = _rhs186
			var _rhs187 _dafny.Int = _dafny.IntOfInt64(852)
			_ = _rhs187
			var _lhs148 _dafny.Array = _1502_v91
			_ = _lhs148
			var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))
			_ = _lhs149
			var _lhs150 *GlobalState = globalState
			_ = _lhs150
			(_lhs148).ArraySet1(_rhs186, (_lhs149).Int())
			_lhs150.F13 = _rhs187
			if _this.F35() {
				(_this).F36_set_(_this.F36())
				var _1504_v92 _dafny.Array
				_ = _1504_v92
				var _nw250 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw250
				_1504_v92 = _nw250
				var _1505_v93 _dafny.Sequence
				_ = _1505_v93
				_1505_v93 = _dafny.SeqOf((_this).F30(), (_dafny.Zero).Minus((_this).F34()))
				var _1506_v94 _dafny.Map
				_ = _1506_v94
				_1506_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1505_v93, true)
				var _1507_v95 _dafny.Array
				_ = _1507_v95
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_31
				var _nw251 _dafny.Array
				_ = _nw251
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw251 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.CodePoint = func(_1508_i7 _dafny.Int) _dafny.CodePoint {
						return _this.F36()
					}
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw251 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw251).ArraySet1CodePoint(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw251).ArraySet1CodePoint(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1507_v95 = _nw251
				var _1509_v96 *C0
				_ = _1509_v96
				var _nw252 *C0 = New_C0_()
				_ = _nw252
				_nw252.Ctor__(_1507_v95)
				_1509_v96 = _nw252
				var _1510_v97 _dafny.Sequence
				_ = _1510_v97
				_1510_v97 = _dafny.SeqOf(_1509_v96)
				var _1511_v98 D9
				_ = _1511_v98
				_1511_v98 = Companion_D9_.Create_DC27_((_1510_v97).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1510_v97).Cardinality()))).Uint32()).(*C0), _1501_v90)
				var _1512_v99 _dafny.MultiSet
				_ = _1512_v99
				_1512_v99 = _dafny.MultiSetOf(_1511_v98)
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1504_v92), 0))
				_ = _index190
				(_1504_v92).ArraySet1(((_1506_v94).Cardinality()).Cmp((_1512_v99).Cardinality()) > 0, (_index190).Int())
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1504_v92), 0))
				_ = _index191
				(_1504_v92).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1501_v90, (_1502_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))).Int()).(_dafny.Sequence)), (_index191).Int())
				var _1513_v100 _dafny.Array
				_ = _1513_v100
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_32
				var _nw253 _dafny.Array
				_ = _nw253
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw253 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) D8 = func(_1514_i8 _dafny.Int) D8 {
						return Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetFromSeq((_this).F32())), _this.F35()))
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw253 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw253).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw253).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1513_v100 = _nw253
				var _1515_v101 _dafny.Map
				_ = _1515_v101
				_1515_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _this.F31())
				var _1516_v102 _dafny.Map
				_ = _1516_v102
				_1516_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1504_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1504_v92), 0))).Int()).(bool))
				var _1517_v103 _dafny.Map
				_ = _1517_v103
				_1517_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1515_v101).Cardinality(), _1516_v102)
				var _1518_v104 _dafny.Map
				_ = _1518_v104
				_1518_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1513_v100, ((_1517_v103).Merge(_1517_v103)).Cardinality())
				_1518_v104 = _1518_v104
				var _1519_v105 _dafny.MultiSet
				_ = _1519_v105
				_1519_v105 = _dafny.MultiSetOf((_this).F30())
				var _1520_v106 T2
				_ = _1520_v106
				var _nw254 *C5 = New_C5_()
				_ = _nw254
				_nw254.Ctor__(_this.F29(), (_this).F34(), (_this).F34(), _this.F35(), Companion_Default___.Fm0((_1519_v105).Cardinality(), p0, _this.F29(), (_this).F34(), globalState), (Companion_D14_.Create_DC36_(Companion_Default___.Fm51(p0, _this.F36(), globalState))).Dtor_cf67(), (_1515_v101).Cardinality(), _this.F28(), _this.F29(), p0)
				_1520_v106 = _nw254
				var _1521_v107 _dafny.Map
				_ = _1521_v107
				_1521_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _1520_v106)
				var _1522_v108 _dafny.Sequence
				_ = _1522_v108
				_1522_v108 = _dafny.SeqOf(_1520_v106.F29(), true, _1520_v106.F28(), _this.F28(), (_1504_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1504_v92), 0))).Int()).(bool))
				var _1523_v109 _dafny.MultiSet
				_ = _1523_v109
				_1523_v109 = _dafny.MultiSetOf((_1520_v106).F32(), _1522_v108)
				var _rhs188 _dafny.Map = _1521_v107
				_ = _rhs188
				var _rhs189 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_1520_v106).F32(), (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_1520_v106).F32()).Cardinality()))).Uint32(), !_dafny.Companion_Sequence_.Contains((_1502_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))).Int()).(_dafny.Sequence), _this.F36()))
				_ = _rhs189
				var _rhs190 bool = (Companion_Default___.Fm52(Companion_D8_.Create_DC21_((_1500_cf40).Update(Companion_Default___.Fm36((_1520_v106).F30(), (_1519_v105).Cardinality(), _this.F29(), globalState), _1520_v106.F29())), globalState)).IsSubsetOf((_dafny.SetOf(_1522_v108, (_this).F32(), (_1520_v106).F32())).Difference(func() _dafny.Set {
					var _coll41 = _dafny.NewBuilder()
					_ = _coll41
					for _iter43 := _dafny.Iterate((_1523_v109).Elements()); ; {
						_compr_41, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _1524_v110 _dafny.Sequence
						_1524_v110 = interface{}(_compr_41).(_dafny.Sequence)
						if (_1523_v109).Contains(_1524_v110) {
							_coll41.Add(_1524_v110)
						}
					}
					return _coll41.ToSet()
				}()))
				_ = _rhs190
				var _lhs151 T2 = _1520_v106
				_ = _lhs151
				_1521_v107 = _rhs188
				_1522_v108 = _rhs189
				_lhs151.F29_set_(_rhs190)
				_1516_v102 = (_1516_v102).Update(!((func() bool {
					if (_1515_v101).Contains((_this).F34()) {
						return (_1515_v101).Get((_this).F34()).(bool)
					}
					return !(_this.F28())
				})()), p0)
			} else {
				var _1525_v111 _dafny.Array
				_ = _1525_v111
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw255
				_1525_v111 = _nw255
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))
				_ = _index192
				(_1525_v111).ArraySet1((_this).F30(), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))
				_ = _index193
				(_1525_v111).ArraySet1(((_this).F34()).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_this).F32()).Cardinality()), (_this).F34())), (_index193).Int())
				var _1526_v112 _dafny.Set
				_ = _1526_v112
				_1526_v112 = _dafny.SetOf((func() bool {
					if p0 {
						return _this.F29()
					}
					return _this.F28()
				})())
				var _1527_v113 _dafny.MultiSet
				_ = _1527_v113
				_1527_v113 = _dafny.MultiSetOf(_this.F28())
				var _1528_v114 _dafny.Sequence
				_ = _1528_v114
				_1528_v114 = _dafny.SeqOf((func() _dafny.Int {
					if (_1527_v113).Contains(_this.F35()) {
						return (_1527_v113).Multiplicity(_this.F35())
					}
					return (_this).F34()
				})(), (_1525_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))).Int()).(_dafny.Int))
				_1526_v112 = _dafny.SetOf((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pk")).Cardinality())).Cmp(_dafny.IntOfUint32((_1528_v114).Cardinality())) >= 0, _this.F28(), false)
				var _1529_v115 _dafny.Int
				_ = _1529_v115
				var _out57 _dafny.Int
				_ = _out57
				_out57 = Companion_Default___.M0(_1528_v114, _this.F29(), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-105)), (_this).F34()), _this.F35(), globalState)
				_1529_v115 = _out57
				var _1530_v116 _dafny.Map
				_ = _1530_v116
				_1530_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(66), _this.F28())
				var _1531_v117 _dafny.MultiSet
				_ = _1531_v117
				_1531_v117 = _dafny.MultiSetOf((_1525_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))).Int()).(_dafny.Int))
				var _1532_v118 _dafny.Map
				_ = _1532_v118
				_1532_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1525_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))).Int()).(_dafny.Int), (_this).F30())
				var _1533_v119 *C4
				_ = _1533_v119
				var _nw256 *C4 = New_C4_()
				_ = _nw256
				_nw256.Ctor__(_dafny.IntOfUint32(((_1502_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))).Int()).(_dafny.Sequence)).Cardinality()), (true) && ((func() bool {
					if (_1530_v116).Contains(_dafny.IntOfUint32(((_this).F32()).Cardinality())) {
						return (_1530_v116).Get(_dafny.IntOfUint32(((_this).F32()).Cardinality())).(bool)
					}
					return (_this).Fm7(_1531_v117, _1529_v115, globalState)
				})()), (_this).F32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg99 _dafny.Int) interface{} {
						return coer99(arg99)
					}
				}(func(_1534_i9 _dafny.Int) _dafny.Sequence {
					return (_this).F32()
				})), (_1527_v113).Cardinality(), _this.F35(), _this.F29(), (_dafny.IntOfUint32(((_1502_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))).Int()).(_dafny.Sequence)).Cardinality())).Cmp((_1532_v118).Cardinality()) >= 0)
				_1533_v119 = _nw256
				var _1535_v120 _dafny.Array
				_ = _1535_v120
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
				_ = _nw257
				_1535_v120 = _nw257
				var _1536_v121 *C0
				_ = _1536_v121
				var _nw258 *C0 = New_C0_()
				_ = _nw258
				_nw258.Ctor__(_1535_v120)
				_1536_v121 = _nw258
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))
				_ = _index194
				var _rhs191 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1501_v90, _dafny.UnicodeSeqOfUtf8Bytes("gobhvtjm"))
				_ = _rhs191
				var _rhs192 *C4 = _1533_v119
				_ = _rhs192
				var _rhs193 _dafny.Int = (Companion_D1_.Create_DC3_(((_this).F32()).Select((Companion_Default___.SafeIndex((_1525_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(302), _1536_v121, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Dtor_cf6()
				_ = _rhs193
				var _lhs152 _dafny.Array = _1502_v91
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))
				_ = _lhs153
				(_lhs152).ArraySet1(_rhs191, (_lhs153).Int())
				_1533_v119 = _rhs192
				r3 = _rhs193
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1525_v111), 0))
				_ = _index195
				(_1525_v111).ArraySet1(((_1526_v112).Intersection((_1526_v112).Difference(_1526_v112))).Cardinality(), (_index195).Int())
			}
			var _1537_v122 _dafny.Array
			_ = _1537_v122
			var _nw259 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw259
			_1537_v122 = _nw259
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_1537_v122), 0))
			_ = _index196
			(_1537_v122).ArraySet1((_this).F34(), (_index196).Int())
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_1537_v122), 0))
			_ = _index197
			var _rhs194 _dafny.Int = ((_dafny.IntOfInt64(-486)).Times(_dafny.IntOfUint32((_1501_v90).Cardinality()))).Minus(Companion_Default___.SafeDivisionInt((_this).F34(), (_this).F34()))
			_ = _rhs194
			var _rhs195 bool = _this.F29()
			_ = _rhs195
			var _rhs196 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1501_v90, (_1502_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1502_v91), 0))).Int()).(_dafny.Sequence)), _1501_v90)
			_ = _rhs196
			var _rhs197 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F30()), (_dafny.Zero).Minus((_this).F30()))
			_ = _rhs197
			var _lhs154 *C6 = _this
			_ = _lhs154
			var _lhs155 _dafny.Array = _1537_v122
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_1537_v122), 0))
			_ = _lhs156
			r2 = _rhs194
			_lhs154.F35_set_(_rhs195)
			_1501_v90 = _rhs196
			(_lhs155).ArraySet1(_rhs197, (_lhs156).Int())
			(_this).F36_set_(_dafny.CodePoint('a'))
		} else {
			var _1538___mcc_h11 D8 = _source22.Get_().(D8_DC25).Cf51
			_ = _1538___mcc_h11
			var _1539_cf51 D8 = _1538___mcc_h11
			_ = _1539_cf51
			(_this).F35_set_(!(_this.F31()) || (p0))
			r1 = !(true) || (!(p0))
			(_this).F28_set_(!(p0))
			(globalState).F26 = p0
		}
		if ((_this).F30()).Cmp((_this).F34()) == 0 {
			var _1540_v123 _dafny.Sequence
			_ = _1540_v123
			_1540_v123 = _dafny.SeqOf((_this).F34(), (_this).F30(), (_this).F30())
			var _1541_v124 _dafny.Int
			_ = _1541_v124
			var _out58 _dafny.Int
			_ = _out58
			_out58 = Companion_Default___.M0(_1540_v123, !(p0), (_this).F30(), _this.F35(), globalState)
			_1541_v124 = _out58
			var _1542_v125 _dafny.Sequence
			_ = _1542_v125
			_1542_v125 = _dafny.UnicodeSeqOfUtf8Bytes("lsw")
			_1542_v125 = _1542_v125
			var _1543_v126 _dafny.Map
			_ = _1543_v126
			_1543_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_this).F34())
			var _1544_v127 _dafny.Sequence
			_ = _1544_v127
			_1544_v127 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg100 _dafny.Int) interface{} {
					return coer100(arg100)
				}
			}(func(_1545_i10 _dafny.Int) _dafny.Int {
				return (_this).F34()
			})))
			var _1546_v128 D14
			_ = _1546_v128
			_1546_v128 = Companion_D14_.Create_DC36_((_this).F33())
			var _1547_v129 _dafny.MultiSet
			_ = _1547_v129
			_1547_v129 = _dafny.MultiSetOf(_1541_v124)
			var _1548_v130 _dafny.Sequence
			_ = _1548_v130
			_1548_v130 = _dafny.SeqOf(_1547_v129, _dafny.MultiSetOf((_this).F30(), (_this).F34()))
			var _1549_v131 *C5
			_ = _1549_v131
			var _nw260 *C5 = New_C5_()
			_ = _nw260
			_nw260.Ctor__(_this.F31(), (func() _dafny.Int {
				if (_1543_v126).Contains(_dafny.IntOfUint32(((_this).F32()).Cardinality())) {
					return (_1543_v126).Get(_dafny.IntOfUint32(((_this).F32()).Cardinality())).(_dafny.Int)
				}
				return Companion_Default___.Fm32(true, globalState)
			})(), (_this).F30(), _dafny.Companion_Sequence_.IsProperPrefixOf(_1540_v123, (_1544_v127).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1544_v127).Cardinality()))).Uint32()).(_dafny.Sequence)), (_this).F32(), (_1546_v128).Dtor_cf67(), Companion_Default___.SafeDivisionInt((_this).F34(), Companion_Default___.Fm46(((_1548_v130).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1548_v130).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), (_this).F34(), false, globalState)), Companion_Default___.Fm1(true, (_1540_v123).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1540_v123).Cardinality()), _dafny.IntOfUint32((_1540_v123).Cardinality()))).Uint32()).(_dafny.Int), globalState), ((_this).F32()).Select((Companion_Default___.SafeIndex(_1541_v124, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), (!(_this.F28())) || (p0))
			_1549_v131 = _nw260
			var _1550_v132 D4
			_ = _1550_v132
			_1550_v132 = Companion_D4_.Create_DC11_(_dafny.IntOfUint32((_1542_v125).Cardinality()), _this.F29(), (_1549_v131).F48())
			var _source25 D4 = _1550_v132
			_ = _source25
			if _source25.Is_DC11() {
				var _1551___mcc_h24 _dafny.Int = _source25.Get_().(D4_DC11).Cf23
				_ = _1551___mcc_h24
				var _1552___mcc_h25 bool = _source25.Get_().(D4_DC11).Cf24
				_ = _1552___mcc_h25
				var _1553___mcc_h26 _dafny.Int = _source25.Get_().(D4_DC11).Cf25
				_ = _1553___mcc_h26
				var _1554_cf25 _dafny.Int = _1553___mcc_h26
				_ = _1554_cf25
				var _1555_cf24 bool = _1552___mcc_h25
				_ = _1555_cf24
				var _1556_cf23 _dafny.Int = _1551___mcc_h24
				_ = _1556_cf23
				var _1557_v133 _dafny.Array
				_ = _1557_v133
				var _nwElement0_55 _dafny.CodePoint = _this.F36()
				_ = _nwElement0_55
				var _nw261 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(10))
				_ = _nw261
				(_nw261).ArraySet1CodePoint(_nwElement0_55, 0)
				(_nw261).ArraySet1CodePoint(_this.F36(), 1)
				(_nw261).ArraySet1CodePoint(_this.F36(), 2)
				(_nw261).ArraySet1CodePoint(_this.F36(), 3)
				(_nw261).ArraySet1CodePoint(_dafny.CodePoint('q'), 4)
				(_nw261).ArraySet1CodePoint(_this.F36(), 5)
				(_nw261).ArraySet1CodePoint(_this.F36(), 6)
				(_nw261).ArraySet1CodePoint(_this.F36(), 7)
				(_nw261).ArraySet1CodePoint(_this.F36(), 8)
				(_nw261).ArraySet1CodePoint(_dafny.CodePoint('q'), 9)
				_1557_v133 = _nw261
				var _1558_v134 *C0
				_ = _1558_v134
				var _nw262 *C0 = New_C0_()
				_ = _nw262
				_nw262.Ctor__(_1557_v133)
				_1558_v134 = _nw262
				var _1559_v135 _dafny.Map
				_ = _1559_v135
				_1559_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v125, _1558_v134)
				_1559_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1542_v125, (func() *C0 {
					if (_1549_v131).F47() {
						return _1558_v134
					}
					return (Companion_D1_.Create_DC3_(p0, (_dafny.Zero).Minus((_this).F34()), _1558_v134, _1541_v124)).Dtor_cf5()
				})())
				var _1560_v136 _dafny.Map
				_ = _1560_v136
				_1560_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1549_v131).F47(), true)
				_1560_v136 = (_1560_v136).Update(_this.F28(), (_1541_v124).Cmp(_1556_cf23) < 0)
				_1542_v125 = _dafny.Companion_Sequence_.Concatenate(_1542_v125, _1542_v125)
				(globalState).F1 = ((_this).F32()).Select((Companion_Default___.SafeIndex(_1554_cf25, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)
			} else {
				var _1561___mcc_h27 _dafny.MultiSet = _source25.Get_().(D4_DC10).Cf22
				_ = _1561___mcc_h27
				var _1562_cf22 _dafny.MultiSet = _1561___mcc_h27
				_ = _1562_cf22
				var _1563_v137 _dafny.Map
				_ = _1563_v137
				_1563_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _this.F31())
				var _1564_v139 *C4
				_ = _1564_v139
				var _nw263 *C4 = New_C4_()
				_ = _nw263
				_nw263.Ctor__(_1541_v124, true, _dafny.SeqOf(!(_this.F29())), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F33(), (Companion_Default___.SafeIndex(_1541_v124, _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32(), (_this).F32()), _dafny.SeqOf((_this).F32(), (_this).F32(), (_this).Fm5((_1549_v131).F48(), (_dafny.Zero).Minus((_1563_v137).Cardinality()), _this.F28(), _1541_v124, globalState))), (_this).F34(), !(func() _dafny.Map {
					var _coll42 = _dafny.NewMapBuilder()
					_ = _coll42
					for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(910), _dafny.IntOfInt64(734))); ; {
						_compr_42, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _1565_v138 _dafny.Int
						_1565_v138 = interface{}(_compr_42).(_dafny.Int)
						if ((_dafny.IntOfInt64(910)).Cmp(_1565_v138) <= 0) && ((_1565_v138).Cmp(_dafny.IntOfInt64(734)) < 0) {
							_coll42.Add(Companion_Default___.SafeModuloInt(_1565_v138, (_this).F34()), _dafny.IntOfInt64(341))
						}
					}
					return _coll42.ToMap()
				}()).Contains((_this).F30()), (_1549_v131).F47(), _this.F29())
				_1564_v139 = _nw263
				r1 = ((_1549_v131).F48()).Cmp((_this).F34()) >= 0
				_1542_v125 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1542_v125, _dafny.UnicodeSeqOfUtf8Bytes("mhye")), _1542_v125)
				var _1566_v140 _dafny.Array
				_ = _1566_v140
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_33
				var _nw264 _dafny.Array
				_ = _nw264
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw264 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1567_p0 bool) func(_dafny.Int) bool {
						return func(_1568_i11 _dafny.Int) bool {
							return _1567_p0
						}
					})(p0)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw264 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw264).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw264).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1566_v140 = _nw264
				var _1569_v141 _dafny.Map
				_ = _1569_v141
				_1569_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_1549_v131).F47())
				var _1570_v142 _dafny.MultiSet
				_ = _1570_v142
				_1570_v142 = _dafny.MultiSetOf(_this.F29())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1566_v140), 0))
				_ = _index198
				(_1566_v140).ArraySet1(((_1569_v141).Cardinality()).Cmp((_1570_v142).Cardinality()) <= 0, (_index198).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1566_v140), 0))
				_ = _index199
				(_1566_v140).ArraySet1((((_this).F30()).Cmp(_dafny.IntOfInt64(166)) == 0) == ((_1549_v131).F47()), (_index199).Int())
			}
			var _1571_v143 _dafny.Array
			_ = _1571_v143
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_34
			var _nw265 _dafny.Array
			_ = _nw265
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw265 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1572_v125 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1573_i12 _dafny.Int) _dafny.Sequence {
						return _1572_v125
					}
				})(_1542_v125)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw265 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw265).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw265).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1571_v143 = _nw265
			var _rhs198 _dafny.Int = (_this).F30()
			_ = _rhs198
			var _rhs199 _dafny.Int = _1541_v124
			_ = _rhs199
			var _rhs200 _dafny.Array = _1571_v143
			_ = _rhs200
			var _rhs201 _dafny.Int = (_this).F34()
			_ = _rhs201
			var _rhs202 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1549_v131).Fm3(_this.F29(), (_1549_v131).F47(), globalState), _1542_v125)
			_ = _rhs202
			var _lhs157 *GlobalState = globalState
			_ = _lhs157
			r3 = _rhs198
			_lhs157.F13 = _rhs199
			_1571_v143 = _rhs200
			r2 = _rhs201
			_1542_v125 = _rhs202
		} else {
			var _1574_v144 _dafny.Sequence
			_ = _1574_v144
			_1574_v144 = _dafny.SeqOf((_this).F34())
			var _1575_v145 _dafny.Int
			_ = _1575_v145
			var _out59 _dafny.Int
			_ = _out59
			_out59 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_1574_v144, _dafny.SeqOf((_this).F30())), _this.F29(), (_this).F34(), p0, globalState)
			_1575_v145 = _out59
			r2 = _1575_v145
			var _1576_v146 _dafny.MultiSet
			_ = _1576_v146
			_1576_v146 = _dafny.MultiSetOf((_this).F30(), _1575_v145, (_this).F30(), _1575_v145)
			var _1577_v147 D4
			_ = _1577_v147
			_1577_v147 = Companion_D4_.Create_DC10_(_1576_v146)
			var _1578_v148 _dafny.Sequence
			_ = _1578_v148
			_1578_v148 = _dafny.SeqOf(_1576_v146, (_1577_v147).Dtor_cf22(), (_1576_v146).Difference(_1576_v146))
			var _1579_v149 _dafny.Array
			_ = _1579_v149
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_35
			var _nw266 _dafny.Array
			_ = _nw266
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw266 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) bool = func(_1580_i13 _dafny.Int) bool {
					return !(_this.F31())
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw266 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw266).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw266).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1579_v149 = _nw266
			var _rhs203 _dafny.Array = _1579_v149
			_ = _rhs203
			var _rhs204 _dafny.Int = _dafny.IntOfInt64(354)
			_ = _rhs204
			var _rhs205 _dafny.Int = Companion_Default___.Fm32(!(p0) || (p0), globalState)
			_ = _rhs205
			var _rhs206 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1578_v148, _1578_v148), _1578_v148)
			_ = _rhs206
			var _lhs158 *GlobalState = globalState
			_ = _lhs158
			var _lhs159 *GlobalState = globalState
			_ = _lhs159
			_lhs158.F21 = _rhs203
			_lhs159.F13 = _rhs204
			_1575_v145 = _rhs205
			_1578_v148 = _rhs206
			_1575_v145 = _1575_v145
			r3 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.MultiSetOf(_this.F29())).Cardinality()), ((_this).F34()).Times((_this).F34()))
		}
		var _1581_v150 _dafny.Array
		_ = _1581_v150
		var _len0_36 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_36
		var _nw267 _dafny.Array
		_ = _nw267
		if _len0_36.Cmp(_dafny.Zero) == 0 {
			_nw267 = _dafny.NewArray(_len0_36)
		} else {
			var _init36 func(_dafny.Int) _dafny.CodePoint = func(_1582_i14 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			}
			_ = _init36
			var _element0_36 = _init36(_dafny.Zero)
			_ = _element0_36
			_nw267 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
			(_nw267).ArraySet1CodePoint(_element0_36, 0)
			var _nativeLen0_36 = (_len0_36).Int()
			_ = _nativeLen0_36
			for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
				(_nw267).ArraySet1CodePoint(_init36(_dafny.IntOf(_i0_36)), _i0_36)
			}
		}
		_1581_v150 = _nw267
		var _1583_v151 *C0
		_ = _1583_v151
		var _nw268 *C0 = New_C0_()
		_ = _nw268
		_nw268.Ctor__(_1581_v150)
		_1583_v151 = _nw268
		var _1584_v152 _dafny.Map
		_ = _1584_v152
		_1584_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F35())
		var _1585_v153 _dafny.Sequence
		_ = _1585_v153
		_1585_v153 = _dafny.SeqOf(_1584_v152, _1584_v152, _1584_v152, _1584_v152, _1584_v152)
		var _1586_v154 D1
		_ = _1586_v154
		_1586_v154 = Companion_D1_.Create_DC3_(_this.F29(), (_this).F30(), _1583_v151, ((_this).F30()).Plus(_dafny.IntOfUint32((_1585_v153).Cardinality())))
		var _source26 D1 = _1586_v154
		_ = _source26
		if _source26.Is_DC3() {
			var _1587___mcc_h28 bool = _source26.Get_().(D1_DC3).Cf3
			_ = _1587___mcc_h28
			var _1588___mcc_h29 _dafny.Int = _source26.Get_().(D1_DC3).Cf4
			_ = _1588___mcc_h29
			var _1589___mcc_h30 *C0 = _source26.Get_().(D1_DC3).Cf5
			_ = _1589___mcc_h30
			var _1590___mcc_h31 _dafny.Int = _source26.Get_().(D1_DC3).Cf6
			_ = _1590___mcc_h31
			var _1591_cf6 _dafny.Int = _1590___mcc_h31
			_ = _1591_cf6
			var _1592_cf5 *C0 = _1589___mcc_h30
			_ = _1592_cf5
			var _1593_cf4 _dafny.Int = _1588___mcc_h29
			_ = _1593_cf4
			var _1594_cf3 bool = _1587___mcc_h28
			_ = _1594_cf3
			var _1595_v155 _dafny.Array
			_ = _1595_v155
			var _nw269 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw269
			_1595_v155 = _nw269
			var _1596_v156 _dafny.Sequence
			_ = _1596_v156
			_1596_v156 = _dafny.UnicodeSeqOfUtf8Bytes("xiirlaq")
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))
			_ = _index200
			(_1595_v155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1596_v156, _1596_v156), (_index200).Int())
			var _1597_v157 _dafny.Sequence
			_ = _1597_v157
			_1597_v157 = _dafny.SeqOf((_this).F30())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))
			_ = _index201
			var _rhs207 bool = ((_this).F32()).Select((Companion_Default___.SafeIndex((_1597_v157).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1597_v157).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)
			_ = _rhs207
			var _rhs208 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg101 _dafny.Int) interface{} {
					return coer101(arg101)
				}
			}(func(_1598_i15 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			}))
			_ = _rhs208
			var _rhs209 bool = _this.F28()
			_ = _rhs209
			var _lhs160 _dafny.Array = _1595_v155
			_ = _lhs160
			var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))
			_ = _lhs161
			var _lhs162 *C6 = _this
			_ = _lhs162
			r0 = _rhs207
			(_lhs160).ArraySet1(_rhs208, (_lhs161).Int())
			_lhs162.F29_set_(_rhs209)
			var _1599_v158 _dafny.Set
			_ = _1599_v158
			_1599_v158 = _dafny.SetOf(((_this).F30()).Cmp(_1593_cf4) == 0, true)
			var _1600_v159 _dafny.Set
			_ = _1600_v159
			_1600_v159 = _dafny.SetOf(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F30()))))
			var _rhs210 _dafny.Int = ((_dafny.Zero).Minus(_1593_cf4)).Times(((_dafny.SetOf(_1597_v157)).Intersection(_1600_v159)).Cardinality())
			_ = _rhs210
			var _rhs211 bool = !(((_this).F30()).Cmp(((_this).F34()).Plus(_1591_cf6)) != 0)
			_ = _rhs211
			var _rhs212 _dafny.CodePoint = _this.F36()
			_ = _rhs212
			var _rhs213 _dafny.Set = (_1599_v158).Difference(_1599_v158)
			_ = _rhs213
			var _rhs214 bool = ((_1599_v158).Cardinality()).Cmp((_this).F34()) >= 0
			_ = _rhs214
			var _lhs163 *GlobalState = globalState
			_ = _lhs163
			var _lhs164 *GlobalState = globalState
			_ = _lhs164
			var _lhs165 *C6 = _this
			_ = _lhs165
			var _lhs166 *C6 = _this
			_ = _lhs166
			_lhs163.F13 = _rhs210
			_lhs164.F26 = _rhs211
			_lhs165.F36_set_(_rhs212)
			_1599_v158 = _rhs213
			_lhs166.F29_set_(_rhs214)
			var _1601_v160 *C2
			_ = _1601_v160
			var _nw270 *C2 = New_C2_()
			_ = _nw270
			_nw270.Ctor__(_1591_cf6, Companion_Default___.Fm32(!(p0), globalState), !(_this.F35()), ((_this).F32()).Select((Companion_Default___.SafeIndex(_1593_cf4, _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), _this.F31())
			_1601_v160 = _nw270
			var _1602_v161 D9
			_ = _1602_v161
			_1602_v161 = Companion_D9_.Create_DC26_(_1601_v160)
			var _1603_v162 D9
			_ = _1603_v162
			_1603_v162 = Companion_D9_.Create_DC26_((_1602_v161).Dtor_cf52())
			var _source27 D9 = _1602_v161
			_ = _source27
			if _source27.Is_DC27() {
				var _1604___mcc_h34 *C0 = _source27.Get_().(D9_DC27).Cf53
				_ = _1604___mcc_h34
				var _1605___mcc_h35 _dafny.Sequence = _source27.Get_().(D9_DC27).Cf54
				_ = _1605___mcc_h35
				var _1606_cf54 _dafny.Sequence = _1605___mcc_h35
				_ = _1606_cf54
				var _1607_cf53 *C0 = _1604___mcc_h34
				_ = _1607_cf53
				var _1608_v163 T3
				_ = _1608_v163
				var _nw271 *C4 = New_C4_()
				_ = _nw271
				_nw271.Ctor__((_1601_v160).F46(), _1594_cf3, (_this).F32(), (_this).F33(), (_1584_v152).Cardinality(), _this.F35(), _this.F35(), _this.F29())
				_1608_v163 = _nw271
				var _1609_v164 _dafny.MultiSet
				_ = _1609_v164
				_1609_v164 = _dafny.MultiSetOf(_1608_v163, _1608_v163, _1608_v163, _1608_v163, _1608_v163)
				var _1610_v165 D10
				_ = _1610_v165
				_1610_v165 = Companion_D10_.Create_DC30_(_this.F36(), (_1609_v164).Update(_1608_v163, Companion_Default___.Abs((_1601_v160).F46())))
				_1610_v165 = _1610_v165
				var _1611_v166 _dafny.Array
				_ = _1611_v166
				var _nw272 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw272
				_1611_v166 = _nw272
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_1611_v166), 0))
				_ = _index202
				(_1611_v166).ArraySet1(_1608_v163.F31(), (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_1611_v166), 0))
				_ = _index203
				(_1611_v166).ArraySet1(p0, (_index203).Int())
				var _1612_v167 _dafny.Map
				_ = _1612_v167
				_1612_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(556), _1608_v163.F28())
				_1612_v167 = _1612_v167
				var _1613_v168 _dafny.Map
				_ = _1613_v168
				_1613_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(_dafny.IntOfInt64(-700))).Cmp((_1597_v157).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1606_cf54, (Companion_Default___.SafeIndex((_1601_v160).F46(), _dafny.IntOfUint32((_1606_cf54).Cardinality()))).Uint32(), _this.F36())).Cardinality()), _dafny.IntOfUint32((_1597_v157).Cardinality()))).Uint32()).(_dafny.Int)) <= 0, (_dafny.MultiSetOf(!(_1608_v163.F35()), false, (_1611_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_1611_v166), 0))).Int()).(bool))).Intersection(_dafny.MultiSetOf(_1608_v163.F31())))
				var _1614_v169 _dafny.MultiSet
				_ = _1614_v169
				_1614_v169 = _dafny.MultiSetOf(false)
				var _1615_v170 _dafny.Map
				_ = _1615_v170
				_1615_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence), _1591_cf6)
				var _1616_v171 D9
				_ = _1616_v171
				_1616_v171 = Companion_D9_.Create_DC27_(_1583_v151, (_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence))
				var _1617_v172 _dafny.Sequence
				_ = _1617_v172
				_1617_v172 = _dafny.SeqOf(_1616_v171)
				var _rhs215 _dafny.Int = ((_1601_v160).F46()).Times(((_1601_v160).Fm29(globalState)).Minus(_dafny.IntOfInt64(564)))
				_ = _rhs215
				var _rhs216 bool = ((_1614_v169).Update(_1608_v163.F29(), Companion_Default___.Abs((_1615_v170).Cardinality()))).IsProperSubsetOf(_dafny.MultiSetOf(_1608_v163.F28(), _this.F28(), (_1611_v166).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_1611_v166), 0))).Int()).(bool)))
				_ = _rhs216
				var _rhs217 bool = _1594_cf3
				_ = _rhs217
				var _rhs218 _dafny.Map = ((_1613_v168).Merge(_1613_v168)).Update(!(_1608_v163.F28()) || (_1608_v163.F29()), _dafny.MultiSetFromSeq((_this).F32()))
				_ = _rhs218
				var _rhs219 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).Fm3(_dafny.Companion_Sequence_.IsPrefixOf(_1617_v172, _1617_v172), _1594_cf3, globalState), (Companion_Default___.SafeIndex((_1601_v160).F46(), _dafny.IntOfUint32(((_this).Fm3(_dafny.Companion_Sequence_.IsPrefixOf(_1617_v172, _1617_v172), _1594_cf3, globalState)).Cardinality()))).Uint32(), _this.F36())).Cardinality())
				_ = _rhs219
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				_1591_cf6 = _rhs215
				_lhs167.F26 = _rhs216
				_lhs168.F1 = _rhs217
				_1613_v168 = _rhs218
				_1591_cf6 = _rhs219
			} else {
				var _1618___mcc_h36 *C2 = _source27.Get_().(D9_DC26).Cf52
				_ = _1618___mcc_h36
				var _1619_cf52 *C2 = _1618___mcc_h36
				_ = _1619_cf52
				_1596_v156 = (_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence)
				(globalState).F12 = true
				var _1620_v173 _dafny.MultiSet
				_ = _1620_v173
				_1620_v173 = _dafny.MultiSetOf(_1596_v156)
				var _1621_v174 _dafny.Map
				_ = _1621_v174
				_1621_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1620_v173).Union(_dafny.MultiSetOf((_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence), (_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence), (_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence), (_1595_v155).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1595_v155), 0))).Int()).(_dafny.Sequence))), (_1601_v160).F46())
				_1621_v174 = (_1621_v174).Update(_1620_v173, _1591_cf6)
				var _1622_v175 _dafny.MultiSet
				_ = _1622_v175
				_1622_v175 = _dafny.MultiSetOf(_this.F28(), _this.F35())
				_1622_v175 = (_1622_v175).Difference(_1622_v175)
			}
			var _1623_v176 _dafny.Sequence
			_ = _1623_v176
			_1623_v176 = _dafny.SeqOf(_1596_v156)
			_1623_v176 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-677))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}((func(_1624_v156 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1625_i16 _dafny.Int) _dafny.Sequence {
					return _1624_v156
				}
			})(_1596_v156))), _dafny.SeqOf(_1596_v156))
		} else if _source26.Is_DC2() {
			var _1626___mcc_h32 _dafny.Array = _source26.Get_().(D1_DC2).Cf2
			_ = _1626___mcc_h32
			var _1627_cf2 _dafny.Array = _1626___mcc_h32
			_ = _1627_cf2
			var _1628_v177 _dafny.Set
			_ = _1628_v177
			_1628_v177 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}(func(_1629_i17 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			}))).Cardinality())), (_this).F34(), (_this).F30(), Companion_Default___.Fm32(!(_this.F35()), globalState))
			_1628_v177 = (_1628_v177).Intersection(_1628_v177)
			(globalState).F1 = _this.F35()
			(globalState).F13 = _dafny.IntOfInt64(804)
			var _1630_v178 _dafny.Map
			_ = _1630_v178
			_1630_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (_this).F30())
			_1630_v178 = _1630_v178
		} else {
			var _1631___mcc_h33 D1 = _source26.Get_().(D1_DC4).Cf7
			_ = _1631___mcc_h33
			var _1632_cf7 D1 = _1631___mcc_h33
			_ = _1632_cf7
			(_this).F31_set_(_this.F29())
			var _1633_v179 _dafny.MultiSet
			_ = _1633_v179
			_1633_v179 = _dafny.MultiSetOf(true)
			var _1634_v180 D0
			_ = _1634_v180
			_1634_v180 = Companion_D0_.Create_DC1_(_1633_v179)
			var _1635_v181 _dafny.Map
			_ = _1635_v181
			_1635_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1634_v180, _this.F29())
			var _1636_v182 D8
			_ = _1636_v182
			_1636_v182 = Companion_D8_.Create_DC21_((_1635_v181).Merge(_1635_v181))
			var _source28 D8 = _1636_v182
			_ = _source28
			if _source28.Is_DC22() {
				var _1637___mcc_h37 _dafny.Int = _source28.Get_().(D8_DC22).Cf41
				_ = _1637___mcc_h37
				var _1638___mcc_h38 _dafny.Int = _source28.Get_().(D8_DC22).Cf42
				_ = _1638___mcc_h38
				var _1639_cf42 _dafny.Int = _1638___mcc_h38
				_ = _1639_cf42
				var _1640_cf41 _dafny.Int = _1637___mcc_h37
				_ = _1640_cf41
				var _1641_v183 _dafny.Array
				_ = _1641_v183
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_37
				var _nw273 _dafny.Array
				_ = _nw273
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw273 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.Int = (func(_1642_cf41 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1643_i18 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1643_i18, _1642_cf41)
						}
					})(_1640_cf41)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw273 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw273).ArraySet1(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw273).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1641_v183 = _nw273
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1641_v183), 0))
				_ = _index204
				(_1641_v183).ArraySet1(_dafny.IntOfInt64(593), (_index204).Int())
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1641_v183), 0))
				_ = _index205
				(_1641_v183).ArraySet1((_1639_cf42).Plus((_1639_cf42).Times((_this).F34())), (_index205).Int())
				var _1644_v184 _dafny.Array
				_ = _1644_v184
				var _nw274 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw274
				_1644_v184 = _nw274
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1644_v184), 0))
				_ = _index206
				(_1644_v184).ArraySet1(p0, (_index206).Int())
				var _1645_v185 _dafny.Set
				_ = _1645_v185
				_1645_v185 = _dafny.SetOf(_this.F28())
				var _1646_v186 D7
				_ = _1646_v186
				_1646_v186 = Companion_D7_.Create_DC19_(_this.F31(), _dafny.IntOfInt64(530), _dafny.IntOfInt64(368))
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1644_v184), 0))
				_ = _index207
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1641_v183), 0))
				_ = _index208
				var _rhs220 bool = (_1645_v185).IsProperSubsetOf((func() _dafny.Set {
					if p0 {
						return _1645_v185
					}
					return _1645_v185
				})())
				_ = _rhs220
				var _rhs221 bool = _this.F31()
				_ = _rhs221
				var _rhs222 _dafny.Int = (_dafny.Zero).Minus((_this).F34())
				_ = _rhs222
				var _rhs223 bool = !((_1646_v186).Dtor_cf36())
				_ = _rhs223
				var _lhs169 _dafny.Array = _1644_v184
				_ = _lhs169
				var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1644_v184), 0))
				_ = _lhs170
				var _lhs171 *C6 = _this
				_ = _lhs171
				var _lhs172 _dafny.Array = _1641_v183
				_ = _lhs172
				var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_1641_v183), 0))
				_ = _lhs173
				var _lhs174 *GlobalState = globalState
				_ = _lhs174
				(_lhs169).ArraySet1(_rhs220, (_lhs170).Int())
				_lhs171.F28_set_(_rhs221)
				(_lhs172).ArraySet1(_rhs222, (_lhs173).Int())
				_lhs174.F26 = _rhs223
				var _1647_v187 _dafny.Sequence
				_ = _1647_v187
				_1647_v187 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F34()))
				var _1648_v188 _dafny.Map
				_ = _1648_v188
				_1648_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1644_v184).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1644_v184), 0))).Int()).(bool), _1640_cf41)
				(globalState).F13 = (((_1647_v187).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-962), _dafny.IntOfUint32((_1647_v187).Cardinality()))).Uint32()).(_dafny.Map)).Merge((_1648_v188).Merge(_1648_v188))).Cardinality()
				var _1649_v189 _dafny.Sequence
				_ = _1649_v189
				_1649_v189 = _dafny.UnicodeSeqOfUtf8Bytes("mwd")
				var _rhs224 _dafny.Int = _dafny.IntOfInt64(742)
				_ = _rhs224
				var _rhs225 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uy"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(863))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg104 _dafny.Int) interface{} {
						return coer104(arg104)
					}
				}(func(_1650_i19 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}))), _dafny.UnicodeSeqOfUtf8Bytes("gvqi"))
				_ = _rhs225
				var _rhs226 bool = (_1644_v184).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1644_v184), 0))).Int()).(bool)
				_ = _rhs226
				var _lhs175 *GlobalState = globalState
				_ = _lhs175
				r3 = _rhs224
				_1649_v189 = _rhs225
				_lhs175.F1 = _rhs226
			} else if _source28.Is_DC23() {
				var _1651___mcc_h39 bool = _source28.Get_().(D8_DC23).Cf43
				_ = _1651___mcc_h39
				var _1652___mcc_h40 bool = _source28.Get_().(D8_DC23).Cf44
				_ = _1652___mcc_h40
				var _1653___mcc_h41 bool = _source28.Get_().(D8_DC23).Cf45
				_ = _1653___mcc_h41
				var _1654_cf45 bool = _1653___mcc_h41
				_ = _1654_cf45
				var _1655_cf44 bool = _1652___mcc_h40
				_ = _1655_cf44
				var _1656_cf43 bool = _1651___mcc_h39
				_ = _1656_cf43
				var _1657_v190 _dafny.Sequence
				_ = _1657_v190
				_1657_v190 = _dafny.SeqOf((_this).F34())
				var _1658_v191 *C4
				_ = _1658_v191
				var _nw275 *C4 = New_C4_()
				_ = _nw275
				_nw275.Ctor__((_1633_v179).Cardinality(), _this.F28(), _dafny.Companion_Sequence_.Concatenate((_this).F32(), ((_this).F33()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg105 _dafny.Int) interface{} {
						return coer105(arg105)
					}
				}(func(_1659_i20 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}))).Cardinality()), _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence)), (_this).F33(), (_this).F30(), false, _1655_cf44, _dafny.Companion_Sequence_.IsProperPrefixOf(_1657_v190, _dafny.Companion_Sequence_.Update(_1657_v190, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.IntOfUint32((_1657_v190).Cardinality()))).Uint32(), (_this).F30())))
				_1658_v191 = _nw275
				var _1660_v192 *C1
				_ = _1660_v192
				var _nw276 *C1 = New_C1_()
				_ = _nw276
				_nw276.Ctor__(false, _this.F31(), _1656_cf43)
				_1660_v192 = _nw276
				_1660_v192 = _1660_v192
				var _1661_v193 _dafny.Map
				_ = _1661_v193
				_1661_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _1655_cf44)
				var _1662_v194 _dafny.Map
				_ = _1662_v194
				_1662_v194 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1661_v193, _1656_cf43)
				var _1663_v195 _dafny.Map
				_ = _1663_v195
				_1663_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1662_v194, _this.F28())
				var _1664_v196 _dafny.Sequence
				_ = _1664_v196
				_1664_v196 = _dafny.UnicodeSeqOfUtf8Bytes("nvyflppiw")
				var _1665_v197 _dafny.MultiSet
				_ = _1665_v197
				_1665_v197 = _dafny.MultiSetOf((_this).F30())
				var _1666_v198 _dafny.Map
				_ = _1666_v198
				_1666_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F35())
				var _1667_v199 _dafny.Array
				_ = _1667_v199
				var _nwElement0_56 bool = _this.F31()
				_ = _nwElement0_56
				var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(28))
				_ = _nw277
				(_nw277).ArraySet1(_nwElement0_56, 0)
				(_nw277).ArraySet1((func() bool {
					if (_1663_v195).Contains(_1662_v194) {
						return (_1663_v195).Get(_1662_v194).(bool)
					}
					return _this.F29()
				})(), 1)
				(_nw277).ArraySet1(_this.F29(), 2)
				(_nw277).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1664_v196, _1664_v196), 3)
				(_nw277).ArraySet1((_1660_v192).F43(), 4)
				(_nw277).ArraySet1(_1654_cf45, 5)
				(_nw277).ArraySet1(!((_1660_v192).F43()) || (p0), 6)
				(_nw277).ArraySet1(_1655_cf44, 7)
				(_nw277).ArraySet1(_1656_cf43, 8)
				(_nw277).ArraySet1(_this.F35(), 9)
				(_nw277).ArraySet1(_this.F28(), 10)
				(_nw277).ArraySet1((_1660_v192).F43(), 11)
				(_nw277).ArraySet1(_this.F31(), 12)
				(_nw277).ArraySet1((_1660_v192).F43(), 13)
				(_nw277).ArraySet1(_dafny.Companion_Sequence_.Contains((_this).F32(), _1654_cf45), 14)
				(_nw277).ArraySet1((_1656_cf43) || ((func() bool {
					if (_1584_v152).Contains(_this.F28()) {
						return (_1584_v152).Get(_this.F28()).(bool)
					}
					return (_this).Fm7(_1665_v197, (_this).F34(), globalState)
				})()), 15)
				(_nw277).ArraySet1((_1660_v192).F43(), 16)
				(_nw277).ArraySet1(_1656_cf43, 17)
				(_nw277).ArraySet1((_1660_v192).F43(), 18)
				(_nw277).ArraySet1(true, 19)
				(_nw277).ArraySet1((_1660_v192).F43(), 20)
				(_nw277).ArraySet1(_this.F29(), 21)
				(_nw277).ArraySet1(_1656_cf43, 22)
				(_nw277).ArraySet1(!((_1660_v192).F43()), 23)
				(_nw277).ArraySet1(_this.F29(), 24)
				(_nw277).ArraySet1((func() bool {
					if (_1666_v198).Contains((_this).F34()) {
						return (_1666_v198).Get((_this).F34()).(bool)
					}
					return _this.F28()
				})(), 25)
				(_nw277).ArraySet1(_this.F35(), 26)
				(_nw277).ArraySet1(true, 27)
				_1667_v199 = _nw277
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1667_v199), 0))
				_ = _index209
				(_1667_v199).ArraySet1(_this.F28(), (_index209).Int())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1667_v199), 0))
				_ = _index210
				var _rhs227 _dafny.Int = (func() _dafny.Int {
					if (_1665_v197).Contains((_this).F34()) {
						return (_1665_v197).Multiplicity((_this).F34())
					}
					return (_dafny.Zero).Minus((_this).F34())
				})()
				_ = _rhs227
				var _rhs228 bool = _1656_cf43
				_ = _rhs228
				var _lhs176 *GlobalState = globalState
				_ = _lhs176
				var _lhs177 _dafny.Array = _1667_v199
				_ = _lhs177
				var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1667_v199), 0))
				_ = _lhs178
				_lhs176.F13 = _rhs227
				(_lhs177).ArraySet1(_rhs228, (_lhs178).Int())
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1667_v199), 0))
				_ = _index211
				(_1667_v199).ArraySet1((_1667_v199).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_1667_v199), 0))).Int()).(bool), (_index211).Int())
			} else if _source28.Is_DC24() {
				var _1668___mcc_h42 bool = _source28.Get_().(D8_DC24).Cf46
				_ = _1668___mcc_h42
				var _1669___mcc_h43 _dafny.Int = _source28.Get_().(D8_DC24).Cf47
				_ = _1669___mcc_h43
				var _1670___mcc_h44 _dafny.Sequence = _source28.Get_().(D8_DC24).Cf48
				_ = _1670___mcc_h44
				var _1671___mcc_h45 bool = _source28.Get_().(D8_DC24).Cf49
				_ = _1671___mcc_h45
				var _1672___mcc_h46 _dafny.Sequence = _source28.Get_().(D8_DC24).Cf50
				_ = _1672___mcc_h46
				var _1673_cf50 _dafny.Sequence = _1672___mcc_h46
				_ = _1673_cf50
				var _1674_cf49 bool = _1671___mcc_h45
				_ = _1674_cf49
				var _1675_cf48 _dafny.Sequence = _1670___mcc_h44
				_ = _1675_cf48
				var _1676_cf47 _dafny.Int = _1669___mcc_h43
				_ = _1676_cf47
				var _1677_cf46 bool = _1668___mcc_h42
				_ = _1677_cf46
				var _1678_v200 _dafny.Array
				_ = _1678_v200
				var _nwElement0_57 bool = _1674_cf49
				_ = _nwElement0_57
				var _nw278 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(2))
				_ = _nw278
				(_nw278).ArraySet1(_nwElement0_57, 0)
				(_nw278).ArraySet1(((_this).F30()).Cmp(_1676_cf47) <= 0, 1)
				_1678_v200 = _nw278
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1678_v200), 0))
				_ = _index212
				(_1678_v200).ArraySet1(_this.F28(), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1678_v200), 0))
				_ = _index213
				(_1678_v200).ArraySet1(!(_this.F35()), (_index213).Int())
				var _1679_v201 _dafny.Array
				_ = _1679_v201
				var _len0_38 _dafny.Int = _dafny.One
				_ = _len0_38
				var _nw279 _dafny.Array
				_ = _nw279
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw279 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Int = func(_1680_i21 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1680_i21, (_this).F30())
					}
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw279 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw279).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw279).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1679_v201 = _nw279
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1678_v200), 0))
				_ = _index214
				var _rhs229 bool = _1674_cf49
				_ = _rhs229
				var _rhs230 _dafny.Int = (Companion_Default___.SafeModuloInt(_1676_cf47, _dafny.IntOfInt64(22))).Minus(((_1633_v179).Cardinality()).Plus(_1676_cf47))
				_ = _rhs230
				var _rhs231 _dafny.Array = _1679_v201
				_ = _rhs231
				var _rhs232 bool = (_dafny.IntOfInt64(645)).Cmp(_1676_cf47) <= 0
				_ = _rhs232
				var _lhs179 _dafny.Array = _1678_v200
				_ = _lhs179
				var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1678_v200), 0))
				_ = _lhs180
				_1677_cf46 = _rhs229
				r3 = _rhs230
				_1679_v201 = _rhs231
				(_lhs179).ArraySet1(_rhs232, (_lhs180).Int())
				var _1681_v202 T3
				_ = _1681_v202
				var _nw280 *C4 = New_C4_()
				_ = _nw280
				_nw280.Ctor__((_this).F30(), _this.F35(), _dafny.SeqOf(_this.F31(), p0, _this.F35(), _1674_cf49, _1677_cf46), (_this).F33(), (_this).F34(), _this.F35(), _this.F31(), (_1678_v200).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_1678_v200), 0))).Int()).(bool))
				_1681_v202 = _nw280
				var _1682_v203 _dafny.MultiSet
				_ = _1682_v203
				_1682_v203 = _dafny.MultiSetOf(_1681_v202, _1681_v202)
				var _1683_v204 _dafny.Map
				_ = _1683_v204
				_1683_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC30_(_this.F36(), _1682_v203), ((_1681_v202).F34()).Minus((_dafny.Zero).Minus(_1676_cf47)))
				_1683_v204 = (_1683_v204).Update(Companion_D10_.Create_DC30_(_this.F36(), _dafny.MultiSetOf(_1681_v202)), _1676_cf47)
				var _1684_v205 _dafny.Sequence
				_ = _1684_v205
				_1684_v205 = _dafny.SeqOf(_1583_v151.F37, _1583_v151.F37, _1583_v151.F37)
				_1673_cf50 = _dafny.SeqOf(_dafny.IntOfUint32((_1684_v205).Cardinality()))
			} else if _source28.Is_DC21() {
				var _1685___mcc_h47 _dafny.Map = _source28.Get_().(D8_DC21).Cf40
				_ = _1685___mcc_h47
				var _1686_cf40 _dafny.Map = _1685___mcc_h47
				_ = _1686_cf40
				var _1687_v206 _dafny.Sequence
				_ = _1687_v206
				_1687_v206 = _dafny.UnicodeSeqOfUtf8Bytes("jkeeiw")
				(globalState).F12 = _dafny.Companion_Sequence_.Contains(_1687_v206, _this.F36())
				var _1688_v207 _dafny.MultiSet
				_ = _1688_v207
				_1688_v207 = _dafny.MultiSetOf((_this).F30())
				var _1689_v208 _dafny.Set
				_ = _1689_v208
				_1689_v208 = _dafny.SetOf((_1584_v152).Cardinality(), (_this).F34(), (_this).F30())
				var _1690_v209 _dafny.Sequence
				_ = _1690_v209
				_1690_v209 = _dafny.SeqOf((_this).F34(), (_1633_v179).Cardinality(), (_this).F30(), (_1633_v179).Cardinality(), (_this).F30())
				var _1691_v210 _dafny.Set
				_ = _1691_v210
				_1691_v210 = _dafny.SetOf(!(!(_this.F28())), _this.F28(), false)
				var _1692_v211 _dafny.Array
				_ = _1692_v211
				var _nwElement0_58 bool = (_1688_v207).Contains((_this).F34())
				_ = _nwElement0_58
				var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(8))
				_ = _nw281
				(_nw281).ArraySet1(_nwElement0_58, 0)
				(_nw281).ArraySet1(_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm42(_1687_v206, (func() _dafny.Int {
					if (_1688_v207).Contains((_1689_v208).Cardinality()) {
						return (_1688_v207).Multiplicity((_1689_v208).Cardinality())
					}
					return (_this).F30()
				})(), _dafny.UnicodeSeqOfUtf8Bytes("yvyovs"), globalState), _1690_v209), 1)
				(_nw281).ArraySet1(_this.F31(), 2)
				(_nw281).ArraySet1((_1691_v210).IsDisjointFrom(_1691_v210), 3)
				(_nw281).ArraySet1(_this.F31(), 4)
				(_nw281).ArraySet1(_this.F29(), 5)
				(_nw281).ArraySet1(p0, 6)
				(_nw281).ArraySet1(_this.F28(), 7)
				_1692_v211 = _nw281
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1692_v211), 0))
				_ = _index215
				(_1692_v211).ArraySet1(((_this).F30()).Cmp((Companion_Default___.Fm53(_dafny.CodePoint('v'), (_this).F30(), globalState)).Dtor_cf62()) > 0, (_index215).Int())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((_1692_v211), 0))
				_ = _index216
				(_1692_v211).ArraySet1(Companion_Default___.Fm1(p0, (_this).F30(), globalState), (_index216).Int())
				var _1693_v212 _dafny.Array
				_ = _1693_v212
				var _nw282 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw282
				_1693_v212 = _nw282
				var _1694_v213 _dafny.Array
				_ = _1694_v213
				var _nw283 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw283
				_1694_v213 = _nw283
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1693_v212), 0))
				_ = _index217
				(_1693_v212).ArraySet1(_1694_v213, (_index217).Int())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1693_v212), 0))
				_ = _index218
				var _rhs233 _dafny.Array = _1694_v213
				_ = _rhs233
				var _rhs234 _dafny.Int = (_this).F30()
				_ = _rhs234
				var _rhs235 _dafny.Int = (_this).F34()
				_ = _rhs235
				var _rhs236 bool = (_this).Fm6(globalState)
				_ = _rhs236
				var _lhs181 _dafny.Array = _1693_v212
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_1693_v212), 0))
				_ = _lhs182
				var _lhs183 *GlobalState = globalState
				_ = _lhs183
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				var _lhs185 *C6 = _this
				_ = _lhs185
				(_lhs181).ArraySet1(_rhs233, (_lhs182).Int())
				_lhs183.F13 = _rhs234
				_lhs184.F13 = _rhs235
				_lhs185.F35_set_(_rhs236)
				var _1695_v214 _dafny.Map
				_ = _1695_v214
				_1695_v214 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F30())
				_1695_v214 = (_1695_v214).Update(_this.F35(), (_this).F34())
			} else {
				var _1696___mcc_h48 D8 = _source28.Get_().(D8_DC25).Cf51
				_ = _1696___mcc_h48
				var _1697_cf51 D8 = _1696___mcc_h48
				_ = _1697_cf51
				(globalState).F13 = (_this).F34()
				(globalState).F1 = p0
				(globalState).F13 = Companion_Default___.SafeModuloInt(((_this).F34()).Minus(_dafny.IntOfInt64(669)), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-218))))
				var _1698_v215 _dafny.Map
				_ = _1698_v215
				_1698_v215 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), p0)
				var _rhs237 bool = !(p0)
				_ = _rhs237
				var _rhs238 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm46((_dafny.Zero).Minus((_this).F34()), (_this).F34(), (func() bool {
					if (_1698_v215).Contains((_this).F30()) {
						return (_1698_v215).Get((_this).F30()).(bool)
					}
					return p0
				})(), globalState))
				_ = _rhs238
				var _lhs186 *C6 = _this
				_ = _lhs186
				_lhs186.F31_set_(_rhs237)
				r3 = _rhs238
			}
			var _1699_v216 _dafny.Sequence
			_ = _1699_v216
			_1699_v216 = _dafny.UnicodeSeqOfUtf8Bytes("ylj")
			_1699_v216 = _1699_v216
			var _rhs239 _dafny.Int = (_this).F34()
			_ = _rhs239
			var _rhs240 _dafny.Int = (_this).F30()
			_ = _rhs240
			var _lhs187 *GlobalState = globalState
			_ = _lhs187
			r2 = _rhs239
			_lhs187.F13 = _rhs240
		}
		var _1700_v217 _dafny.Array
		_ = _1700_v217
		var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw284
		_1700_v217 = _nw284
		var _1701_v218 _dafny.Array
		_ = _1701_v218
		var _nw285 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
		_ = _nw285
		_1701_v218 = _nw285
		var _1702_v219 _dafny.Sequence
		_ = _1702_v219
		_1702_v219 = _dafny.SeqOf(_1701_v218, _1701_v218, _1701_v218, _1701_v218)
		var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))
		_ = _index219
		(_1700_v217).ArraySet1(_dafny.IntOfUint32((_1702_v219).Cardinality()), (_index219).Int())
		var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))
		_ = _index220
		(_1700_v217).ArraySet1((_this).F30(), (_index220).Int())
		r2 = (_1700_v217).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))).Int()).(_dafny.Int)
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1700_v217), 0))); ; {
			_guard_loop_2, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _1703_i22 _dafny.Int
			_1703_i22 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1703_i22).Sign() != -1) && ((_1703_i22).Cmp(_dafny.ArrayLen((_1700_v217), 0)) < 0)) {
				(_1700_v217).ArraySet1((_1703_i22).Times((_this).F30()), (_1703_i22).Int())
			}
		}
		r0 = true
		r1 = _this.F31()
		var _1704_v220 _dafny.Set
		_ = _1704_v220
		_1704_v220 = _dafny.SetOf((_1700_v217).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))).Int()).(_dafny.Int))
		var _1705_v221 _dafny.Map
		_ = _1705_v221
		_1705_v221 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1704_v220)
		var _1706_v222 _dafny.Map
		_ = _1706_v222
		_1706_v222 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1705_v221).Merge(_1705_v221))
		r2 = (_dafny.Zero).Minus(((func() _dafny.Map {
			if (_1706_v222).Contains(_this.F31()) {
				return (_1706_v222).Get(_this.F31()).(_dafny.Map)
			}
			return (_1705_v221).Merge(_1705_v221)
		})()).Cardinality())
		var _1707_v223 _dafny.Sequence
		_ = _1707_v223
		_1707_v223 = _dafny.UnicodeSeqOfUtf8Bytes("urgiph")
		var _1708_v224 _dafny.Sequence
		_ = _1708_v224
		_1708_v224 = _dafny.SeqOf((_this).F34(), _dafny.IntOfUint32((_1707_v223).Cardinality()), ((_1700_v217).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))).Int()).(_dafny.Int)).Times((_1700_v217).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))).Int()).(_dafny.Int)))
		r3 = (_1708_v224).Select((Companion_Default___.SafeIndex((_1700_v217).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_1700_v217), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1708_v224).Cardinality()))).Uint32()).(_dafny.Int)
		return r0, r1, r2, r3
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f30 _dafny.Int
	_f41 _dafny.Int
	_f42 _dafny.MultiSet
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f30 = _dafny.Zero
	_this._f41 = _dafny.Zero
	_this._f42 = _dafny.EmptyMultiSet
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

func (_this *C7) F28() bool {
	return _this._f28
}
func (_this *C7) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C7) F29() bool {
	return _this._f29
}
func (_this *C7) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C7) F31() bool {
	return _this._f31
}
func (_this *C7) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C7) F30() _dafny.Int {
	return _this._f30
}
func (_this *C7) Ctor__(f41 _dafny.Int, f42 _dafny.MultiSet, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f41 = f41
		(_this)._f42 = f42
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C7) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(557))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg106 _dafny.Int) interface{} {
				return coer106(arg106)
			}
		}(func(_1709_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), _dafny.UnicodeSeqOfUtf8Bytes("lytqd")), (func() _dafny.Sequence {
			if _this.F28() {
				return _dafny.UnicodeSeqOfUtf8Bytes("gig")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("cbkwbpt")
		})())
	}
}
func (_this *C7) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F41())
	}
}
func (_this *C7) Fm19(globalState *GlobalState) bool {
	{
		return _this.F29()
	}
}
func (_this *C7) Fm20(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F30()))
	}
}
func (_this *C7) M7(p0 bool, p1 _dafny.Array, p2 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Sequence, _dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		(_this).F29_set_(!(_this.F31()) || (p2))
		(_this).F31_set_(p2)
		var _1710_v0 _dafny.Map
		_ = _1710_v0
		_1710_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false), (_this).F41())
		var _hi8 _dafny.Int = (_this).F41()
		_ = _hi8
		for _1711_i0 := (_1710_v0).Cardinality(); _1711_i0.Cmp(_hi8) < 0; _1711_i0 = _1711_i0.Plus(_dafny.One) {
			var _1712_v1 _dafny.Sequence
			_ = _1712_v1
			_1712_v1 = _dafny.UnicodeSeqOfUtf8Bytes("fiw")
			r3 = _dafny.Companion_Sequence_.Concatenate((_this).Fm3(p0, _this.F31(), globalState), _dafny.Companion_Sequence_.Concatenate(_1712_v1, _dafny.UnicodeSeqOfUtf8Bytes("phvd")))
			(globalState).F13 = ((_this).F30()).Minus(_1711_i0)
			r2 = _dafny.IntOfInt64(445)
			var _1713_v2 _dafny.MultiSet
			_ = _1713_v2
			_1713_v2 = _dafny.MultiSetOf((_this).Fm20((_this).F41(), false, globalState), (_this).F30(), _1711_i0)
			_1713_v2 = (_this).F42()
		}
		var _1714_v3 _dafny.Sequence
		_ = _1714_v3
		_1714_v3 = _dafny.SeqOf((_this).F41())
		var _1715_v4 _dafny.Map
		_ = _1715_v4
		_1715_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), _dafny.IntOfInt64(-371))
		(globalState).F13 = (_1714_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_1715_v4).Contains(_this.F29()) {
				return (_1715_v4).Get(_this.F29()).(_dafny.Int)
			}
			return (_this).F41()
		})(), _dafny.IntOfUint32((_1714_v3).Cardinality()))).Uint32()).(_dafny.Int)
		var _1716_v5 _dafny.CodePoint
		_ = _1716_v5
		_1716_v5 = _dafny.CodePoint('r')
		var _1717_v6 _dafny.Sequence
		_ = _1717_v6
		_1717_v6 = _dafny.UnicodeSeqOfUtf8Bytes("pdeidmrnl")
		var _1718_v7 _dafny.Array
		_ = _1718_v7
		var _nwElement0_59 _dafny.CodePoint = _1716_v5
		_ = _nwElement0_59
		var _nw286 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(28))
		_ = _nw286
		(_nw286).ArraySet1CodePoint(_nwElement0_59, 0)
		(_nw286).ArraySet1CodePoint(_dafny.CodePoint('q'), 1)
		(_nw286).ArraySet1CodePoint(_1716_v5, 2)
		(_nw286).ArraySet1CodePoint(_1716_v5, 3)
		(_nw286).ArraySet1CodePoint(_1716_v5, 4)
		(_nw286).ArraySet1CodePoint(_1716_v5, 5)
		(_nw286).ArraySet1CodePoint(_1716_v5, 6)
		(_nw286).ArraySet1CodePoint(_1716_v5, 7)
		(_nw286).ArraySet1CodePoint(_1716_v5, 8)
		(_nw286).ArraySet1CodePoint(_1716_v5, 9)
		(_nw286).ArraySet1CodePoint(_dafny.CodePoint('e'), 10)
		(_nw286).ArraySet1CodePoint(Companion_Default___.Fm21(_1716_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F41(), false)).Cardinality(), globalState), 11)
		(_nw286).ArraySet1CodePoint(_1716_v5, 12)
		(_nw286).ArraySet1CodePoint(_1716_v5, 13)
		(_nw286).ArraySet1CodePoint(_1716_v5, 14)
		(_nw286).ArraySet1CodePoint(_1716_v5, 15)
		(_nw286).ArraySet1CodePoint(_1716_v5, 16)
		(_nw286).ArraySet1CodePoint(_1716_v5, 17)
		(_nw286).ArraySet1CodePoint(_1716_v5, 18)
		(_nw286).ArraySet1CodePoint(_1716_v5, 19)
		(_nw286).ArraySet1CodePoint(Companion_Default___.Fm21(Companion_Default___.Fm21(_1716_v5, (_this).F30(), globalState), (_this).F30(), globalState), 20)
		(_nw286).ArraySet1CodePoint((_1717_v6).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1717_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), 21)
		(_nw286).ArraySet1CodePoint(_dafny.CodePoint('e'), 22)
		(_nw286).ArraySet1CodePoint(_1716_v5, 23)
		(_nw286).ArraySet1CodePoint(_1716_v5, 24)
		(_nw286).ArraySet1CodePoint(_1716_v5, 25)
		(_nw286).ArraySet1CodePoint(_1716_v5, 26)
		(_nw286).ArraySet1CodePoint(_1716_v5, 27)
		_1718_v7 = _nw286
		var _1719_v8 *C0
		_ = _1719_v8
		var _nw287 *C0 = New_C0_()
		_ = _nw287
		_nw287.Ctor__(_1718_v7)
		_1719_v8 = _nw287
		var _1720_v9 _dafny.Map
		_ = _1720_v9
		_1720_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F41(), (_this).F30())
		var _1721_v10 _dafny.Array
		_ = _1721_v10
		var _nwElement0_60 _dafny.Int = (_this).F41()
		_ = _nwElement0_60
		var _nw288 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(20))
		_ = _nw288
		(_nw288).ArraySet1(_nwElement0_60, 0)
		(_nw288).ArraySet1((func() _dafny.Int {
			if p0 {
				return (_dafny.Zero).Minus((_this).F41())
			}
			return _dafny.IntOfInt64(277)
		})(), 1)
		(_nw288).ArraySet1((_this).F41(), 2)
		(_nw288).ArraySet1((_dafny.Zero).Minus((_this).F41()), 3)
		(_nw288).ArraySet1((_dafny.Zero).Minus((_this).F41()), 4)
		(_nw288).ArraySet1((_this).Fm20((_this).F41(), !(true), globalState), 5)
		(_nw288).ArraySet1((_this).F30(), 6)
		(_nw288).ArraySet1((_this).F30(), 7)
		(_nw288).ArraySet1((_this).F30(), 8)
		(_nw288).ArraySet1((_this).Fm20((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F30())), p0, globalState), 9)
		(_nw288).ArraySet1(((_dafny.Zero).Minus((_1720_v9).Cardinality())).Times((_this).F41()), 10)
		(_nw288).ArraySet1((_this).Fm20((_this).F41(), _this.F31(), globalState), 11)
		(_nw288).ArraySet1((_this).F30(), 12)
		(_nw288).ArraySet1(((_this).F41()).Plus((_this).Fm20((_this).F41(), p0, globalState)), 13)
		(_nw288).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if (_1715_v4).Contains(_this.F28()) {
				return (_1715_v4).Get(_this.F28()).(_dafny.Int)
			}
			return (_this).F30()
		})()), 14)
		(_nw288).ArraySet1((_this).F30(), 15)
		(_nw288).ArraySet1((_dafny.SetOf((_this).F41())).Cardinality(), 16)
		(_nw288).ArraySet1((_dafny.Zero).Minus((_this).F41()), 17)
		(_nw288).ArraySet1((_this).F30(), 18)
		(_nw288).ArraySet1(Companion_Default___.SafeModuloInt((_this).Fm20((_this).F30(), (_this).Fm19(globalState), globalState), (_this).F41()), 19)
		_1721_v10 = _nw288
		var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_1721_v10), 0))
		_ = _index221
		(_1721_v10).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F30(), (_dafny.Zero).Minus((_this).F30())), (_index221).Int())
		var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p1), 0))
		_ = _index222
		(p1).ArraySet1(p0, (_index222).Int())
		var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1721_v10), 0))
		_ = _index223
		(_1721_v10).ArraySet1((_this).F41(), (_index223).Int())
		var _1722_v11 _dafny.Map
		_ = _1722_v11
		_1722_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F41(), p2)
		var _1723_v12 _dafny.MultiSet
		_ = _1723_v12
		_1723_v12 = _dafny.MultiSetOf(_this.F28())
		var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_1721_v10), 0))
		_ = _index224
		var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p1), 0))
		_ = _index225
		var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1721_v10), 0))
		_ = _index226
		var _rhs241 _dafny.Int = (_this).F30()
		_ = _rhs241
		var _rhs242 _dafny.Int = (_this).F41()
		_ = _rhs242
		var _rhs243 bool = !(((_this).F30()).Cmp(((_this).F30()).Minus((_1722_v11).Cardinality())) > 0)
		_ = _rhs243
		var _rhs244 bool = (_dafny.MultiSetOf(p2)).IsProperSubsetOf(_1723_v12)
		_ = _rhs244
		var _rhs245 _dafny.Int = ((_this).F30()).Minus((_this).F30())
		_ = _rhs245
		var _lhs188 *GlobalState = globalState
		_ = _lhs188
		var _lhs189 _dafny.Array = _1721_v10
		_ = _lhs189
		var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_1721_v10), 0))
		_ = _lhs190
		var _lhs191 _dafny.Array = p1
		_ = _lhs191
		var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((p1), 0))
		_ = _lhs192
		var _lhs193 *C7 = _this
		_ = _lhs193
		var _lhs194 _dafny.Array = _1721_v10
		_ = _lhs194
		var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.ArrayLen((_1721_v10), 0))
		_ = _lhs195
		_lhs188.F13 = _rhs241
		(_lhs189).ArraySet1(_rhs242, (_lhs190).Int())
		(_lhs191).ArraySet1(_rhs243, (_lhs192).Int())
		_lhs193.F31_set_(_rhs244)
		(_lhs194).ArraySet1(_rhs245, (_lhs195).Int())
		var _1724_v13 D5
		_ = _1724_v13
		_1724_v13 = Companion_D5_.Create_DC13_(p0, _1715_v4, _1716_v5)
		var _1725_v14 D5
		_ = _1725_v14
		_1725_v14 = Companion_D5_.Create_DC13_(_this.F28(), (_1724_v13).Dtor_cf28(), _1716_v5)
		var _pat_let_tv15 = _1716_v5
		_ = _pat_let_tv15
		var _pat_let_tv16 = _1716_v5
		_ = _pat_let_tv16
		r0 = func(_source29 D5) _dafny.CodePoint {
			if _source29.Is_DC13() {
				var _1726___mcc_h0 bool = _source29.Get_().(D5_DC13).Cf27
				_ = _1726___mcc_h0
				var _1727___mcc_h1 _dafny.Map = _source29.Get_().(D5_DC13).Cf28
				_ = _1727___mcc_h1
				var _1728___mcc_h2 _dafny.CodePoint = _source29.Get_().(D5_DC13).Cf29
				_ = _1728___mcc_h2
				var _1729_cf29 _dafny.CodePoint = _1728___mcc_h2
				_ = _1729_cf29
				var _1730_cf28 _dafny.Map = _1727___mcc_h1
				_ = _1730_cf28
				var _1731_cf27 bool = _1726___mcc_h0
				_ = _1731_cf27
				return _dafny.CodePoint('e')
			} else if _source29.Is_DC14() {
				var _1732___mcc_h3 bool = _source29.Get_().(D5_DC14).Cf30
				_ = _1732___mcc_h3
				var _1733___mcc_h4 _dafny.Int = _source29.Get_().(D5_DC14).Cf31
				_ = _1733___mcc_h4
				var _1734_cf31 _dafny.Int = _1733___mcc_h4
				_ = _1734_cf31
				var _1735_cf30 bool = _1732___mcc_h3
				_ = _1735_cf30
				return _pat_let_tv15
			} else {
				var _1736___mcc_h5 _dafny.Sequence = _source29.Get_().(D5_DC12).Cf26
				_ = _1736___mcc_h5
				var _1737_cf26 _dafny.Sequence = _1736___mcc_h5
				_ = _1737_cf26
				return _pat_let_tv16
			}
		}(_1724_v13)
		var _1738_v15 _dafny.Sequence
		_ = _1738_v15
		_1738_v15 = _dafny.SeqOf(_1720_v9, _1720_v9, _1720_v9, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm20((_this).F41(), p0, globalState), (_1722_v11).Cardinality()))
		r1 = (func() _dafny.Sequence {
			if p0 {
				return _1738_v15
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer107 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}((func(_1739_v9 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1740_i1 _dafny.Int) _dafny.Map {
					return _1739_v9
				}
			})(_1720_v9)))
		})()
		r2 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1717_v6).Cardinality()), (_1721_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_1721_v10), 0))).Int()).(_dafny.Int))
		r3 = _1717_v6
		return r0, r1, r2, r3
	}
}
func (_this *C7) M8(p0 _dafny.Array, globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1741_v0 T0
		_ = _1741_v0
		var _nw289 *C1 = New_C1_()
		_ = _nw289
		_nw289.Ctor__(_this.F29(), _this.F29(), true)
		_1741_v0 = _nw289
		var _1742_v1 _dafny.Map
		_ = _1742_v1
		_1742_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_v0, _dafny.IntOfInt64(161))
		var _1743_v2 _dafny.CodePoint
		_ = _1743_v2
		_1743_v2 = _dafny.CodePoint('w')
		var _1744_v3 _dafny.Map
		_ = _1744_v3
		_1744_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_v0.F28(), (_this).Fm19(globalState))
		var _1745_v4 T2
		_ = _1745_v4
		var _nw290 *C6 = New_C6_()
		_ = _nw290
		_nw290.Ctor__(_1743_v2, (_this).F30(), _this.F29(), _dafny.SeqOf(_this.F29()), Companion_Default___.Fm51(_1741_v0.F29(), _1743_v2, globalState), (_1744_v3).Cardinality(), _this.F29(), _this.F29(), _this.F29())
		_1745_v4 = _nw290
		var _1746_v5 _dafny.Map
		_ = _1746_v5
		_1746_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
			if (_1742_v1).Contains(_1741_v0) {
				return (_1742_v1).Get(_1741_v0).(_dafny.Int)
			}
			return (_this).F30()
		})()).Cmp((_this).F41()) <= 0, (_dafny.IntOfUint32((_dafny.SeqOf(_1745_v4, _1745_v4)).Cardinality())).Plus((_this).F30()))
		_1746_v5 = (_1746_v5).Update(_this.F31(), (_1745_v4).F30())
		var _1747_v6 D5
		_ = _1747_v6
		_1747_v6 = Companion_D5_.Create_DC13_(_1745_v4.F31(), _1746_v5, _1743_v2)
		var _1748_v7 _dafny.Map
		_ = _1748_v7
		_1748_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm24(_1741_v0.F29(), _dafny.IntOfInt64(487), globalState), (_1745_v4).F30())
		var _1749_v8 _dafny.Sequence
		_ = _1749_v8
		_1749_v8 = _dafny.UnicodeSeqOfUtf8Bytes("xbffckyc")
		var _1750_v9 *C5
		_ = _1750_v9
		var _nw291 *C5 = New_C5_()
		_ = _nw291
		_nw291.Ctor__((_1747_v6).Dtor_cf27(), _dafny.IntOfInt64(593), (_this).F41(), _this.F31(), (_1745_v4).F32(), (_1745_v4).F33(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1748_v7, (_dafny.Zero).Minus((_1745_v4).F30()))).Cardinality(), true, ((_1745_v4).F32()).Select((Companion_Default___.SafeIndex(((_this).F42()).Cardinality(), _dafny.IntOfUint32(((_1745_v4).F32()).Cardinality()))).Uint32()).(bool), !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(418))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg108 _dafny.Int) interface{} {
				return coer108(arg108)
			}
		}((func(_1751_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1752_i0 _dafny.Int) _dafny.CodePoint {
				return _1751_v2
			}
		})(_1743_v2))), _1749_v8))
		_1750_v9 = _nw291
		(_1741_v0).F29_set_(((_1745_v4).F32()).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((_1745_v4).F32()).Cardinality()))).Uint32()).(bool))
		var _1753_v10 _dafny.Sequence
		_ = _1753_v10
		_1753_v10 = _dafny.SeqOf((_1746_v5).Cardinality())
		var _hi9 _dafny.Int = (_1750_v9).F48()
		_ = _hi9
		for _1754_i1 := (_dafny.IntOfUint32((_1753_v10).Cardinality())).Minus((_this).F30()); _1754_i1.Cmp(_hi9) < 0; _1754_i1 = _1754_i1.Plus(_dafny.One) {
			var _1755_v11 D8
			_ = _1755_v11
			_1755_v11 = Companion_D8_.Create_DC24_((_1750_v9).F47(), (_this).F41(), (_1745_v4).F32(), (_1750_v9).F47(), _1753_v10)
			var _1756_v12 *C4
			_ = _1756_v12
			var _nw292 *C4 = New_C4_()
			_ = _nw292
			_nw292.Ctor__(_dafny.IntOfUint32((_1749_v8).Cardinality()), (_1750_v9).F47(), (_1755_v11).Dtor_cf48(), _dafny.Companion_Sequence_.Concatenate((_1745_v4).F33(), (_1745_v4).F33()), (_1745_v4).F30(), ((_1745_v4).F30()).Cmp((_1750_v9).F48()) < 0, (_1750_v9).F47(), _1741_v0.F29())
			_1756_v12 = _nw292
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((p0), 0))
			_ = _index227
			(p0).ArraySet1(_this.F29(), (_index227).Int())
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((p0), 0))
			_ = _index228
			(p0).ArraySet1((_1756_v12).Fm44((_dafny.Zero).Minus(_1754_i1), globalState), (_index228).Int())
			var _1757_v13 _dafny.Map
			_ = _1757_v13
			_1757_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1754_i1, _1741_v0.F28())
			_1757_v13 = (_1757_v13).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1745_v4).F32(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-685), _dafny.IntOfUint32(((_1745_v4).F32()).Cardinality()))).Uint32(), _1741_v0.F29())).Cardinality()), _1745_v4.F29())
			var _1758_v14 _dafny.Array
			_ = _1758_v14
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_39
			var _nw293 _dafny.Array
			_ = _nw293
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw293 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Int = func(_1759_i2 _dafny.Int) _dafny.Int {
					return (_1759_i2).Plus(_dafny.IntOfInt64(962))
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw293 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw293).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw293).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1758_v14 = _nw293
			var _1760_v15 _dafny.Map
			_ = _1760_v15
			_1760_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1758_v14, _dafny.IntOfUint32(((_1745_v4).F32()).Cardinality()))
			var _1761_v16 D7
			_ = _1761_v16
			_1761_v16 = Companion_D7_.Create_DC20_(Companion_D7_.Create_DC20_(Companion_D7_.Create_DC18_(_1760_v15)))
			var _1762_v17 D7
			_ = _1762_v17
			_1762_v17 = Companion_D7_.Create_DC20_(_1761_v16)
			var _pat_let_tv17 = _1761_v16
			_ = _pat_let_tv17
			var _1763_v18 _dafny.Array
			_ = _1763_v18
			var _nwElement0_61 D7 = _1762_v17
			_ = _nwElement0_61
			var _nw294 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(19))
			_ = _nw294
			(_nw294).ArraySet1(_nwElement0_61, 0)
			(_nw294).ArraySet1(_1762_v17, 1)
			(_nw294).ArraySet1(_1762_v17, 2)
			(_nw294).ArraySet1(_1762_v17, 3)
			(_nw294).ArraySet1(_1762_v17, 4)
			(_nw294).ArraySet1(_1762_v17, 5)
			(_nw294).ArraySet1(_1762_v17, 6)
			(_nw294).ArraySet1(Companion_D7_.Create_DC20_(_1761_v16), 7)
			(_nw294).ArraySet1(_1762_v17, 8)
			(_nw294).ArraySet1(_1762_v17, 9)
			(_nw294).ArraySet1(Companion_D7_.Create_DC20_(_1761_v16), 10)
			(_nw294).ArraySet1(_1762_v17, 11)
			(_nw294).ArraySet1(_1762_v17, 12)
			(_nw294).ArraySet1(_1762_v17, 13)
			(_nw294).ArraySet1(_1762_v17, 14)
			(_nw294).ArraySet1(_1762_v17, 15)
			(_nw294).ArraySet1(func(_pat_let23_0 D7) D7 {
				return func(_1764_dt__update__tmp_h0 D7) D7 {
					return func(_pat_let24_0 D7) D7 {
						return func(_1765_dt__update_hcf39_h0 D7) D7 {
							return Companion_D7_.Create_DC20_(_1765_dt__update_hcf39_h0)
						}(_pat_let24_0)
					}(_pat_let_tv17)
				}(_pat_let23_0)
			}(_1762_v17), 16)
			(_nw294).ArraySet1(_1762_v17, 17)
			(_nw294).ArraySet1(_1762_v17, 18)
			_1763_v18 = _nw294
			var _1766_v19 _dafny.Array
			_ = _1766_v19
			_1766_v19 = _1763_v18
			var _1767_v20 _dafny.Array
			_ = _1767_v20
			var _nwElement0_62 _dafny.Array = (func() _dafny.Array {
				if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
					return _1763_v18
				}
				return _1763_v18
			})()
			_ = _nwElement0_62
			var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(17))
			_ = _nw295
			(_nw295).ArraySet1(_nwElement0_62, 0)
			(_nw295).ArraySet1(_1763_v18, 1)
			(_nw295).ArraySet1(_1763_v18, 2)
			(_nw295).ArraySet1(_1763_v18, 3)
			(_nw295).ArraySet1(_1763_v18, 4)
			(_nw295).ArraySet1(_1763_v18, 5)
			(_nw295).ArraySet1((_1766_v19), 6)
			(_nw295).ArraySet1(_1763_v18, 7)
			(_nw295).ArraySet1(_1763_v18, 8)
			(_nw295).ArraySet1(_1763_v18, 9)
			(_nw295).ArraySet1(_1763_v18, 10)
			(_nw295).ArraySet1(_1763_v18, 11)
			(_nw295).ArraySet1(_1763_v18, 12)
			(_nw295).ArraySet1(_1763_v18, 13)
			(_nw295).ArraySet1(_1763_v18, 14)
			(_nw295).ArraySet1(_1763_v18, 15)
			(_nw295).ArraySet1(_1763_v18, 16)
			_1767_v20 = _nw295
			var _1768_v21 _dafny.Array
			_ = _1768_v21
			var _nw296 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
			_ = _nw296
			_1768_v21 = _nw296
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1768_v21), 0))
			_ = _index229
			(_1768_v21).ArraySet1CodePoint(_1743_v2, (_index229).Int())
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1768_v21), 0))
			_ = _index230
			var _rhs246 _dafny.Array = _1767_v20
			_ = _rhs246
			var _rhs247 _dafny.Int = Companion_Default___.SafeDivisionInt((_1750_v9).F48(), Companion_Default___.SafeModuloInt((_this).F41(), _1754_i1))
			_ = _rhs247
			var _rhs248 _dafny.CodePoint = _dafny.CodePoint('x')
			_ = _rhs248
			var _rhs249 bool = _1745_v4.F29()
			_ = _rhs249
			var _rhs250 bool = _1745_v4.F31()
			_ = _rhs250
			var _lhs196 _dafny.Array = _1768_v21
			_ = _lhs196
			var _lhs197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(294), _dafny.ArrayLen((_1768_v21), 0))
			_ = _lhs197
			var _lhs198 *GlobalState = globalState
			_ = _lhs198
			var _lhs199 T0 = _1741_v0
			_ = _lhs199
			_1767_v20 = _rhs246
			r1 = _rhs247
			(_lhs196).ArraySet1CodePoint(_rhs248, (_lhs197).Int())
			_lhs198.F12 = _rhs249
			_lhs199.F29_set_(_rhs250)
		}
		if _this.F29() {
			if (_1750_v9).F47() {
				var _1769_v22 _dafny.Array
				_ = _1769_v22
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_40
				var _nw297 _dafny.Array
				_ = _nw297
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw297 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Sequence = func(_1770_i3 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(false)
					}
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw297 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw297).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw297).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1769_v22 = _nw297
				var _rhs251 _dafny.Array = _1769_v22
				_ = _rhs251
				var _rhs252 _dafny.Int = (_1745_v4).F30()
				_ = _rhs252
				var _rhs253 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hgo"), _1749_v8)
				_ = _rhs253
				var _lhs200 *GlobalState = globalState
				_ = _lhs200
				_1769_v22 = _rhs251
				_lhs200.F13 = _rhs252
				_1749_v8 = _rhs253
				var _1771_v23 _dafny.Map
				_ = _1771_v23
				_1771_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1745_v4).F30(), _1746_v5)
				var _rhs254 _dafny.Int = Companion_Default___.Fm46(((func() _dafny.Map {
					if (_1771_v23).Contains(_dafny.IntOfUint32(((Companion_D8_.Create_DC24_((_1750_v9).F47(), (_1750_v9).F48(), (_1745_v4).F32(), _1741_v0.F29(), _1753_v10)).Dtor_cf48()).Cardinality())) {
						return (_1771_v23).Get(_dafny.IntOfUint32(((Companion_D8_.Create_DC24_((_1750_v9).F47(), (_1750_v9).F48(), (_1745_v4).F32(), _1741_v0.F29(), _1753_v10)).Dtor_cf48()).Cardinality())).(_dafny.Map)
					}
					return _1746_v5
				})()).Cardinality(), (_1750_v9).F48(), _1741_v0.F28(), globalState)
				_ = _rhs254
				var _rhs255 _dafny.CodePoint = _1743_v2
				_ = _rhs255
				var _lhs201 *GlobalState = globalState
				_ = _lhs201
				_lhs201.F13 = _rhs254
				_1743_v2 = _rhs255
				(globalState).F26 = false
				var _1772_v24 *C1
				_ = _1772_v24
				var _nw298 *C1 = New_C1_()
				_ = _nw298
				_nw298.Ctor__((_1750_v9).Fm6(globalState), _1741_v0.F28(), _1745_v4.F31())
				_1772_v24 = _nw298
				var _1773_v25 _dafny.Map
				_ = _1773_v25
				_1773_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1772_v24, _dafny.IntOfUint32((_1749_v8).Cardinality()))
				_1773_v25 = (_1773_v25).Update((func() *C1 {
					if true {
						return _1772_v24
					}
					return _1772_v24
				})(), (_1745_v4).F30())
				(globalState).F26 = (_1750_v9).F47()
			} else {
				var _1774_v26 _dafny.Array
				_ = _1774_v26
				var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
				_ = _nw299
				_1774_v26 = _nw299
				_1774_v26 = _1774_v26
				var _rhs256 bool = _1741_v0.F28()
				_ = _rhs256
				var _rhs257 _dafny.CodePoint = _1743_v2
				_ = _rhs257
				var _rhs258 _dafny.CodePoint = _1743_v2
				_ = _rhs258
				var _lhs202 T0 = _1741_v0
				_ = _lhs202
				_lhs202.F28_set_(_rhs256)
				_1743_v2 = _rhs257
				_1743_v2 = _rhs258
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((p0), 0))
				_ = _index231
				(p0).ArraySet1((_1750_v9).F47(), (_index231).Int())
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((p0), 0))
				_ = _index232
				(p0).ArraySet1((_1750_v9).F47(), (_index232).Int())
				var _1775_v27 *C1
				_ = _1775_v27
				var _nw300 *C1 = New_C1_()
				_ = _nw300
				_nw300.Ctor__(_1741_v0.F29(), false, ((_1745_v4).F30()).Cmp((_1750_v9).F48()) >= 0)
				_1775_v27 = _nw300
				var _1776_v28 _dafny.Array
				_ = _1776_v28
				var _nw301 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw301
				_1776_v28 = _nw301
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1776_v28), 0))
				_ = _index233
				(_1776_v28).ArraySet1((_this).F41(), (_index233).Int())
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1776_v28), 0))
				_ = _index234
				(_1776_v28).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_1745_v4).F30())), (_index234).Int())
			}
			var _1777_v29 _dafny.Map
			_ = _1777_v29
			_1777_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.Companion_Sequence_.Concatenate(_1749_v8, _1749_v8))
			_1777_v29 = ((_1777_v29).Merge(_1777_v29)).Merge(_1777_v29)
			var _1778_v30 _dafny.Array
			_ = _1778_v30
			var _len0_41 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_41
			var _nw302 _dafny.Array
			_ = _nw302
			if _len0_41.Cmp(_dafny.Zero) == 0 {
				_nw302 = _dafny.NewArray(_len0_41)
			} else {
				var _init41 func(_dafny.Int) _dafny.CodePoint = (func(_1779_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1780_i4 _dafny.Int) _dafny.CodePoint {
						return _1779_v2
					}
				})(_1743_v2)
				_ = _init41
				var _element0_41 = _init41(_dafny.Zero)
				_ = _element0_41
				_nw302 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
				(_nw302).ArraySet1CodePoint(_element0_41, 0)
				var _nativeLen0_41 = (_len0_41).Int()
				_ = _nativeLen0_41
				for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
					(_nw302).ArraySet1CodePoint(_init41(_dafny.IntOf(_i0_41)), _i0_41)
				}
			}
			_1778_v30 = _nw302
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_1778_v30), 0))
			_ = _index235
			(_1778_v30).ArraySet1CodePoint(_dafny.CodePoint('a'), (_index235).Int())
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_1778_v30), 0))
			_ = _index236
			(_1778_v30).ArraySet1CodePoint(_1743_v2, (_index236).Int())
			var _1781_v31 _dafny.MultiSet
			_ = _1781_v31
			_1781_v31 = _dafny.MultiSetOf(_this.F28())
			var _1782_v32 _dafny.Map
			_ = _1782_v32
			_1782_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1745_v4).F30(), false)
			r1 = (((_1781_v31).Update(_1741_v0.F29(), Companion_Default___.Abs(_dafny.IntOfUint32(((_1745_v4).F32()).Cardinality())))).Cardinality()).Minus(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1746_v5).Contains(!((Companion_D2_.Create_DC6_((_1745_v4).F30(), true, (_1745_v4).F30(), _1745_v4.F29(), (_1745_v4).F30())).Dtor_cf12())) {
					return (_1746_v5).Get(!((Companion_D2_.Create_DC6_((_1745_v4).F30(), true, (_1745_v4).F30(), _1745_v4.F29(), (_1745_v4).F30())).Dtor_cf12())).(_dafny.Int)
				}
				return (_1782_v32).Cardinality()
			})(), (_1745_v4).F30()))
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((p0), 0))
			_ = _index237
			(p0).ArraySet1(_1745_v4.F29(), (_index237).Int())
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((p0), 0))
			_ = _index238
			(p0).ArraySet1(_1741_v0.F29(), (_index238).Int())
		} else {
			var _1783_v33 _dafny.Map
			_ = _1783_v33
			_1783_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(665))).Uint32(), func(coer109 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}(func(_1784_i5 _dafny.Int) _dafny.Int {
				return (_this).F30()
			})), (_1745_v4).F30())
			(globalState).F13 = ((_1783_v33).Merge(_1783_v33)).Cardinality()
			(_1745_v4).F28_set_(_1745_v4.F28())
			var _1785_v34 D4
			_ = _1785_v34
			_1785_v34 = Companion_D4_.Create_DC11_((_1745_v4).F30(), _1741_v0.F28(), (_this).F41())
			var _1786_v35 _dafny.Map
			_ = _1786_v35
			_1786_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1785_v34, (_this).F41())
			var _1787_v36 _dafny.Sequence
			_ = _1787_v36
			_1787_v36 = _dafny.SeqOf(_1786_v35)
			var _rhs259 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(96), (_1750_v9).F48())
			_ = _rhs259
			var _rhs260 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(((_1745_v4).F30()).Minus(Companion_Default___.Fm46(_dafny.IntOfUint32((_1749_v8).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gtbdkgdg")).Cardinality()), _1741_v0.F28(), globalState))), (_1745_v4).F30())
			_ = _rhs260
			var _rhs261 _dafny.Sequence = _1749_v8
			_ = _rhs261
			var _rhs262 bool = Companion_Default___.Fm1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1787_v36, _dafny.Companion_Sequence_.Update(_1787_v36, (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_1787_v36).Cardinality()))).Uint32(), _1786_v35)), (_1750_v9).F48(), globalState)
			_ = _rhs262
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 *GlobalState = globalState
			_ = _lhs204
			var _lhs205 T2 = _1745_v4
			_ = _lhs205
			_lhs203.F13 = _rhs259
			_lhs204.F13 = _rhs260
			_1749_v8 = _rhs261
			_lhs205.F31_set_(_rhs262)
			var _1788_v37 _dafny.CodePoint
			_ = _1788_v37
			var _1789_v38 _dafny.Sequence
			_ = _1789_v38
			var _1790_v39 _dafny.Int
			_ = _1790_v39
			var _1791_v40 _dafny.Sequence
			_ = _1791_v40
			var _out60 _dafny.CodePoint
			_ = _out60
			var _out61 _dafny.Sequence
			_ = _out61
			var _out62 _dafny.Int
			_ = _out62
			var _out63 _dafny.Sequence
			_ = _out63
			_out60, _out61, _out62, _out63 = (_this).M7(!(_1741_v0.F28()), p0, Companion_Default___.Fm1((_1750_v9).F47(), (_this).F30(), globalState), globalState)
			_1788_v37 = _out60
			_1789_v38 = _out61
			_1790_v39 = _out62
			_1791_v40 = _out63
			(_1741_v0).F29_set_(!((func() bool {
				if _1741_v0.F28() {
					return (_1750_v9).F47()
				}
				return (_1745_v4.F29()) || (_1745_v4.F31())
			})()))
		}
		var _1792_v41 _dafny.Array
		_ = _1792_v41
		var _nw303 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw303
		_1792_v41 = _nw303
		for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1792_v41), 0))); ; {
			_guard_loop_3, _ok46 := _iter46()
			if !_ok46 {
				break
			}
			var _1793_i6 _dafny.Int
			_1793_i6 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1793_i6).Sign() != -1) && ((_1793_i6).Cmp(_dafny.ArrayLen((_1792_v41), 0)) < 0)) {
				(_1792_v41).ArraySet1(Companion_Default___.SafeDivisionInt(_1793_i6, Companion_Default___.SafeDivisionInt((_1745_v4).F30(), _dafny.IntOfUint32((_1749_v8).Cardinality()))), (_1793_i6).Int())
			}
		}
		r0 = p0
		var _1794_v42 _dafny.Set
		_ = _1794_v42
		_1794_v42 = _dafny.SetOf(_dafny.IntOfUint32((_1749_v8).Cardinality()))
		var _1795_v43 _dafny.Map
		_ = _1795_v43
		_1795_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1745_v4.F29(), _1794_v42)
		r1 = (_dafny.Zero).Minus((((_1794_v42).Intersection(_1794_v42)).Intersection((func() _dafny.Set {
			if (_1795_v43).Contains(_1741_v0.F29()) {
				return (_1795_v43).Get(_1741_v0.F29()).(_dafny.Set)
			}
			return _dafny.SetOf((_this).F30())
		})())).Cardinality())
		return r0, r1
	}
}
func (_this *C7) F41() _dafny.Int {
	{
		return _this._f41
	}
}
func (_this *C7) F42() _dafny.MultiSet {
	{
		return _this._f42
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f36 _dafny.CodePoint
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
	F40  bool
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f36 = _dafny.CodePoint('D')
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
	_this.F40 = false
	return &_this
}

type CompanionStruct_C8_ struct {
}

var Companion_C8_ = CompanionStruct_C8_{}

func (_this *C8) Equals(other *C8) bool {
	return _this == other
}

func (_this *C8) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C8)
	return ok && _this.Equals(other)
}

func (*C8) String() string {
	return "_module.C8"
}

func Type_C8_() _dafny.TypeDescriptor {
	return type_C8_{}
}

type type_C8_ struct {
}

func (_this type_C8_) Default() interface{} {
	return (*C8)(nil)
}

func (_this type_C8_) String() string {
	return "main.C8"
}
func (_this *C8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T2_.TraitID_, Companion_T3_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C8{}
var _ T2 = &C8{}
var _ T3 = &C8{}
var _ T1 = &C8{}
var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F28() bool {
	return _this._f28
}
func (_this *C8) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C8) F29() bool {
	return _this._f29
}
func (_this *C8) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C8) F31() bool {
	return _this._f31
}
func (_this *C8) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C8) F35() bool {
	return _this._f35
}
func (_this *C8) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C8) F36() _dafny.CodePoint {
	return _this._f36
}
func (_this *C8) F36_set_(value _dafny.CodePoint) {
	_this._f36 = value
}
func (_this *C8) F34() _dafny.Int {
	return _this._f34
}
func (_this *C8) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C8) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C8) F30() _dafny.Int {
	return _this._f30
}
func (_this *C8) Ctor__(f40 bool, f36 _dafny.CodePoint, f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this).F40 = f40
		(_this)._f36 = f36
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C8) Fm7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return _this.F35()
	}
}
func (_this *C8) Fm6(globalState *GlobalState) bool {
	{
		return _this.F35()
	}
}
func (_this *C8) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf((_this).F34())
	}
}
func (_this *C8) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		if _this.F40 {
			return (_this).F32()
		} else {
			return _dafny.Companion_Sequence_.Concatenate((_this).F32(), _dafny.SeqOf(_this.F29()))
		}
	}
}
func (_this *C8) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if _this.F28() {
				return _dafny.UnicodeSeqOfUtf8Bytes("qlucvm")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("ef")
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sgolrgbnu"), _dafny.UnicodeSeqOfUtf8Bytes("tbm")))
	}
}
func (_this *C8) Fm2(globalState *GlobalState) _dafny.Map {
	{
		var _source30 D4 = Companion_D4_.Create_DC10_(_dafny.MultiSetOf(_dafny.IntOfInt64(747), _dafny.IntOfInt64(611)))
		_ = _source30
		if _source30.Is_DC11() {
			var _1796___mcc_h0 _dafny.Int = _source30.Get_().(D4_DC11).Cf23
			_ = _1796___mcc_h0
			var _1797___mcc_h1 bool = _source30.Get_().(D4_DC11).Cf24
			_ = _1797___mcc_h1
			var _1798___mcc_h2 _dafny.Int = _source30.Get_().(D4_DC11).Cf25
			_ = _1798___mcc_h2
			var _1799_cf25 _dafny.Int = _1798___mcc_h2
			_ = _1799_cf25
			var _1800_cf24 bool = _1797___mcc_h1
			_ = _1800_cf24
			var _1801_cf23 _dafny.Int = _1796___mcc_h0
			_ = _1801_cf23
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus((_this).F30()))
		} else {
			var _1802___mcc_h3 _dafny.MultiSet = _source30.Get_().(D4_DC10).Cf22
			_ = _1802___mcc_h3
			var _1803_cf22 _dafny.MultiSet = _1802___mcc_h3
			_ = _1803_cf22
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_dafny.Zero).Minus((_this).F30()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F40, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _this.F31())).Cardinality()))
		}
	}
}
func (_this *C8) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		(_this).F36_set_(_dafny.CodePoint('a'))
		if (((_this).F30()).Cmp((_this).F30()) < 0) && (_this.F28()) {
			var _1804_v0 _dafny.Sequence
			_ = _1804_v0
			_1804_v0 = _dafny.UnicodeSeqOfUtf8Bytes("suoxd")
			var _1805_v1 _dafny.MultiSet
			_ = _1805_v1
			_1805_v1 = _dafny.MultiSetOf((_this).F34(), _dafny.IntOfUint32((_1804_v0).Cardinality()))
			var _1806_v2 _dafny.Sequence
			_ = _1806_v2
			_1806_v2 = _dafny.SeqOf((_this).F30())
			var _1807_v3 _dafny.Sequence
			_ = _1807_v3
			_1807_v3 = _dafny.SeqOf(_dafny.SeqOf((_this).F34(), (_1805_v1).Cardinality(), (_this).F30()), _1806_v2, _1806_v2, _1806_v2)
			var _1808_v4 _dafny.Array
			_ = _1808_v4
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_42
			var _nw304 _dafny.Array
			_ = _nw304
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw304 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.Int = func(_1809_i0 _dafny.Int) _dafny.Int {
					return (_1809_i0).Times((_this).F34())
				}
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw304 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw304).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw304).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1808_v4 = _nw304
			var _1810_v5 _dafny.Map
			_ = _1810_v5
			_1810_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1808_v4, _1804_v0)
			var _1811_v6 _dafny.MultiSet
			_ = _1811_v6
			_1811_v6 = _dafny.MultiSetOf(_this.F31())
			var _1812_v7 T0
			_ = _1812_v7
			var _nw305 *C7 = New_C7_()
			_ = _nw305
			_nw305.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-494), (_this).F30()), (Companion_Default___.Fm49(_1805_v1, _this.F36(), _1807_v3, globalState)).Difference(_1805_v1), (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1808_v4, _1804_v0)).Update(_1808_v4, _1804_v0)).Merge(_1810_v5)).Cardinality(), _this.F40, _dafny.Companion_Sequence_.Contains(_1804_v0, Companion_Default___.Fm34((_this).F34(), (_this).F30(), (_1811_v6).Cardinality(), !(_this.F35()), globalState)), _this.F29())
			_1812_v7 = _nw305
			_1812_v7 = _1812_v7
			r1 = (_dafny.Zero).Minus(Companion_Default___.Fm39(true, _1805_v1, _this.F40, globalState))
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1808_v4), 0))
			_ = _index239
			(_1808_v4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32())).Cardinality()), (_index239).Int())
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1808_v4), 0))
			_ = _index240
			(_1808_v4).ArraySet1((_this).F34(), (_index240).Int())
			var _1813_v9 _dafny.Map
			_ = _1813_v9
			_1813_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1808_v4)
			var _1814_v10 _dafny.Map
			_ = _1814_v10
			_1814_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), true)
			var _1815_v11 _dafny.Map
			_ = _1815_v11
			_1815_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1813_v9).Cardinality(), _1814_v10)
			var _1816_v13 _dafny.Set
			_ = _1816_v13
			_1816_v13 = _dafny.SetOf(func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter47 := _dafny.Iterate(((_1815_v11).Update((_this).F34(), _1814_v10)).Keys().Elements()); ; {
					_compr_43, _ok47 := _iter47()
					if !_ok47 {
						break
					}
					var _1817_v8 _dafny.Int
					_1817_v8 = interface{}(_compr_43).(_dafny.Int)
					if ((_1815_v11).Update((_this).F34(), _1814_v10)).Contains(_1817_v8) {
						_coll43.Add((_1817_v8).Plus(_dafny.IntOfInt64(24)), _this.F29())
					}
				}
				return _coll43.ToMap()
			}(), func() _dafny.Map {
				var _coll44 = _dafny.NewMapBuilder()
				_ = _coll44
				for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(985), _dafny.IntOfInt64(854))); ; {
					_compr_44, _ok48 := _iter48()
					if !_ok48 {
						break
					}
					var _1818_v12 _dafny.Int
					_1818_v12 = interface{}(_compr_44).(_dafny.Int)
					if ((_dafny.IntOfInt64(985)).Cmp(_1818_v12) <= 0) && ((_1818_v12).Cmp(_dafny.IntOfInt64(854)) < 0) {
						_coll44.Add(Companion_Default___.SafeDivisionInt(_1818_v12, (_1808_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1808_v4), 0))).Int()).(_dafny.Int)), false)
					}
				}
				return _coll44.ToMap()
			}(), _1814_v10)
			_1816_v13 = _dafny.SetOf(_1814_v10, func() _dafny.Map {
				var _coll45 = _dafny.NewMapBuilder()
				_ = _coll45
				for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-637), _dafny.IntOfInt64(477))); ; {
					_compr_45, _ok49 := _iter49()
					if !_ok49 {
						break
					}
					var _1819_v14 _dafny.Int
					_1819_v14 = interface{}(_compr_45).(_dafny.Int)
					if ((_dafny.IntOfInt64(-637)).Cmp(_1819_v14) <= 0) && ((_1819_v14).Cmp(_dafny.IntOfInt64(477)) < 0) {
						_coll45.Add((_1819_v14).Times((_this).F30()), _1812_v7.F28())
					}
				}
				return _coll45.ToMap()
			}())
			_1811_v6 = _1811_v6
		} else {
			var _1820_v15 _dafny.Map
			_ = _1820_v15
			_1820_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F35())
			var _1821_v16 _dafny.Sequence
			_ = _1821_v16
			_1821_v16 = _dafny.SeqOf((_this).F30(), (_this).F34())
			var _1822_v17 _dafny.Map
			_ = _1822_v17
			_1822_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1820_v15, _1821_v16)
			var _1823_v18 _dafny.Int
			_ = _1823_v18
			var _out64 _dafny.Int
			_ = _out64
			_out64 = Companion_Default___.M0((func() _dafny.Sequence {
				if (_1822_v17).Contains(_1820_v15) {
					return (_1822_v17).Get(_1820_v15).(_dafny.Sequence)
				}
				return _1821_v16
			})(), _this.F28(), ((_this).F30()).Minus(_dafny.IntOfInt64(-749)), _this.F29(), globalState)
			_1823_v18 = _out64
			(globalState).F13 = (_this).F34()
			var _1824_v19 _dafny.Array
			_ = _1824_v19
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_43
			var _nw306 _dafny.Array
			_ = _nw306
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw306 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.Int = func(_1825_i1 _dafny.Int) _dafny.Int {
					return (_1825_i1).Times(_dafny.IntOfUint32((_dafny.SeqOf(_this.F28(), _this.F29())).Cardinality()))
				}
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw306 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw306).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw306).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1824_v19 = _nw306
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1824_v19), 0))
			_ = _index241
			(_1824_v19).ArraySet1((_this).F30(), (_index241).Int())
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1824_v19), 0))
			_ = _index242
			(_1824_v19).ArraySet1((_this).F34(), (_index242).Int())
			(globalState).F1 = _this.F28()
			_1820_v15 = (_1820_v15).Update(_this.F40, !((func() bool {
				if _this.F28() {
					return _this.F40
				}
				return _this.F31()
			})()))
		}
		var _1826_i2 _dafny.Int
		_ = _1826_i2
		_1826_i2 = _dafny.Zero
		{
			for (func() bool {
				if _this.F29() {
					return !((_this.F40) == (true))
				}
				return true
			})() {
				{
					if (_1826_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1826_i2 = (_1826_i2).Plus(_dafny.One)
					var _1827_v20 _dafny.Sequence
					_ = _1827_v20
					_1827_v20 = _dafny.UnicodeSeqOfUtf8Bytes("idpxt")
					var _1828_v21 _dafny.Sequence
					_ = _1828_v21
					_1828_v21 = _dafny.SeqOf(_this.F35(), _this.F35(), ((_this).F30()).Cmp(_dafny.IntOfUint32((_1827_v20).Cardinality())) >= 0)
					_1828_v21 = _1828_v21
					var _1829_v22 _dafny.Array
					_ = _1829_v22
					var _len0_44 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_44
					var _nw307 _dafny.Array
					_ = _nw307
					if _len0_44.Cmp(_dafny.Zero) == 0 {
						_nw307 = _dafny.NewArray(_len0_44)
					} else {
						var _init44 func(_dafny.Int) _dafny.CodePoint = func(_1830_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						}
						_ = _init44
						var _element0_44 = _init44(_dafny.Zero)
						_ = _element0_44
						_nw307 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
						(_nw307).ArraySet1CodePoint(_element0_44, 0)
						var _nativeLen0_44 = (_len0_44).Int()
						_ = _nativeLen0_44
						for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
							(_nw307).ArraySet1CodePoint(_init44(_dafny.IntOf(_i0_44)), _i0_44)
						}
					}
					_1829_v22 = _nw307
					var _1831_v23 *C0
					_ = _1831_v23
					var _nw308 *C0 = New_C0_()
					_ = _nw308
					_nw308.Ctor__(_1829_v22)
					_1831_v23 = _nw308
					var _source31 D9 = Companion_D9_.Create_DC27_(_1831_v23, _1827_v20)
					_ = _source31
					if _source31.Is_DC27() {
						var _1832___mcc_h0 *C0 = _source31.Get_().(D9_DC27).Cf53
						_ = _1832___mcc_h0
						var _1833___mcc_h1 _dafny.Sequence = _source31.Get_().(D9_DC27).Cf54
						_ = _1833___mcc_h1
						var _1834_cf54 _dafny.Sequence = _1833___mcc_h1
						_ = _1834_cf54
						var _1835_cf53 *C0 = _1832___mcc_h0
						_ = _1835_cf53
						var _1836_v24 D9
						_ = _1836_v24
						_1836_v24 = Companion_D9_.Create_DC27_(_1835_cf53, (func() _dafny.Sequence {
							if ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool) {
								return _1827_v20
							}
							return _1834_cf54
						})())
						var _rhs263 D9 = _1836_v24
						_ = _rhs263
						var _rhs264 _dafny.Sequence = (_this).Fm3(_this.F31(), _this.F40, globalState)
						_ = _rhs264
						var _rhs265 bool = _this.F29()
						_ = _rhs265
						_1836_v24 = _rhs263
						r2 = _rhs264
						r0 = _rhs265
						(globalState).F12 = _this.F35()
						(globalState).F13 = (func() _dafny.Int {
							if _this.F28() {
								return (_this).F34()
							}
							return (_dafny.Zero).Minus((_this).F30())
						})()
						var _1837_v25 _dafny.Set
						_ = _1837_v25
						_1837_v25 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}(func(_1838_i4 _dafny.Int) _dafny.Int {
							return (_this).F30()
						})))
						var _1839_v26 _dafny.Map
						_ = _1839_v26
						_1839_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _1837_v25)
						var _1840_v27 _dafny.Sequence
						_ = _1840_v27
						_1840_v27 = _dafny.SeqOf((_this).F34())
						var _1841_v29 _dafny.Map
						_ = _1841_v29
						_1841_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1834_cf54).Cardinality()), _this.F36())
						var _1842_v31 D16
						_ = _1842_v31
						_1842_v31 = Companion_D16_.Create_DC39_(_1837_v25)
						var _1843_v32 _dafny.Sequence
						_ = _1843_v32
						_1843_v32 = _dafny.SeqOf(_dafny.SeqOf((_this).F30(), (_this).F34()), _1840_v27, _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("emyatjs")).Cardinality())), (_this).F34()))
						var _1844_v34 _dafny.Array
						_ = _1844_v34
						var _nwElement0_63 _dafny.Set = _1837_v25
						_ = _nwElement0_63
						var _nw309 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(23))
						_ = _nw309
						(_nw309).ArraySet1(_nwElement0_63, 0)
						(_nw309).ArraySet1((_1837_v25).Intersection(_1837_v25), 1)
						(_nw309).ArraySet1((func() _dafny.Set {
							if (_1839_v26).Contains((_this).F34()) {
								return (_1839_v26).Get((_this).F34()).(_dafny.Set)
							}
							return _1837_v25
						})(), 2)
						(_nw309).ArraySet1(_1837_v25, 3)
						(_nw309).ArraySet1(_1837_v25, 4)
						(_nw309).ArraySet1(_1837_v25, 5)
						(_nw309).ArraySet1((func() _dafny.Set {
							if _this.F40 {
								return _dafny.SetOf(_1840_v27, _1840_v27)
							}
							return _1837_v25
						})(), 6)
						(_nw309).ArraySet1(_1837_v25, 7)
						(_nw309).ArraySet1(_1837_v25, 8)
						(_nw309).ArraySet1((_dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F34()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-949), _dafny.IntOfUint32((_dafny.SeqOf((_this).F34())).Cardinality()))).Uint32(), (_this).F34()), _1840_v27, _1840_v27)).Intersection(_1837_v25), 9)
						(_nw309).ArraySet1(_1837_v25, 10)
						(_nw309).ArraySet1(_1837_v25, 11)
						(_nw309).ArraySet1(func() _dafny.Set {
							var _coll46 = _dafny.NewBuilder()
							_ = _coll46
							for _iter50 := _dafny.Iterate((_1837_v25).Elements()); ; {
								_compr_46, _ok50 := _iter50()
								if !_ok50 {
									break
								}
								var _1845_v28 _dafny.Sequence
								_1845_v28 = interface{}(_compr_46).(_dafny.Sequence)
								if (_1837_v25).Contains(_1845_v28) {
									_coll46.Add(_1845_v28)
								}
							}
							return _coll46.ToSet()
						}(), 12)
						(_nw309).ArraySet1((_dafny.SetOf(_dafny.SeqOf((_this).F34(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), false)).Cardinality(), (func() _dafny.Set {
							var _coll47 = _dafny.NewBuilder()
							_ = _coll47
							for _iter51 := _dafny.Iterate((_1841_v29).Keys().Elements()); ; {
								_compr_47, _ok51 := _iter51()
								if !_ok51 {
									break
								}
								var _1846_v30 _dafny.Int
								_1846_v30 = interface{}(_compr_47).(_dafny.Int)
								if (_1841_v29).Contains(_1846_v30) {
									_coll47.Add((_1846_v30).Times(_dafny.IntOfInt64(-831)))
								}
							}
							return _coll47.ToSet()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wt")).Cardinality())))).Difference((_1842_v31).Dtor_cf72()), 13)
						(_nw309).ArraySet1(_1837_v25, 14)
						(_nw309).ArraySet1(_1837_v25, 15)
						(_nw309).ArraySet1(_1837_v25, 16)
						(_nw309).ArraySet1(_1837_v25, 17)
						(_nw309).ArraySet1(_1837_v25, 18)
						(_nw309).ArraySet1((Companion_Default___.Fm54(_dafny.IntOfInt64(536), _1840_v27, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg111 _dafny.Int) interface{} {
								return coer111(arg111)
							}
						}(func(_1847_i5 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('w')
						}))).Cardinality()), globalState)).Difference(_1837_v25), 19)
						(_nw309).ArraySet1(_1837_v25, 20)
						(_nw309).ArraySet1(_dafny.SetOf(_1840_v27), 21)
						(_nw309).ArraySet1(func() _dafny.Set {
							var _coll48 = _dafny.NewBuilder()
							_ = _coll48
							for _iter52 := _dafny.Iterate((_1843_v32).Elements()); ; {
								_compr_48, _ok52 := _iter52()
								if !_ok52 {
									break
								}
								var _1848_v33 _dafny.Sequence
								_1848_v33 = interface{}(_compr_48).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_1843_v32, _1848_v33) {
									_coll48.Add(_1848_v33)
								}
							}
							return _coll48.ToSet()
						}(), 22)
						_1844_v34 = _nw309
						var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1844_v34), 0))
						_ = _index243
						(_1844_v34).ArraySet1(_1837_v25, (_index243).Int())
						var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1844_v34), 0))
						_ = _index244
						var _rhs266 _dafny.Set = _1837_v25
						_ = _rhs266
						var _rhs267 _dafny.Int = ((_this).F30()).Minus((_this).F30())
						_ = _rhs267
						var _rhs268 _dafny.Int = Companion_Default___.SafeModuloInt(((_this).F30()).Minus((func() _dafny.Map {
							var _coll49 = _dafny.NewMapBuilder()
							_ = _coll49
							for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(642), _dafny.IntOfInt64(927))); ; {
								_compr_49, _ok53 := _iter53()
								if !_ok53 {
									break
								}
								var _1849_v35 _dafny.Int
								_1849_v35 = interface{}(_compr_49).(_dafny.Int)
								if ((_dafny.IntOfInt64(642)).Cmp(_1849_v35) <= 0) && ((_1849_v35).Cmp(_dafny.IntOfInt64(927)) < 0) {
									_coll49.Add((_1849_v35).Times((_this).F34()), (_this).F30())
								}
							}
							return _coll49.ToMap()
						}()).Cardinality()), (Companion_Default___.Fm32(_this.F40, globalState)).Minus(Companion_Default___.Fm32(_this.F35(), globalState)))
						_ = _rhs268
						var _rhs269 _dafny.Int = (_this).F34()
						_ = _rhs269
						var _lhs206 _dafny.Array = _1844_v34
						_ = _lhs206
						var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_1844_v34), 0))
						_ = _lhs207
						var _lhs208 *GlobalState = globalState
						_ = _lhs208
						var _lhs209 *GlobalState = globalState
						_ = _lhs209
						(_lhs206).ArraySet1(_rhs266, (_lhs207).Int())
						_lhs208.F13 = _rhs267
						r1 = _rhs268
						_lhs209.F13 = _rhs269
					} else {
						var _1850___mcc_h2 *C2 = _source31.Get_().(D9_DC26).Cf52
						_ = _1850___mcc_h2
						var _1851_cf52 *C2 = _1850___mcc_h2
						_ = _1851_cf52
						var _1852_v36 _dafny.Array
						_ = _1852_v36
						var _nw310 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
						_ = _nw310
						_1852_v36 = _nw310
						(globalState).F21 = (func() _dafny.Array {
							if _this.F40 {
								return _1852_v36
							}
							return _1852_v36
						})()
						var _1853_v37 _dafny.Array
						_ = _1853_v37
						var _len0_45 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_45
						var _nw311 _dafny.Array
						_ = _nw311
						if _len0_45.Cmp(_dafny.Zero) == 0 {
							_nw311 = _dafny.NewArray(_len0_45)
						} else {
							var _init45 func(_dafny.Int) _dafny.Int = (func(_1854_v20 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
								return func(_1855_i6 _dafny.Int) _dafny.Int {
									return (_1855_i6).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1854_v20).Cardinality()), (_this).F30())).Cardinality()))
								}
							})(_1827_v20)
							_ = _init45
							var _element0_45 = _init45(_dafny.Zero)
							_ = _element0_45
							_nw311 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
							(_nw311).ArraySet1(_element0_45, 0)
							var _nativeLen0_45 = (_len0_45).Int()
							_ = _nativeLen0_45
							for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
								(_nw311).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
							}
						}
						_1853_v37 = _nw311
						var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1853_v37), 0))
						_ = _index245
						(_1853_v37).ArraySet1(((_this).F34()).Minus((_this).F30()), (_index245).Int())
						var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1853_v37), 0))
						_ = _index246
						var _rhs270 bool = ((_this).F30()).Cmp(Companion_Default___.SafeModuloInt((_1851_cf52).F46(), (_1851_cf52).F46())) > 0
						_ = _rhs270
						var _rhs271 _dafny.Int = (_this).F30()
						_ = _rhs271
						var _lhs210 *GlobalState = globalState
						_ = _lhs210
						var _lhs211 _dafny.Array = _1853_v37
						_ = _lhs211
						var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1853_v37), 0))
						_ = _lhs212
						_lhs210.F1 = _rhs270
						(_lhs211).ArraySet1(_rhs271, (_lhs212).Int())
						var _1856_v38 _dafny.Map
						_ = _1856_v38
						_1856_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F40)
						var _1857_v39 *C2
						_ = _1857_v39
						var _nw312 *C2 = New_C2_()
						_ = _nw312
						_nw312.Ctor__(Companion_Default___.SafeDivisionInt((_1856_v38).Cardinality(), (_1853_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1853_v37), 0))).Int()).(_dafny.Int)), Companion_Default___.Fm24(_this.F31(), _dafny.IntOfUint32((_1827_v20).Cardinality()), globalState), _this.F29(), _this.F35(), true)
						_1857_v39 = _nw312
						(globalState).F13 = ((_1857_v39).F46()).Minus(_dafny.IntOfInt64(514))
					}
					(globalState).F12 = _this.F40
					var _1858_v40 _dafny.Array
					_ = _1858_v40
					var _nw313 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
					_ = _nw313
					_1858_v40 = _nw313
					var _1859_v41 _dafny.Array
					_ = _1859_v41
					var _nw314 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
					_ = _nw314
					_1859_v41 = _nw314
					var _nwElement0_64 _dafny.Array = (func() _dafny.Array {
						if _this.F31() {
							return _1859_v41
						}
						return _1859_v41
					})()
					_ = _nwElement0_64
					var _nw315 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(24))
					_ = _nw315
					(_nw315).ArraySet1(_nwElement0_64, 0)
					(_nw315).ArraySet1(_1859_v41, 1)
					(_nw315).ArraySet1(_1859_v41, 2)
					(_nw315).ArraySet1(_1859_v41, 3)
					(_nw315).ArraySet1(_1859_v41, 4)
					(_nw315).ArraySet1(_1859_v41, 5)
					(_nw315).ArraySet1(_1859_v41, 6)
					(_nw315).ArraySet1(_1859_v41, 7)
					(_nw315).ArraySet1(_1859_v41, 8)
					(_nw315).ArraySet1(_1859_v41, 9)
					(_nw315).ArraySet1(_1859_v41, 10)
					(_nw315).ArraySet1(_1859_v41, 11)
					(_nw315).ArraySet1(_1859_v41, 12)
					(_nw315).ArraySet1(_1859_v41, 13)
					(_nw315).ArraySet1(_1859_v41, 14)
					(_nw315).ArraySet1(_1859_v41, 15)
					(_nw315).ArraySet1(_1859_v41, 16)
					(_nw315).ArraySet1(_1859_v41, 17)
					(_nw315).ArraySet1(_1859_v41, 18)
					(_nw315).ArraySet1(_1859_v41, 19)
					(_nw315).ArraySet1(_1859_v41, 20)
					(_nw315).ArraySet1(_1859_v41, 21)
					(_nw315).ArraySet1(_1859_v41, 22)
					(_nw315).ArraySet1(_1859_v41, 23)
					_1858_v40 = _nw315
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		(_this).F40 = _this.F28()
		(_this).F35_set_(((_this).F30()).Cmp((_this).F30()) != 0)
		var _1860_v42 _dafny.MultiSet
		_ = _1860_v42
		_1860_v42 = _dafny.MultiSetOf(Companion_Default___.Fm1(_this.F29(), (_this).F30(), globalState))
		var _1861_v43 D0
		_ = _1861_v43
		_1861_v43 = Companion_D0_.Create_DC1_(_1860_v42)
		var _1862_v44 _dafny.Map
		_ = _1862_v44
		_1862_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1861_v43, _this.F40)
		var _1863_v45 D8
		_ = _1863_v45
		_1863_v45 = Companion_D8_.Create_DC21_(_1862_v44)
		var _1864_v46 _dafny.Sequence
		_ = _1864_v46
		_1864_v46 = _dafny.SeqOf(_1863_v45)
		var _1865_v47 _dafny.Sequence
		_ = _1865_v47
		_1865_v47 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer112 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
			return func(arg112 _dafny.Int) interface{} {
				return coer112(arg112)
			}
		}(func(_1866_i7 _dafny.Int) D8 {
			return Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_this.F31())), _this.F29()))
		})), _1864_v46, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(47))).Uint32(), func(coer113 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
			return func(arg113 _dafny.Int) interface{} {
				return coer113(arg113)
			}
		}((func(_1867_v45 D8) func(_dafny.Int) D8 {
			return func(_1868_i8 _dafny.Int) D8 {
				return _1867_v45
			}
		})(_1863_v45))), _dafny.SeqOf(_1863_v45, _1863_v45, _1863_v45, _1863_v45))
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_1865_v47, _dafny.SeqOf(_dafny.SeqOf(_1863_v45))) {
			var _1869_v48 _dafny.Array
			_ = _1869_v48
			var _nw316 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw316
			_1869_v48 = _nw316
			var _1870_v49 *C3
			_ = _1870_v49
			var _nw317 *C3 = New_C3_()
			_ = _nw317
			_nw317.Ctor__((_this).F34(), _1869_v48, _dafny.CodePoint('i'), _dafny.IntOfInt64(101), _this.F35(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), (_this).F32()), (_this).F33(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), _this.F28(), _this.F29(), (false) && (false))
			_1870_v49 = _nw317
			var _1871_v50 _dafny.Set
			_ = _1871_v50
			_1871_v50 = _dafny.SetOf(_this.F35())
			var _1872_v51 _dafny.Array
			_ = _1872_v51
			var _len0_46 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_46
			var _nw318 _dafny.Array
			_ = _nw318
			if _len0_46.Cmp(_dafny.Zero) == 0 {
				_nw318 = _dafny.NewArray(_len0_46)
			} else {
				var _init46 func(_dafny.Int) _dafny.CodePoint = func(_1873_i9 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}
				_ = _init46
				var _element0_46 = _init46(_dafny.Zero)
				_ = _element0_46
				_nw318 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
				(_nw318).ArraySet1CodePoint(_element0_46, 0)
				var _nativeLen0_46 = (_len0_46).Int()
				_ = _nativeLen0_46
				for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
					(_nw318).ArraySet1CodePoint(_init46(_dafny.IntOf(_i0_46)), _i0_46)
				}
			}
			_1872_v51 = _nw318
			var _1874_v52 *C0
			_ = _1874_v52
			var _nw319 *C0 = New_C0_()
			_ = _nw319
			_nw319.Ctor__(_1872_v51)
			_1874_v52 = _nw319
			var _1875_v53 _dafny.Map
			_ = _1875_v53
			var _1876_v54 bool
			_ = _1876_v54
			var _1877_v55 bool
			_ = _1877_v55
			var _1878_v56 _dafny.Int
			_ = _1878_v56
			var _out65 _dafny.Map
			_ = _out65
			var _out66 bool
			_ = _out66
			var _out67 bool
			_ = _out67
			var _out68 _dafny.Int
			_ = _out68
			_out65, _out66, _out67, _out68 = (_this).M6(Companion_Default___.SafeDivisionInt((_this).F34(), (_1870_v49).F44()), Companion_Default___.Fm39(true, _dafny.MultiSetOf((_1871_v50).Cardinality()), _this.F31(), globalState), _this.F36(), _1874_v52, globalState)
			_1875_v53 = _out65
			_1876_v54 = _out66
			_1877_v55 = _out67
			_1878_v56 = _out68
			var _1879_v57 _dafny.Sequence
			_ = _1879_v57
			_1879_v57 = _dafny.UnicodeSeqOfUtf8Bytes("fu")
			r2 = _dafny.Companion_Sequence_.Concatenate(_1879_v57, _1879_v57)
			(_this).F36_set_(_this.F36())
			var _1880_v58 _dafny.Map
			_ = _1880_v58
			_1880_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_1860_v42).Cardinality())
			_1880_v58 = (_1880_v58).Update((_this).Fm6(globalState), (_this).F34())
		} else {
			r1 = ((_this).F30()).Times((_this).F30())
			r1 = (_dafny.IntOfUint32(((_this).Fm3(_this.F31(), _this.F35(), globalState)).Cardinality())).Minus((_this).F30())
			(globalState).F1 = _this.F35()
			var _1881_v59 _dafny.Array
			_ = _1881_v59
			var _nw320 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw320
			_1881_v59 = _nw320
			var _1882_v60 *C4
			_ = _1882_v60
			var _nw321 *C4 = New_C4_()
			_ = _nw321
			_nw321.Ctor__(_dafny.IntOfInt64(657), false, (_this).F32(), (_this).F33(), (_this).F34(), false, _this.F31(), _this.F28())
			_1882_v60 = _nw321
			var _1883_v61 _dafny.Sequence
			_ = _1883_v61
			_1883_v61 = _dafny.SeqOf(_1882_v60, _1882_v60, _1882_v60, _1882_v60, _1882_v60)
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1881_v59), 0))
			_ = _index247
			(_1881_v59).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1883_v61, _1883_v61), (_index247).Int())
			var _1884_v62 D17
			_ = _1884_v62
			_1884_v62 = Companion_D17_.Create_DC42_(_1883_v61)
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1881_v59), 0))
			_ = _index248
			var _rhs272 _dafny.Sequence = (_1884_v62).Dtor_cf76()
			_ = _rhs272
			var _rhs273 bool = ((_this).F30()).Cmp((_this).F34()) >= 0
			_ = _rhs273
			var _lhs213 _dafny.Array = _1881_v59
			_ = _lhs213
			var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1881_v59), 0))
			_ = _lhs214
			var _lhs215 *GlobalState = globalState
			_ = _lhs215
			(_lhs213).ArraySet1(_rhs272, (_lhs214).Int())
			_lhs215.F1 = _rhs273
			var _1885_v63 _dafny.Map
			_ = _1885_v63
			_1885_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F30())
			_1885_v63 = (_1885_v63).Update(_this.F28(), (_this).F34())
		}
		r0 = _this.F35()
		var _1886_v64 _dafny.Map
		_ = _1886_v64
		_1886_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _this.F36())
		r1 = ((_1886_v64).Cardinality()).Times((_this).F34())
		var _1887_v65 _dafny.Sequence
		_ = _1887_v65
		_1887_v65 = _dafny.UnicodeSeqOfUtf8Bytes("v")
		r2 = _1887_v65
		return r0, r1, r2
	}
}
func (_this *C8) M6(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 *C0, globalState *GlobalState) (_dafny.Map, bool, bool, _dafny.Int) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1888_v0 _dafny.Set
		_ = _1888_v0
		_1888_v0 = _dafny.SetOf(_this.F35(), _this.F29(), _this.F35())
		var _hi10 _dafny.Int = Companion_Default___.SafeDivisionInt((_1888_v0).Cardinality(), (_this).F30())
		_ = _hi10
		for _1889_i0 := (func() _dafny.Int {
			if _this.F31() {
				return p1
			}
			return (_this).F30()
		})(); _1889_i0.Cmp(_hi10) < 0; _1889_i0 = _1889_i0.Plus(_dafny.One) {
			var _1890_v1 _dafny.Sequence
			_ = _1890_v1
			_1890_v1 = _dafny.UnicodeSeqOfUtf8Bytes("vq")
			var _1891_v2 _dafny.Sequence
			_ = _1891_v2
			_1891_v2 = _dafny.SeqOf(p1, _1889_i0, _dafny.IntOfInt64(43), _1889_i0, p0)
			var _1892_v4 _dafny.Map
			_ = _1892_v4
			_1892_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F34())).Cardinality()), p1)
			var _1893_v5 _dafny.Array
			_ = _1893_v5
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_47
			var _nw322 _dafny.Array
			_ = _nw322
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw322 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) _dafny.Int = func(_1894_i1 _dafny.Int) _dafny.Int {
					return (_1894_i1).Plus((_dafny.Zero).Minus((_this).F34()))
				}
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw322 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw322).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw322).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1893_v5 = _nw322
			var _1895_v6 _dafny.Map
			_ = _1895_v6
			_1895_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1893_v5, _this.F29())
			var _1896_v7 T2
			_ = _1896_v7
			var _nw323 *C5 = New_C5_()
			_ = _nw323
			_nw323.Ctor__(_this.F35(), (func() _dafny.Map {
				var _coll50 = _dafny.NewMapBuilder()
				_ = _coll50
				for _iter54 := _dafny.Iterate((_1892_v4).Keys().Elements()); ; {
					_compr_50, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1897_v3 _dafny.Int
					_1897_v3 = interface{}(_compr_50).(_dafny.Int)
					if (_1892_v4).Contains(_1897_v3) {
						_coll50.Add((_1897_v3).Minus((_this).F30()), _this.F28())
					}
				}
				return _coll50.ToMap()
			}()).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), _this.F28(), Companion_Default___.Fm0((_this).F34(), false, _this.F40, _dafny.IntOfUint32((_1890_v1).Cardinality()), globalState), _dafny.Companion_Sequence_.Update((_this).F33(), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32(), (_this).F32()), (_1895_v6).Cardinality(), false, _this.F31(), true)
			_1896_v7 = _nw323
			var _1898_v8 _dafny.Map
			_ = _1898_v8
			_1898_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1891_v2, _1896_v7)
			(globalState).F13 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1890_v1, (Companion_Default___.SafeIndex((_1898_v8).Cardinality(), _dafny.IntOfUint32((_1890_v1).Cardinality()))).Uint32(), _this.F36())).Cardinality())
			(_this).F35_set_(false)
			(_this).F36_set_(_this.F36())
			var _1899_v9 D16
			_ = _1899_v9
			_1899_v9 = Companion_D16_.Create_DC40_(_dafny.UnicodeSeqOfUtf8Bytes("rcj"), ((_dafny.SetOf((_dafny.MultiSetFromSeq(_1891_v2)).Cardinality(), (_this).F34())).Cardinality()).Plus((_1896_v7).F30()))
			var _source32 D16 = _1899_v9
			_ = _source32
			if _source32.Is_DC40() {
				var _1900___mcc_h0 _dafny.Sequence = _source32.Get_().(D16_DC40).Cf73
				_ = _1900___mcc_h0
				var _1901___mcc_h1 _dafny.Int = _source32.Get_().(D16_DC40).Cf74
				_ = _1901___mcc_h1
				var _1902_cf74 _dafny.Int = _1901___mcc_h1
				_ = _1902_cf74
				var _1903_cf73 _dafny.Sequence = _1900___mcc_h0
				_ = _1903_cf73
				var _1904_v10 _dafny.Array
				_ = _1904_v10
				var _nw324 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw324
				_1904_v10 = _nw324
				var _1905_v11 _dafny.Sequence
				_ = _1905_v11
				_1905_v11 = _dafny.SeqOf(_1904_v10, _1904_v10)
				var _1906_v12 _dafny.Sequence
				_ = _1906_v12
				_1906_v12 = _1905_v11
				_1905_v11 = (_1906_v12)
				(globalState).F26 = _this.F35()
				var _1907_v13 _dafny.Set
				_ = _1907_v13
				_1907_v13 = _dafny.SetOf(Companion_Default___.Fm55((_dafny.Zero).Minus(p0), _1896_v7.F29(), (_dafny.SetOf(true, _1896_v7.F31())).Cardinality(), globalState), _1888_v0)
				_1907_v13 = _1907_v13
				var _1908_v14 _dafny.Sequence
				_ = _1908_v14
				_1908_v14 = _dafny.SeqOf(_1888_v0)
				var _1909_v15 D2
				_ = _1909_v15
				_1909_v15 = Companion_D2_.Create_DC6_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F34())).Cardinality()), _1896_v7.F29(), _dafny.IntOfInt64(423), _this.F29(), ((_1908_v14).Select((Companion_Default___.SafeIndex((_1896_v7).F30(), _dafny.IntOfUint32((_1908_v14).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())
				_1909_v15 = _1909_v15
			} else if _source32.Is_DC39() {
				var _1910___mcc_h2 _dafny.Set = _source32.Get_().(D16_DC39).Cf72
				_ = _1910___mcc_h2
				var _1911_cf72 _dafny.Set = _1910___mcc_h2
				_ = _1911_cf72
				var _1912_v16 _dafny.Array
				_ = _1912_v16
				var _nw325 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw325
				_1912_v16 = _nw325
				var _1913_v17 *C4
				_ = _1913_v17
				var _nw326 *C4 = New_C4_()
				_ = _nw326
				_nw326.Ctor__(_1889_i0, Companion_Default___.Fm1(_1896_v7.F31(), p0, globalState), _dafny.SeqOf(_1896_v7.F31()), (_1896_v7).F33(), _dafny.IntOfUint32((Companion_Default___.Fm0(_dafny.IntOfInt64(553), _this.F35(), _this.F29(), p1, globalState)).Cardinality()), _1896_v7.F31(), false, _1896_v7.F28())
				_1913_v17 = _nw326
				var _1914_v18 *C3
				_ = _1914_v18
				var _nw327 *C3 = New_C3_()
				_ = _nw327
				_nw327.Ctor__((_this).F34(), (func() _dafny.Array {
					if _1896_v7.F31() {
						return _1912_v16
					}
					return _1912_v16
				})(), _this.F36(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-686))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg114 _dafny.Int) interface{} {
						return coer114(arg114)
					}
				}((func(_1915_v7 T2) func(_dafny.Int) _dafny.Int {
					return func(_1916_i2 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32(((_1915_v7).F32()).Cardinality())
					}
				})(_1896_v7)))).Cardinality()), _this.F31(), _dafny.Companion_Sequence_.Concatenate((_1896_v7).F32(), _dafny.SeqOf(_1896_v7.F31(), _this.F31(), !(_1896_v7.F29()))), (_1896_v7).F33(), (_1896_v7).F30(), Companion_Default___.Fm1(_this.F35(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1913_v17)).Update((_dafny.Zero).Minus(_dafny.IntOfUint32((_1890_v1).Cardinality())), _1913_v17)).Cardinality(), globalState), _this.F40, _1896_v7.F29())
				_1914_v18 = _nw327
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1893_v5), 0))
				_ = _index249
				(_1893_v5).ArraySet1((_this).F30(), (_index249).Int())
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_1893_v5), 0))
				_ = _index250
				(_1893_v5).ArraySet1(Companion_Default___.SafeDivisionInt((_1896_v7).F30(), (_1896_v7).F30()), (_index250).Int())
				var _1917_v19 _dafny.Array
				_ = _1917_v19
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_48
				var _nw328 _dafny.Array
				_ = _nw328
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw328 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) bool = (func(_1918_v7 T2) func(_dafny.Int) bool {
						return func(_1919_i3 _dafny.Int) bool {
							return (_dafny.MultiSetOf(_dafny.MultiSetOf(_this.F29(), _this.F40))).IsProperSubsetOf((_dafny.MultiSetOf(_dafny.MultiSetOf(_1918_v7.F28()))))
						}
					})(_1896_v7)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw328 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw328).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw328).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1917_v19 = _nw328
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_1917_v19), 0))
				_ = _index251
				(_1917_v19).ArraySet1(_this.F28(), (_index251).Int())
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_1917_v19), 0))
				_ = _index252
				(_1917_v19).ArraySet1((p1).Cmp(p0) < 0, (_index252).Int())
				_1917_v19 = _1917_v19
			} else {
				var _1920___mcc_h3 D16 = _source32.Get_().(D16_DC41).Cf75
				_ = _1920___mcc_h3
				var _1921_cf75 D16 = _1920___mcc_h3
				_ = _1921_cf75
				(globalState).F1 = _this.F40
				var _1922_v20 _dafny.Array
				_ = _1922_v20
				var _nw329 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw329
				_1922_v20 = _nw329
				var _1923_v21 _dafny.Array
				_ = _1923_v21
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_49
				var _nw330 _dafny.Array
				_ = _nw330
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw330 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) bool = (func(_1924_v7 T2) func(_dafny.Int) bool {
						return func(_1925_i4 _dafny.Int) bool {
							return _1924_v7.F28()
						}
					})(_1896_v7)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw330 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw330).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw330).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1923_v21 = _nw330
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1922_v20), 0))
				_ = _index253
				(_1922_v20).ArraySet1(_1923_v21, (_index253).Int())
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1922_v20), 0))
				_ = _index254
				var _rhs274 _dafny.Int = _dafny.IntOfInt64(522)
				_ = _rhs274
				var _rhs275 _dafny.Array = _1923_v21
				_ = _rhs275
				var _lhs216 _dafny.Array = _1922_v20
				_ = _lhs216
				var _lhs217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(720), _dafny.ArrayLen((_1922_v20), 0))
				_ = _lhs217
				r3 = _rhs274
				(_lhs216).ArraySet1(_rhs275, (_lhs217).Int())
				var _1926_v22 _dafny.Map
				_ = _1926_v22
				_1926_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1923_v21, _dafny.IntOfUint32((_1890_v1).Cardinality()))
				var _1927_v23 *C4
				_ = _1927_v23
				var _nw331 *C4 = New_C4_()
				_ = _nw331
				_nw331.Ctor__((_1896_v7).F30(), _1896_v7.F28(), (_1896_v7).F32(), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_1896_v7).F33(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_1896_v7).F33()).Cardinality()))).Uint32(), (_this).F32()), (_this).F33()), Companion_Default___.SafeModuloInt((_1896_v7).F30(), _dafny.IntOfInt64(143)), ((_1896_v7).F30()).Cmp((_1926_v22).Cardinality()) == 0, _this.F31(), _1896_v7.F28())
				_1927_v23 = _nw331
				var _1928_v24 _dafny.Set
				_ = _1928_v24
				_1928_v24 = _dafny.SetOf(_this.F36(), Companion_Default___.Fm21(p2, p0, globalState))
				var _1929_v25 _dafny.Map
				_ = _1929_v25
				_1929_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())
				var _1930_v26 D2
				_ = _1930_v26
				_1930_v26 = Companion_D2_.Create_DC6_((_1928_v24).Cardinality(), _1896_v7.F29(), _dafny.IntOfInt64(230), _1896_v7.F31(), (func() _dafny.Int {
					if (_1929_v25).Contains(_this.F28()) {
						return (_1929_v25).Get(_this.F28()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_1891_v2).Cardinality())
				})())
				var _1931_v27 _dafny.Array
				_ = _1931_v27
				var _nw332 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
				_ = _nw332
				_1931_v27 = _nw332
				var _1932_v28 _dafny.Sequence
				_ = _1932_v28
				_1932_v28 = _dafny.SeqOf(_1890_v1)
				var _rhs276 D2 = Companion_Default___.Fm56((_this).F30(), _dafny.Companion_Sequence_.IsPrefixOf(_1890_v1, _1890_v1), ((_dafny.Zero).Minus((_1896_v7).F30())).Times(p0), globalState)
				_ = _rhs276
				var _rhs277 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, (_this).F30())
				_ = _rhs277
				var _rhs278 _dafny.Int = p1
				_ = _rhs278
				var _rhs279 _dafny.Array = _1931_v27
				_ = _rhs279
				var _rhs280 bool = !_dafny.Companion_Sequence_.Equal((_1932_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1890_v1).Cardinality()), _dafny.IntOfUint32((_1932_v28).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(443))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg115 _dafny.Int) interface{} {
						return coer115(arg115)
					}
				}((func(_1933_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1934_i5 _dafny.Int) _dafny.CodePoint {
						return _1933_p2
					}
				})(p2))), _1890_v1))
				_ = _rhs280
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				var _lhs220 T2 = _1896_v7
				_ = _lhs220
				_1930_v26 = _rhs276
				_lhs218.F13 = _rhs277
				_lhs219.F13 = _rhs278
				_1931_v27 = _rhs279
				_lhs220.F28_set_(_rhs280)
			}
		}
		var _1935_v29 _dafny.Array
		_ = _1935_v29
		var _nw333 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw333
		_1935_v29 = _nw333
		var _1936_v30 _dafny.Map
		_ = _1936_v30
		_1936_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F29())
		var _1937_v31 _dafny.Map
		_ = _1937_v31
		_1937_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F40, (_1936_v30).Cardinality())
		var _1938_v32 _dafny.Map
		_ = _1938_v32
		_1938_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F40), Companion_Default___.Fm57(p2, (_1937_v31).Cardinality(), globalState))
		var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))
		_ = _index255
		(_1935_v29).ArraySet1(Companion_Default___.Fm1(_this.F29(), Companion_Default___.Fm24(false, (_1938_v32).Cardinality(), globalState), globalState), (_index255).Int())
		var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))
		_ = _index256
		(_1935_v29).ArraySet1(((_this).F32()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F34())), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), (_index256).Int())
		var _1939_i6 _dafny.Int
		_ = _1939_i6
		_1939_i6 = _dafny.Zero
		{
			for !(_this.F28()) {
				{
					if (_1939_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1939_i6 = (_1939_i6).Plus(_dafny.One)
					var _1940_v33 _dafny.MultiSet
					_ = _1940_v33
					_1940_v33 = _dafny.MultiSetOf(_this.F40)
					var _1941_v34 _dafny.MultiSet
					_ = _1941_v34
					_1941_v34 = _dafny.MultiSetOf((_this).F34(), p1)
					var _1942_v35 _dafny.Sequence
					_ = _1942_v35
					_1942_v35 = _dafny.SeqOf((_this).F34())
					var _1943_v36 D4
					_ = _1943_v36
					_1943_v36 = Companion_D4_.Create_DC11_(p1, _this.F40, (func() _dafny.Int {
						if (_1940_v33).Contains(_this.F28()) {
							return (_1940_v33).Multiplicity(_this.F28())
						}
						return p0
					})())
					var _1944_v37 _dafny.MultiSet
					_ = _1944_v37
					_1944_v37 = _dafny.MultiSetOf(_1943_v36)
					var _1945_v38 _dafny.Sequence
					_ = _1945_v38
					_1945_v38 = _dafny.UnicodeSeqOfUtf8Bytes("ff")
					var _1946_v39 T0
					_ = _1946_v39
					var _nw334 *C7 = New_C7_()
					_ = _nw334
					_nw334.Ctor__((func() _dafny.Int {
						if (_1940_v33).Contains((_1935_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))).Int()).(bool)) {
							return (_1940_v33).Multiplicity((_1935_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))).Int()).(bool))
						}
						return p0
					})(), (_1941_v34).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_1942_v35).Cardinality()))), p0, (_1944_v37).IsProperSubsetOf(_1944_v37), ((_this).F32()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1945_v38).Cardinality()), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), _this.F29())
					_1946_v39 = _nw334
					var _1947_v40 _dafny.Array
					_ = _1947_v40
					var _nw335 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
					_ = _nw335
					_1947_v40 = _nw335
					var _1948_v41 _dafny.Sequence
					_ = _1948_v41
					_1948_v41 = _dafny.SeqOf(Companion_Default___.Fm49(_1941_v34, _this.F36(), _dafny.SeqOf(_1942_v35, _dafny.SeqOf(_dafny.IntOfUint32((_1945_v38).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1947_v40, _this.F31())).Cardinality())), globalState), _1941_v34)
					var _rhs281 T0 = _1946_v39
					_ = _rhs281
					var _rhs282 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1948_v41, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(963))).Uint32(), func(coer116 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
						return func(arg116 _dafny.Int) interface{} {
							return coer116(arg116)
						}
					}((func(_1949_v35 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
						return func(_1950_i7 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetFromSeq(_1949_v35)
						}
					})(_1942_v35)))), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1948_v41, (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1948_v41).Cardinality()))).Uint32(), _1941_v34), _dafny.Companion_Sequence_.Update(_1948_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.IntOfUint32((_1948_v41).Cardinality()))).Uint32(), _dafny.MultiSetOf((_this).F34()))))
					_ = _rhs282
					_1946_v39 = _rhs281
					r1 = _rhs282
					var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))
					_ = _index257
					var _rhs283 bool = (_this).Fm7(_1941_v34, (_dafny.Zero).Minus((_this).F30()), globalState)
					_ = _rhs283
					var _rhs284 _dafny.Int = ((_this).F30()).Times((_dafny.Zero).Minus((_this).F34()))
					_ = _rhs284
					var _lhs221 _dafny.Array = _1935_v29
					_ = _lhs221
					var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))
					_ = _lhs222
					var _lhs223 *GlobalState = globalState
					_ = _lhs223
					(_lhs221).ArraySet1(_rhs283, (_lhs222).Int())
					_lhs223.F13 = _rhs284
					(globalState).F21 = _1935_v29
					var _1951_v42 _dafny.Array
					_ = _1951_v42
					var _nw336 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
					_ = _nw336
					_1951_v42 = _nw336
					var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1951_v42), 0))
					_ = _index258
					(_1951_v42).ArraySet1((_this).F32(), (_index258).Int())
					var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1951_v42), 0))
					_ = _index259
					(_1951_v42).ArraySet1((_this).F32(), (_index259).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1952_v43 _dafny.Sequence
		_ = _1952_v43
		_1952_v43 = _dafny.UnicodeSeqOfUtf8Bytes("opebgfkic")
		var _1953_v44 _dafny.Array
		_ = _1953_v44
		var _nwElement0_65 _dafny.Int = p1
		_ = _nwElement0_65
		var _nw337 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(13))
		_ = _nw337
		(_nw337).ArraySet1(_nwElement0_65, 0)
		(_nw337).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F30())), 1)
		(_nw337).ArraySet1(p0, 2)
		(_nw337).ArraySet1(_dafny.IntOfUint32((_1952_v43).Cardinality()), 3)
		(_nw337).ArraySet1(p0, 4)
		(_nw337).ArraySet1((_this).F34(), 5)
		(_nw337).ArraySet1((_this).F30(), 6)
		(_nw337).ArraySet1((_this).F34(), 7)
		(_nw337).ArraySet1((_this).F30(), 8)
		(_nw337).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1952_v43).Cardinality())), 9)
		(_nw337).ArraySet1(p1, 10)
		(_nw337).ArraySet1((_this).F30(), 11)
		(_nw337).ArraySet1(p0, 12)
		_1953_v44 = _nw337
		var _1954_v45 _dafny.Sequence
		_ = _1954_v45
		_1954_v45 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F28())).Cardinality()), (_this).F30())
		var _1955_v46 _dafny.Map
		_ = _1955_v46
		_1955_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1953_v44, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm23((_1935_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))).Int()).(bool), p0, globalState), _1954_v45))
		var _1956_v47 _dafny.Set
		_ = _1956_v47
		_1956_v47 = _dafny.SetOf(_dafny.IntOfUint32((_1954_v45).Cardinality()))
		var _1957_v48 _dafny.Map
		_ = _1957_v48
		_1957_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1956_v47, _1955_v46)
		var _1958_v49 _dafny.Set
		_ = _1958_v49
		_1958_v49 = _1888_v0
		var _1959_v50 _dafny.Map
		_ = _1959_v50
		_1959_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F28()), _1955_v46)
		var _1960_v51 _dafny.Map
		_ = _1960_v51
		_1960_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1958_v49).Cardinality(), (func() _dafny.Map {
			if (_1959_v50).Contains(_this.F40) {
				return (_1959_v50).Get(_this.F40).(_dafny.Map)
			}
			return _1955_v46
		})())
		_1955_v46 = (func() _dafny.Map {
			if (_1957_v48).Contains(_1956_v47) {
				return (_1957_v48).Get(_1956_v47).(_dafny.Map)
			}
			return (func() _dafny.Map {
				if (_1960_v51).Contains(p0) {
					return (_1960_v51).Get(p0).(_dafny.Map)
				}
				return _1955_v46
			})()
		})()
		var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_1953_v44), 0))
		_ = _index260
		(_1953_v44).ArraySet1((_this).F30(), (_index260).Int())
		var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_1953_v44), 0))
		_ = _index261
		(_1953_v44).ArraySet1((_dafny.Zero).Minus(((func() _dafny.Int {
			if _this.F29() {
				return (_this).F30()
			}
			return p0
		})()).Times(_dafny.IntOfInt64(-661))), (_index261).Int())
		(_this).F29_set_(_this.F29())
		var _1961_v52 _dafny.Map
		_ = _1961_v52
		_1961_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SeqOf(_1937_v31, _1937_v31))
		var _1962_v53 _dafny.Map
		_ = _1962_v53
		_1962_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(26), _this.F31())
		var _1963_v54 _dafny.Sequence
		_ = _1963_v54
		_1963_v54 = _dafny.SeqOf(_1937_v31)
		var _1964_v55 D3
		_ = _1964_v55
		_1964_v55 = Companion_D3_.Create_DC7_((func() _dafny.Sequence {
			if (_1961_v52).Contains((_dafny.Zero).Minus((_1962_v53).Cardinality())) {
				return (_1961_v52).Get((_dafny.Zero).Minus((_1962_v53).Cardinality())).(_dafny.Sequence)
			}
			return _1963_v54
		})())
		var _1965_v56 _dafny.Map
		_ = _1965_v56
		_1965_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1964_v55, (_this).F30())
		var _1966_v57 _dafny.Map
		_ = _1966_v57
		_1966_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1965_v56, _1953_v44)
		r0 = _1966_v57
		r1 = (_this).Fm7(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(868))).Uint32(), func(coer117 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg117 _dafny.Int) interface{} {
				return coer117(arg117)
			}
		}((func(_1967_v45 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_1968_i8 _dafny.Int) _dafny.Int {
				return (Companion_D2_.Create_DC6_(_dafny.IntOfInt64(438), _this.F35(), _dafny.IntOfInt64(871), _this.F28(), _dafny.IntOfUint32((_1967_v45).Cardinality()))).Dtor_cf11()
			}
		})(_1954_v45)))), Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf((_1935_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(82), _dafny.ArrayLen((_1935_v29), 0))).Int()).(bool), _this.F40, _this.F35())).Cardinality(), _dafny.IntOfUint32((_1952_v43).Cardinality())), globalState)
		r2 = _this.F28()
		r3 = p1
		return r0, r1, r2, r3
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f36 _dafny.CodePoint
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
	_f38 _dafny.Int
	_f39 _dafny.Map
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f36 = _dafny.CodePoint('D')
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
	_this._f38 = _dafny.Zero
	_this._f39 = _dafny.EmptyMap
	return &_this
}

type CompanionStruct_C9_ struct {
}

var Companion_C9_ = CompanionStruct_C9_{}

func (_this *C9) Equals(other *C9) bool {
	return _this == other
}

func (_this *C9) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C9)
	return ok && _this.Equals(other)
}

func (*C9) String() string {
	return "_module.C9"
}

func Type_C9_() _dafny.TypeDescriptor {
	return type_C9_{}
}

type type_C9_ struct {
}

func (_this type_C9_) Default() interface{} {
	return (*C9)(nil)
}

func (_this type_C9_) String() string {
	return "main.C9"
}
func (_this *C9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C9{}
var _ T3 = &C9{}
var _ T2 = &C9{}
var _ T1 = &C9{}
var _ T0 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) F28() bool {
	return _this._f28
}
func (_this *C9) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C9) F29() bool {
	return _this._f29
}
func (_this *C9) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C9) F31() bool {
	return _this._f31
}
func (_this *C9) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C9) F35() bool {
	return _this._f35
}
func (_this *C9) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C9) F36() _dafny.CodePoint {
	return _this._f36
}
func (_this *C9) F36_set_(value _dafny.CodePoint) {
	_this._f36 = value
}
func (_this *C9) F34() _dafny.Int {
	return _this._f34
}
func (_this *C9) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C9) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C9) F30() _dafny.Int {
	return _this._f30
}
func (_this *C9) Ctor__(f38 _dafny.Int, f39 _dafny.Map, f36 _dafny.CodePoint, f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f38 = f38
		(_this)._f39 = f39
		(_this)._f36 = f36
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C9) Fm7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		if true {
			return false
		} else {
			return (Companion_D0_.Create_DC0_(_this.F35())).Dtor_cf0()
		}
	}
}
func (_this *C9) Fm6(globalState *GlobalState) bool {
	{
		return !((func() _dafny.Map {
			if _this.F29() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _this.F35())
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (Companion_D0_.Create_DC0_(_this.F35())).Dtor_cf0())
		})()).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)), _this.F31())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _this.F28())))
	}
}
func (_this *C9) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf((((Companion_D4_.Create_DC10_(_dafny.MultiSetOf((_this).F34()))).Dtor_cf22()).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(571), ((Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_this.F28()))).Dtor_cf1()).Cardinality(), (_this).F34(), (_this).F30(), (func() _dafny.Map {
			var _coll51 = _dafny.NewMapBuilder()
			_ = _coll51
			for _iter55 := _dafny.Iterate((_dafny.SetOf((_this).F38())).Elements()); ; {
				_compr_51, _ok55 := _iter55()
				if !_ok55 {
					break
				}
				var _1969_v0 _dafny.Int
				_1969_v0 = interface{}(_compr_51).(_dafny.Int)
				if (_dafny.SetOf((_this).F38())).Contains(_1969_v0) {
					_coll51.Add(Companion_Default___.SafeDivisionInt(_1969_v0, (_this).F38()), _this.F31())
				}
			}
			return _coll51.ToMap()
		}()).Cardinality()))).Cardinality())
	}
}
func (_this *C9) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return (_this).F32()
	}
}
func (_this *C9) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nxjjkbq"), _dafny.UnicodeSeqOfUtf8Bytes("buh"))
	}
}
func (_this *C9) Fm2(globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), (_this).F38())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F35()), (_this).F34())))
	}
}
func (_this *C9) Fm16(p0 bool, p1 D0, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf((_dafny.Zero).Minus((_this).F30()))).Intersection(_dafny.SetOf((_this).F38(), (_this).F30()))).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(715), (_this).F30()))
	}
}
func (_this *C9) Fm17(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_this.F35()), _dafny.MultiSetOf(_this.F28(), !(_this.F31()), _this.F31()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true, _this.F31()), _dafny.MultiSetOf(_this.F28())))).Cardinality()).Plus((_this).F34())
	}
}
func (_this *C9) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _1970_v0 _dafny.Set
		_ = _1970_v0
		_1970_v0 = _dafny.SetOf(_this.F35(), _this.F35())
		if (_1970_v0).IsDisjointFrom(_dafny.SetOf(_this.F29(), _this.F35())) {
			var _1971_v1 _dafny.Map
			_ = _1971_v1
			_1971_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_this).F30())
			if ((func() _dafny.Int {
				if (_1971_v1).Contains(_this.F31()) {
					return (_1971_v1).Get(_this.F31()).(_dafny.Int)
				}
				return (_this).F34()
			})()).Cmp((_this).F30()) <= 0 {
				var _1972_v2 _dafny.Array
				_ = _1972_v2
				var _len0_50 _dafny.Int = _dafny.One
				_ = _len0_50
				var _nw338 _dafny.Array
				_ = _nw338
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw338 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) _dafny.CodePoint = func(_1973_i0 _dafny.Int) _dafny.CodePoint {
						return _this.F36()
					}
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw338 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw338).ArraySet1CodePoint(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw338).ArraySet1CodePoint(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_1972_v2 = _nw338
				var _1974_v3 *C0
				_ = _1974_v3
				var _nw339 *C0 = New_C0_()
				_ = _nw339
				_nw339.Ctor__(_1972_v2)
				_1974_v3 = _nw339
				var _1975_v4 D1
				_ = _1975_v4
				_1975_v4 = Companion_D1_.Create_DC3_(_this.F35(), _dafny.IntOfInt64(777), _1974_v3, (_dafny.Zero).Minus((_this).F34()))
				_1975_v4 = _1975_v4
				var _1976_v5 _dafny.Sequence
				_ = _1976_v5
				_1976_v5 = _dafny.UnicodeSeqOfUtf8Bytes("iyotbppo")
				var _1977_v6 D5
				_ = _1977_v6
				_1977_v6 = Companion_D5_.Create_DC12_(_1976_v5)
				var _1978_v7 D0
				_ = _1978_v7
				_1978_v7 = Companion_D0_.Create_DC0_(_this.F35())
				var _1979_v8 D2
				_ = _1979_v8
				_1979_v8 = Companion_D2_.Create_DC6_((_this).F30(), _this.F35(), _dafny.IntOfUint32((_1976_v5).Cardinality()), !((_1978_v7).Dtor_cf0()), _dafny.IntOfInt64(-57))
				var _1980_v9 _dafny.MultiSet
				_ = _1980_v9
				_1980_v9 = _dafny.MultiSetOf(_this.F28(), _this.F35(), _this.F28())
				var _1981_v10 _dafny.Array
				_ = _1981_v10
				var _nwElement0_66 _dafny.Int = (_this).F30()
				_ = _nwElement0_66
				var _nw340 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(10))
				_ = _nw340
				(_nw340).ArraySet1(_nwElement0_66, 0)
				(_nw340).ArraySet1((_this).F30(), 1)
				(_nw340).ArraySet1(((_this).F30()).Times((_this).F34()), 2)
				(_nw340).ArraySet1((_dafny.IntOfUint32(((_1977_v6).Dtor_cf26()).Cardinality())).Plus((_this).F38()), 3)
				(_nw340).ArraySet1((_this).F34(), 4)
				(_nw340).ArraySet1((_1979_v8).Dtor_cf11(), 5)
				(_nw340).ArraySet1((_this).F34(), 6)
				(_nw340).ArraySet1((_1980_v9).Cardinality(), 7)
				(_nw340).ArraySet1((_this).F34(), 8)
				(_nw340).ArraySet1((_this).F38(), 9)
				_1981_v10 = _nw340
				var _1982_v11 _dafny.Set
				_ = _1982_v11
				_1982_v11 = _dafny.SetOf((_this).F34(), (_this).F34())
				var _1983_v12 _dafny.MultiSet
				_ = _1983_v12
				_1983_v12 = _dafny.MultiSetOf((_this).F30())
				var _1984_v13 _dafny.Array
				_ = _1984_v13
				var _nwElement0_67 bool = true
				_ = _nwElement0_67
				var _nw341 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(16))
				_ = _nw341
				(_nw341).ArraySet1(_nwElement0_67, 0)
				(_nw341).ArraySet1(_this.F31(), 1)
				(_nw341).ArraySet1(_this.F28(), 2)
				(_nw341).ArraySet1((_this).Fm16(_this.F31(), _1978_v7, _1982_v11, (_this).F38(), globalState), 3)
				(_nw341).ArraySet1(_this.F35(), 4)
				(_nw341).ArraySet1(_this.F31(), 5)
				(_nw341).ArraySet1(_this.F29(), 6)
				(_nw341).ArraySet1(_this.F35(), 7)
				(_nw341).ArraySet1(_this.F31(), 8)
				(_nw341).ArraySet1(!(true), 9)
				(_nw341).ArraySet1(_this.F28(), 10)
				(_nw341).ArraySet1(_this.F28(), 11)
				(_nw341).ArraySet1((_this).Fm7((_1983_v12).Update((_this).F34(), Companion_Default___.Abs((_this).F38())), (_this).F30(), globalState), 12)
				(_nw341).ArraySet1(_this.F35(), 13)
				(_nw341).ArraySet1(_this.F29(), 14)
				(_nw341).ArraySet1(_this.F28(), 15)
				_1984_v13 = _nw341
				var _1985_v14 D3
				_ = _1985_v14
				_1985_v14 = Companion_D3_.Create_DC9_(_this.F28(), false, _1984_v13, _this.F31())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1981_v10), 0))
				_ = _index262
				(_1981_v10).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1985_v14).Dtor_cf21() {
						return (_this).F30()
					}
					return (_this).F30()
				})()), (_index262).Int())
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1981_v10), 0))
				_ = _index263
				(_1981_v10).ArraySet1((_this).F38(), (_index263).Int())
				var _1986_v15 _dafny.Map
				_ = _1986_v15
				_1986_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _this.F35())
				_1986_v15 = (_1986_v15).Update(_this.F36(), _this.F35())
				var _1987_v16 _dafny.Sequence
				_ = _1987_v16
				_1987_v16 = _dafny.SeqOf(((_this).F34()).Minus((_1981_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1981_v10), 0))).Int()).(_dafny.Int)))
				(globalState).F13 = (_1987_v16).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_1987_v16).Cardinality()))).Uint32()).(_dafny.Int)
				var _1988_v18 _dafny.Sequence
				_ = _1988_v18
				_1988_v18 = _dafny.SeqOf(func() _dafny.Set {
					var _coll52 = _dafny.NewBuilder()
					_ = _coll52
					for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(144), _dafny.IntOfInt64(351))); ; {
						_compr_52, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1989_v17 _dafny.Int
						_1989_v17 = interface{}(_compr_52).(_dafny.Int)
						if ((_dafny.IntOfInt64(144)).Cmp(_1989_v17) <= 0) && ((_1989_v17).Cmp(_dafny.IntOfInt64(351)) < 0) {
							_coll52.Add((_1989_v17).Minus((_this).F38()))
						}
					}
					return _coll52.ToSet()
				}())
				var _1990_v21 _dafny.Map
				_ = _1990_v21
				_1990_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), func() _dafny.Set {
					var _coll53 = _dafny.NewBuilder()
					_ = _coll53
					for _iter57 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1976_v5).Cardinality()), _this.F28())).Keys().Elements()); ; {
						_compr_53, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1991_v20 _dafny.Int
						_1991_v20 = interface{}(_compr_53).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1976_v5).Cardinality()), _this.F28())).Contains(_1991_v20) {
							_coll53.Add((_1991_v20).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(454), _dafny.IntOfInt64(198))).Cardinality())))
						}
					}
					return _coll53.ToSet()
				}())
				var _1992_v22 _dafny.Set
				_ = _1992_v22
				_1992_v22 = _dafny.SetOf(_dafny.SetOf((_this).F38()), (func() _dafny.Set {
					if (_1990_v21).Contains((_this).F38()) {
						return (_1990_v21).Get((_this).F38()).(_dafny.Set)
					}
					return (_1988_v18).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_1988_v18).Cardinality()))).Uint32()).(_dafny.Set)
				})(), _1982_v11, _1982_v11, _dafny.SetOf(_dafny.IntOfInt64(-489), (_this).F38()))
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1981_v10), 0))
				_ = _index264
				(_1981_v10).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_1992_v22).IsSubsetOf(func() _dafny.Set {
					var _coll54 = _dafny.NewBuilder()
					_ = _coll54
					for _iter58 := _dafny.Iterate((_1988_v18).Elements()); ; {
						_compr_54, _ok58 := _iter58()
						if !_ok58 {
							break
						}
						var _1993_v19 _dafny.Set
						_1993_v19 = interface{}(_compr_54).(_dafny.Set)
						if _dafny.Companion_Sequence_.Contains(_1988_v18, _1993_v19) {
							_coll54.Add(_1993_v19)
						}
					}
					return _coll54.ToSet()
				}()), (func() bool {
					if _this.F28() {
						return _this.F31()
					}
					return false
				})(), _this.F31())).Cardinality()), (_index264).Int())
			} else {
				var _1994_v23 _dafny.Set
				_ = _1994_v23
				_1994_v23 = _dafny.SetOf((_this).F30())
				(globalState).F1 = (_this).Fm16(((_this).F38()).Cmp((_this).F34()) >= 0, (func() D0 {
					if _this.F29() {
						return Companion_Default___.Fm18(_this.F31(), _this.F29(), globalState)
					}
					return Companion_D0_.Create_DC0_(_this.F35())
				})(), _1994_v23, (_this).F38(), globalState)
				var _1995_v24 _dafny.Sequence
				_ = _1995_v24
				_1995_v24 = _dafny.UnicodeSeqOfUtf8Bytes("srednb")
				var _1996_v25 _dafny.Map
				_ = _1996_v25
				_1996_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_1995_v24, _1995_v24), !((func() bool {
					if _this.F31() {
						return _this.F35()
					}
					return _this.F28()
				})()))
				_1996_v25 = (_1996_v25).Update(_this.F31(), _this.F29())
				(_this).F36_set_(_this.F36())
				var _1997_v26 _dafny.Map
				_ = _1997_v26
				_1997_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if true {
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-92))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg118 _dafny.Int) interface{} {
								return coer118(arg118)
							}
						}(func(_1998_i1 _dafny.Int) _dafny.CodePoint {
							return _this.F36()
						}))).Cardinality())
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("im")).Cardinality())
				})(), !(_this.F35()))
				r0 = !((func() bool {
					if (_1997_v26).Contains((_this).F38()) {
						return (_1997_v26).Get((_this).F38()).(bool)
					}
					return _this.F31()
				})())
				var _1999_v27 _dafny.Array
				_ = _1999_v27
				var _nw342 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw342
				_1999_v27 = _nw342
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1999_v27), 0))
				_ = _index265
				(_1999_v27).ArraySet1(_this.F29(), (_index265).Int())
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1999_v27), 0))
				_ = _index266
				(_1999_v27).ArraySet1(_this.F28(), (_index266).Int())
			}
			var _2000_v28 _dafny.Sequence
			_ = _2000_v28
			_2000_v28 = _dafny.UnicodeSeqOfUtf8Bytes("bd")
			var _2001_v29 _dafny.Array
			_ = _2001_v29
			var _nwElement0_68 _dafny.Sequence = _2000_v28
			_ = _nwElement0_68
			var _nw343 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(2))
			_ = _nw343
			(_nw343).ArraySet1(_nwElement0_68, 0)
			(_nw343).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2000_v28, _2000_v28), 1)
			_2001_v29 = _nw343
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_2001_v29), 0))
			_ = _index267
			(_2001_v29).ArraySet1(_2000_v28, (_index267).Int())
			var _2002_v30 _dafny.Array
			_ = _2002_v30
			var _nw344 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw344
			_2002_v30 = _nw344
			var _2003_v31 _dafny.Array
			_ = _2003_v31
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_51
			var _nw345 _dafny.Array
			_ = _nw345
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw345 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) _dafny.CodePoint = func(_2004_i2 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw345 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw345).ArraySet1CodePoint(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw345).ArraySet1CodePoint(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_2003_v31 = _nw345
			var _2005_v32 *C0
			_ = _2005_v32
			var _nw346 *C0 = New_C0_()
			_ = _nw346
			_nw346.Ctor__(_2003_v31)
			_2005_v32 = _nw346
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_2002_v30), 0))
			_ = _index268
			(_2002_v30).ArraySet1(_2005_v32, (_index268).Int())
			var _2006_v33 _dafny.Sequence
			_ = _2006_v33
			_2006_v33 = _dafny.SeqOf((_this).F38(), (_this).F30())
			var _2007_v34 T2
			_ = _2007_v34
			var _nw347 *C5 = New_C5_()
			_ = _nw347
			_nw347.Ctor__(_this.F28(), (_this).F34(), (_this).F30(), _this.F31(), (_this).F32(), (_this).F33(), (_2006_v33).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_2006_v33).Cardinality()))).Uint32()).(_dafny.Int), _this.F28(), _this.F35(), !(_this.F35()))
			_2007_v34 = _nw347
			var _2008_v35 _dafny.Map
			_ = _2008_v35
			_2008_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2007_v34)
			var _2009_v36 _dafny.Map
			_ = _2009_v36
			_2009_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2000_v28).Cardinality()), Companion_Default___.Fm46((_this).F34(), (_2007_v34).F30(), _this.F35(), globalState))
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_2001_v29), 0))
			_ = _index269
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_2002_v30), 0))
			_ = _index270
			var _rhs285 _dafny.Sequence = (_this).Fm3(_this.F28(), _this.F28(), globalState)
			_ = _rhs285
			var _rhs286 _dafny.Int = ((_2008_v35).Cardinality()).Times(((_this).F38()).Times((func() _dafny.Int {
				if (_2009_v36).Contains((_2006_v33).Select((Companion_Default___.SafeIndex((_2007_v34).F30(), _dafny.IntOfUint32((_2006_v33).Cardinality()))).Uint32()).(_dafny.Int)) {
					return (_2009_v36).Get((_2006_v33).Select((Companion_Default___.SafeIndex((_2007_v34).F30(), _dafny.IntOfUint32((_2006_v33).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Int)
				}
				return (_2007_v34).F30()
			})()))
			_ = _rhs286
			var _rhs287 _dafny.Int = (_2007_v34).F30()
			_ = _rhs287
			var _rhs288 *C0 = _2005_v32
			_ = _rhs288
			var _lhs224 _dafny.Array = _2001_v29
			_ = _lhs224
			var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_2001_v29), 0))
			_ = _lhs225
			var _lhs226 *GlobalState = globalState
			_ = _lhs226
			var _lhs227 *GlobalState = globalState
			_ = _lhs227
			var _lhs228 _dafny.Array = _2002_v30
			_ = _lhs228
			var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_2002_v30), 0))
			_ = _lhs229
			(_lhs224).ArraySet1(_rhs285, (_lhs225).Int())
			_lhs226.F13 = _rhs286
			_lhs227.F13 = _rhs287
			(_lhs228).ArraySet1(_rhs288, (_lhs229).Int())
			r1 = ((_dafny.IntOfInt64(620)).Minus((_this).F38())).Times((_2007_v34).F30())
			var _2010_v37 _dafny.Array
			_ = _2010_v37
			var _nw348 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
			_ = _nw348
			_2010_v37 = _nw348
			(globalState).F21 = _2010_v37
			var _2011_v38 _dafny.Map
			_ = _2011_v38
			_2011_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_this.F31(), _this.F28(), _2007_v34.F29())), _this.F29())
			var _2012_v39 D8
			_ = _2012_v39
			_2012_v39 = Companion_D8_.Create_DC21_(_2011_v38)
			var _2013_v40 D8
			_ = _2013_v40
			_2013_v40 = Companion_D8_.Create_DC25_(_2012_v39)
			var _rhs289 _dafny.Array = _2010_v37
			_ = _rhs289
			var _rhs290 _dafny.Array = _2010_v37
			_ = _rhs290
			var _rhs291 D8 = _2013_v40
			_ = _rhs291
			var _rhs292 bool = _2007_v34.F28()
			_ = _rhs292
			var _lhs230 *C9 = _this
			_ = _lhs230
			_2010_v37 = _rhs289
			_2010_v37 = _rhs290
			_2013_v40 = _rhs291
			_lhs230.F28_set_(_rhs292)
		} else {
			var _2014_v41 _dafny.MultiSet
			_ = _2014_v41
			_2014_v41 = _dafny.MultiSetOf((_this).F38(), (_this).F38())
			var _2015_v42 _dafny.Array
			_ = _2015_v42
			var _nwElement0_69 _dafny.Int = (_2014_v41).Cardinality()
			_ = _nwElement0_69
			var _nw349 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(4))
			_ = _nw349
			(_nw349).ArraySet1(_nwElement0_69, 0)
			(_nw349).ArraySet1((_this).F34(), 1)
			(_nw349).ArraySet1((_this).F34(), 2)
			(_nw349).ArraySet1((_this).F34(), 3)
			_2015_v42 = _nw349
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
			_ = _index271
			(_2015_v42).ArraySet1((_this).F30(), (_index271).Int())
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
			_ = _index272
			(_2015_v42).ArraySet1((_this).F34(), (_index272).Int())
			(globalState).F1 = false
			(globalState).F26 = _this.F35()
			var _2016_v43 *C8
			_ = _2016_v43
			var _nw350 *C8 = New_C8_()
			_ = _nw350
			_nw350.Ctor__(_this.F35(), _this.F36(), (_this).F34(), _this.F28(), (_this).F32(), (_this).F33(), (_this).F30(), _this.F31(), _this.F29(), _this.F31())
			_2016_v43 = _nw350
			var _2017_v44 _dafny.Map
			_ = _2017_v44
			_2017_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2016_v43, (_this).F30())
			var _2018_v45 _dafny.Set
			_ = _2018_v45
			_2018_v45 = _dafny.SetOf((_dafny.Zero).Minus((_this).F34()), (_this).F30(), Companion_Default___.Fm24(_this.F35(), (_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int), globalState), (_this).F38(), (_2017_v44).Cardinality())
			if ((_dafny.Zero).Minus(((_this).F34()).Plus((_this).F34()))).Cmp((_2018_v45).Cardinality()) >= 0 {
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
				_ = _index273
				var _rhs293 bool = ((_this).F38()).Cmp((_this).F34()) == 0
				_ = _rhs293
				var _rhs294 _dafny.Int = (_this).F38()
				_ = _rhs294
				var _lhs231 *GlobalState = globalState
				_ = _lhs231
				var _lhs232 _dafny.Array = _2015_v42
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
				_ = _lhs233
				_lhs231.F1 = _rhs293
				(_lhs232).ArraySet1(_rhs294, (_lhs233).Int())
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
				_ = _index274
				(_2015_v42).ArraySet1((_this).F34(), (_index274).Int())
				var _2019_v46 _dafny.Array
				_ = _2019_v46
				var _len0_52 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_52
				var _nw351 _dafny.Array
				_ = _nw351
				if _len0_52.Cmp(_dafny.Zero) == 0 {
					_nw351 = _dafny.NewArray(_len0_52)
				} else {
					var _init52 func(_dafny.Int) bool = func(_2020_i3 _dafny.Int) bool {
						return _this.F28()
					}
					_ = _init52
					var _element0_52 = _init52(_dafny.Zero)
					_ = _element0_52
					_nw351 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
					(_nw351).ArraySet1(_element0_52, 0)
					var _nativeLen0_52 = (_len0_52).Int()
					_ = _nativeLen0_52
					for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
						(_nw351).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
					}
				}
				_2019_v46 = _nw351
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2019_v46), 0))
				_ = _index275
				(_2019_v46).ArraySet1(_this.F35(), (_index275).Int())
				var _2021_v47 _dafny.Set
				_ = _2021_v47
				_2021_v47 = _dafny.SetOf(_dafny.SetOf((_this).F30()))
				var _2022_v48 T0
				_ = _2022_v48
				var _nw352 *C7 = New_C7_()
				_ = _nw352
				_nw352.Ctor__((_this).F34(), _2014_v41, (_this).F38(), _this.F35(), _this.F35(), _2016_v43.F40)
				_2022_v48 = _nw352
				var _2023_v49 _dafny.Map
				_ = _2023_v49
				_2023_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if _2022_v48.F29() {
						return (_this).F32()
					}
					return _dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), _2016_v43.F40)
				})()).Cardinality()), _2022_v48)
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2019_v46), 0))
				_ = _index276
				var _rhs295 bool = _this.F29()
				_ = _rhs295
				var _rhs296 _dafny.Set = (_2021_v47).Intersection(_2021_v47)
				_ = _rhs296
				var _rhs297 T0 = (func() T0 {
					if (_2023_v49).Contains(Companion_Default___.SafeDivisionInt((_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("okb")).Cardinality()))) {
						return (_2023_v49).Get(Companion_Default___.SafeDivisionInt((_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("okb")).Cardinality()))).(T0)
					}
					return _2022_v48
				})()
				_ = _rhs297
				var _lhs234 _dafny.Array = _2019_v46
				_ = _lhs234
				var _lhs235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2019_v46), 0))
				_ = _lhs235
				(_lhs234).ArraySet1(_rhs295, (_lhs235).Int())
				_2021_v47 = _rhs296
				_2022_v48 = _rhs297
				var _2024_v50 *C1
				_ = _2024_v50
				var _nw353 *C1 = New_C1_()
				_ = _nw353
				_nw353.Ctor__(_this.F31(), true, _this.F29())
				_2024_v50 = _nw353
				var _2025_v51 _dafny.Map
				_ = _2025_v51
				_2025_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2024_v50, Companion_Default___.Fm39(_2016_v43.F40, _2014_v41, ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), globalState)))
				var _2026_v52 _dafny.Map
				_ = _2026_v52
				_2026_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2024_v50, (_this).F38())
				_2025_v51 = (_2025_v51).Update(_this.F36(), _2026_v52)
				var _2027_v53 _dafny.Sequence
				_ = _2027_v53
				_2027_v53 = _dafny.UnicodeSeqOfUtf8Bytes("mhctyq")
				var _2028_v54 _dafny.Array
				_ = _2028_v54
				var _nwElement0_70 _dafny.CodePoint = _this.F36()
				_ = _nwElement0_70
				var _nw354 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(15))
				_ = _nw354
				(_nw354).ArraySet1CodePoint(_nwElement0_70, 0)
				(_nw354).ArraySet1CodePoint(_this.F36(), 1)
				(_nw354).ArraySet1CodePoint(_this.F36(), 2)
				(_nw354).ArraySet1CodePoint(_dafny.CodePoint('h'), 3)
				(_nw354).ArraySet1CodePoint(_this.F36(), 4)
				(_nw354).ArraySet1CodePoint(Companion_Default___.Fm34((_this).F38(), (_this).F38(), _dafny.IntOfUint32((_2027_v53).Cardinality()), _this.F28(), globalState), 5)
				(_nw354).ArraySet1CodePoint(_this.F36(), 6)
				(_nw354).ArraySet1CodePoint(Companion_Default___.Fm41(_this.F29(), _this.F36(), _this.F29(), (_this).F38(), globalState), 7)
				(_nw354).ArraySet1CodePoint((_2027_v53).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_2027_v53).Cardinality()))).Uint32()).(_dafny.CodePoint), 8)
				(_nw354).ArraySet1CodePoint(_this.F36(), 9)
				(_nw354).ArraySet1CodePoint(_dafny.CodePoint('v'), 10)
				(_nw354).ArraySet1CodePoint(_dafny.CodePoint('a'), 11)
				(_nw354).ArraySet1CodePoint((_2027_v53).Select((Companion_Default___.SafeIndex((_2014_v41).Cardinality(), _dafny.IntOfUint32((_2027_v53).Cardinality()))).Uint32()).(_dafny.CodePoint), 12)
				(_nw354).ArraySet1CodePoint(_this.F36(), 13)
				(_nw354).ArraySet1CodePoint(_this.F36(), 14)
				_2028_v54 = _nw354
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_2028_v54), 0))
				_ = _index277
				(_2028_v54).ArraySet1CodePoint(_this.F36(), (_index277).Int())
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_2028_v54), 0))
				_ = _index278
				var _rhs298 _dafny.CodePoint = _this.F36()
				_ = _rhs298
				var _rhs299 _dafny.CodePoint = _this.F36()
				_ = _rhs299
				var _rhs300 bool = !(_this.F28()) || ((_2019_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_2019_v46), 0))).Int()).(bool))
				_ = _rhs300
				var _lhs236 *C9 = _this
				_ = _lhs236
				var _lhs237 _dafny.Array = _2028_v54
				_ = _lhs237
				var _lhs238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_2028_v54), 0))
				_ = _lhs238
				var _lhs239 *GlobalState = globalState
				_ = _lhs239
				_lhs236.F36_set_(_rhs298)
				(_lhs237).ArraySet1CodePoint(_rhs299, (_lhs238).Int())
				_lhs239.F26 = _rhs300
			} else {
				var _2029_v55 _dafny.Map
				_ = _2029_v55
				_2029_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F35())
				var _2030_v56 _dafny.Set
				_ = _2030_v56
				_2030_v56 = _dafny.SetOf(_this.F36(), Companion_Default___.Fm41(_this.F31(), _this.F36(), _this.F35(), (_this).F30(), globalState))
				var _2031_v57 D0
				_ = _2031_v57
				_2031_v57 = Companion_D0_.Create_DC0_(_this.F28())
				var _2032_v58 _dafny.Map
				_ = _2032_v58
				_2032_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int), _this.F31())
				var _2033_v59 _dafny.Map
				_ = _2033_v59
				_2033_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int), Companion_Default___.Fm1(_2016_v43.F40, (_2032_v58).Cardinality(), globalState))
				var _2034_v60 _dafny.Array
				_ = _2034_v60
				var _nwElement0_71 bool = _2016_v43.F40
				_ = _nwElement0_71
				var _nw355 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(26))
				_ = _nw355
				(_nw355).ArraySet1(_nwElement0_71, 0)
				(_nw355).ArraySet1(_this.F28(), 1)
				(_nw355).ArraySet1(!(_this.F35()), 2)
				(_nw355).ArraySet1(_this.F35(), 3)
				(_nw355).ArraySet1(_this.F35(), 4)
				(_nw355).ArraySet1(_this.F31(), 5)
				(_nw355).ArraySet1((func() bool {
					if (_2029_v55).Contains(_this.F35()) {
						return (_2029_v55).Get(_this.F35()).(bool)
					}
					return true
				})(), 6)
				(_nw355).ArraySet1(_2016_v43.F40, 7)
				(_nw355).ArraySet1(!(_2016_v43.F40), 8)
				(_nw355).ArraySet1((_2014_v41).Contains((_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int)), 9)
				(_nw355).ArraySet1((_2030_v56).Contains(_dafny.CodePoint('s')), 10)
				(_nw355).ArraySet1(_2016_v43.F40, 11)
				(_nw355).ArraySet1(_this.F35(), 12)
				(_nw355).ArraySet1(_this.F31(), 13)
				(_nw355).ArraySet1(_this.F29(), 14)
				(_nw355).ArraySet1((_this).Fm16(!(_this.F28()), _2031_v57, (_this).Fm4(_this.F31(), _dafny.SetOf((_this).F30()), (_this).F30(), _this.F29(), globalState), _dafny.IntOfUint32(((_this).F32()).Cardinality()), globalState), 15)
				(_nw355).ArraySet1(_2016_v43.F40, 16)
				(_nw355).ArraySet1(_this.F35(), 17)
				(_nw355).ArraySet1(_this.F28(), 18)
				(_nw355).ArraySet1(((_this).F34()).Cmp(_dafny.IntOfInt64(848)) == 0, 19)
				(_nw355).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg119 _dafny.Int) interface{} {
						return coer119(arg119)
					}
				}(func(_2035_i4 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				})), _dafny.UnicodeSeqOfUtf8Bytes("vdyxsb")), 20)
				(_nw355).ArraySet1(_this.F29(), 21)
				(_nw355).ArraySet1((_this.F31()) || (_this.F35()), 22)
				(_nw355).ArraySet1((_2032_v58).Contains((_this).F30()), 23)
				(_nw355).ArraySet1((false) || (!(_this.F35())), 24)
				(_nw355).ArraySet1(!(_this.F31()) || (_this.F29()), 25)
				_2034_v60 = _nw355
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_2034_v60), 0))
				_ = _index279
				(_2034_v60).ArraySet1((func() bool {
					if true {
						return _this.F35()
					}
					return true
				})(), (_index279).Int())
				var _2036_v61 _dafny.MultiSet
				_ = _2036_v61
				_2036_v61 = _dafny.MultiSetOf(_this.F31())
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_2034_v60), 0))
				_ = _index280
				(_2034_v60).ArraySet1((func() bool {
					if (_2036_v61).IsProperSubsetOf(_2036_v61) {
						return _this.F31()
					}
					return _this.F29()
				})(), (_index280).Int())
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
				_ = _index281
				(_2015_v42).ArraySet1((_this).F34(), (_index281).Int())
				(_this).F36_set_(_this.F36())
				(globalState).F12 = ((_2029_v55).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), _2016_v43.F40))).Equals((_2029_v55).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2016_v43.F40)))
				(_this).F36_set_(_this.F36())
			}
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
			_ = _index282
			var _rhs301 bool = !((!(_this.F35())) || (_this.F28()))
			_ = _rhs301
			var _rhs302 _dafny.Int = (_2015_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))).Int()).(_dafny.Int)
			_ = _rhs302
			var _rhs303 bool = _2016_v43.F40
			_ = _rhs303
			var _lhs240 *C8 = _2016_v43
			_ = _lhs240
			var _lhs241 _dafny.Array = _2015_v42
			_ = _lhs241
			var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_2015_v42), 0))
			_ = _lhs242
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			_lhs240.F40 = _rhs301
			(_lhs241).ArraySet1(_rhs302, (_lhs242).Int())
			_lhs243.F12 = _rhs303
		}
		var _2037_v62 _dafny.Array
		_ = _2037_v62
		var _len0_53 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_53
		var _nw356 _dafny.Array
		_ = _nw356
		if _len0_53.Cmp(_dafny.Zero) == 0 {
			_nw356 = _dafny.NewArray(_len0_53)
		} else {
			var _init53 func(_dafny.Int) bool = func(_2038_i5 _dafny.Int) bool {
				return (func() bool {
					if _this.F35() {
						return false
					}
					return _this.F28()
				})()
			}
			_ = _init53
			var _element0_53 = _init53(_dafny.Zero)
			_ = _element0_53
			_nw356 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
			(_nw356).ArraySet1(_element0_53, 0)
			var _nativeLen0_53 = (_len0_53).Int()
			_ = _nativeLen0_53
			for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
				(_nw356).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
			}
		}
		_2037_v62 = _nw356
		var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_2037_v62), 0))
		_ = _index283
		(_2037_v62).ArraySet1(_this.F28(), (_index283).Int())
		var _2039_v63 _dafny.Sequence
		_ = _2039_v63
		_2039_v63 = _dafny.UnicodeSeqOfUtf8Bytes("gxqix")
		var _2040_v64 _dafny.Sequence
		_ = _2040_v64
		_2040_v64 = _dafny.SeqOf(_2039_v63, _2039_v63)
		var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_2037_v62), 0))
		_ = _index284
		(_2037_v62).ArraySet1(_dafny.Companion_Sequence_.Equal(_2040_v64, _2040_v64), (_index284).Int())
		var _2041_v65 _dafny.Array
		_ = _2041_v65
		var _nw357 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw357
		_2041_v65 = _nw357
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2041_v65), 0))); ; {
			_guard_loop_4, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _2042_i6 _dafny.Int
			_2042_i6 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_2042_i6).Sign() != -1) && ((_2042_i6).Cmp(_dafny.ArrayLen((_2041_v65), 0)) < 0)) {
				(_2041_v65).ArraySet1((_2042_i6).Minus((_this).F30()), (_2042_i6).Int())
			}
		}
		r1 = (_1970_v0).Cardinality()
		r1 = _dafny.IntOfInt64(-371)
		var _2043_v66 _dafny.Map
		_ = _2043_v66
		_2043_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))
		if (_2043_v66).Equals((_2043_v66).Merge(_2043_v66)) {
			r0 = _this.F28()
			var _2044_v67 _dafny.Array
			_ = _2044_v67
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_54
			var _nw358 _dafny.Array
			_ = _nw358
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw358 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.CodePoint = func(_2045_i7 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw358 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw358).ArraySet1CodePoint(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw358).ArraySet1CodePoint(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_2044_v67 = _nw358
			var _2046_v68 *C0
			_ = _2046_v68
			var _nw359 *C0 = New_C0_()
			_ = _nw359
			_nw359.Ctor__(_2044_v67)
			_2046_v68 = _nw359
			var _2047_v69 D1
			_ = _2047_v69
			_2047_v69 = Companion_D1_.Create_DC3_(_this.F29(), _dafny.IntOfUint32((_2039_v63).Cardinality()), _2046_v68, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xevgdv")).Cardinality())))
			var _2048_v70 *C4
			_ = _2048_v70
			var _nw360 *C4 = New_C4_()
			_ = _nw360
			_nw360.Ctor__((_this).F34(), (_this).Fm6(globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F35()), (_this).F32()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F32()), _dafny.SeqOf((_this).F32(), (_this).F32())), Companion_Default___.SafeModuloInt((_this).F34(), (_this).F34()), (_2047_v69).Dtor_cf3(), _this.F35(), _this.F31())
			_2048_v70 = _nw360
			var _rhs304 *C4 = _2048_v70
			_ = _rhs304
			var _rhs305 bool = _this.F31()
			_ = _rhs305
			var _rhs306 _dafny.Array = _2041_v65
			_ = _rhs306
			var _lhs244 *C9 = _this
			_ = _lhs244
			_2048_v70 = _rhs304
			_lhs244.F31_set_(_rhs305)
			_2041_v65 = _rhs306
			var _2049_v71 _dafny.Map
			_ = _2049_v71
			_2049_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), _2041_v65)
			_2041_v65 = (func() _dafny.Array {
				if (_2049_v71).Contains((_this).F38()) {
					return (_2049_v71).Get((_this).F38()).(_dafny.Array)
				}
				return _2041_v65
			})()
			(_this).F28_set_((_1970_v0).IsProperSubsetOf(_1970_v0))
			var _2050_v72 _dafny.Array
			_ = _2050_v72
			var _nw361 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw361
			_2050_v72 = _nw361
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_2050_v72), 0))
			_ = _index285
			(_2050_v72).ArraySet1(_2037_v62, (_index285).Int())
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_2050_v72), 0))
			_ = _index286
			(_2050_v72).ArraySet1(_2037_v62, (_index286).Int())
		} else {
			var _2051_v73 _dafny.Map
			_ = _2051_v73
			_2051_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2037_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_2037_v62), 0))).Int()).(bool), (_dafny.Zero).Minus((_this).F34()))
			var _2052_v74 _dafny.Sequence
			_ = _2052_v74
			_2052_v74 = _dafny.SeqOf((_2051_v73).Cardinality(), Companion_Default___.SafeDivisionInt((_this).F34(), (_dafny.MultiSetOf(_2041_v65)).Cardinality()), (_this).F38())
			var _2053_v75 _dafny.Sequence
			_ = _2053_v75
			_2053_v75 = _dafny.SeqOf(_2052_v74, _2052_v74)
			_2052_v74 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2052_v74, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(832))).Uint32(), func(coer120 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg120 _dafny.Int) interface{} {
					return coer120(arg120)
				}
			}(func(_2054_i8 _dafny.Int) _dafny.Int {
				return (_this).F38()
			}))), _dafny.Companion_Sequence_.Concatenate(_2052_v74, (_2053_v75).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_2053_v75).Cardinality()))).Uint32()).(_dafny.Sequence)))
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_2041_v65), 0))
			_ = _index287
			(_2041_v65).ArraySet1((_1970_v0).Cardinality(), (_index287).Int())
			var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_2041_v65), 0))
			_ = _index288
			(_2041_v65).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-825), (_this).F34()), (_index288).Int())
			var _2055_v76 _dafny.Map
			_ = _2055_v76
			_2055_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2052_v74).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_2052_v74).Cardinality()))).Uint32()).(_dafny.Int), _this.F29())
			var _2056_v77 _dafny.Map
			_ = _2056_v77
			_2056_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2037_v62, (_2055_v76).Cardinality())
			var _2057_v78 _dafny.Map
			_ = _2057_v78
			_2057_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2041_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_2041_v65), 0))).Int()).(_dafny.Int), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(836))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg121 _dafny.Int) interface{} {
					return coer121(arg121)
				}
			}(func(_2058_i9 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			})))
			var _2059_v79 _dafny.Map
			_ = _2059_v79
			_2059_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _dafny.CodePoint('h'))
			var _2060_v80 _dafny.Array
			_ = _2060_v80
			var _nw362 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw362
			_2060_v80 = _nw362
			var _2061_v81 *C5
			_ = _2061_v81
			var _nw363 *C5 = New_C5_()
			_ = _nw363
			_nw363.Ctor__(((_this).F34()).Cmp(_dafny.IntOfInt64(793)) == 0, (_2056_v77).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_2057_v78).Contains((_this).F38()) {
					return (_2057_v78).Get((_this).F38()).(_dafny.Sequence)
				}
				return _2039_v63
			})()).Cardinality()), (_2037_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_2037_v62), 0))).Int()).(bool), (_this).F32(), _dafny.SeqOf(_dafny.SeqOf(_this.F31(), _this.F29()), _dafny.Companion_Sequence_.Update((_this).F32(), (Companion_Default___.SafeIndex((_2059_v79).Cardinality(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32(), _this.F28()), (_this).F32()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2052_v74, (Companion_Default___.SafeIndex((_dafny.SetOf(_2060_v80)).Cardinality(), _dafny.IntOfUint32((_2052_v74).Cardinality()))).Uint32(), (_2052_v74).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_2052_v74).Cardinality()))).Uint32()).(_dafny.Int))).Cardinality()), (_dafny.IntOfUint32(((_this).F32()).Cardinality())).Cmp((_this).F30()) < 0, _this.F28(), ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool))
			_2061_v81 = _nw363
			var _2062_v82 _dafny.MultiSet
			_ = _2062_v82
			_2062_v82 = _dafny.MultiSetOf((_this).F30(), _dafny.IntOfUint32((_2039_v63).Cardinality()))
			var _2063_v83 _dafny.MultiSet
			_ = _2063_v83
			_2063_v83 = _dafny.MultiSetOf(_dafny.IntOfInt64(290), (_2062_v82).Cardinality())
			(globalState).F12 = (_this).Fm7(_2062_v82, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2052_v74, _2052_v74)).Cardinality())), globalState)
			var _2064_v84 D8
			_ = _2064_v84
			_2064_v84 = Companion_D8_.Create_DC22_((_this).F34(), (_this).F30())
			var _2065_v85 D2
			_ = _2065_v85
			_2065_v85 = Companion_D2_.Create_DC6_((_2064_v84).Dtor_cf41(), _this.F28(), (_2061_v81).F48(), (_2061_v81).F47(), _dafny.IntOfInt64(-48))
			var _pat_let_tv18 = _2055_v76
			_ = _pat_let_tv18
			var _pat_let_tv19 = _2055_v76
			_ = _pat_let_tv19
			_2065_v85 = func(_pat_let25_0 D2) D2 {
				return func(_2066_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let26_0 _dafny.Int) D2 {
						return func(_2067_dt__update_hcf9_h0 _dafny.Int) D2 {
							return func(_pat_let27_0 bool) D2 {
								return func(_2068_dt__update_hcf12_h0 bool) D2 {
									return func(_pat_let28_0 _dafny.Int) D2 {
										return func(_2069_dt__update_hcf13_h0 _dafny.Int) D2 {
											return Companion_D2_.Create_DC6_(_2067_dt__update_hcf9_h0, (_2066_dt__update__tmp_h0).Dtor_cf10(), (_2066_dt__update__tmp_h0).Dtor_cf11(), _2068_dt__update_hcf12_h0, _2069_dt__update_hcf13_h0)
										}(_pat_let28_0)
									}(((_pat_let_tv18).Merge(_pat_let_tv19)).Cardinality())
								}(_pat_let27_0)
							}(!(_this.F29()))
						}(_pat_let26_0)
					}((_this).F38())
				}(_pat_let25_0)
			}(_2065_v85)
		}
		r0 = _this.F31()
		r1 = (_this).F30()
		var _2070_v86 _dafny.MultiSet
		_ = _2070_v86
		_2070_v86 = _dafny.MultiSetOf((_this).F34(), (_dafny.Zero).Minus((_this).F38()))
		var _2071_v87 D14
		_ = _2071_v87
		_2071_v87 = Companion_D14_.Create_DC37_(Companion_Default___.Fm46((_this).F34(), (_this).F38(), _this.F29(), globalState), (_this).Fm7(_2070_v86, (_this).F30(), globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mfybmsjun")).Cardinality()))
		var _pat_let_tv20 = _2039_v63
		_ = _pat_let_tv20
		var _pat_let_tv21 = _2039_v63
		_ = _pat_let_tv21
		r2 = func(_source33 D14) _dafny.Sequence {
			if _source33.Is_DC37() {
				var _2072___mcc_h0 _dafny.Int = _source33.Get_().(D14_DC37).Cf68
				_ = _2072___mcc_h0
				var _2073___mcc_h1 bool = _source33.Get_().(D14_DC37).Cf69
				_ = _2073___mcc_h1
				var _2074___mcc_h2 _dafny.Int = _source33.Get_().(D14_DC37).Cf70
				_ = _2074___mcc_h2
				var _2075_cf70 _dafny.Int = _2074___mcc_h2
				_ = _2075_cf70
				var _2076_cf69 bool = _2073___mcc_h1
				_ = _2076_cf69
				var _2077_cf68 _dafny.Int = _2072___mcc_h0
				_ = _2077_cf68
				return _pat_let_tv20
			} else {
				var _2078___mcc_h3 _dafny.Sequence = _source33.Get_().(D14_DC36).Cf67
				_ = _2078___mcc_h3
				var _2079_cf67 _dafny.Sequence = _2078___mcc_h3
				_ = _2079_cf67
				return _pat_let_tv21
			}
		}(_2071_v87)
		return r0, r1, r2
	}
}
func (_this *C9) M4(globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		(_this).F36_set_(_this.F36())
		(globalState).F13 = (_dafny.Zero).Minus((_this).F38())
		(_this).F28_set_(_this.F29())
		var _2080_v0 D0
		_ = _2080_v0
		_2080_v0 = Companion_D0_.Create_DC0_(_this.F35())
		if (_2080_v0).Dtor_cf0() {
			var _2081_v1 _dafny.Sequence
			_ = _2081_v1
			_2081_v1 = _dafny.SeqOf((_this).F34(), (_this).F34())
			var _2082_v2 _dafny.Set
			_ = _2082_v2
			_2082_v2 = _dafny.SetOf(_2081_v1)
			var _2083_v3 D16
			_ = _2083_v3
			_2083_v3 = Companion_D16_.Create_DC39_(_2082_v2)
			var _2084_v4 D16
			_ = _2084_v4
			_2084_v4 = Companion_D16_.Create_DC41_(_2083_v3)
			var _2085_v5 D16
			_ = _2085_v5
			_2085_v5 = Companion_D16_.Create_DC41_(Companion_D16_.Create_DC41_(_2084_v4))
			var _2086_v6 _dafny.Array
			_ = _2086_v6
			var _nw364 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw364
			_2086_v6 = _nw364
			var _2087_v7 D20
			_ = _2087_v7
			_2087_v7 = Companion_D20_.Create_DC47_(_2086_v6)
			var _pat_let_tv22 = _2086_v6
			_ = _pat_let_tv22
			var _rhs307 D16 = _2085_v5
			_ = _rhs307
			var _rhs308 _dafny.Int = (_this).F38()
			_ = _rhs308
			var _rhs309 bool = ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)
			_ = _rhs309
			var _rhs310 _dafny.Array = (func(_pat_let29_0 D20) D20 {
				return func(_2088_dt__update__tmp_h0 D20) D20 {
					return func(_pat_let30_0 _dafny.Array) D20 {
						return func(_2089_dt__update_hcf81_h0 _dafny.Array) D20 {
							return Companion_D20_.Create_DC47_(_2089_dt__update_hcf81_h0)
						}(_pat_let30_0)
					}(_pat_let_tv22)
				}(_pat_let29_0)
			}(_2087_v7)).Dtor_cf81()
			_ = _rhs310
			var _rhs311 _dafny.Int = (_this).F30()
			_ = _rhs311
			var _lhs245 *GlobalState = globalState
			_ = _lhs245
			var _lhs246 *GlobalState = globalState
			_ = _lhs246
			var _lhs247 *GlobalState = globalState
			_ = _lhs247
			_2085_v5 = _rhs307
			_lhs245.F13 = _rhs308
			_lhs246.F1 = _rhs309
			_2086_v6 = _rhs310
			_lhs247.F13 = _rhs311
			var _2090_v8 _dafny.MultiSet
			_ = _2090_v8
			_2090_v8 = _dafny.MultiSetOf(false, _this.F28())
			if (_dafny.MultiSetOf(_this.F31(), _this.F35(), _this.F35())).IsProperSubsetOf(_2090_v8) {
				(globalState).F13 = (_this).F34()
				var _2091_v9 _dafny.Map
				_ = _2091_v9
				_2091_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _dafny.IntOfUint32((((_this).F33()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F30()), _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				var _2092_v10 _dafny.Sequence
				_ = _2092_v10
				_2092_v10 = _dafny.SeqOf(_2091_v9)
				(_this).F28_set_(((_this).F38()).Cmp(_dafny.IntOfUint32((_2092_v10).Cardinality())) >= 0)
				var _2093_v11 D20
				_ = _2093_v11
				_2093_v11 = Companion_D20_.Create_DC49_(_this.F29(), _this.F36())
				var _2094_v12 *C1
				_ = _2094_v12
				var _nw365 *C1 = New_C1_()
				_ = _nw365
				_nw365.Ctor__(_this.F29(), _dafny.Companion_Sequence_.Contains((_this).F32(), _this.F28()), !((_2093_v11).Dtor_cf85()))
				_2094_v12 = _nw365
				(globalState).F1 = _this.F35()
				var _2095_v13 _dafny.Array
				_ = _2095_v13
				var _nw366 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw366
				_2095_v13 = _nw366
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_2095_v13), 0))
				_ = _index289
				(_2095_v13).ArraySet1(true, (_index289).Int())
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_2095_v13), 0))
				_ = _index290
				(_2095_v13).ArraySet1(_this.F35(), (_index290).Int())
				var _2096_v14 _dafny.Sequence
				_ = _2096_v14
				_2096_v14 = _dafny.UnicodeSeqOfUtf8Bytes("wc")
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_2095_v13), 0))
				_ = _index291
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_2095_v13), 0))
				_ = _index292
				var _rhs312 bool = ((_this).F34()).Cmp((_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_2081_v1).Cardinality()))).Minus((_dafny.Zero).Minus((_this).F38())))) <= 0
				_ = _rhs312
				var _rhs313 bool = _this.F31()
				_ = _rhs313
				var _rhs314 _dafny.Int = (Companion_Default___.SafeDivisionInt((_this).F38(), _dafny.IntOfUint32((_2081_v1).Cardinality()))).Times((_this).F38())
				_ = _rhs314
				var _rhs315 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("e")
				_ = _rhs315
				var _rhs316 _dafny.Int = (_this).F30()
				_ = _rhs316
				var _lhs248 _dafny.Array = _2095_v13
				_ = _lhs248
				var _lhs249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_2095_v13), 0))
				_ = _lhs249
				var _lhs250 _dafny.Array = _2095_v13
				_ = _lhs250
				var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_2095_v13), 0))
				_ = _lhs251
				var _lhs252 *GlobalState = globalState
				_ = _lhs252
				var _lhs253 *GlobalState = globalState
				_ = _lhs253
				(_lhs248).ArraySet1(_rhs312, (_lhs249).Int())
				(_lhs250).ArraySet1(_rhs313, (_lhs251).Int())
				_lhs252.F13 = _rhs314
				_2096_v14 = _rhs315
				_lhs253.F13 = _rhs316
			} else {
				var _2097_v15 _dafny.Array
				_ = _2097_v15
				var _nw367 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
				_ = _nw367
				_2097_v15 = _nw367
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_2097_v15), 0))
				_ = _index293
				(_2097_v15).ArraySet1(_2090_v8, (_index293).Int())
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_2097_v15), 0))
				_ = _index294
				(_2097_v15).ArraySet1(_2090_v8, (_index294).Int())
				(_this).F28_set_(_this.F35())
				var _2098_v16 D8
				_ = _2098_v16
				_2098_v16 = Companion_D8_.Create_DC24_(_this.F31(), (_this).F34(), (_this).F32(), _this.F35(), _2081_v1)
				var _2099_v17 D8
				_ = _2099_v17
				_2099_v17 = Companion_D8_.Create_DC25_(_2098_v16)
				var _2100_v18 _dafny.Sequence
				_ = _2100_v18
				_2100_v18 = _dafny.SeqOf(_2099_v17, _2099_v17)
				var _2101_v19 _dafny.MultiSet
				_ = _2101_v19
				_2101_v19 = _dafny.MultiSetOf(_2100_v18, _2100_v18)
				(globalState).F26 = ((_2101_v19).Difference(_2101_v19)).IsProperSubsetOf(_2101_v19)
				var _2102_v20 _dafny.Array
				_ = _2102_v20
				var _nw368 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
				_ = _nw368
				_2102_v20 = _nw368
				var _2103_v21 _dafny.MultiSet
				_ = _2103_v21
				_2103_v21 = _dafny.MultiSetOf(_2102_v20)
				(globalState).F26 = !((_2103_v21).IsProperSubsetOf(_2103_v21))
				var _2104_v22 _dafny.MultiSet
				_ = _2104_v22
				_2104_v22 = _dafny.MultiSetOf((_this).F34(), (_this).F30())
				var _2105_v23 _dafny.Map
				_ = _2105_v23
				_2105_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(((_2104_v22).Update(_dafny.IntOfInt64(-797), Companion_Default___.Abs((_this).F38()))).Cardinality()), (_this).F38()))
				var _2106_v24 _dafny.Set
				_ = _2106_v24
				_2106_v24 = _dafny.SetOf(_this.F31())
				(globalState).F13 = (func() _dafny.Int {
					if (_2105_v23).Contains(((_2106_v24).Intersection(_dafny.SetOf(_this.F35()))).Cardinality()) {
						return (_2105_v23).Get(((_2106_v24).Intersection(_dafny.SetOf(_this.F35()))).Cardinality()).(_dafny.Int)
					}
					return ((_this).F34()).Times((_this).F38())
				})()
			}
			var _2107_v25 _dafny.Array
			_ = _2107_v25
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_55
			var _nw369 _dafny.Array
			_ = _nw369
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw369 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) _dafny.Map = (func(_2108_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_2109_i0 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _dafny.IntOfInt64(238))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2108_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-294), _dafny.IntOfUint32((_2108_v1).Cardinality()))).Uint32()).(_dafny.Int), (_this).F30()))
					}
				})(_2081_v1)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw369 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw369).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw369).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_2107_v25 = _nw369
			var _2110_v26 _dafny.Map
			_ = _2110_v26
			_2110_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(713), (_this).F34())
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_2107_v25), 0))
			_ = _index295
			(_2107_v25).ArraySet1((_2110_v26).Merge(_2110_v26), (_index295).Int())
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_2107_v25), 0))
			_ = _index296
			(_2107_v25).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F32()).Cardinality()), (_this).F38()), (_index296).Int())
			var _2111_v27 _dafny.MultiSet
			_ = _2111_v27
			_2111_v27 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F30()))
			var _2112_v28 D14
			_ = _2112_v28
			_2112_v28 = Companion_D14_.Create_DC37_((func() _dafny.Int {
				if (_2111_v27).Contains((_this).F38()) {
					return (_2111_v27).Multiplicity((_this).F38())
				}
				return (_this).F34()
			})(), _this.F35(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(625), (_this).F38()))
			_2112_v28 = Companion_Default___.Fm58(_this.F28(), ((_this).F33()).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence), globalState)
			var _2113_v29 _dafny.Set
			_ = _2113_v29
			_2113_v29 = _dafny.SetOf((_this).F30())
			var _2114_v30 _dafny.Int
			_ = _2114_v30
			var _out69 _dafny.Int
			_ = _out69
			_out69 = Companion_Default___.M0(_2081_v1, _this.F28(), (_this).F34(), (_2113_v29).Contains((_this).F38()), globalState)
			_2114_v30 = _out69
		} else {
			var _nw370 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw370
			(globalState).F21 = _nw370
			if _this.F35() {
				var _2115_v31 _dafny.Sequence
				_ = _2115_v31
				_2115_v31 = _dafny.SeqOf((_this).F34())
				var _2116_v32 _dafny.Int
				_ = _2116_v32
				var _out70 _dafny.Int
				_ = _out70
				_out70 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_2115_v31, _2115_v31), _this.F31(), (_this).F38(), _this.F35(), globalState)
				_2116_v32 = _out70
				(globalState).F26 = false
				(_this).F35_set_(!(_this.F31()))
				var _2117_v33 _dafny.Map
				_ = _2117_v33
				_2117_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F38())
				var _2118_v35 _dafny.Map
				_ = _2118_v35
				_2118_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), func() _dafny.Set {
					var _coll55 = _dafny.NewBuilder()
					_ = _coll55
					for _iter60 := _dafny.Iterate((_2117_v33).Keys().Elements()); ; {
						_compr_55, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _2119_v34 _dafny.Int
						_2119_v34 = interface{}(_compr_55).(_dafny.Int)
						if (_2117_v33).Contains(_2119_v34) {
							_coll55.Add(Companion_Default___.SafeDivisionInt(_2119_v34, _dafny.IntOfUint32((_dafny.SeqOf(true, !(true))).Cardinality())))
						}
					}
					return _coll55.ToSet()
				}())
				var _2120_v36 D2
				_ = _2120_v36
				_2120_v36 = Companion_D2_.Create_DC6_((_dafny.Zero).Minus(((func() _dafny.Map {
					if _this.F31() {
						return _2118_v35
					}
					return _2118_v35
				})()).Cardinality()), ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), (_this).F34(), _this.F31(), Companion_Default___.SafeDivisionInt(_2116_v32, _dafny.IntOfInt64(423)))
				var _2121_v37 _dafny.Set
				_ = _2121_v37
				_2121_v37 = _dafny.SetOf(_this.F31())
				var _2122_v38 _dafny.Sequence
				_ = _2122_v38
				_2122_v38 = _dafny.SeqOf(_2121_v37, _dafny.SetOf(_this.F28()), _2121_v37, Companion_Default___.Fm45((_this).F38(), _dafny.MultiSetFromSeq((_this).F32()), _this.F28(), (_dafny.Zero).Minus(_2116_v32), globalState), _2121_v37)
				var _2123_v39 _dafny.Set
				_ = _2123_v39
				_2123_v39 = _dafny.SetOf((_2122_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F32()).Cardinality()), _dafny.IntOfUint32((_2122_v38).Cardinality()))).Uint32()).(_dafny.Set))
				var _2124_v40 _dafny.Array
				_ = _2124_v40
				var _len0_56 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_56
				var _nw371 _dafny.Array
				_ = _nw371
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw371 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) bool = func(_2125_i1 _dafny.Int) bool {
						return _this.F35()
					}
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw371 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw371).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw371).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2124_v40 = _nw371
				var _2126_v41 _dafny.Sequence
				_ = _2126_v41
				_2126_v41 = _dafny.SeqOf(_2124_v40)
				var _2127_v42 _dafny.Sequence
				_ = _2127_v42
				_2127_v42 = _dafny.SeqOf((_2126_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.IntOfUint32((_2126_v41).Cardinality()))).Uint32()).(_dafny.Array), _2124_v40)
				_2120_v36 = Companion_D2_.Create_DC6_((_2123_v39).Cardinality(), _this.F28(), (func() _dafny.Int {
					if !(_this.F29()) {
						return _dafny.IntOfUint32(((_this).F32()).Cardinality())
					}
					return _dafny.IntOfInt64(968)
				})(), _this.F29(), _dafny.IntOfUint32((_2127_v42).Cardinality()))
				var _2128_v43 _dafny.Map
				_ = _2128_v43
				_2128_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2115_v31, true)
				_2128_v43 = ((_2128_v43).Update(_2115_v31, _this.F28())).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer122 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg122 _dafny.Int) interface{} {
						return coer122(arg122)
					}
				}(func(_2129_i2 _dafny.Int) _dafny.Int {
					return (_this).F30()
				})), _2115_v31), _this.F35())
			} else {
				var _2130_v44 _dafny.Set
				_ = _2130_v44
				_2130_v44 = _dafny.SetOf(_this.F35())
				var _2131_v45 _dafny.Map
				_ = _2131_v45
				_2131_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), (_this).F30())
				var _2132_v46 _dafny.Sequence
				_ = _2132_v46
				_2132_v46 = _dafny.UnicodeSeqOfUtf8Bytes("sk")
				var _2133_v47 _dafny.Array
				_ = _2133_v47
				var _nwElement0_72 _dafny.Int = (_this).F38()
				_ = _nwElement0_72
				var _nw372 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(11))
				_ = _nw372
				(_nw372).ArraySet1(_nwElement0_72, 0)
				(_nw372).ArraySet1((_this).F34(), 1)
				(_nw372).ArraySet1((_this).F38(), 2)
				(_nw372).ArraySet1((_2130_v44).Cardinality(), 3)
				(_nw372).ArraySet1((_2131_v45).Cardinality(), 4)
				(_nw372).ArraySet1((_this).F38(), 5)
				(_nw372).ArraySet1(_dafny.IntOfUint32((_2132_v46).Cardinality()), 6)
				(_nw372).ArraySet1((_this).F38(), 7)
				(_nw372).ArraySet1(_dafny.IntOfUint32((_2132_v46).Cardinality()), 8)
				(_nw372).ArraySet1((_this).F30(), 9)
				(_nw372).ArraySet1((_this).F38(), 10)
				_2133_v47 = _nw372
				var _2134_v48 _dafny.MultiSet
				_ = _2134_v48
				_2134_v48 = _dafny.MultiSetOf(_2133_v47)
				var _2135_v50 _dafny.Array
				_ = _2135_v50
				var _nwElement0_73 _dafny.Int = (_this).F34()
				_ = _nwElement0_73
				var _nw373 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(16))
				_ = _nw373
				(_nw373).ArraySet1(_nwElement0_73, 0)
				(_nw373).ArraySet1(((_2134_v48).Cardinality()).Minus((_this).F38()), 1)
				(_nw373).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lyvg")).Cardinality())).Plus(_dafny.IntOfUint32(((_this).F33()).Cardinality())), 2)
				(_nw373).ArraySet1(_dafny.IntOfInt64(140), 3)
				(_nw373).ArraySet1((_this).F38(), 4)
				(_nw373).ArraySet1(((_this).F30()).Times((_this).F30()), 5)
				(_nw373).ArraySet1((_this).F30(), 6)
				(_nw373).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg123 _dafny.Int) interface{} {
						return coer123(arg123)
					}
				}(func(_2136_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				})), _2132_v46)).Cardinality()), 7)
				(_nw373).ArraySet1((_this).F34(), 8)
				(_nw373).ArraySet1((_this).F30(), 9)
				(_nw373).ArraySet1((_this).F38(), 10)
				(_nw373).ArraySet1(_dafny.IntOfUint32((_2132_v46).Cardinality()), 11)
				(_nw373).ArraySet1((_this).F34(), 12)
				(_nw373).ArraySet1((_this).F30(), 13)
				(_nw373).ArraySet1(Companion_Default___.SafeModuloInt((_this).F30(), (_this).F30()), 14)
				(_nw373).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (func() _dafny.Map {
					var _coll56 = _dafny.NewMapBuilder()
					_ = _coll56
					for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-644), _dafny.IntOfInt64(310))); ; {
						_compr_56, _ok61 := _iter61()
						if !_ok61 {
							break
						}
						var _2137_v49 _dafny.Int
						_2137_v49 = interface{}(_compr_56).(_dafny.Int)
						if ((_dafny.IntOfInt64(-644)).Cmp(_2137_v49) <= 0) && ((_2137_v49).Cmp(_dafny.IntOfInt64(310)) < 0) {
							_coll56.Add((_2137_v49).Minus((_this).F30()), _this.F31())
						}
					}
					return _coll56.ToMap()
				}()).Cardinality())).Cardinality(), 15)
				_2135_v50 = _nw373
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _index297
				(_2135_v50).ArraySet1((_this).F38(), (_index297).Int())
				var _2138_v51 _dafny.Set
				_ = _2138_v51
				_2138_v51 = _dafny.SetOf((_this).F30())
				var _2139_v52 D21
				_ = _2139_v52
				_2139_v52 = Companion_D21_.Create_DC51_(Companion_Default___.Fm59(_this.F35(), _this.F35(), _this.F28(), globalState))
				var _2140_v53 _dafny.Map
				_ = _2140_v53
				_2140_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _2138_v51)
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _index298
				var _rhs317 bool = !(_2138_v51).Contains((_this).F34())
				_ = _rhs317
				var _rhs318 bool = false
				_ = _rhs318
				var _rhs319 _dafny.CodePoint = Companion_Default___.Fm48(!((_2139_v52).Dtor_cf88()).Equals((_2140_v53).Update((_this).F38(), func() _dafny.Set {
					var _coll57 = _dafny.NewBuilder()
					_ = _coll57
					for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-564), _dafny.IntOfInt64(901))); ; {
						_compr_57, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _2141_v54 _dafny.Int
						_2141_v54 = interface{}(_compr_57).(_dafny.Int)
						if ((_dafny.IntOfInt64(-564)).Cmp(_2141_v54) <= 0) && ((_2141_v54).Cmp(_dafny.IntOfInt64(901)) < 0) {
							_coll57.Add(Companion_Default___.SafeModuloInt(_2141_v54, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _this.F35())).Cardinality()))
						}
					}
					return _coll57.ToSet()
				}())), _this.F35(), (_this).F38(), globalState)
				_ = _rhs319
				var _rhs320 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(756))).Uint32(), func(coer124 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg124 _dafny.Int) interface{} {
						return coer124(arg124)
					}
				}(func(_2142_i4 _dafny.Int) _dafny.Int {
					return (_this).F38()
				}))).Cardinality())
				_ = _rhs320
				var _lhs254 *GlobalState = globalState
				_ = _lhs254
				var _lhs255 *GlobalState = globalState
				_ = _lhs255
				var _lhs256 *C9 = _this
				_ = _lhs256
				var _lhs257 _dafny.Array = _2135_v50
				_ = _lhs257
				var _lhs258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _lhs258
				_lhs254.F12 = _rhs317
				_lhs255.F12 = _rhs318
				_lhs256.F36_set_(_rhs319)
				(_lhs257).ArraySet1(_rhs320, (_lhs258).Int())
				(_this).F36_set_(_this.F36())
				var _2143_v55 _dafny.Array
				_ = _2143_v55
				var _len0_57 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_57
				var _nw374 _dafny.Array
				_ = _nw374
				if _len0_57.Cmp(_dafny.Zero) == 0 {
					_nw374 = _dafny.NewArray(_len0_57)
				} else {
					var _init57 func(_dafny.Int) _dafny.Sequence = (func(_2144_v46 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2145_i5 _dafny.Int) _dafny.Sequence {
							return _2144_v46
						}
					})(_2132_v46)
					_ = _init57
					var _element0_57 = _init57(_dafny.Zero)
					_ = _element0_57
					_nw374 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
					(_nw374).ArraySet1(_element0_57, 0)
					var _nativeLen0_57 = (_len0_57).Int()
					_ = _nativeLen0_57
					for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
						(_nw374).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
					}
				}
				_2143_v55 = _nw374
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2143_v55), 0))
				_ = _index299
				(_2143_v55).ArraySet1(_2132_v46, (_index299).Int())
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2143_v55), 0))
				_ = _index300
				(_2143_v55).ArraySet1((_this).Fm3(!(!(_this.F35())), false, globalState), (_index300).Int())
				var _2146_v56 _dafny.Array
				_ = _2146_v56
				var _len0_58 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_58
				var _nw375 _dafny.Array
				_ = _nw375
				if _len0_58.Cmp(_dafny.Zero) == 0 {
					_nw375 = _dafny.NewArray(_len0_58)
				} else {
					var _init58 func(_dafny.Int) _dafny.MultiSet = func(_2147_i6 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(_this.F29())
					}
					_ = _init58
					var _element0_58 = _init58(_dafny.Zero)
					_ = _element0_58
					_nw375 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
					(_nw375).ArraySet1(_element0_58, 0)
					var _nativeLen0_58 = (_len0_58).Int()
					_ = _nativeLen0_58
					for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
						(_nw375).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
					}
				}
				_2146_v56 = _nw375
				var _2148_v57 _dafny.MultiSet
				_ = _2148_v57
				_2148_v57 = _dafny.MultiSetOf(_this.F28(), _this.F31(), _this.F31())
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_2146_v56), 0))
				_ = _index301
				(_2146_v56).ArraySet1(_2148_v57, (_index301).Int())
				var _2149_v58 _dafny.Array
				_ = _2149_v58
				var _nw376 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
				_ = _nw376
				_2149_v58 = _nw376
				var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_2149_v58), 0))
				_ = _index302
				(_2149_v58).ArraySet1(_this.F28(), (_index302).Int())
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_2146_v56), 0))
				_ = _index303
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _index304
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_2149_v58), 0))
				_ = _index305
				var _rhs321 _dafny.Sequence = (_2143_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(958), _dafny.ArrayLen((_2143_v55), 0))).Int()).(_dafny.Sequence)
				_ = _rhs321
				var _rhs322 _dafny.MultiSet = (_2148_v57).Update(_this.F31(), Companion_Default___.Abs((_this).F38()))
				_ = _rhs322
				var _rhs323 bool = (_dafny.IntOfInt64(939)).Cmp((_dafny.Zero).Minus((_this).F38())) < 0
				_ = _rhs323
				var _rhs324 _dafny.Int = (_2135_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))).Int()).(_dafny.Int)
				_ = _rhs324
				var _rhs325 bool = !(!(_this.F28()))
				_ = _rhs325
				var _lhs259 _dafny.Array = _2146_v56
				_ = _lhs259
				var _lhs260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_2146_v56), 0))
				_ = _lhs260
				var _lhs261 *GlobalState = globalState
				_ = _lhs261
				var _lhs262 _dafny.Array = _2135_v50
				_ = _lhs262
				var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _lhs263
				var _lhs264 _dafny.Array = _2149_v58
				_ = _lhs264
				var _lhs265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_2149_v58), 0))
				_ = _lhs265
				_2132_v46 = _rhs321
				(_lhs259).ArraySet1(_rhs322, (_lhs260).Int())
				_lhs261.F1 = _rhs323
				(_lhs262).ArraySet1(_rhs324, (_lhs263).Int())
				(_lhs264).ArraySet1(_rhs325, (_lhs265).Int())
				var _2150_v59 _dafny.MultiSet
				_ = _2150_v59
				_2150_v59 = _dafny.MultiSetOf((_this).F34())
				var _2151_v60 D17
				_ = _2151_v60
				_2151_v60 = Companion_D17_.Create_DC44_(_2150_v59, (_this).F34())
				var _2152_v61 _dafny.Map
				_ = _2152_v61
				_2152_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_this.F28()), (_this).F32()), _this.F29())
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _index306
				var _rhs326 D17 = func(_pat_let31_0 D17) D17 {
					return func(_2153_dt__update__tmp_h1 D17) D17 {
						return func(_pat_let32_0 _dafny.Int) D17 {
							return func(_2154_dt__update_hcf78_h0 _dafny.Int) D17 {
								return Companion_D17_.Create_DC44_((_2153_dt__update__tmp_h1).Dtor_cf77(), _2154_dt__update_hcf78_h0)
							}(_pat_let32_0)
						}(_dafny.IntOfInt64(695))
					}(_pat_let31_0)
				}(_2151_v60)
				_ = _rhs326
				var _rhs327 _dafny.Int = (_2152_v61).Cardinality()
				_ = _rhs327
				var _lhs266 _dafny.Array = _2135_v50
				_ = _lhs266
				var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_2135_v50), 0))
				_ = _lhs267
				_2151_v60 = _rhs326
				(_lhs266).ArraySet1(_rhs327, (_lhs267).Int())
			}
			var _2155_v62 _dafny.Sequence
			_ = _2155_v62
			_2155_v62 = _dafny.UnicodeSeqOfUtf8Bytes("rvpm")
			_2155_v62 = _dafny.UnicodeSeqOfUtf8Bytes("nnhqon")
			var _2156_v63 _dafny.Array
			_ = _2156_v63
			var _len0_59 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_59
			var _nw377 _dafny.Array
			_ = _nw377
			if _len0_59.Cmp(_dafny.Zero) == 0 {
				_nw377 = _dafny.NewArray(_len0_59)
			} else {
				var _init59 func(_dafny.Int) bool = (func(_2157_v62 _dafny.Sequence) func(_dafny.Int) bool {
					return func(_2158_i7 _dafny.Int) bool {
						return (_dafny.MultiSetOf(_this.F36())).IsProperSubsetOf(_dafny.MultiSetOf((_2157_v62).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_2157_v62).Cardinality()))).Uint32()).(_dafny.CodePoint)))
					}
				})(_2155_v62)
				_ = _init59
				var _element0_59 = _init59(_dafny.Zero)
				_ = _element0_59
				_nw377 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
				(_nw377).ArraySet1(_element0_59, 0)
				var _nativeLen0_59 = (_len0_59).Int()
				_ = _nativeLen0_59
				for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
					(_nw377).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
				}
			}
			_2156_v63 = _nw377
			var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_2156_v63), 0))
			_ = _index307
			(_2156_v63).ArraySet1(_this.F28(), (_index307).Int())
			var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2156_v63), 0))
			_ = _index308
			(_2156_v63).ArraySet1(_this.F31(), (_index308).Int())
			var _2159_v64 _dafny.MultiSet
			_ = _2159_v64
			_2159_v64 = _dafny.MultiSetOf(_this.F29(), _this.F29())
			var _2160_v65 _dafny.Sequence
			_ = _2160_v65
			_2160_v65 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F34()), (func() _dafny.Int {
				if (_2159_v64).Contains(_this.F31()) {
					return (_2159_v64).Multiplicity(_this.F31())
				}
				return (_dafny.SetOf(false)).Cardinality()
			})(), _dafny.IntOfInt64(431))
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_2156_v63), 0))
			_ = _index309
			var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2156_v63), 0))
			_ = _index310
			var _rhs328 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2160_v65).Cardinality()), (_2160_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.IntOfUint32((_2160_v65).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs328
			var _rhs329 bool = ((_dafny.Zero).Minus((_this).F38())).Cmp((_this).F34()) >= 0
			_ = _rhs329
			var _rhs330 bool = _this.F28()
			_ = _rhs330
			var _lhs268 *GlobalState = globalState
			_ = _lhs268
			var _lhs269 _dafny.Array = _2156_v63
			_ = _lhs269
			var _lhs270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_2156_v63), 0))
			_ = _lhs270
			var _lhs271 _dafny.Array = _2156_v63
			_ = _lhs271
			var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.ArrayLen((_2156_v63), 0))
			_ = _lhs272
			_lhs268.F13 = _rhs328
			(_lhs269).ArraySet1(_rhs329, (_lhs270).Int())
			(_lhs271).ArraySet1(_rhs330, (_lhs272).Int())
			var _2161_v66 _dafny.MultiSet
			_ = _2161_v66
			_2161_v66 = _dafny.MultiSetOf((_this).F30())
			var _rhs331 _dafny.MultiSet = (_dafny.MultiSetOf((_this).F38())).Difference(_2161_v66)
			_ = _rhs331
			var _rhs332 _dafny.Int = _dafny.IntOfInt64(224)
			_ = _rhs332
			var _lhs273 *GlobalState = globalState
			_ = _lhs273
			_2161_v66 = _rhs331
			_lhs273.F13 = _rhs332
		}
		(globalState).F26 = _this.F28()
		var _2162_v67 _dafny.Sequence
		_ = _2162_v67
		_2162_v67 = _dafny.SeqOf(_this.F29())
		var _2163_v68 _dafny.Array
		_ = _2163_v68
		var _len0_60 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_60
		var _nw378 _dafny.Array
		_ = _nw378
		if _len0_60.Cmp(_dafny.Zero) == 0 {
			_nw378 = _dafny.NewArray(_len0_60)
		} else {
			var _init60 func(_dafny.Int) _dafny.Int = func(_2164_i8 _dafny.Int) _dafny.Int {
				return (_2164_i8).Plus((_this).F30())
			}
			_ = _init60
			var _element0_60 = _init60(_dafny.Zero)
			_ = _element0_60
			_nw378 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
			(_nw378).ArraySet1(_element0_60, 0)
			var _nativeLen0_60 = (_len0_60).Int()
			_ = _nativeLen0_60
			for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
				(_nw378).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
			}
		}
		_2163_v68 = _nw378
		var _2165_v69 _dafny.MultiSet
		_ = _2165_v69
		_2165_v69 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F34()), (_this).F38())
		var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_2163_v68), 0))
		_ = _index311
		(_2163_v68).ArraySet1((_2165_v69).Cardinality(), (_index311).Int())
		var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_2163_v68), 0))
		_ = _index312
		var _rhs333 _dafny.Sequence = _2162_v67
		_ = _rhs333
		var _rhs334 _dafny.Int = (_this).F30()
		_ = _rhs334
		var _rhs335 _dafny.Int = (_dafny.Zero).Minus((_this).F34())
		_ = _rhs335
		var _rhs336 _dafny.Int = (_this).F38()
		_ = _rhs336
		var _rhs337 _dafny.Int = (_this).F30()
		_ = _rhs337
		var _lhs274 *GlobalState = globalState
		_ = _lhs274
		var _lhs275 *GlobalState = globalState
		_ = _lhs275
		var _lhs276 *GlobalState = globalState
		_ = _lhs276
		var _lhs277 _dafny.Array = _2163_v68
		_ = _lhs277
		var _lhs278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_2163_v68), 0))
		_ = _lhs278
		_2162_v67 = _rhs333
		_lhs274.F13 = _rhs334
		_lhs275.F13 = _rhs335
		_lhs276.F13 = _rhs336
		(_lhs277).ArraySet1(_rhs337, (_lhs278).Int())
		var _2166_v70 _dafny.Map
		_ = _2166_v70
		_2166_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), _this.F35())
		var _2167_v71 _dafny.Map
		_ = _2167_v71
		_2167_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2166_v70, (_2163_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_2163_v68), 0))).Int()).(_dafny.Int))
		var _2168_v72 _dafny.Sequence
		_ = _2168_v72
		_2168_v72 = _dafny.SeqOf(_2167_v71, (_2167_v71).Merge(_2167_v71), _2167_v71, _2167_v71)
		r0 = (_2168_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.IntOfUint32((_2168_v72).Cardinality()))).Uint32()).(_dafny.Map)
		return r0
	}
}
func (_this *C9) M5(globalState *GlobalState) (_dafny.Int, _dafny.Array, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _2169_v0 _dafny.Array
		_ = _2169_v0
		var _nw379 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw379
		_2169_v0 = _nw379
		var _2170_v1 _dafny.Map
		_ = _2170_v1
		_2170_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(778))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg125 _dafny.Int) interface{} {
				return coer125(arg125)
			}
		}(func(_2171_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F36()
		})), (func() _dafny.Array {
			if false {
				return _2169_v0
			}
			return _2169_v0
		})())
		_2170_v1 = (_2170_v1).Update((_this).Fm3(_this.F28(), _this.F31(), globalState), _2169_v0)
		var _2172_v2 _dafny.Array
		_ = _2172_v2
		var _len0_61 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_61
		var _nw380 _dafny.Array
		_ = _nw380
		if _len0_61.Cmp(_dafny.Zero) == 0 {
			_nw380 = _dafny.NewArray(_len0_61)
		} else {
			var _init61 func(_dafny.Int) bool = func(_2173_i1 _dafny.Int) bool {
				return ((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F34()))).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F31())).Cardinality()) >= 0
			}
			_ = _init61
			var _element0_61 = _init61(_dafny.Zero)
			_ = _element0_61
			_nw380 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
			(_nw380).ArraySet1(_element0_61, 0)
			var _nativeLen0_61 = (_len0_61).Int()
			_ = _nativeLen0_61
			for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
				(_nw380).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
			}
		}
		_2172_v2 = _nw380
		var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))
		_ = _index313
		(_2172_v2).ArraySet1(!((_this.F31()) && (true)), (_index313).Int())
		var _2174_v3 _dafny.MultiSet
		_ = _2174_v3
		_2174_v3 = _dafny.MultiSetOf(_this.F28())
		var _2175_v4 _dafny.Map
		_ = _2175_v4
		_2175_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _2174_v3)
		var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
		_ = _index314
		(_2169_v0).ArraySet1(_dafny.IntOfUint32(((_this).Fm3(_this.F35(), ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), globalState)).Cardinality()), (_index314).Int())
		var _2176_v5 _dafny.MultiSet
		_ = _2176_v5
		_2176_v5 = _dafny.MultiSetOf((_this).F38(), _dafny.IntOfInt64(325))
		var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_2169_v0), 0))
		_ = _index315
		(_2169_v0).ArraySet1(Companion_Default___.Fm39(_this.F31(), _2176_v5, _this.F28(), globalState), (_index315).Int())
		var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))
		_ = _index316
		var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
		_ = _index317
		var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_2169_v0), 0))
		_ = _index318
		var _rhs338 _dafny.CodePoint = _this.F36()
		_ = _rhs338
		var _rhs339 bool = _this.F29()
		_ = _rhs339
		var _rhs340 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-867), _2174_v3)).Merge(_2175_v4)).Merge((_2175_v4).Merge(_2175_v4))
		_ = _rhs340
		var _rhs341 _dafny.Int = (_this).F30()
		_ = _rhs341
		var _rhs342 _dafny.Int = (_this).F38()
		_ = _rhs342
		var _lhs279 *C9 = _this
		_ = _lhs279
		var _lhs280 _dafny.Array = _2172_v2
		_ = _lhs280
		var _lhs281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))
		_ = _lhs281
		var _lhs282 _dafny.Array = _2169_v0
		_ = _lhs282
		var _lhs283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
		_ = _lhs283
		var _lhs284 _dafny.Array = _2169_v0
		_ = _lhs284
		var _lhs285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_2169_v0), 0))
		_ = _lhs285
		_lhs279.F36_set_(_rhs338)
		(_lhs280).ArraySet1(_rhs339, (_lhs281).Int())
		_2175_v4 = _rhs340
		(_lhs282).ArraySet1(_rhs341, (_lhs283).Int())
		(_lhs284).ArraySet1(_rhs342, (_lhs285).Int())
		if _dafny.Companion_Sequence_.IsProperPrefixOf((_this).F32(), (_this).F32()) {
			(globalState).F13 = (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)
			_2169_v0 = _2169_v0
			var _2177_v6 _dafny.Sequence
			_ = _2177_v6
			_2177_v6 = _dafny.SeqOf(_this.F36(), _this.F36())
			var _2178_v7 _dafny.Sequence
			_ = _2178_v7
			_2178_v7 = _dafny.SeqOf(true, false, (func() bool {
				if !(_this.F35()) {
					return _this.F29()
				}
				return _this.F29()
			})(), _this.F28())
			var _2179_v8 _dafny.Set
			_ = _2179_v8
			_2179_v8 = _dafny.SetOf((_this).F34())
			var _rhs343 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F36()), _2177_v6)
			_ = _rhs343
			var _rhs344 bool = ((_this).F34()).Cmp((_dafny.Zero).Minus((_this).F38())) < 0
			_ = _rhs344
			var _rhs345 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F32(), _2178_v7)
			_ = _rhs345
			var _rhs346 _dafny.Int = (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs346
			var _rhs347 _dafny.Set = _2179_v8
			_ = _rhs347
			var _lhs286 *GlobalState = globalState
			_ = _lhs286
			var _lhs287 *GlobalState = globalState
			_ = _lhs287
			_2177_v6 = _rhs343
			_lhs286.F26 = _rhs344
			_2178_v7 = _rhs345
			_lhs287.F13 = _rhs346
			_2179_v8 = _rhs347
			var _2180_v9 _dafny.Map
			_ = _2180_v9
			_2180_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2169_v0, (_this).Fm3(_this.F31(), (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool), globalState))
			_2180_v9 = (_2180_v9).Update(_2169_v0, _2177_v6)
			var _2181_v10 _dafny.Map
			_ = _2181_v10
			_2181_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(995), !((_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool)))
			_2181_v10 = (_2181_v10).Update((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int), _this.F35())
		} else {
			var _2182_v11 _dafny.Sequence
			_ = _2182_v11
			_2182_v11 = _dafny.UnicodeSeqOfUtf8Bytes("aki")
			var _2183_v12 _dafny.Sequence
			_ = _2183_v12
			_2183_v12 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ox")).Cardinality()))
			var _2184_v13 _dafny.Sequence
			_ = _2184_v13
			_2184_v13 = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_2183_v12).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F30()), _dafny.IntOfUint32((_2183_v12).Cardinality()))).Uint32()).(_dafny.Int))))
			var _rhs348 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg126 _dafny.Int) interface{} {
					return coer126(arg126)
				}
			}(func(_2185_i2 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			})), _2182_v11)
			_ = _rhs348
			var _rhs349 bool = !((_this.F28()) || ((_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool))) || (!(Companion_Default___.Fm60(globalState)).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _2183_v12)).Update((_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool), _2183_v12)))
			_ = _rhs349
			var _rhs350 _dafny.Int = (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs350
			var _lhs288 *C9 = _this
			_ = _lhs288
			_2182_v11 = _rhs348
			_lhs288.F35_set_(_rhs349)
			r0 = _rhs350
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
			_ = _index319
			(_2169_v0).ArraySet1((_dafny.Zero).Minus(((_this).F38()).Plus(_dafny.IntOfUint32(((_this).F32()).Cardinality()))), (_index319).Int())
			var _2186_v14 _dafny.Map
			_ = _2186_v14
			var _out71 _dafny.Map
			_ = _out71
			_out71 = (_this).M4(globalState)
			_2186_v14 = _out71
			(globalState).F13 = (_this).F30()
			var _2187_v15 D0
			_ = _2187_v15
			_2187_v15 = Companion_D0_.Create_DC1_(_2174_v3)
			var _pat_let_tv23 = _2174_v3
			_ = _pat_let_tv23
			var _2188_v16 _dafny.Map
			_ = _2188_v16
			_2188_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let33_0 D0) D0 {
				return func(_2189_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let34_0 _dafny.MultiSet) D0 {
						return func(_2190_dt__update_hcf1_h0 _dafny.MultiSet) D0 {
							return Companion_D0_.Create_DC1_(_2190_dt__update_hcf1_h0)
						}(_pat_let34_0)
					}(_pat_let_tv23)
				}(_pat_let33_0)
			}(_2187_v15), true)
			var _2191_v17 D8
			_ = _2191_v17
			_2191_v17 = Companion_D8_.Create_DC21_(_2188_v16)
			var _2192_v18 _dafny.Map
			_ = _2192_v18
			_2192_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F38(), _2188_v16)
			var _2193_v19 _dafny.Set
			_ = _2193_v19
			_2193_v19 = _dafny.SetOf((_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool), _this.F29())
			var _2194_v21 _dafny.MultiSet
			_ = _2194_v21
			_2194_v21 = _dafny.MultiSetOf(_2187_v15)
			var _pat_let_tv24 = _2194_v21
			_ = _pat_let_tv24
			var _pat_let_tv25 = _2194_v21
			_ = _pat_let_tv25
			var _2195_v22 _dafny.Array
			_ = _2195_v22
			var _nwElement0_74 D8 = _2191_v17
			_ = _nwElement0_74
			var _nw381 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(18))
			_ = _nw381
			(_nw381).ArraySet1(_nwElement0_74, 0)
			(_nw381).ArraySet1(func(_pat_let35_0 D8) D8 {
				return func(_2196_dt__update__tmp_h1 D8) D8 {
					return func(_pat_let36_0 _dafny.Map) D8 {
						return func(_2198_dt__update_hcf40_h0 _dafny.Map) D8 {
							return Companion_D8_.Create_DC21_(_2198_dt__update_hcf40_h0)
						}(_pat_let36_0)
					}(func() _dafny.Map {
						var _coll58 = _dafny.NewMapBuilder()
						_ = _coll58
						for _iter63 := _dafny.Iterate((_pat_let_tv24).Elements()); ; {
							_compr_58, _ok63 := _iter63()
							if !_ok63 {
								break
							}
							var _2197_v20 D0
							_2197_v20 = interface{}(_compr_58).(D0)
							if (_pat_let_tv25).Contains(_2197_v20) {
								_coll58.Add(_2197_v20, _this.F29())
							}
						}
						return _coll58.ToMap()
					}())
				}(_pat_let35_0)
			}(Companion_D8_.Create_DC21_((func() _dafny.Map {
				if (_2192_v18).Contains((_2193_v19).Cardinality()) {
					return (_2192_v18).Get((_2193_v19).Cardinality()).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2187_v15, (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool))
			})())), 1)
			(_nw381).ArraySet1(_2191_v17, 2)
			(_nw381).ArraySet1(_2191_v17, 3)
			(_nw381).ArraySet1(_2191_v17, 4)
			(_nw381).ArraySet1(Companion_Default___.Fm61(_2182_v11, (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool), (_this).F30(), globalState), 5)
			(_nw381).ArraySet1(Companion_D8_.Create_DC21_(_2188_v16), 6)
			(_nw381).ArraySet1(_2191_v17, 7)
			(_nw381).ArraySet1(Companion_D8_.Create_DC21_(_2188_v16), 8)
			(_nw381).ArraySet1(_2191_v17, 9)
			(_nw381).ArraySet1(_2191_v17, 10)
			(_nw381).ArraySet1(Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2187_v15, _this.F35())), 11)
			(_nw381).ArraySet1(_2191_v17, 12)
			(_nw381).ArraySet1(_2191_v17, 13)
			(_nw381).ArraySet1(Companion_D8_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_this.F31(), true)), _this.F29())), 14)
			(_nw381).ArraySet1(_2191_v17, 15)
			(_nw381).ArraySet1(_2191_v17, 16)
			(_nw381).ArraySet1(_2191_v17, 17)
			_2195_v22 = _nw381
			var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_2195_v22), 0))
			_ = _index320
			(_2195_v22).ArraySet1(_2191_v17, (_index320).Int())
			var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_2195_v22), 0))
			_ = _index321
			var _rhs351 D8 = _2191_v17
			_ = _rhs351
			var _rhs352 _dafny.MultiSet = ((_2176_v5).Union(_dafny.MultiSetFromSeq(_2184_v13))).Difference((_2176_v5).Difference(_2176_v5))
			_ = _rhs352
			var _lhs289 _dafny.Array = _2195_v22
			_ = _lhs289
			var _lhs290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_2195_v22), 0))
			_ = _lhs290
			(_lhs289).ArraySet1(_rhs351, (_lhs290).Int())
			_2176_v5 = _rhs352
		}
		(globalState).F13 = (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)
		var _2199_i3 _dafny.Int
		_ = _2199_i3
		_2199_i3 = _dafny.Zero
		{
			for (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2169_v0, (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool))).Cardinality()).Times((_this).F34())).Cmp(((_2174_v3).Update(((_this).F32()).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_2174_v3).Contains((_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool)) {
					return (_2174_v3).Multiplicity((_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool))
				}
				return (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)
			})(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), Companion_Default___.Abs((_this).F34()))).Cardinality()) > 0 {
				{
					if (_2199_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_2199_i3 = (_2199_i3).Plus(_dafny.One)
					var _2200_v23 _dafny.Sequence
					_ = _2200_v23
					_2200_v23 = _dafny.UnicodeSeqOfUtf8Bytes("wejhofhx")
					var _2201_v24 _dafny.Map
					_ = _2201_v24
					_2201_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2200_v23, (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool))
					var _2202_v25 _dafny.Array
					_ = _2202_v25
					var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
					_ = _nw382
					_2202_v25 = _nw382
					var _2203_v26 *C0
					_ = _2203_v26
					var _nw383 *C0 = New_C0_()
					_ = _nw383
					_nw383.Ctor__(_2202_v25)
					_2203_v26 = _nw383
					var _2204_v27 D9
					_ = _2204_v27
					_2204_v27 = Companion_D9_.Create_DC27_(_2203_v26, _2200_v23)
					var _2205_v28 _dafny.Map
					_ = _2205_v28
					_2205_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _2204_v27)
					var _2206_v29 _dafny.Map
					_ = _2206_v29
					_2206_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int), _2205_v28)
					var _rhs353 _dafny.Map = _2201_v24
					_ = _rhs353
					var _rhs354 _dafny.Int = Companion_Default___.SafeModuloInt(((_2206_v29).Merge(_2206_v29)).Cardinality(), (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int))
					_ = _rhs354
					var _rhs355 _dafny.Array = _2169_v0
					_ = _rhs355
					var _lhs291 *GlobalState = globalState
					_ = _lhs291
					_2201_v24 = _rhs353
					_lhs291.F13 = _rhs354
					_2169_v0 = _rhs355
					var _2207_v30 _dafny.Sequence
					_ = _2207_v30
					_2207_v30 = _dafny.SeqOf((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int), (_this).F38())
					var _2208_v31 _dafny.Array
					_ = _2208_v31
					var _nwElement0_75 _dafny.Sequence = _2207_v30
					_ = _nwElement0_75
					var _nw384 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(16))
					_ = _nw384
					(_nw384).ArraySet1(_nwElement0_75, 0)
					(_nw384).ArraySet1(_2207_v30, 1)
					(_nw384).ArraySet1(_dafny.Companion_Sequence_.Update(_2207_v30, (Companion_Default___.SafeIndex((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2207_v30).Cardinality()))).Uint32(), (_this).F38()), 2)
					(_nw384).ArraySet1(_2207_v30, 3)
					(_nw384).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(920))).Uint32(), func(coer127 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg127 _dafny.Int) interface{} {
							return coer127(arg127)
						}
					}((func(_2209_v5 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
						return func(_2210_i4 _dafny.Int) _dafny.Int {
							return (_2209_v5).Cardinality()
						}
					})(_2176_v5))), 4)
					(_nw384).ArraySet1(_2207_v30, 5)
					(_nw384).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2207_v30, _2207_v30), 6)
					(_nw384).ArraySet1(_2207_v30, 7)
					(_nw384).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2207_v30, _dafny.SeqOf((_this).F30(), (_this).F30())), 8)
					(_nw384).ArraySet1(Companion_Default___.Fm23(_this.F31(), (_2174_v3).Cardinality(), globalState), 9)
					(_nw384).ArraySet1(_2207_v30, 10)
					(_nw384).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2207_v30, _2207_v30), (Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2207_v30, _2207_v30)).Cardinality()))).Uint32(), (_this).F34()), 11)
					(_nw384).ArraySet1((func() _dafny.Sequence {
						if _this.F29() {
							return Companion_Default___.Fm23(!(_this.F35()), (_this).F38(), globalState)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer128 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg128 _dafny.Int) interface{} {
								return coer128(arg128)
							}
						}(func(_2211_i5 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("idwc")).Cardinality())
						}))
					})(), 12)
					(_nw384).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2207_v30, _2207_v30), 13)
					(_nw384).ArraySet1(_2207_v30, 14)
					(_nw384).ArraySet1(_2207_v30, 15)
					_2208_v31 = _nw384
					var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
					_ = _index322
					var _rhs356 _dafny.Array = _2208_v31
					_ = _rhs356
					var _rhs357 _dafny.Array = _2169_v0
					_ = _rhs357
					var _rhs358 bool = _this.F31()
					_ = _rhs358
					var _rhs359 _dafny.Int = (func() _dafny.Int {
						if (_this).Fm6(globalState) {
							return _dafny.IntOfUint32((_2200_v23).Cardinality())
						}
						return (_dafny.Zero).Minus(((_this).F38()).Times((_this).F34()))
					})()
					_ = _rhs359
					var _lhs292 *GlobalState = globalState
					_ = _lhs292
					var _lhs293 _dafny.Array = _2169_v0
					_ = _lhs293
					var _lhs294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
					_ = _lhs294
					_2208_v31 = _rhs356
					_2169_v0 = _rhs357
					_lhs292.F1 = _rhs358
					(_lhs293).ArraySet1(_rhs359, (_lhs294).Int())
					var _2212_v32 _dafny.Map
					_ = _2212_v32
					_2212_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_2200_v23).Cardinality())), ((_this).F34()).Plus((_this).F38()))
					var _2213_v33 _dafny.Map
					_ = _2213_v33
					_2213_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2212_v32, (_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int))
					_2212_v32 = (_2212_v32).Update(((_2213_v33).Merge(_2213_v33)).Cardinality(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-265), (_this).F38())))
					var _2214_v34 _dafny.Map
					_ = _2214_v34
					_2214_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F31()) || (_this.F31()), Companion_Default___.Fm62(_this.F36(), (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool), _this.F29(), _this.F28(), globalState))
					var _2215_v35 _dafny.Map
					_ = _2215_v35
					_2215_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_this).F34())
					var _2216_v36 _dafny.Set
					_ = _2216_v36
					_2216_v36 = _dafny.SetOf((_this).F38(), (_this).F34())
					var _2217_v37 D21
					_ = _2217_v37
					_2217_v37 = Companion_D21_.Create_DC51_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2215_v35).Cardinality(), _2216_v36))
					_2214_v34 = (_2214_v34).Update(_this.F35(), _2217_v37)
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _hi11 _dafny.Int = (_this).F38()
		_ = _hi11
		for _2218_i6 := (_this).F34(); _2218_i6.Cmp(_hi11) < 0; _2218_i6 = _2218_i6.Plus(_dafny.One) {
			var _2219_v38 _dafny.Sequence
			_ = _2219_v38
			_2219_v38 = _dafny.SeqOf(_2172_v2)
			var _2220_v39 _dafny.Sequence
			_ = _2220_v39
			_2220_v39 = _2219_v38
			_2220_v39 = _2220_v39
			var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
			_ = _index323
			(_2169_v0).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfInt64(88)).Plus(_2218_i6)), (_index323).Int())
			var _2221_v40 _dafny.Map
			_ = _2221_v40
			_2221_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _2172_v2)
			_2221_v40 = (_2221_v40).Update(_this.F36(), _2172_v2)
			if _this.F35() {
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
				_ = _index324
				(_2169_v0).ArraySet1(_dafny.IntOfInt64(55), (_index324).Int())
				(globalState).F21 = _2172_v2
				var _2222_v41 _dafny.Sequence
				_ = _2222_v41
				_2222_v41 = _dafny.UnicodeSeqOfUtf8Bytes("fvq")
				var _2223_v42 *C4
				_ = _2223_v42
				var _nw385 *C4 = New_C4_()
				_ = _nw385
				_nw385.Ctor__(_dafny.IntOfInt64(718), true, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(_this.F31())), (_this).F32()), (_this).F33(), _dafny.IntOfUint32((_2222_v41).Cardinality()), (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool), _this.F35(), _this.F29())
				_2223_v42 = _nw385
				var _2224_v43 _dafny.Map
				_ = _2224_v43
				_2224_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F35())
				var _rhs360 bool = (_2172_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))).Int()).(bool)
				_ = _rhs360
				var _rhs361 _dafny.Sequence = _2222_v41
				_ = _rhs361
				var _rhs362 _dafny.Array = _2169_v0
				_ = _rhs362
				var _rhs363 bool = !((_2174_v3).IsSubsetOf(_dafny.MultiSetOf((func() bool {
					if (_2224_v43).Contains(_this.F28()) {
						return (_2224_v43).Get(_this.F28()).(bool)
					}
					return _this.F29()
				})(), _this.F31(), false, _this.F29())))
				_ = _rhs363
				var _rhs364 *C4 = _2223_v42
				_ = _rhs364
				var _lhs295 *C9 = _this
				_ = _lhs295
				var _lhs296 *C9 = _this
				_ = _lhs296
				_lhs295.F31_set_(_rhs360)
				_2222_v41 = _rhs361
				_2169_v0 = _rhs362
				_lhs296.F29_set_(_rhs363)
				_2223_v42 = _rhs364
				var _2225_v44 _dafny.Sequence
				_ = _2225_v44
				_2225_v44 = _dafny.SeqOf(_2174_v3)
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_2172_v2), 0))
				_ = _index325
				(_2172_v2).ArraySet1((func() bool {
					if (((_2225_v44).Select((Companion_Default___.SafeIndex((_this).F38(), _dafny.IntOfUint32((_2225_v44).Cardinality()))).Uint32()).(_dafny.MultiSet)).Update(_this.F31(), Companion_Default___.Abs((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)))).IsDisjointFrom(_2174_v3) {
						return _this.F35()
					}
					return true
				})(), (_index325).Int())
				var _2226_v45 _dafny.Array
				_ = _2226_v45
				var _len0_62 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_62
				var _nw386 _dafny.Array
				_ = _nw386
				if _len0_62.Cmp(_dafny.Zero) == 0 {
					_nw386 = _dafny.NewArray(_len0_62)
				} else {
					var _init62 func(_dafny.Int) _dafny.Sequence = (func(_2227_v41 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2228_i7 _dafny.Int) _dafny.Sequence {
							return _2227_v41
						}
					})(_2222_v41)
					_ = _init62
					var _element0_62 = _init62(_dafny.Zero)
					_ = _element0_62
					_nw386 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
					(_nw386).ArraySet1(_element0_62, 0)
					var _nativeLen0_62 = (_len0_62).Int()
					_ = _nativeLen0_62
					for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
						(_nw386).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
					}
				}
				_2226_v45 = _nw386
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2226_v45), 0))
				_ = _index326
				(_2226_v45).ArraySet1(_2222_v41, (_index326).Int())
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((_2226_v45), 0))
				_ = _index327
				(_2226_v45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2222_v41, _2222_v41), (_this).Fm3(_this.F28(), _this.F31(), globalState)), (_index327).Int())
			} else {
				var _2229_v46 _dafny.Map
				_ = _2229_v46
				_2229_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F31())
				(globalState).F13 = (_this).Fm17(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F28())).Merge((_2229_v46).Update(_this.F28(), true))).Cardinality(), _2218_i6, globalState)
				var _2230_v47 _dafny.Sequence
				_ = _2230_v47
				_2230_v47 = _dafny.UnicodeSeqOfUtf8Bytes("gmac")
				_2230_v47 = _2230_v47
				var _2231_v48 _dafny.Array
				_ = _2231_v48
				var _nw387 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
				_ = _nw387
				_2231_v48 = _nw387
				var _2232_v49 *C0
				_ = _2232_v49
				var _nw388 *C0 = New_C0_()
				_ = _nw388
				_nw388.Ctor__(_2231_v48)
				_2232_v49 = _nw388
				_2232_v49 = _2232_v49
				(globalState).F13 = _2218_i6
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))
				_ = _index328
				(_2169_v0).ArraySet1((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int), (_index328).Int())
			}
		}
		var _2233_v50 _dafny.MultiSet
		_ = _2233_v50
		_2233_v50 = _dafny.MultiSetOf(_2174_v3, _2174_v3, Companion_Default___.Fm33(_this.F29(), globalState))
		r0 = ((_2233_v50).Union(((_2233_v50).Update(_2174_v3, Companion_Default___.Abs((_2169_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_2169_v0), 0))).Int()).(_dafny.Int)))).Update(_2174_v3, Companion_Default___.Abs(_dafny.IntOfInt64(444))))).Cardinality()
		var _nw389 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
		_ = _nw389
		r1 = _nw389
		r2 = _2172_v2
		return r0, r1, r2
	}
}
func (_this *C9) F38() _dafny.Int {
	{
		return _this._f38
	}
}
func (_this *C9) F39() _dafny.Map {
	{
		return _this._f39
	}
}

// End of class C9

// Definition of class C10
type C10 struct {
	_f28 bool
	_f29 bool
	_f31 bool
	_f35 bool
	_f36 _dafny.CodePoint
	_f34 _dafny.Int
	_f32 _dafny.Sequence
	_f33 _dafny.Sequence
	_f30 _dafny.Int
}

func New_C10_() *C10 {
	_this := C10{}

	_this._f28 = false
	_this._f29 = false
	_this._f31 = false
	_this._f35 = false
	_this._f36 = _dafny.CodePoint('D')
	_this._f34 = _dafny.Zero
	_this._f32 = _dafny.EmptySeq
	_this._f33 = _dafny.EmptySeq
	_this._f30 = _dafny.Zero
	return &_this
}

type CompanionStruct_C10_ struct {
}

var Companion_C10_ = CompanionStruct_C10_{}

func (_this *C10) Equals(other *C10) bool {
	return _this == other
}

func (_this *C10) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C10)
	return ok && _this.Equals(other)
}

func (*C10) String() string {
	return "_module.C10"
}

func Type_C10_() _dafny.TypeDescriptor {
	return type_C10_{}
}

type type_C10_ struct {
}

func (_this type_C10_) Default() interface{} {
	return (*C10)(nil)
}

func (_this type_C10_) String() string {
	return "main.C10"
}
func (_this *C10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C10{}
var _ T3 = &C10{}
var _ T2 = &C10{}
var _ T1 = &C10{}
var _ T0 = &C10{}
var _ _dafny.TraitOffspring = &C10{}

func (_this *C10) F28() bool {
	return _this._f28
}
func (_this *C10) F28_set_(value bool) {
	_this._f28 = value
}
func (_this *C10) F29() bool {
	return _this._f29
}
func (_this *C10) F29_set_(value bool) {
	_this._f29 = value
}
func (_this *C10) F31() bool {
	return _this._f31
}
func (_this *C10) F31_set_(value bool) {
	_this._f31 = value
}
func (_this *C10) F35() bool {
	return _this._f35
}
func (_this *C10) F35_set_(value bool) {
	_this._f35 = value
}
func (_this *C10) F36() _dafny.CodePoint {
	return _this._f36
}
func (_this *C10) F36_set_(value _dafny.CodePoint) {
	_this._f36 = value
}
func (_this *C10) F34() _dafny.Int {
	return _this._f34
}
func (_this *C10) F32() _dafny.Sequence {
	return _this._f32
}
func (_this *C10) F33() _dafny.Sequence {
	return _this._f33
}
func (_this *C10) F30() _dafny.Int {
	return _this._f30
}
func (_this *C10) Ctor__(f36 _dafny.CodePoint, f34 _dafny.Int, f35 bool, f32 _dafny.Sequence, f33 _dafny.Sequence, f30 _dafny.Int, f31 bool, f28 bool, f29 bool) {
	{
		(_this)._f36 = f36
		(_this)._f34 = f34
		(_this)._f35 = f35
		(_this)._f32 = f32
		(_this)._f33 = f33
		(_this)._f30 = f30
		(_this)._f31 = f31
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C10) Fm7(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(!((!(!(_this.F29()))) || (_this.F35())))
	}
}
func (_this *C10) Fm6(globalState *GlobalState) bool {
	{
		return ((func() D0 {
			if _this.F35() {
				return Companion_D0_.Create_DC0_(_this.F28())
			}
			return Companion_D0_.Create_DC0_(_this.F31())
		})()).Dtor_cf0()
	}
}
func (_this *C10) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf((_this).F34(), (_this).F34())).Difference(func() _dafny.Set {
			var _coll59 = _dafny.NewBuilder()
			_ = _coll59
			for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(507), _dafny.IntOfInt64(18))); ; {
				_compr_59, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _2234_v0 _dafny.Int
				_2234_v0 = interface{}(_compr_59).(_dafny.Int)
				if ((_dafny.IntOfInt64(507)).Cmp(_2234_v0) <= 0) && ((_2234_v0).Cmp(_dafny.IntOfInt64(18)) < 0) {
					_coll59.Add((_2234_v0).Minus((_this).F34()))
				}
			}
			return _coll59.ToSet()
		}())
	}
}
func (_this *C10) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var _source34 D0 = Companion_D0_.Create_DC1_(_dafny.MultiSetFromSeq((_this).F32()))
		_ = _source34
		if _source34.Is_DC1() {
			var _2235___mcc_h0 _dafny.MultiSet = _source34.Get_().(D0_DC1).Cf1
			_ = _2235___mcc_h0
			var _2236_cf1 _dafny.MultiSet = _2235___mcc_h0
			_ = _2236_cf1
			return _dafny.Companion_Sequence_.Concatenate((_this).F32(), _dafny.SeqOf(_this.F28()))
		} else {
			var _2237___mcc_h1 bool = _source34.Get_().(D0_DC0).Cf0
			_ = _2237___mcc_h1
			var _2238_cf0 bool = _2237___mcc_h1
			_ = _2238_cf0
			return _dafny.Companion_Sequence_.Concatenate(((_this).F33()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F33()).Cardinality()))).Uint32()).(_dafny.Sequence), (_this).F32())
		}
	}
}
func (_this *C10) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lgpghnx"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer129 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg129 _dafny.Int) interface{} {
				return coer129(arg129)
			}
		}(func(_2239_i0 _dafny.Int) _dafny.CodePoint {
			return _this.F36()
		})))
	}
}
func (_this *C10) Fm2(globalState *GlobalState) _dafny.Map {
	{
		if _this.F31() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_dafny.Zero).Minus((_this).F30()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), (_this).F30()))
		} else {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_dafny.Zero).Minus((_this).F34()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _dafny.IntOfInt64(936)))
		}
	}
}
func (_this *C10) Fm8(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(361)).Times((_dafny.IntOfInt64(369)).Times((_this).F30()))
	}
}
func (_this *C10) M1(globalState *GlobalState) (bool, _dafny.Int, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		(_this).F31_set_(_this.F28())
		var _2240_v0 _dafny.MultiSet
		_ = _2240_v0
		_2240_v0 = _dafny.MultiSetOf(_this.F29(), _this.F31(), _this.F29())
		var _2241_v1 D0
		_ = _2241_v1
		_2241_v1 = Companion_D0_.Create_DC1_(_2240_v0)
		var _source35 D0 = _2241_v1
		_ = _source35
		if _source35.Is_DC1() {
			var _2242___mcc_h0 _dafny.MultiSet = _source35.Get_().(D0_DC1).Cf1
			_ = _2242___mcc_h0
			var _2243_cf1 _dafny.MultiSet = _2242___mcc_h0
			_ = _2243_cf1
			var _2244_v2 _dafny.Array
			_ = _2244_v2
			var _nw390 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
			_ = _nw390
			_2244_v2 = _nw390
			var _2245_v3 *C0
			_ = _2245_v3
			var _nw391 *C0 = New_C0_()
			_ = _nw391
			_nw391.Ctor__(_2244_v2)
			_2245_v3 = _nw391
			var _2246_v4 _dafny.Sequence
			_ = _2246_v4
			_2246_v4 = _dafny.UnicodeSeqOfUtf8Bytes("hdvneponb")
			var _2247_v5 _dafny.Sequence
			_ = _2247_v5
			_2247_v5 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F34()))
			var _2248_v6 _dafny.MultiSet
			_ = _2248_v6
			_2248_v6 = _dafny.MultiSetOf((_this).F30())
			var _2249_v7 _dafny.Array
			_ = _2249_v7
			var _len0_63 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_63
			var _nw392 _dafny.Array
			_ = _nw392
			if _len0_63.Cmp(_dafny.Zero) == 0 {
				_nw392 = _dafny.NewArray(_len0_63)
			} else {
				var _init63 func(_dafny.Int) bool = func(_2250_i0 _dafny.Int) bool {
					return _this.F28()
				}
				_ = _init63
				var _element0_63 = _init63(_dafny.Zero)
				_ = _element0_63
				_nw392 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
				(_nw392).ArraySet1(_element0_63, 0)
				var _nativeLen0_63 = (_len0_63).Int()
				_ = _nativeLen0_63
				for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
					(_nw392).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
				}
			}
			_2249_v7 = _nw392
			var _2251_v8 _dafny.Map
			_ = _2251_v8
			_2251_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2249_v7, _2248_v6)
			var _2252_v9 _dafny.Array
			_ = _2252_v9
			var _nwElement0_76 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _nwElement0_76
			var _nw393 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(12))
			_ = _nw393
			(_nw393).ArraySet1(_nwElement0_76, 0)
			(_nw393).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2246_v4, _2246_v4)).Cardinality()), 1)
			(_nw393).ArraySet1(_dafny.IntOfUint32((_2247_v5).Cardinality()), 2)
			(_nw393).ArraySet1((_2248_v6).Cardinality(), 3)
			(_nw393).ArraySet1((_this).F34(), 4)
			(_nw393).ArraySet1((_this).F34(), 5)
			(_nw393).ArraySet1((_this).F30(), 6)
			(_nw393).ArraySet1((_this).F30(), 7)
			(_nw393).ArraySet1((_this).F30(), 8)
			(_nw393).ArraySet1((_2251_v8).Cardinality(), 9)
			(_nw393).ArraySet1((_this).F30(), 10)
			(_nw393).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F34(), _dafny.IntOfInt64(590)), 11)
			_2252_v9 = _nw393
			var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2252_v9), 0))
			_ = _index329
			(_2252_v9).ArraySet1((_this).F30(), (_index329).Int())
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2252_v9), 0))
			_ = _index330
			(_2252_v9).ArraySet1((_dafny.Zero).Minus(((_this).Fm8(false, _this.F36(), _this.F29(), globalState)).Plus((_2247_v5).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_2247_v5).Cardinality()))).Uint32()).(_dafny.Int))), (_index330).Int())
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2252_v9), 0))
			_ = _index331
			(_2252_v9).ArraySet1((_this).F30(), (_index331).Int())
			var _2253_v10 _dafny.Map
			_ = _2253_v10
			_2253_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (Companion_D1_.Create_DC2_(_2249_v7)).Dtor_cf2())
			var _2254_v11 _dafny.MultiSet
			_ = _2254_v11
			_2254_v11 = _dafny.MultiSetOf((func() _dafny.Array {
				if (_2253_v10).Contains(_this.F28()) {
					return (_2253_v10).Get(_this.F28()).(_dafny.Array)
				}
				return _2249_v7
			})())
			var _2255_v12 _dafny.Set
			_ = _2255_v12
			_2255_v12 = _dafny.SetOf((_2252_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2252_v9), 0))).Int()).(_dafny.Int))
			var _2256_v14 _dafny.Int
			_ = _2256_v14
			var _out72 _dafny.Int
			_ = _out72
			_out72 = Companion_Default___.M0(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((func() _dafny.Int {
				if (_2254_v11).Contains(_2249_v7) {
					return (_2254_v11).Multiplicity(_2249_v7)
				}
				return (_2252_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2252_v9), 0))).Int()).(_dafny.Int)
			})()), _2247_v5), _this.F29(), (func() _dafny.Int {
				if (_2248_v6).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_this.F28())).Cardinality())) {
					return (_2248_v6).Multiplicity(_dafny.IntOfUint32((_dafny.SeqOf(_this.F28())).Cardinality()))
				}
				return _dafny.IntOfUint32((_2247_v5).Cardinality())
			})(), (func() _dafny.Set {
				var _coll60 = _dafny.NewBuilder()
				_ = _coll60
				for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(580), _dafny.IntOfInt64(286))); ; {
					_compr_60, _ok65 := _iter65()
					if !_ok65 {
						break
					}
					var _2257_v13 _dafny.Int
					_2257_v13 = interface{}(_compr_60).(_dafny.Int)
					if ((_dafny.IntOfInt64(580)).Cmp(_2257_v13) <= 0) && ((_2257_v13).Cmp(_dafny.IntOfInt64(286)) < 0) {
						_coll60.Add((_2257_v13).Times(_dafny.IntOfInt64(-415)))
					}
				}
				return _coll60.ToSet()
			}()).IsProperSubsetOf(_2255_v12), globalState)
			_2256_v14 = _out72
		} else {
			var _2258___mcc_h1 bool = _source35.Get_().(D0_DC0).Cf0
			_ = _2258___mcc_h1
			var _2259_cf0 bool = _2258___mcc_h1
			_ = _2259_cf0
			(globalState).F26 = _this.F35()
			var _2260_v15 _dafny.Array
			_ = _2260_v15
			var _len0_64 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_64
			var _nw394 _dafny.Array
			_ = _nw394
			if _len0_64.Cmp(_dafny.Zero) == 0 {
				_nw394 = _dafny.NewArray(_len0_64)
			} else {
				var _init64 func(_dafny.Int) _dafny.Sequence = func(_2261_i1 _dafny.Int) _dafny.Sequence {
					return (_this).F32()
				}
				_ = _init64
				var _element0_64 = _init64(_dafny.Zero)
				_ = _element0_64
				_nw394 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
				(_nw394).ArraySet1(_element0_64, 0)
				var _nativeLen0_64 = (_len0_64).Int()
				_ = _nativeLen0_64
				for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
					(_nw394).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
				}
			}
			_2260_v15 = _nw394
			var _2262_v16 _dafny.Sequence
			_ = _2262_v16
			_2262_v16 = _dafny.UnicodeSeqOfUtf8Bytes("erxeyj")
			var _2263_v17 _dafny.Sequence
			_ = _2263_v17
			_2263_v17 = _dafny.SeqOf((_this).F30(), _dafny.IntOfUint32((_2262_v16).Cardinality()))
			var _2264_v18 _dafny.Set
			_ = _2264_v18
			_2264_v18 = _dafny.SetOf(_dafny.IntOfUint32((_2263_v17).Cardinality()))
			var _rhs365 _dafny.Array = _2260_v15
			_ = _rhs365
			var _rhs366 _dafny.Int = ((((_this).Fm4(_this.F28(), _2264_v18, (_this).F30(), _2259_cf0, globalState)).Union(_2264_v18)).Union((_2264_v18).Difference(_2264_v18))).Cardinality()
			_ = _rhs366
			var _rhs367 D0 = _2241_v1
			_ = _rhs367
			var _rhs368 _dafny.Int = ((_this).F30()).Plus((_this).F34())
			_ = _rhs368
			var _lhs297 *GlobalState = globalState
			_ = _lhs297
			_2260_v15 = _rhs365
			r1 = _rhs366
			_2241_v1 = _rhs367
			_lhs297.F13 = _rhs368
			var _2265_v19 _dafny.Array
			_ = _2265_v19
			var _nw395 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
			_ = _nw395
			_2265_v19 = _nw395
			var _2266_v20 _dafny.Map
			_ = _2266_v20
			_2266_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-720), _dafny.IntOfInt64(212))
			var _2267_v21 _dafny.Map
			_ = _2267_v21
			_2267_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2265_v19, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm9(_this.F36(), (_2266_v20).Cardinality(), (_this).F34(), _this.F35(), globalState), _2263_v17)).Cardinality()))
			_2267_v21 = (_2267_v21).Update((Companion_D2_.Create_DC5_(_2265_v19)).Dtor_cf8(), _dafny.IntOfInt64(206))
			_2259_cf0 = _this.F31()
		}
		var _2268_v22 _dafny.Array
		_ = _2268_v22
		var _len0_65 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_65
		var _nw396 _dafny.Array
		_ = _nw396
		if _len0_65.Cmp(_dafny.Zero) == 0 {
			_nw396 = _dafny.NewArray(_len0_65)
		} else {
			var _init65 func(_dafny.Int) bool = func(_2269_i2 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(866))).Uint32(), func(coer130 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg130 _dafny.Int) interface{} {
						return coer130(arg130)
					}
				}(func(_2270_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				})), _dafny.UnicodeSeqOfUtf8Bytes("fxsini"))
			}
			_ = _init65
			var _element0_65 = _init65(_dafny.Zero)
			_ = _element0_65
			_nw396 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
			(_nw396).ArraySet1(_element0_65, 0)
			var _nativeLen0_65 = (_len0_65).Int()
			_ = _nativeLen0_65
			for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
				(_nw396).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
			}
		}
		_2268_v22 = _nw396
		var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))
		_ = _index332
		(_2268_v22).ArraySet1(false, (_index332).Int())
		var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))
		_ = _index333
		var _rhs369 _dafny.Sequence = func(_source36 D0) _dafny.Sequence {
			if _source36.Is_DC1() {
				var _2271___mcc_h2 _dafny.MultiSet = _source36.Get_().(D0_DC1).Cf1
				_ = _2271___mcc_h2
				var _2272_cf1 _dafny.MultiSet = _2271___mcc_h2
				_ = _2272_cf1
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(627))).Uint32(), func(coer131 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg131 _dafny.Int) interface{} {
						return coer131(arg131)
					}
				}(func(_2273_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))
			} else {
				var _2274___mcc_h3 bool = _source36.Get_().(D0_DC0).Cf0
				_ = _2274___mcc_h3
				var _2275_cf0 bool = _2274___mcc_h3
				_ = _2275_cf0
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("srxfo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(465))).Uint32(), func(coer132 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg132 _dafny.Int) interface{} {
						return coer132(arg132)
					}
				}(func(_2276_i5 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				})))
			}
		}(Companion_D0_.Create_DC0_(_this.F29()))
		_ = _rhs369
		var _rhs370 bool = _this.F31()
		_ = _rhs370
		var _lhs298 _dafny.Array = _2268_v22
		_ = _lhs298
		var _lhs299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))
		_ = _lhs299
		r2 = _rhs369
		(_lhs298).ArraySet1(_rhs370, (_lhs299).Int())
		var _2277_v23 _dafny.Sequence
		_ = _2277_v23
		_2277_v23 = _dafny.SeqOf(Companion_Default___.SafeModuloInt((_this).F34(), _dafny.IntOfInt64(106)), (_this).F34())
		var _2278_v24 _dafny.Sequence
		_ = _2278_v24
		_2278_v24 = _dafny.UnicodeSeqOfUtf8Bytes("p")
		var _rhs371 _dafny.Sequence = _2277_v23
		_ = _rhs371
		var _rhs372 bool = !(!(_this.F28()))
		_ = _rhs372
		var _rhs373 _dafny.Sequence = _2278_v24
		_ = _rhs373
		_2277_v23 = _rhs371
		r0 = _rhs372
		r2 = _rhs373
		var _2279_i6 _dafny.Int
		_ = _2279_i6
		_2279_i6 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(241))).Uint32(), func(coer134 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg134 _dafny.Int) interface{} {
					return coer134(arg134)
				}
			}(func(_2315_i7 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			})), (_this).Fm3((_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool), (_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool), globalState))) {
				{
					if (_2279_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_2279_i6 = (_2279_i6).Plus(_dafny.One)
					var _2280_v25 _dafny.Set
					_ = _2280_v25
					_2280_v25 = _dafny.SetOf(false, _this.F35(), _this.F29())
					_2280_v25 = _dafny.SetOf(_this.F35(), ((_this).F30()).Cmp((_this).F34()) > 0, (_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool), !(_2280_v25).Contains(_this.F28()))
					var _2281_v26 D0
					_ = _2281_v26
					_2281_v26 = Companion_D0_.Create_DC0_(true)
					var _2282_v27 _dafny.Array
					_ = _2282_v27
					var _nw397 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
					_ = _nw397
					_2282_v27 = _nw397
					var _2283_v28 _dafny.Map
					_ = _2283_v28
					_2283_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Array {
						if (_2281_v26).Dtor_cf0() {
							return _2282_v27
						}
						return _2282_v27
					})())
					_2283_v28 = (_2283_v28).Update(_this.F28(), (func() _dafny.Array {
						if (_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool) {
							return _2282_v27
						}
						return (func() _dafny.Array {
							if (_2283_v28).Contains(_this.F35()) {
								return (_2283_v28).Get(_this.F35()).(_dafny.Array)
							}
							return _2282_v27
						})()
					})())
					(globalState).F21 = _2268_v22
					var _2284_v29 _dafny.Array
					_ = _2284_v29
					var _nw398 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
					_ = _nw398
					_2284_v29 = _nw398
					var _2285_v30 *C0
					_ = _2285_v30
					var _nw399 *C0 = New_C0_()
					_ = _nw399
					_nw399.Ctor__(_2284_v29)
					_2285_v30 = _nw399
					var _2286_v31 D1
					_ = _2286_v31
					_2286_v31 = Companion_D1_.Create_DC3_(_this.F35(), ((_this).F30()).Times((_this).F34()), _2285_v30, _dafny.IntOfInt64(929))
					var _source37 D1 = _2286_v31
					_ = _source37
					if _source37.Is_DC3() {
						var _2287___mcc_h4 bool = _source37.Get_().(D1_DC3).Cf3
						_ = _2287___mcc_h4
						var _2288___mcc_h5 _dafny.Int = _source37.Get_().(D1_DC3).Cf4
						_ = _2288___mcc_h5
						var _2289___mcc_h6 *C0 = _source37.Get_().(D1_DC3).Cf5
						_ = _2289___mcc_h6
						var _2290___mcc_h7 _dafny.Int = _source37.Get_().(D1_DC3).Cf6
						_ = _2290___mcc_h7
						var _2291_cf6 _dafny.Int = _2290___mcc_h7
						_ = _2291_cf6
						var _2292_cf5 *C0 = _2289___mcc_h6
						_ = _2292_cf5
						var _2293_cf4 _dafny.Int = _2288___mcc_h5
						_ = _2293_cf4
						var _2294_cf3 bool = _2287___mcc_h4
						_ = _2294_cf3
						var _2295_v32 _dafny.Map
						_ = _2295_v32
						_2295_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F30()).Times((_dafny.Zero).Minus(_2291_cf6)), (_this).F34())
						var _2296_v33 D2
						_ = _2296_v33
						_2296_v33 = Companion_D2_.Create_DC6_((_this).F34(), (_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool), _dafny.IntOfUint32((_2278_v24).Cardinality()), _this.F29(), (_this).F30())
						_2295_v32 = (_2295_v32).Update(_2291_cf6, (_2296_v33).Dtor_cf9())
						(globalState).F1 = (_dafny.SetOf(Companion_Default___.Fm11((_this).F33(), _this.F35(), true, (_this).F34(), globalState), _2280_v25, _2280_v25)).IsProperSubsetOf(Companion_Default___.Fm10(_2280_v25, _dafny.CodePoint('y'), _2293_cf4, !(_this.F31()), globalState))
						r2 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2278_v24, _2278_v24), (Companion_Default___.SafeIndex(_2293_cf4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2278_v24, _2278_v24)).Cardinality()))).Uint32(), _dafny.CodePoint('r'))
						(globalState).F12 = _this.F31()
					} else if _source37.Is_DC2() {
						var _2297___mcc_h8 _dafny.Array = _source37.Get_().(D1_DC2).Cf2
						_ = _2297___mcc_h8
						var _2298_cf2 _dafny.Array = _2297___mcc_h8
						_ = _2298_cf2
						var _2299_v34 D2
						_ = _2299_v34
						_2299_v34 = Companion_D2_.Create_DC6_((_this).F30(), _this.F35(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32(((_this).F32()).Cardinality()), (_this).F30()), (_this).F34())).Cardinality(), !(Companion_Default___.Fm1(true, (_this).F30(), globalState)), (_this).F34())
						var _2300_v35 _dafny.Map
						_ = _2300_v35
						_2300_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), (_this).F30())
						_2299_v34 = Companion_Default___.Fm12(_this.F36(), (_dafny.Zero).Minus((func() _dafny.Int {
							if true {
								return (_this).F30()
							}
							return (_this).F30()
						})()), _2300_v35, _dafny.Companion_Sequence_.IsProperPrefixOf(_2278_v24, _2278_v24), globalState)
						var _2301_v36 _dafny.Sequence
						_ = _2301_v36
						_2301_v36 = _dafny.SeqOf(_2300_v35)
						var _2302_v37 D3
						_ = _2302_v37
						_2302_v37 = Companion_D3_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-446))).Uint32(), func(coer133 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg133 _dafny.Int) interface{} {
								return coer133(arg133)
							}
						}(func(_2303_i8 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), (_dafny.Zero).Minus((_this).F34()))
						})))
						var _2304_v38 _dafny.Array
						_ = _2304_v38
						var _nwElement0_77 _dafny.Sequence = _2301_v36
						_ = _nwElement0_77
						var _nw400 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(6))
						_ = _nw400
						(_nw400).ArraySet1(_nwElement0_77, 0)
						(_nw400).ArraySet1((_2302_v37).Dtor_cf14(), 1)
						(_nw400).ArraySet1(_2301_v36, 2)
						(_nw400).ArraySet1(_2301_v36, 3)
						(_nw400).ArraySet1(_2301_v36, 4)
						(_nw400).ArraySet1(_2301_v36, 5)
						_2304_v38 = _nw400
						var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_2304_v38), 0))
						_ = _index334
						(_2304_v38).ArraySet1(_2301_v36, (_index334).Int())
						var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))
						_ = _index335
						var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_2304_v38), 0))
						_ = _index336
						var _rhs374 bool = ((_this).Fm8(_this.F31(), _dafny.CodePoint('p'), _this.F31(), globalState)).Cmp((_this).F30()) > 0
						_ = _rhs374
						var _rhs375 bool = (_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool)
						_ = _rhs375
						var _rhs376 _dafny.CodePoint = _dafny.CodePoint('d')
						_ = _rhs376
						var _rhs377 _dafny.Int = (_this).F30()
						_ = _rhs377
						var _rhs378 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2301_v36, _2301_v36), _2301_v36)
						_ = _rhs378
						var _lhs300 _dafny.Array = _2268_v22
						_ = _lhs300
						var _lhs301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))
						_ = _lhs301
						var _lhs302 *C10 = _this
						_ = _lhs302
						var _lhs303 *GlobalState = globalState
						_ = _lhs303
						var _lhs304 _dafny.Array = _2304_v38
						_ = _lhs304
						var _lhs305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_2304_v38), 0))
						_ = _lhs305
						r0 = _rhs374
						(_lhs300).ArraySet1(_rhs375, (_lhs301).Int())
						_lhs302.F36_set_(_rhs376)
						_lhs303.F13 = _rhs377
						(_lhs304).ArraySet1(_rhs378, (_lhs305).Int())
						_2300_v35 = (_2300_v35).Update((_2268_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_2268_v22), 0))).Int()).(bool), (_this).F34())
						_2285_v30 = _2285_v30
					} else {
						var _2305___mcc_h9 D1 = _source37.Get_().(D1_DC4).Cf7
						_ = _2305___mcc_h9
						var _2306_cf7 D1 = _2305___mcc_h9
						_ = _2306_cf7
						(globalState).F26 = _this.F31()
						var _2307_v39 _dafny.MultiSet
						_ = _2307_v39
						_2307_v39 = _dafny.MultiSetOf(_this.F36(), _this.F36(), _this.F36())
						var _2308_v40 D3
						_ = _2308_v40
						_2308_v40 = Companion_D3_.Create_DC8_((_2307_v39).Union(_dafny.MultiSetOf(_this.F36(), _this.F36(), _this.F36())), true, (_this).F30())
						var _2309_v41 _dafny.Map
						_ = _2309_v41
						_2309_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F33()).Cardinality()), (_this).F30()), (_this).F34())
						var _2310_v43 _dafny.Map
						_ = _2310_v43
						_2310_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), (_this).F30())
						var _2311_v44 _dafny.Set
						_ = _2311_v44
						_2311_v44 = _dafny.SetOf(_2310_v43)
						var _2312_v46 _dafny.MultiSet
						_ = _2312_v46
						_2312_v46 = _dafny.MultiSetOf(_2280_v25)
						_2308_v40 = Companion_Default___.Fm13(((_2309_v41).Merge(func() _dafny.Map {
							var _coll61 = _dafny.NewMapBuilder()
							_ = _coll61
							for _iter66 := _dafny.Iterate((_2311_v44).Elements()); ; {
								_compr_61, _ok66 := _iter66()
								if !_ok66 {
									break
								}
								var _2313_v42 _dafny.Map
								_2313_v42 = interface{}(_compr_61).(_dafny.Map)
								if (_2311_v44).Contains(_2313_v42) {
									_coll61.Add(_2313_v42, (func() _dafny.Set {
										var _coll62 = _dafny.NewBuilder()
										_ = _coll62
										for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(590), _dafny.IntOfInt64(316))); ; {
											_compr_62, _ok67 := _iter67()
											if !_ok67 {
												break
											}
											var _2314_v45 _dafny.Int
											_2314_v45 = interface{}(_compr_62).(_dafny.Int)
											if ((_dafny.IntOfInt64(590)).Cmp(_2314_v45) <= 0) && ((_2314_v45).Cmp(_dafny.IntOfInt64(316)) < 0) {
												_coll62.Add((_2314_v45).Plus(_dafny.IntOfUint32((_2277_v23).Cardinality())))
											}
										}
										return _coll62.ToSet()
									}()).Cardinality())
								}
							}
							return _coll61.ToMap()
						}())).Cardinality(), (_this).F34(), (_this).F34(), (_2312_v46).Intersection(_2312_v46), globalState)
						(globalState).F13 = (_2277_v23).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_2277_v23).Cardinality()))).Uint32()).(_dafny.Int)
						(globalState).F1 = false
					}
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _2316_i9 _dafny.Int
		_ = _2316_i9
		_2316_i9 = _dafny.Zero
		{
			for _this.F31() {
				{
					if (_2316_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_2316_i9 = (_2316_i9).Plus(_dafny.One)
					(_this).F31_set_(!((_this).Fm7(_dafny.MultiSetOf((_this).F34()), (_this).F30(), globalState)))
					r1 = ((_this).F30()).Plus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _this.F29())).Cardinality()))
					var _2317_v47 _dafny.MultiSet
					_ = _2317_v47
					_2317_v47 = _dafny.MultiSetOf(Companion_Default___.Fm14((_this).F30(), false, _dafny.IntOfInt64(427), Companion_Default___.Fm15((_this).F34(), (_this).F30(), globalState), globalState), _this.F36())
					var _2318_v48 _dafny.Set
					_ = _2318_v48
					_2318_v48 = _dafny.SetOf(_2317_v47)
					_2318_v48 = _dafny.SetOf(_2317_v47)
					var _2319_v49 _dafny.Array
					_ = _2319_v49
					var _nw401 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
					_ = _nw401
					_2319_v49 = _nw401
					var _2320_v50 *C0
					_ = _2320_v50
					var _nw402 *C0 = New_C0_()
					_ = _nw402
					_nw402.Ctor__(_2319_v49)
					_2320_v50 = _nw402
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		r0 = _this.F31()
		r1 = (_this).F34()
		r2 = _dafny.Companion_Sequence_.Concatenate((_this).Fm3(((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), !(true), globalState), _dafny.Companion_Sequence_.Concatenate(_2278_v24, _2278_v24))
		return r0, r1, r2
	}
}
func (_this *C10) M2(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) (bool, D0, D0, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var r2 D0 = Companion_D0_.Default()
		_ = r2
		var r3 bool = false
		_ = r3
		var _2321_v0 _dafny.Array
		_ = _2321_v0
		var _len0_66 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_66
		var _nw403 _dafny.Array
		_ = _nw403
		if _len0_66.Cmp(_dafny.Zero) == 0 {
			_nw403 = _dafny.NewArray(_len0_66)
		} else {
			var _init66 func(_dafny.Int) bool = func(_2322_i0 _dafny.Int) bool {
				return _this.F31()
			}
			_ = _init66
			var _element0_66 = _init66(_dafny.Zero)
			_ = _element0_66
			_nw403 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
			(_nw403).ArraySet1(_element0_66, 0)
			var _nativeLen0_66 = (_len0_66).Int()
			_ = _nativeLen0_66
			for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
				(_nw403).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
			}
		}
		_2321_v0 = _nw403
		var _2323_v1 _dafny.MultiSet
		_ = _2323_v1
		_2323_v1 = _dafny.MultiSetOf(_this.F36(), _this.F36(), _this.F36(), _this.F36())
		var _2324_v2 D3
		_ = _2324_v2
		_2324_v2 = Companion_D3_.Create_DC8_(_2323_v1, _this.F28(), (_this).F34())
		var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))
		_ = _index337
		(_2321_v0).ArraySet1(((_this).F30()).Cmp((_2324_v2).Dtor_cf17()) == 0, (_index337).Int())
		var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))
		_ = _index338
		(_2321_v0).ArraySet1(_this.F29(), (_index338).Int())
		var _2325_i1 _dafny.Int
		_ = _2325_i1
		_2325_i1 = _dafny.Zero
		{
			for _this.F35() {
				{
					if (_2325_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_2325_i1 = (_2325_i1).Plus(_dafny.One)
					var _2326_v3 D3
					_ = _2326_v3
					_2326_v3 = Companion_D3_.Create_DC9_(!(_this.F28()), _this.F35(), _2321_v0, _this.F28())
					var _2327_v4 _dafny.Sequence
					_ = _2327_v4
					_2327_v4 = _dafny.UnicodeSeqOfUtf8Bytes("kddokf")
					var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))
					_ = _index339
					var _rhs379 D3 = Companion_D3_.Create_DC9_(_this.F28(), !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(false, (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), _this.F28(), (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), _this.F35()), !(_this.F35())), _2321_v0, (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool))
					_ = _rhs379
					var _rhs380 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2327_v4, _2327_v4)).Cardinality()), p0)
					_ = _rhs380
					var _rhs381 bool = !(_this.F28()) || (_this.F28())
					_ = _rhs381
					var _rhs382 D3 = _2326_v3
					_ = _rhs382
					var _rhs383 bool = _this.F31()
					_ = _rhs383
					var _lhs306 *GlobalState = globalState
					_ = _lhs306
					var _lhs307 _dafny.Array = _2321_v0
					_ = _lhs307
					var _lhs308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))
					_ = _lhs308
					var _lhs309 *C10 = _this
					_ = _lhs309
					_2326_v3 = _rhs379
					_lhs306.F13 = _rhs380
					(_lhs307).ArraySet1(_rhs381, (_lhs308).Int())
					_2326_v3 = _rhs382
					_lhs309.F31_set_(_rhs383)
					var _2328_v5 _dafny.MultiSet
					_ = _2328_v5
					_2328_v5 = _dafny.MultiSetOf(p2)
					var _2329_v6 _dafny.Map
					_ = _2329_v6
					_2329_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), _dafny.IntOfUint32(((_this).F32()).Cardinality()))
					(globalState).F26 = !((_2328_v5).IsProperSubsetOf(_2328_v5)) || (!(_2329_v6).Equals(_2329_v6))
					var _2330_v7 _dafny.Array
					_ = _2330_v7
					var _len0_67 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_67
					var _nw404 _dafny.Array
					_ = _nw404
					if _len0_67.Cmp(_dafny.Zero) == 0 {
						_nw404 = _dafny.NewArray(_len0_67)
					} else {
						var _init67 func(_dafny.Int) _dafny.Int = (func(_2331_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_2332_i2 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_2332_i2, _2331_p0)
							}
						})(p0)
						_ = _init67
						var _element0_67 = _init67(_dafny.Zero)
						_ = _element0_67
						_nw404 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
						(_nw404).ArraySet1(_element0_67, 0)
						var _nativeLen0_67 = (_len0_67).Int()
						_ = _nativeLen0_67
						for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
							(_nw404).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
						}
					}
					_2330_v7 = _nw404
					_2330_v7 = _2330_v7
					(globalState).F1 = !((func() bool {
						if _this.F28() {
							return _this.F31()
						}
						return (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool)
					})())
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _2333_v8 _dafny.Array
		_ = _2333_v8
		var _nw405 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw405
		_2333_v8 = _nw405
		var _2334_v9 _dafny.Sequence
		_ = _2334_v9
		_2334_v9 = _dafny.SeqOf(_2333_v8)
		var _2335_v10 _dafny.Set
		_ = _2335_v10
		_2335_v10 = _dafny.SetOf((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool))
		var _2336_v11 _dafny.MultiSet
		_ = _2336_v11
		_2336_v11 = _dafny.MultiSetOf(_2335_v10)
		var _2337_v12 _dafny.Array
		_ = _2337_v12
		var _nwElement0_78 _dafny.Array = (_2334_v9).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_2336_v11).Contains(_dafny.SetOf((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool))) {
				return (_2336_v11).Multiplicity(_dafny.SetOf((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool)))
			}
			return (_this).F34()
		})(), _dafny.IntOfUint32((_2334_v9).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _nwElement0_78
		var _nw406 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(28))
		_ = _nw406
		(_nw406).ArraySet1(_nwElement0_78, 0)
		(_nw406).ArraySet1(_2333_v8, 1)
		(_nw406).ArraySet1(_2333_v8, 2)
		(_nw406).ArraySet1(_2333_v8, 3)
		(_nw406).ArraySet1(_2333_v8, 4)
		(_nw406).ArraySet1(_2333_v8, 5)
		(_nw406).ArraySet1(_2333_v8, 6)
		(_nw406).ArraySet1(_2333_v8, 7)
		(_nw406).ArraySet1(_2333_v8, 8)
		(_nw406).ArraySet1((func() _dafny.Array {
			if _this.F31() {
				return _2333_v8
			}
			return _2333_v8
		})(), 9)
		(_nw406).ArraySet1(_2333_v8, 10)
		(_nw406).ArraySet1(_2333_v8, 11)
		(_nw406).ArraySet1(_2333_v8, 12)
		(_nw406).ArraySet1(_2333_v8, 13)
		(_nw406).ArraySet1(_2333_v8, 14)
		(_nw406).ArraySet1(_2333_v8, 15)
		(_nw406).ArraySet1(_2333_v8, 16)
		(_nw406).ArraySet1(_2333_v8, 17)
		(_nw406).ArraySet1(_2333_v8, 18)
		(_nw406).ArraySet1(_2333_v8, 19)
		(_nw406).ArraySet1((func() _dafny.Array {
			if _this.F28() {
				return _2333_v8
			}
			return _2333_v8
		})(), 20)
		(_nw406).ArraySet1(_2333_v8, 21)
		(_nw406).ArraySet1(_2333_v8, 22)
		(_nw406).ArraySet1(_2333_v8, 23)
		(_nw406).ArraySet1(_2333_v8, 24)
		(_nw406).ArraySet1(_2333_v8, 25)
		(_nw406).ArraySet1(_2333_v8, 26)
		(_nw406).ArraySet1(_2333_v8, 27)
		_2337_v12 = _nw406
		var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2337_v12), 0))
		_ = _index340
		(_2337_v12).ArraySet1(_2333_v8, (_index340).Int())
		var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_2337_v12), 0))
		_ = _index341
		(_2337_v12).ArraySet1(_2333_v8, (_index341).Int())
		var _2338_i3 _dafny.Int
		_ = _2338_i3
		_2338_i3 = _dafny.Zero
		{
			for true {
				{
					if (_2338_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_2338_i3 = (_2338_i3).Plus(_dafny.One)
					var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))
					_ = _index342
					(_2321_v0).ArraySet1((((_this).F34()).Minus(p1)).Cmp((_this).F34()) <= 0, (_index342).Int())
					var _2339_v13 _dafny.MultiSet
					_ = _2339_v13
					_2339_v13 = _dafny.MultiSetOf(_this.F35(), _this.F28())
					var _2340_v14 _dafny.Sequence
					_ = _2340_v14
					_2340_v14 = _dafny.SeqOf(_2339_v13)
					var _2341_v15 _dafny.Array
					_ = _2341_v15
					var _nwElement0_79 _dafny.MultiSet = _2339_v13
					_ = _nwElement0_79
					var _nw407 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(14))
					_ = _nw407
					(_nw407).ArraySet1(_nwElement0_79, 0)
					(_nw407).ArraySet1((func() _dafny.MultiSet {
						if _this.F28() {
							return _2339_v13
						}
						return _2339_v13
					})(), 1)
					(_nw407).ArraySet1(_2339_v13, 2)
					(_nw407).ArraySet1(_dafny.MultiSetOf(_this.F31(), ((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), _this.F29(), _this.F28(), _this.F28()), 3)
					(_nw407).ArraySet1(_2339_v13, 4)
					(_nw407).ArraySet1(_2339_v13, 5)
					(_nw407).ArraySet1(_2339_v13, 6)
					(_nw407).ArraySet1((func() _dafny.MultiSet {
						if _this.F31() {
							return _2339_v13
						}
						return _2339_v13
					})(), 7)
					(_nw407).ArraySet1((_2340_v14).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2340_v14).Cardinality()))).Uint32()).(_dafny.MultiSet), 8)
					(_nw407).ArraySet1(_dafny.MultiSetOf(((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)), 9)
					(_nw407).ArraySet1((_2339_v13).Union(_2339_v13), 10)
					(_nw407).ArraySet1(_dafny.MultiSetOf(_this.F35(), _this.F35(), (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), _this.F28(), false), 11)
					(_nw407).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(_this.F31()))), 12)
					(_nw407).ArraySet1(_2339_v13, 13)
					_2341_v15 = _nw407
					var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_2341_v15), 0))
					_ = _index343
					(_2341_v15).ArraySet1((_2339_v13).Update(false, Companion_Default___.Abs((_dafny.Zero).Minus((_this).F30()))), (_index343).Int())
					var _2342_v16 _dafny.Sequence
					_ = _2342_v16
					_2342_v16 = _dafny.UnicodeSeqOfUtf8Bytes("unmvjnend")
					var _2343_v17 _dafny.Map
					_ = _2343_v17
					_2343_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F29(), _2342_v16)
					var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_2341_v15), 0))
					_ = _index344
					(_2341_v15).ArraySet1(_dafny.MultiSetOf(_this.F29(), _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("jfc"), (func() _dafny.Sequence {
						if (_2343_v17).Contains((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool)) {
							return (_2343_v17).Get((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool)).(_dafny.Sequence)
						}
						return _2342_v16
					})())), (_index344).Int())
					var _2344_v18 _dafny.Array
					_ = _2344_v18
					var _nw408 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
					_ = _nw408
					_2344_v18 = _nw408
					var _2345_v19 _dafny.Map
					_ = _2345_v19
					_2345_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _2344_v18)
					var _2346_v20 *C0
					_ = _2346_v20
					var _nw409 *C0 = New_C0_()
					_ = _nw409
					_nw409.Ctor__((func() _dafny.Array {
						if (_2345_v19).Contains(_this.F29()) {
							return (_2345_v19).Get(_this.F29()).(_dafny.Array)
						}
						return _2344_v18
					})())
					_2346_v20 = _nw409
					var _2347_v21 _dafny.Array
					_ = _2347_v21
					var _len0_68 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_68
					var _nw410 _dafny.Array
					_ = _nw410
					if _len0_68.Cmp(_dafny.Zero) == 0 {
						_nw410 = _dafny.NewArray(_len0_68)
					} else {
						var _init68 func(_dafny.Int) _dafny.Sequence = func(_2348_i4 _dafny.Int) _dafny.Sequence {
							return (_this).F32()
						}
						_ = _init68
						var _element0_68 = _init68(_dafny.Zero)
						_ = _element0_68
						_nw410 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
						(_nw410).ArraySet1(_element0_68, 0)
						var _nativeLen0_68 = (_len0_68).Int()
						_ = _nativeLen0_68
						for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
							(_nw410).ArraySet1(_init68(_dafny.IntOf(_i0_68)), _i0_68)
						}
					}
					_2347_v21 = _nw410
					var _2349_v22 D1
					_ = _2349_v22
					_2349_v22 = Companion_D1_.Create_DC3_((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), p0, _2346_v20, (_this).F30())
					var _2350_v23 T4
					_ = _2350_v23
					var _nw411 *C3 = New_C3_()
					_ = _nw411
					_nw411.Ctor__(_dafny.IntOfUint32(((_this).F32()).Cardinality()), _2347_v21, _this.F36(), (_dafny.Zero).Minus((_this).F30()), (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), _dafny.SeqOf(_this.F31(), (_2349_v22).Dtor_cf3()), _dafny.SeqOf(_dafny.SeqOf(_this.F31())), _dafny.IntOfUint32(((_this).F32()).Cardinality()), _this.F35(), _this.F28(), !(_this.F29()))
					_2350_v23 = _nw411
					var _2351_v24 _dafny.Map
					_ = _2351_v24
					_2351_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2350_v23, p1)
					var _2352_v25 _dafny.Map
					_ = _2352_v25
					_2352_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
						if _this.F28() {
							return _this.F36()
						}
						return _this.F36()
					})(), (func() _dafny.Int {
						if (_2351_v24).Contains(_2350_v23) {
							return (_2351_v24).Get(_2350_v23).(_dafny.Int)
						}
						return (_dafny.Zero).Minus(p2)
					})())
					var _2353_v26 _dafny.Sequence
					_ = _2353_v26
					_2353_v26 = _dafny.SeqOf((_this).F30())
					var _2354_v27 _dafny.MultiSet
					_ = _2354_v27
					_2354_v27 = _dafny.MultiSetOf((_2353_v26).Select((Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32((_2353_v26).Cardinality()))).Uint32()).(_dafny.Int), (_2350_v23).F34(), (_2350_v23).F30())
					_2352_v25 = (_2352_v25).Update(Companion_Default___.Fm34((_this).F34(), p0, (func() _dafny.Int {
						if (_2354_v27).Contains(Companion_Default___.Fm46((_2350_v23).F34(), p1, _this.F31(), globalState)) {
							return (_2354_v27).Multiplicity(Companion_Default___.Fm46((_2350_v23).F34(), p1, _this.F31(), globalState))
						}
						return p2
					})(), false, globalState), (_2350_v23).F34())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _2355_v28 _dafny.Array
		_ = _2355_v28
		var _nw412 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw412
		_2355_v28 = _nw412
		var _2356_v29 _dafny.Array
		_ = _2356_v29
		var _nw413 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
		_ = _nw413
		_2356_v29 = _nw413
		var _2357_v30 D22
		_ = _2357_v30
		_2357_v30 = Companion_D22_.Create_DC54_(_2356_v29)
		var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2355_v28), 0))
		_ = _index345
		(_2355_v28).ArraySet1((_2357_v30).Dtor_cf91(), (_index345).Int())
		var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_2355_v28), 0))
		_ = _index346
		(_2355_v28).ArraySet1(_2356_v29, (_index346).Int())
		for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2321_v0), 0))); ; {
			_guard_loop_5, _ok68 := _iter68()
			if !_ok68 {
				break
			}
			var _2358_i5 _dafny.Int
			_2358_i5 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_2358_i5).Sign() != -1) && ((_2358_i5).Cmp(_dafny.ArrayLen((_2321_v0), 0)) < 0)) {
				(_2321_v0).ArraySet1(false, (_2358_i5).Int())
			}
		}
		var _2359_v31 _dafny.Sequence
		_ = _2359_v31
		_2359_v31 = _dafny.UnicodeSeqOfUtf8Bytes("yciqas")
		var _2360_v32 D16
		_ = _2360_v32
		_2360_v32 = Companion_D16_.Create_DC40_(_2359_v31, p1)
		var _2361_v33 D16
		_ = _2361_v33
		_2361_v33 = Companion_D16_.Create_DC41_(_2360_v32)
		var _2362_v34 _dafny.MultiSet
		_ = _2362_v34
		_2362_v34 = _dafny.MultiSetOf(_2361_v33, _2361_v33)
		r0 = ((_dafny.MultiSetOf(_2361_v33)).Update(_2361_v33, Companion_Default___.Abs((_this).F30()))).IsProperSubsetOf((_2362_v34).Update(_2361_v33, Companion_Default___.Abs(p1)))
		var _2363_v35 _dafny.MultiSet
		_ = _2363_v35
		_2363_v35 = _dafny.MultiSetOf(_this.F28(), _this.F28(), _this.F28(), (_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), true)
		r1 = Companion_D0_.Create_DC1_(_2363_v35)
		var _2364_v36 D0
		_ = _2364_v36
		_2364_v36 = Companion_D0_.Create_DC0_(!(!(_2335_v10).Equals(_dafny.SetOf((_2321_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_2321_v0), 0))).Int()).(bool), _this.F31(), _this.F28()))))
		r2 = _2364_v36
		r3 = _this.F28()
		return r0, r1, r2, r3
	}
}
func (_this *C10) M3(p0 bool, p1 bool, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _2365_v0 _dafny.Set
		_ = _2365_v0
		_2365_v0 = _dafny.SetOf(_dafny.IntOfInt64(-512), (_this).F30())
		var _2366_v1 D7
		_ = _2366_v1
		_2366_v1 = Companion_D7_.Create_DC19_(p0, _dafny.IntOfInt64(608), (_this).F30())
		var _2367_v2 _dafny.Array
		_ = _2367_v2
		var _nwElement0_80 _dafny.Int = (_this).F34()
		_ = _nwElement0_80
		var _nw414 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(8))
		_ = _nw414
		(_nw414).ArraySet1(_nwElement0_80, 0)
		(_nw414).ArraySet1((_this).F30(), 1)
		(_nw414).ArraySet1((_2366_v1).Dtor_cf38(), 2)
		(_nw414).ArraySet1((_this).F30(), 3)
		(_nw414).ArraySet1((_this).F34(), 4)
		(_nw414).ArraySet1((func() _dafny.Int {
			if Companion_Default___.Fm1(p0, (_this).F30(), globalState) {
				return (_this).F30()
			}
			return (_this).F34()
		})(), 5)
		(_nw414).ArraySet1((_dafny.Zero).Minus((_this).F30()), 6)
		(_nw414).ArraySet1(_dafny.IntOfInt64(135), 7)
		_2367_v2 = _nw414
		var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
		_ = _index347
		(_2367_v2).ArraySet1((_this).F30(), (_index347).Int())
		var _2368_v3 _dafny.Map
		_ = _2368_v3
		_2368_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))
		var _2369_v4 D5
		_ = _2369_v4
		_2369_v4 = Companion_D5_.Create_DC13_(p0, _2368_v3, _this.F36())
		var _2370_v5 _dafny.MultiSet
		_ = _2370_v5
		_2370_v5 = _dafny.MultiSetOf(_2368_v3, _2368_v3)
		var _2371_v6 _dafny.Sequence
		_ = _2371_v6
		_2371_v6 = _dafny.UnicodeSeqOfUtf8Bytes("vilf")
		var _2372_v7 _dafny.MultiSet
		_ = _2372_v7
		_2372_v7 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2371_v6).Cardinality()))
		var _2373_v8 _dafny.Map
		_ = _2373_v8
		_2373_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), p1)
		var _2374_v9 _dafny.Sequence
		_ = _2374_v9
		_2374_v9 = _dafny.SeqOf((_this).F34(), (_2373_v8).Cardinality())
		var _2375_v10 _dafny.MultiSet
		_ = _2375_v10
		_2375_v10 = _dafny.MultiSetOf(p1)
		var _pat_let_tv26 = _2365_v0
		_ = _pat_let_tv26
		var _pat_let_tv27 = _2365_v0
		_ = _pat_let_tv27
		var _pat_let_tv28 = _2365_v0
		_ = _pat_let_tv28
		var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
		_ = _index348
		var _rhs384 bool = p0
		_ = _rhs384
		var _rhs385 _dafny.Set = func(_source38 D5) _dafny.Set {
			if _source38.Is_DC13() {
				var _2376___mcc_h0 bool = _source38.Get_().(D5_DC13).Cf27
				_ = _2376___mcc_h0
				var _2377___mcc_h1 _dafny.Map = _source38.Get_().(D5_DC13).Cf28
				_ = _2377___mcc_h1
				var _2378___mcc_h2 _dafny.CodePoint = _source38.Get_().(D5_DC13).Cf29
				_ = _2378___mcc_h2
				var _2379_cf29 _dafny.CodePoint = _2378___mcc_h2
				_ = _2379_cf29
				var _2380_cf28 _dafny.Map = _2377___mcc_h1
				_ = _2380_cf28
				var _2381_cf27 bool = _2376___mcc_h0
				_ = _2381_cf27
				return _dafny.SetOf((_this).F30(), _dafny.IntOfInt64(-653))
			} else if _source38.Is_DC14() {
				var _2382___mcc_h3 bool = _source38.Get_().(D5_DC14).Cf30
				_ = _2382___mcc_h3
				var _2383___mcc_h4 _dafny.Int = _source38.Get_().(D5_DC14).Cf31
				_ = _2383___mcc_h4
				var _2384_cf31 _dafny.Int = _2383___mcc_h4
				_ = _2384_cf31
				var _2385_cf30 bool = _2382___mcc_h3
				_ = _2385_cf30
				return _pat_let_tv26
			} else {
				var _2386___mcc_h5 _dafny.Sequence = _source38.Get_().(D5_DC12).Cf26
				_ = _2386___mcc_h5
				var _2387_cf26 _dafny.Sequence = _2386___mcc_h5
				_ = _2387_cf26
				if _this.F35() {
					return _pat_let_tv27
				} else {
					return _pat_let_tv28
				}
			}
		}(_2369_v4)
		_ = _rhs385
		var _rhs386 _dafny.Int = ((_dafny.Zero).Minus(((_2370_v5).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F30())).Update(false, (_this).F34()), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqOf((_this).Fm7(Companion_Default___.Fm49(_2372_v7, _this.F36(), _dafny.SeqOf(_2374_v9), globalState), (_this).F30(), globalState))).Cardinality())))).Cardinality())).Plus((_2375_v10).Cardinality())
		_ = _rhs386
		var _lhs310 *C10 = _this
		_ = _lhs310
		var _lhs311 _dafny.Array = _2367_v2
		_ = _lhs311
		var _lhs312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
		_ = _lhs312
		_lhs310.F31_set_(_rhs384)
		_2365_v0 = _rhs385
		(_lhs311).ArraySet1(_rhs386, (_lhs312).Int())
		if !(p0) {
			var _2388_v11 _dafny.Map
			_ = _2388_v11
			_2388_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _2374_v9)
			var _2389_v12 _dafny.Sequence
			_ = _2389_v12
			_2389_v12 = _dafny.SeqOf(((_2388_v11).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2374_v9, (Companion_Default___.SafeIndex((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2374_v9).Cardinality()))).Uint32(), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Cardinality())) > 0, _this.F29())
			_2389_v12 = Companion_Default___.Fm0(((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)).Minus((_this).F34()), _this.F31(), p0, ((_this).F34()).Plus(_dafny.IntOfUint32(((_this).Fm3(_this.F29(), _this.F29(), globalState)).Cardinality())), globalState)
			(_this).F29_set_(_this.F35())
			var _2390_v13 _dafny.Sequence
			_ = _2390_v13
			_2390_v13 = _dafny.SeqOf(_2365_v0)
			_2365_v0 = ((_2390_v13).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_this.F28(), _this.F28())).Cardinality(), _dafny.IntOfUint32((_2390_v13).Cardinality()))).Uint32()).(_dafny.Set)).Intersection((_2365_v0).Union(_2365_v0))
			var _2391_v14 _dafny.MultiSet
			_ = _2391_v14
			_2391_v14 = _dafny.MultiSetOf((_2375_v10).Update(p0, Companion_Default___.Abs(_dafny.IntOfInt64(178))))
			var _2392_v15 _dafny.MultiSet
			_ = _2392_v15
			_2392_v15 = _2391_v14
			var _source39 _dafny.MultiSet = _2392_v15
			_ = _source39
			var _2393___mcc_h6 _dafny.MultiSet = _source39
			_ = _2393___mcc_h6
			var _2394_cf80 _dafny.MultiSet = _2393___mcc_h6
			_ = _2394_cf80
			var _2395_v16 _dafny.Array
			_ = _2395_v16
			var _nw415 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(22))
			_ = _nw415
			_2395_v16 = _nw415
			var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_2395_v16), 0))
			_ = _index349
			(_2395_v16).ArraySet1(_2375_v10, (_index349).Int())
			var _2396_v17 _dafny.Sequence
			_ = _2396_v17
			_2396_v17 = _dafny.SeqOf(_2375_v10, _2375_v10)
			var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_2395_v16), 0))
			_ = _index350
			(_2395_v16).ArraySet1((_2396_v17).Select((Companion_Default___.SafeIndex(((_this).F30()).Minus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_2396_v17).Cardinality()))).Uint32()).(_dafny.MultiSet), (_index350).Int())
			(_this).F36_set_((func() _dafny.CodePoint {
				if _this.F29() {
					return _this.F36()
				}
				return _this.F36()
			})())
			var _2397_v18 _dafny.Array
			_ = _2397_v18
			var _nw416 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(21))
			_ = _nw416
			_2397_v18 = _nw416
			var _2398_v19 _dafny.Array
			_ = _2398_v19
			_2398_v19 = _2397_v18
			_2398_v19 = _2398_v19
			(globalState).F12 = p1
			var _2399_v20 _dafny.Array
			_ = _2399_v20
			var _len0_69 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_69
			var _nw417 _dafny.Array
			_ = _nw417
			if _len0_69.Cmp(_dafny.Zero) == 0 {
				_nw417 = _dafny.NewArray(_len0_69)
			} else {
				var _init69 func(_dafny.Int) bool = func(_2400_i0 _dafny.Int) bool {
					return _this.F29()
				}
				_ = _init69
				var _element0_69 = _init69(_dafny.Zero)
				_ = _element0_69
				_nw417 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
				(_nw417).ArraySet1(_element0_69, 0)
				var _nativeLen0_69 = (_len0_69).Int()
				_ = _nativeLen0_69
				for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
					(_nw417).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
				}
			}
			_2399_v20 = _nw417
			var _2401_v21 D6
			_ = _2401_v21
			_2401_v21 = Companion_D6_.Create_DC16_((Companion_D3_.Create_DC9_(Companion_Default___.Fm1(p0, (_this).F30(), globalState), Companion_Default___.Fm1(_this.F35(), (_this).F30(), globalState), _2399_v20, false)).Dtor_cf20())
			var _2402_v22 D20
			_ = _2402_v22
			_2402_v22 = Companion_D20_.Create_DC50_(Companion_D20_.Create_DC49_(_this.F35(), _this.F36()))
			var _2403_v23 D3
			_ = _2403_v23
			_2403_v23 = Companion_D3_.Create_DC9_(Companion_Default___.Fm1(p1, (_this).F34(), globalState), false, _2399_v20, p0)
			var _2404_v24 _dafny.Sequence
			_ = _2404_v24
			_2404_v24 = _dafny.SeqOf((_2403_v23).Dtor_cf20())
			var _2405_v25 *C1
			_ = _2405_v25
			var _nw418 *C1 = New_C1_()
			_ = _nw418
			_nw418.Ctor__(_this.F28(), p1, p1)
			_2405_v25 = _nw418
			var _2406_v26 D10
			_ = _2406_v26
			_2406_v26 = Companion_D10_.Create_DC28_(_2405_v25)
			var _pat_let_tv29 = _2405_v25
			_ = _pat_let_tv29
			var _pat_let_tv30 = _2405_v25
			_ = _pat_let_tv30
			var _2407_v27 _dafny.Array
			_ = _2407_v27
			var _nwElement0_81 D10 = _2406_v26
			_ = _nwElement0_81
			var _nw419 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(21))
			_ = _nw419
			(_nw419).ArraySet1(_nwElement0_81, 0)
			(_nw419).ArraySet1(_2406_v26, 1)
			(_nw419).ArraySet1(Companion_D10_.Create_DC28_(_2405_v25), 2)
			(_nw419).ArraySet1(_2406_v26, 3)
			(_nw419).ArraySet1(_2406_v26, 4)
			(_nw419).ArraySet1(_2406_v26, 5)
			(_nw419).ArraySet1(_2406_v26, 6)
			(_nw419).ArraySet1(_2406_v26, 7)
			(_nw419).ArraySet1(Companion_D10_.Create_DC28_(_2405_v25), 8)
			(_nw419).ArraySet1(_2406_v26, 9)
			(_nw419).ArraySet1(Companion_D10_.Create_DC28_(_2405_v25), 10)
			(_nw419).ArraySet1(_2406_v26, 11)
			(_nw419).ArraySet1(_2406_v26, 12)
			(_nw419).ArraySet1(func(_pat_let37_0 D10) D10 {
				return func(_2408_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let38_0 *C1) D10 {
						return func(_2409_dt__update_hcf55_h0 *C1) D10 {
							return Companion_D10_.Create_DC28_(_2409_dt__update_hcf55_h0)
						}(_pat_let38_0)
					}(_pat_let_tv29)
				}(_pat_let37_0)
			}(_2406_v26), 13)
			(_nw419).ArraySet1(func(_pat_let39_0 D10) D10 {
				return func(_2410_dt__update__tmp_h1 D10) D10 {
					return func(_pat_let40_0 *C1) D10 {
						return func(_2411_dt__update_hcf55_h1 *C1) D10 {
							return Companion_D10_.Create_DC28_(_2411_dt__update_hcf55_h1)
						}(_pat_let40_0)
					}(_pat_let_tv30)
				}(_pat_let39_0)
			}(_2406_v26), 14)
			(_nw419).ArraySet1(_2406_v26, 15)
			(_nw419).ArraySet1(_2406_v26, 16)
			(_nw419).ArraySet1(_2406_v26, 17)
			(_nw419).ArraySet1(_2406_v26, 18)
			(_nw419).ArraySet1(_2406_v26, 19)
			(_nw419).ArraySet1(Companion_D10_.Create_DC28_(_2405_v25), 20)
			_2407_v27 = _nw419
			var _2412_v28 D20
			_ = _2412_v28
			_2412_v28 = Companion_D20_.Create_DC48_((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _2407_v27, _this.F28())
			var _2413_v29 D20
			_ = _2413_v29
			_2413_v29 = Companion_D20_.Create_DC50_(_2412_v28)
			var _pat_let_tv31 = _2404_v24
			_ = _pat_let_tv31
			var _pat_let_tv32 = _2404_v24
			_ = _pat_let_tv32
			var _rhs387 D6 = func(_pat_let41_0 D6) D6 {
				return func(_2414_dt__update__tmp_h2 D6) D6 {
					return func(_pat_let42_0 _dafny.Array) D6 {
						return func(_2415_dt__update_hcf33_h0 _dafny.Array) D6 {
							return Companion_D6_.Create_DC16_(_2415_dt__update_hcf33_h0)
						}(_pat_let42_0)
					}((_pat_let_tv31).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_pat_let_tv32).Cardinality()))).Uint32()).(_dafny.Array))
				}(_pat_let41_0)
			}((func() D6 {
				if p0 {
					return _2401_v21
				}
				return _2401_v21
			})())
			_ = _rhs387
			var _rhs388 bool = _dafny.Companion_Sequence_.Contains(_2371_v6, _this.F36())
			_ = _rhs388
			var _rhs389 D20 = (func() D20 {
				if true {
					return Companion_D20_.Create_DC50_(_2413_v29)
				}
				return _2402_v22
			})()
			_ = _rhs389
			var _rhs390 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F34(), _dafny.IntOfInt64(-651))
			_ = _rhs390
			var _lhs313 *GlobalState = globalState
			_ = _lhs313
			_2401_v21 = _rhs387
			_lhs313.F26 = _rhs388
			_2402_v22 = _rhs389
			r1 = _rhs390
		} else {
			var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
			_ = _index351
			var _rhs391 _dafny.Sequence = _2371_v6
			_ = _rhs391
			var _rhs392 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)))
			_ = _rhs392
			var _lhs314 _dafny.Array = _2367_v2
			_ = _lhs314
			var _lhs315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
			_ = _lhs315
			_2371_v6 = _rhs391
			(_lhs314).ArraySet1(_rhs392, (_lhs315).Int())
			var _2416_v30 _dafny.Map
			_ = _2416_v30
			_2416_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _dafny.IntOfInt64(717))
			var _2417_v31 _dafny.Map
			_ = _2417_v31
			_2417_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2416_v30, _2371_v6)
			(globalState).F1 = ((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.SafeModuloInt((_2417_v31).Cardinality(), (_this).F34())) >= 0
			if _this.F35() {
				_2374_v9 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2374_v9, _2374_v9), _dafny.Companion_Sequence_.Concatenate(_2374_v9, _2374_v9))
				var _2418_v32 _dafny.Array
				_ = _2418_v32
				var _len0_70 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_70
				var _nw420 _dafny.Array
				_ = _nw420
				if _len0_70.Cmp(_dafny.Zero) == 0 {
					_nw420 = _dafny.NewArray(_len0_70)
				} else {
					var _init70 func(_dafny.Int) bool = func(_2419_i1 _dafny.Int) bool {
						return true
					}
					_ = _init70
					var _element0_70 = _init70(_dafny.Zero)
					_ = _element0_70
					_nw420 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
					(_nw420).ArraySet1(_element0_70, 0)
					var _nativeLen0_70 = (_len0_70).Int()
					_ = _nativeLen0_70
					for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
						(_nw420).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
					}
				}
				_2418_v32 = _nw420
				var _2420_v33 _dafny.Map
				_ = _2420_v33
				_2420_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F36(), _2373_v8)
				var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_2418_v32), 0))
				_ = _index352
				(_2418_v32).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm0((_this).F34(), true, p0, (_this).F34(), globalState), Companion_Default___.Fm0(((func() _dafny.Map {
					if (_2420_v33).Contains(_dafny.CodePoint('e')) {
						return (_2420_v33).Get(_dafny.CodePoint('e')).(_dafny.Map)
					}
					return _2373_v8
				})()).Cardinality(), false, p0, (_this).F30(), globalState)), (_index352).Int())
				var _2421_v34 _dafny.Map
				_ = _2421_v34
				_2421_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), _this.F36())
				var _2422_v35 _dafny.Array
				_ = _2422_v35
				var _len0_71 _dafny.Int = _dafny.One
				_ = _len0_71
				var _nw421 _dafny.Array
				_ = _nw421
				if _len0_71.Cmp(_dafny.Zero) == 0 {
					_nw421 = _dafny.NewArray(_len0_71)
				} else {
					var _init71 func(_dafny.Int) _dafny.CodePoint = func(_2423_i2 _dafny.Int) _dafny.CodePoint {
						return _this.F36()
					}
					_ = _init71
					var _element0_71 = _init71(_dafny.Zero)
					_ = _element0_71
					_nw421 = _dafny.NewArrayFromExample(_element0_71, nil, _len0_71)
					(_nw421).ArraySet1CodePoint(_element0_71, 0)
					var _nativeLen0_71 = (_len0_71).Int()
					_ = _nativeLen0_71
					for _i0_71 := 1; _i0_71 < _nativeLen0_71; _i0_71++ {
						(_nw421).ArraySet1CodePoint(_init71(_dafny.IntOf(_i0_71)), _i0_71)
					}
				}
				_2422_v35 = _nw421
				var _2424_v36 *C0
				_ = _2424_v36
				var _nw422 *C0 = New_C0_()
				_ = _nw422
				_nw422.Ctor__(_2422_v35)
				_2424_v36 = _nw422
				var _2425_v37 D9
				_ = _2425_v37
				_2425_v37 = Companion_D9_.Create_DC27_(_2424_v36, _2371_v6)
				var _2426_v38 _dafny.Map
				_ = _2426_v38
				_2426_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2425_v37, _this.F29())
				var _2427_v39 *C6
				_ = _2427_v39
				var _nw423 *C6 = New_C6_()
				_ = _nw423
				_nw423.Ctor__((func() _dafny.CodePoint {
					if (_2421_v34).Contains(Companion_Default___.Fm46((_2426_v38).Cardinality(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _this.F28(), globalState)) {
						return (_2421_v34).Get(Companion_Default___.Fm46((_2426_v38).Cardinality(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _this.F28(), globalState)).(_dafny.CodePoint)
					}
					return _this.F36()
				})(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), !(!(p1)), (_this).F32(), (_this).F33(), (_this).F30(), p0, p0, _this.F35())
				_2427_v39 = _nw423
				var _2428_v40 _dafny.Set
				_ = _2428_v40
				_2428_v40 = _dafny.SetOf(_2427_v39, _2427_v39)
				var _2429_v41 _dafny.Map
				_ = _2429_v41
				_2429_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_2428_v40).Cardinality(), (_this).F30()), p0)
				var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_2418_v32), 0))
				_ = _index353
				(_2418_v32).ArraySet1((func() bool {
					if (_2429_v41).Contains((_this).F34()) {
						return (_2429_v41).Get((_this).F34()).(bool)
					}
					return p0
				})(), (_index353).Int())
				(globalState).F21 = _2418_v32
				r2 = (_dafny.Zero).Minus((_this).F30())
				var _2430_v42 _dafny.Array
				_ = _2430_v42
				var _nw424 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
				_ = _nw424
				_2430_v42 = _nw424
				var _2431_v43 _dafny.Map
				_ = _2431_v43
				_2431_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2430_v42)
				_2431_v43 = (_2431_v43).Update(_this.F29(), _2430_v42)
			} else {
				var _2432_v44 _dafny.Array
				_ = _2432_v44
				var _len0_72 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_72
				var _nw425 _dafny.Array
				_ = _nw425
				if _len0_72.Cmp(_dafny.Zero) == 0 {
					_nw425 = _dafny.NewArray(_len0_72)
				} else {
					var _init72 func(_dafny.Int) bool = func(_2433_i3 _dafny.Int) bool {
						return _this.F28()
					}
					_ = _init72
					var _element0_72 = _init72(_dafny.Zero)
					_ = _element0_72
					_nw425 = _dafny.NewArrayFromExample(_element0_72, nil, _len0_72)
					(_nw425).ArraySet1(_element0_72, 0)
					var _nativeLen0_72 = (_len0_72).Int()
					_ = _nativeLen0_72
					for _i0_72 := 1; _i0_72 < _nativeLen0_72; _i0_72++ {
						(_nw425).ArraySet1(_init72(_dafny.IntOf(_i0_72)), _i0_72)
					}
				}
				_2432_v44 = _nw425
				var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_2432_v44), 0))
				_ = _index354
				(_2432_v44).ArraySet1(_this.F28(), (_index354).Int())
				var _2434_v45 _dafny.Sequence
				_ = _2434_v45
				_2434_v45 = _dafny.SeqOf(_2367_v2)
				var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_2432_v44), 0))
				_ = _index355
				var _rhs393 bool = Companion_Default___.Fm1((_this.F28()) || (((_this).F32()).Select((Companion_Default___.SafeIndex((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool)), _dafny.IntOfInt64(96), globalState)
				_ = _rhs393
				var _rhs394 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_2434_v45, _2434_v45)
				_ = _rhs394
				var _rhs395 bool = Companion_Default___.Fm1(_this.F29(), Companion_Default___.SafeModuloInt((_this).F34(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), globalState)
				_ = _rhs395
				var _rhs396 bool = _this.F31()
				_ = _rhs396
				var _rhs397 _dafny.Int = (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)
				_ = _rhs397
				var _lhs316 *GlobalState = globalState
				_ = _lhs316
				var _lhs317 _dafny.Array = _2432_v44
				_ = _lhs317
				var _lhs318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_2432_v44), 0))
				_ = _lhs318
				var _lhs319 *GlobalState = globalState
				_ = _lhs319
				var _lhs320 *C10 = _this
				_ = _lhs320
				var _lhs321 *GlobalState = globalState
				_ = _lhs321
				_lhs316.F26 = _rhs393
				(_lhs317).ArraySet1(_rhs394, (_lhs318).Int())
				_lhs319.F26 = _rhs395
				_lhs320.F28_set_(_rhs396)
				_lhs321.F13 = _rhs397
				var _2435_v47 _dafny.Set
				_ = _2435_v47
				_2435_v47 = func() _dafny.Set {
					var _coll63 = _dafny.NewBuilder()
					_ = _coll63
					for _iter69 := _dafny.Iterate((_2365_v0).Elements()); ; {
						_compr_63, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _2436_v46 _dafny.Int
						_2436_v46 = interface{}(_compr_63).(_dafny.Int)
						if (_2365_v0).Contains(_2436_v46) {
							_coll63.Add(Companion_Default___.SafeModuloInt(_2436_v46, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(672), _dafny.IntOfUint32((_dafny.SeqOf(!(!(true)))).Cardinality()), _dafny.IntOfInt64(866))).Cardinality())))
						}
					}
					return _coll63.ToSet()
				}()
				var _2437_v48 _dafny.Array
				_ = _2437_v48
				var _nwElement0_82 _dafny.Set = _2435_v47
				_ = _nwElement0_82
				var _nw426 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(7))
				_ = _nw426
				(_nw426).ArraySet1(_nwElement0_82, 0)
				(_nw426).ArraySet1(_2435_v47, 1)
				(_nw426).ArraySet1(_2365_v0, 2)
				(_nw426).ArraySet1(_2435_v47, 3)
				(_nw426).ArraySet1((func() _dafny.Set {
					if false {
						return _2435_v47
					}
					return _2435_v47
				})(), 4)
				(_nw426).ArraySet1(_2435_v47, 5)
				(_nw426).ArraySet1(_2435_v47, 6)
				_2437_v48 = _nw426
				var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_2437_v48), 0))
				_ = _index356
				(_2437_v48).ArraySet1(_2435_v47, (_index356).Int())
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_2437_v48), 0))
				_ = _index357
				var _rhs398 _dafny.Int = _dafny.IntOfInt64(292)
				_ = _rhs398
				var _rhs399 _dafny.Sequence = _2371_v6
				_ = _rhs399
				var _rhs400 bool = !(true)
				_ = _rhs400
				var _rhs401 _dafny.Set = Companion_Default___.Fm63((_this.F28()) && ((_2432_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_2432_v44), 0))).Int()).(bool)), p1, p0, globalState)
				_ = _rhs401
				var _lhs322 *GlobalState = globalState
				_ = _lhs322
				var _lhs323 *GlobalState = globalState
				_ = _lhs323
				var _lhs324 _dafny.Array = _2437_v48
				_ = _lhs324
				var _lhs325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_2437_v48), 0))
				_ = _lhs325
				_lhs322.F13 = _rhs398
				_2371_v6 = _rhs399
				_lhs323.F1 = _rhs400
				(_lhs324).ArraySet1(_rhs401, (_lhs325).Int())
				var _2438_v49 _dafny.Sequence
				_ = _2438_v49
				_2438_v49 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate((_this).F32(), (_this).F32()))
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _index358
				var _rhs402 _dafny.Int = _dafny.IntOfInt64(759)
				_ = _rhs402
				var _rhs403 _dafny.Int = ((_this).F30()).Plus(((_this).F30()).Times((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)))
				_ = _rhs403
				var _rhs404 _dafny.Int = (_this).F34()
				_ = _rhs404
				var _rhs405 _dafny.CodePoint = _this.F36()
				_ = _rhs405
				var _rhs406 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2438_v49, (_this).F33())
				_ = _rhs406
				var _lhs326 *GlobalState = globalState
				_ = _lhs326
				var _lhs327 _dafny.Array = _2367_v2
				_ = _lhs327
				var _lhs328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _lhs328
				var _lhs329 *C10 = _this
				_ = _lhs329
				r1 = _rhs402
				_lhs326.F13 = _rhs403
				(_lhs327).ArraySet1(_rhs404, (_lhs328).Int())
				_lhs329.F36_set_(_rhs405)
				_2438_v49 = _rhs406
				var _2439_v50 _dafny.Sequence
				_ = _2439_v50
				_2439_v50 = _dafny.SeqOf(true)
				_2439_v50 = _2439_v50
				var _2440_v51 D4
				_ = _2440_v51
				_2440_v51 = Companion_D4_.Create_DC11_((_2373_v8).Cardinality(), (_2432_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_2432_v44), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tuwhsf")).Cardinality()))
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _index359
				(_2367_v2).ArraySet1((((_2440_v51).Dtor_cf25()).Minus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))).Plus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), (_index359).Int())
			}
			var _2441_v52 D0
			_ = _2441_v52
			_2441_v52 = Companion_D0_.Create_DC0_(p1)
			var _2442_v53 _dafny.Set
			_ = _2442_v53
			_2442_v53 = _dafny.SetOf(p0, (_2441_v52).Dtor_cf0())
			var _2443_v54 _dafny.Sequence
			_ = _2443_v54
			_2443_v54 = _dafny.SeqOf(_2442_v53)
			var _2444_v55 bool
			_ = _2444_v55
			var _2445_v56 D0
			_ = _2445_v56
			var _2446_v57 D0
			_ = _2446_v57
			var _2447_v58 bool
			_ = _2447_v58
			var _out73 bool
			_ = _out73
			var _out74 D0
			_ = _out74
			var _out75 D0
			_ = _out75
			var _out76 bool
			_ = _out76
			_out73, _out74, _out75, _out76 = (_this).M2((_this).F30(), (_this).F34(), (_this).F30(), _2443_v54, globalState)
			_2444_v55 = _out73
			_2445_v56 = _out74
			_2446_v57 = _out75
			_2447_v58 = _out76
			var _2448_v59 _dafny.Array
			_ = _2448_v59
			var _nw427 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw427
			_2448_v59 = _nw427
			var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_2448_v59), 0))
			_ = _index360
			(_2448_v59).ArraySet1(_this.F28(), (_index360).Int())
			var _2449_v60 _dafny.Map
			_ = _2449_v60
			_2449_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F34(), (func() _dafny.Int {
				if (_2375_v10).Contains(_this.F28()) {
					return (_2375_v10).Multiplicity(_this.F28())
				}
				return (_this).F34()
			})())
			var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_2448_v59), 0))
			_ = _index361
			(_2448_v59).ArraySet1(((_this).F32()).Select((Companion_Default___.SafeIndex(((_2449_v60).Cardinality()).Plus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), (_index361).Int())
		}
		if p1 {
			r0 = p0
			if _this.F31() {
				var _2450_v61 _dafny.Array
				_ = _2450_v61
				var _len0_73 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_73
				var _nw428 _dafny.Array
				_ = _nw428
				if _len0_73.Cmp(_dafny.Zero) == 0 {
					_nw428 = _dafny.NewArray(_len0_73)
				} else {
					var _init73 func(_dafny.Int) bool = func(_2451_i4 _dafny.Int) bool {
						return (_this.F28()) == (_this.F29())
					}
					_ = _init73
					var _element0_73 = _init73(_dafny.Zero)
					_ = _element0_73
					_nw428 = _dafny.NewArrayFromExample(_element0_73, nil, _len0_73)
					(_nw428).ArraySet1(_element0_73, 0)
					var _nativeLen0_73 = (_len0_73).Int()
					_ = _nativeLen0_73
					for _i0_73 := 1; _i0_73 < _nativeLen0_73; _i0_73++ {
						(_nw428).ArraySet1(_init73(_dafny.IntOf(_i0_73)), _i0_73)
					}
				}
				_2450_v61 = _nw428
				(globalState).F21 = _2450_v61
				var _2452_v62 D16
				_ = _2452_v62
				_2452_v62 = Companion_D16_.Create_DC40_(_2371_v6, (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))
				var _2453_v63 T1
				_ = _2453_v63
				var _nw429 *C7 = New_C7_()
				_ = _nw429
				_nw429.Ctor__((_this).F34(), (_2372_v7).Intersection(_2372_v7), (_this).F30(), _dafny.Companion_Sequence_.IsProperPrefixOf((_2452_v62).Dtor_cf73(), _2371_v6), _this.F29(), _this.F28())
				_2453_v63 = _nw429
				_2453_v63 = _2453_v63
				var _2454_v64 _dafny.MultiSet
				_ = _2454_v64
				_2454_v64 = _dafny.MultiSetOf(_dafny.SeqOf(_this.F36()))
				var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_2450_v61), 0))
				_ = _index362
				(_2450_v61).ArraySet1((_2454_v64).IsProperSubsetOf(_2454_v64), (_index362).Int())
				var _2455_v65 D3
				_ = _2455_v65
				_2455_v65 = Companion_D3_.Create_DC9_(_2453_v63.F28(), _this.F31(), _2450_v61, p1)
				var _2456_v66 _dafny.Set
				_ = _2456_v66
				_2456_v66 = _dafny.SetOf(!(p1), _this.F28(), !(_2453_v63.F28()))
				var _2457_v67 _dafny.Map
				_ = _2457_v67
				_2457_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _this.F31())
				var _2458_v69 _dafny.Set
				_ = _2458_v69
				_2458_v69 = _dafny.SetOf(_2457_v67, func() _dafny.Map {
					var _coll64 = _dafny.NewMapBuilder()
					_ = _coll64
					for _iter70 := _dafny.Iterate((_dafny.SetOf((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))).Elements()); ; {
						_compr_64, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _2459_v68 _dafny.Int
						_2459_v68 = interface{}(_compr_64).(_dafny.Int)
						if (_dafny.SetOf((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))).Contains(_2459_v68) {
							_coll64.Add((_2459_v68).Times((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), _2453_v63.F28())
						}
					}
					return _coll64.ToMap()
				}())
				var _2460_v70 _dafny.Sequence
				_ = _2460_v70
				_2460_v70 = _dafny.SeqOf(!(_dafny.SetOf(_this.F35(), false)).Equals(_2456_v66), (_2458_v69).IsDisjointFrom(_2458_v69), !(true))
				var _2461_v71 _dafny.Array
				_ = _2461_v71
				var _nwElement0_83 _dafny.CodePoint = _this.F36()
				_ = _nwElement0_83
				var _nw430 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(8))
				_ = _nw430
				(_nw430).ArraySet1CodePoint(_nwElement0_83, 0)
				(_nw430).ArraySet1CodePoint(_dafny.CodePoint('a'), 1)
				(_nw430).ArraySet1CodePoint(_this.F36(), 2)
				(_nw430).ArraySet1CodePoint(_dafny.CodePoint('s'), 3)
				(_nw430).ArraySet1CodePoint(_this.F36(), 4)
				(_nw430).ArraySet1CodePoint(_this.F36(), 5)
				(_nw430).ArraySet1CodePoint((_2371_v6).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), 6)
				(_nw430).ArraySet1CodePoint(_this.F36(), 7)
				_2461_v71 = _nw430
				var _2462_v72 _dafny.Sequence
				_ = _2462_v72
				_2462_v72 = _dafny.SeqOf((_this).F33(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(414))).Uint32(), func(coer135 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg135 _dafny.Int) interface{} {
						return coer135(arg135)
					}
				}(func(_2463_i5 _dafny.Int) _dafny.Sequence {
					return (_this).F32()
				})))
				var _2464_v73 T3
				_ = _2464_v73
				var _nw431 *C5 = New_C5_()
				_ = _nw431
				_nw431.Ctor__(_2453_v63.F29(), (_this).F34(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), p1, (_this).F32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer136 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg136 _dafny.Int) interface{} {
						return coer136(arg136)
					}
				}(func(_2465_i6 _dafny.Int) _dafny.Sequence {
					return (_this).F32()
				})), (_2453_v63).F30(), !(p0), _this.F31(), p1)
				_2464_v73 = _nw431
				var _2466_v74 _dafny.Map
				_ = _2466_v74
				_2466_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2464_v73, (_2368_v3).Cardinality())
				var _2467_v75 T3
				_ = _2467_v75
				var _nw432 *C5 = New_C5_()
				_ = _nw432
				_nw432.Ctor__(_this.F35(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), (_2453_v63).F30(), _this.F35(), _2460_v70, (_2462_v72).Select((Companion_Default___.SafeIndex((_2466_v74).Cardinality(), _dafny.IntOfUint32((_2462_v72).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.IntOfUint32((_2460_v70).Cardinality()), _2464_v73.F28(), p0, p1)
				_2467_v75 = _nw432
				var _2468_v76 _dafny.Map
				_ = _2468_v76
				_2468_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2461_v71, _2467_v75)
				var _pat_let_tv33 = _2450_v61
				_ = _pat_let_tv33
				var _pat_let_tv34 = _2450_v61
				_ = _pat_let_tv34
				var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_2450_v61), 0))
				_ = _index363
				var _rhs407 bool = (_this.F29()) || (p0)
				_ = _rhs407
				var _rhs408 D3 = func(_pat_let43_0 D3) D3 {
					return func(_2469_dt__update__tmp_h3 D3) D3 {
						return func(_pat_let44_0 _dafny.Array) D3 {
							return func(_2470_dt__update_hcf20_h0 _dafny.Array) D3 {
								return Companion_D3_.Create_DC9_((_2469_dt__update__tmp_h3).Dtor_cf18(), (_2469_dt__update__tmp_h3).Dtor_cf19(), _2470_dt__update_hcf20_h0, (_2469_dt__update__tmp_h3).Dtor_cf21())
							}(_pat_let44_0)
						}((func() _dafny.Array {
							if _this.F29() {
								return _pat_let_tv33
							}
							return _pat_let_tv34
						})())
					}(_pat_let43_0)
				}(_2455_v65)
				_ = _rhs408
				var _rhs409 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F34(), (_2468_v76).Cardinality())
				_ = _rhs409
				var _rhs410 _dafny.Sequence = (_2467_v75).F32()
				_ = _rhs410
				var _lhs330 _dafny.Array = _2450_v61
				_ = _lhs330
				var _lhs331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_2450_v61), 0))
				_ = _lhs331
				(_lhs330).ArraySet1(_rhs407, (_lhs331).Int())
				_2455_v65 = _rhs408
				r2 = _rhs409
				_2460_v70 = _rhs410
				var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _index364
				(_2367_v2).ArraySet1(((_2456_v66).Intersection((_dafny.SetOf((_2450_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_2450_v61), 0))).Int()).(bool), _2453_v63.F29(), _2453_v63.F31(), _2453_v63.F29(), (_this).Fm7((_2372_v7).Update((_dafny.Zero).Minus((_this).F34()), Companion_Default___.Abs(_dafny.IntOfInt64(154))), (_2467_v75).F34(), globalState))).Difference(_2456_v66))).Cardinality(), (_index364).Int())
				var _2471_v77 *C4
				_ = _2471_v77
				var _nw433 *C4 = New_C4_()
				_ = _nw433
				_nw433.Ctor__((_2467_v75).F30(), _this.F28(), (_2467_v75).F32(), (_2464_v73).F33(), (_2457_v67).Cardinality(), _2464_v73.F31(), _2467_v75.F31(), _this.F35())
				_2471_v77 = _nw433
				var _2472_v78 _dafny.MultiSet
				_ = _2472_v78
				_2472_v78 = _dafny.MultiSetOf(_2471_v77, _2471_v77)
				var _2473_v79 _dafny.Array
				_ = _2473_v79
				var _nw434 _dafny.Array = _dafny.NewArrayWithValue(Companion_D10_.Default(), _dafny.IntOfInt64(6))
				_ = _nw434
				_2473_v79 = _nw434
				var _2474_v80 D20
				_ = _2474_v80
				_2474_v80 = Companion_D20_.Create_DC48_((_2472_v78).Cardinality(), _2473_v79, _2467_v75.F29())
				(globalState).F13 = (_2474_v80).Dtor_cf82()
			} else {
				var _2475_v81 _dafny.Map
				_ = _2475_v81
				_2475_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2375_v10, p1)
				var _2476_v82 _dafny.Array
				_ = _2476_v82
				var _len0_74 _dafny.Int = _dafny.One
				_ = _len0_74
				var _nw435 _dafny.Array
				_ = _nw435
				if _len0_74.Cmp(_dafny.Zero) == 0 {
					_nw435 = _dafny.NewArray(_len0_74)
				} else {
					var _init74 func(_dafny.Int) _dafny.Sequence = func(_2477_i7 _dafny.Int) _dafny.Sequence {
						return (_this).F32()
					}
					_ = _init74
					var _element0_74 = _init74(_dafny.Zero)
					_ = _element0_74
					_nw435 = _dafny.NewArrayFromExample(_element0_74, nil, _len0_74)
					(_nw435).ArraySet1(_element0_74, 0)
					var _nativeLen0_74 = (_len0_74).Int()
					_ = _nativeLen0_74
					for _i0_74 := 1; _i0_74 < _nativeLen0_74; _i0_74++ {
						(_nw435).ArraySet1(_init74(_dafny.IntOf(_i0_74)), _i0_74)
					}
				}
				_2476_v82 = _nw435
				var _2478_v83 _dafny.Set
				_ = _2478_v83
				_2478_v83 = _dafny.SetOf((func() bool {
					if (_2373_v8).Contains(p0) {
						return (_2373_v8).Get(p0).(bool)
					}
					return _this.F29()
				})(), _this.F28())
				var _2479_v84 *C3
				_ = _2479_v84
				var _nw436 *C3 = New_C3_()
				_ = _nw436
				_nw436.Ctor__((_2478_v83).Cardinality(), _2476_v82, _this.F36(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), false, (_this).F32(), (_this).F33(), (_2365_v0).Cardinality(), Companion_Default___.Fm1(p1, (_this).F30(), globalState), true, p1)
				_2479_v84 = _nw436
				var _2480_v85 _dafny.Map
				_ = _2480_v85
				_2480_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _2479_v84)
				var _2481_v86 _dafny.Sequence
				_ = _2481_v86
				_2481_v86 = _dafny.SeqOf(_2479_v84, (func() *C3 {
					if (_2480_v85).Contains((_this).F30()) {
						return (_2480_v85).Get((_this).F30()).(*C3)
					}
					return _2479_v84
				})())
				var _2482_v87 _dafny.Map
				_ = _2482_v87
				_2482_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _this.F36())).Cardinality(), (_this).F34())
				var _2483_v88 *C3
				_ = _2483_v88
				var _nw437 *C3 = New_C3_()
				_ = _nw437
				_nw437.Ctor__((_2475_v81).Cardinality(), _2476_v82, _this.F36(), ((_this).F34()).Minus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), (_this).Fm7(_dafny.MultiSetOf((_dafny.Zero).Minus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))), _dafny.IntOfUint32((_2481_v86).Cardinality()), globalState), (_this).F32(), (_this).F33(), (func() _dafny.Int {
					if (_2482_v87).Contains((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)) {
						return (_2482_v87).Get((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)).(_dafny.Int)
					}
					return (_2479_v84).F44()
				})(), _this.F28(), ((_this).F34()).Cmp((_2479_v84).F44()) <= 0, (_2375_v10).IsProperSubsetOf(_2375_v10))
				_2483_v88 = _nw437
				var _2484_v89 *C4
				_ = _2484_v89
				var _nw438 *C4 = New_C4_()
				_ = _nw438
				_nw438.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))).Cardinality()), (_2479_v84).F44()), (_dafny.IntOfInt64(213)).Cmp(_dafny.IntOfInt64(452)) <= 0, (_this).F32(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F32(), (_this).F32(), (_this).F32()), (_this).F33()), (_2479_v84).F44(), !(false) || (p1), _this.F28(), p1)
				_2484_v89 = _nw438
				var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _index365
				(_2367_v2).ArraySet1((_dafny.Zero).Minus((_2479_v84).F44()), (_index365).Int())
				(globalState).F26 = (((_this).F30()).Cmp((_this).F34()) == 0) == ((_2484_v89).Fm6(globalState))
				var _2485_v90 _dafny.Sequence
				_ = _2485_v90
				_2485_v90 = _dafny.SeqOf(_2372_v7, _2372_v7)
				var _2486_v91 T1
				_ = _2486_v91
				var _nw439 *C7 = New_C7_()
				_ = _nw439
				_nw439.Ctor__((_2479_v84).F44(), (_2485_v90).Select((Companion_Default___.SafeIndex((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2485_v90).Cardinality()))).Uint32()).(_dafny.MultiSet), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2371_v6, (Companion_Default___.SafeIndex((_2479_v84).F44(), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Uint32(), _this.F36())).Cardinality()), _this.F28(), p1, (_2484_v89).Fm44((_this).F30(), globalState))
				_2486_v91 = _nw439
				var _nw440 *C7 = New_C7_()
				_ = _nw440
				_nw440.Ctor__((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_2486_v91).F30(), (_2483_v88).F44())), Companion_Default___.Fm25(globalState), (_2486_v91).F30(), true, ((_2479_v84).F44()).Cmp((_2483_v88).F44()) != 0, _2486_v91.F28())
				_2486_v91 = _nw440
			}
			_2373_v8 = (_2373_v8).Update(p0, !(!(_this.F31())))
			(globalState).F13 = (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)
			var _2487_v92 D0
			_ = _2487_v92
			_2487_v92 = Companion_D0_.Create_DC1_(_2375_v10)
			var _2488_v93 _dafny.Map
			_ = _2488_v93
			_2488_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _2487_v92)
			var _2489_v94 _dafny.Set
			_ = _2489_v94
			_2489_v94 = _dafny.SetOf(_this.F28())
			var _2490_v96 _dafny.Array
			_ = _2490_v96
			var _nwElement0_84 _dafny.Map = (_2488_v93).Update((_2489_v94).Cardinality(), _2487_v92)
			_ = _nwElement0_84
			var _nw441 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(15))
			_ = _nw441
			(_nw441).ArraySet1(_nwElement0_84, 0)
			(_nw441).ArraySet1(_2488_v93, 1)
			(_nw441).ArraySet1(_2488_v93, 2)
			(_nw441).ArraySet1(_2488_v93, 3)
			(_nw441).ArraySet1(Companion_Default___.Fm64((_this).F30(), _this.F35(), (_dafny.Zero).Minus((_this).F34()), _this.F28(), globalState), 4)
			(_nw441).ArraySet1((func() _dafny.Map {
				if _this.F31() {
					return _2488_v93
				}
				return _2488_v93
			})(), 5)
			(_nw441).ArraySet1(_2488_v93, 6)
			(_nw441).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F34())), _2487_v92), 7)
			(_nw441).ArraySet1(_2488_v93, 8)
			(_nw441).ArraySet1(_2488_v93, 9)
			(_nw441).ArraySet1(func() _dafny.Map {
				var _coll65 = _dafny.NewMapBuilder()
				_ = _coll65
				for _iter71 := _dafny.Iterate((_2374_v9).Elements()); ; {
					_compr_65, _ok71 := _iter71()
					if !_ok71 {
						break
					}
					var _2491_v95 _dafny.Int
					_2491_v95 = interface{}(_compr_65).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_2374_v9, _2491_v95) {
						_coll65.Add((_2491_v95).Minus((_this).F30()), Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_this.F28())))
					}
				}
				return _coll65.ToMap()
			}(), 10)
			(_nw441).ArraySet1((_2488_v93).Merge(_2488_v93), 11)
			(_nw441).ArraySet1(_2488_v93, 12)
			(_nw441).ArraySet1((_2488_v93).Merge(_2488_v93), 13)
			(_nw441).ArraySet1((_2488_v93).Merge(_2488_v93), 14)
			_2490_v96 = _nw441
			var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_2490_v96), 0))
			_ = _index366
			(_2490_v96).ArraySet1(_2488_v93, (_index366).Int())
			var _2492_v97 D8
			_ = _2492_v97
			_2492_v97 = Companion_D8_.Create_DC22_((_this).F30(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2371_v6, (Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Uint32(), _this.F36())).Cardinality()))
			var _2493_v98 _dafny.Map
			_ = _2493_v98
			_2493_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F28(), _2492_v97)
			var _2494_v100 _dafny.Map
			_ = _2494_v100
			_2494_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), p0)
			var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_2490_v96), 0))
			_ = _index367
			var _rhs411 _dafny.Int = (_this).F34()
			_ = _rhs411
			var _rhs412 _dafny.Map = (_2488_v93).Merge(func() _dafny.Map {
				var _coll66 = _dafny.NewMapBuilder()
				_ = _coll66
				for _iter72 := _dafny.Iterate((_2494_v100).Keys().Elements()); ; {
					_compr_66, _ok72 := _iter72()
					if !_ok72 {
						break
					}
					var _2495_v99 _dafny.Int
					_2495_v99 = interface{}(_compr_66).(_dafny.Int)
					if (_2494_v100).Contains(_2495_v99) {
						_coll66.Add(Companion_Default___.SafeDivisionInt(_2495_v99, (_this).F30()), _2487_v92)
					}
				}
				return _coll66.ToMap()
			}())
			_ = _rhs412
			var _rhs413 _dafny.Int = (_this).F34()
			_ = _rhs413
			var _rhs414 _dafny.Map = (func() _dafny.Map {
				if p1 {
					return _2493_v98
				}
				return (Companion_Default___.Fm65((_this).F34(), _this.F35(), (_this).F34(), globalState)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), _2492_v97)).Update(_this.F35(), Companion_D8_.Create_DC22_((_this).F30(), (_this).F34())))
			})()
			_ = _rhs414
			var _rhs415 bool = _this.F28()
			_ = _rhs415
			var _lhs332 _dafny.Array = _2490_v96
			_ = _lhs332
			var _lhs333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_2490_v96), 0))
			_ = _lhs333
			var _lhs334 *GlobalState = globalState
			_ = _lhs334
			var _lhs335 *GlobalState = globalState
			_ = _lhs335
			r1 = _rhs411
			(_lhs332).ArraySet1(_rhs412, (_lhs333).Int())
			_lhs334.F13 = _rhs413
			_2493_v98 = _rhs414
			_lhs335.F1 = _rhs415
		} else {
			if _this.F28() {
				(globalState).F13 = (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)
				(_this).F36_set_((_2371_v6).Select((Companion_Default___.SafeIndex((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _len0_75 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_75
				var _nw442 _dafny.Array
				_ = _nw442
				if _len0_75.Cmp(_dafny.Zero) == 0 {
					_nw442 = _dafny.NewArray(_len0_75)
				} else {
					var _init75 func(_dafny.Int) bool = (func(_2496_p1 bool) func(_dafny.Int) bool {
						return func(_2497_i8 _dafny.Int) bool {
							return _2496_p1
						}
					})(p1)
					_ = _init75
					var _element0_75 = _init75(_dafny.Zero)
					_ = _element0_75
					_nw442 = _dafny.NewArrayFromExample(_element0_75, nil, _len0_75)
					(_nw442).ArraySet1(_element0_75, 0)
					var _nativeLen0_75 = (_len0_75).Int()
					_ = _nativeLen0_75
					for _i0_75 := 1; _i0_75 < _nativeLen0_75; _i0_75++ {
						(_nw442).ArraySet1(_init75(_dafny.IntOf(_i0_75)), _i0_75)
					}
				}
				(globalState).F21 = _nw442
				var _2498_v101 D0
				_ = _2498_v101
				_2498_v101 = Companion_D0_.Create_DC1_(_2375_v10)
				var _2499_v102 _dafny.Map
				_ = _2499_v102
				_2499_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _this.F31())
				var _2500_v103 _dafny.Sequence
				_ = _2500_v103
				_2500_v103 = _dafny.SeqOf(_2499_v102)
				var _2501_v104 _dafny.Set
				_ = _2501_v104
				_2501_v104 = _dafny.SetOf(false)
				var _2502_v105 D2
				_ = _2502_v105
				_2502_v105 = Companion_D2_.Create_DC6_((_this).F34(), _this.F35(), (_2501_v104).Cardinality(), _this.F35(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F31(), (_this).F30())).Cardinality())
				var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _index368
				var _rhs416 D0 = Companion_Default___.Fm15(_dafny.IntOfUint32((_2374_v9).Cardinality()), _dafny.IntOfUint32((_2500_v103).Cardinality()), globalState)
				_ = _rhs416
				var _rhs417 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F35(), (_2502_v105).Dtor_cf10())).Merge(_2373_v8)
				_ = _rhs417
				var _rhs418 _dafny.Int = _dafny.IntOfInt64(121)
				_ = _rhs418
				var _rhs419 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2371_v6, _2371_v6), _2371_v6)
				_ = _rhs419
				var _lhs336 _dafny.Array = _2367_v2
				_ = _lhs336
				var _lhs337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
				_ = _lhs337
				_2498_v101 = _rhs416
				_2373_v8 = _rhs417
				(_lhs336).ArraySet1(_rhs418, (_lhs337).Int())
				r3 = _rhs419
				(_this).F36_set_(_this.F36())
			} else {
				var _2503_v106 _dafny.Map
				_ = _2503_v106
				_2503_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2367_v2, (_this).Fm4(p0, _2365_v0, (_this).F30(), p1, globalState))
				var _2504_v107 _dafny.Map
				_ = _2504_v107
				_2504_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _this.F31())
				var _2505_v108 T2
				_ = _2505_v108
				var _nw443 *C8 = New_C8_()
				_ = _nw443
				_nw443.Ctor__(p1, _this.F36(), _dafny.IntOfUint32((_2371_v6).Cardinality()), _this.F31(), (_this).F32(), (_this).F33(), _dafny.IntOfUint32((_2374_v9).Cardinality()), false, _this.F28(), _this.F35())
				_2505_v108 = _nw443
				var _2506_v109 _dafny.MultiSet
				_ = _2506_v109
				_2506_v109 = _dafny.MultiSetOf(_2505_v108, _2505_v108, _2505_v108, _2505_v108)
				var _2507_v110 T1
				_ = _2507_v110
				var _nw444 *C7 = New_C7_()
				_ = _nw444
				_nw444.Ctor__(((_2503_v106).Cardinality()).Minus((_2372_v7).Cardinality()), (_2372_v7).Intersection(_2372_v7), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-618))).Uint32(), func(coer137 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg137 _dafny.Int) interface{} {
						return coer137(arg137)
					}
				}(func(_2508_i9 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}))).Cardinality())), (func() bool {
					if (_2504_v107).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2371_v6, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F30()), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Uint32(), _this.F36())).Cardinality())) {
						return (_2504_v107).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2371_v6, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F30()), _dafny.IntOfUint32((_2371_v6).Cardinality()))).Uint32(), _this.F36())).Cardinality())).(bool)
					}
					return _this.F28()
				})(), _dafny.Companion_Sequence_.Contains((_this).F32(), _this.F31()), (_2506_v109).IsDisjointFrom(_2506_v109))
				_2507_v110 = _nw444
				_2507_v110 = _2507_v110
				var _2509_v111 _dafny.Array
				_ = _2509_v111
				var _len0_76 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_76
				var _nw445 _dafny.Array
				_ = _nw445
				if _len0_76.Cmp(_dafny.Zero) == 0 {
					_nw445 = _dafny.NewArray(_len0_76)
				} else {
					var _init76 func(_dafny.Int) bool = func(_2510_i10 _dafny.Int) bool {
						return true
					}
					_ = _init76
					var _element0_76 = _init76(_dafny.Zero)
					_ = _element0_76
					_nw445 = _dafny.NewArrayFromExample(_element0_76, nil, _len0_76)
					(_nw445).ArraySet1(_element0_76, 0)
					var _nativeLen0_76 = (_len0_76).Int()
					_ = _nativeLen0_76
					for _i0_76 := 1; _i0_76 < _nativeLen0_76; _i0_76++ {
						(_nw445).ArraySet1(_init76(_dafny.IntOf(_i0_76)), _i0_76)
					}
				}
				_2509_v111 = _nw445
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2509_v111), 0))
				_ = _index369
				(_2509_v111).ArraySet1(p1, (_index369).Int())
				var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2509_v111), 0))
				_ = _index370
				(_2509_v111).ArraySet1(((_2507_v110.F31()) || (!(_2507_v110.F28()))) && (_2507_v110.F31()), (_index370).Int())
				var _2511_v112 *C4
				_ = _2511_v112
				var _nw446 *C4 = New_C4_()
				_ = _nw446
				_nw446.Ctor__(Companion_Default___.SafeDivisionInt((Companion_D10_.Create_DC31_(_2505_v108.F29(), _this.F28(), (_2507_v110).F30())).Dtor_cf62(), _dafny.IntOfInt64(754)), true, (_this).F32(), (_2505_v108).F33(), (_2505_v108).F30(), (_2509_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2509_v111), 0))).Int()).(bool), (_2370_v5).Contains(_2368_v3), !(_this.F31()))
				_2511_v112 = _nw446
				var _2512_v113 _dafny.Sequence
				_ = _2512_v113
				_2512_v113 = _dafny.SeqOf((_dafny.MultiSetOf((_2509_v111).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_2509_v111), 0))).Int()).(bool))).Update(_2507_v110.F31(), Companion_Default___.Abs(Companion_Default___.Fm24(_2507_v110.F31(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), globalState))), (_dafny.MultiSetOf(_2505_v108.F28(), _2507_v110.F29())).Intersection((_dafny.MultiSetOf(_this.F28())).Update(_2505_v108.F28(), Companion_Default___.Abs((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)))), (_2375_v10).Update(_2507_v110.F29(), Companion_Default___.Abs((_2505_v108).F30())))
				_2512_v113 = _dafny.SeqOf((func() _dafny.MultiSet {
					if _2507_v110.F29() {
						return _2375_v10
					}
					return (_2375_v10).Update(false, Companion_Default___.Abs((_this).F30()))
				})())
				var _2513_v114 T3
				_ = _2513_v114
				var _nw447 *C4 = New_C4_()
				_ = _nw447
				_nw447.Ctor__(Companion_Default___.SafeDivisionInt((_2505_v108).F30(), (_this).F30()), p1, (_this).F32(), (_this).F33(), (_dafny.Zero).Minus((_this).F30()), _2505_v108.F28(), _this.F29(), false)
				_2513_v114 = _nw447
				var _rhs420 T3 = _2513_v114
				_ = _rhs420
				var _rhs421 _dafny.Sequence = _2371_v6
				_ = _rhs421
				_2513_v114 = _rhs420
				_2371_v6 = _rhs421
			}
			var _2514_v115 _dafny.Array
			_ = _2514_v115
			var _nw448 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw448
			_2514_v115 = _nw448
			var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_2514_v115), 0))
			_ = _index371
			(_2514_v115).ArraySet1(_2371_v6, (_index371).Int())
			var _2515_v116 D16
			_ = _2515_v116
			_2515_v116 = Companion_D16_.Create_DC40_((_this).Fm3(_this.F29(), p0, globalState), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int))
			var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_2514_v115), 0))
			_ = _index372
			(_2514_v115).ArraySet1((_2515_v116).Dtor_cf73(), (_index372).Int())
			(globalState).F1 = _this.F31()
			var _2516_v117 _dafny.Array
			_ = _2516_v117
			var _nw449 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw449
			_2516_v117 = _nw449
			var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_2516_v117), 0))
			_ = _index373
			(_2516_v117).ArraySet1(((_this).F32()).Select((Companion_Default___.SafeIndex((_this).F34(), _dafny.IntOfUint32(((_this).F32()).Cardinality()))).Uint32()).(bool), (_index373).Int())
			var _2517_v118 _dafny.Array
			_ = _2517_v118
			var _len0_77 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_77
			var _nw450 _dafny.Array
			_ = _nw450
			if _len0_77.Cmp(_dafny.Zero) == 0 {
				_nw450 = _dafny.NewArray(_len0_77)
			} else {
				var _init77 func(_dafny.Int) _dafny.CodePoint = func(_2518_i11 _dafny.Int) _dafny.CodePoint {
					return _this.F36()
				}
				_ = _init77
				var _element0_77 = _init77(_dafny.Zero)
				_ = _element0_77
				_nw450 = _dafny.NewArrayFromExample(_element0_77, nil, _len0_77)
				(_nw450).ArraySet1CodePoint(_element0_77, 0)
				var _nativeLen0_77 = (_len0_77).Int()
				_ = _nativeLen0_77
				for _i0_77 := 1; _i0_77 < _nativeLen0_77; _i0_77++ {
					(_nw450).ArraySet1CodePoint(_init77(_dafny.IntOf(_i0_77)), _i0_77)
				}
			}
			_2517_v118 = _nw450
			var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_2517_v118), 0))
			_ = _index374
			(_2517_v118).ArraySet1CodePoint(_this.F36(), (_index374).Int())
			var _2519_v119 D0
			_ = _2519_v119
			_2519_v119 = Companion_D0_.Create_DC1_(_2375_v10)
			var _2520_v120 _dafny.Map
			_ = _2520_v120
			_2520_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-393), _this.F28())
			var _pat_let_tv35 = _2375_v10
			_ = _pat_let_tv35
			var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_2516_v117), 0))
			_ = _index375
			var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_2517_v118), 0))
			_ = _index376
			var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
			_ = _index377
			var _rhs422 bool = _this.F28()
			_ = _rhs422
			var _rhs423 _dafny.CodePoint = Companion_Default___.Fm14((_this).F34(), !(_this.F29()) || (p0), Companion_Default___.Fm39(_this.F28(), _dafny.MultiSetOf((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), _this.F35(), globalState), func(_pat_let45_0 D0) D0 {
				return func(_2521_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let46_0 _dafny.MultiSet) D0 {
						return func(_2522_dt__update_hcf1_h0 _dafny.MultiSet) D0 {
							return Companion_D0_.Create_DC1_(_2522_dt__update_hcf1_h0)
						}(_pat_let46_0)
					}(_pat_let_tv35)
				}(_pat_let45_0)
			}(_2519_v119), globalState)
			_ = _rhs423
			var _rhs424 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer138 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg138 _dafny.Int) interface{} {
					return coer138(arg138)
				}
			}(func(_2523_i12 _dafny.Int) _dafny.CodePoint {
				return _this.F36()
			})), (_2514_v115).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_2514_v115), 0))).Int()).(_dafny.Sequence))).Cardinality())).Plus((_this).F34())
			_ = _rhs424
			var _rhs425 _dafny.Int = (_this).F34()
			_ = _rhs425
			var _rhs426 _dafny.Int = ((((_2520_v120).Update(_dafny.IntOfInt64(743), false)).Update((_dafny.Zero).Minus((_this).F30()), _this.F35())).Cardinality()).Plus((_this).F34())
			_ = _rhs426
			var _lhs338 _dafny.Array = _2516_v117
			_ = _lhs338
			var _lhs339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_2516_v117), 0))
			_ = _lhs339
			var _lhs340 _dafny.Array = _2517_v118
			_ = _lhs340
			var _lhs341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.ArrayLen((_2517_v118), 0))
			_ = _lhs341
			var _lhs342 _dafny.Array = _2367_v2
			_ = _lhs342
			var _lhs343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))
			_ = _lhs343
			(_lhs338).ArraySet1(_rhs422, (_lhs339).Int())
			(_lhs340).ArraySet1CodePoint(_rhs423, (_lhs341).Int())
			(_lhs342).ArraySet1(_rhs424, (_lhs343).Int())
			r1 = _rhs425
			r1 = _rhs426
			_2367_v2 = _2367_v2
		}
		var _2524_v121 *C5
		_ = _2524_v121
		var _nw451 *C5 = New_C5_()
		_ = _nw451
		_nw451.Ctor__(p1, (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
			if (_2372_v7).Contains((_this).F34()) {
				return (_2372_v7).Multiplicity((_this).F34())
			}
			return _dafny.IntOfInt64(444)
		})(), _this.F28(), (_this).F32(), (_this).F33(), (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), _this.F31(), _this.F31(), !(_this.F29()))
		_2524_v121 = _nw451
		var _2525_v122 _dafny.Set
		_ = _2525_v122
		_2525_v122 = _2365_v0
		_2525_v122 = (func() _dafny.Set {
			if !(p1) {
				return _2525_v122
			}
			return _2525_v122
		})()
		var _2526_v123 _dafny.Map
		_ = _2526_v123
		_2526_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int), (_2373_v8).Update(p0, p0))
		var _2527_v124 _dafny.Array
		_ = _2527_v124
		var _nw452 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
		_ = _nw452
		_2527_v124 = _nw452
		var _2528_v125 _dafny.Map
		_ = _2528_v125
		_2528_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2527_v124, (_this).F33())
		var _2529_v126 D14
		_ = _2529_v126
		_2529_v126 = Companion_D14_.Create_DC36_((_this).F33())
		var _2530_v127 _dafny.Set
		_ = _2530_v127
		_2530_v127 = _dafny.SetOf(p1)
		var _2531_v128 *C9
		_ = _2531_v128
		var _nw453 *C9 = New_C9_()
		_ = _nw453
		_nw453.Ctor__((_dafny.Zero).Minus((_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)), _2526_v123, _this.F36(), _dafny.IntOfUint32(((_this).F32()).Cardinality()), ((_this).F34()).Cmp((_this).F30()) <= 0, (_this).F32(), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_2528_v125).Contains(_2527_v124) {
				return (_2528_v125).Get(_2527_v124).(_dafny.Sequence)
			}
			return (_2529_v126).Dtor_cf67()
		})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer139 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg139 _dafny.Int) interface{} {
				return coer139(arg139)
			}
		}(func(_2532_i13 _dafny.Int) _dafny.Sequence {
			return (_this).F32()
		}))), (_this).F30(), (_2524_v121).F47(), (_2530_v127).IsSubsetOf(Companion_Default___.Fm11((_this).F33(), p1, (_2524_v121).F47(), _dafny.IntOfInt64(195), globalState)), _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_2366_v1), _2366_v1))
		_2531_v128 = _nw453
		r0 = (_2524_v121).Fm6(globalState)
		r1 = (func() _dafny.Int {
			if (_2368_v3).Contains((_2524_v121).F47()) {
				return (_2368_v3).Get((_2524_v121).F47()).(_dafny.Int)
			}
			return (_dafny.IntOfUint32(((_this).F32()).Cardinality())).Minus((_this).F34())
		})()
		r2 = (_2367_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_2367_v2), 0))).Int()).(_dafny.Int)
		var _2533_v129 _dafny.Set
		_ = _2533_v129
		_2533_v129 = _dafny.SetOf(_2367_v2, _2367_v2)
		r3 = (_this).Fm7(_dafny.MultiSetFromSeq(_2374_v9), ((_2533_v129).Difference(_2533_v129)).Cardinality(), globalState)
		return r0, r1, r2, r3
	}
}

// End of class C10
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
