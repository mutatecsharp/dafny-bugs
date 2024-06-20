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
	return _dafny.IntOfInt64(991)
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(409), _dafny.IntOfInt64(-296))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(409)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(-296)) < 0) {
				_coll0.Add((_0_v0).Plus(_dafny.IntOfInt64(260)), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(202))).Cardinality())).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-123))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg0 _dafny.Int) interface{} {
						return coer0(arg0)
					}
				}(func(_1_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(65)
				}))).Cardinality()), !(true))).Cardinality())).Cardinality(), true)))))
			}
		}
		return _coll0.ToMap()
	}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(944), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(false, _dafny.IntOfInt64(754), _dafny.IntOfInt64(561), false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(-113))).Cardinality(), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(true, _dafny.IntOfInt64(111), _dafny.IntOfInt64(-539), !(true)))))))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(466), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(696))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D3_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("xkxmuj"))).Dtor_cf10()
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('n')
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true), _dafny.SeqOf(!(false))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(false))))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(true))).Union(_dafny.SetOf(false))).Difference(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, !(!(false)), false), _dafny.SeqOf(true, false, false, !(false))), _dafny.SeqOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("nw"), _dafny.UnicodeSeqOfUtf8Bytes("fbvst"), _dafny.UnicodeSeqOfUtf8Bytes("dmftxyldf"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})))).IsSubsetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-760))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))))).Difference(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("hgiqbg"), _dafny.UnicodeSeqOfUtf8Bytes("gxnfl"))))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.Map {
	if (func() bool {
		if !(false) {
			return true
		}
		return false
	})() {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-835), _dafny.IntOfUint32((_dafny.SeqOf(!(false), !(false))).Cardinality()))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(156), (_dafny.MultiSetOf(!(true), !(true))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('b'), _dafny.CodePoint('g'))
}
func (_static *CompanionStruct_Default___) M2(p0 bool, p1 bool, globalState *GlobalState) {
	if p0 {
		var _4_v0 _dafny.Int
		_ = _4_v0
		_4_v0 = _dafny.IntOfInt64(274)
		_4_v0 = _4_v0
		var _5_v1 _dafny.Set
		_ = _5_v1
		_5_v1 = _dafny.SetOf(_4_v0)
		if ((_5_v1).IsProperSubsetOf(_5_v1)) || (!(!(p0)) || (p0)) {
			var _6_v2 _dafny.CodePoint
			_ = _6_v2
			_6_v2 = _dafny.CodePoint('e')
			(globalState).F18 = _6_v2
			var _7_v3 _dafny.Sequence
			_ = _7_v3
			_7_v3 = _dafny.SeqOf(_dafny.SeqOf(p0))
			var _8_v4 *C0
			_ = _8_v4
			var _nw0 *C0 = New_C0_()
			_ = _nw0
			_nw0.Ctor__(!(p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_4_v0, (_7_v3).Select((Companion_Default___.SafeIndex(_4_v0, _dafny.IntOfUint32((_7_v3).Cardinality()))).Uint32()).(_dafny.Sequence)))
			_8_v4 = _nw0
			var _9_v5 _dafny.Set
			_ = _9_v5
			_9_v5 = _dafny.SetOf(_8_v4)
			(globalState).F19 = !((_9_v5).IsProperSubsetOf(_9_v5))
			var _10_v6 _dafny.MultiSet
			_ = _10_v6
			_10_v6 = _dafny.MultiSetOf(_8_v4.F24, p1)
			var _11_v7 _dafny.Array
			_ = _11_v7
			var _nwElement0_0 bool = p1
			_ = _nwElement0_0
			var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(22))
			_ = _nw1
			(_nw1).ArraySet1(_nwElement0_0, 0)
			(_nw1).ArraySet1(true, 1)
			(_nw1).ArraySet1(p0, 2)
			(_nw1).ArraySet1(!(true), 3)
			(_nw1).ArraySet1(p1, 4)
			(_nw1).ArraySet1(_8_v4.F24, 5)
			(_nw1).ArraySet1(p0, 6)
			(_nw1).ArraySet1(!(p0), 7)
			(_nw1).ArraySet1(false, 8)
			(_nw1).ArraySet1(p1, 9)
			(_nw1).ArraySet1(p0, 10)
			(_nw1).ArraySet1(_8_v4.F24, 11)
			(_nw1).ArraySet1(Companion_Default___.Fm11(_10_v6, _4_v0, globalState), 12)
			(_nw1).ArraySet1(p0, 13)
			(_nw1).ArraySet1(p0, 14)
			(_nw1).ArraySet1(false, 15)
			(_nw1).ArraySet1(p0, 16)
			(_nw1).ArraySet1(_8_v4.F24, 17)
			(_nw1).ArraySet1(_8_v4.F24, 18)
			(_nw1).ArraySet1(p0, 19)
			(_nw1).ArraySet1(p0, 20)
			(_nw1).ArraySet1(p1, 21)
			_11_v7 = _nw1
			var _12_v8 _dafny.Map
			_ = _12_v8
			_12_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v7, _11_v7)
			_12_v8 = (_12_v8).Update(_11_v7, _11_v7)
			var _13_v9 _dafny.Sequence
			_ = _13_v9
			_13_v9 = _dafny.SeqOf(p1, p0)
			var _14_v10 _dafny.Sequence
			_ = _14_v10
			_14_v10 = _dafny.SeqOf(_dafny.IntOfInt64(-330), _dafny.IntOfUint32((_13_v9).Cardinality()))
			var _15_v11 _dafny.Set
			_ = _15_v11
			_15_v11 = _dafny.SetOf(_14_v10, _dafny.SeqOf((_dafny.Zero).Minus(_4_v0)), _dafny.SeqOf(_4_v0), _14_v10, _dafny.SeqOf(_4_v0, _4_v0, _4_v0))
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_11_v7), 0))
			_ = _index0
			(_11_v7).ArraySet1((_15_v11).IsProperSubsetOf(_15_v11), (_index0).Int())
			var _16_v12 _dafny.Sequence
			_ = _16_v12
			_16_v12 = _dafny.UnicodeSeqOfUtf8Bytes("qlttakuyp")
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_11_v7), 0))
			_ = _index1
			var _rhs0 bool = p1
			_ = _rhs0
			var _rhs1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_16_v12, _16_v12), _dafny.UnicodeSeqOfUtf8Bytes("tfdjcpnpq"))
			_ = _rhs1
			var _rhs2 _dafny.Int = (_dafny.IntOfInt64(575)).Times(_4_v0)
			_ = _rhs2
			var _rhs3 *C0 = _8_v4
			_ = _rhs3
			var _lhs0 _dafny.Array = _11_v7
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_11_v7), 0))
			_ = _lhs1
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			_16_v12 = _rhs1
			_4_v0 = _rhs2
			_8_v4 = _rhs3
			_8_v4 = (func() *C0 {
				if (_8_v4.F24) || ((_11_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_11_v7), 0))).Int()).(bool)) {
					return _8_v4
				}
				return _8_v4
			})()
		} else {
			var _17_v13 _dafny.Sequence
			_ = _17_v13
			_17_v13 = _dafny.SeqOf(p1)
			var _18_v14 _dafny.Map
			_ = _18_v14
			_18_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_4_v0, _17_v13)
			var _19_v15 *C0
			_ = _19_v15
			var _nw2 *C0 = New_C0_()
			_ = _nw2
			_nw2.Ctor__(!(p1) || (p0), _18_v14)
			_19_v15 = _nw2
			var _20_v16 _dafny.Sequence
			_ = _20_v16
			_20_v16 = _dafny.UnicodeSeqOfUtf8Bytes("rqabmwu")
			var _21_v17 _dafny.CodePoint
			_ = _21_v17
			_21_v17 = _dafny.CodePoint('u')
			var _22_v18 _dafny.MultiSet
			_ = _22_v18
			_22_v18 = _dafny.MultiSetOf(p0, p1)
			var _23_v19 _dafny.Map
			_ = _23_v19
			_23_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_v17, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _4_v0)).Update(Companion_Default___.Fm11(_22_v18, _dafny.IntOfInt64(870), globalState), _4_v0)).Cardinality())
			_17_v13 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10(_dafny.IntOfUint32((_20_v16).Cardinality()), _dafny.IntOfInt64(751), globalState), _dafny.Companion_Sequence_.Concatenate(_17_v13, _17_v13)), (Companion_Default___.SafeIndex((_23_v19).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10(_dafny.IntOfUint32((_20_v16).Cardinality()), _dafny.IntOfInt64(751), globalState), _dafny.Companion_Sequence_.Concatenate(_17_v13, _17_v13))).Cardinality()))).Uint32(), !(p0))
			_4_v0 = _dafny.IntOfInt64(825)
			var _24_v20 _dafny.Array
			_ = _24_v20
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw3
			_24_v20 = _nw3
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_24_v20), 0))
			_ = _index2
			(_24_v20).ArraySet1(_4_v0, (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_24_v20), 0))
			_ = _index3
			(_24_v20).ArraySet1(((func() _dafny.Int {
				if p1 {
					return _4_v0
				}
				return _4_v0
			})()).Times(_4_v0), (_index3).Int())
			(globalState).F19 = (_17_v13).Select((Companion_Default___.SafeIndex(_4_v0, _dafny.IntOfUint32((_17_v13).Cardinality()))).Uint32()).(bool)
		}
		(globalState).F6 = p1
		(globalState).F6 = p1
		var _25_v21 _dafny.Map
		_ = _25_v21
		_25_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _26_v22 _dafny.Map
		_ = _26_v22
		_26_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_4_v0, (func() bool {
			if (_25_v21).Contains(false) {
				return (_25_v21).Get(false).(bool)
			}
			return p0
		})())
		var _27_v23 _dafny.Sequence
		_ = _27_v23
		_27_v23 = _dafny.SeqOf(p0, p1, p0, p0)
		(globalState).F19 = (func() bool {
			if !((func() bool {
				if (_26_v22).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(_4_v0))) {
					return (_26_v22).Get((_dafny.Zero).Minus((_dafny.Zero).Minus(_4_v0))).(bool)
				}
				return false
			})()) {
				return (_27_v23).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(p0, p1, globalState), _dafny.IntOfUint32((_27_v23).Cardinality()))).Uint32()).(bool)
			}
			return p1
		})()
	} else {
		var _28_v24 _dafny.Int
		_ = _28_v24
		_28_v24 = _dafny.IntOfInt64(191)
		var _29_v25 _dafny.Sequence
		_ = _29_v25
		_29_v25 = _dafny.UnicodeSeqOfUtf8Bytes("bvmshmd")
		_28_v24 = (_dafny.IntOfUint32((_29_v25).Cardinality())).Plus(_28_v24)
		var _30_v26 _dafny.Array
		_ = _30_v26
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
		_ = _nw4
		_30_v26 = _nw4
		_30_v26 = _30_v26
		var _31_v27 _dafny.Sequence
		_ = _31_v27
		_31_v27 = _dafny.SeqOf(false, p0)
		var _32_v28 D0
		_ = _32_v28
		_32_v28 = Companion_D0_.Create_DC1_(_31_v27, _dafny.IntOfInt64(-405))
		var _33_v29 _dafny.Map
		_ = _33_v29
		_33_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v24, (func() D0 {
			if p0 {
				return _32_v28
			}
			return _32_v28
		})())
		var _34_v30 _dafny.Map
		_ = _34_v30
		_34_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v24, p1)
		var _35_v31 _dafny.MultiSet
		_ = _35_v31
		_35_v31 = _dafny.MultiSetOf(Companion_Default___.Fm0(p0, p0, globalState), (_34_v30).Cardinality())
		var _36_v32 _dafny.Set
		_ = _36_v32
		_36_v32 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()))
		_33_v29 = (_33_v29).Update((func() _dafny.Int {
			if (_35_v31).Contains(_28_v24) {
				return (_35_v31).Multiplicity(_28_v24)
			}
			return (_36_v32).Cardinality()
		})(), _32_v28)
		var _37_v33 _dafny.Array
		_ = _37_v33
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
		_ = _nw5
		_37_v33 = _nw5
		_37_v33 = _37_v33
		var _38_v34 _dafny.Array
		_ = _38_v34
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw6
		_38_v34 = _nw6
		var _39_v35 _dafny.MultiSet
		_ = _39_v35
		_39_v35 = _dafny.MultiSetOf(_38_v34, _38_v34)
		var _40_v36 _dafny.MultiSet
		_ = _40_v36
		_40_v36 = (func() _dafny.MultiSet {
			if !(p1) {
				return _39_v35
			}
			return _39_v35
		})()
		var _source0 _dafny.MultiSet = _40_v36
		_ = _source0
		var _41___mcc_h0 _dafny.MultiSet = _source0
		_ = _41___mcc_h0
		var _42_cf18 _dafny.MultiSet = _41___mcc_h0
		_ = _42_cf18
		var _43_v37 _dafny.Map
		_ = _43_v37
		_43_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-759))
		var _44_v38 _dafny.Sequence
		_ = _44_v38
		_44_v38 = _dafny.SeqOf(_38_v34, _38_v34)
		var _45_v39 D3
		_ = _45_v39
		_45_v39 = Companion_D3_.Create_DC7_(_43_v37, p0, p1, p1, (_44_v38).Select((Companion_Default___.SafeIndex((_32_v28).Dtor_cf2(), _dafny.IntOfUint32((_44_v38).Cardinality()))).Uint32()).(_dafny.Array))
		var _46_v40 *C1
		_ = _46_v40
		var _nw7 *C1 = New_C1_()
		_ = _nw7
		_nw7.Ctor__(p1, (_45_v39).Dtor_cf14())
		_46_v40 = _nw7
		var _47_v41 _dafny.CodePoint
		_ = _47_v41
		_47_v41 = _dafny.CodePoint('t')
		(globalState).F18 = _47_v41
		(globalState).F19 = !((_46_v40).Fm4(_dafny.IntOfInt64(373), _dafny.IntOfInt64(230), _46_v40.F23, globalState)) || ((func() bool {
			if p1 {
				return p0
			}
			return (_45_v39).Dtor_cf12()
		})())
		(globalState).F19 = (_46_v40).F22()
	}
	var _48_v42 _dafny.Array
	_ = _48_v42
	var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
	_ = _nw8
	_48_v42 = _nw8
	for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_48_v42), 0))); ; {
		_guard_loop_0, _ok1 := _iter1()
		if !_ok1 {
			break
		}
		var _49_i0 _dafny.Int
		_49_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_49_i0).Sign() != -1) && ((_49_i0).Cmp(_dafny.ArrayLen((_48_v42), 0)) < 0)) {
			(_48_v42).ArraySet1(Companion_Default___.SafeDivisionInt(_49_i0, _dafny.IntOfInt64(-822)), (_49_i0).Int())
		}
	}
	var _50_i1 _dafny.Int
	_ = _50_i1
	_50_i1 = _dafny.Zero
	{
		for !(p0) {
			{
				if (_50_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_50_i1 = (_50_i1).Plus(_dafny.One)
				var _51_v43 _dafny.Int
				_ = _51_v43
				_51_v43 = _dafny.IntOfInt64(493)
				_51_v43 = _51_v43
				var _52_v44 _dafny.MultiSet
				_ = _52_v44
				_52_v44 = _dafny.MultiSetOf(p1)
				var _53_v45 *C1
				_ = _53_v45
				var _nw9 *C1 = New_C1_()
				_ = _nw9
				_nw9.Ctor__(true, p1)
				_53_v45 = _nw9
				var _54_v46 _dafny.MultiSet
				_ = _54_v46
				_54_v46 = _dafny.MultiSetOf(_53_v45)
				var _55_v47 D0
				_ = _55_v47
				_55_v47 = Companion_D0_.Create_DC2_(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), (Companion_D0_.Create_DC2_(Companion_Default___.Fm11(_52_v44, _51_v43, globalState), _51_v43, _51_v43, true)).Dtor_cf3())).Cardinality(), (_54_v46).Cardinality(), p0)
				var _56_v48 _dafny.Set
				_ = _56_v48
				_56_v48 = _dafny.SetOf((_55_v47).Dtor_cf3(), p1, (_53_v45).F22(), _53_v45.F23)
				var _57_v49 _dafny.Map
				_ = _57_v49
				_57_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v45.F23, p0)
				var _58_v50 _dafny.Map
				_ = _58_v50
				_58_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_56_v48).Cardinality(), (func() bool {
					if (_57_v49).Contains(_53_v45.F23) {
						return (_57_v49).Get(_53_v45.F23).(bool)
					}
					return _53_v45.F23
				})())
				var _59_v51 _dafny.Map
				_ = _59_v51
				_59_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v50, p0)
				_59_v51 = (_59_v51).Update(_58_v50, (_53_v45).F22())
				var _60_v52 _dafny.CodePoint
				_ = _60_v52
				_60_v52 = _dafny.CodePoint('q')
				var _61_v53 _dafny.Sequence
				_ = _61_v53
				_61_v53 = _dafny.SeqOf(_60_v52, _60_v52)
				var _62_v54 _dafny.Map
				_ = _62_v54
				_62_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v45.F23, _61_v53)
				var _63_v55 _dafny.Array
				_ = _63_v55
				var _nwElement0_1 _dafny.Sequence = Companion_Default___.Fm2(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if (_62_v54).Contains(_53_v45.F23) {
						return (_62_v54).Get(_53_v45.F23).(_dafny.Sequence)
					}
					return _dafny.SeqOf(_60_v52)
				})()), globalState)
				_ = _nwElement0_1
				var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
				_ = _nw10
				(_nw10).ArraySet1(_nwElement0_1, 0)
				_63_v55 = _nw10
				var _64_v56 _dafny.Set
				_ = _64_v56
				_64_v56 = _dafny.SetOf(Companion_Default___.Fm5(globalState))
				var _65_v57 _dafny.Sequence
				_ = _65_v57
				_65_v57 = _dafny.SeqOf(_64_v56, _64_v56, Companion_Default___.Fm13(_dafny.IntOfUint32((_61_v53).Cardinality()), _51_v43, globalState))
				var _66_v58 _dafny.Sequence
				_ = _66_v58
				_66_v58 = _dafny.SeqOf(_dafny.IntOfUint32((_65_v57).Cardinality()))
				var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_63_v55), 0))
				_ = _index4
				(_63_v55).ArraySet1(_66_v58, (_index4).Int())
				var _67_v59 _dafny.MultiSet
				_ = _67_v59
				_67_v59 = _dafny.MultiSetOf(_dafny.CodePoint('m'))
				var _68_v60 _dafny.Sequence
				_ = _68_v60
				_68_v60 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_66_v58, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.IntOfUint32((_66_v58).Cardinality()))).Uint32(), _51_v43), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}(func(_69_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(865)
				})))
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_63_v55), 0))
				_ = _index5
				(_63_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_67_v59, globalState), (_68_v60).Select((Companion_Default___.SafeIndex(_51_v43, _dafny.IntOfUint32((_68_v60).Cardinality()))).Uint32()).(_dafny.Sequence)), (_index5).Int())
				if (_53_v45).F22() {
					var _70_v61 _dafny.Map
					_ = _70_v61
					_70_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _51_v43)
					var _71_v62 D3
					_ = _71_v62
					_71_v62 = Companion_D3_.Create_DC8_((func() _dafny.Int {
						if (_70_v61).Contains((_53_v45).F22()) {
							return (_70_v61).Get((_53_v45).F22()).(_dafny.Int)
						}
						return _51_v43
					})(), (func() _dafny.Int {
						if (_53_v45).F22() {
							return _51_v43
						}
						return _51_v43
					})())
					var _pat_let_tv0 = _51_v43
					_ = _pat_let_tv0
					_71_v62 = (func() D3 {
						if (_53_v45).F22() {
							return _71_v62
						}
						return func(_pat_let0_0 D3) D3 {
							return func(_72_dt__update__tmp_h0 D3) D3 {
								return func(_pat_let1_0 _dafny.Int) D3 {
									return func(_73_dt__update_hcf17_h0 _dafny.Int) D3 {
										return Companion_D3_.Create_DC8_((_72_dt__update__tmp_h0).Dtor_cf16(), _73_dt__update_hcf17_h0)
									}(_pat_let1_0)
								}(_pat_let_tv0)
							}(_pat_let0_0)
						}(_71_v62)
					})()
					var _74_v64 T0
					_ = _74_v64
					var _nw11 *C0 = New_C0_()
					_ = _nw11
					_nw11.Ctor__(p1, func() _dafny.Map {
						var _coll1 = _dafny.NewMapBuilder()
						_ = _coll1
						for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-478), _dafny.IntOfInt64(307))); ; {
							_compr_1, _ok2 := _iter2()
							if !_ok2 {
								break
							}
							var _75_v63 _dafny.Int
							_75_v63 = interface{}(_compr_1).(_dafny.Int)
							if ((_dafny.IntOfInt64(-478)).Cmp(_75_v63) <= 0) && ((_75_v63).Cmp(_dafny.IntOfInt64(307)) < 0) {
								_coll1.Add(Companion_Default___.SafeModuloInt(_75_v63, _51_v43), _dafny.SeqOf(false))
							}
						}
						return _coll1.ToMap()
					}())
					_74_v64 = _nw11
					var _76_v65 _dafny.Array
					_ = _76_v65
					var _nwElement0_2 T0 = _74_v64
					_ = _nwElement0_2
					var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
					_ = _nw12
					(_nw12).ArraySet1(_nwElement0_2, 0)
					(_nw12).ArraySet1(_74_v64, 1)
					(_nw12).ArraySet1(_74_v64, 2)
					(_nw12).ArraySet1(_74_v64, 3)
					_76_v65 = _nw12
					_76_v65 = _76_v65
					var _77_v66 *C1
					_ = _77_v66
					var _nw13 *C1 = New_C1_()
					_ = _nw13
					_nw13.Ctor__(!(false), p1)
					_77_v66 = _nw13
					var _rhs4 _dafny.CodePoint = _60_v52
					_ = _rhs4
					var _rhs5 _dafny.Int = _51_v43
					_ = _rhs5
					var _lhs2 *GlobalState = globalState
					_ = _lhs2
					_lhs2.F18 = _rhs4
					_51_v43 = _rhs5
					var _78_v67 _dafny.Array
					_ = _78_v67
					var _nw14 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
					_ = _nw14
					_78_v67 = _nw14
					_78_v67 = _78_v67
				} else {
					_51_v43 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_61_v53, _dafny.UnicodeSeqOfUtf8Bytes("q"))).Cardinality())
					var _79_v68 _dafny.Sequence
					_ = _79_v68
					_79_v68 = _dafny.SeqOf(p0, p0)
					var _80_v69 _dafny.Map
					_ = _80_v69
					_80_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v68, _53_v45.F23)
					_51_v43 = (func() _dafny.Int {
						if _53_v45.F23 {
							return Companion_Default___.SafeModuloInt(_51_v43, _51_v43)
						}
						return (_80_v69).Cardinality()
					})()
					var _81_v70 _dafny.Array
					_ = _81_v70
					var _nw15 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
					_ = _nw15
					_81_v70 = _nw15
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_81_v70), 0))
					_ = _index6
					(_81_v70).ArraySet1(p0, (_index6).Int())
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_81_v70), 0))
					_ = _index7
					(_81_v70).ArraySet1(((_dafny.IntOfInt64(-514)).Cmp(_dafny.IntOfInt64(861)) <= 0) || ((_53_v45).F22()), (_index7).Int())
					var _82_v71 _dafny.Array
					_ = _82_v71
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(22)
					_ = _len0_0
					var _nw16 _dafny.Array
					_ = _nw16
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw16 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.CodePoint = (func(_83_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_84_i3 _dafny.Int) _dafny.CodePoint {
								return _83_v52
							}
						})(_60_v52)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw16 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw16).ArraySet1CodePoint(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw16).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_82_v71 = _nw16
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_82_v71), 0))
					_ = _index8
					(_82_v71).ArraySet1CodePoint(_60_v52, (_index8).Int())
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.ArrayLen((_82_v71), 0))
					_ = _index9
					(_82_v71).ArraySet1CodePoint(_60_v52, (_index9).Int())
					var _85_v72 _dafny.Array
					_ = _85_v72
					var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
					_ = _nw17
					_85_v72 = _nw17
					var _86_v73 _dafny.Array
					_ = _86_v73
					var _nwElement0_3 _dafny.Array = _85_v72
					_ = _nwElement0_3
					var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.One)
					_ = _nw18
					(_nw18).ArraySet1(_nwElement0_3, 0)
					_86_v73 = _nw18
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_86_v73), 0))
					_ = _index10
					(_86_v73).ArraySet1(_85_v72, (_index10).Int())
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_86_v73), 0))
					_ = _index11
					var _rhs6 _dafny.Array = _85_v72
					_ = _rhs6
					var _rhs7 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
						if p0 {
							return _79_v68
						}
						return _dafny.SeqOf(p0, (_53_v45).F22(), (_53_v45).F22())
					})(), _79_v68)).Cardinality()))
					_ = _rhs7
					var _rhs8 bool = _dafny.Companion_Sequence_.Equal((_63_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_63_v55), 0))).Int()).(_dafny.Sequence), _66_v58)
					_ = _rhs8
					var _lhs3 _dafny.Array = _86_v73
					_ = _lhs3
					var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_86_v73), 0))
					_ = _lhs4
					var _lhs5 *C1 = _53_v45
					_ = _lhs5
					(_lhs3).ArraySet1(_rhs6, (_lhs4).Int())
					_51_v43 = _rhs7
					_lhs5.F23 = _rhs8
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _87_i4 _dafny.Int
	_ = _87_i4
	_87_i4 = _dafny.Zero
	{
		for p0 {
			{
				if (_87_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_87_i4 = (_87_i4).Plus(_dafny.One)
				var _88_v74 _dafny.Int
				_ = _88_v74
				_88_v74 = _dafny.IntOfInt64(120)
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))
				_ = _index12
				(_48_v42).ArraySet1(_88_v74, (_index12).Int())
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))
				_ = _index13
				(_48_v42).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_88_v74, _dafny.IntOfInt64(-138))), (_index13).Int())
				var _89_v75 _dafny.Sequence
				_ = _89_v75
				_89_v75 = _dafny.UnicodeSeqOfUtf8Bytes("nqxh")
				var _90_v76 _dafny.Map
				_ = _90_v76
				_90_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _89_v75)
				var _91_v77 _dafny.Map
				_ = _91_v77
				_91_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v74, (_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int))
				var _92_v78 _dafny.CodePoint
				_ = _92_v78
				_92_v78 = _dafny.CodePoint('h')
				var _93_v79 D5
				_ = _93_v79
				_93_v79 = Companion_D5_.Create_DC10_(_92_v78)
				var _94_v80 _dafny.Sequence
				_ = _94_v80
				_94_v80 = _dafny.SeqOf(p0)
				var _95_v81 *C0
				_ = _95_v81
				var _nw19 *C0 = New_C0_()
				_ = _nw19
				_nw19.Ctor__(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), _94_v80))
				_95_v81 = _nw19
				var _96_v82 _dafny.Sequence
				_ = _96_v82
				_96_v82 = _dafny.SeqOf(_95_v81)
				var _97_v83 _dafny.Array
				_ = _97_v83
				var _nwElement0_4 _dafny.Int = _dafny.IntOfInt64(-478)
				_ = _nwElement0_4
				var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(24))
				_ = _nw20
				(_nw20).ArraySet1(_nwElement0_4, 0)
				(_nw20).ArraySet1((_91_v77).Cardinality(), 1)
				(_nw20).ArraySet1(_88_v74, 2)
				(_nw20).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_93_v79).Dtor_cf19())).Cardinality(), 3)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 4)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 5)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 6)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 7)
				(_nw20).ArraySet1(_88_v74, 8)
				(_nw20).ArraySet1(_88_v74, 9)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 10)
				(_nw20).ArraySet1(_dafny.IntOfUint32((_96_v82).Cardinality()), 11)
				(_nw20).ArraySet1(_88_v74, 12)
				(_nw20).ArraySet1(_dafny.IntOfUint32((_89_v75).Cardinality()), 13)
				(_nw20).ArraySet1(_88_v74, 14)
				(_nw20).ArraySet1(_88_v74, 15)
				(_nw20).ArraySet1(_88_v74, 16)
				(_nw20).ArraySet1(_88_v74, 17)
				(_nw20).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), p1)).Cardinality(), 18)
				(_nw20).ArraySet1(Companion_Default___.Fm0(_95_v81.F24, p0, globalState), 19)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 20)
				(_nw20).ArraySet1(_88_v74, 21)
				(_nw20).ArraySet1(_88_v74, 22)
				(_nw20).ArraySet1((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), 23)
				_97_v83 = _nw20
				var _98_v84 _dafny.Map
				_ = _98_v84
				_98_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("no"), _97_v83)
				var _99_v85 _dafny.Set
				_ = _99_v85
				_99_v85 = _dafny.SetOf(_88_v74, (_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), (_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), (_98_v84).Cardinality(), _dafny.IntOfUint32((_94_v80).Cardinality()))
				var _100_v86 _dafny.Sequence
				_ = _100_v86
				_100_v86 = _dafny.SeqOf((_99_v85).Cardinality(), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}((func(_101_v75 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_102_i5 _dafny.Int) _dafny.CodePoint {
						return (_101_v75).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.IntOfUint32((_101_v75).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_89_v75)))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_89_v75).Cardinality())))).Cardinality())
				var _103_v87 T0
				_ = _103_v87
				var _nw21 *C0 = New_C0_()
				_ = _nw21
				_nw21.Ctor__(_95_v81.F24, (_95_v81).F25())
				_103_v87 = _nw21
				var _104_v88 _dafny.Map
				_ = _104_v88
				_104_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int), _103_v87)
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))
				_ = _index14
				(_48_v42).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_90_v76).Contains(p0) {
						return (_90_v76).Get(p0).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("ovddoibvv")
				})()).Cardinality())).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_100_v86, (Companion_Default___.SafeIndex((_104_v88).Cardinality(), _dafny.IntOfUint32((_100_v86).Cardinality()))).Uint32(), _88_v74)).Cardinality())), (_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int)), (_index14).Int())
				var _105_v89 _dafny.Array
				_ = _105_v89
				var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw22
				_105_v89 = _nw22
				var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_105_v89), 0))
				_ = _index15
				(_105_v89).ArraySet1(!(p1), (_index15).Int())
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_105_v89), 0))
				_ = _index16
				(_105_v89).ArraySet1(_95_v81.F24, (_index16).Int())
				var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))
				_ = _index17
				var _rhs9 T0 = _103_v87
				_ = _rhs9
				var _rhs10 bool = _95_v81.F24
				_ = _rhs10
				var _rhs11 _dafny.Int = (_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int)
				_ = _rhs11
				var _rhs12 _dafny.Int = ((_dafny.Zero).Minus(_88_v74)).Plus((_48_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))).Int()).(_dafny.Int))
				_ = _rhs12
				var _lhs6 *GlobalState = globalState
				_ = _lhs6
				var _lhs7 _dafny.Array = _48_v42
				_ = _lhs7
				var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_48_v42), 0))
				_ = _lhs8
				_103_v87 = _rhs9
				_lhs6.F19 = _rhs10
				(_lhs7).ArraySet1(_rhs11, (_lhs8).Int())
				_88_v74 = _rhs12
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _106_v90 _dafny.Int
	_ = _106_v90
	_106_v90 = _dafny.IntOfInt64(568)
	_106_v90 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-273), _106_v90)
	var _107_v91 _dafny.Sequence
	_ = _107_v91
	_107_v91 = _dafny.UnicodeSeqOfUtf8Bytes("obf")
	var _108_v92 _dafny.Array
	_ = _108_v92
	var _nw23 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
	_ = _nw23
	_108_v92 = _nw23
	var _109_v93 _dafny.MultiSet
	_ = _109_v93
	_109_v93 = _dafny.MultiSetOf(p1)
	var _110_v94 _dafny.Sequence
	_ = _110_v94
	_110_v94 = _dafny.SeqOf(Companion_Default___.Fm11(_109_v93, _106_v90, globalState))
	var _111_v95 T0
	_ = _111_v95
	var _nw24 *C0 = New_C0_()
	_ = _nw24
	_nw24.Ctor__(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v90, _110_v94))
	_111_v95 = _nw24
	var _112_v96 D5
	_ = _112_v96
	_112_v96 = Companion_D5_.Create_DC11_(p1, p0, _108_v92, _111_v95)
	var _113_v97 _dafny.Sequence
	_ = _113_v97
	_113_v97 = _dafny.SeqOf(_112_v96, Companion_D5_.Create_DC11_(p1, p0, _108_v92, _111_v95), _112_v96, Companion_D5_.Create_DC11_(p0, p1, _108_v92, _111_v95))
	var _114_v98 _dafny.Map
	_ = _114_v98
	_114_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v91, _113_v97)
	var _115_v99 _dafny.Array
	_ = _115_v99
	var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
	_ = _nw25
	_115_v99 = _nw25
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_115_v99), 0))
	_ = _index18
	(_115_v99).ArraySet1(_48_v42, (_index18).Int())
	var _116_v100 _dafny.Array
	_ = _116_v100
	_116_v100 = _48_v42
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_115_v99), 0))
	_ = _index19
	var _rhs13 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_117_i6 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	})), _113_v97)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rntowef"), _113_v97))
	_ = _rhs13
	var _rhs14 _dafny.Array = (_116_v100)
	_ = _rhs14
	var _lhs9 _dafny.Array = _115_v99
	_ = _lhs9
	var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(278), _dafny.ArrayLen((_115_v99), 0))
	_ = _lhs10
	_114_v98 = _rhs13
	(_lhs9).ArraySet1(_rhs14, (_lhs10).Int())
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _118_v0 bool
	_ = _118_v0
	_118_v0 = true
	var _119_v1 _dafny.Sequence
	_ = _119_v1
	_119_v1 = _dafny.SeqOf(_118_v0)
	var _120_v2 _dafny.Sequence
	_ = _120_v2
	_120_v2 = _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg")
	var _121_v3 D0
	_ = _121_v3
	_121_v3 = Companion_D0_.Create_DC1_(_119_v1, _dafny.IntOfUint32((_120_v2).Cardinality()))
	var _122_v4 _dafny.Map
	_ = _122_v4
	_122_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v0, (_121_v3).Dtor_cf2())
	var _123_v5 _dafny.Map
	_ = _123_v5
	_123_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v3, _dafny.MultiSetOf(_118_v0))
	var _124_v6 _dafny.Set
	_ = _124_v6
	_124_v6 = _dafny.SetOf(_118_v0)
	var _125_v7 _dafny.MultiSet
	_ = _125_v7
	_125_v7 = _dafny.MultiSetOf(_118_v0)
	var _126_v8 _dafny.Sequence
	_ = _126_v8
	_126_v8 = _dafny.SeqOf(_120_v2, _120_v2, _120_v2)
	var _127_v9 _dafny.Array
	_ = _127_v9
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_1
	var _nw26 _dafny.Array
	_ = _nw26
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw26 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = (func(_128_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_129_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_129_i0, _dafny.IntOfUint32((_128_v2).Cardinality()))
			}
		})(_120_v2)
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
	_127_v9 = _nw26
	var _130_v11 _dafny.CodePoint
	_ = _130_v11
	_130_v11 = _dafny.CodePoint('j')
	var _131_v12 _dafny.Set
	_ = _131_v12
	_131_v12 = _dafny.SetOf(_130_v11)
	var _132_v13 _dafny.Map
	_ = _132_v13
	_132_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_125_v7, _118_v0)
	var _133_v14 _dafny.Int
	_ = _133_v14
	_133_v14 = _dafny.IntOfInt64(588)
	var _134_v15 _dafny.Map
	_ = _134_v15
	_134_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v11, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v0, _133_v14)).Update(_118_v0, _133_v14))
	var _135_globalState *GlobalState
	_ = _135_globalState
	var _nw27 *GlobalState = New_GlobalState_()
	_ = _nw27
	_nw27.Ctor__(_122_v4, true, true, (func() _dafny.MultiSet {
		if (_123_v5).Contains(Companion_D0_.Create_DC1_(_119_v1, (_124_v6).Cardinality())) {
			return (_123_v5).Get(Companion_D0_.Create_DC1_(_119_v1, (_124_v6).Cardinality())).(_dafny.MultiSet)
		}
		return _125_v7
	})(), true, _dafny.IntOfInt64(869), false, _dafny.IntOfInt64(924), false, _dafny.IntOfInt64(-924), _dafny.Companion_Sequence_.Concatenate(_126_v8, _126_v8), _120_v2, _dafny.IntOfInt64(380), _127_v9, _dafny.IntOfInt64(486), func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(729), _dafny.IntOfInt64(212))); ; {
			_compr_2, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _136_v10 _dafny.Int
			_136_v10 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(729)).Cmp(_136_v10) <= 0) && ((_136_v10).Cmp(_dafny.IntOfInt64(212)) < 0) {
				_coll2.Add(Companion_Default___.SafeDivisionInt(_136_v10, _dafny.IntOfInt64(-701)), _dafny.SeqOf(_dafny.IntOfInt64(-6), (_dafny.Zero).Minus(_dafny.IntOfInt64(-208))))
			}
		}
		return _coll2.ToMap()
	}(), _131_v12, (_132_v13).Merge(_132_v13), _dafny.CodePoint('y'), false, _134_v15, _dafny.IntOfInt64(964))
	_135_globalState = _nw27
	var _137_v16 _dafny.Set
	_ = _137_v16
	_137_v16 = _dafny.SetOf(_119_v1, _119_v1, _dafny.Companion_Sequence_.Concatenate(_119_v1, _119_v1))
	_137_v16 = (_137_v16).Union((_137_v16).Union(_137_v16))
	var _138_v17 _dafny.Map
	_ = _138_v17
	_138_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if _118_v0 {
			return _118_v0
		}
		return !(_118_v0)
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(243))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}((func(_139_v14 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_140_i1 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_139_v14)
		}
	})(_133_v14))))
	var _141_v18 _dafny.Sequence
	_ = _141_v18
	_141_v18 = _dafny.SeqOf(_125_v7)
	var _142_v19 _dafny.Sequence
	_ = _142_v19
	_142_v19 = _dafny.SeqOf(_133_v14, ((_141_v18).Select((Companion_Default___.SafeIndex(_133_v14, _dafny.IntOfUint32((_141_v18).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality())
	_138_v17 = (_138_v17).Update((_118_v0) || (!(_118_v0)), _142_v19)
	var _143_v20 _dafny.Array
	_ = _143_v20
	var _nw28 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
	_ = _nw28
	_143_v20 = _nw28
	_143_v20 = _143_v20
	var _144_v21 D0
	_ = _144_v21
	_144_v21 = Companion_D0_.Create_DC2_(_118_v0, _133_v14, (_122_v4).Cardinality(), _118_v0)
	var _145_v22 _dafny.Map
	_ = _145_v22
	_145_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v14, Companion_D0_.Create_DC3_(_144_v21))
	var _146_v23 _dafny.Array
	_ = _146_v23
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
	_ = _nw29
	_146_v23 = _nw29
	var _147_v24 _dafny.Map
	_ = _147_v24
	_147_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_130_v11), Companion_Default___.Fm0(_118_v0, _118_v0, _135_globalState))
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_146_v23), 0))
	_ = _index20
	(_146_v23).ArraySet1(_147_v24, (_index20).Int())
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
	_ = _index21
	(_127_v9).ArraySet1(_133_v14, (_index21).Int())
	var _148_v25 _dafny.Set
	_ = _148_v25
	_148_v25 = _dafny.SetOf(_dafny.IntOfUint32((_120_v2).Cardinality()))
	var _149_v26 _dafny.MultiSet
	_ = _149_v26
	_149_v26 = _dafny.MultiSetOf(_130_v11)
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_146_v23), 0))
	_ = _index22
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
	_ = _index23
	var _rhs15 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_120_v2, _120_v2)
	_ = _rhs15
	var _rhs16 _dafny.Map = Companion_Default___.Fm1(_118_v0, _118_v0, Companion_Default___.SafeModuloInt((_148_v25).Cardinality(), _133_v14), _133_v14, _135_globalState)
	_ = _rhs16
	var _rhs17 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_149_v26, (_133_v14).Plus(_133_v14))
	_ = _rhs17
	var _rhs18 _dafny.Set = _124_v6
	_ = _rhs18
	var _rhs19 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt(_133_v14, _133_v14)).Minus(_133_v14))
	_ = _rhs19
	var _lhs11 _dafny.Array = _146_v23
	_ = _lhs11
	var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.ArrayLen((_146_v23), 0))
	_ = _lhs12
	var _lhs13 _dafny.Array = _127_v9
	_ = _lhs13
	var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
	_ = _lhs14
	_120_v2 = _rhs15
	_145_v22 = _rhs16
	(_lhs11).ArraySet1(_rhs17, (_lhs12).Int())
	_124_v6 = _rhs18
	(_lhs13).ArraySet1(_rhs19, (_lhs14).Int())
	_142_v19 = _142_v19
	_120_v2 = _dafny.Companion_Sequence_.Concatenate(_120_v2, _dafny.Companion_Sequence_.Update(_120_v2, (Companion_Default___.SafeIndex(_133_v14, _dafny.IntOfUint32((_120_v2).Cardinality()))).Uint32(), _130_v11))
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
	_ = _index24
	var _rhs20 D0 = _121_v3
	_ = _rhs20
	var _rhs21 _dafny.Int = (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)
	_ = _rhs21
	var _lhs15 _dafny.Array = _127_v9
	_ = _lhs15
	var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
	_ = _lhs16
	_121_v3 = _rhs20
	(_lhs15).ArraySet1(_rhs21, (_lhs16).Int())
	var _150_i2 _dafny.Int
	_ = _150_i2
	_150_i2 = _dafny.Zero
	{
		for !(_118_v0) {
			{
				if (_150_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_150_i2 = (_150_i2).Plus(_dafny.One)
				if !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(138))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}((func(_151_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_152_i3 _dafny.Int) _dafny.CodePoint {
						return _151_v11
					}
				})(_130_v11))), _dafny.Companion_Sequence_.Concatenate(_120_v2, _120_v2)) {
					_142_v19 = Companion_Default___.Fm2((_dafny.MultiSetFromSeq(_120_v2)).Intersection(_149_v26), _135_globalState)
					_120_v2 = Companion_Default___.Fm3(_133_v14, _135_globalState)
					var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_143_v20), 0))
					_ = _index25
					(_143_v20).ArraySet1(_118_v0, (_index25).Int())
					var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_143_v20), 0))
					_ = _index26
					(_143_v20).ArraySet1(_118_v0, (_index26).Int())
					var _153_v27 _dafny.Map
					_ = _153_v27
					_153_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_148_v25).Cardinality(), _119_v1)
					var _154_v28 *C0
					_ = _154_v28
					var _nw30 *C0 = New_C0_()
					_ = _nw30
					_nw30.Ctor__(!(!((_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool))), _153_v27)
					_154_v28 = _nw30
					Companion_Default___.M2((func() bool {
						if _154_v28.F24 {
							return (_119_v1).Select((Companion_Default___.SafeIndex(_133_v14, _dafny.IntOfUint32((_119_v1).Cardinality()))).Uint32()).(bool)
						}
						return _118_v0
					})(), ((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int))) < 0, _135_globalState)
				} else {
					_133_v14 = (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)
					(_135_globalState).F19 = ((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)).Cmp(_133_v14) != 0
					var _155_v29 _dafny.Map
					_ = _155_v29
					_155_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_142_v19).Cardinality()), _dafny.Companion_Sequence_.Update(_119_v1, (Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_122_v4).Contains(_118_v0) {
							return (_122_v4).Get(_118_v0).(_dafny.Int)
						}
						return _133_v14
					})(), _dafny.IntOfUint32((_119_v1).Cardinality()))).Uint32(), _118_v0))
					var _156_v30 T0
					_ = _156_v30
					var _nw31 *C0 = New_C0_()
					_ = _nw31
					_nw31.Ctor__(_118_v0, _155_v29)
					_156_v30 = _nw31
					var _157_v31 _dafny.Map
					_ = _157_v31
					_157_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v11, _156_v30)
					var _158_v32 _dafny.Map
					_ = _158_v32
					_158_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(759), _156_v30)
					var _159_v33 _dafny.Array
					_ = _159_v33
					var _nwElement0_5 T0 = _156_v30
					_ = _nwElement0_5
					var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(20))
					_ = _nw32
					(_nw32).ArraySet1(_nwElement0_5, 0)
					(_nw32).ArraySet1(_156_v30, 1)
					(_nw32).ArraySet1(_156_v30, 2)
					(_nw32).ArraySet1(_156_v30, 3)
					(_nw32).ArraySet1(_156_v30, 4)
					(_nw32).ArraySet1(_156_v30, 5)
					(_nw32).ArraySet1(_156_v30, 6)
					(_nw32).ArraySet1((func() T0 {
						if _118_v0 {
							return (func() T0 {
								if (_157_v31).Contains(_130_v11) {
									return (_157_v31).Get(_130_v11).(T0)
								}
								return _156_v30
							})()
						}
						return _156_v30
					})(), 7)
					(_nw32).ArraySet1(_156_v30, 8)
					(_nw32).ArraySet1(_156_v30, 9)
					(_nw32).ArraySet1(_156_v30, 10)
					(_nw32).ArraySet1(_156_v30, 11)
					(_nw32).ArraySet1(_156_v30, 12)
					(_nw32).ArraySet1((func() T0 {
						if !(_118_v0) {
							return _156_v30
						}
						return _156_v30
					})(), 13)
					(_nw32).ArraySet1(_156_v30, 14)
					(_nw32).ArraySet1(_156_v30, 15)
					(_nw32).ArraySet1(_156_v30, 16)
					(_nw32).ArraySet1((func() T0 {
						if (_158_v32).Contains((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)) {
							return (_158_v32).Get((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)).(T0)
						}
						return _156_v30
					})(), 17)
					(_nw32).ArraySet1(_156_v30, 18)
					(_nw32).ArraySet1(_156_v30, 19)
					_159_v33 = _nw32
					var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_159_v33), 0))
					_ = _index27
					(_159_v33).ArraySet1(_156_v30, (_index27).Int())
					var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_159_v33), 0))
					_ = _index28
					var _nw33 *C0 = New_C0_()
					_ = _nw33
					_nw33.Ctor__(!(_118_v0), (_155_v29).Update(_dafny.IntOfUint32((_119_v1).Cardinality()), _119_v1))
					(_159_v33).ArraySet1(_nw33, (_index28).Int())
					Companion_Default___.M2(_118_v0, !(Companion_Default___.Fm11(_125_v7, _dafny.IntOfUint32((_dafny.SeqOf((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int))).Cardinality()), _135_globalState)), _135_globalState)
					_133_v14 = (_156_v30).Fm8(_135_globalState)
				}
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))
				_ = _index29
				(_143_v20).ArraySet1(!(Companion_Default___.Fm11(_125_v7, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kvyhg")).Cardinality()), _135_globalState)), (_index29).Int())
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))
				_ = _index30
				(_143_v20).ArraySet1(Companion_Default___.Fm11((_dafny.MultiSetOf(Companion_Default___.Fm11(_dafny.MultiSetOf(_118_v0, _118_v0), (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _135_globalState), _118_v0)).Union(_125_v7), _133_v14, _135_globalState), (_index30).Int())
				if !(!((_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool))) {
					var _160_v34 D0
					_ = _160_v34
					_160_v34 = Companion_D0_.Create_DC2_(_118_v0, _133_v14, (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _118_v0)
					var _161_v35 _dafny.Map
					_ = _161_v35
					_161_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), (_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool))
					var _162_v36 _dafny.Array
					_ = _162_v36
					var _nwElement0_6 bool = (_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool)
					_ = _nwElement0_6
					var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(23))
					_ = _nw34
					(_nw34).ArraySet1(_nwElement0_6, 0)
					(_nw34).ArraySet1((_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool), 1)
					(_nw34).ArraySet1((_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool), 2)
					(_nw34).ArraySet1(_118_v0, 3)
					(_nw34).ArraySet1(false, 4)
					(_nw34).ArraySet1(true, 5)
					(_nw34).ArraySet1(_118_v0, 6)
					(_nw34).ArraySet1(_118_v0, 7)
					(_nw34).ArraySet1(_118_v0, 8)
					(_nw34).ArraySet1(_118_v0, 9)
					(_nw34).ArraySet1((_160_v34).Dtor_cf6(), 10)
					(_nw34).ArraySet1(true, 11)
					(_nw34).ArraySet1(_118_v0, 12)
					(_nw34).ArraySet1(true, 13)
					(_nw34).ArraySet1((_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool), 14)
					(_nw34).ArraySet1(_118_v0, 15)
					(_nw34).ArraySet1(_118_v0, 16)
					(_nw34).ArraySet1(true, 17)
					(_nw34).ArraySet1(!(false), 18)
					(_nw34).ArraySet1((_160_v34).Dtor_cf3(), 19)
					(_nw34).ArraySet1(!(_118_v0), 20)
					(_nw34).ArraySet1(_118_v0, 21)
					(_nw34).ArraySet1((func() bool {
						if (_161_v35).Contains(_dafny.IntOfInt64(293)) {
							return (_161_v35).Get(_dafny.IntOfInt64(293)).(bool)
						}
						return _118_v0
					})(), 22)
					_162_v36 = _nw34
					var _163_v37 _dafny.Map
					_ = _163_v37
					_163_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v36, false)
					_163_v37 = (_163_v37).Update(_162_v36, (func() bool {
						if (_161_v35).Contains(_dafny.IntOfInt64(788)) {
							return (_161_v35).Get(_dafny.IntOfInt64(788)).(bool)
						}
						return _118_v0
					})())
					var _164_v38 D3
					_ = _164_v38
					_164_v38 = Companion_D3_.Create_DC8_(_133_v14, _133_v14)
					var _pat_let_tv1 = _127_v9
					_ = _pat_let_tv1
					var _pat_let_tv2 = _127_v9
					_ = _pat_let_tv2
					_133_v14 = (func() _dafny.Int {
						if (_119_v1).Select((Companion_Default___.SafeIndex(_133_v14, _dafny.IntOfUint32((_119_v1).Cardinality()))).Uint32()).(bool) {
							return Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)), (func(_pat_let2_0 D3) D3 {
								return func(_165_dt__update__tmp_h0 D3) D3 {
									return func(_pat_let3_0 _dafny.Int) D3 {
										return func(_166_dt__update_hcf17_h0 _dafny.Int) D3 {
											return Companion_D3_.Create_DC8_((_165_dt__update__tmp_h0).Dtor_cf16(), _166_dt__update_hcf17_h0)
										}(_pat_let3_0)
									}((_pat_let_tv2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_pat_let_tv1), 0))).Int()).(_dafny.Int))
								}(_pat_let2_0)
							}(_164_v38)).Dtor_cf16())
						}
						return _dafny.IntOfInt64(803)
					})()
					(_135_globalState).F6 = ((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)).Cmp(_133_v14) < 0
					var _167_v39 _dafny.Sequence
					_ = _167_v39
					_167_v39 = _dafny.SeqOf(_162_v36)
					_162_v36 = (_167_v39).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_133_v14, (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_167_v39).Cardinality()))).Uint32()).(_dafny.Array)
					var _168_v40 *C1
					_ = _168_v40
					var _nw35 *C1 = New_C1_()
					_ = _nw35
					_nw35.Ctor__(Companion_Default___.Fm11(_125_v7, _dafny.IntOfUint32((_119_v1).Cardinality()), _135_globalState), _118_v0)
					_168_v40 = _nw35
				} else {
					Companion_Default___.M2((func() bool {
						if true {
							return (_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool)
						}
						return _118_v0
					})(), true, _135_globalState)
					(_135_globalState).F19 = _118_v0
					var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))
					_ = _index31
					(_143_v20).ArraySet1(Companion_Default___.Fm11(_dafny.MultiSetOf((_143_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))).Int()).(bool), _118_v0, _118_v0), _133_v14, _135_globalState), (_index31).Int())
					Companion_Default___.M2(_118_v0, _118_v0, _135_globalState)
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v20), 0))
					_ = _index32
					(_143_v20).ArraySet1(_118_v0, (_index32).Int())
				}
				_133_v14 = (Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v14, (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int))).Cardinality(), _dafny.IntOfUint32((_120_v2).Cardinality()))).Minus(((_124_v6).Union(_124_v6)).Cardinality())
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _169_v41 _dafny.Map
	_ = _169_v41
	_169_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _dafny.SeqOf(_118_v0))
	var _170_v42 T0
	_ = _170_v42
	var _nw36 *C0 = New_C0_()
	_ = _nw36
	_nw36.Ctor__(true, _169_v41)
	_170_v42 = _nw36
	var _171_v43 _dafny.Sequence
	_ = _171_v43
	_171_v43 = _dafny.SeqOf(_170_v42, _170_v42, _170_v42)
	(_135_globalState).F19 = _dafny.Companion_Sequence_.IsProperPrefixOf(_171_v43, _171_v43)
	Companion_Default___.M2(_118_v0, _118_v0, _135_globalState)
	var _172_v44 *C1
	_ = _172_v44
	var _nw37 *C1 = New_C1_()
	_ = _nw37
	_nw37.Ctor__((func() bool {
		if _118_v0 {
			return _118_v0
		}
		return _118_v0
	})(), _118_v0)
	_172_v44 = _nw37
	_172_v44 = _172_v44
	var _173_v45 _dafny.Array
	_ = _173_v45
	var _nwElement0_7 _dafny.Array = _127_v9
	_ = _nwElement0_7
	var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(11))
	_ = _nw38
	(_nw38).ArraySet1(_nwElement0_7, 0)
	(_nw38).ArraySet1(_127_v9, 1)
	(_nw38).ArraySet1(_127_v9, 2)
	(_nw38).ArraySet1(_127_v9, 3)
	(_nw38).ArraySet1(_127_v9, 4)
	(_nw38).ArraySet1(_127_v9, 5)
	(_nw38).ArraySet1(_127_v9, 6)
	(_nw38).ArraySet1(_127_v9, 7)
	(_nw38).ArraySet1(_127_v9, 8)
	(_nw38).ArraySet1(_127_v9, 9)
	(_nw38).ArraySet1(_127_v9, 10)
	_173_v45 = _nw38
	var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_173_v45), 0))
	_ = _index33
	(_173_v45).ArraySet1(_127_v9, (_index33).Int())
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_173_v45), 0))
	_ = _index34
	(_173_v45).ArraySet1(_127_v9, (_index34).Int())
	_133_v14 = (func() _dafny.Int {
		if (_dafny.IntOfUint32((Companion_Default___.Fm3((func() _dafny.Int {
			if (_122_v4).Contains((_172_v44).Fm4(_133_v14, (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _172_v44.F23, _135_globalState)) {
				return (_122_v4).Get((_172_v44).Fm4(_133_v14, (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _172_v44.F23, _135_globalState)).(_dafny.Int)
			}
			return (_dafny.Zero).Minus(_133_v14)
		})(), _135_globalState)).Cardinality())).Cmp(_dafny.IntOfInt64(278)) < 0 {
			return (_dafny.Zero).Minus(_133_v14)
		}
		return Companion_Default___.Fm0(_172_v44.F23, _118_v0, _135_globalState)
	})()
	var _hi0 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_133_v14, (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)))
	_ = _hi0
	for _174_i4 := (_133_v14).Minus((_dafny.Zero).Minus((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int))); _174_i4.Cmp(_hi0) < 0; _174_i4 = _174_i4.Plus(_dafny.One) {
		_133_v14 = _dafny.IntOfInt64(673)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
		_ = _index35
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
		_ = _index36
		var _rhs22 bool = ((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)).Cmp((_174_i4).Minus(_174_i4)) != 0
		_ = _rhs22
		var _rhs23 _dafny.Array = _127_v9
		_ = _rhs23
		var _rhs24 _dafny.Int = (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)
		_ = _rhs24
		var _rhs25 _dafny.Int = _174_i4
		_ = _rhs25
		var _lhs17 _dafny.Array = _127_v9
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
		_ = _lhs18
		var _lhs19 _dafny.Array = _127_v9
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))
		_ = _lhs20
		_118_v0 = _rhs22
		_127_v9 = _rhs23
		(_lhs17).ArraySet1(_rhs24, (_lhs18).Int())
		(_lhs19).ArraySet1(_rhs25, (_lhs20).Int())
		(_135_globalState).F18 = _dafny.CodePoint('e')
		_118_v0 = !(!(((_dafny.Zero).Minus((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int))).Cmp((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)) >= 0) || (_118_v0))
	}
	var _175_v47 *C0
	_ = _175_v47
	var _nw39 *C0 = New_C0_()
	_ = _nw39
	_nw39.Ctor__(_172_v44.F23, func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(294))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_176_v19 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_177_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_176_v19).Cardinality())
			}
		})(_142_v19)))).Elements()); ; {
			_compr_3, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _178_v46 _dafny.Int
			_178_v46 = interface{}(_compr_3).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(294))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_179_v19 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_177_i6 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_179_v19).Cardinality())
				}
			})(_142_v19))), _178_v46) {
				_coll3.Add((_178_v46).Plus((_122_v4).Cardinality()), _119_v1)
			}
		}
		return _coll3.ToMap()
	}())
	_175_v47 = _nw39
	var _180_v48 _dafny.Map
	_ = _180_v48
	_180_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v47, (_172_v44).F22())
	var _hi1 _dafny.Int = _dafny.IntOfUint32((_142_v19).Cardinality())
	_ = _hi1
	for _181_i5 := (_180_v48).Cardinality(); _181_i5.Cmp(_hi1) < 0; _181_i5 = _181_i5.Plus(_dafny.One) {
		var _182_v49 _dafny.Sequence
		_ = _182_v49
		var _183_v50 bool
		_ = _183_v50
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 bool
		_ = _out1
		_out0, _out1 = (_172_v44).M1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_120_v2).Cardinality()), _181_i5), _175_v47.F24, (_172_v44).F22(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_120_v2, _120_v2)).Cardinality()), _135_globalState)
		_182_v49 = _out0
		_183_v50 = _out1
		var _184_v51 bool
		_ = _184_v51
		var _185_v52 _dafny.Map
		_ = _185_v52
		var _out2 bool
		_ = _out2
		var _out3 _dafny.Map
		_ = _out3
		_out2, _out3 = (_172_v44).M0(_dafny.IntOfInt64(313), _135_globalState)
		_184_v51 = _out2
		_185_v52 = _out3
		var _186_v53 _dafny.Map
		_ = _186_v53
		_186_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_142_v19).Cardinality()), _143_v20)
		var _187_v54 _dafny.Sequence
		_ = _187_v54
		_187_v54 = _dafny.SeqOf(_143_v20, _143_v20)
		var _188_v55 _dafny.MultiSet
		_ = _188_v55
		_188_v55 = _dafny.MultiSetOf(_143_v20, _143_v20, (func() _dafny.Array {
			if (_186_v53).Contains(_181_i5) {
				return (_186_v53).Get(_181_i5).(_dafny.Array)
			}
			return (_187_v54).Select((Companion_Default___.SafeIndex(_133_v14, _dafny.IntOfUint32((_187_v54).Cardinality()))).Uint32()).(_dafny.Array)
		})(), _143_v20, _143_v20)
		var _189_v56 _dafny.MultiSet
		_ = _189_v56
		_189_v56 = _dafny.MultiSetOf(_143_v20)
		var _190_v57 _dafny.Sequence
		_ = _190_v57
		_190_v57 = _dafny.SeqOf((_189_v56))
		_188_v55 = (_190_v57).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_181_i5, _dafny.IntOfInt64(508)), _dafny.IntOfUint32((_190_v57).Cardinality()))).Uint32()).(_dafny.MultiSet)
		_125_v7 = _125_v7
	}
	if (_172_v44).F22() {
		var _191_v58 _dafny.Map
		_ = _191_v58
		_191_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), (_148_v25).Cardinality())
		var _192_v59 _dafny.Map
		_ = _192_v59
		_192_v59 = _191_v58
		_192_v59 = Companion_Default___.Fm12(_135_globalState)
		var _193_v60 _dafny.Array
		_ = _193_v60
		_193_v60 = _dafny.ArrayCastTo((_173_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_173_v45), 0))).Int()))
		var _194_v61 _dafny.Set
		_ = _194_v61
		_194_v61 = _dafny.SetOf(_193_v60, _193_v60)
		var _195_v62 _dafny.Map
		_ = _195_v62
		_195_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v47, _194_v61)
		(_172_v44).F23 = ((_194_v61).Difference(_194_v61)).IsDisjointFrom((func() _dafny.Set {
			if (_195_v62).Contains(_175_v47) {
				return (_195_v62).Get(_175_v47).(_dafny.Set)
			}
			return _194_v61
		})())
		_133_v14 = ((_133_v14).Times(_dafny.IntOfInt64(52))).Plus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v47, _172_v44.F23)).Merge(_180_v48)).Cardinality())
		_118_v0 = Companion_Default___.Fm11(_125_v7, _133_v14, _135_globalState)
		(_135_globalState).F19 = (_172_v44).F22()
	} else {
		var _196_v63 _dafny.MultiSet
		_ = _196_v63
		_196_v63 = _dafny.MultiSetOf(_143_v20, _143_v20)
		var _197_v64 _dafny.MultiSet
		_ = _197_v64
		_197_v64 = (_196_v63).Union(_196_v63)
		var _source1 _dafny.MultiSet = _197_v64
		_ = _source1
		var _198___mcc_h0 _dafny.MultiSet = _source1
		_ = _198___mcc_h0
		var _199_cf18 _dafny.MultiSet = _198___mcc_h0
		_ = _199_cf18
		var _200_v65 _dafny.Map
		_ = _200_v65
		_200_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(938))
		var _201_v66 _dafny.Map
		_ = _201_v66
		_201_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v65).Cardinality(), _dafny.ArrayCastTo((_173_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_173_v45), 0))).Int())))
		_201_v66 = (_201_v66).Update(_dafny.IntOfInt64(751), _dafny.ArrayCastTo((_173_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_173_v45), 0))).Int())))
		var _nwElement0_8 _dafny.Int = (_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int)
		_ = _nwElement0_8
		var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(14))
		_ = _nw40
		(_nw40).ArraySet1(_nwElement0_8, 0)
		(_nw40).ArraySet1(Companion_Default___.SafeModuloInt((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _133_v14), 1)
		(_nw40).ArraySet1((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), 2)
		(_nw40).ArraySet1(Companion_Default___.Fm0(!(true), _172_v44.F23, _135_globalState), 3)
		(_nw40).ArraySet1((_dafny.IntOfInt64(338)).Times(_133_v14), 4)
		(_nw40).ArraySet1((func() _dafny.Int {
			if (_125_v7).Contains(_172_v44.F23) {
				return (_125_v7).Multiplicity(_172_v44.F23)
			}
			return _133_v14
		})(), 5)
		(_nw40).ArraySet1(((_dafny.Zero).Minus(_133_v14)).Plus(_dafny.IntOfUint32((_120_v2).Cardinality())), 6)
		(_nw40).ArraySet1(_133_v14, 7)
		(_nw40).ArraySet1((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), 8)
		(_nw40).ArraySet1((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), 9)
		(_nw40).ArraySet1(Companion_Default___.SafeModuloInt((_142_v19).Select((Companion_Default___.SafeIndex((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_142_v19).Cardinality()))).Uint32()).(_dafny.Int), _133_v14), 10)
		(_nw40).ArraySet1((_dafny.Zero).Minus(_133_v14), 11)
		(_nw40).ArraySet1((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), 12)
		(_nw40).ArraySet1((_127_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_127_v9), 0))).Int()).(_dafny.Int), 13)
		_127_v9 = _nw40
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_173_v45), 0))
		_ = _index37
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_2
		var _nw41 _dafny.Array
		_ = _nw41
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw41 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_202_v14 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_203_i7 _dafny.Int) _dafny.Int {
					return (_203_i7).Plus(_202_v14)
				}
			})(_133_v14)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw41 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw41).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw41).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		(_173_v45).ArraySet1(_nw41, (_index37).Int())
		var _204_v67 _dafny.Set
		_ = _204_v67
		_204_v67 = _dafny.SetOf(_172_v44)
		_204_v67 = _204_v67
		var _205_v68 *C0
		_ = _205_v68
		var _nw42 *C0 = New_C0_()
		_ = _nw42
		_nw42.Ctor__(_175_v47.F24, (_175_v47).F25())
		_205_v68 = _nw42
		(_175_v47).F24 = (Companion_Default___.SafeDivisionInt((_170_v42).Fm8(_135_globalState), _dafny.IntOfInt64(-823))).Cmp(_133_v14) == 0
		(_175_v47).F24 = _172_v44.F23
		var _206_v69 bool
		_ = _206_v69
		var _207_v70 _dafny.Map
		_ = _207_v70
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Map
		_ = _out5
		_out4, _out5 = (_172_v44).M0(_dafny.IntOfUint32((_142_v19).Cardinality()), _135_globalState)
		_206_v69 = _out4
		_207_v70 = _out5
	}
	_dafny.Print(_118_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_119_v1, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_120_v2.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_121_v3).Dtor_cf1(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v3).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC1_(_dafny.SeqOf(true), _dafny.IntOfInt64(9)), _dafny.MultiSetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v6).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v7).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_126_v8, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v9).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_130_v11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v12).Equals(_dafny.SetOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_132_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_v14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(588)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F0).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(36))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_globalState).F3()).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_135_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_135_globalState).F10(), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"), _dafny.UnicodeSeqOfUtf8Bytes("pvnnsysbg"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F11().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState.F13).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_globalState).F15()).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_globalState).F16()).Equals(_dafny.SetOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_globalState).F17()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_135_globalState.F18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_135_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_globalState).F20()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(588)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v16).Equals(_dafny.SetOf(_dafny.SeqOf(true), _dafny.SeqOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(588), _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_141_v18, _dafny.SeqOf(_dafny.MultiSetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_142_v19, _dafny.SeqOf(_dafny.IntOfInt64(588), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_144_v21).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_144_v21).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_144_v21).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_144_v21).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_v22).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(944), Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(false, _dafny.IntOfInt64(754), _dafny.IntOfInt64(561), false))).UpdateUnsafe(_dafny.One, Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(Companion_D0_.Create_DC2_(true, _dafny.IntOfInt64(111), _dafny.IntOfInt64(-539), false))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_146_v23).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.CodePoint('j')), _dafny.IntOfInt64(1176))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v24).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.CodePoint('j')), _dafny.IntOfInt64(991))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_v25).Equals(_dafny.SetOf(_dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_149_v26).Equals(_dafny.MultiSetOf(_dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_150_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v41).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(588), _dafny.SeqOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_171_v43).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v44).F22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_172_v44.F23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_173_v45).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_v47.F24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_v47).F25()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(3), _dafny.SeqOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v48).Cardinality())
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Sequence, Cf2 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf3 bool
	Cf4 _dafny.Int
	Cf5 _dafny.Int
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf3 bool, Cf4 _dafny.Int, Cf5 _dafny.Int, Cf6 bool) D0 {
	return D0{D0_DC2{Cf3, Cf4, Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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

type D0_DC3 struct {
	Cf7 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf7 D0) D0 {
	return D0{D0_DC3{Cf7}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf7() D0 {
	return _this.Get_().(D0_DC3).Cf7
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
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

type D1_DC4 struct {
	Cf8 _dafny.Map
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.Map) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D1) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf9 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Array) D2 {
	return D2{D2_DC5{Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf9 == data2.Cf9
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
	Cf11 _dafny.Map
	Cf12 bool
	Cf13 bool
	Cf14 bool
	Cf15 _dafny.Array
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Map, Cf12 bool, Cf13 bool, Cf14 bool, Cf15 _dafny.Array) D3 {
	return D3{D3_DC7{Cf11, Cf12, Cf13, Cf14, Cf15}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC8 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf16 _dafny.Int, Cf17 _dafny.Int) D3 {
	return D3{D3_DC8{Cf16, Cf17}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC6 struct {
	Cf10 _dafny.Sequence
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf10 _dafny.Sequence) D3 {
	return D3{D3_DC6{Cf10}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_(_dafny.EmptyMap, false, false, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D3) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC7).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC7).Cf13
}

func (_this D3) Dtor_cf14() bool {
	return _this.Get_().(D3_DC7).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D3_DC7).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf17
}

func (_this D3) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D3_DC6).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + data.Cf10.VerbatimString(true) + ")"
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
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11) && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
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

type D4_DC9 struct {
	Cf18 _dafny.MultiSet
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf18 _dafny.MultiSet) D4 {
	return D4{D4_DC9{Cf18}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D4) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D4_DC9).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D5_DC11 struct {
	Cf20 bool
	Cf21 bool
	Cf22 _dafny.Array
	Cf23 T0
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf20 bool, Cf21 bool, Cf22 _dafny.Array, Cf23 T0) D5 {
	return D5{D5_DC11{Cf20, Cf21, Cf22, Cf23}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC10 struct {
	Cf19 _dafny.CodePoint
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf19 _dafny.CodePoint) D5 {
	return D5{D5_DC10{Cf19}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC11_(false, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), (T0)(nil))
}

func (_this D5) Dtor_cf20() bool {
	return _this.Get_().(D5_DC11).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC11).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC11).Cf22
}

func (_this D5) Dtor_cf23() T0 {
	return _this.Get_().(D5_DC11).Cf23
}

func (_this D5) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D5_DC10).Cf19
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && _dafny.AreEqual(data1.Cf23, data2.Cf23)
		}
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf19 == data2.Cf19
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm8(globalState *GlobalState) _dafny.Int
	Fm9(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.CodePoint
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
	F0   _dafny.Map
	F6   bool
	F13  _dafny.Array
	F18  _dafny.CodePoint
	F19  bool
	_f1  bool
	_f2  bool
	_f3  _dafny.MultiSet
	_f4  bool
	_f5  _dafny.Int
	_f7  _dafny.Int
	_f8  bool
	_f9  _dafny.Int
	_f10 _dafny.Sequence
	_f11 _dafny.Sequence
	_f12 _dafny.Int
	_f14 _dafny.Int
	_f15 _dafny.Map
	_f16 _dafny.Set
	_f17 _dafny.Map
	_f20 _dafny.Map
	_f21 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptyMap
	_this.F6 = false
	_this.F13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F18 = _dafny.CodePoint('D')
	_this.F19 = false
	_this._f1 = false
	_this._f2 = false
	_this._f3 = _dafny.EmptyMultiSet
	_this._f4 = false
	_this._f5 = _dafny.Zero
	_this._f7 = _dafny.Zero
	_this._f8 = false
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.EmptySeq
	_this._f11 = _dafny.EmptySeq
	_this._f12 = _dafny.Zero
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.EmptyMap
	_this._f16 = _dafny.EmptySet
	_this._f17 = _dafny.EmptyMap
	_this._f20 = _dafny.EmptyMap
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 bool, f2 bool, f3 _dafny.MultiSet, f4 bool, f5 _dafny.Int, f6 bool, f7 _dafny.Int, f8 bool, f9 _dafny.Int, f10 _dafny.Sequence, f11 _dafny.Sequence, f12 _dafny.Int, f13 _dafny.Array, f14 _dafny.Int, f15 _dafny.Map, f16 _dafny.Set, f17 _dafny.Map, f18 _dafny.CodePoint, f19 bool, f20 _dafny.Map, f21 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this).F18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.MultiSet {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
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
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Sequence {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Sequence {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Int {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Map {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() _dafny.Set {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Map {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F20() _dafny.Map {
	{
		return _this._f20
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
	F24  bool
	_f25 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F24 = false
	_this._f25 = _dafny.EmptyMap
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

func (_this *C0) Ctor__(f24 bool, f25 _dafny.Map) {
	{
		(_this).F24 = f24
		(_this)._f25 = f25
	}
}
func (_this *C0) Fm8(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(694)
	}
}
func (_this *C0) Fm9(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('h')
	}
}
func (_this *C0) F25() _dafny.Map {
	{
		return _this._f25
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	F23  bool
	_f22 bool
}

func New_C1_() *C1 {
	_this := C1{}

	_this.F23 = false
	_this._f22 = false
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

func (_this *C1) Ctor__(f22 bool, f23 bool) {
	{
		(_this)._f22 = f22
		(_this).F23 = f23
	}
}
func (_this *C1) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		return (_this).F22()
	}
}
func (_this *C1) M0(p0 _dafny.Int, globalState *GlobalState) (bool, _dafny.Map) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var _208_v0 _dafny.CodePoint
		_ = _208_v0
		_208_v0 = _dafny.CodePoint('a')
		var _209_v1 _dafny.Map
		_ = _209_v1
		_209_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_208_v0, false)
		var _210_v2 _dafny.Sequence
		_ = _210_v2
		_210_v2 = _dafny.SeqOf(_209_v1)
		_210_v2 = _210_v2
		var _211_v3 _dafny.Sequence
		_ = _211_v3
		_211_v3 = _dafny.SeqOf((_this).F22())
		var _212_v4 _dafny.Sequence
		_ = _212_v4
		_212_v4 = _dafny.SeqOf(_211_v3, _211_v3)
		var _213_v5 D0
		_ = _213_v5
		_213_v5 = Companion_D0_.Create_DC1_((_212_v4).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_212_v4).Cardinality()))).Uint32()).(_dafny.Sequence), p0)
		var _pat_let_tv3 = p0
		_ = _pat_let_tv3
		var _source2 D0 = func(_pat_let4_0 D0) D0 {
			return func(_214_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let5_0 _dafny.Int) D0 {
					return func(_215_dt__update_hcf2_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_((_214_dt__update__tmp_h0).Dtor_cf1(), _215_dt__update_hcf2_h0)
					}(_pat_let5_0)
				}(_pat_let_tv3)
			}(_pat_let4_0)
		}(_213_v5)
		_ = _source2
		if _source2.Is_DC1() {
			var _216___mcc_h0 _dafny.Sequence = _source2.Get_().(D0_DC1).Cf1
			_ = _216___mcc_h0
			var _217___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
			_ = _217___mcc_h1
			var _218_cf2 _dafny.Int = _217___mcc_h1
			_ = _218_cf2
			var _219_cf1 _dafny.Sequence = _216___mcc_h0
			_ = _219_cf1
			var _220_v6 _dafny.Sequence
			_ = _220_v6
			_220_v6 = _dafny.UnicodeSeqOfUtf8Bytes("kgrcsul")
			var _221_v7 _dafny.Map
			_ = _221_v7
			_221_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v6, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_218_cf2, _dafny.IntOfUint32((_219_cf1).Cardinality()))))
			_221_v7 = (_221_v7).Update(_dafny.Companion_Sequence_.Concatenate(_220_v6, _220_v6), (_dafny.Zero).Minus(_218_cf2))
			var _222_v8 _dafny.Array
			_ = _222_v8
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_3
			var _nw43 _dafny.Array
			_ = _nw43
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw43 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Int = (func(_223_cf2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_224_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_224_i0, _223_cf2)
					}
				})(_218_cf2)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw43 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw43).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw43).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_222_v8 = _nw43
			var _225_v9 _dafny.MultiSet
			_ = _225_v9
			_225_v9 = _dafny.MultiSetOf(_222_v8, _222_v8)
			if ((_225_v9).IsProperSubsetOf(_225_v9)) && (_dafny.Companion_Sequence_.IsPrefixOf(_211_v3, _dafny.SeqOf((_this).Fm4(_dafny.IntOfInt64(-705), _218_cf2, _this.F23, globalState)))) {
				(globalState).F19 = !(!((_this).F22()) || ((_this).F22()))
				var _rhs26 _dafny.Array = _222_v8
				_ = _rhs26
				var _rhs27 _dafny.Int = (func() _dafny.Int {
					if (_221_v7).Contains(_220_v6) {
						return (_221_v7).Get(_220_v6).(_dafny.Int)
					}
					return _218_cf2
				})()
				_ = _rhs27
				var _lhs21 *GlobalState = globalState
				_ = _lhs21
				_lhs21.F13 = _rhs26
				_218_cf2 = _rhs27
				_218_cf2 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_218_cf2, _218_cf2), p0))
				(globalState).F18 = _208_v0
				_220_v6 = _dafny.UnicodeSeqOfUtf8Bytes("upfp")
			} else {
				(_this).F23 = (_this).F22()
				(globalState).F6 = !(false)
				var _226_v10 _dafny.Map
				_ = _226_v10
				_226_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _219_cf1)
				var _227_v11 _dafny.Set
				_ = _227_v11
				_227_v11 = _dafny.SetOf(_208_v0, _208_v0, Companion_Default___.Fm5(globalState), _208_v0)
				var _228_v12 _dafny.Sequence
				_ = _228_v12
				_228_v12 = _dafny.SeqOf(_227_v11, _227_v11)
				var _rhs28 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_211_v3, (func() _dafny.Sequence {
					if (_226_v10).Contains(_this.F23) {
						return (_226_v10).Get(_this.F23).(_dafny.Sequence)
					}
					return _219_cf1
				})()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_228_v12).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_211_v3, (func() _dafny.Sequence {
					if (_226_v10).Contains(_this.F23) {
						return (_226_v10).Get(_this.F23).(_dafny.Sequence)
					}
					return _219_cf1
				})())).Cardinality()))).Uint32(), _this.F23), _dafny.SeqOf(_this.F23, _this.F23))).Cardinality()))
				_ = _rhs28
				var _rhs29 bool = _this.F23
				_ = _rhs29
				var _lhs22 *GlobalState = globalState
				_ = _lhs22
				_218_cf2 = _rhs28
				_lhs22.F19 = _rhs29
				var _229_v13 _dafny.MultiSet
				_ = _229_v13
				_229_v13 = _dafny.MultiSetOf(false)
				var _230_v14 _dafny.Sequence
				_ = _230_v14
				_230_v14 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("vtobwgfoy"))
				var _231_v15 _dafny.Sequence
				_ = _231_v15
				_231_v15 = _dafny.SeqOf(_229_v13, ((_229_v13).Update((_this).F22(), Companion_Default___.Abs(_dafny.IntOfUint32((_230_v14).Cardinality())))).Union(_229_v13), _dafny.MultiSetOf(true, (_this).F22()))
				var _232_v16 _dafny.Set
				_ = _232_v16
				_232_v16 = _dafny.SetOf(_218_cf2)
				var _rhs30 bool = (func() bool {
					if (_232_v16).Contains(_218_cf2) {
						return (_225_v9).IsSubsetOf(_225_v9)
					}
					return (func() bool {
						if _this.F23 {
							return !((_this).F22())
						}
						return _this.F23
					})()
				})()
				_ = _rhs30
				var _rhs31 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_231_v15, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer10 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_233_v13 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_234_i1 _dafny.Int) _dafny.MultiSet {
						return _233_v13
					}
				})(_229_v13))))
				_ = _rhs31
				var _lhs23 *GlobalState = globalState
				_ = _lhs23
				_lhs23.F6 = _rhs30
				_231_v15 = _rhs31
				var _235_v17 _dafny.Array
				_ = _235_v17
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_4
				var _nw44 _dafny.Array
				_ = _nw44
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw44 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) bool = func(_236_i2 _dafny.Int) bool {
						return (func() bool {
							if (_this).F22() {
								return (_this).F22()
							}
							return (_this).F22()
						})()
					}
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw44 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw44).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw44).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_235_v17 = _nw44
				_235_v17 = _235_v17
			}
			var _237_v18 _dafny.Array
			_ = _237_v18
			var _nw45 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw45
			_237_v18 = _nw45
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_237_v18), 0))
			_ = _index38
			(_237_v18).ArraySet1(_this.F23, (_index38).Int())
			var _238_v19 _dafny.Set
			_ = _238_v19
			_238_v19 = _dafny.SetOf(_218_cf2, p0)
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_237_v18), 0))
			_ = _index39
			var _rhs32 _dafny.Int = _dafny.IntOfUint32((_220_v6).Cardinality())
			_ = _rhs32
			var _rhs33 bool = (_211_v3).Select((Companion_Default___.SafeIndex((_238_v19).Cardinality(), _dafny.IntOfUint32((_211_v3).Cardinality()))).Uint32()).(bool)
			_ = _rhs33
			var _rhs34 bool = _this.F23
			_ = _rhs34
			var _rhs35 _dafny.Sequence = _211_v3
			_ = _rhs35
			var _lhs24 *GlobalState = globalState
			_ = _lhs24
			var _lhs25 _dafny.Array = _237_v18
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_237_v18), 0))
			_ = _lhs26
			_218_cf2 = _rhs32
			_lhs24.F19 = _rhs33
			(_lhs25).ArraySet1(_rhs34, (_lhs26).Int())
			_211_v3 = _rhs35
			var _239_v20 _dafny.Sequence
			_ = _239_v20
			_239_v20 = _dafny.SeqOf(p0)
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_222_v8), 0))
			_ = _index40
			(_222_v8).ArraySet1(_dafny.IntOfUint32((_239_v20).Cardinality()), (_index40).Int())
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_237_v18), 0))
			_ = _index41
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_222_v8), 0))
			_ = _index42
			var _rhs36 bool = (_237_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_237_v18), 0))).Int()).(bool)
			_ = _rhs36
			var _rhs37 _dafny.Int = (_dafny.Zero).Minus(_218_cf2)
			_ = _rhs37
			var _rhs38 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_219_cf1, _219_cf1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_220_v6).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_219_cf1, _219_cf1)).Cardinality()))).Uint32(), _this.F23), (_this).F22())
			_ = _rhs38
			var _lhs27 _dafny.Array = _237_v18
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_237_v18), 0))
			_ = _lhs28
			var _lhs29 _dafny.Array = _222_v8
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_222_v8), 0))
			_ = _lhs30
			var _lhs31 *GlobalState = globalState
			_ = _lhs31
			(_lhs27).ArraySet1(_rhs36, (_lhs28).Int())
			(_lhs29).ArraySet1(_rhs37, (_lhs30).Int())
			_lhs31.F19 = _rhs38
		} else if _source2.Is_DC2() {
			var _240___mcc_h2 bool = _source2.Get_().(D0_DC2).Cf3
			_ = _240___mcc_h2
			var _241___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC2).Cf4
			_ = _241___mcc_h3
			var _242___mcc_h4 _dafny.Int = _source2.Get_().(D0_DC2).Cf5
			_ = _242___mcc_h4
			var _243___mcc_h5 bool = _source2.Get_().(D0_DC2).Cf6
			_ = _243___mcc_h5
			var _244_cf6 bool = _243___mcc_h5
			_ = _244_cf6
			var _245_cf5 _dafny.Int = _242___mcc_h4
			_ = _245_cf5
			var _246_cf4 _dafny.Int = _241___mcc_h3
			_ = _246_cf4
			var _247_cf3 bool = _240___mcc_h2
			_ = _247_cf3
			if _this.F23 {
				var _248_v21 _dafny.Array
				_ = _248_v21
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_5
				var _nw46 _dafny.Array
				_ = _nw46
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw46 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.Int = (func(_249_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_250_i3 _dafny.Int) _dafny.Int {
							return (_250_i3).Plus(_249_cf4)
						}
					})(_246_cf4)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw46 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw46).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw46).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_248_v21 = _nw46
				var _251_v22 _dafny.Map
				_ = _251_v22
				_251_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _248_v21)
				_251_v22 = (_251_v22).Update(_244_cf6, _248_v21)
				var _252_v23 _dafny.Array
				_ = _252_v23
				var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
				_ = _nw47
				_252_v23 = _nw47
				var _253_v24 _dafny.Map
				_ = _253_v24
				_253_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_244_cf6, p0)
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_252_v23), 0))
				_ = _index43
				(_252_v23).ArraySet1(_253_v24, (_index43).Int())
				var _254_v25 _dafny.Sequence
				_ = _254_v25
				_254_v25 = _dafny.SeqOf(_253_v24)
				var _255_v26 _dafny.Sequence
				_ = _255_v26
				_255_v26 = _dafny.SeqOf((_253_v24).Cardinality())
				var _256_v27 _dafny.Sequence
				_ = _256_v27
				_256_v27 = _dafny.SeqOf((_255_v26).Select((Companion_Default___.SafeIndex(_246_cf4, _dafny.IntOfUint32((_255_v26).Cardinality()))).Uint32()).(_dafny.Int))
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_252_v23), 0))
				_ = _index44
				(_252_v23).ArraySet1(((_253_v24).Merge(_253_v24)).Merge((_254_v25).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_256_v27).Cardinality()), _dafny.IntOfUint32((_254_v25).Cardinality()))).Uint32()).(_dafny.Map)), (_index44).Int())
				var _257_v28 _dafny.Sequence
				_ = _257_v28
				var _258_v29 bool
				_ = _258_v29
				var _out6 _dafny.Sequence
				_ = _out6
				var _out7 bool
				_ = _out7
				_out6, _out7 = (_this).M1(((_252_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_252_v23), 0))).Int()).(_dafny.Map)).Cardinality(), _this.F23, (_this).F22(), _246_cf4, globalState)
				_257_v28 = _out6
				_258_v29 = _out7
				_246_cf4 = _246_cf4
				_246_cf4 = _dafny.IntOfInt64(669)
			} else {
				_245_cf5 = p0
				var _259_v30 _dafny.Array
				_ = _259_v30
				var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw48
				_259_v30 = _nw48
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_259_v30), 0))
				_ = _index45
				(_259_v30).ArraySet1((_dafny.Zero).Minus((_245_cf5).Plus(_245_cf5)), (_index45).Int())
				var _260_v31 _dafny.Map
				_ = _260_v31
				_260_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_244_cf6, _dafny.IntOfInt64(450))
				var _261_v32 _dafny.Sequence
				_ = _261_v32
				_261_v32 = _dafny.SeqOf((_dafny.Zero).Minus(_245_cf5), p0)
				var _262_v33 _dafny.Map
				_ = _262_v33
				_262_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v31, (_this).Fm4(_dafny.IntOfInt64(-224), (_261_v32).Select((Companion_Default___.SafeIndex(_246_cf4, _dafny.IntOfUint32((_261_v32).Cardinality()))).Uint32()).(_dafny.Int), _244_cf6, globalState))
				var _263_v34 _dafny.Map
				_ = _263_v34
				_263_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_246_cf4, _247_cf3)
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_259_v30), 0))
				_ = _index46
				var _rhs39 _dafny.Int = _246_cf4
				_ = _rhs39
				var _rhs40 _dafny.Int = (_262_v33).Cardinality()
				_ = _rhs40
				var _rhs41 _dafny.Int = ((_263_v34).Cardinality()).Plus((_dafny.Zero).Minus(_245_cf5))
				_ = _rhs41
				var _rhs42 _dafny.Int = (_dafny.IntOfInt64(274)).Times(_dafny.IntOfUint32((_261_v32).Cardinality()))
				_ = _rhs42
				var _rhs43 bool = !(_244_cf6)
				_ = _rhs43
				var _lhs32 _dafny.Array = _259_v30
				_ = _lhs32
				var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_259_v30), 0))
				_ = _lhs33
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				_245_cf5 = _rhs39
				_245_cf5 = _rhs40
				(_lhs32).ArraySet1(_rhs41, (_lhs33).Int())
				_245_cf5 = _rhs42
				_lhs34.F19 = _rhs43
				var _264_v35 _dafny.MultiSet
				_ = _264_v35
				_264_v35 = _dafny.MultiSetOf(_244_cf6)
				var _265_v36 D0
				_ = _265_v36
				_265_v36 = Companion_D0_.Create_DC2_((_this).F22(), p0, (_264_v35).Cardinality(), _244_cf6)
				var _266_v37 _dafny.Map
				_ = _266_v37
				_266_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_265_v36, (_this).Fm4(_245_cf5, (_259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_259_v30), 0))).Int()).(_dafny.Int), _244_cf6, globalState))
				_266_v37 = (_266_v37).Update(_265_v36, (_this).F22())
				var _267_v38 _dafny.Sequence
				_ = _267_v38
				_267_v38 = _dafny.UnicodeSeqOfUtf8Bytes("nvttpt")
				var _268_v39 _dafny.Array
				_ = _268_v39
				var _nwElement0_9 bool = (p0).Cmp((_261_v32).Select((Companion_Default___.SafeIndex((_259_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_259_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_261_v32).Cardinality()))).Uint32()).(_dafny.Int)) == 0
				_ = _nwElement0_9
				var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(13))
				_ = _nw49
				(_nw49).ArraySet1(_nwElement0_9, 0)
				(_nw49).ArraySet1(_244_cf6, 1)
				(_nw49).ArraySet1((_247_cf3) == (_244_cf6), 2)
				(_nw49).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_267_v38, _267_v38), 3)
				(_nw49).ArraySet1(!(_244_cf6), 4)
				(_nw49).ArraySet1((_this).F22(), 5)
				(_nw49).ArraySet1((_this).Fm4(_245_cf5, Companion_Default___.Fm0(!(!(_244_cf6)), _this.F23, globalState), false, globalState), 6)
				(_nw49).ArraySet1(_dafny.Companion_Sequence_.Equal(_211_v3, _211_v3), 7)
				(_nw49).ArraySet1(_this.F23, 8)
				(_nw49).ArraySet1((_this).F22(), 9)
				(_nw49).ArraySet1(_247_cf3, 10)
				(_nw49).ArraySet1(_this.F23, 11)
				(_nw49).ArraySet1(_this.F23, 12)
				_268_v39 = _nw49
				var _269_v40 D0
				_ = _269_v40
				_269_v40 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(554))
				var _270_v41 _dafny.Sequence
				_ = _270_v41
				_270_v41 = _dafny.SeqOf(_268_v39, _268_v39)
				var _rhs44 _dafny.Int = (_dafny.Zero).Minus((_269_v40).Dtor_cf0())
				_ = _rhs44
				var _rhs45 bool = true
				_ = _rhs45
				var _rhs46 bool = !((func() bool {
					if (_this).F22() {
						return (_this).Fm4(_245_cf5, (_dafny.MultiSetFromSeq(_261_v32)).Cardinality(), _244_cf6, globalState)
					}
					return _244_cf6
				})())
				_ = _rhs46
				var _rhs47 _dafny.Array = (_270_v41).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_270_v41).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs47
				var _lhs35 *C1 = _this
				_ = _lhs35
				var _lhs36 *C1 = _this
				_ = _lhs36
				_245_cf5 = _rhs44
				_lhs35.F23 = _rhs45
				_lhs36.F23 = _rhs46
				_268_v39 = _rhs47
				var _271_v42 _dafny.MultiSet
				_ = _271_v42
				_271_v42 = _dafny.MultiSetOf(_dafny.IntOfUint32((_211_v3).Cardinality()), ((Companion_Default___.Fm6(_246_cf4, _247_cf3, globalState)).Difference((_264_v35).Update(false, Companion_Default___.Abs(p0)))).Cardinality(), _245_cf5)
				_271_v42 = _271_v42
			}
			var _272_v43 _dafny.Sequence
			_ = _272_v43
			_272_v43 = _dafny.UnicodeSeqOfUtf8Bytes("tbjypvnv")
			var _273_v44 _dafny.Sequence
			_ = _273_v44
			_273_v44 = _dafny.SeqOf((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_272_v43, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_272_v43).Cardinality()))).Uint32(), _208_v0)).Cardinality())).Times(_246_cf4), p0, _245_cf5, _245_cf5, p0)
			_273_v44 = _dafny.Companion_Sequence_.Concatenate(_273_v44, _273_v44)
			if !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_272_v43, _272_v43), _272_v43) {
				var _274_v45 _dafny.MultiSet
				_ = _274_v45
				_274_v45 = _dafny.MultiSetOf(_247_cf3, !(!(_this.F23)))
				var _275_v46 _dafny.Map
				_ = _275_v46
				_275_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_274_v45, _208_v0)
				_275_v46 = (_275_v46).Update((_274_v45).Union(_274_v45), _208_v0)
				var _276_v47 _dafny.Array
				_ = _276_v47
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw50
				_276_v47 = _nw50
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_276_v47), 0))
				_ = _index47
				(_276_v47).ArraySet1((_245_cf5).Cmp(p0) == 0, (_index47).Int())
				var _277_v48 _dafny.Array
				_ = _277_v48
				var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
				_ = _nw51
				_277_v48 = _nw51
				var _278_v49 _dafny.Map
				_ = _278_v49
				_278_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_273_v44)).Cardinality())), p0)
				var _279_v50 _dafny.Map
				_ = _279_v50
				_279_v50 = _278_v49
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_277_v48), 0))
				_ = _index48
				(_277_v48).ArraySet1((_278_v49).Merge((_279_v50)), (_index48).Int())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_276_v47), 0))
				_ = _index49
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_277_v48), 0))
				_ = _index50
				var _rhs48 bool = !((_this).F22())
				_ = _rhs48
				var _rhs49 _dafny.Map = _278_v49
				_ = _rhs49
				var _lhs37 _dafny.Array = _276_v47
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_276_v47), 0))
				_ = _lhs38
				var _lhs39 _dafny.Array = _277_v48
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_277_v48), 0))
				_ = _lhs40
				(_lhs37).ArraySet1(_rhs48, (_lhs38).Int())
				(_lhs39).ArraySet1(_rhs49, (_lhs40).Int())
				r0 = (_this).Fm4(_dafny.IntOfInt64(485), (_dafny.Zero).Minus((_245_cf5).Times(_245_cf5)), _this.F23, globalState)
				var _280_v51 _dafny.Map
				_ = _280_v51
				_280_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_cf3, _244_cf6)
				var _281_v52 _dafny.Map
				_ = _281_v52
				_281_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v5, ((func() bool {
					if (_280_v51).Contains(_244_cf6) {
						return (_280_v51).Get(_244_cf6).(bool)
					}
					return !(_244_cf6)
				})()) == (_this.F23))
				r0 = (func() bool {
					if (_281_v52).Contains(_213_v5) {
						return (_281_v52).Get(_213_v5).(bool)
					}
					return _this.F23
				})()
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_276_v47), 0))
				_ = _index51
				(_276_v47).ArraySet1(_244_cf6, (_index51).Int())
			} else {
				var _282_v53 _dafny.Set
				_ = _282_v53
				_282_v53 = _dafny.SetOf(_dafny.IntOfInt64(738), _245_cf5, _246_cf4, _dafny.IntOfUint32((_272_v43).Cardinality()))
				var _283_v54 D0
				_ = _283_v54
				_283_v54 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(818))
				_282_v53 = _dafny.SetOf((_283_v54).Dtor_cf0(), _246_cf4, p0)
				_245_cf5 = _245_cf5
				(globalState).F19 = _this.F23
				_244_cf6 = (_this).Fm4(Companion_Default___.Fm0(false, (_this).F22(), globalState), _246_cf4, _dafny.Companion_Sequence_.IsProperPrefixOf(_273_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(157))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_284_i4 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality())
				}))), globalState)
				var _285_v55 _dafny.Set
				_ = _285_v55
				_285_v55 = _dafny.SetOf((_this).F22())
				var _286_v56 _dafny.MultiSet
				_ = _286_v56
				_286_v56 = _dafny.MultiSetOf(_247_cf3, _244_cf6)
				var _287_v57 D0
				_ = _287_v57
				_287_v57 = Companion_D0_.Create_DC2_(!((_this).F22()), p0, _dafny.IntOfInt64(301), (_this).F22())
				var _288_v58 _dafny.Array
				_ = _288_v58
				var _nwElement0_10 _dafny.Int = _245_cf5
				_ = _nwElement0_10
				var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(21))
				_ = _nw52
				(_nw52).ArraySet1(_nwElement0_10, 0)
				(_nw52).ArraySet1(_245_cf5, 1)
				(_nw52).ArraySet1((Companion_Default___.Fm7(globalState)).Cardinality(), 2)
				(_nw52).ArraySet1(_245_cf5, 3)
				(_nw52).ArraySet1((_285_v55).Cardinality(), 4)
				(_nw52).ArraySet1(p0, 5)
				(_nw52).ArraySet1(_245_cf5, 6)
				(_nw52).ArraySet1(_246_cf4, 7)
				(_nw52).ArraySet1((_286_v56).Cardinality(), 8)
				(_nw52).ArraySet1(p0, 9)
				(_nw52).ArraySet1(_245_cf5, 10)
				(_nw52).ArraySet1((_273_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.IntOfUint32((_273_v44).Cardinality()))).Uint32()).(_dafny.Int), 11)
				(_nw52).ArraySet1(_246_cf4, 12)
				(_nw52).ArraySet1(_245_cf5, 13)
				(_nw52).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).Fm4(_246_cf4, _dafny.IntOfInt64(923), (_this).F22(), globalState))).Cardinality()), 14)
				(_nw52).ArraySet1(_245_cf5, 15)
				(_nw52).ArraySet1(_246_cf4, 16)
				(_nw52).ArraySet1(_245_cf5, 17)
				(_nw52).ArraySet1(Companion_Default___.Fm0(!(_this.F23), (_287_v57).Dtor_cf6(), globalState), 18)
				(_nw52).ArraySet1(_245_cf5, 19)
				(_nw52).ArraySet1(_246_cf4, 20)
				_288_v58 = _nw52
				var _289_v59 _dafny.Map
				_ = _289_v59
				_289_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_288_v58, _286_v56)
				_289_v59 = (_289_v59).Update(_288_v58, _286_v56)
			}
			var _290_v60 _dafny.Array
			_ = _290_v60
			var _nwElement0_11 bool = _244_cf6
			_ = _nwElement0_11
			var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(4))
			_ = _nw53
			(_nw53).ArraySet1(_nwElement0_11, 0)
			(_nw53).ArraySet1(_244_cf6, 1)
			(_nw53).ArraySet1((_this).F22(), 2)
			(_nw53).ArraySet1(!(false), 3)
			_290_v60 = _nw53
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_290_v60), 0))
			_ = _index52
			(_290_v60).ArraySet1(_244_cf6, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_290_v60), 0))
			_ = _index53
			(_290_v60).ArraySet1(true, (_index53).Int())
		} else if _source2.Is_DC0() {
			var _291___mcc_h6 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
			_ = _291___mcc_h6
			var _292_cf0 _dafny.Int = _291___mcc_h6
			_ = _292_cf0
			var _293_v61 _dafny.Set
			_ = _293_v61
			_293_v61 = _dafny.SetOf(_211_v3, _211_v3)
			(globalState).F6 = (_293_v61).IsProperSubsetOf(_293_v61)
			if _this.F23 {
				var _294_v62 _dafny.Sequence
				_ = _294_v62
				_294_v62 = _dafny.SeqOf(_292_cf0, _292_cf0)
				var _295_v63 _dafny.Map
				_ = _295_v63
				_295_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _294_v62)
				var _296_v64 _dafny.Set
				_ = _296_v64
				_296_v64 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(923))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg12 _dafny.Int) interface{} {
						return coer12(arg12)
					}
				}((func(_297_cf0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_298_i5 _dafny.Int) _dafny.Int {
						return _297_cf0
					}
				})(_292_cf0))), _294_v62, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_295_v63).Contains(true) {
						return (_295_v63).Get(true).(_dafny.Sequence)
					}
					return _294_v62
				})(), (Companion_Default___.SafeIndex(_292_cf0, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_295_v63).Contains(true) {
						return (_295_v63).Get(true).(_dafny.Sequence)
					}
					return _294_v62
				})()).Cardinality()))).Uint32(), _dafny.IntOfInt64(279)))
				var _299_v65 _dafny.Map
				_ = _299_v65
				_299_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _this.F23)
				var _300_v66 _dafny.Map
				_ = _300_v66
				_300_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v64, (_299_v65).Merge(_299_v65))
				_300_v66 = (_300_v66).Update((_296_v64).Union(_296_v64), _299_v65)
				var _301_v67 _dafny.Sequence
				_ = _301_v67
				_301_v67 = _dafny.UnicodeSeqOfUtf8Bytes("jowiusrj")
				var _302_v68 _dafny.Array
				_ = _302_v68
				var _nw54 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw54
				_302_v68 = _nw54
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_302_v68), 0))
				_ = _index54
				(_302_v68).ArraySet1((_this).F22(), (_index54).Int())
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_302_v68), 0))
				_ = _index55
				var _rhs50 _dafny.Sequence = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if _this.F23 {
						return _301_v67
					}
					return _301_v67
				})(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if _this.F23 {
						return _301_v67
					}
					return _301_v67
				})()).Cardinality()))).Uint32(), _208_v0)
				_ = _rhs50
				var _rhs51 _dafny.CodePoint = _208_v0
				_ = _rhs51
				var _rhs52 bool = _this.F23
				_ = _rhs52
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 _dafny.Array = _302_v68
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_302_v68), 0))
				_ = _lhs43
				_301_v67 = _rhs50
				_lhs41.F18 = _rhs51
				(_lhs42).ArraySet1(_rhs52, (_lhs43).Int())
				var _303_v69 _dafny.MultiSet
				_ = _303_v69
				_303_v69 = _dafny.MultiSetOf(_this.F23)
				r0 = (_dafny.IntOfUint32((_301_v67).Cardinality())).Cmp(((func() _dafny.Int {
					if (_303_v69).Contains(_this.F23) {
						return (_303_v69).Multiplicity(_this.F23)
					}
					return p0
				})()).Plus(_292_cf0)) != 0
				_292_cf0 = _292_cf0
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_302_v68), 0))
				_ = _index56
				(_302_v68).ArraySet1((_this).F22(), (_index56).Int())
			} else {
				var _304_v70 _dafny.Sequence
				_ = _304_v70
				_304_v70 = _dafny.SeqOf(p0, p0, _292_cf0, p0, _dafny.IntOfInt64(850))
				_292_cf0 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(700), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_304_v70, _304_v70)).Cardinality()))
				var _305_v71 _dafny.Array
				_ = _305_v71
				var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw55
				_305_v71 = _nw55
				_305_v71 = _305_v71
				_292_cf0 = (p0).Times(_292_cf0)
				_292_cf0 = (p0).Times(_292_cf0)
				var _306_v72 _dafny.MultiSet
				_ = _306_v72
				_306_v72 = _dafny.MultiSetOf(_this.F23, true)
				var _307_v73 _dafny.Array
				_ = _307_v73
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw56
				_307_v73 = _nw56
				var _308_v74 _dafny.Array
				_ = _308_v74
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_6
				var _nw57 _dafny.Array
				_ = _nw57
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw57 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Int = (func(_309_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_310_i6 _dafny.Int) _dafny.Int {
							return (_310_i6).Plus(_309_p0)
						}
					})(p0)
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
				_308_v74 = _nw57
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_307_v73), 0))
				_ = _index57
				(_307_v73).ArraySet1(_308_v74, (_index57).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_307_v73), 0))
				_ = _index58
				var _rhs53 _dafny.MultiSet = _306_v72
				_ = _rhs53
				var _rhs54 _dafny.Int = _292_cf0
				_ = _rhs54
				var _rhs55 D0 = _213_v5
				_ = _rhs55
				var _rhs56 _dafny.Array = _308_v74
				_ = _rhs56
				var _rhs57 _dafny.Array = _305_v71
				_ = _rhs57
				var _lhs44 _dafny.Array = _307_v73
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(517), _dafny.ArrayLen((_307_v73), 0))
				_ = _lhs45
				_306_v72 = _rhs53
				_292_cf0 = _rhs54
				_213_v5 = _rhs55
				(_lhs44).ArraySet1(_rhs56, (_lhs45).Int())
				_305_v71 = _rhs57
			}
			(globalState).F19 = (func() bool {
				if (_this).F22() {
					return _this.F23
				}
				return (_this).F22()
			})()
			var _311_v75 _dafny.Array
			_ = _311_v75
			var _nw58 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw58
			_311_v75 = _nw58
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_311_v75), 0))
			_ = _index59
			(_311_v75).ArraySet1(false, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_311_v75), 0))
			_ = _index60
			(_311_v75).ArraySet1((Companion_Default___.Fm0(_this.F23, (_this).Fm4(_dafny.IntOfUint32((_211_v3).Cardinality()), (_dafny.SetOf((_this).F22(), (_this).F22())).Cardinality(), !((_this).F22()), globalState), globalState)).Cmp((_dafny.MultiSetOf(_this.F23)).Cardinality()) < 0, (_index60).Int())
		} else {
			var _312___mcc_h7 D0 = _source2.Get_().(D0_DC3).Cf7
			_ = _312___mcc_h7
			var _313_cf7 D0 = _312___mcc_h7
			_ = _313_cf7
			var _314_v76 _dafny.Sequence
			_ = _314_v76
			_314_v76 = _dafny.UnicodeSeqOfUtf8Bytes("rfhewcmh")
			if (func() bool {
				if !(_this.F23) {
					return !(_dafny.Companion_Sequence_.Equal(_314_v76, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(754))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}((func(_315_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_316_i7 _dafny.Int) _dafny.CodePoint {
							return _315_v0
						}
					})(_208_v0)))))
				}
				return _this.F23
			})() {
				var _317_v77 _dafny.Map
				_ = _317_v77
				_317_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(945), _211_v3)
				var _318_v78 *C0
				_ = _318_v78
				var _nw59 *C0 = New_C0_()
				_ = _nw59
				_nw59.Ctor__(false, (_317_v77).Merge(_317_v77))
				_318_v78 = _nw59
				var _319_v79 _dafny.Sequence
				_ = _319_v79
				var _320_v80 bool
				_ = _320_v80
				var _out8 _dafny.Sequence
				_ = _out8
				var _out9 bool
				_ = _out9
				_out8, _out9 = (_this).M1(p0, (_this).F22(), ((_dafny.Zero).Minus(p0)).Cmp(p0) >= 0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}((func(_321_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_322_i8 _dafny.Int) _dafny.CodePoint {
						return _321_v0
					}
				})(_208_v0)))).Cardinality()), globalState)
				_319_v79 = _out8
				_320_v80 = _out9
				var _323_v81 *C0
				_ = _323_v81
				var _nw60 *C0 = New_C0_()
				_ = _nw60
				_nw60.Ctor__((_this).Fm4(p0, p0, _318_v78.F24, globalState), _317_v77)
				_323_v81 = _nw60
				var _324_v82 _dafny.MultiSet
				_ = _324_v82
				_324_v82 = _dafny.MultiSetOf(p0, Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_323_v81.F24, _320_v80, globalState), p0))
				var _rhs58 *C0 = _318_v78
				_ = _rhs58
				var _rhs59 _dafny.MultiSet = (_324_v82).Update((_213_v5).Dtor_cf2(), Companion_Default___.Abs(p0))
				_ = _rhs59
				var _rhs60 bool = !((_this).F22())
				_ = _rhs60
				var _rhs61 _dafny.CodePoint = _208_v0
				_ = _rhs61
				var _rhs62 bool = (_this).Fm4(_dafny.IntOfUint32((_314_v76).Cardinality()), p0, !(_323_v81.F24), globalState)
				_ = _rhs62
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				var _lhs48 *C1 = _this
				_ = _lhs48
				_323_v81 = _rhs58
				_324_v82 = _rhs59
				_lhs46.F19 = _rhs60
				_lhs47.F18 = _rhs61
				_lhs48.F23 = _rhs62
				_319_v79 = _314_v76
				var _325_v83 _dafny.Map
				_ = _325_v83
				_325_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _dafny.IntOfInt64(271))
				r1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v83, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(306))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_326_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_327_i9 _dafny.Int) _dafny.CodePoint {
						return _326_v0
					}
				})(_208_v0))))
			} else {
				var _328_v84 _dafny.Int
				_ = _328_v84
				_328_v84 = _dafny.IntOfInt64(-212)
				_328_v84 = _328_v84
				var _329_v85 _dafny.Array
				_ = _329_v85
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_7
				var _nw61 _dafny.Array
				_ = _nw61
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw61 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Int = (func(_330_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_331_i10 _dafny.Int) _dafny.Int {
							return (_331_i10).Times(_330_p0)
						}
					})(p0)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw61 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw61).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw61).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_329_v85 = _nw61
				var _332_v86 _dafny.Sequence
				_ = _332_v86
				_332_v86 = _dafny.SeqOf(p0)
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_329_v85), 0))
				_ = _index61
				(_329_v85).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_332_v86).Cardinality())), (_index61).Int())
				var _333_v87 _dafny.Array
				_ = _333_v87
				_333_v87 = _329_v85
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_329_v85), 0))
				_ = _index62
				var _rhs63 _dafny.Int = (_328_v84).Times((func() _dafny.Int {
					if (_this).F22() {
						return p0
					}
					return _dafny.IntOfInt64(407)
				})())
				_ = _rhs63
				var _rhs64 _dafny.Array = (_333_v87)
				_ = _rhs64
				var _rhs65 _dafny.Int = _dafny.IntOfInt64(35)
				_ = _rhs65
				var _lhs49 _dafny.Array = _329_v85
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_329_v85), 0))
				_ = _lhs50
				(_lhs49).ArraySet1(_rhs63, (_lhs50).Int())
				_329_v85 = _rhs64
				_328_v84 = _rhs65
				var _334_v88 _dafny.Map
				_ = _334_v88
				_334_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gjfewtn"), (_this).F22())
				_334_v88 = (_334_v88).Update(_314_v76, (_this).F22())
				var _335_v89 _dafny.Array
				_ = _335_v89
				var _nwElement0_12 bool = _this.F23
				_ = _nwElement0_12
				var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(4))
				_ = _nw62
				(_nw62).ArraySet1(_nwElement0_12, 0)
				(_nw62).ArraySet1(_this.F23, 1)
				(_nw62).ArraySet1(!((_this).F22()), 2)
				(_nw62).ArraySet1(_this.F23, 3)
				_335_v89 = _nw62
				var _336_v90 _dafny.Set
				_ = _336_v90
				_336_v90 = _dafny.SetOf(_213_v5)
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_335_v89), 0))
				_ = _index63
				(_335_v89).ArraySet1((_336_v90).IsSubsetOf(_336_v90), (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_335_v89), 0))
				_ = _index64
				(_335_v89).ArraySet1((_this).Fm4(_328_v84, Companion_Default___.Fm0(!((_this).F22()), _this.F23, globalState), (_this).F22(), globalState), (_index64).Int())
				(globalState).F6 = _this.F23
			}
			var _337_v91 _dafny.Map
			_ = _337_v91
			_337_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cmp((_213_v5).Dtor_cf2()) > 0, _314_v76)
			_337_v91 = _337_v91
			var _338_v92 D3
			_ = _338_v92
			_338_v92 = Companion_D3_.Create_DC6_(_314_v76)
			var _339_v93 _dafny.Sequence
			_ = _339_v93
			_339_v93 = _dafny.SeqOf(_314_v76, _314_v76, _314_v76, _dafny.UnicodeSeqOfUtf8Bytes("j"), (_338_v92).Dtor_cf10())
			if _dafny.Companion_Sequence_.IsPrefixOf(_314_v76, (_339_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-908), _dafny.IntOfUint32((_339_v93).Cardinality()))).Uint32()).(_dafny.Sequence)) {
				r0 = !((_this).F22())
				(globalState).F19 = _this.F23
				var _340_v94 _dafny.Map
				_ = _340_v94
				_340_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}((func(_341_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_342_i11 _dafny.Int) _dafny.Int {
						return _341_p0
					}
				})(p0))))
				var _343_v95 _dafny.Sequence
				_ = _343_v95
				_343_v95 = _dafny.SeqOf(p0)
				var _344_v96 _dafny.Map
				_ = _344_v96
				_344_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_340_v94).Contains((_this).F22()) {
						return (_340_v94).Get((_this).F22()).(_dafny.Sequence)
					}
					return _343_v95
				})()).Cardinality()), _211_v3)
				var _345_v97 *C0
				_ = _345_v97
				var _nw63 *C0 = New_C0_()
				_ = _nw63
				_nw63.Ctor__(_this.F23, _344_v96)
				_345_v97 = _nw63
				(globalState).F6 = _this.F23
				var _346_v98 _dafny.Array
				_ = _346_v98
				var _nw64 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw64
				_346_v98 = _nw64
				var _347_v99 _dafny.Map
				_ = _347_v99
				_347_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v98, _345_v97.F24), (_dafny.MultiSetOf(p0)).Cardinality())
				var _348_v100 _dafny.Map
				_ = _348_v100
				_348_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v98, _this.F23)
				_347_v99 = (_347_v99).Update(_348_v100, p0)
			} else {
				(globalState).F6 = (_this).F22()
				(globalState).F6 = !(((_this).F22()) == (_this.F23))
				var _349_v101 D0
				_ = _349_v101
				_349_v101 = Companion_D0_.Create_DC0_(p0)
				(globalState).F19 = (p0).Cmp((_349_v101).Dtor_cf0()) <= 0
				(globalState).F6 = (_this).F22()
				var _350_v102 _dafny.Map
				_ = _350_v102
				_350_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _211_v3)
				var _351_v103 *C0
				_ = _351_v103
				var _nw65 *C0 = New_C0_()
				_ = _nw65
				_nw65.Ctor__(false, _350_v102)
				_351_v103 = _nw65
				var _352_v104 _dafny.Sequence
				_ = _352_v104
				_352_v104 = _dafny.SeqOf(_351_v103, _351_v103, _351_v103, _351_v103)
				var _353_v105 _dafny.MultiSet
				_ = _353_v105
				_353_v105 = _dafny.MultiSetOf(p0)
				var _354_v106 _dafny.Set
				_ = _354_v106
				_354_v106 = _dafny.SetOf(p0)
				var _355_v107 _dafny.Map
				_ = _355_v107
				_355_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_351_v103.F24, (_354_v106).Cardinality())
				var _356_v108 _dafny.Array
				_ = _356_v108
				var _nwElement0_13 _dafny.Int = p0
				_ = _nwElement0_13
				var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(29))
				_ = _nw66
				(_nw66).ArraySet1(_nwElement0_13, 0)
				(_nw66).ArraySet1(_dafny.IntOfInt64(173), 1)
				(_nw66).ArraySet1(p0, 2)
				(_nw66).ArraySet1(p0, 3)
				(_nw66).ArraySet1(_dafny.IntOfInt64(978), 4)
				(_nw66).ArraySet1(p0, 5)
				(_nw66).ArraySet1(p0, 6)
				(_nw66).ArraySet1(p0, 7)
				(_nw66).ArraySet1(p0, 8)
				(_nw66).ArraySet1(p0, 9)
				(_nw66).ArraySet1(_dafny.IntOfUint32((_352_v104).Cardinality()), 10)
				(_nw66).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 11)
				(_nw66).ArraySet1((p0).Plus(p0), 12)
				(_nw66).ArraySet1(p0, 13)
				(_nw66).ArraySet1(p0, 14)
				(_nw66).ArraySet1(((_dafny.Zero).Minus(p0)).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(929))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}((func(_357_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_358_i12 _dafny.Int) _dafny.CodePoint {
						return _357_v0
					}
				})(_208_v0)))).Cardinality())), 15)
				(_nw66).ArraySet1(p0, 16)
				(_nw66).ArraySet1((_dafny.Zero).Minus(p0), 17)
				(_nw66).ArraySet1(((_353_v105).Update(p0, Companion_Default___.Abs(_dafny.IntOfInt64(-355)))).Cardinality(), 18)
				(_nw66).ArraySet1(p0, 19)
				(_nw66).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_211_v3, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-292), _dafny.IntOfUint32((_211_v3).Cardinality()))).Uint32(), _351_v103.F24)).Cardinality()), 20)
				(_nw66).ArraySet1(p0, 21)
				(_nw66).ArraySet1(p0, 22)
				(_nw66).ArraySet1(p0, 23)
				(_nw66).ArraySet1(p0, 24)
				(_nw66).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-346), p0), 25)
				(_nw66).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), 26)
				(_nw66).ArraySet1(p0, 27)
				(_nw66).ArraySet1((func() _dafny.Int {
					if (_355_v107).Contains(_351_v103.F24) {
						return (_355_v107).Get(_351_v103.F24).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_314_v76).Cardinality())
				})(), 28)
				_356_v108 = _nw66
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_356_v108), 0))
				_ = _index65
				(_356_v108).ArraySet1(p0, (_index65).Int())
				var _359_v109 _dafny.Sequence
				_ = _359_v109
				_359_v109 = _dafny.SeqOf(_dafny.IntOfInt64(-72), p0, (_dafny.Zero).Minus(p0), p0)
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_356_v108), 0))
				_ = _index66
				(_356_v108).ArraySet1(_dafny.IntOfUint32((_359_v109).Cardinality()), (_index66).Int())
				var _360_v110 _dafny.Map
				_ = _360_v110
				_360_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F22())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_356_v108), 0))
				_ = _index67
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_356_v108), 0))
				_ = _index68
				var _rhs66 bool = (func() bool {
					if (_this).F22() {
						return !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p0))).Cardinality()), p0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_314_v76, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_314_v76).Cardinality()))).Uint32(), _208_v0)).Cardinality()))
					}
					return _351_v103.F24
				})()
				_ = _rhs66
				var _rhs67 _dafny.Int = p0
				_ = _rhs67
				var _rhs68 _dafny.Int = (_360_v110).Cardinality()
				_ = _rhs68
				var _rhs69 _dafny.Sequence = (func() _dafny.Sequence {
					if (_this).Fm4((_349_v101).Dtor_cf0(), p0, _this.F23, globalState) {
						return Companion_Default___.Fm2(_dafny.MultiSetFromSeq(_314_v76), globalState)
					}
					return _dafny.SeqOf(p0)
				})()
				_ = _rhs69
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				var _lhs52 _dafny.Array = _356_v108
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(195), _dafny.ArrayLen((_356_v108), 0))
				_ = _lhs53
				var _lhs54 _dafny.Array = _356_v108
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_356_v108), 0))
				_ = _lhs55
				_lhs51.F19 = _rhs66
				(_lhs52).ArraySet1(_rhs67, (_lhs53).Int())
				(_lhs54).ArraySet1(_rhs68, (_lhs55).Int())
				_359_v109 = _rhs69
			}
			r0 = (!(true)) && ((_dafny.IntOfInt64(75)).Cmp(p0) > 0)
		}
		_208_v0 = Companion_Default___.Fm5(globalState)
		(globalState).F18 = _208_v0
		if (func() bool {
			if (_this).F22() {
				return true
			}
			return _dafny.Companion_Sequence_.IsPrefixOf(_211_v3, _211_v3)
		})() {
			var _361_v111 _dafny.Map
			_ = _361_v111
			_361_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_this).Fm4(p0, p0, (_this).F22(), globalState), false, globalState), _dafny.Companion_Sequence_.Update(_211_v3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_211_v3).Cardinality()))).Uint32(), _this.F23))
			var _362_v112 *C0
			_ = _362_v112
			var _nw67 *C0 = New_C0_()
			_ = _nw67
			_nw67.Ctor__(false, _361_v111)
			_362_v112 = _nw67
			var _363_v113 D3
			_ = _363_v113
			_363_v113 = Companion_D3_.Create_DC6_(Companion_Default___.Fm3(p0, globalState))
			var _364_v114 _dafny.Map
			_ = _364_v114
			_364_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _363_v113)
			_364_v114 = (_364_v114).Update(_362_v112.F24, _363_v113)
			(globalState).F6 = _362_v112.F24
			var _365_v115 _dafny.Array
			_ = _365_v115
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_8
			var _nw68 _dafny.Array
			_ = _nw68
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw68 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = (func(_366_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_367_i13 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(132))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}((func(_368_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_369_i14 _dafny.Int) _dafny.CodePoint {
								return _368_v0
							}
						})(_366_v0)))
					}
				})(_208_v0)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw68 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw68).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw68).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_365_v115 = _nw68
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_365_v115), 0))
			_ = _index69
			(_365_v115).ArraySet1((_363_v113).Dtor_cf10(), (_index69).Int())
			var _370_v116 _dafny.Sequence
			_ = _370_v116
			_370_v116 = _dafny.UnicodeSeqOfUtf8Bytes("ccagyoi")
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_365_v115), 0))
			_ = _index70
			(_365_v115).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_370_v116, (_363_v113).Dtor_cf10()), (_index70).Int())
			var _371_v117 *C0
			_ = _371_v117
			var _nw69 *C0 = New_C0_()
			_ = _nw69
			_nw69.Ctor__((p0).Cmp(p0) != 0, (_362_v112).F25())
			_371_v117 = _nw69
		} else {
			_208_v0 = _208_v0
			var _372_v118 _dafny.Sequence
			_ = _372_v118
			_372_v118 = _dafny.UnicodeSeqOfUtf8Bytes("sbspy")
			(_this).F23 = (_dafny.IntOfUint32((_372_v118).Cardinality())).Cmp(_dafny.IntOfUint32((_372_v118).Cardinality())) < 0
			var _373_v119 _dafny.MultiSet
			_ = _373_v119
			_373_v119 = _dafny.MultiSetOf(_372_v118, _372_v118, _dafny.Companion_Sequence_.Concatenate(_372_v118, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(469))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_374_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_375_i15 _dafny.Int) _dafny.CodePoint {
					return _374_v0
				}
			})(_208_v0)))))
			var _376_v120 _dafny.MultiSet
			_ = _376_v120
			_376_v120 = _dafny.MultiSetOf(p0)
			var _rhs70 bool = ((_dafny.IntOfInt64(582)).Minus(p0)).Cmp((_dafny.Zero).Minus(p0)) != 0
			_ = _rhs70
			var _rhs71 bool = false
			_ = _rhs71
			var _rhs72 bool = (_376_v120).IsProperSubsetOf((_376_v120).Intersection(_dafny.MultiSetOf(p0)))
			_ = _rhs72
			var _rhs73 _dafny.MultiSet = _373_v119
			_ = _rhs73
			var _lhs56 *GlobalState = globalState
			_ = _lhs56
			var _lhs57 *GlobalState = globalState
			_ = _lhs57
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			_lhs56.F6 = _rhs70
			_lhs57.F6 = _rhs71
			_lhs58.F19 = _rhs72
			_373_v119 = _rhs73
			var _377_v121 _dafny.Array
			_ = _377_v121
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw70
			_377_v121 = _nw70
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_377_v121), 0))
			_ = _index71
			(_377_v121).ArraySet1((_this).F22(), (_index71).Int())
			var _378_v122 _dafny.Array
			_ = _378_v122
			var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw71
			_378_v122 = _nw71
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_377_v121), 0))
			_ = _index72
			var _rhs74 _dafny.Array = _378_v122
			_ = _rhs74
			var _rhs75 bool = _this.F23
			_ = _rhs75
			var _lhs59 *GlobalState = globalState
			_ = _lhs59
			var _lhs60 _dafny.Array = _377_v121
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_377_v121), 0))
			_ = _lhs61
			_lhs59.F13 = _rhs74
			(_lhs60).ArraySet1(_rhs75, (_lhs61).Int())
			var _379_v123 _dafny.Sequence
			_ = _379_v123
			_379_v123 = _dafny.SeqOf(p0)
			var _380_v124 _dafny.Map
			_ = _380_v124
			_380_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_377_v121).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_377_v121), 0))).Int()).(bool), (_379_v123).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_379_v123).Cardinality()))).Uint32()).(_dafny.Int))
			_380_v124 = (_380_v124).Update(_this.F23, (func() _dafny.Int {
				if (_this).F22() {
					return p0
				}
				return p0
			})())
		}
		var _381_v125 _dafny.Map
		_ = _381_v125
		_381_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_211_v3).Cardinality()), _211_v3)
		var _382_v126 *C0
		_ = _382_v126
		var _nw72 *C0 = New_C0_()
		_ = _nw72
		_nw72.Ctor__((_this).F22(), _381_v125)
		_382_v126 = _nw72
		var _383_v127 _dafny.Sequence
		_ = _383_v127
		_383_v127 = _dafny.SeqOf(_382_v126)
		var _384_v128 _dafny.Array
		_ = _384_v128
		var _nwElement0_14 bool = (_this).F22()
		_ = _nwElement0_14
		var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(8))
		_ = _nw73
		(_nw73).ArraySet1(_nwElement0_14, 0)
		(_nw73).ArraySet1(false, 1)
		(_nw73).ArraySet1((_this).F22(), 2)
		(_nw73).ArraySet1(!(((_this).F22()) == ((_this).Fm4(_dafny.IntOfUint32((_383_v127).Cardinality()), p0, (_this).F22(), globalState))), 3)
		(_nw73).ArraySet1(_this.F23, 4)
		(_nw73).ArraySet1((_this).F22(), 5)
		(_nw73).ArraySet1(_this.F23, 6)
		(_nw73).ArraySet1((p0).Cmp(p0) == 0, 7)
		_384_v128 = _nw73
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_384_v128), 0))
		_ = _index73
		(_384_v128).ArraySet1((_this).F22(), (_index73).Int())
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_384_v128), 0))
		_ = _index74
		(_384_v128).ArraySet1(((p0).Plus((_dafny.Zero).Minus(p0))).Cmp(p0) == 0, (_index74).Int())
		r0 = (_this).F22()
		var _385_v129 _dafny.Map
		_ = _385_v129
		_385_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), p0)
		r1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_385_v129).Merge(_385_v129), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_386_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_387_i16 _dafny.Int) _dafny.CodePoint {
				return _386_v0
			}
		})(_208_v0))))
		return r0, r1
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) (_dafny.Sequence, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		(globalState).F0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _388_v0 _dafny.Array
		_ = _388_v0
		var _nw74 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(29))
		_ = _nw74
		_388_v0 = _nw74
		_388_v0 = _388_v0
		if _this.F23 {
			var _389_v1 _dafny.Map
			_ = _389_v1
			_389_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, p2)
			_389_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(true) || ((_this).F22()))
			var _390_v2 _dafny.Sequence
			_ = _390_v2
			_390_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ahbq")
			var _391_v3 _dafny.CodePoint
			_ = _391_v3
			_391_v3 = _dafny.CodePoint('h')
			var _392_v4 _dafny.Sequence
			_ = _392_v4
			_392_v4 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsPrefixOf(_390_v2, _dafny.Companion_Sequence_.Update(_390_v2, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_390_v2).Cardinality()))).Uint32(), _391_v3)))
			_392_v4 = _dafny.Companion_Sequence_.Concatenate(_392_v4, _392_v4)
			var _393_v5 *C0
			_ = _393_v5
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(789), _392_v4))
			_393_v5 = _nw75
			var _394_v6 _dafny.Int
			_ = _394_v6
			_394_v6 = _dafny.IntOfInt64(-403)
			var _395_v7 _dafny.Sequence
			_ = _395_v7
			_395_v7 = _dafny.SeqOf(_394_v6)
			var _rhs76 *C0 = _393_v5
			_ = _rhs76
			var _rhs77 _dafny.Int = Companion_Default___.SafeModuloInt(p3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_395_v7).Cardinality())))
			_ = _rhs77
			_393_v5 = _rhs76
			_394_v6 = _rhs77
			(_393_v5).F24 = _this.F23
			var _396_v9 _dafny.Array
			_ = _396_v9
			var _nwElement0_15 bool = (_this).F22()
			_ = _nwElement0_15
			var _nw76 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(25))
			_ = _nw76
			(_nw76).ArraySet1(_nwElement0_15, 0)
			(_nw76).ArraySet1(true, 1)
			(_nw76).ArraySet1(!(!(p1)), 2)
			(_nw76).ArraySet1(_393_v5.F24, 3)
			(_nw76).ArraySet1(p1, 4)
			(_nw76).ArraySet1(true, 5)
			(_nw76).ArraySet1(false, 6)
			(_nw76).ArraySet1(p2, 7)
			(_nw76).ArraySet1(p2, 8)
			(_nw76).ArraySet1(_this.F23, 9)
			(_nw76).ArraySet1(false, 10)
			(_nw76).ArraySet1(p1, 11)
			(_nw76).ArraySet1(_393_v5.F24, 12)
			(_nw76).ArraySet1(p2, 13)
			(_nw76).ArraySet1((_this).F22(), 14)
			(_nw76).ArraySet1((_this).F22(), 15)
			(_nw76).ArraySet1(_this.F23, 16)
			(_nw76).ArraySet1((_this).F22(), 17)
			(_nw76).ArraySet1(_393_v5.F24, 18)
			(_nw76).ArraySet1(p1, 19)
			(_nw76).ArraySet1(true, 20)
			(_nw76).ArraySet1(p1, 21)
			(_nw76).ArraySet1(p2, 22)
			(_nw76).ArraySet1(p1, 23)
			(_nw76).ArraySet1((_this).Fm4((func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(36), _dafny.IntOfInt64(816))); ; {
					_compr_4, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _397_v8 _dafny.Int
					_397_v8 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(36)).Cmp(_397_v8) <= 0) && ((_397_v8).Cmp(_dafny.IntOfInt64(816)) < 0) {
						_coll4.Add((_397_v8).Minus(_dafny.IntOfUint32((_390_v2).Cardinality())), true)
					}
				}
				return _coll4.ToMap()
			}()).Cardinality(), p3, !(p2), globalState), 24)
			_396_v9 = _nw76
			var _398_v10 _dafny.MultiSet
			_ = _398_v10
			_398_v10 = _dafny.MultiSetOf(_396_v9)
			_394_v6 = ((_398_v10).Union(_398_v10)).Cardinality()
		} else {
			var _399_v11 _dafny.Int
			_ = _399_v11
			_399_v11 = _dafny.IntOfInt64(478)
			var _400_v12 _dafny.MultiSet
			_ = _400_v12
			_400_v12 = _dafny.MultiSetOf(p3)
			var _401_v13 _dafny.MultiSet
			_ = _401_v13
			_401_v13 = _dafny.MultiSetOf(_this.F23)
			var _402_v14 _dafny.Map
			_ = _402_v14
			_402_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_v13, _this.F23)
			var _403_v15 _dafny.Map
			_ = _403_v15
			_403_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _399_v11)
			var _404_v17 _dafny.Sequence
			_ = _404_v17
			_404_v17 = _dafny.SeqOf((_403_v15).Cardinality(), (func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-789), _dafny.IntOfInt64(579))); ; {
					_compr_5, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _405_v16 _dafny.Int
					_405_v16 = interface{}(_compr_5).(_dafny.Int)
					if ((_dafny.IntOfInt64(-789)).Cmp(_405_v16) <= 0) && ((_405_v16).Cmp(_dafny.IntOfInt64(579)) < 0) {
						_coll5.Add((_405_v16).Times(_dafny.IntOfInt64(-899)), p0)
					}
				}
				return _coll5.ToMap()
			}()).Cardinality(), p0, _399_v11)
			var _406_v18 _dafny.Sequence
			_ = _406_v18
			_406_v18 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)).Cardinality(), (_402_v14).Cardinality(), (_404_v17).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_404_v17).Cardinality()))).Uint32()).(_dafny.Int))
			_399_v11 = ((_400_v12).Cardinality()).Plus((p3).Times(_dafny.IntOfUint32((_406_v18).Cardinality())))
			var _407_v19 _dafny.Set
			_ = _407_v19
			_407_v19 = _dafny.SetOf(p0, p3, p3, _399_v11, (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_399_v11, p0)))
			_407_v19 = _dafny.SetOf(p3)
			var _408_v20 D0
			_ = _408_v20
			_408_v20 = Companion_D0_.Create_DC2_((_this).F22(), _dafny.IntOfInt64(-817), _399_v11, p1)
			var _409_v21 _dafny.Sequence
			_ = _409_v21
			_409_v21 = _dafny.SeqOf(_408_v20, _408_v20, _408_v20, _408_v20)
			var _410_v23 _dafny.Map
			_ = _410_v23
			_410_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), p2)
			var _411_v24 _dafny.Sequence
			_ = _411_v24
			_411_v24 = _dafny.SeqOf(true)
			var _rhs78 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(999))).Uint32(), func(coer21 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_412_v20 D0) func(_dafny.Int) D0 {
				return func(_413_i0 _dafny.Int) D0 {
					return _412_v20
				}
			})(_408_v20)))
			_ = _rhs78
			var _rhs79 bool = p2
			_ = _rhs79
			var _rhs80 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if false {
					return _dafny.SeqOf((_this).Fm4(((func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(263), _dafny.IntOfInt64(685))); ; {
							_compr_6, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _414_v22 _dafny.Int
							_414_v22 = interface{}(_compr_6).(_dafny.Int)
							if ((_dafny.IntOfInt64(263)).Cmp(_414_v22) <= 0) && ((_414_v22).Cmp(_dafny.IntOfInt64(685)) < 0) {
								_coll6.Add(Companion_Default___.SafeModuloInt(_414_v22, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_399_v11, _399_v11)).Cardinality()), p3)
							}
						}
						return _coll6.ToMap()
					}()).Update(p0, (_410_v23).Cardinality())).Cardinality(), _399_v11, true, globalState), p2, p2, false)
				}
				return _411_v24
			})(), _dafny.Companion_Sequence_.Concatenate(_411_v24, _dafny.SeqOf((_this).F22())))).Cardinality())
			_ = _rhs80
			var _lhs62 *C1 = _this
			_ = _lhs62
			_409_v21 = _rhs78
			_lhs62.F23 = _rhs79
			_399_v11 = _rhs80
			var _415_v25 _dafny.Map
			_ = _415_v25
			_415_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm10(p3, _399_v11, globalState))
			var _416_v26 *C0
			_ = _416_v26
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(p2, _415_v25)
			_416_v26 = _nw77
			var _417_v27 _dafny.Array
			_ = _417_v27
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_9
			var _nw78 _dafny.Array
			_ = _nw78
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw78 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Int = (func(_418_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_419_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_419_i1, _418_p3)
					}
				})(p3)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw78 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw78).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw78).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_417_v27 = _nw78
			var _420_v28 _dafny.Array
			_ = _420_v28
			var _nwElement0_16 _dafny.Array = _417_v27
			_ = _nwElement0_16
			var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(13))
			_ = _nw79
			(_nw79).ArraySet1(_nwElement0_16, 0)
			(_nw79).ArraySet1(_417_v27, 1)
			(_nw79).ArraySet1(_417_v27, 2)
			(_nw79).ArraySet1(_417_v27, 3)
			(_nw79).ArraySet1(_417_v27, 4)
			(_nw79).ArraySet1(_417_v27, 5)
			(_nw79).ArraySet1(_417_v27, 6)
			(_nw79).ArraySet1(_417_v27, 7)
			(_nw79).ArraySet1(_417_v27, 8)
			(_nw79).ArraySet1(_417_v27, 9)
			(_nw79).ArraySet1(_417_v27, 10)
			(_nw79).ArraySet1(_417_v27, 11)
			(_nw79).ArraySet1(_417_v27, 12)
			_420_v28 = _nw79
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_420_v28), 0))
			_ = _index75
			(_420_v28).ArraySet1(_417_v27, (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.ArrayLen((_420_v28), 0))
			_ = _index76
			(_420_v28).ArraySet1(_417_v27, (_index76).Int())
		}
		var _421_v29 _dafny.Array
		_ = _421_v29
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_10
		var _nw80 _dafny.Array
		_ = _nw80
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw80 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) D0 = (func(_422_p0 _dafny.Int) func(_dafny.Int) D0 {
				return func(_423_i2 _dafny.Int) D0 {
					return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_422_p0))
				}
			})(p0)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw80 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw80).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw80).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_421_v29 = _nw80
		var _424_v30 _dafny.Map
		_ = _424_v30
		_424_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), (_dafny.Zero).Minus(p0))
		var _425_v31 _dafny.Map
		_ = _425_v31
		_425_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_424_v30, _421_v29)
		var _426_v32 _dafny.Sequence
		_ = _426_v32
		_426_v32 = _dafny.SeqOf(_421_v29, _421_v29, _421_v29)
		var _427_v33 _dafny.Sequence
		_ = _427_v33
		_427_v33 = _dafny.SeqOf(_this.F23, p2)
		var _428_v34 _dafny.Map
		_ = _428_v34
		_428_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _427_v33)
		var _429_v35 T0
		_ = _429_v35
		var _nw81 *C0 = New_C0_()
		_ = _nw81
		_nw81.Ctor__(p1, _428_v34)
		_429_v35 = _nw81
		var _430_v36 _dafny.Map
		_ = _430_v36
		_430_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _429_v35)
		_421_v29 = (func() _dafny.Array {
			if (_425_v31).Contains(_424_v30) {
				return (_425_v31).Get(_424_v30).(_dafny.Array)
			}
			return (_426_v32).Select((Companion_Default___.SafeIndex((_430_v36).Cardinality(), _dafny.IntOfUint32((_426_v32).Cardinality()))).Uint32()).(_dafny.Array)
		})()
		var _431_v37 _dafny.Array
		_ = _431_v37
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw82
		_431_v37 = _nw82
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_431_v37), 0))); ; {
			_guard_loop_1, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _432_i3 _dafny.Int
			_432_i3 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_432_i3).Sign() != -1) && ((_432_i3).Cmp(_dafny.ArrayLen((_431_v37), 0)) < 0)) {
				(_431_v37).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_427_v33), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(true)), _dafny.SeqOf(_427_v33))), (_432_i3).Int())
			}
		}
		_431_v37 = _431_v37
		var _433_v38 _dafny.Sequence
		_ = _433_v38
		_433_v38 = _dafny.UnicodeSeqOfUtf8Bytes("crax")
		r0 = _433_v38
		r1 = (p2) || (p1)
		return r0, r1
	}
}
func (_this *C1) F22() bool {
	{
		return _this._f22
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
