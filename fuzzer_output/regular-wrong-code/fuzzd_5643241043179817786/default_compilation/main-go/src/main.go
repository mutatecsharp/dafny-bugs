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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf((Companion_D4_.Create_DC13_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(47))).Cardinality()), _dafny.IntOfInt64(-267), true)).Dtor_cf20(), false, false, !(true), false)).Cardinality(), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(39))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("hrdfpvvp")
	}))).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, (Companion_D0_.Create_DC0_(true)).Dtor_cf0())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) D0 {
	var _source0 D0 = Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_dafny.SeqOf(true, true)), _dafny.IntOfInt64(484)))
	_ = _source0
	if _source0.Is_DC1() {
		var _1___mcc_h0 _dafny.MultiSet = _source0.Get_().(D0_DC1).Cf1
		_ = _1___mcc_h0
		var _2___mcc_h1 _dafny.Int = _source0.Get_().(D0_DC1).Cf2
		_ = _2___mcc_h1
		var _3_cf2 _dafny.Int = _2___mcc_h1
		_ = _3_cf2
		var _4_cf1 _dafny.MultiSet = _1___mcc_h0
		_ = _4_cf1
		return Companion_D0_.Create_DC1_(_4_cf1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _3_cf2)).Cardinality())
	} else if _source0.Is_DC0() {
		var _5___mcc_h2 bool = _source0.Get_().(D0_DC0).Cf0
		_ = _5___mcc_h2
		var _6_cf0 bool = _5___mcc_h2
		_ = _6_cf0
		return Companion_D0_.Create_DC1_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(_6_cf0))), _dafny.IntOfInt64(170))
	} else {
		var _7___mcc_h3 D0 = _source0.Get_().(D0_DC2).Cf3
		_ = _7___mcc_h3
		var _8_cf3 D0 = _7___mcc_h3
		_ = _8_cf3
		return Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_dafny.SeqOf(!(false))), _dafny.IntOfInt64(261))
	}
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Map, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-262))).Cardinality())))).Union(func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(481), _dafny.IntOfInt64(21))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _9_v0 _dafny.Int
			_9_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(481)).Cmp(_9_v0) <= 0) && ((_9_v0).Cmp(_dafny.IntOfInt64(21)) < 0) {
				_coll0.Add((_9_v0).Minus(_dafny.IntOfInt64(-863)))
			}
		}
		return _coll0.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('n')), false)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.CodePoint('i')), true))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.CodePoint, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D3_.Create_DC9_(_dafny.SeqOf(false))).Dtor_cf10()
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), ((Companion_D11_.Create_DC32_(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(702))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_10_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("re")
	}))))).Dtor_cf55()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(517), _dafny.IntOfInt64(44))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(833), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(false))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge((Companion_D12_.Create_DC35_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))).Dtor_cf58())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(_dafny.SeqOf((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-10), _dafny.IntOfInt64(691))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _11_v0 _dafny.Int
			_11_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(-10)).Cmp(_11_v0) <= 0) && ((_11_v0).Cmp(_dafny.IntOfInt64(691)) < 0) {
				_coll1.Add((_11_v0).Plus(_dafny.IntOfInt64(155)), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('f'))).Cardinality()))
			}
		}
		return _coll1.ToMap()
	}()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hrux"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(346))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('l')
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D3_.Create_DC11_(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ihqlrtng")).Cardinality()), true, !(true))).Dtor_cf15())))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-488))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfInt64(-168))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-929))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-715))))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.CodePoint('u'))).Difference((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('b'), _dafny.CodePoint('d'))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _13_v0 _dafny.CodePoint
			_13_v0 = interface{}(_compr_2).(_dafny.CodePoint)
			if (_dafny.MultiSetOf(_dafny.CodePoint('b'), _dafny.CodePoint('d'))).Contains(_13_v0) {
				_coll2.Add(_13_v0)
			}
		}
		return _coll2.ToSet()
	}()).Union(_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('n'))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(738), (true) == (!(!(false))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-350), (_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('l'))).Cardinality())).Merge((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(260))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _14_v0 _dafny.Int
			_14_v0 = interface{}(_compr_3).(_dafny.Int)
			if (_dafny.SetOf(_dafny.IntOfInt64(260))).Contains(_14_v0) {
				_coll3.Add(Companion_Default___.SafeModuloInt(_14_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uvbomwi")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(786))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}(func(_15_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				}))).Cardinality()))
			}
		}
		return _coll3.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(970), _dafny.IntOfInt64(-978))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _16_v1 _dafny.Int
			_16_v1 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(970)).Cmp(_16_v1) <= 0) && ((_16_v1).Cmp(_dafny.IntOfInt64(-978)) < 0) {
				_coll4.Add((_16_v1).Times(_dafny.IntOfInt64(137)), _dafny.IntOfInt64(484))
			}
		}
		return _coll4.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(true))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.MultiSet, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var _hi0 _dafny.Int = p0
	_ = _hi0
	for _17_i0 := p0; _17_i0.Cmp(_hi0) < 0; _17_i0 = _17_i0.Plus(_dafny.One) {
		var _18_v0 _dafny.Array
		_ = _18_v0
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Sequence = (func(_19_p0 _dafny.Int, _20_i0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_21_i1 _dafny.Int) _dafny.Sequence {
					return (func() _dafny.Sequence {
						if false {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(899))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg4 _dafny.Int) interface{} {
									return coer4(arg4)
								}
							}((func(_22_p0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
								return func(_23_i2 _dafny.Int) _dafny.Sequence {
									return _dafny.SeqOf(_22_p0, (_dafny.SetOf(_22_p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wq")).Cardinality()))).Cardinality(), (_dafny.MultiSetOf(_22_p0)).Cardinality())
								}
							})(_19_p0)))
						}
						return _dafny.SeqOf(_dafny.SeqOf(_20_i0, _20_i0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg5 _dafny.Int) interface{} {
								return coer5(arg5)
							}
						}(func(_24_i3 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(143)
						})), _dafny.SeqOf(_19_p0))
					})()
				}
			})(p0, _17_i0)
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
		_18_v0 = _nw0
		var _25_v1 _dafny.Set
		_ = _25_v1
		_25_v1 = _dafny.SetOf((_dafny.Zero).Minus(p0))
		var _26_v2 _dafny.Sequence
		_ = _26_v2
		_26_v2 = _dafny.SeqOf(p0, (_25_v1).Cardinality(), p0)
		var _27_v3 _dafny.Sequence
		_ = _27_v3
		_27_v3 = _dafny.SeqOf(_26_v2)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_18_v0), 0))
		_ = _index0
		(_18_v0).ArraySet1(_27_v3, (_index0).Int())
		var _28_v4 bool
		_ = _28_v4
		_28_v4 = false
		var _29_v5 _dafny.Array
		_ = _29_v5
		var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(26))
		_ = _nw1
		_29_v5 = _nw1
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_29_v5), 0))
		_ = _index1
		(_29_v5).ArraySet1(_25_v1, (_index1).Int())
		var _30_v6 _dafny.CodePoint
		_ = _30_v6
		_30_v6 = _dafny.CodePoint('n')
		var _31_v7 _dafny.Map
		_ = _31_v7
		_31_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _30_v6)
		var _32_v8 *C0
		_ = _32_v8
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__(_17_i0, (_dafny.Zero).Minus(_17_i0), true)
		_32_v8 = _nw2
		var _33_v9 D2
		_ = _33_v9
		_33_v9 = Companion_D2_.Create_DC7_((_31_v7).Cardinality(), _32_v8)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_18_v0), 0))
		_ = _index2
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_29_v5), 0))
		_ = _index3
		var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}((func(_34_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_35_i4 _dafny.Int) _dafny.Sequence {
				return _34_v2
			}
		})(_26_v2))), _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-914))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_36_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_37_i5 _dafny.Int) _dafny.Int {
				return _36_p0
			}
		})(p0)))))
		_ = _rhs0
		var _rhs1 bool = _28_v4
		_ = _rhs1
		var _rhs2 _dafny.Int = (_33_v9).Dtor_cf7()
		_ = _rhs2
		var _rhs3 _dafny.Int = (_32_v8).F9()
		_ = _rhs3
		var _rhs4 _dafny.Set = _25_v1
		_ = _rhs4
		var _lhs0 _dafny.Array = _18_v0
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((_18_v0), 0))
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		var _lhs4 _dafny.Array = _29_v5
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_29_v5), 0))
		_ = _lhs5
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		_28_v4 = _rhs1
		_lhs2.F3 = _rhs2
		_lhs3.F6 = _rhs3
		(_lhs4).ArraySet1(_rhs4, (_lhs5).Int())
		var _38_v10 _dafny.Array
		_ = _38_v10
		var _nw3 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
		_ = _nw3
		_38_v10 = _nw3
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))
		_ = _index4
		(_38_v10).ArraySet1(_32_v8, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))
		_ = _index5
		(_38_v10).ArraySet1(_32_v8, (_index5).Int())
		_28_v4 = _28_v4
		if _28_v4 {
			var _39_v11 _dafny.Array
			_ = _39_v11
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_1
			var _nw4 _dafny.Array
			_ = _nw4
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw4 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_40_p1 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
					return func(_41_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_41_i6, (_40_p1).Cardinality())
					}
				})(p1)
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
			_39_v11 = _nw4
			var _42_v12 _dafny.Sequence
			_ = _42_v12
			_42_v12 = _dafny.UnicodeSeqOfUtf8Bytes("kn")
			var _43_v13 _dafny.Set
			_ = _43_v13
			_43_v13 = _dafny.SetOf(_32_v8)
			var _44_v14 _dafny.MultiSet
			_ = _44_v14
			_44_v14 = _dafny.MultiSetOf(_43_v13)
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_39_v11), 0))
			_ = _index6
			(_39_v11).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_42_v12).Cardinality()), (func() _dafny.Int {
				if (_44_v14).Contains(_dafny.SetOf((_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0), _32_v8)) {
					return (_44_v14).Multiplicity(_dafny.SetOf((_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0), _32_v8))
				}
				return (_26_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_42_v12).Cardinality()), _dafny.IntOfUint32((_26_v2).Cardinality()))).Uint32()).(_dafny.Int)
			})()), (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_39_v11), 0))
			_ = _index7
			(_39_v11).ArraySet1(_dafny.IntOfInt64(551), (_index7).Int())
			var _45_v15 _dafny.Map
			_ = _45_v15
			_45_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_32_v8).F9(), (_32_v8).F9())
			(globalState).F6 = (_dafny.IntOfInt64(42)).Times(Companion_Default___.SafeDivisionInt((_32_v8).F9(), (_45_v15).Cardinality()))
			var _46_v16 _dafny.Sequence
			_ = _46_v16
			_46_v16 = _dafny.SeqOf(_28_v4, _28_v4, !(p1).Contains(!(!(_28_v4))))
			var _47_v17 D4
			_ = _47_v17
			_47_v17 = Companion_D4_.Create_DC14_((_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0), _28_v4)
			var _48_v18 _dafny.Map
			_ = _48_v18
			_48_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0), _30_v6)
			var _49_v19 _dafny.Map
			_ = _49_v19
			_49_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_47_v17, (func() _dafny.CodePoint {
				if (_48_v18).Contains((_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0)) {
					return (_48_v18).Get((_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0)).(_dafny.CodePoint)
				}
				return _30_v6
			})())
			var _50_v20 _dafny.MultiSet
			_ = _50_v20
			_50_v20 = _dafny.MultiSetOf((_49_v19).Cardinality())
			var _51_v21 T0
			_ = _51_v21
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__(_17_i0, _dafny.IntOfUint32((_46_v16).Cardinality()), _28_v4)
			_51_v21 = _nw5
			var _52_v22 _dafny.Sequence
			_ = _52_v22
			_52_v22 = _dafny.SeqOf(_51_v21, _51_v21, _51_v21)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_39_v11), 0))
			_ = _index8
			var _rhs5 _dafny.Sequence = _46_v16
			_ = _rhs5
			var _rhs6 _dafny.CodePoint = _30_v6
			_ = _rhs6
			var _rhs7 _dafny.Int = (_dafny.Zero).Minus((_50_v20).Cardinality())
			_ = _rhs7
			var _rhs8 _dafny.Int = (((_dafny.Zero).Minus(_dafny.IntOfUint32((_52_v22).Cardinality()))).Times((_39_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_39_v11), 0))).Int()).(_dafny.Int))).Plus((_32_v8).F8())
			_ = _rhs8
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			var _lhs7 _dafny.Array = _39_v11
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_39_v11), 0))
			_ = _lhs8
			_46_v16 = _rhs5
			_30_v6 = _rhs6
			_lhs6.F6 = _rhs7
			(_lhs7).ArraySet1(_rhs8, (_lhs8).Int())
			var _53_v23 D6
			_ = _53_v23
			_53_v23 = Companion_D6_.Create_DC17_(_51_v21.F7(), false, _28_v4)
			var _54_v24 _dafny.Sequence
			_ = _54_v24
			_54_v24 = _dafny.SeqOf(_53_v23, _53_v23, _53_v23)
			_28_v4 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_54_v24, _dafny.SeqOf(_53_v23)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(576))).Uint32(), func(coer8 func(_dafny.Int) D6) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_55_v23 D6) func(_dafny.Int) D6 {
				return func(_56_i7 _dafny.Int) D6 {
					return _55_v23
				}
			})(_53_v23))), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_54_v24, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ljh")).Cardinality()), _dafny.IntOfUint32((_54_v24).Cardinality()))).Uint32(), _53_v23), (Companion_Default___.SafeIndex((_32_v8).F8(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_54_v24, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ljh")).Cardinality()), _dafny.IntOfUint32((_54_v24).Cardinality()))).Uint32(), _53_v23)).Cardinality()))).Uint32(), Companion_D6_.Create_DC17_(_51_v21.F7(), _51_v21.F7(), _51_v21.F7()))))
			var _57_v25 _dafny.Sequence
			_ = _57_v25
			_57_v25 = _dafny.SeqOf(_46_v16)
			var _58_v26 _dafny.MultiSet
			_ = _58_v26
			_58_v26 = _dafny.MultiSetOf(_46_v16, _46_v16, p2, p2)
			var _59_v27 _dafny.Sequence
			_ = _59_v27
			_59_v27 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_57_v25), _58_v26)
			var _60_v28 D0
			_ = _60_v28
			_60_v28 = Companion_D0_.Create_DC1_((_59_v27).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_42_v12).Cardinality()), _dafny.IntOfUint32((_59_v27).Cardinality()))).Uint32()).(_dafny.MultiSet), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(233), p0))
			var _rhs9 _dafny.Sequence = _46_v16
			_ = _rhs9
			var _rhs10 *C0 = (func() *C0 {
				if Companion_Default___.Fm3((_39_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_39_v11), 0))).Int()).(_dafny.Int), globalState) {
					return _32_v8
				}
				return (_38_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_38_v10), 0))).Int()).(*C0)
			})()
			_ = _rhs10
			var _rhs11 D0 = _60_v28
			_ = _rhs11
			_46_v16 = _rhs9
			_32_v8 = _rhs10
			_60_v28 = _rhs11
		} else {
			_25_v1 = (_29_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_29_v5), 0))).Int()).(_dafny.Set)
			var _61_v29 _dafny.Sequence
			_ = _61_v29
			_61_v29 = _dafny.UnicodeSeqOfUtf8Bytes("dled")
			var _62_v30 _dafny.Sequence
			_ = _62_v30
			_62_v30 = _dafny.SeqOf(_61_v29, _dafny.UnicodeSeqOfUtf8Bytes("vsxctdar"), _61_v29, _dafny.UnicodeSeqOfUtf8Bytes("gmyfssb"))
			_62_v30 = _62_v30
			var _63_v31 *C0
			_ = _63_v31
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__((_32_v8).F8(), (_32_v8).F9(), (p0).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(641))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_64_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_65_i8 _dafny.Int) _dafny.CodePoint {
					return _64_v6
				}
			})(_30_v6)))).Cardinality())) > 0)
			_63_v31 = _nw6
			var _66_v32 *C0
			_ = _66_v32
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_67_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_68_i9 _dafny.Int) _dafny.CodePoint {
					return _67_v6
				}
			})(_30_v6))), (Companion_Default___.SafeIndex((_32_v8).F8(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_69_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_70_i9 _dafny.Int) _dafny.CodePoint {
					return _69_v6
				}
			})(_30_v6)))).Cardinality()))).Uint32(), _30_v6)).Cardinality()), _dafny.IntOfInt64(403), _28_v4)
			_66_v32 = _nw7
			var _71_v33 _dafny.Map
			_ = _71_v33
			_71_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _72_v34 _dafny.Map
			_ = _72_v34
			_72_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_71_v33).Contains((_dafny.Zero).Minus(_17_i0)) {
					return (_71_v33).Get((_dafny.Zero).Minus(_17_i0)).(bool)
				}
				return _28_v4
			})(), (_26_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_26_v2).Cardinality()))).Uint32()).(_dafny.Int))
			var _73_v35 D3
			_ = _73_v35
			_73_v35 = Companion_D3_.Create_DC10_(_30_v6, _28_v4)
			_72_v34 = (_72_v34).Update((_73_v35).Dtor_cf12(), Companion_Default___.SafeModuloInt((_32_v8).F9(), _dafny.IntOfInt64(-391)))
		}
	}
	var _74_v36 bool
	_ = _74_v36
	_74_v36 = true
	var _75_v37 _dafny.Sequence
	_ = _75_v37
	_75_v37 = _dafny.UnicodeSeqOfUtf8Bytes("deqbibs")
	var _hi1 _dafny.Int = p0
	_ = _hi1
	for _76_i10 := (func() _dafny.Int {
		if !(_74_v36) {
			return _dafny.IntOfUint32((_75_v37).Cardinality())
		}
		return _dafny.IntOfUint32((_75_v37).Cardinality())
	})(); _76_i10.Cmp(_hi1) < 0; _76_i10 = _76_i10.Plus(_dafny.One) {
		var _77_v38 _dafny.Array
		_ = _77_v38
		var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(8))
		_ = _nw8
		_77_v38 = _nw8
		var _78_v39 _dafny.Sequence
		_ = _78_v39
		_78_v39 = _dafny.SeqOf(p0)
		var _79_v40 _dafny.CodePoint
		_ = _79_v40
		_79_v40 = _dafny.CodePoint('x')
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_77_v38), 0))
		_ = _index9
		(_77_v38).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.MultiSetOf(Companion_Default___.Fm0(_dafny.IntOfInt64(340), _74_v36, _78_v39, _79_v40, globalState), p0, _76_i10))), (_index9).Int())
		var _80_v41 _dafny.Array
		_ = _80_v41
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_2
		var _nw9 _dafny.Array
		_ = _nw9
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw9 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_81_v40 _dafny.CodePoint) func(_dafny.Int) bool {
				return func(_82_i11 _dafny.Int) bool {
					return _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("uyntxwc"), _81_v40)
				}
			})(_79_v40)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw9 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw9).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw9).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_80_v41 = _nw9
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_80_v41), 0))
		_ = _index10
		(_80_v41).ArraySet1((p0).Cmp(p0) == 0, (_index10).Int())
		var _83_v42 _dafny.Array
		_ = _83_v42
		var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(10))
		_ = _nw10
		_83_v42 = _nw10
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_83_v42), 0))
		_ = _index11
		(_83_v42).ArraySet1(p1, (_index11).Int())
		var _84_v43 D6
		_ = _84_v43
		_84_v43 = Companion_D6_.Create_DC17_(_74_v36, _74_v36, _74_v36)
		var _85_v44 _dafny.Map
		_ = _85_v44
		_85_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v43, p0)
		var _86_v45 _dafny.MultiSet
		_ = _86_v45
		_86_v45 = _dafny.MultiSetOf(_85_v44, _85_v44, _85_v44, _85_v44)
		var _87_v46 *C0
		_ = _87_v46
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__(_dafny.IntOfInt64(666), _76_i10, false)
		_87_v46 = _nw11
		var _88_v47 D2
		_ = _88_v47
		_88_v47 = Companion_D2_.Create_DC7_(p0, _87_v46)
		var _89_v48 _dafny.MultiSet
		_ = _89_v48
		_89_v48 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0), (_88_v47).Dtor_cf7(), p0, p0)
		var _90_v49 D1
		_ = _90_v49
		_90_v49 = Companion_D1_.Create_DC3_(_78_v39)
		var _91_v50 D7
		_ = _91_v50
		_91_v50 = Companion_D7_.Create_DC19_(_dafny.MultiSetFromSeq(_78_v39))
		var _92_v51 _dafny.MultiSet
		_ = _92_v51
		_92_v51 = _dafny.MultiSetOf(_89_v48, _dafny.MultiSetFromSeq((_90_v49).Dtor_cf4()), (_91_v50).Dtor_cf29(), _89_v48)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_77_v38), 0))
		_ = _index12
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_80_v41), 0))
		_ = _index13
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_83_v42), 0))
		_ = _index14
		var _rhs12 bool = !(!(false)) || ((_86_v45).IsDisjointFrom(_86_v45))
		_ = _rhs12
		var _rhs13 _dafny.MultiSet = _92_v51
		_ = _rhs13
		var _rhs14 _dafny.Array = _80_v41
		_ = _rhs14
		var _rhs15 bool = !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm7(globalState), _78_v39)
		_ = _rhs15
		var _rhs16 _dafny.MultiSet = _dafny.MultiSetFromSeq(p2)
		_ = _rhs16
		var _lhs9 _dafny.Array = _77_v38
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.ArrayLen((_77_v38), 0))
		_ = _lhs10
		var _lhs11 _dafny.Array = _80_v41
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_80_v41), 0))
		_ = _lhs12
		var _lhs13 _dafny.Array = _83_v42
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_83_v42), 0))
		_ = _lhs14
		_74_v36 = _rhs12
		(_lhs9).ArraySet1(_rhs13, (_lhs10).Int())
		_80_v41 = _rhs14
		(_lhs11).ArraySet1(_rhs15, (_lhs12).Int())
		(_lhs13).ArraySet1(_rhs16, (_lhs14).Int())
		var _93_v52 T0
		_ = _93_v52
		var _nw12 *C0 = New_C0_()
		_ = _nw12
		_nw12.Ctor__((_87_v46).F9(), p0, !((_80_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_80_v41), 0))).Int()).(bool)))
		_93_v52 = _nw12
		var _94_v53 _dafny.Sequence
		_ = _94_v53
		_94_v53 = _dafny.SeqOf(_93_v52)
		var _95_v54 _dafny.Array
		_ = _95_v54
		var _nwElement0_0 _dafny.Int = p0
		_ = _nwElement0_0
		var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(4))
		_ = _nw13
		(_nw13).ArraySet1(_nwElement0_0, 0)
		(_nw13).ArraySet1(Companion_Default___.SafeDivisionInt((_87_v46).F8(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xpqepkn"), (_94_v53).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.IntOfUint32((_94_v53).Cardinality()))).Uint32()).(T0))).Cardinality()), 1)
		(_nw13).ArraySet1((_87_v46).F9(), 2)
		(_nw13).ArraySet1(Companion_Default___.SafeModuloInt(_76_i10, (_87_v46).F9()), 3)
		_95_v54 = _nw13
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_95_v54), 0))
		_ = _index15
		(_95_v54).ArraySet1(p0, (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_95_v54), 0))
		_ = _index16
		(_95_v54).ArraySet1((_87_v46).F8(), (_index16).Int())
		var _96_v55 _dafny.Sequence
		_ = _96_v55
		_96_v55 = _dafny.SeqOf(_87_v46)
		var _97_v56 _dafny.Set
		_ = _97_v56
		_97_v56 = _dafny.SetOf(_76_i10, _dafny.IntOfUint32((_96_v55).Cardinality()))
		var _98_v57 _dafny.Set
		_ = _98_v57
		_98_v57 = _dafny.SetOf((_97_v56).Union(_dafny.SetOf(_dafny.IntOfUint32((p2).Cardinality()))))
		_98_v57 = _dafny.SetOf(_97_v56)
		(globalState).F3 = _dafny.IntOfInt64(-293)
	}
	var _99_v58 _dafny.Map
	_ = _99_v58
	_99_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _75_v37)
	_75_v37 = (func() _dafny.Sequence {
		if (_99_v58).Contains((_dafny.Zero).Minus((func() _dafny.Int {
			if (p1).Contains(false) {
				return (p1).Multiplicity(false)
			}
			return p0
		})())) {
			return (_99_v58).Get((_dafny.Zero).Minus((func() _dafny.Int {
				if (p1).Contains(false) {
					return (p1).Multiplicity(false)
				}
				return p0
			})())).(_dafny.Sequence)
		}
		return _75_v37
	})()
	if (p0).Cmp(_dafny.IntOfInt64(784)) != 0 {
		var _100_v59 _dafny.Set
		_ = _100_v59
		_100_v59 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lmjjoibd"), _75_v37), _75_v37, _75_v37, _75_v37)
		_100_v59 = ((_100_v59).Intersection(_100_v59)).Difference(_100_v59)
		if (_dafny.IntOfUint32((_75_v37).Cardinality())).Cmp(_dafny.IntOfInt64(106)) >= 0 {
			var _101_v60 _dafny.Sequence
			_ = _101_v60
			_101_v60 = _dafny.SeqOf(p0)
			var _102_v61 _dafny.CodePoint
			_ = _102_v61
			_102_v61 = _dafny.CodePoint('d')
			var _103_v62 T0
			_ = _103_v62
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(Companion_Default___.Fm0(p0, _74_v36, _101_v60, _102_v61, globalState), _dafny.IntOfUint32((_75_v37).Cardinality()), _74_v36)
			_103_v62 = _nw14
			var _104_v63 _dafny.MultiSet
			_ = _104_v63
			_104_v63 = _dafny.MultiSetOf(_103_v62, _103_v62)
			_74_v36 = ((func() _dafny.MultiSet {
				if _74_v36 {
					return _104_v63
				}
				return _104_v63
			})()).IsProperSubsetOf(_104_v63)
			var _105_v64 *C0
			_ = _105_v64
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(_dafny.IntOfInt64(520), p0, _74_v36)
			_105_v64 = _nw15
			var _106_v65 _dafny.Sequence
			_ = _106_v65
			_106_v65 = _dafny.SeqOf(_105_v64)
			var _107_v66 _dafny.Array
			_ = _107_v66
			var _nwElement0_1 _dafny.Int = Companion_Default___.Fm0(p0, _74_v36, _dafny.SeqOf(_dafny.IntOfUint32((p2).Cardinality())), _102_v61, globalState)
			_ = _nwElement0_1
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(16))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_1, 0)
			(_nw16).ArraySet1(_dafny.IntOfInt64(483), 1)
			(_nw16).ArraySet1(_dafny.IntOfUint32((_106_v65).Cardinality()), 2)
			(_nw16).ArraySet1(p0, 3)
			(_nw16).ArraySet1((_105_v64).F9(), 4)
			(_nw16).ArraySet1((_105_v64).F8(), 5)
			(_nw16).ArraySet1((_105_v64).F9(), 6)
			(_nw16).ArraySet1((_105_v64).F8(), 7)
			(_nw16).ArraySet1(_dafny.IntOfInt64(-166), 8)
			(_nw16).ArraySet1((_105_v64).F9(), 9)
			(_nw16).ArraySet1((_105_v64).F9(), 10)
			(_nw16).ArraySet1((_105_v64).F8(), 11)
			(_nw16).ArraySet1((_105_v64).F9(), 12)
			(_nw16).ArraySet1((_105_v64).F8(), 13)
			(_nw16).ArraySet1((_dafny.Zero).Minus((_101_v60).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_101_v60).Cardinality()))).Uint32()).(_dafny.Int)), 14)
			(_nw16).ArraySet1((_dafny.Zero).Minus((_105_v64).F8()), 15)
			_107_v66 = _nw16
			var _108_v67 D7
			_ = _108_v67
			_108_v67 = Companion_D7_.Create_DC20_(_74_v36, _103_v62.F7(), _105_v64)
			var _109_v68 _dafny.Map
			_ = _109_v68
			_109_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v66, (func() D7 {
				if _74_v36 {
					return _108_v67
				}
				return _108_v67
			})())
			var _110_v69 _dafny.Map
			_ = _110_v69
			_110_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v36, _107_v66)
			var _111_v70 _dafny.Sequence
			_ = _111_v70
			_111_v70 = _dafny.SeqOf(_109_v68)
			_109_v68 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_110_v69).Contains(_74_v36) {
					return (_110_v69).Get(_74_v36).(_dafny.Array)
				}
				return _107_v66
			})(), _108_v67)).Merge((_111_v70).Select((Companion_Default___.SafeIndex((_105_v64).F8(), _dafny.IntOfUint32((_111_v70).Cardinality()))).Uint32()).(_dafny.Map))
			var _112_v71 _dafny.Map
			_ = _112_v71
			_112_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v61, _103_v62)
			var _113_v72 _dafny.Map
			_ = _113_v72
			_113_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v62.F7(), (_105_v64).F9())
			var _114_v73 _dafny.MultiSet
			_ = _114_v73
			_114_v73 = _dafny.MultiSetOf(_113_v72, _113_v72)
			var _115_v74 _dafny.Array
			_ = _115_v74
			var _nwElement0_2 _dafny.Map = (_112_v71).Update(Companion_Default___.Fm11(_114_v73, p0, _75_v37, _dafny.UnicodeSeqOfUtf8Bytes("aicyyk"), globalState), _103_v62)
			_ = _nwElement0_2
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(15))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_2, 0)
			(_nw17).ArraySet1(_112_v71, 1)
			(_nw17).ArraySet1(_112_v71, 2)
			(_nw17).ArraySet1(_112_v71, 3)
			(_nw17).ArraySet1(_112_v71, 4)
			(_nw17).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_102_v61, _103_v62), 5)
			(_nw17).ArraySet1(_112_v71, 6)
			(_nw17).ArraySet1(_112_v71, 7)
			(_nw17).ArraySet1(_112_v71, 8)
			(_nw17).ArraySet1(_112_v71, 9)
			(_nw17).ArraySet1(_112_v71, 10)
			(_nw17).ArraySet1(_112_v71, 11)
			(_nw17).ArraySet1(_112_v71, 12)
			(_nw17).ArraySet1(_112_v71, 13)
			(_nw17).ArraySet1((_112_v71).Update(_102_v61, _103_v62), 14)
			_115_v74 = _nw17
			var _116_v75 _dafny.Set
			_ = _116_v75
			_116_v75 = _dafny.SetOf(_101_v60, _101_v60)
			var _117_v76 _dafny.Map
			_ = _117_v76
			_117_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v74, _116_v75)
			var _118_v77 D8
			_ = _118_v77
			_118_v77 = Companion_D8_.Create_DC22_(_116_v75)
			_117_v76 = (_117_v76).Update(_115_v74, ((_118_v77).Dtor_cf34()).Intersection(_116_v75))
			var _119_v78 _dafny.Array
			_ = _119_v78
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_3
			var _nw18 _dafny.Array
			_ = _nw18
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw18 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_120_v62 T0) func(_dafny.Int) bool {
					return func(_121_i12 _dafny.Int) bool {
						return _120_v62.F7()
					}
				})(_103_v62)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw18 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw18).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw18).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_119_v78 = _nw18
			var _122_v79 _dafny.Map
			_ = _122_v79
			_122_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v37, _119_v78)
			_122_v79 = ((_122_v79).Merge(_122_v79)).Merge(_122_v79)
			var _123_v80 _dafny.Sequence
			_ = _123_v80
			_123_v80 = _dafny.SeqOf(_119_v78, _119_v78)
			_123_v80 = _123_v80
		} else {
			var _124_v81 _dafny.Map
			_ = _124_v81
			_124_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _74_v36)
			var _125_v82 _dafny.Sequence
			_ = _125_v82
			_125_v82 = _dafny.SeqOf(_124_v81, Companion_Default___.Fm15(true, _74_v36, p0, globalState))
			_74_v36 = !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_124_v81), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_124_v81, _124_v81, Companion_Default___.Fm15(_74_v36, Companion_Default___.Fm3(p0, globalState), _dafny.IntOfInt64(137), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _74_v36)), _125_v82))
			var _126_v83 _dafny.Array
			_ = _126_v83
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
			_ = _nw19
			_126_v83 = _nw19
			var _127_v84 _dafny.CodePoint
			_ = _127_v84
			_127_v84 = _dafny.CodePoint('f')
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_126_v83), 0))
			_ = _index17
			(_126_v83).ArraySet1CodePoint(_127_v84, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_126_v83), 0))
			_ = _index18
			(_126_v83).ArraySet1CodePoint(_127_v84, (_index18).Int())
			var _128_v85 _dafny.Map
			_ = _128_v85
			_128_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v36, false)
			_128_v85 = (_128_v85).Update(true, (_74_v36) || (_74_v36))
			var _129_v86 _dafny.Set
			_ = _129_v86
			_129_v86 = _dafny.SetOf(p0, _dafny.IntOfInt64(-138))
			var _130_v88 _dafny.MultiSet
			_ = _130_v88
			_130_v88 = _dafny.MultiSetOf((_129_v86).Intersection(func() _dafny.Set {
				var _coll5 = _dafny.NewBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_129_v86).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _131_v87 _dafny.Int
					_131_v87 = interface{}(_compr_5).(_dafny.Int)
					if (_129_v86).Contains(_131_v87) {
						_coll5.Add(Companion_Default___.SafeDivisionInt(_131_v87, _dafny.IntOfInt64(948)))
					}
				}
				return _coll5.ToSet()
			}()))
			var _132_v90 _dafny.Set
			_ = _132_v90
			_132_v90 = _dafny.SetOf(_74_v36)
			var _133_v91 _dafny.MultiSet
			_ = _133_v91
			_133_v91 = _dafny.MultiSetOf(_132_v90)
			var _134_v92 D9
			_ = _134_v92
			_134_v92 = Companion_D9_.Create_DC27_(_133_v91)
			var _135_v93 _dafny.Sequence
			_ = _135_v93
			_135_v93 = _dafny.SeqOf(_133_v91, _133_v91, (_134_v92).Dtor_cf43(), _133_v91, _133_v91)
			var _136_v94 _dafny.MultiSet
			_ = _136_v94
			_136_v94 = _dafny.MultiSetOf(p0)
			var _137_v95 _dafny.Map
			_ = _137_v95
			_137_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_136_v94).Contains(_dafny.IntOfInt64(123)) {
					return (_136_v94).Multiplicity(_dafny.IntOfInt64(123))
				}
				return p0
			})(), p0)
			var _rhs17 bool = _74_v36
			_ = _rhs17
			var _rhs18 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("uownf")
			_ = _rhs18
			var _rhs19 _dafny.Int = (func() _dafny.Int {
				if (_130_v88).Contains((_129_v86).Union(func() _dafny.Set {
					var _coll6 = _dafny.NewBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(113), _dafny.IntOfInt64(926))); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _138_v89 _dafny.Int
						_138_v89 = interface{}(_compr_6).(_dafny.Int)
						if ((_dafny.IntOfInt64(113)).Cmp(_138_v89) <= 0) && ((_138_v89).Cmp(_dafny.IntOfInt64(926)) < 0) {
							_coll6.Add((_138_v89).Minus(p0))
						}
					}
					return _coll6.ToSet()
				}())) {
					return (_130_v88).Multiplicity((_129_v86).Union(func() _dafny.Set {
						var _coll7 = _dafny.NewBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(113), _dafny.IntOfInt64(926))); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _139_v89 _dafny.Int
							_139_v89 = interface{}(_compr_7).(_dafny.Int)
							if ((_dafny.IntOfInt64(113)).Cmp(_139_v89) <= 0) && ((_139_v89).Cmp(_dafny.IntOfInt64(926)) < 0) {
								_coll7.Add((_139_v89).Minus(p0))
							}
						}
						return _coll7.ToSet()
					}()))
				}
				return ((_135_v93).Select((Companion_Default___.SafeIndex((_137_v95).Cardinality(), _dafny.IntOfUint32((_135_v93).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()
			})()
			_ = _rhs19
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			_74_v36 = _rhs17
			r0 = _rhs18
			_lhs15.F6 = _rhs19
			var _140_v96 *C0
			_ = _140_v96
			var _nw20 *C0 = New_C0_()
			_ = _nw20
			_nw20.Ctor__((Companion_Default___.Fm16(_74_v36, _dafny.IntOfUint32((_75_v37).Cardinality()), _74_v36, _74_v36, globalState)).Cardinality(), p0, _74_v36)
			_140_v96 = _nw20
			var _141_v97 D6
			_ = _141_v97
			_141_v97 = Companion_D6_.Create_DC18_(_140_v96)
			var _142_v98 _dafny.Array
			_ = _142_v98
			var _nwElement0_3 D6 = _141_v97
			_ = _nwElement0_3
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(6))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_3, 0)
			(_nw21).ArraySet1(_141_v97, 1)
			(_nw21).ArraySet1(Companion_D6_.Create_DC18_(_140_v96), 2)
			(_nw21).ArraySet1(_141_v97, 3)
			(_nw21).ArraySet1(_141_v97, 4)
			(_nw21).ArraySet1(_141_v97, 5)
			_142_v98 = _nw21
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_142_v98), 0))
			_ = _index19
			(_142_v98).ArraySet1(Companion_D6_.Create_DC18_(_140_v96), (_index19).Int())
			var _pat_let_tv0 = _140_v96
			_ = _pat_let_tv0
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_142_v98), 0))
			_ = _index20
			(_142_v98).ArraySet1(func(_pat_let0_0 D6) D6 {
				return func(_143_dt__update__tmp_h0 D6) D6 {
					return func(_pat_let1_0 *C0) D6 {
						return func(_144_dt__update_hcf28_h0 *C0) D6 {
							return Companion_D6_.Create_DC18_(_144_dt__update_hcf28_h0)
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_141_v97), (_index20).Int())
		}
		var _145_v99 T0
		_ = _145_v99
		var _nw22 *C0 = New_C0_()
		_ = _nw22
		_nw22.Ctor__((_dafny.Zero).Minus(p0), p0, _74_v36)
		_145_v99 = _nw22
		var _146_v100 *C0
		_ = _146_v100
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_145_v99, _145_v99)).Cardinality()), (func() _dafny.Int {
			if (p1).Contains(_145_v99.F7()) {
				return (p1).Multiplicity(_145_v99.F7())
			}
			return p0
		})()), p0, _145_v99.F7())
		_146_v100 = _nw23
		(globalState).F6 = (_146_v100).F8()
		(_145_v99).F7_set_(_74_v36)
	} else {
		(globalState).F3 = p0
		_74_v36 = _74_v36
		var _147_v101 T0
		_ = _147_v101
		var _nw24 *C0 = New_C0_()
		_ = _nw24
		_nw24.Ctor__((p1).Cardinality(), p0, Companion_Default___.Fm3(p0, globalState))
		_147_v101 = _nw24
		var _148_v102 _dafny.Set
		_ = _148_v102
		_148_v102 = _dafny.SetOf(_147_v101)
		if !(_148_v102).Contains(_147_v101) {
			(globalState).F3 = (p0).Minus(p0)
			var _149_v103 _dafny.Array
			_ = _149_v103
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_4
			var _nw25 _dafny.Array
			_ = _nw25
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw25 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_150_v101 T0) func(_dafny.Int) bool {
					return func(_151_i13 _dafny.Int) bool {
						return _150_v101.F7()
					}
				})(_147_v101)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw25 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw25).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw25).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_149_v103 = _nw25
			var _152_v104 _dafny.MultiSet
			_ = _152_v104
			_152_v104 = _dafny.MultiSetOf(_149_v103)
			_152_v104 = (_dafny.MultiSetOf(_149_v103)).Union(_152_v104)
			var _153_v105 *C0
			_ = _153_v105
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__(p0, (_dafny.Zero).Minus(p0), _74_v36)
			_153_v105 = _nw26
			var _154_v106 D7
			_ = _154_v106
			_154_v106 = Companion_D7_.Create_DC20_(true, false, _153_v105)
			_153_v105 = (_154_v106).Dtor_cf32()
			var _155_v107 *C0
			_ = _155_v107
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(Companion_Default___.SafeDivisionInt(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _147_v101)).Cardinality()), Companion_Default___.SafeDivisionInt((_153_v105).F9(), (_153_v105).F8()), Companion_Default___.Fm3((_153_v105).F8(), globalState))
			_155_v107 = _nw27
			_153_v105 = _153_v105
		} else {
			(_147_v101).F7_set_((p2).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool))
			var _156_v108 _dafny.Array
			_ = _156_v108
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw28
			_156_v108 = _nw28
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_156_v108), 0))
			_ = _index21
			(_156_v108).ArraySet1(false, (_index21).Int())
			var _157_v109 _dafny.Map
			_ = _157_v109
			_157_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_156_v108), 0))
			_ = _index22
			(_156_v108).ArraySet1(!(p1).Equals((func() _dafny.MultiSet {
				if !(_147_v101.F7()) {
					return (func() _dafny.MultiSet {
						if (_157_v109).Contains(p0) {
							return (_157_v109).Get(p0).(_dafny.MultiSet)
						}
						return p1
					})()
				}
				return p1
			})()), (_index22).Int())
			var _158_v110 _dafny.Map
			_ = _158_v110
			_158_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_156_v108).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_156_v108), 0))).Int()).(bool))
			var _159_v111 _dafny.Sequence
			_ = _159_v111
			_159_v111 = _dafny.SeqOf((_158_v110).Cardinality())
			var _160_v112 _dafny.Sequence
			_ = _160_v112
			_160_v112 = _dafny.SeqOf(_dafny.IntOfUint32((_159_v111).Cardinality()))
			var _161_v113 D8
			_ = _161_v113
			_161_v113 = Companion_D8_.Create_DC22_(_dafny.SetOf(_159_v111))
			var _162_v114 _dafny.Set
			_ = _162_v114
			_162_v114 = _dafny.SetOf(_160_v112, _159_v111)
			_161_v113 = Companion_D8_.Create_DC22_(_162_v114)
			var _163_v115 *C0
			_ = _163_v115
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__((_147_v101).Fm4(globalState), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p2)).Cardinality(), (_156_v108).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_156_v108), 0))).Int()).(bool))
			_163_v115 = _nw29
			var _164_v116 _dafny.Map
			_ = _164_v116
			_164_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v115, p0)
			_164_v116 = (_164_v116).Merge(_164_v116)
			_147_v101 = _147_v101
		}
		var _165_v117 *C0
		_ = _165_v117
		var _nw30 *C0 = New_C0_()
		_ = _nw30
		_nw30.Ctor__(_dafny.IntOfUint32((_75_v37).Cardinality()), p0, _147_v101.F7())
		_165_v117 = _nw30
		var _166_v118 D4
		_ = _166_v118
		_166_v118 = Companion_D4_.Create_DC14_(_165_v117, _147_v101.F7())
		var _167_v119 *C0
		_ = _167_v119
		var _nw31 *C0 = New_C0_()
		_ = _nw31
		_nw31.Ctor__(p0, (p0).Minus(p0), (_166_v118).Dtor_cf22())
		_167_v119 = _nw31
		var _168_v120 D0
		_ = _168_v120
		_168_v120 = Companion_D0_.Create_DC0_(_74_v36)
		var _pat_let_tv1 = _147_v101
		_ = _pat_let_tv1
		(_147_v101).F7_set_((func(_pat_let2_0 D0) D0 {
			return func(_169_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let3_0 bool) D0 {
					return func(_170_dt__update_hcf0_h0 bool) D0 {
						return Companion_D0_.Create_DC0_(_170_dt__update_hcf0_h0)
					}(_pat_let3_0)
				}(_pat_let_tv1.F7())
			}(_pat_let2_0)
		}(_168_v120)).Dtor_cf0())
	}
	var _171_v121 _dafny.Array
	_ = _171_v121
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(7)
	_ = _len0_5
	var _nw32 _dafny.Array
	_ = _nw32
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw32 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Int = func(_172_i14 _dafny.Int) _dafny.Int {
			return (_172_i14).Times(_dafny.IntOfInt64(-771))
		}
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw32 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw32).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw32).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_171_v121 = _nw32
	var _173_v122 _dafny.Map
	_ = _173_v122
	_173_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v36, _171_v121)
	_173_v122 = (_173_v122).Update(Companion_Default___.Fm3(p0, globalState), _171_v121)
	var _174_v123 _dafny.MultiSet
	_ = _174_v123
	_174_v123 = _dafny.MultiSetOf(p2)
	var _175_v124 D0
	_ = _175_v124
	_175_v124 = Companion_D0_.Create_DC1_(_dafny.MultiSetOf(p2, p2), (_dafny.Zero).Minus(p0))
	var _176_v125 _dafny.MultiSet
	_ = _176_v125
	_176_v125 = _dafny.MultiSetOf(_175_v124)
	var _hi2 _dafny.Int = ((_dafny.MultiSetOf(Companion_D0_.Create_DC1_(_174_v123, p0))).Union(_176_v125)).Cardinality()
	_ = _hi2
	for _177_i15 := p0; _177_i15.Cmp(_hi2) < 0; _177_i15 = _177_i15.Plus(_dafny.One) {
		var _178_v126 _dafny.Sequence
		_ = _178_v126
		_178_v126 = _dafny.SeqOf(p0)
		var _source1 D1 = Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Concatenate(_178_v126, _178_v126))
		_ = _source1
		if _source1.Is_DC4() {
			r0 = _75_v37
			var _179_v127 _dafny.Array
			_ = _179_v127
			var _nw33 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw33
			_179_v127 = _nw33
			var _180_v128 *C0
			_ = _180_v128
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__(_177_i15, (_178_v126).Select((Companion_Default___.SafeIndex(_177_i15, _dafny.IntOfUint32((_178_v126).Cardinality()))).Uint32()).(_dafny.Int), _74_v36)
			_180_v128 = _nw34
			var _181_v129 _dafny.Map
			_ = _181_v129
			_181_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_180_v128, _180_v128)
			var _182_v130 _dafny.Map
			_ = _182_v130
			_182_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v36, _181_v129)
			var _183_v131 _dafny.CodePoint
			_ = _183_v131
			_183_v131 = _dafny.CodePoint('q')
			var _184_v132 _dafny.MultiSet
			_ = _184_v132
			_184_v132 = _dafny.MultiSetOf(_dafny.IntOfInt64(480), Companion_Default___.Fm0(((func() _dafny.Map {
				if (_182_v130).Contains(_74_v36) {
					return (_182_v130).Get(_74_v36).(_dafny.Map)
				}
				return _181_v129
			})()).Cardinality(), _74_v36, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(927))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_185_v128 *C0) func(_dafny.Int) _dafny.Int {
				return func(_186_i16 _dafny.Int) _dafny.Int {
					return (_185_v128).F9()
				}
			})(_180_v128))), _183_v131, globalState), (_dafny.Zero).Minus(_dafny.IntOfUint32((_178_v126).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(474))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_187_v131 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_188_i17 _dafny.Int) _dafny.CodePoint {
					return _187_v131
				}
			})(_183_v131)))).Cardinality()))
			var _189_v133 *C0
			_ = _189_v133
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(p0, (_184_v132).Cardinality(), _74_v36)
			_189_v133 = _nw35
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_179_v127), 0))
			_ = _index23
			(_179_v127).ArraySet1(_180_v128, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_179_v127), 0))
			_ = _index24
			(_179_v127).ArraySet1(_189_v133, (_index24).Int())
			(globalState).F6 = ((_180_v128).F9()).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-448), _177_i15))
			(globalState).F3 = (_189_v133).F9()
		} else if _source1.Is_DC3() {
			var _190___mcc_h0 _dafny.Sequence = _source1.Get_().(D1_DC3).Cf4
			_ = _190___mcc_h0
			var _191_cf4 _dafny.Sequence = _190___mcc_h0
			_ = _191_cf4
			var _192_v134 _dafny.Sequence
			_ = _192_v134
			_192_v134 = _75_v37
			var _193_v135 _dafny.MultiSet
			_ = _193_v135
			_193_v135 = _dafny.MultiSetOf(_191_cf4, _191_cf4)
			var _194_v136 _dafny.CodePoint
			_ = _194_v136
			_194_v136 = _dafny.CodePoint('t')
			var _195_v137 _dafny.Array
			_ = _195_v137
			var _nwElement0_4 _dafny.Sequence = _75_v37
			_ = _nwElement0_4
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(29))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_4, 0)
			(_nw36).ArraySet1(_75_v37, 1)
			(_nw36).ArraySet1(_75_v37, 2)
			(_nw36).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-246))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_196_i18 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), 3)
			(_nw36).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("iqglrvh"), 4)
			(_nw36).ArraySet1(_75_v37, 5)
			(_nw36).ArraySet1((func() _dafny.Sequence {
				if _74_v36 {
					return _75_v37
				}
				return _75_v37
			})(), 6)
			(_nw36).ArraySet1((_192_v134), 7)
			(_nw36).ArraySet1(_75_v37, 8)
			(_nw36).ArraySet1(_75_v37, 9)
			(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_75_v37, _75_v37), 10)
			(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_75_v37, _75_v37), 11)
			(_nw36).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ugemjeiny"), 12)
			(_nw36).ArraySet1(_75_v37, 13)
			(_nw36).ArraySet1(_75_v37, 14)
			(_nw36).ArraySet1(_75_v37, 15)
			(_nw36).ArraySet1(_75_v37, 16)
			(_nw36).ArraySet1(_75_v37, 17)
			(_nw36).ArraySet1(_75_v37, 18)
			(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10(p0, _74_v36, _177_i15, _193_v135, globalState), _75_v37), 19)
			(_nw36).ArraySet1(_75_v37, 20)
			(_nw36).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("foi"), _75_v37), 21)
			(_nw36).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ptuxtb"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ptuxtb")).Cardinality()))).Uint32(), _194_v136), 22)
			(_nw36).ArraySet1(_75_v37, 23)
			(_nw36).ArraySet1(_75_v37, 24)
			(_nw36).ArraySet1(_75_v37, 25)
			(_nw36).ArraySet1(_75_v37, 26)
			(_nw36).ArraySet1(_75_v37, 27)
			(_nw36).ArraySet1(_75_v37, 28)
			_195_v137 = _nw36
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_195_v137), 0))
			_ = _index25
			(_195_v137).ArraySet1(_75_v37, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_195_v137), 0))
			_ = _index26
			(_195_v137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jquu"), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_171_v121), 0))
			_ = _index27
			(_171_v121).ArraySet1(p0, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_171_v121), 0))
			_ = _index28
			(_171_v121).ArraySet1(_177_i15, (_index28).Int())
			var _197_v138 _dafny.Map
			_ = _197_v138
			_197_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_171_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_171_v121), 0))).Int()).(_dafny.Int))
			var _198_v139 _dafny.Map
			_ = _198_v139
			_198_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_197_v138, _194_v136)
			_194_v136 = (func() _dafny.CodePoint {
				if (_198_v139).Contains(func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate((Companion_Default___.Fm17(_74_v36, _74_v36, _74_v36, globalState)).Elements()); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _199_v140 _dafny.Sequence
						_199_v140 = interface{}(_compr_8).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm17(_74_v36, _74_v36, _74_v36, globalState), _199_v140) {
							_coll8.Add(_199_v140, _177_i15)
						}
					}
					return _coll8.ToMap()
				}()) {
					return (_198_v139).Get(func() _dafny.Map {
						var _coll9 = _dafny.NewMapBuilder()
						_ = _coll9
						for _iter9 := _dafny.Iterate((Companion_Default___.Fm17(_74_v36, _74_v36, _74_v36, globalState)).Elements()); ; {
							_compr_9, _ok9 := _iter9()
							if !_ok9 {
								break
							}
							var _200_v140 _dafny.Sequence
							_200_v140 = interface{}(_compr_9).(_dafny.Sequence)
							if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm17(_74_v36, _74_v36, _74_v36, globalState), _200_v140) {
								_coll9.Add(_200_v140, _177_i15)
							}
						}
						return _coll9.ToMap()
					}()).(_dafny.CodePoint)
				}
				return _dafny.CodePoint('g')
			})()
			var _201_v141 *C0
			_ = _201_v141
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__(_dafny.IntOfUint32((p2).Cardinality()), _177_i15, false)
			_201_v141 = _nw37
		} else {
			var _202___mcc_h1 D1 = _source1.Get_().(D1_DC5).Cf5
			_ = _202___mcc_h1
			var _203_cf5 D1 = _202___mcc_h1
			_ = _203_cf5
			_74_v36 = _74_v36
			_74_v36 = !(_dafny.MultiSetOf(_75_v37, _75_v37, _75_v37, _75_v37, _75_v37)).Contains(_75_v37)
			var _204_v142 _dafny.MultiSet
			_ = _204_v142
			_204_v142 = _dafny.MultiSetOf(_74_v36)
			_204_v142 = p1
			var _205_v143 _dafny.Map
			_ = _205_v143
			_205_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_74_v36, _177_i15)
			_204_v142 = ((Companion_Default___.Fm12(_177_i15, _177_i15, _205_v143, globalState)).Update(_74_v36, Companion_Default___.Abs(_177_i15))).Union(p1)
		}
		var _206_v144 _dafny.Set
		_ = _206_v144
		_206_v144 = _dafny.SetOf(_177_i15)
		var _207_v145 _dafny.MultiSet
		_ = _207_v145
		_207_v145 = _dafny.MultiSetOf(_206_v144)
		var _208_v146 T0
		_ = _208_v146
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__(((func() _dafny.Int {
			if (_207_v145).Contains(_206_v144) {
				return (_207_v145).Multiplicity(_206_v144)
			}
			return p0
		})()).Plus(_177_i15), (_dafny.Zero).Minus(_177_i15), _74_v36)
		_208_v146 = _nw38
		var _209_v147 _dafny.Map
		_ = _209_v147
		_209_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_i15, _208_v146)
		_208_v146 = (func() T0 {
			if (_209_v147).Contains(_177_i15) {
				return (_209_v147).Get(_177_i15).(T0)
			}
			return _208_v146
		})()
		var _210_v148 _dafny.Array
		_ = _210_v148
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
		_ = _nw39
		_210_v148 = _nw39
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_210_v148), 0))
		_ = _index29
		(_210_v148).ArraySet1(_171_v121, (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_210_v148), 0))
		_ = _index30
		var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw40
		(_210_v148).ArraySet1(_nw40, (_index30).Int())
		(_208_v146).F7_set_(_208_v146.F7())
	}
	r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_75_v37, _75_v37), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("oc"), _75_v37))
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _211_v0 _dafny.Sequence
	_ = _211_v0
	_211_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qfxbbb")
	var _212_v1 _dafny.Int
	_ = _212_v1
	_212_v1 = _dafny.IntOfInt64(779)
	var _213_v2 _dafny.CodePoint
	_ = _213_v2
	_213_v2 = _dafny.CodePoint('p')
	var _214_v3 _dafny.Map
	_ = _214_v3
	_214_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v1, _213_v2)
	var _215_v4 bool
	_ = _215_v4
	_215_v4 = false
	var _216_v5 _dafny.Map
	_ = _216_v5
	_216_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
		if (_214_v3).Contains(_dafny.IntOfUint32((_211_v0).Cardinality())) {
			return (_214_v3).Get(_dafny.IntOfUint32((_211_v0).Cardinality())).(_dafny.CodePoint)
		}
		return _213_v2
	})(), _215_v4)
	var _217_v6 _dafny.MultiSet
	_ = _217_v6
	_217_v6 = _dafny.MultiSetOf(_dafny.IntOfInt64(-166))
	var _218_v7 _dafny.MultiSet
	_ = _218_v7
	_218_v7 = _dafny.MultiSetOf(_dafny.MultiSetOf((_217_v6).Cardinality()))
	var _219_globalState *GlobalState
	_ = _219_globalState
	var _nw41 *GlobalState = New_GlobalState_()
	_ = _nw41
	_nw41.Ctor__(true, _211_v0, true, _dafny.IntOfInt64(225), (_216_v5).Merge(_216_v5), _218_v7, _dafny.IntOfInt64(852))
	_219_globalState = _nw41
	(_219_globalState).F6 = _dafny.IntOfInt64(155)
	var _hi3 _dafny.Int = _212_v1
	_ = _hi3
	for _220_i0 := (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_212_v1), (Companion_Default___.SafeIndex((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_211_v0).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _253_v8 _dafny.CodePoint
			_253_v8 = interface{}(_compr_10).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_211_v0, _253_v8) {
				_coll10.Add(_253_v8, _212_v1)
			}
		}
		return _coll10.ToMap()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_212_v1)).Cardinality()))).Uint32(), _212_v1)).Cardinality())).Plus(_dafny.IntOfInt64(83)); _220_i0.Cmp(_hi3) < 0; _220_i0 = _220_i0.Plus(_dafny.One) {
		var _221_v9 _dafny.Sequence
		_ = _221_v9
		_221_v9 = _dafny.SeqOf(_220_i0, _dafny.IntOfInt64(832), _212_v1, _212_v1)
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_221_v9, _221_v9) {
			var _222_v10 _dafny.Array
			_ = _222_v10
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_6
			var _nw42 _dafny.Array
			_ = _nw42
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw42 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.MultiSet = (func(_223_v6 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_224_i1 _dafny.Int) _dafny.MultiSet {
						return _223_v6
					}
				})(_217_v6)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw42 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw42).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw42).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_222_v10 = _nw42
			_222_v10 = _222_v10
			var _225_v11 _dafny.MultiSet
			_ = _225_v11
			_225_v11 = _dafny.MultiSetOf(_215_v4, !(_215_v4))
			var _226_v12 _dafny.Sequence
			_ = _226_v12
			var _out0 _dafny.Sequence
			_ = _out0
			_out0 = Companion_Default___.M0(_220_i0, (_225_v11).Intersection(_225_v11), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_215_v4), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_221_v9).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_215_v4)).Cardinality()))).Uint32(), _215_v4), _219_globalState)
			_226_v12 = _out0
			var _rhs20 _dafny.Int = (func() _dafny.Int {
				if _215_v4 {
					return _220_i0
				}
				return _dafny.IntOfUint32((_211_v0).Cardinality())
			})()
			_ = _rhs20
			var _rhs21 _dafny.Int = _220_i0
			_ = _rhs21
			var _lhs16 *GlobalState = _219_globalState
			_ = _lhs16
			var _lhs17 *GlobalState = _219_globalState
			_ = _lhs17
			_lhs16.F6 = _rhs20
			_lhs17.F3 = _rhs21
			_215_v4 = _215_v4
			var _227_v13 D0
			_ = _227_v13
			_227_v13 = Companion_D0_.Create_DC0_(_215_v4)
			_215_v4 = (_227_v13).Dtor_cf0()
		} else {
			var _228_v14 _dafny.Set
			_ = _228_v14
			_228_v14 = _dafny.SetOf(_220_i0, _212_v1)
			var _rhs22 _dafny.Int = Companion_Default___.Fm0((Companion_Default___.Fm1(_219_globalState)).Dtor_cf2(), true, _221_v9, _213_v2, _219_globalState)
			_ = _rhs22
			var _rhs23 _dafny.Int = Companion_Default___.SafeDivisionInt(_212_v1, (_228_v14).Cardinality())
			_ = _rhs23
			var _lhs18 *GlobalState = _219_globalState
			_ = _lhs18
			var _lhs19 *GlobalState = _219_globalState
			_ = _lhs19
			_lhs18.F3 = _rhs22
			_lhs19.F3 = _rhs23
			var _229_v15 _dafny.Map
			_ = _229_v15
			_229_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, _dafny.MultiSetOf(_dafny.IntOfInt64(928), _212_v1))
			var _230_v16 _dafny.Map
			_ = _230_v16
			_230_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
				if (_229_v15).Contains(_215_v4) {
					return (_229_v15).Get(_215_v4).(_dafny.MultiSet)
				}
				return _217_v6
			})(), false)
			_228_v14 = Companion_Default___.Fm2(_230_v16, _dafny.CodePoint('d'), _219_globalState)
			var _231_v17 _dafny.Map
			_ = _231_v17
			_231_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_i0, _212_v1)
			var _232_v18 _dafny.Sequence
			_ = _232_v18
			_232_v18 = _dafny.SeqOf(_215_v4, _215_v4)
			var _233_v19 _dafny.MultiSet
			_ = _233_v19
			_233_v19 = _dafny.MultiSetOf(_232_v18)
			var _234_v20 _dafny.Map
			_ = _234_v20
			_234_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, false)
			var _235_v21 D0
			_ = _235_v21
			_235_v21 = Companion_D0_.Create_DC1_(_233_v19, (_234_v20).Cardinality())
			var _236_v22 _dafny.Map
			_ = _236_v22
			_236_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v0, _215_v4)
			var _237_v23 _dafny.Map
			_ = _237_v23
			_237_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v22, _215_v4)
			var _238_v24 _dafny.Set
			_ = _238_v24
			_238_v24 = _dafny.SetOf(_215_v4)
			var _rhs24 _dafny.Int = (func() _dafny.Int {
				if Companion_Default___.Fm3(_212_v1, _219_globalState) {
					return (func() _dafny.Int {
						if _215_v4 {
							return _212_v1
						}
						return _220_i0
					})()
				}
				return (_235_v21).Dtor_cf2()
			})()
			_ = _rhs24
			var _rhs25 _dafny.Map = _231_v17
			_ = _rhs25
			var _rhs26 bool = (func() bool {
				if (_237_v23).Contains(_236_v22) {
					return (_237_v23).Get(_236_v22).(bool)
				}
				return (_238_v24).IsProperSubsetOf(_238_v24)
			})()
			_ = _rhs26
			var _rhs27 bool = _215_v4
			_ = _rhs27
			var _lhs20 *GlobalState = _219_globalState
			_ = _lhs20
			_lhs20.F3 = _rhs24
			_231_v17 = _rhs25
			_215_v4 = _rhs26
			_215_v4 = _rhs27
			_215_v4 = _215_v4
			var _239_v25 _dafny.Array
			_ = _239_v25
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw43
			_239_v25 = _nw43
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_239_v25), 0))
			_ = _index31
			(_239_v25).ArraySet1(_220_i0, (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_239_v25), 0))
			_ = _index32
			(_239_v25).ArraySet1(_212_v1, (_index32).Int())
		}
		var _240_v26 _dafny.Sequence
		_ = _240_v26
		_240_v26 = _dafny.SeqOf(_215_v4)
		var _241_v27 D0
		_ = _241_v27
		_241_v27 = Companion_D0_.Create_DC0_(false)
		var _242_v28 _dafny.Array
		_ = _242_v28
		var _nwElement0_5 bool = _215_v4
		_ = _nwElement0_5
		var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(10))
		_ = _nw44
		(_nw44).ArraySet1(_nwElement0_5, 0)
		(_nw44).ArraySet1((_240_v26).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.IntOfUint32((_240_v26).Cardinality()))).Uint32()).(bool), 1)
		(_nw44).ArraySet1(_215_v4, 2)
		(_nw44).ArraySet1(_215_v4, 3)
		(_nw44).ArraySet1(_215_v4, 4)
		(_nw44).ArraySet1(_215_v4, 5)
		(_nw44).ArraySet1(true, 6)
		(_nw44).ArraySet1(_215_v4, 7)
		(_nw44).ArraySet1(_215_v4, 8)
		(_nw44).ArraySet1((_241_v27).Dtor_cf0(), 9)
		_242_v28 = _nw44
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_242_v28), 0))
		_ = _index33
		(_242_v28).ArraySet1(_215_v4, (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_242_v28), 0))
		_ = _index34
		(_242_v28).ArraySet1(_215_v4, (_index34).Int())
		var _243_v29 _dafny.Set
		_ = _243_v29
		_243_v29 = _dafny.SetOf(_213_v2)
		var _244_v30 _dafny.MultiSet
		_ = _244_v30
		_244_v30 = _dafny.MultiSetOf((_243_v29).Intersection(_243_v29), _243_v29)
		var _245_v31 _dafny.Sequence
		_ = _245_v31
		_245_v31 = _dafny.SeqOf(_243_v29)
		_244_v30 = ((_dafny.MultiSetFromSeq(_245_v31)).Union(_244_v30)).Update(_243_v29, Companion_Default___.Abs(_220_i0))
		var _246_v32 _dafny.Map
		_ = _246_v32
		_246_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _220_i0)
		var _247_v33 _dafny.Sequence
		_ = _247_v33
		_247_v33 = _dafny.SeqOf(_242_v28, _242_v28)
		var _248_v34 _dafny.Sequence
		_ = _248_v34
		_248_v34 = _dafny.SeqOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_212_v1)), _217_v6)
		var _249_v35 D1
		_ = _249_v35
		_249_v35 = Companion_D1_.Create_DC3_(_221_v9)
		var _250_v36 _dafny.Array
		_ = _250_v36
		var _nwElement0_6 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-332))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_251_i2 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality())
		})), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_211_v0).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-332))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_252_i2 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality())
		}))).Cardinality()))).Uint32(), (_246_v32).Cardinality())).Cardinality())
		_ = _nwElement0_6
		var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(21))
		_ = _nw45
		(_nw45).ArraySet1(_nwElement0_6, 0)
		(_nw45).ArraySet1(_dafny.IntOfInt64(742), 1)
		(_nw45).ArraySet1((_246_v32).Cardinality(), 2)
		(_nw45).ArraySet1(_212_v1, 3)
		(_nw45).ArraySet1(_220_i0, 4)
		(_nw45).ArraySet1(_212_v1, 5)
		(_nw45).ArraySet1(_220_i0, 6)
		(_nw45).ArraySet1(_dafny.IntOfInt64(967), 7)
		(_nw45).ArraySet1(_220_i0, 8)
		(_nw45).ArraySet1(_dafny.IntOfInt64(568), 9)
		(_nw45).ArraySet1(_dafny.IntOfUint32((_247_v33).Cardinality()), 10)
		(_nw45).ArraySet1(_212_v1, 11)
		(_nw45).ArraySet1(_dafny.IntOfInt64(869), 12)
		(_nw45).ArraySet1(Companion_Default___.SafeModuloInt(_212_v1, _220_i0), 13)
		(_nw45).ArraySet1(_212_v1, 14)
		(_nw45).ArraySet1((_220_i0).Minus(_212_v1), 15)
		(_nw45).ArraySet1(_212_v1, 16)
		(_nw45).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_dafny.IntOfUint32((_248_v34).Cardinality()), (_242_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_242_v28), 0))).Int()).(bool), _221_v9, _213_v2, _219_globalState), Companion_Default___.Fm0(_dafny.IntOfUint32((_211_v0).Cardinality()), (_242_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_242_v28), 0))).Int()).(bool), (_249_v35).Dtor_cf4(), _213_v2, _219_globalState)), 17)
		(_nw45).ArraySet1((_dafny.Zero).Minus(_220_i0), 18)
		(_nw45).ArraySet1(_dafny.IntOfInt64(975), 19)
		(_nw45).ArraySet1(((_dafny.Zero).Minus(_220_i0)).Plus(_220_i0), 20)
		_250_v36 = _nw45
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_250_v36), 0))
		_ = _index35
		(_250_v36).ArraySet1(_212_v1, (_index35).Int())
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_250_v36), 0))
		_ = _index36
		var _rhs28 _dafny.Int = _220_i0
		_ = _rhs28
		var _rhs29 _dafny.Int = _220_i0
		_ = _rhs29
		var _rhs30 _dafny.Int = _220_i0
		_ = _rhs30
		var _lhs21 _dafny.Array = _250_v36
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_250_v36), 0))
		_ = _lhs22
		var _lhs23 *GlobalState = _219_globalState
		_ = _lhs23
		(_lhs21).ArraySet1(_rhs28, (_lhs22).Int())
		_212_v1 = _rhs29
		_lhs23.F6 = _rhs30
	}
	var _254_v37 _dafny.Set
	_ = _254_v37
	_254_v37 = _dafny.SetOf(_213_v2)
	var _255_v38 _dafny.Set
	_ = _255_v38
	_255_v38 = _dafny.SetOf(_254_v37, _254_v37, _254_v37, _dafny.SetOf(_213_v2, _213_v2), _254_v37)
	var _256_v39 _dafny.Sequence
	_ = _256_v39
	_256_v39 = _dafny.SeqOf(_255_v38)
	var _257_v40 _dafny.Map
	_ = _257_v40
	_257_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v1, _212_v1)
	var _hi4 _dafny.Int = (_212_v1).Minus((_257_v40).Cardinality())
	_ = _hi4
	for _258_i3 := ((_256_v39).Select((Companion_Default___.SafeIndex(_212_v1, _dafny.IntOfUint32((_256_v39).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(); _258_i3.Cmp(_hi4) < 0; _258_i3 = _258_i3.Plus(_dafny.One) {
		if _215_v4 {
			var _259_v41 _dafny.Array
			_ = _259_v41
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw46
			_259_v41 = _nw46
			var _260_v42 _dafny.Array
			_ = _260_v42
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_7
			var _nw47 _dafny.Array
			_ = _nw47
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw47 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = (func(_261_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_262_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_262_i4, _261_v1)
					}
				})(_212_v1)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw47 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw47).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw47).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_260_v42 = _nw47
			var _263_v43 _dafny.Map
			_ = _263_v43
			_263_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_259_v41, _260_v42)
			_263_v43 = (_263_v43).Update(_259_v41, _260_v42)
			var _264_v44 *C0
			_ = _264_v44
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__(_212_v1, Companion_Default___.SafeModuloInt(_258_i3, _dafny.IntOfInt64(163)), _215_v4)
			_264_v44 = _nw48
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_260_v42), 0))
			_ = _index37
			(_260_v42).ArraySet1(Companion_Default___.SafeDivisionInt((_264_v44).F8(), _258_i3), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_260_v42), 0))
			_ = _index38
			(_260_v42).ArraySet1((_dafny.Zero).Minus(_212_v1), (_index38).Int())
			var _265_v45 _dafny.Set
			_ = _265_v45
			_265_v45 = _dafny.SetOf(_215_v4)
			var _266_v46 _dafny.MultiSet
			_ = _266_v46
			_266_v46 = _dafny.MultiSetOf(Companion_Default___.Fm3(_212_v1, _219_globalState))
			var _267_v47 D1
			_ = _267_v47
			_267_v47 = Companion_D1_.Create_DC4_()
			var _268_v48 _dafny.Map
			_ = _268_v48
			_268_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v47, _215_v4)
			var _269_v49 _dafny.Sequence
			_ = _269_v49
			var _out1 _dafny.Sequence
			_ = _out1
			_out1 = Companion_Default___.M0(((_265_v45).Cardinality()).Plus(_258_i3), _266_v46, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm6((_264_v44).F9(), _dafny.CodePoint('p'), (func() bool {
				if (_268_v48).Contains(_267_v47) {
					return (_268_v48).Get(_267_v47).(bool)
				}
				return _215_v4
			})(), _215_v4, _219_globalState), Companion_Default___.Fm6(_212_v1, _213_v2, _215_v4, _215_v4, _219_globalState)), _219_globalState)
			_269_v49 = _out1
			var _270_v50 _dafny.Sequence
			_ = _270_v50
			_270_v50 = _dafny.SeqOf(_215_v4, _215_v4)
			var _271_v51 _dafny.Sequence
			_ = _271_v51
			_271_v51 = _dafny.SeqOf((_270_v50).Select((Companion_Default___.SafeIndex((_264_v44).F8(), _dafny.IntOfUint32((_270_v50).Cardinality()))).Uint32()).(bool))
			var _272_v52 _dafny.Sequence
			_ = _272_v52
			var _out2 _dafny.Sequence
			_ = _out2
			_out2 = Companion_Default___.M0(_dafny.IntOfInt64(496), _266_v46, _271_v51, _219_globalState)
			_272_v52 = _out2
		} else {
			_215_v4 = (_217_v6).IsDisjointFrom(_217_v6)
			var _273_v53 _dafny.Array
			_ = _273_v53
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_8
			var _nw49 _dafny.Array
			_ = _nw49
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw49 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) bool = (func(_274_v4 bool) func(_dafny.Int) bool {
					return func(_275_i5 _dafny.Int) bool {
						return _274_v4
					}
				})(_215_v4)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw49 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw49).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw49).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_273_v53 = _nw49
			var _276_v54 _dafny.Map
			_ = _276_v54
			_276_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_i3, _273_v53)
			_212_v1 = (((_276_v54).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v1, _273_v53))).Merge(_276_v54)).Cardinality()
			var _277_v55 _dafny.Map
			_ = _277_v55
			_277_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, _258_i3)
			var _278_v56 _dafny.Sequence
			_ = _278_v56
			_278_v56 = _dafny.SeqOf(!(true))
			var _279_v57 _dafny.Sequence
			_ = _279_v57
			_279_v57 = _dafny.SeqOf(_212_v1, _212_v1)
			var _280_v58 _dafny.Map
			_ = _280_v58
			_280_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_277_v55).Cardinality(), (_278_v56).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_211_v0).Cardinality())), _dafny.IntOfUint32((_278_v56).Cardinality()))).Uint32()).(bool), _279_v57, _213_v2, _219_globalState), _215_v4)
			var _281_v59 *C0
			_ = _281_v59
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__((_dafny.IntOfInt64(969)).Minus(_258_i3), Companion_Default___.SafeDivisionInt(_212_v1, (_277_v55).Cardinality()), (func() bool {
				if (_280_v58).Contains(_212_v1) {
					return (_280_v58).Get(_212_v1).(bool)
				}
				return _215_v4
			})())
			_281_v59 = _nw50
			var _282_v60 _dafny.Array
			_ = _282_v60
			var _nw51 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(12))
			_ = _nw51
			_282_v60 = _nw51
			var _283_v61 D1
			_ = _283_v61
			_283_v61 = Companion_D1_.Create_DC3_(_dafny.SeqOf(_dafny.IntOfInt64(146)))
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_282_v60), 0))
			_ = _index39
			(_282_v60).ArraySet1(_283_v61, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_282_v60), 0))
			_ = _index40
			(_282_v60).ArraySet1(_283_v61, (_index40).Int())
			var _284_v62 _dafny.MultiSet
			_ = _284_v62
			_284_v62 = _dafny.MultiSetOf(_215_v4)
			var _285_v63 _dafny.Sequence
			_ = _285_v63
			var _out3 _dafny.Sequence
			_ = _out3
			_out3 = Companion_Default___.M0((func() _dafny.Int {
				if _215_v4 {
					return (_281_v59).F9()
				}
				return _dafny.IntOfUint32((_278_v56).Cardinality())
			})(), (_dafny.MultiSetFromSeq(_278_v56)).Intersection(_284_v62), _dafny.Companion_Sequence_.Update(_278_v56, (Companion_Default___.SafeIndex(_258_i3, _dafny.IntOfUint32((_278_v56).Cardinality()))).Uint32(), _215_v4), _219_globalState)
			_285_v63 = _out3
		}
		(_219_globalState).F6 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_258_i3), _dafny.IntOfUint32((_211_v0).Cardinality())), _212_v1)
		var _286_v64 _dafny.Sequence
		_ = _286_v64
		_286_v64 = _dafny.SeqOf(_212_v1)
		var _287_v65 _dafny.Array
		_ = _287_v65
		var _nwElement0_7 bool = true
		_ = _nwElement0_7
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(11))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_7, 0)
		(_nw52).ArraySet1(_215_v4, 1)
		(_nw52).ArraySet1(_215_v4, 2)
		(_nw52).ArraySet1(false, 3)
		(_nw52).ArraySet1(!(_215_v4), 4)
		(_nw52).ArraySet1(_215_v4, 5)
		(_nw52).ArraySet1(!(_215_v4), 6)
		(_nw52).ArraySet1(_215_v4, 7)
		(_nw52).ArraySet1(Companion_Default___.Fm3((_dafny.Zero).Minus(_258_i3), _219_globalState), 8)
		(_nw52).ArraySet1(!(false), 9)
		(_nw52).ArraySet1(_215_v4, 10)
		_287_v65 = _nw52
		var _288_v66 _dafny.Map
		_ = _288_v66
		_288_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_286_v64).Select((Companion_Default___.SafeIndex(_258_i3, _dafny.IntOfUint32((_286_v64).Cardinality()))).Uint32()).(_dafny.Int), _287_v65)
		var _289_v67 _dafny.Set
		_ = _289_v67
		_289_v67 = _dafny.SetOf(_215_v4)
		_288_v66 = (_288_v66).Update((_289_v67).Cardinality(), _287_v65)
		if _215_v4 {
			var _290_v68 _dafny.MultiSet
			_ = _290_v68
			_290_v68 = _dafny.MultiSetOf(_215_v4)
			var _rhs31 _dafny.Int = (_212_v1).Times((_dafny.SetOf((_dafny.Zero).Minus(_258_i3), _212_v1, _212_v1)).Cardinality())
			_ = _rhs31
			var _rhs32 bool = !(_290_v68).Equals(_dafny.MultiSetOf(_215_v4))
			_ = _rhs32
			var _lhs24 *GlobalState = _219_globalState
			_ = _lhs24
			_lhs24.F3 = _rhs31
			_215_v4 = _rhs32
			var _291_v69 *C0
			_ = _291_v69
			var _nw53 *C0 = New_C0_()
			_ = _nw53
			_nw53.Ctor__(_258_i3, _258_i3, (_215_v4) == (_215_v4))
			_291_v69 = _nw53
			_215_v4 = !((Companion_D0_.Create_DC0_(_215_v4)).Dtor_cf0())
			var _292_v70 _dafny.Sequence
			_ = _292_v70
			_292_v70 = _dafny.SeqOf(_290_v68)
			var _293_v71 _dafny.Sequence
			_ = _293_v71
			_293_v71 = _dafny.SeqOf(_215_v4)
			var _294_v72 _dafny.Sequence
			_ = _294_v72
			var _out4 _dafny.Sequence
			_ = _out4
			_out4 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt(_258_i3, Companion_Default___.Fm0(_212_v1, _215_v4, _286_v64, _213_v2, _219_globalState)), ((_290_v68).Update(_215_v4, Companion_Default___.Abs((_dafny.Zero).Minus(_258_i3)))).Difference((_292_v70).Select((Companion_Default___.SafeIndex((_291_v69).F9(), _dafny.IntOfUint32((_292_v70).Cardinality()))).Uint32()).(_dafny.MultiSet)), _293_v71, _219_globalState)
			_294_v72 = _out4
			_215_v4 = Companion_Default___.Fm3((_291_v69).F9(), _219_globalState)
		} else {
			_215_v4 = _215_v4
			var _295_v73 _dafny.Sequence
			_ = _295_v73
			_295_v73 = _dafny.SeqOf(_215_v4)
			var _296_v74 _dafny.Sequence
			_ = _296_v74
			_296_v74 = _dafny.SeqOf(_215_v4, (_295_v73).Select((Companion_Default___.SafeIndex((_217_v6).Cardinality(), _dafny.IntOfUint32((_295_v73).Cardinality()))).Uint32()).(bool))
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_287_v65), 0))
			_ = _index41
			(_287_v65).ArraySet1((_295_v73).Select((Companion_Default___.SafeIndex(_212_v1, _dafny.IntOfUint32((_295_v73).Cardinality()))).Uint32()).(bool), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_287_v65), 0))
			_ = _index42
			(_287_v65).ArraySet1(_215_v4, (_index42).Int())
			var _297_v75 _dafny.MultiSet
			_ = _297_v75
			_297_v75 = _dafny.MultiSetOf(_295_v73)
			var _298_v76 D0
			_ = _298_v76
			_298_v76 = Companion_D0_.Create_DC1_(_297_v75, _212_v1)
			var _299_v77 D0
			_ = _299_v77
			_299_v77 = Companion_D0_.Create_DC2_(_298_v76)
			var _pat_let_tv2 = _219_globalState
			_ = _pat_let_tv2
			_299_v77 = func(_pat_let4_0 D0) D0 {
				return func(_300_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let5_0 D0) D0 {
						return func(_301_dt__update_hcf3_h0 D0) D0 {
							return Companion_D0_.Create_DC2_(_301_dt__update_hcf3_h0)
						}(_pat_let5_0)
					}(Companion_Default___.Fm1(_pat_let_tv2))
				}(_pat_let4_0)
			}((func() D0 {
				if _215_v4 {
					return _299_v77
				}
				return _299_v77
			})())
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_287_v65), 0))
			_ = _index43
			(_287_v65).ArraySet1(Companion_Default___.Fm3((func() _dafny.Int {
				if _215_v4 {
					return _258_i3
				}
				return _258_i3
			})(), _219_globalState), (_index43).Int())
			(_219_globalState).F3 = _212_v1
		}
	}
	var _hi5 _dafny.Int = _212_v1
	_ = _hi5
	for _302_i6 := _212_v1; _302_i6.Cmp(_hi5) < 0; _302_i6 = _302_i6.Plus(_dafny.One) {
		var _303_v78 *C0
		_ = _303_v78
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__((_dafny.Zero).Minus((_dafny.Zero).Minus(_212_v1)), _302_i6, _215_v4)
		_303_v78 = _nw54
		_214_v3 = (_214_v3).Update((_303_v78).F9(), _213_v2)
		var _304_v79 _dafny.Sequence
		_ = _304_v79
		_304_v79 = _dafny.SeqOf(_215_v4, _215_v4, _215_v4)
		var _305_v80 _dafny.Sequence
		_ = _305_v80
		var _out5 _dafny.Sequence
		_ = _out5
		_out5 = Companion_Default___.M0((_dafny.Zero).Minus((_303_v78).F8()), _dafny.MultiSetOf(_215_v4, false), _304_v79, _219_globalState)
		_305_v80 = _out5
		(_219_globalState).F3 = (_dafny.Zero).Minus(_302_i6)
	}
	var _306_v81 _dafny.Sequence
	_ = _306_v81
	_306_v81 = _dafny.SeqOf(_212_v1)
	var _307_v82 _dafny.Sequence
	_ = _307_v82
	_307_v82 = _dafny.SeqOf(_212_v1, (_306_v81).Select((Companion_Default___.SafeIndex(_212_v1, _dafny.IntOfUint32((_306_v81).Cardinality()))).Uint32()).(_dafny.Int))
	var _308_v83 D1
	_ = _308_v83
	_308_v83 = Companion_D1_.Create_DC3_(_306_v81)
	var _309_i7 _dafny.Int
	_ = _309_i7
	_309_i7 = _dafny.Zero
	{
		var _pat_let_tv3 = _215_v4
		_ = _pat_let_tv3
		var _pat_let_tv4 = _215_v4
		_ = _pat_let_tv4
		var _pat_let_tv5 = _212_v1
		_ = _pat_let_tv5
		for func(_source2 D1) bool {
			if _source2.Is_DC4() {
				return _pat_let_tv3
			} else if _source2.Is_DC3() {
				var _322___mcc_h0 _dafny.Sequence = _source2.Get_().(D1_DC3).Cf4
				_ = _322___mcc_h0
				var _323_cf4 _dafny.Sequence = _322___mcc_h0
				_ = _323_cf4
				return _pat_let_tv4
			} else {
				var _324___mcc_h1 D1 = _source2.Get_().(D1_DC5).Cf5
				_ = _324___mcc_h1
				var _325_cf5 D1 = _324___mcc_h1
				_ = _325_cf5
				return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nol")).Cardinality())).Cmp(_pat_let_tv5) > 0
			}
		}(_308_v83) {
			{
				if (_309_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_309_i7 = (_309_i7).Plus(_dafny.One)
				var _310_v84 _dafny.Array
				_ = _310_v84
				var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw55
				_310_v84 = _nw55
				var _311_v85 _dafny.Sequence
				_ = _311_v85
				_311_v85 = _dafny.SeqOf(_215_v4)
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))
				_ = _index44
				(_310_v84).ArraySet1(!_dafny.Companion_Sequence_.Equal(_311_v85, _311_v85), (_index44).Int())
				var _312_v86 D0
				_ = _312_v86
				_312_v86 = Companion_D0_.Create_DC0_(_215_v4)
				var _313_v87 *C0
				_ = _313_v87
				var _nw56 *C0 = New_C0_()
				_ = _nw56
				_nw56.Ctor__(_212_v1, _dafny.IntOfInt64(715), (_312_v86).Dtor_cf0())
				_313_v87 = _nw56
				var _314_v88 _dafny.Map
				_ = _314_v88
				_314_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_313_v87, _215_v4)
				var _315_v89 _dafny.MultiSet
				_ = _315_v89
				_315_v89 = _dafny.MultiSetOf(_dafny.SeqOf(!(_215_v4), _215_v4), _311_v85)
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))
				_ = _index45
				(_310_v84).ArraySet1((func() bool {
					if (_314_v88).Contains(_313_v87) {
						return (_314_v88).Get(_313_v87).(bool)
					}
					return (_315_v89).IsSubsetOf(_315_v89)
				})(), (_index45).Int())
				var _316_v90 _dafny.Array
				_ = _316_v90
				var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw57
				_316_v90 = _nw57
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_316_v90), 0))
				_ = _index46
				(_316_v90).ArraySet1((_313_v87).F9(), (_index46).Int())
				var _317_v91 _dafny.Map
				_ = _317_v91
				_317_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, (_310_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))).Int()).(bool))
				var _318_v92 _dafny.Set
				_ = _318_v92
				_318_v92 = _dafny.SetOf(_215_v4, (_312_v86).Dtor_cf0(), !(_317_v91).Equals(_317_v91), Companion_Default___.Fm3((_313_v87).F9(), _219_globalState), _dafny.Companion_Sequence_.Contains(_311_v85, (_310_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))).Int()).(bool)))
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))
				_ = _index47
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_316_v90), 0))
				_ = _index48
				var _rhs33 bool = _215_v4
				_ = _rhs33
				var _rhs34 _dafny.Int = (_313_v87).F9()
				_ = _rhs34
				var _rhs35 _dafny.Int = (_318_v92).Cardinality()
				_ = _rhs35
				var _lhs25 _dafny.Array = _310_v84
				_ = _lhs25
				var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))
				_ = _lhs26
				var _lhs27 *GlobalState = _219_globalState
				_ = _lhs27
				var _lhs28 _dafny.Array = _316_v90
				_ = _lhs28
				var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_316_v90), 0))
				_ = _lhs29
				(_lhs25).ArraySet1(_rhs33, (_lhs26).Int())
				_lhs27.F6 = _rhs34
				(_lhs28).ArraySet1(_rhs35, (_lhs29).Int())
				var _319_v93 _dafny.Array
				_ = _319_v93
				var _nw58 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
				_ = _nw58
				_319_v93 = _nw58
				var _nw59 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
				_ = _nw59
				_319_v93 = _nw59
				var _320_v94 _dafny.MultiSet
				_ = _320_v94
				_320_v94 = _dafny.MultiSetOf(!(false), !(!(!(_215_v4))), _215_v4, !(true), (_310_v84).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_310_v84), 0))).Int()).(bool))
				var _321_v95 _dafny.Sequence
				_ = _321_v95
				var _out6 _dafny.Sequence
				_ = _out6
				_out6 = Companion_Default___.M0((_dafny.Zero).Minus(_212_v1), _320_v94, _311_v85, _219_globalState)
				_321_v95 = _out6
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_215_v4 = _215_v4
	var _326_v96 D0
	_ = _326_v96
	_326_v96 = Companion_D0_.Create_DC0_(Companion_Default___.Fm3(_212_v1, _219_globalState))
	if (_326_v96).Dtor_cf0() {
		var _327_v97 _dafny.Array
		_ = _327_v97
		var _nw60 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw60
		_327_v97 = _nw60
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
		_ = _index49
		(_327_v97).ArraySet1(_215_v4, (_index49).Int())
		var _328_v98 _dafny.Map
		_ = _328_v98
		_328_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, _212_v1)
		var _329_v99 _dafny.Array
		_ = _329_v99
		var _nwElement0_8 _dafny.Int = (_dafny.Zero).Minus(_212_v1)
		_ = _nwElement0_8
		var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(4))
		_ = _nw61
		(_nw61).ArraySet1(_nwElement0_8, 0)
		(_nw61).ArraySet1(_212_v1, 1)
		(_nw61).ArraySet1((func() _dafny.Int {
			if (_328_v98).Contains(_215_v4) {
				return (_328_v98).Get(_215_v4).(_dafny.Int)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v97, _212_v1)).Cardinality()
		})(), 2)
		(_nw61).ArraySet1(_212_v1, 3)
		_329_v99 = _nw61
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
		_ = _index50
		(_329_v99).ArraySet1(_dafny.IntOfInt64(736), (_index50).Int())
		var _330_v100 _dafny.Sequence
		_ = _330_v100
		_330_v100 = _dafny.SeqOf(_215_v4, _215_v4)
		var _331_v101 _dafny.MultiSet
		_ = _331_v101
		_331_v101 = _dafny.MultiSetOf(_330_v100)
		var _332_v102 D0
		_ = _332_v102
		_332_v102 = Companion_D0_.Create_DC1_(_331_v101, _212_v1)
		var _333_v103 _dafny.Map
		_ = _333_v103
		_333_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_307_v82, _212_v1)
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
		_ = _index51
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
		_ = _index52
		var _rhs36 _dafny.Int = (_332_v102).Dtor_cf2()
		_ = _rhs36
		var _rhs37 bool = !(_333_v103).Contains(Companion_Default___.Fm7(_219_globalState))
		_ = _rhs37
		var _rhs38 _dafny.Int = _212_v1
		_ = _rhs38
		var _lhs30 *GlobalState = _219_globalState
		_ = _lhs30
		var _lhs31 _dafny.Array = _327_v97
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
		_ = _lhs32
		var _lhs33 _dafny.Array = _329_v99
		_ = _lhs33
		var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
		_ = _lhs34
		_lhs30.F6 = _rhs36
		(_lhs31).ArraySet1(_rhs37, (_lhs32).Int())
		(_lhs33).ArraySet1(_rhs38, (_lhs34).Int())
		var _334_v104 _dafny.Array
		_ = _334_v104
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw62
		_334_v104 = _nw62
		var _335_v105 _dafny.Array
		_ = _335_v105
		var _nwElement0_9 _dafny.Array = _334_v104
		_ = _nwElement0_9
		var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(4))
		_ = _nw63
		(_nw63).ArraySet1(_nwElement0_9, 0)
		(_nw63).ArraySet1(_334_v104, 1)
		(_nw63).ArraySet1(_334_v104, 2)
		(_nw63).ArraySet1(_334_v104, 3)
		_335_v105 = _nw63
		var _336_v106 _dafny.Map
		_ = _336_v106
		_336_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int), _334_v104)
		var _337_v107 _dafny.MultiSet
		_ = _337_v107
		_337_v107 = _dafny.MultiSetOf(false, (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool))
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_335_v105), 0))
		_ = _index53
		(_335_v105).ArraySet1((func() _dafny.Array {
			if (_336_v106).Contains((_337_v107).Cardinality()) {
				return (_336_v106).Get((_337_v107).Cardinality()).(_dafny.Array)
			}
			return _334_v104
		})(), (_index53).Int())
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_335_v105), 0))
		_ = _index54
		var _rhs39 bool = (_330_v100).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(268), _dafny.IntOfUint32((_330_v100).Cardinality()))).Uint32()).(bool)
		_ = _rhs39
		var _rhs40 _dafny.Array = _334_v104
		_ = _rhs40
		var _lhs35 _dafny.Array = _335_v105
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_335_v105), 0))
		_ = _lhs36
		_215_v4 = _rhs39
		(_lhs35).ArraySet1(_rhs40, (_lhs36).Int())
		var _338_v108 *C0
		_ = _338_v108
		var _nw64 *C0 = New_C0_()
		_ = _nw64
		_nw64.Ctor__((_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int), _212_v1, _215_v4)
		_338_v108 = _nw64
		var _339_v109 D1
		_ = _339_v109
		_339_v109 = Companion_D1_.Create_DC4_()
		var _340_v110 _dafny.Set
		_ = _340_v110
		_340_v110 = _dafny.SetOf(_339_v109)
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
		_ = _index55
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
		_ = _index56
		var _rhs41 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_338_v108).F9(), (_340_v110).Cardinality()), (_338_v108).F9())
		_ = _rhs41
		var _rhs42 bool = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_306_v81, _307_v82))).IsDisjointFrom(_217_v6)
		_ = _rhs42
		var _lhs37 _dafny.Array = _329_v99
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
		_ = _lhs38
		var _lhs39 _dafny.Array = _327_v97
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
		_ = _lhs40
		(_lhs37).ArraySet1(_rhs41, (_lhs38).Int())
		(_lhs39).ArraySet1(_rhs42, (_lhs40).Int())
		if (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool) {
			var _341_v111 _dafny.Map
			_ = _341_v111
			_341_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool))
			var _342_v112 _dafny.Map
			_ = _342_v112
			_342_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v111, _213_v2)
			_342_v112 = (_342_v112).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool), _215_v4), _213_v2)
			var _343_v113 T0
			_ = _343_v113
			var _nw65 *C0 = New_C0_()
			_ = _nw65
			_nw65.Ctor__((_338_v108).F9(), (_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int), _215_v4)
			_343_v113 = _nw65
			var _344_v114 _dafny.Map
			_ = _344_v114
			_344_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_343_v113, (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool))
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _index57
			(_327_v97).ArraySet1((func() bool {
				if (_344_v114).Contains(_343_v113) {
					return (_344_v114).Get(_343_v113).(bool)
				}
				return true
			})(), (_index57).Int())
			var _345_v115 _dafny.Sequence
			_ = _345_v115
			var _out7 _dafny.Sequence
			_ = _out7
			_out7 = Companion_Default___.M0(_dafny.IntOfInt64(907), _dafny.MultiSetOf(_215_v4, (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool)), _330_v100, _219_globalState)
			_345_v115 = _out7
			var _rhs43 _dafny.Int = (_338_v108).F8()
			_ = _rhs43
			var _rhs44 _dafny.Array = _329_v99
			_ = _rhs44
			var _rhs45 _dafny.Map = _341_v111
			_ = _rhs45
			var _rhs46 _dafny.Int = (_338_v108).F9()
			_ = _rhs46
			var _rhs47 *C0 = _338_v108
			_ = _rhs47
			var _lhs41 *GlobalState = _219_globalState
			_ = _lhs41
			var _lhs42 *GlobalState = _219_globalState
			_ = _lhs42
			_lhs41.F6 = _rhs43
			_329_v99 = _rhs44
			_341_v111 = _rhs45
			_lhs42.F3 = _rhs46
			_338_v108 = _rhs47
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
			_ = _index58
			(_329_v99).ArraySet1((_338_v108).F8(), (_index58).Int())
		} else {
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _index59
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
			_ = _index60
			var _rhs48 bool = _215_v4
			_ = _rhs48
			var _rhs49 _dafny.Int = (_338_v108).F9()
			_ = _rhs49
			var _lhs43 _dafny.Array = _327_v97
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _lhs44
			var _lhs45 _dafny.Array = _329_v99
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))
			_ = _lhs46
			(_lhs43).ArraySet1(_rhs48, (_lhs44).Int())
			(_lhs45).ArraySet1(_rhs49, (_lhs46).Int())
			var _346_v116 _dafny.Map
			_ = _346_v116
			_346_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool), _215_v4)
			(_219_globalState).F3 = (Companion_Default___.SafeModuloInt((_346_v116).Cardinality(), (_307_v82).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_215_v4, _215_v4)).Cardinality(), _dafny.IntOfUint32((_307_v82).Cardinality()))).Uint32()).(_dafny.Int))).Times(Companion_Default___.Fm0((_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int), (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool), _dafny.SeqOf((_338_v108).F9(), (_338_v108).F8(), (_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int)), _213_v2, _219_globalState))
			var _347_v117 _dafny.Sequence
			_ = _347_v117
			_347_v117 = _dafny.SeqOf(_211_v0)
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _index61
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _index62
			var _rhs50 bool = ((_338_v108).F8()).Cmp((_338_v108).Fm5(Companion_Default___.Fm8((_338_v108).F9(), (_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(356), (_329_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_329_v99), 0))).Int()).(_dafny.Int), _219_globalState), _213_v2, _212_v1, _215_v4, _219_globalState)) < 0
			_ = _rhs50
			var _rhs51 D1 = (func() D1 {
				if !_dafny.Companion_Sequence_.Equal((_347_v117).Select((Companion_Default___.SafeIndex((_338_v108).F9(), _dafny.IntOfUint32((_347_v117).Cardinality()))).Uint32()).(_dafny.Sequence), _211_v0) {
					return (func() D1 {
						if (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool) {
							return _308_v83
						}
						return _308_v83
					})()
				}
				return Companion_Default___.Fm9(_219_globalState)
			})()
			_ = _rhs51
			var _rhs52 bool = !((_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool)) || (_215_v4)
			_ = _rhs52
			var _lhs47 _dafny.Array = _327_v97
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _lhs48
			var _lhs49 _dafny.Array = _327_v97
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))
			_ = _lhs50
			(_lhs47).ArraySet1(_rhs50, (_lhs48).Int())
			_308_v83 = _rhs51
			(_lhs49).ArraySet1(_rhs52, (_lhs50).Int())
			_346_v116 = (_346_v116).Update((_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool), (_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool))
			var _348_v118 _dafny.Sequence
			_ = _348_v118
			var _out8 _dafny.Sequence
			_ = _out8
			_out8 = Companion_Default___.M0((_338_v108).F8(), _dafny.MultiSetOf((_327_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(29), _dafny.ArrayLen((_327_v97), 0))).Int()).(bool)), _330_v100, _219_globalState)
			_348_v118 = _out8
		}
	} else {
		var _349_v119 _dafny.Sequence
		_ = _349_v119
		_349_v119 = _dafny.SeqOf(_215_v4)
		var _350_v120 _dafny.MultiSet
		_ = _350_v120
		_350_v120 = _dafny.MultiSetOf(_349_v119, _349_v119)
		var _351_v121 D0
		_ = _351_v121
		_351_v121 = Companion_D0_.Create_DC1_(_350_v120, _212_v1)
		var _352_v122 _dafny.MultiSet
		_ = _352_v122
		_352_v122 = _dafny.MultiSetOf(Companion_D0_.Create_DC1_(_350_v120, _dafny.IntOfUint32((_349_v119).Cardinality())), _351_v121)
		_215_v4 = !((_dafny.MultiSetFromSeq(_dafny.SeqOf(_351_v121))).IsProperSubsetOf(_352_v122))
		var _353_v123 _dafny.MultiSet
		_ = _353_v123
		_353_v123 = _dafny.MultiSetOf(_215_v4, _215_v4)
		var _354_v124 _dafny.Sequence
		_ = _354_v124
		var _out9 _dafny.Sequence
		_ = _out9
		_out9 = Companion_Default___.M0(_dafny.IntOfUint32((_211_v0).Cardinality()), _353_v123, _dafny.Companion_Sequence_.Update(_349_v119, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_215_v4)).Cardinality()), _dafny.IntOfUint32((_349_v119).Cardinality()))).Uint32(), _215_v4), _219_globalState)
		_354_v124 = _out9
		var _355_v125 _dafny.Map
		_ = _355_v125
		_355_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v1, false)
		_215_v4 = !((Companion_Default___.Fm3(Companion_Default___.Fm0(_212_v1, _215_v4, _306_v81, _213_v2, _219_globalState), _219_globalState)) && ((func() bool {
			if (_355_v125).Contains(_212_v1) {
				return (_355_v125).Get(_212_v1).(bool)
			}
			return _215_v4
		})()))
		var _356_v126 _dafny.Sequence
		_ = _356_v126
		var _out10 _dafny.Sequence
		_ = _out10
		_out10 = Companion_Default___.M0(_212_v1, (_353_v123).Difference(_353_v123), Companion_Default___.Fm6(_212_v1, _213_v2, _215_v4, _215_v4, _219_globalState), _219_globalState)
		_356_v126 = _out10
		var _357_v127 _dafny.Map
		_ = _357_v127
		_357_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, (_dafny.Zero).Minus(_212_v1))
		var _358_v128 _dafny.Set
		_ = _358_v128
		_358_v128 = _dafny.SetOf((func() _dafny.Int {
			if (_357_v127).Contains(!(_215_v4)) {
				return (_357_v127).Get(!(_215_v4)).(_dafny.Int)
			}
			return _212_v1
		})())
		_358_v128 = _358_v128
	}
	var _359_v129 _dafny.MultiSet
	_ = _359_v129
	_359_v129 = _dafny.MultiSetOf(_215_v4, _215_v4, _215_v4, _215_v4, Companion_Default___.Fm3(_212_v1, _219_globalState))
	var _360_v130 _dafny.Map
	_ = _360_v130
	_360_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v2, _359_v129)
	_360_v130 = (_360_v130).Update(_213_v2, _359_v129)
	_215_v4 = _215_v4
	_215_v4 = (_326_v96).Dtor_cf0()
	var _361_v132 *C0
	_ = _361_v132
	var _nw66 *C0 = New_C0_()
	_ = _nw66
	_nw66.Ctor__(_212_v1, (func() _dafny.Set {
		var _coll11 = _dafny.NewBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(416), _dafny.IntOfInt64(754))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _362_v131 _dafny.Int
			_362_v131 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(416)).Cmp(_362_v131) <= 0) && ((_362_v131).Cmp(_dafny.IntOfInt64(754)) < 0) {
				_coll11.Add((_362_v131).Times(_212_v1))
			}
		}
		return _coll11.ToSet()
	}()).Cardinality(), _215_v4)
	_361_v132 = _nw66
	var _363_v133 D2
	_ = _363_v133
	_363_v133 = Companion_D2_.Create_DC6_(_361_v132)
	var _364_v134 _dafny.Array
	_ = _364_v134
	var _nwElement0_10 *C0 = _361_v132
	_ = _nwElement0_10
	var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(29))
	_ = _nw67
	(_nw67).ArraySet1(_nwElement0_10, 0)
	(_nw67).ArraySet1(_361_v132, 1)
	(_nw67).ArraySet1((_363_v133).Dtor_cf6(), 2)
	(_nw67).ArraySet1(_361_v132, 3)
	(_nw67).ArraySet1(_361_v132, 4)
	(_nw67).ArraySet1(_361_v132, 5)
	(_nw67).ArraySet1((func() *C0 {
		if false {
			return _361_v132
		}
		return _361_v132
	})(), 6)
	(_nw67).ArraySet1(_361_v132, 7)
	(_nw67).ArraySet1(_361_v132, 8)
	(_nw67).ArraySet1(_361_v132, 9)
	(_nw67).ArraySet1(_361_v132, 10)
	(_nw67).ArraySet1(_361_v132, 11)
	(_nw67).ArraySet1(_361_v132, 12)
	(_nw67).ArraySet1(_361_v132, 13)
	(_nw67).ArraySet1(_361_v132, 14)
	(_nw67).ArraySet1(_361_v132, 15)
	(_nw67).ArraySet1(_361_v132, 16)
	(_nw67).ArraySet1(_361_v132, 17)
	(_nw67).ArraySet1(_361_v132, 18)
	(_nw67).ArraySet1(_361_v132, 19)
	(_nw67).ArraySet1(_361_v132, 20)
	(_nw67).ArraySet1(_361_v132, 21)
	(_nw67).ArraySet1(_361_v132, 22)
	(_nw67).ArraySet1(_361_v132, 23)
	(_nw67).ArraySet1(_361_v132, 24)
	(_nw67).ArraySet1((func() *C0 {
		if _215_v4 {
			return _361_v132
		}
		return _361_v132
	})(), 25)
	(_nw67).ArraySet1((func() *C0 {
		if _215_v4 {
			return _361_v132
		}
		return _361_v132
	})(), 26)
	(_nw67).ArraySet1(_361_v132, 27)
	(_nw67).ArraySet1(_361_v132, 28)
	_364_v134 = _nw67
	var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_364_v134), 0))
	_ = _index63
	(_364_v134).ArraySet1(_361_v132, (_index63).Int())
	var _365_v135 _dafny.Array
	_ = _365_v135
	var _len0_9 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_9
	var _nw68 _dafny.Array
	_ = _nw68
	if _len0_9.Cmp(_dafny.Zero) == 0 {
		_nw68 = _dafny.NewArray(_len0_9)
	} else {
		var _init9 func(_dafny.Int) D0 = (func(_366_v4 bool, _367_v2 _dafny.CodePoint, _368_v132 *C0) func(_dafny.Int) D0 {
			return func(_369_i8 _dafny.Int) D0 {
				return Companion_D0_.Create_DC1_(_dafny.MultiSetOf(_dafny.SeqOf(_366_v4, _366_v4), (Companion_D3_.Create_DC9_(_dafny.SeqOf(_366_v4, _366_v4))).Dtor_cf10(), _dafny.SeqOf((Companion_D3_.Create_DC10_(_367_v2, _366_v4)).Dtor_cf12())), (_368_v132).F9())
			}
		})(_215_v4, _213_v2, _361_v132)
		_ = _init9
		var _element0_9 = _init9(_dafny.Zero)
		_ = _element0_9
		_nw68 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
		(_nw68).ArraySet1(_element0_9, 0)
		var _nativeLen0_9 = (_len0_9).Int()
		_ = _nativeLen0_9
		for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
			(_nw68).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
		}
	}
	_365_v135 = _nw68
	var _370_v136 _dafny.Sequence
	_ = _370_v136
	_370_v136 = _dafny.SeqOf(_215_v4)
	var _371_v137 _dafny.MultiSet
	_ = _371_v137
	_371_v137 = _dafny.MultiSetOf(_370_v136)
	var _372_v138 D0
	_ = _372_v138
	_372_v138 = Companion_D0_.Create_DC1_(_371_v137, (_361_v132).F9())
	var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_365_v135), 0))
	_ = _index64
	(_365_v135).ArraySet1(_372_v138, (_index64).Int())
	var _373_v139 _dafny.Array
	_ = _373_v139
	var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
	_ = _nw69
	_373_v139 = _nw69
	var _374_v140 _dafny.Sequence
	_ = _374_v140
	_374_v140 = _dafny.SeqOf(_373_v139)
	var _375_v141 _dafny.MultiSet
	_ = _375_v141
	_375_v141 = _dafny.MultiSetOf(_306_v81)
	var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_364_v134), 0))
	_ = _index65
	var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_365_v135), 0))
	_ = _index66
	var _rhs53 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_374_v140, _dafny.SeqOf(_373_v139))).Cardinality())
	_ = _rhs53
	var _rhs54 *C0 = _361_v132
	_ = _rhs54
	var _rhs55 D0 = _372_v138
	_ = _rhs55
	var _rhs56 _dafny.Sequence = Companion_Default___.Fm10(_dafny.IntOfUint32((_370_v136).Cardinality()), !(((_361_v132).F8()).Cmp((_361_v132).F8()) <= 0), _212_v1, (_375_v141).Difference(_375_v141), _219_globalState)
	_ = _rhs56
	var _lhs51 *GlobalState = _219_globalState
	_ = _lhs51
	var _lhs52 _dafny.Array = _364_v134
	_ = _lhs52
	var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_364_v134), 0))
	_ = _lhs53
	var _lhs54 _dafny.Array = _365_v135
	_ = _lhs54
	var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.ArrayLen((_365_v135), 0))
	_ = _lhs55
	_lhs51.F6 = _rhs53
	(_lhs52).ArraySet1(_rhs54, (_lhs53).Int())
	(_lhs54).ArraySet1(_rhs55, (_lhs55).Int())
	_211_v0 = _rhs56
	var _376_v142 D3
	_ = _376_v142
	_376_v142 = Companion_D3_.Create_DC10_(_213_v2, _215_v4)
	var _source3 D3 = _376_v142
	_ = _source3
	if _source3.Is_DC10() {
		var _377___mcc_h2 _dafny.CodePoint = _source3.Get_().(D3_DC10).Cf11
		_ = _377___mcc_h2
		var _378___mcc_h3 bool = _source3.Get_().(D3_DC10).Cf12
		_ = _378___mcc_h3
		var _379_cf12 bool = _378___mcc_h3
		_ = _379_cf12
		var _380_cf11 _dafny.CodePoint = _377___mcc_h2
		_ = _380_cf11
		var _381_v143 _dafny.MultiSet
		_ = _381_v143
		_381_v143 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, (_361_v132).F8()))
		_380_cf11 = Companion_Default___.Fm11(_381_v143, ((_217_v6).Cardinality()).Plus(_212_v1), _211_v0, _211_v0, _219_globalState)
		_215_v4 = _215_v4
		var _382_v144 _dafny.Map
		_ = _382_v144
		_382_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, _212_v1)
		var _383_v145 _dafny.Map
		_ = _383_v145
		_383_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_373_v139, _382_v144)
		var _384_v146 _dafny.Array
		_ = _384_v146
		var _nwElement0_11 _dafny.CodePoint = _213_v2
		_ = _nwElement0_11
		var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(8))
		_ = _nw70
		(_nw70).ArraySet1CodePoint(_nwElement0_11, 0)
		(_nw70).ArraySet1CodePoint(Companion_Default___.Fm11(_dafny.MultiSetOf(_382_v144, (func() _dafny.Map {
			if (_383_v145).Contains(_373_v139) {
				return (_383_v145).Get(_373_v139).(_dafny.Map)
			}
			return _382_v144
		})()), _212_v1, _dafny.UnicodeSeqOfUtf8Bytes("unvm"), _211_v0, _219_globalState), 1)
		(_nw70).ArraySet1CodePoint(_dafny.CodePoint('l'), 2)
		(_nw70).ArraySet1CodePoint(_380_cf11, 3)
		(_nw70).ArraySet1CodePoint(_dafny.CodePoint('a'), 4)
		(_nw70).ArraySet1CodePoint(_380_cf11, 5)
		(_nw70).ArraySet1CodePoint(_380_cf11, 6)
		(_nw70).ArraySet1CodePoint(_380_cf11, 7)
		_384_v146 = _nw70
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_384_v146), 0))
		_ = _index67
		(_384_v146).ArraySet1CodePoint(_213_v2, (_index67).Int())
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_384_v146), 0))
		_ = _index68
		(_384_v146).ArraySet1CodePoint(_380_cf11, (_index68).Int())
		var _source4 D0 = _326_v96
		_ = _source4
		if _source4.Is_DC1() {
			var _385___mcc_h9 _dafny.MultiSet = _source4.Get_().(D0_DC1).Cf1
			_ = _385___mcc_h9
			var _386___mcc_h10 _dafny.Int = _source4.Get_().(D0_DC1).Cf2
			_ = _386___mcc_h10
			var _387_cf2 _dafny.Int = _386___mcc_h10
			_ = _387_cf2
			var _388_cf1 _dafny.MultiSet = _385___mcc_h9
			_ = _388_cf1
			(_219_globalState).F3 = (_361_v132).F9()
			_387_cf2 = ((_361_v132).F9()).Plus((_dafny.Zero).Minus(_212_v1))
			var _389_v147 _dafny.Map
			_ = _389_v147
			_389_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v40, _306_v81)
			var _390_v148 _dafny.Set
			_ = _390_v148
			_390_v148 = _dafny.SetOf(_dafny.IntOfUint32((_307_v82).Cardinality()))
			_389_v147 = (_389_v147).Update((_257_v40).Update((_361_v132).F9(), (_390_v148).Cardinality()), _307_v82)
			(_219_globalState).F6 = (_361_v132).F8()
		} else if _source4.Is_DC0() {
			var _391___mcc_h11 bool = _source4.Get_().(D0_DC0).Cf0
			_ = _391___mcc_h11
			var _392_cf0 bool = _391___mcc_h11
			_ = _392_cf0
			var _393_v149 _dafny.Sequence
			_ = _393_v149
			var _out11 _dafny.Sequence
			_ = _out11
			_out11 = Companion_Default___.M0(Companion_Default___.SafeModuloInt((_361_v132).F9(), _212_v1), _359_v129, _370_v136, _219_globalState)
			_393_v149 = _out11
			(_219_globalState).F6 = Companion_Default___.SafeModuloInt(_212_v1, (_361_v132).F9())
			var _394_v150 _dafny.Sequence
			_ = _394_v150
			var _out12 _dafny.Sequence
			_ = _out12
			_out12 = Companion_Default___.M0((_361_v132).F9(), _359_v129, _370_v136, _219_globalState)
			_394_v150 = _out12
			(_219_globalState).F6 = (_361_v132).F8()
		} else {
			var _395___mcc_h12 D0 = _source4.Get_().(D0_DC2).Cf3
			_ = _395___mcc_h12
			var _396_cf3 D0 = _395___mcc_h12
			_ = _396_cf3
			_380_cf11 = _380_cf11
			var _397_v151 _dafny.Map
			_ = _397_v151
			_397_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_361_v132).F8(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_373_v139, _212_v1))
			var _398_v152 _dafny.Map
			_ = _398_v152
			_398_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_373_v139, _212_v1)
			_397_v151 = (_397_v151).Update((_361_v132).F9(), _398_v152)
			var _399_v153 _dafny.Sequence
			_ = _399_v153
			var _out13 _dafny.Sequence
			_ = _out13
			_out13 = Companion_Default___.M0((_361_v132).F9(), Companion_Default___.Fm12((_361_v132).F8(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(546))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_400_v146 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_401_i9 _dafny.Int) _dafny.CodePoint {
					return (_400_v146).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_400_v146), 0))).Int())
				}
			})(_384_v146)))).Cardinality()), Companion_Default___.Fm13((_361_v132).F8(), _219_globalState), _219_globalState), _370_v136, _219_globalState)
			_399_v153 = _out13
			(_219_globalState).F3 = _dafny.IntOfUint32((_399_v153).Cardinality())
		}
	} else if _source3.Is_DC11() {
		var _402___mcc_h4 _dafny.Int = _source3.Get_().(D3_DC11).Cf13
		_ = _402___mcc_h4
		var _403___mcc_h5 _dafny.Int = _source3.Get_().(D3_DC11).Cf14
		_ = _403___mcc_h5
		var _404___mcc_h6 bool = _source3.Get_().(D3_DC11).Cf15
		_ = _404___mcc_h6
		var _405___mcc_h7 bool = _source3.Get_().(D3_DC11).Cf16
		_ = _405___mcc_h7
		var _406_cf16 bool = _405___mcc_h7
		_ = _406_cf16
		var _407_cf15 bool = _404___mcc_h6
		_ = _407_cf15
		var _408_cf14 _dafny.Int = _403___mcc_h5
		_ = _408_cf14
		var _409_cf13 _dafny.Int = _402___mcc_h4
		_ = _409_cf13
		var _410_v154 _dafny.Array
		_ = _410_v154
		var _len0_10 _dafny.Int = _dafny.One
		_ = _len0_10
		var _nw71 _dafny.Array
		_ = _nw71
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw71 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) bool = (func(_411_cf16 bool) func(_dafny.Int) bool {
				return func(_412_i10 _dafny.Int) bool {
					return (_411_cf16) && (_411_cf16)
				}
			})(_406_cf16)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw71 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw71).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw71).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_410_v154 = _nw71
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))
		_ = _index69
		(_410_v154).ArraySet1(_215_v4, (_index69).Int())
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))
		_ = _index70
		(_410_v154).ArraySet1(!((_dafny.IntOfInt64(390)).Cmp(((_361_v132).F8()).Times(_212_v1)) >= 0), (_index70).Int())
		var _413_v155 D4
		_ = _413_v155
		_413_v155 = Companion_D4_.Create_DC12_(_373_v139)
		var _414_v156 _dafny.Array
		_ = _414_v156
		var _nwElement0_12 _dafny.Array = _373_v139
		_ = _nwElement0_12
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(12))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_12, 0)
		(_nw72).ArraySet1(_373_v139, 1)
		(_nw72).ArraySet1(_373_v139, 2)
		(_nw72).ArraySet1(_373_v139, 3)
		(_nw72).ArraySet1(_373_v139, 4)
		(_nw72).ArraySet1((_413_v155).Dtor_cf17(), 5)
		(_nw72).ArraySet1(_373_v139, 6)
		(_nw72).ArraySet1((func() _dafny.Array {
			if _215_v4 {
				return _373_v139
			}
			return _373_v139
		})(), 7)
		(_nw72).ArraySet1(_373_v139, 8)
		(_nw72).ArraySet1(_373_v139, 9)
		(_nw72).ArraySet1(_373_v139, 10)
		(_nw72).ArraySet1(_373_v139, 11)
		_414_v156 = _nw72
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_414_v156), 0))
		_ = _index71
		(_414_v156).ArraySet1(_373_v139, (_index71).Int())
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_414_v156), 0))
		_ = _index72
		var _rhs57 _dafny.Array = _373_v139
		_ = _rhs57
		var _rhs58 _dafny.Int = (_361_v132).F8()
		_ = _rhs58
		var _lhs56 _dafny.Array = _414_v156
		_ = _lhs56
		var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_414_v156), 0))
		_ = _lhs57
		(_lhs56).ArraySet1(_rhs57, (_lhs57).Int())
		_409_cf13 = _rhs58
		var _415_v157 _dafny.Map
		_ = _415_v157
		_415_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), Companion_Default___.Fm3(_212_v1, _219_globalState))
		var _416_v158 _dafny.Set
		_ = _416_v158
		_416_v158 = _dafny.SetOf(_257_v40, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_216_v5).Cardinality(), (_415_v157).Cardinality()), _257_v40)
		if !((_416_v158).IsDisjointFrom(_416_v158)) {
			(_219_globalState).F3 = (func() _dafny.Int {
				if _406_cf16 {
					return (_361_v132).F9()
				}
				return (_361_v132).F8()
			})()
			var _417_v159 _dafny.Map
			_ = _417_v159
			_417_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_410_v154, (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool))
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_364_v134), 0))
			_ = _index73
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__(_212_v1, (_417_v159).Cardinality(), _407_cf15)
			(_364_v134).ArraySet1(_nw73, (_index73).Int())
			_211_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_211_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-203))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_418_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uw"), _211_v0))
			var _419_v160 _dafny.Array
			_ = _419_v160
			var _nw74 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
			_ = _nw74
			_419_v160 = _nw74
			var _420_v161 T0
			_ = _420_v161
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__((_361_v132).F8(), _dafny.IntOfInt64(-634), (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool))
			_420_v161 = _nw75
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_419_v160), 0))
			_ = _index74
			(_419_v160).ArraySet1(_420_v161, (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_419_v160), 0))
			_ = _index75
			(_419_v160).ArraySet1(_420_v161, (_index75).Int())
			_211_v0 = _211_v0
		} else {
			var _421_v162 _dafny.Array
			_ = _421_v162
			var _nw76 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.One)
			_ = _nw76
			_421_v162 = _nw76
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_421_v162), 0))
			_ = _index76
			(_421_v162).ArraySet1(_308_v83, (_index76).Int())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_421_v162), 0))
			_ = _index77
			(_421_v162).ArraySet1(Companion_D1_.Create_DC3_(_307_v82), (_index77).Int())
			(_219_globalState).F6 = (((_361_v132).F8()).Plus((_361_v132).F9())).Minus(Companion_Default___.SafeModuloInt((_361_v132).F9(), _409_cf13))
			var _422_v163 _dafny.Map
			_ = _422_v163
			_422_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_361_v132).F9(), (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool))
			_422_v163 = (_422_v163).Update((_361_v132).F9(), _407_cf15)
			_407_cf15 = _407_cf15
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))
			_ = _index78
			(_410_v154).ArraySet1((_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool), (_index78).Int())
		}
		var _423_v164 _dafny.Set
		_ = _423_v164
		_423_v164 = _dafny.SetOf((_361_v132).F9(), (_dafny.Zero).Minus((_361_v132).F8()))
		if (func() _dafny.Set {
			var _coll12 = _dafny.NewBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_423_v164).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _424_v165 _dafny.Int
				_424_v165 = interface{}(_compr_12).(_dafny.Int)
				if (_423_v164).Contains(_424_v165) {
					_coll12.Add((_424_v165).Plus(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())))
				}
			}
			return _coll12.ToSet()
		}()).Equals((_423_v164).Difference(_423_v164)) {
			var _425_v166 *C0
			_ = _425_v166
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(_212_v1, (_361_v132).F8(), (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool))
			_425_v166 = _nw77
			_212_v1 = ((_dafny.IntOfInt64(620)).Times((_361_v132).F8())).Plus((_372_v138).Dtor_cf2())
			var _426_v167 D3
			_ = _426_v167
			_426_v167 = Companion_D3_.Create_DC11_((_415_v157).Cardinality(), _dafny.IntOfInt64(642), Companion_Default___.Fm3((_425_v166).Fm4(_219_globalState), _219_globalState), (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool))
			var _427_v168 _dafny.Sequence
			_ = _427_v168
			_427_v168 = _dafny.SeqOf(_426_v167)
			_427_v168 = _427_v168
			_407_cf15 = (!((func() bool {
				if _406_cf16 {
					return (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool)
				}
				return _407_cf15
			})())) && (true)
			var _428_v169 _dafny.Map
			_ = _428_v169
			_428_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_425_v166).F9())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))
			_ = _index79
			(_410_v154).ArraySet1((func() bool {
				if (_370_v136).Select((Companion_Default___.SafeIndex(_408_cf14, _dafny.IntOfUint32((_370_v136).Cardinality()))).Uint32()).(bool) {
					return ((_428_v169).Cardinality()).Cmp((_425_v166).F9()) == 0
				}
				return _407_cf15
			})(), (_index79).Int())
		} else {
			var _429_v170 _dafny.Map
			_ = _429_v170
			_429_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_cf13, _dafny.Companion_Sequence_.Update(_211_v0, (Companion_Default___.SafeIndex(_408_cf14, _dafny.IntOfUint32((_211_v0).Cardinality()))).Uint32(), _213_v2))
			var _430_v171 _dafny.Sequence
			_ = _430_v171
			var _out14 _dafny.Sequence
			_ = _out14
			_out14 = Companion_Default___.M0((_361_v132).F9(), _359_v129, Companion_Default___.Fm6(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_429_v170).Contains((_361_v132).F8()) {
					return (_429_v170).Get((_361_v132).F8()).(_dafny.Sequence)
				}
				return _211_v0
			})()).Cardinality()), _213_v2, !(false), Companion_Default___.Fm3(_409_cf13, _219_globalState), _219_globalState), _219_globalState)
			_430_v171 = _out14
			var _rhs59 _dafny.MultiSet = _359_v129
			_ = _rhs59
			var _rhs60 _dafny.Int = (_212_v1).Plus(_dafny.IntOfUint32((_dafny.SeqOf(true, _215_v4, (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool))).Cardinality()))
			_ = _rhs60
			var _lhs58 *GlobalState = _219_globalState
			_ = _lhs58
			_359_v129 = _rhs59
			_lhs58.F3 = _rhs60
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))
			_ = _index80
			(_410_v154).ArraySet1(_407_cf15, (_index80).Int())
			var _431_v172 _dafny.Array
			_ = _431_v172
			var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw78
			_431_v172 = _nw78
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_431_v172), 0))
			_ = _index81
			(_431_v172).ArraySet1(_410_v154, (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_431_v172), 0))
			_ = _index82
			var _rhs61 _dafny.Int = _408_cf14
			_ = _rhs61
			var _rhs62 _dafny.Array = _410_v154
			_ = _rhs62
			var _rhs63 _dafny.MultiSet = _359_v129
			_ = _rhs63
			var _lhs59 *GlobalState = _219_globalState
			_ = _lhs59
			var _lhs60 _dafny.Array = _431_v172
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(306), _dafny.ArrayLen((_431_v172), 0))
			_ = _lhs61
			_lhs59.F3 = _rhs61
			(_lhs60).ArraySet1(_rhs62, (_lhs61).Int())
			_359_v129 = _rhs63
			var _432_v173 _dafny.Sequence
			_ = _432_v173
			var _out15 _dafny.Sequence
			_ = _out15
			_out15 = Companion_Default___.M0((_361_v132).Fm5(_415_v157, _213_v2, _dafny.IntOfInt64(213), (_410_v154).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_410_v154), 0))).Int()).(bool), _219_globalState), _359_v129, _dafny.SeqOf(_215_v4), _219_globalState)
			_432_v173 = _out15
		}
	} else {
		var _433___mcc_h8 _dafny.Sequence = _source3.Get_().(D3_DC9).Cf10
		_ = _433___mcc_h8
		var _434_cf10 _dafny.Sequence = _433___mcc_h8
		_ = _434_cf10
		if !(!(!((_215_v4) == (_215_v4)))) {
			var _435_v174 _dafny.Array
			_ = _435_v174
			var _nw79 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw79
			_435_v174 = _nw79
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_435_v174), 0))
			_ = _index83
			(_435_v174).ArraySet1((!(_215_v4)) == (_215_v4), (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_435_v174), 0))
			_ = _index84
			(_435_v174).ArraySet1(((_257_v40).Cardinality()).Cmp((_361_v132).F8()) != 0, (_index84).Int())
			(_219_globalState).F6 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_211_v0, _211_v0)).Cardinality())
			var _436_v175 _dafny.Array
			_ = _436_v175
			var _nw80 _dafny.Array = _dafny.NewArray(_dafny.One)
			_ = _nw80
			_436_v175 = _nw80
			var _437_v176 T0
			_ = _437_v176
			var _nw81 *C0 = New_C0_()
			_ = _nw81
			_nw81.Ctor__((_dafny.Zero).Minus((_361_v132).F8()), (_361_v132).F8(), _215_v4)
			_437_v176 = _nw81
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_436_v175), 0))
			_ = _index85
			(_436_v175).ArraySet1(_437_v176, (_index85).Int())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_436_v175), 0))
			_ = _index86
			(_436_v175).ArraySet1(_437_v176, (_index86).Int())
			var _438_v177 _dafny.Sequence
			_ = _438_v177
			var _out16 _dafny.Sequence
			_ = _out16
			_out16 = Companion_Default___.M0((_361_v132).F8(), _359_v129, _434_cf10, _219_globalState)
			_438_v177 = _out16
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_435_v174), 0))
			_ = _index87
			(_435_v174).ArraySet1(Companion_Default___.Fm3(((_361_v132).F8()).Times((_361_v132).F9()), _219_globalState), (_index87).Int())
		} else {
			var _439_v178 _dafny.Sequence
			_ = _439_v178
			_439_v178 = _dafny.SeqOf(_361_v132, _361_v132)
			var _440_v179 T0
			_ = _440_v179
			var _nw82 *C0 = New_C0_()
			_ = _nw82
			_nw82.Ctor__(_dafny.IntOfUint32((_211_v0).Cardinality()), _dafny.IntOfUint32((_439_v178).Cardinality()), _215_v4)
			_440_v179 = _nw82
			var _441_v180 _dafny.Sequence
			_ = _441_v180
			_441_v180 = _dafny.SeqOf(_440_v179)
			var _442_v181 _dafny.MultiSet
			_ = _442_v181
			_442_v181 = _dafny.MultiSetOf(_440_v179)
			var _443_v182 _dafny.Array
			_ = _443_v182
			var _nwElement0_13 T0 = _440_v179
			_ = _nwElement0_13
			var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(23))
			_ = _nw83
			(_nw83).ArraySet1(_nwElement0_13, 0)
			(_nw83).ArraySet1(_440_v179, 1)
			(_nw83).ArraySet1(_440_v179, 2)
			(_nw83).ArraySet1(_440_v179, 3)
			(_nw83).ArraySet1((_441_v180).Select((Companion_Default___.SafeIndex(((_442_v181).Update(_440_v179, Companion_Default___.Abs((_361_v132).F8()))).Cardinality(), _dafny.IntOfUint32((_441_v180).Cardinality()))).Uint32()).(T0), 4)
			(_nw83).ArraySet1(_440_v179, 5)
			(_nw83).ArraySet1(_440_v179, 6)
			(_nw83).ArraySet1(_440_v179, 7)
			(_nw83).ArraySet1(_440_v179, 8)
			(_nw83).ArraySet1(_440_v179, 9)
			(_nw83).ArraySet1(_440_v179, 10)
			(_nw83).ArraySet1(_440_v179, 11)
			(_nw83).ArraySet1(_440_v179, 12)
			(_nw83).ArraySet1(_440_v179, 13)
			(_nw83).ArraySet1(_440_v179, 14)
			(_nw83).ArraySet1(_440_v179, 15)
			(_nw83).ArraySet1(_440_v179, 16)
			(_nw83).ArraySet1(_440_v179, 17)
			(_nw83).ArraySet1(_440_v179, 18)
			(_nw83).ArraySet1((func() T0 {
				if true {
					return _440_v179
				}
				return _440_v179
			})(), 19)
			(_nw83).ArraySet1(_440_v179, 20)
			(_nw83).ArraySet1(_440_v179, 21)
			(_nw83).ArraySet1(_440_v179, 22)
			_443_v182 = _nw83
			var _444_v183 T0
			_ = _444_v183
			_444_v183 = _440_v179
			var _445_v184 _dafny.Map
			_ = _445_v184
			_445_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(793), _440_v179)
			var _nwElement0_14 T0 = _440_v179
			_ = _nwElement0_14
			var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(20))
			_ = _nw84
			(_nw84).ArraySet1(_nwElement0_14, 0)
			(_nw84).ArraySet1(_440_v179, 1)
			(_nw84).ArraySet1(_440_v179, 2)
			(_nw84).ArraySet1(_440_v179, 3)
			(_nw84).ArraySet1(_440_v179, 4)
			(_nw84).ArraySet1(_440_v179, 5)
			(_nw84).ArraySet1(_440_v179, 6)
			(_nw84).ArraySet1(_440_v179, 7)
			(_nw84).ArraySet1(_440_v179, 8)
			(_nw84).ArraySet1(_440_v179, 9)
			(_nw84).ArraySet1(_440_v179, 10)
			(_nw84).ArraySet1(_440_v179, 11)
			(_nw84).ArraySet1((_444_v183), 12)
			(_nw84).ArraySet1(_440_v179, 13)
			(_nw84).ArraySet1(_440_v179, 14)
			(_nw84).ArraySet1(_440_v179, 15)
			(_nw84).ArraySet1(_440_v179, 16)
			(_nw84).ArraySet1((func() T0 {
				if (_445_v184).Contains((_361_v132).F8()) {
					return (_445_v184).Get((_361_v132).F8()).(T0)
				}
				return _440_v179
			})(), 17)
			(_nw84).ArraySet1(_440_v179, 18)
			(_nw84).ArraySet1(_440_v179, 19)
			_443_v182 = _nw84
			var _446_v185 _dafny.Sequence
			_ = _446_v185
			var _out17 _dafny.Sequence
			_ = _out17
			_out17 = Companion_Default___.M0(_212_v1, _359_v129, _370_v136, _219_globalState)
			_446_v185 = _out17
			var _447_v186 _dafny.Array
			_ = _447_v186
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_11
			var _nw85 _dafny.Array
			_ = _nw85
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw85 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = func(_448_i12 _dafny.Int) bool {
					return true
				}
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw85 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw85).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw85).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_447_v186 = _nw85
			var _449_v187 _dafny.Map
			_ = _449_v187
			_449_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_361_v132, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(144), _440_v179)).Cardinality())
			var _rhs64 _dafny.Array = _447_v186
			_ = _rhs64
			var _rhs65 bool = _440_v179.F7()
			_ = _rhs65
			var _rhs66 _dafny.Int = (_212_v1).Times(((_361_v132).F9()).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_449_v187)).Cardinality()))))
			_ = _rhs66
			var _rhs67 _dafny.Array = _373_v139
			_ = _rhs67
			var _rhs68 _dafny.Int = ((_361_v132).F8()).Times((_dafny.Zero).Minus((_361_v132).F8()))
			_ = _rhs68
			var _lhs62 *GlobalState = _219_globalState
			_ = _lhs62
			_447_v186 = _rhs64
			_215_v4 = _rhs65
			_lhs62.F6 = _rhs66
			_373_v139 = _rhs67
			_212_v1 = _rhs68
			var _450_v188 _dafny.MultiSet
			_ = _450_v188
			_450_v188 = _dafny.MultiSetOf(_447_v186)
			var _451_v189 D3
			_ = _451_v189
			_451_v189 = Companion_D3_.Create_DC11_((_361_v132).F9(), (_440_v179).Fm4(_219_globalState), (_450_v188).IsSubsetOf(_450_v188), true)
			_451_v189 = _451_v189
			_434_cf10 = _dafny.Companion_Sequence_.Concatenate(_434_cf10, _370_v136)
		}
		_215_v4 = !(!(_215_v4)) || (((_361_v132).F9()).Cmp((_361_v132).F9()) == 0)
		var _452_v190 _dafny.Sequence
		_ = _452_v190
		var _out18 _dafny.Sequence
		_ = _out18
		_out18 = Companion_Default___.M0((_361_v132).F9(), _359_v129, _434_cf10, _219_globalState)
		_452_v190 = _out18
		var _rhs69 bool = _215_v4
		_ = _rhs69
		var _rhs70 _dafny.Int = (_361_v132).F8()
		_ = _rhs70
		var _lhs63 *GlobalState = _219_globalState
		_ = _lhs63
		_215_v4 = _rhs69
		_lhs63.F6 = _rhs70
	}
	_215_v4 = _215_v4
	var _rhs71 bool = (_370_v136).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
		if _215_v4 {
			return (_361_v132).F8()
		}
		return (_361_v132).F9()
	})(), _dafny.IntOfUint32((_370_v136).Cardinality()))).Uint32()).(bool)
	_ = _rhs71
	var _rhs72 bool = true
	_ = _rhs72
	var _rhs73 _dafny.Int = (_dafny.Zero).Minus((_212_v1).Plus(_212_v1))
	_ = _rhs73
	var _rhs74 _dafny.Int = (_361_v132).F9()
	_ = _rhs74
	var _lhs64 *GlobalState = _219_globalState
	_ = _lhs64
	var _lhs65 *GlobalState = _219_globalState
	_ = _lhs65
	_215_v4 = _rhs71
	_215_v4 = _rhs72
	_lhs64.F3 = _rhs73
	_lhs65.F3 = _rhs74
	if !(_215_v4) || (true) {
		var _453_v191 _dafny.Sequence
		_ = _453_v191
		var _out19 _dafny.Sequence
		_ = _out19
		_out19 = Companion_Default___.M0(_212_v1, _dafny.MultiSetFromSeq(_370_v136), _dafny.Companion_Sequence_.Concatenate(_370_v136, _dafny.SeqOf(_215_v4, _215_v4, _215_v4)), _219_globalState)
		_453_v191 = _out19
		_215_v4 = true
		_215_v4 = !_dafny.Companion_Sequence_.Equal(_211_v0, _453_v191)
		var _454_v192 _dafny.Array
		_ = _454_v192
		var _nw86 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw86
		_454_v192 = _nw86
		var _455_v193 D1
		_ = _455_v193
		_455_v193 = Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Update(_307_v82, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_454_v192, _212_v1)).Cardinality(), _dafny.IntOfUint32((_307_v82).Cardinality()))).Uint32(), (_361_v132).F8()))
		var _456_v194 D1
		_ = _456_v194
		_456_v194 = Companion_D1_.Create_DC5_(_455_v193)
		var _source5 D1 = _456_v194
		_ = _source5
		if _source5.Is_DC4() {
			var _457_v195 _dafny.Array
			_ = _457_v195
			var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
			_ = _nw87
			_457_v195 = _nw87
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_457_v195), 0))
			_ = _index88
			(_457_v195).ArraySet1(_306_v81, (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_457_v195), 0))
			_ = _index89
			(_457_v195).ArraySet1(_306_v81, (_index89).Int())
			var _458_v196 _dafny.Set
			_ = _458_v196
			_458_v196 = _dafny.SetOf(_dafny.IntOfInt64(-939), (_361_v132).F8())
			var _459_v197 D6
			_ = _459_v197
			_459_v197 = Companion_D6_.Create_DC16_(_458_v196)
			(_219_globalState).F3 = (((_459_v197).Dtor_cf24()).Cardinality()).Plus(_212_v1)
			var _460_v198 _dafny.Map
			_ = _460_v198
			_460_v198 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_361_v132).F8(), false)
			_460_v198 = (_460_v198).Update(_212_v1, false)
			var _461_v199 _dafny.Set
			_ = _461_v199
			_461_v199 = _dafny.SetOf(_215_v4)
			var _462_v200 _dafny.Map
			_ = _462_v200
			_462_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_370_v136, (Companion_Default___.SafeIndex((_361_v132).F8(), _dafny.IntOfUint32((_370_v136).Cardinality()))).Uint32(), _215_v4)).Cardinality()), _dafny.IntOfUint32((_211_v0).Cardinality())), _461_v199)
			_462_v200 = (_462_v200).Update(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_307_v82).Select((Companion_Default___.SafeIndex(_212_v1, _dafny.IntOfUint32((_307_v82).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.IntOfInt64(-792)), (_461_v199).Difference(_461_v199))
		} else if _source5.Is_DC3() {
			var _463___mcc_h13 _dafny.Sequence = _source5.Get_().(D1_DC3).Cf4
			_ = _463___mcc_h13
			var _464_cf4 _dafny.Sequence = _463___mcc_h13
			_ = _464_cf4
			_215_v4 = !(_215_v4)
			var _465_v201 T0
			_ = _465_v201
			var _nw88 *C0 = New_C0_()
			_ = _nw88
			_nw88.Ctor__(_dafny.IntOfInt64(965), (_361_v132).F8(), Companion_Default___.Fm3(_212_v1, _219_globalState))
			_465_v201 = _nw88
			var _466_v202 _dafny.Set
			_ = _466_v202
			_466_v202 = _dafny.SetOf(_465_v201.F7(), _465_v201.F7(), _465_v201.F7(), false, _465_v201.F7())
			_465_v201 = (func() T0 {
				if ((_361_v132).F9()).Cmp(Companion_Default___.Fm0((_466_v202).Cardinality(), true, _dafny.SeqOf((_361_v132).F9(), _212_v1, (_361_v132).F9()), _213_v2, _219_globalState)) != 0 {
					return _465_v201
				}
				return _465_v201
			})()
			var _467_v203 _dafny.Map
			_ = _467_v203
			_467_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_465_v201.F7(), (_dafny.Zero).Minus(_212_v1))
			var _468_v204 *C0
			_ = _468_v204
			var _nw89 *C0 = New_C0_()
			_ = _nw89
			_nw89.Ctor__(((_467_v203).Merge(_467_v203)).Cardinality(), (_361_v132).F8(), false)
			_468_v204 = _nw89
			(_219_globalState).F3 = (((_468_v204).Fm4(_219_globalState)).Minus((_361_v132).F8())).Plus((_361_v132).F9())
		} else {
			var _469___mcc_h14 D1 = _source5.Get_().(D1_DC5).Cf5
			_ = _469___mcc_h14
			var _470_cf5 D1 = _469___mcc_h14
			_ = _470_cf5
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_364_v134), 0))
			_ = _index90
			(_364_v134).ArraySet1((func() *C0 {
				if ((_361_v132).F9()).Cmp((_361_v132).F8()) == 0 {
					return _361_v132
				}
				return _361_v132
			})(), (_index90).Int())
			var _471_v205 _dafny.Map
			_ = _471_v205
			_471_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_361_v132).F9(), _454_v192)
			_471_v205 = (_471_v205).Update((func() _dafny.Int {
				if _215_v4 {
					return (_361_v132).F9()
				}
				return (_361_v132).F9()
			})(), _454_v192)
			_215_v4 = (Companion_Default___.Fm14(_219_globalState)).IsSubsetOf(_254_v37)
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_373_v139), 0))
			_ = _index91
			(_373_v139).ArraySet1((_361_v132).F8(), (_index91).Int())
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_373_v139), 0))
			_ = _index92
			(_373_v139).ArraySet1((_361_v132).F9(), (_index92).Int())
		}
		_215_v4 = _215_v4
	} else {
		(_219_globalState).F3 = _212_v1
		var _472_v206 _dafny.Map
		_ = _472_v206
		_472_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(!(_215_v4)), _373_v139)
		var _473_v207 _dafny.Set
		_ = _473_v207
		_473_v207 = _dafny.SetOf(_215_v4)
		_472_v206 = (_472_v206).Update(_473_v207, _373_v139)
		(_219_globalState).F3 = (_361_v132).F8()
		if _215_v4 {
			var _474_v208 *C0
			_ = _474_v208
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__(((_361_v132).F8()).Times((_361_v132).F9()), _212_v1, _215_v4)
			_474_v208 = _nw90
			var _475_v209 _dafny.MultiSet
			_ = _475_v209
			_475_v209 = _dafny.MultiSetOf(_361_v132, _361_v132)
			var _476_v210 _dafny.Map
			_ = _476_v210
			_476_v210 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_475_v209, (_dafny.SetOf(true, _215_v4, _215_v4)).Cardinality())
			_215_v4 = !(!(_476_v210).Contains((_475_v209).Intersection(_475_v209)))
			(_219_globalState).F6 = _dafny.IntOfInt64(65)
			var _477_v212 _dafny.Set
			_ = _477_v212
			_477_v212 = _dafny.SetOf(_dafny.IntOfUint32((_211_v0).Cardinality()), Companion_Default___.SafeModuloInt((func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate((_306_v81).Elements()); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _478_v211 _dafny.Int
					_478_v211 = interface{}(_compr_13).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_306_v81, _478_v211) {
						_coll13.Add((_478_v211).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.IntOfInt64(931))).Cardinality()))
					}
				}
				return _coll13.ToSet()
			}()).Cardinality(), (_361_v132).F8()))
			_477_v212 = ((_477_v212).Intersection(_477_v212)).Intersection(_477_v212)
			_213_v2 = _213_v2
		} else {
			_211_v0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10((_361_v132).F8(), _215_v4, ((_dafny.MultiSetOf(_215_v4, _215_v4)).Update(Companion_Default___.Fm3((_361_v132).F8(), _219_globalState), Companion_Default___.Abs((_361_v132).F8()))).Cardinality(), _375_v141, _219_globalState), _211_v0)
			_361_v132 = (_364_v134).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_364_v134), 0))).Int()).(*C0)
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_373_v139), 0))
			_ = _index93
			(_373_v139).ArraySet1((_dafny.Zero).Minus((_361_v132).F8()), (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_373_v139), 0))
			_ = _index94
			(_373_v139).ArraySet1((_212_v1).Minus((_361_v132).F9()), (_index94).Int())
			var _479_v213 _dafny.Map
			_ = _479_v213
			_479_v213 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_215_v4), _dafny.SetOf(_215_v4))
			var _480_v214 _dafny.Sequence
			_ = _480_v214
			var _out20 _dafny.Sequence
			_ = _out20
			_out20 = Companion_Default___.M0(((func() _dafny.Set {
				if (_479_v213).Contains(_215_v4) {
					return (_479_v213).Get(_215_v4).(_dafny.Set)
				}
				return _473_v207
			})()).Cardinality(), _359_v129, _370_v136, _219_globalState)
			_480_v214 = _out20
			var _481_v215 *C0
			_ = _481_v215
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__((_361_v132).F8(), Companion_Default___.SafeModuloInt((_361_v132).F9(), (_361_v132).F8()), ((_373_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_373_v139), 0))).Int()).(_dafny.Int)).Cmp((_361_v132).F8()) >= 0)
			_481_v215 = _nw91
		}
		var _482_v216 _dafny.Map
		_ = _482_v216
		_482_v216 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, _215_v4)
		var _483_v217 _dafny.Map
		_ = _483_v217
		_483_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v4, !(!((_370_v136).Select((Companion_Default___.SafeIndex((_482_v216).Cardinality(), _dafny.IntOfUint32((_370_v136).Cardinality()))).Uint32()).(bool))))
		var _484_v218 D3
		_ = _484_v218
		_484_v218 = Companion_D3_.Create_DC9_(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_370_v136, (Companion_Default___.SafeIndex((_361_v132).F9(), _dafny.IntOfUint32((_370_v136).Cardinality()))).Uint32(), false), (Companion_Default___.SafeIndex((_361_v132).Fm5(_482_v216, _213_v2, _dafny.IntOfInt64(472), _215_v4, _219_globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_370_v136, (Companion_Default___.SafeIndex((_361_v132).F9(), _dafny.IntOfUint32((_370_v136).Cardinality()))).Uint32(), false)).Cardinality()))).Uint32(), _215_v4))
		var _485_v219 _dafny.Sequence
		_ = _485_v219
		var _out21 _dafny.Sequence
		_ = _out21
		_out21 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt((_361_v132).F8(), _212_v1), _dafny.MultiSetOf(_215_v4, _215_v4), (_484_v218).Dtor_cf10(), _219_globalState)
		_485_v219 = _out21
	}
	_212_v1 = (_212_v1).Minus((_375_v141).Cardinality())
	_dafny.Print(_211_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_212_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_213_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(779), _dafny.CodePoint('p'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_215_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_216_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_217_v6).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-166))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v7).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_219_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_globalState.F4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_219_globalState).F5()).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_219_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v37).Equals(_dafny.SetOf(_dafny.CodePoint('p'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v38).Equals(_dafny.SetOf(_dafny.SetOf(_dafny.CodePoint('p')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_256_v39, _dafny.SeqOf(_dafny.SetOf(_dafny.SetOf(_dafny.CodePoint('p'))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v40).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(778), _dafny.IntOfInt64(778))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_306_v81, _dafny.SeqOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_307_v82, _dafny.SeqOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_308_v83).Dtor_cf4(), _dafny.SeqOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_309_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_326_v96).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_359_v129).Equals(_dafny.MultiSetOf(true, true, true, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_360_v130).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.MultiSetOf(true, true, true, true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_361_v132).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_361_v132).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_363_v133).Dtor_cf6()).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_363_v133).Dtor_cf6()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_363_v133).Dtor_cf6().F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.Zero).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.Zero).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.Zero).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.One).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.One).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.One).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(*C0)).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_364_v134).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(*C0)).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_364_v134).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(*C0).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_365_v135).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf1()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_365_v135).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_365_v135).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf1()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false, false), _dafny.SeqOf(false, false), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_365_v135).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_365_v135).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf1()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false, false), _dafny.SeqOf(false, false), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_365_v135).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_365_v135).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf1()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false, false), _dafny.SeqOf(false, false), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_365_v135).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_370_v136, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_371_v137).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_372_v138).Dtor_cf1()).Equals(_dafny.MultiSetOf(_dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_372_v138).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_373_v139).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_374_v140).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_375_v141).Equals(_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_376_v142).Dtor_cf11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_376_v142).Dtor_cf12())
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
	Cf2 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.MultiSet, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
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

type D0_DC2 struct {
	Cf3 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf3 D0) D0 {
	return D0{D0_DC2{Cf3}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.EmptyMultiSet, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.MultiSet {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1.Equals(data2.Cf1) && data1.Cf2.Cmp(data2.Cf2) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_() D1 {
	return D1{D1_DC4{}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC5 struct {
	Cf5 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf5 D1) D1 {
	return D1{D1_DC5{Cf5}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_()
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() D1 {
	return _this.Get_().(D1_DC5).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			_, ok := other.Get_().(D1_DC4)
			return ok
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
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

type D2_DC7 struct {
	Cf7 _dafny.Int
	Cf8 *C0
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf7 _dafny.Int, Cf8 *C0) D2 {
	return D2{D2_DC7{Cf7, Cf8}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf6 *C0
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 *C0) D2 {
	return D2{D2_DC6{Cf6}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf9 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf9 D2) D2 {
	return D2{D2_DC8{Cf9}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.Zero, (*C0)(nil))
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) Dtor_cf8() *C0 {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf6() *C0 {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) Dtor_cf9() D2 {
	return _this.Get_().(D2_DC8).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf6 == data2.Cf6
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf11 _dafny.CodePoint
	Cf12 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf11 _dafny.CodePoint, Cf12 bool) D3 {
	return D3{D3_DC10{Cf11, Cf12}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf13 _dafny.Int
	Cf14 _dafny.Int
	Cf15 bool
	Cf16 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf13 _dafny.Int, Cf14 _dafny.Int, Cf15 bool, Cf16 bool) D3 {
	return D3{D3_DC11{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf10 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf10 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf10}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(_dafny.CodePoint('D'), false)
}

func (_this D3) Dtor_cf11() _dafny.CodePoint {
	return _this.Get_().(D3_DC10).Cf11
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf14
}

func (_this D3) Dtor_cf15() bool {
	return _this.Get_().(D3_DC11).Cf15
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC11).Cf16
}

func (_this D3) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15 && data1.Cf16 == data2.Cf16
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
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

type D4_DC13 struct {
	Cf18 _dafny.Int
	Cf19 _dafny.Int
	Cf20 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf18 _dafny.Int, Cf19 _dafny.Int, Cf20 bool) D4 {
	return D4{D4_DC13{Cf18, Cf19, Cf20}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf21 *C0
	Cf22 bool
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf21 *C0, Cf22 bool) D4 {
	return D4{D4_DC14{Cf21, Cf22}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC12 struct {
	Cf17 _dafny.Array
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf17 _dafny.Array) D4 {
	return D4{D4_DC12{Cf17}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) Dtor_cf21() *C0 {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC14).Cf22
}

func (_this D4) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf17) + ")"
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
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf17 == data2.Cf17
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
	Cf23 T0
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf23 T0) D5 {
	return D5{D5_DC15{Cf23}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() T0 {
	return (T0)(nil)
}

func (_this D5) Dtor_cf23() T0 {
	return _this.Get_().(D5_DC15).Cf23
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf23, data2.Cf23)
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
	Cf25 bool
	Cf26 bool
	Cf27 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf25 bool, Cf26 bool, Cf27 bool) D6 {
	return D6{D6_DC17{Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf28 *C0
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf28 *C0) D6 {
	return D6{D6_DC18{Cf28}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC16 struct {
	Cf24 _dafny.Set
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf24 _dafny.Set) D6 {
	return D6{D6_DC16{Cf24}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false, false, false)
}

func (_this D6) Dtor_cf25() bool {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC17).Cf26
}

func (_this D6) Dtor_cf27() bool {
	return _this.Get_().(D6_DC17).Cf27
}

func (_this D6) Dtor_cf28() *C0 {
	return _this.Get_().(D6_DC18).Cf28
}

func (_this D6) Dtor_cf24() _dafny.Set {
	return _this.Get_().(D6_DC16).Cf24
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf28 == data2.Cf28
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf30 bool
	Cf31 bool
	Cf32 *C0
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf30 bool, Cf31 bool, Cf32 *C0) D7 {
	return D7{D7_DC20{Cf30, Cf31, Cf32}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC19 struct {
	Cf29 _dafny.MultiSet
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf29 _dafny.MultiSet) D7 {
	return D7{D7_DC19{Cf29}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC21 struct {
	Cf33 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf33 D7) D7 {
	return D7{D7_DC21{Cf33}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(false, false, (*C0)(nil))
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC20).Cf30
}

func (_this D7) Dtor_cf31() bool {
	return _this.Get_().(D7_DC20).Cf31
}

func (_this D7) Dtor_cf32() *C0 {
	return _this.Get_().(D7_DC20).Cf32
}

func (_this D7) Dtor_cf29() _dafny.MultiSet {
	return _this.Get_().(D7_DC19).Cf29
}

func (_this D7) Dtor_cf33() D7 {
	return _this.Get_().(D7_DC21).Cf33
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf33) + ")"
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
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf33.Equals(data2.Cf33)
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
	Cf35 bool
	Cf36 bool
	Cf37 *C0
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf35 bool, Cf36 bool, Cf37 *C0) D8 {
	return D8{D8_DC23{Cf35, Cf36, Cf37}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC24 struct {
	Cf38 bool
	Cf39 _dafny.Array
	Cf40 bool
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf38 bool, Cf39 _dafny.Array, Cf40 bool) D8 {
	return D8{D8_DC24{Cf38, Cf39, Cf40}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC25 struct {
	Cf41 D6
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf41 D6) D8 {
	return D8{D8_DC25{Cf41}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC22 struct {
	Cf34 _dafny.Set
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf34 _dafny.Set) D8 {
	return D8{D8_DC22{Cf34}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC26 struct {
	Cf42 D8
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf42 D8) D8 {
	return D8{D8_DC26{Cf42}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(false, false, (*C0)(nil))
}

func (_this D8) Dtor_cf35() bool {
	return _this.Get_().(D8_DC23).Cf35
}

func (_this D8) Dtor_cf36() bool {
	return _this.Get_().(D8_DC23).Cf36
}

func (_this D8) Dtor_cf37() *C0 {
	return _this.Get_().(D8_DC23).Cf37
}

func (_this D8) Dtor_cf38() bool {
	return _this.Get_().(D8_DC24).Cf38
}

func (_this D8) Dtor_cf39() _dafny.Array {
	return _this.Get_().(D8_DC24).Cf39
}

func (_this D8) Dtor_cf40() bool {
	return _this.Get_().(D8_DC24).Cf40
}

func (_this D8) Dtor_cf41() D6 {
	return _this.Get_().(D8_DC25).Cf41
}

func (_this D8) Dtor_cf34() _dafny.Set {
	return _this.Get_().(D8_DC22).Cf34
}

func (_this D8) Dtor_cf42() D8 {
	return _this.Get_().(D8_DC26).Cf42
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf41.Equals(data2.Cf41)
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

type D9_DC28 struct {
	Cf44 bool
	Cf45 _dafny.Int
	Cf46 _dafny.Array
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf44 bool, Cf45 _dafny.Int, Cf46 _dafny.Array) D9 {
	return D9{D9_DC28{Cf44, Cf45, Cf46}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

type D9_DC29 struct {
	Cf47 _dafny.Int
	Cf48 bool
	Cf49 bool
}

func (D9_DC29) isD9() {}

func (CompanionStruct_D9_) Create_DC29_(Cf47 _dafny.Int, Cf48 bool, Cf49 bool) D9 {
	return D9{D9_DC29{Cf47, Cf48, Cf49}}
}

func (_this D9) Is_DC29() bool {
	_, ok := _this.Get_().(D9_DC29)
	return ok
}

type D9_DC30 struct {
	Cf50 _dafny.Int
	Cf51 _dafny.Int
	Cf52 _dafny.Int
	Cf53 _dafny.Array
}

func (D9_DC30) isD9() {}

func (CompanionStruct_D9_) Create_DC30_(Cf50 _dafny.Int, Cf51 _dafny.Int, Cf52 _dafny.Int, Cf53 _dafny.Array) D9 {
	return D9{D9_DC30{Cf50, Cf51, Cf52, Cf53}}
}

func (_this D9) Is_DC30() bool {
	_, ok := _this.Get_().(D9_DC30)
	return ok
}

type D9_DC27 struct {
	Cf43 _dafny.MultiSet
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf43 _dafny.MultiSet) D9 {
	return D9{D9_DC27{Cf43}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC28_(false, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D9) Dtor_cf44() bool {
	return _this.Get_().(D9_DC28).Cf44
}

func (_this D9) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D9_DC28).Cf45
}

func (_this D9) Dtor_cf46() _dafny.Array {
	return _this.Get_().(D9_DC28).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC29).Cf47
}

func (_this D9) Dtor_cf48() bool {
	return _this.Get_().(D9_DC29).Cf48
}

func (_this D9) Dtor_cf49() bool {
	return _this.Get_().(D9_DC29).Cf49
}

func (_this D9) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D9_DC30).Cf50
}

func (_this D9) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D9_DC30).Cf51
}

func (_this D9) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D9_DC30).Cf52
}

func (_this D9) Dtor_cf53() _dafny.Array {
	return _this.Get_().(D9_DC30).Cf53
}

func (_this D9) Dtor_cf43() _dafny.MultiSet {
	return _this.Get_().(D9_DC27).Cf43
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D9_DC29:
		{
			return "D9.DC29" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ")"
		}
	case D9_DC30:
		{
			return "D9.DC30" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf44 == data2.Cf44 && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46
		}
	case D9_DC29:
		{
			data2, ok := other.Get_().(D9_DC29)
			return ok && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48 == data2.Cf48 && data1.Cf49 == data2.Cf49
		}
	case D9_DC30:
		{
			data2, ok := other.Get_().(D9_DC30)
			return ok && data1.Cf50.Cmp(data2.Cf50) == 0 && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53 == data2.Cf53
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D10_DC31 struct {
	Cf54 _dafny.Sequence
}

func (D10_DC31) isD10() {}

func (CompanionStruct_D10_) Create_DC31_(Cf54 _dafny.Sequence) D10 {
	return D10{D10_DC31{Cf54}}
}

func (_this D10) Is_DC31() bool {
	_, ok := _this.Get_().(D10_DC31)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D10) Dtor_cf54() _dafny.Sequence {
	return _this.Get_().(D10_DC31).Cf54
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC31:
		{
			return "D10.DC31" + "(" + data.Cf54.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC31:
		{
			data2, ok := other.Get_().(D10_DC31)
			return ok && data1.Cf54.Equals(data2.Cf54)
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
	Cf56 bool
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf56 bool) D11 {
	return D11{D11_DC33{Cf56}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

type D11_DC32 struct {
	Cf55 _dafny.MultiSet
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf55 _dafny.MultiSet) D11 {
	return D11{D11_DC32{Cf55}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC34 struct {
	Cf57 D11
}

func (D11_DC34) isD11() {}

func (CompanionStruct_D11_) Create_DC34_(Cf57 D11) D11 {
	return D11{D11_DC34{Cf57}}
}

func (_this D11) Is_DC34() bool {
	_, ok := _this.Get_().(D11_DC34)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC33_(false)
}

func (_this D11) Dtor_cf56() bool {
	return _this.Get_().(D11_DC33).Cf56
}

func (_this D11) Dtor_cf55() _dafny.MultiSet {
	return _this.Get_().(D11_DC32).Cf55
}

func (_this D11) Dtor_cf57() D11 {
	return _this.Get_().(D11_DC34).Cf57
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D11_DC34:
		{
			return "D11.DC34" + "(" + _dafny.String(data.Cf57) + ")"
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
			return ok && data1.Cf56 == data2.Cf56
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf55.Equals(data2.Cf55)
		}
	case D11_DC34:
		{
			data2, ok := other.Get_().(D11_DC34)
			return ok && data1.Cf57.Equals(data2.Cf57)
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

type D12_DC36 struct {
	Cf59 _dafny.Int
	Cf60 _dafny.Int
	Cf61 bool
	Cf62 _dafny.Int
	Cf63 bool
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf59 _dafny.Int, Cf60 _dafny.Int, Cf61 bool, Cf62 _dafny.Int, Cf63 bool) D12 {
	return D12{D12_DC36{Cf59, Cf60, Cf61, Cf62, Cf63}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC37 struct {
	Cf64 _dafny.Int
	Cf65 bool
	Cf66 bool
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf64 _dafny.Int, Cf65 bool, Cf66 bool) D12 {
	return D12{D12_DC37{Cf64, Cf65, Cf66}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

type D12_DC35 struct {
	Cf58 _dafny.Map
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf58 _dafny.Map) D12 {
	return D12{D12_DC35{Cf58}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC36_(_dafny.Zero, _dafny.Zero, false, _dafny.Zero, false)
}

func (_this D12) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf59
}

func (_this D12) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf60
}

func (_this D12) Dtor_cf61() bool {
	return _this.Get_().(D12_DC36).Cf61
}

func (_this D12) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf62
}

func (_this D12) Dtor_cf63() bool {
	return _this.Get_().(D12_DC36).Cf63
}

func (_this D12) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D12_DC37).Cf64
}

func (_this D12) Dtor_cf65() bool {
	return _this.Get_().(D12_DC37).Cf65
}

func (_this D12) Dtor_cf66() bool {
	return _this.Get_().(D12_DC37).Cf66
}

func (_this D12) Dtor_cf58() _dafny.Map {
	return _this.Get_().(D12_DC35).Cf58
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ")"
		}
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf58) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60.Cmp(data2.Cf60) == 0 && data1.Cf61 == data2.Cf61 && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63 == data2.Cf63
		}
	case D12_DC37:
		{
			data2, ok := other.Get_().(D12_DC37)
			return ok && data1.Cf64.Cmp(data2.Cf64) == 0 && data1.Cf65 == data2.Cf65 && data1.Cf66 == data2.Cf66
		}
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf58.Equals(data2.Cf58)
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

// Definition of trait T0
type T0 interface {
	String() string
	F7() bool
	F7_set_(value bool)
	Fm4(globalState *GlobalState) _dafny.Int
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
	F3  _dafny.Int
	F4  _dafny.Map
	F6  _dafny.Int
	_f0 bool
	_f1 _dafny.Sequence
	_f2 bool
	_f5 _dafny.MultiSet
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.Zero
	_this.F4 = _dafny.EmptyMap
	_this.F6 = _dafny.Zero
	_this._f0 = false
	_this._f1 = _dafny.EmptySeq
	_this._f2 = false
	_this._f5 = _dafny.EmptyMultiSet
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Sequence, f2 bool, f3 _dafny.Int, f4 _dafny.Map, f5 _dafny.MultiSet, f6 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F5() _dafny.MultiSet {
	{
		return _this._f5
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f7 bool
	_f8 _dafny.Int
	_f9 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f7 = false
	_this._f8 = _dafny.Zero
	_this._f9 = _dafny.Zero
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

func (_this *C0) F7() bool {
	return _this._f7
}
func (_this *C0) F7_set_(value bool) {
	_this._f7 = value
}
func (_this *C0) Ctor__(f8 _dafny.Int, f9 _dafny.Int, f7 bool) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f7 = f7
	}
}
func (_this *C0) Fm4(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(-635)).Minus((_this).F9())
	}
}
func (_this *C0) Fm5(p0 _dafny.Map, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if _this.F7() {
				return _this.F7()
			}
			return _this.F7()
		})(), _this.F7())).Cardinality()
	}
}
func (_this *C0) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *C0) F9() _dafny.Int {
	{
		return _this._f9
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
