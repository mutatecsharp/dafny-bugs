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
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('v')), _dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('w')))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('s')
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 bool, p2 D0, p3 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(false))).Intersection(_dafny.SetOf(!(false)))).Union(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 bool, p3 bool, globalState *GlobalState) bool {
	return ((_dafny.IntOfInt64(548)).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(412))).Cardinality()))).Cmp(_dafny.IntOfInt64(-314)) >= 0
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(421), _dafny.IntOfInt64(-17))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false), false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfInt64(100)).Plus((Companion_D1_.Create_DC3_(_dafny.IntOfInt64(339), !(!(!(!(true)))))).Dtor_cf6())
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, globalState *GlobalState) D1 {
	var _source0 D3 = Companion_D3_.Create_DC8_((_dafny.MultiSetOf(_dafny.IntOfInt64(-463), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("u")).Cardinality())), _dafny.IntOfInt64(851), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cdurb"), _dafny.UnicodeSeqOfUtf8Bytes("amh"))).Cardinality()), _dafny.IntOfInt64(331))).Cardinality())
	_ = _source0
	if _source0.Is_DC8() {
		var _0___mcc_h0 _dafny.Int = _source0.Get_().(D3_DC8).Cf13
		_ = _0___mcc_h0
		var _1_cf13 _dafny.Int = _0___mcc_h0
		_ = _1_cf13
		return Companion_D1_.Create_DC3_(_1_cf13, false)
	} else if _source0.Is_DC7() {
		var _2___mcc_h1 _dafny.Map = _source0.Get_().(D3_DC7).Cf12
		_ = _2___mcc_h1
		var _3_cf12 _dafny.Map = _2___mcc_h1
		_ = _3_cf12
		return Companion_D1_.Create_DC3_((_dafny.MultiSetOf(true, true)).Cardinality(), false)
	} else {
		var _4___mcc_h2 D3 = _source0.Get_().(D3_DC9).Cf14
		_ = _4___mcc_h2
		var _5_cf14 D3 = _4___mcc_h2
		_ = _5_cf14
		return Companion_D1_.Create_DC3_((_dafny.MultiSetOf((_dafny.SetOf(!(false))).Cardinality())).Cardinality(), false)
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(767), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nux")).Cardinality())), ((_dafny.SetOf(_dafny.IntOfInt64(-871), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rqqxre")).Cardinality()))).Cardinality()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	var _source1 D2 = Companion_D2_.Create_DC6_()
	_ = _source1
	if _source1.Is_DC5() {
		var _6___mcc_h0 _dafny.Int = _source1.Get_().(D2_DC5).Cf9
		_ = _6___mcc_h0
		var _7___mcc_h1 _dafny.Int = _source1.Get_().(D2_DC5).Cf10
		_ = _7___mcc_h1
		var _8___mcc_h2 D0 = _source1.Get_().(D2_DC5).Cf11
		_ = _8___mcc_h2
		var _9_cf11 D0 = _8___mcc_h2
		_ = _9_cf11
		var _10_cf10 _dafny.Int = _7___mcc_h1
		_ = _10_cf10
		var _11_cf9 _dafny.Int = _6___mcc_h0
		_ = _11_cf9
		return _dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.CodePoint('y')), Companion_D0_.Create_DC0_(_dafny.CodePoint('j')), Companion_D0_.Create_DC0_(_dafny.CodePoint('w')), Companion_D0_.Create_DC0_(_dafny.CodePoint('h')))
	} else if _source1.Is_DC6() {
		return _dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.CodePoint('b')), Companion_D0_.Create_DC0_(_dafny.CodePoint('p')))
	} else {
		var _12___mcc_h3 _dafny.Array = _source1.Get_().(D2_DC4).Cf8
		_ = _12___mcc_h3
		var _13_cf8 _dafny.Array = _12___mcc_h3
		_ = _13_cf8
		return _dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.CodePoint('y')), Companion_D0_.Create_DC0_(_dafny.CodePoint('r')))
	}
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (bool, _dafny.Set, _dafny.Sequence) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Set = _dafny.EmptySet
	_ = r1
	var r2 _dafny.Sequence = _dafny.EmptySeq
	_ = r2
	var _14_v0 _dafny.Int
	_ = _14_v0
	_14_v0 = _dafny.IntOfInt64(506)
	var _15_v1 _dafny.Sequence
	_ = _15_v1
	_15_v1 = _dafny.UnicodeSeqOfUtf8Bytes("skwjs")
	var _16_v2 _dafny.Map
	_ = _16_v2
	_16_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_14_v0, _dafny.IntOfUint32((_15_v1).Cardinality())), _dafny.UnicodeSeqOfUtf8Bytes("k"))
	var _17_v3 _dafny.Sequence
	_ = _17_v3
	_17_v3 = _dafny.SeqOf(_15_v1)
	_16_v2 = (_16_v2).Update((_14_v0).Times(_dafny.IntOfInt64(61)), _dafny.Companion_Sequence_.Concatenate((_17_v3).Select((Companion_Default___.SafeIndex(_14_v0, _dafny.IntOfUint32((_17_v3).Cardinality()))).Uint32()).(_dafny.Sequence), _15_v1))
	var _18_v4 _dafny.Array
	_ = _18_v4
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
	_ = _nw0
	_18_v4 = _nw0
	var _19_v5 bool
	_ = _19_v5
	_19_v5 = true
	var _20_v6 D1
	_ = _20_v6
	_20_v6 = Companion_D1_.Create_DC3_((_dafny.Zero).Minus(_14_v0), _19_v5)
	var _21_v7 _dafny.Map
	_ = _21_v7
	_21_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _dafny.SetOf(_20_v6))
	var _pat_let_tv0 = _14_v0
	_ = _pat_let_tv0
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_18_v4), 0))
	_ = _index0
	(_18_v4).ArraySet1((func() _dafny.Set {
		if (_21_v7).Contains(_19_v5) {
			return (_21_v7).Get(_19_v5).(_dafny.Set)
		}
		return _dafny.SetOf(_20_v6, func(_pat_let0_0 D1) D1 {
			return func(_22_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 _dafny.Int) D1 {
					return func(_23_dt__update_hcf6_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC3_(_23_dt__update_hcf6_h0, (_22_dt__update__tmp_h0).Dtor_cf7())
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_20_v6), _20_v6)
	})(), (_index0).Int())
	var _24_v8 _dafny.Set
	_ = _24_v8
	_24_v8 = _dafny.SetOf(_20_v6)
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_18_v4), 0))
	_ = _index1
	(_18_v4).ArraySet1(_24_v8, (_index1).Int())
	var _25_i0 _dafny.Int
	_ = _25_i0
	_25_i0 = _dafny.Zero
	{
		for (_19_v5) == (_19_v5) {
			{
				if (_25_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_25_i0 = (_25_i0).Plus(_dafny.One)
				_14_v0 = (_dafny.Zero).Minus(_14_v0)
				var _26_v9 _dafny.Array
				_ = _26_v9
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_0
				var _nw1 _dafny.Array
				_ = _nw1
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw1 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.Int = (func(_27_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_28_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_28_i1, _27_v0)
						}
					})(_14_v0)
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
				_26_v9 = _nw1
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_26_v9), 0))
				_ = _index2
				(_26_v9).ArraySet1(_14_v0, (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_26_v9), 0))
				_ = _index3
				(_26_v9).ArraySet1(_14_v0, (_index3).Int())
				var _29_v10 _dafny.CodePoint
				_ = _29_v10
				_29_v10 = _dafny.CodePoint('k')
				var _30_v11 D0
				_ = _30_v11
				_30_v11 = Companion_D0_.Create_DC0_(_29_v10)
				var _31_v12 _dafny.Sequence
				_ = _31_v12
				_31_v12 = _dafny.SeqOf(_30_v11)
				var _32_v13 _dafny.Sequence
				_ = _32_v13
				_32_v13 = _dafny.SeqOf((_26_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_26_v9), 0))).Int()).(_dafny.Int))
				var _33_v14 _dafny.Sequence
				_ = _33_v14
				_33_v14 = _dafny.SeqOf(true, _19_v5, _19_v5, !(_19_v5))
				var _34_v15 _dafny.Sequence
				_ = _34_v15
				_34_v15 = _dafny.SeqOf((_26_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_26_v9), 0))).Int()).(_dafny.Int), (_32_v13).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(650), _19_v5)).Cardinality(), _dafny.IntOfUint32((_32_v13).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_33_v14).Cardinality()), _14_v0, _dafny.IntOfUint32((_15_v1).Cardinality()))
				var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_31_v12, Companion_Default___.Fm9(_14_v0, (_26_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_26_v9), 0))).Int()).(_dafny.Int), _19_v5, _19_v5, globalState)), _31_v12)
				_ = _rhs0
				var _rhs1 bool = Companion_Default___.Fm4(_15_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(811))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg0 _dafny.Int) interface{} {
						return coer0(arg0)
					}
				}((func(_35_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_36_i2 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_35_v1).Cardinality())
					}
				})(_15_v1))), _32_v13), _19_v5, _19_v5, globalState)
				_ = _rhs1
				_31_v12 = _rhs0
				_19_v5 = _rhs1
				r0 = _19_v5
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _37_v16 _dafny.Sequence
	_ = _37_v16
	_37_v16 = _dafny.SeqOf(_14_v0)
	(globalState).F6 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(706), _14_v0), (_14_v0).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm4(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(476))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_38_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	})), _37_v16, _19_v5, _19_v5, globalState)), _14_v0)).Cardinality())))
	var _39_v17 *C0
	_ = _39_v17
	var _nw2 *C0 = New_C0_()
	_ = _nw2
	_nw2.Ctor__()
	_39_v17 = _nw2
	var _40_v18 _dafny.Map
	_ = _40_v18
	_40_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v0, _39_v17)
	var _41_v19 _dafny.MultiSet
	_ = _41_v19
	_41_v19 = _dafny.MultiSetOf(_40_v18, _40_v18)
	var _42_v20 _dafny.Set
	_ = _42_v20
	_42_v20 = _dafny.SetOf(_19_v5)
	var _43_v21 _dafny.Map
	_ = _43_v21
	_43_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_41_v19).Contains(_40_v18) {
			return (_41_v19).Multiplicity(_40_v18)
		}
		return Companion_Default___.Fm6(_14_v0, _19_v5, globalState)
	})(), (_42_v20).IsProperSubsetOf(_dafny.SetOf(!(_19_v5))))
	if (func() bool {
		if (_43_v21).Contains((_14_v0).Times(_14_v0)) {
			return (_43_v21).Get((_14_v0).Times(_14_v0)).(bool)
		}
		return true
	})() {
		var _44_v22 *C0
		_ = _44_v22
		var _nw3 *C0 = New_C0_()
		_ = _nw3
		_nw3.Ctor__()
		_44_v22 = _nw3
		var _45_v23 D0
		_ = _45_v23
		_45_v23 = Companion_D0_.Create_DC1_((_14_v0).Cmp(_14_v0) == 0, _14_v0, _19_v5, (_14_v0).Times(_14_v0))
		_45_v23 = _45_v23
		var _46_v24 _dafny.Sequence
		_ = _46_v24
		_46_v24 = _dafny.SeqOf(false, _19_v5)
		var _47_v25 _dafny.Sequence
		_ = _47_v25
		_47_v25 = _dafny.SeqOf(_dafny.SeqOf(_19_v5, _19_v5), _dafny.SeqOf(false))
		var _48_v26 _dafny.Sequence
		_ = _48_v26
		_48_v26 = _dafny.SeqOf(_39_v17)
		var _rhs2 bool = !_dafny.Companion_Sequence_.Contains(_47_v25, _46_v24)
		_ = _rhs2
		var _rhs3 bool = _19_v5
		_ = _rhs3
		var _rhs4 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(883))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_49_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality()), _dafny.IntOfUint32((_48_v26).Cardinality()))
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		_19_v5 = _rhs2
		r0 = _rhs3
		_lhs0.F6 = _rhs4
		var _50_v27 _dafny.MultiSet
		_ = _50_v27
		_50_v27 = _dafny.MultiSetOf(_14_v0, _14_v0, _14_v0)
		var _51_v28 _dafny.MultiSet
		_ = _51_v28
		_51_v28 = _dafny.MultiSetOf(_19_v5)
		var _52_v29 _dafny.CodePoint
		_ = _52_v29
		_52_v29 = _dafny.CodePoint('g')
		var _53_v30 _dafny.Map
		_ = _53_v30
		_53_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_52_v29, !(true))
		var _54_v31 _dafny.Array
		_ = _54_v31
		var _nwElement0_0 bool = _19_v5
		_ = _nwElement0_0
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(23))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_0, 0)
		(_nw4).ArraySet1((_20_v6).Dtor_cf7(), 1)
		(_nw4).ArraySet1(_dafny.Companion_Sequence_.Equal(_15_v1, _15_v1), 2)
		(_nw4).ArraySet1((func() bool {
			if !(_19_v5) {
				return true
			}
			return _19_v5
		})(), 3)
		(_nw4).ArraySet1(_19_v5, 4)
		(_nw4).ArraySet1(_19_v5, 5)
		(_nw4).ArraySet1(_19_v5, 6)
		(_nw4).ArraySet1(_19_v5, 7)
		(_nw4).ArraySet1((_14_v0).Cmp(_14_v0) != 0, 8)
		(_nw4).ArraySet1((_dafny.IntOfUint32((_37_v16).Cardinality())).Cmp(_14_v0) < 0, 9)
		(_nw4).ArraySet1(_19_v5, 10)
		(_nw4).ArraySet1(_19_v5, 11)
		(_nw4).ArraySet1(_19_v5, 12)
		(_nw4).ArraySet1(!((_19_v5) || (false)), 13)
		(_nw4).ArraySet1(((_dafny.Zero).Minus(_14_v0)).Cmp((_42_v20).Cardinality()) > 0, 14)
		(_nw4).ArraySet1((_50_v27).IsDisjointFrom(_50_v27), 15)
		(_nw4).ArraySet1(_19_v5, 16)
		(_nw4).ArraySet1((_51_v28).IsProperSubsetOf(_51_v28), 17)
		(_nw4).ArraySet1(_19_v5, 18)
		(_nw4).ArraySet1(true, 19)
		(_nw4).ArraySet1((func() bool {
			if (_53_v30).Contains(_52_v29) {
				return (_53_v30).Get(_52_v29).(bool)
			}
			return _19_v5
		})(), 20)
		(_nw4).ArraySet1(!(_19_v5), 21)
		(_nw4).ArraySet1(_19_v5, 22)
		_54_v31 = _nw4
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))
		_ = _index4
		(_54_v31).ArraySet1(_19_v5, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))
		_ = _index5
		(_54_v31).ArraySet1(!(_19_v5) || (_19_v5), (_index5).Int())
		var _55_v32 _dafny.Map
		_ = _55_v32
		_55_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _52_v29)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))
		_ = _index6
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))
		_ = _index7
		var _rhs5 bool = (_54_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))).Int()).(bool)
		_ = _rhs5
		var _rhs6 bool = (((_dafny.Zero).Minus(_14_v0)).Minus(_14_v0)).Cmp(_14_v0) < 0
		_ = _rhs6
		var _rhs7 bool = (_dafny.IntOfUint32((_15_v1).Cardinality())).Cmp(_14_v0) >= 0
		_ = _rhs7
		var _rhs8 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_15_v1, (Companion_Default___.SafeIndex(_14_v0, _dafny.IntOfUint32((_15_v1).Cardinality()))).Uint32(), _52_v29)
		_ = _rhs8
		var _rhs9 _dafny.Map = _55_v32
		_ = _rhs9
		var _lhs1 _dafny.Array = _54_v31
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))
		_ = _lhs2
		var _lhs3 _dafny.Array = _54_v31
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_54_v31), 0))
		_ = _lhs4
		r0 = _rhs5
		(_lhs1).ArraySet1(_rhs6, (_lhs2).Int())
		(_lhs3).ArraySet1(_rhs7, (_lhs4).Int())
		r2 = _rhs8
		_55_v32 = _rhs9
	} else {
		(globalState).F7 = (_20_v6).Dtor_cf6()
		(globalState).F6 = (_dafny.Zero).Minus((_37_v16).Select((Companion_Default___.SafeIndex(_14_v0, _dafny.IntOfUint32((_37_v16).Cardinality()))).Uint32()).(_dafny.Int))
		var _56_v33 _dafny.Array
		_ = _56_v33
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
		_ = _len0_1
		var _nw5 _dafny.Array
		_ = _nw5
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw5 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = (func(_57_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_58_i5 _dafny.Int) _dafny.Int {
					return (_58_i5).Minus(_57_v0)
				}
			})(_14_v0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw5 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw5).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw5).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_56_v33 = _nw5
		var _59_v34 _dafny.MultiSet
		_ = _59_v34
		_59_v34 = _dafny.MultiSetOf(_56_v33, _56_v33)
		if (_dafny.MultiSetOf(_56_v33)).Equals(_59_v34) {
			var _60_v35 _dafny.Sequence
			_ = _60_v35
			_60_v35 = _dafny.SeqOf(_19_v5, _19_v5)
			_60_v35 = _dafny.Companion_Sequence_.Concatenate(_60_v35, _60_v35)
			_19_v5 = ((_14_v0).Minus(_14_v0)).Cmp(_14_v0) < 0
			(globalState).F8 = ((_dafny.IntOfUint32((_60_v35).Cardinality())).Times(_14_v0)).Cmp(_dafny.IntOfInt64(245)) >= 0
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_56_v33), 0))
			_ = _index8
			(_56_v33).ArraySet1((_14_v0).Minus(_14_v0), (_index8).Int())
			var _61_v36 _dafny.Map
			_ = _61_v36
			_61_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v33, Companion_Default___.SafeDivisionInt(_14_v0, _14_v0))
			var _62_v37 _dafny.Sequence
			_ = _62_v37
			_62_v37 = _dafny.SeqOf(_56_v33)
			var _63_v38 _dafny.Map
			_ = _63_v38
			_63_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_62_v37).Cardinality()), _14_v0)
			var _64_v39 _dafny.Set
			_ = _64_v39
			_64_v39 = _dafny.SetOf(_63_v38, _63_v38)
			var _65_v40 _dafny.Map
			_ = _65_v40
			_65_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _19_v5)
			var _66_v41 _dafny.Map
			_ = _66_v41
			_66_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _39_v17)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_56_v33), 0))
			_ = _index9
			var _rhs10 bool = _19_v5
			_ = _rhs10
			var _rhs11 _dafny.Int = (func() _dafny.Int {
				if (_61_v36).Contains(_56_v33) {
					return (_61_v36).Get(_56_v33).(_dafny.Int)
				}
				return _14_v0
			})()
			_ = _rhs11
			var _rhs12 _dafny.Int = _14_v0
			_ = _rhs12
			var _rhs13 bool = _19_v5
			_ = _rhs13
			var _rhs14 bool = (Companion_D0_.Create_DC1_(_19_v5, (_64_v39).Cardinality(), (func() bool {
				if (_65_v40).Contains(Companion_Default___.Fm4(_15_v1, _dafny.SeqOf(_14_v0), _19_v5, _19_v5, globalState)) {
					return (_65_v40).Get(Companion_Default___.Fm4(_15_v1, _dafny.SeqOf(_14_v0), _19_v5, _19_v5, globalState)).(bool)
				}
				return _19_v5
			})(), (_66_v41).Cardinality())).Dtor_cf1()
			_ = _rhs14
			var _lhs5 _dafny.Array = _56_v33
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_56_v33), 0))
			_ = _lhs6
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			r0 = _rhs10
			(_lhs5).ArraySet1(_rhs11, (_lhs6).Int())
			_lhs7.F6 = _rhs12
			r0 = _rhs13
			_lhs8.F8 = _rhs14
			_60_v35 = _dafny.Companion_Sequence_.Concatenate(_60_v35, _60_v35)
		} else {
			(globalState).F6 = _14_v0
			var _67_v42 _dafny.Sequence
			_ = _67_v42
			_67_v42 = _dafny.SeqOf(_39_v17, _39_v17)
			_39_v17 = (_67_v42).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(473))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_68_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality()), _dafny.IntOfUint32((_67_v42).Cardinality()))).Uint32()).(*C0)
			var _69_v43 _dafny.Sequence
			_ = _69_v43
			_69_v43 = _dafny.SeqOf(_19_v5, _19_v5)
			var _70_v44 _dafny.MultiSet
			_ = _70_v44
			_70_v44 = _dafny.MultiSetOf(_14_v0)
			var _71_v45 _dafny.Map
			_ = _71_v45
			_71_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v43, (_70_v44).IsDisjointFrom(_70_v44))
			var _rhs15 _dafny.Int = Companion_Default___.SafeDivisionInt(_14_v0, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-380), _14_v0))
			_ = _rhs15
			var _rhs16 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_15_v1, _15_v1)).Cardinality())).Times(_14_v0)
			_ = _rhs16
			var _rhs17 _dafny.Map = _71_v45
			_ = _rhs17
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			var _lhs10 *GlobalState = globalState
			_ = _lhs10
			_lhs9.F3 = _rhs15
			_lhs10.F3 = _rhs16
			_71_v45 = _rhs17
			var _72_v46 *C0
			_ = _72_v46
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__()
			_72_v46 = _nw6
			(globalState).F3 = _14_v0
		}
		var _73_v47 _dafny.Array
		_ = _73_v47
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw7
		_73_v47 = _nw7
		var _74_v48 _dafny.Array
		_ = _74_v48
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_2
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_75_v5 bool) func(_dafny.Int) bool {
				return func(_76_i7 _dafny.Int) bool {
					return _75_v5
				}
			})(_19_v5)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw8 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw8).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw8).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_74_v48 = _nw8
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_73_v47), 0))
		_ = _index10
		(_73_v47).ArraySet1(_74_v48, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_73_v47), 0))
		_ = _index11
		(_73_v47).ArraySet1(_74_v48, (_index11).Int())
		var _77_v49 D0
		_ = _77_v49
		_77_v49 = Companion_D0_.Create_DC1_(_19_v5, _dafny.IntOfInt64(34), !(_19_v5), (_dafny.Zero).Minus(((_42_v20).Union(_42_v20)).Cardinality()))
		var _source2 D0 = _77_v49
		_ = _source2
		if _source2.Is_DC1() {
			var _78___mcc_h0 bool = _source2.Get_().(D0_DC1).Cf1
			_ = _78___mcc_h0
			var _79___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
			_ = _79___mcc_h1
			var _80___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf3
			_ = _80___mcc_h2
			var _81___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC1).Cf4
			_ = _81___mcc_h3
			var _82_cf4 _dafny.Int = _81___mcc_h3
			_ = _82_cf4
			var _83_cf3 bool = _80___mcc_h2
			_ = _83_cf3
			var _84_cf2 _dafny.Int = _79___mcc_h1
			_ = _84_cf2
			var _85_cf1 bool = _78___mcc_h0
			_ = _85_cf1
			r2 = _15_v1
			_74_v48 = _dafny.ArrayCastTo((_73_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_73_v47), 0))).Int()))
			(globalState).F3 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vspevqskp")).Cardinality())
			var _86_v50 _dafny.Map
			_ = _86_v50
			_86_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_cf3, _82_cf4)
			_16_v2 = (_16_v2).Update(((_86_v50).Cardinality()).Times(_14_v0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_87_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			})))
		} else {
			var _88___mcc_h4 _dafny.CodePoint = _source2.Get_().(D0_DC0).Cf0
			_ = _88___mcc_h4
			var _89_cf0 _dafny.CodePoint = _88___mcc_h4
			_ = _89_cf0
			(globalState).F3 = (_37_v16).Select((Companion_Default___.SafeIndex((_42_v20).Cardinality(), _dafny.IntOfUint32((_37_v16).Cardinality()))).Uint32()).(_dafny.Int)
			var _90_v51 _dafny.MultiSet
			_ = _90_v51
			_90_v51 = _dafny.MultiSetOf(_19_v5)
			var _91_v52 _dafny.Set
			_ = _91_v52
			_91_v52 = _dafny.SetOf(_90_v51, _90_v51, _90_v51)
			var _92_v53 _dafny.Map
			_ = _92_v53
			_92_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _91_v52)
			_91_v52 = (func() _dafny.Set {
				if (_92_v53).Contains(Companion_Default___.Fm4(_dafny.UnicodeSeqOfUtf8Bytes("x"), _37_v16, !(!(_19_v5)), !(_19_v5), globalState)) {
					return (_92_v53).Get(Companion_Default___.Fm4(_dafny.UnicodeSeqOfUtf8Bytes("x"), _37_v16, !(!(_19_v5)), !(_19_v5), globalState)).(_dafny.Set)
				}
				return _91_v52
			})()
			(globalState).F8 = _19_v5
			var _93_v54 _dafny.Array
			_ = _93_v54
			var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw9
			_93_v54 = _nw9
			var _94_v55 _dafny.Array
			_ = _94_v55
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_3
			var _nw10 _dafny.Array
			_ = _nw10
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw10 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = func(_95_i9 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(false)
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw10 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw10).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw10).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_94_v55 = _nw10
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_93_v54), 0))
			_ = _index12
			(_93_v54).ArraySet1(_94_v55, (_index12).Int())
			var _96_v56 D1
			_ = _96_v56
			_96_v56 = Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_15_v1, _37_v16, _19_v5, _19_v5, globalState), _14_v0))
			var _97_v57 _dafny.Sequence
			_ = _97_v57
			_97_v57 = _dafny.SeqOf(_94_v55, _94_v55)
			var _98_v58 _dafny.Sequence
			_ = _98_v58
			_98_v58 = _dafny.SeqOf(_94_v55, (_97_v57).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_14_v0), _dafny.IntOfUint32((_97_v57).Cardinality()))).Uint32()).(_dafny.Array), _94_v55)
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_93_v54), 0))
			_ = _index13
			var _rhs18 _dafny.Array = (_97_v57).Select((Companion_Default___.SafeIndex(_14_v0, _dafny.IntOfUint32((_97_v57).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs18
			var _rhs19 _dafny.Int = (Companion_Default___.Fm6(_14_v0, _19_v5, globalState)).Times(Companion_Default___.Fm6(_14_v0, _19_v5, globalState))
			_ = _rhs19
			var _rhs20 D1 = _96_v56
			_ = _rhs20
			var _lhs11 _dafny.Array = _93_v54
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_93_v54), 0))
			_ = _lhs12
			var _lhs13 *GlobalState = globalState
			_ = _lhs13
			(_lhs11).ArraySet1(_rhs18, (_lhs12).Int())
			_lhs13.F6 = _rhs19
			_96_v56 = _rhs20
		}
	}
	var _99_v59 _dafny.Sequence
	_ = _99_v59
	_99_v59 = _dafny.SeqOf(_42_v20, _42_v20)
	var _100_i10 _dafny.Int
	_ = _100_i10
	_100_i10 = _dafny.Zero
	{
		for (_dafny.IntOfUint32((_37_v16).Cardinality())).Cmp(((_99_v59).Select((Companion_Default___.SafeIndex(_14_v0, _dafny.IntOfUint32((_99_v59).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()) <= 0 {
			{
				if (_100_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_100_i10 = (_100_i10).Plus(_dafny.One)
				var _101_v60 _dafny.Array
				_ = _101_v60
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_4
				var _nw11 _dafny.Array
				_ = _nw11
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw11 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Int = (func(_102_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_103_i11 _dafny.Int) _dafny.Int {
							return (_103_i11).Times(_102_v0)
						}
					})(_14_v0)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw11 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw11).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw11).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_101_v60 = _nw11
				var _104_v61 _dafny.MultiSet
				_ = _104_v61
				_104_v61 = _dafny.MultiSetOf(_101_v60, _101_v60, _101_v60, _101_v60)
				var _rhs21 _dafny.Array = (Companion_D2_.Create_DC4_(_101_v60)).Dtor_cf8()
				_ = _rhs21
				var _rhs22 _dafny.Int = _14_v0
				_ = _rhs22
				var _rhs23 bool = _19_v5
				_ = _rhs23
				var _rhs24 bool = _19_v5
				_ = _rhs24
				var _rhs25 _dafny.Int = ((_104_v61).Update(_101_v60, Companion_Default___.Abs((_14_v0).Minus(_14_v0)))).Cardinality()
				_ = _rhs25
				var _lhs14 *GlobalState = globalState
				_ = _lhs14
				var _lhs15 *GlobalState = globalState
				_ = _lhs15
				var _lhs16 *GlobalState = globalState
				_ = _lhs16
				_101_v60 = _rhs21
				_14_v0 = _rhs22
				_lhs14.F8 = _rhs23
				_lhs15.F8 = _rhs24
				_lhs16.F6 = _rhs25
				if _19_v5 {
					var _105_v62 _dafny.Map
					_ = _105_v62
					_105_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(292), _14_v0)
					var _106_v63 _dafny.CodePoint
					_ = _106_v63
					_106_v63 = _dafny.CodePoint('p')
					var _107_v64 _dafny.MultiSet
					_ = _107_v64
					_107_v64 = _dafny.MultiSetOf(_106_v63)
					var _108_v65 _dafny.Array
					_ = _108_v65
					var _nwElement0_1 bool = _19_v5
					_ = _nwElement0_1
					var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(21))
					_ = _nw12
					(_nw12).ArraySet1(_nwElement0_1, 0)
					(_nw12).ArraySet1(_19_v5, 1)
					(_nw12).ArraySet1(Companion_Default___.Fm4(_dafny.UnicodeSeqOfUtf8Bytes("owms"), _37_v16, _19_v5, _19_v5, globalState), 2)
					(_nw12).ArraySet1(_19_v5, 3)
					(_nw12).ArraySet1(_19_v5, 4)
					(_nw12).ArraySet1(_19_v5, 5)
					(_nw12).ArraySet1(_19_v5, 6)
					(_nw12).ArraySet1(((_105_v62).Cardinality()).Cmp(_14_v0) > 0, 7)
					(_nw12).ArraySet1(false, 8)
					(_nw12).ArraySet1(_19_v5, 9)
					(_nw12).ArraySet1((_107_v64).IsProperSubsetOf(_107_v64), 10)
					(_nw12).ArraySet1(_19_v5, 11)
					(_nw12).ArraySet1(_19_v5, 12)
					(_nw12).ArraySet1(_19_v5, 13)
					(_nw12).ArraySet1(_19_v5, 14)
					(_nw12).ArraySet1(_19_v5, 15)
					(_nw12).ArraySet1(_19_v5, 16)
					(_nw12).ArraySet1(_19_v5, 17)
					(_nw12).ArraySet1(_19_v5, 18)
					(_nw12).ArraySet1(_19_v5, 19)
					(_nw12).ArraySet1(_19_v5, 20)
					_108_v65 = _nw12
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_108_v65), 0))
					_ = _index14
					(_108_v65).ArraySet1((func() bool {
						if _19_v5 {
							return _19_v5
						}
						return true
					})(), (_index14).Int())
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_108_v65), 0))
					_ = _index15
					(_108_v65).ArraySet1(_19_v5, (_index15).Int())
					_39_v17 = _39_v17
					var _109_v66 _dafny.Array
					_ = _109_v66
					var _nw13 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
					_ = _nw13
					_109_v66 = _nw13
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_109_v66), 0))
					_ = _index16
					(_109_v66).ArraySet1(_39_v17, (_index16).Int())
					var _110_v67 _dafny.Sequence
					_ = _110_v67
					_110_v67 = _dafny.SeqOf(_19_v5, (_108_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(836), _dafny.ArrayLen((_108_v65), 0))).Int()).(bool))
					var _111_v68 _dafny.MultiSet
					_ = _111_v68
					_111_v68 = _dafny.MultiSetOf(_19_v5, false)
					var _112_v69 _dafny.Set
					_ = _112_v69
					_112_v69 = _dafny.SetOf(_dafny.IntOfUint32((_110_v67).Cardinality()), (_111_v68).Cardinality(), _14_v0)
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_109_v66), 0))
					_ = _index17
					var _rhs26 *C0 = _39_v17
					_ = _rhs26
					var _rhs27 *C0 = _39_v17
					_ = _rhs27
					var _rhs28 _dafny.Int = (((_112_v69).Cardinality()).Times(_14_v0)).Plus(_14_v0)
					_ = _rhs28
					var _rhs29 _dafny.Int = (_dafny.IntOfInt64(708)).Times((_43_v21).Cardinality())
					_ = _rhs29
					var _lhs17 _dafny.Array = _109_v66
					_ = _lhs17
					var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(624), _dafny.ArrayLen((_109_v66), 0))
					_ = _lhs18
					var _lhs19 *GlobalState = globalState
					_ = _lhs19
					var _lhs20 *GlobalState = globalState
					_ = _lhs20
					_39_v17 = _rhs26
					(_lhs17).ArraySet1(_rhs27, (_lhs18).Int())
					_lhs19.F3 = _rhs28
					_lhs20.F3 = _rhs29
					(globalState).F6 = _14_v0
					var _113_v70 *C0
					_ = _113_v70
					var _nw14 *C0 = New_C0_()
					_ = _nw14
					_nw14.Ctor__()
					_113_v70 = _nw14
				} else {
					(globalState).F8 = _19_v5
					(globalState).F3 = _14_v0
					var _114_v71 _dafny.Set
					_ = _114_v71
					_114_v71 = _dafny.SetOf(_14_v0, (_dafny.Zero).Minus(_14_v0))
					_19_v5 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_37_v16, _37_v16), (_114_v71).Cardinality())
					var _115_v72 D0
					_ = _115_v72
					_115_v72 = Companion_D0_.Create_DC1_(false, (_dafny.Zero).Minus(_14_v0), !(_19_v5), _14_v0)
					var _116_v73 D2
					_ = _116_v73
					_116_v73 = Companion_D2_.Create_DC5_(_14_v0, _14_v0, _115_v72)
					var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_101_v60), 0))
					_ = _index18
					(_101_v60).ArraySet1(Companion_Default___.SafeModuloInt(_14_v0, (_116_v73).Dtor_cf9()), (_index18).Int())
					var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_101_v60), 0))
					_ = _index19
					(_101_v60).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("poflqgduh")).Cardinality()), (_index19).Int())
					var _117_v74 _dafny.Array
					_ = _117_v74
					var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
					_ = _nw15
					_117_v74 = _nw15
					var _118_v75 _dafny.Map
					_ = _118_v75
					_118_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _117_v74)
					var _119_v76 _dafny.Sequence
					_ = _119_v76
					_119_v76 = _dafny.SeqOf(_118_v75, _118_v75)
					var _120_v77 D3
					_ = _120_v77
					_120_v77 = Companion_D3_.Create_DC7_(_118_v75)
					var _121_v78 _dafny.Sequence
					_ = _121_v78
					_121_v78 = _dafny.SeqOf(_19_v5)
					var _122_v79 _dafny.Array
					_ = _122_v79
					var _nwElement0_2 _dafny.Map = _118_v75
					_ = _nwElement0_2
					var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(20))
					_ = _nw16
					(_nw16).ArraySet1(_nwElement0_2, 0)
					(_nw16).ArraySet1(_118_v75, 1)
					(_nw16).ArraySet1(_118_v75, 2)
					(_nw16).ArraySet1(_118_v75, 3)
					(_nw16).ArraySet1(_118_v75, 4)
					(_nw16).ArraySet1(_118_v75, 5)
					(_nw16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-688))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg5 _dafny.Int) interface{} {
							return coer5(arg5)
						}
					}(func(_123_i12 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					})), _dafny.SeqOf((_43_v21).Cardinality(), _14_v0), _19_v5, _19_v5, globalState), _117_v74), 6)
					(_nw16).ArraySet1(_118_v75, 7)
					(_nw16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _117_v74), 8)
					(_nw16).ArraySet1((_118_v75).Merge((_119_v76).Select((Companion_Default___.SafeIndex(_14_v0, _dafny.IntOfUint32((_119_v76).Cardinality()))).Uint32()).(_dafny.Map)), 9)
					(_nw16).ArraySet1(_118_v75, 10)
					(_nw16).ArraySet1(_118_v75, 11)
					(_nw16).ArraySet1((_120_v77).Dtor_cf12(), 12)
					(_nw16).ArraySet1(_118_v75, 13)
					(_nw16).ArraySet1(_118_v75, 14)
					(_nw16).ArraySet1(_118_v75, 15)
					(_nw16).ArraySet1((_118_v75).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _117_v74)), 16)
					(_nw16).ArraySet1((_118_v75).Merge(_118_v75), 17)
					(_nw16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_121_v78).Select((Companion_Default___.SafeIndex((_101_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_101_v60), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_121_v78).Cardinality()))).Uint32()).(bool), _117_v74), 18)
					(_nw16).ArraySet1(_118_v75, 19)
					_122_v79 = _nw16
					var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_122_v79), 0))
					_ = _index20
					(_122_v79).ArraySet1((_118_v75).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _117_v74)), (_index20).Int())
					var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_122_v79), 0))
					_ = _index21
					(_122_v79).ArraySet1(_118_v75, (_index21).Int())
				}
				var _124_v80 D0
				_ = _124_v80
				_124_v80 = Companion_D0_.Create_DC1_(true, _14_v0, _19_v5, _dafny.IntOfInt64(848))
				(globalState).F7 = Companion_Default___.SafeModuloInt((_124_v80).Dtor_cf2(), _dafny.IntOfInt64(34))
				var _125_v81 *C0
				_ = _125_v81
				var _nw17 *C0 = New_C0_()
				_ = _nw17
				_nw17.Ctor__()
				_125_v81 = _nw17
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	r0 = _19_v5
	r1 = func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v0, _14_v0)).Cardinality()), _37_v16)).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _126_v82 _dafny.Int
			_126_v82 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v0, _14_v0)).Cardinality()), _37_v16), _126_v82) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_126_v82, _dafny.IntOfInt64(669)))
			}
		}
		return _coll0.ToSet()
	}()
	var _127_v83 _dafny.Map
	_ = _127_v83
	_127_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_19_v5, _14_v0)
	r2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_127_v83).Cardinality(), _19_v5, globalState), _15_v1), _15_v1)
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _128_v0 _dafny.Int
	_ = _128_v0
	_128_v0 = _dafny.IntOfInt64(-183)
	var _129_v1 _dafny.MultiSet
	_ = _129_v1
	_129_v1 = _dafny.MultiSetOf(_128_v0)
	var _130_v2 _dafny.Array
	_ = _130_v2
	var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
	_ = _nw18
	_130_v2 = _nw18
	var _131_v3 _dafny.Sequence
	_ = _131_v3
	_131_v3 = _dafny.UnicodeSeqOfUtf8Bytes("cqq")
	var _132_v4 _dafny.Set
	_ = _132_v4
	_132_v4 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("rwijw"), _131_v3)
	var _133_globalState *GlobalState
	_ = _133_globalState
	var _nw19 *GlobalState = New_GlobalState_()
	_ = _nw19
	_nw19.Ctor__(_129_v1, _130_v2, false, _dafny.IntOfInt64(434), _dafny.IntOfInt64(588), _dafny.IntOfInt64(870), _dafny.IntOfInt64(77), _dafny.IntOfInt64(-390), false, true, _132_v4, true, _dafny.IntOfInt64(980), _dafny.IntOfInt64(718))
	_133_globalState = _nw19
	var _134_v5 bool
	_ = _134_v5
	_134_v5 = false
	var _135_v6 _dafny.Map
	_ = _135_v6
	_135_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, !(_134_v5)), _128_v0)
	var _136_v7 _dafny.Map
	_ = _136_v7
	_136_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-421), false)
	_135_v6 = (_135_v6).Update(_136_v7, (_dafny.Zero).Minus(_128_v0))
	var _137_v8 *C0
	_ = _137_v8
	var _nw20 *C0 = New_C0_()
	_ = _nw20
	_nw20.Ctor__()
	_137_v8 = _nw20
	var _138_v9 _dafny.Array
	_ = _138_v9
	var _nw21 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
	_ = _nw21
	_138_v9 = _nw21
	for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_138_v9), 0))); ; {
		_guard_loop_0, _ok1 := _iter1()
		if !_ok1 {
			break
		}
		var _139_i0 _dafny.Int
		_139_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_139_i0).Sign() != -1) && ((_139_i0).Cmp(_dafny.ArrayLen((_138_v9), 0)) < 0)) {
			(_138_v9).ArraySet1(_134_v5, (_139_i0).Int())
		}
	}
	var _140_v10 bool
	_ = _140_v10
	var _141_v11 _dafny.Set
	_ = _141_v11
	var _142_v12 _dafny.Sequence
	_ = _142_v12
	var _out0 bool
	_ = _out0
	var _out1 _dafny.Set
	_ = _out1
	var _out2 _dafny.Sequence
	_ = _out2
	_out0, _out1, _out2 = Companion_Default___.M0(_133_globalState)
	_140_v10 = _out0
	_141_v11 = _out1
	_142_v12 = _out2
	var _143_i1 _dafny.Int
	_ = _143_i1
	_143_i1 = _dafny.Zero
	{
		for _140_v10 {
			{
				if (_143_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_143_i1 = (_143_i1).Plus(_dafny.One)
				_140_v10 = !(!(_134_v5)) || (!((_128_v0).Cmp(_128_v0) < 0))
				var _144_v13 bool
				_ = _144_v13
				var _145_v14 _dafny.Set
				_ = _145_v14
				var _146_v15 _dafny.Sequence
				_ = _146_v15
				var _out3 bool
				_ = _out3
				var _out4 _dafny.Set
				_ = _out4
				var _out5 _dafny.Sequence
				_ = _out5
				_out3, _out4, _out5 = Companion_Default___.M0(_133_globalState)
				_144_v13 = _out3
				_145_v14 = _out4
				_146_v15 = _out5
				(_133_globalState).F3 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tiuwpcrys")).Cardinality())
				var _147_v16 _dafny.Map
				_ = _147_v16
				_147_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _142_v12)
				_142_v12 = _dafny.Companion_Sequence_.Concatenate(_146_v15, (func() _dafny.Sequence {
					if (_147_v16).Contains(_128_v0) {
						return (_147_v16).Get(_128_v0).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(747))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}(func(_148_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					}))
				})())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _149_i3 _dafny.Int
	_ = _149_i3
	_149_i3 = _dafny.Zero
	{
		for (func() bool {
			if (_136_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, true)) {
				return _140_v10
			}
			return _140_v10
		})() {
			{
				if (_149_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_149_i3 = (_149_i3).Plus(_dafny.One)
				(_133_globalState).F3 = _128_v0
				var _150_v17 *C0
				_ = _150_v17
				var _nw22 *C0 = New_C0_()
				_ = _nw22
				_nw22.Ctor__()
				_150_v17 = _nw22
				_131_v3 = _dafny.Companion_Sequence_.Concatenate(_131_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-155))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}(func(_151_i4 _dafny.Int) _dafny.CodePoint {
					return (Companion_D0_.Create_DC0_(_dafny.CodePoint('x'))).Dtor_cf0()
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-596))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_152_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))))
				var _153_v18 _dafny.Map
				_ = _153_v18
				_153_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v5, (_134_v5) == (_140_v10))
				(_133_globalState).F6 = (_153_v18).Cardinality()
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _154_i6 _dafny.Int
	_ = _154_i6
	_154_i6 = _dafny.Zero
	{
		for _134_v5 {
			{
				if (_154_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_154_i6 = (_154_i6).Plus(_dafny.One)
				var _155_v19 _dafny.CodePoint
				_ = _155_v19
				_155_v19 = _dafny.CodePoint('v')
				var _156_v20 D0
				_ = _156_v20
				_156_v20 = Companion_D0_.Create_DC0_(_155_v19)
				var _rhs30 D0 = _156_v20
				_ = _rhs30
				var _rhs31 bool = false
				_ = _rhs31
				var _lhs21 *GlobalState = _133_globalState
				_ = _lhs21
				_156_v20 = _rhs30
				_lhs21.F8 = _rhs31
				_131_v3 = _131_v3
				var _157_v21 _dafny.Array
				_ = _157_v21
				var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
				_ = _nw23
				_157_v21 = _nw23
				var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_157_v21), 0))
				_ = _index22
				(_157_v21).ArraySet1(_130_v2, (_index22).Int())
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(128), _dafny.ArrayLen((_157_v21), 0))
				_ = _index23
				(_157_v21).ArraySet1(_130_v2, (_index23).Int())
				_134_v5 = _134_v5
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	var _158_v22 _dafny.CodePoint
	_ = _158_v22
	_158_v22 = _dafny.CodePoint('q')
	var _159_v23 D0
	_ = _159_v23
	_159_v23 = Companion_D0_.Create_DC0_(_158_v22)
	_159_v23 = _159_v23
	_140_v10 = _134_v5
	var _160_v24 D0
	_ = _160_v24
	_160_v24 = Companion_D0_.Create_DC1_(_140_v10, _128_v0, true, _128_v0)
	var _161_i7 _dafny.Int
	_ = _161_i7
	_161_i7 = _dafny.Zero
	{
		var _pat_let_tv1 = _140_v10
		_ = _pat_let_tv1
		var _pat_let_tv2 = _134_v5
		_ = _pat_let_tv2
		for func(_source3 D0) bool {
			if _source3.Is_DC1() {
				var _169___mcc_h0 bool = _source3.Get_().(D0_DC1).Cf1
				_ = _169___mcc_h0
				var _170___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
				_ = _170___mcc_h1
				var _171___mcc_h2 bool = _source3.Get_().(D0_DC1).Cf3
				_ = _171___mcc_h2
				var _172___mcc_h3 _dafny.Int = _source3.Get_().(D0_DC1).Cf4
				_ = _172___mcc_h3
				var _173_cf4 _dafny.Int = _172___mcc_h3
				_ = _173_cf4
				var _174_cf3 bool = _171___mcc_h2
				_ = _174_cf3
				var _175_cf2 _dafny.Int = _170___mcc_h1
				_ = _175_cf2
				var _176_cf1 bool = _169___mcc_h0
				_ = _176_cf1
				return _pat_let_tv1
			} else {
				var _177___mcc_h4 _dafny.CodePoint = _source3.Get_().(D0_DC0).Cf0
				_ = _177___mcc_h4
				var _178_cf0 _dafny.CodePoint = _177___mcc_h4
				_ = _178_cf0
				return _pat_let_tv2
			}
		}(_160_v24) {
			{
				if (_161_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L5
				}
				_161_i7 = (_161_i7).Plus(_dafny.One)
				(_133_globalState).F6 = _128_v0
				var _162_v25 *C0
				_ = _162_v25
				var _nw24 *C0 = New_C0_()
				_ = _nw24
				_nw24.Ctor__()
				_162_v25 = _nw24
				var _163_v26 _dafny.Map
				_ = _163_v26
				_163_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v0, _162_v25)
				(_133_globalState).F8 = ((_163_v26).Merge(_163_v26)).Contains(_128_v0)
				if _134_v5 {
					(_133_globalState).F8 = _140_v10
					var _164_v27 _dafny.Array
					_ = _164_v27
					var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
					_ = _nw25
					_164_v27 = _nw25
					var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_164_v27), 0))
					_ = _index24
					(_164_v27).ArraySet1(_131_v3, (_index24).Int())
					var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_164_v27), 0))
					_ = _index25
					var _rhs32 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(_128_v0, _140_v10, _133_globalState), (Companion_Default___.SafeIndex(_128_v0, _dafny.IntOfUint32((Companion_Default___.Fm1(_128_v0, _140_v10, _133_globalState)).Cardinality()))).Uint32(), Companion_Default___.Fm2(_142_v12, _134_v5, _140_v10, _133_globalState))
					_ = _rhs32
					var _rhs33 bool = _140_v10
					_ = _rhs33
					var _rhs34 _dafny.Int = ((_128_v0).Minus(_128_v0)).Minus(_dafny.IntOfUint32((_131_v3).Cardinality()))
					_ = _rhs34
					var _lhs22 _dafny.Array = _164_v27
					_ = _lhs22
					var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_164_v27), 0))
					_ = _lhs23
					var _lhs24 *GlobalState = _133_globalState
					_ = _lhs24
					var _lhs25 *GlobalState = _133_globalState
					_ = _lhs25
					(_lhs22).ArraySet1(_rhs32, (_lhs23).Int())
					_lhs24.F8 = _rhs33
					_lhs25.F3 = _rhs34
					_134_v5 = _140_v10
					var _165_v28 _dafny.Array
					_ = _165_v28
					var _nwElement0_3 D0 = _159_v23
					_ = _nwElement0_3
					var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(15))
					_ = _nw26
					(_nw26).ArraySet1(_nwElement0_3, 0)
					(_nw26).ArraySet1(Companion_D0_.Create_DC0_(_158_v22), 1)
					(_nw26).ArraySet1(_159_v23, 2)
					(_nw26).ArraySet1(_159_v23, 3)
					(_nw26).ArraySet1(_159_v23, 4)
					(_nw26).ArraySet1(_159_v23, 5)
					(_nw26).ArraySet1(_159_v23, 6)
					(_nw26).ArraySet1(_159_v23, 7)
					(_nw26).ArraySet1(Companion_D0_.Create_DC0_(_158_v22), 8)
					(_nw26).ArraySet1(_159_v23, 9)
					(_nw26).ArraySet1(Companion_D0_.Create_DC0_(_158_v22), 10)
					(_nw26).ArraySet1(_159_v23, 11)
					(_nw26).ArraySet1(Companion_D0_.Create_DC0_(_dafny.CodePoint('c')), 12)
					(_nw26).ArraySet1(_159_v23, 13)
					(_nw26).ArraySet1(_159_v23, 14)
					_165_v28 = _nw26
					var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_165_v28), 0))
					_ = _index26
					(_165_v28).ArraySet1(Companion_D0_.Create_DC0_(_158_v22), (_index26).Int())
					var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_164_v27), 0))
					_ = _index27
					var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_165_v28), 0))
					_ = _index28
					var _rhs35 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_131_v3, _131_v3)
					_ = _rhs35
					var _rhs36 D0 = _159_v23
					_ = _rhs36
					var _lhs26 _dafny.Array = _164_v27
					_ = _lhs26
					var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_164_v27), 0))
					_ = _lhs27
					var _lhs28 _dafny.Array = _165_v28
					_ = _lhs28
					var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_165_v28), 0))
					_ = _lhs29
					(_lhs26).ArraySet1(_rhs35, (_lhs27).Int())
					(_lhs28).ArraySet1(_rhs36, (_lhs29).Int())
					(_133_globalState).F7 = _128_v0
				} else {
					_160_v24 = _160_v24
					var _166_v29 *C0
					_ = _166_v29
					var _nw27 *C0 = New_C0_()
					_ = _nw27
					_nw27.Ctor__()
					_166_v29 = _nw27
					_158_v22 = _158_v22
					var _167_v30 _dafny.Sequence
					_ = _167_v30
					_167_v30 = _dafny.SeqOf(_134_v5)
					(_133_globalState).F8 = !_dafny.Companion_Sequence_.Equal(_167_v30, (_137_v8).Fm0(_140_v10, _131_v3, _133_globalState))
					var _168_v31 *C0
					_ = _168_v31
					var _nw28 *C0 = New_C0_()
					_ = _nw28
					_nw28.Ctor__()
					_168_v31 = _nw28
				}
				goto C5
			}
		C5:
		}
		goto L5
	}
L5:
	var _179_v32 _dafny.Set
	_ = _179_v32
	_179_v32 = _dafny.SetOf(_134_v5)
	var _180_v33 _dafny.Sequence
	_ = _180_v33
	_180_v33 = _dafny.SeqOf(_dafny.SetOf(_134_v5, _140_v10, _140_v10, _140_v10, _140_v10), _179_v32, _179_v32)
	if !_dafny.Companion_Sequence_.Equal(_180_v33, _dafny.Companion_Sequence_.Concatenate(_180_v33, _dafny.SeqOf(_179_v32))) {
		var _181_v34 _dafny.Map
		_ = _181_v34
		_181_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v3, _140_v10)
		_181_v34 = _181_v34
		if !((_179_v32).IsSubsetOf(Companion_Default___.Fm3(_140_v10, _140_v10, _160_v24, _134_v5, _133_globalState))) {
			_179_v32 = _179_v32
			var _182_v35 _dafny.Sequence
			_ = _182_v35
			_182_v35 = _dafny.SeqOf(_128_v0, (_dafny.Zero).Minus(_128_v0), _128_v0)
			(_133_globalState).F8 = Companion_Default___.Fm4(_142_v12, _182_v35, false, _134_v5, _133_globalState)
			var _183_v37 _dafny.Array
			_ = _183_v37
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_5
			var _nw29 _dafny.Array
			_ = _nw29
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw29 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Map = (func(_184_v5 bool, _185_v0 _dafny.Int, _186_v10 bool, _187_v12 _dafny.Sequence, _188_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
					return func(_189_i8 _dafny.Int) _dafny.Map {
						return ((Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v5, _185_v0))).Dtor_cf5()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v10, (func() _dafny.Map {
							var _coll1 = _dafny.NewMapBuilder()
							_ = _coll1
							for _iter2 := _dafny.Iterate((_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_187_v12, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_185_v0), _dafny.IntOfUint32((_187_v12).Cardinality()))).Uint32(), _188_v22), _dafny.SeqOf(_188_v22))).Elements()); ; {
								_compr_1, _ok2 := _iter2()
								if !_ok2 {
									break
								}
								var _190_v36 _dafny.Sequence
								_190_v36 = interface{}(_compr_1).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_187_v12, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_185_v0), _dafny.IntOfUint32((_187_v12).Cardinality()))).Uint32(), _188_v22), _dafny.SeqOf(_188_v22)), _190_v36) {
									_coll1.Add(_190_v36, Companion_D0_.Create_DC1_(!(_184_v5), _185_v0, _184_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v22, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality()))
								}
							}
							return _coll1.ToMap()
						}()).Cardinality()))
					}
				})(_134_v5, _128_v0, _140_v10, _142_v12, _158_v22)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw29 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw29).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw29).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_183_v37 = _nw29
			_183_v37 = _183_v37
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_138_v9), 0))
			_ = _index29
			(_138_v9).ArraySet1(_140_v10, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_138_v9), 0))
			_ = _index30
			(_138_v9).ArraySet1(_134_v5, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_130_v2), 0))
			_ = _index31
			(_130_v2).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ngxstq"), _142_v12), (Companion_Default___.SafeIndex(_128_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ngxstq"), _142_v12)).Cardinality()))).Uint32(), _158_v22)).Cardinality()), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(833), _dafny.ArrayLen((_130_v2), 0))
			_ = _index32
			(_130_v2).ArraySet1((_dafny.Zero).Minus(_128_v0), (_index32).Int())
		} else {
			var _191_v38 _dafny.Array
			_ = _191_v38
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_6
			var _nw30 _dafny.Array
			_ = _nw30
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw30 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.CodePoint = (func(_192_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_193_i9 _dafny.Int) _dafny.CodePoint {
						return _192_v22
					}
				})(_158_v22)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw30 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw30).ArraySet1CodePoint(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw30).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_191_v38 = _nw30
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_191_v38), 0))
			_ = _index33
			(_191_v38).ArraySet1CodePoint(_158_v22, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_191_v38), 0))
			_ = _index34
			(_191_v38).ArraySet1CodePoint(_158_v22, (_index34).Int())
			(_133_globalState).F6 = (func() _dafny.Int {
				if true {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm2(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}((func(_194_v38 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
						return func(_195_i10 _dafny.Int) _dafny.CodePoint {
							return (_194_v38).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_194_v38), 0))).Int())
						}
					})(_191_v38))), _140_v10, (func() bool {
						if (_181_v34).Contains(_142_v12) {
							return (_181_v34).Get(_142_v12).(bool)
						}
						return Companion_Default___.Fm4(_142_v12, _dafny.SeqOf(_128_v0), _140_v10, true, _133_globalState)
					})(), _133_globalState))).Cardinality()))
				}
				return _128_v0
			})()
			var _196_v39 _dafny.Sequence
			_ = _196_v39
			_196_v39 = _dafny.SeqOf(_131_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_197_v38 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_198_i11 _dafny.Int) _dafny.CodePoint {
					return (_197_v38).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_197_v38), 0))).Int())
				}
			})(_191_v38))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-219))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_199_v38 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_200_i12 _dafny.Int) _dafny.CodePoint {
					return (_199_v38).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_199_v38), 0))).Int())
				}
			})(_191_v38)))))
			_196_v39 = _dafny.SeqOf(_131_v3, _131_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(279))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_201_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_202_i13 _dafny.Int) _dafny.CodePoint {
					return _201_v22
				}
			})(_158_v22))), _dafny.UnicodeSeqOfUtf8Bytes("ngtw")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_203_v38 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_204_i14 _dafny.Int) _dafny.CodePoint {
					return (_203_v38).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_203_v38), 0))).Int())
				}
			})(_191_v38))), _142_v12)
			var _205_v40 _dafny.Sequence
			_ = _205_v40
			_205_v40 = _dafny.SeqOf(!(_134_v5))
			(_133_globalState).F8 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_205_v40, _205_v40), !(_179_v32).Equals(_dafny.SetOf(!(_140_v10))))
			var _206_v41 _dafny.Array
			_ = _206_v41
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(2))
			_ = _nw31
			_206_v41 = _nw31
			var _207_v42 _dafny.MultiSet
			_ = _207_v42
			_207_v42 = _dafny.MultiSetOf(_129_v1)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_206_v41), 0))
			_ = _index35
			(_206_v41).ArraySet1((_207_v42).Difference(_dafny.MultiSetOf(_129_v1, _129_v1)), (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_206_v41), 0))
			_ = _index36
			(_206_v41).ArraySet1(Companion_Default___.Fm5(_dafny.IntOfInt64(330), !(_179_v32).Contains(_140_v10), (_134_v5) && (true), _133_globalState), (_index36).Int())
		}
		var _208_v43 bool
		_ = _208_v43
		var _209_v44 _dafny.Set
		_ = _209_v44
		var _210_v45 _dafny.Sequence
		_ = _210_v45
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Set
		_ = _out7
		var _out8 _dafny.Sequence
		_ = _out8
		_out6, _out7, _out8 = Companion_Default___.M0(_133_globalState)
		_208_v43 = _out6
		_209_v44 = _out7
		_210_v45 = _out8
		_208_v43 = _134_v5
		var _211_v46 _dafny.Array
		_ = _211_v46
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(12))
		_ = _nw32
		_211_v46 = _nw32
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_211_v46), 0))
		_ = _index37
		(_211_v46).ArraySet1(_160_v24, (_index37).Int())
		var _212_v47 _dafny.Sequence
		_ = _212_v47
		_212_v47 = _dafny.SeqOf(_128_v0)
		var _213_v48 _dafny.Sequence
		_ = _213_v48
		_213_v48 = _dafny.SeqOf(Companion_Default___.Fm4(_dafny.UnicodeSeqOfUtf8Bytes("okbhypfhd"), _212_v47, _134_v5, _134_v5, _133_globalState))
		var _214_v49 _dafny.Sequence
		_ = _214_v49
		_214_v49 = _dafny.SeqOf(_138_v9, _138_v9)
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_211_v46), 0))
		_ = _index38
		var _rhs37 _dafny.Int = ((_dafny.IntOfUint32((_210_v45).Cardinality())).Times((_dafny.MultiSetFromSeq(_213_v48)).Cardinality())).Minus((func() _dafny.Int {
			if (_129_v1).Contains(_128_v0) {
				return (_129_v1).Multiplicity(_128_v0)
			}
			return Companion_Default___.Fm6(_dafny.IntOfInt64(899), _134_v5, _133_globalState)
		})())
		_ = _rhs37
		var _rhs38 _dafny.CodePoint = _dafny.CodePoint('t')
		_ = _rhs38
		var _rhs39 bool = !(_208_v43) || (_208_v43)
		_ = _rhs39
		var _rhs40 D0 = Companion_D0_.Create_DC1_(_208_v43, _128_v0, _134_v5, _128_v0)
		_ = _rhs40
		var _rhs41 _dafny.Array = (_214_v49).Select((Companion_Default___.SafeIndex(_128_v0, _dafny.IntOfUint32((_214_v49).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs41
		var _lhs30 *GlobalState = _133_globalState
		_ = _lhs30
		var _lhs31 _dafny.Array = _211_v46
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(27), _dafny.ArrayLen((_211_v46), 0))
		_ = _lhs32
		_lhs30.F7 = _rhs37
		_158_v22 = _rhs38
		_134_v5 = _rhs39
		(_lhs31).ArraySet1(_rhs40, (_lhs32).Int())
		_138_v9 = _rhs41
	} else {
		(_133_globalState).F6 = (Companion_Default___.Fm6(Companion_Default___.Fm6(_128_v0, _140_v10, _133_globalState), _140_v10, _133_globalState)).Plus((func() _dafny.Int {
			if _134_v5 {
				return _128_v0
			}
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nowxxf")).Cardinality())
		})())
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_138_v9), 0))
		_ = _index39
		(_138_v9).ArraySet1(!(_140_v10) || (_134_v5), (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_138_v9), 0))
		_ = _index40
		(_138_v9).ArraySet1(false, (_index40).Int())
		var _215_v50 _dafny.Array
		_ = _215_v50
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_7
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Sequence = (func(_216_v9 _dafny.Array, _217_v5 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_218_i15 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_216_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_216_v9), 0))).Int()).(bool), (_216_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_216_v9), 0))).Int()).(bool), _217_v5)
				}
			})(_138_v9, _134_v5)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw33 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw33).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw33).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_215_v50 = _nw33
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_215_v50), 0))
		_ = _index41
		(_215_v50).ArraySet1(_dafny.SeqOf(_140_v10), (_index41).Int())
		var _219_v51 _dafny.Sequence
		_ = _219_v51
		_219_v51 = _dafny.SeqOf(_140_v10)
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_215_v50), 0))
		_ = _index42
		(_215_v50).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_219_v51, _219_v51), (_index42).Int())
		var _220_v52 _dafny.Sequence
		_ = _220_v52
		_220_v52 = _dafny.SeqOf(_137_v8, _137_v8)
		_220_v52 = _220_v52
		(_133_globalState).F3 = (_dafny.Zero).Minus(_128_v0)
	}
	var _221_v53 _dafny.Array
	_ = _221_v53
	var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
	_ = _nw34
	_221_v53 = _nw34
	var _222_v54 _dafny.Array
	_ = _222_v54
	var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
	_ = _nw35
	_222_v54 = _nw35
	var _223_v55 _dafny.Sequence
	_ = _223_v55
	_223_v55 = _dafny.SeqOf(_137_v8)
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_222_v54), 0))
	_ = _index43
	(_222_v54).ArraySet1((func() _dafny.Sequence {
		if _140_v10 {
			return _223_v55
		}
		return _223_v55
	})(), (_index43).Int())
	var _224_v56 _dafny.Map
	_ = _224_v56
	_224_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_158_v22, _158_v22)).Cardinality()), _128_v0)
	var _225_v57 _dafny.Sequence
	_ = _225_v57
	_225_v57 = _dafny.SeqOf(_128_v0)
	var _pat_let_tv3 = _140_v10
	_ = _pat_let_tv3
	var _pat_let_tv4 = _141_v11
	_ = _pat_let_tv4
	var _pat_let_tv5 = _141_v11
	_ = _pat_let_tv5
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_222_v54), 0))
	_ = _index44
	var _rhs42 bool = func(_source4 D1) bool {
		if _source4.Is_DC3() {
			var _226___mcc_h5 _dafny.Int = _source4.Get_().(D1_DC3).Cf6
			_ = _226___mcc_h5
			var _227___mcc_h6 bool = _source4.Get_().(D1_DC3).Cf7
			_ = _227___mcc_h6
			var _228_cf7 bool = _227___mcc_h6
			_ = _228_cf7
			var _229_cf6 _dafny.Int = _226___mcc_h5
			_ = _229_cf6
			return _pat_let_tv3
		} else {
			var _230___mcc_h7 _dafny.Map = _source4.Get_().(D1_DC2).Cf5
			_ = _230___mcc_h7
			var _231_cf5 _dafny.Map = _230___mcc_h7
			_ = _231_cf5
			return (_pat_let_tv4).IsDisjointFrom(_pat_let_tv5)
		}
	}(Companion_Default___.Fm7(_134_v5, _134_v5, _133_globalState))
	_ = _rhs42
	var _rhs43 _dafny.Array = _221_v53
	_ = _rhs43
	var _rhs44 bool = _134_v5
	_ = _rhs44
	var _rhs45 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_223_v55, _223_v55), _223_v55), (Companion_Default___.SafeIndex(_128_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_223_v55, _223_v55), _223_v55)).Cardinality()))).Uint32(), _137_v8)
	_ = _rhs45
	var _rhs46 bool = (func() bool {
		if Companion_Default___.Fm4(Companion_Default___.Fm1(_128_v0, _140_v10, _133_globalState), Companion_Default___.Fm8(_140_v10, (_224_v56).Update(_dafny.IntOfInt64(316), _128_v0), _128_v0, _133_globalState), _134_v5, Companion_Default___.Fm4(_142_v12, _225_v57, _140_v10, false, _133_globalState), _133_globalState) {
			return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_128_v0), _225_v57)
		}
		return !(_134_v5)
	})()
	_ = _rhs46
	var _lhs33 _dafny.Array = _222_v54
	_ = _lhs33
	var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_222_v54), 0))
	_ = _lhs34
	_134_v5 = _rhs42
	_221_v53 = _rhs43
	_140_v10 = _rhs44
	(_lhs33).ArraySet1(_rhs45, (_lhs34).Int())
	_134_v5 = _rhs46
	var _232_v58 bool
	_ = _232_v58
	var _233_v59 _dafny.Set
	_ = _233_v59
	var _234_v60 _dafny.Sequence
	_ = _234_v60
	var _out9 bool
	_ = _out9
	var _out10 _dafny.Set
	_ = _out10
	var _out11 _dafny.Sequence
	_ = _out11
	_out9, _out10, _out11 = Companion_Default___.M0(_133_globalState)
	_232_v58 = _out9
	_233_v59 = _out10
	_234_v60 = _out11
	_128_v0 = (_dafny.Zero).Minus(_128_v0)
	var _235_v61 *C0
	_ = _235_v61
	var _nw36 *C0 = New_C0_()
	_ = _nw36
	_nw36.Ctor__()
	_235_v61 = _nw36
	var _236_v62 _dafny.MultiSet
	_ = _236_v62
	_236_v62 = _dafny.MultiSetOf(_232_v58)
	var _237_v63 D1
	_ = _237_v63
	_237_v63 = Companion_D1_.Create_DC3_((_236_v62).Cardinality(), _232_v58)
	var _238_v64 _dafny.Array
	_ = _238_v64
	var _nwElement0_4 _dafny.Set = _179_v32
	_ = _nwElement0_4
	var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(6))
	_ = _nw37
	(_nw37).ArraySet1(_nwElement0_4, 0)
	(_nw37).ArraySet1(_179_v32, 1)
	(_nw37).ArraySet1((_179_v32).Intersection((_180_v33).Select((Companion_Default___.SafeIndex(_128_v0, _dafny.IntOfUint32((_180_v33).Cardinality()))).Uint32()).(_dafny.Set)), 2)
	(_nw37).ArraySet1(_179_v32, 3)
	(_nw37).ArraySet1(_179_v32, 4)
	(_nw37).ArraySet1((_179_v32).Difference(Companion_Default___.Fm3(_140_v10, !(_134_v5), _160_v24, (_237_v63).Dtor_cf7(), _133_globalState)), 5)
	_238_v64 = _nw37
	for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_238_v64), 0))); ; {
		_guard_loop_1, _ok3 := _iter3()
		if !_ok3 {
			break
		}
		var _239_i16 _dafny.Int
		_239_i16 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_239_i16).Sign() != -1) && ((_239_i16).Cmp(_dafny.ArrayLen((_238_v64), 0)) < 0)) {
			(_238_v64).ArraySet1(_179_v32, (_239_i16).Int())
		}
	}
	_dafny.Print(_128_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_129_v1).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-183))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v2).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_v3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_132_v4).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("rwijw"), _dafny.UnicodeSeqOfUtf8Bytes("cqq"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_133_globalState).F0()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-183))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_133_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_133_globalState).F10()).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("rwijw"), _dafny.UnicodeSeqOfUtf8Bytes("cqq"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_134_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-183), true), _dafny.IntOfInt64(-183)).UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-421), false), _dafny.IntOfInt64(183))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-421), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v9).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_140_v10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v11).Equals(_dafny.SetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_142_v12.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_143_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_149_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_154_i6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_v22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v23).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v24).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v24).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v24).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v24).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_161_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v32).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_180_v33, _dafny.SeqOf(_dafny.SetOf(false), _dafny.SetOf(false), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_222_v54).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence)).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_223_v55).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v56).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.IntOfInt64(-183))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_225_v57, _dafny.SeqOf(_dafny.IntOfInt64(-183))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_232_v58)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v59).Equals(_dafny.SetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_234_v60.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v62).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v63).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v63).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v64).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v64).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v64).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v64).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v64).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_238_v64).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
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
	Cf3 bool
	Cf4 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int, Cf3 bool, Cf4 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.CodePoint
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.CodePoint) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.Zero, false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf0() _dafny.CodePoint {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf6 _dafny.Int
	Cf7 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.Int, Cf7 bool) D1 {
	return D1{D1_DC3{Cf6, Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf5 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf5 _dafny.Map) D1 {
	return D1{D1_DC2{Cf5}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.Zero, false)
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
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

type D2_DC5 struct {
	Cf9  _dafny.Int
	Cf10 _dafny.Int
	Cf11 D0
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Int, Cf10 _dafny.Int, Cf11 D0) D2 {
	return D2{D2_DC5{Cf9, Cf10, Cf11}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC6 struct {
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_() D2 {
	return D2{D2_DC6{}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC4 struct {
	Cf8 _dafny.Array
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf8 _dafny.Array) D2 {
	return D2{D2_DC4{Cf8}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(_dafny.Zero, _dafny.Zero, Companion_D0_.Default())
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf11() D0 {
	return _this.Get_().(D2_DC5).Cf11
}

func (_this D2) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6"
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
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Equals(data2.Cf11)
		}
	case D2_DC6:
		{
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
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
	Cf13 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf13 _dafny.Int) D3 {
	return D3{D3_DC8{Cf13}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf12 _dafny.Map
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf12 _dafny.Map) D3 {
	return D3{D3_DC7{Cf12}}
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
	return Companion_D3_.Create_DC8_(_dafny.Zero)
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf12
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
			return "D3.DC8" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
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

// Definition of class GlobalState
type GlobalState struct {
	F3   _dafny.Int
	F6   _dafny.Int
	F7   _dafny.Int
	F8   bool
	_f0  _dafny.MultiSet
	_f1  _dafny.Array
	_f2  bool
	_f4  _dafny.Int
	_f5  _dafny.Int
	_f9  bool
	_f10 _dafny.Set
	_f11 bool
	_f12 _dafny.Int
	_f13 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F3 = _dafny.Zero
	_this.F6 = _dafny.Zero
	_this.F7 = _dafny.Zero
	_this.F8 = false
	_this._f0 = _dafny.EmptyMultiSet
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = false
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f9 = false
	_this._f10 = _dafny.EmptySet
	_this._f11 = false
	_this._f12 = _dafny.Zero
	_this._f13 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.MultiSet, f1 _dafny.Array, f2 bool, f3 _dafny.Int, f4 _dafny.Int, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.Int, f8 bool, f9 bool, f10 _dafny.Set, f11 bool, f12 _dafny.Int, f13 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *GlobalState) F0() _dafny.MultiSet {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Set {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
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
func (_this *C0) Fm0(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if false {
				return _dafny.SeqOf(true)
			}
			return _dafny.SeqOf(false)
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
