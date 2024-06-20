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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-469), false)).Cardinality()).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-411), (_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(94))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(366)
	})))).Cardinality())).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(850), _dafny.IntOfInt64(668))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(850)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(668)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_1_v0, _dafny.IntOfInt64(-868)), (func() _dafny.Set {
					var _coll1 = _dafny.NewBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), false)).Keys().Elements()); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _2_v1 _dafny.CodePoint
						_2_v1 = interface{}(_compr_1).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), false)).Contains(_2_v1) {
							_coll1.Add(_2_v1)
						}
					}
					return _coll1.ToSet()
				}()).Cardinality())
			}
		}
		return _coll0.ToMap()
	}())).Cardinality()))).Union(_dafny.SetOf(_dafny.IntOfInt64(176)))).Union((_dafny.SetOf(_dafny.IntOfInt64(166))).Union(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("henpihgvp")).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(91))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()), _dafny.IntOfInt64(936))).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer2 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_4_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC4_(Companion_D1_.Create_DC3_(true, true, false, !(!(!(true)))))
	}))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((Companion_D4_.Create_DC11_(_dafny.CodePoint('b'), true, _dafny.MultiSetOf(_dafny.CodePoint('t')), true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-835))).Uint32(), func(coer3 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_5_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.IntOfInt64(493))
	})))).Dtor_cf20())
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('g')
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) bool {
	var _source0 D3 = Companion_D3_.Create_DC7_()
	_ = _source0
	if _source0.Is_DC7() {
		return false
	} else if _source0.Is_DC8() {
		var _6___mcc_h0 _dafny.Int = _source0.Get_().(D3_DC8).Cf14
		_ = _6___mcc_h0
		var _7_cf14 _dafny.Int = _6___mcc_h0
		_ = _7_cf14
		return false
	} else {
		var _8___mcc_h1 _dafny.Sequence = _source0.Get_().(D3_DC6).Cf13
		_ = _8___mcc_h1
		var _9_cf13 _dafny.Sequence = _8___mcc_h1
		_ = _9_cf13
		return !(!(false))
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('x'))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if false {
			return _dafny.SeqOf((_dafny.SetOf(false, false)).Cardinality())
		}
		return _dafny.SeqOf((func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(939), _dafny.IntOfInt64(466))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _10_v0 _dafny.Int
				_10_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(939)).Cmp(_10_v0) <= 0) && ((_10_v0).Cmp(_dafny.IntOfInt64(466)) < 0) {
					_coll2.Add((_10_v0).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality()), false)
				}
			}
			return _coll2.ToMap()
		}()).Cardinality(), _dafny.IntOfInt64(844), _dafny.IntOfInt64(518))
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_11_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(244)
	})))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC7_()
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	var _source1 D4 = Companion_D4_.Create_DC11_(_dafny.CodePoint('h'), false, _dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('d'), _dafny.CodePoint('d')), true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer5 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_12_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf((func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(421), _dafny.IntOfInt64(551))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _13_v0 _dafny.Int
				_13_v0 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(421)).Cmp(_13_v0) <= 0) && ((_13_v0).Cmp(_dafny.IntOfInt64(551)) < 0) {
					_coll3.Add((_13_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dxbtxxe")).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gi")).Cardinality()))
				}
			}
			return _coll3.ToMap()
		}()).Cardinality())
	})))
	_ = _source1
	if _source1.Is_DC10() {
		var _14___mcc_h0 _dafny.Int = _source1.Get_().(D4_DC10).Cf16
		_ = _14___mcc_h0
		var _15___mcc_h1 bool = _source1.Get_().(D4_DC10).Cf17
		_ = _15___mcc_h1
		var _16___mcc_h2 bool = _source1.Get_().(D4_DC10).Cf18
		_ = _16___mcc_h2
		var _17_cf18 bool = _16___mcc_h2
		_ = _17_cf18
		var _18_cf17 bool = _15___mcc_h1
		_ = _18_cf17
		var _19_cf16 _dafny.Int = _14___mcc_h0
		_ = _19_cf16
		return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_18_cf17, true))
	} else if _source1.Is_DC11() {
		var _20___mcc_h3 _dafny.CodePoint = _source1.Get_().(D4_DC11).Cf19
		_ = _20___mcc_h3
		var _21___mcc_h4 bool = _source1.Get_().(D4_DC11).Cf20
		_ = _21___mcc_h4
		var _22___mcc_h5 _dafny.MultiSet = _source1.Get_().(D4_DC11).Cf21
		_ = _22___mcc_h5
		var _23___mcc_h6 bool = _source1.Get_().(D4_DC11).Cf22
		_ = _23___mcc_h6
		var _24___mcc_h7 _dafny.Sequence = _source1.Get_().(D4_DC11).Cf23
		_ = _24___mcc_h7
		var _25_cf23 _dafny.Sequence = _24___mcc_h7
		_ = _25_cf23
		var _26_cf22 bool = _23___mcc_h6
		_ = _26_cf22
		var _27_cf21 _dafny.MultiSet = _22___mcc_h5
		_ = _27_cf21
		var _28_cf20 bool = _21___mcc_h4
		_ = _28_cf20
		var _29_cf19 _dafny.CodePoint = _20___mcc_h3
		_ = _29_cf19
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_cf22, _26_cf22), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_cf20, _28_cf20))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _26_cf22), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_cf20, _26_cf22))))
	} else if _source1.Is_DC12() {
		var _30___mcc_h8 _dafny.Int = _source1.Get_().(D4_DC12).Cf24
		_ = _30___mcc_h8
		var _31___mcc_h9 _dafny.Array = _source1.Get_().(D4_DC12).Cf25
		_ = _31___mcc_h9
		var _32___mcc_h10 bool = _source1.Get_().(D4_DC12).Cf26
		_ = _32___mcc_h10
		var _33___mcc_h11 _dafny.Int = _source1.Get_().(D4_DC12).Cf27
		_ = _33___mcc_h11
		var _34_cf27 _dafny.Int = _33___mcc_h11
		_ = _34_cf27
		var _35_cf26 bool = _32___mcc_h10
		_ = _35_cf26
		var _36_cf25 _dafny.Array = _31___mcc_h9
		_ = _36_cf25
		var _37_cf24 _dafny.Int = _30___mcc_h8
		_ = _37_cf24
		return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_cf26, _35_cf26), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_cf26, (Companion_D4_.Create_DC10_(_37_cf24, false, _35_cf26)).Dtor_cf18()))
	} else {
		var _38___mcc_h12 *C0 = _source1.Get_().(D4_DC9).Cf15
		_ = _38___mcc_h12
		var _39_cf15 *C0 = _38___mcc_h12
		_ = _39_cf15
		return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_40_i1 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
		}))))
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) (_dafny.Set, bool, _dafny.MultiSet) {
	var r0 _dafny.Set = _dafny.EmptySet
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
	_ = r2
	if p1 {
		(globalState).F15 = Companion_Default___.SafeModuloInt(p0, Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(639)))
		var _41_v0 _dafny.Sequence
		_ = _41_v0
		_41_v0 = _dafny.SeqOf(p0, p0, (_dafny.Zero).Minus(Companion_Default___.Fm0(p1, p3, p2, p0, globalState)), (p0).Minus(p0))
		var _42_v1 _dafny.CodePoint
		_ = _42_v1
		_42_v1 = _dafny.CodePoint('y')
		var _43_v2 *C0
		_ = _43_v2
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__(p0, _42_v1)
		_43_v2 = _nw0
		var _44_v3 _dafny.Sequence
		_ = _44_v3
		_44_v3 = _dafny.UnicodeSeqOfUtf8Bytes("deyue")
		var _45_v4 D0
		_ = _45_v4
		_45_v4 = Companion_D0_.Create_DC0_(_44_v3)
		var _rhs0 _dafny.Sequence = Companion_Default___.Fm8(!(p2) || (p2), globalState)
		_ = _rhs0
		var _rhs1 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7(false, (_43_v2).F19(), _dafny.IntOfUint32((_41_v0).Cardinality()), globalState), _44_v3), (func() _dafny.Sequence {
			if !(p2) {
				return _44_v3
			}
			return (_45_v4).Dtor_cf0()
		})())
		_ = _rhs1
		var _rhs2 _dafny.Int = (_43_v2).F19()
		_ = _rhs2
		var _rhs3 *C0 = _43_v2
		_ = _rhs3
		var _rhs4 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hkykdidt"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(74))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_46_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_47_i0 _dafny.Int) _dafny.CodePoint {
				return _46_v1
			}
		})(_42_v1))))
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		_41_v0 = _rhs0
		_lhs0.F6 = _rhs1
		_lhs1.F13 = _rhs2
		_43_v2 = _rhs3
		_44_v3 = _rhs4
		var _48_v5 _dafny.Array
		_ = _48_v5
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_49_v2 *C0) func(_dafny.Int) _dafny.Int {
				return func(_50_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_50_i1, (_49_v2).F19())
				}
			})(_43_v2)
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
		_48_v5 = _nw1
		var _51_v6 _dafny.Map
		_ = _51_v6
		_51_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _48_v5)
		var _52_v7 _dafny.Map
		_ = _52_v7
		_52_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _53_v8 _dafny.Map
		_ = _53_v8
		_53_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if (_51_v6).Contains(true) {
				return (_51_v6).Get(true).(_dafny.Array)
			}
			return _48_v5
		})(), _52_v7)
		_53_v8 = (_53_v8).Update(_48_v5, _52_v7)
		var _54_v9 _dafny.Map
		_ = _54_v9
		_54_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v0, _48_v5)
		_54_v9 = (_54_v9).Merge(_54_v9)
		var _55_v10 _dafny.Array
		_ = _55_v10
		var _nwElement0_0 _dafny.CodePoint = _42_v1
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(14))
		_ = _nw2
		(_nw2).ArraySet1CodePoint(_nwElement0_0, 0)
		(_nw2).ArraySet1CodePoint(_42_v1, 1)
		(_nw2).ArraySet1CodePoint((p3).Select((Companion_Default___.SafeIndex((_43_v2).F19(), _dafny.IntOfUint32((p3).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
		(_nw2).ArraySet1CodePoint((_43_v2).F20(), 3)
		(_nw2).ArraySet1CodePoint((_43_v2).F20(), 4)
		(_nw2).ArraySet1CodePoint((_43_v2).F20(), 5)
		(_nw2).ArraySet1CodePoint(_42_v1, 6)
		(_nw2).ArraySet1CodePoint(_dafny.CodePoint('c'), 7)
		(_nw2).ArraySet1CodePoint(_42_v1, 8)
		(_nw2).ArraySet1CodePoint(_42_v1, 9)
		(_nw2).ArraySet1CodePoint((_43_v2).F20(), 10)
		(_nw2).ArraySet1CodePoint((_43_v2).F20(), 11)
		(_nw2).ArraySet1CodePoint(_42_v1, 12)
		(_nw2).ArraySet1CodePoint((_43_v2).F20(), 13)
		_55_v10 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_55_v10), 0))
		_ = _index0
		(_55_v10).ArraySet1CodePoint(_42_v1, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_55_v10), 0))
		_ = _index1
		(_55_v10).ArraySet1CodePoint((_43_v2).F20(), (_index1).Int())
	} else {
		var _56_v11 _dafny.CodePoint
		_ = _56_v11
		_56_v11 = _dafny.CodePoint('a')
		(globalState).F6 = _dafny.Companion_Sequence_.Contains(p3, (func() _dafny.CodePoint {
			if p2 {
				return _56_v11
			}
			return Companion_Default___.Fm5(globalState)
		})())
		var _57_v12 D3
		_ = _57_v12
		_57_v12 = Companion_D3_.Create_DC7_()
		_57_v12 = _57_v12
		var _58_v13 _dafny.MultiSet
		_ = _58_v13
		_58_v13 = _dafny.MultiSetOf((func() bool {
			if p1 {
				return !(p1)
			}
			return p1
		})(), (false) && (p1))
		_58_v13 = Companion_Default___.Fm4(p2, globalState)
		(globalState).F13 = (_dafny.Zero).Minus(p0)
		var _59_v14 _dafny.Array
		_ = _59_v14
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_1
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = (func(_60_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_61_i2 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_61_i2, _60_p0)
				}
			})(p0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw3 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw3).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw3).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_59_v14 = _nw3
		var _62_v15 _dafny.Array
		_ = _62_v15
		var _nwElement0_1 bool = p2
		_ = _nwElement0_1
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(19))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_1, 0)
		(_nw4).ArraySet1(!(p1), 1)
		(_nw4).ArraySet1(p2, 2)
		(_nw4).ArraySet1(p2, 3)
		(_nw4).ArraySet1(p1, 4)
		(_nw4).ArraySet1(p1, 5)
		(_nw4).ArraySet1(p1, 6)
		(_nw4).ArraySet1(p2, 7)
		(_nw4).ArraySet1(p2, 8)
		(_nw4).ArraySet1(true, 9)
		(_nw4).ArraySet1(p2, 10)
		(_nw4).ArraySet1(!(p1), 11)
		(_nw4).ArraySet1(p2, 12)
		(_nw4).ArraySet1(false, 13)
		(_nw4).ArraySet1(p2, 14)
		(_nw4).ArraySet1(p1, 15)
		(_nw4).ArraySet1(p2, 16)
		(_nw4).ArraySet1(p2, 17)
		(_nw4).ArraySet1(p1, 18)
		_62_v15 = _nw4
		var _63_v16 _dafny.Sequence
		_ = _63_v16
		_63_v16 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(691))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_64_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_65_i3 _dafny.Int) _dafny.CodePoint {
				return _64_v11
			}
		})(_56_v11))), p3)
		var _66_v17 D0
		_ = _66_v17
		_66_v17 = Companion_D0_.Create_DC1_(_59_v14, _62_v15, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_63_v16, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_63_v16).Cardinality()))).Uint32(), _dafny.SeqOf(_56_v11))), p0, p0)
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_59_v14), 0))
		_ = _index2
		(_59_v14).ArraySet1((_66_v17).Dtor_cf4(), (_index2).Int())
		var _67_v18 D1
		_ = _67_v18
		_67_v18 = Companion_D1_.Create_DC3_(p2, p1, true, Companion_Default___.Fm6(p0, Companion_Default___.Fm5(globalState), (_dafny.Zero).Minus(p0), globalState))
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_59_v14), 0))
		_ = _index3
		(_59_v14).ArraySet1((_dafny.Zero).Minus((_dafny.MultiSetOf(p1, (_67_v18).Dtor_cf8())).Cardinality()), (_index3).Int())
	}
	var _68_v19 _dafny.Array
	_ = _68_v19
	var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
	_ = _nw5
	_68_v19 = _nw5
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_68_v19), 0))); ; {
		_guard_loop_0, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _69_i4 _dafny.Int
		_69_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_69_i4).Sign() != -1) && ((_69_i4).Cmp(_dafny.ArrayLen((_68_v19), 0)) < 0)) {
			(_68_v19).ArraySet1(Companion_Default___.SafeModuloInt(_69_i4, p0), (_69_i4).Int())
		}
	}
	var _hi0 _dafny.Int = p0
	_ = _hi0
	for _70_i5 := p0; _70_i5.Cmp(_hi0) < 0; _70_i5 = _70_i5.Plus(_dafny.One) {
		var _71_v20 _dafny.Array
		_ = _71_v20
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw6
		_71_v20 = _nw6
		var _72_v21 _dafny.Set
		_ = _72_v21
		_72_v21 = _dafny.SetOf(_71_v20)
		_72_v21 = _72_v21
		var _73_v22 _dafny.Array
		_ = _73_v22
		var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
		_ = _nw7
		_73_v22 = _nw7
		_73_v22 = _73_v22
		var _74_v23 _dafny.Map
		_ = _74_v23
		_74_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _source2 D3 = Companion_Default___.Fm9(true, ((_74_v23).Merge(_74_v23)).Cardinality(), globalState)
		_ = _source2
		if _source2.Is_DC7() {
			(globalState).F6 = true
			var _75_v24 _dafny.Sequence
			_ = _75_v24
			_75_v24 = _dafny.SeqOf(p1)
			var _76_v25 _dafny.Sequence
			_ = _76_v25
			_76_v25 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), _dafny.CodePoint('i')), (Companion_Default___.SafeIndex(_70_i5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), _dafny.CodePoint('i'))).Cardinality()))).Uint32(), Companion_Default___.Fm5(globalState)), p3, _dafny.Companion_Sequence_.Concatenate(p3, p3), Companion_Default___.Fm7(!((_75_v24).Select((Companion_Default___.SafeIndex(_70_i5, _dafny.IntOfUint32((_75_v24).Cardinality()))).Uint32()).(bool)), p0, p0, globalState))
			_76_v25 = _dafny.Companion_Sequence_.Concatenate(_76_v25, _76_v25)
			var _77_v26 _dafny.Map
			_ = _77_v26
			_77_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
			var _78_v27 _dafny.Map
			_ = _78_v27
			_78_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2) || (p1), _77_v26)
			_78_v27 = (_78_v27).Update(p1, (_77_v26).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)))
			var _79_v28 _dafny.CodePoint
			_ = _79_v28
			_79_v28 = _dafny.CodePoint('w')
			var _80_v29 *C0
			_ = _80_v29
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__(p0, _79_v28)
			_80_v29 = _nw8
			var _81_v30 D4
			_ = _81_v30
			_81_v30 = Companion_D4_.Create_DC9_(_80_v29)
			var _82_v31 _dafny.Array
			_ = _82_v31
			var _nwElement0_2 *C0 = _80_v29
			_ = _nwElement0_2
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(19))
			_ = _nw9
			(_nw9).ArraySet1(_nwElement0_2, 0)
			(_nw9).ArraySet1(_80_v29, 1)
			(_nw9).ArraySet1(_80_v29, 2)
			(_nw9).ArraySet1(_80_v29, 3)
			(_nw9).ArraySet1((_81_v30).Dtor_cf15(), 4)
			(_nw9).ArraySet1(_80_v29, 5)
			(_nw9).ArraySet1(_80_v29, 6)
			(_nw9).ArraySet1(_80_v29, 7)
			(_nw9).ArraySet1(_80_v29, 8)
			(_nw9).ArraySet1(_80_v29, 9)
			(_nw9).ArraySet1(_80_v29, 10)
			(_nw9).ArraySet1(_80_v29, 11)
			(_nw9).ArraySet1(_80_v29, 12)
			(_nw9).ArraySet1(_80_v29, 13)
			(_nw9).ArraySet1(_80_v29, 14)
			(_nw9).ArraySet1(_80_v29, 15)
			(_nw9).ArraySet1(_80_v29, 16)
			(_nw9).ArraySet1(_80_v29, 17)
			(_nw9).ArraySet1(_80_v29, 18)
			_82_v31 = _nw9
			var _83_v32 _dafny.Set
			_ = _83_v32
			_83_v32 = _dafny.SetOf(_82_v31, _82_v31, _82_v31, _82_v31, _82_v31)
			_83_v32 = _dafny.SetOf((func() _dafny.Array {
				if p2 {
					return _82_v31
				}
				return _82_v31
			})())
		} else if _source2.Is_DC8() {
			var _84___mcc_h0 _dafny.Int = _source2.Get_().(D3_DC8).Cf14
			_ = _84___mcc_h0
			var _85_cf14 _dafny.Int = _84___mcc_h0
			_ = _85_cf14
			var _86_v33 _dafny.CodePoint
			_ = _86_v33
			_86_v33 = _dafny.CodePoint('c')
			var _87_v34 _dafny.Map
			_ = _87_v34
			_87_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v33, p1)
			var _88_v35 _dafny.Map
			_ = _88_v35
			_88_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v20, _87_v34)
			var _89_v36 _dafny.Sequence
			_ = _89_v36
			_89_v36 = _dafny.SeqOf((func() _dafny.Map {
				if p1 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v20, _87_v34)
				}
				return _88_v35
			})())
			var _90_v37 *C0
			_ = _90_v37
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__(_85_cf14, _86_v33)
			_90_v37 = _nw10
			var _91_v38 _dafny.Map
			_ = _91_v38
			_91_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v37, _70_i5)
			var _rhs5 _dafny.Map = (_89_v36).Select((Companion_Default___.SafeIndex(_85_cf14, _dafny.IntOfUint32((_89_v36).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs5
			var _rhs6 bool = (((_91_v38).Update(_90_v37, _85_cf14)).Cardinality()).Cmp(_dafny.IntOfInt64(908)) < 0
			_ = _rhs6
			_88_v35 = _rhs5
			r1 = _rhs6
			var _92_v39 _dafny.Map
			_ = _92_v39
			_92_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(470), _dafny.IntOfUint32(((Companion_D0_.Create_DC0_(p3)).Dtor_cf0()).Cardinality()))
			r1 = (func() bool {
				if (p1) == (true) {
					return ((_90_v37).F19()).Cmp((_92_v39).Cardinality()) > 0
				}
				return p2
			})()
			_71_v20 = _71_v20
			(globalState).F15 = _85_cf14
		} else {
			var _93___mcc_h1 _dafny.Sequence = _source2.Get_().(D3_DC6).Cf13
			_ = _93___mcc_h1
			var _94_cf13 _dafny.Sequence = _93___mcc_h1
			_ = _94_cf13
			var _95_v40 _dafny.CodePoint
			_ = _95_v40
			_95_v40 = _dafny.CodePoint('t')
			var _96_v41 *C0
			_ = _96_v41
			var _nw11 *C0 = New_C0_()
			_ = _nw11
			_nw11.Ctor__(p0, _95_v40)
			_96_v41 = _nw11
			var _97_v42 _dafny.Sequence
			_ = _97_v42
			_97_v42 = _dafny.SeqOf(p0, _dafny.IntOfUint32((p3).Cardinality()), _70_i5)
			var _98_v43 _dafny.Map
			_ = _98_v43
			_98_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v42, _96_v41)
			var _99_v44 _dafny.MultiSet
			_ = _99_v44
			_99_v44 = _dafny.MultiSetOf((_96_v41).F20(), (_96_v41).F20())
			var _100_v45 _dafny.MultiSet
			_ = _100_v45
			_100_v45 = _dafny.MultiSetOf((_dafny.SetOf(!(p2))).Cardinality(), _dafny.IntOfInt64(328))
			var _101_v46 _dafny.Sequence
			_ = _101_v46
			_101_v46 = _dafny.SeqOf(_100_v45, _100_v45)
			var _102_v47 D4
			_ = _102_v47
			_102_v47 = Companion_D4_.Create_DC11_((_96_v41).F20(), p1, _99_v44, p1, _101_v46)
			var _rhs7 bool = (_102_v47).Dtor_cf22()
			_ = _rhs7
			var _rhs8 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_v42, _96_v41)
			_ = _rhs8
			r1 = _rhs7
			_98_v43 = _rhs8
			(globalState).F5 = (_dafny.Zero).Minus((((_96_v41).F19()).Times((_96_v41).F19())).Plus(_70_i5))
			_97_v42 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_97_v42, _97_v42), (func() _dafny.Sequence {
				if p2 {
					return _97_v42
				}
				return _dafny.SeqOf((_96_v41).F19())
			})())
		}
		var _103_v48 _dafny.Map
		_ = _103_v48
		_103_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _104_v49 _dafny.CodePoint
		_ = _104_v49
		_104_v49 = _dafny.CodePoint('c')
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_68_v19), 0))
		_ = _index4
		(_68_v19).ArraySet1(Companion_Default___.SafeDivisionInt(_70_i5, Companion_Default___.Fm0(p1, p3, Companion_Default___.Fm6((_103_v48).Cardinality(), _104_v49, _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()), globalState), p0, globalState)), (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_68_v19), 0))
		_ = _index5
		(_68_v19).ArraySet1((_dafny.Zero).Minus(p0), (_index5).Int())
	}
	var _105_v50 _dafny.Array
	_ = _105_v50
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(25)
	_ = _len0_2
	var _nw12 _dafny.Array
	_ = _nw12
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw12 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Sequence = (func(_106_p3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_107_i6 _dafny.Int) _dafny.Sequence {
				return _106_p3
			}
		})(p3)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw12 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw12).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw12).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_105_v50 = _nw12
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_105_v50), 0))
	_ = _index6
	(_105_v50).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("d"), (_index6).Int())
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_105_v50), 0))
	_ = _index7
	(_105_v50).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("arixobu"), (_index7).Int())
	var _108_v51 _dafny.Array
	_ = _108_v51
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
	_ = _nw13
	_108_v51 = _nw13
	for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_108_v51), 0))); ; {
		_guard_loop_1, _ok5 := _iter5()
		if !_ok5 {
			break
		}
		var _109_i7 _dafny.Int
		_109_i7 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_109_i7).Sign() != -1) && ((_109_i7).Cmp(_dafny.ArrayLen((_108_v51), 0)) < 0)) {
			(_108_v51).ArraySet1(p1, (_109_i7).Int())
		}
	}
	var _110_v52 _dafny.CodePoint
	_ = _110_v52
	_110_v52 = _dafny.CodePoint('x')
	var _111_v53 *C0
	_ = _111_v53
	var _nw14 *C0 = New_C0_()
	_ = _nw14
	_nw14.Ctor__(p0, _110_v52)
	_111_v53 = _nw14
	_111_v53 = _111_v53
	var _112_v54 _dafny.Set
	_ = _112_v54
	_112_v54 = _dafny.SetOf((_111_v53).F19(), p0)
	var _113_v56 _dafny.Set
	_ = _113_v56
	_113_v56 = _dafny.SetOf(_112_v54, (_112_v54).Union(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-145), _dafny.IntOfInt64(464))); ; {
			_compr_4, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _114_v55 _dafny.Int
			_114_v55 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-145)).Cmp(_114_v55) <= 0) && ((_114_v55).Cmp(_dafny.IntOfInt64(464)) < 0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_114_v55, _dafny.IntOfUint32((_dafny.SeqOf(p2, p1, true, p1, p1)).Cardinality())))
			}
		}
		return _coll4.ToSet()
	}()), Companion_Default___.Fm1(p1, (_111_v53).F19(), p0, _dafny.IntOfInt64(503), globalState))
	r0 = _113_v56
	r1 = p1
	var _115_v57 D4
	_ = _115_v57
	_115_v57 = Companion_D4_.Create_DC10_(p0, p2, p1)
	var _116_v58 _dafny.Map
	_ = _116_v58
	_116_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(828))
	var _117_v60 _dafny.MultiSet
	_ = _117_v60
	_117_v60 = _dafny.MultiSetOf(_116_v58, func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(335), _dafny.IntOfInt64(915))); ; {
			_compr_5, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _118_v59 _dafny.Int
			_118_v59 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(335)).Cmp(_118_v59) <= 0) && ((_118_v59).Cmp(_dafny.IntOfInt64(915)) < 0) {
				_coll5.Add((_118_v59).Minus((_112_v54).Cardinality()), (_111_v53).F19())
			}
		}
		return _coll5.ToMap()
	}())
	r2 = Companion_Default___.Fm10(p0, (_dafny.IntOfInt64(435)).Minus((_115_v57).Dtor_cf16()), _117_v60, p0, globalState)
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _119_v0 _dafny.Int
	_ = _119_v0
	_119_v0 = _dafny.IntOfInt64(327)
	var _120_v1 _dafny.Sequence
	_ = _120_v1
	_120_v1 = _dafny.UnicodeSeqOfUtf8Bytes("qljvqw")
	var _121_v2 bool
	_ = _121_v2
	_121_v2 = true
	var _122_v3 _dafny.Sequence
	_ = _122_v3
	_122_v3 = _dafny.SeqOf(_121_v2)
	var _123_v4 _dafny.MultiSet
	_ = _123_v4
	_123_v4 = _dafny.MultiSetOf(_120_v1, _120_v1, _120_v1)
	var _124_v5 _dafny.Array
	_ = _124_v5
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_3
	var _nw15 _dafny.Array
	_ = _nw15
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw15 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) bool = (func(_125_v2 bool) func(_dafny.Int) bool {
			return func(_126_i0 _dafny.Int) bool {
				return _125_v2
			}
		})(_121_v2)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw15 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw15).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw15).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_124_v5 = _nw15
	var _127_globalState *GlobalState
	_ = _127_globalState
	var _nw16 *GlobalState = New_GlobalState_()
	_ = _nw16
	_nw16.Ctor__(_dafny.IntOfInt64(624), _dafny.IntOfInt64(743), false, _dafny.IntOfInt64(-990), false, _dafny.IntOfInt64(913), false, _dafny.SeqOf(_119_v0, _dafny.IntOfUint32((_120_v1).Cardinality()), _119_v0, (_dafny.Zero).Minus(_119_v0)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v3, _121_v2), _dafny.IntOfInt64(580), _dafny.IntOfInt64(647), _dafny.CodePoint('r'), _123_v4, _dafny.IntOfInt64(481), _dafny.SeqOf(_124_v5, _124_v5), _dafny.IntOfInt64(31), true, _dafny.IntOfInt64(-830), false)
	_127_globalState = _nw16
	var _128_v6 _dafny.CodePoint
	_ = _128_v6
	_128_v6 = _dafny.CodePoint('m')
	var _129_v7 _dafny.Sequence
	_ = _129_v7
	_129_v7 = _dafny.SeqOf(_dafny.IntOfInt64(-349))
	var _130_v8 _dafny.Set
	_ = _130_v8
	_130_v8 = _dafny.SetOf(_119_v0, (_129_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-857), _dafny.IntOfUint32((_129_v7).Cardinality()))).Uint32()).(_dafny.Int), _119_v0, _119_v0, _119_v0)
	var _131_v9 _dafny.MultiSet
	_ = _131_v9
	_131_v9 = _dafny.MultiSetOf(Companion_Default___.Fm1(_121_v2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, _121_v2, true, _121_v2, _121_v2)).Cardinality())), _119_v0, (_dafny.Zero).Minus(_119_v0), _127_globalState), _130_v8)
	var _132_v11 _dafny.Map
	_ = _132_v11
	_132_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v0, (_dafny.Zero).Minus(_119_v0))
	var _133_v12 _dafny.Array
	_ = _133_v12
	var _nwElement0_3 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v0, _121_v2)).Cardinality()
	_ = _nwElement0_3
	var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(14))
	_ = _nw17
	(_nw17).ArraySet1(_nwElement0_3, 0)
	(_nw17).ArraySet1(Companion_Default___.SafeModuloInt(_119_v0, _119_v0), 1)
	(_nw17).ArraySet1(_119_v0, 2)
	(_nw17).ArraySet1(_dafny.IntOfInt64(-498), 3)
	(_nw17).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_120_v1, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_121_v2, _120_v1, _121_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v2, _128_v6)).Cardinality(), _127_globalState), _dafny.IntOfUint32((_120_v1).Cardinality()))).Uint32(), _dafny.CodePoint('q'))).Cardinality()), 4)
	(_nw17).ArraySet1(_119_v0, 5)
	(_nw17).ArraySet1(_119_v0, 6)
	(_nw17).ArraySet1((func() _dafny.Int {
		if (_131_v9).Contains(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(218), _dafny.IntOfInt64(451))); ; {
				_compr_6, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _134_v10 _dafny.Int
				_134_v10 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(218)).Cmp(_134_v10) <= 0) && ((_134_v10).Cmp(_dafny.IntOfInt64(451)) < 0) {
					_coll6.Add(Companion_Default___.SafeDivisionInt(_134_v10, (_dafny.Zero).Minus((_130_v8).Cardinality())))
				}
			}
			return _coll6.ToSet()
		}()) {
			return (_131_v9).Multiplicity(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(218), _dafny.IntOfInt64(451))); ; {
					_compr_7, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _135_v10 _dafny.Int
					_135_v10 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(218)).Cmp(_135_v10) <= 0) && ((_135_v10).Cmp(_dafny.IntOfInt64(451)) < 0) {
						_coll7.Add(Companion_Default___.SafeDivisionInt(_135_v10, (_dafny.Zero).Minus((_130_v8).Cardinality())))
					}
				}
				return _coll7.ToSet()
			}())
		}
		return _dafny.IntOfInt64(157)
	})(), 7)
	(_nw17).ArraySet1(_119_v0, 8)
	(_nw17).ArraySet1((_119_v0).Minus(_119_v0), 9)
	(_nw17).ArraySet1(_119_v0, 10)
	(_nw17).ArraySet1(Companion_Default___.SafeDivisionInt(_119_v0, _119_v0), 11)
	(_nw17).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_119_v0), _119_v0), 12)
	(_nw17).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
		if _121_v2 {
			return (_132_v11).Cardinality()
		}
		return (_dafny.Zero).Minus(_119_v0)
	})()), 13)
	_133_v12 = _nw17
	for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_133_v12), 0))); ; {
		_guard_loop_2, _ok10 := _iter10()
		if !_ok10 {
			break
		}
		var _136_i1 _dafny.Int
		_136_i1 = interface{}(_guard_loop_2).(_dafny.Int)
		if (true) && (((_136_i1).Sign() != -1) && ((_136_i1).Cmp(_dafny.ArrayLen((_133_v12), 0)) < 0)) {
			(_133_v12).ArraySet1(Companion_Default___.SafeDivisionInt(_136_i1, _119_v0), (_136_i1).Int())
		}
	}
	var _137_v13 _dafny.MultiSet
	_ = _137_v13
	_137_v13 = _dafny.MultiSetOf(_119_v0, _119_v0)
	var _rhs9 _dafny.CodePoint = _128_v6
	_ = _rhs9
	var _rhs10 _dafny.Int = _119_v0
	_ = _rhs10
	var _rhs11 bool = (_137_v13).IsSubsetOf(_137_v13)
	_ = _rhs11
	var _rhs12 _dafny.Array = _133_v12
	_ = _rhs12
	var _lhs2 *GlobalState = _127_globalState
	_ = _lhs2
	var _lhs3 *GlobalState = _127_globalState
	_ = _lhs3
	_128_v6 = _rhs9
	_lhs2.F5 = _rhs10
	_lhs3.F6 = _rhs11
	_133_v12 = _rhs12
	var _138_v14 _dafny.Set
	_ = _138_v14
	var _139_v15 bool
	_ = _139_v15
	var _140_v16 _dafny.MultiSet
	_ = _140_v16
	var _out0 _dafny.Set
	_ = _out0
	var _out1 bool
	_ = _out1
	var _out2 _dafny.MultiSet
	_ = _out2
	_out0, _out1, _out2 = Companion_Default___.M0(_119_v0, false, _121_v2, _120_v1, _127_globalState)
	_138_v14 = _out0
	_139_v15 = _out1
	_140_v16 = _out2
	(_127_globalState).F5 = _119_v0
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
	_ = _index8
	(_124_v5).ArraySet1(!(_139_v15), (_index8).Int())
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
	_ = _index9
	(_124_v5).ArraySet1(_139_v15, (_index9).Int())
	var _141_v17 _dafny.Array
	_ = _141_v17
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(10)
	_ = _len0_4
	var _nw18 _dafny.Array
	_ = _nw18
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw18 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Sequence = (func(_142_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_143_i2 _dafny.Int) _dafny.Sequence {
				return _142_v1
			}
		})(_120_v1)
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
	_141_v17 = _nw18
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))
	_ = _index10
	(_141_v17).ArraySet1(_120_v1, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))
	_ = _index11
	(_141_v17).ArraySet1((Companion_D0_.Create_DC0_(_120_v1)).Dtor_cf0(), (_index11).Int())
	var _144_i3 _dafny.Int
	_ = _144_i3
	_144_i3 = _dafny.Zero
	{
		for !(_139_v15) {
			{
				if (_144_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_144_i3 = (_144_i3).Plus(_dafny.One)
				var _145_v18 D1
				_ = _145_v18
				_145_v18 = Companion_D1_.Create_DC2_(_128_v6)
				var _146_v19 *C0
				_ = _146_v19
				var _nw19 *C0 = New_C0_()
				_ = _nw19
				_nw19.Ctor__(_119_v0, (_145_v18).Dtor_cf6())
				_146_v19 = _nw19
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))
				_ = _index12
				(_133_v12).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ofwu")).Cardinality())).Plus(_dafny.IntOfUint32((_120_v1).Cardinality())), (_index12).Int())
				var _147_v20 _dafny.Map
				_ = _147_v20
				_147_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oijvun")).Cardinality()))
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))
				_ = _index13
				(_133_v12).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm3((func() _dafny.Int {
					if false {
						return (func() _dafny.Int {
							if (_147_v20).Contains(_121_v2) {
								return (_147_v20).Get(_121_v2).(_dafny.Int)
							}
							return _dafny.IntOfInt64(-524)
						})()
					}
					return (_146_v19).F19()
				})(), ((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)) || ((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)), _122_v3, _147_v20, _127_globalState)).Cardinality())), (_index13).Int())
				var _148_v21 _dafny.Map
				_ = _148_v21
				_148_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), !(_139_v15))
				var _source3 D1 = Companion_D1_.Create_DC3_(_139_v15, !_dafny.Companion_Sequence_.Contains(_122_v3, true), _139_v15, _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_129_v7, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_139_v15, (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence), _139_v15, _119_v0, _127_globalState), _dafny.IntOfUint32((_129_v7).Cardinality()))).Uint32(), (_146_v19).F19()), _dafny.SeqOf((_148_v21).Cardinality())))
				_ = _source3
				if _source3.Is_DC3() {
					var _149___mcc_h0 bool = _source3.Get_().(D1_DC3).Cf7
					_ = _149___mcc_h0
					var _150___mcc_h1 bool = _source3.Get_().(D1_DC3).Cf8
					_ = _150___mcc_h1
					var _151___mcc_h2 bool = _source3.Get_().(D1_DC3).Cf9
					_ = _151___mcc_h2
					var _152___mcc_h3 bool = _source3.Get_().(D1_DC3).Cf10
					_ = _152___mcc_h3
					var _153_cf10 bool = _152___mcc_h3
					_ = _153_cf10
					var _154_cf9 bool = _151___mcc_h2
					_ = _154_cf9
					var _155_cf8 bool = _150___mcc_h1
					_ = _155_cf8
					var _156_cf7 bool = _149___mcc_h0
					_ = _156_cf7
					(_127_globalState).F5 = (_146_v19).F19()
					var _157_v22 _dafny.Set
					_ = _157_v22
					var _158_v23 bool
					_ = _158_v23
					var _159_v24 _dafny.MultiSet
					_ = _159_v24
					var _out3 _dafny.Set
					_ = _out3
					var _out4 bool
					_ = _out4
					var _out5 _dafny.MultiSet
					_ = _out5
					_out3, _out4, _out5 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), (_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int)), (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate((_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence), _120_v1), _127_globalState)
					_157_v22 = _out3
					_158_v23 = _out4
					_159_v24 = _out5
					var _160_v25 *C0
					_ = _160_v25
					var _nw20 *C0 = New_C0_()
					_ = _nw20
					_nw20.Ctor__(_119_v0, _128_v6)
					_160_v25 = _nw20
					var _161_v26 _dafny.MultiSet
					_ = _161_v26
					_161_v26 = _dafny.MultiSetOf(_146_v19)
					var _162_v27 _dafny.Array
					_ = _162_v27
					var _nwElement0_4 _dafny.Int = (_160_v25).F19()
					_ = _nwElement0_4
					var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(27))
					_ = _nw21
					(_nw21).ArraySet1(_nwElement0_4, 0)
					(_nw21).ArraySet1((_dafny.SetOf(_158_v23)).Cardinality(), 1)
					(_nw21).ArraySet1((_146_v19).F19(), 2)
					(_nw21).ArraySet1((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), 3)
					(_nw21).ArraySet1((_146_v19).F19(), 4)
					(_nw21).ArraySet1((_146_v19).F19(), 5)
					(_nw21).ArraySet1((_146_v19).F19(), 6)
					(_nw21).ArraySet1((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), 7)
					(_nw21).ArraySet1(_dafny.IntOfInt64(179), 8)
					(_nw21).ArraySet1(_119_v0, 9)
					(_nw21).ArraySet1(_119_v0, 10)
					(_nw21).ArraySet1((_146_v19).F19(), 11)
					(_nw21).ArraySet1((_146_v19).F19(), 12)
					(_nw21).ArraySet1((_148_v21).Cardinality(), 13)
					(_nw21).ArraySet1(_119_v0, 14)
					(_nw21).ArraySet1((_160_v25).F19(), 15)
					(_nw21).ArraySet1(_dafny.IntOfUint32(((_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence)).Cardinality()), 16)
					(_nw21).ArraySet1((_146_v19).F19(), 17)
					(_nw21).ArraySet1((_146_v19).F19(), 18)
					(_nw21).ArraySet1(_119_v0, 19)
					(_nw21).ArraySet1(_dafny.IntOfInt64(-124), 20)
					(_nw21).ArraySet1((_146_v19).F19(), 21)
					(_nw21).ArraySet1((_146_v19).F19(), 22)
					(_nw21).ArraySet1((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), 23)
					(_nw21).ArraySet1((func() _dafny.Int {
						if (_137_v13).Contains((_161_v26).Cardinality()) {
							return (_137_v13).Multiplicity((_161_v26).Cardinality())
						}
						return (_146_v19).F19()
					})(), 24)
					(_nw21).ArraySet1((_146_v19).F19(), 25)
					(_nw21).ArraySet1(_dafny.IntOfInt64(-897), 26)
					_162_v27 = _nw21
					var _163_v28 _dafny.Map
					_ = _163_v28
					_163_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_146_v19).F20(), _162_v27)
					var _164_v29 _dafny.Set
					_ = _164_v29
					var _165_v30 bool
					_ = _165_v30
					var _166_v31 _dafny.MultiSet
					_ = _166_v31
					var _out6 _dafny.Set
					_ = _out6
					var _out7 bool
					_ = _out7
					var _out8 _dafny.MultiSet
					_ = _out8
					_out6, _out7, _out8 = Companion_Default___.M0(_119_v0, !(_163_v28).Equals((_163_v28)), _139_v15, (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence), _127_globalState)
					_164_v29 = _out6
					_165_v30 = _out7
					_166_v31 = _out8
				} else if _source3.Is_DC2() {
					var _167___mcc_h4 _dafny.CodePoint = _source3.Get_().(D1_DC2).Cf6
					_ = _167___mcc_h4
					var _168_cf6 _dafny.CodePoint = _167___mcc_h4
					_ = _168_cf6
					var _169_v32 _dafny.Sequence
					_ = _169_v32
					_169_v32 = _dafny.SeqOf(_129_v7)
					var _170_v33 _dafny.Map
					_ = _170_v33
					_170_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_169_v32).Select((Companion_Default___.SafeIndex((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_169_v32).Cardinality()))).Uint32()).(_dafny.Sequence), (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence))
					var _171_v34 _dafny.Set
					_ = _171_v34
					_171_v34 = _dafny.SetOf(!((_170_v33).Equals(_170_v33)))
					var _172_v35 _dafny.Map
					_ = _172_v35
					_172_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v19, _119_v0)
					var _173_v36 _dafny.MultiSet
					_ = _173_v36
					_173_v36 = _dafny.MultiSetOf((_172_v35).Contains(_146_v19), false)
					var _174_v37 _dafny.Sequence
					_ = _174_v37
					_174_v37 = _dafny.SeqOf(_146_v19)
					var _175_v38 _dafny.Sequence
					_ = _175_v38
					_175_v38 = _dafny.SeqOf(_171_v34)
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))
					_ = _index14
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))
					_ = _index15
					var _rhs13 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_174_v37, (Companion_Default___.SafeIndex((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_174_v37).Cardinality()))).Uint32(), _146_v19)).Cardinality()), _119_v0)
					_ = _rhs13
					var _rhs14 _dafny.Set = (_175_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_122_v3).Cardinality()), _dafny.IntOfUint32((_175_v38).Cardinality()))).Uint32()).(_dafny.Set)
					_ = _rhs14
					var _rhs15 _dafny.MultiSet = Companion_Default___.Fm4(_139_v15, _127_globalState)
					_ = _rhs15
					var _rhs16 _dafny.Int = (_146_v19).F19()
					_ = _rhs16
					var _lhs4 _dafny.Array = _133_v12
					_ = _lhs4
					var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))
					_ = _lhs5
					var _lhs6 _dafny.Array = _133_v12
					_ = _lhs6
					var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))
					_ = _lhs7
					(_lhs4).ArraySet1(_rhs13, (_lhs5).Int())
					_171_v34 = _rhs14
					_173_v36 = _rhs15
					(_lhs6).ArraySet1(_rhs16, (_lhs7).Int())
					_119_v0 = _119_v0
					_130_v8 = func() _dafny.Set {
						var _coll8 = _dafny.NewBuilder()
						_ = _coll8
						for _iter11 := _dafny.Iterate((_129_v7).Elements()); ; {
							_compr_8, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _176_v39 _dafny.Int
							_176_v39 = interface{}(_compr_8).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_129_v7, _176_v39) {
								_coll8.Add(Companion_Default___.SafeDivisionInt(_176_v39, _dafny.IntOfInt64(-557)))
							}
						}
						return _coll8.ToSet()
					}()
					var _177_v40 _dafny.Array
					_ = _177_v40
					var _nw22 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
					_ = _nw22
					_177_v40 = _nw22
					_177_v40 = _177_v40
				} else {
					var _178___mcc_h5 D1 = _source3.Get_().(D1_DC4).Cf11
					_ = _178___mcc_h5
					var _179_cf11 D1 = _178___mcc_h5
					_ = _179_cf11
					(_127_globalState).F15 = (_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int)
					var _180_v41 _dafny.Array
					_ = _180_v41
					var _len0_5 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_5
					var _nw23 _dafny.Array
					_ = _nw23
					if _len0_5.Cmp(_dafny.Zero) == 0 {
						_nw23 = _dafny.NewArray(_len0_5)
					} else {
						var _init5 func(_dafny.Int) _dafny.Int = (func(_181_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_182_i4 _dafny.Int) _dafny.Int {
								return (_182_i4).Times(_181_v0)
							}
						})(_119_v0)
						_ = _init5
						var _element0_5 = _init5(_dafny.Zero)
						_ = _element0_5
						_nw23 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
						(_nw23).ArraySet1(_element0_5, 0)
						var _nativeLen0_5 = (_len0_5).Int()
						_ = _nativeLen0_5
						for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
							(_nw23).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
						}
					}
					_180_v41 = _nw23
					var _183_v42 _dafny.Map
					_ = _183_v42
					_183_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_146_v19).F20(), _180_v41)
					var _184_v43 _dafny.Map
					_ = _184_v43
					_184_v43 = _183_v42
					_184_v43 = _183_v42
					_128_v6 = (_146_v19).F20()
					_184_v43 = _184_v43
				}
				var _185_v44 _dafny.Set
				_ = _185_v44
				var _186_v45 bool
				_ = _186_v45
				var _187_v46 _dafny.MultiSet
				_ = _187_v46
				var _out9 _dafny.Set
				_ = _out9
				var _out10 bool
				_ = _out10
				var _out11 _dafny.MultiSet
				_ = _out11
				_out9, _out10, _out11 = Companion_Default___.M0((_146_v19).F19(), _121_v2, _139_v15, _dafny.Companion_Sequence_.Concatenate((_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence), _120_v1), _127_globalState)
				_185_v44 = _out9
				_186_v45 = _out10
				_187_v46 = _out11
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_133_v12), 0))); ; {
		_guard_loop_3, _ok12 := _iter12()
		if !_ok12 {
			break
		}
		var _188_i5 _dafny.Int
		_188_i5 = interface{}(_guard_loop_3).(_dafny.Int)
		if (true) && (((_188_i5).Sign() != -1) && ((_188_i5).Cmp(_dafny.ArrayLen((_133_v12), 0)) < 0)) {
			(_133_v12).ArraySet1((_188_i5).Minus((func() _dafny.Int {
				if _139_v15 {
					return _119_v0
				}
				return _119_v0
			})()), (_188_i5).Int())
		}
	}
	var _189_v47 D1
	_ = _189_v47
	_189_v47 = Companion_D1_.Create_DC3_((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _121_v2, _121_v2, _139_v15)
	(_127_globalState).F6 = !((_189_v47).Dtor_cf7())
	var _hi1 _dafny.Int = (_dafny.SetOf(_119_v0, _119_v0)).Cardinality()
	_ = _hi1
	for _190_i6 := _119_v0; _190_i6.Cmp(_hi1) < 0; _190_i6 = _190_i6.Plus(_dafny.One) {
		var _191_v48 D0
		_ = _191_v48
		_191_v48 = Companion_D0_.Create_DC1_(_133_v12, _124_v5, _123_v4, _190_i6, _190_i6)
		var _source4 D0 = _191_v48
		_ = _source4
		if _source4.Is_DC1() {
			var _192___mcc_h6 _dafny.Array = _source4.Get_().(D0_DC1).Cf1
			_ = _192___mcc_h6
			var _193___mcc_h7 _dafny.Array = _source4.Get_().(D0_DC1).Cf2
			_ = _193___mcc_h7
			var _194___mcc_h8 _dafny.MultiSet = _source4.Get_().(D0_DC1).Cf3
			_ = _194___mcc_h8
			var _195___mcc_h9 _dafny.Int = _source4.Get_().(D0_DC1).Cf4
			_ = _195___mcc_h9
			var _196___mcc_h10 _dafny.Int = _source4.Get_().(D0_DC1).Cf5
			_ = _196___mcc_h10
			var _197_cf5 _dafny.Int = _196___mcc_h10
			_ = _197_cf5
			var _198_cf4 _dafny.Int = _195___mcc_h9
			_ = _198_cf4
			var _199_cf3 _dafny.MultiSet = _194___mcc_h8
			_ = _199_cf3
			var _200_cf2 _dafny.Array = _193___mcc_h7
			_ = _200_cf2
			var _201_cf1 _dafny.Array = _192___mcc_h6
			_ = _201_cf1
			_128_v6 = (_120_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_202_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_203_i7 _dafny.Int) _dafny.CodePoint {
					return _202_v6
				}
			})(_128_v6)))).Cardinality()), _dafny.IntOfUint32((_120_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			var _204_v49 _dafny.Set
			_ = _204_v49
			var _205_v50 bool
			_ = _205_v50
			var _206_v51 _dafny.MultiSet
			_ = _206_v51
			var _out12 _dafny.Set
			_ = _out12
			var _out13 bool
			_ = _out13
			var _out14 _dafny.MultiSet
			_ = _out14
			_out12, _out13, _out14 = Companion_Default___.M0(_198_cf4, true, _139_v15, _120_v1, _127_globalState)
			_204_v49 = _out12
			_205_v50 = _out13
			_206_v51 = _out14
			(_127_globalState).F5 = _197_cf5
			_197_cf5 = _198_cf4
		} else {
			var _207___mcc_h11 _dafny.Sequence = _source4.Get_().(D0_DC0).Cf0
			_ = _207___mcc_h11
			var _208_cf0 _dafny.Sequence = _207___mcc_h11
			_ = _208_cf0
			_122_v3 = _122_v3
			var _209_v52 _dafny.Array
			_ = _209_v52
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_6
			var _nw24 _dafny.Array
			_ = _nw24
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw24 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) D0 = (func(_210_v1 _dafny.Sequence) func(_dafny.Int) D0 {
					return func(_211_i8 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_210_v1)
					}
				})(_120_v1)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw24 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw24).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw24).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_209_v52 = _nw24
			var _212_v53 D0
			_ = _212_v53
			_212_v53 = Companion_D0_.Create_DC0_(_120_v1)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_209_v52), 0))
			_ = _index16
			(_209_v52).ArraySet1(_212_v53, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_209_v52), 0))
			_ = _index17
			var _rhs17 _dafny.Sequence = (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence)
			_ = _rhs17
			var _rhs18 _dafny.Int = _119_v0
			_ = _rhs18
			var _rhs19 bool = (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)
			_ = _rhs19
			var _rhs20 bool = !(false)
			_ = _rhs20
			var _rhs21 D0 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_208_cf0, _dafny.UnicodeSeqOfUtf8Bytes("fy")))
			_ = _rhs21
			var _lhs8 *GlobalState = _127_globalState
			_ = _lhs8
			var _lhs9 *GlobalState = _127_globalState
			_ = _lhs9
			var _lhs10 *GlobalState = _127_globalState
			_ = _lhs10
			var _lhs11 _dafny.Array = _209_v52
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_209_v52), 0))
			_ = _lhs12
			_208_cf0 = _rhs17
			_lhs8.F13 = _rhs18
			_lhs9.F6 = _rhs19
			_lhs10.F6 = _rhs20
			(_lhs11).ArraySet1(_rhs21, (_lhs12).Int())
			var _213_v54 _dafny.Map
			_ = _213_v54
			_213_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v6, _133_v12)
			var _214_v55 _dafny.Map
			_ = _214_v55
			_214_v55 = _213_v54
			var _215_v56 _dafny.Map
			_ = _215_v56
			_215_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v0, _214_v55)
			_121_v2 = (_190_i6).Cmp(((_215_v56).Update(_190_i6, _214_v55)).Cardinality()) >= 0
			var _216_v57 _dafny.Set
			_ = _216_v57
			var _217_v58 bool
			_ = _217_v58
			var _218_v59 _dafny.MultiSet
			_ = _218_v59
			var _out15 _dafny.Set
			_ = _out15
			var _out16 bool
			_ = _out16
			var _out17 _dafny.MultiSet
			_ = _out17
			_out15, _out16, _out17 = Companion_Default___.M0(_dafny.IntOfInt64(849), ((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)) == ((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)), (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _120_v1, _127_globalState)
			_216_v57 = _out15
			_217_v58 = _out16
			_218_v59 = _out17
		}
		var _219_v60 _dafny.Array
		_ = _219_v60
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_7
		var _nw25 _dafny.Array
		_ = _nw25
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw25 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Sequence = (func(_220_v7 _dafny.Sequence, _221_i6 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_222_i9 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_220_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.IntOfUint32((_220_v7).Cardinality()))).Uint32(), _221_i6)
				}
			})(_129_v7, _190_i6)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw25 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw25).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw25).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_219_v60 = _nw25
		_219_v60 = _219_v60
		(_127_globalState).F6 = !(!(!(_139_v15)))
		var _223_v61 *C0
		_ = _223_v61
		var _nw26 *C0 = New_C0_()
		_ = _nw26
		_nw26.Ctor__(_dafny.IntOfInt64(204), _128_v6)
		_223_v61 = _nw26
		var _224_v62 _dafny.Array
		_ = _224_v62
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_8
		var _nw27 _dafny.Array
		_ = _nw27
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw27 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) D1 = (func(_225_v5 _dafny.Array) func(_dafny.Int) D1 {
				return func(_226_i10 _dafny.Int) D1 {
					return Companion_D1_.Create_DC3_((_225_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_225_v5), 0))).Int()).(bool), (_225_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_225_v5), 0))).Int()).(bool), true, (_225_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_225_v5), 0))).Int()).(bool))
				}
			})(_124_v5)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw27 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw27).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw27).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_224_v62 = _nw27
		var _227_v63 _dafny.Map
		_ = _227_v63
		_227_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v61, _224_v62)
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
		_ = _index18
		var _rhs22 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_119_v0, (_227_v63).Cardinality()), _190_i6)
		_ = _rhs22
		var _rhs23 bool = _121_v2
		_ = _rhs23
		var _rhs24 _dafny.Int = _119_v0
		_ = _rhs24
		var _rhs25 _dafny.Array = _141_v17
		_ = _rhs25
		var _rhs26 _dafny.Int = _dafny.IntOfInt64(827)
		_ = _rhs26
		var _lhs13 _dafny.Array = _124_v5
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
		_ = _lhs14
		var _lhs15 *GlobalState = _127_globalState
		_ = _lhs15
		_119_v0 = _rhs22
		(_lhs13).ArraySet1(_rhs23, (_lhs14).Int())
		_119_v0 = _rhs24
		_141_v17 = _rhs25
		_lhs15.F15 = _rhs26
	}
	var _hi2 _dafny.Int = _119_v0
	_ = _hi2
	for _228_i11 := Companion_Default___.SafeDivisionInt(_119_v0, _119_v0); _228_i11.Cmp(_hi2) < 0; _228_i11 = _228_i11.Plus(_dafny.One) {
		if _121_v2 {
			var _229_v64 *C0
			_ = _229_v64
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__(_119_v0, _128_v6)
			_229_v64 = _nw28
			var _230_v65 _dafny.Map
			_ = _230_v65
			_230_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v2, _229_v64)
			var _231_v66 _dafny.Array
			_ = _231_v66
			var _nwElement0_5 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_139_v15, _229_v64)).Merge(_230_v65)
			_ = _nwElement0_5
			var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(3))
			_ = _nw29
			(_nw29).ArraySet1(_nwElement0_5, 0)
			(_nw29).ArraySet1(_230_v65, 1)
			(_nw29).ArraySet1((_230_v65).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_139_v15), _229_v64)), 2)
			_231_v66 = _nw29
			_231_v66 = _231_v66
			var _232_v67 _dafny.Array
			_ = _232_v67
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_9
			var _nw30 _dafny.Array
			_ = _nw30
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw30 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.MultiSet = (func(_233_v3 _dafny.Sequence, _234_v15 bool) func(_dafny.Int) _dafny.MultiSet {
					return func(_235_i12 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetFromSeq(_233_v3)).Difference(_dafny.MultiSetOf(_234_v15))
					}
				})(_122_v3, _139_v15)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw30 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw30).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw30).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_232_v67 = _nw30
			var _236_v68 _dafny.MultiSet
			_ = _236_v68
			_236_v68 = _dafny.MultiSetOf(_139_v15)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_232_v67), 0))
			_ = _index19
			(_232_v67).ArraySet1((_236_v68).Difference(_dafny.MultiSetOf((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _139_v15, _121_v2)), (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_232_v67), 0))
			_ = _index20
			(_232_v67).ArraySet1(_236_v68, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_133_v12), 0))
			_ = _index21
			(_133_v12).ArraySet1((_229_v64).F19(), (_index21).Int())
			var _237_v69 _dafny.Map
			_ = _237_v69
			_237_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_236_v68).Cardinality(), _139_v15)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_133_v12), 0))
			_ = _index22
			var _rhs27 _dafny.Int = Companion_Default___.SafeModuloInt((_237_v69).Cardinality(), _228_i11)
			_ = _rhs27
			var _rhs28 bool = _121_v2
			_ = _rhs28
			var _lhs16 _dafny.Array = _133_v12
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_133_v12), 0))
			_ = _lhs17
			var _lhs18 *GlobalState = _127_globalState
			_ = _lhs18
			(_lhs16).ArraySet1(_rhs27, (_lhs17).Int())
			_lhs18.F6 = _rhs28
			var _238_v70 _dafny.Sequence
			_ = _238_v70
			_238_v70 = _dafny.SeqOf(_124_v5)
			_238_v70 = _dafny.Companion_Sequence_.Update(_238_v70, (Companion_Default___.SafeIndex((_133_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(807), _dafny.ArrayLen((_133_v12), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_238_v70).Cardinality()))).Uint32(), _124_v5)
			(_127_globalState).F5 = Companion_Default___.SafeDivisionInt(_228_i11, _228_i11)
		} else {
			(_127_globalState).F6 = (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)
			var _239_v71 D0
			_ = _239_v71
			_239_v71 = Companion_D0_.Create_DC1_(_133_v12, _124_v5, _123_v4, _119_v0, (_129_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_120_v1).Cardinality()), _dafny.IntOfUint32((_129_v7).Cardinality()))).Uint32()).(_dafny.Int))
			var _240_v72 _dafny.Sequence
			_ = _240_v72
			_240_v72 = _dafny.SeqOf(_239_v71)
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))
			_ = _index23
			var _rhs29 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("vw"), (Companion_Default___.SafeIndex(_228_i11, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vw")).Cardinality()))).Uint32(), _128_v6), _dafny.UnicodeSeqOfUtf8Bytes("pmpe"))
			_ = _rhs29
			var _rhs30 bool = !(_139_v15) || (_139_v15)
			_ = _rhs30
			var _rhs31 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_240_v72, _240_v72), _240_v72)
			_ = _rhs31
			var _rhs32 _dafny.CodePoint = Companion_Default___.Fm5(_127_globalState)
			_ = _rhs32
			var _lhs19 _dafny.Array = _141_v17
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))
			_ = _lhs20
			var _lhs21 *GlobalState = _127_globalState
			_ = _lhs21
			(_lhs19).ArraySet1(_rhs29, (_lhs20).Int())
			_lhs21.F6 = _rhs30
			_240_v72 = _rhs31
			_128_v6 = _rhs32
			(_127_globalState).F15 = _228_i11
			(_127_globalState).F6 = _121_v2
			var _241_v73 _dafny.Map
			_ = _241_v73
			_241_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_239_v71, ((_137_v13).Cardinality()).Plus(_228_i11))
			_241_v73 = (_241_v73).Update(_239_v71, (_dafny.Zero).Minus(_119_v0))
		}
		var _242_v74 _dafny.Sequence
		_ = _242_v74
		_242_v74 = _dafny.SeqOf(_120_v1)
		_242_v74 = _242_v74
		var _243_v75 _dafny.Set
		_ = _243_v75
		_243_v75 = _dafny.SetOf(true)
		var _244_v76 _dafny.Set
		_ = _244_v76
		_244_v76 = _dafny.SetOf(_243_v75)
		var _245_v77 _dafny.Map
		_ = _245_v77
		_245_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v6, (_244_v76).Cardinality())
		_245_v77 = (_245_v77).Update(_128_v6, (_dafny.Zero).Minus(_119_v0))
		(_127_globalState).F15 = (_119_v0).Times(_228_i11)
	}
	var _246_i13 _dafny.Int
	_ = _246_i13
	_246_i13 = _dafny.Zero
	{
		for !(_139_v15) {
			{
				if (_246_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_246_i13 = (_246_i13).Plus(_dafny.One)
				var _247_v78 _dafny.Set
				_ = _247_v78
				var _248_v79 bool
				_ = _248_v79
				var _249_v80 _dafny.MultiSet
				_ = _249_v80
				var _out18 _dafny.Set
				_ = _out18
				var _out19 bool
				_ = _out19
				var _out20 _dafny.MultiSet
				_ = _out20
				_out18, _out19, _out20 = Companion_Default___.M0((func() _dafny.Int {
					if (_132_v11).Contains(_dafny.IntOfUint32(((_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence)).Cardinality())) {
						return (_132_v11).Get(_dafny.IntOfUint32(((_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence)).Cardinality())).(_dafny.Int)
					}
					return _119_v0
				})(), !(_139_v15), (_119_v0).Cmp(_119_v0) != 0, _120_v1, _127_globalState)
				_247_v78 = _out18
				_248_v79 = _out19
				_249_v80 = _out20
				var _250_v81 *C0
				_ = _250_v81
				var _nw31 *C0 = New_C0_()
				_ = _nw31
				_nw31.Ctor__(_119_v0, _128_v6)
				_250_v81 = _nw31
				var _251_v82 _dafny.Sequence
				_ = _251_v82
				_251_v82 = _dafny.SeqOf(_130_v8, Companion_Default___.Fm1((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vityokux")).Cardinality()), _dafny.IntOfUint32((_129_v7).Cardinality()), _119_v0, _127_globalState))
				var _252_v83 _dafny.Map
				_ = _252_v83
				_252_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(719), (_250_v81).F19()), _130_v8), _251_v82), !((_119_v0).Cmp(_119_v0) < 0))
				_252_v83 = (_252_v83).Update(_248_v79, (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool))
				var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_133_v12), 0))
				_ = _index24
				(_133_v12).ArraySet1(_119_v0, (_index24).Int())
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen((_133_v12), 0))
				_ = _index25
				(_133_v12).ArraySet1(_119_v0, (_index25).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _253_v84 _dafny.Set
	_ = _253_v84
	_253_v84 = _dafny.SetOf(_139_v15, (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool))
	var _hi3 _dafny.Int = Companion_Default___.SafeModuloInt(_119_v0, _119_v0)
	_ = _hi3
	for _254_i14 := ((_253_v84).Union(_253_v84)).Cardinality(); _254_i14.Cmp(_hi3) < 0; _254_i14 = _254_i14.Plus(_dafny.One) {
		if _139_v15 {
			_189_v47 = _189_v47
			(_127_globalState).F15 = (_dafny.Zero).Minus((_254_i14).Minus(_dafny.IntOfInt64(881)))
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
			_ = _index26
			(_124_v5).ArraySet1(!(!(_139_v15)), (_index26).Int())
			(_127_globalState).F13 = _254_i14
			var _255_v85 _dafny.Sequence
			_ = _255_v85
			_255_v85 = _dafny.SeqOf(_124_v5)
			var _256_v86 D3
			_ = _256_v86
			_256_v86 = Companion_D3_.Create_DC6_(_255_v85)
			(_127_globalState).F5 = _dafny.IntOfUint32(((_256_v86).Dtor_cf13()).Cardinality())
		} else {
			(_127_globalState).F6 = _121_v2
			_189_v47 = _189_v47
			(_127_globalState).F6 = (_dafny.IntOfInt64(-512)).Cmp(_119_v0) < 0
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
			_ = _index27
			(_124_v5).ArraySet1((_253_v84).IsDisjointFrom(_253_v84), (_index27).Int())
			var _257_v87 _dafny.Array
			_ = _257_v87
			var _nw32 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
			_ = _nw32
			_257_v87 = _nw32
			var _258_v88 _dafny.Map
			_ = _258_v88
			_258_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v0, _121_v2)
			var _259_v89 _dafny.Map
			_ = _259_v89
			_259_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6((_dafny.Zero).Minus((_258_v88).Cardinality()), _128_v6, (_dafny.Zero).Minus(_119_v0), _127_globalState), _257_v87)
			_257_v87 = (func() _dafny.Array {
				if (_259_v89).Contains((_dafny.SetOf(_121_v2)).IsProperSubsetOf(_253_v84)) {
					return (_259_v89).Get((_dafny.SetOf(_121_v2)).IsProperSubsetOf(_253_v84)).(_dafny.Array)
				}
				return _257_v87
			})()
		}
		(_127_globalState).F15 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_119_v0, _254_i14), _dafny.IntOfUint32((_122_v3).Cardinality()))
		var _260_v90 _dafny.Map
		_ = _260_v90
		_260_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (Companion_D0_.Create_DC1_(_133_v12, _124_v5, _123_v4, _119_v0, _254_i14)).Dtor_cf2())
		var _261_v91 _dafny.Map
		_ = _261_v91
		_261_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v6, (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence))
		var _262_v92 _dafny.Map
		_ = _262_v92
		_262_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_254_i14, _139_v15)
		var _263_v93 D0
		_ = _263_v93
		_263_v93 = Companion_D0_.Create_DC1_(_133_v12, (func() _dafny.Array {
			if (_260_v90).Contains((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)) {
				return (_260_v90).Get((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)).(_dafny.Array)
			}
			return _124_v5
		})(), _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}((func(_264_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_265_i15 _dafny.Int) _dafny.CodePoint {
				return _264_v6
			}
		})(_128_v6))), (func() _dafny.Sequence {
			if (_261_v91).Contains(_128_v6) {
				return (_261_v91).Get(_128_v6).(_dafny.Sequence)
			}
			return Companion_Default___.Fm7((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _254_i14, Companion_Default___.Fm0((func() bool {
				if (_262_v92).Contains(Companion_Default___.Fm0((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _120_v1, (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _254_i14, _127_globalState)) {
					return (_262_v92).Get(Companion_Default___.Fm0((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _120_v1, (_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _254_i14, _127_globalState)).(bool)
				}
				return !(_121_v2)
			})(), _dafny.UnicodeSeqOfUtf8Bytes("ytn"), _139_v15, _119_v0, _127_globalState), _127_globalState)
		})(), (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence), (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence)), _254_i14, _119_v0)
		_263_v93 = _263_v93
		var _266_v94 _dafny.Set
		_ = _266_v94
		var _267_v95 bool
		_ = _267_v95
		var _268_v96 _dafny.MultiSet
		_ = _268_v96
		var _out21 _dafny.Set
		_ = _out21
		var _out22 bool
		_ = _out22
		var _out23 _dafny.MultiSet
		_ = _out23
		_out21, _out22, _out23 = Companion_Default___.M0((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_254_i14, _119_v0)), _139_v15, !(_130_v8).Contains(_dafny.IntOfInt64(232)), _dafny.UnicodeSeqOfUtf8Bytes("ckhhe"), _127_globalState)
		_266_v94 = _out21
		_267_v95 = _out22
		_268_v96 = _out23
	}
	var _269_v97 _dafny.Array
	_ = _269_v97
	var _nw33 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
	_ = _nw33
	_269_v97 = _nw33
	var _270_v98 *C0
	_ = _270_v98
	var _nw34 *C0 = New_C0_()
	_ = _nw34
	_nw34.Ctor__(_119_v0, _128_v6)
	_270_v98 = _nw34
	var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_269_v97), 0))
	_ = _index28
	(_269_v97).ArraySet1(_270_v98, (_index28).Int())
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_269_v97), 0))
	_ = _index29
	(_269_v97).ArraySet1(_270_v98, (_index29).Int())
	var _271_v99 _dafny.Sequence
	_ = _271_v99
	_271_v99 = _dafny.SeqOf((_269_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_269_v97), 0))).Int()).(*C0))
	var _272_i16 _dafny.Int
	_ = _272_i16
	_272_i16 = _dafny.Zero
	{
		for !_dafny.Companion_Sequence_.Equal(_271_v99, _271_v99) {
			{
				if (_272_i16).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_272_i16 = (_272_i16).Plus(_dafny.One)
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))
				_ = _index30
				(_141_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence), (_141_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_141_v17), 0))).Int()).(_dafny.Sequence)), (_index30).Int())
				var _273_v100 _dafny.Set
				_ = _273_v100
				var _274_v101 bool
				_ = _274_v101
				var _275_v102 _dafny.MultiSet
				_ = _275_v102
				var _out24 _dafny.Set
				_ = _out24
				var _out25 bool
				_ = _out25
				var _out26 _dafny.MultiSet
				_ = _out26
				_out24, _out25, _out26 = Companion_Default___.M0((func() _dafny.Int {
					if (_137_v13).Contains(_119_v0) {
						return (_137_v13).Multiplicity(_119_v0)
					}
					return (_270_v98).F19()
				})(), !(_139_v15) || (!(true)), _121_v2, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_120_v1, _dafny.Companion_Sequence_.Update(_120_v1, (Companion_Default___.SafeIndex((_270_v98).F19(), _dafny.IntOfUint32((_120_v1).Cardinality()))).Uint32(), _128_v6)), (Companion_Default___.SafeIndex((_270_v98).F19(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_120_v1, _dafny.Companion_Sequence_.Update(_120_v1, (Companion_Default___.SafeIndex((_270_v98).F19(), _dafny.IntOfUint32((_120_v1).Cardinality()))).Uint32(), _128_v6))).Cardinality()))).Uint32(), _128_v6), _127_globalState)
				_273_v100 = _out24
				_274_v101 = _out25
				_275_v102 = _out26
				_274_v101 = _121_v2
				if _121_v2 {
					var _276_v103 _dafny.Array
					_ = _276_v103
					var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
					_ = _nw35
					_276_v103 = _nw35
					var _277_v104 _dafny.Sequence
					_ = _277_v104
					_277_v104 = _dafny.SeqOf(_124_v5)
					var _278_v105 D3
					_ = _278_v105
					_278_v105 = Companion_D3_.Create_DC6_(_277_v104)
					var _279_v106 _dafny.Map
					_ = _279_v106
					_279_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_278_v105, (_270_v98).F19())
					var _280_v107 _dafny.Map
					_ = _280_v107
					_280_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_279_v106, _139_v15)
					var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_276_v103), 0))
					_ = _index31
					(_276_v103).ArraySet1(_280_v107, (_index31).Int())
					var _281_v108 _dafny.Map
					_ = _281_v108
					_281_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _280_v107)
					var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_276_v103), 0))
					_ = _index32
					var _rhs33 _dafny.Map = ((func() _dafny.Map {
						if (_281_v108).Contains(false) {
							return (_281_v108).Get(false).(_dafny.Map)
						}
						return _280_v107
					})()).Merge(_280_v107)
					_ = _rhs33
					var _rhs34 bool = (_dafny.IntOfInt64(138)).Cmp(_dafny.IntOfInt64(62)) == 0
					_ = _rhs34
					var _lhs22 _dafny.Array = _276_v103
					_ = _lhs22
					var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(196), _dafny.ArrayLen((_276_v103), 0))
					_ = _lhs23
					var _lhs24 *GlobalState = _127_globalState
					_ = _lhs24
					(_lhs22).ArraySet1(_rhs33, (_lhs23).Int())
					_lhs24.F6 = _rhs34
					_141_v17 = _141_v17
					(_127_globalState).F5 = Companion_Default___.SafeModuloInt(((_270_v98).F19()).Plus((_270_v98).F19()), (_270_v98).F19())
					var _282_v109 D1
					_ = _282_v109
					_282_v109 = Companion_D1_.Create_DC2_((_270_v98).F20())
					var _pat_let_tv0 = _128_v6
					_ = _pat_let_tv0
					_282_v109 = func(_pat_let0_0 D1) D1 {
						return func(_283_dt__update__tmp_h1 D1) D1 {
							return func(_pat_let1_0 _dafny.CodePoint) D1 {
								return func(_284_dt__update_hcf6_h0 _dafny.CodePoint) D1 {
									return Companion_D1_.Create_DC2_(_284_dt__update_hcf6_h0)
								}(_pat_let1_0)
							}(_pat_let_tv0)
						}(_pat_let0_0)
					}(_282_v109)
					(_127_globalState).F15 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus((_270_v98).F19())).Plus(_dafny.IntOfInt64(259)), (_270_v98).Fm2(_119_v0, (_270_v98).F19(), (_270_v98).F19(), _127_globalState))))
				} else {
					_132_v11 = (_132_v11).Update(_dafny.IntOfUint32((_122_v3).Cardinality()), (_270_v98).F19())
					_129_v7 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool), _127_globalState), _129_v7)
					_120_v1 = _dafny.Companion_Sequence_.Concatenate(_120_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg11 _dafny.Int) interface{} {
							return coer11(arg11)
						}
					}((func(_285_v98 *C0) func(_dafny.Int) _dafny.CodePoint {
						return func(_286_i17 _dafny.Int) _dafny.CodePoint {
							return (_285_v98).F20()
						}
					})(_270_v98))))
					var _287_v110 _dafny.Array
					_ = _287_v110
					var _len0_10 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_10
					var _nw36 _dafny.Array
					_ = _nw36
					if _len0_10.Cmp(_dafny.Zero) == 0 {
						_nw36 = _dafny.NewArray(_len0_10)
					} else {
						var _init10 func(_dafny.Int) _dafny.Map = (func(_288_v2 bool, _289_v101 bool) func(_dafny.Int) _dafny.Map {
							return func(_290_i18 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_288_v2, _289_v101)
							}
						})(_121_v2, _274_v101)
						_ = _init10
						var _element0_10 = _init10(_dafny.Zero)
						_ = _element0_10
						_nw36 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
						(_nw36).ArraySet1(_element0_10, 0)
						var _nativeLen0_10 = (_len0_10).Int()
						_ = _nativeLen0_10
						for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
							(_nw36).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
						}
					}
					_287_v110 = _nw36
					var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
					_ = _index33
					var _rhs35 bool = ((_270_v98).F19()).Cmp(_119_v0) < 0
					_ = _rhs35
					var _rhs36 _dafny.Array = _287_v110
					_ = _rhs36
					var _lhs25 _dafny.Array = _124_v5
					_ = _lhs25
					var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))
					_ = _lhs26
					(_lhs25).ArraySet1(_rhs35, (_lhs26).Int())
					_287_v110 = _rhs36
					var _291_v111 _dafny.Set
					_ = _291_v111
					var _292_v112 bool
					_ = _292_v112
					var _293_v113 _dafny.MultiSet
					_ = _293_v113
					var _out27 _dafny.Set
					_ = _out27
					var _out28 bool
					_ = _out28
					var _out29 _dafny.MultiSet
					_ = _out29
					_out27, _out28, _out29 = Companion_Default___.M0((_270_v98).F19(), _274_v101, !((_124_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_124_v5), 0))).Int()).(bool)), _dafny.UnicodeSeqOfUtf8Bytes("svi"), _127_globalState)
					_291_v111 = _out27
					_292_v112 = _out28
					_293_v113 = _out29
				}
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_269_v97), 0))
	_ = _index34
	(_269_v97).ArraySet1((_269_v97).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen((_269_v97), 0))).Int()).(*C0), (_index34).Int())
	_dafny.Print(_119_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_120_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_121_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_122_v3, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v4).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qljvqw"), _dafny.UnicodeSeqOfUtf8Bytes("qljvqw"), _dafny.UnicodeSeqOfUtf8Bytes("qljvqw"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v5).ArrayGet1((_dafny.IntOfInt64(28)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_127_globalState).F7(), _dafny.SeqOf(_dafny.IntOfInt64(327), _dafny.IntOfInt64(6), _dafny.IntOfInt64(327), _dafny.IntOfInt64(-327))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_globalState).F8()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_127_globalState).F12()).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qljvqw"), _dafny.UnicodeSeqOfUtf8Bytes("qljvqw"), _dafny.UnicodeSeqOfUtf8Bytes("qljvqw"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_127_globalState).F14()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_128_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_129_v7, _dafny.SeqOf(_dafny.IntOfInt64(-349))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v8).Equals(_dafny.SetOf(_dafny.IntOfInt64(327), _dafny.IntOfInt64(-349))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v9).Equals(_dafny.MultiSetOf(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(176), _dafny.IntOfInt64(166), _dafny.IntOfInt64(-3)), _dafny.SetOf(_dafny.IntOfInt64(327), _dafny.IntOfInt64(-349)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_132_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(327), _dafny.IntOfInt64(-327))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v12).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v13).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(327), _dafny.IntOfInt64(327))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_138_v14).Equals(_dafny.SetOf(_dafny.SetOf(_dafny.IntOfInt64(327)), _dafny.SetOf(_dafny.IntOfInt64(327), _dafny.IntOfInt64(-29), _dafny.IntOfInt64(-28), _dafny.IntOfInt64(-27), _dafny.IntOfInt64(-26), _dafny.IntOfInt64(-25), _dafny.IntOfInt64(-24), _dafny.IntOfInt64(-23), _dafny.IntOfInt64(-22), _dafny.IntOfInt64(-21), _dafny.IntOfInt64(-20), _dafny.IntOfInt64(-19), _dafny.IntOfInt64(-18), _dafny.IntOfInt64(-17), _dafny.IntOfInt64(-16), _dafny.IntOfInt64(-15), _dafny.IntOfInt64(-14), _dafny.IntOfInt64(-13), _dafny.IntOfInt64(-12), _dafny.IntOfInt64(-11), _dafny.IntOfInt64(-10), _dafny.IntOfInt64(-9), _dafny.IntOfInt64(-8), _dafny.IntOfInt64(-7), _dafny.IntOfInt64(-6), _dafny.IntOfInt64(-5), _dafny.IntOfInt64(-4), _dafny.IntOfInt64(-3), _dafny.IntOfInt64(-2), _dafny.IntOfInt64(-1), _dafny.Zero, _dafny.One, _dafny.IntOfInt64(2), _dafny.IntOfInt64(3), _dafny.IntOfInt64(4), _dafny.IntOfInt64(5), _dafny.IntOfInt64(6), _dafny.IntOfInt64(7), _dafny.IntOfInt64(8), _dafny.IntOfInt64(9), _dafny.IntOfInt64(10), _dafny.IntOfInt64(11), _dafny.IntOfInt64(12), _dafny.IntOfInt64(13), _dafny.IntOfInt64(14), _dafny.IntOfInt64(15), _dafny.IntOfInt64(16), _dafny.IntOfInt64(17), _dafny.IntOfInt64(18), _dafny.IntOfInt64(19), _dafny.IntOfInt64(20), _dafny.IntOfInt64(21), _dafny.IntOfInt64(22), _dafny.IntOfInt64(23), _dafny.IntOfInt64(24), _dafny.IntOfInt64(25), _dafny.IntOfInt64(26), _dafny.IntOfInt64(27), _dafny.IntOfInt64(28), _dafny.IntOfInt64(29), _dafny.IntOfInt64(30), _dafny.IntOfInt64(31), _dafny.IntOfInt64(32), _dafny.IntOfInt64(33), _dafny.IntOfInt64(34), _dafny.IntOfInt64(35), _dafny.IntOfInt64(36), _dafny.IntOfInt64(37), _dafny.IntOfInt64(38), _dafny.IntOfInt64(39), _dafny.IntOfInt64(40), _dafny.IntOfInt64(41), _dafny.IntOfInt64(42), _dafny.IntOfInt64(43), _dafny.IntOfInt64(44), _dafny.IntOfInt64(45), _dafny.IntOfInt64(46), _dafny.IntOfInt64(47), _dafny.IntOfInt64(48), _dafny.IntOfInt64(49), _dafny.IntOfInt64(50), _dafny.IntOfInt64(51), _dafny.IntOfInt64(52), _dafny.IntOfInt64(53), _dafny.IntOfInt64(54), _dafny.IntOfInt64(55), _dafny.IntOfInt64(56), _dafny.IntOfInt64(57), _dafny.IntOfInt64(58), _dafny.IntOfInt64(59), _dafny.IntOfInt64(60), _dafny.IntOfInt64(61), _dafny.IntOfInt64(62), _dafny.IntOfInt64(63), _dafny.IntOfInt64(64), _dafny.IntOfInt64(65), _dafny.IntOfInt64(66), _dafny.IntOfInt64(67), _dafny.IntOfInt64(68), _dafny.IntOfInt64(69), _dafny.IntOfInt64(70), _dafny.IntOfInt64(71), _dafny.IntOfInt64(72), _dafny.IntOfInt64(73), _dafny.IntOfInt64(74), _dafny.IntOfInt64(75), _dafny.IntOfInt64(76), _dafny.IntOfInt64(77), _dafny.IntOfInt64(78), _dafny.IntOfInt64(79), _dafny.IntOfInt64(80), _dafny.IntOfInt64(81), _dafny.IntOfInt64(82), _dafny.IntOfInt64(83), _dafny.IntOfInt64(84), _dafny.IntOfInt64(85), _dafny.IntOfInt64(86), _dafny.IntOfInt64(87), _dafny.IntOfInt64(88), _dafny.IntOfInt64(89), _dafny.IntOfInt64(90), _dafny.IntOfInt64(91), _dafny.IntOfInt64(92)), _dafny.SetOf(_dafny.One, _dafny.IntOfInt64(176), _dafny.IntOfInt64(166), _dafny.IntOfInt64(-3)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_139_v15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_140_v16).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v17).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_144_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v47).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v47).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v47).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v47).Dtor_cf10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_i13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v84).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v97).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C0)).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_269_v97).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(*C0)).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v98).F19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_270_v98).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_271_v99).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_272_i16)
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
	Cf2 _dafny.Array
	Cf3 _dafny.MultiSet
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Array, Cf2 _dafny.Array, Cf3 _dafny.MultiSet, Cf4 _dafny.Int, Cf5 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
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
	return Companion_D0_.Create_DC1_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMultiSet, _dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.MultiSet {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf5
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + data.Cf0.VerbatimString(true) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3.Equals(data2.Cf3) && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
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
	Cf7  bool
	Cf8  bool
	Cf9  bool
	Cf10 bool
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 bool, Cf8 bool, Cf9 bool, Cf10 bool) D1 {
	return D1{D1_DC3{Cf7, Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf6 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf6 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf6}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC4 struct {
	Cf11 D1
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf11 D1) D1 {
	return D1{D1_DC4{Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false, false, false, false)
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC3).Cf8
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC3).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC3).Cf10
}

func (_this D1) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf6
}

func (_this D1) Dtor_cf11() D1 {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8 && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf6 == data2.Cf6
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf12 _dafny.Map
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf12 _dafny.Map) D2 {
	return D2{D2_DC5{Cf12}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D2) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D2_DC5).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf12.Equals(data2.Cf12)
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

type D3_DC8 struct {
	Cf14 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 _dafny.Int) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC6 struct {
	Cf13 _dafny.Sequence
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf13 _dafny.Sequence) D3 {
	return D3{D3_DC6{Cf13}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC7_()
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D3_DC6).Cf13
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC7:
		{
			return "D3.DC7"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf13) + ")"
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
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D3_DC6:
		{
			data2, ok := other.Get_().(D3_DC6)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf16 _dafny.Int
	Cf17 bool
	Cf18 bool
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf16 _dafny.Int, Cf17 bool, Cf18 bool) D4 {
	return D4{D4_DC10{Cf16, Cf17, Cf18}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC11 struct {
	Cf19 _dafny.CodePoint
	Cf20 bool
	Cf21 _dafny.MultiSet
	Cf22 bool
	Cf23 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf19 _dafny.CodePoint, Cf20 bool, Cf21 _dafny.MultiSet, Cf22 bool, Cf23 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf19, Cf20, Cf21, Cf22, Cf23}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf24 _dafny.Int
	Cf25 _dafny.Array
	Cf26 bool
	Cf27 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf24 _dafny.Int, Cf25 _dafny.Array, Cf26 bool, Cf27 _dafny.Int) D4 {
	return D4{D4_DC12{Cf24, Cf25, Cf26, Cf27}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC9 struct {
	Cf15 *C0
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf15 *C0) D4 {
	return D4{D4_DC9{Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(_dafny.Zero, false, false)
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC10).Cf17
}

func (_this D4) Dtor_cf18() bool {
	return _this.Get_().(D4_DC10).Cf18
}

func (_this D4) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) Dtor_cf20() bool {
	return _this.Get_().(D4_DC11).Cf20
}

func (_this D4) Dtor_cf21() _dafny.MultiSet {
	return _this.Get_().(D4_DC11).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC11).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf25
}

func (_this D4) Dtor_cf26() bool {
	return _this.Get_().(D4_DC12).Cf26
}

func (_this D4) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf27
}

func (_this D4) Dtor_cf15() *C0 {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
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
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17 == data2.Cf17 && data1.Cf18 == data2.Cf18
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf19 == data2.Cf19 && data1.Cf20 == data2.Cf20 && data1.Cf21.Equals(data2.Cf21) && data1.Cf22 == data2.Cf22 && data1.Cf23.Equals(data2.Cf23)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25 == data2.Cf25 && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15 == data2.Cf15
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

// Definition of class GlobalState
type GlobalState struct {
	F5   _dafny.Int
	F6   bool
	F13  _dafny.Int
	F15  _dafny.Int
	_f0  _dafny.Int
	_f1  _dafny.Int
	_f2  bool
	_f3  _dafny.Int
	_f4  bool
	_f7  _dafny.Sequence
	_f8  _dafny.Map
	_f9  _dafny.Int
	_f10 _dafny.Int
	_f11 _dafny.CodePoint
	_f12 _dafny.MultiSet
	_f14 _dafny.Sequence
	_f16 bool
	_f17 _dafny.Int
	_f18 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F5 = _dafny.Zero
	_this.F6 = false
	_this.F13 = _dafny.Zero
	_this.F15 = _dafny.Zero
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f2 = false
	_this._f3 = _dafny.Zero
	_this._f4 = false
	_this._f7 = _dafny.EmptySeq
	_this._f8 = _dafny.EmptyMap
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.CodePoint('D')
	_this._f12 = _dafny.EmptyMultiSet
	_this._f14 = _dafny.EmptySeq
	_this._f16 = false
	_this._f17 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 bool, f3 _dafny.Int, f4 bool, f5 _dafny.Int, f6 bool, f7 _dafny.Sequence, f8 _dafny.Map, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.CodePoint, f12 _dafny.MultiSet, f13 _dafny.Int, f14 _dafny.Sequence, f15 _dafny.Int, f16 bool, f17 _dafny.Int, f18 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F7() _dafny.Sequence {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Map {
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
func (_this *GlobalState) F11() _dafny.CodePoint {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.MultiSet {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Sequence {
	{
		return _this._f14
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
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f19 _dafny.Int
	_f20 _dafny.CodePoint
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f19 = _dafny.Zero
	_this._f20 = _dafny.CodePoint('D')
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

func (_this *C0) Ctor__(f19 _dafny.Int, f20 _dafny.CodePoint) {
	{
		(_this)._f19 = f19
		(_this)._f20 = f20
	}
}
func (_this *C0) Fm2(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F19()
	}
}
func (_this *C0) F19() _dafny.Int {
	{
		return _this._f19
	}
}
func (_this *C0) F20() _dafny.CodePoint {
	{
		return _this._f20
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
