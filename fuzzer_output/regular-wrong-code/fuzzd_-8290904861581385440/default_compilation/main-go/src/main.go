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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfInt64(856)).Minus((_dafny.IntOfInt64(964)).Plus(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("jxdfowin"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("imhcdljya"), _dafny.UnicodeSeqOfUtf8Bytes("gng"))))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(-386))
}
func (_static *CompanionStruct_Default___) Fm4(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('w')
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	return ((func() _dafny.Int {
		if true {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oydcocpi")).Cardinality())
		}
		return _dafny.IntOfInt64(904)
	})()).Cmp(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(-604))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.CodePoint
			_0_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(-604))).Contains(_0_v0) {
				_coll0.Add(_0_v0, !(false))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(820))).Cardinality()))) != 0
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Sequence, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(176))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	})), _dafny.SeqOf(_dafny.CodePoint('d')))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D0_.Create_DC2_(!(true), (_dafny.SetOf(_dafny.IntOfInt64(185))).Cardinality(), _dafny.SeqOf(true, false, true), _dafny.IntOfInt64(273), _dafny.IntOfUint32((_dafny.SeqOf(!(true), false)).Cardinality()))).Dtor_cf6()
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if true {
		return _dafny.MultiSetOf(_dafny.IntOfInt64(676), (func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(492), _dafny.IntOfInt64(328))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _2_v0 _dafny.Int
				_2_v0 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(492)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(328)) < 0) {
					_coll1.Add((_2_v0).Plus(_dafny.IntOfInt64(956)), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), _dafny.IntOfInt64(49))
	} else {
		return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(-771)), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('y')
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('i'))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(491)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(true))), _dafny.IntOfInt64(644))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, !(false), true, true)).Difference((_dafny.SetOf(!(true), !(true), false)).Intersection(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_D3_.Create_DC9_(_dafny.UnicodeSeqOfUtf8Bytes("hsnhlv")))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_3_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
		}))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_2).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_3_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())
			})), _4_v0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_4_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pkhneayht")).Cardinality())), _dafny.IntOfInt64(583))
			}
		}
		return _coll2.ToMap()
	}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-379), _dafny.IntOfInt64(887))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D5_.Create_DC17_((_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(-257), _dafny.IntOfInt64(187), _dafny.IntOfInt64(911), _dafny.IntOfInt64(21))).Cardinality(), Companion_D0_.Create_DC1_(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, false, false)).Cardinality()), false)).Cardinality(), _dafny.CodePoint('w'))).Cardinality())).Cardinality()), _dafny.IntOfInt64(729))).Cardinality()), _dafny.IntOfInt64(798)), false, true))).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, true))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true, true, true, true)))).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(false), false, false, false), _dafny.SeqOf(false, true))))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, globalState *GlobalState) D0 {
	if !((_dafny.IntOfInt64(192)).Cmp(_dafny.IntOfInt64(358)) == 0) {
		return Companion_D0_.Create_DC0_(_dafny.MultiSetOf(false, false))
	} else {
		return Companion_D0_.Create_DC0_(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false)))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(false, true)
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D5_.Create_DC17_((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()))), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("naexdwq")).Cardinality()), Companion_D0_.Create_DC1_(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(376))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(752), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("famhky")).Cardinality()))).Cardinality())).Cardinality()), true, true)).Dtor_cf33()))
	})(), (_dafny.IntOfInt64(889)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwit")).Cardinality()))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D2 {
	var _source0 D1 = Companion_D1_.Create_DC4_(_dafny.SeqOf(Companion_D0_.Create_DC2_(true, (_dafny.SetOf(true, false)).Cardinality(), _dafny.SeqOf(true), _dafny.IntOfInt64(467), _dafny.IntOfInt64(-669)), Companion_D0_.Create_DC2_(true, _dafny.IntOfInt64(195), _dafny.SeqOf(false), _dafny.IntOfInt64(-382), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqavgwus")).Cardinality()))))
	_ = _source0
	if _source0.Is_DC5() {
		var _5___mcc_h0 bool = _source0.Get_().(D1_DC5).Cf11
		_ = _5___mcc_h0
		var _6_cf11 bool = _5___mcc_h0
		_ = _6_cf11
		return Companion_D2_.Create_DC8_(_dafny.IntOfInt64(231), (_dafny.Zero).Minus(_dafny.IntOfInt64(-575)))
	} else if _source0.Is_DC4() {
		var _7___mcc_h1 _dafny.Sequence = _source0.Get_().(D1_DC4).Cf10
		_ = _7___mcc_h1
		var _8_cf10 _dafny.Sequence = _7___mcc_h1
		_ = _8_cf10
		return Companion_D2_.Create_DC8_(_dafny.IntOfInt64(527), _dafny.IntOfInt64(49))
	} else {
		var _9___mcc_h2 D1 = _source0.Get_().(D1_DC6).Cf12
		_ = _9___mcc_h2
		var _10_cf12 D1 = _9___mcc_h2
		_ = _10_cf12
		return Companion_D2_.Create_DC8_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("etcx")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("djn"), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vdomg")).Cardinality()))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC2_(true, (_dafny.MultiSetOf(false, true, true, true, !(false))).Cardinality(), _dafny.SeqOf(false), _dafny.IntOfInt64(74), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nrcsueoj")).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(762))).Uint32(), func(coer3 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_11_i0 _dafny.Int) D0 {
		return Companion_D0_.Create_DC2_(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-536), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jva")).Cardinality()))).Cardinality())), _dafny.SeqOf(false, false), _dafny.IntOfInt64(546), _dafny.IntOfInt64(-596))
	})))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, (Companion_D0_.Create_DC2_(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(541))).Cardinality()), _dafny.SeqOf(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(true)).Cardinality())).Cardinality(), _dafny.IntOfInt64(810))).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ats")).Cardinality()))).Dtor_cf4(), !(!(false)), true)).Cardinality()), false)).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(660), _dafny.IntOfInt64(885))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _12_v0 _dafny.Int
			_12_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(660)).Cmp(_12_v0) <= 0) && ((_12_v0).Cmp(_dafny.IntOfInt64(885)) < 0) {
				_coll3.Add((_12_v0).Minus(_dafny.IntOfInt64(641)), false)
			}
		}
		return _coll3.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Map, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(!(false) || (false))
}
func (_static *CompanionStruct_Default___) Fm27(p0 bool, p1 _dafny.Map, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((false) == (false), _dafny.IntOfInt64(192), (func() _dafny.Int {
		if false {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hnqykwbvy")).Cardinality())
		}
		return _dafny.IntOfInt64(443)
	})())
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.SetOf((Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(2), _dafny.IntOfInt64(-271))).Dtor_cf1(), false, false, true, true)).Intersection(_dafny.SetOf(!(true))), (_dafny.SetOf(true, false)).Union(_dafny.SetOf(false, true)), (_dafny.SetOf(false)).Difference(_dafny.SetOf(true, true)), (_dafny.SetOf(!(true))).Union(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _13_v0 _dafny.Int
	_ = _13_v0
	_13_v0 = _dafny.IntOfInt64(888)
	var _14_v1 _dafny.Sequence
	_ = _14_v1
	_14_v1 = _dafny.SeqOf(_13_v0, _13_v0)
	var _15_v2 _dafny.Sequence
	_ = _15_v2
	_15_v2 = _dafny.SeqOf(_13_v0, _dafny.IntOfUint32((_14_v1).Cardinality()))
	var _16_v3 _dafny.Array
	_ = _16_v3
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
	_ = _nw0
	_16_v3 = _nw0
	var _17_v4 _dafny.Set
	_ = _17_v4
	_17_v4 = _dafny.SetOf(_13_v0)
	var _18_v5 _dafny.Map
	_ = _18_v5
	_18_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_v4, _13_v0)
	var _19_v6 _dafny.Array
	_ = _19_v6
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_0
	var _nw1 _dafny.Array
	_ = _nw1
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw1 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = func(_20_i0 _dafny.Int) bool {
			return true
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
	_19_v6 = _nw1
	var _21_v7 _dafny.Sequence
	_ = _21_v7
	_21_v7 = _dafny.SeqOf(_16_v3)
	var _22_v8 _dafny.Array
	_ = _22_v8
	var _nwElement0_0 _dafny.Array = _16_v3
	_ = _nwElement0_0
	var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(3))
	_ = _nw2
	(_nw2).ArraySet1(_nwElement0_0, 0)
	(_nw2).ArraySet1(_16_v3, 1)
	(_nw2).ArraySet1((_21_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.IntOfUint32((_21_v7).Cardinality()))).Uint32()).(_dafny.Array), 2)
	_22_v8 = _nw2
	var _23_globalState *GlobalState
	_ = _23_globalState
	var _nw3 *GlobalState = New_GlobalState_()
	_ = _nw3
	_nw3.Ctor__(true, _15_v2, _dafny.IntOfInt64(957), _16_v3, _18_v5, _dafny.IntOfInt64(116), _dafny.IntOfInt64(-467), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v6, _13_v0), true, _22_v8, _16_v3, true, false, _dafny.IntOfInt64(-352), _16_v3, _dafny.IntOfInt64(833), false, true, _dafny.IntOfInt64(294))
	_23_globalState = _nw3
	var _24_v9 bool
	_ = _24_v9
	_24_v9 = true
	var _hi0 _dafny.Int = _dafny.IntOfInt64(379)
	_ = _hi0
	for _25_i1 := Companion_Default___.Fm0(_24_v9, _24_v9, _23_globalState); _25_i1.Cmp(_hi0) < 0; _25_i1 = _25_i1.Plus(_dafny.One) {
		(_23_globalState).F18 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(_24_v9, _24_v9, _23_globalState)))
		var _26_v10 *C0
		_ = _26_v10
		var _nw4 *C0 = New_C0_()
		_ = _nw4
		_nw4.Ctor__(_13_v0)
		_26_v10 = _nw4
		var _27_v11 *C2
		_ = _27_v11
		var _nw5 *C2 = New_C2_()
		_ = _nw5
		_nw5.Ctor__(_26_v10, _24_v9, _19_v6, _24_v9)
		_27_v11 = _nw5
		var _28_v12 _dafny.Map
		_ = _28_v12
		_28_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_27_v11.F26), _24_v9)
		_24_v9 = !((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v11.F26, Companion_Default___.Fm9(_24_v9, _23_globalState))).Equals((_28_v12).Merge(_28_v12)))
		(_23_globalState).F13 = (_dafny.Zero).Minus(_25_i1)
	}
	(_23_globalState).F0 = Companion_Default___.Fm7(_13_v0, _13_v0, _13_v0, _23_globalState)
	var _29_v13 _dafny.CodePoint
	_ = _29_v13
	_29_v13 = _dafny.CodePoint('j')
	var _30_v14 _dafny.Sequence
	_ = _30_v14
	_30_v14 = _dafny.SeqOf(_29_v13)
	var _31_v15 *C3
	_ = _31_v15
	var _nw6 *C3 = New_C3_()
	_ = _nw6
	_nw6.Ctor__(Companion_Default___.Fm8(_30_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _24_v9), _13_v0, _23_globalState), _19_v6)
	_31_v15 = _nw6
	var _32_v16 _dafny.Sequence
	_ = _32_v16
	_32_v16 = _dafny.SeqOf(_31_v15, _31_v15, _31_v15, _31_v15)
	var _hi1 _dafny.Int = (_17_v4).Cardinality()
	_ = _hi1
	for _33_i2 := (_13_v0).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_32_v16, (Companion_Default___.SafeIndex(_13_v0, _dafny.IntOfUint32((_32_v16).Cardinality()))).Uint32(), _31_v15)).Cardinality())); _33_i2.Cmp(_hi1) < 0; _33_i2 = _33_i2.Plus(_dafny.One) {
		(_23_globalState).F0 = !(false) || (_24_v9)
		(_31_v15).F22 = _31_v15.F22
		var _34_v17 _dafny.Map
		_ = _34_v17
		_34_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v9, _24_v9)
		_34_v17 = _34_v17
		(_23_globalState).F0 = !((_13_v0).Cmp(_13_v0) != 0)
	}
	var _35_i3 _dafny.Int
	_ = _35_i3
	_35_i3 = _dafny.Zero
	{
		for Companion_Default___.Fm2(_24_v9, _13_v0, _13_v0, _23_globalState) {
			{
				if (_35_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_35_i3 = (_35_i3).Plus(_dafny.One)
				var _36_v18 _dafny.Array
				_ = _36_v18
				var _len0_1 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_1
				var _nw7 _dafny.Array
				_ = _nw7
				if _len0_1.Cmp(_dafny.Zero) == 0 {
					_nw7 = _dafny.NewArray(_len0_1)
				} else {
					var _init1 func(_dafny.Int) _dafny.Sequence = (func(_37_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_38_i4 _dafny.Int) _dafny.Sequence {
							return _37_v1
						}
					})(_14_v1)
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
				_36_v18 = _nw7
				var _39_v19 D4
				_ = _39_v19
				_39_v19 = Companion_D4_.Create_DC12_(_29_v13)
				var _40_v20 _dafny.Map
				_ = _40_v20
				_40_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_v18, _39_v19)
				_40_v20 = (_40_v20).Update(_36_v18, _39_v19)
				var _41_v21 _dafny.Sequence
				_ = _41_v21
				_41_v21 = _dafny.SeqOf(_24_v9)
				var _42_v22 T0
				_ = _42_v22
				var _nw8 *C1 = New_C1_()
				_ = _nw8
				_nw8.Ctor__(_31_v15.F22, (_31_v15).F23(), _24_v9)
				_42_v22 = _nw8
				var _43_v23 _dafny.Map
				_ = _43_v23
				_43_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_v22, _13_v0)
				var _44_v24 *C4
				_ = _44_v24
				var _nw9 *C4 = New_C4_()
				_ = _nw9
				_nw9.Ctor__((_31_v15).Fm5(Companion_Default___.Fm7(_dafny.IntOfUint32((_41_v21).Cardinality()), (func() _dafny.Int {
					if (_43_v23).Contains(_42_v22) {
						return (_43_v23).Get(_42_v22).(_dafny.Int)
					}
					return _13_v0
				})(), _dafny.IntOfUint32((_30_v14).Cardinality()), _23_globalState), _23_globalState), _19_v6, (_42_v22).F20())
				_44_v24 = _nw9
				_44_v24 = _44_v24
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_31_v15).F23()), 0))
				_ = _index0
				((_31_v15).F23()).ArraySet1((_24_v9) == ((_42_v22).F20()), (_index0).Int())
				var _45_v25 _dafny.Map
				_ = _45_v25
				_45_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _24_v9)
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen(((_31_v15).F23()), 0))
				_ = _index1
				((_31_v15).F23()).ArraySet1(!_dafny.Companion_Sequence_.Equal(_30_v14, Companion_Default___.Fm8(_30_v14, _45_v25, (_44_v24).F21(), _23_globalState)), (_index1).Int())
				var _46_v26 _dafny.Array
				_ = _46_v26
				var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw10
				_46_v26 = _nw10
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_46_v26), 0))
				_ = _index2
				(_46_v26).ArraySet1(_31_v15.F22, (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_46_v26), 0))
				_ = _index3
				(_46_v26).ArraySet1(_31_v15.F22, (_index3).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _47_v27 _dafny.Sequence
	_ = _47_v27
	var _48_v28 bool
	_ = _48_v28
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 bool
	_ = _out1
	_out0, _out1 = (_31_v15).M2(_23_globalState)
	_47_v27 = _out0
	_48_v28 = _out1
	var _49_i5 _dafny.Int
	_ = _49_i5
	_49_i5 = _dafny.Zero
	{
		for _24_v9 {
			{
				if (_49_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_49_i5 = (_49_i5).Plus(_dafny.One)
				(_23_globalState).F0 = (Companion_Default___.SafeModuloInt(_13_v0, _dafny.IntOfInt64(155))).Cmp(_dafny.IntOfInt64(845)) <= 0
				var _50_v29 _dafny.Map
				_ = _50_v29
				_50_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v28, _48_v28)
				(_23_globalState).F18 = Companion_Default___.SafeDivisionInt(_13_v0, (_31_v15).Fm5((func() bool {
					if (_50_v29).Contains(_48_v28) {
						return (_50_v29).Get(_48_v28).(bool)
					}
					return _24_v9
				})(), _23_globalState))
				if !_dafny.Companion_Sequence_.Contains(_14_v1, _13_v0) {
					(_23_globalState).F13 = _13_v0
					var _51_v30 T0
					_ = _51_v30
					var _nw11 *C1 = New_C1_()
					_ = _nw11
					_nw11.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(963))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg4 _dafny.Int) interface{} {
							return coer4(arg4)
						}
					}((func(_52_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_53_i6 _dafny.Int) _dafny.CodePoint {
							return _52_v13
						}
					})(_29_v13))), (_31_v15).F23(), _24_v9)
					_51_v30 = _nw11
					var _54_v31 _dafny.Sequence
					_ = _54_v31
					_54_v31 = _dafny.SeqOf(_51_v30, _51_v30)
					var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_51_v30, _51_v30, _51_v30), _54_v31)
					_ = _rhs0
					var _rhs1 _dafny.Int = (_dafny.Zero).Minus(_13_v0)
					_ = _rhs1
					var _lhs0 *GlobalState = _23_globalState
					_ = _lhs0
					_54_v31 = _rhs0
					_lhs0.F18 = _rhs1
					var _55_v32 _dafny.Sequence
					_ = _55_v32
					var _56_v33 bool
					_ = _56_v33
					var _out2 _dafny.Sequence
					_ = _out2
					var _out3 bool
					_ = _out3
					_out2, _out3 = (_31_v15).M2(_23_globalState)
					_55_v32 = _out2
					_56_v33 = _out3
					(_23_globalState).F18 = (Companion_Default___.Fm0(_24_v9, _24_v9, _23_globalState)).Times(Companion_Default___.SafeModuloInt(_13_v0, _13_v0))
					_30_v14 = (func() _dafny.Sequence {
						if (_13_v0).Cmp(_13_v0) <= 0 {
							return _30_v14
						}
						return _31_v15.F22
					})()
				} else {
					var _57_v34 *C1
					_ = _57_v34
					var _nw12 *C1 = New_C1_()
					_ = _nw12
					_nw12.Ctor__(_dafny.Companion_Sequence_.Concatenate(_30_v14, Companion_Default___.Fm8(_dafny.SeqOf(Companion_Default___.Fm4(_23_globalState)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _24_v9), _dafny.IntOfUint32((_30_v14).Cardinality()), _23_globalState)), (_31_v15).F23(), _24_v9)
					_57_v34 = _nw12
					_48_v28 = _48_v28
					var _58_v35 *C1
					_ = _58_v35
					var _nw13 *C1 = New_C1_()
					_ = _nw13
					_nw13.Ctor__(_30_v14, (_31_v15).F23(), _24_v9)
					_58_v35 = _nw13
					var _59_v36 *C0
					_ = _59_v36
					var _nw14 *C0 = New_C0_()
					_ = _nw14
					_nw14.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_47_v27).Cardinality()), _13_v0)).Cardinality()))
					_59_v36 = _nw14
					var _60_v37 *C2
					_ = _60_v37
					var _nw15 *C2 = New_C2_()
					_ = _nw15
					_nw15.Ctor__(_59_v36, _24_v9, (_31_v15).F23(), _48_v28)
					_60_v37 = _nw15
					var _61_v38 D8
					_ = _61_v38
					_61_v38 = Companion_D8_.Create_DC22_(_60_v37)
					var _62_v39 _dafny.Set
					_ = _62_v39
					_62_v39 = _dafny.SetOf((_61_v38).Dtor_cf42(), _60_v37, _60_v37, _60_v37, _60_v37)
					_62_v39 = (_62_v39).Intersection(_62_v39)
					var _63_v40 bool
					_ = _63_v40
					var _64_v41 _dafny.Int
					_ = _64_v41
					var _65_v42 _dafny.Int
					_ = _65_v42
					var _66_v43 _dafny.Int
					_ = _66_v43
					var _out4 bool
					_ = _out4
					var _out5 _dafny.Int
					_ = _out5
					var _out6 _dafny.Int
					_ = _out6
					var _out7 _dafny.Int
					_ = _out7
					_out4, _out5, _out6, _out7 = (_60_v37).M4(((_59_v36).F24()).Cmp((_59_v36).F24()) < 0, _23_globalState)
					_63_v40 = _out4
					_64_v41 = _out5
					_65_v42 = _out6
					_66_v43 = _out7
				}
				(_31_v15).M3(_23_globalState)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _67_v44 _dafny.Array
	_ = _67_v44
	var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.One)
	_ = _nw16
	_67_v44 = _nw16
	var _68_v45 _dafny.MultiSet
	_ = _68_v45
	_68_v45 = _dafny.MultiSetOf(_13_v0, _dafny.IntOfInt64(-409))
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_67_v44), 0))
	_ = _index4
	(_67_v44).ArraySet1(_68_v45, (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_67_v44), 0))
	_ = _index5
	(_67_v44).ArraySet1((_68_v45).Intersection((_68_v45).Difference(_68_v45)), (_index5).Int())
	var _69_v46 _dafny.Map
	_ = _69_v46
	_69_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v9, _13_v0)
	var _70_v47 D1
	_ = _70_v47
	_70_v47 = Companion_D1_.Create_DC6_(Companion_Default___.Fm26(_13_v0, _24_v9, _24_v9, _69_v46, _23_globalState))
	var _71_v48 D1
	_ = _71_v48
	_71_v48 = Companion_D1_.Create_DC6_(_70_v47)
	var _pat_let_tv0 = _17_v4
	_ = _pat_let_tv0
	var _pat_let_tv1 = _17_v4
	_ = _pat_let_tv1
	var _pat_let_tv2 = _17_v4
	_ = _pat_let_tv2
	var _pat_let_tv3 = _67_v44
	_ = _pat_let_tv3
	var _pat_let_tv4 = _67_v44
	_ = _pat_let_tv4
	var _pat_let_tv5 = _67_v44
	_ = _pat_let_tv5
	var _pat_let_tv6 = _67_v44
	_ = _pat_let_tv6
	_17_v4 = func(_source1 D1) _dafny.Set {
		if _source1.Is_DC5() {
			var _72___mcc_h0 bool = _source1.Get_().(D1_DC5).Cf11
			_ = _72___mcc_h0
			var _73_cf11 bool = _72___mcc_h0
			_ = _73_cf11
			return _pat_let_tv0
		} else if _source1.Is_DC4() {
			var _74___mcc_h1 _dafny.Sequence = _source1.Get_().(D1_DC4).Cf10
			_ = _74___mcc_h1
			var _75_cf10 _dafny.Sequence = _74___mcc_h1
			_ = _75_cf10
			return _pat_let_tv1
		} else {
			var _76___mcc_h2 D1 = _source1.Get_().(D1_DC6).Cf12
			_ = _76___mcc_h2
			var _77_cf12 D1 = _76___mcc_h2
			_ = _77_cf12
			return (_pat_let_tv2).Union(func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(((_pat_let_tv4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_pat_let_tv3), 0))).Int()).(_dafny.MultiSet)).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _78_v49 _dafny.Int
					_78_v49 = interface{}(_compr_4).(_dafny.Int)
					if ((_pat_let_tv6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_pat_let_tv5), 0))).Int()).(_dafny.MultiSet)).Contains(_78_v49) {
						_coll4.Add((_78_v49).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(258))).Cardinality())))
					}
				}
				return _coll4.ToSet()
			}())
		}
	}(_71_v48)
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
	_ = _index6
	((_31_v15).F23()).ArraySet1(false, (_index6).Int())
	var _79_v50 *C0
	_ = _79_v50
	var _nw17 *C0 = New_C0_()
	_ = _nw17
	_nw17.Ctor__(_13_v0)
	_79_v50 = _nw17
	var _80_v51 T0
	_ = _80_v51
	var _nw18 *C2 = New_C2_()
	_ = _nw18
	_nw18.Ctor__(_79_v50, false, (_31_v15).F23(), false)
	_80_v51 = _nw18
	var _81_v52 _dafny.Map
	_ = _81_v52
	_81_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v51, Companion_Default___.Fm7((_dafny.Zero).Minus(_dafny.IntOfUint32((_30_v14).Cardinality())), (_dafny.Zero).Minus((_79_v50).F24()), (_79_v50).F24(), _23_globalState))
	var _82_v53 _dafny.MultiSet
	_ = _82_v53
	_82_v53 = _dafny.MultiSetOf(_29_v13, _dafny.CodePoint('o'), Companion_Default___.Fm15((_79_v50).F24(), (_80_v51).F20(), (_79_v50).F24(), _23_globalState))
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
	_ = _index7
	((_31_v15).F23()).ArraySet1(((_81_v52).Cardinality()).Cmp((_82_v53).Cardinality()) != 0, (_index7).Int())
	var _source2 D0 = Companion_Default___.Fm27(!(_48_v28), func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(712), _dafny.IntOfInt64(713))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _83_v54 _dafny.Int
			_83_v54 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(712)).Cmp(_83_v54) <= 0) && ((_83_v54).Cmp(_dafny.IntOfInt64(713)) < 0) {
				_coll5.Add(Companion_Default___.SafeDivisionInt(_83_v54, (_dafny.Zero).Minus((_dafny.Zero).Minus(_13_v0))), Companion_D3_.Create_DC11_(_13_v0, _13_v0, _48_v28))
			}
		}
		return _coll5.ToMap()
	}(), _23_globalState)
	_ = _source2
	if _source2.Is_DC1() {
		var _84___mcc_h3 bool = _source2.Get_().(D0_DC1).Cf1
		_ = _84___mcc_h3
		var _85___mcc_h4 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
		_ = _85___mcc_h4
		var _86___mcc_h5 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
		_ = _86___mcc_h5
		var _87_cf3 _dafny.Int = _86___mcc_h5
		_ = _87_cf3
		var _88_cf2 _dafny.Int = _85___mcc_h4
		_ = _88_cf2
		var _89_cf1 bool = _84___mcc_h3
		_ = _89_cf1
		_14_v1 = _dafny.Companion_Sequence_.Concatenate(_15_v2, _dafny.Companion_Sequence_.Concatenate(_15_v2, _dafny.SeqOf(_dafny.IntOfInt64(648), (func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(828), _dafny.IntOfInt64(775))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _90_v55 _dafny.Int
				_90_v55 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(828)).Cmp(_90_v55) <= 0) && ((_90_v55).Cmp(_dafny.IntOfInt64(775)) < 0) {
					_coll6.Add((_90_v55).Minus((_79_v50).F24()), _dafny.SetOf(_29_v13))
				}
			}
			return _coll6.ToMap()
		}()).Cardinality())))
		_87_cf3 = _88_cf2
		var _91_v56 *C2
		_ = _91_v56
		var _nw19 *C2 = New_C2_()
		_ = _nw19
		_nw19.Ctor__(_79_v50, (_80_v51).F20(), (_31_v15).F23(), ((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool))
		_91_v56 = _nw19
		if (_91_v56.F26) == ((_80_v51).F20()) {
			var _92_v57 _dafny.MultiSet
			_ = _92_v57
			_92_v57 = _dafny.MultiSetOf((_80_v51).F20())
			_88_cf2 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_92_v57).Update(_91_v56.F26, Companion_Default___.Abs((_14_v1).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_cf1, true)).Cardinality(), _dafny.IntOfUint32((_14_v1).Cardinality()))).Uint32()).(_dafny.Int)))).Cardinality(), (_79_v50).F24()))
			var _93_v58 _dafny.Map
			_ = _93_v58
			_93_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _dafny.UnicodeSeqOfUtf8Bytes("gvadfbvx"))
			var _94_v59 _dafny.Map
			_ = _94_v59
			_94_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v58, (_dafny.IntOfUint32((_31_v15.F22).Cardinality())).Minus((_79_v50).F24()))
			_94_v59 = (_94_v59).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_79_v50).F24(), _31_v15.F22), (_88_cf2).Plus((_79_v50).F24()))
			var _95_v60 _dafny.Array
			_ = _95_v60
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_2
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Set = (func(_96_v4 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_97_i7 _dafny.Int) _dafny.Set {
						return _96_v4
					}
				})(_17_v4)
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
			_95_v60 = _nw20
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_95_v60), 0))
			_ = _index8
			(_95_v60).ArraySet1(_17_v4, (_index8).Int())
			var _98_v61 _dafny.Set
			_ = _98_v61
			_98_v61 = _dafny.SetOf(_13_v0)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
			_ = _index9
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_95_v60), 0))
			_ = _index10
			var _rhs2 _dafny.Int = Companion_Default___.SafeModuloInt((_79_v50).F24(), _13_v0)
			_ = _rhs2
			var _rhs3 _dafny.Int = (_79_v50).F24()
			_ = _rhs3
			var _rhs4 bool = (_80_v51).F20()
			_ = _rhs4
			var _rhs5 _dafny.Set = (_98_v61)
			_ = _rhs5
			var _lhs1 *GlobalState = _23_globalState
			_ = _lhs1
			var _lhs2 _dafny.Array = (_31_v15).F23()
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
			_ = _lhs3
			var _lhs4 _dafny.Array = _95_v60
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_95_v60), 0))
			_ = _lhs5
			_87_cf3 = _rhs2
			_lhs1.F13 = _rhs3
			(_lhs2).ArraySet1(_rhs4, (_lhs3).Int())
			(_lhs4).ArraySet1(_rhs5, (_lhs5).Int())
			var _99_v62 _dafny.Sequence
			_ = _99_v62
			_99_v62 = _dafny.SeqOf((_31_v15).F23())
			var _100_v63 *C3
			_ = _100_v63
			var _nw21 *C3 = New_C3_()
			_ = _nw21
			_nw21.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("kfcjm"), (_99_v62).Select((Companion_Default___.SafeIndex(_13_v0, _dafny.IntOfUint32((_99_v62).Cardinality()))).Uint32()).(_dafny.Array))
			_100_v63 = _nw21
			var _101_v64 *C1
			_ = _101_v64
			var _nw22 *C1 = New_C1_()
			_ = _nw22
			_nw22.Ctor__(_31_v15.F22, (_31_v15).F23(), true)
			_101_v64 = _nw22
			var _102_v65 _dafny.Sequence
			_ = _102_v65
			_102_v65 = _dafny.SeqOf(_dafny.SeqOf(_101_v64, _101_v64))
			_17_v4 = ((_95_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_95_v60), 0))).Int()).(_dafny.Set)).Difference(_dafny.SetOf(_dafny.IntOfUint32(((_102_v65).Select((Companion_Default___.SafeIndex((_79_v50).F24(), _dafny.IntOfUint32((_102_v65).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
		} else {
			var _103_v66 _dafny.Array
			_ = _103_v66
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_3
			var _nw23 _dafny.Array
			_ = _nw23
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw23 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_104_v15 *C3) func(_dafny.Int) _dafny.Sequence {
					return func(_105_i8 _dafny.Int) _dafny.Sequence {
						return _104_v15.F22
					}
				})(_31_v15)
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
			_103_v66 = _nw23
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_103_v66), 0))
			_ = _index11
			(_103_v66).ArraySet1(_30_v14, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_103_v66), 0))
			_ = _index12
			(_103_v66).ArraySet1(_31_v15.F22, (_index12).Int())
			var _106_v67 _dafny.Set
			_ = _106_v67
			_106_v67 = _dafny.SetOf(_19_v6)
			_106_v67 = (_106_v67).Intersection(_106_v67)
			_24_v9 = _24_v9
			(_91_v56).F26 = !(!(_dafny.SetOf((_80_v51).F20(), _89_cf1, _91_v56.F26)).Contains(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_88_cf2, (_79_v50).F24()), _14_v1))))
			var _107_v68 _dafny.Array
			_ = _107_v68
			var _len0_4 _dafny.Int = _dafny.One
			_ = _len0_4
			var _nw24 _dafny.Array
			_ = _nw24
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw24 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_108_v27 _dafny.Sequence, _109_v50 *C0) func(_dafny.Int) _dafny.Sequence {
					return func(_110_i9 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update(_108_v27, (Companion_Default___.SafeIndex((_109_v50).F24(), _dafny.IntOfUint32((_108_v27).Cardinality()))).Uint32(), true)
					}
				})(_47_v27, _79_v50)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw24 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw24).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw24).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_107_v68 = _nw24
			_107_v68 = _107_v68
		}
	} else if _source2.Is_DC2() {
		var _111___mcc_h6 bool = _source2.Get_().(D0_DC2).Cf4
		_ = _111___mcc_h6
		var _112___mcc_h7 _dafny.Int = _source2.Get_().(D0_DC2).Cf5
		_ = _112___mcc_h7
		var _113___mcc_h8 _dafny.Sequence = _source2.Get_().(D0_DC2).Cf6
		_ = _113___mcc_h8
		var _114___mcc_h9 _dafny.Int = _source2.Get_().(D0_DC2).Cf7
		_ = _114___mcc_h9
		var _115___mcc_h10 _dafny.Int = _source2.Get_().(D0_DC2).Cf8
		_ = _115___mcc_h10
		var _116_cf8 _dafny.Int = _115___mcc_h10
		_ = _116_cf8
		var _117_cf7 _dafny.Int = _114___mcc_h9
		_ = _117_cf7
		var _118_cf6 _dafny.Sequence = _113___mcc_h8
		_ = _118_cf6
		var _119_cf5 _dafny.Int = _112___mcc_h7
		_ = _119_cf5
		var _120_cf4 bool = _111___mcc_h6
		_ = _120_cf4
		(_23_globalState).F0 = _120_cf4
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_16_v3), 0))
		_ = _index13
		(_16_v3).ArraySet1((_79_v50).F24(), (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_16_v3), 0))
		_ = _index14
		(_16_v3).ArraySet1((_79_v50).F24(), (_index14).Int())
		(_23_globalState).F18 = (_79_v50).F24()
		(_31_v15).F22 = _31_v15.F22
	} else if _source2.Is_DC3() {
		var _121___mcc_h11 _dafny.Map = _source2.Get_().(D0_DC3).Cf9
		_ = _121___mcc_h11
		var _122_cf9 _dafny.Map = _121___mcc_h11
		_ = _122_cf9
		_19_v6 = (_31_v15).F23()
		var _123_v69 _dafny.Set
		_ = _123_v69
		_123_v69 = _dafny.SetOf((_80_v51).F20())
		var _124_v70 _dafny.Sequence
		_ = _124_v70
		_124_v70 = _dafny.SeqOf(_123_v69, (_123_v69).Difference(_123_v69), (_123_v69).Intersection(_123_v69), (Companion_Default___.Fm17(_dafny.IntOfUint32((_31_v15.F22).Cardinality()), _23_globalState)).Intersection(_123_v69))
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
		_ = _index15
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
		_ = _index16
		var _rhs6 _dafny.Sequence = _124_v70
		_ = _rhs6
		var _rhs7 bool = (_80_v51).F20()
		_ = _rhs7
		var _rhs8 bool = !(_24_v9)
		_ = _rhs8
		var _rhs9 bool = true
		_ = _rhs9
		var _rhs10 bool = ((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool)
		_ = _rhs10
		var _lhs6 *GlobalState = _23_globalState
		_ = _lhs6
		var _lhs7 _dafny.Array = (_31_v15).F23()
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
		_ = _lhs8
		var _lhs9 _dafny.Array = (_31_v15).F23()
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
		_ = _lhs10
		_124_v70 = _rhs6
		_lhs6.F0 = _rhs7
		(_lhs7).ArraySet1(_rhs8, (_lhs8).Int())
		(_lhs9).ArraySet1(_rhs9, (_lhs10).Int())
		_24_v9 = _rhs10
		var _125_v71 *C0
		_ = _125_v71
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__((_13_v0).Minus(_13_v0))
		_125_v71 = _nw25
		(_31_v15).M3(_23_globalState)
	} else {
		var _126___mcc_h12 _dafny.MultiSet = _source2.Get_().(D0_DC0).Cf0
		_ = _126___mcc_h12
		var _127_cf0 _dafny.MultiSet = _126___mcc_h12
		_ = _127_cf0
		var _128_v72 _dafny.Map
		_ = _128_v72
		_128_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((Companion_Default___.Fm19(_24_v9, _23_globalState)).Cardinality(), (_79_v50).F24()), false)
		_128_v72 = (_128_v72).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_30_v14, (Companion_Default___.SafeIndex((_31_v15).Fm5((_80_v51).F20(), _23_globalState), _dafny.IntOfUint32((_30_v14).Cardinality()))).Uint32(), _29_v13)).Cardinality()), !(_48_v28))
		var _129_v73 _dafny.Set
		_ = _129_v73
		_129_v73 = _dafny.SetOf(!((_80_v51).F20()), _24_v9, _48_v28, _24_v9)
		var _130_v74 D4
		_ = _130_v74
		_130_v74 = Companion_D4_.Create_DC14_(_24_v9, _129_v73)
		_130_v74 = _130_v74
		var _131_v75 D1
		_ = _131_v75
		_131_v75 = Companion_D1_.Create_DC5_((_80_v51).F20())
		(_23_globalState).F0 = Companion_Default___.Fm2((_131_v75).Dtor_cf11(), (func() _dafny.Int {
			if (_80_v51).F20() {
				return _13_v0
			}
			return (_79_v50).F24()
		})(), ((_79_v50).F24()).Times((_79_v50).F24()), _23_globalState)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_16_v3), 0))
		_ = _index17
		(_16_v3).ArraySet1((_79_v50).F24(), (_index17).Int())
		var _132_v76 _dafny.Set
		_ = _132_v76
		_132_v76 = _dafny.SetOf(_31_v15)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_16_v3), 0))
		_ = _index18
		(_16_v3).ArraySet1((((func() _dafny.Int {
			if (_127_cf0).Contains(_24_v9) {
				return (_127_cf0).Multiplicity(_24_v9)
			}
			return (_17_v4).Cardinality()
		})()).Plus((_132_v76).Cardinality())).Plus((_dafny.Zero).Minus(_13_v0)), (_index18).Int())
	}
	_15_v2 = _14_v1
	(_31_v15).M3(_23_globalState)
	var _133_v77 _dafny.Map
	_ = _133_v77
	_133_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v15.F22, _80_v51), ((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool))
	var _134_v78 _dafny.Map
	_ = _134_v78
	_134_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_80_v51).F20(), _80_v51)
	var _135_v79 _dafny.Map
	_ = _135_v79
	_135_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("sqqlffck"), (func() T0 {
		if (_134_v78).Contains((_80_v51).F20()) {
			return (_134_v78).Get((_80_v51).F20()).(T0)
		}
		return _80_v51
	})())
	var _136_v80 *C4
	_ = _136_v80
	var _nw26 *C4 = New_C4_()
	_ = _nw26
	_nw26.Ctor__(Companion_Default___.SafeModuloInt(_13_v0, _13_v0), _19_v6, (func() bool {
		if (_133_v77).Contains(_135_v79) {
			return (_133_v77).Get(_135_v79).(bool)
		}
		return _48_v28
	})())
	_136_v80 = _nw26
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
	_ = _index19
	var _rhs11 _dafny.CodePoint = _29_v13
	_ = _rhs11
	var _rhs12 _dafny.Int = _13_v0
	_ = _rhs12
	var _rhs13 *C4 = _136_v80
	_ = _rhs13
	var _rhs14 bool = _48_v28
	_ = _rhs14
	var _lhs11 *GlobalState = _23_globalState
	_ = _lhs11
	var _lhs12 _dafny.Array = (_31_v15).F23()
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
	_ = _lhs13
	_29_v13 = _rhs11
	_lhs11.F13 = _rhs12
	_136_v80 = _rhs13
	(_lhs12).ArraySet1(_rhs14, (_lhs13).Int())
	_16_v3 = _16_v3
	var _hi2 _dafny.Int = _13_v0
	_ = _hi2
	for _137_i10 := _13_v0; _137_i10.Cmp(_hi2) < 0; _137_i10 = _137_i10.Plus(_dafny.One) {
		var _138_v81 *C2
		_ = _138_v81
		var _nw27 *C2 = New_C2_()
		_ = _nw27
		_nw27.Ctor__(_79_v50, (_80_v51).F20(), (_80_v51).F19(), true)
		_138_v81 = _nw27
		var _139_v82 _dafny.Sequence
		_ = _139_v82
		_139_v82 = _dafny.SeqOf(_138_v81, _138_v81)
		var _140_v83 _dafny.Map
		_ = _140_v83
		_140_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, _139_v82)
		var _141_v84 _dafny.Sequence
		_ = _141_v84
		_141_v84 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_139_v82, _139_v82), _139_v82, (func() _dafny.Sequence {
			if (_140_v83).Contains(_13_v0) {
				return (_140_v83).Get(_13_v0).(_dafny.Sequence)
			}
			return _139_v82
		})())
		_141_v84 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_138_v81), _139_v82, _139_v82, _139_v82, (_141_v84).Select((Companion_Default___.SafeIndex(_13_v0, _dafny.IntOfUint32((_141_v84).Cardinality()))).Uint32()).(_dafny.Sequence)), _dafny.Companion_Sequence_.Concatenate(_141_v84, _141_v84))
		var _142_v85 D0
		_ = _142_v85
		_142_v85 = Companion_D0_.Create_DC2_(_138_v81.F26, (_dafny.SetOf(_31_v15)).Cardinality(), _dafny.SeqOf(((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool)), (_136_v80).F21(), _dafny.IntOfInt64(945))
		var _143_v86 D0
		_ = _143_v86
		_143_v86 = Companion_D0_.Create_DC2_((_142_v85).Dtor_cf4(), _137_i10, _dafny.Companion_Sequence_.Update(_47_v27, (Companion_Default___.SafeIndex(_137_i10, _dafny.IntOfUint32((_47_v27).Cardinality()))).Uint32(), _24_v9), (_136_v80).F21(), (_79_v50).F24())
		var _144_v87 _dafny.Map
		_ = _144_v87
		_144_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v15, _138_v81)
		var _145_v88 _dafny.Sequence
		_ = _145_v88
		_145_v88 = _dafny.SeqOf(_142_v85, _142_v85, _142_v85, Companion_D0_.Create_DC2_(((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool), _137_i10, _47_v27, (_dafny.Zero).Minus(_dafny.IntOfUint32((_30_v14).Cardinality())), _137_i10), Companion_D0_.Create_DC2_(_138_v81.F26, _137_i10, _dafny.SeqOf(false), (_144_v87).Cardinality(), _dafny.IntOfUint32((_15_v2).Cardinality())))
		var _source3 D1 = Companion_D1_.Create_DC4_(_145_v88)
		_ = _source3
		if _source3.Is_DC5() {
			var _146___mcc_h13 bool = _source3.Get_().(D1_DC5).Cf11
			_ = _146___mcc_h13
			var _147_cf11 bool = _146___mcc_h13
			_ = _147_cf11
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw28
			_22_v8 = _nw28
			(_23_globalState).F13 = (func() _dafny.Int {
				if (_80_v51).F20() {
					return (_31_v15).Fm6((_80_v51).F20(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_47_v27).Cardinality()), _dafny.IntOfInt64(40))).Update((_136_v80).F21(), (_136_v80).F21()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg5 _dafny.Int) interface{} {
							return coer5(arg5)
						}
					}((func(_148_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_149_i11 _dafny.Int) _dafny.CodePoint {
							return _148_v13
						}
					})(_29_v13))), false, _23_globalState)
				}
				return (_79_v50).F24()
			})()
			var _150_v89 *C1
			_ = _150_v89
			var _nw29 *C1 = New_C1_()
			_ = _nw29
			_nw29.Ctor__(_31_v15.F22, _19_v6, ((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool))
			_150_v89 = _nw29
			var _151_v90 _dafny.Map
			_ = _151_v90
			_151_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D8_.Create_DC23_(_48_v28, _dafny.IntOfUint32((_15_v2).Cardinality()), (_136_v80).F21(), _24_v9)).Dtor_cf45(), _150_v89)
			(_23_globalState).F18 = (((_79_v50).F24()).Times((_151_v90).Cardinality())).Times(_13_v0)
			_13_v0 = _dafny.IntOfInt64(75)
		} else if _source3.Is_DC4() {
			var _152___mcc_h14 _dafny.Sequence = _source3.Get_().(D1_DC4).Cf10
			_ = _152___mcc_h14
			var _153_cf10 _dafny.Sequence = _152___mcc_h14
			_ = _153_cf10
			var _154_v91 _dafny.Sequence
			_ = _154_v91
			_154_v91 = _dafny.SeqOf(_136_v80)
			var _155_v92 _dafny.Array
			_ = _155_v92
			var _nwElement0_1 *C4 = _136_v80
			_ = _nwElement0_1
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_1, 0)
			(_nw30).ArraySet1(_136_v80, 1)
			(_nw30).ArraySet1(_136_v80, 2)
			(_nw30).ArraySet1(_136_v80, 3)
			(_nw30).ArraySet1(_136_v80, 4)
			(_nw30).ArraySet1(_136_v80, 5)
			(_nw30).ArraySet1((_154_v91).Select((Companion_Default___.SafeIndex((_79_v50).F24(), _dafny.IntOfUint32((_154_v91).Cardinality()))).Uint32()).(*C4), 6)
			(_nw30).ArraySet1(_136_v80, 7)
			_155_v92 = _nw30
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_155_v92), 0))
			_ = _index20
			(_155_v92).ArraySet1(_136_v80, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_155_v92), 0))
			_ = _index21
			(_155_v92).ArraySet1(_136_v80, (_index21).Int())
			(_23_globalState).F0 = ((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool)
			var _156_v93 bool
			_ = _156_v93
			var _157_v94 _dafny.Int
			_ = _157_v94
			var _158_v95 _dafny.Int
			_ = _158_v95
			var _159_v96 _dafny.Int
			_ = _159_v96
			var _out8 bool
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = (_138_v81).M4((_80_v51).F20(), _23_globalState)
			_156_v93 = _out8
			_157_v94 = _out9
			_158_v95 = _out10
			_159_v96 = _out11
			_24_v9 = (_80_v51).F20()
		} else {
			var _160___mcc_h15 D1 = _source3.Get_().(D1_DC6).Cf12
			_ = _160___mcc_h15
			var _161_cf12 D1 = _160___mcc_h15
			_ = _161_cf12
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))
			_ = _index22
			((_31_v15).F23()).ArraySet1(_48_v28, (_index22).Int())
			var _162_v97 *C2
			_ = _162_v97
			var _nw31 *C2 = New_C2_()
			_ = _nw31
			_nw31.Ctor__(_138_v81.F25, (_80_v51).F20(), (_80_v51).F19(), _138_v81.F26)
			_162_v97 = _nw31
			var _163_v98 _dafny.Set
			_ = _163_v98
			_163_v98 = _dafny.SetOf((_80_v51).F20())
			var _164_v99 _dafny.Set
			_ = _164_v99
			_164_v99 = _dafny.SetOf((_dafny.SetOf((_80_v51).F20())).Difference(_163_v98))
			_164_v99 = Companion_Default___.Fm28(_137_i10, Companion_Default___.SafeModuloInt(Companion_Default___.Fm0((_80_v51).F20(), _162_v97.F26, _23_globalState), (_136_v80).F21()), _23_globalState)
			(_23_globalState).F13 = _13_v0
		}
		var _165_v100 bool
		_ = _165_v100
		var _166_v101 _dafny.Int
		_ = _166_v101
		var _167_v102 _dafny.Int
		_ = _167_v102
		var _168_v103 _dafny.Int
		_ = _168_v103
		var _out12 bool
		_ = _out12
		var _out13 _dafny.Int
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out12, _out13, _out14, _out15 = (_138_v81).M4(!(_48_v28) || (_138_v81.F26), _23_globalState)
		_165_v100 = _out12
		_166_v101 = _out13
		_167_v102 = _out14
		_168_v103 = _out15
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_16_v3), 0))
		_ = _index23
		(_16_v3).ArraySet1(_168_v103, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_16_v3), 0))
		_ = _index24
		var _rhs15 *C3 = (_31_v15)
		_ = _rhs15
		var _rhs16 _dafny.Int = _168_v103
		_ = _rhs16
		var _lhs14 _dafny.Array = _16_v3
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_16_v3), 0))
		_ = _lhs15
		_31_v15 = _rhs15
		(_lhs14).ArraySet1(_rhs16, (_lhs15).Int())
	}
	var _169_v104 _dafny.Map
	_ = _169_v104
	_169_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v0, Companion_Default___.Fm7(_13_v0, _13_v0, _dafny.IntOfUint32((_30_v14).Cardinality()), _23_globalState))
	var _170_v107 D0
	_ = _170_v107
	_170_v107 = Companion_D0_.Create_DC1_(true, (_136_v80).F21(), (_79_v50).F24())
	var _171_v108 _dafny.MultiSet
	_ = _171_v108
	_171_v108 = _dafny.MultiSetOf((_169_v104).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_79_v50).F24(), _48_v28)), func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(929), _dafny.IntOfInt64(294))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _172_v105 _dafny.Int
			_172_v105 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(929)).Cmp(_172_v105) <= 0) && ((_172_v105).Cmp(_dafny.IntOfInt64(294)) < 0) {
				_coll7.Add((_172_v105).Plus((_136_v80).F21()), ((_31_v15).F23()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen(((_31_v15).F23()), 0))).Int()).(bool))
			}
		}
		return _coll7.ToMap()
	}(), (func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_14_v1).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _173_v106 _dafny.Int
			_173_v106 = interface{}(_compr_8).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_14_v1, _173_v106) {
				_coll8.Add((_173_v106).Times((_136_v80).F21()), _48_v28)
			}
		}
		return _coll8.ToMap()
	}()).Merge((_169_v104).Update((_79_v50).F24(), (_170_v107).Dtor_cf1())))
	(_23_globalState).F13 = (func() _dafny.Int {
		if (_171_v108).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_79_v50).F24()), !(_48_v28))).Merge(_169_v104)) {
			return (_171_v108).Multiplicity((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_79_v50).F24()), !(_48_v28))).Merge(_169_v104))
		}
		return (_136_v80).F21()
	})()
	_dafny.Print(_13_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_14_v1, _dafny.SeqOf(_dafny.IntOfInt64(888), _dafny.IntOfInt64(2), _dafny.IntOfInt64(888), _dafny.IntOfInt64(2), _dafny.IntOfInt64(648), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_15_v2, _dafny.SeqOf(_dafny.IntOfInt64(888), _dafny.IntOfInt64(2), _dafny.IntOfInt64(888), _dafny.IntOfInt64(2), _dafny.IntOfInt64(648), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_17_v4).Equals(_dafny.SetOf(_dafny.IntOfInt64(888))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_18_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(888)), _dafny.IntOfInt64(888))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_19_v6).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_21_v7).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_23_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_23_globalState).F1(), _dafny.SeqOf(_dafny.IntOfInt64(888), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState.F4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(888)), _dafny.IntOfInt64(888))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_23_globalState).F7()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_23_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_23_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_23_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_24_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_29_v13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_30_v14, _dafny.SeqOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_31_v15.F22, _dafny.SeqOf(_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_31_v15).F23()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_32_v16).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_35_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_47_v27, _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_48_v28)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_49_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_67_v44).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_68_v45).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(888), _dafny.IntOfInt64(-409))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v46).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(888))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_70_v47).Dtor_cf12()).Dtor_cf11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_v48).Dtor_cf12()).Dtor_cf12()).Dtor_cf11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v50).F24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v51).F19()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v51).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v52).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v53).Equals(_dafny.MultiSetOf(_dafny.CodePoint('j'), _dafny.CodePoint('o'), _dafny.CodePoint('y'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v77).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v78).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v79).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v80).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v104).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(888), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v107).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v107).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v107).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v108).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(888), true), _dafny.NewMapBuilder().ToMap(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, true).UpdateUnsafe(_dafny.IntOfInt64(888), true))))
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
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 bool
	Cf5 _dafny.Int
	Cf6 _dafny.Sequence
	Cf7 _dafny.Int
	Cf8 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 bool, Cf5 _dafny.Int, Cf6 _dafny.Sequence, Cf7 _dafny.Int, Cf8 _dafny.Int) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6, Cf7, Cf8}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf9 _dafny.Map
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf9 _dafny.Map) D0 {
	return D0{D0_DC3{Cf9}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.MultiSet
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.MultiSet) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf8
}

func (_this D0) Dtor_cf9() _dafny.Map {
	return _this.Get_().(D0_DC3).Cf9
}

func (_this D0) Dtor_cf0() _dafny.MultiSet {
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
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Equals(data2.Cf6) && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf9.Equals(data2.Cf9)
		}
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

type D1_DC5 struct {
	Cf11 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf11 bool) D1 {
	return D1{D1_DC5{Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf10 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf10 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf10}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC6 struct {
	Cf12 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf12 D1) D1 {
	return D1{D1_DC6{Cf12}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false)
}

func (_this D1) Dtor_cf11() bool {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf12() D1 {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf11 == data2.Cf11
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf14 _dafny.Int
	Cf15 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf14 _dafny.Int, Cf15 _dafny.Int) D2 {
	return D2{D2_DC8{Cf14, Cf15}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf13 _dafny.Sequence
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 _dafny.Sequence) D2 {
	return D2{D2_DC7{Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf17 _dafny.Array
	Cf18 _dafny.Map
	Cf19 _dafny.Set
	Cf20 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf17 _dafny.Array, Cf18 _dafny.Map, Cf19 _dafny.Set, Cf20 _dafny.Int) D3 {
	return D3{D3_DC10{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf21 _dafny.Int
	Cf22 _dafny.Int
	Cf23 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf21 _dafny.Int, Cf22 _dafny.Int, Cf23 bool) D3 {
	return D3{D3_DC11{Cf21, Cf22, Cf23}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf16 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf16 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf16}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMap, _dafny.EmptySet, _dafny.Zero)
}

func (_this D3) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf19() _dafny.Set {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf20
}

func (_this D3) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf21
}

func (_this D3) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf22
}

func (_this D3) Dtor_cf23() bool {
	return _this.Get_().(D3_DC11).Cf23
}

func (_this D3) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + data.Cf16.VerbatimString(true) + ")"
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
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18) && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
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

type D4_DC13 struct {
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_() D4 {
	return D4{D4_DC13{}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf25 bool
	Cf26 _dafny.Set
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf25 bool, Cf26 _dafny.Set) D4 {
	return D4{D4_DC14{Cf25, Cf26}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC12 struct {
	Cf24 _dafny.CodePoint
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf24 _dafny.CodePoint) D4 {
	return D4{D4_DC12{Cf24}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC15 struct {
	Cf27 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf27 D4) D4 {
	return D4{D4_DC15{Cf27}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_()
}

func (_this D4) Dtor_cf25() bool {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) Dtor_cf26() _dafny.Set {
	return _this.Get_().(D4_DC14).Cf26
}

func (_this D4) Dtor_cf24() _dafny.CodePoint {
	return _this.Get_().(D4_DC12).Cf24
}

func (_this D4) Dtor_cf27() D4 {
	return _this.Get_().(D4_DC15).Cf27
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf27) + ")"
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
			_, ok := other.Get_().(D4_DC13)
			return ok
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26.Equals(data2.Cf26)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf27.Equals(data2.Cf27)
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

type D5_DC17 struct {
	Cf29 _dafny.Int
	Cf30 _dafny.Int
	Cf31 D0
	Cf32 bool
	Cf33 bool
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf29 _dafny.Int, Cf30 _dafny.Int, Cf31 D0, Cf32 bool, Cf33 bool) D5 {
	return D5{D5_DC17{Cf29, Cf30, Cf31, Cf32, Cf33}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC18 struct {
	Cf34 _dafny.Sequence
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf34 _dafny.Sequence) D5 {
	return D5{D5_DC18{Cf34}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC19 struct {
	Cf35 _dafny.Int
	Cf36 bool
	Cf37 _dafny.Array
	Cf38 bool
	Cf39 _dafny.CodePoint
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf35 _dafny.Int, Cf36 bool, Cf37 _dafny.Array, Cf38 bool, Cf39 _dafny.CodePoint) D5 {
	return D5{D5_DC19{Cf35, Cf36, Cf37, Cf38, Cf39}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC16 struct {
	Cf28 _dafny.Array
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf28 _dafny.Array) D5 {
	return D5{D5_DC16{Cf28}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(_dafny.Zero, _dafny.Zero, Companion_D0_.Default(), false, false)
}

func (_this D5) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf30
}

func (_this D5) Dtor_cf31() D0 {
	return _this.Get_().(D5_DC17).Cf31
}

func (_this D5) Dtor_cf32() bool {
	return _this.Get_().(D5_DC17).Cf32
}

func (_this D5) Dtor_cf33() bool {
	return _this.Get_().(D5_DC17).Cf33
}

func (_this D5) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D5_DC18).Cf34
}

func (_this D5) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D5_DC19).Cf35
}

func (_this D5) Dtor_cf36() bool {
	return _this.Get_().(D5_DC19).Cf36
}

func (_this D5) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D5_DC19).Cf37
}

func (_this D5) Dtor_cf38() bool {
	return _this.Get_().(D5_DC19).Cf38
}

func (_this D5) Dtor_cf39() _dafny.CodePoint {
	return _this.Get_().(D5_DC19).Cf39
}

func (_this D5) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D5_DC16).Cf28
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0 && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Equals(data2.Cf31) && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf28 == data2.Cf28
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

type D6_DC20 struct {
	Cf40 _dafny.Array
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf40 _dafny.Array) D6 {
	return D6{D6_DC20{Cf40}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D6) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D6_DC20).Cf40
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC20:
		{
			data2, ok := other.Get_().(D6_DC20)
			return ok && data1.Cf40 == data2.Cf40
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
	Cf41 _dafny.Array
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf41 _dafny.Array) D7 {
	return D7{D7_DC21{Cf41}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D7) Dtor_cf41() _dafny.Array {
	return _this.Get_().(D7_DC21).Cf41
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf41 == data2.Cf41
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
	Cf43 bool
	Cf44 _dafny.Int
	Cf45 _dafny.Int
	Cf46 bool
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf43 bool, Cf44 _dafny.Int, Cf45 _dafny.Int, Cf46 bool) D8 {
	return D8{D8_DC23{Cf43, Cf44, Cf45, Cf46}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf47 bool
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf47 bool) D8 {
	return D8{D8_DC24{Cf47}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC22 struct {
	Cf42 *C2
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf42 *C2) D8 {
	return D8{D8_DC22{Cf42}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC25 struct {
	Cf48 D8
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf48 D8) D8 {
	return D8{D8_DC25{Cf48}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(false, _dafny.Zero, _dafny.Zero, false)
}

func (_this D8) Dtor_cf43() bool {
	return _this.Get_().(D8_DC23).Cf43
}

func (_this D8) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf44
}

func (_this D8) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf45
}

func (_this D8) Dtor_cf46() bool {
	return _this.Get_().(D8_DC23).Cf46
}

func (_this D8) Dtor_cf47() bool {
	return _this.Get_().(D8_DC24).Cf47
}

func (_this D8) Dtor_cf42() *C2 {
	return _this.Get_().(D8_DC22).Cf42
}

func (_this D8) Dtor_cf48() D8 {
	return _this.Get_().(D8_DC25).Cf48
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf42) + ")"
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
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf43 == data2.Cf43 && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf47 == data2.Cf47
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf42 == data2.Cf42
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
	Cf49 _dafny.Set
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf49 _dafny.Set) D9 {
	return D9{D9_DC26{Cf49}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D9) Dtor_cf49() _dafny.Set {
	return _this.Get_().(D9_DC26).Cf49
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D10_DC27 struct {
	Cf50 *C3
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf50 *C3) D10 {
	return D10{D10_DC27{Cf50}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

func (CompanionStruct_D10_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D10) Dtor_cf50() *C3 {
	return _this.Get_().(D10_DC27).Cf50
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf50 == data2.Cf50
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
	Cf51 _dafny.Map
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf51 _dafny.Map) D11 {
	return D11{D11_DC28{Cf51}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D11) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D11_DC28).Cf51
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf51) + ")"
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
			return ok && data1.Cf51.Equals(data2.Cf51)
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

// Definition of trait T0
type T0 interface {
	String() string
	F19() _dafny.Array
	F20() bool
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
	F0   bool
	F3   _dafny.Array
	F4   _dafny.Map
	F10  _dafny.Array
	F13  _dafny.Int
	F18  _dafny.Int
	_f1  _dafny.Sequence
	_f2  _dafny.Int
	_f5  _dafny.Int
	_f6  _dafny.Int
	_f7  _dafny.Map
	_f8  bool
	_f9  _dafny.Array
	_f11 bool
	_f12 bool
	_f14 _dafny.Array
	_f15 _dafny.Int
	_f16 bool
	_f17 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F4 = _dafny.EmptyMap
	_this.F10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F13 = _dafny.Zero
	_this.F18 = _dafny.Zero
	_this._f1 = _dafny.EmptySeq
	_this._f2 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.EmptyMap
	_this._f8 = false
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = false
	_this._f12 = false
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.Zero
	_this._f16 = false
	_this._f17 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Sequence, f2 _dafny.Int, f3 _dafny.Array, f4 _dafny.Map, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.Map, f8 bool, f9 _dafny.Array, f10 _dafny.Array, f11 bool, f12 bool, f13 _dafny.Int, f14 _dafny.Array, f15 _dafny.Int, f16 bool, f17 bool, f18 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this).F18 = f18
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Map {
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
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Array {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f24 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f24 = _dafny.Zero
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

func (_this *C0) Ctor__(f24 _dafny.Int) {
	{
		(_this)._f24 = f24
	}
}
func (_this *C0) F24() _dafny.Int {
	{
		return _this._f24
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f19 _dafny.Array
	_f20 bool
	_f27 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f20 = false
	_this._f27 = _dafny.EmptySeq
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

func (_this *C1) F19() _dafny.Array {
	return _this._f19
}
func (_this *C1) F20() bool {
	return _this._f20
}
func (_this *C1) Ctor__(f27 _dafny.Sequence, f19 _dafny.Array, f20 bool) {
	{
		(_this)._f27 = f27
		(_this)._f19 = f19
		(_this)._f20 = f20
	}
}
func (_this *C1) Fm13(globalState *GlobalState) bool {
	{
		return (_this).F20()
	}
}
func (_this *C1) Fm14(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(780))).Uint32(), func(coer6 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_174_i0 _dafny.Int) D0 {
			return Companion_D0_.Create_DC2_((_this).F20(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _dafny.IntOfUint32(((_this).F27()).Cardinality()))).Cardinality(), _dafny.SeqOf(!((_this).F20()), (_this).F20()), _dafny.IntOfInt64(572), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(293), Companion_D0_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F20(), (_this).F20())), (_dafny.MultiSetOf((_this).F20(), (_this).F20(), (_this).F20(), (_this).F20())).Cardinality())))).Cardinality())
		}))), (_this).F20())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(398))).Uint32(), func(coer7 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_175_i1 _dafny.Int) D0 {
			return Companion_D0_.Create_DC2_((_this).F20(), _dafny.IntOfInt64(756), _dafny.SeqOf((_this).F20()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(319))).Cardinality()), _dafny.IntOfInt64(449))
		}))), (_this).F20())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(802))).Uint32(), func(coer8 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_176_i2 _dafny.Int) D0 {
			return Companion_D0_.Create_DC2_((_this).F20(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (Companion_D0_.Create_DC1_((_this).F20(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F27()).Cardinality()), (_dafny.MultiSetOf((_this).F20())).Cardinality())).Cardinality(), _dafny.IntOfInt64(-88))).Cardinality()), _dafny.IntOfUint32(((_this).F27()).Cardinality()))).Dtor_cf1())).Cardinality(), _dafny.SeqOf(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(951), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yhnf")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())).Cardinality())
		}))), !((_this).F20()))))
	}
}
func (_this *C1) F27() _dafny.Sequence {
	{
		return _this._f27
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f19 _dafny.Array
	_f20 bool
	F25  *C0
	F26  bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f20 = false
	_this.F25 = (*C0)(nil)
	_this.F26 = false
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

func (_this *C2) F19() _dafny.Array {
	return _this._f19
}
func (_this *C2) F20() bool {
	return _this._f20
}
func (_this *C2) Ctor__(f25 *C0, f26 bool, f19 _dafny.Array, f20 bool) {
	{
		(_this).F25 = f25
		(_this).F26 = f26
		(_this)._f19 = f19
		(_this)._f20 = f20
	}
}
func (_this *C2) M4(p0 bool, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _177_v0 _dafny.Sequence
		_ = _177_v0
		_177_v0 = _dafny.SeqOf((_this.F25).F24())
		var _hi3 _dafny.Int = ((_this.F25).F24()).Plus(_dafny.IntOfUint32((_177_v0).Cardinality()))
		_ = _hi3
		for _178_i0 := (_this.F25).F24(); _178_i0.Cmp(_hi3) < 0; _178_i0 = _178_i0.Plus(_dafny.One) {
			var _179_v1 _dafny.Sequence
			_ = _179_v1
			_179_v1 = _dafny.SeqOf(_dafny.CodePoint('x'))
			var _180_v2 _dafny.Map
			_ = _180_v2
			_180_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(424), (_this).F20())
			_179_v1 = Companion_Default___.Fm8(_179_v1, _180_v2, _178_i0, globalState)
			r1 = (_dafny.Zero).Minus(((_this.F25).F24()).Plus((_this.F25).F24()))
			var _181_v3 _dafny.Sequence
			_ = _181_v3
			_181_v3 = _dafny.SeqOf(_this.F26, p0)
			var _182_v4 *C0
			_ = _182_v4
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__(_dafny.IntOfUint32((_181_v3).Cardinality()))
			_182_v4 = _nw32
			(globalState).F0 = _this.F26
		}
		var _183_v5 _dafny.Sequence
		_ = _183_v5
		_183_v5 = _dafny.SeqOf(_this.F26)
		var _184_v6 _dafny.Map
		_ = _184_v6
		_184_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this.F25).F24())
		var _185_v7 _dafny.MultiSet
		_ = _185_v7
		_185_v7 = _dafny.MultiSetOf(_184_v6, _184_v6)
		var _186_v8 D0
		_ = _186_v8
		_186_v8 = Companion_D0_.Create_DC2_(true, (_this.F25).F24(), _183_v5, (func() _dafny.Int {
			if (_185_v7).Contains(_184_v6) {
				return (_185_v7).Multiplicity(_184_v6)
			}
			return (_dafny.SetOf(_dafny.IntOfInt64(137))).Cardinality()
		})(), (_this.F25).F24())
		var _187_v9 _dafny.Sequence
		_ = _187_v9
		_187_v9 = _dafny.SeqOf(_186_v8, _186_v8)
		var _188_v10 _dafny.Map
		_ = _188_v10
		_188_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_this.F26, (_this).F20())).Cardinality(), (_this.F25).F24())
		var _189_v11 _dafny.Map
		_ = _189_v11
		_189_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_188_v10).Cardinality(), p0)
		var _190_v12 D1
		_ = _190_v12
		_190_v12 = Companion_D1_.Create_DC4_(_dafny.SeqOf(Companion_D0_.Create_DC2_((_this).F20(), (_this.F25).F24(), _183_v5, (_dafny.Zero).Minus((_this.F25).F24()), (_189_v11).Cardinality())))
		var _191_v13 _dafny.Array
		_ = _191_v13
		var _nwElement0_2 D1 = (func() D1 {
			if (_this).F20() {
				return Companion_D1_.Create_DC4_(_dafny.Companion_Sequence_.Update(_187_v9, (Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_187_v9).Cardinality()))).Uint32(), _186_v8))
			}
			return _190_v12
		})()
		_ = _nwElement0_2
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_2, 0)
		(_nw33).ArraySet1(_190_v12, 1)
		_191_v13 = _nw33
		var _pat_let_tv7 = _187_v9
		_ = _pat_let_tv7
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_191_v13), 0))
		_ = _index25
		(_191_v13).ArraySet1(func(_pat_let0_0 D1) D1 {
			return func(_192_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 _dafny.Sequence) D1 {
					return func(_193_dt__update_hcf10_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC4_(_193_dt__update_hcf10_h0)
					}(_pat_let1_0)
				}(_pat_let_tv7)
			}(_pat_let0_0)
		}(_190_v12), (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_191_v13), 0))
		_ = _index26
		(_191_v13).ArraySet1(_190_v12, (_index26).Int())
		if (_this).F20() {
			r0 = Companion_Default___.Fm9((_183_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-479), _dafny.IntOfUint32((_183_v5).Cardinality()))).Uint32()).(bool), globalState)
			var _194_v14 _dafny.Set
			_ = _194_v14
			_194_v14 = _dafny.SetOf((_this).F20(), _this.F26)
			var _195_v15 D1
			_ = _195_v15
			_195_v15 = Companion_D1_.Create_DC5_(false)
			var _196_v16 _dafny.MultiSet
			_ = _196_v16
			_196_v16 = _dafny.MultiSetOf(_195_v15)
			var _197_v17 _dafny.Map
			_ = _197_v17
			_197_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v14, _196_v16)
			var _198_v18 _dafny.Array
			_ = _198_v18
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw34
			_198_v18 = _nw34
			var _199_v19 _dafny.Sequence
			_ = _199_v19
			_199_v19 = _dafny.SeqOf(_198_v18, _198_v18)
			var _200_v20 _dafny.Map
			_ = _200_v20
			_200_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_197_v17, (_199_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.IntOfUint32((_199_v19).Cardinality()))).Uint32()).(_dafny.Array))
			var _201_v21 _dafny.Map
			_ = _201_v21
			_201_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F25).F24(), _194_v14)
			_200_v20 = (_200_v20).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				if (_201_v21).Contains((_this.F25).F24()) {
					return (_201_v21).Get((_this.F25).F24()).(_dafny.Set)
				}
				return _194_v14
			})(), _dafny.MultiSetOf(_195_v15, Companion_D1_.Create_DC5_((_this).F20()), _195_v15)), _198_v18)
			(globalState).F13 = ((_this.F25).F24()).Times((_this.F25).F24())
			var _202_v22 *C0
			_ = _202_v22
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__((_this.F25).F24())
			_202_v22 = _nw35
			var _203_v23 _dafny.Array
			_ = _203_v23
			var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw36
			_203_v23 = _nw36
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_203_v23), 0))
			_ = _index27
			(_203_v23).ArraySet1((_202_v22).F24(), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_203_v23), 0))
			_ = _index28
			(_203_v23).ArraySet1((_this.F25).F24(), (_index28).Int())
		} else {
			var _204_v24 _dafny.Array
			_ = _204_v24
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
			_ = _nw37
			_204_v24 = _nw37
			var _205_v25 _dafny.Map
			_ = _205_v25
			_205_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen((_204_v24), 0))
			_ = _index29
			(_204_v24).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F26)).Merge(_205_v25), (_index29).Int())
			var _206_v26 _dafny.Sequence
			_ = _206_v26
			_206_v26 = _dafny.SeqOf((Companion_Default___.Fm10((_this.F25).F24(), _dafny.IntOfUint32((_177_v0).Cardinality()), (_this.F25).F24(), globalState)).Merge(_205_v25))
			var _207_v27 _dafny.MultiSet
			_ = _207_v27
			_207_v27 = _dafny.MultiSetOf(p0, !(_this.F26))
			var _208_v28 _dafny.Array
			_ = _208_v28
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw38
			_208_v28 = _nw38
			var _209_v29 _dafny.MultiSet
			_ = _209_v29
			_209_v29 = _dafny.MultiSetOf(_208_v28)
			var _210_v30 _dafny.Set
			_ = _210_v30
			_210_v30 = _dafny.SetOf((_this.F25).F24(), (func() _dafny.Int {
				if (_209_v29).Contains(_208_v28) {
					return (_209_v29).Multiplicity(_208_v28)
				}
				return (_this.F25).F24()
			})())
			var _211_v31 _dafny.Map
			_ = _211_v31
			_211_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_207_v27).Cardinality()), _210_v30)
			var _212_v32 _dafny.Sequence
			_ = _212_v32
			_212_v32 = _dafny.UnicodeSeqOfUtf8Bytes("r")
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.ArrayLen((_204_v24), 0))
			_ = _index30
			(_204_v24).ArraySet1((_206_v26).Select((Companion_Default___.SafeIndex(((_211_v31).Cardinality()).Times(_dafny.IntOfUint32((_212_v32).Cardinality())), _dafny.IntOfUint32((_206_v26).Cardinality()))).Uint32()).(_dafny.Map), (_index30).Int())
			_188_v10 = (_188_v10).Merge((func() _dafny.Map {
				if p0 {
					return _188_v10
				}
				return (_188_v10).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this.F25).F24()), (Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_dafny.SeqOf((_this.F25).F24())).Cardinality()))).Uint32(), (_this.F25).F24())).Cardinality()), (_this.F25).F24())
			})())
			var _213_v33 D2
			_ = _213_v33
			_213_v33 = Companion_D2_.Create_DC7_(_dafny.SeqOf((_this.F25).F24()))
			var _214_v34 _dafny.Set
			_ = _214_v34
			_214_v34 = _dafny.SetOf(_213_v33)
			var _215_v35 _dafny.Sequence
			_ = _215_v35
			_215_v35 = _dafny.SeqOf(_214_v34, _214_v34)
			(_this).F26 = (_dafny.SetOf(_213_v33, _213_v33)).IsDisjointFrom((_215_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_212_v32).Cardinality()), _dafny.IntOfUint32((_215_v35).Cardinality()))).Uint32()).(_dafny.Set))
			var _216_v36 _dafny.Array
			_ = _216_v36
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_5
			var _nw39 _dafny.Array
			_ = _nw39
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw39 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.CodePoint = func(_217_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				}
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw39 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw39).ArraySet1CodePoint(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw39).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_216_v36 = _nw39
			_216_v36 = _216_v36
			(globalState).F0 = (_this).F20()
		}
		var _218_v37 _dafny.Array
		_ = _218_v37
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_6
		var _nw40 _dafny.Array
		_ = _nw40
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw40 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = func(_219_i2 _dafny.Int) _dafny.Int {
				return (_219_i2).Plus((_this.F25).F24())
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw40 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw40).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw40).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_218_v37 = _nw40
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_218_v37), 0))
		_ = _index31
		(_218_v37).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm11((_this.F25).F24(), globalState)).Cardinality()), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_218_v37), 0))
		_ = _index32
		(_218_v37).ArraySet1((_this.F25).F24(), (_index32).Int())
		var _220_v38 _dafny.Sequence
		_ = _220_v38
		_220_v38 = _dafny.UnicodeSeqOfUtf8Bytes("nmahe")
		var _221_v39 _dafny.CodePoint
		_ = _221_v39
		_221_v39 = _dafny.CodePoint('m')
		var _222_v40 D3
		_ = _222_v40
		_222_v40 = Companion_D3_.Create_DC9_(_220_v38)
		var _223_v42 _dafny.Set
		_ = _223_v42
		_223_v42 = _dafny.SetOf((_218_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_218_v37), 0))).Int()).(_dafny.Int))
		var _224_v43 _dafny.Array
		_ = _224_v43
		var _nwElement0_3 _dafny.Sequence = _220_v38
		_ = _nwElement0_3
		var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
		_ = _nw41
		(_nw41).ArraySet1(_nwElement0_3, 0)
		(_nw41).ArraySet1(_220_v38, 1)
		(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_220_v38, _220_v38), 2)
		(_nw41).ArraySet1(_220_v38, 3)
		(_nw41).ArraySet1(_220_v38, 4)
		(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cnp"), 5)
		(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("iqfjf"), (Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqfjf")).Cardinality()))).Uint32(), _221_v39), _220_v38), 6)
		(_nw41).ArraySet1(_220_v38, 7)
		(_nw41).ArraySet1((_222_v40).Dtor_cf16(), 8)
		(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_220_v38, _220_v38), 9)
		(_nw41).ArraySet1(_dafny.Companion_Sequence_.Update(_220_v38, (Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_220_v38).Cardinality()))).Uint32(), _221_v39), 10)
		(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cbiytnt"), 11)
		(_nw41).ArraySet1(_220_v38, 12)
		(_nw41).ArraySet1(_220_v38, 13)
		(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uc"), _dafny.UnicodeSeqOfUtf8Bytes("b")), 14)
		(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uusrltbmy"), 15)
		(_nw41).ArraySet1(Companion_Default___.Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(88))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_225_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_226_i4 _dafny.Int) _dafny.CodePoint {
				return _225_v39
			}
		})(_221_v39))), (func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_177_v0).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _227_v41 _dafny.Int
				_227_v41 = interface{}(_compr_9).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_177_v0, _227_v41) {
					_coll9.Add((_227_v41).Times((_this.F25).F24()), (_this).F20())
				}
			}
			return _coll9.ToMap()
		}()).Update((_223_v42).Cardinality(), _this.F26), (_218_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_218_v37), 0))).Int()).(_dafny.Int), globalState), 16)
		_224_v43 = _nw41
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_224_v43), 0))); ; {
			_guard_loop_0, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _228_i3 _dafny.Int
			_228_i3 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_228_i3).Sign() != -1) && ((_228_i3).Cmp(_dafny.ArrayLen((_224_v43), 0)) < 0)) {
				(_224_v43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _220_v38), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(40))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_229_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_230_i5 _dafny.Int) _dafny.CodePoint {
						return _229_v39
					}
				})(_221_v39)))), (_228_i3).Int())
			}
		}
		var _hi4 _dafny.Int = (_218_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_218_v37), 0))).Int()).(_dafny.Int)
		_ = _hi4
		for _231_i6 := ((_this.F25).F24()).Minus((_this.F25).F24()); _231_i6.Cmp(_hi4) < 0; _231_i6 = _231_i6.Plus(_dafny.One) {
			r0 = _this.F26
			(globalState).F13 = (func() _dafny.Int {
				if (_this.F26) == (!((_this).F20())) {
					return (_this.F25).F24()
				}
				return _dafny.IntOfInt64(128)
			})()
			var _232_v44 _dafny.Map
			_ = _232_v44
			_232_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_i6, (_this).F19())
			var _233_v45 _dafny.Array
			_ = _233_v45
			var _nwElement0_4 _dafny.Array = (_this).F19()
			_ = _nwElement0_4
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(13))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_4, 0)
			(_nw42).ArraySet1((_this).F19(), 1)
			(_nw42).ArraySet1((_this).F19(), 2)
			(_nw42).ArraySet1((_this).F19(), 3)
			(_nw42).ArraySet1((func() _dafny.Array {
				if (_232_v44).Contains(Companion_Default___.Fm0(p0, _this.F26, globalState)) {
					return (_232_v44).Get(Companion_Default___.Fm0(p0, _this.F26, globalState)).(_dafny.Array)
				}
				return (_this).F19()
			})(), 4)
			(_nw42).ArraySet1((_this).F19(), 5)
			(_nw42).ArraySet1((_this).F19(), 6)
			(_nw42).ArraySet1((_this).F19(), 7)
			(_nw42).ArraySet1((_this).F19(), 8)
			(_nw42).ArraySet1((_this).F19(), 9)
			(_nw42).ArraySet1((func() _dafny.Array {
				if (_this).F20() {
					return (_this).F19()
				}
				return (_this).F19()
			})(), 10)
			(_nw42).ArraySet1((_this).F19(), 11)
			(_nw42).ArraySet1((_this).F19(), 12)
			_233_v45 = _nw42
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_233_v45), 0))
			_ = _index33
			(_233_v45).ArraySet1((_this).F19(), (_index33).Int())
			var _234_v46 _dafny.Set
			_ = _234_v46
			_234_v46 = _dafny.SetOf(p0)
			var _235_v47 _dafny.Map
			_ = _235_v47
			_235_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _236_v48 _dafny.Map
			_ = _236_v48
			_236_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_177_v0).Cardinality()), _235_v47)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_233_v45), 0))
			_ = _index34
			var _nwElement0_5 bool = _this.F26
			_ = _nwElement0_5
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(7))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_5, 0)
			(_nw43).ArraySet1((_234_v46).IsProperSubsetOf(_234_v46), 1)
			(_nw43).ArraySet1(_dafny.Companion_Sequence_.Contains(_220_v38, _221_v39), 2)
			(_nw43).ArraySet1((_this).F20(), 3)
			(_nw43).ArraySet1((_this).F20(), 4)
			(_nw43).ArraySet1((_223_v42).IsDisjointFrom(_dafny.SetOf(((func() _dafny.Map {
				if (_236_v48).Contains(_dafny.IntOfInt64(120)) {
					return (_236_v48).Get(_dafny.IntOfInt64(120)).(_dafny.Map)
				}
				return _235_v47
			})()).Cardinality())), 5)
			(_nw43).ArraySet1(p0, 6)
			(_233_v45).ArraySet1(_nw43, (_index34).Int())
			var _237_v49 _dafny.MultiSet
			_ = _237_v49
			_237_v49 = _dafny.MultiSetOf(Companion_Default___.Fm0(_this.F26, _this.F26, globalState), (_this.F25).F24(), (_this.F25).F24())
			var _238_v50 _dafny.MultiSet
			_ = _238_v50
			_238_v50 = _dafny.MultiSetOf(_this.F26, p0)
			(globalState).F0 = (Companion_Default___.Fm12(_237_v49, (func() _dafny.Int {
				if (_238_v50).Contains(!((_this).F20())) {
					return (_238_v50).Multiplicity(!((_this).F20()))
				}
				return (_this.F25).F24()
			})(), globalState)).Contains((_234_v46).Cardinality())
		}
		r0 = Companion_Default___.Fm9(p0, globalState)
		r1 = _dafny.IntOfInt64(893)
		r2 = (func() _dafny.Int {
			if (_184_v6).Contains(!(((_this.F25).F24()).Cmp((_this.F25).F24()) >= 0)) {
				return (_184_v6).Get(!(((_this.F25).F24()).Cmp((_this.F25).F24()) >= 0)).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_188_v10).Contains((_this.F25).F24()) {
					return (_188_v10).Get((_this.F25).F24()).(_dafny.Int)
				}
				return (_this.F25).F24()
			})()
		})()
		r3 = (_218_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_218_v37), 0))).Int()).(_dafny.Int)
		return r0, r1, r2, r3
	}
}
func (_this *C2) M5(globalState *GlobalState) (bool, bool, _dafny.Int, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _239_v0 _dafny.Sequence
		_ = _239_v0
		_239_v0 = _dafny.SeqOf(_this.F26)
		var _240_v1 _dafny.Sequence
		_ = _240_v1
		_240_v1 = _dafny.SeqOf(Companion_D0_.Create_DC2_((_this).F20(), (_this.F25).F24(), _239_v0, (_this.F25).F24(), (_this.F25).F24()))
		var _241_v2 D1
		_ = _241_v2
		_241_v2 = Companion_D1_.Create_DC4_(_240_v1)
		var _pat_let_tv8 = _240_v1
		_ = _pat_let_tv8
		var _source4 D1 = func(_pat_let2_0 D1) D1 {
			return func(_242_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let3_0 _dafny.Sequence) D1 {
					return func(_243_dt__update_hcf10_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC4_(_243_dt__update_hcf10_h0)
					}(_pat_let3_0)
				}(_pat_let_tv8)
			}(_pat_let2_0)
		}(_241_v2)
		_ = _source4
		if _source4.Is_DC5() {
			var _244___mcc_h0 bool = _source4.Get_().(D1_DC5).Cf11
			_ = _244___mcc_h0
			var _245_cf11 bool = _244___mcc_h0
			_ = _245_cf11
			var _246_v3 *C0
			_ = _246_v3
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__(Companion_Default___.Fm0((_this).F20(), _245_cf11, globalState))
			_246_v3 = _nw44
			var _247_v4 D1
			_ = _247_v4
			_247_v4 = Companion_D1_.Create_DC5_(Companion_Default___.Fm9(_this.F26, globalState))
			var _248_v5 D1
			_ = _248_v5
			_248_v5 = Companion_D1_.Create_DC6_(_247_v4)
			_248_v5 = _248_v5
			_245_cf11 = _245_cf11
			var _249_v6 _dafny.Sequence
			_ = _249_v6
			_249_v6 = _dafny.UnicodeSeqOfUtf8Bytes("tclfl")
			var _250_v7 T0
			_ = _250_v7
			var _nw45 *C1 = New_C1_()
			_ = _nw45
			_nw45.Ctor__(_249_v6, (_this).F19(), _245_cf11)
			_250_v7 = _nw45
			var _251_v8 _dafny.Map
			_ = _251_v8
			_251_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _250_v7)
			_251_v8 = (_251_v8).Update(true, (func() T0 {
				if _245_cf11 {
					return _250_v7
				}
				return _250_v7
			})())
		} else if _source4.Is_DC4() {
			var _252___mcc_h1 _dafny.Sequence = _source4.Get_().(D1_DC4).Cf10
			_ = _252___mcc_h1
			var _253_cf10 _dafny.Sequence = _252___mcc_h1
			_ = _253_cf10
			var _254_v9 _dafny.Array
			_ = _254_v9
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw46
			_254_v9 = _nw46
			var _255_v10 _dafny.CodePoint
			_ = _255_v10
			_255_v10 = _dafny.CodePoint('i')
			var _256_v11 _dafny.MultiSet
			_ = _256_v11
			_256_v11 = _dafny.MultiSetOf((_this).F20(), _this.F26)
			var _257_v12 _dafny.Array
			_ = _257_v12
			var _nwElement0_6 _dafny.CodePoint = _255_v10
			_ = _nwElement0_6
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(26))
			_ = _nw47
			(_nw47).ArraySet1CodePoint(_nwElement0_6, 0)
			(_nw47).ArraySet1CodePoint(_255_v10, 1)
			(_nw47).ArraySet1CodePoint(_255_v10, 2)
			(_nw47).ArraySet1CodePoint(_255_v10, 3)
			(_nw47).ArraySet1CodePoint(_255_v10, 4)
			(_nw47).ArraySet1CodePoint(_255_v10, 5)
			(_nw47).ArraySet1CodePoint(_255_v10, 6)
			(_nw47).ArraySet1CodePoint(_255_v10, 7)
			(_nw47).ArraySet1CodePoint(_255_v10, 8)
			(_nw47).ArraySet1CodePoint(_255_v10, 9)
			(_nw47).ArraySet1CodePoint(_255_v10, 10)
			(_nw47).ArraySet1CodePoint(_255_v10, 11)
			(_nw47).ArraySet1CodePoint(_255_v10, 12)
			(_nw47).ArraySet1CodePoint(_255_v10, 13)
			(_nw47).ArraySet1CodePoint(_255_v10, 14)
			(_nw47).ArraySet1CodePoint(_dafny.CodePoint('x'), 15)
			(_nw47).ArraySet1CodePoint(_255_v10, 16)
			(_nw47).ArraySet1CodePoint(Companion_Default___.Fm15((_256_v11).Cardinality(), _this.F26, (_this.F25).F24(), globalState), 17)
			(_nw47).ArraySet1CodePoint(_255_v10, 18)
			(_nw47).ArraySet1CodePoint(_255_v10, 19)
			(_nw47).ArraySet1CodePoint(_dafny.CodePoint('i'), 20)
			(_nw47).ArraySet1CodePoint(_dafny.CodePoint('u'), 21)
			(_nw47).ArraySet1CodePoint(_255_v10, 22)
			(_nw47).ArraySet1CodePoint(_255_v10, 23)
			(_nw47).ArraySet1CodePoint(_255_v10, 24)
			(_nw47).ArraySet1CodePoint(_255_v10, 25)
			_257_v12 = _nw47
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_254_v9), 0))
			_ = _index35
			(_254_v9).ArraySet1(_257_v12, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_254_v9), 0))
			_ = _index36
			(_254_v9).ArraySet1(_257_v12, (_index36).Int())
			var _258_v13 _dafny.Array
			_ = _258_v13
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
			_ = _nw48
			_258_v13 = _nw48
			_258_v13 = _258_v13
			var _259_v14 _dafny.Sequence
			_ = _259_v14
			_259_v14 = _dafny.UnicodeSeqOfUtf8Bytes("w")
			var _260_v15 D3
			_ = _260_v15
			_260_v15 = Companion_D3_.Create_DC9_(_259_v14)
			var _261_v16 D3
			_ = _261_v16
			_261_v16 = Companion_D3_.Create_DC9_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(273))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_262_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_263_i0 _dafny.Int) _dafny.CodePoint {
					return _262_v10
				}
			})(_255_v10))), (_260_v15).Dtor_cf16()))
			var _source5 D3 = _260_v15
			_ = _source5
			if _source5.Is_DC10() {
				var _264___mcc_h3 _dafny.Array = _source5.Get_().(D3_DC10).Cf17
				_ = _264___mcc_h3
				var _265___mcc_h4 _dafny.Map = _source5.Get_().(D3_DC10).Cf18
				_ = _265___mcc_h4
				var _266___mcc_h5 _dafny.Set = _source5.Get_().(D3_DC10).Cf19
				_ = _266___mcc_h5
				var _267___mcc_h6 _dafny.Int = _source5.Get_().(D3_DC10).Cf20
				_ = _267___mcc_h6
				var _268_cf20 _dafny.Int = _267___mcc_h6
				_ = _268_cf20
				var _269_cf19 _dafny.Set = _266___mcc_h5
				_ = _269_cf19
				var _270_cf18 _dafny.Map = _265___mcc_h4
				_ = _270_cf18
				var _271_cf17 _dafny.Array = _264___mcc_h3
				_ = _271_cf17
				var _272_v17 _dafny.Sequence
				_ = _272_v17
				_272_v17 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(132))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}((func(_273_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_274_i1 _dafny.Int) _dafny.CodePoint {
						return _273_v10
					}
				})(_255_v10)))).Cardinality()), (_this.F25).F24())
				var _275_v18 _dafny.Map
				_ = _275_v18
				_275_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_272_v17, _272_v17), ((_this.F25).F24()).Times(_dafny.IntOfUint32((_259_v14).Cardinality())))
				_275_v18 = (_275_v18).Update(_272_v17, (_this.F25).F24())
				var _276_v19 _dafny.Map
				_ = _276_v19
				_276_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_255_v10, (_this.F25).F24())
				var _277_v20 _dafny.MultiSet
				_ = _277_v20
				_277_v20 = _dafny.MultiSetOf((_this.F25).F24(), ((_270_cf18).Update(_this.F26, _dafny.IntOfInt64(850))).Cardinality(), (func() _dafny.Int {
					if (_256_v11).Contains((_this).F20()) {
						return (_256_v11).Multiplicity((_this).F20())
					}
					return _268_cf20
				})(), (_this.F25).F24(), (func() _dafny.Int {
					if (_276_v19).Contains(_255_v10) {
						return (_276_v19).Get(_255_v10).(_dafny.Int)
					}
					return (_this.F25).F24()
				})())
				r1 = (_277_v20).IsDisjointFrom((_277_v20).Intersection(_dafny.MultiSetOf((_270_cf18).Cardinality(), (_this.F25).F24())))
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index37
				((_this).F19()).ArraySet1((_this).F20(), (_index37).Int())
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index38
				((_this).F19()).ArraySet1(!((_this).F20()), (_index38).Int())
				var _278_v21 _dafny.Set
				_ = _278_v21
				_278_v21 = _dafny.SetOf(_dafny.Companion_Sequence_.Contains(_259_v14, _255_v10), ((_this).F20()) && (((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)))
				_278_v21 = _dafny.SetOf(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), (_239_v0).Select((Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_239_v0).Cardinality()))).Uint32()).(bool), true)
			} else if _source5.Is_DC11() {
				var _279___mcc_h7 _dafny.Int = _source5.Get_().(D3_DC11).Cf21
				_ = _279___mcc_h7
				var _280___mcc_h8 _dafny.Int = _source5.Get_().(D3_DC11).Cf22
				_ = _280___mcc_h8
				var _281___mcc_h9 bool = _source5.Get_().(D3_DC11).Cf23
				_ = _281___mcc_h9
				var _282_cf23 bool = _281___mcc_h9
				_ = _282_cf23
				var _283_cf22 _dafny.Int = _280___mcc_h8
				_ = _283_cf22
				var _284_cf21 _dafny.Int = _279___mcc_h7
				_ = _284_cf21
				_283_cf22 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_239_v0, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_284_cf21, (_this.F25).F24()), _dafny.IntOfUint32((_239_v0).Cardinality()))).Uint32(), (_283_cf22).Cmp((_this.F25).F24()) <= 0)).Cardinality())
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index39
				((_this).F19()).ArraySet1(((_this).F20()) && (_this.F26), (_index39).Int())
				var _285_v22 *C1
				_ = _285_v22
				var _nw49 *C1 = New_C1_()
				_ = _nw49
				_nw49.Ctor__(_259_v14, (_this).F19(), _282_cf23)
				_285_v22 = _nw49
				var _286_v23 _dafny.Sequence
				_ = _286_v23
				_286_v23 = _dafny.SeqOf(_285_v22, _285_v22)
				var _287_v24 _dafny.Array
				_ = _287_v24
				var _nwElement0_7 _dafny.Int = (_this.F25).F24()
				_ = _nwElement0_7
				var _nw50 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(26))
				_ = _nw50
				(_nw50).ArraySet1(_nwElement0_7, 0)
				(_nw50).ArraySet1((_this.F25).F24(), 1)
				(_nw50).ArraySet1(_dafny.IntOfInt64(858), 2)
				(_nw50).ArraySet1(_dafny.IntOfUint32((_259_v14).Cardinality()), 3)
				(_nw50).ArraySet1(_dafny.IntOfInt64(46), 4)
				(_nw50).ArraySet1((_this.F25).F24(), 5)
				(_nw50).ArraySet1(_284_cf21, 6)
				(_nw50).ArraySet1((_this.F25).F24(), 7)
				(_nw50).ArraySet1(_dafny.IntOfUint32((_259_v14).Cardinality()), 8)
				(_nw50).ArraySet1((_this.F25).F24(), 9)
				(_nw50).ArraySet1((_dafny.MultiSetOf(_255_v10, _255_v10)).Cardinality(), 10)
				(_nw50).ArraySet1((_this.F25).F24(), 11)
				(_nw50).ArraySet1(_dafny.IntOfUint32((_286_v23).Cardinality()), 12)
				(_nw50).ArraySet1(_dafny.IntOfInt64(479), 13)
				(_nw50).ArraySet1((_this.F25).F24(), 14)
				(_nw50).ArraySet1(_dafny.IntOfInt64(708), 15)
				(_nw50).ArraySet1((_this.F25).F24(), 16)
				(_nw50).ArraySet1(_283_cf22, 17)
				(_nw50).ArraySet1(_dafny.IntOfInt64(777), 18)
				(_nw50).ArraySet1(_284_cf21, 19)
				(_nw50).ArraySet1((_this.F25).F24(), 20)
				(_nw50).ArraySet1((_this.F25).F24(), 21)
				(_nw50).ArraySet1(_dafny.IntOfUint32((_239_v0).Cardinality()), 22)
				(_nw50).ArraySet1(_283_cf22, 23)
				(_nw50).ArraySet1(_dafny.IntOfInt64(29), 24)
				(_nw50).ArraySet1((_this.F25).F24(), 25)
				_287_v24 = _nw50
				var _288_v25 D0
				_ = _288_v25
				_288_v25 = Companion_D0_.Create_DC0_(_256_v11)
				var _289_v26 _dafny.Map
				_ = _289_v26
				_289_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F25).F24(), _288_v25)
				var _290_v27 _dafny.Map
				_ = _290_v27
				_290_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_287_v24, _289_v26)
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index40
				((_this).F19()).ArraySet1(!(_290_v27).Equals(_290_v27), (_index40).Int())
				var _291_v28 _dafny.Sequence
				_ = _291_v28
				_291_v28 = _dafny.SeqOf(Companion_Default___.Fm8(_259_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(944), _this.F26), (_this.F25).F24(), globalState))
				var _292_v29 *C0
				_ = _292_v29
				var _nw51 *C0 = New_C0_()
				_ = _nw51
				_nw51.Ctor__(_dafny.IntOfUint32((_291_v28).Cardinality()))
				_292_v29 = _nw51
				var _293_v30 D4
				_ = _293_v30
				_293_v30 = Companion_D4_.Create_DC12_(_255_v10)
				_255_v10 = (_293_v30).Dtor_cf24()
			} else {
				var _294___mcc_h10 _dafny.Sequence = _source5.Get_().(D3_DC9).Cf16
				_ = _294___mcc_h10
				var _295_cf16 _dafny.Sequence = _294___mcc_h10
				_ = _295_cf16
				var _296_v31 *C1
				_ = _296_v31
				var _nw52 *C1 = New_C1_()
				_ = _nw52
				_nw52.Ctor__(_259_v14, (_this).F19(), ((_this.F25).F24()).Cmp((_this.F25).F24()) > 0)
				_296_v31 = _nw52
				_296_v31 = _296_v31
				var _297_v32 _dafny.Map
				_ = _297_v32
				_297_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v31, _this.F26)
				var _298_v33 _dafny.Map
				_ = _298_v33
				_298_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_297_v32).Cardinality(), _this.F26)
				var _299_v34 _dafny.Map
				_ = _299_v34
				_299_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_298_v33).Cardinality(), (_this.F25).F24())
				_299_v34 = (_299_v34).Update((_this.F25).F24(), ((_this.F25).F24()).Plus((func() _dafny.Int {
					if (_256_v11).Contains(_this.F26) {
						return (_256_v11).Multiplicity(_this.F26)
					}
					return _dafny.IntOfUint32(((_296_v31).F27()).Cardinality())
				})()))
				(globalState).F0 = _this.F26
				var _300_v35 _dafny.Set
				_ = _300_v35
				_300_v35 = _dafny.SetOf((_this.F25).F24())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index41
				((_this).F19()).ArraySet1(!((_dafny.IntOfInt64(670)).Cmp((func() _dafny.Int {
					if (_256_v11).Contains((_this).F20()) {
						return (_256_v11).Multiplicity((_this).F20())
					}
					return (_300_v35).Cardinality()
				})()) != 0), (_index41).Int())
				var _301_v36 _dafny.Array
				_ = _301_v36
				var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw53
				_301_v36 = _nw53
				var _302_v37 _dafny.Sequence
				_ = _302_v37
				_302_v37 = _dafny.SeqOf((_this.F25).F24())
				var _303_v38 _dafny.Map
				_ = _303_v38
				_303_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_this).F20() {
						return _dafny.IntOfUint32((_302_v37).Cardinality())
					}
					return (_this.F25).F24()
				})(), _301_v36)
				var _304_v39 _dafny.Sequence
				_ = _304_v39
				_304_v39 = _dafny.SeqOf(_259_v14)
				var _305_v40 _dafny.Map
				_ = _305_v40
				_305_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_304_v39, _301_v36)
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index42
				var _rhs17 _dafny.Int = Companion_Default___.Fm0((_256_v11).IsSubsetOf(_256_v11), _this.F26, globalState)
				_ = _rhs17
				var _rhs18 bool = (_this).F20()
				_ = _rhs18
				var _rhs19 _dafny.Int = (_this.F25).F24()
				_ = _rhs19
				var _rhs20 _dafny.Array = (func() _dafny.Array {
					if (_303_v38).Contains((_this.F25).F24()) {
						return (_303_v38).Get((_this.F25).F24()).(_dafny.Array)
					}
					return (func() _dafny.Array {
						if (_305_v40).Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("mti"))) {
							return (_305_v40).Get(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("mti"))).(_dafny.Array)
						}
						return _301_v36
					})()
				})()
				_ = _rhs20
				var _lhs16 *GlobalState = globalState
				_ = _lhs16
				var _lhs17 _dafny.Array = (_this).F19()
				_ = _lhs17
				var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _lhs18
				var _lhs19 *GlobalState = globalState
				_ = _lhs19
				_lhs16.F18 = _rhs17
				(_lhs17).ArraySet1(_rhs18, (_lhs18).Int())
				_lhs19.F18 = _rhs19
				_301_v36 = _rhs20
			}
			(globalState).F0 = _this.F26
		} else {
			var _306___mcc_h2 D1 = _source4.Get_().(D1_DC6).Cf12
			_ = _306___mcc_h2
			var _307_cf12 D1 = _306___mcc_h2
			_ = _307_cf12
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))
			_ = _index43
			((_this).F19()).ArraySet1((_this).F20(), (_index43).Int())
			var _308_v41 _dafny.Sequence
			_ = _308_v41
			_308_v41 = _dafny.SeqOf((_this.F25).F24())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))
			_ = _index44
			((_this).F19()).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_308_v41, _308_v41), (_index44).Int())
			var _309_v42 _dafny.Map
			_ = _309_v42
			_309_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _this.F25)
			var _310_v43 _dafny.Map
			_ = _310_v43
			_310_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F26, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _this.F25)).Update((_this).F20(), _this.F25))
			var _311_v44 _dafny.Array
			_ = _311_v44
			var _nwElement0_8 _dafny.Map = _309_v42
			_ = _nwElement0_8
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(28))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_8, 0)
			(_nw54).ArraySet1(_309_v42, 1)
			(_nw54).ArraySet1(_309_v42, 2)
			(_nw54).ArraySet1(_309_v42, 3)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F25), 4)
			(_nw54).ArraySet1(_309_v42, 5)
			(_nw54).ArraySet1(_309_v42, 6)
			(_nw54).ArraySet1(_309_v42, 7)
			(_nw54).ArraySet1(_309_v42, 8)
			(_nw54).ArraySet1(_309_v42, 9)
			(_nw54).ArraySet1(_309_v42, 10)
			(_nw54).ArraySet1(_309_v42, 11)
			(_nw54).ArraySet1(_309_v42, 12)
			(_nw54).ArraySet1((_309_v42).Update(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _this.F25), 13)
			(_nw54).ArraySet1(_309_v42, 14)
			(_nw54).ArraySet1(_309_v42, 15)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F25), 16)
			(_nw54).ArraySet1(_309_v42, 17)
			(_nw54).ArraySet1((func() _dafny.Map {
				if (_310_v43).Contains((_this).F20()) {
					return (_310_v43).Get((_this).F20()).(_dafny.Map)
				}
				return _309_v42
			})(), 18)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _this.F25), 19)
			(_nw54).ArraySet1(_309_v42, 20)
			(_nw54).ArraySet1(_309_v42, 21)
			(_nw54).ArraySet1(_309_v42, 22)
			(_nw54).ArraySet1(_309_v42, 23)
			(_nw54).ArraySet1(_309_v42, 24)
			(_nw54).ArraySet1(_309_v42, 25)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_239_v0).Select((Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_239_v0).Cardinality()))).Uint32()).(bool), _this.F25), 26)
			(_nw54).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F25), 27)
			_311_v44 = _nw54
			var _312_v45 _dafny.Map
			_ = _312_v45
			_312_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), Companion_Default___.Fm0(_this.F26, ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), globalState))
			var _313_v46 _dafny.Sequence
			_ = _313_v46
			_313_v46 = _dafny.SeqOf(_312_v45)
			var _314_v47 D3
			_ = _314_v47
			_314_v47 = Companion_D3_.Create_DC10_(_311_v44, (_313_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-519), _dafny.IntOfUint32((_313_v46).Cardinality()))).Uint32()).(_dafny.Map), _dafny.SetOf(_241_v2), (_dafny.IntOfInt64(932)).Minus((_this.F25).F24()))
			var _315_v48 _dafny.CodePoint
			_ = _315_v48
			_315_v48 = _dafny.CodePoint('l')
			var _316_v49 _dafny.Sequence
			_ = _316_v49
			_316_v49 = _dafny.SeqOf((func() _dafny.CodePoint {
				if _this.F26 {
					return _315_v48
				}
				return _315_v48
			})(), _315_v48)
			var _317_v50 _dafny.MultiSet
			_ = _317_v50
			_317_v50 = _dafny.MultiSetOf(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool))
			var _318_v51 _dafny.Set
			_ = _318_v51
			_318_v51 = _dafny.SetOf((_317_v50).Intersection(_317_v50))
			var _319_v52 _dafny.MultiSet
			_ = _319_v52
			_319_v52 = _dafny.MultiSetOf(_241_v2, _241_v2)
			var _320_v53 _dafny.MultiSet
			_ = _320_v53
			_320_v53 = _dafny.MultiSetOf((_dafny.Zero).Minus((_this.F25).F24()), (_this.F25).F24())
			var _321_v54 _dafny.Array
			_ = _321_v54
			var _nwElement0_9 _dafny.Int = (func() _dafny.Int {
				if (_319_v52).Contains(_241_v2) {
					return (_319_v52).Multiplicity(_241_v2)
				}
				return (_this.F25).F24()
			})()
			_ = _nwElement0_9
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(24))
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_9, 0)
			(_nw55).ArraySet1((_dafny.Zero).Minus((_this.F25).F24()), 1)
			(_nw55).ArraySet1(Companion_Default___.SafeModuloInt((_this.F25).F24(), (_this.F25).F24()), 2)
			(_nw55).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this.F25).F24()), (_dafny.Zero).Minus((_this.F25).F24()))), 3)
			(_nw55).ArraySet1((_this.F25).F24(), 4)
			(_nw55).ArraySet1((_this.F25).F24(), 5)
			(_nw55).ArraySet1(_dafny.IntOfInt64(40), 6)
			(_nw55).ArraySet1(_dafny.IntOfInt64(670), 7)
			(_nw55).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this.F25).F24()), (_this.F25).F24()), 8)
			(_nw55).ArraySet1(_dafny.IntOfUint32((_316_v49).Cardinality()), 9)
			(_nw55).ArraySet1(_dafny.IntOfInt64(337), 10)
			(_nw55).ArraySet1((_this.F25).F24(), 11)
			(_nw55).ArraySet1(((_this.F25).F24()).Plus((_this.F25).F24()), 12)
			(_nw55).ArraySet1((_this.F25).F24(), 13)
			(_nw55).ArraySet1(_dafny.IntOfUint32((_316_v49).Cardinality()), 14)
			(_nw55).ArraySet1((_this.F25).F24(), 15)
			(_nw55).ArraySet1((func() _dafny.Int {
				if ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool) {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mxgapglw")).Cardinality())
				}
				return (_this.F25).F24()
			})(), 16)
			(_nw55).ArraySet1((_this.F25).F24(), 17)
			(_nw55).ArraySet1((func() _dafny.Int {
				if _this.F26 {
					return _dafny.IntOfInt64(611)
				}
				return (_dafny.Zero).Minus(Companion_Default___.Fm0(_this.F26, _this.F26, globalState))
			})(), 18)
			(_nw55).ArraySet1(((_this.F25).F24()).Times((_this.F25).F24()), 19)
			(_nw55).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("x")).Cardinality()), 20)
			(_nw55).ArraySet1((_dafny.IntOfInt64(283)).Minus((_this.F25).F24()), 21)
			(_nw55).ArraySet1((_this.F25).F24(), 22)
			(_nw55).ArraySet1(((_317_v50).Cardinality()).Plus((_320_v53).Cardinality()), 23)
			_321_v54 = _nw55
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_321_v54), 0))
			_ = _index45
			(_321_v54).ArraySet1((_this.F25).F24(), (_index45).Int())
			var _322_v55 _dafny.Sequence
			_ = _322_v55
			_322_v55 = _dafny.SeqOf(_311_v44)
			var _323_v56 D0
			_ = _323_v56
			_323_v56 = Companion_D0_.Create_DC2_((_this).F20(), _dafny.IntOfInt64(682), _239_v0, _dafny.IntOfUint32((_316_v49).Cardinality()), (_this.F25).F24())
			var _324_v57 _dafny.Map
			_ = _324_v57
			_324_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_323_v56).Dtor_cf7(), ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool))
			var _325_v58 _dafny.Map
			_ = _325_v58
			_325_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_324_v57).Cardinality(), _dafny.SeqOf(_315_v48, _315_v48, _315_v48))
			var _pat_let_tv9 = _312_v45
			_ = _pat_let_tv9
			var _pat_let_tv10 = _322_v55
			_ = _pat_let_tv10
			var _pat_let_tv11 = _322_v55
			_ = _pat_let_tv11
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_321_v54), 0))
			_ = _index46
			var _rhs21 D3 = func(_pat_let4_0 D3) D3 {
				return func(_326_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let5_0 _dafny.Map) D3 {
						return func(_327_dt__update_hcf18_h0 _dafny.Map) D3 {
							return func(_pat_let6_0 _dafny.Array) D3 {
								return func(_328_dt__update_hcf17_h0 _dafny.Array) D3 {
									return Companion_D3_.Create_DC10_(_328_dt__update_hcf17_h0, _327_dt__update_hcf18_h0, (_326_dt__update__tmp_h1).Dtor_cf19(), (_326_dt__update__tmp_h1).Dtor_cf20())
								}(_pat_let6_0)
							}((_pat_let_tv10).Select((Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_pat_let_tv11).Cardinality()))).Uint32()).(_dafny.Array))
						}(_pat_let5_0)
					}(_pat_let_tv9)
				}(_pat_let4_0)
			}(_314_v47)
			_ = _rhs21
			var _rhs22 _dafny.Sequence = (func() _dafny.Sequence {
				if (_325_v58).Contains(_dafny.IntOfUint32((_239_v0).Cardinality())) {
					return (_325_v58).Get(_dafny.IntOfUint32((_239_v0).Cardinality())).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('e')), _dafny.SeqOf(_315_v48, _315_v48))
			})()
			_ = _rhs22
			var _rhs23 _dafny.Set = _318_v51
			_ = _rhs23
			var _rhs24 _dafny.Int = (_dafny.Zero).Minus((_this.F25).F24())
			_ = _rhs24
			var _lhs20 _dafny.Array = _321_v54
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_321_v54), 0))
			_ = _lhs21
			_314_v47 = _rhs21
			_316_v49 = _rhs22
			_318_v51 = _rhs23
			(_lhs20).ArraySet1(_rhs24, (_lhs21).Int())
			var _329_v59 D3
			_ = _329_v59
			_329_v59 = Companion_D3_.Create_DC11_((_this.F25).F24(), (_this.F25).F24(), !(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)) || (((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)))
			var _source6 D3 = _329_v59
			_ = _source6
			if _source6.Is_DC10() {
				var _330___mcc_h11 _dafny.Array = _source6.Get_().(D3_DC10).Cf17
				_ = _330___mcc_h11
				var _331___mcc_h12 _dafny.Map = _source6.Get_().(D3_DC10).Cf18
				_ = _331___mcc_h12
				var _332___mcc_h13 _dafny.Set = _source6.Get_().(D3_DC10).Cf19
				_ = _332___mcc_h13
				var _333___mcc_h14 _dafny.Int = _source6.Get_().(D3_DC10).Cf20
				_ = _333___mcc_h14
				var _334_cf20 _dafny.Int = _333___mcc_h14
				_ = _334_cf20
				var _335_cf19 _dafny.Set = _332___mcc_h13
				_ = _335_cf19
				var _336_cf18 _dafny.Map = _331___mcc_h12
				_ = _336_cf18
				var _337_cf17 _dafny.Array = _330___mcc_h11
				_ = _337_cf17
				var _338_v60 D4
				_ = _338_v60
				_338_v60 = Companion_D4_.Create_DC12_(_315_v48)
				var _pat_let_tv12 = _315_v48
				_ = _pat_let_tv12
				var _339_v61 _dafny.Map
				_ = _339_v61
				_339_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let7_0 D4) D4 {
					return func(_340_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let8_0 _dafny.CodePoint) D4 {
							return func(_341_dt__update_hcf24_h0 _dafny.CodePoint) D4 {
								return Companion_D4_.Create_DC12_(_341_dt__update_hcf24_h0)
							}(_pat_let8_0)
						}(_pat_let_tv12)
					}(_pat_let7_0)
				}(_338_v60), (_this.F25).F24())
				_339_v61 = (_339_v61).Update(_338_v60, Companion_Default___.SafeModuloInt((_this.F25).F24(), (_this.F25).F24()))
				(globalState).F0 = !((func() bool {
					if (_324_v57).Contains((_dafny.Zero).Minus((_dafny.IntOfUint32((_239_v0).Cardinality())).Plus((_dafny.MultiSetOf((_this.F25).F24())).Cardinality()))) {
						return (_324_v57).Get((_dafny.Zero).Minus((_dafny.IntOfUint32((_239_v0).Cardinality())).Plus((_dafny.MultiSetOf((_this.F25).F24())).Cardinality()))).(bool)
					}
					return (_this).F20()
				})())
				(_this).F26 = ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)
				var _342_v62 _dafny.Set
				_ = _342_v62
				_342_v62 = _dafny.SetOf(_this.F26, !((_317_v50).Update(_this.F26, Companion_Default___.Abs((_this.F25).F24()))).Equals(_317_v50), (_this).F20(), ((_this.F25).F24()).Cmp((_this.F25).F24()) >= 0)
				var _343_v63 D4
				_ = _343_v63
				_343_v63 = Companion_D4_.Create_DC14_(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _342_v62)
				var _pat_let_tv13 = _342_v62
				_ = _pat_let_tv13
				_342_v62 = (func(_pat_let9_0 D4) D4 {
					return func(_344_dt__update__tmp_h3 D4) D4 {
						return func(_pat_let10_0 _dafny.Set) D4 {
							return func(_345_dt__update_hcf26_h0 _dafny.Set) D4 {
								return Companion_D4_.Create_DC14_((_344_dt__update__tmp_h3).Dtor_cf25(), _345_dt__update_hcf26_h0)
							}(_pat_let10_0)
						}(_pat_let_tv13)
					}(_pat_let9_0)
				}(_343_v63)).Dtor_cf26()
			} else if _source6.Is_DC11() {
				var _346___mcc_h15 _dafny.Int = _source6.Get_().(D3_DC11).Cf21
				_ = _346___mcc_h15
				var _347___mcc_h16 _dafny.Int = _source6.Get_().(D3_DC11).Cf22
				_ = _347___mcc_h16
				var _348___mcc_h17 bool = _source6.Get_().(D3_DC11).Cf23
				_ = _348___mcc_h17
				var _349_cf23 bool = _348___mcc_h17
				_ = _349_cf23
				var _350_cf22 _dafny.Int = _347___mcc_h16
				_ = _350_cf22
				var _351_cf21 _dafny.Int = _346___mcc_h15
				_ = _351_cf21
				var _352_v64 _dafny.Array
				_ = _352_v64
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_7
				var _nw56 _dafny.Array
				_ = _nw56
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw56 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = (func(_353_cf22 _dafny.Int) func(_dafny.Int) bool {
						return func(_354_i2 _dafny.Int) bool {
							return ((_this.F25).F24()).Cmp(_353_cf22) > 0
						}
					})(_350_cf22)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw56 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw56).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw56).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_352_v64 = _nw56
				_352_v64 = _352_v64
				(globalState).F18 = (_321_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_321_v54), 0))).Int()).(_dafny.Int)
				var _355_v65 *C1
				_ = _355_v65
				var _nw57 *C1 = New_C1_()
				_ = _nw57
				_nw57.Ctor__(_316_v49, _352_v64, _this.F26)
				_355_v65 = _nw57
				var _356_v66 _dafny.Map
				_ = _356_v66
				_356_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_355_v65).F27(), _this.F26)
				_356_v66 = (_356_v66).Update((_355_v65).F27(), !((_this).F20()))
			} else {
				var _357___mcc_h18 _dafny.Sequence = _source6.Get_().(D3_DC9).Cf16
				_ = _357___mcc_h18
				var _358_cf16 _dafny.Sequence = _357___mcc_h18
				_ = _358_cf16
				var _359_v68 _dafny.Array
				_ = _359_v68
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_8
				var _nw58 _dafny.Array
				_ = _nw58
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw58 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) bool = (func(_360_v48 _dafny.CodePoint) func(_dafny.Int) bool {
						return func(_361_i3 _dafny.Int) bool {
							return (_dafny.SetOf(_360_v48)).IsDisjointFrom(func() _dafny.Set {
								var _coll10 = _dafny.NewBuilder()
								_ = _coll10
								for _iter11 := _dafny.Iterate((_dafny.SetOf(_360_v48, _360_v48)).Elements()); ; {
									_compr_10, _ok11 := _iter11()
									if !_ok11 {
										break
									}
									var _362_v67 _dafny.CodePoint
									_362_v67 = interface{}(_compr_10).(_dafny.CodePoint)
									if (_dafny.SetOf(_360_v48, _360_v48)).Contains(_362_v67) {
										_coll10.Add(_362_v67)
									}
								}
								return _coll10.ToSet()
							}())
						}
					})(_315_v48)
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
				_359_v68 = _nw58
				_359_v68 = _359_v68
				(globalState).F10 = _321_v54
				var _363_v69 *C0
				_ = _363_v69
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(((_this.F25).F24()).Minus((_this.F25).F24()))
				_363_v69 = _nw59
				var _364_v70 _dafny.Map
				_ = _364_v70
				_364_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F26)
				_364_v70 = _364_v70
			}
			var _365_v71 bool
			_ = _365_v71
			var _366_v72 _dafny.Int
			_ = _366_v72
			var _367_v73 _dafny.Int
			_ = _367_v73
			var _368_v74 _dafny.Int
			_ = _368_v74
			var _out16 bool
			_ = _out16
			var _out17 _dafny.Int
			_ = _out17
			var _out18 _dafny.Int
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			_out16, _out17, _out18, _out19 = (_this).M4(!((_this).F20()), globalState)
			_365_v71 = _out16
			_366_v72 = _out17
			_367_v73 = _out18
			_368_v74 = _out19
		}
		var _369_v75 _dafny.Sequence
		_ = _369_v75
		_369_v75 = _dafny.UnicodeSeqOfUtf8Bytes("eneiyaye")
		_369_v75 = _dafny.UnicodeSeqOfUtf8Bytes("rbiwky")
		if _this.F26 {
			var _370_v76 _dafny.Map
			_ = _370_v76
			_370_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(955), (_this.F25).F24())
			var _371_v77 _dafny.MultiSet
			_ = _371_v77
			_371_v77 = _dafny.MultiSetOf((_this.F25).F24(), (func() _dafny.Int {
				if (_370_v76).Contains((_this.F25).F24()) {
					return (_370_v76).Get((_this.F25).F24()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_369_v75).Cardinality())
			})(), _dafny.IntOfInt64(403), (_this.F25).F24())
			var _372_v78 _dafny.Map
			_ = _372_v78
			_372_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_371_v77).Update((_dafny.MultiSetFromSeq(_239_v0)).Cardinality(), Companion_Default___.Abs(_dafny.IntOfUint32((_239_v0).Cardinality()))))
			_371_v77 = ((func() _dafny.MultiSet {
				if (_372_v78).Contains(_this.F26) {
					return (_372_v78).Get(_this.F26).(_dafny.MultiSet)
				}
				return _371_v77
			})()).Difference(_371_v77)
			var _373_v79 *C1
			_ = _373_v79
			var _nw60 *C1 = New_C1_()
			_ = _nw60
			_nw60.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(276))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_374_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})), (_this).F19(), _this.F26)
			_373_v79 = _nw60
			var _375_v80 _dafny.Sequence
			_ = _375_v80
			_375_v80 = _dafny.SeqOf(_373_v79)
			var _376_v81 _dafny.Map
			_ = _376_v81
			_376_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F26), (_this).F20())
			var _377_v82 D0
			_ = _377_v82
			_377_v82 = Companion_D0_.Create_DC2_(Companion_Default___.Fm9((_this).F20(), globalState), (_376_v81).Cardinality(), _dafny.SeqOf((_this).F20(), _this.F26), (_this.F25).F24(), (_this.F25).F24())
			var _378_v83 _dafny.Array
			_ = _378_v83
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw61
			_378_v83 = _nw61
			var _pat_let_tv14 = _239_v0
			_ = _pat_let_tv14
			var _379_v84 _dafny.Map
			_ = _379_v84
			_379_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_375_v80).Cardinality())).Minus((func(_pat_let11_0 D0) D0 {
				return func(_380_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let12_0 _dafny.Int) D0 {
						return func(_381_dt__update_hcf5_h0 _dafny.Int) D0 {
							return func(_pat_let13_0 bool) D0 {
								return func(_382_dt__update_hcf4_h0 bool) D0 {
									return func(_pat_let14_0 _dafny.Sequence) D0 {
										return func(_383_dt__update_hcf6_h0 _dafny.Sequence) D0 {
											return Companion_D0_.Create_DC2_(_382_dt__update_hcf4_h0, _381_dt__update_hcf5_h0, _383_dt__update_hcf6_h0, (_380_dt__update__tmp_h4).Dtor_cf7(), (_380_dt__update__tmp_h4).Dtor_cf8())
										}(_pat_let14_0)
									}(_pat_let_tv14)
								}(_pat_let13_0)
							}(true)
						}(_pat_let12_0)
					}((_this.F25).F24())
				}(_pat_let11_0)
			}(_377_v82)).Dtor_cf7()), _378_v83)
			var _384_v85 D3
			_ = _384_v85
			_384_v85 = Companion_D3_.Create_DC9_(_369_v75)
			var _385_v86 _dafny.MultiSet
			_ = _385_v86
			_385_v86 = _dafny.MultiSetOf(Companion_D3_.Create_DC9_((_373_v79).F27()), _384_v85, Companion_D3_.Create_DC9_(_369_v75), _384_v85, _384_v85)
			_379_v84 = (_379_v84).Update((_385_v86).Cardinality(), _378_v83)
			var _386_v87 _dafny.MultiSet
			_ = _386_v87
			_386_v87 = _dafny.MultiSetOf((_this).F20())
			_386_v87 = _386_v87
			if !((_this.F26) || ((_this.F26) || (!((_this).F20())))) {
				(globalState).F0 = (_this).F20()
				var _387_v88 _dafny.Set
				_ = _387_v88
				_387_v88 = _dafny.SetOf((_dafny.Zero).Minus((_this.F25).F24()))
				var _rhs25 bool = (_387_v88).IsProperSubsetOf((_387_v88).Difference(_387_v88))
				_ = _rhs25
				var _rhs26 _dafny.Int = (_this.F25).F24()
				_ = _rhs26
				var _lhs22 *C2 = _this
				_ = _lhs22
				var _lhs23 *GlobalState = globalState
				_ = _lhs23
				_lhs22.F26 = _rhs25
				_lhs23.F13 = _rhs26
				var _388_v89 D2
				_ = _388_v89
				_388_v89 = Companion_D2_.Create_DC8_((_this.F25).F24(), Companion_Default___.Fm0(_this.F26, _this.F26, globalState))
				var _rhs27 _dafny.Int = (_388_v89).Dtor_cf15()
				_ = _rhs27
				var _rhs28 _dafny.Array = (func() _dafny.Array {
					if (_this).F20() {
						return _378_v83
					}
					return _378_v83
				})()
				_ = _rhs28
				r3 = _rhs27
				_378_v83 = _rhs28
				(globalState).F18 = (_this.F25).F24()
				var _389_v90 _dafny.Map
				_ = _389_v90
				_389_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F25).F24(), (_this).F20())
				(globalState).F18 = Companion_Default___.SafeDivisionInt((_this.F25).F24(), (Companion_Default___.Fm12(_371_v77, ((_389_v90).Update(_dafny.IntOfUint32(((_373_v79).F27()).Cardinality()), (_this).F20())).Cardinality(), globalState)).Cardinality())
			} else {
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index47
				((_this).F19()).ArraySet1(_this.F26, (_index47).Int())
				var _390_v91 D1
				_ = _390_v91
				_390_v91 = Companion_D1_.Create_DC5_(true)
				var _391_v92 _dafny.CodePoint
				_ = _391_v92
				_391_v92 = _dafny.CodePoint('o')
				var _392_v93 T0
				_ = _392_v93
				var _nw62 *C1 = New_C1_()
				_ = _nw62
				_nw62.Ctor__(_dafny.Companion_Sequence_.Update((_373_v79).F27(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this.F25).F24()), _dafny.IntOfUint32(((_373_v79).F27()).Cardinality()))).Uint32(), _391_v92), (_this).F19(), _this.F26)
				_392_v93 = _nw62
				var _393_v94 _dafny.Sequence
				_ = _393_v94
				_393_v94 = _dafny.SeqOf(_392_v93, _392_v93)
				var _394_v95 _dafny.MultiSet
				_ = _394_v95
				_394_v95 = _dafny.MultiSetOf(_392_v93, (_393_v94).Select((Companion_Default___.SafeIndex((_this.F25).F24(), _dafny.IntOfUint32((_393_v94).Cardinality()))).Uint32()).(T0), _392_v93)
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index48
				var _rhs29 bool = ((_394_v95).Cardinality()).Cmp(((_this.F25).F24()).Times((_this.F25).F24())) != 0
				_ = _rhs29
				var _rhs30 _dafny.Int = (_this.F25).F24()
				_ = _rhs30
				var _rhs31 D1 = (func() D1 {
					if _this.F26 {
						return Companion_D1_.Create_DC5_((_392_v93).F20())
					}
					return _390_v91
				})()
				_ = _rhs31
				var _lhs24 _dafny.Array = (_this).F19()
				_ = _lhs24
				var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _lhs25
				(_lhs24).ArraySet1(_rhs29, (_lhs25).Int())
				r3 = _rhs30
				_390_v91 = _rhs31
				r3 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_369_v75).Cardinality()), (_dafny.Zero).Minus((_this.F25).F24()))).Minus((_this.F25).F24())
				var _395_v96 _dafny.Set
				_ = _395_v96
				_395_v96 = _dafny.SetOf((_this).F20())
				var _396_v97 *C0
				_ = _396_v97
				var _nw63 *C0 = New_C0_()
				_ = _nw63
				_nw63.Ctor__(Companion_Default___.SafeDivisionInt((_395_v96).Cardinality(), (_this.F25).F24()))
				_396_v97 = _nw63
				var _397_v98 _dafny.Array
				_ = _397_v98
				var _nw64 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(26))
				_ = _nw64
				_397_v98 = _nw64
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_397_v98), 0))
				_ = _index49
				(_397_v98).ArraySet1(Companion_D1_.Create_DC4_(_240_v1), (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_397_v98), 0))
				_ = _index50
				(_397_v98).ArraySet1(_241_v2, (_index50).Int())
				(globalState).F10 = _378_v83
			}
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F19()), 0))
			_ = _index51
			((_this).F19()).ArraySet1((_this).F20(), (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen(((_this).F19()), 0))
			_ = _index52
			((_this).F19()).ArraySet1((_this).F20(), (_index52).Int())
		} else {
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen(((_this).F19()), 0))
			_ = _index53
			((_this).F19()).ArraySet1(true, (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen(((_this).F19()), 0))
			_ = _index54
			((_this).F19()).ArraySet1(_this.F26, (_index54).Int())
			var _398_v99 _dafny.Array
			_ = _398_v99
			var _nwElement0_10 _dafny.Int = (_dafny.Zero).Minus((_this.F25).F24())
			_ = _nwElement0_10
			var _nw65 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(10))
			_ = _nw65
			(_nw65).ArraySet1(_nwElement0_10, 0)
			(_nw65).ArraySet1((_this.F25).F24(), 1)
			(_nw65).ArraySet1((_this.F25).F24(), 2)
			(_nw65).ArraySet1(_dafny.IntOfInt64(226), 3)
			(_nw65).ArraySet1((_dafny.MultiSetOf((_this.F25).F24())).Cardinality(), 4)
			(_nw65).ArraySet1((_this.F25).F24(), 5)
			(_nw65).ArraySet1(_dafny.IntOfInt64(67), 6)
			(_nw65).ArraySet1(_dafny.IntOfInt64(418), 7)
			(_nw65).ArraySet1((_this.F25).F24(), 8)
			(_nw65).ArraySet1((_this.F25).F24(), 9)
			_398_v99 = _nw65
			(globalState).F10 = (Companion_D5_.Create_DC16_(_398_v99)).Dtor_cf28()
			var _399_v100 _dafny.Array
			_ = _399_v100
			var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw66
			_399_v100 = _nw66
			var _400_v101 _dafny.Array
			_ = _400_v101
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_9
			var _nw67 _dafny.Array
			_ = _nw67
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw67 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) D0 = func(_401_i5 _dafny.Int) D0 {
					return Companion_D0_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), (_this).F20(), (_this).F20(), ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)), _dafny.IntOfInt64(118)))
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw67 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw67).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw67).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_400_v101 = _nw67
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_399_v100), 0))
			_ = _index55
			(_399_v100).ArraySet1(_400_v101, (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_399_v100), 0))
			_ = _index56
			(_399_v100).ArraySet1(_400_v101, (_index56).Int())
			var _402_v102 _dafny.Array
			_ = _402_v102
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_10
			var _nw68 _dafny.Array
			_ = _nw68
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw68 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Map = func(_403_i6 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), (_this).F20())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)))
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
			_402_v102 = _nw68
			_402_v102 = _402_v102
			var _404_v103 _dafny.Array
			_ = _404_v103
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
			_ = _nw69
			_404_v103 = _nw69
			var _405_v104 _dafny.Set
			_ = _405_v104
			_405_v104 = _dafny.SetOf(_241_v2)
			var _406_v105 D3
			_ = _406_v105
			_406_v105 = Companion_D3_.Create_DC10_(_404_v103, Companion_Default___.Fm16(_this.F26, _this.F26, _dafny.IntOfUint32((_369_v75).Cardinality()), globalState), _405_v104, (_dafny.Zero).Minus((_this.F25).F24()))
			r2 = (_dafny.Zero).Minus((_406_v105).Dtor_cf20())
		}
		var _407_v106 _dafny.Array
		_ = _407_v106
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw70
		_407_v106 = _nw70
		(globalState).F3 = _407_v106
		var _408_v107 _dafny.Array
		_ = _408_v107
		var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw71
		_408_v107 = _nw71
		var _409_v108 _dafny.Array
		_ = _409_v108
		var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
		_ = _nw72
		_409_v108 = _nw72
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_408_v107), 0))
		_ = _index57
		(_408_v107).ArraySet1(_409_v108, (_index57).Int())
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_408_v107), 0))
		_ = _index58
		(_408_v107).ArraySet1(_409_v108, (_index58).Int())
		r3 = ((_this.F25).F24()).Minus(((_this.F25).F24()).Plus((_this.F25).F24()))
		r0 = _this.F26
		r1 = (func() bool {
			if false {
				return (_this).F20()
			}
			return _this.F26
		})()
		r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_369_v75, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_410_i7 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})))).Cardinality())
		r3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_369_v75, _dafny.Companion_Sequence_.Concatenate(_369_v75, _369_v75))).Cardinality())
		return r0, r1, r2, r3
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	F22  _dafny.Sequence
	_f23 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this.F22 = _dafny.EmptySeq
	_this._f23 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C3) Ctor__(f22 _dafny.Sequence, f23 _dafny.Array) {
	{
		(_this).F22 = f22
		(_this)._f23 = f23
	}
}
func (_this *C3) Fm5(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var _source7 D0 = Companion_D0_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true)), _dafny.IntOfUint32((_dafny.SeqOf(true, false, true)).Cardinality())))
		_ = _source7
		if _source7.Is_DC1() {
			var _411___mcc_h0 bool = _source7.Get_().(D0_DC1).Cf1
			_ = _411___mcc_h0
			var _412___mcc_h1 _dafny.Int = _source7.Get_().(D0_DC1).Cf2
			_ = _412___mcc_h1
			var _413___mcc_h2 _dafny.Int = _source7.Get_().(D0_DC1).Cf3
			_ = _413___mcc_h2
			var _414_cf3 _dafny.Int = _413___mcc_h2
			_ = _414_cf3
			var _415_cf2 _dafny.Int = _412___mcc_h1
			_ = _415_cf2
			var _416_cf1 bool = _411___mcc_h0
			_ = _416_cf1
			return _dafny.IntOfInt64(543)
		} else if _source7.Is_DC2() {
			var _417___mcc_h3 bool = _source7.Get_().(D0_DC2).Cf4
			_ = _417___mcc_h3
			var _418___mcc_h4 _dafny.Int = _source7.Get_().(D0_DC2).Cf5
			_ = _418___mcc_h4
			var _419___mcc_h5 _dafny.Sequence = _source7.Get_().(D0_DC2).Cf6
			_ = _419___mcc_h5
			var _420___mcc_h6 _dafny.Int = _source7.Get_().(D0_DC2).Cf7
			_ = _420___mcc_h6
			var _421___mcc_h7 _dafny.Int = _source7.Get_().(D0_DC2).Cf8
			_ = _421___mcc_h7
			var _422_cf8 _dafny.Int = _421___mcc_h7
			_ = _422_cf8
			var _423_cf7 _dafny.Int = _420___mcc_h6
			_ = _423_cf7
			var _424_cf6 _dafny.Sequence = _419___mcc_h5
			_ = _424_cf6
			var _425_cf5 _dafny.Int = _418___mcc_h4
			_ = _425_cf5
			var _426_cf4 bool = _417___mcc_h3
			_ = _426_cf4
			return (_dafny.IntOfUint32((_dafny.SeqOf(_426_cf4)).Cardinality())).Minus(_425_cf5)
		} else if _source7.Is_DC3() {
			var _427___mcc_h8 _dafny.Map = _source7.Get_().(D0_DC3).Cf9
			_ = _427___mcc_h8
			var _428_cf9 _dafny.Map = _427___mcc_h8
			_ = _428_cf9
			return _dafny.IntOfInt64(4)
		} else {
			var _429___mcc_h9 _dafny.MultiSet = _source7.Get_().(D0_DC0).Cf0
			_ = _429___mcc_h9
			var _430_cf0 _dafny.MultiSet = _429___mcc_h9
			_ = _430_cf0
			return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(206), _dafny.IntOfInt64(-935))
		}
	}
}
func (_this *C3) Fm6(p0 bool, p1 _dafny.Map, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(13))).Uint32(), func(coer15 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_431_i0 _dafny.Int) D0 {
			return Companion_D0_.Create_DC2_(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(283), !(true))).Cardinality(), _dafny.SeqOf(true), _dafny.IntOfInt64(78), _dafny.IntOfInt64(231))
		})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC2_(!(true), _dafny.IntOfInt64(872), _dafny.SeqOf(false, true), _dafny.IntOfInt64(945), (_dafny.SetOf(true)).Cardinality())), (Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(462))).Uint32(), func(coer16 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_432_i1 _dafny.Int) D0 {
			return Companion_D0_.Create_DC2_(false, (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()))).Cardinality())).Cardinality(), _dafny.SeqOf(false), _dafny.IntOfInt64(-264), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(428))).Cardinality()))
		})))).Dtor_cf10()))).Cardinality())
	}
}
func (_this *C3) M2(globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var _433_v0 _dafny.Int
		_ = _433_v0
		_433_v0 = _dafny.IntOfInt64(419)
		var _434_v1 bool
		_ = _434_v1
		_434_v1 = true
		var _435_v2 _dafny.Set
		_ = _435_v2
		_435_v2 = _dafny.SetOf(false, _434_v1, _434_v1, _434_v1)
		var _436_v3 _dafny.Map
		_ = _436_v3
		_436_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-146), _dafny.SetOf(_434_v1))
		var _437_v4 _dafny.Sequence
		_ = _437_v4
		_437_v4 = _dafny.SeqOf((_435_v2).Cardinality(), ((func() _dafny.Set {
			if (_436_v3).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_435_v2).Cardinality(), _this.F22)).Cardinality()) {
				return (_436_v3).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_435_v2).Cardinality(), _this.F22)).Cardinality()).(_dafny.Set)
			}
			return _435_v2
		})()).Cardinality(), _433_v0, _433_v0)
		var _438_v5 _dafny.Array
		_ = _438_v5
		var _nwElement0_11 _dafny.Int = _433_v0
		_ = _nwElement0_11
		var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(4))
		_ = _nw73
		(_nw73).ArraySet1(_nwElement0_11, 0)
		(_nw73).ArraySet1((_dafny.IntOfInt64(581)).Plus(_433_v0), 1)
		(_nw73).ArraySet1((_437_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_433_v0), _dafny.IntOfUint32((_437_v4).Cardinality()))).Uint32()).(_dafny.Int), 2)
		(_nw73).ArraySet1(_433_v0, 3)
		_438_v5 = _nw73
		for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_438_v5), 0))); ; {
			_guard_loop_1, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _439_i0 _dafny.Int
			_439_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_439_i0).Sign() != -1) && ((_439_i0).Cmp(_dafny.ArrayLen((_438_v5), 0)) < 0)) {
				(_438_v5).ArraySet1(Companion_Default___.SafeDivisionInt(_439_i0, _433_v0), (_439_i0).Int())
			}
		}
		var _440_v6 *C0
		_ = _440_v6
		var _nw74 *C0 = New_C0_()
		_ = _nw74
		_nw74.Ctor__(Companion_Default___.SafeModuloInt(_433_v0, _433_v0))
		_440_v6 = _nw74
		var _441_v7 _dafny.MultiSet
		_ = _441_v7
		_441_v7 = _dafny.MultiSetOf(_434_v1)
		var _442_v8 D0
		_ = _442_v8
		_442_v8 = Companion_D0_.Create_DC0_(_441_v7)
		_442_v8 = Companion_D0_.Create_DC0_(_441_v7)
		var _443_v9 D2
		_ = _443_v9
		_443_v9 = Companion_D2_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_444_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_this.F22).Cardinality())
		})))
		_437_v4 = (_443_v9).Dtor_cf13()
		(globalState).F18 = (_433_v0).Minus(_433_v0)
		r1 = ((_435_v2).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(false, _434_v1, true, _434_v1, _434_v1)).Cardinality())) < 0
		var _445_v10 _dafny.Sequence
		_ = _445_v10
		_445_v10 = _dafny.SeqOf(Companion_Default___.Fm7(_433_v0, (_dafny.Zero).Minus(_433_v0), _433_v0, globalState), _dafny.Companion_Sequence_.IsProperPrefixOf(_437_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(595))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_446_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_447_i2 _dafny.Int) _dafny.Int {
				return _446_v0
			}
		})(_433_v0)))))
		r0 = _445_v10
		r1 = _434_v1
		return r0, r1
	}
}
func (_this *C3) M3(globalState *GlobalState) {
	{
		var _448_v0 bool
		_ = _448_v0
		_448_v0 = true
		var _449_v1 _dafny.Sequence
		_ = _449_v1
		_449_v1 = _dafny.SeqOf(_448_v0, false, _448_v0, _448_v0)
		(globalState).F13 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _448_v0 {
				return _449_v1
			}
			return _449_v1
		})()).Cardinality())
		var _450_v2 _dafny.Array
		_ = _450_v2
		var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(4))
		_ = _nw75
		_450_v2 = _nw75
		for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_450_v2), 0))); ; {
			_guard_loop_2, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _451_i0 _dafny.Int
			_451_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_451_i0).Sign() != -1) && ((_451_i0).Cmp(_dafny.ArrayLen((_450_v2), 0)) < 0)) {
				(_450_v2).ArraySet1CodePoint(_dafny.CodePoint('d'), (_451_i0).Int())
			}
		}
		var _452_v3 T0
		_ = _452_v3
		var _nw76 *C1 = New_C1_()
		_ = _nw76
		_nw76.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("xxxd"), (_this).F23(), _448_v0)
		_452_v3 = _nw76
		var _453_v4 _dafny.Map
		_ = _453_v4
		_453_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_452_v3, (_this).F23())
		_453_v4 = _453_v4
		var _454_v5 _dafny.Array
		_ = _454_v5
		_454_v5 = (_this).F23()
		var _455_v6 _dafny.Set
		_ = _455_v6
		_455_v6 = _dafny.SetOf((_this).F23(), (_this).F23())
		var _456_v7 _dafny.Int
		_ = _456_v7
		_456_v7 = _dafny.IntOfInt64(-719)
		var _457_v8 _dafny.Map
		_ = _457_v8
		_457_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_455_v6).IsProperSubsetOf(_dafny.SetOf((_this).F23(), (_454_v5), (_this).F23(), (_this).F23(), (_452_v3).F19())), _456_v7)
		var _458_v9 _dafny.Sequence
		_ = _458_v9
		_458_v9 = _dafny.SeqOf(_457_v8)
		var _459_v10 _dafny.Map
		_ = _459_v10
		_459_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_v8, (_458_v9).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_458_v9).Cardinality()))).Uint32()).(_dafny.Map))
		var _460_v11 _dafny.Map
		_ = _460_v11
		_460_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _457_v8)
		var _461_v12 _dafny.Array
		_ = _461_v12
		var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
		_ = _nw77
		_461_v12 = _nw77
		var _462_v13 _dafny.Set
		_ = _462_v13
		_462_v13 = _dafny.SetOf(Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(724))).Uint32(), func(coer19 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_463_v3 T0, _464_v7 _dafny.Int, _465_v1 _dafny.Sequence) func(_dafny.Int) D0 {
			return func(_466_i1 _dafny.Int) D0 {
				return Companion_D0_.Create_DC2_(!((_463_v3).F20()), _464_v7, _dafny.SeqOf((_465_v1).Select((Companion_Default___.SafeIndex(_464_v7, _dafny.IntOfUint32((_465_v1).Cardinality()))).Uint32()).(bool)), _464_v7, _464_v7)
			}
		})(_452_v3, _456_v7, _449_v1)))))
		var _467_v14 D3
		_ = _467_v14
		_467_v14 = Companion_D3_.Create_DC10_(_461_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_456_v7, _dafny.IntOfInt64(844), _456_v7, globalState), _456_v7), _462_v13, _456_v7)
		_457_v8 = ((func() _dafny.Map {
			if (_452_v3).F20() {
				return (func() _dafny.Map {
					if (_459_v10).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v3).F20(), _456_v7)) {
						return (_459_v10).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v3).F20(), _456_v7)).(_dafny.Map)
					}
					return Companion_Default___.Fm16((_452_v3).F20(), !(_448_v0), _456_v7, globalState)
				})()
			}
			return (func() _dafny.Map {
				if (_460_v11).Contains((_452_v3).F20()) {
					return (_460_v11).Get((_452_v3).F20()).(_dafny.Map)
				}
				return _457_v8
			})()
		})()).Merge((_467_v14).Dtor_cf18())
		if (_452_v3).F20() {
			(globalState).F18 = _456_v7
			(globalState).F0 = _448_v0
			var _rhs32 _dafny.Int = (_dafny.IntOfInt64(-257)).Minus(_456_v7)
			_ = _rhs32
			var _rhs33 _dafny.Int = _456_v7
			_ = _rhs33
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			_456_v7 = _rhs32
			_lhs26.F18 = _rhs33
			var _468_v15 D1
			_ = _468_v15
			_468_v15 = Companion_D1_.Create_DC5_((_452_v3).F20())
			var _rhs34 bool = _448_v0
			_ = _rhs34
			var _rhs35 D1 = _468_v15
			_ = _rhs35
			var _rhs36 bool = !(false) || ((_452_v3).F20())
			_ = _rhs36
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			_lhs27.F0 = _rhs34
			_468_v15 = _rhs35
			_lhs28.F0 = _rhs36
			var _469_v16 _dafny.Sequence
			_ = _469_v16
			_469_v16 = _dafny.SeqOf((_452_v3).F19(), (_452_v3).F19(), (_452_v3).F19())
			var _470_v17 _dafny.CodePoint
			_ = _470_v17
			_470_v17 = _dafny.CodePoint('o')
			var _471_v18 _dafny.Map
			_ = _471_v18
			_471_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_470_v17, _456_v7)
			var _472_v19 _dafny.Sequence
			_ = _472_v19
			_472_v19 = _dafny.SeqOf(_471_v18)
			var _473_v20 D0
			_ = _473_v20
			_473_v20 = Companion_D0_.Create_DC1_((_452_v3).F20(), _dafny.IntOfInt64(514), _456_v7)
			var _474_v21 D5
			_ = _474_v21
			_474_v21 = Companion_D5_.Create_DC17_(_dafny.IntOfUint32((_449_v1).Cardinality()), _456_v7, _473_v20, (_452_v3).F20(), (_452_v3).F20())
			var _475_v22 _dafny.Array
			_ = _475_v22
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_11
			var _nw78 _dafny.Array
			_ = _nw78
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw78 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Int = func(_476_i2 _dafny.Int) _dafny.Int {
					return (_476_i2).Times(_dafny.IntOfUint32((_this.F22).Cardinality()))
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw78 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw78).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw78).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_475_v22 = _nw78
			var _477_v23 _dafny.MultiSet
			_ = _477_v23
			_477_v23 = _dafny.MultiSetOf(_475_v22, _475_v22)
			var _478_v24 _dafny.MultiSet
			_ = _478_v24
			_478_v24 = _dafny.MultiSetOf((_477_v23).Cardinality(), _456_v7)
			var _479_v25 _dafny.Sequence
			_ = _479_v25
			_479_v25 = _dafny.SeqOf(_456_v7)
			var _480_v26 _dafny.Sequence
			_ = _480_v26
			_480_v26 = _dafny.SeqOf(_dafny.SeqOf((func() _dafny.Int {
				if (_478_v24).Contains((_457_v8).Cardinality()) {
					return (_478_v24).Multiplicity((_457_v8).Cardinality())
				}
				return _456_v7
			})(), _456_v7, _dafny.IntOfUint32((_479_v25).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(453))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_481_v8 _dafny.Map) func(_dafny.Int) _dafny.Int {
				return func(_482_i3 _dafny.Int) _dafny.Int {
					return (_481_v8).Cardinality()
				}
			})(_457_v8))))
			var _pat_let_tv15 = _456_v7
			_ = _pat_let_tv15
			var _pat_let_tv16 = _449_v1
			_ = _pat_let_tv16
			var _483_v27 _dafny.Sequence
			_ = _483_v27
			_483_v27 = _dafny.SeqOf(func(_pat_let15_0 D0) D0 {
				return func(_484_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let16_0 _dafny.Int) D0 {
						return func(_485_dt__update_hcf8_h0 _dafny.Int) D0 {
							return func(_pat_let17_0 _dafny.Sequence) D0 {
								return func(_486_dt__update_hcf6_h0 _dafny.Sequence) D0 {
									return Companion_D0_.Create_DC2_((_484_dt__update__tmp_h0).Dtor_cf4(), (_484_dt__update__tmp_h0).Dtor_cf5(), _486_dt__update_hcf6_h0, (_484_dt__update__tmp_h0).Dtor_cf7(), _485_dt__update_hcf8_h0)
								}(_pat_let17_0)
							}(_pat_let_tv16)
						}(_pat_let16_0)
					}(_pat_let_tv15)
				}(_pat_let15_0)
			}(Companion_D0_.Create_DC2_((_452_v3).F20(), _456_v7, _449_v1, _456_v7, _dafny.IntOfUint32((_this.F22).Cardinality()))))
			var _487_v28 D1
			_ = _487_v28
			_487_v28 = Companion_D1_.Create_DC4_(_483_v27)
			var _488_v29 _dafny.Sequence
			_ = _488_v29
			_488_v29 = _dafny.SeqOf(_this.F22)
			var _pat_let_tv17 = _452_v3
			_ = _pat_let_tv17
			var _pat_let_tv18 = _456_v7
			_ = _pat_let_tv18
			var _pat_let_tv19 = _449_v1
			_ = _pat_let_tv19
			var _pat_let_tv20 = _456_v7
			_ = _pat_let_tv20
			var _pat_let_tv21 = _488_v29
			_ = _pat_let_tv21
			var _pat_let_tv22 = _456_v7
			_ = _pat_let_tv22
			var _pat_let_tv23 = _488_v29
			_ = _pat_let_tv23
			var _489_v30 _dafny.MultiSet
			_ = _489_v30
			_489_v30 = _dafny.MultiSetOf(func(_pat_let18_0 D1) D1 {
				return func(_490_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let19_0 _dafny.Sequence) D1 {
						return func(_491_dt__update_hcf10_h0 _dafny.Sequence) D1 {
							return Companion_D1_.Create_DC4_(_491_dt__update_hcf10_h0)
						}(_pat_let19_0)
					}(_dafny.SeqOf(Companion_D0_.Create_DC2_(!((_pat_let_tv17).F20()), _pat_let_tv18, _pat_let_tv19, _pat_let_tv20, _dafny.IntOfUint32(((_pat_let_tv21).Select((Companion_Default___.SafeIndex(_pat_let_tv22, _dafny.IntOfUint32((_pat_let_tv23).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))))
				}(_pat_let18_0)
			}(_487_v28), _487_v28)
			var _492_v31 _dafny.Array
			_ = _492_v31
			var _nwElement0_12 _dafny.Int = _dafny.IntOfUint32((_469_v16).Cardinality())
			_ = _nwElement0_12
			var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(27))
			_ = _nw79
			(_nw79).ArraySet1(_nwElement0_12, 0)
			(_nw79).ArraySet1(_dafny.IntOfUint32((_this.F22).Cardinality()), 1)
			(_nw79).ArraySet1((func() _dafny.Int {
				if _448_v0 {
					return _456_v7
				}
				return _dafny.IntOfInt64(851)
			})(), 2)
			(_nw79).ArraySet1((Companion_Default___.Fm0(false, (_452_v3).F20(), globalState)).Times(_456_v7), 3)
			(_nw79).ArraySet1((_dafny.Zero).Minus(_456_v7), 4)
			(_nw79).ArraySet1((_456_v7).Plus(_456_v7), 5)
			(_nw79).ArraySet1((_dafny.Zero).Minus(_456_v7), 6)
			(_nw79).ArraySet1(_dafny.IntOfInt64(398), 7)
			(_nw79).ArraySet1(((_dafny.SetOf((_452_v3).F20())).Difference(Companion_Default___.Fm17(_dafny.IntOfUint32((_this.F22).Cardinality()), globalState))).Cardinality(), 8)
			(_nw79).ArraySet1(_456_v7, 9)
			(_nw79).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qnxjphmtm")).Cardinality()), 10)
			(_nw79).ArraySet1(_dafny.IntOfUint32((_472_v19).Cardinality()), 11)
			(_nw79).ArraySet1(_456_v7, 12)
			(_nw79).ArraySet1(Companion_Default___.Fm0((_452_v3).F20(), (_474_v21).Dtor_cf33(), globalState), 13)
			(_nw79).ArraySet1(Companion_Default___.SafeModuloInt(_456_v7, _456_v7), 14)
			(_nw79).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus(_456_v7)).Times(_456_v7)), 15)
			(_nw79).ArraySet1(_456_v7, 16)
			(_nw79).ArraySet1(_456_v7, 17)
			(_nw79).ArraySet1(_dafny.IntOfUint32(((_480_v26).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_480_v26).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 18)
			(_nw79).ArraySet1(_456_v7, 19)
			(_nw79).ArraySet1(_456_v7, 20)
			(_nw79).ArraySet1(_456_v7, 21)
			(_nw79).ArraySet1(_456_v7, 22)
			(_nw79).ArraySet1(_456_v7, 23)
			(_nw79).ArraySet1((_456_v7).Times((_489_v30).Cardinality()), 24)
			(_nw79).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F22, _dafny.UnicodeSeqOfUtf8Bytes("w"))).Cardinality()), 25)
			(_nw79).ArraySet1(_456_v7, 26)
			_492_v31 = _nw79
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_475_v22), 0))
			_ = _index59
			(_475_v22).ArraySet1(_456_v7, (_index59).Int())
			var _493_v32 _dafny.Array
			_ = _493_v32
			var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
			_ = _nw80
			_493_v32 = _nw80
			var _494_v33 D5
			_ = _494_v33
			_494_v33 = Companion_D5_.Create_DC19_((_478_v24).Cardinality(), (_452_v3).F20(), _493_v32, (_452_v3).F20(), _470_v17)
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_475_v22), 0))
			_ = _index60
			var _rhs37 _dafny.Int = (_456_v7).Minus(_456_v7)
			_ = _rhs37
			var _rhs38 bool = (_449_v1).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_449_v1).Cardinality()))).Uint32()).(bool)
			_ = _rhs38
			var _rhs39 bool = (_494_v33).Dtor_cf36()
			_ = _rhs39
			var _rhs40 _dafny.Int = _456_v7
			_ = _rhs40
			var _lhs29 _dafny.Array = _475_v22
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(209), _dafny.ArrayLen((_475_v22), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			var _lhs32 *GlobalState = globalState
			_ = _lhs32
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
			_lhs31.F0 = _rhs38
			_lhs32.F0 = _rhs39
			_lhs33.F13 = _rhs40
		} else {
			(globalState).F0 = ((_dafny.MultiSetOf((_452_v3).F20())).Update((_452_v3).F20(), Companion_Default___.Abs(_456_v7))).IsDisjointFrom(_dafny.MultiSetOf(_448_v0, !((_452_v3).F20()), _448_v0))
			if _448_v0 {
				var _495_v34 *C0
				_ = _495_v34
				var _nw81 *C0 = New_C0_()
				_ = _nw81
				_nw81.Ctor__(_456_v7)
				_495_v34 = _nw81
				var _496_v35 _dafny.CodePoint
				_ = _496_v35
				_496_v35 = _dafny.CodePoint('p')
				var _497_v36 _dafny.Map
				_ = _497_v36
				_497_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v35, (_452_v3).F19())
				var _498_v37 *C2
				_ = _498_v37
				var _nw82 *C2 = New_C2_()
				_ = _nw82
				_nw82.Ctor__(_495_v34, (_452_v3).F20(), (func() _dafny.Array {
					if !((_452_v3).F20()) {
						return (_452_v3).F19()
					}
					return (func() _dafny.Array {
						if (_497_v36).Contains(_496_v35) {
							return (_497_v36).Get(_496_v35).(_dafny.Array)
						}
						return (_452_v3).F19()
					})()
				})(), Companion_Default___.Fm9((_452_v3).F20(), globalState))
				_498_v37 = _nw82
				var _499_v38 _dafny.Array
				_ = _499_v38
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw83
				_499_v38 = _nw83
				var _500_v39 _dafny.MultiSet
				_ = _500_v39
				_500_v39 = _dafny.MultiSetOf(false, (_452_v3).F20())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_499_v38), 0))
				_ = _index61
				(_499_v38).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_500_v39)).Cardinality()), (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_499_v38), 0))
				_ = _index62
				(_499_v38).ArraySet1(_456_v7, (_index62).Int())
				_456_v7 = Companion_Default___.SafeModuloInt((_456_v7).Plus((_495_v34).F24()), (_499_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_499_v38), 0))).Int()).(_dafny.Int))
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen(((_452_v3).F19()), 0))
				_ = _index63
				((_452_v3).F19()).ArraySet1(!(((_452_v3).F20()) || ((_452_v3).F20())), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen(((_452_v3).F19()), 0))
				_ = _index64
				((_452_v3).F19()).ArraySet1(false, (_index64).Int())
				var _501_v40 _dafny.Sequence
				_ = _501_v40
				_501_v40 = _dafny.SeqOf(((_499_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_499_v38), 0))).Int()).(_dafny.Int)).Times((_499_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_499_v38), 0))).Int()).(_dafny.Int)), _456_v7)
				(globalState).F13 = (_501_v40).Select((Companion_Default___.SafeIndex((_495_v34).F24(), _dafny.IntOfUint32((_501_v40).Cardinality()))).Uint32()).(_dafny.Int)
			} else {
				var _502_v41 D3
				_ = _502_v41
				_502_v41 = Companion_D3_.Create_DC9_(_this.F22)
				var _503_v42 _dafny.Map
				_ = _503_v42
				_503_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_448_v0, _dafny.MultiSetOf(_502_v41, _502_v41))
				var _504_v43 _dafny.MultiSet
				_ = _504_v43
				_504_v43 = _dafny.MultiSetOf(_502_v41)
				var _505_v44 _dafny.Sequence
				_ = _505_v44
				_505_v44 = _dafny.SeqOf(_456_v7, _dafny.IntOfInt64(448), _456_v7, _456_v7)
				(globalState).F0 = (Companion_Default___.Fm18(_505_v44, globalState)).IsSubsetOf(((func() _dafny.MultiSet {
					if (_503_v42).Contains((_452_v3).F20()) {
						return (_503_v42).Get((_452_v3).F20()).(_dafny.MultiSet)
					}
					return _504_v43
				})()).Union(_504_v43))
				var _506_v45 _dafny.Map
				_ = _506_v45
				_506_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_448_v0, (_452_v3).F20())
				_506_v45 = (_506_v45).Update((_452_v3).F20(), (_452_v3).F20())
				var _507_v46 _dafny.Map
				_ = _507_v46
				_507_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, !((_452_v3).F20()))
				var _508_v47 *C0
				_ = _508_v47
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__((_507_v46).Cardinality())
				_508_v47 = _nw84
				var _509_v48 *C2
				_ = _509_v48
				var _nw85 *C2 = New_C2_()
				_ = _nw85
				_nw85.Ctor__(_508_v47, (_452_v3).F20(), (_452_v3).F19(), _dafny.Companion_Sequence_.IsProperPrefixOf(_this.F22, _this.F22))
				_509_v48 = _nw85
				var _510_v49 _dafny.Array
				_ = _510_v49
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.One)
				_ = _nw86
				_510_v49 = _nw86
				var _511_v50 _dafny.Array
				_ = _511_v50
				_511_v50 = _510_v49
				var _512_v51 _dafny.Array
				_ = _512_v51
				var _nwElement0_13 _dafny.Array = _510_v49
				_ = _nwElement0_13
				var _nw87 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(6))
				_ = _nw87
				(_nw87).ArraySet1(_nwElement0_13, 0)
				(_nw87).ArraySet1(_510_v49, 1)
				(_nw87).ArraySet1(_510_v49, 2)
				(_nw87).ArraySet1(_510_v49, 3)
				(_nw87).ArraySet1(_510_v49, 4)
				(_nw87).ArraySet1((_511_v50), 5)
				_512_v51 = _nw87
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_512_v51), 0))
				_ = _index65
				(_512_v51).ArraySet1(_510_v49, (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_512_v51), 0))
				_ = _index66
				(_512_v51).ArraySet1(_510_v49, (_index66).Int())
				var _513_v52 _dafny.MultiSet
				_ = _513_v52
				_513_v52 = _dafny.MultiSetOf(!(_448_v0), (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20())
				(globalState).F0 = ((func() _dafny.Int {
					if (_513_v52).Contains(!(_448_v0)) {
						return (_513_v52).Multiplicity(!(_448_v0))
					}
					return (Companion_Default___.Fm19(true, globalState)).Cardinality()
				})()).Cmp((_508_v47).F24()) < 0
			}
			var _514_v53 _dafny.Sequence
			_ = _514_v53
			var _515_v54 bool
			_ = _515_v54
			var _out20 _dafny.Sequence
			_ = _out20
			var _out21 bool
			_ = _out21
			_out20, _out21 = (_this).M2(globalState)
			_514_v53 = _out20
			_515_v54 = _out21
			var _516_v55 _dafny.CodePoint
			_ = _516_v55
			_516_v55 = _dafny.CodePoint('o')
			_516_v55 = _516_v55
			(globalState).F13 = ((_dafny.Zero).Minus(_456_v7)).Minus((_456_v7).Plus(_456_v7))
		}
		if !(_448_v0) {
			_456_v7 = _dafny.IntOfInt64(-455)
			var _517_v56 _dafny.Set
			_ = _517_v56
			_517_v56 = _dafny.SetOf(_456_v7)
			var _rhs41 _dafny.Int = _456_v7
			_ = _rhs41
			var _rhs42 bool = (_452_v3).F20()
			_ = _rhs42
			var _rhs43 bool = (_452_v3).F20()
			_ = _rhs43
			var _rhs44 bool = ((func() _dafny.Set {
				if (_452_v3).F20() {
					return _517_v56
				}
				return _517_v56
			})()).IsSubsetOf(_517_v56)
			_ = _rhs44
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			var _lhs35 *GlobalState = globalState
			_ = _lhs35
			var _lhs36 *GlobalState = globalState
			_ = _lhs36
			var _lhs37 *GlobalState = globalState
			_ = _lhs37
			_lhs34.F13 = _rhs41
			_lhs35.F0 = _rhs42
			_lhs36.F0 = _rhs43
			_lhs37.F0 = _rhs44
			(globalState).F13 = (func() _dafny.Int {
				if (_452_v3).F20() {
					return (_457_v8).Cardinality()
				}
				return _456_v7
			})()
			var _518_v57 _dafny.CodePoint
			_ = _518_v57
			_518_v57 = _dafny.CodePoint('p')
			(globalState).F13 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_this.F22, _dafny.Companion_Sequence_.Update(_this.F22, (Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_this.F22).Cardinality()))).Uint32(), _518_v57)), _dafny.UnicodeSeqOfUtf8Bytes("usturyl"))).Cardinality())
			if (_452_v3).F20() {
				var _519_v58 D4
				_ = _519_v58
				_519_v58 = Companion_D4_.Create_DC13_()
				_519_v58 = _519_v58
				var _520_v59 _dafny.Map
				_ = _520_v59
				_520_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v3).F20(), _518_v57)
				var _521_v60 _dafny.Map
				_ = _521_v60
				_521_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v3).F20(), (_520_v59).Contains(_448_v0))
				_521_v60 = ((_521_v60).Merge(_521_v60)).Merge(_521_v60)
				var _522_v61 _dafny.Sequence
				_ = _522_v61
				_522_v61 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("i"))
				var _523_v62 _dafny.Array
				_ = _523_v62
				var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw88
				_523_v62 = _nw88
				var _524_v63 _dafny.Map
				_ = _524_v63
				_524_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_522_v61).Cardinality()), _523_v62)
				var _525_v64 _dafny.Map
				_ = _525_v64
				_525_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, (func() _dafny.Array {
					if (_524_v63).Contains(_dafny.IntOfInt64(680)) {
						return (_524_v63).Get(_dafny.IntOfInt64(680)).(_dafny.Array)
					}
					return _523_v62
				})())
				_525_v64 = (_525_v64).Update((func() _dafny.Int {
					if (_452_v3).F20() {
						return _456_v7
					}
					return _456_v7
				})(), _523_v62)
				(globalState).F0 = (func() bool {
					if Companion_Default___.Fm7((Companion_Default___.Fm19(true, globalState)).Cardinality(), _456_v7, _456_v7, globalState) {
						return !((_452_v3).F20())
					}
					return _dafny.Companion_Sequence_.Contains(_449_v1, _448_v0)
				})()
				var _526_v65 D0
				_ = _526_v65
				_526_v65 = Companion_D0_.Create_DC1_((_452_v3).F20(), _456_v7, _456_v7)
				var _527_v66 _dafny.Array
				_ = _527_v66
				var _nwElement0_14 _dafny.Int = _456_v7
				_ = _nwElement0_14
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(9))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_14, 0)
				(_nw89).ArraySet1(Companion_Default___.Fm0((_526_v65).Dtor_cf1(), (_452_v3).F20(), globalState), 1)
				(_nw89).ArraySet1(_456_v7, 2)
				(_nw89).ArraySet1(_456_v7, 3)
				(_nw89).ArraySet1(_dafny.IntOfUint32((_449_v1).Cardinality()), 4)
				(_nw89).ArraySet1(_dafny.IntOfInt64(139), 5)
				(_nw89).ArraySet1(_456_v7, 6)
				(_nw89).ArraySet1(_456_v7, 7)
				(_nw89).ArraySet1(_456_v7, 8)
				_527_v66 = _nw89
				var _528_v67 D5
				_ = _528_v67
				_528_v67 = Companion_D5_.Create_DC16_(_527_v66)
				var _529_v68 _dafny.Map
				_ = _529_v68
				_529_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_528_v67, (_452_v3).F19())
				var _530_v69 *C1
				_ = _530_v69
				var _nw90 *C1 = New_C1_()
				_ = _nw90
				_nw90.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("sjlkirkt"), (func() _dafny.Array {
					if (_529_v68).Contains(_528_v67) {
						return (_529_v68).Get(_528_v67).(_dafny.Array)
					}
					return (_452_v3).F19()
				})(), (_449_v1).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_449_v1).Cardinality()))).Uint32()).(bool))
				_530_v69 = _nw90
			} else {
				(globalState).F18 = (func() _dafny.Int {
					if (_457_v8).Contains((_452_v3).F20()) {
						return (_457_v8).Get((_452_v3).F20()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.SeqOf((_452_v3).F20())).Cardinality())
				})()
				_456_v7 = _dafny.IntOfUint32((_this.F22).Cardinality())
				var _531_v70 D0
				_ = _531_v70
				_531_v70 = Companion_D0_.Create_DC1_((_452_v3).F20(), _dafny.IntOfInt64(310), (_dafny.Zero).Minus(_456_v7))
				var _532_v71 _dafny.MultiSet
				_ = _532_v71
				_532_v71 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_((_452_v3).F20(), _456_v7, _456_v7), _531_v70, _531_v70, _531_v70)
				var _533_v72 *C1
				_ = _533_v72
				var _nw91 *C1 = New_C1_()
				_ = _nw91
				_nw91.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}((func(_534_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_535_i4 _dafny.Int) _dafny.CodePoint {
						return _534_v57
					}
				})(_518_v57))), (_this).F23(), (_452_v3).F20())
				_533_v72 = _nw91
				var _536_v73 _dafny.Map
				_ = _536_v73
				_536_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, _533_v72)
				var _537_v74 _dafny.Map
				_ = _537_v74
				_537_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, (_536_v73).Cardinality())
				var _538_v75 _dafny.Sequence
				_ = _538_v75
				_538_v75 = _dafny.SeqOf((func() _dafny.Int {
					if (_532_v71).Contains(Companion_D0_.Create_DC1_((_452_v3).F20(), (func() _dafny.Int {
						if (_537_v74).Contains(_456_v7) {
							return (_537_v74).Get(_456_v7).(_dafny.Int)
						}
						return _456_v7
					})(), _456_v7)) {
						return (_532_v71).Multiplicity(Companion_D0_.Create_DC1_((_452_v3).F20(), (func() _dafny.Int {
							if (_537_v74).Contains(_456_v7) {
								return (_537_v74).Get(_456_v7).(_dafny.Int)
							}
							return _456_v7
						})(), _456_v7))
					}
					return _456_v7
				})(), _456_v7, (_dafny.MultiSetOf(_456_v7)).Cardinality())
				var _539_v76 *C0
				_ = _539_v76
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__(_456_v7)
				_539_v76 = _nw92
				var _540_v77 _dafny.Sequence
				_ = _540_v77
				_540_v77 = _dafny.SeqOf(_539_v76)
				var _541_v78 D5
				_ = _541_v78
				_541_v78 = Companion_D5_.Create_DC18_(_540_v77)
				var _542_v79 _dafny.Map
				_ = _542_v79
				_542_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v78, (_452_v3).F20())
				var _543_v80 _dafny.Array
				_ = _543_v80
				var _nwElement0_15 _dafny.Int = _456_v7
				_ = _nwElement0_15
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(16))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_15, 0)
				(_nw93).ArraySet1(_456_v7, 1)
				(_nw93).ArraySet1(_456_v7, 2)
				(_nw93).ArraySet1(_456_v7, 3)
				(_nw93).ArraySet1(_456_v7, 4)
				(_nw93).ArraySet1((_538_v75).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_538_v75).Cardinality()))).Uint32()).(_dafny.Int), 5)
				(_nw93).ArraySet1(_456_v7, 6)
				(_nw93).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_456_v7, _456_v7)), 7)
				(_nw93).ArraySet1((_456_v7).Minus(_456_v7), 8)
				(_nw93).ArraySet1((_542_v79).Cardinality(), 9)
				(_nw93).ArraySet1((_this).Fm5((_452_v3).F20(), globalState), 10)
				(_nw93).ArraySet1((_dafny.IntOfUint32((_449_v1).Cardinality())).Plus((_539_v76).F24()), 11)
				(_nw93).ArraySet1(_456_v7, 12)
				(_nw93).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_452_v3).F20() {
						return _449_v1
					}
					return _449_v1
				})()).Cardinality()), 13)
				(_nw93).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_this.F22, _this.F22), (Companion_Default___.SafeIndex((_539_v76).F24(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_this.F22, _this.F22)).Cardinality()))).Uint32(), _518_v57)).Cardinality()), 14)
				(_nw93).ArraySet1((_456_v7).Times((_dafny.Zero).Minus((_539_v76).F24())), 15)
				_543_v80 = _nw93
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_543_v80), 0))
				_ = _index67
				(_543_v80).ArraySet1((_539_v76).F24(), (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_543_v80), 0))
				_ = _index68
				(_543_v80).ArraySet1(Companion_Default___.SafeDivisionInt((_539_v76).F24(), _dafny.IntOfUint32((_dafny.SeqOf((_452_v3).F20(), _448_v0)).Cardinality())), (_index68).Int())
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_543_v80), 0))
				_ = _index69
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_543_v80), 0))
				_ = _index70
				var _rhs45 _dafny.Int = (_456_v7).Minus((_539_v76).F24())
				_ = _rhs45
				var _rhs46 _dafny.Int = _456_v7
				_ = _rhs46
				var _rhs47 _dafny.Int = _456_v7
				_ = _rhs47
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				var _lhs39 _dafny.Array = _543_v80
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_543_v80), 0))
				_ = _lhs40
				var _lhs41 _dafny.Array = _543_v80
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_543_v80), 0))
				_ = _lhs42
				_lhs38.F13 = _rhs45
				(_lhs39).ArraySet1(_rhs46, (_lhs40).Int())
				(_lhs41).ArraySet1(_rhs47, (_lhs42).Int())
				var _544_v81 D4
				_ = _544_v81
				_544_v81 = Companion_D4_.Create_DC14_((_452_v3).F20(), _dafny.SetOf((_452_v3).F20()))
				(globalState).F0 = (func() bool {
					if (_452_v3).F20() {
						return (_544_v81).Dtor_cf25()
					}
					return (_452_v3).F20()
				})()
				(globalState).F13 = _456_v7
			}
		} else {
			var _545_v82 _dafny.Set
			_ = _545_v82
			_545_v82 = _dafny.SetOf((_452_v3).F20())
			var _546_v83 _dafny.MultiSet
			_ = _546_v83
			_546_v83 = _dafny.MultiSetOf(_456_v7, _456_v7)
			var _547_v84 _dafny.Sequence
			_ = _547_v84
			_547_v84 = _dafny.SeqOf((_546_v83).Cardinality())
			var _548_v85 _dafny.Map
			_ = _548_v85
			_548_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_v82, (_547_v84).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_547_v84).Cardinality()))).Uint32()).(_dafny.Int))
			var _549_v86 _dafny.Map
			_ = _549_v86
			_549_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, _456_v7)
			var _550_v87 _dafny.Map
			_ = _550_v87
			_550_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_549_v86, true)
			if ((func() _dafny.Int {
				if (_548_v85).Contains(_545_v82) {
					return (_548_v85).Get(_545_v82).(_dafny.Int)
				}
				return _456_v7
			})()).Cmp((_dafny.Zero).Minus((_550_v87).Cardinality())) < 0 {
				var _551_v88 _dafny.Set
				_ = _551_v88
				_551_v88 = _dafny.SetOf(Companion_Default___.SafeModuloInt(_456_v7, _456_v7), _456_v7, Companion_Default___.Fm0((_452_v3).F20(), (_452_v3).F20(), globalState), _456_v7)
				_551_v88 = (func() _dafny.Set {
					var _coll11 = _dafny.NewBuilder()
					_ = _coll11
					for _iter14 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_this.F22).Cardinality()))).Elements()); ; {
						_compr_11, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _552_v89 _dafny.Int
						_552_v89 = interface{}(_compr_11).(_dafny.Int)
						if (_dafny.MultiSetOf(_dafny.IntOfUint32((_this.F22).Cardinality()))).Contains(_552_v89) {
							_coll11.Add(Companion_Default___.SafeModuloInt(_552_v89, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(true), _dafny.SeqOf(false)))).Cardinality()))
						}
					}
					return _coll11.ToSet()
				}()).Intersection(_551_v88)
				var _553_v90 _dafny.CodePoint
				_ = _553_v90
				_553_v90 = _dafny.CodePoint('p')
				var _554_v91 _dafny.Map
				_ = _554_v91
				_554_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v3).F20(), _553_v90)
				_554_v91 = (_554_v91).Update(_448_v0, _553_v90)
				var _555_v92 _dafny.Map
				_ = _555_v92
				_555_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, (_this).F23())
				var _rhs48 _dafny.Map = _555_v92
				_ = _rhs48
				var _rhs49 bool = !(!((_452_v3).F20()))
				_ = _rhs49
				_555_v92 = _rhs48
				_448_v0 = _rhs49
				var _556_v93 _dafny.Set
				_ = _556_v93
				_556_v93 = _dafny.SetOf(_545_v82)
				var _557_v94 _dafny.MultiSet
				_ = _557_v94
				_557_v94 = _dafny.MultiSetOf(_553_v90)
				var _558_v95 _dafny.Map
				_ = _558_v95
				_558_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, _557_v94)
				_449_v1 = _dafny.SeqOf((_556_v93).IsProperSubsetOf(_556_v93), (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20(), !((func() _dafny.MultiSet {
					if (_558_v95).Contains(_dafny.IntOfUint32((_547_v84).Cardinality())) {
						return (_558_v95).Get(_dafny.IntOfUint32((_547_v84).Cardinality())).(_dafny.MultiSet)
					}
					return _557_v94
				})()).Equals(_557_v94))
				var _559_v96 _dafny.Array
				_ = _559_v96
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_12
				var _nw94 _dafny.Array
				_ = _nw94
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw94 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Sequence = (func(_560_v1 _dafny.Sequence, _561_v3 T0) func(_dafny.Int) _dafny.Sequence {
						return func(_562_i5 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_560_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_560_v1).Cardinality()), _dafny.IntOfUint32((_560_v1).Cardinality()))).Uint32(), (_561_v3).F20()))), Companion_D0_.Create_DC0_(_dafny.MultiSetOf((_561_v3).F20(), (_561_v3).F20(), (_561_v3).F20(), !((_561_v3).F20()), (_561_v3).F20())), Companion_D0_.Create_DC0_(_dafny.MultiSetFromSeq(_560_v1)))
						}
					})(_449_v1, _452_v3)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw94 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw94).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw94).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_559_v96 = _nw94
				var _563_v97 _dafny.MultiSet
				_ = _563_v97
				_563_v97 = _dafny.MultiSetOf(_448_v0)
				var _564_v98 D0
				_ = _564_v98
				_564_v98 = Companion_D0_.Create_DC0_(_563_v97)
				var _565_v99 _dafny.Sequence
				_ = _565_v99
				_565_v99 = _dafny.SeqOf(Companion_D0_.Create_DC0_(_563_v97), _564_v98)
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_559_v96), 0))
				_ = _index71
				(_559_v96).ArraySet1((func() _dafny.Sequence {
					if (_452_v3).F20() {
						return _565_v99
					}
					return _565_v99
				})(), (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_559_v96), 0))
				_ = _index72
				(_559_v96).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_565_v99, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_456_v7), _dafny.IntOfUint32((_565_v99).Cardinality()))).Uint32(), Companion_D0_.Create_DC0_(Companion_Default___.Fm20((_452_v3).F20(), _456_v7, globalState))), (Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_565_v99, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_456_v7), _dafny.IntOfUint32((_565_v99).Cardinality()))).Uint32(), Companion_D0_.Create_DC0_(Companion_Default___.Fm20((_452_v3).F20(), _456_v7, globalState)))).Cardinality()))).Uint32(), (func() D0 {
					if (_452_v3).F20() {
						return _564_v98
					}
					return Companion_Default___.Fm21(_456_v7, globalState)
				})()), (_index72).Int())
			} else {
				var _566_v100 _dafny.MultiSet
				_ = _566_v100
				_566_v100 = _dafny.MultiSetOf(_this.F22, _this.F22)
				var _567_v102 *C1
				_ = _567_v102
				var _nw95 *C1 = New_C1_()
				_ = _nw95
				_nw95.Ctor__(_this.F22, (_452_v3).F19(), _448_v0)
				_567_v102 = _nw95
				var _568_v103 _dafny.MultiSet
				_ = _568_v103
				_568_v103 = _dafny.MultiSetOf(_567_v102)
				var _569_v104 _dafny.Map
				_ = _569_v104
				_569_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm22((func() _dafny.Int {
					if (_568_v103).Contains(_567_v102) {
						return (_568_v103).Multiplicity(_567_v102)
					}
					return _456_v7
				})(), _456_v7, globalState), _dafny.IntOfUint32(((_567_v102).F27()).Cardinality()))
				var _570_v105 _dafny.MultiSet
				_ = _570_v105
				_570_v105 = _dafny.MultiSetOf(_448_v0, (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20())
				var _571_v106 _dafny.CodePoint
				_ = _571_v106
				_571_v106 = _dafny.CodePoint('d')
				var _572_v107 _dafny.Map
				_ = _572_v107
				_572_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_571_v106, _448_v0)
				var _573_v108 _dafny.Set
				_ = _573_v108
				_573_v108 = _dafny.SetOf((_572_v107).Cardinality())
				var _574_v109 _dafny.Map
				_ = _574_v109
				_574_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_570_v105, (_573_v108).Cardinality())
				var _575_v110 _dafny.Array
				_ = _575_v110
				var _nwElement0_16 _dafny.Int = _dafny.IntOfUint32((_this.F22).Cardinality())
				_ = _nwElement0_16
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(28))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_16, 0)
				(_nw96).ArraySet1(_456_v7, 1)
				(_nw96).ArraySet1((func() _dafny.Int {
					if (_452_v3).F20() {
						return _456_v7
					}
					return (_dafny.Zero).Minus(_456_v7)
				})(), 2)
				(_nw96).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_456_v7), _dafny.IntOfInt64(645)), 3)
				(_nw96).ArraySet1((_dafny.Zero).Minus(_456_v7), 4)
				(_nw96).ArraySet1(_456_v7, 5)
				(_nw96).ArraySet1(_dafny.IntOfInt64(-369), 6)
				(_nw96).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("brlgftm"), false)).Cardinality(), 7)
				(_nw96).ArraySet1(_456_v7, 8)
				(_nw96).ArraySet1(_dafny.IntOfInt64(122), 9)
				(_nw96).ArraySet1(_456_v7, 10)
				(_nw96).ArraySet1((_456_v7).Times((_dafny.Zero).Minus(_456_v7)), 11)
				(_nw96).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_456_v7), _456_v7), 12)
				(_nw96).ArraySet1(_456_v7, 13)
				(_nw96).ArraySet1((_456_v7).Times((func() _dafny.Set {
					var _coll12 = _dafny.NewBuilder()
					_ = _coll12
					for _iter15 := _dafny.Iterate((_566_v100).Elements()); ; {
						_compr_12, _ok15 := _iter15()
						if !_ok15 {
							break
						}
						var _576_v101 _dafny.Sequence
						_576_v101 = interface{}(_compr_12).(_dafny.Sequence)
						if (_566_v100).Contains(_576_v101) {
							_coll12.Add(_576_v101)
						}
					}
					return _coll12.ToSet()
				}()).Cardinality()), 14)
				(_nw96).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_452_v3).F20() {
						return _456_v7
					}
					return _456_v7
				})())), 15)
				(_nw96).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_this.F22).Cardinality()), (_460_v11).Cardinality()), 16)
				(_nw96).ArraySet1(_456_v7, 17)
				(_nw96).ArraySet1((_566_v100).Cardinality(), 18)
				(_nw96).ArraySet1(_456_v7, 19)
				(_nw96).ArraySet1((_dafny.IntOfInt64(942)).Minus((_dafny.Zero).Minus(_456_v7)), 20)
				(_nw96).ArraySet1(_456_v7, 21)
				(_nw96).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(406), _456_v7), 22)
				(_nw96).ArraySet1(_456_v7, 23)
				(_nw96).ArraySet1((func() _dafny.Int {
					if (_569_v104).Contains(_574_v109) {
						return (_569_v104).Get(_574_v109).(_dafny.Int)
					}
					return _456_v7
				})(), 24)
				(_nw96).ArraySet1((_456_v7).Minus(_456_v7), 25)
				(_nw96).ArraySet1(_456_v7, 26)
				(_nw96).ArraySet1((_456_v7).Times(_456_v7), 27)
				_575_v110 = _nw96
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_575_v110), 0))
				_ = _index73
				(_575_v110).ArraySet1(_456_v7, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_575_v110), 0))
				_ = _index74
				(_575_v110).ArraySet1(_456_v7, (_index74).Int())
				var _577_v111 *C1
				_ = _577_v111
				var _nw97 *C1 = New_C1_()
				_ = _nw97
				_nw97.Ctor__((_567_v102).F27(), (_this).F23(), (_452_v3).F20())
				_577_v111 = _nw97
				var _578_v112 _dafny.Map
				_ = _578_v112
				_578_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_575_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_575_v110), 0))).Int()).(_dafny.Int)).Plus((_575_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_575_v110), 0))).Int()).(_dafny.Int)), false)
				_578_v112 = (_578_v112).Update((_456_v7).Times((_575_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_575_v110), 0))).Int()).(_dafny.Int)), (_452_v3).F20())
				_545_v82 = ((_545_v82).Intersection(_545_v82)).Intersection(_545_v82)
				_457_v8 = (_457_v8).Update((_452_v3).F20(), (_575_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_575_v110), 0))).Int()).(_dafny.Int))
			}
			var _579_v113 D4
			_ = _579_v113
			_579_v113 = Companion_D4_.Create_DC14_((_452_v3).F20(), _545_v82)
			var _580_v114 _dafny.Sequence
			_ = _580_v114
			_580_v114 = _dafny.SeqOf(_dafny.SetOf(_448_v0, (_579_v113).Dtor_cf25()))
			var _581_v115 _dafny.CodePoint
			_ = _581_v115
			_581_v115 = _dafny.CodePoint('y')
			var _582_v116 _dafny.Map
			_ = _582_v116
			_582_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_581_v115, (_452_v3).F20())
			var _583_v117 _dafny.Map
			_ = _583_v117
			_583_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v3).F20(), (_452_v3).F20())
			var _584_v118 _dafny.Map
			_ = _584_v118
			_584_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(917), (_452_v3).F20())
			var _585_v119 _dafny.Map
			_ = _585_v119
			_585_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_456_v7, _dafny.SetOf(!((func() bool {
				if (_584_v118).Contains(_dafny.IntOfInt64(-341)) {
					return (_584_v118).Get(_dafny.IntOfInt64(-341)).(bool)
				}
				return (_449_v1).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_449_v1).Cardinality()))).Uint32()).(bool)
			})())))
			var _586_v120 _dafny.Array
			_ = _586_v120
			var _nwElement0_17 _dafny.Set = (_580_v114).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_456_v7), _dafny.IntOfUint32((_580_v114).Cardinality()))).Uint32()).(_dafny.Set)
			_ = _nwElement0_17
			var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(19))
			_ = _nw98
			(_nw98).ArraySet1(_nwElement0_17, 0)
			(_nw98).ArraySet1(_dafny.SetOf((_452_v3).F20(), (func() bool {
				if (_582_v116).Contains(_dafny.CodePoint('w')) {
					return (_582_v116).Get(_dafny.CodePoint('w')).(bool)
				}
				return (_452_v3).F20()
			})()), 1)
			(_nw98).ArraySet1(_545_v82, 2)
			(_nw98).ArraySet1(_545_v82, 3)
			(_nw98).ArraySet1(_545_v82, 4)
			(_nw98).ArraySet1(_545_v82, 5)
			(_nw98).ArraySet1(_545_v82, 6)
			(_nw98).ArraySet1(Companion_Default___.Fm17((_583_v117).Cardinality(), globalState), 7)
			(_nw98).ArraySet1((_545_v82).Difference(_545_v82), 8)
			(_nw98).ArraySet1(_dafny.SetOf((_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20()), 9)
			(_nw98).ArraySet1((_545_v82).Union((func() _dafny.Set {
				if (_585_v119).Contains(_dafny.IntOfUint32((_this.F22).Cardinality())) {
					return (_585_v119).Get(_dafny.IntOfUint32((_this.F22).Cardinality())).(_dafny.Set)
				}
				return _545_v82
			})()), 10)
			(_nw98).ArraySet1(_545_v82, 11)
			(_nw98).ArraySet1(_545_v82, 12)
			(_nw98).ArraySet1((Companion_Default___.Fm17(_456_v7, globalState)).Difference(_dafny.SetOf((_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20(), (_452_v3).F20())), 13)
			(_nw98).ArraySet1(_545_v82, 14)
			(_nw98).ArraySet1((_580_v114).Select((Companion_Default___.SafeIndex(_456_v7, _dafny.IntOfUint32((_580_v114).Cardinality()))).Uint32()).(_dafny.Set), 15)
			(_nw98).ArraySet1(_545_v82, 16)
			(_nw98).ArraySet1((_545_v82).Difference(_545_v82), 17)
			(_nw98).ArraySet1((_545_v82).Difference(_545_v82), 18)
			_586_v120 = _nw98
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_586_v120), 0))
			_ = _index75
			(_586_v120).ArraySet1(_545_v82, (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_586_v120), 0))
			_ = _index76
			(_586_v120).ArraySet1(((_545_v82).Union(_545_v82)).Intersection(_dafny.SetOf((_452_v3).F20(), _448_v0)), (_index76).Int())
			var _587_v121 _dafny.Array
			_ = _587_v121
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_13
			var _nw99 _dafny.Array
			_ = _nw99
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw99 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Sequence = (func(_588_v7 _dafny.Int, _589_v118 _dafny.Map, _590_v3 T0, _591_v84 _dafny.Sequence, _592_v117 _dafny.Map) func(_dafny.Int) _dafny.Sequence {
					return func(_593_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(Companion_D3_.Create_DC11_(_588_v7, (_589_v118).Cardinality(), (_590_v3).F20()), Companion_D3_.Create_DC11_(_dafny.IntOfUint32((_591_v84).Cardinality()), (_592_v117).Cardinality(), (_590_v3).F20()), Companion_D3_.Create_DC11_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-641))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg22 _dafny.Int) interface{} {
								return coer22(arg22)
							}
						}((func(_594_v117 _dafny.Map) func(_dafny.Int) _dafny.Int {
							return func(_595_i7 _dafny.Int) _dafny.Int {
								return (_594_v117).Cardinality()
							}
						})(_592_v117)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(830))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg23 _dafny.Int) interface{} {
								return coer23(arg23)
							}
						}(func(_596_i8 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('a')
						}))).Cardinality()), (_590_v3).F20()), Companion_D3_.Create_DC11_(_588_v7, _588_v7, true), Companion_D3_.Create_DC11_(_588_v7, _588_v7, (_590_v3).F20()))
					}
				})(_456_v7, _584_v118, _452_v3, _547_v84, _583_v117)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw99 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw99).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw99).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_587_v121 = _nw99
			var _597_v122 D5
			_ = _597_v122
			_597_v122 = Companion_D5_.Create_DC19_(_456_v7, (_452_v3).F20(), _587_v121, _448_v0, _581_v115)
			_583_v117 = (_583_v117).Update((_597_v122).Dtor_cf38(), !(true))
			var _rhs50 _dafny.Int = _456_v7
			_ = _rhs50
			var _rhs51 bool = (_449_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F22).Cardinality()), _dafny.IntOfUint32((_449_v1).Cardinality()))).Uint32()).(bool)
			_ = _rhs51
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			_lhs43.F18 = _rhs50
			_lhs44.F0 = _rhs51
			_456_v7 = _456_v7
		}
	}
}
func (_this *C3) F23() _dafny.Array {
	{
		return _this._f23
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f19 _dafny.Array
	_f20 bool
	_f21 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f19 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f20 = false
	_this._f21 = _dafny.Zero
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

func (_this *C4) F19() _dafny.Array {
	return _this._f19
}
func (_this *C4) F20() bool {
	return _this._f20
}
func (_this *C4) Ctor__(f21 _dafny.Int, f19 _dafny.Array, f20 bool) {
	{
		(_this)._f21 = f21
		(_this)._f19 = f19
		(_this)._f20 = f20
	}
}
func (_this *C4) Fm1(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		if (_this).F20() {
			return _dafny.MultiSetOf((_this).F20())
		} else {
			return ((Companion_D0_.Create_DC0_(_dafny.MultiSetOf((_this).F20()))).Dtor_cf0()).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F20())))
		}
	}
}
func (_this *C4) M0(p0 bool, p1 _dafny.Set, globalState *GlobalState) {
	{
		var _598_v0 _dafny.Sequence
		_ = _598_v0
		_598_v0 = _dafny.UnicodeSeqOfUtf8Bytes("jrrji")
		(globalState).F13 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_598_v0, _dafny.UnicodeSeqOfUtf8Bytes("qqvxm"))).Cardinality())
		_598_v0 = _598_v0
		(globalState).F0 = !((_this).F20())
		var _599_v1 _dafny.MultiSet
		_ = _599_v1
		_599_v1 = _dafny.MultiSetOf((_this).F20(), p0)
		var _600_v2 D0
		_ = _600_v2
		_600_v2 = Companion_D0_.Create_DC0_(_599_v1)
		var _pat_let_tv24 = p0
		_ = _pat_let_tv24
		var _pat_let_tv25 = p0
		_ = _pat_let_tv25
		var _pat_let_tv26 = p0
		_ = _pat_let_tv26
		if func(_source8 D0) bool {
			if _source8.Is_DC1() {
				var _601___mcc_h0 bool = _source8.Get_().(D0_DC1).Cf1
				_ = _601___mcc_h0
				var _602___mcc_h1 _dafny.Int = _source8.Get_().(D0_DC1).Cf2
				_ = _602___mcc_h1
				var _603___mcc_h2 _dafny.Int = _source8.Get_().(D0_DC1).Cf3
				_ = _603___mcc_h2
				var _604_cf3 _dafny.Int = _603___mcc_h2
				_ = _604_cf3
				var _605_cf2 _dafny.Int = _602___mcc_h1
				_ = _605_cf2
				var _606_cf1 bool = _601___mcc_h0
				_ = _606_cf1
				return _pat_let_tv24
			} else if _source8.Is_DC2() {
				var _607___mcc_h3 bool = _source8.Get_().(D0_DC2).Cf4
				_ = _607___mcc_h3
				var _608___mcc_h4 _dafny.Int = _source8.Get_().(D0_DC2).Cf5
				_ = _608___mcc_h4
				var _609___mcc_h5 _dafny.Sequence = _source8.Get_().(D0_DC2).Cf6
				_ = _609___mcc_h5
				var _610___mcc_h6 _dafny.Int = _source8.Get_().(D0_DC2).Cf7
				_ = _610___mcc_h6
				var _611___mcc_h7 _dafny.Int = _source8.Get_().(D0_DC2).Cf8
				_ = _611___mcc_h7
				var _612_cf8 _dafny.Int = _611___mcc_h7
				_ = _612_cf8
				var _613_cf7 _dafny.Int = _610___mcc_h6
				_ = _613_cf7
				var _614_cf6 _dafny.Sequence = _609___mcc_h5
				_ = _614_cf6
				var _615_cf5 _dafny.Int = _608___mcc_h4
				_ = _615_cf5
				var _616_cf4 bool = _607___mcc_h3
				_ = _616_cf4
				return _616_cf4
			} else if _source8.Is_DC3() {
				var _617___mcc_h8 _dafny.Map = _source8.Get_().(D0_DC3).Cf9
				_ = _617___mcc_h8
				var _618_cf9 _dafny.Map = _617___mcc_h8
				_ = _618_cf9
				return _pat_let_tv25
			} else {
				var _619___mcc_h9 _dafny.MultiSet = _source8.Get_().(D0_DC0).Cf0
				_ = _619___mcc_h9
				var _620_cf0 _dafny.MultiSet = _619___mcc_h9
				_ = _620_cf0
				return _pat_let_tv26
			}
		}(_600_v2) {
			(globalState).F0 = p0
			(globalState).F0 = !_dafny.Companion_Sequence_.Equal(_598_v0, _dafny.UnicodeSeqOfUtf8Bytes("sqfpsses"))
			var _621_v3 _dafny.Map
			_ = _621_v3
			_621_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
			var _622_v4 _dafny.Map
			_ = _622_v4
			_622_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_598_v0, _598_v0), _621_v3)
			_622_v4 = _622_v4
			if p0 {
				(globalState).F13 = (_this).F21()
				var _623_v5 _dafny.Sequence
				_ = _623_v5
				_623_v5 = _dafny.SeqOf(p0)
				var _624_v6 D0
				_ = _624_v6
				_624_v6 = Companion_D0_.Create_DC2_((_this).F20(), Companion_Default___.Fm0((_this).F20(), true, globalState), _623_v5, (_this).F21(), (_this).F21())
				(globalState).F13 = (_624_v6).Dtor_cf5()
				(globalState).F0 = !((true) && (((_this).F21()).Cmp((_this).F21()) > 0))
				(globalState).F0 = p0
				var _625_v7 _dafny.Map
				_ = _625_v7
				_625_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_598_v0).Cardinality())), _dafny.MultiSetOf((_this).F20(), p0, p0, p0))
				var _626_v8 _dafny.Sequence
				_ = _626_v8
				_626_v8 = _dafny.SeqOf(_599_v1, _599_v1, _599_v1, _599_v1, _599_v1)
				var _627_v9 _dafny.Sequence
				_ = _627_v9
				_627_v9 = _dafny.SeqOf((_this).F21())
				var _628_v10 _dafny.Array
				_ = _628_v10
				var _nwElement0_18 _dafny.MultiSet = _dafny.MultiSetOf(p0, (_this).F20(), p0)
				_ = _nwElement0_18
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(26))
				_ = _nw100
				(_nw100).ArraySet1(_nwElement0_18, 0)
				(_nw100).ArraySet1(((func() _dafny.MultiSet {
					if (_625_v7).Contains((_this).F21()) {
						return (_625_v7).Get((_this).F21()).(_dafny.MultiSet)
					}
					return _599_v1
				})()).Update(!((_this).F20()), Companion_Default___.Abs(_dafny.IntOfInt64(138))), 1)
				(_nw100).ArraySet1((_599_v1).Difference(_599_v1), 2)
				(_nw100).ArraySet1((_dafny.MultiSetFromSeq(_623_v5)).Update(p0, Companion_Default___.Abs((_this).F21())), 3)
				(_nw100).ArraySet1(_599_v1, 4)
				(_nw100).ArraySet1(_599_v1, 5)
				(_nw100).ArraySet1(_dafny.MultiSetFromSeq(_623_v5), 6)
				(_nw100).ArraySet1(_599_v1, 7)
				(_nw100).ArraySet1(_599_v1, 8)
				(_nw100).ArraySet1(_599_v1, 9)
				(_nw100).ArraySet1(((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs((_this).F21()))).Difference(_599_v1), 10)
				(_nw100).ArraySet1(_599_v1, 11)
				(_nw100).ArraySet1((func() _dafny.MultiSet {
					if (_this).F20() {
						return _599_v1
					}
					return _599_v1
				})(), 12)
				(_nw100).ArraySet1((_626_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_627_v9).Cardinality()), _dafny.IntOfUint32((_626_v8).Cardinality()))).Uint32()).(_dafny.MultiSet), 13)
				(_nw100).ArraySet1((_dafny.MultiSetOf(Companion_Default___.Fm2((_this).F20(), (_this).F21(), (_this).F21(), globalState))).Union(_599_v1), 14)
				(_nw100).ArraySet1((_this).Fm1((_this).F21(), (_this).F21(), (_this).F20(), globalState), 15)
				(_nw100).ArraySet1(((_this).Fm1((_this).F21(), (p1).Cardinality(), true, globalState)).Intersection(_dafny.MultiSetOf((_this).F20(), false)), 16)
				(_nw100).ArraySet1((_600_v2).Dtor_cf0(), 17)
				(_nw100).ArraySet1(_dafny.MultiSetFromSeq(_623_v5), 18)
				(_nw100).ArraySet1((_this).Fm1((_this).F21(), (_this).F21(), p0, globalState), 19)
				(_nw100).ArraySet1((_this).Fm1((_this).F21(), (_this).F21(), (_this).F20(), globalState), 20)
				(_nw100).ArraySet1(_dafny.MultiSetOf((_this).F20(), (_this).F20(), !(true), false, p0), 21)
				(_nw100).ArraySet1((_dafny.MultiSetOf(false, (_this).F20())).Union(_599_v1), 22)
				(_nw100).ArraySet1((_626_v8).Select((Companion_Default___.SafeIndex((_599_v1).Cardinality(), _dafny.IntOfUint32((_626_v8).Cardinality()))).Uint32()).(_dafny.MultiSet), 23)
				(_nw100).ArraySet1(_599_v1, 24)
				(_nw100).ArraySet1((_600_v2).Dtor_cf0(), 25)
				_628_v10 = _nw100
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_628_v10), 0))
				_ = _index77
				(_628_v10).ArraySet1((_this).Fm1((_dafny.Zero).Minus((_this).F21()), (_this).F21(), Companion_Default___.Fm2(p0, (_this).F21(), (_this).F21(), globalState), globalState), (_index77).Int())
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_628_v10), 0))
				_ = _index78
				(_628_v10).ArraySet1(_599_v1, (_index78).Int())
			} else {
				var _629_v11 _dafny.Map
				_ = _629_v11
				_629_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index79
				((_this).F19()).ArraySet1(p0, (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index80
				var _rhs52 _dafny.Map = _629_v11
				_ = _rhs52
				var _rhs53 bool = _dafny.Companion_Sequence_.Equal(_598_v0, _dafny.Companion_Sequence_.Concatenate(_598_v0, _598_v0))
				_ = _rhs53
				var _rhs54 bool = true
				_ = _rhs54
				var _rhs55 bool = !((_this).F20()) || (!((_this).F20()))
				_ = _rhs55
				var _lhs45 *GlobalState = globalState
				_ = _lhs45
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				var _lhs47 _dafny.Array = (_this).F19()
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _lhs48
				_629_v11 = _rhs52
				_lhs45.F0 = _rhs53
				_lhs46.F0 = _rhs54
				(_lhs47).ArraySet1(_rhs55, (_lhs48).Int())
				var _630_v12 _dafny.Sequence
				_ = _630_v12
				_630_v12 = _dafny.SeqOf((_this).F21())
				var _631_v13 _dafny.Map
				_ = _631_v13
				_631_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _630_v12)
				var _632_v14 _dafny.Set
				_ = _632_v14
				_632_v14 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_631_v13).Contains(p0) {
						return (_631_v13).Get(p0).(_dafny.Sequence)
					}
					return _dafny.SeqOf((_this).F21(), (_this).F21())
				})(), _630_v12), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(292))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}(func(_633_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(430)
				})), (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(292))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg25 _dafny.Int) interface{} {
						return coer25(arg25)
					}
				}(func(_634_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(430)
				}))).Cardinality()))).Uint32(), (_this).F21()), _630_v12))
				var _rhs56 bool = (((_629_v11).Merge(_629_v11)).Cardinality()).Cmp(((_dafny.Zero).Minus((_this).F21())).Minus((_this).F21())) > 0
				_ = _rhs56
				var _rhs57 _dafny.Set = _dafny.SetOf(Companion_Default___.Fm3(globalState))
				_ = _rhs57
				var _rhs58 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F21(), (_621_v3).Cardinality())
				_ = _rhs58
				var _lhs49 *GlobalState = globalState
				_ = _lhs49
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				_lhs49.F0 = _rhs56
				_632_v14 = _rhs57
				_lhs50.F13 = _rhs58
				var _635_v15 _dafny.Array
				_ = _635_v15
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
				_ = _nw101
				_635_v15 = _nw101
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_635_v15), 0))
				_ = _index81
				(_635_v15).ArraySet1(_dafny.IntOfInt64(789), (_index81).Int())
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_635_v15), 0))
				_ = _index82
				var _rhs59 bool = p0
				_ = _rhs59
				var _rhs60 _dafny.Int = (_this).F21()
				_ = _rhs60
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				var _lhs52 _dafny.Array = _635_v15
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_635_v15), 0))
				_ = _lhs53
				_lhs51.F0 = _rhs59
				(_lhs52).ArraySet1(_rhs60, (_lhs53).Int())
				(globalState).F0 = (_this).F20()
				_598_v0 = _598_v0
			}
			var _636_v16 _dafny.Map
			_ = _636_v16
			_636_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)
			var _637_v17 _dafny.CodePoint
			_ = _637_v17
			_637_v17 = _dafny.CodePoint('j')
			(globalState).F0 = (func() bool {
				if (_636_v16).Contains(_637_v17) {
					return (_636_v16).Get(_637_v17).(bool)
				}
				return (_this).F20()
			})()
		} else {
			var _638_v18 _dafny.CodePoint
			_ = _638_v18
			_638_v18 = _dafny.CodePoint('x')
			var _639_v19 _dafny.Map
			_ = _639_v19
			_639_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_638_v18, p0)
			var _640_v20 D0
			_ = _640_v20
			_640_v20 = Companion_D0_.Create_DC1_(p0, (_dafny.Zero).Minus((_599_v1).Cardinality()), _dafny.IntOfInt64(360))
			var _641_v21 _dafny.MultiSet
			_ = _641_v21
			_641_v21 = _dafny.MultiSetOf(_640_v20)
			var _642_v22 _dafny.Sequence
			_ = _642_v22
			_642_v22 = _dafny.SeqOf(p0, (_this).F20())
			var _643_v23 D0
			_ = _643_v23
			_643_v23 = Companion_D0_.Create_DC2_((func() bool {
				if (_639_v19).Contains(Companion_Default___.Fm4(globalState)) {
					return (_639_v19).Get(Companion_Default___.Fm4(globalState)).(bool)
				}
				return (_this).F20()
			})(), (_641_v21).Cardinality(), _642_v22, _dafny.IntOfUint32((_598_v0).Cardinality()), (_this).F21())
			_643_v23 = _643_v23
			(globalState).F0 = !(!((_this).F20())) || ((func() bool {
				if (_this).F20() {
					return (_this).F20()
				}
				return (_this).F20()
			})())
			var _644_v24 _dafny.Array
			_ = _644_v24
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw102
			_644_v24 = _nw102
			(globalState).F3 = _644_v24
			if (_this).F20() {
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index83
				((_this).F19()).ArraySet1((true) == (!(p0)), (_index83).Int())
				var _645_v25 _dafny.Map
				_ = _645_v25
				_645_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F21())
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _index84
				var _rhs61 bool = p0
				_ = _rhs61
				var _rhs62 bool = !(_645_v25).Equals(_645_v25)
				_ = _rhs62
				var _rhs63 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F20()), _642_v22)
				_ = _rhs63
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				var _lhs55 _dafny.Array = (_this).F19()
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))
				_ = _lhs56
				_lhs54.F0 = _rhs61
				(_lhs55).ArraySet1(_rhs62, (_lhs56).Int())
				_642_v22 = _rhs63
				var _646_v26 _dafny.Array
				_ = _646_v26
				var _nwElement0_19 bool = ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool)
				_ = _nwElement0_19
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(23))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_19, 0)
				(_nw103).ArraySet1((_this).F20(), 1)
				(_nw103).ArraySet1(false, 2)
				(_nw103).ArraySet1(true, 3)
				(_nw103).ArraySet1(false, 4)
				(_nw103).ArraySet1(p0, 5)
				(_nw103).ArraySet1(false, 6)
				(_nw103).ArraySet1((_this).F20(), 7)
				(_nw103).ArraySet1((_this).F20(), 8)
				(_nw103).ArraySet1(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), 9)
				(_nw103).ArraySet1(p0, 10)
				(_nw103).ArraySet1(p0, 11)
				(_nw103).ArraySet1(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), 12)
				(_nw103).ArraySet1(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), 13)
				(_nw103).ArraySet1((_this).F20(), 14)
				(_nw103).ArraySet1(true, 15)
				(_nw103).ArraySet1((_this).F20(), 16)
				(_nw103).ArraySet1((_this).F20(), 17)
				(_nw103).ArraySet1(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), 18)
				(_nw103).ArraySet1(p0, 19)
				(_nw103).ArraySet1(((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), 20)
				(_nw103).ArraySet1(p0, 21)
				(_nw103).ArraySet1((_this).F20(), 22)
				_646_v26 = _nw103
				var _647_v27 *C1
				_ = _647_v27
				var _nw104 *C1 = New_C1_()
				_ = _nw104
				_nw104.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ucn"), _598_v0), _646_v26, false)
				_647_v27 = _nw104
				var _648_v28 _dafny.Map
				_ = _648_v28
				_648_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v24, _dafny.CodePoint('e'))
				_648_v28 = (_648_v28).Update(_644_v24, _dafny.CodePoint('g'))
				var _649_v29 D4
				_ = _649_v29
				_649_v29 = Companion_D4_.Create_DC13_()
				var _650_v30 D4
				_ = _650_v30
				_650_v30 = Companion_D4_.Create_DC15_(_649_v29)
				_650_v30 = _650_v30
				var _651_v31 *C0
				_ = _651_v31
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__((_this).F21())
				_651_v31 = _nw105
				var _652_v32 *C2
				_ = _652_v32
				var _nw106 *C2 = New_C2_()
				_ = _nw106
				_nw106.Ctor__(_651_v31, ((_this).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen(((_this).F19()), 0))).Int()).(bool), _646_v26, (_this).F20())
				_652_v32 = _nw106
			} else {
				var _653_v33 _dafny.MultiSet
				_ = _653_v33
				_653_v33 = _dafny.MultiSetOf((_this).F21())
				var _654_v34 _dafny.Map
				_ = _654_v34
				_654_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_653_v33).Update((_this).F21(), Companion_Default___.Abs((_this).F21()))).Cardinality()), (_this).F20())
				var _655_v35 _dafny.Map
				_ = _655_v35
				_655_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _dafny.IntOfUint32((Companion_Default___.Fm8(_598_v0, _654_v34, (_this).F21(), globalState)).Cardinality()))
				var _656_v36 *C0
				_ = _656_v36
				var _nw107 *C0 = New_C0_()
				_ = _nw107
				_nw107.Ctor__((_this).F21())
				_656_v36 = _nw107
				var _657_v37 T0
				_ = _657_v37
				var _nw108 *C2 = New_C2_()
				_ = _nw108
				_nw108.Ctor__(_656_v36, (_this).F20(), (_this).F19(), p0)
				_657_v37 = _nw108
				var _658_v38 _dafny.Sequence
				_ = _658_v38
				_658_v38 = _dafny.SeqOf(_657_v37, _657_v37, _657_v37, _657_v37, _657_v37)
				var _659_v39 _dafny.MultiSet
				_ = _659_v39
				_659_v39 = _dafny.MultiSetOf(_dafny.IntOfUint32((_642_v22).Cardinality()), (_this).F21(), (_655_v35).Cardinality(), _dafny.IntOfUint32((_658_v38).Cardinality()))
				(globalState).F13 = ((_659_v39).Cardinality()).Minus((_656_v36).F24())
				var _660_v40 _dafny.Map
				_ = _660_v40
				_660_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(126))
				_660_v40 = (_660_v40).Update((_this).F20(), (_656_v36).F24())
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_644_v24), 0))
				_ = _index85
				(_644_v24).ArraySet1(Companion_Default___.SafeModuloInt((_this).F21(), (_this).F21()), (_index85).Int())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen(((_657_v37).F19()), 0))
				_ = _index86
				((_657_v37).F19()).ArraySet1(p0, (_index86).Int())
				var _661_v41 _dafny.Map
				_ = _661_v41
				_661_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_657_v37).F20(), _598_v0)
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_644_v24), 0))
				_ = _index87
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen(((_657_v37).F19()), 0))
				_ = _index88
				var _rhs64 _dafny.Int = _dafny.IntOfInt64(948)
				_ = _rhs64
				var _rhs65 _dafny.Int = _dafny.IntOfInt64(542)
				_ = _rhs65
				var _rhs66 _dafny.Int = (_this).F21()
				_ = _rhs66
				var _rhs67 bool = Companion_Default___.Fm7(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_661_v41).Contains(Companion_Default___.Fm7((_this).F21(), (_this).F21(), (_this).F21(), globalState)) {
						return (_661_v41).Get(Companion_Default___.Fm7((_this).F21(), (_this).F21(), (_this).F21(), globalState)).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("bxledcrb")
				})()).Cardinality())), (func() _dafny.Int {
					if (_655_v35).Contains((_this).F21()) {
						return (_655_v35).Get((_this).F21()).(_dafny.Int)
					}
					return (_this).F21()
				})()), (_656_v36).F24(), (_this).F21(), globalState)
				_ = _rhs67
				var _rhs68 _dafny.Int = (_dafny.Zero).Minus((_this).F21())
				_ = _rhs68
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				var _lhs59 _dafny.Array = _644_v24
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_644_v24), 0))
				_ = _lhs60
				var _lhs61 _dafny.Array = (_657_v37).F19()
				_ = _lhs61
				var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen(((_657_v37).F19()), 0))
				_ = _lhs62
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				_lhs57.F18 = _rhs64
				_lhs58.F13 = _rhs65
				(_lhs59).ArraySet1(_rhs66, (_lhs60).Int())
				(_lhs61).ArraySet1(_rhs67, (_lhs62).Int())
				_lhs63.F18 = _rhs68
				var _662_v42 D3
				_ = _662_v42
				_662_v42 = Companion_D3_.Create_DC9_(_598_v0)
				var _663_v43 _dafny.Array
				_ = _663_v43
				var _nwElement0_20 _dafny.Int = Companion_Default___.Fm0(((_657_v37).F19()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen(((_657_v37).F19()), 0))).Int()).(bool), (_this).F20(), globalState)
				_ = _nwElement0_20
				var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(17))
				_ = _nw109
				(_nw109).ArraySet1(_nwElement0_20, 0)
				(_nw109).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xw")).Cardinality()), 1)
				(_nw109).ArraySet1(_dafny.IntOfInt64(836), 2)
				(_nw109).ArraySet1(_dafny.IntOfInt64(350), 3)
				(_nw109).ArraySet1((_this).F21(), 4)
				(_nw109).ArraySet1(_dafny.IntOfUint32((_598_v0).Cardinality()), 5)
				(_nw109).ArraySet1((_644_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_644_v24), 0))).Int()).(_dafny.Int), 6)
				(_nw109).ArraySet1((_this).F21(), 7)
				(_nw109).ArraySet1((_this).F21(), 8)
				(_nw109).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(233))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_664_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_665_i1 _dafny.Int) _dafny.CodePoint {
						return _664_v18
					}
				})(_638_v18)))).Cardinality()))), 9)
				(_nw109).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yqgtfnndn")).Cardinality()), 10)
				(_nw109).ArraySet1(Companion_Default___.SafeModuloInt((_656_v36).F24(), (_656_v36).F24()), 11)
				(_nw109).ArraySet1((_dafny.IntOfUint32(((_662_v42).Dtor_cf16()).Cardinality())).Plus((p1).Cardinality()), 12)
				(_nw109).ArraySet1((_644_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_644_v24), 0))).Int()).(_dafny.Int), 13)
				(_nw109).ArraySet1((_dafny.Zero).Minus((_this).F21()), 14)
				(_nw109).ArraySet1((_656_v36).F24(), 15)
				(_nw109).ArraySet1(((_this).F21()).Minus((_644_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(978), _dafny.ArrayLen((_644_v24), 0))).Int()).(_dafny.Int)), 16)
				_663_v43 = _nw109
				(globalState).F10 = _663_v43
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen(((_657_v37).F19()), 0))
				_ = _index89
				((_657_v37).F19()).ArraySet1(!((p0) == (Companion_Default___.Fm7((_this).F21(), (_dafny.SetOf((_657_v37).F20(), (_657_v37).F20(), false)).Cardinality(), (_656_v36).F24(), globalState))), (_index89).Int())
			}
			var _666_v44 D3
			_ = _666_v44
			_666_v44 = Companion_D3_.Create_DC9_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(990))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_667_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_668_i2 _dafny.Int) _dafny.CodePoint {
					return _667_v18
				}
			})(_638_v18))))
			var _pat_let_tv27 = _666_v44
			_ = _pat_let_tv27
			var _pat_let_tv28 = _598_v0
			_ = _pat_let_tv28
			var _pat_let_tv29 = _598_v0
			_ = _pat_let_tv29
			var _pat_let_tv30 = _638_v18
			_ = _pat_let_tv30
			var _669_v45 _dafny.Array
			_ = _669_v45
			var _nwElement0_21 D3 = _666_v44
			_ = _nwElement0_21
			var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(22))
			_ = _nw110
			(_nw110).ArraySet1(_nwElement0_21, 0)
			(_nw110).ArraySet1(_666_v44, 1)
			(_nw110).ArraySet1(_666_v44, 2)
			(_nw110).ArraySet1(_666_v44, 3)
			(_nw110).ArraySet1(_666_v44, 4)
			(_nw110).ArraySet1(_666_v44, 5)
			(_nw110).ArraySet1(Companion_D3_.Create_DC9_(_598_v0), 6)
			(_nw110).ArraySet1(_666_v44, 7)
			(_nw110).ArraySet1(_666_v44, 8)
			(_nw110).ArraySet1(func(_pat_let20_0 D3) D3 {
				return func(_670_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let21_0 _dafny.Sequence) D3 {
						return func(_671_dt__update_hcf16_h0 _dafny.Sequence) D3 {
							return Companion_D3_.Create_DC9_(_671_dt__update_hcf16_h0)
						}(_pat_let21_0)
					}((_pat_let_tv27).Dtor_cf16())
				}(_pat_let20_0)
			}(_666_v44), 9)
			(_nw110).ArraySet1(_666_v44, 10)
			(_nw110).ArraySet1(_666_v44, 11)
			(_nw110).ArraySet1(_666_v44, 12)
			(_nw110).ArraySet1(_666_v44, 13)
			(_nw110).ArraySet1(_666_v44, 14)
			(_nw110).ArraySet1(_666_v44, 15)
			(_nw110).ArraySet1(_666_v44, 16)
			(_nw110).ArraySet1((func() D3 {
				if p0 {
					return _666_v44
				}
				return _666_v44
			})(), 17)
			(_nw110).ArraySet1(_666_v44, 18)
			(_nw110).ArraySet1(_666_v44, 19)
			(_nw110).ArraySet1(func(_pat_let22_0 D3) D3 {
				return func(_672_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let23_0 _dafny.Sequence) D3 {
						return func(_673_dt__update_hcf16_h1 _dafny.Sequence) D3 {
							return Companion_D3_.Create_DC9_(_673_dt__update_hcf16_h1)
						}(_pat_let23_0)
					}(_dafny.Companion_Sequence_.Update(_pat_let_tv28, (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_pat_let_tv29).Cardinality()))).Uint32(), _pat_let_tv30))
				}(_pat_let22_0)
			}(_666_v44), 20)
			(_nw110).ArraySet1(_666_v44, 21)
			_669_v45 = _nw110
			_669_v45 = _669_v45
		}
		(globalState).F18 = (_this).F21()
		var _674_v46 *C0
		_ = _674_v46
		var _nw111 *C0 = New_C0_()
		_ = _nw111
		_nw111.Ctor__(_dafny.IntOfInt64(107))
		_674_v46 = _nw111
		var _675_v47 T0
		_ = _675_v47
		var _nw112 *C2 = New_C2_()
		_ = _nw112
		_nw112.Ctor__(_674_v46, (_this).F20(), (_this).F19(), (_this).F20())
		_675_v47 = _nw112
		var _676_v48 _dafny.Map
		_ = _676_v48
		_676_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_675_v47, (_this).F20())
		var _677_v49 _dafny.Map
		_ = _677_v49
		_677_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_676_v48).Contains(_675_v47) {
				return (_676_v48).Get(_675_v47).(bool)
			}
			return !(p0)
		})(), ((_674_v46).F24()).Cmp((_674_v46).F24()) == 0)
		var _678_v50 D4
		_ = _678_v50
		_678_v50 = Companion_D4_.Create_DC14_(p0, p1)
		var _pat_let_tv31 = p1
		_ = _pat_let_tv31
		_677_v49 = (_677_v49).Update((func(_pat_let24_0 D4) D4 {
			return func(_679_dt__update__tmp_h2 D4) D4 {
				return func(_pat_let25_0 _dafny.Set) D4 {
					return func(_680_dt__update_hcf26_h0 _dafny.Set) D4 {
						return Companion_D4_.Create_DC14_((_679_dt__update__tmp_h2).Dtor_cf25(), _680_dt__update_hcf26_h0)
					}(_pat_let25_0)
				}(_pat_let_tv31)
			}(_pat_let24_0)
		}(_678_v50)).Dtor_cf25(), p0)
	}
}
func (_this *C4) M1(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		(globalState).F13 = ((_this).F21()).Times((_this).F21())
		var _hi5 _dafny.Int = (_this).F21()
		_ = _hi5
		for _681_i0 := (_this).F21(); _681_i0.Cmp(_hi5) < 0; _681_i0 = _681_i0.Plus(_dafny.One) {
			var _682_v0 _dafny.Sequence
			_ = _682_v0
			_682_v0 = _dafny.UnicodeSeqOfUtf8Bytes("v")
			var _683_v1 *C1
			_ = _683_v1
			var _nw113 *C1 = New_C1_()
			_ = _nw113
			_nw113.Ctor__(_682_v0, (_this).F19(), (_this).F20())
			_683_v1 = _nw113
			var _684_v2 *C3
			_ = _684_v2
			var _nw114 *C3 = New_C3_()
			_ = _nw114
			_nw114.Ctor__(_dafny.Companion_Sequence_.Concatenate((_683_v1).F27(), _682_v0), (_this).F19())
			_684_v2 = _nw114
			var _685_v3 D2
			_ = _685_v3
			_685_v3 = Companion_D2_.Create_DC8_((func() _dafny.Int {
				if (_this).F20() {
					return (_this).F21()
				}
				return _681_i0
			})(), (_this).F21())
			var _686_v4 _dafny.Array
			_ = _686_v4
			var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
			_ = _nw115
			_686_v4 = _nw115
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))
			_ = _index90
			(_686_v4).ArraySet1(_681_i0, (_index90).Int())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_686_v4), 0))
			_ = _index91
			(_686_v4).ArraySet1(_681_i0, (_index91).Int())
			var _687_v5 _dafny.Sequence
			_ = _687_v5
			_687_v5 = _dafny.SeqOf((_this).F20(), (_this).F20())
			var _688_v6 _dafny.MultiSet
			_ = _688_v6
			_688_v6 = _dafny.MultiSetOf((_this).F21())
			var _pat_let_tv32 = _688_v6
			_ = _pat_let_tv32
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))
			_ = _index92
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_686_v4), 0))
			_ = _index93
			var _rhs69 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_687_v5, _687_v5)).Cardinality()))
			_ = _rhs69
			var _rhs70 D2 = func(_pat_let26_0 D2) D2 {
				return func(_689_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let27_0 _dafny.Int) D2 {
						return func(_690_dt__update_hcf15_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC8_((_689_dt__update__tmp_h0).Dtor_cf14(), _690_dt__update_hcf15_h0)
						}(_pat_let27_0)
					}(((_pat_let_tv32).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(28)))).Cardinality())
				}(_pat_let26_0)
			}(Companion_Default___.Fm23(globalState))
			_ = _rhs70
			var _rhs71 _dafny.Int = _dafny.IntOfUint32((_684_v2.F22).Cardinality())
			_ = _rhs71
			var _rhs72 _dafny.Int = Companion_Default___.SafeModuloInt(_681_i0, (_this).F21())
			_ = _rhs72
			var _rhs73 bool = (_this).F20()
			_ = _rhs73
			var _lhs64 *GlobalState = globalState
			_ = _lhs64
			var _lhs65 _dafny.Array = _686_v4
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))
			_ = _lhs66
			var _lhs67 _dafny.Array = _686_v4
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_686_v4), 0))
			_ = _lhs68
			var _lhs69 *GlobalState = globalState
			_ = _lhs69
			_lhs64.F13 = _rhs69
			_685_v3 = _rhs70
			(_lhs65).ArraySet1(_rhs71, (_lhs66).Int())
			(_lhs67).ArraySet1(_rhs72, (_lhs68).Int())
			_lhs69.F0 = _rhs73
			if (_this).F20() {
				var _691_v7 _dafny.Map
				_ = _691_v7
				_691_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_this).F20())
				var _692_v8 _dafny.Sequence
				_ = _692_v8
				_692_v8 = _dafny.SeqOf((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int), (_this).F21(), (_this).F21(), (_this).F21(), (_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int))
				var _693_v9 _dafny.Array
				_ = _693_v9
				var _nwElement0_22 bool = (_this).F20()
				_ = _nwElement0_22
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(28))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_22, 0)
				(_nw116).ArraySet1(!((func() bool {
					if (_691_v7).Contains((_this).F20()) {
						return (_691_v7).Get((_this).F20()).(bool)
					}
					return (_this).F20()
				})()), 1)
				(_nw116).ArraySet1(!_dafny.Companion_Sequence_.Equal(_684_v2.F22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(563))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}(func(_694_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))), 2)
				(_nw116).ArraySet1((_this).F20(), 3)
				(_nw116).ArraySet1((_this).F20(), 4)
				(_nw116).ArraySet1(!(true), 5)
				(_nw116).ArraySet1((_this).F20(), 6)
				(_nw116).ArraySet1(!((_dafny.MultiSetFromSeq(_692_v8)).IsSubsetOf(_dafny.MultiSetOf((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int)))), 7)
				(_nw116).ArraySet1((_this).F20(), 8)
				(_nw116).ArraySet1((_this).F20(), 9)
				(_nw116).ArraySet1((_this).F20(), 10)
				(_nw116).ArraySet1(false, 11)
				(_nw116).ArraySet1((_683_v1).Fm13(globalState), 12)
				(_nw116).ArraySet1((_this).F20(), 13)
				(_nw116).ArraySet1((_this).F20(), 14)
				(_nw116).ArraySet1((_this).F20(), 15)
				(_nw116).ArraySet1((_688_v6).IsDisjointFrom(_688_v6), 16)
				(_nw116).ArraySet1(((_this).F21()).Cmp(_681_i0) > 0, 17)
				(_nw116).ArraySet1((_this).F20(), 18)
				(_nw116).ArraySet1((_this).F20(), 19)
				(_nw116).ArraySet1((_this).F20(), 20)
				(_nw116).ArraySet1((_this).F20(), 21)
				(_nw116).ArraySet1((_this).F20(), 22)
				(_nw116).ArraySet1((_this).F20(), 23)
				(_nw116).ArraySet1((_this).F20(), 24)
				(_nw116).ArraySet1(!(true), 25)
				(_nw116).ArraySet1((_this).F20(), 26)
				(_nw116).ArraySet1((_this).F20(), 27)
				_693_v9 = _nw116
				_693_v9 = (_this).F19()
				var _695_v10 _dafny.CodePoint
				_ = _695_v10
				_695_v10 = _dafny.CodePoint('q')
				var _696_v12 _dafny.Sequence
				_ = _696_v12
				_696_v12 = _dafny.SeqOf(func() _dafny.Map {
					var _coll13 = _dafny.NewMapBuilder()
					_ = _coll13
					for _iter16 := _dafny.Iterate((_692_v8).Elements()); ; {
						_compr_13, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _697_v11 _dafny.Int
						_697_v11 = interface{}(_compr_13).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_692_v8, _697_v11) {
							_coll13.Add(Companion_Default___.SafeModuloInt(_697_v11, (_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int)), (_this).F20())
						}
					}
					return _coll13.ToMap()
				}())
				var _698_v13 _dafny.Map
				_ = _698_v13
				_698_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F20())
				var _699_v14 *C3
				_ = _699_v14
				var _nw117 *C3 = New_C3_()
				_ = _nw117
				_nw117.Ctor__(Companion_Default___.Fm8(Companion_Default___.Fm8(_dafny.SeqOf(_695_v10), (_696_v12).Select((Companion_Default___.SafeIndex(_681_i0, _dafny.IntOfUint32((_696_v12).Cardinality()))).Uint32()).(_dafny.Map), _dafny.IntOfUint32((_684_v2.F22).Cardinality()), globalState), _698_v13, (_this).F21(), globalState), _693_v9)
				_699_v14 = _nw117
				var _700_v15 _dafny.Set
				_ = _700_v15
				_700_v15 = _dafny.SetOf((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int))
				var _701_v16 _dafny.Map
				_ = _701_v16
				_701_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v10, (_700_v15).IsDisjointFrom(_dafny.SetOf((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_684_v2.F22).Cardinality()))))
				_701_v16 = _701_v16
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))
				_ = _index94
				(_686_v4).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-178)), (_index94).Int())
				var _702_v17 _dafny.Array
				_ = _702_v17
				var _nw118 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(2))
				_ = _nw118
				_702_v17 = _nw118
				var _703_v18 D1
				_ = _703_v18
				_703_v18 = Companion_D1_.Create_DC5_((_this).F20())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_702_v17), 0))
				_ = _index95
				(_702_v17).ArraySet1(_703_v18, (_index95).Int())
				var _704_v19 D0
				_ = _704_v19
				_704_v19 = Companion_D0_.Create_DC2_((_this).F20(), (_this).F21(), _687_v5, Companion_Default___.SafeDivisionInt((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(853)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_684_v2.F22).Cardinality())))
				var _705_v20 _dafny.Map
				_ = _705_v20
				_705_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int), _681_i0)
				var _706_v21 _dafny.MultiSet
				_ = _706_v21
				_706_v21 = _dafny.MultiSetOf(!_dafny.Companion_Sequence_.Equal((_683_v1).F27(), _684_v2.F22), (_687_v5).Select((Companion_Default___.SafeIndex(_681_i0, _dafny.IntOfUint32((_687_v5).Cardinality()))).Uint32()).(bool), !((_this).F20()), (_681_i0).Cmp((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int)) == 0, (_this).F20())
				var _707_v22 _dafny.MultiSet
				_ = _707_v22
				_707_v22 = _dafny.MultiSetOf(_695_v10, _695_v10)
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_702_v17), 0))
				_ = _index96
				var _rhs74 D1 = Companion_D1_.Create_DC5_((_this).F20())
				_ = _rhs74
				var _rhs75 D0 = _704_v19
				_ = _rhs75
				var _rhs76 _dafny.Map = (_705_v20).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_707_v22).Contains(_695_v10) {
						return (_707_v22).Multiplicity(_695_v10)
					}
					return _dafny.IntOfInt64(817)
				})(), _681_i0))
				_ = _rhs76
				var _rhs77 bool = (_this).F20()
				_ = _rhs77
				var _rhs78 _dafny.MultiSet = (_706_v21).Intersection((_706_v21).Intersection(_706_v21))
				_ = _rhs78
				var _lhs70 _dafny.Array = _702_v17
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_702_v17), 0))
				_ = _lhs71
				var _lhs72 *GlobalState = globalState
				_ = _lhs72
				(_lhs70).ArraySet1(_rhs74, (_lhs71).Int())
				_704_v19 = _rhs75
				_705_v20 = _rhs76
				_lhs72.F0 = _rhs77
				_706_v21 = _rhs78
			} else {
				r0 = _681_i0
				var _rhs79 _dafny.Sequence = _687_v5
				_ = _rhs79
				var _rhs80 bool = (_this).F20()
				_ = _rhs80
				var _lhs73 *GlobalState = globalState
				_ = _lhs73
				_687_v5 = _rhs79
				_lhs73.F0 = _rhs80
				(_684_v2).M3(globalState)
				var _708_v23 _dafny.Map
				_ = _708_v23
				_708_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("iqw"), (_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int))
				var _709_v24 _dafny.Map
				_ = _709_v24
				_709_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_683_v1).F27())
				var _710_v25 _dafny.Map
				_ = _710_v25
				_710_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int))
				var _711_v26 _dafny.MultiSet
				_ = _711_v26
				_711_v26 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int)), (_710_v25).Update((_this).F20(), _681_i0))
				var _712_v27 _dafny.Map
				_ = _712_v27
				_712_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-831), (func() _dafny.Int {
					if (_711_v26).Contains(_710_v25) {
						return (_711_v26).Multiplicity(_710_v25)
					}
					return _dafny.IntOfInt64(-915)
				})())
				_708_v23 = (_708_v23).Update(_dafny.Companion_Sequence_.Concatenate(_682_v0, (func() _dafny.Sequence {
					if (_709_v24).Contains(_681_i0) {
						return (_709_v24).Get(_681_i0).(_dafny.Sequence)
					}
					return _682_v0
				})()), ((_712_v27).Update((_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int), (_686_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_686_v4), 0))).Int()).(_dafny.Int))).Cardinality())
				var _713_v28 _dafny.CodePoint
				_ = _713_v28
				_713_v28 = _dafny.CodePoint('y')
				_713_v28 = _713_v28
			}
		}
		var _714_v29 _dafny.Array
		_ = _714_v29
		var _nwElement0_23 _dafny.Int = (_this).F21()
		_ = _nwElement0_23
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(4))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_23, 0)
		(_nw119).ArraySet1((_this).F21(), 1)
		(_nw119).ArraySet1((_this).F21(), 2)
		(_nw119).ArraySet1(Companion_Default___.Fm0((_this).F20(), (_this).F20(), globalState), 3)
		_714_v29 = _nw119
		var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
		_ = _index97
		(_714_v29).ArraySet1((_this).F21(), (_index97).Int())
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
		_ = _index98
		(_714_v29).ArraySet1((_this).F21(), (_index98).Int())
		var _715_v30 _dafny.CodePoint
		_ = _715_v30
		_715_v30 = _dafny.CodePoint('a')
		var _716_v31 D4
		_ = _716_v31
		_716_v31 = Companion_D4_.Create_DC12_(_715_v30)
		var _717_v32 _dafny.Sequence
		_ = _717_v32
		_717_v32 = _dafny.SeqOf(_715_v30, (_716_v31).Dtor_cf24())
		var _718_v33 _dafny.Map
		_ = _718_v33
		_718_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(968), false)
		var _719_v34 T0
		_ = _719_v34
		var _nw120 *C1 = New_C1_()
		_ = _nw120
		_nw120.Ctor__(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(_717_v32, _718_v33, (_this).F21(), globalState), _717_v32), (func() _dafny.Array {
			if false {
				return (_this).F19()
			}
			return (_this).F19()
		})(), (_this).F20())
		_719_v34 = _nw120
		var _720_v35 D0
		_ = _720_v35
		_720_v35 = Companion_D0_.Create_DC1_(true, (_this).F21(), (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))
		var _nw121 *C1 = New_C1_()
		_ = _nw121
		_nw121.Ctor__(_717_v32, (func() _dafny.Array {
			if (Companion_D5_.Create_DC17_((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qedpt")).Cardinality()), _720_v35, (_this).F20(), true)).Dtor_cf32() {
				return (_719_v34).F19()
			}
			return (_this).F19()
		})(), (_719_v34).F20())
		_719_v34 = _nw121
		var _hi6 _dafny.Int = (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)
		_ = _hi6
		for _721_i2 := (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-527))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_729_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_730_i3 _dafny.Int) _dafny.CodePoint {
				return _729_v30
			}
		})(_715_v30)))).Cardinality())).Minus((_this).F21()); _721_i2.Cmp(_hi6) < 0; _721_i2 = _721_i2.Plus(_dafny.One) {
			var _722_v36 _dafny.MultiSet
			_ = _722_v36
			_722_v36 = _dafny.MultiSetOf((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0((_719_v34).F20(), (_this).F20(), globalState))
			var _723_v37 _dafny.Map
			_ = _723_v37
			_723_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), _721_i2, globalState), _722_v36)
			_723_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_722_v36).Union(_722_v36))
			(globalState).F0 = (_this).F20()
			var _724_v38 _dafny.MultiSet
			_ = _724_v38
			_724_v38 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_717_v32, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_725_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_726_i4 _dafny.Int) _dafny.CodePoint {
					return _725_v30
				}
			})(_715_v30)))), _dafny.Companion_Sequence_.Concatenate(_717_v32, _717_v32), _717_v32, _717_v32, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_719_v34).F20() {
					return _717_v32
				}
				return _717_v32
			})(), (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_719_v34).F20() {
					return _717_v32
				}
				return _717_v32
			})()).Cardinality()))).Uint32(), _715_v30))
			_724_v38 = _724_v38
			var _727_v39 _dafny.Array
			_ = _727_v39
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw122
			_727_v39 = _nw122
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen(((_719_v34).F19()), 0))
			_ = _index99
			((_719_v34).F19()).ArraySet1(!((_this).F20()), (_index99).Int())
			var _728_v40 _dafny.MultiSet
			_ = _728_v40
			_728_v40 = _dafny.MultiSetOf((_719_v34).F20(), (_719_v34).F20(), false)
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
			_ = _index100
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen(((_719_v34).F19()), 0))
			_ = _index101
			var _rhs81 bool = (_719_v34).F20()
			_ = _rhs81
			var _rhs82 _dafny.Array = (_719_v34).F19()
			_ = _rhs82
			var _rhs83 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F21(), (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))
			_ = _rhs83
			var _rhs84 bool = (_728_v40).IsDisjointFrom((_728_v40).Update((_719_v34).F20(), Companion_Default___.Abs(_721_i2)))
			_ = _rhs84
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			var _lhs75 _dafny.Array = _714_v29
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
			_ = _lhs76
			var _lhs77 _dafny.Array = (_719_v34).F19()
			_ = _lhs77
			var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen(((_719_v34).F19()), 0))
			_ = _lhs78
			_lhs74.F0 = _rhs81
			_727_v39 = _rhs82
			(_lhs75).ArraySet1(_rhs83, (_lhs76).Int())
			(_lhs77).ArraySet1(_rhs84, (_lhs78).Int())
		}
		var _hi7 _dafny.Int = (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)
		_ = _hi7
		for _731_i5 := (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int); _731_i5.Cmp(_hi7) < 0; _731_i5 = _731_i5.Plus(_dafny.One) {
			if (((_this).F21()).Cmp((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)) <= 0) == ((_this).F20()) {
				var _732_v41 _dafny.Map
				_ = _732_v41
				_732_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_719_v34).F20(), (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))
				var _733_v42 _dafny.Map
				_ = _733_v42
				_733_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
				(globalState).F13 = (func() _dafny.Int {
					if (_732_v41).Contains((func() bool {
						if false {
							return true
						}
						return (_719_v34).F20()
					})()) {
						return (_732_v41).Get((func() bool {
							if false {
								return true
							}
							return (_719_v34).F20()
						})()).(_dafny.Int)
					}
					return ((_733_v42).Merge(_733_v42)).Cardinality()
				})()
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
				_ = _index102
				(_714_v29).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)), (_dafny.Zero).Minus((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))), (_index102).Int())
				(globalState).F0 = (_this).F20()
				var _734_v43 _dafny.Array
				_ = _734_v43
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_14
				var _nw123 _dafny.Array
				_ = _nw123
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw123 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) bool = (func(_735_v34 T0) func(_dafny.Int) bool {
						return func(_736_i6 _dafny.Int) bool {
							return (_735_v34).F20()
						}
					})(_719_v34)
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw123 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw123).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw123).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_734_v43 = _nw123
				var _737_v44 *C0
				_ = _737_v44
				var _nw124 *C0 = New_C0_()
				_ = _nw124
				_nw124.Ctor__((_this).F21())
				_737_v44 = _nw124
				var _738_v45 *C2
				_ = _738_v45
				var _nw125 *C2 = New_C2_()
				_ = _nw125
				_nw125.Ctor__(_737_v44, (_719_v34).F20(), _734_v43, (_719_v34).F20())
				_738_v45 = _nw125
				var _739_v46 _dafny.Array
				_ = _739_v46
				var _nw126 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(2))
				_ = _nw126
				_739_v46 = _nw126
				var _740_v47 _dafny.Array
				_ = _740_v47
				_740_v47 = _739_v46
				var _741_v48 _dafny.Array
				_ = _741_v48
				var _nwElement0_24 _dafny.Array = _740_v47
				_ = _nwElement0_24
				var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(18))
				_ = _nw127
				(_nw127).ArraySet1(_nwElement0_24, 0)
				(_nw127).ArraySet1(_739_v46, 1)
				(_nw127).ArraySet1(_739_v46, 2)
				(_nw127).ArraySet1(_740_v47, 3)
				(_nw127).ArraySet1(_740_v47, 4)
				(_nw127).ArraySet1(_740_v47, 5)
				(_nw127).ArraySet1(_740_v47, 6)
				(_nw127).ArraySet1(_740_v47, 7)
				(_nw127).ArraySet1(_740_v47, 8)
				(_nw127).ArraySet1(_740_v47, 9)
				(_nw127).ArraySet1(_740_v47, 10)
				(_nw127).ArraySet1(_740_v47, 11)
				(_nw127).ArraySet1(_740_v47, 12)
				(_nw127).ArraySet1(_739_v46, 13)
				(_nw127).ArraySet1((func() _dafny.Array {
					if (_719_v34).F20() {
						return _740_v47
					}
					return _739_v46
				})(), 14)
				(_nw127).ArraySet1(_739_v46, 15)
				(_nw127).ArraySet1(_740_v47, 16)
				(_nw127).ArraySet1((func() _dafny.Array {
					if (_719_v34).F20() {
						return _739_v46
					}
					return _739_v46
				})(), 17)
				_741_v48 = _nw127
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_741_v48), 0))
				_ = _index103
				(_741_v48).ArraySet1(_740_v47, (_index103).Int())
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
				_ = _index104
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_741_v48), 0))
				_ = _index105
				var _rhs85 _dafny.Array = _734_v43
				_ = _rhs85
				var _rhs86 *C2 = _738_v45
				_ = _rhs86
				var _rhs87 _dafny.Int = (_this).F21()
				_ = _rhs87
				var _rhs88 _dafny.Array = _740_v47
				_ = _rhs88
				var _lhs79 _dafny.Array = _714_v29
				_ = _lhs79
				var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
				_ = _lhs80
				var _lhs81 _dafny.Array = _741_v48
				_ = _lhs81
				var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(250), _dafny.ArrayLen((_741_v48), 0))
				_ = _lhs82
				_734_v43 = _rhs85
				_738_v45 = _rhs86
				(_lhs79).ArraySet1(_rhs87, (_lhs80).Int())
				(_lhs81).ArraySet1(_rhs88, (_lhs82).Int())
				_732_v41 = (_732_v41).Update(_dafny.Companion_Sequence_.IsPrefixOf(_717_v32, _dafny.UnicodeSeqOfUtf8Bytes("jnij")), (_737_v44).F24())
			} else {
				var _742_v49 _dafny.MultiSet
				_ = _742_v49
				_742_v49 = _dafny.MultiSetOf((_this).F20(), (_this).F20())
				var _743_v50 D0
				_ = _743_v50
				_743_v50 = Companion_D0_.Create_DC0_(_742_v49)
				_743_v50 = _743_v50
				(globalState).F0 = (_this).F20()
				var _744_v51 _dafny.Array
				_ = _744_v51
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw128
				_744_v51 = _nw128
				_744_v51 = _744_v51
				var _745_v52 _dafny.Map
				_ = _745_v52
				_745_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_719_v34).F20(), (_dafny.Zero).Minus((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)))
				_742_v49 = ((_742_v49).Update((_719_v34).F20(), Companion_Default___.Abs((_745_v52).Cardinality()))).Union(_742_v49)
				var _746_v53 *C0
				_ = _746_v53
				var _nw129 *C0 = New_C0_()
				_ = _nw129
				_nw129.Ctor__((Companion_Default___.Fm19(false, globalState)).Cardinality())
				_746_v53 = _nw129
			}
			var _747_v54 _dafny.Sequence
			_ = _747_v54
			_747_v54 = _dafny.SeqOf((_this).F20())
			var _748_v55 _dafny.Map
			_ = _748_v55
			_748_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_717_v32, _717_v32), (Companion_Default___.Fm0((_this).F20(), (_719_v34).F20(), globalState)).Plus((Companion_Default___.Fm10(_dafny.IntOfUint32((_747_v54).Cardinality()), _dafny.IntOfInt64(462), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-433))), globalState)).Cardinality()))
			var _rhs89 _dafny.Map = _748_v55
			_ = _rhs89
			var _rhs90 bool = (_719_v34).F20()
			_ = _rhs90
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			_748_v55 = _rhs89
			_lhs83.F0 = _rhs90
			var _749_v56 _dafny.Sequence
			_ = _749_v56
			_749_v56 = _dafny.SeqOf(_717_v32, _717_v32)
			var _750_v57 _dafny.Map
			_ = _750_v57
			_750_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v56, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(182))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_751_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_752_i7 _dafny.Int) _dafny.CodePoint {
					return _751_v30
				}
			})(_715_v30))))
			_750_v57 = (_750_v57).Update((func() _dafny.Sequence {
				if Companion_Default___.Fm9((_this).F20(), globalState) {
					return _749_v56
				}
				return _749_v56
			})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-291))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_753_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_754_i8 _dafny.Int) _dafny.CodePoint {
					return _753_v30
				}
			})(_715_v30))))
			if (_this).F20() {
				var _755_v58 _dafny.MultiSet
				_ = _755_v58
				_755_v58 = _dafny.MultiSetOf((_this).F20())
				var _756_v59 D0
				_ = _756_v59
				_756_v59 = Companion_D0_.Create_DC0_(_755_v58)
				var _pat_let_tv33 = _755_v58
				_ = _pat_let_tv33
				var _757_v60 _dafny.Map
				_ = _757_v60
				_757_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func(_pat_let28_0 D0) D0 {
					return func(_758_dt__update__tmp_h4 D0) D0 {
						return func(_pat_let29_0 _dafny.MultiSet) D0 {
							return func(_759_dt__update_hcf0_h0 _dafny.MultiSet) D0 {
								return Companion_D0_.Create_DC0_(_759_dt__update_hcf0_h0)
							}(_pat_let29_0)
						}(_pat_let_tv33)
					}(_pat_let28_0)
				}(_756_v59))
				var _760_v61 _dafny.Sequence
				_ = _760_v61
				_760_v61 = _dafny.SeqOf(_757_v60, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _756_v59))
				(globalState).F0 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_760_v61, _760_v61), _760_v61)
				(globalState).F18 = (_dafny.Zero).Minus(((_dafny.Zero).Minus(Companion_Default___.Fm0((_719_v34).F20(), true, globalState))).Minus((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)))
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
				_ = _index106
				(_714_v29).ArraySet1(_dafny.IntOfInt64(260), (_index106).Int())
				var _761_v62 _dafny.Array
				_ = _761_v62
				var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
				_ = _nw130
				_761_v62 = _nw130
				var _762_v63 _dafny.Map
				_ = _762_v63
				_762_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm9((_719_v34).F20(), globalState), _731_i5)
				var _763_v64 D0
				_ = _763_v64
				_763_v64 = Companion_D0_.Create_DC2_((_this).F20(), _731_i5, _747_v54, (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_this).F20())).Cardinality()))
				var _764_v65 _dafny.Sequence
				_ = _764_v65
				_764_v65 = _dafny.SeqOf(_763_v64)
				var _765_v66 _dafny.Set
				_ = _765_v66
				_765_v66 = _dafny.SetOf(Companion_D1_.Create_DC4_(_764_v65))
				var _766_v67 _dafny.Set
				_ = _766_v67
				_766_v67 = _dafny.SetOf((_this).F20(), (_this).F20(), (_719_v34).F20())
				var _767_v68 D4
				_ = _767_v68
				_767_v68 = Companion_D4_.Create_DC14_(!(false), _766_v67)
				_718_v33 = (_718_v33).Update((Companion_D3_.Create_DC10_(_761_v62, _762_v63, _765_v66, (_this).F21())).Dtor_cf20(), (_767_v68).Dtor_cf25())
				var _768_v69 _dafny.Array
				_ = _768_v69
				var _nwElement0_25 _dafny.CodePoint = _715_v30
				_ = _nwElement0_25
				var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(5))
				_ = _nw131
				(_nw131).ArraySet1CodePoint(_nwElement0_25, 0)
				(_nw131).ArraySet1CodePoint(_715_v30, 1)
				(_nw131).ArraySet1CodePoint(_dafny.CodePoint('o'), 2)
				(_nw131).ArraySet1CodePoint(_715_v30, 3)
				(_nw131).ArraySet1CodePoint((_717_v32).Select((Companion_Default___.SafeIndex((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_717_v32).Cardinality()))).Uint32()).(_dafny.CodePoint), 4)
				_768_v69 = _nw131
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_768_v69), 0))
				_ = _index107
				(_768_v69).ArraySet1CodePoint(_715_v30, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_768_v69), 0))
				_ = _index108
				(_768_v69).ArraySet1CodePoint(_dafny.CodePoint('u'), (_index108).Int())
			} else {
				var _769_v70 _dafny.MultiSet
				_ = _769_v70
				_769_v70 = _dafny.MultiSetOf((_731_i5).Cmp(_dafny.IntOfInt64(548)) < 0, (_719_v34).F20(), (_719_v34).F20(), ((_719_v34).F20()) == ((_this).F20()))
				_769_v70 = (_769_v70).Union((func() _dafny.MultiSet {
					if (_719_v34).F20() {
						return _769_v70
					}
					return _769_v70
				})())
				var _770_v71 *C0
				_ = _770_v71
				var _nw132 *C0 = New_C0_()
				_ = _nw132
				_nw132.Ctor__((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))
				_770_v71 = _nw132
				var _771_v72 _dafny.Map
				_ = _771_v72
				_771_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_719_v34).F20(), _770_v71)
				var _772_v73 _dafny.Sequence
				_ = _772_v73
				_772_v73 = _dafny.SeqOf(_770_v71)
				var _773_v74 _dafny.Map
				_ = _773_v74
				_773_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))
				var _774_v75 _dafny.Array
				_ = _774_v75
				var _nwElement0_26 _dafny.Map = _771_v72
				_ = _nwElement0_26
				var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(17))
				_ = _nw133
				(_nw133).ArraySet1(_nwElement0_26, 0)
				(_nw133).ArraySet1(_771_v72, 1)
				(_nw133).ArraySet1(_771_v72, 2)
				(_nw133).ArraySet1(_771_v72, 3)
				(_nw133).ArraySet1(_771_v72, 4)
				(_nw133).ArraySet1(_771_v72, 5)
				(_nw133).ArraySet1(_771_v72, 6)
				(_nw133).ArraySet1(_771_v72, 7)
				(_nw133).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_719_v34).F20(), (_772_v73).Select((Companion_Default___.SafeIndex((_773_v74).Cardinality(), _dafny.IntOfUint32((_772_v73).Cardinality()))).Uint32()).(*C0)), 8)
				(_nw133).ArraySet1(_771_v72, 9)
				(_nw133).ArraySet1(_771_v72, 10)
				(_nw133).ArraySet1(_771_v72, 11)
				(_nw133).ArraySet1(_771_v72, 12)
				(_nw133).ArraySet1(_771_v72, 13)
				(_nw133).ArraySet1(_771_v72, 14)
				(_nw133).ArraySet1(_771_v72, 15)
				(_nw133).ArraySet1(_771_v72, 16)
				_774_v75 = _nw133
				var _775_v76 D1
				_ = _775_v76
				_775_v76 = Companion_D1_.Create_DC4_(Companion_Default___.Fm24(globalState))
				var _776_v77 D0
				_ = _776_v77
				_776_v77 = Companion_D0_.Create_DC2_((_719_v34).F20(), (_770_v71).F24(), _747_v54, _dafny.IntOfInt64(604), (_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int))
				var _777_v78 _dafny.Sequence
				_ = _777_v78
				_777_v78 = _dafny.SeqOf(_776_v77)
				var _pat_let_tv34 = _777_v78
				_ = _pat_let_tv34
				var _778_v79 _dafny.Set
				_ = _778_v79
				_778_v79 = _dafny.SetOf(func(_pat_let30_0 D1) D1 {
					return func(_779_dt__update__tmp_h5 D1) D1 {
						return func(_pat_let31_0 _dafny.Sequence) D1 {
							return func(_780_dt__update_hcf10_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC4_(_780_dt__update_hcf10_h0)
							}(_pat_let31_0)
						}(_pat_let_tv34)
					}(_pat_let30_0)
				}(_775_v76), Companion_D1_.Create_DC4_(_777_v78))
				var _781_v80 D3
				_ = _781_v80
				_781_v80 = Companion_D3_.Create_DC10_(_774_v75, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F21()), _778_v79, _dafny.IntOfInt64(-727))
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
				_ = _index109
				(_714_v29).ArraySet1((_781_v80).Dtor_cf20(), (_index109).Int())
				(globalState).F0 = (func() bool {
					if !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_Default___.Fm25(globalState)), func() _dafny.Map {
						var _coll14 = _dafny.NewMapBuilder()
						_ = _coll14
						for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(527), _dafny.IntOfInt64(117))); ; {
							_compr_14, _ok17 := _iter17()
							if !_ok17 {
								break
							}
							var _782_v81 _dafny.Int
							_782_v81 = interface{}(_compr_14).(_dafny.Int)
							if ((_dafny.IntOfInt64(527)).Cmp(_782_v81) <= 0) && ((_782_v81).Cmp(_dafny.IntOfInt64(117)) < 0) {
								_coll14.Add((_782_v81).Times((_770_v71).F24()), (_719_v34).F20())
							}
						}
						return _coll14.ToMap()
					}()) {
						return (_719_v34).F20()
					}
					return (_719_v34).F20()
				})()
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))
				_ = _index110
				(_714_v29).ArraySet1((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int), (_index110).Int())
				_769_v70 = _dafny.MultiSetOf((_this).F20())
			}
		}
		r0 = Companion_Default___.SafeModuloInt(((_this).F21()).Plus((_714_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_714_v29), 0))).Int()).(_dafny.Int)), (_this).F21())
		return r0
	}
}
func (_this *C4) F21() _dafny.Int {
	{
		return _this._f21
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
