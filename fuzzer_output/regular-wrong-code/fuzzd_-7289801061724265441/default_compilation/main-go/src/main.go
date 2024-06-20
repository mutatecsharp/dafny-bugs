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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Set, globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(_dafny.IntOfInt64(-300))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(905))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-446)
	})))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-787))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fivufydg")).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(390))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i1 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-437)), _dafny.IntOfInt64(967))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pevdpvb")).Cardinality()))
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(523))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i2 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), _dafny.One)
	})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-719), (_dafny.MultiSetOf(true, true, false, false, false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(!(true)))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.MultiSet {
	var _source0 D2 = Companion_D2_.Create_DC8_(_dafny.SeqOf(false, false, false))
	_ = _source0
	if _source0.Is_DC9() {
		var _4___mcc_h0 _dafny.Int = _source0.Get_().(D2_DC9).Cf14
		_ = _4___mcc_h0
		var _5___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC9).Cf15
		_ = _5___mcc_h1
		var _6_cf15 _dafny.Int = _5___mcc_h1
		_ = _6_cf15
		var _7_cf14 _dafny.Int = _4___mcc_h0
		_ = _7_cf14
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("gm"), _dafny.UnicodeSeqOfUtf8Bytes("byjjmnp"), _dafny.UnicodeSeqOfUtf8Bytes("l"), _dafny.UnicodeSeqOfUtf8Bytes("dkpsd"))
	} else {
		var _8___mcc_h2 _dafny.Sequence = _source0.Get_().(D2_DC8).Cf13
		_ = _8___mcc_h2
		var _9_cf13 _dafny.Sequence = _8___mcc_h2
		_ = _9_cf13
		return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ihjrffkga"), _dafny.UnicodeSeqOfUtf8Bytes("xiypy"), _dafny.UnicodeSeqOfUtf8Bytes("kxg"))
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ruoqh"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	})), _dafny.UnicodeSeqOfUtf8Bytes("irmgafus")))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('n'))).Cardinality())).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('c'))).Cardinality())))).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(879))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_11_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_12_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality())
	}))).Cardinality()))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(768), _dafny.IntOfInt64(353))))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) _dafny.Map {
	var r0 _dafny.Map = _dafny.EmptyMap
	_ = r0
	var _13_v0 _dafny.Int
	_ = _13_v0
	_13_v0 = _dafny.IntOfInt64(234)
	var _14_v1 _dafny.Sequence
	_ = _14_v1
	_14_v1 = _dafny.UnicodeSeqOfUtf8Bytes("kmmfs")
	var _15_v2 _dafny.MultiSet
	_ = _15_v2
	_15_v2 = _dafny.MultiSetOf(_dafny.IntOfUint32((_14_v1).Cardinality()))
	var _16_v3 bool
	_ = _16_v3
	_16_v3 = false
	var _17_v4 _dafny.Map
	_ = _17_v4
	_17_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _13_v0)
	var _18_v5 _dafny.MultiSet
	_ = _18_v5
	_18_v5 = _dafny.MultiSetOf(_16_v3)
	var _19_v6 _dafny.Sequence
	_ = _19_v6
	_19_v6 = _dafny.SeqOf(_16_v3, _16_v3)
	var _20_v7 _dafny.Set
	_ = _20_v7
	_20_v7 = _dafny.SetOf(_13_v0, (_dafny.Zero).Minus(_13_v0), _dafny.IntOfUint32((_19_v6).Cardinality()), _dafny.IntOfInt64(417))
	var _21_v8 _dafny.Array
	_ = _21_v8
	var _nwElement0_0 _dafny.Int = _13_v0
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(16))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_13_v0, 1)
	(_nw0).ArraySet1(_13_v0, 2)
	(_nw0).ArraySet1(_dafny.IntOfInt64(43), 3)
	(_nw0).ArraySet1(_13_v0, 4)
	(_nw0).ArraySet1(_13_v0, 5)
	(_nw0).ArraySet1(_13_v0, 6)
	(_nw0).ArraySet1((func() _dafny.Int {
		if (_15_v2).Contains(_13_v0) {
			return (_15_v2).Multiplicity(_13_v0)
		}
		return (func() _dafny.Int {
			if (_17_v4).Contains(Companion_Default___.Fm0(_13_v0, _16_v3, _16_v3, _14_v1, globalState)) {
				return (_17_v4).Get(Companion_Default___.Fm0(_13_v0, _16_v3, _16_v3, _14_v1, globalState)).(_dafny.Int)
			}
			return _13_v0
		})()
	})(), 7)
	(_nw0).ArraySet1(_dafny.IntOfInt64(554), 8)
	(_nw0).ArraySet1(((_18_v5).Cardinality()).Plus(_dafny.IntOfUint32((_14_v1).Cardinality())), 9)
	(_nw0).ArraySet1((func() _dafny.Int {
		if (_18_v5).Contains(_16_v3) {
			return (_18_v5).Multiplicity(_16_v3)
		}
		return _13_v0
	})(), 10)
	(_nw0).ArraySet1(_13_v0, 11)
	(_nw0).ArraySet1(_13_v0, 12)
	(_nw0).ArraySet1(_13_v0, 13)
	(_nw0).ArraySet1((_dafny.IntOfInt64(380)).Times(Companion_Default___.Fm1((_dafny.Zero).Minus(_dafny.IntOfUint32((_19_v6).Cardinality())), _16_v3, _20_v7, globalState)), 14)
	(_nw0).ArraySet1(_13_v0, 15)
	_21_v8 = _nw0
	var _22_v9 _dafny.Map
	_ = _22_v9
	_22_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _21_v8)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))
	_ = _index0
	(_21_v8).ArraySet1((_22_v9).Cardinality(), (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))
	_ = _index1
	(_21_v8).ArraySet1(_13_v0, (_index1).Int())
	var _23_v10 _dafny.Array
	_ = _23_v10
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
	_ = _nw1
	_23_v10 = _nw1
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))
	_ = _index2
	(_23_v10).ArraySet1(_16_v3, (_index2).Int())
	var _24_v11 D3
	_ = _24_v11
	_24_v11 = Companion_D3_.Create_DC10_(_20_v7)
	var _25_v12 _dafny.Sequence
	_ = _25_v12
	_25_v12 = _dafny.SeqOf(Companion_Default___.Fm1(_13_v0, _16_v3, (_24_v11).Dtor_cf16(), globalState), _13_v0)
	var _26_v13 _dafny.Map
	_ = _26_v13
	_26_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(985), _17_v4)
	var _27_v14 _dafny.CodePoint
	_ = _27_v14
	_27_v14 = _dafny.CodePoint('p')
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))
	_ = _index3
	var _rhs0 _dafny.Sequence = _19_v6
	_ = _rhs0
	var _rhs1 _dafny.Int = _13_v0
	_ = _rhs1
	var _rhs2 bool = _16_v3
	_ = _rhs2
	var _rhs3 _dafny.MultiSet = ((_dafny.MultiSetFromSeq(_25_v12)).Union(Companion_Default___.Fm8((_25_v12).Select((Companion_Default___.SafeIndex(((func() _dafny.Map {
		if (_26_v13).Contains((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int)) {
			return (_26_v13).Get((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int)).(_dafny.Map)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _13_v0)
	})()).Cardinality(), _dafny.IntOfUint32((_25_v12).Cardinality()))).Uint32()).(_dafny.Int), _16_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_16_v3, _27_v14), _dafny.IntOfUint32((_19_v6).Cardinality()), globalState))).Intersection(_dafny.MultiSetOf(_13_v0))
	_ = _rhs3
	var _rhs4 bool = _16_v3
	_ = _rhs4
	var _lhs0 *GlobalState = globalState
	_ = _lhs0
	var _lhs1 _dafny.Array = _23_v10
	_ = _lhs1
	var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))
	_ = _lhs2
	_19_v6 = _rhs0
	_lhs0.F20 = _rhs1
	_16_v3 = _rhs2
	_15_v2 = _rhs3
	(_lhs1).ArraySet1(_rhs4, (_lhs2).Int())
	var _hi0 _dafny.Int = (_13_v0).Plus((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int))
	_ = _hi0
	for _28_i0 := _dafny.IntOfUint32((_dafny.SeqOf(_13_v0, _13_v0, (_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int))).Cardinality()); _28_i0.Cmp(_hi0) < 0; _28_i0 = _28_i0.Plus(_dafny.One) {
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))
		_ = _index4
		(_23_v10).ArraySet1((((_dafny.MultiSetOf(_13_v0)).Cardinality()).Cmp(_28_i0) < 0) && (((_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool)) == (_16_v3)), (_index4).Int())
		var _29_v15 _dafny.Map
		_ = _29_v15
		_29_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool), _23_v10)
		_29_v15 = (_29_v15).Update(((_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool)) || (_16_v3), _23_v10)
		var _30_v16 _dafny.Map
		_ = _30_v16
		_30_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v8, (_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int))
		var _31_v17 _dafny.Map
		_ = _31_v17
		_31_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _16_v3)
		var _32_v18 _dafny.Map
		_ = _32_v18
		_32_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_31_v17, _16_v3)
		var _33_v19 _dafny.Sequence
		_ = _33_v19
		_33_v19 = _dafny.SeqOf(_32_v18, _32_v18)
		var _34_v20 _dafny.Array
		_ = _34_v20
		var _nwElement0_1 _dafny.Map = (_30_v16).Merge(_30_v16)
		_ = _nwElement0_1
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(13))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_1, 0)
		(_nw2).ArraySet1(_30_v16, 1)
		(_nw2).ArraySet1(_30_v16, 2)
		(_nw2).ArraySet1(_30_v16, 3)
		(_nw2).ArraySet1(_30_v16, 4)
		(_nw2).ArraySet1(_30_v16, 5)
		(_nw2).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v8, _dafny.IntOfInt64(96)), 6)
		(_nw2).ArraySet1(((_30_v16).Update(_21_v8, (_15_v2).Cardinality())).Update((func() _dafny.Array {
			if (_22_v9).Contains(_16_v3) {
				return (_22_v9).Get(_16_v3).(_dafny.Array)
			}
			return _21_v8
		})(), (_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int)), 7)
		(_nw2).ArraySet1(_30_v16, 8)
		(_nw2).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v8, (_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int)), 9)
		(_nw2).ArraySet1(_30_v16, 10)
		(_nw2).ArraySet1(((_30_v16).Update(_21_v8, ((_33_v19).Select((Companion_Default___.SafeIndex((_25_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_25_v12).Cardinality()), _dafny.IntOfUint32((_25_v12).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_33_v19).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())).Merge(_30_v16), 11)
		(_nw2).ArraySet1(_30_v16, 12)
		_34_v20 = _nw2
		var _35_v21 *C0
		_ = _35_v21
		var _nw3 *C0 = New_C0_()
		_ = _nw3
		_nw3.Ctor__((_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool))
		_35_v21 = _nw3
		var _36_v22 _dafny.MultiSet
		_ = _36_v22
		_36_v22 = _dafny.MultiSetOf(_35_v21)
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_34_v20), 0))
		_ = _index5
		(_34_v20).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v8, (func() _dafny.Int {
			if (_36_v22).Contains(_35_v21) {
				return (_36_v22).Multiplicity(_35_v21)
			}
			return _13_v0
		})())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v8, (_31_v17).Cardinality())), (_index5).Int())
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_34_v20), 0))
		_ = _index6
		var _rhs5 bool = (_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool)
		_ = _rhs5
		var _rhs6 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_14_v1, _14_v1), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("g"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_25_v12).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()))).Uint32(), _dafny.CodePoint('e')))
		_ = _rhs6
		var _rhs7 _dafny.Int = Companion_Default___.Fm1(((_dafny.Zero).Minus((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int))).Plus(_28_i0), !_dafny.Companion_Sequence_.Equal(_19_v6, _dafny.SeqOf(_16_v3)), _20_v7, globalState)
		_ = _rhs7
		var _rhs8 _dafny.Map = _30_v16
		_ = _rhs8
		var _lhs3 *GlobalState = globalState
		_ = _lhs3
		var _lhs4 _dafny.Array = _34_v20
		_ = _lhs4
		var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_34_v20), 0))
		_ = _lhs5
		_16_v3 = _rhs5
		_16_v3 = _rhs6
		_lhs3.F20 = _rhs7
		(_lhs4).ArraySet1(_rhs8, (_lhs5).Int())
		var _37_v23 D0
		_ = _37_v23
		_37_v23 = Companion_D0_.Create_DC0_(_25_v12)
		var _38_v24 _dafny.Array
		_ = _38_v24
		var _nwElement0_2 D0 = Companion_D0_.Create_DC0_(_25_v12)
		_ = _nwElement0_2
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(3))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_2, 0)
		(_nw4).ArraySet1(_37_v23, 1)
		(_nw4).ArraySet1(Companion_Default___.Fm2((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int), globalState), 2)
		_38_v24 = _nw4
		var _39_v25 _dafny.Map
		_ = _39_v25
		_39_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v21, _38_v24)
		_39_v25 = (_39_v25).Update(_35_v21, _38_v24)
	}
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))
	_ = _index7
	(_23_v10).ArraySet1((_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool), (_index7).Int())
	var _40_v26 _dafny.Map
	_ = _40_v26
	_40_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int), (_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int))
	var _41_v27 _dafny.Map
	_ = _41_v27
	_41_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if !(_16_v3) {
			return _13_v0
		}
		return _13_v0
	})(), (func() bool {
		if Companion_Default___.Fm0((_40_v26).Cardinality(), _16_v3, _16_v3, _dafny.UnicodeSeqOfUtf8Bytes("whejvaw"), globalState) {
			return _16_v3
		}
		return true
	})())
	_16_v3 = !((func() bool {
		if (_41_v27).Contains(_13_v0) {
			return (_41_v27).Get(_13_v0).(bool)
		}
		return (_23_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_23_v10), 0))).Int()).(bool)
	})())
	var _42_v28 _dafny.Map
	_ = _42_v28
	_42_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v10, Companion_Default___.Fm0((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_21_v8), 0))).Int()).(_dafny.Int), _16_v3, false, _14_v1, globalState))
	_16_v3 = (func() bool {
		if (_42_v28).Contains(_23_v10) {
			return (_42_v28).Get(_23_v10).(bool)
		}
		return (_16_v3) == (_16_v3)
	})()
	var _43_v29 *C0
	_ = _43_v29
	var _nw5 *C0 = New_C0_()
	_ = _nw5
	_nw5.Ctor__(_16_v3)
	_43_v29 = _nw5
	var _44_v30 _dafny.MultiSet
	_ = _44_v30
	_44_v30 = _dafny.MultiSetOf(_43_v29)
	var _45_v31 _dafny.Map
	_ = _45_v31
	_45_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_44_v30).Cardinality(), _21_v8)
	r0 = _45_v31
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _46_v0 _dafny.Sequence
	_ = _46_v0
	_46_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ljlofb")
	var _47_v1 _dafny.Array
	_ = _47_v1
	var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
	_ = _nw6
	_47_v1 = _nw6
	var _48_v2 _dafny.Int
	_ = _48_v2
	_48_v2 = _dafny.IntOfInt64(135)
	var _49_v3 _dafny.Sequence
	_ = _49_v3
	_49_v3 = _dafny.SeqOf(_48_v2)
	var _50_v4 _dafny.Array
	_ = _50_v4
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
	_ = _nw7
	_50_v4 = _nw7
	var _51_v5 bool
	_ = _51_v5
	_51_v5 = false
	var _52_v6 _dafny.Map
	_ = _52_v6
	_52_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v5, _dafny.IntOfInt64(584))
	var _53_v7 _dafny.CodePoint
	_ = _53_v7
	_53_v7 = _dafny.CodePoint('t')
	var _54_v8 _dafny.Sequence
	_ = _54_v8
	_54_v8 = _dafny.SeqOf(_51_v5)
	var _55_v9 _dafny.Map
	_ = _55_v9
	_55_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_46_v0).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_46_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("apnhjorfd")).Cardinality()))
	var _56_v10 _dafny.Sequence
	_ = _56_v10
	_56_v10 = _dafny.SeqOf(_55_v9, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v7, _48_v2))
	var _57_v11 _dafny.Map
	_ = _57_v11
	_57_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, (_56_v10).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_56_v10).Cardinality()))).Uint32()).(_dafny.Map))
	var _58_v12 _dafny.Sequence
	_ = _58_v12
	_58_v12 = _dafny.SeqOf(_46_v0)
	var _59_globalState *GlobalState
	_ = _59_globalState
	var _nw8 *GlobalState = New_GlobalState_()
	_ = _nw8
	_nw8.Ctor__(_46_v0, _46_v0, false, _47_v1, _49_v3, true, _dafny.IntOfInt64(604), false, _50_v4, (_52_v6).Merge(_52_v6), _dafny.IntOfInt64(503), _dafny.Companion_Sequence_.Update(_46_v0, (Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_46_v0).Cardinality()))).Uint32(), _53_v7), _dafny.Companion_Sequence_.Concatenate(_54_v8, _dafny.Companion_Sequence_.Update(_54_v8, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.IntOfUint32((_54_v8).Cardinality()))).Uint32(), _51_v5)), _dafny.IntOfInt64(996), _49_v3, _50_v4, false, _dafny.IntOfInt64(938), (_57_v11).Merge(_57_v11), _58_v12, _dafny.IntOfInt64(465), _dafny.IntOfInt64(48))
	_59_globalState = _nw8
	var _60_v13 _dafny.Array
	_ = _60_v13
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_0
	var _nw9 _dafny.Array
	_ = _nw9
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw9 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = func(_61_i0 _dafny.Int) bool {
			return false
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw9 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw9).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw9).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_60_v13 = _nw9
	var _62_v14 _dafny.Sequence
	_ = _62_v14
	_62_v14 = _dafny.SeqOf(_60_v13)
	_60_v13 = (_62_v14).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_62_v14).Cardinality()))).Uint32()).(_dafny.Array)
	var _hi1 _dafny.Int = _48_v2
	_ = _hi1
	for _63_i1 := Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_48_v2, _48_v2, _48_v2, _48_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_46_v0).Cardinality()), _48_v2)).Cardinality()), (Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_dafny.SeqOf(_48_v2, _48_v2, _48_v2, _48_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_46_v0).Cardinality()), _48_v2)).Cardinality())).Cardinality()))).Uint32(), _48_v2)).Cardinality()), _48_v2); _63_i1.Cmp(_hi1) < 0; _63_i1 = _63_i1.Plus(_dafny.One) {
		var _64_v15 _dafny.Map
		_ = _64_v15
		_64_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _dafny.IntOfInt64(174))
		var _rhs9 _dafny.Map = _64_v15
		_ = _rhs9
		var _rhs10 _dafny.Int = (_48_v2).Times(_63_i1)
		_ = _rhs10
		var _lhs6 *GlobalState = _59_globalState
		_ = _lhs6
		_64_v15 = _rhs9
		_lhs6.F20 = _rhs10
		var _65_v16 _dafny.Sequence
		_ = _65_v16
		_65_v16 = _dafny.SeqOf(_50_v4)
		var _66_v17 _dafny.Map
		_ = _66_v17
		_66_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, (_65_v16).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_65_v16).Cardinality()))).Uint32()).(_dafny.Array))
		var _67_v18 _dafny.Sequence
		_ = _67_v18
		_67_v18 = _dafny.SeqOf(_50_v4, _50_v4, (func() _dafny.Array {
			if (_66_v17).Contains(_48_v2) {
				return (_66_v17).Get(_48_v2).(_dafny.Array)
			}
			return _50_v4
		})())
		_67_v18 = _dafny.Companion_Sequence_.Update(_65_v16, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vpxkaki")).Cardinality()), _dafny.IntOfUint32((_65_v16).Cardinality()))).Uint32(), _50_v4)
		_54_v8 = _54_v8
		_51_v5 = _51_v5
	}
	_52_v6 = (_52_v6).Update(_51_v5, _48_v2)
	var _68_v19 _dafny.Map
	_ = _68_v19
	var _out0 _dafny.Map
	_ = _out0
	_out0 = Companion_Default___.M0(_59_globalState)
	_68_v19 = _out0
	if (false) || (false) {
		var _69_v20 _dafny.Sequence
		_ = _69_v20
		_69_v20 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_49_v3, _49_v3), _49_v3, _49_v3)
		_69_v20 = _69_v20
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))
		_ = _index8
		(_60_v13).ArraySet1(_51_v5, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))
		_ = _index9
		(_60_v13).ArraySet1(_51_v5, (_index9).Int())
		var _70_v21 _dafny.MultiSet
		_ = _70_v21
		_70_v21 = _dafny.MultiSetOf((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool))
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))
		_ = _index10
		(_60_v13).ArraySet1(!(!(!(_70_v21).Contains(_51_v5))), (_index10).Int())
		_46_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dyxqgdxc"), _46_v0), (Companion_Default___.SafeIndex((func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(938), _dafny.IntOfInt64(402))); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _71_v22 _dafny.Int
				_71_v22 = interface{}(_compr_0).(_dafny.Int)
				if ((_dafny.IntOfInt64(938)).Cmp(_71_v22) <= 0) && ((_71_v22).Cmp(_dafny.IntOfInt64(402)) < 0) {
					_coll0.Add((_71_v22).Times(_48_v2), (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool))
				}
			}
			return _coll0.ToMap()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dyxqgdxc"), _46_v0)).Cardinality()))).Uint32(), _53_v7), _46_v0)
		var _72_v23 _dafny.Array
		_ = _72_v23
		var _nwElement0_3 bool = Companion_Default___.Fm0(_48_v2, (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), _51_v5, _46_v0, _59_globalState)
		_ = _nwElement0_3
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_3, 0)
		(_nw10).ArraySet1(_51_v5, 1)
		(_nw10).ArraySet1(_51_v5, 2)
		(_nw10).ArraySet1(!(_51_v5), 3)
		(_nw10).ArraySet1(_51_v5, 4)
		(_nw10).ArraySet1(_51_v5, 5)
		(_nw10).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfInt64(-251), false, _51_v5, _46_v0, _59_globalState), 6)
		(_nw10).ArraySet1(_51_v5, 7)
		(_nw10).ArraySet1((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), 8)
		(_nw10).ArraySet1(_51_v5, 9)
		(_nw10).ArraySet1((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), 10)
		(_nw10).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfUint32((_46_v0).Cardinality()), (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), _51_v5, _46_v0, _59_globalState), 11)
		(_nw10).ArraySet1(_51_v5, 12)
		(_nw10).ArraySet1((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), 13)
		(_nw10).ArraySet1((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), 14)
		(_nw10).ArraySet1((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), 15)
		(_nw10).ArraySet1(_51_v5, 16)
		_72_v23 = _nw10
		var _73_v24 _dafny.Map
		_ = _73_v24
		_73_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_v23, _48_v2)
		_73_v24 = ((_73_v24).Merge(_73_v24)).Merge(_73_v24)
	} else {
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_1
		var _nw11 _dafny.Array
		_ = _nw11
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw11 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = (func(_74_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_75_i2 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_75_i2, _74_v2)
				}
			})(_48_v2)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw11 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw11).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw11).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_50_v4 = _nw11
		var _76_v25 _dafny.Array
		_ = _76_v25
		var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
		_ = _nw12
		_76_v25 = _nw12
		var _77_v26 _dafny.Map
		_ = _77_v26
		_77_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_49_v3, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_46_v0).Cardinality())), _dafny.IntOfUint32((_49_v3).Cardinality()))).Uint32(), _48_v2), _76_v25)
		var _78_v27 _dafny.Map
		_ = _78_v27
		_78_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_54_v8), _49_v3)
		var _79_v28 _dafny.Sequence
		_ = _79_v28
		_79_v28 = _dafny.SeqOf(_54_v8)
		_77_v26 = (_77_v26).Update((func() _dafny.Sequence {
			if (_78_v27).Contains(_79_v28) {
				return (_78_v27).Get(_79_v28).(_dafny.Sequence)
			}
			return _49_v3
		})(), _76_v25)
		var _80_v29 _dafny.Set
		_ = _80_v29
		_80_v29 = _dafny.SetOf(_60_v13)
		_48_v2 = (_80_v29).Cardinality()
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))
		_ = _index11
		(_60_v13).ArraySet1(_51_v5, (_index11).Int())
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))
		_ = _index12
		(_60_v13).ArraySet1(_51_v5, (_index12).Int())
		var _81_v30 _dafny.Array
		_ = _81_v30
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw13
		_81_v30 = _nw13
		var _82_v31 _dafny.Array
		_ = _82_v31
		var _nwElement0_4 _dafny.Array = _50_v4
		_ = _nwElement0_4
		var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(7))
		_ = _nw14
		(_nw14).ArraySet1(_nwElement0_4, 0)
		(_nw14).ArraySet1(_50_v4, 1)
		(_nw14).ArraySet1(_50_v4, 2)
		(_nw14).ArraySet1(_50_v4, 3)
		(_nw14).ArraySet1(_50_v4, 4)
		(_nw14).ArraySet1(_50_v4, 5)
		(_nw14).ArraySet1(_50_v4, 6)
		_82_v31 = _nw14
		var _83_v32 _dafny.Sequence
		_ = _83_v32
		_83_v32 = _dafny.SeqOf(_82_v31)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_81_v30), 0))
		_ = _index13
		(_81_v30).ArraySet1((_83_v32).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_83_v32).Cardinality()))).Uint32()).(_dafny.Array), (_index13).Int())
		var _84_v33 _dafny.Map
		_ = _84_v33
		_84_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _dafny.IntOfUint32((_54_v8).Cardinality()))
		var _85_v34 _dafny.MultiSet
		_ = _85_v34
		_85_v34 = _dafny.MultiSetOf((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool))
		var _86_v35 _dafny.Map
		_ = _86_v35
		_86_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_85_v34).Contains(Companion_Default___.Fm0(_48_v2, (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), _46_v0, _59_globalState)) {
				return (_85_v34).Multiplicity(Companion_Default___.Fm0(_48_v2, (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool), _46_v0, _59_globalState))
			}
			return _48_v2
		})(), (_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool))
		var _87_v36 _dafny.Set
		_ = _87_v36
		_87_v36 = _dafny.SetOf((_52_v6).Cardinality())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))
		_ = _index14
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_81_v30), 0))
		_ = _index15
		var _rhs11 bool = true
		_ = _rhs11
		var _rhs12 _dafny.Array = _82_v31
		_ = _rhs12
		var _rhs13 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_48_v2, (_86_v35).Cardinality()), _48_v2)
		_ = _rhs13
		var _rhs14 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(633), (_48_v2).Minus(Companion_Default___.Fm1(_48_v2, false, _87_v36, _59_globalState)))
		_ = _rhs14
		var _lhs7 _dafny.Array = _60_v13
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_60_v13), 0))
		_ = _lhs8
		var _lhs9 _dafny.Array = _81_v30
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(588), _dafny.ArrayLen((_81_v30), 0))
		_ = _lhs10
		var _lhs11 *GlobalState = _59_globalState
		_ = _lhs11
		(_lhs7).ArraySet1(_rhs11, (_lhs8).Int())
		(_lhs9).ArraySet1(_rhs12, (_lhs10).Int())
		_84_v33 = _rhs13
		_lhs11.F20 = _rhs14
	}
	var _hi2 _dafny.Int = _48_v2
	_ = _hi2
	for _88_i3 := (_dafny.IntOfInt64(941)).Plus(_48_v2); _88_i3.Cmp(_hi2) < 0; _88_i3 = _88_i3.Plus(_dafny.One) {
		var _89_v37 D0
		_ = _89_v37
		_89_v37 = Companion_D0_.Create_DC0_(_49_v3)
		var _pat_let_tv0 = _49_v3
		_ = _pat_let_tv0
		_49_v3 = (func(_pat_let0_0 D0) D0 {
			return func(_90_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 _dafny.Sequence) D0 {
					return func(_91_dt__update_hcf0_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_91_dt__update_hcf0_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_89_v37)).Dtor_cf0()
		var _92_v38 D0
		_ = _92_v38
		_92_v38 = Companion_D0_.Create_DC3_(_50_v4, _88_i3)
		var _source1 D0 = _92_v38
		_ = _source1
		if _source1.Is_DC1() {
			var _93___mcc_h0 _dafny.Array = _source1.Get_().(D0_DC1).Cf1
			_ = _93___mcc_h0
			var _94___mcc_h1 bool = _source1.Get_().(D0_DC1).Cf2
			_ = _94___mcc_h1
			var _95___mcc_h2 _dafny.Array = _source1.Get_().(D0_DC1).Cf3
			_ = _95___mcc_h2
			var _96_cf3 _dafny.Array = _95___mcc_h2
			_ = _96_cf3
			var _97_cf2 bool = _94___mcc_h1
			_ = _97_cf2
			var _98_cf1 _dafny.Array = _93___mcc_h0
			_ = _98_cf1
			_50_v4 = _50_v4
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_60_v13), 0))
			_ = _index16
			(_60_v13).ArraySet1(!(_97_cf2) || (_97_cf2), (_index16).Int())
			var _99_v39 _dafny.MultiSet
			_ = _99_v39
			_99_v39 = _dafny.MultiSetOf(!(_51_v5), _dafny.Companion_Sequence_.Contains(_49_v3, (_dafny.Zero).Minus(_48_v2)))
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_60_v13), 0))
			_ = _index17
			var _rhs15 bool = !(_51_v5) || (_97_cf2)
			_ = _rhs15
			var _rhs16 _dafny.MultiSet = (_99_v39).Union(_99_v39)
			_ = _rhs16
			var _lhs12 _dafny.Array = _60_v13
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_60_v13), 0))
			_ = _lhs13
			(_lhs12).ArraySet1(_rhs15, (_lhs13).Int())
			_99_v39 = _rhs16
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_60_v13), 0))
			_ = _index18
			(_60_v13).ArraySet1((_51_v5) == ((_88_i3).Cmp(_88_i3) >= 0), (_index18).Int())
			_89_v37 = (func() D0 {
				if _97_cf2 {
					return _89_v37
				}
				return Companion_Default___.Fm2(_dafny.IntOfUint32((_54_v8).Cardinality()), _59_globalState)
			})()
		} else if _source1.Is_DC2() {
			var _100___mcc_h3 bool = _source1.Get_().(D0_DC2).Cf4
			_ = _100___mcc_h3
			var _101_cf4 bool = _100___mcc_h3
			_ = _101_cf4
			var _102_v40 _dafny.Sequence
			_ = _102_v40
			_102_v40 = _dafny.SeqOf(_50_v4, _50_v4)
			var _rhs17 _dafny.Sequence = _dafny.SeqOf(_50_v4, _50_v4, _50_v4, _50_v4)
			_ = _rhs17
			var _rhs18 _dafny.Array = _60_v13
			_ = _rhs18
			var _rhs19 bool = false
			_ = _rhs19
			var _rhs20 _dafny.Int = (_48_v2).Plus((_88_i3).Minus(_48_v2))
			_ = _rhs20
			var _lhs14 *GlobalState = _59_globalState
			_ = _lhs14
			_102_v40 = _rhs17
			_60_v13 = _rhs18
			_51_v5 = _rhs19
			_lhs14.F20 = _rhs20
			_51_v5 = !(_51_v5) || ((func() bool {
				if _51_v5 {
					return !(_51_v5)
				}
				return _51_v5
			})())
			(_59_globalState).F20 = _48_v2
			var _103_v41 _dafny.Set
			_ = _103_v41
			_103_v41 = _dafny.SetOf(_89_v37, _89_v37)
			_103_v41 = (_103_v41).Union(_103_v41)
		} else if _source1.Is_DC3() {
			var _104___mcc_h4 _dafny.Array = _source1.Get_().(D0_DC3).Cf5
			_ = _104___mcc_h4
			var _105___mcc_h5 _dafny.Int = _source1.Get_().(D0_DC3).Cf6
			_ = _105___mcc_h5
			var _106_cf6 _dafny.Int = _105___mcc_h5
			_ = _106_cf6
			var _107_cf5 _dafny.Array = _104___mcc_h4
			_ = _107_cf5
			_51_v5 = (!(true)) || (!(!(_51_v5)) || (_51_v5))
			(_59_globalState).F20 = (_88_i3).Times((_dafny.Zero).Minus((_88_i3).Minus(_dafny.IntOfInt64(538))))
			var _rhs21 _dafny.Int = (_dafny.IntOfUint32((_49_v3).Cardinality())).Plus(_106_cf6)
			_ = _rhs21
			var _rhs22 bool = _51_v5
			_ = _rhs22
			_106_cf6 = _rhs21
			_51_v5 = _rhs22
			var _108_v42 _dafny.Set
			_ = _108_v42
			_108_v42 = _dafny.SetOf(_48_v2)
			(_59_globalState).F20 = (Companion_Default___.Fm1(_88_i3, _51_v5, _108_v42, _59_globalState)).Minus(_88_i3)
		} else if _source1.Is_DC0() {
			var _109___mcc_h6 _dafny.Sequence = _source1.Get_().(D0_DC0).Cf0
			_ = _109___mcc_h6
			var _110_cf0 _dafny.Sequence = _109___mcc_h6
			_ = _110_cf0
			var _111_v43 _dafny.Set
			_ = _111_v43
			_111_v43 = _dafny.SetOf(_48_v2, _48_v2, _88_i3)
			_111_v43 = _111_v43
			var _112_v44 _dafny.Set
			_ = _112_v44
			_112_v44 = _dafny.SetOf(_51_v5, (_54_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.IntOfUint32((_54_v8).Cardinality()))).Uint32()).(bool), !((_54_v8).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_54_v8).Cardinality()))).Uint32()).(bool)), _51_v5)
			_51_v5 = (_112_v44).IsSubsetOf(_112_v44)
			var _113_v45 _dafny.Array
			_ = _113_v45
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw15
			_113_v45 = _nw15
			var _114_v46 _dafny.Map
			_ = _114_v46
			_114_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _dafny.IntOfInt64(303))
			var _115_v47 _dafny.Sequence
			_ = _115_v47
			_115_v47 = _dafny.SeqOf(_114_v46, _114_v46)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_113_v45), 0))
			_ = _index19
			(_113_v45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_115_v47, _115_v47), (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_113_v45), 0))
			_ = _index20
			(_113_v45).ArraySet1(Companion_Default___.Fm3(_59_globalState), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_60_v13), 0))
			_ = _index21
			(_60_v13).ArraySet1(!((func() bool {
				if _51_v5 {
					return _51_v5
				}
				return _51_v5
			})()), (_index21).Int())
			var _116_v48 D0
			_ = _116_v48
			_116_v48 = Companion_D0_.Create_DC2_(_51_v5)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_60_v13), 0))
			_ = _index22
			(_60_v13).ArraySet1((_116_v48).Dtor_cf4(), (_index22).Int())
		} else {
			var _117___mcc_h7 D0 = _source1.Get_().(D0_DC4).Cf7
			_ = _117___mcc_h7
			var _118_cf7 D0 = _117___mcc_h7
			_ = _118_cf7
			var _119_v49 _dafny.Map
			_ = _119_v49
			_119_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_i3, _88_i3)
			var _120_v50 *C0
			_ = _120_v50
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__((_119_v49).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_i3, _48_v2)))
			_120_v50 = _nw16
			var _121_v51 _dafny.MultiSet
			_ = _121_v51
			_121_v51 = _dafny.MultiSetOf((_120_v50).F22())
			_51_v5 = (_121_v51).IsProperSubsetOf((_dafny.MultiSetFromSeq(Companion_Default___.Fm5(_59_globalState))).Update(_51_v5, Companion_Default___.Abs(_48_v2)))
			_68_v19 = _68_v19
			(_59_globalState).F20 = _88_i3
		}
		(_59_globalState).F20 = (func() _dafny.Int {
			if _51_v5 {
				return Companion_Default___.SafeDivisionInt((_52_v6).Cardinality(), _48_v2)
			}
			return _48_v2
		})()
		var _122_v52 _dafny.Set
		_ = _122_v52
		_122_v52 = _dafny.SetOf(_52_v6)
		var _123_v53 D1
		_ = _123_v53
		_123_v53 = Companion_D1_.Create_DC5_(_46_v0)
		var _124_v54 _dafny.Map
		_ = _124_v54
		_124_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v52).IsProperSubsetOf(_122_v52), Companion_D0_.Create_DC3_(_50_v4, _dafny.IntOfUint32(((_123_v53).Dtor_cf8()).Cardinality())))
		var _125_v55 _dafny.MultiSet
		_ = _125_v55
		_125_v55 = _dafny.MultiSetOf((_51_v5) == (_51_v5))
		var _126_v56 _dafny.Array
		_ = _126_v56
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_2
		var _nw17 _dafny.Array
		_ = _nw17
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw17 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Sequence = (func(_127_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_128_i4 _dafny.Int) _dafny.Sequence {
					return _127_v0
				}
			})(_46_v0)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw17 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw17).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw17).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_126_v56 = _nw17
		var _129_v57 D0
		_ = _129_v57
		_129_v57 = Companion_D0_.Create_DC1_(_50_v4, _51_v5, _126_v56)
		var _rhs23 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v5, (func() D0 {
			if false {
				return _92_v38
			}
			return _92_v38
		})())
		_ = _rhs23
		var _rhs24 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_54_v8)).Union(_dafny.MultiSetOf((_129_v57).Dtor_cf2(), _51_v5))
		_ = _rhs24
		_124_v54 = _rhs23
		_125_v55 = _rhs24
	}
	_46_v0 = _46_v0
	var _130_v58 _dafny.Map
	_ = _130_v58
	var _out1 _dafny.Map
	_ = _out1
	_out1 = Companion_Default___.M0(_59_globalState)
	_130_v58 = _out1
	(_59_globalState).F20 = _48_v2
	var _131_v59 D0
	_ = _131_v59
	_131_v59 = Companion_D0_.Create_DC2_(false)
	var _132_i5 _dafny.Int
	_ = _132_i5
	_132_i5 = _dafny.Zero
	{
		for !((_131_v59).Dtor_cf4()) || (true) {
			{
				if (_132_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_132_i5 = (_132_i5).Plus(_dafny.One)
				var _133_v60 _dafny.Array
				_ = _133_v60
				var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
				_ = _nw18
				_133_v60 = _nw18
				var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_133_v60), 0))
				_ = _index23
				(_133_v60).ArraySet1(Companion_Default___.Fm6(_59_globalState), (_index23).Int())
				var _134_v61 _dafny.Map
				_ = _134_v61
				_134_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v5, _46_v0)
				var _135_v62 _dafny.MultiSet
				_ = _135_v62
				_135_v62 = _dafny.MultiSetOf((func() _dafny.Sequence {
					if (_134_v61).Contains(_51_v5) {
						return (_134_v61).Get(_51_v5).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("lrw")
				})())
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_133_v60), 0))
				_ = _index24
				(_133_v60).ArraySet1(((_135_v62).Union(_135_v62)).Update(_46_v0, Companion_Default___.Abs(_48_v2)), (_index24).Int())
				(_59_globalState).F11 = _46_v0
				_51_v5 = _51_v5
				(_59_globalState).F20 = (_dafny.Zero).Minus((_49_v3).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_49_v3).Cardinality()))).Uint32()).(_dafny.Int))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _136_v63 _dafny.Map
	_ = _136_v63
	var _out2 _dafny.Map
	_ = _out2
	_out2 = Companion_Default___.M0(_59_globalState)
	_136_v63 = _out2
	if !(!((_48_v2).Cmp(_dafny.IntOfInt64(27)) != 0)) {
		var _137_v64 _dafny.Map
		_ = _137_v64
		var _out3 _dafny.Map
		_ = _out3
		_out3 = Companion_Default___.M0(_59_globalState)
		_137_v64 = _out3
		if _51_v5 {
			_50_v4 = _50_v4
			_51_v5 = _51_v5
			var _138_v65 D0
			_ = _138_v65
			_138_v65 = Companion_D0_.Create_DC0_(_49_v3)
			var _139_v66 _dafny.Sequence
			_ = _139_v66
			_139_v66 = _dafny.SeqOf(Companion_D0_.Create_DC4_(_138_v65), _138_v65, Companion_D0_.Create_DC4_(_138_v65), _138_v65, _138_v65)
			var _140_v67 _dafny.Map
			_ = _140_v67
			_140_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v7, (_139_v66).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.IntOfUint32((_139_v66).Cardinality()))).Uint32()).(D0))
			var _141_v68 D0
			_ = _141_v68
			_141_v68 = Companion_D0_.Create_DC4_((func() D0 {
				if (_140_v67).Contains(_53_v7) {
					return (_140_v67).Get(_53_v7).(D0)
				}
				return _138_v65
			})())
			var _142_v69 _dafny.Array
			_ = _142_v69
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw19
			_142_v69 = _nw19
			var _143_v70 D0
			_ = _143_v70
			_143_v70 = Companion_D0_.Create_DC4_((func() D0 {
				if _51_v5 {
					return _138_v65
				}
				return Companion_D0_.Create_DC1_(_50_v4, _51_v5, _142_v69)
			})())
			_143_v70 = Companion_D0_.Create_DC4_(_138_v65)
			var _144_v71 _dafny.Set
			_ = _144_v71
			_144_v71 = _dafny.SetOf(_51_v5)
			var _145_v72 _dafny.Sequence
			_ = _145_v72
			_145_v72 = _dafny.SeqOf(_144_v71, _144_v71)
			var _146_v73 _dafny.Array
			_ = _146_v73
			var _nwElement0_5 _dafny.Sequence = _dafny.SeqOf(_48_v2)
			_ = _nwElement0_5
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(9))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_5, 0)
			(_nw20).ArraySet1(_dafny.SeqOf(_48_v2), 1)
			(_nw20).ArraySet1(_49_v3, 2)
			(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_49_v3, _dafny.SeqOf((_144_v71).Cardinality())), 3)
			(_nw20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_49_v3, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_48_v2), (Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_dafny.SeqOf(_48_v2)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(_54_v8, _54_v8, _54_v8, _54_v8)).Cardinality()))), 4)
			(_nw20).ArraySet1(_dafny.SeqOf(_48_v2, _48_v2), 5)
			(_nw20).ArraySet1(_49_v3, 6)
			(_nw20).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(579))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_147_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_148_i6 _dafny.Int) _dafny.Int {
					return _147_v2
				}
			})(_48_v2))), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_48_v2), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(579))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_149_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_150_i6 _dafny.Int) _dafny.Int {
					return _149_v2
				}
			})(_48_v2)))).Cardinality()))).Uint32(), _48_v2), 7)
			(_nw20).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_48_v2), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_48_v2), _dafny.IntOfUint32((_dafny.SeqOf(_48_v2)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_145_v72).Cardinality())), 8)
			_146_v73 = _nw20
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_146_v73), 0))
			_ = _index25
			(_146_v73).ArraySet1(_49_v3, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(705), _dafny.ArrayLen((_146_v73), 0))
			_ = _index26
			(_146_v73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_49_v3, (Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_49_v3).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.SeqOf(_48_v2)).Cardinality())), _49_v3), (_index26).Int())
			var _151_v74 _dafny.Map
			_ = _151_v74
			_151_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _48_v2)
			var _152_v75 _dafny.Set
			_ = _152_v75
			_152_v75 = _dafny.SetOf(_48_v2, _48_v2)
			(_59_globalState).F20 = (_dafny.Zero).Minus(Companion_Default___.Fm1((func() _dafny.Int {
				if (_151_v74).Contains(_48_v2) {
					return (_151_v74).Get(_48_v2).(_dafny.Int)
				}
				return Companion_Default___.Fm1(_48_v2, true, _152_v75, _59_globalState)
			})(), !(_dafny.SetOf(_51_v5)).Equals(_144_v71), _152_v75, _59_globalState))
		} else {
			var _153_v76 *C0
			_ = _153_v76
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__(_51_v5)
			_153_v76 = _nw21
			var _154_v77 _dafny.Map
			_ = _154_v77
			_154_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_49_v3).Cardinality()), _153_v76)
			(_59_globalState).F20 = (((func() _dafny.Map {
				if _51_v5 {
					return _154_v77
				}
				return _154_v77
			})()).Cardinality()).Minus(_dafny.IntOfInt64(61))
			var _nw22 *C0 = New_C0_()
			_ = _nw22
			_nw22.Ctor__((_153_v76).F22())
			_153_v76 = _nw22
			var _155_v78 _dafny.Map
			_ = _155_v78
			var _out4 _dafny.Map
			_ = _out4
			_out4 = Companion_Default___.M0(_59_globalState)
			_155_v78 = _out4
			var _156_v79 _dafny.Map
			_ = _156_v79
			var _out5 _dafny.Map
			_ = _out5
			_out5 = Companion_Default___.M0(_59_globalState)
			_156_v79 = _out5
			var _157_v80 _dafny.Array
			_ = _157_v80
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(5))
			_ = _nw23
			_157_v80 = _nw23
			var _158_v81 _dafny.Set
			_ = _158_v81
			_158_v81 = _dafny.SetOf(_48_v2)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_157_v80), 0))
			_ = _index27
			(_157_v80).ArraySet1((_dafny.SetOf(_48_v2, _dafny.IntOfUint32((_49_v3).Cardinality()))).Difference(_158_v81), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_157_v80), 0))
			_ = _index28
			(_157_v80).ArraySet1((_158_v81).Difference(_158_v81), (_index28).Int())
		}
		_51_v5 = _51_v5
		if _51_v5 {
			var _159_v82 _dafny.Map
			_ = _159_v82
			var _out6 _dafny.Map
			_ = _out6
			_out6 = Companion_Default___.M0(_59_globalState)
			_159_v82 = _out6
			var _160_v83 _dafny.MultiSet
			_ = _160_v83
			_160_v83 = _dafny.MultiSetOf(_50_v4)
			var _161_v84 _dafny.Map
			_ = _161_v84
			_161_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_48_v2), (func() _dafny.Int {
				if (_160_v83).Contains(_50_v4) {
					return (_160_v83).Multiplicity(_50_v4)
				}
				return (_dafny.Zero).Minus(_48_v2)
			})()), _51_v5)
			_161_v84 = (_161_v84).Update(_48_v2, _51_v5)
			_51_v5 = _51_v5
			var _162_v85 *C0
			_ = _162_v85
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__((_54_v8).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_54_v8).Cardinality()))).Uint32()).(bool))
			_162_v85 = _nw24
			var _163_v86 _dafny.Map
			_ = _163_v86
			_163_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v85, _51_v5)
			var _164_v87 _dafny.Set
			_ = _164_v87
			_164_v87 = _dafny.SetOf((_163_v86).Update(_162_v85, _51_v5), _163_v86, _163_v86)
			var _165_v88 _dafny.Set
			_ = _165_v88
			_165_v88 = _dafny.SetOf(_48_v2)
			var _166_v89 _dafny.Map
			_ = _166_v89
			_166_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(Companion_Default___.Fm1((_164_v87).Cardinality(), _51_v5, _165_v88, _59_globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(328))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_167_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_168_i7 _dafny.Int) _dafny.CodePoint {
					return _167_v7
				}
			})(_53_v7)))).Cardinality())), (_48_v2).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gkst")).Cardinality())))
			_166_v89 = _166_v89
			_51_v5 = _dafny.Companion_Sequence_.IsPrefixOf(_46_v0, _dafny.SeqOf(_53_v7))
		} else {
			(_59_globalState).F20 = (_dafny.Zero).Minus(_48_v2)
			_51_v5 = _51_v5
			var _169_v90 _dafny.Map
			_ = _169_v90
			var _out7 _dafny.Map
			_ = _out7
			_out7 = Companion_Default___.M0(_59_globalState)
			_169_v90 = _out7
			var _170_v91 _dafny.Map
			_ = _170_v91
			_170_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_52_v6).Cardinality(), _48_v2), _53_v7)
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_60_v13), 0))
			_ = _index29
			(_60_v13).ArraySet1(_51_v5, (_index29).Int())
			var _171_v92 *C0
			_ = _171_v92
			var _nw25 *C0 = New_C0_()
			_ = _nw25
			_nw25.Ctor__(_51_v5)
			_171_v92 = _nw25
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_60_v13), 0))
			_ = _index30
			(_60_v13).ArraySet1(((_dafny.Zero).Minus(_48_v2)).Cmp(_48_v2) != 0, (_index30).Int())
			var _172_v94 _dafny.MultiSet
			_ = _172_v94
			_172_v94 = _dafny.MultiSetOf(_48_v2)
			var _173_v95 _dafny.Map
			_ = _173_v95
			_173_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v4, _171_v92)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_60_v13), 0))
			_ = _index31
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_60_v13), 0))
			_ = _index32
			var _rhs25 _dafny.Map = func() _dafny.Map {
				var _coll1 = _dafny.NewMapBuilder()
				_ = _coll1
				for _iter1 := _dafny.Iterate((_172_v94).Elements()); ; {
					_compr_1, _ok1 := _iter1()
					if !_ok1 {
						break
					}
					var _174_v93 _dafny.Int
					_174_v93 = interface{}(_compr_1).(_dafny.Int)
					if (_172_v94).Contains(_174_v93) {
						_coll1.Add(Companion_Default___.SafeModuloInt(_174_v93, _48_v2), _53_v7)
					}
				}
				return _coll1.ToMap()
			}()
			_ = _rhs25
			var _rhs26 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7((_171_v92).F22(), _48_v2, _59_globalState), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("oqkka"), (Companion_Default___.SafeIndex((_52_v6).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oqkka")).Cardinality()))).Uint32(), (_46_v0).Select((Companion_Default___.SafeIndex(_48_v2, _dafny.IntOfUint32((_46_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)))
			_ = _rhs26
			var _rhs27 bool = _51_v5
			_ = _rhs27
			var _rhs28 *C0 = _171_v92
			_ = _rhs28
			var _rhs29 bool = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v4, _171_v92)).Equals(_173_v95)
			_ = _rhs29
			var _lhs15 *GlobalState = _59_globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _60_v13
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen((_60_v13), 0))
			_ = _lhs17
			var _lhs18 _dafny.Array = _60_v13
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_60_v13), 0))
			_ = _lhs19
			_170_v91 = _rhs25
			_lhs15.F0 = _rhs26
			(_lhs16).ArraySet1(_rhs27, (_lhs17).Int())
			_171_v92 = _rhs28
			(_lhs18).ArraySet1(_rhs29, (_lhs19).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_50_v4), 0))
			_ = _index33
			(_50_v4).ArraySet1(_48_v2, (_index33).Int())
			var _175_v96 _dafny.Set
			_ = _175_v96
			_175_v96 = _dafny.SetOf(_48_v2)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_50_v4), 0))
			_ = _index34
			(_50_v4).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfInt64(-810), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-288))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_176_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_177_i8 _dafny.Int) _dafny.Int {
					return _176_v2
				}
			})(_48_v2))), _49_v3), _175_v96, _59_globalState), (_index34).Int())
		}
		_48_v2 = _48_v2
	} else {
		var _rhs30 _dafny.CodePoint = _53_v7
		_ = _rhs30
		var _rhs31 bool = _51_v5
		_ = _rhs31
		var _rhs32 _dafny.Int = Companion_Default___.SafeModuloInt(_48_v2, _48_v2)
		_ = _rhs32
		var _lhs20 *GlobalState = _59_globalState
		_ = _lhs20
		_53_v7 = _rhs30
		_51_v5 = _rhs31
		_lhs20.F20 = _rhs32
		_51_v5 = _51_v5
		_48_v2 = _48_v2
		_48_v2 = _48_v2
		if ((func() _dafny.Int {
			if _51_v5 {
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(_48_v2))
			}
			return _48_v2
		})()).Cmp(_dafny.IntOfUint32((_46_v0).Cardinality())) < 0 {
			_53_v7 = _53_v7
			var _178_v97 _dafny.Map
			_ = _178_v97
			_178_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_48_v2, _48_v2), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-178), _48_v2))
			_178_v97 = _178_v97
			var _179_v98 D0
			_ = _179_v98
			_179_v98 = Companion_D0_.Create_DC3_(_50_v4, (_dafny.Zero).Minus((_48_v2).Times(_48_v2)))
			var _180_v99 _dafny.Array
			_ = _180_v99
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw26
			_180_v99 = _nw26
			var _181_v100 _dafny.Set
			_ = _181_v100
			_181_v100 = _dafny.SetOf(_48_v2, _48_v2)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_50_v4), 0))
			_ = _index35
			(_50_v4).ArraySet1((Companion_Default___.Fm1(_48_v2, _51_v5, _181_v100, _59_globalState)).Minus(_48_v2), (_index35).Int())
			var _182_v101 _dafny.Map
			_ = _182_v101
			_182_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v3, _48_v2)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_50_v4), 0))
			_ = _index36
			var _rhs33 D0 = _179_v98
			_ = _rhs33
			var _rhs34 _dafny.Array = (func() _dafny.Array {
				if !(!(_51_v5)) || (_51_v5) {
					return _180_v99
				}
				return _180_v99
			})()
			_ = _rhs34
			var _rhs35 _dafny.Int = Companion_Default___.Fm1(((_182_v101).Cardinality()).Plus(_48_v2), !_dafny.Companion_Sequence_.Equal(_46_v0, _46_v0), _181_v100, _59_globalState)
			_ = _rhs35
			var _rhs36 _dafny.Int = (_48_v2).Times(_48_v2)
			_ = _rhs36
			var _rhs37 _dafny.Int = Companion_Default___.SafeDivisionInt(_48_v2, _48_v2)
			_ = _rhs37
			var _lhs21 _dafny.Array = _50_v4
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_50_v4), 0))
			_ = _lhs22
			var _lhs23 *GlobalState = _59_globalState
			_ = _lhs23
			_179_v98 = _rhs33
			_180_v99 = _rhs34
			(_lhs21).ArraySet1(_rhs35, (_lhs22).Int())
			_lhs23.F20 = _rhs36
			_48_v2 = _rhs37
			var _183_v102 *C0
			_ = _183_v102
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(true)
			_183_v102 = _nw27
			var _184_v103 _dafny.Map
			_ = _184_v103
			_184_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _54_v8)
			var _185_v104 D2
			_ = _185_v104
			_185_v104 = Companion_D2_.Create_DC8_(_dafny.SeqOf(_51_v5, _51_v5))
			var _186_v105 _dafny.Array
			_ = _186_v105
			var _nwElement0_6 _dafny.Sequence = _54_v8
			_ = _nwElement0_6
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(27))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_6, 0)
			(_nw28).ArraySet1(_dafny.SeqOf(_51_v5), 1)
			(_nw28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_54_v8, _54_v8), 2)
			(_nw28).ArraySet1(_54_v8, 3)
			(_nw28).ArraySet1(_54_v8, 4)
			(_nw28).ArraySet1(_54_v8, 5)
			(_nw28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_59_globalState), _54_v8), 6)
			(_nw28).ArraySet1(_54_v8, 7)
			(_nw28).ArraySet1(_54_v8, 8)
			(_nw28).ArraySet1(_dafny.SeqOf(_51_v5), 9)
			(_nw28).ArraySet1(_54_v8, 10)
			(_nw28).ArraySet1(_dafny.SeqOf(true), 11)
			(_nw28).ArraySet1(_54_v8, 12)
			(_nw28).ArraySet1(_dafny.Companion_Sequence_.Update(_54_v8, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fllbype")).Cardinality()), _dafny.IntOfUint32((_54_v8).Cardinality()))).Uint32(), _51_v5), 13)
			(_nw28).ArraySet1(_dafny.SeqOf(_51_v5, (_183_v102).F22(), !(!(!(_51_v5))), _51_v5, _51_v5), 14)
			(_nw28).ArraySet1(_54_v8, 15)
			(_nw28).ArraySet1(_dafny.SeqOf(_51_v5), 16)
			(_nw28).ArraySet1((func() _dafny.Sequence {
				if (_184_v103).Contains((_183_v102).F22()) {
					return (_184_v103).Get((_183_v102).F22()).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_51_v5)
			})(), 17)
			(_nw28).ArraySet1(_dafny.SeqOf(!(_51_v5)), 18)
			(_nw28).ArraySet1((_185_v104).Dtor_cf13(), 19)
			(_nw28).ArraySet1(_54_v8, 20)
			(_nw28).ArraySet1(_54_v8, 21)
			(_nw28).ArraySet1(_54_v8, 22)
			(_nw28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_54_v8, _54_v8), 23)
			(_nw28).ArraySet1(_54_v8, 24)
			(_nw28).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_54_v8, _dafny.SeqOf(_51_v5, false, !(_51_v5), _51_v5)), 25)
			(_nw28).ArraySet1(_54_v8, 26)
			_186_v105 = _nw28
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_186_v105), 0))
			_ = _index37
			(_186_v105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_59_globalState), _54_v8), (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_186_v105), 0))
			_ = _index38
			(_186_v105).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_59_globalState), _dafny.Companion_Sequence_.Concatenate(_54_v8, _54_v8)), (_index38).Int())
		} else {
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_50_v4), 0))
			_ = _index39
			(_50_v4).ArraySet1(_dafny.IntOfUint32((_54_v8).Cardinality()), (_index39).Int())
			var _187_v106 _dafny.Array
			_ = _187_v106
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_3
			var _nw29 _dafny.Array
			_ = _nw29
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw29 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_188_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_189_i9 _dafny.Int) _dafny.Sequence {
						return _188_v0
					}
				})(_46_v0)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw29 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw29).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw29).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_187_v106 = _nw29
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_187_v106), 0))
			_ = _index40
			(_187_v106).ArraySet1(Companion_Default___.Fm7(_51_v5, _48_v2, _59_globalState), (_index40).Int())
			var _190_v107 _dafny.Set
			_ = _190_v107
			_190_v107 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v5, _48_v2))
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_50_v4), 0))
			_ = _index41
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_187_v106), 0))
			_ = _index42
			var _rhs38 _dafny.Sequence = _46_v0
			_ = _rhs38
			var _rhs39 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_190_v107).Cardinality(), _48_v2), _dafny.IntOfUint32((_46_v0).Cardinality()))
			_ = _rhs39
			var _rhs40 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(723))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_191_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_192_i10 _dafny.Int) _dafny.CodePoint {
					return _191_v7
				}
			})(_53_v7)))
			_ = _rhs40
			var _rhs41 _dafny.Array = _50_v4
			_ = _rhs41
			var _rhs42 _dafny.Int = _48_v2
			_ = _rhs42
			var _lhs24 _dafny.Array = _50_v4
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_50_v4), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _187_v106
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_187_v106), 0))
			_ = _lhs27
			var _lhs28 *GlobalState = _59_globalState
			_ = _lhs28
			_46_v0 = _rhs38
			(_lhs24).ArraySet1(_rhs39, (_lhs25).Int())
			(_lhs26).ArraySet1(_rhs40, (_lhs27).Int())
			_50_v4 = _rhs41
			_lhs28.F20 = _rhs42
			var _193_v108 _dafny.Map
			_ = _193_v108
			var _out8 _dafny.Map
			_ = _out8
			_out8 = Companion_Default___.M0(_59_globalState)
			_193_v108 = _out8
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_60_v13), 0))
			_ = _index43
			(_60_v13).ArraySet1(_51_v5, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_60_v13), 0))
			_ = _index44
			(_60_v13).ArraySet1(_51_v5, (_index44).Int())
			var _194_v109 _dafny.MultiSet
			_ = _194_v109
			_194_v109 = _dafny.MultiSetOf(_48_v2)
			_194_v109 = _194_v109
			var _195_v110 _dafny.Set
			_ = _195_v110
			_195_v110 = _dafny.SetOf(_131_v59)
			var _pat_let_tv1 = _60_v13
			_ = _pat_let_tv1
			var _pat_let_tv2 = _60_v13
			_ = _pat_let_tv2
			var _196_v111 _dafny.Map
			_ = _196_v111
			_196_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_195_v110).Contains(func(_pat_let2_0 D0) D0 {
				return func(_197_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let3_0 bool) D0 {
						return func(_198_dt__update_hcf4_h0 bool) D0 {
							return Companion_D0_.Create_DC2_(_198_dt__update_hcf4_h0)
						}(_pat_let3_0)
					}((_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(bool))
				}(_pat_let2_0)
			}(_131_v59)), !(!(!((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool)))) || (!(false)))
			_196_v111 = (_196_v111).Update(false, Companion_Default___.Fm0((_dafny.Zero).Minus(_48_v2), !((_60_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_60_v13), 0))).Int()).(bool)), Companion_Default___.Fm0(_48_v2, _51_v5, _51_v5, _dafny.UnicodeSeqOfUtf8Bytes("t"), _59_globalState), (_187_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_187_v106), 0))).Int()).(_dafny.Sequence), _59_globalState))
		}
	}
	var _199_v112 _dafny.Map
	_ = _199_v112
	var _out9 _dafny.Map
	_ = _out9
	_out9 = Companion_Default___.M0(_59_globalState)
	_199_v112 = _out9
	var _200_v113 _dafny.Map
	_ = _200_v113
	_200_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _49_v3)
	var _201_v114 _dafny.Map
	_ = _201_v114
	_201_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _48_v2)
	var _rhs43 _dafny.Map = (_200_v113).Update(_48_v2, _dafny.SeqOf((func() _dafny.Int {
		if (_201_v114).Contains(_48_v2) {
			return (_201_v114).Get(_48_v2).(_dafny.Int)
		}
		return _48_v2
	})()))
	_ = _rhs43
	var _rhs44 _dafny.Int = _48_v2
	_ = _rhs44
	var _rhs45 bool = !_dafny.Companion_Sequence_.Contains(_58_v12, _dafny.Companion_Sequence_.Concatenate(_46_v0, _46_v0))
	_ = _rhs45
	var _lhs29 *GlobalState = _59_globalState
	_ = _lhs29
	_200_v113 = _rhs43
	_lhs29.F20 = _rhs44
	_51_v5 = _rhs45
	for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_50_v4), 0))); ; {
		_guard_loop_0, _ok2 := _iter2()
		if !_ok2 {
			break
		}
		var _202_i11 _dafny.Int
		_202_i11 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_202_i11).Sign() != -1) && ((_202_i11).Cmp(_dafny.ArrayLen((_50_v4), 0)) < 0)) {
			(_50_v4).ArraySet1((_202_i11).Times(_48_v2), (_202_i11).Int())
		}
	}
	var _203_i12 _dafny.Int
	_ = _203_i12
	_203_i12 = _dafny.Zero
	{
		for _51_v5 {
			{
				if (_203_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_203_i12 = (_203_i12).Plus(_dafny.One)
				var _204_v115 *C0
				_ = _204_v115
				var _nw30 *C0 = New_C0_()
				_ = _nw30
				_nw30.Ctor__(_51_v5)
				_204_v115 = _nw30
				(_59_globalState).F20 = _dafny.IntOfInt64(612)
				(_59_globalState).F20 = _48_v2
				if (_204_v115).F22() {
					var _205_v116 _dafny.Sequence
					_ = _205_v116
					_205_v116 = _dafny.SeqOf(_204_v115, _204_v115)
					var _206_v117 _dafny.Set
					_ = _206_v117
					_206_v117 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_204_v115), _205_v116)).Cardinality()), _dafny.IntOfInt64(73), _48_v2)
					_206_v117 = _206_v117
					(_59_globalState).F20 = (_48_v2).Minus(Companion_Default___.SafeDivisionInt(_48_v2, _48_v2))
					var _207_v118 _dafny.MultiSet
					_ = _207_v118
					_207_v118 = _dafny.MultiSetOf(true)
					var _208_v119 _dafny.Map
					_ = _208_v119
					_208_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _207_v118)
					_208_v119 = (func() _dafny.Map {
						if (_48_v2).Cmp(_dafny.IntOfInt64(50)) >= 0 {
							return _208_v119
						}
						return _208_v119
					})()
					var _209_v120 *C0
					_ = _209_v120
					var _nw31 *C0 = New_C0_()
					_ = _nw31
					_nw31.Ctor__(true)
					_209_v120 = _nw31
					var _210_v121 _dafny.Map
					_ = _210_v121
					var _out10 _dafny.Map
					_ = _out10
					_out10 = Companion_Default___.M0(_59_globalState)
					_210_v121 = _out10
				} else {
					var _211_v122 _dafny.Array
					_ = _211_v122
					var _len0_4 _dafny.Int = _dafny.IntOfInt64(12)
					_ = _len0_4
					var _nw32 _dafny.Array
					_ = _nw32
					if _len0_4.Cmp(_dafny.Zero) == 0 {
						_nw32 = _dafny.NewArray(_len0_4)
					} else {
						var _init4 func(_dafny.Int) _dafny.Sequence = (func(_212_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_213_i13 _dafny.Int) _dafny.Sequence {
								return _212_v0
							}
						})(_46_v0)
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
					_211_v122 = _nw32
					var _214_v123 D0
					_ = _214_v123
					_214_v123 = Companion_D0_.Create_DC1_(_50_v4, !(false), _211_v122)
					var _rhs46 bool = (((_204_v115).F22()) && ((_204_v115).F22())) || ((_214_v123).Dtor_cf2())
					_ = _rhs46
					var _rhs47 _dafny.Int = _48_v2
					_ = _rhs47
					var _rhs48 _dafny.Array = _60_v13
					_ = _rhs48
					var _rhs49 _dafny.Int = _48_v2
					_ = _rhs49
					var _lhs30 *GlobalState = _59_globalState
					_ = _lhs30
					var _lhs31 *GlobalState = _59_globalState
					_ = _lhs31
					_51_v5 = _rhs46
					_lhs30.F20 = _rhs47
					_60_v13 = _rhs48
					_lhs31.F20 = _rhs49
					(_59_globalState).F20 = _48_v2
					var _215_v125 _dafny.Sequence
					_ = _215_v125
					_215_v125 = _dafny.SeqOf((func() _dafny.Map {
						var _coll2 = _dafny.NewMapBuilder()
						_ = _coll2
						for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(266), _dafny.IntOfInt64(17))); ; {
							_compr_2, _ok3 := _iter3()
							if !_ok3 {
								break
							}
							var _216_v124 _dafny.Int
							_216_v124 = interface{}(_compr_2).(_dafny.Int)
							if ((_dafny.IntOfInt64(266)).Cmp(_216_v124) <= 0) && ((_216_v124).Cmp(_dafny.IntOfInt64(17)) < 0) {
								_coll2.Add((_216_v124).Minus(_dafny.IntOfUint32((_46_v0).Cardinality())), _48_v2)
							}
						}
						return _coll2.ToMap()
					}()).Merge(_201_v114), (_201_v114).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, (_dafny.SetOf(_48_v2)).Cardinality())), _201_v114)
					var _217_v126 _dafny.Map
					_ = _217_v126
					_217_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC2_((_204_v115).F22()), _215_v125)
					_215_v125 = (func() _dafny.Sequence {
						if (_217_v126).Contains(Companion_D0_.Create_DC2_((_204_v115).F22())) {
							return (_217_v126).Get(Companion_D0_.Create_DC2_((_204_v115).F22())).(_dafny.Sequence)
						}
						return _215_v125
					})()
					var _218_v127 _dafny.Map
					_ = _218_v127
					_218_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v2, _60_v13)
					_218_v127 = _218_v127
					var _219_v128 D0
					_ = _219_v128
					_219_v128 = Companion_D0_.Create_DC0_(_49_v3)
					var _220_v129 D0
					_ = _220_v129
					_220_v129 = Companion_D0_.Create_DC4_(_219_v128)
					var _pat_let_tv3 = _204_v115
					_ = _pat_let_tv3
					var _pat_let_tv4 = _219_v128
					_ = _pat_let_tv4
					var _221_v130 _dafny.Array
					_ = _221_v130
					var _nwElement0_7 D0 = _220_v129
					_ = _nwElement0_7
					var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(17))
					_ = _nw33
					(_nw33).ArraySet1(_nwElement0_7, 0)
					(_nw33).ArraySet1(_220_v129, 1)
					(_nw33).ArraySet1(_220_v129, 2)
					(_nw33).ArraySet1(Companion_D0_.Create_DC4_(_219_v128), 3)
					(_nw33).ArraySet1(_220_v129, 4)
					(_nw33).ArraySet1(_220_v129, 5)
					(_nw33).ArraySet1(_220_v129, 6)
					(_nw33).ArraySet1(_220_v129, 7)
					(_nw33).ArraySet1(_220_v129, 8)
					(_nw33).ArraySet1(_220_v129, 9)
					(_nw33).ArraySet1(func(_pat_let4_0 D0) D0 {
						return func(_222_dt__update__tmp_h2 D0) D0 {
							return func(_pat_let5_0 D0) D0 {
								return func(_223_dt__update_hcf7_h0 D0) D0 {
									return Companion_D0_.Create_DC4_(_223_dt__update_hcf7_h0)
								}(_pat_let5_0)
							}(Companion_D0_.Create_DC2_((_pat_let_tv3).F22()))
						}(_pat_let4_0)
					}(_220_v129), 10)
					(_nw33).ArraySet1(func(_pat_let6_0 D0) D0 {
						return func(_224_dt__update__tmp_h3 D0) D0 {
							return func(_pat_let7_0 D0) D0 {
								return func(_225_dt__update_hcf7_h1 D0) D0 {
									return Companion_D0_.Create_DC4_(_225_dt__update_hcf7_h1)
								}(_pat_let7_0)
							}(_pat_let_tv4)
						}(_pat_let6_0)
					}(_220_v129), 11)
					(_nw33).ArraySet1(_220_v129, 12)
					(_nw33).ArraySet1(_220_v129, 13)
					(_nw33).ArraySet1(_220_v129, 14)
					(_nw33).ArraySet1(Companion_D0_.Create_DC4_(_219_v128), 15)
					(_nw33).ArraySet1(_220_v129, 16)
					_221_v130 = _nw33
					var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_221_v130), 0))
					_ = _index45
					(_221_v130).ArraySet1(_220_v129, (_index45).Int())
					var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_221_v130), 0))
					_ = _index46
					(_221_v130).ArraySet1((func() D0 {
						if _51_v5 {
							return (func() D0 {
								if (_204_v115).F22() {
									return _220_v129
								}
								return _220_v129
							})()
						}
						return _220_v129
					})(), (_index46).Int())
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_dafny.Print(_46_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_48_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_49_v3, _dafny.SeqOf(_dafny.IntOfInt64(135))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_50_v4).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_51_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_52_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(135))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_53_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_54_v8, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_55_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_56_v10, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(9)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), _dafny.IntOfInt64(135)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(135), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), _dafny.IntOfInt64(135)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_58_v12, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ljlofb"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_59_globalState.F0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_59_globalState).F4(), _dafny.SeqOf(_dafny.IntOfInt64(135))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_59_globalState).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(584))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_59_globalState.F11.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_59_globalState).F12(), _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_59_globalState).F14(), _dafny.SeqOf(_dafny.IntOfInt64(135))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_59_globalState).F18()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(135), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), _dafny.IntOfInt64(135)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_59_globalState.F19, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ljlofb"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_59_globalState.F20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_62_v14).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_68_v19).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v58).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v59).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_132_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v63).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v112).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_200_v113).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.SeqOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_201_v114).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_203_i12)
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
	Cf1 _dafny.Array
	Cf2 bool
	Cf3 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Array, Cf2 bool, Cf3 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 bool) D0 {
	return D0{D0_DC2{Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf5 _dafny.Array
	Cf6 _dafny.Int
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf5 _dafny.Array, Cf6 _dafny.Int) D0 {
	return D0{D0_DC3{Cf5, Cf6}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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

type D0_DC4 struct {
	Cf7 D0
}

func (D0_DC4) isD0() {}

func (CompanionStruct_D0_) Create_DC4_(Cf7 D0) D0 {
	return D0{D0_DC4{Cf7}}
}

func (_this D0) Is_DC4() bool {
	_, ok := _this.Get_().(D0_DC4)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D0_DC3).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf6
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf7() D0 {
	return _this.Get_().(D0_DC4).Cf7
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
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC4:
		{
			return "D0.DC4" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC4:
		{
			data2, ok := other.Get_().(D0_DC4)
			return ok && data1.Cf7.Equals(data2.Cf7)
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

type D1_DC6 struct {
	Cf9  _dafny.Int
	Cf10 _dafny.Int
	Cf11 _dafny.Int
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf9 _dafny.Int, Cf10 _dafny.Int, Cf11 _dafny.Int) D1 {
	return D1{D1_DC6{Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC5 struct {
	Cf8 _dafny.Sequence
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf8 _dafny.Sequence) D1 {
	return D1{D1_DC5{Cf8}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC7 struct {
	Cf12 D1
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_(Cf12 D1) D1 {
	return D1{D1_DC7{Cf12}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC6_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf11
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC5).Cf8
}

func (_this D1) Dtor_cf12() D1 {
	return _this.Get_().(D1_DC7).Cf12
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + data.Cf8.VerbatimString(true) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf8.Equals(data2.Cf8)
		}
	case D1_DC7:
		{
			data2, ok := other.Get_().(D1_DC7)
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

type D2_DC9 struct {
	Cf14 _dafny.Int
	Cf15 _dafny.Int
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf14 _dafny.Int, Cf15 _dafny.Int) D2 {
	return D2{D2_DC9{Cf14, Cf15}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC8 struct {
	Cf13 _dafny.Sequence
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Sequence) D2 {
	return D2{D2_DC8{Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC9_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf15
}

func (_this D2) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf13) + ")"
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
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
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

type D3_DC11 struct {
	Cf17 *C0
	Cf18 *C0
	Cf19 bool
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf17 *C0, Cf18 *C0, Cf19 bool) D3 {
	return D3{D3_DC11{Cf17, Cf18, Cf19}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf16 _dafny.Set
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 _dafny.Set) D3 {
	return D3{D3_DC10{Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_((*C0)(nil), (*C0)(nil), false)
}

func (_this D3) Dtor_cf17() *C0 {
	return _this.Get_().(D3_DC11).Cf17
}

func (_this D3) Dtor_cf18() *C0 {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf16() _dafny.Set {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Sequence
	F11  _dafny.Sequence
	F19  _dafny.Sequence
	F20  _dafny.Int
	_f1  _dafny.Sequence
	_f2  bool
	_f3  _dafny.Array
	_f4  _dafny.Sequence
	_f5  bool
	_f6  _dafny.Int
	_f7  bool
	_f8  _dafny.Array
	_f9  _dafny.Map
	_f10 _dafny.Int
	_f12 _dafny.Sequence
	_f13 _dafny.Int
	_f14 _dafny.Sequence
	_f15 _dafny.Array
	_f16 bool
	_f17 _dafny.Int
	_f18 _dafny.Map
	_f21 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptySeq
	_this.F11 = _dafny.EmptySeq
	_this.F19 = _dafny.EmptySeq
	_this.F20 = _dafny.Zero
	_this._f1 = _dafny.EmptySeq
	_this._f2 = false
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = _dafny.EmptySeq
	_this._f5 = false
	_this._f6 = _dafny.Zero
	_this._f7 = false
	_this._f8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f9 = _dafny.EmptyMap
	_this._f10 = _dafny.Zero
	_this._f12 = _dafny.EmptySeq
	_this._f13 = _dafny.Zero
	_this._f14 = _dafny.EmptySeq
	_this._f15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f16 = false
	_this._f17 = _dafny.Zero
	_this._f18 = _dafny.EmptyMap
	_this._f21 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 _dafny.Sequence, f2 bool, f3 _dafny.Array, f4 _dafny.Sequence, f5 bool, f6 _dafny.Int, f7 bool, f8 _dafny.Array, f9 _dafny.Map, f10 _dafny.Int, f11 _dafny.Sequence, f12 _dafny.Sequence, f13 _dafny.Int, f14 _dafny.Sequence, f15 _dafny.Array, f16 bool, f17 _dafny.Int, f18 _dafny.Map, f19 _dafny.Sequence, f20 _dafny.Int, f21 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this).F20 = f20
		(_this)._f21 = f21
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
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
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
func (_this *GlobalState) F9() _dafny.Map {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
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
func (_this *GlobalState) F14() _dafny.Sequence {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Array {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Map {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F21() _dafny.Int {
	{
		return _this._f21
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f22 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f22 = false
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

func (_this *C0) Ctor__(f22 bool) {
	{
		(_this)._f22 = f22
	}
}
func (_this *C0) Fm4(p0 _dafny.Sequence, p1 bool, p2 _dafny.Map, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F22()
	}
}
func (_this *C0) F22() bool {
	{
		return _this._f22
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
