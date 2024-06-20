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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.IntOfInt64(864)).Cmp(((_dafny.SetOf(_dafny.IntOfInt64(572))).Cardinality()).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-483))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()))) != 0
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Map, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(587), false)).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(587), false)).Contains(_1_v0) {
				_coll0.Add((_1_v0).Minus((_dafny.MultiSetOf(!(true))).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}()).Union(func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(182), _dafny.IntOfInt64(877))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(182)).Cmp(_2_v1) <= 0) && ((_2_v1).Cmp(_dafny.IntOfInt64(877)) < 0) {
				_coll1.Add((_2_v1).Plus((_dafny.SetOf(false)).Cardinality()))
			}
		}
		return _coll1.ToSet()
	}())).Intersection(_dafny.SetOf(_dafny.IntOfInt64(-216)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt(((_dafny.SetOf(false)).Intersection(_dafny.SetOf(false))).Cardinality(), _dafny.IntOfInt64(81))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("nargnu")
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D2 = Companion_D2_.Create_DC2_(_dafny.CodePoint('a'))
	_ = _source0
	if _source0.Is_DC3() {
		var _3___mcc_h0 bool = _source0.Get_().(D2_DC3).Cf3
		_ = _3___mcc_h0
		var _4___mcc_h1 bool = _source0.Get_().(D2_DC3).Cf4
		_ = _4___mcc_h1
		var _5_cf4 bool = _4___mcc_h1
		_ = _5_cf4
		var _6_cf3 bool = _3___mcc_h0
		_ = _6_cf3
		return _dafny.CodePoint('x')
	} else if _source0.Is_DC4() {
		var _7___mcc_h2 _dafny.Int = _source0.Get_().(D2_DC4).Cf5
		_ = _7___mcc_h2
		var _8___mcc_h3 _dafny.CodePoint = _source0.Get_().(D2_DC4).Cf6
		_ = _8___mcc_h3
		var _9_cf6 _dafny.CodePoint = _8___mcc_h3
		_ = _9_cf6
		var _10_cf5 _dafny.Int = _7___mcc_h2
		_ = _10_cf5
		return (Companion_D2_.Create_DC2_(_dafny.CodePoint('l'))).Dtor_cf2()
	} else {
		var _11___mcc_h4 _dafny.CodePoint = _source0.Get_().(D2_DC2).Cf2
		_ = _11___mcc_h4
		var _12_cf2 _dafny.CodePoint = _11___mcc_h4
		_ = _12_cf2
		return _12_cf2
	}
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(false), true, !(true))).Difference(_dafny.MultiSetOf(false, false))).Difference((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, true))).Difference(_dafny.MultiSetOf(false, false, !(!(!(true))))))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-644), _dafny.IntOfInt64(320))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-380))).Cardinality()))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-569), _dafny.IntOfInt64(-889))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(432), _dafny.IntOfInt64(210))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_13_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(531)
	})), _dafny.SeqOf(_dafny.IntOfInt64(238))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(779))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_14_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(209)
	})))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(827))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality(), _dafny.SetOf(Companion_D2_.Create_DC3_(false, true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-47), func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC3_(!(!(false)), true), _dafny.SetOf(false))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _15_v0 D2
			_15_v0 = interface{}(_compr_2).(D2)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC3_(!(!(false)), true), _dafny.SetOf(false))).Contains(_15_v0) {
				_coll2.Add(_15_v0)
			}
		}
		return _coll2.ToSet()
	}()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-487), _dafny.SetOf(Companion_D2_.Create_DC3_(false, true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-654), !(true))).Cardinality(), _dafny.SetOf(Companion_D2_.Create_DC3_(true, true)))))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_16_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(493))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_17_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))
	}))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("vwsaa"), _dafny.UnicodeSeqOfUtf8Bytes("dmyeoj")), !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), false)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(808), _dafny.IntOfInt64(-796))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _18_v0 _dafny.Int
			_18_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(808)).Cmp(_18_v0) <= 0) && ((_18_v0).Cmp(_dafny.IntOfInt64(-796)) < 0) {
				_coll3.Add((_18_v0).Plus(_dafny.IntOfInt64(10)))
			}
		}
		return _coll3.ToSet()
	}()).Cardinality()), _dafny.IntOfInt64(754), _dafny.IntOfInt64(462), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tkpurq")).Cardinality()))).Cardinality()), !(false))), !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfInt64(898), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("agfup")).Cardinality()))).Cardinality()), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-752))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.MultiSetOf((_dafny.SetOf(!(true))).Cardinality(), (Companion_D4_.Create_DC7_(_dafny.IntOfInt64(102), true)).Dtor_cf9()))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(86), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(741))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, globalState *GlobalState) D2 {
	if !(((Companion_D8_.Create_DC17_(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(977)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(101))))).Dtor_cf28()).IsDisjointFrom(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(857), _dafny.IntOfInt64(355))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _19_v0 _dafny.Int
			_19_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(857)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(355)) < 0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_19_v0, _dafny.IntOfInt64(886)), _dafny.SeqOf(true))
			}
		}
		return _coll4.ToMap()
	}()).Cardinality())))) {
		return Companion_D2_.Create_DC3_(!(true), false)
	} else {
		return Companion_D2_.Create_DC3_(!(true), false)
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Set {
	if true {
		return (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-551))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})), _dafny.UnicodeSeqOfUtf8Bytes("w"))).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("l")))
	} else {
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ll"))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.Int, p2 _dafny.Set, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), !(!(false)))).Merge(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf(!(true)))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _21_v0 _dafny.MultiSet
			_21_v0 = interface{}(_compr_5).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf(!(true))), _21_v0) {
				_coll5.Add(_21_v0, false)
			}
		}
		return _coll5.ToMap()
	}())).Merge((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)), _dafny.MultiSetOf(true), _dafny.MultiSetOf(true, false), _dafny.MultiSetOf(false))).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _22_v1 _dafny.MultiSet
			_22_v1 = interface{}(_compr_6).(_dafny.MultiSet)
			if (_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)), _dafny.MultiSetOf(true), _dafny.MultiSetOf(true, false), _dafny.MultiSetOf(false))).Contains(_22_v1) {
				_coll6.Add(_22_v1, false)
			}
		}
		return _coll6.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((Companion_D8_.Create_DC18_((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(-998))).Cardinality()), true)).Dtor_cf30()), !(true))))
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Map {
	var _source1 D6 = Companion_D6_.Create_DC13_()
	_ = _source1
	if _source1.Is_DC11() {
		var _23___mcc_h0 _dafny.Sequence = _source1.Get_().(D6_DC11).Cf17
		_ = _23___mcc_h0
		var _24___mcc_h1 *C0 = _source1.Get_().(D6_DC11).Cf18
		_ = _24___mcc_h1
		var _25___mcc_h2 _dafny.Int = _source1.Get_().(D6_DC11).Cf19
		_ = _25___mcc_h2
		var _26___mcc_h3 _dafny.Sequence = _source1.Get_().(D6_DC11).Cf20
		_ = _26___mcc_h3
		var _27___mcc_h4 _dafny.Int = _source1.Get_().(D6_DC11).Cf21
		_ = _27___mcc_h4
		var _28_cf21 _dafny.Int = _27___mcc_h4
		_ = _28_cf21
		var _29_cf20 _dafny.Sequence = _26___mcc_h3
		_ = _29_cf20
		var _30_cf19 _dafny.Int = _25___mcc_h2
		_ = _30_cf19
		var _31_cf18 *C0 = _24___mcc_h1
		_ = _31_cf18
		var _32_cf17 _dafny.Sequence = _23___mcc_h0
		_ = _32_cf17
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true, true), _28_cf21)
	} else if _source1.Is_DC12() {
		var _33___mcc_h5 bool = _source1.Get_().(D6_DC12).Cf22
		_ = _33___mcc_h5
		var _34_cf22 bool = _33___mcc_h5
		_ = _34_cf22
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_34_cf22), _dafny.IntOfInt64(-495))
	} else if _source1.Is_DC13() {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(!(false), true, true, true, false), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hw")).Cardinality())))
	} else if _source1.Is_DC10() {
		var _35___mcc_h6 _dafny.Map = _source1.Get_().(D6_DC10).Cf16
		_ = _35___mcc_h6
		var _36_cf16 _dafny.Map = _35___mcc_h6
		_ = _36_cf16
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), _dafny.IntOfInt64(414))
	} else {
		var _37___mcc_h7 D6 = _source1.Get_().(D6_DC14).Cf23
		_ = _37___mcc_h7
		var _38_cf23 D6 = _37___mcc_h7
		_ = _38_cf23
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), _dafny.IntOfInt64(-560))
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true)
}
func (_static *CompanionStruct_Default___) M1(p0 _dafny.Int, p1 _dafny.Array, globalState *GlobalState) _dafny.Array {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var _39_v0 bool
	_ = _39_v0
	_39_v0 = false
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))
	_ = _index0
	(p1).ArraySet1(_39_v0, (_index0).Int())
	var _40_v1 _dafny.Sequence
	_ = _40_v1
	_40_v1 = _dafny.UnicodeSeqOfUtf8Bytes("xdfamt")
	var _41_v2 _dafny.MultiSet
	_ = _41_v2
	_41_v2 = _dafny.MultiSetOf(p0, (_dafny.Zero).Minus(p0))
	var _42_v3 _dafny.Sequence
	_ = _42_v3
	_42_v3 = _dafny.SeqOf(p0, p0, p0, _dafny.IntOfUint32((_40_v1).Cardinality()), (func() _dafny.Int {
		if (_41_v2).Contains(p0) {
			return (_41_v2).Multiplicity(p0)
		}
		return p0
	})())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))
	_ = _index1
	(p1).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(p0, p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("benlgnvi")).Cardinality()), p0), _42_v3)) || (_39_v0), (_index1).Int())
	var _43_v4 D2
	_ = _43_v4
	_43_v4 = Companion_D2_.Create_DC3_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
	var _44_v5 _dafny.Array
	_ = _44_v5
	var _nwElement0_0 D2 = Companion_D2_.Create_DC3_((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(3))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_43_v4, 1)
	(_nw0).ArraySet1(_43_v4, 2)
	_44_v5 = _nw0
	var _45_v6 _dafny.Sequence
	_ = _45_v6
	_45_v6 = _dafny.SeqOf(_44_v5)
	_45_v6 = _45_v6
	var _46_v7 _dafny.Sequence
	_ = _46_v7
	_46_v7 = _dafny.SeqOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
	if (!_dafny.Companion_Sequence_.Contains(_46_v7, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))) && (!(!(Companion_Default___.Fm0(_39_v0, true, _dafny.IntOfInt64(444), p0, globalState)) || (!((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))))) {
		var _47_v8 bool
		_ = _47_v8
		_47_v8 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)
		var _48_v9 *C1
		_ = _48_v9
		var _nw1 *C1 = New_C1_()
		_ = _nw1
		_nw1.Ctor__(_47_v8)
		_48_v9 = _nw1
		var _49_v10 _dafny.CodePoint
		_ = _49_v10
		_49_v10 = _dafny.CodePoint('i')
		var _50_v11 D5
		_ = _50_v11
		_50_v11 = Companion_D5_.Create_DC9_(_40_v1, _39_v0, _48_v9, _49_v10)
		var _pat_let_tv0 = _48_v9
		_ = _pat_let_tv0
		var _source2 D5 = func(_pat_let0_0 D5) D5 {
			return func(_51_dt__update__tmp_h0 D5) D5 {
				return func(_pat_let1_0 *C1) D5 {
					return func(_52_dt__update_hcf14_h0 *C1) D5 {
						return Companion_D5_.Create_DC9_((_51_dt__update__tmp_h0).Dtor_cf12(), (_51_dt__update__tmp_h0).Dtor_cf13(), _52_dt__update_hcf14_h0, (_51_dt__update__tmp_h0).Dtor_cf15())
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}((func() D5 {
			if !(_39_v0) {
				return Companion_D5_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("j"), _39_v0, _48_v9, _49_v10)
			}
			return _50_v11
		})())
		_ = _source2
		if _source2.Is_DC9() {
			var _53___mcc_h0 _dafny.Sequence = _source2.Get_().(D5_DC9).Cf12
			_ = _53___mcc_h0
			var _54___mcc_h1 bool = _source2.Get_().(D5_DC9).Cf13
			_ = _54___mcc_h1
			var _55___mcc_h2 *C1 = _source2.Get_().(D5_DC9).Cf14
			_ = _55___mcc_h2
			var _56___mcc_h3 _dafny.CodePoint = _source2.Get_().(D5_DC9).Cf15
			_ = _56___mcc_h3
			var _57_cf15 _dafny.CodePoint = _56___mcc_h3
			_ = _57_cf15
			var _58_cf14 *C1 = _55___mcc_h2
			_ = _58_cf14
			var _59_cf13 bool = _54___mcc_h1
			_ = _59_cf13
			var _60_cf12 _dafny.Sequence = _53___mcc_h0
			_ = _60_cf12
			var _61_v12 _dafny.Set
			_ = _61_v12
			_61_v12 = _dafny.SetOf(_59_cf13)
			var _62_v13 _dafny.Map
			_ = _62_v13
			_62_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_61_v12, (_dafny.MultiSetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))).Cardinality())
			var _63_v14 *C0
			_ = _63_v14
			var _nw2 *C0 = New_C0_()
			_ = _nw2
			_nw2.Ctor__((_60_cf12).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_60_cf12).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_63_v14 = _nw2
			var _64_v15 _dafny.Map
			_ = _64_v15
			_64_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm16(globalState)).Equals(_62_v13), _63_v14)
			var _rhs0 bool = !(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(294), _dafny.IntOfInt64(485))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _65_v16 _dafny.Int
					_65_v16 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(294)).Cmp(_65_v16) <= 0) && ((_65_v16).Cmp(_dafny.IntOfInt64(485)) < 0) {
						_coll7.Add((_65_v16).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_40_v1, _60_cf12)).Cardinality())))
					}
				}
				return _coll7.ToSet()
			}()).Contains(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
					return _46_v7
				}
				return _46_v7
			})()).Cardinality()))
			_ = _rhs0
			var _rhs1 bool = (_39_v0) == (Companion_Default___.Fm0(_39_v0, _39_v0, p0, p0, globalState))
			_ = _rhs1
			var _rhs2 _dafny.Map = _64_v15
			_ = _rhs2
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			_39_v0 = _rhs0
			_lhs0.F9 = _rhs1
			_64_v15 = _rhs2
			var _66_v17 *C0
			_ = _66_v17
			var _nw3 *C0 = New_C0_()
			_ = _nw3
			_nw3.Ctor__(_49_v10)
			_66_v17 = _nw3
			var _67_v18 _dafny.Map
			_ = _67_v18
			_67_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_41_v2).Cardinality(), p1)
			var _68_v19 D5
			_ = _68_v19
			_68_v19 = Companion_D5_.Create_DC8_((_67_v18).Merge(_67_v18))
			_68_v19 = (func() D5 {
				if _39_v0 {
					return _68_v19
				}
				return _68_v19
			})()
			var _69_v20 _dafny.Map
			_ = _69_v20
			_69_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.Fm2(_39_v0, p0, p0, _60_cf12, globalState)), _49_v10)
			var _70_v21 _dafny.Map
			_ = _70_v21
			_70_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v0, p0)
			var _71_v22 _dafny.Map
			_ = _71_v22
			_71_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _70_v21)
			var _72_v23 _dafny.Array
			_ = _72_v23
			var _nwElement0_1 _dafny.CodePoint = _dafny.CodePoint('k')
			_ = _nwElement0_1
			var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(28))
			_ = _nw4
			(_nw4).ArraySet1CodePoint(_nwElement0_1, 0)
			(_nw4).ArraySet1CodePoint(_57_cf15, 1)
			(_nw4).ArraySet1CodePoint(_57_cf15, 2)
			(_nw4).ArraySet1CodePoint(_57_cf15, 3)
			(_nw4).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_69_v20).Contains(_dafny.IntOfInt64(285)) {
					return (_69_v20).Get(_dafny.IntOfInt64(285)).(_dafny.CodePoint)
				}
				return _49_v10
			})(), 4)
			(_nw4).ArraySet1CodePoint(_57_cf15, 5)
			(_nw4).ArraySet1CodePoint(_dafny.CodePoint('e'), 6)
			(_nw4).ArraySet1CodePoint((_63_v14).F12(), 7)
			(_nw4).ArraySet1CodePoint((_66_v17).F12(), 8)
			(_nw4).ArraySet1CodePoint(_49_v10, 9)
			(_nw4).ArraySet1CodePoint((_66_v17).F12(), 10)
			(_nw4).ArraySet1CodePoint(_49_v10, 11)
			(_nw4).ArraySet1CodePoint((_66_v17).F12(), 12)
			(_nw4).ArraySet1CodePoint(_57_cf15, 13)
			(_nw4).ArraySet1CodePoint((_40_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-569), _dafny.IntOfUint32((_40_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), 14)
			(_nw4).ArraySet1CodePoint(_49_v10, 15)
			(_nw4).ArraySet1CodePoint((_63_v14).F12(), 16)
			(_nw4).ArraySet1CodePoint((_63_v14).F12(), 17)
			(_nw4).ArraySet1CodePoint(Companion_Default___.Fm5((_71_v22).Cardinality(), globalState), 18)
			(_nw4).ArraySet1CodePoint(_49_v10, 19)
			(_nw4).ArraySet1CodePoint(_49_v10, 20)
			(_nw4).ArraySet1CodePoint((_66_v17).F12(), 21)
			(_nw4).ArraySet1CodePoint(_49_v10, 22)
			(_nw4).ArraySet1CodePoint(_57_cf15, 23)
			(_nw4).ArraySet1CodePoint(_49_v10, 24)
			(_nw4).ArraySet1CodePoint((_63_v14).F12(), 25)
			(_nw4).ArraySet1CodePoint(_dafny.CodePoint('f'), 26)
			(_nw4).ArraySet1CodePoint((_63_v14).F12(), 27)
			_72_v23 = _nw4
			var _73_v24 _dafny.Map
			_ = _73_v24
			_73_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_cf13, _72_v23)
			var _74_v25 _dafny.Sequence
			_ = _74_v25
			_74_v25 = _dafny.SeqOf((func() _dafny.Array {
				if (_73_v24).Contains((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)) {
					return (_73_v24).Get((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _72_v23
			})())
			_39_v0 = !(true) || (!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_74_v25, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_74_v25).Cardinality()))).Uint32(), _72_v23), _72_v23))
		} else {
			var _75___mcc_h4 _dafny.Map = _source2.Get_().(D5_DC8).Cf11
			_ = _75___mcc_h4
			var _76_cf11 _dafny.Map = _75___mcc_h4
			_ = _76_cf11
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))
			_ = _index2
			(p1).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool), (_index2).Int())
			(globalState).F4 = (p0).Times(p0)
			var _77_v26 _dafny.Array
			_ = _77_v26
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
			_ = _nw5
			_77_v26 = _nw5
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_77_v26), 0))
			_ = _index3
			(_77_v26).ArraySet1CodePoint(_49_v10, (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_77_v26), 0))
			_ = _index4
			var _rhs3 _dafny.Int = _dafny.IntOfInt64(-951)
			_ = _rhs3
			var _rhs4 _dafny.Sequence = _40_v1
			_ = _rhs4
			var _rhs5 _dafny.CodePoint = _49_v10
			_ = _rhs5
			var _lhs1 *GlobalState = globalState
			_ = _lhs1
			var _lhs2 _dafny.Array = _77_v26
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_77_v26), 0))
			_ = _lhs3
			_lhs1.F4 = _rhs3
			_40_v1 = _rhs4
			(_lhs2).ArraySet1CodePoint(_rhs5, (_lhs3).Int())
			var _78_v27 _dafny.Map
			_ = _78_v27
			_78_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_77_v26).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_77_v26), 0))).Int()))
			var _79_v28 _dafny.Map
			_ = _79_v28
			_79_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _78_v27)
			(globalState).F4 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(p0, p0), (func() _dafny.Int {
				if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
					return p0
				}
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(((func() _dafny.Map {
					if (_79_v28).Contains(p0) {
						return (_79_v28).Get(p0).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _49_v10)
				})()).Cardinality()))
			})())
		}
		var _80_v29 _dafny.MultiSet
		_ = _80_v29
		_80_v29 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hqp"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(12))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}((func(_81_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_82_i0 _dafny.Int) _dafny.CodePoint {
				return _81_v10
			}
		})(_49_v10))))
		var _83_v30 _dafny.MultiSet
		_ = _83_v30
		_83_v30 = _dafny.MultiSetOf(_39_v0)
		var _84_v31 _dafny.Map
		_ = _84_v31
		_84_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _83_v30)
		var _85_v32 _dafny.Map
		_ = _85_v32
		_85_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool), _39_v0)
		var _86_v33 _dafny.Map
		_ = _86_v33
		_86_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
		var _87_v34 _dafny.Map
		_ = _87_v34
		_87_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_86_v33).Cardinality(), p0)
		var _88_v35 _dafny.Array
		_ = _88_v35
		var _nwElement0_2 _dafny.Int = p0
		_ = _nwElement0_2
		var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(21))
		_ = _nw6
		(_nw6).ArraySet1(_nwElement0_2, 0)
		(_nw6).ArraySet1((p0).Times(p0), 1)
		(_nw6).ArraySet1(p0, 2)
		(_nw6).ArraySet1((func() _dafny.Int {
			if (_80_v29).Contains(_40_v1) {
				return (_80_v29).Multiplicity(_40_v1)
			}
			return _dafny.IntOfUint32((_40_v1).Cardinality())
		})(), 3)
		(_nw6).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(255))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_89_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_90_i1 _dafny.Int) _dafny.Sequence {
				return _89_v1
			}
		})(_40_v1)))).Cardinality()), 4)
		(_nw6).ArraySet1(p0, 5)
		(_nw6).ArraySet1(p0, 6)
		(_nw6).ArraySet1((_dafny.Zero).Minus(p0), 7)
		(_nw6).ArraySet1((p0).Minus(p0), 8)
		(_nw6).ArraySet1(p0, 9)
		(_nw6).ArraySet1((_84_v31).Cardinality(), 10)
		(_nw6).ArraySet1(_dafny.IntOfInt64(936), 11)
		(_nw6).ArraySet1(p0, 12)
		(_nw6).ArraySet1((_dafny.Zero).Minus(((_85_v32).Update(_39_v0, true)).Cardinality()), 13)
		(_nw6).ArraySet1(p0, 14)
		(_nw6).ArraySet1(p0, 15)
		(_nw6).ArraySet1(p0, 16)
		(_nw6).ArraySet1(p0, 17)
		(_nw6).ArraySet1((p0).Times((func() _dafny.Int {
			if (_83_v30).Contains(_39_v0) {
				return (_83_v30).Multiplicity(_39_v0)
			}
			return (_87_v34).Cardinality()
		})()), 18)
		(_nw6).ArraySet1(_dafny.IntOfInt64(405), 19)
		(_nw6).ArraySet1((func() _dafny.Int {
			if _39_v0 {
				return p0
			}
			return p0
		})(), 20)
		_88_v35 = _nw6
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))
		_ = _index5
		(_88_v35).ArraySet1(p0, (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))
		_ = _index6
		(_88_v35).ArraySet1((func() _dafny.Int {
			if (_87_v34).Contains((Companion_D4_.Create_DC7_((func() _dafny.Int {
				if (_83_v30).Contains(_39_v0) {
					return (_83_v30).Multiplicity(_39_v0)
				}
				return p0
			})(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))).Dtor_cf9()) {
				return (_87_v34).Get((Companion_D4_.Create_DC7_((func() _dafny.Int {
					if (_83_v30).Contains(_39_v0) {
						return (_83_v30).Multiplicity(_39_v0)
					}
					return p0
				})(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))).Dtor_cf9()).(_dafny.Int)
			}
			return p0
		})(), (_index6).Int())
		var _91_v36 *C0
		_ = _91_v36
		var _nw7 *C0 = New_C0_()
		_ = _nw7
		_nw7.Ctor__((func() _dafny.CodePoint {
			if Companion_Default___.Fm0((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool), true, (_dafny.Zero).Minus(p0), (_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int), globalState) {
				return Companion_Default___.Fm5(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_40_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.IntOfUint32((_40_v1).Cardinality()))).Uint32(), _49_v10)).Cardinality()), globalState)
			}
			return _49_v10
		})())
		_91_v36 = _nw7
		(globalState).F10 = !((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
		var _92_v37 _dafny.Set
		_ = _92_v37
		_92_v37 = _dafny.SetOf(_40_v1, _40_v1)
		var _nwElement0_3 _dafny.Int = (_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int)
		_ = _nwElement0_3
		var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(12))
		_ = _nw8
		(_nw8).ArraySet1(_nwElement0_3, 0)
		(_nw8).ArraySet1((_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int), 1)
		(_nw8).ArraySet1(p0, 2)
		(_nw8).ArraySet1(p0, 3)
		(_nw8).ArraySet1(p0, 4)
		(_nw8).ArraySet1((_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int), 5)
		(_nw8).ArraySet1((_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int), 6)
		(_nw8).ArraySet1((_92_v37).Cardinality(), 7)
		(_nw8).ArraySet1((_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int), 8)
		(_nw8).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 9)
		(_nw8).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_46_v7, _46_v7)).Cardinality()), 10)
		(_nw8).ArraySet1((_88_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_88_v35), 0))).Int()).(_dafny.Int), 11)
		_88_v35 = _nw8
	} else {
		var _93_v38 bool
		_ = _93_v38
		_93_v38 = !(_39_v0)
		var _94_v39 *C1
		_ = _94_v39
		var _nw9 *C1 = New_C1_()
		_ = _nw9
		_nw9.Ctor__(_93_v38)
		_94_v39 = _nw9
		var _95_v40 _dafny.Map
		_ = _95_v40
		_95_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v39, _39_v0)
		var _96_v41 D6
		_ = _96_v41
		_96_v41 = Companion_D6_.Create_DC10_(_95_v40)
		var _97_v42 _dafny.Map
		_ = _97_v42
		_97_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v2, (_96_v41).Dtor_cf16())
		var _98_v43 _dafny.Sequence
		_ = _98_v43
		_98_v43 = _dafny.SeqOf(_95_v40)
		var _99_v44 _dafny.Map
		_ = _99_v44
		_99_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-989), _95_v40)
		var _100_v45 _dafny.Map
		_ = _100_v45
		_100_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool), _94_v39)
		var _101_v46 _dafny.Sequence
		_ = _101_v46
		_101_v46 = _dafny.SeqOf(_94_v39)
		var _102_v47 _dafny.Array
		_ = _102_v47
		var _nwElement0_4 _dafny.Map = (func() _dafny.Map {
			if (_97_v42).Contains(_41_v2) {
				return (_97_v42).Get(_41_v2).(_dafny.Map)
			}
			return _95_v40
		})()
		_ = _nwElement0_4
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(25))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_4, 0)
		(_nw10).ArraySet1(_95_v40, 1)
		(_nw10).ArraySet1(_95_v40, 2)
		(_nw10).ArraySet1(_95_v40, 3)
		(_nw10).ArraySet1(_95_v40, 4)
		(_nw10).ArraySet1((_95_v40).Merge(_95_v40), 5)
		(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v39, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)), 6)
		(_nw10).ArraySet1(_95_v40, 7)
		(_nw10).ArraySet1((_98_v43).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_98_v43).Cardinality()))).Uint32()).(_dafny.Map), 8)
		(_nw10).ArraySet1(_95_v40, 9)
		(_nw10).ArraySet1((_95_v40).Merge(_95_v40), 10)
		(_nw10).ArraySet1(_95_v40, 11)
		(_nw10).ArraySet1(_95_v40, 12)
		(_nw10).ArraySet1(((_95_v40).Update(_94_v39, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v39, _39_v0)), 13)
		(_nw10).ArraySet1(_95_v40, 14)
		(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v39, Companion_Default___.Fm0(_39_v0, Companion_Default___.Fm0(true, _39_v0, p0, p0, globalState), p0, p0, globalState)), 15)
		(_nw10).ArraySet1(_95_v40, 16)
		(_nw10).ArraySet1(((func() _dafny.Map {
			if (_99_v44).Contains(_dafny.IntOfUint32((_42_v3).Cardinality())) {
				return (_99_v44).Get(_dafny.IntOfUint32((_42_v3).Cardinality())).(_dafny.Map)
			}
			return _95_v40
		})()).Merge(_95_v40), 17)
		(_nw10).ArraySet1(_95_v40, 18)
		(_nw10).ArraySet1(_95_v40, 19)
		(_nw10).ArraySet1(_95_v40, 20)
		(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C1 {
			if (_100_v45).Contains(_39_v0) {
				return (_100_v45).Get(_39_v0).(*C1)
			}
			return _94_v39
		})(), _39_v0), 21)
		(_nw10).ArraySet1(_95_v40, 22)
		(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_101_v46).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_101_v46).Cardinality()))).Uint32()).(*C1), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)), 23)
		(_nw10).ArraySet1((_95_v40).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v39, _39_v0)), 24)
		_102_v47 = _nw10
		var _103_v48 _dafny.Map
		_ = _103_v48
		_103_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(p0))
		var _104_v49 _dafny.Array
		_ = _104_v49
		var _nwElement0_5 _dafny.Int = p0
		_ = _nwElement0_5
		var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(3))
		_ = _nw11
		(_nw11).ArraySet1(_nwElement0_5, 0)
		(_nw11).ArraySet1((func() _dafny.Int {
			if (_103_v48).Contains(p0) {
				return (_103_v48).Get(p0).(_dafny.Int)
			}
			return p0
		})(), 1)
		(_nw11).ArraySet1(p0, 2)
		_104_v49 = _nw11
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_104_v49), 0))
		_ = _index7
		(_104_v49).ArraySet1(Companion_Default___.Fm2(_39_v0, p0, p0, _40_v1, globalState), (_index7).Int())
		var _105_v50 _dafny.Array
		_ = _105_v50
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(7))
		_ = _nw12
		_105_v50 = _nw12
		var _106_v51 _dafny.Set
		_ = _106_v51
		_106_v51 = _dafny.SetOf(_39_v0)
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_105_v50), 0))
		_ = _index8
		(_105_v50).ArraySet1(_106_v51, (_index8).Int())
		var _107_v52 _dafny.Map
		_ = _107_v52
		_107_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v2, _102_v47)
		var _108_v53 _dafny.CodePoint
		_ = _108_v53
		_108_v53 = _dafny.CodePoint('h')
		var _109_v54 *C0
		_ = _109_v54
		var _nw13 *C0 = New_C0_()
		_ = _nw13
		_nw13.Ctor__(_108_v53)
		_109_v54 = _nw13
		var _110_v55 _dafny.Map
		_ = _110_v55
		_110_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)))
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_104_v49), 0))
		_ = _index9
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_105_v50), 0))
		_ = _index10
		var _rhs6 _dafny.Array = (func() _dafny.Array {
			if (_107_v52).Contains(_dafny.MultiSetOf(_dafny.IntOfUint32(((Companion_D6_.Create_DC11_(_40_v1, _109_v54, _dafny.IntOfUint32((_dafny.SeqOf(_109_v54)).Cardinality()), _dafny.SeqOf(p0, p0), p0)).Dtor_cf20()).Cardinality()), p0, p0, p0, _dafny.IntOfInt64(-582))) {
				return (_107_v52).Get(_dafny.MultiSetOf(_dafny.IntOfUint32(((Companion_D6_.Create_DC11_(_40_v1, _109_v54, _dafny.IntOfUint32((_dafny.SeqOf(_109_v54)).Cardinality()), _dafny.SeqOf(p0, p0), p0)).Dtor_cf20()).Cardinality()), p0, p0, p0, _dafny.IntOfInt64(-582))).(_dafny.Array)
			}
			return _102_v47
		})()
		_ = _rhs6
		var _rhs7 _dafny.Int = p0
		_ = _rhs7
		var _rhs8 _dafny.Int = (_dafny.Zero).Minus(p0)
		_ = _rhs8
		var _rhs9 bool = _39_v0
		_ = _rhs9
		var _rhs10 _dafny.Set = Companion_Default___.Fm17(_110_v55, _40_v1, p0, globalState)
		_ = _rhs10
		var _lhs4 _dafny.Array = _104_v49
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_104_v49), 0))
		_ = _lhs5
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		var _lhs8 _dafny.Array = _105_v50
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_105_v50), 0))
		_ = _lhs9
		_102_v47 = _rhs6
		(_lhs4).ArraySet1(_rhs7, (_lhs5).Int())
		_lhs6.F4 = _rhs8
		_lhs7.F9 = _rhs9
		(_lhs8).ArraySet1(_rhs10, (_lhs9).Int())
		(globalState).F9 = _39_v0
		(globalState).F4 = (p0).Plus(p0)
		var _111_v56 _dafny.Array
		_ = _111_v56
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_0
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = (func(_112_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_113_i2 _dafny.Int) _dafny.Sequence {
					return _112_v7
				}
			})(_46_v7)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw14 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw14).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw14).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_111_v56 = _nw14
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_111_v56), 0))
		_ = _index11
		(_111_v56).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_46_v7, _46_v7), (_index11).Int())
		var _114_v57 D7
		_ = _114_v57
		_114_v57 = Companion_D7_.Create_DC15_(_46_v7)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_111_v56), 0))
		_ = _index12
		(_111_v56).ArraySet1((_114_v57).Dtor_cf24(), (_index12).Int())
		var _115_v58 _dafny.Array
		_ = _115_v58
		var _nwElement0_6 _dafny.Array = p1
		_ = _nwElement0_6
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(17))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_6, 0)
		(_nw15).ArraySet1((func() _dafny.Array {
			if _39_v0 {
				return p1
			}
			return p1
		})(), 1)
		(_nw15).ArraySet1(p1, 2)
		(_nw15).ArraySet1(p1, 3)
		(_nw15).ArraySet1(p1, 4)
		(_nw15).ArraySet1(p1, 5)
		(_nw15).ArraySet1(p1, 6)
		(_nw15).ArraySet1(p1, 7)
		(_nw15).ArraySet1(p1, 8)
		(_nw15).ArraySet1(p1, 9)
		(_nw15).ArraySet1(p1, 10)
		(_nw15).ArraySet1(p1, 11)
		(_nw15).ArraySet1(p1, 12)
		(_nw15).ArraySet1(p1, 13)
		(_nw15).ArraySet1(p1, 14)
		(_nw15).ArraySet1(p1, 15)
		(_nw15).ArraySet1(p1, 16)
		_115_v58 = _nw15
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_115_v58), 0))
		_ = _index13
		(_115_v58).ArraySet1(p1, (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_115_v58), 0))
		_ = _index14
		(_115_v58).ArraySet1(p1, (_index14).Int())
	}
	var _116_v59 _dafny.Array
	_ = _116_v59
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_1
	var _nw16 _dafny.Array
	_ = _nw16
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw16 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Sequence = (func(_117_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_118_i3 _dafny.Int) _dafny.Sequence {
				return _117_v3
			}
		})(_42_v3)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw16 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw16).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw16).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_116_v59 = _nw16
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_116_v59), 0))
	_ = _index15
	(_116_v59).ArraySet1(_42_v3, (_index15).Int())
	var _119_v60 _dafny.Array
	_ = _119_v60
	var _nw17 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
	_ = _nw17
	_119_v60 = _nw17
	var _120_v61 bool
	_ = _120_v61
	_120_v61 = false
	var _121_v62 *C1
	_ = _121_v62
	var _nw18 *C1 = New_C1_()
	_ = _nw18
	_nw18.Ctor__(_120_v61)
	_121_v62 = _nw18
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_119_v60), 0))
	_ = _index16
	(_119_v60).ArraySet1(_121_v62, (_index16).Int())
	var _122_v63 _dafny.Set
	_ = _122_v63
	_122_v63 = _dafny.SetOf(p0, (_41_v2).Cardinality())
	var _123_v64 _dafny.MultiSet
	_ = _123_v64
	_123_v64 = _dafny.MultiSetOf(_122_v63)
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_116_v59), 0))
	_ = _index17
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_119_v60), 0))
	_ = _index18
	var _rhs11 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_42_v3, _dafny.Companion_Sequence_.Concatenate(_42_v3, _42_v3))
	_ = _rhs11
	var _rhs12 *C1 = _121_v62
	_ = _rhs12
	var _rhs13 *C1 = _121_v62
	_ = _rhs13
	var _rhs14 _dafny.MultiSet = _123_v64
	_ = _rhs14
	var _lhs10 _dafny.Array = _116_v59
	_ = _lhs10
	var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_116_v59), 0))
	_ = _lhs11
	var _lhs12 _dafny.Array = _119_v60
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_119_v60), 0))
	_ = _lhs13
	(_lhs10).ArraySet1(_rhs11, (_lhs11).Int())
	(_lhs12).ArraySet1(_rhs12, (_lhs13).Int())
	_121_v62 = _rhs13
	_123_v64 = _rhs14
	var _124_v65 _dafny.Set
	_ = _124_v65
	_124_v65 = _dafny.SetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
	if (_124_v65).IsSubsetOf(_124_v65) {
		var _125_v66 _dafny.Map
		_ = _125_v66
		_125_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v61, !((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)))
		_125_v66 = (_125_v66).Update((_121_v62).F11(), _39_v0)
		_40_v1 = (func() _dafny.Sequence {
			if _39_v0 {
				return (func() _dafny.Sequence {
					if true {
						return _40_v1
					}
					return _40_v1
				})()
			}
			return _dafny.Companion_Sequence_.Concatenate(_40_v1, _40_v1)
		})()
		var _126_v67 _dafny.Map
		_ = _126_v67
		_126_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_46_v7, _46_v7), _39_v0)
		_126_v67 = (_126_v67).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))), _46_v7), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool))
		(globalState).F3 = p1
		(globalState).F4 = _dafny.IntOfInt64(31)
	} else {
		var _127_v68 *C1
		_ = _127_v68
		var _nw19 *C1 = New_C1_()
		_ = _nw19
		_nw19.Ctor__((_121_v62).F11())
		_127_v68 = _nw19
		var _128_v69 _dafny.Map
		_ = _128_v69
		_128_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _39_v0)
		var _129_v70 _dafny.Set
		_ = _129_v70
		_129_v70 = _dafny.SetOf(_128_v69, _128_v69)
		_129_v70 = _129_v70
		var _130_v71 D6
		_ = _130_v71
		_130_v71 = Companion_D6_.Create_DC13_()
		var _131_v72 _dafny.Sequence
		_ = _131_v72
		_131_v72 = _dafny.SeqOf(_130_v71, _130_v71, _130_v71)
		var _132_v73 D6
		_ = _132_v73
		_132_v73 = Companion_D6_.Create_DC14_((_131_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_40_v1).Cardinality()), _dafny.IntOfUint32((_131_v72).Cardinality()))).Uint32()).(D6))
		_132_v73 = _132_v73
		_40_v1 = _dafny.Companion_Sequence_.Concatenate(_40_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(771))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_133_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})))
		(globalState).F9 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((p1), 0))).Int()).(bool)
	}
	var _134_v74 _dafny.Sequence
	_ = _134_v74
	var _out0 _dafny.Sequence
	_ = _out0
	_out0 = (_121_v62).M0(_dafny.IntOfUint32((_40_v1).Cardinality()), _40_v1, _39_v0, p0, globalState)
	_134_v74 = _out0
	r0 = p1
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _135_v0 _dafny.Array
	_ = _135_v0
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
	_ = _nw20
	_135_v0 = _nw20
	var _136_globalState *GlobalState
	_ = _136_globalState
	var _nw21 *GlobalState = New_GlobalState_()
	_ = _nw21
	_nw21.Ctor__(_135_v0, _dafny.IntOfInt64(319), false, _135_v0, _dafny.IntOfInt64(250), _dafny.IntOfInt64(888), false, _dafny.IntOfInt64(746), false, false, true)
	_136_globalState = _nw21
	var _137_v1 bool
	_ = _137_v1
	_137_v1 = true
	var _138_v2 _dafny.MultiSet
	_ = _138_v2
	_138_v2 = _dafny.MultiSetOf(_137_v1)
	var _139_v3 _dafny.Map
	_ = _139_v3
	_139_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_138_v2).Union(_138_v2), _137_v1)
	_139_v3 = (_139_v3).Merge(_139_v3)
	var _140_v4 _dafny.Int
	_ = _140_v4
	_140_v4 = _dafny.IntOfInt64(384)
	var _hi0 _dafny.Int = ((_dafny.Zero).Minus(_140_v4)).Minus(_140_v4)
	_ = _hi0
	for _141_i0 := _140_v4; _141_i0.Cmp(_hi0) < 0; _141_i0 = _141_i0.Plus(_dafny.One) {
		var _142_v5 _dafny.Sequence
		_ = _142_v5
		_142_v5 = _dafny.SeqOf(_137_v1)
		_137_v1 = (func() bool {
			if Companion_Default___.Fm0(_137_v1, _137_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bfjvu")).Cardinality()), _dafny.IntOfUint32((_142_v5).Cardinality()), _136_globalState) {
				return _137_v1
			}
			return Companion_Default___.Fm0(_137_v1, !(true), _140_v4, _141_i0, _136_globalState)
		})()
		var _143_v6 _dafny.Sequence
		_ = _143_v6
		_143_v6 = _dafny.SeqOf(_135_v0)
		_143_v6 = _dafny.SeqOf(_135_v0, _135_v0)
		var _144_v7 _dafny.Array
		_ = _144_v7
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_2
		var _nw22 _dafny.Array
		_ = _nw22
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw22 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Map = (func(_145_i0 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_146_i1 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_i0, _dafny.CodePoint('g'))
				}
			})(_141_i0)
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
		_144_v7 = _nw22
		var _147_v8 _dafny.CodePoint
		_ = _147_v8
		_147_v8 = _dafny.CodePoint('o')
		var _148_v9 _dafny.Map
		_ = _148_v9
		_148_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v4, _147_v8)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_144_v7), 0))
		_ = _index19
		(_144_v7).ArraySet1(_148_v9, (_index19).Int())
		var _149_v10 bool
		_ = _149_v10
		_149_v10 = false
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_144_v7), 0))
		_ = _index20
		(_144_v7).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_141_i0).Minus((Companion_Default___.Fm1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_i0, _140_v4), (_149_v10), _141_i0, _140_v4, _136_globalState)).Cardinality()), _147_v8), (_index20).Int())
		var _150_v11 _dafny.Sequence
		_ = _150_v11
		_150_v11 = _dafny.UnicodeSeqOfUtf8Bytes("ewk")
		_150_v11 = _150_v11
	}
	var _151_v12 _dafny.Map
	_ = _151_v12
	_151_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_140_v4).Minus(_dafny.IntOfInt64(261)), !(Companion_Default___.Fm0(_137_v1, _137_v1, _140_v4, _dafny.IntOfInt64(789), _136_globalState)))
	_151_v12 = _151_v12
	var _152_v13 _dafny.Map
	_ = _152_v13
	_152_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v1, _dafny.UnicodeSeqOfUtf8Bytes("du"))
	var _153_v14 _dafny.Sequence
	_ = _153_v14
	_153_v14 = _dafny.SeqOf(_140_v4, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_152_v13).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_154_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality())))
	var _155_v15 bool
	_ = _155_v15
	_155_v15 = _137_v1
	var _156_v16 _dafny.Map
	_ = _156_v16
	_156_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v1, _140_v4)
	var _157_v17 _dafny.Sequence
	_ = _157_v17
	_157_v17 = _dafny.UnicodeSeqOfUtf8Bytes("ajftdewh")
	(_136_globalState).F4 = (_153_v14).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm2(_137_v1, (func() _dafny.Int {
		if (_156_v16).Contains(_155_v15) {
			return (_156_v16).Get(_155_v15).(_dafny.Int)
		}
		return _140_v4
	})(), _140_v4, _157_v17, _136_globalState), _dafny.IntOfUint32((_153_v14).Cardinality()))).Uint32()).(_dafny.Int)
	var _158_v18 _dafny.Map
	_ = _158_v18
	_158_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v1, (func() bool {
		if _137_v1 {
			return false
		}
		return _137_v1
	})())
	_158_v18 = (_158_v18).Update(_137_v1, _137_v1)
	var _source3 bool = _155_v15
	_ = _source3
	var _159___mcc_h0 bool = _source3
	_ = _159___mcc_h0
	var _160_cf0 bool = _159___mcc_h0
	_ = _160_cf0
	var _161_v19 _dafny.CodePoint
	_ = _161_v19
	_161_v19 = _dafny.CodePoint('t')
	var _162_v20 *C0
	_ = _162_v20
	var _nw23 *C0 = New_C0_()
	_ = _nw23
	_nw23.Ctor__(_161_v19)
	_162_v20 = _nw23
	var _163_v21 _dafny.Sequence
	_ = _163_v21
	_163_v21 = _dafny.SeqOf(_157_v17)
	var _164_v22 _dafny.Array
	_ = _164_v22
	var _nwElement0_7 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_157_v17, _157_v17)
	_ = _nwElement0_7
	var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(18))
	_ = _nw24
	(_nw24).ArraySet1(_nwElement0_7, 0)
	(_nw24).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kwmdrfli"), 1)
	(_nw24).ArraySet1(_157_v17, 2)
	(_nw24).ArraySet1((_163_v21).Select((Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_163_v21).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
	(_nw24).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(101))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}((func(_165_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_166_i3 _dafny.Int) _dafny.CodePoint {
			return _165_v19
		}
	})(_161_v19))), 4)
	(_nw24).ArraySet1(_157_v17, 5)
	(_nw24).ArraySet1((func() _dafny.Sequence {
		if (_152_v13).Contains(Companion_Default___.Fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState)) {
			return (_152_v13).Get(Companion_Default___.Fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState)).(_dafny.Sequence)
		}
		return _157_v17
	})(), 6)
	(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(195))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}((func(_167_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_168_i4 _dafny.Int) _dafny.CodePoint {
			return _167_v19
		}
	})(_161_v19))), _157_v17), 7)
	(_nw24).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("krisnrwt"), 8)
	(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(496))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}((func(_169_v20 *C0) func(_dafny.Int) _dafny.CodePoint {
		return func(_170_i5 _dafny.Int) _dafny.CodePoint {
			return (_169_v20).F12()
		}
	})(_162_v20)))), 9)
	(_nw24).ArraySet1(_157_v17, 10)
	(_nw24).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(649))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}((func(_171_v20 *C0) func(_dafny.Int) _dafny.CodePoint {
		return func(_172_i6 _dafny.Int) _dafny.CodePoint {
			return (_171_v20).F12()
		}
	})(_162_v20))), 11)
	(_nw24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_157_v17, _157_v17), 12)
	(_nw24).ArraySet1(_157_v17, 13)
	(_nw24).ArraySet1(_157_v17, 14)
	(_nw24).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dkm"), 15)
	(_nw24).ArraySet1(Companion_Default___.Fm4((_dafny.Zero).Minus(_140_v4), _160_cf0, _160_cf0, (_dafny.Zero).Minus(_140_v4), _136_globalState), 16)
	(_nw24).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_157_v17, _dafny.UnicodeSeqOfUtf8Bytes("qfkthhxah")), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_157_v17, _dafny.UnicodeSeqOfUtf8Bytes("qfkthhxah"))).Cardinality()))).Uint32(), _161_v19), 17)
	_164_v22 = _nw24
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_164_v22), 0))
	_ = _index21
	(_164_v22).ArraySet1(_157_v17, (_index21).Int())
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_164_v22), 0))
	_ = _index22
	(_164_v22).ArraySet1(_157_v17, (_index22).Int())
	var _173_v23 _dafny.Array
	_ = _173_v23
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
	_ = _nw25
	_173_v23 = _nw25
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_173_v23), 0))
	_ = _index23
	(_173_v23).ArraySet1CodePoint((_162_v20).F12(), (_index23).Int())
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_173_v23), 0))
	_ = _index24
	(_173_v23).ArraySet1CodePoint(_161_v19, (_index24).Int())
	(_136_globalState).F10 = (!(!(_137_v1)) || (_137_v1)) && (_137_v1)
	(_136_globalState).F10 = _137_v1
	var _174_v24 _dafny.Array
	_ = _174_v24
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_3
	var _nw26 _dafny.Array
	_ = _nw26
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw26 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Int = (func(_175_v17 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_176_i8 _dafny.Int) _dafny.Int {
				return (_176_i8).Minus(_dafny.IntOfUint32((_175_v17).Cardinality()))
			}
		})(_157_v17)
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
	_174_v24 = _nw26
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_174_v24), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _177_i7 _dafny.Int
		_177_i7 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_177_i7).Sign() != -1) && ((_177_i7).Cmp(_dafny.ArrayLen((_174_v24), 0)) < 0)) {
			(_174_v24).ArraySet1((_177_i7).Times(_140_v4), (_177_i7).Int())
		}
	}
	if _137_v1 {
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))
		_ = _index25
		(_174_v24).ArraySet1((_dafny.Zero).Minus(_140_v4), (_index25).Int())
		var _178_v25 _dafny.Sequence
		_ = _178_v25
		_178_v25 = _dafny.SeqOf(_137_v1, _137_v1)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))
		_ = _index26
		var _rhs15 _dafny.Int = (_140_v4).Minus(_140_v4)
		_ = _rhs15
		var _rhs16 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_178_v25, (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_178_v25).Cardinality()))).Uint32(), _137_v1), _137_v1)
		_ = _rhs16
		var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt(_140_v4, _dafny.IntOfUint32((Companion_Default___.Fm11((_dafny.Zero).Minus(_140_v4), _136_globalState)).Cardinality()))
		_ = _rhs17
		var _rhs18 bool = _137_v1
		_ = _rhs18
		var _lhs14 _dafny.Array = _174_v24
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))
		_ = _lhs15
		var _lhs16 *GlobalState = _136_globalState
		_ = _lhs16
		var _lhs17 *GlobalState = _136_globalState
		_ = _lhs17
		(_lhs14).ArraySet1(_rhs15, (_lhs15).Int())
		_137_v1 = _rhs16
		_lhs16.F4 = _rhs17
		_lhs17.F9 = _rhs18
		var _source4 D2 = Companion_Default___.Fm12(_140_v4, _136_globalState)
		_ = _source4
		if _source4.Is_DC3() {
			var _179___mcc_h1 bool = _source4.Get_().(D2_DC3).Cf3
			_ = _179___mcc_h1
			var _180___mcc_h2 bool = _source4.Get_().(D2_DC3).Cf4
			_ = _180___mcc_h2
			var _181_cf4 bool = _180___mcc_h2
			_ = _181_cf4
			var _182_cf3 bool = _179___mcc_h1
			_ = _182_cf3
			var _183_v26 _dafny.Set
			_ = _183_v26
			_183_v26 = _dafny.SetOf((func() bool {
				if (_151_v12).Contains(_140_v4) {
					return (_151_v12).Get(_140_v4).(bool)
				}
				return _182_cf3
			})())
			(_136_globalState).F9 = !(_183_v26).Equals(_183_v26)
			_153_v14 = _153_v14
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))
			_ = _index27
			(_174_v24).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm11((_174_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))).Int()).(_dafny.Int), _136_globalState)).Cardinality()), (_index27).Int())
			var _184_v27 D2
			_ = _184_v27
			_184_v27 = Companion_D2_.Create_DC3_((_178_v25).Select((Companion_Default___.SafeIndex((_174_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_178_v25).Cardinality()))).Uint32()).(bool), false)
			var _185_v28 *C0
			_ = _185_v28
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__((func() _dafny.CodePoint {
				if (_184_v27).Dtor_cf4() {
					return _dafny.CodePoint('l')
				}
				return _dafny.CodePoint('k')
			})())
			_185_v28 = _nw27
		} else if _source4.Is_DC4() {
			var _186___mcc_h3 _dafny.Int = _source4.Get_().(D2_DC4).Cf5
			_ = _186___mcc_h3
			var _187___mcc_h4 _dafny.CodePoint = _source4.Get_().(D2_DC4).Cf6
			_ = _187___mcc_h4
			var _188_cf6 _dafny.CodePoint = _187___mcc_h4
			_ = _188_cf6
			var _189_cf5 _dafny.Int = _186___mcc_h3
			_ = _189_cf5
			var _190_v29 *C1
			_ = _190_v29
			var _nw28 *C1 = New_C1_()
			_ = _nw28
			_nw28.Ctor__(_155_v15)
			_190_v29 = _nw28
			_190_v29 = _190_v29
			var _191_v30 _dafny.Array
			_ = _191_v30
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw29
			_191_v30 = _nw29
			var _rhs19 _dafny.Array = _191_v30
			_ = _rhs19
			var _rhs20 bool = _137_v1
			_ = _rhs20
			var _rhs21 _dafny.Sequence = _178_v25
			_ = _rhs21
			var _rhs22 bool = _137_v1
			_ = _rhs22
			var _lhs18 *GlobalState = _136_globalState
			_ = _lhs18
			var _lhs19 *GlobalState = _136_globalState
			_ = _lhs19
			_191_v30 = _rhs19
			_lhs18.F9 = _rhs20
			_178_v25 = _rhs21
			_lhs19.F9 = _rhs22
			_137_v1 = _137_v1
			var _192_v31 _dafny.Set
			_ = _192_v31
			_192_v31 = _dafny.SetOf(_138_v2, _138_v2)
			var _193_v32 _dafny.MultiSet
			_ = _193_v32
			_193_v32 = _dafny.MultiSetOf(_138_v2, _138_v2)
			_192_v31 = func() _dafny.Set {
				var _coll8 = _dafny.NewBuilder()
				_ = _coll8
				for _iter9 := _dafny.Iterate((_193_v32).Elements()); ; {
					_compr_8, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _194_v33 _dafny.MultiSet
					_194_v33 = interface{}(_compr_8).(_dafny.MultiSet)
					if (_193_v32).Contains(_194_v33) {
						_coll8.Add(_194_v33)
					}
				}
				return _coll8.ToSet()
			}()
		} else {
			var _195___mcc_h5 _dafny.CodePoint = _source4.Get_().(D2_DC2).Cf2
			_ = _195___mcc_h5
			var _196_cf2 _dafny.CodePoint = _195___mcc_h5
			_ = _196_cf2
			_157_v17 = _dafny.Companion_Sequence_.Concatenate(_157_v17, _157_v17)
			var _197_v34 *C1
			_ = _197_v34
			var _nw30 *C1 = New_C1_()
			_ = _nw30
			_nw30.Ctor__(_155_v15)
			_197_v34 = _nw30
			_197_v34 = _197_v34
			var _198_v35 *C0
			_ = _198_v35
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(_196_cf2)
			_198_v35 = _nw31
			_198_v35 = _198_v35
			var _199_v36 _dafny.Map
			_ = _199_v36
			_199_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v2, _140_v4)
			var _200_v37 _dafny.Sequence
			_ = _200_v37
			var _out1 _dafny.Sequence
			_ = _out1
			_out1 = (_197_v34).M0((_199_v36).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_157_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_201_cf2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_202_i9 _dafny.Int) _dafny.CodePoint {
					return _201_cf2
				}
			})(_196_cf2)))), !(_137_v1) || (_137_v1), (_174_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))).Int()).(_dafny.Int), _136_globalState)
			_200_v37 = _out1
		}
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))
		_ = _index28
		(_174_v24).ArraySet1((_174_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_174_v24), 0))).Int()).(_dafny.Int), (_index28).Int())
		var _203_v38 _dafny.CodePoint
		_ = _203_v38
		_203_v38 = _dafny.CodePoint('f')
		var _204_v39 D2
		_ = _204_v39
		_204_v39 = Companion_D2_.Create_DC3_(_dafny.Companion_Sequence_.IsPrefixOf(_157_v17, _dafny.Companion_Sequence_.Update(_157_v17, (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_157_v17).Cardinality()))).Uint32(), _203_v38)), _dafny.Companion_Sequence_.IsProperPrefixOf(_157_v17, _dafny.Companion_Sequence_.Update(_157_v17, (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_157_v17).Cardinality()))).Uint32(), _203_v38)))
		_204_v39 = _204_v39
		var _205_v40 _dafny.Array
		_ = _205_v40
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
		_ = _nw32
		_205_v40 = _nw32
		_205_v40 = _205_v40
	} else {
		var _206_v41 _dafny.Set
		_ = _206_v41
		_206_v41 = _dafny.SetOf(_157_v17, _157_v17, _157_v17)
		_206_v41 = Companion_Default___.Fm13(_dafny.IntOfInt64(839), _140_v4, _157_v17, _137_v1, _136_globalState)
		var _207_v42 _dafny.CodePoint
		_ = _207_v42
		_207_v42 = _dafny.CodePoint('w')
		var _208_v43 *C0
		_ = _208_v43
		var _nw33 *C0 = New_C0_()
		_ = _nw33
		_nw33.Ctor__(_207_v42)
		_208_v43 = _nw33
		var _209_v44 _dafny.Sequence
		_ = _209_v44
		_209_v44 = _dafny.SeqOf(_208_v43)
		var _210_v45 _dafny.Array
		_ = _210_v45
		var _nwElement0_8 *C0 = _208_v43
		_ = _nwElement0_8
		var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(21))
		_ = _nw34
		(_nw34).ArraySet1(_nwElement0_8, 0)
		(_nw34).ArraySet1(_208_v43, 1)
		(_nw34).ArraySet1(_208_v43, 2)
		(_nw34).ArraySet1(_208_v43, 3)
		(_nw34).ArraySet1(_208_v43, 4)
		(_nw34).ArraySet1(_208_v43, 5)
		(_nw34).ArraySet1((_209_v44).Select((Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_209_v44).Cardinality()))).Uint32()).(*C0), 6)
		(_nw34).ArraySet1(_208_v43, 7)
		(_nw34).ArraySet1(_208_v43, 8)
		(_nw34).ArraySet1(_208_v43, 9)
		(_nw34).ArraySet1(_208_v43, 10)
		(_nw34).ArraySet1(_208_v43, 11)
		(_nw34).ArraySet1(_208_v43, 12)
		(_nw34).ArraySet1(_208_v43, 13)
		(_nw34).ArraySet1(_208_v43, 14)
		(_nw34).ArraySet1(_208_v43, 15)
		(_nw34).ArraySet1(_208_v43, 16)
		(_nw34).ArraySet1(_208_v43, 17)
		(_nw34).ArraySet1(_208_v43, 18)
		(_nw34).ArraySet1(_208_v43, 19)
		(_nw34).ArraySet1(_208_v43, 20)
		_210_v45 = _nw34
		_210_v45 = _210_v45
		(_136_globalState).F4 = (Companion_Default___.Fm2(_137_v1, _dafny.IntOfInt64(942), _140_v4, _dafny.UnicodeSeqOfUtf8Bytes("p"), _136_globalState)).Plus(_140_v4)
		var _211_v46 _dafny.Array
		_ = _211_v46
		var _out2 _dafny.Array
		_ = _out2
		_out2 = Companion_Default___.M1(_140_v4, _135_v0, _136_globalState)
		_211_v46 = _out2
		var _212_v47 D4
		_ = _212_v47
		_212_v47 = Companion_D4_.Create_DC6_(_211_v46)
		var _213_v48 _dafny.Array
		_ = _213_v48
		var _out3 _dafny.Array
		_ = _out3
		_out3 = Companion_Default___.M1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_140_v4), _140_v4), (_212_v47).Dtor_cf8(), _136_globalState)
		_213_v48 = _out3
	}
	var _214_v49 *C1
	_ = _214_v49
	var _nw35 *C1 = New_C1_()
	_ = _nw35
	_nw35.Ctor__(_155_v15)
	_214_v49 = _nw35
	_214_v49 = _214_v49
	var _215_v50 _dafny.Sequence
	_ = _215_v50
	_215_v50 = _dafny.SeqOf(_137_v1, _137_v1)
	_137_v1 = !_dafny.Companion_Sequence_.Equal(_215_v50, _215_v50)
	if _137_v1 {
		(_136_globalState).F10 = _137_v1
		var _216_v52 _dafny.Set
		_ = _216_v52
		_216_v52 = _dafny.SetOf(((func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-52), _dafny.IntOfInt64(-65))); ; {
				_compr_9, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _217_v51 _dafny.Int
				_217_v51 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(-52)).Cmp(_217_v51) <= 0) && ((_217_v51).Cmp(_dafny.IntOfInt64(-65)) < 0) {
					_coll9.Add((_217_v51).Plus(_140_v4), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_157_v17).Cardinality()))).Cardinality())
				}
			}
			return _coll9.ToMap()
		}()).Cardinality()).Times(_140_v4))
		_216_v52 = (_216_v52).Intersection(_216_v52)
		var _218_v53 _dafny.CodePoint
		_ = _218_v53
		_218_v53 = _dafny.CodePoint('y')
		var _219_v54 D2
		_ = _219_v54
		_219_v54 = Companion_D2_.Create_DC4_(_140_v4, _218_v53)
		var _220_v55 _dafny.Map
		_ = _220_v55
		_220_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(14)), (func() _dafny.Int {
			if _137_v1 {
				return _140_v4
			}
			return (_219_v54).Dtor_cf5()
		})())
		_220_v55 = (_220_v55).Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_140_v4), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.SeqOf(_140_v4)).Cardinality()))).Uint32(), _140_v4), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_140_v4), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.SeqOf(_140_v4)).Cardinality()))).Uint32(), _140_v4)).Cardinality()))).Uint32(), _140_v4), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_137_v1)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_140_v4), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.SeqOf(_140_v4)).Cardinality()))).Uint32(), _140_v4), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_140_v4), (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_dafny.SeqOf(_140_v4)).Cardinality()))).Uint32(), _140_v4)).Cardinality()))).Uint32(), _140_v4)).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v1, _137_v1)).Cardinality()), _dafny.IntOfInt64(-114))
		var _221_v56 *C0
		_ = _221_v56
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__(_218_v53)
		_221_v56 = _nw36
		var _222_v57 _dafny.Array
		_ = _222_v57
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_4
		var _nw37 _dafny.Array
		_ = _nw37
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw37 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Sequence = (func(_223_v14 _dafny.Sequence, _224_v4 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_225_i10 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_223_v14, (Companion_Default___.SafeIndex(_224_v4, _dafny.IntOfUint32((_223_v14).Cardinality()))).Uint32(), _224_v4)
				}
			})(_153_v14, _140_v4)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw37 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw37).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw37).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_222_v57 = _nw37
		var _226_v58 _dafny.Sequence
		_ = _226_v58
		_226_v58 = _dafny.SeqOf(_222_v57)
		var _227_v59 _dafny.Array
		_ = _227_v59
		var _nwElement0_9 _dafny.Array = _222_v57
		_ = _nwElement0_9
		var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(18))
		_ = _nw38
		(_nw38).ArraySet1(_nwElement0_9, 0)
		(_nw38).ArraySet1(_222_v57, 1)
		(_nw38).ArraySet1(_222_v57, 2)
		(_nw38).ArraySet1((_226_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.IntOfUint32((_226_v58).Cardinality()))).Uint32()).(_dafny.Array), 3)
		(_nw38).ArraySet1(_222_v57, 4)
		(_nw38).ArraySet1(_222_v57, 5)
		(_nw38).ArraySet1(_222_v57, 6)
		(_nw38).ArraySet1(_222_v57, 7)
		(_nw38).ArraySet1(_222_v57, 8)
		(_nw38).ArraySet1(_222_v57, 9)
		(_nw38).ArraySet1(_222_v57, 10)
		(_nw38).ArraySet1(_222_v57, 11)
		(_nw38).ArraySet1(_222_v57, 12)
		(_nw38).ArraySet1((func() _dafny.Array {
			if !(_137_v1) {
				return _222_v57
			}
			return _222_v57
		})(), 13)
		(_nw38).ArraySet1(_222_v57, 14)
		(_nw38).ArraySet1(_222_v57, 15)
		(_nw38).ArraySet1(_222_v57, 16)
		(_nw38).ArraySet1(_222_v57, 17)
		_227_v59 = _nw38
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_227_v59), 0))
		_ = _index29
		(_227_v59).ArraySet1(_222_v57, (_index29).Int())
		var _228_v60 _dafny.Array
		_ = _228_v60
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw39
		_228_v60 = _nw39
		var _229_v61 _dafny.Map
		_ = _229_v61
		_229_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_140_v4), _221_v56)
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_227_v59), 0))
		_ = _index30
		var _rhs23 *C0 = (func() *C0 {
			if (_229_v61).Contains(_dafny.IntOfUint32((Companion_Default___.Fm8(_137_v1, _157_v17, _136_globalState)).Cardinality())) {
				return (_229_v61).Get(_dafny.IntOfUint32((Companion_Default___.Fm8(_137_v1, _157_v17, _136_globalState)).Cardinality())).(*C0)
			}
			return (func() *C0 {
				if _137_v1 {
					return _221_v56
				}
				return _221_v56
			})()
		})()
		_ = _rhs23
		var _rhs24 bool = (func() bool {
			if (_158_v18).Contains((func() bool {
				if _137_v1 {
					return _137_v1
				}
				return _137_v1
			})()) {
				return (_158_v18).Get((func() bool {
					if _137_v1 {
						return _137_v1
					}
					return _137_v1
				})()).(bool)
			}
			return _137_v1
		})()
		_ = _rhs24
		var _rhs25 _dafny.Array = _222_v57
		_ = _rhs25
		var _rhs26 _dafny.Array = _228_v60
		_ = _rhs26
		var _rhs27 _dafny.Int = _140_v4
		_ = _rhs27
		var _lhs20 *GlobalState = _136_globalState
		_ = _lhs20
		var _lhs21 _dafny.Array = _227_v59
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen((_227_v59), 0))
		_ = _lhs22
		var _lhs23 *GlobalState = _136_globalState
		_ = _lhs23
		_221_v56 = _rhs23
		_lhs20.F10 = _rhs24
		(_lhs21).ArraySet1(_rhs25, (_lhs22).Int())
		_228_v60 = _rhs26
		_lhs23.F4 = _rhs27
		var _230_v62 *C1
		_ = _230_v62
		var _nw40 *C1 = New_C1_()
		_ = _nw40
		_nw40.Ctor__((_214_v49).F11())
		_230_v62 = _nw40
	} else {
		var _231_v63 _dafny.Array
		_ = _231_v63
		var _nwElement0_10 _dafny.Array = _135_v0
		_ = _nwElement0_10
		var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
		_ = _nw41
		(_nw41).ArraySet1(_nwElement0_10, 0)
		(_nw41).ArraySet1(_135_v0, 1)
		(_nw41).ArraySet1(_135_v0, 2)
		(_nw41).ArraySet1(_135_v0, 3)
		(_nw41).ArraySet1(_135_v0, 4)
		(_nw41).ArraySet1(_135_v0, 5)
		(_nw41).ArraySet1(_135_v0, 6)
		(_nw41).ArraySet1(_135_v0, 7)
		(_nw41).ArraySet1(_135_v0, 8)
		(_nw41).ArraySet1(_135_v0, 9)
		(_nw41).ArraySet1(_135_v0, 10)
		(_nw41).ArraySet1(_135_v0, 11)
		(_nw41).ArraySet1(_135_v0, 12)
		_231_v63 = _nw41
		_231_v63 = _231_v63
		var _232_v64 _dafny.Sequence
		_ = _232_v64
		_232_v64 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("jfb"))
		var _233_v65 _dafny.Map
		_ = _233_v65
		_233_v65 = _151_v12
		if Companion_Default___.Fm0(_137_v1, _dafny.Companion_Sequence_.IsProperPrefixOf(_157_v17, (_232_v64).Select((Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_232_v64).Cardinality()))).Uint32()).(_dafny.Sequence)), _dafny.IntOfInt64(-756), Companion_Default___.Fm2(Companion_Default___.Fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState), _dafny.IntOfUint32((_dafny.SeqOf(_233_v65)).Cardinality()), _dafny.IntOfUint32((_215_v50).Cardinality()), _157_v17, _136_globalState), _136_globalState) {
			var _234_v66 _dafny.CodePoint
			_ = _234_v66
			_234_v66 = _dafny.CodePoint('w')
			var _235_v67 *C0
			_ = _235_v67
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__(_234_v66)
			_235_v67 = _nw42
			(_136_globalState).F4 = _140_v4
			var _236_v69 _dafny.Array
			_ = _236_v69
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_5
			var _nw43 _dafny.Array
			_ = _nw43
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw43 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Map = (func(_237_v4 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_238_i11 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_237_v4, _237_v4)).Merge(func() _dafny.Map {
							var _coll10 = _dafny.NewMapBuilder()
							_ = _coll10
							for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-874), _dafny.IntOfInt64(754))); ; {
								_compr_10, _ok11 := _iter11()
								if !_ok11 {
									break
								}
								var _239_v68 _dafny.Int
								_239_v68 = interface{}(_compr_10).(_dafny.Int)
								if ((_dafny.IntOfInt64(-874)).Cmp(_239_v68) <= 0) && ((_239_v68).Cmp(_dafny.IntOfInt64(754)) < 0) {
									_coll10.Add(Companion_Default___.SafeModuloInt(_239_v68, _237_v4), _237_v4)
								}
							}
							return _coll10.ToMap()
						}())
					}
				})(_140_v4)
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
			_236_v69 = _nw43
			var _240_v70 _dafny.Map
			_ = _240_v70
			_240_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v4, _140_v4)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_236_v69), 0))
			_ = _index31
			(_236_v69).ArraySet1(_240_v70, (_index31).Int())
			var _241_v71 _dafny.Sequence
			_ = _241_v71
			_241_v71 = _dafny.SeqOf(_214_v49)
			var _242_v72 _dafny.Map
			_ = _242_v72
			_242_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v4, _135_v0)
			var _243_v73 D5
			_ = _243_v73
			_243_v73 = Companion_D5_.Create_DC8_(_242_v72)
			var _244_v74 _dafny.Map
			_ = _244_v74
			_244_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_241_v71), (((_243_v73).Dtor_cf11()).Merge(_242_v72)).Cardinality())
			var _245_v77 _dafny.Map
			_ = _245_v77
			_245_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_235_v67).F12(), _137_v1)
			var _246_v78 _dafny.Map
			_ = _246_v78
			_246_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v2, func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter12 := _dafny.Iterate((_245_v77).Keys().Elements()); ; {
					_compr_11, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _247_v76 _dafny.CodePoint
					_247_v76 = interface{}(_compr_11).(_dafny.CodePoint)
					if (_245_v77).Contains(_247_v76) {
						_coll11.Add(_247_v76, _140_v4)
					}
				}
				return _coll11.ToMap()
			}())
			var _248_v79 _dafny.Set
			_ = _248_v79
			_248_v79 = _dafny.SetOf(_dafny.IntOfUint32((_157_v17).Cardinality()), _140_v4, _140_v4)
			var _249_v80 _dafny.Array
			_ = _249_v80
			var _nw44 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
			_ = _nw44
			_249_v80 = _nw44
			var _250_v81 _dafny.Map
			_ = _250_v81
			_250_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v1, _249_v80)
			var _251_v82 _dafny.Map
			_ = _251_v82
			_251_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_250_v81, _244_v74)
			var _252_v83 _dafny.MultiSet
			_ = _252_v83
			_252_v83 = _dafny.MultiSetOf(_214_v49)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_236_v69), 0))
			_ = _index32
			var _rhs28 _dafny.Map = _240_v70
			_ = _rhs28
			var _rhs29 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v2, _137_v1)).Merge((func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter13 := _dafny.Iterate((_246_v78).Keys().Elements()); ; {
					_compr_12, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _253_v75 _dafny.MultiSet
					_253_v75 = interface{}(_compr_12).(_dafny.MultiSet)
					if (_246_v78).Contains(_253_v75) {
						_coll12.Add(_253_v75, _137_v1)
					}
				}
				return _coll12.ToMap()
			}()).Merge(Companion_Default___.Fm14(true, _140_v4, _248_v79, _136_globalState)))
			_ = _rhs29
			var _rhs30 _dafny.Map = (func() _dafny.Map {
				if (_251_v82).Contains(_250_v81) {
					return (_251_v82).Get(_250_v81).(_dafny.Map)
				}
				return (_244_v74).Update(_252_v83, _140_v4)
			})()
			_ = _rhs30
			var _lhs24 _dafny.Array = _236_v69
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_236_v69), 0))
			_ = _lhs25
			(_lhs24).ArraySet1(_rhs28, (_lhs25).Int())
			_139_v3 = _rhs29
			_244_v74 = _rhs30
			var _254_v84 _dafny.Array
			_ = _254_v84
			var _out4 _dafny.Array
			_ = _out4
			_out4 = Companion_Default___.M1(_140_v4, _135_v0, _136_globalState)
			_254_v84 = _out4
			var _255_v85 D2
			_ = _255_v85
			_255_v85 = Companion_D2_.Create_DC2_((_235_v67).F12())
			var _256_v86 _dafny.MultiSet
			_ = _256_v86
			_256_v86 = _dafny.MultiSetOf(_255_v85)
			(_136_globalState).F10 = !((_256_v86).IsProperSubsetOf((_256_v86).Intersection((_256_v86).Update(_255_v85, Companion_Default___.Abs(_140_v4)))))
		} else {
			var _257_v87 _dafny.Map
			_ = _257_v87
			_257_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if _137_v1 {
					return _137_v1
				}
				return _137_v1
			})(), _140_v4)
			var _258_v88 _dafny.CodePoint
			_ = _258_v88
			_258_v88 = _dafny.CodePoint('n')
			var _rhs31 bool = !(_137_v1)
			_ = _rhs31
			var _rhs32 _dafny.Int = _140_v4
			_ = _rhs32
			var _rhs33 bool = _dafny.Companion_Sequence_.Equal(_157_v17, _dafny.Companion_Sequence_.Update(_157_v17, (Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_157_v17).Cardinality()))).Uint32(), _258_v88))
			_ = _rhs33
			var _rhs34 _dafny.Array = _174_v24
			_ = _rhs34
			var _rhs35 _dafny.Map = (_257_v87).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v1, _140_v4))
			_ = _rhs35
			var _lhs26 *GlobalState = _136_globalState
			_ = _lhs26
			var _lhs27 *GlobalState = _136_globalState
			_ = _lhs27
			var _lhs28 *GlobalState = _136_globalState
			_ = _lhs28
			_lhs26.F10 = _rhs31
			_lhs27.F4 = _rhs32
			_lhs28.F10 = _rhs33
			_174_v24 = _rhs34
			_257_v87 = _rhs35
			var _259_v89 _dafny.Array
			_ = _259_v89
			var _nw45 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
			_ = _nw45
			_259_v89 = _nw45
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_259_v89), 0))
			_ = _index33
			(_259_v89).ArraySet1((func() *C1 {
				if _137_v1 {
					return _214_v49
				}
				return _214_v49
			})(), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_174_v24), 0))
			_ = _index34
			(_174_v24).ArraySet1((Companion_Default___.Fm15(_136_globalState)).Cardinality(), (_index34).Int())
			var _260_v90 _dafny.Set
			_ = _260_v90
			_260_v90 = _dafny.SetOf(_135_v0)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_259_v89), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_174_v24), 0))
			_ = _index36
			var _rhs36 *C1 = _214_v49
			_ = _rhs36
			var _rhs37 bool = Companion_Default___.Fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState)
			_ = _rhs37
			var _rhs38 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.Fm2(_137_v1, _140_v4, _140_v4, _157_v17, _136_globalState), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-324))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_261_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_262_i12 _dafny.Int) _dafny.Int {
					return _261_v4
				}
			})(_140_v4)))).Cardinality())).Times(_140_v4))
			_ = _rhs38
			var _rhs39 _dafny.Set = _260_v90
			_ = _rhs39
			var _rhs40 bool = true
			_ = _rhs40
			var _lhs29 _dafny.Array = _259_v89
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(264), _dafny.ArrayLen((_259_v89), 0))
			_ = _lhs30
			var _lhs31 _dafny.Array = _174_v24
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_174_v24), 0))
			_ = _lhs32
			var _lhs33 *GlobalState = _136_globalState
			_ = _lhs33
			(_lhs29).ArraySet1(_rhs36, (_lhs30).Int())
			_137_v1 = _rhs37
			(_lhs31).ArraySet1(_rhs38, (_lhs32).Int())
			_260_v90 = _rhs39
			_lhs33.F10 = _rhs40
			var _263_v91 _dafny.Array
			_ = _263_v91
			var _out5 _dafny.Array
			_ = _out5
			_out5 = Companion_Default___.M1((_174_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_174_v24), 0))).Int()).(_dafny.Int), _135_v0, _136_globalState)
			_263_v91 = _out5
			var _264_v92 _dafny.Sequence
			_ = _264_v92
			_264_v92 = _dafny.SeqOf(_135_v0, _263_v91, _263_v91, _135_v0)
			var _265_v93 _dafny.Array
			_ = _265_v93
			var _out6 _dafny.Array
			_ = _out6
			_out6 = Companion_Default___.M1(_dafny.IntOfUint32((_157_v17).Cardinality()), (_264_v92).Select((Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_264_v92).Cardinality()))).Uint32()).(_dafny.Array), _136_globalState)
			_265_v93 = _out6
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_174_v24), 0))
			_ = _index37
			(_174_v24).ArraySet1(_dafny.IntOfInt64(412), (_index37).Int())
		}
		_140_v4 = (_153_v14).Select((Companion_Default___.SafeIndex(_140_v4, _dafny.IntOfUint32((_153_v14).Cardinality()))).Uint32()).(_dafny.Int)
		var _266_v94 _dafny.MultiSet
		_ = _266_v94
		_266_v94 = _dafny.MultiSetOf(_140_v4, _140_v4)
		var _267_v95 _dafny.Sequence
		_ = _267_v95
		_267_v95 = _dafny.SeqOf(_266_v94)
		var _268_v96 _dafny.Map
		_ = _268_v96
		_268_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v95, !(!((func() bool {
			if _137_v1 {
				return _137_v1
			}
			return _137_v1
		})())))
		_268_v96 = (_268_v96).Update(_267_v95, _137_v1)
		_140_v4 = _140_v4
	}
	_140_v4 = _140_v4
	(_136_globalState).F4 = _140_v4
	if _137_v1 {
		var _269_v97 _dafny.Array
		_ = _269_v97
		var _out7 _dafny.Array
		_ = _out7
		_out7 = Companion_Default___.M1(_140_v4, _135_v0, _136_globalState)
		_269_v97 = _out7
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_174_v24), 0))
		_ = _index38
		(_174_v24).ArraySet1(_140_v4, (_index38).Int())
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_174_v24), 0))
		_ = _index39
		(_174_v24).ArraySet1(_140_v4, (_index39).Int())
		(_136_globalState).F10 = _137_v1
		var _270_v98 *C1
		_ = _270_v98
		var _nw46 *C1 = New_C1_()
		_ = _nw46
		_nw46.Ctor__((_214_v49).F11())
		_270_v98 = _nw46
		(_136_globalState).F4 = (func() _dafny.Int {
			if _137_v1 {
				return _140_v4
			}
			return _140_v4
		})()
	} else {
		var _271_v99 _dafny.CodePoint
		_ = _271_v99
		_271_v99 = _dafny.CodePoint('r')
		var _272_v100 _dafny.Map
		_ = _272_v100
		_272_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(_140_v4, _136_globalState), _dafny.IntOfInt64(692))
		var _273_v101 _dafny.Map
		_ = _273_v101
		_273_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v0, _158_v18)
		var _274_v102 _dafny.Map
		_ = _274_v102
		_274_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			if _137_v1 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_271_v99, _dafny.IntOfInt64(-509))
			}
			return _272_v100
		})(), (_273_v101).Cardinality())
		_274_v102 = (_274_v102).Update(_272_v100, _140_v4)
		(_136_globalState).F4 = _140_v4
		(_136_globalState).F3 = _135_v0
		var _275_v103 _dafny.Sequence
		_ = _275_v103
		var _out8 _dafny.Sequence
		_ = _out8
		_out8 = (_214_v49).M0(_140_v4, _dafny.Companion_Sequence_.Concatenate(_157_v17, _157_v17), !(_137_v1), _140_v4, _136_globalState)
		_275_v103 = _out8
		(_136_globalState).F10 = ((_214_v49).F11())
	}
	(_136_globalState).F4 = Companion_Default___.Fm2(_137_v1, _140_v4, _140_v4, _157_v17, _136_globalState)
	_dafny.Print((_135_v0).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_136_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState.F3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_136_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_136_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_136_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v2).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), false).UpdateUnsafe(_dafny.MultiSetOf(false), false).UpdateUnsafe(_dafny.MultiSetOf(true, false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_140_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(123), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("du"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_153_v14, _dafny.SeqOf(_dafny.IntOfInt64(384), _dafny.IntOfInt64(670))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v15))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(384))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_v17.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v18).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_174_v24).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_214_v49).F11()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_215_v50, _dafny.SeqOf(false, false)))
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

func (CompanionStruct_D0_) Default() bool {
	return false
}

func (_this D0) Dtor_cf0() bool {
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

type D1_DC1 struct {
	Cf1 _dafny.Map
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Map) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D1) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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

type D2_DC3 struct {
	Cf3 bool
	Cf4 bool
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf3 bool, Cf4 bool) D2 {
	return D2{D2_DC3{Cf3, Cf4}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

type D2_DC4 struct {
	Cf5 _dafny.Int
	Cf6 _dafny.CodePoint
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.Int, Cf6 _dafny.CodePoint) D2 {
	return D2{D2_DC4{Cf5, Cf6}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC2 struct {
	Cf2 _dafny.CodePoint
}

func (D2_DC2) isD2() {}

func (CompanionStruct_D2_) Create_DC2_(Cf2 _dafny.CodePoint) D2 {
	return D2{D2_DC2{Cf2}}
}

func (_this D2) Is_DC2() bool {
	_, ok := _this.Get_().(D2_DC2)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC3_(false, false)
}

func (_this D2) Dtor_cf3() bool {
	return _this.Get_().(D2_DC3).Cf3
}

func (_this D2) Dtor_cf4() bool {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf2() _dafny.CodePoint {
	return _this.Get_().(D2_DC2).Cf2
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC2:
		{
			return "D2.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6
		}
	case D2_DC2:
		{
			data2, ok := other.Get_().(D2_DC2)
			return ok && data1.Cf2 == data2.Cf2
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

type D3_DC5 struct {
	Cf7 _dafny.Map
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf7 _dafny.Map) D3 {
	return D3{D3_DC5{Cf7}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D3_DC5).Cf7
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf7) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf7.Equals(data2.Cf7)
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

type D4_DC7 struct {
	Cf9  _dafny.Int
	Cf10 bool
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf9 _dafny.Int, Cf10 bool) D4 {
	return D4{D4_DC7{Cf9, Cf10}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

type D4_DC6 struct {
	Cf8 _dafny.Array
}

func (D4_DC6) isD4() {}

func (CompanionStruct_D4_) Create_DC6_(Cf8 _dafny.Array) D4 {
	return D4{D4_DC6{Cf8}}
}

func (_this D4) Is_DC6() bool {
	_, ok := _this.Get_().(D4_DC6)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC7_(_dafny.Zero, false)
}

func (_this D4) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D4_DC7).Cf9
}

func (_this D4) Dtor_cf10() bool {
	return _this.Get_().(D4_DC7).Cf10
}

func (_this D4) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D4_DC6).Cf8
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC7:
		{
			return "D4.DC7" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D4_DC6:
		{
			return "D4.DC6" + "(" + _dafny.String(data.Cf8) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10
		}
	case D4_DC6:
		{
			data2, ok := other.Get_().(D4_DC6)
			return ok && data1.Cf8 == data2.Cf8
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

type D5_DC9 struct {
	Cf12 _dafny.Sequence
	Cf13 bool
	Cf14 *C1
	Cf15 _dafny.CodePoint
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf12 _dafny.Sequence, Cf13 bool, Cf14 *C1, Cf15 _dafny.CodePoint) D5 {
	return D5{D5_DC9{Cf12, Cf13, Cf14, Cf15}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

type D5_DC8 struct {
	Cf11 _dafny.Map
}

func (D5_DC8) isD5() {}

func (CompanionStruct_D5_) Create_DC8_(Cf11 _dafny.Map) D5 {
	return D5{D5_DC8{Cf11}}
}

func (_this D5) Is_DC8() bool {
	_, ok := _this.Get_().(D5_DC8)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC9_(_dafny.EmptySeq, false, (*C1)(nil), _dafny.CodePoint('D'))
}

func (_this D5) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D5_DC9).Cf12
}

func (_this D5) Dtor_cf13() bool {
	return _this.Get_().(D5_DC9).Cf13
}

func (_this D5) Dtor_cf14() *C1 {
	return _this.Get_().(D5_DC9).Cf14
}

func (_this D5) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D5_DC9).Cf15
}

func (_this D5) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D5_DC8).Cf11
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC9:
		{
			return "D5.DC9" + "(" + data.Cf12.VerbatimString(true) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D5_DC8:
		{
			return "D5.DC8" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC9:
		{
			data2, ok := other.Get_().(D5_DC9)
			return ok && data1.Cf12.Equals(data2.Cf12) && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D5_DC8:
		{
			data2, ok := other.Get_().(D5_DC8)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D6_DC11 struct {
	Cf17 _dafny.Sequence
	Cf18 *C0
	Cf19 _dafny.Int
	Cf20 _dafny.Sequence
	Cf21 _dafny.Int
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf17 _dafny.Sequence, Cf18 *C0, Cf19 _dafny.Int, Cf20 _dafny.Sequence, Cf21 _dafny.Int) D6 {
	return D6{D6_DC11{Cf17, Cf18, Cf19, Cf20, Cf21}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

type D6_DC12 struct {
	Cf22 bool
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf22 bool) D6 {
	return D6{D6_DC12{Cf22}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC13 struct {
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_() D6 {
	return D6{D6_DC13{}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC10 struct {
	Cf16 _dafny.Map
}

func (D6_DC10) isD6() {}

func (CompanionStruct_D6_) Create_DC10_(Cf16 _dafny.Map) D6 {
	return D6{D6_DC10{Cf16}}
}

func (_this D6) Is_DC10() bool {
	_, ok := _this.Get_().(D6_DC10)
	return ok
}

type D6_DC14 struct {
	Cf23 D6
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf23 D6) D6 {
	return D6{D6_DC14{Cf23}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC11_(_dafny.EmptySeq, (*C0)(nil), _dafny.Zero, _dafny.EmptySeq, _dafny.Zero)
}

func (_this D6) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D6_DC11).Cf17
}

func (_this D6) Dtor_cf18() *C0 {
	return _this.Get_().(D6_DC11).Cf18
}

func (_this D6) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf19
}

func (_this D6) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D6_DC11).Cf20
}

func (_this D6) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D6_DC11).Cf21
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC12).Cf22
}

func (_this D6) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D6_DC10).Cf16
}

func (_this D6) Dtor_cf23() D6 {
	return _this.Get_().(D6_DC14).Cf23
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC11:
		{
			return "D6.DC11" + "(" + data.Cf17.VerbatimString(true) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13"
		}
	case D6_DC10:
		{
			return "D6.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17) && data1.Cf18 == data2.Cf18 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Equals(data2.Cf20) && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf22 == data2.Cf22
		}
	case D6_DC13:
		{
			_, ok := other.Get_().(D6_DC13)
			return ok
		}
	case D6_DC10:
		{
			data2, ok := other.Get_().(D6_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
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

type D7_DC16 struct {
	Cf25 _dafny.Sequence
	Cf26 _dafny.MultiSet
	Cf27 bool
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf25 _dafny.Sequence, Cf26 _dafny.MultiSet, Cf27 bool) D7 {
	return D7{D7_DC16{Cf25, Cf26, Cf27}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC15 struct {
	Cf24 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf24 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf24}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC16_(_dafny.EmptySeq, _dafny.EmptyMultiSet, false)
}

func (_this D7) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D7_DC16).Cf25
}

func (_this D7) Dtor_cf26() _dafny.MultiSet {
	return _this.Get_().(D7_DC16).Cf26
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC16).Cf27
}

func (_this D7) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf24
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC16:
		{
			return "D7.DC16" + "(" + data.Cf25.VerbatimString(true) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf25.Equals(data2.Cf25) && data1.Cf26.Equals(data2.Cf26) && data1.Cf27 == data2.Cf27
		}
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
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
	Cf29 _dafny.Int
	Cf30 bool
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf29 _dafny.Int, Cf30 bool) D8 {
	return D8{D8_DC18{Cf29, Cf30}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

type D8_DC17 struct {
	Cf28 _dafny.MultiSet
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf28 _dafny.MultiSet) D8 {
	return D8{D8_DC17{Cf28}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC19 struct {
	Cf31 D8
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf31 D8) D8 {
	return D8{D8_DC19{Cf31}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC18_(_dafny.Zero, false)
}

func (_this D8) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D8_DC18).Cf29
}

func (_this D8) Dtor_cf30() bool {
	return _this.Get_().(D8_DC18).Cf30
}

func (_this D8) Dtor_cf28() _dafny.MultiSet {
	return _this.Get_().(D8_DC17).Cf28
}

func (_this D8) Dtor_cf31() D8 {
	return _this.Get_().(D8_DC19).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30 == data2.Cf30
		}
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	F3  _dafny.Array
	F4  _dafny.Int
	F9  bool
	F10 bool
	_f0 _dafny.Array
	_f1 _dafny.Int
	_f2 bool
	_f5 _dafny.Int
	_f6 bool
	_f7 _dafny.Int
	_f8 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F4 = _dafny.Zero
	_this.F9 = false
	_this.F10 = false
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f8 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Int, f2 bool, f3 _dafny.Array, f4 _dafny.Int, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 bool, f9 bool, f10 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this).F10 = f10
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
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 _dafny.CodePoint
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = _dafny.CodePoint('D')
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

func (_this *C0) Ctor__(f12 _dafny.CodePoint) {
	{
		(_this)._f12 = f12
	}
}
func (_this *C0) Fm3(p0 _dafny.Set, globalState *GlobalState) bool {
	{
		if (func() bool {
			if !(true) {
				return false
			}
			return true
		})() {
			return false
		} else {
			return false
		}
	}
}
func (_this *C0) F12() _dafny.CodePoint {
	{
		return _this._f12
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f11 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f11 = false
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

func (_this *C1) Ctor__(f11 bool) {
	{
		(_this)._f11 = f11
	}
}
func (_this *C1) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _276_v0 _dafny.Sequence
		_ = _276_v0
		_276_v0 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()), p3)
		var _277_i0 _dafny.Int
		_ = _277_i0
		_277_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsProperPrefixOf(_276_v0, _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-33)))) {
				{
					if (_277_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_277_i0 = (_277_i0).Plus(_dafny.One)
					(globalState).F4 = (_dafny.Zero).Minus(p3)
					var _278_v1 _dafny.Map
					_ = _278_v1
					_278_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-768), p1)
					_278_v1 = (_278_v1).Update(p0, p1)
					(globalState).F10 = Companion_Default___.Fm0(p2, (func() bool {
						if p2 {
							return p2
						}
						return true
					})(), p3, Companion_Default___.SafeDivisionInt(p3, p3), globalState)
					var _279_v2 _dafny.Array
					_ = _279_v2
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(26)
					_ = _len0_6
					var _nw47 _dafny.Array
					_ = _nw47
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw47 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Sequence = (func(_280_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_281_i1 _dafny.Int) _dafny.Sequence {
								return _280_p1
							}
						})(p1)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw47 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw47).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw47).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_279_v2 = _nw47
					var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_279_v2), 0))
					_ = _index40
					(_279_v2).ArraySet1(p1, (_index40).Int())
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_279_v2), 0))
					_ = _index41
					(_279_v2).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tqgbnrs"), (_index41).Int())
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		(globalState).F10 = p2
		var _hi1 _dafny.Int = p3
		_ = _hi1
		for _282_i2 := p3; _282_i2.Cmp(_hi1) < 0; _282_i2 = _282_i2.Plus(_dafny.One) {
			var _283_v3 _dafny.CodePoint
			_ = _283_v3
			_283_v3 = _dafny.CodePoint('p')
			var _284_v4 _dafny.Map
			_ = _284_v4
			_284_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v3, p2)).Cardinality(), p2)
			var _285_v5 _dafny.Sequence
			_ = _285_v5
			_285_v5 = _dafny.SeqOf(p1)
			var _286_v6 _dafny.Array
			_ = _286_v6
			var _nwElement0_11 _dafny.Int = p3
			_ = _nwElement0_11
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(6))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_11, 0)
			(_nw48).ArraySet1(_282_i2, 1)
			(_nw48).ArraySet1(_282_i2, 2)
			(_nw48).ArraySet1(Companion_Default___.Fm2(p2, (_dafny.MultiSetOf(p2, p2, p2)).Cardinality(), (_284_v4).Cardinality(), (_285_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_285_v5).Cardinality()))).Uint32()).(_dafny.Sequence), globalState), 3)
			(_nw48).ArraySet1(p3, 4)
			(_nw48).ArraySet1((p0).Minus(p3), 5)
			_286_v6 = _nw48
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))
			_ = _index42
			(_286_v6).ArraySet1(_282_i2, (_index42).Int())
			var _287_v7 _dafny.Sequence
			_ = _287_v7
			_287_v7 = _dafny.SeqOf(p2)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))
			_ = _index43
			(_286_v6).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_287_v7).Cardinality())).Minus(_282_i2)), (_index43).Int())
			if p2 {
				var _288_v8 _dafny.MultiSet
				_ = _288_v8
				_288_v8 = _dafny.MultiSetOf(_dafny.IntOfInt64(866), _282_i2)
				(globalState).F4 = ((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus(((_dafny.Zero).Minus((func() _dafny.Int {
					if (_288_v8).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _283_v3)).Cardinality())) {
						return (_288_v8).Multiplicity(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _283_v3)).Cardinality()))
					}
					return _282_i2
				})())).Times((_dafny.Zero).Minus(p3))))
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))
				_ = _index44
				(_286_v6).ArraySet1(p3, (_index44).Int())
				(globalState).F9 = !(p2) || (Companion_Default___.Fm0(p2, false, (_276_v0).Select((Companion_Default___.SafeIndex((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_276_v0).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32((_287_v7).Cardinality())), globalState))
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))
				_ = _index45
				(_286_v6).ArraySet1(((func() _dafny.Int {
					if (_288_v8).Contains(_dafny.IntOfUint32((p1).Cardinality())) {
						return (_288_v8).Multiplicity(_dafny.IntOfUint32((p1).Cardinality()))
					}
					return (_dafny.Zero).Minus(p3)
				})()).Times((_276_v0).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_276_v0).Cardinality()))).Uint32()).(_dafny.Int)), (_index45).Int())
				var _289_v9 _dafny.Map
				_ = _289_v9
				_289_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _276_v0)
				(globalState).F10 = _dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
					if (_289_v9).Contains(p2) {
						return (_289_v9).Get(p2).(_dafny.Sequence)
					}
					return _276_v0
				})(), _276_v0)
			} else {
				var _290_v10 *C0
				_ = _290_v10
				var _nw49 *C0 = New_C0_()
				_ = _nw49
				_nw49.Ctor__(_283_v3)
				_290_v10 = _nw49
				var _291_v11 _dafny.Array
				_ = _291_v11
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
				_ = _nw50
				_291_v11 = _nw50
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_291_v11), 0))
				_ = _index46
				(_291_v11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_i2, p2), (_index46).Int())
				var _292_v12 _dafny.Map
				_ = _292_v12
				_292_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_282_i2, p2)
				var _293_v13 _dafny.Sequence
				_ = _293_v13
				_293_v13 = _dafny.SeqOf(_284_v4, (_292_v12).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)), (_284_v4).Merge(_284_v4), _284_v4, _284_v4)
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_291_v11), 0))
				_ = _index47
				(_291_v11).ArraySet1((_293_v13).Select((Companion_Default___.SafeIndex(((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int)).Times(p3), _dafny.IntOfUint32((_293_v13).Cardinality()))).Uint32()).(_dafny.Map), (_index47).Int())
				var _294_v14 *C0
				_ = _294_v14
				var _nw51 *C0 = New_C0_()
				_ = _nw51
				_nw51.Ctor__((_290_v10).F12())
				_294_v14 = _nw51
				var _295_v15 _dafny.Array
				_ = _295_v15
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_7
				var _nw52 _dafny.Array
				_ = _nw52
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw52 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Sequence = (func(_296_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_297_i3 _dafny.Int) _dafny.Sequence {
							return _296_p1
						}
					})(p1)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw52 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw52).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw52).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_295_v15 = _nw52
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_295_v15), 0))
				_ = _index48
				(_295_v15).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-706))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}((func(_298_p3 _dafny.Int, _299_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_300_i4 _dafny.Int) _dafny.CodePoint {
						return (Companion_D2_.Create_DC4_(_298_p3, _299_v3)).Dtor_cf6()
					}
				})(p3, _283_v3))), (_index48).Int())
				var _301_v16 _dafny.Set
				_ = _301_v16
				_301_v16 = _dafny.SetOf(p3)
				var _302_v17 _dafny.Sequence
				_ = _302_v17
				_302_v17 = _dafny.SeqOf(_286_v6)
				var _303_v18 _dafny.Map
				_ = _303_v18
				_303_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v17).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_302_v17).Cardinality()))).Uint32()).(_dafny.Array), _dafny.IntOfUint32((_276_v0).Cardinality()))
				var _304_v19 _dafny.MultiSet
				_ = _304_v19
				_304_v19 = _dafny.MultiSetOf(p2)
				var _305_v20 _dafny.Map
				_ = _305_v20
				_305_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(615), _282_i2)
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_295_v15), 0))
				_ = _index49
				var _rhs41 _dafny.Sequence = p1
				_ = _rhs41
				var _rhs42 bool = (_301_v16).IsSubsetOf(_301_v16)
				_ = _rhs42
				var _rhs43 _dafny.Int = (func() _dafny.Int {
					if !(_303_v18).Equals(_303_v18) {
						return (func() _dafny.Int {
							if (_304_v19).Contains((_287_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_276_v0).Cardinality()), _dafny.IntOfUint32((_287_v7).Cardinality()))).Uint32()).(bool)) {
								return (_304_v19).Multiplicity((_287_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_276_v0).Cardinality()), _dafny.IntOfUint32((_287_v7).Cardinality()))).Uint32()).(bool))
							}
							return _282_i2
						})()
					}
					return (func() _dafny.Int {
						if (_305_v20).Contains(p0) {
							return (_305_v20).Get(p0).(_dafny.Int)
						}
						return p3
					})()
				})()
				_ = _rhs43
				var _rhs44 bool = !(p2)
				_ = _rhs44
				var _lhs34 _dafny.Array = _295_v15
				_ = _lhs34
				var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_295_v15), 0))
				_ = _lhs35
				var _lhs36 *GlobalState = globalState
				_ = _lhs36
				var _lhs37 *GlobalState = globalState
				_ = _lhs37
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				(_lhs34).ArraySet1(_rhs41, (_lhs35).Int())
				_lhs36.F9 = _rhs42
				_lhs37.F4 = _rhs43
				_lhs38.F9 = _rhs44
				var _306_v21 _dafny.MultiSet
				_ = _306_v21
				_306_v21 = _dafny.MultiSetOf(_290_v10, _290_v10)
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_295_v15), 0))
				_ = _index50
				(_295_v15).ArraySet1(Companion_Default___.Fm4((_dafny.Zero).Minus((_306_v21).Cardinality()), p2, !(p2), ((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int)).Times((_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int)), globalState), (_index50).Int())
			}
			var _307_v22 _dafny.Map
			_ = _307_v22
			_307_v22 = _284_v4
			var _308_v23 _dafny.Sequence
			_ = _308_v23
			_308_v23 = _dafny.SeqOf(_307_v22, _307_v22)
			(globalState).F10 = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
				if p2 {
					return _308_v23
				}
				return _308_v23
			})(), _308_v23)
			(globalState).F4 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm2(p2, (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(199))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_309_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_310_i5 _dafny.Int) _dafny.CodePoint {
					return _309_v3
				}
			})(_283_v3)))).Cardinality()), p1, globalState), _282_i2), (_286_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.ArrayLen((_286_v6), 0))).Int()).(_dafny.Int))
		}
		if p2 {
			var _311_v24 *C0
			_ = _311_v24
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__(Companion_Default___.Fm5((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm4(p3, p2, p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aroliv")).Cardinality()), globalState)).Cardinality())), globalState))
			_311_v24 = _nw53
			var _312_v25 _dafny.Sequence
			_ = _312_v25
			_312_v25 = _dafny.SeqOf(_276_v0)
			_312_v25 = _312_v25
			(globalState).F10 = !((false) == ((_dafny.IntOfInt64(636)).Cmp(p0) > 0))
			var _313_v26 _dafny.Array
			_ = _313_v26
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw54
			_313_v26 = _nw54
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_313_v26), 0))
			_ = _index51
			(_313_v26).ArraySet1(p0, (_index51).Int())
			var _314_v27 _dafny.Set
			_ = _314_v27
			_314_v27 = _dafny.SetOf(p2, p2, p2, !(p2))
			var _315_v28 _dafny.Array
			_ = _315_v28
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw55
			_315_v28 = _nw55
			var _316_v29 _dafny.Sequence
			_ = _316_v29
			_316_v29 = _dafny.SeqOf(p2)
			var _317_v30 D2
			_ = _317_v30
			_317_v30 = Companion_D2_.Create_DC3_(p2, p2)
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_313_v26), 0))
			_ = _index52
			var _rhs45 _dafny.Int = p3
			_ = _rhs45
			var _rhs46 bool = (_314_v27).IsSubsetOf(_314_v27)
			_ = _rhs46
			var _rhs47 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p0, p0))
			_ = _rhs47
			var _rhs48 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, p0)
			_ = _rhs48
			var _rhs49 bool = (_dafny.IntOfUint32((Companion_Default___.Fm4((_dafny.MultiSetOf(_315_v28)).Cardinality(), p2, (_316_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_316_v29).Cardinality()))).Uint32()).(bool), p0, globalState)).Cardinality())).Cmp((func() _dafny.Int {
				if (_317_v30).Dtor_cf4() {
					return p3
				}
				return p0
			})()) > 0
			_ = _rhs49
			var _lhs39 _dafny.Array = _313_v26
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_313_v26), 0))
			_ = _lhs40
			var _lhs41 *GlobalState = globalState
			_ = _lhs41
			var _lhs42 *GlobalState = globalState
			_ = _lhs42
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			(_lhs39).ArraySet1(_rhs45, (_lhs40).Int())
			_lhs41.F9 = _rhs46
			_lhs42.F4 = _rhs47
			_lhs43.F4 = _rhs48
			_lhs44.F9 = _rhs49
			var _318_v31 _dafny.MultiSet
			_ = _318_v31
			_318_v31 = _dafny.MultiSetOf(p2, p2, p2, p2, !(p2))
			var _319_v32 _dafny.Map
			_ = _319_v32
			_319_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_318_v31, _311_v24)
			var _320_v33 _dafny.Map
			_ = _320_v33
			_320_v33 = _319_v32
			var _321_v34 _dafny.Map
			_ = _321_v34
			_321_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_320_v33))
			var _322_v35 _dafny.Array
			_ = _322_v35
			var _nwElement0_12 _dafny.Map = (func() _dafny.Map {
				if (_321_v34).Contains(p2) {
					return (_321_v34).Get(p2).(_dafny.Map)
				}
				return _319_v32
			})()
			_ = _nwElement0_12
			var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(8))
			_ = _nw56
			(_nw56).ArraySet1(_nwElement0_12, 0)
			(_nw56).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6(true, (_313_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_313_v26), 0))).Int()).(_dafny.Int), globalState), _311_v24), 1)
			(_nw56).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_318_v31, _311_v24), 2)
			(_nw56).ArraySet1(_319_v32, 3)
			(_nw56).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_318_v31, _311_v24), 4)
			(_nw56).ArraySet1(_319_v32, 5)
			(_nw56).ArraySet1(_319_v32, 6)
			(_nw56).ArraySet1(_319_v32, 7)
			_322_v35 = _nw56
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_322_v35), 0))
			_ = _index53
			(_322_v35).ArraySet1(_319_v32, (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_322_v35), 0))
			_ = _index54
			(_322_v35).ArraySet1((_320_v33), (_index54).Int())
		} else {
			(globalState).F10 = p2
			var _323_v36 D2
			_ = _323_v36
			_323_v36 = Companion_D2_.Create_DC3_(p2, p2)
			var _source5 D2 = _323_v36
			_ = _source5
			if _source5.Is_DC3() {
				var _324___mcc_h0 bool = _source5.Get_().(D2_DC3).Cf3
				_ = _324___mcc_h0
				var _325___mcc_h1 bool = _source5.Get_().(D2_DC3).Cf4
				_ = _325___mcc_h1
				var _326_cf4 bool = _325___mcc_h1
				_ = _326_cf4
				var _327_cf3 bool = _324___mcc_h0
				_ = _327_cf3
				var _328_v37 _dafny.CodePoint
				_ = _328_v37
				_328_v37 = _dafny.CodePoint('b')
				var _329_v38 _dafny.Array
				_ = _329_v38
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_8
				var _nw57 _dafny.Array
				_ = _nw57
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw57 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) bool = (func(_330_p2 bool) func(_dafny.Int) bool {
						return func(_331_i6 _dafny.Int) bool {
							return _330_p2
						}
					})(p2)
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
				_329_v38 = _nw57
				var _332_v39 _dafny.Map
				_ = _332_v39
				_332_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v37, _329_v38)
				_332_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
					if p2 {
						return Companion_Default___.Fm5(_dafny.IntOfUint32((_276_v0).Cardinality()), globalState)
					}
					return _328_v37
				})(), _329_v38)
				var _333_v40 _dafny.Sequence
				_ = _333_v40
				_333_v40 = _dafny.UnicodeSeqOfUtf8Bytes("ovmhyj")
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_329_v38), 0))
				_ = _index55
				(_329_v38).ArraySet1(_327_cf3, (_index55).Int())
				var _334_v41 _dafny.MultiSet
				_ = _334_v41
				_334_v41 = _dafny.MultiSetOf(p2)
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_329_v38), 0))
				_ = _index56
				var _rhs50 bool = _326_cf4
				_ = _rhs50
				var _rhs51 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p1, _dafny.Companion_Sequence_.Concatenate(_333_v40, p1))
				_ = _rhs51
				var _rhs52 bool = ((_334_v41).Contains(false)) == (p2)
				_ = _rhs52
				var _lhs45 *GlobalState = globalState
				_ = _lhs45
				var _lhs46 _dafny.Array = _329_v38
				_ = _lhs46
				var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(415), _dafny.ArrayLen((_329_v38), 0))
				_ = _lhs47
				_lhs45.F10 = _rhs50
				_333_v40 = _rhs51
				(_lhs46).ArraySet1(_rhs52, (_lhs47).Int())
				var _335_v42 _dafny.Array
				_ = _335_v42
				var _nwElement0_13 _dafny.Int = Companion_Default___.Fm2(!(p2), (_dafny.Zero).Minus((_dafny.SetOf(p0, p0)).Cardinality()), p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-331))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_336_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_337_i7 _dafny.Int) _dafny.CodePoint {
						return _336_v37
					}
				})(_328_v37))), globalState)
				_ = _nwElement0_13
				var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(14))
				_ = _nw58
				(_nw58).ArraySet1(_nwElement0_13, 0)
				(_nw58).ArraySet1(_dafny.IntOfInt64(-31), 1)
				(_nw58).ArraySet1(p3, 2)
				(_nw58).ArraySet1(p3, 3)
				(_nw58).ArraySet1(p3, 4)
				(_nw58).ArraySet1(p3, 5)
				(_nw58).ArraySet1(p3, 6)
				(_nw58).ArraySet1(p0, 7)
				(_nw58).ArraySet1(p3, 8)
				(_nw58).ArraySet1(p0, 9)
				(_nw58).ArraySet1(p0, 10)
				(_nw58).ArraySet1((_dafny.Zero).Minus(p0), 11)
				(_nw58).ArraySet1(p3, 12)
				(_nw58).ArraySet1(p3, 13)
				_335_v42 = _nw58
				var _338_v43 _dafny.Map
				_ = _338_v43
				_338_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _335_v42)
				var _339_v44 _dafny.Sequence
				_ = _339_v44
				_339_v44 = _dafny.SeqOf(_327_cf3)
				var _340_v45 _dafny.Map
				_ = _340_v45
				_340_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v37, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))
				var _341_v46 _dafny.MultiSet
				_ = _341_v46
				_341_v46 = _dafny.MultiSetOf(_dafny.IntOfInt64(43))
				var _342_v47 _dafny.Sequence
				_ = _342_v47
				_342_v47 = _dafny.SeqOf(false, Companion_Default___.Fm0(Companion_Default___.Fm0(_327_cf3, (_339_v44).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_339_v44).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(808), p0, globalState), p2, (func() _dafny.Int {
					if (_340_v45).Contains(_dafny.CodePoint('r')) {
						return (_340_v45).Get(_dafny.CodePoint('r')).(_dafny.Int)
					}
					return _dafny.IntOfInt64(550)
				})(), (_341_v46).Cardinality(), globalState), _326_cf4, _327_cf3, p2)
				_338_v43 = (_338_v43).Update((_339_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-217), _dafny.IntOfUint32((_339_v44).Cardinality()))).Uint32()).(bool), _335_v42)
				(globalState).F9 = _327_cf3
			} else if _source5.Is_DC4() {
				var _343___mcc_h2 _dafny.Int = _source5.Get_().(D2_DC4).Cf5
				_ = _343___mcc_h2
				var _344___mcc_h3 _dafny.CodePoint = _source5.Get_().(D2_DC4).Cf6
				_ = _344___mcc_h3
				var _345_cf6 _dafny.CodePoint = _344___mcc_h3
				_ = _345_cf6
				var _346_cf5 _dafny.Int = _343___mcc_h2
				_ = _346_cf5
				var _347_v48 D2
				_ = _347_v48
				_347_v48 = Companion_D2_.Create_DC4_(_346_cf5, _345_cf6)
				var _348_v49 _dafny.Map
				_ = _348_v49
				_348_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let2_0 D2) D2 {
					return func(_349_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let3_0 _dafny.CodePoint) D2 {
							return func(_350_dt__update_hcf6_h0 _dafny.CodePoint) D2 {
								return Companion_D2_.Create_DC4_((_349_dt__update__tmp_h0).Dtor_cf5(), _350_dt__update_hcf6_h0)
							}(_pat_let3_0)
						}(_dafny.CodePoint('d'))
					}(_pat_let2_0)
				}(_347_v48), true)
				_348_v49 = (_348_v49).Update(_347_v48, _dafny.Companion_Sequence_.IsProperPrefixOf(p1, p1))
				var _351_v50 *C0
				_ = _351_v50
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(_345_cf6)
				_351_v50 = _nw59
				var _352_v51 _dafny.MultiSet
				_ = _352_v51
				_352_v51 = _dafny.MultiSetOf(_351_v50, _351_v50, _351_v50, _351_v50)
				_352_v51 = (_352_v51).Update(_351_v50, Companion_Default___.Abs(_346_cf5))
				_345_cf6 = (_351_v50).F12()
				var _353_v52 _dafny.Set
				_ = _353_v52
				_353_v52 = _dafny.SetOf(!(p2), p2, p2, Companion_Default___.Fm0(true, p2, _dafny.IntOfInt64(585), _346_cf5, globalState), true)
				var _354_v53 _dafny.MultiSet
				_ = _354_v53
				_354_v53 = _dafny.MultiSetOf(p3)
				var _355_v54 _dafny.Map
				_ = _355_v54
				_355_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_351_v50, p3)
				var _356_v55 _dafny.Map
				_ = _356_v55
				_356_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(p2), _351_v50)
				var _357_v56 _dafny.Map
				_ = _357_v56
				_357_v56 = _356_v55
				var _358_v57 _dafny.Map
				_ = _358_v57
				_358_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_323_v36).Dtor_cf3(), _357_v56)
				var _359_v58 _dafny.Sequence
				_ = _359_v58
				_359_v58 = _dafny.SeqOf(p2)
				var _360_v59 _dafny.Array
				_ = _360_v59
				var _nwElement0_14 _dafny.Int = _dafny.IntOfUint32((_276_v0).Cardinality())
				_ = _nwElement0_14
				var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(17))
				_ = _nw60
				(_nw60).ArraySet1(_nwElement0_14, 0)
				(_nw60).ArraySet1(p0, 1)
				(_nw60).ArraySet1(p3, 2)
				(_nw60).ArraySet1(p0, 3)
				(_nw60).ArraySet1((_353_v52).Cardinality(), 4)
				(_nw60).ArraySet1(_dafny.IntOfInt64(608), 5)
				(_nw60).ArraySet1((_354_v53).Cardinality(), 6)
				(_nw60).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 7)
				(_nw60).ArraySet1((func() _dafny.Int {
					if (_354_v53).Contains(_346_cf5) {
						return (_354_v53).Multiplicity(_346_cf5)
					}
					return (_dafny.Zero).Minus(Companion_Default___.Fm2(true, p0, p0, p1, globalState))
				})(), 8)
				(_nw60).ArraySet1(_346_cf5, 9)
				(_nw60).ArraySet1((_355_v54).Cardinality(), 10)
				(_nw60).ArraySet1(((_358_v57).Update(p2, _357_v56)).Cardinality(), 11)
				(_nw60).ArraySet1(_dafny.IntOfUint32((_359_v58).Cardinality()), 12)
				(_nw60).ArraySet1(p0, 13)
				(_nw60).ArraySet1(p3, 14)
				(_nw60).ArraySet1(_346_cf5, 15)
				(_nw60).ArraySet1(Companion_Default___.Fm2(p2, p0, _346_cf5, p1, globalState), 16)
				_360_v59 = _nw60
				var _361_v60 _dafny.Array
				_ = _361_v60
				var _nwElement0_15 _dafny.Int = _dafny.IntOfUint32((p1).Cardinality())
				_ = _nwElement0_15
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(7))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_15, 0)
				(_nw61).ArraySet1(_346_cf5, 1)
				(_nw61).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_351_v50), _360_v59)).Cardinality(), 2)
				(_nw61).ArraySet1(p3, 3)
				(_nw61).ArraySet1(_dafny.IntOfInt64(671), 4)
				(_nw61).ArraySet1(_346_cf5, 5)
				(_nw61).ArraySet1(p3, 6)
				_361_v60 = _nw61
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_361_v60), 0))
				_ = _index57
				(_361_v60).ArraySet1((_dafny.IntOfInt64(-38)).Plus(p3), (_index57).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(655), _dafny.ArrayLen((_361_v60), 0))
				_ = _index58
				(_361_v60).ArraySet1(p0, (_index58).Int())
			} else {
				var _362___mcc_h4 _dafny.CodePoint = _source5.Get_().(D2_DC2).Cf2
				_ = _362___mcc_h4
				var _363_cf2 _dafny.CodePoint = _362___mcc_h4
				_ = _363_cf2
				var _364_v61 _dafny.Set
				_ = _364_v61
				_364_v61 = _dafny.SetOf(p3)
				var _365_v62 _dafny.Map
				_ = _365_v62
				_365_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), p3)
				var _366_v63 _dafny.Map
				_ = _366_v63
				_366_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
				var _367_v64 *C0
				_ = _367_v64
				var _nw62 *C0 = New_C0_()
				_ = _nw62
				_nw62.Ctor__(_363_cf2)
				_367_v64 = _nw62
				var _368_v65 _dafny.Map
				_ = _368_v65
				_368_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v64, p2)
				var _369_v66 _dafny.MultiSet
				_ = _369_v66
				_369_v66 = _dafny.MultiSetOf(p0, p3)
				var _370_v67 _dafny.MultiSet
				_ = _370_v67
				_370_v67 = _dafny.MultiSetOf(p2, p2)
				var _371_v68 _dafny.Array
				_ = _371_v68
				var _nwElement0_16 _dafny.Int = Companion_Default___.Fm2(p2, p3, _dafny.IntOfUint32((p1).Cardinality()), p1, globalState)
				_ = _nwElement0_16
				var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(21))
				_ = _nw63
				(_nw63).ArraySet1(_nwElement0_16, 0)
				(_nw63).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm2(p2, p3, p0, p1, globalState))), 1)
				(_nw63).ArraySet1((func() _dafny.Int {
					if p2 {
						return (_dafny.Zero).Minus(p0)
					}
					return p3
				})(), 2)
				(_nw63).ArraySet1(p0, 3)
				(_nw63).ArraySet1((func() _dafny.Int {
					if !(p2) {
						return (_364_v61).Cardinality()
					}
					return p0
				})(), 4)
				(_nw63).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(807), p0), 5)
				(_nw63).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((p1).Cardinality()), p3), 6)
				(_nw63).ArraySet1(p3, 7)
				(_nw63).ArraySet1(p3, 8)
				(_nw63).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_372_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_373_i8 _dafny.Int) _dafny.Int {
						return _372_p0
					}
				})(p0))))).Cardinality(), 9)
				(_nw63).ArraySet1(p3, 10)
				(_nw63).ArraySet1(p3, 11)
				(_nw63).ArraySet1((func() _dafny.Int {
					if p2 {
						return p0
					}
					return (func() _dafny.Int {
						if (_365_v62).Contains((_this).F11()) {
							return (_365_v62).Get((_this).F11()).(_dafny.Int)
						}
						return p3
					})()
				})(), 12)
				(_nw63).ArraySet1(p0, 13)
				(_nw63).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 14)
				(_nw63).ArraySet1(_dafny.IntOfInt64(-629), 15)
				(_nw63).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-111), (_366_v63).Cardinality())), 16)
				(_nw63).ArraySet1((_368_v65).Cardinality(), 17)
				(_nw63).ArraySet1((_dafny.Zero).Minus((_369_v66).Cardinality()), 18)
				(_nw63).ArraySet1(p3, 19)
				(_nw63).ArraySet1((func() _dafny.Int {
					if p2 {
						return p3
					}
					return _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf((func() _dafny.Int {
						if (_370_v67).Contains(p2) {
							return (_370_v67).Multiplicity(p2)
						}
						return p0
					})())).Cardinality())).Cardinality())
				})(), 20)
				_371_v68 = _nw63
				var _374_v69 _dafny.Array
				_ = _374_v69
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_9
				var _nw64 _dafny.Array
				_ = _nw64
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw64 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.CodePoint = (func(_375_v64 *C0) func(_dafny.Int) _dafny.CodePoint {
						return func(_376_i9 _dafny.Int) _dafny.CodePoint {
							return (_375_v64).F12()
						}
					})(_367_v64)
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
				_374_v69 = _nw64
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_374_v69), 0))
				_ = _index59
				(_374_v69).ArraySet1CodePoint(_363_cf2, (_index59).Int())
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_374_v69), 0))
				_ = _index60
				var _rhs53 _dafny.Array = _371_v68
				_ = _rhs53
				var _rhs54 _dafny.CodePoint = _363_cf2
				_ = _rhs54
				var _rhs55 bool = p2
				_ = _rhs55
				var _rhs56 bool = p2
				_ = _rhs56
				var _lhs48 _dafny.Array = _374_v69
				_ = _lhs48
				var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_374_v69), 0))
				_ = _lhs49
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				_371_v68 = _rhs53
				(_lhs48).ArraySet1CodePoint(_rhs54, (_lhs49).Int())
				_lhs50.F10 = _rhs55
				_lhs51.F9 = _rhs56
				var _nw65 *C0 = New_C0_()
				_ = _nw65
				_nw65.Ctor__((_367_v64).F12())
				_367_v64 = _nw65
				var _377_v70 _dafny.Map
				_ = _377_v70
				_377_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm0(p2, p2, p3, p3, globalState))
				_377_v70 = func() _dafny.Map {
					var _coll13 = _dafny.NewMapBuilder()
					_ = _coll13
					for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(531), _dafny.IntOfInt64(273))); ; {
						_compr_13, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _378_v71 _dafny.Int
						_378_v71 = interface{}(_compr_13).(_dafny.Int)
						if ((_dafny.IntOfInt64(531)).Cmp(_378_v71) <= 0) && ((_378_v71).Cmp(_dafny.IntOfInt64(273)) < 0) {
							_coll13.Add((_378_v71).Minus(_dafny.IntOfUint32((_276_v0).Cardinality())), !(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(p2), _dafny.SeqOf(p2, p2))))
						}
					}
					return _coll13.ToMap()
				}()
				var _379_v72 *C0
				_ = _379_v72
				var _nw66 *C0 = New_C0_()
				_ = _nw66
				_nw66.Ctor__((_367_v64).F12())
				_379_v72 = _nw66
			}
			var _380_v73 _dafny.Map
			_ = _380_v73
			_380_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.UnicodeSeqOfUtf8Bytes("nqmai"))
			var _381_v74 _dafny.Map
			_ = _381_v74
			_381_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _382_v75 _dafny.Set
			_ = _382_v75
			_382_v75 = _dafny.SetOf((_381_v74).Cardinality())
			var _383_v76 _dafny.CodePoint
			_ = _383_v76
			_383_v76 = _dafny.CodePoint('o')
			var _384_v77 _dafny.MultiSet
			_ = _384_v77
			_384_v77 = _dafny.MultiSetOf(p2)
			var _385_v78 _dafny.Sequence
			_ = _385_v78
			_385_v78 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("twnjp"), p1)
			var _386_v79 _dafny.Array
			_ = _386_v79
			var _nwElement0_17 _dafny.Int = (_dafny.SetOf(p3)).Cardinality()
			_ = _nwElement0_17
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(28))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_17, 0)
			(_nw67).ArraySet1(p3, 1)
			(_nw67).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p3, p3)), 2)
			(_nw67).ArraySet1(p0, 3)
			(_nw67).ArraySet1((func() _dafny.Int {
				if p2 {
					return p0
				}
				return (_dafny.Zero).Minus(p3)
			})(), 4)
			(_nw67).ArraySet1(((Companion_Default___.Fm7(p2, p0, p2, (_380_v73).Cardinality(), globalState)).Cardinality()).Minus((_dafny.Zero).Minus(p3)), 5)
			(_nw67).ArraySet1(p0, 6)
			(_nw67).ArraySet1(p3, 7)
			(_nw67).ArraySet1(p3, 8)
			(_nw67).ArraySet1(p3, 9)
			(_nw67).ArraySet1((_382_v75).Cardinality(), 10)
			(_nw67).ArraySet1(p0, 11)
			(_nw67).ArraySet1(Companion_Default___.SafeDivisionInt(p0, (_381_v74).Cardinality()), 12)
			(_nw67).ArraySet1((_dafny.Zero).Minus(p0), 13)
			(_nw67).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("m"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()))).Uint32(), _383_v76)).Cardinality()), 14)
			(_nw67).ArraySet1((_276_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.IntOfUint32((_276_v0).Cardinality()))).Uint32()).(_dafny.Int), 15)
			(_nw67).ArraySet1(p3, 16)
			(_nw67).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm8(p2, p1, globalState)).Cardinality()), 17)
			(_nw67).ArraySet1((_dafny.Zero).Minus(p3), 18)
			(_nw67).ArraySet1(p3, 19)
			(_nw67).ArraySet1((_dafny.IntOfInt64(-652)).Minus(p0), 20)
			(_nw67).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xlpfc")).Cardinality()), 21)
			(_nw67).ArraySet1(_dafny.IntOfInt64(739), 22)
			(_nw67).ArraySet1((((_dafny.MultiSetOf(p2)).Update(p2, Companion_Default___.Abs(p3))).Union(_384_v77)).Cardinality(), 23)
			(_nw67).ArraySet1(p3, 24)
			(_nw67).ArraySet1(p3, 25)
			(_nw67).ArraySet1(_dafny.IntOfInt64(655), 26)
			(_nw67).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, (_385_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_385_v78).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()), 27)
			_386_v79 = _nw67
			var _387_v80 _dafny.MultiSet
			_ = _387_v80
			_387_v80 = _dafny.MultiSetOf(_dafny.IntOfUint32((p1).Cardinality()), p3)
			var _388_v81 *C0
			_ = _388_v81
			var _nw68 *C0 = New_C0_()
			_ = _nw68
			_nw68.Ctor__(_dafny.CodePoint('j'))
			_388_v81 = _nw68
			var _389_v82 _dafny.Map
			_ = _389_v82
			_389_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v77, _388_v81)
			var _390_v83 _dafny.Map
			_ = _390_v83
			_390_v83 = _389_v82
			var _391_v84 _dafny.Set
			_ = _391_v84
			_391_v84 = _dafny.SetOf(_390_v83)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_386_v79), 0))
			_ = _index61
			(_386_v79).ArraySet1((func() _dafny.Int {
				if (_387_v80).Contains(p0) {
					return (_387_v80).Multiplicity(p0)
				}
				return (_dafny.Zero).Minus((_391_v84).Cardinality())
			})(), (_index61).Int())
			var _392_v85 _dafny.Map
			_ = _392_v85
			_392_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _393_v86 _dafny.Map
			_ = _393_v86
			_393_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (func() _dafny.Int {
				if (_392_v85).Contains(p0) {
					return (_392_v85).Get(p0).(_dafny.Int)
				}
				return p3
			})())
			var _394_v87 _dafny.Map
			_ = _394_v87
			_394_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.SetOf(Companion_D2_.Create_DC3_(p2, p2), _323_v36, _323_v36, _323_v36))
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_386_v79), 0))
			_ = _index62
			var _rhs57 _dafny.Int = (Companion_Default___.Fm1(_392_v85, !(p2), _dafny.IntOfUint32((p1).Cardinality()), (_dafny.Zero).Minus(p0), globalState)).Cardinality()
			_ = _rhs57
			var _rhs58 _dafny.CodePoint = (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs58
			var _rhs59 bool = !(Companion_Default___.Fm9(p2, globalState)).Equals(_394_v87)
			_ = _rhs59
			var _lhs52 _dafny.Array = _386_v79
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_386_v79), 0))
			_ = _lhs53
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			(_lhs52).ArraySet1(_rhs57, (_lhs53).Int())
			_383_v76 = _rhs58
			_lhs54.F10 = _rhs59
			(globalState).F9 = p2
			var _395_v88 _dafny.Map
			_ = _395_v88
			_395_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			_395_v88 = (_395_v88).Update(p2, (_dafny.IntOfInt64(-902)).Plus((_395_v88).Cardinality()))
		}
		var _396_v89 _dafny.MultiSet
		_ = _396_v89
		_396_v89 = _dafny.MultiSetOf(p2)
		var _397_v90 _dafny.CodePoint
		_ = _397_v90
		_397_v90 = _dafny.CodePoint('y')
		var _398_v91 *C0
		_ = _398_v91
		var _nw69 *C0 = New_C0_()
		_ = _nw69
		_nw69.Ctor__(_397_v90)
		_398_v91 = _nw69
		var _399_v92 _dafny.Map
		_ = _399_v92
		_399_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_396_v89, _398_v91)
		var _400_v93 _dafny.Map
		_ = _400_v93
		_400_v93 = _399_v92
		var _401_v94 _dafny.Sequence
		_ = _401_v94
		_401_v94 = _dafny.SeqOf(_400_v93)
		var _hi2 _dafny.Int = Companion_Default___.Fm2(p2, (_dafny.Zero).Minus(_dafny.IntOfInt64(-285)), p0, p1, globalState)
		_ = _hi2
		for _402_i10 := Companion_Default___.SafeDivisionInt((_dafny.MultiSetFromSeq(_401_v94)).Cardinality(), p0); _402_i10.Cmp(_hi2) < 0; _402_i10 = _402_i10.Plus(_dafny.One) {
			var _403_v95 _dafny.Map
			_ = _403_v95
			_403_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			var _404_v96 _dafny.Map
			_ = _404_v96
			_404_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_403_v95).Cardinality())
			var _405_v97 _dafny.Map
			_ = _405_v97
			_405_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_404_v96).Contains(p0) {
					return (_404_v96).Get(p0).(_dafny.Int)
				}
				return p3
			})(), p2)
			var _406_v98 _dafny.Sequence
			_ = _406_v98
			_406_v98 = _dafny.SeqOf((p2) == (!((func() bool {
				if (_405_v97).Contains(p3) {
					return (_405_v97).Get(p3).(bool)
				}
				return p2
			})())))
			(globalState).F10 = (_406_v98).Select((Companion_Default___.SafeIndex(_402_i10, _dafny.IntOfUint32((_406_v98).Cardinality()))).Uint32()).(bool)
			var _407_v99 _dafny.Array
			_ = _407_v99
			var _nwElement0_18 _dafny.Map = _400_v93
			_ = _nwElement0_18
			var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(21))
			_ = _nw70
			(_nw70).ArraySet1(_nwElement0_18, 0)
			(_nw70).ArraySet1((func() _dafny.Map {
				if p2 {
					return _400_v93
				}
				return _400_v93
			})(), 1)
			(_nw70).ArraySet1(_400_v93, 2)
			(_nw70).ArraySet1(_400_v93, 3)
			(_nw70).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_396_v89, _398_v91), 4)
			(_nw70).ArraySet1(_400_v93, 5)
			(_nw70).ArraySet1(_400_v93, 6)
			(_nw70).ArraySet1(_400_v93, 7)
			(_nw70).ArraySet1(_400_v93, 8)
			(_nw70).ArraySet1(_400_v93, 9)
			(_nw70).ArraySet1(_400_v93, 10)
			(_nw70).ArraySet1(_400_v93, 11)
			(_nw70).ArraySet1(_400_v93, 12)
			(_nw70).ArraySet1(_399_v92, 13)
			(_nw70).ArraySet1(_399_v92, 14)
			(_nw70).ArraySet1(_400_v93, 15)
			(_nw70).ArraySet1(_400_v93, 16)
			(_nw70).ArraySet1(_399_v92, 17)
			(_nw70).ArraySet1(_399_v92, 18)
			(_nw70).ArraySet1(_399_v92, 19)
			(_nw70).ArraySet1((_399_v92), 20)
			_407_v99 = _nw70
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_407_v99), 0))
			_ = _index63
			(_407_v99).ArraySet1(_400_v93, (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_407_v99), 0))
			_ = _index64
			(_407_v99).ArraySet1(_399_v92, (_index64).Int())
			var _408_v100 _dafny.Map
			_ = _408_v100
			_408_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _402_i10)
			var _409_v101 _dafny.Array
			_ = _409_v101
			var _nw71 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw71
			_409_v101 = _nw71
			var _410_v102 _dafny.Map
			_ = _410_v102
			_410_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v100, _409_v101)
			_410_v102 = (_410_v102).Update(_408_v100, _409_v101)
			var _411_v103 *C0
			_ = _411_v103
			var _nw72 *C0 = New_C0_()
			_ = _nw72
			_nw72.Ctor__(_397_v90)
			_411_v103 = _nw72
		}
		var _412_v104 _dafny.Sequence
		_ = _412_v104
		_412_v104 = _dafny.SeqOf(p2)
		var _413_v105 _dafny.Array
		_ = _413_v105
		var _nwElement0_19 bool = p2
		_ = _nwElement0_19
		var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(26))
		_ = _nw73
		(_nw73).ArraySet1(_nwElement0_19, 0)
		(_nw73).ArraySet1(p2, 1)
		(_nw73).ArraySet1((_this).F11(), 2)
		(_nw73).ArraySet1((_this).F11(), 3)
		(_nw73).ArraySet1((_this).F11(), 4)
		(_nw73).ArraySet1((_this).F11(), 5)
		(_nw73).ArraySet1((_this).F11(), 6)
		(_nw73).ArraySet1((_this).F11(), 7)
		(_nw73).ArraySet1((_this).F11(), 8)
		(_nw73).ArraySet1(true, 9)
		(_nw73).ArraySet1((_this).F11(), 10)
		(_nw73).ArraySet1((_this).F11(), 11)
		(_nw73).ArraySet1((_this).F11(), 12)
		(_nw73).ArraySet1(p2, 13)
		(_nw73).ArraySet1((_this).F11(), 14)
		(_nw73).ArraySet1((_this).F11(), 15)
		(_nw73).ArraySet1((_this).F11(), 16)
		(_nw73).ArraySet1((_this).F11(), 17)
		(_nw73).ArraySet1((_this).F11(), 18)
		(_nw73).ArraySet1((_this).F11(), 19)
		(_nw73).ArraySet1((_412_v104).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_412_v104).Cardinality()))).Uint32()).(bool), 20)
		(_nw73).ArraySet1((_this).F11(), 21)
		(_nw73).ArraySet1((_this).F11(), 22)
		(_nw73).ArraySet1((_this).F11(), 23)
		(_nw73).ArraySet1(p2, 24)
		(_nw73).ArraySet1((_this).F11(), 25)
		_413_v105 = _nw73
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_413_v105), 0))
		_ = _index65
		(_413_v105).ArraySet1((_this).F11(), (_index65).Int())
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_413_v105), 0))
		_ = _index66
		(_413_v105).ArraySet1((_this).F11(), (_index66).Int())
		r0 = Companion_Default___.Fm10(globalState)
		return r0
	}
}
func (_this *C1) F11() bool {
	{
		return _this._f11
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
