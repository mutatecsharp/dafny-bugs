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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(997)), _dafny.SeqOf(_dafny.IntOfInt64(-776), (_dafny.MultiSetOf(_dafny.IntOfInt64(-746), _dafny.IntOfInt64(499))).Cardinality(), _dafny.IntOfInt64(549)))).Cardinality()), _dafny.IntOfInt64(38))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, globalState *GlobalState) bool {
	return (Companion_D1_.Create_DC4_(!(false), false, true)).Dtor_cf6()
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(39), (_dafny.MultiSetOf(_dafny.IntOfInt64(633), (_dafny.Zero).Minus(_dafny.IntOfInt64(-268)))).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-431))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vjsfret"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bm"), _dafny.UnicodeSeqOfUtf8Bytes("pwkap")))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Sequence, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-185)
	}))).Cardinality()), _dafny.IntOfInt64(886))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tysocxty")).Cardinality()), (_dafny.MultiSetOf(false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false, !(false), !(false), false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false, true, !(true), true))))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(542), _dafny.IntOfInt64(145)), false)).Merge((Companion_D13_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(531), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pq")).Cardinality())), true))).Dtor_cf45())
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(413)).Cmp((_dafny.MultiSetOf(_dafny.IntOfInt64(603))).Cardinality()) <= 0, _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(_dafny.SetOf(_dafny.IntOfInt64(686)), _dafny.SetOf(_dafny.IntOfInt64(815)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-851))).Uint32(), func(coer2 func(_dafny.Int) D11) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) D11 {
		return Companion_D11_.Create_DC30_(false)
	}))).Cardinality())), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC5_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uvfc")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(((_dafny.MultiSetOf(false, !(false))).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("epbbayn")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ap")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm13(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v0 _dafny.CodePoint
			_4_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), _4_v0) {
				_coll0.Add(_4_v0, true)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Cmp(_dafny.IntOfInt64(587)) > 0, (func() bool {
		if false {
			return true
		}
		return true
	})(), (false) || (false))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('h'))
}
func (_static *CompanionStruct_Default___) Fm15(p0 D2, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_5_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
	}))).Cardinality()))).Cardinality())).Cmp(_dafny.IntOfInt64(231)) > 0 {
		return _dafny.CodePoint('k')
	} else {
		return _dafny.CodePoint('c')
	}
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("pyah"))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) D6 {
	if false {
		return Companion_D6_.Create_DC17_(_dafny.IntOfInt64(380), true)
	} else {
		return Companion_D6_.Create_DC17_(_dafny.IntOfInt64(-419), true)
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC13_(((_dafny.MultiSetOf(false)).Cardinality()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(854), _dafny.CodePoint('q'))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(338))).Cardinality())), Companion_D1_.Create_DC3_(_dafny.IntOfInt64(553)), Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())), Companion_D1_.Create_DC3_((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-257)))).Cardinality()), Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(295), _dafny.IntOfInt64(780))).Cardinality())))).Intersection(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(567))))).Intersection((_dafny.MultiSetOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(602)), Companion_D1_.Create_DC3_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("htapck")).Cardinality())))).Cardinality()), Companion_D1_.Create_DC3_(_dafny.IntOfInt64(-138)))).Intersection(_dafny.MultiSetOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(967)))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Map) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var _hi0 _dafny.Int = p0
	_ = _hi0
	for _6_i0 := Companion_Default___.SafeDivisionInt(p0, p0); _6_i0.Cmp(_hi0) < 0; _6_i0 = _6_i0.Plus(_dafny.One) {
		var _7_v0 bool
		_ = _7_v0
		_7_v0 = false
		var _8_v1 _dafny.Map
		_ = _8_v1
		_8_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-239), _7_v0)
		var _9_v2 D0
		_ = _9_v2
		_9_v2 = Companion_D0_.Create_DC0_(_7_v0)
		var _10_v3 _dafny.Set
		_ = _10_v3
		_10_v3 = _dafny.SetOf(_6_i0)
		var _11_v4 *C1
		_ = _11_v4
		var _nw0 *C1 = New_C1_()
		_ = _nw0
		_nw0.Ctor__(_7_v0, (p1).Minus(((_8_v1).Update(p0, (_9_v2).Dtor_cf0())).Cardinality()), p0, (_10_v3).IsProperSubsetOf(_dafny.SetOf(p0, p0)))
		_11_v4 = _nw0
		r0 = Companion_Default___.SafeModuloInt(p1, p0)
		var _12_v6 _dafny.Map
		_ = _12_v6
		_12_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(947), _dafny.IntOfInt64(518))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _13_v5 _dafny.Int
				_13_v5 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(947)).Cmp(_13_v5) <= 0) && ((_13_v5).Cmp(_dafny.IntOfInt64(518)) < 0) {
					_coll1.Add(Companion_Default___.SafeDivisionInt(_13_v5, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(491))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg6 _dafny.Int) interface{} {
							return coer6(arg6)
						}
					}(func(_14_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('v')
					}))).Cardinality())))
				}
			}
			return _coll1.ToSet()
		}())
		var _15_v7 _dafny.Map
		_ = _15_v7
		_15_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v4.F21, (_12_v6).Cardinality())
		var _16_v8 _dafny.Sequence
		_ = _16_v8
		_16_v8 = _dafny.SeqOf(_dafny.IntOfInt64(-998), p1)
		var _17_v9 D5
		_ = _17_v9
		_17_v9 = Companion_D5_.Create_DC13_((_16_v8).Select((Companion_Default___.SafeIndex(_6_i0, _dafny.IntOfUint32((_16_v8).Cardinality()))).Uint32()).(_dafny.Int))
		var _18_v10 _dafny.Set
		_ = _18_v10
		_18_v10 = _dafny.SetOf(_17_v9)
		var _19_v11 _dafny.Set
		_ = _19_v11
		_19_v11 = _dafny.SetOf(_18_v10)
		var _20_v12 _dafny.Map
		_ = _20_v12
		_20_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v7, _19_v11)
		var _21_v13 _dafny.MultiSet
		_ = _21_v13
		_21_v13 = _dafny.MultiSetOf(_11_v4.F21)
		_20_v12 = (_20_v12).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v0, (_21_v13).Cardinality()), _19_v11)
		var _22_v14 *C1
		_ = _22_v14
		var _nw1 *C1 = New_C1_()
		_ = _nw1
		_nw1.Ctor__(_7_v0, _dafny.IntOfInt64(350), p1, _7_v0)
		_22_v14 = _nw1
	}
	var _23_v15 bool
	_ = _23_v15
	_23_v15 = false
	var _24_v16 _dafny.Sequence
	_ = _24_v16
	_24_v16 = _dafny.SeqOf(_23_v15, _23_v15)
	if !((_24_v16).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_24_v16).Cardinality()))).Uint32()).(bool)) {
		var _25_v17 _dafny.CodePoint
		_ = _25_v17
		_25_v17 = _dafny.CodePoint('m')
		var _26_v18 _dafny.Map
		_ = _26_v18
		_26_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v15, _dafny.IntOfInt64(837))
		var _27_v19 _dafny.MultiSet
		_ = _27_v19
		_27_v19 = _dafny.MultiSetOf(p0, (func() _dafny.Int {
			if (_26_v18).Contains(_23_v15) {
				return (_26_v18).Get(_23_v15).(_dafny.Int)
			}
			return p1
		})())
		var _28_v20 D2
		_ = _28_v20
		_28_v20 = Companion_D2_.Create_DC6_(_27_v19)
		var _29_v21 D2
		_ = _29_v21
		_29_v21 = Companion_D2_.Create_DC7_(_28_v20)
		_25_v17 = Companion_Default___.Fm15(_29_v21, !(_23_v15), globalState)
		var _30_v22 _dafny.Array
		_ = _30_v22
		var _nwElement0_0 _dafny.CodePoint = _25_v17
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(2))
		_ = _nw2
		(_nw2).ArraySet1CodePoint(_nwElement0_0, 0)
		(_nw2).ArraySet1CodePoint(_25_v17, 1)
		_30_v22 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_30_v22), 0))
		_ = _index0
		(_30_v22).ArraySet1CodePoint(_dafny.CodePoint('a'), (_index0).Int())
		var _31_v23 _dafny.MultiSet
		_ = _31_v23
		_31_v23 = _dafny.MultiSetOf(_23_v15, _23_v15)
		var _32_v24 _dafny.Sequence
		_ = _32_v24
		_32_v24 = _dafny.SeqOf(_31_v23)
		var _33_v25 _dafny.Sequence
		_ = _33_v25
		_33_v25 = _dafny.SeqOf(p0)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_30_v22), 0))
		_ = _index1
		var _rhs0 bool = !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfUint32((_32_v24).Cardinality())), _33_v25)
		_ = _rhs0
		var _rhs1 bool = !(!(_31_v23).Contains(_23_v15))
		_ = _rhs1
		var _rhs2 _dafny.CodePoint = _25_v17
		_ = _rhs2
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 _dafny.Array = _30_v22
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(540), _dafny.ArrayLen((_30_v22), 0))
		_ = _lhs3
		_lhs0.F13 = _rhs0
		_lhs1.F13 = _rhs1
		(_lhs2).ArraySet1CodePoint(_rhs2, (_lhs3).Int())
		var _34_v26 T1
		_ = _34_v26
		var _nw3 *C1 = New_C1_()
		_ = _nw3
		_nw3.Ctor__(!(_23_v15), p0, p0, _23_v15)
		_34_v26 = _nw3
		var _35_v27 _dafny.Sequence
		_ = _35_v27
		_35_v27 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _36_v28 _dafny.MultiSet
		_ = _36_v28
		_36_v28 = _dafny.MultiSetOf(_35_v27)
		var _37_v29 D0
		_ = _37_v29
		_37_v29 = Companion_D0_.Create_DC1_((_34_v26).F19(), _dafny.IntOfUint32((_24_v16).Cardinality()), _23_v15)
		var _38_v30 _dafny.Map
		_ = _38_v30
		_38_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _35_v27)
		var _39_v31 _dafny.Array
		_ = _39_v31
		var _nwElement0_1 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(128))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_40_i2 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("nfnc")
		}))).Cardinality())
		_ = _nwElement0_1
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(21))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_1, 0)
		(_nw4).ArraySet1(_34_v26.F20(), 1)
		(_nw4).ArraySet1(_34_v26.F20(), 2)
		(_nw4).ArraySet1((_34_v26).F19(), 3)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_24_v16, _24_v16)).Cardinality()), 4)
		(_nw4).ArraySet1((_dafny.IntOfInt64(-199)).Minus(_dafny.IntOfUint32((_35_v27).Cardinality())), 5)
		(_nw4).ArraySet1(_34_v26.F20(), 6)
		(_nw4).ArraySet1(_34_v26.F20(), 7)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(579), (_34_v26).F19()), _33_v25)).Cardinality()), 8)
		(_nw4).ArraySet1(((_36_v28).Intersection(Companion_Default___.Fm16((_37_v29).Dtor_cf3(), globalState))).Cardinality(), 9)
		(_nw4).ArraySet1(_34_v26.F20(), 10)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_35_v27).Cardinality()), 11)
		(_nw4).ArraySet1((_38_v30).Cardinality(), 12)
		(_nw4).ArraySet1((_34_v26).F19(), 13)
		(_nw4).ArraySet1(_34_v26.F20(), 14)
		(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("b")).Cardinality()), 15)
		(_nw4).ArraySet1(p1, 16)
		(_nw4).ArraySet1(_34_v26.F20(), 17)
		(_nw4).ArraySet1(((Companion_Default___.Fm12(globalState)).Update(p1, Companion_Default___.Abs(p1))).Cardinality(), 18)
		(_nw4).ArraySet1(_dafny.IntOfInt64(-579), 19)
		(_nw4).ArraySet1(_34_v26.F20(), 20)
		_39_v31 = _nw4
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_39_v31), 0))
		_ = _index2
		(_39_v31).ArraySet1(_dafny.IntOfInt64(-751), (_index2).Int())
		var _41_v32 T0
		_ = _41_v32
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__(Companion_D2_.Create_DC6_(_27_v19), _dafny.IntOfInt64(-841), (_34_v26).F18())
		_41_v32 = _nw5
		var _42_v33 _dafny.Map
		_ = _42_v33
		_42_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v32, p1)
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_39_v31), 0))
		_ = _index3
		var _rhs3 bool = _23_v15
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _dafny.SeqOf((func() _dafny.Int {
			if (_42_v33).Contains(_41_v32) {
				return (_42_v33).Get(_41_v32).(_dafny.Int)
			}
			return p1
		})(), (_34_v26).F19(), (func() _dafny.Int {
			if _23_v15 {
				return p1
			}
			return _34_v26.F20()
		})(), (_34_v26).F19())
		_ = _rhs4
		var _rhs5 T1 = _34_v26
		_ = _rhs5
		var _rhs6 _dafny.Int = (p0).Plus(p0)
		_ = _rhs6
		var _lhs4 *GlobalState = globalState
		_ = _lhs4
		var _lhs5 _dafny.Array = _39_v31
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_39_v31), 0))
		_ = _lhs6
		_lhs4.F13 = _rhs3
		_33_v25 = _rhs4
		_34_v26 = _rhs5
		(_lhs5).ArraySet1(_rhs6, (_lhs6).Int())
		var _43_v34 _dafny.Array
		_ = _43_v34
		var _nw6 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
		_ = _nw6
		_43_v34 = _nw6
		var _44_v35 _dafny.Map
		_ = _44_v35
		_44_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_43_v34, (_34_v26).F18())
		_44_v35 = (_44_v35).Update(_43_v34, _23_v15)
		var _45_v36 _dafny.Set
		_ = _45_v36
		_45_v36 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_33_v25, (Companion_Default___.SafeIndex((_34_v26).F19(), _dafny.IntOfUint32((_33_v25).Cardinality()))).Uint32(), p1)).Cardinality()))
		(_34_v26).F20_set_((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_33_v25).Cardinality()), Companion_Default___.SafeDivisionInt(p1, (_45_v36).Cardinality())))))
	} else {
		(globalState).F13 = _23_v15
		var _46_v37 _dafny.Array
		_ = _46_v37
		var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
		_ = _nw7
		_46_v37 = _nw7
		var _47_v38 D6
		_ = _47_v38
		_47_v38 = Companion_D6_.Create_DC16_(p1)
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_46_v37), 0))
		_ = _index4
		(_46_v37).ArraySet1(_47_v38, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_46_v37), 0))
		_ = _index5
		(_46_v37).ArraySet1(Companion_D6_.Create_DC16_(p1), (_index5).Int())
		var _48_v39 _dafny.Array
		_ = _48_v39
		var _nw8 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
		_ = _nw8
		_48_v39 = _nw8
		var _49_v40 _dafny.MultiSet
		_ = _49_v40
		_49_v40 = _dafny.MultiSetOf(p1, p1)
		var _50_v41 D2
		_ = _50_v41
		_50_v41 = Companion_D2_.Create_DC6_(_49_v40)
		var _51_v42 T0
		_ = _51_v42
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__(_50_v41, p1, _23_v15)
		_51_v42 = _nw9
		var _52_v43 D6
		_ = _52_v43
		_52_v43 = Companion_D6_.Create_DC15_(_51_v42)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_48_v39), 0))
		_ = _index6
		(_48_v39).ArraySet1(_52_v43, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_48_v39), 0))
		_ = _index7
		(_48_v39).ArraySet1(_52_v43, (_index7).Int())
		(globalState).F13 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_24_v16, _24_v16), _24_v16)
		var _53_v44 *C0
		_ = _53_v44
		var _nw10 *C0 = New_C0_()
		_ = _nw10
		_nw10.Ctor__(Companion_D2_.Create_DC6_(_49_v40), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uhrsgbmu")).Cardinality()), _23_v15)
		_53_v44 = _nw10
		var _54_v45 _dafny.Sequence
		_ = _54_v45
		_54_v45 = _dafny.SeqOf(_53_v44)
		var _55_v46 _dafny.Array
		_ = _55_v46
		var _nwElement0_2 _dafny.Sequence = _54_v45
		_ = _nwElement0_2
		var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
		_ = _nw11
		(_nw11).ArraySet1(_nwElement0_2, 0)
		(_nw11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_54_v45, _54_v45), 1)
		(_nw11).ArraySet1(_54_v45, 2)
		(_nw11).ArraySet1(_54_v45, 3)
		_55_v46 = _nw11
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_55_v46), 0))
		_ = _index8
		(_55_v46).ArraySet1(_54_v45, (_index8).Int())
		var _56_v47 _dafny.Sequence
		_ = _56_v47
		_56_v47 = _dafny.SeqOf(_53_v44)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(530), _dafny.ArrayLen((_55_v46), 0))
		_ = _index9
		(_55_v46).ArraySet1(((func() _dafny.Sequence {
			if (_51_v42).F18() {
				return _56_v47
			}
			return _54_v45
		})()), (_index9).Int())
	}
	var _57_v48 *C1
	_ = _57_v48
	var _nw12 *C1 = New_C1_()
	_ = _nw12
	_nw12.Ctor__(_23_v15, p1, (_dafny.Zero).Minus(p1), _23_v15)
	_57_v48 = _nw12
	var _58_v49 D4
	_ = _58_v49
	_58_v49 = Companion_D4_.Create_DC9_(_57_v48)
	var _59_v50 _dafny.Sequence
	_ = _59_v50
	_59_v50 = _dafny.SeqOf((_58_v49).Dtor_cf13(), _57_v48)
	var _60_v51 _dafny.Array
	_ = _60_v51
	var _nwElement0_3 *C1 = _57_v48
	_ = _nwElement0_3
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(20))
	_ = _nw13
	(_nw13).ArraySet1(_nwElement0_3, 0)
	(_nw13).ArraySet1(_57_v48, 1)
	(_nw13).ArraySet1((_59_v50).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_59_v50).Cardinality()))).Uint32()).(*C1), 2)
	(_nw13).ArraySet1(_57_v48, 3)
	(_nw13).ArraySet1(_57_v48, 4)
	(_nw13).ArraySet1(_57_v48, 5)
	(_nw13).ArraySet1(_57_v48, 6)
	(_nw13).ArraySet1(_57_v48, 7)
	(_nw13).ArraySet1(_57_v48, 8)
	(_nw13).ArraySet1(_57_v48, 9)
	(_nw13).ArraySet1(_57_v48, 10)
	(_nw13).ArraySet1(_57_v48, 11)
	(_nw13).ArraySet1(_57_v48, 12)
	(_nw13).ArraySet1(_57_v48, 13)
	(_nw13).ArraySet1(_57_v48, 14)
	(_nw13).ArraySet1(_57_v48, 15)
	(_nw13).ArraySet1(_57_v48, 16)
	(_nw13).ArraySet1(_57_v48, 17)
	(_nw13).ArraySet1(_57_v48, 18)
	(_nw13).ArraySet1(_57_v48, 19)
	_60_v51 = _nw13
	var _61_v52 _dafny.Sequence
	_ = _61_v52
	_61_v52 = _dafny.SeqOf(_60_v51, _60_v51, _60_v51)
	var _62_v53 D6
	_ = _62_v53
	_62_v53 = Companion_D6_.Create_DC16_(p1)
	var _63_v54 _dafny.MultiSet
	_ = _63_v54
	_63_v54 = _dafny.MultiSetOf(p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(507))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}((func(_64_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_65_i3 _dafny.Int) _dafny.Int {
			return _64_p0
		}
	})(p0)))).Cardinality()))
	var _66_v55 D2
	_ = _66_v55
	_66_v55 = Companion_D2_.Create_DC6_(_63_v54)
	var _67_v56 *C0
	_ = _67_v56
	var _nw14 *C0 = New_C0_()
	_ = _nw14
	_nw14.Ctor__(_66_v55, p1, _23_v15)
	_67_v56 = _nw14
	var _68_v57 _dafny.Sequence
	_ = _68_v57
	_68_v57 = _dafny.SeqOf(_67_v56)
	var _69_v58 _dafny.Sequence
	_ = _69_v58
	_69_v58 = _dafny.SeqOf(_dafny.IntOfInt64(479), (_dafny.SetOf((_62_v53).Dtor_cf22(), _dafny.IntOfUint32((_68_v57).Cardinality()))).Cardinality())
	var _rhs7 _dafny.Array = (_61_v52).Select((Companion_Default___.SafeIndex(((_69_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.IntOfUint32((_69_v58).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p0), _dafny.IntOfUint32((_61_v52).Cardinality()))).Uint32()).(_dafny.Array)
	_ = _rhs7
	var _rhs8 _dafny.Int = p0
	_ = _rhs8
	var _rhs9 _dafny.Int = (_67_v56).F23()
	_ = _rhs9
	var _rhs10 bool = Companion_Default___.Fm1(p0, (_57_v48.F21) || (_23_v15), (_dafny.MultiSetOf(_57_v48.F21)).Intersection(_dafny.MultiSetOf(!(_23_v15))), globalState)
	_ = _rhs10
	var _lhs7 *GlobalState = globalState
	_ = _lhs7
	var _lhs8 *GlobalState = globalState
	_ = _lhs8
	var _lhs9 *GlobalState = globalState
	_ = _lhs9
	_60_v51 = _rhs7
	_lhs7.F9 = _rhs8
	_lhs8.F5 = _rhs9
	_lhs9.F13 = _rhs10
	if (Companion_Default___.SafeDivisionInt((_67_v56).F23(), (_dafny.Zero).Minus((Companion_Default___.Fm9(_57_v48.F21, true, globalState)).Cardinality()))).Cmp(_dafny.IntOfInt64(36)) >= 0 {
		var _70_v59 _dafny.Set
		_ = _70_v59
		_70_v59 = _dafny.SetOf(_23_v15, _57_v48.F21, _23_v15)
		var _71_v60 D9
		_ = _71_v60
		_71_v60 = Companion_D9_.Create_DC23_(_70_v59)
		var _pat_let_tv0 = _57_v48
		_ = _pat_let_tv0
		var _pat_let_tv1 = _57_v48
		_ = _pat_let_tv1
		_70_v59 = (_70_v59).Union((func(_pat_let0_0 D9) D9 {
			return func(_72_dt__update__tmp_h0 D9) D9 {
				return func(_pat_let1_0 _dafny.Set) D9 {
					return func(_73_dt__update_hcf32_h0 _dafny.Set) D9 {
						return Companion_D9_.Create_DC23_(_73_dt__update_hcf32_h0)
					}(_pat_let1_0)
				}(_dafny.SetOf(_pat_let_tv0.F21, _pat_let_tv1.F21))
			}(_pat_let0_0)
		}(_71_v60)).Dtor_cf32())
		(globalState).F3 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_dafny.IntOfInt64(929), p1, _dafny.IntOfInt64(503), p0, globalState), _dafny.IntOfInt64(263))
		var _74_v61 *C1
		_ = _74_v61
		var _nw15 *C1 = New_C1_()
		_ = _nw15
		_nw15.Ctor__(_57_v48.F21, p0, _dafny.IntOfInt64(537), !((_57_v48).Fm5(false, _dafny.IntOfInt64(-758), globalState)))
		_74_v61 = _nw15
		var _75_v62 _dafny.Map
		_ = _75_v62
		_75_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v56, _70_v59)
		var _76_v63 D1
		_ = _76_v63
		_76_v63 = Companion_D1_.Create_DC3_(((_75_v62).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_74_v61.F21, _57_v48.F21, _74_v61.F21)).Cardinality())))
		var _source0 D1 = _76_v63
		_ = _source0
		if _source0.Is_DC3() {
			var _77___mcc_h0 _dafny.Int = _source0.Get_().(D1_DC3).Cf5
			_ = _77___mcc_h0
			var _78_cf5 _dafny.Int = _77___mcc_h0
			_ = _78_cf5
			var _79_v64 _dafny.Map
			_ = _79_v64
			_79_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_67_v56).F23(), _74_v61.F21)
			_79_v64 = (_79_v64).Update(p0, _74_v61.F21)
			var _80_v65 _dafny.Array
			_ = _80_v65
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw16
			_80_v65 = _nw16
			var _81_v66 _dafny.Array
			_ = _81_v66
			var _nwElement0_4 _dafny.Int = (_67_v56).F23()
			_ = _nwElement0_4
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.One)
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_4, 0)
			_81_v66 = _nw17
			var _82_v67 _dafny.Map
			_ = _82_v67
			_82_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v65, _81_v66)
			var _83_v68 D7
			_ = _83_v68
			_83_v68 = Companion_D7_.Create_DC18_(_82_v67)
			_83_v68 = Companion_D7_.Create_DC18_((_82_v67).Merge(_82_v67))
			_23_v15 = !(((_67_v56).F23()).Cmp((_67_v56).F23()) >= 0)
			(_74_v61).F21 = !(_57_v48.F21)
		} else if _source0.Is_DC4() {
			var _84___mcc_h1 bool = _source0.Get_().(D1_DC4).Cf6
			_ = _84___mcc_h1
			var _85___mcc_h2 bool = _source0.Get_().(D1_DC4).Cf7
			_ = _85___mcc_h2
			var _86___mcc_h3 bool = _source0.Get_().(D1_DC4).Cf8
			_ = _86___mcc_h3
			var _87_cf8 bool = _86___mcc_h3
			_ = _87_cf8
			var _88_cf7 bool = _85___mcc_h2
			_ = _88_cf7
			var _89_cf6 bool = _84___mcc_h1
			_ = _89_cf6
			var _90_v70 _dafny.Array
			_ = _90_v70
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_0
			var _nw18 _dafny.Array
			_ = _nw18
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw18 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Map = (func(_91_cf7 bool, _92_v48 *C1, _93_v56 *C0) func(_dafny.Int) _dafny.Map {
					return func(_94_i4 _dafny.Int) _dafny.Map {
						return (func() _dafny.Map {
							if _91_cf7 {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nqlnui")).Cardinality()), _92_v48.F21)
							}
							return func() _dafny.Map {
								var _coll2 = _dafny.NewMapBuilder()
								_ = _coll2
								for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(418), _dafny.IntOfInt64(981))); ; {
									_compr_2, _ok2 := _iter2()
									if !_ok2 {
										break
									}
									var _95_v69 _dafny.Int
									_95_v69 = interface{}(_compr_2).(_dafny.Int)
									if ((_dafny.IntOfInt64(418)).Cmp(_95_v69) <= 0) && ((_95_v69).Cmp(_dafny.IntOfInt64(981)) < 0) {
										_coll2.Add((_95_v69).Minus((_93_v56).F23()), _92_v48.F21)
									}
								}
								return _coll2.ToMap()
							}()
						})()
					}
				})(_88_cf7, _57_v48, _67_v56)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw18 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw18).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw18).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_90_v70 = _nw18
			var _96_v71 _dafny.Map
			_ = _96_v71
			_96_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_67_v56).F23(), _57_v48.F21)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_90_v70), 0))
			_ = _index10
			(_90_v70).ArraySet1(_96_v71, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_90_v70), 0))
			_ = _index11
			(_90_v70).ArraySet1(_96_v71, (_index11).Int())
			var _97_v72 _dafny.Map
			_ = _97_v72
			_97_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, _57_v48.F21)
			var _98_v73 *C1
			_ = _98_v73
			var _nw19 *C1 = New_C1_()
			_ = _nw19
			_nw19.Ctor__(_57_v48.F21, (_dafny.IntOfInt64(615)).Times(p1), (_97_v72).Cardinality(), _57_v48.F21)
			_98_v73 = _nw19
			(globalState).F9 = (_67_v56).F23()
			var _99_v74 _dafny.Sequence
			_ = _99_v74
			_99_v74 = _dafny.UnicodeSeqOfUtf8Bytes("qaky")
			(globalState).F4 = _99_v74
		} else {
			var _100___mcc_h4 _dafny.Sequence = _source0.Get_().(D1_DC2).Cf4
			_ = _100___mcc_h4
			var _101_cf4 _dafny.Sequence = _100___mcc_h4
			_ = _101_cf4
			var _102_v75 _dafny.Int
			_ = _102_v75
			var _103_v76 bool
			_ = _103_v76
			var _104_v77 _dafny.Int
			_ = _104_v77
			var _105_v78 _dafny.Int
			_ = _105_v78
			var _out0 _dafny.Int
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			var _out3 _dafny.Int
			_ = _out3
			_out0, _out1, _out2, _out3 = (_57_v48).M1(_57_v48.F21, _74_v61.F21, globalState)
			_102_v75 = _out0
			_103_v76 = _out1
			_104_v77 = _out2
			_105_v78 = _out3
			var _106_v79 _dafny.Array
			_ = _106_v79
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw20
			_106_v79 = _nw20
			var _107_v80 _dafny.Sequence
			_ = _107_v80
			_107_v80 = _dafny.UnicodeSeqOfUtf8Bytes("p")
			var _108_v81 _dafny.Map
			_ = _108_v81
			_108_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v77, _dafny.IntOfUint32((_24_v16).Cardinality()))
			var _109_v82 _dafny.Map
			_ = _109_v82
			_109_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, _108_v81)
			var _110_v83 _dafny.Array
			_ = _110_v83
			var _nwElement0_5 _dafny.Int = (_67_v56).F23()
			_ = _nwElement0_5
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(17))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_5, 0)
			(_nw21).ArraySet1(p0, 1)
			(_nw21).ArraySet1((_67_v56).F23(), 2)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_107_v80).Cardinality()), 3)
			(_nw21).ArraySet1((_67_v56).F23(), 4)
			(_nw21).ArraySet1((_109_v82).Cardinality(), 5)
			(_nw21).ArraySet1((_dafny.Zero).Minus((_67_v56).F23()), 6)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_69_v58).Cardinality()), 7)
			(_nw21).ArraySet1((_67_v56).F23(), 8)
			(_nw21).ArraySet1(_104_v77, 9)
			(_nw21).ArraySet1((_67_v56).F23(), 10)
			(_nw21).ArraySet1(p0, 11)
			(_nw21).ArraySet1(p0, 12)
			(_nw21).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(330))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}(func(_111_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}))).Cardinality()), 13)
			(_nw21).ArraySet1(_104_v77, 14)
			(_nw21).ArraySet1((_dafny.Zero).Minus((_67_v56).F23()), 15)
			(_nw21).ArraySet1(_104_v77, 16)
			_110_v83 = _nw21
			var _112_v84 _dafny.Map
			_ = _112_v84
			_112_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v79, _110_v83)
			var _113_v85 D7
			_ = _113_v85
			_113_v85 = Companion_D7_.Create_DC18_(_112_v84)
			_113_v85 = _113_v85
			var _114_v86 D6
			_ = _114_v86
			_114_v86 = Companion_D6_.Create_DC17_((_67_v56).F23(), true)
			var _115_v87 _dafny.Set
			_ = _115_v87
			_115_v87 = _dafny.SetOf(_114_v86)
			var _116_v88 _dafny.Map
			_ = _116_v88
			_116_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_104_v77, _115_v87)
			_116_v88 = (_116_v88).Update(_104_v77, _dafny.SetOf(_114_v86, Companion_Default___.Fm17(globalState), Companion_D6_.Create_DC17_(p1, _23_v15)))
			var _117_v89 _dafny.Map
			_ = _117_v89
			_117_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_110_v83, _107_v80)
			(globalState).F5 = ((_117_v89).Merge((_117_v89).Update(_110_v83, _107_v80))).Cardinality()
		}
		r0 = (_67_v56).F23()
	} else {
		var _118_v90 _dafny.Array
		_ = _118_v90
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw22
		_118_v90 = _nw22
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))
		_ = _index12
		(_118_v90).ArraySet1(_dafny.IntOfInt64(250), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))
		_ = _index13
		(_118_v90).ArraySet1(_dafny.IntOfUint32((_24_v16).Cardinality()), (_index13).Int())
		_57_v48 = _57_v48
		var _119_v91 _dafny.Map
		_ = _119_v91
		_119_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_67_v56).F23(), p1)
		var _120_v93 D9
		_ = _120_v93
		_120_v93 = Companion_D9_.Create_DC24_(!(_119_v91).Equals(func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(665), _dafny.IntOfInt64(12))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _121_v92 _dafny.Int
				_121_v92 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(665)).Cmp(_121_v92) <= 0) && ((_121_v92).Cmp(_dafny.IntOfInt64(12)) < 0) {
					_coll3.Add((_121_v92).Minus(_dafny.IntOfInt64(330)), (_67_v56).F23())
				}
			}
			return _coll3.ToMap()
		}()), (_dafny.Zero).Minus((_67_v56).F23()), !(_57_v48.F21))
		var _source1 D9 = _120_v93
		_ = _source1
		if _source1.Is_DC24() {
			var _122___mcc_h5 bool = _source1.Get_().(D9_DC24).Cf33
			_ = _122___mcc_h5
			var _123___mcc_h6 _dafny.Int = _source1.Get_().(D9_DC24).Cf34
			_ = _123___mcc_h6
			var _124___mcc_h7 bool = _source1.Get_().(D9_DC24).Cf35
			_ = _124___mcc_h7
			var _125_cf35 bool = _124___mcc_h7
			_ = _125_cf35
			var _126_cf34 _dafny.Int = _123___mcc_h6
			_ = _126_cf34
			var _127_cf33 bool = _122___mcc_h5
			_ = _127_cf33
			var _128_v94 _dafny.Array
			_ = _128_v94
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_1
			var _nw23 _dafny.Array
			_ = _nw23
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw23 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_129_v48 *C1) func(_dafny.Int) bool {
					return func(_130_i6 _dafny.Int) bool {
						return _129_v48.F21
					}
				})(_57_v48)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw23 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw23).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw23).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_128_v94 = _nw23
			var _131_v95 _dafny.Map
			_ = _131_v95
			_131_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, _128_v94)
			var _132_v96 _dafny.Array
			_ = _132_v96
			var _nwElement0_6 _dafny.Array = _128_v94
			_ = _nwElement0_6
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(4))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_6, 0)
			(_nw24).ArraySet1(_128_v94, 1)
			(_nw24).ArraySet1((func() _dafny.Array {
				if (_131_v95).Contains(_57_v48.F21) {
					return (_131_v95).Get(_57_v48.F21).(_dafny.Array)
				}
				return _128_v94
			})(), 2)
			(_nw24).ArraySet1(_128_v94, 3)
			_132_v96 = _nw24
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_132_v96), 0))
			_ = _index14
			(_132_v96).ArraySet1(_128_v94, (_index14).Int())
			var _133_v97 _dafny.Sequence
			_ = _133_v97
			_133_v97 = _dafny.SeqOf((Companion_D10_.Create_DC25_(_128_v94)).Dtor_cf36())
			var _134_v98 _dafny.Set
			_ = _134_v98
			_134_v98 = _dafny.SetOf(_57_v48.F21)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_132_v96), 0))
			_ = _index15
			(_132_v96).ArraySet1((_133_v97).Select((Companion_Default___.SafeIndex((_134_v98).Cardinality(), _dafny.IntOfUint32((_133_v97).Cardinality()))).Uint32()).(_dafny.Array), (_index15).Int())
			var _135_v99 _dafny.Sequence
			_ = _135_v99
			_135_v99 = _dafny.UnicodeSeqOfUtf8Bytes("qggfdxe")
			var _136_v100 _dafny.MultiSet
			_ = _136_v100
			_136_v100 = _dafny.MultiSetOf(true)
			var _rhs11 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(663))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_137_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_138_i7 _dafny.Int) _dafny.Int {
					return _137_p0
				}
			})(p0)))
			_ = _rhs11
			var _rhs12 _dafny.Int = (_dafny.Zero).Minus(_126_cf34)
			_ = _rhs12
			var _rhs13 bool = (_24_v16).Select((Companion_Default___.SafeIndex(((_118_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfUint32((_135_v99).Cardinality())), _dafny.IntOfUint32((_24_v16).Cardinality()))).Uint32()).(bool)
			_ = _rhs13
			var _rhs14 _dafny.Int = (_118_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))).Int()).(_dafny.Int)
			_ = _rhs14
			var _rhs15 _dafny.Int = (func() _dafny.Int {
				if (_136_v100).Contains((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ursmty")).Cardinality())).Cmp((_67_v56).F23()) < 0) {
					return (_136_v100).Multiplicity((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ursmty")).Cardinality())).Cmp((_67_v56).F23()) < 0)
				}
				return ((_67_v56).F23()).Times((_67_v56).F23())
			})()
			_ = _rhs15
			var _lhs10 *GlobalState = globalState
			_ = _lhs10
			var _lhs11 *C1 = _57_v48
			_ = _lhs11
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			_69_v58 = _rhs11
			_lhs10.F3 = _rhs12
			_lhs11.F21 = _rhs13
			_lhs12.F3 = _rhs14
			_126_cf34 = _rhs15
			var _139_v101 _dafny.Map
			_ = _139_v101
			_139_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_cf33, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_67_v56).F23())).Cardinality(), _57_v48)).Cardinality())
			_125_cf35 = (((_139_v101).Cardinality()).Minus((_136_v100).Cardinality())).Cmp((_118_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))).Int()).(_dafny.Int)) <= 0
			var _140_v102 _dafny.Map
			_ = _140_v102
			_140_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))).Int()).(_dafny.Int), _135_v99)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))
			_ = _index16
			var _rhs16 _dafny.Array = _132_v96
			_ = _rhs16
			var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_133_v97, _133_v97)).Cardinality()), (_dafny.Zero).Minus(((_140_v102).Merge(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(410), _dafny.IntOfInt64(451))); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _141_v103 _dafny.Int
					_141_v103 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(410)).Cmp(_141_v103) <= 0) && ((_141_v103).Cmp(_dafny.IntOfInt64(451)) < 0) {
						_coll4.Add((_141_v103).Plus(_dafny.IntOfInt64(169)), _135_v99)
					}
				}
				return _coll4.ToMap()
			}())).Cardinality()))
			_ = _rhs17
			var _lhs13 _dafny.Array = _118_v90
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))
			_ = _lhs14
			_132_v96 = _rhs16
			(_lhs13).ArraySet1(_rhs17, (_lhs14).Int())
		} else {
			var _142___mcc_h8 _dafny.Set = _source1.Get_().(D9_DC23).Cf32
			_ = _142___mcc_h8
			var _143_cf32 _dafny.Set = _142___mcc_h8
			_ = _143_cf32
			var _144_v104 _dafny.Array
			_ = _144_v104
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_2
			var _nw25 _dafny.Array
			_ = _nw25
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw25 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Map = (func(_145_v56 *C0, _146_v48 *C1) func(_dafny.Int) _dafny.Map {
					return func(_147_i8 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC13_((_145_v56).F23()), _146_v48.F21)
					}
				})(_67_v56, _57_v48)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw25 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw25).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw25).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_144_v104 = _nw25
			var _148_v105 D5
			_ = _148_v105
			_148_v105 = Companion_D5_.Create_DC13_(p0)
			var _149_v106 _dafny.Map
			_ = _149_v106
			_149_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v105, _57_v48.F21)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_144_v104), 0))
			_ = _index17
			(_144_v104).ArraySet1(_149_v106, (_index17).Int())
			var _150_v108 _dafny.Sequence
			_ = _150_v108
			_150_v108 = _dafny.SeqOf(_148_v105, Companion_Default___.Fm18(_23_v15, globalState), _148_v105)
			var _pat_let_tv2 = _67_v56
			_ = _pat_let_tv2
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_144_v104), 0))
			_ = _index18
			var _rhs18 bool = _57_v48.F21
			_ = _rhs18
			var _rhs19 bool = _57_v48.F21
			_ = _rhs19
			var _rhs20 _dafny.Map = (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_150_v108).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _151_v107 D5
					_151_v107 = interface{}(_compr_5).(D5)
					if _dafny.Companion_Sequence_.Contains(_150_v108, _151_v107) {
						_coll5.Add(_151_v107, _23_v15)
					}
				}
				return _coll5.ToMap()
			}()).Update(func(_pat_let2_0 D5) D5 {
				return func(_152_dt__update__tmp_h1 D5) D5 {
					return func(_pat_let3_0 _dafny.Int) D5 {
						return func(_153_dt__update_hcf19_h0 _dafny.Int) D5 {
							return Companion_D5_.Create_DC13_(_153_dt__update_hcf19_h0)
						}(_pat_let3_0)
					}((_pat_let_tv2).F23())
				}(_pat_let2_0)
			}(_148_v105), _57_v48.F21)
			_ = _rhs20
			var _lhs15 *GlobalState = globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _144_v104
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_144_v104), 0))
			_ = _lhs17
			_lhs15.F13 = _rhs18
			_23_v15 = _rhs19
			(_lhs16).ArraySet1(_rhs20, (_lhs17).Int())
			var _nw26 *C1 = New_C1_()
			_ = _nw26
			_nw26.Ctor__(_57_v48.F21, (_67_v56).F23(), (_67_v56).F23(), (func() bool {
				if _57_v48.F21 {
					return _57_v48.F21
				}
				return _57_v48.F21
			})())
			_57_v48 = _nw26
			(globalState).F13 = _57_v48.F21
			(globalState).F9 = (_67_v56).F23()
		}
		var _154_v109 _dafny.Set
		_ = _154_v109
		_154_v109 = _dafny.SetOf(_67_v56, _67_v56, _67_v56)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))
		_ = _index19
		(_118_v90).ArraySet1((_dafny.Zero).Minus((_154_v109).Cardinality()), (_index19).Int())
		if _23_v15 {
			var _155_v110 _dafny.Map
			_ = _155_v110
			_155_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_67_v56).F23(), _57_v48)
			_155_v110 = (_155_v110).Merge(_155_v110)
			var _156_v111 _dafny.Map
			_ = _156_v111
			_156_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v15, p0)
			_156_v111 = _156_v111
			(globalState).F13 = _57_v48.F21
			var _157_v112 T1
			_ = _157_v112
			var _nw27 *C1 = New_C1_()
			_ = _nw27
			_nw27.Ctor__(_57_v48.F21, (_67_v56).F23(), (_67_v56).F23(), true)
			_157_v112 = _nw27
			var _158_v113 _dafny.CodePoint
			_ = _158_v113
			_158_v113 = _dafny.CodePoint('t')
			var _159_v114 _dafny.Map
			_ = _159_v114
			_159_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, _57_v48)
			var _160_v115 _dafny.MultiSet
			_ = _160_v115
			_160_v115 = _dafny.MultiSetOf(_118_v90, _118_v90)
			var _rhs21 T1 = _157_v112
			_ = _rhs21
			var _rhs22 bool = !(_57_v48.F21) || ((_157_v112).F18())
			_ = _rhs22
			var _rhs23 _dafny.CodePoint = _158_v113
			_ = _rhs23
			var _rhs24 *C1 = (func() *C1 {
				if (_159_v114).Contains((_dafny.MultiSetOf(_118_v90)).IsProperSubsetOf(_160_v115)) {
					return (_159_v114).Get((_dafny.MultiSetOf(_118_v90)).IsProperSubsetOf(_160_v115)).(*C1)
				}
				return _57_v48
			})()
			_ = _rhs24
			var _lhs18 *GlobalState = globalState
			_ = _lhs18
			_157_v112 = _rhs21
			_lhs18.F13 = _rhs22
			_158_v113 = _rhs23
			_57_v48 = _rhs24
			var _161_v116 _dafny.Array
			_ = _161_v116
			var _nwElement0_7 _dafny.CodePoint = _158_v113
			_ = _nwElement0_7
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(6))
			_ = _nw28
			(_nw28).ArraySet1CodePoint(_nwElement0_7, 0)
			(_nw28).ArraySet1CodePoint(_158_v113, 1)
			(_nw28).ArraySet1CodePoint(_158_v113, 2)
			(_nw28).ArraySet1CodePoint(_158_v113, 3)
			(_nw28).ArraySet1CodePoint(_158_v113, 4)
			(_nw28).ArraySet1CodePoint(_158_v113, 5)
			_161_v116 = _nw28
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_161_v116), 0))
			_ = _index20
			(_161_v116).ArraySet1CodePoint(_158_v113, (_index20).Int())
			var _162_v117 D2
			_ = _162_v117
			_162_v117 = Companion_D2_.Create_DC6_(_dafny.MultiSetOf(_dafny.IntOfInt64(789)))
			var _163_v118 D2
			_ = _163_v118
			_163_v118 = Companion_D2_.Create_DC7_(_162_v117)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(388), _dafny.ArrayLen((_161_v116), 0))
			_ = _index21
			(_161_v116).ArraySet1CodePoint(Companion_Default___.Fm15(_163_v118, (_57_v48).Fm5((_157_v112).F18(), Companion_Default___.Fm0(_dafny.IntOfInt64(760), p0, (_67_v56).F23(), _dafny.IntOfInt64(189), globalState), globalState), globalState), (_index21).Int())
		} else {
			var _164_v119 _dafny.CodePoint
			_ = _164_v119
			_164_v119 = _dafny.CodePoint('k')
			var _165_v120 D2
			_ = _165_v120
			_165_v120 = Companion_D2_.Create_DC6_(_63_v54)
			var _166_v121 D2
			_ = _166_v121
			_166_v121 = Companion_D2_.Create_DC7_(Companion_D2_.Create_DC7_(_165_v120))
			var _167_v122 _dafny.Array
			_ = _167_v122
			var _nwElement0_8 _dafny.CodePoint = _164_v119
			_ = _nwElement0_8
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(6))
			_ = _nw29
			(_nw29).ArraySet1CodePoint(_nwElement0_8, 0)
			(_nw29).ArraySet1CodePoint(_164_v119, 1)
			(_nw29).ArraySet1CodePoint(Companion_Default___.Fm15(_166_v121, _57_v48.F21, globalState), 2)
			(_nw29).ArraySet1CodePoint(_164_v119, 3)
			(_nw29).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _57_v48.F21 {
					return _164_v119
				}
				return _164_v119
			})(), 4)
			(_nw29).ArraySet1CodePoint(_164_v119, 5)
			_167_v122 = _nw29
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_167_v122), 0))
			_ = _index22
			(_167_v122).ArraySet1CodePoint(_164_v119, (_index22).Int())
			var _168_v123 _dafny.Map
			_ = _168_v123
			_168_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _57_v48)
			var _169_v124 T1
			_ = _169_v124
			var _nw30 *C1 = New_C1_()
			_ = _nw30
			_nw30.Ctor__((_57_v48).Fm5(_57_v48.F21, p1, globalState), (_dafny.MultiSetFromSeq(_24_v16)).Cardinality(), (_168_v123).Cardinality(), _57_v48.F21)
			_169_v124 = _nw30
			var _170_v125 D10
			_ = _170_v125
			_170_v125 = Companion_D10_.Create_DC26_(_169_v124)
			var _171_v126 _dafny.Array
			_ = _171_v126
			var _nwElement0_9 D10 = _170_v125
			_ = _nwElement0_9
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(4))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_9, 0)
			(_nw31).ArraySet1(Companion_D10_.Create_DC26_(_169_v124), 1)
			(_nw31).ArraySet1(_170_v125, 2)
			(_nw31).ArraySet1(Companion_D10_.Create_DC26_(_169_v124), 3)
			_171_v126 = _nw31
			var _172_v127 _dafny.Array
			_ = _172_v127
			var _nwElement0_10 _dafny.Array = _118_v90
			_ = _nwElement0_10
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(3))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_10, 0)
			(_nw32).ArraySet1(_118_v90, 1)
			(_nw32).ArraySet1(_118_v90, 2)
			_172_v127 = _nw32
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_172_v127), 0))
			_ = _index23
			(_172_v127).ArraySet1(_118_v90, (_index23).Int())
			var _173_v128 D11
			_ = _173_v128
			_173_v128 = Companion_D11_.Create_DC29_(_118_v90)
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_167_v122), 0))
			_ = _index24
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_172_v127), 0))
			_ = _index25
			var _rhs25 _dafny.Sequence = _69_v58
			_ = _rhs25
			var _rhs26 _dafny.CodePoint = _164_v119
			_ = _rhs26
			var _rhs27 _dafny.Array = _171_v126
			_ = _rhs27
			var _rhs28 _dafny.Array = (_173_v128).Dtor_cf40()
			_ = _rhs28
			var _lhs19 _dafny.Array = _167_v122
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_167_v122), 0))
			_ = _lhs20
			var _lhs21 _dafny.Array = _172_v127
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_172_v127), 0))
			_ = _lhs22
			_69_v58 = _rhs25
			(_lhs19).ArraySet1CodePoint(_rhs26, (_lhs20).Int())
			_171_v126 = _rhs27
			(_lhs21).ArraySet1(_rhs28, (_lhs22).Int())
			var _174_v129 _dafny.Int
			_ = _174_v129
			var _175_v130 bool
			_ = _175_v130
			var _176_v131 _dafny.Int
			_ = _176_v131
			var _177_v132 _dafny.Int
			_ = _177_v132
			var _out4 _dafny.Int
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out4, _out5, _out6, _out7 = (_169_v124).M1((_24_v16).Select((Companion_Default___.SafeIndex(_169_v124.F20(), _dafny.IntOfUint32((_24_v16).Cardinality()))).Uint32()).(bool), _23_v15, globalState)
			_174_v129 = _out4
			_175_v130 = _out5
			_176_v131 = _out6
			_177_v132 = _out7
			var _178_v133 _dafny.MultiSet
			_ = _178_v133
			_178_v133 = _dafny.MultiSetOf(_57_v48.F21, (_169_v124).F18())
			var _179_v134 _dafny.Sequence
			_ = _179_v134
			_179_v134 = _dafny.UnicodeSeqOfUtf8Bytes("rihyj")
			var _180_v135 _dafny.Map
			_ = _180_v135
			_180_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_178_v133).Cardinality(), _dafny.IntOfInt64(600), _dafny.IntOfInt64(-288), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_181_v119 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_182_i9 _dafny.Int) _dafny.CodePoint {
					return _181_v119
				}
			})(_164_v119)))).Cardinality()), globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_179_v134, (Companion_Default___.SafeIndex((_118_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(983), _dafny.ArrayLen((_118_v90), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_179_v134).Cardinality()))).Uint32(), (_167_v122).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_167_v122), 0))).Int())), _179_v134))
			_180_v135 = (_180_v135).Update((func() _dafny.Int {
				if (_178_v133).Contains(_57_v48.F21) {
					return (_178_v133).Multiplicity(_57_v48.F21)
				}
				return _176_v131
			})(), _179_v134)
			var _183_v136 T0
			_ = _183_v136
			var _nw33 *C0 = New_C0_()
			_ = _nw33
			_nw33.Ctor__(_67_v56.F22, p1, _175_v130)
			_183_v136 = _nw33
			var _184_v137 D6
			_ = _184_v137
			_184_v137 = Companion_D6_.Create_DC15_(_183_v136)
			var _185_v138 _dafny.Set
			_ = _185_v138
			_185_v138 = _dafny.SetOf((_184_v137).Dtor_cf21())
			var _186_v139 _dafny.Map
			_ = _186_v139
			_186_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v138).Cardinality(), _57_v48.F21)
			(globalState).F3 = ((_dafny.Zero).Minus((_69_v58).Select((Companion_Default___.SafeIndex((_186_v139).Cardinality(), _dafny.IntOfUint32((_69_v58).Cardinality()))).Uint32()).(_dafny.Int))).Plus(_176_v131)
			var _187_v140 _dafny.Array
			_ = _187_v140
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_3
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_188_v124 T1) func(_dafny.Int) bool {
					return func(_189_i10 _dafny.Int) bool {
						return !(!((_188_v124).F18()))
					}
				})(_169_v124)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw34 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw34).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw34).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_187_v140 = _nw34
			var _190_v141 _dafny.Map
			_ = _190_v141
			_190_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v140, (_119_v91).Cardinality())
			var _191_v142 _dafny.Map
			_ = _191_v142
			_191_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _190_v141)
			var _rhs29 bool = (_24_v16).Select((Companion_Default___.SafeIndex(_169_v124.F20(), _dafny.IntOfUint32((_24_v16).Cardinality()))).Uint32()).(bool)
			_ = _rhs29
			var _rhs30 _dafny.Sequence = _24_v16
			_ = _rhs30
			var _rhs31 bool = (_190_v141).Equals(((func() _dafny.Map {
				if (_191_v142).Contains(_57_v48.F21) {
					return (_191_v142).Get(_57_v48.F21).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v140, _177_v132)
			})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v140, Companion_Default___.Fm0(_176_v131, (_67_v56).F23(), _dafny.IntOfUint32((_24_v16).Cardinality()), _169_v124.F20(), globalState))))
			_ = _rhs31
			var _rhs32 bool = (_183_v136).F18()
			_ = _rhs32
			var _rhs33 _dafny.Int = (_174_v129).Minus((_67_v56).F23())
			_ = _rhs33
			var _lhs23 *C1 = _57_v48
			_ = _lhs23
			var _lhs24 *C1 = _57_v48
			_ = _lhs24
			var _lhs25 *GlobalState = globalState
			_ = _lhs25
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			_lhs23.F21 = _rhs29
			_24_v16 = _rhs30
			_lhs24.F21 = _rhs31
			_lhs25.F13 = _rhs32
			_lhs26.F9 = _rhs33
		}
	}
	var _192_v143 _dafny.Map
	_ = _192_v143
	_192_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(366))
	var _193_v144 _dafny.MultiSet
	_ = _193_v144
	_193_v144 = _dafny.MultiSetOf(_192_v143, _192_v143)
	var _194_v145 D7
	_ = _194_v145
	_194_v145 = Companion_D7_.Create_DC20_(_193_v144)
	var _pat_let_tv3 = _57_v48
	_ = _pat_let_tv3
	var _pat_let_tv4 = _193_v144
	_ = _pat_let_tv4
	var _pat_let_tv5 = _193_v144
	_ = _pat_let_tv5
	(_57_v48).F21 = func(_source2 D7) bool {
		if _source2.Is_DC19() {
			var _195___mcc_h9 bool = _source2.Get_().(D7_DC19).Cf26
			_ = _195___mcc_h9
			var _196___mcc_h10 _dafny.CodePoint = _source2.Get_().(D7_DC19).Cf27
			_ = _196___mcc_h10
			var _197___mcc_h11 bool = _source2.Get_().(D7_DC19).Cf28
			_ = _197___mcc_h11
			var _198_cf28 bool = _197___mcc_h11
			_ = _198_cf28
			var _199_cf27 _dafny.CodePoint = _196___mcc_h10
			_ = _199_cf27
			var _200_cf26 bool = _195___mcc_h9
			_ = _200_cf26
			return _198_cf28
		} else if _source2.Is_DC20() {
			var _201___mcc_h12 _dafny.MultiSet = _source2.Get_().(D7_DC20).Cf29
			_ = _201___mcc_h12
			var _202_cf29 _dafny.MultiSet = _201___mcc_h12
			_ = _202_cf29
			return true
		} else if _source2.Is_DC18() {
			var _203___mcc_h13 _dafny.Map = _source2.Get_().(D7_DC18).Cf25
			_ = _203___mcc_h13
			var _204_cf25 _dafny.Map = _203___mcc_h13
			_ = _204_cf25
			return _pat_let_tv3.F21
		} else {
			var _205___mcc_h14 D7 = _source2.Get_().(D7_DC21).Cf30
			_ = _205___mcc_h14
			var _206_cf30 D7 = _205___mcc_h14
			_ = _206_cf30
			return (_pat_let_tv4).IsSubsetOf(_pat_let_tv5)
		}
	}(_194_v145)
	var _207_v146 _dafny.Sequence
	_ = _207_v146
	_207_v146 = _dafny.UnicodeSeqOfUtf8Bytes("w")
	if !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ev"), _dafny.UnicodeSeqOfUtf8Bytes("xlhryx")), _207_v146)) {
		if _57_v48.F21 {
			var _208_v147 _dafny.Map
			_ = _208_v147
			_208_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_dafny.MultiSetOf(true)).Cardinality(), !(!(_57_v48.F21)), _dafny.MultiSetOf(_23_v15), globalState), _dafny.IntOfUint32((_69_v58).Cardinality()))
			var _209_v148 T1
			_ = _209_v148
			var _nw35 *C1 = New_C1_()
			_ = _nw35
			_nw35.Ctor__(_23_v15, p1, (_69_v58).Select((Companion_Default___.SafeIndex((_67_v56).F23(), _dafny.IntOfUint32((_69_v58).Cardinality()))).Uint32()).(_dafny.Int), false)
			_209_v148 = _nw35
			var _210_v149 D10
			_ = _210_v149
			_210_v149 = Companion_D10_.Create_DC26_(_209_v148)
			var _211_v150 _dafny.Map
			_ = _211_v150
			_211_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_208_v147, _210_v149)
			_211_v150 = (_211_v150).Update(_208_v147, _210_v149)
			var _212_v151 D4
			_ = _212_v151
			_212_v151 = Companion_D4_.Create_DC10_(_57_v48.F21, (func() _dafny.Int {
				if (_192_v143).Contains((_dafny.Zero).Minus(p1)) {
					return (_192_v143).Get((_dafny.Zero).Minus(p1)).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-796)
			})(), p1)
			var _213_v152 D1
			_ = _213_v152
			_213_v152 = Companion_D1_.Create_DC3_((_212_v151).Dtor_cf15())
			var _214_v153 _dafny.MultiSet
			_ = _214_v153
			_214_v153 = _dafny.MultiSetOf(_213_v152, _213_v152)
			var _215_v154 D9
			_ = _215_v154
			_215_v154 = Companion_D9_.Create_DC24_(false, (_209_v148).F19(), _57_v48.F21)
			var _rhs34 _dafny.Int = (func() _dafny.Int {
				if _57_v48.F21 {
					return ((_209_v148).F19()).Minus(p1)
				}
				return (_209_v148).F19()
			})()
			_ = _rhs34
			var _rhs35 bool = (_215_v154).Dtor_cf35()
			_ = _rhs35
			var _rhs36 _dafny.MultiSet = Companion_Default___.Fm19(_209_v148.F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_69_v58, _69_v58)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-197))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_216_v56 *C0) func(_dafny.Int) _dafny.Int {
				return func(_217_i11 _dafny.Int) _dafny.Int {
					return (_216_v56).F23()
				}
			})(_67_v56)))).Cardinality()), globalState)
			_ = _rhs36
			var _rhs37 bool = _57_v48.F21
			_ = _rhs37
			var _rhs38 bool = _57_v48.F21
			_ = _rhs38
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			var _lhs28 *C1 = _57_v48
			_ = _lhs28
			var _lhs29 *GlobalState = globalState
			_ = _lhs29
			var _lhs30 *C1 = _57_v48
			_ = _lhs30
			_lhs27.F3 = _rhs34
			_lhs28.F21 = _rhs35
			_214_v153 = _rhs36
			_lhs29.F13 = _rhs37
			_lhs30.F21 = _rhs38
			_57_v48 = _57_v48
			var _218_v155 _dafny.Array
			_ = _218_v155
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_4
			var _nw36 _dafny.Array
			_ = _nw36
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw36 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Int = func(_219_i12 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_219_i12, _dafny.IntOfInt64(110))
				}
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw36 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw36).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw36).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_218_v155 = _nw36
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_218_v155), 0))
			_ = _index26
			(_218_v155).ArraySet1(p0, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_218_v155), 0))
			_ = _index27
			(_218_v155).ArraySet1((_dafny.IntOfInt64(-685)).Plus(_209_v148.F20()), (_index27).Int())
			(globalState).F13 = (_209_v148).Fm5(_57_v48.F21, (_209_v148).F19(), globalState)
		} else {
			_57_v48 = _57_v48
			(globalState).F13 = _23_v15
			var _220_v156 _dafny.MultiSet
			_ = _220_v156
			_220_v156 = _dafny.MultiSetOf(_23_v15, _57_v48.F21)
			var _221_v157 *C1
			_ = _221_v157
			var _nw37 *C1 = New_C1_()
			_ = _nw37
			_nw37.Ctor__((_220_v156).Equals(_220_v156), (_dafny.Zero).Minus(((_67_v56).F23()).Times(_dafny.IntOfUint32((_207_v146).Cardinality()))), p0, (func() bool {
				if _57_v48.F21 {
					return _57_v48.F21
				}
				return _57_v48.F21
			})())
			_221_v157 = _nw37
			var _222_v158 _dafny.Array
			_ = _222_v158
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw38
			_222_v158 = _nw38
			_222_v158 = _222_v158
			var _223_v159 _dafny.CodePoint
			_ = _223_v159
			_223_v159 = _dafny.CodePoint('k')
			var _224_v160 _dafny.Set
			_ = _224_v160
			_224_v160 = _dafny.SetOf(_57_v48, _221_v157, _57_v48, _57_v48, _57_v48)
			_223_v159 = (func() _dafny.CodePoint {
				if _57_v48.F21 {
					return _223_v159
				}
				return (func() _dafny.CodePoint {
					if _221_v157.F21 {
						return (_207_v146).Select((Companion_Default___.SafeIndex((_224_v160).Cardinality(), _dafny.IntOfUint32((_207_v146).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _223_v159
				})()
			})()
		}
		(globalState).F9 = (func() _dafny.Int {
			if (_dafny.IntOfInt64(263)).Cmp(_dafny.IntOfUint32((_207_v146).Cardinality())) < 0 {
				return p0
			}
			return (_67_v56).F23()
		})()
		if (p1).Cmp(_dafny.IntOfInt64(914)) <= 0 {
			var _225_v161 _dafny.Int
			_ = _225_v161
			var _226_v162 bool
			_ = _226_v162
			var _227_v163 _dafny.Int
			_ = _227_v163
			var _228_v164 _dafny.Int
			_ = _228_v164
			var _out8 _dafny.Int
			_ = _out8
			var _out9 bool
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = (_57_v48).M1((_24_v16).Select((Companion_Default___.SafeIndex((_67_v56).F23(), _dafny.IntOfUint32((_24_v16).Cardinality()))).Uint32()).(bool), true, globalState)
			_225_v161 = _out8
			_226_v162 = _out9
			_227_v163 = _out10
			_228_v164 = _out11
			var _229_v165 _dafny.CodePoint
			_ = _229_v165
			_229_v165 = _dafny.CodePoint('x')
			_229_v165 = _229_v165
			var _230_v166 _dafny.Map
			_ = _230_v166
			_230_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, _dafny.IntOfUint32((_207_v146).Cardinality()))
			var _231_v167 _dafny.MultiSet
			_ = _231_v167
			_231_v167 = _dafny.MultiSetOf(_57_v48.F21, _23_v15)
			var _232_v168 *C1
			_ = _232_v168
			var _nw39 *C1 = New_C1_()
			_ = _nw39
			_nw39.Ctor__(Companion_Default___.Fm1((_230_v166).Cardinality(), _57_v48.F21, _231_v167, globalState), (_dafny.Zero).Minus((_67_v56).F23()), p0, Companion_Default___.Fm1((_67_v56).F23(), _226_v162, _231_v167, globalState))
			_232_v168 = _nw39
			var _233_v169 _dafny.Array
			_ = _233_v169
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_5
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_234_v164 _dafny.Int, _235_v15 bool) func(_dafny.Int) _dafny.Int {
					return func(_236_i13 _dafny.Int) _dafny.Int {
						return (_236_i13).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v164, _235_v15)).Cardinality())
					}
				})(_228_v164, _23_v15)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw40 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw40).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw40).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_233_v169 = _nw40
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_233_v169), 0))
			_ = _index28
			(_233_v169).ArraySet1((_67_v56).F23(), (_index28).Int())
			var _237_v170 _dafny.MultiSet
			_ = _237_v170
			_237_v170 = _dafny.MultiSetOf(_69_v58)
			var _238_v171 _dafny.Array
			_ = _238_v171
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_6
			var _nw41 _dafny.Array
			_ = _nw41
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw41 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) bool = func(_239_i14 _dafny.Int) bool {
					return false
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw41 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw41).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw41).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_238_v171 = _nw41
			var _240_v172 _dafny.Map
			_ = _240_v172
			_240_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_238_v171, _57_v48.F21)
			var _241_v173 _dafny.Set
			_ = _241_v173
			_241_v173 = _dafny.SetOf(_dafny.MultiSetOf(!(_23_v15)), _231_v167, _231_v167, _231_v167, _231_v167)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_233_v169), 0))
			_ = _index29
			(_233_v169).ArraySet1(((_237_v170).Cardinality()).Minus(Companion_Default___.Fm0(_dafny.IntOfUint32((_207_v146).Cardinality()), _dafny.IntOfInt64(785), (_dafny.Zero).Minus((_240_v172).Cardinality()), (_241_v173).Cardinality(), globalState)), (_index29).Int())
			var _242_v174 _dafny.Array
			_ = _242_v174
			var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
			_ = _nw42
			_242_v174 = _nw42
			var _243_v175 _dafny.Set
			_ = _243_v175
			_243_v175 = _dafny.SetOf((_67_v56).F23())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_242_v174), 0))
			_ = _index30
			(_242_v174).ArraySet1(_243_v175, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.ArrayLen((_242_v174), 0))
			_ = _index31
			(_242_v174).ArraySet1((_243_v175).Difference(_243_v175), (_index31).Int())
		} else {
			var _244_v176 _dafny.Array
			_ = _244_v176
			var _nwElement0_11 bool = !(((_69_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_207_v146).Cardinality()), _dafny.IntOfUint32((_69_v58).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_67_v56).F23()) != 0)
			_ = _nwElement0_11
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(7))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_11, 0)
			(_nw43).ArraySet1(!(_57_v48.F21), 1)
			(_nw43).ArraySet1(_57_v48.F21, 2)
			(_nw43).ArraySet1(_57_v48.F21, 3)
			(_nw43).ArraySet1(_57_v48.F21, 4)
			(_nw43).ArraySet1(_23_v15, 5)
			(_nw43).ArraySet1(_57_v48.F21, 6)
			_244_v176 = _nw43
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_244_v176), 0))
			_ = _index32
			(_244_v176).ArraySet1(!(_57_v48.F21) || (_57_v48.F21), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_244_v176), 0))
			_ = _index33
			(_244_v176).ArraySet1(_23_v15, (_index33).Int())
			var _245_v177 *C0
			_ = _245_v177
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__(Companion_D2_.Create_DC6_(_63_v54), (_dafny.IntOfUint32((_207_v146).Cardinality())).Plus((_67_v56).F23()), _57_v48.F21)
			_245_v177 = _nw44
			var _246_v178 D5
			_ = _246_v178
			_246_v178 = Companion_D5_.Create_DC13_((_67_v56).F23())
			var _pat_let_tv6 = _245_v177
			_ = _pat_let_tv6
			var _pat_let_tv7 = p0
			_ = _pat_let_tv7
			var _pat_let_tv8 = _67_v56
			_ = _pat_let_tv8
			var _pat_let_tv9 = _67_v56
			_ = _pat_let_tv9
			var _pat_let_tv10 = globalState
			_ = _pat_let_tv10
			_246_v178 = func(_pat_let4_0 D5) D5 {
				return func(_247_dt__update__tmp_h2 D5) D5 {
					return func(_pat_let5_0 _dafny.Int) D5 {
						return func(_248_dt__update_hcf19_h1 _dafny.Int) D5 {
							return Companion_D5_.Create_DC13_(_248_dt__update_hcf19_h1)
						}(_pat_let5_0)
					}((_dafny.Zero).Minus(Companion_Default___.Fm0((_pat_let_tv6).F23(), _pat_let_tv7, (_pat_let_tv8).F23(), (_pat_let_tv9).F23(), _pat_let_tv10)))
				}(_pat_let4_0)
			}(_246_v178)
			var _249_v179 _dafny.Sequence
			_ = _249_v179
			_249_v179 = _dafny.SeqOf(_245_v177)
			_249_v179 = _249_v179
			var _250_v180 _dafny.Int
			_ = _250_v180
			var _251_v181 bool
			_ = _251_v181
			var _252_v182 _dafny.Int
			_ = _252_v182
			var _253_v183 _dafny.Int
			_ = _253_v183
			var _out12 _dafny.Int
			_ = _out12
			var _out13 bool
			_ = _out13
			var _out14 _dafny.Int
			_ = _out14
			var _out15 _dafny.Int
			_ = _out15
			_out12, _out13, _out14, _out15 = (_57_v48).M1(_23_v15, _57_v48.F21, globalState)
			_250_v180 = _out12
			_251_v181 = _out13
			_252_v182 = _out14
			_253_v183 = _out15
		}
		var _254_v184 _dafny.MultiSet
		_ = _254_v184
		_254_v184 = _dafny.MultiSetOf(_57_v48.F21)
		var _255_v185 _dafny.Set
		_ = _255_v185
		_255_v185 = _dafny.SetOf(_57_v48.F21, _23_v15)
		var _256_v186 *C1
		_ = _256_v186
		var _nw45 *C1 = New_C1_()
		_ = _nw45
		_nw45.Ctor__(!((Companion_Default___.Fm1((_67_v56).F23(), _57_v48.F21, _254_v184, globalState)) == (_57_v48.F21)), (_dafny.IntOfInt64(366)).Plus(p1), (_67_v56).F23(), (_255_v185).IsProperSubsetOf(_255_v185))
		_256_v186 = _nw45
		var _257_v187 D6
		_ = _257_v187
		_257_v187 = Companion_D6_.Create_DC17_((func() _dafny.Int {
			if (_63_v54).Contains(_dafny.IntOfInt64(123)) {
				return (_63_v54).Multiplicity(_dafny.IntOfInt64(123))
			}
			return (_67_v56).F23()
		})(), _256_v186.F21)
		var _258_v188 *C1
		_ = _258_v188
		var _nw46 *C1 = New_C1_()
		_ = _nw46
		_nw46.Ctor__(_57_v48.F21, (_257_v187).Dtor_cf23(), p1, (_dafny.IntOfUint32((_207_v146).Cardinality())).Cmp(p0) == 0)
		_258_v188 = _nw46
	} else {
		_23_v15 = (_57_v48.F21) || (true)
		if _57_v48.F21 {
			var _259_v189 _dafny.MultiSet
			_ = _259_v189
			_259_v189 = _dafny.MultiSetOf(_57_v48.F21)
			var _260_v190 *C1
			_ = _260_v190
			var _nw47 *C1 = New_C1_()
			_ = _nw47
			_nw47.Ctor__(_57_v48.F21, Companion_Default___.SafeModuloInt((_259_v189).Cardinality(), p0), (_67_v56).F23(), !(_57_v48.F21))
			_260_v190 = _nw47
			(_57_v48).F21 = false
			var _261_v191 _dafny.Map
			_ = _261_v191
			_261_v191 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, p1)
			var _262_v192 _dafny.Map
			_ = _262_v192
			_262_v192 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_57_v48).Fm5(_57_v48.F21, _dafny.IntOfInt64(308), globalState), _261_v191)
			_262_v192 = (_262_v192).Update(_57_v48.F21, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, p0))
			(_260_v190).F21 = _260_v190.F21
			var _263_v193 _dafny.Map
			_ = _263_v193
			_263_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (Companion_D6_.Create_DC14_(_67_v56)).Dtor_cf20())
			var _264_v194 T1
			_ = _264_v194
			var _nw48 *C1 = New_C1_()
			_ = _nw48
			_nw48.Ctor__(false, (_67_v56).F23(), ((_259_v189).Update(_57_v48.F21, Companion_Default___.Abs((_67_v56).F23()))).Cardinality(), _57_v48.F21)
			_264_v194 = _nw48
			var _265_v195 _dafny.Map
			_ = _265_v195
			_265_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_264_v194, _57_v48.F21)
			var _266_v196 _dafny.Map
			_ = _266_v196
			_266_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v56, (_68_v57).Select((Companion_Default___.SafeIndex((_67_v56).F23(), _dafny.IntOfUint32((_68_v57).Cardinality()))).Uint32()).(*C0))
			var _267_v197 _dafny.Array
			_ = _267_v197
			var _nwElement0_12 *C0 = (func() *C0 {
				if (_263_v193).Contains((_265_v195).Cardinality()) {
					return (_263_v193).Get((_265_v195).Cardinality()).(*C0)
				}
				return _67_v56
			})()
			_ = _nwElement0_12
			var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(19))
			_ = _nw49
			(_nw49).ArraySet1(_nwElement0_12, 0)
			(_nw49).ArraySet1(_67_v56, 1)
			(_nw49).ArraySet1(_67_v56, 2)
			(_nw49).ArraySet1(_67_v56, 3)
			(_nw49).ArraySet1(_67_v56, 4)
			(_nw49).ArraySet1((func() *C0 {
				if (_266_v196).Contains(_67_v56) {
					return (_266_v196).Get(_67_v56).(*C0)
				}
				return _67_v56
			})(), 5)
			(_nw49).ArraySet1((func() *C0 {
				if _57_v48.F21 {
					return _67_v56
				}
				return _67_v56
			})(), 6)
			(_nw49).ArraySet1(_67_v56, 7)
			(_nw49).ArraySet1(_67_v56, 8)
			(_nw49).ArraySet1(_67_v56, 9)
			(_nw49).ArraySet1(_67_v56, 10)
			(_nw49).ArraySet1(_67_v56, 11)
			(_nw49).ArraySet1((_68_v57).Select((Companion_Default___.SafeIndex((_264_v194).F19(), _dafny.IntOfUint32((_68_v57).Cardinality()))).Uint32()).(*C0), 12)
			(_nw49).ArraySet1(_67_v56, 13)
			(_nw49).ArraySet1(_67_v56, 14)
			(_nw49).ArraySet1(_67_v56, 15)
			(_nw49).ArraySet1(_67_v56, 16)
			(_nw49).ArraySet1(_67_v56, 17)
			(_nw49).ArraySet1(_67_v56, 18)
			_267_v197 = _nw49
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_267_v197), 0))
			_ = _index34
			(_267_v197).ArraySet1(_67_v56, (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(12), _dafny.ArrayLen((_267_v197), 0))
			_ = _index35
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(_67_v56.F22, (_264_v194).F19(), !(_23_v15))
			(_267_v197).ArraySet1(_nw50, (_index35).Int())
		} else {
			var _268_v198 T0
			_ = _268_v198
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__(_66_v55, p1, _57_v48.F21)
			_268_v198 = _nw51
			_268_v198 = _268_v198
			(globalState).F3 = (_dafny.Zero).Minus((_67_v56).F23())
			var _269_v199 _dafny.Map
			_ = _269_v199
			_269_v199 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _69_v58)
			_269_v199 = (_269_v199).Update(Companion_Default___.SafeModuloInt(p0, (_67_v56).F23()), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_67_v56).F23()), (Companion_Default___.SafeIndex((_67_v56).F23(), _dafny.IntOfUint32((_dafny.SeqOf((_67_v56).F23())).Cardinality()))).Uint32(), p1))
			_69_v58 = _dafny.SeqOf((_67_v56).F23(), (_67_v56).F23())
			var _270_v200 _dafny.Array
			_ = _270_v200
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_7
			var _nw52 _dafny.Array
			_ = _nw52
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw52 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = (func(_271_v48 *C1, _272_v198 T0) func(_dafny.Int) bool {
					return func(_273_i15 _dafny.Int) bool {
						return (_dafny.MultiSetOf(_271_v48.F21, _271_v48.F21)).IsSubsetOf(_dafny.MultiSetOf((_272_v198).F18()))
					}
				})(_57_v48, _268_v198)
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
			_270_v200 = _nw52
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_270_v200), 0))
			_ = _index36
			(_270_v200).ArraySet1(_57_v48.F21, (_index36).Int())
			var _274_v201 _dafny.Set
			_ = _274_v201
			_274_v201 = _dafny.SetOf(_270_v200)
			var _275_v202 _dafny.Set
			_ = _275_v202
			_275_v202 = _274_v201
			var _276_v203 T1
			_ = _276_v203
			var _nw53 *C1 = New_C1_()
			_ = _nw53
			_nw53.Ctor__(_57_v48.F21, p1, (_67_v56).F23(), true)
			_276_v203 = _nw53
			var _277_v204 _dafny.Map
			_ = _277_v204
			_277_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_v203, (_276_v203).F19())
			var _278_v205 _dafny.Sequence
			_ = _278_v205
			_278_v205 = _dafny.SeqOf(_276_v203, _276_v203, _276_v203)
			var _279_v206 _dafny.Map
			_ = _279_v206
			_279_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, (_277_v204).Update((_278_v205).Select((Companion_Default___.SafeIndex((_67_v56).F23(), _dafny.IntOfUint32((_278_v205).Cardinality()))).Uint32()).(T1), p1))
			var _280_v207 _dafny.Map
			_ = _280_v207
			_280_v207 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Map {
				if (_279_v206).Contains(_57_v48.F21) {
					return (_279_v206).Get(_57_v48.F21).(_dafny.Map)
				}
				return _277_v204
			})()).Cardinality(), _192_v143)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_270_v200), 0))
			_ = _index37
			var _rhs39 bool = !(((func() _dafny.Set {
				if _57_v48.F21 {
					return _dafny.SetOf(_270_v200, _270_v200, _270_v200, _270_v200, _270_v200)
				}
				return (_275_v202)
			})()).IsDisjointFrom((_dafny.SetOf(_270_v200)).Difference(_274_v201)))
			_ = _rhs39
			var _rhs40 _dafny.Int = p1
			_ = _rhs40
			var _rhs41 bool = !((Companion_Default___.SafeModuloInt(p1, p0)).Cmp((_280_v207).Cardinality()) >= 0)
			_ = _rhs41
			var _lhs31 _dafny.Array = _270_v200
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_270_v200), 0))
			_ = _lhs32
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			(_lhs31).ArraySet1(_rhs39, (_lhs32).Int())
			_lhs33.F5 = _rhs40
			_lhs34.F13 = _rhs41
		}
		var _281_v208 D2
		_ = _281_v208
		_281_v208 = Companion_D2_.Create_DC5_((func() _dafny.Map {
			if _57_v48.F21 {
				return _192_v143
			}
			return _192_v143
		})())
		var _rhs42 _dafny.Map = _192_v143
		_ = _rhs42
		var _rhs43 D2 = Companion_D2_.Create_DC5_(_192_v143)
		_ = _rhs43
		var _rhs44 bool = _23_v15
		_ = _rhs44
		var _lhs35 *GlobalState = globalState
		_ = _lhs35
		_192_v143 = _rhs42
		_281_v208 = _rhs43
		_lhs35.F13 = _rhs44
		var _282_v209 _dafny.Array
		_ = _282_v209
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_8
		var _nw54 _dafny.Array
		_ = _nw54
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw54 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Int = (func(_283_v16 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_284_i16 _dafny.Int) _dafny.Int {
					return (_284_i16).Plus(_dafny.IntOfUint32((_283_v16).Cardinality()))
				}
			})(_24_v16)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw54 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw54).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw54).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_282_v209 = _nw54
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_282_v209), 0))
		_ = _index38
		(_282_v209).ArraySet1(p0, (_index38).Int())
		var _285_v210 _dafny.Array
		_ = _285_v210
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_9
		var _nw55 _dafny.Array
		_ = _nw55
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw55 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Sequence = (func(_286_v146 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_287_i17 _dafny.Int) _dafny.Sequence {
					return _286_v146
				}
			})(_207_v146)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw55 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw55).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw55).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_285_v210 = _nw55
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_285_v210), 0))
		_ = _index39
		(_285_v210).ArraySet1(_207_v146, (_index39).Int())
		var _288_v211 _dafny.Array
		_ = _288_v211
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_10
		var _nw56 _dafny.Array
		_ = _nw56
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw56 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) bool = (func(_289_v48 *C1) func(_dafny.Int) bool {
				return func(_290_i18 _dafny.Int) bool {
					return _289_v48.F21
				}
			})(_57_v48)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw56 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw56).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw56).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_288_v211 = _nw56
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_288_v211), 0))
		_ = _index40
		(_288_v211).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_69_v58, _69_v58), (_index40).Int())
		var _291_v212 _dafny.Set
		_ = _291_v212
		_291_v212 = _dafny.SetOf(_57_v48.F21)
		var _292_v213 T0
		_ = _292_v213
		var _nw57 *C0 = New_C0_()
		_ = _nw57
		_nw57.Ctor__(_67_v56.F22, (_291_v212).Cardinality(), !(_23_v15))
		_292_v213 = _nw57
		var _293_v214 _dafny.Set
		_ = _293_v214
		_293_v214 = _dafny.SetOf(_292_v213, _292_v213)
		var _294_v215 _dafny.CodePoint
		_ = _294_v215
		_294_v215 = _dafny.CodePoint('j')
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_282_v209), 0))
		_ = _index41
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_285_v210), 0))
		_ = _index42
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_288_v211), 0))
		_ = _index43
		var _rhs45 _dafny.Int = p1
		_ = _rhs45
		var _rhs46 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(468))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_295_i19 _dafny.Int) _dafny.CodePoint {
			return (func() _dafny.CodePoint {
				if false {
					return _dafny.CodePoint('v')
				}
				return _dafny.CodePoint('e')
			})()
		}))).Cardinality())
		_ = _rhs46
		var _rhs47 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_207_v146, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_207_v146).Cardinality()))).Uint32(), _294_v215)
		_ = _rhs47
		var _rhs48 bool = _57_v48.F21
		_ = _rhs48
		var _rhs49 _dafny.Set = (func() _dafny.Set {
			if _dafny.Companion_Sequence_.Contains(_207_v146, _294_v215) {
				return _dafny.SetOf(_292_v213, _292_v213, _292_v213)
			}
			return _dafny.SetOf(_292_v213)
		})()
		_ = _rhs49
		var _lhs36 _dafny.Array = _282_v209
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_282_v209), 0))
		_ = _lhs37
		var _lhs38 *GlobalState = globalState
		_ = _lhs38
		var _lhs39 _dafny.Array = _285_v210
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_285_v210), 0))
		_ = _lhs40
		var _lhs41 _dafny.Array = _288_v211
		_ = _lhs41
		var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_288_v211), 0))
		_ = _lhs42
		(_lhs36).ArraySet1(_rhs45, (_lhs37).Int())
		_lhs38.F5 = _rhs46
		(_lhs39).ArraySet1(_rhs47, (_lhs40).Int())
		(_lhs41).ArraySet1(_rhs48, (_lhs42).Int())
		_293_v214 = _rhs49
		var _296_v216 _dafny.Map
		_ = _296_v216
		_296_v216 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v48.F21, _57_v48.F21)
		var _297_v217 D1
		_ = _297_v217
		_297_v217 = Companion_D1_.Create_DC4_(!((_292_v213).F18()), (func() bool {
			if (_24_v16).Select((Companion_Default___.SafeIndex((_67_v56).F23(), _dafny.IntOfUint32((_24_v16).Cardinality()))).Uint32()).(bool) {
				return (_57_v48).Fm5(_23_v15, (_67_v56).F23(), globalState)
			}
			return _57_v48.F21
		})(), (p0).Cmp((_296_v216).Cardinality()) > 0)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_282_v209), 0))
		_ = _index44
		var _rhs50 _dafny.Int = (_282_v209).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_282_v209), 0))).Int()).(_dafny.Int)
		_ = _rhs50
		var _rhs51 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32(((_285_v210).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_285_v210), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Cardinality())
		_ = _rhs51
		var _rhs52 D1 = _297_v217
		_ = _rhs52
		var _lhs43 *GlobalState = globalState
		_ = _lhs43
		var _lhs44 _dafny.Array = _282_v209
		_ = _lhs44
		var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_282_v209), 0))
		_ = _lhs45
		_lhs43.F5 = _rhs50
		(_lhs44).ArraySet1(_rhs51, (_lhs45).Int())
		_297_v217 = _rhs52
	}
	var _298_v218 T0
	_ = _298_v218
	var _nw58 *C0 = New_C0_()
	_ = _nw58
	_nw58.Ctor__(_66_v55, p0, _57_v48.F21)
	_298_v218 = _nw58
	var _299_v219 _dafny.Map
	_ = _299_v219
	_299_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v56, _298_v218)
	r0 = (((_299_v219).Update(_67_v56, _298_v218)).Merge(_299_v219)).Cardinality()
	var _300_v220 _dafny.Array
	_ = _300_v220
	var _len0_11 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_11
	var _nw59 _dafny.Array
	_ = _nw59
	if _len0_11.Cmp(_dafny.Zero) == 0 {
		_nw59 = _dafny.NewArray(_len0_11)
	} else {
		var _init11 func(_dafny.Int) _dafny.Int = (func(_301_v56 *C0) func(_dafny.Int) _dafny.Int {
			return func(_302_i20 _dafny.Int) _dafny.Int {
				return (_302_i20).Minus((_301_v56).F23())
			}
		})(_67_v56)
		_ = _init11
		var _element0_11 = _init11(_dafny.Zero)
		_ = _element0_11
		_nw59 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
		(_nw59).ArraySet1(_element0_11, 0)
		var _nativeLen0_11 = (_len0_11).Int()
		_ = _nativeLen0_11
		for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
			(_nw59).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
		}
	}
	_300_v220 = _nw59
	var _303_v221 _dafny.Map
	_ = _303_v221
	_303_v221 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), _300_v220)
	var _304_v222 _dafny.MultiSet
	_ = _304_v222
	_304_v222 = _dafny.MultiSetOf(_57_v48.F21)
	r1 = (_303_v221).Update((_304_v222).Cardinality(), _300_v220)
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _305_v0 _dafny.Array
	_ = _305_v0
	var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
	_ = _nw60
	_305_v0 = _nw60
	var _306_v1 _dafny.Array
	_ = _306_v1
	var _nw61 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
	_ = _nw61
	_306_v1 = _nw61
	var _307_v2 _dafny.Sequence
	_ = _307_v2
	_307_v2 = _dafny.UnicodeSeqOfUtf8Bytes("mqrmmy")
	var _308_v3 bool
	_ = _308_v3
	_308_v3 = false
	var _309_globalState *GlobalState
	_ = _309_globalState
	var _nw62 *GlobalState = New_GlobalState_()
	_ = _nw62
	_nw62.Ctor__(_305_v0, true, _306_v1, _dafny.Zero, _dafny.UnicodeSeqOfUtf8Bytes("okaofhw"), _dafny.IntOfInt64(-352), _dafny.IntOfInt64(893), _307_v2, _dafny.IntOfInt64(868), _dafny.IntOfInt64(685), true, _dafny.IntOfInt64(229), false, false, _dafny.IntOfInt64(157), false, true, (func() _dafny.Array {
		if _308_v3 {
			return _306_v1
		}
		return _306_v1
	})())
	_309_globalState = _nw62
	var _310_v4 _dafny.Array
	_ = _310_v4
	var _len0_12 _dafny.Int = _dafny.IntOfInt64(14)
	_ = _len0_12
	var _nw63 _dafny.Array
	_ = _nw63
	if _len0_12.Cmp(_dafny.Zero) == 0 {
		_nw63 = _dafny.NewArray(_len0_12)
	} else {
		var _init12 func(_dafny.Int) _dafny.Int = (func(_311_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_312_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_312_i0, _dafny.IntOfUint32((_311_v2).Cardinality()))
			}
		})(_307_v2)
		_ = _init12
		var _element0_12 = _init12(_dafny.Zero)
		_ = _element0_12
		_nw63 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
		(_nw63).ArraySet1(_element0_12, 0)
		var _nativeLen0_12 = (_len0_12).Int()
		_ = _nativeLen0_12
		for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
			(_nw63).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
		}
	}
	_310_v4 = _nw63
	var _313_v5 _dafny.Map
	_ = _313_v5
	_313_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _310_v4)
	var _314_v6 _dafny.Sequence
	_ = _314_v6
	_314_v6 = _dafny.SeqOf((func() _dafny.Array {
		if (_313_v5).Contains(_308_v3) {
			return (_313_v5).Get(_308_v3).(_dafny.Array)
		}
		return _310_v4
	})(), _310_v4, _310_v4)
	_314_v6 = _dafny.Companion_Sequence_.Concatenate(_314_v6, _314_v6)
	var _315_v7 _dafny.Int
	_ = _315_v7
	_315_v7 = _dafny.IntOfInt64(604)
	(_309_globalState).F9 = ((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_308_v3)).Cardinality())))).Times(Companion_Default___.Fm0(_dafny.IntOfUint32((_307_v2).Cardinality()), _dafny.IntOfUint32((_307_v2).Cardinality()), _315_v7, _315_v7, _309_globalState))
	var _316_v8 _dafny.MultiSet
	_ = _316_v8
	_316_v8 = _dafny.MultiSetOf(_310_v4, _310_v4)
	if !(((_dafny.MultiSetOf(_310_v4)).Update(_310_v4, Companion_Default___.Abs(_315_v7))).Equals((_316_v8).Union(_316_v8))) {
		var _317_v9 _dafny.Array
		_ = _317_v9
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_13
		var _nw64 _dafny.Array
		_ = _nw64
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw64 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Set = (func(_318_v7 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_319_i1 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_318_v7)
				}
			})(_315_v7)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw64 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw64).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw64).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_317_v9 = _nw64
		var _320_v10 _dafny.CodePoint
		_ = _320_v10
		_320_v10 = _dafny.CodePoint('f')
		var _321_v11 _dafny.Sequence
		_ = _321_v11
		_321_v11 = _dafny.SeqOf(_308_v3, _308_v3, _308_v3, (_315_v7).Cmp(_315_v7) != 0, (_dafny.IntOfUint32((_307_v2).Cardinality())).Cmp(_315_v7) > 0)
		var _rhs53 _dafny.Int = (((_dafny.Zero).Minus(_315_v7)).Times((_dafny.Zero).Minus(_315_v7))).Minus(_315_v7)
		_ = _rhs53
		var _rhs54 _dafny.Array = (func() _dafny.Array {
			if (_315_v7).Cmp(_315_v7) <= 0 {
				return _317_v9
			}
			return _317_v9
		})()
		_ = _rhs54
		var _rhs55 _dafny.CodePoint = _320_v10
		_ = _rhs55
		var _rhs56 _dafny.Array = _305_v0
		_ = _rhs56
		var _rhs57 _dafny.Sequence = (func() _dafny.Sequence {
			if _308_v3 {
				return _dafny.SeqOf(_308_v3, _308_v3, Companion_Default___.Fm1(_315_v7, _308_v3, _dafny.MultiSetOf(_308_v3), _309_globalState))
			}
			return _321_v11
		})()
		_ = _rhs57
		var _lhs46 *GlobalState = _309_globalState
		_ = _lhs46
		_lhs46.F9 = _rhs53
		_317_v9 = _rhs54
		_320_v10 = _rhs55
		_305_v0 = _rhs56
		_321_v11 = _rhs57
		var _322_v12 _dafny.MultiSet
		_ = _322_v12
		_322_v12 = _dafny.MultiSetOf(_308_v3)
		var _323_v13 _dafny.MultiSet
		_ = _323_v13
		_323_v13 = _dafny.MultiSetOf(_306_v1)
		var _324_v14 _dafny.Int
		_ = _324_v14
		var _325_v15 _dafny.Map
		_ = _325_v15
		var _out16 _dafny.Int
		_ = _out16
		var _out17 _dafny.Map
		_ = _out17
		_out16, _out17 = Companion_Default___.M0((func() _dafny.Int {
			if (_322_v12).Contains(_308_v3) {
				return (_322_v12).Multiplicity(_308_v3)
			}
			return (_323_v13).Cardinality()
		})(), _315_v7, _309_globalState)
		_324_v14 = _out16
		_325_v15 = _out17
		var _326_v16 _dafny.Set
		_ = _326_v16
		_326_v16 = _dafny.SetOf(_308_v3, _308_v3)
		var _327_v17 _dafny.Sequence
		_ = _327_v17
		_327_v17 = _dafny.SeqOf(_dafny.IntOfInt64(615))
		var _328_v18 _dafny.Map
		_ = _328_v18
		_328_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_315_v7), (_321_v11).Select((Companion_Default___.SafeIndex(_324_v14, _dafny.IntOfUint32((_321_v11).Cardinality()))).Uint32()).(bool))
		var _329_v19 _dafny.Map
		_ = _329_v19
		_329_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_326_v16).Cardinality(), (_327_v17).Select((Companion_Default___.SafeIndex(_315_v7, _dafny.IntOfUint32((_327_v17).Cardinality()))).Uint32()).(_dafny.Int), _315_v7, (_328_v18).Cardinality(), _309_globalState), _308_v3)
		var _330_v20 _dafny.Sequence
		_ = _330_v20
		_330_v20 = _dafny.SeqOf((_329_v19).Cardinality())
		var _331_v21 _dafny.Sequence
		_ = _331_v21
		_331_v21 = _dafny.SeqOf(_315_v7, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_327_v17).Cardinality()), _324_v14), _315_v7)
		var _332_v22 _dafny.Map
		_ = _332_v22
		_332_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _308_v3)
		var _333_v23 _dafny.MultiSet
		_ = _333_v23
		_333_v23 = _dafny.MultiSetOf(_315_v7)
		var _334_v24 _dafny.Sequence
		_ = _334_v24
		_334_v24 = _dafny.SeqOf(_333_v23)
		var _335_v25 _dafny.Sequence
		_ = _335_v25
		_335_v25 = _dafny.SeqOf(_334_v24, _334_v24)
		var _336_v26 _dafny.Map
		_ = _336_v26
		_336_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _315_v7)
		var _rhs58 _dafny.Sequence = Companion_Default___.Fm2(_320_v10, _308_v3, _309_globalState)
		_ = _rhs58
		var _rhs59 _dafny.Int = (func() _dafny.Int {
			if (_dafny.IntOfInt64(568)).Cmp((_332_v22).Cardinality()) != 0 {
				return _315_v7
			}
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_335_v25).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_336_v26).Contains(_dafny.IntOfInt64(515)) {
					return (_336_v26).Get(_dafny.IntOfInt64(515)).(_dafny.Int)
				}
				return _324_v14
			})(), _dafny.IntOfUint32((_335_v25).Cardinality()))).Uint32()).(_dafny.Sequence), _334_v24)).Cardinality())
		})()
		_ = _rhs59
		_330_v20 = _rhs58
		_324_v14 = _rhs59
		(_309_globalState).F13 = _308_v3
		(_309_globalState).F3 = _324_v14
	} else {
		if (false) == (_308_v3) {
			(_309_globalState).F2 = _306_v1
			var _337_v27 _dafny.CodePoint
			_ = _337_v27
			_337_v27 = _dafny.CodePoint('r')
			var _338_v28 _dafny.Map
			_ = _338_v28
			_338_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _337_v27)
			var _339_v29 _dafny.Sequence
			_ = _339_v29
			_339_v29 = _dafny.SeqOf(_338_v28)
			(_309_globalState).F5 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_339_v29, _339_v29)).Cardinality())
			var _340_v30 _dafny.Map
			_ = _340_v30
			_340_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC0_(_308_v3)).Dtor_cf0(), _315_v7)
			var _341_v31 _dafny.Sequence
			_ = _341_v31
			_341_v31 = _dafny.SeqOf(_308_v3, _308_v3)
			var _rhs60 _dafny.Map = _340_v30
			_ = _rhs60
			var _rhs61 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if false {
					return _341_v31
				}
				return _341_v31
			})(), _341_v31)
			_ = _rhs61
			var _rhs62 bool = _308_v3
			_ = _rhs62
			var _lhs47 *GlobalState = _309_globalState
			_ = _lhs47
			_340_v30 = _rhs60
			_341_v31 = _rhs61
			_lhs47.F13 = _rhs62
			(_309_globalState).F3 = _dafny.IntOfInt64(176)
			var _342_v32 _dafny.Array
			_ = _342_v32
			var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
			_ = _nw65
			_342_v32 = _nw65
			var _343_v33 _dafny.Set
			_ = _343_v33
			_343_v33 = _dafny.SetOf(_315_v7, _315_v7, _315_v7, _dafny.IntOfInt64(-714), (_dafny.Zero).Minus((_dafny.Zero).Minus(_315_v7)))
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_342_v32), 0))
			_ = _index45
			(_342_v32).ArraySet1(_343_v33, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_342_v32), 0))
			_ = _index46
			(_342_v32).ArraySet1(_343_v33, (_index46).Int())
		} else {
			var _344_v34 _dafny.Sequence
			_ = _344_v34
			_344_v34 = _dafny.SeqOf(_308_v3)
			var _345_v35 _dafny.Map
			_ = _345_v35
			_345_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _308_v3)
			var _346_v36 _dafny.Sequence
			_ = _346_v36
			_346_v36 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_344_v34, _344_v34), _dafny.Companion_Sequence_.Concatenate((Companion_D1_.Create_DC2_(_dafny.SeqOf(_308_v3, _308_v3))).Dtor_cf4(), _dafny.SeqOf((func() bool {
				if (_345_v35).Contains(_308_v3) {
					return (_345_v35).Get(_308_v3).(bool)
				}
				return _308_v3
			})())))
			_344_v34 = (_346_v36).Select((Companion_Default___.SafeIndex(_315_v7, _dafny.IntOfUint32((_346_v36).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _347_v37 _dafny.Map
			_ = _347_v37
			_347_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_315_v7).Plus(_315_v7), _308_v3)
			_347_v37 = (_347_v37).Update(_315_v7, !_dafny.Companion_Sequence_.Equal(_307_v2, _307_v2))
			(_309_globalState).F9 = (_dafny.Zero).Minus((func() _dafny.Int {
				if _308_v3 {
					return (_dafny.Zero).Minus((_315_v7).Times(_315_v7))
				}
				return _315_v7
			})())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_306_v1), 0))
			_ = _index47
			(_306_v1).ArraySet1(false, (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_306_v1), 0))
			_ = _index48
			(_306_v1).ArraySet1((func() bool {
				if _308_v3 {
					return _308_v3
				}
				return false
			})(), (_index48).Int())
			var _348_v38 D1
			_ = _348_v38
			_348_v38 = Companion_D1_.Create_DC4_(_308_v3, (_306_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_306_v1), 0))).Int()).(bool), _308_v3)
			var _349_v39 _dafny.MultiSet
			_ = _349_v39
			_349_v39 = _dafny.MultiSetOf(_348_v38, _348_v38)
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_310_v4), 0))
			_ = _index49
			(_310_v4).ArraySet1((func() _dafny.Int {
				if (_349_v39).Contains(_348_v38) {
					return (_349_v39).Multiplicity(_348_v38)
				}
				return _315_v7
			})(), (_index49).Int())
			var _350_v40 _dafny.Array
			_ = _350_v40
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_14
			var _nw66 _dafny.Array
			_ = _nw66
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw66 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Sequence = (func(_351_v1 _dafny.Array, _352_v7 _dafny.Int, _353_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_354_i2 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if (_351_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_351_v1), 0))).Int()).(bool) {
								return _dafny.UnicodeSeqOfUtf8Bytes("ghaimy")
							}
							return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nvj"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_352_v7), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nvj")).Cardinality()))).Uint32(), _dafny.CodePoint('k'))
						})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_353_v2).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_351_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_351_v1), 0))).Int()).(bool) {
								return _dafny.UnicodeSeqOfUtf8Bytes("ghaimy")
							}
							return _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nvj"), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_352_v7), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nvj")).Cardinality()))).Uint32(), _dafny.CodePoint('k'))
						})()).Cardinality()))).Uint32(), _dafny.CodePoint('l'))
					}
				})(_306_v1, _315_v7, _307_v2)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw66 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw66).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw66).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_350_v40 = _nw66
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_350_v40), 0))
			_ = _index50
			(_350_v40).ArraySet1(_307_v2, (_index50).Int())
			var _355_v41 _dafny.MultiSet
			_ = _355_v41
			_355_v41 = _dafny.MultiSetOf((_306_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_306_v1), 0))).Int()).(bool))
			var _356_v42 _dafny.Map
			_ = _356_v42
			_356_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _307_v2)
			var _357_v43 _dafny.Sequence
			_ = _357_v43
			_357_v43 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_307_v2, _307_v2), (func() _dafny.Sequence {
				if (_356_v42).Contains(_308_v3) {
					return (_356_v42).Get(_308_v3).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}(func(_358_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))
			})(), _307_v2)
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_310_v4), 0))
			_ = _index51
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_350_v40), 0))
			_ = _index52
			var _rhs63 bool = Companion_Default___.Fm1(_315_v7, _dafny.Companion_Sequence_.IsProperPrefixOf(_307_v2, _307_v2), (_dafny.MultiSetOf(_308_v3, (_306_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_306_v1), 0))).Int()).(bool))).Difference(_355_v41), _309_globalState)
			_ = _rhs63
			var _rhs64 _dafny.Int = Companion_Default___.SafeDivisionInt(_315_v7, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_307_v2).Cardinality())), _315_v7))
			_ = _rhs64
			var _rhs65 bool = _308_v3
			_ = _rhs65
			var _rhs66 _dafny.Sequence = (_357_v43).Select((Companion_Default___.SafeIndex(_315_v7, _dafny.IntOfUint32((_357_v43).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs66
			var _lhs48 *GlobalState = _309_globalState
			_ = _lhs48
			var _lhs49 _dafny.Array = _310_v4
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_310_v4), 0))
			_ = _lhs50
			var _lhs51 _dafny.Array = _350_v40
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_350_v40), 0))
			_ = _lhs52
			_lhs48.F13 = _rhs63
			(_lhs49).ArraySet1(_rhs64, (_lhs50).Int())
			_308_v3 = _rhs65
			(_lhs51).ArraySet1(_rhs66, (_lhs52).Int())
		}
		var _359_v44 _dafny.Sequence
		_ = _359_v44
		_359_v44 = _dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_315_v7, _315_v7)))
		_359_v44 = (func() _dafny.Sequence {
			if !((func() bool {
				if _308_v3 {
					return !(_308_v3)
				}
				return _308_v3
			})()) {
				return _359_v44
			}
			return _dafny.Companion_Sequence_.Concatenate(_359_v44, _359_v44)
		})()
		if (func() bool {
			if (_315_v7).Cmp(_dafny.IntOfInt64(810)) > 0 {
				return !(_308_v3) || (_308_v3)
			}
			return (_315_v7).Cmp(_315_v7) > 0
		})() {
			var _360_v45 _dafny.MultiSet
			_ = _360_v45
			_360_v45 = _dafny.MultiSetOf(_308_v3)
			var _361_v46 _dafny.Set
			_ = _361_v46
			_361_v46 = _dafny.SetOf(true, Companion_Default___.Fm1(_315_v7, _308_v3, _360_v45, _309_globalState))
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_310_v4), 0))
			_ = _index53
			(_310_v4).ArraySet1(Companion_Default___.Fm0((_361_v46).Cardinality(), _315_v7, _315_v7, _315_v7, _309_globalState), (_index53).Int())
			var _362_v47 _dafny.Set
			_ = _362_v47
			_362_v47 = _dafny.SetOf(_307_v2, _307_v2, Companion_Default___.Fm3(_309_globalState), _307_v2, _307_v2)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(261), _dafny.ArrayLen((_310_v4), 0))
			_ = _index54
			(_310_v4).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm0(_315_v7, _dafny.IntOfInt64(780), _315_v7, ((func() _dafny.Set {
				if _308_v3 {
					return _362_v47
				}
				return _362_v47
			})()).Cardinality(), _309_globalState)), (_index54).Int())
			(_309_globalState).F13 = _308_v3
			_308_v3 = !(_308_v3)
			var _363_v48 _dafny.CodePoint
			_ = _363_v48
			_363_v48 = _dafny.CodePoint('b')
			var _rhs67 bool = _308_v3
			_ = _rhs67
			var _rhs68 _dafny.Int = _315_v7
			_ = _rhs68
			var _rhs69 _dafny.CodePoint = _363_v48
			_ = _rhs69
			var _lhs53 *GlobalState = _309_globalState
			_ = _lhs53
			var _lhs54 *GlobalState = _309_globalState
			_ = _lhs54
			_lhs53.F13 = _rhs67
			_lhs54.F9 = _rhs68
			_363_v48 = _rhs69
			var _364_v49 _dafny.Int
			_ = _364_v49
			var _365_v50 _dafny.Map
			_ = _365_v50
			var _out18 _dafny.Int
			_ = _out18
			var _out19 _dafny.Map
			_ = _out19
			_out18, _out19 = Companion_Default___.M0(_315_v7, _315_v7, _309_globalState)
			_364_v49 = _out18
			_365_v50 = _out19
		} else {
			var _366_v51 _dafny.CodePoint
			_ = _366_v51
			_366_v51 = _dafny.CodePoint('a')
			_366_v51 = (func() _dafny.CodePoint {
				if _308_v3 {
					return (_307_v2).Select((Companion_Default___.SafeIndex(_315_v7, _dafny.IntOfUint32((_307_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
				return (func() _dafny.CodePoint {
					if _308_v3 {
						return _366_v51
					}
					return _366_v51
				})()
			})()
			var _367_v52 _dafny.Map
			_ = _367_v52
			_367_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _dafny.IntOfInt64(628))
			var _368_v53 _dafny.MultiSet
			_ = _368_v53
			_368_v53 = _dafny.MultiSetOf(_367_v52, _367_v52, (_367_v52).Update(_308_v3, _315_v7))
			(_309_globalState).F13 = (_368_v53).IsDisjointFrom((_368_v53).Union(_368_v53))
			var _369_v54 _dafny.MultiSet
			_ = _369_v54
			_369_v54 = _dafny.MultiSetOf(_306_v1, _306_v1)
			var _370_v55 _dafny.MultiSet
			_ = _370_v55
			_370_v55 = _dafny.MultiSetOf(_308_v3, _308_v3)
			(_309_globalState).F5 = (((_369_v54).Union((_369_v54).Update(_306_v1, Companion_Default___.Abs(_315_v7)))).Cardinality()).Plus(((_367_v52).Cardinality()).Minus((_370_v55).Cardinality()))
			var _371_v56 _dafny.Sequence
			_ = _371_v56
			_371_v56 = _dafny.SeqOf(!(_308_v3))
			var _372_v57 _dafny.Map
			_ = _372_v57
			_372_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_315_v7, _315_v7, _315_v7, _dafny.IntOfUint32((_371_v56).Cardinality()), _309_globalState), _315_v7)
			var _373_v59 _dafny.MultiSet
			_ = _373_v59
			_373_v59 = _dafny.MultiSetOf(_315_v7, _315_v7)
			var _374_v60 _dafny.Sequence
			_ = _374_v60
			_374_v60 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(248), _315_v7))
			var _375_v61 D2
			_ = _375_v61
			_375_v61 = Companion_D2_.Create_DC5_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, (func() _dafny.Int {
				if (_367_v52).Contains(_308_v3) {
					return (_367_v52).Get(_308_v3).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-956)
			})()))
			var _376_v64 _dafny.Array
			_ = _376_v64
			var _nwElement0_13 _dafny.Map = _372_v57
			_ = _nwElement0_13
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(27))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_13, 0)
			(_nw67).ArraySet1(_372_v57, 1)
			(_nw67).ArraySet1((_372_v57).Merge(_372_v57), 2)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_373_v59).Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _377_v58 _dafny.Int
					_377_v58 = interface{}(_compr_6).(_dafny.Int)
					if (_373_v59).Contains(_377_v58) {
						_coll6.Add(Companion_Default___.SafeDivisionInt(_377_v58, _315_v7), (_dafny.Zero).Minus(_315_v7))
					}
				}
				return _coll6.ToMap()
			}(), 3)
			(_nw67).ArraySet1(_372_v57, 4)
			(_nw67).ArraySet1(_372_v57, 5)
			(_nw67).ArraySet1(_372_v57, 6)
			(_nw67).ArraySet1(_372_v57, 7)
			(_nw67).ArraySet1((_372_v57).Merge(_372_v57), 8)
			(_nw67).ArraySet1((_374_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(498), _dafny.IntOfUint32((_374_v60).Cardinality()))).Uint32()).(_dafny.Map), 9)
			(_nw67).ArraySet1(_372_v57, 10)
			(_nw67).ArraySet1((_372_v57).Merge(_372_v57), 11)
			(_nw67).ArraySet1(_372_v57, 12)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _315_v7), 13)
			(_nw67).ArraySet1((func() _dafny.Map {
				if _308_v3 {
					return _372_v57
				}
				return (_374_v60).Select((Companion_Default___.SafeIndex(_315_v7, _dafny.IntOfUint32((_374_v60).Cardinality()))).Uint32()).(_dafny.Map)
			})(), 14)
			(_nw67).ArraySet1((_375_v61).Dtor_cf9(), 15)
			(_nw67).ArraySet1((_372_v57).Merge(_372_v57), 16)
			(_nw67).ArraySet1(_372_v57, 17)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_307_v2).Cardinality()), Companion_Default___.Fm0(_315_v7, (_372_v57).Cardinality(), _315_v7, _dafny.IntOfUint32((_307_v2).Cardinality()), _309_globalState)), 18)
			(_nw67).ArraySet1(_372_v57, 19)
			(_nw67).ArraySet1(_372_v57, 20)
			(_nw67).ArraySet1(Companion_Default___.Fm4(_dafny.UnicodeSeqOfUtf8Bytes("n"), true, _308_v3, _308_v3, _309_globalState), 21)
			(_nw67).ArraySet1(_372_v57, 22)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-970), _dafny.IntOfInt64(932))); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _378_v62 _dafny.Int
					_378_v62 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(-970)).Cmp(_378_v62) <= 0) && ((_378_v62).Cmp(_dafny.IntOfInt64(932)) < 0) {
						_coll7.Add((_378_v62).Minus(_315_v7), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(723))).Cardinality()))
					}
				}
				return _coll7.ToMap()
			}(), 23)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-721), _dafny.IntOfInt64(577))); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _379_v63 _dafny.Int
					_379_v63 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(-721)).Cmp(_379_v63) <= 0) && ((_379_v63).Cmp(_dafny.IntOfInt64(577)) < 0) {
						_coll8.Add((_379_v63).Times((_dafny.SetOf(_308_v3)).Cardinality()), _315_v7)
					}
				}
				return _coll8.ToMap()
			}(), 24)
			(_nw67).ArraySet1(_372_v57, 25)
			(_nw67).ArraySet1((func() _dafny.Map {
				if false {
					return Companion_Default___.Fm4(_307_v2, _308_v3, !(false), _308_v3, _309_globalState)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _dafny.IntOfUint32((_307_v2).Cardinality()))).Update(_315_v7, (_367_v52).Cardinality())
			})(), 26)
			_376_v64 = _nw67
			_376_v64 = _376_v64
			var _380_v65 *C1
			_ = _380_v65
			var _nw68 *C1 = New_C1_()
			_ = _nw68
			_nw68.Ctor__(_308_v3, (_315_v7).Times(_315_v7), _dafny.IntOfInt64(-522), _308_v3)
			_380_v65 = _nw68
		}
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_310_v4), 0))
		_ = _index55
		(_310_v4).ArraySet1(_315_v7, (_index55).Int())
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_310_v4), 0))
		_ = _index56
		(_310_v4).ArraySet1(Companion_Default___.SafeDivisionInt(_315_v7, _dafny.IntOfInt64(425)), (_index56).Int())
		var _381_v66 _dafny.Map
		_ = _381_v66
		_381_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _315_v7)
		var _382_v67 _dafny.Sequence
		_ = _382_v67
		_382_v67 = _dafny.SeqOf(_308_v3, _308_v3)
		var _383_v68 _dafny.Int
		_ = _383_v68
		var _384_v69 _dafny.Map
		_ = _384_v69
		var _out20 _dafny.Int
		_ = _out20
		var _out21 _dafny.Map
		_ = _out21
		_out20, _out21 = Companion_Default___.M0(Companion_Default___.Fm0((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int)), Companion_Default___.Fm0((_381_v66).Cardinality(), _315_v7, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(263))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_385_v4 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_386_i4 _dafny.Int) _dafny.Int {
				return (_385_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_385_v4), 0))).Int()).(_dafny.Int)
			}
		})(_310_v4)))).Cardinality()), _dafny.IntOfUint32((_382_v67).Cardinality()), _309_globalState), _dafny.IntOfUint32((_307_v2).Cardinality()), _309_globalState), _dafny.IntOfInt64(659), _309_globalState)
		_383_v68 = _out20
		_384_v69 = _out21
	}
	var _387_v70 _dafny.Map
	_ = _387_v70
	_387_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _315_v7)
	(_309_globalState).F3 = (((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_387_v70).Keys().Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _388_v71 _dafny.Int
			_388_v71 = interface{}(_compr_9).(_dafny.Int)
			if (_387_v70).Contains(_388_v71) {
				_coll9.Add(Companion_Default___.SafeModuloInt(_388_v71, _dafny.IntOfInt64(240)))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality()).Times(_315_v7)).Plus(_315_v7)
	(_309_globalState).F9 = _dafny.IntOfInt64(-648)
	var _389_v72 D0
	_ = _389_v72
	_389_v72 = Companion_D0_.Create_DC0_(true)
	var _source3 D0 = _389_v72
	_ = _source3
	if _source3.Is_DC1() {
		var _390___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC1).Cf1
		_ = _390___mcc_h0
		var _391___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC1).Cf2
		_ = _391___mcc_h1
		var _392___mcc_h2 bool = _source3.Get_().(D0_DC1).Cf3
		_ = _392___mcc_h2
		var _393_cf3 bool = _392___mcc_h2
		_ = _393_cf3
		var _394_cf2 _dafny.Int = _391___mcc_h1
		_ = _394_cf2
		var _395_cf1 _dafny.Int = _390___mcc_h0
		_ = _395_cf1
		var _396_v73 _dafny.Map
		_ = _396_v73
		_396_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _308_v3)
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_310_v4), 0))
		_ = _index57
		(_310_v4).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.SetOf((func() bool {
			if (_396_v73).Contains(_dafny.IntOfInt64(415)) {
				return (_396_v73).Get(_dafny.IntOfInt64(415)).(bool)
			}
			return _393_cf3
		})())).Cardinality(), _394_cf2), (_index57).Int())
		var _397_v74 _dafny.MultiSet
		_ = _397_v74
		_397_v74 = _dafny.MultiSetOf(_308_v3, _308_v3, _308_v3, _393_cf3)
		var _398_v75 _dafny.Sequence
		_ = _398_v75
		_398_v75 = _dafny.SeqOf(_393_cf3, _308_v3)
		var _399_v76 _dafny.Sequence
		_ = _399_v76
		_399_v76 = _dafny.SeqOf(_394_cf2)
		var _400_v77 _dafny.Sequence
		_ = _400_v77
		_400_v77 = _dafny.SeqOf(_399_v76)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_310_v4), 0))
		_ = _index58
		(_310_v4).ArraySet1((func() _dafny.Int {
			if (_dafny.MultiSetFromSeq(_398_v75)).IsSubsetOf(_397_v74) {
				return Companion_Default___.SafeModuloInt(_395_cf1, _dafny.IntOfInt64(750))
			}
			return (_dafny.IntOfInt64(987)).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_400_v77).Select((Companion_Default___.SafeIndex(_395_cf1, _dafny.IntOfUint32((_400_v77).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
		})(), (_index58).Int())
		var _401_v78 _dafny.Map
		_ = _401_v78
		_401_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _308_v3)
		var _402_v79 _dafny.Int
		_ = _402_v79
		var _403_v80 _dafny.Map
		_ = _403_v80
		var _out22 _dafny.Int
		_ = _out22
		var _out23 _dafny.Map
		_ = _out23
		_out22, _out23 = Companion_Default___.M0((_dafny.Zero).Minus(((_397_v74).Cardinality()).Minus((_401_v78).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lmuinekm")).Cardinality()), _309_globalState)
		_402_v79 = _out22
		_403_v80 = _out23
		var _404_v81 _dafny.Map
		_ = _404_v81
		_404_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _315_v7)
		var _405_v82 _dafny.Map
		_ = _405_v82
		_405_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_404_v81, _308_v3)
		_398_v75 = _dafny.Companion_Sequence_.Update(_398_v75, (Companion_Default___.SafeIndex((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_398_v75).Cardinality()))).Uint32(), (_402_v79).Cmp((_405_v82).Cardinality()) == 0)
		_401_v78 = (_401_v78).Update(false, false)
	} else {
		var _406___mcc_h3 bool = _source3.Get_().(D0_DC0).Cf0
		_ = _406___mcc_h3
		var _407_cf0 bool = _406___mcc_h3
		_ = _407_cf0
		(_309_globalState).F4 = _307_v2
		(_309_globalState).F4 = _dafny.Companion_Sequence_.Concatenate(_307_v2, _307_v2)
		var _408_v83 T1
		_ = _408_v83
		var _nw69 *C1 = New_C1_()
		_ = _nw69
		_nw69.Ctor__(!(_407_cf0), _315_v7, _315_v7, _407_cf0)
		_408_v83 = _nw69
		var _409_v84 _dafny.Map
		_ = _409_v84
		_409_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v83, _315_v7)
		var _410_v85 _dafny.Sequence
		_ = _410_v85
		_410_v85 = _dafny.SeqOf((_315_v7).Minus(_315_v7), _315_v7, _dafny.IntOfInt64(516), (_409_v84).Cardinality(), _408_v83.F20())
		_410_v85 = _410_v85
		var _411_v86 D1
		_ = _411_v86
		_411_v86 = Companion_D1_.Create_DC3_((_408_v83).F19())
		var _412_v87 _dafny.Map
		_ = _412_v87
		_412_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _411_v86)
		var _413_v88 _dafny.Sequence
		_ = _413_v88
		_413_v88 = _dafny.SeqOf(_412_v87, _412_v87, _412_v87)
		(_309_globalState).F13 = (_408_v83).Fm5(false, _dafny.IntOfUint32((_413_v88).Cardinality()), _309_globalState)
	}
	var _414_v89 *C1
	_ = _414_v89
	var _nw70 *C1 = New_C1_()
	_ = _nw70
	_nw70.Ctor__(false, (Companion_Default___.Fm8(_309_globalState)).Cardinality(), _315_v7, !(_308_v3))
	_414_v89 = _nw70
	var _415_v90 _dafny.Map
	_ = _415_v90
	_415_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v89, _307_v2)
	var _416_v91 D4
	_ = _416_v91
	_416_v91 = Companion_D4_.Create_DC9_(_414_v89)
	_415_v90 = (_415_v90).Update((_416_v91).Dtor_cf13(), _307_v2)
	var _417_v92 _dafny.CodePoint
	_ = _417_v92
	_417_v92 = _dafny.CodePoint('t')
	var _418_v93 T1
	_ = _418_v93
	var _nw71 *C1 = New_C1_()
	_ = _nw71
	_nw71.Ctor__(_308_v3, Companion_Default___.SafeModuloInt(_315_v7, _315_v7), _315_v7, _414_v89.F21)
	_418_v93 = _nw71
	var _419_v94 _dafny.Sequence
	_ = _419_v94
	_419_v94 = _dafny.SeqOf(!((_418_v93).F18()), _414_v89.F21)
	var _420_v95 _dafny.Map
	_ = _420_v95
	_420_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v94, _308_v3)
	var _421_v96 _dafny.Map
	_ = _421_v96
	_421_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_305_v0, (_418_v93).F18())
	var _422_v97 _dafny.MultiSet
	_ = _422_v97
	_422_v97 = _dafny.MultiSetOf((func() bool {
		if (_420_v95).Contains(_419_v94) {
			return (_420_v95).Get(_419_v94).(bool)
		}
		return _414_v89.F21
	})(), _308_v3, (func() bool {
		if (_421_v96).Contains(_305_v0) {
			return (_421_v96).Get(_305_v0).(bool)
		}
		return false
	})())
	var _423_v98 _dafny.Map
	_ = _423_v98
	_423_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v3, _308_v3)
	var _424_v99 _dafny.Sequence
	_ = _424_v99
	_424_v99 = _dafny.SeqOf(_423_v98, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_418_v93).F18(), _308_v3))
	var _425_v100 _dafny.Sequence
	_ = _425_v100
	_425_v100 = _dafny.SeqOf(_dafny.IntOfInt64(-186))
	var _426_v101 _dafny.Map
	_ = _426_v101
	_426_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_418_v93).F18(), (_418_v93).F19())
	var _427_v102 D0
	_ = _427_v102
	_427_v102 = Companion_D0_.Create_DC1_(_315_v7, (_425_v100).Select((Companion_Default___.SafeIndex((_426_v101).Cardinality(), _dafny.IntOfUint32((_425_v100).Cardinality()))).Uint32()).(_dafny.Int), _308_v3)
	var _428_v103 _dafny.Map
	_ = _428_v103
	_428_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _418_v93)
	var _rhs70 _dafny.CodePoint = _dafny.CodePoint('s')
	_ = _rhs70
	var _rhs71 _dafny.Sequence = (func() _dafny.Sequence {
		if false {
			return _307_v2
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("nd")
	})()
	_ = _rhs71
	var _rhs72 bool = (_418_v93).Fm5(_dafny.Companion_Sequence_.Equal(_424_v99, _dafny.Companion_Sequence_.Update(_424_v99, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v92, (_387_v70).Cardinality())).Cardinality(), _dafny.IntOfUint32((_424_v99).Cardinality()))).Uint32(), Companion_Default___.Fm9(_414_v89.F21, (_418_v93).F18(), _309_globalState))), (_427_v102).Dtor_cf2(), _309_globalState)
	_ = _rhs72
	var _rhs73 T1 = (func() T1 {
		if (_428_v103).Contains(Companion_Default___.SafeDivisionInt(_418_v93.F20(), _315_v7)) {
			return (_428_v103).Get(Companion_Default___.SafeDivisionInt(_418_v93.F20(), _315_v7)).(T1)
		}
		return _418_v93
	})()
	_ = _rhs73
	var _rhs74 _dafny.MultiSet = _422_v97
	_ = _rhs74
	_417_v92 = _rhs70
	_307_v2 = _rhs71
	_308_v3 = _rhs72
	_418_v93 = _rhs73
	_422_v97 = _rhs74
	var _429_v104 _dafny.MultiSet
	_ = _429_v104
	_429_v104 = _dafny.MultiSetOf(_414_v89)
	(_414_v89).F21 = (_308_v3) && (Companion_Default___.Fm1((func() _dafny.Int {
		if (_429_v104).Contains(_414_v89) {
			return (_429_v104).Multiplicity(_414_v89)
		}
		return (_dafny.Zero).Minus(_418_v93.F20())
	})(), _308_v3, _422_v97, _309_globalState))
	var _430_v105 D2
	_ = _430_v105
	_430_v105 = Companion_D2_.Create_DC5_(_387_v70)
	_430_v105 = Companion_Default___.Fm10(_315_v7, _309_globalState)
	if !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm2(_417_v92, _308_v3, _309_globalState), _dafny.Companion_Sequence_.Concatenate(_425_v100, _425_v100)) {
		var _431_v106 *C1
		_ = _431_v106
		var _nw72 *C1 = New_C1_()
		_ = _nw72
		_nw72.Ctor__((_418_v93).F18(), (_418_v93).F19(), _315_v7, _414_v89.F21)
		_431_v106 = _nw72
		(_309_globalState).F13 = (_418_v93).F18()
		if (_418_v93).Fm5(_308_v3, _315_v7, _309_globalState) {
			(_309_globalState).F3 = (_dafny.Zero).Minus((_418_v93).F19())
			(_414_v89).F21 = false
			(_414_v89).F21 = _431_v106.F21
			var _432_v107 _dafny.MultiSet
			_ = _432_v107
			_432_v107 = _dafny.MultiSetOf(_418_v93.F20())
			var _433_v108 D2
			_ = _433_v108
			_433_v108 = Companion_D2_.Create_DC6_(_432_v107)
			var _434_v109 _dafny.Set
			_ = _434_v109
			_434_v109 = _dafny.SetOf((_418_v93).F19(), (_418_v93).F19(), _315_v7)
			var _435_v110 T0
			_ = _435_v110
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__(Companion_D2_.Create_DC6_(_dafny.MultiSetOf((_434_v109).Cardinality(), (_418_v93).F19(), _418_v93.F20())), (_418_v93).F19(), (_418_v93).F18())
			_435_v110 = _nw73
			var _436_v111 _dafny.Sequence
			_ = _436_v111
			_436_v111 = _dafny.SeqOf(_435_v110, _435_v110, _435_v110, _435_v110, _435_v110)
			var _437_v112 *C0
			_ = _437_v112
			var _nw74 *C0 = New_C0_()
			_ = _nw74
			_nw74.Ctor__(_433_v108, _418_v93.F20(), (_414_v89).Fm5(_414_v89.F21, _dafny.IntOfUint32((_436_v111).Cardinality()), _309_globalState))
			_437_v112 = _nw74
			var _438_v113 _dafny.Sequence
			_ = _438_v113
			_438_v113 = _dafny.SeqOf(_437_v112)
			var _439_v114 *C1
			_ = _439_v114
			var _nw75 *C1 = New_C1_()
			_ = _nw75
			_nw75.Ctor__(_414_v89.F21, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_438_v113, _438_v113)).Cardinality()), (_315_v7).Times((_418_v93).F19()), ((_437_v112).F23()).Cmp((Companion_Default___.Fm11(_309_globalState)).Dtor_cf5()) > 0)
			_439_v114 = _nw75
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_310_v4), 0))
			_ = _index59
			(_310_v4).ArraySet1(_418_v93.F20(), (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_310_v4), 0))
			_ = _index60
			var _rhs75 _dafny.Int = (_418_v93.F20()).Times(_418_v93.F20())
			_ = _rhs75
			var _rhs76 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-158), _dafny.IntOfInt64(654)), _418_v93.F20())
			_ = _rhs76
			var _rhs77 bool = (_dafny.SetOf((_418_v93).F19(), (_422_v97).Cardinality())).IsDisjointFrom((_434_v109).Intersection(_434_v109))
			_ = _rhs77
			var _rhs78 _dafny.Array = _305_v0
			_ = _rhs78
			var _rhs79 bool = false
			_ = _rhs79
			var _lhs55 _dafny.Array = _310_v4
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_310_v4), 0))
			_ = _lhs56
			var _lhs57 *GlobalState = _309_globalState
			_ = _lhs57
			var _lhs58 *GlobalState = _309_globalState
			_ = _lhs58
			var _lhs59 *GlobalState = _309_globalState
			_ = _lhs59
			(_lhs55).ArraySet1(_rhs75, (_lhs56).Int())
			_lhs57.F9 = _rhs76
			_lhs58.F13 = _rhs77
			_305_v0 = _rhs78
			_lhs59.F13 = _rhs79
		} else {
			var _440_v115 _dafny.Array
			_ = _440_v115
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_15
			var _nw76 _dafny.Array
			_ = _nw76
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw76 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Map = (func(_441_v101 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_442_i5 _dafny.Int) _dafny.Map {
						return _441_v101
					}
				})(_426_v101)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw76 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw76).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw76).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_440_v115 = _nw76
			var _443_v116 _dafny.Map
			_ = _443_v116
			_443_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v94, _440_v115)
			_443_v116 = (_443_v116).Update(_dafny.SeqOf((_418_v93).F18(), (_418_v93).F18(), (_418_v93).F18(), _414_v89.F21, !(_431_v106.F21)), _440_v115)
			var _444_v117 D1
			_ = _444_v117
			_444_v117 = Companion_D1_.Create_DC2_(_dafny.SeqOf((_418_v93).F18(), _414_v89.F21, _431_v106.F21))
			var _445_v118 D2
			_ = _445_v118
			_445_v118 = Companion_D2_.Create_DC6_(_dafny.MultiSetOf(_dafny.IntOfInt64(115), (_418_v93).F19()))
			var _446_v119 _dafny.MultiSet
			_ = _446_v119
			_446_v119 = _dafny.MultiSetOf((_418_v93).F19())
			var _pat_let_tv11 = _446_v119
			_ = _pat_let_tv11
			var _447_v120 *C0
			_ = _447_v120
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(func(_pat_let6_0 D2) D2 {
				return func(_448_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let7_0 _dafny.MultiSet) D2 {
						return func(_449_dt__update_hcf10_h0 _dafny.MultiSet) D2 {
							return Companion_D2_.Create_DC6_(_449_dt__update_hcf10_h0)
						}(_pat_let7_0)
					}(_pat_let_tv11)
				}(_pat_let6_0)
			}(_445_v118), (_418_v93).F19(), _431_v106.F21)
			_447_v120 = _nw77
			var _450_v121 _dafny.Map
			_ = _450_v121
			_450_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_447_v120, _418_v93.F20())
			var _451_v122 _dafny.Map
			_ = _451_v122
			_451_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v117, _450_v121)
			_451_v122 = _451_v122
			_447_v120 = _447_v120
			var _452_v123 _dafny.Map
			_ = _452_v123
			_452_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v106, _418_v93.F20())
			var _453_v124 _dafny.Array
			_ = _453_v124
			var _454_v125 _dafny.Set
			_ = _454_v125
			var _out24 _dafny.Array
			_ = _out24
			var _out25 _dafny.Set
			_ = _out25
			_out24, _out25 = (_418_v93).M2((_446_v119).Difference(_dafny.MultiSetOf(_315_v7, _418_v93.F20())), (func() _dafny.Int {
				if (_452_v123).Contains(_414_v89) {
					return (_452_v123).Get(_414_v89).(_dafny.Int)
				}
				return (_418_v93).F19()
			})(), _306_v1, _307_v2, _309_globalState)
			_453_v124 = _out24
			_454_v125 = _out25
			var _455_v126 _dafny.Map
			_ = _455_v126
			_455_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v89.F21, _306_v1)
			_455_v126 = (_455_v126).Update(_414_v89.F21, _306_v1)
		}
		var _456_v127 _dafny.Set
		_ = _456_v127
		_456_v127 = _dafny.SetOf(_306_v1, _306_v1)
		if (((_418_v93).F19()).Times((_456_v127).Cardinality())).Cmp(Companion_Default___.SafeDivisionInt((_418_v93).F19(), (_418_v93).F19())) == 0 {
			var _457_v128 _dafny.Array
			_ = _457_v128
			var _nw78 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
			_ = _nw78
			_457_v128 = _nw78
			var _458_v129 _dafny.Map
			_ = _458_v129
			_458_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_457_v128, _310_v4)
			var _459_v130 _dafny.Set
			_ = _459_v130
			_459_v130 = _dafny.SetOf(_dafny.CodePoint('v'), _dafny.CodePoint('d'))
			var _460_v131 _dafny.MultiSet
			_ = _460_v131
			_460_v131 = _dafny.MultiSetOf(Companion_Default___.Fm0(_418_v93.F20(), _418_v93.F20(), (_418_v93).F19(), (_459_v130).Cardinality(), _309_globalState))
			var _461_v132 *C1
			_ = _461_v132
			var _nw79 *C1 = New_C1_()
			_ = _nw79
			_nw79.Ctor__(_414_v89.F21, (_418_v93).F19(), Companion_Default___.Fm0(_418_v93.F20(), (_418_v93).F19(), (_458_v129).Cardinality(), _315_v7, _309_globalState), (Companion_D0_.Create_DC1_((_460_v131).Cardinality(), (_418_v93).F19(), _431_v106.F21)).Dtor_cf3())
			_461_v132 = _nw79
			(_309_globalState).F9 = _315_v7
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_306_v1), 0))
			_ = _index61
			(_306_v1).ArraySet1(!(_414_v89.F21), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_306_v1), 0))
			_ = _index62
			(_306_v1).ArraySet1(((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jfbbox")).Cardinality())).Plus(_418_v93.F20())).Cmp((_418_v93).F19()) < 0, (_index62).Int())
			_387_v70 = (_387_v70).Update(_418_v93.F20(), (_dafny.Zero).Minus((_423_v98).Cardinality()))
			_307_v2 = (func() _dafny.Sequence {
				if (_415_v90).Contains(_431_v106) {
					return (_415_v90).Get(_431_v106).(_dafny.Sequence)
				}
				return _307_v2
			})()
		} else {
			(_309_globalState).F5 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_418_v93.F20(), ((_422_v97).Difference(_422_v97)).Cardinality()))
			(_414_v89).F21 = ((_dafny.MultiSetOf((_418_v93).F19(), _418_v93.F20())).Cardinality()).Cmp((_dafny.Zero).Minus(_315_v7)) != 0
			_308_v3 = (_387_v70).Contains(_418_v93.F20())
			var _462_v133 _dafny.Int
			_ = _462_v133
			var _463_v134 bool
			_ = _463_v134
			var _464_v135 _dafny.Int
			_ = _464_v135
			var _465_v136 _dafny.Int
			_ = _465_v136
			var _out26 _dafny.Int
			_ = _out26
			var _out27 bool
			_ = _out27
			var _out28 _dafny.Int
			_ = _out28
			var _out29 _dafny.Int
			_ = _out29
			_out26, _out27, _out28, _out29 = (_414_v89).M1(_431_v106.F21, (_431_v106.F21) || (false), _309_globalState)
			_462_v133 = _out26
			_463_v134 = _out27
			_464_v135 = _out28
			_465_v136 = _out29
			var _rhs80 _dafny.Int = (_418_v93).F19()
			_ = _rhs80
			var _rhs81 _dafny.Int = (_dafny.Zero).Minus(_418_v93.F20())
			_ = _rhs81
			var _lhs60 *GlobalState = _309_globalState
			_ = _lhs60
			_464_v135 = _rhs80
			_lhs60.F3 = _rhs81
		}
		(_418_v93).F20_set_(_dafny.IntOfInt64(-525))
	} else {
		var _466_v137 _dafny.Sequence
		_ = _466_v137
		_466_v137 = _dafny.SeqOf((func() _dafny.Sequence {
			if (_418_v93).F18() {
				return _419_v94
			}
			return _dafny.SeqOf((_418_v93).F18())
		})(), _dafny.SeqOf(_414_v89.F21))
		var _467_v138 _dafny.Array
		_ = _467_v138
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw80
		_467_v138 = _nw80
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_467_v138), 0))
		_ = _index63
		(_467_v138).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_419_v94, _419_v94), (_index63).Int())
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_467_v138), 0))
		_ = _index64
		var _rhs82 _dafny.Int = (((_418_v93).F19()).Plus((_418_v93).F19())).Plus(_315_v7)
		_ = _rhs82
		var _rhs83 bool = _414_v89.F21
		_ = _rhs83
		var _rhs84 _dafny.Sequence = _466_v137
		_ = _rhs84
		var _rhs85 _dafny.Sequence = _419_v94
		_ = _rhs85
		var _rhs86 bool = (func() bool {
			if !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_468_v92 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_469_i6 _dafny.Int) _dafny.CodePoint {
					return _468_v92
				}
			})(_417_v92))), _307_v2) {
				return (_418_v93.F20()).Cmp((_418_v93).F19()) >= 0
			}
			return ((_418_v93).F18()) || ((_418_v93).F18())
		})()
		_ = _rhs86
		var _lhs61 *GlobalState = _309_globalState
		_ = _lhs61
		var _lhs62 *C1 = _414_v89
		_ = _lhs62
		var _lhs63 _dafny.Array = _467_v138
		_ = _lhs63
		var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_467_v138), 0))
		_ = _lhs64
		var _lhs65 *C1 = _414_v89
		_ = _lhs65
		_lhs61.F9 = _rhs82
		_lhs62.F21 = _rhs83
		_466_v137 = _rhs84
		(_lhs63).ArraySet1(_rhs85, (_lhs64).Int())
		_lhs65.F21 = _rhs86
		var _470_v139 D1
		_ = _470_v139
		_470_v139 = Companion_D1_.Create_DC3_((_418_v93).F19())
		var _pat_let_tv12 = _418_v93
		_ = _pat_let_tv12
		_470_v139 = func(_pat_let8_0 D1) D1 {
			return func(_471_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let9_0 _dafny.Int) D1 {
					return func(_472_dt__update_hcf5_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC3_(_472_dt__update_hcf5_h0)
					}(_pat_let9_0)
				}(_pat_let_tv12.F20())
			}(_pat_let8_0)
		}(_470_v139)
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_467_v138), 0))
		_ = _index65
		var _rhs87 _dafny.Int = (func() _dafny.Int {
			if _414_v89.F21 {
				return _dafny.IntOfUint32((_425_v100).Cardinality())
			}
			return (_418_v93).F19()
		})()
		_ = _rhs87
		var _rhs88 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_307_v2, _307_v2)
		_ = _rhs88
		var _rhs89 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_419_v94, _dafny.Companion_Sequence_.Concatenate((_467_v138).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_467_v138), 0))).Int()).(_dafny.Sequence), _419_v94))
		_ = _rhs89
		var _rhs90 _dafny.Sequence = _307_v2
		_ = _rhs90
		var _lhs66 *GlobalState = _309_globalState
		_ = _lhs66
		var _lhs67 *C1 = _414_v89
		_ = _lhs67
		var _lhs68 _dafny.Array = _467_v138
		_ = _lhs68
		var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_467_v138), 0))
		_ = _lhs69
		_lhs66.F3 = _rhs87
		_lhs67.F21 = _rhs88
		(_lhs68).ArraySet1(_rhs89, (_lhs69).Int())
		_307_v2 = _rhs90
		if _414_v89.F21 {
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_306_v1), 0))
			_ = _index66
			(_306_v1).ArraySet1(!((_418_v93).F18()), (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_306_v1), 0))
			_ = _index67
			(_306_v1).ArraySet1(_308_v3, (_index67).Int())
			var _473_v140 _dafny.Array
			_ = _473_v140
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw81
			_473_v140 = _nw81
			(_309_globalState).F2 = _473_v140
			var _474_v141 _dafny.Array
			_ = _474_v141
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_16
			var _nw82 _dafny.Array
			_ = _nw82
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw82 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) D0 = (func(_475_v102 D0) func(_dafny.Int) D0 {
					return func(_476_i7 _dafny.Int) D0 {
						return _475_v102
					}
				})(_427_v102)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw82 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw82).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw82).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_474_v141 = _nw82
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_474_v141), 0))
			_ = _index68
			(_474_v141).ArraySet1(_427_v102, (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_474_v141), 0))
			_ = _index69
			(_474_v141).ArraySet1(Companion_D0_.Create_DC1_((_418_v93).F19(), Companion_Default___.Fm0((_418_v93).F19(), (_418_v93).F19(), _418_v93.F20(), (_418_v93).F19(), _309_globalState), ((_387_v70).Cardinality()).Cmp(_418_v93.F20()) <= 0), (_index69).Int())
			var _477_v142 _dafny.Array
			_ = _477_v142
			var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw83
			_477_v142 = _nw83
			_477_v142 = _477_v142
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_310_v4), 0))
			_ = _index70
			(_310_v4).ArraySet1((_418_v93).F19(), (_index70).Int())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_310_v4), 0))
			_ = _index71
			(_310_v4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_424_v99, _424_v99)).Cardinality()), (_index71).Int())
		} else {
			(_309_globalState).F3 = (_418_v93).F19()
			var _478_v143 _dafny.Int
			_ = _478_v143
			var _479_v144 _dafny.Map
			_ = _479_v144
			var _out30 _dafny.Int
			_ = _out30
			var _out31 _dafny.Map
			_ = _out31
			_out30, _out31 = Companion_Default___.M0(_418_v93.F20(), _dafny.IntOfUint32((Companion_Default___.Fm2(_417_v92, _414_v89.F21, _309_globalState)).Cardinality()), _309_globalState)
			_478_v143 = _out30
			_479_v144 = _out31
			(_309_globalState).F9 = _478_v143
			var _rhs91 bool = _414_v89.F21
			_ = _rhs91
			var _rhs92 _dafny.Array = _310_v4
			_ = _rhs92
			var _rhs93 _dafny.Int = _478_v143
			_ = _rhs93
			var _lhs70 *GlobalState = _309_globalState
			_ = _lhs70
			_308_v3 = _rhs91
			_310_v4 = _rhs92
			_lhs70.F3 = _rhs93
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_306_v1), 0))
			_ = _index72
			(_306_v1).ArraySet1((_418_v93).F18(), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_306_v1), 0))
			_ = _index73
			(_306_v1).ArraySet1(_414_v89.F21, (_index73).Int())
		}
		var _480_v145 _dafny.Map
		_ = _480_v145
		_480_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v94, (_dafny.Zero).Minus((_418_v93).F19()))
		var _481_v146 _dafny.MultiSet
		_ = _481_v146
		_481_v146 = _dafny.MultiSetOf(_dafny.IntOfUint32((_425_v100).Cardinality()), (_480_v145).Cardinality())
		var _482_v147 _dafny.Array
		_ = _482_v147
		var _483_v148 _dafny.Set
		_ = _483_v148
		var _out32 _dafny.Array
		_ = _out32
		var _out33 _dafny.Set
		_ = _out33
		_out32, _out33 = (_414_v89).M2(_481_v146, (_418_v93).F19(), _306_v1, Companion_Default___.Fm3(_309_globalState), _309_globalState)
		_482_v147 = _out32
		_483_v148 = _out33
	}
	var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
	_ = _index74
	(_310_v4).ArraySet1((_418_v93).F19(), (_index74).Int())
	var _484_v149 _dafny.Array
	_ = _484_v149
	var _len0_17 _dafny.Int = _dafny.IntOfInt64(28)
	_ = _len0_17
	var _nw84 _dafny.Array
	_ = _nw84
	if _len0_17.Cmp(_dafny.Zero) == 0 {
		_nw84 = _dafny.NewArray(_len0_17)
	} else {
		var _init17 func(_dafny.Int) _dafny.Map = (func(_485_v101 _dafny.Map) func(_dafny.Int) _dafny.Map {
			return func(_486_i8 _dafny.Int) _dafny.Map {
				return _485_v101
			}
		})(_426_v101)
		_ = _init17
		var _element0_17 = _init17(_dafny.Zero)
		_ = _element0_17
		_nw84 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
		(_nw84).ArraySet1(_element0_17, 0)
		var _nativeLen0_17 = (_len0_17).Int()
		_ = _nativeLen0_17
		for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
			(_nw84).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
		}
	}
	_484_v149 = _nw84
	var _487_v150 _dafny.Map
	_ = _487_v150
	_487_v150 = _426_v101
	var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_484_v149), 0))
	_ = _index75
	(_484_v149).ArraySet1((func() _dafny.Map {
		if !(_414_v89.F21) {
			return _487_v150
		}
		return _487_v150
	})(), (_index75).Int())
	var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
	_ = _index76
	var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_484_v149), 0))
	_ = _index77
	var _rhs94 _dafny.Int = _418_v93.F20()
	_ = _rhs94
	var _rhs95 _dafny.Map = _487_v150
	_ = _rhs95
	var _rhs96 bool = (Companion_Default___.Fm0(_418_v93.F20(), (_dafny.Zero).Minus(_418_v93.F20()), Companion_Default___.Fm0(_dafny.IntOfInt64(-511), _315_v7, _418_v93.F20(), _315_v7, _309_globalState), _418_v93.F20(), _309_globalState)).Cmp((func() _dafny.Int {
		if _308_v3 {
			return _dafny.IntOfUint32((_419_v94).Cardinality())
		}
		return _418_v93.F20()
	})()) < 0
	_ = _rhs96
	var _lhs71 _dafny.Array = _310_v4
	_ = _lhs71
	var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
	_ = _lhs72
	var _lhs73 _dafny.Array = _484_v149
	_ = _lhs73
	var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_484_v149), 0))
	_ = _lhs74
	var _lhs75 *C1 = _414_v89
	_ = _lhs75
	(_lhs71).ArraySet1(_rhs94, (_lhs72).Int())
	(_lhs73).ArraySet1(_rhs95, (_lhs74).Int())
	_lhs75.F21 = _rhs96
	var _pat_let_tv13 = _417_v92
	_ = _pat_let_tv13
	var _pat_let_tv14 = _417_v92
	_ = _pat_let_tv14
	var _pat_let_tv15 = _417_v92
	_ = _pat_let_tv15
	_417_v92 = func(_source4 D2) _dafny.CodePoint {
		if _source4.Is_DC6() {
			var _488___mcc_h4 _dafny.MultiSet = _source4.Get_().(D2_DC6).Cf10
			_ = _488___mcc_h4
			var _489_cf10 _dafny.MultiSet = _488___mcc_h4
			_ = _489_cf10
			return _pat_let_tv13
		} else if _source4.Is_DC5() {
			var _490___mcc_h5 _dafny.Map = _source4.Get_().(D2_DC5).Cf9
			_ = _490___mcc_h5
			var _491_cf9 _dafny.Map = _490___mcc_h5
			_ = _491_cf9
			return _pat_let_tv14
		} else {
			var _492___mcc_h6 D2 = _source4.Get_().(D2_DC7).Cf11
			_ = _492___mcc_h6
			var _493_cf11 D2 = _492___mcc_h6
			_ = _493_cf11
			return _pat_let_tv15
		}
	}((func() D2 {
		if true {
			return _430_v105
		}
		return _430_v105
	})())
	(_309_globalState).F3 = _dafny.IntOfInt64(440)
	var _494_v151 _dafny.Array
	_ = _494_v151
	var _495_v152 _dafny.Set
	_ = _495_v152
	var _out34 _dafny.Array
	_ = _out34
	var _out35 _dafny.Set
	_ = _out35
	_out34, _out35 = (_414_v89).M2(Companion_Default___.Fm12(_309_globalState), (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _306_v1, (func() _dafny.Sequence {
		if _414_v89.F21 {
			return _dafny.UnicodeSeqOfUtf8Bytes("xmgvn")
		}
		return _307_v2
	})(), _309_globalState)
	_494_v151 = _out34
	_495_v152 = _out35
	if ((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int)).Cmp(_418_v93.F20()) < 0 {
		var _496_v153 _dafny.MultiSet
		_ = _496_v153
		_496_v153 = _dafny.MultiSetOf(_418_v93.F20())
		var _497_v154 _dafny.MultiSet
		_ = _497_v154
		_497_v154 = _dafny.MultiSetOf((_418_v93).F19(), Companion_Default___.Fm0((_496_v153).Cardinality(), (_418_v93).F19(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_307_v2).Cardinality())), (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _309_globalState))
		(_418_v93).F20_set_(Companion_Default___.Fm0(_315_v7, _dafny.IntOfInt64(790), (func() _dafny.Int {
			if (_497_v154).Contains(_dafny.IntOfUint32((_425_v100).Cardinality())) {
				return (_497_v154).Multiplicity(_dafny.IntOfUint32((_425_v100).Cardinality()))
			}
			return Companion_Default___.Fm0(_315_v7, _315_v7, _dafny.IntOfInt64(250), _dafny.IntOfUint32((_425_v100).Cardinality()), _309_globalState)
		})(), (_418_v93).F19(), _309_globalState))
		var _498_v155 _dafny.Map
		_ = _498_v155
		_498_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v89, _414_v89.F21)
		var _499_v156 D1
		_ = _499_v156
		_499_v156 = Companion_D1_.Create_DC3_((_495_v152).Cardinality())
		var _source5 D1 = (func() D1 {
			if (_498_v155).Equals(_498_v155) {
				return _499_v156
			}
			return _499_v156
		})()
		_ = _source5
		if _source5.Is_DC3() {
			var _500___mcc_h7 _dafny.Int = _source5.Get_().(D1_DC3).Cf5
			_ = _500___mcc_h7
			var _501_cf5 _dafny.Int = _500___mcc_h7
			_ = _501_cf5
			var _502_v157 D5
			_ = _502_v157
			_502_v157 = Companion_D5_.Create_DC12_(_418_v93)
			var _503_v158 _dafny.Map
			_ = _503_v158
			_503_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v89.F21, _418_v93)
			var _504_v159 _dafny.Array
			_ = _504_v159
			var _nwElement0_14 T1 = _418_v93
			_ = _nwElement0_14
			var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(18))
			_ = _nw85
			(_nw85).ArraySet1(_nwElement0_14, 0)
			(_nw85).ArraySet1(_418_v93, 1)
			(_nw85).ArraySet1(_418_v93, 2)
			(_nw85).ArraySet1(_418_v93, 3)
			(_nw85).ArraySet1(_418_v93, 4)
			(_nw85).ArraySet1(_418_v93, 5)
			(_nw85).ArraySet1((_502_v157).Dtor_cf18(), 6)
			(_nw85).ArraySet1(_418_v93, 7)
			(_nw85).ArraySet1((func() T1 {
				if (_503_v158).Contains(_414_v89.F21) {
					return (_503_v158).Get(_414_v89.F21).(T1)
				}
				return _418_v93
			})(), 8)
			(_nw85).ArraySet1(_418_v93, 9)
			(_nw85).ArraySet1(_418_v93, 10)
			(_nw85).ArraySet1(_418_v93, 11)
			(_nw85).ArraySet1(_418_v93, 12)
			(_nw85).ArraySet1(_418_v93, 13)
			(_nw85).ArraySet1(_418_v93, 14)
			(_nw85).ArraySet1(_418_v93, 15)
			(_nw85).ArraySet1(_418_v93, 16)
			(_nw85).ArraySet1((func() T1 {
				if _414_v89.F21 {
					return _418_v93
				}
				return _418_v93
			})(), 17)
			_504_v159 = _nw85
			var _rhs97 _dafny.Array = _504_v159
			_ = _rhs97
			var _rhs98 _dafny.Sequence = _307_v2
			_ = _rhs98
			var _rhs99 *C1 = _414_v89
			_ = _rhs99
			var _lhs76 *GlobalState = _309_globalState
			_ = _lhs76
			_504_v159 = _rhs97
			_lhs76.F4 = _rhs98
			_414_v89 = _rhs99
			(_309_globalState).F3 = _501_cf5
			(_414_v89).F21 = _308_v3
			(_309_globalState).F4 = (func() _dafny.Sequence {
				if _414_v89.F21 {
					return _307_v2
				}
				return _307_v2
			})()
		} else if _source5.Is_DC4() {
			var _505___mcc_h8 bool = _source5.Get_().(D1_DC4).Cf6
			_ = _505___mcc_h8
			var _506___mcc_h9 bool = _source5.Get_().(D1_DC4).Cf7
			_ = _506___mcc_h9
			var _507___mcc_h10 bool = _source5.Get_().(D1_DC4).Cf8
			_ = _507___mcc_h10
			var _508_cf8 bool = _507___mcc_h10
			_ = _508_cf8
			var _509_cf7 bool = _506___mcc_h9
			_ = _509_cf7
			var _510_cf6 bool = _505___mcc_h8
			_ = _510_cf6
			_426_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_418_v93).F18(), (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int))
			_509_cf7 = (_418_v93).F18()
			_510_cf6 = _509_cf7
			var _511_v160 _dafny.Sequence
			_ = _511_v160
			_511_v160 = _dafny.SeqOf(_307_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(410))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_512_v92 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_513_i9 _dafny.Int) _dafny.CodePoint {
					return _512_v92
				}
			})(_417_v92))), _307_v2)
			_419_v94 = _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_308_v3) == (_414_v89.F21), _308_v3), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_511_v160).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_308_v3) == (_414_v89.F21), _308_v3)).Cardinality()))).Uint32(), ((Companion_Default___.Fm13((_418_v93).F18(), (func() bool {
				if (_423_v98).Contains(_308_v3) {
					return (_423_v98).Get(_308_v3).(bool)
				}
				return _508_cf8
			})(), _418_v93.F20(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(389))).Cardinality()), _309_globalState)).Cardinality()).Cmp(_418_v93.F20()) < 0)
		} else {
			var _514___mcc_h11 _dafny.Sequence = _source5.Get_().(D1_DC2).Cf4
			_ = _514___mcc_h11
			var _515_cf4 _dafny.Sequence = _514___mcc_h11
			_ = _515_cf4
			(_414_v89).F21 = !(_dafny.Companion_Sequence_.IsPrefixOf(_307_v2, _307_v2))
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _index78
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _index79
			var _rhs100 _dafny.Int = _418_v93.F20()
			_ = _rhs100
			var _rhs101 D0 = _389_v72
			_ = _rhs101
			var _rhs102 _dafny.Int = (_418_v93.F20()).Minus(_315_v7)
			_ = _rhs102
			var _lhs77 _dafny.Array = _310_v4
			_ = _lhs77
			var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _lhs78
			var _lhs79 _dafny.Array = _310_v4
			_ = _lhs79
			var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _lhs80
			(_lhs77).ArraySet1(_rhs100, (_lhs78).Int())
			_389_v72 = _rhs101
			(_lhs79).ArraySet1(_rhs102, (_lhs80).Int())
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _index80
			(_310_v4).ArraySet1(((_dafny.Zero).Minus((_418_v93).F19())).Plus(_418_v93.F20()), (_index80).Int())
			var _516_v161 _dafny.Sequence
			_ = _516_v161
			_516_v161 = _dafny.SeqOf(_307_v2)
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _index81
			var _rhs103 _dafny.Array = _310_v4
			_ = _rhs103
			var _rhs104 bool = _414_v89.F21
			_ = _rhs104
			var _rhs105 _dafny.Sequence = _516_v161
			_ = _rhs105
			var _rhs106 _dafny.Int = _dafny.IntOfInt64(88)
			_ = _rhs106
			var _rhs107 _dafny.CodePoint = _417_v92
			_ = _rhs107
			var _lhs81 _dafny.Array = _310_v4
			_ = _lhs81
			var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _lhs82
			_310_v4 = _rhs103
			_308_v3 = _rhs104
			_516_v161 = _rhs105
			(_lhs81).ArraySet1(_rhs106, (_lhs82).Int())
			_417_v92 = _rhs107
		}
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_306_v1), 0))
		_ = _index82
		(_306_v1).ArraySet1((!(_414_v89.F21)) == ((_418_v93).F18()), (_index82).Int())
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_306_v1), 0))
		_ = _index83
		(_306_v1).ArraySet1(_414_v89.F21, (_index83).Int())
		var _517_v162 *C1
		_ = _517_v162
		var _nw86 *C1 = New_C1_()
		_ = _nw86
		_nw86.Ctor__((func() bool {
			if true {
				return (_418_v93).F18()
			}
			return (Companion_D0_.Create_DC0_(_308_v3)).Dtor_cf0()
		})(), _418_v93.F20(), _418_v93.F20(), _414_v89.F21)
		_517_v162 = _nw86
		if (_dafny.IntOfUint32((_307_v2).Cardinality())).Cmp((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int)) > 0 {
			(_309_globalState).F5 = _315_v7
			(_418_v93).F20_set_(((_418_v93).F19()).Times(_418_v93.F20()))
			var _518_v163 D2
			_ = _518_v163
			_518_v163 = Companion_D2_.Create_DC6_(_497_v154)
			var _519_v164 D2
			_ = _519_v164
			_519_v164 = Companion_D2_.Create_DC6_((_518_v163).Dtor_cf10())
			var _520_v165 T0
			_ = _520_v165
			var _nw87 *C0 = New_C0_()
			_ = _nw87
			_nw87.Ctor__(_519_v164, _418_v93.F20(), (_418_v93).F18())
			_520_v165 = _nw87
			var _521_v166 _dafny.Sequence
			_ = _521_v166
			_521_v166 = _dafny.SeqOf(_520_v165, _520_v165, _520_v165)
			var _522_v167 _dafny.Map
			_ = _522_v167
			_522_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_521_v166).Select((Companion_Default___.SafeIndex((_418_v93).F19(), _dafny.IntOfUint32((_521_v166).Cardinality()))).Uint32()).(T0), _414_v89.F21)
			_522_v167 = (_522_v167).Update(_520_v165, (_418_v93).F18())
			var _523_v168 D2
			_ = _523_v168
			_523_v168 = Companion_D2_.Create_DC5_(_387_v70)
			var _524_v169 D2
			_ = _524_v169
			_524_v169 = Companion_D2_.Create_DC7_(_523_v168)
			var _525_v170 D2
			_ = _525_v170
			_525_v170 = Companion_D2_.Create_DC7_(_523_v168)
			var _526_v171 _dafny.Map
			_ = _526_v171
			_526_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_525_v170, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_517_v162.F21, _dafny.IntOfUint32((_307_v2).Cardinality())))
			_526_v171 = (_526_v171).Update((func() D2 {
				if false {
					return _525_v170
				}
				return _525_v170
			})(), _426_v101)
			var _527_v172 _dafny.Map
			_ = _527_v172
			_527_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_418_v93.F20(), (_418_v93).F18())
			_527_v172 = (_527_v172).Update(Companion_Default___.SafeDivisionInt(_418_v93.F20(), (_418_v93).F19()), !(!(Companion_Default___.Fm1((_418_v93).F19(), _517_v162.F21, (_dafny.MultiSetOf((_418_v93).F18())).Update((_418_v93).Fm5(_517_v162.F21, _418_v93.F20(), _309_globalState), Companion_Default___.Abs((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int))), _309_globalState))))
		} else {
			var _528_v173 _dafny.Array
			_ = _528_v173
			var _nw88 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw88
			_528_v173 = _nw88
			var _529_v174 _dafny.Array
			_ = _529_v174
			var _530_v175 _dafny.Set
			_ = _530_v175
			var _out36 _dafny.Array
			_ = _out36
			var _out37 _dafny.Set
			_ = _out37
			_out36, _out37 = (_418_v93).M2((_497_v154).Update((_dafny.Zero).Minus(_418_v93.F20()), Companion_Default___.Abs(_418_v93.F20())), _dafny.IntOfUint32((_dafny.SeqOf(_414_v89.F21)).Cardinality()), _528_v173, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_531_v92 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_532_i10 _dafny.Int) _dafny.CodePoint {
					return _531_v92
				}
			})(_417_v92))), _309_globalState)
			_529_v174 = _out36
			_530_v175 = _out37
			var _533_v176 _dafny.Array
			_ = _533_v176
			var _nw89 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
			_ = _nw89
			_533_v176 = _nw89
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_533_v176), 0))
			_ = _index84
			(_533_v176).ArraySet1(_414_v89, (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_533_v176), 0))
			_ = _index85
			(_533_v176).ArraySet1(_517_v162, (_index85).Int())
			var _534_v177 _dafny.Map
			_ = _534_v177
			_534_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_427_v102).Dtor_cf3() {
					return _307_v2
				}
				return _307_v2
			})(), _414_v89.F21)
			_534_v177 = (_534_v177).Update(_dafny.Companion_Sequence_.Concatenate(_307_v2, _307_v2), _414_v89.F21)
			var _535_v178 D2
			_ = _535_v178
			_535_v178 = Companion_D2_.Create_DC6_(Companion_Default___.Fm12(_309_globalState))
			var _536_v179 _dafny.Map
			_ = _536_v179
			_536_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_315_v7, _dafny.MultiSetFromSeq(_419_v94))
			var _537_v180 *C0
			_ = _537_v180
			var _nw90 *C0 = New_C0_()
			_ = _nw90
			_nw90.Ctor__(Companion_D2_.Create_DC6_(_497_v154), (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), Companion_Default___.Fm1(_dafny.IntOfUint32((_307_v2).Cardinality()), _414_v89.F21, (func() _dafny.MultiSet {
				if (_536_v179).Contains(_315_v7) {
					return (_536_v179).Get(_315_v7).(_dafny.MultiSet)
				}
				return _422_v97
			})(), _309_globalState))
			_537_v180 = _nw90
			var _538_v181 _dafny.Map
			_ = _538_v181
			_538_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_537_v180, !((_418_v93).F18()))
			var _539_v182 T0
			_ = _539_v182
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__(_535_v178, ((_538_v181).Cardinality()).Minus((_422_v97).Cardinality()), !(_414_v89.F21))
			_539_v182 = _nw91
			_539_v182 = _539_v182
			(_309_globalState).F13 = _517_v162.F21
		}
	} else {
		if (func() bool {
			if _414_v89.F21 {
				return _308_v3
			}
			return false
		})() {
			var _540_v183 _dafny.Array
			_ = _540_v183
			var _541_v184 _dafny.Set
			_ = _541_v184
			var _out38 _dafny.Array
			_ = _out38
			var _out39 _dafny.Set
			_ = _out39
			_out38, _out39 = (_414_v89).M2(_dafny.MultiSetOf(_dafny.IntOfInt64(608)), _418_v93.F20(), _306_v1, _307_v2, _309_globalState)
			_540_v183 = _out38
			_541_v184 = _out39
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_18
			var _nw92 _dafny.Array
			_ = _nw92
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw92 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = (func(_542_v89 *C1) func(_dafny.Int) bool {
					return func(_543_i11 _dafny.Int) bool {
						return _542_v89.F21
					}
				})(_414_v89)
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw92 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw92).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw92).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_306_v1 = _nw92
			var _544_v185 _dafny.Map
			_ = _544_v185
			_544_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_419_v94, _425_v100)
			_544_v185 = (_544_v185).Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_308_v3, _414_v89.F21, true, _308_v3, _414_v89.F21), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.IntOfUint32((_dafny.SeqOf(_308_v3, _414_v89.F21, true, _308_v3, _414_v89.F21)).Cardinality()))).Uint32(), _414_v89.F21), _425_v100)
			var _545_v186 _dafny.Set
			_ = _545_v186
			_545_v186 = _dafny.SetOf((_418_v93).F18())
			var _546_v187 _dafny.MultiSet
			_ = _546_v187
			_546_v187 = _dafny.MultiSetOf(_418_v93.F20(), (_545_v186).Cardinality())
			var _pat_let_tv16 = _546_v187
			_ = _pat_let_tv16
			var _547_v188 *C0
			_ = _547_v188
			var _nw93 *C0 = New_C0_()
			_ = _nw93
			_nw93.Ctor__(func(_pat_let10_0 D2) D2 {
				return func(_550_dt__update__tmp_h3 D2) D2 {
					return func(_pat_let13_0 _dafny.MultiSet) D2 {
						return func(_551_dt__update_hcf10_h2 _dafny.MultiSet) D2 {
							return Companion_D2_.Create_DC6_(_551_dt__update_hcf10_h2)
						}(_pat_let13_0)
					}(_dafny.MultiSetOf(_dafny.IntOfInt64(286)))
				}(_pat_let10_0)
			}(func(_pat_let11_0 D2) D2 {
				return func(_548_dt__update__tmp_h2 D2) D2 {
					return func(_pat_let12_0 _dafny.MultiSet) D2 {
						return func(_549_dt__update_hcf10_h1 _dafny.MultiSet) D2 {
							return Companion_D2_.Create_DC6_(_549_dt__update_hcf10_h1)
						}(_pat_let12_0)
					}(_pat_let_tv16)
				}(_pat_let11_0)
			}(Companion_D2_.Create_DC6_(_546_v187))), (_418_v93).F19(), (_418_v93).F18())
			_547_v188 = _nw93
			var _552_v189 _dafny.Set
			_ = _552_v189
			_552_v189 = _dafny.SetOf((Companion_D6_.Create_DC14_(_547_v188)).Dtor_cf20())
			var _553_v190 _dafny.MultiSet
			_ = _553_v190
			_553_v190 = _dafny.MultiSetOf((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), (_418_v93).F19(), (_552_v189).Cardinality(), _dafny.IntOfInt64(360))
			var _554_v191 _dafny.Array
			_ = _554_v191
			var _555_v192 _dafny.Set
			_ = _555_v192
			var _out40 _dafny.Array
			_ = _out40
			var _out41 _dafny.Set
			_ = _out41
			_out40, _out41 = (_414_v89).M2(_546_v187, _418_v93.F20(), _306_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xduc"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-917))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_556_v92 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_557_i12 _dafny.Int) _dafny.CodePoint {
					return _556_v92
				}
			})(_417_v92)))), _309_globalState)
			_554_v191 = _out40
			_555_v192 = _out41
			var _558_v193 _dafny.Map
			_ = _558_v193
			_558_v193 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_418_v93).F18(), _417_v92)
			var _559_v194 _dafny.Map
			_ = _559_v194
			_559_v194 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v89.F21, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qafu"), (Companion_Default___.SafeIndex((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qafu")).Cardinality()))).Uint32(), _417_v92))
			var _560_v195 _dafny.Array
			_ = _560_v195
			var _561_v196 _dafny.Set
			_ = _561_v196
			var _out42 _dafny.Array
			_ = _out42
			var _out43 _dafny.Set
			_ = _out43
			_out42, _out43 = (_414_v89).M2((_dafny.MultiSetOf(_418_v93.F20(), Companion_Default___.Fm0((_418_v93).F19(), (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(_dafny.IntOfInt64(81), _315_v7, (_418_v93).F19(), (_418_v93).F19(), _309_globalState), (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _309_globalState))).Intersection(_553_v190), Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0((_555_v192).Cardinality(), _dafny.IntOfInt64(892), Companion_Default___.Fm0((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), (_418_v93).F19(), (_552_v189).Cardinality(), _418_v93.F20(), _309_globalState), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_418_v93).F19(), true)).Cardinality(), _309_globalState), (_558_v193).Cardinality()), _306_v1, (func() _dafny.Sequence {
				if (_559_v194).Contains((_418_v93).F18()) {
					return (_559_v194).Get((_418_v93).F18()).(_dafny.Sequence)
				}
				return _307_v2
			})(), _309_globalState)
			_560_v195 = _out42
			_561_v196 = _out43
		} else {
			var _562_v197 _dafny.Set
			_ = _562_v197
			_562_v197 = _dafny.SetOf(_417_v92)
			var _pat_let_tv17 = _562_v197
			_ = _pat_let_tv17
			var _pat_let_tv18 = _414_v89
			_ = _pat_let_tv18
			var _pat_let_tv19 = _418_v93
			_ = _pat_let_tv19
			var _pat_let_tv20 = _418_v93
			_ = _pat_let_tv20
			var _pat_let_tv21 = _309_globalState
			_ = _pat_let_tv21
			var _pat_let_tv22 = _418_v93
			_ = _pat_let_tv22
			_427_v102 = func(_pat_let14_0 D0) D0 {
				return func(_563_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let15_0 bool) D0 {
						return func(_564_dt__update_hcf3_h0 bool) D0 {
							return func(_pat_let16_0 _dafny.Int) D0 {
								return func(_565_dt__update_hcf1_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC1_(_565_dt__update_hcf1_h0, (_563_dt__update__tmp_h4).Dtor_cf2(), _564_dt__update_hcf3_h0)
								}(_pat_let16_0)
							}(_pat_let_tv22.F20())
						}(_pat_let15_0)
					}((_pat_let_tv17).IsSubsetOf(Companion_Default___.Fm14(_pat_let_tv18.F21, (_pat_let_tv19).F18(), (_pat_let_tv20).F19(), _pat_let_tv21)))
				}(_pat_let14_0)
			}(_427_v102)
			_426_v101 = (_426_v101).Update(false, (_418_v93).F19())
			var _566_v198 _dafny.MultiSet
			_ = _566_v198
			_566_v198 = _dafny.MultiSetOf((_495_v152).Cardinality(), _dafny.IntOfInt64(122))
			var _567_v200 _dafny.MultiSet
			_ = _567_v200
			_567_v200 = _dafny.MultiSetOf(_307_v2, _307_v2)
			var _568_v201 _dafny.Array
			_ = _568_v201
			var _569_v202 _dafny.Set
			_ = _569_v202
			var _out44 _dafny.Array
			_ = _out44
			var _out45 _dafny.Set
			_ = _out45
			_out44, _out45 = (_414_v89).M2(_566_v198, (func() _dafny.Map {
				var _coll10 = _dafny.NewMapBuilder()
				_ = _coll10
				for _iter10 := _dafny.Iterate((_567_v200).Elements()); ; {
					_compr_10, _ok10 := _iter10()
					if !_ok10 {
						break
					}
					var _570_v199 _dafny.Sequence
					_570_v199 = interface{}(_compr_10).(_dafny.Sequence)
					if (_567_v200).Contains(_570_v199) {
						_coll10.Add(_570_v199, (_418_v93).F18())
					}
				}
				return _coll10.ToMap()
			}()).Cardinality(), _306_v1, _307_v2, _309_globalState)
			_568_v201 = _out44
			_569_v202 = _out45
			var _571_v203 _dafny.Map
			_ = _571_v203
			_571_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v1, _310_v4)
			var _572_v204 D7
			_ = _572_v204
			_572_v204 = Companion_D7_.Create_DC18_(_571_v203)
			_571_v203 = ((_572_v204).Dtor_cf25()).Merge((_571_v203).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_v1, _310_v4)))
			(_414_v89).F21 = ((Companion_Default___.Fm7((_418_v93).F19(), (_418_v93).F19(), (_418_v93).F18(), _309_globalState)).Intersection(_422_v97)).IsSubsetOf(_dafny.MultiSetOf(_308_v3, (_418_v93).F18()))
		}
		var _573_v205 _dafny.Sequence
		_ = _573_v205
		_573_v205 = _dafny.SeqOf(_306_v1, _306_v1, _306_v1)
		var _574_v206 _dafny.Array
		_ = _574_v206
		var _nwElement0_15 _dafny.Array = (_573_v205).Select((Companion_Default___.SafeIndex((_422_v97).Cardinality(), _dafny.IntOfUint32((_573_v205).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _nwElement0_15
		var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(15))
		_ = _nw94
		(_nw94).ArraySet1(_nwElement0_15, 0)
		(_nw94).ArraySet1(_306_v1, 1)
		(_nw94).ArraySet1(_306_v1, 2)
		(_nw94).ArraySet1(_306_v1, 3)
		(_nw94).ArraySet1(_306_v1, 4)
		(_nw94).ArraySet1(_306_v1, 5)
		(_nw94).ArraySet1(_306_v1, 6)
		(_nw94).ArraySet1(_306_v1, 7)
		(_nw94).ArraySet1(_306_v1, 8)
		(_nw94).ArraySet1(_306_v1, 9)
		(_nw94).ArraySet1(_306_v1, 10)
		(_nw94).ArraySet1(_306_v1, 11)
		(_nw94).ArraySet1((_573_v205).Select((Companion_Default___.SafeIndex(_418_v93.F20(), _dafny.IntOfUint32((_573_v205).Cardinality()))).Uint32()).(_dafny.Array), 12)
		(_nw94).ArraySet1(_306_v1, 13)
		(_nw94).ArraySet1(_306_v1, 14)
		_574_v206 = _nw94
		_574_v206 = _574_v206
		var _575_v207 D6
		_ = _575_v207
		_575_v207 = Companion_D6_.Create_DC17_((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _308_v3)
		var _rhs108 _dafny.Int = _418_v93.F20()
		_ = _rhs108
		var _rhs109 bool = (_dafny.IntOfInt64(925)).Cmp(((_418_v93).F19()).Times((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int))) > 0
		_ = _rhs109
		var _rhs110 _dafny.Int = _315_v7
		_ = _rhs110
		var _rhs111 _dafny.Int = (_575_v207).Dtor_cf23()
		_ = _rhs111
		var _lhs83 T1 = _418_v93
		_ = _lhs83
		var _lhs84 *GlobalState = _309_globalState
		_ = _lhs84
		var _lhs85 *GlobalState = _309_globalState
		_ = _lhs85
		var _lhs86 *GlobalState = _309_globalState
		_ = _lhs86
		_lhs83.F20_set_(_rhs108)
		_lhs84.F13 = _rhs109
		_lhs85.F5 = _rhs110
		_lhs86.F9 = _rhs111
		if !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_307_v2, (Companion_Default___.SafeIndex((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_307_v2).Cardinality()))).Uint32(), _417_v92), _dafny.UnicodeSeqOfUtf8Bytes("fa")), _dafny.UnicodeSeqOfUtf8Bytes("a"))) {
			(_309_globalState).F13 = ((Companion_D1_.Create_DC3_(_315_v7)).Dtor_cf5()).Cmp(_418_v93.F20()) != 0
			(_309_globalState).F9 = _418_v93.F20()
			var _576_v208 _dafny.Int
			_ = _576_v208
			var _577_v209 bool
			_ = _577_v209
			var _578_v210 _dafny.Int
			_ = _578_v210
			var _579_v211 _dafny.Int
			_ = _579_v211
			var _out46 _dafny.Int
			_ = _out46
			var _out47 bool
			_ = _out47
			var _out48 _dafny.Int
			_ = _out48
			var _out49 _dafny.Int
			_ = _out49
			_out46, _out47, _out48, _out49 = (_418_v93).M1(_414_v89.F21, (_418_v93).F18(), _309_globalState)
			_576_v208 = _out46
			_577_v209 = _out47
			_578_v210 = _out48
			_579_v211 = _out49
			var _580_v212 D7
			_ = _580_v212
			_580_v212 = Companion_D7_.Create_DC19_(_414_v89.F21, _417_v92, _577_v209)
			var _581_v213 *C1
			_ = _581_v213
			var _nw95 *C1 = New_C1_()
			_ = _nw95
			_nw95.Ctor__(_414_v89.F21, (_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int), (_418_v93).F19(), (_580_v212).Dtor_cf28())
			_581_v213 = _nw95
			var _582_v214 _dafny.MultiSet
			_ = _582_v214
			_582_v214 = _dafny.MultiSetOf((_418_v93).F19())
			var _583_v215 D2
			_ = _583_v215
			_583_v215 = Companion_D2_.Create_DC6_(_582_v214)
			var _584_v216 *C0
			_ = _584_v216
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__(_583_v215, (_418_v93).F19(), _308_v3)
			_584_v216 = _nw96
			var _585_v217 _dafny.Map
			_ = _585_v217
			_585_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_584_v216, _dafny.IntOfInt64(-175))
			var _586_v218 _dafny.Map
			_ = _586_v218
			_586_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(888), _307_v2)
			var _587_v219 _dafny.MultiSet
			_ = _587_v219
			_587_v219 = _dafny.MultiSetOf(_585_v217, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_584_v216, (_586_v218).Cardinality()))
			var _588_v220 *C1
			_ = _588_v220
			var _nw97 *C1 = New_C1_()
			_ = _nw97
			_nw97.Ctor__((_418_v93).F18(), _315_v7, ((Companion_Default___.Fm13(_581_v213.F21, (_418_v93).F18(), (func() _dafny.Int {
				if (_587_v219).Contains(_585_v217) {
					return (_587_v219).Multiplicity(_585_v217)
				}
				return _418_v93.F20()
			})(), (_dafny.Zero).Minus((_584_v216).F23()), _309_globalState)).Cardinality()).Minus((_dafny.MultiSetFromSeq(_425_v100)).Cardinality()), _414_v89.F21)
			_588_v220 = _nw97
		} else {
			(_309_globalState).F9 = (func() _dafny.Int {
				if Companion_Default___.Fm1((_425_v100).Select((Companion_Default___.SafeIndex(_418_v93.F20(), _dafny.IntOfUint32((_425_v100).Cardinality()))).Uint32()).(_dafny.Int), _308_v3, _422_v97, _309_globalState) {
					return (_418_v93).F19()
				}
				return (_418_v93).F19()
			})()
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))
			_ = _index86
			(_310_v4).ArraySet1(((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int)).Plus(_418_v93.F20()), (_index86).Int())
			(_309_globalState).F5 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_418_v93.F20()), func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(800), _dafny.IntOfInt64(926))); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _589_v221 _dafny.Int
					_589_v221 = interface{}(_compr_11).(_dafny.Int)
					if ((_dafny.IntOfInt64(800)).Cmp(_589_v221) <= 0) && ((_589_v221).Cmp(_dafny.IntOfInt64(926)) < 0) {
						_coll11.Add((_589_v221).Times((_418_v93).F19()), _315_v7)
					}
				}
				return _coll11.ToMap()
			}())).Cardinality()
			var _590_v222 _dafny.MultiSet
			_ = _590_v222
			_590_v222 = _dafny.MultiSetOf((_495_v152).Cardinality())
			var _591_v223 D2
			_ = _591_v223
			_591_v223 = Companion_D2_.Create_DC6_(_590_v222)
			var _592_v224 T0
			_ = _592_v224
			var _nw98 *C0 = New_C0_()
			_ = _nw98
			_nw98.Ctor__(_591_v223, _418_v93.F20(), true)
			_592_v224 = _nw98
			var _593_v225 _dafny.Sequence
			_ = _593_v225
			_593_v225 = _dafny.SeqOf(_592_v224, _592_v224)
			var _594_v226 T0
			_ = _594_v226
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__(_591_v223, _dafny.IntOfUint32((_593_v225).Cardinality()), _308_v3)
			_594_v226 = _nw99
			var _595_v227 _dafny.Set
			_ = _595_v227
			_595_v227 = _dafny.SetOf(_594_v226)
			(_309_globalState).F3 = (_595_v227).Cardinality()
			var _596_v228 _dafny.Set
			_ = _596_v228
			_596_v228 = _dafny.SetOf(_307_v2, _dafny.UnicodeSeqOfUtf8Bytes("njxobfr"), _307_v2, _307_v2)
			var _597_v230 _dafny.Map
			_ = _597_v230
			_597_v230 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_596_v228).Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _598_v229 _dafny.Sequence
					_598_v229 = interface{}(_compr_12).(_dafny.Sequence)
					if (_596_v228).Contains(_598_v229) {
						_coll12.Add(_598_v229)
					}
				}
				return _coll12.ToSet()
			}(), _310_v4)
			_597_v230 = (_597_v230).Update(_596_v228, _310_v4)
		}
		var _599_v231 _dafny.MultiSet
		_ = _599_v231
		_599_v231 = _dafny.MultiSetOf(_418_v93.F20(), (_418_v93.F20()).Plus((_418_v93).F19()), _dafny.IntOfUint32((_425_v100).Cardinality()))
		_599_v231 = ((_599_v231).Difference(_599_v231)).Update(_dafny.IntOfUint32((_419_v94).Cardinality()), Companion_Default___.Abs((_310_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_310_v4), 0))).Int()).(_dafny.Int)))
	}
	_dafny.Print((_306_v1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_307_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_308_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState.F2).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_309_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_309_globalState.F4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_309_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F7().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_309_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_309_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_309_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_309_globalState).F17()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_310_v4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_313_v5).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_314_v6).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_315_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_316_v8).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_387_v70).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(604), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_389_v72).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_414_v89.F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_415_v90).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_416_v91).Dtor_cf13().F21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_416_v91).Dtor_cf13()).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_416_v91).Dtor_cf13().F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_416_v91).Dtor_cf13()).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_417_v92)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_418_v93).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_418_v93.F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_418_v93).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_419_v94, _dafny.SeqOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_420_v95).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_421_v96).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_422_v97).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_423_v98).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_424_v99, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_425_v100, _dafny.SeqOf(_dafny.IntOfInt64(-186))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_426_v101).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_427_v102).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_427_v102).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_427_v102).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_428_v103).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_429_v104).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_430_v105).Dtor_cf9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(4), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_484_v149).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_487_v150).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.Zero).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.One).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_494_v151).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(D2)).Dtor_cf10()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_495_v152).Equals(_dafny.SetOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(-525))))
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
	Cf2 _dafny.Int
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
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
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
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
	Cf5 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Int) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf6 bool
	Cf7 bool
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 bool, Cf7 bool, Cf8 bool) D1 {
	return D1{D1_DC4{Cf6, Cf7, Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf4 _dafny.Sequence
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Sequence) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.Zero)
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
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
	Cf10 _dafny.MultiSet
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.MultiSet) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf9 _dafny.Map
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Map) D2 {
	return D2{D2_DC5{Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC7 struct {
	Cf11 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 D2) D2 {
	return D2{D2_DC7{Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(_dafny.EmptyMultiSet)
}

func (_this D2) Dtor_cf10() _dafny.MultiSet {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf9() _dafny.Map {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf11() D2 {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf9.Equals(data2.Cf9)
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
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

type D3_DC8 struct {
	Cf12 _dafny.Map
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 _dafny.Map) D3 {
	return D3{D3_DC8{Cf12}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf14 bool
	Cf15 _dafny.Int
	Cf16 _dafny.Int
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf14 bool, Cf15 _dafny.Int, Cf16 _dafny.Int) D4 {
	return D4{D4_DC10{Cf14, Cf15, Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf13 *C1
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf13 *C1) D4 {
	return D4{D4_DC9{Cf13}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC11 struct {
	Cf17 D4
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 D4) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf14() bool {
	return _this.Get_().(D4_DC10).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf13() *C1 {
	return _this.Get_().(D4_DC9).Cf13
}

func (_this D4) Dtor_cf17() D4 {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf17) + ")"
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
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf13 == data2.Cf13
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17)
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
	Cf19 _dafny.Int
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf19 _dafny.Int) D5 {
	return D5{D5_DC13{Cf19}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC12 struct {
	Cf18 T1
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 T1) D5 {
	return D5{D5_DC12{Cf18}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(_dafny.Zero)
}

func (_this D5) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf18() T1 {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf19.Cmp(data2.Cf19) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && _dafny.AreEqual(data1.Cf18, data2.Cf18)
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
	Cf21 T0
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf21 T0) D6 {
	return D6{D6_DC15{Cf21}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC16 struct {
	Cf22 _dafny.Int
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 _dafny.Int) D6 {
	return D6{D6_DC16{Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC17 struct {
	Cf23 _dafny.Int
	Cf24 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf23 _dafny.Int, Cf24 bool) D6 {
	return D6{D6_DC17{Cf23, Cf24}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC14 struct {
	Cf20 *C0
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf20 *C0) D6 {
	return D6{D6_DC14{Cf20}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC15_((T0)(nil))
}

func (_this D6) Dtor_cf21() T0 {
	return _this.Get_().(D6_DC15).Cf21
}

func (_this D6) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC17).Cf24
}

func (_this D6) Dtor_cf20() *C0 {
	return _this.Get_().(D6_DC14).Cf20
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf21, data2.Cf21)
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf20 == data2.Cf20
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
	Cf26 bool
	Cf27 _dafny.CodePoint
	Cf28 bool
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf26 bool, Cf27 _dafny.CodePoint, Cf28 bool) D7 {
	return D7{D7_DC19{Cf26, Cf27, Cf28}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf29 _dafny.MultiSet
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf29 _dafny.MultiSet) D7 {
	return D7{D7_DC20{Cf29}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC18 struct {
	Cf25 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf25 _dafny.Map) D7 {
	return D7{D7_DC18{Cf25}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC21 struct {
	Cf30 D7
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf30 D7) D7 {
	return D7{D7_DC21{Cf30}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(false, _dafny.CodePoint('D'), false)
}

func (_this D7) Dtor_cf26() bool {
	return _this.Get_().(D7_DC19).Cf26
}

func (_this D7) Dtor_cf27() _dafny.CodePoint {
	return _this.Get_().(D7_DC19).Cf27
}

func (_this D7) Dtor_cf28() bool {
	return _this.Get_().(D7_DC19).Cf28
}

func (_this D7) Dtor_cf29() _dafny.MultiSet {
	return _this.Get_().(D7_DC20).Cf29
}

func (_this D7) Dtor_cf25() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf25
}

func (_this D7) Dtor_cf30() D7 {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf30) + ")"
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
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28 == data2.Cf28
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf30.Equals(data2.Cf30)
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
	Cf31 _dafny.Sequence
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf31 _dafny.Sequence) D8 {
	return D8{D8_DC22{Cf31}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D8_DC22).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf31) + ")"
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
	Cf33 bool
	Cf34 _dafny.Int
	Cf35 bool
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf33 bool, Cf34 _dafny.Int, Cf35 bool) D9 {
	return D9{D9_DC24{Cf33, Cf34, Cf35}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC23 struct {
	Cf32 _dafny.Set
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf32 _dafny.Set) D9 {
	return D9{D9_DC23{Cf32}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC24_(false, _dafny.Zero, false)
}

func (_this D9) Dtor_cf33() bool {
	return _this.Get_().(D9_DC24).Cf33
}

func (_this D9) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D9_DC24).Cf34
}

func (_this D9) Dtor_cf35() bool {
	return _this.Get_().(D9_DC24).Cf35
}

func (_this D9) Dtor_cf32() _dafny.Set {
	return _this.Get_().(D9_DC23).Cf32
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf32) + ")"
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
			return ok && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf32.Equals(data2.Cf32)
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
	Cf37 T1
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf37 T1) D10 {
	return D10{D10_DC26{Cf37}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC27 struct {
	Cf38 _dafny.Int
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf38 _dafny.Int) D10 {
	return D10{D10_DC27{Cf38}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

type D10_DC25 struct {
	Cf36 _dafny.Array
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf36 _dafny.Array) D10 {
	return D10{D10_DC25{Cf36}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

type D10_DC28 struct {
	Cf39 D10
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf39 D10) D10 {
	return D10{D10_DC28{Cf39}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC26_((T1)(nil))
}

func (_this D10) Dtor_cf37() T1 {
	return _this.Get_().(D10_DC26).Cf37
}

func (_this D10) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D10_DC27).Cf38
}

func (_this D10) Dtor_cf36() _dafny.Array {
	return _this.Get_().(D10_DC25).Cf36
}

func (_this D10) Dtor_cf39() D10 {
	return _this.Get_().(D10_DC28).Cf39
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf38) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf39) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf37, data2.Cf37)
		}
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf36 == data2.Cf36
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf39.Equals(data2.Cf39)
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
	Cf41 bool
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf41 bool) D11 {
	return D11{D11_DC30{Cf41}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC31 struct {
	Cf42 _dafny.MultiSet
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf42 _dafny.MultiSet) D11 {
	return D11{D11_DC31{Cf42}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC29 struct {
	Cf40 _dafny.Array
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf40 _dafny.Array) D11 {
	return D11{D11_DC29{Cf40}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC32 struct {
	Cf43 D11
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf43 D11) D11 {
	return D11{D11_DC32{Cf43}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_(false)
}

func (_this D11) Dtor_cf41() bool {
	return _this.Get_().(D11_DC30).Cf41
}

func (_this D11) Dtor_cf42() _dafny.MultiSet {
	return _this.Get_().(D11_DC31).Cf42
}

func (_this D11) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D11_DC29).Cf40
}

func (_this D11) Dtor_cf43() D11 {
	return _this.Get_().(D11_DC32).Cf43
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf43) + ")"
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
			return ok && data1.Cf41 == data2.Cf41
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf40 == data2.Cf40
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D12_DC33 struct {
	Cf44 _dafny.Set
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf44 _dafny.Set) D12 {
	return D12{D12_DC33{Cf44}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

func (CompanionStruct_D12_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D12) Dtor_cf44() _dafny.Set {
	return _this.Get_().(D12_DC33).Cf44
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf44.Equals(data2.Cf44)
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
	Cf46 _dafny.CodePoint
	Cf47 D9
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf46 _dafny.CodePoint, Cf47 D9) D13 {
	return D13{D13_DC35{Cf46, Cf47}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

type D13_DC34 struct {
	Cf45 _dafny.Map
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf45 _dafny.Map) D13 {
	return D13{D13_DC34{Cf45}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

type D13_DC36 struct {
	Cf48 D13
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf48 D13) D13 {
	return D13{D13_DC36{Cf48}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC35_(_dafny.CodePoint('D'), Companion_D9_.Default())
}

func (_this D13) Dtor_cf46() _dafny.CodePoint {
	return _this.Get_().(D13_DC35).Cf46
}

func (_this D13) Dtor_cf47() D9 {
	return _this.Get_().(D13_DC35).Cf47
}

func (_this D13) Dtor_cf45() _dafny.Map {
	return _this.Get_().(D13_DC34).Cf45
}

func (_this D13) Dtor_cf48() D13 {
	return _this.Get_().(D13_DC36).Cf48
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D13_DC36:
		{
			return "D13.DC36" + "(" + _dafny.String(data.Cf48) + ")"
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
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Equals(data2.Cf47)
		}
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

// Definition of trait T0
type T0 interface {
	String() string
	F18() bool
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
	F18() bool
	F20() _dafny.Int
	F20_set_(value _dafny.Int)
	Fm5(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool
	Fm6(p0 _dafny.Map, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map
	M1(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int, _dafny.Int)
	M2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Array, p3 _dafny.Sequence, globalState *GlobalState) (_dafny.Array, _dafny.Set)
	F19() _dafny.Int
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

// Definition of class GlobalState
type GlobalState struct {
	F2   _dafny.Array
	F3   _dafny.Int
	F4   _dafny.Sequence
	F5   _dafny.Int
	F9   _dafny.Int
	F13  bool
	_f0  _dafny.Array
	_f1  bool
	_f6  _dafny.Int
	_f7  _dafny.Sequence
	_f8  _dafny.Int
	_f10 bool
	_f11 _dafny.Int
	_f12 bool
	_f14 _dafny.Int
	_f15 bool
	_f16 bool
	_f17 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F3 = _dafny.Zero
	_this.F4 = _dafny.EmptySeq
	_this.F5 = _dafny.Zero
	_this.F9 = _dafny.Zero
	_this.F13 = false
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = false
	_this._f6 = _dafny.Zero
	_this._f7 = _dafny.EmptySeq
	_this._f8 = _dafny.Zero
	_this._f10 = false
	_this._f11 = _dafny.Zero
	_this._f12 = false
	_this._f14 = _dafny.Zero
	_this._f15 = false
	_this._f16 = false
	_this._f17 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 bool, f2 _dafny.Array, f3 _dafny.Int, f4 _dafny.Sequence, f5 _dafny.Int, f6 _dafny.Int, f7 _dafny.Sequence, f8 _dafny.Int, f9 _dafny.Int, f10 bool, f11 _dafny.Int, f12 bool, f13 bool, f14 _dafny.Int, f15 bool, f16 bool, f17 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
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
func (_this *GlobalState) F8() _dafny.Int {
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
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Array {
	{
		return _this._f17
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f18 bool
	F22  D2
	_f23 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f18 = false
	_this.F22 = Companion_D2_.Default()
	_this._f23 = _dafny.Zero
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

func (_this *C0) F18() bool {
	return _this._f18
}
func (_this *C0) Ctor__(f22 D2, f23 _dafny.Int, f18 bool) {
	{
		(_this).F22 = f22
		(_this)._f23 = f23
		(_this)._f18 = f18
	}
}
func (_this *C0) F23() _dafny.Int {
	{
		return _this._f23
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f20 _dafny.Int
	_f19 _dafny.Int
	_f18 bool
	F21  bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f20 = _dafny.Zero
	_this._f19 = _dafny.Zero
	_this._f18 = false
	_this.F21 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F20() _dafny.Int {
	return _this._f20
}
func (_this *C1) F20_set_(value _dafny.Int) {
	_this._f20 = value
}
func (_this *C1) F19() _dafny.Int {
	return _this._f19
}
func (_this *C1) F18() bool {
	return _this._f18
}
func (_this *C1) Ctor__(f21 bool, f19 _dafny.Int, f20 _dafny.Int, f18 bool) {
	{
		(_this).F21 = f21
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this)._f18 = f18
	}
}
func (_this *C1) Fm5(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_dafny.MultiSetOf(_dafny.IntOfInt64(354), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(654))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_600_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		}))).Cardinality()), _this.F20())).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(444))))).IsDisjointFrom(_dafny.MultiSetOf((_this).F19()))
	}
}
func (_this *C1) Fm6(p0 _dafny.Map, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(989), _this.F20())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F20()), _this.F20()))).Merge((func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus(_this.F20()), (_this).F19())).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _601_v0 _dafny.Int
				_601_v0 = interface{}(_compr_13).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus(_this.F20()), (_this).F19()), _601_v0) {
					_coll13.Add((_601_v0).Minus((_this).F19()), (_this).F19())
				}
			}
			return _coll13.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F20(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F21)).Cardinality()))))
	}
}
func (_this *C1) M1(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _602_v0 _dafny.Array
		_ = _602_v0
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_19
		var _nw100 _dafny.Array
		_ = _nw100
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw100 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) _dafny.Int = func(_603_i0 _dafny.Int) _dafny.Int {
				return (_603_i0).Minus((_this).F19())
			}
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw100 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw100).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw100).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_602_v0 = _nw100
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))
		_ = _index87
		(_602_v0).ArraySet1(Companion_Default___.SafeModuloInt(_this.F20(), _this.F20()), (_index87).Int())
		var _604_v1 _dafny.Map
		_ = _604_v1
		_604_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F20(), _this.F20())
		var _605_v2 _dafny.MultiSet
		_ = _605_v2
		_605_v2 = _dafny.MultiSetOf(_this.F20())
		var _606_v3 D2
		_ = _606_v3
		_606_v3 = Companion_D2_.Create_DC6_(_605_v2)
		var _607_v4 _dafny.MultiSet
		_ = _607_v4
		_607_v4 = _dafny.MultiSetOf(_this.F21)
		var _608_v5 T0
		_ = _608_v5
		var _nw101 *C0 = New_C0_()
		_ = _nw101
		_nw101.Ctor__(_606_v3, (func() _dafny.Int {
			if (_607_v4).Contains(!(p1)) {
				return (_607_v4).Multiplicity(!(p1))
			}
			return _this.F20()
		})(), (_this).F18())
		_608_v5 = _nw101
		var _609_v6 _dafny.Map
		_ = _609_v6
		_609_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_608_v5, _602_v0)
		var _610_v7 _dafny.Sequence
		_ = _610_v7
		_610_v7 = _dafny.UnicodeSeqOfUtf8Bytes("oelkbtu")
		var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))
		_ = _index88
		(_602_v0).ArraySet1(((func() _dafny.Int {
			if (_604_v1).Contains((_this).F19()) {
				return (_604_v1).Get((_this).F19()).(_dafny.Int)
			}
			return ((_609_v6).Update(_608_v5, _602_v0)).Cardinality()
		})()).Minus(_dafny.IntOfUint32((_610_v7).Cardinality())), (_index88).Int())
		var _611_v8 _dafny.Array
		_ = _611_v8
		var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(6))
		_ = _nw102
		_611_v8 = _nw102
		var _ingredients0 = _dafny.NewBuilder()
		_ = _ingredients0
		for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_611_v8), 0))); ; {
			_guard_loop_0, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _612_i1 _dafny.Int
			_612_i1 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_612_i1).Sign() != -1) && ((_612_i1).Cmp(_dafny.ArrayLen((_611_v8), 0)) < 0)) {
				_ingredients0.Add(_dafny.TupleOf(_611_v8, (_612_i1).Int(), (_605_v2).Update(_this.F20(), Companion_Default___.Abs((_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int)))))
			}
		}
		for _iter15 := _dafny.Iterate(_ingredients0); ; {
			_tup0, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
		}
		var _hi1 _dafny.Int = _dafny.IntOfUint32((_610_v7).Cardinality())
		_ = _hi1
		for _613_i2 := _dafny.IntOfInt64(719); _613_i2.Cmp(_hi1) < 0; _613_i2 = _613_i2.Plus(_dafny.One) {
			var _614_v9 _dafny.Set
			_ = _614_v9
			_614_v9 = _dafny.SetOf(p0, p0, true)
			var _615_v10 _dafny.MultiSet
			_ = _615_v10
			_615_v10 = _dafny.MultiSetOf(_614_v9, _614_v9, _614_v9)
			(globalState).F5 = (func() _dafny.Int {
				if (_604_v1).Contains((func() _dafny.Int {
					if (_615_v10).Contains(_614_v9) {
						return (_615_v10).Multiplicity(_614_v9)
					}
					return (_this).F19()
				})()) {
					return (_604_v1).Get((func() _dafny.Int {
						if (_615_v10).Contains(_614_v9) {
							return (_615_v10).Multiplicity(_614_v9)
						}
						return (_this).F19()
					})()).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_604_v1).Contains(_this.F20()) {
						return (_604_v1).Get(_this.F20()).(_dafny.Int)
					}
					return _613_i2
				})()
			})()
			var _616_v11 D1
			_ = _616_v11
			_616_v11 = Companion_D1_.Create_DC3_(_dafny.IntOfInt64(980))
			var _617_v12 _dafny.Map
			_ = _617_v12
			_617_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_608_v5).F18() {
					return (_this).F18()
				}
				return (_this).F18()
			})(), (_dafny.Zero).Minus((_616_v11).Dtor_cf5()))
			var _618_v13 _dafny.Map
			_ = _618_v13
			_618_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_608_v5).F18(), Companion_Default___.Fm1((_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int), (_608_v5).F18(), _607_v4, globalState))
			var _619_v14 _dafny.Map
			_ = _619_v14
			_619_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_608_v5, (_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int))
			var _620_v15 _dafny.Map
			_ = _620_v15
			_620_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v2, _619_v14)
			_617_v12 = (_617_v12).Update((_this).F18(), Companion_Default___.Fm0(_613_i2, _613_i2, (_618_v13).Cardinality(), (_620_v15).Cardinality(), globalState))
			r0 = (_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int)
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))
			_ = _index89
			(_602_v0).ArraySet1((_616_v11).Dtor_cf5(), (_index89).Int())
		}
		var _621_v16 D0
		_ = _621_v16
		_621_v16 = Companion_D0_.Create_DC0_((_this).F18())
		var _622_v17 _dafny.Sequence
		_ = _622_v17
		_622_v17 = _dafny.SeqOf((_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int), _this.F20())
		var _623_v18 _dafny.Array
		_ = _623_v18
		var _nwElement0_16 bool = true
		_ = _nwElement0_16
		var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(11))
		_ = _nw103
		(_nw103).ArraySet1(_nwElement0_16, 0)
		(_nw103).ArraySet1(false, 1)
		(_nw103).ArraySet1(p1, 2)
		(_nw103).ArraySet1(!((_621_v16).Dtor_cf0()), 3)
		(_nw103).ArraySet1(!((_this.F20()).Cmp(_this.F20()) == 0), 4)
		(_nw103).ArraySet1(p1, 5)
		(_nw103).ArraySet1(!(false), 6)
		(_nw103).ArraySet1(p0, 7)
		(_nw103).ArraySet1((_608_v5).F18(), 8)
		(_nw103).ArraySet1(p1, 9)
		(_nw103).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_624_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_625_i3 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_624_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_624_v0), 0))).Int()).(_dafny.Int))
			}
		})(_602_v0))), _622_v17), 10)
		_623_v18 = _nw103
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_623_v18), 0))
		_ = _index90
		(_623_v18).ArraySet1((_this).Fm5(false, _this.F20(), globalState), (_index90).Int())
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_623_v18), 0))
		_ = _index91
		(_623_v18).ArraySet1((func() bool {
			if (_607_v4).IsSubsetOf(_607_v4) {
				return p0
			}
			return p0
		})(), (_index91).Int())
		_607_v4 = (_dafny.MultiSetOf(true, p0)).Union((Companion_Default___.Fm7((_this).F19(), (_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int), p1, globalState)).Union(_607_v4))
		var _626_v19 _dafny.Sequence
		_ = _626_v19
		_626_v19 = _dafny.SeqOf((_608_v5).F18(), (_623_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(650), _dafny.ArrayLen((_623_v18), 0))).Int()).(bool))
		r3 = ((func() _dafny.Int {
			if (_this).F18() {
				return _this.F20()
			}
			return (func() _dafny.Int {
				if (_605_v2).Contains(_dafny.IntOfUint32((_626_v19).Cardinality())) {
					return (_605_v2).Multiplicity(_dafny.IntOfUint32((_626_v19).Cardinality()))
				}
				return _this.F20()
			})()
		})()).Minus((_this).F19())
		r0 = (_this).F19()
		var _627_v20 _dafny.Sequence
		_ = _627_v20
		_627_v20 = _dafny.SeqOf(_608_v5)
		r1 = _dafny.Companion_Sequence_.Contains(_627_v20, _608_v5)
		r2 = (_602_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(449), _dafny.ArrayLen((_602_v0), 0))).Int()).(_dafny.Int)
		r3 = _this.F20()
		return r0, r1, r2, r3
	}
}
func (_this *C1) M2(p0 _dafny.MultiSet, p1 _dafny.Int, p2 _dafny.Array, p3 _dafny.Sequence, globalState *GlobalState) (_dafny.Array, _dafny.Set) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 _dafny.Set = _dafny.EmptySet
		_ = r1
		var _628_v0 _dafny.Array
		_ = _628_v0
		var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
		_ = _nw104
		_628_v0 = _nw104
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_628_v0), 0))); ; {
			_guard_loop_1, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _629_i0 _dafny.Int
			_629_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_629_i0).Sign() != -1) && ((_629_i0).Cmp(_dafny.ArrayLen((_628_v0), 0)) < 0)) {
				(_628_v0).ArraySet1((_629_i0).Minus(_this.F20()), (_629_i0).Int())
			}
		}
		var _630_v1 D2
		_ = _630_v1
		_630_v1 = Companion_D2_.Create_DC6_(p0)
		var _631_v2 _dafny.Sequence
		_ = _631_v2
		_631_v2 = _dafny.SeqOf(p1)
		var _pat_let_tv23 = _630_v1
		_ = _pat_let_tv23
		var _632_v3 _dafny.MultiSet
		_ = _632_v3
		_632_v3 = _dafny.MultiSetOf(_630_v1, _630_v1, func(_pat_let17_0 D2) D2 {
			return func(_633_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let18_0 _dafny.MultiSet) D2 {
					return func(_634_dt__update_hcf10_h0 _dafny.MultiSet) D2 {
						return Companion_D2_.Create_DC6_(_634_dt__update_hcf10_h0)
					}(_pat_let18_0)
				}((_pat_let_tv23).Dtor_cf10())
			}(_pat_let17_0)
		}(Companion_D2_.Create_DC6_(_dafny.MultiSetFromSeq(_631_v2))))
		var _635_i1 _dafny.Int
		_ = _635_i1
		_635_i1 = _dafny.Zero
		{
			for (func() bool {
				if (_this).F18() {
					return _this.F21
				}
				return (_632_v3).IsDisjointFrom(_632_v3)
			})() {
				{
					if (_635_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_635_i1 = (_635_i1).Plus(_dafny.One)
					var _636_v6 _dafny.Array
					_ = _636_v6
					var _len0_20 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_20
					var _nw105 _dafny.Array
					_ = _nw105
					if _len0_20.Cmp(_dafny.Zero) == 0 {
						_nw105 = _dafny.NewArray(_len0_20)
					} else {
						var _init20 func(_dafny.Int) _dafny.Map = func(_637_i2 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(977), _this.F20()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _this.F21))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
								var _coll14 = _dafny.NewBuilder()
								_ = _coll14
								for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-686), _dafny.IntOfInt64(230))); ; {
									_compr_14, _ok17 := _iter17()
									if !_ok17 {
										break
									}
									var _638_v4 _dafny.Int
									_638_v4 = interface{}(_compr_14).(_dafny.Int)
									if ((_dafny.IntOfInt64(-686)).Cmp(_638_v4) <= 0) && ((_638_v4).Cmp(_dafny.IntOfInt64(230)) < 0) {
										_coll14.Add((_638_v4).Minus(_this.F20()))
									}
								}
								return _coll14.ToSet()
							}()).Cardinality(), _this.F20()), func() _dafny.Map {
								var _coll15 = _dafny.NewMapBuilder()
								_ = _coll15
								for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(955), _dafny.IntOfInt64(-330))); ; {
									_compr_15, _ok18 := _iter18()
									if !_ok18 {
										break
									}
									var _639_v5 _dafny.Int
									_639_v5 = interface{}(_compr_15).(_dafny.Int)
									if ((_dafny.IntOfInt64(955)).Cmp(_639_v5) <= 0) && ((_639_v5).Cmp(_dafny.IntOfInt64(-330)) < 0) {
										_coll15.Add((_639_v5).Minus(_dafny.IntOfInt64(32)), false)
									}
								}
								return _coll15.ToMap()
							}()))
						}
						_ = _init20
						var _element0_20 = _init20(_dafny.Zero)
						_ = _element0_20
						_nw105 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
						(_nw105).ArraySet1(_element0_20, 0)
						var _nativeLen0_20 = (_len0_20).Int()
						_ = _nativeLen0_20
						for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
							(_nw105).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
						}
					}
					_636_v6 = _nw105
					var _640_v7 _dafny.CodePoint
					_ = _640_v7
					_640_v7 = _dafny.CodePoint('e')
					var _641_v8 _dafny.Map
					_ = _641_v8
					_641_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_640_v7)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("yh"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yh")).Cardinality()))).Uint32(), _640_v7)).Cardinality()))
					var _642_v9 _dafny.Map
					_ = _642_v9
					_642_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _this.F21)
					var _643_v10 _dafny.Map
					_ = _643_v10
					_643_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_641_v8, _642_v9)
					var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_636_v6), 0))
					_ = _index92
					(_636_v6).ArraySet1(_643_v10, (_index92).Int())
					var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_636_v6), 0))
					_ = _index93
					(_636_v6).ArraySet1((func() _dafny.Map {
						if _this.F21 {
							return (func() _dafny.Map {
								if true {
									return _643_v10
								}
								return _643_v10
							})()
						}
						return func() _dafny.Map {
							var _coll16 = _dafny.NewMapBuilder()
							_ = _coll16
							for _iter19 := _dafny.Iterate((_dafny.SetOf(_641_v8)).Elements()); ; {
								_compr_16, _ok19 := _iter19()
								if !_ok19 {
									break
								}
								var _644_v11 _dafny.Map
								_644_v11 = interface{}(_compr_16).(_dafny.Map)
								if (_dafny.SetOf(_641_v8)).Contains(_644_v11) {
									_coll16.Add(_644_v11, _642_v9)
								}
							}
							return _coll16.ToMap()
						}()
					})(), (_index93).Int())
					if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm3(globalState), _640_v7) {
						(globalState).F5 = p1
						var _645_v12 _dafny.Set
						_ = _645_v12
						_645_v12 = _dafny.SetOf((_this).F19())
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((p2), 0))
						_ = _index94
						(p2).ArraySet1((_645_v12).IsSubsetOf(_645_v12), (_index94).Int())
						var _646_v13 D0
						_ = _646_v13
						_646_v13 = Companion_D0_.Create_DC0_((_this).F18())
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((p2), 0))
						_ = _index95
						(p2).ArraySet1((_646_v13).Dtor_cf0(), (_index95).Int())
						(globalState).F5 = (_dafny.Zero).Minus(_this.F20())
						var _647_v14 _dafny.Set
						_ = _647_v14
						_647_v14 = _dafny.SetOf(_this.F21, _this.F21)
						_647_v14 = (_647_v14).Intersection((_647_v14).Intersection(_647_v14))
						var _648_v15 _dafny.Map
						_ = _648_v15
						_648_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_631_v2, (_this).F19())
						_648_v15 = (_648_v15).Update(_631_v2, (_dafny.IntOfInt64(706)).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((p3).Cardinality()))))
					} else {
						(globalState).F13 = !((_this).F18())
						var _649_v16 _dafny.Set
						_ = _649_v16
						_649_v16 = _dafny.SetOf(_this.F21, _this.F21, (_this).F18())
						var _650_v17 _dafny.Map
						_ = _650_v17
						_650_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(811), _649_v16)
						_650_v17 = _650_v17
						var _651_v18 D1
						_ = _651_v18
						_651_v18 = Companion_D1_.Create_DC4_((_this).F18(), false, _this.F21)
						(globalState).F13 = (func() bool {
							if (_651_v18).Dtor_cf7() {
								return (_this).F18()
							}
							return _this.F21
						})()
						var _652_v19 _dafny.Set
						_ = _652_v19
						_652_v19 = _dafny.SetOf(_this.F20(), (_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), Companion_Default___.Fm0(p1, _dafny.IntOfInt64(-806), p1, p1, globalState))
						var _653_v20 *C0
						_ = _653_v20
						var _nw106 *C0 = New_C0_()
						_ = _nw106
						_nw106.Ctor__(Companion_D2_.Create_DC6_(p0), (_this).F19(), _this.F21)
						_653_v20 = _nw106
						var _654_v21 _dafny.Map
						_ = _654_v21
						_654_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_652_v19, _653_v20)
						_654_v21 = (_654_v21).Update((_652_v19).Difference(_652_v19), _653_v20)
						var _655_v22 _dafny.Sequence
						_ = _655_v22
						_655_v22 = _dafny.SeqOf(p3)
						var _656_v23 _dafny.Map
						_ = _656_v23
						_656_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.Companion_Sequence_.IsProperPrefixOf(_655_v22, _655_v22))
						_656_v23 = _656_v23
					}
					var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_628_v0), 0))
					_ = _index96
					(_628_v0).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F19(), _this.F20()), (_index96).Int())
					var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_628_v0), 0))
					_ = _index97
					(_628_v0).ArraySet1((p1).Times((_this).F19()), (_index97).Int())
					var _657_v24 *C0
					_ = _657_v24
					var _nw107 *C0 = New_C0_()
					_ = _nw107
					_nw107.Ctor__(_630_v1, (_this).F19(), (_this).F18())
					_657_v24 = _nw107
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _658_v25 *C0
		_ = _658_v25
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__(Companion_D2_.Create_DC6_(p0), p1, _this.F21)
		_658_v25 = _nw108
		var _659_v26 _dafny.Set
		_ = _659_v26
		_659_v26 = _dafny.SetOf(_658_v25, _658_v25)
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_628_v0), 0))
		_ = _index98
		(_628_v0).ArraySet1(((_659_v26).Cardinality()).Minus(_dafny.IntOfInt64(825)), (_index98).Int())
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_628_v0), 0))
		_ = _index99
		(_628_v0).ArraySet1(p1, (_index99).Int())
		var _660_v27 _dafny.CodePoint
		_ = _660_v27
		_660_v27 = _dafny.CodePoint('y')
		var _661_v28 _dafny.Set
		_ = _661_v28
		_661_v28 = _dafny.SetOf(_660_v27, _660_v27, _660_v27)
		var _662_v29 _dafny.MultiSet
		_ = _662_v29
		_662_v29 = _dafny.MultiSetOf(_661_v28)
		var _663_v30 _dafny.Map
		_ = _663_v30
		_663_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_662_v29).Update(_661_v28, Companion_Default___.Abs(_dafny.IntOfUint32((_631_v2).Cardinality()))))
		var _664_v31 _dafny.Sequence
		_ = _664_v31
		_664_v31 = _dafny.SeqOf((func() _dafny.MultiSet {
			if (_663_v30).Contains(_this.F21) {
				return (_663_v30).Get(_this.F21).(_dafny.MultiSet)
			}
			return _662_v29
		})(), _dafny.MultiSetOf(_661_v28, _dafny.SetOf(_660_v27)))
		_662_v29 = (_664_v31).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((p3).Cardinality())).Times(_this.F20()), _dafny.IntOfUint32((_664_v31).Cardinality()))).Uint32()).(_dafny.MultiSet)
		var _665_v32 _dafny.Array
		_ = _665_v32
		var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw109
		_665_v32 = _nw109
		for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_665_v32), 0))); ; {
			_guard_loop_2, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _666_i3 _dafny.Int
			_666_i3 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_666_i3).Sign() != -1) && ((_666_i3).Cmp(_dafny.ArrayLen((_665_v32), 0)) < 0)) {
				(_665_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-898))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_667_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_668_i4 _dafny.Int) _dafny.CodePoint {
						return _667_v27
					}
				})(_660_v27)))), (_666_i3).Int())
			}
		}
		_628_v0 = _628_v0
		var _669_v33 _dafny.Array
		_ = _669_v33
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_21
		var _nw110 _dafny.Array
		_ = _nw110
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw110 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) D2 = (func(_670_v25 *C0) func(_dafny.Int) D2 {
				return func(_671_i5 _dafny.Int) D2 {
					return _670_v25.F22
				}
			})(_658_v25)
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw110 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw110).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw110).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_669_v33 = _nw110
		r0 = _669_v33
		r1 = _dafny.SetOf((_this).F19(), (_658_v25).F23())
		return r0, r1
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
