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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(617))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	})), _dafny.UnicodeSeqOfUtf8Bytes("ocmloxxh")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nwbfxj"), _dafny.UnicodeSeqOfUtf8Bytes("ctmkds")))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	return (((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(428), true)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-698), !(false))
	})()).Cardinality()).Cmp(((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(907), _dafny.IntOfInt64(423))); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _1_v0 _dafny.Int
				_1_v0 = interface{}(_compr_1).(_dafny.Int)
				if ((_dafny.IntOfInt64(907)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(423)) < 0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_1_v0, (_dafny.MultiSetOf(_dafny.IntOfInt64(849), _dafny.IntOfInt64(974))).Cardinality()), _dafny.CodePoint('v'))
				}
			}
			return _coll1.ToMap()
		}()).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_0).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(907), _dafny.IntOfInt64(423))); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _1_v0 _dafny.Int
					_1_v0 = interface{}(_compr_2).(_dafny.Int)
					if ((_dafny.IntOfInt64(907)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(423)) < 0) {
						_coll2.Add(Companion_Default___.SafeModuloInt(_1_v0, (_dafny.MultiSetOf(_dafny.IntOfInt64(849), _dafny.IntOfInt64(974))).Cardinality()), _dafny.CodePoint('v'))
					}
				}
				return _coll2.ToMap()
			}()).Contains(_2_v1) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_2_v1, _dafny.IntOfInt64(47)))
			}
		}
		return _coll0.ToSet()
	}()).Union(_dafny.SetOf(_dafny.IntOfInt64(956), _dafny.IntOfInt64(701)))).Cardinality()) == 0
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if ((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _4_v0 _dafny.CodePoint
			_4_v0 = interface{}(_compr_3).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), _4_v0) {
				_coll3.Add(_4_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vaujf")).Cardinality()))).Cardinality()))
			}
		}
		return _coll3.ToMap()
	}()).Cardinality()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uat")).Cardinality())))) == 0 {
		return _dafny.CodePoint('d')
	} else if false {
		return _dafny.CodePoint('u')
	} else {
		return _dafny.CodePoint('w')
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(false)))).Cardinality(), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), !(false))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-421), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(532), false)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rnwxeunx")).Cardinality()))).Cardinality()), _dafny.IntOfInt64(-846), _dafny.IntOfInt64(246), _dafny.IntOfInt64(938)))).Cardinality()).Cmp(_dafny.IntOfInt64(-324)) >= 0)
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if !(true) {
			return _dafny.SeqOf(true, false, true, true, false)
		}
		return _dafny.SeqOf(true)
	})(), _dafny.SeqOf(true))).Cardinality())
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 bool, globalState *GlobalState) {
	var _5_i0 _dafny.Int
	_ = _5_i0
	_5_i0 = _dafny.Zero
	{
		for p1 {
			{
				if (_5_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_5_i0 = (_5_i0).Plus(_dafny.One)
				var _6_v0 _dafny.Array
				_ = _6_v0
				var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw0
				_6_v0 = _nw0
				_6_v0 = _6_v0
				var _7_v1 _dafny.Int
				_ = _7_v1
				_7_v1 = _dafny.IntOfInt64(35)
				var _8_v2 _dafny.Array
				_ = _8_v2
				var _len0_0 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_0
				var _nw1 _dafny.Array
				_ = _nw1
				if _len0_0.Cmp(_dafny.Zero) == 0 {
					_nw1 = _dafny.NewArray(_len0_0)
				} else {
					var _init0 func(_dafny.Int) _dafny.CodePoint = func(_9_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('j')
					}
					_ = _init0
					var _element0_0 = _init0(_dafny.Zero)
					_ = _element0_0
					_nw1 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
					(_nw1).ArraySet1CodePoint(_element0_0, 0)
					var _nativeLen0_0 = (_len0_0).Int()
					_ = _nativeLen0_0
					for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
						(_nw1).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
					}
				}
				_8_v2 = _nw1
				var _10_v3 _dafny.Map
				_ = _10_v3
				_10_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v1, _dafny.SeqOf(_8_v2))
				var _11_v4 _dafny.Map
				_ = _11_v4
				_11_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v1, p1)
				var _12_v5 _dafny.Sequence
				_ = _12_v5
				_12_v5 = _dafny.SeqOf(_8_v2, _8_v2)
				var _13_v6 *C0
				_ = _13_v6
				var _nw2 *C0 = New_C0_()
				_ = _nw2
				_nw2.Ctor__((func() _dafny.Sequence {
					if (_10_v3).Contains((_dafny.Zero).Minus(Companion_Default___.Fm5((func() bool {
						if (_11_v4).Contains((_dafny.Zero).Minus(_7_v1)) {
							return (_11_v4).Get((_dafny.Zero).Minus(_7_v1)).(bool)
						}
						return p0
					})(), globalState))) {
						return (_10_v3).Get((_dafny.Zero).Minus(Companion_Default___.Fm5((func() bool {
							if (_11_v4).Contains((_dafny.Zero).Minus(_7_v1)) {
								return (_11_v4).Get((_dafny.Zero).Minus(_7_v1)).(bool)
							}
							return p0
						})(), globalState))).(_dafny.Sequence)
					}
					return _12_v5
				})())
				_13_v6 = _nw2
				var _14_v7 bool
				_ = _14_v7
				_14_v7 = true
				var _15_v8 _dafny.Sequence
				_ = _15_v8
				_15_v8 = _dafny.UnicodeSeqOfUtf8Bytes("k")
				_14_v7 = !((_7_v1).Cmp(_dafny.IntOfUint32((_15_v8).Cardinality())) >= 0)
				_14_v7 = p1
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _16_v9 _dafny.Sequence
	_ = _16_v9
	_16_v9 = _dafny.UnicodeSeqOfUtf8Bytes("peut")
	var _17_v10 _dafny.Array
	_ = _17_v10
	var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
	_ = _nw3
	_17_v10 = _nw3
	var _18_v11 _dafny.Sequence
	_ = _18_v11
	_18_v11 = _dafny.SeqOf(_17_v10, _17_v10, _17_v10, _17_v10, _17_v10)
	var _19_v12 *C0
	_ = _19_v12
	var _nw4 *C0 = New_C0_()
	_ = _nw4
	_nw4.Ctor__(_18_v11)
	_19_v12 = _nw4
	var _20_v13 _dafny.Map
	_ = _20_v13
	_20_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-877))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_21_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})), _16_v9), _19_v12)
	_20_v13 = (_20_v13).Update(p0, _19_v12)
	var _22_v14 D2
	_ = _22_v14
	_22_v14 = Companion_D2_.Create_DC9_(p1)
	var _pat_let_tv0 = _16_v9
	_ = _pat_let_tv0
	var _pat_let_tv1 = _16_v9
	_ = _pat_let_tv1
	var _pat_let_tv2 = _16_v9
	_ = _pat_let_tv2
	var _pat_let_tv3 = _16_v9
	_ = _pat_let_tv3
	_16_v9 = func(_source0 D2) _dafny.Sequence {
		if _source0.Is_DC7() {
			var _23___mcc_h0 _dafny.Sequence = _source0.Get_().(D2_DC7).Cf11
			_ = _23___mcc_h0
			var _24___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC7).Cf12
			_ = _24___mcc_h1
			var _25___mcc_h2 _dafny.Int = _source0.Get_().(D2_DC7).Cf13
			_ = _25___mcc_h2
			var _26___mcc_h3 bool = _source0.Get_().(D2_DC7).Cf14
			_ = _26___mcc_h3
			var _27_cf14 bool = _26___mcc_h3
			_ = _27_cf14
			var _28_cf13 _dafny.Int = _25___mcc_h2
			_ = _28_cf13
			var _29_cf12 _dafny.Int = _24___mcc_h1
			_ = _29_cf12
			var _30_cf11 _dafny.Sequence = _23___mcc_h0
			_ = _30_cf11
			return _dafny.UnicodeSeqOfUtf8Bytes("ba")
		} else if _source0.Is_DC8() {
			var _31___mcc_h4 bool = _source0.Get_().(D2_DC8).Cf15
			_ = _31___mcc_h4
			var _32___mcc_h5 T0 = _source0.Get_().(D2_DC8).Cf16
			_ = _32___mcc_h5
			var _33_cf16 T0 = _32___mcc_h5
			_ = _33_cf16
			var _34_cf15 bool = _31___mcc_h4
			_ = _34_cf15
			return _pat_let_tv0
		} else if _source0.Is_DC9() {
			var _35___mcc_h6 bool = _source0.Get_().(D2_DC9).Cf17
			_ = _35___mcc_h6
			var _36_cf17 bool = _35___mcc_h6
			_ = _36_cf17
			return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv1, _pat_let_tv2)
		} else if _source0.Is_DC6() {
			var _37___mcc_h7 _dafny.Sequence = _source0.Get_().(D2_DC6).Cf10
			_ = _37___mcc_h7
			var _38_cf10 _dafny.Sequence = _37___mcc_h7
			_ = _38_cf10
			return _dafny.Companion_Sequence_.Concatenate(_38_cf10, _pat_let_tv3)
		} else {
			var _39___mcc_h8 D2 = _source0.Get_().(D2_DC10).Cf18
			_ = _39___mcc_h8
			var _40_cf18 D2 = _39___mcc_h8
			_ = _40_cf18
			return _dafny.UnicodeSeqOfUtf8Bytes("eohfqvsl")
		}
	}(_22_v14)
	_16_v9 = _16_v9
	if !(!(p0) || (p0)) || (p0) {
		var _41_v15 bool
		_ = _41_v15
		_41_v15 = true
		var _42_v16 _dafny.Int
		_ = _42_v16
		_42_v16 = _dafny.IntOfInt64(348)
		var _43_v17 _dafny.MultiSet
		_ = _43_v17
		_43_v17 = _dafny.MultiSetOf(_42_v16)
		_41_v15 = (_43_v17).Contains((_dafny.IntOfUint32((_16_v9).Cardinality())).Times(_42_v16))
		_41_v15 = p1
		_41_v15 = _41_v15
		var _44_v18 T0
		_ = _44_v18
		var _nw5 *C0 = New_C0_()
		_ = _nw5
		_nw5.Ctor__(_18_v11)
		_44_v18 = _nw5
		var _45_v19 D1
		_ = _45_v19
		_45_v19 = Companion_D1_.Create_DC3_(_44_v18)
		var _46_v20 D1
		_ = _46_v20
		_46_v20 = Companion_D1_.Create_DC5_(_45_v19)
		var _47_v21 _dafny.Map
		_ = _47_v21
		_47_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v15, _dafny.SeqOf(_46_v20))
		var _48_v22 _dafny.Sequence
		_ = _48_v22
		_48_v22 = _dafny.SeqOf(Companion_D1_.Create_DC5_(_45_v19), _46_v20, _46_v20, Companion_D1_.Create_DC5_(_45_v19), _46_v20)
		var _49_v23 _dafny.Sequence
		_ = _49_v23
		_49_v23 = _dafny.SeqOf(p1, p1)
		var _50_v24 _dafny.Sequence
		_ = _50_v24
		_50_v24 = _dafny.SeqOf((func() _dafny.Sequence {
			if (_47_v21).Contains(!(p0)) {
				return (_47_v21).Get(!(p0)).(_dafny.Sequence)
			}
			return _48_v22
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_46_v20), _48_v22), (func() _dafny.Sequence {
			if (_49_v23).Select((Companion_Default___.SafeIndex(_42_v16, _dafny.IntOfUint32((_49_v23).Cardinality()))).Uint32()).(bool) {
				return _48_v22
			}
			return _48_v22
		})())
		_42_v16 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_50_v24).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm5(_41_v15, globalState), _dafny.IntOfUint32((_50_v24).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())))
		_44_v18 = _44_v18
	} else {
		var _51_v25 _dafny.Int
		_ = _51_v25
		_51_v25 = _dafny.IntOfInt64(534)
		var _52_v26 bool
		_ = _52_v26
		_52_v26 = true
		var _53_v27 _dafny.Map
		_ = _53_v27
		_53_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
		var _rhs0 _dafny.Int = _51_v25
		_ = _rhs0
		var _rhs1 _dafny.Int = (func() _dafny.Int {
			if (p0) || ((func() bool {
				if (_53_v27).Contains(_52_v26) {
					return (_53_v27).Get(_52_v26).(bool)
				}
				return p0
			})()) {
				return _51_v25
			}
			return _51_v25
		})()
		_ = _rhs1
		var _rhs2 bool = true
		_ = _rhs2
		_51_v25 = _rhs0
		_51_v25 = _rhs1
		_52_v26 = _rhs2
		if (func() bool {
			if _52_v26 {
				return !(p1) || (p1)
			}
			return (Companion_Default___.Fm5(_52_v26, globalState)).Cmp(_51_v25) < 0
		})() {
			var _54_v28 _dafny.Array
			_ = _54_v28
			var _nw6 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
			_ = _nw6
			_54_v28 = _nw6
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_54_v28), 0))
			_ = _index0
			(_54_v28).ArraySet1(_19_v12, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(184), _dafny.ArrayLen((_54_v28), 0))
			_ = _index1
			(_54_v28).ArraySet1(_19_v12, (_index1).Int())
			var _55_v29 _dafny.Sequence
			_ = _55_v29
			_55_v29 = _dafny.SeqOf(_20_v13, _20_v13)
			_51_v25 = _dafny.IntOfUint32((_55_v29).Cardinality())
			_51_v25 = _51_v25
			_52_v26 = p0
			var _56_v30 _dafny.Array
			_ = _56_v30
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_1
			var _nw7 _dafny.Array
			_ = _nw7
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw7 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_57_v25 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_58_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_58_i3, _57_v25)
					}
				})(_51_v25)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw7 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw7).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw7).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_56_v30 = _nw7
			_56_v30 = _56_v30
		} else {
			_52_v26 = p0
			var _59_v31 *C0
			_ = _59_v31
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__(_dafny.Companion_Sequence_.Concatenate(_18_v11, _18_v11))
			_59_v31 = _nw8
			var _60_v32 _dafny.CodePoint
			_ = _60_v32
			_60_v32 = _dafny.CodePoint('u')
			var _61_v33 _dafny.MultiSet
			_ = _61_v33
			_61_v33 = _dafny.MultiSetOf(_60_v32)
			_51_v25 = (func() _dafny.Int {
				if (_61_v33).Contains(_60_v32) {
					return (_61_v33).Multiplicity(_60_v32)
				}
				return (_59_v31).Fm1(p0, _52_v26, true, p1, globalState)
			})()
			_19_v12 = _19_v12
			var _62_v34 T0
			_ = _62_v34
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__(_dafny.SeqOf(_17_v10))
			_62_v34 = _nw9
			_62_v34 = _62_v34
		}
		var _63_v35 _dafny.Set
		_ = _63_v35
		_63_v35 = _dafny.SetOf(false)
		var _64_v36 D2
		_ = _64_v36
		_64_v36 = Companion_D2_.Create_DC7_(_16_v9, (_63_v35).Cardinality(), _51_v25, !(true))
		var _65_v37 _dafny.Map
		_ = _65_v37
		_65_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v36, _52_v26)
		_65_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_64_v36, _52_v26)
		var _66_v38 _dafny.CodePoint
		_ = _66_v38
		_66_v38 = _dafny.CodePoint('t')
		var _67_v39 D0
		_ = _67_v39
		_67_v39 = Companion_D0_.Create_DC1_(_52_v26, Companion_Default___.Fm5(_52_v26, globalState), _51_v25, false, _51_v25)
		var _68_v40 _dafny.Map
		_ = _68_v40
		_68_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_66_v38, (_67_v39).Dtor_cf3())
		_68_v40 = (_68_v40).Update(_dafny.CodePoint('h'), _51_v25)
		if _52_v26 {
			_51_v25 = (_51_v25).Times((_dafny.IntOfUint32((_16_v9).Cardinality())).Minus(_dafny.IntOfInt64(436)))
			_66_v38 = _66_v38
			var _69_v41 *C0
			_ = _69_v41
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__(_18_v11)
			_69_v41 = _nw10
			var _70_v42 _dafny.Set
			_ = _70_v42
			_70_v42 = _dafny.SetOf(_51_v25)
			var _71_v43 _dafny.Sequence
			_ = _71_v43
			_71_v43 = _dafny.SeqOf(_69_v41, _69_v41)
			var _72_v44 _dafny.Array
			_ = _72_v44
			var _nwElement0_0 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_51_v25, (_dafny.Zero).Minus(_51_v25)))
			_ = _nwElement0_0
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(15))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_0, 0)
			(_nw11).ArraySet1(((_70_v42).Cardinality()).Minus(_51_v25), 1)
			(_nw11).ArraySet1(_51_v25, 2)
			(_nw11).ArraySet1(_51_v25, 3)
			(_nw11).ArraySet1(_dafny.IntOfUint32((_16_v9).Cardinality()), 4)
			(_nw11).ArraySet1(_51_v25, 5)
			(_nw11).ArraySet1(_dafny.IntOfUint32((_71_v43).Cardinality()), 6)
			(_nw11).ArraySet1(_51_v25, 7)
			(_nw11).ArraySet1((_64_v36).Dtor_cf13(), 8)
			(_nw11).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_16_v9).Cardinality())), 9)
			(_nw11).ArraySet1(_51_v25, 10)
			(_nw11).ArraySet1((_64_v36).Dtor_cf13(), 11)
			(_nw11).ArraySet1(_51_v25, 12)
			(_nw11).ArraySet1(_51_v25, 13)
			(_nw11).ArraySet1(_51_v25, 14)
			_72_v44 = _nw11
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_72_v44), 0))
			_ = _index2
			(_72_v44).ArraySet1(_51_v25, (_index2).Int())
			var _73_v45 _dafny.Sequence
			_ = _73_v45
			_73_v45 = _dafny.SeqOf(_52_v26, p0)
			var _74_v46 D4
			_ = _74_v46
			_74_v46 = Companion_D4_.Create_DC14_(_73_v45, _dafny.CodePoint('l'), _51_v25, _51_v25)
			var _75_v47 _dafny.Set
			_ = _75_v47
			_75_v47 = _dafny.SetOf((_74_v46).Dtor_cf22(), _66_v38, _66_v38)
			var _76_v48 D3
			_ = _76_v48
			_76_v48 = Companion_D3_.Create_DC11_(_75_v47)
			var _pat_let_tv4 = _75_v47
			_ = _pat_let_tv4
			var _pat_let_tv5 = _75_v47
			_ = _pat_let_tv5
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_72_v44), 0))
			_ = _index3
			var _rhs3 bool = !(p0) || (true)
			_ = _rhs3
			var _rhs4 _dafny.Int = ((func(_pat_let0_0 D3) D3 {
				return func(_77_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let1_0 _dafny.Set) D3 {
						return func(_79_dt__update_hcf19_h0 _dafny.Set) D3 {
							return Companion_D3_.Create_DC11_(_79_dt__update_hcf19_h0)
						}(_pat_let1_0)
					}(func() _dafny.Set {
						var _coll4 = _dafny.NewBuilder()
						_ = _coll4
						for _iter4 := _dafny.Iterate((_pat_let_tv4).Elements()); ; {
							_compr_4, _ok4 := _iter4()
							if !_ok4 {
								break
							}
							var _78_v49 _dafny.CodePoint
							_78_v49 = interface{}(_compr_4).(_dafny.CodePoint)
							if (_pat_let_tv5).Contains(_78_v49) {
								_coll4.Add(_78_v49)
							}
						}
						return _coll4.ToSet()
					}())
				}(_pat_let0_0)
			}(_76_v48)).Dtor_cf19()).Cardinality()
			_ = _rhs4
			var _lhs0 _dafny.Array = _72_v44
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_72_v44), 0))
			_ = _lhs1
			_52_v26 = _rhs3
			(_lhs0).ArraySet1(_rhs4, (_lhs1).Int())
			_73_v45 = _73_v45
		} else {
			_51_v25 = _51_v25
			var _80_v50 T0
			_ = _80_v50
			var _nw12 *C0 = New_C0_()
			_ = _nw12
			_nw12.Ctor__(_18_v11)
			_80_v50 = _nw12
			var _81_v51 D2
			_ = _81_v51
			_81_v51 = Companion_D2_.Create_DC8_(p0, _80_v50)
			_52_v26 = ((func() D2 {
				if true {
					return _81_v51
				}
				return _81_v51
			})()).Dtor_cf15()
			var _82_v52 _dafny.MultiSet
			_ = _82_v52
			_82_v52 = _dafny.MultiSetOf(_dafny.CodePoint('v'), _66_v38)
			var _83_v53 _dafny.Map
			_ = _83_v53
			_83_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_67_v39).Dtor_cf2(), _16_v9)
			var _84_v54 _dafny.Sequence
			_ = _84_v54
			_84_v54 = _dafny.SeqOf((_82_v52).IsSubsetOf(_dafny.MultiSetOf(_66_v38, _66_v38)), !(false), (_51_v25).Cmp(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_83_v53).Contains(_51_v25) {
					return (_83_v53).Get(_51_v25).(_dafny.Sequence)
				}
				return _16_v9
			})()).Cardinality())) <= 0)
			_84_v54 = _dafny.Companion_Sequence_.Concatenate(_84_v54, _dafny.Companion_Sequence_.Update(_84_v54, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_16_v9).Cardinality()), _dafny.IntOfUint32((_84_v54).Cardinality()))).Uint32(), p1))
			_51_v25 = (_51_v25).Plus(_51_v25)
			_51_v25 = _51_v25
		}
	}
	var _85_v55 _dafny.Int
	_ = _85_v55
	_85_v55 = _dafny.IntOfInt64(115)
	var _86_v56 _dafny.Set
	_ = _86_v56
	_86_v56 = _dafny.SetOf(_85_v55, _85_v55)
	var _87_v57 _dafny.Map
	_ = _87_v57
	_87_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v56, _16_v9)
	_87_v57 = (_87_v57).Merge(_87_v57)
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _88_v0 _dafny.Int
	_ = _88_v0
	_88_v0 = _dafny.IntOfInt64(159)
	var _89_v1 _dafny.Sequence
	_ = _89_v1
	_89_v1 = _dafny.UnicodeSeqOfUtf8Bytes("vqomatru")
	var _90_v2 _dafny.Sequence
	_ = _90_v2
	_90_v2 = _dafny.SeqOf(_88_v0, (_dafny.Zero).Minus((_dafny.Zero).Minus(_88_v0)), _88_v0, _dafny.IntOfUint32((_89_v1).Cardinality()))
	var _91_globalState *GlobalState
	_ = _91_globalState
	var _nw13 *GlobalState = New_GlobalState_()
	_ = _nw13
	_nw13.Ctor__(_dafny.IntOfInt64(862), _dafny.IntOfInt64(771), _dafny.IntOfInt64(785), func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_90_v2).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _92_v3 _dafny.Int
			_92_v3 = interface{}(_compr_5).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_90_v2, _92_v3) {
				_coll5.Add((_92_v3).Plus(_dafny.IntOfInt64(477)))
			}
		}
		return _coll5.ToSet()
	}(), _dafny.IntOfInt64(-185), _dafny.CodePoint('u'), false, _dafny.IntOfInt64(385), true)
	_91_globalState = _nw13
	var _93_v4 D0
	_ = _93_v4
	_93_v4 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm0(_91_globalState), _89_v1)).Cardinality()))
	var _source1 D0 = _93_v4
	_ = _source1
	if _source1.Is_DC0() {
		var _94___mcc_h0 _dafny.Int = _source1.Get_().(D0_DC0).Cf0
		_ = _94___mcc_h0
		var _95_cf0 _dafny.Int = _94___mcc_h0
		_ = _95_cf0
		_88_v0 = _95_cf0
		var _96_v5 bool
		_ = _96_v5
		_96_v5 = false
		_96_v5 = _96_v5
		var _97_v6 _dafny.Set
		_ = _97_v6
		_97_v6 = _dafny.SetOf(_89_v1, _dafny.UnicodeSeqOfUtf8Bytes("erdcbq"), _dafny.Companion_Sequence_.Concatenate(_89_v1, _89_v1))
		var _98_v7 _dafny.Array
		_ = _98_v7
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw14
		_98_v7 = _nw14
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_98_v7), 0))
		_ = _index4
		(_98_v7).ArraySet1(_88_v0, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_98_v7), 0))
		_ = _index5
		var _rhs5 _dafny.Set = (_97_v6).Difference(_dafny.SetOf(_89_v1))
		_ = _rhs5
		var _rhs6 _dafny.Int = _88_v0
		_ = _rhs6
		var _lhs2 _dafny.Array = _98_v7
		_ = _lhs2
		var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_98_v7), 0))
		_ = _lhs3
		_97_v6 = _rhs5
		(_lhs2).ArraySet1(_rhs6, (_lhs3).Int())
		var _99_v8 _dafny.Sequence
		_ = _99_v8
		_99_v8 = _dafny.SeqOf(_96_v5, (_96_v5) && (_96_v5))
		_96_v5 = (_99_v8).Select((Companion_Default___.SafeIndex(_95_cf0, _dafny.IntOfUint32((_99_v8).Cardinality()))).Uint32()).(bool)
	} else if _source1.Is_DC1() {
		var _100___mcc_h1 bool = _source1.Get_().(D0_DC1).Cf1
		_ = _100___mcc_h1
		var _101___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC1).Cf2
		_ = _101___mcc_h2
		var _102___mcc_h3 _dafny.Int = _source1.Get_().(D0_DC1).Cf3
		_ = _102___mcc_h3
		var _103___mcc_h4 bool = _source1.Get_().(D0_DC1).Cf4
		_ = _103___mcc_h4
		var _104___mcc_h5 _dafny.Int = _source1.Get_().(D0_DC1).Cf5
		_ = _104___mcc_h5
		var _105_cf5 _dafny.Int = _104___mcc_h5
		_ = _105_cf5
		var _106_cf4 bool = _103___mcc_h4
		_ = _106_cf4
		var _107_cf3 _dafny.Int = _102___mcc_h3
		_ = _107_cf3
		var _108_cf2 _dafny.Int = _101___mcc_h2
		_ = _108_cf2
		var _109_cf1 bool = _100___mcc_h1
		_ = _109_cf1
		var _110_v9 _dafny.Set
		_ = _110_v9
		_110_v9 = _dafny.SetOf(false)
		var _111_v10 _dafny.Set
		_ = _111_v10
		_111_v10 = _dafny.SetOf(_110_v9, _110_v9)
		_106_cf4 = !((_111_v10).IsSubsetOf((_111_v10).Union(_111_v10)))
		Companion_Default___.M0(_109_cf1, _109_cf1, _91_globalState)
		var _112_v11 _dafny.Array
		_ = _112_v11
		var _nw15 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
		_ = _nw15
		_112_v11 = _nw15
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_112_v11), 0))
		_ = _index6
		(_112_v11).ArraySet1(_109_cf1, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_112_v11), 0))
		_ = _index7
		(_112_v11).ArraySet1((func() bool {
			if _109_cf1 {
				return (_88_v0).Cmp(_108_cf2) > 0
			}
			return !(_106_cf4) || (false)
		})(), (_index7).Int())
		var _113_v12 _dafny.Array
		_ = _113_v12
		var _nwElement0_1 _dafny.CodePoint = _dafny.CodePoint('o')
		_ = _nwElement0_1
		var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
		_ = _nw16
		(_nw16).ArraySet1CodePoint(_nwElement0_1, 0)
		_113_v12 = _nw16
		var _114_v13 _dafny.Sequence
		_ = _114_v13
		_114_v13 = _dafny.SeqOf(_113_v12)
		var _115_v14 *C0
		_ = _115_v14
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__(_dafny.Companion_Sequence_.Update(_114_v13, (Companion_Default___.SafeIndex(_105_cf5, _dafny.IntOfUint32((_114_v13).Cardinality()))).Uint32(), _113_v12))
		_115_v14 = _nw17
	} else {
		var _116___mcc_h6 D0 = _source1.Get_().(D0_DC2).Cf6
		_ = _116___mcc_h6
		var _117_cf6 D0 = _116___mcc_h6
		_ = _117_cf6
		var _118_v15 bool
		_ = _118_v15
		_118_v15 = true
		var _119_v16 _dafny.Map
		_ = _119_v16
		_119_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v15) == (_118_v15), _88_v0)
		_119_v16 = (_119_v16).Update(_118_v15, _88_v0)
		var _120_v17 _dafny.CodePoint
		_ = _120_v17
		_120_v17 = _dafny.CodePoint('d')
		_118_v15 = Companion_Default___.Fm2(!(_118_v15), _120_v17, _88_v0, (func() bool {
			if !(_118_v15) {
				return _118_v15
			}
			return _118_v15
		})(), _91_globalState)
		var _121_v18 _dafny.Map
		_ = _121_v18
		_121_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_88_v0), (_118_v15) || (_118_v15))
		_121_v18 = (_121_v18).Update(_dafny.IntOfInt64(-558), (_118_v15) || (false))
		var _122_v19 _dafny.Array
		_ = _122_v19
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
		_ = _nw18
		_122_v19 = _nw18
		var _123_v20 _dafny.Array
		_ = _123_v20
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(9))
		_ = _nw19
		_123_v20 = _nw19
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_122_v19), 0))
		_ = _index8
		(_122_v19).ArraySet1(_123_v20, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_122_v19), 0))
		_ = _index9
		(_122_v19).ArraySet1(_123_v20, (_index9).Int())
	}
	var _124_v21 _dafny.Array
	_ = _124_v21
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
	_ = _nw20
	_124_v21 = _nw20
	var _125_v22 _dafny.CodePoint
	_ = _125_v22
	_125_v22 = _dafny.CodePoint('h')
	var _126_v23 _dafny.Array
	_ = _126_v23
	var _nwElement0_2 _dafny.CodePoint = _125_v22
	_ = _nwElement0_2
	var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(12))
	_ = _nw21
	(_nw21).ArraySet1CodePoint(_nwElement0_2, 0)
	(_nw21).ArraySet1CodePoint(_125_v22, 1)
	(_nw21).ArraySet1CodePoint(_125_v22, 2)
	(_nw21).ArraySet1CodePoint(_125_v22, 3)
	(_nw21).ArraySet1CodePoint(_125_v22, 4)
	(_nw21).ArraySet1CodePoint(_125_v22, 5)
	(_nw21).ArraySet1CodePoint(_125_v22, 6)
	(_nw21).ArraySet1CodePoint(_dafny.CodePoint('q'), 7)
	(_nw21).ArraySet1CodePoint(_125_v22, 8)
	(_nw21).ArraySet1CodePoint(_125_v22, 9)
	(_nw21).ArraySet1CodePoint(_125_v22, 10)
	(_nw21).ArraySet1CodePoint(_125_v22, 11)
	_126_v23 = _nw21
	var _127_v24 _dafny.Sequence
	_ = _127_v24
	_127_v24 = _dafny.SeqOf(_126_v23)
	var _128_v25 *C0
	_ = _128_v25
	var _nw22 *C0 = New_C0_()
	_ = _nw22
	_nw22.Ctor__(_127_v24)
	_128_v25 = _nw22
	var _129_v26 _dafny.Set
	_ = _129_v26
	_129_v26 = _dafny.SetOf(_128_v25, _128_v25)
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_124_v21), 0))
	_ = _index10
	(_124_v21).ArraySet1(_129_v26, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_124_v21), 0))
	_ = _index11
	(_124_v21).ArraySet1(_129_v26, (_index11).Int())
	var _source2 D0 = _93_v4
	_ = _source2
	if _source2.Is_DC0() {
		var _130___mcc_h7 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
		_ = _130___mcc_h7
		var _131_cf0 _dafny.Int = _130___mcc_h7
		_ = _131_cf0
		_89_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rllcb"), _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("t"), (Companion_Default___.SafeIndex(_88_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()))).Uint32(), _125_v22))
		var _132_v27 T0
		_ = _132_v27
		var _nw23 *C0 = New_C0_()
		_ = _nw23
		_nw23.Ctor__(_dafny.SeqOf(_126_v23, _126_v23))
		_132_v27 = _nw23
		_132_v27 = _132_v27
		_88_v0 = _88_v0
		var _133_v28 _dafny.Array
		_ = _133_v28
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_2
		var _nw24 _dafny.Array
		_ = _nw24
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw24 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = func(_134_i0 _dafny.Int) bool {
				return (false) == (true)
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw24 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw24).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw24).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_133_v28 = _nw24
		var _135_v29 bool
		_ = _135_v29
		_135_v29 = true
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_133_v28), 0))
		_ = _index12
		(_133_v28).ArraySet1(_135_v29, (_index12).Int())
		var _136_v30 _dafny.Set
		_ = _136_v30
		_136_v30 = _dafny.SetOf(_135_v29)
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_133_v28), 0))
		_ = _index13
		(_133_v28).ArraySet1((func() bool {
			if _135_v29 {
				return !(_135_v29) || ((Companion_D0_.Create_DC1_(_135_v29, (_129_v26).Cardinality(), (_dafny.Zero).Minus(_88_v0), _135_v29, _dafny.IntOfInt64(556))).Dtor_cf4())
			}
			return ((_136_v30).Cardinality()).Cmp(_88_v0) <= 0
		})(), (_index13).Int())
	} else if _source2.Is_DC1() {
		var _137___mcc_h8 bool = _source2.Get_().(D0_DC1).Cf1
		_ = _137___mcc_h8
		var _138___mcc_h9 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
		_ = _138___mcc_h9
		var _139___mcc_h10 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
		_ = _139___mcc_h10
		var _140___mcc_h11 bool = _source2.Get_().(D0_DC1).Cf4
		_ = _140___mcc_h11
		var _141___mcc_h12 _dafny.Int = _source2.Get_().(D0_DC1).Cf5
		_ = _141___mcc_h12
		var _142_cf5 _dafny.Int = _141___mcc_h12
		_ = _142_cf5
		var _143_cf4 bool = _140___mcc_h11
		_ = _143_cf4
		var _144_cf3 _dafny.Int = _139___mcc_h10
		_ = _144_cf3
		var _145_cf2 _dafny.Int = _138___mcc_h9
		_ = _145_cf2
		var _146_cf1 bool = _137___mcc_h8
		_ = _146_cf1
		var _147_v31 _dafny.Array
		_ = _147_v31
		var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
		_ = _nw25
		_147_v31 = _nw25
		var _148_v32 _dafny.Array
		_ = _148_v32
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw26
		_148_v32 = _nw26
		var _149_v33 _dafny.Map
		_ = _149_v33
		_149_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_v31, _148_v32)
		_149_v33 = (_149_v33).Update(_147_v31, _148_v32)
		_128_v25 = _128_v25
		_89_v1 = _dafny.Companion_Sequence_.Concatenate(_89_v1, _89_v1)
		if _143_cf4 {
			var _150_v34 D0
			_ = _150_v34
			_150_v34 = Companion_D0_.Create_DC1_(false, _142_cf5, _88_v0, _146_cf1, _145_cf2)
			_144_cf3 = (_150_v34).Dtor_cf2()
			var _151_v35 _dafny.Sequence
			_ = _151_v35
			_151_v35 = _dafny.SeqOf(_148_v32, _148_v32)
			var _152_v36 _dafny.Sequence
			_ = _152_v36
			_152_v36 = _dafny.SeqOf((_151_v35).Select((Companion_Default___.SafeIndex(_144_cf3, _dafny.IntOfUint32((_151_v35).Cardinality()))).Uint32()).(_dafny.Array), _148_v32, _148_v32)
			var _153_v37 _dafny.Map
			_ = _153_v37
			_153_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_cf1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_151_v35, _151_v35)).Cardinality()))
			var _154_v38 _dafny.Sequence
			_ = _154_v38
			_154_v38 = _dafny.SeqOf(_146_cf1)
			var _155_v39 _dafny.Map
			_ = _155_v39
			_155_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v1, _154_v38)
			_153_v37 = (_153_v37).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v1, _154_v38)).Equals(_155_v39), _88_v0)
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_147_v31), 0))
			_ = _index14
			(_147_v31).ArraySet1((func() _dafny.Int {
				if _143_cf4 {
					return _145_cf2
				}
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(_88_v0))
			})(), (_index14).Int())
			var _156_v40 _dafny.Set
			_ = _156_v40
			_156_v40 = _dafny.SetOf((_128_v25).Fm1(Companion_Default___.Fm2(_146_cf1, _dafny.CodePoint('w'), _dafny.IntOfUint32((_90_v2).Cardinality()), !(_143_cf4), _91_globalState), _146_cf1, _146_cf1, _146_cf1, _91_globalState), _144_cf3, _145_cf2)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_147_v31), 0))
			_ = _index15
			(_147_v31).ArraySet1((_88_v0).Plus((_156_v40).Cardinality()), (_index15).Int())
			_146_cf1 = (_150_v34).Dtor_cf1()
			_146_cf1 = !(!(Companion_Default___.Fm2(_143_cf4, (func() _dafny.CodePoint {
				if _146_cf1 {
					return _125_v22
				}
				return Companion_Default___.Fm3((_150_v34).Dtor_cf3(), _91_globalState)
			})(), (_144_cf3).Minus((_dafny.Zero).Minus(_88_v0)), _146_cf1, _91_globalState)))
		} else {
			Companion_Default___.M0((_143_cf4) == (false), _146_cf1, _91_globalState)
			_88_v0 = ((_90_v2).Select((Companion_Default___.SafeIndex(_142_cf5, _dafny.IntOfUint32((_90_v2).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_dafny.IntOfInt64(-475))
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_148_v32), 0))
			_ = _index16
			(_148_v32).ArraySet1(_146_cf1, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_148_v32), 0))
			_ = _index17
			(_148_v32).ArraySet1((_129_v26).Contains(_128_v25), (_index17).Int())
			var _157_v41 _dafny.Array
			_ = _157_v41
			var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
			_ = _nw27
			_157_v41 = _nw27
			_157_v41 = _157_v41
			var _158_v42 _dafny.Array
			_ = _158_v42
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw28
			_158_v42 = _nw28
			_158_v42 = _158_v42
		}
	} else {
		var _159___mcc_h13 D0 = _source2.Get_().(D0_DC2).Cf6
		_ = _159___mcc_h13
		var _160_cf6 D0 = _159___mcc_h13
		_ = _160_cf6
		var _161_v43 bool
		_ = _161_v43
		_161_v43 = false
		if !(_161_v43) {
			_88_v0 = _88_v0
			var _162_v44 _dafny.Sequence
			_ = _162_v44
			_162_v44 = _dafny.SeqOf(_161_v43, _161_v43, _161_v43, _161_v43, _161_v43)
			var _163_v45 _dafny.MultiSet
			_ = _163_v45
			_163_v45 = _dafny.MultiSetOf(Companion_Default___.Fm2(_161_v43, _125_v22, _88_v0, !(_161_v43), _91_globalState))
			var _164_v46 _dafny.Map
			_ = _164_v46
			_164_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v0, _88_v0)
			var _165_v47 D0
			_ = _165_v47
			_165_v47 = Companion_D0_.Create_DC1_(_161_v43, (_163_v45).Cardinality(), (_164_v46).Cardinality(), _161_v43, _88_v0)
			var _166_v48 T0
			_ = _166_v48
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__(_127_v24)
			_166_v48 = _nw29
			var _167_v49 _dafny.Map
			_ = _167_v49
			_167_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v48, _161_v43)
			var _168_v50 _dafny.MultiSet
			_ = _168_v50
			_168_v50 = _dafny.MultiSetOf(_88_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(529))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}((func(_169_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_170_i1 _dafny.Int) _dafny.CodePoint {
					return _169_v22
				}
			})(_125_v22)))).Cardinality()))
			var _171_v51 D0
			_ = _171_v51
			_171_v51 = Companion_D0_.Create_DC1_((_165_v47).Dtor_cf4(), _88_v0, ((_167_v49).Update(_166_v48, _161_v43)).Cardinality(), _161_v43, (func() _dafny.Int {
				if (_168_v50).Contains(_88_v0) {
					return (_168_v50).Multiplicity(_88_v0)
				}
				return (_90_v2).Select((Companion_Default___.SafeIndex(_88_v0, _dafny.IntOfUint32((_90_v2).Cardinality()))).Uint32()).(_dafny.Int)
			})())
			var _172_v52 _dafny.MultiSet
			_ = _172_v52
			_172_v52 = _dafny.MultiSetOf(_128_v25)
			var _173_v53 _dafny.Array
			_ = _173_v53
			var _nwElement0_3 _dafny.Int = _88_v0
			_ = _nwElement0_3
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(12))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_3, 0)
			(_nw30).ArraySet1(_dafny.IntOfUint32((_162_v44).Cardinality()), 1)
			(_nw30).ArraySet1(_88_v0, 2)
			(_nw30).ArraySet1(_dafny.IntOfInt64(916), 3)
			(_nw30).ArraySet1((_165_v47).Dtor_cf5(), 4)
			(_nw30).ArraySet1(_88_v0, 5)
			(_nw30).ArraySet1(Companion_Default___.SafeDivisionInt((_172_v52).Cardinality(), (_dafny.Zero).Minus((_166_v48).Fm1(_161_v43, _161_v43, _161_v43, false, _91_globalState))), 6)
			(_nw30).ArraySet1(_dafny.IntOfUint32((_89_v1).Cardinality()), 7)
			(_nw30).ArraySet1(_88_v0, 8)
			(_nw30).ArraySet1(_dafny.IntOfInt64(810), 9)
			(_nw30).ArraySet1(_88_v0, 10)
			(_nw30).ArraySet1(_88_v0, 11)
			_173_v53 = _nw30
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_173_v53), 0))
			_ = _index18
			(_173_v53).ArraySet1(_88_v0, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_173_v53), 0))
			_ = _index19
			(_173_v53).ArraySet1((_166_v48).Fm1((_88_v0).Cmp(_88_v0) < 0, _161_v43, _161_v43, (func() bool {
				if Companion_Default___.Fm2(Companion_Default___.Fm2(_161_v43, _125_v22, _88_v0, _161_v43, _91_globalState), _125_v22, _88_v0, _161_v43, _91_globalState) {
					return Companion_Default___.Fm2(_161_v43, _125_v22, _dafny.IntOfUint32((_89_v1).Cardinality()), _161_v43, _91_globalState)
				}
				return false
			})(), _91_globalState), (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_173_v53), 0))
			_ = _index20
			(_173_v53).ArraySet1((_128_v25).Fm1(_161_v43, Companion_Default___.Fm2(_161_v43, _125_v22, _dafny.IntOfInt64(-962), _161_v43, _91_globalState), _161_v43, !((_168_v50).IsSubsetOf(_168_v50)), _91_globalState), (_index20).Int())
			_88_v0 = _dafny.IntOfUint32((_89_v1).Cardinality())
			_161_v43 = !(_161_v43)
		} else {
			var _174_v54 T0
			_ = _174_v54
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(_127_v24)
			_174_v54 = _nw31
			var _175_v55 _dafny.Sequence
			_ = _175_v55
			_175_v55 = _dafny.SeqOf(_174_v54, _174_v54, _174_v54, _174_v54)
			var _176_v56 D1
			_ = _176_v56
			_176_v56 = Companion_D1_.Create_DC3_((_175_v55).Select((Companion_Default___.SafeIndex(_88_v0, _dafny.IntOfUint32((_175_v55).Cardinality()))).Uint32()).(T0))
			var _177_v57 _dafny.Array
			_ = _177_v57
			var _nwElement0_4 T0 = _174_v54
			_ = _nwElement0_4
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(21))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_4, 0)
			(_nw32).ArraySet1((Companion_D1_.Create_DC3_(_174_v54)).Dtor_cf7(), 1)
			(_nw32).ArraySet1(_174_v54, 2)
			(_nw32).ArraySet1((_176_v56).Dtor_cf7(), 3)
			(_nw32).ArraySet1(_174_v54, 4)
			(_nw32).ArraySet1(_174_v54, 5)
			(_nw32).ArraySet1(_174_v54, 6)
			(_nw32).ArraySet1(_174_v54, 7)
			(_nw32).ArraySet1(_174_v54, 8)
			(_nw32).ArraySet1(_174_v54, 9)
			(_nw32).ArraySet1(_174_v54, 10)
			(_nw32).ArraySet1(_174_v54, 11)
			(_nw32).ArraySet1(_174_v54, 12)
			(_nw32).ArraySet1(_174_v54, 13)
			(_nw32).ArraySet1(_174_v54, 14)
			(_nw32).ArraySet1(_174_v54, 15)
			(_nw32).ArraySet1(_174_v54, 16)
			(_nw32).ArraySet1(_174_v54, 17)
			(_nw32).ArraySet1(_174_v54, 18)
			(_nw32).ArraySet1(_174_v54, 19)
			(_nw32).ArraySet1(_174_v54, 20)
			_177_v57 = _nw32
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_177_v57), 0))
			_ = _index21
			(_177_v57).ArraySet1(_174_v54, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_177_v57), 0))
			_ = _index22
			(_177_v57).ArraySet1((func() T0 {
				if _161_v43 {
					return _174_v54
				}
				return (func() T0 {
					if _161_v43 {
						return _174_v54
					}
					return _174_v54
				})()
			})(), (_index22).Int())
			var _178_v58 _dafny.Map
			_ = _178_v58
			_178_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_161_v43, false)
			_178_v58 = (_178_v58).Update(_161_v43, _161_v43)
			var _179_v59 _dafny.Array
			_ = _179_v59
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw33
			_179_v59 = _nw33
			_179_v59 = _179_v59
			var _180_v60 _dafny.Array
			_ = _180_v60
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
			_ = _nw34
			_180_v60 = _nw34
			var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
			_ = _nw35
			_180_v60 = _nw35
			var _181_v61 _dafny.Sequence
			_ = _181_v61
			_181_v61 = _dafny.SeqOf(_161_v43, _161_v43, _161_v43)
			_181_v61 = _181_v61
		}
		var _182_v62 _dafny.Array
		_ = _182_v62
		var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
		_ = _nw36
		_182_v62 = _nw36
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_182_v62), 0))
		_ = _index23
		(_182_v62).ArraySet1(_88_v0, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_182_v62), 0))
		_ = _index24
		(_182_v62).ArraySet1(_88_v0, (_index24).Int())
		_161_v43 = _161_v43
		_88_v0 = _88_v0
	}
	var _183_v63 bool
	_ = _183_v63
	_183_v63 = true
	if _183_v63 {
		_183_v63 = _183_v63
		var _184_v64 _dafny.Map
		_ = _184_v64
		_184_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_v0, false)
		_88_v0 = (((_184_v64).Cardinality()).Times(_88_v0)).Plus(_88_v0)
		_89_v1 = Companion_Default___.Fm0(_91_globalState)
		if !(_183_v63) {
			_183_v63 = _183_v63
			var _185_v65 _dafny.Array
			_ = _185_v65
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw37
			_185_v65 = _nw37
			var _186_v66 _dafny.Sequence
			_ = _186_v66
			_186_v66 = _dafny.SeqOf(_183_v63)
			var _187_v67 _dafny.Set
			_ = _187_v67
			_187_v67 = _dafny.SetOf(_dafny.IntOfUint32((_186_v66).Cardinality()), _88_v0)
			var _188_v68 D0
			_ = _188_v68
			_188_v68 = Companion_D0_.Create_DC1_(_183_v63, _88_v0, _88_v0, _183_v63, (_187_v67).Cardinality())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_185_v65), 0))
			_ = _index25
			(_185_v65).ArraySet1((_188_v68).Dtor_cf1(), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_185_v65), 0))
			_ = _index26
			(_185_v65).ArraySet1(true, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_185_v65), 0))
			_ = _index27
			(_185_v65).ArraySet1(_183_v63, (_index27).Int())
			var _189_v69 _dafny.Map
			_ = _189_v69
			_189_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_185_v65), 0))).Int()).(bool), _186_v66)
			_189_v69 = _189_v69
			var _190_v70 T0
			_ = _190_v70
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__(_127_v24)
			_190_v70 = _nw38
			var _191_v71 _dafny.Map
			_ = _191_v71
			_191_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_190_v70, _88_v0)
			_191_v71 = (_191_v71).Update(_190_v70, _88_v0)
		} else {
			var _192_v72 _dafny.Map
			_ = _192_v72
			_192_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v63, _dafny.CodePoint('m'))
			_192_v72 = _192_v72
			var _193_v73 _dafny.MultiSet
			_ = _193_v73
			_193_v73 = _dafny.MultiSetOf(_183_v63)
			var _194_v74 _dafny.Sequence
			_ = _194_v74
			_194_v74 = _dafny.SeqOf(_183_v63)
			var _195_v75 _dafny.Array
			_ = _195_v75
			var _nwElement0_5 _dafny.MultiSet = _193_v73
			_ = _nwElement0_5
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(26))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_5, 0)
			(_nw39).ArraySet1(Companion_Default___.Fm4(false, _183_v63, _91_globalState), 1)
			(_nw39).ArraySet1(_193_v73, 2)
			(_nw39).ArraySet1(_dafny.MultiSetOf(false, _183_v63, _183_v63, _183_v63, _183_v63), 3)
			(_nw39).ArraySet1(_193_v73, 4)
			(_nw39).ArraySet1(_193_v73, 5)
			(_nw39).ArraySet1(_193_v73, 6)
			(_nw39).ArraySet1(_dafny.MultiSetOf(_183_v63), 7)
			(_nw39).ArraySet1(_193_v73, 8)
			(_nw39).ArraySet1(_193_v73, 9)
			(_nw39).ArraySet1(Companion_Default___.Fm4(_183_v63, false, _91_globalState), 10)
			(_nw39).ArraySet1((_193_v73).Update(_183_v63, Companion_Default___.Abs(_88_v0)), 11)
			(_nw39).ArraySet1(_193_v73, 12)
			(_nw39).ArraySet1(_193_v73, 13)
			(_nw39).ArraySet1(_193_v73, 14)
			(_nw39).ArraySet1(_193_v73, 15)
			(_nw39).ArraySet1(_193_v73, 16)
			(_nw39).ArraySet1(_193_v73, 17)
			(_nw39).ArraySet1(_dafny.MultiSetOf(_183_v63), 18)
			(_nw39).ArraySet1(_193_v73, 19)
			(_nw39).ArraySet1(_dafny.MultiSetOf(_183_v63), 20)
			(_nw39).ArraySet1(_193_v73, 21)
			(_nw39).ArraySet1(_193_v73, 22)
			(_nw39).ArraySet1(Companion_Default___.Fm4(Companion_Default___.Fm2(_183_v63, (_89_v1).Select((Companion_Default___.SafeIndex((_184_v64).Cardinality(), _dafny.IntOfUint32((_89_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), _88_v0, _183_v63, _91_globalState), (_194_v74).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(96), _dafny.IntOfUint32((_194_v74).Cardinality()))).Uint32()).(bool), _91_globalState), 23)
			(_nw39).ArraySet1(_193_v73, 24)
			(_nw39).ArraySet1(_193_v73, 25)
			_195_v75 = _nw39
			var _196_v76 _dafny.Map
			_ = _196_v76
			_196_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v75, !(_183_v63) || (_183_v63))
			_183_v63 = (func() bool {
				if (_196_v76).Contains(_195_v75) {
					return (_196_v76).Get(_195_v75).(bool)
				}
				return _183_v63
			})()
			_88_v0 = _88_v0
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_124_v21), 0))
			_ = _index28
			(_124_v21).ArraySet1(((_124_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_124_v21), 0))).Int()).(_dafny.Set)).Difference(_129_v26), (_index28).Int())
			var _197_v77 _dafny.Array
			_ = _197_v77
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_3
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Sequence = (func(_198_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_199_i2 _dafny.Int) _dafny.Sequence {
						return _198_v1
					}
				})(_89_v1)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw40 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw40).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw40).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_197_v77 = _nw40
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_197_v77), 0))
			_ = _index29
			(_197_v77).ArraySet1(_89_v1, (_index29).Int())
			var _200_v78 _dafny.Array
			_ = _200_v78
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
			_ = _nw41
			_200_v78 = _nw41
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_200_v78), 0))
			_ = _index30
			(_200_v78).ArraySet1(!(_183_v63), (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_197_v77), 0))
			_ = _index31
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_200_v78), 0))
			_ = _index32
			var _rhs7 _dafny.Sequence = _89_v1
			_ = _rhs7
			var _rhs8 bool = (!(_183_v63)) && (_183_v63)
			_ = _rhs8
			var _lhs4 _dafny.Array = _197_v77
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(809), _dafny.ArrayLen((_197_v77), 0))
			_ = _lhs5
			var _lhs6 _dafny.Array = _200_v78
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(965), _dafny.ArrayLen((_200_v78), 0))
			_ = _lhs7
			(_lhs4).ArraySet1(_rhs7, (_lhs5).Int())
			(_lhs6).ArraySet1(_rhs8, (_lhs7).Int())
		}
		var _201_v79 _dafny.Sequence
		_ = _201_v79
		_201_v79 = _dafny.SeqOf(_183_v63)
		_183_v63 = !(Companion_Default___.Fm2(_183_v63, _125_v22, _dafny.IntOfUint32((_90_v2).Cardinality()), Companion_Default___.Fm2(_183_v63, _125_v22, _dafny.IntOfUint32((_201_v79).Cardinality()), _183_v63, _91_globalState), _91_globalState)) || (!(!(_183_v63)))
	} else {
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_126_v23), 0))
		_ = _index33
		(_126_v23).ArraySet1CodePoint((_89_v1).Select((Companion_Default___.SafeIndex(_88_v0, _dafny.IntOfUint32((_89_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_126_v23), 0))
		_ = _index34
		var _rhs9 bool = false
		_ = _rhs9
		var _rhs10 _dafny.Int = _88_v0
		_ = _rhs10
		var _rhs11 _dafny.CodePoint = _125_v22
		_ = _rhs11
		var _rhs12 _dafny.Int = (func() _dafny.Int {
			if _183_v63 {
				return _88_v0
			}
			return _88_v0
		})()
		_ = _rhs12
		var _rhs13 bool = true
		_ = _rhs13
		var _lhs8 _dafny.Array = _126_v23
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_126_v23), 0))
		_ = _lhs9
		_183_v63 = _rhs9
		_88_v0 = _rhs10
		(_lhs8).ArraySet1CodePoint(_rhs11, (_lhs9).Int())
		_88_v0 = _rhs12
		_183_v63 = _rhs13
		var _202_v80 _dafny.Array
		_ = _202_v80
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_4
		var _nw42 _dafny.Array
		_ = _nw42
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw42 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Sequence = (func(_203_v2 _dafny.Sequence, _204_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_205_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_203_v2, _dafny.SeqOf(_204_v0, _204_v0, _204_v0))
				}
			})(_90_v2, _88_v0)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw42 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw42).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw42).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_202_v80 = _nw42
		_202_v80 = _202_v80
		var _206_v81 _dafny.Map
		_ = _206_v81
		_206_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v2, _89_v1)
		_206_v81 = _206_v81
		_183_v63 = _183_v63
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_126_v23), 0))
		_ = _index35
		(_126_v23).ArraySet1CodePoint((_126_v23).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_126_v23), 0))).Int()), (_index35).Int())
	}
	Companion_Default___.M0(_183_v63, (_183_v63) == (_183_v63), _91_globalState)
	_183_v63 = _183_v63
	_183_v63 = _183_v63
	var _207_v82 _dafny.Array
	_ = _207_v82
	var _len0_5 _dafny.Int = _dafny.One
	_ = _len0_5
	var _nw43 _dafny.Array
	_ = _nw43
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw43 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Sequence = (func(_208_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_209_i4 _dafny.Int) _dafny.Sequence {
				return _208_v2
			}
		})(_90_v2)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw43 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw43).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw43).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_207_v82 = _nw43
	var _210_v83 _dafny.MultiSet
	_ = _210_v83
	_210_v83 = _dafny.MultiSetOf(_183_v63, _183_v63)
	var _211_v84 _dafny.Map
	_ = _211_v84
	_211_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_207_v82, _210_v83)
	var _212_v85 _dafny.Sequence
	_ = _212_v85
	_212_v85 = _dafny.SeqOf((func() _dafny.MultiSet {
		if (_211_v84).Contains(_207_v82) {
			return (_211_v84).Get(_207_v82).(_dafny.MultiSet)
		}
		return _210_v83
	})(), _210_v83, _210_v83)
	_212_v85 = _212_v85
	var _213_v86 *C0
	_ = _213_v86
	var _nw44 *C0 = New_C0_()
	_ = _nw44
	_nw44.Ctor__(_127_v24)
	_213_v86 = _nw44
	Companion_Default___.M0(_183_v63, _183_v63, _91_globalState)
	var _214_v87 _dafny.Array
	_ = _214_v87
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_6
	var _nw45 _dafny.Array
	_ = _nw45
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw45 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) _dafny.Int = (func(_215_v2 _dafny.Sequence, _216_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_217_i5 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_217_i5, (_215_v2).Select((Companion_Default___.SafeIndex(_216_v0, _dafny.IntOfUint32((_215_v2).Cardinality()))).Uint32()).(_dafny.Int))
			}
		})(_90_v2, _88_v0)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw45 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw45).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw45).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_214_v87 = _nw45
	var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
	_ = _index36
	(_214_v87).ArraySet1(_88_v0, (_index36).Int())
	var _218_v88 _dafny.Map
	_ = _218_v88
	_218_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v63, _dafny.IntOfInt64(211))
	var _219_v89 _dafny.Set
	_ = _219_v89
	_219_v89 = _dafny.SetOf(_dafny.IntOfUint32((_90_v2).Cardinality()))
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
	_ = _index37
	(_214_v87).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
		if (_218_v88).Contains(_183_v63) {
			return (_218_v88).Get(_183_v63).(_dafny.Int)
		}
		return _88_v0
	})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v89).Cardinality(), (_128_v25).Fm1(false, _183_v63, !(_183_v63), !(_183_v63), _91_globalState))).Cardinality())), (_index37).Int())
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
	_ = _index38
	(_214_v87).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_88_v0, (_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int))), (_index38).Int())
	var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
	_ = _index39
	(_214_v87).ArraySet1((_88_v0).Minus((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int)), (_index39).Int())
	var _220_v90 _dafny.Array
	_ = _220_v90
	var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
	_ = _nw46
	_220_v90 = _nw46
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))
	_ = _index40
	(_220_v90).ArraySet1(_89_v1, (_index40).Int())
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))
	_ = _index41
	(_220_v90).ArraySet1(_89_v1, (_index41).Int())
	var _221_v91 _dafny.Map
	_ = _221_v91
	_221_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int), (_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int))
	_221_v91 = (func() _dafny.Map {
		if _183_v63 {
			return func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_90_v2).Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _222_v92 _dafny.Int
					_222_v92 = interface{}(_compr_6).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_90_v2, _222_v92) {
						_coll6.Add((_222_v92).Times(_88_v0), (_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int))
					}
				}
				return _coll6.ToMap()
			}()
		}
		return _221_v91
	})()
	var _223_v93 _dafny.Set
	_ = _223_v93
	_223_v93 = _dafny.SetOf(_210_v83, (_210_v83).Update(_183_v63, Companion_Default___.Abs((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int))), _dafny.MultiSetOf(_183_v63), _210_v83)
	if (_dafny.SetOf(_210_v83, _210_v83)).IsDisjointFrom(_223_v93) {
		var _224_v94 _dafny.Sequence
		_ = _224_v94
		_224_v94 = _dafny.SeqOf(_89_v1, _dafny.UnicodeSeqOfUtf8Bytes("rqfejscgi"), _89_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-358))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_225_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_226_i6 _dafny.Int) _dafny.CodePoint {
				return _225_v22
			}
		})(_125_v22))))
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))
		_ = _index42
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
		_ = _index43
		var _rhs14 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_224_v94).Select((Companion_Default___.SafeIndex((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_224_v94).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.Companion_Sequence_.Concatenate(_89_v1, (_220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))).Int()).(_dafny.Sequence)))
		_ = _rhs14
		var _rhs15 _dafny.Int = ((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int)).Minus((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int))
		_ = _rhs15
		var _lhs10 _dafny.Array = _220_v90
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))
		_ = _lhs11
		var _lhs12 _dafny.Array = _214_v87
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
		_ = _lhs13
		(_lhs10).ArraySet1(_rhs14, (_lhs11).Int())
		(_lhs12).ArraySet1(_rhs15, (_lhs13).Int())
		var _227_v95 _dafny.MultiSet
		_ = _227_v95
		_227_v95 = _dafny.MultiSetOf(_88_v0)
		var _228_v97 _dafny.Map
		_ = _228_v97
		_228_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v95, _219_v89)
		Companion_Default___.M0((_dafny.MultiSetFromSeq(_dafny.SeqOf((_213_v86).Fm1(_183_v63, _183_v63, _183_v63, _183_v63, _91_globalState), _88_v0))).IsSubsetOf(_227_v95), (func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-62), _dafny.IntOfInt64(728))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _229_v96 _dafny.Int
				_229_v96 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(-62)).Cmp(_229_v96) <= 0) && ((_229_v96).Cmp(_dafny.IntOfInt64(728)) < 0) {
					_coll7.Add((_229_v96).Minus(_88_v0))
				}
			}
			return _coll7.ToSet()
		}()).IsDisjointFrom((func() _dafny.Set {
			if (_228_v97).Contains(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wbacwiua")).Cardinality()))) {
				return (_228_v97).Get(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wbacwiua")).Cardinality()))).(_dafny.Set)
			}
			return _219_v89
		})()), _91_globalState)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))
		_ = _index44
		(_214_v87).ArraySet1((_dafny.Zero).Minus(_88_v0), (_index44).Int())
		var _230_v98 D2
		_ = _230_v98
		_230_v98 = Companion_D2_.Create_DC7_((_220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))).Int()).(_dafny.Sequence), _dafny.IntOfUint32((_89_v1).Cardinality()), (_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int), _183_v63)
		var _231_v99 _dafny.MultiSet
		_ = _231_v99
		_231_v99 = _dafny.MultiSetOf((_230_v98).Dtor_cf11(), _89_v1, (_220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))).Int()).(_dafny.Sequence))
		var _232_v100 _dafny.Map
		_ = _232_v100
		_232_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v99, _183_v63)
		_232_v100 = (_232_v100).Update((_dafny.MultiSetOf(_89_v1)).Difference(_231_v99), false)
		var _233_v101 _dafny.Map
		_ = _233_v101
		_233_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int), _183_v63)
		_233_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_213_v86).Fm1(_183_v63, _183_v63, _183_v63, _183_v63, _91_globalState), _183_v63)
	} else {
		_183_v63 = _183_v63
		var _234_v102 T0
		_ = _234_v102
		var _nw47 *C0 = New_C0_()
		_ = _nw47
		_nw47.Ctor__(_dafny.Companion_Sequence_.Concatenate(_127_v24, _127_v24))
		_234_v102 = _nw47
		var _235_v103 _dafny.Sequence
		_ = _235_v103
		_235_v103 = _dafny.SeqOf(_183_v63, _183_v63, false, _183_v63)
		var _236_v104 D2
		_ = _236_v104
		_236_v104 = Companion_D2_.Create_DC7_(_dafny.Companion_Sequence_.Update((_220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_88_v0, _dafny.IntOfUint32(((_220_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_220_v90), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), _125_v22), _88_v0, (_dafny.Zero).Minus((_dafny.SetOf(_183_v63)).Cardinality()), _183_v63)
		var _237_v105 _dafny.Map
		_ = _237_v105
		_237_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_89_v1).Cardinality()), _183_v63)
		var _238_v106 D1
		_ = _238_v106
		_238_v106 = Companion_D1_.Create_DC4_(_214_v87)
		var _239_v107 _dafny.Sequence
		_ = _239_v107
		_239_v107 = _dafny.SeqOf(_214_v87, (_238_v106).Dtor_cf8())
		var _rhs16 _dafny.Int = (_90_v2).Select((Companion_Default___.SafeIndex((_234_v102).Fm1(_183_v63, (_235_v103).Select((Companion_Default___.SafeIndex((_236_v104).Dtor_cf13(), _dafny.IntOfUint32((_235_v103).Cardinality()))).Uint32()).(bool), _183_v63, !(!(_183_v63)), _91_globalState), _dafny.IntOfUint32((_90_v2).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs16
		var _rhs17 T0 = _234_v102
		_ = _rhs17
		var _rhs18 bool = (func() bool {
			if !(Companion_Default___.Fm2(!(_183_v63), Companion_Default___.Fm3((_237_v105).Cardinality(), _91_globalState), (_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int), _183_v63, _91_globalState)) {
				return !(((_214_v87).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_214_v87), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_239_v107).Cardinality())) > 0)
			}
			return _183_v63
		})()
		_ = _rhs18
		var _rhs19 T0 = _234_v102
		_ = _rhs19
		_88_v0 = _rhs16
		_234_v102 = _rhs17
		_183_v63 = _rhs18
		_234_v102 = _rhs19
		_125_v22 = (func() _dafny.CodePoint {
			if (_183_v63) == (_183_v63) {
				return _125_v22
			}
			return _125_v22
		})()
		_128_v25 = _128_v25
		_221_v91 = (_221_v91).Update(_88_v0, (_234_v102).Fm1(_183_v63, _183_v63, true, _183_v63, _91_globalState))
	}
	_dafny.Print(_88_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_89_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_90_v2, _dafny.SeqOf(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_91_globalState).F3()).Equals(_dafny.SetOf(_dafny.IntOfInt64(636), _dafny.IntOfInt64(485))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_93_v4).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_124_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_125_v22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_126_v23).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_127_v24).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_129_v26).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_183_v63)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_207_v82).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(159), _dafny.IntOfInt64(8))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_210_v83).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_211_v84).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_212_v85, _dafny.SeqOf(_dafny.MultiSetOf(true, true), _dafny.MultiSetOf(true, true), _dafny.MultiSetOf(true, true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_214_v87).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v88).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(211))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v89).Equals(_dafny.SetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v90).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_221_v91).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(205110), _dafny.IntOfInt64(2580)).UpdateUnsafe(_dafny.IntOfInt64(10320), _dafny.IntOfInt64(2580)).UpdateUnsafe(_dafny.IntOfInt64(159), _dafny.IntOfInt64(1105))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v93).Equals(_dafny.SetOf(_dafny.MultiSetOf(true, true), _dafny.MultiSetOf(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true), _dafny.MultiSetOf(true))))
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

type D0_DC1 struct {
	Cf1 bool
	Cf2 _dafny.Int
	Cf3 _dafny.Int
	Cf4 bool
	Cf5 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 bool, Cf5 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf6 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf6 D0) D0 {
	return D0{D0_DC2{Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf5
}

func (_this D0) Dtor_cf6() D0 {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf6) + ")"
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
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf6.Equals(data2.Cf6)
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
	Cf8 _dafny.Array
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.Array) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC3 struct {
	Cf7 T0
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 T0) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC5 struct {
	Cf9 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 D1) D1 {
	return D1{D1_DC5{Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D1) Dtor_cf8() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf7() T0 {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf9() D1 {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf8 == data2.Cf8
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && _dafny.AreEqual(data1.Cf7, data2.Cf7)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf11 _dafny.Sequence
	Cf12 _dafny.Int
	Cf13 _dafny.Int
	Cf14 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 _dafny.Sequence, Cf12 _dafny.Int, Cf13 _dafny.Int, Cf14 bool) D2 {
	return D2{D2_DC7{Cf11, Cf12, Cf13, Cf14}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf15 bool
	Cf16 T0
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf15 bool, Cf16 T0) D2 {
	return D2{D2_DC8{Cf15, Cf16}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf17 bool
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf17 bool) D2 {
	return D2{D2_DC9{Cf17}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC6 struct {
	Cf10 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC10 struct {
	Cf18 D2
}

func (D2_DC10) isD2() {}

func (CompanionStruct_D2_) Create_DC10_(Cf18 D2) D2 {
	return D2{D2_DC10{Cf18}}
}

func (_this D2) Is_DC10() bool {
	_, ok := _this.Get_().(D2_DC10)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero, false)
}

func (_this D2) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() bool {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf15() bool {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf16() T0 {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf17() bool {
	return _this.Get_().(D2_DC9).Cf17
}

func (_this D2) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf18() D2 {
	return _this.Get_().(D2_DC10).Cf18
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + data.Cf11.VerbatimString(true) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + data.Cf10.VerbatimString(true) + ")"
		}
	case D2_DC10:
		{
			return "D2.DC10" + "(" + _dafny.String(data.Cf18) + ")"
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
			return ok && data1.Cf11.Equals(data2.Cf11) && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14 == data2.Cf14
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf15 == data2.Cf15 && _dafny.AreEqual(data1.Cf16, data2.Cf16)
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf17 == data2.Cf17
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	case D2_DC10:
		{
			data2, ok := other.Get_().(D2_DC10)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D3_DC12 struct {
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_() D3 {
	return D3{D3_DC12{}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC11 struct {
	Cf19 _dafny.Set
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf19 _dafny.Set) D3 {
	return D3{D3_DC11{Cf19}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC12_()
}

func (_this D3) Dtor_cf19() _dafny.Set {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC12:
		{
			return "D3.DC12"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC12:
		{
			_, ok := other.Get_().(D3_DC12)
			return ok
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf19.Equals(data2.Cf19)
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

type D4_DC14 struct {
	Cf21 _dafny.Sequence
	Cf22 _dafny.CodePoint
	Cf23 _dafny.Int
	Cf24 _dafny.Int
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf21 _dafny.Sequence, Cf22 _dafny.CodePoint, Cf23 _dafny.Int, Cf24 _dafny.Int) D4 {
	return D4{D4_DC14{Cf21, Cf22, Cf23, Cf24}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf20 _dafny.CodePoint
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf20 _dafny.CodePoint) D4 {
	return D4{D4_DC13{Cf20}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC15 struct {
	Cf25 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf25 D4) D4 {
	return D4{D4_DC15{Cf25}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(_dafny.EmptySeq, _dafny.CodePoint('D'), _dafny.Zero, _dafny.Zero)
}

func (_this D4) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() _dafny.CodePoint {
	return _this.Get_().(D4_DC14).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf24
}

func (_this D4) Dtor_cf20() _dafny.CodePoint {
	return _this.Get_().(D4_DC13).Cf20
}

func (_this D4) Dtor_cf25() D4 {
	return _this.Get_().(D4_DC15).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf21.Equals(data2.Cf21) && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf20 == data2.Cf20
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf25.Equals(data2.Cf25)
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

// Definition of trait T0
type T0 interface {
	String() string
	F9() _dafny.Sequence
	F9_set_(value _dafny.Sequence)
	Fm1(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int
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
	_f0 _dafny.Int
	_f1 _dafny.Int
	_f2 _dafny.Int
	_f3 _dafny.Set
	_f4 _dafny.Int
	_f5 _dafny.CodePoint
	_f6 bool
	_f7 _dafny.Int
	_f8 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.Zero
	_this._f3 = _dafny.EmptySet
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f8 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Set, f4 _dafny.Int, f5 _dafny.CodePoint, f6 bool, f7 _dafny.Int, f8 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
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
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Set {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.CodePoint {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f9 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f9 = _dafny.EmptySeq
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

func (_this *C0) F9() _dafny.Sequence {
	return _this._f9
}
func (_this *C0) F9_set_(value _dafny.Sequence) {
	_this._f9 = value
}
func (_this *C0) Ctor__(f9 _dafny.Sequence) {
	{
		(_this)._f9 = f9
	}
}
func (_this *C0) Fm1(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(134))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_240_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(141)
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-967))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_241_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(81)
		})))).Cardinality())).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, !(false), (Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(486), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), true, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(308)))).Cardinality())).Dtor_cf4()), _dafny.SeqOf(true))).Cardinality())))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
