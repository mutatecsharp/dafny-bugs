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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Int {
	return (Companion_D0_.Create_DC2_(false, false, _dafny.IntOfInt64(-941))).Dtor_cf4()
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("csbalwk"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-85))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(394))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})))
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	if !(!(true)) || (false) {
		return (_dafny.SetOf(_dafny.IntOfInt64(-71))).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfInt64(-563)))
	} else {
		return false
	}
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, !(!(!(false)))), _dafny.SeqOf(true)), _dafny.SeqOf(!(!(true)), true))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(450), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-488), false))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false, !(true))).Difference(_dafny.SetOf(false, true))).Intersection(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) D0 {
	if !(!(false)) {
		if true {
			return Companion_D0_.Create_DC2_(false, true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(981), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(758))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			}))).Cardinality()))).Cardinality()))
		} else {
			return Companion_D0_.Create_DC2_(!(false), (Companion_D3_.Create_DC8_(_dafny.IntOfInt64(431), _dafny.IntOfInt64(-307), true)).Dtor_cf19(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())
		}
	} else {
		return Companion_D0_.Create_DC2_((Companion_D0_.Create_DC2_(!(true), true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(943))).Cardinality()))).Dtor_cf3(), true, _dafny.IntOfInt64(567))
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.CodePoint, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("ft"), ((_dafny.SetOf(true)).Cardinality()).Plus((Companion_D3_.Create_DC8_((_dafny.SetOf((_dafny.SetOf(true, true)).Cardinality())).Cardinality(), _dafny.IntOfInt64(798), !(!(true)))).Dtor_cf17()), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).IsProperSubsetOf(func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _3_v0 _dafny.Map
			_3_v0 = interface{}(_compr_0).(_dafny.Map)
			if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Contains(_3_v0) {
				_coll0.Add(_3_v0)
			}
		}
		return _coll0.ToSet()
	}()), ((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(true, true)
		}
		return _dafny.MultiSetOf((Companion_D3_.Create_DC8_(_dafny.IntOfInt64(171), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("otasatp")).Cardinality()), false)).Dtor_cf19(), false, false)
	})()).Cardinality(), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality())), _dafny.SeqOf(false))).Cardinality(), _dafny.IntOfInt64(118))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfInt64(587))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfInt64(239))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) D4 {
	if !((false) == (!(!(true)))) {
		return Companion_D4_.Create_DC9_(func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D3_.Create_DC8_(_dafny.IntOfInt64(764), _dafny.IntOfInt64(168), !(false)))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _4_v0 D3
				_4_v0 = interface{}(_compr_1).(D3)
				if (_dafny.MultiSetOf(Companion_D3_.Create_DC8_(_dafny.IntOfInt64(764), _dafny.IntOfInt64(168), !(false)))).Contains(_4_v0) {
					_coll1.Add(_4_v0, _dafny.SeqOf(_dafny.CodePoint('u')))
				}
			}
			return _coll1.ToMap()
		}())
	} else if true {
		return Companion_D4_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.IntOfInt64(-664), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(90))).Cardinality()), true), _dafny.SeqOf(_dafny.CodePoint('x'))))
	} else {
		return Companion_D4_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-837), true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('i'))).Cardinality()))).Cardinality(), true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(82))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfInt64(305))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.SeqOf(!(false), true))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(true), !((!(!((Companion_D0_.Create_DC2_(true, false, _dafny.IntOfInt64(735))).Dtor_cf2()))) && (false)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D2 = (func() D2 {
		if !(false) {
			return Companion_D2_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("mjffy"), (_dafny.MultiSetOf(false)).Cardinality(), false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false, false, !(false)))).Cardinality(), false)
		}
		return Companion_D2_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("l"), _dafny.IntOfInt64(139), !(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dop")).Cardinality()), false)
	})()
	_ = _source0
	if _source0.Is_DC6() {
		var _6___mcc_h0 _dafny.Sequence = _source0.Get_().(D2_DC6).Cf11
		_ = _6___mcc_h0
		var _7___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC6).Cf12
		_ = _7___mcc_h1
		var _8___mcc_h2 bool = _source0.Get_().(D2_DC6).Cf13
		_ = _8___mcc_h2
		var _9___mcc_h3 _dafny.Int = _source0.Get_().(D2_DC6).Cf14
		_ = _9___mcc_h3
		var _10___mcc_h4 bool = _source0.Get_().(D2_DC6).Cf15
		_ = _10___mcc_h4
		var _11_cf15 bool = _10___mcc_h4
		_ = _11_cf15
		var _12_cf14 _dafny.Int = _9___mcc_h3
		_ = _12_cf14
		var _13_cf13 bool = _8___mcc_h2
		_ = _13_cf13
		var _14_cf12 _dafny.Int = _7___mcc_h1
		_ = _14_cf12
		var _15_cf11 _dafny.Sequence = _6___mcc_h0
		_ = _15_cf11
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(780))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}((func(_16_cf14 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_17_i0 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll2 = _dafny.NewMapBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(790), _dafny.IntOfInt64(970))); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _18_v0 _dafny.Int
						_18_v0 = interface{}(_compr_2).(_dafny.Int)
						if ((_dafny.IntOfInt64(790)).Cmp(_18_v0) <= 0) && ((_18_v0).Cmp(_dafny.IntOfInt64(970)) < 0) {
							_coll2.Add(Companion_Default___.SafeDivisionInt(_18_v0, _16_cf14), _dafny.IntOfInt64(857))
						}
					}
					return _coll2.ToMap()
				}()
			}
		})(_12_cf14))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(701))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_19_cf12 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_20_i1 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-290), _19_cf12)
			}
		})(_14_cf12))))
	} else {
		var _21___mcc_h5 _dafny.Sequence = _source0.Get_().(D2_DC5).Cf10
		_ = _21___mcc_h5
		var _22_cf10 _dafny.Sequence = _21___mcc_h5
		_ = _22_cf10
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(980))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_23_i2 _dafny.Int) _dafny.Map {
			return (func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-768), _dafny.IntOfInt64(394))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _24_v1 _dafny.Int
					_24_v1 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(-768)).Cmp(_24_v1) <= 0) && ((_24_v1).Cmp(_dafny.IntOfInt64(394)) < 0) {
						_coll3.Add((_24_v1).Times(_dafny.IntOfInt64(-880)), _dafny.IntOfInt64(586))
					}
				}
				return _coll3.ToMap()
			}())
		}))
	}
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) {
	if p0 {
		var _25_v0 bool
		_ = _25_v0
		_25_v0 = true
		var _26_v1 _dafny.Sequence
		_ = _26_v1
		_26_v1 = _dafny.SeqOf(p3)
		_25_v0 = !_dafny.Companion_Sequence_.Contains(_26_v1, p1)
		var _27_v2 D3
		_ = _27_v2
		_27_v2 = Companion_D3_.Create_DC8_(_dafny.IntOfInt64(-73), p3, _25_v0)
		var _28_v3 *C0
		_ = _28_v3
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__(_27_v2)
		_28_v3 = _nw0
		var _29_v4 _dafny.Sequence
		_ = _29_v4
		_29_v4 = _dafny.SeqOf(_25_v0)
		var _30_v5 _dafny.Map
		_ = _30_v5
		_30_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_29_v4).Cardinality()), _dafny.IntOfUint32((_29_v4).Cardinality()))
		var _31_v6 _dafny.Map
		_ = _31_v6
		_31_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, ((_dafny.MultiSetOf(p0, p2)).Update(p0, Companion_Default___.Abs(p3))).Cardinality())
		var _32_v7 _dafny.Array
		_ = _32_v7
		var _nw1 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw1
		_32_v7 = _nw1
		var _33_v8 D4
		_ = _33_v8
		_33_v8 = Companion_D4_.Create_DC12_(_25_v0, _dafny.SeqOf(_30_v5, _30_v5, _30_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_31_v6).Cardinality()), _30_v5), p1, _32_v7, p0)
		var _34_v9 _dafny.MultiSet
		_ = _34_v9
		_34_v9 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v0, (_33_v8).Dtor_cf28()), _31_v6, _31_v6)
		var _35_v10 D4
		_ = _35_v10
		_35_v10 = Companion_D4_.Create_DC11_(_28_v3, (func() _dafny.Int {
			if (_34_v9).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v0, (_28_v3).Fm6(p0, p1, !(true), globalState))) {
				return (_34_v9).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v0, (_28_v3).Fm6(p0, p1, !(true), globalState)))
			}
			return p3
		})(), p3)
		_35_v10 = _35_v10
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_32_v7), 0))
		_ = _index0
		(_32_v7).ArraySet1(false, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_32_v7), 0))
		_ = _index1
		(_32_v7).ArraySet1(Companion_Default___.Fm2(globalState), (_index1).Int())
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_32_v7), 0))
		_ = _index2
		(_32_v7).ArraySet1(p0, (_index2).Int())
		if (p2) && (Companion_Default___.Fm2(globalState)) {
			var _36_v11 *C1
			_ = _36_v11
			var _nw2 *C1 = New_C1_()
			_ = _nw2
			_nw2.Ctor__()
			_36_v11 = _nw2
			var _37_v12 _dafny.Array
			_ = _37_v12
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
			_ = _nw3
			_37_v12 = _nw3
			var _38_v13 _dafny.CodePoint
			_ = _38_v13
			_38_v13 = _dafny.CodePoint('j')
			var _39_v14 _dafny.Map
			_ = _39_v14
			_39_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v13, _29_v4)
			var _40_v15 D0
			_ = _40_v15
			_40_v15 = Companion_D0_.Create_DC0_(_dafny.SeqOf(!(p2)))
			var _41_v16 _dafny.Array
			_ = _41_v16
			var _nwElement0_0 _dafny.Sequence = _dafny.SeqOf(p0)
			_ = _nwElement0_0
			var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(16))
			_ = _nw4
			(_nw4).ArraySet1(_nwElement0_0, 0)
			(_nw4).ArraySet1(_29_v4, 1)
			(_nw4).ArraySet1(_29_v4, 2)
			(_nw4).ArraySet1(_dafny.SeqOf(p2, (_32_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_32_v7), 0))).Int()).(bool)), 3)
			(_nw4).ArraySet1(_29_v4, 4)
			(_nw4).ArraySet1(_29_v4, 5)
			(_nw4).ArraySet1(_dafny.SeqOf(false, true, _25_v0), 6)
			(_nw4).ArraySet1(_29_v4, 7)
			(_nw4).ArraySet1(_dafny.SeqOf(true), 8)
			(_nw4).ArraySet1(_29_v4, 9)
			(_nw4).ArraySet1(_29_v4, 10)
			(_nw4).ArraySet1(_dafny.SeqOf(_25_v0, (_32_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_32_v7), 0))).Int()).(bool)), 11)
			(_nw4).ArraySet1(_29_v4, 12)
			(_nw4).ArraySet1(_29_v4, 13)
			(_nw4).ArraySet1(_29_v4, 14)
			(_nw4).ArraySet1((func() _dafny.Sequence {
				if (_39_v14).Contains(_38_v13) {
					return (_39_v14).Get(_38_v13).(_dafny.Sequence)
				}
				return (_40_v15).Dtor_cf0()
			})(), 15)
			_41_v16 = _nw4
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_37_v12), 0))
			_ = _index3
			(_37_v12).ArraySet1((func() _dafny.Array {
				if !((_32_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_32_v7), 0))).Int()).(bool)) {
					return _41_v16
				}
				return _41_v16
			})(), (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_37_v12), 0))
			_ = _index4
			(_37_v12).ArraySet1(_41_v16, (_index4).Int())
			var _42_v17 _dafny.Sequence
			_ = _42_v17
			_42_v17 = _dafny.UnicodeSeqOfUtf8Bytes("d")
			var _43_v18 _dafny.Map
			_ = _43_v18
			_43_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Cmp(_dafny.IntOfUint32((_42_v17).Cardinality())) > 0, _dafny.CodePoint('i'))
			var _44_v19 _dafny.CodePoint
			_ = _44_v19
			_44_v19 = _38_v13
			_43_v18 = (_43_v18).Update(p2, (_44_v19))
			_27_v2 = _27_v2
			var _45_v20 *C0
			_ = _45_v20
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(_27_v2)
			_45_v20 = _nw5
		} else {
			var _46_v21 *C1
			_ = _46_v21
			var _nw6 *C1 = New_C1_()
			_ = _nw6
			_nw6.Ctor__()
			_46_v21 = _nw6
			var _47_v22 *C1
			_ = _47_v22
			var _nw7 *C1 = New_C1_()
			_ = _nw7
			_nw7.Ctor__()
			_47_v22 = _nw7
			(globalState).F0 = p3
			var _48_v23 _dafny.Sequence
			_ = _48_v23
			_48_v23 = _dafny.UnicodeSeqOfUtf8Bytes("mmk")
			_48_v23 = _dafny.Companion_Sequence_.Update(_48_v23, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_48_v23).Cardinality()))).Uint32(), (_48_v23).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_48_v23).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _49_v24 _dafny.Sequence
			_ = _49_v24
			_49_v24 = _dafny.SeqOf(_28_v3)
			var _50_v25 D1
			_ = _50_v25
			_50_v25 = Companion_D1_.Create_DC4_(p3, !(p0), true, false)
			var _51_v26 _dafny.Array
			_ = _51_v26
			var _nwElement0_1 *C0 = _28_v3
			_ = _nwElement0_1
			var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(19))
			_ = _nw8
			(_nw8).ArraySet1(_nwElement0_1, 0)
			(_nw8).ArraySet1(_28_v3, 1)
			(_nw8).ArraySet1(_28_v3, 2)
			(_nw8).ArraySet1(_28_v3, 3)
			(_nw8).ArraySet1((_49_v24).Select((Companion_Default___.SafeIndex((_50_v25).Dtor_cf6(), _dafny.IntOfUint32((_49_v24).Cardinality()))).Uint32()).(*C0), 4)
			(_nw8).ArraySet1(_28_v3, 5)
			(_nw8).ArraySet1(_28_v3, 6)
			(_nw8).ArraySet1(_28_v3, 7)
			(_nw8).ArraySet1(_28_v3, 8)
			(_nw8).ArraySet1(_28_v3, 9)
			(_nw8).ArraySet1(_28_v3, 10)
			(_nw8).ArraySet1(_28_v3, 11)
			(_nw8).ArraySet1(_28_v3, 12)
			(_nw8).ArraySet1(_28_v3, 13)
			(_nw8).ArraySet1(_28_v3, 14)
			(_nw8).ArraySet1(_28_v3, 15)
			(_nw8).ArraySet1(_28_v3, 16)
			(_nw8).ArraySet1(_28_v3, 17)
			(_nw8).ArraySet1(_28_v3, 18)
			_51_v26 = _nw8
			var _52_v27 _dafny.Map
			_ = _52_v27
			_52_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v26, _25_v0)
			_52_v27 = (_52_v27).Update(_51_v26, p2)
		}
	} else {
		var _53_v28 _dafny.Map
		_ = _53_v28
		_53_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D3_.Create_DC8_(p3, p3, p0)).Dtor_cf18(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("slgt")).Cardinality()))
		var _54_v29 _dafny.MultiSet
		_ = _54_v29
		_54_v29 = _dafny.MultiSetOf(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v28, Companion_D1_.Create_DC4_(p1, p2, !(p2), false))).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(334), p3)), p2, p0, p0)
		_54_v29 = (_54_v29).Difference(_54_v29)
		if p2 {
			(globalState).F0 = _dafny.IntOfInt64(70)
			var _55_v30 *C0
			_ = _55_v30
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__(Companion_D3_.Create_DC8_(p3, Companion_Default___.Fm0(globalState), p2))
			_55_v30 = _nw9
			var _56_v31 _dafny.CodePoint
			_ = _56_v31
			_56_v31 = _dafny.CodePoint('c')
			var _57_v32 _dafny.Array
			_ = _57_v32
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw10
			_57_v32 = _nw10
			var _58_v33 _dafny.Map
			_ = _58_v33
			_58_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v31, _57_v32)
			_58_v33 = (_58_v33).Update(_56_v31, _57_v32)
			var _59_v34 bool
			_ = _59_v34
			_59_v34 = true
			_59_v34 = !((_59_v34) || (p2))
			var _60_v35 _dafny.Sequence
			_ = _60_v35
			_60_v35 = _dafny.SeqOf(!(p0), p0)
			var _61_v36 D4
			_ = _61_v36
			_61_v36 = Companion_D4_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v30).F14(), _dafny.SeqOf(_56_v31, _56_v31)))
			var _62_v37 _dafny.Sequence
			_ = _62_v37
			_62_v37 = _dafny.UnicodeSeqOfUtf8Bytes("osydgy")
			var _63_v38 _dafny.Map
			_ = _63_v38
			_63_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_55_v30).F14(), _62_v37)
			var _64_v39 _dafny.MultiSet
			_ = _64_v39
			_64_v39 = _dafny.MultiSetOf(p1, _dafny.IntOfUint32((Companion_Default___.Fm1(p0, (_dafny.Zero).Minus(p1), globalState)).Cardinality()), p1)
			var _pat_let_tv0 = _63_v38
			_ = _pat_let_tv0
			var _pat_let_tv1 = _63_v38
			_ = _pat_let_tv1
			var _pat_let_tv2 = _63_v38
			_ = _pat_let_tv2
			var _pat_let_tv3 = _63_v38
			_ = _pat_let_tv3
			var _pat_let_tv4 = _63_v38
			_ = _pat_let_tv4
			var _65_v40 _dafny.Array
			_ = _65_v40
			var _nwElement0_2 D4 = _61_v36
			_ = _nwElement0_2
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(20))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_2, 0)
			(_nw11).ArraySet1(_61_v36, 1)
			(_nw11).ArraySet1(Companion_Default___.Fm13(p3, _62_v37, _59_v34, globalState), 2)
			(_nw11).ArraySet1(_61_v36, 3)
			(_nw11).ArraySet1(_61_v36, 4)
			(_nw11).ArraySet1(_61_v36, 5)
			(_nw11).ArraySet1(Companion_D4_.Create_DC9_(_63_v38), 6)
			(_nw11).ArraySet1(_61_v36, 7)
			(_nw11).ArraySet1(_61_v36, 8)
			(_nw11).ArraySet1(Companion_D4_.Create_DC9_(_63_v38), 9)
			(_nw11).ArraySet1((func() D4 {
				if true {
					return _61_v36
				}
				return _61_v36
			})(), 10)
			(_nw11).ArraySet1(func(_pat_let0_0 D4) D4 {
				return func(_68_dt__update__tmp_h1 D4) D4 {
					return func(_pat_let3_0 _dafny.Map) D4 {
						return func(_69_dt__update_hcf20_h1 _dafny.Map) D4 {
							return Companion_D4_.Create_DC9_(_69_dt__update_hcf20_h1)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let0_0)
			}(func(_pat_let1_0 D4) D4 {
				return func(_66_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let2_0 _dafny.Map) D4 {
						return func(_67_dt__update_hcf20_h0 _dafny.Map) D4 {
							return Companion_D4_.Create_DC9_(_67_dt__update_hcf20_h0)
						}(_pat_let2_0)
					}(_pat_let_tv0)
				}(_pat_let1_0)
			}(_61_v36)), 11)
			(_nw11).ArraySet1(_61_v36, 12)
			(_nw11).ArraySet1(func(_pat_let4_0 D4) D4 {
				return func(_72_dt__update__tmp_h3 D4) D4 {
					return func(_pat_let7_0 _dafny.Map) D4 {
						return func(_73_dt__update_hcf20_h3 _dafny.Map) D4 {
							return Companion_D4_.Create_DC9_(_73_dt__update_hcf20_h3)
						}(_pat_let7_0)
					}(_pat_let_tv3)
				}(_pat_let4_0)
			}(func(_pat_let5_0 D4) D4 {
				return func(_70_dt__update__tmp_h2 D4) D4 {
					return func(_pat_let6_0 _dafny.Map) D4 {
						return func(_71_dt__update_hcf20_h2 _dafny.Map) D4 {
							return Companion_D4_.Create_DC9_(_71_dt__update_hcf20_h2)
						}(_pat_let6_0)
					}(_pat_let_tv2)
				}(_pat_let5_0)
			}(_61_v36)), 13)
			(_nw11).ArraySet1((func() D4 {
				if p0 {
					return func(_pat_let8_0 D4) D4 {
						return func(_74_dt__update__tmp_h4 D4) D4 {
							return func(_pat_let9_0 _dafny.Map) D4 {
								return func(_75_dt__update_hcf20_h4 _dafny.Map) D4 {
									return Companion_D4_.Create_DC9_(_75_dt__update_hcf20_h4)
								}(_pat_let9_0)
							}(_pat_let_tv4)
						}(_pat_let8_0)
					}(_61_v36)
				}
				return _61_v36
			})(), 14)
			(_nw11).ArraySet1(_61_v36, 15)
			(_nw11).ArraySet1(_61_v36, 16)
			(_nw11).ArraySet1(Companion_D4_.Create_DC9_(_63_v38), 17)
			(_nw11).ArraySet1(Companion_D4_.Create_DC9_(_63_v38), 18)
			(_nw11).ArraySet1(Companion_Default___.Fm13((_64_v39).Cardinality(), _62_v37, _59_v34, globalState), 19)
			_65_v40 = _nw11
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_65_v40), 0))
			_ = _index5
			(_65_v40).ArraySet1(_61_v36, (_index5).Int())
			var _76_v41 _dafny.Map
			_ = _76_v41
			_76_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm1(p0, _dafny.IntOfUint32((_62_v37).Cardinality()), globalState))
			var _77_v42 *C1
			_ = _77_v42
			var _nw12 *C1 = New_C1_()
			_ = _nw12
			_nw12.Ctor__()
			_77_v42 = _nw12
			var _78_v43 _dafny.Map
			_ = _78_v43
			_78_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _77_v42)
			var _79_v44 _dafny.Map
			_ = _79_v44
			_79_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_78_v43).Cardinality(), _59_v34)
			var _pat_let_tv5 = _63_v38
			_ = _pat_let_tv5
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_65_v40), 0))
			_ = _index6
			var _rhs0 _dafny.Sequence = Companion_Default___.Fm3(p1, Companion_Default___.SafeModuloInt(p1, p3), ((_76_v41).Cardinality()).Minus(p1), (func() _dafny.Map {
				if p0 {
					return _79_v44
				}
				return _79_v44
			})(), globalState)
			_ = _rhs0
			var _rhs1 D4 = (func() D4 {
				if p0 {
					return func(_pat_let10_0 D4) D4 {
						return func(_80_dt__update__tmp_h5 D4) D4 {
							return func(_pat_let11_0 _dafny.Map) D4 {
								return func(_81_dt__update_hcf20_h5 _dafny.Map) D4 {
									return Companion_D4_.Create_DC9_(_81_dt__update_hcf20_h5)
								}(_pat_let11_0)
							}(_pat_let_tv5)
						}(_pat_let10_0)
					}(_61_v36)
				}
				return _61_v36
			})()
			_ = _rhs1
			var _lhs0 _dafny.Array = _65_v40
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_65_v40), 0))
			_ = _lhs1
			_60_v35 = _rhs0
			(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
		} else {
			var _82_v45 bool
			_ = _82_v45
			_82_v45 = true
			var _83_v46 _dafny.Array
			_ = _83_v46
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw13
			_83_v46 = _nw13
			var _84_v47 _dafny.Array
			_ = _84_v47
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_0
			var _nw14 _dafny.Array
			_ = _nw14
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw14 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = (func(_85_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_86_i0 _dafny.Int) _dafny.Int {
						return (_86_i0).Minus((_dafny.Zero).Minus(_85_p3))
					}
				})(p3)
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
			_84_v47 = _nw14
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_83_v46), 0))
			_ = _index7
			(_83_v46).ArraySet1(_84_v47, (_index7).Int())
			var _87_v48 D3
			_ = _87_v48
			_87_v48 = Companion_D3_.Create_DC8_(p1, p3, p0)
			var _88_v49 *C0
			_ = _88_v49
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_87_v48)
			_88_v49 = _nw15
			var _89_v50 _dafny.Sequence
			_ = _89_v50
			_89_v50 = _dafny.SeqOf(_88_v49)
			var _90_v51 _dafny.Sequence
			_ = _90_v51
			_90_v51 = _dafny.SeqOf(p2, p2, p2)
			var _91_v52 _dafny.Set
			_ = _91_v52
			_91_v52 = _dafny.SetOf(p1, _dafny.IntOfUint32((_90_v51).Cardinality()))
			var _92_v53 _dafny.Sequence
			_ = _92_v53
			_92_v53 = _dafny.SeqOf(_91_v52)
			var _93_v54 D3
			_ = _93_v54
			_93_v54 = Companion_D3_.Create_DC8_(p3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_89_v50, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_92_v53).Cardinality()), _dafny.IntOfUint32((_89_v50).Cardinality()))).Uint32(), _88_v49)).Cardinality()), p0)
			var _94_v55 _dafny.Array
			_ = _94_v55
			var _nwElement0_3 bool = _82_v45
			_ = _nwElement0_3
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(12))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_3, 0)
			(_nw16).ArraySet1(p2, 1)
			(_nw16).ArraySet1(_82_v45, 2)
			(_nw16).ArraySet1(p2, 3)
			(_nw16).ArraySet1(_82_v45, 4)
			(_nw16).ArraySet1(p0, 5)
			(_nw16).ArraySet1(_82_v45, 6)
			(_nw16).ArraySet1(_82_v45, 7)
			(_nw16).ArraySet1((_87_v48).Dtor_cf19(), 8)
			(_nw16).ArraySet1(!(p2), 9)
			(_nw16).ArraySet1(p0, 10)
			(_nw16).ArraySet1(p0, 11)
			_94_v55 = _nw16
			var _95_v56 _dafny.Sequence
			_ = _95_v56
			_95_v56 = _dafny.SeqOf(_94_v55)
			var _96_v57 _dafny.Map
			_ = _96_v57
			_96_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v55, _94_v55)
			var _97_v58 *C1
			_ = _97_v58
			var _nw17 *C1 = New_C1_()
			_ = _nw17
			_nw17.Ctor__()
			_97_v58 = _nw17
			var _98_v59 _dafny.Map
			_ = _98_v59
			_98_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v58, _94_v55)
			var _99_v60 _dafny.Array
			_ = _99_v60
			var _nwElement0_4 _dafny.Array = (_95_v56).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_95_v56).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _nwElement0_4
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_4, 0)
			(_nw18).ArraySet1(_94_v55, 1)
			(_nw18).ArraySet1(_94_v55, 2)
			(_nw18).ArraySet1(_94_v55, 3)
			(_nw18).ArraySet1(_94_v55, 4)
			(_nw18).ArraySet1((func() _dafny.Array {
				if (_96_v57).Contains(_94_v55) {
					return (_96_v57).Get(_94_v55).(_dafny.Array)
				}
				return _94_v55
			})(), 5)
			(_nw18).ArraySet1(_94_v55, 6)
			(_nw18).ArraySet1(_94_v55, 7)
			(_nw18).ArraySet1(_94_v55, 8)
			(_nw18).ArraySet1(_94_v55, 9)
			(_nw18).ArraySet1((func() _dafny.Array {
				if true {
					return _94_v55
				}
				return _94_v55
			})(), 10)
			(_nw18).ArraySet1((func() _dafny.Array {
				if (_98_v59).Contains(_97_v58) {
					return (_98_v59).Get(_97_v58).(_dafny.Array)
				}
				return _94_v55
			})(), 11)
			_99_v60 = _nw18
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_99_v60), 0))
			_ = _index8
			(_99_v60).ArraySet1(_94_v55, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_83_v46), 0))
			_ = _index9
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_99_v60), 0))
			_ = _index10
			var _rhs2 bool = _82_v45
			_ = _rhs2
			var _rhs3 _dafny.Array = _84_v47
			_ = _rhs3
			var _rhs4 _dafny.Array = _94_v55
			_ = _rhs4
			var _rhs5 bool = p2
			_ = _rhs5
			var _lhs2 _dafny.Array = _83_v46
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(803), _dafny.ArrayLen((_83_v46), 0))
			_ = _lhs3
			var _lhs4 _dafny.Array = _99_v60
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_99_v60), 0))
			_ = _lhs5
			_82_v45 = _rhs2
			(_lhs2).ArraySet1(_rhs3, (_lhs3).Int())
			(_lhs4).ArraySet1(_rhs4, (_lhs5).Int())
			_82_v45 = _rhs5
			var _100_v61 _dafny.Array
			_ = _100_v61
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_1
			var _nw19 _dafny.Array
			_ = _nw19
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw19 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Set = (func(_101_v52 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_102_i1 _dafny.Int) _dafny.Set {
						return _101_v52
					}
				})(_91_v52)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw19 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw19).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw19).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_100_v61 = _nw19
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_100_v61), 0))
			_ = _index11
			(_100_v61).ArraySet1(_91_v52, (_index11).Int())
			var _103_v62 _dafny.CodePoint
			_ = _103_v62
			_103_v62 = _dafny.CodePoint('e')
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_100_v61), 0))
			_ = _index12
			(_100_v61).ArraySet1(((Companion_Default___.Fm14(_103_v62, (_dafny.Zero).Minus(p1), globalState)).Intersection(_91_v52)).Union(_91_v52), (_index12).Int())
			var _104_v63 _dafny.MultiSet
			_ = _104_v63
			_104_v63 = _dafny.MultiSetOf(p3)
			var _105_v64 _dafny.MultiSet
			_ = _105_v64
			_105_v64 = _dafny.MultiSetOf(Companion_Default___.Fm15((_dafny.Zero).Minus((_104_v63).Cardinality()), p1, globalState), Companion_D0_.Create_DC0_(_90_v51))
			var _106_v65 _dafny.Map
			_ = _106_v65
			_106_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_82_v45) || (p0), _105_v64)
			_106_v65 = (_106_v65).Update(!(p0), _105_v64)
			_82_v45 = (_90_v51).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_90_v51).Cardinality()))).Uint32()).(bool)
			var _107_v66 _dafny.Sequence
			_ = _107_v66
			_107_v66 = _dafny.SeqOf(_54_v29)
			_107_v66 = (func() _dafny.Sequence {
				if p0 {
					return _107_v66
				}
				return _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(p0)), _54_v29, _54_v29, Companion_Default___.Fm16(p3, p1, globalState), _54_v29)
			})()
		}
		var _108_v67 _dafny.Sequence
		_ = _108_v67
		_108_v67 = _dafny.UnicodeSeqOfUtf8Bytes("iwvnrgqkx")
		var _109_v68 _dafny.Map
		_ = _109_v68
		_109_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jgktoekf")).Cardinality()), true)
		(globalState).F0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_108_v67).Cardinality()), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(489), (_109_v68).Cardinality())))
		var _110_v69 _dafny.Map
		_ = _110_v69
		_110_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _54_v29)
		var _111_v70 _dafny.Sequence
		_ = _111_v70
		_111_v70 = _dafny.SeqOf(p0)
		_110_v69 = (_110_v69).Update(p3, _dafny.MultiSetFromSeq(_111_v70))
		var _112_v71 bool
		_ = _112_v71
		_112_v71 = true
		_112_v71 = (Companion_Default___.SafeModuloInt(p1, p3)).Cmp(p3) >= 0
	}
	(globalState).F0 = p1
	var _113_v72 _dafny.Sequence
	_ = _113_v72
	_113_v72 = _dafny.UnicodeSeqOfUtf8Bytes("oncu")
	var _114_v73 _dafny.Sequence
	_ = _114_v73
	_114_v73 = _dafny.SeqOf(_dafny.IntOfInt64(3), _dafny.IntOfInt64(236))
	var _hi0 _dafny.Int = (_dafny.Zero).Minus((_114_v73).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_114_v73).Cardinality()))).Uint32()).(_dafny.Int))
	_ = _hi0
	for _115_i2 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_113_v72).Cardinality())); _115_i2.Cmp(_hi0) < 0; _115_i2 = _115_i2.Plus(_dafny.One) {
		var _116_v74 _dafny.Map
		_ = _116_v74
		_116_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _113_v72)
		var _117_v75 _dafny.CodePoint
		_ = _117_v75
		_117_v75 = _dafny.CodePoint('i')
		var _118_v76 _dafny.Map
		_ = _118_v76
		_118_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v74, _117_v75)
		_118_v76 = (_118_v76).Update(_116_v74, _117_v75)
		_113_v72 = _113_v72
		(globalState).F0 = p3
		var _119_v77 _dafny.Array
		_ = _119_v77
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_2
		var _nw20 _dafny.Array
		_ = _nw20
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw20 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_120_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_121_i3 _dafny.Int) _dafny.Int {
					return (_121_i3).Minus(_120_p1)
				}
			})(p1)
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
		_119_v77 = _nw20
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_119_v77), 0))
		_ = _index13
		(_119_v77).ArraySet1(_115_i2, (_index13).Int())
		var _122_v78 _dafny.MultiSet
		_ = _122_v78
		_122_v78 = _dafny.MultiSetOf(p3, p3)
		var _123_v79 _dafny.Sequence
		_ = _123_v79
		_123_v79 = _dafny.SeqOf(_122_v78)
		var _124_v80 *C1
		_ = _124_v80
		var _nw21 *C1 = New_C1_()
		_ = _nw21
		_nw21.Ctor__()
		_124_v80 = _nw21
		var _125_v81 _dafny.Map
		_ = _125_v81
		_125_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v80, _124_v80)
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_119_v77), 0))
		_ = _index14
		(_119_v77).ArraySet1((_dafny.Zero).Minus((p3).Minus((func() _dafny.Int {
			if p0 {
				return _dafny.IntOfUint32((_123_v79).Cardinality())
			}
			return (_dafny.Zero).Minus((_125_v81).Cardinality())
		})())), (_index14).Int())
	}
	(globalState).F0 = p1
	var _126_v82 _dafny.Array
	_ = _126_v82
	var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
	_ = _nw22
	_126_v82 = _nw22
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_126_v82), 0))); ; {
		_guard_loop_0, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _127_i4 _dafny.Int
		_127_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_127_i4).Sign() != -1) && ((_127_i4).Cmp(_dafny.ArrayLen((_126_v82), 0)) < 0)) {
			(_126_v82).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_114_v73, _114_v73), _114_v73), (_127_i4).Int())
		}
	}
	var _128_v83 _dafny.Array
	_ = _128_v83
	var _nwElement0_5 bool = false
	_ = _nwElement0_5
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(27))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_5, 0)
	(_nw23).ArraySet1(p0, 1)
	(_nw23).ArraySet1(p2, 2)
	(_nw23).ArraySet1(!(p0), 3)
	(_nw23).ArraySet1(p2, 4)
	(_nw23).ArraySet1(p0, 5)
	(_nw23).ArraySet1(p0, 6)
	(_nw23).ArraySet1(p2, 7)
	(_nw23).ArraySet1(p2, 8)
	(_nw23).ArraySet1(p0, 9)
	(_nw23).ArraySet1(p0, 10)
	(_nw23).ArraySet1(p0, 11)
	(_nw23).ArraySet1(p0, 12)
	(_nw23).ArraySet1(p2, 13)
	(_nw23).ArraySet1(p2, 14)
	(_nw23).ArraySet1(p0, 15)
	(_nw23).ArraySet1(p2, 16)
	(_nw23).ArraySet1(Companion_Default___.Fm2(globalState), 17)
	(_nw23).ArraySet1(p0, 18)
	(_nw23).ArraySet1(p2, 19)
	(_nw23).ArraySet1(p0, 20)
	(_nw23).ArraySet1(p0, 21)
	(_nw23).ArraySet1(p2, 22)
	(_nw23).ArraySet1(p0, 23)
	(_nw23).ArraySet1(!(p0), 24)
	(_nw23).ArraySet1(p0, 25)
	(_nw23).ArraySet1(p0, 26)
	_128_v83 = _nw23
	var _129_v84 D4
	_ = _129_v84
	_129_v84 = Companion_D4_.Create_DC12_(p2, Companion_Default___.Fm17(p3, p2, false, globalState), p1, _128_v83, p2)
	var _130_v85 _dafny.Map
	_ = _130_v85
	_130_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_dafny.Zero).Minus(Companion_Default___.Fm0(globalState)))
	var _131_v86 _dafny.Sequence
	_ = _131_v86
	_131_v86 = _dafny.SeqOf(_130_v85)
	var _pat_let_tv6 = _128_v83
	_ = _pat_let_tv6
	var _pat_let_tv7 = _114_v73
	_ = _pat_let_tv7
	var _pat_let_tv8 = _131_v86
	_ = _pat_let_tv8
	var _132_v87 _dafny.Array
	_ = _132_v87
	var _nwElement0_6 D4 = func(_pat_let12_0 D4) D4 {
		return func(_133_dt__update__tmp_h6 D4) D4 {
			return func(_pat_let13_0 _dafny.Array) D4 {
				return func(_134_dt__update_hcf29_h0 _dafny.Array) D4 {
					return func(_pat_let14_0 _dafny.Int) D4 {
						return func(_135_dt__update_hcf28_h0 _dafny.Int) D4 {
							return func(_pat_let15_0 _dafny.Sequence) D4 {
								return func(_136_dt__update_hcf27_h0 _dafny.Sequence) D4 {
									return Companion_D4_.Create_DC12_((_133_dt__update__tmp_h6).Dtor_cf26(), _136_dt__update_hcf27_h0, _135_dt__update_hcf28_h0, _134_dt__update_hcf29_h0, (_133_dt__update__tmp_h6).Dtor_cf30())
								}(_pat_let15_0)
							}(_pat_let_tv8)
						}(_pat_let14_0)
					}(_dafny.IntOfUint32((_pat_let_tv7).Cardinality()))
				}(_pat_let13_0)
			}(_pat_let_tv6)
		}(_pat_let12_0)
	}(_129_v84)
	_ = _nwElement0_6
	var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
	_ = _nw24
	(_nw24).ArraySet1(_nwElement0_6, 0)
	_132_v87 = _nw24
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_132_v87), 0))
	_ = _index15
	(_132_v87).ArraySet1(_129_v84, (_index15).Int())
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_132_v87), 0))
	_ = _index16
	(_132_v87).ArraySet1(_129_v84, (_index16).Int())
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _137_v0 bool
	_ = _137_v0
	_137_v0 = true
	var _138_v1 _dafny.Set
	_ = _138_v1
	_138_v1 = _dafny.SetOf(_137_v0, _137_v0, _137_v0)
	var _139_v2 _dafny.Array
	_ = _139_v2
	var _len0_3 _dafny.Int = _dafny.One
	_ = _len0_3
	var _nw25 _dafny.Array
	_ = _nw25
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw25 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Sequence = (func(_140_v0 bool) func(_dafny.Int) _dafny.Sequence {
			return func(_141_i0 _dafny.Int) _dafny.Sequence {
				return (Companion_D0_.Create_DC0_(_dafny.SeqOf(_140_v0, _140_v0))).Dtor_cf0()
			}
		})(_137_v0)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw25 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw25).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw25).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_139_v2 = _nw25
	var _142_v3 _dafny.Int
	_ = _142_v3
	_142_v3 = _dafny.IntOfInt64(-730)
	var _143_v4 _dafny.Map
	_ = _143_v4
	_143_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _142_v3)
	var _144_v5 _dafny.Array
	_ = _144_v5
	var _nwElement0_7 _dafny.Map = _143_v4
	_ = _nwElement0_7
	var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(21))
	_ = _nw26
	(_nw26).ArraySet1(_nwElement0_7, 0)
	(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _142_v3), 1)
	(_nw26).ArraySet1(_143_v4, 2)
	(_nw26).ArraySet1(_143_v4, 3)
	(_nw26).ArraySet1(_143_v4, 4)
	(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _142_v3), 5)
	(_nw26).ArraySet1(_143_v4, 6)
	(_nw26).ArraySet1(_143_v4, 7)
	(_nw26).ArraySet1(_143_v4, 8)
	(_nw26).ArraySet1(_143_v4, 9)
	(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _142_v3), 10)
	(_nw26).ArraySet1(_143_v4, 11)
	(_nw26).ArraySet1(_143_v4, 12)
	(_nw26).ArraySet1(_143_v4, 13)
	(_nw26).ArraySet1(_143_v4, 14)
	(_nw26).ArraySet1(_143_v4, 15)
	(_nw26).ArraySet1(_143_v4, 16)
	(_nw26).ArraySet1(_143_v4, 17)
	(_nw26).ArraySet1(_143_v4, 18)
	(_nw26).ArraySet1(_143_v4, 19)
	(_nw26).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _142_v3), 20)
	_144_v5 = _nw26
	var _145_globalState *GlobalState
	_ = _145_globalState
	var _nw27 *GlobalState = New_GlobalState_()
	_ = _nw27
	_nw27.Ctor__(_dafny.IntOfInt64(317), (_dafny.SetOf(_137_v0)).Difference(_138_v1), _139_v2, _144_v5, _dafny.CodePoint('u'), _dafny.UnicodeSeqOfUtf8Bytes("tinqkjdyt"), _dafny.IntOfInt64(892), _dafny.IntOfInt64(437), true, _dafny.CodePoint('o'), false, _dafny.IntOfInt64(-34), false, false)
	_145_globalState = _nw27
	var _146_v6 _dafny.Sequence
	_ = _146_v6
	_146_v6 = _dafny.UnicodeSeqOfUtf8Bytes("whhg")
	var _hi1 _dafny.Int = _dafny.IntOfUint32((_146_v6).Cardinality())
	_ = _hi1
	for _147_i1 := (_138_v1).Cardinality(); _147_i1.Cmp(_hi1) < 0; _147_i1 = _147_i1.Plus(_dafny.One) {
		var _148_v7 _dafny.Array
		_ = _148_v7
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw28
		_148_v7 = _nw28
		var _149_v8 _dafny.Sequence
		_ = _149_v8
		_149_v8 = _dafny.SeqOf(_148_v7)
		var _150_v9 _dafny.Map
		_ = _150_v9
		_150_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_142_v3).Plus(_dafny.IntOfUint32((_146_v6).Cardinality())), _149_v8)
		_150_v9 = (_150_v9).Update((_dafny.Zero).Minus(Companion_Default___.Fm0(_145_globalState)), _149_v8)
		_142_v3 = (_dafny.Zero).Minus(_142_v3)
		var _151_v10 _dafny.Sequence
		_ = _151_v10
		_151_v10 = _dafny.SeqOf(_142_v3, _147_i1)
		if (Companion_Default___.SafeDivisionInt(_142_v3, ((_dafny.MultiSetFromSeq(_151_v10)).Update(_142_v3, Companion_Default___.Abs(_147_i1))).Cardinality())).Cmp(_142_v3) > 0 {
			var _152_v11 _dafny.Sequence
			_ = _152_v11
			_152_v11 = _dafny.SeqOf(_137_v0)
			var _153_v12 D0
			_ = _153_v12
			_153_v12 = Companion_D0_.Create_DC2_(_137_v0, _137_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_152_v11).Cardinality())))
			_142_v3 = (_153_v12).Dtor_cf4()
			var _154_v13 _dafny.Array
			_ = _154_v13
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw29
			_154_v13 = _nw29
			var _155_v14 _dafny.Sequence
			_ = _155_v14
			_155_v14 = _dafny.SeqOf(_154_v13)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_154_v13), 0))
			_ = _index17
			(_154_v13).ArraySet1(!_dafny.Companion_Sequence_.Equal(_155_v14, _155_v14), (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_154_v13), 0))
			_ = _index18
			(_154_v13).ArraySet1((_137_v0) && (_137_v0), (_index18).Int())
			var _156_v15 _dafny.Map
			_ = _156_v15
			_156_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_146_v6).Cardinality()), _137_v0)
			Companion_Default___.M0((_142_v3).Cmp(_147_i1) < 0, (_156_v15).Cardinality(), (_154_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_154_v13), 0))).Int()).(bool), _142_v3, _145_globalState)
			var _157_v16 _dafny.MultiSet
			_ = _157_v16
			_157_v16 = _dafny.MultiSetOf(_147_i1)
			var _158_v17 _dafny.Map
			_ = _158_v17
			_158_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_i1, Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_157_v16).Contains(_142_v3) {
					return (_157_v16).Multiplicity(_142_v3)
				}
				return _147_i1
			})(), _147_i1))
			var _159_v18 _dafny.CodePoint
			_ = _159_v18
			_159_v18 = _dafny.CodePoint('i')
			_158_v17 = (_158_v17).Update(_dafny.IntOfUint32((_146_v6).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_160_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})), (Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_161_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))).Cardinality()))).Uint32(), _159_v18)).Cardinality()))
			(_145_globalState).F0 = Companion_Default___.SafeModuloInt(_147_i1, _147_i1)
		} else {
			_142_v3 = _142_v3
			Companion_Default___.M0(!(_137_v0), _142_v3, (_142_v3).Cmp(_142_v3) > 0, _147_i1, _145_globalState)
			Companion_Default___.M0(_137_v0, _147_i1, false, (_142_v3).Times(_142_v3), _145_globalState)
			var _162_v19 _dafny.Map
			_ = _162_v19
			_162_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _146_v6)
			var _163_v20 _dafny.Map
			_ = _163_v20
			_163_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _dafny.Companion_Sequence_.IsProperPrefixOf(_146_v6, (func() _dafny.Sequence {
				if (_162_v19).Contains(_147_i1) {
					return (_162_v19).Get(_147_i1).(_dafny.Sequence)
				}
				return _146_v6
			})()))
			var _164_v21 _dafny.Map
			_ = _164_v21
			_164_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _137_v0)
			_137_v0 = (func() bool {
				if (_163_v20).Contains(_147_i1) {
					return (_163_v20).Get(_147_i1).(bool)
				}
				return (func() bool {
					if (_164_v21).Contains(_137_v0) {
						return (_164_v21).Get(_137_v0).(bool)
					}
					return _137_v0
				})()
			})()
			_163_v20 = _163_v20
		}
		var _165_v22 D0
		_ = _165_v22
		_165_v22 = Companion_D0_.Create_DC2_(_137_v0, _137_v0, _147_i1)
		var _166_v23 _dafny.Map
		_ = _166_v23
		_166_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, (_165_v22).Dtor_cf2())
		Companion_Default___.M0((func() bool {
			if (_166_v23).Contains(_137_v0) {
				return (_166_v23).Get(_137_v0).(bool)
			}
			return _137_v0
		})(), _142_v3, _137_v0, (_142_v3).Plus((_dafny.Zero).Minus(_142_v3)), _145_globalState)
	}
	Companion_Default___.M0((_142_v3).Cmp(_142_v3) <= 0, (_142_v3).Minus(_142_v3), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("bffs"), _146_v6), _142_v3, _145_globalState)
	var _hi2 _dafny.Int = (_dafny.Zero).Minus(_142_v3)
	_ = _hi2
	for _167_i3 := _142_v3; _167_i3.Cmp(_hi2) < 0; _167_i3 = _167_i3.Plus(_dafny.One) {
		_137_v0 = _137_v0
		Companion_Default___.M0(_137_v0, (_167_i3).Minus(Companion_Default___.Fm0(_145_globalState)), false, _142_v3, _145_globalState)
		var _168_v24 _dafny.Sequence
		_ = _168_v24
		_168_v24 = _dafny.SeqOf(_137_v0, _137_v0)
		var _169_v25 _dafny.Map
		_ = _169_v25
		_169_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_i3, _142_v3)
		var _170_v26 _dafny.Sequence
		_ = _170_v26
		_170_v26 = _dafny.SeqOf(_169_v25)
		var _171_v27 _dafny.CodePoint
		_ = _171_v27
		_171_v27 = _dafny.CodePoint('g')
		var _172_v28 _dafny.Sequence
		_ = _172_v28
		_172_v28 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(_137_v0, _142_v3, _145_globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_170_v26).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm1(_137_v0, _142_v3, _145_globalState)).Cardinality()))).Uint32(), _171_v27)).Cardinality()))
		var _173_v29 _dafny.Map
		_ = _173_v29
		_173_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_i3, !(_137_v0))
		Companion_Default___.M0(_dafny.Companion_Sequence_.Contains(_168_v24, _137_v0), (_172_v28).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_172_v28).Cardinality()))).Uint32()).(_dafny.Int), (func() bool {
			if (_173_v29).Contains(_167_i3) {
				return (_173_v29).Get(_167_i3).(bool)
			}
			return _137_v0
		})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_168_v24, (Companion_Default___.SafeIndex(_167_i3, _dafny.IntOfUint32((_168_v24).Cardinality()))).Uint32(), _137_v0)).Cardinality()), _145_globalState)
		var _174_v30 _dafny.Map
		_ = _174_v30
		_174_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, Companion_Default___.Fm2(_145_globalState))
		_142_v3 = ((_174_v30).Cardinality()).Times(_167_i3)
	}
	var _175_v31 D0
	_ = _175_v31
	_175_v31 = Companion_D0_.Create_DC1_((_142_v3).Plus(_142_v3))
	var _source1 D0 = _175_v31
	_ = _source1
	if _source1.Is_DC1() {
		var _176___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC1).Cf1
		_ = _176___mcc_h0
		var _177_cf1 _dafny.Int = _176___mcc_h0
		_ = _177_cf1
		_137_v0 = Companion_Default___.Fm2(_145_globalState)
		var _178_v32 _dafny.Sequence
		_ = _178_v32
		_178_v32 = _dafny.SeqOf(false)
		_178_v32 = _178_v32
		var _179_v33 _dafny.Map
		_ = _179_v33
		_179_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_cf1, _137_v0)
		var _180_v34 _dafny.Sequence
		_ = _180_v34
		_180_v34 = _dafny.SeqOf(_142_v3)
		_178_v32 = _dafny.Companion_Sequence_.Concatenate(_178_v32, _dafny.Companion_Sequence_.Concatenate(_178_v32, Companion_Default___.Fm3((_dafny.Zero).Minus(_177_cf1), _dafny.IntOfInt64(-816), _177_cf1, (_179_v33).Update(_dafny.IntOfUint32((_180_v34).Cardinality()), _137_v0), _145_globalState)))
		var _181_v35 D0
		_ = _181_v35
		_181_v35 = Companion_D0_.Create_DC2_(_137_v0, _137_v0, _142_v3)
		var _182_v36 _dafny.Array
		_ = _182_v36
		var _nwElement0_8 bool = _137_v0
		_ = _nwElement0_8
		var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(12))
		_ = _nw30
		(_nw30).ArraySet1(_nwElement0_8, 0)
		(_nw30).ArraySet1(!(_137_v0), 1)
		(_nw30).ArraySet1((_181_v35).Dtor_cf2(), 2)
		(_nw30).ArraySet1(_137_v0, 3)
		(_nw30).ArraySet1(_137_v0, 4)
		(_nw30).ArraySet1(_137_v0, 5)
		(_nw30).ArraySet1(_137_v0, 6)
		(_nw30).ArraySet1(_137_v0, 7)
		(_nw30).ArraySet1(false, 8)
		(_nw30).ArraySet1(_137_v0, 9)
		(_nw30).ArraySet1(!(!((_181_v35).Dtor_cf2())), 10)
		(_nw30).ArraySet1(_137_v0, 11)
		_182_v36 = _nw30
		var _183_v37 _dafny.MultiSet
		_ = _183_v37
		_183_v37 = _dafny.MultiSetOf(_182_v36, _182_v36, _182_v36)
		_137_v0 = !(!((_183_v37).Contains(_182_v36))) || (_137_v0)
	} else if _source1.Is_DC2() {
		var _184___mcc_h1 bool = _source1.Get_().(D0_DC2).Cf2
		_ = _184___mcc_h1
		var _185___mcc_h2 bool = _source1.Get_().(D0_DC2).Cf3
		_ = _185___mcc_h2
		var _186___mcc_h3 _dafny.Int = _source1.Get_().(D0_DC2).Cf4
		_ = _186___mcc_h3
		var _187_cf4 _dafny.Int = _186___mcc_h3
		_ = _187_cf4
		var _188_cf3 bool = _185___mcc_h2
		_ = _188_cf3
		var _189_cf2 bool = _184___mcc_h1
		_ = _189_cf2
		var _190_v38 _dafny.Sequence
		_ = _190_v38
		_190_v38 = _dafny.SeqOf(_137_v0, true, _137_v0)
		Companion_Default___.M0(_dafny.Companion_Sequence_.IsPrefixOf(_190_v38, _190_v38), _142_v3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_cf3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_146_v6).Cardinality())))).Contains(_188_cf3), _142_v3, _145_globalState)
		var _191_v39 *C1
		_ = _191_v39
		var _nw31 *C1 = New_C1_()
		_ = _nw31
		_nw31.Ctor__()
		_191_v39 = _nw31
		var _192_v40 _dafny.Array
		_ = _192_v40
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_4
		var _nw32 _dafny.Array
		_ = _nw32
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw32 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Map = (func(_193_v0 bool, _194_cf3 bool, _195_cf2 bool) func(_dafny.Int) _dafny.Map {
				return func(_196_i4 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v0, _194_cf3)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC2_(_193_v0, _195_cf2, (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.CodePoint('y'))).Cardinality())).Cardinality())).Dtor_cf3(), _194_cf3))
				}
			})(_137_v0, _188_cf3, _189_cf2)
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
		_192_v40 = _nw32
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_192_v40), 0))
		_ = _index19
		(_192_v40).ArraySet1(Companion_Default___.Fm7(true, _145_globalState), (_index19).Int())
		var _197_v41 _dafny.Map
		_ = _197_v41
		_197_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_cf3, _137_v0)
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_192_v40), 0))
		_ = _index20
		(_192_v40).ArraySet1((_197_v41).Merge(Companion_Default___.Fm7(_137_v0, _145_globalState)), (_index20).Int())
		var _198_v42 _dafny.Array
		_ = _198_v42
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_5
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) bool = (func(_199_cf3 bool) func(_dafny.Int) bool {
				return func(_200_i5 _dafny.Int) bool {
					return _199_cf3
				}
			})(_188_cf3)
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
		_198_v42 = _nw33
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_198_v42), 0))
		_ = _index21
		(_198_v42).ArraySet1(_188_cf3, (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_198_v42), 0))
		_ = _index22
		(_198_v42).ArraySet1((_142_v3).Cmp(_187_cf4) > 0, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_198_v42), 0))
		_ = _index23
		(_198_v42).ArraySet1((_187_cf4).Cmp(_142_v3) < 0, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_198_v42), 0))
		_ = _index24
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_198_v42), 0))
		_ = _index25
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_198_v42), 0))
		_ = _index26
		var _rhs6 bool = (_187_cf4).Cmp(_187_cf4) <= 0
		_ = _rhs6
		var _rhs7 bool = _188_cf3
		_ = _rhs7
		var _rhs8 bool = false
		_ = _rhs8
		var _rhs9 bool = _188_cf3
		_ = _rhs9
		var _lhs6 _dafny.Array = _198_v42
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_198_v42), 0))
		_ = _lhs7
		var _lhs8 _dafny.Array = _198_v42
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_198_v42), 0))
		_ = _lhs9
		var _lhs10 _dafny.Array = _198_v42
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_198_v42), 0))
		_ = _lhs11
		(_lhs6).ArraySet1(_rhs6, (_lhs7).Int())
		_189_cf2 = _rhs7
		(_lhs8).ArraySet1(_rhs8, (_lhs9).Int())
		(_lhs10).ArraySet1(_rhs9, (_lhs11).Int())
	} else {
		var _201___mcc_h4 _dafny.Sequence = _source1.Get_().(D0_DC0).Cf0
		_ = _201___mcc_h4
		var _202_cf0 _dafny.Sequence = _201___mcc_h4
		_ = _202_cf0
		Companion_Default___.M0(_137_v0, _142_v3, _137_v0, (_142_v3).Plus(_142_v3), _145_globalState)
		var _203_v43 _dafny.MultiSet
		_ = _203_v43
		_203_v43 = _dafny.MultiSetOf(_143_v4)
		Companion_Default___.M0(_137_v0, ((_203_v43).Cardinality()).Minus(_142_v3), _137_v0, _dafny.IntOfUint32((_146_v6).Cardinality()), _145_globalState)
		var _204_v44 _dafny.Array
		_ = _204_v44
		var _nwElement0_9 bool = _137_v0
		_ = _nwElement0_9
		var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(15))
		_ = _nw34
		(_nw34).ArraySet1(_nwElement0_9, 0)
		(_nw34).ArraySet1(_137_v0, 1)
		(_nw34).ArraySet1(_137_v0, 2)
		(_nw34).ArraySet1(_137_v0, 3)
		(_nw34).ArraySet1(_137_v0, 4)
		(_nw34).ArraySet1(_137_v0, 5)
		(_nw34).ArraySet1(_137_v0, 6)
		(_nw34).ArraySet1(Companion_Default___.Fm2(_145_globalState), 7)
		(_nw34).ArraySet1(_137_v0, 8)
		(_nw34).ArraySet1(_137_v0, 9)
		(_nw34).ArraySet1(_137_v0, 10)
		(_nw34).ArraySet1(_137_v0, 11)
		(_nw34).ArraySet1(_137_v0, 12)
		(_nw34).ArraySet1(_137_v0, 13)
		(_nw34).ArraySet1(_137_v0, 14)
		_204_v44 = _nw34
		var _205_v45 _dafny.Map
		_ = _205_v45
		_205_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v44, _142_v3)
		var _206_v46 _dafny.Map
		_ = _206_v46
		_206_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_205_v45).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v44, _142_v3)), _dafny.Companion_Sequence_.Concatenate(_146_v6, _146_v6))
		_206_v46 = (_206_v46).Update((_205_v45).Merge(_205_v45), _dafny.UnicodeSeqOfUtf8Bytes("jbs"))
		var _207_v47 _dafny.CodePoint
		_ = _207_v47
		_207_v47 = _dafny.CodePoint('j')
		(_145_globalState).F9 = _207_v47
	}
	_137_v0 = _137_v0
	var _208_v48 D1
	_ = _208_v48
	_208_v48 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(423), _137_v0, _137_v0, _137_v0)
	var _pat_let_tv9 = _142_v3
	_ = _pat_let_tv9
	var _pat_let_tv10 = _143_v4
	_ = _pat_let_tv10
	if func(_source2 D1) bool {
		if _source2.Is_DC4() {
			var _209___mcc_h5 _dafny.Int = _source2.Get_().(D1_DC4).Cf6
			_ = _209___mcc_h5
			var _210___mcc_h6 bool = _source2.Get_().(D1_DC4).Cf7
			_ = _210___mcc_h6
			var _211___mcc_h7 bool = _source2.Get_().(D1_DC4).Cf8
			_ = _211___mcc_h7
			var _212___mcc_h8 bool = _source2.Get_().(D1_DC4).Cf9
			_ = _212___mcc_h8
			var _213_cf9 bool = _212___mcc_h8
			_ = _213_cf9
			var _214_cf8 bool = _211___mcc_h7
			_ = _214_cf8
			var _215_cf7 bool = _210___mcc_h6
			_ = _215_cf7
			var _216_cf6 _dafny.Int = _209___mcc_h5
			_ = _216_cf6
			return !((_dafny.SetOf(_pat_let_tv9, (_pat_let_tv10).Cardinality())).IsSubsetOf(_dafny.SetOf(_216_cf6)))
		} else {
			var _217___mcc_h9 _dafny.Map = _source2.Get_().(D1_DC3).Cf5
			_ = _217___mcc_h9
			var _218_cf5 _dafny.Map = _217___mcc_h9
			_ = _218_cf5
			return false
		}
	}(_208_v48) {
		var _219_v49 _dafny.Sequence
		_ = _219_v49
		_219_v49 = _dafny.SeqOf(_142_v3, (func() _dafny.Int {
			if (_143_v4).Contains(_137_v0) {
				return (_143_v4).Get(_137_v0).(_dafny.Int)
			}
			return _142_v3
		})(), _142_v3, _dafny.IntOfInt64(390), _142_v3)
		var _rhs10 _dafny.Int = Companion_Default___.SafeDivisionInt(_142_v3, _dafny.IntOfUint32((_219_v49).Cardinality()))
		_ = _rhs10
		var _rhs11 _dafny.Int = _142_v3
		_ = _rhs11
		var _rhs12 _dafny.Int = _142_v3
		_ = _rhs12
		var _lhs12 *GlobalState = _145_globalState
		_ = _lhs12
		var _lhs13 *GlobalState = _145_globalState
		_ = _lhs13
		_lhs12.F0 = _rhs10
		_lhs13.F0 = _rhs11
		_142_v3 = _rhs12
		var _220_v50 _dafny.Array
		_ = _220_v50
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw35
		_220_v50 = _nw35
		var _221_v51 _dafny.Sequence
		_ = _221_v51
		_221_v51 = _dafny.SeqOf(_220_v50, _220_v50, _220_v50)
		_221_v51 = _221_v51
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))
		_ = _index27
		(_220_v50).ArraySet1(_142_v3, (_index27).Int())
		var _222_v52 _dafny.MultiSet
		_ = _222_v52
		_222_v52 = _dafny.MultiSetOf(_137_v0)
		var _223_v53 _dafny.Map
		_ = _223_v53
		_223_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_145_globalState), _142_v3)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))
		_ = _index28
		(_220_v50).ArraySet1((func() _dafny.Int {
			if (_222_v52).Contains((func() bool {
				if _137_v0 {
					return _137_v0
				}
				return _137_v0
			})()) {
				return (_222_v52).Multiplicity((func() bool {
					if _137_v0 {
						return _137_v0
					}
					return _137_v0
				})())
			}
			return (_223_v53).Cardinality()
		})(), (_index28).Int())
		if Companion_Default___.Fm2(_145_globalState) {
			var _224_v54 _dafny.Map
			_ = _224_v54
			_224_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _137_v0)
			_224_v54 = (Companion_Default___.Fm8(_142_v3, _145_globalState)).Merge(Companion_Default___.Fm8(_dafny.IntOfInt64(473), _145_globalState))
			var _225_v55 _dafny.Array
			_ = _225_v55
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_6
			var _nw36 _dafny.Array
			_ = _nw36
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw36 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.CodePoint = func(_226_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw36 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw36).ArraySet1CodePoint(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw36).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_225_v55 = _nw36
			var _227_v56 _dafny.CodePoint
			_ = _227_v56
			_227_v56 = _dafny.CodePoint('k')
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_225_v55), 0))
			_ = _index29
			(_225_v55).ArraySet1CodePoint(_227_v56, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_225_v55), 0))
			_ = _index30
			(_225_v55).ArraySet1CodePoint(_dafny.CodePoint('g'), (_index30).Int())
			(_145_globalState).F0 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_220_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))).Int()).(_dafny.Int), _142_v3), (_220_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))).Int()).(_dafny.Int))
			(_145_globalState).F0 = _142_v3
			var _228_v57 _dafny.Array
			_ = _228_v57
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw37
			_228_v57 = _nw37
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_228_v57), 0))
			_ = _index31
			(_228_v57).ArraySet1(_dafny.SeqOf(_223_v53, _223_v53, _223_v53), (_index31).Int())
			var _229_v58 _dafny.Sequence
			_ = _229_v58
			_229_v58 = _dafny.SeqOf(_223_v53, _223_v53)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_228_v57), 0))
			_ = _index32
			(_228_v57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_229_v58, _229_v58), (_index32).Int())
		} else {
			var _230_v59 _dafny.Map
			_ = _230_v59
			_230_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_137_v0)), _dafny.UnicodeSeqOfUtf8Bytes("jlxmbhbd"))
			_230_v59 = (_230_v59).Merge((_230_v59).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _146_v6)))
			var _231_v60 _dafny.Set
			_ = _231_v60
			_231_v60 = _dafny.SetOf(_142_v3)
			_231_v60 = _231_v60
			var _232_v61 _dafny.Map
			_ = _232_v61
			_232_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_142_v3, _142_v3, _137_v0), _137_v0)
			var _233_v63 _dafny.Map
			_ = _233_v63
			_233_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm9(_145_globalState), _146_v6)
			var _234_v64 _dafny.Set
			_ = _234_v64
			_234_v64 = _dafny.SetOf(Companion_D1_.Create_DC3_(_233_v63))
			var _235_v65 D3
			_ = _235_v65
			_235_v65 = Companion_D3_.Create_DC8_((_220_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))).Int()).(_dafny.Int), (_234_v64).Cardinality(), _137_v0)
			var _236_v66 _dafny.Array
			_ = _236_v66
			var _nwElement0_10 _dafny.Map = _232_v61
			_ = _nwElement0_10
			var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(2))
			_ = _nw38
			(_nw38).ArraySet1(_nwElement0_10, 0)
			(_nw38).ArraySet1(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter5 := _dafny.Iterate((_dafny.SeqOf(_235_v65)).Elements()); ; {
					_compr_4, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _237_v62 D3
					_237_v62 = interface{}(_compr_4).(D3)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_235_v65), _237_v62) {
						_coll4.Add(_237_v62, _137_v0)
					}
				}
				return _coll4.ToMap()
			}(), 1)
			_236_v66 = _nw38
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_236_v66), 0))
			_ = _index33
			(_236_v66).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v65, true), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_236_v66), 0))
			_ = _index34
			(_236_v66).ArraySet1((_232_v61).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v65, Companion_Default___.Fm2(_145_globalState))), (_index34).Int())
			_137_v0 = true
			var _238_v67 D0
			_ = _238_v67
			_238_v67 = Companion_D0_.Create_DC2_(_137_v0, _137_v0, (_220_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))).Int()).(_dafny.Int))
			var _pat_let_tv11 = _137_v0
			_ = _pat_let_tv11
			var _239_v68 _dafny.Array
			_ = _239_v68
			var _nwElement0_11 D0 = _238_v67
			_ = _nwElement0_11
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(11))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_11, 0)
			(_nw39).ArraySet1(_238_v67, 1)
			(_nw39).ArraySet1(_238_v67, 2)
			(_nw39).ArraySet1(_238_v67, 3)
			(_nw39).ArraySet1(Companion_D0_.Create_DC2_(_137_v0, _137_v0, Companion_Default___.Fm0(_145_globalState)), 4)
			(_nw39).ArraySet1(_238_v67, 5)
			(_nw39).ArraySet1(_238_v67, 6)
			(_nw39).ArraySet1(func(_pat_let16_0 D0) D0 {
				return func(_240_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let17_0 bool) D0 {
						return func(_241_dt__update_hcf2_h0 bool) D0 {
							return Companion_D0_.Create_DC2_(_241_dt__update_hcf2_h0, (_240_dt__update__tmp_h0).Dtor_cf3(), (_240_dt__update__tmp_h0).Dtor_cf4())
						}(_pat_let17_0)
					}(_pat_let_tv11)
				}(_pat_let16_0)
			}(_238_v67), 7)
			(_nw39).ArraySet1(_238_v67, 8)
			(_nw39).ArraySet1(_238_v67, 9)
			(_nw39).ArraySet1(func(_pat_let18_0 D0) D0 {
				return func(_242_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let19_0 bool) D0 {
						return func(_243_dt__update_hcf3_h0 bool) D0 {
							return Companion_D0_.Create_DC2_((_242_dt__update__tmp_h1).Dtor_cf2(), _243_dt__update_hcf3_h0, (_242_dt__update__tmp_h1).Dtor_cf4())
						}(_pat_let19_0)
					}(false)
				}(_pat_let18_0)
			}(Companion_Default___.Fm10(_137_v0, _dafny.UnicodeSeqOfUtf8Bytes("oysattfu"), _145_globalState)), 10)
			_239_v68 = _nw39
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_239_v68), 0))
			_ = _index35
			(_239_v68).ArraySet1(_238_v67, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_239_v68), 0))
			_ = _index36
			var _rhs13 _dafny.Int = ((_138_v1).Cardinality()).Plus(_142_v3)
			_ = _rhs13
			var _rhs14 _dafny.CodePoint = _dafny.CodePoint('e')
			_ = _rhs14
			var _rhs15 D0 = Companion_Default___.Fm10(_137_v0, _146_v6, _145_globalState)
			_ = _rhs15
			var _rhs16 _dafny.Int = (_dafny.Zero).Minus((_220_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_220_v50), 0))).Int()).(_dafny.Int))
			_ = _rhs16
			var _lhs14 *GlobalState = _145_globalState
			_ = _lhs14
			var _lhs15 *GlobalState = _145_globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _239_v68
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_239_v68), 0))
			_ = _lhs17
			_lhs14.F0 = _rhs13
			_lhs15.F4 = _rhs14
			(_lhs16).ArraySet1(_rhs15, (_lhs17).Int())
			_142_v3 = _rhs16
		}
		(_145_globalState).F0 = _142_v3
	} else {
		(_145_globalState).F0 = _dafny.IntOfInt64(660)
		var _244_v69 _dafny.Array
		_ = _244_v69
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw40
		_244_v69 = _nw40
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_244_v69), 0))
		_ = _index37
		(_244_v69).ArraySet1(!(_137_v0), (_index37).Int())
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_244_v69), 0))
		_ = _index38
		(_244_v69).ArraySet1((_142_v3).Cmp(_142_v3) == 0, (_index38).Int())
		var _245_v70 _dafny.Sequence
		_ = _245_v70
		_245_v70 = _dafny.SeqOf(_146_v6)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_244_v69), 0))
		_ = _index39
		var _rhs17 bool = (_142_v3).Cmp(_142_v3) != 0
		_ = _rhs17
		var _rhs18 bool = !((_142_v3).Cmp(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_145_globalState), _dafny.IntOfUint32((_245_v70).Cardinality()))) != 0)
		_ = _rhs18
		var _lhs18 _dafny.Array = _244_v69
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_244_v69), 0))
		_ = _lhs19
		_137_v0 = _rhs17
		(_lhs18).ArraySet1(_rhs18, (_lhs19).Int())
		_137_v0 = (_142_v3).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(486))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_246_v3 _dafny.Int, _247_v0 bool, _248_v69 _dafny.Array) func(_dafny.Int) _dafny.Map {
			return func(_249_i7 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_246_v3, (_dafny.MultiSetOf(_247_v0, (_248_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_248_v69), 0))).Int()).(bool))).Cardinality())
			}
		})(_142_v3, _137_v0, _244_v69)))).Cardinality())) <= 0
		var _250_v71 _dafny.Sequence
		_ = _250_v71
		_250_v71 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v6, _244_v69))
		(_145_globalState).F0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_142_v3), ((_250_v71).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(576))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_251_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()), _dafny.IntOfUint32((_250_v71).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
	}
	var _252_v72 _dafny.Map
	_ = _252_v72
	_252_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, _137_v0)
	_252_v72 = (_252_v72).Merge(_252_v72)
	var _253_v73 D0
	_ = _253_v73
	_253_v73 = Companion_D0_.Create_DC2_(_137_v0, _137_v0, _142_v3)
	var _pat_let_tv12 = _137_v0
	_ = _pat_let_tv12
	var _pat_let_tv13 = _142_v3
	_ = _pat_let_tv13
	_142_v3 = func(_source3 D0) _dafny.Int {
		if _source3.Is_DC1() {
			var _254___mcc_h10 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
			_ = _254___mcc_h10
			var _255_cf1 _dafny.Int = _254___mcc_h10
			_ = _255_cf1
			return _255_cf1
		} else if _source3.Is_DC2() {
			var _256___mcc_h11 bool = _source3.Get_().(D0_DC2).Cf2
			_ = _256___mcc_h11
			var _257___mcc_h12 bool = _source3.Get_().(D0_DC2).Cf3
			_ = _257___mcc_h12
			var _258___mcc_h13 _dafny.Int = _source3.Get_().(D0_DC2).Cf4
			_ = _258___mcc_h13
			var _259_cf4 _dafny.Int = _258___mcc_h13
			_ = _259_cf4
			var _260_cf3 bool = _257___mcc_h12
			_ = _260_cf3
			var _261_cf2 bool = _256___mcc_h11
			_ = _261_cf2
			return (_dafny.MultiSetOf(_pat_let_tv12)).Cardinality()
		} else {
			var _262___mcc_h14 _dafny.Sequence = _source3.Get_().(D0_DC0).Cf0
			_ = _262___mcc_h14
			var _263_cf0 _dafny.Sequence = _262___mcc_h14
			_ = _263_cf0
			return (_dafny.IntOfInt64(667)).Times(_pat_let_tv13)
		}
	}(_253_v73)
	_137_v0 = _137_v0
	var _264_v74 _dafny.Map
	_ = _264_v74
	_264_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_208_v48, _146_v6)
	var _265_v75 D2
	_ = _265_v75
	_265_v75 = Companion_D2_.Create_DC6_(_dafny.Companion_Sequence_.Concatenate(_146_v6, (func() _dafny.Sequence {
		if (_264_v74).Contains(Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_dafny.IntOfUint32((_146_v6).Cardinality())), _137_v0, _137_v0, _137_v0)) {
			return (_264_v74).Get(Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_dafny.IntOfUint32((_146_v6).Cardinality())), _137_v0, _137_v0, _137_v0)).(_dafny.Sequence)
		}
		return _146_v6
	})()), _142_v3, _137_v0, _142_v3, _137_v0)
	_265_v75 = _265_v75
	var _266_v76 _dafny.Map
	_ = _266_v76
	_266_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v3, (Companion_D0_.Create_DC1_(_142_v3)).Dtor_cf1())
	var _267_v77 _dafny.Sequence
	_ = _267_v77
	_267_v77 = _dafny.SeqOf(_142_v3)
	var _268_v78 _dafny.Sequence
	_ = _268_v78
	_268_v78 = _dafny.SeqOf(_266_v76, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_146_v6).Cardinality()), _dafny.IntOfUint32((_267_v77).Cardinality())), _266_v76, _266_v76, _266_v76)
	var _269_v79 _dafny.Sequence
	_ = _269_v79
	_269_v79 = _dafny.SeqOf(_137_v0)
	var _270_v80 D3
	_ = _270_v80
	_270_v80 = Companion_D3_.Create_DC8_(_142_v3, _142_v3, false)
	var _271_v81 *C0
	_ = _271_v81
	var _nw41 *C0 = New_C0_()
	_ = _nw41
	_nw41.Ctor__(_270_v80)
	_271_v81 = _nw41
	var _272_v82 _dafny.Sequence
	_ = _272_v82
	_272_v82 = _dafny.SeqOf(_271_v81, _271_v81)
	var _273_v83 _dafny.MultiSet
	_ = _273_v83
	_273_v83 = _dafny.MultiSetOf((_272_v82).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_272_v82).Cardinality()))).Uint32()).(*C0))
	var _274_v84 _dafny.MultiSet
	_ = _274_v84
	_274_v84 = _dafny.MultiSetOf(_142_v3, (_273_v83).Cardinality(), (_dafny.Zero).Minus(_142_v3))
	var _275_v85 _dafny.Array
	_ = _275_v85
	var _nwElement0_12 bool = _137_v0
	_ = _nwElement0_12
	var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(21))
	_ = _nw42
	(_nw42).ArraySet1(_nwElement0_12, 0)
	(_nw42).ArraySet1(_137_v0, 1)
	(_nw42).ArraySet1(true, 2)
	(_nw42).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_268_v78, _268_v78), 3)
	(_nw42).ArraySet1(_137_v0, 4)
	(_nw42).ArraySet1(!((_138_v1).Contains(_137_v0)), 5)
	(_nw42).ArraySet1(_137_v0, 6)
	(_nw42).ArraySet1(!(_137_v0), 7)
	(_nw42).ArraySet1(_137_v0, 8)
	(_nw42).ArraySet1(_137_v0, 9)
	(_nw42).ArraySet1(_137_v0, 10)
	(_nw42).ArraySet1(((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ryigcb"), _146_v6, _146_v6)).Cardinality()).Cmp(_142_v3) == 0, 11)
	(_nw42).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfUint32((_269_v79).Cardinality()))).IsProperSubsetOf(_274_v84), 12)
	(_nw42).ArraySet1((_137_v0) || (_137_v0), 13)
	(_nw42).ArraySet1(_137_v0, 14)
	(_nw42).ArraySet1(_137_v0, 15)
	(_nw42).ArraySet1(_137_v0, 16)
	(_nw42).ArraySet1((_137_v0) || (_137_v0), 17)
	(_nw42).ArraySet1(_137_v0, 18)
	(_nw42).ArraySet1(!(_137_v0), 19)
	(_nw42).ArraySet1(!(_137_v0), 20)
	_275_v85 = _nw42
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))
	_ = _index40
	(_275_v85).ArraySet1(_137_v0, (_index40).Int())
	var _276_v86 _dafny.CodePoint
	_ = _276_v86
	_276_v86 = _dafny.CodePoint('d')
	var _277_v87 _dafny.Map
	_ = _277_v87
	_277_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_v86, (_dafny.Zero).Minus((_271_v81).Fm6(_137_v0, (Companion_D2_.Create_DC6_(_146_v6, _142_v3, (func() bool {
		if (_252_v72).Contains(_142_v3) {
			return (_252_v72).Get(_142_v3).(bool)
		}
		return _137_v0
	})(), _142_v3, (func() bool {
		if (_252_v72).Contains(_dafny.IntOfInt64(540)) {
			return (_252_v72).Get(_dafny.IntOfInt64(540)).(bool)
		}
		return _137_v0
	})())).Dtor_cf12(), !(_137_v0), _145_globalState)))
	var _278_v88 _dafny.MultiSet
	_ = _278_v88
	_278_v88 = _dafny.MultiSetOf(_137_v0)
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))
	_ = _index41
	(_275_v85).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lthcc")).Cardinality())).Cmp((func() _dafny.Int {
		if (_277_v87).Contains(_276_v86) {
			return (_277_v87).Get(_276_v86).(_dafny.Int)
		}
		return (func() _dafny.Int {
			if (_278_v88).Contains(_137_v0) {
				return (_278_v88).Multiplicity(_137_v0)
			}
			return _142_v3
		})()
	})()) > 0, (_index41).Int())
	Companion_Default___.M0(_137_v0, (_142_v3).Times(_142_v3), _137_v0, _142_v3, _145_globalState)
	var _279_v89 _dafny.Map
	_ = _279_v89
	_279_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_v80, _146_v6)
	var _280_v90 D4
	_ = _280_v90
	_280_v90 = Companion_D4_.Create_DC9_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_271_v81).F14(), _146_v6))
	_279_v89 = (_280_v90).Dtor_cf20()
	var _281_v91 _dafny.MultiSet
	_ = _281_v91
	_281_v91 = _dafny.MultiSetOf(_269_v79)
	var _rhs19 _dafny.Int = ((_281_v91).Union(_dafny.MultiSetOf(_dafny.SeqOf(_137_v0), _269_v79))).Cardinality()
	_ = _rhs19
	var _rhs20 _dafny.Sequence = _267_v77
	_ = _rhs20
	var _rhs21 _dafny.Int = (Companion_Default___.Fm11(_276_v86, _145_globalState)).Dtor_cf14()
	_ = _rhs21
	var _lhs20 *GlobalState = _145_globalState
	_ = _lhs20
	var _lhs21 *GlobalState = _145_globalState
	_ = _lhs21
	_lhs20.F0 = _rhs19
	_267_v77 = _rhs20
	_lhs21.F0 = _rhs21
	var _282_i9 _dafny.Int
	_ = _282_i9
	_282_i9 = _dafny.Zero
	{
		for _137_v0 {
			{
				if (_282_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_282_i9 = (_282_i9).Plus(_dafny.One)
				Companion_Default___.M0((_275_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))).Int()).(bool), _142_v3, (_275_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))).Int()).(bool), _142_v3, _145_globalState)
				var _283_v92 D0
				_ = _283_v92
				_283_v92 = Companion_D0_.Create_DC0_(_dafny.SeqOf((_275_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))).Int()).(bool), _137_v0))
				var _source4 D0 = _283_v92
				_ = _source4
				if _source4.Is_DC1() {
					var _284___mcc_h15 _dafny.Int = _source4.Get_().(D0_DC1).Cf1
					_ = _284___mcc_h15
					var _285_cf1 _dafny.Int = _284___mcc_h15
					_ = _285_cf1
					var _286_v93 _dafny.Array
					_ = _286_v93
					var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(21))
					_ = _nw43
					_286_v93 = _nw43
					var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_286_v93), 0))
					_ = _index42
					(_286_v93).ArraySet1(_274_v84, (_index42).Int())
					var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))
					_ = _index43
					var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_286_v93), 0))
					_ = _index44
					var _rhs22 _dafny.Int = _285_cf1
					_ = _rhs22
					var _rhs23 bool = _137_v0
					_ = _rhs23
					var _rhs24 _dafny.CodePoint = _dafny.CodePoint('o')
					_ = _rhs24
					var _rhs25 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_267_v77, _267_v77))).Intersection(Companion_Default___.Fm12(_276_v86, (_275_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))).Int()).(bool), _285_cf1, _145_globalState))
					_ = _rhs25
					var _rhs26 _dafny.Sequence = _272_v82
					_ = _rhs26
					var _lhs22 _dafny.Array = _275_v85
					_ = _lhs22
					var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))
					_ = _lhs23
					var _lhs24 *GlobalState = _145_globalState
					_ = _lhs24
					var _lhs25 _dafny.Array = _286_v93
					_ = _lhs25
					var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_286_v93), 0))
					_ = _lhs26
					_142_v3 = _rhs22
					(_lhs22).ArraySet1(_rhs23, (_lhs23).Int())
					_lhs24.F9 = _rhs24
					(_lhs25).ArraySet1(_rhs25, (_lhs26).Int())
					_272_v82 = _rhs26
					var _287_v94 *C1
					_ = _287_v94
					var _nw44 *C1 = New_C1_()
					_ = _nw44
					_nw44.Ctor__()
					_287_v94 = _nw44
					var _288_v95 *C0
					_ = _288_v95
					var _nw45 *C0 = New_C0_()
					_ = _nw45
					_nw45.Ctor__((_271_v81).F14())
					_288_v95 = _nw45
					var _289_v96 *C0
					_ = _289_v96
					var _nw46 *C0 = New_C0_()
					_ = _nw46
					_nw46.Ctor__((_271_v81).F14())
					_289_v96 = _nw46
				} else if _source4.Is_DC2() {
					var _290___mcc_h16 bool = _source4.Get_().(D0_DC2).Cf2
					_ = _290___mcc_h16
					var _291___mcc_h17 bool = _source4.Get_().(D0_DC2).Cf3
					_ = _291___mcc_h17
					var _292___mcc_h18 _dafny.Int = _source4.Get_().(D0_DC2).Cf4
					_ = _292___mcc_h18
					var _293_cf4 _dafny.Int = _292___mcc_h18
					_ = _293_cf4
					var _294_cf3 bool = _291___mcc_h17
					_ = _294_cf3
					var _295_cf2 bool = _290___mcc_h16
					_ = _295_cf2
					var _296_v97 _dafny.Array
					_ = _296_v97
					var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
					_ = _nw47
					_296_v97 = _nw47
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_296_v97), 0))
					_ = _index45
					(_296_v97).ArraySet1(_275_v85, (_index45).Int())
					var _297_v98 _dafny.Array
					_ = _297_v98
					var _nwElement0_13 _dafny.Sequence = _146_v6
					_ = _nwElement0_13
					var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(2))
					_ = _nw48
					(_nw48).ArraySet1(_nwElement0_13, 0)
					(_nw48).ArraySet1(_146_v6, 1)
					_297_v98 = _nw48
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_297_v98), 0))
					_ = _index46
					(_297_v98).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_146_v6, _146_v6), (_index46).Int())
					var _298_v99 _dafny.Array
					_ = _298_v99
					var _len0_7 _dafny.Int = _dafny.IntOfInt64(24)
					_ = _len0_7
					var _nw49 _dafny.Array
					_ = _nw49
					if _len0_7.Cmp(_dafny.Zero) == 0 {
						_nw49 = _dafny.NewArray(_len0_7)
					} else {
						var _init7 func(_dafny.Int) _dafny.Int = (func(_299_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_300_i10 _dafny.Int) _dafny.Int {
								return (_300_i10).Minus(_299_cf4)
							}
						})(_293_cf4)
						_ = _init7
						var _element0_7 = _init7(_dafny.Zero)
						_ = _element0_7
						_nw49 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
						(_nw49).ArraySet1(_element0_7, 0)
						var _nativeLen0_7 = (_len0_7).Int()
						_ = _nativeLen0_7
						for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
							(_nw49).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
						}
					}
					_298_v99 = _nw49
					var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_298_v99), 0))
					_ = _index47
					(_298_v99).ArraySet1(_dafny.IntOfInt64(-345), (_index47).Int())
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_296_v97), 0))
					_ = _index48
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_297_v98), 0))
					_ = _index49
					var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_298_v99), 0))
					_ = _index50
					var _rhs27 _dafny.Array = _275_v85
					_ = _rhs27
					var _rhs28 _dafny.Sequence = _146_v6
					_ = _rhs28
					var _rhs29 _dafny.Int = _293_cf4
					_ = _rhs29
					var _rhs30 bool = (Companion_Default___.SafeModuloInt(_142_v3, _142_v3)).Cmp(_293_cf4) > 0
					_ = _rhs30
					var _lhs27 _dafny.Array = _296_v97
					_ = _lhs27
					var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_296_v97), 0))
					_ = _lhs28
					var _lhs29 _dafny.Array = _297_v98
					_ = _lhs29
					var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_297_v98), 0))
					_ = _lhs30
					var _lhs31 _dafny.Array = _298_v99
					_ = _lhs31
					var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_298_v99), 0))
					_ = _lhs32
					(_lhs27).ArraySet1(_rhs27, (_lhs28).Int())
					(_lhs29).ArraySet1(_rhs28, (_lhs30).Int())
					(_lhs31).ArraySet1(_rhs29, (_lhs32).Int())
					_295_cf2 = _rhs30
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_298_v99), 0))
					_ = _index51
					var _rhs31 _dafny.Int = Companion_Default___.Fm0(_145_globalState)
					_ = _rhs31
					var _rhs32 bool = !((Companion_D4_.Create_DC12_(_137_v0, _268_v78, (_175_v31).Dtor_cf1(), _275_v85, _137_v0)).Dtor_cf26())
					_ = _rhs32
					var _lhs33 _dafny.Array = _298_v99
					_ = _lhs33
					var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(40), _dafny.ArrayLen((_298_v99), 0))
					_ = _lhs34
					(_lhs33).ArraySet1(_rhs31, (_lhs34).Int())
					_294_cf3 = _rhs32
					var _301_v100 _dafny.Array
					_ = _301_v100
					var _nw50 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
					_ = _nw50
					_301_v100 = _nw50
					var _302_v101 _dafny.Map
					_ = _302_v101
					_302_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_294_cf3, _301_v100)
					_302_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_275_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))).Int()).(bool), _301_v100)
					_142_v3 = _293_cf4
				} else {
					var _303___mcc_h19 _dafny.Sequence = _source4.Get_().(D0_DC0).Cf0
					_ = _303___mcc_h19
					var _304_cf0 _dafny.Sequence = _303___mcc_h19
					_ = _304_cf0
					var _305_v102 _dafny.Array
					_ = _305_v102
					var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
					_ = _nw51
					_305_v102 = _nw51
					var _306_v103 _dafny.Map
					_ = _306_v103
					_306_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v0, _305_v102)
					var _307_v104 D3
					_ = _307_v104
					_307_v104 = Companion_D3_.Create_DC7_((func() _dafny.Array {
						if (_306_v103).Contains((_269_v79).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_269_v79).Cardinality()))).Uint32()).(bool)) {
							return (_306_v103).Get((_269_v79).Select((Companion_Default___.SafeIndex(_142_v3, _dafny.IntOfUint32((_269_v79).Cardinality()))).Uint32()).(bool)).(_dafny.Array)
						}
						return _305_v102
					})())
					_307_v104 = _307_v104
					var _308_v105 *C1
					_ = _308_v105
					var _nw52 *C1 = New_C1_()
					_ = _nw52
					_nw52.Ctor__()
					_308_v105 = _nw52
					var _309_v106 _dafny.Sequence
					_ = _309_v106
					_309_v106 = _dafny.SeqOf(_308_v105)
					var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))
					_ = _index52
					(_275_v85).ArraySet1((Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_304_cf0).Cardinality()), _142_v3)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_309_v106, _dafny.SeqOf(_308_v105, _308_v105))).Cardinality())) < 0, (_index52).Int())
					(_145_globalState).F0 = Companion_Default___.SafeModuloInt(_142_v3, _142_v3)
					(_145_globalState).F0 = _142_v3
				}
				var _310_v107 *C1
				_ = _310_v107
				var _nw53 *C1 = New_C1_()
				_ = _nw53
				_nw53.Ctor__()
				_310_v107 = _nw53
				var _311_v108 _dafny.Sequence
				_ = _311_v108
				_311_v108 = _dafny.SeqOf(_310_v107, _310_v107, _310_v107, _310_v107)
				(_145_globalState).F0 = (_142_v3).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_311_v108, _311_v108)).Cardinality()))
				(_145_globalState).F0 = (_142_v3).Times(_142_v3)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _312_v109 _dafny.Map
	_ = _312_v109
	_312_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v6, (_275_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_275_v85), 0))).Int()).(bool))
	var _313_v110 _dafny.Map
	_ = _313_v110
	_313_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v1, (func() _dafny.Int {
		if (_266_v76).Contains(_142_v3) {
			return (_266_v76).Get(_142_v3).(_dafny.Int)
		}
		return _142_v3
	})())
	_312_v109 = (_312_v109).Update(_146_v6, !(_313_v110).Contains(_dafny.SetOf(_137_v0)))
	_dafny.Print(_137_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v1).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_139_v2).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_142_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_143_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_144_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState).F1()).Equals(_dafny.SetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_145_globalState).F2()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_145_globalState.F3).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-730))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F5().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_145_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_146_v6.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_v31).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v48).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v48).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v48).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v48).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v72).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v73).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v73).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v73).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v74).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.IntOfInt64(423), true, true, true), _dafny.UnicodeSeqOfUtf8Bytes("whhg"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v75).Dtor_cf11().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v75).Dtor_cf12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v75).Dtor_cf13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v75).Dtor_cf14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v75).Dtor_cf15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_266_v76).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_267_v77, _dafny.SeqOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_268_v78, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(4), _dafny.One), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_269_v79, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v80).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v80).Dtor_cf18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v80).Dtor_cf19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_271_v81).F14()).Dtor_cf17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_271_v81).F14()).Dtor_cf18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_271_v81).F14()).Dtor_cf19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_272_v82).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v83).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v84).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.One, _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_275_v85).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_276_v86)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v87).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfInt64(615))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_278_v88).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v89).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.One, _dafny.One, false), _dafny.UnicodeSeqOfUtf8Bytes("whhg"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_280_v90).Dtor_cf20()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.One, _dafny.One, false), _dafny.UnicodeSeqOfUtf8Bytes("whhg"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v91).Equals(_dafny.MultiSetOf(_dafny.SeqOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_282_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_312_v109).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("whhg"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_313_v110).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), _dafny.One)))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
	Cf3 bool
	Cf4 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool, Cf3 bool, Cf4 _dafny.Int) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4}}
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
	return Companion_D0_.Create_DC1_(_dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf6 _dafny.Int
	Cf7 bool
	Cf8 bool
	Cf9 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.Int, Cf7 bool, Cf8 bool, Cf9 bool) D1 {
	return D1{D1_DC4{Cf6, Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf5 _dafny.Map
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Map) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero, false, false, false)
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
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

type D2_DC6 struct {
	Cf11 _dafny.Sequence
	Cf12 _dafny.Int
	Cf13 bool
	Cf14 _dafny.Int
	Cf15 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf11 _dafny.Sequence, Cf12 _dafny.Int, Cf13 bool, Cf14 _dafny.Int, Cf15 bool) D2 {
	return D2{D2_DC6{Cf11, Cf12, Cf13, Cf14, Cf15}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf10 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf10 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf10}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.EmptySeq, _dafny.Zero, false, _dafny.Zero, false)
}

func (_this D2) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf14
}

func (_this D2) Dtor_cf15() bool {
	return _this.Get_().(D2_DC6).Cf15
}

func (_this D2) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + data.Cf11.VerbatimString(true) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + data.Cf10.VerbatimString(true) + ")"
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
			return ok && data1.Cf11.Equals(data2.Cf11) && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
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
	Cf17 _dafny.Int
	Cf18 _dafny.Int
	Cf19 bool
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf17 _dafny.Int, Cf18 _dafny.Int, Cf19 bool) D3 {
	return D3{D3_DC8{Cf17, Cf18, Cf19}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf16 _dafny.Array
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf16 _dafny.Array) D3 {
	return D3{D3_DC7{Cf16}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC8).Cf19
}

func (_this D3) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D3_DC7).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf16 == data2.Cf16
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

type D4_DC10 struct {
	Cf21 _dafny.Sequence
	Cf22 bool
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf21 _dafny.Sequence, Cf22 bool) D4 {
	return D4{D4_DC10{Cf21, Cf22}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf23 *C0
	Cf24 _dafny.Int
	Cf25 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf23 *C0, Cf24 _dafny.Int, Cf25 _dafny.Int) D4 {
	return D4{D4_DC11{Cf23, Cf24, Cf25}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf26 bool
	Cf27 _dafny.Sequence
	Cf28 _dafny.Int
	Cf29 _dafny.Array
	Cf30 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf26 bool, Cf27 _dafny.Sequence, Cf28 _dafny.Int, Cf29 _dafny.Array, Cf30 bool) D4 {
	return D4{D4_DC12{Cf26, Cf27, Cf28, Cf29, Cf30}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC9 struct {
	Cf20 _dafny.Map
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf20 _dafny.Map) D4 {
	return D4{D4_DC9{Cf20}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(_dafny.EmptySeq, false)
}

func (_this D4) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D4_DC10).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC10).Cf22
}

func (_this D4) Dtor_cf23() *C0 {
	return _this.Get_().(D4_DC11).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf25
}

func (_this D4) Dtor_cf26() bool {
	return _this.Get_().(D4_DC12).Cf26
}

func (_this D4) Dtor_cf27() _dafny.Sequence {
	return _this.Get_().(D4_DC12).Cf27
}

func (_this D4) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf28
}

func (_this D4) Dtor_cf29() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf29
}

func (_this D4) Dtor_cf30() bool {
	return _this.Get_().(D4_DC12).Cf30
}

func (_this D4) Dtor_cf20() _dafny.Map {
	return _this.Get_().(D4_DC9).Cf20
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf21.Equals(data2.Cf21) && data1.Cf22 == data2.Cf22
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27.Equals(data2.Cf27) && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf20.Equals(data2.Cf20)
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
	Cf31 _dafny.CodePoint
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf31 _dafny.CodePoint) D5 {
	return D5{D5_DC13{Cf31}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D5) Dtor_cf31() _dafny.CodePoint {
	return _this.Get_().(D5_DC13).Cf31
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf31 == data2.Cf31
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

type D6_DC14 struct {
	Cf32 _dafny.Map
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf32 _dafny.Map) D6 {
	return D6{D6_DC14{Cf32}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D6) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D6_DC14).Cf32
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Int
	F3   _dafny.Array
	F4   _dafny.CodePoint
	F9   _dafny.CodePoint
	_f1  _dafny.Set
	_f2  _dafny.Array
	_f5  _dafny.Sequence
	_f6  _dafny.Int
	_f7  _dafny.Int
	_f8  bool
	_f10 bool
	_f11 _dafny.Int
	_f12 bool
	_f13 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F4 = _dafny.CodePoint('D')
	_this.F9 = _dafny.CodePoint('D')
	_this._f1 = _dafny.EmptySet
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f5 = _dafny.EmptySeq
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.Zero
	_this._f8 = false
	_this._f10 = false
	_this._f11 = _dafny.Zero
	_this._f12 = false
	_this._f13 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Set, f2 _dafny.Array, f3 _dafny.Array, f4 _dafny.CodePoint, f5 _dafny.Sequence, f6 _dafny.Int, f7 _dafny.Int, f8 bool, f9 _dafny.CodePoint, f10 bool, f11 _dafny.Int, f12 bool, f13 bool) {
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
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *GlobalState) F1() _dafny.Set {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F5() _dafny.Sequence {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
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
func (_this *GlobalState) F10() bool {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f14 D3
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f14 = Companion_D3_.Default()
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

func (_this *C0) Ctor__(f14 D3) {
	{
		(_this)._f14 = f14
	}
}
func (_this *C0) Fm6(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-847), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uqfinko")).Cardinality()))).Cardinality()), _dafny.IntOfInt64(15))).Minus(_dafny.IntOfInt64(615))
	}
}
func (_this *C0) F14() D3 {
	{
		return _this._f14
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	dummy byte
}

func New_C1_() *C1 {
	_this := C1{}

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

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm4(globalState *GlobalState) _dafny.Map {
	{
		return ((Companion_D1_.Create_DC3_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true, false, false), _dafny.UnicodeSeqOfUtf8Bytes("saf")))).Dtor_cf5()).Merge(func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), false)).Keys().Elements()); ; {
				_compr_5, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _314_v0 _dafny.Set
				_314_v0 = interface{}(_compr_5).(_dafny.Set)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), false)).Contains(_314_v0) {
					_coll5.Add(_314_v0, _dafny.UnicodeSeqOfUtf8Bytes("byiwjxmen"))
				}
			}
			return _coll5.ToMap()
		}())
	}
}
func (_this *C1) Fm5(globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_315_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("kipi"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vbvdrrtth"), _dafny.UnicodeSeqOfUtf8Bytes("n")))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D2_.Create_DC5_(_dafny.UnicodeSeqOfUtf8Bytes("rt"))).Dtor_cf10(), _dafny.UnicodeSeqOfUtf8Bytes("goeperti")))
	}
}
func (_this *C1) M1(p0 _dafny.Map, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) (_dafny.Int, bool, _dafny.Set, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _316_v0 _dafny.Int
		_ = _316_v0
		_316_v0 = _dafny.IntOfInt64(52)
		var _317_v1 D0
		_ = _317_v1
		_317_v1 = Companion_D0_.Create_DC1_(_316_v0)
		var _pat_let_tv14 = _316_v0
		_ = _pat_let_tv14
		r3 = (func(_pat_let20_0 D0) D0 {
			return func(_318_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let21_0 _dafny.Int) D0 {
					return func(_319_dt__update_hcf1_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_(_319_dt__update_hcf1_h0)
					}(_pat_let21_0)
				}(_pat_let_tv14)
			}(_pat_let20_0)
		}(_317_v1)).Dtor_cf1()
		var _hi3 _dafny.Int = (_316_v0).Plus(_316_v0)
		_ = _hi3
		for _320_i0 := _dafny.IntOfUint32((p1).Cardinality()); _320_i0.Cmp(_hi3) < 0; _320_i0 = _320_i0.Plus(_dafny.One) {
			r1 = !(p2)
			var _rhs33 bool = p2
			_ = _rhs33
			var _rhs34 _dafny.Int = _320_i0
			_ = _rhs34
			var _rhs35 bool = p2
			_ = _rhs35
			r1 = _rhs33
			_316_v0 = _rhs34
			r1 = _rhs35
			r1 = p2
			var _321_v2 _dafny.Sequence
			_ = _321_v2
			_321_v2 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()))
			var _rhs36 _dafny.Int = (_321_v2).Select((Companion_Default___.SafeIndex(_316_v0, _dafny.IntOfUint32((_321_v2).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs36
			var _rhs37 bool = Companion_Default___.Fm2(globalState)
			_ = _rhs37
			var _lhs35 *GlobalState = globalState
			_ = _lhs35
			_lhs35.F0 = _rhs36
			r1 = _rhs37
		}
		var _322_v3 _dafny.Array
		_ = _322_v3
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw54
		_322_v3 = _nw54
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
		_ = _index53
		(_322_v3).ArraySet1(p2, (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
		_ = _index54
		(_322_v3).ArraySet1(!(!(p2)) || (p2), (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
		_ = _index55
		var _rhs38 bool = !((_316_v0).Cmp(_dafny.IntOfUint32((p1).Cardinality())) >= 0)
		_ = _rhs38
		var _rhs39 bool = p2
		_ = _rhs39
		var _lhs36 _dafny.Array = _322_v3
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
		_ = _lhs37
		r1 = _rhs38
		(_lhs36).ArraySet1(_rhs39, (_lhs37).Int())
		var _hi4 _dafny.Int = _316_v0
		_ = _hi4
		for _323_i1 := (_316_v0).Minus(_316_v0); _323_i1.Cmp(_hi4) < 0; _323_i1 = _323_i1.Plus(_dafny.One) {
			var _324_v4 D0
			_ = _324_v4
			_324_v4 = Companion_D0_.Create_DC2_((_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool), (_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool), _323_i1)
			var _325_v5 _dafny.Sequence
			_ = _325_v5
			_325_v5 = _dafny.SeqOf(_324_v4, _324_v4)
			var _326_v6 _dafny.Map
			_ = _326_v6
			_326_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_316_v0, _dafny.IntOfInt64(178))
			var _327_v7 _dafny.MultiSet
			_ = _327_v7
			_327_v7 = _dafny.MultiSetOf(_326_v6, _326_v6)
			var _328_v8 _dafny.Sequence
			_ = _328_v8
			_328_v8 = _dafny.SeqOf(p2, (_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool))
			var _329_v9 _dafny.Set
			_ = _329_v9
			_329_v9 = _dafny.SetOf(_328_v8)
			var _330_v10 _dafny.Sequence
			_ = _330_v10
			_330_v10 = _dafny.SeqOf(_dafny.IntOfUint32((_325_v5).Cardinality()), (_dafny.Zero).Minus((_327_v7).Cardinality()), _323_i1, (_329_v9).Cardinality())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
			_ = _index56
			var _rhs40 _dafny.Int = _316_v0
			_ = _rhs40
			var _rhs41 bool = _dafny.Companion_Sequence_.Equal(_330_v10, _dafny.Companion_Sequence_.Concatenate(_330_v10, _330_v10))
			_ = _rhs41
			var _lhs38 _dafny.Array = _322_v3
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
			_ = _lhs39
			r3 = _rhs40
			(_lhs38).ArraySet1(_rhs41, (_lhs39).Int())
			if Companion_Default___.Fm2(globalState) {
				var _331_v11 _dafny.Map
				_ = _331_v11
				_331_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, Companion_Default___.Fm2(globalState))
				var _332_v12 _dafny.MultiSet
				_ = _332_v12
				_332_v12 = _dafny.MultiSetOf(Companion_Default___.Fm2(globalState), p2)
				_331_v11 = (_331_v11).Update((_dafny.MultiSetOf((_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool))).IsProperSubsetOf(_332_v12), (_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool))
				_331_v11 = (_331_v11).Update(p2, p2)
				var _333_v13 _dafny.Array
				_ = _333_v13
				var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw55
				_333_v13 = _nw55
				var _334_v14 D3
				_ = _334_v14
				_334_v14 = Companion_D3_.Create_DC7_(_333_v13)
				var _335_v15 _dafny.Map
				_ = _335_v15
				_335_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool), (_334_v14).Dtor_cf16())
				var _336_v16 _dafny.Sequence
				_ = _336_v16
				_336_v16 = _dafny.SeqOf(_333_v13, _333_v13, _333_v13, _333_v13)
				_335_v15 = (_335_v15).Update((_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool), (_336_v16).Select((Companion_Default___.SafeIndex(_323_i1, _dafny.IntOfUint32((_336_v16).Cardinality()))).Uint32()).(_dafny.Array))
				var _337_v17 _dafny.Array
				_ = _337_v17
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw56
				_337_v17 = _nw56
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_337_v17), 0))
				_ = _index57
				(_337_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_336_v16, _336_v16), (_index57).Int())
				var _338_v18 _dafny.Map
				_ = _338_v18
				_338_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_316_v0, p2)
				var _339_v19 _dafny.Sequence
				_ = _339_v19
				_339_v19 = _dafny.SeqOf(_338_v18, _338_v18)
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_337_v17), 0))
				_ = _index58
				(_337_v17).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_336_v16, _336_v16), _336_v16), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_339_v19).Cardinality()), _dafny.IntOfUint32((p1).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_336_v16, _336_v16), _336_v16)).Cardinality()))).Uint32(), (func() _dafny.Array {
					if (_335_v15).Contains(p2) {
						return (_335_v15).Get(p2).(_dafny.Array)
					}
					return _333_v13
				})()), (_index58).Int())
				_331_v11 = (_331_v11).Update(p2, p2)
			} else {
				_316_v0 = _dafny.IntOfUint32((_328_v8).Cardinality())
				var _340_v20 _dafny.CodePoint
				_ = _340_v20
				_340_v20 = _dafny.CodePoint('q')
				(globalState).F4 = _340_v20
				_326_v6 = (_326_v6).Update(_323_i1, _316_v0)
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))
				_ = _index59
				(_322_v3).ArraySet1(Companion_Default___.Fm2(globalState), (_index59).Int())
				var _341_v21 _dafny.Array
				_ = _341_v21
				var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw57
				_341_v21 = _nw57
				var _342_v22 _dafny.Sequence
				_ = _342_v22
				_342_v22 = _dafny.SeqOf(_341_v21, _341_v21)
				var _343_v23 _dafny.Map
				_ = _343_v23
				_343_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC7_((_342_v22).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.IntOfUint32((_342_v22).Cardinality()))).Uint32()).(_dafny.Array)), _316_v0)
				var _344_v24 _dafny.Array
				_ = _344_v24
				var _nwElement0_14 _dafny.Map = _343_v23
				_ = _nwElement0_14
				var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.One)
				_ = _nw58
				(_nw58).ArraySet1(_nwElement0_14, 0)
				_344_v24 = _nw58
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_344_v24), 0))
				_ = _index60
				(_344_v24).ArraySet1((_343_v23).Merge(_343_v23), (_index60).Int())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_344_v24), 0))
				_ = _index61
				(_344_v24).ArraySet1(_343_v23, (_index61).Int())
			}
			(globalState).F0 = _316_v0
			var _345_v25 D3
			_ = _345_v25
			_345_v25 = Companion_D3_.Create_DC8_(_323_i1, _316_v0, (_322_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_322_v3), 0))).Int()).(bool))
			var _346_v26 *C0
			_ = _346_v26
			var _nw59 *C0 = New_C0_()
			_ = _nw59
			_nw59.Ctor__(_345_v25)
			_346_v26 = _nw59
		}
		(globalState).F0 = Companion_Default___.SafeDivisionInt(_316_v0, (_316_v0).Times(_dafny.IntOfUint32((p1).Cardinality())))
		r0 = (_dafny.Zero).Minus(_dafny.IntOfUint32((p1).Cardinality()))
		r1 = !(p2)
		var _347_v27 _dafny.Set
		_ = _347_v27
		_347_v27 = _dafny.SetOf(_316_v0)
		r2 = (_347_v27).Union(_347_v27)
		r3 = Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus(_316_v0)).Minus(_316_v0), _316_v0)
		return r0, r1, r2, r3
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
