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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Map, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) bool {
	return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false, false), (_dafny.SeqOf(true)))).Cardinality())).Cmp((_dafny.Zero).Minus(((_dafny.SetOf(true, !(false))).Cardinality()).Plus(_dafny.IntOfInt64(559)))) > 0
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Intersection(_dafny.SetOf(false, false, false, false, true))).Union((_dafny.SetOf(!(false))).Union(_dafny.SetOf(true, !(true))))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uybruf")).Cardinality())), _dafny.IntOfInt64(695), _dafny.IntOfInt64(-92), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("txqeg")).Cardinality()), _dafny.IntOfInt64(559)), _dafny.SeqOf(_dafny.IntOfInt64(657), _dafny.IntOfInt64(-678)))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-387)), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(93), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xuqnfvfjl")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.MultiSet, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(false)), false))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("trfmv")).Cardinality()))), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ataulj")).Cardinality())), true))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Set, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mclcxtfx"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	}))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(161), _dafny.IntOfUint32(((func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("wvcutvx")
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("entn")
	})()).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if !(false) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))
	})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Set, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.CodePoint {
	if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("h"), _dafny.UnicodeSeqOfUtf8Bytes("ehe")) {
		return _dafny.CodePoint('h')
	} else {
		return _dafny.CodePoint('i')
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) (bool, bool) {
	var r0 bool = false
	_ = r0
	var r1 bool = false
	_ = r1
	var _2_v0 bool
	_ = _2_v0
	_2_v0 = false
	var _3_i0 _dafny.Int
	_ = _3_i0
	_3_i0 = _dafny.Zero
	{
		for _2_v0 {
			{
				if (_3_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_3_i0 = (_3_i0).Plus(_dafny.One)
				var _4_v1 *C0
				_ = _4_v1
				var _nw0 *C0 = New_C0_()
				_ = _nw0
				_nw0.Ctor__(p1)
				_4_v1 = _nw0
				var _5_v2 D3
				_ = _5_v2
				_5_v2 = Companion_D3_.Create_DC8_(Companion_D3_.Create_DC6_(_4_v1))
				var _source0 D3 = (func() D3 {
					if true {
						return _5_v2
					}
					return _5_v2
				})()
				_ = _source0
				if _source0.Is_DC7() {
					var _6_v3 _dafny.Set
					_ = _6_v3
					_6_v3 = _dafny.SetOf(_2_v0, _2_v0)
					var _7_v4 _dafny.Set
					_ = _7_v4
					_7_v4 = _dafny.SetOf(_2_v0)
					(globalState).F6 = (_7_v4).IsSubsetOf(_6_v3)
					var _8_v5 *C0
					_ = _8_v5
					var _nw1 *C0 = New_C0_()
					_ = _nw1
					_nw1.Ctor__((_4_v1).F22())
					_8_v5 = _nw1
					var _9_v6 _dafny.Sequence
					_ = _9_v6
					_9_v6 = _dafny.UnicodeSeqOfUtf8Bytes("ack")
					var _10_v7 D0
					_ = _10_v7
					_10_v7 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_9_v6).Cardinality()), _2_v0)
					var _11_v8 _dafny.Sequence
					_ = _11_v8
					_11_v8 = _dafny.SeqOf(((_10_v7).Dtor_cf1()).Cmp(_dafny.IntOfInt64(693)) >= 0, (_6_v3).IsDisjointFrom(_6_v3))
					_11_v8 = _dafny.Companion_Sequence_.Concatenate(_11_v8, _dafny.SeqOf(false, _2_v0))
					(globalState).F6 = !(_2_v0)
				} else if _source0.Is_DC6() {
					var _12___mcc_h0 *C0 = _source0.Get_().(D3_DC6).Cf8
					_ = _12___mcc_h0
					var _13_cf8 *C0 = _12___mcc_h0
					_ = _13_cf8
					var _14_v9 _dafny.Sequence
					_ = _14_v9
					_14_v9 = _dafny.UnicodeSeqOfUtf8Bytes("jyy")
					_14_v9 = _14_v9
					(globalState).F6 = !(_2_v0)
					var _15_v10 _dafny.Map
					_ = _15_v10
					_15_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-222), _2_v0)
					var _16_v11 _dafny.Map
					_ = _16_v11
					_16_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("y"), (_13_cf8).Fm1((func() bool {
						if (_15_v10).Contains(p2) {
							return (_15_v10).Get(p2).(bool)
						}
						return _2_v0
					})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(984))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg2 _dafny.Int) interface{} {
							return coer2(arg2)
						}
					}(func(_17_i1 _dafny.Int) _dafny.Sequence {
						return (_dafny.UnicodeSeqOfUtf8Bytes("kqfg"))
					})), globalState))
					var _18_v12 _dafny.CodePoint
					_ = _18_v12
					_18_v12 = _dafny.CodePoint('m')
					_16_v11 = (_16_v11).Update(_14_v9, !_dafny.Companion_Sequence_.Contains(_14_v9, _18_v12))
					var _19_v13 *C0
					_ = _19_v13
					var _nw2 *C0 = New_C0_()
					_ = _nw2
					_nw2.Ctor__(((_4_v1).F22()).Update(_2_v0, p0))
					_19_v13 = _nw2
				} else {
					var _20___mcc_h1 D3 = _source0.Get_().(D3_DC8).Cf9
					_ = _20___mcc_h1
					var _21_cf9 D3 = _20___mcc_h1
					_ = _21_cf9
					var _22_v14 _dafny.Array
					_ = _22_v14
					var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
					_ = _nw3
					_22_v14 = _nw3
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_22_v14), 0))
					_ = _index0
					(_22_v14).ArraySet1(p0, (_index0).Int())
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_22_v14), 0))
					_ = _index1
					(_22_v14).ArraySet1(p0, (_index1).Int())
					var _23_v15 _dafny.Sequence
					_ = _23_v15
					_23_v15 = _dafny.UnicodeSeqOfUtf8Bytes("u")
					var _24_v16 _dafny.Sequence
					_ = _24_v16
					_24_v16 = _dafny.SeqOf(_4_v1, _4_v1, _4_v1, _4_v1, _4_v1)
					var _25_v17 _dafny.Map
					_ = _25_v17
					_25_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_23_v15).Cardinality()))
					var _26_v18 D3
					_ = _26_v18
					_26_v18 = Companion_D3_.Create_DC7_()
					var _27_v19 _dafny.Map
					_ = _27_v19
					_27_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2_v0, _26_v18)
					var _28_v20 _dafny.Set
					_ = _28_v20
					_28_v20 = _dafny.SetOf(_2_v0, _2_v0)
					var _29_v21 _dafny.CodePoint
					_ = _29_v21
					_29_v21 = _dafny.CodePoint('t')
					var _30_v22 _dafny.Array
					_ = _30_v22
					var _nwElement0_0 _dafny.Int = _dafny.IntOfInt64(740)
					_ = _nwElement0_0
					var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(26))
					_ = _nw4
					(_nw4).ArraySet1(_nwElement0_0, 0)
					(_nw4).ArraySet1(p2, 1)
					(_nw4).ArraySet1(p2, 2)
					(_nw4).ArraySet1(p2, 3)
					(_nw4).ArraySet1((_dafny.IntOfInt64(992)).Minus(p2), 4)
					(_nw4).ArraySet1(Companion_Default___.SafeDivisionInt(p2, p2), 5)
					(_nw4).ArraySet1(p2, 6)
					(_nw4).ArraySet1(Companion_Default___.SafeDivisionInt(p2, p2), 7)
					(_nw4).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-561)), 8)
					(_nw4).ArraySet1(p2, 9)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_24_v16).Cardinality()), 10)
					(_nw4).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), 11)
					(_nw4).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), 12)
					(_nw4).ArraySet1(p2, 13)
					(_nw4).ArraySet1((_dafny.Zero).Minus((_25_v17).Cardinality()), 14)
					(_nw4).ArraySet1((p2).Minus(p2), 15)
					(_nw4).ArraySet1(((_27_v19).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2_v0, Companion_D3_.Create_DC7_()))).Cardinality(), 16)
					(_nw4).ArraySet1(p2, 17)
					(_nw4).ArraySet1((_dafny.Zero).Minus(p2), 18)
					(_nw4).ArraySet1(p2, 19)
					(_nw4).ArraySet1(p2, 20)
					(_nw4).ArraySet1((_dafny.Zero).Minus((p2).Times(p2)), 21)
					(_nw4).ArraySet1((p2).Times((_28_v20).Cardinality()), 22)
					(_nw4).ArraySet1(p2, 23)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eybkhsaop"), _dafny.UnicodeSeqOfUtf8Bytes("ddyvcpun")), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eybkhsaop"), _dafny.UnicodeSeqOfUtf8Bytes("ddyvcpun"))).Cardinality()))).Uint32(), _29_v21)).Cardinality()), 24)
					(_nw4).ArraySet1(_dafny.IntOfInt64(-209), 25)
					_30_v22 = _nw4
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_30_v22), 0))
					_ = _index2
					(_30_v22).ArraySet1(p2, (_index2).Int())
					var _31_v23 _dafny.MultiSet
					_ = _31_v23
					_31_v23 = _dafny.MultiSetOf(_4_v1)
					var _32_v24 D3
					_ = _32_v24
					_32_v24 = Companion_D3_.Create_DC6_(_4_v1)
					var _33_v25 _dafny.Map
					_ = _33_v25
					_33_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v24, _2_v0)
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_30_v22), 0))
					_ = _index3
					var _rhs0 bool = ((_31_v23).Intersection(_31_v23)).Equals(_dafny.MultiSetOf(_4_v1))
					_ = _rhs0
					var _rhs1 _dafny.Sequence = _23_v15
					_ = _rhs1
					var _rhs2 _dafny.Int = (func() _dafny.Int {
						if _2_v0 {
							return (_33_v25).Cardinality()
						}
						return (p2).Times(p2)
					})()
					_ = _rhs2
					var _rhs3 _dafny.Int = _dafny.IntOfInt64(650)
					_ = _rhs3
					var _lhs0 _dafny.Array = _30_v22
					_ = _lhs0
					var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_30_v22), 0))
					_ = _lhs1
					var _lhs2 *GlobalState = globalState
					_ = _lhs2
					r0 = _rhs0
					_23_v15 = _rhs1
					(_lhs0).ArraySet1(_rhs2, (_lhs1).Int())
					_lhs2.F7 = _rhs3
					var _34_v26 _dafny.Sequence
					_ = _34_v26
					_34_v26 = _dafny.SeqOf(_23_v15)
					var _35_v27 _dafny.MultiSet
					_ = _35_v27
					_35_v27 = _dafny.MultiSetOf(_2_v0, !(_2_v0), _2_v0, !(_2_v0), (func() bool {
						if _2_v0 {
							return (_4_v1).Fm1(_2_v0, _34_v26, globalState)
						}
						return _2_v0
					})())
					var _rhs4 _dafny.MultiSet = _35_v27
					_ = _rhs4
					var _rhs5 _dafny.Int = (_30_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(783), _dafny.ArrayLen((_30_v22), 0))).Int()).(_dafny.Int)
					_ = _rhs5
					var _lhs3 *GlobalState = globalState
					_ = _lhs3
					_35_v27 = _rhs4
					_lhs3.F20 = _rhs5
					var _36_v28 _dafny.Sequence
					_ = _36_v28
					_36_v28 = _dafny.SeqOf(false, _2_v0, _2_v0, _2_v0)
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((p0), 0))
					_ = _index4
					(p0).ArraySet1(!_dafny.Companion_Sequence_.Contains(_36_v28, _2_v0), (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((p0), 0))
					_ = _index5
					(p0).ArraySet1(_2_v0, (_index5).Int())
				}
				var _37_v29 _dafny.Array
				_ = _37_v29
				var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw5
				_37_v29 = _nw5
				var _38_v30 _dafny.MultiSet
				_ = _38_v30
				_38_v30 = _dafny.MultiSetOf(_2_v0)
				var _39_v31 _dafny.Sequence
				_ = _39_v31
				_39_v31 = _dafny.SeqOf(_4_v1)
				var _40_v32 _dafny.Sequence
				_ = _40_v32
				_40_v32 = _dafny.UnicodeSeqOfUtf8Bytes("nb")
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_37_v29), 0))
				_ = _index6
				(_37_v29).ArraySet1(((func() _dafny.Int {
					if (_38_v30).Contains(_2_v0) {
						return (_38_v30).Multiplicity(_2_v0)
					}
					return _dafny.IntOfUint32((_39_v31).Cardinality())
				})()).Times(Companion_Default___.Fm8(p2, _dafny.IntOfUint32((_40_v32).Cardinality()), _2_v0, _2_v0, globalState)), (_index6).Int())
				var _41_v33 _dafny.Map
				_ = _41_v33
				_41_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_4_v1, _2_v0)
				var _42_v34 _dafny.CodePoint
				_ = _42_v34
				_42_v34 = _dafny.CodePoint('p')
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_37_v29), 0))
				_ = _index7
				var _rhs6 bool = _2_v0
				_ = _rhs6
				var _rhs7 bool = (!(!(_2_v0)) || (_2_v0)) == (_2_v0)
				_ = _rhs7
				var _rhs8 _dafny.Int = (((Companion_Default___.Fm9(p2, p2, _2_v0, globalState)).Cardinality()).Times(p2)).Plus(p2)
				_ = _rhs8
				var _rhs9 _dafny.Int = (func() _dafny.Int {
					if !_dafny.Companion_Sequence_.Equal(_40_v32, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("necjelxnj"), (Companion_Default___.SafeIndex((_41_v33).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("necjelxnj")).Cardinality()))).Uint32(), _42_v34)) {
						return (_dafny.Zero).Minus(p2)
					}
					return p2
				})()
				_ = _rhs9
				var _rhs10 _dafny.Int = (p2).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality())).Times(p2))
				_ = _rhs10
				var _lhs4 *GlobalState = globalState
				_ = _lhs4
				var _lhs5 *GlobalState = globalState
				_ = _lhs5
				var _lhs6 *GlobalState = globalState
				_ = _lhs6
				var _lhs7 _dafny.Array = _37_v29
				_ = _lhs7
				var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_37_v29), 0))
				_ = _lhs8
				var _lhs9 *GlobalState = globalState
				_ = _lhs9
				_lhs4.F6 = _rhs6
				_lhs5.F4 = _rhs7
				_lhs6.F14 = _rhs8
				(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
				_lhs9.F7 = _rhs10
				_2_v0 = _2_v0
				var _43_v35 _dafny.Sequence
				_ = _43_v35
				_43_v35 = _dafny.SeqOf(_2_v0, _2_v0)
				var _44_v36 _dafny.Set
				_ = _44_v36
				_44_v36 = _dafny.SetOf(_2_v0, !((_43_v35).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_37_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_37_v29), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_43_v35).Cardinality()))).Uint32()).(bool)))
				var _45_v37 _dafny.Map
				_ = _45_v37
				_45_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2_v0, (_44_v36).Equals(_44_v36))
				var _46_v38 _dafny.Sequence
				_ = _46_v38
				_46_v38 = _dafny.SeqOf(Companion_Default___.Fm8(_dafny.IntOfUint32((_40_v32).Cardinality()), p2, _2_v0, _2_v0, globalState))
				var _47_v39 D1
				_ = _47_v39
				_47_v39 = Companion_D1_.Create_DC2_(p1)
				var _48_v40 D1
				_ = _48_v40
				_48_v40 = Companion_D1_.Create_DC4_(_47_v39)
				var _49_v41 _dafny.Sequence
				_ = _49_v41
				_49_v41 = _dafny.SeqOf(_48_v40)
				var _50_v42 _dafny.Set
				_ = _50_v42
				_50_v42 = _dafny.SetOf(_dafny.IntOfUint32((_49_v41).Cardinality()))
				var _51_v43 _dafny.Map
				_ = _51_v43
				_51_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_2_v0), _dafny.CodePoint('n'))
				var _52_v44 _dafny.MultiSet
				_ = _52_v44
				_52_v44 = _dafny.MultiSetOf(_51_v43, _51_v43)
				var _rhs11 *C0 = _4_v1
				_ = _rhs11
				var _rhs12 _dafny.Map = (_45_v37).Merge((_45_v37).Update(_2_v0, _2_v0))
				_ = _rhs12
				var _rhs13 _dafny.Int = (_37_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_37_v29), 0))).Int()).(_dafny.Int)
				_ = _rhs13
				var _rhs14 _dafny.CodePoint = Companion_Default___.Fm10(_50_v42, (_52_v44).Difference(_52_v44), globalState)
				_ = _rhs14
				var _rhs15 _dafny.Sequence = _46_v38
				_ = _rhs15
				var _lhs10 *GlobalState = globalState
				_ = _lhs10
				_4_v1 = _rhs11
				_45_v37 = _rhs12
				_lhs10.F14 = _rhs13
				_42_v34 = _rhs14
				_46_v38 = _rhs15
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	r1 = _2_v0
	(globalState).F20 = p2
	var _53_v45 _dafny.Sequence
	_ = _53_v45
	_53_v45 = _dafny.UnicodeSeqOfUtf8Bytes("eqgwm")
	var _54_v46 _dafny.Sequence
	_ = _54_v46
	_54_v46 = _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_53_v45, _53_v45), (p2).Cmp((_dafny.Zero).Minus(p2)) < 0, (_2_v0) || (_2_v0), _2_v0, _2_v0)
	var _55_v48 _dafny.Sequence
	_ = _55_v48
	_55_v48 = _dafny.SeqOf((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(209), _dafny.IntOfInt64(925))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _56_v47 _dafny.Int
			_56_v47 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(209)).Cmp(_56_v47) <= 0) && ((_56_v47).Cmp(_dafny.IntOfInt64(925)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_56_v47, p2), _2_v0)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())
	r0 = (_54_v46).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm8(_dafny.IntOfUint32((_55_v48).Cardinality()), p2, _2_v0, false, globalState), _dafny.IntOfUint32((_54_v46).Cardinality()))).Uint32()).(bool)
	var _57_v49 _dafny.Array
	_ = _57_v49
	var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
	_ = _nw6
	_57_v49 = _nw6
	_57_v49 = p0
	var _58_v50 _dafny.Map
	_ = _58_v50
	_58_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_57_v49, (_55_v48).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_55_v48).Cardinality()))).Uint32()).(_dafny.Int))
	_58_v50 = (_58_v50).Merge(_58_v50)
	var _59_v51 D0
	_ = _59_v51
	_59_v51 = Companion_D0_.Create_DC1_(p2, _2_v0)
	var _pat_let_tv0 = _2_v0
	_ = _pat_let_tv0
	r0 = func(_source1 D0) bool {
		if _source1.Is_DC1() {
			var _60___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC1).Cf1
			_ = _60___mcc_h2
			var _61___mcc_h3 bool = _source1.Get_().(D0_DC1).Cf2
			_ = _61___mcc_h3
			var _62_cf2 bool = _61___mcc_h3
			_ = _62_cf2
			var _63_cf1 _dafny.Int = _60___mcc_h2
			_ = _63_cf1
			return _pat_let_tv0
		} else {
			var _64___mcc_h4 _dafny.Map = _source1.Get_().(D0_DC0).Cf0
			_ = _64___mcc_h4
			var _65_cf0 _dafny.Map = _64___mcc_h4
			_ = _65_cf0
			return false
		}
	}(_59_v51)
	var _66_v52 _dafny.MultiSet
	_ = _66_v52
	_66_v52 = _dafny.MultiSetOf(p2, p2)
	r1 = (_66_v52).IsDisjointFrom(_66_v52)
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _67_v0 _dafny.CodePoint
	_ = _67_v0
	_67_v0 = _dafny.CodePoint('w')
	var _68_v1 _dafny.Int
	_ = _68_v1
	_68_v1 = _dafny.IntOfInt64(557)
	var _69_v2 _dafny.Array
	_ = _69_v2
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
	_ = _nw7
	_69_v2 = _nw7
	var _70_v3 _dafny.Map
	_ = _70_v3
	_70_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v2, _68_v1)
	var _71_v4 _dafny.Set
	_ = _71_v4
	_71_v4 = _dafny.SetOf((_70_v3).Cardinality(), _68_v1)
	var _72_v5 _dafny.MultiSet
	_ = _72_v5
	_72_v5 = _dafny.MultiSetOf(_71_v4)
	var _73_v6 bool
	_ = _73_v6
	_73_v6 = true
	var _74_v7 _dafny.MultiSet
	_ = _74_v7
	_74_v7 = _dafny.MultiSetOf(_73_v6)
	var _75_v8 _dafny.Map
	_ = _75_v8
	_75_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v6, _68_v1)
	var _76_v9 _dafny.Array
	_ = _76_v9
	var _nwElement0_1 _dafny.Int = _68_v1
	_ = _nwElement0_1
	var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(22))
	_ = _nw8
	(_nw8).ArraySet1(_nwElement0_1, 0)
	(_nw8).ArraySet1((_72_v5).Cardinality(), 1)
	(_nw8).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-906))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}((func(_77_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_78_i1 _dafny.Int) _dafny.CodePoint {
			return _77_v0
		}
	})(_67_v0)))).Cardinality()), 2)
	(_nw8).ArraySet1(_68_v1, 3)
	(_nw8).ArraySet1(_68_v1, 4)
	(_nw8).ArraySet1(_68_v1, 5)
	(_nw8).ArraySet1(_68_v1, 6)
	(_nw8).ArraySet1(_68_v1, 7)
	(_nw8).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _73_v6)).Cardinality(), 8)
	(_nw8).ArraySet1(_68_v1, 9)
	(_nw8).ArraySet1(_68_v1, 10)
	(_nw8).ArraySet1(_68_v1, 11)
	(_nw8).ArraySet1(_68_v1, 12)
	(_nw8).ArraySet1(_68_v1, 13)
	(_nw8).ArraySet1(_68_v1, 14)
	(_nw8).ArraySet1(_68_v1, 15)
	(_nw8).ArraySet1(_68_v1, 16)
	(_nw8).ArraySet1((func() _dafny.Int {
		if (_74_v7).Contains(_73_v6) {
			return (_74_v7).Multiplicity(_73_v6)
		}
		return _68_v1
	})(), 17)
	(_nw8).ArraySet1((_75_v8).Cardinality(), 18)
	(_nw8).ArraySet1(_68_v1, 19)
	(_nw8).ArraySet1(_68_v1, 20)
	(_nw8).ArraySet1(_dafny.IntOfInt64(875), 21)
	_76_v9 = _nw8
	var _79_v10 _dafny.Map
	_ = _79_v10
	_79_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v6, !(_73_v6))
	var _80_v11 _dafny.Set
	_ = _80_v11
	_80_v11 = _dafny.SetOf(_75_v8, _75_v8, _75_v8)
	var _81_v12 _dafny.Array
	_ = _81_v12
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
	_ = _nw9
	_81_v12 = _nw9
	var _82_v13 _dafny.Set
	_ = _82_v13
	_82_v13 = _dafny.SetOf(_73_v6)
	var _83_v14 _dafny.Sequence
	_ = _83_v14
	_83_v14 = _dafny.SeqOf(_82_v13)
	var _84_v15 _dafny.Sequence
	_ = _84_v15
	_84_v15 = _dafny.UnicodeSeqOfUtf8Bytes("je")
	var _85_globalState *GlobalState
	_ = _85_globalState
	var _nw10 *GlobalState = New_GlobalState_()
	_ = _nw10
	_nw10.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-417))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_86_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), _67_v0), _76_v9, _79_v10, _80_v11, true, _81_v12, true, _dafny.IntOfInt64(929), _74_v7, _dafny.IntOfInt64(319), _dafny.IntOfInt64(-434), true, (_71_v4).Intersection(_71_v4), _76_v9, _dafny.IntOfInt64(-685), true, ((_83_v14).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_83_v14).Cardinality()))).Uint32()).(_dafny.Set)).Union(_dafny.SetOf(_73_v6)), _69_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_84_v15).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _68_v1)), _dafny.IntOfInt64(581), _dafny.IntOfInt64(37), _84_v15)
	_85_globalState = _nw10
	if true {
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_69_v2), 0))
		_ = _index8
		(_69_v2).ArraySet1((_dafny.Zero).Minus(_68_v1), (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_69_v2), 0))
		_ = _index9
		(_69_v2).ArraySet1(_68_v1, (_index9).Int())
		_75_v8 = (_75_v8).Update((_74_v7).IsSubsetOf(_74_v7), ((_69_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_69_v2), 0))).Int()).(_dafny.Int)).Times((_69_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_69_v2), 0))).Int()).(_dafny.Int)))
		var _87_v16 _dafny.Array
		_ = _87_v16
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
		_ = _nw11
		_87_v16 = _nw11
		_87_v16 = _87_v16
		_84_v15 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(971))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_88_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_89_i2 _dafny.Int) _dafny.CodePoint {
				return _88_v0
			}
		})(_67_v0)))
		_73_v6 = !(true) || (_73_v6)
	} else {
		var _90_v17 _dafny.Array
		_ = _90_v17
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_0
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_91_v15 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_92_i3 _dafny.Int) bool {
					return _dafny.Companion_Sequence_.Equal(_91_v15, _91_v15)
				}
			})(_84_v15)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw12 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw12).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw12).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_90_v17 = _nw12
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_90_v17), 0))
		_ = _index10
		(_90_v17).ArraySet1(_73_v6, (_index10).Int())
		var _93_v18 _dafny.MultiSet
		_ = _93_v18
		_93_v18 = _dafny.MultiSetOf(_dafny.IntOfInt64(805), _68_v1)
		var _94_v19 _dafny.Sequence
		_ = _94_v19
		_94_v19 = _dafny.SeqOf(_68_v1)
		var _95_v20 _dafny.MultiSet
		_ = _95_v20
		_95_v20 = _dafny.MultiSetOf(((func() _dafny.Int {
			if (_93_v18).Contains(_68_v1) {
				return (_93_v18).Multiplicity(_68_v1)
			}
			return _68_v1
		})()).Minus(_68_v1), _68_v1, (_dafny.IntOfUint32((_94_v19).Cardinality())).Plus(_68_v1), _68_v1)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_69_v2), 0))
		_ = _index11
		(_69_v2).ArraySet1(Companion_Default___.SafeDivisionInt(_68_v1, _68_v1), (_index11).Int())
		var _96_v21 _dafny.Map
		_ = _96_v21
		_96_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_84_v15).Cardinality()), _68_v1), _dafny.MultiSetOf(_68_v1))
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_90_v17), 0))
		_ = _index12
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_69_v2), 0))
		_ = _index13
		var _rhs16 bool = _73_v6
		_ = _rhs16
		var _rhs17 _dafny.MultiSet = (func() _dafny.MultiSet {
			if (_96_v21).Contains(_94_v19) {
				return (_96_v21).Get(_94_v19).(_dafny.MultiSet)
			}
			return _95_v20
		})()
		_ = _rhs17
		var _rhs18 _dafny.Int = (_dafny.IntOfInt64(-971)).Minus(_68_v1)
		_ = _rhs18
		var _rhs19 _dafny.Int = _68_v1
		_ = _rhs19
		var _lhs11 _dafny.Array = _90_v17
		_ = _lhs11
		var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_90_v17), 0))
		_ = _lhs12
		var _lhs13 _dafny.Array = _69_v2
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_69_v2), 0))
		_ = _lhs14
		var _lhs15 *GlobalState = _85_globalState
		_ = _lhs15
		(_lhs11).ArraySet1(_rhs16, (_lhs12).Int())
		_95_v20 = _rhs17
		(_lhs13).ArraySet1(_rhs18, (_lhs14).Int())
		_lhs15.F20 = _rhs19
		(_85_globalState).F6 = Companion_Default___.Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, (_69_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_69_v2), 0))).Int()).(_dafny.Int)), _68_v1, (_90_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_90_v17), 0))).Int()).(bool), _73_v6, _85_globalState)
		_76_v9 = _76_v9
		var _97_v22 _dafny.Map
		_ = _97_v22
		_97_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_69_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_69_v2), 0))).Int()).(_dafny.Int), _68_v1)
		var _98_v23 _dafny.Map
		_ = _98_v23
		_98_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_97_v22).Cardinality(), (_90_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_90_v17), 0))).Int()).(bool))
		var _99_v24 _dafny.Map
		_ = _99_v24
		_99_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v23, _dafny.IntOfInt64(446))
		_73_v6 = !(_99_v24).Equals(_99_v24)
		_82_v13 = _82_v13
	}
	var _100_v25 _dafny.Array
	_ = _100_v25
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
	_ = _nw13
	_100_v25 = _nw13
	var _101_v26 _dafny.Array
	_ = _101_v26
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
	_ = _nw14
	_101_v26 = _nw14
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_100_v25), 0))
	_ = _index14
	(_100_v25).ArraySet1(_101_v26, (_index14).Int())
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_100_v25), 0))
	_ = _index15
	(_100_v25).ArraySet1(_101_v26, (_index15).Int())
	(_85_globalState).F4 = _73_v6
	var _102_v27 _dafny.Map
	_ = _102_v27
	_102_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v6, _dafny.UnicodeSeqOfUtf8Bytes("vfucpd"))
	_84_v15 = (func() _dafny.Sequence {
		if (_102_v27).Contains(_73_v6) {
			return (_102_v27).Get(_73_v6).(_dafny.Sequence)
		}
		return _84_v15
	})()
	(_85_globalState).F4 = (_68_v1).Cmp(_68_v1) < 0
	if !(_73_v6) || (!(_73_v6)) {
		(_85_globalState).F14 = _dafny.IntOfInt64(-584)
		var _103_v28 _dafny.Sequence
		_ = _103_v28
		_103_v28 = _dafny.SeqOf(_73_v6, _73_v6, _73_v6)
		var _104_v29 _dafny.Map
		_ = _104_v29
		_104_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, (_71_v4).Cardinality())
		var _105_v30 _dafny.Array
		_ = _105_v30
		var _nwElement0_2 bool = _73_v6
		_ = _nwElement0_2
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(11))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_2, 0)
		(_nw15).ArraySet1(_73_v6, 1)
		(_nw15).ArraySet1(_73_v6, 2)
		(_nw15).ArraySet1(_73_v6, 3)
		(_nw15).ArraySet1((_103_v28).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_103_v28).Cardinality()))).Uint32()).(bool), 4)
		(_nw15).ArraySet1(_73_v6, 5)
		(_nw15).ArraySet1(_73_v6, 6)
		(_nw15).ArraySet1(_73_v6, 7)
		(_nw15).ArraySet1(Companion_Default___.Fm0(_104_v29, _68_v1, _73_v6, _73_v6, _85_globalState), 8)
		(_nw15).ArraySet1(_73_v6, 9)
		(_nw15).ArraySet1(!(!(_73_v6)), 10)
		_105_v30 = _nw15
		var _106_v31 _dafny.Map
		_ = _106_v31
		_106_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v6, _105_v30)
		var _107_v32 _dafny.Map
		_ = _107_v32
		_107_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _74_v7)
		var _108_v33 _dafny.Map
		_ = _108_v33
		_108_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(Companion_Default___.Fm0(_104_v29, _68_v1, _73_v6, _73_v6, _85_globalState), false), (func() _dafny.MultiSet {
			if (_107_v32).Contains(true) {
				return (_107_v32).Get(true).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(_73_v6)
		})())
		var _109_v34 bool
		_ = _109_v34
		var _110_v35 bool
		_ = _110_v35
		var _out0 bool
		_ = _out0
		var _out1 bool
		_ = _out1
		_out0, _out1 = Companion_Default___.M0(_105_v30, (_106_v31).Update(true, _105_v30), ((func() _dafny.Map {
			if _73_v6 {
				return _108_v33
			}
			return _108_v33
		})()).Cardinality(), _85_globalState)
		_109_v34 = _out0
		_110_v35 = _out1
		var _111_v36 _dafny.Map
		_ = _111_v36
		_111_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v9, false)
		var _112_v37 _dafny.Sequence
		_ = _112_v37
		_112_v37 = _dafny.SeqOf((_68_v1).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(236))).Cardinality()), _68_v1, Companion_Default___.SafeModuloInt(((_111_v36).Update(_76_v9, _109_v34)).Cardinality(), _68_v1))
		var _113_v38 D0
		_ = _113_v38
		_113_v38 = Companion_D0_.Create_DC0_(_79_v10)
		var _114_v39 _dafny.Map
		_ = _114_v39
		_114_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _110_v35)
		var _115_v40 _dafny.Map
		_ = _115_v40
		_115_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_113_v38).Dtor_cf0()).Cardinality(), _114_v39)
		_112_v37 = _dafny.Companion_Sequence_.Update(_112_v37, (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_112_v37).Cardinality()))).Uint32(), ((_115_v40).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _114_v39))).Cardinality())
		var _116_v41 *C0
		_ = _116_v41
		var _nw16 *C0 = New_C0_()
		_ = _nw16
		_nw16.Ctor__(_106_v31)
		_116_v41 = _nw16
		var _117_v42 *C0
		_ = _117_v42
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__(_106_v31)
		_117_v42 = _nw17
	} else {
		(_85_globalState).F20 = Companion_Default___.SafeDivisionInt((_68_v1).Minus((_dafny.Zero).Minus(_68_v1)), _68_v1)
		var _118_v43 _dafny.Array
		_ = _118_v43
		var _nwElement0_3 bool = _73_v6
		_ = _nwElement0_3
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(7))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_3, 0)
		(_nw18).ArraySet1(_73_v6, 1)
		(_nw18).ArraySet1(true, 2)
		(_nw18).ArraySet1(_73_v6, 3)
		(_nw18).ArraySet1(_73_v6, 4)
		(_nw18).ArraySet1(true, 5)
		(_nw18).ArraySet1(_73_v6, 6)
		_118_v43 = _nw18
		var _119_v44 _dafny.Map
		_ = _119_v44
		_119_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v6, _118_v43)
		var _120_v45 *C0
		_ = _120_v45
		var _nw19 *C0 = New_C0_()
		_ = _nw19
		_nw19.Ctor__(_119_v44)
		_120_v45 = _nw19
		_120_v45 = _120_v45
		var _121_v46 _dafny.Map
		_ = _121_v46
		_121_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _68_v1)
		var _122_v47 bool
		_ = _122_v47
		var _123_v48 bool
		_ = _123_v48
		var _out2 bool
		_ = _out2
		var _out3 bool
		_ = _out3
		_out2, _out3 = Companion_Default___.M0(_118_v43, (_120_v45).F22(), (func() _dafny.Int {
			if (_121_v46).Contains(_68_v1) {
				return (_121_v46).Get(_68_v1).(_dafny.Int)
			}
			return _68_v1
		})(), _85_globalState)
		_122_v47 = _out2
		_123_v48 = _out3
		var _124_v49 _dafny.Map
		_ = _124_v49
		_124_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v10, _dafny.IntOfUint32((_84_v15).Cardinality()))
		(_85_globalState).F14 = (func() _dafny.Int {
			if (_124_v49).Contains(_79_v10) {
				return (_124_v49).Get(_79_v10).(_dafny.Int)
			}
			return _68_v1
		})()
		var _125_v50 _dafny.Map
		_ = _125_v50
		_125_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_68_v1), _123_v48)
		_125_v50 = (_125_v50).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-604), _122_v47))
	}
	(_85_globalState).F14 = _68_v1
	var _126_v51 _dafny.Array
	_ = _126_v51
	var _nwElement0_4 bool = _73_v6
	_ = _nwElement0_4
	var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(23))
	_ = _nw20
	(_nw20).ArraySet1(_nwElement0_4, 0)
	(_nw20).ArraySet1(_73_v6, 1)
	(_nw20).ArraySet1(false, 2)
	(_nw20).ArraySet1(_73_v6, 3)
	(_nw20).ArraySet1(_73_v6, 4)
	(_nw20).ArraySet1(_73_v6, 5)
	(_nw20).ArraySet1(_73_v6, 6)
	(_nw20).ArraySet1(_73_v6, 7)
	(_nw20).ArraySet1(_73_v6, 8)
	(_nw20).ArraySet1(_73_v6, 9)
	(_nw20).ArraySet1(_73_v6, 10)
	(_nw20).ArraySet1(_73_v6, 11)
	(_nw20).ArraySet1(_73_v6, 12)
	(_nw20).ArraySet1(_73_v6, 13)
	(_nw20).ArraySet1(_73_v6, 14)
	(_nw20).ArraySet1(_73_v6, 15)
	(_nw20).ArraySet1(_73_v6, 16)
	(_nw20).ArraySet1(_73_v6, 17)
	(_nw20).ArraySet1(_73_v6, 18)
	(_nw20).ArraySet1(!((func() bool {
		if (_79_v10).Contains(true) {
			return (_79_v10).Get(true).(bool)
		}
		return _73_v6
	})()), 19)
	(_nw20).ArraySet1(_73_v6, 20)
	(_nw20).ArraySet1(_73_v6, 21)
	(_nw20).ArraySet1(_73_v6, 22)
	_126_v51 = _nw20
	var _127_v52 _dafny.Map
	_ = _127_v52
	_127_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _126_v51)
	var _128_v53 bool
	_ = _128_v53
	var _129_v54 bool
	_ = _129_v54
	var _out4 bool
	_ = _out4
	var _out5 bool
	_ = _out5
	_out4, _out5 = Companion_Default___.M0(_126_v51, _127_v52, (_68_v1).Times(_68_v1), _85_globalState)
	_128_v53 = _out4
	_129_v54 = _out5
	var _130_v56 _dafny.Sequence
	_ = _130_v56
	_130_v56 = _dafny.SeqOf(_79_v10, _79_v10)
	var _131_v57 _dafny.Map
	_ = _131_v57
	_131_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_130_v56).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _132_v55 _dafny.Map
			_132_v55 = interface{}(_compr_1).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_130_v56, _132_v55) {
				_coll1.Add(_132_v55, !(_128_v53))
			}
		}
		return _coll1.ToMap()
	}(), _128_v53)
	var _133_v58 _dafny.Map
	_ = _133_v58
	_133_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dyot")).Cardinality()), _68_v1)
	var _134_v59 _dafny.Map
	_ = _134_v59
	_134_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v54, _129_v54), Companion_Default___.Fm0(_133_v58, _68_v1, _128_v53, false, _85_globalState))
	(_85_globalState).F4 = (func() bool {
		if (_131_v57).Contains(_134_v59) {
			return (_131_v57).Get(_134_v59).(bool)
		}
		return Companion_Default___.Fm0(_133_v58, _68_v1, _129_v54, _129_v54, _85_globalState)
	})()
	(_85_globalState).F4 = _129_v54
	var _135_v60 D1
	_ = _135_v60
	_135_v60 = Companion_D1_.Create_DC2_(_127_v52)
	var _136_v61 *C0
	_ = _136_v61
	var _nw21 *C0 = New_C0_()
	_ = _nw21
	_nw21.Ctor__((_135_v60).Dtor_cf3())
	_136_v61 = _nw21
	_136_v61 = _136_v61
	var _hi0 _dafny.Int = (_dafny.Zero).Minus(_68_v1)
	_ = _hi0
	for _137_i4 := _68_v1; _137_i4.Cmp(_hi0) < 0; _137_i4 = _137_i4.Plus(_dafny.One) {
		var _138_v62 _dafny.Array
		_ = _138_v62
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw22
		_138_v62 = _nw22
		var _139_v63 _dafny.Map
		_ = _139_v63
		_139_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v53, _82_v13)
		var _140_v64 _dafny.Sequence
		_ = _140_v64
		_140_v64 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(282))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}((func(_141_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_142_i5 _dafny.Int) _dafny.CodePoint {
				return _141_v0
			}
		})(_67_v0))))
		var _143_v65 _dafny.Array
		_ = _143_v65
		var _nwElement0_5 _dafny.Set = _dafny.SetOf(_129_v54)
		_ = _nwElement0_5
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(14))
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_5, 0)
		(_nw23).ArraySet1(_82_v13, 1)
		(_nw23).ArraySet1(_82_v13, 2)
		(_nw23).ArraySet1(_82_v13, 3)
		(_nw23).ArraySet1(_82_v13, 4)
		(_nw23).ArraySet1(_82_v13, 5)
		(_nw23).ArraySet1(_82_v13, 6)
		(_nw23).ArraySet1(_82_v13, 7)
		(_nw23).ArraySet1(_82_v13, 8)
		(_nw23).ArraySet1((func() _dafny.Set {
			if (_139_v63).Contains(_129_v54) {
				return (_139_v63).Get(_129_v54).(_dafny.Set)
			}
			return Companion_Default___.Fm2(_129_v54, (_136_v61).Fm1(_129_v54, _dafny.Companion_Sequence_.Update(_140_v64, (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_140_v64).Cardinality()))).Uint32(), _84_v15), _85_globalState), _dafny.IntOfUint32((_84_v15).Cardinality()), _85_globalState)
		})(), 9)
		(_nw23).ArraySet1(_dafny.SetOf(_128_v53), 10)
		(_nw23).ArraySet1(_82_v13, 11)
		(_nw23).ArraySet1(_82_v13, 12)
		(_nw23).ArraySet1(_82_v13, 13)
		_143_v65 = _nw23
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_138_v62), 0))
		_ = _index16
		(_138_v62).ArraySet1(_143_v65, (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_138_v62), 0))
		_ = _index17
		(_138_v62).ArraySet1((func() _dafny.Array {
			if _129_v54 {
				return _143_v65
			}
			return _143_v65
		})(), (_index17).Int())
		var _144_v66 _dafny.Array
		_ = _144_v66
		_144_v66 = _69_v2
		var _145_v67 _dafny.MultiSet
		_ = _145_v67
		_145_v67 = _dafny.MultiSetOf((_144_v66), _76_v9, _76_v9, _69_v2, _76_v9)
		var _146_v68 _dafny.Sequence
		_ = _146_v68
		_146_v68 = _dafny.SeqOf(!(true), _129_v54)
		var _147_v69 _dafny.Map
		_ = _147_v69
		_147_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v67, _146_v68)
		_147_v69 = (_147_v69).Update(_dafny.MultiSetOf(_76_v9, _69_v2), _dafny.Companion_Sequence_.Concatenate(_146_v68, _146_v68))
		(_85_globalState).F4 = (func() bool {
			if _73_v6 {
				return (_128_v53) || ((_136_v61).Fm1(true, _140_v64, _85_globalState))
			}
			return _128_v53
		})()
		var _148_v70 D0
		_ = _148_v70
		_148_v70 = Companion_D0_.Create_DC0_(_79_v10)
		var _pat_let_tv1 = _79_v10
		_ = _pat_let_tv1
		var _149_v71 _dafny.Sequence
		_ = _149_v71
		_149_v71 = _dafny.SeqOf(Companion_D0_.Create_DC0_(_79_v10), func(_pat_let0_0 D0) D0 {
			return func(_150_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 _dafny.Map) D0 {
					return func(_151_dt__update_hcf0_h0 _dafny.Map) D0 {
						return Companion_D0_.Create_DC0_(_151_dt__update_hcf0_h0)
					}(_pat_let1_0)
				}(_pat_let_tv1)
			}(_pat_let0_0)
		}(_148_v70))
		var _rhs20 _dafny.Int = _dafny.IntOfInt64(495)
		_ = _rhs20
		var _rhs21 _dafny.Int = _137_i4
		_ = _rhs21
		var _rhs22 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_148_v70), _149_v71)
		_ = _rhs22
		var _rhs23 _dafny.Int = ((_68_v1).Plus(_137_i4)).Times((_68_v1).Times(_137_i4))
		_ = _rhs23
		var _rhs24 _dafny.Int = _68_v1
		_ = _rhs24
		var _lhs16 *GlobalState = _85_globalState
		_ = _lhs16
		var _lhs17 *GlobalState = _85_globalState
		_ = _lhs17
		var _lhs18 *GlobalState = _85_globalState
		_ = _lhs18
		_68_v1 = _rhs20
		_lhs16.F14 = _rhs21
		_149_v71 = _rhs22
		_lhs17.F7 = _rhs23
		_lhs18.F20 = _rhs24
	}
	if Companion_Default___.Fm0(_133_v58, _68_v1, _129_v54, _73_v6, _85_globalState) {
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_81_v12), 0))
		_ = _index18
		(_81_v12).ArraySet1(_84_v15, (_index18).Int())
		var _152_v72 _dafny.MultiSet
		_ = _152_v72
		_152_v72 = _dafny.MultiSetOf(_dafny.IntOfInt64(746), _68_v1)
		var _153_v73 _dafny.Sequence
		_ = _153_v73
		_153_v73 = _dafny.SeqOf((_152_v72).Cardinality())
		var _154_v74 _dafny.Sequence
		_ = _154_v74
		_154_v74 = _dafny.SeqOf(_153_v73, _153_v73, _153_v73)
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_81_v12), 0))
		_ = _index19
		var _rhs25 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((_154_v74).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_154_v74).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.Fm3(_85_globalState))
		_ = _rhs25
		var _rhs26 _dafny.Int = (_68_v1).Minus(_68_v1)
		_ = _rhs26
		var _rhs27 bool = !(((func() _dafny.Int {
			if (_152_v72).Contains(_68_v1) {
				return (_152_v72).Multiplicity(_68_v1)
			}
			return _68_v1
		})()).Cmp(_68_v1) == 0)
		_ = _rhs27
		var _rhs28 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('f')), _84_v15), (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('f')), _84_v15)).Cardinality()))).Uint32(), _67_v0), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_68_v1, _68_v1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('f')), _84_v15), (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('f')), _84_v15)).Cardinality()))).Uint32(), _67_v0)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
			if _129_v54 {
				return _67_v0
			}
			return _67_v0
		})())
		_ = _rhs28
		var _lhs19 *GlobalState = _85_globalState
		_ = _lhs19
		var _lhs20 _dafny.Array = _81_v12
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_81_v12), 0))
		_ = _lhs21
		_129_v54 = _rhs25
		_lhs19.F7 = _rhs26
		_129_v54 = _rhs27
		(_lhs20).ArraySet1(_rhs28, (_lhs21).Int())
		var _155_v75 _dafny.Sequence
		_ = _155_v75
		_155_v75 = _dafny.SeqOf(_127_v52, _127_v52)
		var _156_v76 *C0
		_ = _156_v76
		var _nw24 *C0 = New_C0_()
		_ = _nw24
		_nw24.Ctor__((_155_v75).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_155_v75).Cardinality()))).Uint32()).(_dafny.Map))
		_156_v76 = _nw24
		var _157_v77 *C0
		_ = _157_v77
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__((_136_v61).F22())
		_157_v77 = _nw25
		_133_v58 = Companion_Default___.Fm4((func() _dafny.Int {
			if _128_v53 {
				return _68_v1
			}
			return _68_v1
		})(), _68_v1, _85_globalState)
		var _158_v78 _dafny.Set
		_ = _158_v78
		_158_v78 = _dafny.SetOf(_67_v0, _67_v0)
		var _159_v79 _dafny.Map
		_ = _159_v79
		_159_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_79_v10), (_158_v78).Cardinality())
		var _160_v80 _dafny.Map
		_ = _160_v80
		_160_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v54, _152_v72)
		var _161_v81 bool
		_ = _161_v81
		var _162_v82 bool
		_ = _162_v82
		var _out6 bool
		_ = _out6
		var _out7 bool
		_ = _out7
		_out6, _out7 = Companion_Default___.M0(_126_v51, (_156_v76).F22(), ((func() _dafny.Int {
			if (_159_v79).Contains(Companion_Default___.Fm5(_152_v72, _85_globalState)) {
				return (_159_v79).Get(Companion_Default___.Fm5(_152_v72, _85_globalState)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(315)
		})()).Minus(((func() _dafny.MultiSet {
			if (_160_v80).Contains(_73_v6) {
				return (_160_v80).Get(_73_v6).(_dafny.MultiSet)
			}
			return _152_v72
		})()).Cardinality()), _85_globalState)
		_161_v81 = _out6
		_162_v82 = _out7
	} else {
		if ((_136_v61).Fm1(_128_v53, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-683))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_163_i6 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("u")
		})), _85_globalState)) && ((_68_v1).Cmp(_68_v1) >= 0) {
			var _164_v83 _dafny.Map
			_ = _164_v83
			_164_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _84_v15)
			_84_v15 = (func() _dafny.Sequence {
				if (_164_v83).Contains(_68_v1) {
					return (_164_v83).Get(_68_v1).(_dafny.Sequence)
				}
				return _84_v15
			})()
			var _165_v84 *C0
			_ = _165_v84
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__((_136_v61).F22())
			_165_v84 = _nw26
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_76_v9), 0))
			_ = _index20
			(_76_v9).ArraySet1(_68_v1, (_index20).Int())
			var _166_v85 D0
			_ = _166_v85
			_166_v85 = Companion_D0_.Create_DC1_(_68_v1, _129_v54)
			var _167_v86 _dafny.MultiSet
			_ = _167_v86
			_167_v86 = _dafny.MultiSetOf(_166_v85)
			var _168_v87 _dafny.Map
			_ = _168_v87
			_168_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v86, _129_v54)
			var _169_v88 _dafny.Array
			_ = _169_v88
			var _nwElement0_6 _dafny.Map = _168_v87
			_ = _nwElement0_6
			var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(2))
			_ = _nw27
			(_nw27).ArraySet1(_nwElement0_6, 0)
			(_nw27).ArraySet1(_168_v87, 1)
			_169_v88 = _nw27
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_169_v88), 0))
			_ = _index21
			(_169_v88).ArraySet1((_168_v87).Merge(_168_v87), (_index21).Int())
			var _170_v89 _dafny.Sequence
			_ = _170_v89
			_170_v89 = _dafny.SeqOf(_73_v6, _129_v54, !(!(true)))
			var _171_v90 _dafny.Array
			_ = _171_v90
			_171_v90 = _69_v2
			var _172_v91 _dafny.Array
			_ = _172_v91
			_172_v91 = (_171_v90)
			var _173_v92 _dafny.Sequence
			_ = _173_v92
			_173_v92 = _dafny.SeqOf(_dafny.IntOfInt64(-376), _dafny.IntOfInt64(212))
			var _174_v93 _dafny.Map
			_ = _174_v93
			_174_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v91, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_173_v92, (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_173_v92).Cardinality()))).Uint32(), _68_v1)).Cardinality())))
			var _175_v94 D0
			_ = _175_v94
			_175_v94 = Companion_D0_.Create_DC0_(_79_v10)
			var _176_v95 _dafny.Map
			_ = _176_v95
			_176_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v94, _136_v61)
			var _177_v96 _dafny.Map
			_ = _177_v96
			_177_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_174_v93).Merge(_174_v93)).Cardinality(), (func() *C0 {
				if (_176_v95).Contains(_175_v94) {
					return (_176_v95).Get(_175_v94).(*C0)
				}
				return _165_v84
			})())
			var _178_v97 D3
			_ = _178_v97
			_178_v97 = Companion_D3_.Create_DC6_(_165_v84)
			var _179_v98 _dafny.MultiSet
			_ = _179_v98
			_179_v98 = _dafny.MultiSetOf(_82_v13)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_76_v9), 0))
			_ = _index22
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_169_v88), 0))
			_ = _index23
			var _rhs29 _dafny.Int = _68_v1
			_ = _rhs29
			var _rhs30 *C0 = (func() *C0 {
				if (_177_v96).Contains(_68_v1) {
					return (_177_v96).Get(_68_v1).(*C0)
				}
				return (_178_v97).Dtor_cf8()
			})()
			_ = _rhs30
			var _rhs31 _dafny.Int = (_179_v98).Cardinality()
			_ = _rhs31
			var _rhs32 _dafny.Map = ((_168_v87).Merge(_168_v87)).Merge(_168_v87)
			_ = _rhs32
			var _rhs33 _dafny.Sequence = _170_v89
			_ = _rhs33
			var _lhs22 _dafny.Array = _76_v9
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_76_v9), 0))
			_ = _lhs23
			var _lhs24 *GlobalState = _85_globalState
			_ = _lhs24
			var _lhs25 _dafny.Array = _169_v88
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_169_v88), 0))
			_ = _lhs26
			(_lhs22).ArraySet1(_rhs29, (_lhs23).Int())
			_136_v61 = _rhs30
			_lhs24.F7 = _rhs31
			(_lhs25).ArraySet1(_rhs32, (_lhs26).Int())
			_170_v89 = _rhs33
			(_85_globalState).F4 = (!(_129_v54)) || (_129_v54)
			_84_v15 = _84_v15
		} else {
			var _180_v99 *C0
			_ = _180_v99
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__((_136_v61).F22())
			_180_v99 = _nw28
			_84_v15 = _84_v15
			_129_v54 = _129_v54
			var _181_v100 *C0
			_ = _181_v100
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__((_136_v61).F22())
			_181_v100 = _nw29
			var _182_v101 _dafny.Sequence
			_ = _182_v101
			_182_v101 = _dafny.SeqOf(_181_v100)
			var _183_v102 _dafny.Array
			_ = _183_v102
			var _nwElement0_7 *C0 = _136_v61
			_ = _nwElement0_7
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(21))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_7, 0)
			(_nw30).ArraySet1((_182_v101).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_182_v101).Cardinality()))).Uint32()).(*C0), 1)
			(_nw30).ArraySet1(_136_v61, 2)
			(_nw30).ArraySet1(_181_v100, 3)
			(_nw30).ArraySet1(_181_v100, 4)
			(_nw30).ArraySet1(_180_v99, 5)
			(_nw30).ArraySet1(_136_v61, 6)
			(_nw30).ArraySet1(_136_v61, 7)
			(_nw30).ArraySet1(_181_v100, 8)
			(_nw30).ArraySet1(_181_v100, 9)
			(_nw30).ArraySet1(_136_v61, 10)
			(_nw30).ArraySet1(_181_v100, 11)
			(_nw30).ArraySet1(_180_v99, 12)
			(_nw30).ArraySet1(_181_v100, 13)
			(_nw30).ArraySet1(_180_v99, 14)
			(_nw30).ArraySet1(_136_v61, 15)
			(_nw30).ArraySet1(_136_v61, 16)
			(_nw30).ArraySet1(_180_v99, 17)
			(_nw30).ArraySet1(_181_v100, 18)
			(_nw30).ArraySet1(_181_v100, 19)
			(_nw30).ArraySet1(_181_v100, 20)
			_183_v102 = _nw30
			var _184_v103 _dafny.Set
			_ = _184_v103
			_184_v103 = _dafny.SetOf(_183_v102)
			var _rhs34 bool = !(_128_v53)
			_ = _rhs34
			var _rhs35 bool = _73_v6
			_ = _rhs35
			var _rhs36 _dafny.Int = (((_184_v103).Union(_184_v103)).Union((_184_v103).Difference(_dafny.SetOf(_183_v102)))).Cardinality()
			_ = _rhs36
			var _lhs27 *GlobalState = _85_globalState
			_ = _lhs27
			var _lhs28 *GlobalState = _85_globalState
			_ = _lhs28
			_73_v6 = _rhs34
			_lhs27.F6 = _rhs35
			_lhs28.F14 = _rhs36
		}
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))
		_ = _index24
		(_76_v9).ArraySet1(_68_v1, (_index24).Int())
		var _185_v104 _dafny.Map
		_ = _185_v104
		_185_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v54, _76_v9)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))
		_ = _index25
		(_76_v9).ArraySet1((_185_v104).Cardinality(), (_index25).Int())
		var _186_v105 _dafny.Map
		_ = _186_v105
		_186_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_84_v15).Cardinality()), _129_v54)
		var _187_v106 _dafny.Map
		_ = _187_v106
		_187_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-955), !((func() bool {
			if (_186_v105).Contains(_68_v1) {
				return (_186_v105).Get(_68_v1).(bool)
			}
			return _73_v6
		})()))
		_187_v106 = _187_v106
		(_85_globalState).F6 = true
		var _188_v107 _dafny.Sequence
		_ = _188_v107
		_188_v107 = _dafny.SeqOf(_129_v54, _128_v53, _129_v54)
		var _189_v108 _dafny.Sequence
		_ = _189_v108
		_189_v108 = _dafny.SeqOf((_76_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))).Int()).(_dafny.Int))
		var _190_v109 _dafny.Map
		_ = _190_v109
		_190_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v107, (Companion_Default___.Fm6(_189_v108, _128_v53, _73_v6, _85_globalState)).Cardinality())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))
		_ = _index26
		var _rhs37 _dafny.Int = (func() _dafny.Int {
			if _129_v54 {
				return ((_76_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))).Int()).(_dafny.Int)).Minus((_76_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))).Int()).(_dafny.Int))
			}
			return (_76_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))).Int()).(_dafny.Int)
		})()
		_ = _rhs37
		var _rhs38 bool = !(_73_v6)
		_ = _rhs38
		var _rhs39 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_190_v109).Contains(_dafny.Companion_Sequence_.Concatenate(_188_v107, _188_v107)) {
				return (_190_v109).Get(_dafny.Companion_Sequence_.Concatenate(_188_v107, _188_v107)).(_dafny.Int)
			}
			return (_76_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))).Int()).(_dafny.Int)
		})())
		_ = _rhs39
		var _rhs40 _dafny.Int = (_76_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))).Int()).(_dafny.Int)
		_ = _rhs40
		var _lhs29 _dafny.Array = _76_v9
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_76_v9), 0))
		_ = _lhs30
		var _lhs31 *GlobalState = _85_globalState
		_ = _lhs31
		var _lhs32 *GlobalState = _85_globalState
		_ = _lhs32
		var _lhs33 *GlobalState = _85_globalState
		_ = _lhs33
		(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
		_lhs31.F4 = _rhs38
		_lhs32.F14 = _rhs39
		_lhs33.F14 = _rhs40
	}
	(_85_globalState).F14 = (_68_v1).Times(_68_v1)
	var _191_i7 _dafny.Int
	_ = _191_i7
	_191_i7 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.IsProperPrefixOf(_84_v15, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pt"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(834))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_207_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_208_i8 _dafny.Int) _dafny.CodePoint {
				return _207_v0
			}
		})(_67_v0))))) {
			{
				if (_191_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_191_i7 = (_191_i7).Plus(_dafny.One)
				var _192_v110 D3
				_ = _192_v110
				_192_v110 = Companion_D3_.Create_DC6_(_136_v61)
				_192_v110 = Companion_D3_.Create_DC6_((_192_v110).Dtor_cf8())
				var _193_v111 *C0
				_ = _193_v111
				var _nw31 *C0 = New_C0_()
				_ = _nw31
				_nw31.Ctor__(_127_v52)
				_193_v111 = _nw31
				var _194_v112 _dafny.Sequence
				_ = _194_v112
				_194_v112 = (func() _dafny.Sequence {
					if (_102_v27).Contains(_73_v6) {
						return (_102_v27).Get(_73_v6).(_dafny.Sequence)
					}
					return _84_v15
				})()
				_84_v15 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_194_v112), _84_v15), _84_v15)
				var _195_v113 _dafny.Sequence
				_ = _195_v113
				_195_v113 = _dafny.SeqOf(_84_v15, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(285))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}(func(_196_i10 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), _dafny.UnicodeSeqOfUtf8Bytes("ikudn"))
				if Companion_Default___.Fm0(_133_v58, ((_dafny.Zero).Minus((_dafny.Zero).Minus(_68_v1))).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(6))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_197_i9 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(90)
				}))).Cardinality())), _128_v53, (_193_v111).Fm1(_73_v6, _195_v113, _85_globalState), _85_globalState) {
					(_85_globalState).F14 = (_68_v1).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(915), _68_v1))
					_84_v15 = _dafny.UnicodeSeqOfUtf8Bytes("sv")
					_84_v15 = _dafny.UnicodeSeqOfUtf8Bytes("jbjlpyup")
					var _198_v114 *C0
					_ = _198_v114
					var _nw32 *C0 = New_C0_()
					_ = _nw32
					_nw32.Ctor__(((_136_v61).F22()).Merge(((_136_v61).F22()).Update(_73_v6, _126_v51)))
					_198_v114 = _nw32
					var _199_v115 _dafny.Sequence
					_ = _199_v115
					_199_v115 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v111, _193_v111)).Cardinality(), _68_v1)
					var _200_v117 _dafny.Map
					_ = _200_v117
					_200_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_84_v15, Companion_Default___.Fm7(func() _dafny.Set {
						var _coll2 = _dafny.NewBuilder()
						_ = _coll2
						for _iter2 := _dafny.Iterate((_199_v115).Elements()); ; {
							_compr_2, _ok2 := _iter2()
							if !_ok2 {
								break
							}
							var _201_v116 _dafny.Int
							_201_v116 = interface{}(_compr_2).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_199_v115, _201_v116) {
								_coll2.Add((_201_v116).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(783))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg10 _dafny.Int) interface{} {
										return coer10(arg10)
									}
								}(func(_202_i11 _dafny.Int) _dafny.Int {
									return (_dafny.Zero).Minus(_dafny.IntOfInt64(-117))
								}))).Cardinality())))
							}
						}
						return _coll2.ToSet()
					}(), _85_globalState)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _68_v1))
					var _203_v118 D3
					_ = _203_v118
					_203_v118 = Companion_D3_.Create_DC8_(Companion_D3_.Create_DC7_())
					var _rhs41 bool = _73_v6
					_ = _rhs41
					var _rhs42 _dafny.Map = _200_v117
					_ = _rhs42
					var _rhs43 _dafny.Array = _76_v9
					_ = _rhs43
					var _rhs44 D3 = _203_v118
					_ = _rhs44
					var _lhs34 *GlobalState = _85_globalState
					_ = _lhs34
					_lhs34.F4 = _rhs41
					_200_v117 = _rhs42
					_76_v9 = _rhs43
					_203_v118 = _rhs44
				} else {
					_84_v15 = _84_v15
					var _204_v119 _dafny.Array
					_ = _204_v119
					var _nw33 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(3))
					_ = _nw33
					_204_v119 = _nw33
					var _rhs45 *C0 = (Companion_D3_.Create_DC6_(_193_v111)).Dtor_cf8()
					_ = _rhs45
					var _rhs46 _dafny.Array = _204_v119
					_ = _rhs46
					var _rhs47 _dafny.Sequence = _84_v15
					_ = _rhs47
					_136_v61 = _rhs45
					_204_v119 = _rhs46
					_84_v15 = _rhs47
					var _205_v120 *C0
					_ = _205_v120
					var _nw34 *C0 = New_C0_()
					_ = _nw34
					_nw34.Ctor__(_127_v52)
					_205_v120 = _nw34
					var _206_v121 *C0
					_ = _206_v121
					var _nw35 *C0 = New_C0_()
					_ = _nw35
					_nw35.Ctor__((_136_v61).F22())
					_206_v121 = _nw35
					_75_v8 = (_75_v8).Update(_129_v54, _68_v1)
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _hi1 _dafny.Int = _dafny.IntOfInt64(492)
	_ = _hi1
	for _209_i12 := (func() _dafny.Int {
		if _73_v6 {
			return (func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(372), _dafny.IntOfInt64(739))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _225_v122 _dafny.Int
					_225_v122 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(372)).Cmp(_225_v122) <= 0) && ((_225_v122).Cmp(_dafny.IntOfInt64(739)) < 0) {
						_coll3.Add(Companion_Default___.SafeDivisionInt(_225_v122, _68_v1), _128_v53)
					}
				}
				return _coll3.ToMap()
			}()).Cardinality()
		}
		return _68_v1
	})(); _209_i12.Cmp(_hi1) < 0; _209_i12 = _209_i12.Plus(_dafny.One) {
		var _210_v123 D3
		_ = _210_v123
		_210_v123 = Companion_D3_.Create_DC6_(_136_v61)
		var _211_v124 D3
		_ = _211_v124
		_211_v124 = Companion_D3_.Create_DC8_(_210_v123)
		var _pat_let_tv2 = _210_v123
		_ = _pat_let_tv2
		var _pat_let_tv3 = _210_v123
		_ = _pat_let_tv3
		var _212_v125 _dafny.Array
		_ = _212_v125
		var _nwElement0_8 D3 = _211_v124
		_ = _nwElement0_8
		var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(6))
		_ = _nw36
		(_nw36).ArraySet1(_nwElement0_8, 0)
		(_nw36).ArraySet1(_211_v124, 1)
		(_nw36).ArraySet1(_211_v124, 2)
		(_nw36).ArraySet1(func(_pat_let2_0 D3) D3 {
			return func(_213_dt__update__tmp_h1 D3) D3 {
				return func(_pat_let3_0 D3) D3 {
					return func(_214_dt__update_hcf9_h0 D3) D3 {
						return Companion_D3_.Create_DC8_(_214_dt__update_hcf9_h0)
					}(_pat_let3_0)
				}(_pat_let_tv2)
			}(_pat_let2_0)
		}(_211_v124), 3)
		(_nw36).ArraySet1(func(_pat_let4_0 D3) D3 {
			return func(_215_dt__update__tmp_h2 D3) D3 {
				return func(_pat_let5_0 D3) D3 {
					return func(_216_dt__update_hcf9_h1 D3) D3 {
						return Companion_D3_.Create_DC8_(_216_dt__update_hcf9_h1)
					}(_pat_let5_0)
				}(_pat_let_tv3)
			}(_pat_let4_0)
		}(_211_v124), 4)
		(_nw36).ArraySet1(_211_v124, 5)
		_212_v125 = _nw36
		var _217_v126 _dafny.Map
		_ = _217_v126
		_217_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v125, _76_v9)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _index27
		(_126_v51).ArraySet1(_129_v54, (_index27).Int())
		var _218_v127 _dafny.Map
		_ = _218_v127
		_218_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_67_v0, _136_v61)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _index28
		var _rhs48 _dafny.Map = ((_217_v126).Update(_212_v125, _69_v2)).Merge((_217_v126).Merge(_217_v126))
		_ = _rhs48
		var _rhs49 _dafny.Array = _126_v51
		_ = _rhs49
		var _rhs50 *C0 = (func() *C0 {
			if (_218_v127).Contains(_dafny.CodePoint('t')) {
				return (_218_v127).Get(_dafny.CodePoint('t')).(*C0)
			}
			return _136_v61
		})()
		_ = _rhs50
		var _rhs51 bool = (_129_v54) || (_73_v6)
		_ = _rhs51
		var _lhs35 _dafny.Array = _126_v51
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _lhs36
		_217_v126 = _rhs48
		_126_v51 = _rhs49
		_136_v61 = _rhs50
		(_lhs35).ArraySet1(_rhs51, (_lhs36).Int())
		var _219_v128 *C0
		_ = _219_v128
		var _nw37 *C0 = New_C0_()
		_ = _nw37
		_nw37.Ctor__((_136_v61).F22())
		_219_v128 = _nw37
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_69_v2), 0))
		_ = _index29
		(_69_v2).ArraySet1(_68_v1, (_index29).Int())
		var _220_v129 _dafny.MultiSet
		_ = _220_v129
		_220_v129 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(198))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_221_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_222_i13 _dafny.Int) _dafny.CodePoint {
				return _221_v0
			}
		})(_67_v0)))).Cardinality()), _209_i12)
		var _223_v130 _dafny.Map
		_ = _223_v130
		_223_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_220_v129).Difference(_220_v129), _dafny.SeqOf(_dafny.IntOfInt64(653)))
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_69_v2), 0))
		_ = _index30
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _index31
		var _rhs52 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqOf(_68_v1, _209_i12, _68_v1, _209_i12, _dafny.IntOfInt64(658))).Cardinality())).Minus(Companion_Default___.SafeModuloInt(_68_v1, _68_v1))
		_ = _rhs52
		var _rhs53 _dafny.Map = _223_v130
		_ = _rhs53
		var _rhs54 *C0 = (func() *C0 {
			if !(_73_v6) {
				return _136_v61
			}
			return _219_v128
		})()
		_ = _rhs54
		var _rhs55 bool = !((_126_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))).Int()).(bool))
		_ = _rhs55
		var _lhs37 _dafny.Array = _69_v2
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_69_v2), 0))
		_ = _lhs38
		var _lhs39 _dafny.Array = _126_v51
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _lhs40
		(_lhs37).ArraySet1(_rhs52, (_lhs38).Int())
		_223_v130 = _rhs53
		_219_v128 = _rhs54
		(_lhs39).ArraySet1(_rhs55, (_lhs40).Int())
		var _224_v131 _dafny.Sequence
		_ = _224_v131
		_224_v131 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("uiiomhn"))
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _index32
		var _rhs56 bool = (func() bool {
			if (_126_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))).Int()).(bool) {
				return Companion_Default___.Fm0(_133_v58, _68_v1, (_219_v128).Fm1(!(_128_v53), _dafny.SeqOf(Companion_Default___.Fm7(_71_v4, _85_globalState), _dafny.UnicodeSeqOfUtf8Bytes("odeavklj")), _85_globalState), (_219_v128).Fm1(true, _224_v131, _85_globalState), _85_globalState)
			}
			return _128_v53
		})()
		_ = _rhs56
		var _rhs57 _dafny.Int = (_69_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_69_v2), 0))).Int()).(_dafny.Int)
		_ = _rhs57
		var _lhs41 _dafny.Array = _126_v51
		_ = _lhs41
		var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_126_v51), 0))
		_ = _lhs42
		var _lhs43 *GlobalState = _85_globalState
		_ = _lhs43
		(_lhs41).ArraySet1(_rhs56, (_lhs42).Int())
		_lhs43.F20 = _rhs57
	}
	_dafny.Print(_67_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v2).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_70_v3).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_v4).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(557))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_72_v5).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(557)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_73_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_74_v7).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(310249))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_76_v9).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_79_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_80_v11).Equals(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(557)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v13).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_83_v14, _dafny.SeqOf(_dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_84_v15.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x')), _dafny.CodePoint('w'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F1()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F2()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F3()).Equals(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(557)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F5()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F8()).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F12()).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(557))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F13()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F16()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F17()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F17()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_85_globalState).F18()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(2)).UpdateUnsafe(true, _dafny.IntOfInt64(557))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_globalState).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_globalState.F20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_85_globalState).F21().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("vfucpd"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v51).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v52).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_128_v53)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_129_v54)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_130_v56, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v57).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), true), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v58).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(387), _dafny.IntOfInt64(-10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v59).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v60).Dtor_cf3()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_136_v61).F22()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_191_i7)
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Map
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf0() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf0
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2 == data2.Cf2
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

type D1_DC3 struct {
	Cf4 bool
	Cf5 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 bool, Cf5 bool) D1 {
	return D1{D1_DC3{Cf4, Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf3 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf3 _dafny.Map) D1 {
	return D1{D1_DC2{Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC4 struct {
	Cf6 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 D1) D1 {
	return D1{D1_DC4{Cf6}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false, false)
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf3() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf6() D1 {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf6) + ")"
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
			return ok && data1.Cf4 == data2.Cf4 && data1.Cf5 == data2.Cf5
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf3.Equals(data2.Cf3)
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf7 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf7 _dafny.Array) D2 {
	return D2{D2_DC5{Cf7}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf7) + ")"
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
			return ok && data1.Cf7 == data2.Cf7
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
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_() D3 {
	return D3{D3_DC7{}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC6 struct {
	Cf8 *C0
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf8 *C0) D3 {
	return D3{D3_DC6{Cf8}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC8 struct {
	Cf9 D3
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf9 D3) D3 {
	return D3{D3_DC8{Cf9}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_()
}

func (_this D3) Dtor_cf8() *C0 {
	return _this.Get_().(D3_DC6).Cf8
}

func (_this D3) Dtor_cf9() D3 {
	return _this.Get_().(D3_DC8).Cf9
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf9) + ")"
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
			_, ok := other.Get_().(D3_DC7)
			return ok
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf8 == data2.Cf8
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf9.Equals(data2.Cf9)
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

type D4_DC9 struct {
	Cf10 _dafny.Sequence
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf10 _dafny.Sequence) D4 {
	return D4{D4_DC9{Cf10}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D4_DC9).Cf10
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + data.Cf10.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D5_DC10 struct {
	Cf11 _dafny.Set
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf11 _dafny.Set) D5 {
	return D5{D5_DC10{Cf11}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D5) Dtor_cf11() _dafny.Set {
	return _this.Get_().(D5_DC10).Cf11
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D6_DC11 struct {
	Cf12 _dafny.Sequence
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf12 _dafny.Sequence) D6 {
	return D6{D6_DC11{Cf12}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D6) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D6_DC11).Cf12
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	F4   bool
	F6   bool
	F7   _dafny.Int
	F14  _dafny.Int
	F20  _dafny.Int
	_f0  _dafny.Map
	_f1  _dafny.Array
	_f2  _dafny.Map
	_f3  _dafny.Set
	_f5  _dafny.Array
	_f8  _dafny.MultiSet
	_f9  _dafny.Int
	_f10 _dafny.Int
	_f11 bool
	_f12 _dafny.Set
	_f13 _dafny.Array
	_f15 bool
	_f16 _dafny.Set
	_f17 _dafny.Array
	_f18 _dafny.Map
	_f19 _dafny.Int
	_f21 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = false
	_this.F6 = false
	_this.F7 = _dafny.Zero
	_this.F14 = _dafny.Zero
	_this.F20 = _dafny.Zero
	_this._f0 = _dafny.EmptyMap
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = _dafny.EmptyMap
	_this._f3 = _dafny.EmptySet
	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptyMultiSet
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f12 = _dafny.EmptySet
	_this._f13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = false
	_this._f16 = _dafny.EmptySet
	_this._f17 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f18 = _dafny.EmptyMap
	_this._f19 = _dafny.Zero
	_this._f21 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 _dafny.Array, f2 _dafny.Map, f3 _dafny.Set, f4 bool, f5 _dafny.Array, f6 bool, f7 _dafny.Int, f8 _dafny.MultiSet, f9 _dafny.Int, f10 _dafny.Int, f11 bool, f12 _dafny.Set, f13 _dafny.Array, f14 _dafny.Int, f15 bool, f16 _dafny.Set, f17 _dafny.Array, f18 _dafny.Map, f19 _dafny.Int, f20 _dafny.Int, f21 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this).F20 = f20
		(_this)._f21 = f21
	}
}
func (_this *GlobalState) F0() _dafny.Map {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Map {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Set {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.Array {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F8() _dafny.MultiSet {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
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
func (_this *GlobalState) F12() _dafny.Set {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Array {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Set {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Array {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Map {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() _dafny.Int {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F21() _dafny.Sequence {
	{
		return _this._f21
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f22 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f22 = _dafny.EmptyMap
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

func (_this *C0) Ctor__(f22 _dafny.Map) {
	{
		(_this)._f22 = f22
	}
}
func (_this *C0) Fm1(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) bool {
	{
		var _source2 D0 = Companion_D0_.Create_DC1_((_dafny.SetOf(!(false), true, true)).Cardinality(), !(false))
		_ = _source2
		if _source2.Is_DC1() {
			var _226___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
			_ = _226___mcc_h0
			var _227___mcc_h1 bool = _source2.Get_().(D0_DC1).Cf2
			_ = _227___mcc_h1
			var _228_cf2 bool = _227___mcc_h1
			_ = _228_cf2
			var _229_cf1 _dafny.Int = _226___mcc_h0
			_ = _229_cf1
			return (Companion_D0_.Create_DC1_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mpsgaxsnn")).Cardinality())), _228_cf2)).Dtor_cf2()
		} else {
			var _230___mcc_h2 _dafny.Map = _source2.Get_().(D0_DC0).Cf0
			_ = _230___mcc_h2
			var _231_cf0 _dafny.Map = _230___mcc_h2
			_ = _231_cf0
			return !(!(false)) || (true)
		}
	}
}
func (_this *C0) F22() _dafny.Map {
	{
		return _this._f22
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
