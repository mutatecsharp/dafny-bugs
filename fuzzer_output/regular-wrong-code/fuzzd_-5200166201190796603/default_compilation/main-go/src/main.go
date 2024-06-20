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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
		if true {
			return (func() _dafny.Int {
				if !(!(!(false))) {
					return _dafny.IntOfInt64(581)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-3), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xhl")).Cardinality()))).Cardinality()))).Cardinality()
			})()
		}
		return ((_dafny.MultiSetOf(_dafny.IntOfInt64(981), _dafny.IntOfInt64(-818))).Difference(_dafny.MultiSetOf((func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-229)), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Elements()); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _0_v0 _dafny.Int
				_0_v0 = interface{}(_compr_0).(_dafny.Int)
				if (_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-229)), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Contains(_0_v0) {
					_coll0.Add(Companion_Default___.SafeModuloInt(_0_v0, _dafny.IntOfInt64(-693)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-902), _dafny.CodePoint('d'))).Cardinality())
				}
			}
			return _coll0.ToMap()
		}()).Cardinality(), (_dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(430), _dafny.IntOfInt64(54))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _1_v1 _dafny.Int
				_1_v1 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(430)).Cmp(_1_v1) <= 0) && ((_1_v1).Cmp(_dafny.IntOfInt64(54)) < 0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_1_v1, _dafny.IntOfInt64(738)))
				}
			}
			return _coll1.ToSet()
		}()).Cardinality()))).Cardinality()))).Cardinality()
	})()))
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tjfujd"), _dafny.UnicodeSeqOfUtf8Bytes("cnhbwhio")))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(_dafny.IntOfInt64(977), false, _dafny.IntOfInt64(-930))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true, !(false))).Union((_dafny.MultiSetOf(false, true, false)).Difference(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(953), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-771)), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Map, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("txsljff")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(302))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_4_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("cuu")
	})))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	var _source0 D0 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(369), false, _dafny.IntOfInt64(448))
	_ = _source0
	if _source0.Is_DC1() {
		var _5___mcc_h0 _dafny.Int = _source0.Get_().(D0_DC1).Cf1
		_ = _5___mcc_h0
		var _6___mcc_h1 bool = _source0.Get_().(D0_DC1).Cf2
		_ = _6___mcc_h1
		var _7___mcc_h2 _dafny.Int = _source0.Get_().(D0_DC1).Cf3
		_ = _7___mcc_h2
		var _8_cf3 _dafny.Int = _7___mcc_h2
		_ = _8_cf3
		var _9_cf2 bool = _6___mcc_h1
		_ = _9_cf2
		var _10_cf1 _dafny.Int = _5___mcc_h0
		_ = _10_cf1
		return (_8_cf3).Cmp(_10_cf1) < 0
	} else if _source0.Is_DC2() {
		return (false) == (true)
	} else if _source0.Is_DC3() {
		return !(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(false))).Contains(true))
	} else {
		var _11___mcc_h3 _dafny.Int = _source0.Get_().(D0_DC0).Cf0
		_ = _11___mcc_h3
		var _12_cf0 _dafny.Int = _11___mcc_h3
		_ = _12_cf0
		return ((_dafny.Zero).Minus(_12_cf0)).Cmp(_dafny.IntOfInt64(917)) > 0
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.CodePoint('h'))).Union(_dafny.SetOf(_dafny.CodePoint('q')))).Union((Companion_D11_.Create_DC29_(_dafny.SetOf(_dafny.CodePoint('m'), _dafny.CodePoint('u')))).Dtor_cf39())
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(636))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_13_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(437)
	})), _dafny.SeqOf(_dafny.IntOfInt64(-774)))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((_dafny.MultiSetOf(!(false))).Cardinality(), _dafny.IntOfInt64(768), _dafny.IntOfInt64(863))).Difference(_dafny.SetOf((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('d'))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _14_v0 _dafny.CodePoint
			_14_v0 = interface{}(_compr_2).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('d')), _14_v0) {
				_coll2.Add(_14_v0, Companion_D9_.Create_DC23_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jtgcdb")).Cardinality())))
			}
		}
		return _coll2.ToMap()
	}()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, !(false), true), _dafny.SeqOf(true, !(true))))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, globalState *GlobalState) {
	var _15_v0 _dafny.Array
	_ = _15_v0
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(28)
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = func(_16_i0 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeDivisionInt(_16_i0, _dafny.IntOfInt64(-52))
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
	_15_v0 = _nw0
	var _17_v1 _dafny.Int
	_ = _17_v1
	_17_v1 = _dafny.IntOfInt64(473)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
	_ = _index0
	(_15_v0).ArraySet1(Companion_Default___.Fm0(_17_v1, _17_v1, globalState), (_index0).Int())
	var _18_v2 _dafny.Sequence
	_ = _18_v2
	_18_v2 = _dafny.SeqOf(p0)
	var _19_v3 _dafny.Sequence
	_ = _19_v3
	_19_v3 = _dafny.SeqOf(_17_v1, _dafny.IntOfUint32((_18_v2).Cardinality()))
	var _20_v4 _dafny.Sequence
	_ = _20_v4
	_20_v4 = _dafny.SeqOf(_19_v3)
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
	_ = _index1
	(_15_v0).ArraySet1(((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_17_v1, _17_v1))).Plus(_dafny.IntOfUint32((_20_v4).Cardinality())), (_index1).Int())
	var _21_i1 _dafny.Int
	_ = _21_i1
	_21_i1 = _dafny.Zero
	{
		for (_18_v2).Select((Companion_Default___.SafeIndex((_17_v1).Minus(_17_v1), _dafny.IntOfUint32((_18_v2).Cardinality()))).Uint32()).(bool) {
			{
				if (_21_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_21_i1 = (_21_i1).Plus(_dafny.One)
				var _22_v5 _dafny.Array
				_ = _22_v5
				var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_1
				var _nw1 _dafny.Array
				_ = _nw1
				if _len0_1.Cmp(_dafny.Zero) == 0 {
					_nw1 = _dafny.NewArray(_len0_1)
				} else {
					var _init1 func(_dafny.Int) _dafny.Sequence = func(_23_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("kvewhxik")
					}
					_ = _init1
					var _element0_1 = _init1(_dafny.Zero)
					_ = _element0_1
					_nw1 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
					(_nw1).ArraySet1(_element0_1, 0)
					var _nativeLen0_1 = (_len0_1).Int()
					_ = _nativeLen0_1
					for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
						(_nw1).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
					}
				}
				_22_v5 = _nw1
				var _24_v6 _dafny.Sequence
				_ = _24_v6
				_24_v6 = _dafny.UnicodeSeqOfUtf8Bytes("mqwfqaq")
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_22_v5), 0))
				_ = _index2
				(_22_v5).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_24_v6, _24_v6), (_index2).Int())
				var _25_v7 _dafny.Array
				_ = _25_v7
				var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
				_ = _nw2
				_25_v7 = _nw2
				var _26_v8 D0
				_ = _26_v8
				_26_v8 = Companion_D0_.Create_DC0_(_17_v1)
				var _27_v9 _dafny.MultiSet
				_ = _27_v9
				_27_v9 = _dafny.MultiSetOf(p0, p0, p0, true, p0)
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_25_v7), 0))
				_ = _index3
				(_25_v7).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v8, _27_v9), (_index3).Int())
				var _28_v10 _dafny.MultiSet
				_ = _28_v10
				_28_v10 = _dafny.MultiSetOf(_dafny.IntOfInt64(607))
				var _29_v11 _dafny.Map
				_ = _29_v11
				_29_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v8, Companion_Default___.Fm3(p0, _28_v10, p0, globalState))
				var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_22_v5), 0))
				_ = _index4
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_25_v7), 0))
				_ = _index5
				var _rhs0 _dafny.Sequence = _24_v6
				_ = _rhs0
				var _rhs1 _dafny.Int = (_dafny.IntOfInt64(106)).Plus(Companion_Default___.SafeDivisionInt((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)))
				_ = _rhs1
				var _rhs2 _dafny.Map = (_29_v11).Merge(_29_v11)
				_ = _rhs2
				var _rhs3 bool = !((_27_v9).IsProperSubsetOf(_27_v9))
				_ = _rhs3
				var _lhs0 _dafny.Array = _22_v5
				_ = _lhs0
				var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_22_v5), 0))
				_ = _lhs1
				var _lhs2 *GlobalState = globalState
				_ = _lhs2
				var _lhs3 _dafny.Array = _25_v7
				_ = _lhs3
				var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_25_v7), 0))
				_ = _lhs4
				var _lhs5 *GlobalState = globalState
				_ = _lhs5
				(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
				_lhs2.F4 = _rhs1
				(_lhs3).ArraySet1(_rhs2, (_lhs4).Int())
				_lhs5.F2 = _rhs3
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
				_ = _index6
				(_15_v0).ArraySet1(_17_v1, (_index6).Int())
				var _30_v13 _dafny.Array
				_ = _30_v13
				var _len0_2 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_2
				var _nw3 _dafny.Array
				_ = _nw3
				if _len0_2.Cmp(_dafny.Zero) == 0 {
					_nw3 = _dafny.NewArray(_len0_2)
				} else {
					var _init2 func(_dafny.Int) bool = (func(_31_v1 _dafny.Int, _32_v10 _dafny.MultiSet, _33_p0 bool, _34_v0 _dafny.Array) func(_dafny.Int) bool {
						return func(_35_i3 _dafny.Int) bool {
							return (func() _dafny.Set {
								var _coll3 = _dafny.NewBuilder()
								_ = _coll3
								for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(671), _dafny.IntOfInt64(-131))); ; {
									_compr_3, _ok3 := _iter3()
									if !_ok3 {
										break
									}
									var _36_v12 _dafny.Int
									_36_v12 = interface{}(_compr_3).(_dafny.Int)
									if ((_dafny.IntOfInt64(671)).Cmp(_36_v12) <= 0) && ((_36_v12).Cmp(_dafny.IntOfInt64(-131)) < 0) {
										_coll3.Add((_36_v12).Times((Companion_D0_.Create_DC1_((_32_v10).Cardinality(), _33_p0, (_34_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_34_v0), 0))).Int()).(_dafny.Int))).Dtor_cf3()))
									}
								}
								return _coll3.ToSet()
							}()).IsDisjointFrom(_dafny.SetOf(_31_v1))
						}
					})(_17_v1, _28_v10, p0, _15_v0)
					_ = _init2
					var _element0_2 = _init2(_dafny.Zero)
					_ = _element0_2
					_nw3 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
					(_nw3).ArraySet1(_element0_2, 0)
					var _nativeLen0_2 = (_len0_2).Int()
					_ = _nativeLen0_2
					for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
						(_nw3).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
					}
				}
				_30_v13 = _nw3
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_30_v13), 0))
				_ = _index7
				(_30_v13).ArraySet1(true, (_index7).Int())
				var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_30_v13), 0))
				_ = _index8
				(_30_v13).ArraySet1(!(Companion_Default___.Fm6(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(698))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_37_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (_22_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_22_v5), 0))).Int()).(_dafny.Sequence)), (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), globalState)), (_index8).Int())
				var _38_v14 D0
				_ = _38_v14
				_38_v14 = Companion_D0_.Create_DC1_(_17_v1, p0, (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
				var _39_v15 _dafny.Map
				_ = _39_v15
				_39_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v14, (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
				var _40_v16 *C0
				_ = _40_v16
				var _nw4 *C0 = New_C0_()
				_ = _nw4
				_nw4.Ctor__(_39_v15)
				_40_v16 = _nw4
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _hi0 _dafny.Int = Companion_Default___.SafeDivisionInt((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
	_ = _hi0
	for _41_i5 := (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int); _41_i5.Cmp(_hi0) < 0; _41_i5 = _41_i5.Plus(_dafny.One) {
		var _42_v17 _dafny.Set
		_ = _42_v17
		_42_v17 = _dafny.SetOf(p0, p0)
		var _43_v18 _dafny.CodePoint
		_ = _43_v18
		_43_v18 = _dafny.CodePoint('y')
		var _44_v19 _dafny.Map
		_ = _44_v19
		_44_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_v17, _43_v18)
		var _45_v20 D9
		_ = _45_v20
		_45_v20 = Companion_D9_.Create_DC22_(_42_v17)
		_44_v19 = (_44_v19).Update((_45_v20).Dtor_cf25(), _43_v18)
		var _46_v21 D0
		_ = _46_v21
		_46_v21 = Companion_D0_.Create_DC1_(_41_i5, p0, _17_v1)
		var _47_v22 _dafny.Map
		_ = _47_v22
		_47_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v21, _dafny.IntOfInt64(-29))
		var _48_v23 *C0
		_ = _48_v23
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__(_47_v22)
		_48_v23 = _nw5
		(globalState).F3 = _41_i5
		var _49_v24 _dafny.Sequence
		_ = _49_v24
		_49_v24 = _dafny.UnicodeSeqOfUtf8Bytes("lfxeijna")
		var _50_v25 _dafny.MultiSet
		_ = _50_v25
		_50_v25 = _dafny.MultiSetOf(_49_v24, _49_v24, _dafny.UnicodeSeqOfUtf8Bytes("mv"), _49_v24)
		var _51_v26 _dafny.MultiSet
		_ = _51_v26
		_51_v26 = _dafny.MultiSetOf(_41_i5, (_19_v3).Select((Companion_Default___.SafeIndex((_50_v25).Cardinality(), _dafny.IntOfUint32((_19_v3).Cardinality()))).Uint32()).(_dafny.Int))
		var _52_v27 _dafny.Map
		_ = _52_v27
		_52_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_17_v1).Cmp((func() _dafny.Int {
			if (_51_v26).Contains(_17_v1) {
				return (_51_v26).Multiplicity(_17_v1)
			}
			return _17_v1
		})()) == 0, Companion_Default___.SafeModuloInt(_41_i5, _17_v1))
		_52_v27 = (_52_v27).Update(p0, _17_v1)
	}
	if p0 {
		var _53_v29 _dafny.Map
		_ = _53_v29
		_53_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_v1, p0)
		var _54_v30 D0
		_ = _54_v30
		_54_v30 = Companion_D0_.Create_DC1_((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), (func() bool {
			if (_53_v29).Contains((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)) {
				return (_53_v29).Get((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)).(bool)
			}
			return p0
		})(), _17_v1)
		var _55_v31 _dafny.Map
		_ = _55_v31
		_55_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v30, (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
		var _56_v32 *C0
		_ = _56_v32
		var _nw6 *C0 = New_C0_()
		_ = _nw6
		_nw6.Ctor__(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_55_v31).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _57_v28 D0
				_57_v28 = interface{}(_compr_4).(D0)
				if (_55_v31).Contains(_57_v28) {
					_coll4.Add(_57_v28, _17_v1)
				}
			}
			return _coll4.ToMap()
		}())
		_56_v32 = _nw6
		var _58_v33 _dafny.CodePoint
		_ = _58_v33
		_58_v33 = _dafny.CodePoint('l')
		var _59_v34 D9
		_ = _59_v34
		_59_v34 = Companion_D9_.Create_DC24_(Companion_Default___.Fm0((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), globalState), _56_v32, _58_v33, _17_v1)
		var _60_v35 _dafny.Map
		_ = _60_v35
		_60_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v34, (_dafny.Zero).Minus(_17_v1))
		var _61_v36 _dafny.MultiSet
		_ = _61_v36
		_61_v36 = _dafny.MultiSetOf(p0)
		var _62_v37 _dafny.Sequence
		_ = _62_v37
		_62_v37 = _dafny.UnicodeSeqOfUtf8Bytes("qaoscxjb")
		_60_v35 = (_60_v35).Update(_59_v34, Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_61_v36).Contains(p0) {
				return (_61_v36).Multiplicity(p0)
			}
			return _dafny.IntOfUint32((_62_v37).Cardinality())
		})(), (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)))
		(globalState).F2 = p0
		(globalState).F2 = p0
		var _63_v38 _dafny.Map
		_ = _63_v38
		_63_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _58_v33)
		var _64_v39 D8
		_ = _64_v39
		_64_v39 = Companion_D8_.Create_DC20_(_58_v33)
		var _65_v40 _dafny.Array
		_ = _65_v40
		var _nwElement0_0 _dafny.CodePoint = _dafny.CodePoint('q')
		_ = _nwElement0_0
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(26))
		_ = _nw7
		(_nw7).ArraySet1CodePoint(_nwElement0_0, 0)
		(_nw7).ArraySet1CodePoint(_58_v33, 1)
		(_nw7).ArraySet1CodePoint(_58_v33, 2)
		(_nw7).ArraySet1CodePoint(_58_v33, 3)
		(_nw7).ArraySet1CodePoint((Companion_D9_.Create_DC24_((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), _56_v32, _dafny.CodePoint('g'), _17_v1)).Dtor_cf29(), 4)
		(_nw7).ArraySet1CodePoint(_58_v33, 5)
		(_nw7).ArraySet1CodePoint(_58_v33, 6)
		(_nw7).ArraySet1CodePoint(_58_v33, 7)
		(_nw7).ArraySet1CodePoint(_dafny.CodePoint('x'), 8)
		(_nw7).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_63_v38).Contains(p0) {
				return (_63_v38).Get(p0).(_dafny.CodePoint)
			}
			return _58_v33
		})(), 9)
		(_nw7).ArraySet1CodePoint(_dafny.CodePoint('s'), 10)
		(_nw7).ArraySet1CodePoint(_58_v33, 11)
		(_nw7).ArraySet1CodePoint(_58_v33, 12)
		(_nw7).ArraySet1CodePoint(_58_v33, 13)
		(_nw7).ArraySet1CodePoint(_58_v33, 14)
		(_nw7).ArraySet1CodePoint(_dafny.CodePoint('w'), 15)
		(_nw7).ArraySet1CodePoint(_58_v33, 16)
		(_nw7).ArraySet1CodePoint(_58_v33, 17)
		(_nw7).ArraySet1CodePoint((_64_v39).Dtor_cf22(), 18)
		(_nw7).ArraySet1CodePoint(_58_v33, 19)
		(_nw7).ArraySet1CodePoint(_58_v33, 20)
		(_nw7).ArraySet1CodePoint(_58_v33, 21)
		(_nw7).ArraySet1CodePoint(_58_v33, 22)
		(_nw7).ArraySet1CodePoint(_dafny.CodePoint('k'), 23)
		(_nw7).ArraySet1CodePoint(_58_v33, 24)
		(_nw7).ArraySet1CodePoint(_58_v33, 25)
		_65_v40 = _nw7
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_65_v40), 0))
		_ = _index9
		(_65_v40).ArraySet1CodePoint(_58_v33, (_index9).Int())
		var _66_v41 _dafny.Set
		_ = _66_v41
		_66_v41 = _dafny.SetOf((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
		_ = _index10
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_65_v40), 0))
		_ = _index11
		var _rhs4 _dafny.Int = (func() _dafny.Int {
			if p0 {
				return _dafny.IntOfInt64(-33)
			}
			return (_66_v41).Cardinality()
		})()
		_ = _rhs4
		var _rhs5 _dafny.CodePoint = _58_v33
		_ = _rhs5
		var _rhs6 _dafny.Int = _17_v1
		_ = _rhs6
		var _lhs6 _dafny.Array = _15_v0
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
		_ = _lhs7
		var _lhs8 _dafny.Array = _65_v40
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_65_v40), 0))
		_ = _lhs9
		var _lhs10 *GlobalState = globalState
		_ = _lhs10
		(_lhs6).ArraySet1(_rhs4, (_lhs7).Int())
		(_lhs8).ArraySet1CodePoint(_rhs5, (_lhs9).Int())
		_lhs10.F3 = _rhs6
		var _67_v42 _dafny.Map
		_ = _67_v42
		_67_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
			if true {
				return _62_v37
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-146))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_68_v40 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_69_i6 _dafny.Int) _dafny.CodePoint {
					return (_68_v40).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_68_v40), 0))).Int())
				}
			})(_65_v40)))
		})(), ((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(476)))
		_67_v42 = _67_v42
	} else {
		if p0 {
			var _70_v43 _dafny.Sequence
			_ = _70_v43
			_70_v43 = _dafny.UnicodeSeqOfUtf8Bytes("cnu")
			var _71_v44 _dafny.CodePoint
			_ = _71_v44
			_71_v44 = _dafny.CodePoint('n')
			var _72_v45 D0
			_ = _72_v45
			_72_v45 = Companion_D0_.Create_DC1_((_dafny.IntOfInt64(343)).Minus(_17_v1), _dafny.Companion_Sequence_.IsPrefixOf(_70_v43, _dafny.Companion_Sequence_.Update(_70_v43, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.IntOfUint32((_70_v43).Cardinality()))).Uint32(), _71_v44)), _dafny.IntOfUint32((_18_v2).Cardinality()))
			_72_v45 = _72_v45
			(globalState).F4 = (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)
			var _73_v46 D0
			_ = _73_v46
			_73_v46 = Companion_D0_.Create_DC3_()
			var _74_v47 _dafny.Map
			_ = _74_v47
			_74_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v46, _15_v0)
			_74_v47 = (_74_v47).Update(_73_v46, _15_v0)
			var _75_v48 _dafny.Map
			_ = _75_v48
			_75_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _76_v50 _dafny.Map
			_ = _76_v50
			_76_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v45, _19_v3)
			var _77_v51 *C0
			_ = _77_v51
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__(func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_76_v50).Keys().Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _78_v49 D0
					_78_v49 = interface{}(_compr_5).(D0)
					if (_76_v50).Contains(_78_v49) {
						_coll5.Add(_78_v49, _dafny.IntOfUint32((_70_v43).Cardinality()))
					}
				}
				return _coll5.ToMap()
			}())
			_77_v51 = _nw8
			var _79_v52 D10
			_ = _79_v52
			_79_v52 = Companion_D10_.Create_DC26_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v48, _77_v51))
			_17_v1 = (((_79_v52).Dtor_cf31()).Cardinality()).Times(_17_v1)
			_71_v44 = _71_v44
		} else {
			var _80_v53 _dafny.Sequence
			_ = _80_v53
			_80_v53 = _dafny.UnicodeSeqOfUtf8Bytes("kqlo")
			var _81_v54 _dafny.Array
			_ = _81_v54
			var _len0_3 _dafny.Int = _dafny.One
			_ = _len0_3
			var _nw9 _dafny.Array
			_ = _nw9
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw9 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_82_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_83_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_82_v2, _82_v2)
					}
				})(_18_v2)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw9 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw9).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw9).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_81_v54 = _nw9
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_81_v54), 0))
			_ = _index12
			(_81_v54).ArraySet1(_dafny.SeqOf(!(true), p0, p0, p0, p0), (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_81_v54), 0))
			_ = _index13
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
			_ = _index14
			var _rhs7 _dafny.Sequence = (func() _dafny.Sequence {
				if p0 {
					return _80_v53
				}
				return _80_v53
			})()
			_ = _rhs7
			var _rhs8 _dafny.Int = (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs8
			var _rhs9 _dafny.Sequence = _18_v2
			_ = _rhs9
			var _rhs10 _dafny.Int = (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs10
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 _dafny.Array = _81_v54
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_81_v54), 0))
			_ = _lhs13
			var _lhs14 _dafny.Array = _15_v0
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
			_ = _lhs15
			_80_v53 = _rhs7
			_lhs11.F4 = _rhs8
			(_lhs12).ArraySet1(_rhs9, (_lhs13).Int())
			(_lhs14).ArraySet1(_rhs10, (_lhs15).Int())
			var _84_v55 _dafny.Map
			_ = _84_v55
			_84_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_17_v1, (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
			var _85_v57 _dafny.Sequence
			_ = _85_v57
			_85_v57 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(923), globalState), (_dafny.MultiSetOf(_17_v1, _dafny.IntOfInt64(-135))).Cardinality()))
			var _86_v58 _dafny.MultiSet
			_ = _86_v58
			_86_v58 = _dafny.MultiSetOf(_84_v55, _84_v55, func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(686), _dafny.IntOfInt64(758))); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _87_v56 _dafny.Int
					_87_v56 = interface{}(_compr_6).(_dafny.Int)
					if ((_dafny.IntOfInt64(686)).Cmp(_87_v56) <= 0) && ((_87_v56).Cmp(_dafny.IntOfInt64(758)) < 0) {
						_coll6.Add((_87_v56).Plus(_dafny.IntOfInt64(387)), _17_v1)
					}
				}
				return _coll6.ToMap()
			}(), (_85_v57).Select((Companion_Default___.SafeIndex(_17_v1, _dafny.IntOfUint32((_85_v57).Cardinality()))).Uint32()).(_dafny.Map), _84_v55)
			(globalState).F2 = ((_86_v58).Union(_86_v58)).IsProperSubsetOf(_86_v58)
			var _88_v59 _dafny.Array
			_ = _88_v59
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw10
			_88_v59 = _nw10
			var _89_v60 D7
			_ = _89_v60
			_89_v60 = Companion_D7_.Create_DC17_(_15_v0)
			var _90_v61 _dafny.Map
			_ = _90_v61
			_90_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(451), (_89_v60).Dtor_cf19())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_88_v59), 0))
			_ = _index15
			(_88_v59).ArraySet1(!(_90_v61).Equals(_90_v61), (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_88_v59), 0))
			_ = _index16
			(_88_v59).ArraySet1(p0, (_index16).Int())
			(globalState).F2 = (_88_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_88_v59), 0))).Int()).(bool)
			(globalState).F2 = ((_dafny.Zero).Minus((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfInt64(621)) != 0
		}
		_19_v3 = _dafny.SeqOf(_dafny.IntOfInt64(494))
		_15_v0 = _15_v0
		var _91_v62 D0
		_ = _91_v62
		_91_v62 = Companion_D0_.Create_DC1_((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int), (_18_v2).Select((Companion_Default___.SafeIndex(_17_v1, _dafny.IntOfUint32((_18_v2).Cardinality()))).Uint32()).(bool), (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
		var _92_v63 _dafny.Map
		_ = _92_v63
		_92_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v62, _17_v1)
		var _93_v64 *C0
		_ = _93_v64
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__(_92_v63)
		_93_v64 = _nw11
		var _94_v65 _dafny.CodePoint
		_ = _94_v65
		_94_v65 = _dafny.CodePoint('k')
		var _95_v66 D9
		_ = _95_v66
		_95_v66 = Companion_D9_.Create_DC24_(_17_v1, _93_v64, _94_v65, (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int))
		var _96_v67 _dafny.Map
		_ = _96_v67
		_96_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v66, p0)
		var _97_v68 _dafny.Array
		_ = _97_v68
		var _nwElement0_1 _dafny.Sequence = _18_v2
		_ = _nwElement0_1
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(9))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_1, 0)
		(_nw12).ArraySet1(_18_v2, 1)
		(_nw12).ArraySet1(_18_v2, 2)
		(_nw12).ArraySet1(_dafny.SeqOf(p0, p0), 3)
		(_nw12).ArraySet1((func() _dafny.Sequence {
			if p0 {
				return _18_v2
			}
			return _18_v2
		})(), 4)
		(_nw12).ArraySet1(_dafny.SeqOf(!(false), (func() bool {
			if (_96_v67).Contains(_95_v66) {
				return (_96_v67).Get(_95_v66).(bool)
			}
			return p0
		})()), 5)
		(_nw12).ArraySet1(_dafny.SeqOf(p0, p0), 6)
		(_nw12).ArraySet1(_18_v2, 7)
		(_nw12).ArraySet1(_18_v2, 8)
		_97_v68 = _nw12
		var _98_v69 _dafny.Array
		_ = _98_v69
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(21))
		_ = _nw13
		_98_v69 = _nw13
		var _99_v70 D6
		_ = _99_v70
		_99_v70 = Companion_D6_.Create_DC15_(_17_v1, _17_v1)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_98_v69), 0))
		_ = _index17
		(_98_v69).ArraySet1(_99_v70, (_index17).Int())
		var _pat_let_tv0 = _17_v1
		_ = _pat_let_tv0
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_98_v69), 0))
		_ = _index18
		var _rhs11 _dafny.Array = _97_v68
		_ = _rhs11
		var _rhs12 D6 = func(_pat_let0_0 D6) D6 {
			return func(_100_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let1_0 _dafny.Int) D6 {
					return func(_101_dt__update_hcf17_h0 _dafny.Int) D6 {
						return Companion_D6_.Create_DC15_((_100_dt__update__tmp_h0).Dtor_cf16(), _101_dt__update_hcf17_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_99_v70)
		_ = _rhs12
		var _lhs16 _dafny.Array = _98_v69
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_98_v69), 0))
		_ = _lhs17
		_97_v68 = _rhs11
		(_lhs16).ArraySet1(_rhs12, (_lhs17).Int())
		var _102_v71 _dafny.Array
		_ = _102_v71
		var _nw14 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
		_ = _nw14
		_102_v71 = _nw14
		var _103_v72 _dafny.Map
		_ = _103_v72
		_103_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v71, p0)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
		_ = _index19
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
		_ = _index20
		var _rhs13 _dafny.Int = (_dafny.Zero).Minus((((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus(_17_v1))).Minus((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)))
		_ = _rhs13
		var _rhs14 bool = (p0) == (!(((_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)).Cmp((_103_v72).Cardinality()) <= 0))
		_ = _rhs14
		var _rhs15 _dafny.Int = (_15_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))).Int()).(_dafny.Int)
		_ = _rhs15
		var _rhs16 _dafny.Int = Companion_Default___.Fm0(_17_v1, _17_v1, globalState)
		_ = _rhs16
		var _lhs18 *GlobalState = globalState
		_ = _lhs18
		var _lhs19 *GlobalState = globalState
		_ = _lhs19
		var _lhs20 _dafny.Array = _15_v0
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
		_ = _lhs21
		var _lhs22 _dafny.Array = _15_v0
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_15_v0), 0))
		_ = _lhs23
		_lhs18.F4 = _rhs13
		_lhs19.F2 = _rhs14
		(_lhs20).ArraySet1(_rhs15, (_lhs21).Int())
		(_lhs22).ArraySet1(_rhs16, (_lhs23).Int())
	}
	var _104_v73 _dafny.CodePoint
	_ = _104_v73
	_104_v73 = _dafny.CodePoint('f')
	var _105_v74 _dafny.Sequence
	_ = _105_v74
	_105_v74 = _dafny.SeqOf(_104_v73)
	var _hi1 _dafny.Int = _17_v1
	_ = _hi1
	for _106_i8 := (func() _dafny.Int {
		if p0 {
			return _17_v1
		}
		return _dafny.IntOfUint32((_105_v74).Cardinality())
	})(); _106_i8.Cmp(_hi1) < 0; _106_i8 = _106_i8.Plus(_dafny.One) {
		(globalState).F2 = !(p0) || (p0)
		var _107_v75 _dafny.Set
		_ = _107_v75
		_107_v75 = _dafny.SetOf((func() bool {
			if p0 {
				return p0
			}
			return p0
		})())
		_107_v75 = _107_v75
		var _108_v76 _dafny.MultiSet
		_ = _108_v76
		_108_v76 = _dafny.MultiSetOf(false)
		_108_v76 = (_dafny.MultiSetOf(p0, p0, p0, p0, true)).Union((_108_v76).Intersection(_108_v76))
		var _109_v77 _dafny.Set
		_ = _109_v77
		_109_v77 = _dafny.SetOf(_17_v1, _106_i8)
		var _110_v78 D0
		_ = _110_v78
		_110_v78 = Companion_D0_.Create_DC1_(_106_i8, p0, (_109_v77).Cardinality())
		var _111_v79 _dafny.Map
		_ = _111_v79
		_111_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let2_0 D0) D0 {
			return func(_112_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let3_0 _dafny.Int) D0 {
					return func(_113_dt__update_hcf3_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_((_112_dt__update__tmp_h1).Dtor_cf1(), (_112_dt__update__tmp_h1).Dtor_cf2(), _113_dt__update_hcf3_h0)
					}(_pat_let3_0)
				}(_106_i8)
			}(_pat_let2_0)
		}(_110_v78), _17_v1)
		var _114_v80 T0
		_ = _114_v80
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__(_111_v79)
		_114_v80 = _nw15
		var _115_v81 _dafny.Array
		_ = _115_v81
		var _nwElement0_2 T0 = _114_v80
		_ = _nwElement0_2
		var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
		_ = _nw16
		(_nw16).ArraySet1(_nwElement0_2, 0)
		(_nw16).ArraySet1(_114_v80, 1)
		_115_v81 = _nw16
		_115_v81 = (func() _dafny.Array {
			if p0 {
				return _115_v81
			}
			return _115_v81
		})()
	}
	var _116_v82 _dafny.Set
	_ = _116_v82
	_116_v82 = _dafny.SetOf(_19_v3)
	(globalState).F2 = (func() bool {
		if (_116_v82).IsProperSubsetOf(_116_v82) {
			return p0
		}
		return (_17_v1).Cmp(_17_v1) > 0
	})()
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _117_v0 _dafny.Int
	_ = _117_v0
	_117_v0 = _dafny.IntOfInt64(877)
	var _118_v1 _dafny.Set
	_ = _118_v1
	_118_v1 = _dafny.SetOf((_dafny.Zero).Minus(_117_v0))
	var _119_v2 _dafny.Sequence
	_ = _119_v2
	_119_v2 = _dafny.UnicodeSeqOfUtf8Bytes("yxsvjmbqp")
	var _120_v3 _dafny.Set
	_ = _120_v3
	_120_v3 = _dafny.SetOf(_119_v2, _119_v2)
	var _121_v4 _dafny.Array
	_ = _121_v4
	var _nwElement0_3 _dafny.Int = _117_v0
	_ = _nwElement0_3
	var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(20))
	_ = _nw17
	(_nw17).ArraySet1(_nwElement0_3, 0)
	(_nw17).ArraySet1(_117_v0, 1)
	(_nw17).ArraySet1(_117_v0, 2)
	(_nw17).ArraySet1(_117_v0, 3)
	(_nw17).ArraySet1(_117_v0, 4)
	(_nw17).ArraySet1(_dafny.IntOfInt64(782), 5)
	(_nw17).ArraySet1(_117_v0, 6)
	(_nw17).ArraySet1(_117_v0, 7)
	(_nw17).ArraySet1(_117_v0, 8)
	(_nw17).ArraySet1((Companion_D0_.Create_DC0_(_117_v0)).Dtor_cf0(), 9)
	(_nw17).ArraySet1(_117_v0, 10)
	(_nw17).ArraySet1(_117_v0, 11)
	(_nw17).ArraySet1((_118_v1).Cardinality(), 12)
	(_nw17).ArraySet1((_dafny.Zero).Minus((_120_v3).Cardinality()), 13)
	(_nw17).ArraySet1(_117_v0, 14)
	(_nw17).ArraySet1(_117_v0, 15)
	(_nw17).ArraySet1(_117_v0, 16)
	(_nw17).ArraySet1((_dafny.Zero).Minus(_117_v0), 17)
	(_nw17).ArraySet1(_117_v0, 18)
	(_nw17).ArraySet1(_dafny.IntOfInt64(217), 19)
	_121_v4 = _nw17
	var _122_v5 _dafny.Array
	_ = _122_v5
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_4
	var _nw18 _dafny.Array
	_ = _nw18
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw18 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) bool = (func(_123_v0 _dafny.Int) func(_dafny.Int) bool {
			return func(_124_i0 _dafny.Int) bool {
				return (Companion_D0_.Create_DC1_(_123_v0, true, _123_v0)).Dtor_cf2()
			}
		})(_117_v0)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw18 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw18).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw18).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_122_v5 = _nw18
	var _125_globalState *GlobalState
	_ = _125_globalState
	var _nw19 *GlobalState = New_GlobalState_()
	_ = _nw19
	_nw19.Ctor__(false, true, true, _dafny.IntOfInt64(-701), _dafny.IntOfInt64(757), _dafny.IntOfInt64(943), _121_v4, true, _121_v4, _122_v5, _dafny.IntOfInt64(685), false, _119_v2, _dafny.IntOfInt64(826), true)
	_125_globalState = _nw19
	var _126_i1 _dafny.Int
	_ = _126_i1
	_126_i1 = _dafny.Zero
	{
		for (_117_v0).Cmp(_117_v0) == 0 {
			{
				if (_126_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_126_i1 = (_126_i1).Plus(_dafny.One)
				var _127_v6 _dafny.Map
				_ = _127_v6
				_127_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _117_v0)
				var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
				_ = _index21
				(_121_v4).ArraySet1(((_127_v6).Merge(_127_v6)).Cardinality(), (_index21).Int())
				var _128_v7 bool
				_ = _128_v7
				_128_v7 = true
				var _129_v8 _dafny.CodePoint
				_ = _129_v8
				_129_v8 = _dafny.CodePoint('i')
				var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
				_ = _index22
				(_121_v4).ArraySet1(Companion_Default___.Fm0((func() _dafny.Int {
					if _128_v7 {
						return _117_v0
					}
					return _dafny.IntOfInt64(-650)
				})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if true {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(438))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg6 _dafny.Int) interface{} {
								return coer6(arg6)
							}
						}(func(_130_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						}))
					}
					return _119_v2
				})(), (Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if true {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(438))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg7 _dafny.Int) interface{} {
								return coer7(arg7)
							}
						}(func(_131_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						}))
					}
					return _119_v2
				})()).Cardinality()))).Uint32(), _129_v8)).Cardinality()), _125_globalState), (_index22).Int())
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
				_ = _index23
				(_121_v4).ArraySet1(Companion_Default___.Fm0(_117_v0, _117_v0, _125_globalState), (_index23).Int())
				if (_117_v0).Cmp(_117_v0) < 0 {
					var _132_v9 _dafny.MultiSet
					_ = _132_v9
					_132_v9 = _dafny.MultiSetOf(_117_v0)
					var _133_v10 _dafny.Sequence
					_ = _133_v10
					_133_v10 = _dafny.SeqOf(_119_v2)
					var _134_v11 _dafny.MultiSet
					_ = _134_v11
					_134_v11 = _dafny.MultiSetOf(Companion_Default___.Fm0((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), (_132_v9).Cardinality(), _125_globalState), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_133_v10).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_132_v9).Contains(_117_v0) {
							return (_132_v9).Multiplicity(_117_v0)
						}
						return _117_v0
					})(), _dafny.IntOfUint32((_133_v10).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _117_v0)
					var _135_v12 _dafny.Array
					_ = _135_v12
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(20)
					_ = _len0_5
					var _nw20 _dafny.Array
					_ = _nw20
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw20 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.Sequence = (func(_136_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_137_i3 _dafny.Int) _dafny.Sequence {
								return _136_v2
							}
						})(_119_v2)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw20 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw20).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw20).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_135_v12 = _nw20
					var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_135_v12), 0))
					_ = _index24
					(_135_v12).ArraySet1(_119_v2, (_index24).Int())
					var _138_v13 _dafny.Sequence
					_ = _138_v13
					_138_v13 = _119_v2
					var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
					_ = _index25
					var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_135_v12), 0))
					_ = _index26
					var _rhs17 _dafny.Int = _117_v0
					_ = _rhs17
					var _rhs18 _dafny.MultiSet = _134_v11
					_ = _rhs18
					var _rhs19 _dafny.Sequence = _119_v2
					_ = _rhs19
					var _rhs20 bool = _128_v7
					_ = _rhs20
					var _rhs21 _dafny.Sequence = (_138_v13)
					_ = _rhs21
					var _lhs24 _dafny.Array = _121_v4
					_ = _lhs24
					var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
					_ = _lhs25
					var _lhs26 *GlobalState = _125_globalState
					_ = _lhs26
					var _lhs27 _dafny.Array = _135_v12
					_ = _lhs27
					var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_135_v12), 0))
					_ = _lhs28
					(_lhs24).ArraySet1(_rhs17, (_lhs25).Int())
					_132_v9 = _rhs18
					_119_v2 = _rhs19
					_lhs26.F2 = _rhs20
					(_lhs27).ArraySet1(_rhs21, (_lhs28).Int())
					var _139_v14 _dafny.Sequence
					_ = _139_v14
					_139_v14 = _dafny.SeqOf((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _117_v0)
					(_125_globalState).F2 = !_dafny.Companion_Sequence_.Contains(_139_v14, _117_v0)
					var _140_v15 _dafny.MultiSet
					_ = _140_v15
					_140_v15 = _dafny.MultiSetOf(_119_v2, (_135_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_135_v12), 0))).Int()).(_dafny.Sequence))
					var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_135_v12), 0))
					_ = _index27
					var _rhs22 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_119_v2, (Companion_Default___.SafeIndex((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_119_v2).Cardinality()))).Uint32(), _129_v8), (_138_v13))
					_ = _rhs22
					var _rhs23 bool = (_117_v0).Cmp((_117_v0).Times((_139_v14).Select((Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_139_v14).Cardinality()))).Uint32()).(_dafny.Int))) != 0
					_ = _rhs23
					var _rhs24 bool = (_140_v15).IsProperSubsetOf(_140_v15)
					_ = _rhs24
					var _rhs25 _dafny.Int = _117_v0
					_ = _rhs25
					var _lhs29 _dafny.Array = _135_v12
					_ = _lhs29
					var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_135_v12), 0))
					_ = _lhs30
					var _lhs31 *GlobalState = _125_globalState
					_ = _lhs31
					var _lhs32 *GlobalState = _125_globalState
					_ = _lhs32
					var _lhs33 *GlobalState = _125_globalState
					_ = _lhs33
					(_lhs29).ArraySet1(_rhs22, (_lhs30).Int())
					_lhs31.F2 = _rhs23
					_lhs32.F2 = _rhs24
					_lhs33.F4 = _rhs25
					var _141_v16 D0
					_ = _141_v16
					_141_v16 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)), _128_v7, _dafny.IntOfUint32((_119_v2).Cardinality()))
					var _142_v17 *C0
					_ = _142_v17
					var _nw21 *C0 = New_C0_()
					_ = _nw21
					_nw21.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v16, (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)))
					_142_v17 = _nw21
					(_125_globalState).F3 = ((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)).Minus((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int))
				} else {
					Companion_Default___.M0(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("lkmnn"), _119_v2), _125_globalState)
					var _143_v18 _dafny.Array
					_ = _143_v18
					var _nwElement0_4 _dafny.Array = (func() _dafny.Array {
						if _128_v7 {
							return _122_v5
						}
						return _122_v5
					})()
					_ = _nwElement0_4
					var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(3))
					_ = _nw22
					(_nw22).ArraySet1(_nwElement0_4, 0)
					(_nw22).ArraySet1(_122_v5, 1)
					(_nw22).ArraySet1(_122_v5, 2)
					_143_v18 = _nw22
					var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_143_v18), 0))
					_ = _index28
					(_143_v18).ArraySet1(_122_v5, (_index28).Int())
					var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_143_v18), 0))
					_ = _index29
					(_143_v18).ArraySet1(_122_v5, (_index29).Int())
					var _144_v19 D0
					_ = _144_v19
					_144_v19 = Companion_D0_.Create_DC2_()
					_144_v19 = _144_v19
					var _145_v20 _dafny.Array
					_ = _145_v20
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
					_ = _len0_6
					var _nw23 _dafny.Array
					_ = _nw23
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw23 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Sequence = (func(_146_v7 bool, _147_v8 _dafny.CodePoint, _148_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_149_i4 _dafny.Int) _dafny.Sequence {
								return (func() _dafny.Sequence {
									if _146_v7 {
										return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg8 _dafny.Int) interface{} {
												return coer8(arg8)
											}
										}((func(_150_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
											return func(_151_i5 _dafny.Int) _dafny.CodePoint {
												return _150_v8
											}
										})(_147_v8)))
									}
									return _148_v2
								})()
							}
						})(_128_v7, _129_v8, _119_v2)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw23 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw23).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw23).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_145_v20 = _nw23
					var _152_v21 _dafny.Map
					_ = _152_v21
					_152_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_128_v7), _dafny.UnicodeSeqOfUtf8Bytes("hcnuuyv"))
					var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_145_v20), 0))
					_ = _index30
					(_145_v20).ArraySet1((func() _dafny.Sequence {
						if (_152_v21).Contains(_128_v7) {
							return (_152_v21).Get(_128_v7).(_dafny.Sequence)
						}
						return Companion_Default___.Fm1(_125_globalState)
					})(), (_index30).Int())
					var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_145_v20), 0))
					_ = _index31
					(_145_v20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_119_v2, _119_v2), (_index31).Int())
					var _153_v22 _dafny.MultiSet
					_ = _153_v22
					_153_v22 = _dafny.MultiSetOf((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int))
					Companion_Default___.M0(!(_153_v22).Contains(_dafny.IntOfInt64(577)), _125_globalState)
				}
				if (_117_v0).Cmp((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)) == 0 {
					(_125_globalState).F3 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _117_v0)).Plus(_117_v0))))
					(_125_globalState).F4 = _117_v0
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
					_ = _index32
					(_121_v4).ArraySet1((_dafny.Zero).Minus(_117_v0), (_index32).Int())
					var _154_v23 _dafny.Array
					_ = _154_v23
					var _len0_7 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_7
					var _nw24 _dafny.Array
					_ = _nw24
					if _len0_7.Cmp(_dafny.Zero) == 0 {
						_nw24 = _dafny.NewArray(_len0_7)
					} else {
						var _init7 func(_dafny.Int) _dafny.Map = (func(_155_v7 bool, _156_v4 _dafny.Array) func(_dafny.Int) _dafny.Map {
							return func(_157_i6 _dafny.Int) _dafny.Map {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-2))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_155_v7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_156_v4), 0))).Int()).(_dafny.Int), _155_v7)).Cardinality()))
							}
						})(_128_v7, _121_v4)
						_ = _init7
						var _element0_7 = _init7(_dafny.Zero)
						_ = _element0_7
						_nw24 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
						(_nw24).ArraySet1(_element0_7, 0)
						var _nativeLen0_7 = (_len0_7).Int()
						_ = _nativeLen0_7
						for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
							(_nw24).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
						}
					}
					_154_v23 = _nw24
					var _158_v24 _dafny.Map
					_ = _158_v24
					_158_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v7, _117_v0)
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_154_v23), 0))
					_ = _index33
					(_154_v23).ArraySet1(_158_v24, (_index33).Int())
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_154_v23), 0))
					_ = _index34
					(_154_v23).ArraySet1((_158_v24).Update(_128_v7, _dafny.IntOfInt64(410)), (_index34).Int())
					var _159_v25 _dafny.Map
					_ = _159_v25
					_159_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _dafny.MultiSetOf(_dafny.IntOfInt64(300), _dafny.IntOfInt64(58)))
					var _160_v26 _dafny.Sequence
					_ = _160_v26
					_160_v26 = _dafny.SeqOf(_117_v0, _117_v0)
					var _161_v27 D2
					_ = _161_v27
					_161_v27 = Companion_D2_.Create_DC5_(_dafny.MultiSetFromSeq(_160_v26))
					_159_v25 = (_159_v25).Update(_117_v0, (_161_v27).Dtor_cf5())
				} else {
					var _162_v29 _dafny.MultiSet
					_ = _162_v29
					_162_v29 = _dafny.MultiSetOf(_128_v7)
					var _163_v30 D0
					_ = _163_v30
					_163_v30 = Companion_D0_.Create_DC1_((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _128_v7, (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int))
					var _164_v31 _dafny.Map
					_ = _164_v31
					_164_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v30, _117_v0)
					var _165_v32 T0
					_ = _165_v32
					var _nw25 *C0 = New_C0_()
					_ = _nw25
					_nw25.Ctor__(_164_v31)
					_165_v32 = _nw25
					var _166_v33 _dafny.MultiSet
					_ = _166_v33
					_166_v33 = _dafny.MultiSetOf(_165_v32, _165_v32)
					var _167_v34 D0
					_ = _167_v34
					_167_v34 = Companion_D0_.Create_DC1_((_166_v33).Cardinality(), _128_v7, _117_v0)
					var _168_v35 _dafny.MultiSet
					_ = _168_v35
					_168_v35 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _128_v7, (func() _dafny.Int {
						if (_162_v29).Contains(_128_v7) {
							return (_162_v29).Multiplicity(_128_v7)
						}
						return _117_v0
					})()), _167_v34)
					var _169_v36 T0
					_ = _169_v36
					var _nw26 *C0 = New_C0_()
					_ = _nw26
					_nw26.Ctor__(func() _dafny.Map {
						var _coll7 = _dafny.NewMapBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate((_168_v35).Elements()); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _170_v28 D0
							_170_v28 = interface{}(_compr_7).(D0)
							if (_168_v35).Contains(_170_v28) {
								_coll7.Add(_170_v28, (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int))
							}
						}
						return _coll7.ToMap()
					}())
					_169_v36 = _nw26
					var _171_v37 _dafny.Array
					_ = _171_v37
					var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
					_ = _nw27
					_171_v37 = _nw27
					var _172_v38 _dafny.Map
					_ = _172_v38
					_172_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v32, _171_v37)
					var _173_v39 *C0
					_ = _173_v39
					var _nw28 *C0 = New_C0_()
					_ = _nw28
					_nw28.Ctor__(_164_v31)
					_173_v39 = _nw28
					var _174_v40 _dafny.Set
					_ = _174_v40
					_174_v40 = _dafny.SetOf(_173_v39)
					var _rhs26 bool = _128_v7
					_ = _rhs26
					var _rhs27 bool = _128_v7
					_ = _rhs27
					var _rhs28 bool = ((_174_v40).Difference(_174_v40)).IsDisjointFrom((_174_v40).Intersection(_dafny.SetOf(_173_v39, _173_v39, _173_v39)))
					_ = _rhs28
					var _rhs29 bool = _128_v7
					_ = _rhs29
					var _rhs30 _dafny.Map = (_172_v38).Merge((_172_v38))
					_ = _rhs30
					var _lhs34 *GlobalState = _125_globalState
					_ = _lhs34
					var _lhs35 *GlobalState = _125_globalState
					_ = _lhs35
					_128_v7 = _rhs26
					_lhs34.F2 = _rhs27
					_128_v7 = _rhs28
					_lhs35.F2 = _rhs29
					_172_v38 = _rhs30
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
					_ = _index35
					var _rhs31 bool = _128_v7
					_ = _rhs31
					var _rhs32 _dafny.Int = ((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)))
					_ = _rhs32
					var _lhs36 *GlobalState = _125_globalState
					_ = _lhs36
					var _lhs37 _dafny.Array = _121_v4
					_ = _lhs37
					var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))
					_ = _lhs38
					_lhs36.F2 = _rhs31
					(_lhs37).ArraySet1(_rhs32, (_lhs38).Int())
					var _175_v41 _dafny.Set
					_ = _175_v41
					_175_v41 = _dafny.SetOf(_128_v7, _128_v7, _128_v7, _128_v7, _128_v7)
					var _176_v42 _dafny.Sequence
					_ = _176_v42
					_176_v42 = _dafny.SeqOf(_175_v41, _175_v41, _175_v41, _175_v41)
					var _177_v43 _dafny.Map
					_ = _177_v43
					_177_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v7, _dafny.MultiSetFromSeq(_176_v42))
					var _178_v44 _dafny.MultiSet
					_ = _178_v44
					_178_v44 = _dafny.MultiSetOf(_175_v41, _175_v41, _175_v41)
					var _179_v45 _dafny.Map
					_ = _179_v45
					_179_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_175_v41).Cardinality(), _178_v44)
					_177_v43 = (_177_v43).Update(_128_v7, (func() _dafny.MultiSet {
						if (_179_v45).Contains(_117_v0) {
							return (_179_v45).Get(_117_v0).(_dafny.MultiSet)
						}
						return _178_v44
					})())
					var _180_v46 _dafny.Array
					_ = _180_v46
					var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(4))
					_ = _nw29
					_180_v46 = _nw29
					_180_v46 = _180_v46
					var _181_v47 *C0
					_ = _181_v47
					var _nw30 *C0 = New_C0_()
					_ = _nw30
					_nw30.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_128_v7, _125_globalState), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)))
					_181_v47 = _nw30
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _182_v48 bool
	_ = _182_v48
	_182_v48 = false
	var _183_v49 _dafny.Sequence
	_ = _183_v49
	_183_v49 = _dafny.SeqOf(_182_v48)
	if (_182_v48) == (_dafny.Companion_Sequence_.IsProperPrefixOf(_183_v49, _dafny.SeqOf(true))) {
		var _184_v50 _dafny.MultiSet
		_ = _184_v50
		_184_v50 = _dafny.MultiSetOf(_119_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eueosg"), _119_v2), _119_v2)
		_117_v0 = (_184_v50).Cardinality()
		var _185_v51 _dafny.CodePoint
		_ = _185_v51
		_185_v51 = _dafny.CodePoint('y')
		_185_v51 = _185_v51
		var _186_v52 D0
		_ = _186_v52
		_186_v52 = Companion_D0_.Create_DC1_(_117_v0, _182_v48, Companion_Default___.Fm0(_117_v0, _117_v0, _125_globalState))
		var _187_v54 _dafny.Map
		_ = _187_v54
		_187_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v52, _117_v0)
		var _188_v55 *C0
		_ = _188_v55
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v52, _117_v0)).Merge(func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_187_v54).Keys().Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _189_v53 D0
				_189_v53 = interface{}(_compr_8).(D0)
				if (_187_v54).Contains(_189_v53) {
					_coll8.Add(_189_v53, (_186_v52).Dtor_cf3())
				}
			}
			return _coll8.ToMap()
		}()))
		_188_v55 = _nw31
		var _190_v56 _dafny.Sequence
		_ = _190_v56
		_190_v56 = _dafny.SeqOf(_117_v0)
		if !_dafny.Companion_Sequence_.Contains(_190_v56, _117_v0) {
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))
			_ = _index36
			(_122_v5).ArraySet1(_182_v48, (_index36).Int())
			var _191_v57 _dafny.Map
			_ = _191_v57
			_191_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _117_v0)
			var _192_v58 T0
			_ = _192_v58
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__(_187_v54)
			_192_v58 = _nw32
			var _193_v59 _dafny.Sequence
			_ = _193_v59
			_193_v59 = _dafny.SeqOf(_192_v58, _192_v58)
			var _194_v60 _dafny.Map
			_ = _194_v60
			_194_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_191_v57).Cardinality(), _117_v0), _dafny.Companion_Sequence_.Concatenate((Companion_D4_.Create_DC9_(_193_v59)).Dtor_cf11(), _193_v59))
			var _195_v61 _dafny.Map
			_ = _195_v61
			_195_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v58, _117_v0)
			var _196_v62 _dafny.Map
			_ = _196_v62
			_196_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v58, (func() _dafny.Map {
				if _182_v48 {
					return _195_v61
				}
				return _195_v61
			})())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))
			_ = _index37
			var _rhs33 bool = _182_v48
			_ = _rhs33
			var _rhs34 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _193_v59)).Merge(_194_v60)
			_ = _rhs34
			var _rhs35 _dafny.Map = _196_v62
			_ = _rhs35
			var _lhs39 _dafny.Array = _122_v5
			_ = _lhs39
			var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))
			_ = _lhs40
			(_lhs39).ArraySet1(_rhs33, (_lhs40).Int())
			_194_v60 = _rhs34
			_196_v62 = _rhs35
			var _197_v63 _dafny.Map
			_ = _197_v63
			_197_v63 = _187_v54
			var _198_v64 *C0
			_ = _198_v64
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__((_197_v63))
			_198_v64 = _nw33
			_198_v64 = _198_v64
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))
			_ = _index38
			(_122_v5).ArraySet1((func() bool {
				if (_117_v0).Cmp((_dafny.MultiSetOf(_117_v0)).Cardinality()) == 0 {
					return (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)
				}
				return _182_v48
			})(), (_index38).Int())
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__((_188_v55).F15())
			_192_v58 = _nw34
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))
			_ = _index39
			(_122_v5).ArraySet1((func() bool {
				if !_dafny.Companion_Sequence_.Contains(_119_v2, _185_v51) {
					return false
				}
				return (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)
			})(), (_index39).Int())
		} else {
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))
			_ = _index40
			(_121_v4).ArraySet1(_dafny.IntOfUint32((_190_v56).Cardinality()), (_index40).Int())
			var _199_v65 _dafny.MultiSet
			_ = _199_v65
			_199_v65 = _dafny.MultiSetOf(_182_v48)
			var _200_v66 _dafny.Map
			_ = _200_v66
			_200_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _dafny.MultiSetOf(_182_v48))
			var _201_v67 _dafny.MultiSet
			_ = _201_v67
			_201_v67 = _dafny.MultiSetOf(_117_v0)
			var _202_v68 _dafny.Array
			_ = _202_v68
			var _nwElement0_5 _dafny.MultiSet = _dafny.MultiSetOf(!(true), _182_v48)
			_ = _nwElement0_5
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(22))
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_5, 0)
			(_nw35).ArraySet1(_199_v65, 1)
			(_nw35).ArraySet1(_199_v65, 2)
			(_nw35).ArraySet1(_199_v65, 3)
			(_nw35).ArraySet1(_199_v65, 4)
			(_nw35).ArraySet1(_dafny.MultiSetFromSeq(_183_v49), 5)
			(_nw35).ArraySet1(_199_v65, 6)
			(_nw35).ArraySet1(_199_v65, 7)
			(_nw35).ArraySet1(_199_v65, 8)
			(_nw35).ArraySet1(_dafny.MultiSetFromSeq(_183_v49), 9)
			(_nw35).ArraySet1((func() _dafny.MultiSet {
				if (_200_v66).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}((func(_203_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_204_i7 _dafny.Int) _dafny.Int {
						return _203_v0
					}
				})(_117_v0))), (Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_205_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_206_i7 _dafny.Int) _dafny.Int {
						return _205_v0
					}
				})(_117_v0)))).Cardinality()))).Uint32(), _117_v0)).Cardinality())) {
					return (_200_v66).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}((func(_207_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_208_i7 _dafny.Int) _dafny.Int {
							return _207_v0
						}
					})(_117_v0))), (Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}((func(_209_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_210_i7 _dafny.Int) _dafny.Int {
							return _209_v0
						}
					})(_117_v0)))).Cardinality()))).Uint32(), _117_v0)).Cardinality())).(_dafny.MultiSet)
				}
				return _199_v65
			})(), 10)
			(_nw35).ArraySet1(_199_v65, 11)
			(_nw35).ArraySet1(_199_v65, 12)
			(_nw35).ArraySet1(_199_v65, 13)
			(_nw35).ArraySet1(_199_v65, 14)
			(_nw35).ArraySet1(_199_v65, 15)
			(_nw35).ArraySet1(_199_v65, 16)
			(_nw35).ArraySet1(_dafny.MultiSetFromSeq(_183_v49), 17)
			(_nw35).ArraySet1(_199_v65, 18)
			(_nw35).ArraySet1(Companion_Default___.Fm3(true, _201_v67, _182_v48, _125_globalState), 19)
			(_nw35).ArraySet1(_dafny.MultiSetFromSeq(_183_v49), 20)
			(_nw35).ArraySet1(_199_v65, 21)
			_202_v68 = _nw35
			var _211_v69 _dafny.Map
			_ = _211_v69
			_211_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _122_v5)
			var _212_v70 _dafny.Map
			_ = _212_v70
			_212_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_202_v68, (func() _dafny.Array {
				if (_211_v69).Contains(_182_v48) {
					return (_211_v69).Get(_182_v48).(_dafny.Array)
				}
				return _122_v5
			})())
			var _213_v71 _dafny.Sequence
			_ = _213_v71
			_213_v71 = _dafny.SeqOf(_122_v5)
			var _214_v72 _dafny.Array
			_ = _214_v72
			var _nwElement0_6 _dafny.Array = _122_v5
			_ = _nwElement0_6
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(25))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_6, 0)
			(_nw36).ArraySet1(_122_v5, 1)
			(_nw36).ArraySet1((func() _dafny.Array {
				if (_212_v70).Contains(_202_v68) {
					return (_212_v70).Get(_202_v68).(_dafny.Array)
				}
				return _122_v5
			})(), 2)
			(_nw36).ArraySet1(_122_v5, 3)
			(_nw36).ArraySet1(_122_v5, 4)
			(_nw36).ArraySet1(_122_v5, 5)
			(_nw36).ArraySet1(_122_v5, 6)
			(_nw36).ArraySet1(_122_v5, 7)
			(_nw36).ArraySet1(_122_v5, 8)
			(_nw36).ArraySet1(_122_v5, 9)
			(_nw36).ArraySet1(_122_v5, 10)
			(_nw36).ArraySet1((_213_v71).Select((Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_213_v71).Cardinality()))).Uint32()).(_dafny.Array), 11)
			(_nw36).ArraySet1(_122_v5, 12)
			(_nw36).ArraySet1(_122_v5, 13)
			(_nw36).ArraySet1(_122_v5, 14)
			(_nw36).ArraySet1(_122_v5, 15)
			(_nw36).ArraySet1((Companion_D6_.Create_DC14_(_122_v5)).Dtor_cf15(), 16)
			(_nw36).ArraySet1(_122_v5, 17)
			(_nw36).ArraySet1((func() _dafny.Array {
				if _182_v48 {
					return _122_v5
				}
				return _122_v5
			})(), 18)
			(_nw36).ArraySet1(_122_v5, 19)
			(_nw36).ArraySet1(_122_v5, 20)
			(_nw36).ArraySet1(_122_v5, 21)
			(_nw36).ArraySet1((func() _dafny.Array {
				if _182_v48 {
					return _122_v5
				}
				return _122_v5
			})(), 22)
			(_nw36).ArraySet1(_122_v5, 23)
			(_nw36).ArraySet1(_122_v5, 24)
			_214_v72 = _nw36
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_214_v72), 0))
			_ = _index41
			(_214_v72).ArraySet1(_122_v5, (_index41).Int())
			var _215_v73 _dafny.Map
			_ = _215_v73
			_215_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v2, _117_v0)
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))
			_ = _index42
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_214_v72), 0))
			_ = _index43
			var _rhs36 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_117_v0), _117_v0), (func() _dafny.Int {
				if (_215_v73).Contains(_119_v2) {
					return (_215_v73).Get(_119_v2).(_dafny.Int)
				}
				return _117_v0
			})())
			_ = _rhs36
			var _rhs37 _dafny.Sequence = _119_v2
			_ = _rhs37
			var _rhs38 _dafny.Int = (_dafny.IntOfInt64(209)).Minus(_dafny.IntOfInt64(602))
			_ = _rhs38
			var _rhs39 _dafny.Array = _122_v5
			_ = _rhs39
			var _lhs41 *GlobalState = _125_globalState
			_ = _lhs41
			var _lhs42 _dafny.Array = _121_v4
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))
			_ = _lhs43
			var _lhs44 _dafny.Array = _214_v72
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_214_v72), 0))
			_ = _lhs45
			_lhs41.F3 = _rhs36
			_119_v2 = _rhs37
			(_lhs42).ArraySet1(_rhs38, (_lhs43).Int())
			(_lhs44).ArraySet1(_rhs39, (_lhs45).Int())
			var _216_v74 *C0
			_ = _216_v74
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v52, _117_v0))
			_216_v74 = _nw37
			var _217_v75 _dafny.Map
			_ = _217_v75
			_217_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _119_v2)
			_217_v75 = (_217_v75).Update((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(644))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_218_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_219_i8 _dafny.Int) _dafny.CodePoint {
					return _218_v51
				}
			})(_185_v51))))
			(_125_globalState).F4 = (_dafny.Zero).Minus(((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(820)))
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))
			_ = _index44
			var _rhs40 _dafny.Sequence = _190_v56
			_ = _rhs40
			var _rhs41 _dafny.Int = (Companion_Default___.Fm4(!(_182_v48), _182_v48, _125_globalState)).Cardinality()
			_ = _rhs41
			var _lhs46 _dafny.Array = _121_v4
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_121_v4), 0))
			_ = _lhs47
			_190_v56 = _rhs40
			(_lhs46).ArraySet1(_rhs41, (_lhs47).Int())
		}
		(_125_globalState).F2 = _182_v48
	} else {
		var _220_v76 _dafny.CodePoint
		_ = _220_v76
		_220_v76 = _dafny.CodePoint('x')
		var _221_v77 _dafny.Map
		_ = _221_v77
		_221_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v76, _dafny.Companion_Sequence_.Equal(_119_v2, _119_v2))
		var _222_v78 _dafny.Map
		_ = _222_v78
		_222_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, _182_v48)
		var _223_v79 D0
		_ = _223_v79
		_223_v79 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(116), _182_v48, _117_v0)
		var _224_v80 _dafny.Map
		_ = _224_v80
		_224_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v79, _117_v0)
		var _225_v81 *C0
		_ = _225_v81
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__((_224_v80).Update(_223_v79, _117_v0))
		_225_v81 = _nw38
		var _226_v82 _dafny.Map
		_ = _226_v82
		_226_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_225_v81, _225_v81)
		var _227_v83 _dafny.Map
		_ = _227_v83
		_227_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_225_v81, _225_v81), _117_v0)
		var _228_v84 _dafny.Sequence
		_ = _228_v84
		_228_v84 = _dafny.SeqOf(_dafny.IntOfUint32((_183_v49).Cardinality()))
		var _rhs42 _dafny.Map = (func() _dafny.Map {
			if !(_227_v83).Contains((_226_v82).Update(_225_v81, _225_v81)) {
				return _221_v77
			}
			return _221_v77
		})()
		_ = _rhs42
		var _rhs43 _dafny.Map = (_222_v78).Update(_117_v0, (func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(17), _dafny.IntOfInt64(599))); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _229_v86 _dafny.Int
				_229_v86 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(17)).Cmp(_229_v86) <= 0) && ((_229_v86).Cmp(_dafny.IntOfInt64(599)) < 0) {
					_coll9.Add((_229_v86).Times(_dafny.IntOfInt64(904)))
				}
			}
			return _coll9.ToSet()
		}()).IsProperSubsetOf(func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_dafny.SetOf((_228_v84).Select((Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_228_v84).Cardinality()))).Uint32()).(_dafny.Int))).Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _230_v85 _dafny.Int
				_230_v85 = interface{}(_compr_10).(_dafny.Int)
				if (_dafny.SetOf((_228_v84).Select((Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_228_v84).Cardinality()))).Uint32()).(_dafny.Int))).Contains(_230_v85) {
					_coll10.Add(Companion_Default___.SafeModuloInt(_230_v85, _dafny.IntOfInt64(-93)))
				}
			}
			return _coll10.ToSet()
		}()))
		_ = _rhs43
		_221_v77 = _rhs42
		_222_v78 = _rhs43
		var _231_v87 _dafny.Map
		_ = _231_v87
		_231_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_119_v2, _dafny.UnicodeSeqOfUtf8Bytes("jlvlgk")), _121_v4)
		_231_v87 = (_231_v87).Update(Companion_Default___.Fm5(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _117_v0), _182_v48, _117_v0, _117_v0, _125_globalState), _121_v4)
		Companion_Default___.M0(_182_v48, _125_globalState)
		var _232_v88 _dafny.Map
		_ = _232_v88
		_232_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v4, _117_v0)
		_232_v88 = _232_v88
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_121_v4), 0))
		_ = _index45
		(_121_v4).ArraySet1((_117_v0).Times(_117_v0), (_index45).Int())
		var _233_v89 _dafny.Array
		_ = _233_v89
		var _len0_8 _dafny.Int = _dafny.One
		_ = _len0_8
		var _nw39 _dafny.Array
		_ = _nw39
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw39 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Sequence = (func(_234_v84 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_235_i9 _dafny.Int) _dafny.Sequence {
					return _234_v84
				}
			})(_228_v84)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw39 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw39).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw39).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_233_v89 = _nw39
		var _236_v90 _dafny.Map
		_ = _236_v90
		_236_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6(Companion_Default___.Fm6(_182_v48, (_dafny.SetOf(_121_v4, _121_v4, _121_v4, (Companion_D7_.Create_DC17_(_121_v4)).Dtor_cf19(), _121_v4)).Cardinality(), _125_globalState), _117_v0, _125_globalState), _dafny.SeqOf((_dafny.Zero).Minus(_117_v0), _117_v0))
		var _237_v91 _dafny.Map
		_ = _237_v91
		_237_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, !(_182_v48))
		var _238_v92 D2
		_ = _238_v92
		_238_v92 = Companion_D2_.Create_DC6_(_237_v91, _182_v48)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_233_v89), 0))
		_ = _index46
		(_233_v89).ArraySet1((func() _dafny.Sequence {
			if (_236_v90).Contains(!((_238_v92).Dtor_cf7())) {
				return (_236_v90).Get(!((_238_v92).Dtor_cf7())).(_dafny.Sequence)
			}
			return _dafny.SeqOf(_117_v0)
		})(), (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_121_v4), 0))
		_ = _index47
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_233_v89), 0))
		_ = _index48
		var _rhs44 _dafny.Int = (_117_v0).Times(_117_v0)
		_ = _rhs44
		var _rhs45 _dafny.Int = ((_117_v0).Times(_117_v0)).Plus(_117_v0)
		_ = _rhs45
		var _rhs46 _dafny.Sequence = _228_v84
		_ = _rhs46
		var _lhs48 _dafny.Array = _121_v4
		_ = _lhs48
		var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_121_v4), 0))
		_ = _lhs49
		var _lhs50 *GlobalState = _125_globalState
		_ = _lhs50
		var _lhs51 _dafny.Array = _233_v89
		_ = _lhs51
		var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_233_v89), 0))
		_ = _lhs52
		(_lhs48).ArraySet1(_rhs44, (_lhs49).Int())
		_lhs50.F3 = _rhs45
		(_lhs51).ArraySet1(_rhs46, (_lhs52).Int())
	}
	var _239_i10 _dafny.Int
	_ = _239_i10
	_239_i10 = _dafny.Zero
	{
		for false {
			{
				if (_239_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_239_i10 = (_239_i10).Plus(_dafny.One)
				var _240_v93 D0
				_ = _240_v93
				_240_v93 = Companion_D0_.Create_DC1_(_117_v0, _182_v48, _117_v0)
				var _241_v94 _dafny.Map
				_ = _241_v94
				_241_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v93, (_dafny.Zero).Minus(_117_v0))
				var _242_v95 _dafny.Map
				_ = _242_v95
				_242_v95 = _241_v94
				var _source1 _dafny.Map = _242_v95
				_ = _source1
				var _243___mcc_h0 _dafny.Map = _source1
				_ = _243___mcc_h0
				var _244_cf14 _dafny.Map = _243___mcc_h0
				_ = _244_cf14
				var _245_v96 _dafny.Map
				_ = _245_v96
				_245_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _182_v48)
				_182_v48 = (((_245_v96).Cardinality()).Cmp(_117_v0) <= 0) == (_182_v48)
				Companion_Default___.M0(_182_v48, _125_globalState)
				_117_v0 = (func() _dafny.Int {
					if !(_182_v48) {
						return _117_v0
					}
					return _117_v0
				})()
				var _246_v97 D2
				_ = _246_v97
				_246_v97 = Companion_D2_.Create_DC6_(_245_v96, _182_v48)
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_122_v5), 0))
				_ = _index49
				(_122_v5).ArraySet1(((_246_v97).Dtor_cf7()) == (_182_v48), (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_122_v5), 0))
				_ = _index50
				(_122_v5).ArraySet1(_182_v48, (_index50).Int())
				(_125_globalState).F2 = !(_182_v48)
				var _rhs47 _dafny.Int = _dafny.IntOfUint32((_119_v2).Cardinality())
				_ = _rhs47
				var _rhs48 _dafny.Int = _117_v0
				_ = _rhs48
				var _lhs53 *GlobalState = _125_globalState
				_ = _lhs53
				_117_v0 = _rhs47
				_lhs53.F4 = _rhs48
				if _182_v48 {
					var _247_v98 D0
					_ = _247_v98
					_247_v98 = Companion_D0_.Create_DC0_(_117_v0)
					(_125_globalState).F2 = ((_247_v98).Dtor_cf0()).Cmp(_117_v0) > 0
					_119_v2 = _dafny.UnicodeSeqOfUtf8Bytes("uuglv")
					var _248_v99 *C0
					_ = _248_v99
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__(_241_v94)
					_248_v99 = _nw40
					(_125_globalState).F4 = (_dafny.IntOfUint32((Companion_Default___.Fm1(_125_globalState)).Cardinality())).Plus(_dafny.IntOfInt64(759))
					var _249_v100 _dafny.Sequence
					_ = _249_v100
					_249_v100 = _dafny.SeqOf(_117_v0, _dafny.IntOfInt64(609), _117_v0)
					var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_121_v4), 0))
					_ = _index51
					(_121_v4).ArraySet1((func() _dafny.Int {
						if _182_v48 {
							return _117_v0
						}
						return _dafny.IntOfUint32((_249_v100).Cardinality())
					})(), (_index51).Int())
					var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_121_v4), 0))
					_ = _index52
					(_121_v4).ArraySet1(_117_v0, (_index52).Int())
				} else {
					var _250_v101 _dafny.Map
					_ = _250_v101
					_250_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_117_v0).Times(_117_v0), (_117_v0).Minus(_117_v0))
					_250_v101 = _250_v101
					(_125_globalState).F3 = _117_v0
					Companion_Default___.M0(_dafny.Companion_Sequence_.IsPrefixOf(_183_v49, _dafny.SeqOf(_182_v48)), _125_globalState)
					(_125_globalState).F3 = _117_v0
					var _pat_let_tv1 = _117_v0
					_ = _pat_let_tv1
					var _251_v102 *C0
					_ = _251_v102
					var _nw41 *C0 = New_C0_()
					_ = _nw41
					_nw41.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let4_0 D0) D0 {
						return func(_252_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let5_0 _dafny.Int) D0 {
								return func(_253_dt__update_hcf3_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC1_((_252_dt__update__tmp_h0).Dtor_cf1(), (_252_dt__update__tmp_h0).Dtor_cf2(), _253_dt__update_hcf3_h0)
								}(_pat_let5_0)
							}(_pat_let_tv1)
						}(_pat_let4_0)
					}(_240_v93), _117_v0))
					_251_v102 = _nw41
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _source2 D0 = Companion_D0_.Create_DC0_(_117_v0)
	_ = _source2
	if _source2.Is_DC1() {
		var _254___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
		_ = _254___mcc_h1
		var _255___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf2
		_ = _255___mcc_h2
		var _256___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
		_ = _256___mcc_h3
		var _257_cf3 _dafny.Int = _256___mcc_h3
		_ = _257_cf3
		var _258_cf2 bool = _255___mcc_h2
		_ = _258_cf2
		var _259_cf1 _dafny.Int = _254___mcc_h1
		_ = _259_cf1
		var _260_v103 _dafny.CodePoint
		_ = _260_v103
		_260_v103 = _dafny.CodePoint('o')
		_260_v103 = _dafny.CodePoint('u')
		var _261_v104 _dafny.Array
		_ = _261_v104
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw42
		_261_v104 = _nw42
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_261_v104), 0))
		_ = _index53
		(_261_v104).ArraySet1(_121_v4, (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_261_v104), 0))
		_ = _index54
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
		_ = _nw43
		(_261_v104).ArraySet1(_nw43, (_index54).Int())
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_122_v5), 0))
		_ = _index55
		(_122_v5).ArraySet1(_182_v48, (_index55).Int())
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_122_v5), 0))
		_ = _index56
		(_122_v5).ArraySet1(_182_v48, (_index56).Int())
		_259_cf1 = Companion_Default___.Fm0(Companion_Default___.SafeModuloInt(_257_cf3, _259_cf1), Companion_Default___.Fm0(_257_cf3, _257_cf3, _125_globalState), _125_globalState)
	} else if _source2.Is_DC2() {
		(_125_globalState).F2 = _182_v48
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_122_v5), 0))
		_ = _index57
		(_122_v5).ArraySet1(false, (_index57).Int())
		var _262_v105 _dafny.Map
		_ = _262_v105
		_262_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _182_v48)
		var _263_v106 _dafny.MultiSet
		_ = _263_v106
		_263_v106 = _dafny.MultiSetOf((_262_v105).Cardinality())
		var _264_v107 D2
		_ = _264_v107
		_264_v107 = Companion_D2_.Create_DC5_(_263_v106)
		var _265_v108 _dafny.Map
		_ = _265_v108
		_265_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D2 {
			if _182_v48 {
				return Companion_D2_.Create_DC5_(_263_v106)
			}
			return _264_v107
		})(), (_117_v0).Cmp(_dafny.IntOfInt64(953)) == 0)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_122_v5), 0))
		_ = _index58
		var _rhs49 bool = !(!((_dafny.IntOfInt64(956)).Cmp((_dafny.Zero).Minus(_117_v0)) < 0))
		_ = _rhs49
		var _rhs50 bool = true
		_ = _rhs50
		var _rhs51 _dafny.Int = _117_v0
		_ = _rhs51
		var _rhs52 bool = (func() bool {
			if (_265_v108).Contains(_264_v107) {
				return (_265_v108).Get(_264_v107).(bool)
			}
			return _182_v48
		})()
		_ = _rhs52
		var _lhs54 *GlobalState = _125_globalState
		_ = _lhs54
		var _lhs55 _dafny.Array = _122_v5
		_ = _lhs55
		var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_122_v5), 0))
		_ = _lhs56
		var _lhs57 *GlobalState = _125_globalState
		_ = _lhs57
		var _lhs58 *GlobalState = _125_globalState
		_ = _lhs58
		_lhs54.F2 = _rhs49
		(_lhs55).ArraySet1(_rhs50, (_lhs56).Int())
		_lhs57.F4 = _rhs51
		_lhs58.F2 = _rhs52
		_117_v0 = _117_v0
		var _266_v109 D2
		_ = _266_v109
		_266_v109 = Companion_D2_.Create_DC6_(_262_v105, (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool))
		var _267_v110 _dafny.CodePoint
		_ = _267_v110
		_267_v110 = _dafny.CodePoint('x')
		var _268_v111 _dafny.Map
		_ = _268_v111
		_268_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_266_v109, _267_v110)
		var _269_v112 _dafny.Set
		_ = _269_v112
		_269_v112 = _dafny.SetOf((func() _dafny.CodePoint {
			if (_268_v111).Contains(_266_v109) {
				return (_268_v111).Get(_266_v109).(_dafny.CodePoint)
			}
			return _267_v110
		})())
		var _270_v113 _dafny.Sequence
		_ = _270_v113
		_270_v113 = _dafny.SeqOf(Companion_Default___.Fm0(_117_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jfgpdshm")).Cardinality()), _125_globalState), _117_v0)
		var _271_v114 _dafny.Sequence
		_ = _271_v114
		_271_v114 = _dafny.SeqOf((_dafny.MultiSetFromSeq(_270_v113)).Cardinality())
		var _rhs53 _dafny.Set = Companion_Default___.Fm7(_267_v110, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(_117_v0, _125_globalState), _271_v114)).Cardinality()), _117_v0, _182_v48, _125_globalState)
		_ = _rhs53
		var _rhs54 bool = (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)
		_ = _rhs54
		var _rhs55 _dafny.CodePoint = _267_v110
		_ = _rhs55
		var _lhs59 *GlobalState = _125_globalState
		_ = _lhs59
		_269_v112 = _rhs53
		_lhs59.F2 = _rhs54
		_267_v110 = _rhs55
	} else if _source2.Is_DC3() {
		(_125_globalState).F2 = _182_v48
		var _272_v115 D0
		_ = _272_v115
		_272_v115 = Companion_D0_.Create_DC0_(_117_v0)
		_272_v115 = _272_v115
		(_125_globalState).F2 = _182_v48
		if _182_v48 {
			var _273_v116 D0
			_ = _273_v116
			_273_v116 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_119_v2).Cardinality()), !(_182_v48), _117_v0)
			var _274_v117 _dafny.Map
			_ = _274_v117
			_274_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_273_v116, _dafny.IntOfUint32((_183_v49).Cardinality()))
			var _275_v118 *C0
			_ = _275_v118
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__(_274_v117)
			_275_v118 = _nw44
			_119_v2 = _dafny.Companion_Sequence_.Concatenate(_119_v2, _dafny.Companion_Sequence_.Concatenate(_119_v2, _119_v2))
			var _276_v119 *C0
			_ = _276_v119
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__((_275_v118).F15())
			_276_v119 = _nw45
			_276_v119 = _275_v118
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_122_v5), 0))
			_ = _index59
			(_122_v5).ArraySet1((_117_v0).Cmp(_117_v0) < 0, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_122_v5), 0))
			_ = _index60
			(_122_v5).ArraySet1(Companion_Default___.Fm6(false, _dafny.IntOfInt64(603), _125_globalState), (_index60).Int())
		} else {
			_119_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_119_v2, _119_v2), _119_v2)
			var _277_v120 D0
			_ = _277_v120
			_277_v120 = Companion_D0_.Create_DC1_((_dafny.Zero).Minus((_dafny.Zero).Minus(_117_v0)), _182_v48, _117_v0)
			var _278_v121 _dafny.Map
			_ = _278_v121
			_278_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v120, _117_v0)
			var _279_v122 *C0
			_ = _279_v122
			var _nw46 *C0 = New_C0_()
			_ = _nw46
			_nw46.Ctor__(_278_v121)
			_279_v122 = _nw46
			var _280_v124 _dafny.Sequence
			_ = _280_v124
			_280_v124 = _dafny.SeqOf(_277_v120, _277_v120)
			var _281_v125 T0
			_ = _281_v125
			var _nw47 *C0 = New_C0_()
			_ = _nw47
			_nw47.Ctor__(func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate((_280_v124).Elements()); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _282_v123 D0
					_282_v123 = interface{}(_compr_11).(D0)
					if _dafny.Companion_Sequence_.Contains(_280_v124, _282_v123) {
						_coll11.Add(_282_v123, _117_v0)
					}
				}
				return _coll11.ToMap()
			}())
			_281_v125 = _nw47
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_121_v4), 0))
			_ = _index61
			(_121_v4).ArraySet1(_117_v0, (_index61).Int())
			var _283_v126 _dafny.Map
			_ = _283_v126
			_283_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _182_v48)
			var _284_v127 D2
			_ = _284_v127
			_284_v127 = Companion_D2_.Create_DC6_(_283_v126, (_182_v48) && (_182_v48))
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_121_v4), 0))
			_ = _index62
			var _rhs56 T0 = _281_v125
			_ = _rhs56
			var _rhs57 _dafny.Int = Companion_Default___.SafeModuloInt(_117_v0, _117_v0)
			_ = _rhs57
			var _rhs58 bool = _182_v48
			_ = _rhs58
			var _rhs59 _dafny.Int = _117_v0
			_ = _rhs59
			var _rhs60 D2 = _284_v127
			_ = _rhs60
			var _lhs60 _dafny.Array = _121_v4
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_121_v4), 0))
			_ = _lhs61
			var _lhs62 *GlobalState = _125_globalState
			_ = _lhs62
			var _lhs63 *GlobalState = _125_globalState
			_ = _lhs63
			_281_v125 = _rhs56
			(_lhs60).ArraySet1(_rhs57, (_lhs61).Int())
			_lhs62.F2 = _rhs58
			_lhs63.F3 = _rhs59
			_284_v127 = _rhs60
			var _285_v128 _dafny.CodePoint
			_ = _285_v128
			_285_v128 = _dafny.CodePoint('c')
			var _286_v129 D8
			_ = _286_v129
			_286_v129 = Companion_D8_.Create_DC20_(_285_v128)
			_285_v128 = (_286_v129).Dtor_cf22()
			var _287_v130 _dafny.Array
			_ = _287_v130
			var _nw48 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(6))
			_ = _nw48
			_287_v130 = _nw48
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_287_v130), 0))
			_ = _index63
			(_287_v130).ArraySet1(_281_v125, (_index63).Int())
			var _288_v131 _dafny.Map
			_ = _288_v131
			_288_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _281_v125)
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_287_v130), 0))
			_ = _index64
			(_287_v130).ArraySet1((func() T0 {
				if (_288_v131).Contains(_182_v48) {
					return (_288_v131).Get(_182_v48).(T0)
				}
				return _281_v125
			})(), (_index64).Int())
		}
	} else {
		var _289___mcc_h4 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
		_ = _289___mcc_h4
		var _290_cf0 _dafny.Int = _289___mcc_h4
		_ = _290_cf0
		var _291_v132 _dafny.CodePoint
		_ = _291_v132
		_291_v132 = _dafny.CodePoint('i')
		var _292_v133 _dafny.MultiSet
		_ = _292_v133
		_292_v133 = _dafny.MultiSetOf(_291_v132, _291_v132, _291_v132, _291_v132)
		(_125_globalState).F2 = Companion_Default___.Fm6(_182_v48, (_292_v133).Cardinality(), _125_globalState)
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))
		_ = _index65
		(_122_v5).ArraySet1((_183_v49).Select((Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_183_v49).Cardinality()))).Uint32()).(bool), (_index65).Int())
		var _293_v134 _dafny.Sequence
		_ = _293_v134
		_293_v134 = _dafny.SeqOf(_117_v0)
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))
		_ = _index66
		(_122_v5).ArraySet1((Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(694), _dafny.IntOfUint32((_dafny.SeqOf(_182_v48)).Cardinality()))).Cmp((_293_v134).Select((Companion_Default___.SafeIndex(_290_cf0, _dafny.IntOfUint32((_293_v134).Cardinality()))).Uint32()).(_dafny.Int)) <= 0, (_index66).Int())
		var _294_v135 _dafny.Set
		_ = _294_v135
		_294_v135 = _dafny.SetOf(_182_v48)
		var _295_v136 _dafny.Map
		_ = _295_v136
		_295_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_((_294_v135).Cardinality(), _182_v48, _117_v0), (_dafny.Zero).Minus(_290_cf0))
		var _296_v137 *C0
		_ = _296_v137
		var _nw49 *C0 = New_C0_()
		_ = _nw49
		_nw49.Ctor__(_295_v136)
		_296_v137 = _nw49
		var _297_v138 _dafny.Map
		_ = _297_v138
		_297_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, _296_v137)
		if (_118_v1).IsSubsetOf(Companion_Default___.Fm9((_297_v138).Cardinality(), Companion_Default___.Fm0(_290_cf0, _117_v0, _125_globalState), _125_globalState)) {
			(_125_globalState).F4 = (Companion_Default___.SafeModuloInt(_290_cf0, _290_cf0)).Minus(_dafny.IntOfInt64(-167))
			var _298_v139 _dafny.Map
			_ = _298_v139
			_298_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool))
			Companion_Default___.M0((func() bool {
				if (_298_v139).Contains(_291_v132) {
					return (_298_v139).Get(_291_v132).(bool)
				}
				return true
			})(), _125_globalState)
			_296_v137 = _296_v137
			var _299_v141 _dafny.MultiSet
			_ = _299_v141
			_299_v141 = _dafny.MultiSetOf(_290_cf0)
			var _300_v142 D0
			_ = _300_v142
			_300_v142 = Companion_D0_.Create_DC1_((func() _dafny.Int {
				if (_299_v141).Contains(_290_cf0) {
					return (_299_v141).Multiplicity(_290_cf0)
				}
				return _117_v0
			})(), (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool), _290_cf0)
			var _301_v143 _dafny.Map
			_ = _301_v143
			_301_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_v142, (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool))
			var _302_v144 *C0
			_ = _302_v144
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__((func() _dafny.Map {
				var _coll12 = _dafny.NewMapBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_301_v143).Keys().Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _303_v140 D0
					_303_v140 = interface{}(_compr_12).(D0)
					if (_301_v143).Contains(_303_v140) {
						_coll12.Add(_303_v140, _117_v0)
					}
				}
				return _coll12.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_v142, _117_v0)))
			_302_v144 = _nw50
			(_125_globalState).F2 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_304_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})), _291_v132)
		} else {
			_291_v132 = (Companion_D8_.Create_DC20_(_291_v132)).Dtor_cf22()
			Companion_Default___.M0((_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool), _125_globalState)
			(_125_globalState).F3 = _117_v0
			var _305_v145 _dafny.Map
			_ = _305_v145
			_305_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool), _dafny.IntOfInt64(189))
			var _306_v146 _dafny.Map
			_ = _306_v146
			_306_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, (func() _dafny.Int {
				if (_305_v145).Contains(true) {
					return (_305_v145).Get(true).(_dafny.Int)
				}
				return _290_cf0
			})())
			var _307_v147 _dafny.Map
			_ = _307_v147
			_307_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_306_v146).Contains(_290_cf0) {
					return (_306_v146).Get(_290_cf0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yeoger")).Cardinality())
			})(), _290_cf0)
			var _308_v148 _dafny.Sequence
			_ = _308_v148
			_308_v148 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(796), _dafny.IntOfInt64(480)))
			_307_v147 = (((_308_v148).Select((Companion_Default___.SafeIndex(_290_cf0, _dafny.IntOfUint32((_308_v148).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_119_v2).Cardinality()), _dafny.IntOfInt64(503)))).Merge(_307_v147)
			_119_v2 = _dafny.UnicodeSeqOfUtf8Bytes("vfpkjomp")
		}
		var _309_v149 _dafny.Map
		_ = _309_v149
		_309_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm10(_125_globalState), _296_v137)
		_290_cf0 = (_309_v149).Cardinality()
	}
	var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))
	_ = _index67
	(_121_v4).ArraySet1(_117_v0, (_index67).Int())
	var _310_v150 _dafny.Set
	_ = _310_v150
	_310_v150 = _dafny.SetOf(_122_v5, _122_v5, _122_v5, _122_v5)
	var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))
	_ = _index68
	(_121_v4).ArraySet1(Companion_Default___.SafeModuloInt((_310_v150).Cardinality(), _dafny.IntOfInt64(328)), (_index68).Int())
	var _311_v151 _dafny.Map
	_ = _311_v151
	_311_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v0, Companion_Default___.Fm6(_182_v48, (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), _125_globalState))
	var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
	_ = _index69
	(_122_v5).ArraySet1(_182_v48, (_index69).Int())
	var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))
	_ = _index70
	var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
	_ = _index71
	var _rhs61 _dafny.Int = ((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus(_117_v0))
	_ = _rhs61
	var _rhs62 bool = _182_v48
	_ = _rhs62
	var _rhs63 _dafny.Map = _311_v151
	_ = _rhs63
	var _rhs64 _dafny.Array = _121_v4
	_ = _rhs64
	var _rhs65 bool = (_dafny.IntOfInt64(956)).Cmp(_117_v0) != 0
	_ = _rhs65
	var _lhs64 _dafny.Array = _121_v4
	_ = _lhs64
	var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))
	_ = _lhs65
	var _lhs66 _dafny.Array = _122_v5
	_ = _lhs66
	var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
	_ = _lhs67
	(_lhs64).ArraySet1(_rhs61, (_lhs65).Int())
	_182_v48 = _rhs62
	_311_v151 = _rhs63
	_121_v4 = _rhs64
	(_lhs66).ArraySet1(_rhs65, (_lhs67).Int())
	var _312_v152 _dafny.MultiSet
	_ = _312_v152
	_312_v152 = _dafny.MultiSetOf(_118_v1)
	var _313_v153 _dafny.Sequence
	_ = _313_v153
	_313_v153 = _dafny.SeqOf(_312_v152)
	var _314_i12 _dafny.Int
	_ = _314_i12
	_314_i12 = _dafny.Zero
	{
		for (((_313_v153).Select((Companion_Default___.SafeIndex(_117_v0, _dafny.IntOfUint32((_313_v153).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_312_v152)).IsProperSubsetOf((_312_v152).Union(_312_v152)) {
			{
				if (_314_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_314_i12 = (_314_i12).Plus(_dafny.One)
				var _315_v154 D0
				_ = _315_v154
				_315_v154 = Companion_D0_.Create_DC1_((_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int), (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int))
				var _316_v155 _dafny.Map
				_ = _316_v155
				_316_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v154, _117_v0)
				var _317_v156 *C0
				_ = _317_v156
				var _nw51 *C0 = New_C0_()
				_ = _nw51
				_nw51.Ctor__(_316_v155)
				_317_v156 = _nw51
				var _318_v157 _dafny.Map
				_ = _318_v157
				_318_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_317_v156, _119_v2)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
				_ = _index72
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
				_ = _index73
				var _rhs66 bool = (_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)
				_ = _rhs66
				var _rhs67 _dafny.Sequence = (func() _dafny.Sequence {
					if (_318_v157).Contains(_317_v156) {
						return (_318_v157).Get(_317_v156).(_dafny.Sequence)
					}
					return _119_v2
				})()
				_ = _rhs67
				var _rhs68 bool = (true) && (!(_182_v48))
				_ = _rhs68
				var _lhs68 _dafny.Array = _122_v5
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
				_ = _lhs69
				var _lhs70 _dafny.Array = _122_v5
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))
				_ = _lhs71
				(_lhs68).ArraySet1(_rhs66, (_lhs69).Int())
				_119_v2 = _rhs67
				(_lhs70).ArraySet1(_rhs68, (_lhs71).Int())
				var _319_v158 _dafny.CodePoint
				_ = _319_v158
				_319_v158 = _dafny.CodePoint('d')
				var _320_v159 D8
				_ = _320_v159
				_320_v159 = Companion_D8_.Create_DC20_(_319_v158)
				_182_v48 = !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer15 func(_dafny.Int) D8) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_321_v159 D8) func(_dafny.Int) D8 {
					return func(_322_i13 _dafny.Int) D8 {
						return _321_v159
					}
				})(_320_v159))), _320_v159)
				_317_v156 = _317_v156
				var _323_v160 _dafny.Map
				_ = _323_v160
				_323_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)), _317_v156)
				_323_v160 = _323_v160
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _rhs69 bool = _182_v48
	_ = _rhs69
	var _rhs70 _dafny.Sequence = _dafny.SeqOf(_182_v48)
	_ = _rhs70
	var _lhs72 *GlobalState = _125_globalState
	_ = _lhs72
	_lhs72.F2 = _rhs69
	_183_v49 = _rhs70
	(_125_globalState).F2 = _182_v48
	var _324_v161 _dafny.Array
	_ = _324_v161
	var _len0_9 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_9
	var _nw52 _dafny.Array
	_ = _nw52
	if _len0_9.Cmp(_dafny.Zero) == 0 {
		_nw52 = _dafny.NewArray(_len0_9)
	} else {
		var _init9 func(_dafny.Int) _dafny.Map = (func(_325_v48 bool, _326_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_327_i15 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v48, _326_v0)
			}
		})(_182_v48, _117_v0)
		_ = _init9
		var _element0_9 = _init9(_dafny.Zero)
		_ = _element0_9
		_nw52 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
		(_nw52).ArraySet1(_element0_9, 0)
		var _nativeLen0_9 = (_len0_9).Int()
		_ = _nativeLen0_9
		for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
			(_nw52).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
		}
	}
	_324_v161 = _nw52
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_324_v161), 0))); ; {
		_guard_loop_0, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		var _328_i14 _dafny.Int
		_328_i14 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_328_i14).Sign() != -1) && ((_328_i14).Cmp(_dafny.ArrayLen((_324_v161), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_324_v161, (_328_i14).Int(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _117_v0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool), (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int)))))
		}
	}
	for _iter14 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok14 := _iter14()
		if !_ok14 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	Companion_Default___.M0(!(_182_v48) || ((_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)), _125_globalState)
	var _329_v162 _dafny.Sequence
	_ = _329_v162
	_329_v162 = (_119_v2)
	var _pat_let_tv2 = _117_v0
	_ = _pat_let_tv2
	var _pat_let_tv3 = _121_v4
	_ = _pat_let_tv3
	var _pat_let_tv4 = _121_v4
	_ = _pat_let_tv4
	var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))
	_ = _index74
	(_121_v4).ArraySet1(func(_source3 _dafny.Sequence) _dafny.Int {
		var _330___mcc_h5 _dafny.Sequence = _source3
		_ = _330___mcc_h5
		var _331_cf4 _dafny.Sequence = _330___mcc_h5
		_ = _331_cf4
		return (_dafny.Zero).Minus((_pat_let_tv2).Times((_dafny.MultiSetOf((_dafny.Zero).Minus((_pat_let_tv4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_pat_let_tv3), 0))).Int()).(_dafny.Int)))).Cardinality()))
	}(_329_v162), (_index74).Int())
	_117_v0 = _117_v0
	var _332_v163 _dafny.CodePoint
	_ = _332_v163
	_332_v163 = _dafny.CodePoint('e')
	var _333_v164 _dafny.Map
	_ = _333_v164
	_333_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_332_v163, _182_v48)
	var _334_v165 _dafny.Map
	_ = _334_v165
	_334_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v48, (_333_v164).Cardinality())
	var _335_v166 D0
	_ = _335_v166
	_335_v166 = Companion_D0_.Create_DC1_(_117_v0, _182_v48, (_dafny.Zero).Minus((_334_v165).Cardinality()))
	var _336_v167 _dafny.Map
	_ = _336_v167
	_336_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_335_v166, _117_v0)
	var _337_v168 *C0
	_ = _337_v168
	var _nw53 *C0 = New_C0_()
	_ = _nw53
	_nw53.Ctor__(_336_v167)
	_337_v168 = _nw53
	var _338_v169 _dafny.MultiSet
	_ = _338_v169
	_338_v169 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v168, _337_v168))
	var _339_v170 D8
	_ = _339_v170
	_339_v170 = Companion_D8_.Create_DC21_(_117_v0, _337_v168)
	var _340_v171 _dafny.Map
	_ = _340_v171
	_340_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_339_v170).Dtor_cf24(), _337_v168)
	_338_v169 = (_338_v169).Update(_340_v171, Companion_Default___.Abs(_117_v0))
	var _341_i16 _dafny.Int
	_ = _341_i16
	_341_i16 = _dafny.Zero
	{
		for ((_122_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_122_v5), 0))).Int()).(bool)) || (_182_v48) {
			{
				if (_341_i16).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_341_i16 = (_341_i16).Plus(_dafny.One)
				_334_v165 = (_334_v165).Update(_182_v48, _117_v0)
				_117_v0 = (_dafny.IntOfInt64(-967)).Plus(((_118_v1).Union(_118_v1)).Cardinality())
				_119_v2 = _119_v2
				_334_v165 = (_334_v165).Update(_182_v48, (_121_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_121_v4), 0))).Int()).(_dafny.Int))
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	var _342_v172 _dafny.Array
	_ = _342_v172
	var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
	_ = _nw54
	_342_v172 = _nw54
	var _343_v173 _dafny.Map
	_ = _343_v173
	_343_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v172, _182_v48)
	_343_v173 = (_343_v173).Update(_342_v172, _182_v48)
	_dafny.Print(_117_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v1).Equals(_dafny.SetOf(_dafny.IntOfInt64(-877))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_119_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("yxsvjmbqp"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v4).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_125_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_125_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_125_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_globalState).F9()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F12().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_126_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_v48)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_183_v49, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_239_i10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v150).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_311_v151).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_312_v152).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-877)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_313_v153, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-877))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_314_i12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_324_v161).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3)).UpdateUnsafe(true, _dafny.IntOfInt64(-2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_324_v161).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3)).UpdateUnsafe(true, _dafny.IntOfInt64(-2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_324_v161).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3)).UpdateUnsafe(true, _dafny.IntOfInt64(-2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_324_v161).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(3)).UpdateUnsafe(true, _dafny.IntOfInt64(-2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_329_v162).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_332_v163)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_333_v164).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_334_v165).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v166).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v166).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v166).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_336_v167).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(3), false, _dafny.IntOfInt64(-1)), _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_337_v168).F15()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(3), false, _dafny.IntOfInt64(-1)), _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_338_v169).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_339_v170).Dtor_cf23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_339_v170).Dtor_cf24()).F15()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.IntOfInt64(3), false, _dafny.IntOfInt64(-1)), _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_340_v171).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_341_i16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_343_v173).Cardinality())
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

type D0_DC2 struct {
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_() D0 {
	return D0{D0_DC2{}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_() D0 {
	return D0{D0_DC3{}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

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

func (_this D0) Dtor_cf0() _dafny.Int {
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
			return "D0.DC2"
		}
	case D0_DC3:
		{
			return "D0.DC3"
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
	case D0_DC2:
		{
			_, ok := other.Get_().(D0_DC2)
			return ok
		}
	case D0_DC3:
		{
			_, ok := other.Get_().(D0_DC3)
			return ok
		}
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

type D1_DC4 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + data.Cf4.VerbatimString(true) + ")"
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

type D2_DC6 struct {
	Cf6 _dafny.Map
	Cf7 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 _dafny.Map, Cf7 bool) D2 {
	return D2{D2_DC6{Cf6, Cf7}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf8 _dafny.Int
	Cf9 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf8 _dafny.Int, Cf9 _dafny.Array) D2 {
	return D2{D2_DC7{Cf8, Cf9}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf5 _dafny.MultiSet
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 _dafny.MultiSet) D2 {
	return D2{D2_DC5{Cf5}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.EmptyMap, false)
}

func (_this D2) Dtor_cf6() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC6).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf5() _dafny.MultiSet {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7 == data2.Cf7
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
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

type D3_DC8 struct {
	Cf10 _dafny.Map
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf10 _dafny.Map) D3 {
	return D3{D3_DC8{Cf10}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf10() _dafny.Map {
	return _this.Get_().(D3_DC8).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf10.Equals(data2.Cf10)
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
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_() D4 {
	return D4{D4_DC10{}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf12 _dafny.Int
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf12 _dafny.Int) D4 {
	return D4{D4_DC11{Cf12}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC9 struct {
	Cf11 _dafny.Sequence
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf11 _dafny.Sequence) D4 {
	return D4{D4_DC9{Cf11}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC12 struct {
	Cf13 D4
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf13 D4) D4 {
	return D4{D4_DC12{Cf13}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_()
}

func (_this D4) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D4_DC11).Cf12
}

func (_this D4) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D4_DC9).Cf11
}

func (_this D4) Dtor_cf13() D4 {
	return _this.Get_().(D4_DC12).Cf13
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf13) + ")"
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
			_, ok := other.Get_().(D4_DC10)
			return ok
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf14 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf14 _dafny.Map) D5 {
	return D5{D5_DC13{Cf14}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D5) Dtor_cf14() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf14
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D6_DC15 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.Int
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf16 _dafny.Int, Cf17 _dafny.Int) D6 {
	return D6{D6_DC15{Cf16, Cf17}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC14 struct {
	Cf15 _dafny.Array
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf15 _dafny.Array) D6 {
	return D6{D6_DC14{Cf15}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC16 struct {
	Cf18 D6
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf18 D6) D6 {
	return D6{D6_DC16{Cf18}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC15_(_dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf16
}

func (_this D6) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf17
}

func (_this D6) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D6_DC14).Cf15
}

func (_this D6) Dtor_cf18() D6 {
	return _this.Get_().(D6_DC16).Cf18
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf18.Equals(data2.Cf18)
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
	Cf20 _dafny.Int
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf20 _dafny.Int) D7 {
	return D7{D7_DC18{Cf20}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC17 struct {
	Cf19 _dafny.Array
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf19 _dafny.Array) D7 {
	return D7{D7_DC17{Cf19}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC19 struct {
	Cf21 D7
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf21 D7) D7 {
	return D7{D7_DC19{Cf21}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_(_dafny.Zero)
}

func (_this D7) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D7_DC18).Cf20
}

func (_this D7) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D7_DC17).Cf19
}

func (_this D7) Dtor_cf21() D7 {
	return _this.Get_().(D7_DC19).Cf21
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf19 == data2.Cf19
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 _dafny.Int
	Cf24 *C0
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf23 _dafny.Int, Cf24 *C0) D8 {
	return D8{D8_DC21{Cf23, Cf24}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC20 struct {
	Cf22 _dafny.CodePoint
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf22 _dafny.CodePoint) D8 {
	return D8{D8_DC20{Cf22}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(_dafny.Zero, (*C0)(nil))
}

func (_this D8) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf23
}

func (_this D8) Dtor_cf24() *C0 {
	return _this.Get_().(D8_DC21).Cf24
}

func (_this D8) Dtor_cf22() _dafny.CodePoint {
	return _this.Get_().(D8_DC20).Cf22
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf22 == data2.Cf22
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
	Cf26 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf26 _dafny.Int) D9 {
	return D9{D9_DC23{Cf26}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC24 struct {
	Cf27 _dafny.Int
	Cf28 *C0
	Cf29 _dafny.CodePoint
	Cf30 _dafny.Int
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf27 _dafny.Int, Cf28 *C0, Cf29 _dafny.CodePoint, Cf30 _dafny.Int) D9 {
	return D9{D9_DC24{Cf27, Cf28, Cf29, Cf30}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC25 struct {
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_() D9 {
	return D9{D9_DC25{}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC22 struct {
	Cf25 _dafny.Set
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf25 _dafny.Set) D9 {
	return D9{D9_DC22{Cf25}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(_dafny.Zero)
}

func (_this D9) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf26
}

func (_this D9) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf27
}

func (_this D9) Dtor_cf28() *C0 {
	return _this.Get_().(D9_DC24).Cf28
}

func (_this D9) Dtor_cf29() _dafny.CodePoint {
	return _this.Get_().(D9_DC24).Cf29
}

func (_this D9) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf30
}

func (_this D9) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D9_DC22).Cf25
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf26) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28 && data1.Cf29 == data2.Cf29 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D9_DC25:
		{
			_, ok := other.Get_().(D9_DC25)
			return ok
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf32 bool
	Cf33 bool
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf32 bool, Cf33 bool) D10 {
	return D10{D10_DC27{Cf32, Cf33}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

type D10_DC28 struct {
	Cf34 _dafny.Int
	Cf35 bool
	Cf36 _dafny.Int
	Cf37 T0
	Cf38 _dafny.Int
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf34 _dafny.Int, Cf35 bool, Cf36 _dafny.Int, Cf37 T0, Cf38 _dafny.Int) D10 {
	return D10{D10_DC28{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC26 struct {
	Cf31 _dafny.Map
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf31 _dafny.Map) D10 {
	return D10{D10_DC26{Cf31}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC27_(false, false)
}

func (_this D10) Dtor_cf32() bool {
	return _this.Get_().(D10_DC27).Cf32
}

func (_this D10) Dtor_cf33() bool {
	return _this.Get_().(D10_DC27).Cf33
}

func (_this D10) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D10_DC28).Cf34
}

func (_this D10) Dtor_cf35() bool {
	return _this.Get_().(D10_DC28).Cf35
}

func (_this D10) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D10_DC28).Cf36
}

func (_this D10) Dtor_cf37() T0 {
	return _this.Get_().(D10_DC28).Cf37
}

func (_this D10) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D10_DC28).Cf38
}

func (_this D10) Dtor_cf31() _dafny.Map {
	return _this.Get_().(D10_DC26).Cf31
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf31) + ")"
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
			return ok && data1.Cf32 == data2.Cf32 && data1.Cf33 == data2.Cf33
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35 && data1.Cf36.Cmp(data2.Cf36) == 0 && _dafny.AreEqual(data1.Cf37, data2.Cf37) && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf31.Equals(data2.Cf31)
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
	Cf40 bool
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf40 bool) D11 {
	return D11{D11_DC30{Cf40}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC29 struct {
	Cf39 _dafny.Set
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf39 _dafny.Set) D11 {
	return D11{D11_DC29{Cf39}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_(false)
}

func (_this D11) Dtor_cf40() bool {
	return _this.Get_().(D11_DC30).Cf40
}

func (_this D11) Dtor_cf39() _dafny.Set {
	return _this.Get_().(D11_DC29).Cf39
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && data1.Cf40 == data2.Cf40
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	F2   bool
	F3   _dafny.Int
	F4   _dafny.Int
	_f0  bool
	_f1  bool
	_f5  _dafny.Int
	_f6  _dafny.Array
	_f7  bool
	_f8  _dafny.Array
	_f9  _dafny.Array
	_f10 _dafny.Int
	_f11 bool
	_f12 _dafny.Sequence
	_f13 _dafny.Int
	_f14 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this.F3 = _dafny.Zero
	_this.F4 = _dafny.Zero
	_this._f0 = false
	_this._f1 = false
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f7 = false
	_this._f8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = _dafny.EmptySeq
	_this._f13 = _dafny.Zero
	_this._f14 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 bool, f3 _dafny.Int, f4 _dafny.Int, f5 _dafny.Int, f6 _dafny.Array, f7 bool, f8 _dafny.Array, f9 _dafny.Array, f10 _dafny.Int, f11 bool, f12 _dafny.Sequence, f13 _dafny.Int, f14 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
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
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Array {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() bool {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Array {
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
func (_this *GlobalState) F12() _dafny.Sequence {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f15 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f15 = _dafny.EmptyMap
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

func (_this *C0) Ctor__(f15 _dafny.Map) {
	{
		(_this)._f15 = f15
	}
}
func (_this *C0) F15() _dafny.Map {
	{
		return _this._f15
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
