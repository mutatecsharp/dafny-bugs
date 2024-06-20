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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(833), (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(565), _dafny.IntOfInt64(642))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(565)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(642)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_0_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("doerfd")).Cardinality())), (func() _dafny.Map {
					var _coll1 = _dafny.NewMapBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqOf(_dafny.CodePoint('f')))).Elements()); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _1_v1 _dafny.Sequence
						_1_v1 = interface{}(_compr_1).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.CodePoint('f')))).Contains(_1_v1) {
							_coll1.Add(_1_v1, _dafny.IntOfInt64(24))
						}
					}
					return _coll1.ToMap()
				}()).Cardinality())
			}
		}
		return _coll0.ToMap()
	}()).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('w')
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Sequence, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfInt64(-329), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality())), _dafny.IntOfInt64(696), (_dafny.SetOf(true, true, !(true))).Cardinality())).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(982))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("omngle")
	}))).Cardinality()))).Intersection(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-597)), !(false))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, globalState *GlobalState) bool {
	return ((func() _dafny.Int {
		if true {
			return _dafny.IntOfInt64(-92)
		}
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(605))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality())))).Cardinality())).Cardinality()
	})()).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("esnepq")).Cardinality()), _dafny.IntOfInt64(414))) <= 0
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aahphab")).Cardinality())).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-978)))) == 0, !(true), (_dafny.IntOfInt64(-747)).Cmp(_dafny.IntOfInt64(-992)) >= 0, (true) && (!(!(true))))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("jkgbfxm")
	_ = _source0
	var _4___mcc_h0 _dafny.Sequence = _source0
	_ = _4___mcc_h0
	var _5_cf5 _dafny.Sequence = _4___mcc_h0
	_ = _5_cf5
	return _5_cf5
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetOf(true, true, !(false), true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(-334))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cudaub")).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(650))))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(false)).Union(_dafny.MultiSetOf(true, false, true))).Difference(_dafny.MultiSetOf(true, true))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	(globalState).F2 = p0
	var _6_v0 bool
	_ = _6_v0
	_6_v0 = false
	var _7_v1 _dafny.Sequence
	_ = _7_v1
	_7_v1 = _dafny.SeqOf(_6_v0, _6_v0, _6_v0)
	_7_v1 = _7_v1
	var _8_v2 _dafny.Set
	_ = _8_v2
	_8_v2 = _dafny.SetOf(p1)
	var _9_i0 _dafny.Int
	_ = _9_i0
	_9_i0 = _dafny.Zero
	{
		for !(!(((_8_v2).IsSubsetOf(_8_v2)) || (true))) {
			{
				if (_9_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_9_i0 = (_9_i0).Plus(_dafny.One)
				var _10_v3 _dafny.CodePoint
				_ = _10_v3
				_10_v3 = _dafny.CodePoint('r')
				var _11_v4 _dafny.Map
				_ = _11_v4
				_11_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_v0, _10_v3)
				var _12_v5 _dafny.Array
				_ = _12_v5
				var _nwElement0_0 _dafny.CodePoint = _10_v3
				_ = _nwElement0_0
				var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(15))
				_ = _nw0
				(_nw0).ArraySet1CodePoint(_nwElement0_0, 0)
				(_nw0).ArraySet1CodePoint(Companion_Default___.Fm1(globalState), 1)
				(_nw0).ArraySet1CodePoint(_10_v3, 2)
				(_nw0).ArraySet1CodePoint(_10_v3, 3)
				(_nw0).ArraySet1CodePoint(_10_v3, 4)
				(_nw0).ArraySet1CodePoint(_10_v3, 5)
				(_nw0).ArraySet1CodePoint(_10_v3, 6)
				(_nw0).ArraySet1CodePoint(_10_v3, 7)
				(_nw0).ArraySet1CodePoint((func() _dafny.CodePoint {
					if _6_v0 {
						return _10_v3
					}
					return _10_v3
				})(), 8)
				(_nw0).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_11_v4).Contains(_6_v0) {
						return (_11_v4).Get(_6_v0).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('b')
				})(), 9)
				(_nw0).ArraySet1CodePoint(_10_v3, 10)
				(_nw0).ArraySet1CodePoint(_10_v3, 11)
				(_nw0).ArraySet1CodePoint(_10_v3, 12)
				(_nw0).ArraySet1CodePoint(_10_v3, 13)
				(_nw0).ArraySet1CodePoint(_10_v3, 14)
				_12_v5 = _nw0
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_12_v5), 0))
				_ = _index0
				(_12_v5).ArraySet1CodePoint(_10_v3, (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_12_v5), 0))
				_ = _index1
				(_12_v5).ArraySet1CodePoint(_10_v3, (_index1).Int())
				r1 = Companion_Default___.Fm0(_6_v0, p0, globalState)
				var _13_v6 *C0
				_ = _13_v6
				var _nw1 *C0 = New_C0_()
				_ = _nw1
				_nw1.Ctor__()
				_13_v6 = _nw1
				var _14_v7 _dafny.MultiSet
				_ = _14_v7
				_14_v7 = _dafny.MultiSetOf(true, _6_v0, _6_v0)
				_6_v0 = (((_14_v7).Union(_14_v7)).Cardinality()).Cmp(p0) != 0
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _15_v8 _dafny.Sequence
	_ = _15_v8
	_15_v8 = _dafny.UnicodeSeqOfUtf8Bytes("bhpgj")
	var _hi0 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(380), _dafny.IntOfInt64(60))
	_ = _hi0
	for _16_i1 := _dafny.IntOfUint32((_15_v8).Cardinality()); _16_i1.Cmp(_hi0) < 0; _16_i1 = _16_i1.Plus(_dafny.One) {
		var _17_v9 _dafny.Array
		_ = _17_v9
		var _nwElement0_1 _dafny.Int = _16_i1
		_ = _nwElement0_1
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_1, 0)
		_17_v9 = _nw2
		var _18_v10 _dafny.Sequence
		_ = _18_v10
		_18_v10 = _dafny.SeqOf(_17_v9, _17_v9, _17_v9)
		(globalState).F2 = _dafny.IntOfUint32((_18_v10).Cardinality())
		_15_v8 = _dafny.Companion_Sequence_.Concatenate(_15_v8, _15_v8)
		var _19_v11 _dafny.CodePoint
		_ = _19_v11
		_19_v11 = _dafny.CodePoint('n')
		_15_v8 = _dafny.Companion_Sequence_.Concatenate(_15_v8, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm6(!(false), _6_v0, globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm6(!(false), _6_v0, globalState)).Cardinality()))).Uint32(), _19_v11))
		if _6_v0 {
			var _20_v12 _dafny.Map
			_ = _20_v12
			_20_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(830)).Times(_16_i1), _15_v8)
			_20_v12 = (_20_v12).Update(p1, _15_v8)
			(globalState).F2 = p1
			var _21_v13 *C0
			_ = _21_v13
			var _nw3 *C0 = New_C0_()
			_ = _nw3
			_nw3.Ctor__()
			_21_v13 = _nw3
			var _22_v14 _dafny.Array
			_ = _22_v14
			var _nwElement0_2 *C0 = _21_v13
			_ = _nwElement0_2
			var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(24))
			_ = _nw4
			(_nw4).ArraySet1(_nwElement0_2, 0)
			(_nw4).ArraySet1(_21_v13, 1)
			(_nw4).ArraySet1(_21_v13, 2)
			(_nw4).ArraySet1(_21_v13, 3)
			(_nw4).ArraySet1(_21_v13, 4)
			(_nw4).ArraySet1(_21_v13, 5)
			(_nw4).ArraySet1(_21_v13, 6)
			(_nw4).ArraySet1(_21_v13, 7)
			(_nw4).ArraySet1(_21_v13, 8)
			(_nw4).ArraySet1(_21_v13, 9)
			(_nw4).ArraySet1((func() *C0 {
				if _6_v0 {
					return _21_v13
				}
				return _21_v13
			})(), 10)
			(_nw4).ArraySet1(_21_v13, 11)
			(_nw4).ArraySet1(_21_v13, 12)
			(_nw4).ArraySet1(_21_v13, 13)
			(_nw4).ArraySet1(_21_v13, 14)
			(_nw4).ArraySet1(_21_v13, 15)
			(_nw4).ArraySet1((func() *C0 {
				if _6_v0 {
					return _21_v13
				}
				return _21_v13
			})(), 16)
			(_nw4).ArraySet1(_21_v13, 17)
			(_nw4).ArraySet1(_21_v13, 18)
			(_nw4).ArraySet1(_21_v13, 19)
			(_nw4).ArraySet1(_21_v13, 20)
			(_nw4).ArraySet1(_21_v13, 21)
			(_nw4).ArraySet1((func() *C0 {
				if _6_v0 {
					return _21_v13
				}
				return _21_v13
			})(), 22)
			(_nw4).ArraySet1(_21_v13, 23)
			_22_v14 = _nw4
			var _nw5 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
			_ = _nw5
			_22_v14 = _nw5
			var _23_v15 _dafny.MultiSet
			_ = _23_v15
			_23_v15 = _dafny.MultiSetOf(_6_v0, _6_v0)
			(globalState).F2 = (func() _dafny.Int {
				if (_23_v15).Contains(!(!_dafny.Companion_Sequence_.Contains(_15_v8, _19_v11))) {
					return (_23_v15).Multiplicity(!(!_dafny.Companion_Sequence_.Contains(_15_v8, _19_v11)))
				}
				return _16_i1
			})()
			var _24_v16 _dafny.Sequence
			_ = _24_v16
			_24_v16 = _dafny.SeqOf(p1, Companion_Default___.SafeDivisionInt(p0, p0))
			_24_v16 = Companion_Default___.Fm7(_6_v0, (func() _dafny.Int {
				if _6_v0 {
					return _dafny.IntOfInt64(569)
				}
				return _16_i1
			})(), globalState)
		} else {
			var _25_v17 _dafny.Array
			_ = _25_v17
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_0
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_26_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_27_i2 _dafny.Int) _dafny.Sequence {
						return _26_v8
					}
				})(_15_v8)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw6 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw6).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw6).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_25_v17 = _nw6
			var _28_v18 D0
			_ = _28_v18
			_28_v18 = Companion_D0_.Create_DC2_(_6_v0, _25_v17, p0)
			var _29_v19 _dafny.Map
			_ = _29_v19
			_29_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_28_v18).Dtor_cf2())
			var _30_v20 _dafny.MultiSet
			_ = _30_v20
			_30_v20 = _dafny.MultiSetOf(_16_i1, _16_i1, p1)
			_6_v0 = (func() bool {
				if (_29_v19).Contains(Companion_Default___.SafeModuloInt(_16_i1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_15_v8, (Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_30_v20).Contains(_dafny.IntOfUint32((_15_v8).Cardinality())) {
						return (_30_v20).Multiplicity(_dafny.IntOfUint32((_15_v8).Cardinality()))
					}
					return p0
				})(), _dafny.IntOfUint32((_15_v8).Cardinality()))).Uint32(), _19_v11)).Cardinality()))) {
					return (_29_v19).Get(Companion_Default___.SafeModuloInt(_16_i1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_15_v8, (Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_30_v20).Contains(_dafny.IntOfUint32((_15_v8).Cardinality())) {
							return (_30_v20).Multiplicity(_dafny.IntOfUint32((_15_v8).Cardinality()))
						}
						return p0
					})(), _dafny.IntOfUint32((_15_v8).Cardinality()))).Uint32(), _19_v11)).Cardinality()))).(bool)
				}
				return _6_v0
			})()
			var _31_v21 _dafny.Sequence
			_ = _31_v21
			_31_v21 = _dafny.SeqOf(Companion_Default___.Fm0(_6_v0, p0, globalState), p1)
			var _32_v22 _dafny.Set
			_ = _32_v22
			_32_v22 = _dafny.SetOf(_19_v11, _19_v11)
			var _33_v23 _dafny.Sequence
			_ = _33_v23
			_33_v23 = _dafny.SeqOf(_31_v21)
			var _34_v24 D0
			_ = _34_v24
			_34_v24 = Companion_D0_.Create_DC1_(p0, _dafny.Companion_Sequence_.Update(_31_v21, (Companion_Default___.SafeIndex((_32_v22).Cardinality(), _dafny.IntOfUint32((_31_v21).Cardinality()))).Uint32(), _dafny.IntOfUint32((_33_v23).Cardinality())))
			(globalState).F2 = (_34_v24).Dtor_cf0()
			_6_v0 = _6_v0
			var _35_v25 D0
			_ = _35_v25
			_35_v25 = Companion_D0_.Create_DC0_()
			var _rhs0 _dafny.Int = _16_i1
			_ = _rhs0
			var _rhs1 D0 = _35_v25
			_ = _rhs1
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			_lhs0.F2 = _rhs0
			_35_v25 = _rhs1
			var _36_v26 _dafny.MultiSet
			_ = _36_v26
			_36_v26 = _dafny.MultiSetOf(_6_v0)
			_6_v0 = ((Companion_Default___.Fm8(globalState)).Intersection(_36_v26)).IsDisjointFrom((_36_v26).Union(_dafny.MultiSetOf(_6_v0)))
		}
	}
	var _37_v27 _dafny.Array
	_ = _37_v27
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
	_ = _nw7
	_37_v27 = _nw7
	var _38_v28 _dafny.Sequence
	_ = _38_v28
	_38_v28 = _dafny.SeqOf(p0, p0)
	var _39_v29 D0
	_ = _39_v29
	_39_v29 = Companion_D0_.Create_DC2_(_6_v0, _37_v27, (Companion_D0_.Create_DC1_(p1, _38_v28)).Dtor_cf0())
	var _40_i3 _dafny.Int
	_ = _40_i3
	_40_i3 = _dafny.Zero
	{
		for !((_39_v29).Dtor_cf2()) {
			{
				if (_40_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_40_i3 = (_40_i3).Plus(_dafny.One)
				var _41_v30 _dafny.Array
				_ = _41_v30
				var _nw8 _dafny.Array = _dafny.NewArray(_dafny.One)
				_ = _nw8
				_41_v30 = _nw8
				var _42_v31 *C0
				_ = _42_v31
				var _nw9 *C0 = New_C0_()
				_ = _nw9
				_nw9.Ctor__()
				_42_v31 = _nw9
				var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_41_v30), 0))
				_ = _index2
				(_41_v30).ArraySet1(_42_v31, (_index2).Int())
				var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_41_v30), 0))
				_ = _index3
				(_41_v30).ArraySet1(_42_v31, (_index3).Int())
				var _43_v32 D0
				_ = _43_v32
				_43_v32 = Companion_D0_.Create_DC1_(p1, _38_v28)
				(globalState).F2 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_6_v0, _6_v0)).Cardinality()), (p1).Minus((_43_v32).Dtor_cf0()))
				_38_v28 = _38_v28
				var _44_v33 _dafny.Map
				_ = _44_v33
				_44_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_38_v28, _dafny.Companion_Sequence_.Concatenate(_15_v8, _dafny.UnicodeSeqOfUtf8Bytes("odihkxdux")))
				r1 = (_44_v33).Cardinality()
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _45_v35 _dafny.Array
	_ = _45_v35
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(29)
	_ = _len0_1
	var _nw10 _dafny.Array
	_ = _nw10
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw10 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Set = (func(_46_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
			return func(_47_i5 _dafny.Int) _dafny.Set {
				return (_dafny.SetOf(_dafny.CodePoint('b'))).Difference(func() _dafny.Set {
					var _coll2 = _dafny.NewBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate((_dafny.MultiSetFromSeq(_46_v8)).Elements()); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _48_v34 _dafny.CodePoint
						_48_v34 = interface{}(_compr_2).(_dafny.CodePoint)
						if (_dafny.MultiSetFromSeq(_46_v8)).Contains(_48_v34) {
							_coll2.Add(_48_v34)
						}
					}
					return _coll2.ToSet()
				}())
			}
		})(_15_v8)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw10 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw10).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw10).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_45_v35 = _nw10
	for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_45_v35), 0))); ; {
		_guard_loop_0, _ok3 := _iter3()
		if !_ok3 {
			break
		}
		var _49_i4 _dafny.Int
		_49_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_49_i4).Sign() != -1) && ((_49_i4).Cmp(_dafny.ArrayLen((_45_v35), 0)) < 0)) {
			(_45_v35).ArraySet1((_dafny.SetOf(_dafny.CodePoint('b'))).Difference(_dafny.SetOf(_dafny.CodePoint('f'))), (_49_i4).Int())
		}
	}
	var _50_v36 _dafny.Array
	_ = _50_v36
	var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
	_ = _nw11
	_50_v36 = _nw11
	r0 = _50_v36
	r1 = p1
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _51_v0 _dafny.Sequence
	_ = _51_v0
	_51_v0 = _dafny.UnicodeSeqOfUtf8Bytes("rhvwcv")
	var _52_globalState *GlobalState
	_ = _52_globalState
	var _nw12 *GlobalState = New_GlobalState_()
	_ = _nw12
	_nw12.Ctor__(_dafny.Companion_Sequence_.Concatenate(_51_v0, _51_v0), true, _dafny.IntOfInt64(411), _51_v0)
	_52_globalState = _nw12
	var _53_v1 _dafny.Int
	_ = _53_v1
	_53_v1 = _dafny.IntOfInt64(387)
	(_52_globalState).F2 = _53_v1
	var _54_v2 bool
	_ = _54_v2
	_54_v2 = false
	_54_v2 = _54_v2
	if _54_v2 {
		var _55_v3 _dafny.Array
		_ = _55_v3
		var _56_v4 _dafny.Int
		_ = _56_v4
		var _out0 _dafny.Array
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		_out0, _out1 = Companion_Default___.M0(_53_v1, (_53_v1).Minus(_53_v1), _52_globalState)
		_55_v3 = _out0
		_56_v4 = _out1
		var _57_v5 _dafny.MultiSet
		_ = _57_v5
		_57_v5 = _dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(-8), _56_v4))
		var _58_v6 _dafny.Sequence
		_ = _58_v6
		_58_v6 = _dafny.SeqOf(_57_v5, _57_v5)
		_54_v2 = ((_58_v6).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_56_v4), _dafny.IntOfUint32((_58_v6).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_57_v5)
		var _59_v7 _dafny.Array
		_ = _59_v7
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_2
		var _nw13 _dafny.Array
		_ = _nw13
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw13 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Map = (func(_60_v2 bool, _61_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
				return func(_62_i0 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_60_v2, _dafny.IntOfUint32((_61_v0).Cardinality()))
				}
			})(_54_v2, _51_v0)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw13 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw13).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw13).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_59_v7 = _nw13
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_55_v3), 0))
		_ = _index4
		(_55_v3).ArraySet1(_54_v2, (_index4).Int())
		var _63_v8 _dafny.Map
		_ = _63_v8
		_63_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v1, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_51_v0).Cardinality()), _56_v4))
		var _64_v9 _dafny.Set
		_ = _64_v9
		_64_v9 = _dafny.SetOf(_53_v1, _56_v4, _53_v1)
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_55_v3), 0))
		_ = _index5
		var _rhs2 _dafny.Int = (_63_v8).Cardinality()
		_ = _rhs2
		var _rhs3 _dafny.Array = _59_v7
		_ = _rhs3
		var _rhs4 bool = (_64_v9).IsDisjointFrom((_64_v9).Intersection(_64_v9))
		_ = _rhs4
		var _lhs1 *GlobalState = _52_globalState
		_ = _lhs1
		var _lhs2 _dafny.Array = _55_v3
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_55_v3), 0))
		_ = _lhs3
		_lhs1.F2 = _rhs2
		_59_v7 = _rhs3
		(_lhs2).ArraySet1(_rhs4, (_lhs3).Int())
		var _65_v10 _dafny.Array
		_ = _65_v10
		var _66_v11 _dafny.Int
		_ = _66_v11
		var _out2 _dafny.Array
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out2, _out3 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt(_53_v1, _56_v4), _53_v1, _52_globalState)
		_65_v10 = _out2
		_66_v11 = _out3
		_51_v0 = _dafny.UnicodeSeqOfUtf8Bytes("nwstnawx")
	} else {
		var _67_v12 _dafny.Array
		_ = _67_v12
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw14
		_67_v12 = _nw14
		var _68_v13 _dafny.Map
		_ = _68_v13
		_68_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v1, _67_v12)
		var _69_v14 _dafny.Array
		_ = _69_v14
		var _nwElement0_3 _dafny.Array = (func() _dafny.Array {
			if (_68_v13).Contains(_53_v1) {
				return (_68_v13).Get(_53_v1).(_dafny.Array)
			}
			return _67_v12
		})()
		_ = _nwElement0_3
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.One)
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_3, 0)
		_69_v14 = _nw15
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_69_v14), 0))
		_ = _index6
		(_69_v14).ArraySet1(_67_v12, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_69_v14), 0))
		_ = _index7
		(_69_v14).ArraySet1(_67_v12, (_index7).Int())
		(_52_globalState).F2 = _53_v1
		var _70_v15 _dafny.Array
		_ = _70_v15
		var _71_v16 _dafny.Int
		_ = _71_v16
		var _out4 _dafny.Array
		_ = _out4
		var _out5 _dafny.Int
		_ = _out5
		_out4, _out5 = Companion_Default___.M0(_53_v1, (_dafny.Zero).Minus(_dafny.IntOfInt64(-856)), _52_globalState)
		_70_v15 = _out4
		_71_v16 = _out5
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_69_v14), 0))
		_ = _index8
		(_69_v14).ArraySet1(_67_v12, (_index8).Int())
		var _72_v17 _dafny.Map
		_ = _72_v17
		_72_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v15, (_71_v16).Cmp(_71_v16) >= 0)
		_72_v17 = (_72_v17).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v15, false))
	}
	var _hi1 _dafny.Int = _53_v1
	_ = _hi1
	for _73_i1 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("suibphoto"), _51_v0)).Cardinality()); _73_i1.Cmp(_hi1) < 0; _73_i1 = _73_i1.Plus(_dafny.One) {
		var _74_v18 _dafny.Array
		_ = _74_v18
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_3
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_75_v1 _dafny.Int, _76_i1 _dafny.Int) func(_dafny.Int) bool {
				return func(_77_i2 _dafny.Int) bool {
					return (_75_v1).Cmp(_76_i1) != 0
				}
			})(_53_v1, _73_i1)
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
		_74_v18 = _nw16
		_74_v18 = _74_v18
		_53_v1 = _73_i1
		var _78_v19 _dafny.Sequence
		_ = _78_v19
		_78_v19 = _dafny.SeqOf(false, _54_v2)
		var _79_v20 _dafny.Array
		_ = _79_v20
		var _80_v21 _dafny.Int
		_ = _80_v21
		var _out6 _dafny.Array
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out6, _out7 = Companion_Default___.M0((_dafny.IntOfInt64(663)).Times(Companion_Default___.Fm0(_54_v2, _dafny.IntOfUint32((_78_v19).Cardinality()), _52_globalState)), _73_i1, _52_globalState)
		_79_v20 = _out6
		_80_v21 = _out7
		var _rhs5 bool = _54_v2
		_ = _rhs5
		var _rhs6 _dafny.Int = _dafny.IntOfInt64(-337)
		_ = _rhs6
		var _lhs4 *GlobalState = _52_globalState
		_ = _lhs4
		_54_v2 = _rhs5
		_lhs4.F2 = _rhs6
	}
	var _81_v22 _dafny.Set
	_ = _81_v22
	_81_v22 = _dafny.SetOf(_54_v2, _54_v2, true)
	var _82_v23 _dafny.Array
	_ = _82_v23
	var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
	_ = _nw17
	_82_v23 = _nw17
	var _83_v24 D0
	_ = _83_v24
	_83_v24 = Companion_D0_.Create_DC2_(!(_81_v22).Equals(_81_v22), _82_v23, _53_v1)
	var _source1 D0 = _83_v24
	_ = _source1
	if _source1.Is_DC0() {
		var _84_v25 _dafny.Map
		_ = _84_v25
		_84_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v1, _54_v2)
		var _85_v26 _dafny.Sequence
		_ = _85_v26
		_85_v26 = _dafny.SeqOf(_54_v2)
		var _86_v27 _dafny.Map
		_ = _86_v27
		_86_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v26, true)
		var _87_v28 _dafny.Array
		_ = _87_v28
		var _nwElement0_4 _dafny.Int = _53_v1
		_ = _nwElement0_4
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_4, 0)
		(_nw18).ArraySet1(Companion_Default___.Fm0(_54_v2, _53_v1, _52_globalState), 1)
		(_nw18).ArraySet1(_dafny.IntOfInt64(-526), 2)
		(_nw18).ArraySet1(_dafny.IntOfInt64(-912), 3)
		(_nw18).ArraySet1(_53_v1, 4)
		(_nw18).ArraySet1(_53_v1, 5)
		(_nw18).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-894)), 6)
		(_nw18).ArraySet1(_53_v1, 7)
		(_nw18).ArraySet1((_53_v1).Minus((_84_v25).Cardinality()), 8)
		(_nw18).ArraySet1((_86_v27).Cardinality(), 9)
		(_nw18).ArraySet1((func() _dafny.Int {
			if true {
				return _53_v1
			}
			return _53_v1
		})(), 10)
		(_nw18).ArraySet1(_53_v1, 11)
		_87_v28 = _nw18
		var _88_v29 _dafny.Array
		_ = _88_v29
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_4
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.CodePoint = func(_89_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw19 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw19).ArraySet1CodePoint(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw19).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_88_v29 = _nw19
		var _90_v30 _dafny.Map
		_ = _90_v30
		_90_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v29, _53_v1)
		var _91_v31 _dafny.Map
		_ = _91_v31
		_91_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v1, (_90_v30).Cardinality())
		var _92_v32 _dafny.Map
		_ = _92_v32
		_92_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _91_v31)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))
		_ = _index9
		(_87_v28).ArraySet1(((func() _dafny.Map {
			if (_92_v32).Contains(_54_v2) {
				return (_92_v32).Get(_54_v2).(_dafny.Map)
			}
			return func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(955), _dafny.IntOfInt64(328))); ; {
					_compr_3, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _93_v33 _dafny.Int
					_93_v33 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(955)).Cmp(_93_v33) <= 0) && ((_93_v33).Cmp(_dafny.IntOfInt64(328)) < 0) {
						_coll3.Add((_93_v33).Times(_53_v1), _53_v1)
					}
				}
				return _coll3.ToMap()
			}()
		})()).Cardinality(), (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))
		_ = _index10
		(_87_v28).ArraySet1(Companion_Default___.Fm0(_54_v2, _dafny.IntOfInt64(348), _52_globalState), (_index10).Int())
		var _94_v34 _dafny.CodePoint
		_ = _94_v34
		_94_v34 = _dafny.CodePoint('o')
		var _95_v35 _dafny.Array
		_ = _95_v35
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw20
		_95_v35 = _nw20
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_95_v35), 0))
		_ = _index11
		(_95_v35).ArraySet1(_54_v2, (_index11).Int())
		var _96_v36 _dafny.Map
		_ = _96_v36
		_96_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D0 {
			if _54_v2 {
				return _83_v24
			}
			return Companion_D0_.Create_DC2_(false, _82_v23, _53_v1)
		})(), _94_v34)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))
		_ = _index12
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_95_v35), 0))
		_ = _index13
		var _rhs7 bool = _54_v2
		_ = _rhs7
		var _rhs8 _dafny.Int = Companion_Default___.SafeModuloInt((_87_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))).Int()).(_dafny.Int), (_87_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))).Int()).(_dafny.Int))
		_ = _rhs8
		var _rhs9 _dafny.CodePoint = (func() _dafny.CodePoint {
			if (_96_v36).Contains((func() D0 {
				if false {
					return Companion_D0_.Create_DC2_(_54_v2, _82_v23, (_87_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))).Int()).(_dafny.Int))
				}
				return _83_v24
			})()) {
				return (_96_v36).Get((func() D0 {
					if false {
						return Companion_D0_.Create_DC2_(_54_v2, _82_v23, (_87_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))).Int()).(_dafny.Int))
					}
					return _83_v24
				})()).(_dafny.CodePoint)
			}
			return Companion_Default___.Fm1(_52_globalState)
		})()
		_ = _rhs9
		var _rhs10 bool = !((func() bool {
			if !(_54_v2) {
				return _54_v2
			}
			return _54_v2
		})()) || (_54_v2)
		_ = _rhs10
		var _lhs5 _dafny.Array = _87_v28
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))
		_ = _lhs6
		var _lhs7 _dafny.Array = _95_v35
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_95_v35), 0))
		_ = _lhs8
		_54_v2 = _rhs7
		(_lhs5).ArraySet1(_rhs8, (_lhs6).Int())
		_94_v34 = _rhs9
		(_lhs7).ArraySet1(_rhs10, (_lhs8).Int())
		var _97_v37 *C0
		_ = _97_v37
		var _nw21 *C0 = New_C0_()
		_ = _nw21
		_nw21.Ctor__()
		_97_v37 = _nw21
		var _98_v38 _dafny.Sequence
		_ = _98_v38
		_98_v38 = _dafny.SeqOf(_dafny.IntOfInt64(-296), _53_v1)
		var _rhs11 _dafny.Sequence = _51_v0
		_ = _rhs11
		var _rhs12 _dafny.Int = (_98_v38).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_87_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_87_v28), 0))).Int()).(_dafny.Int), _53_v1), _dafny.IntOfUint32((_98_v38).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs12
		var _lhs9 *GlobalState = _52_globalState
		_ = _lhs9
		_51_v0 = _rhs11
		_lhs9.F2 = _rhs12
	} else if _source1.Is_DC1() {
		var _99___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC1).Cf0
		_ = _99___mcc_h0
		var _100___mcc_h1 _dafny.Sequence = _source1.Get_().(D0_DC1).Cf1
		_ = _100___mcc_h1
		var _101_cf1 _dafny.Sequence = _100___mcc_h1
		_ = _101_cf1
		var _102_cf0 _dafny.Int = _99___mcc_h0
		_ = _102_cf0
		var _103_v39 _dafny.Array
		_ = _103_v39
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
		_ = _nw22
		_103_v39 = _nw22
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_103_v39), 0))
		_ = _index14
		(_103_v39).ArraySet1(_53_v1, (_index14).Int())
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_103_v39), 0))
		_ = _index15
		(_103_v39).ArraySet1((_102_cf0).Minus(_102_cf0), (_index15).Int())
		_54_v2 = _54_v2
		_103_v39 = _103_v39
		_54_v2 = _54_v2
	} else {
		var _104___mcc_h2 bool = _source1.Get_().(D0_DC2).Cf2
		_ = _104___mcc_h2
		var _105___mcc_h3 _dafny.Array = _source1.Get_().(D0_DC2).Cf3
		_ = _105___mcc_h3
		var _106___mcc_h4 _dafny.Int = _source1.Get_().(D0_DC2).Cf4
		_ = _106___mcc_h4
		var _107_cf4 _dafny.Int = _106___mcc_h4
		_ = _107_cf4
		var _108_cf3 _dafny.Array = _105___mcc_h3
		_ = _108_cf3
		var _109_cf2 bool = _104___mcc_h2
		_ = _109_cf2
		_51_v0 = _51_v0
		_109_cf2 = (_83_v24).Dtor_cf2()
		(_52_globalState).F2 = (_53_v1).Plus((_53_v1).Minus(_53_v1))
		var _110_v41 *C0
		_ = _110_v41
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__()
		_110_v41 = _nw23
		var _111_v42 _dafny.Map
		_ = _111_v42
		_111_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(927), _dafny.IntOfInt64(-180))); ; {
				_compr_4, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _112_v40 _dafny.Int
				_112_v40 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(927)).Cmp(_112_v40) <= 0) && ((_112_v40).Cmp(_dafny.IntOfInt64(-180)) < 0) {
					_coll4.Add(Companion_Default___.SafeDivisionInt(_112_v40, _53_v1), _54_v2)
				}
			}
			return _coll4.ToMap()
		}(), _110_v41)
		var _113_v43 _dafny.Array
		_ = _113_v43
		var _114_v44 _dafny.Int
		_ = _114_v44
		var _out8 _dafny.Array
		_ = _out8
		var _out9 _dafny.Int
		_ = _out9
		_out8, _out9 = Companion_Default___.M0((_111_v42).Cardinality(), (_dafny.Zero).Minus((func() _dafny.Int {
			if _109_cf2 {
				return _53_v1
			}
			return _dafny.IntOfInt64(506)
		})()), _52_globalState)
		_113_v43 = _out8
		_114_v44 = _out9
	}
	var _115_v45 *C0
	_ = _115_v45
	var _nw24 *C0 = New_C0_()
	_ = _nw24
	_nw24.Ctor__()
	_115_v45 = _nw24
	var _116_v46 _dafny.MultiSet
	_ = _116_v46
	_116_v46 = _dafny.MultiSetOf(_115_v45, _115_v45)
	var _117_v47 _dafny.Sequence
	_ = _117_v47
	_117_v47 = _dafny.SeqOf(_53_v1, _53_v1, _53_v1, (_116_v46).Cardinality())
	var _118_v48 _dafny.MultiSet
	_ = _118_v48
	_118_v48 = _dafny.MultiSetOf(_dafny.IntOfInt64(-254), (_117_v47).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_51_v0).Cardinality()), _dafny.IntOfUint32((_117_v47).Cardinality()))).Uint32()).(_dafny.Int), _53_v1)
	var _hi2 _dafny.Int = (func() _dafny.Int {
		if (_118_v48).Contains(_dafny.IntOfInt64(-779)) {
			return (_118_v48).Multiplicity(_dafny.IntOfInt64(-779))
		}
		return _53_v1
	})()
	_ = _hi2
	for _119_i4 := _53_v1; _119_i4.Cmp(_hi2) < 0; _119_i4 = _119_i4.Plus(_dafny.One) {
		var _120_v49 _dafny.Sequence
		_ = _120_v49
		_120_v49 = _dafny.SeqOf(_54_v2)
		var _121_v50 _dafny.Map
		_ = _121_v50
		_121_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v49, _54_v2)
		var _122_v51 _dafny.CodePoint
		_ = _122_v51
		_122_v51 = _dafny.CodePoint('s')
		_51_v0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_123_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _51_v0), (Companion_Default___.SafeIndex(((_121_v50).Cardinality()).Plus(_119_i4), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_124_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _51_v0)).Cardinality()))).Uint32(), _122_v51), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_53_v1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_125_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _51_v0), (Companion_Default___.SafeIndex(((_121_v50).Cardinality()).Plus(_119_i4), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_126_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})), _51_v0)).Cardinality()))).Uint32(), _122_v51)).Cardinality()))).Uint32(), _dafny.CodePoint('l'))
		var _127_v52 _dafny.Map
		_ = _127_v52
		_127_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_53_v1).Plus(_119_i4), _dafny.IntOfInt64(807))
		_127_v52 = (_127_v52).Update(((_dafny.SetOf(_119_i4, _119_i4, _53_v1)).Cardinality()).Plus(_53_v1), Companion_Default___.SafeDivisionInt(_119_i4, _53_v1))
		var _128_v53 _dafny.Array
		_ = _128_v53
		var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw25
		_128_v53 = _nw25
		var _129_v54 _dafny.Set
		_ = _129_v54
		_129_v54 = _dafny.SetOf(_128_v53, _128_v53)
		_54_v2 = (_129_v54).IsSubsetOf(_129_v54)
		_54_v2 = (_120_v49).Select((Companion_Default___.SafeIndex(_53_v1, _dafny.IntOfUint32((_120_v49).Cardinality()))).Uint32()).(bool)
	}
	var _130_v55 _dafny.Array
	_ = _130_v55
	var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
	_ = _nw26
	_130_v55 = _nw26
	_130_v55 = _130_v55
	var _131_v56 _dafny.Sequence
	_ = _131_v56
	_131_v56 = _dafny.SeqOf(_54_v2, _54_v2)
	if _dafny.Companion_Sequence_.Contains(_117_v47, (_53_v1).Times(_dafny.IntOfUint32((_131_v56).Cardinality()))) {
		_53_v1 = (func() _dafny.Int {
			if false {
				return _53_v1
			}
			return (_53_v1).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_132_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))).Cardinality()))
		})()
		var _133_v57 _dafny.Map
		_ = _133_v57
		_133_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_52_globalState), _54_v2)
		var _134_v58 _dafny.CodePoint
		_ = _134_v58
		_134_v58 = _dafny.CodePoint('a')
		_54_v2 = !((func() bool {
			if (_133_v57).Contains(_134_v58) {
				return (_133_v57).Get(_134_v58).(bool)
			}
			return !(_dafny.Companion_Sequence_.Equal(_51_v0, _51_v0))
		})())
		(_52_globalState).F2 = _53_v1
		_83_v24 = _83_v24
		var _135_v59 _dafny.Map
		_ = _135_v59
		_135_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(993), _53_v1)
		_54_v2 = (_dafny.IntOfInt64(774)).Cmp(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_51_v0).Cardinality()))), (_135_v59).Cardinality())) <= 0
	} else {
		var _136_v60 *C0
		_ = _136_v60
		var _nw27 *C0 = New_C0_()
		_ = _nw27
		_nw27.Ctor__()
		_136_v60 = _nw27
		_54_v2 = (_dafny.IntOfUint32((_117_v47).Cardinality())).Cmp(_53_v1) != 0
		var _137_v61 _dafny.Array
		_ = _137_v61
		var _138_v62 _dafny.Int
		_ = _138_v62
		var _out10 _dafny.Array
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out10, _out11 = Companion_Default___.M0(_53_v1, _53_v1, _52_globalState)
		_137_v61 = _out10
		_138_v62 = _out11
		if _54_v2 {
			var _139_v63 _dafny.Array
			_ = _139_v63
			var _nwElement0_5 _dafny.Int = Companion_Default___.Fm0(!(_54_v2), _53_v1, _52_globalState)
			_ = _nwElement0_5
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_5, 0)
			(_nw28).ArraySet1(_53_v1, 1)
			(_nw28).ArraySet1(_53_v1, 2)
			(_nw28).ArraySet1(_53_v1, 3)
			(_nw28).ArraySet1((_138_v62).Plus(_53_v1), 4)
			(_nw28).ArraySet1((_138_v62).Minus(_53_v1), 5)
			(_nw28).ArraySet1(_53_v1, 6)
			(_nw28).ArraySet1(_53_v1, 7)
			(_nw28).ArraySet1((_dafny.IntOfInt64(10)).Plus(_53_v1), 8)
			(_nw28).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fatyy")).Cardinality()), 9)
			(_nw28).ArraySet1(_138_v62, 10)
			(_nw28).ArraySet1(_53_v1, 11)
			(_nw28).ArraySet1(_138_v62, 12)
			(_nw28).ArraySet1(_53_v1, 13)
			(_nw28).ArraySet1(_53_v1, 14)
			(_nw28).ArraySet1(_53_v1, 15)
			(_nw28).ArraySet1((_138_v62).Times((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(_54_v2, (_81_v22).Cardinality(), _52_globalState)))), 16)
			(_nw28).ArraySet1((_dafny.IntOfUint32((_131_v56).Cardinality())).Times(_dafny.IntOfInt64(7)), 17)
			(_nw28).ArraySet1((_dafny.Zero).Minus(_138_v62), 18)
			(_nw28).ArraySet1((_dafny.Zero).Minus(_138_v62), 19)
			(_nw28).ArraySet1(Companion_Default___.SafeModuloInt(_138_v62, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iasqrm")).Cardinality())), 20)
			(_nw28).ArraySet1((_dafny.Zero).Minus(_138_v62), 21)
			(_nw28).ArraySet1(_53_v1, 22)
			(_nw28).ArraySet1((_117_v47).Select((Companion_Default___.SafeIndex(_53_v1, _dafny.IntOfUint32((_117_v47).Cardinality()))).Uint32()).(_dafny.Int), 23)
			(_nw28).ArraySet1(_138_v62, 24)
			_139_v63 = _nw28
			_139_v63 = _139_v63
			_54_v2 = !((_118_v48).IsSubsetOf((Companion_Default___.Fm3(_dafny.SeqOf(_53_v1), _54_v2, _54_v2, _54_v2, _52_globalState)).Difference(_118_v48)))
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_139_v63), 0))
			_ = _index16
			(_139_v63).ArraySet1(_53_v1, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_139_v63), 0))
			_ = _index17
			(_139_v63).ArraySet1((_dafny.SetOf((func() bool {
				if _54_v2 {
					return _54_v2
				}
				return true
			})(), ((_118_v48).Update(_138_v62, Companion_Default___.Abs(_53_v1))).IsSubsetOf(_118_v48), _dafny.Companion_Sequence_.IsPrefixOf(_51_v0, _51_v0))).Cardinality(), (_index17).Int())
			var _140_v64 _dafny.Map
			_ = _140_v64
			_140_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v2, Companion_Default___.Fm4(_54_v2, _52_globalState))
			var _141_v65 _dafny.Map
			_ = _141_v65
			_141_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_54_v2, _138_v62, _52_globalState), _54_v2)
			_54_v2 = (func() bool {
				if (_140_v64).Contains(((_141_v65).Cardinality()).Cmp((_139_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_139_v63), 0))).Int()).(_dafny.Int)) >= 0) {
					return (_140_v64).Get(((_141_v65).Cardinality()).Cmp((_139_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_139_v63), 0))).Int()).(_dafny.Int)) >= 0).(bool)
				}
				return !(_54_v2)
			})()
			var _142_v66 *C0
			_ = _142_v66
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__()
			_142_v66 = _nw29
		} else {
			var _143_v67 _dafny.Array
			_ = _143_v67
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw30
			_143_v67 = _nw30
			var _144_v68 _dafny.Array
			_ = _144_v68
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_5
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_145_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_146_i7 _dafny.Int) _dafny.Int {
						return (_146_i7).Minus(_145_v1)
					}
				})(_53_v1)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw31 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw31).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw31).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_144_v68 = _nw31
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))
			_ = _index18
			(_144_v68).ArraySet1(_53_v1, (_index18).Int())
			var _147_v69 _dafny.Map
			_ = _147_v69
			_147_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_51_v0).Cardinality()), (_81_v22).Cardinality())
			var _148_v70 _dafny.Map
			_ = _148_v70
			_148_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v61, _147_v69)
			var _149_v71 _dafny.Map
			_ = _149_v71
			_149_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v68, _143_v67)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))
			_ = _index19
			var _rhs13 bool = !(_148_v70).Equals(_148_v70)
			_ = _rhs13
			var _rhs14 bool = _54_v2
			_ = _rhs14
			var _rhs15 _dafny.Array = (func() _dafny.Array {
				if (_149_v71).Contains(_144_v68) {
					return (_149_v71).Get(_144_v68).(_dafny.Array)
				}
				return _143_v67
			})()
			_ = _rhs15
			var _rhs16 _dafny.Int = (_dafny.Zero).Minus(_138_v62)
			_ = _rhs16
			var _lhs10 _dafny.Array = _144_v68
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))
			_ = _lhs11
			_54_v2 = _rhs13
			_54_v2 = _rhs14
			_143_v67 = _rhs15
			(_lhs10).ArraySet1(_rhs16, (_lhs11).Int())
			var _150_v72 _dafny.Set
			_ = _150_v72
			_150_v72 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_131_v56, (Companion_Default___.SafeIndex(_53_v1, _dafny.IntOfUint32((_131_v56).Cardinality()))).Uint32(), _54_v2), _131_v56), _dafny.Companion_Sequence_.Concatenate(_131_v56, _131_v56))
			_150_v72 = _dafny.SetOf(_131_v56, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5((_144_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))).Int()).(_dafny.Int), _54_v2, _54_v2, _52_globalState), _dafny.SeqOf(_54_v2)), _131_v56)
			var _151_v73 _dafny.CodePoint
			_ = _151_v73
			_151_v73 = _dafny.CodePoint('a')
			var _152_v74 _dafny.Map
			_ = _152_v74
			_152_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v48, _151_v73)
			_152_v74 = (_152_v74).Update(_118_v48, (func() _dafny.CodePoint {
				if false {
					return _151_v73
				}
				return _151_v73
			})())
			var _153_v75 _dafny.Array
			_ = _153_v75
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(24))
			_ = _nw32
			_153_v75 = _nw32
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_153_v75), 0))
			_ = _index20
			(_153_v75).ArraySet1(_83_v24, (_index20).Int())
			var _pat_let_tv0 = _82_v23
			_ = _pat_let_tv0
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(483), _dafny.ArrayLen((_153_v75), 0))
			_ = _index21
			(_153_v75).ArraySet1(func(_pat_let0_0 D0) D0 {
				return func(_154_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 _dafny.Array) D0 {
						return func(_155_dt__update_hcf3_h0 _dafny.Array) D0 {
							return Companion_D0_.Create_DC2_((_154_dt__update__tmp_h0).Dtor_cf2(), _155_dt__update_hcf3_h0, (_154_dt__update__tmp_h0).Dtor_cf4())
						}(_pat_let1_0)
					}(_pat_let_tv0)
				}(_pat_let0_0)
			}(_83_v24), (_index21).Int())
			var _156_v76 _dafny.Array
			_ = _156_v76
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_6
			var _nw33 _dafny.Array
			_ = _nw33
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw33 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.CodePoint = func(_157_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw33 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw33).ArraySet1CodePoint(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw33).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_156_v76 = _nw33
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_156_v76), 0))
			_ = _index22
			(_156_v76).ArraySet1CodePoint(_151_v73, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_156_v76), 0))
			_ = _index23
			var _rhs17 _dafny.CodePoint = _151_v73
			_ = _rhs17
			var _rhs18 _dafny.MultiSet = (_118_v48).Update((_144_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))).Int()).(_dafny.Int), Companion_Default___.Abs((_144_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))).Int()).(_dafny.Int)))
			_ = _rhs18
			var _rhs19 _dafny.Int = (_144_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_144_v68), 0))).Int()).(_dafny.Int)
			_ = _rhs19
			var _lhs12 _dafny.Array = _156_v76
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_156_v76), 0))
			_ = _lhs13
			(_lhs12).ArraySet1CodePoint(_rhs17, (_lhs13).Int())
			_118_v48 = _rhs18
			_138_v62 = _rhs19
		}
		var _158_v77 _dafny.MultiSet
		_ = _158_v77
		_158_v77 = _dafny.MultiSetOf(!(_54_v2))
		_158_v77 = _158_v77
	}
	_115_v45 = _115_v45
	var _159_v78 _dafny.CodePoint
	_ = _159_v78
	_159_v78 = _dafny.CodePoint('i')
	var _rhs20 bool = _54_v2
	_ = _rhs20
	var _rhs21 bool = _54_v2
	_ = _rhs21
	var _rhs22 _dafny.CodePoint = _159_v78
	_ = _rhs22
	var _rhs23 _dafny.Int = ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_53_v1, _53_v1))).Intersection((func() _dafny.MultiSet {
		if _54_v2 {
			return _118_v48
		}
		return _118_v48
	})())).Cardinality()
	_ = _rhs23
	_54_v2 = _rhs20
	_54_v2 = _rhs21
	_159_v78 = _rhs22
	_53_v1 = _rhs23
	var _160_v79 D0
	_ = _160_v79
	_160_v79 = Companion_D0_.Create_DC0_()
	var _161_v80 _dafny.Array
	_ = _161_v80
	var _nwElement0_6 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_53_v1, _160_v79)).Cardinality()
	_ = _nwElement0_6
	var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
	_ = _nw34
	(_nw34).ArraySet1(_nwElement0_6, 0)
	_161_v80 = _nw34
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_161_v80), 0))); ; {
		_guard_loop_1, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _162_i9 _dafny.Int
		_162_i9 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_162_i9).Sign() != -1) && ((_162_i9).Cmp(_dafny.ArrayLen((_161_v80), 0)) < 0)) {
			(_161_v80).ArraySet1(Companion_Default___.SafeModuloInt(_162_i9, _53_v1), (_162_i9).Int())
		}
	}
	var _163_v81 *C0
	_ = _163_v81
	var _nw35 *C0 = New_C0_()
	_ = _nw35
	_nw35.Ctor__()
	_163_v81 = _nw35
	_54_v2 = _54_v2
	var _164_v82 _dafny.Map
	_ = _164_v82
	_164_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v2, _53_v1)
	if (func() bool {
		if _54_v2 {
			return !(_164_v82).Equals(_164_v82)
		}
		return !(_54_v2)
	})() {
		var _165_v83 _dafny.Array
		_ = _165_v83
		var _166_v84 _dafny.Int
		_ = _166_v84
		var _out12 _dafny.Array
		_ = _out12
		var _out13 _dafny.Int
		_ = _out13
		_out12, _out13 = Companion_Default___.M0(_53_v1, _53_v1, _52_globalState)
		_165_v83 = _out12
		_166_v84 = _out13
		_54_v2 = !(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_167_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_168_i10 _dafny.Int) _dafny.CodePoint {
				return _167_v78
			}
		})(_159_v78))), _51_v0))
		var _169_v85 _dafny.Array
		_ = _169_v85
		var _170_v86 _dafny.Int
		_ = _170_v86
		var _out14 _dafny.Array
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out14, _out15 = Companion_Default___.M0(Companion_Default___.SafeDivisionInt((_118_v48).Cardinality(), _53_v1), _53_v1, _52_globalState)
		_169_v85 = _out14
		_170_v86 = _out15
		_54_v2 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_51_v0, _51_v0), _159_v78)
		var _171_v87 *C0
		_ = _171_v87
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__()
		_171_v87 = _nw36
	} else {
		_53_v1 = _53_v1
		var _rhs24 _dafny.Int = (_dafny.Zero).Minus(_53_v1)
		_ = _rhs24
		var _rhs25 _dafny.Int = _53_v1
		_ = _rhs25
		var _rhs26 _dafny.Map = _164_v82
		_ = _rhs26
		var _lhs14 *GlobalState = _52_globalState
		_ = _lhs14
		_lhs14.F2 = _rhs24
		_53_v1 = _rhs25
		_164_v82 = _rhs26
		var _172_v88 _dafny.Array
		_ = _172_v88
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw37
		_172_v88 = _nw37
		var _173_v89 _dafny.Array
		_ = _173_v89
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_7
		var _nw38 _dafny.Array
		_ = _nw38
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw38 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.CodePoint = (func(_174_v78 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_175_i11 _dafny.Int) _dafny.CodePoint {
					return _174_v78
				}
			})(_159_v78)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw38 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw38).ArraySet1CodePoint(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw38).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_173_v89 = _nw38
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_172_v88), 0))
		_ = _index24
		(_172_v88).ArraySet1(_173_v89, (_index24).Int())
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_172_v88), 0))
		_ = _index25
		(_172_v88).ArraySet1(_173_v89, (_index25).Int())
		var _176_v90 _dafny.Array
		_ = _176_v90
		var _177_v91 _dafny.Int
		_ = _177_v91
		var _out16 _dafny.Array
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		_out16, _out17 = Companion_Default___.M0(_dafny.IntOfUint32((_51_v0).Cardinality()), (_53_v1).Plus(_53_v1), _52_globalState)
		_176_v90 = _out16
		_177_v91 = _out17
		_51_v0 = _51_v0
	}
	var _178_v92 _dafny.Array
	_ = _178_v92
	var _nw39 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
	_ = _nw39
	_178_v92 = _nw39
	_178_v92 = _178_v92
	var _179_v93 _dafny.Sequence
	_ = _179_v93
	_179_v93 = _dafny.SeqOf(_161_v80)
	var _180_v94 _dafny.Map
	_ = _180_v94
	_180_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v47, _dafny.Companion_Sequence_.Concatenate(_179_v93, _179_v93))
	_180_v94 = (_180_v94).Update(_117_v47, _179_v93)
	_dafny.Print(_51_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_52_globalState).F0().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_52_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_52_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_52_globalState).F3().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_53_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_54_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v22).Equals(_dafny.SetOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_v24).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_v24).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v46).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_117_v47, _dafny.SeqOf(_dafny.IntOfInt64(386), _dafny.IntOfInt64(386), _dafny.IntOfInt64(386), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v48).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-254), _dafny.IntOfInt64(386), _dafny.IntOfInt64(386))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_131_v56, _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_v78)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v80).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_164_v82).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_179_v93).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v94).Cardinality())
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
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_() D0 {
	return D0{D0_DC0{}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf0 _dafny.Int
	Cf1 _dafny.Sequence
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf0 _dafny.Int, Cf1 _dafny.Sequence) D0 {
	return D0{D0_DC1{Cf0, Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 bool
	Cf3 _dafny.Array
	Cf4 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 bool, Cf3 _dafny.Array, Cf4 _dafny.Int) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_()
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Array {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
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
			_, ok := other.Get_().(D0_DC0)
			return ok
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1.Equals(data2.Cf1)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf5 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + data.Cf5.VerbatimString(true) + ")"
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

// Definition of class GlobalState
type GlobalState struct {
	F2  _dafny.Int
	_f0 _dafny.Sequence
	_f1 bool
	_f3 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this._f0 = _dafny.EmptySeq
	_this._f1 = false
	_this._f3 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 bool, f2 _dafny.Int, f3 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
	}
}
func (_this *GlobalState) F0() _dafny.Sequence {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
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
func (_this *C0) Fm2(p0 _dafny.Int, p1 D0, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ieh"), _dafny.UnicodeSeqOfUtf8Bytes("ewuihk")), _dafny.UnicodeSeqOfUtf8Bytes("h"))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
