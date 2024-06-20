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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	return ((_dafny.IntOfInt64(448)).Times(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(393))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-484))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false, !(true))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((Companion_D7_.Create_DC22_(true, _dafny.IntOfInt64(629), _dafny.IntOfInt64(716))).Dtor_cf35()), _dafny.SeqOf(false, true)))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Map, globalState *GlobalState) D0 {
	var _source0 D7 = Companion_D7_.Create_DC23_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-437)), Companion_D7_.Create_DC21_(_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D2_.Create_DC7_(_dafny.CodePoint('i')), Companion_D2_.Create_DC7_(_dafny.CodePoint('f')), Companion_D2_.Create_DC7_(_dafny.CodePoint('o')), Companion_D2_.Create_DC7_(_dafny.CodePoint('r')), Companion_D2_.Create_DC7_(_dafny.CodePoint('m'))))))).Cardinality())
	_ = _source0
	if _source0.Is_DC22() {
		var _2___mcc_h0 bool = _source0.Get_().(D7_DC22).Cf35
		_ = _2___mcc_h0
		var _3___mcc_h1 _dafny.Int = _source0.Get_().(D7_DC22).Cf36
		_ = _3___mcc_h1
		var _4___mcc_h2 _dafny.Int = _source0.Get_().(D7_DC22).Cf37
		_ = _4___mcc_h2
		var _5_cf37 _dafny.Int = _4___mcc_h2
		_ = _5_cf37
		var _6_cf36 _dafny.Int = _3___mcc_h1
		_ = _6_cf36
		var _7_cf35 bool = _2___mcc_h0
		_ = _7_cf35
		return Companion_D0_.Create_DC1_(Companion_D0_.Create_DC0_(_6_cf36, _dafny.SeqOf(true)))
	} else if _source0.Is_DC23() {
		var _8___mcc_h3 bool = _source0.Get_().(D7_DC23).Cf38
		_ = _8___mcc_h3
		var _9___mcc_h4 _dafny.Int = _source0.Get_().(D7_DC23).Cf39
		_ = _9___mcc_h4
		var _10_cf39 _dafny.Int = _9___mcc_h4
		_ = _10_cf39
		var _11_cf38 bool = _8___mcc_h3
		_ = _11_cf38
		return Companion_D0_.Create_DC1_((Companion_D0_.Create_DC1_(Companion_D0_.Create_DC0_(_10_cf39, _dafny.SeqOf(_11_cf38)))).Dtor_cf2())
	} else {
		var _12___mcc_h5 _dafny.MultiSet = _source0.Get_().(D7_DC21).Cf34
		_ = _12___mcc_h5
		var _13_cf34 _dafny.MultiSet = _12___mcc_h5
		_ = _13_cf34
		return Companion_D0_.Create_DC1_(Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(224), true), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(186))).Cardinality()))).Cardinality(), _dafny.SeqOf(true, false, !(!(false)))))
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_(_dafny.CodePoint('e'))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((func() bool {
		if true {
			return true
		}
		return true
	})(), true)
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Map, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	if true {
		return func() _dafny.Set {
			var _coll0 = _dafny.NewBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), true)).Cardinality(), _dafny.SeqOf(true, false)))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _14_v0 D0
				_14_v0 = interface{}(_compr_0).(D0)
				if (_dafny.MultiSetOf(Companion_D0_.Create_DC0_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), true)).Cardinality(), _dafny.SeqOf(true, false)))).Contains(_14_v0) {
					_coll0.Add(_14_v0)
				}
			}
			return _coll0.ToSet()
		}()
	} else {
		return _dafny.SetOf(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngymbxll")).Cardinality()), _dafny.SeqOf(false)), Companion_D0_.Create_DC0_(_dafny.IntOfInt64(531), _dafny.SeqOf(!(!(true)), !(false), false)))
	}
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.CodePoint {
	var _source1 D0 = Companion_D0_.Create_DC0_((_dafny.SetOf(_dafny.IntOfInt64(934), _dafny.IntOfInt64(600))).Cardinality(), _dafny.SeqOf(true, !(false)))
	_ = _source1
	if _source1.Is_DC0() {
		var _15___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC0).Cf0
		_ = _15___mcc_h0
		var _16___mcc_h1 _dafny.Sequence = _source1.Get_().(D0_DC0).Cf1
		_ = _16___mcc_h1
		var _17_cf1 _dafny.Sequence = _16___mcc_h1
		_ = _17_cf1
		var _18_cf0 _dafny.Int = _15___mcc_h0
		_ = _18_cf0
		return (Companion_D1_.Create_DC2_(_dafny.CodePoint('y'))).Dtor_cf3()
	} else {
		var _19___mcc_h2 D0 = _source1.Get_().(D0_DC1).Cf2
		_ = _19___mcc_h2
		var _20_cf2 D0 = _19___mcc_h2
		_ = _20_cf2
		return _dafny.CodePoint('y')
	}
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(490))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(969)))).Difference((Companion_D8_.Create_DC24_(_dafny.MultiSetOf(_dafny.IntOfInt64(246), (func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false, false)).Cardinality(), true)).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _21_v0 _dafny.Int
			_21_v0 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false, false)).Cardinality(), true)).Contains(_21_v0) {
				_coll1.Add((_21_v0).Plus(_dafny.IntOfInt64(-936)), false)
			}
		}
		return _coll1.ToMap()
	}()).Cardinality()))).Dtor_cf40())
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(451), _dafny.IntOfInt64(803)), false)
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	var r0 bool = false
	_ = r0
	var _22_v0 _dafny.Sequence
	_ = _22_v0
	_22_v0 = _dafny.SeqOf(p1, p0, p1, p2, p3)
	var _23_v1 _dafny.Int
	_ = _23_v1
	_23_v1 = _dafny.IntOfInt64(761)
	var _24_v2 _dafny.Sequence
	_ = _24_v2
	_24_v2 = _dafny.UnicodeSeqOfUtf8Bytes("aeoyscw")
	var _25_v3 _dafny.Array
	_ = _25_v3
	var _nwElement0_0 _dafny.Sequence = _22_v0
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(21))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_22_v0, 1)
	(_nw0).ArraySet1(Companion_Default___.Fm3(_23_v1, _23_v1, true, globalState), 2)
	(_nw0).ArraySet1(_22_v0, 3)
	(_nw0).ArraySet1(_22_v0, 4)
	(_nw0).ArraySet1(_22_v0, 5)
	(_nw0).ArraySet1((func() _dafny.Sequence {
		if p0 {
			return _22_v0
		}
		return _22_v0
	})(), 6)
	(_nw0).ArraySet1(_dafny.SeqOf(true), 7)
	(_nw0).ArraySet1(_22_v0, 8)
	(_nw0).ArraySet1((func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf(!(p0))
		}
		return _22_v0
	})(), 9)
	(_nw0).ArraySet1(_dafny.SeqOf(p3), 10)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_22_v0, Companion_Default___.Fm3(_dafny.IntOfInt64(669), _23_v1, p0, globalState)), 11)
	(_nw0).ArraySet1(_22_v0, 12)
	(_nw0).ArraySet1(_22_v0, 13)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Update(_22_v0, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(!(p2), _24_v2, globalState), _dafny.IntOfUint32((_22_v0).Cardinality()))).Uint32(), Companion_Default___.Fm1(_dafny.UnicodeSeqOfUtf8Bytes("tpcnkcvm"), p3, _23_v1, globalState)), 14)
	(_nw0).ArraySet1(_22_v0, 15)
	(_nw0).ArraySet1(_22_v0, 16)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p3, p1), _22_v0), 17)
	(_nw0).ArraySet1(_22_v0, 18)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_22_v0, _22_v0), 19)
	(_nw0).ArraySet1(_22_v0, 20)
	_25_v3 = _nw0
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_25_v3), 0))
	_ = _index0
	(_25_v3).ArraySet1(_dafny.SeqOf(p2), (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_25_v3), 0))
	_ = _index1
	(_25_v3).ArraySet1(_22_v0, (_index1).Int())
	var _26_v4 _dafny.CodePoint
	_ = _26_v4
	_26_v4 = _dafny.CodePoint('j')
	var _27_v5 _dafny.Array
	_ = _27_v5
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
	_ = _nw1
	_27_v5 = _nw1
	var _28_v6 _dafny.Map
	_ = _28_v6
	_28_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v4, _27_v5)
	var _29_v7 _dafny.Map
	_ = _29_v7
	_29_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
	var _30_v8 _dafny.Sequence
	_ = _30_v8
	_30_v8 = _dafny.SeqOf(_29_v7)
	var _31_v9 _dafny.Map
	_ = _31_v9
	_31_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_30_v8, _27_v5)
	_28_v6 = (_28_v6).Update(_26_v4, (func() _dafny.Array {
		if (_31_v9).Contains(_30_v8) {
			return (_31_v9).Get(_30_v8).(_dafny.Array)
		}
		return _27_v5
	})())
	var _32_v10 _dafny.MultiSet
	_ = _32_v10
	_32_v10 = _dafny.MultiSetOf(p0, !(!(p0)))
	var _rhs0 bool = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
		if p2 {
			return _24_v2
		}
		return _dafny.Companion_Sequence_.Update(_24_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.IntOfUint32((_24_v2).Cardinality()))).Uint32(), (_24_v2).Select((Companion_Default___.SafeIndex(_23_v1, _dafny.IntOfUint32((_24_v2).Cardinality()))).Uint32()).(_dafny.CodePoint))
	})(), _dafny.Companion_Sequence_.Concatenate(_24_v2, _24_v2))
	_ = _rhs0
	var _rhs1 bool = p2
	_ = _rhs1
	var _rhs2 _dafny.MultiSet = _32_v10
	_ = _rhs2
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 *GlobalState = globalState
	_ = _lhs1
	_lhs0.F7 = _rhs0
	_lhs1.F6 = _rhs1
	_32_v10 = _rhs2
	(globalState).F7 = (_23_v1).Cmp((_23_v1).Minus(_23_v1)) != 0
	(globalState).F7 = (func() bool {
		if p0 {
			return !_dafny.Companion_Sequence_.Contains(Companion_Default___.Fm3((_dafny.Zero).Minus(_23_v1), _23_v1, p0, globalState), p2)
		}
		return p2
	})()
	(globalState).F7 = (_23_v1).Cmp(Companion_Default___.SafeDivisionInt((Companion_Default___.Fm13(p0, globalState)).Cardinality(), _23_v1)) != 0
	r0 = !(p1) || (p0)
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _33_v0 _dafny.CodePoint
	_ = _33_v0
	_33_v0 = _dafny.CodePoint('n')
	var _34_v1 _dafny.MultiSet
	_ = _34_v1
	_34_v1 = _dafny.MultiSetOf(_33_v0, _33_v0)
	var _35_v2 _dafny.Int
	_ = _35_v2
	_35_v2 = _dafny.IntOfInt64(863)
	var _36_v3 _dafny.Sequence
	_ = _36_v3
	_36_v3 = _dafny.SeqOf((func() _dafny.Int {
		if (_34_v1).Contains(_33_v0) {
			return (_34_v1).Multiplicity(_33_v0)
		}
		return (_dafny.Zero).Minus(_35_v2)
	})())
	var _37_v4 _dafny.Sequence
	_ = _37_v4
	_37_v4 = _dafny.UnicodeSeqOfUtf8Bytes("yw")
	var _38_v8 bool
	_ = _38_v8
	_38_v8 = false
	var _39_v9 _dafny.Sequence
	_ = _39_v9
	_39_v9 = _dafny.SeqOf(_38_v8, false)
	var _40_v10 _dafny.Sequence
	_ = _40_v10
	_40_v10 = _dafny.SeqOf(_39_v9, _39_v9)
	var _41_globalState *GlobalState
	_ = _41_globalState
	var _nw2 *GlobalState = New_GlobalState_()
	_ = _nw2
	_nw2.Ctor__(_dafny.IntOfInt64(-585), _36_v3, _dafny.IntOfInt64(991), false, _dafny.IntOfInt64(-718), false, true, false, _dafny.IntOfInt64(30), true, true, _37_v4, func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}((func(_42_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_43_i0 _dafny.Int) _dafny.Int {
				return _42_v2
			}
		})(_35_v2)))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _44_v5 _dafny.Int
			_44_v5 = interface{}(_compr_2).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}((func(_45_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_43_i0 _dafny.Int) _dafny.Int {
					return _45_v2
				}
			})(_35_v2))), _44_v5) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_44_v5, _dafny.IntOfUint32((_37_v4).Cardinality())), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _35_v2)).Cardinality())
			}
		}
		return _coll2.ToMap()
	}(), true, func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-414))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}((func(_46_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_47_i1 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll4 = _dafny.NewMapBuilder()
					_ = _coll4
					for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(827), _dafny.IntOfInt64(741))); ; {
						_compr_4, _ok4 := _iter4()
						if !_ok4 {
							break
						}
						var _48_v7 _dafny.Int
						_48_v7 = interface{}(_compr_4).(_dafny.Int)
						if ((_dafny.IntOfInt64(827)).Cmp(_48_v7) <= 0) && ((_48_v7).Cmp(_dafny.IntOfInt64(741)) < 0) {
							_coll4.Add((_48_v7).Plus(_46_v2), false)
						}
					}
					return _coll4.ToMap()
				}()
			}
		})(_35_v2)))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _49_v6 _dafny.Map
			_49_v6 = interface{}(_compr_3).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-414))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_50_v2 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_47_i1 _dafny.Int) _dafny.Map {
					return func() _dafny.Map {
						var _coll5 = _dafny.NewMapBuilder()
						_ = _coll5
						for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(827), _dafny.IntOfInt64(741))); ; {
							_compr_5, _ok5 := _iter5()
							if !_ok5 {
								break
							}
							var _48_v7 _dafny.Int
							_48_v7 = interface{}(_compr_5).(_dafny.Int)
							if ((_dafny.IntOfInt64(827)).Cmp(_48_v7) <= 0) && ((_48_v7).Cmp(_dafny.IntOfInt64(741)) < 0) {
								_coll5.Add((_48_v7).Plus(_50_v2), false)
							}
						}
						return _coll5.ToMap()
					}()
				}
			})(_35_v2))), _49_v6) {
				_coll3.Add(_49_v6, _35_v2)
			}
		}
		return _coll3.ToMap()
	}(), true, _dafny.IntOfInt64(897), _40_v10, true)
	_41_globalState = _nw2
	var _51_v11 D0
	_ = _51_v11
	_51_v11 = Companion_D0_.Create_DC0_(_35_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_38_v8, _38_v8), _39_v9))
	var _source2 D0 = _51_v11
	_ = _source2
	if _source2.Is_DC0() {
		var _52___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
		_ = _52___mcc_h0
		var _53___mcc_h1 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf1
		_ = _53___mcc_h1
		var _54_cf1 _dafny.Sequence = _53___mcc_h1
		_ = _54_cf1
		var _55_cf0 _dafny.Int = _52___mcc_h0
		_ = _55_cf0
		(_41_globalState).F0 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_37_v4).Cardinality()), _35_v2)).Minus(_35_v2)
		_33_v0 = _33_v0
		(_41_globalState).F7 = !(!(_38_v8) || (!(_38_v8)))
		_51_v11 = _51_v11
	} else {
		var _56___mcc_h2 D0 = _source2.Get_().(D0_DC1).Cf2
		_ = _56___mcc_h2
		var _57_cf2 D0 = _56___mcc_h2
		_ = _57_cf2
		var _58_v12 _dafny.Array
		_ = _58_v12
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw3
		_58_v12 = _nw3
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))
		_ = _index2
		(_58_v12).ArraySet1(_35_v2, (_index2).Int())
		var _59_v13 _dafny.Array
		_ = _59_v13
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw4
		_59_v13 = _nw4
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
		_ = _index3
		(_59_v13).ArraySet1(true, (_index3).Int())
		var _60_v14 _dafny.Map
		_ = _60_v14
		_60_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v12, (_35_v2).Plus(_35_v2))
		var _61_v15 _dafny.MultiSet
		_ = _61_v15
		_61_v15 = _dafny.MultiSetOf(_38_v8, _38_v8)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))
		_ = _index4
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
		_ = _index5
		var _rhs3 _dafny.Int = _35_v2
		_ = _rhs3
		var _rhs4 _dafny.Int = (func() _dafny.Int {
			if (_60_v14).Contains(_58_v12) {
				return (_60_v14).Get(_58_v12).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_61_v15).Contains(_38_v8) {
					return (_61_v15).Multiplicity(_38_v8)
				}
				return _35_v2
			})()
		})()
		_ = _rhs4
		var _rhs5 bool = _38_v8
		_ = _rhs5
		var _rhs6 _dafny.Int = (_35_v2).Times(_35_v2)
		_ = _rhs6
		var _lhs2 *GlobalState = _41_globalState
		_ = _lhs2
		var _lhs3 _dafny.Array = _58_v12
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))
		_ = _lhs4
		var _lhs5 _dafny.Array = _59_v13
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
		_ = _lhs6
		var _lhs7 *GlobalState = _41_globalState
		_ = _lhs7
		_lhs2.F16 = _rhs3
		(_lhs3).ArraySet1(_rhs4, (_lhs4).Int())
		(_lhs5).ArraySet1(_rhs5, (_lhs6).Int())
		_lhs7.F16 = _rhs6
		var _62_v16 D0
		_ = _62_v16
		_62_v16 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_37_v4).Cardinality()), _39_v9)
		var _63_v17 D0
		_ = _63_v17
		_63_v17 = Companion_D0_.Create_DC1_(_62_v16)
		var _source3 D0 = _63_v17
		_ = _source3
		if _source3.Is_DC0() {
			var _64___mcc_h3 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
			_ = _64___mcc_h3
			var _65___mcc_h4 _dafny.Sequence = _source3.Get_().(D0_DC0).Cf1
			_ = _65___mcc_h4
			var _66_cf1 _dafny.Sequence = _65___mcc_h4
			_ = _66_cf1
			var _67_cf0 _dafny.Int = _64___mcc_h3
			_ = _67_cf0
			var _68_v18 _dafny.Array
			_ = _68_v18
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw5
			_68_v18 = _nw5
			_68_v18 = _68_v18
			var _69_v19 _dafny.Array
			_ = _69_v19
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_0
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_70_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_71_i2 _dafny.Int) _dafny.Sequence {
						return _70_v4
					}
				})(_37_v4)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw6 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw6).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw6).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_69_v19 = _nw6
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_69_v19), 0))
			_ = _index6
			(_69_v19).ArraySet1(_37_v4, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
			_ = _index7
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_69_v19), 0))
			_ = _index8
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
			_ = _index9
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))
			_ = _index10
			var _rhs7 _dafny.Int = Companion_Default___.Fm0(_dafny.Companion_Sequence_.IsPrefixOf(_66_cf1, _66_cf1), _37_v4, _41_globalState)
			_ = _rhs7
			var _rhs8 bool = ((_59_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))).Int()).(bool)) || (_38_v8)
			_ = _rhs8
			var _rhs9 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_37_v4, (Companion_Default___.SafeIndex((_61_v15).Cardinality(), _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32(), (_37_v4).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_39_v9)).Cardinality(), _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32()).(_dafny.CodePoint))
			_ = _rhs9
			var _rhs10 bool = (_59_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))).Int()).(bool)
			_ = _rhs10
			var _rhs11 _dafny.Int = (_dafny.Zero).Minus((_51_v11).Dtor_cf0())
			_ = _rhs11
			var _lhs8 _dafny.Array = _59_v13
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
			_ = _lhs9
			var _lhs10 _dafny.Array = _69_v19
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_69_v19), 0))
			_ = _lhs11
			var _lhs12 _dafny.Array = _59_v13
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _58_v12
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))
			_ = _lhs15
			_67_cf0 = _rhs7
			(_lhs8).ArraySet1(_rhs8, (_lhs9).Int())
			(_lhs10).ArraySet1(_rhs9, (_lhs11).Int())
			(_lhs12).ArraySet1(_rhs10, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs11, (_lhs15).Int())
			_33_v0 = _33_v0
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
			_ = _index11
			(_59_v13).ArraySet1(_dafny.Companion_Sequence_.Contains(_66_cf1, true), (_index11).Int())
		} else {
			var _72___mcc_h5 D0 = _source3.Get_().(D0_DC1).Cf2
			_ = _72___mcc_h5
			var _73_cf2 D0 = _72___mcc_h5
			_ = _73_cf2
			_38_v8 = Companion_Default___.Fm1(_37_v4, (_38_v8) == (false), (_58_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))).Int()).(_dafny.Int), _41_globalState)
			(_41_globalState).F0 = _dafny.IntOfInt64(273)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))
			_ = _index12
			(_58_v12).ArraySet1((_dafny.Zero).Minus((_58_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))).Int()).(_dafny.Int)), (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))
			_ = _index13
			(_59_v13).ArraySet1(_38_v8, (_index13).Int())
		}
		var _74_v20 _dafny.Map
		_ = _74_v20
		_74_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_58_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))).Int()).(_dafny.Int)).Plus((_58_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))).Int()).(_dafny.Int)), (_59_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))).Int()).(bool))
		_74_v20 = (_74_v20).Update(Companion_Default___.Fm0((_59_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_59_v13), 0))).Int()).(bool), _37_v4, _41_globalState), false)
		(_41_globalState).F6 = Companion_Default___.Fm1(_dafny.UnicodeSeqOfUtf8Bytes("xjtqkbv"), _38_v8, (_58_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_58_v12), 0))).Int()).(_dafny.Int), _41_globalState)
	}
	_35_v2 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("yyv"), _37_v4)).Cardinality())).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_37_v4).Cardinality()), _dafny.IntOfUint32((_37_v4).Cardinality())))
	var _hi0 _dafny.Int = (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if _38_v8 {
			return _dafny.SeqOf(_38_v8)
		}
		return _39_v9
	})())).Cardinality()
	_ = _hi0
	for _75_i3 := _35_v2; _75_i3.Cmp(_hi0) < 0; _75_i3 = _75_i3.Plus(_dafny.One) {
		var _76_v21 _dafny.Set
		_ = _76_v21
		_76_v21 = _dafny.SetOf(_75_i3)
		var _77_v22 _dafny.Map
		_ = _77_v22
		_77_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v21, _38_v8)
		_35_v2 = ((func() _dafny.Map {
			if _38_v8 {
				return (_77_v22).Merge(_77_v22)
			}
			return _77_v22
		})()).Cardinality()
		var _78_v23 _dafny.Map
		_ = _78_v23
		_78_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _dafny.IntOfInt64(109))
		var _79_v24 _dafny.Map
		_ = _79_v24
		_79_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v0, (_78_v23).Update(_38_v8, _35_v2))
		var _80_v25 _dafny.Array
		_ = _80_v25
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw7
		_80_v25 = _nw7
		var _81_v26 _dafny.Array
		_ = _81_v26
		var _nw8 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw8
		_81_v26 = _nw8
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
		_ = _index14
		(_81_v26).ArraySet1(_38_v8, (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
		_ = _index15
		var _rhs12 _dafny.Map = ((func() _dafny.Map {
			if _38_v8 {
				return func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate((_34_v1).Elements()); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _82_v27 _dafny.CodePoint
						_82_v27 = interface{}(_compr_6).(_dafny.CodePoint)
						if (_34_v1).Contains(_82_v27) {
							_coll6.Add(_82_v27, _78_v23)
						}
					}
					return _coll6.ToMap()
				}()
			}
			return _79_v24
		})()).Merge(_79_v24)
		_ = _rhs12
		var _rhs13 _dafny.Array = _80_v25
		_ = _rhs13
		var _rhs14 _dafny.Int = (_75_i3).Minus((_75_i3).Plus((_dafny.Zero).Minus(_75_i3)))
		_ = _rhs14
		var _rhs15 bool = _38_v8
		_ = _rhs15
		var _rhs16 bool = !(Companion_Default___.Fm1(_37_v4, _38_v8, _35_v2, _41_globalState))
		_ = _rhs16
		var _lhs16 *GlobalState = _41_globalState
		_ = _lhs16
		var _lhs17 *GlobalState = _41_globalState
		_ = _lhs17
		var _lhs18 _dafny.Array = _81_v26
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
		_ = _lhs19
		_79_v24 = _rhs12
		_80_v25 = _rhs13
		_lhs16.F16 = _rhs14
		_lhs17.F6 = _rhs15
		(_lhs18).ArraySet1(_rhs16, (_lhs19).Int())
		(_41_globalState).F0 = _75_i3
		if !((_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool)) {
			var _83_v28 bool
			_ = _83_v28
			var _out0 bool
			_ = _out0
			_out0 = Companion_Default___.M0(!(_38_v8), (_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool), !(_38_v8), Companion_Default___.Fm1(_37_v4, (_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool), _75_i3, _41_globalState), _41_globalState)
			_83_v28 = _out0
			_80_v25 = _80_v25
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
			_ = _index16
			(_81_v26).ArraySet1(_83_v28, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_80_v25), 0))
			_ = _index17
			(_80_v25).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dcuajj")).Cardinality()), (_index17).Int())
			var _84_v29 _dafny.Array
			_ = _84_v29
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw9
			_84_v29 = _nw9
			var _85_v30 _dafny.Set
			_ = _85_v30
			_85_v30 = _dafny.SetOf(_83_v28)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_80_v25), 0))
			_ = _index18
			var _rhs17 _dafny.Int = _75_i3
			_ = _rhs17
			var _rhs18 bool = !(Companion_Default___.Fm1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_38_v8, _38_v8, _35_v2, _85_v30, _41_globalState), _37_v4), (_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool), _35_v2, _41_globalState))
			_ = _rhs18
			var _rhs19 _dafny.Array = _84_v29
			_ = _rhs19
			var _lhs20 _dafny.Array = _80_v25
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_80_v25), 0))
			_ = _lhs21
			var _lhs22 *GlobalState = _41_globalState
			_ = _lhs22
			(_lhs20).ArraySet1(_rhs17, (_lhs21).Int())
			_lhs22.F6 = _rhs18
			_84_v29 = _rhs19
			var _86_v31 _dafny.Array
			_ = _86_v31
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw10
			_86_v31 = _nw10
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_86_v31), 0))
			_ = _index19
			(_86_v31).ArraySet1(_37_v4, (_index19).Int())
			var _87_v32 _dafny.MultiSet
			_ = _87_v32
			_87_v32 = _dafny.MultiSetOf(false)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_80_v25), 0))
			_ = _index20
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
			_ = _index21
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_86_v31), 0))
			_ = _index22
			var _rhs20 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_75_i3, (_80_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_80_v25), 0))).Int()).(_dafny.Int))).Plus(_75_i3))
			_ = _rhs20
			var _rhs21 bool = (_87_v32).IsDisjointFrom(_87_v32)
			_ = _rhs21
			var _rhs22 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-715))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_88_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_89_i4 _dafny.Int) _dafny.CodePoint {
					return _88_v0
				}
			})(_33_v0)))
			_ = _rhs22
			var _lhs23 _dafny.Array = _80_v25
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_80_v25), 0))
			_ = _lhs24
			var _lhs25 _dafny.Array = _81_v26
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
			_ = _lhs26
			var _lhs27 _dafny.Array = _86_v31
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_86_v31), 0))
			_ = _lhs28
			(_lhs23).ArraySet1(_rhs20, (_lhs24).Int())
			(_lhs25).ArraySet1(_rhs21, (_lhs26).Int())
			(_lhs27).ArraySet1(_rhs22, (_lhs28).Int())
		} else {
			_78_v23 = (_78_v23).Merge(_78_v23)
			var _90_v33 D0
			_ = _90_v33
			_90_v33 = Companion_D0_.Create_DC0_(_35_v2, _39_v9)
			var _91_v34 D0
			_ = _91_v34
			_91_v34 = Companion_D0_.Create_DC1_(_90_v33)
			var _92_v35 D0
			_ = _92_v35
			_92_v35 = Companion_D0_.Create_DC1_(_91_v34)
			var _93_v36 D0
			_ = _93_v36
			_93_v36 = Companion_D0_.Create_DC1_(_91_v34)
			var _94_v37 D0
			_ = _94_v37
			_94_v37 = Companion_D0_.Create_DC1_(_92_v35)
			var _95_v38 D0
			_ = _95_v38
			_95_v38 = Companion_D0_.Create_DC1_(_93_v36)
			var _96_v39 _dafny.MultiSet
			_ = _96_v39
			_96_v39 = _dafny.MultiSetOf(_95_v38, Companion_D0_.Create_DC1_(_93_v36))
			var _97_v40 _dafny.Sequence
			_ = _97_v40
			_97_v40 = _dafny.SeqOf(_95_v38, _95_v38, _95_v38)
			_38_v8 = !(((_96_v39).Union(_dafny.MultiSetFromSeq(_97_v40))).IsDisjointFrom((_96_v39).Update(Companion_D0_.Create_DC1_(_94_v37), Companion_Default___.Abs(_dafny.IntOfInt64(646)))))
			var _98_v41 _dafny.MultiSet
			_ = _98_v41
			_98_v41 = _dafny.MultiSetOf(_37_v4)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))
			_ = _index23
			(_81_v26).ArraySet1((_98_v41).IsSubsetOf((_dafny.MultiSetOf(_37_v4)).Update(_37_v4, Companion_Default___.Abs((_76_v21).Cardinality()))), (_index23).Int())
			var _rhs23 _dafny.Int = _35_v2
			_ = _rhs23
			var _rhs24 D0 = _95_v38
			_ = _rhs24
			var _lhs29 *GlobalState = _41_globalState
			_ = _lhs29
			_lhs29.F16 = _rhs23
			_95_v38 = _rhs24
			var _99_v42 _dafny.Map
			_ = _99_v42
			_99_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_35_v2), _78_v23)
			var _100_v43 _dafny.Set
			_ = _100_v43
			_100_v43 = _dafny.SetOf((_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool), (_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool))
			var _101_v44 _dafny.Map
			_ = _101_v44
			_101_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool), _99_v42)
			var _102_v46 _dafny.Map
			_ = _102_v46
			_102_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_39_v9, func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(542), _dafny.IntOfInt64(652))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _103_v45 _dafny.Int
					_103_v45 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(542)).Cmp(_103_v45) <= 0) && ((_103_v45).Cmp(_dafny.IntOfInt64(652)) < 0) {
						_coll7.Add((_103_v45).Times(_75_i3), _78_v23)
					}
				}
				return _coll7.ToMap()
			}())
			var _rhs25 bool = !(_100_v43).Equals(_100_v43)
			_ = _rhs25
			var _rhs26 _dafny.Map = ((_99_v42).Merge((func() _dafny.Map {
				if (_101_v44).Contains((_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool)) {
					return (_101_v44).Get((_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool)).(_dafny.Map)
				}
				return (func() _dafny.Map {
					if (_102_v46).Contains(_dafny.Companion_Sequence_.Update(_39_v9, (Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_39_v9).Cardinality()))).Uint32(), (_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool))) {
						return (_102_v46).Get(_dafny.Companion_Sequence_.Update(_39_v9, (Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_39_v9).Cardinality()))).Uint32(), (_81_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_81_v26), 0))).Int()).(bool))).(_dafny.Map)
					}
					return _99_v42
				})()
			})())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v2, _78_v23))
			_ = _rhs26
			var _lhs30 *GlobalState = _41_globalState
			_ = _lhs30
			_lhs30.F7 = _rhs25
			_99_v42 = _rhs26
		}
	}
	_38_v8 = (_dafny.IntOfInt64(979)).Cmp(_dafny.IntOfInt64(283)) < 0
	var _104_v47 _dafny.Array
	_ = _104_v47
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
	_ = _nw11
	_104_v47 = _nw11
	var _105_v48 D0
	_ = _105_v48
	_105_v48 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(636), Companion_Default___.Fm3(_dafny.IntOfInt64(569), _dafny.IntOfUint32((_36_v3).Cardinality()), _38_v8, _41_globalState))
	var _106_v49 D0
	_ = _106_v49
	_106_v49 = Companion_D0_.Create_DC1_(_105_v48)
	var _107_v50 D0
	_ = _107_v50
	_107_v50 = Companion_D0_.Create_DC1_(_106_v49)
	var _108_v51 _dafny.Map
	_ = _108_v51
	_108_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v0, _35_v2)
	var _109_v52 _dafny.MultiSet
	_ = _109_v52
	_109_v52 = _dafny.MultiSetOf(_104_v47)
	var _rhs27 _dafny.Array = _104_v47
	_ = _rhs27
	var _rhs28 D0 = Companion_Default___.Fm4(_108_v51, _41_globalState)
	_ = _rhs28
	var _rhs29 bool = _38_v8
	_ = _rhs29
	var _rhs30 bool = (_dafny.MultiSetOf(_104_v47, _104_v47, _104_v47)).IsDisjointFrom((_109_v52).Intersection(_109_v52))
	_ = _rhs30
	var _lhs31 *GlobalState = _41_globalState
	_ = _lhs31
	_104_v47 = _rhs27
	_107_v50 = _rhs28
	_lhs31.F7 = _rhs29
	_38_v8 = _rhs30
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_104_v47), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _110_i5 _dafny.Int
		_110_i5 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_110_i5).Sign() != -1) && ((_110_i5).Cmp(_dafny.ArrayLen((_104_v47), 0)) < 0)) {
			(_104_v47).ArraySet1(_38_v8, (_110_i5).Int())
		}
	}
	var _111_v53 _dafny.Array
	_ = _111_v53
	var _nwElement0_1 _dafny.Int = _35_v2
	_ = _nwElement0_1
	var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
	_ = _nw12
	(_nw12).ArraySet1(_nwElement0_1, 0)
	(_nw12).ArraySet1(_35_v2, 1)
	(_nw12).ArraySet1((_dafny.IntOfInt64(721)).Times(_35_v2), 2)
	(_nw12).ArraySet1(_dafny.One, 3)
	(_nw12).ArraySet1((_dafny.Zero).Minus(_35_v2), 4)
	(_nw12).ArraySet1(_35_v2, 5)
	(_nw12).ArraySet1(Companion_Default___.Fm0(_38_v8, _37_v4, _41_globalState), 6)
	(_nw12).ArraySet1(_35_v2, 7)
	_111_v53 = _nw12
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _index24
	(_111_v53).ArraySet1((_35_v2).Times(_35_v2), (_index24).Int())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_111_v53), 0))
	_ = _index25
	(_111_v53).ArraySet1(_dafny.IntOfUint32((_37_v4).Cardinality()), (_index25).Int())
	var _112_v54 _dafny.MultiSet
	_ = _112_v54
	_112_v54 = _dafny.MultiSetOf(_38_v8)
	var _113_v55 _dafny.Map
	_ = _113_v55
	_113_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v54, _104_v47)
	var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _index26
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_111_v53), 0))
	_ = _index27
	var _rhs31 _dafny.Int = Companion_Default___.Fm0(_38_v8, _37_v4, _41_globalState)
	_ = _rhs31
	var _rhs32 _dafny.CodePoint = (Companion_Default___.Fm5(_38_v8, _33_v0, _35_v2, _38_v8, _41_globalState)).Dtor_cf3()
	_ = _rhs32
	var _rhs33 _dafny.Int = (_dafny.Zero).Minus(((_113_v55).Cardinality()).Minus(_dafny.IntOfInt64(-525)))
	_ = _rhs33
	var _rhs34 bool = !(_38_v8)
	_ = _rhs34
	var _lhs32 _dafny.Array = _111_v53
	_ = _lhs32
	var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _lhs33
	var _lhs34 _dafny.Array = _111_v53
	_ = _lhs34
	var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_111_v53), 0))
	_ = _lhs35
	var _lhs36 *GlobalState = _41_globalState
	_ = _lhs36
	(_lhs32).ArraySet1(_rhs31, (_lhs33).Int())
	_33_v0 = _rhs32
	(_lhs34).ArraySet1(_rhs33, (_lhs35).Int())
	_lhs36.F6 = _rhs34
	if (_112_v54).Contains((_38_v8) && (_38_v8)) {
		var _114_v56 *C1
		_ = _114_v56
		var _nw13 *C1 = New_C1_()
		_ = _nw13
		_nw13.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (Companion_Default___.Fm9((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _38_v8, false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(282))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_115_v53 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_116_i6 _dafny.Int) _dafny.Int {
				return (_115_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_115_v53), 0))).Int()).(_dafny.Int)
			}
		})(_111_v53)))).Cardinality()), _41_globalState)).Cardinality()), _107_v50)
		_114_v56 = _nw13
		var _117_v57 D2
		_ = _117_v57
		_117_v57 = Companion_D2_.Create_DC8_(_38_v8)
		var _118_v58 _dafny.Map
		_ = _118_v58
		_118_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v57, Companion_Default___.Fm1(_dafny.UnicodeSeqOfUtf8Bytes("uteyy"), _38_v8, (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _41_globalState))
		var _119_v59 _dafny.MultiSet
		_ = _119_v59
		_119_v59 = _dafny.MultiSetOf(_35_v2)
		(_41_globalState).F6 = !((Companion_Default___.Fm10(_118_v58, _38_v8, _119_v59, _41_globalState)).IsDisjointFrom(_dafny.SetOf(_51_v11)))
		if (true) && (_38_v8) {
			(_41_globalState).F7 = ((_112_v54).Cardinality()).Cmp((_35_v2).Plus(_35_v2)) <= 0
			var _rhs35 bool = ((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_37_v4, (Companion_Default___.SafeIndex((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32(), _33_v0)).Cardinality())) == 0
			_ = _rhs35
			var _rhs36 _dafny.Int = (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)
			_ = _rhs36
			_38_v8 = _rhs35
			_35_v2 = _rhs36
			_37_v4 = _dafny.UnicodeSeqOfUtf8Bytes("vsbt")
			_38_v8 = false
			var _120_v60 *C0
			_ = _120_v60
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__()
			_120_v60 = _nw14
		} else {
			var _121_v61 _dafny.Array
			_ = _121_v61
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
			_ = _nw15
			_121_v61 = _nw15
			var _rhs37 bool = !(_38_v8)
			_ = _rhs37
			var _rhs38 _dafny.Int = Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if _38_v8 {
					return _35_v2
				}
				return _35_v2
			})(), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
			_ = _rhs38
			var _rhs39 _dafny.Int = (_114_v56).Fm6((_37_v4).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _38_v8, !(_38_v8), _38_v8, _41_globalState)
			_ = _rhs39
			var _rhs40 _dafny.Array = _121_v61
			_ = _rhs40
			var _lhs37 *GlobalState = _41_globalState
			_ = _lhs37
			var _lhs38 *GlobalState = _41_globalState
			_ = _lhs38
			var _lhs39 *GlobalState = _41_globalState
			_ = _lhs39
			_lhs37.F6 = _rhs37
			_lhs38.F16 = _rhs38
			_lhs39.F16 = _rhs39
			_121_v61 = _rhs40
			var _122_v62 _dafny.Sequence
			_ = _122_v62
			_122_v62 = _dafny.SeqOf(_37_v4)
			(_41_globalState).F6 = _dafny.Companion_Sequence_.Contains(_122_v62, _dafny.UnicodeSeqOfUtf8Bytes("loukb"))
			var _123_v63 _dafny.Set
			_ = _123_v63
			_123_v63 = _dafny.SetOf(_dafny.IntOfInt64(811), (_dafny.SetOf(_38_v8, _38_v8)).Cardinality(), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_114_v56.F19).Contains(_38_v8) {
					return (_114_v56.F19).Get(_38_v8).(_dafny.Int)
				}
				return _35_v2
			})(), _35_v2)
			var _124_v64 _dafny.Array
			_ = _124_v64
			var _125_v65 _dafny.Map
			_ = _125_v65
			var _out1 _dafny.Array
			_ = _out1
			var _out2 _dafny.Map
			_ = _out2
			_out1, _out2 = (_114_v56).M1(_dafny.IntOfInt64(2), (_dafny.SetOf(_35_v2)).Intersection(_123_v63), _35_v2, _41_globalState)
			_124_v64 = _out1
			_125_v65 = _out2
			var _126_v66 _dafny.Sequence
			_ = _126_v66
			_126_v66 = _dafny.SeqOf(_123_v63, _123_v63)
			var _127_v68 _dafny.Array
			_ = _127_v68
			var _128_v69 _dafny.Map
			_ = _128_v69
			var _out3 _dafny.Array
			_ = _out3
			var _out4 _dafny.Map
			_ = _out4
			_out3, _out4 = (_114_v56).M1(Companion_Default___.SafeModuloInt(_35_v2, _dafny.IntOfInt64(789)), ((_126_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))).Cardinality()), _dafny.IntOfUint32((_126_v66).Cardinality()))).Uint32()).(_dafny.Set)).Union(_dafny.SetOf(_dafny.IntOfInt64(-395), (func() _dafny.Set {
				var _coll8 = _dafny.NewBuilder()
				_ = _coll8
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(913), _dafny.IntOfInt64(-188))); ; {
					_compr_8, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _129_v67 _dafny.Int
					_129_v67 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(913)).Cmp(_129_v67) <= 0) && ((_129_v67).Cmp(_dafny.IntOfInt64(-188)) < 0) {
						_coll8.Add((_129_v67).Plus((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll8.ToSet()
			}()).Cardinality())), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _41_globalState)
			_127_v68 = _out3
			_128_v69 = _out4
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_127_v68), 0))
			_ = _index28
			(_127_v68).ArraySet1((_38_v8) == (_38_v8), (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_127_v68), 0))
			_ = _index29
			(_127_v68).ArraySet1(_38_v8, (_index29).Int())
		}
		(_41_globalState).F16 = (_dafny.Zero).Minus((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
		var _130_v70 *C1
		_ = _130_v70
		var _nw16 *C1 = New_C1_()
		_ = _nw16
		_nw16.Ctor__(_114_v56.F19, Companion_D0_.Create_DC1_(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(352))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_131_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_132_i7 _dafny.Int) _dafny.CodePoint {
				return _131_v0
			}
		})(_33_v0)))).Cardinality()), _39_v9)))
		_130_v70 = _nw16
		_130_v70 = _130_v70
	} else {
		(_41_globalState).F7 = _38_v8
		var _133_v71 D2
		_ = _133_v71
		_133_v71 = Companion_D2_.Create_DC7_(_33_v0)
		var _pat_let_tv0 = _33_v0
		_ = _pat_let_tv0
		_133_v71 = func(_pat_let0_0 D2) D2 {
			return func(_134_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let1_0 _dafny.CodePoint) D2 {
					return func(_135_dt__update_hcf12_h0 _dafny.CodePoint) D2 {
						return Companion_D2_.Create_DC7_(_135_dt__update_hcf12_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(Companion_D2_.Create_DC7_(_dafny.CodePoint('q')))
		var _136_v72 _dafny.Array
		_ = _136_v72
		var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
		_ = _nw17
		_136_v72 = _nw17
		var _137_v73 _dafny.Sequence
		_ = _137_v73
		_137_v73 = _dafny.SeqOf(_136_v72, _136_v72)
		_136_v72 = (_137_v73).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_137_v73).Cardinality()))).Uint32()).(_dafny.Array)
		var _source4 D2 = Companion_D2_.Create_DC7_(_33_v0)
		_ = _source4
		if _source4.Is_DC7() {
			var _138___mcc_h6 _dafny.CodePoint = _source4.Get_().(D2_DC7).Cf12
			_ = _138___mcc_h6
			var _139_cf12 _dafny.CodePoint = _138___mcc_h6
			_ = _139_cf12
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
			_ = _index30
			(_111_v53).ArraySet1((_dafny.Zero).Minus((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)), (_index30).Int())
			var _140_v74 _dafny.Map
			_ = _140_v74
			_140_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v2, _36_v3)
			var _141_v75 _dafny.Map
			_ = _141_v75
			_141_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Concatenate(_36_v3, (func() _dafny.Sequence {
				if (_140_v74).Contains((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)) {
					return (_140_v74).Get((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _36_v3
			})()))
			_141_v75 = (_141_v75).Update(_35_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(999))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_142_v53 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_143_i8 _dafny.Int) _dafny.Int {
					return (_142_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_142_v53), 0))).Int()).(_dafny.Int)
				}
			})(_111_v53))))
			var _144_v76 *C0
			_ = _144_v76
			var _nw18 *C0 = New_C0_()
			_ = _nw18
			_nw18.Ctor__()
			_144_v76 = _nw18
			var _145_v77 _dafny.MultiSet
			_ = _145_v77
			_145_v77 = _dafny.MultiSetOf(_144_v76)
			var _146_v78 bool
			_ = _146_v78
			var _out5 bool
			_ = _out5
			_out5 = Companion_Default___.M0(_38_v8, (_35_v2).Cmp((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)) != 0, (_145_v77).IsDisjointFrom(_145_v77), _38_v8, _41_globalState)
			_146_v78 = _out5
			var _147_v79 _dafny.Map
			_ = _147_v79
			_147_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v78, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(8))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_148_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_149_i9 _dafny.Int) _dafny.CodePoint {
					return _148_v0
				}
			})(_33_v0)))).Cardinality()))
			var _150_v80 *C1
			_ = _150_v80
			var _nw19 *C1 = New_C1_()
			_ = _nw19
			_nw19.Ctor__(_147_v79, _107_v50)
			_150_v80 = _nw19
			var _151_v81 _dafny.MultiSet
			_ = _151_v81
			_151_v81 = _dafny.MultiSetOf(_150_v80, _150_v80, _150_v80)
			var _152_v82 _dafny.Set
			_ = _152_v82
			_152_v82 = _dafny.SetOf((_151_v81).Cardinality())
			var _153_v83 _dafny.Map
			_ = _153_v83
			_153_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if _146_v78 {
					return _38_v8
				}
				return !(_146_v78)
			})(), _152_v82)
			_153_v83 = _153_v83
		} else if _source4.Is_DC8() {
			var _154___mcc_h7 bool = _source4.Get_().(D2_DC8).Cf13
			_ = _154___mcc_h7
			var _155_cf13 bool = _154___mcc_h7
			_ = _155_cf13
			var _156_v84 _dafny.Map
			_ = _156_v84
			_156_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _35_v2)
			var _157_v85 *C1
			_ = _157_v85
			var _nw20 *C1 = New_C1_()
			_ = _nw20
			_nw20.Ctor__(_156_v84, Companion_D0_.Create_DC1_(Companion_D0_.Create_DC1_(_105_v48)))
			_157_v85 = _nw20
			var _158_v86 _dafny.Sequence
			_ = _158_v86
			_158_v86 = _dafny.SeqOf(_157_v85)
			var _159_v87 bool
			_ = _159_v87
			var _out6 bool
			_ = _out6
			_out6 = Companion_Default___.M0(_38_v8, _38_v8, _155_cf13, Companion_Default___.Fm1(_37_v4, _38_v8, _dafny.IntOfUint32(((Companion_D4_.Create_DC11_(_158_v86)).Dtor_cf16()).Cardinality()), _41_globalState), _41_globalState)
			_159_v87 = _out6
			(_41_globalState).F16 = (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)
			_33_v0 = _33_v0
			var _160_v88 *C0
			_ = _160_v88
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__()
			_160_v88 = _nw21
		} else {
			var _161___mcc_h8 _dafny.Set = _source4.Get_().(D2_DC6).Cf11
			_ = _161___mcc_h8
			var _162_cf11 _dafny.Set = _161___mcc_h8
			_ = _162_cf11
			var _163_v89 D1
			_ = _163_v89
			_163_v89 = Companion_D1_.Create_DC3_(_35_v2)
			_163_v89 = Companion_D1_.Create_DC3_((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
			_38_v8 = (_38_v8) || (_38_v8)
			var _164_v90 _dafny.Sequence
			_ = _164_v90
			_164_v90 = _dafny.SeqOf(_111_v53, _111_v53)
			var _165_v91 _dafny.Map
			_ = _165_v91
			_165_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _dafny.IntOfInt64(-63))
			var _166_v92 *C1
			_ = _166_v92
			var _nw22 *C1 = New_C1_()
			_ = _nw22
			_nw22.Ctor__(_165_v91, _107_v50)
			_166_v92 = _nw22
			var _167_v93 _dafny.Sequence
			_ = _167_v93
			_167_v93 = _dafny.SeqOf(_111_v53, _111_v53, (_164_v90).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_164_v90).Cardinality()))).Uint32()).(_dafny.Array), (Companion_D5_.Create_DC17_(_112_v54, _111_v53, _166_v92, _dafny.IntOfUint32((_37_v4).Cardinality()), _dafny.Companion_Sequence_.Update(_37_v4, (Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32(), _33_v0))).Dtor_cf25(), _111_v53)
			var _rhs41 _dafny.Int = _35_v2
			_ = _rhs41
			var _rhs42 _dafny.CodePoint = Companion_Default___.Fm11(_41_globalState)
			_ = _rhs42
			var _rhs43 _dafny.Array = (_167_v93).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_167_v93).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs43
			var _lhs40 *GlobalState = _41_globalState
			_ = _lhs40
			_lhs40.F0 = _rhs41
			_33_v0 = _rhs42
			_111_v53 = _rhs43
			var _168_v94 _dafny.Map
			_ = _168_v94
			_168_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, Companion_D0_.Create_DC1_(_105_v48))
			_168_v94 = (_168_v94).Update(_38_v8, (func() D0 {
				if _38_v8 {
					return _107_v50
				}
				return _107_v50
			})())
		}
		var _169_v95 bool
		_ = _169_v95
		var _out7 bool
		_ = _out7
		_out7 = Companion_Default___.M0(_38_v8, _38_v8, _38_v8, (_38_v8) || (!(_38_v8)), _41_globalState)
		_169_v95 = _out7
	}
	var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
	_ = _index31
	(_104_v47).ArraySet1(_38_v8, (_index31).Int())
	var _170_v96 _dafny.Map
	_ = _170_v96
	_170_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _dafny.IntOfInt64(676))
	var _171_v97 *C1
	_ = _171_v97
	var _nw23 *C1 = New_C1_()
	_ = _nw23
	_nw23.Ctor__(_170_v96, Companion_D0_.Create_DC1_(_105_v48))
	_171_v97 = _nw23
	var _172_v98 _dafny.Sequence
	_ = _172_v98
	_172_v98 = _dafny.SeqOf(_171_v97)
	var _173_v99 D4
	_ = _173_v99
	_173_v99 = Companion_D4_.Create_DC11_(_172_v98)
	var _174_v100 D4
	_ = _174_v100
	_174_v100 = Companion_D4_.Create_DC14_(_173_v99)
	var _pat_let_tv1 = _173_v99
	_ = _pat_let_tv1
	var _175_v101 _dafny.Map
	_ = _175_v101
	_175_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func(_pat_let2_0 D4) D4 {
		return func(_176_dt__update__tmp_h1 D4) D4 {
			return func(_pat_let3_0 D4) D4 {
				return func(_177_dt__update_hcf21_h0 D4) D4 {
					return Companion_D4_.Create_DC14_(_177_dt__update_hcf21_h0)
				}(_pat_let3_0)
			}(_pat_let_tv1)
		}(_pat_let2_0)
	}(_174_v100))
	var _178_v102 _dafny.Sequence
	_ = _178_v102
	_178_v102 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, Companion_D4_.Create_DC14_(_173_v99)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _174_v100), (_175_v101).Update(true, Companion_D4_.Create_DC14_(Companion_D4_.Create_DC12_())))
	var _179_v103 _dafny.MultiSet
	_ = _179_v103
	_179_v103 = _dafny.MultiSetOf((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
	var _180_v104 _dafny.Map
	_ = _180_v104
	_180_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _179_v103)
	var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
	_ = _index32
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _index33
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _index34
	var _rhs44 bool = !(!((_179_v103).IsProperSubsetOf((func() _dafny.MultiSet {
		if (_180_v104).Contains(_38_v8) {
			return (_180_v104).Get(_38_v8).(_dafny.MultiSet)
		}
		return Companion_Default___.Fm12(_dafny.IntOfInt64(238), _41_globalState)
	})()))) || (_38_v8)
	_ = _rhs44
	var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_178_v102, _178_v102)
	_ = _rhs45
	var _rhs46 bool = _38_v8
	_ = _rhs46
	var _rhs47 _dafny.Int = Companion_Default___.Fm0((func() bool {
		if _38_v8 {
			return _38_v8
		}
		return _38_v8
	})(), _dafny.UnicodeSeqOfUtf8Bytes("hrcu"), _41_globalState)
	_ = _rhs47
	var _rhs48 _dafny.Int = ((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)).Plus(Companion_Default___.SafeModuloInt((_171_v97).Fm6(_33_v0, _38_v8, _38_v8, _38_v8, _41_globalState), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)))
	_ = _rhs48
	var _lhs41 _dafny.Array = _104_v47
	_ = _lhs41
	var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
	_ = _lhs42
	var _lhs43 *GlobalState = _41_globalState
	_ = _lhs43
	var _lhs44 _dafny.Array = _111_v53
	_ = _lhs44
	var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _lhs45
	var _lhs46 _dafny.Array = _111_v53
	_ = _lhs46
	var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _lhs47
	(_lhs41).ArraySet1(_rhs44, (_lhs42).Int())
	_178_v102 = _rhs45
	_lhs43.F6 = _rhs46
	(_lhs44).ArraySet1(_rhs47, (_lhs45).Int())
	(_lhs46).ArraySet1(_rhs48, (_lhs47).Int())
	var _181_v105 _dafny.Sequence
	_ = _181_v105
	_181_v105 = _dafny.SeqOf(_104_v47)
	var _182_v106 _dafny.Set
	_ = _182_v106
	_182_v106 = _dafny.SetOf(true, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))
	var _rhs49 _dafny.Int = Companion_Default___.SafeDivisionInt(_35_v2, (_35_v2).Times(_35_v2))
	_ = _rhs49
	var _rhs50 bool = _dafny.Companion_Sequence_.IsPrefixOf(_36_v3, _36_v3)
	_ = _rhs50
	var _rhs51 bool = (_35_v2).Cmp(Companion_Default___.SafeModuloInt((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
		if (_171_v97.F19).Contains(!((_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))) {
			return (_171_v97.F19).Get(!((_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))).(_dafny.Int)
		}
		return (_170_v96).Cardinality()
	})())) <= 0
	_ = _rhs51
	var _rhs52 _dafny.Sequence = Companion_Default___.Fm2(_dafny.Companion_Sequence_.IsProperPrefixOf(_181_v105, _181_v105), !(false), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), (_182_v106).Union(_182_v106), _41_globalState)
	_ = _rhs52
	var _rhs53 bool = (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool)
	_ = _rhs53
	var _lhs48 *GlobalState = _41_globalState
	_ = _lhs48
	var _lhs49 *GlobalState = _41_globalState
	_ = _lhs49
	var _lhs50 *GlobalState = _41_globalState
	_ = _lhs50
	_lhs48.F0 = _rhs49
	_38_v8 = _rhs50
	_lhs49.F7 = _rhs51
	_37_v4 = _rhs52
	_lhs50.F6 = _rhs53
	var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
	_ = _index35
	(_111_v53).ArraySet1((func() _dafny.Int {
		if (true) && (_38_v8) {
			return ((_182_v106).Cardinality()).Minus((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
		}
		return _35_v2
	})(), (_index35).Int())
	var _183_v107 _dafny.Map
	_ = _183_v107
	_183_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
	var _184_v108 _dafny.Sequence
	_ = _184_v108
	_184_v108 = _dafny.SeqOf(_111_v53)
	var _185_v109 *C0
	_ = _185_v109
	var _nw24 *C0 = New_C0_()
	_ = _nw24
	_nw24.Ctor__()
	_185_v109 = _nw24
	var _186_v110 _dafny.Map
	_ = _186_v110
	_186_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_185_v109, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))
	var _187_v111 _dafny.Map
	_ = _187_v111
	_187_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v110, _38_v8)
	var _188_v112 D5
	_ = _188_v112
	_188_v112 = Companion_D5_.Create_DC17_(_dafny.MultiSetOf(_38_v8), _111_v53, _171_v97, _dafny.IntOfUint32((_39_v9).Cardinality()), _37_v4)
	var _189_v113 _dafny.Array
	_ = _189_v113
	var _nwElement0_2 _dafny.Array = _111_v53
	_ = _nwElement0_2
	var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(17))
	_ = _nw25
	(_nw25).ArraySet1(_nwElement0_2, 0)
	(_nw25).ArraySet1(_111_v53, 1)
	(_nw25).ArraySet1(_111_v53, 2)
	(_nw25).ArraySet1((func() _dafny.Array {
		if (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool) {
			return (_184_v108).Select((Companion_Default___.SafeIndex((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_184_v108).Cardinality()))).Uint32()).(_dafny.Array)
		}
		return _111_v53
	})(), 3)
	(_nw25).ArraySet1(_111_v53, 4)
	(_nw25).ArraySet1(_111_v53, 5)
	(_nw25).ArraySet1(_111_v53, 6)
	(_nw25).ArraySet1(_111_v53, 7)
	(_nw25).ArraySet1(_111_v53, 8)
	(_nw25).ArraySet1(_111_v53, 9)
	(_nw25).ArraySet1(_111_v53, 10)
	(_nw25).ArraySet1(_111_v53, 11)
	(_nw25).ArraySet1((_184_v108).Select((Companion_Default___.SafeIndex((_187_v111).Cardinality(), _dafny.IntOfUint32((_184_v108).Cardinality()))).Uint32()).(_dafny.Array), 12)
	(_nw25).ArraySet1((_188_v112).Dtor_cf25(), 13)
	(_nw25).ArraySet1(_111_v53, 14)
	(_nw25).ArraySet1(_111_v53, 15)
	(_nw25).ArraySet1(_111_v53, 16)
	_189_v113 = _nw25
	var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))
	_ = _index36
	(_189_v113).ArraySet1(_111_v53, (_index36).Int())
	var _190_v114 _dafny.Map
	_ = _190_v114
	_190_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v47, _35_v2)
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))
	_ = _index37
	var _rhs54 bool = ((_35_v2).Minus(_35_v2)).Cmp((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)) != 0
	_ = _rhs54
	var _rhs55 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_37_v4).Cardinality()), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
	_ = _rhs55
	var _rhs56 _dafny.Array = _111_v53
	_ = _rhs56
	var _rhs57 _dafny.Map = (_190_v114).Update(_104_v47, _35_v2)
	_ = _rhs57
	var _rhs58 bool = _dafny.Companion_Sequence_.Contains(_37_v4, _33_v0)
	_ = _rhs58
	var _lhs51 *GlobalState = _41_globalState
	_ = _lhs51
	var _lhs52 _dafny.Array = _189_v113
	_ = _lhs52
	var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))
	_ = _lhs53
	var _lhs54 *GlobalState = _41_globalState
	_ = _lhs54
	_lhs51.F7 = _rhs54
	_183_v107 = _rhs55
	(_lhs52).ArraySet1(_rhs56, (_lhs53).Int())
	_190_v114 = _rhs57
	_lhs54.F6 = _rhs58
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_dafny.ArrayCastTo((_189_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))).Int()))), 0))); ; {
		_guard_loop_1, _ok10 := _iter10()
		if !_ok10 {
			break
		}
		var _191_i10 _dafny.Int
		_191_i10 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_191_i10).Sign() != -1) && ((_191_i10).Cmp(_dafny.ArrayLen((_dafny.ArrayCastTo((_189_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))).Int()))), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_dafny.ArrayCastTo((_189_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))).Int())), (_191_i10).Int(), (_191_i10).Minus((_dafny.Zero).Minus(_35_v2))))
		}
	}
	for _iter11 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok11 := _iter11()
		if !_ok11 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))
	_ = _index38
	var _rhs59 _dafny.Array = _dafny.ArrayCastTo((_189_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))).Int()))
	_ = _rhs59
	var _rhs60 bool = _38_v8
	_ = _rhs60
	var _lhs55 _dafny.Array = _189_v113
	_ = _lhs55
	var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))
	_ = _lhs56
	(_lhs55).ArraySet1(_rhs59, (_lhs56).Int())
	_38_v8 = _rhs60
	if (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool) {
		var _192_v115 _dafny.Map
		_ = _192_v115
		_192_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_37_v4, (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
		var _193_v116 _dafny.Sequence
		_ = _193_v116
		_193_v116 = _dafny.SeqOf(_37_v4, _dafny.Companion_Sequence_.Update(_37_v4, (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_192_v115).Contains(_37_v4) {
				return (_192_v115).Get(_37_v4).(_dafny.Int)
			}
			return _35_v2
		})(), _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32(), _33_v0))
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
		_ = _index39
		var _rhs61 _dafny.Sequence = _193_v116
		_ = _rhs61
		var _rhs62 _dafny.Sequence = _36_v3
		_ = _rhs62
		var _rhs63 _dafny.Int = (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)
		_ = _rhs63
		var _lhs57 _dafny.Array = _111_v53
		_ = _lhs57
		var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
		_ = _lhs58
		_193_v116 = _rhs61
		_36_v3 = _rhs62
		(_lhs57).ArraySet1(_rhs63, (_lhs58).Int())
		var _194_v117 D2
		_ = _194_v117
		_194_v117 = Companion_D2_.Create_DC7_(_33_v0)
		var _source5 D2 = _194_v117
		_ = _source5
		if _source5.Is_DC7() {
			var _195___mcc_h9 _dafny.CodePoint = _source5.Get_().(D2_DC7).Cf12
			_ = _195___mcc_h9
			var _196_cf12 _dafny.CodePoint = _195___mcc_h9
			_ = _196_cf12
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
			_ = _index40
			(_104_v47).ArraySet1((_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), (_index40).Int())
			(_41_globalState).F0 = _35_v2
			var _197_v118 _dafny.Array
			_ = _197_v118
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_1
			var _nw26 _dafny.Array
			_ = _nw26
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw26 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = (func(_198_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_199_i11 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_198_v4, _dafny.UnicodeSeqOfUtf8Bytes("ymqrueuc"))
					}
				})(_37_v4)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw26 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw26).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw26).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_197_v118 = _nw26
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_197_v118), 0))
			_ = _index41
			(_197_v118).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dgxpoeh"), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_197_v118), 0))
			_ = _index42
			(_197_v118).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-478))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_200_cf12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_201_i12 _dafny.Int) _dafny.CodePoint {
					return _200_cf12
				}
			})(_196_cf12))), (_index42).Int())
			var _rhs64 *C0 = _185_v109
			_ = _rhs64
			var _rhs65 _dafny.Int = (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)
			_ = _rhs65
			var _lhs59 *GlobalState = _41_globalState
			_ = _lhs59
			_185_v109 = _rhs64
			_lhs59.F0 = _rhs65
		} else if _source5.Is_DC8() {
			var _202___mcc_h10 bool = _source5.Get_().(D2_DC8).Cf13
			_ = _202___mcc_h10
			var _203_cf13 bool = _202___mcc_h10
			_ = _203_cf13
			var _204_v119 _dafny.Array
			_ = _204_v119
			var _nwElement0_3 _dafny.Sequence = _37_v4
			_ = _nwElement0_3
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(10))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_3, 0)
			(_nw27).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_205_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_206_i13 _dafny.Int) _dafny.CodePoint {
					return _205_v0
				}
			})(_33_v0))), 1)
			(_nw27).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("levgfjrvj"), 2)
			(_nw27).ArraySet1(_37_v4, 3)
			(_nw27).ArraySet1(_37_v4, 4)
			(_nw27).ArraySet1(_dafny.Companion_Sequence_.Update(_37_v4, (Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_37_v4).Cardinality()))).Uint32(), _dafny.CodePoint('v')), 5)
			(_nw27).ArraySet1(_37_v4, 6)
			(_nw27).ArraySet1(_37_v4, 7)
			(_nw27).ArraySet1(_37_v4, 8)
			(_nw27).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_193_v116).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_193_v116).Cardinality()))).Uint32()).(_dafny.Sequence), _37_v4), 9)
			_204_v119 = _nw27
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_204_v119), 0))
			_ = _index43
			(_204_v119).ArraySet1(_37_v4, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_204_v119), 0))
			_ = _index44
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
			_ = _index45
			var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_37_v4, _37_v4)
			_ = _rhs66
			var _rhs67 *C0 = _185_v109
			_ = _rhs67
			var _rhs68 _dafny.Int = ((_dafny.Zero).Minus(((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)).Plus((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)))).Times(_dafny.IntOfUint32((_39_v9).Cardinality()))
			_ = _rhs68
			var _lhs60 _dafny.Array = _204_v119
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_204_v119), 0))
			_ = _lhs61
			var _lhs62 _dafny.Array = _111_v53
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))
			_ = _lhs63
			(_lhs60).ArraySet1(_rhs66, (_lhs61).Int())
			_185_v109 = _rhs67
			(_lhs62).ArraySet1(_rhs68, (_lhs63).Int())
			var _207_v120 *C1
			_ = _207_v120
			var _nw28 *C1 = New_C1_()
			_ = _nw28
			_nw28.Ctor__((_171_v97.F19).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(916))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_208_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_209_i14 _dafny.Int) _dafny.CodePoint {
					return _208_v0
				}
			})(_33_v0)))).Cardinality()))), (_171_v97).F20())
			_207_v120 = _nw28
			var _210_v121 *C1
			_ = _210_v121
			var _nw29 *C1 = New_C1_()
			_ = _nw29
			_nw29.Ctor__(_207_v120.F19, (_171_v97).F20())
			_210_v121 = _nw29
			var _211_v122 *C1
			_ = _211_v122
			var _nw30 *C1 = New_C1_()
			_ = _nw30
			_nw30.Ctor__(_170_v96, (_171_v97).F20())
			_211_v122 = _nw30
		} else {
			var _212___mcc_h11 _dafny.Set = _source5.Get_().(D2_DC6).Cf11
			_ = _212___mcc_h11
			var _213_cf11 _dafny.Set = _212___mcc_h11
			_ = _213_cf11
			_213_cf11 = _213_cf11
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
			_ = _index46
			(_104_v47).ArraySet1(Companion_Default___.Fm1(_dafny.Companion_Sequence_.Concatenate(_37_v4, _37_v4), (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _41_globalState), (_index46).Int())
			var _214_v123 _dafny.Map
			_ = _214_v123
			_214_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), _38_v8)
			var _215_v124 _dafny.Map
			_ = _215_v124
			_215_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_214_v123, (_183_v107).Cardinality())
			var _216_v125 bool
			_ = _216_v125
			var _out8 bool
			_ = _out8
			_out8 = Companion_Default___.M0(_38_v8, Companion_Default___.Fm1(_37_v4, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), (_213_cf11).Cardinality(), _41_globalState), !(_dafny.SetOf((func() _dafny.Int {
				if (_215_v124).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))) {
					return (_215_v124).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))).(_dafny.Int)
				}
				return _35_v2
			})(), (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))).Contains(_dafny.IntOfUint32((_36_v3).Cardinality())), (_39_v9).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_39_v9).Cardinality()))).Uint32()).(bool), _41_globalState)
			_216_v125 = _out8
			(_41_globalState).F7 = (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool)
		}
		var _217_v126 _dafny.Map
		_ = _217_v126
		_217_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v2, _182_v106)
		var _rhs69 _dafny.Set = ((func() _dafny.Set {
			if (_217_v126).Contains((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)) {
				return (_217_v126).Get((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)).(_dafny.Set)
			}
			return _dafny.SetOf(_38_v8, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), _38_v8, _38_v8)
		})()).Intersection((_182_v106).Intersection(_182_v106))
		_ = _rhs69
		var _rhs70 _dafny.Int = ((_dafny.Zero).Minus(_35_v2)).Minus(_35_v2)
		_ = _rhs70
		_182_v106 = _rhs69
		_35_v2 = _rhs70
		var _218_v127 D5
		_ = _218_v127
		_218_v127 = Companion_D5_.Create_DC15_(_111_v53)
		_218_v127 = _218_v127
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(352), _dafny.ArrayLen((_189_v113), 0))
		_ = _index47
		(_189_v113).ArraySet1(_111_v53, (_index47).Int())
	} else {
		(_41_globalState).F7 = _38_v8
		var _219_v128 _dafny.Array
		_ = _219_v128
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_2
		var _nw31 _dafny.Array
		_ = _nw31
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw31 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Map = (func(_220_v47 _dafny.Array, _221_v8 bool, _222_v106 _dafny.Set) func(_dafny.Int) _dafny.Map {
				return func(_223_i15 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_dafny.SetOf((_220_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_220_v47), 0))).Int()).(bool), _221_v8)), _221_v8)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_222_v106), (_220_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_220_v47), 0))).Int()).(bool)))
				}
			})(_104_v47, _38_v8, _182_v106)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw31 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw31).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw31).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_219_v128 = _nw31
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_3
		var _nw32 _dafny.Array
		_ = _nw32
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw32 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Map = (func(_224_v106 _dafny.Set, _225_v8 bool) func(_dafny.Int) _dafny.Map {
				return func(_226_i16 _dafny.Int) _dafny.Map {
					return func() _dafny.Map {
						var _coll9 = _dafny.NewMapBuilder()
						_ = _coll9
						for _iter12 := _dafny.Iterate((_dafny.SeqOf(Companion_D2_.Create_DC6_(_224_v106), Companion_D2_.Create_DC6_(_224_v106), Companion_D2_.Create_DC6_(_224_v106), Companion_D2_.Create_DC6_(_224_v106))).Elements()); ; {
							_compr_9, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _227_v129 D2
							_227_v129 = interface{}(_compr_9).(D2)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D2_.Create_DC6_(_224_v106), Companion_D2_.Create_DC6_(_224_v106), Companion_D2_.Create_DC6_(_224_v106), Companion_D2_.Create_DC6_(_224_v106)), _227_v129) {
								_coll9.Add(_227_v129, _225_v8)
							}
						}
						return _coll9.ToMap()
					}()
				}
			})(_182_v106, _38_v8)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw32 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw32).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw32).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_219_v128 = _nw32
		var _228_v130 *C0
		_ = _228_v130
		var _nw33 *C0 = New_C0_()
		_ = _nw33
		_nw33.Ctor__()
		_228_v130 = _nw33
		(_41_globalState).F16 = (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)
		var _229_v131 D1
		_ = _229_v131
		_229_v131 = Companion_D1_.Create_DC4_((_179_v103).Cardinality(), _108_v51, _35_v2, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool))
		(_41_globalState).F7 = (_229_v131).Dtor_cf8()
	}
	var _230_v132 D3
	_ = _230_v132
	_230_v132 = Companion_D3_.Create_DC9_(_185_v109)
	var _source6 D3 = _230_v132
	_ = _source6
	if _source6.Is_DC10() {
		var _231___mcc_h12 _dafny.Int = _source6.Get_().(D3_DC10).Cf15
		_ = _231___mcc_h12
		var _232_cf15 _dafny.Int = _231___mcc_h12
		_ = _232_cf15
		var _233_v133 _dafny.Set
		_ = _233_v133
		_233_v133 = _dafny.SetOf(_dafny.IntOfInt64(867))
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
		_ = _index48
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
		_ = _index49
		var _rhs71 bool = _38_v8
		_ = _rhs71
		var _rhs72 _dafny.Int = ((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)).Times((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
		_ = _rhs72
		var _rhs73 _dafny.Int = _232_cf15
		_ = _rhs73
		var _rhs74 bool = (_233_v133).Equals((func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(669), _dafny.IntOfInt64(51))); ; {
				_compr_10, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _234_v134 _dafny.Int
				_234_v134 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(669)).Cmp(_234_v134) <= 0) && ((_234_v134).Cmp(_dafny.IntOfInt64(51)) < 0) {
					_coll10.Add((_234_v134).Plus(_232_cf15))
				}
			}
			return _coll10.ToSet()
		}()).Intersection(_233_v133))
		_ = _rhs74
		var _rhs75 bool = (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool)
		_ = _rhs75
		var _lhs64 _dafny.Array = _104_v47
		_ = _lhs64
		var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
		_ = _lhs65
		var _lhs66 _dafny.Array = _104_v47
		_ = _lhs66
		var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))
		_ = _lhs67
		var _lhs68 *GlobalState = _41_globalState
		_ = _lhs68
		(_lhs64).ArraySet1(_rhs71, (_lhs65).Int())
		_35_v2 = _rhs72
		_232_cf15 = _rhs73
		(_lhs66).ArraySet1(_rhs74, (_lhs67).Int())
		_lhs68.F7 = _rhs75
		(_41_globalState).F6 = _38_v8
		var _235_v135 _dafny.Map
		_ = _235_v135
		_235_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v97, _104_v47)
		var _236_v136 _dafny.Array
		_ = _236_v136
		var _nwElement0_4 _dafny.Array = (func() _dafny.Array {
			if (_235_v135).Contains(_171_v97) {
				return (_235_v135).Get(_171_v97).(_dafny.Array)
			}
			return _104_v47
		})()
		_ = _nwElement0_4
		var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(2))
		_ = _nw34
		(_nw34).ArraySet1(_nwElement0_4, 0)
		(_nw34).ArraySet1(_104_v47, 1)
		_236_v136 = _nw34
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_236_v136), 0))
		_ = _index50
		(_236_v136).ArraySet1(_104_v47, (_index50).Int())
		var _237_v137 D2
		_ = _237_v137
		_237_v137 = Companion_D2_.Create_DC7_(_33_v0)
		var _238_v138 _dafny.MultiSet
		_ = _238_v138
		_238_v138 = _dafny.MultiSetOf(_237_v137)
		var _239_v139 _dafny.Set
		_ = _239_v139
		_239_v139 = _dafny.SetOf(_171_v97)
		var _240_v140 _dafny.Map
		_ = _240_v140
		_240_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v8, _239_v139)
		var _241_v141 D6
		_ = _241_v141
		_241_v141 = Companion_D6_.Create_DC19_(_239_v139)
		var _242_v142 _dafny.Array
		_ = _242_v142
		var _nwElement0_5 _dafny.Set = _239_v139
		_ = _nwElement0_5
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_5, 0)
		(_nw35).ArraySet1(_239_v139, 1)
		(_nw35).ArraySet1((func() _dafny.Set {
			if (_240_v140).Contains(true) {
				return (_240_v140).Get(true).(_dafny.Set)
			}
			return _dafny.SetOf(_171_v97)
		})(), 2)
		(_nw35).ArraySet1((_239_v139).Union(_239_v139), 3)
		(_nw35).ArraySet1(_239_v139, 4)
		(_nw35).ArraySet1((_239_v139).Union(_239_v139), 5)
		(_nw35).ArraySet1(_239_v139, 6)
		(_nw35).ArraySet1(_239_v139, 7)
		(_nw35).ArraySet1((_dafny.SetOf(_171_v97, _171_v97)).Union(_239_v139), 8)
		(_nw35).ArraySet1(_239_v139, 9)
		(_nw35).ArraySet1(((_241_v141).Dtor_cf30()).Intersection(_dafny.SetOf(_171_v97, _171_v97, _171_v97)), 10)
		(_nw35).ArraySet1((_239_v139).Intersection(_239_v139), 11)
		(_nw35).ArraySet1(_239_v139, 12)
		(_nw35).ArraySet1(_239_v139, 13)
		(_nw35).ArraySet1(_239_v139, 14)
		(_nw35).ArraySet1(_239_v139, 15)
		(_nw35).ArraySet1(_dafny.SetOf(_171_v97, _171_v97, _171_v97), 16)
		(_nw35).ArraySet1(_239_v139, 17)
		(_nw35).ArraySet1(_dafny.SetOf(_171_v97), 18)
		(_nw35).ArraySet1((_239_v139).Union(_239_v139), 19)
		(_nw35).ArraySet1(_239_v139, 20)
		(_nw35).ArraySet1((_239_v139).Intersection(_239_v139), 21)
		(_nw35).ArraySet1(_dafny.SetOf(_171_v97, _171_v97, _171_v97, _171_v97, (_172_v98).Select((Companion_Default___.SafeIndex((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_172_v98).Cardinality()))).Uint32()).(*C1)), 22)
		(_nw35).ArraySet1((_239_v139).Difference(_239_v139), 23)
		(_nw35).ArraySet1(_239_v139, 24)
		_242_v142 = _nw35
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_242_v142), 0))
		_ = _index51
		(_242_v142).ArraySet1(_dafny.SetOf(_171_v97), (_index51).Int())
		var _243_v143 _dafny.Sequence
		_ = _243_v143
		_243_v143 = _dafny.SeqOf(_237_v137, _237_v137)
		var _pat_let_tv2 = _33_v0
		_ = _pat_let_tv2
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_236_v136), 0))
		_ = _index52
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_242_v142), 0))
		_ = _index53
		var _rhs76 _dafny.Array = (_181_v105).Select((Companion_Default___.SafeIndex(_35_v2, _dafny.IntOfUint32((_181_v105).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs76
		var _rhs77 _dafny.MultiSet = (func() _dafny.MultiSet {
			if Companion_Default___.Fm1(_37_v4, _38_v8, _35_v2, _41_globalState) {
				return (Companion_D7_.Create_DC21_(_dafny.MultiSetFromSeq(_243_v143))).Dtor_cf34()
			}
			return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_243_v143, (Companion_Default___.SafeIndex(_232_cf15, _dafny.IntOfUint32((_243_v143).Cardinality()))).Uint32(), func(_pat_let4_0 D2) D2 {
				return func(_244_dt__update__tmp_h2 D2) D2 {
					return func(_pat_let5_0 _dafny.CodePoint) D2 {
						return func(_245_dt__update_hcf12_h1 _dafny.CodePoint) D2 {
							return Companion_D2_.Create_DC7_(_245_dt__update_hcf12_h1)
						}(_pat_let5_0)
					}(_pat_let_tv2)
				}(_pat_let4_0)
			}(_237_v137)))
		})()
		_ = _rhs77
		var _rhs78 bool = _38_v8
		_ = _rhs78
		var _rhs79 bool = _38_v8
		_ = _rhs79
		var _rhs80 _dafny.Set = _dafny.SetOf(_171_v97, _171_v97)
		_ = _rhs80
		var _lhs69 _dafny.Array = _236_v136
		_ = _lhs69
		var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(471), _dafny.ArrayLen((_236_v136), 0))
		_ = _lhs70
		var _lhs71 *GlobalState = _41_globalState
		_ = _lhs71
		var _lhs72 *GlobalState = _41_globalState
		_ = _lhs72
		var _lhs73 _dafny.Array = _242_v142
		_ = _lhs73
		var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(604), _dafny.ArrayLen((_242_v142), 0))
		_ = _lhs74
		(_lhs69).ArraySet1(_rhs76, (_lhs70).Int())
		_238_v138 = _rhs77
		_lhs71.F6 = _rhs78
		_lhs72.F7 = _rhs79
		(_lhs73).ArraySet1(_rhs80, (_lhs74).Int())
		_232_cf15 = _232_cf15
	} else {
		var _246___mcc_h13 *C0 = _source6.Get_().(D3_DC9).Cf14
		_ = _246___mcc_h13
		var _247_cf14 *C0 = _246___mcc_h13
		_ = _247_cf14
		var _248_v144 _dafny.Set
		_ = _248_v144
		_248_v144 = _dafny.SetOf((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int))
		var _249_v145 _dafny.Array
		_ = _249_v145
		var _250_v146 _dafny.Map
		_ = _250_v146
		var _out9 _dafny.Array
		_ = _out9
		var _out10 _dafny.Map
		_ = _out10
		_out9, _out10 = (_171_v97).M1((_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int), _248_v144, (_183_v107).Cardinality(), _41_globalState)
		_249_v145 = _out9
		_250_v146 = _out10
		(_41_globalState).F0 = (_111_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(741), _dafny.ArrayLen((_111_v53), 0))).Int()).(_dafny.Int)
		var _251_v147 bool
		_ = _251_v147
		var _out11 bool
		_ = _out11
		_out11 = Companion_Default___.M0(false, (_104_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_104_v47), 0))).Int()).(bool), _38_v8, _38_v8, _41_globalState)
		_251_v147 = _out11
		(_41_globalState).F0 = _35_v2
	}
	_dafny.Print(_33_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_34_v1).Equals(_dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('n'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_35_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_36_v3, _dafny.SeqOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_37_v4, _dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_38_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_39_v9, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_40_v10, _dafny.SeqOf(_dafny.SeqOf(false, false), _dafny.SeqOf(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_41_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_41_globalState).F1(), _dafny.SeqOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_41_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_41_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_41_globalState.F11.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_41_globalState).F12()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_41_globalState).F14()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap(), _dafny.IntOfInt64(863))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_41_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_41_globalState).F17(), _dafny.SeqOf(_dafny.SeqOf(false, false), _dafny.SeqOf(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_41_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v11).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_51_v11).Dtor_cf1(), _dafny.SeqOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v47).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v48).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_105_v48).Dtor_cf1(), _dafny.SeqOf(true, true, false, false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_106_v49).Dtor_cf2()).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_106_v49).Dtor_cf2()).Dtor_cf1(), _dafny.SeqOf(true, true, false, false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_107_v50).Dtor_cf2()).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_107_v50).Dtor_cf2()).Dtor_cf1(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_108_v51).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), _dafny.IntOfInt64(5))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v52).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_111_v53).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_112_v54).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v55).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_170_v96).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(676))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v97.F19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(676))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_171_v97).F20()).Dtor_cf2()).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((((_171_v97).F20()).Dtor_cf2()).Dtor_cf1(), _dafny.SeqOf(true, true, false, false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_172_v98).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_173_v99).Dtor_cf16()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_174_v100).Dtor_cf21()).Dtor_cf16()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_v101).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_178_v102).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v103).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(896))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v104).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(_dafny.IntOfInt64(896)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_181_v105).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_v106).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v107).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(484), _dafny.IntOfInt64(-1362))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_184_v108).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_186_v110).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_187_v111).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf24()).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf25()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_188_v112).Dtor_cf26().F19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(676))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_188_v112).Dtor_cf26()).F20()).Dtor_cf2()).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((((_188_v112).Dtor_cf26()).F20()).Dtor_cf2()).Dtor_cf1(), _dafny.SeqOf(true, true, false, false, true, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_188_v112).Dtor_cf27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_188_v112).Dtor_cf28(), _dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_189_v113).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v114).Cardinality())
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
	Cf1 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf2 D0
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf2 D0) D0 {
	return D0{D0_DC1{Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, _dafny.EmptySeq)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() D0 {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1.Equals(data2.Cf1)
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf2.Equals(data2.Cf2)
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
	Cf4 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Int) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf5 _dafny.Int
	Cf6 _dafny.Map
	Cf7 _dafny.Int
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf5 _dafny.Int, Cf6 _dafny.Map, Cf7 _dafny.Int, Cf8 bool) D1 {
	return D1{D1_DC4{Cf5, Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf9  _dafny.Array
	Cf10 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 _dafny.Array, Cf10 _dafny.Int) D1 {
	return D1{D1_DC5{Cf9, Cf10}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC2 struct {
	Cf3 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf3 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.Zero)
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6.Equals(data2.Cf6) && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf9 == data2.Cf9 && data1.Cf10.Cmp(data2.Cf10) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf3 == data2.Cf3
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
	Cf12 _dafny.CodePoint
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.CodePoint) D2 {
	return D2{D2_DC7{Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf13 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 bool) D2 {
	return D2{D2_DC8{Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC6 struct {
	Cf11 _dafny.Set
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Set) D2 {
	return D2{D2_DC6{Cf11}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.CodePoint('D'))
}

func (_this D2) Dtor_cf12() _dafny.CodePoint {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf11() _dafny.Set {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf11) + ")"
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
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf12 == data2.Cf12
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf13 == data2.Cf13
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf15 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf15 _dafny.Int) D3 {
	return D3{D3_DC10{Cf15}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf14 *C0
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf14 *C0) D3 {
	return D3{D3_DC9{Cf14}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.Zero)
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf14() *C0 {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf15) + ")"
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
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
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

type D4_DC12 struct {
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_() D4 {
	return D4{D4_DC12{}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf17 bool
	Cf18 _dafny.Int
	Cf19 _dafny.Array
	Cf20 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf17 bool, Cf18 _dafny.Int, Cf19 _dafny.Array, Cf20 bool) D4 {
	return D4{D4_DC13{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC11 struct {
	Cf16 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf16 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf16}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC14 struct {
	Cf21 D4
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf21 D4) D4 {
	return D4{D4_DC14{Cf21}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_()
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC13).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf16
}

func (_this D4) Dtor_cf21() D4 {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf21) + ")"
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
			_, ok := other.Get_().(D4_DC12)
			return ok
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf16.Equals(data2.Cf16)
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 _dafny.Int
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf23 _dafny.Int) D5 {
	return D5{D5_DC16{Cf23}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC17 struct {
	Cf24 _dafny.MultiSet
	Cf25 _dafny.Array
	Cf26 *C1
	Cf27 _dafny.Int
	Cf28 _dafny.Sequence
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf24 _dafny.MultiSet, Cf25 _dafny.Array, Cf26 *C1, Cf27 _dafny.Int, Cf28 _dafny.Sequence) D5 {
	return D5{D5_DC17{Cf24, Cf25, Cf26, Cf27, Cf28}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC15 struct {
	Cf22 _dafny.Array
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf22 _dafny.Array) D5 {
	return D5{D5_DC15{Cf22}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC18 struct {
	Cf29 D5
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf29 D5) D5 {
	return D5{D5_DC18{Cf29}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC16_(_dafny.Zero)
}

func (_this D5) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf23
}

func (_this D5) Dtor_cf24() _dafny.MultiSet {
	return _this.Get_().(D5_DC17).Cf24
}

func (_this D5) Dtor_cf25() _dafny.Array {
	return _this.Get_().(D5_DC17).Cf25
}

func (_this D5) Dtor_cf26() *C1 {
	return _this.Get_().(D5_DC17).Cf26
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D5_DC17).Cf28
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf22
}

func (_this D5) Dtor_cf29() D5 {
	return _this.Get_().(D5_DC18).Cf29
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + data.Cf28.VerbatimString(true) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf29) + ")"
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D5_DC17:
		{
			data2, ok := other.Get_().(D5_DC17)
			return ok && data1.Cf24.Equals(data2.Cf24) && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28.Equals(data2.Cf28)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf22 == data2.Cf22
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
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

type D6_DC20 struct {
	Cf31 bool
	Cf32 D5
	Cf33 _dafny.Set
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf31 bool, Cf32 D5, Cf33 _dafny.Set) D6 {
	return D6{D6_DC20{Cf31, Cf32, Cf33}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

type D6_DC19 struct {
	Cf30 _dafny.Set
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_(Cf30 _dafny.Set) D6 {
	return D6{D6_DC19{Cf30}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC20_(false, Companion_D5_.Default(), _dafny.EmptySet)
}

func (_this D6) Dtor_cf31() bool {
	return _this.Get_().(D6_DC20).Cf31
}

func (_this D6) Dtor_cf32() D5 {
	return _this.Get_().(D6_DC20).Cf32
}

func (_this D6) Dtor_cf33() _dafny.Set {
	return _this.Get_().(D6_DC20).Cf33
}

func (_this D6) Dtor_cf30() _dafny.Set {
	return _this.Get_().(D6_DC19).Cf30
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19" + "(" + _dafny.String(data.Cf30) + ")"
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
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32.Equals(data2.Cf32) && data1.Cf33.Equals(data2.Cf33)
		}
	case D6_DC19:
		{
			data2, ok := other.Get_().(D6_DC19)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf35 bool
	Cf36 _dafny.Int
	Cf37 _dafny.Int
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf35 bool, Cf36 _dafny.Int, Cf37 _dafny.Int) D7 {
	return D7{D7_DC22{Cf35, Cf36, Cf37}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC23 struct {
	Cf38 bool
	Cf39 _dafny.Int
}

func (D7_DC23) isD7() {}

func (CompanionStruct_D7_) Create_DC23_(Cf38 bool, Cf39 _dafny.Int) D7 {
	return D7{D7_DC23{Cf38, Cf39}}
}

func (_this D7) Is_DC23() bool {
	_, ok := _this.Get_().(D7_DC23)
	return ok
}

type D7_DC21 struct {
	Cf34 _dafny.MultiSet
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf34 _dafny.MultiSet) D7 {
	return D7{D7_DC21{Cf34}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC22_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D7) Dtor_cf35() bool {
	return _this.Get_().(D7_DC22).Cf35
}

func (_this D7) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf36
}

func (_this D7) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D7_DC22).Cf37
}

func (_this D7) Dtor_cf38() bool {
	return _this.Get_().(D7_DC23).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf39
}

func (_this D7) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D7_DC21).Cf34
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D7_DC23:
		{
			return "D7.DC23" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf34) + ")"
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
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37.Cmp(data2.Cf37) == 0
		}
	case D7_DC23:
		{
			data2, ok := other.Get_().(D7_DC23)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D8_DC25 struct {
	Cf41 _dafny.Int
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf41 _dafny.Int) D8 {
	return D8{D8_DC25{Cf41}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC26 struct {
	Cf42 _dafny.Int
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf42 _dafny.Int) D8 {
	return D8{D8_DC26{Cf42}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC24 struct {
	Cf40 _dafny.MultiSet
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf40 _dafny.MultiSet) D8 {
	return D8{D8_DC24{Cf40}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC25_(_dafny.Zero)
}

func (_this D8) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D8_DC25).Cf41
}

func (_this D8) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D8_DC26).Cf42
}

func (_this D8) Dtor_cf40() _dafny.MultiSet {
	return _this.Get_().(D8_DC24).Cf40
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf41.Cmp(data2.Cf41) == 0
		}
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F6   bool
	F7   bool
	F11  _dafny.Sequence
	F16  _dafny.Int
	_f1  _dafny.Sequence
	_f2  _dafny.Int
	_f3  bool
	_f4  _dafny.Int
	_f5  bool
	_f8  _dafny.Int
	_f9  bool
	_f10 bool
	_f12 _dafny.Map
	_f13 bool
	_f14 _dafny.Map
	_f15 bool
	_f17 _dafny.Sequence
	_f18 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F6 = false
	_this.F7 = false
	_this.F11 = _dafny.EmptySeq
	_this.F16 = _dafny.Zero
	_this._f1 = _dafny.EmptySeq
	_this._f2 = _dafny.Zero
	_this._f3 = false
	_this._f4 = _dafny.Zero
	_this._f5 = false
	_this._f8 = _dafny.Zero
	_this._f9 = false
	_this._f10 = false
	_this._f12 = _dafny.EmptyMap
	_this._f13 = false
	_this._f14 = _dafny.EmptyMap
	_this._f15 = false
	_this._f17 = _dafny.EmptySeq
	_this._f18 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 _dafny.Int, f3 bool, f4 _dafny.Int, f5 bool, f6 bool, f7 bool, f8 _dafny.Int, f9 bool, f10 bool, f11 _dafny.Sequence, f12 _dafny.Map, f13 bool, f14 _dafny.Map, f15 bool, f16 _dafny.Int, f17 _dafny.Sequence, f18 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
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
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Map {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Map {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
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
func (_this *C0) Fm7(p0 D0, p1 D0, globalState *GlobalState) _dafny.MultiSet {
	{
		if !(false) || (!(true)) {
			return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(538), Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-515))).Cardinality()), _dafny.SeqOf(true, true)))).Cardinality())).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(490), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-631), false)).Cardinality()))
		} else {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(666)))
		}
	}
}
func (_this *C0) Fm8(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var _source7 D0 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(856), _dafny.SeqOf(!(true)))
		_ = _source7
		if _source7.Is_DC0() {
			var _252___mcc_h0 _dafny.Int = _source7.Get_().(D0_DC0).Cf0
			_ = _252___mcc_h0
			var _253___mcc_h1 _dafny.Sequence = _source7.Get_().(D0_DC0).Cf1
			_ = _253___mcc_h1
			var _254_cf1 _dafny.Sequence = _253___mcc_h1
			_ = _254_cf1
			var _255_cf0 _dafny.Int = _252___mcc_h0
			_ = _255_cf0
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)))
		} else {
			var _256___mcc_h2 D0 = _source7.Get_().(D0_DC1).Cf2
			_ = _256___mcc_h2
			var _257_cf2 D0 = _256___mcc_h2
			_ = _257_cf2
			if false {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
			} else {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
			}
		}
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F19  _dafny.Map
	_f20 D0
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F19 = _dafny.EmptyMap
	_this._f20 = Companion_D0_.Default()
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

func (_this *C1) Ctor__(f19 _dafny.Map, f20 D0) {
	{
		(_this).F19 = f19
		(_this)._f20 = f20
	}
}
func (_this *C1) Fm6(p0 _dafny.CodePoint, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vc")).Cardinality())).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Minus(_dafny.IntOfInt64(471)))
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 _dafny.Set, p2 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Map) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _258_v0 _dafny.Sequence
		_ = _258_v0
		_258_v0 = _dafny.UnicodeSeqOfUtf8Bytes("iuyubyqfn")
		var _259_v1 _dafny.CodePoint
		_ = _259_v1
		_259_v1 = _dafny.CodePoint('y')
		var _hi1 _dafny.Int = p0
		_ = _hi1
		for _260_i0 := (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_258_v0, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_258_v0).Cardinality()))).Uint32(), _259_v1)).Cardinality())).Times(p0); _260_i0.Cmp(_hi1) < 0; _260_i0 = _260_i0.Plus(_dafny.One) {
			(globalState).F0 = _260_i0
			var _261_v2 bool
			_ = _261_v2
			_261_v2 = false
			var _262_v3 _dafny.Sequence
			_ = _262_v3
			_262_v3 = _dafny.SeqOf(_261_v2)
			(globalState).F7 = !((_262_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_262_v3).Cardinality()))).Uint32()).(bool)) || (_261_v2)
			(globalState).F6 = (_dafny.SetOf(p0)).Contains(p0)
			var _263_v4 *C0
			_ = _263_v4
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__()
			_263_v4 = _nw36
		}
		var _264_v5 _dafny.Array
		_ = _264_v5
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw37
		_264_v5 = _nw37
		var _265_v6 bool
		_ = _265_v6
		_265_v6 = true
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))
		_ = _index54
		(_264_v5).ArraySet1(_265_v6, (_index54).Int())
		var _266_v7 _dafny.Map
		_ = _266_v7
		_266_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v1, p0)
		var _267_v8 D1
		_ = _267_v8
		_267_v8 = Companion_D1_.Create_DC4_(p2, _266_v7, p2, _265_v6)
		var _268_v9 _dafny.Sequence
		_ = _268_v9
		_268_v9 = _dafny.SeqOf(_265_v6, _265_v6, _265_v6, (_267_v8).Dtor_cf8())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))
		_ = _index55
		(_264_v5).ArraySet1(!((_268_v9).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_268_v9).Cardinality()))).Uint32()).(bool)) || (_265_v6), (_index55).Int())
		var _269_v10 _dafny.Array
		_ = _269_v10
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_4
		var _nw38 _dafny.Array
		_ = _nw38
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw38 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = (func(_270_p0 _dafny.Int, _271_v5 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_272_i2 _dafny.Int) _dafny.Int {
					return (_272_i2).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_p0, !((_271_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_271_v5), 0))).Int()).(bool)))).Cardinality())
				}
			})(p0, _264_v5)
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
		_269_v10 = _nw38
		for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_269_v10), 0))); ; {
			_guard_loop_2, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _273_i1 _dafny.Int
			_273_i1 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_273_i1).Sign() != -1) && ((_273_i1).Cmp(_dafny.ArrayLen((_269_v10), 0)) < 0)) {
				(_269_v10).ArraySet1((_273_i1).Minus(_dafny.IntOfUint32((_258_v0).Cardinality())), (_273_i1).Int())
			}
		}
		var _274_v11 _dafny.Array
		_ = _274_v11
		var _nwElement0_6 _dafny.Array = _269_v10
		_ = _nwElement0_6
		var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(4))
		_ = _nw39
		(_nw39).ArraySet1(_nwElement0_6, 0)
		(_nw39).ArraySet1(_269_v10, 1)
		(_nw39).ArraySet1(_269_v10, 2)
		(_nw39).ArraySet1(_269_v10, 3)
		_274_v11 = _nw39
		var _275_v12 _dafny.Map
		_ = _275_v12
		_275_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_264_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))).Int()).(bool), _274_v11)
		_275_v12 = (_275_v12).Update(_265_v6, _274_v11)
		var _276_v13 D1
		_ = _276_v13
		_276_v13 = Companion_D1_.Create_DC5_(_264_v5, p2)
		var _277_v14 D2
		_ = _277_v14
		_277_v14 = Companion_D2_.Create_DC6_(_dafny.SetOf(_265_v6))
		var _278_v15 D1
		_ = _278_v15
		_278_v15 = Companion_D1_.Create_DC5_((_276_v13).Dtor_cf9(), ((_277_v14).Dtor_cf11()).Cardinality())
		var _279_v16 _dafny.Sequence
		_ = _279_v16
		_279_v16 = _dafny.SeqOf(_264_v5, _264_v5, _264_v5)
		var _280_v17 _dafny.Array
		_ = _280_v17
		var _nwElement0_7 _dafny.Array = _264_v5
		_ = _nwElement0_7
		var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(28))
		_ = _nw40
		(_nw40).ArraySet1(_nwElement0_7, 0)
		(_nw40).ArraySet1(_264_v5, 1)
		(_nw40).ArraySet1(_264_v5, 2)
		(_nw40).ArraySet1(_264_v5, 3)
		(_nw40).ArraySet1(_264_v5, 4)
		(_nw40).ArraySet1(_264_v5, 5)
		(_nw40).ArraySet1((_278_v15).Dtor_cf9(), 6)
		(_nw40).ArraySet1(_264_v5, 7)
		(_nw40).ArraySet1(_264_v5, 8)
		(_nw40).ArraySet1(_264_v5, 9)
		(_nw40).ArraySet1(_264_v5, 10)
		(_nw40).ArraySet1(_264_v5, 11)
		(_nw40).ArraySet1(_264_v5, 12)
		(_nw40).ArraySet1(_264_v5, 13)
		(_nw40).ArraySet1(_264_v5, 14)
		(_nw40).ArraySet1(_264_v5, 15)
		(_nw40).ArraySet1(_264_v5, 16)
		(_nw40).ArraySet1(_264_v5, 17)
		(_nw40).ArraySet1(_264_v5, 18)
		(_nw40).ArraySet1((_279_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.IntOfUint32((_279_v16).Cardinality()))).Uint32()).(_dafny.Array), 19)
		(_nw40).ArraySet1(_264_v5, 20)
		(_nw40).ArraySet1(_264_v5, 21)
		(_nw40).ArraySet1(_264_v5, 22)
		(_nw40).ArraySet1(_264_v5, 23)
		(_nw40).ArraySet1(_264_v5, 24)
		(_nw40).ArraySet1(_264_v5, 25)
		(_nw40).ArraySet1(_264_v5, 26)
		(_nw40).ArraySet1(_264_v5, 27)
		_280_v17 = _nw40
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_280_v17), 0))
		_ = _index56
		(_280_v17).ArraySet1(_264_v5, (_index56).Int())
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_280_v17), 0))
		_ = _index57
		(_280_v17).ArraySet1(_264_v5, (_index57).Int())
		var _281_v18 _dafny.Sequence
		_ = _281_v18
		_281_v18 = _dafny.SeqOf(_258_v0, _258_v0)
		var _282_v19 *C0
		_ = _282_v19
		var _nw41 *C0 = New_C0_()
		_ = _nw41
		_nw41.Ctor__()
		_282_v19 = _nw41
		var _283_v20 _dafny.Sequence
		_ = _283_v20
		_283_v20 = _dafny.SeqOf(p2)
		var _284_v21 _dafny.MultiSet
		_ = _284_v21
		_284_v21 = _dafny.MultiSetOf(p2, _dafny.IntOfUint32((_283_v20).Cardinality()))
		var _285_v22 _dafny.Map
		_ = _285_v22
		_285_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_284_v21).Cardinality(), _282_v19)
		var _286_v23 _dafny.Sequence
		_ = _286_v23
		_286_v23 = _dafny.SeqOf(_282_v19, _282_v19)
		var _287_v24 _dafny.Array
		_ = _287_v24
		var _nwElement0_8 *C0 = _282_v19
		_ = _nwElement0_8
		var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(28))
		_ = _nw42
		(_nw42).ArraySet1(_nwElement0_8, 0)
		(_nw42).ArraySet1((func() *C0 {
			if (_285_v22).Contains(p2) {
				return (_285_v22).Get(p2).(*C0)
			}
			return _282_v19
		})(), 1)
		(_nw42).ArraySet1(_282_v19, 2)
		(_nw42).ArraySet1(_282_v19, 3)
		(_nw42).ArraySet1(_282_v19, 4)
		(_nw42).ArraySet1(_282_v19, 5)
		(_nw42).ArraySet1(_282_v19, 6)
		(_nw42).ArraySet1(_282_v19, 7)
		(_nw42).ArraySet1((Companion_D3_.Create_DC9_(_282_v19)).Dtor_cf14(), 8)
		(_nw42).ArraySet1(_282_v19, 9)
		(_nw42).ArraySet1(_282_v19, 10)
		(_nw42).ArraySet1((_286_v23).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_286_v23).Cardinality()))).Uint32()).(*C0), 11)
		(_nw42).ArraySet1(_282_v19, 12)
		(_nw42).ArraySet1(_282_v19, 13)
		(_nw42).ArraySet1(_282_v19, 14)
		(_nw42).ArraySet1(_282_v19, 15)
		(_nw42).ArraySet1(_282_v19, 16)
		(_nw42).ArraySet1(_282_v19, 17)
		(_nw42).ArraySet1(_282_v19, 18)
		(_nw42).ArraySet1(_282_v19, 19)
		(_nw42).ArraySet1(_282_v19, 20)
		(_nw42).ArraySet1(_282_v19, 21)
		(_nw42).ArraySet1(_282_v19, 22)
		(_nw42).ArraySet1(_282_v19, 23)
		(_nw42).ArraySet1(_282_v19, 24)
		(_nw42).ArraySet1(_282_v19, 25)
		(_nw42).ArraySet1(_282_v19, 26)
		(_nw42).ArraySet1(_282_v19, 27)
		_287_v24 = _nw42
		var _288_v25 _dafny.Array
		_ = _288_v25
		var _nwElement0_9 _dafny.Array = _287_v24
		_ = _nwElement0_9
		var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(3))
		_ = _nw43
		(_nw43).ArraySet1(_nwElement0_9, 0)
		(_nw43).ArraySet1(_287_v24, 1)
		(_nw43).ArraySet1(_287_v24, 2)
		_288_v25 = _nw43
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_288_v25), 0))
		_ = _index58
		(_288_v25).ArraySet1(_287_v24, (_index58).Int())
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_288_v25), 0))
		_ = _index59
		var _rhs81 _dafny.Sequence = _281_v18
		_ = _rhs81
		var _rhs82 _dafny.Int = p0
		_ = _rhs82
		var _rhs83 _dafny.Array = _287_v24
		_ = _rhs83
		var _rhs84 bool = (_264_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))).Int()).(bool)
		_ = _rhs84
		var _lhs75 *GlobalState = globalState
		_ = _lhs75
		var _lhs76 _dafny.Array = _288_v25
		_ = _lhs76
		var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_288_v25), 0))
		_ = _lhs77
		_281_v18 = _rhs81
		_lhs75.F0 = _rhs82
		(_lhs76).ArraySet1(_rhs83, (_lhs77).Int())
		_265_v6 = _rhs84
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_5
		var _nw44 _dafny.Array
		_ = _nw44
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw44 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) bool = (func(_289_v6 bool) func(_dafny.Int) bool {
				return func(_290_i3 _dafny.Int) bool {
					return _289_v6
				}
			})(_265_v6)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw44 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw44).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw44).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		r0 = _nw44
		var _291_v26 _dafny.Sequence
		_ = _291_v26
		_291_v26 = _dafny.SeqOf(_283_v20, _283_v20, _283_v20, _283_v20, _283_v20)
		var _292_v27 _dafny.Map
		_ = _292_v27
		_292_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_291_v26).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_291_v26).Cardinality()))).Uint32()).(_dafny.Sequence), _265_v6)
		var _293_v28 _dafny.Map
		_ = _293_v28
		_293_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_264_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))).Int()).(bool), _283_v20)
		var _294_v29 _dafny.Map
		_ = _294_v29
		_294_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-977), p2)
		r1 = ((_292_v27).Update((func() _dafny.Sequence {
			if (_293_v28).Contains((_264_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))).Int()).(bool)) {
				return (_293_v28).Get((_264_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))).Int()).(bool)).(_dafny.Sequence)
			}
			return _283_v20
		})(), (_264_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_264_v5), 0))).Int()).(bool))).Update(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(722))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_295_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_296_i4 _dafny.Int) _dafny.Int {
				return _295_p2
			}
		})(p2))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(722))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_297_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_298_i4 _dafny.Int) _dafny.Int {
				return _297_p2
			}
		})(p2)))).Cardinality()))).Uint32(), (_294_v29).Cardinality()), _265_v6)
		return r0, r1
	}
}
func (_this *C1) F20() D0 {
	{
		return _this._f20
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
