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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.CodePoint, globalState *GlobalState) bool {
	return (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(763), _dafny.IntOfInt64(294))).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(934))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})), _dafny.UnicodeSeqOfUtf8Bytes("goxbqw"))).Cardinality())) > 0
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	if (true) == (true) {
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(882))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), _dafny.UnicodeSeqOfUtf8Bytes("ov"))).Cardinality())
	} else {
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('f'), _dafny.CodePoint('v'), _dafny.CodePoint('d'), _dafny.CodePoint('y')), _dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('p'), _dafny.CodePoint('w'), _dafny.CodePoint('p')))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(710), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(816), _dafny.CodePoint('l'))).Cardinality())).Cardinality(), true)).Cardinality())).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("irm")).Cardinality())) < 0, (Companion_D9_.Create_DC25_(_dafny.UnicodeSeqOfUtf8Bytes("ap"), false, _dafny.IntOfInt64(214))).Dtor_cf60(), false)
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	if !(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("sfmw"))).Cardinality()).Cmp(_dafny.IntOfInt64(306)) >= 0) {
		return _dafny.SeqOf(true, false)
	} else {
		return _dafny.SeqOf(true, !(!(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("x")
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if (true) || (false) {
		return _dafny.SeqOf(false)
	} else {
		return _dafny.SeqOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-834), _dafny.IntOfInt64(-776))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-834)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(-776)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_2_v0, (_dafny.MultiSetOf(true)).Cardinality()), _dafny.IntOfInt64(54))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), false)).Cardinality(), (_dafny.SetOf(!((Companion_D8_.Create_DC20_(true, false, false)).Dtor_cf48()))).Cardinality())).Intersection(_dafny.MultiSetOf((_dafny.SetOf(false, true)).Cardinality(), _dafny.Zero, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC18_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(552))).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("dvkfi"), true, !(false), func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(-683))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v1 _dafny.Int
			_3_v1 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.SetOf(_dafny.IntOfInt64(-683))).Contains(_3_v1) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_3_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mqnecwgoe")).Cardinality())))
			}
		}
		return _coll1.ToSet()
	}())).Dtor_cf43(), _dafny.CodePoint('r'))).Cardinality())).Cardinality())), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(951), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-863), true)).Cardinality(), _dafny.IntOfInt64(159), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())))).Union((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-677)))).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(861))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("od")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.One)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 D1, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('n')
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((Companion_D9_.Create_DC25_(_dafny.UnicodeSeqOfUtf8Bytes("klphfov"), !(true), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("arkhat")).Cardinality()), false)).Cardinality()))).Dtor_cf60(), true, false, (func() _dafny.CodePoint {
		if true {
			return _dafny.CodePoint('k')
		}
		return _dafny.CodePoint('g')
	})(), _dafny.CodePoint('k'))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), true))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_4_i0 _dafny.Int) _dafny.Set {
		return _dafny.SetOf(Companion_D1_.Create_DC3_(Companion_D1_.Create_DC1_(_dafny.SetOf(true))))
	})))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, false, !(true), true, false)).Union(_dafny.SetOf(true, !(false)))).Union(_dafny.SetOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((Companion_D10_.Create_DC27_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(143)))).Dtor_cf63(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(176))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(694))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(514), _dafny.IntOfInt64(-236))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _5_v0 _dafny.Int
			_5_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(514)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(-236)) < 0) {
				_coll2.Add((_5_v0).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-183))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}(func(_6_i0 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (func() _dafny.Map {
						var _coll3 = _dafny.NewMapBuilder()
						_ = _coll3
						for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(740), _dafny.IntOfInt64(303))); ; {
							_compr_3, _ok3 := _iter3()
							if !_ok3 {
								break
							}
							var _7_v1 _dafny.Int
							_7_v1 = interface{}(_compr_3).(_dafny.Int)
							if ((_dafny.IntOfInt64(740)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(303)) < 0) {
								_coll3.Add((_7_v1).Plus((_dafny.MultiSetOf(_dafny.IntOfInt64(278))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wag")).Cardinality()))
							}
						}
						return _coll3.ToMap()
					}()).Cardinality())).Cardinality()
				}))).Cardinality())))
			}
		}
		return _coll2.ToSet()
	}()).Difference((func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-46), _dafny.IntOfInt64(53))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _8_v2 _dafny.Int
			_8_v2 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-46)).Cmp(_8_v2) <= 0) && ((_8_v2).Cmp(_dafny.IntOfInt64(53)) < 0) {
				_coll4.Add((_8_v2).Times(_dafny.IntOfInt64(370)))
			}
		}
		return _coll4.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfInt64(-604), (_dafny.Zero).Minus((_dafny.SetOf((Companion_D1_.Create_DC2_(!(true), true, false, _dafny.CodePoint('g'), _dafny.CodePoint('i'))).Dtor_cf2())).Cardinality()), _dafny.IntOfInt64(806))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (Companion_D15_.Create_DC36_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(123)))).Dtor_cf78()
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.IntOfInt64(492)).Cmp(_dafny.IntOfInt64(754)) <= 0)
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oot")).Cardinality()), false)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(120), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality(), _dafny.IntOfInt64(712))), _dafny.IntOfInt64(-813), (func() _dafny.Int {
		if !(true) {
			return _dafny.IntOfInt64(941)
		}
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ldpqylslw")).Cardinality())
	})())
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.CodePoint, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("aua"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _10_v0 _dafny.Sequence
			_10_v0 = interface{}(_compr_5).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("aua"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			})))).Contains(_10_v0) {
				_coll5.Add(_10_v0, _dafny.UnicodeSeqOfUtf8Bytes("jhdbecap"))
			}
		}
		return _coll5.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-65))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_11_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_12_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("iqvwwh"), _dafny.UnicodeSeqOfUtf8Bytes("oudm")))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC20_((func() bool {
		if false {
			return !(true)
		}
		return true
	})(), true, _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("qoxnse"), _dafny.UnicodeSeqOfUtf8Bytes("rjxfs")))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.MultiSet, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(396))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(410), _dafny.IntOfInt64(166))))
}
func (_static *CompanionStruct_Default___) Fm28(globalState *GlobalState) bool {
	if false {
		return true
	} else {
		return true
	}
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-549), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-220)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-451), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(540))))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) {
	var _13_v0 bool
	_ = _13_v0
	_13_v0 = false
	(globalState).F3 = _13_v0
	var _14_v1 _dafny.CodePoint
	_ = _14_v1
	_14_v1 = _dafny.CodePoint('f')
	var _15_v2 _dafny.Map
	_ = _15_v2
	_15_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _14_v1)
	var _16_v3 _dafny.Int
	_ = _16_v3
	_16_v3 = _dafny.IntOfInt64(692)
	_15_v2 = (_15_v2).Update((_16_v3).Cmp(_dafny.IntOfInt64(390)) >= 0, _14_v1)
	var _17_v4 _dafny.Sequence
	_ = _17_v4
	_17_v4 = _dafny.SeqOf(_16_v3, _dafny.IntOfInt64(164))
	var _18_i0 _dafny.Int
	_ = _18_i0
	_18_i0 = _dafny.Zero
	{
		for (_16_v3).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_17_v4).Cardinality()))).Cardinality()), _16_v3)) <= 0 {
			{
				if (_18_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_18_i0 = (_18_i0).Plus(_dafny.One)
				(globalState).F9 = _16_v3
				if !(!(!(!(_13_v0)))) {
					var _19_v5 _dafny.Array
					_ = _19_v5
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
					_ = _len0_0
					var _nw0 _dafny.Array
					_ = _nw0
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw0 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) bool = (func(_20_v3 _dafny.Int) func(_dafny.Int) bool {
							return func(_21_i1 _dafny.Int) bool {
								return (_20_v3).Cmp(_20_v3) != 0
							}
						})(_16_v3)
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
					_19_v5 = _nw0
					var _22_v6 _dafny.Sequence
					_ = _22_v6
					_22_v6 = _dafny.UnicodeSeqOfUtf8Bytes("luog")
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_19_v5), 0))
					_ = _index0
					(_19_v5).ArraySet1(!_dafny.Companion_Sequence_.Equal(_22_v6, _22_v6), (_index0).Int())
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_19_v5), 0))
					_ = _index1
					(_19_v5).ArraySet1(Companion_Default___.Fm0(_14_v1, globalState), (_index1).Int())
					var _23_v7 _dafny.Array
					_ = _23_v7
					var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
					_ = _nw1
					_23_v7 = _nw1
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_19_v5), 0))
					_ = _index2
					var _rhs0 bool = (_dafny.SetOf(_23_v7)).IsProperSubsetOf(_dafny.SetOf(_23_v7))
					_ = _rhs0
					var _rhs1 bool = (func() bool {
						if (_19_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_19_v5), 0))).Int()).(bool) {
							return true
						}
						return (_19_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_19_v5), 0))).Int()).(bool)
					})()
					_ = _rhs1
					var _lhs0 *GlobalState = globalState
					_ = _lhs0
					var _lhs1 _dafny.Array = _19_v5
					_ = _lhs1
					var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_19_v5), 0))
					_ = _lhs2
					_lhs0.F3 = _rhs0
					(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
					(globalState).F7 = false
					(globalState).F9 = _16_v3
					(globalState).F9 = _16_v3
				} else {
					var _24_v8 _dafny.MultiSet
					_ = _24_v8
					_24_v8 = _dafny.MultiSetOf(_13_v0, false)
					var _25_v9 _dafny.Sequence
					_ = _25_v9
					_25_v9 = _dafny.SeqOf(_13_v0, _13_v0)
					var _26_v10 _dafny.Set
					_ = _26_v10
					_26_v10 = _dafny.SetOf(_13_v0, _13_v0)
					var _27_v11 _dafny.Map
					_ = _27_v11
					_27_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(130), _16_v3), Companion_Default___.Fm1(_13_v0, (_26_v10).Cardinality(), _16_v3, globalState))
					var _28_v12 _dafny.Set
					_ = _28_v12
					_28_v12 = _dafny.SetOf(_16_v3)
					var _29_v13 _dafny.Sequence
					_ = _29_v13
					_29_v13 = _dafny.UnicodeSeqOfUtf8Bytes("jlgqgkipu")
					_16_v3 = ((func() _dafny.Int {
						if (_24_v8).Contains((_25_v9).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_25_v9).Cardinality()))).Uint32()).(bool)) {
							return (_24_v8).Multiplicity((_25_v9).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_25_v9).Cardinality()))).Uint32()).(bool))
						}
						return _16_v3
					})()).Plus((func() _dafny.Int {
						if (_27_v11).Contains(_28_v12) {
							return (_27_v11).Get(_28_v12).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_29_v13).Cardinality())
					})())
					var _30_v14 *C2
					_ = _30_v14
					var _nw2 *C2 = New_C2_()
					_ = _nw2
					_nw2.Ctor__(_13_v0, _24_v8, Companion_Default___.Fm0((_29_v13).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_29_v13).Cardinality()))).Uint32()).(_dafny.CodePoint), globalState))
					_30_v14 = _nw2
					var _rhs2 _dafny.Int = (_16_v3).Minus((func() _dafny.Int {
						if _30_v14.F15 {
							return _dafny.IntOfUint32((_29_v13).Cardinality())
						}
						return _16_v3
					})())
					_ = _rhs2
					var _rhs3 _dafny.Int = _dafny.IntOfInt64(-264)
					_ = _rhs3
					var _rhs4 _dafny.Int = (_dafny.Zero).Minus(_16_v3)
					_ = _rhs4
					var _rhs5 bool = _13_v0
					_ = _rhs5
					var _lhs3 *GlobalState = globalState
					_ = _lhs3
					_lhs3.F9 = _rhs2
					_16_v3 = _rhs3
					_16_v3 = _rhs4
					_13_v0 = _rhs5
					var _31_v15 _dafny.MultiSet
					_ = _31_v15
					_31_v15 = _dafny.MultiSetOf(_16_v3, _16_v3, _16_v3, _16_v3, _dafny.IntOfUint32((_29_v13).Cardinality()))
					(globalState).F9 = (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
						if (_31_v15).Contains(_16_v3) {
							return (_31_v15).Multiplicity(_16_v3)
						}
						return _dafny.IntOfUint32((_17_v4).Cardinality())
					})())), (_dafny.SetOf(_30_v14.F15)).Cardinality())).Minus(_16_v3)
					(globalState).F7 = _30_v14.F15
				}
				var _32_v16 _dafny.Map
				_ = _32_v16
				_32_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _16_v3)
				var _33_v17 D10
				_ = _33_v17
				_33_v17 = Companion_D10_.Create_DC27_(_32_v16)
				var _pat_let_tv0 = _13_v0
				_ = _pat_let_tv0
				var _pat_let_tv1 = _16_v3
				_ = _pat_let_tv1
				var _pat_let_tv2 = _32_v16
				_ = _pat_let_tv2
				var _source0 D10 = func(_pat_let0_0 D10) D10 {
					return func(_34_dt__update__tmp_h0 D10) D10 {
						return func(_pat_let1_0 _dafny.Map) D10 {
							return func(_35_dt__update_hcf63_h0 _dafny.Map) D10 {
								return Companion_D10_.Create_DC27_(_35_dt__update_hcf63_h0)
							}(_pat_let1_0)
						}((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_pat_let_tv0), _pat_let_tv1)).Merge(_pat_let_tv2))
					}(_pat_let0_0)
				}(_33_v17)
				_ = _source0
				if _source0.Is_DC28() {
					var _36___mcc_h0 *C2 = _source0.Get_().(D10_DC28).Cf64
					_ = _36___mcc_h0
					var _37___mcc_h1 bool = _source0.Get_().(D10_DC28).Cf65
					_ = _37___mcc_h1
					var _38___mcc_h2 bool = _source0.Get_().(D10_DC28).Cf66
					_ = _38___mcc_h2
					var _39___mcc_h3 bool = _source0.Get_().(D10_DC28).Cf67
					_ = _39___mcc_h3
					var _40_cf67 bool = _39___mcc_h3
					_ = _40_cf67
					var _41_cf66 bool = _38___mcc_h2
					_ = _41_cf66
					var _42_cf65 bool = _37___mcc_h1
					_ = _42_cf65
					var _43_cf64 *C2 = _36___mcc_h0
					_ = _43_cf64
					var _44_v18 _dafny.Sequence
					_ = _44_v18
					_44_v18 = _dafny.SeqOf(_41_cf66)
					var _45_v19 _dafny.Map
					_ = _45_v19
					_45_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v1, _44_v18)
					_45_v19 = (_45_v19).Update(_14_v1, _dafny.SeqOf(_40_cf67, _13_v0))
					var _46_v20 _dafny.Set
					_ = _46_v20
					_46_v20 = _dafny.SetOf(_13_v0)
					_40_cf67 = (_46_v20).IsProperSubsetOf((_46_v20).Difference(_46_v20))
					var _47_v21 _dafny.Set
					_ = _47_v21
					_47_v21 = _dafny.SetOf(Companion_Default___.Fm24(globalState))
					var _48_v22 _dafny.MultiSet
					_ = _48_v22
					_48_v22 = _dafny.MultiSetOf(_14_v1)
					var _49_v23 _dafny.MultiSet
					_ = _49_v23
					_49_v23 = _48_v22
					var _50_v24 _dafny.Sequence
					_ = _50_v24
					_50_v24 = _dafny.UnicodeSeqOfUtf8Bytes("kjr")
					var _51_v26 _dafny.Map
					_ = _51_v26
					_51_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _47_v21)
					var _52_v28 _dafny.Map
					_ = _52_v28
					_52_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm24(globalState), _14_v1)
					var _53_v30 _dafny.Array
					_ = _53_v30
					var _nwElement0_0 _dafny.Set = _47_v21
					_ = _nwElement0_0
					var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(15))
					_ = _nw3
					(_nw3).ArraySet1(_nwElement0_0, 0)
					(_nw3).ArraySet1(Companion_Default___.Fm27(_49_v23, _50_v24, globalState), 1)
					(_nw3).ArraySet1(_47_v21, 2)
					(_nw3).ArraySet1(_dafny.SetOf(_17_v4, _17_v4), 3)
					(_nw3).ArraySet1(Companion_Default___.Fm27(_48_v22, _50_v24, globalState), 4)
					(_nw3).ArraySet1(_47_v21, 5)
					(_nw3).ArraySet1(_47_v21, 6)
					(_nw3).ArraySet1(func() _dafny.Set {
						var _coll6 = _dafny.NewBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-30))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg8 _dafny.Int) interface{} {
								return coer8(arg8)
							}
						}((func(_54_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_55_i2 _dafny.Int) _dafny.Sequence {
								return _54_v4
							}
						})(_17_v4)))).Elements()); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _56_v25 _dafny.Sequence
							_56_v25 = interface{}(_compr_6).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-30))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg9 _dafny.Int) interface{} {
									return coer9(arg9)
								}
							}((func(_57_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_55_i2 _dafny.Int) _dafny.Sequence {
									return _57_v4
								}
							})(_17_v4))), _56_v25) {
								_coll6.Add(_56_v25)
							}
						}
						return _coll6.ToSet()
					}(), 7)
					(_nw3).ArraySet1(_47_v21, 8)
					(_nw3).ArraySet1(((func() _dafny.Set {
						if (_51_v26).Contains(_16_v3) {
							return (_51_v26).Get(_16_v3).(_dafny.Set)
						}
						return _dafny.SetOf(_17_v4, _17_v4)
					})()).Difference(_dafny.SetOf(_dafny.Companion_Sequence_.Update(_17_v4, (Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_17_v4).Cardinality()))).Uint32(), _16_v3))), 9)
					(_nw3).ArraySet1(_47_v21, 10)
					(_nw3).ArraySet1(func() _dafny.Set {
						var _coll7 = _dafny.NewBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate((_47_v21).Elements()); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _58_v27 _dafny.Sequence
							_58_v27 = interface{}(_compr_7).(_dafny.Sequence)
							if (_47_v21).Contains(_58_v27) {
								_coll7.Add(_58_v27)
							}
						}
						return _coll7.ToSet()
					}(), 11)
					(_nw3).ArraySet1(_47_v21, 12)
					(_nw3).ArraySet1(func() _dafny.Set {
						var _coll8 = _dafny.NewBuilder()
						_ = _coll8
						for _iter8 := _dafny.Iterate((_52_v28).Keys().Elements()); ; {
							_compr_8, _ok8 := _iter8()
							if !_ok8 {
								break
							}
							var _59_v29 _dafny.Sequence
							_59_v29 = interface{}(_compr_8).(_dafny.Sequence)
							if (_52_v28).Contains(_59_v29) {
								_coll8.Add(_59_v29)
							}
						}
						return _coll8.ToSet()
					}(), 13)
					(_nw3).ArraySet1(_47_v21, 14)
					_53_v30 = _nw3
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_53_v30), 0))
					_ = _index3
					(_53_v30).ArraySet1(Companion_Default___.Fm27(_48_v22, _50_v24, globalState), (_index3).Int())
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_53_v30), 0))
					_ = _index4
					(_53_v30).ArraySet1(_47_v21, (_index4).Int())
					(globalState).F9 = _dafny.IntOfInt64(162)
				} else if _source0.Is_DC29() {
					var _60___mcc_h4 bool = _source0.Get_().(D10_DC29).Cf68
					_ = _60___mcc_h4
					var _61_cf68 bool = _60___mcc_h4
					_ = _61_cf68
					_14_v1 = _14_v1
					(globalState).F9 = (_17_v4).Select((Companion_Default___.SafeIndex((_16_v3).Plus(_16_v3), _dafny.IntOfUint32((_17_v4).Cardinality()))).Uint32()).(_dafny.Int)
					var _62_v31 _dafny.Array
					_ = _62_v31
					var _len0_1 _dafny.Int = _dafny.IntOfInt64(11)
					_ = _len0_1
					var _nw4 _dafny.Array
					_ = _nw4
					if _len0_1.Cmp(_dafny.Zero) == 0 {
						_nw4 = _dafny.NewArray(_len0_1)
					} else {
						var _init1 func(_dafny.Int) _dafny.Map = (func(_63_v3 _dafny.Int, _64_v0 bool, _65_cf68 bool) func(_dafny.Int) _dafny.Map {
							return func(_66_i3 _dafny.Int) _dafny.Map {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(), _63_v3)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(), (_dafny.MultiSetOf(_64_v0, _65_cf68)).Cardinality()))
							}
						})(_16_v3, _13_v0, _61_cf68)
						_ = _init1
						var _element0_1 = _init1(_dafny.Zero)
						_ = _element0_1
						_nw4 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
						(_nw4).ArraySet1(_element0_1, 0)
						var _nativeLen0_1 = (_len0_1).Int()
						_ = _nativeLen0_1
						for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
							(_nw4).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
						}
					}
					_62_v31 = _nw4
					var _67_v32 D5
					_ = _67_v32
					_67_v32 = Companion_D5_.Create_DC14_()
					var _68_v33 _dafny.Map
					_ = _68_v33
					_68_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v32, _16_v3)
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_62_v31), 0))
					_ = _index5
					(_62_v31).ArraySet1(_68_v33, (_index5).Int())
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_62_v31), 0))
					_ = _index6
					(_62_v31).ArraySet1((_68_v33).Merge(_68_v33), (_index6).Int())
					var _69_v34 _dafny.Sequence
					_ = _69_v34
					_69_v34 = _dafny.SeqOf(true)
					var _70_v35 D1
					_ = _70_v35
					_70_v35 = Companion_D1_.Create_DC2_(_61_cf68, _13_v0, _61_cf68, _14_v1, _14_v1)
					var _71_v36 D1
					_ = _71_v36
					_71_v36 = Companion_D1_.Create_DC2_((_dafny.IntOfUint32((_69_v34).Cardinality())).Cmp(_16_v3) == 0, _13_v0, _61_cf68, Companion_Default___.Fm14(_70_v35, _dafny.IntOfUint32((_69_v34).Cardinality()), globalState), _14_v1)
					var _pat_let_tv3 = _13_v0
					_ = _pat_let_tv3
					var _pat_let_tv4 = _14_v1
					_ = _pat_let_tv4
					var _pat_let_tv5 = _61_cf68
					_ = _pat_let_tv5
					_70_v35 = func(_pat_let2_0 D1) D1 {
						return func(_75_dt__update__tmp_h3 D1) D1 {
							return func(_pat_let6_0 bool) D1 {
								return func(_76_dt__update_hcf4_h0 bool) D1 {
									return func(_pat_let7_0 bool) D1 {
										return func(_77_dt__update_hcf3_h1 bool) D1 {
											return Companion_D1_.Create_DC2_((_75_dt__update__tmp_h3).Dtor_cf2(), _77_dt__update_hcf3_h1, _76_dt__update_hcf4_h0, (_75_dt__update__tmp_h3).Dtor_cf5(), (_75_dt__update__tmp_h3).Dtor_cf6())
										}(_pat_let7_0)
									}(_pat_let_tv5)
								}(_pat_let6_0)
							}(false)
						}(_pat_let2_0)
					}(func(_pat_let3_0 D1) D1 {
						return func(_72_dt__update__tmp_h2 D1) D1 {
							return func(_pat_let4_0 bool) D1 {
								return func(_73_dt__update_hcf3_h0 bool) D1 {
									return func(_pat_let5_0 _dafny.CodePoint) D1 {
										return func(_74_dt__update_hcf6_h0 _dafny.CodePoint) D1 {
											return Companion_D1_.Create_DC2_((_72_dt__update__tmp_h2).Dtor_cf2(), _73_dt__update_hcf3_h0, (_72_dt__update__tmp_h2).Dtor_cf4(), (_72_dt__update__tmp_h2).Dtor_cf5(), _74_dt__update_hcf6_h0)
										}(_pat_let5_0)
									}(_pat_let_tv4)
								}(_pat_let4_0)
							}(_pat_let_tv3)
						}(_pat_let3_0)
					}(_70_v35))
				} else {
					var _78___mcc_h5 _dafny.Map = _source0.Get_().(D10_DC27).Cf63
					_ = _78___mcc_h5
					var _79_cf63 _dafny.Map = _78___mcc_h5
					_ = _79_cf63
					var _80_v37 _dafny.Array
					_ = _80_v37
					var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
					_ = _nw5
					_80_v37 = _nw5
					(globalState).F1 = (func() _dafny.Array {
						if _13_v0 {
							return _80_v37
						}
						return _80_v37
					})()
					var _81_v38 _dafny.MultiSet
					_ = _81_v38
					_81_v38 = _dafny.MultiSetOf(_13_v0)
					var _82_v39 *C2
					_ = _82_v39
					var _nw6 *C2 = New_C2_()
					_ = _nw6
					_nw6.Ctor__(true, _81_v38, _13_v0)
					_82_v39 = _nw6
					var _83_v40 _dafny.Map
					_ = _83_v40
					_83_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_82_v39, _16_v3)
					var _84_v41 _dafny.Set
					_ = _84_v41
					_84_v41 = _dafny.SetOf((func() _dafny.Int {
						if _13_v0 {
							return ((_83_v40).Update(_82_v39, _16_v3)).Cardinality()
						}
						return _16_v3
					})(), (_16_v3).Plus(_16_v3))
					_84_v41 = (_84_v41).Union(_84_v41)
					(globalState).F9 = (_16_v3).Plus(_16_v3)
					var _85_v42 _dafny.Sequence
					_ = _85_v42
					_85_v42 = _dafny.UnicodeSeqOfUtf8Bytes("bdgl")
					_85_v42 = _85_v42
				}
				var _86_v43 _dafny.MultiSet
				_ = _86_v43
				_86_v43 = _dafny.MultiSetOf(_13_v0, false, _13_v0, false)
				var _87_v44 T0
				_ = _87_v44
				var _nw7 *C2 = New_C2_()
				_ = _nw7
				_nw7.Ctor__(_13_v0, _86_v43, _13_v0)
				_87_v44 = _nw7
				var _88_v45 _dafny.MultiSet
				_ = _88_v45
				_88_v45 = _dafny.MultiSetOf(_87_v44, _87_v44)
				var _89_v46 _dafny.Map
				_ = _89_v46
				_89_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _87_v44.F12())
				_16_v3 = Companion_Default___.Fm1(Companion_Default___.Fm0(_14_v1, globalState), (_dafny.Zero).Minus((_88_v45).Cardinality()), (_89_v46).Cardinality(), globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _90_i4 _dafny.Int
	_ = _90_i4
	_90_i4 = _dafny.Zero
	{
		for _13_v0 {
			{
				if (_90_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_90_i4 = (_90_i4).Plus(_dafny.One)
				var _91_v47 _dafny.Sequence
				_ = _91_v47
				_91_v47 = _dafny.SeqOf((_13_v0) == (_13_v0), _13_v0, _13_v0)
				if (_91_v47).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_91_v47).Cardinality()))).Uint32()).(bool) {
					var _92_v48 _dafny.MultiSet
					_ = _92_v48
					_92_v48 = _dafny.MultiSetOf(_13_v0, _13_v0)
					var _93_v49 *C2
					_ = _93_v49
					var _nw8 *C2 = New_C2_()
					_ = _nw8
					_nw8.Ctor__(_13_v0, _92_v48, false)
					_93_v49 = _nw8
					_92_v48 = (func() _dafny.MultiSet {
						if !((_16_v3).Cmp(_16_v3) < 0) {
							return _92_v48
						}
						return _92_v48
					})()
					(globalState).F9 = _16_v3
					var _94_v50 _dafny.Sequence
					_ = _94_v50
					_94_v50 = _dafny.UnicodeSeqOfUtf8Bytes("s")
					var _95_v51 _dafny.Map
					_ = _95_v51
					_95_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _93_v49.F15)
					var _96_v52 _dafny.Sequence
					_ = _96_v52
					_96_v52 = _dafny.SeqOf(_95_v51)
					var _97_v53 _dafny.Map
					_ = _97_v53
					_97_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_94_v50, _94_v50), (Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_94_v50, _94_v50)).Cardinality()))).Uint32(), _14_v1)).Cardinality())), (_16_v3).Minus(((_96_v52).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_96_v52).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()))
					_97_v53 = (_97_v53).Update(_16_v3, _16_v3)
					var _98_v54 T0
					_ = _98_v54
					var _nw9 *C1 = New_C1_()
					_ = _nw9
					_nw9.Ctor__(_16_v3, _92_v48, _93_v49.F15)
					_98_v54 = _nw9
					var _99_v55 _dafny.Map
					_ = _99_v55
					_99_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v54, _17_v4)
					var _100_v56 _dafny.Set
					_ = _100_v56
					_100_v56 = _dafny.SetOf(_16_v3, _16_v3, (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
						if (_99_v55).Contains(_98_v54) {
							return (_99_v55).Get(_98_v54).(_dafny.Sequence)
						}
						return _17_v4
					})())).Cardinality(), (_dafny.SetOf(_14_v1, _14_v1)).Cardinality())
					var _101_v57 _dafny.Map
					_ = _101_v57
					_101_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _92_v48)
					var _102_v58 *C3
					_ = _102_v58
					var _nw10 *C3 = New_C3_()
					_ = _nw10
					_nw10.Ctor__((_dafny.MultiSetOf(_13_v0, _13_v0)).Cardinality(), (_100_v56).IsProperSubsetOf(_100_v56), (func() _dafny.MultiSet {
						if (_101_v57).Contains(_dafny.CodePoint('y')) {
							return (_101_v57).Get(_dafny.CodePoint('y')).(_dafny.MultiSet)
						}
						return _92_v48
					})(), (_dafny.SetOf(_16_v3)).IsProperSubsetOf(_dafny.SetOf(_16_v3)))
					_102_v58 = _nw10
				} else {
					var _103_v59 bool
					_ = _103_v59
					_103_v59 = !(_13_v0) || (true)
					_103_v59 = !(_13_v0) || (_13_v0)
					(globalState).F7 = _13_v0
					var _104_v60 *C0
					_ = _104_v60
					var _nw11 *C0 = New_C0_()
					_ = _nw11
					_nw11.Ctor__((func() bool {
						if _13_v0 {
							return _13_v0
						}
						return _13_v0
					})(), _16_v3, _dafny.MultiSetOf(true, _13_v0), !_dafny.Companion_Sequence_.Contains(_17_v4, _16_v3))
					_104_v60 = _nw11
					_104_v60 = _104_v60
					var _105_v61 _dafny.MultiSet
					_ = _105_v61
					_105_v61 = _dafny.MultiSetOf(_13_v0)
					var _106_v62 *C2
					_ = _106_v62
					var _nw12 *C2 = New_C2_()
					_ = _nw12
					_nw12.Ctor__(_13_v0, (_105_v61).Union(_dafny.MultiSetOf((_104_v60).F17())), _13_v0)
					_106_v62 = _nw12
					var _107_v63 _dafny.Sequence
					_ = _107_v63
					_107_v63 = _dafny.UnicodeSeqOfUtf8Bytes("r")
					var _108_v64 _dafny.MultiSet
					_ = _108_v64
					_108_v64 = _dafny.MultiSetOf(Companion_Default___.Fm1((_104_v60).F17(), (_104_v60).F18(), (_15_v2).Cardinality(), globalState))
					var _109_v65 _dafny.Map
					_ = _109_v65
					_109_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _16_v3)
					var _rhs6 _dafny.Sequence = Companion_Default___.Fm9(_16_v3, !((_104_v60).F17()) || ((_104_v60).Fm2(_108_v64, (_dafny.Zero).Minus(_16_v3), _109_v65, _13_v0, globalState)), globalState)
					_ = _rhs6
					var _rhs7 _dafny.Int = _16_v3
					_ = _rhs7
					var _rhs8 _dafny.Int = ((_dafny.IntOfInt64(350)).Plus(_16_v3)).Times(_16_v3)
					_ = _rhs8
					var _lhs4 *GlobalState = globalState
					_ = _lhs4
					_107_v63 = _rhs6
					_lhs4.F9 = _rhs7
					_16_v3 = _rhs8
				}
				var _110_v66 _dafny.Array
				_ = _110_v66
				var _len0_2 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_2
				var _nw13 _dafny.Array
				_ = _nw13
				if _len0_2.Cmp(_dafny.Zero) == 0 {
					_nw13 = _dafny.NewArray(_len0_2)
				} else {
					var _init2 func(_dafny.Int) bool = (func(_111_v0 bool) func(_dafny.Int) bool {
						return func(_112_i5 _dafny.Int) bool {
							return (func() bool {
								if false {
									return _111_v0
								}
								return _111_v0
							})()
						}
					})(_13_v0)
					_ = _init2
					var _element0_2 = _init2(_dafny.Zero)
					_ = _element0_2
					_nw13 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
					(_nw13).ArraySet1(_element0_2, 0)
					var _nativeLen0_2 = (_len0_2).Int()
					_ = _nativeLen0_2
					for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
						(_nw13).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
					}
				}
				_110_v66 = _nw13
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_110_v66), 0))
				_ = _index7
				(_110_v66).ArraySet1(!(_13_v0), (_index7).Int())
				var _113_v67 _dafny.Map
				_ = _113_v67
				_113_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _16_v3)
				var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_110_v66), 0))
				_ = _index8
				(_110_v66).ArraySet1(((_113_v67).Merge(_113_v67)).Contains(true), (_index8).Int())
				(globalState).F6 = _91_v47
				var _114_v68 _dafny.MultiSet
				_ = _114_v68
				_114_v68 = _dafny.MultiSetOf(_14_v1)
				var _115_v69 _dafny.Sequence
				_ = _115_v69
				_115_v69 = _dafny.SeqOf(_14_v1)
				var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_110_v66), 0))
				_ = _index9
				(_110_v66).ArraySet1((_dafny.MultiSetFromSeq(_115_v69)).IsProperSubsetOf(_114_v68), (_index9).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _116_v70 _dafny.Sequence
	_ = _116_v70
	_116_v70 = _dafny.UnicodeSeqOfUtf8Bytes("dslsfkrf")
	var _117_v71 _dafny.Map
	_ = _117_v71
	_117_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v70, _dafny.IntOfUint32((_116_v70).Cardinality()))
	var _118_v72 _dafny.Set
	_ = _118_v72
	_118_v72 = _dafny.SetOf(Companion_Default___.Fm0(_14_v1, globalState))
	var _119_v73 _dafny.MultiSet
	_ = _119_v73
	_119_v73 = _dafny.MultiSetOf(_13_v0)
	var _120_v74 _dafny.Array
	_ = _120_v74
	var _nwElement0_1 bool = _13_v0
	_ = _nwElement0_1
	var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
	_ = _nw14
	(_nw14).ArraySet1(_nwElement0_1, 0)
	(_nw14).ArraySet1(_13_v0, 1)
	(_nw14).ArraySet1((true) == (_13_v0), 2)
	(_nw14).ArraySet1(_13_v0, 3)
	(_nw14).ArraySet1(_13_v0, 4)
	(_nw14).ArraySet1(_13_v0, 5)
	(_nw14).ArraySet1(!(_117_v71).Equals((_117_v71).Update(_116_v70, _16_v3)), 6)
	(_nw14).ArraySet1(_13_v0, 7)
	(_nw14).ArraySet1(!(_118_v72).Contains(_13_v0), 8)
	(_nw14).ArraySet1(false, 9)
	(_nw14).ArraySet1(_13_v0, 10)
	(_nw14).ArraySet1(!(_13_v0), 11)
	(_nw14).ArraySet1((_118_v72).IsDisjointFrom(_118_v72), 12)
	(_nw14).ArraySet1(true, 13)
	(_nw14).ArraySet1(Companion_Default___.Fm0(_14_v1, globalState), 14)
	(_nw14).ArraySet1(true, 15)
	(_nw14).ArraySet1((func() bool {
		if _13_v0 {
			return _13_v0
		}
		return false
	})(), 16)
	(_nw14).ArraySet1(_13_v0, 17)
	(_nw14).ArraySet1(_13_v0, 18)
	(_nw14).ArraySet1(_13_v0, 19)
	(_nw14).ArraySet1(((_119_v73).Cardinality()).Cmp(_dafny.IntOfUint32((_116_v70).Cardinality())) == 0, 20)
	(_nw14).ArraySet1((_13_v0), 21)
	(_nw14).ArraySet1(_13_v0, 22)
	(_nw14).ArraySet1(_13_v0, 23)
	(_nw14).ArraySet1(_13_v0, 24)
	(_nw14).ArraySet1(_13_v0, 25)
	_120_v74 = _nw14
	var _121_v75 _dafny.Sequence
	_ = _121_v75
	_121_v75 = _dafny.SeqOf(_120_v74, _120_v74, (func() _dafny.Array {
		if _13_v0 {
			return _120_v74
		}
		return _120_v74
	})())
	var _122_v76 _dafny.MultiSet
	_ = _122_v76
	_122_v76 = _dafny.MultiSetOf(_16_v3, _16_v3, _16_v3, Companion_Default___.Fm1(_13_v0, _16_v3, _dafny.IntOfInt64(496), globalState))
	var _rhs9 _dafny.Array = (_121_v75).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}((func(_123_v70 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
		return func(_124_i6 _dafny.Int) _dafny.CodePoint {
			return (_123_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.IntOfUint32((_123_v70).Cardinality()))).Uint32()).(_dafny.CodePoint)
		}
	})(_116_v70)))).Cardinality()), _16_v3), _dafny.IntOfUint32((_121_v75).Cardinality()))).Uint32()).(_dafny.Array)
	_ = _rhs9
	var _rhs10 _dafny.Array = _120_v74
	_ = _rhs10
	var _rhs11 bool = (func() bool {
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_116_v70, _116_v70) {
			return !(_13_v0)
		}
		return (_122_v76).Contains(_16_v3)
	})()
	_ = _rhs11
	_120_v74 = _rhs9
	_120_v74 = _rhs10
	_13_v0 = _rhs11
	var _125_v77 _dafny.Array
	_ = _125_v77
	var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
	_ = _nw15
	_125_v77 = _nw15
	var _126_v78 D5
	_ = _126_v78
	_126_v78 = Companion_D5_.Create_DC13_(_16_v3, _13_v0, _16_v3, _125_v77, _13_v0)
	var _hi0 _dafny.Int = (_126_v78).Dtor_cf30()
	_ = _hi0
	for _127_i7 := _16_v3; _127_i7.Cmp(_hi0) < 0; _127_i7 = _127_i7.Plus(_dafny.One) {
		var _128_v79 _dafny.MultiSet
		_ = _128_v79
		_128_v79 = (_dafny.MultiSetOf(_14_v1)).Intersection(_dafny.MultiSetOf(_dafny.CodePoint('g')))
		var _source1 _dafny.MultiSet = _128_v79
		_ = _source1
		var _129___mcc_h6 _dafny.MultiSet = _source1
		_ = _129___mcc_h6
		var _130_cf77 _dafny.MultiSet = _129___mcc_h6
		_ = _130_cf77
		var _131_v80 _dafny.Sequence
		_ = _131_v80
		_131_v80 = _dafny.SeqOf(true)
		var _132_v81 _dafny.Sequence
		_ = _132_v81
		_132_v81 = _dafny.SeqOf(_13_v0, _13_v0, false, (_131_v80).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jcvbfjwrl")).Cardinality())), _dafny.IntOfUint32((_131_v80).Cardinality()))).Uint32()).(bool))
		(globalState).F9 = (func() _dafny.Int {
			if (_122_v76).Contains(_127_i7) {
				return _16_v3
			}
			return (_127_i7).Plus(_dafny.IntOfUint32((_132_v81).Cardinality()))
		})()
		var _133_v82 _dafny.Array
		_ = _133_v82
		var _len0_3 _dafny.Int = _dafny.One
		_ = _len0_3
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_134_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_135_i8 _dafny.Int) _dafny.Int {
					return (_135_i8).Plus(_134_v3)
				}
			})(_16_v3)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw16 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw16).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw16).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_133_v82 = _nw16
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_133_v82), 0))
		_ = _index10
		(_133_v82).ArraySet1((_16_v3).Times(_16_v3), (_index10).Int())
		var _136_v83 _dafny.MultiSet
		_ = _136_v83
		_136_v83 = _dafny.MultiSetOf(_133_v82, _133_v82)
		var _137_v84 _dafny.Set
		_ = _137_v84
		_137_v84 = _dafny.SetOf(_127_i7, _127_i7, _16_v3)
		var _138_v85 D2
		_ = _138_v85
		_138_v85 = Companion_D2_.Create_DC5_((_dafny.Zero).Minus(_127_i7), (_136_v83).Cardinality(), (_137_v84).Cardinality(), ((_dafny.MultiSetFromSeq(_131_v80)).Update(_13_v0, Companion_Default___.Abs(Companion_Default___.Fm1(_13_v0, _127_i7, _16_v3, globalState)))).Cardinality(), _132_v81)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_133_v82), 0))
		_ = _index11
		var _rhs12 _dafny.Array = _133_v82
		_ = _rhs12
		var _rhs13 _dafny.Int = Companion_Default___.Fm1(_13_v0, (_16_v3).Times(_16_v3), _16_v3, globalState)
		_ = _rhs13
		var _rhs14 D2 = _138_v85
		_ = _rhs14
		var _lhs5 *GlobalState = globalState
		_ = _lhs5
		var _lhs6 _dafny.Array = _133_v82
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_133_v82), 0))
		_ = _lhs7
		_lhs5.F4 = _rhs12
		(_lhs6).ArraySet1(_rhs13, (_lhs7).Int())
		_138_v85 = _rhs14
		_14_v1 = (func() _dafny.CodePoint {
			if true {
				return _14_v1
			}
			return _14_v1
		})()
		var _139_v86 _dafny.Map
		_ = _139_v86
		_139_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v1, _16_v3)
		_139_v86 = (_139_v86).Update(_14_v1, Companion_Default___.Fm1(_13_v0, (_133_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_133_v82), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_133_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_133_v82), 0))).Int()).(_dafny.Int)), globalState))
		_13_v0 = false
		var _140_v87 _dafny.Sequence
		_ = _140_v87
		_140_v87 = _dafny.SeqOf(_dafny.SeqOf(_13_v0))
		var _141_v88 _dafny.Map
		_ = _141_v88
		_141_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_140_v87).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_140_v87).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(529))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_142_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_143_i9 _dafny.Int) _dafny.CodePoint {
				return _142_v1
			}
		})(_14_v1))), _116_v70, _116_v70, _dafny.UnicodeSeqOfUtf8Bytes("bultv"), _dafny.UnicodeSeqOfUtf8Bytes("yvdibk")), _dafny.SeqOf(_116_v70, _116_v70))))
		var _144_v89 _dafny.Sequence
		_ = _144_v89
		_144_v89 = _dafny.SeqOf(_13_v0)
		var _145_v90 _dafny.MultiSet
		_ = _145_v90
		_145_v90 = _dafny.MultiSetOf(_116_v70, _116_v70, _116_v70)
		_141_v88 = (_141_v88).Update(_144_v89, _145_v90)
		if !(!(_13_v0)) {
			var _146_v91 *C4
			_ = _146_v91
			var _nw17 *C4 = New_C4_()
			_ = _nw17
			_nw17.Ctor__((_dafny.MultiSetFromSeq(_144_v89)).Union(_119_v73), !(!(_13_v0)))
			_146_v91 = _nw17
			(globalState).F9 = Companion_Default___.SafeDivisionInt(_127_i7, _127_i7)
			_116_v70 = _116_v70
			var _147_v92 _dafny.Set
			_ = _147_v92
			_147_v92 = _dafny.SetOf(_127_i7)
			var _148_v93 D3
			_ = _148_v93
			_148_v93 = Companion_D3_.Create_DC8_((_dafny.Zero).Minus((_dafny.Zero).Minus(_16_v3)), _127_i7)
			var _149_v94 _dafny.Map
			_ = _149_v94
			_149_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_147_v92).Contains((_147_v92).Cardinality()), _148_v93)
			_149_v94 = (_149_v94).Update(_13_v0, _148_v93)
			var _150_v95 _dafny.Map
			_ = _150_v95
			_150_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("o"), _dafny.UnicodeSeqOfUtf8Bytes("uli"))
			var _151_v96 _dafny.Map
			_ = _151_v96
			_151_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29((_144_v89).Select((Companion_Default___.SafeIndex((_147_v92).Cardinality(), _dafny.IntOfUint32((_144_v89).Cardinality()))).Uint32()).(bool), _16_v3, globalState), (_dafny.SetOf(_dafny.IntOfInt64(917), _dafny.IntOfUint32((_17_v4).Cardinality()), (_146_v91).Fm3(_13_v0, _150_v95, _13_v0, _127_i7, globalState))).Cardinality())
			var _152_v97 _dafny.Map
			_ = _152_v97
			_152_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _127_i7)
			var _153_v98 _dafny.Map
			_ = _153_v98
			_153_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_i7, _152_v97)
			var _154_v99 _dafny.Map
			_ = _154_v99
			_154_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _13_v0)
			_151_v96 = (_151_v96).Update((_153_v98).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _152_v97)), ((_154_v99).Cardinality()).Times(_127_i7))
		} else {
			var _155_v100 _dafny.Array
			_ = _155_v100
			var _nw18 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
			_ = _nw18
			_155_v100 = _nw18
			var _156_v101 D9
			_ = _156_v101
			_156_v101 = Companion_D9_.Create_DC23_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _155_v100))
			var _157_v102 _dafny.Map
			_ = _157_v102
			_157_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _155_v100))
			var _158_v103 _dafny.Map
			_ = _158_v103
			_158_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _155_v100)
			var _pat_let_tv6 = _157_v102
			_ = _pat_let_tv6
			var _pat_let_tv7 = _16_v3
			_ = _pat_let_tv7
			var _pat_let_tv8 = _157_v102
			_ = _pat_let_tv8
			var _pat_let_tv9 = _16_v3
			_ = _pat_let_tv9
			var _pat_let_tv10 = _158_v103
			_ = _pat_let_tv10
			var _159_v104 _dafny.Set
			_ = _159_v104
			_159_v104 = _dafny.SetOf(func(_pat_let8_0 D9) D9 {
				return func(_160_dt__update__tmp_h6 D9) D9 {
					return func(_pat_let9_0 _dafny.Map) D9 {
						return func(_161_dt__update_hcf55_h0 _dafny.Map) D9 {
							return Companion_D9_.Create_DC23_(_161_dt__update_hcf55_h0)
						}(_pat_let9_0)
					}((func() _dafny.Map {
						if (_pat_let_tv6).Contains(_pat_let_tv7) {
							return (_pat_let_tv8).Get(_pat_let_tv9).(_dafny.Map)
						}
						return _pat_let_tv10
					})())
				}(_pat_let8_0)
			}(_156_v101), _156_v101, (func() D9 {
				if true {
					return _156_v101
				}
				return Companion_D9_.Create_DC23_(_158_v103)
			})(), _156_v101, _156_v101)
			_159_v104 = (_159_v104).Union(_159_v104)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))
			_ = _index12
			(_120_v74).ArraySet1(_13_v0, (_index12).Int())
			var _162_v105 _dafny.Map
			_ = _162_v105
			_162_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
				if _13_v0 {
					return _13_v0
				}
				return !(_13_v0)
			})()), !((_13_v0) && (_13_v0)))
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))
			_ = _index13
			(_120_v74).ArraySet1((func() bool {
				if (_162_v105).Contains((func() bool {
					if !(_13_v0) {
						return _13_v0
					}
					return _13_v0
				})()) {
					return (_162_v105).Get((func() bool {
						if !(_13_v0) {
							return _13_v0
						}
						return _13_v0
					})()).(bool)
				}
				return _13_v0
			})(), (_index13).Int())
			var _163_v106 D9
			_ = _163_v106
			_163_v106 = Companion_D9_.Create_DC24_(_116_v70, _13_v0, _16_v3)
			var _164_v107 _dafny.Array
			_ = _164_v107
			var _nwElement0_2 bool = (_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool)
			_ = _nwElement0_2
			var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(20))
			_ = _nw19
			(_nw19).ArraySet1(_nwElement0_2, 0)
			(_nw19).ArraySet1(_13_v0, 1)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 2)
			(_nw19).ArraySet1(false, 3)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 4)
			(_nw19).ArraySet1((_163_v106).Dtor_cf57(), 5)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 6)
			(_nw19).ArraySet1(!(_13_v0), 7)
			(_nw19).ArraySet1((_144_v89).Select((Companion_Default___.SafeIndex(_16_v3, _dafny.IntOfUint32((_144_v89).Cardinality()))).Uint32()).(bool), 8)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 9)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 10)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 11)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 12)
			(_nw19).ArraySet1((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), 13)
			(_nw19).ArraySet1(_13_v0, 14)
			(_nw19).ArraySet1(_13_v0, 15)
			(_nw19).ArraySet1(_13_v0, 16)
			(_nw19).ArraySet1(_13_v0, 17)
			(_nw19).ArraySet1(_13_v0, 18)
			(_nw19).ArraySet1(_13_v0, 19)
			_164_v107 = _nw19
			var _165_v108 _dafny.Set
			_ = _165_v108
			_165_v108 = _dafny.SetOf(_164_v107, _164_v107)
			var _166_v109 _dafny.Map
			_ = _166_v109
			_166_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_i7, (_162_v105).Cardinality())
			var _167_v110 _dafny.Map
			_ = _167_v110
			_167_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _166_v109)
			var _168_v111 _dafny.Map
			_ = _168_v111
			_168_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_162_v105).Cardinality()).Cmp((_165_v108).Cardinality()) != 0, ((_167_v110).Update((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(305))).Cardinality())))).Merge(_167_v110))
			_168_v111 = (_168_v111).Update((func() bool {
				if (_144_v89).Select((Companion_Default___.SafeIndex(_127_i7, _dafny.IntOfUint32((_144_v89).Cardinality()))).Uint32()).(bool) {
					return (_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool)
				}
				return _13_v0
			})(), (_167_v110).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _166_v109)))
			var _169_v112 *C1
			_ = _169_v112
			var _nw20 *C1 = New_C1_()
			_ = _nw20
			_nw20.Ctor__(_dafny.IntOfInt64(300), _dafny.MultiSetOf((_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool)), _13_v0)
			_169_v112 = _nw20
			var _170_v113 _dafny.Map
			_ = _170_v113
			_170_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _16_v3)
			var _171_v114 _dafny.Array
			_ = _171_v114
			var _nwElement0_3 _dafny.Int = (_170_v113).Cardinality()
			_ = _nwElement0_3
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(10))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_3, 0)
			(_nw21).ArraySet1(_16_v3, 1)
			(_nw21).ArraySet1(_16_v3, 2)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_144_v89).Cardinality()), 3)
			(_nw21).ArraySet1(_127_i7, 4)
			(_nw21).ArraySet1(_dafny.IntOfInt64(93), 5)
			(_nw21).ArraySet1((_17_v4).Select((Companion_Default___.SafeIndex(_169_v112.F16, _dafny.IntOfUint32((_17_v4).Cardinality()))).Uint32()).(_dafny.Int), 6)
			(_nw21).ArraySet1(_169_v112.F16, 7)
			(_nw21).ArraySet1((_dafny.Zero).Minus(_16_v3), 8)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_17_v4).Cardinality()), 9)
			_171_v114 = _nw21
			var _172_v115 _dafny.Sequence
			_ = _172_v115
			_172_v115 = _dafny.SeqOf(_171_v114)
			var _173_v116 _dafny.Array
			_ = _173_v116
			var _nwElement0_4 _dafny.Array = _164_v107
			_ = _nwElement0_4
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(17))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_4, 0)
			(_nw22).ArraySet1(_164_v107, 1)
			(_nw22).ArraySet1(_164_v107, 2)
			(_nw22).ArraySet1(_164_v107, 3)
			(_nw22).ArraySet1(_164_v107, 4)
			(_nw22).ArraySet1(_164_v107, 5)
			(_nw22).ArraySet1(_164_v107, 6)
			(_nw22).ArraySet1(_164_v107, 7)
			(_nw22).ArraySet1(_164_v107, 8)
			(_nw22).ArraySet1(_164_v107, 9)
			(_nw22).ArraySet1(_164_v107, 10)
			(_nw22).ArraySet1(_164_v107, 11)
			(_nw22).ArraySet1(_164_v107, 12)
			(_nw22).ArraySet1(_164_v107, 13)
			(_nw22).ArraySet1(_164_v107, 14)
			(_nw22).ArraySet1(_164_v107, 15)
			(_nw22).ArraySet1(_164_v107, 16)
			_173_v116 = _nw22
			var _174_v117 D8
			_ = _174_v117
			_174_v117 = Companion_D8_.Create_DC21_(!(_13_v0), _173_v116, _16_v3, _164_v107, _16_v3)
			var _rhs15 *C1 = _169_v112
			_ = _rhs15
			var _rhs16 _dafny.Array = (func() _dafny.Array {
				if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_172_v115, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_169_v112.F16), _dafny.IntOfUint32((_172_v115).Cardinality()))).Uint32(), _171_v114), _172_v115) {
					return (_174_v117).Dtor_cf52()
				}
				return _164_v107
			})()
			_ = _rhs16
			var _rhs17 bool = (_120_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_120_v74), 0))).Int()).(bool)
			_ = _rhs17
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			_169_v112 = _rhs15
			_164_v107 = _rhs16
			_lhs8.F8 = _rhs17
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_171_v114), 0))
			_ = _index14
			(_171_v114).ArraySet1(_169_v112.F16, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_171_v114), 0))
			_ = _index15
			(_171_v114).ArraySet1(_127_i7, (_index15).Int())
		}
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _175_v1 _dafny.Int
	_ = _175_v1
	_175_v1 = _dafny.IntOfInt64(-351)
	var _176_v2 _dafny.Sequence
	_ = _176_v2
	_176_v2 = _dafny.SeqOf(_175_v1, _dafny.IntOfInt64(231), _175_v1)
	var _177_v3 _dafny.CodePoint
	_ = _177_v3
	_177_v3 = _dafny.CodePoint('p')
	var _178_v4 _dafny.Map
	_ = _178_v4
	_178_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_176_v2).Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _179_v0 _dafny.Int
			_179_v0 = interface{}(_compr_9).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_176_v2, _179_v0) {
				_coll9.Add((_179_v0).Minus(_175_v1), false)
			}
		}
		return _coll9.ToMap()
	}(), _177_v3)
	var _180_v5 _dafny.Array
	_ = _180_v5
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_4
	var _nw23 _dafny.Array
	_ = _nw23
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw23 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = (func(_181_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_182_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_182_i0, _181_v1)
			}
		})(_175_v1)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw23 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw23).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw23).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_180_v5 = _nw23
	var _183_v6 bool
	_ = _183_v6
	_183_v6 = false
	var _184_v7 _dafny.Sequence
	_ = _184_v7
	_184_v7 = _dafny.SeqOf(_183_v6, true)
	var _185_globalState *GlobalState
	_ = _185_globalState
	var _nw24 *GlobalState = New_GlobalState_()
	_ = _nw24
	_nw24.Ctor__(_178_v4, _180_v5, true, false, _180_v5, _dafny.IntOfInt64(724), _184_v7, true, true, _dafny.IntOfInt64(-663), false)
	_185_globalState = _nw24
	Companion_Default___.M0(_185_globalState)
	var _186_i1 _dafny.Int
	_ = _186_i1
	_186_i1 = _dafny.Zero
	{
		for (func() bool {
			if false {
				return _183_v6
			}
			return !(!(!(_183_v6)))
		})() {
			{
				if (_186_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_186_i1 = (_186_i1).Plus(_dafny.One)
				(_185_globalState).F8 = Companion_Default___.Fm0(_177_v3, _185_globalState)
				var _187_v8 _dafny.Map
				_ = _187_v8
				_187_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_175_v1, _175_v1), Companion_Default___.Fm0(_177_v3, _185_globalState))
				var _188_v10 _dafny.Map
				_ = _188_v10
				_188_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, _175_v1)
				_187_v8 = func() _dafny.Map {
					var _coll10 = _dafny.NewMapBuilder()
					_ = _coll10
					for _iter10 := _dafny.Iterate((_dafny.SeqOf(Companion_Default___.Fm1(_183_v6, _175_v1, (_188_v10).Cardinality(), _185_globalState))).Elements()); ; {
						_compr_10, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _189_v9 _dafny.Int
						_189_v9 = interface{}(_compr_10).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_Default___.Fm1(_183_v6, _175_v1, (_188_v10).Cardinality(), _185_globalState)), _189_v9) {
							_coll10.Add(Companion_Default___.SafeDivisionInt(_189_v9, (_dafny.MultiSetOf(_175_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kengs")).Cardinality()), (func() _dafny.Set {
								var _coll11 = _dafny.NewBuilder()
								_ = _coll11
								for _iter11 := _dafny.Iterate((_187_v8).Keys().Elements()); ; {
									_compr_11, _ok11 := _iter11()
									if !_ok11 {
										break
									}
									var _190_v11 _dafny.Int
									_190_v11 = interface{}(_compr_11).(_dafny.Int)
									if (_187_v8).Contains(_190_v11) {
										_coll11.Add((_190_v11).Minus(_dafny.IntOfInt64(206)))
									}
								}
								return _coll11.ToSet()
							}()).Cardinality())).Cardinality()), (_dafny.MultiSetOf(_183_v6, _183_v6)).IsProperSubsetOf(_dafny.MultiSetOf(_183_v6, _183_v6)))
						}
					}
					return _coll10.ToMap()
				}()
				var _191_v12 _dafny.MultiSet
				_ = _191_v12
				_191_v12 = _dafny.MultiSetOf(_183_v6)
				var _192_v13 *C4
				_ = _192_v13
				var _nw25 *C4 = New_C4_()
				_ = _nw25
				_nw25.Ctor__(_191_v12, _183_v6)
				_192_v13 = _nw25
				if _183_v6 {
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_180_v5), 0))
					_ = _index16
					(_180_v5).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}((func(_193_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_194_i2 _dafny.Int) _dafny.CodePoint {
							return _193_v3
						}
					})(_177_v3)))).Cardinality())).Minus(_175_v1), (_index16).Int())
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_180_v5), 0))
					_ = _index17
					(_180_v5).ArraySet1(_175_v1, (_index17).Int())
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_180_v5), 0))
					_ = _index18
					(_180_v5).ArraySet1(_175_v1, (_index18).Int())
					(_185_globalState).F9 = _175_v1
					var _195_v14 *C4
					_ = _195_v14
					var _nw26 *C4 = New_C4_()
					_ = _nw26
					_nw26.Ctor__(_191_v12, _183_v6)
					_195_v14 = _nw26
					(_185_globalState).F7 = _183_v6
				} else {
					var _196_v15 _dafny.Sequence
					_ = _196_v15
					_196_v15 = _dafny.UnicodeSeqOfUtf8Bytes("fuaink")
					var _197_v16 _dafny.MultiSet
					_ = _197_v16
					_197_v16 = _dafny.MultiSetOf(_196_v15, (func() _dafny.Sequence {
						if _183_v6 {
							return _196_v15
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg13 _dafny.Int) interface{} {
								return coer13(arg13)
							}
						}((func(_198_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_199_i3 _dafny.Int) _dafny.CodePoint {
								return _198_v3
							}
						})(_177_v3)))
					})(), _dafny.Companion_Sequence_.Concatenate(_196_v15, _dafny.UnicodeSeqOfUtf8Bytes("tdvl")), _196_v15, (func() _dafny.Sequence {
						if (func() bool {
							if (_187_v8).Contains((_dafny.Zero).Minus(_175_v1)) {
								return (_187_v8).Get((_dafny.Zero).Minus(_175_v1)).(bool)
							}
							return _183_v6
						})() {
							return _196_v15
						}
						return (Companion_D9_.Create_DC25_(_196_v15, _183_v6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gsmf")).Cardinality()))).Dtor_cf59()
					})())
					(_185_globalState).F9 = (_197_v16).Cardinality()
					_188_v10 = (_188_v10).Update(_183_v6, (_192_v13).Fm3(_183_v6, Companion_Default___.Fm25(_177_v3, _177_v3, !(_183_v6), _185_globalState), !(_183_v6), _175_v1, _185_globalState))
					var _200_v17 *C2
					_ = _200_v17
					var _nw27 *C2 = New_C2_()
					_ = _nw27
					_nw27.Ctor__(_183_v6, _191_v12, _183_v6)
					_200_v17 = _nw27
					var _201_v18 T0
					_ = _201_v18
					var _nw28 *C4 = New_C4_()
					_ = _nw28
					_nw28.Ctor__(_dafny.MultiSetOf(_200_v17.F15), _200_v17.F15)
					_201_v18 = _nw28
					var _202_v19 _dafny.Map
					_ = _202_v19
					_202_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _201_v18)
					var _203_v20 D10
					_ = _203_v20
					_203_v20 = Companion_D10_.Create_DC28_(_200_v17, _200_v17.F15, _201_v18.F12(), _201_v18.F12())
					var _204_v21 _dafny.Array
					_ = _204_v21
					var _nwElement0_5 bool = true
					_ = _nwElement0_5
					var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(8))
					_ = _nw29
					(_nw29).ArraySet1(_nwElement0_5, 0)
					(_nw29).ArraySet1(_200_v17.F15, 1)
					(_nw29).ArraySet1(_200_v17.F15, 2)
					(_nw29).ArraySet1((_175_v1).Cmp((_202_v19).Cardinality()) != 0, 3)
					(_nw29).ArraySet1((_203_v20).Dtor_cf67(), 4)
					(_nw29).ArraySet1(true, 5)
					(_nw29).ArraySet1(_200_v17.F15, 6)
					(_nw29).ArraySet1(_201_v18.F12(), 7)
					_204_v21 = _nw29
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_204_v21), 0))
					_ = _index19
					(_204_v21).ArraySet1((_175_v1).Cmp(_dafny.IntOfInt64(489)) >= 0, (_index19).Int())
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_204_v21), 0))
					_ = _index20
					(_204_v21).ArraySet1(_201_v18.F12(), (_index20).Int())
					(_185_globalState).F7 = _200_v17.F15
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _205_i4 _dafny.Int
	_ = _205_i4
	_205_i4 = _dafny.Zero
	{
		for _183_v6 {
			{
				if (_205_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_205_i4 = (_205_i4).Plus(_dafny.One)
				var _206_v22 _dafny.Sequence
				_ = _206_v22
				_206_v22 = _dafny.UnicodeSeqOfUtf8Bytes("lrpi")
				var _207_v23 _dafny.MultiSet
				_ = _207_v23
				_207_v23 = _dafny.MultiSetOf(_183_v6)
				var _208_v24 T0
				_ = _208_v24
				var _nw30 *C3 = New_C3_()
				_ = _nw30
				_nw30.Ctor__(_dafny.IntOfUint32((_206_v22).Cardinality()), !(_183_v6), _207_v23, _183_v6)
				_208_v24 = _nw30
				_208_v24 = (func() T0 {
					if _208_v24.F12() {
						return _208_v24
					}
					return _208_v24
				})()
				var _209_v25 _dafny.Array
				_ = _209_v25
				var _len0_5 _dafny.Int = _dafny.One
				_ = _len0_5
				var _nw31 _dafny.Array
				_ = _nw31
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw31 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_210_v24 T0) func(_dafny.Int) _dafny.MultiSet {
						return func(_211_i5 _dafny.Int) _dafny.MultiSet {
							return _210_v24.F11()
						}
					})(_208_v24)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw31 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw31).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw31).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_209_v25 = _nw31
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_209_v25), 0))
				_ = _index21
				(_209_v25).ArraySet1(_dafny.MultiSetOf(!(true), _208_v24.F12()), (_index21).Int())
				var _212_v26 *C2
				_ = _212_v26
				var _nw32 *C2 = New_C2_()
				_ = _nw32
				_nw32.Ctor__(_208_v24.F12(), Companion_Default___.Fm22(_175_v1, !(_208_v24.F12()), _185_globalState), _183_v6)
				_212_v26 = _nw32
				var _213_v27 _dafny.Sequence
				_ = _213_v27
				_213_v27 = _dafny.SeqOf(_212_v26, _212_v26, _212_v26)
				var _214_v28 *C0
				_ = _214_v28
				var _nw33 *C0 = New_C0_()
				_ = _nw33
				_nw33.Ctor__(_183_v6, _175_v1, _dafny.MultiSetOf(_212_v26.F15, _208_v24.F12()), _208_v24.F12())
				_214_v28 = _nw33
				var _215_v29 _dafny.Set
				_ = _215_v29
				_215_v29 = _dafny.SetOf((_214_v28).F18())
				var _216_v30 _dafny.Map
				_ = _216_v30
				_216_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v29, _214_v28)
				var _217_v31 _dafny.Array
				_ = _217_v31
				var _nwElement0_6 *C0 = _214_v28
				_ = _nwElement0_6
				var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(22))
				_ = _nw34
				(_nw34).ArraySet1(_nwElement0_6, 0)
				(_nw34).ArraySet1(_214_v28, 1)
				(_nw34).ArraySet1(_214_v28, 2)
				(_nw34).ArraySet1(_214_v28, 3)
				(_nw34).ArraySet1(_214_v28, 4)
				(_nw34).ArraySet1(_214_v28, 5)
				(_nw34).ArraySet1(_214_v28, 6)
				(_nw34).ArraySet1(_214_v28, 7)
				(_nw34).ArraySet1((func() *C0 {
					if (_216_v30).Contains(_dafny.SetOf(_175_v1, (_214_v28).F18())) {
						return (_216_v30).Get(_dafny.SetOf(_175_v1, (_214_v28).F18())).(*C0)
					}
					return _214_v28
				})(), 8)
				(_nw34).ArraySet1(_214_v28, 9)
				(_nw34).ArraySet1(_214_v28, 10)
				(_nw34).ArraySet1(_214_v28, 11)
				(_nw34).ArraySet1(_214_v28, 12)
				(_nw34).ArraySet1(_214_v28, 13)
				(_nw34).ArraySet1(_214_v28, 14)
				(_nw34).ArraySet1(_214_v28, 15)
				(_nw34).ArraySet1(_214_v28, 16)
				(_nw34).ArraySet1(_214_v28, 17)
				(_nw34).ArraySet1(_214_v28, 18)
				(_nw34).ArraySet1(_214_v28, 19)
				(_nw34).ArraySet1(_214_v28, 20)
				(_nw34).ArraySet1(_214_v28, 21)
				_217_v31 = _nw34
				var _218_v32 _dafny.Map
				_ = _218_v32
				_218_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, _217_v31)
				var _219_v33 _dafny.Map
				_ = _219_v33
				_219_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC23_(_218_v32), _208_v24.F12())
				var _220_v34 _dafny.Map
				_ = _220_v34
				_220_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v26.F15, _218_v32)
				var _221_v35 _dafny.Map
				_ = _221_v35
				_221_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v26.F15, _175_v1)
				var _222_v36 D9
				_ = _222_v36
				_222_v36 = Companion_D9_.Create_DC24_(_206_v22, (_214_v28).F17(), _175_v1)
				var _223_v37 _dafny.Map
				_ = _223_v37
				_223_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_214_v28).Fm2(_dafny.MultiSetOf((_208_v24.F11()).Cardinality()), _175_v1, _221_v35, false, _185_globalState), (_222_v36).Dtor_cf57())
				var _224_v38 _dafny.MultiSet
				_ = _224_v38
				_224_v38 = _dafny.MultiSetOf(Companion_Default___.Fm16((_214_v28).F18(), _185_globalState), _223_v37)
				var _225_v39 _dafny.MultiSet
				_ = _225_v39
				_225_v39 = _dafny.MultiSetOf(_177_v3, _177_v3, _177_v3, _177_v3, _177_v3)
				var _226_v40 _dafny.MultiSet
				_ = _226_v40
				_226_v40 = _dafny.MultiSetOf((_214_v28).F18(), _dafny.IntOfInt64(80), (_214_v28).F18(), _175_v1, (func() _dafny.Int {
					if (_225_v39).Contains(_177_v3) {
						return (_225_v39).Multiplicity(_177_v3)
					}
					return (_214_v28).F18()
				})())
				var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_209_v25), 0))
				_ = _index22
				var _rhs18 bool = (func() bool {
					if (_219_v33).Contains(Companion_D9_.Create_DC23_((func() _dafny.Map {
						if (_220_v34).Contains((_214_v28).F17()) {
							return (_220_v34).Get((_214_v28).F17()).(_dafny.Map)
						}
						return _218_v32
					})())) {
						return (_219_v33).Get(Companion_D9_.Create_DC23_((func() _dafny.Map {
							if (_220_v34).Contains((_214_v28).F17()) {
								return (_220_v34).Get((_214_v28).F17()).(_dafny.Map)
							}
							return _218_v32
						})())).(bool)
					}
					return (_183_v6) == ((_214_v28).F17())
				})()
				_ = _rhs18
				var _rhs19 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_206_v22, _dafny.UnicodeSeqOfUtf8Bytes("svi")), _206_v22)
				_ = _rhs19
				var _rhs20 _dafny.MultiSet = Companion_Default___.Fm22((func() _dafny.Int {
					if (_224_v38).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_214_v28).F17(), (_214_v28).F17())) {
						return (_224_v38).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_214_v28).F17(), (_214_v28).F17()))
					}
					return (_214_v28).F18()
				})(), _208_v24.F12(), _185_globalState)
				_ = _rhs20
				var _rhs21 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_212_v26), _213_v27), _213_v27)
				_ = _rhs21
				var _rhs22 bool = !(!((_208_v24).Fm2(_226_v40, (_214_v28).F18(), Companion_Default___.Fm23((_214_v28).F18(), _185_globalState), _212_v26.F15, _185_globalState)))
				_ = _rhs22
				var _lhs9 *GlobalState = _185_globalState
				_ = _lhs9
				var _lhs10 *GlobalState = _185_globalState
				_ = _lhs10
				var _lhs11 _dafny.Array = _209_v25
				_ = _lhs11
				var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_209_v25), 0))
				_ = _lhs12
				var _lhs13 *GlobalState = _185_globalState
				_ = _lhs13
				_lhs9.F3 = _rhs18
				_lhs10.F7 = _rhs19
				(_lhs11).ArraySet1(_rhs20, (_lhs12).Int())
				_213_v27 = _rhs21
				_lhs13.F7 = _rhs22
				var _227_v41 D6
				_ = _227_v41
				_227_v41 = Companion_D6_.Create_DC16_(_175_v1, (_225_v39).IsProperSubsetOf((_225_v39).Update(_177_v3, Companion_Default___.Abs((func() _dafny.Int {
					if (_208_v24.F11()).Contains(_212_v26.F15) {
						return (_208_v24.F11()).Multiplicity(_212_v26.F15)
					}
					return _175_v1
				})()))), !(_183_v6))
				var _source2 D6 = _227_v41
				_ = _source2
				if _source2.Is_DC16() {
					var _228___mcc_h0 _dafny.Int = _source2.Get_().(D6_DC16).Cf36
					_ = _228___mcc_h0
					var _229___mcc_h1 bool = _source2.Get_().(D6_DC16).Cf37
					_ = _229___mcc_h1
					var _230___mcc_h2 bool = _source2.Get_().(D6_DC16).Cf38
					_ = _230___mcc_h2
					var _231_cf38 bool = _230___mcc_h2
					_ = _231_cf38
					var _232_cf37 bool = _229___mcc_h1
					_ = _232_cf37
					var _233_cf36 _dafny.Int = _228___mcc_h0
					_ = _233_cf36
					_175_v1 = (_214_v28).F18()
					var _234_v42 _dafny.Map
					_ = _234_v42
					_234_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_232_cf37, _223_v37)
					var _235_v43 _dafny.Map
					_ = _235_v43
					_235_v43 = _223_v37
					var _rhs23 T0 = _208_v24
					_ = _rhs23
					var _rhs24 _dafny.Map = ((func() _dafny.Map {
						if Companion_Default___.Fm0(_177_v3, _185_globalState) {
							return (((func() _dafny.Map {
								if (_234_v42).Contains((_214_v28).F17()) {
									return (_234_v42).Get((_214_v28).F17()).(_dafny.Map)
								}
								return _223_v37
							})()).Update(_212_v26.F15, _208_v24.F12())).Update(_232_cf37, _208_v24.F12())
						}
						return _223_v37
					})()).Merge((_235_v43))
					_ = _rhs24
					_208_v24 = _rhs23
					_223_v37 = _rhs24
					var _236_v44 _dafny.Array
					_ = _236_v44
					var _nwElement0_7 bool = (_214_v28).F17()
					_ = _nwElement0_7
					var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(19))
					_ = _nw35
					(_nw35).ArraySet1(_nwElement0_7, 0)
					(_nw35).ArraySet1((_214_v28).F17(), 1)
					(_nw35).ArraySet1(_231_cf38, 2)
					(_nw35).ArraySet1(_208_v24.F12(), 3)
					(_nw35).ArraySet1(_212_v26.F15, 4)
					(_nw35).ArraySet1(_212_v26.F15, 5)
					(_nw35).ArraySet1((_214_v28).F17(), 6)
					(_nw35).ArraySet1(_232_cf37, 7)
					(_nw35).ArraySet1(_212_v26.F15, 8)
					(_nw35).ArraySet1(_212_v26.F15, 9)
					(_nw35).ArraySet1((_214_v28).F17(), 10)
					(_nw35).ArraySet1((_214_v28).F17(), 11)
					(_nw35).ArraySet1(_208_v24.F12(), 12)
					(_nw35).ArraySet1((_214_v28).F17(), 13)
					(_nw35).ArraySet1(!(_212_v26.F15), 14)
					(_nw35).ArraySet1(_208_v24.F12(), 15)
					(_nw35).ArraySet1((_214_v28).F17(), 16)
					(_nw35).ArraySet1(_231_cf38, 17)
					(_nw35).ArraySet1((_212_v26).Fm2(_dafny.MultiSetOf((_214_v28).F18(), _175_v1), (_214_v28).F18(), _221_v35, _212_v26.F15, _185_globalState), 18)
					_236_v44 = _nw35
					var _237_v45 _dafny.Sequence
					_ = _237_v45
					_237_v45 = _dafny.SeqOf(_236_v44, _236_v44, _236_v44, _236_v44, _236_v44)
					var _238_v46 _dafny.Set
					_ = _238_v46
					_238_v46 = _dafny.SetOf(_236_v44, _236_v44, (_237_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(849), _dafny.IntOfUint32((_237_v45).Cardinality()))).Uint32()).(_dafny.Array))
					_238_v46 = (_238_v46).Intersection(_238_v46)
					var _239_v47 D1
					_ = _239_v47
					_239_v47 = Companion_D1_.Create_DC2_(_231_cf38, _212_v26.F15, (_214_v28).F17(), _177_v3, _177_v3)
					var _240_v48 _dafny.Map
					_ = _240_v48
					_240_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_214_v28).F18(), Companion_Default___.Fm14(_239_v47, (_208_v24.F11()).Cardinality(), _185_globalState))
					var _241_v49 D1
					_ = _241_v49
					_241_v49 = Companion_D1_.Create_DC2_(_231_cf38, (func() bool {
						if (_223_v37).Contains(_208_v24.F12()) {
							return (_223_v37).Get(_208_v24.F12()).(bool)
						}
						return true
					})(), false, _177_v3, (func() _dafny.CodePoint {
						if (_240_v48).Contains((_214_v28).F18()) {
							return (_240_v48).Get((_214_v28).F18()).(_dafny.CodePoint)
						}
						return _177_v3
					})())
					_177_v3 = Companion_Default___.Fm14(_241_v49, _175_v1, _185_globalState)
				} else {
					var _242___mcc_h3 _dafny.MultiSet = _source2.Get_().(D6_DC15).Cf35
					_ = _242___mcc_h3
					var _243_cf35 _dafny.MultiSet = _242___mcc_h3
					_ = _243_cf35
					var _244_v50 bool
					_ = _244_v50
					var _245_v51 _dafny.Int
					_ = _245_v51
					var _out0 bool
					_ = _out0
					var _out1 _dafny.Int
					_ = _out1
					_out0, _out1 = (_212_v26).M5(Companion_Default___.SafeModuloInt((_214_v28).F18(), (_214_v28).F18()), _206_v22, _185_globalState)
					_244_v50 = _out0
					_245_v51 = _out1
					_206_v22 = _dafny.Companion_Sequence_.Concatenate(_206_v22, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("alxsdoim"), _206_v22))
					_245_v51 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_184_v7, _dafny.SeqOf(_183_v6, _212_v26.F15)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_208_v24.F12(), _208_v24.F12(), _212_v26.F15), _184_v7))).Cardinality())
					(_185_globalState).F3 = ((_214_v28).F18()).Cmp(_dafny.IntOfUint32((_206_v22).Cardinality())) != 0
				}
				var _246_v52 bool
				_ = _246_v52
				var _247_v53 _dafny.Int
				_ = _247_v53
				var _out2 bool
				_ = _out2
				var _out3 _dafny.Int
				_ = _out3
				_out2, _out3 = (_212_v26).M5(Companion_Default___.Fm1(_212_v26.F15, _175_v1, (_214_v28).F18(), _185_globalState), _206_v22, _185_globalState)
				_246_v52 = _out2
				_247_v53 = _out3
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	_176_v2 = _dafny.Companion_Sequence_.Concatenate(_176_v2, _176_v2)
	(_185_globalState).F3 = _183_v6
	var _248_v54 _dafny.Sequence
	_ = _248_v54
	_248_v54 = _dafny.UnicodeSeqOfUtf8Bytes("eyxck")
	var _249_v55 _dafny.Set
	_ = _249_v55
	_249_v55 = _dafny.SetOf(_175_v1)
	var _250_v56 D7
	_ = _250_v56
	_250_v56 = Companion_D7_.Create_DC18_(_dafny.IntOfUint32((_184_v7).Cardinality()), _248_v54, _183_v6, _183_v6, _249_v55)
	var _251_v57 _dafny.Sequence
	_ = _251_v57
	_251_v57 = _dafny.SeqOf((_250_v56).Dtor_cf41())
	if (_175_v1).Cmp(_dafny.IntOfUint32(((_251_v57).Select((Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_251_v57).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) <= 0 {
		var _252_v58 _dafny.MultiSet
		_ = _252_v58
		_252_v58 = _dafny.MultiSetOf(_183_v6)
		var _253_v59 T0
		_ = _253_v59
		var _nw36 *C2 = New_C2_()
		_ = _nw36
		_nw36.Ctor__(_183_v6, _252_v58, _183_v6)
		_253_v59 = _nw36
		var _254_v60 _dafny.MultiSet
		_ = _254_v60
		_254_v60 = _dafny.MultiSetOf(_253_v59, _253_v59, _253_v59, _253_v59, _253_v59)
		var _255_v61 _dafny.Map
		_ = _255_v61
		_255_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _183_v6)
		var _256_v62 _dafny.Map
		_ = _256_v62
		_256_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_175_v1), _175_v1)
		var _257_v63 _dafny.MultiSet
		_ = _257_v63
		_257_v63 = _dafny.MultiSetOf(_175_v1, (_256_v62).Cardinality(), _175_v1)
		var _258_v64 *C1
		_ = _258_v64
		var _nw37 *C1 = New_C1_()
		_ = _nw37
		_nw37.Ctor__((_dafny.SetOf(_175_v1, _175_v1, _dafny.IntOfInt64(466), (_254_v60).Cardinality(), (_255_v61).Cardinality())).Cardinality(), _dafny.MultiSetFromSeq(Companion_Default___.Fm10(_253_v59.F12(), _253_v59.F12(), _183_v6, _185_globalState)), ((_257_v63).Cardinality()).Cmp(Companion_Default___.Fm1(_183_v6, (_249_v55).Cardinality(), _175_v1, _185_globalState)) < 0)
		_258_v64 = _nw37
		var _259_v66 _dafny.Array
		_ = _259_v66
		var _nwElement0_8 _dafny.Sequence = _248_v54
		_ = _nwElement0_8
		var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
		_ = _nw38
		(_nw38).ArraySet1(_nwElement0_8, 0)
		(_nw38).ArraySet1(_248_v54, 1)
		(_nw38).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_260_v54 _dafny.Sequence, _261_v64 *C1, _262_v1 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
			return func(_263_i6 _dafny.Int) _dafny.CodePoint {
				return (_260_v54).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_dafny.SeqOf(_261_v64.F16)).Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _264_v65 _dafny.Int
						_264_v65 = interface{}(_compr_12).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_261_v64.F16), _264_v65) {
							_coll12.Add((_264_v65).Minus(_262_v1), _dafny.IntOfInt64(-244))
						}
					}
					return _coll12.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_260_v54).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_248_v54, _258_v64, _175_v1))), 2)
		(_nw38).ArraySet1(_248_v54, 3)
		(_nw38).ArraySet1(_248_v54, 4)
		(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("r"), 5)
		(_nw38).ArraySet1(_dafny.Companion_Sequence_.Update(_248_v54, (Companion_Default___.SafeIndex(_258_v64.F16, _dafny.IntOfUint32((_248_v54).Cardinality()))).Uint32(), _177_v3), 6)
		(_nw38).ArraySet1(_248_v54, 7)
		(_nw38).ArraySet1(_248_v54, 8)
		(_nw38).ArraySet1(_248_v54, 9)
		(_nw38).ArraySet1(_248_v54, 10)
		(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("foxyq"), 11)
		(_nw38).ArraySet1(_248_v54, 12)
		(_nw38).ArraySet1(_248_v54, 13)
		_259_v66 = _nw38
		var _265_v67 D5
		_ = _265_v67
		_265_v67 = Companion_D5_.Create_DC13_(_258_v64.F16, _183_v6, (_dafny.Zero).Minus(_258_v64.F16), _259_v66, _253_v59.F12())
		_175_v1 = (_258_v64.F16).Times((_265_v67).Dtor_cf30())
		var _rhs25 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qtkfwgn")).Cardinality())).Minus((func() _dafny.Int {
			if (_257_v63).Contains((_255_v61).Cardinality()) {
				return (_257_v63).Multiplicity((_255_v61).Cardinality())
			}
			return _175_v1
		})())
		_ = _rhs25
		var _rhs26 _dafny.Int = _175_v1
		_ = _rhs26
		var _lhs14 *C1 = _258_v64
		_ = _lhs14
		_lhs14.F16 = _rhs25
		_175_v1 = _rhs26
		var _266_v68 D9
		_ = _266_v68
		_266_v68 = Companion_D9_.Create_DC25_(_dafny.Companion_Sequence_.Concatenate(_248_v54, _dafny.UnicodeSeqOfUtf8Bytes("ak")), _253_v59.F12(), _258_v64.F16)
		var _source3 D9 = _266_v68
		_ = _source3
		if _source3.Is_DC24() {
			var _267___mcc_h4 _dafny.Sequence = _source3.Get_().(D9_DC24).Cf56
			_ = _267___mcc_h4
			var _268___mcc_h5 bool = _source3.Get_().(D9_DC24).Cf57
			_ = _268___mcc_h5
			var _269___mcc_h6 _dafny.Int = _source3.Get_().(D9_DC24).Cf58
			_ = _269___mcc_h6
			var _270_cf58 _dafny.Int = _269___mcc_h6
			_ = _270_cf58
			var _271_cf57 bool = _268___mcc_h5
			_ = _271_cf57
			var _272_cf56 _dafny.Sequence = _267___mcc_h4
			_ = _272_cf56
			var _273_v69 _dafny.Map
			_ = _273_v69
			_273_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v59.F12(), false)
			_273_v69 = (_273_v69).Update(((_dafny.SetOf(_271_cf57, _253_v59.F12())).Cardinality()).Cmp((_dafny.Zero).Minus(_dafny.IntOfInt64(-496))) > 0, (func() bool {
				if _253_v59.F12() {
					return _253_v59.F12()
				}
				return _253_v59.F12()
			})())
			var _274_v70 _dafny.Map
			_ = _274_v70
			_274_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v59.F12(), _258_v64.F16)
			var _275_v71 _dafny.Map
			_ = _275_v71
			_275_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v64.F16, _274_v70)
			_275_v71 = (_275_v71).Update(_dafny.IntOfInt64(12), _274_v70)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_180_v5), 0))
			_ = _index23
			(_180_v5).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_272_cf56, _272_cf56)).Cardinality()), (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_180_v5), 0))
			_ = _index24
			(_180_v5).ArraySet1(_258_v64.F16, (_index24).Int())
			_177_v3 = _177_v3
		} else if _source3.Is_DC25() {
			var _276___mcc_h7 _dafny.Sequence = _source3.Get_().(D9_DC25).Cf59
			_ = _276___mcc_h7
			var _277___mcc_h8 bool = _source3.Get_().(D9_DC25).Cf60
			_ = _277___mcc_h8
			var _278___mcc_h9 _dafny.Int = _source3.Get_().(D9_DC25).Cf61
			_ = _278___mcc_h9
			var _279_cf61 _dafny.Int = _278___mcc_h9
			_ = _279_cf61
			var _280_cf60 bool = _277___mcc_h8
			_ = _280_cf60
			var _281_cf59 _dafny.Sequence = _276___mcc_h7
			_ = _281_cf59
			var _282_v72 _dafny.Array
			_ = _282_v72
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw39
			_282_v72 = _nw39
			var _283_v73 _dafny.Map
			_ = _283_v73
			_283_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_255_v61).Contains(_258_v64.F16) {
					return (_255_v61).Get(_258_v64.F16).(bool)
				}
				return _253_v59.F12()
			})(), _282_v72)
			_283_v73 = (_283_v73).Update(_253_v59.F12(), _282_v72)
			var _284_v74 _dafny.Map
			_ = _284_v74
			_284_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v59.F12(), _175_v1)
			var _285_v75 _dafny.Map
			_ = _285_v75
			_285_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_256_v62).Contains(_258_v64.F16) {
					return (_256_v62).Get(_258_v64.F16).(_dafny.Int)
				}
				return (_284_v74).Cardinality()
			})(), _dafny.SeqOf(_280_cf60))
			var _286_v76 *C3
			_ = _286_v76
			var _nw40 *C3 = New_C3_()
			_ = _nw40
			_nw40.Ctor__(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_285_v75).Contains(_dafny.IntOfInt64(679)) {
					return (_285_v75).Get(_dafny.IntOfInt64(679)).(_dafny.Sequence)
				}
				return _184_v7
			})()).Cardinality()), _183_v6, Companion_Default___.Fm22(_258_v64.F16, _253_v59.F12(), _185_globalState), _253_v59.F12())
			_286_v76 = _nw40
			var _287_v77 _dafny.Map
			_ = _287_v77
			_287_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_286_v76, (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(283))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_288_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_289_i7 _dafny.Int) _dafny.CodePoint {
					return _288_v3
				}
			})(_177_v3)))).Cardinality())).Minus(_258_v64.F16))
			(_185_globalState).F9 = (_287_v77).Cardinality()
			(_185_globalState).F8 = _253_v59.F12()
			var _290_v78 D6
			_ = _290_v78
			_290_v78 = Companion_D6_.Create_DC16_(_258_v64.F16, _280_cf60, _183_v6)
			(_253_v59).F12_set_((_290_v78).Dtor_cf38())
		} else if _source3.Is_DC26() {
			var _291___mcc_h10 _dafny.Int = _source3.Get_().(D9_DC26).Cf62
			_ = _291___mcc_h10
			var _292_cf62 _dafny.Int = _291___mcc_h10
			_ = _292_cf62
			(_185_globalState).F8 = _253_v59.F12()
			var _293_v79 *C0
			_ = _293_v79
			var _nw41 *C0 = New_C0_()
			_ = _nw41
			_nw41.Ctor__(!(_253_v59.F12()), Companion_Default___.SafeModuloInt(_258_v64.F16, _258_v64.F16), _252_v58, !(_183_v6))
			_293_v79 = _nw41
			_292_cf62 = _dafny.IntOfInt64(-697)
			var _294_v80 _dafny.Array
			_ = _294_v80
			var _nw42 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
			_ = _nw42
			_294_v80 = _nw42
			var _295_v81 _dafny.Sequence
			_ = _295_v81
			_295_v81 = _dafny.SeqOf(_253_v59.F11())
			var _296_v82 *C3
			_ = _296_v82
			var _nw43 *C3 = New_C3_()
			_ = _nw43
			_nw43.Ctor__(_258_v64.F16, _183_v6, (_295_v81).Select((Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_295_v81).Cardinality()))).Uint32()).(_dafny.MultiSet), _253_v59.F12())
			_296_v82 = _nw43
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_294_v80), 0))
			_ = _index25
			(_294_v80).ArraySet1(_296_v82, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_294_v80), 0))
			_ = _index26
			(_294_v80).ArraySet1(_296_v82, (_index26).Int())
		} else {
			var _297___mcc_h11 _dafny.Map = _source3.Get_().(D9_DC23).Cf55
			_ = _297___mcc_h11
			var _298_cf55 _dafny.Map = _297___mcc_h11
			_ = _298_cf55
			var _299_v83 *C4
			_ = _299_v83
			var _nw44 *C4 = New_C4_()
			_ = _nw44
			_nw44.Ctor__(_253_v59.F11(), _253_v59.F12())
			_299_v83 = _nw44
			var _300_v84 _dafny.Array
			_ = _300_v84
			var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw45
			_300_v84 = _nw45
			var _301_v85 _dafny.Map
			_ = _301_v85
			_301_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, _300_v84)
			var _302_v86 _dafny.Sequence
			_ = _302_v86
			_302_v86 = _dafny.SeqOf(_252_v58)
			var _303_v87 _dafny.Map
			_ = _303_v87
			_303_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_301_v85).Contains(_183_v6) {
					return (_301_v85).Get(_183_v6).(_dafny.Array)
				}
				return _300_v84
			})(), _302_v86)
			_303_v87 = (_303_v87).Update(_300_v84, _302_v86)
			var _304_v88 D8
			_ = _304_v88
			_304_v88 = Companion_D8_.Create_DC20_(_253_v59.F12(), _253_v59.F12(), false)
			_304_v88 = Companion_D8_.Create_DC20_(_253_v59.F12(), _183_v6, (_299_v83).Fm4(_258_v64.F16, _258_v64.F16, _185_globalState))
			var _305_v89 _dafny.Set
			_ = _305_v89
			_305_v89 = _dafny.SetOf(_253_v59.F12(), true, _253_v59.F12(), _183_v6, !(_253_v59.F12()))
			var _306_v90 _dafny.Array
			_ = _306_v90
			var _307_v91 _dafny.Array
			_ = _307_v91
			var _308_v92 _dafny.Int
			_ = _308_v92
			var _out4 _dafny.Array
			_ = _out4
			var _out5 _dafny.Array
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			_out4, _out5, _out6 = (_299_v83).M1(_305_v89, _185_globalState)
			_306_v90 = _out4
			_307_v91 = _out5
			_308_v92 = _out6
		}
		var _rhs27 _dafny.Int = _258_v64.F16
		_ = _rhs27
		var _rhs28 _dafny.Int = (func() _dafny.Int {
			if (_175_v1).Cmp(_258_v64.F16) >= 0 {
				return _258_v64.F16
			}
			return _175_v1
		})()
		_ = _rhs28
		var _rhs29 _dafny.Int = _175_v1
		_ = _rhs29
		var _rhs30 bool = _183_v6
		_ = _rhs30
		var _rhs31 bool = _183_v6
		_ = _rhs31
		var _lhs15 *C1 = _258_v64
		_ = _lhs15
		var _lhs16 *GlobalState = _185_globalState
		_ = _lhs16
		var _lhs17 *GlobalState = _185_globalState
		_ = _lhs17
		var _lhs18 *GlobalState = _185_globalState
		_ = _lhs18
		var _lhs19 *GlobalState = _185_globalState
		_ = _lhs19
		_lhs15.F16 = _rhs27
		_lhs16.F9 = _rhs28
		_lhs17.F9 = _rhs29
		_lhs18.F3 = _rhs30
		_lhs19.F8 = _rhs31
	} else {
		var _309_v93 _dafny.MultiSet
		_ = _309_v93
		_309_v93 = _dafny.MultiSetOf(!(_183_v6), false, _183_v6)
		var _310_v94 _dafny.Map
		_ = _310_v94
		_310_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Equals(_309_v93), _176_v2)
		_310_v94 = (_310_v94).Update((_184_v7).Select((Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_184_v7).Cardinality()))).Uint32()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(226))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_311_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_312_i8 _dafny.Int) _dafny.Int {
				return _311_v1
			}
		})(_175_v1))))
		(_185_globalState).F8 = (_175_v1).Cmp(_175_v1) > 0
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))
		_ = _index27
		(_180_v5).ArraySet1((_175_v1).Times(_175_v1), (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))
		_ = _index28
		var _rhs32 _dafny.Int = (func() _dafny.Int {
			if _183_v6 {
				return _175_v1
			}
			return _175_v1
		})()
		_ = _rhs32
		var _rhs33 _dafny.Int = (_176_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.IntOfUint32((_176_v2).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs33
		var _lhs20 *GlobalState = _185_globalState
		_ = _lhs20
		var _lhs21 _dafny.Array = _180_v5
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))
		_ = _lhs22
		_lhs20.F9 = _rhs32
		(_lhs21).ArraySet1(_rhs33, (_lhs22).Int())
		Companion_Default___.M0(_185_globalState)
		var _313_v95 _dafny.Map
		_ = _313_v95
		_313_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int))
		var _314_v96 _dafny.Map
		_ = _314_v96
		_314_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), (_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int))
		(_185_globalState).F9 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
			if (_313_v95).Contains(!(true)) {
				return (_313_v95).Get(!(true)).(_dafny.Int)
			}
			return Companion_Default___.Fm1(_183_v6, (_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_314_v96).Cardinality()), _185_globalState)
		})(), (_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int))
	}
	(_185_globalState).F3 = (_dafny.IntOfUint32(((func() _dafny.Sequence {
		if _183_v6 {
			return _248_v54
		}
		return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("gul"), (Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gul")).Cardinality()))).Uint32(), _dafny.CodePoint('d'))
	})()).Cardinality())).Cmp(_175_v1) == 0
	if _183_v6 {
		var _315_v97 _dafny.Set
		_ = _315_v97
		_315_v97 = _dafny.SetOf(_183_v6, _183_v6)
		if (_315_v97).IsProperSubsetOf(_315_v97) {
			var _316_v98 _dafny.Map
			_ = _316_v98
			_316_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, Companion_Default___.Fm0(_177_v3, _185_globalState))
			var _317_v99 _dafny.Map
			_ = _317_v99
			_317_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _316_v98)
			var _318_v100 D10
			_ = _318_v100
			_318_v100 = Companion_D10_.Create_DC29_(_183_v6)
			var _pat_let_tv11 = _183_v6
			_ = _pat_let_tv11
			var _319_v101 _dafny.Set
			_ = _319_v101
			_319_v101 = _dafny.SetOf(_315_v97, Companion_Default___.Fm18(((_317_v99).Update(_dafny.SeqOf(_dafny.IntOfUint32((_251_v57).Cardinality())), _316_v98)).Cardinality(), (func(_pat_let10_0 D10) D10 {
				return func(_320_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let11_0 bool) D10 {
						return func(_321_dt__update_hcf68_h0 bool) D10 {
							return Companion_D10_.Create_DC29_(_321_dt__update_hcf68_h0)
						}(_pat_let11_0)
					}(_pat_let_tv11)
				}(_pat_let10_0)
			}(_318_v100)).Dtor_cf68(), _185_globalState), _dafny.SetOf(_183_v6, _183_v6))
			var _rhs34 bool = !(Companion_Default___.Fm0(_177_v3, _185_globalState))
			_ = _rhs34
			var _rhs35 _dafny.Int = (_175_v1).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vomrh")).Cardinality()))
			_ = _rhs35
			var _rhs36 _dafny.Set = _319_v101
			_ = _rhs36
			var _rhs37 bool = _183_v6
			_ = _rhs37
			var _rhs38 bool = (_175_v1).Cmp((_dafny.IntOfInt64(-312)).Times(_175_v1)) > 0
			_ = _rhs38
			var _lhs23 *GlobalState = _185_globalState
			_ = _lhs23
			var _lhs24 *GlobalState = _185_globalState
			_ = _lhs24
			var _lhs25 *GlobalState = _185_globalState
			_ = _lhs25
			_lhs23.F3 = _rhs34
			_175_v1 = _rhs35
			_319_v101 = _rhs36
			_lhs24.F7 = _rhs37
			_lhs25.F7 = _rhs38
			var _322_v102 _dafny.Array
			_ = _322_v102
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_6
			var _nw46 _dafny.Array
			_ = _nw46
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw46 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Set = (func(_323_v55 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_324_i9 _dafny.Int) _dafny.Set {
						return _323_v55
					}
				})(_249_v55)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw46 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw46).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw46).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_322_v102 = _nw46
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_322_v102), 0))
			_ = _index29
			(_322_v102).ArraySet1(_249_v55, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_322_v102), 0))
			_ = _index30
			(_322_v102).ArraySet1(_249_v55, (_index30).Int())
			var _325_v103 _dafny.Map
			_ = _325_v103
			_325_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _183_v6)
			var _326_v104 _dafny.Map
			_ = _326_v104
			_326_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v103, _175_v1)
			(_185_globalState).F9 = ((_326_v104).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _183_v6), (_dafny.Zero).Minus(((_322_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(838), _dafny.ArrayLen((_322_v102), 0))).Int()).(_dafny.Set)).Cardinality())))).Cardinality()
			_175_v1 = _175_v1
			(_185_globalState).F7 = _183_v6
		} else {
			var _327_v105 _dafny.Map
			_ = _327_v105
			_327_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v97, _175_v1)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _index31
			(_180_v5).ArraySet1((func() _dafny.Int {
				if (_327_v105).Contains(_315_v97) {
					return (_327_v105).Get(_315_v97).(_dafny.Int)
				}
				return _175_v1
			})(), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_180_v5), 0))
			_ = _index32
			(_180_v5).ArraySet1(_175_v1, (_index32).Int())
			var _328_v106 *C4
			_ = _328_v106
			var _nw47 *C4 = New_C4_()
			_ = _nw47
			_nw47.Ctor__(_dafny.MultiSetOf(_183_v6, _183_v6), _183_v6)
			_328_v106 = _nw47
			var _329_v107 _dafny.MultiSet
			_ = _329_v107
			_329_v107 = _dafny.MultiSetOf(_328_v106, _328_v106)
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _index33
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_180_v5), 0))
			_ = _index34
			var _rhs39 bool = _183_v6
			_ = _rhs39
			var _rhs40 _dafny.Int = _175_v1
			_ = _rhs40
			var _rhs41 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(574), _175_v1)
			_ = _rhs41
			var _rhs42 _dafny.Int = Companion_Default___.SafeDivisionInt(_175_v1, _dafny.IntOfInt64(-498))
			_ = _rhs42
			var _rhs43 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_329_v107, _183_v6)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_329_v107, _183_v6)) {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}((func(_330_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_331_i10 _dafny.Int) _dafny.Int {
							return _330_v1
						}
					})(_175_v1)))
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(817))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_332_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_333_i11 _dafny.Int) _dafny.Int {
						return _332_v1
					}
				})(_175_v1)))
			})()).Cardinality())
			_ = _rhs43
			var _lhs26 *GlobalState = _185_globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _180_v5
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs28
			var _lhs29 _dafny.Array = _180_v5
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = _185_globalState
			_ = _lhs31
			var _lhs32 *GlobalState = _185_globalState
			_ = _lhs32
			_lhs26.F3 = _rhs39
			(_lhs27).ArraySet1(_rhs40, (_lhs28).Int())
			(_lhs29).ArraySet1(_rhs41, (_lhs30).Int())
			_lhs31.F9 = _rhs42
			_lhs32.F9 = _rhs43
			(_185_globalState).F9 = _dafny.IntOfInt64(760)
			var _334_v108 _dafny.MultiSet
			_ = _334_v108
			_334_v108 = _dafny.MultiSetOf(_183_v6, _183_v6, false, _183_v6)
			var _335_v109 _dafny.Map
			_ = _335_v109
			_335_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _334_v108)
			var _336_v110 D1
			_ = _336_v110
			_336_v110 = Companion_D1_.Create_DC2_(true, _183_v6, _183_v6, _177_v3, _177_v3)
			var _pat_let_tv12 = _177_v3
			_ = _pat_let_tv12
			var _337_v111 *C0
			_ = _337_v111
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__(_183_v6, _175_v1, (func() _dafny.MultiSet {
				if (_335_v109).Contains((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int)) {
					return (_335_v109).Get((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int)).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_183_v6, _183_v6, (func(_pat_let12_0 D1) D1 {
					return func(_338_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let13_0 _dafny.CodePoint) D1 {
							return func(_339_dt__update_hcf6_h0 _dafny.CodePoint) D1 {
								return func(_pat_let14_0 bool) D1 {
									return func(_340_dt__update_hcf2_h0 bool) D1 {
										return Companion_D1_.Create_DC2_(_340_dt__update_hcf2_h0, (_338_dt__update__tmp_h1).Dtor_cf3(), (_338_dt__update__tmp_h1).Dtor_cf4(), (_338_dt__update__tmp_h1).Dtor_cf5(), _339_dt__update_hcf6_h0)
									}(_pat_let14_0)
								}(false)
							}(_pat_let13_0)
						}(_pat_let_tv12)
					}(_pat_let12_0)
				}(_336_v110)).Dtor_cf4())
			})(), _183_v6)
			_337_v111 = _nw48
			var _341_v112 _dafny.Map
			_ = _341_v112
			_341_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v111, _177_v3)
			_341_v112 = (_341_v112).Update(_337_v111, _dafny.CodePoint('d'))
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _index36
			var _rhs44 _dafny.Int = ((_337_v111).F18()).Times(((_337_v111).F18()).Times(_dafny.IntOfInt64(55)))
			_ = _rhs44
			var _rhs45 _dafny.Int = Companion_Default___.SafeDivisionInt((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_334_v108).Contains(_183_v6) {
					return (_334_v108).Multiplicity(_183_v6)
				}
				return _dafny.IntOfInt64(542)
			})())
			_ = _rhs45
			var _rhs46 _dafny.Int = (_dafny.Zero).Minus((_337_v111).F18())
			_ = _rhs46
			var _lhs33 _dafny.Array = _180_v5
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs34
			var _lhs35 _dafny.Array = _180_v5
			_ = _lhs35
			var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs36
			var _lhs37 *GlobalState = _185_globalState
			_ = _lhs37
			(_lhs33).ArraySet1(_rhs44, (_lhs34).Int())
			(_lhs35).ArraySet1(_rhs45, (_lhs36).Int())
			_lhs37.F9 = _rhs46
			var _342_v113 _dafny.Array
			_ = _342_v113
			var _343_v114 _dafny.Array
			_ = _343_v114
			var _344_v115 _dafny.Int
			_ = _344_v115
			var _out7 _dafny.Array
			_ = _out7
			var _out8 _dafny.Array
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			_out7, _out8, _out9 = (_328_v106).M1(Companion_Default___.Fm18(_dafny.IntOfInt64(269), (_337_v111).F17(), _185_globalState), _185_globalState)
			_342_v113 = _out7
			_343_v114 = _out8
			_344_v115 = _out9
		}
		(_185_globalState).F3 = (_175_v1).Cmp(Companion_Default___.SafeDivisionInt(_175_v1, _175_v1)) != 0
		var _345_v116 _dafny.Array
		_ = _345_v116
		var _nw49 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
		_ = _nw49
		_345_v116 = _nw49
		var _346_v117 _dafny.Sequence
		_ = _346_v117
		_346_v117 = _dafny.SeqOf(_345_v116)
		var _347_v118 _dafny.Array
		_ = _347_v118
		var _nwElement0_9 _dafny.Array = _345_v116
		_ = _nwElement0_9
		var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(14))
		_ = _nw50
		(_nw50).ArraySet1(_nwElement0_9, 0)
		(_nw50).ArraySet1(_345_v116, 1)
		(_nw50).ArraySet1(_345_v116, 2)
		(_nw50).ArraySet1(_345_v116, 3)
		(_nw50).ArraySet1(_345_v116, 4)
		(_nw50).ArraySet1((_346_v117).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_175_v1), _dafny.IntOfUint32((_346_v117).Cardinality()))).Uint32()).(_dafny.Array), 5)
		(_nw50).ArraySet1(_345_v116, 6)
		(_nw50).ArraySet1(_345_v116, 7)
		(_nw50).ArraySet1(_345_v116, 8)
		(_nw50).ArraySet1(_345_v116, 9)
		(_nw50).ArraySet1(_345_v116, 10)
		(_nw50).ArraySet1(_345_v116, 11)
		(_nw50).ArraySet1(_345_v116, 12)
		(_nw50).ArraySet1(_345_v116, 13)
		_347_v118 = _nw50
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_347_v118), 0))
		_ = _index37
		(_347_v118).ArraySet1((func() _dafny.Array {
			if _183_v6 {
				return _345_v116
			}
			return _345_v116
		})(), (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_347_v118), 0))
		_ = _index38
		(_347_v118).ArraySet1(_345_v116, (_index38).Int())
		var _348_v119 _dafny.Map
		_ = _348_v119
		_348_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v3, _175_v1)
		var _349_v120 _dafny.Set
		_ = _349_v120
		_349_v120 = _dafny.SetOf(_348_v119)
		var _350_v121 _dafny.Map
		_ = _350_v121
		_350_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_349_v120).Intersection(_349_v120), (_249_v55).Difference(_249_v55))
		_249_v55 = (func() _dafny.Set {
			if (_350_v121).Contains(_dafny.SetOf(_348_v119, _348_v119, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v3, _175_v1))) {
				return (_350_v121).Get(_dafny.SetOf(_348_v119, _348_v119, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v3, _175_v1))).(_dafny.Set)
			}
			return ((_250_v56).Dtor_cf44()).Difference(_249_v55)
		})()
		if Companion_Default___.Fm0(_177_v3, _185_globalState) {
			(_185_globalState).F9 = _dafny.IntOfInt64(678)
			var _351_v122 _dafny.Array
			_ = _351_v122
			var _nw51 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw51
			_351_v122 = _nw51
			_351_v122 = _351_v122
			var _352_v123 _dafny.Array
			_ = _352_v123
			var _nw52 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(29))
			_ = _nw52
			_352_v123 = _nw52
			var _353_v124 _dafny.Array
			_ = _353_v124
			var _nwElement0_10 _dafny.Array = _352_v123
			_ = _nwElement0_10
			var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(8))
			_ = _nw53
			(_nw53).ArraySet1(_nwElement0_10, 0)
			(_nw53).ArraySet1(_352_v123, 1)
			(_nw53).ArraySet1(_352_v123, 2)
			(_nw53).ArraySet1(_352_v123, 3)
			(_nw53).ArraySet1(_352_v123, 4)
			(_nw53).ArraySet1(_352_v123, 5)
			(_nw53).ArraySet1(_352_v123, 6)
			(_nw53).ArraySet1(_352_v123, 7)
			_353_v124 = _nw53
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_353_v124), 0))
			_ = _index39
			(_353_v124).ArraySet1(_352_v123, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_353_v124), 0))
			_ = _index40
			(_353_v124).ArraySet1(_352_v123, (_index40).Int())
			var _354_v125 _dafny.Map
			_ = _354_v125
			_354_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v3, _183_v6)
			_354_v125 = (_354_v125).Update(_177_v3, _183_v6)
			var _355_v126 _dafny.Array
			_ = _355_v126
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw54
			_355_v126 = _nw54
			var _356_v127 D8
			_ = _356_v127
			_356_v127 = Companion_D8_.Create_DC21_(_183_v6, _355_v126, _175_v1, _351_v122, _dafny.IntOfInt64(325))
			var _357_v128 _dafny.MultiSet
			_ = _357_v128
			_357_v128 = _dafny.MultiSetOf(!(Companion_Default___.Fm0(_177_v3, _185_globalState)) || (!((_356_v127).Dtor_cf49())), _183_v6, _183_v6, _183_v6)
			_357_v128 = ((Companion_Default___.Fm22(_175_v1, _183_v6, _185_globalState)).Union(_357_v128)).Intersection((_357_v128).Intersection(_357_v128))
		} else {
			(_185_globalState).F3 = _183_v6
			(_185_globalState).F9 = _dafny.IntOfInt64(-58)
			Companion_Default___.M0(_185_globalState)
			var _358_v129 _dafny.Array
			_ = _358_v129
			var _nw55 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw55
			_358_v129 = _nw55
			var _359_v130 _dafny.Map
			_ = _359_v130
			_359_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_358_v129, _183_v6)
			_359_v130 = (_359_v130).Update(_358_v129, _183_v6)
			_176_v2 = _dafny.SeqOf((_175_v1).Times(_175_v1), _dafny.IntOfInt64(975), _175_v1)
		}
	} else {
		(_185_globalState).F7 = _183_v6
		var _360_v131 _dafny.MultiSet
		_ = _360_v131
		_360_v131 = _dafny.MultiSetOf(_183_v6)
		var _361_v132 *C4
		_ = _361_v132
		var _nw56 *C4 = New_C4_()
		_ = _nw56
		_nw56.Ctor__(_360_v131, _183_v6)
		_361_v132 = _nw56
		var _362_v133 T0
		_ = _362_v133
		var _nw57 *C0 = New_C0_()
		_ = _nw57
		_nw57.Ctor__(_183_v6, _175_v1, (_360_v131).Update(_183_v6, Companion_Default___.Abs(_175_v1)), _183_v6)
		_362_v133 = _nw57
		_362_v133 = _362_v133
		var _363_v134 D8
		_ = _363_v134
		_363_v134 = Companion_D8_.Create_DC20_(_183_v6, _183_v6, _362_v133.F12())
		var _364_v135 *C0
		_ = _364_v135
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__(!(_362_v133.F12()), _175_v1, _362_v133.F11(), (_363_v134).Dtor_cf48())
		_364_v135 = _nw58
		var _365_v136 _dafny.Map
		_ = _365_v136
		_365_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, _175_v1)
		var _366_v137 _dafny.Set
		_ = _366_v137
		_366_v137 = _dafny.SetOf(_365_v136)
		var _367_v138 _dafny.Set
		_ = _367_v138
		_367_v138 = _dafny.SetOf(true, (_364_v135).F17())
		var _368_v139 *C0
		_ = _368_v139
		var _nw59 *C0 = New_C0_()
		_ = _nw59
		_nw59.Ctor__(_183_v6, (_366_v137).Cardinality(), _362_v133.F11(), ((_364_v135).F18()).Cmp((_367_v138).Cardinality()) >= 0)
		_368_v139 = _nw59
	}
	var _369_v140 _dafny.Map
	_ = _369_v140
	_369_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("muwivve"), (Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("muwivve")).Cardinality()))).Uint32(), _177_v3), (Companion_Default___.Fm26(_175_v1, !(false), _185_globalState)).Dtor_cf46())
	if (func() bool {
		if (_369_v140).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_370_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_371_i12 _dafny.Int) _dafny.CodePoint {
				return _370_v3
			}
		})(_177_v3)))) {
			return (_369_v140).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(820))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_372_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_373_i12 _dafny.Int) _dafny.CodePoint {
					return _372_v3
				}
			})(_177_v3)))).(bool)
		}
		return (_249_v55).IsProperSubsetOf(func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(567), _dafny.IntOfInt64(942))); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _374_v141 _dafny.Int
				_374_v141 = interface{}(_compr_13).(_dafny.Int)
				if ((_dafny.IntOfInt64(567)).Cmp(_374_v141) <= 0) && ((_374_v141).Cmp(_dafny.IntOfInt64(942)) < 0) {
					_coll13.Add((_374_v141).Plus(_175_v1))
				}
			}
			return _coll13.ToSet()
		}())
	})() {
		var _375_v142 _dafny.Array
		_ = _375_v142
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_7
		var _nw60 _dafny.Array
		_ = _nw60
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw60 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) bool = (func(_376_v6 bool) func(_dafny.Int) bool {
				return func(_377_i13 _dafny.Int) bool {
					return _376_v6
				}
			})(_183_v6)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw60 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw60).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw60).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_375_v142 = _nw60
		var _378_v143 _dafny.Sequence
		_ = _378_v143
		_378_v143 = _dafny.SeqOf(_375_v142)
		var _379_v144 _dafny.Array
		_ = _379_v144
		var _nwElement0_11 _dafny.Array = _375_v142
		_ = _nwElement0_11
		var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(23))
		_ = _nw61
		(_nw61).ArraySet1(_nwElement0_11, 0)
		(_nw61).ArraySet1(_375_v142, 1)
		(_nw61).ArraySet1(_375_v142, 2)
		(_nw61).ArraySet1((_378_v143).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_183_v6, Companion_Default___.Fm1(_183_v6, _175_v1, _175_v1, _185_globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(810))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_380_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_381_i14 _dafny.Int) _dafny.CodePoint {
				return _380_v3
			}
		})(_177_v3)))).Cardinality()), _185_globalState), _dafny.IntOfUint32((_378_v143).Cardinality()))).Uint32()).(_dafny.Array), 3)
		(_nw61).ArraySet1(_375_v142, 4)
		(_nw61).ArraySet1(_375_v142, 5)
		(_nw61).ArraySet1(_375_v142, 6)
		(_nw61).ArraySet1(_375_v142, 7)
		(_nw61).ArraySet1(_375_v142, 8)
		(_nw61).ArraySet1(_375_v142, 9)
		(_nw61).ArraySet1(_375_v142, 10)
		(_nw61).ArraySet1(_375_v142, 11)
		(_nw61).ArraySet1(_375_v142, 12)
		(_nw61).ArraySet1(_375_v142, 13)
		(_nw61).ArraySet1(_375_v142, 14)
		(_nw61).ArraySet1(_375_v142, 15)
		(_nw61).ArraySet1(_375_v142, 16)
		(_nw61).ArraySet1(_375_v142, 17)
		(_nw61).ArraySet1(_375_v142, 18)
		(_nw61).ArraySet1(_375_v142, 19)
		(_nw61).ArraySet1(_375_v142, 20)
		(_nw61).ArraySet1(_375_v142, 21)
		(_nw61).ArraySet1(_375_v142, 22)
		_379_v144 = _nw61
		var _382_v145 D8
		_ = _382_v145
		_382_v145 = Companion_D8_.Create_DC21_(_183_v6, _379_v144, _175_v1, _375_v142, Companion_Default___.Fm1(_183_v6, _dafny.IntOfInt64(759), _175_v1, _185_globalState))
		var _383_v146 _dafny.MultiSet
		_ = _383_v146
		_383_v146 = _dafny.MultiSetOf(_183_v6, (_382_v145).Dtor_cf49())
		var _384_v147 *C2
		_ = _384_v147
		var _nw62 *C2 = New_C2_()
		_ = _nw62
		_nw62.Ctor__(true, _383_v146, !(_183_v6))
		_384_v147 = _nw62
		var _385_v148 _dafny.Map
		_ = _385_v148
		_385_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v147, _375_v142)
		var _pat_let_tv13 = _379_v144
		_ = _pat_let_tv13
		var _386_v149 _dafny.Array
		_ = _386_v149
		var _nwElement0_12 _dafny.Array = _375_v142
		_ = _nwElement0_12
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(23))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_12, 0)
		(_nw63).ArraySet1(_375_v142, 1)
		(_nw63).ArraySet1(_375_v142, 2)
		(_nw63).ArraySet1(_375_v142, 3)
		(_nw63).ArraySet1((func() _dafny.Array {
			if (_385_v148).Contains(_384_v147) {
				return (_385_v148).Get(_384_v147).(_dafny.Array)
			}
			return _375_v142
		})(), 4)
		(_nw63).ArraySet1(_375_v142, 5)
		(_nw63).ArraySet1(_375_v142, 6)
		(_nw63).ArraySet1(_375_v142, 7)
		(_nw63).ArraySet1((func(_pat_let15_0 D8) D8 {
			return func(_387_dt__update__tmp_h2 D8) D8 {
				return func(_pat_let16_0 _dafny.Array) D8 {
					return func(_388_dt__update_hcf50_h0 _dafny.Array) D8 {
						return Companion_D8_.Create_DC21_((_387_dt__update__tmp_h2).Dtor_cf49(), _388_dt__update_hcf50_h0, (_387_dt__update__tmp_h2).Dtor_cf51(), (_387_dt__update__tmp_h2).Dtor_cf52(), (_387_dt__update__tmp_h2).Dtor_cf53())
					}(_pat_let16_0)
				}(_pat_let_tv13)
			}(_pat_let15_0)
		}(_382_v145)).Dtor_cf52(), 8)
		(_nw63).ArraySet1(_375_v142, 9)
		(_nw63).ArraySet1(_375_v142, 10)
		(_nw63).ArraySet1(_375_v142, 11)
		(_nw63).ArraySet1(_375_v142, 12)
		(_nw63).ArraySet1(_375_v142, 13)
		(_nw63).ArraySet1(_375_v142, 14)
		(_nw63).ArraySet1(_375_v142, 15)
		(_nw63).ArraySet1(_375_v142, 16)
		(_nw63).ArraySet1(_375_v142, 17)
		(_nw63).ArraySet1(_375_v142, 18)
		(_nw63).ArraySet1(_375_v142, 19)
		(_nw63).ArraySet1(_375_v142, 20)
		(_nw63).ArraySet1(_375_v142, 21)
		(_nw63).ArraySet1(_375_v142, 22)
		_386_v149 = _nw63
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))
		_ = _index41
		(_379_v144).ArraySet1(_375_v142, (_index41).Int())
		var _389_v150 _dafny.Map
		_ = _389_v150
		_389_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _384_v147.F15)
		var _390_v151 _dafny.Map
		_ = _390_v151
		_390_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_389_v150).Contains(_175_v1) {
				return (_389_v150).Get(_175_v1).(bool)
			}
			return _384_v147.F15
		})(), _375_v142)
		var _391_v152 _dafny.MultiSet
		_ = _391_v152
		_391_v152 = _dafny.MultiSetOf(_175_v1, _175_v1, _dafny.IntOfUint32((_184_v7).Cardinality()))
		var _392_v153 _dafny.Map
		_ = _392_v153
		_392_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, _175_v1)
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))
		_ = _index42
		(_379_v144).ArraySet1((func() _dafny.Array {
			if (_390_v151).Contains((_384_v147).Fm2(_391_v152, _175_v1, _392_v153, !(_183_v6), _185_globalState)) {
				return (_390_v151).Get((_384_v147).Fm2(_391_v152, _175_v1, _392_v153, !(_183_v6), _185_globalState)).(_dafny.Array)
			}
			return _375_v142
		})(), (_index42).Int())
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))
		_ = _arr0
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))
		_ = _index43
		(_arr0).ArraySet1(_384_v147.F15, (_index43).Int())
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))
		_ = _arr1
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))
		_ = _index44
		(_arr1).ArraySet1(_183_v6, (_index44).Int())
		var _393_v154 *C0
		_ = _393_v154
		var _nw64 *C0 = New_C0_()
		_ = _nw64
		_nw64.Ctor__((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))).Int()).(bool), _175_v1, _383_v146, _183_v6)
		_393_v154 = _nw64
		if _384_v147.F15 {
			_183_v6 = (((_392_v153).Cardinality()).Times((_393_v154).F18())).Cmp((_393_v154).F18()) >= 0
			(_185_globalState).F8 = (_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))).Int()).(bool)
			var _394_v155 _dafny.Set
			_ = _394_v155
			_394_v155 = _dafny.SetOf(_177_v3, _177_v3)
			var _395_v156 _dafny.Array
			_ = _395_v156
			var _nwElement0_13 _dafny.Sequence = _248_v54
			_ = _nwElement0_13
			var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(27))
			_ = _nw65
			(_nw65).ArraySet1(_nwElement0_13, 0)
			(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-298))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_396_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_397_i15 _dafny.Int) _dafny.CodePoint {
					return _396_v3
				}
			})(_177_v3))), 1)
			(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ecbkysbx"), 2)
			(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mekb"), 3)
			(_nw65).ArraySet1(_248_v54, 4)
			(_nw65).ArraySet1(_248_v54, 5)
			(_nw65).ArraySet1(_248_v54, 6)
			(_nw65).ArraySet1(_248_v54, 7)
			(_nw65).ArraySet1(_248_v54, 8)
			(_nw65).ArraySet1(_248_v54, 9)
			(_nw65).ArraySet1(_248_v54, 10)
			(_nw65).ArraySet1(_248_v54, 11)
			(_nw65).ArraySet1(_248_v54, 12)
			(_nw65).ArraySet1(_248_v54, 13)
			(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cnflkujmg"), 14)
			(_nw65).ArraySet1(_248_v54, 15)
			(_nw65).ArraySet1(_248_v54, 16)
			(_nw65).ArraySet1(_248_v54, 17)
			(_nw65).ArraySet1(Companion_Default___.Fm9((_394_v155).Cardinality(), true, _185_globalState), 18)
			(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(432))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_398_i16 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), 19)
			(_nw65).ArraySet1(_248_v54, 20)
			(_nw65).ArraySet1(_248_v54, 21)
			(_nw65).ArraySet1(_248_v54, 22)
			(_nw65).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("c"), 23)
			(_nw65).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(325))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_399_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_400_i17 _dafny.Int) _dafny.CodePoint {
					return _399_v3
				}
			})(_177_v3))), 24)
			(_nw65).ArraySet1(_248_v54, 25)
			(_nw65).ArraySet1(_248_v54, 26)
			_395_v156 = _nw65
			var _401_v157 D5
			_ = _401_v157
			_401_v157 = Companion_D5_.Create_DC13_((_392_v153).Cardinality(), true, _175_v1, _395_v156, (_393_v154).F17())
			var _402_v158 _dafny.Set
			_ = _402_v158
			_402_v158 = _dafny.SetOf(_401_v157, _401_v157)
			var _403_v159 _dafny.Sequence
			_ = _403_v159
			_403_v159 = _dafny.SeqOf(_402_v158)
			var _404_v160 D12
			_ = _404_v160
			_404_v160 = Companion_D12_.Create_DC31_(_403_v159)
			var _405_v161 _dafny.Sequence
			_ = _405_v161
			_405_v161 = _dafny.SeqOf(_403_v159, _dafny.SeqOf(_402_v158))
			(_384_v147).F15 = _dafny.Companion_Sequence_.Equal((_404_v160).Dtor_cf70(), (_405_v161).Select((Companion_Default___.SafeIndex((_393_v154).F18(), _dafny.IntOfUint32((_405_v161).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _406_v162 *C1
			_ = _406_v162
			var _nw66 *C1 = New_C1_()
			_ = _nw66
			_nw66.Ctor__((_dafny.Zero).Minus((_393_v154).F18()), _dafny.MultiSetOf(_183_v6), ((_393_v154).F18()).Cmp((_393_v154).F18()) == 0)
			_406_v162 = _nw66
			var _407_v163 _dafny.MultiSet
			_ = _407_v163
			_407_v163 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ygwxn"))
			var _408_v164 _dafny.Map
			_ = _408_v164
			_408_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_393_v154).F17(), _407_v163)
			_407_v163 = ((func() _dafny.MultiSet {
				if (_408_v164).Contains(_183_v6) {
					return (_408_v164).Get(_183_v6).(_dafny.MultiSet)
				}
				return _407_v163
			})()).Difference(_407_v163)
		} else {
			Companion_Default___.M0(_185_globalState)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_180_v5), 0))
			_ = _index45
			(_180_v5).ArraySet1((_393_v154).F18(), (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_180_v5), 0))
			_ = _index46
			(_180_v5).ArraySet1((func() _dafny.Int {
				if (_383_v146).Contains((_384_v147).Fm2(_391_v152, (_393_v154).F18(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_393_v154).F18()), _183_v6, _185_globalState)) {
					return (_383_v146).Multiplicity((_384_v147).Fm2(_391_v152, (_393_v154).F18(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_393_v154).F18()), _183_v6, _185_globalState))
				}
				return (_393_v154).F18()
			})(), (_index46).Int())
			var _409_v165 _dafny.MultiSet
			_ = _409_v165
			_409_v165 = _dafny.MultiSetOf(_177_v3)
			var _410_v166 _dafny.Map
			_ = _410_v166
			_410_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))).Int()).(bool), _409_v165)
			var _411_v167 _dafny.Set
			_ = _411_v167
			_411_v167 = _dafny.SetOf(!(_183_v6))
			var _412_v168 _dafny.Array
			_ = _412_v168
			var _nwElement0_14 _dafny.Int = (_dafny.Zero).Minus((_393_v154).F18())
			_ = _nwElement0_14
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(14))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_14, 0)
			(_nw67).ArraySet1((func() _dafny.Int {
				if (_393_v154).F17() {
					return _dafny.IntOfInt64(-420)
				}
				return ((func() _dafny.MultiSet {
					if (_410_v166).Contains(false) {
						return (_410_v166).Get(false).(_dafny.MultiSet)
					}
					return _409_v165
				})()).Cardinality()
			})(), 1)
			(_nw67).ArraySet1((_dafny.Zero).Minus(((_dafny.MultiSetOf((_393_v154).F17(), _183_v6)).Cardinality()).Minus((_393_v154).F18())), 2)
			(_nw67).ArraySet1(_175_v1, 3)
			(_nw67).ArraySet1((_411_v167).Cardinality(), 4)
			(_nw67).ArraySet1(_175_v1, 5)
			(_nw67).ArraySet1((_175_v1).Minus(_dafny.IntOfInt64(844)), 6)
			(_nw67).ArraySet1(((_393_v154).Fm12(_185_globalState)).Times((_393_v154).F18()), 7)
			(_nw67).ArraySet1(_dafny.IntOfInt64(620), 8)
			(_nw67).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(610), (_393_v154).F18())).Cardinality(), 9)
			(_nw67).ArraySet1((_393_v154).F18(), 10)
			(_nw67).ArraySet1(Companion_Default___.SafeModuloInt(_175_v1, (_393_v154).F18()), 11)
			(_nw67).ArraySet1((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), 12)
			(_nw67).ArraySet1(Companion_Default___.SafeDivisionInt(_175_v1, _dafny.IntOfUint32((_248_v54).Cardinality())), 13)
			_412_v168 = _nw67
			(_185_globalState).F1 = _412_v168
			var _413_v169 _dafny.Set
			_ = _413_v169
			_413_v169 = _dafny.SetOf(_177_v3, _177_v3)
			var _414_v170 _dafny.Map
			_ = _414_v170
			_414_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_413_v169).Cardinality(), _391_v152)
			var _415_v171 _dafny.Map
			_ = _415_v171
			_415_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _383_v146)
			var _416_v172 _dafny.Array
			_ = _416_v172
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw68
			_416_v172 = _nw68
			var _417_v173 D5
			_ = _417_v173
			_417_v173 = Companion_D5_.Create_DC13_(_175_v1, false, (_415_v171).Cardinality(), _416_v172, (_393_v154).F17())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_180_v5), 0))
			_ = _index47
			var _rhs47 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_248_v54, _248_v54)
			_ = _rhs47
			var _rhs48 _dafny.Int = (_393_v154).F18()
			_ = _rhs48
			var _rhs49 bool = (_384_v147).Fm2(_dafny.MultiSetOf((_393_v154).F18(), ((func() _dafny.MultiSet {
				if (_414_v170).Contains((_249_v55).Cardinality()) {
					return (_414_v170).Get((_249_v55).Cardinality()).(_dafny.MultiSet)
				}
				return _391_v152
			})()).Cardinality()), (_417_v173).Dtor_cf32(), (Companion_Default___.Fm23((_dafny.Zero).Minus((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int)), _185_globalState)).Update(false, _175_v1), (_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))).Int()).(bool), _185_globalState)
			_ = _rhs49
			var _rhs50 _dafny.MultiSet = ((_dafny.MultiSetFromSeq(_176_v2)).Intersection(Companion_Default___.Fm13((_248_v54).Select((Companion_Default___.SafeIndex((_393_v154).F18(), _dafny.IntOfUint32((_248_v54).Cardinality()))).Uint32()).(_dafny.CodePoint), false, _185_globalState))).Union(_391_v152)
			_ = _rhs50
			var _lhs38 _dafny.Array = _180_v5
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(527), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs39
			var _lhs40 *GlobalState = _185_globalState
			_ = _lhs40
			_248_v54 = _rhs47
			(_lhs38).ArraySet1(_rhs48, (_lhs39).Int())
			_lhs40.F7 = _rhs49
			_391_v152 = _rhs50
			(_185_globalState).F9 = (_393_v154).Fm12(_185_globalState)
		}
		var _418_v174 D5
		_ = _418_v174
		_418_v174 = Companion_D5_.Create_DC14_()
		var _source4 D5 = _418_v174
		_ = _source4
		if _source4.Is_DC13() {
			var _419___mcc_h12 _dafny.Int = _source4.Get_().(D5_DC13).Cf30
			_ = _419___mcc_h12
			var _420___mcc_h13 bool = _source4.Get_().(D5_DC13).Cf31
			_ = _420___mcc_h13
			var _421___mcc_h14 _dafny.Int = _source4.Get_().(D5_DC13).Cf32
			_ = _421___mcc_h14
			var _422___mcc_h15 _dafny.Array = _source4.Get_().(D5_DC13).Cf33
			_ = _422___mcc_h15
			var _423___mcc_h16 bool = _source4.Get_().(D5_DC13).Cf34
			_ = _423___mcc_h16
			var _424_cf34 bool = _423___mcc_h16
			_ = _424_cf34
			var _425_cf33 _dafny.Array = _422___mcc_h15
			_ = _425_cf33
			var _426_cf32 _dafny.Int = _421___mcc_h14
			_ = _426_cf32
			var _427_cf31 bool = _420___mcc_h13
			_ = _427_cf31
			var _428_cf30 _dafny.Int = _419___mcc_h12
			_ = _428_cf30
			var _429_v175 *C4
			_ = _429_v175
			var _nw69 *C4 = New_C4_()
			_ = _nw69
			_nw69.Ctor__(_383_v146, (_249_v55).Contains((_393_v154).F18()))
			_429_v175 = _nw69
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))
			_ = _index48
			(_180_v5).ArraySet1(_dafny.IntOfInt64(135), (_index48).Int())
			var _430_v176 *C3
			_ = _430_v176
			var _nw70 *C3 = New_C3_()
			_ = _nw70
			_nw70.Ctor__(_426_cf32, (_393_v154).F17(), _383_v146, false)
			_430_v176 = _nw70
			var _431_v177 _dafny.Map
			_ = _431_v177
			_431_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_430_v176).F13(), _428_cf30)
			var _432_v178 _dafny.Map
			_ = _432_v178
			_432_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_430_v176, (_431_v177).Cardinality())
			var _433_v179 _dafny.MultiSet
			_ = _433_v179
			_433_v179 = _dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Int {
				if (_432_v178).Contains(_430_v176) {
					return (_432_v178).Get(_430_v176).(_dafny.Int)
				}
				return (_393_v154).F18()
			})(), Companion_Default___.Fm1(!(_384_v147.F15), _175_v1, _175_v1, _185_globalState), _175_v1), (_391_v152).Difference(_391_v152), _391_v152, _391_v152)
			var _434_v180 D9
			_ = _434_v180
			_434_v180 = Companion_D9_.Create_DC24_(_248_v54, _384_v147.F15, _175_v1)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))
			_ = _index49
			var _rhs51 _dafny.CodePoint = _177_v3
			_ = _rhs51
			var _rhs52 *C4 = _429_v175
			_ = _rhs52
			var _rhs53 _dafny.Set = (_249_v55).Intersection(_249_v55)
			_ = _rhs53
			var _rhs54 _dafny.Int = _dafny.IntOfUint32(((_434_v180).Dtor_cf56()).Cardinality())
			_ = _rhs54
			var _rhs55 _dafny.MultiSet = _433_v179
			_ = _rhs55
			var _lhs41 _dafny.Array = _180_v5
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs42
			_177_v3 = _rhs51
			_429_v175 = _rhs52
			_249_v55 = _rhs53
			(_lhs41).ArraySet1(_rhs54, (_lhs42).Int())
			_433_v179 = _rhs55
			(_185_globalState).F9 = (_393_v154).F18()
			var _435_v181 _dafny.Sequence
			_ = _435_v181
			_435_v181 = _dafny.SeqOf(_383_v146)
			var _436_v182 T0
			_ = _436_v182
			var _nw71 *C4 = New_C4_()
			_ = _nw71
			_nw71.Ctor__((_435_v181).Select((Companion_Default___.SafeIndex((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_435_v181).Cardinality()))).Uint32()).(_dafny.MultiSet), !(_424_cf34))
			_436_v182 = _nw71
			var _437_v183 _dafny.Map
			_ = _437_v183
			_437_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_436_v182, _dafny.Companion_Sequence_.Update(_248_v54, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_430_v176).F13()), _dafny.IntOfUint32((_248_v54).Cardinality()))).Uint32(), _177_v3))
			var _438_v184 _dafny.Sequence
			_ = _438_v184
			_438_v184 = _dafny.SeqOf(Companion_Default___.Fm23(_dafny.IntOfUint32((_176_v2).Cardinality()), _185_globalState), Companion_Default___.Fm23((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), _185_globalState))
			var _439_v185 _dafny.MultiSet
			_ = _439_v185
			_439_v185 = _dafny.MultiSetOf(_436_v182, _436_v182)
			var _440_v186 _dafny.Map
			_ = _440_v186
			_440_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(467), _386_v149)
			var _441_v187 _dafny.Set
			_ = _441_v187
			_441_v187 = _dafny.SetOf(_436_v182.F12())
			var _442_v188 _dafny.Set
			_ = _442_v188
			_442_v188 = _dafny.SetOf(_441_v187)
			var _arr2 _dafny.Array = _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))
			_ = _arr2
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))
			_ = _index50
			var _rhs56 _dafny.Map = _437_v183
			_ = _rhs56
			var _rhs57 _dafny.Int = (func() _dafny.Int {
				if (_439_v185).Contains(_436_v182) {
					return (_439_v185).Multiplicity(_436_v182)
				}
				return Companion_Default___.SafeDivisionInt((_393_v154).F18(), _dafny.IntOfInt64(222))
			})()
			_ = _rhs57
			var _rhs58 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_248_v54, _248_v54), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(805))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_443_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_444_i18 _dafny.Int) _dafny.CodePoint {
					return _443_v3
				}
			})(_177_v3))))
			_ = _rhs58
			var _rhs59 D8 = Companion_D8_.Create_DC21_(_384_v147.F15, (func() _dafny.Array {
				if (_440_v186).Contains((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int)) {
					return (_440_v186).Get((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int)).(_dafny.Array)
				}
				return _379_v144
			})(), Companion_Default___.SafeModuloInt((_180_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_180_v5), 0))).Int()).(_dafny.Int), (_393_v154).F18()), _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int())), ((_442_v188).Cardinality()).Times(Companion_Default___.Fm1(_427_cf31, _428_cf30, _428_cf30, _185_globalState)))
			_ = _rhs59
			var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_438_v184, _438_v184)
			_ = _rhs60
			var _lhs43 _dafny.Array = _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))
			_ = _lhs44
			_437_v183 = _rhs56
			_428_cf30 = _rhs57
			(_lhs43).ArraySet1(_rhs58, (_lhs44).Int())
			_382_v145 = _rhs59
			_438_v184 = _rhs60
			var _445_v189 _dafny.Int
			_ = _445_v189
			var _446_v190 bool
			_ = _446_v190
			var _out10 _dafny.Int
			_ = _out10
			var _out11 bool
			_ = _out11
			_out10, _out11 = (_430_v176).M2(_dafny.UnicodeSeqOfUtf8Bytes("jb"), (_430_v176).F14(), _425_cf33, _185_globalState)
			_445_v189 = _out10
			_446_v190 = _out11
		} else if _source4.Is_DC14() {
			(_384_v147).F15 = true
			(_185_globalState).F3 = _384_v147.F15
			_392_v153 = (_392_v153).Update((_393_v154).Fm2(_391_v152, _dafny.IntOfUint32((_378_v143).Cardinality()), _392_v153, _183_v6, _185_globalState), (_393_v154).F18())
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))
			_ = _arr3
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))
			_ = _index51
			var _rhs61 bool = !(!(false))
			_ = _rhs61
			var _rhs62 _dafny.MultiSet = (_391_v152).Difference((_391_v152).Union(_391_v152))
			_ = _rhs62
			var _rhs63 bool = (_393_v154).F17()
			_ = _rhs63
			var _rhs64 bool = !((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))).Int()).(bool))
			_ = _rhs64
			var _lhs45 *GlobalState = _185_globalState
			_ = _lhs45
			var _lhs46 _dafny.Array = _dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))
			_ = _lhs47
			_183_v6 = _rhs61
			_391_v152 = _rhs62
			_lhs45.F3 = _rhs63
			(_lhs46).ArraySet1(_rhs64, (_lhs47).Int())
		} else {
			var _447___mcc_h17 _dafny.Sequence = _source4.Get_().(D5_DC12).Cf29
			_ = _447___mcc_h17
			var _448_cf29 _dafny.Sequence = _447___mcc_h17
			_ = _448_cf29
			var _449_v191 *C2
			_ = _449_v191
			var _nw72 *C2 = New_C2_()
			_ = _nw72
			_nw72.Ctor__(_384_v147.F15, _383_v146, (_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_dafny.ArrayCastTo((_379_v144).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_379_v144), 0))).Int()))), 0))).Int()).(bool))
			_449_v191 = _nw72
			(_185_globalState).F9 = _dafny.IntOfUint32((_248_v54).Cardinality())
			var _450_v192 bool
			_ = _450_v192
			var _451_v193 _dafny.Int
			_ = _451_v193
			var _out12 bool
			_ = _out12
			var _out13 _dafny.Int
			_ = _out13
			_out12, _out13 = (_449_v191).M5(_dafny.IntOfInt64(765), _dafny.Companion_Sequence_.Concatenate(_248_v54, _248_v54), _185_globalState)
			_450_v192 = _out12
			_451_v193 = _out13
			_451_v193 = (_dafny.Zero).Minus(_175_v1)
		}
	} else {
		var _452_v194 _dafny.Array
		_ = _452_v194
		var _nw73 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(6))
		_ = _nw73
		_452_v194 = _nw73
		var _453_v195 _dafny.Map
		_ = _453_v195
		_453_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v6, _183_v6)
		var _454_v196 *C1
		_ = _454_v196
		var _nw74 *C1 = New_C1_()
		_ = _nw74
		_nw74.Ctor__(_175_v1, (_dafny.MultiSetFromSeq(_184_v7)).Update(_183_v6, Companion_Default___.Abs((_453_v195).Cardinality())), _183_v6)
		_454_v196 = _nw74
		var _455_v197 _dafny.MultiSet
		_ = _455_v197
		_455_v197 = _dafny.MultiSetOf(_175_v1, _454_v196.F16)
		var _456_v198 _dafny.Map
		_ = _456_v198
		_456_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_454_v196, _455_v197)
		var _457_v199 D6
		_ = _457_v199
		_457_v199 = Companion_D6_.Create_DC15_((func() _dafny.MultiSet {
			if (_456_v198).Contains(_454_v196) {
				return (_456_v198).Get(_454_v196).(_dafny.MultiSet)
			}
			return _455_v197
		})())
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_452_v194), 0))
		_ = _index52
		(_452_v194).ArraySet1(_457_v199, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_452_v194), 0))
		_ = _index53
		(_452_v194).ArraySet1(_457_v199, (_index53).Int())
		var _458_v200 D1
		_ = _458_v200
		_458_v200 = Companion_D1_.Create_DC2_(false, _183_v6, _183_v6, _177_v3, _177_v3)
		var _459_v201 D1
		_ = _459_v201
		_459_v201 = Companion_D1_.Create_DC3_(_458_v200)
		var _460_v202 D1
		_ = _460_v202
		_460_v202 = Companion_D1_.Create_DC3_(_459_v201)
		var _461_v203 D1
		_ = _461_v203
		_461_v203 = Companion_D1_.Create_DC3_(_460_v202)
		var _462_v204 D1
		_ = _462_v204
		_462_v204 = Companion_D1_.Create_DC3_(_461_v203)
		_462_v204 = _462_v204
		var _463_v206 _dafny.Map
		_ = _463_v206
		_463_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_453_v195).Cardinality(), _453_v195)
		_175_v1 = Companion_Default___.SafeDivisionInt(((func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(359), _dafny.IntOfInt64(-879))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _464_v205 _dafny.Int
				_464_v205 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(359)).Cmp(_464_v205) <= 0) && ((_464_v205).Cmp(_dafny.IntOfInt64(-879)) < 0) {
					_coll14.Add(Companion_Default___.SafeDivisionInt(_464_v205, (_dafny.MultiSetFromSeq(_184_v7)).Cardinality()), (_453_v195).Update(!(_183_v6), _183_v6))
				}
			}
			return _coll14.ToMap()
		}()).Merge((_463_v206).Update(_dafny.IntOfUint32((_248_v54).Cardinality()), _453_v195))).Cardinality(), _dafny.IntOfUint32((_248_v54).Cardinality()))
		var _465_v207 _dafny.Array
		_ = _465_v207
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_8
		var _nw75 _dafny.Array
		_ = _nw75
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw75 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = (func(_466_v7 _dafny.Sequence, _467_v1 _dafny.Int) func(_dafny.Int) bool {
				return func(_468_i19 _dafny.Int) bool {
					return (_466_v7).Select((Companion_Default___.SafeIndex(_467_v1, _dafny.IntOfUint32((_466_v7).Cardinality()))).Uint32()).(bool)
				}
			})(_184_v7, _175_v1)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw75 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw75).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw75).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_465_v207 = _nw75
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))
		_ = _index54
		(_465_v207).ArraySet1(_183_v6, (_index54).Int())
		var _469_v208 _dafny.Set
		_ = _469_v208
		_469_v208 = _dafny.SetOf(true)
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))
		_ = _index55
		(_465_v207).ArraySet1(!(!(_dafny.SetOf(_183_v6, _183_v6, _183_v6)).Equals(_469_v208)) || ((func() bool {
			if _183_v6 {
				return _183_v6
			}
			return _183_v6
		})()), (_index55).Int())
		var _470_v209 _dafny.Set
		_ = _470_v209
		_470_v209 = _dafny.SetOf((_251_v57).Select((Companion_Default___.SafeIndex(_454_v196.F16, _dafny.IntOfUint32((_251_v57).Cardinality()))).Uint32()).(_dafny.Sequence))
		if ((_470_v209).Intersection(func() _dafny.Set {
			var _coll15 = _dafny.NewBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v54, _183_v6)).Keys().Elements()); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _471_v210 _dafny.Sequence
				_471_v210 = interface{}(_compr_15).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v54, _183_v6)).Contains(_471_v210) {
					_coll15.Add(_471_v210)
				}
			}
			return _coll15.ToSet()
		}())).IsProperSubsetOf(_470_v209) {
			_184_v7 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_184_v7, _184_v7), (Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_184_v7, _184_v7)).Cardinality()))).Uint32(), _183_v6)
			_248_v54 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_472_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_473_i20 _dafny.Int) _dafny.CodePoint {
					return _472_v3
				}
			})(_177_v3)))
			var _474_v211 _dafny.MultiSet
			_ = _474_v211
			_474_v211 = _dafny.MultiSetOf((_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool), _183_v6, (_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool))
			var _475_v212 *C4
			_ = _475_v212
			var _nw76 *C4 = New_C4_()
			_ = _nw76
			_nw76.Ctor__(_474_v211, false)
			_475_v212 = _nw76
			var _476_v213 *C0
			_ = _476_v213
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__((_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool), _dafny.IntOfInt64(449), _474_v211, (_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool))
			_476_v213 = _nw77
			_183_v6 = (_476_v213).F17()
		} else {
			var _477_v214 _dafny.MultiSet
			_ = _477_v214
			_477_v214 = _dafny.MultiSetOf((_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool))
			var _478_v215 _dafny.Array
			_ = _478_v215
			var _nwElement0_15 _dafny.Array = _465_v207
			_ = _nwElement0_15
			var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(20))
			_ = _nw78
			(_nw78).ArraySet1(_nwElement0_15, 0)
			(_nw78).ArraySet1(_465_v207, 1)
			(_nw78).ArraySet1(_465_v207, 2)
			(_nw78).ArraySet1(_465_v207, 3)
			(_nw78).ArraySet1(_465_v207, 4)
			(_nw78).ArraySet1(_465_v207, 5)
			(_nw78).ArraySet1(_465_v207, 6)
			(_nw78).ArraySet1(_465_v207, 7)
			(_nw78).ArraySet1(_465_v207, 8)
			(_nw78).ArraySet1(_465_v207, 9)
			(_nw78).ArraySet1(_465_v207, 10)
			(_nw78).ArraySet1(_465_v207, 11)
			(_nw78).ArraySet1(_465_v207, 12)
			(_nw78).ArraySet1(_465_v207, 13)
			(_nw78).ArraySet1(_465_v207, 14)
			(_nw78).ArraySet1(_465_v207, 15)
			(_nw78).ArraySet1(_465_v207, 16)
			(_nw78).ArraySet1(_465_v207, 17)
			(_nw78).ArraySet1(_465_v207, 18)
			(_nw78).ArraySet1(_465_v207, 19)
			_478_v215 = _nw78
			var _479_v216 D8
			_ = _479_v216
			_479_v216 = Companion_D8_.Create_DC21_(false, _478_v215, _175_v1, _465_v207, (_dafny.SetOf((_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool), (_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool))).Cardinality())
			var _480_v217 *C4
			_ = _480_v217
			var _nw79 *C4 = New_C4_()
			_ = _nw79
			_nw79.Ctor__(_477_v214, (_479_v216).Dtor_cf49())
			_480_v217 = _nw79
			var _481_v218 _dafny.Map
			_ = _481_v218
			_481_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _480_v217)
			var _482_v219 _dafny.Map
			_ = _482_v219
			_482_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_454_v196.F16, (_465_v207).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_465_v207), 0))).Int()).(bool))
			_481_v218 = (_481_v218).Update((!(_183_v6)) && ((func() bool {
				if (_482_v219).Contains(_454_v196.F16) {
					return (_482_v219).Get(_454_v196.F16).(bool)
				}
				return _183_v6
			})()), _480_v217)
			_477_v214 = _477_v214
			(_185_globalState).F1 = (func() _dafny.Array {
				if (_480_v217).Fm4(_454_v196.F16, _454_v196.F16, _185_globalState) {
					return _180_v5
				}
				return _180_v5
			})()
			(_185_globalState).F9 = _454_v196.F16
			(_454_v196).F16 = _175_v1
		}
	}
	var _483_v220 _dafny.Set
	_ = _483_v220
	_483_v220 = _dafny.SetOf(_183_v6)
	var _484_v221 *C0
	_ = _484_v221
	var _nw80 *C0 = New_C0_()
	_ = _nw80
	_nw80.Ctor__(_183_v6, (_483_v220).Cardinality(), _dafny.MultiSetOf(_183_v6, _183_v6), _183_v6)
	_484_v221 = _nw80
	var _485_v222 _dafny.MultiSet
	_ = _485_v222
	_485_v222 = _dafny.MultiSetOf((_484_v221).F17(), _183_v6, _183_v6)
	var _486_v223 *C3
	_ = _486_v223
	var _nw81 *C3 = New_C3_()
	_ = _nw81
	_nw81.Ctor__(_175_v1, (_484_v221).F17(), _485_v222, _183_v6)
	_486_v223 = _nw81
	var _487_v224 _dafny.Map
	_ = _487_v224
	_487_v224 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _486_v223)
	var _488_v225 _dafny.Sequence
	_ = _488_v225
	_488_v225 = _dafny.SeqOf(_485_v222, _485_v222, _485_v222, Companion_Default___.Fm22((_487_v224).Cardinality(), _183_v6, _185_globalState))
	var _489_v226 _dafny.Map
	_ = _489_v226
	_489_v226 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_v221, ((_488_v225).Select((Companion_Default___.SafeIndex((_484_v221).F18(), _dafny.IntOfUint32((_488_v225).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsDisjointFrom(Companion_Default___.Fm22((_486_v223).F13(), (_486_v223).F14(), _185_globalState)))
	_489_v226 = (_489_v226).Update(_484_v221, !((_486_v223).F14()))
	var _490_v227 _dafny.Array
	_ = _490_v227
	var _len0_9 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_9
	var _nw82 _dafny.Array
	_ = _nw82
	if _len0_9.Cmp(_dafny.Zero) == 0 {
		_nw82 = _dafny.NewArray(_len0_9)
	} else {
		var _init9 func(_dafny.Int) _dafny.Sequence = (func(_491_v54 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_492_i22 _dafny.Int) _dafny.Sequence {
				return _491_v54
			}
		})(_248_v54)
		_ = _init9
		var _element0_9 = _init9(_dafny.Zero)
		_ = _element0_9
		_nw82 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
		(_nw82).ArraySet1(_element0_9, 0)
		var _nativeLen0_9 = (_len0_9).Int()
		_ = _nativeLen0_9
		for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
			(_nw82).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
		}
	}
	_490_v227 = _nw82
	for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_490_v227), 0))); ; {
		_guard_loop_0, _ok16 := _iter16()
		if !_ok16 {
			break
		}
		var _493_i21 _dafny.Int
		_493_i21 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_493_i21).Sign() != -1) && ((_493_i21).Cmp(_dafny.ArrayLen((_490_v227), 0)) < 0)) {
			(_490_v227).ArraySet1(_248_v54, (_493_i21).Int())
		}
	}
	var _494_v228 *C1
	_ = _494_v228
	var _nw83 *C1 = New_C1_()
	_ = _nw83
	_nw83.Ctor__((_486_v223).F13(), _485_v222, (_484_v221).F17())
	_494_v228 = _nw83
	var _495_v229 _dafny.Sequence
	_ = _495_v229
	_495_v229 = _dafny.SeqOf(_494_v228, _494_v228)
	var _496_i23 _dafny.Int
	_ = _496_i23
	_496_i23 = _dafny.Zero
	{
		for !(!(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_495_v229, _dafny.SeqOf(_494_v228)), (Companion_Default___.SafeIndex(_175_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_495_v229, _dafny.SeqOf(_494_v228))).Cardinality()))).Uint32(), _494_v228), _495_v229))) {
			{
				if (_496_i23).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_496_i23 = (_496_i23).Plus(_dafny.One)
				var _497_v230 T0
				_ = _497_v230
				var _nw84 *C1 = New_C1_()
				_ = _nw84
				_nw84.Ctor__((_486_v223).F13(), _dafny.MultiSetOf(_183_v6), (_484_v221).F17())
				_497_v230 = _nw84
				var _498_v231 _dafny.Set
				_ = _498_v231
				_498_v231 = _dafny.SetOf(_497_v230)
				_251_v57 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(485))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}((func(_499_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_500_i24 _dafny.Int) _dafny.CodePoint {
						return _499_v3
					}
				})(_177_v3))), _248_v54, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ot"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_498_v231).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ot")).Cardinality()))).Uint32(), _dafny.CodePoint('r')), _248_v54)
				var _501_v233 _dafny.Array
				_ = _501_v233
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_10
				var _nw85 _dafny.Array
				_ = _nw85
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw85 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Map = (func(_502_v221 *C0, _503_v1 _dafny.Int, _504_v223 *C3, _505_v2 _dafny.Sequence, _506_v228 *C1) func(_dafny.Int) _dafny.Map {
						return func(_507_i25 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_502_v221).F18(), _503_v1), (_504_v223).F13())).Merge(func() _dafny.Map {
								var _coll16 = _dafny.NewMapBuilder()
								_ = _coll16
								for _iter17 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_505_v2, (Companion_Default___.SafeIndex(_506_v228.F16, _dafny.IntOfUint32((_505_v2).Cardinality()))).Uint32(), _503_v1), _505_v2)).Elements()); ; {
									_compr_16, _ok17 := _iter17()
									if !_ok17 {
										break
									}
									var _508_v232 _dafny.Sequence
									_508_v232 = interface{}(_compr_16).(_dafny.Sequence)
									if (_dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_505_v2, (Companion_Default___.SafeIndex(_506_v228.F16, _dafny.IntOfUint32((_505_v2).Cardinality()))).Uint32(), _503_v1), _505_v2)).Contains(_508_v232) {
										_coll16.Add(_508_v232, (_504_v223).F13())
									}
								}
								return _coll16.ToMap()
							}())
						}
					})(_484_v221, _175_v1, _486_v223, _176_v2, _494_v228)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw85 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw85).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw85).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_501_v233 = _nw85
				var _509_v234 _dafny.Map
				_ = _509_v234
				_509_v234 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_248_v54).Cardinality())), _175_v1)
				var _510_v235 _dafny.Map
				_ = _510_v235
				_510_v235 = _509_v234
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_501_v233), 0))
				_ = _index56
				(_501_v233).ArraySet1((_510_v235), (_index56).Int())
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_501_v233), 0))
				_ = _index57
				(_501_v233).ArraySet1(((_509_v234).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _175_v1))).Merge((_509_v234).Merge(func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter18 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}((func(_511_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_512_i26 _dafny.Int) _dafny.Sequence {
							return _511_v2
						}
					})(_176_v2)))).Elements()); ; {
						_compr_17, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _513_v236 _dafny.Sequence
						_513_v236 = interface{}(_compr_17).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg29 _dafny.Int) interface{} {
								return coer29(arg29)
							}
						}((func(_514_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_512_i26 _dafny.Int) _dafny.Sequence {
								return _514_v2
							}
						})(_176_v2))), _513_v236) {
							_coll17.Add(_513_v236, (_484_v221).F18())
						}
					}
					return _coll17.ToMap()
				}())), (_index57).Int())
				var _515_v237 _dafny.MultiSet
				_ = _515_v237
				_515_v237 = _dafny.MultiSetOf(_494_v228.F16, _dafny.IntOfInt64(813), (_484_v221).F18(), _dafny.IntOfUint32((_184_v7).Cardinality()))
				var _516_v238 _dafny.Map
				_ = _516_v238
				_516_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_497_v230.F12(), (_486_v223).F13())
				var _517_v239 _dafny.Sequence
				_ = _517_v239
				_517_v239 = _dafny.SeqOf(_176_v2, (func() _dafny.Sequence {
					if (_494_v228).Fm2(_515_v237, _dafny.IntOfInt64(673), _516_v238, (_486_v223).Fm6((_486_v223).F14(), _175_v1, _185_globalState), _185_globalState) {
						return _176_v2
					}
					return _176_v2
				})(), _dafny.Companion_Sequence_.Concatenate(_176_v2, _176_v2))
				var _518_v240 _dafny.Sequence
				_ = _518_v240
				_518_v240 = _dafny.SeqOf(_517_v239, _dafny.Companion_Sequence_.Concatenate(_517_v239, _517_v239))
				var _519_v241 D9
				_ = _519_v241
				_519_v241 = Companion_D9_.Create_DC24_(_248_v54, true, _dafny.IntOfInt64(707))
				_517_v239 = _dafny.Companion_Sequence_.Update((_518_v240).Select((Companion_Default___.SafeIndex((_486_v223).F13(), _dafny.IntOfUint32((_518_v240).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_484_v221).F18(), _dafny.IntOfUint32(((_518_v240).Select((Companion_Default___.SafeIndex((_486_v223).F13(), _dafny.IntOfUint32((_518_v240).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_519_v241).Dtor_cf58()), _176_v2))
				(_497_v230).F12_set_(!(_183_v6))
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	Companion_Default___.M0(_185_globalState)
	if _183_v6 {
		var _rhs65 _dafny.Sequence = _248_v54
		_ = _rhs65
		var _rhs66 _dafny.Sequence = _176_v2
		_ = _rhs66
		_248_v54 = _rhs65
		_176_v2 = _rhs66
		var _520_v242 _dafny.Map
		_ = _520_v242
		_520_v242 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_248_v54).Cardinality()), (_484_v221).F17())
		var _521_v243 T0
		_ = _521_v243
		var _nw86 *C0 = New_C0_()
		_ = _nw86
		_nw86.Ctor__((_486_v223).F14(), (_484_v221).F18(), _485_v222, _183_v6)
		_521_v243 = _nw86
		var _522_v244 _dafny.Map
		_ = _522_v244
		_522_v244 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v3, _521_v243)
		var _523_v245 _dafny.Sequence
		_ = _523_v245
		_523_v245 = _dafny.SeqOf((_522_v244).Update(_177_v3, _521_v243))
		var _524_v246 *C1
		_ = _524_v246
		var _nw87 *C1 = New_C1_()
		_ = _nw87
		_nw87.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(847), _dafny.IntOfUint32((_248_v54).Cardinality())), _dafny.MultiSetOf((func() bool {
			if (_520_v242).Contains(_dafny.IntOfUint32((_523_v245).Cardinality())) {
				return (_520_v242).Get(_dafny.IntOfUint32((_523_v245).Cardinality())).(bool)
			}
			return _521_v243.F12()
		})(), (_486_v223).F14()), !((_484_v221).F17()))
		_524_v246 = _nw87
		if !((_486_v223).Fm6(!(!(_183_v6)), (_484_v221).Fm12(_185_globalState), _185_globalState)) || (_521_v243.F12()) {
			var _525_v247 _dafny.Array
			_ = _525_v247
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw88
			_525_v247 = _nw88
			var _526_v248 _dafny.Sequence
			_ = _526_v248
			_526_v248 = _dafny.SeqOf(_525_v247)
			var _527_v249 _dafny.MultiSet
			_ = _527_v249
			_527_v249 = _dafny.MultiSetOf(_494_v228.F16, _175_v1, _dafny.IntOfUint32((_526_v248).Cardinality()), (_dafny.Zero).Minus(_494_v228.F16), _494_v228.F16)
			var _528_v250 _dafny.Map
			_ = _528_v250
			_528_v250 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v1, _527_v249)
			var _529_v251 _dafny.Array
			_ = _529_v251
			var _nwElement0_16 _dafny.MultiSet = Companion_Default___.Fm13(_177_v3, !((_484_v221).F17()), _185_globalState)
			_ = _nwElement0_16
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(12))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_16, 0)
			(_nw89).ArraySet1(_527_v249, 1)
			(_nw89).ArraySet1(_527_v249, 2)
			(_nw89).ArraySet1(_527_v249, 3)
			(_nw89).ArraySet1((func() _dafny.MultiSet {
				if (_528_v250).Contains((_484_v221).F18()) {
					return (_528_v250).Get((_484_v221).F18()).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_494_v228.F16)
			})(), 4)
			(_nw89).ArraySet1(_527_v249, 5)
			(_nw89).ArraySet1((_527_v249).Update((_486_v223).F13(), Companion_Default___.Abs((_486_v223).F13())), 6)
			(_nw89).ArraySet1(_527_v249, 7)
			(_nw89).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfUint32((_176_v2).Cardinality()))).Union(_527_v249), 8)
			(_nw89).ArraySet1(_dafny.MultiSetOf(_494_v228.F16, _524_v246.F16, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("co")).Cardinality())), 9)
			(_nw89).ArraySet1((func() _dafny.MultiSet {
				if _521_v243.F12() {
					return _527_v249
				}
				return _527_v249
			})(), 10)
			(_nw89).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfInt64(162), _494_v228.F16)).Update((_dafny.Zero).Minus(_175_v1), Companion_Default___.Abs(_494_v228.F16)), 11)
			_529_v251 = _nw89
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_529_v251), 0))
			_ = _index58
			(_529_v251).ArraySet1(Companion_Default___.Fm13(_dafny.CodePoint('h'), (_486_v223).F14(), _185_globalState), (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_529_v251), 0))
			_ = _index59
			(_529_v251).ArraySet1(_527_v249, (_index59).Int())
			var _530_v252 _dafny.Map
			_ = _530_v252
			_530_v252 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_176_v2, (Companion_Default___.SafeIndex(_524_v246.F16, _dafny.IntOfUint32((_176_v2).Cardinality()))).Uint32(), (_484_v221).F18()), _dafny.IntOfInt64(354))
			var _531_v253 _dafny.Map
			_ = _531_v253
			_531_v253 = _530_v252
			var _532_v254 _dafny.Map
			_ = _532_v254
			_532_v254 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_531_v253, (_486_v223).F14())
			var _533_v255 _dafny.Array
			_ = _533_v255
			var _nw90 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(10))
			_ = _nw90
			_533_v255 = _nw90
			var _534_v256 _dafny.Map
			_ = _534_v256
			_534_v256 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_532_v254, _533_v255)
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_180_v5), 0))
			_ = _index60
			(_180_v5).ArraySet1(Companion_Default___.SafeModuloInt((_484_v221).F18(), _494_v228.F16), (_index60).Int())
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_180_v5), 0))
			_ = _index61
			var _rhs67 _dafny.Map = _534_v256
			_ = _rhs67
			var _rhs68 bool = ((_486_v223).F13()).Cmp((_484_v221).F18()) != 0
			_ = _rhs68
			var _rhs69 *C3 = (func() *C3 {
				if (_184_v7).Select((Companion_Default___.SafeIndex((_484_v221).F18(), _dafny.IntOfUint32((_184_v7).Cardinality()))).Uint32()).(bool) {
					return (func() *C3 {
						if (_484_v221).F17() {
							return _486_v223
						}
						return _486_v223
					})()
				}
				return _486_v223
			})()
			_ = _rhs69
			var _rhs70 _dafny.Int = _494_v228.F16
			_ = _rhs70
			var _lhs48 T0 = _521_v243
			_ = _lhs48
			var _lhs49 _dafny.Array = _180_v5
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_180_v5), 0))
			_ = _lhs50
			_534_v256 = _rhs67
			_lhs48.F12_set_(_rhs68)
			_486_v223 = _rhs69
			(_lhs49).ArraySet1(_rhs70, (_lhs50).Int())
			var _535_v257 *C2
			_ = _535_v257
			var _nw91 *C2 = New_C2_()
			_ = _nw91
			_nw91.Ctor__(_521_v243.F12(), _521_v243.F11(), _521_v243.F12())
			_535_v257 = _nw91
			var _536_v258 D12
			_ = _536_v258
			_536_v258 = Companion_D12_.Create_DC32_(Companion_Default___.SafeDivisionInt((_484_v221).F18(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-861), _524_v246)).Cardinality()), _535_v257, _484_v221, (_486_v223).F13())
			var _537_v259 _dafny.Map
			_ = _537_v259
			_537_v259 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_524_v246.F16, _484_v221)
			var _538_v260 _dafny.Sequence
			_ = _538_v260
			_538_v260 = _dafny.SeqOf(_484_v221, (func() *C0 {
				if (_537_v259).Contains((_484_v221).F18()) {
					return (_537_v259).Get((_484_v221).F18()).(*C0)
				}
				return _484_v221
			})(), _484_v221)
			_536_v258 = Companion_D12_.Create_DC32_(_494_v228.F16, _535_v257, (_538_v260).Select((Companion_Default___.SafeIndex((_483_v220).Cardinality(), _dafny.IntOfUint32((_538_v260).Cardinality()))).Uint32()).(*C0), (_dafny.Zero).Minus(_dafny.IntOfUint32((_248_v54).Cardinality())))
			var _539_v261 *C3
			_ = _539_v261
			var _nw92 *C3 = New_C3_()
			_ = _nw92
			_nw92.Ctor__((_dafny.Zero).Minus((_486_v223).F13()), _521_v243.F12(), (_dafny.MultiSetOf(false, (_484_v221).F17())).Union(_485_v222), (_484_v221).F17())
			_539_v261 = _nw92
			var _540_v262 _dafny.Array
			_ = _540_v262
			var _nwElement0_17 _dafny.Int = ((_486_v223).F13()).Plus(_494_v228.F16)
			_ = _nwElement0_17
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(13))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_17, 0)
			(_nw93).ArraySet1((_486_v223).F13(), 1)
			(_nw93).ArraySet1((_494_v228.F16).Plus((_dafny.Zero).Minus(_524_v246.F16)), 2)
			(_nw93).ArraySet1(Companion_Default___.Fm1(_535_v257.F15, (_486_v223).F13(), (_249_v55).Cardinality(), _185_globalState), 3)
			(_nw93).ArraySet1((_dafny.Zero).Minus(((_486_v223).F13()).Times((_484_v221).F18())), 4)
			(_nw93).ArraySet1(_494_v228.F16, 5)
			(_nw93).ArraySet1((_486_v223).F13(), 6)
			(_nw93).ArraySet1((_dafny.Zero).Minus(_494_v228.F16), 7)
			(_nw93).ArraySet1((_486_v223).F13(), 8)
			(_nw93).ArraySet1(_494_v228.F16, 9)
			(_nw93).ArraySet1((func() _dafny.Int {
				if !(!(_535_v257.F15)) {
					return _dafny.IntOfUint32((_248_v54).Cardinality())
				}
				return _175_v1
			})(), 10)
			(_nw93).ArraySet1(_dafny.IntOfUint32((_184_v7).Cardinality()), 11)
			(_nw93).ArraySet1((_dafny.IntOfUint32((_248_v54).Cardinality())).Minus((_484_v221).F18()), 12)
			_540_v262 = _nw93
			var _rhs71 bool = !(!(true) || (_521_v243.F12()))
			_ = _rhs71
			var _rhs72 _dafny.Array = _540_v262
			_ = _rhs72
			var _rhs73 _dafny.Int = (func() _dafny.Int {
				if ((_539_v261).F13()).Cmp(_dafny.IntOfInt64(-876)) <= 0 {
					return (_175_v1).Plus((_484_v221).F18())
				}
				return _dafny.IntOfInt64(979)
			})()
			_ = _rhs73
			var _lhs51 *GlobalState = _185_globalState
			_ = _lhs51
			var _lhs52 *GlobalState = _185_globalState
			_ = _lhs52
			_lhs51.F3 = _rhs71
			_lhs52.F1 = _rhs72
			_175_v1 = _rhs73
		} else {
			_183_v6 = _521_v243.F12()
			var _541_v263 _dafny.Map
			_ = _541_v263
			_541_v263 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _494_v228.F16)
			var _542_v264 _dafny.Map
			_ = _542_v264
			_542_v264 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _524_v246.F16)
			var _543_v265 _dafny.Array
			_ = _543_v265
			var _nwElement0_18 _dafny.Map = _541_v263
			_ = _nwElement0_18
			var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(4))
			_ = _nw94
			(_nw94).ArraySet1(_nwElement0_18, 0)
			(_nw94).ArraySet1(_542_v264, 1)
			(_nw94).ArraySet1(_541_v263, 2)
			(_nw94).ArraySet1(_541_v263, 3)
			_543_v265 = _nw94
			var _544_v266 _dafny.Map
			_ = _544_v266
			_544_v266 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v223).F14(), _543_v265)
			var _545_v267 _dafny.Array
			_ = _545_v267
			var _nwElement0_19 bool = (_dafny.SetOf((_483_v220).Cardinality(), _175_v1, _494_v228.F16)).IsProperSubsetOf(_249_v55)
			_ = _nwElement0_19
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(10))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_19, 0)
			(_nw95).ArraySet1(((_486_v223).F13()).Cmp(_dafny.IntOfInt64(533)) > 0, 1)
			(_nw95).ArraySet1(((_484_v221).F17()) && ((_486_v223).F14()), 2)
			(_nw95).ArraySet1(_183_v6, 3)
			(_nw95).ArraySet1((_486_v223).F14(), 4)
			(_nw95).ArraySet1((_486_v223).F14(), 5)
			(_nw95).ArraySet1(_521_v243.F12(), 6)
			(_nw95).ArraySet1(!(_544_v266).Contains(!(_521_v243.F12())), 7)
			(_nw95).ArraySet1((_184_v7).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_183_v6, _175_v1, _494_v228.F16, _185_globalState), _dafny.IntOfUint32((_184_v7).Cardinality()))).Uint32()).(bool), 8)
			(_nw95).ArraySet1((_484_v221).F17(), 9)
			_545_v267 = _nw95
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_545_v267), 0))
			_ = _index62
			(_545_v267).ArraySet1((_486_v223).F14(), (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_545_v267), 0))
			_ = _index63
			var _rhs74 bool = (_dafny.Companion_Sequence_.IsPrefixOf(_176_v2, _176_v2)) || ((_484_v221).F17())
			_ = _rhs74
			var _rhs75 _dafny.Int = (_484_v221).F18()
			_ = _rhs75
			var _rhs76 bool = (_484_v221).F17()
			_ = _rhs76
			var _rhs77 _dafny.CodePoint = _177_v3
			_ = _rhs77
			var _lhs53 *GlobalState = _185_globalState
			_ = _lhs53
			var _lhs54 *GlobalState = _185_globalState
			_ = _lhs54
			var _lhs55 _dafny.Array = _545_v267
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_545_v267), 0))
			_ = _lhs56
			_lhs53.F7 = _rhs74
			_lhs54.F9 = _rhs75
			(_lhs55).ArraySet1(_rhs76, (_lhs56).Int())
			_177_v3 = _rhs77
			var _546_v268 _dafny.Int
			_ = _546_v268
			var _out14 _dafny.Int
			_ = _out14
			_out14 = (_494_v228).M6(_185_globalState)
			_546_v268 = _out14
			var _547_v269 _dafny.Map
			_ = _547_v269
			_547_v269 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_486_v223).F13())
			var _548_v270 _dafny.Map
			_ = _548_v270
			_548_v270 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v223).F14(), (_486_v223).F14())
			var _549_v271 _dafny.MultiSet
			_ = _549_v271
			_549_v271 = _dafny.MultiSetOf(_dafny.IntOfUint32((_176_v2).Cardinality()), (_484_v221).Fm12(_185_globalState), (_548_v270).Cardinality(), _494_v228.F16, (_484_v221).F18())
			var _550_v272 _dafny.Map
			_ = _550_v272
			_550_v272 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v5, true)
			var _551_v273 _dafny.Map
			_ = _551_v273
			_551_v273 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_547_v269).Contains((_521_v243).Fm2(_549_v271, (_dafny.Zero).Minus((_548_v270).Cardinality()), _547_v269, (_545_v267).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_545_v267), 0))).Int()).(bool), _185_globalState)) {
					return (_547_v269).Get((_521_v243).Fm2(_549_v271, (_dafny.Zero).Minus((_548_v270).Cardinality()), _547_v269, (_545_v267).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_545_v267), 0))).Int()).(bool), _185_globalState)).(_dafny.Int)
				}
				return (_550_v272).Cardinality()
			})(), _545_v267)
			var _552_v274 _dafny.Set
			_ = _552_v274
			_552_v274 = _dafny.SetOf(_177_v3)
			var _553_v275 _dafny.Map
			_ = _553_v275
			_553_v275 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_552_v274).Equals(_552_v274), (func() _dafny.Array {
				if (_486_v223).F14() {
					return _545_v267
				}
				return _545_v267
			})())
			var _rhs78 _dafny.Map = _551_v273
			_ = _rhs78
			var _rhs79 bool = (func() bool {
				if (_486_v223).F14() {
					return (_545_v267).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_545_v267), 0))).Int()).(bool)
				}
				return _521_v243.F12()
			})()
			_ = _rhs79
			var _rhs80 _dafny.Int = _dafny.IntOfUint32((_248_v54).Cardinality())
			_ = _rhs80
			var _rhs81 _dafny.Int = ((_484_v221).F18()).Plus((_dafny.Zero).Minus(_494_v228.F16))
			_ = _rhs81
			var _rhs82 _dafny.Map = (_553_v275).Merge(_553_v275)
			_ = _rhs82
			var _lhs57 *GlobalState = _185_globalState
			_ = _lhs57
			var _lhs58 *C1 = _494_v228
			_ = _lhs58
			var _lhs59 *C1 = _494_v228
			_ = _lhs59
			_551_v273 = _rhs78
			_lhs57.F3 = _rhs79
			_lhs58.F16 = _rhs80
			_lhs59.F16 = _rhs81
			_553_v275 = _rhs82
			_548_v270 = (_548_v270).Update((_484_v221).F17(), (_486_v223).F14())
		}
		_486_v223 = _486_v223
		var _554_v276 D1
		_ = _554_v276
		_554_v276 = Companion_D1_.Create_DC2_((_484_v221).F17(), (_486_v223).F14(), (_484_v221).F17(), _177_v3, _177_v3)
		var _pat_let_tv14 = _486_v223
		_ = _pat_let_tv14
		var _555_v277 _dafny.Array
		_ = _555_v277
		var _nwElement0_20 _dafny.CodePoint = Companion_Default___.Fm14(_554_v276, _524_v246.F16, _185_globalState)
		_ = _nwElement0_20
		var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(12))
		_ = _nw96
		(_nw96).ArraySet1CodePoint(_nwElement0_20, 0)
		(_nw96).ArraySet1CodePoint(_177_v3, 1)
		(_nw96).ArraySet1CodePoint(_177_v3, 2)
		(_nw96).ArraySet1CodePoint(_177_v3, 3)
		(_nw96).ArraySet1CodePoint(Companion_Default___.Fm14(func(_pat_let17_0 D1) D1 {
			return func(_556_dt__update__tmp_h3 D1) D1 {
				return func(_pat_let18_0 bool) D1 {
					return func(_557_dt__update_hcf3_h0 bool) D1 {
						return Companion_D1_.Create_DC2_((_556_dt__update__tmp_h3).Dtor_cf2(), _557_dt__update_hcf3_h0, (_556_dt__update__tmp_h3).Dtor_cf4(), (_556_dt__update__tmp_h3).Dtor_cf5(), (_556_dt__update__tmp_h3).Dtor_cf6())
					}(_pat_let18_0)
				}((_pat_let_tv14).F14())
			}(_pat_let17_0)
		}(_554_v276), (_484_v221).F18(), _185_globalState), 4)
		(_nw96).ArraySet1CodePoint(_177_v3, 5)
		(_nw96).ArraySet1CodePoint(_177_v3, 6)
		(_nw96).ArraySet1CodePoint(_177_v3, 7)
		(_nw96).ArraySet1CodePoint(_177_v3, 8)
		(_nw96).ArraySet1CodePoint(_177_v3, 9)
		(_nw96).ArraySet1CodePoint((func() _dafny.CodePoint {
			if !(!((_486_v223).F14())) {
				return (_248_v54).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v223).Fm6(_183_v6, (_486_v223).F13(), _185_globalState), _175_v1)).Cardinality(), _dafny.IntOfUint32((_248_v54).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
			return _177_v3
		})(), 10)
		(_nw96).ArraySet1CodePoint(_177_v3, 11)
		_555_v277 = _nw96
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_555_v277), 0))
		_ = _index64
		(_555_v277).ArraySet1CodePoint(Companion_Default___.Fm14(_554_v276, (_484_v221).F18(), _185_globalState), (_index64).Int())
		var _558_v278 *C2
		_ = _558_v278
		var _nw97 *C2 = New_C2_()
		_ = _nw97
		_nw97.Ctor__((_520_v242).Equals(_520_v242), _521_v243.F11(), _521_v243.F12())
		_558_v278 = _nw97
		var _559_v279 _dafny.MultiSet
		_ = _559_v279
		_559_v279 = _dafny.MultiSetOf(_dafny.IntOfUint32((_184_v7).Cardinality()), _dafny.IntOfUint32((_176_v2).Cardinality()), _494_v228.F16, _524_v246.F16)
		var _560_v280 _dafny.Map
		_ = _560_v280
		_560_v280 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-467), _559_v279)
		var _561_v281 _dafny.Array
		_ = _561_v281
		var _nwElement0_21 _dafny.MultiSet = (_559_v279).Update(_494_v228.F16, Companion_Default___.Abs(_175_v1))
		_ = _nwElement0_21
		var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(12))
		_ = _nw98
		(_nw98).ArraySet1(_nwElement0_21, 0)
		(_nw98).ArraySet1((_559_v279).Intersection(_559_v279), 1)
		(_nw98).ArraySet1(_559_v279, 2)
		(_nw98).ArraySet1(_559_v279, 3)
		(_nw98).ArraySet1(_559_v279, 4)
		(_nw98).ArraySet1(_dafny.MultiSetFromSeq(_176_v2), 5)
		(_nw98).ArraySet1((func() _dafny.MultiSet {
			if (_560_v280).Contains((_484_v221).F18()) {
				return (_560_v280).Get((_484_v221).F18()).(_dafny.MultiSet)
			}
			return _559_v279
		})(), 6)
		(_nw98).ArraySet1(_559_v279, 7)
		(_nw98).ArraySet1(_559_v279, 8)
		(_nw98).ArraySet1((Companion_Default___.Fm13(_177_v3, (_484_v221).F17(), _185_globalState)).Update((_484_v221).F18(), Companion_Default___.Abs((_dafny.Zero).Minus(_dafny.IntOfUint32((_248_v54).Cardinality())))), 9)
		(_nw98).ArraySet1((_dafny.MultiSetOf(_524_v246.F16)).Union(_559_v279), 10)
		(_nw98).ArraySet1(_559_v279, 11)
		_561_v281 = _nw98
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(81), _dafny.ArrayLen((_561_v281), 0))
		_ = _index65
		(_561_v281).ArraySet1(_559_v279, (_index65).Int())
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_555_v277), 0))
		_ = _index66
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(81), _dafny.ArrayLen((_561_v281), 0))
		_ = _index67
		var _rhs83 _dafny.Int = _494_v228.F16
		_ = _rhs83
		var _rhs84 _dafny.CodePoint = _177_v3
		_ = _rhs84
		var _rhs85 *C2 = _558_v278
		_ = _rhs85
		var _rhs86 _dafny.MultiSet = _559_v279
		_ = _rhs86
		var _lhs60 *GlobalState = _185_globalState
		_ = _lhs60
		var _lhs61 _dafny.Array = _555_v277
		_ = _lhs61
		var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_555_v277), 0))
		_ = _lhs62
		var _lhs63 _dafny.Array = _561_v281
		_ = _lhs63
		var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(81), _dafny.ArrayLen((_561_v281), 0))
		_ = _lhs64
		_lhs60.F9 = _rhs83
		(_lhs61).ArraySet1CodePoint(_rhs84, (_lhs62).Int())
		_558_v278 = _rhs85
		(_lhs63).ArraySet1(_rhs86, (_lhs64).Int())
	} else {
		var _562_v282 T0
		_ = _562_v282
		var _nw99 *C1 = New_C1_()
		_ = _nw99
		_nw99.Ctor__((_dafny.Zero).Minus((_485_v222).Cardinality()), _485_v222, true)
		_562_v282 = _nw99
		var _563_v283 _dafny.Map
		_ = _563_v283
		_563_v283 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_562_v282, (func() _dafny.Int {
			if (_484_v221).F17() {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_251_v57).Select((Companion_Default___.SafeIndex((_249_v55).Cardinality(), _dafny.IntOfUint32((_251_v57).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_484_v221).F18(), _dafny.IntOfUint32(((_251_v57).Select((Companion_Default___.SafeIndex((_249_v55).Cardinality(), _dafny.IntOfUint32((_251_v57).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _177_v3)).Cardinality())
			}
			return _dafny.IntOfUint32((_184_v7).Cardinality())
		})())
		_563_v283 = (_563_v283).Update(_562_v282, Companion_Default___.Fm1((_184_v7).Select((Companion_Default___.SafeIndex((_484_v221).F18(), _dafny.IntOfUint32((_184_v7).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus((_484_v221).F18()), (_486_v223).F13(), _185_globalState))
		var _564_v284 _dafny.Map
		_ = _564_v284
		_564_v284 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_486_v223).F14()), false)
		var _565_v285 _dafny.Sequence
		_ = _565_v285
		_565_v285 = _dafny.SeqOf(_564_v284, (_564_v284).Update(_562_v282.F12(), !((_486_v223).Fm6((_486_v223).F14(), (func() _dafny.Int {
			if (_485_v222).Contains((_486_v223).F14()) {
				return (_485_v222).Multiplicity((_486_v223).F14())
			}
			return _175_v1
		})(), _185_globalState))), _564_v284)
		_565_v285 = _dafny.SeqOf((_564_v284).Merge((_564_v284).Update((_484_v221).F17(), _183_v6)), _564_v284, _564_v284, (_564_v284).Merge(_564_v284))
		var _566_v286 _dafny.Array
		_ = _566_v286
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_11
		var _nw100 _dafny.Array
		_ = _nw100
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw100 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = func(_567_i27 _dafny.Int) bool {
				return true
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw100 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw100).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw100).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_566_v286 = _nw100
		var _568_v287 _dafny.Map
		_ = _568_v287
		_568_v287 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _566_v286)
		var _569_v288 _dafny.Array
		_ = _569_v288
		var _nwElement0_22 _dafny.Array = _566_v286
		_ = _nwElement0_22
		var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(22))
		_ = _nw101
		(_nw101).ArraySet1(_nwElement0_22, 0)
		(_nw101).ArraySet1(_566_v286, 1)
		(_nw101).ArraySet1(_566_v286, 2)
		(_nw101).ArraySet1(_566_v286, 3)
		(_nw101).ArraySet1(_566_v286, 4)
		(_nw101).ArraySet1(_566_v286, 5)
		(_nw101).ArraySet1(_566_v286, 6)
		(_nw101).ArraySet1(_566_v286, 7)
		(_nw101).ArraySet1(_566_v286, 8)
		(_nw101).ArraySet1(_566_v286, 9)
		(_nw101).ArraySet1(_566_v286, 10)
		(_nw101).ArraySet1(_566_v286, 11)
		(_nw101).ArraySet1(_566_v286, 12)
		(_nw101).ArraySet1(_566_v286, 13)
		(_nw101).ArraySet1(_566_v286, 14)
		(_nw101).ArraySet1((func() _dafny.Array {
			if (_568_v287).Contains(_183_v6) {
				return (_568_v287).Get(_183_v6).(_dafny.Array)
			}
			return _566_v286
		})(), 15)
		(_nw101).ArraySet1(_566_v286, 16)
		(_nw101).ArraySet1(_566_v286, 17)
		(_nw101).ArraySet1(_566_v286, 18)
		(_nw101).ArraySet1(_566_v286, 19)
		(_nw101).ArraySet1(_566_v286, 20)
		(_nw101).ArraySet1(_566_v286, 21)
		_569_v288 = _nw101
		var _570_v289 D8
		_ = _570_v289
		_570_v289 = Companion_D8_.Create_DC21_((_484_v221).F17(), _569_v288, (_484_v221).F18(), _566_v286, Companion_Default___.SafeModuloInt((_486_v223).F13(), (_484_v221).F18()))
		var _source5 D8 = _570_v289
		_ = _source5
		if _source5.Is_DC20() {
			var _571___mcc_h18 bool = _source5.Get_().(D8_DC20).Cf46
			_ = _571___mcc_h18
			var _572___mcc_h19 bool = _source5.Get_().(D8_DC20).Cf47
			_ = _572___mcc_h19
			var _573___mcc_h20 bool = _source5.Get_().(D8_DC20).Cf48
			_ = _573___mcc_h20
			var _574_cf48 bool = _573___mcc_h20
			_ = _574_cf48
			var _575_cf47 bool = _572___mcc_h19
			_ = _575_cf47
			var _576_cf46 bool = _571___mcc_h18
			_ = _576_cf46
			var _577_v290 _dafny.Sequence
			_ = _577_v290
			_577_v290 = _dafny.SeqOf(_483_v220)
			_175_v1 = (_dafny.Zero).Minus((((_577_v290).Select((Companion_Default___.SafeIndex(_494_v228.F16, _dafny.IntOfUint32((_577_v290).Cardinality()))).Uint32()).(_dafny.Set)).Difference((_483_v220).Intersection(_483_v220))).Cardinality())
			var _578_v291 _dafny.Map
			_ = _578_v291
			_578_v291 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(803), _177_v3)
			_177_v3 = (func() _dafny.CodePoint {
				if (_578_v291).Contains(_494_v228.F16) {
					return (_578_v291).Get(_494_v228.F16).(_dafny.CodePoint)
				}
				return _177_v3
			})()
			var _579_v292 _dafny.MultiSet
			_ = _579_v292
			_579_v292 = _dafny.MultiSetOf(_177_v3, _177_v3, _177_v3, (_248_v54).Select((Companion_Default___.SafeIndex((_484_v221).F18(), _dafny.IntOfUint32((_248_v54).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _580_v293 _dafny.MultiSet
			_ = _580_v293
			_580_v293 = _dafny.MultiSetOf(_177_v3)
			var _581_v294 _dafny.Sequence
			_ = _581_v294
			_581_v294 = _dafny.SeqOf(_579_v292)
			_575_cf47 = ((_581_v294).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_248_v54).Cardinality()), _dafny.IntOfUint32((_581_v294).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsSubsetOf((func() _dafny.MultiSet {
				if true {
					return _579_v292
				}
				return (_580_v293)
			})())
			(_185_globalState).F3 = (_486_v223).F14()
		} else if _source5.Is_DC21() {
			var _582___mcc_h21 bool = _source5.Get_().(D8_DC21).Cf49
			_ = _582___mcc_h21
			var _583___mcc_h22 _dafny.Array = _source5.Get_().(D8_DC21).Cf50
			_ = _583___mcc_h22
			var _584___mcc_h23 _dafny.Int = _source5.Get_().(D8_DC21).Cf51
			_ = _584___mcc_h23
			var _585___mcc_h24 _dafny.Array = _source5.Get_().(D8_DC21).Cf52
			_ = _585___mcc_h24
			var _586___mcc_h25 _dafny.Int = _source5.Get_().(D8_DC21).Cf53
			_ = _586___mcc_h25
			var _587_cf53 _dafny.Int = _586___mcc_h25
			_ = _587_cf53
			var _588_cf52 _dafny.Array = _585___mcc_h24
			_ = _588_cf52
			var _589_cf51 _dafny.Int = _584___mcc_h23
			_ = _589_cf51
			var _590_cf50 _dafny.Array = _583___mcc_h22
			_ = _590_cf50
			var _591_cf49 bool = _582___mcc_h21
			_ = _591_cf49
			_589_cf51 = _dafny.IntOfUint32(((_251_v57).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if _562_v282.F12() {
					return _175_v1
				}
				return (_484_v221).F18()
			})(), _dafny.IntOfUint32((_251_v57).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			(_562_v282).F12_set_(_591_cf49)
			_588_cf52 = _566_v286
			var _592_v295 _dafny.MultiSet
			_ = _592_v295
			_592_v295 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(235))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}((func(_593_v228 *C1) func(_dafny.Int) _dafny.Int {
				return func(_594_i28 _dafny.Int) _dafny.Int {
					return _593_v228.F16
				}
			})(_494_v228))))
			_589_cf51 = (func() _dafny.Int {
				if (_592_v295).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(899))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_595_i29 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfInt64(-198))
				}))) {
					return (_592_v295).Multiplicity(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(899))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg32 _dafny.Int) interface{} {
							return coer32(arg32)
						}
					}(func(_596_i29 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_dafny.IntOfInt64(-198))
					})))
				}
				return Companion_Default___.SafeDivisionInt((_484_v221).F18(), (_484_v221).F18())
			})()
		} else if _source5.Is_DC19() {
			var _597___mcc_h26 _dafny.MultiSet = _source5.Get_().(D8_DC19).Cf45
			_ = _597___mcc_h26
			var _598_cf45 _dafny.MultiSet = _597___mcc_h26
			_ = _598_cf45
			Companion_Default___.M0(_185_globalState)
			(_185_globalState).F9 = (func() _dafny.Int {
				if _562_v282.F12() {
					return (_486_v223).F13()
				}
				return (_486_v223).F13()
			})()
			(_185_globalState).F9 = ((_484_v221).F18()).Times(_494_v228.F16)
			_562_v282 = _562_v282
		} else {
			var _599___mcc_h27 D8 = _source5.Get_().(D8_DC22).Cf54
			_ = _599___mcc_h27
			var _600_cf54 D8 = _599___mcc_h27
			_ = _600_cf54
			(_562_v282).F12_set_(_562_v282.F12())
			var _601_v296 D1
			_ = _601_v296
			_601_v296 = Companion_D1_.Create_DC2_(_562_v282.F12(), (_484_v221).F17(), (_486_v223).F14(), _177_v3, _177_v3)
			_177_v3 = Companion_Default___.Fm14(_601_v296, _175_v1, _185_globalState)
			var _602_v297 _dafny.Map
			_ = _602_v297
			_602_v297 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v223).F14(), (_484_v221).F18())
			var _603_v298 _dafny.Array
			_ = _603_v298
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw102
			_603_v298 = _nw102
			var _604_v299 _dafny.Map
			_ = _604_v299
			_604_v299 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_603_v298, _602_v297)
			var _605_v300 _dafny.Sequence
			_ = _605_v300
			_605_v300 = _dafny.SeqOf(_602_v297, _602_v297)
			var _606_v301 _dafny.Map
			_ = _606_v301
			_606_v301 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_248_v54).Cardinality()), (_484_v221).F18())
			var _607_v302 _dafny.MultiSet
			_ = _607_v302
			_607_v302 = _dafny.MultiSetOf(_494_v228.F16, (_486_v223).F13())
			var _608_v303 _dafny.Map
			_ = _608_v303
			_608_v303 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((func() _dafny.Int {
				if (_606_v301).Contains((_607_v302).Cardinality()) {
					return (_606_v301).Get((_607_v302).Cardinality()).(_dafny.Int)
				}
				return _494_v228.F16
			})(), _175_v1), (_484_v221).F18())
			var _609_v304 _dafny.Array
			_ = _609_v304
			var _nwElement0_23 _dafny.Map = _602_v297
			_ = _nwElement0_23
			var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(19))
			_ = _nw103
			(_nw103).ArraySet1(_nwElement0_23, 0)
			(_nw103).ArraySet1(_602_v297, 1)
			(_nw103).ArraySet1((func() _dafny.Map {
				if (_604_v299).Contains(_603_v298) {
					return (_604_v299).Get(_603_v298).(_dafny.Map)
				}
				return _602_v297
			})(), 2)
			(_nw103).ArraySet1(_602_v297, 3)
			(_nw103).ArraySet1(Companion_Default___.Fm23(_494_v228.F16, _185_globalState), 4)
			(_nw103).ArraySet1(_602_v297, 5)
			(_nw103).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_484_v221).F17(), (_486_v223).F13()), 6)
			(_nw103).ArraySet1(_602_v297, 7)
			(_nw103).ArraySet1((_602_v297).Merge(_602_v297), 8)
			(_nw103).ArraySet1((_602_v297).Merge(_602_v297), 9)
			(_nw103).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_484_v221).F17(), (_484_v221).F18()), 10)
			(_nw103).ArraySet1(_602_v297, 11)
			(_nw103).ArraySet1(_602_v297, 12)
			(_nw103).ArraySet1(_602_v297, 13)
			(_nw103).ArraySet1(_602_v297, 14)
			(_nw103).ArraySet1((_602_v297).Merge(_602_v297), 15)
			(_nw103).ArraySet1(_602_v297, 16)
			(_nw103).ArraySet1(_602_v297, 17)
			(_nw103).ArraySet1((_602_v297).Merge((_605_v300).Select((Companion_Default___.SafeIndex((_608_v303).Cardinality(), _dafny.IntOfUint32((_605_v300).Cardinality()))).Uint32()).(_dafny.Map)), 18)
			_609_v304 = _nw103
			var _610_v305 _dafny.Array
			_ = _610_v305
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_12
			var _nw104 _dafny.Array
			_ = _nw104
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw104 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) D9 = (func(_611_v1 _dafny.Int) func(_dafny.Int) D9 {
					return func(_612_i30 _dafny.Int) D9 {
						return Companion_D9_.Create_DC26_(_611_v1)
					}
				})(_175_v1)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw104 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw104).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw104).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_610_v305 = _nw104
			var _rhs87 _dafny.Int = (_dafny.MultiSetOf(_610_v305)).Cardinality()
			_ = _rhs87
			var _rhs88 *C1 = _494_v228
			_ = _rhs88
			var _rhs89 _dafny.Int = (_484_v221).F18()
			_ = _rhs89
			var _rhs90 bool = (_486_v223).F14()
			_ = _rhs90
			var _rhs91 _dafny.Array = _609_v304
			_ = _rhs91
			var _lhs65 *C1 = _494_v228
			_ = _lhs65
			var _lhs66 *C1 = _494_v228
			_ = _lhs66
			var _lhs67 *GlobalState = _185_globalState
			_ = _lhs67
			_lhs65.F16 = _rhs87
			_494_v228 = _rhs88
			_lhs66.F16 = _rhs89
			_lhs67.F8 = _rhs90
			_609_v304 = _rhs91
			var _613_v306 _dafny.Map
			_ = _613_v306
			_613_v306 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _494_v228)
			var _614_v307 _dafny.Array
			_ = _614_v307
			var _nwElement0_24 *C1 = (func() *C1 {
				if (_613_v306).Contains(_183_v6) {
					return (_613_v306).Get(_183_v6).(*C1)
				}
				return _494_v228
			})()
			_ = _nwElement0_24
			var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(4))
			_ = _nw105
			(_nw105).ArraySet1(_nwElement0_24, 0)
			(_nw105).ArraySet1(_494_v228, 1)
			(_nw105).ArraySet1(_494_v228, 2)
			(_nw105).ArraySet1(_494_v228, 3)
			_614_v307 = _nw105
			_614_v307 = _614_v307
		}
		Companion_Default___.M0(_185_globalState)
		_566_v286 = _566_v286
	}
	(_185_globalState).F7 = _183_v6
	var _615_i31 _dafny.Int
	_ = _615_i31
	_615_i31 = _dafny.Zero
	{
		for (_484_v221).F17() {
			{
				if (_615_i31).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L5
				}
				_615_i31 = (_615_i31).Plus(_dafny.One)
				(_494_v228).F16 = _175_v1
				var _616_v308 *C2
				_ = _616_v308
				var _nw106 *C2 = New_C2_()
				_ = _nw106
				_nw106.Ctor__((_484_v221).F17(), _485_v222, (_486_v223).F14())
				_616_v308 = _nw106
				var _617_v309 D12
				_ = _617_v309
				_617_v309 = Companion_D12_.Create_DC32_(_dafny.IntOfUint32((_248_v54).Cardinality()), _616_v308, _484_v221, (_486_v223).F13())
				var _618_v310 _dafny.Map
				_ = _618_v310
				_618_v310 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v2, _617_v309)
				var _619_v311 _dafny.Map
				_ = _619_v311
				_619_v311 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_v228.F16, (_dafny.SetOf(_616_v308.F15)).Cardinality())
				var _620_v312 _dafny.Map
				_ = _620_v312
				_620_v312 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v223).F13(), (_619_v311).Cardinality())
				_618_v310 = (_618_v310).Update(_176_v2, Companion_D12_.Create_DC32_(_494_v228.F16, _616_v308, _484_v221, (func() _dafny.Int {
					if (_619_v311).Contains(_dafny.IntOfInt64(-252)) {
						return (_619_v311).Get(_dafny.IntOfInt64(-252)).(_dafny.Int)
					}
					return (_484_v221).F18()
				})()))
				var _621_v313 _dafny.Sequence
				_ = _621_v313
				_621_v313 = _dafny.SeqOf(_484_v221)
				var _622_v314 _dafny.MultiSet
				_ = _622_v314
				_622_v314 = _dafny.MultiSetOf(_dafny.IntOfUint32((_621_v313).Cardinality()))
				var _623_v315 D9
				_ = _623_v315
				_623_v315 = Companion_D9_.Create_DC25_(_248_v54, _616_v308.F15, (_484_v221).F18())
				var _624_v316 _dafny.Map
				_ = _624_v316
				_624_v316 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_494_v228).Fm2(_622_v314, _494_v228.F16, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_623_v315).Dtor_cf60(), _494_v228.F16), (_486_v223).F14(), _185_globalState), _dafny.IntOfInt64(373))
				var _625_v317 *C4
				_ = _625_v317
				var _nw107 *C4 = New_C4_()
				_ = _nw107
				_nw107.Ctor__(_485_v222, (_484_v221).F17())
				_625_v317 = _nw107
				var _626_v318 _dafny.Map
				_ = _626_v318
				_626_v318 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_624_v316).Contains(_616_v308.F15) {
						return (_624_v316).Get(_616_v308.F15).(_dafny.Int)
					}
					return _dafny.IntOfInt64(312)
				})(), _625_v317)
				_626_v318 = (_626_v318).Update(_dafny.IntOfInt64(615), (func() *C4 {
					if (_486_v223).F14() {
						return _625_v317
					}
					return _625_v317
				})())
				var _627_v319 _dafny.Int
				_ = _627_v319
				var _628_v320 bool
				_ = _628_v320
				var _out15 _dafny.Int
				_ = _out15
				var _out16 bool
				_ = _out16
				_out15, _out16 = (_486_v223).M2(_248_v54, (_486_v223).F14(), _490_v227, _185_globalState)
				_627_v319 = _out15
				_628_v320 = _out16
				goto C5
			}
		C5:
		}
		goto L5
	}
L5:
	_dafny.Print(_175_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_176_v2, _dafny.SeqOf(_dafny.IntOfInt64(-351), _dafny.IntOfInt64(231), _dafny.IntOfInt64(-351), _dafny.IntOfInt64(-351), _dafny.IntOfInt64(231), _dafny.IntOfInt64(-351))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_177_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false).UpdateUnsafe(_dafny.IntOfInt64(582), false), _dafny.CodePoint('p'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v5).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_183_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_184_v7, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_185_globalState).F0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false).UpdateUnsafe(_dafny.IntOfInt64(582), false), _dafny.CodePoint('p'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F1).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_185_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState.F4).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_185_globalState.F6, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_185_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_185_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_185_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_186_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_205_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_248_v54.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_249_v55).Equals(_dafny.SetOf(_dafny.IntOfInt64(-351))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v56).Dtor_cf40())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v56).Dtor_cf41().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v56).Dtor_cf42())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_250_v56).Dtor_cf43())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_250_v56).Dtor_cf44()).Equals(_dafny.SetOf(_dafny.IntOfInt64(-351))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_251_v57, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("eyxck"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_369_v140).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mupivve"), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_483_v220).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_484_v221).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_484_v221).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_485_v222).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_486_v223).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_486_v223).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_487_v224).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_488_v225, _dafny.SeqOf(_dafny.MultiSetOf(false, false, false), _dafny.MultiSetOf(false, false, false), _dafny.MultiSetOf(false, false, false), _dafny.MultiSetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_489_v226).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_490_v227).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_494_v228.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_495_v229).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_496_i23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_615_i31)
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

type D1_DC2 struct {
	Cf2 bool
	Cf3 bool
	Cf4 bool
	Cf5 _dafny.CodePoint
	Cf6 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 bool, Cf3 bool, Cf4 bool, Cf5 _dafny.CodePoint, Cf6 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4, Cf5, Cf6}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Set
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Set) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

type D1_DC3 struct {
	Cf7 D1
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 D1) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(false, false, false, _dafny.CodePoint('D'), _dafny.CodePoint('D'))
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf5() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf6
}

func (_this D1) Dtor_cf1() _dafny.Set {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) Dtor_cf7() D1 {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
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

type D2_DC5 struct {
	Cf9  _dafny.Int
	Cf10 _dafny.Int
	Cf11 _dafny.Int
	Cf12 _dafny.Int
	Cf13 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Int, Cf10 _dafny.Int, Cf11 _dafny.Int, Cf12 _dafny.Int, Cf13 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf9, Cf10, Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC4 struct {
	Cf8 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf8 _dafny.Int) D2 {
	return D2{D2_DC4{Cf8}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf13
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13.Equals(data2.Cf13)
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0
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

type D3_DC7 struct {
	Cf15 _dafny.Int
	Cf16 _dafny.Int
	Cf17 *C0
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf15 _dafny.Int, Cf16 _dafny.Int, Cf17 *C0) D3 {
	return D3{D3_DC7{Cf15, Cf16, Cf17}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC8 struct {
	Cf18 _dafny.Int
	Cf19 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf18 _dafny.Int, Cf19 _dafny.Int) D3 {
	return D3{D3_DC8{Cf18, Cf19}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC9 struct {
	Cf20 _dafny.Int
	Cf21 _dafny.Int
	Cf22 *C0
	Cf23 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf20 _dafny.Int, Cf21 _dafny.Int, Cf22 *C0, Cf23 _dafny.Int) D3 {
	return D3{D3_DC9{Cf20, Cf21, Cf22, Cf23}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC6 struct {
	Cf14 _dafny.Array
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf14 _dafny.Array) D3 {
	return D3{D3_DC6{Cf14}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_(_dafny.Zero, _dafny.Zero, (*C0)(nil))
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC7).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC7).Cf16
}

func (_this D3) Dtor_cf17() *C0 {
	return _this.Get_().(D3_DC7).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf20
}

func (_this D3) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf21
}

func (_this D3) Dtor_cf22() *C0 {
	return _this.Get_().(D3_DC9).Cf22
}

func (_this D3) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf23
}

func (_this D3) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D3_DC6).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Cmp(data2.Cf19) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
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

type D4_DC11 struct {
	Cf25 bool
	Cf26 _dafny.Int
	Cf27 _dafny.Int
	Cf28 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf25 bool, Cf26 _dafny.Int, Cf27 _dafny.Int, Cf28 _dafny.Int) D4 {
	return D4{D4_DC11{Cf25, Cf26, Cf27, Cf28}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf24 _dafny.Array
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf24 _dafny.Array) D4 {
	return D4{D4_DC10{Cf24}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(false, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf25() bool {
	return _this.Get_().(D4_DC11).Cf25
}

func (_this D4) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf26
}

func (_this D4) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf27
}

func (_this D4) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf28
}

func (_this D4) Dtor_cf24() _dafny.Array {
	return _this.Get_().(D4_DC10).Cf24
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28.Cmp(data2.Cf28) == 0
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf24 == data2.Cf24
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
	Cf30 _dafny.Int
	Cf31 bool
	Cf32 _dafny.Int
	Cf33 _dafny.Array
	Cf34 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf30 _dafny.Int, Cf31 bool, Cf32 _dafny.Int, Cf33 _dafny.Array, Cf34 bool) D5 {
	return D5{D5_DC13{Cf30, Cf31, Cf32, Cf33, Cf34}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_() D5 {
	return D5{D5_DC14{}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf29 _dafny.Sequence
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf29 _dafny.Sequence) D5 {
	return D5{D5_DC12{Cf29}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(_dafny.Zero, false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D5) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf30
}

func (_this D5) Dtor_cf31() bool {
	return _this.Get_().(D5_DC13).Cf31
}

func (_this D5) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf32
}

func (_this D5) Dtor_cf33() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf33
}

func (_this D5) Dtor_cf34() bool {
	return _this.Get_().(D5_DC13).Cf34
}

func (_this D5) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf29
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33 == data2.Cf33 && data1.Cf34 == data2.Cf34
		}
	case D5_DC14:
		{
			_, ok := other.Get_().(D5_DC14)
			return ok
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
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

type D6_DC16 struct {
	Cf36 _dafny.Int
	Cf37 bool
	Cf38 bool
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf36 _dafny.Int, Cf37 bool, Cf38 bool) D6 {
	return D6{D6_DC16{Cf36, Cf37, Cf38}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf35 _dafny.MultiSet
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf35 _dafny.MultiSet) D6 {
	return D6{D6_DC15{Cf35}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.Zero, false, false)
}

func (_this D6) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf36
}

func (_this D6) Dtor_cf37() bool {
	return _this.Get_().(D6_DC16).Cf37
}

func (_this D6) Dtor_cf38() bool {
	return _this.Get_().(D6_DC16).Cf38
}

func (_this D6) Dtor_cf35() _dafny.MultiSet {
	return _this.Get_().(D6_DC15).Cf35
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf35.Equals(data2.Cf35)
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
	Cf40 _dafny.Int
	Cf41 _dafny.Sequence
	Cf42 bool
	Cf43 bool
	Cf44 _dafny.Set
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf40 _dafny.Int, Cf41 _dafny.Sequence, Cf42 bool, Cf43 bool, Cf44 _dafny.Set) D7 {
	return D7{D7_DC18{Cf40, Cf41, Cf42, Cf43, Cf44}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC17 struct {
	Cf39 _dafny.Map
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf39 _dafny.Map) D7 {
	return D7{D7_DC17{Cf39}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_(_dafny.Zero, _dafny.EmptySeq, false, false, _dafny.EmptySet)
}

func (_this D7) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D7_DC18).Cf40
}

func (_this D7) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D7_DC18).Cf41
}

func (_this D7) Dtor_cf42() bool {
	return _this.Get_().(D7_DC18).Cf42
}

func (_this D7) Dtor_cf43() bool {
	return _this.Get_().(D7_DC18).Cf43
}

func (_this D7) Dtor_cf44() _dafny.Set {
	return _this.Get_().(D7_DC18).Cf44
}

func (_this D7) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D7_DC17).Cf39
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf40) + ", " + data.Cf41.VerbatimString(true) + ", " + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Equals(data2.Cf41) && data1.Cf42 == data2.Cf42 && data1.Cf43 == data2.Cf43 && data1.Cf44.Equals(data2.Cf44)
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
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

type D8_DC20 struct {
	Cf46 bool
	Cf47 bool
	Cf48 bool
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf46 bool, Cf47 bool, Cf48 bool) D8 {
	return D8{D8_DC20{Cf46, Cf47, Cf48}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC21 struct {
	Cf49 bool
	Cf50 _dafny.Array
	Cf51 _dafny.Int
	Cf52 _dafny.Array
	Cf53 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf49 bool, Cf50 _dafny.Array, Cf51 _dafny.Int, Cf52 _dafny.Array, Cf53 _dafny.Int) D8 {
	return D8{D8_DC21{Cf49, Cf50, Cf51, Cf52, Cf53}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC19 struct {
	Cf45 _dafny.MultiSet
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf45 _dafny.MultiSet) D8 {
	return D8{D8_DC19{Cf45}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC22 struct {
	Cf54 D8
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf54 D8) D8 {
	return D8{D8_DC22{Cf54}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_(false, false, false)
}

func (_this D8) Dtor_cf46() bool {
	return _this.Get_().(D8_DC20).Cf46
}

func (_this D8) Dtor_cf47() bool {
	return _this.Get_().(D8_DC20).Cf47
}

func (_this D8) Dtor_cf48() bool {
	return _this.Get_().(D8_DC20).Cf48
}

func (_this D8) Dtor_cf49() bool {
	return _this.Get_().(D8_DC21).Cf49
}

func (_this D8) Dtor_cf50() _dafny.Array {
	return _this.Get_().(D8_DC21).Cf50
}

func (_this D8) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf51
}

func (_this D8) Dtor_cf52() _dafny.Array {
	return _this.Get_().(D8_DC21).Cf52
}

func (_this D8) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf53
}

func (_this D8) Dtor_cf45() _dafny.MultiSet {
	return _this.Get_().(D8_DC19).Cf45
}

func (_this D8) Dtor_cf54() D8 {
	return _this.Get_().(D8_DC22).Cf54
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf54) + ")"
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
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50 == data2.Cf50 && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52 && data1.Cf53.Cmp(data2.Cf53) == 0
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D9_DC24 struct {
	Cf56 _dafny.Sequence
	Cf57 bool
	Cf58 _dafny.Int
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf56 _dafny.Sequence, Cf57 bool, Cf58 _dafny.Int) D9 {
	return D9{D9_DC24{Cf56, Cf57, Cf58}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC25 struct {
	Cf59 _dafny.Sequence
	Cf60 bool
	Cf61 _dafny.Int
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf59 _dafny.Sequence, Cf60 bool, Cf61 _dafny.Int) D9 {
	return D9{D9_DC25{Cf59, Cf60, Cf61}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf62 _dafny.Int
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf62 _dafny.Int) D9 {
	return D9{D9_DC26{Cf62}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC23 struct {
	Cf55 _dafny.Map
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf55 _dafny.Map) D9 {
	return D9{D9_DC23{Cf55}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(_dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D9) Dtor_cf56() _dafny.Sequence {
	return _this.Get_().(D9_DC24).Cf56
}

func (_this D9) Dtor_cf57() bool {
	return _this.Get_().(D9_DC24).Cf57
}

func (_this D9) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf58
}

func (_this D9) Dtor_cf59() _dafny.Sequence {
	return _this.Get_().(D9_DC25).Cf59
}

func (_this D9) Dtor_cf60() bool {
	return _this.Get_().(D9_DC25).Cf60
}

func (_this D9) Dtor_cf61() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf61
}

func (_this D9) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf62
}

func (_this D9) Dtor_cf55() _dafny.Map {
	return _this.Get_().(D9_DC23).Cf55
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + data.Cf56.VerbatimString(true) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + data.Cf59.VerbatimString(true) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf62) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf56.Equals(data2.Cf56) && data1.Cf57 == data2.Cf57 && data1.Cf58.Cmp(data2.Cf58) == 0
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf59.Equals(data2.Cf59) && data1.Cf60 == data2.Cf60 && data1.Cf61.Cmp(data2.Cf61) == 0
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D10_DC28 struct {
	Cf64 *C2
	Cf65 bool
	Cf66 bool
	Cf67 bool
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf64 *C2, Cf65 bool, Cf66 bool, Cf67 bool) D10 {
	return D10{D10_DC28{Cf64, Cf65, Cf66, Cf67}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC29 struct {
	Cf68 bool
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf68 bool) D10 {
	return D10{D10_DC29{Cf68}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

type D10_DC27 struct {
	Cf63 _dafny.Map
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf63 _dafny.Map) D10 {
	return D10{D10_DC27{Cf63}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC28_((*C2)(nil), false, false, false)
}

func (_this D10) Dtor_cf64() *C2 {
	return _this.Get_().(D10_DC28).Cf64
}

func (_this D10) Dtor_cf65() bool {
	return _this.Get_().(D10_DC28).Cf65
}

func (_this D10) Dtor_cf66() bool {
	return _this.Get_().(D10_DC28).Cf66
}

func (_this D10) Dtor_cf67() bool {
	return _this.Get_().(D10_DC28).Cf67
}

func (_this D10) Dtor_cf68() bool {
	return _this.Get_().(D10_DC29).Cf68
}

func (_this D10) Dtor_cf63() _dafny.Map {
	return _this.Get_().(D10_DC27).Cf63
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf63) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf64 == data2.Cf64 && data1.Cf65 == data2.Cf65 && data1.Cf66 == data2.Cf66 && data1.Cf67 == data2.Cf67
		}
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf68 == data2.Cf68
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
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

type D11_DC30 struct {
	Cf69 _dafny.Map
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf69 _dafny.Map) D11 {
	return D11{D11_DC30{Cf69}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D11) Dtor_cf69() _dafny.Map {
	return _this.Get_().(D11_DC30).Cf69
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf69) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf69.Equals(data2.Cf69)
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

type D12_DC32 struct {
	Cf71 _dafny.Int
	Cf72 *C2
	Cf73 *C0
	Cf74 _dafny.Int
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf71 _dafny.Int, Cf72 *C2, Cf73 *C0, Cf74 _dafny.Int) D12 {
	return D12{D12_DC32{Cf71, Cf72, Cf73, Cf74}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC31 struct {
	Cf70 _dafny.Sequence
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf70 _dafny.Sequence) D12 {
	return D12{D12_DC31{Cf70}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC33 struct {
	Cf75 D12
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf75 D12) D12 {
	return D12{D12_DC33{Cf75}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC32_(_dafny.Zero, (*C2)(nil), (*C0)(nil), _dafny.Zero)
}

func (_this D12) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf71
}

func (_this D12) Dtor_cf72() *C2 {
	return _this.Get_().(D12_DC32).Cf72
}

func (_this D12) Dtor_cf73() *C0 {
	return _this.Get_().(D12_DC32).Cf73
}

func (_this D12) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf74
}

func (_this D12) Dtor_cf70() _dafny.Sequence {
	return _this.Get_().(D12_DC31).Cf70
}

func (_this D12) Dtor_cf75() D12 {
	return _this.Get_().(D12_DC33).Cf75
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf70) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72 == data2.Cf72 && data1.Cf73 == data2.Cf73 && data1.Cf74.Cmp(data2.Cf74) == 0
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf70.Equals(data2.Cf70)
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf75.Equals(data2.Cf75)
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

type D13_DC34 struct {
	Cf76 _dafny.Map
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf76 _dafny.Map) D13 {
	return D13{D13_DC34{Cf76}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D13) Dtor_cf76() _dafny.Map {
	return _this.Get_().(D13_DC34).Cf76
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf76) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf76.Equals(data2.Cf76)
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
	Cf77 _dafny.MultiSet
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf77 _dafny.MultiSet) D14 {
	return D14{D14_DC35{Cf77}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D14) Dtor_cf77() _dafny.MultiSet {
	return _this.Get_().(D14_DC35).Cf77
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf77) + ")"
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
			return ok && data1.Cf77.Equals(data2.Cf77)
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

type D15_DC37 struct {
	Cf79 _dafny.Sequence
	Cf80 bool
	Cf81 bool
	Cf82 _dafny.CodePoint
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf79 _dafny.Sequence, Cf80 bool, Cf81 bool, Cf82 _dafny.CodePoint) D15 {
	return D15{D15_DC37{Cf79, Cf80, Cf81, Cf82}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

type D15_DC36 struct {
	Cf78 _dafny.Map
}

func (D15_DC36) isD15() {}

func (CompanionStruct_D15_) Create_DC36_(Cf78 _dafny.Map) D15 {
	return D15{D15_DC36{Cf78}}
}

func (_this D15) Is_DC36() bool {
	_, ok := _this.Get_().(D15_DC36)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC37_(_dafny.EmptySeq, false, false, _dafny.CodePoint('D'))
}

func (_this D15) Dtor_cf79() _dafny.Sequence {
	return _this.Get_().(D15_DC37).Cf79
}

func (_this D15) Dtor_cf80() bool {
	return _this.Get_().(D15_DC37).Cf80
}

func (_this D15) Dtor_cf81() bool {
	return _this.Get_().(D15_DC37).Cf81
}

func (_this D15) Dtor_cf82() _dafny.CodePoint {
	return _this.Get_().(D15_DC37).Cf82
}

func (_this D15) Dtor_cf78() _dafny.Map {
	return _this.Get_().(D15_DC36).Cf78
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D15_DC36:
		{
			return "D15.DC36" + "(" + _dafny.String(data.Cf78) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf79.Equals(data2.Cf79) && data1.Cf80 == data2.Cf80 && data1.Cf81 == data2.Cf81 && data1.Cf82 == data2.Cf82
		}
	case D15_DC36:
		{
			data2, ok := other.Get_().(D15_DC36)
			return ok && data1.Cf78.Equals(data2.Cf78)
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

// Definition of trait T0
type T0 interface {
	String() string
	F11() _dafny.MultiSet
	F11_set_(value _dafny.MultiSet)
	F12() bool
	F12_set_(value bool)
	Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool
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
	F1   _dafny.Array
	F3   bool
	F4   _dafny.Array
	F6   _dafny.Sequence
	F7   bool
	F8   bool
	F9   _dafny.Int
	_f0  _dafny.Map
	_f2  bool
	_f5  _dafny.Int
	_f10 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F3 = false
	_this.F4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F6 = _dafny.EmptySeq
	_this.F7 = false
	_this.F8 = false
	_this.F9 = _dafny.Zero
	_this._f0 = _dafny.EmptyMap
	_this._f2 = false
	_this._f5 = _dafny.Zero
	_this._f10 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 _dafny.Array, f2 bool, f3 bool, f4 _dafny.Array, f5 _dafny.Int, f6 _dafny.Sequence, f7 bool, f8 bool, f9 _dafny.Int, f10 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
	}
}
func (_this *GlobalState) F0() _dafny.Map {
	{
		return _this._f0
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
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f11 _dafny.MultiSet
	_f12 bool
	_f17 bool
	_f18 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f11 = _dafny.EmptyMultiSet
	_this._f12 = false
	_this._f17 = false
	_this._f18 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F11() _dafny.MultiSet {
	return _this._f11
}
func (_this *C0) F11_set_(value _dafny.MultiSet) {
	_this._f11 = value
}
func (_this *C0) F12() bool {
	return _this._f12
}
func (_this *C0) F12_set_(value bool) {
	_this._f12 = value
}
func (_this *C0) Ctor__(f17 bool, f18 _dafny.Int, f11 _dafny.MultiSet, f12 bool) {
	{
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C0) Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return _this.F12()
	}
}
func (_this *C0) Fm11(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).F18(), (_this).F18()), Companion_D1_.Create_DC2_(_this.F12(), false, (_this).F17(), _dafny.CodePoint('f'), _dafny.CodePoint('x')))
	}
}
func (_this *C0) Fm12(globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F18(), _dafny.IntOfInt64(291))
	}
}
func (_this *C0) F17() bool {
	{
		return _this._f17
	}
}
func (_this *C0) F18() _dafny.Int {
	{
		return _this._f18
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f11 _dafny.MultiSet
	_f12 bool
	F16  _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f11 = _dafny.EmptyMultiSet
	_this._f12 = false
	_this.F16 = _dafny.Zero
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

func (_this *C1) F11() _dafny.MultiSet {
	return _this._f11
}
func (_this *C1) F11_set_(value _dafny.MultiSet) {
	_this._f11 = value
}
func (_this *C1) F12() bool {
	return _this._f12
}
func (_this *C1) F12_set_(value bool) {
	_this._f12 = value
}
func (_this *C1) Ctor__(f16 _dafny.Int, f11 _dafny.MultiSet, f12 bool) {
	{
		(_this).F16 = f16
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C1) Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		var _source6 D2 = Companion_D2_.Create_DC5_(_this.F16, _this.F16, _this.F16, _dafny.IntOfInt64(-458), _dafny.SeqOf(_this.F12()))
		_ = _source6
		if _source6.Is_DC5() {
			var _629___mcc_h0 _dafny.Int = _source6.Get_().(D2_DC5).Cf9
			_ = _629___mcc_h0
			var _630___mcc_h1 _dafny.Int = _source6.Get_().(D2_DC5).Cf10
			_ = _630___mcc_h1
			var _631___mcc_h2 _dafny.Int = _source6.Get_().(D2_DC5).Cf11
			_ = _631___mcc_h2
			var _632___mcc_h3 _dafny.Int = _source6.Get_().(D2_DC5).Cf12
			_ = _632___mcc_h3
			var _633___mcc_h4 _dafny.Sequence = _source6.Get_().(D2_DC5).Cf13
			_ = _633___mcc_h4
			var _634_cf13 _dafny.Sequence = _633___mcc_h4
			_ = _634_cf13
			var _635_cf12 _dafny.Int = _632___mcc_h3
			_ = _635_cf12
			var _636_cf11 _dafny.Int = _631___mcc_h2
			_ = _636_cf11
			var _637_cf10 _dafny.Int = _630___mcc_h1
			_ = _637_cf10
			var _638_cf9 _dafny.Int = _629___mcc_h0
			_ = _638_cf9
			return _this.F12()
		} else {
			var _639___mcc_h5 _dafny.Int = _source6.Get_().(D2_DC4).Cf8
			_ = _639___mcc_h5
			var _640_cf8 _dafny.Int = _639___mcc_h5
			_ = _640_cf8
			return !((func() bool {
				if _this.F12() {
					return _this.F12()
				}
				return false
			})())
		}
	}
}
func (_this *C1) M6(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _641_i0 _dafny.Int
		_ = _641_i0
		_641_i0 = _dafny.Zero
		{
			for _this.F12() {
				{
					if (_641_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_641_i0 = (_641_i0).Plus(_dafny.One)
					if ((func() _dafny.Int {
						if _this.F12() {
							return _this.F16
						}
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F12())).Cardinality()
					})()).Cmp(Companion_Default___.SafeDivisionInt(_this.F16, _this.F16)) == 0 {
						var _642_v0 D2
						_ = _642_v0
						_642_v0 = Companion_D2_.Create_DC4_(_dafny.IntOfInt64(43))
						(globalState).F9 = Companion_Default___.SafeModuloInt((_642_v0).Dtor_cf8(), _this.F16)
						var _643_v1 _dafny.Array
						_ = _643_v1
						var _len0_13 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_13
						var _nw108 _dafny.Array
						_ = _nw108
						if _len0_13.Cmp(_dafny.Zero) == 0 {
							_nw108 = _dafny.NewArray(_len0_13)
						} else {
							var _init13 func(_dafny.Int) _dafny.Int = func(_644_i1 _dafny.Int) _dafny.Int {
								return (_644_i1).Plus(_this.F16)
							}
							_ = _init13
							var _element0_13 = _init13(_dafny.Zero)
							_ = _element0_13
							_nw108 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
							(_nw108).ArraySet1(_element0_13, 0)
							var _nativeLen0_13 = (_len0_13).Int()
							_ = _nativeLen0_13
							for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
								(_nw108).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
							}
						}
						_643_v1 = _nw108
						var _645_v2 _dafny.Sequence
						_ = _645_v2
						_645_v2 = _dafny.SeqOf(_643_v1)
						var _646_v3 _dafny.Map
						_ = _646_v3
						_646_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_645_v2, _645_v2), (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_645_v2, _645_v2)).Cardinality()))).Uint32(), (_645_v2).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_645_v2).Cardinality()))).Uint32()).(_dafny.Array))).Cardinality()))
						_646_v3 = (_646_v3).Merge(_646_v3)
						var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))
						_ = _index68
						(_643_v1).ArraySet1((_dafny.Zero).Minus(_this.F16), (_index68).Int())
						var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))
						_ = _index69
						(_643_v1).ArraySet1(_dafny.IntOfInt64(276), (_index69).Int())
						var _647_v4 _dafny.Array
						_ = _647_v4
						var _nw109 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
						_ = _nw109
						_647_v4 = _nw109
						_647_v4 = _647_v4
						var _648_v5 _dafny.Array
						_ = _648_v5
						var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
						_ = _nw110
						_648_v5 = _nw110
						var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_648_v5), 0))
						_ = _index70
						(_648_v5).ArraySet1(_647_v4, (_index70).Int())
						var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_647_v4), 0))
						_ = _index71
						(_647_v4).ArraySet1(false, (_index71).Int())
						var _649_v6 _dafny.Array
						_ = _649_v6
						var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
						_ = _nw111
						_649_v6 = _nw111
						var _650_v7 _dafny.CodePoint
						_ = _650_v7
						_650_v7 = _dafny.CodePoint('t')
						var _651_v8 D1
						_ = _651_v8
						_651_v8 = Companion_D1_.Create_DC2_(_this.F12(), _this.F12(), true, _650_v7, _650_v7)
						var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_649_v6), 0))
						_ = _index72
						(_649_v6).ArraySet1(Companion_Default___.Fm9((_643_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))).Int()).(_dafny.Int), (_651_v8).Dtor_cf3(), globalState), (_index72).Int())
						var _652_v9 _dafny.Sequence
						_ = _652_v9
						_652_v9 = _dafny.UnicodeSeqOfUtf8Bytes("xcdmomb")
						var _653_v10 _dafny.MultiSet
						_ = _653_v10
						_653_v10 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(248))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg33 _dafny.Int) interface{} {
								return coer33(arg33)
							}
						}((func(_654_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_655_i2 _dafny.Int) _dafny.CodePoint {
								return _654_v7
							}
						})(_650_v7))), _dafny.UnicodeSeqOfUtf8Bytes("cqudam"), _652_v9)
						var _656_v11 _dafny.Map
						_ = _656_v11
						_656_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F12())
						var _657_v12 _dafny.Sequence
						_ = _657_v12
						_657_v12 = _dafny.SeqOf((_656_v11).Cardinality(), (_643_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_652_v9).Cardinality()), (_643_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))).Int()).(_dafny.Int))
						var _658_v13 _dafny.Sequence
						_ = _658_v13
						_658_v13 = _dafny.SeqOf(_652_v9)
						var _659_v14 _dafny.MultiSet
						_ = _659_v14
						_659_v14 = _dafny.MultiSetOf(_this.F16)
						var _660_v15 _dafny.Map
						_ = _660_v15
						_660_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_643_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))).Int()).(_dafny.Int), _this.F16)
						var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_648_v5), 0))
						_ = _index73
						var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_647_v4), 0))
						_ = _index74
						var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_649_v6), 0))
						_ = _index75
						var _rhs92 _dafny.Array = _647_v4
						_ = _rhs92
						var _rhs93 bool = (_653_v10).IsProperSubsetOf((_653_v10).Difference(_653_v10))
						_ = _rhs93
						var _rhs94 bool = _dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_656_v11).Cardinality()), _657_v12)
						_ = _rhs94
						var _rhs95 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_652_v9, (_658_v13).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1((_this).Fm2(_659_v14, _this.F16, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _dafny.IntOfInt64(-762)), _this.F12(), globalState), (_643_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_643_v1), 0))).Int()).(_dafny.Int), (_660_v15).Cardinality(), globalState), _dafny.IntOfUint32((_658_v13).Cardinality()))).Uint32()).(_dafny.Sequence))
						_ = _rhs95
						var _lhs68 _dafny.Array = _648_v5
						_ = _lhs68
						var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_648_v5), 0))
						_ = _lhs69
						var _lhs70 *C1 = _this
						_ = _lhs70
						var _lhs71 _dafny.Array = _647_v4
						_ = _lhs71
						var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_647_v4), 0))
						_ = _lhs72
						var _lhs73 _dafny.Array = _649_v6
						_ = _lhs73
						var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_649_v6), 0))
						_ = _lhs74
						(_lhs68).ArraySet1(_rhs92, (_lhs69).Int())
						_lhs70.F12_set_(_rhs93)
						(_lhs71).ArraySet1(_rhs94, (_lhs72).Int())
						(_lhs73).ArraySet1(_rhs95, (_lhs74).Int())
					} else {
						var _661_v16 _dafny.CodePoint
						_ = _661_v16
						_661_v16 = _dafny.CodePoint('t')
						var _662_v17 _dafny.Sequence
						_ = _662_v17
						_662_v17 = _dafny.SeqOf(_661_v16)
						var _663_v18 _dafny.Sequence
						_ = _663_v18
						_663_v18 = _dafny.SeqOf(_this.F12(), !_dafny.Companion_Sequence_.Contains(_662_v17, _dafny.CodePoint('p')), _this.F12(), _this.F12())
						(globalState).F6 = _dafny.Companion_Sequence_.Update(_663_v18, (Companion_Default___.SafeIndex((_this.F16).Times(_this.F16), _dafny.IntOfUint32((_663_v18).Cardinality()))).Uint32(), true)
						var _664_v19 _dafny.MultiSet
						_ = _664_v19
						_664_v19 = _dafny.MultiSetOf(_dafny.IntOfUint32((_662_v17).Cardinality()))
						var _665_v20 _dafny.Array
						_ = _665_v20
						var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
						_ = _nw112
						_665_v20 = _nw112
						var _666_v21 _dafny.Map
						_ = _666_v21
						_666_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v18, _665_v20)
						(globalState).F9 = (func() _dafny.Int {
							if !((_663_v18).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_663_v18).Cardinality()))).Uint32()).(bool)) {
								return (func() _dafny.Int {
									if (_664_v19).Contains((_666_v21).Cardinality()) {
										return (_664_v19).Multiplicity((_666_v21).Cardinality())
									}
									return _dafny.IntOfInt64(354)
								})()
							}
							return _dafny.IntOfInt64(-844)
						})()
						(globalState).F3 = !(_this.F12())
						(globalState).F9 = _this.F16
						var _667_v22 D1
						_ = _667_v22
						_667_v22 = Companion_D1_.Create_DC2_(_this.F12(), _this.F12(), _this.F12(), _661_v16, _661_v16)
						(globalState).F9 = (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if _this.F12() {
								return _this.F12()
							}
							return _this.F12()
						})(), (_667_v22).Dtor_cf5())).Cardinality())
					}
					var _668_v23 _dafny.CodePoint
					_ = _668_v23
					_668_v23 = _dafny.CodePoint('q')
					var _669_v24 _dafny.Sequence
					_ = _669_v24
					_669_v24 = _dafny.SeqOf(Companion_Default___.Fm0(_668_v23, globalState))
					var _670_v25 _dafny.Sequence
					_ = _670_v25
					_670_v25 = _dafny.UnicodeSeqOfUtf8Bytes("yfreelwss")
					var _671_v26 _dafny.Array
					_ = _671_v26
					var _nwElement0_25 bool = _this.F12()
					_ = _nwElement0_25
					var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(23))
					_ = _nw113
					(_nw113).ArraySet1(_nwElement0_25, 0)
					(_nw113).ArraySet1((func() bool {
						if _this.F12() {
							return _this.F12()
						}
						return false
					})(), 1)
					(_nw113).ArraySet1(_this.F12(), 2)
					(_nw113).ArraySet1(_this.F12(), 3)
					(_nw113).ArraySet1(_this.F12(), 4)
					(_nw113).ArraySet1(true, 5)
					(_nw113).ArraySet1((_this.F12()) == (!(_this.F12())), 6)
					(_nw113).ArraySet1(_dafny.Companion_Sequence_.Equal(_669_v24, _669_v24), 7)
					(_nw113).ArraySet1(_this.F12(), 8)
					(_nw113).ArraySet1(!(_this.F12()) || (_this.F12()), 9)
					(_nw113).ArraySet1((_this.F12()) || (!(_this.F12())), 10)
					(_nw113).ArraySet1(_this.F12(), 11)
					(_nw113).ArraySet1(_this.F12(), 12)
					(_nw113).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F12())).Cardinality()).Cmp(_this.F16) > 0, 13)
					(_nw113).ArraySet1(_this.F12(), 14)
					(_nw113).ArraySet1(_this.F12(), 15)
					(_nw113).ArraySet1((_dafny.IntOfUint32((_670_v25).Cardinality())).Cmp(_this.F16) > 0, 16)
					(_nw113).ArraySet1(!(_this.F12()), 17)
					(_nw113).ArraySet1(_this.F12(), 18)
					(_nw113).ArraySet1((func() bool {
						if _this.F12() {
							return _this.F12()
						}
						return _this.F12()
					})(), 19)
					(_nw113).ArraySet1(_this.F12(), 20)
					(_nw113).ArraySet1(_this.F12(), 21)
					(_nw113).ArraySet1(_this.F12(), 22)
					_671_v26 = _nw113
					var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_671_v26), 0))
					_ = _index76
					(_671_v26).ArraySet1(_this.F12(), (_index76).Int())
					var _672_v27 _dafny.Map
					_ = _672_v27
					_672_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F12()), _dafny.Companion_Sequence_.Equal(_670_v25, _670_v25))
					var _673_v28 _dafny.Sequence
					_ = _673_v28
					_673_v28 = _dafny.SeqOf(_669_v24, _669_v24, _dafny.SeqOf(_this.F12(), _this.F12()), _669_v24, Companion_Default___.Fm10(_this.F12(), _this.F12(), true, globalState))
					var _674_v29 _dafny.MultiSet
					_ = _674_v29
					_674_v29 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_673_v28).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_673_v28).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
					var _675_v30 _dafny.Map
					_ = _675_v30
					_675_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F16)
					var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_671_v26), 0))
					_ = _index77
					var _rhs96 _dafny.Int = (_this.F16).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ddcprsn")).Cardinality()))
					_ = _rhs96
					var _rhs97 bool = (func() bool {
						if (_672_v27).Contains(_this.F12()) {
							return (_672_v27).Get(_this.F12()).(bool)
						}
						return !(_this.F12()) || ((_this).Fm2(_674_v29, _this.F16, _675_v30, _this.F12(), globalState))
					})()
					_ = _rhs97
					var _lhs75 *GlobalState = globalState
					_ = _lhs75
					var _lhs76 _dafny.Array = _671_v26
					_ = _lhs76
					var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_671_v26), 0))
					_ = _lhs77
					_lhs75.F9 = _rhs96
					(_lhs76).ArraySet1(_rhs97, (_lhs77).Int())
					var _676_v31 _dafny.Array
					_ = _676_v31
					var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
					_ = _nw114
					_676_v31 = _nw114
					var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_676_v31), 0))
					_ = _index78
					(_676_v31).ArraySet1(_671_v26, (_index78).Int())
					var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_676_v31), 0))
					_ = _index79
					var _rhs98 _dafny.MultiSet = _674_v29
					_ = _rhs98
					var _rhs99 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nvme"), _670_v25), (Companion_Default___.SafeIndex((_this.F16).Minus(_this.F16), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nvme"), _670_v25)).Cardinality()))).Uint32(), _668_v23)
					_ = _rhs99
					var _rhs100 _dafny.Array = _671_v26
					_ = _rhs100
					var _rhs101 _dafny.Array = _671_v26
					_ = _rhs101
					var _lhs78 _dafny.Array = _676_v31
					_ = _lhs78
					var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(811), _dafny.ArrayLen((_676_v31), 0))
					_ = _lhs79
					_674_v29 = _rhs98
					_670_v25 = _rhs99
					(_lhs78).ArraySet1(_rhs100, (_lhs79).Int())
					_671_v26 = _rhs101
					var _677_v32 _dafny.MultiSet
					_ = _677_v32
					_677_v32 = _dafny.MultiSetOf(_668_v23, _668_v23, _668_v23)
					var _678_v33 _dafny.Sequence
					_ = _678_v33
					_678_v33 = _dafny.SeqOf(_this.F16, _this.F16)
					var _679_v34 _dafny.Map
					_ = _679_v34
					_679_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_677_v32, _678_v33)
					_679_v34 = (_679_v34).Update(_677_v32, _dafny.Companion_Sequence_.Update(_678_v33, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_678_v33).Cardinality()))).Uint32(), _this.F16))
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _680_v35 _dafny.Array
		_ = _680_v35
		var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw115
		_680_v35 = _nw115
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))
		_ = _index80
		(_680_v35).ArraySet1(_this.F16, (_index80).Int())
		var _681_v36 _dafny.Sequence
		_ = _681_v36
		_681_v36 = _dafny.UnicodeSeqOfUtf8Bytes("bso")
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))
		_ = _index81
		(_680_v35).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_681_v36).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(52))).Uint32(), func(coer34 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_682_i3 _dafny.Int) D1 {
			return Companion_D1_.Create_DC3_(Companion_D1_.Create_DC2_(_this.F12(), _this.F12(), _this.F12(), _dafny.CodePoint('k'), _dafny.CodePoint('d')))
		}))).Cardinality())), (_index81).Int())
		var _683_v37 _dafny.Map
		_ = _683_v37
		_683_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F12())
		_683_v37 = _683_v37
		var _684_v38 D2
		_ = _684_v38
		_684_v38 = Companion_D2_.Create_DC4_((_dafny.Zero).Minus(_this.F16))
		var _685_v39 _dafny.Array
		_ = _685_v39
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_14
		var _nw116 _dafny.Array
		_ = _nw116
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw116 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) bool = func(_686_i4 _dafny.Int) bool {
				return false
			}
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw116 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw116).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw116).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_685_v39 = _nw116
		var _687_v40 _dafny.Map
		_ = _687_v40
		_687_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D2 {
			if _this.F12() {
				return _684_v38
			}
			return _684_v38
		})(), _685_v39)
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))
		_ = _index82
		var _rhs102 _dafny.Int = (_this.F16).Plus((_this.F16).Times((_680_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))).Int()).(_dafny.Int)))
		_ = _rhs102
		var _rhs103 _dafny.Map = _687_v40
		_ = _rhs103
		var _lhs80 _dafny.Array = _680_v35
		_ = _lhs80
		var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))
		_ = _lhs81
		(_lhs80).ArraySet1(_rhs102, (_lhs81).Int())
		_687_v40 = _rhs103
		var _688_v41 _dafny.Sequence
		_ = _688_v41
		_688_v41 = _dafny.SeqOf(_this.F12())
		var _689_v42 _dafny.Map
		_ = _689_v42
		_689_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _688_v41)
		var _690_v43 _dafny.Array
		_ = _690_v43
		var _nwElement0_26 _dafny.Sequence = _688_v41
		_ = _nwElement0_26
		var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(10))
		_ = _nw117
		(_nw117).ArraySet1(_nwElement0_26, 0)
		(_nw117).ArraySet1(_688_v41, 1)
		(_nw117).ArraySet1(_688_v41, 2)
		(_nw117).ArraySet1(_dafny.Companion_Sequence_.Update(_688_v41, (Companion_Default___.SafeIndex((_680_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_688_v41).Cardinality()))).Uint32(), _this.F12()), 3)
		(_nw117).ArraySet1(_688_v41, 4)
		(_nw117).ArraySet1(_688_v41, 5)
		(_nw117).ArraySet1(Companion_Default___.Fm10(!(_this.F12()), _this.F12(), (_688_v41).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_688_v41).Cardinality()))).Uint32()).(bool), globalState), 6)
		(_nw117).ArraySet1(_688_v41, 7)
		(_nw117).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_689_v42).Contains(Companion_Default___.Fm1(_this.F12(), _this.F16, _this.F16, globalState)) {
				return (_689_v42).Get(Companion_Default___.Fm1(_this.F12(), _this.F16, _this.F16, globalState)).(_dafny.Sequence)
			}
			return _688_v41
		})(), _688_v41), 8)
		(_nw117).ArraySet1(_dafny.SeqOf(_this.F12()), 9)
		_690_v43 = _nw117
		var _691_v44 bool
		_ = _691_v44
		_691_v44 = _this.F12()
		var _692_v45 _dafny.Set
		_ = _692_v45
		_692_v45 = _dafny.SetOf(_this.F12())
		var _693_v46 _dafny.Sequence
		_ = _693_v46
		_693_v46 = _dafny.SeqOf(_688_v41, _688_v41, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F12(), (_691_v44), _this.F12()), (Companion_Default___.SafeIndex((_692_v45).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F12(), (_691_v44), _this.F12())).Cardinality()))).Uint32(), _this.F12()), _688_v41, _688_v41)
		var _694_v47 _dafny.Map
		_ = _694_v47
		_694_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_680_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))).Int()).(_dafny.Int), _this.F12())
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_690_v43), 0))
		_ = _index83
		(_690_v43).ArraySet1((_693_v46).Select((Companion_Default___.SafeIndex((_694_v47).Cardinality(), _dafny.IntOfUint32((_693_v46).Cardinality()))).Uint32()).(_dafny.Sequence), (_index83).Int())
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_690_v43), 0))
		_ = _index84
		(_690_v43).ArraySet1(_688_v41, (_index84).Int())
		_691_v44 = _691_v44
		r0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F12()), (_680_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_680_v35), 0))).Int()).(_dafny.Int))).Cardinality()
		return r0
	}
}
func (_this *C1) M7(p0 _dafny.MultiSet, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _695_v0 bool
		_ = _695_v0
		_695_v0 = _this.F12()
		var _696_v1 _dafny.CodePoint
		_ = _696_v1
		_696_v1 = _dafny.CodePoint('s')
		var _697_v2 D1
		_ = _697_v2
		_697_v2 = Companion_D1_.Create_DC2_((_695_v0), _this.F12(), _this.F12(), _696_v1, _dafny.CodePoint('w'))
		var _source7 D1 = _697_v2
		_ = _source7
		if _source7.Is_DC2() {
			var _698___mcc_h0 bool = _source7.Get_().(D1_DC2).Cf2
			_ = _698___mcc_h0
			var _699___mcc_h1 bool = _source7.Get_().(D1_DC2).Cf3
			_ = _699___mcc_h1
			var _700___mcc_h2 bool = _source7.Get_().(D1_DC2).Cf4
			_ = _700___mcc_h2
			var _701___mcc_h3 _dafny.CodePoint = _source7.Get_().(D1_DC2).Cf5
			_ = _701___mcc_h3
			var _702___mcc_h4 _dafny.CodePoint = _source7.Get_().(D1_DC2).Cf6
			_ = _702___mcc_h4
			var _703_cf6 _dafny.CodePoint = _702___mcc_h4
			_ = _703_cf6
			var _704_cf5 _dafny.CodePoint = _701___mcc_h3
			_ = _704_cf5
			var _705_cf4 bool = _700___mcc_h2
			_ = _705_cf4
			var _706_cf3 bool = _699___mcc_h1
			_ = _706_cf3
			var _707_cf2 bool = _698___mcc_h0
			_ = _707_cf2
			(globalState).F9 = _this.F16
			var _708_v3 _dafny.Sequence
			_ = _708_v3
			_708_v3 = _dafny.SeqOf(_this.F16, _this.F16, _this.F16)
			(_this).F16 = (_708_v3).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_708_v3).Cardinality()))).Uint32()).(_dafny.Int)
			var _rhs104 bool = _706_cf3
			_ = _rhs104
			var _rhs105 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_708_v3, _dafny.Companion_Sequence_.Concatenate(_708_v3, _708_v3))
			_ = _rhs105
			var _rhs106 bool = false
			_ = _rhs106
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			_lhs82.F7 = _rhs104
			_708_v3 = _rhs105
			_lhs83.F3 = _rhs106
			if ((p0).Difference(_dafny.MultiSetOf(_this.F16))).IsProperSubsetOf(p0) {
				(globalState).F7 = (func() bool {
					if _706_cf3 {
						return _705_cf4
					}
					return !(_this.F12()) || (_this.F12())
				})()
				var _709_v4 _dafny.Map
				_ = _709_v4
				_709_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm9(_this.F16, _706_cf3, globalState)).Cardinality()), _706_cf3)
				var _710_v5 _dafny.Array
				_ = _710_v5
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_15
				var _nw118 _dafny.Array
				_ = _nw118
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw118 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) bool = (func(_711_cf4 bool) func(_dafny.Int) bool {
						return func(_712_i0 _dafny.Int) bool {
							return _711_cf4
						}
					})(_705_cf4)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw118 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw118).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw118).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_710_v5 = _nw118
				var _713_v6 _dafny.Map
				_ = _713_v6
				_713_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_710_v5, _705_cf4)
				_709_v4 = (_709_v4).Update((_dafny.Zero).Minus(_this.F16), (func() bool {
					if (_713_v6).Contains(_710_v5) {
						return (_713_v6).Get(_710_v5).(bool)
					}
					return _705_cf4
				})())
				var _714_v7 _dafny.Sequence
				_ = _714_v7
				_714_v7 = _dafny.SeqOf(_705_cf4)
				var _rhs107 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm9(_dafny.IntOfInt64(-404), _706_cf3, globalState)).Cardinality())
				_ = _rhs107
				var _rhs108 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F16, _dafny.IntOfInt64(302))
				_ = _rhs108
				var _rhs109 bool = _707_cf2
				_ = _rhs109
				var _rhs110 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_714_v7)).Union((func() _dafny.MultiSet {
					if _705_cf4 {
						return _this.F11()
					}
					return _dafny.MultiSetOf(_707_cf2, _706_cf3, _706_cf3)
				})())
				_ = _rhs110
				var _rhs111 bool = _706_cf3
				_ = _rhs111
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *C1 = _this
				_ = _lhs86
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				_lhs84.F9 = _rhs107
				_lhs85.F9 = _rhs108
				_706_cf3 = _rhs109
				_lhs86.F11_set_(_rhs110)
				_lhs87.F3 = _rhs111
				(globalState).F9 = _this.F16
				var _715_v8 _dafny.Sequence
				_ = _715_v8
				_715_v8 = _dafny.SeqOf(p1)
				var _716_v9 _dafny.Set
				_ = _716_v9
				_716_v9 = _dafny.SetOf(_this.F16)
				var _717_v10 _dafny.Array
				_ = _717_v10
				var _nwElement0_27 _dafny.Sequence = p1
				_ = _nwElement0_27
				var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(23))
				_ = _nw119
				(_nw119).ArraySet1(_nwElement0_27, 0)
				(_nw119).ArraySet1(p1, 1)
				(_nw119).ArraySet1((_715_v8).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_715_v8).Cardinality()))).Uint32()).(_dafny.Sequence), 2)
				(_nw119).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("exlcjc"), 3)
				(_nw119).ArraySet1(p1, 4)
				(_nw119).ArraySet1(p1, 5)
				(_nw119).ArraySet1(p1, 6)
				(_nw119).ArraySet1(Companion_Default___.Fm9((_716_v9).Cardinality(), _706_cf3, globalState), 7)
				(_nw119).ArraySet1(p1, 8)
				(_nw119).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(527))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}((func(_718_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_719_i1 _dafny.Int) _dafny.CodePoint {
						return _718_v1
					}
				})(_696_v1))), 9)
				(_nw119).ArraySet1(p1, 10)
				(_nw119).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}((func(_720_cf6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_721_i2 _dafny.Int) _dafny.CodePoint {
						return _720_cf6
					}
				})(_703_cf6))), 11)
				(_nw119).ArraySet1(p1, 12)
				(_nw119).ArraySet1(p1, 13)
				(_nw119).ArraySet1(p1, 14)
				(_nw119).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("mw"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_714_v7).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mw")).Cardinality()))).Uint32(), _703_cf6), 15)
				(_nw119).ArraySet1(p1, 16)
				(_nw119).ArraySet1(p1, 17)
				(_nw119).ArraySet1(p1, 18)
				(_nw119).ArraySet1(p1, 19)
				(_nw119).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_722_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_723_i3 _dafny.Int) _dafny.CodePoint {
						return _722_v1
					}
				})(_696_v1))), 20)
				(_nw119).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(233))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_724_cf6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_725_i4 _dafny.Int) _dafny.CodePoint {
						return _724_cf6
					}
				})(_703_cf6))), 21)
				(_nw119).ArraySet1(p1, 22)
				_717_v10 = _nw119
				var _726_v11 _dafny.Map
				_ = _726_v11
				_726_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_704_cf5, _717_v10)
				var _727_v12 _dafny.Map
				_ = _727_v12
				_727_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_707_cf2, _717_v10)
				var _728_v13 _dafny.Sequence
				_ = _728_v13
				_728_v13 = _dafny.SeqOf(_717_v10)
				var _729_v14 _dafny.Array
				_ = _729_v14
				var _nwElement0_28 _dafny.Array = _717_v10
				_ = _nwElement0_28
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(19))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_28, 0)
				(_nw120).ArraySet1(_717_v10, 1)
				(_nw120).ArraySet1(_717_v10, 2)
				(_nw120).ArraySet1((func() _dafny.Array {
					if (_726_v11).Contains(_696_v1) {
						return (_726_v11).Get(_696_v1).(_dafny.Array)
					}
					return _717_v10
				})(), 3)
				(_nw120).ArraySet1((func() _dafny.Array {
					if _this.F12() {
						return _717_v10
					}
					return _717_v10
				})(), 4)
				(_nw120).ArraySet1(_717_v10, 5)
				(_nw120).ArraySet1(_717_v10, 6)
				(_nw120).ArraySet1(_717_v10, 7)
				(_nw120).ArraySet1(_717_v10, 8)
				(_nw120).ArraySet1((func() _dafny.Array {
					if (_727_v12).Contains(_705_cf4) {
						return (_727_v12).Get(_705_cf4).(_dafny.Array)
					}
					return _717_v10
				})(), 9)
				(_nw120).ArraySet1(_717_v10, 10)
				(_nw120).ArraySet1(_717_v10, 11)
				(_nw120).ArraySet1(_717_v10, 12)
				(_nw120).ArraySet1(_717_v10, 13)
				(_nw120).ArraySet1((func() _dafny.Array {
					if _705_cf4 {
						return (_728_v13).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_728_v13).Cardinality()))).Uint32()).(_dafny.Array)
					}
					return _717_v10
				})(), 14)
				(_nw120).ArraySet1(_717_v10, 15)
				(_nw120).ArraySet1(_717_v10, 16)
				(_nw120).ArraySet1((func() _dafny.Array {
					if (_727_v12).Contains(_707_cf2) {
						return (_727_v12).Get(_707_cf2).(_dafny.Array)
					}
					return _717_v10
				})(), 17)
				(_nw120).ArraySet1(_717_v10, 18)
				_729_v14 = _nw120
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_729_v14), 0))
				_ = _index85
				(_729_v14).ArraySet1(_717_v10, (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_729_v14), 0))
				_ = _index86
				(_729_v14).ArraySet1(_717_v10, (_index86).Int())
			} else {
				var _730_v15 _dafny.Map
				_ = _730_v15
				_730_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F16)
				(globalState).F7 = (_this).Fm2(((p0).Update(_this.F16, Companion_Default___.Abs(_this.F16))).Union(p0), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(768), (_dafny.Zero).Minus(_this.F16)), _730_v15, (_this.F16).Cmp(_this.F16) == 0, globalState)
				_696_v1 = _696_v1
				(globalState).F9 = (_this.F16).Minus(_this.F16)
				var _731_v16 _dafny.Sequence
				_ = _731_v16
				_731_v16 = _dafny.SeqOf(_705_cf4, _707_cf2)
				var _732_v17 _dafny.Map
				_ = _732_v17
				_732_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(374))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_733_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_734_i5 _dafny.Int) _dafny.CodePoint {
						return _733_v1
					}
				})(_696_v1))), _dafny.Companion_Sequence_.Concatenate(_731_v16, _731_v16))
				var _735_v18 T0
				_ = _735_v18
				var _nw121 *C0 = New_C0_()
				_ = _nw121
				_nw121.Ctor__(_706_cf3, _this.F16, _this.F11(), _707_cf2)
				_735_v18 = _nw121
				var _736_v19 _dafny.Map
				_ = _736_v19
				_736_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_735_v18, _731_v16)
				_732_v17 = (_732_v17).Update((func() _dafny.Sequence {
					if _706_cf3 {
						return p1
					}
					return p1
				})(), (func() _dafny.Sequence {
					if (_736_v19).Contains(_735_v18) {
						return (_736_v19).Get(_735_v18).(_dafny.Sequence)
					}
					return _731_v16
				})())
				var _737_v20 _dafny.Array
				_ = _737_v20
				var _nw122 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
				_ = _nw122
				_737_v20 = _nw122
				var _738_v21 _dafny.Map
				_ = _738_v21
				_738_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _this.F16)
				var _739_v22 _dafny.Map
				_ = _739_v22
				_739_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_703_cf6, p1)
				var _rhs112 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_731_v16, _731_v16)).Cardinality()), _this.F16)
				_ = _rhs112
				var _rhs113 bool = (_731_v16).Select((Companion_Default___.SafeIndex((_this.F16).Times(_this.F16), _dafny.IntOfUint32((_731_v16).Cardinality()))).Uint32()).(bool)
				_ = _rhs113
				var _rhs114 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _705_cf4)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_707_cf2)).Cardinality())), (_this.F16).Plus(_this.F16))
				_ = _rhs114
				var _rhs115 _dafny.Array = _737_v20
				_ = _rhs115
				var _rhs116 _dafny.Int = (func() _dafny.Int {
					if _707_cf2 {
						return _this.F16
					}
					return (_dafny.Zero).Minus((_this.F16).Minus((func() _dafny.Int {
						if (_738_v21).Contains(_this.F16) {
							return (_738_v21).Get(_this.F16).(_dafny.Int)
						}
						return (_739_v22).Cardinality()
					})()))
				})()
				_ = _rhs116
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 *C1 = _this
				_ = _lhs90
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				_lhs88.F9 = _rhs112
				_lhs89.F7 = _rhs113
				_lhs90.F16 = _rhs114
				_737_v20 = _rhs115
				_lhs91.F9 = _rhs116
			}
		} else if _source7.Is_DC1() {
			var _740___mcc_h5 _dafny.Set = _source7.Get_().(D1_DC1).Cf1
			_ = _740___mcc_h5
			var _741_cf1 _dafny.Set = _740___mcc_h5
			_ = _741_cf1
			var _742_v23 D1
			_ = _742_v23
			_742_v23 = Companion_D1_.Create_DC1_(_741_cf1)
			var _743_v24 _dafny.Map
			_ = _743_v24
			_743_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_742_v23, _741_cf1)
			var _744_v25 _dafny.Array
			_ = _744_v25
			var _nwElement0_29 _dafny.Int = (p0).Cardinality()
			_ = _nwElement0_29
			var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(8))
			_ = _nw123
			(_nw123).ArraySet1(_nwElement0_29, 0)
			(_nw123).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("q"))).Cardinality()), 1)
			(_nw123).ArraySet1(_dafny.IntOfInt64(135), 2)
			(_nw123).ArraySet1(_this.F16, 3)
			(_nw123).ArraySet1((_this.F16).Times(Companion_Default___.Fm1(_this.F12(), _dafny.IntOfInt64(467), _this.F16, globalState)), 4)
			(_nw123).ArraySet1(_this.F16, 5)
			(_nw123).ArraySet1(_this.F16, 6)
			(_nw123).ArraySet1(((func() _dafny.Set {
				if (_743_v24).Contains(_742_v23) {
					return (_743_v24).Get(_742_v23).(_dafny.Set)
				}
				return _741_cf1
			})()).Cardinality(), 7)
			_744_v25 = _nw123
			(globalState).F1 = _744_v25
			var _745_v27 _dafny.Map
			_ = _745_v27
			_745_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll18 = _dafny.NewMapBuilder()
				_ = _coll18
				for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(893), _dafny.IntOfInt64(501))); ; {
					_compr_18, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _746_v26 _dafny.Int
					_746_v26 = interface{}(_compr_18).(_dafny.Int)
					if ((_dafny.IntOfInt64(893)).Cmp(_746_v26) <= 0) && ((_746_v26).Cmp(_dafny.IntOfInt64(501)) < 0) {
						_coll18.Add((_746_v26).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _696_v1)).Cardinality()), _this.F12())
					}
				}
				return _coll18.ToMap()
			}()).Update(_dafny.IntOfInt64(73), _this.F12()), _this.F16))
			var _747_v29 _dafny.Set
			_ = _747_v29
			_747_v29 = _dafny.SetOf(_this.F16)
			if (func() bool {
				if _this.F12() {
					return _this.F12()
				}
				return (_this.F16).Cmp(((func() _dafny.Map {
					if (_745_v27).Contains(_dafny.IntOfInt64(-670)) {
						return (_745_v27).Get(_dafny.IntOfInt64(-670)).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
						var _coll19 = _dafny.NewMapBuilder()
						_ = _coll19
						for _iter20 := _dafny.Iterate((_747_v29).Elements()); ; {
							_compr_19, _ok20 := _iter20()
							if !_ok20 {
								break
							}
							var _748_v28 _dafny.Int
							_748_v28 = interface{}(_compr_19).(_dafny.Int)
							if (_747_v29).Contains(_748_v28) {
								_coll19.Add(Companion_Default___.SafeModuloInt(_748_v28, _dafny.IntOfInt64(553)), _this.F12())
							}
						}
						return _coll19.ToMap()
					}(), _this.F16)
				})()).Cardinality()) < 0
			})() {
				(globalState).F7 = _this.F12()
				(_this).F12_set_(_this.F12())
				var _rhs117 _dafny.Int = Companion_Default___.Fm1(_this.F12(), _this.F16, _this.F16, globalState)
				_ = _rhs117
				var _rhs118 bool = false
				_ = _rhs118
				var _lhs92 *C1 = _this
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				_lhs92.F16 = _rhs117
				_lhs93.F3 = _rhs118
				(globalState).F9 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm1(_this.F12(), _dafny.IntOfInt64(834), (func() _dafny.Int {
					if (_this.F11()).Contains(_this.F12()) {
						return (_this.F11()).Multiplicity(_this.F12())
					}
					return _dafny.IntOfInt64(849)
				})(), globalState), (_this.F16).Plus(_dafny.IntOfInt64(945)))
				var _749_v30 _dafny.Array
				_ = _749_v30
				var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
				_ = _nw124
				_749_v30 = _nw124
				var _750_v31 _dafny.Map
				_ = _750_v31
				_750_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_this.F16), _749_v30)
				var _751_v32 D2
				_ = _751_v32
				_751_v32 = Companion_D2_.Create_DC4_(_this.F16)
				var _752_v33 _dafny.Map
				_ = _752_v33
				_752_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _751_v32)
				_750_v31 = (_750_v31).Update((p0).Intersection((p0).Update((_752_v33).Cardinality(), Companion_Default___.Abs(_this.F16))), _749_v30)
			} else {
				var _753_v34 _dafny.Sequence
				_ = _753_v34
				_753_v34 = _dafny.UnicodeSeqOfUtf8Bytes("pekwurnd")
				var _754_v35 _dafny.Map
				_ = _754_v35
				_754_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_696_v1, _dafny.UnicodeSeqOfUtf8Bytes("tfxhr"))
				_753_v34 = (func() _dafny.Sequence {
					if (_754_v35).Contains(_dafny.CodePoint('w')) {
						return (_754_v35).Get(_dafny.CodePoint('w')).(_dafny.Sequence)
					}
					return p1
				})()
				var _755_v36 _dafny.Array
				_ = _755_v36
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_16
				var _nw125 _dafny.Array
				_ = _nw125
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw125 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) bool = func(_756_i6 _dafny.Int) bool {
						return !((func() bool {
							if true {
								return _this.F12()
							}
							return true
						})())
					}
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw125 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw125).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw125).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_755_v36 = _nw125
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_755_v36), 0))
				_ = _index87
				(_755_v36).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gdkv")).Cardinality())).Cmp(_this.F16) > 0, (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_755_v36), 0))
				_ = _index88
				(_755_v36).ArraySet1(_this.F12(), (_index88).Int())
				var _757_v37 D3
				_ = _757_v37
				_757_v37 = Companion_D3_.Create_DC6_(_755_v36)
				_755_v36 = (_757_v37).Dtor_cf14()
				var _758_v38 _dafny.Array
				_ = _758_v38
				var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
				_ = _nw126
				_758_v38 = _nw126
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_758_v38), 0))
				_ = _index89
				(_758_v38).ArraySet1(Companion_Default___.Fm13(_696_v1, _this.F12(), globalState), (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_758_v38), 0))
				_ = _index90
				(_758_v38).ArraySet1((p0).Intersection(p0), (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_755_v36), 0))
				_ = _index91
				(_755_v36).ArraySet1(_this.F12(), (_index91).Int())
			}
			var _759_v39 _dafny.Map
			_ = _759_v39
			_759_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _this.F16)
			var _760_v40 _dafny.Sequence
			_ = _760_v40
			_760_v40 = _dafny.SeqOf(_759_v39, _759_v39)
			_760_v40 = (func() _dafny.Sequence {
				if true {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_759_v39), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_759_v39)).Cardinality()))).Uint32(), _759_v39), _dafny.SeqOf(_759_v39))
				}
				return _760_v40
			})()
			var _761_v41 _dafny.Array
			_ = _761_v41
			var _nw127 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw127
			_761_v41 = _nw127
			_761_v41 = _761_v41
		} else {
			var _762___mcc_h6 D1 = _source7.Get_().(D1_DC3).Cf7
			_ = _762___mcc_h6
			var _763_cf7 D1 = _762___mcc_h6
			_ = _763_cf7
			if !(Companion_Default___.Fm0(_696_v1, globalState)) {
				var _764_v42 _dafny.Array
				_ = _764_v42
				var _nwElement0_30 _dafny.CodePoint = _696_v1
				_ = _nwElement0_30
				var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(24))
				_ = _nw128
				(_nw128).ArraySet1CodePoint(_nwElement0_30, 0)
				(_nw128).ArraySet1CodePoint(_696_v1, 1)
				(_nw128).ArraySet1CodePoint(_696_v1, 2)
				(_nw128).ArraySet1CodePoint(_696_v1, 3)
				(_nw128).ArraySet1CodePoint(_dafny.CodePoint('f'), 4)
				(_nw128).ArraySet1CodePoint(_dafny.CodePoint('w'), 5)
				(_nw128).ArraySet1CodePoint(_696_v1, 6)
				(_nw128).ArraySet1CodePoint(_696_v1, 7)
				(_nw128).ArraySet1CodePoint(_696_v1, 8)
				(_nw128).ArraySet1CodePoint(_696_v1, 9)
				(_nw128).ArraySet1CodePoint(_696_v1, 10)
				(_nw128).ArraySet1CodePoint(_dafny.CodePoint('c'), 11)
				(_nw128).ArraySet1CodePoint(_696_v1, 12)
				(_nw128).ArraySet1CodePoint(_696_v1, 13)
				(_nw128).ArraySet1CodePoint(_696_v1, 14)
				(_nw128).ArraySet1CodePoint(_dafny.CodePoint('r'), 15)
				(_nw128).ArraySet1CodePoint(_696_v1, 16)
				(_nw128).ArraySet1CodePoint(_696_v1, 17)
				(_nw128).ArraySet1CodePoint(_696_v1, 18)
				(_nw128).ArraySet1CodePoint(_696_v1, 19)
				(_nw128).ArraySet1CodePoint(_696_v1, 20)
				(_nw128).ArraySet1CodePoint(_696_v1, 21)
				(_nw128).ArraySet1CodePoint(_696_v1, 22)
				(_nw128).ArraySet1CodePoint(_dafny.CodePoint('w'), 23)
				_764_v42 = _nw128
				var _765_v43 _dafny.Sequence
				_ = _765_v43
				_765_v43 = _dafny.SeqOf(_764_v42)
				_765_v43 = _765_v43
				var _766_v44 _dafny.Map
				_ = _766_v44
				_766_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(588), _this.F12())
				_766_v44 = (_766_v44).Update(_this.F16, _this.F12())
				var _767_v45 _dafny.Set
				_ = _767_v45
				_767_v45 = _dafny.SetOf(_this.F12(), _this.F12(), _this.F12(), _this.F12())
				(globalState).F8 = (_this.F16).Cmp((_767_v45).Cardinality()) == 0
				var _768_v46 _dafny.Array
				_ = _768_v46
				var _nw129 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw129
				_768_v46 = _nw129
				_768_v46 = _768_v46
				var _769_v47 _dafny.Sequence
				_ = _769_v47
				_769_v47 = _dafny.UnicodeSeqOfUtf8Bytes("njgjri")
				var _770_v48 _dafny.Array
				_ = _770_v48
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_17
				var _nw130 _dafny.Array
				_ = _nw130
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw130 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = func(_771_i7 _dafny.Int) _dafny.Int {
						return (_771_i7).Times((_dafny.Zero).Minus(_this.F16))
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw130 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw130).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw130).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_770_v48 = _nw130
				var _772_v49 D4
				_ = _772_v49
				_772_v49 = Companion_D4_.Create_DC10_(_770_v48)
				var _773_v50 _dafny.Sequence
				_ = _773_v50
				_773_v50 = _dafny.SeqOf(_770_v48)
				var _pat_let_tv15 = _770_v48
				_ = _pat_let_tv15
				var _774_v51 _dafny.Array
				_ = _774_v51
				var _nwElement0_31 _dafny.Array = _770_v48
				_ = _nwElement0_31
				var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(11))
				_ = _nw131
				(_nw131).ArraySet1(_nwElement0_31, 0)
				(_nw131).ArraySet1(_770_v48, 1)
				(_nw131).ArraySet1((func() _dafny.Array {
					if _this.F12() {
						return _770_v48
					}
					return _770_v48
				})(), 2)
				(_nw131).ArraySet1((func(_pat_let19_0 D4) D4 {
					return func(_775_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let20_0 _dafny.Array) D4 {
							return func(_776_dt__update_hcf24_h0 _dafny.Array) D4 {
								return Companion_D4_.Create_DC10_(_776_dt__update_hcf24_h0)
							}(_pat_let20_0)
						}(_pat_let_tv15)
					}(_pat_let19_0)
				}(_772_v49)).Dtor_cf24(), 3)
				(_nw131).ArraySet1((_773_v50).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_773_v50).Cardinality()))).Uint32()).(_dafny.Array), 4)
				(_nw131).ArraySet1(_770_v48, 5)
				(_nw131).ArraySet1(_770_v48, 6)
				(_nw131).ArraySet1(_770_v48, 7)
				(_nw131).ArraySet1(_770_v48, 8)
				(_nw131).ArraySet1(_770_v48, 9)
				(_nw131).ArraySet1(_770_v48, 10)
				_774_v51 = _nw131
				var _777_v52 _dafny.Map
				_ = _777_v52
				_777_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _696_v1)
				var _778_v53 _dafny.Map
				_ = _778_v53
				_778_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F16)
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_768_v46), 0))
				_ = _index92
				(_768_v46).ArraySet1(((_this).Fm2(_dafny.MultiSetOf((func() _dafny.Int {
					if (_this.F11()).Contains(!(_this.F12())) {
						return (_this.F11()).Multiplicity(!(_this.F12()))
					}
					return (_777_v52).Cardinality()
				})()), _this.F16, _778_v53, _this.F12(), globalState)) && (_this.F12()), (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_768_v46), 0))
				_ = _index93
				var _rhs119 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_769_v47, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_769_v47).Cardinality()))).Uint32(), Companion_Default___.Fm14(_697_v2, _this.F16, globalState))
				_ = _rhs119
				var _rhs120 _dafny.Array = _774_v51
				_ = _rhs120
				var _rhs121 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_769_v47, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_769_v47).Cardinality()))).Uint32(), _dafny.CodePoint('c')), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_769_v47, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_769_v47).Cardinality()))).Uint32(), _696_v1), p1))
				_ = _rhs121
				var _lhs94 _dafny.Array = _768_v46
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_768_v46), 0))
				_ = _lhs95
				_769_v47 = _rhs119
				_774_v51 = _rhs120
				(_lhs94).ArraySet1(_rhs121, (_lhs95).Int())
			} else {
				var _779_v54 _dafny.Map
				_ = _779_v54
				_779_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, _this.F12())
				_696_v1 = (func() _dafny.CodePoint {
					if (_779_v54).Equals(_779_v54) {
						return _696_v1
					}
					return _696_v1
				})()
				var _780_v55 _dafny.Array
				_ = _780_v55
				var _len0_18 _dafny.Int = _dafny.One
				_ = _len0_18
				var _nw132 _dafny.Array
				_ = _nw132
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw132 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = func(_781_i8 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("nol")
					}
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw132 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw132).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw132).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_780_v55 = _nw132
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_780_v55), 0))
				_ = _index94
				(_780_v55).ArraySet1(p1, (_index94).Int())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_780_v55), 0))
				_ = _index95
				(_780_v55).ArraySet1(p1, (_index95).Int())
				(globalState).F9 = (_this.F16).Times(_this.F16)
				var _782_v56 _dafny.Sequence
				_ = _782_v56
				_782_v56 = _dafny.SeqOf(_this.F16, _dafny.IntOfInt64(278))
				var _783_v57 D5
				_ = _783_v57
				_783_v57 = Companion_D5_.Create_DC12_(_782_v56)
				_782_v56 = _dafny.Companion_Sequence_.Update((_783_v57).Dtor_cf29(), (Companion_Default___.SafeIndex((func() _dafny.Int {
					if _this.F12() {
						return (_dafny.MultiSetOf(_this.F12(), _this.F12())).Cardinality()
					}
					return (_dafny.Zero).Minus(_dafny.IntOfUint32(((_780_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_780_v55), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				})(), _dafny.IntOfUint32(((_783_v57).Dtor_cf29()).Cardinality()))).Uint32(), _this.F16)
				var _784_v58 _dafny.Sequence
				_ = _784_v58
				_784_v58 = _dafny.SeqOf(_this.F12(), _this.F12())
				(globalState).F9 = Companion_Default___.SafeDivisionInt(_this.F16, (_this.F16).Minus(_dafny.IntOfUint32((_784_v58).Cardinality())))
			}
			var _785_v59 _dafny.Array
			_ = _785_v59
			var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw133
			_785_v59 = _nw133
			r0 = _785_v59
			var _source8 bool = _695_v0
			_ = _source8
			var _786___mcc_h7 bool = _source8
			_ = _786___mcc_h7
			var _787_cf0 bool = _786___mcc_h7
			_ = _787_cf0
			var _788_v60 _dafny.Array
			_ = _788_v60
			var _nw134 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw134
			_788_v60 = _nw134
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_788_v60), 0))
			_ = _index96
			(_788_v60).ArraySet1(_this.F16, (_index96).Int())
			var _789_v61 _dafny.Sequence
			_ = _789_v61
			_789_v61 = _dafny.UnicodeSeqOfUtf8Bytes("jljlfkl")
			var _790_v62 _dafny.Array
			_ = _790_v62
			var _nw135 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(21))
			_ = _nw135
			_790_v62 = _nw135
			var _791_v63 D1
			_ = _791_v63
			_791_v63 = Companion_D1_.Create_DC3_(Companion_Default___.Fm15(_dafny.IntOfUint32((_789_v61).Cardinality()), _this.F16, _787_cf0, globalState))
			var _792_v64 D1
			_ = _792_v64
			_792_v64 = Companion_D1_.Create_DC3_(_791_v63)
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_790_v62), 0))
			_ = _index97
			(_790_v62).ArraySet1(_792_v64, (_index97).Int())
			var _793_v65 D3
			_ = _793_v65
			_793_v65 = Companion_D3_.Create_DC8_(_this.F16, Companion_Default___.Fm1(true, _this.F16, _this.F16, globalState))
			var _794_v66 _dafny.Sequence
			_ = _794_v66
			_794_v66 = _dafny.SeqOf(true, _this.F12())
			var _795_v67 _dafny.Map
			_ = _795_v67
			_795_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_this.F16, _dafny.IntOfUint32((_dafny.SeqOf(_this.F16, _dafny.IntOfUint32((_794_v66).Cardinality()))).Cardinality())), _this.F16)
			var _796_v68 _dafny.Sequence
			_ = _796_v68
			_796_v68 = _dafny.SeqOf(_this.F16)
			var _797_v69 _dafny.Set
			_ = _797_v69
			_797_v69 = _dafny.SetOf(Companion_Default___.Fm9(_this.F16, _this.F12(), globalState))
			var _798_v70 _dafny.MultiSet
			_ = _798_v70
			_798_v70 = _dafny.MultiSetOf(p1, _dafny.UnicodeSeqOfUtf8Bytes("sarchm"), p1)
			var _799_v71 _dafny.Array
			_ = _799_v71
			var _nwElement0_32 _dafny.CodePoint = _696_v1
			_ = _nwElement0_32
			var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
			_ = _nw136
			(_nw136).ArraySet1CodePoint(_nwElement0_32, 0)
			(_nw136).ArraySet1CodePoint(_696_v1, 1)
			(_nw136).ArraySet1CodePoint(_696_v1, 2)
			(_nw136).ArraySet1CodePoint(_696_v1, 3)
			(_nw136).ArraySet1CodePoint(_696_v1, 4)
			(_nw136).ArraySet1CodePoint(_696_v1, 5)
			(_nw136).ArraySet1CodePoint(Companion_Default___.Fm14(_697_v2, _this.F16, globalState), 6)
			(_nw136).ArraySet1CodePoint(_696_v1, 7)
			(_nw136).ArraySet1CodePoint(_696_v1, 8)
			(_nw136).ArraySet1CodePoint(_696_v1, 9)
			(_nw136).ArraySet1CodePoint(_696_v1, 10)
			(_nw136).ArraySet1CodePoint(_696_v1, 11)
			(_nw136).ArraySet1CodePoint(_696_v1, 12)
			(_nw136).ArraySet1CodePoint(_696_v1, 13)
			(_nw136).ArraySet1CodePoint(_696_v1, 14)
			(_nw136).ArraySet1CodePoint(Companion_Default___.Fm14(Companion_Default___.Fm15(_this.F16, _this.F16, _this.F12(), globalState), _dafny.IntOfInt64(651), globalState), 15)
			(_nw136).ArraySet1CodePoint(_dafny.CodePoint('o'), 16)
			(_nw136).ArraySet1CodePoint(_696_v1, 17)
			(_nw136).ArraySet1CodePoint(_696_v1, 18)
			(_nw136).ArraySet1CodePoint(_dafny.CodePoint('p'), 19)
			(_nw136).ArraySet1CodePoint(_696_v1, 20)
			(_nw136).ArraySet1CodePoint(_696_v1, 21)
			(_nw136).ArraySet1CodePoint(_696_v1, 22)
			(_nw136).ArraySet1CodePoint((_789_v61).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_789_v61).Cardinality()))).Uint32()).(_dafny.CodePoint), 23)
			(_nw136).ArraySet1CodePoint(_696_v1, 24)
			_799_v71 = _nw136
			var _800_v72 _dafny.Map
			_ = _800_v72
			_800_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_799_v71, false)
			var _801_v73 _dafny.Set
			_ = _801_v73
			_801_v73 = _dafny.SetOf(false, _787_cf0, !(_787_cf0))
			var _802_v74 _dafny.Set
			_ = _802_v74
			_802_v74 = _dafny.SetOf(_this.F16)
			var _803_v75 D4
			_ = _803_v75
			_803_v75 = Companion_D4_.Create_DC11_(_787_cf0, _this.F16, Companion_Default___.Fm1(_this.F12(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_this.F16)).Cardinality())), _this.F16, globalState), (_802_v74).Cardinality())
			var _804_v76 _dafny.Map
			_ = _804_v76
			_804_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v61, _dafny.IntOfUint32((_796_v68).Cardinality()))
			var _805_v77 *C0
			_ = _805_v77
			var _nw137 *C0 = New_C0_()
			_ = _nw137
			_nw137.Ctor__(_787_cf0, (_804_v76).Cardinality(), _this.F11(), (_794_v66).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (p0).Contains((_this.F11()).Cardinality()) {
					return (p0).Multiplicity((_this.F11()).Cardinality())
				}
				return _this.F16
			})(), _dafny.IntOfUint32((_794_v66).Cardinality()))).Uint32()).(bool))
			_805_v77 = _nw137
			var _806_v78 _dafny.Set
			_ = _806_v78
			_806_v78 = _dafny.SetOf(_805_v77, _805_v77, _805_v77, _805_v77)
			var _807_v79 _dafny.Array
			_ = _807_v79
			var _nwElement0_33 bool = (_795_v67).Contains(_793_v65)
			_ = _nwElement0_33
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(11))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_33, 0)
			(_nw138).ArraySet1((_this.F16).Cmp(_dafny.IntOfUint32((_796_v68).Cardinality())) <= 0, 1)
			(_nw138).ArraySet1((_dafny.IntOfInt64(-92)).Cmp((_797_v69).Cardinality()) == 0, 2)
			(_nw138).ArraySet1(_this.F12(), 3)
			(_nw138).ArraySet1((_798_v70).Contains(_789_v61), 4)
			(_nw138).ArraySet1((func() bool {
				if (_800_v72).Contains(_799_v71) {
					return (_800_v72).Get(_799_v71).(bool)
				}
				return _this.F12()
			})(), 5)
			(_nw138).ArraySet1(_787_cf0, 6)
			(_nw138).ArraySet1((_801_v73).IsSubsetOf(_dafny.SetOf((_803_v75).Dtor_cf25())), 7)
			(_nw138).ArraySet1(_787_cf0, 8)
			(_nw138).ArraySet1(_787_cf0, 9)
			(_nw138).ArraySet1(((_806_v78).Cardinality()).Cmp((_805_v77).F18()) > 0, 10)
			_807_v79 = _nw138
			var _808_v80 _dafny.Map
			_ = _808_v80
			_808_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_805_v77).F17(), _dafny.IntOfInt64(53))
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_807_v79), 0))
			_ = _index98
			(_807_v79).ArraySet1((_805_v77).Fm2(p0, _this.F16, _808_v80, (_805_v77).F17(), globalState), (_index98).Int())
			var _809_v81 _dafny.Map
			_ = _809_v81
			_809_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_805_v77).F17())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_788_v60), 0))
			_ = _index99
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_790_v62), 0))
			_ = _index100
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_807_v79), 0))
			_ = _index101
			var _rhs122 _dafny.Int = _this.F16
			_ = _rhs122
			var _rhs123 _dafny.Sequence = _789_v61
			_ = _rhs123
			var _rhs124 D1 = _792_v64
			_ = _rhs124
			var _rhs125 _dafny.Array = _807_v79
			_ = _rhs125
			var _rhs126 bool = !((func() bool {
				if (_809_v81).Contains(p1) {
					return (_809_v81).Get(p1).(bool)
				}
				return _787_cf0
			})())
			_ = _rhs126
			var _lhs96 _dafny.Array = _788_v60
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_788_v60), 0))
			_ = _lhs97
			var _lhs98 _dafny.Array = _790_v62
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(188), _dafny.ArrayLen((_790_v62), 0))
			_ = _lhs99
			var _lhs100 _dafny.Array = _807_v79
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_807_v79), 0))
			_ = _lhs101
			(_lhs96).ArraySet1(_rhs122, (_lhs97).Int())
			_789_v61 = _rhs123
			(_lhs98).ArraySet1(_rhs124, (_lhs99).Int())
			_807_v79 = _rhs125
			(_lhs100).ArraySet1(_rhs126, (_lhs101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_788_v60), 0))
			_ = _index102
			(_788_v60).ArraySet1(_dafny.IntOfInt64(-18), (_index102).Int())
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_785_v59), 0))
			_ = _index103
			(_785_v59).ArraySet1(p1, (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_785_v59), 0))
			_ = _index104
			(_785_v59).ArraySet1(Companion_Default___.Fm9((_788_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_788_v60), 0))).Int()).(_dafny.Int), ((_dafny.Zero).Minus((_788_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_788_v60), 0))).Int()).(_dafny.Int))).Cmp((func() _dafny.Int {
				if (_808_v80).Contains((_805_v77).F17()) {
					return (_808_v80).Get((_805_v77).F17()).(_dafny.Int)
				}
				return _this.F16
			})()) >= 0, globalState), (_index104).Int())
			_697_v2 = _697_v2
			var _810_v82 _dafny.Array
			_ = _810_v82
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
			_ = _nw139
			_810_v82 = _nw139
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_810_v82), 0))
			_ = _index105
			(_810_v82).ArraySet1CodePoint(_696_v1, (_index105).Int())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(113), _dafny.ArrayLen((_810_v82), 0))
			_ = _index106
			(_810_v82).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _this.F12() {
					return _696_v1
				}
				return _696_v1
			})(), (_index106).Int())
		}
		var _hi1 _dafny.Int = _this.F16
		_ = _hi1
		for _811_i9 := _this.F16; _811_i9.Cmp(_hi1) < 0; _811_i9 = _811_i9.Plus(_dafny.One) {
			var _812_v83 _dafny.Map
			_ = _812_v83
			_812_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}(func(_813_i10 _dafny.Int) _dafny.Int {
				return _this.F16
			}))).Cardinality()), _this.F12())
			var _814_v84 D4
			_ = _814_v84
			_814_v84 = Companion_D4_.Create_DC11_(!(false), _811_i9, _811_i9, (_812_v83).Cardinality())
			var _815_v85 *C0
			_ = _815_v85
			var _nw140 *C0 = New_C0_()
			_ = _nw140
			_nw140.Ctor__(((_814_v84).Dtor_cf25()) == (_this.F12()), (_dafny.Zero).Minus(_811_i9), _this.F11(), false)
			_815_v85 = _nw140
			var _816_v86 _dafny.Array
			_ = _816_v86
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
			_ = _nw141
			_816_v86 = _nw141
			var _817_v87 _dafny.Sequence
			_ = _817_v87
			_817_v87 = _dafny.SeqOf(_this.F16, (_815_v85).F18())
			var _818_v88 _dafny.Sequence
			_ = _818_v88
			_818_v88 = _dafny.SeqOf(_dafny.IntOfUint32((_817_v87).Cardinality()))
			var _819_v89 _dafny.Map
			_ = _819_v89
			_819_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _818_v88)
			var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_816_v86), 0))
			_ = _index107
			(_816_v86).ArraySet1((_819_v89).Merge(_819_v89), (_index107).Int())
			var _820_v90 _dafny.Array
			_ = _820_v90
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw142
			_820_v90 = _nw142
			var _821_v91 _dafny.Sequence
			_ = _821_v91
			_821_v91 = _dafny.SeqOf(_820_v90, _820_v90)
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(225), _dafny.ArrayLen((_816_v86), 0))
			_ = _index108
			(_816_v86).ArraySet1((func() _dafny.Map {
				if _dafny.Companion_Sequence_.Equal(_821_v91, _821_v91) {
					return (func() _dafny.Map {
						if _this.F12() {
							return _819_v89
						}
						return _819_v89
					})()
				}
				return (_819_v89).Merge(_819_v89)
			})(), (_index108).Int())
			(_this).F12_set_(((_815_v85).F17()) && (((_815_v85).F17()) && (_this.F12())))
			var _822_v92 _dafny.Map
			_ = _822_v92
			_822_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_815_v85).F18(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_823_v85 *C0) func(_dafny.Int) _dafny.Int {
				return func(_824_i11 _dafny.Int) _dafny.Int {
					return (_823_v85).F18()
				}
			})(_815_v85))))
			var _825_v93 _dafny.Set
			_ = _825_v93
			_825_v93 = _dafny.SetOf(_817_v87, _817_v87, (func() _dafny.Sequence {
				if (_822_v92).Contains(_811_i9) {
					return (_822_v92).Get(_811_i9).(_dafny.Sequence)
				}
				return _818_v88
			})())
			var _826_v94 _dafny.Map
			_ = _826_v94
			_826_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_825_v93, _this.F12())
			_826_v94 = (_826_v94).Update(_825_v93, (_815_v85).F17())
		}
		var _827_v95 _dafny.Array
		_ = _827_v95
		var _nw143 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw143
		_827_v95 = _nw143
		var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
		_ = _index109
		(_827_v95).ArraySet1(_this.F12(), (_index109).Int())
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
		_ = _index110
		(_827_v95).ArraySet1((_this.F11()).IsProperSubsetOf((_this.F11()).Union(_dafny.MultiSetOf(_this.F12(), true))), (_index110).Int())
		var _828_v96 _dafny.Array
		_ = _828_v96
		var _nwElement0_34 _dafny.CodePoint = _dafny.CodePoint('b')
		_ = _nwElement0_34
		var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(16))
		_ = _nw144
		(_nw144).ArraySet1CodePoint(_nwElement0_34, 0)
		(_nw144).ArraySet1CodePoint(_696_v1, 1)
		(_nw144).ArraySet1CodePoint(_696_v1, 2)
		(_nw144).ArraySet1CodePoint(_696_v1, 3)
		(_nw144).ArraySet1CodePoint(_696_v1, 4)
		(_nw144).ArraySet1CodePoint(_696_v1, 5)
		(_nw144).ArraySet1CodePoint(_696_v1, 6)
		(_nw144).ArraySet1CodePoint(Companion_Default___.Fm14(_697_v2, (_dafny.Zero).Minus(_this.F16), globalState), 7)
		(_nw144).ArraySet1CodePoint(_696_v1, 8)
		(_nw144).ArraySet1CodePoint(_696_v1, 9)
		(_nw144).ArraySet1CodePoint(_696_v1, 10)
		(_nw144).ArraySet1CodePoint(_696_v1, 11)
		(_nw144).ArraySet1CodePoint(_696_v1, 12)
		(_nw144).ArraySet1CodePoint(_696_v1, 13)
		(_nw144).ArraySet1CodePoint(_696_v1, 14)
		(_nw144).ArraySet1CodePoint(_696_v1, 15)
		_828_v96 = _nw144
		for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_828_v96), 0))); ; {
			_guard_loop_1, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _829_i12 _dafny.Int
			_829_i12 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_829_i12).Sign() != -1) && ((_829_i12).Cmp(_dafny.ArrayLen((_828_v96), 0)) < 0)) {
				(_828_v96).ArraySet1CodePoint(_696_v1, (_829_i12).Int())
			}
		}
		var _830_v97 _dafny.Array
		_ = _830_v97
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_19
		var _nw145 _dafny.Array
		_ = _nw145
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw145 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) _dafny.Sequence = (func(_831_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_832_i13 _dafny.Int) _dafny.Sequence {
					return _831_p1
				}
			})(p1)
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
		_830_v97 = _nw145
		var _833_v98 D5
		_ = _833_v98
		_833_v98 = Companion_D5_.Create_DC13_(_this.F16, (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F16, _830_v97, (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))
		var _834_v99 _dafny.Sequence
		_ = _834_v99
		_834_v99 = _dafny.SeqOf((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F12(), false)
		var _835_v100 _dafny.Array
		_ = _835_v100
		var _nwElement0_35 _dafny.MultiSet = p0
		_ = _nwElement0_35
		var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(3))
		_ = _nw146
		(_nw146).ArraySet1(_nwElement0_35, 0)
		(_nw146).ArraySet1(p0, 1)
		(_nw146).ArraySet1(((p0).Update((_833_v98).Dtor_cf32(), Companion_Default___.Abs((_dafny.Zero).Minus(_this.F16)))).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_834_v99, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.IntOfUint32((_834_v99).Cardinality()))).Uint32(), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))).Cardinality()), Companion_Default___.Abs(_this.F16)), 2)
		_835_v100 = _nw146
		var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))
		_ = _index111
		(_835_v100).ArraySet1((p0).Difference(p0), (_index111).Int())
		var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))
		_ = _index112
		(_835_v100).ArraySet1((p0).Intersection(p0), (_index112).Int())
		if (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool) {
			(_this).F12_set_((Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))).Cardinality(), _this.F16)).Cmp(_this.F16) <= 0)
			(globalState).F9 = _this.F16
			(globalState).F7 = Companion_Default___.Fm0(_696_v1, globalState)
			var _836_v101 D3
			_ = _836_v101
			_836_v101 = Companion_D3_.Create_DC6_(_827_v95)
			var _source9 D3 = _836_v101
			_ = _source9
			if _source9.Is_DC7() {
				var _837___mcc_h8 _dafny.Int = _source9.Get_().(D3_DC7).Cf15
				_ = _837___mcc_h8
				var _838___mcc_h9 _dafny.Int = _source9.Get_().(D3_DC7).Cf16
				_ = _838___mcc_h9
				var _839___mcc_h10 *C0 = _source9.Get_().(D3_DC7).Cf17
				_ = _839___mcc_h10
				var _840_cf17 *C0 = _839___mcc_h10
				_ = _840_cf17
				var _841_cf16 _dafny.Int = _838___mcc_h9
				_ = _841_cf16
				var _842_cf15 _dafny.Int = _837___mcc_h8
				_ = _842_cf15
				(globalState).F7 = true
				var _rhs127 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eokqyem")).Cardinality()), _841_cf16))
				_ = _rhs127
				var _rhs128 _dafny.Int = Companion_Default___.SafeModuloInt((_841_cf16).Plus(_842_cf15), _this.F16)
				_ = _rhs128
				var _rhs129 *C0 = _840_cf17
				_ = _rhs129
				var _rhs130 _dafny.Int = (_dafny.Zero).Minus((_840_cf17).F18())
				_ = _rhs130
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				var _lhs103 *GlobalState = globalState
				_ = _lhs103
				_842_cf15 = _rhs127
				_lhs102.F9 = _rhs128
				_840_cf17 = _rhs129
				_lhs103.F9 = _rhs130
				var _843_v102 _dafny.Map
				_ = _843_v102
				_843_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_842_cf15, _this.F16)
				var _844_v103 _dafny.Set
				_ = _844_v103
				_844_v103 = _dafny.SetOf(_842_cf15, (_840_cf17).F18(), _841_cf16, (_dafny.IntOfInt64(228)).Times((_843_v102).Cardinality()), (_840_cf17).Fm12(globalState))
				_844_v103 = _844_v103
				var _845_v104 _dafny.Set
				_ = _845_v104
				_845_v104 = _dafny.SetOf((_840_cf17).F17())
				var _846_v105 _dafny.Map
				_ = _846_v105
				_846_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F12())
				var _847_v106 _dafny.Set
				_ = _847_v106
				_847_v106 = _dafny.SetOf(_this.F12(), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), !(_this.F12()), !(Companion_Default___.Fm16((_845_v104).Cardinality(), globalState)).Equals(_846_v105))
				var _848_v107 _dafny.Sequence
				_ = _848_v107
				_848_v107 = _dafny.SeqOf(_842_cf15, ((_835_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))).Int()).(_dafny.MultiSet)).Cardinality())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
				_ = _index113
				var _rhs131 bool = ((_844_v103).Difference(_844_v103)).IsProperSubsetOf(func() _dafny.Set {
					var _coll20 = _dafny.NewBuilder()
					_ = _coll20
					for _iter22 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_848_v107, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_848_v107).Cardinality()))).Uint32(), _842_cf15)).Elements()); ; {
						_compr_20, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _849_v108 _dafny.Int
						_849_v108 = interface{}(_compr_20).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_848_v107, (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_848_v107).Cardinality()))).Uint32(), _842_cf15), _849_v108) {
							_coll20.Add((_849_v108).Minus(_dafny.IntOfInt64(992)))
						}
					}
					return _coll20.ToSet()
				}())
				_ = _rhs131
				var _rhs132 _dafny.Int = (_843_v102).Cardinality()
				_ = _rhs132
				var _rhs133 _dafny.Set = (_847_v106).Intersection(_dafny.SetOf((_840_cf17).F17()))
				_ = _rhs133
				var _rhs134 bool = (_845_v104).IsSubsetOf(_dafny.SetOf(_this.F12()))
				_ = _rhs134
				var _lhs104 _dafny.Array = _827_v95
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
				_ = _lhs105
				var _lhs106 *GlobalState = globalState
				_ = _lhs106
				(_lhs104).ArraySet1(_rhs131, (_lhs105).Int())
				_841_cf16 = _rhs132
				_847_v106 = _rhs133
				_lhs106.F8 = _rhs134
			} else if _source9.Is_DC8() {
				var _850___mcc_h11 _dafny.Int = _source9.Get_().(D3_DC8).Cf18
				_ = _850___mcc_h11
				var _851___mcc_h12 _dafny.Int = _source9.Get_().(D3_DC8).Cf19
				_ = _851___mcc_h12
				var _852_cf19 _dafny.Int = _851___mcc_h12
				_ = _852_cf19
				var _853_cf18 _dafny.Int = _850___mcc_h11
				_ = _853_cf18
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_830_v97), 0))
				_ = _index114
				(_830_v97).ArraySet1(p1, (_index114).Int())
				var _pat_let_tv16 = _827_v95
				_ = _pat_let_tv16
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_830_v97), 0))
				_ = _index115
				var _rhs135 _dafny.Sequence = p1
				_ = _rhs135
				var _rhs136 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if !(!(false)) {
						return _dafny.SeqOf((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F12())
					}
					return _834_v99
				})(), _834_v99)
				_ = _rhs136
				var _rhs137 bool = !(Companion_Default___.Fm0(_696_v1, globalState))
				_ = _rhs137
				var _rhs138 bool = _this.F12()
				_ = _rhs138
				var _rhs139 _dafny.Array = (func(_pat_let21_0 D3) D3 {
					return func(_854_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let22_0 _dafny.Array) D3 {
							return func(_855_dt__update_hcf14_h0 _dafny.Array) D3 {
								return Companion_D3_.Create_DC6_(_855_dt__update_hcf14_h0)
							}(_pat_let22_0)
						}(_pat_let_tv16)
					}(_pat_let21_0)
				}(_836_v101)).Dtor_cf14()
				_ = _rhs139
				var _lhs107 _dafny.Array = _830_v97
				_ = _lhs107
				var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.ArrayLen((_830_v97), 0))
				_ = _lhs108
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				(_lhs107).ArraySet1(_rhs135, (_lhs108).Int())
				_lhs109.F6 = _rhs136
				_lhs110.F7 = _rhs137
				_lhs111.F3 = _rhs138
				_827_v95 = _rhs139
				var _856_v109 T0
				_ = _856_v109
				var _nw147 *C0 = New_C0_()
				_ = _nw147
				_nw147.Ctor__(false, _this.F16, _this.F11(), (_834_v99).Select((Companion_Default___.SafeIndex(_852_cf19, _dafny.IntOfUint32((_834_v99).Cardinality()))).Uint32()).(bool))
				_856_v109 = _nw147
				var _857_v110 _dafny.Map
				_ = _857_v110
				_857_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_852_cf19, _856_v109)
				var _858_v111 _dafny.Map
				_ = _858_v111
				_858_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F16)
				(globalState).F7 = (_this).Fm2(Companion_Default___.Fm13(_696_v1, (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), globalState), Companion_Default___.SafeModuloInt((_857_v110).Cardinality(), _853_cf18), (_858_v111).Merge(_858_v111), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), globalState)
				var _859_v112 _dafny.Sequence
				_ = _859_v112
				_859_v112 = _dafny.SeqOf(_853_cf18, _853_cf18)
				var _860_v113 _dafny.Array
				_ = _860_v113
				var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw148
				_860_v113 = _nw148
				var _861_v114 _dafny.Map
				_ = _861_v114
				_861_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if _this.F12() {
						return _859_v112
					}
					return _859_v112
				})(), (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _this.F12() {
						return _859_v112
					}
					return _859_v112
				})()).Cardinality()))).Uint32(), _this.F16)), _860_v113)
				_861_v114 = (_861_v114).Update((p0).Difference((_835_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))).Int()).(_dafny.MultiSet)), _860_v113)
				var _862_v115 *C0
				_ = _862_v115
				var _nw149 *C0 = New_C0_()
				_ = _nw149
				_nw149.Ctor__((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F16, (_856_v109.F11()).Union(_this.F11()), ((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)) || ((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)))
				_862_v115 = _nw149
				_862_v115 = _862_v115
			} else if _source9.Is_DC9() {
				var _863___mcc_h13 _dafny.Int = _source9.Get_().(D3_DC9).Cf20
				_ = _863___mcc_h13
				var _864___mcc_h14 _dafny.Int = _source9.Get_().(D3_DC9).Cf21
				_ = _864___mcc_h14
				var _865___mcc_h15 *C0 = _source9.Get_().(D3_DC9).Cf22
				_ = _865___mcc_h15
				var _866___mcc_h16 _dafny.Int = _source9.Get_().(D3_DC9).Cf23
				_ = _866___mcc_h16
				var _867_cf23 _dafny.Int = _866___mcc_h16
				_ = _867_cf23
				var _868_cf22 *C0 = _865___mcc_h15
				_ = _868_cf22
				var _869_cf21 _dafny.Int = _864___mcc_h14
				_ = _869_cf21
				var _870_cf20 _dafny.Int = _863___mcc_h13
				_ = _870_cf20
				var _871_v116 _dafny.Array
				_ = _871_v116
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw150
				_871_v116 = _nw150
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_871_v116), 0))
				_ = _index116
				(_871_v116).ArraySet1(_this.F16, (_index116).Int())
				var _872_v117 _dafny.Map
				_ = _872_v117
				_872_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F12())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_871_v116), 0))
				_ = _index117
				(_871_v116).ArraySet1(Companion_Default___.SafeModuloInt((_872_v117).Cardinality(), (_868_cf22).F18()), (_index117).Int())
				var _873_v118 D1
				_ = _873_v118
				_873_v118 = Companion_D1_.Create_DC2_(_this.F12(), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F12(), _696_v1, _696_v1)
				var _874_v119 D1
				_ = _874_v119
				_874_v119 = Companion_D1_.Create_DC3_(_873_v118)
				var _875_v120 D1
				_ = _875_v120
				_875_v120 = Companion_D1_.Create_DC3_(_874_v119)
				var _876_v121 D1
				_ = _876_v121
				_876_v121 = Companion_D1_.Create_DC3_(_873_v118)
				var _pat_let_tv17 = _873_v118
				_ = _pat_let_tv17
				var _877_v122 _dafny.Set
				_ = _877_v122
				_877_v122 = _dafny.SetOf(func(_pat_let23_0 D1) D1 {
					return func(_878_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let24_0 D1) D1 {
							return func(_879_dt__update_hcf7_h0 D1) D1 {
								return Companion_D1_.Create_DC3_(_879_dt__update_hcf7_h0)
							}(_pat_let24_0)
						}(_pat_let_tv17)
					}(_pat_let23_0)
				}(Companion_D1_.Create_DC3_(_875_v120)), _876_v121)
				var _880_v123 _dafny.MultiSet
				_ = _880_v123
				_880_v123 = _dafny.MultiSetOf((_877_v122).Union(_877_v122), _877_v122)
				_880_v123 = (_880_v123).Difference((_880_v123).Difference(Companion_Default___.Fm17((_868_cf22).F17(), !(false), (_868_cf22).F18(), _dafny.IntOfInt64(190), globalState)))
				_870_cf20 = (_dafny.Zero).Minus(_870_cf20)
				var _881_v124 _dafny.Map
				_ = _881_v124
				_881_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _827_v95)
				_881_v124 = (_881_v124).Update((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _827_v95)
			} else {
				var _882___mcc_h17 _dafny.Array = _source9.Get_().(D3_DC6).Cf14
				_ = _882___mcc_h17
				var _883_cf14 _dafny.Array = _882___mcc_h17
				_ = _883_cf14
				var _884_v125 _dafny.Set
				_ = _884_v125
				_884_v125 = _dafny.SetOf(_this.F16)
				var _885_v126 _dafny.Map
				_ = _885_v126
				_885_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F16)
				var _886_v127 _dafny.Set
				_ = _886_v127
				_886_v127 = _dafny.SetOf(_this.F12(), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))
				var _887_v128 D2
				_ = _887_v128
				_887_v128 = Companion_D2_.Create_DC4_(_this.F16)
				var _888_v129 _dafny.Array
				_ = _888_v129
				var _nwElement0_36 _dafny.Int = _this.F16
				_ = _nwElement0_36
				var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(6))
				_ = _nw151
				(_nw151).ArraySet1(_nwElement0_36, 0)
				(_nw151).ArraySet1(_this.F16, 1)
				(_nw151).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_834_v99).Cardinality()), _this.F16, _this.F16, _this.F16, _834_v99)).Dtor_cf9(), _this.F12())).Cardinality()), 2)
				(_nw151).ArraySet1(((_884_v125).Difference(_dafny.SetOf((_dafny.Zero).Minus(_this.F16)))).Cardinality(), 3)
				(_nw151).ArraySet1((_dafny.IntOfInt64(-623)).Minus(((_885_v126).Update(!(_this.F12()), Companion_Default___.Fm1(!((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)), (_886_v127).Cardinality(), _this.F16, globalState))).Cardinality()), 4)
				(_nw151).ArraySet1(Companion_Default___.Fm1(_this.F12(), (_887_v128).Dtor_cf8(), _dafny.IntOfInt64(797), globalState), 5)
				_888_v129 = _nw151
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_888_v129), 0))
				_ = _index118
				(_888_v129).ArraySet1(_this.F16, (_index118).Int())
				var _889_v130 _dafny.Map
				_ = _889_v130
				_889_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))
				var _890_v131 _dafny.MultiSet
				_ = _890_v131
				_890_v131 = _dafny.MultiSetOf(_696_v1, (_697_v2).Dtor_cf5(), _696_v1, (p1).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _891_v132 _dafny.Sequence
				_ = _891_v132
				_891_v132 = _dafny.SeqOf(_this.F16, ((_889_v130).Merge(_889_v130)).Cardinality(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_834_v99, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.IntOfUint32((_834_v99).Cardinality()))).Uint32(), (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))).Cardinality()), _this.F16), _this.F16, ((_890_v131).Difference(_890_v131)).Cardinality())
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_888_v129), 0))
				_ = _index119
				(_888_v129).ArraySet1((_891_v132).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_891_v132).Cardinality()))).Uint32()).(_dafny.Int), (_index119).Int())
				var _892_v133 *C0
				_ = _892_v133
				var _nw152 *C0 = New_C0_()
				_ = _nw152
				_nw152.Ctor__(_this.F12(), _this.F16, (_dafny.MultiSetOf((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), _this.F12())).Union(_this.F11()), (_this.F16).Cmp(_this.F16) == 0)
				_892_v133 = _nw152
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_888_v129), 0))
				_ = _index120
				(_888_v129).ArraySet1(_this.F16, (_index120).Int())
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_888_v129), 0))
				_ = _index121
				(_888_v129).ArraySet1((_888_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen((_888_v129), 0))).Int()).(_dafny.Int), (_index121).Int())
			}
			var _893_v135 _dafny.Set
			_ = _893_v135
			_893_v135 = _dafny.SetOf(_this.F16)
			var _894_v136 *C0
			_ = _894_v136
			var _nw153 *C0 = New_C0_()
			_ = _nw153
			_nw153.Ctor__(false, _this.F16, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)), (Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_dafny.SeqOf((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))).Cardinality()))).Uint32(), _this.F12())), (_893_v135).IsProperSubsetOf(func() _dafny.Set {
				var _coll21 = _dafny.NewBuilder()
				_ = _coll21
				for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(447), _dafny.IntOfInt64(623))); ; {
					_compr_21, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _895_v134 _dafny.Int
					_895_v134 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(447)).Cmp(_895_v134) <= 0) && ((_895_v134).Cmp(_dafny.IntOfInt64(623)) < 0) {
						_coll21.Add(Companion_Default___.SafeModuloInt(_895_v134, _this.F16))
					}
				}
				return _coll21.ToSet()
			}()))
			_894_v136 = _nw153
		} else {
			(globalState).F9 = _this.F16
			var _896_v137 _dafny.Map
			_ = _896_v137
			_896_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)), _this.F16)
			(globalState).F3 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-201), (func() _dafny.Int {
				if (_896_v137).Contains((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)) {
					return (_896_v137).Get((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)).(_dafny.Int)
				}
				return _this.F16
			})())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(606))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}(func(_897_i14 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_this.F16)
			}))).Cardinality())) != 0
			if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F16)).Contains((_this.F16).Cmp(Companion_Default___.Fm1(_this.F12(), _this.F16, _this.F16, globalState)) == 0) {
				var _898_v138 _dafny.Set
				_ = _898_v138
				_898_v138 = _dafny.SetOf(_this.F12(), _this.F12())
				var _899_v139 *C0
				_ = _899_v139
				var _nw154 *C0 = New_C0_()
				_ = _nw154
				_nw154.Ctor__(_this.F12(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, (_898_v138).Cardinality())).Cardinality(), (_this.F11()).Update(_this.F12(), Companion_Default___.Abs(_dafny.IntOfUint32((p1).Cardinality()))), _this.F12())
				_899_v139 = _nw154
				var _900_v140 D3
				_ = _900_v140
				_900_v140 = Companion_D3_.Create_DC9_(_this.F16, (Companion_Default___.Fm18(_this.F16, _this.F12(), globalState)).Cardinality(), _899_v139, (_899_v139).F18())
				var _901_v141 _dafny.Map
				_ = _901_v141
				_901_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16, Companion_Default___.SafeDivisionInt(_this.F16, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_900_v140, (_899_v139).F17())).Cardinality()))
				_901_v141 = ((_901_v141).Merge(_901_v141)).Merge(_901_v141)
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
				_ = _index122
				(_827_v95).ArraySet1(true, (_index122).Int())
				var _902_v142 _dafny.Array
				_ = _902_v142
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_20
				var _nw155 _dafny.Array
				_ = _nw155
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw155 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Int = func(_903_i15 _dafny.Int) _dafny.Int {
						return (_903_i15).Times(_dafny.IntOfInt64(263))
					}
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw155 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw155).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw155).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_902_v142 = _nw155
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_902_v142), 0))
				_ = _index123
				(_902_v142).ArraySet1(_this.F16, (_index123).Int())
				var _904_v143 D3
				_ = _904_v143
				_904_v143 = Companion_D3_.Create_DC6_(_827_v95)
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_902_v142), 0))
				_ = _index124
				(_902_v142).ArraySet1(_this.F16, (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_902_v142), 0))
				_ = _index125
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_902_v142), 0))
				_ = _index126
				var _rhs140 _dafny.CodePoint = _696_v1
				_ = _rhs140
				var _rhs141 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_this.F16, (_this.F16).Plus((_this.F11()).Cardinality())))
				_ = _rhs141
				var _rhs142 D3 = _904_v143
				_ = _rhs142
				var _rhs143 _dafny.Int = _this.F16
				_ = _rhs143
				var _lhs112 _dafny.Array = _902_v142
				_ = _lhs112
				var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(122), _dafny.ArrayLen((_902_v142), 0))
				_ = _lhs113
				var _lhs114 _dafny.Array = _902_v142
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_902_v142), 0))
				_ = _lhs115
				_696_v1 = _rhs140
				(_lhs112).ArraySet1(_rhs141, (_lhs113).Int())
				_904_v143 = _rhs142
				(_lhs114).ArraySet1(_rhs143, (_lhs115).Int())
				(globalState).F3 = (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
				_ = _index127
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))
				_ = _index128
				var _rhs144 bool = _this.F12()
				_ = _rhs144
				var _rhs145 _dafny.MultiSet = p0
				_ = _rhs145
				var _rhs146 bool = (_899_v139).Fm2(Companion_Default___.Fm13(_696_v1, false, globalState), (_dafny.Zero).Minus(_this.F16), _896_v137, !(_this.F12()) || ((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)), globalState)
				_ = _rhs146
				var _rhs147 bool = (_899_v139).F17()
				_ = _rhs147
				var _lhs116 _dafny.Array = _827_v95
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))
				_ = _lhs117
				var _lhs118 _dafny.Array = _835_v100
				_ = _lhs118
				var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))
				_ = _lhs119
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				var _lhs121 *C1 = _this
				_ = _lhs121
				(_lhs116).ArraySet1(_rhs144, (_lhs117).Int())
				(_lhs118).ArraySet1(_rhs145, (_lhs119).Int())
				_lhs120.F7 = _rhs146
				_lhs121.F12_set_(_rhs147)
			} else {
				var _905_v144 _dafny.Array
				_ = _905_v144
				var _nw156 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
				_ = _nw156
				_905_v144 = _nw156
				var _906_v145 D3
				_ = _906_v145
				_906_v145 = Companion_D3_.Create_DC6_(_827_v95)
				var _pat_let_tv18 = _827_v95
				_ = _pat_let_tv18
				var _pat_let_tv19 = _827_v95
				_ = _pat_let_tv19
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_905_v144), 0))
				_ = _index129
				(_905_v144).ArraySet1(func(_pat_let25_0 D3) D3 {
					return func(_909_dt__update__tmp_h4 D3) D3 {
						return func(_pat_let28_0 _dafny.Array) D3 {
							return func(_910_dt__update_hcf14_h2 _dafny.Array) D3 {
								return Companion_D3_.Create_DC6_(_910_dt__update_hcf14_h2)
							}(_pat_let28_0)
						}(_pat_let_tv19)
					}(_pat_let25_0)
				}(func(_pat_let26_0 D3) D3 {
					return func(_907_dt__update__tmp_h3 D3) D3 {
						return func(_pat_let27_0 _dafny.Array) D3 {
							return func(_908_dt__update_hcf14_h1 _dafny.Array) D3 {
								return Companion_D3_.Create_DC6_(_908_dt__update_hcf14_h1)
							}(_pat_let27_0)
						}(_pat_let_tv18)
					}(_pat_let26_0)
				}(_906_v145)), (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_905_v144), 0))
				_ = _index130
				(_905_v144).ArraySet1(_906_v145, (_index130).Int())
				var _911_v146 *C0
				_ = _911_v146
				var _nw157 *C0 = New_C0_()
				_ = _nw157
				_nw157.Ctor__(_this.F12(), ((_835_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_835_v100), 0))).Int()).(_dafny.MultiSet)).Cardinality(), (_this.F11()).Intersection(_this.F11()), !(!((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)) || ((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool))))
				_911_v146 = _nw157
				var _912_v147 T0
				_ = _912_v147
				var _nw158 *C0 = New_C0_()
				_ = _nw158
				_nw158.Ctor__((func() bool {
					if (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool) {
						return (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)
					}
					return _this.F12()
				})(), ((_911_v146).F18()).Times(_this.F16), (_this.F11()).Update(_this.F12(), Companion_Default___.Abs((_911_v146).F18())), _this.F12())
				_912_v147 = _nw158
				_912_v147 = _912_v147
				var _913_v148 *C0
				_ = _913_v148
				var _nw159 *C0 = New_C0_()
				_ = _nw159
				_nw159.Ctor__((_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool), (_911_v146).F18(), (_912_v147.F11()).Union(_912_v147.F11()), _this.F12())
				_913_v148 = _nw159
				var _914_v149 _dafny.Sequence
				_ = _914_v149
				_914_v149 = _dafny.UnicodeSeqOfUtf8Bytes("ggsc")
				_914_v149 = p1
			}
			var _915_v150 _dafny.Sequence
			_ = _915_v150
			_915_v150 = _dafny.SeqOf(_827_v95)
			var _916_v151 _dafny.MultiSet
			_ = _916_v151
			_916_v151 = _dafny.MultiSetOf(_696_v1)
			var _917_v152 _dafny.Map
			_ = _917_v152
			_917_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _916_v151)
			var _918_v153 _dafny.Array
			_ = _918_v153
			var _nwElement0_37 _dafny.Array = _827_v95
			_ = _nwElement0_37
			var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(10))
			_ = _nw160
			(_nw160).ArraySet1(_nwElement0_37, 0)
			(_nw160).ArraySet1(_827_v95, 1)
			(_nw160).ArraySet1(_827_v95, 2)
			(_nw160).ArraySet1(_827_v95, 3)
			(_nw160).ArraySet1((_915_v150).Select((Companion_Default___.SafeIndex((_917_v152).Cardinality(), _dafny.IntOfUint32((_915_v150).Cardinality()))).Uint32()).(_dafny.Array), 4)
			(_nw160).ArraySet1(_827_v95, 5)
			(_nw160).ArraySet1(_827_v95, 6)
			(_nw160).ArraySet1((_915_v150).Select((Companion_Default___.SafeIndex(_this.F16, _dafny.IntOfUint32((_915_v150).Cardinality()))).Uint32()).(_dafny.Array), 7)
			(_nw160).ArraySet1(_827_v95, 8)
			(_nw160).ArraySet1(_827_v95, 9)
			_918_v153 = _nw160
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_918_v153), 0))
			_ = _index131
			(_918_v153).ArraySet1(_827_v95, (_index131).Int())
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_918_v153), 0))
			_ = _index132
			(_918_v153).ArraySet1(_827_v95, (_index132).Int())
			(globalState).F3 = (_827_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_827_v95), 0))).Int()).(bool)
		}
		r0 = _830_v97
		return r0
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f11 _dafny.MultiSet
	_f12 bool
	F15  bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f11 = _dafny.EmptyMultiSet
	_this._f12 = false
	_this.F15 = false
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

func (_this *C2) F11() _dafny.MultiSet {
	return _this._f11
}
func (_this *C2) F11_set_(value _dafny.MultiSet) {
	_this._f11 = value
}
func (_this *C2) F12() bool {
	return _this._f12
}
func (_this *C2) F12_set_(value bool) {
	_this._f12 = value
}
func (_this *C2) Ctor__(f15 bool, f11 _dafny.MultiSet, f12 bool) {
	{
		(_this).F15 = f15
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C2) Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return _this.F12()
	}
}
func (_this *C2) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(622)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(392))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg43 _dafny.Int) interface{} {
				return coer43(arg43)
			}
		}(func(_919_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(959)
		})))
	}
}
func (_this *C2) M4(p0 _dafny.Int, p1 bool, p2 D1, globalState *GlobalState) (bool, _dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		if (_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality())).Cmp(p0) == 0 {
			if p1 {
				var _920_v0 _dafny.Array
				_ = _920_v0
				var _nw161 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
				_ = _nw161
				_920_v0 = _nw161
				_920_v0 = (func() _dafny.Array {
					if (_this.F12()) && (p1) {
						return _920_v0
					}
					return _920_v0
				})()
				r2 = Companion_Default___.SafeDivisionInt(p0, p0)
				var _921_v1 _dafny.Sequence
				_ = _921_v1
				_921_v1 = _dafny.UnicodeSeqOfUtf8Bytes("nbffrd")
				(globalState).F7 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}(func(_922_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				})), _921_v1), _dafny.UnicodeSeqOfUtf8Bytes("xycvusgt"))
				var _923_v2 _dafny.Array
				_ = _923_v2
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw162
				_923_v2 = _nw162
				var _924_v3 _dafny.Map
				_ = _924_v3
				_924_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_923_v2, _this.F15)
				_924_v3 = (_924_v3).Update(_923_v2, _this.F15)
				var _925_v4 _dafny.Map
				_ = _925_v4
				_925_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _this.F12())
				var _926_v5 _dafny.Set
				_ = _926_v5
				_926_v5 = _dafny.SetOf(_920_v0)
				var _927_v6 _dafny.Map
				_ = _927_v6
				_927_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_925_v4).Cardinality(), (_926_v5).IsProperSubsetOf(_926_v5))
				var _rhs148 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_921_v1, _921_v1)
				_ = _rhs148
				var _rhs149 _dafny.Map = _927_v6
				_ = _rhs149
				var _rhs150 _dafny.Array = _920_v0
				_ = _rhs150
				var _rhs151 bool = _this.F12()
				_ = _rhs151
				var _rhs152 bool = !(_this.F15) || (_this.F12())
				_ = _rhs152
				var _lhs122 *GlobalState = globalState
				_ = _lhs122
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				_921_v1 = _rhs148
				_927_v6 = _rhs149
				_920_v0 = _rhs150
				_lhs122.F7 = _rhs151
				_lhs123.F3 = _rhs152
			} else {
				var _928_v7 _dafny.Sequence
				_ = _928_v7
				_928_v7 = _dafny.SeqOf(p0, p0)
				_928_v7 = _928_v7
				var _929_v8 _dafny.Array
				_ = _929_v8
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw163
				_929_v8 = _nw163
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_929_v8), 0))
				_ = _index133
				(_929_v8).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(23))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}(func(_930_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}(func(_931_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				}))), (_index133).Int())
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_929_v8), 0))
				_ = _index134
				(_929_v8).ArraySet1(_this.F12(), (_index134).Int())
				var _932_v9 _dafny.Array
				_ = _932_v9
				var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
				_ = _nw164
				_932_v9 = _nw164
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_932_v9), 0))
				_ = _index135
				(_932_v9).ArraySet1(Companion_Default___.Fm1(_this.F15, p0, p0, globalState), (_index135).Int())
				var _933_v10 _dafny.Map
				_ = _933_v10
				_933_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F12())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_932_v9), 0))
				_ = _index136
				(_932_v9).ArraySet1((Companion_Default___.SafeModuloInt(p0, p0)).Minus((_933_v10).Cardinality()), (_index136).Int())
				Companion_Default___.M0(globalState)
				var _934_v11 _dafny.Map
				_ = _934_v11
				_934_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_929_v8, _932_v9)
				_934_v11 = (_934_v11).Update(_929_v8, _932_v9)
			}
			(globalState).F9 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xswkq")).Cardinality()))
			(globalState).F9 = p0
			r2 = (func() _dafny.Int {
				if (_this.F11()).Contains(!(_this.F12()) || (_this.F15)) {
					return (_this.F11()).Multiplicity(!(_this.F12()) || (_this.F15))
				}
				return p0
			})()
			var _935_v12 _dafny.Set
			_ = _935_v12
			_935_v12 = _dafny.SetOf(_dafny.IntOfInt64(-479), p0, p0)
			var _936_v13 _dafny.Map
			_ = _936_v13
			_936_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1) || (_this.F12()), _dafny.SetOf((_935_v12).Cardinality()))
			_935_v12 = (func() _dafny.Set {
				if (_936_v13).Contains(false) {
					return (_936_v13).Get(false).(_dafny.Set)
				}
				return _935_v12
			})()
		} else {
			var _937_v14 _dafny.Sequence
			_ = _937_v14
			_937_v14 = _dafny.SeqOf(p0)
			var _938_v16 *C0
			_ = _938_v16
			var _nw165 *C0 = New_C0_()
			_ = _nw165
			_nw165.Ctor__(_this.F12(), p0, _this.F11(), !(_dafny.SetOf(_dafny.IntOfUint32((_937_v14).Cardinality()))).Equals(func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(279), _dafny.IntOfInt64(299))); ; {
					_compr_22, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _939_v15 _dafny.Int
					_939_v15 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(279)).Cmp(_939_v15) <= 0) && ((_939_v15).Cmp(_dafny.IntOfInt64(299)) < 0) {
						_coll22.Add(Companion_Default___.SafeModuloInt(_939_v15, p0))
					}
				}
				return _coll22.ToSet()
			}()))
			_938_v16 = _nw165
			var _940_v17 _dafny.Map
			_ = _940_v17
			_940_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_938_v16).F17(), p0)
			_940_v17 = (_940_v17).Update(p1, p0)
			var _941_v18 _dafny.Sequence
			_ = _941_v18
			_941_v18 = _dafny.SeqOf(_this.F15)
			var _942_v19 _dafny.Array
			_ = _942_v19
			var _nwElement0_38 bool = (_938_v16).F17()
			_ = _nwElement0_38
			var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(2))
			_ = _nw166
			(_nw166).ArraySet1(_nwElement0_38, 0)
			(_nw166).ArraySet1(!((_dafny.IntOfUint32((_941_v18).Cardinality())).Cmp((_938_v16).F18()) > 0), 1)
			_942_v19 = _nw166
			var _943_v20 _dafny.Map
			_ = _943_v20
			_943_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_938_v16).F18(), _this.F12())
			var _944_v21 _dafny.Map
			_ = _944_v21
			_944_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, (func() bool {
				if (_943_v20).Contains(_dafny.IntOfInt64(694)) {
					return (_943_v20).Get(_dafny.IntOfInt64(694)).(bool)
				}
				return p1
			})())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_942_v19), 0))
			_ = _index137
			(_942_v19).ArraySet1((func() bool {
				if (_944_v21).Contains(!(!(true))) {
					return (_944_v21).Get(!(!(true))).(bool)
				}
				return (_938_v16).F17()
			})(), (_index137).Int())
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_942_v19), 0))
			_ = _index138
			(_942_v19).ArraySet1(_this.F15, (_index138).Int())
			(globalState).F3 = (p0).Cmp(Companion_Default___.SafeDivisionInt(p0, p0)) == 0
			_940_v17 = (_940_v17).Update((_942_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_942_v19), 0))).Int()).(bool), (_938_v16).F18())
		}
		var _945_i3 _dafny.Int
		_ = _945_i3
		_945_i3 = _dafny.Zero
		{
			for p1 {
				{
					if (_945_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_945_i3 = (_945_i3).Plus(_dafny.One)
					var _946_v22 _dafny.Sequence
					_ = _946_v22
					_946_v22 = _dafny.SeqOf(p0)
					var _947_v23 _dafny.Map
					_ = _947_v23
					_947_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), true)
					_946_v22 = (_this).Fm8((func() bool {
						if (_947_v23).Contains(p0) {
							return (_947_v23).Get(p0).(bool)
						}
						return _this.F12()
					})(), p1, globalState)
					var _948_v24 _dafny.Set
					_ = _948_v24
					_948_v24 = _dafny.SetOf(_this.F12())
					var _949_v25 _dafny.CodePoint
					_ = _949_v25
					_949_v25 = _dafny.CodePoint('q')
					var _950_v26 _dafny.Map
					_ = _950_v26
					_950_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_948_v24).Cardinality()).Minus(p0), _949_v25)
					_950_v26 = _950_v26
					var _951_v27 _dafny.MultiSet
					_ = _951_v27
					_951_v27 = _dafny.MultiSetOf(_dafny.IntOfInt64(100), p0, p0, _dafny.IntOfInt64(494), p0)
					(globalState).F9 = ((func() _dafny.Int {
						if !(_this.F15) {
							return (_this.F11()).Cardinality()
						}
						return (func() _dafny.Int {
							if (_951_v27).Contains(p0) {
								return (_951_v27).Multiplicity(p0)
							}
							return _dafny.IntOfInt64(316)
						})()
					})()).Times(p0)
					var _952_v28 _dafny.Map
					_ = _952_v28
					_952_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_946_v22, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-740), _dafny.IntOfUint32((_946_v22).Cardinality()))).Uint32(), p0), _dafny.IntOfInt64(954))
					var _953_v30 _dafny.MultiSet
					_ = _953_v30
					_953_v30 = _dafny.MultiSetOf(_946_v22)
					var _954_v31 _dafny.Sequence
					_ = _954_v31
					_954_v31 = _dafny.SeqOf(_952_v28)
					_952_v28 = ((_952_v28).Merge(func() _dafny.Map {
						var _coll23 = _dafny.NewMapBuilder()
						_ = _coll23
						for _iter25 := _dafny.Iterate((_953_v30).Elements()); ; {
							_compr_23, _ok25 := _iter25()
							if !_ok25 {
								break
							}
							var _955_v29 _dafny.Sequence
							_955_v29 = interface{}(_compr_23).(_dafny.Sequence)
							if (_953_v30).Contains(_955_v29) {
								_coll23.Add(_955_v29, p0)
							}
						}
						return _coll23.ToMap()
					}())).Merge((_954_v31).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_954_v31).Cardinality()))).Uint32()).(_dafny.Map))
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		(_this).F15 = _this.F15
		var _956_v32 _dafny.Sequence
		_ = _956_v32
		_956_v32 = _dafny.SeqOf(_this.F12())
		var _957_v33 _dafny.Sequence
		_ = _957_v33
		_957_v33 = _dafny.SeqOf(Companion_Default___.Fm1(p1, p0, p0, globalState), p0, p0)
		var _958_v34 _dafny.Map
		_ = _958_v34
		_958_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this.F11()).Cardinality()).Times((func() _dafny.Int {
			if (_this.F11()).Contains((_956_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_957_v33).Cardinality()), _dafny.IntOfUint32((_956_v32).Cardinality()))).Uint32()).(bool)) {
				return (_this.F11()).Multiplicity((_956_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_957_v33).Cardinality()), _dafny.IntOfUint32((_956_v32).Cardinality()))).Uint32()).(bool))
			}
			return p0
		})()), p0)
		var _959_v35 _dafny.Sequence
		_ = _959_v35
		_959_v35 = _dafny.UnicodeSeqOfUtf8Bytes("hptlhvxc")
		var _960_v36 _dafny.CodePoint
		_ = _960_v36
		_960_v36 = _dafny.CodePoint('c')
		_958_v34 = (_958_v34).Update((func() _dafny.Int {
			if _this.F12() {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_959_v35, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_959_v35).Cardinality()))).Uint32(), _960_v36)).Cardinality())
			}
			return p0
		})(), _dafny.IntOfUint32((_959_v35).Cardinality()))
		var _961_v37 _dafny.Array
		_ = _961_v37
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_21
		var _nw167 _dafny.Array
		_ = _nw167
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw167 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) _dafny.Int = func(_962_i5 _dafny.Int) _dafny.Int {
				return (_962_i5).Plus(_dafny.IntOfInt64(758))
			}
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw167 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw167).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw167).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_961_v37 = _nw167
		for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_961_v37), 0))); ; {
			_guard_loop_2, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _963_i4 _dafny.Int
			_963_i4 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_963_i4).Sign() != -1) && ((_963_i4).Cmp(_dafny.ArrayLen((_961_v37), 0)) < 0)) {
				(_961_v37).ArraySet1((_963_i4).Times((Companion_D2_.Create_DC4_(_dafny.IntOfInt64(-115))).Dtor_cf8()), (_963_i4).Int())
			}
		}
		r2 = ((_dafny.Zero).Minus(p0)).Minus((p0).Plus((_dafny.Zero).Minus(p0)))
		var _964_v38 _dafny.MultiSet
		_ = _964_v38
		_964_v38 = _dafny.MultiSetOf(p0)
		var _965_v39 D6
		_ = _965_v39
		_965_v39 = Companion_D6_.Create_DC15_(_964_v38)
		var _966_v40 _dafny.Set
		_ = _966_v40
		_966_v40 = _dafny.SetOf(p1, true)
		var _967_v41 _dafny.Map
		_ = _967_v41
		_967_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), p0)
		r0 = (_this).Fm2(((_965_v39).Dtor_cf35()).Intersection(_dafny.MultiSetOf(p0, _dafny.IntOfUint32((_959_v35).Cardinality()), (_966_v40).Cardinality())), p0, _967_v41, _this.F15, globalState)
		r1 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg47 _dafny.Int) interface{} {
				return coer47(arg47)
			}
		}((func(_968_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_969_i6 _dafny.Int) _dafny.CodePoint {
				return _968_v36
			}
		})(_960_v36)))
		r2 = Companion_Default___.SafeModuloInt(p0, Companion_Default___.Fm1(p1, (func() _dafny.Int {
			if (_this.F11()).Contains(_this.F15) {
				return (_this.F11()).Multiplicity(_this.F15)
			}
			return _dafny.IntOfInt64(-55)
		})(), (_966_v40).Cardinality(), globalState))
		r3 = !(_this.F12())
		return r0, r1, r2, r3
	}
}
func (_this *C2) M5(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		r1 = p0
		var _970_v0 _dafny.CodePoint
		_ = _970_v0
		_970_v0 = _dafny.CodePoint('a')
		var _971_v1 D1
		_ = _971_v1
		_971_v1 = Companion_D1_.Create_DC2_(_this.F15, _this.F12(), true, _970_v0, _970_v0)
		var _972_v2 _dafny.Map
		_ = _972_v2
		_972_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, Companion_Default___.Fm14(_971_v1, _dafny.IntOfInt64(344), globalState))
		_972_v2 = (_972_v2).Update(_this.F15, _970_v0)
		var _973_v3 _dafny.Array
		_ = _973_v3
		var _nwElement0_39 bool = _this.F15
		_ = _nwElement0_39
		var _nw168 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(3))
		_ = _nw168
		(_nw168).ArraySet1(_nwElement0_39, 0)
		(_nw168).ArraySet1(_this.F12(), 1)
		(_nw168).ArraySet1(_this.F15, 2)
		_973_v3 = _nw168
		for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_973_v3), 0))); ; {
			_guard_loop_3, _ok27 := _iter27()
			if !_ok27 {
				break
			}
			var _974_i0 _dafny.Int
			_974_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_974_i0).Sign() != -1) && ((_974_i0).Cmp(_dafny.ArrayLen((_973_v3), 0)) < 0)) {
				(_973_v3).ArraySet1(!(_this.F15), (_974_i0).Int())
			}
		}
		var _975_v4 _dafny.Set
		_ = _975_v4
		_975_v4 = _dafny.SetOf(_this.F12(), _this.F15)
		var _976_v5 D1
		_ = _976_v5
		_976_v5 = Companion_D1_.Create_DC1_(_975_v4)
		var _977_v6 _dafny.Set
		_ = _977_v6
		_977_v6 = _dafny.SetOf(_976_v5)
		var _978_v7 _dafny.Map
		_ = _978_v7
		_978_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, (_977_v6).IsProperSubsetOf(_dafny.SetOf(Companion_D1_.Create_DC1_(_975_v4))))
		if (func() bool {
			if (_978_v7).Contains(_this.F15) {
				return (_978_v7).Get(_this.F15).(bool)
			}
			return _this.F15
		})() {
			(globalState).F9 = p0
			(_this).F12_set_(_this.F12())
			var _979_v8 _dafny.Map
			_ = _979_v8
			_979_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
			var _980_v9 _dafny.Sequence
			_ = _980_v9
			_980_v9 = _dafny.SeqOf(p1, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_981_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_979_v8).Contains(_this.F15) {
					return (_979_v8).Get(_this.F15).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}(func(_982_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))).Cardinality()))).Uint32(), _970_v0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}(func(_983_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			})), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_979_v8).Contains(_this.F15) {
					return (_979_v8).Get(_this.F15).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(591))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}(func(_984_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))).Cardinality()))).Uint32(), _970_v0)).Cardinality()))).Uint32(), _970_v0))
			var _985_v10 _dafny.Sequence
			_ = _985_v10
			_985_v10 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())), (_dafny.IntOfUint32((_980_v9).Cardinality())).Minus(p0), (func() _dafny.Int {
				if _this.F15 {
					return p0
				}
				return p0
			})(), p0, p0)
			var _986_v11 *C0
			_ = _986_v11
			var _nw169 *C0 = New_C0_()
			_ = _nw169
			_nw169.Ctor__(_this.F12(), p0, _this.F11(), _this.F15)
			_986_v11 = _nw169
			(globalState).F9 = (_985_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_986_v11, _986_v11)).Cardinality()), _dafny.IntOfUint32((_985_v10).Cardinality()))).Uint32()).(_dafny.Int)
			var _987_v12 _dafny.Sequence
			_ = _987_v12
			_987_v12 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_985_v10, _985_v10), _985_v10)
			_985_v10 = (_987_v12).Select((Companion_Default___.SafeIndex((_986_v11).F18(), _dafny.IntOfUint32((_987_v12).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _988_v13 _dafny.Set
			_ = _988_v13
			_988_v13 = _dafny.SetOf(_970_v0)
			var _989_v14 _dafny.Map
			_ = _989_v14
			_989_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC14_(), !(false))
			var _990_v15 D7
			_ = _990_v15
			_990_v15 = Companion_D7_.Create_DC17_(_989_v14)
			var _991_v16 _dafny.Array
			_ = _991_v16
			var _nwElement0_40 _dafny.Int = p0
			_ = _nwElement0_40
			var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(26))
			_ = _nw170
			(_nw170).ArraySet1(_nwElement0_40, 0)
			(_nw170).ArraySet1(p0, 1)
			(_nw170).ArraySet1((_986_v11).F18(), 2)
			(_nw170).ArraySet1(p0, 3)
			(_nw170).ArraySet1((_dafny.IntOfInt64(300)).Times((_986_v11).F18()), 4)
			(_nw170).ArraySet1((_dafny.IntOfInt64(631)).Minus((_986_v11).F18()), 5)
			(_nw170).ArraySet1((p0).Plus(p0), 6)
			(_nw170).ArraySet1(p0, 7)
			(_nw170).ArraySet1(p0, 8)
			(_nw170).ArraySet1(((_986_v11).F18()).Plus(p0), 9)
			(_nw170).ArraySet1((_975_v4).Cardinality(), 10)
			(_nw170).ArraySet1(_dafny.Zero, 11)
			(_nw170).ArraySet1(Companion_Default___.SafeModuloInt((_988_v13).Cardinality(), p0), 12)
			(_nw170).ArraySet1((_dafny.Zero).Minus(((_990_v15).Dtor_cf39()).Cardinality()), 13)
			(_nw170).ArraySet1(p0, 14)
			(_nw170).ArraySet1(p0, 15)
			(_nw170).ArraySet1(Companion_Default___.Fm1(_this.F12(), (_986_v11).F18(), (_986_v11).F18(), globalState), 16)
			(_nw170).ArraySet1((_986_v11).F18(), 17)
			(_nw170).ArraySet1((func() _dafny.Int {
				if (_this.F11()).Contains(_this.F12()) {
					return (_this.F11()).Multiplicity(_this.F12())
				}
				return (_986_v11).F18()
			})(), 18)
			(_nw170).ArraySet1((_986_v11).F18(), 19)
			(_nw170).ArraySet1((_986_v11).F18(), 20)
			(_nw170).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 21)
			(_nw170).ArraySet1(((_dafny.Zero).Minus((_986_v11).F18())).Times((_986_v11).F18()), 22)
			(_nw170).ArraySet1(((_986_v11).F18()).Plus((_this.F11()).Cardinality()), 23)
			(_nw170).ArraySet1((_986_v11).F18(), 24)
			(_nw170).ArraySet1((_986_v11).F18(), 25)
			_991_v16 = _nw170
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_991_v16), 0))
			_ = _index139
			(_991_v16).ArraySet1((_985_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_985_v10).Cardinality()))).Uint32()).(_dafny.Int), (_index139).Int())
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_991_v16), 0))
			_ = _index140
			(_991_v16).ArraySet1((_dafny.Zero).Minus((p0).Plus(_dafny.IntOfInt64(635))), (_index140).Int())
		} else {
			var _992_v17 _dafny.Array
			_ = _992_v17
			var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw171
			_992_v17 = _nw171
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_992_v17), 0))
			_ = _index141
			(_992_v17).ArraySet1((_this).Fm8(_this.F15, _this.F12(), globalState), (_index141).Int())
			var _993_v18 _dafny.Sequence
			_ = _993_v18
			_993_v18 = _dafny.SeqOf(p0)
			var _994_v19 D5
			_ = _994_v19
			_994_v19 = Companion_D5_.Create_DC12_(_993_v18)
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_992_v17), 0))
			_ = _index142
			(_992_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate((Companion_D5_.Create_DC12_(_993_v18)).Dtor_cf29(), (_994_v19).Dtor_cf29()), (_index142).Int())
			var _995_v20 _dafny.Array
			_ = _995_v20
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_22
			var _nw172 _dafny.Array
			_ = _nw172
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw172 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = (func(_996_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_997_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_997_i2, _996_p0)
					}
				})(p0)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw172 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw172).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw172).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_995_v20 = _nw172
			var _998_v21 _dafny.Map
			_ = _998_v21
			_998_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _995_v20)
			var _999_v22 _dafny.Sequence
			_ = _999_v22
			_999_v22 = _dafny.SeqOf((func() _dafny.Array {
				if (_998_v21).Contains(_this.F15) {
					return (_998_v21).Get(_this.F15).(_dafny.Array)
				}
				return _995_v20
			})())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_973_v3), 0))
			_ = _index143
			(_973_v3).ArraySet1(_this.F12(), (_index143).Int())
			var _1000_v23 _dafny.Sequence
			_ = _1000_v23
			_1000_v23 = _dafny.UnicodeSeqOfUtf8Bytes("muarlkhi")
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_973_v3), 0))
			_ = _index144
			var _rhs153 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_999_v22, _999_v22), _999_v22)
			_ = _rhs153
			var _rhs154 bool = (_978_v7).Equals(_978_v7)
			_ = _rhs154
			var _rhs155 _dafny.Sequence = p1
			_ = _rhs155
			var _lhs124 _dafny.Array = _973_v3
			_ = _lhs124
			var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_973_v3), 0))
			_ = _lhs125
			_999_v22 = _rhs153
			(_lhs124).ArraySet1(_rhs154, (_lhs125).Int())
			_1000_v23 = _rhs155
			var _1001_v24 _dafny.Map
			_ = _1001_v24
			_1001_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1000_v23)
			var _1002_v25 _dafny.MultiSet
			_ = _1002_v25
			_1002_v25 = _dafny.MultiSetOf(p0)
			var _rhs156 bool = _dafny.Companion_Sequence_.IsPrefixOf(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}(func(_1003_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})))
			_ = _rhs156
			var _rhs157 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_1002_v25).Contains(p0) {
					return (_1002_v25).Multiplicity(p0)
				}
				return p0
			})(), p0), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(p0))), _dafny.IntOfUint32((p1).Cardinality())))
			_ = _rhs157
			var _rhs158 _dafny.Map = ((_1001_v24).Update(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_1004_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))).Cardinality()), p1)).Merge(_1001_v24)
			_ = _rhs158
			var _lhs126 *GlobalState = globalState
			_ = _lhs126
			var _lhs127 *GlobalState = globalState
			_ = _lhs127
			_lhs126.F3 = _rhs156
			_lhs127.F9 = _rhs157
			_1001_v24 = _rhs158
			var _1005_v26 _dafny.Map
			_ = _1005_v26
			_1005_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_993_v18).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_993_v18).Cardinality()))).Uint32()).(_dafny.Int))
			var _1006_v27 _dafny.Set
			_ = _1006_v27
			_1006_v27 = _dafny.SetOf(_1005_v26, _1005_v26)
			var _1007_v28 _dafny.Map
			_ = _1007_v28
			_1007_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), Companion_Default___.Fm1(!(_this.F12()), p0, (_1006_v27).Cardinality(), globalState))
			_1007_v28 = _1007_v28
			var _1008_v29 _dafny.Map
			_ = _1008_v29
			_1008_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC15_(_1002_v25), _1002_v25)
			var _1009_v30 D6
			_ = _1009_v30
			_1009_v30 = Companion_D6_.Create_DC15_(_1002_v25)
			var _1010_v31 _dafny.Set
			_ = _1010_v31
			_1010_v31 = _dafny.SetOf((_1007_v28).Cardinality())
			var _1011_v32 D4
			_ = _1011_v32
			_1011_v32 = Companion_D4_.Create_DC11_((_this).Fm2((func() _dafny.MultiSet {
				if (_1008_v29).Contains(_1009_v30) {
					return (_1008_v29).Get(_1009_v30).(_dafny.MultiSet)
				}
				return _1002_v25
			})(), p0, _1007_v28, (_973_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_973_v3), 0))).Int()).(bool), globalState), p0, p0, Companion_Default___.Fm1((_973_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_973_v3), 0))).Int()).(bool), (_dafny.Zero).Minus((_1010_v31).Cardinality()), (_978_v7).Cardinality(), globalState))
			var _1012_v33 _dafny.Map
			_ = _1012_v33
			_1012_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_995_v20, (_1011_v32).Dtor_cf25())
			_1012_v33 = (_1012_v33).Update(_995_v20, _this.F12())
		}
		var _1013_v34 D8
		_ = _1013_v34
		_1013_v34 = Companion_D8_.Create_DC19_(_this.F11())
		var _1014_v35 _dafny.Sequence
		_ = _1014_v35
		_1014_v35 = _dafny.SeqOf(_this.F15)
		var _1015_v36 *C1
		_ = _1015_v36
		var _nw173 *C1 = New_C1_()
		_ = _nw173
		_nw173.Ctor__(_dafny.IntOfUint32((p1).Cardinality()), ((_1013_v34).Dtor_cf45()).Update((_1014_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-857), _dafny.IntOfUint32((_1014_v35).Cardinality()))).Uint32()).(bool), Companion_Default___.Abs(_dafny.IntOfInt64(-462))), _this.F12())
		_1015_v36 = _nw173
		(globalState).F3 = Companion_Default___.Fm0(_970_v0, globalState)
		var _1016_v37 _dafny.Set
		_ = _1016_v37
		_1016_v37 = _dafny.SetOf(p0)
		var _1017_v38 bool
		_ = _1017_v38
		_1017_v38 = false
		r0 = (func() bool {
			if (_dafny.SetOf(p0, _1015_v36.F16)).IsSubsetOf(_1016_v37) {
				return _1017_v38
			}
			return _1017_v38
		})()
		r1 = (_1015_v36.F16).Plus(p0)
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f11 _dafny.MultiSet
	_f12 bool
	_f13 _dafny.Int
	_f14 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f11 = _dafny.EmptyMultiSet
	_this._f12 = false
	_this._f13 = _dafny.Zero
	_this._f14 = false
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

func (_this *C3) F11() _dafny.MultiSet {
	return _this._f11
}
func (_this *C3) F11_set_(value _dafny.MultiSet) {
	_this._f11 = value
}
func (_this *C3) F12() bool {
	return _this._f12
}
func (_this *C3) F12_set_(value bool) {
	_this._f12 = value
}
func (_this *C3) Ctor__(f13 _dafny.Int, f14 bool, f11 _dafny.MultiSet, f12 bool) {
	{
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C3) Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return !((_this).F14()) || (_this.F12())
	}
}
func (_this *C3) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		if (func() bool {
			if true {
				return _this.F12()
			}
			return _this.F12()
		})() {
			if _this.F12() {
				return (_this).F14()
			} else {
				return !(_this.F12())
			}
		} else {
			return ((_this).F13()).Cmp((_this).F13()) < 0
		}
	}
}
func (_this *C3) M2(p0 _dafny.Sequence, p1 bool, p2 _dafny.Array, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		if (_this).F14() {
			var _1018_v0 _dafny.Array
			_ = _1018_v0
			var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
			_ = _nw174
			_1018_v0 = _nw174
			var _1019_v1 _dafny.CodePoint
			_ = _1019_v1
			_1019_v1 = _dafny.CodePoint('d')
			var _1020_v2 _dafny.MultiSet
			_ = _1020_v2
			_1020_v2 = _dafny.MultiSetOf(_dafny.CodePoint('a'), _1019_v1)
			var _1021_v3 _dafny.Sequence
			_ = _1021_v3
			_1021_v3 = _dafny.SeqOf(_1020_v2)
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_1018_v0), 0))
			_ = _index145
			(_1018_v0).ArraySet1((_1021_v3).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1021_v3).Cardinality()))).Uint32()).(_dafny.MultiSet), (_index145).Int())
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_1018_v0), 0))
			_ = _index146
			(_1018_v0).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Concatenate(p0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1019_v1, _1019_v1), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_dafny.SeqOf(_1019_v1, _1019_v1)).Cardinality()))).Uint32(), _1019_v1), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1019_v1, _1019_v1), (Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_dafny.SeqOf(_1019_v1, _1019_v1)).Cardinality()))).Uint32(), _1019_v1)).Cardinality()))).Uint32(), _dafny.CodePoint('i'))))), (_index146).Int())
			if p1 {
				var _1022_v4 _dafny.Set
				_ = _1022_v4
				_1022_v4 = _dafny.SetOf((_this).F13(), (_this).F13())
				var _1023_v5 _dafny.Array
				_ = _1023_v5
				var _nwElement0_41 bool = _this.F12()
				_ = _nwElement0_41
				var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(5))
				_ = _nw175
				(_nw175).ArraySet1(_nwElement0_41, 0)
				(_nw175).ArraySet1(((_this).F13()).Cmp((_this).F13()) != 0, 1)
				(_nw175).ArraySet1((_1022_v4).IsDisjointFrom(_1022_v4), 2)
				(_nw175).ArraySet1((_this).F14(), 3)
				(_nw175).ArraySet1(p1, 4)
				_1023_v5 = _nw175
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1023_v5), 0))
				_ = _index147
				(_1023_v5).ArraySet1((p1) == ((_this).F14()), (_index147).Int())
				var _1024_v6 _dafny.Sequence
				_ = _1024_v6
				_1024_v6 = _dafny.SeqOf(_1023_v5, _1023_v5, _1023_v5)
				var _1025_v7 _dafny.Map
				_ = _1025_v7
				_1025_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1024_v6, _this.F12())
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1023_v5), 0))
				_ = _index148
				(_1023_v5).ArraySet1((func() bool {
					if (_1025_v7).Contains(_1024_v6) {
						return (_1025_v7).Get(_1024_v6).(bool)
					}
					return (func() bool {
						if p1 {
							return false
						}
						return p1
					})()
				})(), (_index148).Int())
				var _1026_v8 _dafny.Array
				_ = _1026_v8
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw176
				_1026_v8 = _nw176
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1026_v8), 0))
				_ = _index149
				(_1026_v8).ArraySet1((_this).F13(), (_index149).Int())
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1026_v8), 0))
				_ = _index150
				(_1026_v8).ArraySet1((_this).F13(), (_index150).Int())
				var _1027_v9 _dafny.Map
				_ = _1027_v9
				_1027_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_1026_v8, _1026_v8, _1026_v8)).Cardinality()), p0)
				_1027_v9 = (_1027_v9).Update(((_1026_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1026_v8), 0))).Int()).(_dafny.Int)).Times((_this).F13()), p0)
				var _1028_v10 D1
				_ = _1028_v10
				_1028_v10 = Companion_D1_.Create_DC2_(true, (_this).F14(), (_1023_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_1023_v5), 0))).Int()).(bool), _1019_v1, _dafny.CodePoint('t'))
				var _1029_v11 D1
				_ = _1029_v11
				_1029_v11 = Companion_D1_.Create_DC3_(_1028_v10)
				var _1030_v12 _dafny.Map
				_ = _1030_v12
				_1030_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-76), _1029_v11)
				_1030_v12 = (_1030_v12).Update((_this).F13(), Companion_D1_.Create_DC3_(_1028_v10))
				var _1031_v13 bool
				_ = _1031_v13
				_1031_v13 = p1
				var _rhs159 bool = _1031_v13
				_ = _rhs159
				var _rhs160 bool = p1
				_ = _rhs160
				var _lhs128 *GlobalState = globalState
				_ = _lhs128
				_1031_v13 = _rhs159
				_lhs128.F7 = _rhs160
			} else {
				var _1032_v14 _dafny.Sequence
				_ = _1032_v14
				_1032_v14 = _dafny.SeqOf(_this.F12())
				var _1033_v15 D2
				_ = _1033_v15
				_1033_v15 = Companion_D2_.Create_DC5_((_this).F13(), (_this).F13(), (_this).F13(), (_this).F13(), _1032_v14)
				var _pat_let_tv20 = globalState
				_ = _pat_let_tv20
				r0 = (_dafny.IntOfUint32((p0).Cardinality())).Times((func(_pat_let29_0 D2) D2 {
					return func(_1034_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let30_0 _dafny.Sequence) D2 {
							return func(_1035_dt__update_hcf13_h0 _dafny.Sequence) D2 {
								return func(_pat_let31_0 _dafny.Int) D2 {
									return func(_1036_dt__update_hcf10_h0 _dafny.Int) D2 {
										return Companion_D2_.Create_DC5_((_1034_dt__update__tmp_h0).Dtor_cf9(), _1036_dt__update_hcf10_h0, (_1034_dt__update__tmp_h0).Dtor_cf11(), (_1034_dt__update__tmp_h0).Dtor_cf12(), _1035_dt__update_hcf13_h0)
									}(_pat_let31_0)
								}((_this).F13())
							}(_pat_let30_0)
						}(Companion_Default___.Fm7(_pat_let_tv20))
					}(_pat_let29_0)
				}(_1033_v15)).Dtor_cf10())
				(globalState).F9 = (_this).F13()
				var _1037_v16 *C1
				_ = _1037_v16
				var _nw177 *C1 = New_C1_()
				_ = _nw177
				_nw177.Ctor__((_this).F13(), _this.F11(), p1)
				_1037_v16 = _nw177
				var _1038_v17 _dafny.Int
				_ = _1038_v17
				var _out17 _dafny.Int
				_ = _out17
				_out17 = (_1037_v16).M6(globalState)
				_1038_v17 = _out17
				var _1039_v18 *C1
				_ = _1039_v18
				var _nw178 *C1 = New_C1_()
				_ = _nw178
				_nw178.Ctor__((_this).F13(), _this.F11(), (_this).F14())
				_1039_v18 = _nw178
			}
			if !((_this).F14()) {
				var _1040_v19 _dafny.Sequence
				_ = _1040_v19
				_1040_v19 = _dafny.SeqOf(_dafny.CodePoint('k'))
				var _1041_v20 _dafny.Array
				_ = _1041_v20
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw179
				_1041_v20 = _nw179
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))
				_ = _index151
				(_1041_v20).ArraySet1((_this).F13(), (_index151).Int())
				var _1042_v21 _dafny.Set
				_ = _1042_v21
				_1042_v21 = _dafny.SetOf((_this).F13(), (_this).F13(), _dafny.Zero)
				var _1043_v23 D7
				_ = _1043_v23
				_1043_v23 = Companion_D7_.Create_DC18_((_this).F13(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(787))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1044_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1045_i0 _dafny.Int) _dafny.CodePoint {
						return _1044_v1
					}
				})(_1019_v1))), _this.F12(), (_this).F14(), func() _dafny.Set {
					var _coll24 = _dafny.NewBuilder()
					_ = _coll24
					for _iter28 := _dafny.Iterate((_1042_v21).Elements()); ; {
						_compr_24, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _1046_v22 _dafny.Int
						_1046_v22 = interface{}(_compr_24).(_dafny.Int)
						if (_1042_v21).Contains(_1046_v22) {
							_coll24.Add((_1046_v22).Plus((_dafny.MultiSetOf(false)).Cardinality()))
						}
					}
					return _coll24.ToSet()
				}())
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))
				_ = _index152
				var _rhs161 _dafny.Int = (func() _dafny.Int {
					if p1 {
						return _dafny.IntOfInt64(43)
					}
					return (_this).F13()
				})()
				_ = _rhs161
				var _rhs162 _dafny.Sequence = (_1043_v23).Dtor_cf41()
				_ = _rhs162
				var _rhs163 _dafny.Int = ((_this).F13()).Times((_this).F13())
				_ = _rhs163
				var _rhs164 bool = !(_this.F12())
				_ = _rhs164
				var _lhs129 _dafny.Array = _1041_v20
				_ = _lhs129
				var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))
				_ = _lhs130
				var _lhs131 *GlobalState = globalState
				_ = _lhs131
				r0 = _rhs161
				_1040_v19 = _rhs162
				(_lhs129).ArraySet1(_rhs163, (_lhs130).Int())
				_lhs131.F7 = _rhs164
				var _1047_v24 _dafny.Map
				_ = _1047_v24
				_1047_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F14())
				var _1048_v25 _dafny.Array
				_ = _1048_v25
				var _nw180 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
				_ = _nw180
				_1048_v25 = _nw180
				var _1049_v26 _dafny.Map
				_ = _1049_v26
				_1049_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(p1)).IsProperSubsetOf(_this.F11()), _1048_v25)
				var _1050_v27 D4
				_ = _1050_v27
				_1050_v27 = Companion_D4_.Create_DC11_(false, (_1041_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))).Int()).(_dafny.Int), (_this).F13(), (_this).F13())
				var _1051_v28 _dafny.Map
				_ = _1051_v28
				_1051_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1041_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))).Int()).(_dafny.Int), _1041_v20)
				var _rhs165 _dafny.Map = _1047_v24
				_ = _rhs165
				var _rhs166 _dafny.Map = _1049_v26
				_ = _rhs166
				var _rhs167 D4 = Companion_D4_.Create_DC11_((Companion_D8_.Create_DC20_(p1, _this.F12(), p1)).Dtor_cf48(), (_1041_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))).Int()).(_dafny.Int), (_1041_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))).Int()).(_dafny.Int), ((func() _dafny.Map {
					if (_this).F14() {
						return _1051_v28
					}
					return _1051_v28
				})()).Cardinality())
				_ = _rhs167
				_1047_v24 = _rhs165
				_1049_v26 = _rhs166
				_1050_v27 = _rhs167
				var _1052_v29 _dafny.Array
				_ = _1052_v29
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
				_ = _nw181
				_1052_v29 = _nw181
				var _1053_v30 _dafny.Array
				_ = _1053_v30
				var _nw182 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(2))
				_ = _nw182
				_1053_v30 = _nw182
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1052_v29), 0))
				_ = _index153
				(_1052_v29).ArraySet1(_dafny.SeqOf(_1053_v30, _1053_v30, _1053_v30, _1053_v30, _1053_v30), (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1052_v29), 0))
				_ = _index154
				(_1052_v29).ArraySet1(_dafny.SeqOf(_1053_v30, _1053_v30), (_index154).Int())
				var _1054_v31 bool
				_ = _1054_v31
				_1054_v31 = _this.F12()
				var _1055_v32 _dafny.Map
				_ = _1055_v32
				_1055_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1054_v31), _1019_v1)
				(_this).M3(((_this).F13()).Minus((_1055_v32).Cardinality()), globalState)
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))
				_ = _index155
				var _rhs168 _dafny.Int = (_this).F13()
				_ = _rhs168
				var _rhs169 _dafny.Int = _dafny.IntOfInt64(386)
				_ = _rhs169
				var _rhs170 _dafny.MultiSet = _this.F11()
				_ = _rhs170
				var _rhs171 _dafny.Int = ((_1041_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))).Int()).(_dafny.Int)).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(918), (_1050_v27).Dtor_cf28()))
				_ = _rhs171
				var _lhs132 *GlobalState = globalState
				_ = _lhs132
				var _lhs133 *C3 = _this
				_ = _lhs133
				var _lhs134 _dafny.Array = _1041_v20
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1041_v20), 0))
				_ = _lhs135
				r0 = _rhs168
				_lhs132.F9 = _rhs169
				_lhs133.F11_set_(_rhs170)
				(_lhs134).ArraySet1(_rhs171, (_lhs135).Int())
			} else {
				r1 = false
				(globalState).F9 = (_this).F13()
				var _1056_v33 _dafny.Array
				_ = _1056_v33
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_23
				var _nw183 _dafny.Array
				_ = _nw183
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw183 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Set = (func(_1057_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
						return func(_1058_i1 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf(_dafny.SetOf(_dafny.IntOfUint32((_1057_p0).Cardinality())))).Union(_dafny.SetOf(_dafny.SetOf(_dafny.IntOfUint32((_1057_p0).Cardinality()), (_this).F13())))
						}
					})(p0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw183 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw183).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw183).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1056_v33 = _nw183
				var _1059_v34 _dafny.Map
				_ = _1059_v34
				_1059_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _this.F12())
				var _1060_v35 _dafny.Set
				_ = _1060_v35
				_1060_v35 = _dafny.SetOf((_this).F13(), (_1059_v34).Cardinality())
				var _1061_v36 _dafny.Sequence
				_ = _1061_v36
				_1061_v36 = _dafny.SeqOf((_this).F14())
				var _1062_v37 T0
				_ = _1062_v37
				var _nw184 *C1 = New_C1_()
				_ = _nw184
				_nw184.Ctor__((_this).F13(), _dafny.MultiSetOf(_this.F12(), true), true)
				_1062_v37 = _nw184
				var _1063_v38 _dafny.Sequence
				_ = _1063_v38
				_1063_v38 = _dafny.SeqOf(_1062_v37)
				var _1064_v39 _dafny.Map
				_ = _1064_v39
				_1064_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1063_v38).Select((Companion_Default___.SafeIndex((_this).F13(), _dafny.IntOfUint32((_1063_v38).Cardinality()))).Uint32()).(T0))
				var _1065_v40 D2
				_ = _1065_v40
				_1065_v40 = Companion_D2_.Create_DC5_(_dafny.IntOfUint32((_1061_v36).Cardinality()), (_1064_v39).Cardinality(), (_this).F13(), (_this).F13(), _dafny.SeqOf(_1062_v37.F12(), _1062_v37.F12()))
				var _1066_v41 _dafny.Set
				_ = _1066_v41
				_1066_v41 = _dafny.SetOf(_1060_v35, _1060_v35, _dafny.SetOf((_1065_v40).Dtor_cf12(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(706))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_1067_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1068_i2 _dafny.Int) _dafny.CodePoint {
						return _1067_v1
					}
				})(_1019_v1)))).Cardinality())))
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1056_v33), 0))
				_ = _index156
				(_1056_v33).ArraySet1(_1066_v41, (_index156).Int())
				var _1069_v42 _dafny.Array
				_ = _1069_v42
				var _nw185 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw185
				_1069_v42 = _nw185
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1056_v33), 0))
				_ = _index157
				var _rhs172 _dafny.Set = _1066_v41
				_ = _rhs172
				var _rhs173 _dafny.Array = _1069_v42
				_ = _rhs173
				var _rhs174 bool = (func() bool {
					if p1 {
						return p1
					}
					return _this.F12()
				})()
				_ = _rhs174
				var _lhs136 _dafny.Array = _1056_v33
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1056_v33), 0))
				_ = _lhs137
				var _lhs138 T0 = _1062_v37
				_ = _lhs138
				(_lhs136).ArraySet1(_rhs172, (_lhs137).Int())
				_1069_v42 = _rhs173
				_lhs138.F12_set_(_rhs174)
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1069_v42), 0))
				_ = _index158
				(_1069_v42).ArraySet1(_1062_v37.F12(), (_index158).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1069_v42), 0))
				_ = _index159
				(_1069_v42).ArraySet1(p1, (_index159).Int())
				var _1070_v43 _dafny.MultiSet
				_ = _1070_v43
				_1070_v43 = _dafny.MultiSetOf(_1065_v40, _1065_v40)
				var _1071_v44 _dafny.Sequence
				_ = _1071_v44
				_1071_v44 = _dafny.SeqOf(_dafny.MultiSetOf(_1065_v40, _1065_v40), (_1070_v43).Update(_1065_v40, Companion_Default___.Abs((_this).F13())), _1070_v43)
				_1071_v44 = _dafny.Companion_Sequence_.Concatenate(_1071_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-912))).Uint32(), func(coer56 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_1072_v43 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_1073_i3 _dafny.Int) _dafny.MultiSet {
						return _1072_v43
					}
				})(_1070_v43))))
			}
			var _1074_v45 _dafny.Array
			_ = _1074_v45
			var _nw186 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw186
			_1074_v45 = _nw186
			var _1075_v46 _dafny.Map
			_ = _1075_v46
			_1075_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1074_v45)
			var _1076_v47 _dafny.Sequence
			_ = _1076_v47
			_1076_v47 = _dafny.SeqOf((_this).F14())
			var _1077_v48 _dafny.Map
			_ = _1077_v48
			_1077_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_this).F14(), (_this).F13(), _dafny.IntOfUint32((_1076_v47).Cardinality()), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _1074_v45))
			var _1078_v49 D9
			_ = _1078_v49
			_1078_v49 = Companion_D9_.Create_DC23_(_1075_v46)
			var _1079_v50 _dafny.Array
			_ = _1079_v50
			var _nwElement0_42 _dafny.Map = _1075_v46
			_ = _nwElement0_42
			var _nw187 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(25))
			_ = _nw187
			(_nw187).ArraySet1(_nwElement0_42, 0)
			(_nw187).ArraySet1((func() _dafny.Map {
				if _this.F12() {
					return _1075_v46
				}
				return _1075_v46
			})(), 1)
			(_nw187).ArraySet1((_1075_v46).Update(true, _1074_v45), 2)
			(_nw187).ArraySet1(_1075_v46, 3)
			(_nw187).ArraySet1((_1075_v46).Merge(_1075_v46), 4)
			(_nw187).ArraySet1(_1075_v46, 5)
			(_nw187).ArraySet1(_1075_v46, 6)
			(_nw187).ArraySet1(_1075_v46, 7)
			(_nw187).ArraySet1(_1075_v46, 8)
			(_nw187).ArraySet1((func() _dafny.Map {
				if (_1077_v48).Contains((func() _dafny.Int {
					if (_this.F11()).Contains(true) {
						return (_this.F11()).Multiplicity(true)
					}
					return (_this).F13()
				})()) {
					return (_1077_v48).Get((func() _dafny.Int {
						if (_this.F11()).Contains(true) {
							return (_this.F11()).Multiplicity(true)
						}
						return (_this).F13()
					})()).(_dafny.Map)
				}
				return _1075_v46
			})(), 9)
			(_nw187).ArraySet1((_1078_v49).Dtor_cf55(), 10)
			(_nw187).ArraySet1((func() _dafny.Map {
				if _this.F12() {
					return _1075_v46
				}
				return _1075_v46
			})(), 11)
			(_nw187).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _1074_v45), 12)
			(_nw187).ArraySet1(_1075_v46, 13)
			(_nw187).ArraySet1(_1075_v46, 14)
			(_nw187).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _1074_v45)).Update(_this.F12(), _1074_v45), 15)
			(_nw187).ArraySet1(_1075_v46, 16)
			(_nw187).ArraySet1(_1075_v46, 17)
			(_nw187).ArraySet1(_1075_v46, 18)
			(_nw187).ArraySet1((_1075_v46).Merge(_1075_v46), 19)
			(_nw187).ArraySet1(_1075_v46, 20)
			(_nw187).ArraySet1(_1075_v46, 21)
			(_nw187).ArraySet1(_1075_v46, 22)
			(_nw187).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1074_v45), 23)
			(_nw187).ArraySet1((_1075_v46).Update(_this.F12(), _1074_v45), 24)
			_1079_v50 = _nw187
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1079_v50), 0))
			_ = _index160
			(_1079_v50).ArraySet1((func() _dafny.Map {
				if p1 {
					return _1075_v46
				}
				return _1075_v46
			})(), (_index160).Int())
			var _1080_v51 _dafny.MultiSet
			_ = _1080_v51
			_1080_v51 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this).F13()))
			var _1081_v52 _dafny.Sequence
			_ = _1081_v52
			_1081_v52 = _dafny.SeqOf(_1080_v51, _1080_v51)
			var _1082_v53 _dafny.Map
			_ = _1082_v53
			_1082_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F13())
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1079_v50), 0))
			_ = _index161
			var _rhs175 _dafny.Int = (_this).F13()
			_ = _rhs175
			var _rhs176 _dafny.Map = (_1075_v46).Merge((_1078_v49).Dtor_cf55())
			_ = _rhs176
			var _rhs177 bool = !((_this).Fm2((_1081_v52).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_1076_v47).Cardinality())), _dafny.IntOfUint32((_1081_v52).Cardinality()))).Uint32()).(_dafny.MultiSet), (_this).F13(), (_1082_v53).Update(false, (_this).F13()), (_this).F14(), globalState))
			_ = _rhs177
			var _rhs178 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_this).F13(), (_this).F13()), ((_this.F11()).Cardinality()).Plus((_this).F13()))
			_ = _rhs178
			var _lhs139 _dafny.Array = _1079_v50
			_ = _lhs139
			var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1079_v50), 0))
			_ = _lhs140
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			r0 = _rhs175
			(_lhs139).ArraySet1(_rhs176, (_lhs140).Int())
			_lhs141.F3 = _rhs177
			r0 = _rhs178
			var _1083_v54 _dafny.Array
			_ = _1083_v54
			var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw188
			_1083_v54 = _nw188
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_1083_v54), 0))
			_ = _index162
			(_1083_v54).ArraySet1((_this).F13(), (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_1083_v54), 0))
			_ = _index163
			(_1083_v54).ArraySet1((_this).F13(), (_index163).Int())
		} else {
			var _1084_v55 _dafny.CodePoint
			_ = _1084_v55
			_1084_v55 = _dafny.CodePoint('y')
			var _1085_v56 _dafny.MultiSet
			_ = _1085_v56
			_1085_v56 = _dafny.MultiSetOf(_1084_v55, _1084_v55)
			var _1086_v57 _dafny.Array
			_ = _1086_v57
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw189
			_1086_v57 = _nw189
			var _1087_v58 _dafny.Sequence
			_ = _1087_v58
			_1087_v58 = _dafny.SeqOf((_this).F13())
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))
			_ = _index164
			(_1086_v57).ArraySet1(_1087_v58, (_index164).Int())
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))
			_ = _index165
			var _rhs179 _dafny.MultiSet = _dafny.MultiSetOf(_1084_v55)
			_ = _rhs179
			var _rhs180 bool = (_this).F14()
			_ = _rhs180
			var _rhs181 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F13()), (_this).F13())
			_ = _rhs181
			var _rhs182 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1087_v58, _1087_v58), _1087_v58)
			_ = _rhs182
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			var _lhs143 _dafny.Array = _1086_v57
			_ = _lhs143
			var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))
			_ = _lhs144
			_1085_v56 = _rhs179
			r1 = _rhs180
			_lhs142.F9 = _rhs181
			(_lhs143).ArraySet1(_rhs182, (_lhs144).Int())
			var _1088_v59 _dafny.MultiSet
			_ = _1088_v59
			_1088_v59 = _dafny.MultiSetOf((_this).F13())
			var _1089_v60 D1
			_ = _1089_v60
			_1089_v60 = Companion_D1_.Create_DC2_(_this.F12(), _this.F12(), _this.F12(), _1084_v55, _1084_v55)
			var _1090_v61 D1
			_ = _1090_v61
			_1090_v61 = Companion_D1_.Create_DC2_(p1, false, !(p1), _1084_v55, Companion_Default___.Fm14(_1089_v60, (_this).F13(), globalState))
			var _1091_v62 _dafny.Array
			_ = _1091_v62
			var _nwElement0_43 _dafny.Sequence = p0
			_ = _nwElement0_43
			var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(23))
			_ = _nw190
			(_nw190).ArraySet1(_nwElement0_43, 0)
			(_nw190).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(833))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}((func(_1092_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1093_i4 _dafny.Int) _dafny.CodePoint {
					return _1092_v55
				}
			})(_1084_v55))), 1)
			(_nw190).ArraySet1(p0, 2)
			(_nw190).ArraySet1(p0, 3)
			(_nw190).ArraySet1(p0, 4)
			(_nw190).ArraySet1(p0, 5)
			(_nw190).ArraySet1(p0, 6)
			(_nw190).ArraySet1(p0, 7)
			(_nw190).ArraySet1(p0, 8)
			(_nw190).ArraySet1(p0, 9)
			(_nw190).ArraySet1(p0, 10)
			(_nw190).ArraySet1(p0, 11)
			(_nw190).ArraySet1(p0, 12)
			(_nw190).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ppym"), 13)
			(_nw190).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 14)
			(_nw190).ArraySet1(p0, 15)
			(_nw190).ArraySet1(p0, 16)
			(_nw190).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 17)
			(_nw190).ArraySet1((func() _dafny.Sequence {
				if (_this).F14() {
					return p0
				}
				return _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_1088_v59).Update((_this).F13(), Companion_Default___.Abs((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1088_v59).Contains((_this).F13()) {
						return (_1088_v59).Multiplicity((_this).F13())
					}
					return (_this).F13()
				})())))).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_1090_v61).Dtor_cf6())
			})(), 18)
			(_nw190).ArraySet1(p0, 19)
			(_nw190).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-542))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}((func(_1094_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1095_i5 _dafny.Int) _dafny.CodePoint {
					return _1094_v55
				}
			})(_1084_v55))), 20)
			(_nw190).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("drfnfyqmg"), 21)
			(_nw190).ArraySet1(p0, 22)
			_1091_v62 = _nw190
			var _1096_v63 _dafny.Map
			_ = _1096_v63
			_1096_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("whjni")).Cardinality()))
			var _1097_v64 _dafny.Sequence
			_ = _1097_v64
			_1097_v64 = _dafny.SeqOf(true)
			var _1098_v65 D2
			_ = _1098_v65
			_1098_v65 = Companion_D2_.Create_DC5_((_this).F13(), (_1085_v56).Cardinality(), (_this).F13(), (func() _dafny.Int {
				if (_1096_v63).Contains((_this).F14()) {
					return (_1096_v63).Get((_this).F14()).(_dafny.Int)
				}
				return (_this).F13()
			})(), _1097_v64)
			var _pat_let_tv21 = _1097_v64
			_ = _pat_let_tv21
			var _1099_v66 _dafny.Map
			_ = _1099_v66
			_1099_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D2 {
				if _this.F12() {
					return _1098_v65
				}
				return func(_pat_let32_0 D2) D2 {
					return func(_1100_dt__update__tmp_h1 D2) D2 {
						return func(_pat_let33_0 _dafny.Sequence) D2 {
							return func(_1101_dt__update_hcf13_h1 _dafny.Sequence) D2 {
								return func(_pat_let34_0 _dafny.Int) D2 {
									return func(_1102_dt__update_hcf9_h0 _dafny.Int) D2 {
										return Companion_D2_.Create_DC5_(_1102_dt__update_hcf9_h0, (_1100_dt__update__tmp_h1).Dtor_cf10(), (_1100_dt__update__tmp_h1).Dtor_cf11(), (_1100_dt__update__tmp_h1).Dtor_cf12(), _1101_dt__update_hcf13_h1)
									}(_pat_let34_0)
								}((_this).F13())
							}(_pat_let33_0)
						}(_pat_let_tv21)
					}(_pat_let32_0)
				}(_1098_v65)
			})(), _1091_v62)
			var _1103_v67 _dafny.Array
			_ = _1103_v67
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw191
			_1103_v67 = _nw191
			var _1104_v68 _dafny.Sequence
			_ = _1104_v68
			_1104_v68 = _dafny.SeqOf(_1103_v67, _1103_v67)
			var _1105_v70 _dafny.Set
			_ = _1105_v70
			_1105_v70 = _dafny.SetOf((_this).F13(), (_this).F13())
			var _1106_v71 _dafny.Sequence
			_ = _1106_v71
			_1106_v71 = _dafny.SeqOf(func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter29 := _dafny.Iterate((_1088_v59).Elements()); ; {
					_compr_25, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _1107_v69 _dafny.Int
					_1107_v69 = interface{}(_compr_25).(_dafny.Int)
					if (_1088_v59).Contains(_1107_v69) {
						_coll25.Add((_1107_v69).Times(_dafny.IntOfInt64(-388)))
					}
				}
				return _coll25.ToSet()
			}(), _1105_v70, _dafny.SetOf(_dafny.IntOfUint32((_1097_v64).Cardinality()), (_this).F13()))
			var _1108_v72 _dafny.Sequence
			_ = _1108_v72
			_1108_v72 = _dafny.SeqOf(_1103_v67, _1103_v67, (_1104_v68).Select((Companion_Default___.SafeIndex(((_1086_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(p1, (_this).F13(), _dafny.IntOfUint32((_1106_v71).Cardinality()), globalState), _dafny.IntOfUint32(((_1086_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1104_v68).Cardinality()))).Uint32()).(_dafny.Array))
			var _1109_v73 D9
			_ = _1109_v73
			_1109_v73 = Companion_D9_.Create_DC24_(_dafny.UnicodeSeqOfUtf8Bytes("c"), _this.F12(), (_this).F13())
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))
			_ = _index166
			var _rhs183 _dafny.Int = _dafny.IntOfInt64(-676)
			_ = _rhs183
			var _rhs184 _dafny.Array = (func() _dafny.Array {
				if (_1099_v66).Contains(func(_pat_let35_0 D2) D2 {
					return func(_1110_dt__update__tmp_h2 D2) D2 {
						return func(_pat_let36_0 _dafny.Int) D2 {
							return func(_1111_dt__update_hcf9_h1 _dafny.Int) D2 {
								return Companion_D2_.Create_DC5_(_1111_dt__update_hcf9_h1, (_1110_dt__update__tmp_h2).Dtor_cf10(), (_1110_dt__update__tmp_h2).Dtor_cf11(), (_1110_dt__update__tmp_h2).Dtor_cf12(), (_1110_dt__update__tmp_h2).Dtor_cf13())
							}(_pat_let36_0)
						}(_dafny.IntOfInt64(697))
					}(_pat_let35_0)
				}(_1098_v65)) {
					return (_1099_v66).Get(func(_pat_let37_0 D2) D2 {
						return func(_1112_dt__update__tmp_h3 D2) D2 {
							return func(_pat_let38_0 _dafny.Int) D2 {
								return func(_1113_dt__update_hcf9_h2 _dafny.Int) D2 {
									return Companion_D2_.Create_DC5_(_1113_dt__update_hcf9_h2, (_1112_dt__update__tmp_h3).Dtor_cf10(), (_1112_dt__update__tmp_h3).Dtor_cf11(), (_1112_dt__update__tmp_h3).Dtor_cf12(), (_1112_dt__update__tmp_h3).Dtor_cf13())
								}(_pat_let38_0)
							}(_dafny.IntOfInt64(697))
						}(_pat_let37_0)
					}(_1098_v65)).(_dafny.Array)
				}
				return p2
			})()
			_ = _rhs184
			var _rhs185 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1087_v58, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-730))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}(func(_1114_i6 _dafny.Int) _dafny.Int {
				return (_this).F13()
			})))
			_ = _rhs185
			var _rhs186 _dafny.Array = (_1104_v68).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F13(), (_this).F13()), _dafny.IntOfUint32((_1104_v68).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs186
			var _rhs187 bool = (func(_pat_let39_0 D9) D9 {
				return func(_1115_dt__update__tmp_h4 D9) D9 {
					return func(_pat_let40_0 _dafny.Int) D9 {
						return func(_1116_dt__update_hcf58_h0 _dafny.Int) D9 {
							return Companion_D9_.Create_DC24_((_1115_dt__update__tmp_h4).Dtor_cf56(), (_1115_dt__update__tmp_h4).Dtor_cf57(), _1116_dt__update_hcf58_h0)
						}(_pat_let40_0)
					}((_this).F13())
				}(_pat_let39_0)
			}(_1109_v73)).Dtor_cf57()
			_ = _rhs187
			var _lhs145 *GlobalState = globalState
			_ = _lhs145
			var _lhs146 _dafny.Array = _1086_v57
			_ = _lhs146
			var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))
			_ = _lhs147
			var _lhs148 *GlobalState = globalState
			_ = _lhs148
			var _lhs149 *GlobalState = globalState
			_ = _lhs149
			_lhs145.F9 = _rhs183
			_1091_v62 = _rhs184
			(_lhs146).ArraySet1(_rhs185, (_lhs147).Int())
			_lhs148.F4 = _rhs186
			_lhs149.F7 = _rhs187
			var _1117_v74 _dafny.Map
			_ = _1117_v74
			_1117_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _this.F12())
			_1117_v74 = (_1117_v74).Update(Companion_Default___.SafeDivisionInt((_this).F13(), (_this).F13()), !(p1) || (_this.F12()))
			r0 = (_dafny.Zero).Minus((_this).F13())
			var _1118_v75 *C0
			_ = _1118_v75
			var _nw192 *C0 = New_C0_()
			_ = _nw192
			_nw192.Ctor__((_this).F14(), (_this).F13(), (_this.F11()).Update(_this.F12(), Companion_Default___.Abs((_this).F13())), (_this).F14())
			_1118_v75 = _nw192
			var _1119_v76 D3
			_ = _1119_v76
			_1119_v76 = Companion_D3_.Create_DC9_((_this).F13(), Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_1088_v59).Contains(Companion_Default___.Fm1(_this.F12(), (_this).F13(), (_this).F13(), globalState)) {
					return (_1088_v59).Multiplicity(Companion_Default___.Fm1(_this.F12(), (_this).F13(), (_this).F13(), globalState))
				}
				return (_this).F13()
			})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_1086_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1086_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32(((_1086_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_1086_v57), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _dafny.IntOfInt64(-968))).Cardinality()))), _1118_v75, (_1118_v75).F18())
			_1119_v76 = _1119_v76
		}
		var _1120_v77 _dafny.Sequence
		_ = _1120_v77
		_1120_v77 = _dafny.SeqOf(!((_this).F14()), _this.F12(), p1)
		var _rhs188 bool = (_this).F14()
		_ = _rhs188
		var _rhs189 bool = (_1120_v77).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())).Cardinality(), _dafny.IntOfUint32((_1120_v77).Cardinality()))).Uint32()).(bool)
		_ = _rhs189
		var _lhs150 *GlobalState = globalState
		_ = _lhs150
		r1 = _rhs188
		_lhs150.F3 = _rhs189
		var _1121_v78 _dafny.Map
		_ = _1121_v78
		_1121_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F12()), (_this).F13())
		var _1122_v79 _dafny.MultiSet
		_ = _1122_v79
		_1122_v79 = _dafny.MultiSetOf(_1121_v78)
		var _1123_v80 D10
		_ = _1123_v80
		_1123_v80 = Companion_D10_.Create_DC27_(_1121_v78)
		var _1124_v81 _dafny.Sequence
		_ = _1124_v81
		_1124_v81 = _dafny.SeqOf(((_1121_v78).Update(p1, (_dafny.Zero).Minus((_this).F13()))).Cardinality(), (_this).F13(), (_this).F13())
		var _1125_v82 _dafny.Array
		_ = _1125_v82
		var _nwElement0_44 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _dafny.IntOfUint32((p0).Cardinality())), _1121_v78, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), (_this).F13()))
		_ = _nwElement0_44
		var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(13))
		_ = _nw193
		(_nw193).ArraySet1(_nwElement0_44, 0)
		(_nw193).ArraySet1(_dafny.MultiSetOf(_1121_v78, _1121_v78), 1)
		(_nw193).ArraySet1((_1122_v79).Intersection(_1122_v79), 2)
		(_nw193).ArraySet1(_1122_v79, 3)
		(_nw193).ArraySet1(_1122_v79, 4)
		(_nw193).ArraySet1(_dafny.MultiSetOf((_1123_v80).Dtor_cf63()), 5)
		(_nw193).ArraySet1(_1122_v79, 6)
		(_nw193).ArraySet1(_1122_v79, 7)
		(_nw193).ArraySet1(_1122_v79, 8)
		(_nw193).ArraySet1(_1122_v79, 9)
		(_nw193).ArraySet1(_1122_v79, 10)
		(_nw193).ArraySet1((_1122_v79).Update(_1121_v78, Companion_Default___.Abs(_dafny.IntOfUint32((_1124_v81).Cardinality()))), 11)
		(_nw193).ArraySet1((_1122_v79).Difference(Companion_Default___.Fm19((_this).F13(), globalState)), 12)
		_1125_v82 = _nw193
		var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_1125_v82), 0))
		_ = _index167
		(_1125_v82).ArraySet1(_1122_v79, (_index167).Int())
		var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_1125_v82), 0))
		_ = _index168
		(_1125_v82).ArraySet1((_1122_v79).Difference((_1122_v79).Union(_1122_v79)), (_index168).Int())
		var _1126_v83 _dafny.Map
		_ = _1126_v83
		_1126_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F13()), (_this).F13())
		if (_1126_v83).Contains(_1124_v81) {
			var _1127_v84 _dafny.Map
			_ = _1127_v84
			_1127_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1120_v77, p1)
			_1127_v84 = (_1127_v84).Update(_dafny.Companion_Sequence_.Concatenate(_1120_v77, _dafny.Companion_Sequence_.Update(_1120_v77, (Companion_Default___.SafeIndex((_this.F11()).Cardinality(), _dafny.IntOfUint32((_1120_v77).Cardinality()))).Uint32(), !(p1))), _this.F12())
			var _1128_v85 *C2
			_ = _1128_v85
			var _nw194 *C2 = New_C2_()
			_ = _nw194
			_nw194.Ctor__((_this).F14(), _this.F11(), (_this).F14())
			_1128_v85 = _nw194
			_1128_v85 = _1128_v85
			var _1129_v86 _dafny.Sequence
			_ = _1129_v86
			_1129_v86 = _dafny.UnicodeSeqOfUtf8Bytes("kfcc")
			_1129_v86 = _dafny.Companion_Sequence_.Concatenate(_1129_v86, p0)
			var _1130_v87 D5
			_ = _1130_v87
			_1130_v87 = Companion_D5_.Create_DC14_()
			var _1131_v88 D7
			_ = _1131_v88
			_1131_v88 = Companion_D7_.Create_DC17_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1130_v87, p1))
			var _1132_v89 _dafny.Map
			_ = _1132_v89
			_1132_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1130_v87, _this.F12())
			var _pat_let_tv22 = _1132_v89
			_ = _pat_let_tv22
			var _pat_let_tv23 = _1132_v89
			_ = _pat_let_tv23
			_1131_v88 = func(_pat_let41_0 D7) D7 {
				return func(_1133_dt__update__tmp_h5 D7) D7 {
					return func(_pat_let42_0 _dafny.Map) D7 {
						return func(_1134_dt__update_hcf39_h0 _dafny.Map) D7 {
							return Companion_D7_.Create_DC17_(_1134_dt__update_hcf39_h0)
						}(_pat_let42_0)
					}((_pat_let_tv22).Merge(_pat_let_tv23))
				}(_pat_let41_0)
			}(_1131_v88)
			(globalState).F7 = (func() bool {
				if _1128_v85.F15 {
					return _1128_v85.F15
				}
				return _this.F12()
			})()
		} else {
			(globalState).F8 = (_this).F14()
			var _1135_v90 _dafny.CodePoint
			_ = _1135_v90
			_1135_v90 = _dafny.CodePoint('s')
			var _1136_v91 D1
			_ = _1136_v91
			_1136_v91 = Companion_D1_.Create_DC2_(_this.F12(), false, (_this).F14(), _1135_v90, _1135_v90)
			var _1137_v92 _dafny.Map
			_ = _1137_v92
			_1137_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1124_v81).Cardinality()), Companion_Default___.Fm14(_1136_v91, (_this).F13(), globalState))
			var _1138_v93 _dafny.Set
			_ = _1138_v93
			_1138_v93 = _dafny.SetOf((_this).F14())
			var _1139_v94 _dafny.MultiSet
			_ = _1139_v94
			_1139_v94 = _dafny.MultiSetOf((_this).F13())
			var _1140_v95 _dafny.Map
			_ = _1140_v95
			_1140_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F13()), (_this).F13())
			var _1141_v96 _dafny.Set
			_ = _1141_v96
			_1141_v96 = _dafny.SetOf((_this).F13(), (_1140_v95).Cardinality(), (_this).F13())
			var _1142_v97 _dafny.Map
			_ = _1142_v97
			_1142_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(_1139_v94, (_this).F13(), _1121_v78, p1, globalState), _1141_v96)
			var _1143_v98 _dafny.Array
			_ = _1143_v98
			var _nwElement0_45 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1124_v81).Cardinality()))
			_ = _nwElement0_45
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(24))
			_ = _nw195
			(_nw195).ArraySet1(_nwElement0_45, 0)
			(_nw195).ArraySet1(Companion_Default___.Fm1((_this).F14(), (_this).F13(), _dafny.IntOfInt64(375), globalState), 1)
			(_nw195).ArraySet1((_1137_v92).Cardinality(), 2)
			(_nw195).ArraySet1((_this).F13(), 3)
			(_nw195).ArraySet1((_this).F13(), 4)
			(_nw195).ArraySet1((_this).F13(), 5)
			(_nw195).ArraySet1((_this).F13(), 6)
			(_nw195).ArraySet1((_1138_v93).Cardinality(), 7)
			(_nw195).ArraySet1((_this).F13(), 8)
			(_nw195).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mwn")).Cardinality()), 9)
			(_nw195).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 10)
			(_nw195).ArraySet1((_this).F13(), 11)
			(_nw195).ArraySet1((_this).F13(), 12)
			(_nw195).ArraySet1((_this).F13(), 13)
			(_nw195).ArraySet1((_this).F13(), 14)
			(_nw195).ArraySet1((_1142_v97).Cardinality(), 15)
			(_nw195).ArraySet1((_this).F13(), 16)
			(_nw195).ArraySet1((_this).F13(), 17)
			(_nw195).ArraySet1(Companion_Default___.Fm1(p1, (_1121_v78).Cardinality(), (_this).F13(), globalState), 18)
			(_nw195).ArraySet1((_this).F13(), 19)
			(_nw195).ArraySet1((_dafny.Zero).Minus((_this).F13()), 20)
			(_nw195).ArraySet1((_this).F13(), 21)
			(_nw195).ArraySet1((_this).F13(), 22)
			(_nw195).ArraySet1((_this).F13(), 23)
			_1143_v98 = _nw195
			var _1144_v99 _dafny.Set
			_ = _1144_v99
			_1144_v99 = _dafny.SetOf(_1143_v98)
			_1144_v99 = (func() _dafny.Set {
				if (_this).F14() {
					return (_1144_v99).Difference(_1144_v99)
				}
				return (_1144_v99).Union(_1144_v99)
			})()
			var _1145_v100 _dafny.Map
			_ = _1145_v100
			_1145_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _1143_v98)
			_1145_v100 = (_1145_v100).Update(_this.F12(), _1143_v98)
			(_this).F12_set_(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(558))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}(func(_1146_i7 _dafny.Int) _dafny.Int {
				return (_this).F13()
			})), _1124_v81), _1124_v81))
			var _1147_v101 _dafny.Sequence
			_ = _1147_v101
			_1147_v101 = _dafny.UnicodeSeqOfUtf8Bytes("m")
			_1147_v101 = p0
		}
		var _1148_v102 *C0
		_ = _1148_v102
		var _nw196 *C0 = New_C0_()
		_ = _nw196
		_nw196.Ctor__(!((_this).F14()), (_this).F13(), _this.F11(), p1)
		_1148_v102 = _nw196
		(_this).F12_set_(_this.F12())
		r0 = (_dafny.Zero).Minus(Companion_Default___.Fm1(_this.F12(), (_this).F13(), Companion_Default___.SafeModuloInt((_this).F13(), _dafny.IntOfInt64(120)), globalState))
		var _1149_v103 D5
		_ = _1149_v103
		_1149_v103 = Companion_D5_.Create_DC12_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(179))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}(func(_1150_i8 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(382)
		})))
		var _1151_v104 _dafny.MultiSet
		_ = _1151_v104
		_1151_v104 = _dafny.MultiSetOf(_1149_v103, _1149_v103, _1149_v103)
		r1 = (_1151_v104).IsSubsetOf(_1151_v104)
		return r0, r1
	}
}
func (_this *C3) M3(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _1152_v0 _dafny.Array
		_ = _1152_v0
		var _nw197 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw197
		_1152_v0 = _nw197
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
		_ = _index169
		(_1152_v0).ArraySet1(p0, (_index169).Int())
		var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
		_ = _index170
		(_1152_v0).ArraySet1(p0, (_index170).Int())
		var _1153_v1 _dafny.Array
		_ = _1153_v1
		var _nw198 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(13))
		_ = _nw198
		_1153_v1 = _nw198
		var _1154_v2 D4
		_ = _1154_v2
		_1154_v2 = Companion_D4_.Create_DC10_(_1152_v0)
		var _pat_let_tv24 = _1152_v0
		_ = _pat_let_tv24
		var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1153_v1), 0))
		_ = _index171
		(_1153_v1).ArraySet1(func(_pat_let43_0 D4) D4 {
			return func(_1155_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let44_0 _dafny.Array) D4 {
					return func(_1156_dt__update_hcf24_h0 _dafny.Array) D4 {
						return Companion_D4_.Create_DC10_(_1156_dt__update_hcf24_h0)
					}(_pat_let44_0)
				}(_pat_let_tv24)
			}(_pat_let43_0)
		}(_1154_v2), (_index171).Int())
		var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1153_v1), 0))
		_ = _index172
		(_1153_v1).ArraySet1(Companion_D4_.Create_DC10_(_1152_v0), (_index172).Int())
		if (_this).F14() {
			var _1157_v3 _dafny.CodePoint
			_ = _1157_v3
			_1157_v3 = _dafny.CodePoint('m')
			var _1158_v4 D1
			_ = _1158_v4
			_1158_v4 = Companion_D1_.Create_DC2_((_this).F14(), (_this).F14(), _this.F12(), _1157_v3, _1157_v3)
			var _1159_v5 _dafny.Map
			_ = _1159_v5
			_1159_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm14(_1158_v4, p0, globalState))
			var _1160_v6 _dafny.Set
			_ = _1160_v6
			_1160_v6 = _dafny.SetOf(p0, (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
			var _rhs190 bool = !((_1160_v6).IsSubsetOf(Companion_Default___.Fm20(_1157_v3, _dafny.IntOfInt64(678), _dafny.IntOfInt64(388), _1159_v5, globalState)))
			_ = _rhs190
			var _rhs191 _dafny.Int = _dafny.IntOfInt64(-307)
			_ = _rhs191
			var _lhs151 *C3 = _this
			_ = _lhs151
			var _lhs152 *GlobalState = globalState
			_ = _lhs152
			_lhs151.F12_set_(_rhs190)
			_lhs152.F9 = _rhs191
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
			_ = _index173
			(_1152_v0).ArraySet1((_this).F13(), (_index173).Int())
			var _1161_v7 _dafny.Map
			_ = _1161_v7
			_1161_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1152_v0, (_1160_v6).Cardinality())
			var _1162_v8 _dafny.Sequence
			_ = _1162_v8
			_1162_v8 = _dafny.SeqOf(_1152_v0, _1152_v0)
			var _1163_v9 _dafny.Set
			_ = _1163_v9
			_1163_v9 = _dafny.SetOf(_this.F12())
			var _1164_v10 _dafny.Sequence
			_ = _1164_v10
			_1164_v10 = _dafny.SeqOf((_1163_v9).Cardinality(), p0, (_this).F13(), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
			_1161_v7 = (_1161_v7).Update((_1162_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1164_v10).Cardinality()), _dafny.IntOfUint32((_1162_v8).Cardinality()))).Uint32()).(_dafny.Array), p0)
			var _1165_v11 _dafny.Sequence
			_ = _1165_v11
			_1165_v11 = _dafny.UnicodeSeqOfUtf8Bytes("hoftygcb")
			_1165_v11 = _1165_v11
			(globalState).F9 = (p0).Times((_this).F13())
		} else {
			(globalState).F9 = _dafny.IntOfInt64(701)
			var _1166_v12 _dafny.Array
			_ = _1166_v12
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_24
			var _nw199 _dafny.Array
			_ = _nw199
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw199 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Sequence = func(_1167_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("ipdww")
				}
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw199 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw199).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw199).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1166_v12 = _nw199
			var _1168_v13 D5
			_ = _1168_v13
			_1168_v13 = Companion_D5_.Create_DC13_((_this).F13(), _this.F12(), (p0).Minus(p0), _1166_v12, (_this).F14())
			var _source10 D5 = _1168_v13
			_ = _source10
			if _source10.Is_DC13() {
				var _1169___mcc_h0 _dafny.Int = _source10.Get_().(D5_DC13).Cf30
				_ = _1169___mcc_h0
				var _1170___mcc_h1 bool = _source10.Get_().(D5_DC13).Cf31
				_ = _1170___mcc_h1
				var _1171___mcc_h2 _dafny.Int = _source10.Get_().(D5_DC13).Cf32
				_ = _1171___mcc_h2
				var _1172___mcc_h3 _dafny.Array = _source10.Get_().(D5_DC13).Cf33
				_ = _1172___mcc_h3
				var _1173___mcc_h4 bool = _source10.Get_().(D5_DC13).Cf34
				_ = _1173___mcc_h4
				var _1174_cf34 bool = _1173___mcc_h4
				_ = _1174_cf34
				var _1175_cf33 _dafny.Array = _1172___mcc_h3
				_ = _1175_cf33
				var _1176_cf32 _dafny.Int = _1171___mcc_h2
				_ = _1176_cf32
				var _1177_cf31 bool = _1170___mcc_h1
				_ = _1177_cf31
				var _1178_cf30 _dafny.Int = _1169___mcc_h0
				_ = _1178_cf30
				var _1179_v14 *C2
				_ = _1179_v14
				var _nw200 *C2 = New_C2_()
				_ = _nw200
				_nw200.Ctor__(_this.F12(), (_this.F11()).Difference(_dafny.MultiSetOf((_this).F14())), _1174_cf34)
				_1179_v14 = _nw200
				_1179_v14 = _1179_v14
				_1176_cf32 = Companion_Default___.SafeModuloInt((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("klkgrol")).Cardinality())).Times(_1176_cf32))
				var _1180_v15 _dafny.Array
				_ = _1180_v15
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_25
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) bool = func(_1181_i1 _dafny.Int) bool {
						return _this.F12()
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw201 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw201).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw201).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1180_v15 = _nw201
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1180_v15), 0))
				_ = _index174
				(_1180_v15).ArraySet1(_this.F12(), (_index174).Int())
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_1180_v15), 0))
				_ = _index175
				(_1180_v15).ArraySet1((_this).F14(), (_index175).Int())
				var _1182_v16 _dafny.Map
				_ = _1182_v16
				_1182_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1179_v14.F15, _1180_v15)
				(globalState).F3 = ((_1182_v16).Merge(_1182_v16)).Equals(_1182_v16)
			} else if _source10.Is_DC14() {
				(globalState).F9 = Companion_Default___.Fm1((_this).F14(), _dafny.IntOfInt64(12), (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(86), _dafny.IntOfInt64(617), (_this.F11()).Cardinality())).Cardinality())).Minus((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)), globalState)
				var _1183_v17 _dafny.Sequence
				_ = _1183_v17
				_1183_v17 = _dafny.SeqOf(_this.F12())
				var _1184_v18 _dafny.Map
				_ = _1184_v18
				_1184_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), p0)
				var _1185_v19 _dafny.Sequence
				_ = _1185_v19
				_1185_v19 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				var _1186_v20 _dafny.Set
				_ = _1186_v20
				_1186_v20 = _dafny.SetOf(_dafny.IntOfUint32((_1185_v19).Cardinality()))
				var _1187_v21 *C0
				_ = _1187_v21
				var _nw202 *C0 = New_C0_()
				_ = _nw202
				_nw202.Ctor__(_this.F12(), p0, (_this.F11()).Difference((_dafny.MultiSetFromSeq(_1183_v17)).Update((_this).F14(), Companion_Default___.Abs(_dafny.IntOfUint32((_1183_v17).Cardinality())))), (_dafny.SetOf((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1184_v18).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((false), (_this).F14())) {
						return (_1184_v18).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((false), (_this).F14())).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-628)
				})()), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))).IsSubsetOf(_1186_v20))
				_1187_v21 = _nw202
				var _nw203 *C0 = New_C0_()
				_ = _nw203
				_nw203.Ctor__((_this).F14(), (_this).F13(), _dafny.MultiSetOf((_1187_v21).F17()), (_this).F14())
				_1187_v21 = _nw203
				var _1188_v22 D9
				_ = _1188_v22
				_1188_v22 = Companion_D9_.Create_DC25_(_1185_v19, (_1187_v21).F17(), p0)
				var _1189_v23 _dafny.CodePoint
				_ = _1189_v23
				_1189_v23 = _dafny.CodePoint('y')
				_1185_v19 = _dafny.Companion_Sequence_.Update((_1188_v22).Dtor_cf59(), (Companion_Default___.SafeIndex((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_1188_v22).Dtor_cf59()).Cardinality()))).Uint32(), _1189_v23)
				var _1190_v24 D2
				_ = _1190_v24
				_1190_v24 = Companion_D2_.Create_DC5_(p0, (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(223), p0, _1183_v17)
				var _rhs192 bool = _this.F12()
				_ = _rhs192
				var _rhs193 _dafny.CodePoint = _1189_v23
				_ = _rhs193
				var _rhs194 bool = (Companion_Default___.SafeModuloInt((_1187_v21).F18(), (_1187_v21).Fm12(globalState))).Cmp((_dafny.IntOfUint32(((_1190_v24).Dtor_cf13()).Cardinality())).Plus((_this).F13())) < 0
				_ = _rhs194
				var _rhs195 _dafny.Int = Companion_Default___.SafeModuloInt((_1186_v20).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1189_v23, _this.F12())).Cardinality())
				_ = _rhs195
				var _lhs153 *GlobalState = globalState
				_ = _lhs153
				var _lhs154 *GlobalState = globalState
				_ = _lhs154
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				_lhs153.F8 = _rhs192
				_1189_v23 = _rhs193
				_lhs154.F7 = _rhs194
				_lhs155.F9 = _rhs195
			} else {
				var _1191___mcc_h5 _dafny.Sequence = _source10.Get_().(D5_DC12).Cf29
				_ = _1191___mcc_h5
				var _1192_cf29 _dafny.Sequence = _1191___mcc_h5
				_ = _1192_cf29
				var _1193_v25 T0
				_ = _1193_v25
				var _nw204 *C2 = New_C2_()
				_ = _nw204
				_nw204.Ctor__((_this).F14(), _this.F11(), _this.F12())
				_1193_v25 = _nw204
				_1193_v25 = _1193_v25
				var _1194_v26 _dafny.CodePoint
				_ = _1194_v26
				_1194_v26 = _dafny.CodePoint('c')
				var _1195_v27 _dafny.Sequence
				_ = _1195_v27
				_1195_v27 = _dafny.SeqOf((_this).F14())
				var _1196_v28 _dafny.Map
				_ = _1196_v28
				_1196_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1195_v27).Cardinality()), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
				var _1197_v29 *C0
				_ = _1197_v29
				var _nw205 *C0 = New_C0_()
				_ = _nw205
				_nw205.Ctor__(Companion_Default___.Fm0(_1194_v26, globalState), Companion_Default___.Fm1(_1193_v25.F12(), (_this).F13(), (_this).F13(), globalState), _dafny.MultiSetOf(!((_this).F14()), (_this).F14(), _this.F12(), _this.F12(), !((_this).F14())), !(((func() _dafny.Int {
					if (_1196_v28).Contains((_this).F13()) {
						return (_1196_v28).Get((_this).F13()).(_dafny.Int)
					}
					return (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)
				})()).Cmp(_dafny.IntOfInt64(-424)) == 0))
				_1197_v29 = _nw205
				var _1198_v30 _dafny.Set
				_ = _1198_v30
				_1198_v30 = _dafny.SetOf(_this.F12(), (_this).F14())
				var _1199_v31 *C0
				_ = _1199_v31
				var _nw206 *C0 = New_C0_()
				_ = _nw206
				_nw206.Ctor__((_dafny.SetOf(_this.F12(), (_this).F14())).IsProperSubsetOf(_1198_v30), _dafny.IntOfInt64(502), _dafny.MultiSetOf((_this).F14(), (_1197_v29).F17()), !((_1197_v29).F17()))
				_1199_v31 = _nw206
				var _1200_v32 _dafny.Array
				_ = _1200_v32
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_26
				var _nw207 _dafny.Array
				_ = _nw207
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw207 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = (func(_1201_v31 *C0, _1202_v25 T0) func(_dafny.Int) bool {
						return func(_1203_i2 _dafny.Int) bool {
							return !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v25.F12(), (_1201_v31).F17())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1201_v31).F17(), (_1201_v31).F17()))
						}
					})(_1199_v31, _1193_v25)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw207 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw207).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw207).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1200_v32 = _nw207
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1200_v32), 0))
				_ = _index176
				(_1200_v32).ArraySet1((_1199_v31).F17(), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_1200_v32), 0))
				_ = _index177
				(_1200_v32).ArraySet1(!(Companion_Default___.Fm0(_1194_v26, globalState)) || (!((_1197_v29).F17())), (_index177).Int())
			}
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
			_ = _index178
			(_1152_v0).ArraySet1((Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F13()), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))).Plus((_this).F13()), (_index178).Int())
			var _1204_v33 _dafny.Sequence
			_ = _1204_v33
			_1204_v33 = _dafny.SeqOf(_1152_v0)
			var _1205_v34 _dafny.CodePoint
			_ = _1205_v34
			_1205_v34 = _dafny.CodePoint('l')
			var _1206_v35 D1
			_ = _1206_v35
			_1206_v35 = Companion_D1_.Create_DC2_((_this).F14(), !((_this).F14()), _this.F12(), _dafny.CodePoint('c'), _1205_v34)
			var _source11 D5 = (func() D5 {
				if _this.F12() {
					return Companion_D5_.Create_DC13_((_this.F11()).Cardinality(), _this.F12(), _dafny.IntOfUint32((_1204_v33).Cardinality()), _1166_v12, (_1206_v35).Dtor_cf2())
				}
				return _1168_v13
			})()
			_ = _source11
			if _source11.Is_DC13() {
				var _1207___mcc_h6 _dafny.Int = _source11.Get_().(D5_DC13).Cf30
				_ = _1207___mcc_h6
				var _1208___mcc_h7 bool = _source11.Get_().(D5_DC13).Cf31
				_ = _1208___mcc_h7
				var _1209___mcc_h8 _dafny.Int = _source11.Get_().(D5_DC13).Cf32
				_ = _1209___mcc_h8
				var _1210___mcc_h9 _dafny.Array = _source11.Get_().(D5_DC13).Cf33
				_ = _1210___mcc_h9
				var _1211___mcc_h10 bool = _source11.Get_().(D5_DC13).Cf34
				_ = _1211___mcc_h10
				var _1212_cf34 bool = _1211___mcc_h10
				_ = _1212_cf34
				var _1213_cf33 _dafny.Array = _1210___mcc_h9
				_ = _1213_cf33
				var _1214_cf32 _dafny.Int = _1209___mcc_h8
				_ = _1214_cf32
				var _1215_cf31 bool = _1208___mcc_h7
				_ = _1215_cf31
				var _1216_cf30 _dafny.Int = _1207___mcc_h6
				_ = _1216_cf30
				(globalState).F3 = _1212_cf34
				var _1217_v36 _dafny.Array
				_ = _1217_v36
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_27
				var _nw208 _dafny.Array
				_ = _nw208
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw208 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = (func(_1218_cf34 bool) func(_dafny.Int) bool {
						return func(_1219_i3 _dafny.Int) bool {
							return _1218_cf34
						}
					})(_1212_cf34)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw208 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw208).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw208).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1217_v36 = _nw208
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1217_v36), 0))
				_ = _index179
				(_1217_v36).ArraySet1(!(_1212_cf34), (_index179).Int())
				var _1220_v38 _dafny.Set
				_ = _1220_v38
				_1220_v38 = _dafny.SetOf(_dafny.CodePoint('c'), _1205_v34, _1205_v34)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_1217_v36), 0))
				_ = _index180
				(_1217_v36).ArraySet1((func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter30 := _dafny.Iterate((Companion_Default___.Fm21(_1212_cf34, true, _1215_cf31, globalState)).Keys().Elements()); ; {
						_compr_26, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _1221_v37 _dafny.CodePoint
						_1221_v37 = interface{}(_compr_26).(_dafny.CodePoint)
						if (Companion_Default___.Fm21(_1212_cf34, true, _1215_cf31, globalState)).Contains(_1221_v37) {
							_coll26.Add(_1221_v37)
						}
					}
					return _coll26.ToSet()
				}()).IsSubsetOf(_1220_v38), (_index180).Int())
				var _1222_v39 _dafny.Sequence
				_ = _1222_v39
				_1222_v39 = _dafny.SeqOf(_1217_v36, _1217_v36, _1217_v36, _1217_v36, _1217_v36)
				_1222_v39 = _1222_v39
				var _1223_v40 _dafny.Map
				_ = _1223_v40
				_1223_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1152_v0, (_this).F13())
				_1205_v34 = (func() _dafny.CodePoint {
					if (_1223_v40).Equals(_1223_v40) {
						return _1205_v34
					}
					return _1205_v34
				})()
			} else if _source11.Is_DC14() {
				var _1224_v41 _dafny.Sequence
				_ = _1224_v41
				_1224_v41 = _dafny.SeqOf(_this.F12())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _index181
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1153_v1), 0))
				_ = _index182
				var _rhs196 _dafny.Int = (_this).F13()
				_ = _rhs196
				var _rhs197 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1224_v41).Cardinality()), Companion_Default___.Fm1(_this.F12(), p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("scarno")).Cardinality()), globalState))
				_ = _rhs197
				var _rhs198 D4 = Companion_D4_.Create_DC10_(_1152_v0)
				_ = _rhs198
				var _lhs156 _dafny.Array = _1152_v0
				_ = _lhs156
				var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				var _lhs159 _dafny.Array = _1153_v1
				_ = _lhs159
				var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1153_v1), 0))
				_ = _lhs160
				(_lhs156).ArraySet1(_rhs196, (_lhs157).Int())
				_lhs158.F9 = _rhs197
				(_lhs159).ArraySet1(_rhs198, (_lhs160).Int())
				var _1225_v42 _dafny.Sequence
				_ = _1225_v42
				_1225_v42 = _dafny.SeqOf(_1224_v41, _1224_v41, _dafny.Companion_Sequence_.Update(_1224_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.IntOfUint32((_1224_v41).Cardinality()))).Uint32(), _this.F12()))
				var _1226_v43 _dafny.Map
				_ = _1226_v43
				_1226_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1225_v42)
				_1226_v43 = (_1226_v43).Update((_this).F14(), _dafny.Companion_Sequence_.Concatenate(_1225_v42, _1225_v42))
				var _1227_v44 _dafny.Array
				_ = _1227_v44
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw209
				_1227_v44 = _nw209
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1227_v44), 0))
				_ = _index183
				(_1227_v44).ArraySet1((_this).F14(), (_index183).Int())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1227_v44), 0))
				_ = _index184
				var _rhs199 _dafny.Int = p0
				_ = _rhs199
				var _rhs200 bool = _this.F12()
				_ = _rhs200
				var _rhs201 bool = _this.F12()
				_ = _rhs201
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				var _lhs162 _dafny.Array = _1227_v44
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1227_v44), 0))
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				_lhs161.F9 = _rhs199
				(_lhs162).ArraySet1(_rhs200, (_lhs163).Int())
				_lhs164.F8 = _rhs201
				var _1228_v45 _dafny.Array
				_ = _1228_v45
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw210
				_1228_v45 = _nw210
				_1228_v45 = _1228_v45
			} else {
				var _1229___mcc_h11 _dafny.Sequence = _source11.Get_().(D5_DC12).Cf29
				_ = _1229___mcc_h11
				var _1230_cf29 _dafny.Sequence = _1229___mcc_h11
				_ = _1230_cf29
				var _1231_v46 _dafny.Sequence
				_ = _1231_v46
				_1231_v46 = _dafny.SeqOf(_this.F12())
				var _1232_v47 T0
				_ = _1232_v47
				var _nw211 *C0 = New_C0_()
				_ = _nw211
				_nw211.Ctor__(true, (_this).F13(), _dafny.MultiSetFromSeq(_1231_v46), false)
				_1232_v47 = _nw211
				var _1233_v48 _dafny.Sequence
				_ = _1233_v48
				_1233_v48 = _dafny.SeqOf(_1232_v47)
				var _1234_v49 _dafny.Map
				_ = _1234_v49
				_1234_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.Companion_Sequence_.Concatenate(_1233_v48, _1233_v48))
				_1234_v49 = (_1234_v49).Update(p0, _dafny.Companion_Sequence_.Concatenate(_1233_v48, _1233_v48))
				var _1235_v50 *C2
				_ = _1235_v50
				var _nw212 *C2 = New_C2_()
				_ = _nw212
				_nw212.Ctor__(_1232_v47.F12(), _this.F11(), (_this).F14())
				_1235_v50 = _nw212
				var _1236_v51 _dafny.Array
				_ = _1236_v51
				var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
				_ = _nw213
				_1236_v51 = _nw213
				var _1237_v52 _dafny.Sequence
				_ = _1237_v52
				_1237_v52 = _dafny.UnicodeSeqOfUtf8Bytes("tdx")
				var _1238_v53 _dafny.Set
				_ = _1238_v53
				_1238_v53 = _dafny.SetOf(_dafny.IntOfUint32((_1237_v52).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(583))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1239_v34 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1240_i4 _dafny.Int) _dafny.CodePoint {
						return _1239_v34
					}
				})(_1205_v34)))).Cardinality()))
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1236_v51), 0))
				_ = _index185
				(_1236_v51).ArraySet1(_1238_v53, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_1236_v51), 0))
				_ = _index186
				(_1236_v51).ArraySet1(func() _dafny.Set {
					var _coll27 = _dafny.NewBuilder()
					_ = _coll27
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(188), _dafny.IntOfInt64(25))); ; {
						_compr_27, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _1241_v54 _dafny.Int
						_1241_v54 = interface{}(_compr_27).(_dafny.Int)
						if ((_dafny.IntOfInt64(188)).Cmp(_1241_v54) <= 0) && ((_1241_v54).Cmp(_dafny.IntOfInt64(25)) < 0) {
							_coll27.Add((_1241_v54).Minus((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)))
						}
					}
					return _coll27.ToSet()
				}(), (_index186).Int())
				(globalState).F8 = (func() bool {
					if Companion_Default___.Fm0(_1205_v34, globalState) {
						return _1235_v50.F15
					}
					return false
				})()
			}
			(globalState).F7 = (_this).Fm6(!(_this.F12()), (_this).F13(), globalState)
		}
		var _1242_i5 _dafny.Int
		_ = _1242_i5
		_1242_i5 = _dafny.Zero
		{
			for _this.F12() {
				{
					if (_1242_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1242_i5 = (_1242_i5).Plus(_dafny.One)
					Companion_Default___.M0(globalState)
					var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
					_ = _index187
					(_1152_v0).ArraySet1((_dafny.IntOfInt64(-310)).Times(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F13()), p0)), (_index187).Int())
					var _1243_v55 _dafny.CodePoint
					_ = _1243_v55
					_1243_v55 = _dafny.CodePoint('p')
					var _1244_v56 _dafny.Map
					_ = _1244_v56
					_1244_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), _1243_v55)
					if Companion_Default___.Fm0((func() _dafny.CodePoint {
						if (_1244_v56).Contains(p0) {
							return (_1244_v56).Get(p0).(_dafny.CodePoint)
						}
						return _1243_v55
					})(), globalState) {
						var _1245_v57 _dafny.Array
						_ = _1245_v57
						var _len0_28 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_28
						var _nw214 _dafny.Array
						_ = _nw214
						if _len0_28.Cmp(_dafny.Zero) == 0 {
							_nw214 = _dafny.NewArray(_len0_28)
						} else {
							var _init28 func(_dafny.Int) bool = func(_1246_i6 _dafny.Int) bool {
								return ((_this).F13()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ad")).Cardinality())))) < 0
							}
							_ = _init28
							var _element0_28 = _init28(_dafny.Zero)
							_ = _element0_28
							_nw214 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
							(_nw214).ArraySet1(_element0_28, 0)
							var _nativeLen0_28 = (_len0_28).Int()
							_ = _nativeLen0_28
							for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
								(_nw214).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
							}
						}
						_1245_v57 = _nw214
						var _1247_v58 _dafny.Sequence
						_ = _1247_v58
						_1247_v58 = _dafny.UnicodeSeqOfUtf8Bytes("xmboh")
						var _1248_v59 _dafny.Set
						_ = _1248_v59
						_1248_v59 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(134))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg63 _dafny.Int) interface{} {
								return coer63(arg63)
							}
						}((func(_1249_v55 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1250_i7 _dafny.Int) _dafny.CodePoint {
								return _1249_v55
							}
						})(_1243_v55)))).Cardinality()))
						var _1251_v60 D7
						_ = _1251_v60
						_1251_v60 = Companion_D7_.Create_DC18_(p0, _1247_v58, (_this).F14(), (_this).F14(), _1248_v59)
						var _1252_v61 _dafny.Sequence
						_ = _1252_v61
						_1252_v61 = _dafny.SeqOf(_this.F12())
						var _pat_let_tv25 = _1248_v59
						_ = _pat_let_tv25
						var _nwElement0_46 bool = !((_this).F14())
						_ = _nwElement0_46
						var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(11))
						_ = _nw215
						(_nw215).ArraySet1(_nwElement0_46, 0)
						(_nw215).ArraySet1((_this).F14(), 1)
						(_nw215).ArraySet1((_this).F14(), 2)
						(_nw215).ArraySet1(false, 3)
						(_nw215).ArraySet1(_this.F12(), 4)
						(_nw215).ArraySet1(!_dafny.Companion_Sequence_.Contains(_1247_v58, _1243_v55), 5)
						(_nw215).ArraySet1((_this).F14(), 6)
						(_nw215).ArraySet1(!((func(_pat_let45_0 D7) D7 {
							return func(_1253_dt__update__tmp_h1 D7) D7 {
								return func(_pat_let46_0 _dafny.Set) D7 {
									return func(_1254_dt__update_hcf44_h0 _dafny.Set) D7 {
										return Companion_D7_.Create_DC18_((_1253_dt__update__tmp_h1).Dtor_cf40(), (_1253_dt__update__tmp_h1).Dtor_cf41(), (_1253_dt__update__tmp_h1).Dtor_cf42(), (_1253_dt__update__tmp_h1).Dtor_cf43(), _1254_dt__update_hcf44_h0)
									}(_pat_let46_0)
								}(_pat_let_tv25)
							}(_pat_let45_0)
						}(_1251_v60)).Dtor_cf43()), 7)
						(_nw215).ArraySet1((_this).F14(), 8)
						(_nw215).ArraySet1(_this.F12(), 9)
						(_nw215).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1252_v61, _1252_v61), 10)
						_1245_v57 = _nw215
						var _1255_v62 *C2
						_ = _1255_v62
						var _nw216 *C2 = New_C2_()
						_ = _nw216
						_nw216.Ctor__(_this.F12(), (_dafny.MultiSetOf((_this).F14())).Update(Companion_Default___.Fm0(_dafny.CodePoint('c'), globalState), Companion_Default___.Abs((_this).F13())), (_this).F14())
						_1255_v62 = _nw216
						var _1256_v63 _dafny.Map
						_ = _1256_v63
						_1256_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1245_v57)
						(globalState).F8 = !((_1256_v63).Equals(_1256_v63)) || (_1255_v62.F15)
						var _1257_v64 *C2
						_ = _1257_v64
						var _nw217 *C2 = New_C2_()
						_ = _nw217
						_nw217.Ctor__(_this.F12(), _this.F11(), (_this).F14())
						_1257_v64 = _nw217
						(globalState).F9 = (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)
					} else {
						var _1258_v65 _dafny.Map
						_ = _1258_v65
						_1258_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
						var _1259_v66 _dafny.Sequence
						_ = _1259_v66
						_1259_v66 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tpnnpacy")).Cardinality()))
						var _1260_v67 _dafny.Sequence
						_ = _1260_v67
						_1260_v67 = _dafny.SeqOf(_1259_v66, _1259_v66)
						var _1261_v68 _dafny.Map
						_ = _1261_v68
						_1261_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1258_v65, _dafny.Companion_Sequence_.Concatenate((_1260_v67).Select((Companion_Default___.SafeIndex((_1259_v66).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1259_v66).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1260_v67).Cardinality()))).Uint32()).(_dafny.Sequence), _1259_v66))
						_1261_v68 = _1261_v68
						var _1262_v69 _dafny.Array
						_ = _1262_v69
						var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
						_ = _nw218
						_1262_v69 = _nw218
						var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1262_v69), 0))
						_ = _index188
						(_1262_v69).ArraySet1(_1152_v0, (_index188).Int())
						var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1262_v69), 0))
						_ = _index189
						(_1262_v69).ArraySet1(_1152_v0, (_index189).Int())
						var _1263_v70 *C2
						_ = _1263_v70
						var _nw219 *C2 = New_C2_()
						_ = _nw219
						_nw219.Ctor__(Companion_Default___.Fm0(_1243_v55, globalState), (Companion_Default___.Fm22(p0, _this.F12(), globalState)).Update((_this).F14(), Companion_Default___.Abs((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))), (_this).F14())
						_1263_v70 = _nw219
						var _1264_v71 _dafny.Map
						_ = _1264_v71
						_1264_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F11()).Update(_1263_v70.F15, Companion_Default___.Abs(p0)), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
						var _1265_v72 _dafny.Map
						_ = _1265_v72
						_1265_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v55, (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
						var _1266_v74 _dafny.Array
						_ = _1266_v74
						var _nwElement0_47 _dafny.Map = (_1265_v72).Update(_1243_v55, (_this).F13())
						_ = _nwElement0_47
						var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(8))
						_ = _nw220
						(_nw220).ArraySet1(_nwElement0_47, 0)
						(_nw220).ArraySet1(_1265_v72, 1)
						(_nw220).ArraySet1(_1265_v72, 2)
						(_nw220).ArraySet1(_1265_v72, 3)
						(_nw220).ArraySet1(func() _dafny.Map {
							var _coll28 = _dafny.NewMapBuilder()
							_ = _coll28
							for _iter32 := _dafny.Iterate((_dafny.SetOf(_1243_v55)).Elements()); ; {
								_compr_28, _ok32 := _iter32()
								if !_ok32 {
									break
								}
								var _1267_v73 _dafny.CodePoint
								_1267_v73 = interface{}(_compr_28).(_dafny.CodePoint)
								if (_dafny.SetOf(_1243_v55)).Contains(_1267_v73) {
									_coll28.Add(_1267_v73, (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))
								}
							}
							return _coll28.ToMap()
						}(), 4)
						(_nw220).ArraySet1(_1265_v72, 5)
						(_nw220).ArraySet1(_1265_v72, 6)
						(_nw220).ArraySet1(_1265_v72, 7)
						_1266_v74 = _nw220
						var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1266_v74), 0))
						_ = _index190
						(_1266_v74).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1243_v55, (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)), (_index190).Int())
						var _1268_v76 _dafny.MultiSet
						_ = _1268_v76
						_1268_v76 = _dafny.MultiSetOf(_this.F11())
						var _1269_v78 _dafny.Sequence
						_ = _1269_v78
						_1269_v78 = _dafny.SeqOf(_this.F11(), _this.F11())
						var _1270_v79 _dafny.Set
						_ = _1270_v79
						_1270_v79 = _dafny.SetOf(_this.F11(), (_1269_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1269_v78).Cardinality()))).Uint32()).(_dafny.MultiSet))
						var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
						_ = _index191
						var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
						_ = _index192
						var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1266_v74), 0))
						_ = _index193
						var _rhs202 *C2 = _1263_v70
						_ = _rhs202
						var _rhs203 _dafny.Int = (Companion_Default___.SafeDivisionInt(Companion_Default___.Fm1((_this).F14(), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), (_this).F13(), globalState), (_dafny.Zero).Minus((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)))).Times(p0)
						_ = _rhs203
						var _rhs204 _dafny.Map = (func() _dafny.Map {
							if ((_this).F13()).Cmp(_dafny.IntOfInt64(411)) != 0 {
								return (func() _dafny.Map {
									var _coll29 = _dafny.NewMapBuilder()
									_ = _coll29
									for _iter33 := _dafny.Iterate((_1268_v76).Elements()); ; {
										_compr_29, _ok33 := _iter33()
										if !_ok33 {
											break
										}
										var _1271_v75 _dafny.MultiSet
										_1271_v75 = interface{}(_compr_29).(_dafny.MultiSet)
										if (_1268_v76).Contains(_1271_v75) {
											_coll29.Add(_1271_v75, (_1258_v65).Cardinality())
										}
									}
									return _coll29.ToMap()
								}()).Merge(func() _dafny.Map {
									var _coll30 = _dafny.NewMapBuilder()
									_ = _coll30
									for _iter34 := _dafny.Iterate((_1270_v79).Elements()); ; {
										_compr_30, _ok34 := _iter34()
										if !_ok34 {
											break
										}
										var _1272_v77 _dafny.MultiSet
										_1272_v77 = interface{}(_compr_30).(_dafny.MultiSet)
										if (_1270_v79).Contains(_1272_v77) {
											_coll30.Add(_1272_v77, (_this).F13())
										}
									}
									return _coll30.ToMap()
								}())
							}
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11(), (_this).F13())).Merge(_1264_v71)
						})()
						_ = _rhs204
						var _rhs205 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(815), (_1258_v65).Cardinality())
						_ = _rhs205
						var _rhs206 _dafny.Map = _1265_v72
						_ = _rhs206
						var _lhs165 _dafny.Array = _1152_v0
						_ = _lhs165
						var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
						_ = _lhs166
						var _lhs167 _dafny.Array = _1152_v0
						_ = _lhs167
						var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
						_ = _lhs168
						var _lhs169 _dafny.Array = _1266_v74
						_ = _lhs169
						var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1266_v74), 0))
						_ = _lhs170
						_1263_v70 = _rhs202
						(_lhs165).ArraySet1(_rhs203, (_lhs166).Int())
						_1264_v71 = _rhs204
						(_lhs167).ArraySet1(_rhs205, (_lhs168).Int())
						(_lhs169).ArraySet1(_rhs206, (_lhs170).Int())
						var _1273_v80 T0
						_ = _1273_v80
						var _nw221 *C0 = New_C0_()
						_ = _nw221
						_nw221.Ctor__(_this.F12(), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), _this.F11(), _this.F12())
						_1273_v80 = _nw221
						var _1274_v81 _dafny.Sequence
						_ = _1274_v81
						_1274_v81 = _dafny.UnicodeSeqOfUtf8Bytes("xcisubu")
						var _1275_v82 _dafny.Map
						_ = _1275_v82
						_1275_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1273_v80, _1274_v81)
						_1275_v82 = (_1275_v82).Update(_1273_v80, _1274_v81)
						(globalState).F3 = Companion_Default___.Fm0(_1243_v55, globalState)
					}
					var _1276_v83 _dafny.Set
					_ = _1276_v83
					_1276_v83 = _dafny.SetOf(p0)
					var _1277_v84 _dafny.Array
					_ = _1277_v84
					var _nwElement0_48 bool = !(_this.F12()) || ((_this).F14())
					_ = _nwElement0_48
					var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(2))
					_ = _nw222
					(_nw222).ArraySet1(_nwElement0_48, 0)
					(_nw222).ArraySet1((_1276_v83).IsSubsetOf(_1276_v83), 1)
					_1277_v84 = _nw222
					var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_1277_v84), 0))
					_ = _index194
					(_1277_v84).ArraySet1((func() bool {
						if _this.F12() {
							return Companion_Default___.Fm0(_1243_v55, globalState)
						}
						return (_this).F14()
					})(), (_index194).Int())
					var _1278_v85 _dafny.Sequence
					_ = _1278_v85
					_1278_v85 = _dafny.UnicodeSeqOfUtf8Bytes("hbi")
					var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_1277_v84), 0))
					_ = _index195
					(_1277_v84).ArraySet1(_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
						if (_this).Fm6((_this).F14(), p0, globalState) {
							return Companion_Default___.Fm9(p0, false, globalState)
						}
						return _1278_v85
					})(), _1278_v85), (_index195).Int())
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1279_v86 _dafny.Set
		_ = _1279_v86
		_1279_v86 = _dafny.SetOf(_this.F12(), false, (_this).F14())
		var _1280_v87 _dafny.Sequence
		_ = _1280_v87
		_1280_v87 = _dafny.UnicodeSeqOfUtf8Bytes("kjje")
		var _1281_v88 _dafny.Set
		_ = _1281_v88
		_1281_v88 = _dafny.SetOf((_this).F13())
		var _hi2 _dafny.Int = (Companion_D7_.Create_DC18_((_1279_v86).Cardinality(), _1280_v87, false, (_this).F14(), _1281_v88)).Dtor_cf40()
		_ = _hi2
		for _1282_i8 := (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int); _1282_i8.Cmp(_hi2) < 0; _1282_i8 = _1282_i8.Plus(_dafny.One) {
			if (_this).F14() {
				var _1283_v89 _dafny.Array
				_ = _1283_v89
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_29
				var _nw223 _dafny.Array
				_ = _nw223
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw223 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) bool = func(_1284_i9 _dafny.Int) bool {
						return (_this).F14()
					}
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw223 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw223).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw223).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1283_v89 = _nw223
				var _1285_v90 _dafny.Map
				_ = _1285_v90
				_1285_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_i8, _1283_v89)
				_1285_v90 = ((_1285_v90).Merge(_1285_v90)).Merge(_1285_v90)
				var _1286_v91 _dafny.Sequence
				_ = _1286_v91
				_1286_v91 = _dafny.SeqOf(_1282_i8)
				(globalState).F9 = (func() _dafny.Int {
					if _this.F12() {
						return (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)
					}
					return ((_1286_v91).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1286_v91).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_this).F13())
				})()
				_1285_v90 = _1285_v90
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _index196
				(_1152_v0).ArraySet1(Companion_Default___.SafeModuloInt((_this).F13(), (_this).F13()), (_index196).Int())
				var _1287_v92 _dafny.Map
				_ = _1287_v92
				_1287_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)
				var _1288_v93 _dafny.Sequence
				_ = _1288_v93
				_1288_v93 = _dafny.SeqOf(_1287_v92)
				var _1289_v94 _dafny.Map
				_ = _1289_v94
				_1289_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _1280_v87)
				var _1290_v95 _dafny.Map
				_ = _1290_v95
				_1290_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1287_v92)
				var _1291_v96 _dafny.Map
				_ = _1291_v96
				_1291_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), _this.F12())
				var _1292_v97 _dafny.Array
				_ = _1292_v97
				var _nwElement0_49 _dafny.Map = _1287_v92
				_ = _nwElement0_49
				var _nw224 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(19))
				_ = _nw224
				(_nw224).ArraySet1(_nwElement0_49, 0)
				(_nw224).ArraySet1(_1287_v92, 1)
				(_nw224).ArraySet1(((_1288_v93).Select((Companion_Default___.SafeIndex((_1289_v94).Cardinality(), _dafny.IntOfUint32((_1288_v93).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1287_v92), 2)
				(_nw224).ArraySet1(_1287_v92, 3)
				(_nw224).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int))).Merge(_1287_v92), 4)
				(_nw224).ArraySet1(_1287_v92, 5)
				(_nw224).ArraySet1((_1287_v92).Merge(_1287_v92), 6)
				(_nw224).ArraySet1((_1287_v92).Update(false, (_this).F13()), 7)
				(_nw224).ArraySet1((Companion_Default___.Fm23(p0, globalState)).Merge(_1287_v92), 8)
				(_nw224).ArraySet1(_1287_v92, 9)
				(_nw224).ArraySet1(_1287_v92, 10)
				(_nw224).ArraySet1((Companion_Default___.Fm23((_this).F13(), globalState)).Update((_this).F14(), (_this).F13()), 11)
				(_nw224).ArraySet1((_1287_v92).Merge(_1287_v92), 12)
				(_nw224).ArraySet1(_1287_v92, 13)
				(_nw224).ArraySet1((_1287_v92).Merge(_1287_v92), 14)
				(_nw224).ArraySet1((_1287_v92).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)), 15)
				(_nw224).ArraySet1((func() _dafny.Map {
					if (_1290_v95).Contains((_1291_v96).Cardinality()) {
						return (_1290_v95).Get((_1291_v96).Cardinality()).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1282_i8)
				})(), 16)
				(_nw224).ArraySet1((_1287_v92).Merge(_1287_v92), 17)
				(_nw224).ArraySet1(_1287_v92, 18)
				_1292_v97 = _nw224
				var _rhs207 _dafny.Array = _1292_v97
				_ = _rhs207
				var _rhs208 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("l")
				_ = _rhs208
				_1292_v97 = _rhs207
				_1280_v87 = _rhs208
			} else {
				var _1293_v98 _dafny.Array
				_ = _1293_v98
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_30
				var _nw225 _dafny.Array
				_ = _nw225
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw225 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) bool = func(_1294_i10 _dafny.Int) bool {
						return true
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw225 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw225).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw225).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1293_v98 = _nw225
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1293_v98), 0))
				_ = _index197
				(_1293_v98).ArraySet1(_this.F12(), (_index197).Int())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1293_v98), 0))
				_ = _index198
				(_1293_v98).ArraySet1((_this).F14(), (_index198).Int())
				(globalState).F1 = _1152_v0
				var _1295_v99 _dafny.Sequence
				_ = _1295_v99
				_1295_v99 = _dafny.SeqOf(_1279_v86, (_1279_v86).Difference(_1279_v86))
				var _1296_v100 _dafny.CodePoint
				_ = _1296_v100
				_1296_v100 = _dafny.CodePoint('q')
				var _1297_v101 _dafny.Sequence
				_ = _1297_v101
				_1297_v101 = _dafny.SeqOf(_1280_v87)
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1293_v98), 0))
				_ = _index199
				var _rhs209 _dafny.Int = ((_1295_v99).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.Fm1(true, _1282_i8, (_dafny.Zero).Minus(_1282_i8), globalState)), _dafny.IntOfUint32((_1295_v99).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
				_ = _rhs209
				var _rhs210 bool = (func() bool {
					if (_this).F14() {
						return _this.F12()
					}
					return (_1293_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1293_v98), 0))).Int()).(bool)
				})()
				_ = _rhs210
				var _rhs211 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1280_v87, (Companion_Default___.SafeIndex(((_this.F11()).Union(_this.F11())).Cardinality(), _dafny.IntOfUint32((_1280_v87).Cardinality()))).Uint32(), _1296_v100)
				_ = _rhs211
				var _rhs212 bool = _dafny.Companion_Sequence_.Contains((_1297_v101).Select((Companion_Default___.SafeIndex((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1297_v101).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.CodePoint('d'))
				_ = _rhs212
				var _rhs213 bool = !(_this.F12())
				_ = _rhs213
				var _lhs171 *GlobalState = globalState
				_ = _lhs171
				var _lhs172 *C3 = _this
				_ = _lhs172
				var _lhs173 *GlobalState = globalState
				_ = _lhs173
				var _lhs174 _dafny.Array = _1293_v98
				_ = _lhs174
				var _lhs175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1293_v98), 0))
				_ = _lhs175
				_lhs171.F9 = _rhs209
				_lhs172.F12_set_(_rhs210)
				_1280_v87 = _rhs211
				_lhs173.F8 = _rhs212
				(_lhs174).ArraySet1(_rhs213, (_lhs175).Int())
				Companion_Default___.M0(globalState)
				var _1298_v102 _dafny.Array
				_ = _1298_v102
				var _nw226 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
				_ = _nw226
				_1298_v102 = _nw226
				var _1299_v103 _dafny.Map
				_ = _1299_v103
				_1299_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1293_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_1293_v98), 0))).Int()).(bool), _1298_v102)
				var _1300_v104 D9
				_ = _1300_v104
				_1300_v104 = Companion_D9_.Create_DC23_(_1299_v103)
				var _pat_let_tv26 = _1299_v103
				_ = _pat_let_tv26
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _index200
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _index201
				var _rhs214 D9 = func(_pat_let47_0 D9) D9 {
					return func(_1301_dt__update__tmp_h2 D9) D9 {
						return func(_pat_let48_0 _dafny.Map) D9 {
							return func(_1302_dt__update_hcf55_h0 _dafny.Map) D9 {
								return Companion_D9_.Create_DC23_(_1302_dt__update_hcf55_h0)
							}(_pat_let48_0)
						}(_pat_let_tv26)
					}(_pat_let47_0)
				}(_1300_v104)
				_ = _rhs214
				var _rhs215 _dafny.Int = (_this).F13()
				_ = _rhs215
				var _rhs216 _dafny.Int = p0
				_ = _rhs216
				var _rhs217 _dafny.Int = _dafny.IntOfInt64(-637)
				_ = _rhs217
				var _lhs176 _dafny.Array = _1152_v0
				_ = _lhs176
				var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _lhs177
				var _lhs178 *GlobalState = globalState
				_ = _lhs178
				var _lhs179 _dafny.Array = _1152_v0
				_ = _lhs179
				var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
				_ = _lhs180
				_1300_v104 = _rhs214
				(_lhs176).ArraySet1(_rhs215, (_lhs177).Int())
				_lhs178.F9 = _rhs216
				(_lhs179).ArraySet1(_rhs217, (_lhs180).Int())
			}
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
			_ = _index202
			(_1152_v0).ArraySet1((_this).F13(), (_index202).Int())
			var _1303_v105 _dafny.Sequence
			_ = _1303_v105
			_1303_v105 = _dafny.SeqOf(_this.F12(), _this.F12())
			var _1304_v106 _dafny.Map
			_ = _1304_v106
			_1304_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F14())
			var _1305_v107 _dafny.Array
			_ = _1305_v107
			var _nwElement0_50 bool = (_this).F14()
			_ = _nwElement0_50
			var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(8))
			_ = _nw227
			(_nw227).ArraySet1(_nwElement0_50, 0)
			(_nw227).ArraySet1(!((func() bool {
				if (_this).F14() {
					return true
				}
				return _this.F12()
			})()), 1)
			(_nw227).ArraySet1(!(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1303_v105, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1303_v105).Cardinality()))).Uint32(), _this.F12())).Cardinality()))).Cmp((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)) == 0), 2)
			(_nw227).ArraySet1((func() bool {
				if (_1304_v106).Contains((_this).F14()) {
					return (_1304_v106).Get((_this).F14()).(bool)
				}
				return _this.F12()
			})(), 3)
			(_nw227).ArraySet1(_this.F12(), 4)
			(_nw227).ArraySet1((_this).F14(), 5)
			(_nw227).ArraySet1(_this.F12(), 6)
			(_nw227).ArraySet1(_this.F12(), 7)
			_1305_v107 = _nw227
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1305_v107), 0))
			_ = _index203
			(_1305_v107).ArraySet1(_dafny.Companion_Sequence_.Equal(_1280_v87, _1280_v87), (_index203).Int())
			var _1306_v108 _dafny.Array
			_ = _1306_v108
			var _nw228 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(9))
			_ = _nw228
			_1306_v108 = _nw228
			var _1307_v109 D8
			_ = _1307_v109
			_1307_v109 = Companion_D8_.Create_DC19_(_this.F11())
			var _1308_v110 D8
			_ = _1308_v110
			_1308_v110 = Companion_D8_.Create_DC22_(_1307_v109)
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1306_v108), 0))
			_ = _index204
			(_1306_v108).ArraySet1(_1308_v110, (_index204).Int())
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1305_v107), 0))
			_ = _index205
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
			_ = _index206
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1306_v108), 0))
			_ = _index207
			var _rhs218 bool = !((_1281_v88).IsSubsetOf(_1281_v88))
			_ = _rhs218
			var _rhs219 _dafny.Int = (func() _dafny.Int {
				if _this.F12() {
					return ((_this).F13()).Plus(_dafny.IntOfInt64(466))
				}
				return _1282_i8
			})()
			_ = _rhs219
			var _rhs220 D8 = _1308_v110
			_ = _rhs220
			var _rhs221 _dafny.Int = (_this).F13()
			_ = _rhs221
			var _rhs222 _dafny.Int = p0
			_ = _rhs222
			var _lhs181 _dafny.Array = _1305_v107
			_ = _lhs181
			var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_1305_v107), 0))
			_ = _lhs182
			var _lhs183 _dafny.Array = _1152_v0
			_ = _lhs183
			var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
			_ = _lhs184
			var _lhs185 _dafny.Array = _1306_v108
			_ = _lhs185
			var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1306_v108), 0))
			_ = _lhs186
			var _lhs187 *GlobalState = globalState
			_ = _lhs187
			var _lhs188 *GlobalState = globalState
			_ = _lhs188
			(_lhs181).ArraySet1(_rhs218, (_lhs182).Int())
			(_lhs183).ArraySet1(_rhs219, (_lhs184).Int())
			(_lhs185).ArraySet1(_rhs220, (_lhs186).Int())
			_lhs187.F9 = _rhs221
			_lhs188.F9 = _rhs222
			var _1309_v111 _dafny.Map
			_ = _1309_v111
			_1309_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1152_v0, _this.F12())
			_1309_v111 = (_1309_v111).Update(_1152_v0, _dafny.Companion_Sequence_.Equal(_1280_v87, _1280_v87))
		}
		var _1310_v112 _dafny.CodePoint
		_ = _1310_v112
		_1310_v112 = _dafny.CodePoint('x')
		var _1311_v113 _dafny.Sequence
		_ = _1311_v113
		_1311_v113 = _dafny.SeqOf(!(_this.F12()), Companion_Default___.Fm0(_1310_v112, globalState))
		var _1312_v114 _dafny.Map
		_ = _1312_v114
		_1312_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F12()), _dafny.IntOfUint32((_1311_v113).Cardinality()))
		var _1313_v115 _dafny.Sequence
		_ = _1313_v115
		_1313_v115 = _dafny.SeqOf(p0, (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tyfnjxj")).Cardinality()), (_this).F13(), (func() _dafny.Int {
			if (_1312_v114).Contains(_this.F12()) {
				return (_1312_v114).Get(_this.F12()).(_dafny.Int)
			}
			return (_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)
		})())
		var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))
		_ = _index208
		(_1152_v0).ArraySet1((((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_1313_v115).Cardinality()))).Plus((_1152_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_1152_v0), 0))).Int()).(_dafny.Int)), (_index208).Int())
	}
}
func (_this *C3) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *C3) F14() bool {
	{
		return _this._f14
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f11 _dafny.MultiSet
	_f12 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f11 = _dafny.EmptyMultiSet
	_this._f12 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F11() _dafny.MultiSet {
	return _this._f11
}
func (_this *C4) F11_set_(value _dafny.MultiSet) {
	_this._f11 = value
}
func (_this *C4) F12() bool {
	return _this._f12
}
func (_this *C4) F12_set_(value bool) {
	_this._f12 = value
}
func (_this *C4) Ctor__(f11 _dafny.MultiSet, f12 bool) {
	{
		(_this)._f11 = f11
		(_this)._f12 = f12
	}
}
func (_this *C4) Fm2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Map, p3 bool, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(896)).Cmp(((_dafny.Zero).Minus(_dafny.IntOfInt64(-394))).Times(_dafny.IntOfInt64(46))) != 0
	}
}
func (_this *C4) Fm3(p0 bool, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(9)
	}
}
func (_this *C4) Fm4(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Int {
			if _this.F12() {
				return _dafny.IntOfInt64(-538)
			}
			return _dafny.IntOfInt64(-626)
		})()).Cmp(((_dafny.MultiSetOf(_dafny.IntOfInt64(908))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-317), _dafny.IntOfInt64(163)))).Cardinality()) == 0
	}
}
func (_this *C4) M1(p0 _dafny.Set, globalState *GlobalState) (_dafny.Array, _dafny.Array, _dafny.Int) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1314_v0 _dafny.Int
		_ = _1314_v0
		_1314_v0 = _dafny.IntOfInt64(-835)
		var _1315_v1 _dafny.Sequence
		_ = _1315_v1
		_1315_v1 = _dafny.UnicodeSeqOfUtf8Bytes("fahubdda")
		var _rhs223 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm1(_this.F12(), _1314_v0, _1314_v0, globalState)).Times(_1314_v0))
		_ = _rhs223
		var _rhs224 bool = _this.F12()
		_ = _rhs224
		var _rhs225 _dafny.MultiSet = (_this.F11()).Update(_this.F12(), Companion_Default___.Abs((_dafny.Zero).Minus((_dafny.IntOfInt64(36)).Plus(_1314_v0))))
		_ = _rhs225
		var _rhs226 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_this.F12(), _this.F12(), _this.F12(), _1314_v0, globalState), Companion_Default___.Fm5(_this.F12(), _this.F12(), _this.F12(), _dafny.IntOfUint32((_1315_v1).Cardinality()), globalState))
		_ = _rhs226
		var _lhs189 *GlobalState = globalState
		_ = _lhs189
		var _lhs190 *GlobalState = globalState
		_ = _lhs190
		var _lhs191 *C4 = _this
		_ = _lhs191
		var _lhs192 *GlobalState = globalState
		_ = _lhs192
		_lhs189.F9 = _rhs223
		_lhs190.F7 = _rhs224
		_lhs191.F11_set_(_rhs225)
		_lhs192.F6 = _rhs226
		var _1316_v2 _dafny.Array
		_ = _1316_v2
		var _nw229 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw229
		_1316_v2 = _nw229
		var _1317_v3 _dafny.Set
		_ = _1317_v3
		_1317_v3 = _dafny.SetOf(_1316_v2, _1316_v2, _1316_v2, _1316_v2)
		var _1318_v4 _dafny.Sequence
		_ = _1318_v4
		_1318_v4 = _dafny.SeqOf(_1314_v0, Companion_Default___.Fm1(_this.F12(), _1314_v0, _1314_v0, globalState), _1314_v0, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kd")).Cardinality())).Plus(_1314_v0))
		var _1319_v5 _dafny.Sequence
		_ = _1319_v5
		_1319_v5 = _dafny.SeqOf(p0, p0)
		var _1320_v6 D1
		_ = _1320_v6
		_1320_v6 = Companion_D1_.Create_DC1_(p0)
		var _1321_v7 _dafny.Array
		_ = _1321_v7
		var _nwElement0_51 _dafny.Set = p0
		_ = _nwElement0_51
		var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(29))
		_ = _nw230
		(_nw230).ArraySet1(_nwElement0_51, 0)
		(_nw230).ArraySet1(p0, 1)
		(_nw230).ArraySet1((p0).Intersection(p0), 2)
		(_nw230).ArraySet1((p0).Union(p0), 3)
		(_nw230).ArraySet1(_dafny.SetOf(_this.F12(), (_this.F12()), _this.F12(), !(_this.F12())), 4)
		(_nw230).ArraySet1((p0).Union(p0), 5)
		(_nw230).ArraySet1(p0, 6)
		(_nw230).ArraySet1((p0).Union(p0), 7)
		(_nw230).ArraySet1(p0, 8)
		(_nw230).ArraySet1((p0).Difference((_1319_v5).Select((Companion_Default___.SafeIndex(_1314_v0, _dafny.IntOfUint32((_1319_v5).Cardinality()))).Uint32()).(_dafny.Set)), 9)
		(_nw230).ArraySet1(p0, 10)
		(_nw230).ArraySet1((_1320_v6).Dtor_cf1(), 11)
		(_nw230).ArraySet1(p0, 12)
		(_nw230).ArraySet1(p0, 13)
		(_nw230).ArraySet1(p0, 14)
		(_nw230).ArraySet1(p0, 15)
		(_nw230).ArraySet1(p0, 16)
		(_nw230).ArraySet1((_1319_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-249), _dafny.IntOfUint32((_1319_v5).Cardinality()))).Uint32()).(_dafny.Set), 17)
		(_nw230).ArraySet1(p0, 18)
		(_nw230).ArraySet1((p0).Intersection(p0), 19)
		(_nw230).ArraySet1(p0, 20)
		(_nw230).ArraySet1((p0).Difference(_dafny.SetOf(_this.F12(), _this.F12())), 21)
		(_nw230).ArraySet1((func() _dafny.Set {
			if _this.F12() {
				return p0
			}
			return p0
		})(), 22)
		(_nw230).ArraySet1((_dafny.SetOf(!(false))).Union(p0), 23)
		(_nw230).ArraySet1((p0).Intersection(_dafny.SetOf(_this.F12())), 24)
		(_nw230).ArraySet1(p0, 25)
		(_nw230).ArraySet1(_dafny.SetOf(_this.F12(), _this.F12()), 26)
		(_nw230).ArraySet1((p0).Difference(p0), 27)
		(_nw230).ArraySet1(p0, 28)
		_1321_v7 = _nw230
		var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))
		_ = _index209
		(_1321_v7).ArraySet1((p0).Difference(p0), (_index209).Int())
		var _1322_v8 _dafny.Sequence
		_ = _1322_v8
		_1322_v8 = _dafny.SeqOf(_1317_v3)
		var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))
		_ = _index210
		var _rhs227 _dafny.Set = (_1322_v8).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(662), (_dafny.Zero).Minus(_1314_v0)), _dafny.IntOfUint32((_1322_v8).Cardinality()))).Uint32()).(_dafny.Set)
		_ = _rhs227
		var _rhs228 _dafny.Sequence = _1318_v4
		_ = _rhs228
		var _rhs229 _dafny.Set = p0
		_ = _rhs229
		var _rhs230 bool = _this.F12()
		_ = _rhs230
		var _rhs231 _dafny.Int = _dafny.IntOfInt64(922)
		_ = _rhs231
		var _lhs193 _dafny.Array = _1321_v7
		_ = _lhs193
		var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))
		_ = _lhs194
		var _lhs195 *GlobalState = globalState
		_ = _lhs195
		_1317_v3 = _rhs227
		_1318_v4 = _rhs228
		(_lhs193).ArraySet1(_rhs229, (_lhs194).Int())
		_lhs195.F3 = _rhs230
		_1314_v0 = _rhs231
		_1316_v2 = _1316_v2
		var _1323_i0 _dafny.Int
		_ = _1323_i0
		_1323_i0 = _dafny.Zero
		{
			for _this.F12() {
				{
					if (_1323_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1323_i0 = (_1323_i0).Plus(_dafny.One)
					(globalState).F9 = (_1314_v0).Minus(_1314_v0)
					(globalState).F3 = !(_this.F12())
					if _this.F12() {
						var _1324_v9 D1
						_ = _1324_v9
						_1324_v9 = Companion_D1_.Create_DC1_((_1321_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))).Int()).(_dafny.Set))
						var _1325_v10 _dafny.MultiSet
						_ = _1325_v10
						_1325_v10 = _dafny.MultiSetOf(Companion_D1_.Create_DC3_(_1324_v9))
						var _1326_v11 D1
						_ = _1326_v11
						_1326_v11 = Companion_D1_.Create_DC3_(_1324_v9)
						(globalState).F3 = ((_1325_v10).Difference(_1325_v10)).IsProperSubsetOf((_1325_v10).Update(_1326_v11, Companion_Default___.Abs(_1314_v0)))
						var _1327_v12 *C3
						_ = _1327_v12
						var _nw231 *C3 = New_C3_()
						_ = _nw231
						_nw231.Ctor__((_dafny.Zero).Minus((func() _dafny.Int {
							if _this.F12() {
								return _1314_v0
							}
							return _1314_v0
						})()), _this.F12(), (_dafny.MultiSetOf(_this.F12(), _this.F12())).Difference(_dafny.MultiSetOf(!(_this.F12()), _this.F12())), (_this.F11()).IsDisjointFrom(_this.F11()))
						_1327_v12 = _nw231
						var _1328_v13 _dafny.Map
						_ = _1328_v13
						_1328_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1327_v12).F13(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_1327_v12).F13())))
						var _1329_v14 _dafny.Set
						_ = _1329_v14
						_1329_v14 = _dafny.SetOf((_1327_v12).F13(), (_1327_v12).F13(), (func() _dafny.Int {
							if (_1328_v13).Contains(_dafny.IntOfInt64(-708)) {
								return (_1328_v13).Get(_dafny.IntOfInt64(-708)).(_dafny.Int)
							}
							return (_1327_v12).F13()
						})())
						var _rhs232 _dafny.Set = _1329_v14
						_ = _rhs232
						var _rhs233 _dafny.Int = _1314_v0
						_ = _rhs233
						var _rhs234 _dafny.Int = _1314_v0
						_ = _rhs234
						var _rhs235 _dafny.Sequence = (func() _dafny.Sequence {
							if (_1327_v12).F14() {
								return _1318_v4
							}
							return _dafny.Companion_Sequence_.Concatenate(_1318_v4, _1318_v4)
						})()
						_ = _rhs235
						var _lhs196 *GlobalState = globalState
						_ = _lhs196
						var _lhs197 *GlobalState = globalState
						_ = _lhs197
						_1329_v14 = _rhs232
						_lhs196.F9 = _rhs233
						_lhs197.F9 = _rhs234
						_1318_v4 = _rhs235
						var _1330_v15 _dafny.Array
						_ = _1330_v15
						var _nw232 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
						_ = _nw232
						_1330_v15 = _nw232
						var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1330_v15), 0))
						_ = _index211
						(_1330_v15).ArraySet1((func() _dafny.Array {
							if _this.F12() {
								return _1316_v2
							}
							return _1316_v2
						})(), (_index211).Int())
						var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1330_v15), 0))
						_ = _index212
						var _rhs236 _dafny.Int = (_1314_v0).Minus((func() _dafny.Int {
							if (_1327_v12).F14() {
								return (_1327_v12).F13()
							}
							return _1314_v0
						})())
						_ = _rhs236
						var _rhs237 _dafny.Array = _1316_v2
						_ = _rhs237
						var _lhs198 _dafny.Array = _1330_v15
						_ = _lhs198
						var _lhs199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_1330_v15), 0))
						_ = _lhs199
						_1314_v0 = _rhs236
						(_lhs198).ArraySet1(_rhs237, (_lhs199).Int())
						var _1331_v16 _dafny.Map
						_ = _1331_v16
						_1331_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1318_v4).Select((Companion_Default___.SafeIndex((_1327_v12).F13(), _dafny.IntOfUint32((_1318_v4).Cardinality()))).Uint32()).(_dafny.Int), _1315_v1)
						var _rhs238 _dafny.Int = (_1327_v12).F13()
						_ = _rhs238
						var _rhs239 bool = (_1327_v12).F14()
						_ = _rhs239
						var _rhs240 bool = (_1327_v12).F14()
						_ = _rhs240
						var _rhs241 _dafny.Sequence = (func() _dafny.Sequence {
							if (_1331_v16).Contains(_1314_v0) {
								return (_1331_v16).Get(_1314_v0).(_dafny.Sequence)
							}
							return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tghqo"), _1315_v1)
						})()
						_ = _rhs241
						var _lhs200 *GlobalState = globalState
						_ = _lhs200
						var _lhs201 *GlobalState = globalState
						_ = _lhs201
						var _lhs202 *GlobalState = globalState
						_ = _lhs202
						_lhs200.F9 = _rhs238
						_lhs201.F8 = _rhs239
						_lhs202.F8 = _rhs240
						_1315_v1 = _rhs241
					} else {
						var _1332_v17 D9
						_ = _1332_v17
						_1332_v17 = Companion_D9_.Create_DC25_(_1315_v1, _this.F12(), _1314_v0)
						var _1333_v18 _dafny.Sequence
						_ = _1333_v18
						_1333_v18 = _dafny.SeqOf(_1332_v17, _1332_v17)
						var _1334_v19 *C3
						_ = _1334_v19
						var _nw233 *C3 = New_C3_()
						_ = _nw233
						_nw233.Ctor__(_1314_v0, _this.F12(), _this.F11(), _this.F12())
						_1334_v19 = _nw233
						var _1335_v20 _dafny.Sequence
						_ = _1335_v20
						_1335_v20 = _dafny.SeqOf(_1334_v19, _1334_v19)
						var _1336_v21 _dafny.Map
						_ = _1336_v21
						_1336_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1334_v19).F14()), _this.F12())
						var _1337_v22 _dafny.Sequence
						_ = _1337_v22
						_1337_v22 = _dafny.SeqOf(true, _this.F12(), (_1334_v19).F14())
						var _1338_v23 _dafny.CodePoint
						_ = _1338_v23
						_1338_v23 = _dafny.CodePoint('v')
						var _1339_v24 D1
						_ = _1339_v24
						_1339_v24 = Companion_D1_.Create_DC2_((_1334_v19).F14(), _this.F12(), false, _1338_v23, _1338_v23)
						var _1340_v25 T0
						_ = _1340_v25
						var _nw234 *C1 = New_C1_()
						_ = _nw234
						_nw234.Ctor__(_dafny.IntOfUint32((_1337_v22).Cardinality()), _dafny.MultiSetOf((_1339_v24).Dtor_cf4()), !(true))
						_1340_v25 = _nw234
						var _1341_v26 _dafny.Set
						_ = _1341_v26
						_1341_v26 = _dafny.SetOf(_1314_v0)
						var _1342_v27 _dafny.Array
						_ = _1342_v27
						var _len0_31 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_31
						var _nw235 _dafny.Array
						_ = _nw235
						if _len0_31.Cmp(_dafny.Zero) == 0 {
							_nw235 = _dafny.NewArray(_len0_31)
						} else {
							var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1343_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1344_i1 _dafny.Int) _dafny.Sequence {
									return _1343_v1
								}
							})(_1315_v1)
							_ = _init31
							var _element0_31 = _init31(_dafny.Zero)
							_ = _element0_31
							_nw235 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
							(_nw235).ArraySet1(_element0_31, 0)
							var _nativeLen0_31 = (_len0_31).Int()
							_ = _nativeLen0_31
							for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
								(_nw235).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
							}
						}
						_1342_v27 = _nw235
						var _1345_v28 D5
						_ = _1345_v28
						_1345_v28 = Companion_D5_.Create_DC13_((_1334_v19).F13(), true, (_1334_v19).F13(), _1342_v27, _1340_v25.F12())
						var _1346_v29 _dafny.Map
						_ = _1346_v29
						_1346_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1345_v28).Dtor_cf34(), (_1334_v19).F13())
						var _1347_v30 *C0
						_ = _1347_v30
						var _nw236 *C0 = New_C0_()
						_ = _nw236
						_nw236.Ctor__(true, _1314_v0, _this.F11(), _1340_v25.F12())
						_1347_v30 = _nw236
						var _1348_v31 _dafny.Sequence
						_ = _1348_v31
						_1348_v31 = _dafny.SeqOf(_1347_v30)
						var _1349_v32 _dafny.Array
						_ = _1349_v32
						var _nwElement0_52 _dafny.Int = _1314_v0
						_ = _nwElement0_52
						var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(19))
						_ = _nw237
						(_nw237).ArraySet1(_nwElement0_52, 0)
						(_nw237).ArraySet1(_1314_v0, 1)
						(_nw237).ArraySet1(((func() _dafny.Int {
							if (_this.F11()).Contains(_this.F12()) {
								return (_this.F11()).Multiplicity(_this.F12())
							}
							return _1314_v0
						})()).Plus(_1314_v0), 2)
						(_nw237).ArraySet1(_dafny.IntOfUint32((_1333_v18).Cardinality()), 3)
						(_nw237).ArraySet1(_1314_v0, 4)
						(_nw237).ArraySet1(_1314_v0, 5)
						(_nw237).ArraySet1(_1314_v0, 6)
						(_nw237).ArraySet1(_1314_v0, 7)
						(_nw237).ArraySet1(_dafny.IntOfUint32((_1335_v20).Cardinality()), 8)
						(_nw237).ArraySet1((_1336_v21).Cardinality(), 9)
						(_nw237).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1334_v19).F13(), _1340_v25)).Cardinality(), 10)
						(_nw237).ArraySet1(Companion_Default___.SafeDivisionInt((_1334_v19).F13(), (_1334_v19).F13()), 11)
						(_nw237).ArraySet1(((_1341_v26).Difference(_1341_v26)).Cardinality(), 12)
						(_nw237).ArraySet1((_dafny.Zero).Minus((_1334_v19).F13()), 13)
						(_nw237).ArraySet1((func() _dafny.Int {
							if (_1346_v29).Contains(_1340_v25.F12()) {
								return (_1346_v29).Get(_1340_v25.F12()).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_1348_v31).Cardinality())
						})(), 14)
						(_nw237).ArraySet1((_1347_v30).F18(), 15)
						(_nw237).ArraySet1(_dafny.IntOfInt64(278), 16)
						(_nw237).ArraySet1((_1334_v19).F13(), 17)
						(_nw237).ArraySet1((_1334_v19).F13(), 18)
						_1349_v32 = _nw237
						var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1349_v32), 0))
						_ = _index213
						(_1349_v32).ArraySet1(Companion_Default___.SafeDivisionInt(_1314_v0, _1314_v0), (_index213).Int())
						var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1349_v32), 0))
						_ = _index214
						(_1349_v32).ArraySet1(_1314_v0, (_index214).Int())
						var _1350_v33 *C0
						_ = _1350_v33
						var _nw238 *C0 = New_C0_()
						_ = _nw238
						_nw238.Ctor__((_1334_v19).F14(), Companion_Default___.SafeModuloInt((p0).Cardinality(), (_1334_v19).F13()), _this.F11(), ((_dafny.MultiSetOf(_1318_v4)).Cardinality()).Cmp(_1314_v0) != 0)
						_1350_v33 = _nw238
						var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_1316_v2), 0))
						_ = _index215
						(_1316_v2).ArraySet1((_1347_v30).F17(), (_index215).Int())
						var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_1316_v2), 0))
						_ = _index216
						(_1316_v2).ArraySet1((_1347_v30).F17(), (_index216).Int())
						(globalState).F8 = (func() bool {
							if ((_1334_v19).F13()).Cmp(_dafny.IntOfInt64(591)) == 0 {
								return _1340_v25.F12()
							}
							return _this.F12()
						})()
						var _1351_v34 _dafny.Map
						_ = _1351_v34
						_1351_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_1349_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1349_v32), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg64 _dafny.Int) interface{} {
								return coer64(arg64)
							}
						}((func(_1352_v23 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1353_i2 _dafny.Int) _dafny.CodePoint {
								return _1352_v23
							}
						})(_1338_v23)))).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(336))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}((func(_1354_v19 *C3) func(_dafny.Int) _dafny.Int {
							return func(_1355_i3 _dafny.Int) _dafny.Int {
								return (_1354_v19).F13()
							}
						})(_1334_v19))))
						_1351_v34 = (func() _dafny.Map {
							if (func() bool {
								if !((_1316_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen((_1316_v2), 0))).Int()).(bool)) {
									return (_1350_v33).F17()
								}
								return (_1350_v33).F17()
							})() {
								return _1351_v34
							}
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1350_v33).F18(), _1318_v4)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1349_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_1349_v32), 0))).Int()).(_dafny.Int), _1318_v4)).Update(_dafny.IntOfInt64(907), _1318_v4))
						})()
					}
					_1315_v1 = Companion_Default___.Fm9((_dafny.MultiSetOf(_1314_v0)).Cardinality(), _this.F12(), globalState)
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1356_v35 _dafny.CodePoint
		_ = _1356_v35
		_1356_v35 = _dafny.CodePoint('m')
		var _1357_v36 _dafny.Array
		_ = _1357_v36
		var _nwElement0_53 _dafny.CodePoint = _1356_v35
		_ = _nwElement0_53
		var _nw239 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(5))
		_ = _nw239
		(_nw239).ArraySet1CodePoint(_nwElement0_53, 0)
		(_nw239).ArraySet1CodePoint(_1356_v35, 1)
		(_nw239).ArraySet1CodePoint(_1356_v35, 2)
		(_nw239).ArraySet1CodePoint(_1356_v35, 3)
		(_nw239).ArraySet1CodePoint(_1356_v35, 4)
		_1357_v36 = _nw239
		var _1358_v37 D1
		_ = _1358_v37
		_1358_v37 = Companion_D1_.Create_DC2_(_this.F12(), _this.F12(), _this.F12(), _1356_v35, _1356_v35)
		var _1359_v38 _dafny.Array
		_ = _1359_v38
		var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw240
		_1359_v38 = _nw240
		var _1360_v39 T0
		_ = _1360_v39
		var _nw241 *C3 = New_C3_()
		_ = _nw241
		_nw241.Ctor__(_1314_v0, _this.F12(), _this.F11(), _this.F12())
		_1360_v39 = _nw241
		var _1361_v40 _dafny.Set
		_ = _1361_v40
		_1361_v40 = _dafny.SetOf(_1360_v39, _1360_v39)
		var _1362_v41 D8
		_ = _1362_v41
		_1362_v41 = Companion_D8_.Create_DC21_(true, _1359_v38, _dafny.IntOfUint32((_1315_v1).Cardinality()), _1316_v2, (_1361_v40).Cardinality())
		var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1357_v36), 0))
		_ = _index217
		(_1357_v36).ArraySet1CodePoint(Companion_Default___.Fm14(_1358_v37, (_dafny.Zero).Minus((_1362_v41).Dtor_cf53()), globalState), (_index217).Int())
		var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))
		_ = _index218
		var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1357_v36), 0))
		_ = _index219
		var _rhs242 _dafny.Int = _dafny.IntOfUint32((_1318_v4).Cardinality())
		_ = _rhs242
		var _rhs243 _dafny.Set = (p0).Intersection(_dafny.SetOf(true))
		_ = _rhs243
		var _rhs244 _dafny.CodePoint = _1356_v35
		_ = _rhs244
		var _rhs245 _dafny.Sequence = Companion_Default___.Fm24(globalState)
		_ = _rhs245
		var _lhs203 *GlobalState = globalState
		_ = _lhs203
		var _lhs204 _dafny.Array = _1321_v7
		_ = _lhs204
		var _lhs205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))
		_ = _lhs205
		var _lhs206 _dafny.Array = _1357_v36
		_ = _lhs206
		var _lhs207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1357_v36), 0))
		_ = _lhs207
		_lhs203.F9 = _rhs242
		(_lhs204).ArraySet1(_rhs243, (_lhs205).Int())
		(_lhs206).ArraySet1CodePoint(_rhs244, (_lhs207).Int())
		_1318_v4 = _rhs245
		var _hi3 _dafny.Int = _1314_v0
		_ = _hi3
		for _1363_i4 := Companion_Default___.SafeDivisionInt(_1314_v0, _1314_v0); _1363_i4.Cmp(_hi3) < 0; _1363_i4 = _1363_i4.Plus(_dafny.One) {
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(185), _dafny.ArrayLen((_1321_v7), 0))
			_ = _index220
			(_1321_v7).ArraySet1(p0, (_index220).Int())
			var _1364_v42 _dafny.Array
			_ = _1364_v42
			var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw242
			_1364_v42 = _nw242
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1364_v42), 0))
			_ = _index221
			(_1364_v42).ArraySet1(_1315_v1, (_index221).Int())
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1364_v42), 0))
			_ = _index222
			(_1364_v42).ArraySet1(_1315_v1, (_index222).Int())
			var _1365_v43 _dafny.Sequence
			_ = _1365_v43
			_1365_v43 = _dafny.SeqOf(_this.F12())
			var _1366_v44 _dafny.Map
			_ = _1366_v44
			_1366_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1360_v39.F12(), _this.F12())), _1365_v43)
			var _1367_v45 _dafny.Map
			_ = _1367_v45
			_1367_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1357_v36).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1357_v36), 0))).Int()), _1360_v39.F12())
			var _rhs246 _dafny.Sequence = (func() _dafny.Sequence {
				if (_1366_v44).Contains(_1360_v39.F11()) {
					return (_1366_v44).Get(_1360_v39.F11()).(_dafny.Sequence)
				}
				return _1365_v43
			})()
			_ = _rhs246
			var _rhs247 _dafny.Int = _1314_v0
			_ = _rhs247
			var _rhs248 bool = ((_1367_v45).Merge(_1367_v45)).Contains((_1357_v36).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1357_v36), 0))).Int()))
			_ = _rhs248
			var _lhs208 *GlobalState = globalState
			_ = _lhs208
			var _lhs209 *GlobalState = globalState
			_ = _lhs209
			_lhs208.F6 = _rhs246
			_1314_v0 = _rhs247
			_lhs209.F7 = _rhs248
			var _1368_v46 _dafny.Map
			_ = _1368_v46
			_1368_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1360_v39, _1360_v39)
			_1368_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1360_v39, _1360_v39)
		}
		var _1369_v47 _dafny.Array
		_ = _1369_v47
		var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
		_ = _nw243
		_1369_v47 = _nw243
		r0 = _1369_v47
		var _1370_v48 _dafny.Array
		_ = _1370_v48
		var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw244
		_1370_v48 = _nw244
		r1 = _1370_v48
		r2 = _1314_v0
		return r0, r1, r2
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
