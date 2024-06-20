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
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-771), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(523), false))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Map, p1 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("lvxthx")
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (((_dafny.MultiSetOf(_dafny.IntOfInt64(693))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(533), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('j'))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.CodePoint
			_0_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('j'))).Contains(_0_v0) {
				_coll0.Add(_0_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vnemwohk")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(762))).Cardinality()))).Cardinality()))))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32(((Companion_D1_.Create_DC4_(!((Companion_D1_.Create_DC4_(false, true, _dafny.UnicodeSeqOfUtf8Bytes("rookixm"), _dafny.IntOfInt64(276), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Dtor_cf9()), true, _dafny.UnicodeSeqOfUtf8Bytes("kowjt"), _dafny.IntOfInt64(344), _dafny.IntOfInt64(452))).Dtor_cf10()).Cardinality()), _dafny.IntOfInt64(-249), _dafny.IntOfInt64(293), (Companion_D1_.Create_DC4_(false, true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(299))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.IntOfInt64(361), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-739))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()))).Dtor_cf12()))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("khaefkx"), Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(false, false, false, false, (Companion_D9_.Create_DC23_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((Companion_D10_.Create_DC25_(_dafny.SetOf(false, false))).Dtor_cf49()).Cardinality()), false)).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.CodePoint('o'))).Cardinality()), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(934), _dafny.SetOf(true)), _dafny.IntOfInt64(-844))).Dtor_cf45())).Cardinality(), _dafny.IntOfInt64(112)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(117)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(724)), _dafny.SeqOf(_dafny.IntOfInt64(-729))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((Companion_D5_.Create_DC15_(_dafny.MultiSetOf(false, !(!(true)), false, !(!(false))))).Dtor_cf31()).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false, false))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.CodePoint, p1 bool, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('w')
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false)).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(678))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ucej")).Cardinality())), _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(66))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.SetOf(true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nmj")).Cardinality())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(718))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-50))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, globalState *GlobalState) D1 {
	if (func() bool {
		if true {
			return false
		}
		return true
	})() {
		return Companion_D1_.Create_DC7_(Companion_D1_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(true, true)).Cardinality())).Cardinality(), false)))
	} else {
		return Companion_D1_.Create_DC7_(Companion_D1_.Create_DC5_(true, false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gc")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("r")).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-246)).Cmp(_dafny.IntOfInt64(36)) < 0, _dafny.SeqOf(_dafny.IntOfInt64(-678)))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('t'))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.CodePoint('q'))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(631))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.CodePoint('a'))
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('x'))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(814))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_4_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.CodePoint('b'))
	}))))
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) D6 {
	if (_dafny.MultiSetOf(true)).IsDisjointFrom(_dafny.MultiSetOf(!(!(!(!(true)))), false, !(false), false, false)) {
		return Companion_D6_.Create_DC17_(_dafny.SeqOf(true))
	} else {
		return Companion_D6_.Create_DC17_(_dafny.SeqOf(!(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_5_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(208)
	}))), _dafny.MultiSetOf(_dafny.IntOfInt64(-582), _dafny.IntOfInt64(809)), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_6_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	}))).Cardinality())))).Union((_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("seyfsdjs")).Cardinality()), _dafny.IntOfInt64(136))), _dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pwmmx")).Cardinality()))).Cardinality()), _dafny.MultiSetOf((_dafny.SetOf(true, true)).Cardinality()))).Intersection(_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnlhtu")).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 bool, globalState *GlobalState) D1 {
	var _source0 D4 = Companion_D4_.Create_DC13_((Companion_D4_.Create_DC13_(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('c'))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _7_v0 _dafny.CodePoint
			_7_v0 = interface{}(_compr_1).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('c')), _7_v0) {
				_coll1.Add(_7_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("penvbw")).Cardinality()))
			}
		}
		return _coll1.ToMap()
	}())).Dtor_cf30())
	_ = _source0
	if _source0.Is_DC14() {
		return Companion_D1_.Create_DC6_(true, _dafny.IntOfInt64(19), _dafny.IntOfInt64(-603), false)
	} else {
		var _8___mcc_h0 _dafny.Map = _source0.Get_().(D4_DC13).Cf30
		_ = _8___mcc_h0
		var _9_cf30 _dafny.Map = _8___mcc_h0
		_ = _9_cf30
		return Companion_D1_.Create_DC6_(false, (Companion_D1_.Create_DC5_(false, !(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-945))).Cardinality(), _dafny.IntOfInt64(284))).Dtor_cf15(), _dafny.IntOfInt64(915), !(!(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-357), _dafny.IntOfInt64(743))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _10_v0 _dafny.Int
			_10_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(-357)).Cmp(_10_v0) <= 0) && ((_10_v0).Cmp(_dafny.IntOfInt64(743)) < 0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_10_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Cardinality())))
			}
		}
		return _coll2.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC15_((_dafny.MultiSetOf(!(!(true)))).Intersection(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D2 {
	if true {
		return Companion_D2_.Create_DC8_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-344))))
	} else {
		return Companion_D2_.Create_DC8_(_dafny.MultiSetOf(_dafny.IntOfInt64(852)))
	}
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) bool {
	return (!((_dafny.SetOf(_dafny.IntOfInt64(-500))).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(458), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), false)).Cardinality()))))) == ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjjecxhfo")).Cardinality())).Cmp(_dafny.IntOfInt64(53)) > 0)
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _dafny.IntOfInt64(796))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.SetOf(true)).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, _dafny.Array, _dafny.Int, _dafny.Int) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 _dafny.Int = _dafny.Zero
	_ = r3
	var _11_v0 bool
	_ = _11_v0
	_11_v0 = true
	var _12_v1 _dafny.Sequence
	_ = _12_v1
	_12_v1 = _dafny.SeqOf(_11_v0, _11_v0, _11_v0, _11_v0, _11_v0)
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_12_v1).Cardinality())
	_ = _hi0
	for _13_i0 := p1; _13_i0.Cmp(_hi0) < 0; _13_i0 = _13_i0.Plus(_dafny.One) {
		var _14_v3 _dafny.Set
		_ = _14_v3
		_14_v3 = _dafny.SetOf(_dafny.IntOfInt64(-156), p0)
		var _15_v4 _dafny.Map
		_ = _15_v4
		_15_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_i0, _dafny.IntOfInt64(983))
		_11_v0 = !(func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_14_v3).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _16_v2 _dafny.Int
				_16_v2 = interface{}(_compr_3).(_dafny.Int)
				if (_14_v3).Contains(_16_v2) {
					_coll3.Add(Companion_Default___.SafeDivisionInt(_16_v2, _13_i0), p0)
				}
			}
			return _coll3.ToMap()
		}()).Equals(_15_v4)
		var _17_v5 _dafny.Array
		_ = _17_v5
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_18_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_19_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_19_i1, _18_p0)
				}
			})(p0)
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
		_17_v5 = _nw0
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(500), _dafny.ArrayLen((_17_v5), 0))
		_ = _index0
		(_17_v5).ArraySet1(p0, (_index0).Int())
		var _20_v6 _dafny.Sequence
		_ = _20_v6
		_20_v6 = _dafny.UnicodeSeqOfUtf8Bytes("riotpp")
		var _21_v7 D1
		_ = _21_v7
		_21_v7 = Companion_D1_.Create_DC4_(_11_v0, _11_v0, _20_v6, p1, p0)
		var _22_v8 _dafny.Sequence
		_ = _22_v8
		_22_v8 = _dafny.SeqOf(_12_v1)
		var _23_v9 *C3
		_ = _23_v9
		var _nw1 *C3 = New_C3_()
		_ = _nw1
		_nw1.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_22_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_22_v8).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.IntOfUint32(((_22_v8).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_22_v8).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _11_v0)).Cardinality()), _13_i0, _11_v0)
		_23_v9 = _nw1
		var _24_v10 _dafny.Sequence
		_ = _24_v10
		_24_v10 = _dafny.SeqOf(_23_v9, _23_v9, _23_v9)
		var _25_v11 _dafny.MultiSet
		_ = _25_v11
		_25_v11 = _dafny.MultiSetOf(_13_i0, p0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_24_v10).Cardinality()))), _13_i0, _dafny.IntOfUint32((_20_v6).Cardinality()))
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(500), _dafny.ArrayLen((_17_v5), 0))
		_ = _index1
		var _rhs0 _dafny.Int = (p1).Minus((_21_v7).Dtor_cf12())
		_ = _rhs0
		var _rhs1 _dafny.Int = ((Companion_D2_.Create_DC8_(_25_v11)).Dtor_cf22()).Cardinality()
		_ = _rhs1
		var _rhs2 _dafny.Int = _13_i0
		_ = _rhs2
		var _lhs0 _dafny.Array = _17_v5
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(500), _dafny.ArrayLen((_17_v5), 0))
		_ = _lhs1
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		r2 = _rhs1
		r3 = _rhs2
		r2 = ((p1).Times(_dafny.IntOfUint32((_20_v6).Cardinality()))).Plus(p0)
		var _26_v12 *C2
		_ = _26_v12
		var _nw2 *C2 = New_C2_()
		_ = _nw2
		_nw2.Ctor__(Companion_Default___.SafeDivisionInt(p1, (_17_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(500), _dafny.ArrayLen((_17_v5), 0))).Int()).(_dafny.Int)))
		_26_v12 = _nw2
	}
	r2 = p0
	var _27_v13 _dafny.CodePoint
	_ = _27_v13
	_27_v13 = _dafny.CodePoint('q')
	_27_v13 = func(_source1 _dafny.CodePoint) _dafny.CodePoint {
		var _28___mcc_h0 _dafny.CodePoint = _source1
		_ = _28___mcc_h0
		var _29_cf40 _dafny.CodePoint = _28___mcc_h0
		_ = _29_cf40
		return _dafny.CodePoint('c')
	}(_27_v13)
	var _30_v14 _dafny.Array
	_ = _30_v14
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_1
	var _nw3 _dafny.Array
	_ = _nw3
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw3 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = func(_31_i2 _dafny.Int) _dafny.Int {
			return (_31_i2).Plus(_dafny.IntOfInt64(834))
		}
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
	_30_v14 = _nw3
	var _32_v15 T0
	_ = _32_v15
	var _nw4 *C3 = New_C3_()
	_ = _nw4
	_nw4.Ctor__(p0, p0, _11_v0)
	_32_v15 = _nw4
	var _33_v16 _dafny.MultiSet
	_ = _33_v16
	_33_v16 = _dafny.MultiSetOf(_32_v15)
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_30_v14), 0))
	_ = _index2
	(_30_v14).ArraySet1((_33_v16).Cardinality(), (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_30_v14), 0))
	_ = _index3
	(_30_v14).ArraySet1(p0, (_index3).Int())
	var _34_v17 D4
	_ = _34_v17
	_34_v17 = Companion_D4_.Create_DC13_(Companion_Default___.Fm25(globalState))
	var _pat_let_tv0 = _32_v15
	_ = _pat_let_tv0
	var _pat_let_tv1 = _11_v0
	_ = _pat_let_tv1
	var _pat_let_tv2 = _32_v15
	_ = _pat_let_tv2
	_11_v0 = func(_source2 D4) bool {
		if _source2.Is_DC14() {
			return (_pat_let_tv0).F6()
		} else {
			var _35___mcc_h1 _dafny.Map = _source2.Get_().(D4_DC13).Cf30
			_ = _35___mcc_h1
			var _36_cf30 _dafny.Map = _35___mcc_h1
			_ = _36_cf30
			return !(_pat_let_tv1) || ((_pat_let_tv2).F6())
		}
	}(_34_v17)
	var _37_v18 _dafny.Array
	_ = _37_v18
	var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
	_ = _nw5
	_37_v18 = _nw5
	var _38_v19 _dafny.MultiSet
	_ = _38_v19
	_38_v19 = _dafny.MultiSetOf((_32_v15).F6(), (_32_v15).F6())
	var _39_v20 *C0
	_ = _39_v20
	var _nw6 *C0 = New_C0_()
	_ = _nw6
	_nw6.Ctor__(_dafny.IntOfInt64(867), (func() _dafny.Int {
		if (_38_v19).Contains(_11_v0) {
			return (_38_v19).Multiplicity(_11_v0)
		}
		return (_30_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_30_v14), 0))).Int()).(_dafny.Int)
	})())
	_39_v20 = _nw6
	var _40_v21 _dafny.Map
	_ = _40_v21
	_40_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v20, (_32_v15).F6())
	var _41_v22 _dafny.Map
	_ = _41_v22
	_41_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
		if _11_v0 {
			return _37_v18
		}
		return _37_v18
	})(), (((_40_v21).Update(_39_v20, (_32_v15).F6())).Cardinality()).Cmp(_dafny.IntOfInt64(831)) <= 0)
	_41_v22 = (_41_v22).Update(_37_v18, _11_v0)
	var _42_v23 _dafny.Sequence
	_ = _42_v23
	_42_v23 = _dafny.UnicodeSeqOfUtf8Bytes("tjkieshnd")
	r0 = _42_v23
	var _43_v24 _dafny.Map
	_ = _43_v24
	_43_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_32_v15).F6(), _11_v0)
	var _44_v25 _dafny.Sequence
	_ = _44_v25
	_44_v25 = _dafny.SeqOf(p0, (_38_v19).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_42_v23, _42_v23, _42_v23, _dafny.UnicodeSeqOfUtf8Bytes("eao"), _dafny.UnicodeSeqOfUtf8Bytes("iolhb"))).Cardinality()), (((_43_v24).Update((_32_v15).F6(), _11_v0)).Update(_11_v0, _11_v0)).Cardinality())
	var _nwElement0_0 _dafny.Int = (_30_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_30_v14), 0))).Int()).(_dafny.Int)
	_ = _nwElement0_0
	var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
	_ = _nw7
	(_nw7).ArraySet1(_nwElement0_0, 0)
	(_nw7).ArraySet1((_dafny.Zero).Minus(p1), 1)
	(_nw7).ArraySet1(_dafny.IntOfInt64(112), 2)
	(_nw7).ArraySet1(p0, 3)
	(_nw7).ArraySet1((p1).Times(Companion_Default___.Fm6((_39_v20).F12(), (_32_v15).F5(), (_32_v15).F6(), (_39_v20).F13(), globalState)), 4)
	(_nw7).ArraySet1(p1, 5)
	(_nw7).ArraySet1((_44_v25).Select((Companion_Default___.SafeIndex((_39_v20).F12(), _dafny.IntOfUint32((_44_v25).Cardinality()))).Uint32()).(_dafny.Int), 6)
	(_nw7).ArraySet1((_39_v20).F12(), 7)
	(_nw7).ArraySet1(((_30_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_30_v14), 0))).Int()).(_dafny.Int)).Plus((_39_v20).F13()), 8)
	(_nw7).ArraySet1((_39_v20).F12(), 9)
	r1 = _nw7
	r2 = (_30_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_30_v14), 0))).Int()).(_dafny.Int)
	r3 = _dafny.IntOfUint32((_dafny.SeqOf(_30_v14)).Cardinality())
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _45_v0 bool
	_ = _45_v0
	_45_v0 = false
	var _46_v1 _dafny.Int
	_ = _46_v1
	_46_v1 = _dafny.IntOfInt64(82)
	var _47_v2 _dafny.CodePoint
	_ = _47_v2
	_47_v2 = _dafny.CodePoint('d')
	var _48_v3 _dafny.MultiSet
	_ = _48_v3
	_48_v3 = _dafny.MultiSetOf(_47_v2)
	var _49_v4 _dafny.Map
	_ = _49_v4
	_49_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _46_v1)
	var _50_v5 _dafny.Map
	_ = _50_v5
	_50_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _45_v0)
	var _51_v6 _dafny.Map
	_ = _51_v6
	_51_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_50_v5).Cardinality(), _45_v0)
	var _52_v7 _dafny.Sequence
	_ = _52_v7
	_52_v7 = _dafny.UnicodeSeqOfUtf8Bytes("oxasa")
	var _53_v8 _dafny.Sequence
	_ = _53_v8
	_53_v8 = _dafny.SeqOf(_45_v0)
	var _54_v9 _dafny.Array
	_ = _54_v9
	var _nwElement0_1 _dafny.Int = _46_v1
	_ = _nwElement0_1
	var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(29))
	_ = _nw8
	(_nw8).ArraySet1(_nwElement0_1, 0)
	(_nw8).ArraySet1((func() _dafny.Int {
		if (_48_v3).Contains(_47_v2) {
			return (_48_v3).Multiplicity(_47_v2)
		}
		return _dafny.IntOfInt64(87)
	})(), 1)
	(_nw8).ArraySet1(_dafny.IntOfInt64(872), 2)
	(_nw8).ArraySet1((_49_v4).Cardinality(), 3)
	(_nw8).ArraySet1(_46_v1, 4)
	(_nw8).ArraySet1(_46_v1, 5)
	(_nw8).ArraySet1(_46_v1, 6)
	(_nw8).ArraySet1((_dafny.Zero).Minus(_46_v1), 7)
	(_nw8).ArraySet1((_51_v6).Cardinality(), 8)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_46_v1, _46_v1)).Cardinality()), 9)
	(_nw8).ArraySet1((_dafny.Zero).Minus(_46_v1), 10)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_52_v7).Cardinality()), 11)
	(_nw8).ArraySet1(_46_v1, 12)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_52_v7).Cardinality()), 13)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(325))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}((func(_55_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_56_i0 _dafny.Int) _dafny.CodePoint {
			return _55_v2
		}
	})(_47_v2)))).Cardinality()), 14)
	(_nw8).ArraySet1(_46_v1, 15)
	(_nw8).ArraySet1(_46_v1, 16)
	(_nw8).ArraySet1(_46_v1, 17)
	(_nw8).ArraySet1(_46_v1, 18)
	(_nw8).ArraySet1(_46_v1, 19)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_53_v8).Cardinality()), 20)
	(_nw8).ArraySet1(_46_v1, 21)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(47))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}((func(_57_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_58_i1 _dafny.Int) _dafny.CodePoint {
			return _57_v2
		}
	})(_47_v2)))).Cardinality()), 22)
	(_nw8).ArraySet1(_46_v1, 23)
	(_nw8).ArraySet1(_46_v1, 24)
	(_nw8).ArraySet1((_dafny.Zero).Minus(_46_v1), 25)
	(_nw8).ArraySet1(_46_v1, 26)
	(_nw8).ArraySet1(_46_v1, 27)
	(_nw8).ArraySet1(_46_v1, 28)
	_54_v9 = _nw8
	var _59_v10 _dafny.Map
	_ = _59_v10
	_59_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _54_v9)
	var _60_v11 _dafny.Set
	_ = _60_v11
	_60_v11 = _dafny.SetOf(_46_v1)
	var _61_v12 _dafny.Sequence
	_ = _61_v12
	_61_v12 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(475))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}((func(_62_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_63_i2 _dafny.Int) _dafny.CodePoint {
			return _62_v2
		}
	})(_47_v2))))
	var _64_globalState *GlobalState
	_ = _64_globalState
	var _nw9 *GlobalState = New_GlobalState_()
	_ = _nw9
	_nw9.Ctor__(_dafny.IntOfInt64(766), _59_v10, _dafny.SeqOf(_60_v11), _61_v12, true)
	_64_globalState = _nw9
	var _65_v13 _dafny.Sequence
	_ = _65_v13
	var _66_v14 _dafny.Array
	_ = _66_v14
	var _67_v15 _dafny.Int
	_ = _67_v15
	var _68_v16 _dafny.Int
	_ = _68_v16
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 _dafny.Array
	_ = _out1
	var _out2 _dafny.Int
	_ = _out2
	var _out3 _dafny.Int
	_ = _out3
	_out0, _out1, _out2, _out3 = Companion_Default___.M0((_dafny.IntOfInt64(861)).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}((func(_69_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_70_i3 _dafny.Int) _dafny.CodePoint {
			return _69_v2
		}
	})(_47_v2)))).Cardinality())), _46_v1, _64_globalState)
	_65_v13 = _out0
	_66_v14 = _out1
	_67_v15 = _out2
	_68_v16 = _out3
	var _71_v17 _dafny.MultiSet
	_ = _71_v17
	_71_v17 = _dafny.MultiSetOf(_68_v16, _67_v15, _67_v15)
	_68_v16 = Companion_Default___.SafeModuloInt(_68_v16, ((_71_v17).Intersection(_71_v17)).Cardinality())
	var _72_v18 _dafny.Array
	_ = _72_v18
	var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
	_ = _nw10
	_72_v18 = _nw10
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))
	_ = _index4
	(_72_v18).ArraySet1(_71_v17, (_index4).Int())
	var _73_v19 _dafny.Sequence
	_ = _73_v19
	_73_v19 = _dafny.SeqOf(_71_v17)
	var _74_v20 _dafny.Sequence
	_ = _74_v20
	_74_v20 = _dafny.SeqOf(_dafny.IntOfInt64(720))
	var _pat_let_tv3 = _52_v7
	_ = _pat_let_tv3
	var _pat_let_tv4 = _49_v4
	_ = _pat_let_tv4
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))
	_ = _index5
	(_72_v18).ArraySet1(((_71_v17).Intersection((_73_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_65_v13).Cardinality()), _dafny.IntOfUint32((_73_v19).Cardinality()))).Uint32()).(_dafny.MultiSet))).Intersection((_dafny.MultiSetFromSeq(_74_v20)).Union(_dafny.MultiSetFromSeq((func(_pat_let0_0 D0) D0 {
		return func(_75_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 _dafny.Sequence) D0 {
				return func(_76_dt__update_hcf1_h0 _dafny.Sequence) D0 {
					return func(_pat_let2_0 _dafny.Map) D0 {
						return func(_77_dt__update_hcf3_h0 _dafny.Map) D0 {
							return Companion_D0_.Create_DC1_(_76_dt__update_hcf1_h0, (_75_dt__update__tmp_h0).Dtor_cf2(), _77_dt__update_hcf3_h0, (_75_dt__update__tmp_h0).Dtor_cf4())
						}(_pat_let2_0)
					}(_pat_let_tv4)
				}(_pat_let1_0)
			}(_pat_let_tv3)
		}(_pat_let0_0)
	}(Companion_D0_.Create_DC1_(_65_v13, _67_v15, _49_v4, _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_60_v11).Cardinality(), _46_v1), (Companion_Default___.SafeIndex(_67_v15, _dafny.IntOfUint32((_dafny.SeqOf((_60_v11).Cardinality(), _46_v1)).Cardinality()))).Uint32(), _67_v15)))).Dtor_cf4()))), (_index5).Int())
	var _78_v21 _dafny.Map
	_ = _78_v21
	_78_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v16, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(125), _68_v16))
	var _79_v22 D0
	_ = _79_v22
	_79_v22 = Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("vewema"), _46_v1, _49_v4, _74_v20)
	_78_v21 = (_78_v21).Update((_67_v15).Minus((_79_v22).Dtor_cf2()), (_dafny.Zero).Minus((_dafny.IntOfInt64(741)).Times(_68_v16)))
	_67_v15 = _68_v16
	var _80_v23 _dafny.Array
	_ = _80_v23
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_2
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Map = (func(_81_v5 _dafny.Map) func(_dafny.Int) _dafny.Map {
			return func(_82_i4 _dafny.Int) _dafny.Map {
				return _81_v5
			}
		})(_50_v5)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw11 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw11).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw11).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_80_v23 = _nw11
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_80_v23), 0))
	_ = _index6
	(_80_v23).ArraySet1((_50_v5).Merge(_50_v5), (_index6).Int())
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_80_v23), 0))
	_ = _index7
	(_80_v23).ArraySet1((_50_v5).Merge((_50_v5).Merge(_50_v5)), (_index7).Int())
	_47_v2 = _dafny.CodePoint('y')
	var _83_v24 _dafny.Array
	_ = _83_v24
	var _nw12 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
	_ = _nw12
	_83_v24 = _nw12
	var _84_v25 _dafny.Map
	_ = _84_v25
	_84_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v24, !(_45_v0))
	_84_v25 = (_84_v25).Update(_83_v24, !(!(_45_v0) || (_45_v0)))
	if ((_60_v11).Cardinality()).Cmp(_46_v1) >= 0 {
		var _85_v26 _dafny.Map
		_ = _85_v26
		_85_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _74_v20)
		var _86_v27 D0
		_ = _86_v27
		_86_v27 = Companion_D0_.Create_DC2_(_85_v26, !(_45_v0))
		var _source3 D0 = _86_v27
		_ = _source3
		if _source3.Is_DC1() {
			var _87___mcc_h0 _dafny.Sequence = _source3.Get_().(D0_DC1).Cf1
			_ = _87___mcc_h0
			var _88___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
			_ = _88___mcc_h1
			var _89___mcc_h2 _dafny.Map = _source3.Get_().(D0_DC1).Cf3
			_ = _89___mcc_h2
			var _90___mcc_h3 _dafny.Sequence = _source3.Get_().(D0_DC1).Cf4
			_ = _90___mcc_h3
			var _91_cf4 _dafny.Sequence = _90___mcc_h3
			_ = _91_cf4
			var _92_cf3 _dafny.Map = _89___mcc_h2
			_ = _92_cf3
			var _93_cf2 _dafny.Int = _88___mcc_h1
			_ = _93_cf2
			var _94_cf1 _dafny.Sequence = _87___mcc_h0
			_ = _94_cf1
			_45_v0 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-964))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_95_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_96_i5 _dafny.Int) _dafny.CodePoint {
					return _95_v2
				}
			})(_47_v2))), (Companion_Default___.SafeIndex(_68_v16, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-964))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_97_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_98_i5 _dafny.Int) _dafny.CodePoint {
					return _97_v2
				}
			})(_47_v2)))).Cardinality()))).Uint32(), _47_v2), _dafny.UnicodeSeqOfUtf8Bytes("l")), (Companion_Default___.SafeIndex(_67_v15, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-964))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_99_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_100_i5 _dafny.Int) _dafny.CodePoint {
					return _99_v2
				}
			})(_47_v2))), (Companion_Default___.SafeIndex(_68_v16, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-964))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_101_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_102_i5 _dafny.Int) _dafny.CodePoint {
					return _101_v2
				}
			})(_47_v2)))).Cardinality()))).Uint32(), _47_v2), _dafny.UnicodeSeqOfUtf8Bytes("l"))).Cardinality()))).Uint32(), _47_v2), _52_v7)
			var _103_v28 *C0
			_ = _103_v28
			var _nw13 *C0 = New_C0_()
			_ = _nw13
			_nw13.Ctor__(_93_cf2, (_dafny.Zero).Minus(_46_v1))
			_103_v28 = _nw13
			_93_cf2 = Companion_Default___.SafeDivisionInt((_103_v28).F13(), _93_cf2)
			_67_v15 = (_68_v16).Minus((_103_v28).F13())
		} else if _source3.Is_DC2() {
			var _104___mcc_h4 _dafny.Map = _source3.Get_().(D0_DC2).Cf5
			_ = _104___mcc_h4
			var _105___mcc_h5 bool = _source3.Get_().(D0_DC2).Cf6
			_ = _105___mcc_h5
			var _106_cf6 bool = _105___mcc_h5
			_ = _106_cf6
			var _107_cf5 _dafny.Map = _104___mcc_h4
			_ = _107_cf5
			_67_v15 = _68_v16
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_83_v24), 0))
			_ = _index8
			(_83_v24).ArraySet1(_106_cf6, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_54_v9), 0))
			_ = _index9
			(_54_v9).ArraySet1(_67_v15, (_index9).Int())
			var _108_v29 *C0
			_ = _108_v29
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_46_v1, _67_v15)
			_108_v29 = _nw14
			var _109_v30 _dafny.Map
			_ = _109_v30
			_109_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v29, true)
			var _110_v31 _dafny.Map
			_ = _110_v31
			_110_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_109_v30).Cardinality(), _49_v4)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_83_v24), 0))
			_ = _index10
			(_83_v24).ArraySet1((_67_v15).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm5(_49_v4, _110_v31, _64_globalState)).Cardinality())) > 0, (_index10).Int())
			var _111_v32 _dafny.Set
			_ = _111_v32
			_111_v32 = _dafny.SetOf(true)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_83_v24), 0))
			_ = _index11
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))
			_ = _index12
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_54_v9), 0))
			_ = _index13
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_83_v24), 0))
			_ = _index14
			var _rhs3 bool = (_45_v0) == (false)
			_ = _rhs3
			var _rhs4 _dafny.MultiSet = (_71_v17).Union((_72_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))).Int()).(_dafny.MultiSet))
			_ = _rhs4
			var _rhs5 _dafny.Int = (_108_v29).F12()
			_ = _rhs5
			var _rhs6 _dafny.Array = _83_v24
			_ = _rhs6
			var _rhs7 bool = !((func() bool {
				if false {
					return (_111_v32).IsSubsetOf(_111_v32)
				}
				return _45_v0
			})())
			_ = _rhs7
			var _lhs2 _dafny.Array = _83_v24
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_83_v24), 0))
			_ = _lhs3
			var _lhs4 _dafny.Array = _72_v18
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))
			_ = _lhs5
			var _lhs6 _dafny.Array = _54_v9
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_54_v9), 0))
			_ = _lhs7
			var _lhs8 _dafny.Array = _83_v24
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_83_v24), 0))
			_ = _lhs9
			(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
			(_lhs4).ArraySet1(_rhs4, (_lhs5).Int())
			(_lhs6).ArraySet1(_rhs5, (_lhs7).Int())
			_83_v24 = _rhs6
			(_lhs8).ArraySet1(_rhs7, (_lhs9).Int())
			_47_v2 = Companion_Default___.Fm11(_47_v2, (!((_83_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_83_v24), 0))).Int()).(bool))) || ((_83_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_83_v24), 0))).Int()).(bool)), (_108_v29).Fm8((_83_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_83_v24), 0))).Int()).(bool), _52_v7, _47_v2, _64_globalState), _64_globalState)
			var _112_v33 T0
			_ = _112_v33
			var _nw15 *C4 = New_C4_()
			_ = _nw15
			_nw15.Ctor__((_108_v29).F12(), (_54_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_54_v9), 0))).Int()).(_dafny.Int), _106_cf6)
			_112_v33 = _nw15
			_112_v33 = _112_v33
		} else {
			var _113___mcc_h6 _dafny.Sequence = _source3.Get_().(D0_DC0).Cf0
			_ = _113___mcc_h6
			var _114_cf0 _dafny.Sequence = _113___mcc_h6
			_ = _114_cf0
			var _115_v34 _dafny.Sequence
			_ = _115_v34
			var _116_v35 _dafny.Array
			_ = _116_v35
			var _117_v36 _dafny.Int
			_ = _117_v36
			var _118_v37 _dafny.Int
			_ = _118_v37
			var _out4 _dafny.Sequence
			_ = _out4
			var _out5 _dafny.Array
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0(_67_v15, _46_v1, _64_globalState)
			_115_v34 = _out4
			_116_v35 = _out5
			_117_v36 = _out6
			_118_v37 = _out7
			var _119_v38 _dafny.Array
			_ = _119_v38
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw16
			_119_v38 = _nw16
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_119_v38), 0))
			_ = _index15
			(_119_v38).ArraySet1(_65_v13, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(373), _dafny.ArrayLen((_119_v38), 0))
			_ = _index16
			(_119_v38).ArraySet1(_65_v13, (_index16).Int())
			var _120_v39 _dafny.Array
			_ = _120_v39
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_3
			var _nw17 _dafny.Array
			_ = _nw17
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw17 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = func(_121_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw17 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw17).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw17).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_120_v39 = _nw17
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_120_v39), 0))
			_ = _index17
			(_120_v39).ArraySet1CodePoint(_47_v2, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_120_v39), 0))
			_ = _index18
			(_120_v39).ArraySet1CodePoint(_47_v2, (_index18).Int())
			_67_v15 = _46_v1
		}
		var _122_v40 _dafny.CodePoint
		_ = _122_v40
		_122_v40 = _47_v2
		_122_v40 = _47_v2
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw18
		_66_v14 = _nw18
		_47_v2 = Companion_Default___.Fm11(_47_v2, _45_v0, !(_45_v0) || (Companion_Default___.Fm24(_64_globalState)), _64_globalState)
		var _123_v41 T0
		_ = _123_v41
		var _nw19 *C1 = New_C1_()
		_ = _nw19
		_nw19.Ctor__(_67_v15, _67_v15, _67_v15, _45_v0)
		_123_v41 = _nw19
		var _124_v42 _dafny.Array
		_ = _124_v42
		var _nwElement0_2 T0 = (func() T0 {
			if _45_v0 {
				return _123_v41
			}
			return _123_v41
		})()
		_ = _nwElement0_2
		var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(3))
		_ = _nw20
		(_nw20).ArraySet1(_nwElement0_2, 0)
		(_nw20).ArraySet1(_123_v41, 1)
		(_nw20).ArraySet1(_123_v41, 2)
		_124_v42 = _nw20
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_124_v42), 0))
		_ = _index19
		(_124_v42).ArraySet1(_123_v41, (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_124_v42), 0))
		_ = _index20
		(_124_v42).ArraySet1((func() T0 {
			if (_123_v41).F6() {
				return _123_v41
			}
			return _123_v41
		})(), (_index20).Int())
	} else {
		var _125_v43 *C3
		_ = _125_v43
		var _nw21 *C3 = New_C3_()
		_ = _nw21
		_nw21.Ctor__(_68_v16, ((_dafny.Zero).Minus(_dafny.IntOfInt64(-469))).Minus(_46_v1), _45_v0)
		_125_v43 = _nw21
		_68_v16 = (_125_v43).F8()
		var _126_v44 _dafny.Set
		_ = _126_v44
		_126_v44 = _dafny.SetOf(_45_v0)
		var _127_v45 _dafny.Map
		_ = _127_v45
		_127_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v44, _52_v7)
		var _128_v46 T0
		_ = _128_v46
		var _nw22 *C1 = New_C1_()
		_ = _nw22
		_nw22.Ctor__((_125_v43).F8(), _dafny.IntOfInt64(431), _67_v15, _45_v0)
		_128_v46 = _nw22
		var _129_v47 _dafny.Set
		_ = _129_v47
		_129_v47 = _dafny.SetOf(_128_v46, _128_v46)
		_45_v0 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_127_v45).Contains(_126_v44) {
				return (_127_v45).Get(_126_v44).(_dafny.Sequence)
			}
			return _65_v13
		})(), (Companion_Default___.SafeIndex((_129_v47).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_127_v45).Contains(_126_v44) {
				return (_127_v45).Get(_126_v44).(_dafny.Sequence)
			}
			return _65_v13
		})()).Cardinality()))).Uint32(), _47_v2), _47_v2)
		var _rhs8 bool = (func() bool {
			if _45_v0 {
				return true
			}
			return (_128_v46).F6()
		})()
		_ = _rhs8
		var _rhs9 _dafny.Int = ((_67_v15).Plus(_46_v1)).Times((_dafny.Zero).Minus((_dafny.IntOfUint32((_53_v8).Cardinality())).Minus(_dafny.IntOfUint32((_53_v8).Cardinality()))))
		_ = _rhs9
		_45_v0 = _rhs8
		_68_v16 = _rhs9
		var _rhs10 _dafny.CodePoint = _47_v2
		_ = _rhs10
		var _rhs11 _dafny.Int = (_dafny.Zero).Minus(_68_v16)
		_ = _rhs11
		_47_v2 = _rhs10
		_68_v16 = _rhs11
	}
	_46_v1 = Companion_Default___.SafeModuloInt((_74_v20).Select((Companion_Default___.SafeIndex(_67_v15, _dafny.IntOfUint32((_74_v20).Cardinality()))).Uint32()).(_dafny.Int), _67_v15)
	if Companion_Default___.Fm24(_64_globalState) {
		_45_v0 = _45_v0
		var _130_v49 D1
		_ = _130_v49
		_130_v49 = Companion_D1_.Create_DC3_(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SetOf(_67_v15)).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _131_v48 _dafny.Int
				_131_v48 = interface{}(_compr_4).(_dafny.Int)
				if (_dafny.SetOf(_67_v15)).Contains(_131_v48) {
					_coll4.Add(Companion_Default___.SafeModuloInt(_131_v48, _dafny.IntOfInt64(236)), _45_v0)
				}
			}
			return _coll4.ToMap()
		}())
		var _132_v50 *C3
		_ = _132_v50
		var _nw23 *C3 = New_C3_()
		_ = _nw23
		_nw23.Ctor__(_dafny.IntOfInt64(-302), _dafny.IntOfUint32((_74_v20).Cardinality()), _45_v0)
		_132_v50 = _nw23
		var _133_v51 T0
		_ = _133_v51
		var _nw24 *C1 = New_C1_()
		_ = _nw24
		_nw24.Ctor__(_68_v16, _67_v15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v49, _132_v50)).Cardinality(), _45_v0)
		_133_v51 = _nw24
		var _134_v52 D5
		_ = _134_v52
		_134_v52 = Companion_D5_.Create_DC16_(_133_v51, (_133_v51).F6(), _68_v16)
		var _135_v53 *C0
		_ = _135_v53
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__((Companion_Default___.Fm6((_dafny.MultiSetOf(_45_v0)).Cardinality(), _46_v1, _45_v0, _46_v1, _64_globalState)).Plus(Companion_Default___.Fm6(_dafny.IntOfUint32((_74_v20).Cardinality()), _68_v16, _45_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("puoy")).Cardinality()), _64_globalState)), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_68_v16), (_134_v52).Dtor_cf34()))
		_135_v53 = _nw25
		_47_v2 = Companion_Default___.Fm11(_47_v2, (_45_v0) && ((_133_v51).F6()), (_133_v51).F6(), _64_globalState)
		_47_v2 = _47_v2
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_54_v9), 0))
		_ = _index21
		(_54_v9).ArraySet1(((_135_v53).F13()).Minus((_135_v53).F13()), (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_54_v9), 0))
		_ = _index22
		var _rhs12 _dafny.Sequence = _65_v13
		_ = _rhs12
		var _rhs13 _dafny.Int = ((_135_v53).F12()).Plus(_67_v15)
		_ = _rhs13
		var _rhs14 _dafny.Int = _46_v1
		_ = _rhs14
		var _lhs10 _dafny.Array = _54_v9
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_54_v9), 0))
		_ = _lhs11
		_65_v13 = _rhs12
		_46_v1 = _rhs13
		(_lhs10).ArraySet1(_rhs14, (_lhs11).Int())
	} else {
		var _136_v54 _dafny.Array
		_ = _136_v54
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
		_ = _nw26
		_136_v54 = _nw26
		_136_v54 = _136_v54
		_83_v24 = _83_v24
		_45_v0 = _45_v0
		var _137_v55 _dafny.Map
		_ = _137_v55
		_137_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _74_v20)
		var _pat_let_tv5 = _45_v0
		_ = _pat_let_tv5
		var _source4 D3 = func(_pat_let3_0 D3) D3 {
			return func(_138_dt__update__tmp_h1 D3) D3 {
				return func(_pat_let4_0 bool) D3 {
					return func(_139_dt__update_hcf28_h0 bool) D3 {
						return Companion_D3_.Create_DC12_(_139_dt__update_hcf28_h0, (_138_dt__update__tmp_h1).Dtor_cf29())
					}(_pat_let4_0)
				}(_pat_let_tv5)
			}(_pat_let3_0)
		}(Companion_D3_.Create_DC12_((Companion_D0_.Create_DC2_(_137_v55, Companion_Default___.Fm24(_64_globalState))).Dtor_cf6(), (_71_v17).Cardinality()))
		_ = _source4
		if _source4.Is_DC12() {
			var _140___mcc_h7 bool = _source4.Get_().(D3_DC12).Cf28
			_ = _140___mcc_h7
			var _141___mcc_h8 _dafny.Int = _source4.Get_().(D3_DC12).Cf29
			_ = _141___mcc_h8
			var _142_cf29 _dafny.Int = _141___mcc_h8
			_ = _142_cf29
			var _143_cf28 bool = _140___mcc_h7
			_ = _143_cf28
			var _144_v56 _dafny.Map
			_ = _144_v56
			_144_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("skjvslqs"), _67_v15)
			var _145_v57 _dafny.Map
			_ = _145_v57
			_145_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v56, _83_v24)
			var _146_v58 _dafny.Sequence
			_ = _146_v58
			_146_v58 = _dafny.SeqOf(_83_v24)
			_145_v57 = (_145_v57).Update(_144_v56, (_146_v58).Select((Companion_Default___.SafeIndex(_142_cf29, _dafny.IntOfUint32((_146_v58).Cardinality()))).Uint32()).(_dafny.Array))
			var _147_v59 _dafny.Sequence
			_ = _147_v59
			var _148_v60 _dafny.Array
			_ = _148_v60
			var _149_v61 _dafny.Int
			_ = _149_v61
			var _150_v62 _dafny.Int
			_ = _150_v62
			var _out8 _dafny.Sequence
			_ = _out8
			var _out9 _dafny.Array
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0(_67_v15, _46_v1, _64_globalState)
			_147_v59 = _out8
			_148_v60 = _out9
			_149_v61 = _out10
			_150_v62 = _out11
			var _151_v63 *C0
			_ = _151_v63
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(_46_v1, _68_v16)
			_151_v63 = _nw27
			var _152_v64 _dafny.Map
			_ = _152_v64
			_152_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v63, ((_72_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))).Int()).(_dafny.MultiSet)).Cardinality())
			var _153_v65 _dafny.MultiSet
			_ = _153_v65
			_153_v65 = _dafny.MultiSetOf(_152_v64)
			_51_v6 = (func() _dafny.Map {
				if Companion_Default___.Fm24(_64_globalState) {
					return (_51_v6).Update(_142_cf29, _45_v0)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_153_v65).Contains(_152_v64) {
						return (_153_v65).Multiplicity(_152_v64)
					}
					return _149_v61
				})(), _45_v0)
			})()
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_54_v9), 0))
			_ = _index23
			(_54_v9).ArraySet1(_142_cf29, (_index23).Int())
			var _154_v66 _dafny.Set
			_ = _154_v66
			_154_v66 = _dafny.SetOf(_45_v0)
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_54_v9), 0))
			_ = _index24
			(_54_v9).ArraySet1((_154_v66).Cardinality(), (_index24).Int())
		} else {
			var _155___mcc_h9 _dafny.Array = _source4.Get_().(D3_DC11).Cf27
			_ = _155___mcc_h9
			var _156_cf27 _dafny.Array = _155___mcc_h9
			_ = _156_cf27
			var _157_v67 _dafny.Map
			_ = _157_v67
			_157_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v11).Intersection(_60_v11), !(Companion_Default___.Fm24(_64_globalState)))
			_157_v67 = (_157_v67).Update(_60_v11, _45_v0)
			_46_v1 = (((_79_v22).Dtor_cf3()).Merge(Companion_Default___.Fm13(_dafny.UnicodeSeqOfUtf8Bytes("ptcfinu"), _68_v16, (_73_v19).Select((Companion_Default___.SafeIndex(_67_v15, _dafny.IntOfUint32((_73_v19).Cardinality()))).Uint32()).(_dafny.MultiSet), _64_globalState))).Cardinality()
			var _158_v68 _dafny.Map
			_ = _158_v68
			_158_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _68_v16))
			var _159_v69 *C0
			_ = _159_v69
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__((_67_v15).Plus(_dafny.IntOfInt64(722)), _dafny.IntOfUint32((Companion_Default___.Fm5(_49_v4, _158_v68, _64_globalState)).Cardinality()))
			_159_v69 = _nw28
			var _160_v70 T0
			_ = _160_v70
			var _nw29 *C3 = New_C3_()
			_ = _nw29
			_nw29.Ctor__((_159_v69).F12(), _68_v16, _45_v0)
			_160_v70 = _nw29
			var _rhs15 *C0 = _159_v69
			_ = _rhs15
			var _rhs16 T0 = _160_v70
			_ = _rhs16
			_159_v69 = _rhs15
			_160_v70 = _rhs16
			var _161_v71 _dafny.Array
			_ = _161_v71
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw30
			_161_v71 = _nw30
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_161_v71), 0))
			_ = _index25
			(_161_v71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_52_v7, _52_v7), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_161_v71), 0))
			_ = _index26
			(_161_v71).ArraySet1(_52_v7, (_index26).Int())
		}
		var _162_v72 _dafny.Sequence
		_ = _162_v72
		var _163_v73 _dafny.Array
		_ = _163_v73
		var _164_v74 _dafny.Int
		_ = _164_v74
		var _165_v75 _dafny.Int
		_ = _165_v75
		var _out12 _dafny.Sequence
		_ = _out12
		var _out13 _dafny.Array
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out12, _out13, _out14, _out15 = Companion_Default___.M0(Companion_Default___.SafeModuloInt(_46_v1, _67_v15), _68_v16, _64_globalState)
		_162_v72 = _out12
		_163_v73 = _out13
		_164_v74 = _out14
		_165_v75 = _out15
	}
	_54_v9 = _66_v14
	var _166_v76 *C4
	_ = _166_v76
	var _nw31 *C4 = New_C4_()
	_ = _nw31
	_nw31.Ctor__(_68_v16, _67_v15, _45_v0)
	_166_v76 = _nw31
	var _167_v77 _dafny.Sequence
	_ = _167_v77
	_167_v77 = _dafny.SeqOf(_166_v76)
	var _168_v78 D9
	_ = _168_v78
	_168_v78 = Companion_D9_.Create_DC22_(_167_v77)
	var _169_v79 _dafny.MultiSet
	_ = _169_v79
	_169_v79 = _dafny.MultiSetOf(_166_v76)
	var _170_v81 D2
	_ = _170_v81
	_170_v81 = Companion_D2_.Create_DC10_(Companion_D2_.Create_DC8_((_71_v17).Update(_46_v1, Companion_Default___.Abs((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_52_v7).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _171_v80 _dafny.CodePoint
			_171_v80 = interface{}(_compr_5).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_52_v7, _171_v80) {
				_coll5.Add(_171_v80, !(_45_v0))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality())))))
	var _172_v82 D2
	_ = _172_v82
	_172_v82 = Companion_D2_.Create_DC8_((_72_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))).Int()).(_dafny.MultiSet))
	var _pat_let_tv6 = _167_v77
	_ = _pat_let_tv6
	var _source5 D2 = (func() D2 {
		if (_169_v79).IsProperSubsetOf(_dafny.MultiSetFromSeq((func(_pat_let5_0 D9) D9 {
			return func(_173_dt__update__tmp_h2 D9) D9 {
				return func(_pat_let6_0 _dafny.Sequence) D9 {
					return func(_174_dt__update_hcf42_h0 _dafny.Sequence) D9 {
						return Companion_D9_.Create_DC22_(_174_dt__update_hcf42_h0)
					}(_pat_let6_0)
				}(_pat_let_tv6)
			}(_pat_let5_0)
		}(_168_v78)).Dtor_cf42())) {
			return _170_v81
		}
		return Companion_D2_.Create_DC10_(_172_v82)
	})()
	_ = _source5
	if _source5.Is_DC9() {
		var _175___mcc_h10 bool = _source5.Get_().(D2_DC9).Cf23
		_ = _175___mcc_h10
		var _176___mcc_h11 _dafny.Int = _source5.Get_().(D2_DC9).Cf24
		_ = _176___mcc_h11
		var _177___mcc_h12 _dafny.Int = _source5.Get_().(D2_DC9).Cf25
		_ = _177___mcc_h12
		var _178_cf25 _dafny.Int = _177___mcc_h12
		_ = _178_cf25
		var _179_cf24 _dafny.Int = _176___mcc_h11
		_ = _179_cf24
		var _180_cf23 bool = _175___mcc_h10
		_ = _180_cf23
		_46_v1 = _67_v15
		_65_v13 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-45))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_181_v7 _dafny.Sequence, _182_v76 *C4) func(_dafny.Int) _dafny.CodePoint {
			return func(_183_i7 _dafny.Int) _dafny.CodePoint {
				return (_181_v7).Select((Companion_Default___.SafeIndex(_182_v76.F7, _dafny.IntOfUint32((_181_v7).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		})(_52_v7, _166_v76)))
		_180_cf23 = _180_cf23
		var _184_v83 *C0
		_ = _184_v83
		var _nw32 *C0 = New_C0_()
		_ = _nw32
		_nw32.Ctor__(_46_v1, _67_v15)
		_184_v83 = _nw32
		var _185_v84 _dafny.Map
		_ = _185_v84
		_185_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _184_v83)
		var _186_v85 *C2
		_ = _186_v85
		var _nw33 *C2 = New_C2_()
		_ = _nw33
		_nw33.Ctor__((_185_v84).Cardinality())
		_186_v85 = _nw33
	} else if _source5.Is_DC8() {
		var _187___mcc_h13 _dafny.MultiSet = _source5.Get_().(D2_DC8).Cf22
		_ = _187___mcc_h13
		var _188_cf22 _dafny.MultiSet = _187___mcc_h13
		_ = _188_cf22
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_66_v14), 0))
		_ = _index27
		(_66_v14).ArraySet1((_166_v76.F7).Plus((_dafny.Zero).Minus(_166_v76.F7)), (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_66_v14), 0))
		_ = _index28
		(_66_v14).ArraySet1(_166_v76.F7, (_index28).Int())
		(_166_v76).M1(_166_v76.F7, _64_globalState)
		var _189_v86 _dafny.Array
		_ = _189_v86
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw34
		_189_v86 = _nw34
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_189_v86), 0))
		_ = _index29
		(_189_v86).ArraySet1(_dafny.Companion_Sequence_.Update(_53_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_53_v8).Cardinality()), _dafny.IntOfUint32((_53_v8).Cardinality()))).Uint32(), _45_v0), (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_189_v86), 0))
		_ = _index30
		(_189_v86).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if _45_v0 {
				return _dafny.SeqOf(_45_v0, true)
			}
			return _53_v8
		})(), _53_v8), (_index30).Int())
		var _190_v87 *C0
		_ = _190_v87
		var _nw35 *C0 = New_C0_()
		_ = _nw35
		_nw35.Ctor__(_46_v1, _dafny.IntOfUint32((_53_v8).Cardinality()))
		_190_v87 = _nw35
		var _191_v88 _dafny.Map
		_ = _191_v88
		_191_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_190_v87), (_190_v87).F13())
		var _192_v89 _dafny.Set
		_ = _192_v89
		_192_v89 = _dafny.SetOf(_190_v87)
		var _193_v90 _dafny.Map
		_ = _193_v90
		_193_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_191_v88).Update(_192_v89, (_78_v21).Cardinality())).Update(_192_v89, (_60_v11).Cardinality()), !(!(_45_v0)))
		_193_v90 = (_193_v90).Update(_191_v88, !((_190_v87).Fm8(_45_v0, (_79_v22).Dtor_cf1(), _47_v2, _64_globalState)))
	} else {
		var _194___mcc_h14 D2 = _source5.Get_().(D2_DC10).Cf26
		_ = _194___mcc_h14
		var _195_cf26 D2 = _194___mcc_h14
		_ = _195_cf26
		_67_v15 = Companion_Default___.Fm6(_67_v15, Companion_Default___.SafeModuloInt(_166_v76.F7, _166_v76.F7), !((_166_v76).Fm0(false, _64_globalState)), ((_72_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))).Int()).(_dafny.MultiSet)).Cardinality(), _64_globalState)
		var _196_v91 _dafny.Array
		_ = _196_v91
		var _nwElement0_3 *C4 = _166_v76
		_ = _nwElement0_3
		var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(4))
		_ = _nw36
		(_nw36).ArraySet1(_nwElement0_3, 0)
		(_nw36).ArraySet1(_166_v76, 1)
		(_nw36).ArraySet1(_166_v76, 2)
		(_nw36).ArraySet1(_166_v76, 3)
		_196_v91 = _nw36
		var _197_v92 D3
		_ = _197_v92
		_197_v92 = Companion_D3_.Create_DC11_(_66_v14)
		var _198_v93 _dafny.Map
		_ = _198_v93
		_198_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_196_v91, _197_v92)
		_198_v93 = (_198_v93).Update(_196_v91, _197_v92)
		_78_v21 = (_78_v21).Update(_67_v15, (func() _dafny.Int {
			if true {
				return _166_v76.F7
			}
			return _67_v15
		})())
		var _199_v94 _dafny.Map
		_ = _199_v94
		_199_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v0, _83_v24)
		var _200_v95 T0
		_ = _200_v95
		var _nw37 *C3 = New_C3_()
		_ = _nw37
		_nw37.Ctor__(_68_v16, _46_v1, _45_v0)
		_200_v95 = _nw37
		var _201_v96 _dafny.Set
		_ = _201_v96
		_201_v96 = _dafny.SetOf(_200_v95)
		var _202_v97 _dafny.Map
		_ = _202_v97
		_202_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_199_v94, _201_v96)
		_202_v97 = (_202_v97).Update(_199_v94, _201_v96)
	}
	var _203_i8 _dafny.Int
	_ = _203_i8
	_203_i8 = _dafny.Zero
	{
		for (_166_v76).Fm0((_46_v1).Cmp(_dafny.IntOfInt64(383)) <= 0, _64_globalState) {
			{
				if (_203_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_203_i8 = (_203_i8).Plus(_dafny.One)
				var _204_v98 _dafny.Set
				_ = _204_v98
				_204_v98 = _dafny.SetOf(_83_v24, _83_v24)
				var _205_v99 _dafny.Sequence
				_ = _205_v99
				var _206_v100 _dafny.Array
				_ = _206_v100
				var _207_v101 _dafny.Int
				_ = _207_v101
				var _208_v102 _dafny.Int
				_ = _208_v102
				var _out16 _dafny.Sequence
				_ = _out16
				var _out17 _dafny.Array
				_ = _out17
				var _out18 _dafny.Int
				_ = _out18
				var _out19 _dafny.Int
				_ = _out19
				_out16, _out17, _out18, _out19 = Companion_Default___.M0(Companion_Default___.SafeModuloInt((_204_v98).Cardinality(), _166_v76.F7), (func() _dafny.Int {
					if _45_v0 {
						return _67_v15
					}
					return _166_v76.F7
				})(), _64_globalState)
				_205_v99 = _out16
				_206_v100 = _out17
				_207_v101 = _out18
				_208_v102 = _out19
				_45_v0 = (func() bool {
					if _dafny.Companion_Sequence_.IsProperPrefixOf(_205_v99, _65_v13) {
						return Companion_Default___.Fm24(_64_globalState)
					}
					return _45_v0
				})()
				var _209_v103 *C0
				_ = _209_v103
				var _nw38 *C0 = New_C0_()
				_ = _nw38
				_nw38.Ctor__(_207_v101, _67_v15)
				_209_v103 = _nw38
				var _210_v104 *C1
				_ = _210_v104
				var _nw39 *C1 = New_C1_()
				_ = _nw39
				_nw39.Ctor__(_67_v15, Companion_Default___.Fm6(_68_v16, _208_v102, _45_v0, _207_v101, _64_globalState), (_209_v103).F12(), _45_v0)
				_210_v104 = _nw39
				var _211_v105 D4
				_ = _211_v105
				_211_v105 = Companion_D4_.Create_DC14_()
				var _212_v106 _dafny.Map
				_ = _212_v106
				_212_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v104, _211_v105)
				(_166_v76).F7 = Companion_Default___.Fm6(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(148), _166_v76.F7), (_212_v106).Cardinality(), (_68_v16).Cmp(_166_v76.F7) != 0, (_68_v16).Times(_46_v1), _64_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _213_v107 _dafny.Sequence
	_ = _213_v107
	var _214_v108 _dafny.Array
	_ = _214_v108
	var _215_v109 _dafny.Int
	_ = _215_v109
	var _216_v110 _dafny.Int
	_ = _216_v110
	var _out20 _dafny.Sequence
	_ = _out20
	var _out21 _dafny.Array
	_ = _out21
	var _out22 _dafny.Int
	_ = _out22
	var _out23 _dafny.Int
	_ = _out23
	_out20, _out21, _out22, _out23 = Companion_Default___.M0((func() _dafny.Int {
		if ((_72_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))).Int()).(_dafny.MultiSet)).Contains(_68_v16) {
			return ((_72_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_72_v18), 0))).Int()).(_dafny.MultiSet)).Multiplicity(_68_v16)
		}
		return _68_v16
	})(), _67_v15, _64_globalState)
	_213_v107 = _out20
	_214_v108 = _out21
	_215_v109 = _out22
	_216_v110 = _out23
	var _217_v111 _dafny.MultiSet
	_ = _217_v111
	_217_v111 = _dafny.MultiSetOf(_45_v0, !(_45_v0))
	_45_v0 = (!(_217_v111).Contains(_45_v0)) == (_45_v0)
	_dafny.Print(_45_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_46_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_47_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_48_v3).Equals(_dafny.MultiSetOf(_dafny.CodePoint('d'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_49_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(82))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_52_v7.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_53_v8, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_54_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v11).Equals(_dafny.SetOf(_dafny.IntOfInt64(82))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_61_v12, _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_64_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_64_globalState).F1()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_64_globalState).F2(), _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(82)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_64_globalState).F3(), _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_64_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_65_v13.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v14).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_67_v15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v17).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(1352), _dafny.IntOfInt64(1352))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_72_v18).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(1352))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_73_v19, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(1352), _dafny.IntOfInt64(1352)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_74_v20, _dafny.SeqOf(_dafny.IntOfInt64(720))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_78_v21).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.Zero).UpdateUnsafe(_dafny.IntOfInt64(1270), _dafny.IntOfInt64(-741)).UpdateUnsafe(_dafny.IntOfInt64(4), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v22).Dtor_cf1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v22).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_79_v22).Dtor_cf3()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(82))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_79_v22).Dtor_cf4(), _dafny.SeqOf(_dafny.IntOfInt64(720))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_80_v23).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_84_v25).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_166_v76.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_167_v77).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_168_v78).Dtor_cf42()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v79).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_170_v81).Dtor_cf26()).Dtor_cf22()).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(1352), _dafny.IntOfInt64(1352), _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_172_v82).Dtor_cf22()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(1352))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_203_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_213_v107.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v108).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_215_v109)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_216_v110)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_217_v111).Equals(_dafny.MultiSetOf(false, true)))
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
	Cf1 _dafny.Sequence
	Cf2 _dafny.Int
	Cf3 _dafny.Map
	Cf4 _dafny.Sequence
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Sequence, Cf2 _dafny.Int, Cf3 _dafny.Map, Cf4 _dafny.Sequence) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 _dafny.Map
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 _dafny.Map, Cf6 bool) D0 {
	return D0{D0_DC2{Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.EmptySeq, _dafny.Zero, _dafny.EmptyMap, _dafny.EmptySeq)
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + data.Cf1.VerbatimString(true) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf1.Equals(data2.Cf1) && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Equals(data2.Cf3) && data1.Cf4.Equals(data2.Cf4)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6
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

type D1_DC4 struct {
	Cf8  bool
	Cf9  bool
	Cf10 _dafny.Sequence
	Cf11 _dafny.Int
	Cf12 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 bool, Cf9 bool, Cf10 _dafny.Sequence, Cf11 _dafny.Int, Cf12 _dafny.Int) D1 {
	return D1{D1_DC4{Cf8, Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf13 bool
	Cf14 bool
	Cf15 _dafny.Int
	Cf16 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf13 bool, Cf14 bool, Cf15 _dafny.Int, Cf16 _dafny.Int) D1 {
	return D1{D1_DC5{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf17 bool
	Cf18 _dafny.Int
	Cf19 _dafny.Int
	Cf20 bool
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf17 bool, Cf18 _dafny.Int, Cf19 _dafny.Int, Cf20 bool) D1 {
	return D1{D1_DC6{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC3 struct {
	Cf7 _dafny.Map
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 _dafny.Map) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC7 struct {
	Cf21 D1
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_(Cf21 D1) D1 {
	return D1{D1_DC7{Cf21}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, false, _dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf12
}

func (_this D1) Dtor_cf13() bool {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf14() bool {
	return _this.Get_().(D1_DC5).Cf14
}

func (_this D1) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf15
}

func (_this D1) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf16
}

func (_this D1) Dtor_cf17() bool {
	return _this.Get_().(D1_DC6).Cf17
}

func (_this D1) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf18
}

func (_this D1) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf19
}

func (_this D1) Dtor_cf20() bool {
	return _this.Get_().(D1_DC6).Cf20
}

func (_this D1) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf21() D1 {
	return _this.Get_().(D1_DC7).Cf21
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + data.Cf10.VerbatimString(true) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9 && data1.Cf10.Equals(data2.Cf10) && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf7.Equals(data2.Cf7)
		}
	case D1_DC7:
		{
			data2, ok := other.Get_().(D1_DC7)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

type D2_DC9 struct {
	Cf23 bool
	Cf24 _dafny.Int
	Cf25 _dafny.Int
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf23 bool, Cf24 _dafny.Int, Cf25 _dafny.Int) D2 {
	return D2{D2_DC9{Cf23, Cf24, Cf25}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC8 struct {
	Cf22 _dafny.MultiSet
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf22 _dafny.MultiSet) D2 {
	return D2{D2_DC8{Cf22}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC10 struct {
	Cf26 D2
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_(Cf26 D2) D2 {
	return D2{D2_DC10{Cf26}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC9_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf23() bool {
	return _this.Get_().(D2_DC9).Cf23
}

func (_this D2) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf24
}

func (_this D2) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf25
}

func (_this D2) Dtor_cf22() _dafny.MultiSet {
	return _this.Get_().(D2_DC8).Cf22
}

func (_this D2) Dtor_cf26() D2 {
	return _this.Get_().(D2_DC10).Cf26
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D2_DC10:
		{
			return "D2.DC10" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D2_DC10:
		{
			data2, ok := other.Get_().(D2_DC10)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D3_DC12 struct {
	Cf28 bool
	Cf29 _dafny.Int
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf28 bool, Cf29 _dafny.Int) D3 {
	return D3{D3_DC12{Cf28, Cf29}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC11 struct {
	Cf27 _dafny.Array
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf27 _dafny.Array) D3 {
	return D3{D3_DC11{Cf27}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC12_(false, _dafny.Zero)
}

func (_this D3) Dtor_cf28() bool {
	return _this.Get_().(D3_DC12).Cf28
}

func (_this D3) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D3_DC12).Cf29
}

func (_this D3) Dtor_cf27() _dafny.Array {
	return _this.Get_().(D3_DC11).Cf27
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf28 == data2.Cf28 && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf27 == data2.Cf27
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

type D4_DC14 struct {
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_() D4 {
	return D4{D4_DC14{}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf30 _dafny.Map
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf30 _dafny.Map) D4 {
	return D4{D4_DC13{Cf30}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_()
}

func (_this D4) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D4_DC13).Cf30
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC14:
		{
			_, ok := other.Get_().(D4_DC14)
			return ok
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf32 T0
	Cf33 bool
	Cf34 _dafny.Int
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf32 T0, Cf33 bool, Cf34 _dafny.Int) D5 {
	return D5{D5_DC16{Cf32, Cf33, Cf34}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC15 struct {
	Cf31 _dafny.MultiSet
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf31 _dafny.MultiSet) D5 {
	return D5{D5_DC15{Cf31}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC16_((T0)(nil), false, _dafny.Zero)
}

func (_this D5) Dtor_cf32() T0 {
	return _this.Get_().(D5_DC16).Cf32
}

func (_this D5) Dtor_cf33() bool {
	return _this.Get_().(D5_DC16).Cf33
}

func (_this D5) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf34
}

func (_this D5) Dtor_cf31() _dafny.MultiSet {
	return _this.Get_().(D5_DC15).Cf31
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf32, data2.Cf32) && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf36 bool
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf36 bool) D6 {
	return D6{D6_DC18{Cf36}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC19 struct {
	Cf37 _dafny.Int
	Cf38 *C0
	Cf39 _dafny.Int
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf37 _dafny.Int, Cf38 *C0, Cf39 _dafny.Int) D6 {
	return D6{D6_DC19{Cf37, Cf38, Cf39}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC17 struct {
	Cf35 _dafny.Sequence
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf35 _dafny.Sequence) D6 {
	return D6{D6_DC17{Cf35}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_(false)
}

func (_this D6) Dtor_cf36() bool {
	return _this.Get_().(D6_DC18).Cf36
}

func (_this D6) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf37
}

func (_this D6) Dtor_cf38() *C0 {
	return _this.Get_().(D6_DC19).Cf38
}

func (_this D6) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D6_DC19).Cf39
}

func (_this D6) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf35
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf36 == data2.Cf36
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38 == data2.Cf38 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
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

type D7_DC20 struct {
	Cf40 _dafny.CodePoint
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf40 _dafny.CodePoint) D7 {
	return D7{D7_DC20{Cf40}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D7) Dtor_cf40() _dafny.CodePoint {
	return _this.Get_().(D7_DC20).Cf40
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf40) + ")"
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
			return ok && data1.Cf40 == data2.Cf40
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

type D8_DC21 struct {
	Cf41 _dafny.Array
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf41 _dafny.Array) D8 {
	return D8{D8_DC21{Cf41}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D8) Dtor_cf41() _dafny.Array {
	return _this.Get_().(D8_DC21).Cf41
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf41 == data2.Cf41
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

type D9_DC23 struct {
	Cf43 _dafny.Int
	Cf44 _dafny.Int
	Cf45 bool
	Cf46 _dafny.Map
	Cf47 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf43 _dafny.Int, Cf44 _dafny.Int, Cf45 bool, Cf46 _dafny.Map, Cf47 _dafny.Int) D9 {
	return D9{D9_DC23{Cf43, Cf44, Cf45, Cf46, Cf47}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC22 struct {
	Cf42 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf42 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf42}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC24 struct {
	Cf48 D9
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf48 D9) D9 {
	return D9{D9_DC24{Cf48}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(_dafny.Zero, _dafny.Zero, false, _dafny.EmptyMap, _dafny.Zero)
}

func (_this D9) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf43
}

func (_this D9) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf44
}

func (_this D9) Dtor_cf45() bool {
	return _this.Get_().(D9_DC23).Cf45
}

func (_this D9) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D9_DC23).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf47
}

func (_this D9) Dtor_cf42() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf42
}

func (_this D9) Dtor_cf48() D9 {
	return _this.Get_().(D9_DC24).Cf48
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45 == data2.Cf45 && data1.Cf46.Equals(data2.Cf46) && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
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

type D10_DC26 struct {
	Cf50 bool
	Cf51 _dafny.Int
	Cf52 bool
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf50 bool, Cf51 _dafny.Int, Cf52 bool) D10 {
	return D10{D10_DC26{Cf50, Cf51, Cf52}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC25 struct {
	Cf49 _dafny.Set
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf49 _dafny.Set) D10 {
	return D10{D10_DC25{Cf49}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC26_(false, _dafny.Zero, false)
}

func (_this D10) Dtor_cf50() bool {
	return _this.Get_().(D10_DC26).Cf50
}

func (_this D10) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D10_DC26).Cf51
}

func (_this D10) Dtor_cf52() bool {
	return _this.Get_().(D10_DC26).Cf52
}

func (_this D10) Dtor_cf49() _dafny.Set {
	return _this.Get_().(D10_DC25).Cf49
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf49) + ")"
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
			return ok && data1.Cf50 == data2.Cf50 && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52
		}
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf49.Equals(data2.Cf49)
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
	Fm0(p0 bool, globalState *GlobalState) bool
	M1(p0 _dafny.Int, globalState *GlobalState)
	F5() _dafny.Int
	F6() bool
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
	_f0 _dafny.Int
	_f1 _dafny.Map
	_f2 _dafny.Sequence
	_f3 _dafny.Sequence
	_f4 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.EmptyMap
	_this._f2 = _dafny.EmptySeq
	_this._f3 = _dafny.EmptySeq
	_this._f4 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Map, f2 _dafny.Sequence, f3 _dafny.Sequence, f4 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Map {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Sequence {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 _dafny.Int
	_f13 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.Zero
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

func (_this *C0) Ctor__(f12 _dafny.Int, f13 _dafny.Int) {
	{
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C0) Fm7(p0 _dafny.Sequence, p1 _dafny.Map, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	{
		return (func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(925), _dafny.IntOfInt64(394))); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _218_v0 _dafny.Int
				_218_v0 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(925)).Cmp(_218_v0) <= 0) && ((_218_v0).Cmp(_dafny.IntOfInt64(394)) < 0) {
					_coll6.Add((_218_v0).Times((_this).F12()), true)
				}
			}
			return _coll6.ToMap()
		}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), true)).Merge((Companion_D1_.Create_DC3_(func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(397), _dafny.IntOfInt64(116))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _219_v1 _dafny.Int
				_219_v1 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(397)).Cmp(_219_v1) <= 0) && ((_219_v1).Cmp(_dafny.IntOfInt64(116)) < 0) {
					_coll7.Add((_219_v1).Plus((_this).F13()), false)
				}
			}
			return _coll7.ToMap()
		}())).Dtor_cf7()))
	}
}
func (_this *C0) Fm8(p0 bool, p1 _dafny.Sequence, p2 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return (!(true)) || (false)
	}
}
func (_this *C0) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *C0) F13() _dafny.Int {
	{
		return _this._f13
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f5  _dafny.Int
	_f6  bool
	F10  _dafny.Int
	_f11 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this.F10 = _dafny.Zero
	_this._f11 = _dafny.Zero
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
func (_this *C1) F6() bool {
	return _this._f6
}
func (_this *C1) Ctor__(f10 _dafny.Int, f11 _dafny.Int, f5 _dafny.Int, f6 bool) {
	{
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C1) Fm0(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C1) Fm2(p0 _dafny.Set, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_this).F11()).Cmp((_dafny.Zero).Minus(((_this).F5()).Plus((_this).F11()))) == 0
	}
}
func (_this *C1) Fm3(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F10)
	}
}
func (_this *C1) M1(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _220_v0 _dafny.Sequence
		_ = _220_v0
		_220_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ju")
		var _221_v1 _dafny.MultiSet
		_ = _221_v1
		_221_v1 = _dafny.MultiSetOf((func() _dafny.Sequence {
			if false {
				return _220_v0
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("kcobgiw")
		})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_222_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})), _dafny.Companion_Sequence_.Concatenate(_220_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.Zero)).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_223_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))))
		(_this).F10 = (func() _dafny.Int {
			if (_221_v1).Contains(_220_v0) {
				return (_221_v1).Multiplicity(_220_v0)
			}
			return p0
		})()
		var _224_v2 _dafny.Array
		_ = _224_v2
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw40
		_224_v2 = _nw40
		var _225_v3 _dafny.Map
		_ = _225_v3
		_225_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p0)
		var _226_v4 _dafny.Map
		_ = _226_v4
		_226_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_225_v3).Cardinality(), _this.F10)
		var _227_v5 _dafny.Set
		_ = _227_v5
		_227_v5 = _dafny.SetOf(true, (_this).F6())
		var _228_v6 _dafny.Sequence
		_ = _228_v6
		_228_v6 = _dafny.SeqOf((_227_v5).Cardinality())
		var _229_v7 _dafny.Map
		_ = _229_v7
		_229_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _228_v6)
		var _230_v8 D0
		_ = _230_v8
		_230_v8 = Companion_D0_.Create_DC2_(_229_v7, !((_this).F6()))
		var _231_v9 _dafny.Set
		_ = _231_v9
		_231_v9 = _dafny.SetOf(_230_v8)
		var _hi1 _dafny.Int = (func() _dafny.Int {
			if (_226_v4).Contains(_dafny.IntOfUint32((_220_v0).Cardinality())) {
				return (_226_v4).Get(_dafny.IntOfUint32((_220_v0).Cardinality())).(_dafny.Int)
			}
			return (_231_v9).Cardinality()
		})()
		_ = _hi1
		for _232_i2 := Companion_Default___.SafeDivisionInt((_this).F5(), (Companion_Default___.Fm4(_dafny.IntOfUint32((_dafny.SeqOf(_224_v2)).Cardinality()), globalState)).Cardinality()); _232_i2.Cmp(_hi1) < 0; _232_i2 = _232_i2.Plus(_dafny.One) {
			var _233_v10 _dafny.Sequence
			_ = _233_v10
			_233_v10 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _228_v6))
			var _234_v11 _dafny.Map
			_ = _234_v11
			_234_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), false)
			_225_v3 = (_225_v3).Update((Companion_D0_.Create_DC2_((_233_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_233_v10).Cardinality()))).Uint32()).(_dafny.Map), (_this).F6())).Dtor_cf6(), (_234_v11).Cardinality())
			var _235_v12 _dafny.Map
			_ = _235_v12
			_235_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_232_i2, _220_v0)
			var _236_v13 _dafny.CodePoint
			_ = _236_v13
			_236_v13 = _dafny.CodePoint('b')
			var _237_v14 D0
			_ = _237_v14
			_237_v14 = Companion_D0_.Create_DC1_(_dafny.Companion_Sequence_.Update(_220_v0, (Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_220_v0).Cardinality()))).Uint32(), _236_v13), (_this).F5(), _225_v3, _228_v6)
			var _238_v15 _dafny.Sequence
			_ = _238_v15
			_238_v15 = _dafny.SeqOf(_220_v0, _220_v0)
			var _239_v16 _dafny.Map
			_ = _239_v16
			_239_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(652), _225_v3)
			var _240_v17 _dafny.Array
			_ = _240_v17
			var _nwElement0_4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_220_v0, _dafny.UnicodeSeqOfUtf8Bytes("md"))
			_ = _nwElement0_4
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(24))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_4, 0)
			(_nw41).ArraySet1((func() _dafny.Sequence {
				if (_235_v12).Contains(p0) {
					return (_235_v12).Get(p0).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(70))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}(func(_241_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))
			})(), 1)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_220_v0, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_220_v0).Cardinality()))).Uint32(), _236_v13), _220_v0), 2)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Update((_237_v14).Dtor_cf1(), (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32(((_237_v14).Dtor_cf1()).Cardinality()))).Uint32(), _236_v13), 3)
			(_nw41).ArraySet1(_220_v0, 4)
			(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("o"), 5)
			(_nw41).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("th"), 6)
			(_nw41).ArraySet1((_238_v15).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_238_v15).Cardinality()))).Uint32()).(_dafny.Sequence), 7)
			(_nw41).ArraySet1(_220_v0, 8)
			(_nw41).ArraySet1(_220_v0, 9)
			(_nw41).ArraySet1(_220_v0, 10)
			(_nw41).ArraySet1(_220_v0, 11)
			(_nw41).ArraySet1(Companion_Default___.Fm5(_225_v3, _239_v16, globalState), 12)
			(_nw41).ArraySet1((func() _dafny.Sequence {
				if (_235_v12).Contains((_this).F11()) {
					return (_235_v12).Get((_this).F11()).(_dafny.Sequence)
				}
				return _220_v0
			})(), 13)
			(_nw41).ArraySet1(_220_v0, 14)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("chm"), _220_v0), 15)
			(_nw41).ArraySet1(_220_v0, 16)
			(_nw41).ArraySet1(_220_v0, 17)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Update(_220_v0, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_220_v0).Cardinality()))).Uint32(), _236_v13), 18)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _220_v0), 19)
			(_nw41).ArraySet1((Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("ylxgswe"), Companion_Default___.Fm6((_this).F5(), (_this).F5(), (_this).F6(), (_this).F5(), globalState), _225_v3, _228_v6)).Dtor_cf1(), 20)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_220_v0, _220_v0), 21)
			(_nw41).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_220_v0, _dafny.UnicodeSeqOfUtf8Bytes("soi")), 22)
			(_nw41).ArraySet1(_220_v0, 23)
			_240_v17 = _nw41
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_240_v17), 0))
			_ = _index31
			(_240_v17).ArraySet1(_220_v0, (_index31).Int())
			var _242_v18 bool
			_ = _242_v18
			_242_v18 = false
			var _243_v19 _dafny.Sequence
			_ = _243_v19
			_243_v19 = _dafny.SeqOf((_this).Fm0((_this).Fm0(true, globalState), globalState), _242_v18)
			var _244_v20 _dafny.Sequence
			_ = _244_v20
			_244_v20 = _dafny.SeqOf(_243_v19)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_240_v17), 0))
			_ = _index32
			var _rhs17 _dafny.Int = (_232_i2).Times((_this.F10).Times(p0))
			_ = _rhs17
			var _rhs18 _dafny.Int = p0
			_ = _rhs18
			var _rhs19 _dafny.Sequence = _220_v0
			_ = _rhs19
			var _rhs20 bool = (_this.F10).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_244_v20, _244_v20)).Cardinality())) < 0
			_ = _rhs20
			var _lhs12 *C1 = _this
			_ = _lhs12
			var _lhs13 *C1 = _this
			_ = _lhs13
			var _lhs14 _dafny.Array = _240_v17
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_240_v17), 0))
			_ = _lhs15
			_lhs12.F10 = _rhs17
			_lhs13.F10 = _rhs18
			(_lhs14).ArraySet1(_rhs19, (_lhs15).Int())
			_242_v18 = _rhs20
			(_this).F10 = _232_i2
			(_this).F10 = (_dafny.Zero).Minus(((_this).F11()).Minus(p0))
		}
		var _245_v22 _dafny.Map
		_ = _245_v22
		_245_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm5(_225_v3, func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(111), _dafny.IntOfInt64(211))); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _246_v21 _dafny.Int
				_246_v21 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.IntOfInt64(111)).Cmp(_246_v21) <= 0) && ((_246_v21).Cmp(_dafny.IntOfInt64(211)) < 0) {
					_coll8.Add((_246_v21).Times((_this).F11()), _225_v3)
				}
			}
			return _coll8.ToMap()
		}(), globalState), (_this).F11())
		(_this).F10 = (_245_v22).Cardinality()
		var _hi2 _dafny.Int = _this.F10
		_ = _hi2
		for _247_i4 := (_this).F5(); _247_i4.Cmp(_hi2) < 0; _247_i4 = _247_i4.Plus(_dafny.One) {
			var _248_v23 *C0
			_ = _248_v23
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__(Companion_Default___.Fm6((_dafny.Zero).Minus(_dafny.IntOfUint32((_220_v0).Cardinality())), _247_i4, (_this).F6(), (_this).F5(), globalState), (_this).F11())
			_248_v23 = _nw42
			if (func() bool {
				if (_this).F6() {
					return (_this).F6()
				}
				return ((_this).F6()) == ((_this).F6())
			})() {
				var _249_v24 _dafny.Sequence
				_ = _249_v24
				var _250_v25 _dafny.Array
				_ = _250_v25
				var _251_v26 _dafny.Int
				_ = _251_v26
				var _252_v27 _dafny.Int
				_ = _252_v27
				var _out24 _dafny.Sequence
				_ = _out24
				var _out25 _dafny.Array
				_ = _out25
				var _out26 _dafny.Int
				_ = _out26
				var _out27 _dafny.Int
				_ = _out27
				_out24, _out25, _out26, _out27 = Companion_Default___.M0(_this.F10, (_228_v6).Select((Companion_Default___.SafeIndex((_248_v23).F13(), _dafny.IntOfUint32((_228_v6).Cardinality()))).Uint32()).(_dafny.Int), globalState)
				_249_v24 = _out24
				_250_v25 = _out25
				_251_v26 = _out26
				_252_v27 = _out27
				var _253_v28 _dafny.MultiSet
				_ = _253_v28
				_253_v28 = _dafny.MultiSetOf((_this).F6())
				var _254_v29 _dafny.Map
				_ = _254_v29
				_254_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_253_v28, (_dafny.Zero).Minus((_248_v23).F12()))
				var _255_v30 _dafny.MultiSet
				_ = _255_v30
				_255_v30 = _dafny.MultiSetOf((_248_v23).F12(), p0, _dafny.IntOfUint32((_249_v24).Cardinality()), _this.F10, _dafny.IntOfInt64(612))
				var _256_v31 _dafny.Map
				_ = _256_v31
				_256_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _253_v28)
				var _257_v32 D2
				_ = _257_v32
				_257_v32 = Companion_D2_.Create_DC8_(_255_v30)
				var _258_v33 _dafny.Map
				_ = _258_v33
				_258_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())))
				var _259_v34 _dafny.Array
				_ = _259_v34
				var _nwElement0_5 _dafny.MultiSet = _255_v30
				_ = _nwElement0_5
				var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(15))
				_ = _nw43
				(_nw43).ArraySet1(_nwElement0_5, 0)
				(_nw43).ArraySet1((_255_v30).Difference(_255_v30), 1)
				(_nw43).ArraySet1(_dafny.MultiSetFromSeq(_228_v6), 2)
				(_nw43).ArraySet1(_255_v30, 3)
				(_nw43).ArraySet1(_255_v30, 4)
				(_nw43).ArraySet1(_255_v30, 5)
				(_nw43).ArraySet1(_255_v30, 6)
				(_nw43).ArraySet1((_255_v30).Union(_255_v30), 7)
				(_nw43).ArraySet1(_255_v30, 8)
				(_nw43).ArraySet1((_255_v30).Difference(_255_v30), 9)
				(_nw43).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(180))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_260_v23 *C0) func(_dafny.Int) _dafny.Int {
					return func(_261_i5 _dafny.Int) _dafny.Int {
						return (_260_v23).F13()
					}
				})(_248_v23))), (Companion_Default___.Fm9((_248_v23).F13(), globalState)).Dtor_cf4())), 10)
				(_nw43).ArraySet1(_dafny.MultiSetFromSeq(_228_v6), 11)
				(_nw43).ArraySet1(_255_v30, 12)
				(_nw43).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_228_v6, _dafny.SeqOf(((func() _dafny.MultiSet {
					if (_256_v31).Contains(_247_i4) {
						return (_256_v31).Get(_247_i4).(_dafny.MultiSet)
					}
					return _253_v28
				})()).Cardinality(), (_248_v23).F13(), _dafny.IntOfUint32((_220_v0).Cardinality())))), 13)
				(_nw43).ArraySet1(((_257_v32).Dtor_cf22()).Update((func() _dafny.Int {
					if (_258_v33).Contains(_224_v2) {
						return (_258_v33).Get(_224_v2).(_dafny.Int)
					}
					return (_225_v3).Cardinality()
				})(), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(457))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}(func(_262_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))).Cardinality()))), 14)
				_259_v34 = _nw43
				var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_259_v34), 0))
				_ = _index33
				(_259_v34).ArraySet1(_255_v30, (_index33).Int())
				var _263_v35 bool
				_ = _263_v35
				_263_v35 = true
				var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_224_v2), 0))
				_ = _index34
				(_224_v2).ArraySet1(_263_v35, (_index34).Int())
				var _264_v36 _dafny.Sequence
				_ = _264_v36
				_264_v36 = _dafny.SeqOf(true, (_this).F6(), false)
				var _265_v37 D0
				_ = _265_v37
				_265_v37 = Companion_D0_.Create_DC1_(_249_v24, _252_v27, _225_v3, _dafny.SeqOf((_dafny.Zero).Minus((_248_v23).F13())))
				var _266_v38 _dafny.CodePoint
				_ = _266_v38
				_266_v38 = _dafny.CodePoint('d')
				var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_259_v34), 0))
				_ = _index35
				var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_224_v2), 0))
				_ = _index36
				var _rhs21 _dafny.Map = _254_v29
				_ = _rhs21
				var _rhs22 _dafny.MultiSet = _255_v30
				_ = _rhs22
				var _rhs23 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_264_v36, _dafny.SeqOf((_this).F6(), !((_this).F6()))), _dafny.SeqOf(_263_v35))
				_ = _rhs23
				var _rhs24 bool = (_248_v23).Fm8(_263_v35, (_265_v37).Dtor_cf1(), _266_v38, globalState)
				_ = _rhs24
				var _lhs16 _dafny.Array = _259_v34
				_ = _lhs16
				var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_259_v34), 0))
				_ = _lhs17
				var _lhs18 _dafny.Array = _224_v2
				_ = _lhs18
				var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_224_v2), 0))
				_ = _lhs19
				_254_v29 = _rhs21
				(_lhs16).ArraySet1(_rhs22, (_lhs17).Int())
				_263_v35 = _rhs23
				(_lhs18).ArraySet1(_rhs24, (_lhs19).Int())
				_249_v24 = _249_v24
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_259_v34), 0))
				_ = _index37
				(_259_v34).ArraySet1(_255_v30, (_index37).Int())
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_224_v2), 0))
				_ = _index38
				var _rhs25 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F5(), (_this).F11())
				_ = _rhs25
				var _rhs26 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_249_v24, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg20 _dafny.Int) interface{} {
						return coer20(arg20)
					}
				}((func(_267_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_268_i7 _dafny.Int) _dafny.CodePoint {
						return _267_v38
					}
				})(_266_v38))), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(125))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}(func(_269_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (Companion_Default___.SafeIndex((_248_v23).F12(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(125))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_270_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				}))).Cardinality()))).Uint32(), _266_v38)))
				_ = _rhs26
				var _rhs27 _dafny.Array = _250_v25
				_ = _rhs27
				var _rhs28 _dafny.Int = _252_v27
				_ = _rhs28
				var _rhs29 bool = _263_v35
				_ = _rhs29
				var _lhs20 _dafny.Array = _224_v2
				_ = _lhs20
				var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_224_v2), 0))
				_ = _lhs21
				var _lhs22 *C1 = _this
				_ = _lhs22
				_252_v27 = _rhs25
				(_lhs20).ArraySet1(_rhs26, (_lhs21).Int())
				_250_v25 = _rhs27
				_lhs22.F10 = _rhs28
				_263_v35 = _rhs29
			} else {
				var _271_v39 _dafny.CodePoint
				_ = _271_v39
				_271_v39 = _dafny.CodePoint('w')
				_271_v39 = _271_v39
				(_this).F10 = (func() _dafny.Int {
					if !((_this).F6()) {
						return _dafny.IntOfUint32((_220_v0).Cardinality())
					}
					return (_248_v23).F12()
				})()
				var _272_v40 _dafny.Array
				_ = _272_v40
				var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
				_ = _nw44
				_272_v40 = _nw44
				var _273_v41 _dafny.MultiSet
				_ = _273_v41
				_273_v41 = _dafny.MultiSetOf((_this).F6())
				var _274_v42 _dafny.Sequence
				_ = _274_v42
				_274_v42 = _dafny.SeqOf((_this).F6(), (_this).F6())
				var _275_v43 _dafny.MultiSet
				_ = _275_v43
				_275_v43 = _dafny.MultiSetOf(_dafny.IntOfUint32((_220_v0).Cardinality()), p0, _this.F10, _dafny.IntOfUint32((_274_v42).Cardinality()))
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_272_v40), 0))
				_ = _index39
				(_272_v40).ArraySet1((_273_v41).Difference(Companion_Default___.Fm10(_275_v43, (_248_v23).F13(), globalState)), (_index39).Int())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_272_v40), 0))
				_ = _index40
				(_272_v40).ArraySet1(_273_v41, (_index40).Int())
				_271_v39 = _271_v39
				var _276_v44 bool
				_ = _276_v44
				_276_v44 = true
				_276_v44 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_220_v0, _dafny.UnicodeSeqOfUtf8Bytes("qchhltun"))).Cardinality())).Cmp((_dafny.Zero).Minus((_247_i4).Plus(_this.F10))) < 0
			}
			if (_this).F6() {
				var _277_v45 *C0
				_ = _277_v45
				var _nw45 *C0 = New_C0_()
				_ = _nw45
				_nw45.Ctor__(_dafny.IntOfInt64(401), (func() _dafny.Int {
					if (_this).F6() {
						return _247_i4
					}
					return _dafny.IntOfInt64(-2)
				})())
				_277_v45 = _nw45
				(_this).F10 = (_dafny.Zero).Minus((_this).F11())
				var _278_v46 _dafny.Array
				_ = _278_v46
				var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw46
				_278_v46 = _nw46
				var _279_v47 _dafny.Set
				_ = _279_v47
				_279_v47 = _dafny.SetOf(_278_v46, _278_v46, _278_v46)
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_278_v46), 0))
				_ = _index41
				(_278_v46).ArraySet1((_this).F5(), (_index41).Int())
				var _280_v48 _dafny.Map
				_ = _280_v48
				_280_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v45).F12(), _220_v0)
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_278_v46), 0))
				_ = _index42
				var _rhs30 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(478), (_277_v45).F12())).Plus(((_280_v48).Merge(_280_v48)).Cardinality())
				_ = _rhs30
				var _rhs31 _dafny.Set = _279_v47
				_ = _rhs31
				var _rhs32 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ihvp")
				_ = _rhs32
				var _rhs33 _dafny.Int = Companion_Default___.SafeModuloInt((_277_v45).F13(), (_248_v23).F12())
				_ = _rhs33
				var _lhs23 *C1 = _this
				_ = _lhs23
				var _lhs24 _dafny.Array = _278_v46
				_ = _lhs24
				var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_278_v46), 0))
				_ = _lhs25
				_lhs23.F10 = _rhs30
				_279_v47 = _rhs31
				_220_v0 = _rhs32
				(_lhs24).ArraySet1(_rhs33, (_lhs25).Int())
				_224_v2 = _224_v2
				var _281_v49 _dafny.Array
				_ = _281_v49
				var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw47
				_281_v49 = _nw47
				var _282_v50 _dafny.Sequence
				_ = _282_v50
				_282_v50 = _dafny.SeqOf((_this).F6())
				var _rhs34 _dafny.Array = (func() _dafny.Array {
					if ((_278_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_278_v46), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(115)) >= 0 {
						return _281_v49
					}
					return _281_v49
				})()
				_ = _rhs34
				var _rhs35 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_282_v50, _282_v50), _282_v50)
				_ = _rhs35
				_281_v49 = _rhs34
				_282_v50 = _rhs35
			} else {
				var _283_v51 _dafny.Array
				_ = _283_v51
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_4
				var _nw48 _dafny.Array
				_ = _nw48
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw48 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Int = (func(_284_v23 *C0) func(_dafny.Int) _dafny.Int {
						return func(_285_i9 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_285_i9, (_284_v23).F12())
						}
					})(_248_v23)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw48 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw48).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw48).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_283_v51 = _nw48
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_283_v51), 0))
				_ = _index43
				(_283_v51).ArraySet1(p0, (_index43).Int())
				var _286_v52 _dafny.Sequence
				_ = _286_v52
				_286_v52 = _dafny.SeqOf(!((_this).F6()))
				var _287_v53 _dafny.CodePoint
				_ = _287_v53
				_287_v53 = _dafny.CodePoint('e')
				var _288_v54 _dafny.Map
				_ = _288_v54
				_288_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_226_v4).Cardinality(), (_this).F6())
				var _289_v55 _dafny.Map
				_ = _289_v55
				_289_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_248_v23).F13(), _225_v3)
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_283_v51), 0))
				_ = _index44
				(_283_v51).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_220_v0, _220_v0), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_286_v52).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_220_v0, _220_v0)).Cardinality()))).Uint32(), Companion_Default___.Fm11(_287_v53, true, (_this).F6(), globalState)), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5((_225_v3).Update((_this).F6(), (_288_v54).Cardinality()), _289_v55, globalState), _220_v0))).Cardinality()), (_index44).Int())
				(_this).F10 = (_248_v23).F12()
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_224_v2), 0))
				_ = _index45
				(_224_v2).ArraySet1(!((_this).F6()) || ((_this).F6()), (_index45).Int())
				var _290_v56 bool
				_ = _290_v56
				_290_v56 = true
				var _291_v57 _dafny.Array
				_ = _291_v57
				var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
				_ = _nw49
				_291_v57 = _nw49
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_291_v57), 0))
				_ = _index46
				(_291_v57).ArraySet1(_283_v51, (_index46).Int())
				var _292_v58 _dafny.MultiSet
				_ = _292_v58
				_292_v58 = _dafny.MultiSetOf(_290_v56, (_this).F6())
				var _293_v59 _dafny.MultiSet
				_ = _293_v59
				_293_v59 = _dafny.MultiSetOf((_248_v23).F12())
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_224_v2), 0))
				_ = _index47
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_291_v57), 0))
				_ = _index48
				var _rhs36 bool = ((_this).F6()) == ((Companion_Default___.Fm10(_293_v59, _dafny.IntOfInt64(100), globalState)).IsProperSubsetOf(_292_v58))
				_ = _rhs36
				var _rhs37 bool = _290_v56
				_ = _rhs37
				var _rhs38 _dafny.Array = _283_v51
				_ = _rhs38
				var _lhs26 _dafny.Array = _224_v2
				_ = _lhs26
				var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_224_v2), 0))
				_ = _lhs27
				var _lhs28 _dafny.Array = _291_v57
				_ = _lhs28
				var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_291_v57), 0))
				_ = _lhs29
				(_lhs26).ArraySet1(_rhs36, (_lhs27).Int())
				_290_v56 = _rhs37
				(_lhs28).ArraySet1(_rhs38, (_lhs29).Int())
				var _294_v60 _dafny.Map
				_ = _294_v60
				_294_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v56, false)
				_287_v53 = (func() _dafny.CodePoint {
					if !(((_294_v60).Cardinality()).Cmp(_dafny.IntOfInt64(294)) <= 0) {
						return _287_v53
					}
					return _287_v53
				})()
				_290_v56 = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_286_v52, _dafny.SeqOf((_286_v52).Select((Companion_Default___.SafeIndex((_248_v23).F13(), _dafny.IntOfUint32((_286_v52).Cardinality()))).Uint32()).(bool))))
			}
			var _295_v61 *C0
			_ = _295_v61
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(950), (_this).F5()), p0)
			_295_v61 = _nw50
		}
		var _296_v62 _dafny.Map
		_ = _296_v62
		_296_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		(_this).F10 = ((_296_v62).Cardinality()).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_228_v6, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_228_v6).Cardinality()))).Uint32(), _this.F10), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(101))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_297_i10 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(711)
		})))).Cardinality()))
		var _298_v63 bool
		_ = _298_v63
		_298_v63 = false
		_298_v63 = !((_this).F6())
	}
}
func (_this *C1) M4(p0 _dafny.Map, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) {
	{
		var _299_v0 _dafny.Sequence
		_ = _299_v0
		_299_v0 = _dafny.SeqOf(true, (_this).F6(), (_this).Fm0((_this).F6(), globalState))
		var _300_v1 _dafny.Map
		_ = _300_v1
		_300_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), _299_v0)
		_300_v1 = (_300_v1).Update(!(false), _299_v0)
		var _hi3 _dafny.Int = (_this).F5()
		_ = _hi3
		for _301_i0 := p3; _301_i0.Cmp(_hi3) < 0; _301_i0 = _301_i0.Plus(_dafny.One) {
			var _302_v2 _dafny.MultiSet
			_ = _302_v2
			_302_v2 = _dafny.MultiSetOf(p2)
			var _303_v3 _dafny.Sequence
			_ = _303_v3
			_303_v3 = _dafny.SeqOf((_302_v2).Cardinality(), Companion_Default___.Fm6(p3, (_this).F5(), (_this).F6(), _dafny.IntOfInt64(-771), globalState))
			var _304_v4 _dafny.MultiSet
			_ = _304_v4
			_304_v4 = _dafny.MultiSetOf(_this.F10, (_303_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_303_v3).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(929))
			var _305_v5 _dafny.Sequence
			_ = _305_v5
			_305_v5 = _dafny.SeqOf(_303_v3)
			var _306_v6 *C0
			_ = _306_v6
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__((_this).F5(), (func() _dafny.Int {
				if (_304_v4).Contains(_dafny.IntOfUint32((_305_v5).Cardinality())) {
					return (_304_v4).Multiplicity(_dafny.IntOfUint32((_305_v5).Cardinality()))
				}
				return _this.F10
			})())
			_306_v6 = _nw51
			var _307_v7 D0
			_ = _307_v7
			_307_v7 = Companion_D0_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _303_v3), (_this).F6())
			var _308_v8 _dafny.Map
			_ = _308_v8
			_308_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqOf(p3))
			var _309_v9 _dafny.MultiSet
			_ = _309_v9
			_309_v9 = _dafny.MultiSetOf(_307_v7, (func() D0 {
				if (_this).F6() {
					return _307_v7
				}
				return func(_pat_let7_0 D0) D0 {
					return func(_310_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let8_0 bool) D0 {
							return func(_311_dt__update_hcf6_h0 bool) D0 {
								return Companion_D0_.Create_DC2_((_310_dt__update__tmp_h0).Dtor_cf5(), _311_dt__update_hcf6_h0)
							}(_pat_let8_0)
						}((_this).F6())
					}(_pat_let7_0)
				}(Companion_D0_.Create_DC2_(_308_v8, false))
			})())
			_309_v9 = _309_v9
			var _312_v10 bool
			_ = _312_v10
			_312_v10 = false
			_312_v10 = _312_v10
			var _313_v11 D0
			_ = _313_v11
			_313_v11 = Companion_D0_.Create_DC0_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(455))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_314_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_315_i1 _dafny.Int) _dafny.Int {
					return _314_p3
				}
			})(p3))))
			var _source6 D0 = _313_v11
			_ = _source6
			if _source6.Is_DC1() {
				var _316___mcc_h0 _dafny.Sequence = _source6.Get_().(D0_DC1).Cf1
				_ = _316___mcc_h0
				var _317___mcc_h1 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
				_ = _317___mcc_h1
				var _318___mcc_h2 _dafny.Map = _source6.Get_().(D0_DC1).Cf3
				_ = _318___mcc_h2
				var _319___mcc_h3 _dafny.Sequence = _source6.Get_().(D0_DC1).Cf4
				_ = _319___mcc_h3
				var _320_cf4 _dafny.Sequence = _319___mcc_h3
				_ = _320_cf4
				var _321_cf3 _dafny.Map = _318___mcc_h2
				_ = _321_cf3
				var _322_cf2 _dafny.Int = _317___mcc_h1
				_ = _322_cf2
				var _323_cf1 _dafny.Sequence = _316___mcc_h0
				_ = _323_cf1
				var _324_v12 _dafny.Array
				_ = _324_v12
				var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw52
				_324_v12 = _nw52
				_324_v12 = _324_v12
				var _rhs39 bool = (_this).F6()
				_ = _rhs39
				var _rhs40 _dafny.Sequence = _320_cf4
				_ = _rhs40
				var _rhs41 bool = _312_v10
				_ = _rhs41
				_312_v10 = _rhs39
				_303_v3 = _rhs40
				_312_v10 = _rhs41
				var _325_v13 _dafny.MultiSet
				_ = _325_v13
				_325_v13 = _dafny.MultiSetOf(_312_v10)
				var _326_v14 *C0
				_ = _326_v14
				var _nw53 *C0 = New_C0_()
				_ = _nw53
				_nw53.Ctor__(Companion_Default___.Fm6(_this.F10, (func() _dafny.Int {
					if (_325_v13).Contains(false) {
						return (_325_v13).Multiplicity(false)
					}
					return p2
				})(), _312_v10, (_306_v6).F13(), globalState), (_306_v6).F12())
				_326_v14 = _nw53
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_324_v12), 0))
				_ = _index49
				(_324_v12).ArraySet1((_this).F5(), (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_324_v12), 0))
				_ = _index50
				(_324_v12).ArraySet1(((_this).F11()).Times((_306_v6).F13()), (_index50).Int())
			} else if _source6.Is_DC2() {
				var _327___mcc_h4 _dafny.Map = _source6.Get_().(D0_DC2).Cf5
				_ = _327___mcc_h4
				var _328___mcc_h5 bool = _source6.Get_().(D0_DC2).Cf6
				_ = _328___mcc_h5
				var _329_cf6 bool = _328___mcc_h5
				_ = _329_cf6
				var _330_cf5 _dafny.Map = _327___mcc_h4
				_ = _330_cf5
				var _331_v15 _dafny.Set
				_ = _331_v15
				_331_v15 = _dafny.SetOf(_329_cf6)
				var _332_v16 _dafny.Set
				_ = _332_v16
				_332_v16 = _dafny.SetOf(_this.F10, p3, (_306_v6).F12(), (_331_v15).Cardinality(), _301_i0)
				var _333_v17 _dafny.Array
				_ = _333_v17
				var _nwElement0_6 _dafny.Int = (_332_v16).Cardinality()
				_ = _nwElement0_6
				var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(11))
				_ = _nw54
				(_nw54).ArraySet1(_nwElement0_6, 0)
				(_nw54).ArraySet1(p3, 1)
				(_nw54).ArraySet1(_dafny.IntOfUint32((_303_v3).Cardinality()), 2)
				(_nw54).ArraySet1((_this).F11(), 3)
				(_nw54).ArraySet1((_306_v6).F12(), 4)
				(_nw54).ArraySet1((_306_v6).F13(), 5)
				(_nw54).ArraySet1((_this.F10).Times((_306_v6).F12()), 6)
				(_nw54).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wecoqytg")).Cardinality()), 7)
				(_nw54).ArraySet1((_this).F11(), 8)
				(_nw54).ArraySet1((_306_v6).F12(), 9)
				(_nw54).ArraySet1(_this.F10, 10)
				_333_v17 = _nw54
				_333_v17 = _333_v17
				var _334_v18 _dafny.Array
				_ = _334_v18
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_5
				var _nw55 _dafny.Array
				_ = _nw55
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw55 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) bool = (func(_335_v10 bool) func(_dafny.Int) bool {
						return func(_336_i2 _dafny.Int) bool {
							return _335_v10
						}
					})(_312_v10)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw55 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw55).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw55).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_334_v18 = _nw55
				var _337_v19 _dafny.Array
				_ = _337_v19
				var _nwElement0_7 _dafny.Array = _334_v18
				_ = _nwElement0_7
				var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(29))
				_ = _nw56
				(_nw56).ArraySet1(_nwElement0_7, 0)
				(_nw56).ArraySet1(_334_v18, 1)
				(_nw56).ArraySet1(_334_v18, 2)
				(_nw56).ArraySet1(_334_v18, 3)
				(_nw56).ArraySet1(_334_v18, 4)
				(_nw56).ArraySet1(_334_v18, 5)
				(_nw56).ArraySet1(_334_v18, 6)
				(_nw56).ArraySet1(_334_v18, 7)
				(_nw56).ArraySet1(_334_v18, 8)
				(_nw56).ArraySet1(_334_v18, 9)
				(_nw56).ArraySet1(_334_v18, 10)
				(_nw56).ArraySet1(_334_v18, 11)
				(_nw56).ArraySet1(_334_v18, 12)
				(_nw56).ArraySet1(_334_v18, 13)
				(_nw56).ArraySet1(_334_v18, 14)
				(_nw56).ArraySet1(_334_v18, 15)
				(_nw56).ArraySet1(_334_v18, 16)
				(_nw56).ArraySet1(_334_v18, 17)
				(_nw56).ArraySet1(_334_v18, 18)
				(_nw56).ArraySet1(_334_v18, 19)
				(_nw56).ArraySet1(_334_v18, 20)
				(_nw56).ArraySet1(_334_v18, 21)
				(_nw56).ArraySet1(_334_v18, 22)
				(_nw56).ArraySet1(_334_v18, 23)
				(_nw56).ArraySet1(_334_v18, 24)
				(_nw56).ArraySet1(_334_v18, 25)
				(_nw56).ArraySet1(_334_v18, 26)
				(_nw56).ArraySet1(_334_v18, 27)
				(_nw56).ArraySet1(_334_v18, 28)
				_337_v19 = _nw56
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_337_v19), 0))
				_ = _index51
				(_337_v19).ArraySet1(_334_v18, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_337_v19), 0))
				_ = _index52
				(_337_v19).ArraySet1(_334_v18, (_index52).Int())
				var _338_v20 _dafny.Sequence
				_ = _338_v20
				_338_v20 = _dafny.SeqOf(_dafny.ArrayCastTo((_337_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_337_v19), 0))).Int())))
				var _339_v21 _dafny.Map
				_ = _339_v21
				_339_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_306_v6).F12())).Cardinality()), _dafny.Companion_Sequence_.Update(_338_v20, (Companion_Default___.SafeIndex((_this).F11(), _dafny.IntOfUint32((_338_v20).Cardinality()))).Uint32(), _334_v18))
				var _340_v22 _dafny.Sequence
				_ = _340_v22
				var _341_v23 _dafny.Array
				_ = _341_v23
				var _342_v24 _dafny.Int
				_ = _342_v24
				var _343_v25 _dafny.Int
				_ = _343_v25
				var _out28 _dafny.Sequence
				_ = _out28
				var _out29 _dafny.Array
				_ = _out29
				var _out30 _dafny.Int
				_ = _out30
				var _out31 _dafny.Int
				_ = _out31
				_out28, _out29, _out30, _out31 = Companion_Default___.M0((_this.F10).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_339_v21).Contains((_306_v6).F13()) {
						return (_339_v21).Get((_306_v6).F13()).(_dafny.Sequence)
					}
					return _338_v20
				})()).Cardinality())), _dafny.IntOfInt64(425), globalState)
				_340_v22 = _out28
				_341_v23 = _out29
				_342_v24 = _out30
				_343_v25 = _out31
				_312_v10 = (_this).Fm2(_dafny.SetOf((_this).F6(), _312_v10, _329_cf6, _329_cf6, _312_v10), (_dafny.Zero).Minus(_342_v24), _dafny.IntOfUint32((_340_v22).Cardinality()), globalState)
			} else {
				var _344___mcc_h6 _dafny.Sequence = _source6.Get_().(D0_DC0).Cf0
				_ = _344___mcc_h6
				var _345_cf0 _dafny.Sequence = _344___mcc_h6
				_ = _345_cf0
				var _346_v26 _dafny.Map
				_ = _346_v26
				_346_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v6, _312_v10)
				var _347_v27 _dafny.Set
				_ = _347_v27
				_347_v27 = _dafny.SetOf(_346_v26, _346_v26)
				var _348_v28 _dafny.CodePoint
				_ = _348_v28
				_348_v28 = _dafny.CodePoint('n')
				var _349_v29 _dafny.Sequence
				_ = _349_v29
				_349_v29 = _dafny.UnicodeSeqOfUtf8Bytes("xcfy")
				var _350_v30 D1
				_ = _350_v30
				_350_v30 = Companion_D1_.Create_DC6_((_this).F6(), (_347_v27).Cardinality(), (_306_v6).F12(), !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_349_v29, (Companion_Default___.SafeIndex((_306_v6).F13(), _dafny.IntOfUint32((_349_v29).Cardinality()))).Uint32(), _348_v28), _348_v28))
				_350_v30 = _350_v30
				_312_v10 = _312_v10
				var _351_v31 _dafny.MultiSet
				_ = _351_v31
				_351_v31 = _dafny.MultiSetOf(false)
				var _352_v32 _dafny.Map
				_ = _352_v32
				_352_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()) || (_312_v10), (_351_v31).IsProperSubsetOf(_351_v31))
				_352_v32 = (_352_v32).Update(_312_v10, _312_v10)
				(_this).F10 = p2
			}
		}
		var _353_v33 _dafny.Sequence
		_ = _353_v33
		_353_v33 = _dafny.SeqOf(p2, _this.F10, _this.F10)
		var _354_v34 _dafny.Set
		_ = _354_v34
		_354_v34 = _dafny.SetOf((_this).F11())
		var _355_v35 _dafny.Map
		_ = _355_v35
		_355_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		var _356_v36 _dafny.Array
		_ = _356_v36
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_6
		var _nw57 _dafny.Array
		_ = _nw57
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw57 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = (func(_357_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_358_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_358_i3, _357_p3)
				}
			})(p3)
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
		_356_v36 = _nw57
		var _359_v37 _dafny.Map
		_ = _359_v37
		_359_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, _356_v36)
		var _360_v38 D3
		_ = _360_v38
		_360_v38 = Companion_D3_.Create_DC11_(_356_v36)
		_353_v33 = Companion_Default___.Fm12(((_354_v34).Cardinality()).Times(_this.F10), !(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_355_v35).Cardinality(), (func() _dafny.Array {
			if (_359_v37).Contains(_this.F10) {
				return (_359_v37).Get(_this.F10).(_dafny.Array)
			}
			return (_360_v38).Dtor_cf27()
		})())).Cardinality()).Cmp(_dafny.IntOfInt64(431)) == 0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-214))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_361_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), globalState)
		var _362_v39 _dafny.Map
		_ = _362_v39
		_362_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vbenawu")).Cardinality()))
		var _363_v40 _dafny.MultiSet
		_ = _363_v40
		_363_v40 = _dafny.MultiSetOf(p2, (_this).F5())
		var _364_v41 _dafny.Map
		_ = _364_v41
		_364_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_362_v39).Cardinality(), (_363_v40).Intersection(_363_v40))
		var _365_v42 _dafny.Map
		_ = _365_v42
		_365_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _353_v33)
		var _366_v43 D0
		_ = _366_v43
		_366_v43 = Companion_D0_.Create_DC2_(_365_v42, (_this).F6())
		var _367_v44 _dafny.Set
		_ = _367_v44
		_367_v44 = _dafny.SetOf(!(true))
		_364_v41 = (_364_v41).Update((func() _dafny.Int {
			if (_366_v43).Dtor_cf6() {
				return (_this).F11()
			}
			return Companion_Default___.Fm6((_367_v44).Cardinality(), p3, false, _dafny.IntOfInt64(393), globalState)
		})(), (_363_v40).Update((_this).F5(), Companion_Default___.Abs(_dafny.IntOfInt64(515))))
		var _368_v45 _dafny.CodePoint
		_ = _368_v45
		_368_v45 = _dafny.CodePoint('i')
		var _369_v46 _dafny.Map
		_ = _369_v46
		_369_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v45, p3)
		var _370_v47 D4
		_ = _370_v47
		_370_v47 = Companion_D4_.Create_DC13_(_369_v46)
		var _pat_let_tv7 = _369_v46
		_ = _pat_let_tv7
		var _371_v48 *C0
		_ = _371_v48
		var _nw58 *C0 = New_C0_()
		_ = _nw58
		_nw58.Ctor__((_this).F11(), (func() _dafny.Int {
			if !(false) {
				return ((func(_pat_let9_0 D4) D4 {
					return func(_372_dt__update__tmp_h1 D4) D4 {
						return func(_pat_let10_0 _dafny.Map) D4 {
							return func(_373_dt__update_hcf30_h0 _dafny.Map) D4 {
								return Companion_D4_.Create_DC13_(_373_dt__update_hcf30_h0)
							}(_pat_let10_0)
						}(_pat_let_tv7)
					}(_pat_let9_0)
				}(_370_v47)).Dtor_cf30()).Cardinality()
			}
			return p2
		})())
		_371_v48 = _nw58
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_356_v36), 0))); ; {
			_guard_loop_0, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _374_i5 _dafny.Int
			_374_i5 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_374_i5).Sign() != -1) && ((_374_i5).Cmp(_dafny.ArrayLen((_356_v36), 0)) < 0)) {
				(_356_v36).ArraySet1((_374_i5).Plus(Companion_Default___.SafeModuloInt((_this).F11(), (_371_v48).F12())), (_374_i5).Int())
			}
		}
	}
}
func (_this *C1) F11() _dafny.Int {
	{
		return _this._f11
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F9 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F9 = _dafny.Zero
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f9 _dafny.Int) {
	{
		(_this).F9 = f9
	}
}
func (_this *C2) Fm1(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), true))).Cardinality())).Cmp(_this.F9) < 0
	}
}
func (_this *C2) M3(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		(_this).F9 = _this.F9
		var _375_v0 bool
		_ = _375_v0
		_375_v0 = false
		var _376_i0 _dafny.Int
		_ = _376_i0
		_376_i0 = _dafny.Zero
		{
			for (_375_v0) || (_375_v0) {
				{
					if (_376_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_376_i0 = (_376_i0).Plus(_dafny.One)
					var _377_v1 T0
					_ = _377_v1
					var _nw59 *C1 = New_C1_()
					_ = _nw59
					_nw59.Ctor__(_this.F9, _dafny.IntOfUint32((_dafny.SeqOf(_375_v0, _375_v0, _375_v0, _375_v0)).Cardinality()), _this.F9, _375_v0)
					_377_v1 = _nw59
					var _378_v2 _dafny.Set
					_ = _378_v2
					_378_v2 = _dafny.SetOf((func() T0 {
						if _375_v0 {
							return _377_v1
						}
						return _377_v1
					})())
					var _379_v3 _dafny.Map
					_ = _379_v3
					_379_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(_377_v1))
					var _380_v4 _dafny.Sequence
					_ = _380_v4
					_380_v4 = _dafny.SeqOf(_378_v2)
					_378_v2 = ((func() _dafny.Set {
						if (_379_v3).Contains(p0) {
							return (_379_v3).Get(p0).(_dafny.Set)
						}
						return (_380_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.IntOfUint32((_380_v4).Cardinality()))).Uint32()).(_dafny.Set)
					})()).Difference((_378_v2).Difference(_378_v2))
					var _381_v5 _dafny.Map
					_ = _381_v5
					_381_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_377_v1).F5(), _dafny.IntOfInt64(613)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}(func(_382_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('d')
					})))
					_381_v5 = (_381_v5).Update((_377_v1).F5(), (func() _dafny.Sequence {
						if (_377_v1).F6() {
							return _dafny.UnicodeSeqOfUtf8Bytes("byj")
						}
						return p0
					})())
					var _383_v6 _dafny.Array
					_ = _383_v6
					var _nw60 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
					_ = _nw60
					_383_v6 = _nw60
					var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))
					_ = _index53
					(_383_v6).ArraySet1(_375_v0, (_index53).Int())
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))
					_ = _index54
					(_383_v6).ArraySet1(!(!(((_377_v1).F5()).Cmp(Companion_Default___.Fm6((_377_v1).F5(), _this.F9, _375_v0, (_377_v1).F5(), globalState)) == 0)), (_index54).Int())
					var _384_v7 _dafny.Map
					_ = _384_v7
					_384_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_375_v0, _this.F9)
					var _385_v8 _dafny.Map
					_ = _385_v8
					_385_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_383_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))).Int()).(bool), _384_v7)
					var _386_v9 _dafny.Sequence
					_ = _386_v9
					_386_v9 = _dafny.SeqOf((_377_v1).F5())
					var _387_v10 D0
					_ = _387_v10
					_387_v10 = Companion_D0_.Create_DC1_(p0, _this.F9, _384_v7, _386_v9)
					var _388_v11 _dafny.Map
					_ = _388_v11
					_388_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9, _384_v7)
					var _389_v12 _dafny.MultiSet
					_ = _389_v12
					_389_v12 = _dafny.MultiSetOf((_377_v1).F5())
					var _390_v13 D1
					_ = _390_v13
					_390_v13 = Companion_D1_.Create_DC6_(_375_v0, (_389_v12).Cardinality(), (_386_v9).Select((Companion_Default___.SafeIndex((_377_v1).F5(), _dafny.IntOfUint32((_386_v9).Cardinality()))).Uint32()).(_dafny.Int), (_377_v1).F6())
					var _391_v14 _dafny.Sequence
					_ = _391_v14
					_391_v14 = _dafny.SeqOf((_383_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))).Int()).(bool), (_383_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))).Int()).(bool), (_383_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))).Int()).(bool), (_377_v1).F6(), (_390_v13).Dtor_cf20())
					var _392_v15 _dafny.MultiSet
					_ = _392_v15
					_392_v15 = _dafny.MultiSetOf(_dafny.IntOfUint32((_391_v14).Cardinality()), (_377_v1).F5(), _this.F9, (_377_v1).F5(), _this.F9)
					var _393_v16 D2
					_ = _393_v16
					_393_v16 = Companion_D2_.Create_DC8_(_392_v15)
					var _394_v17 _dafny.Array
					_ = _394_v17
					var _nwElement0_8 _dafny.Map = (func() _dafny.Map {
						if (_385_v8).Contains((_377_v1).F6()) {
							return (_385_v8).Get((_377_v1).F6()).(_dafny.Map)
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_383_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_383_v6), 0))).Int()).(bool), (_377_v1).F5())
					})()
					_ = _nwElement0_8
					var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(19))
					_ = _nw61
					(_nw61).ArraySet1(_nwElement0_8, 0)
					(_nw61).ArraySet1(_384_v7, 1)
					(_nw61).ArraySet1(_384_v7, 2)
					(_nw61).ArraySet1(((_387_v10).Dtor_cf3()).Merge(_384_v7), 3)
					(_nw61).ArraySet1(_384_v7, 4)
					(_nw61).ArraySet1(_384_v7, 5)
					(_nw61).ArraySet1(_384_v7, 6)
					(_nw61).ArraySet1(_384_v7, 7)
					(_nw61).ArraySet1(_384_v7, 8)
					(_nw61).ArraySet1(_384_v7, 9)
					(_nw61).ArraySet1(_384_v7, 10)
					(_nw61).ArraySet1(_384_v7, 11)
					(_nw61).ArraySet1((_384_v7).Merge((func() _dafny.Map {
						if (_388_v11).Contains((_377_v1).F5()) {
							return (_388_v11).Get((_377_v1).F5()).(_dafny.Map)
						}
						return _384_v7
					})()), 12)
					(_nw61).ArraySet1((_384_v7).Merge(_384_v7), 13)
					(_nw61).ArraySet1(_384_v7, 14)
					(_nw61).ArraySet1((_384_v7).Update((_377_v1).F6(), (_377_v1).F5()), 15)
					(_nw61).ArraySet1(_384_v7, 16)
					(_nw61).ArraySet1(Companion_Default___.Fm13(p0, (func() _dafny.Int {
						if (_392_v15).Contains((_dafny.Zero).Minus((_377_v1).F5())) {
							return (_392_v15).Multiplicity((_dafny.Zero).Minus((_377_v1).F5()))
						}
						return _dafny.IntOfInt64(144)
					})(), (_393_v16).Dtor_cf22(), globalState), 17)
					(_nw61).ArraySet1(_384_v7, 18)
					_394_v17 = _nw61
					_394_v17 = _394_v17
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _395_v18 _dafny.Map
		_ = _395_v18
		_395_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9, _this.F9)
		var _396_v19 _dafny.Sequence
		_ = _396_v19
		_396_v19 = _dafny.SeqOf(_dafny.IntOfInt64(-537), Companion_Default___.Fm6((_395_v18).Cardinality(), _dafny.IntOfInt64(210), _375_v0, _this.F9, globalState))
		var _397_i2 _dafny.Int
		_ = _397_i2
		_397_i2 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_396_v19, (func() _dafny.Sequence {
				if _375_v0 {
					return _396_v19
				}
				return _dafny.SeqOf(_this.F9)
			})()) {
				{
					if (_397_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_397_i2 = (_397_i2).Plus(_dafny.One)
					var _398_v20 _dafny.Sequence
					_ = _398_v20
					_398_v20 = _dafny.SeqOf(_375_v0)
					var _source7 D1 = Companion_Default___.Fm14(!_dafny.Companion_Sequence_.Contains(_398_v20, _375_v0), globalState)
					_ = _source7
					if _source7.Is_DC4() {
						var _399___mcc_h0 bool = _source7.Get_().(D1_DC4).Cf8
						_ = _399___mcc_h0
						var _400___mcc_h1 bool = _source7.Get_().(D1_DC4).Cf9
						_ = _400___mcc_h1
						var _401___mcc_h2 _dafny.Sequence = _source7.Get_().(D1_DC4).Cf10
						_ = _401___mcc_h2
						var _402___mcc_h3 _dafny.Int = _source7.Get_().(D1_DC4).Cf11
						_ = _402___mcc_h3
						var _403___mcc_h4 _dafny.Int = _source7.Get_().(D1_DC4).Cf12
						_ = _403___mcc_h4
						var _404_cf12 _dafny.Int = _403___mcc_h4
						_ = _404_cf12
						var _405_cf11 _dafny.Int = _402___mcc_h3
						_ = _405_cf11
						var _406_cf10 _dafny.Sequence = _401___mcc_h2
						_ = _406_cf10
						var _407_cf9 bool = _400___mcc_h1
						_ = _407_cf9
						var _408_cf8 bool = _399___mcc_h0
						_ = _408_cf8
						var _rhs42 bool = true
						_ = _rhs42
						var _rhs43 bool = _407_cf9
						_ = _rhs43
						_407_cf9 = _rhs42
						_407_cf9 = _rhs43
						var _409_v21 _dafny.MultiSet
						_ = _409_v21
						_409_v21 = _dafny.MultiSetOf(_375_v0, _407_cf9, false, !(!(_407_cf9) || (_408_cf8)))
						_409_v21 = _409_v21
						_407_cf9 = _408_cf8
						var _410_v22 _dafny.Map
						_ = _410_v22
						_410_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_cf8, Companion_Default___.Fm12(_405_cf11, _375_v0, _406_cf10, globalState))
						_407_cf9 = !(_410_v22).Equals((Companion_Default___.Fm15(_405_cf11, globalState)).Merge(_410_v22))
					} else if _source7.Is_DC5() {
						var _411___mcc_h5 bool = _source7.Get_().(D1_DC5).Cf13
						_ = _411___mcc_h5
						var _412___mcc_h6 bool = _source7.Get_().(D1_DC5).Cf14
						_ = _412___mcc_h6
						var _413___mcc_h7 _dafny.Int = _source7.Get_().(D1_DC5).Cf15
						_ = _413___mcc_h7
						var _414___mcc_h8 _dafny.Int = _source7.Get_().(D1_DC5).Cf16
						_ = _414___mcc_h8
						var _415_cf16 _dafny.Int = _414___mcc_h8
						_ = _415_cf16
						var _416_cf15 _dafny.Int = _413___mcc_h7
						_ = _416_cf15
						var _417_cf14 bool = _412___mcc_h6
						_ = _417_cf14
						var _418_cf13 bool = _411___mcc_h5
						_ = _418_cf13
						_418_cf13 = false
						_398_v20 = Companion_Default___.Fm16(globalState)
						var _419_v23 _dafny.MultiSet
						_ = _419_v23
						_419_v23 = _dafny.MultiSetOf(_dafny.IntOfInt64(575))
						var _420_v24 _dafny.MultiSet
						_ = _420_v24
						_420_v24 = _dafny.MultiSetOf(_375_v0, false, _418_cf13, _375_v0)
						var _421_v25 D2
						_ = _421_v25
						_421_v25 = Companion_D2_.Create_DC8_((_419_v23).Union(_dafny.MultiSetOf(_this.F9, (func() _dafny.Int {
							if (_420_v24).Contains(_375_v0) {
								return (_420_v24).Multiplicity(_375_v0)
							}
							return _415_cf16
						})(), _this.F9)))
						var _pat_let_tv8 = _419_v23
						_ = _pat_let_tv8
						var _pat_let_tv9 = _415_cf16
						_ = _pat_let_tv9
						var _pat_let_tv10 = _416_cf15
						_ = _pat_let_tv10
						var _pat_let_tv11 = _415_cf16
						_ = _pat_let_tv11
						_421_v25 = func(_pat_let11_0 D2) D2 {
							return func(_424_dt__update__tmp_h1 D2) D2 {
								return func(_pat_let14_0 _dafny.MultiSet) D2 {
									return func(_425_dt__update_hcf22_h1 _dafny.MultiSet) D2 {
										return Companion_D2_.Create_DC8_(_425_dt__update_hcf22_h1)
									}(_pat_let14_0)
								}(_dafny.MultiSetOf(_pat_let_tv9, _pat_let_tv10, _pat_let_tv11))
							}(_pat_let11_0)
						}(func(_pat_let12_0 D2) D2 {
							return func(_422_dt__update__tmp_h0 D2) D2 {
								return func(_pat_let13_0 _dafny.MultiSet) D2 {
									return func(_423_dt__update_hcf22_h0 _dafny.MultiSet) D2 {
										return Companion_D2_.Create_DC8_(_423_dt__update_hcf22_h0)
									}(_pat_let13_0)
								}(_pat_let_tv8)
							}(_pat_let12_0)
						}(_421_v25))
						var _426_v26 _dafny.CodePoint
						_ = _426_v26
						_426_v26 = _dafny.CodePoint('c')
						_426_v26 = _426_v26
					} else if _source7.Is_DC6() {
						var _427___mcc_h9 bool = _source7.Get_().(D1_DC6).Cf17
						_ = _427___mcc_h9
						var _428___mcc_h10 _dafny.Int = _source7.Get_().(D1_DC6).Cf18
						_ = _428___mcc_h10
						var _429___mcc_h11 _dafny.Int = _source7.Get_().(D1_DC6).Cf19
						_ = _429___mcc_h11
						var _430___mcc_h12 bool = _source7.Get_().(D1_DC6).Cf20
						_ = _430___mcc_h12
						var _431_cf20 bool = _430___mcc_h12
						_ = _431_cf20
						var _432_cf19 _dafny.Int = _429___mcc_h11
						_ = _432_cf19
						var _433_cf18 _dafny.Int = _428___mcc_h10
						_ = _433_cf18
						var _434_cf17 bool = _427___mcc_h9
						_ = _434_cf17
						_375_v0 = _431_cf20
						var _435_v27 _dafny.Sequence
						_ = _435_v27
						var _436_v28 _dafny.Array
						_ = _436_v28
						var _437_v29 _dafny.Int
						_ = _437_v29
						var _438_v30 _dafny.Int
						_ = _438_v30
						var _out32 _dafny.Sequence
						_ = _out32
						var _out33 _dafny.Array
						_ = _out33
						var _out34 _dafny.Int
						_ = _out34
						var _out35 _dafny.Int
						_ = _out35
						_out32, _out33, _out34, _out35 = Companion_Default___.M0(_433_cf18, Companion_Default___.Fm6(_this.F9, _this.F9, _431_cf20, _this.F9, globalState), globalState)
						_435_v27 = _out32
						_436_v28 = _out33
						_437_v29 = _out34
						_438_v30 = _out35
						var _439_v31 _dafny.Map
						_ = _439_v31
						_439_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_cf17, _375_v0)
						var _440_v32 _dafny.Map
						_ = _440_v32
						_440_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(647), _439_v31)
						var _441_v33 *C1
						_ = _441_v33
						var _nw62 *C1 = New_C1_()
						_ = _nw62
						_nw62.Ctor__((_438_v30).Times(((func() _dafny.Map {
							if (_440_v32).Contains(_dafny.IntOfInt64(984)) {
								return (_440_v32).Get(_dafny.IntOfInt64(984)).(_dafny.Map)
							}
							return _439_v31
						})()).Cardinality()), _dafny.IntOfInt64(365), _437_v29, _431_cf20)
						_441_v33 = _nw62
						_441_v33 = _441_v33
						var _442_v34 D2
						_ = _442_v34
						_442_v34 = Companion_D2_.Create_DC9_(_375_v0, (_441_v33).F11(), _432_cf19)
						var _443_v35 _dafny.Set
						_ = _443_v35
						_443_v35 = _dafny.SetOf(_375_v0)
						var _444_v36 D1
						_ = _444_v36
						_444_v36 = Companion_D1_.Create_DC6_(_434_cf17, (_443_v35).Cardinality(), _dafny.IntOfUint32((_435_v27).Cardinality()), _375_v0)
						var _445_v37 _dafny.Map
						_ = _445_v37
						_445_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_444_v36).Dtor_cf17(), _396_v19)
						var _446_v38 D0
						_ = _446_v38
						_446_v38 = Companion_D0_.Create_DC2_(_445_v37, _431_cf20)
						var _447_v39 _dafny.MultiSet
						_ = _447_v39
						_447_v39 = _dafny.MultiSetOf(_431_cf20)
						var _448_v40 _dafny.CodePoint
						_ = _448_v40
						_448_v40 = _dafny.CodePoint('u')
						var _449_v41 _dafny.Array
						_ = _449_v41
						var _nwElement0_9 bool = (!(_375_v0)) == (_431_cf20)
						_ = _nwElement0_9
						var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(28))
						_ = _nw63
						(_nw63).ArraySet1(_nwElement0_9, 0)
						(_nw63).ArraySet1(_375_v0, 1)
						(_nw63).ArraySet1((Companion_D2_.Create_DC9_((_442_v34).Dtor_cf23(), _dafny.IntOfInt64(809), _dafny.IntOfUint32((p0).Cardinality()))).Dtor_cf23(), 2)
						(_nw63).ArraySet1(!(_431_cf20), 3)
						(_nw63).ArraySet1(_375_v0, 4)
						(_nw63).ArraySet1(_375_v0, 5)
						(_nw63).ArraySet1(!(_375_v0) || (_434_cf17), 6)
						(_nw63).ArraySet1(_434_cf17, 7)
						(_nw63).ArraySet1(!(!(!(_375_v0) || (_431_cf20))), 8)
						(_nw63).ArraySet1(_431_cf20, 9)
						(_nw63).ArraySet1(_434_cf17, 10)
						(_nw63).ArraySet1(!(_439_v31).Contains(_375_v0), 11)
						(_nw63).ArraySet1(_434_cf17, 12)
						(_nw63).ArraySet1((_442_v34).Dtor_cf23(), 13)
						(_nw63).ArraySet1(_431_cf20, 14)
						(_nw63).ArraySet1(_375_v0, 15)
						(_nw63).ArraySet1(_434_cf17, 16)
						(_nw63).ArraySet1(_375_v0, 17)
						(_nw63).ArraySet1(!((_446_v38).Dtor_cf6()) || (_375_v0), 18)
						(_nw63).ArraySet1(_431_cf20, 19)
						(_nw63).ArraySet1((_dafny.MultiSetOf(!(_431_cf20))).IsDisjointFrom(_447_v39), 20)
						(_nw63).ArraySet1(_431_cf20, 21)
						(_nw63).ArraySet1(_dafny.Companion_Sequence_.Equal(_396_v19, _396_v19), 22)
						(_nw63).ArraySet1((_441_v33).Fm2(_443_v35, (_441_v33).F11(), _433_cf18, globalState), 23)
						(_nw63).ArraySet1(_431_cf20, 24)
						(_nw63).ArraySet1((_this).Fm1(_441_v33.F10, _448_v40, globalState), 25)
						(_nw63).ArraySet1((_398_v20).Select((Companion_Default___.SafeIndex(_438_v30, _dafny.IntOfUint32((_398_v20).Cardinality()))).Uint32()).(bool), 26)
						(_nw63).ArraySet1(_375_v0, 27)
						_449_v41 = _nw63
						var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_449_v41), 0))
						_ = _index55
						(_449_v41).ArraySet1(_434_cf17, (_index55).Int())
						var _450_v42 _dafny.Map
						_ = _450_v42
						_450_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_375_v0, _dafny.IntOfUint32((_435_v27).Cardinality()))
						var _451_v43 _dafny.Sequence
						_ = _451_v43
						_451_v43 = _dafny.SeqOf(_450_v42)
						var _452_v44 D0
						_ = _452_v44
						_452_v44 = Companion_D0_.Create_DC1_(p0, _437_v29, (_451_v43).Select((Companion_Default___.SafeIndex(_441_v33.F10, _dafny.IntOfUint32((_451_v43).Cardinality()))).Uint32()).(_dafny.Map), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(557))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg27 _dafny.Int) interface{} {
								return coer27(arg27)
							}
						}((func(_453_v30 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_454_i3 _dafny.Int) _dafny.Int {
								return _453_v30
							}
						})(_438_v30))))
						var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_449_v41), 0))
						_ = _index56
						(_449_v41).ArraySet1(_dafny.Companion_Sequence_.Contains((_452_v44).Dtor_cf1(), _448_v40), (_index56).Int())
					} else if _source7.Is_DC3() {
						var _455___mcc_h13 _dafny.Map = _source7.Get_().(D1_DC3).Cf7
						_ = _455___mcc_h13
						var _456_cf7 _dafny.Map = _455___mcc_h13
						_ = _456_cf7
						var _457_v45 _dafny.Array
						_ = _457_v45
						var _len0_7 _dafny.Int = _dafny.IntOfInt64(19)
						_ = _len0_7
						var _nw64 _dafny.Array
						_ = _nw64
						if _len0_7.Cmp(_dafny.Zero) == 0 {
							_nw64 = _dafny.NewArray(_len0_7)
						} else {
							var _init7 func(_dafny.Int) _dafny.CodePoint = func(_458_i4 _dafny.Int) _dafny.CodePoint {
								return (func() _dafny.CodePoint {
									if true {
										return _dafny.CodePoint('w')
									}
									return _dafny.CodePoint('l')
								})()
							}
							_ = _init7
							var _element0_7 = _init7(_dafny.Zero)
							_ = _element0_7
							_nw64 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
							(_nw64).ArraySet1CodePoint(_element0_7, 0)
							var _nativeLen0_7 = (_len0_7).Int()
							_ = _nativeLen0_7
							for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
								(_nw64).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
							}
						}
						_457_v45 = _nw64
						var _459_v46 _dafny.CodePoint
						_ = _459_v46
						_459_v46 = _dafny.CodePoint('f')
						var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_457_v45), 0))
						_ = _index57
						(_457_v45).ArraySet1CodePoint(_459_v46, (_index57).Int())
						var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_457_v45), 0))
						_ = _index58
						(_457_v45).ArraySet1CodePoint(_459_v46, (_index58).Int())
						var _460_v47 _dafny.Array
						_ = _460_v47
						var _nw65 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
						_ = _nw65
						_460_v47 = _nw65
						var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_460_v47), 0))
						_ = _index59
						(_460_v47).ArraySet1(!(_375_v0), (_index59).Int())
						var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_460_v47), 0))
						_ = _index60
						(_460_v47).ArraySet1(_375_v0, (_index60).Int())
						var _461_v48 _dafny.MultiSet
						_ = _461_v48
						_461_v48 = _dafny.MultiSetOf((_dafny.Zero).Minus(_this.F9), _this.F9, _this.F9, _this.F9)
						var _462_v49 D2
						_ = _462_v49
						_462_v49 = Companion_D2_.Create_DC8_(_461_v48)
						var _pat_let_tv12 = _461_v48
						_ = _pat_let_tv12
						_462_v49 = func(_pat_let15_0 D2) D2 {
							return func(_463_dt__update__tmp_h2 D2) D2 {
								return func(_pat_let16_0 _dafny.MultiSet) D2 {
									return func(_464_dt__update_hcf22_h2 _dafny.MultiSet) D2 {
										return Companion_D2_.Create_DC8_(_464_dt__update_hcf22_h2)
									}(_pat_let16_0)
								}(_pat_let_tv12)
							}(_pat_let15_0)
						}(_462_v49)
						var _465_v50 _dafny.Array
						_ = _465_v50
						var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
						_ = _nw66
						_465_v50 = _nw66
						var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_465_v50), 0))
						_ = _index61
						(_465_v50).ArraySet1(_460_v47, (_index61).Int())
						var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_465_v50), 0))
						_ = _index62
						var _nw67 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
						_ = _nw67
						(_465_v50).ArraySet1(_nw67, (_index62).Int())
					} else {
						var _466___mcc_h14 D1 = _source7.Get_().(D1_DC7).Cf21
						_ = _466___mcc_h14
						var _467_cf21 D1 = _466___mcc_h14
						_ = _467_cf21
						var _468_v51 _dafny.Map
						_ = _468_v51
						_468_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(500), _375_v0)
						var _469_v52 D1
						_ = _469_v52
						_469_v52 = Companion_D1_.Create_DC3_(_468_v51)
						var _470_v53 D1
						_ = _470_v53
						_470_v53 = Companion_D1_.Create_DC7_(_469_v52)
						_470_v53 = _470_v53
						var _471_v54 _dafny.Sequence
						_ = _471_v54
						var _472_v55 _dafny.Array
						_ = _472_v55
						var _473_v56 _dafny.Int
						_ = _473_v56
						var _474_v57 _dafny.Int
						_ = _474_v57
						var _out36 _dafny.Sequence
						_ = _out36
						var _out37 _dafny.Array
						_ = _out37
						var _out38 _dafny.Int
						_ = _out38
						var _out39 _dafny.Int
						_ = _out39
						_out36, _out37, _out38, _out39 = Companion_Default___.M0(_this.F9, _this.F9, globalState)
						_471_v54 = _out36
						_472_v55 = _out37
						_473_v56 = _out38
						_474_v57 = _out39
						var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_472_v55), 0))
						_ = _index63
						(_472_v55).ArraySet1((func() _dafny.Map {
							var _coll9 = _dafny.NewMapBuilder()
							_ = _coll9
							for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(809), _dafny.IntOfInt64(186))); ; {
								_compr_9, _ok10 := _iter10()
								if !_ok10 {
									break
								}
								var _475_v58 _dafny.Int
								_475_v58 = interface{}(_compr_9).(_dafny.Int)
								if ((_dafny.IntOfInt64(809)).Cmp(_475_v58) <= 0) && ((_475_v58).Cmp(_dafny.IntOfInt64(186)) < 0) {
									_coll9.Add(Companion_Default___.SafeModuloInt(_475_v58, _this.F9), _471_v54)
								}
							}
							return _coll9.ToMap()
						}()).Cardinality(), (_index63).Int())
						var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_472_v55), 0))
						_ = _index64
						(_472_v55).ArraySet1(_474_v57, (_index64).Int())
						var _476_v59 _dafny.Set
						_ = _476_v59
						_476_v59 = _dafny.SetOf(_375_v0, _375_v0)
						var _477_v60 D1
						_ = _477_v60
						_477_v60 = Companion_D1_.Create_DC5_(_375_v0, _375_v0, (_476_v59).Cardinality(), _dafny.IntOfInt64(90))
						var _478_v61 _dafny.Map
						_ = _478_v61
						_478_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_477_v60, _396_v19)
						_478_v61 = (_478_v61).Update(_477_v60, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_471_v54).Cardinality())), _396_v19))
					}
					var _479_v62 _dafny.Array
					_ = _479_v62
					var _nw68 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
					_ = _nw68
					_479_v62 = _nw68
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_479_v62), 0))
					_ = _index65
					(_479_v62).ArraySet1((_398_v20).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((p0).Cardinality())), _dafny.IntOfUint32((_398_v20).Cardinality()))).Uint32()).(bool), (_index65).Int())
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_479_v62), 0))
					_ = _index66
					(_479_v62).ArraySet1(_375_v0, (_index66).Int())
					var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_479_v62), 0))
					_ = _index67
					(_479_v62).ArraySet1(((_479_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_479_v62), 0))).Int()).(bool)) || ((_479_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_479_v62), 0))).Int()).(bool)), (_index67).Int())
					var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_479_v62), 0))
					_ = _index68
					(_479_v62).ArraySet1((_this.F9).Cmp(_this.F9) == 0, (_index68).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		_375_v0 = (_375_v0) == (_375_v0)
		var _480_v63 D4
		_ = _480_v63
		_480_v63 = Companion_D4_.Create_DC14_()
		_480_v63 = _480_v63
		var _481_v64 _dafny.Map
		_ = _481_v64
		_481_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9, _375_v0)
		_481_v64 = (_481_v64).Update(_this.F9, _375_v0)
		var _482_v65 _dafny.CodePoint
		_ = _482_v65
		_482_v65 = _dafny.CodePoint('h')
		var _483_v66 D0
		_ = _483_v66
		_483_v66 = Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.IntOfInt64(-920)))
		var _484_v67 _dafny.Map
		_ = _484_v67
		_484_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_482_v65, _483_v66)
		var _485_v68 _dafny.Set
		_ = _485_v68
		_485_v68 = _dafny.SetOf(_375_v0, _375_v0)
		var _486_v69 _dafny.Map
		_ = _486_v69
		_486_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_v67, (_485_v68).Cardinality())
		var _487_v71 _dafny.Map
		_ = _487_v71
		_487_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_375_v0), _486_v69)
		var _488_v74 _dafny.MultiSet
		_ = _488_v74
		_488_v74 = _dafny.MultiSetOf(_482_v65)
		var _pat_let_tv13 = _396_v19
		_ = _pat_let_tv13
		var _489_v75 _dafny.Set
		_ = _489_v75
		_489_v75 = _dafny.SetOf(func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter11 := _dafny.Iterate((_488_v74).Elements()); ; {
				_compr_10, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _490_v73 _dafny.CodePoint
				_490_v73 = interface{}(_compr_10).(_dafny.CodePoint)
				if (_488_v74).Contains(_490_v73) {
					_coll10.Add(_490_v73, _483_v66)
				}
			}
			return _coll10.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_482_v65, func(_pat_let17_0 D0) D0 {
			return func(_491_dt__update__tmp_h3 D0) D0 {
				return func(_pat_let18_0 _dafny.Sequence) D0 {
					return func(_492_dt__update_hcf0_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_492_dt__update_hcf0_h0)
					}(_pat_let18_0)
				}(_pat_let_tv13)
			}(_pat_let17_0)
		}(_483_v66)))
		var _493_v76 _dafny.Map
		_ = _493_v76
		_493_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F9, _489_v75)
		var _494_v77 _dafny.Sequence
		_ = _494_v77
		_494_v77 = _dafny.SeqOf(_486_v69, func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter12 := _dafny.Iterate(((func() _dafny.Map {
				if (_487_v71).Contains(_375_v0) {
					return (_487_v71).Get(_375_v0).(_dafny.Map)
				}
				return func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter13 := _dafny.Iterate(((func() _dafny.Set {
						if (_493_v76).Contains(_this.F9) {
							return (_493_v76).Get(_this.F9).(_dafny.Set)
						}
						return _489_v75
					})()).Elements()); ; {
						_compr_12, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _495_v72 _dafny.Map
						_495_v72 = interface{}(_compr_12).(_dafny.Map)
						if ((func() _dafny.Set {
							if (_493_v76).Contains(_this.F9) {
								return (_493_v76).Get(_this.F9).(_dafny.Set)
							}
							return _489_v75
						})()).Contains(_495_v72) {
							_coll12.Add(_495_v72, _this.F9)
						}
					}
					return _coll12.ToMap()
				}()
			})()).Keys().Elements()); ; {
				_compr_11, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _496_v70 _dafny.Map
				_496_v70 = interface{}(_compr_11).(_dafny.Map)
				if ((func() _dafny.Map {
					if (_487_v71).Contains(_375_v0) {
						return (_487_v71).Get(_375_v0).(_dafny.Map)
					}
					return func() _dafny.Map {
						var _coll13 = _dafny.NewMapBuilder()
						_ = _coll13
						for _iter14 := _dafny.Iterate(((func() _dafny.Set {
							if (_493_v76).Contains(_this.F9) {
								return (_493_v76).Get(_this.F9).(_dafny.Set)
							}
							return _489_v75
						})()).Elements()); ; {
							_compr_13, _ok14 := _iter14()
							if !_ok14 {
								break
							}
							var _495_v72 _dafny.Map
							_495_v72 = interface{}(_compr_13).(_dafny.Map)
							if ((func() _dafny.Set {
								if (_493_v76).Contains(_this.F9) {
									return (_493_v76).Get(_this.F9).(_dafny.Set)
								}
								return _489_v75
							})()).Contains(_495_v72) {
								_coll13.Add(_495_v72, _this.F9)
							}
						}
						return _coll13.ToMap()
					}()
				})()).Contains(_496_v70) {
					_coll11.Add(_496_v70, _this.F9)
				}
			}
			return _coll11.ToMap()
		}(), _486_v69)
		r0 = (_494_v77).Select((Companion_Default___.SafeIndex(_this.F9, _dafny.IntOfUint32((_494_v77).Cardinality()))).Uint32()).(_dafny.Map)
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f5 _dafny.Int
	_f6 bool
	_f8 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f8 = _dafny.Zero
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

func (_this *C3) F5() _dafny.Int {
	return _this._f5
}
func (_this *C3) F6() bool {
	return _this._f6
}
func (_this *C3) Ctor__(f8 _dafny.Int, f5 _dafny.Int, f6 bool) {
	{
		(_this)._f8 = f8
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C3) Fm0(p0 bool, globalState *GlobalState) bool {
	{
		return !(((_this).F8()).Cmp((_this).F5()) >= 0)
	}
}
func (_this *C3) M1(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _497_i0 _dafny.Int
		_ = _497_i0
		_497_i0 = _dafny.Zero
		{
			for (_this).F6() {
				{
					if (_497_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_497_i0 = (_497_i0).Plus(_dafny.One)
					var _498_v0 bool
					_ = _498_v0
					_498_v0 = false
					_498_v0 = !(_498_v0)
					var _499_v1 _dafny.MultiSet
					_ = _499_v1
					_499_v1 = _dafny.MultiSetOf((_this).F6(), _498_v0)
					var _500_v2 D5
					_ = _500_v2
					_500_v2 = Companion_D5_.Create_DC15_(_499_v1)
					var _501_v3 _dafny.Sequence
					_ = _501_v3
					_501_v3 = _dafny.SeqOf((_dafny.MultiSetOf((_this).F6(), (_this).F6())).Cardinality(), (_this).F8(), ((_500_v2).Dtor_cf31()).Cardinality(), p0)
					var _502_v4 *C2
					_ = _502_v4
					var _nw69 *C2 = New_C2_()
					_ = _nw69
					_nw69.Ctor__(((_501_v3).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_501_v3).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p0))
					_502_v4 = _nw69
					var _503_v5 _dafny.Set
					_ = _503_v5
					_503_v5 = _dafny.SetOf(p0, (_this).F5(), p0, (_this).F5())
					var _504_v6 _dafny.Map
					_ = _504_v6
					_504_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_502_v4.F9, _dafny.IntOfInt64(408))
					var _505_v7 _dafny.Map
					_ = _505_v7
					_505_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_498_v0, _498_v0)
					var _506_v8 _dafny.Array
					_ = _506_v8
					var _nwElement0_10 _dafny.Int = Companion_Default___.Fm6(_dafny.IntOfInt64(56), _502_v4.F9, _498_v0, p0, globalState)
					_ = _nwElement0_10
					var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(15))
					_ = _nw70
					(_nw70).ArraySet1(_nwElement0_10, 0)
					(_nw70).ArraySet1((_this).F8(), 1)
					(_nw70).ArraySet1(_502_v4.F9, 2)
					(_nw70).ArraySet1(_502_v4.F9, 3)
					(_nw70).ArraySet1(_502_v4.F9, 4)
					(_nw70).ArraySet1(Companion_Default___.SafeModuloInt((_this).F8(), (_this).F5()), 5)
					(_nw70).ArraySet1((_dafny.Zero).Minus(_502_v4.F9), 6)
					(_nw70).ArraySet1((_502_v4.F9).Times((_503_v5).Cardinality()), 7)
					(_nw70).ArraySet1(Companion_Default___.Fm6(_502_v4.F9, (_this).F8(), _498_v0, (_this).F5(), globalState), 8)
					(_nw70).ArraySet1((_this).F8(), 9)
					(_nw70).ArraySet1((func() _dafny.Int {
						if (_504_v6).Contains((_this).F5()) {
							return (_504_v6).Get((_this).F5()).(_dafny.Int)
						}
						return (_501_v3).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_501_v3).Cardinality()))).Uint32()).(_dafny.Int)
					})(), 10)
					(_nw70).ArraySet1(_502_v4.F9, 11)
					(_nw70).ArraySet1(Companion_Default___.SafeModuloInt(_502_v4.F9, p0), 12)
					(_nw70).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(889))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}((func(_507_v4 *C2) func(_dafny.Int) _dafny.Int {
						return func(_508_i1 _dafny.Int) _dafny.Int {
							return _507_v4.F9
						}
					})(_502_v4))), _501_v3)).Cardinality()), 13)
					(_nw70).ArraySet1((func() _dafny.Int {
						if _498_v0 {
							return (_dafny.Zero).Minus((_505_v7).Cardinality())
						}
						return p0
					})(), 14)
					_506_v8 = _nw70
					var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_506_v8), 0))
					_ = _index69
					(_506_v8).ArraySet1(_502_v4.F9, (_index69).Int())
					var _509_v9 _dafny.Array
					_ = _509_v9
					var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
					_ = _nw71
					_509_v9 = _nw71
					var _510_v10 _dafny.Array
					_ = _510_v10
					var _nw72 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
					_ = _nw72
					_510_v10 = _nw72
					var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))
					_ = _index70
					(_509_v9).ArraySet1(_510_v10, (_index70).Int())
					var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_506_v8), 0))
					_ = _index71
					var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))
					_ = _index72
					var _rhs44 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F5()))
					_ = _rhs44
					var _rhs45 _dafny.Array = _510_v10
					_ = _rhs45
					var _rhs46 _dafny.Int = (_this).F5()
					_ = _rhs46
					var _lhs30 _dafny.Array = _506_v8
					_ = _lhs30
					var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_506_v8), 0))
					_ = _lhs31
					var _lhs32 _dafny.Array = _509_v9
					_ = _lhs32
					var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))
					_ = _lhs33
					var _lhs34 *C2 = _502_v4
					_ = _lhs34
					(_lhs30).ArraySet1(_rhs44, (_lhs31).Int())
					(_lhs32).ArraySet1(_rhs45, (_lhs33).Int())
					_lhs34.F9 = _rhs46
					var _arr0 _dafny.Array = _dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))).Int()))
					_ = _arr0
					var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))
					_ = _index73
					(_arr0).ArraySet1((_this).F6(), (_index73).Int())
					var _511_v11 T0
					_ = _511_v11
					var _nw73 *C1 = New_C1_()
					_ = _nw73
					_nw73.Ctor__((_506_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_506_v8), 0))).Int()).(_dafny.Int), (_this).F5(), _502_v4.F9, (_this).F6())
					_511_v11 = _nw73
					var _512_v12 _dafny.CodePoint
					_ = _512_v12
					_512_v12 = _dafny.CodePoint('b')
					var _513_v13 _dafny.Sequence
					_ = _513_v13
					_513_v13 = _dafny.SeqOf(_512_v12)
					var _514_v14 D5
					_ = _514_v14
					_514_v14 = Companion_D5_.Create_DC16_(_511_v11, false, _dafny.IntOfUint32((_513_v13).Cardinality()))
					var _pat_let_tv14 = _498_v0
					_ = _pat_let_tv14
					var _arr1 _dafny.Array = _dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))).Int()))
					_ = _arr1
					var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))
					_ = _index74
					var _rhs47 bool = (func(_pat_let19_0 D5) D5 {
						return func(_515_dt__update__tmp_h0 D5) D5 {
							return func(_pat_let20_0 bool) D5 {
								return func(_516_dt__update_hcf33_h0 bool) D5 {
									return Companion_D5_.Create_DC16_((_515_dt__update__tmp_h0).Dtor_cf32(), _516_dt__update_hcf33_h0, (_515_dt__update__tmp_h0).Dtor_cf34())
								}(_pat_let20_0)
							}(_pat_let_tv14)
						}(_pat_let19_0)
					}(_514_v14)).Dtor_cf33()
					_ = _rhs47
					var _rhs48 bool = (_this).F6()
					_ = _rhs48
					var _lhs35 _dafny.Array = _dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))).Int()))
					_ = _lhs35
					var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_dafny.ArrayCastTo((_509_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_509_v9), 0))).Int()))), 0))
					_ = _lhs36
					_498_v0 = _rhs47
					(_lhs35).ArraySet1(_rhs48, (_lhs36).Int())
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		if (p0).Cmp(p0) <= 0 {
			var _517_v15 _dafny.CodePoint
			_ = _517_v15
			_517_v15 = _dafny.CodePoint('o')
			var _518_v16 _dafny.Map
			_ = _518_v16
			_518_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _517_v15)
			var _519_v17 _dafny.MultiSet
			_ = _519_v17
			_519_v17 = _dafny.MultiSetOf((_this).F6(), (_this).F6(), (_this).F6(), (_this).Fm0((_this).F6(), globalState), (_this).F6())
			var _520_v18 *C0
			_ = _520_v18
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__((_519_v17).Cardinality(), (_this).F8())
			_520_v18 = _nw74
			var _521_v19 _dafny.Map
			_ = _521_v19
			_521_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_520_v18, _517_v15)
			var _522_v20 _dafny.MultiSet
			_ = _522_v20
			_522_v20 = _dafny.MultiSetOf(_518_v16, (_518_v16).Update((_521_v19).Cardinality(), Companion_Default___.Fm11(_517_v15, (_this).F6(), (_this).F6(), globalState)))
			var _523_v21 _dafny.Sequence
			_ = _523_v21
			_523_v21 = _dafny.SeqOf(_522_v20)
			var _524_v22 _dafny.MultiSet
			_ = _524_v22
			_524_v22 = _dafny.MultiSetOf((_520_v18).F13())
			var _525_v23 _dafny.MultiSet
			_ = _525_v23
			_525_v23 = _dafny.MultiSetOf((_524_v22).Cardinality(), Companion_Default___.Fm6((_520_v18).F12(), (_this).F5(), (_this).F6(), (_520_v18).F12(), globalState), Companion_Default___.Fm6((_520_v18).F13(), (_this).F8(), (_this).F6(), (_520_v18).F13(), globalState), (_this).F5(), (_this).F8())
			_522_v20 = (func() _dafny.MultiSet {
				if (_this).F6() {
					return (_523_v21).Select((Companion_Default___.SafeIndex((_525_v23).Cardinality(), _dafny.IntOfUint32((_523_v21).Cardinality()))).Uint32()).(_dafny.MultiSet)
				}
				return _522_v20
			})()
			var _526_v24 _dafny.Sequence
			_ = _526_v24
			_526_v24 = _dafny.SeqOf(_517_v15)
			var _527_v25 _dafny.Sequence
			_ = _527_v25
			_527_v25 = _dafny.SeqOf(_526_v24)
			var _528_v26 _dafny.Map
			_ = _528_v26
			_528_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), !((_this).F6()))
			var _529_v27 _dafny.Sequence
			_ = _529_v27
			_529_v27 = _dafny.SeqOf(Companion_Default___.Fm17((func() bool {
				if (_528_v26).Contains(p0) {
					return (_528_v26).Get(p0).(bool)
				}
				return (_this).F6()
			})(), globalState))
			_527_v25 = (_529_v27).Select((Companion_Default___.SafeIndex(((func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(119), _dafny.IntOfInt64(115))); ; {
					_compr_14, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _530_v28 _dafny.Int
					_530_v28 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(119)).Cmp(_530_v28) <= 0) && ((_530_v28).Cmp(_dafny.IntOfInt64(115)) < 0) {
						_coll14.Add(Companion_Default___.SafeDivisionInt(_530_v28, (_520_v18).F13()), !((_this).F6()))
					}
				}
				return _coll14.ToMap()
			}()).Merge(func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(223), _dafny.IntOfInt64(441))); ; {
					_compr_15, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _531_v29 _dafny.Int
					_531_v29 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(223)).Cmp(_531_v29) <= 0) && ((_531_v29).Cmp(_dafny.IntOfInt64(441)) < 0) {
						_coll15.Add((_531_v29).Plus((_520_v18).F13()), true)
					}
				}
				return _coll15.ToMap()
			}())).Cardinality(), _dafny.IntOfUint32((_529_v27).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _532_v30 _dafny.Array
			_ = _532_v30
			var _nw75 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
			_ = _nw75
			_532_v30 = _nw75
			var _533_v31 _dafny.Set
			_ = _533_v31
			_533_v31 = _dafny.SetOf(_517_v15, _517_v15, _517_v15, _517_v15)
			var _534_v32 *C2
			_ = _534_v32
			var _nw76 *C2 = New_C2_()
			_ = _nw76
			_nw76.Ctor__((_533_v31).Cardinality())
			_534_v32 = _nw76
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_532_v30), 0))
			_ = _index75
			(_532_v30).ArraySet1(_534_v32, (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_532_v30), 0))
			_ = _index76
			(_532_v30).ArraySet1(_534_v32, (_index76).Int())
			var _535_v33 bool
			_ = _535_v33
			_535_v33 = true
			_535_v33 = (_this).F6()
			_517_v15 = _dafny.CodePoint('r')
		} else {
			var _536_v34 _dafny.Map
			_ = _536_v34
			_536_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), (_this).F6())
			var _537_v35 _dafny.Sequence
			_ = _537_v35
			_537_v35 = _dafny.SeqOf(_536_v34, _536_v34)
			var _538_v36 _dafny.Map
			_ = _538_v36
			_538_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), (_this).F6())).Cardinality()))
			var _539_v37 *C1
			_ = _539_v37
			var _nw77 *C1 = New_C1_()
			_ = _nw77
			_nw77.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(274))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}(func(_540_i2 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
			})), _537_v35)).Cardinality()), (_this).F5(), (_538_v36).Cardinality(), (func() bool {
				if (_this).F6() {
					return (_this).F6()
				}
				return (_this).F6()
			})())
			_539_v37 = _nw77
			var _541_v38 _dafny.Array
			_ = _541_v38
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw78
			_541_v38 = _nw78
			var _542_v39 _dafny.Sequence
			_ = _542_v39
			_542_v39 = _dafny.SeqOf(true)
			var _543_v40 _dafny.Map
			_ = _543_v40
			_543_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _542_v39)
			var _544_v41 D6
			_ = _544_v41
			_544_v41 = Companion_D6_.Create_DC17_(_542_v39)
			var _nwElement0_11 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_543_v40).Contains(_539_v37.F10) {
					return (_543_v40).Get(_539_v37.F10).(_dafny.Sequence)
				}
				return _542_v39
			})()).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39)).Cardinality()))).Uint32(), (_this).F6())
			_ = _nwElement0_11
			var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(28))
			_ = _nw79
			(_nw79).ArraySet1(_nwElement0_11, 0)
			(_nw79).ArraySet1(_542_v39, 1)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), 2)
			(_nw79).ArraySet1((Companion_D6_.Create_DC17_(_542_v39)).Dtor_cf35(), 3)
			(_nw79).ArraySet1(_dafny.SeqOf((_this).F6(), (_this).F6()), 4)
			(_nw79).ArraySet1(Companion_Default___.Fm16(globalState), 5)
			(_nw79).ArraySet1(Companion_Default___.Fm16(globalState), 6)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), 7)
			(_nw79).ArraySet1(_542_v39, 8)
			(_nw79).ArraySet1(_542_v39, 9)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), 10)
			(_nw79).ArraySet1((Companion_Default___.Fm18(globalState)).Dtor_cf35(), 11)
			(_nw79).ArraySet1(_dafny.SeqOf(true), 12)
			(_nw79).ArraySet1(_542_v39, 13)
			(_nw79).ArraySet1(_542_v39, 14)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_542_v39, (_544_v41).Dtor_cf35()), 15)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), 16)
			(_nw79).ArraySet1(_542_v39, 17)
			(_nw79).ArraySet1(_dafny.SeqOf((_this).F6()), 18)
			(_nw79).ArraySet1(_dafny.SeqOf((_this).F6()), 19)
			(_nw79).ArraySet1(_dafny.SeqOf((_this).F6(), (_this).F6()), 20)
			(_nw79).ArraySet1(_542_v39, 21)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), (Companion_Default___.SafeIndex(_539_v37.F10, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39)).Cardinality()))).Uint32(), false), 22)
			(_nw79).ArraySet1(_542_v39, 23)
			(_nw79).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_542_v39, _542_v39), 24)
			(_nw79).ArraySet1(_542_v39, 25)
			(_nw79).ArraySet1(_dafny.SeqOf(true), 26)
			(_nw79).ArraySet1(_542_v39, 27)
			_541_v38 = _nw79
			var _545_v42 _dafny.Array
			_ = _545_v42
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_8
			var _nw80 _dafny.Array
			_ = _nw80
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw80 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) bool = func(_546_i3 _dafny.Int) bool {
					return (_this).F6()
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw80 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw80).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw80).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_545_v42 = _nw80
			var _547_v43 _dafny.Sequence
			_ = _547_v43
			_547_v43 = _dafny.SeqOf((_this).F5())
			var _548_v44 _dafny.Sequence
			_ = _548_v44
			_548_v44 = _dafny.SeqOf(_dafny.IntOfUint32((_547_v43).Cardinality()))
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_545_v42), 0))
			_ = _index77
			(_545_v42).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_548_v44, _547_v43), (_index77).Int())
			var _549_v45 _dafny.Array
			_ = _549_v45
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw81
			_549_v45 = _nw81
			var _550_v46 D3
			_ = _550_v46
			_550_v46 = Companion_D3_.Create_DC11_(_549_v45)
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_545_v42), 0))
			_ = _index78
			var _rhs49 bool = false
			_ = _rhs49
			var _rhs50 D3 = _550_v46
			_ = _rhs50
			var _lhs37 _dafny.Array = _545_v42
			_ = _lhs37
			var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_545_v42), 0))
			_ = _lhs38
			(_lhs37).ArraySet1(_rhs49, (_lhs38).Int())
			_550_v46 = _rhs50
			var _551_v47 _dafny.CodePoint
			_ = _551_v47
			_551_v47 = _dafny.CodePoint('n')
			_551_v47 = _551_v47
			var _552_v48 _dafny.MultiSet
			_ = _552_v48
			_552_v48 = _dafny.MultiSetOf(_dafny.IntOfInt64(622))
			var _553_v49 _dafny.Sequence
			_ = _553_v49
			_553_v49 = _dafny.SeqOf(_552_v48)
			var _554_v50 _dafny.Sequence
			_ = _554_v50
			_554_v50 = _dafny.UnicodeSeqOfUtf8Bytes("fmwqpjm")
			var _555_v51 _dafny.Sequence
			_ = _555_v51
			_555_v51 = _dafny.SeqOf(_554_v50)
			var _556_v52 _dafny.Map
			_ = _556_v52
			_556_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_553_v49).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_553_v49).Cardinality()))).Uint32()).(_dafny.MultiSet), (_555_v51).Select((Companion_Default___.SafeIndex((_539_v37).F11(), _dafny.IntOfUint32((_555_v51).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_545_v42), 0))
			_ = _index79
			(_545_v42).ArraySet1((func() bool {
				if (_539_v37).Fm0((_this).F6(), globalState) {
					return !((_539_v37.F10).Cmp((_556_v52).Cardinality()) > 0)
				}
				return (_552_v48).IsProperSubsetOf(_552_v48)
			})(), (_index79).Int())
		}
		var _557_v53 _dafny.Sequence
		_ = _557_v53
		_557_v53 = _dafny.UnicodeSeqOfUtf8Bytes("whxnlgjoh")
		var _hi4 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_557_v53).Cardinality()))
		_ = _hi4
		for _558_i4 := ((_this).F5()).Minus(Companion_Default___.Fm6(p0, (_this).F5(), (_this).F6(), p0, globalState)); _558_i4.Cmp(_hi4) < 0; _558_i4 = _558_i4.Plus(_dafny.One) {
			var _559_v54 _dafny.Int
			_ = _559_v54
			_559_v54 = _dafny.IntOfInt64(352)
			var _560_v55 _dafny.Sequence
			_ = _560_v55
			_560_v55 = _dafny.SeqOf(false)
			var _561_v56 _dafny.Set
			_ = _561_v56
			_561_v56 = _dafny.SetOf((_this).F6())
			var _562_v57 _dafny.Sequence
			_ = _562_v57
			_562_v57 = _dafny.SeqOf(_dafny.IntOfUint32((_560_v55).Cardinality()), _dafny.IntOfUint32((_557_v53).Cardinality()), (_561_v56).Cardinality())
			_559_v54 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_562_v57, _562_v57)).Cardinality()), _559_v54)
			var _563_v58 bool
			_ = _563_v58
			_563_v58 = true
			var _564_v59 _dafny.MultiSet
			_ = _564_v59
			_564_v59 = _dafny.MultiSetOf(_dafny.IntOfUint32((_557_v53).Cardinality()))
			var _565_v60 _dafny.Map
			_ = _565_v60
			_565_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_564_v59, (_this).F6())
			_563_v58 = (func() bool {
				if (_this).F6() {
					return !(((_565_v60).Cardinality()).Cmp(_dafny.IntOfInt64(924)) >= 0)
				}
				return (_563_v58) == ((_this).F6())
			})()
			var _566_v61 _dafny.CodePoint
			_ = _566_v61
			_566_v61 = _dafny.CodePoint('t')
			var _567_v62 _dafny.Array
			_ = _567_v62
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw82
			_567_v62 = _nw82
			var _568_v63 D6
			_ = _568_v63
			_568_v63 = Companion_D6_.Create_DC18_((_this).F6())
			var _rhs51 bool = (_this).Fm0(((_568_v63).Dtor_cf36()) == ((_this).F6()), globalState)
			_ = _rhs51
			var _rhs52 _dafny.Int = p0
			_ = _rhs52
			var _rhs53 _dafny.CodePoint = _566_v61
			_ = _rhs53
			var _rhs54 _dafny.Array = _567_v62
			_ = _rhs54
			var _rhs55 _dafny.Int = p0
			_ = _rhs55
			_563_v58 = _rhs51
			_559_v54 = _rhs52
			_566_v61 = _rhs53
			_567_v62 = _rhs54
			_559_v54 = _rhs55
			_563_v58 = (_this).F6()
		}
		var _569_v64 _dafny.CodePoint
		_ = _569_v64
		_569_v64 = _dafny.CodePoint('o')
		var _570_v65 _dafny.Array
		_ = _570_v65
		var _nwElement0_12 _dafny.CodePoint = _569_v64
		_ = _nwElement0_12
		var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
		_ = _nw83
		(_nw83).ArraySet1CodePoint(_nwElement0_12, 0)
		(_nw83).ArraySet1CodePoint(_569_v64, 1)
		(_nw83).ArraySet1CodePoint(_569_v64, 2)
		(_nw83).ArraySet1CodePoint(_569_v64, 3)
		_570_v65 = _nw83
		var _571_v66 _dafny.Sequence
		_ = _571_v66
		_571_v66 = _dafny.SeqOf(_570_v65)
		var _572_v67 _dafny.MultiSet
		_ = _572_v67
		_572_v67 = _dafny.MultiSetOf((_this).F6(), true)
		var _573_v68 D5
		_ = _573_v68
		_573_v68 = Companion_D5_.Create_DC15_(_572_v67)
		var _574_v69 _dafny.Array
		_ = _574_v69
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_9
		var _nw84 _dafny.Array
		_ = _nw84
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw84 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Sequence = func(_575_i5 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F6(), (_this).F6()), _dafny.SeqOf((_this).F6(), (_this).F6()))
			}
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw84 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw84).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw84).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_574_v69 = _nw84
		var _576_v70 _dafny.Sequence
		_ = _576_v70
		_576_v70 = _dafny.SeqOf((_this).F6(), (_this).F6())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_574_v69), 0))
		_ = _index80
		(_574_v69).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_576_v70, _576_v70), (_index80).Int())
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_574_v69), 0))
		_ = _index81
		(_574_v69).ArraySet1(_576_v70, (_index81).Int())
		var _577_v71 bool
		_ = _577_v71
		_577_v71 = false
		var _578_v72 D1
		_ = _578_v72
		_578_v72 = Companion_D1_.Create_DC7_(Companion_D1_.Create_DC4_(_577_v71, _577_v71, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}((func(_579_v64 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_580_i6 _dafny.Int) _dafny.CodePoint {
				return _579_v64
			}
		})(_569_v64))), p0, Companion_Default___.Fm6((_this).F8(), _dafny.IntOfInt64(-992), _577_v71, _dafny.IntOfInt64(918), globalState)))
		var _581_v73 _dafny.Map
		_ = _581_v73
		_581_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_578_v72, _572_v67)
		var _pat_let_tv15 = _581_v73
		_ = _pat_let_tv15
		var _pat_let_tv16 = _577_v71
		_ = _pat_let_tv16
		var _pat_let_tv17 = globalState
		_ = _pat_let_tv17
		var _pat_let_tv18 = _581_v73
		_ = _pat_let_tv18
		var _pat_let_tv19 = _577_v71
		_ = _pat_let_tv19
		var _pat_let_tv20 = globalState
		_ = _pat_let_tv20
		var _pat_let_tv21 = _572_v67
		_ = _pat_let_tv21
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_574_v69), 0))
		_ = _index82
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_574_v69), 0))
		_ = _index83
		var _rhs56 _dafny.Sequence = _571_v66
		_ = _rhs56
		var _rhs57 D5 = func(_pat_let21_0 D5) D5 {
			return func(_582_dt__update__tmp_h1 D5) D5 {
				return func(_pat_let22_0 _dafny.MultiSet) D5 {
					return func(_583_dt__update_hcf31_h0 _dafny.MultiSet) D5 {
						return Companion_D5_.Create_DC15_(_583_dt__update_hcf31_h0)
					}(_pat_let22_0)
				}((func() _dafny.MultiSet {
					if (_pat_let_tv15).Contains(Companion_D1_.Create_DC7_(Companion_Default___.Fm14(_pat_let_tv16, _pat_let_tv17))) {
						return (_pat_let_tv18).Get(Companion_D1_.Create_DC7_(Companion_Default___.Fm14(_pat_let_tv19, _pat_let_tv20))).(_dafny.MultiSet)
					}
					return _pat_let_tv21
				})())
			}(_pat_let21_0)
		}(Companion_D5_.Create_DC15_(_572_v67))
		_ = _rhs57
		var _rhs58 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_576_v70, _576_v70)
		_ = _rhs58
		var _rhs59 _dafny.Sequence = Companion_Default___.Fm16(globalState)
		_ = _rhs59
		var _rhs60 bool = _577_v71
		_ = _rhs60
		var _lhs39 _dafny.Array = _574_v69
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_574_v69), 0))
		_ = _lhs40
		var _lhs41 _dafny.Array = _574_v69
		_ = _lhs41
		var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(680), _dafny.ArrayLen((_574_v69), 0))
		_ = _lhs42
		_571_v66 = _rhs56
		_573_v68 = _rhs57
		(_lhs39).ArraySet1(_rhs58, (_lhs40).Int())
		(_lhs41).ArraySet1(_rhs59, (_lhs42).Int())
		_577_v71 = _rhs60
		var _584_v74 _dafny.Array
		_ = _584_v74
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_10
		var _nw85 _dafny.Array
		_ = _nw85
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw85 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Int = (func(_585_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_586_i8 _dafny.Int) _dafny.Int {
					return (_586_i8).Minus(_585_p0)
				}
			})(p0)
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
		_584_v74 = _nw85
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_584_v74), 0))); ; {
			_guard_loop_1, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _587_i7 _dafny.Int
			_587_i7 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_587_i7).Sign() != -1) && ((_587_i7).Cmp(_dafny.ArrayLen((_584_v74), 0)) < 0)) {
				(_584_v74).ArraySet1(Companion_Default___.SafeDivisionInt(_587_i7, p0), (_587_i7).Int())
			}
		}
		var _588_v75 _dafny.Set
		_ = _588_v75
		_588_v75 = _dafny.SetOf(_577_v71, (_this).F6())
		var _589_v76 _dafny.Sequence
		_ = _589_v76
		_589_v76 = _dafny.SeqOf((_dafny.Zero).Minus((_588_v75).Cardinality()))
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_584_v74), 0))
		_ = _index84
		(_584_v74).ArraySet1((_589_v76).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_557_v53).Cardinality()), _dafny.IntOfUint32((_589_v76).Cardinality()))).Uint32()).(_dafny.Int), (_index84).Int())
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_584_v74), 0))
		_ = _index85
		(_584_v74).ArraySet1(((_this).F5()).Plus(((_588_v75).Difference(_588_v75)).Cardinality()), (_index85).Int())
	}
}
func (_this *C3) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Array, D0) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 D0 = Companion_D0_.Default()
		_ = r1
		var _590_v0 _dafny.Int
		_ = _590_v0
		_590_v0 = _dafny.IntOfInt64(-393)
		var _591_v1 _dafny.Set
		_ = _591_v1
		_591_v1 = _dafny.SetOf(p2, _dafny.IntOfInt64(430))
		var _592_v2 _dafny.MultiSet
		_ = _592_v2
		_592_v2 = _dafny.MultiSetOf((_591_v1).Cardinality(), _590_v0)
		var _593_v3 T0
		_ = _593_v3
		var _nw86 *C1 = New_C1_()
		_ = _nw86
		_nw86.Ctor__((_592_v2).Cardinality(), (_this).F8(), _590_v0, (_this).F6())
		_593_v3 = _nw86
		var _594_v4 D5
		_ = _594_v4
		_594_v4 = Companion_D5_.Create_DC16_(_593_v3, (_this).F6(), p0)
		var _595_v5 T0
		_ = _595_v5
		var _nw87 *C1 = New_C1_()
		_ = _nw87
		_nw87.Ctor__(_dafny.IntOfInt64(-929), (_594_v4).Dtor_cf34(), p0, (_593_v3).F6())
		_595_v5 = _nw87
		var _596_v6 _dafny.Sequence
		_ = _596_v6
		_596_v6 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfInt64(968))
		var _597_v7 _dafny.Set
		_ = _597_v7
		_597_v7 = _dafny.SetOf((_this).F6(), (_this).F6(), (_595_v5).F6(), (_this).F6())
		var _598_v8 _dafny.Sequence
		_ = _598_v8
		_598_v8 = _dafny.SeqOf((_595_v5).F6(), true)
		var _599_v9 _dafny.Map
		_ = _599_v9
		_599_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_597_v7).Cardinality(), _dafny.IntOfUint32((_598_v8).Cardinality()))
		var _600_v10 _dafny.Map
		_ = _600_v10
		_600_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(11), (_593_v3).F6())
		var _601_v11 _dafny.Map
		_ = _601_v11
		_601_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_595_v5).F6(), (func() bool {
			if (_600_v10).Contains((_595_v5).F5()) {
				return (_600_v10).Get((_595_v5).F5()).(bool)
			}
			return (_595_v5).F6()
		})())
		var _602_v12 _dafny.Array
		_ = _602_v12
		var _nwElement0_13 _dafny.Int = p0
		_ = _nwElement0_13
		var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(9))
		_ = _nw88
		(_nw88).ArraySet1(_nwElement0_13, 0)
		(_nw88).ArraySet1(_dafny.IntOfUint32((_596_v6).Cardinality()), 1)
		(_nw88).ArraySet1(_590_v0, 2)
		(_nw88).ArraySet1(((_599_v9).Merge(_599_v9)).Cardinality(), 3)
		(_nw88).ArraySet1((_dafny.IntOfInt64(665)).Plus(_dafny.IntOfInt64(780)), 4)
		(_nw88).ArraySet1((_595_v5).F5(), 5)
		(_nw88).ArraySet1(Companion_Default___.SafeDivisionInt((_593_v3).F5(), (_this).F5()), 6)
		(_nw88).ArraySet1((_593_v3).F5(), 7)
		(_nw88).ArraySet1(Companion_Default___.Fm6(p0, (_601_v11).Cardinality(), (_595_v5).F6(), Companion_Default___.Fm6(Companion_Default___.Fm6((_593_v3).F5(), (_this).F8(), (_593_v3).F6(), (_595_v5).F5(), globalState), _590_v0, p3, p0, globalState), globalState), 8)
		_602_v12 = _nw88
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
		_ = _index86
		(_602_v12).ArraySet1(Companion_Default___.SafeModuloInt((_this).F8(), (_593_v3).F5()), (_index86).Int())
		var _603_v13 D1
		_ = _603_v13
		_603_v13 = Companion_D1_.Create_DC6_(((_593_v3).F6()) || ((_this).F6()), (_dafny.Zero).Minus((_593_v3).F5()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("a")).Cardinality()), (_595_v5).F6())
		var _604_v14 _dafny.MultiSet
		_ = _604_v14
		_604_v14 = _dafny.MultiSetOf(!((_593_v3).F6()))
		var _605_v15 _dafny.Map
		_ = _605_v15
		_605_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_593_v3).F6()), _595_v5)
		var _606_v16 _dafny.Map
		_ = _606_v16
		_606_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _596_v6)
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
		_ = _index87
		var _rhs61 _dafny.Int = Companion_Default___.Fm6((Companion_Default___.Fm19((_604_v14).Cardinality(), globalState)).Cardinality(), (_dafny.Zero).Minus((_dafny.IntOfUint32((p1).Cardinality())).Times(p0)), p3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("doyklvo")).Cardinality()), globalState)
		_ = _rhs61
		var _rhs62 T0 = (func() T0 {
			if (_605_v15).Contains((_this).F6()) {
				return (_605_v15).Get((_this).F6()).(T0)
			}
			return _593_v3
		})()
		_ = _rhs62
		var _rhs63 _dafny.Int = Companion_Default___.Fm6((_593_v3).F5(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_596_v6, (func() _dafny.Sequence {
			if (_606_v16).Contains((_593_v3).F6()) {
				return (_606_v16).Get((_593_v3).F6()).(_dafny.Sequence)
			}
			return _596_v6
		})())).Cardinality()), (_593_v3).Fm0(p3, globalState), Companion_Default___.SafeModuloInt(_590_v0, p2), globalState)
		_ = _rhs63
		var _rhs64 D1 = (func() D1 {
			if (_this).Fm0(p3, globalState) {
				return _603_v13
			}
			return Companion_Default___.Fm20((_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality())), _604_v14, (_this).F6(), (_this).F6(), globalState)
		})()
		_ = _rhs64
		var _rhs65 T0 = _593_v3
		_ = _rhs65
		var _lhs43 _dafny.Array = _602_v12
		_ = _lhs43
		var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
		_ = _lhs44
		_590_v0 = _rhs61
		_595_v5 = _rhs62
		(_lhs43).ArraySet1(_rhs63, (_lhs44).Int())
		_603_v13 = _rhs64
		_595_v5 = _rhs65
		if (func() bool {
			if (_601_v11).Contains((_593_v3).F6()) {
				return (_601_v11).Get((_593_v3).F6()).(bool)
			}
			return (_this).Fm0((_595_v5).F6(), globalState)
		})() {
			var _607_v17 *C0
			_ = _607_v17
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__(p2, (_597_v7).Cardinality())
			_607_v17 = _nw89
			if (_604_v14).IsSubsetOf(_604_v14) {
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _index88
				(_602_v12).ArraySet1((_607_v17).F12(), (_index88).Int())
				_590_v0 = Companion_Default___.Fm6((_590_v0).Times((_595_v5).F5()), (_dafny.Zero).Minus(p2), (_593_v3).F6(), ((_607_v17).F13()).Times(p2), globalState)
				var _608_v18 _dafny.Map
				_ = _608_v18
				_608_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_607_v17).F13(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_609_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})))
				var _610_v19 _dafny.Array
				_ = _610_v19
				var _nwElement0_14 _dafny.Sequence = (func() _dafny.Sequence {
					if (_608_v18).Contains((_607_v17).F12()) {
						return (_608_v18).Get((_607_v17).F12()).(_dafny.Sequence)
					}
					return p1
				})()
				_ = _nwElement0_14
				var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(17))
				_ = _nw90
				(_nw90).ArraySet1(_nwElement0_14, 0)
				(_nw90).ArraySet1(p1, 1)
				(_nw90).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xvmci"), 2)
				(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(962))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_611_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				})), 3)
				(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(908))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}(func(_612_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				})), 4)
				(_nw90).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ea"), 5)
				(_nw90).ArraySet1(p1, 6)
				(_nw90).ArraySet1(p1, 7)
				(_nw90).ArraySet1(p1, 8)
				(_nw90).ArraySet1(p1, 9)
				(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(852))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}(func(_613_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				})), 10)
				(_nw90).ArraySet1(p1, 11)
				(_nw90).ArraySet1(p1, 12)
				(_nw90).ArraySet1(p1, 13)
				(_nw90).ArraySet1(p1, 14)
				(_nw90).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(519))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}(func(_614_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), 15)
				(_nw90).ArraySet1(p1, 16)
				_610_v19 = _nw90
				var _615_v20 _dafny.Map
				_ = _615_v20
				_615_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_595_v5).F5(), _610_v19)
				_615_v20 = (_615_v20).Update((_dafny.Zero).Minus((_this).F8()), _610_v19)
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_610_v19), 0))
				_ = _index89
				(_610_v19).ArraySet1(p1, (_index89).Int())
				var _616_v21 _dafny.Array
				_ = _616_v21
				var _nw91 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
				_ = _nw91
				_616_v21 = _nw91
				var _617_v22 _dafny.CodePoint
				_ = _617_v22
				_617_v22 = _dafny.CodePoint('g')
				var _618_v23 _dafny.CodePoint
				_ = _618_v23
				_618_v23 = _617_v22
				var _619_v24 _dafny.Array
				_ = _619_v24
				var _nwElement0_15 bool = (_593_v3).F6()
				_ = _nwElement0_15
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(22))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_15, 0)
				(_nw92).ArraySet1((_593_v3).F6(), 1)
				(_nw92).ArraySet1((_this).F6(), 2)
				(_nw92).ArraySet1(!((_593_v3).F6()) || (!((func() bool {
					if (_601_v11).Contains((_593_v3).F6()) {
						return (_601_v11).Get((_593_v3).F6()).(bool)
					}
					return (_593_v3).F6()
				})())), 3)
				(_nw92).ArraySet1((_593_v3).F6(), 4)
				(_nw92).ArraySet1(((_dafny.Zero).Minus((_595_v5).F5())).Cmp((_607_v17).F12()) != 0, 5)
				(_nw92).ArraySet1((func() bool {
					if (_593_v3).F6() {
						return !((_593_v3).F6())
					}
					return (_593_v3).F6()
				})(), 6)
				(_nw92).ArraySet1((_593_v3).F6(), 7)
				(_nw92).ArraySet1(!((_595_v5).F6()) || ((_593_v3).F6()), 8)
				(_nw92).ArraySet1((_593_v3).F6(), 9)
				(_nw92).ArraySet1(true, 10)
				(_nw92).ArraySet1((_593_v3).F6(), 11)
				(_nw92).ArraySet1((_593_v3).F6(), 12)
				(_nw92).ArraySet1((_593_v3).F6(), 13)
				(_nw92).ArraySet1(((_595_v5).F6()) && ((_this).F6()), 14)
				(_nw92).ArraySet1((_593_v3).F6(), 15)
				(_nw92).ArraySet1((_595_v5).Fm0((_595_v5).F6(), globalState), 16)
				(_nw92).ArraySet1((_593_v3).F6(), 17)
				(_nw92).ArraySet1((_593_v3).Fm0((_595_v5).F6(), globalState), 18)
				(_nw92).ArraySet1(p3, 19)
				(_nw92).ArraySet1((_595_v5).F6(), 20)
				(_nw92).ArraySet1(!((func() bool {
					if (_601_v11).Contains((_607_v17).Fm8(p3, _dafny.UnicodeSeqOfUtf8Bytes("vjkgdpeu"), (_618_v23), globalState)) {
						return (_601_v11).Get((_607_v17).Fm8(p3, _dafny.UnicodeSeqOfUtf8Bytes("vjkgdpeu"), (_618_v23), globalState)).(bool)
					}
					return (_593_v3).F6()
				})()), 21)
				_619_v24 = _nw92
				var _620_v25 _dafny.Map
				_ = _620_v25
				_620_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(129), _619_v24)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_610_v19), 0))
				_ = _index90
				var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p1, p1)
				_ = _rhs66
				var _rhs67 _dafny.Int = _dafny.IntOfInt64(481)
				_ = _rhs67
				var _rhs68 _dafny.Array = _616_v21
				_ = _rhs68
				var _rhs69 _dafny.Array = (func() _dafny.Array {
					if (_620_v25).Contains(_590_v0) {
						return (_620_v25).Get(_590_v0).(_dafny.Array)
					}
					return _619_v24
				})()
				_ = _rhs69
				var _lhs45 _dafny.Array = _610_v19
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_610_v19), 0))
				_ = _lhs46
				(_lhs45).ArraySet1(_rhs66, (_lhs46).Int())
				_590_v0 = _rhs67
				_616_v21 = _rhs68
				_619_v24 = _rhs69
				var _621_v26 D6
				_ = _621_v26
				_621_v26 = Companion_D6_.Create_DC17_(_598_v8)
				var _pat_let_tv22 = _598_v8
				_ = _pat_let_tv22
				var _pat_let_tv23 = _598_v8
				_ = _pat_let_tv23
				_621_v26 = func(_pat_let23_0 D6) D6 {
					return func(_622_dt__update__tmp_h0 D6) D6 {
						return func(_pat_let24_0 _dafny.Sequence) D6 {
							return func(_623_dt__update_hcf35_h0 _dafny.Sequence) D6 {
								return Companion_D6_.Create_DC17_(_623_dt__update_hcf35_h0)
							}(_pat_let24_0)
						}(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv22, _pat_let_tv23))
					}(_pat_let23_0)
				}(_621_v26)
			} else {
				var _624_v27 D1
				_ = _624_v27
				_624_v27 = Companion_D1_.Create_DC4_((_595_v5).F6(), p3, _dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("ahcbwma")), (_595_v5).F5(), Companion_Default___.SafeDivisionInt((_600_v10).Cardinality(), (_607_v17).F12()))
				_624_v27 = _624_v27
				var _625_v29 _dafny.Sequence
				_ = _625_v29
				_625_v29 = _dafny.SeqOf(_591_v1, func() _dafny.Set {
					var _coll16 = _dafny.NewBuilder()
					_ = _coll16
					for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(626), _dafny.IntOfInt64(843))); ; {
						_compr_16, _ok18 := _iter18()
						if !_ok18 {
							break
						}
						var _626_v28 _dafny.Int
						_626_v28 = interface{}(_compr_16).(_dafny.Int)
						if ((_dafny.IntOfInt64(626)).Cmp(_626_v28) <= 0) && ((_626_v28).Cmp(_dafny.IntOfInt64(843)) < 0) {
							_coll16.Add((_626_v28).Plus((_595_v5).F5()))
						}
					}
					return _coll16.ToSet()
				}(), _591_v1)
				var _627_v31 _dafny.Array
				_ = _627_v31
				var _nwElement0_16 _dafny.Set = (_591_v1).Difference(_591_v1)
				_ = _nwElement0_16
				var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(20))
				_ = _nw93
				(_nw93).ArraySet1(_nwElement0_16, 0)
				(_nw93).ArraySet1(_591_v1, 1)
				(_nw93).ArraySet1(_591_v1, 2)
				(_nw93).ArraySet1(Companion_Default___.Fm21(globalState), 3)
				(_nw93).ArraySet1(_591_v1, 4)
				(_nw93).ArraySet1((_625_v29).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_625_v29).Cardinality()))).Uint32()).(_dafny.Set), 5)
				(_nw93).ArraySet1((_dafny.SetOf((_593_v3).F5())).Union(_591_v1), 6)
				(_nw93).ArraySet1((_591_v1).Difference(_dafny.SetOf(_dafny.IntOfInt64(-219))), 7)
				(_nw93).ArraySet1(_591_v1, 8)
				(_nw93).ArraySet1(_591_v1, 9)
				(_nw93).ArraySet1(_591_v1, 10)
				(_nw93).ArraySet1(func() _dafny.Set {
					var _coll17 = _dafny.NewBuilder()
					_ = _coll17
					for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-558), _dafny.IntOfInt64(-764))); ; {
						_compr_17, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _628_v30 _dafny.Int
						_628_v30 = interface{}(_compr_17).(_dafny.Int)
						if ((_dafny.IntOfInt64(-558)).Cmp(_628_v30) <= 0) && ((_628_v30).Cmp(_dafny.IntOfInt64(-764)) < 0) {
							_coll17.Add(Companion_Default___.SafeDivisionInt(_628_v30, (_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int)))
						}
					}
					return _coll17.ToSet()
				}(), 11)
				(_nw93).ArraySet1((_591_v1).Union(_591_v1), 12)
				(_nw93).ArraySet1(_591_v1, 13)
				(_nw93).ArraySet1(_591_v1, 14)
				(_nw93).ArraySet1(_591_v1, 15)
				(_nw93).ArraySet1(_591_v1, 16)
				(_nw93).ArraySet1(_591_v1, 17)
				(_nw93).ArraySet1(_591_v1, 18)
				(_nw93).ArraySet1(_591_v1, 19)
				_627_v31 = _nw93
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_627_v31), 0))
				_ = _index91
				(_627_v31).ArraySet1(_591_v1, (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_627_v31), 0))
				_ = _index92
				(_627_v31).ArraySet1((_591_v1).Intersection((_591_v1).Union(_591_v1)), (_index92).Int())
				var _629_v32 _dafny.Array
				_ = _629_v32
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_11
				var _nw94 _dafny.Array
				_ = _nw94
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw94 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = (func(_630_v3 T0) func(_dafny.Int) bool {
						return func(_631_i5 _dafny.Int) bool {
							return (_630_v3).F6()
						}
					})(_593_v3)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw94 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw94).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw94).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_629_v32 = _nw94
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_629_v32), 0))
				_ = _index93
				(_629_v32).ArraySet1((_593_v3).F6(), (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_629_v32), 0))
				_ = _index94
				(_629_v32).ArraySet1(!(false) || (!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-432))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}(func(_632_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				})), p1)), (_index94).Int())
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_629_v32), 0))
				_ = _index95
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _index96
				var _rhs70 bool = !(false) || ((_629_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_629_v32), 0))).Int()).(bool))
				_ = _rhs70
				var _rhs71 _dafny.Int = (Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(672))).Plus((_593_v3).F5())
				_ = _rhs71
				var _lhs47 _dafny.Array = _629_v32
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_629_v32), 0))
				_ = _lhs48
				var _lhs49 _dafny.Array = _602_v12
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _lhs50
				(_lhs47).ArraySet1(_rhs70, (_lhs48).Int())
				(_lhs49).ArraySet1(_rhs71, (_lhs50).Int())
				var _633_v33 D3
				_ = _633_v33
				_633_v33 = Companion_D3_.Create_DC11_(_602_v12)
				_633_v33 = _633_v33
			}
			var _634_v34 _dafny.Array
			_ = _634_v34
			var _nwElement0_17 bool = (_593_v3).F6()
			_ = _nwElement0_17
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(9))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_17, 0)
			(_nw95).ArraySet1((_593_v3).F6(), 1)
			(_nw95).ArraySet1(p3, 2)
			(_nw95).ArraySet1(p3, 3)
			(_nw95).ArraySet1((_593_v3).F6(), 4)
			(_nw95).ArraySet1((_593_v3).F6(), 5)
			(_nw95).ArraySet1((_593_v3).F6(), 6)
			(_nw95).ArraySet1(false, 7)
			(_nw95).ArraySet1((_595_v5).F6(), 8)
			_634_v34 = _nw95
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_634_v34), 0))
			_ = _index97
			(_634_v34).ArraySet1((_595_v5).F6(), (_index97).Int())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_634_v34), 0))
			_ = _index98
			(_634_v34).ArraySet1(p3, (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_634_v34), 0))
			_ = _index99
			(_634_v34).ArraySet1((_595_v5).F6(), (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
			_ = _index100
			(_602_v12).ArraySet1((_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int), (_index100).Int())
		} else {
			var _635_v35 _dafny.Map
			_ = _635_v35
			_635_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_596_v6, (_593_v3).F6())
			var _636_v36 _dafny.Array
			_ = _636_v36
			var _nwElement0_18 bool = true
			_ = _nwElement0_18
			var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(23))
			_ = _nw96
			(_nw96).ArraySet1(_nwElement0_18, 0)
			(_nw96).ArraySet1((_595_v5).F6(), 1)
			(_nw96).ArraySet1((_593_v3).F6(), 2)
			(_nw96).ArraySet1((_591_v1).IsDisjointFrom(_591_v1), 3)
			(_nw96).ArraySet1((_595_v5).F6(), 4)
			(_nw96).ArraySet1(true, 5)
			(_nw96).ArraySet1((_593_v3).F6(), 6)
			(_nw96).ArraySet1((_595_v5).F6(), 7)
			(_nw96).ArraySet1((_593_v3).F6(), 8)
			(_nw96).ArraySet1(p3, 9)
			(_nw96).ArraySet1((_595_v5).Fm0((_593_v3).F6(), globalState), 10)
			(_nw96).ArraySet1(((_this).F8()).Cmp(_590_v0) < 0, 11)
			(_nw96).ArraySet1((func() bool {
				if p3 {
					return (_593_v3).F6()
				}
				return p3
			})(), 12)
			(_nw96).ArraySet1((_595_v5).F6(), 13)
			(_nw96).ArraySet1((_595_v5).F6(), 14)
			(_nw96).ArraySet1((_this).F6(), 15)
			(_nw96).ArraySet1((func() bool {
				if (_635_v35).Contains(_dafny.SeqOf(p2, (_595_v5).F5(), (_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int), (_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int))) {
					return (_635_v35).Get(_dafny.SeqOf(p2, (_595_v5).F5(), (_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int), (_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int))).(bool)
				}
				return (_this).F6()
			})(), 16)
			(_nw96).ArraySet1(true, 17)
			(_nw96).ArraySet1((_593_v3).F6(), 18)
			(_nw96).ArraySet1((_this).F6(), 19)
			(_nw96).ArraySet1((_this).F6(), 20)
			(_nw96).ArraySet1(!((_595_v5).F6()) || (p3), 21)
			(_nw96).ArraySet1(!((_598_v8).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.MultiSetOf((_this).Fm0((_593_v3).F6(), globalState), (_this).F6())).Cardinality()), _dafny.IntOfUint32((_598_v8).Cardinality()))).Uint32()).(bool)) || ((_593_v3).F6()), 22)
			_636_v36 = _nw96
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
			_ = _index101
			(_636_v36).ArraySet1(p3, (_index101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
			_ = _index102
			(_636_v36).ArraySet1(p3, (_index102).Int())
			var _637_v37 _dafny.Map
			_ = _637_v37
			_637_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _636_v36)
			_637_v37 = (_637_v37).Update((_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int), _636_v36)
			_598_v8 = Companion_Default___.Fm16(globalState)
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
			_ = _index103
			(_636_v36).ArraySet1((_595_v5).F6(), (_index103).Int())
			var _638_v38 _dafny.Sequence
			_ = _638_v38
			_638_v38 = _dafny.SeqOf(p1)
			var _639_v39 D1
			_ = _639_v39
			_639_v39 = Companion_D1_.Create_DC4_((p0).Cmp(_590_v0) > 0, (_593_v3).Fm0(p3, globalState), p1, ((_593_v3).F5()).Minus(_dafny.IntOfUint32(((_638_v38).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_638_v38).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), (_593_v3).F5())
			var _source8 D1 = _639_v39
			_ = _source8
			if _source8.Is_DC4() {
				var _640___mcc_h0 bool = _source8.Get_().(D1_DC4).Cf8
				_ = _640___mcc_h0
				var _641___mcc_h1 bool = _source8.Get_().(D1_DC4).Cf9
				_ = _641___mcc_h1
				var _642___mcc_h2 _dafny.Sequence = _source8.Get_().(D1_DC4).Cf10
				_ = _642___mcc_h2
				var _643___mcc_h3 _dafny.Int = _source8.Get_().(D1_DC4).Cf11
				_ = _643___mcc_h3
				var _644___mcc_h4 _dafny.Int = _source8.Get_().(D1_DC4).Cf12
				_ = _644___mcc_h4
				var _645_cf12 _dafny.Int = _644___mcc_h4
				_ = _645_cf12
				var _646_cf11 _dafny.Int = _643___mcc_h3
				_ = _646_cf11
				var _647_cf10 _dafny.Sequence = _642___mcc_h2
				_ = _647_cf10
				var _648_cf9 bool = _641___mcc_h1
				_ = _648_cf9
				var _649_cf8 bool = _640___mcc_h0
				_ = _649_cf8
				var _650_v40 _dafny.Map
				_ = _650_v40
				_650_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_648_cf9, (_593_v3).F5())
				var _651_v41 _dafny.Map
				_ = _651_v41
				_651_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_601_v11, ((_dafny.Zero).Minus(Companion_Default___.Fm6((_650_v40).Cardinality(), (_593_v3).F5(), p3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("htbyp")).Cardinality()), globalState))).Cmp((_593_v3).F5()) != 0)
				_651_v41 = _651_v41
				var _652_v42 _dafny.CodePoint
				_ = _652_v42
				_652_v42 = _dafny.CodePoint('e')
				_652_v42 = _652_v42
				var _653_v43 _dafny.Array
				_ = _653_v43
				var _nw97 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
				_ = _nw97
				_653_v43 = _nw97
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_653_v43), 0))
				_ = _index104
				(_653_v43).ArraySet1(_595_v5, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(199), _dafny.ArrayLen((_653_v43), 0))
				_ = _index105
				(_653_v43).ArraySet1((func() T0 {
					if (_605_v15).Contains(p3) {
						return (_605_v15).Get(p3).(T0)
					}
					return _595_v5
				})(), (_index105).Int())
				_649_cf8 = (_593_v3).F6()
			} else if _source8.Is_DC5() {
				var _654___mcc_h5 bool = _source8.Get_().(D1_DC5).Cf13
				_ = _654___mcc_h5
				var _655___mcc_h6 bool = _source8.Get_().(D1_DC5).Cf14
				_ = _655___mcc_h6
				var _656___mcc_h7 _dafny.Int = _source8.Get_().(D1_DC5).Cf15
				_ = _656___mcc_h7
				var _657___mcc_h8 _dafny.Int = _source8.Get_().(D1_DC5).Cf16
				_ = _657___mcc_h8
				var _658_cf16 _dafny.Int = _657___mcc_h8
				_ = _658_cf16
				var _659_cf15 _dafny.Int = _656___mcc_h7
				_ = _659_cf15
				var _660_cf14 bool = _655___mcc_h6
				_ = _660_cf14
				var _661_cf13 bool = _654___mcc_h5
				_ = _661_cf13
				_659_cf15 = (_595_v5).F5()
				var _662_v44 _dafny.CodePoint
				_ = _662_v44
				_662_v44 = _dafny.CodePoint('e')
				var _rhs72 _dafny.Array = _602_v12
				_ = _rhs72
				var _rhs73 bool = (Companion_D1_.Create_DC4_((_595_v5).F6(), (func() bool {
					if (_601_v11).Contains(false) {
						return (_601_v11).Get(false).(bool)
					}
					return p3
				})(), p1, p0, (_595_v5).F5())).Dtor_cf8()
				_ = _rhs73
				var _rhs74 bool = !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("vhrqm"), _662_v44)
				_ = _rhs74
				_602_v12 = _rhs72
				_660_cf14 = _rhs73
				_661_cf13 = _rhs74
				_658_cf16 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(138))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_663_v3 T0) func(_dafny.Int) _dafny.Int {
					return func(_664_i7 _dafny.Int) _dafny.Int {
						return (_663_v3).F5()
					}
				})(_593_v3)))).Cardinality())
				_660_cf14 = (_595_v5).F6()
			} else if _source8.Is_DC6() {
				var _665___mcc_h9 bool = _source8.Get_().(D1_DC6).Cf17
				_ = _665___mcc_h9
				var _666___mcc_h10 _dafny.Int = _source8.Get_().(D1_DC6).Cf18
				_ = _666___mcc_h10
				var _667___mcc_h11 _dafny.Int = _source8.Get_().(D1_DC6).Cf19
				_ = _667___mcc_h11
				var _668___mcc_h12 bool = _source8.Get_().(D1_DC6).Cf20
				_ = _668___mcc_h12
				var _669_cf20 bool = _668___mcc_h12
				_ = _669_cf20
				var _670_cf19 _dafny.Int = _667___mcc_h11
				_ = _670_cf19
				var _671_cf18 _dafny.Int = _666___mcc_h10
				_ = _671_cf18
				var _672_cf17 bool = _665___mcc_h9
				_ = _672_cf17
				var _673_v45 *C2
				_ = _673_v45
				var _nw98 *C2 = New_C2_()
				_ = _nw98
				_nw98.Ctor__((_593_v3).F5())
				_673_v45 = _nw98
				var _674_v46 _dafny.Map
				_ = _674_v46
				_674_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_602_v12, _673_v45)
				var _675_v47 _dafny.Sequence
				_ = _675_v47
				_675_v47 = _dafny.SeqOf(_602_v12)
				var _676_v48 _dafny.Map
				_ = _676_v48
				_676_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C2 {
					if (_674_v46).Contains((_675_v47).Select((Companion_Default___.SafeIndex((_593_v3).F5(), _dafny.IntOfUint32((_675_v47).Cardinality()))).Uint32()).(_dafny.Array)) {
						return (_674_v46).Get((_675_v47).Select((Companion_Default___.SafeIndex((_593_v3).F5(), _dafny.IntOfUint32((_675_v47).Cardinality()))).Uint32()).(_dafny.Array)).(*C2)
					}
					return _673_v45
				})(), _dafny.SetOf((_593_v3).F5(), p2))
				_676_v48 = _676_v48
				_636_v36 = _636_v36
				var _677_v49 _dafny.CodePoint
				_ = _677_v49
				_677_v49 = _dafny.CodePoint('k')
				_669_cf20 = !((_673_v45).Fm1(p0, _677_v49, globalState))
				var _678_v50 _dafny.Array
				_ = _678_v50
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw99
				_678_v50 = _nw99
				var _679_v52 _dafny.Sequence
				_ = _679_v52
				_679_v52 = _dafny.SeqOf(_599_v9)
				var _680_v53 _dafny.Array
				_ = _680_v53
				var _nwElement0_19 _dafny.Map = _599_v9
				_ = _nwElement0_19
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(15))
				_ = _nw100
				(_nw100).ArraySet1(_nwElement0_19, 0)
				(_nw100).ArraySet1(_599_v9, 1)
				(_nw100).ArraySet1(_599_v9, 2)
				(_nw100).ArraySet1(_599_v9, 3)
				(_nw100).ArraySet1(func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter20 := _dafny.Iterate((_599_v9).Keys().Elements()); ; {
						_compr_18, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _681_v51 _dafny.Int
						_681_v51 = interface{}(_compr_18).(_dafny.Int)
						if (_599_v9).Contains(_681_v51) {
							_coll18.Add((_681_v51).Times(_670_cf19), _670_cf19)
						}
					}
					return _coll18.ToMap()
				}(), 4)
				(_nw100).ArraySet1((_599_v9).Update((_593_v3).F5(), (_595_v5).F5()), 5)
				(_nw100).ArraySet1(_599_v9, 6)
				(_nw100).ArraySet1(_599_v9, 7)
				(_nw100).ArraySet1(_599_v9, 8)
				(_nw100).ArraySet1((_679_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-264), _dafny.IntOfUint32((_679_v52).Cardinality()))).Uint32()).(_dafny.Map), 9)
				(_nw100).ArraySet1(_599_v9, 10)
				(_nw100).ArraySet1((_599_v9).Update(_dafny.IntOfInt64(967), (_this).F8()), 11)
				(_nw100).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(515), _590_v0), 12)
				(_nw100).ArraySet1(_599_v9, 13)
				(_nw100).ArraySet1(_599_v9, 14)
				_680_v53 = _nw100
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_678_v50), 0))
				_ = _index106
				(_678_v50).ArraySet1(_680_v53, (_index106).Int())
				var _682_v54 _dafny.Sequence
				_ = _682_v54
				_682_v54 = _dafny.SeqOf(_680_v53, _680_v53, _680_v53, _680_v53)
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_678_v50), 0))
				_ = _index107
				(_678_v50).ArraySet1((_682_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-573), _dafny.IntOfUint32((_682_v54).Cardinality()))).Uint32()).(_dafny.Array), (_index107).Int())
			} else if _source8.Is_DC3() {
				var _683___mcc_h13 _dafny.Map = _source8.Get_().(D1_DC3).Cf7
				_ = _683___mcc_h13
				var _684_cf7 _dafny.Map = _683___mcc_h13
				_ = _684_cf7
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
				_ = _index108
				(_636_v36).ArraySet1((_595_v5).F6(), (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _index109
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _index110
				var _rhs75 _dafny.Int = (_595_v5).F5()
				_ = _rhs75
				var _rhs76 _dafny.Int = (_593_v3).F5()
				_ = _rhs76
				var _lhs51 _dafny.Array = _602_v12
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _lhs52
				var _lhs53 _dafny.Array = _602_v12
				_ = _lhs53
				var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _lhs54
				(_lhs51).ArraySet1(_rhs75, (_lhs52).Int())
				(_lhs53).ArraySet1(_rhs76, (_lhs54).Int())
				var _685_v55 D3
				_ = _685_v55
				_685_v55 = Companion_D3_.Create_DC11_(_602_v12)
				var _686_v56 _dafny.Map
				_ = _686_v56
				_686_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_685_v55, (_595_v5).F5())
				_686_v56 = (_686_v56).Update((func() D3 {
					if (_this).F6() {
						return _685_v55
					}
					return Companion_D3_.Create_DC11_(_602_v12)
				})(), (_593_v3).F5())
				var _687_v57 _dafny.Map
				_ = _687_v57
				_687_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_595_v5).F5(), p1)
				_687_v57 = (_687_v57).Update(((_595_v5).F5()).Plus(p0), p1)
			} else {
				var _688___mcc_h14 D1 = _source8.Get_().(D1_DC7).Cf21
				_ = _688___mcc_h14
				var _689_cf21 D1 = _688___mcc_h14
				_ = _689_cf21
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
				_ = _index111
				(_636_v36).ArraySet1((_593_v3).F6(), (_index111).Int())
				_590_v0 = p2
				var _690_v58 _dafny.Array
				_ = _690_v58
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
				_ = _nw101
				_690_v58 = _nw101
				var _691_v59 _dafny.Sequence
				_ = _691_v59
				_691_v59 = _dafny.SeqOf(_636_v36, _636_v36)
				var _692_v60 _dafny.Map
				_ = _692_v60
				_692_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_603_v13).Dtor_cf19())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
				_ = _index112
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _index113
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _index114
				var _rhs77 bool = !((true) == (!_dafny.Companion_Sequence_.Contains(_691_v59, _636_v36)))
				_ = _rhs77
				var _rhs78 _dafny.Array = _690_v58
				_ = _rhs78
				var _rhs79 _dafny.Int = Companion_Default___.SafeModuloInt(((_604_v14).Difference(_604_v14)).Cardinality(), (_692_v60).Cardinality())
				_ = _rhs79
				var _rhs80 _dafny.Int = Companion_Default___.Fm6(_590_v0, ((_593_v3).F5()).Plus(p2), true, _dafny.IntOfInt64(197), globalState)
				_ = _rhs80
				var _lhs55 _dafny.Array = _636_v36
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
				_ = _lhs56
				var _lhs57 _dafny.Array = _602_v12
				_ = _lhs57
				var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _lhs58
				var _lhs59 _dafny.Array = _602_v12
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))
				_ = _lhs60
				(_lhs55).ArraySet1(_rhs77, (_lhs56).Int())
				_690_v58 = _rhs78
				(_lhs57).ArraySet1(_rhs79, (_lhs58).Int())
				(_lhs59).ArraySet1(_rhs80, (_lhs60).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_636_v36), 0))
				_ = _index115
				(_636_v36).ArraySet1(true, (_index115).Int())
			}
		}
		var _693_v61 *C2
		_ = _693_v61
		var _nw102 *C2 = New_C2_()
		_ = _nw102
		_nw102.Ctor__(((_602_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_602_v12), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus(_590_v0)))
		_693_v61 = _nw102
		var _694_i8 _dafny.Int
		_ = _694_i8
		_694_i8 = _dafny.Zero
		{
			for !((_595_v5).F6()) {
				{
					if (_694_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_694_i8 = (_694_i8).Plus(_dafny.One)
					var _695_v62 _dafny.Array
					_ = _695_v62
					var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
					_ = _nw103
					_695_v62 = _nw103
					var _696_v63 _dafny.Map
					_ = _696_v63
					_696_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
						if !(false) {
							return p1
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(441))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg38 _dafny.Int) interface{} {
								return coer38(arg38)
							}
						}(func(_697_i9 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						}))
					})(), _695_v62)
					_696_v63 = (_696_v63).Update(_dafny.Companion_Sequence_.Concatenate(p1, p1), _695_v62)
					var _698_v64 _dafny.Map
					_ = _698_v64
					_698_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_593_v3).F6(), (_593_v3).F5())
					var _699_v65 _dafny.Sequence
					_ = _699_v65
					_699_v65 = _dafny.SeqOf(_596_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(141))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg39 _dafny.Int) interface{} {
							return coer39(arg39)
						}
					}((func(_700_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_701_i10 _dafny.Int) _dafny.Int {
							return _700_p2
						}
					})(p2))))
					_698_v64 = (_698_v64).Update(_dafny.Companion_Sequence_.IsPrefixOf((_699_v65).Select((Companion_Default___.SafeIndex((_595_v5).F5(), _dafny.IntOfUint32((_699_v65).Cardinality()))).Uint32()).(_dafny.Sequence), _596_v6), (_dafny.Zero).Minus((_595_v5).F5()))
					var _702_v66 *C2
					_ = _702_v66
					var _nw104 *C2 = New_C2_()
					_ = _nw104
					_nw104.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(585), (_dafny.Zero).Minus((_this).F8())))
					_702_v66 = _nw104
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_695_v62), 0))
					_ = _index116
					(_695_v62).ArraySet1(p1, (_index116).Int())
					var _703_v67 bool
					_ = _703_v67
					_703_v67 = true
					var _704_v68 _dafny.Map
					_ = _704_v68
					_704_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_595_v5).F5(), _698_v64)
					var _705_v69 _dafny.Sequence
					_ = _705_v69
					_705_v69 = _dafny.SeqOf(Companion_Default___.Fm5(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((_dafny.SeqOf((_595_v5).F5())).Cardinality())), _704_v68, globalState), p1)
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_695_v62), 0))
					_ = _index117
					var _rhs81 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p1, (_705_v69).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_598_v8).Cardinality()), _dafny.IntOfUint32((_705_v69).Cardinality()))).Uint32()).(_dafny.Sequence))
					_ = _rhs81
					var _rhs82 _dafny.Int = (_this).F5()
					_ = _rhs82
					var _rhs83 bool = p3
					_ = _rhs83
					var _rhs84 bool = !(((_593_v3).F5()).Cmp((_593_v3).F5()) <= 0) || ((_this).F6())
					_ = _rhs84
					var _rhs85 bool = (_693_v61.F9).Cmp((_595_v5).F5()) != 0
					_ = _rhs85
					var _lhs61 _dafny.Array = _695_v62
					_ = _lhs61
					var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_695_v62), 0))
					_ = _lhs62
					var _lhs63 *C2 = _693_v61
					_ = _lhs63
					(_lhs61).ArraySet1(_rhs81, (_lhs62).Int())
					_lhs63.F9 = _rhs82
					_703_v67 = _rhs83
					_703_v67 = _rhs84
					_703_v67 = _rhs85
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _706_v70 D3
		_ = _706_v70
		_706_v70 = Companion_D3_.Create_DC11_(_602_v12)
		var _707_v71 _dafny.Sequence
		_ = _707_v71
		_707_v71 = _dafny.SeqOf(_706_v70)
		var _708_v72 _dafny.MultiSet
		_ = _708_v72
		_708_v72 = _dafny.MultiSetOf(_706_v70)
		var _709_v73 _dafny.Array
		_ = _709_v73
		var _nwElement0_20 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_707_v71, _707_v71))
		_ = _nwElement0_20
		var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(6))
		_ = _nw105
		(_nw105).ArraySet1(_nwElement0_20, 0)
		(_nw105).ArraySet1((_dafny.MultiSetOf(_706_v70)).Difference(_708_v72), 1)
		(_nw105).ArraySet1(_708_v72, 2)
		(_nw105).ArraySet1(_708_v72, 3)
		(_nw105).ArraySet1(_708_v72, 4)
		(_nw105).ArraySet1((_708_v72).Difference(_708_v72), 5)
		_709_v73 = _nw105
		_709_v73 = (_709_v73)
		var _710_v74 _dafny.Array
		_ = _710_v74
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_12
		var _nw106 _dafny.Array
		_ = _nw106
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw106 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Map = (func(_711_v9 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_712_i11 _dafny.Int) _dafny.Map {
					return _711_v9
				}
			})(_599_v9)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw106 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw106).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw106).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_710_v74 = _nw106
		_710_v74 = _710_v74
		var _713_v75 _dafny.Array
		_ = _713_v75
		var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
		_ = _nw107
		_713_v75 = _nw107
		var _714_v76 _dafny.Sequence
		_ = _714_v76
		_714_v76 = _dafny.SeqOf(_713_v75, _713_v75)
		r0 = (_714_v76).Select((Companion_Default___.SafeIndex(_590_v0, _dafny.IntOfUint32((_714_v76).Cardinality()))).Uint32()).(_dafny.Array)
		var _715_v77 D0
		_ = _715_v77
		_715_v77 = Companion_D0_.Create_DC0_(_596_v6)
		r1 = _715_v77
		return r0, r1
	}
}
func (_this *C3) F8() _dafny.Int {
	{
		return _this._f8
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f5 _dafny.Int
	_f6 bool
	F7  _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this.F7 = _dafny.Zero
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

func (_this *C4) F5() _dafny.Int {
	return _this._f5
}
func (_this *C4) F6() bool {
	return _this._f6
}
func (_this *C4) Ctor__(f7 _dafny.Int, f5 _dafny.Int, f6 bool) {
	{
		(_this).F7 = f7
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C4) Fm0(p0 bool, globalState *GlobalState) bool {
	{
		var _source9 D0 = Companion_D0_.Create_DC0_(_dafny.SeqOf(_this.F7))
		_ = _source9
		if _source9.Is_DC1() {
			var _716___mcc_h0 _dafny.Sequence = _source9.Get_().(D0_DC1).Cf1
			_ = _716___mcc_h0
			var _717___mcc_h1 _dafny.Int = _source9.Get_().(D0_DC1).Cf2
			_ = _717___mcc_h1
			var _718___mcc_h2 _dafny.Map = _source9.Get_().(D0_DC1).Cf3
			_ = _718___mcc_h2
			var _719___mcc_h3 _dafny.Sequence = _source9.Get_().(D0_DC1).Cf4
			_ = _719___mcc_h3
			var _720_cf4 _dafny.Sequence = _719___mcc_h3
			_ = _720_cf4
			var _721_cf3 _dafny.Map = _718___mcc_h2
			_ = _721_cf3
			var _722_cf2 _dafny.Int = _717___mcc_h1
			_ = _722_cf2
			var _723_cf1 _dafny.Sequence = _716___mcc_h0
			_ = _723_cf1
			return true
		} else if _source9.Is_DC2() {
			var _724___mcc_h4 _dafny.Map = _source9.Get_().(D0_DC2).Cf5
			_ = _724___mcc_h4
			var _725___mcc_h5 bool = _source9.Get_().(D0_DC2).Cf6
			_ = _725___mcc_h5
			var _726_cf6 bool = _725___mcc_h5
			_ = _726_cf6
			var _727_cf5 _dafny.Map = _724___mcc_h4
			_ = _727_cf5
			return (_dafny.SetOf((_this).F6())).Contains((_this).F6())
		} else {
			var _728___mcc_h6 _dafny.Sequence = _source9.Get_().(D0_DC0).Cf0
			_ = _728___mcc_h6
			var _729_cf0 _dafny.Sequence = _728___mcc_h6
			_ = _729_cf0
			return !((_this).F6())
		}
	}
}
func (_this *C4) M1(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _730_v0 _dafny.Sequence
		_ = _730_v0
		_730_v0 = _dafny.SeqOf((_this).F6(), (_this).F6())
		var _731_v1 _dafny.Sequence
		_ = _731_v1
		_731_v1 = _dafny.SeqOf(_this.F7, p0, p0, _dafny.IntOfUint32((_730_v0).Cardinality()), p0)
		var _732_v2 D0
		_ = _732_v2
		_732_v2 = Companion_D0_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _731_v1), (_this).F6())
		var _733_v3 _dafny.Set
		_ = _733_v3
		_733_v3 = _dafny.SetOf(_this.F7, _this.F7)
		var _734_v4 D1
		_ = _734_v4
		_734_v4 = Companion_D1_.Create_DC6_((_this).F6(), _this.F7, (_733_v3).Cardinality(), false)
		var _735_v5 *C0
		_ = _735_v5
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fcehnr")).Cardinality()), (_734_v4).Dtor_cf19())
		_735_v5 = _nw108
		var _736_v6 _dafny.Map
		_ = _736_v6
		_736_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_735_v5, (_735_v5).F12())
		var _737_v7 _dafny.MultiSet
		_ = _737_v7
		_737_v7 = _dafny.MultiSetOf(_736_v6, _736_v6)
		var _738_v8 _dafny.Map
		_ = _738_v8
		_738_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-102), _this.F7)
		var _739_v9 *C3
		_ = _739_v9
		var _nw109 *C3 = New_C3_()
		_ = _nw109
		_nw109.Ctor__((_735_v5).F12(), (_dafny.Zero).Minus((_738_v8).Cardinality()), (_this).F6())
		_739_v9 = _nw109
		var _740_v10 T0
		_ = _740_v10
		var _nw110 *C3 = New_C3_()
		_ = _nw110
		_nw110.Ctor__((_737_v7).Cardinality(), (_dafny.SetOf(_739_v9, _739_v9)).Cardinality(), (_this).F6())
		_740_v10 = _nw110
		var _741_v11 _dafny.Set
		_ = _741_v11
		_741_v11 = _dafny.SetOf(_740_v10)
		var _742_v12 _dafny.Sequence
		_ = _742_v12
		_742_v12 = _dafny.UnicodeSeqOfUtf8Bytes("nhstj")
		var _743_v13 _dafny.Array
		_ = _743_v13
		var _nwElement0_21 bool = (_732_v2).Dtor_cf6()
		_ = _nwElement0_21
		var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(17))
		_ = _nw111
		(_nw111).ArraySet1(_nwElement0_21, 0)
		(_nw111).ArraySet1((_dafny.IntOfInt64(56)).Cmp(p0) > 0, 1)
		(_nw111).ArraySet1((_this).F6(), 2)
		(_nw111).ArraySet1(_dafny.Companion_Sequence_.Equal(_731_v1, _731_v1), 3)
		(_nw111).ArraySet1((_this).F6(), 4)
		(_nw111).ArraySet1((_this).F6(), 5)
		(_nw111).ArraySet1((_this).F6(), 6)
		(_nw111).ArraySet1(!_dafny.Companion_Sequence_.Contains(_731_v1, _this.F7), 7)
		(_nw111).ArraySet1((_this).F6(), 8)
		(_nw111).ArraySet1(!((_this).F6()), 9)
		(_nw111).ArraySet1((_this).F6(), 10)
		(_nw111).ArraySet1((_this).F6(), 11)
		(_nw111).ArraySet1((_741_v11).IsDisjointFrom(_741_v11), 12)
		(_nw111).ArraySet1(_dafny.Companion_Sequence_.Equal(_742_v12, _742_v12), 13)
		(_nw111).ArraySet1(((_this).F6()) && ((_this).F6()), 14)
		(_nw111).ArraySet1((_740_v10).F6(), 15)
		(_nw111).ArraySet1((_this).F6(), 16)
		_743_v13 = _nw111
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
		_ = _index118
		(_743_v13).ArraySet1((_740_v10).F6(), (_index118).Int())
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
		_ = _index119
		(_743_v13).ArraySet1((_740_v10).F6(), (_index119).Int())
		var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
		_ = _index120
		(_743_v13).ArraySet1((_740_v10).F6(), (_index120).Int())
		if (_this).F6() {
			var _744_v14 _dafny.Array
			_ = _744_v14
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(17))
			_ = _nw112
			_744_v14 = _nw112
			var _745_v15 _dafny.Array
			_ = _745_v15
			_745_v15 = _744_v14
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
			_ = _index121
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
			_ = _index122
			var _rhs86 bool = false
			_ = _rhs86
			var _rhs87 _dafny.Array = _745_v15
			_ = _rhs87
			var _rhs88 bool = !((_740_v10).F6())
			_ = _rhs88
			var _rhs89 _dafny.Sequence = _742_v12
			_ = _rhs89
			var _rhs90 _dafny.Int = (_739_v9).F8()
			_ = _rhs90
			var _lhs64 _dafny.Array = _743_v13
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
			_ = _lhs65
			var _lhs66 _dafny.Array = _743_v13
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
			_ = _lhs67
			var _lhs68 *C4 = _this
			_ = _lhs68
			(_lhs64).ArraySet1(_rhs86, (_lhs65).Int())
			_745_v15 = _rhs87
			(_lhs66).ArraySet1(_rhs88, (_lhs67).Int())
			_742_v12 = _rhs89
			_lhs68.F7 = _rhs90
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))
			_ = _index123
			(_743_v13).ArraySet1((func() bool {
				if (_743_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))).Int()).(bool) {
					return (_743_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))).Int()).(bool)
				}
				return (_740_v10).F6()
			})(), (_index123).Int())
			var _746_v16 _dafny.MultiSet
			_ = _746_v16
			_746_v16 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jijr"))
			_746_v16 = (_746_v16).Difference(_746_v16)
			var _747_v17 _dafny.Sequence
			_ = _747_v17
			var _748_v18 _dafny.Array
			_ = _748_v18
			var _749_v19 _dafny.Int
			_ = _749_v19
			var _750_v20 _dafny.Int
			_ = _750_v20
			var _out40 _dafny.Sequence
			_ = _out40
			var _out41 _dafny.Array
			_ = _out41
			var _out42 _dafny.Int
			_ = _out42
			var _out43 _dafny.Int
			_ = _out43
			_out40, _out41, _out42, _out43 = Companion_Default___.M0(((_735_v5).F12()).Times(_this.F7), _dafny.IntOfInt64(253), globalState)
			_747_v17 = _out40
			_748_v18 = _out41
			_749_v19 = _out42
			_750_v20 = _out43
			_749_v19 = _this.F7
		} else {
			var _751_v21 D5
			_ = _751_v21
			_751_v21 = Companion_D5_.Create_DC16_(_740_v10, (_740_v10).F6(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_740_v10).F5())))
			var _752_v22 _dafny.Map
			_ = _752_v22
			_752_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_751_v21).Dtor_cf33(), (_731_v1).Select((Companion_Default___.SafeIndex(_this.F7, _dafny.IntOfUint32((_731_v1).Cardinality()))).Uint32()).(_dafny.Int))
			var _753_v23 _dafny.Map
			_ = _753_v23
			_753_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_752_v22).Cardinality(), true)
			var _754_v24 *C3
			_ = _754_v24
			var _nw113 *C3 = New_C3_()
			_ = _nw113
			_nw113.Ctor__((_739_v9).F8(), (_753_v23).Cardinality(), false)
			_754_v24 = _nw113
			var _755_v25 *C3
			_ = _755_v25
			var _nw114 *C3 = New_C3_()
			_ = _nw114
			_nw114.Ctor__((_739_v9).F8(), (_this.F7).Times((func() _dafny.Int {
				if (_752_v22).Contains((_743_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))).Int()).(bool)) {
					return (_752_v22).Get((_743_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))).Int()).(bool)).(_dafny.Int)
				}
				return (_dafny.MultiSetOf((_743_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))).Int()).(bool))).Cardinality()
			})()), ((_754_v24).F8()).Cmp((_754_v24).F8()) > 0)
			_755_v25 = _nw114
			var _756_v26 _dafny.CodePoint
			_ = _756_v26
			_756_v26 = _dafny.CodePoint('e')
			_742_v12 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_742_v12, _742_v12), _dafny.Companion_Sequence_.Update(_742_v12, (Companion_Default___.SafeIndex(Companion_Default___.Fm6(_dafny.IntOfInt64(-546), (_735_v5).F12(), (_740_v10).F6(), (_735_v5).F13(), globalState), _dafny.IntOfUint32((_742_v12).Cardinality()))).Uint32(), _756_v26))
			var _757_v27 *C0
			_ = _757_v27
			var _nw115 *C0 = New_C0_()
			_ = _nw115
			_nw115.Ctor__((_755_v25).F8(), (_740_v10).F5())
			_757_v27 = _nw115
			(_this).F7 = Companion_Default___.Fm6((_757_v27).F12(), _this.F7, (_740_v10).F6(), (_this).F5(), globalState)
		}
		var _758_i0 _dafny.Int
		_ = _758_i0
		_758_i0 = _dafny.Zero
		{
			for (_740_v10).F6() {
				{
					if (_758_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_758_i0 = (_758_i0).Plus(_dafny.One)
					var _759_v28 _dafny.Array
					_ = _759_v28
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_13
					var _nw116 _dafny.Array
					_ = _nw116
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw116 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) D5 = (func(_760_v0 _dafny.Sequence) func(_dafny.Int) D5 {
							return func(_761_i1 _dafny.Int) D5 {
								return Companion_D5_.Create_DC15_(_dafny.MultiSetFromSeq(_760_v0))
							}
						})(_730_v0)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw116 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw116).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw116).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_759_v28 = _nw116
					var _762_v29 _dafny.MultiSet
					_ = _762_v29
					_762_v29 = _dafny.MultiSetOf((_740_v10).F6())
					var _763_v30 D5
					_ = _763_v30
					_763_v30 = Companion_D5_.Create_DC15_((_762_v29).Update((_this).F6(), Companion_Default___.Abs(_dafny.IntOfUint32((_742_v12).Cardinality()))))
					var _pat_let_tv24 = _735_v5
					_ = _pat_let_tv24
					var _pat_let_tv25 = _735_v5
					_ = _pat_let_tv25
					var _pat_let_tv26 = globalState
					_ = _pat_let_tv26
					var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_759_v28), 0))
					_ = _index124
					(_759_v28).ArraySet1(func(_pat_let25_0 D5) D5 {
						return func(_764_dt__update__tmp_h0 D5) D5 {
							return func(_pat_let26_0 _dafny.MultiSet) D5 {
								return func(_765_dt__update_hcf31_h0 _dafny.MultiSet) D5 {
									return Companion_D5_.Create_DC15_(_765_dt__update_hcf31_h0)
								}(_pat_let26_0)
							}(Companion_Default___.Fm10(_dafny.MultiSetOf((_pat_let_tv24).F12(), (_pat_let_tv25).F12(), _dafny.IntOfInt64(-863)), (_this).F5(), _pat_let_tv26))
						}(_pat_let25_0)
					}(_763_v30), (_index124).Int())
					var _766_v31 *C1
					_ = _766_v31
					var _nw117 *C1 = New_C1_()
					_ = _nw117
					_nw117.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('i'))).Cardinality()), (_this).F5(), (_739_v9).F8(), false)
					_766_v31 = _nw117
					var _767_v32 _dafny.Map
					_ = _767_v32
					_767_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_740_v10).F6(), _766_v31)
					var _768_v33 _dafny.Map
					_ = _768_v33
					_768_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F7)
					var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_759_v28), 0))
					_ = _index125
					(_759_v28).ArraySet1(Companion_Default___.Fm22(((_767_v32).Merge(_767_v32)).Cardinality(), _742_v12, ((_766_v31).F11()).Minus((Companion_D0_.Create_DC1_(_742_v12, (_739_v9).F8(), _768_v33, _731_v1)).Dtor_cf2()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(126))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg40 _dafny.Int) interface{} {
							return coer40(arg40)
						}
					}(func(_769_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('o')
					})), globalState), (_index125).Int())
					var _770_v34 _dafny.MultiSet
					_ = _770_v34
					_770_v34 = _dafny.MultiSetOf((_dafny.Zero).Minus((_735_v5).F12()), (_dafny.SetOf((_this).F5())).Cardinality())
					var _771_v35 D2
					_ = _771_v35
					_771_v35 = Companion_D2_.Create_DC8_(_770_v34)
					_771_v35 = Companion_Default___.Fm23(_this.F7, (_dafny.IntOfUint32((_742_v12).Cardinality())).Cmp(_dafny.IntOfUint32((_731_v1).Cardinality())) < 0, _dafny.IntOfInt64(237), ((_735_v5).F12()).Plus(_dafny.IntOfUint32((_742_v12).Cardinality())), globalState)
					var _772_v36 _dafny.Map
					_ = _772_v36
					_772_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_739_v9).F8(), !_dafny.Companion_Sequence_.Equal(_742_v12, _742_v12))
					var _773_v37 D5
					_ = _773_v37
					_773_v37 = Companion_D5_.Create_DC16_(_740_v10, (_this).F6(), (_735_v5).F13())
					_772_v36 = (_772_v36).Update((_773_v37).Dtor_cf34(), (_this).F6())
					var _rhs91 _dafny.Int = (_dafny.Zero).Minus((_735_v5).F12())
					_ = _rhs91
					var _rhs92 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_740_v10).F5(), (_740_v10).F6())).Update(((_740_v10).F5()).Times((_735_v5).F13()), (_743_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_743_v13), 0))).Int()).(bool))
					_ = _rhs92
					var _lhs69 *C4 = _this
					_ = _lhs69
					_lhs69.F7 = _rhs91
					_772_v36 = _rhs92
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		_742_v12 = _742_v12
		var _774_v38 _dafny.Array
		_ = _774_v38
		var _nwElement0_22 _dafny.Int = (_this).F5()
		_ = _nwElement0_22
		var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(5))
		_ = _nw118
		(_nw118).ArraySet1(_nwElement0_22, 0)
		(_nw118).ArraySet1(_this.F7, 1)
		(_nw118).ArraySet1((_735_v5).F12(), 2)
		(_nw118).ArraySet1((_740_v10).F5(), 3)
		(_nw118).ArraySet1((_this).F5(), 4)
		_774_v38 = _nw118
		var _775_v39 _dafny.MultiSet
		_ = _775_v39
		_775_v39 = _dafny.MultiSetOf(_774_v38)
		var _776_v40 _dafny.Array
		_ = _776_v40
		var _nwElement0_23 _dafny.MultiSet = _775_v39
		_ = _nwElement0_23
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(7))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_23, 0)
		(_nw119).ArraySet1((_775_v39).Intersection(_775_v39), 1)
		(_nw119).ArraySet1(_dafny.MultiSetOf(_774_v38, _774_v38, _774_v38, _774_v38, _774_v38), 2)
		(_nw119).ArraySet1((_775_v39).Union(_dafny.MultiSetOf(_774_v38)), 3)
		(_nw119).ArraySet1(_775_v39, 4)
		(_nw119).ArraySet1(_dafny.MultiSetOf(_774_v38, _774_v38), 5)
		(_nw119).ArraySet1(_775_v39, 6)
		_776_v40 = _nw119
		var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_776_v40), 0))
		_ = _index126
		(_776_v40).ArraySet1(_775_v39, (_index126).Int())
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_776_v40), 0))
		_ = _index127
		(_776_v40).ArraySet1(_775_v39, (_index127).Int())
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
