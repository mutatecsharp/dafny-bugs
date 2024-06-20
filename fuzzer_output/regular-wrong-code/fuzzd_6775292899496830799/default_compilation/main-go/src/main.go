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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(true), true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ydsmrsl")).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, false), _dafny.IntOfInt64(730))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), _dafny.IntOfInt64(17))))).Cardinality()
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) bool {
	return ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ikpsqij")).Cardinality())).Minus(_dafny.IntOfInt64(740))).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(819))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("ufwlwbkag")
	})), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xw")))).Cardinality())) == 0
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (Companion_D6_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Dtor_cf22()
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(490), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(252))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(920)
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(true, true)).Cardinality()), _dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(308), _dafny.IntOfInt64(811), _dafny.IntOfInt64(692), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), (_dafny.SetOf(_dafny.IntOfInt64(985), _dafny.IntOfInt64(-50), _dafny.IntOfInt64(-950))).Cardinality())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, false, false, false), _dafny.SeqOf(!(true), false, false))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(150), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-678))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))).Cardinality()), _dafny.IntOfInt64(-739), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(5))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(640), _dafny.IntOfInt64(226), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()))).Cardinality()))).Difference(_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality(), _dafny.IntOfInt64(-783)))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(191), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yik"), false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gfvif"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(311))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	})))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D3 = Companion_D3_.Create_DC4_(true)
	_ = _source0
	if _source0.Is_DC4() {
		var _4___mcc_h0 bool = _source0.Get_().(D3_DC4).Cf4
		_ = _4___mcc_h0
		var _5_cf4 bool = _4___mcc_h0
		_ = _5_cf4
		if _5_cf4 {
			return _dafny.CodePoint('x')
		} else {
			return _dafny.CodePoint('b')
		}
	} else if _source0.Is_DC3() {
		var _6___mcc_h1 _dafny.Array = _source0.Get_().(D3_DC3).Cf3
		_ = _6___mcc_h1
		var _7_cf3 _dafny.Array = _6___mcc_h1
		_ = _7_cf3
		return _dafny.CodePoint('j')
	} else {
		var _8___mcc_h2 D3 = _source0.Get_().(D3_DC5).Cf5
		_ = _8___mcc_h2
		var _9_cf5 D3 = _8___mcc_h2
		_ = _9_cf5
		return _dafny.CodePoint('x')
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(false), false)).Union(_dafny.MultiSetOf(!(true)))).Difference(_dafny.MultiSetOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(false), false, true)).Difference((_dafny.SetOf(!(false), true, true)).Intersection(_dafny.SetOf(!(true))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, globalState *GlobalState) {
	var _10_v0 *C0
	_ = _10_v0
	var _nw0 *C0 = New_C0_()
	_ = _nw0
	_nw0.Ctor__()
	_10_v0 = _nw0
	var _11_v1 bool
	_ = _11_v1
	_11_v1 = true
	var _12_i0 _dafny.Int
	_ = _12_i0
	_12_i0 = _dafny.Zero
	{
		for _11_v1 {
			{
				if (_12_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_12_i0 = (_12_i0).Plus(_dafny.One)
				var _13_v2 _dafny.Int
				_ = _13_v2
				_13_v2 = _dafny.IntOfInt64(620)
				_13_v2 = _13_v2
				(globalState).F2 = !(false) || (_11_v1)
				(globalState).F2 = (Companion_Default___.SafeDivisionInt(_13_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("umrnxguo")).Cardinality()))).Cmp(Companion_Default___.SafeDivisionInt(_13_v2, _13_v2)) != 0
				var _14_v3 _dafny.Set
				_ = _14_v3
				_14_v3 = _dafny.SetOf(_13_v2)
				var _15_v4 _dafny.Set
				_ = _15_v4
				_15_v4 = _14_v3
				var _16_v5 _dafny.Map
				_ = _16_v5
				_16_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v2, _11_v1)
				if !((_15_v4).Equals(_dafny.SetOf(_dafny.IntOfInt64(608), _13_v2, _13_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_17_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}))).Cardinality()), (_16_v5).Cardinality()))) {
					var _18_v6 *C0
					_ = _18_v6
					var _nw1 *C0 = New_C0_()
					_ = _nw1
					_nw1.Ctor__()
					_18_v6 = _nw1
					_18_v6 = _10_v0
					var _19_v7 D3
					_ = _19_v7
					_19_v7 = Companion_D3_.Create_DC4_(_11_v1)
					var _20_v8 _dafny.Sequence
					_ = _20_v8
					_20_v8 = _dafny.SeqOf((_19_v7).Dtor_cf4(), _11_v1, _11_v1)
					var _21_v9 _dafny.Map
					_ = _21_v9
					_21_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _13_v2)
					var _22_v10 _dafny.Sequence
					_ = _22_v10
					_22_v10 = _dafny.SeqOf(_21_v9)
					var _rhs0 _dafny.Int = _13_v2
					_ = _rhs0
					var _rhs1 bool = (func() bool {
						if false {
							return _11_v1
						}
						return (_20_v8).Select((Companion_Default___.SafeIndex(_13_v2, _dafny.IntOfUint32((_20_v8).Cardinality()))).Uint32()).(bool)
					})()
					_ = _rhs1
					var _rhs2 bool = !(_21_v9).Equals((_22_v10).Select((Companion_Default___.SafeIndex(_13_v2, _dafny.IntOfUint32((_22_v10).Cardinality()))).Uint32()).(_dafny.Map))
					_ = _rhs2
					var _rhs3 bool = _11_v1
					_ = _rhs3
					var _rhs4 bool = _11_v1
					_ = _rhs4
					var _lhs0 *GlobalState = globalState
					_ = _lhs0
					var _lhs1 *GlobalState = globalState
					_ = _lhs1
					_13_v2 = _rhs0
					_lhs0.F2 = _rhs1
					_lhs1.F2 = _rhs2
					_11_v1 = _rhs3
					_11_v1 = _rhs4
					var _23_v11 _dafny.Sequence
					_ = _23_v11
					_23_v11 = _dafny.UnicodeSeqOfUtf8Bytes("o")
					var _24_v12 _dafny.MultiSet
					_ = _24_v12
					_24_v12 = _dafny.MultiSetOf(true)
					var _25_v13 _dafny.Int
					_ = _25_v13
					_25_v13 = _13_v2
					var _26_v14 _dafny.Sequence
					_ = _26_v14
					_26_v14 = _dafny.SeqOf(_13_v2)
					var _27_v15 _dafny.Array
					_ = _27_v15
					var _nwElement0_0 _dafny.Int = _13_v2
					_ = _nwElement0_0
					var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(23))
					_ = _nw2
					(_nw2).ArraySet1(_nwElement0_0, 0)
					(_nw2).ArraySet1(_13_v2, 1)
					(_nw2).ArraySet1(_13_v2, 2)
					(_nw2).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(917))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg5 _dafny.Int) interface{} {
							return coer5(arg5)
						}
					}(func(_28_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('s')
					}))).Cardinality())).Times(_13_v2), 3)
					(_nw2).ArraySet1(((Companion_D5_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v11, _11_v1))).Dtor_cf16()).Cardinality(), 4)
					(_nw2).ArraySet1(_13_v2, 5)
					(_nw2).ArraySet1(((_dafny.MultiSetOf(_11_v1)).Cardinality()), 6)
					(_nw2).ArraySet1((_dafny.Zero).Minus(_13_v2), 7)
					(_nw2).ArraySet1(_dafny.IntOfInt64(775), 8)
					(_nw2).ArraySet1((_13_v2).Times(_13_v2), 9)
					(_nw2).ArraySet1((func() _dafny.Int {
						if (_24_v12).Contains(!(_11_v1)) {
							return (_24_v12).Multiplicity(!(_11_v1))
						}
						return _13_v2
					})(), 10)
					(_nw2).ArraySet1(Companion_Default___.SafeDivisionInt(_13_v2, _13_v2), 11)
					(_nw2).ArraySet1(_13_v2, 12)
					(_nw2).ArraySet1(_13_v2, 13)
					(_nw2).ArraySet1(Companion_Default___.Fm0(_25_v13, globalState), 14)
					(_nw2).ArraySet1((_26_v14).Select((Companion_Default___.SafeIndex(_13_v2, _dafny.IntOfUint32((_26_v14).Cardinality()))).Uint32()).(_dafny.Int), 15)
					(_nw2).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfInt64(991), globalState), 16)
					(_nw2).ArraySet1(_13_v2, 17)
					(_nw2).ArraySet1(_13_v2, 18)
					(_nw2).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_26_v14).Cardinality()), _dafny.IntOfUint32((_23_v11).Cardinality())), 19)
					(_nw2).ArraySet1((_13_v2).Minus(_dafny.IntOfUint32((_23_v11).Cardinality())), 20)
					(_nw2).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ym"), _23_v11)).Cardinality()), 21)
					(_nw2).ArraySet1(_13_v2, 22)
					_27_v15 = _nw2
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_27_v15), 0))
					_ = _index0
					(_27_v15).ArraySet1((_13_v2).Minus(_13_v2), (_index0).Int())
					var _29_v16 _dafny.CodePoint
					_ = _29_v16
					_29_v16 = _dafny.CodePoint('m')
					var _30_v17 _dafny.Array
					_ = _30_v17
					var _nwElement0_1 _dafny.CodePoint = Companion_Default___.Fm8(!((_18_v6).Fm1(globalState)), _11_v1, globalState)
					_ = _nwElement0_1
					var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
					_ = _nw3
					(_nw3).ArraySet1CodePoint(_nwElement0_1, 0)
					(_nw3).ArraySet1CodePoint((_23_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.IntOfUint32((_23_v11).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
					(_nw3).ArraySet1CodePoint(_29_v16, 2)
					(_nw3).ArraySet1CodePoint(_29_v16, 3)
					(_nw3).ArraySet1CodePoint(_dafny.CodePoint('s'), 4)
					(_nw3).ArraySet1CodePoint(_29_v16, 5)
					(_nw3).ArraySet1CodePoint(_29_v16, 6)
					(_nw3).ArraySet1CodePoint(_dafny.CodePoint('r'), 7)
					(_nw3).ArraySet1CodePoint(_29_v16, 8)
					(_nw3).ArraySet1CodePoint(_29_v16, 9)
					(_nw3).ArraySet1CodePoint(_29_v16, 10)
					(_nw3).ArraySet1CodePoint(_dafny.CodePoint('i'), 11)
					(_nw3).ArraySet1CodePoint(_dafny.CodePoint('h'), 12)
					(_nw3).ArraySet1CodePoint(_29_v16, 13)
					(_nw3).ArraySet1CodePoint(_29_v16, 14)
					(_nw3).ArraySet1CodePoint(_29_v16, 15)
					(_nw3).ArraySet1CodePoint(_29_v16, 16)
					(_nw3).ArraySet1CodePoint(_29_v16, 17)
					(_nw3).ArraySet1CodePoint((_23_v11).Select((Companion_Default___.SafeIndex(_13_v2, _dafny.IntOfUint32((_23_v11).Cardinality()))).Uint32()).(_dafny.CodePoint), 18)
					(_nw3).ArraySet1CodePoint(_29_v16, 19)
					(_nw3).ArraySet1CodePoint(_29_v16, 20)
					(_nw3).ArraySet1CodePoint(_29_v16, 21)
					(_nw3).ArraySet1CodePoint(_29_v16, 22)
					(_nw3).ArraySet1CodePoint(_29_v16, 23)
					(_nw3).ArraySet1CodePoint(_29_v16, 24)
					(_nw3).ArraySet1CodePoint(_29_v16, 25)
					_30_v17 = _nw3
					var _31_v18 D4
					_ = _31_v18
					_31_v18 = Companion_D4_.Create_DC7_(_23_v11, _13_v2, _30_v17, (func() bool {
						if true {
							return _11_v1
						}
						return _11_v1
					})(), _13_v2)
					var _pat_let_tv0 = _13_v2
					_ = _pat_let_tv0
					var _pat_let_tv1 = _30_v17
					_ = _pat_let_tv1
					var _pat_let_tv2 = _11_v1
					_ = _pat_let_tv2
					var _pat_let_tv3 = _13_v2
					_ = _pat_let_tv3
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_27_v15), 0))
					_ = _index1
					var _rhs5 _dafny.Int = _dafny.IntOfInt64(-509)
					_ = _rhs5
					var _rhs6 D4 = func(_pat_let0_0 D4) D4 {
						return func(_32_dt__update__tmp_h0 D4) D4 {
							return func(_pat_let1_0 _dafny.Int) D4 {
								return func(_33_dt__update_hcf8_h0 _dafny.Int) D4 {
									return func(_pat_let2_0 _dafny.Array) D4 {
										return func(_34_dt__update_hcf9_h0 _dafny.Array) D4 {
											return func(_pat_let3_0 bool) D4 {
												return func(_35_dt__update_hcf10_h0 bool) D4 {
													return func(_pat_let4_0 _dafny.Int) D4 {
														return func(_36_dt__update_hcf11_h0 _dafny.Int) D4 {
															return Companion_D4_.Create_DC7_((_32_dt__update__tmp_h0).Dtor_cf7(), _33_dt__update_hcf8_h0, _34_dt__update_hcf9_h0, _35_dt__update_hcf10_h0, _36_dt__update_hcf11_h0)
														}(_pat_let4_0)
													}(_pat_let_tv3)
												}(_pat_let3_0)
											}(_pat_let_tv2)
										}(_pat_let2_0)
									}(_pat_let_tv1)
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_31_v18)
					_ = _rhs6
					var _rhs7 bool = !((func() bool {
						if (_16_v5).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _20_v8)).Cardinality())) {
							return (_16_v5).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _20_v8)).Cardinality())).(bool)
						}
						return _11_v1
					})())
					_ = _rhs7
					var _lhs2 _dafny.Array = _27_v15
					_ = _lhs2
					var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_27_v15), 0))
					_ = _lhs3
					var _lhs4 *GlobalState = globalState
					_ = _lhs4
					(_lhs2).ArraySet1(_rhs5, (_lhs3).Int())
					_31_v18 = _rhs6
					_lhs4.F2 = _rhs7
					_13_v2 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_27_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_27_v15), 0))).Int()).(_dafny.Int)), ((_dafny.Zero).Minus(_13_v2)).Minus(_13_v2))
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_27_v15), 0))
					_ = _index2
					(_27_v15).ArraySet1((_25_v13), (_index2).Int())
				} else {
					var _37_v19 _dafny.Sequence
					_ = _37_v19
					_37_v19 = _dafny.SeqOf(_13_v2)
					var _38_v20 _dafny.Sequence
					_ = _38_v20
					_38_v20 = _dafny.SeqOf((_dafny.Zero).Minus((_37_v19).Select((Companion_Default___.SafeIndex(_13_v2, _dafny.IntOfUint32((_37_v19).Cardinality()))).Uint32()).(_dafny.Int)), _13_v2)
					var _39_v21 _dafny.Map
					_ = _39_v21
					_39_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _37_v19)
					var _40_v22 _dafny.Sequence
					_ = _40_v22
					_40_v22 = _dafny.SeqOf(_39_v21)
					_13_v2 = (((_39_v21).Merge(_39_v21)).Merge(((_40_v22).Select((Companion_Default___.SafeIndex(_13_v2, _dafny.IntOfUint32((_40_v22).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_39_v21))).Cardinality()
					_13_v2 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v2, _10_v0)).Cardinality()).Minus(_dafny.IntOfInt64(504))).Minus(_13_v2)
					_13_v2 = _13_v2
					var _41_v23 *C0
					_ = _41_v23
					var _nw4 *C0 = New_C0_()
					_ = _nw4
					_nw4.Ctor__()
					_41_v23 = _nw4
					(globalState).F2 = _11_v1
				}
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _42_v24 _dafny.Int
	_ = _42_v24
	_42_v24 = _dafny.IntOfInt64(892)
	var _43_v25 _dafny.Array
	_ = _43_v25
	var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
	_ = _nw5
	_43_v25 = _nw5
	var _44_v26 _dafny.MultiSet
	_ = _44_v26
	_44_v26 = _dafny.MultiSetOf(_11_v1)
	var _45_v27 D4
	_ = _45_v27
	_45_v27 = Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("frcgkbmfp"), _42_v24, _43_v25, _11_v1, (func() _dafny.Int {
		if (_44_v26).Contains(_11_v1) {
			return (_44_v26).Multiplicity(_11_v1)
		}
		return (_dafny.SetOf(!(_11_v1))).Cardinality()
	})())
	var _46_v28 _dafny.Sequence
	_ = _46_v28
	_46_v28 = _dafny.UnicodeSeqOfUtf8Bytes("xurjjvlua")
	var _47_v29 _dafny.MultiSet
	_ = _47_v29
	_47_v29 = _dafny.MultiSetOf(_45_v27, Companion_D4_.Create_DC7_(_46_v28, _42_v24, _43_v25, _11_v1, _42_v24), _45_v27)
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_46_v28, _46_v28)).Cardinality())
	_ = _hi0
	for _48_i3 := (_dafny.Zero).Minus((func() _dafny.Int {
		if _11_v1 {
			return _42_v24
		}
		return (_47_v29).Cardinality()
	})()); _48_i3.Cmp(_hi0) < 0; _48_i3 = _48_i3.Plus(_dafny.One) {
		_42_v24 = _dafny.IntOfUint32((_46_v28).Cardinality())
		var _49_v30 _dafny.Map
		_ = _49_v30
		_49_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_i3, _42_v24)
		var _50_v31 _dafny.Array
		_ = _50_v31
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw6
		_50_v31 = _nw6
		var _51_v32 _dafny.Sequence
		_ = _51_v32
		_51_v32 = _dafny.SeqOf(_50_v31)
		_49_v30 = (_49_v30).Update(Companion_Default___.Fm0(_48_i3, globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_51_v32, (Companion_Default___.SafeIndex(_48_i3, _dafny.IntOfUint32((_51_v32).Cardinality()))).Uint32(), _50_v31)).Cardinality()))
		var _52_v33 _dafny.CodePoint
		_ = _52_v33
		_52_v33 = _dafny.CodePoint('l')
		var _53_v34 _dafny.Array
		_ = _53_v34
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_0
		var _nw7 _dafny.Array
		_ = _nw7
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw7 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.MultiSet = (func(_54_v26 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_55_i4 _dafny.Int) _dafny.MultiSet {
					return _54_v26
				}
			})(_44_v26)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw7 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw7).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw7).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_53_v34 = _nw7
		var _rhs8 _dafny.CodePoint = _52_v33
		_ = _rhs8
		var _rhs9 _dafny.Array = _53_v34
		_ = _rhs9
		_52_v33 = _rhs8
		_53_v34 = _rhs9
		var _56_v35 _dafny.Int
		_ = _56_v35
		_56_v35 = _48_i3
		_42_v24 = Companion_Default___.SafeDivisionInt(_42_v24, (Companion_Default___.Fm0(_56_v35, globalState)).Minus(Companion_Default___.Fm0(_56_v35, globalState)))
	}
	(globalState).F2 = (_10_v0).Fm1(globalState)
	var _57_v36 _dafny.Map
	_ = _57_v36
	_57_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v1, _11_v1)
	var _58_v37 _dafny.Sequence
	_ = _58_v37
	_58_v37 = _dafny.SeqOf(_11_v1, !(_11_v1))
	var _59_v38 _dafny.Array
	_ = _59_v38
	var _nwElement0_2 bool = !((func() bool {
		if (_57_v36).Contains(true) {
			return (_57_v36).Get(true).(bool)
		}
		return _11_v1
	})())
	_ = _nwElement0_2
	var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(23))
	_ = _nw8
	(_nw8).ArraySet1(_nwElement0_2, 0)
	(_nw8).ArraySet1(_11_v1, 1)
	(_nw8).ArraySet1(_11_v1, 2)
	(_nw8).ArraySet1(_11_v1, 3)
	(_nw8).ArraySet1(_11_v1, 4)
	(_nw8).ArraySet1(_11_v1, 5)
	(_nw8).ArraySet1(_11_v1, 6)
	(_nw8).ArraySet1(_11_v1, 7)
	(_nw8).ArraySet1(_11_v1, 8)
	(_nw8).ArraySet1(Companion_Default___.Fm2(globalState), 9)
	(_nw8).ArraySet1(_11_v1, 10)
	(_nw8).ArraySet1(true, 11)
	(_nw8).ArraySet1(_11_v1, 12)
	(_nw8).ArraySet1(_11_v1, 13)
	(_nw8).ArraySet1(_11_v1, 14)
	(_nw8).ArraySet1(_11_v1, 15)
	(_nw8).ArraySet1(_11_v1, 16)
	(_nw8).ArraySet1(_11_v1, 17)
	(_nw8).ArraySet1(_11_v1, 18)
	(_nw8).ArraySet1(true, 19)
	(_nw8).ArraySet1((_58_v37).Select((Companion_Default___.SafeIndex(_42_v24, _dafny.IntOfUint32((_58_v37).Cardinality()))).Uint32()).(bool), 20)
	(_nw8).ArraySet1(true, 21)
	(_nw8).ArraySet1(_11_v1, 22)
	_59_v38 = _nw8
	var _60_v39 _dafny.Map
	_ = _60_v39
	_60_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v38, !(_11_v1))
	var _61_v40 _dafny.Map
	_ = _61_v40
	_61_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_42_v24, _11_v1)
	var _62_v41 _dafny.Array
	_ = _62_v41
	var _nwElement0_3 bool = _11_v1
	_ = _nwElement0_3
	var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(20))
	_ = _nw9
	(_nw9).ArraySet1(_nwElement0_3, 0)
	(_nw9).ArraySet1(_11_v1, 1)
	(_nw9).ArraySet1(_11_v1, 2)
	(_nw9).ArraySet1(_11_v1, 3)
	(_nw9).ArraySet1(_11_v1, 4)
	(_nw9).ArraySet1(!((func() bool {
		if _11_v1 {
			return _11_v1
		}
		return _11_v1
	})()), 5)
	(_nw9).ArraySet1(_11_v1, 6)
	(_nw9).ArraySet1(_11_v1, 7)
	(_nw9).ArraySet1(_11_v1, 8)
	(_nw9).ArraySet1(_11_v1, 9)
	(_nw9).ArraySet1((func() bool {
		if _11_v1 {
			return _11_v1
		}
		return _11_v1
	})(), 10)
	(_nw9).ArraySet1(!((func() bool {
		if (_60_v39).Contains(_59_v38) {
			return (_60_v39).Get(_59_v38).(bool)
		}
		return _11_v1
	})()), 11)
	(_nw9).ArraySet1(_11_v1, 12)
	(_nw9).ArraySet1(_11_v1, 13)
	(_nw9).ArraySet1((_58_v37).Select((Companion_Default___.SafeIndex(_42_v24, _dafny.IntOfUint32((_58_v37).Cardinality()))).Uint32()).(bool), 14)
	(_nw9).ArraySet1(_dafny.Companion_Sequence_.Equal(_46_v28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(716))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_63_i5 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))), 15)
	(_nw9).ArraySet1(_11_v1, 16)
	(_nw9).ArraySet1(_11_v1, 17)
	(_nw9).ArraySet1(((_61_v40).Cardinality()).Cmp(_42_v24) > 0, 18)
	(_nw9).ArraySet1((false) && (_11_v1), 19)
	_62_v41 = _nw9
	_59_v38 = _59_v38
	var _hi1 _dafny.Int = (_dafny.IntOfInt64(715)).Minus(_42_v24)
	_ = _hi1
	for _64_i6 := (func() _dafny.Int {
		if _11_v1 {
			return _42_v24
		}
		return _42_v24
	})(); _64_i6.Cmp(_hi1) < 0; _64_i6 = _64_i6.Plus(_dafny.One) {
		var _65_v42 _dafny.Array
		_ = _65_v42
		var _nwElement0_4 *C0 = _10_v0
		_ = _nwElement0_4
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.One)
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_4, 0)
		_65_v42 = _nw10
		_65_v42 = _65_v42
		_11_v1 = _11_v1
		var _66_v43 _dafny.Sequence
		_ = _66_v43
		_66_v43 = _dafny.SeqOf(_65_v42, _65_v42, _65_v42)
		_66_v43 = _66_v43
		(globalState).F2 = Companion_Default___.Fm2(globalState)
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _67_v0 _dafny.Sequence
	_ = _67_v0
	_67_v0 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("smr"))
	var _68_v1 _dafny.Int
	_ = _68_v1
	_68_v1 = _dafny.IntOfInt64(456)
	var _69_v2 _dafny.Int
	_ = _69_v2
	_69_v2 = _68_v1
	var _70_v3 _dafny.Set
	_ = _70_v3
	_70_v3 = _dafny.SetOf(_dafny.IntOfUint32((_67_v0).Cardinality()), _68_v1, (_69_v2), _68_v1, _68_v1)
	var _71_globalState *GlobalState
	_ = _71_globalState
	var _nw11 *GlobalState = New_GlobalState_()
	_ = _nw11
	_nw11.Ctor__(_dafny.CodePoint('w'), _70_v3, true)
	_71_globalState = _nw11
	var _72_v4 _dafny.Array
	_ = _72_v4
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(24)
	_ = _len0_1
	var _nw12 _dafny.Array
	_ = _nw12
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw12 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Sequence = func(_73_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(779))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_74_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw12 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw12).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw12).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_72_v4 = _nw12
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))
	_ = _index3
	(_72_v4).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uwc"), (_index3).Int())
	var _75_v5 _dafny.Sequence
	_ = _75_v5
	_75_v5 = _dafny.UnicodeSeqOfUtf8Bytes("lyuag")
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))
	_ = _index4
	(_72_v4).ArraySet1(_75_v5, (_index4).Int())
	var _source1 _dafny.Int = Companion_Default___.Fm0(_69_v2, _71_globalState)
	_ = _source1
	var _76___mcc_h0 _dafny.Int = _source1
	_ = _76___mcc_h0
	var _77_cf0 _dafny.Int = _76___mcc_h0
	_ = _77_cf0
	var _78_v6 _dafny.MultiSet
	_ = _78_v6
	_78_v6 = _dafny.MultiSetOf(_75_v5, (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("lvyvvfan"), _dafny.UnicodeSeqOfUtf8Bytes("e"), _dafny.UnicodeSeqOfUtf8Bytes("p"))
	_78_v6 = _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(635))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_79_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	}))))
	var _80_v7 *C0
	_ = _80_v7
	var _nw13 *C0 = New_C0_()
	_ = _nw13
	_nw13.Ctor__()
	_80_v7 = _nw13
	var _81_v8 bool
	_ = _81_v8
	_81_v8 = false
	var _rhs10 bool = _81_v8
	_ = _rhs10
	var _rhs11 _dafny.Int = _77_cf0
	_ = _rhs11
	var _lhs5 *GlobalState = _71_globalState
	_ = _lhs5
	_lhs5.F2 = _rhs10
	_77_cf0 = _rhs11
	var _82_v9 _dafny.CodePoint
	_ = _82_v9
	_82_v9 = _dafny.CodePoint('y')
	var _83_v10 _dafny.MultiSet
	_ = _83_v10
	_83_v10 = _dafny.MultiSetOf(_82_v9, _82_v9)
	_81_v8 = !(_83_v10).Equals(_dafny.MultiSetFromSeq((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)))
	if true {
		var _84_v11 _dafny.Array
		_ = _84_v11
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_2
		var _nw14 _dafny.Array
		_ = _nw14
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw14 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_85_v1 _dafny.Int) func(_dafny.Int) bool {
				return func(_86_i3 _dafny.Int) bool {
					return (_85_v1).Cmp(_85_v1) == 0
				}
			})(_68_v1)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw14 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw14).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw14).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_84_v11 = _nw14
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))
		_ = _index5
		(_84_v11).ArraySet1((_68_v1).Cmp(_68_v1) != 0, (_index5).Int())
		var _87_v12 bool
		_ = _87_v12
		_87_v12 = true
		var _88_v13 _dafny.Sequence
		_ = _88_v13
		_88_v13 = _dafny.SeqOf(_87_v12)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))
		_ = _index6
		(_84_v11).ArraySet1(!_dafny.Companion_Sequence_.Contains(_88_v13, Companion_Default___.Fm2(_71_globalState)), (_index6).Int())
		var _source2 _dafny.Int = _69_v2
		_ = _source2
		var _89___mcc_h1 _dafny.Int = _source2
		_ = _89___mcc_h1
		var _90_cf0 _dafny.Int = _89___mcc_h1
		_ = _90_cf0
		var _91_v14 _dafny.Array
		_ = _91_v14
		var _nwElement0_5 _dafny.Int = _dafny.IntOfInt64(-632)
		_ = _nwElement0_5
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.One)
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_5, 0)
		_91_v14 = _nw15
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_91_v14), 0))
		_ = _index7
		(_91_v14).ArraySet1(_dafny.IntOfInt64(560), (_index7).Int())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_91_v14), 0))
		_ = _index8
		(_91_v14).ArraySet1(Companion_Default___.Fm0(_69_v2, _71_globalState), (_index8).Int())
		var _92_v15 _dafny.Map
		_ = _92_v15
		_92_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_69_v2, _71_globalState), _84_v11)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_91_v14), 0))
		_ = _index9
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_91_v14), 0))
		_ = _index10
		var _rhs12 _dafny.Array = (func() _dafny.Array {
			if (_92_v15).Contains(_68_v1) {
				return (_92_v15).Get(_68_v1).(_dafny.Array)
			}
			return _84_v11
		})()
		_ = _rhs12
		var _rhs13 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_90_cf0, Companion_Default___.Fm0(_69_v2, _71_globalState)), _90_cf0))
		_ = _rhs13
		var _rhs14 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ivmmsgbn")).Cardinality()), _90_cf0)
		_ = _rhs14
		var _rhs15 _dafny.Int = (Companion_Default___.Fm0(_69_v2, _71_globalState)).Times((_dafny.Zero).Minus(_68_v1))
		_ = _rhs15
		var _lhs6 _dafny.Array = _91_v14
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_91_v14), 0))
		_ = _lhs7
		var _lhs8 _dafny.Array = _91_v14
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_91_v14), 0))
		_ = _lhs9
		_84_v11 = _rhs12
		(_lhs6).ArraySet1(_rhs13, (_lhs7).Int())
		_90_cf0 = _rhs14
		(_lhs8).ArraySet1(_rhs15, (_lhs9).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))
		_ = _index11
		(_72_v4).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_75_v5, (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)), (_index11).Int())
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))
		_ = _index12
		(_84_v11).ArraySet1(false, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_91_v14), 0))
		_ = _index13
		(_91_v14).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(842))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_93_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})))).Cardinality()), (_index13).Int())
		var _94_v16 _dafny.Array
		_ = _94_v16
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_3
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_95_v13 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_96_i5 _dafny.Int) _dafny.Sequence {
					return _95_v13
				}
			})(_88_v13)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw16 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw16).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw16).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_94_v16 = _nw16
		var _97_v17 _dafny.Map
		_ = _97_v17
		_97_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_84_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))).Int()).(bool), (_84_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))).Int()).(bool))
		var _98_v18 _dafny.Sequence
		_ = _98_v18
		_98_v18 = _dafny.SeqOf(_68_v1)
		var _99_v19 _dafny.Sequence
		_ = _99_v19
		_99_v19 = _dafny.SeqOf(_98_v18, _98_v18, _dafny.SeqOf(_68_v1, (_98_v18).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_98_v18).Cardinality()))).Uint32()).(_dafny.Int), _68_v1, _68_v1, _68_v1), _98_v18)
		var _100_v20 _dafny.Array
		_ = _100_v20
		var _nwElement0_6 _dafny.Map = _97_v17
		_ = _nwElement0_6
		var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw17
		(_nw17).ArraySet1(_nwElement0_6, 0)
		(_nw17).ArraySet1((_97_v17).Merge(_97_v17), 1)
		(_nw17).ArraySet1(_97_v17, 2)
		(_nw17).ArraySet1(_97_v17, 3)
		(_nw17).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v12, _87_v12), 4)
		(_nw17).ArraySet1(_97_v17, 5)
		(_nw17).ArraySet1(_97_v17, 6)
		(_nw17).ArraySet1((_97_v17).Update(true, (_84_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))).Int()).(bool)), 7)
		(_nw17).ArraySet1(_97_v17, 8)
		(_nw17).ArraySet1(_97_v17, 9)
		(_nw17).ArraySet1(Companion_Default___.Fm3((_dafny.MultiSetFromSeq(_99_v19)).Cardinality(), _71_globalState), 10)
		(_nw17).ArraySet1(_97_v17, 11)
		(_nw17).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_84_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))).Int()).(bool), _87_v12), 12)
		_100_v20 = _nw17
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_100_v20), 0))
		_ = _index14
		(_100_v20).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v12, (_84_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))).Int()).(bool))).Merge(_97_v17), (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_100_v20), 0))
		_ = _index15
		var _rhs16 _dafny.Array = _94_v16
		_ = _rhs16
		var _rhs17 _dafny.Int = _68_v1
		_ = _rhs17
		var _rhs18 _dafny.Map = _97_v17
		_ = _rhs18
		var _lhs10 _dafny.Array = _100_v20
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_100_v20), 0))
		_ = _lhs11
		_94_v16 = _rhs16
		_68_v1 = _rhs17
		(_lhs10).ArraySet1(_rhs18, (_lhs11).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_84_v11), 0))
		_ = _index16
		(_84_v11).ArraySet1(!(_87_v12), (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))
		_ = _index17
		(_72_v4).ArraySet1((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), (_index17).Int())
	} else {
		var _101_v21 _dafny.Array
		_ = _101_v21
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
		_ = _nw18
		_101_v21 = _nw18
		var _102_v22 _dafny.Array
		_ = _102_v22
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_4
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = (func(_103_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_104_i6 _dafny.Int) _dafny.Int {
					return (_104_i6).Plus(_103_v1)
				}
			})(_68_v1)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw19 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw19).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw19).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_102_v22 = _nw19
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_101_v21), 0))
		_ = _index18
		(_101_v21).ArraySet1(_102_v22, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(73), _dafny.ArrayLen((_101_v21), 0))
		_ = _index19
		(_101_v21).ArraySet1(_102_v22, (_index19).Int())
		var _105_v23 bool
		_ = _105_v23
		_105_v23 = true
		if !(_105_v23) {
			var _106_v24 _dafny.Array
			_ = _106_v24
			var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw20
			_106_v24 = _nw20
			var _107_v25 _dafny.Map
			_ = _107_v25
			_107_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v23, !(_105_v23))
			var _108_v26 _dafny.CodePoint
			_ = _108_v26
			_108_v26 = _dafny.CodePoint('r')
			var _109_v27 _dafny.Map
			_ = _109_v27
			_109_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm4((func() bool {
				if (_107_v25).Contains(!(_105_v23)) {
					return (_107_v25).Get(!(_105_v23)).(bool)
				}
				return _105_v23
			})(), _105_v23, _108_v26, _71_globalState)).Cardinality(), _dafny.IntOfInt64(423))
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_106_v24), 0))
			_ = _index20
			(_106_v24).ArraySet1(_109_v27, (_index20).Int())
			var _110_v28 _dafny.Map
			_ = _110_v28
			_110_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, Companion_Default___.Fm2(_71_globalState))
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_106_v24), 0))
			_ = _index21
			(_106_v24).ArraySet1((func() _dafny.Map {
				if (func() bool {
					if (_110_v28).Contains(_68_v1) {
						return (_110_v28).Get(_68_v1).(bool)
					}
					return Companion_Default___.Fm2(_71_globalState)
				})() {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _68_v1)
				}
				return (_109_v27).Merge(_109_v27)
			})(), (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_106_v24), 0))
			_ = _index22
			(_106_v24).ArraySet1(((_106_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_106_v24), 0))).Int()).(_dafny.Map)).Update(_68_v1, _68_v1), (_index22).Int())
			_69_v2 = _69_v2
			var _111_v29 *C0
			_ = _111_v29
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__()
			_111_v29 = _nw21
			var _112_v30 _dafny.Map
			_ = _112_v30
			_112_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_105_v23), _111_v29)
			var _113_v31 _dafny.Set
			_ = _113_v31
			_113_v31 = _dafny.SetOf((func() *C0 {
				if (_112_v30).Contains(!(_105_v23)) {
					return (_112_v30).Get(!(_105_v23)).(*C0)
				}
				return _111_v29
			})())
			var _114_v32 _dafny.Sequence
			_ = _114_v32
			_114_v32 = _dafny.SeqOf(_68_v1)
			var _rhs19 _dafny.Set = ((_113_v31).Intersection(_113_v31)).Intersection(_113_v31)
			_ = _rhs19
			var _rhs20 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_68_v1), _dafny.IntOfInt64(340))
			_ = _rhs20
			var _rhs21 _dafny.Int = (_68_v1).Times(Companion_Default___.SafeModuloInt((_dafny.MultiSetFromSeq(_114_v32)).Cardinality(), _68_v1))
			_ = _rhs21
			_113_v31 = _rhs19
			_68_v1 = _rhs20
			_68_v1 = _rhs21
			var _115_v33 _dafny.MultiSet
			_ = _115_v33
			_115_v33 = _dafny.MultiSetOf(_68_v1, _68_v1)
			var _116_v34 _dafny.Map
			_ = _116_v34
			_116_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v3, _dafny.MultiSetFromSeq(_dafny.SeqOf(_68_v1, _68_v1)))
			var _117_v35 _dafny.Set
			_ = _117_v35
			_117_v35 = _70_v3
			var _118_v36 _dafny.Map
			_ = _118_v36
			_118_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _114_v32)
			var _119_v37 _dafny.Map
			_ = _119_v37
			_119_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_108_v26, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_75_v5, (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_75_v5).Cardinality()))).Uint32(), _108_v26)).Cardinality()))
			var _120_v38 _dafny.Sequence
			_ = _120_v38
			_120_v38 = _dafny.SeqOf(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if (_118_v36).Contains((_dafny.SetOf((_119_v37).Cardinality(), _68_v1, _68_v1)).Cardinality()) {
					return (_118_v36).Get((_dafny.SetOf((_119_v37).Cardinality(), _68_v1, _68_v1)).Cardinality()).(_dafny.Sequence)
				}
				return _114_v32
			})()), _115_v33, _115_v33)
			var _121_v39 _dafny.Array
			_ = _121_v39
			var _nwElement0_7 _dafny.MultiSet = (_115_v33).Update(((_115_v33).Update(_68_v1, Companion_Default___.Abs(Companion_Default___.Fm0(_68_v1, _71_globalState)))).Cardinality(), Companion_Default___.Abs((_dafny.Zero).Minus(_68_v1)))
			_ = _nwElement0_7
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(20))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_7, 0)
			(_nw22).ArraySet1(_115_v33, 1)
			(_nw22).ArraySet1(_115_v33, 2)
			(_nw22).ArraySet1(_115_v33, 3)
			(_nw22).ArraySet1(_115_v33, 4)
			(_nw22).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqneu")).Cardinality())), 5)
			(_nw22).ArraySet1(_115_v33, 6)
			(_nw22).ArraySet1(_115_v33, 7)
			(_nw22).ArraySet1(_115_v33, 8)
			(_nw22).ArraySet1(_115_v33, 9)
			(_nw22).ArraySet1(_115_v33, 10)
			(_nw22).ArraySet1(_dafny.MultiSetOf(_68_v1, _dafny.IntOfInt64(695), (_dafny.SetOf(_105_v23)).Cardinality()), 11)
			(_nw22).ArraySet1(_115_v33, 12)
			(_nw22).ArraySet1(_115_v33, 13)
			(_nw22).ArraySet1(_115_v33, 14)
			(_nw22).ArraySet1(_115_v33, 15)
			(_nw22).ArraySet1((func() _dafny.MultiSet {
				if (_116_v34).Contains((_117_v35)) {
					return (_116_v34).Get((_117_v35)).(_dafny.MultiSet)
				}
				return (_120_v38).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_120_v38).Cardinality()))).Uint32()).(_dafny.MultiSet)
			})(), 16)
			(_nw22).ArraySet1(_115_v33, 17)
			(_nw22).ArraySet1(_115_v33, 18)
			(_nw22).ArraySet1(_115_v33, 19)
			_121_v39 = _nw22
			Companion_Default___.M0(_121_v39, _71_globalState)
		} else {
			var _122_v40 _dafny.Map
			_ = _122_v40
			_122_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_105_v23)).Cardinality(), _105_v23)
			var _123_v41 _dafny.Sequence
			_ = _123_v41
			_123_v41 = _dafny.SeqOf(_105_v23, _105_v23)
			var _124_v42 _dafny.Map
			_ = _124_v42
			_124_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_123_v41, _105_v23)
			var _125_v43 _dafny.Sequence
			_ = _125_v43
			_125_v43 = _dafny.SeqOf(_105_v23, _105_v23, _105_v23)
			var _126_v44 _dafny.Map
			_ = _126_v44
			_126_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_105_v23), (_125_v43))
			var _127_v45 _dafny.CodePoint
			_ = _127_v45
			_127_v45 = _dafny.CodePoint('n')
			var _128_v46 _dafny.Array
			_ = _128_v46
			var _nwElement0_8 bool = _105_v23
			_ = _nwElement0_8
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(21))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_8, 0)
			(_nw23).ArraySet1(_105_v23, 1)
			(_nw23).ArraySet1(!(true), 2)
			(_nw23).ArraySet1(_105_v23, 3)
			(_nw23).ArraySet1((func() bool {
				if _105_v23 {
					return false
				}
				return true
			})(), 4)
			(_nw23).ArraySet1(!(_105_v23), 5)
			(_nw23).ArraySet1((func() bool {
				if (_122_v40).Contains(_68_v1) {
					return (_122_v40).Get(_68_v1).(bool)
				}
				return !(_105_v23)
			})(), 6)
			(_nw23).ArraySet1(_105_v23, 7)
			(_nw23).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((func() bool {
				if (_124_v42).Contains((func() _dafny.Sequence {
					if (_126_v44).Contains(_105_v23) {
						return (_126_v44).Get(_105_v23).(_dafny.Sequence)
					}
					return _123_v41
				})()) {
					return (_124_v42).Get((func() _dafny.Sequence {
						if (_126_v44).Contains(_105_v23) {
							return (_126_v44).Get(_105_v23).(_dafny.Sequence)
						}
						return _123_v41
					})()).(bool)
				}
				return _105_v23
			})(), false, _105_v23), Companion_Default___.Fm5(_68_v1, _71_globalState)), 8)
			(_nw23).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(261))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_129_v45 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_130_i7 _dafny.Int) _dafny.CodePoint {
					return _129_v45
				}
			})(_127_v45))), _127_v45), 9)
			(_nw23).ArraySet1((_68_v1).Cmp(_68_v1) <= 0, 10)
			(_nw23).ArraySet1(_105_v23, 11)
			(_nw23).ArraySet1(_105_v23, 12)
			(_nw23).ArraySet1((_68_v1).Cmp(_68_v1) == 0, 13)
			(_nw23).ArraySet1(_105_v23, 14)
			(_nw23).ArraySet1(_105_v23, 15)
			(_nw23).ArraySet1(true, 16)
			(_nw23).ArraySet1(_105_v23, 17)
			(_nw23).ArraySet1((func() bool {
				if !(true) {
					return false
				}
				return _105_v23
			})(), 18)
			(_nw23).ArraySet1(_105_v23, 19)
			(_nw23).ArraySet1(_105_v23, 20)
			_128_v46 = _nw23
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_128_v46), 0))
			_ = _index23
			(_128_v46).ArraySet1(_105_v23, (_index23).Int())
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_128_v46), 0))
			_ = _index24
			(_128_v46).ArraySet1(Companion_Default___.Fm2(_71_globalState), (_index24).Int())
			var _131_v47 _dafny.Array
			_ = _131_v47
			var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
			_ = _nw24
			_131_v47 = _nw24
			var _132_v48 _dafny.Set
			_ = _132_v48
			_132_v48 = _70_v3
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_131_v47), 0))
			_ = _index25
			(_131_v47).ArraySet1(_132_v48, (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_131_v47), 0))
			_ = _index26
			(_131_v47).ArraySet1(_70_v3, (_index26).Int())
			var _133_v49 _dafny.Set
			_ = _133_v49
			_133_v49 = _dafny.SetOf((_128_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_128_v46), 0))).Int()).(bool))
			var _134_v50 _dafny.MultiSet
			_ = _134_v50
			_134_v50 = _dafny.MultiSetOf(_133_v49)
			_68_v1 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_68_v1), (func() _dafny.Int {
				if (_134_v50).Contains(_133_v49) {
					return (_134_v50).Multiplicity(_133_v49)
				}
				return _68_v1
			})())
			var _135_v51 _dafny.Array
			_ = _135_v51
			var _nwElement0_9 _dafny.Set = _70_v3
			_ = _nwElement0_9
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(7))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_9, 0)
			(_nw25).ArraySet1(_70_v3, 1)
			(_nw25).ArraySet1(_70_v3, 2)
			(_nw25).ArraySet1(_70_v3, 3)
			(_nw25).ArraySet1(_70_v3, 4)
			(_nw25).ArraySet1(_70_v3, 5)
			(_nw25).ArraySet1(_70_v3, 6)
			_135_v51 = _nw25
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_135_v51), 0))
			_ = _index27
			(_135_v51).ArraySet1(_70_v3, (_index27).Int())
			var _136_v52 _dafny.Sequence
			_ = _136_v52
			_136_v52 = _dafny.SeqOf(_68_v1)
			var _137_v53 _dafny.Sequence
			_ = _137_v53
			_137_v53 = _dafny.SeqOf(_136_v52, _136_v52, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v23, (_128_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_128_v46), 0))).Int()).(bool))).Cardinality()))
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_135_v51), 0))
			_ = _index28
			(_135_v51).ArraySet1(func() _dafny.Set {
				var _coll0 = _dafny.NewBuilder()
				_ = _coll0
				for _iter0 := _dafny.Iterate(((func() _dafny.Sequence {
					if _105_v23 {
						return (_137_v53).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_137_v53).Cardinality()))).Uint32()).(_dafny.Sequence)
					}
					return _136_v52
				})()).Elements()); ; {
					_compr_0, _ok0 := _iter0()
					if !_ok0 {
						break
					}
					var _138_v54 _dafny.Int
					_138_v54 = interface{}(_compr_0).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
						if _105_v23 {
							return (_137_v53).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_137_v53).Cardinality()))).Uint32()).(_dafny.Sequence)
						}
						return _136_v52
					})(), _138_v54) {
						_coll0.Add((_138_v54).Minus((_dafny.IntOfInt64(729)).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(917))).Cardinality(), _dafny.IntOfInt64(139), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ewhrvia")).Cardinality())), (_dafny.SetOf(!(true))).Cardinality())).Cardinality())).Cardinality()))))
					}
				}
				return _coll0.ToSet()
			}(), (_index28).Int())
			_105_v23 = _105_v23
		}
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))
		_ = _index29
		(_102_v22).ArraySet1(_68_v1, (_index29).Int())
		var _139_v55 _dafny.Array
		_ = _139_v55
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
		_ = _nw26
		_139_v55 = _nw26
		var _140_v56 _dafny.CodePoint
		_ = _140_v56
		_140_v56 = _dafny.CodePoint('q')
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_139_v55), 0))
		_ = _index30
		(_139_v55).ArraySet1CodePoint(_140_v56, (_index30).Int())
		var _141_v57 _dafny.Array
		_ = _141_v57
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_5
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) bool = (func(_142_v5 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_143_i8 _dafny.Int) bool {
					return !_dafny.Companion_Sequence_.Equal(_142_v5, _142_v5)
				}
			})(_75_v5)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw27 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw27).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw27).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_141_v57 = _nw27
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_141_v57), 0))
		_ = _index31
		(_141_v57).ArraySet1(_105_v23, (_index31).Int())
		var _144_v58 _dafny.Map
		_ = _144_v58
		_144_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if _105_v23 {
				return _105_v23
			}
			return _105_v23
		})(), _140_v56)
		var _145_v59 _dafny.Sequence
		_ = _145_v59
		_145_v59 = _dafny.SeqOf(_105_v23, _105_v23)
		var _146_v60 _dafny.Sequence
		_ = _146_v60
		_146_v60 = _dafny.SeqOf(_dafny.IntOfUint32((_145_v59).Cardinality()))
		var _147_v61 _dafny.Sequence
		_ = _147_v61
		_147_v61 = _dafny.SeqOf(_146_v60)
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))
		_ = _index32
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_139_v55), 0))
		_ = _index33
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_141_v57), 0))
		_ = _index34
		var _rhs22 _dafny.Int = (_68_v1).Minus(_68_v1)
		_ = _rhs22
		var _rhs23 _dafny.Array = _101_v21
		_ = _rhs23
		var _rhs24 _dafny.CodePoint = (func() _dafny.CodePoint {
			if (_144_v58).Contains(_105_v23) {
				return (_144_v58).Get(_105_v23).(_dafny.CodePoint)
			}
			return _140_v56
		})()
		_ = _rhs24
		var _rhs25 _dafny.Array = _141_v57
		_ = _rhs25
		var _rhs26 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
			if _105_v23 {
				return _147_v61
			}
			return _dafny.SeqOf(_146_v60, _146_v60)
		})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(563))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_148_v60 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_149_i9 _dafny.Int) _dafny.Sequence {
				return _148_v60
			}
		})(_146_v60)))))
		_ = _rhs26
		var _lhs12 _dafny.Array = _102_v22
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))
		_ = _lhs13
		var _lhs14 _dafny.Array = _139_v55
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_139_v55), 0))
		_ = _lhs15
		var _lhs16 _dafny.Array = _141_v57
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_141_v57), 0))
		_ = _lhs17
		(_lhs12).ArraySet1(_rhs22, (_lhs13).Int())
		_101_v21 = _rhs23
		(_lhs14).ArraySet1CodePoint(_rhs24, (_lhs15).Int())
		_141_v57 = _rhs25
		(_lhs16).ArraySet1(_rhs26, (_lhs17).Int())
		var _150_v62 _dafny.MultiSet
		_ = _150_v62
		_150_v62 = _dafny.MultiSetOf(_dafny.IntOfUint32((_145_v59).Cardinality()), _dafny.IntOfUint32(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()), _68_v1)
		var _151_v63 _dafny.Map
		_ = _151_v63
		_151_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_105_v23, (_141_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_141_v57), 0))).Int()).(bool))
		var _152_v64 *C0
		_ = _152_v64
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__()
		_152_v64 = _nw28
		var _153_v65 _dafny.Map
		_ = _153_v65
		_153_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v64, _dafny.IntOfInt64(960))
		var _154_v66 _dafny.Array
		_ = _154_v66
		var _nwElement0_10 _dafny.MultiSet = _150_v62
		_ = _nwElement0_10
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(26))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_10, 0)
		(_nw29).ArraySet1(_dafny.MultiSetOf((_102_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))).Int()).(_dafny.Int)), 1)
		(_nw29).ArraySet1(_150_v62, 2)
		(_nw29).ArraySet1((_150_v62).Update((_102_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_102_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))).Int()).(_dafny.Int))), 3)
		(_nw29).ArraySet1(Companion_Default___.Fm6((_141_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_141_v57), 0))).Int()).(bool), _71_globalState), 4)
		(_nw29).ArraySet1(_dafny.MultiSetOf((_151_v63).Cardinality(), (_102_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))).Int()).(_dafny.Int)), 5)
		(_nw29).ArraySet1(_150_v62, 6)
		(_nw29).ArraySet1(_150_v62, 7)
		(_nw29).ArraySet1(_150_v62, 8)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_68_v1, (_153_v65).Cardinality()), 9)
		(_nw29).ArraySet1(Companion_Default___.Fm6(_105_v23, _71_globalState), 10)
		(_nw29).ArraySet1(_150_v62, 11)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_68_v1, (_102_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_102_v22), 0))).Int()).(_dafny.Int)), 12)
		(_nw29).ArraySet1(_150_v62, 13)
		(_nw29).ArraySet1(_150_v62, 14)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_68_v1, _68_v1), 15)
		(_nw29).ArraySet1(_150_v62, 16)
		(_nw29).ArraySet1(_150_v62, 17)
		(_nw29).ArraySet1(_150_v62, 18)
		(_nw29).ArraySet1(Companion_Default___.Fm6((_141_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_141_v57), 0))).Int()).(bool), _71_globalState), 19)
		(_nw29).ArraySet1(_150_v62, 20)
		(_nw29).ArraySet1(_150_v62, 21)
		(_nw29).ArraySet1(_150_v62, 22)
		(_nw29).ArraySet1(_150_v62, 23)
		(_nw29).ArraySet1(_150_v62, 24)
		(_nw29).ArraySet1(_150_v62, 25)
		_154_v66 = _nw29
		Companion_Default___.M0(_154_v66, _71_globalState)
		var _155_v67 _dafny.MultiSet
		_ = _155_v67
		_155_v67 = _dafny.MultiSetOf(_69_v2)
		var _156_v68 _dafny.Array
		_ = _156_v68
		var _nwElement0_11 _dafny.MultiSet = _155_v67
		_ = _nwElement0_11
		var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(6))
		_ = _nw30
		(_nw30).ArraySet1(_nwElement0_11, 0)
		(_nw30).ArraySet1(_155_v67, 1)
		(_nw30).ArraySet1((_dafny.MultiSetOf(_69_v2)).Union(_155_v67), 2)
		(_nw30).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfInt64(925)), 3)
		(_nw30).ArraySet1(_155_v67, 4)
		(_nw30).ArraySet1(_155_v67, 5)
		_156_v68 = _nw30
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_156_v68), 0))
		_ = _index35
		(_156_v68).ArraySet1(_dafny.MultiSetOf(_68_v1, Companion_Default___.Fm0(_69_v2, _71_globalState)), (_index35).Int())
		var _157_v69 _dafny.Array
		_ = _157_v69
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
		_ = _nw31
		_157_v69 = _nw31
		var _158_v70 D3
		_ = _158_v70
		_158_v70 = Companion_D3_.Create_DC3_(_157_v69)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_156_v68), 0))
		_ = _index36
		var _rhs27 bool = _105_v23
		_ = _rhs27
		var _rhs28 _dafny.MultiSet = _155_v67
		_ = _rhs28
		var _rhs29 _dafny.Array = (_158_v70).Dtor_cf3()
		_ = _rhs29
		var _lhs18 _dafny.Array = _156_v68
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(339), _dafny.ArrayLen((_156_v68), 0))
		_ = _lhs19
		_105_v23 = _rhs27
		(_lhs18).ArraySet1(_rhs28, (_lhs19).Int())
		_157_v69 = _rhs29
	}
	var _159_v71 *C0
	_ = _159_v71
	var _nw32 *C0 = New_C0_()
	_ = _nw32
	_nw32.Ctor__()
	_159_v71 = _nw32
	_159_v71 = _159_v71
	var _160_v72 bool
	_ = _160_v72
	_160_v72 = true
	var _161_v73 _dafny.Map
	_ = _161_v73
	_161_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v72, _160_v72)
	var _162_v74 _dafny.Array
	_ = _162_v74
	var _nwElement0_12 _dafny.Map = (_161_v73).Merge(_161_v73)
	_ = _nwElement0_12
	var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(5))
	_ = _nw33
	(_nw33).ArraySet1(_nwElement0_12, 0)
	(_nw33).ArraySet1(_161_v73, 1)
	(_nw33).ArraySet1(_161_v73, 2)
	(_nw33).ArraySet1(Companion_Default___.Fm3(_68_v1, _71_globalState), 3)
	(_nw33).ArraySet1(Companion_Default___.Fm3(_68_v1, _71_globalState), 4)
	_162_v74 = _nw33
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_162_v74), 0))
	_ = _index37
	(_162_v74).ArraySet1(_161_v73, (_index37).Int())
	var _163_v75 _dafny.Array
	_ = _163_v75
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_6
	var _nw34 _dafny.Array
	_ = _nw34
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw34 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = (func(_164_v72 bool) func(_dafny.Int) bool {
			return func(_165_i10 _dafny.Int) bool {
				return _164_v72
			}
		})(_160_v72)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw34 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw34).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw34).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_163_v75 = _nw34
	var _166_v76 _dafny.Sequence
	_ = _166_v76
	_166_v76 = _dafny.SeqOf(_160_v72)
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
	_ = _index38
	(_163_v75).ArraySet1((_166_v76).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_166_v76).Cardinality()))).Uint32()).(bool), (_index38).Int())
	var _167_v77 _dafny.Sequence
	_ = _167_v77
	_167_v77 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v72, _160_v72), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v72, true))
	var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_162_v74), 0))
	_ = _index39
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
	_ = _index40
	var _rhs30 _dafny.Map = ((_161_v73).Merge((_167_v77).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_68_v1), _dafny.IntOfUint32((_167_v77).Cardinality()))).Uint32()).(_dafny.Map))).Merge(_161_v73)
	_ = _rhs30
	var _rhs31 bool = _160_v72
	_ = _rhs31
	var _lhs20 _dafny.Array = _162_v74
	_ = _lhs20
	var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_162_v74), 0))
	_ = _lhs21
	var _lhs22 _dafny.Array = _163_v75
	_ = _lhs22
	var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
	_ = _lhs23
	(_lhs20).ArraySet1(_rhs30, (_lhs21).Int())
	(_lhs22).ArraySet1(_rhs31, (_lhs23).Int())
	if (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool) {
		var _168_v78 _dafny.Map
		_ = _168_v78
		_168_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v76, _dafny.IntOfInt64(142))
		_68_v1 = (_68_v1).Times((func() _dafny.Int {
			if _160_v72 {
				return (_168_v78).Cardinality()
			}
			return _68_v1
		})())
		_163_v75 = _163_v75
		_68_v1 = _68_v1
		(_71_globalState).F2 = (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)
		var _169_v79 _dafny.Map
		_ = _169_v79
		_169_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_69_v2, _71_globalState), (_68_v1).Minus((_161_v73).Cardinality()))
		var _170_v80 _dafny.Set
		_ = _170_v80
		_170_v80 = _dafny.SetOf(_160_v72, _160_v72, true, _160_v72, (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
		_169_v79 = (_169_v79).Update(_68_v1, (func() _dafny.Int {
			if (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool) {
				return (_170_v80).Cardinality()
			}
			return _dafny.IntOfInt64(-598)
		})())
	} else {
		_159_v71 = _159_v71
		var _171_v81 _dafny.Array
		_ = _171_v81
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_7
		var _nw35 _dafny.Array
		_ = _nw35
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw35 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Int = (func(_172_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_173_i11 _dafny.Int) _dafny.Int {
					return (_173_i11).Times(_172_v1)
				}
			})(_68_v1)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw35 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw35).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw35).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_171_v81 = _nw35
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
		_ = _index41
		(_171_v81).ArraySet1((_68_v1).Minus(_68_v1), (_index41).Int())
		var _174_v82 _dafny.CodePoint
		_ = _174_v82
		_174_v82 = _dafny.CodePoint('m')
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
		_ = _index42
		(_171_v81).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm7(_68_v1, (_159_v71).Fm1(_71_globalState), !((_68_v1).Cmp(_68_v1) < 0), _174_v82, _71_globalState)).Cardinality()), (_index42).Int())
		var _175_v83 _dafny.MultiSet
		_ = _175_v83
		_175_v83 = _dafny.MultiSetOf(_160_v72, _160_v72)
		var _176_v84 _dafny.Sequence
		_ = _176_v84
		_176_v84 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_166_v76), _dafny.MultiSetFromSeq(_166_v76))
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
		_ = _index43
		(_163_v75).ArraySet1(((_176_v84).Select((Companion_Default___.SafeIndex((_171_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_176_v84).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_175_v83), (_index43).Int())
		var _177_v85 _dafny.Array
		_ = _177_v85
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
		_ = _nw36
		_177_v85 = _nw36
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_177_v85), 0))
		_ = _index44
		(_177_v85).ArraySet1CodePoint(_174_v82, (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_177_v85), 0))
		_ = _index45
		(_177_v85).ArraySet1CodePoint(_174_v82, (_index45).Int())
		if (_68_v1).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_67_v0, _67_v0)).Cardinality())) != 0 {
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_177_v85), 0))
			_ = _index46
			(_177_v85).ArraySet1CodePoint(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_171_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index46).Int())
			_160_v72 = _160_v72
			var _178_v86 _dafny.Array
			_ = _178_v86
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(17))
			_ = _nw37
			_178_v86 = _nw37
			var _179_v87 _dafny.Array
			_ = _179_v87
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw38
			_179_v87 = _nw38
			var _180_v88 D3
			_ = _180_v88
			_180_v88 = Companion_D3_.Create_DC3_(_179_v87)
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_178_v86), 0))
			_ = _index47
			(_178_v86).ArraySet1(_180_v88, (_index47).Int())
			var _181_v89 _dafny.Sequence
			_ = _181_v89
			_181_v89 = _dafny.SeqOf((_171_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(_69_v2, _71_globalState))
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _index48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _index49
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_178_v86), 0))
			_ = _index50
			var _rhs32 _dafny.Int = _68_v1
			_ = _rhs32
			var _rhs33 bool = (_dafny.IntOfInt64(697)).Cmp(_dafny.IntOfUint32((_181_v89).Cardinality())) == 0
			_ = _rhs33
			var _rhs34 _dafny.Int = Companion_Default___.SafeDivisionInt(((_171_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))).Int()).(_dafny.Int)).Minus(_68_v1), (_69_v2))
			_ = _rhs34
			var _rhs35 _dafny.Array = _171_v81
			_ = _rhs35
			var _rhs36 D3 = _180_v88
			_ = _rhs36
			var _lhs24 _dafny.Array = _171_v81
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _lhs25
			var _lhs26 *GlobalState = _71_globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _171_v81
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _lhs28
			var _lhs29 _dafny.Array = _178_v86
			_ = _lhs29
			var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_178_v86), 0))
			_ = _lhs30
			(_lhs24).ArraySet1(_rhs32, (_lhs25).Int())
			_lhs26.F2 = _rhs33
			(_lhs27).ArraySet1(_rhs34, (_lhs28).Int())
			_171_v81 = _rhs35
			(_lhs29).ArraySet1(_rhs36, (_lhs30).Int())
			_68_v1 = _68_v1
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _index51
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _index52
			var _rhs37 _dafny.Int = (_171_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))).Int()).(_dafny.Int)
			_ = _rhs37
			var _rhs38 _dafny.Int = _68_v1
			_ = _rhs38
			var _lhs31 _dafny.Array = _171_v81
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _171_v81
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))
			_ = _lhs34
			(_lhs31).ArraySet1(_rhs37, (_lhs32).Int())
			(_lhs33).ArraySet1(_rhs38, (_lhs34).Int())
		} else {
			_68_v1 = _68_v1
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
			_ = _index53
			(_163_v75).ArraySet1((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), (_index53).Int())
			var _182_v90 _dafny.Array
			_ = _182_v90
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
			_ = _nw39
			_182_v90 = _nw39
			Companion_Default___.M0(_182_v90, _71_globalState)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
			_ = _index54
			var _rhs39 bool = _160_v72
			_ = _rhs39
			var _rhs40 bool = _160_v72
			_ = _rhs40
			var _lhs35 *GlobalState = _71_globalState
			_ = _lhs35
			var _lhs36 _dafny.Array = _163_v75
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
			_ = _lhs37
			_lhs35.F2 = _rhs39
			(_lhs36).ArraySet1(_rhs40, (_lhs37).Int())
			_68_v1 = (_171_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_171_v81), 0))).Int()).(_dafny.Int)
		}
	}
	var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
	_ = _index55
	(_163_v75).ArraySet1((_160_v72) && (_160_v72), (_index55).Int())
	_68_v1 = _dafny.IntOfUint32(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Cardinality())
	var _183_v91 _dafny.MultiSet
	_ = _183_v91
	_183_v91 = _dafny.MultiSetOf(_68_v1)
	var _184_v92 _dafny.Map
	_ = _184_v92
	_184_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_183_v91).Contains(_68_v1) {
			return (_183_v91).Multiplicity(_68_v1)
		}
		return _dafny.IntOfUint32((_75_v5).Cardinality())
	})(), _68_v1)
	var _hi2 _dafny.Int = _68_v1
	_ = _hi2
	for _185_i12 := (_68_v1).Plus((_184_v92).Cardinality()); _185_i12.Cmp(_hi2) < 0; _185_i12 = _185_i12.Plus(_dafny.One) {
		_160_v72 = (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)
		_184_v92 = (_184_v92).Update(_68_v1, _185_i12)
		var _186_v93 _dafny.Map
		_ = _186_v93
		_186_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, _160_v72)
		var _187_v94 _dafny.MultiSet
		_ = _187_v94
		_187_v94 = _dafny.MultiSetOf(_160_v72)
		if (_160_v72) || ((func() bool {
			if (_186_v93).Contains((_187_v94).Cardinality()) {
				return (_186_v93).Get((_187_v94).Cardinality()).(bool)
			}
			return _160_v72
		})()) {
			var _188_v95 *C0
			_ = _188_v95
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__()
			_188_v95 = _nw40
			var _189_v96 _dafny.Set
			_ = _189_v96
			_189_v96 = _dafny.SetOf((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
			var _190_v97 _dafny.Sequence
			_ = _190_v97
			_190_v97 = _dafny.SeqOf(_185_i12)
			var _191_v98 _dafny.Array
			_ = _191_v98
			var _nwElement0_13 _dafny.Int = (_dafny.MultiSetOf(_68_v1)).Cardinality()
			_ = _nwElement0_13
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(15))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_13, 0)
			(_nw41).ArraySet1(_dafny.IntOfInt64(918), 1)
			(_nw41).ArraySet1((func() _dafny.Int {
				if !((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)) {
					return _185_i12
				}
				return _68_v1
			})(), 2)
			(_nw41).ArraySet1(((_189_v96).Cardinality()).Times(_dafny.IntOfInt64(773)), 3)
			(_nw41).ArraySet1(_dafny.IntOfInt64(-623), 4)
			(_nw41).ArraySet1((_183_v91).Cardinality(), 5)
			(_nw41).ArraySet1(_185_i12, 6)
			(_nw41).ArraySet1(_185_i12, 7)
			(_nw41).ArraySet1((_dafny.IntOfUint32((_190_v97).Cardinality())).Plus(_68_v1), 8)
			(_nw41).ArraySet1((func() _dafny.Int {
				if (_184_v92).Contains(_185_i12) {
					return (_184_v92).Get(_185_i12).(_dafny.Int)
				}
				return _185_i12
			})(), 9)
			(_nw41).ArraySet1((_187_v94).Cardinality(), 10)
			(_nw41).ArraySet1((_68_v1).Times(_185_i12), 11)
			(_nw41).ArraySet1(_185_i12, 12)
			(_nw41).ArraySet1((_185_i12).Plus(_185_i12), 13)
			(_nw41).ArraySet1(_68_v1, 14)
			_191_v98 = _nw41
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_191_v98), 0))
			_ = _index56
			(_191_v98).ArraySet1(_185_i12, (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_191_v98), 0))
			_ = _index57
			(_191_v98).ArraySet1(_185_i12, (_index57).Int())
			var _192_v99 _dafny.Map
			_ = _192_v99
			_192_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v94, _160_v72)
			_192_v99 = _192_v99
			var _193_v100 _dafny.Array
			_ = _193_v100
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_8
			var _nw42 _dafny.Array
			_ = _nw42
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw42 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = func(_194_i13 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg12 _dafny.Int) interface{} {
							return coer12(arg12)
						}
					}(func(_195_i14 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(-612)
					}))
				}
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw42 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw42).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw42).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_193_v100 = _nw42
			var _196_v101 _dafny.Sequence
			_ = _196_v101
			_196_v101 = _dafny.SeqOf(_188_v95)
			var _197_v102 _dafny.Map
			_ = _197_v102
			_197_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_188_v95, _159_v71)
			var _198_v103 _dafny.Array
			_ = _198_v103
			var _nwElement0_14 *C0 = _159_v71
			_ = _nwElement0_14
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(23))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_14, 0)
			(_nw43).ArraySet1((func() *C0 {
				if (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool) {
					return _188_v95
				}
				return _188_v95
			})(), 1)
			(_nw43).ArraySet1(_188_v95, 2)
			(_nw43).ArraySet1(_188_v95, 3)
			(_nw43).ArraySet1(_188_v95, 4)
			(_nw43).ArraySet1((_196_v101).Select((Companion_Default___.SafeIndex(_185_i12, _dafny.IntOfUint32((_196_v101).Cardinality()))).Uint32()).(*C0), 5)
			(_nw43).ArraySet1(_159_v71, 6)
			(_nw43).ArraySet1(_188_v95, 7)
			(_nw43).ArraySet1(_188_v95, 8)
			(_nw43).ArraySet1(_159_v71, 9)
			(_nw43).ArraySet1(_159_v71, 10)
			(_nw43).ArraySet1(_159_v71, 11)
			(_nw43).ArraySet1(_188_v95, 12)
			(_nw43).ArraySet1(_188_v95, 13)
			(_nw43).ArraySet1(_188_v95, 14)
			(_nw43).ArraySet1(_188_v95, 15)
			(_nw43).ArraySet1(_159_v71, 16)
			(_nw43).ArraySet1(_159_v71, 17)
			(_nw43).ArraySet1(_159_v71, 18)
			(_nw43).ArraySet1(_188_v95, 19)
			(_nw43).ArraySet1((func() *C0 {
				if (_197_v102).Contains(_159_v71) {
					return (_197_v102).Get(_159_v71).(*C0)
				}
				return _159_v71
			})(), 20)
			(_nw43).ArraySet1(_159_v71, 21)
			(_nw43).ArraySet1(_159_v71, 22)
			_198_v103 = _nw43
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_198_v103), 0))
			_ = _index58
			(_198_v103).ArraySet1(_159_v71, (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_198_v103), 0))
			_ = _index59
			var _rhs41 _dafny.Int = (_186_v93).Cardinality()
			_ = _rhs41
			var _rhs42 _dafny.Array = _193_v100
			_ = _rhs42
			var _rhs43 _dafny.Array = _163_v75
			_ = _rhs43
			var _rhs44 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(382))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_199_v5 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_200_i15 _dafny.Int) _dafny.Sequence {
					return _199_v5
				}
			})(_75_v5)))
			_ = _rhs44
			var _rhs45 *C0 = _188_v95
			_ = _rhs45
			var _lhs38 _dafny.Array = _198_v103
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_198_v103), 0))
			_ = _lhs39
			_69_v2 = _rhs41
			_193_v100 = _rhs42
			_163_v75 = _rhs43
			_67_v0 = _rhs44
			(_lhs38).ArraySet1(_rhs45, (_lhs39).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_191_v98), 0))
			_ = _index60
			(_191_v98).ArraySet1(Companion_Default___.SafeDivisionInt((_191_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_191_v98), 0))).Int()).(_dafny.Int), (_dafny.MultiSetOf(Companion_Default___.Fm8(_160_v72, _160_v72, _71_globalState))).Cardinality()), (_index60).Int())
		} else {
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))
			_ = _index61
			(_72_v4).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), _75_v5), (_index61).Int())
			_68_v1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), _75_v5), (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Int {
				if _160_v72 {
					return (_dafny.Zero).Minus(_185_i12)
				}
				return _dafny.IntOfInt64(166)
			})()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), _75_v5), (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence))).Cardinality()))).Uint32(), Companion_Default___.Fm8(_160_v72, (_159_v71).Fm1(_71_globalState), _71_globalState))).Cardinality())
			var _201_v104 _dafny.Array
			_ = _201_v104
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_9
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.CodePoint = func(_202_i16 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw44 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw44).ArraySet1CodePoint(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw44).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_201_v104 = _nw44
			var _203_v105 _dafny.CodePoint
			_ = _203_v105
			_203_v105 = _dafny.CodePoint('n')
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(890), _dafny.ArrayLen((_201_v104), 0))
			_ = _index62
			(_201_v104).ArraySet1CodePoint(_203_v105, (_index62).Int())
			var _204_v106 _dafny.Map
			_ = _204_v106
			_204_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), _203_v105)
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
			_ = _index63
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(890), _dafny.ArrayLen((_201_v104), 0))
			_ = _index64
			var _rhs46 _dafny.Int = (_dafny.IntOfInt64(471))
			_ = _rhs46
			var _rhs47 bool = (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)
			_ = _rhs47
			var _rhs48 _dafny.Int = _dafny.IntOfInt64(995)
			_ = _rhs48
			var _rhs49 _dafny.Int = (_204_v106).Cardinality()
			_ = _rhs49
			var _rhs50 _dafny.CodePoint = _203_v105
			_ = _rhs50
			var _lhs40 _dafny.Array = _163_v75
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
			_ = _lhs41
			var _lhs42 _dafny.Array = _201_v104
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(890), _dafny.ArrayLen((_201_v104), 0))
			_ = _lhs43
			_68_v1 = _rhs46
			(_lhs40).ArraySet1(_rhs47, (_lhs41).Int())
			_68_v1 = _rhs48
			_68_v1 = _rhs49
			(_lhs42).ArraySet1CodePoint(_rhs50, (_lhs43).Int())
			(_71_globalState).F2 = !((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
			var _205_v107 *C0
			_ = _205_v107
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__()
			_205_v107 = _nw45
		}
		var _206_v108 _dafny.Array
		_ = _206_v108
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_10
		var _nw46 _dafny.Array
		_ = _nw46
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw46 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.MultiSet = (func(_207_v91 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_208_i17 _dafny.Int) _dafny.MultiSet {
					return _207_v91
				}
			})(_183_v91)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw46 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw46).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw46).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_206_v108 = _nw46
		var _209_v109 _dafny.Sequence
		_ = _209_v109
		_209_v109 = _dafny.SeqOf(_206_v108, _206_v108)
		Companion_Default___.M0((_209_v109).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_209_v109).Cardinality()))).Uint32()).(_dafny.Array), _71_globalState)
	}
	var _210_v110 *C0
	_ = _210_v110
	var _nw47 *C0 = New_C0_()
	_ = _nw47
	_nw47.Ctor__()
	_210_v110 = _nw47
	var _211_v111 *C0
	_ = _211_v111
	var _nw48 *C0 = New_C0_()
	_ = _nw48
	_nw48.Ctor__()
	_211_v111 = _nw48
	var _212_i18 _dafny.Int
	_ = _212_i18
	_212_i18 = _dafny.Zero
	{
		for (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool) {
			{
				if (_212_i18).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_212_i18 = (_212_i18).Plus(_dafny.One)
				_68_v1 = Companion_Default___.SafeDivisionInt(_68_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), _160_v72)).Cardinality())
				var _213_v112 _dafny.Array
				_ = _213_v112
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_11
				var _nw49 _dafny.Array
				_ = _nw49
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw49 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.MultiSet = (func(_214_v91 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_215_i19 _dafny.Int) _dafny.MultiSet {
							return _214_v91
						}
					})(_183_v91)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw49 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw49).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw49).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_213_v112 = _nw49
				Companion_Default___.M0(_213_v112, _71_globalState)
				var _216_v113 _dafny.Sequence
				_ = _216_v113
				_216_v113 = _166_v76
				_216_v113 = _216_v113
				_68_v1 = _68_v1
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_68_v1 = Companion_Default___.SafeDivisionInt((_68_v1).Plus((_dafny.Zero).Minus(_68_v1)), _dafny.IntOfInt64(598))
	var _rhs51 *C0 = _210_v110
	_ = _rhs51
	var _rhs52 _dafny.Int = _68_v1
	_ = _rhs52
	_210_v110 = _rhs51
	_68_v1 = _rhs52
	if _dafny.Companion_Sequence_.Equal(_75_v5, (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)) {
		var _217_v114 _dafny.Array
		_ = _217_v114
		var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw50
		_217_v114 = _nw50
		_217_v114 = _217_v114
		var _218_v115 _dafny.MultiSet
		_ = _218_v115
		_218_v115 = _dafny.MultiSetOf(_160_v72)
		_218_v115 = Companion_Default___.Fm9((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), _71_globalState)
		(_71_globalState).F2 = !((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
		var _219_v116 _dafny.Array
		_ = _219_v116
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_12
		var _nw51 _dafny.Array
		_ = _nw51
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw51 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Map = (func(_220_v1 _dafny.Int, _221_v75 _dafny.Array) func(_dafny.Int) _dafny.Map {
				return func(_222_i20 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v1, (_221_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_221_v75), 0))).Int()).(bool))
				}
			})(_68_v1, _163_v75)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw51 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw51).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw51).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_219_v116 = _nw51
		var _223_v117 _dafny.Map
		_ = _223_v117
		_223_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v116, (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
		_223_v117 = (_223_v117).Update(_219_v116, _160_v72)
		var _source3 _dafny.Sequence = _166_v76
		_ = _source3
		var _224___mcc_h2 _dafny.Sequence = _source3
		_ = _224___mcc_h2
		var _225_cf2 _dafny.Sequence = _224___mcc_h2
		_ = _225_cf2
		_68_v1 = _68_v1
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_217_v114), 0))
		_ = _index65
		(_217_v114).ArraySet1((_68_v1).Minus(_68_v1), (_index65).Int())
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_217_v114), 0))
		_ = _index66
		(_217_v114).ArraySet1((func() _dafny.Int {
			if !(false) {
				return _68_v1
			}
			return _dafny.IntOfUint32((_75_v5).Cardinality())
		})(), (_index66).Int())
		var _226_v118 _dafny.CodePoint
		_ = _226_v118
		_226_v118 = _dafny.CodePoint('f')
		_75_v5 = Companion_Default___.Fm7((_217_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_217_v114), 0))).Int()).(_dafny.Int), _160_v72, (func() bool {
			if _160_v72 {
				return false
			}
			return _160_v72
		})(), _226_v118, _71_globalState)
		var _227_v120 _dafny.Array
		_ = _227_v120
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_13
		var _nw52 _dafny.Array
		_ = _nw52
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw52 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.CodePoint = (func(_228_v4 _dafny.Array, _229_v114 _dafny.Array, _230_v1 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_231_i21 _dafny.Int) _dafny.CodePoint {
					return ((_228_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_228_v4), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SetOf((_229_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_229_v114), 0))).Int()).(_dafny.Int)), func() _dafny.Set {
						var _coll1 = _dafny.NewBuilder()
						_ = _coll1
						for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(881), _dafny.IntOfInt64(625))); ; {
							_compr_1, _ok1 := _iter1()
							if !_ok1 {
								break
							}
							var _232_v119 _dafny.Int
							_232_v119 = interface{}(_compr_1).(_dafny.Int)
							if ((_dafny.IntOfInt64(881)).Cmp(_232_v119) <= 0) && ((_232_v119).Cmp(_dafny.IntOfInt64(625)) < 0) {
								_coll1.Add((_232_v119).Minus(_230_v1))
							}
						}
						return _coll1.ToSet()
					}())).Cardinality()), _dafny.IntOfUint32(((_228_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_228_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_72_v4, _217_v114, _68_v1)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw52 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw52).ArraySet1CodePoint(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw52).ArraySet1CodePoint(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_227_v120 = _nw52
		var _233_v121 _dafny.Map
		_ = _233_v121
		_233_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v120, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v1, (_159_v71).Fm1(_71_globalState)))
		var _234_v122 _dafny.Sequence
		_ = _234_v122
		_234_v122 = _dafny.SeqOf(_233_v121, _233_v121)
		var _235_v123 _dafny.Map
		_ = _235_v123
		_235_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_217_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_217_v114), 0))).Int()).(_dafny.Int), _160_v72)
		var _236_v124 _dafny.Array
		_ = _236_v124
		var _nwElement0_15 _dafny.Map = _233_v121
		_ = _nwElement0_15
		var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(6))
		_ = _nw53
		(_nw53).ArraySet1(_nwElement0_15, 0)
		(_nw53).ArraySet1((_234_v122).Select((Companion_Default___.SafeIndex((_217_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_217_v114), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_234_v122).Cardinality()))).Uint32()).(_dafny.Map), 1)
		(_nw53).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v120, _235_v123)).Update(_227_v120, (_235_v123).Update((_217_v114).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_217_v114), 0))).Int()).(_dafny.Int), _160_v72))).Merge(_233_v121), 2)
		(_nw53).ArraySet1((_233_v121).Merge(_233_v121), 3)
		(_nw53).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v120, _235_v123), 4)
		(_nw53).ArraySet1((_233_v121).Merge(_233_v121), 5)
		_236_v124 = _nw53
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_236_v124), 0))
		_ = _index67
		(_236_v124).ArraySet1(_233_v121, (_index67).Int())
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_236_v124), 0))
		_ = _index68
		(_236_v124).ArraySet1(_233_v121, (_index68).Int())
	} else {
		var _237_v125 _dafny.Map
		_ = _237_v125
		_237_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v91, _211_v111)
		_237_v125 = (_237_v125).Update((func() _dafny.MultiSet {
			if (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool) {
				return _dafny.MultiSetOf(_68_v1)
			}
			return _183_v91
		})(), _211_v111)
		_210_v110 = _159_v71
		if (_68_v1).Cmp(Companion_Default___.SafeModuloInt(_68_v1, _68_v1)) == 0 {
			var _238_v126 *C0
			_ = _238_v126
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__()
			_238_v126 = _nw54
			var _239_v127 _dafny.CodePoint
			_ = _239_v127
			_239_v127 = _dafny.CodePoint('s')
			var _rhs53 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_75_v5, (Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_75_v5).Cardinality()))).Uint32(), _239_v127), (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence))
			_ = _rhs53
			var _rhs54 _dafny.Int = _69_v2
			_ = _rhs54
			var _rhs55 bool = (_68_v1).Cmp(_68_v1) != 0
			_ = _rhs55
			var _lhs44 *GlobalState = _71_globalState
			_ = _lhs44
			_75_v5 = _rhs53
			_69_v2 = _rhs54
			_lhs44.F2 = _rhs55
			(_71_globalState).F2 = !((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
			var _240_v128 _dafny.Array
			_ = _240_v128
			var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(6))
			_ = _nw55
			_240_v128 = _nw55
			var _241_v129 _dafny.Array
			_ = _241_v129
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(14))
			_ = _nw56
			_241_v129 = _nw56
			var _242_v130 _dafny.Set
			_ = _242_v130
			_242_v130 = _dafny.SetOf(_241_v129, _241_v129, _241_v129, _241_v129, _241_v129)
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_240_v128), 0))
			_ = _index69
			(_240_v128).ArraySet1(_242_v130, (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_240_v128), 0))
			_ = _index70
			(_240_v128).ArraySet1((_242_v130).Difference(_dafny.SetOf(_241_v129, _241_v129, _241_v129)), (_index70).Int())
			var _243_v131 *C0
			_ = _243_v131
			var _nw57 *C0 = New_C0_()
			_ = _nw57
			_nw57.Ctor__()
			_243_v131 = _nw57
		} else {
			var _244_v132 D4
			_ = _244_v132
			_244_v132 = Companion_D4_.Create_DC6_(_183_v91)
			var _245_v133 _dafny.Array
			_ = _245_v133
			var _nwElement0_16 _dafny.MultiSet = _183_v91
			_ = _nwElement0_16
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(25))
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_16, 0)
			(_nw58).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_166_v76).Cardinality())), 1)
			(_nw58).ArraySet1(_183_v91, 2)
			(_nw58).ArraySet1(_183_v91, 3)
			(_nw58).ArraySet1(_183_v91, 4)
			(_nw58).ArraySet1(_183_v91, 5)
			(_nw58).ArraySet1((_244_v132).Dtor_cf6(), 6)
			(_nw58).ArraySet1(_183_v91, 7)
			(_nw58).ArraySet1(_183_v91, 8)
			(_nw58).ArraySet1(_183_v91, 9)
			(_nw58).ArraySet1((_183_v91).Update(_dafny.IntOfInt64(840), Companion_Default___.Abs(_68_v1)), 10)
			(_nw58).ArraySet1(_183_v91, 11)
			(_nw58).ArraySet1(_183_v91, 12)
			(_nw58).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_166_v76).Cardinality()), _68_v1), 13)
			(_nw58).ArraySet1(_183_v91, 14)
			(_nw58).ArraySet1(_183_v91, 15)
			(_nw58).ArraySet1(_183_v91, 16)
			(_nw58).ArraySet1(_183_v91, 17)
			(_nw58).ArraySet1(_183_v91, 18)
			(_nw58).ArraySet1(_183_v91, 19)
			(_nw58).ArraySet1(_183_v91, 20)
			(_nw58).ArraySet1(_183_v91, 21)
			(_nw58).ArraySet1(_183_v91, 22)
			(_nw58).ArraySet1(_183_v91, 23)
			(_nw58).ArraySet1(_183_v91, 24)
			_245_v133 = _nw58
			Companion_Default___.M0(_245_v133, _71_globalState)
			_183_v91 = _183_v91
			(_71_globalState).F2 = (_160_v72) || (_160_v72)
			_70_v3 = _70_v3
			_68_v1 = (_68_v1).Minus(_68_v1)
		}
		var _246_v134 _dafny.Map
		_ = _246_v134
		_246_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(_75_v5, (_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)))
		_246_v134 = (_246_v134).Update(_160_v72, _dafny.UnicodeSeqOfUtf8Bytes("mxmseuuv"))
		var _247_v135 _dafny.Map
		_ = _247_v135
		_247_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v110, _68_v1)
		var _248_v136 _dafny.Sequence
		_ = _248_v136
		_248_v136 = _dafny.SeqOf(_211_v111, _210_v110, _210_v110, _211_v111)
		if (((Companion_Default___.Fm6(_160_v72, _71_globalState)).Update(_dafny.IntOfUint32((_248_v136).Cardinality()), Companion_Default___.Abs(_dafny.IntOfUint32((_166_v76).Cardinality())))).IsProperSubsetOf(_dafny.MultiSetOf(_68_v1, _68_v1, (_247_v135).Cardinality()))) == (_160_v72) {
			var _249_v137 _dafny.Array
			_ = _249_v137
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw59
			_249_v137 = _nw59
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))
			_ = _index71
			(_249_v137).ArraySet1(_dafny.IntOfUint32(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()), (_index71).Int())
			var _250_v138 _dafny.Map
			_ = _250_v138
			_250_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_166_v76).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_166_v76).Cardinality()))).Uint32()).(bool)), _68_v1)
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))
			_ = _index72
			(_249_v137).ArraySet1((_dafny.Zero).Minus((_68_v1).Times(((func() _dafny.Int {
				if (_250_v138).Contains(true) {
					return (_250_v138).Get(true).(_dafny.Int)
				}
				return (func() _dafny.Map {
					var _coll2 = _dafny.NewMapBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-503), _dafny.IntOfInt64(961))); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _251_v139 _dafny.Int
						_251_v139 = interface{}(_compr_2).(_dafny.Int)
						if ((_dafny.IntOfInt64(-503)).Cmp(_251_v139) <= 0) && ((_251_v139).Cmp(_dafny.IntOfInt64(961)) < 0) {
							_coll2.Add(Companion_Default___.SafeModuloInt(_251_v139, (func() _dafny.Int {
								if (_184_v92).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_68_v1, _68_v1)).Cardinality())) {
									return (_184_v92).Get(_dafny.IntOfUint32((_dafny.SeqOf(_68_v1, _68_v1)).Cardinality())).(_dafny.Int)
								}
								return _68_v1
							})()), (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
						}
					}
					return _coll2.ToMap()
				}()).Cardinality()
			})()).Minus(_68_v1))), (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))
			_ = _index73
			(_249_v137).ArraySet1(Companion_Default___.SafeDivisionInt(_68_v1, (_249_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))).Int()).(_dafny.Int)), (_index73).Int())
			var _252_v140 _dafny.Array
			_ = _252_v140
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
			_ = _nw60
			_252_v140 = _nw60
			Companion_Default___.M0(_252_v140, _71_globalState)
			var _253_v141 _dafny.Array
			_ = _253_v141
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw61
			_253_v141 = _nw61
			var _254_v142 _dafny.Sequence
			_ = _254_v142
			_254_v142 = _dafny.SeqOf(_68_v1)
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_253_v141), 0))
			_ = _index74
			(_253_v141).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_254_v142, _254_v142), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_250_v138).Contains((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)) {
					return (_250_v138).Get((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)).(_dafny.Int)
				}
				return (_249_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))).Int()).(_dafny.Int)
			})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_254_v142, _254_v142)).Cardinality()))).Uint32(), (_249_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))).Int()).(_dafny.Int)), (_index74).Int())
			var _255_v143 _dafny.Map
			_ = _255_v143
			_255_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_75_v5).Cardinality()), (_249_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))).Int()).(_dafny.Int)), (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_253_v141), 0))
			_ = _index75
			var _rhs56 _dafny.Sequence = _254_v142
			_ = _rhs56
			var _rhs57 _dafny.Map = (_255_v143).Merge((_255_v143).Update((_249_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_249_v137), 0))).Int()).(_dafny.Int), (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)))
			_ = _rhs57
			var _lhs45 _dafny.Array = _253_v141
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_253_v141), 0))
			_ = _lhs46
			(_lhs45).ArraySet1(_rhs56, (_lhs46).Int())
			_255_v143 = _rhs57
			_184_v92 = _184_v92
		} else {
			var _256_v144 _dafny.CodePoint
			_ = _256_v144
			_256_v144 = _dafny.CodePoint('y')
			var _257_v145 _dafny.Array
			_ = _257_v145
			var _nwElement0_17 _dafny.CodePoint = _dafny.CodePoint('e')
			_ = _nwElement0_17
			var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(20))
			_ = _nw62
			(_nw62).ArraySet1CodePoint(_nwElement0_17, 0)
			(_nw62).ArraySet1CodePoint(_256_v144, 1)
			(_nw62).ArraySet1CodePoint(_256_v144, 2)
			(_nw62).ArraySet1CodePoint(_dafny.CodePoint('p'), 3)
			(_nw62).ArraySet1CodePoint(Companion_Default___.Fm8(false, (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), _71_globalState), 4)
			(_nw62).ArraySet1CodePoint(_256_v144, 5)
			(_nw62).ArraySet1CodePoint(_256_v144, 6)
			(_nw62).ArraySet1CodePoint(_256_v144, 7)
			(_nw62).ArraySet1CodePoint(_256_v144, 8)
			(_nw62).ArraySet1CodePoint(_256_v144, 9)
			(_nw62).ArraySet1CodePoint(_256_v144, 10)
			(_nw62).ArraySet1CodePoint(_256_v144, 11)
			(_nw62).ArraySet1CodePoint(_256_v144, 12)
			(_nw62).ArraySet1CodePoint(_256_v144, 13)
			(_nw62).ArraySet1CodePoint(_256_v144, 14)
			(_nw62).ArraySet1CodePoint(_256_v144, 15)
			(_nw62).ArraySet1CodePoint(_dafny.CodePoint('x'), 16)
			(_nw62).ArraySet1CodePoint(_256_v144, 17)
			(_nw62).ArraySet1CodePoint(_256_v144, 18)
			(_nw62).ArraySet1CodePoint(_256_v144, 19)
			_257_v145 = _nw62
			var _258_v146 _dafny.Sequence
			_ = _258_v146
			_258_v146 = _dafny.SeqOf(_68_v1)
			var _259_v147 D4
			_ = _259_v147
			_259_v147 = Companion_D4_.Create_DC7_(_75_v5, _68_v1, _257_v145, _160_v72, (_258_v146).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_258_v146).Cardinality()))).Uint32()).(_dafny.Int))
			var _260_v148 _dafny.Sequence
			_ = _260_v148
			_260_v148 = _dafny.SeqOf(Companion_D4_.Create_DC8_((_259_v147).Dtor_cf10(), (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool)))
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))
			_ = _index76
			(_163_v75).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_260_v148, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer14 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_261_v75 _dafny.Array, _262_v72 bool) func(_dafny.Int) D4 {
				return func(_263_i22 _dafny.Int) D4 {
					return Companion_D4_.Create_DC8_((_261_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_261_v75), 0))).Int()).(bool), (_261_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_261_v75), 0))).Int()).(bool), _262_v72)
				}
			})(_163_v75, _160_v72)))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_260_v148, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer15 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_264_v75 _dafny.Array, _265_v72 bool) func(_dafny.Int) D4 {
				return func(_266_i22 _dafny.Int) D4 {
					return Companion_D4_.Create_DC8_((_264_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_264_v75), 0))).Int()).(bool), (_264_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_264_v75), 0))).Int()).(bool), _265_v72)
				}
			})(_163_v75, _160_v72))))).Cardinality()))).Uint32(), Companion_D4_.Create_DC8_(_160_v72, (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), _160_v72)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(65))).Uint32(), func(coer16 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_267_v75 _dafny.Array) func(_dafny.Int) D4 {
				return func(_268_i23 _dafny.Int) D4 {
					return Companion_D4_.Create_DC8_((_267_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_267_v75), 0))).Int()).(bool), (_267_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_267_v75), 0))).Int()).(bool), (_267_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_267_v75), 0))).Int()).(bool))
				}
			})(_163_v75)))), (_index76).Int())
			var _269_v149 _dafny.Set
			_ = _269_v149
			_269_v149 = _dafny.SetOf(_160_v72, _160_v72)
			var _rhs58 _dafny.CodePoint = ((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs58
			var _rhs59 bool = (_68_v1).Cmp(((_269_v149).Union(_269_v149)).Cardinality()) < 0
			_ = _rhs59
			var _lhs47 *GlobalState = _71_globalState
			_ = _lhs47
			_256_v144 = _rhs58
			_lhs47.F2 = _rhs59
			_161_v73 = (_161_v73).Update((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), (_68_v1).Cmp(_dafny.IntOfUint32(((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).Cardinality())) == 0)
			var _270_v150 _dafny.MultiSet
			_ = _270_v150
			_270_v150 = _dafny.MultiSetOf(_dafny.CodePoint('t'), _256_v144, _256_v144, _256_v144, _256_v144)
			var _271_v151 _dafny.Map
			_ = _271_v151
			_271_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence), (_270_v150).Cardinality())
			var _272_v152 _dafny.Map
			_ = _272_v152
			_272_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v75, (_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool))
			var _273_v153 _dafny.Array
			_ = _273_v153
			var _nwElement0_18 _dafny.Int = _68_v1
			_ = _nwElement0_18
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(21))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_18, 0)
			(_nw63).ArraySet1((_70_v3).Cardinality(), 1)
			(_nw63).ArraySet1(_68_v1, 2)
			(_nw63).ArraySet1(_68_v1, 3)
			(_nw63).ArraySet1(_68_v1, 4)
			(_nw63).ArraySet1(_dafny.IntOfUint32((_75_v5).Cardinality()), 5)
			(_nw63).ArraySet1(_68_v1, 6)
			(_nw63).ArraySet1((_dafny.Zero).Minus(_68_v1), 7)
			(_nw63).ArraySet1((func() _dafny.Int {
				if (_271_v151).Contains((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)) {
					return (_271_v151).Get((_72_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_72_v4), 0))).Int()).(_dafny.Sequence)).(_dafny.Int)
				}
				return _68_v1
			})(), 8)
			(_nw63).ArraySet1(_68_v1, 9)
			(_nw63).ArraySet1(_68_v1, 10)
			(_nw63).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_68_v1)), 11)
			(_nw63).ArraySet1(_68_v1, 12)
			(_nw63).ArraySet1((_269_v149).Cardinality(), 13)
			(_nw63).ArraySet1((_68_v1).Plus(_68_v1), 14)
			(_nw63).ArraySet1(_dafny.IntOfInt64(346), 15)
			(_nw63).ArraySet1(_68_v1, 16)
			(_nw63).ArraySet1((Companion_Default___.Fm10((_163_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_163_v75), 0))).Int()).(bool), !(_160_v72), (func() bool {
				if (_272_v152).Contains(_163_v75) {
					return (_272_v152).Get(_163_v75).(bool)
				}
				return _160_v72
			})(), _71_globalState)).Cardinality(), 17)
			(_nw63).ArraySet1((_258_v146).Select((Companion_Default___.SafeIndex(_68_v1, _dafny.IntOfUint32((_258_v146).Cardinality()))).Uint32()).(_dafny.Int), 18)
			(_nw63).ArraySet1((_dafny.IntOfInt64(491)).Plus(_68_v1), 19)
			(_nw63).ArraySet1(_68_v1, 20)
			_273_v153 = _nw63
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_273_v153), 0))
			_ = _index77
			(_273_v153).ArraySet1(_68_v1, (_index77).Int())
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_273_v153), 0))
			_ = _index78
			(_273_v153).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_68_v1)), (_index78).Int())
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(301), _dafny.ArrayLen((_273_v153), 0))
			_ = _index79
			(_273_v153).ArraySet1(_68_v1, (_index79).Int())
		}
	}
	var _274_v154 *C0
	_ = _274_v154
	var _nw64 *C0 = New_C0_()
	_ = _nw64
	_nw64.Ctor__()
	_274_v154 = _nw64
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_67_v0, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("smr"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_68_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_69_v2))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_70_v3).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(456))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F1()).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(456))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_72_v4).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_72_v4).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_75_v5.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_160_v72)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v73).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v74).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v74).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v74).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v74).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_162_v74).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v75).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_166_v76, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_167_v77, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v91).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v92).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(10))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_212_i18)
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
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D1_DC1 struct {
	Cf1 _dafny.Set
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Set) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D1) Dtor_cf1() _dafny.Set {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC1:
		{
			return "D1.DC1" + "(" + _dafny.String(data.Cf1) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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

type D2_DC2 struct {
	Cf2 _dafny.Sequence
}

func (D2_DC2) isD2() {}

func (CompanionStruct_D2_) Create_DC2_(Cf2 _dafny.Sequence) D2 {
	return D2{D2_DC2{Cf2}}
}

func (_this D2) Is_DC2() bool {
	_, ok := _this.Get_().(D2_DC2)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D2) Dtor_cf2() _dafny.Sequence {
	return _this.Get_().(D2_DC2).Cf2
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC2:
		{
			return "D2.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC2:
		{
			data2, ok := other.Get_().(D2_DC2)
			return ok && data1.Cf2.Equals(data2.Cf2)
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

type D3_DC4 struct {
	Cf4 bool
}

func (D3_DC4) isD3() {}

func (CompanionStruct_D3_) Create_DC4_(Cf4 bool) D3 {
	return D3{D3_DC4{Cf4}}
}

func (_this D3) Is_DC4() bool {
	_, ok := _this.Get_().(D3_DC4)
	return ok
}

type D3_DC3 struct {
	Cf3 _dafny.Array
}

func (D3_DC3) isD3() {}

func (CompanionStruct_D3_) Create_DC3_(Cf3 _dafny.Array) D3 {
	return D3{D3_DC3{Cf3}}
}

func (_this D3) Is_DC3() bool {
	_, ok := _this.Get_().(D3_DC3)
	return ok
}

type D3_DC5 struct {
	Cf5 D3
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf5 D3) D3 {
	return D3{D3_DC5{Cf5}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC4_(false)
}

func (_this D3) Dtor_cf4() bool {
	return _this.Get_().(D3_DC4).Cf4
}

func (_this D3) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D3_DC3).Cf3
}

func (_this D3) Dtor_cf5() D3 {
	return _this.Get_().(D3_DC5).Cf5
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC4:
		{
			return "D3.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D3_DC3:
		{
			return "D3.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf5) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC4:
		{
			data2, ok := other.Get_().(D3_DC4)
			return ok && data1.Cf4 == data2.Cf4
		}
	case D3_DC3:
		{
			data2, ok := other.Get_().(D3_DC3)
			return ok && data1.Cf3 == data2.Cf3
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf5.Equals(data2.Cf5)
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

type D4_DC7 struct {
	Cf7  _dafny.Sequence
	Cf8  _dafny.Int
	Cf9  _dafny.Array
	Cf10 bool
	Cf11 _dafny.Int
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf7 _dafny.Sequence, Cf8 _dafny.Int, Cf9 _dafny.Array, Cf10 bool, Cf11 _dafny.Int) D4 {
	return D4{D4_DC7{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

type D4_DC8 struct {
	Cf12 bool
	Cf13 bool
	Cf14 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf12 bool, Cf13 bool, Cf14 bool) D4 {
	return D4{D4_DC8{Cf12, Cf13, Cf14}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC6 struct {
	Cf6 _dafny.MultiSet
}

func (D4_DC6) isD4() {}

func (CompanionStruct_D4_) Create_DC6_(Cf6 _dafny.MultiSet) D4 {
	return D4{D4_DC6{Cf6}}
}

func (_this D4) Is_DC6() bool {
	_, ok := _this.Get_().(D4_DC6)
	return ok
}

type D4_DC9 struct {
	Cf15 D4
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf15 D4) D4 {
	return D4{D4_DC9{Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC7_(_dafny.EmptySeq, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero)
}

func (_this D4) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D4_DC7).Cf7
}

func (_this D4) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D4_DC7).Cf8
}

func (_this D4) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D4_DC7).Cf9
}

func (_this D4) Dtor_cf10() bool {
	return _this.Get_().(D4_DC7).Cf10
}

func (_this D4) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D4_DC7).Cf11
}

func (_this D4) Dtor_cf12() bool {
	return _this.Get_().(D4_DC8).Cf12
}

func (_this D4) Dtor_cf13() bool {
	return _this.Get_().(D4_DC8).Cf13
}

func (_this D4) Dtor_cf14() bool {
	return _this.Get_().(D4_DC8).Cf14
}

func (_this D4) Dtor_cf6() _dafny.MultiSet {
	return _this.Get_().(D4_DC6).Cf6
}

func (_this D4) Dtor_cf15() D4 {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC7:
		{
			return "D4.DC7" + "(" + data.Cf7.VerbatimString(true) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D4_DC6:
		{
			return "D4.DC6" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf7.Equals(data2.Cf7) && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14
		}
	case D4_DC6:
		{
			data2, ok := other.Get_().(D4_DC6)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15.Equals(data2.Cf15)
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
	Cf17 bool
	Cf18 _dafny.MultiSet
	Cf19 bool
	Cf20 _dafny.Int
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf17 bool, Cf18 _dafny.MultiSet, Cf19 bool, Cf20 _dafny.Int) D5 {
	return D5{D5_DC11{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC12 struct {
	Cf21 D4
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf21 D4) D5 {
	return D5{D5_DC12{Cf21}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC10 struct {
	Cf16 _dafny.Map
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf16 _dafny.Map) D5 {
	return D5{D5_DC10{Cf16}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC11_(false, _dafny.EmptyMultiSet, false, _dafny.Zero)
}

func (_this D5) Dtor_cf17() bool {
	return _this.Get_().(D5_DC11).Cf17
}

func (_this D5) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D5_DC11).Cf18
}

func (_this D5) Dtor_cf19() bool {
	return _this.Get_().(D5_DC11).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC11).Cf20
}

func (_this D5) Dtor_cf21() D4 {
	return _this.Get_().(D5_DC12).Cf21
}

func (_this D5) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D5_DC10).Cf16
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18) && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf21.Equals(data2.Cf21)
		}
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
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
	Cf23 _dafny.Int
	Cf24 _dafny.Sequence
	Cf25 _dafny.Int
	Cf26 _dafny.Sequence
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf23 _dafny.Int, Cf24 _dafny.Sequence, Cf25 _dafny.Int, Cf26 _dafny.Sequence) D6 {
	return D6{D6_DC14{Cf23, Cf24, Cf25, Cf26}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC13 struct {
	Cf22 _dafny.Map
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf22 _dafny.Map) D6 {
	return D6{D6_DC13{Cf22}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC15 struct {
	Cf27 D6
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf27 D6) D6 {
	return D6{D6_DC15{Cf27}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(_dafny.Zero, _dafny.EmptySeq, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D6_DC14).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D6_DC14).Cf26
}

func (_this D6) Dtor_cf22() _dafny.Map {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) Dtor_cf27() D6 {
	return _this.Get_().(D6_DC15).Cf27
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf23) + ", " + data.Cf24.VerbatimString(true) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf27) + ")"
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Equals(data2.Cf24) && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26.Equals(data2.Cf26)
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf22.Equals(data2.Cf22)
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	F2  bool
	_f0 _dafny.CodePoint
	_f1 _dafny.Set
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this._f0 = _dafny.CodePoint('D')
	_this._f1 = _dafny.EmptySet
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

func (_this *GlobalState) Ctor__(f0 _dafny.CodePoint, f1 _dafny.Set, f2 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
	}
}
func (_this *GlobalState) F0() _dafny.CodePoint {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Set {
	{
		return _this._f1
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
func (_this *C0) Fm1(globalState *GlobalState) bool {
	{
		return true
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
