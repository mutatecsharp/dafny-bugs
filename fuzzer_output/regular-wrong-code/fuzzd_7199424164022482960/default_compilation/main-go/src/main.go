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
func (_static *CompanionStruct_Default___) Fm0(p0 D0, globalState *GlobalState) _dafny.Int {
	if true {
		return _dafny.IntOfInt64(772)
	} else {
		return (_dafny.IntOfInt64(914)).Times(_dafny.IntOfInt64(333))
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, globalState *GlobalState) bool {
	return ((Companion_D0_.Create_DC0_((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(137)
	}))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tq"), _dafny.IntOfInt64(-356)), _dafny.MultiSetOf(_dafny.IntOfInt64(-192), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()))).Dtor_cf0()).Cmp(((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-610))))).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())) <= 0
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.IntOfInt64(392))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("m")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("xnjhhmdfk"), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tslm")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("lhmcdhe"), _dafny.IntOfInt64(798))), (_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xoksdycy")).Cardinality()))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(298))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(!(false)), false, false, true)).Difference((_dafny.SetOf(false, true, false)).Difference(_dafny.SetOf(!(true), false)))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('l')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(780))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('o')))
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('c')
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true, false)).Difference(_dafny.MultiSetOf(false, true, false))).Union((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true, true, !(true)))).Intersection(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Map {
	return ((func() D4 {
		if false {
			return Companion_D4_.Create_DC11_(func() _dafny.Map {
				var _coll0 = _dafny.NewMapBuilder()
				_ = _coll0
				for _iter0 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(649))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_2_i0 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(204), _dafny.IntOfInt64(26))).Cardinality())
				}))).Elements()); ; {
					_compr_0, _ok0 := _iter0()
					if !_ok0 {
						break
					}
					var _3_v0 _dafny.Int
					_3_v0 = interface{}(_compr_0).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(649))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg3 _dafny.Int) interface{} {
							return coer3(arg3)
						}
					}(func(_2_i0 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(204), _dafny.IntOfInt64(26))).Cardinality())
					})), _3_v0) {
						_coll0.Add(Companion_Default___.SafeDivisionInt(_3_v0, _dafny.IntOfInt64(308)), true)
					}
				}
				return _coll0.ToMap()
			}())
		}
		return Companion_D4_.Create_DC11_(func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())).Keys().Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _4_v2 _dafny.Int
					_4_v2 = interface{}(_compr_2).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())).Contains(_4_v2) {
						_coll2.Add((_4_v2).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qc")).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mrjysrs")).Cardinality()))
					}
				}
				return _coll2.ToMap()
			}()).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _5_v1 _dafny.Int
				_5_v1 = interface{}(_compr_1).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll3 = _dafny.NewMapBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())).Keys().Elements()); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _4_v2 _dafny.Int
						_4_v2 = interface{}(_compr_3).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())).Contains(_4_v2) {
							_coll3.Add((_4_v2).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qc")).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mrjysrs")).Cardinality()))
						}
					}
					return _coll3.ToMap()
				}()).Contains(_5_v1) {
					_coll1.Add((_5_v1).Minus(_dafny.IntOfInt64(432)), true)
				}
			}
			return _coll1.ToMap()
		}())
	})()).Dtor_cf19()
}
func (_static *CompanionStruct_Default___) M0(p0 D0, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Sequence, _dafny.Set, bool) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Sequence = _dafny.EmptySeq
	_ = r1
	var r2 _dafny.Set = _dafny.EmptySet
	_ = r2
	var r3 bool = false
	_ = r3
	var _6_v0 _dafny.Sequence
	_ = _6_v0
	_6_v0 = _dafny.UnicodeSeqOfUtf8Bytes("hkx")
	var _7_v1 _dafny.Sequence
	_ = _7_v1
	_7_v1 = _dafny.SeqOf(_6_v0)
	var _8_v2 bool
	_ = _8_v2
	_8_v2 = true
	var _9_v3 _dafny.Map
	_ = _9_v3
	_9_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(902))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality()))
	var _11_v4 _dafny.Map
	_ = _11_v4
	_11_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()), _8_v2)
	var _12_v5 _dafny.Sequence
	_ = _12_v5
	_12_v5 = _dafny.SeqOf(_8_v2, _8_v2)
	var _13_v6 _dafny.Map
	_ = _13_v6
	_13_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_v0, p1)
	var _14_v7 _dafny.CodePoint
	_ = _14_v7
	_14_v7 = _dafny.CodePoint('i')
	var _15_v8 _dafny.Map
	_ = _15_v8
	_15_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v7, true)
	var _16_v9 _dafny.MultiSet
	_ = _16_v9
	_16_v9 = _dafny.MultiSetOf((_15_v8).Cardinality(), _dafny.IntOfInt64(-545), p1, p1)
	var _17_v10 D0
	_ = _17_v10
	_17_v10 = Companion_D0_.Create_DC0_(p1, _13_v6, _16_v9)
	var _18_v11 _dafny.Map
	_ = _18_v11
	_18_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_17_v10, globalState), _6_v0)
	var _19_v12 _dafny.Map
	_ = _19_v12
	_19_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
		if (_11_v4).Contains(_dafny.IntOfUint32((_12_v5).Cardinality())) {
			return (_11_v4).Get(_dafny.IntOfUint32((_12_v5).Cardinality())).(bool)
		}
		return _8_v2
	})(), (func() _dafny.Sequence {
		if (_18_v11).Contains(Companion_Default___.Fm0(_17_v10, globalState)) {
			return (_18_v11).Get(Companion_Default___.Fm0(_17_v10, globalState)).(_dafny.Sequence)
		}
		return _6_v0
	})())
	var _20_v13 D1
	_ = _20_v13
	_20_v13 = Companion_D1_.Create_DC4_(_6_v0, _8_v2, _7_v1, _6_v0, _14_v7)
	var _21_v14 _dafny.Array
	_ = _21_v14
	var _nwElement0_0 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_7_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_7_v1).Cardinality()))).Uint32()).(_dafny.Sequence), _6_v0)
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(28))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(_6_v0, 1)
	(_nw0).ArraySet1(Companion_Default___.Fm5(_dafny.IntOfUint32((_6_v0).Cardinality()), p1, p1, (_9_v3).Cardinality(), globalState), 2)
	(_nw0).ArraySet1(_6_v0, 3)
	(_nw0).ArraySet1(_6_v0, 4)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_6_v0, _6_v0), 5)
	(_nw0).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}((func(_22_v0 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
		return func(_23_i1 _dafny.Int) _dafny.CodePoint {
			return (_22_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.IntOfUint32((_22_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
		}
	})(_6_v0))), 6)
	(_nw0).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hrekd"), 7)
	(_nw0).ArraySet1(_6_v0, 8)
	(_nw0).ArraySet1(_6_v0, 9)
	(_nw0).ArraySet1(_6_v0, 10)
	(_nw0).ArraySet1(_6_v0, 11)
	(_nw0).ArraySet1((func() _dafny.Sequence {
		if (_19_v12).Contains(_8_v2) {
			return (_19_v12).Get(_8_v2).(_dafny.Sequence)
		}
		return _dafny.Companion_Sequence_.Update(_6_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.IntOfUint32((_6_v0).Cardinality()))).Uint32(), (_20_v13).Dtor_cf10())
	})(), 12)
	(_nw0).ArraySet1(_6_v0, 13)
	(_nw0).ArraySet1(_6_v0, 14)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_6_v0, _6_v0), 15)
	(_nw0).ArraySet1(_6_v0, 16)
	(_nw0).ArraySet1(_6_v0, 17)
	(_nw0).ArraySet1(_6_v0, 18)
	(_nw0).ArraySet1(_6_v0, 19)
	(_nw0).ArraySet1(_6_v0, 20)
	(_nw0).ArraySet1(_6_v0, 21)
	(_nw0).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mxsg"), 22)
	(_nw0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_6_v0, _6_v0), 23)
	(_nw0).ArraySet1(_6_v0, 24)
	(_nw0).ArraySet1(_6_v0, 25)
	(_nw0).ArraySet1(_6_v0, 26)
	(_nw0).ArraySet1(_6_v0, 27)
	_21_v14 = _nw0
	_21_v14 = _21_v14
	var _24_v15 *C0
	_ = _24_v15
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__((p1).Plus(p1))
	_24_v15 = _nw1
	var _25_v16 _dafny.Set
	_ = _25_v16
	_25_v16 = _dafny.SetOf(_8_v2, _8_v2)
	var _26_v17 D1
	_ = _26_v17
	_26_v17 = Companion_D1_.Create_DC3_(_25_v16)
	var _pat_let_tv0 = _6_v0
	_ = _pat_let_tv0
	var _pat_let_tv1 = _6_v0
	_ = _pat_let_tv1
	var _pat_let_tv2 = _8_v2
	_ = _pat_let_tv2
	var _pat_let_tv3 = _8_v2
	_ = _pat_let_tv3
	var _pat_let_tv4 = _8_v2
	_ = _pat_let_tv4
	r3 = func(_source0 D1) bool {
		if _source0.Is_DC3() {
			var _27___mcc_h0 _dafny.Set = _source0.Get_().(D1_DC3).Cf5
			_ = _27___mcc_h0
			var _28_cf5 _dafny.Set = _27___mcc_h0
			_ = _28_cf5
			return _dafny.Companion_Sequence_.IsProperPrefixOf(_pat_let_tv0, _pat_let_tv1)
		} else if _source0.Is_DC4() {
			var _29___mcc_h1 _dafny.Sequence = _source0.Get_().(D1_DC4).Cf6
			_ = _29___mcc_h1
			var _30___mcc_h2 bool = _source0.Get_().(D1_DC4).Cf7
			_ = _30___mcc_h2
			var _31___mcc_h3 _dafny.Sequence = _source0.Get_().(D1_DC4).Cf8
			_ = _31___mcc_h3
			var _32___mcc_h4 _dafny.Sequence = _source0.Get_().(D1_DC4).Cf9
			_ = _32___mcc_h4
			var _33___mcc_h5 _dafny.CodePoint = _source0.Get_().(D1_DC4).Cf10
			_ = _33___mcc_h5
			var _34_cf10 _dafny.CodePoint = _33___mcc_h5
			_ = _34_cf10
			var _35_cf9 _dafny.Sequence = _32___mcc_h4
			_ = _35_cf9
			var _36_cf8 _dafny.Sequence = _31___mcc_h3
			_ = _36_cf8
			var _37_cf7 bool = _30___mcc_h2
			_ = _37_cf7
			var _38_cf6 _dafny.Sequence = _29___mcc_h1
			_ = _38_cf6
			return _37_cf7
		} else if _source0.Is_DC2() {
			var _39___mcc_h6 _dafny.Map = _source0.Get_().(D1_DC2).Cf4
			_ = _39___mcc_h6
			var _40_cf4 _dafny.Map = _39___mcc_h6
			_ = _40_cf4
			return _pat_let_tv2
		} else {
			var _41___mcc_h7 D1 = _source0.Get_().(D1_DC5).Cf11
			_ = _41___mcc_h7
			var _42_cf11 D1 = _41___mcc_h7
			_ = _42_cf11
			return (_pat_let_tv3) || (_pat_let_tv4)
		}
	}(_26_v17)
	var _43_v18 _dafny.Array
	_ = _43_v18
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_44_v15 *C0) func(_dafny.Int) _dafny.Int {
			return func(_45_i2 _dafny.Int) _dafny.Int {
				return (_45_i2).Times((_44_v15).F18())
			}
		})(_24_v15)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw2 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw2).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw2).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_43_v18 = _nw2
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(17)
	_ = _len0_1
	var _nw3 _dafny.Array
	_ = _nw3
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw3 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = func(_46_i3 _dafny.Int) _dafny.Int {
			return (_46_i3).Times(_dafny.IntOfInt64(-986))
		}
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
	_43_v18 = _nw3
	var _hi0 _dafny.Int = (_24_v15).F18()
	_ = _hi0
	for _47_i4 := (_24_v15).F18(); _47_i4.Cmp(_hi0) < 0; _47_i4 = _47_i4.Plus(_dafny.One) {
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))
		_ = _index0
		(_43_v18).ArraySet1((_24_v15).F18(), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))
		_ = _index1
		(_43_v18).ArraySet1((func() _dafny.Int {
			if _8_v2 {
				return (_24_v15).F18()
			}
			return (p1).Times(p1)
		})(), (_index1).Int())
		_6_v0 = _dafny.Companion_Sequence_.Concatenate(_6_v0, _6_v0)
		var _48_v19 _dafny.Sequence
		_ = _48_v19
		_48_v19 = _dafny.SeqOf(_dafny.IntOfInt64(95), p1)
		var _49_v20 _dafny.Array
		_ = _49_v20
		var _nwElement0_1 bool = _8_v2
		_ = _nwElement0_1
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(21))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_1, 0)
		(_nw4).ArraySet1(_8_v2, 1)
		(_nw4).ArraySet1(_8_v2, 2)
		(_nw4).ArraySet1(_8_v2, 3)
		(_nw4).ArraySet1(_8_v2, 4)
		(_nw4).ArraySet1(_8_v2, 5)
		(_nw4).ArraySet1(_8_v2, 6)
		(_nw4).ArraySet1(_8_v2, 7)
		(_nw4).ArraySet1(_8_v2, 8)
		(_nw4).ArraySet1(_8_v2, 9)
		(_nw4).ArraySet1(_8_v2, 10)
		(_nw4).ArraySet1(true, 11)
		(_nw4).ArraySet1(_8_v2, 12)
		(_nw4).ArraySet1(_8_v2, 13)
		(_nw4).ArraySet1(_8_v2, 14)
		(_nw4).ArraySet1(_8_v2, 15)
		(_nw4).ArraySet1(_8_v2, 16)
		(_nw4).ArraySet1(_8_v2, 17)
		(_nw4).ArraySet1(_8_v2, 18)
		(_nw4).ArraySet1(_8_v2, 19)
		(_nw4).ArraySet1(_8_v2, 20)
		_49_v20 = _nw4
		var _50_v21 _dafny.Sequence
		_ = _50_v21
		_50_v21 = _dafny.SeqOf(_49_v20, _49_v20, _49_v20, _49_v20)
		_12_v5 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_12_v5, _dafny.Companion_Sequence_.Concatenate(_12_v5, _12_v5)), (Companion_Default___.SafeIndex((_17_v10).Dtor_cf0(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_12_v5, _dafny.Companion_Sequence_.Concatenate(_12_v5, _12_v5))).Cardinality()))).Uint32(), (_dafny.IntOfUint32((_48_v19).Cardinality())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_50_v21).Cardinality()))) < 0)
		var _source1 D1 = _20_v13
		_ = _source1
		if _source1.Is_DC3() {
			var _51___mcc_h8 _dafny.Set = _source1.Get_().(D1_DC3).Cf5
			_ = _51___mcc_h8
			var _52_cf5 _dafny.Set = _51___mcc_h8
			_ = _52_cf5
			(globalState).F13 = (_dafny.IntOfInt64(721)).Cmp(_47_i4) > 0
			(globalState).F13 = _8_v2
			(globalState).F10 = (_24_v15).F18()
			r3 = _8_v2
		} else if _source1.Is_DC4() {
			var _53___mcc_h9 _dafny.Sequence = _source1.Get_().(D1_DC4).Cf6
			_ = _53___mcc_h9
			var _54___mcc_h10 bool = _source1.Get_().(D1_DC4).Cf7
			_ = _54___mcc_h10
			var _55___mcc_h11 _dafny.Sequence = _source1.Get_().(D1_DC4).Cf8
			_ = _55___mcc_h11
			var _56___mcc_h12 _dafny.Sequence = _source1.Get_().(D1_DC4).Cf9
			_ = _56___mcc_h12
			var _57___mcc_h13 _dafny.CodePoint = _source1.Get_().(D1_DC4).Cf10
			_ = _57___mcc_h13
			var _58_cf10 _dafny.CodePoint = _57___mcc_h13
			_ = _58_cf10
			var _59_cf9 _dafny.Sequence = _56___mcc_h12
			_ = _59_cf9
			var _60_cf8 _dafny.Sequence = _55___mcc_h11
			_ = _60_cf8
			var _61_cf7 bool = _54___mcc_h10
			_ = _61_cf7
			var _62_cf6 _dafny.Sequence = _53___mcc_h9
			_ = _62_cf6
			var _63_v22 _dafny.Sequence
			_ = _63_v22
			_63_v22 = _dafny.SeqOf(_24_v15, _24_v15, _24_v15)
			_63_v22 = _dafny.Companion_Sequence_.Concatenate(_63_v22, _63_v22)
			var _64_v23 _dafny.Map
			_ = _64_v23
			_64_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v19, (_24_v15).F18())
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))
			_ = _index2
			var _rhs0 _dafny.Int = Companion_Default___.SafeDivisionInt((_24_v15).F18(), (_dafny.Zero).Minus(Companion_Default___.Fm0(_17_v10, globalState)))
			_ = _rhs0
			var _rhs1 _dafny.Map = _64_v23
			_ = _rhs1
			var _rhs2 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-273))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_65_v15 *C0, _66_v9 _dafny.MultiSet, _67_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_68_i5 _dafny.Int) _dafny.Int {
					return (_dafny.SetOf((func() _dafny.Int {
						if (_66_v9).Contains((_65_v15).F18()) {
							return (_66_v9).Multiplicity((_65_v15).F18())
						}
						return _dafny.IntOfUint32((_67_v0).Cardinality())
					})())).Cardinality()
				}
			})(_24_v15, _16_v9, _6_v0))), _48_v19), _dafny.Companion_Sequence_.Concatenate(_48_v19, _48_v19))).Cardinality())
			_ = _rhs2
			var _lhs0 _dafny.Array = _43_v18
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
			_64_v23 = _rhs1
			_lhs2.F4 = _rhs2
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))
			_ = _index3
			(_43_v18).ArraySet1(_47_i4, (_index3).Int())
			var _69_v24 _dafny.Array
			_ = _69_v24
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_2
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.CodePoint = (func(_70_v0 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_71_i6 _dafny.Int) _dafny.CodePoint {
						return (_70_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-327), _dafny.IntOfUint32((_70_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_6_v0)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw5 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw5).ArraySet1CodePoint(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw5).ArraySet1CodePoint(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_69_v24 = _nw5
			var _72_v25 _dafny.MultiSet
			_ = _72_v25
			_72_v25 = _dafny.MultiSetOf(_69_v24)
			_72_v25 = _72_v25
		} else if _source1.Is_DC2() {
			var _73___mcc_h14 _dafny.Map = _source1.Get_().(D1_DC2).Cf4
			_ = _73___mcc_h14
			var _74_cf4 _dafny.Map = _73___mcc_h14
			_ = _74_cf4
			var _75_v26 *C0
			_ = _75_v26
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__((_47_i4).Minus(Companion_Default___.Fm0(_17_v10, globalState)))
			_75_v26 = _nw6
			r3 = ((_24_v15).F18()).Cmp(p1) > 0
			var _76_v27 _dafny.Map
			_ = _76_v27
			_76_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v26, _8_v2)
			var _77_v28 _dafny.Array
			_ = _77_v28
			var _nwElement0_2 _dafny.Map = _76_v27
			_ = _nwElement0_2
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(11))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_2, 0)
			(_nw7).ArraySet1(_76_v27, 1)
			(_nw7).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v26, !(_8_v2)), 2)
			(_nw7).ArraySet1(_76_v27, 3)
			(_nw7).ArraySet1((_76_v27).Update(_75_v26, _8_v2), 4)
			(_nw7).ArraySet1((_76_v27).Merge(_76_v27), 5)
			(_nw7).ArraySet1(_76_v27, 6)
			(_nw7).ArraySet1(_76_v27, 7)
			(_nw7).ArraySet1(_76_v27, 8)
			(_nw7).ArraySet1(_76_v27, 9)
			(_nw7).ArraySet1(_76_v27, 10)
			_77_v28 = _nw7
			var _78_v29 _dafny.MultiSet
			_ = _78_v29
			_78_v29 = _dafny.MultiSetOf(_8_v2)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_77_v28), 0))
			_ = _index4
			(_77_v28).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_75_v26, (_12_v5).Select((Companion_Default___.SafeIndex((_78_v29).Cardinality(), _dafny.IntOfUint32((_12_v5).Cardinality()))).Uint32()).(bool))).Merge(_76_v27), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_77_v28), 0))
			_ = _index5
			(_77_v28).ArraySet1(_76_v27, (_index5).Int())
			var _79_v30 _dafny.Map
			_ = _79_v30
			_79_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_24_v15).F18(), (_24_v15).F18())
			var _80_v31 _dafny.Array
			_ = _80_v31
			var _nwElement0_3 _dafny.Int = _47_i4
			_ = _nwElement0_3
			var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(14))
			_ = _nw8
			(_nw8).ArraySet1(_nwElement0_3, 0)
			(_nw8).ArraySet1((_43_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))).Int()).(_dafny.Int), 1)
			(_nw8).ArraySet1((_75_v26).F18(), 2)
			(_nw8).ArraySet1((_24_v15).F18(), 3)
			(_nw8).ArraySet1((_75_v26).F18(), 4)
			(_nw8).ArraySet1((_43_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))).Int()).(_dafny.Int), 5)
			(_nw8).ArraySet1((_dafny.Zero).Minus((_24_v15).F18()), 6)
			(_nw8).ArraySet1((_43_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))).Int()).(_dafny.Int), 7)
			(_nw8).ArraySet1((_dafny.Zero).Minus((_24_v15).F18()), 8)
			(_nw8).ArraySet1((_24_v15).F18(), 9)
			(_nw8).ArraySet1((_24_v15).F18(), 10)
			(_nw8).ArraySet1(((_79_v30).Update(p1, _dafny.IntOfUint32((_12_v5).Cardinality()))).Cardinality(), 11)
			(_nw8).ArraySet1((_43_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_43_v18), 0))).Int()).(_dafny.Int), 12)
			(_nw8).ArraySet1((_24_v15).F18(), 13)
			_80_v31 = _nw8
			var _81_v32 D3
			_ = _81_v32
			_81_v32 = Companion_D3_.Create_DC10_((func() bool {
				if _8_v2 {
					return true
				}
				return _8_v2
			})(), _80_v31, _dafny.Companion_Sequence_.Concatenate(_6_v0, _6_v0))
			_81_v32 = _81_v32
		} else {
			var _82___mcc_h15 D1 = _source1.Get_().(D1_DC5).Cf11
			_ = _82___mcc_h15
			var _83_cf11 D1 = _82___mcc_h15
			_ = _83_cf11
			(globalState).F0 = _47_i4
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_49_v20), 0))
			_ = _index6
			(_49_v20).ArraySet1(_8_v2, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_49_v20), 0))
			_ = _index7
			(_49_v20).ArraySet1(_8_v2, (_index7).Int())
			var _84_v33 _dafny.Array
			_ = _84_v33
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_3
			var _nw9 _dafny.Array
			_ = _nw9
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw9 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_85_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_86_i7 _dafny.Int) _dafny.CodePoint {
						return _85_v7
					}
				})(_14_v7)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw9 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw9).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw9).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_84_v33 = _nw9
			_84_v33 = _84_v33
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_49_v20), 0))
			_ = _index8
			(_49_v20).ArraySet1(!((_49_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_49_v20), 0))).Int()).(bool)) || ((_49_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_49_v20), 0))).Int()).(bool)), (_index8).Int())
		}
	}
	var _pat_let_tv5 = _7_v1
	_ = _pat_let_tv5
	var _pat_let_tv6 = _14_v7
	_ = _pat_let_tv6
	var _source2 D1 = func(_pat_let0_0 D1) D1 {
		return func(_87_dt__update__tmp_h0 D1) D1 {
			return func(_pat_let1_0 _dafny.Sequence) D1 {
				return func(_88_dt__update_hcf8_h0 _dafny.Sequence) D1 {
					return func(_pat_let2_0 _dafny.Sequence) D1 {
						return func(_91_dt__update_hcf6_h0 _dafny.Sequence) D1 {
							return Companion_D1_.Create_DC4_(_91_dt__update_hcf6_h0, (_87_dt__update__tmp_h0).Dtor_cf7(), _88_dt__update_hcf8_h0, (_87_dt__update__tmp_h0).Dtor_cf9(), (_87_dt__update__tmp_h0).Dtor_cf10())
						}(_pat_let2_0)
					}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(335))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg7 _dafny.Int) interface{} {
							return coer7(arg7)
						}
					}((func(_89_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_90_i8 _dafny.Int) _dafny.CodePoint {
							return _89_v7
						}
					})(_pat_let_tv6))))
				}(_pat_let1_0)
			}(_pat_let_tv5)
		}(_pat_let0_0)
	}(_20_v13)
	_ = _source2
	if _source2.Is_DC3() {
		var _92___mcc_h16 _dafny.Set = _source2.Get_().(D1_DC3).Cf5
		_ = _92___mcc_h16
		var _93_cf5 _dafny.Set = _92___mcc_h16
		_ = _93_cf5
		var _94_v34 _dafny.Array
		_ = _94_v34
		var _nw10 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw10
		_94_v34 = _nw10
		var _95_v35 D2
		_ = _95_v35
		_95_v35 = Companion_D2_.Create_DC7_(_94_v34)
		var _96_v36 D2
		_ = _96_v36
		_96_v36 = Companion_D2_.Create_DC8_(_95_v35)
		_96_v36 = _96_v36
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_94_v34), 0))
		_ = _index9
		(_94_v34).ArraySet1(!(_8_v2) || (_8_v2), (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(898), _dafny.ArrayLen((_94_v34), 0))
		_ = _index10
		(_94_v34).ArraySet1(_8_v2, (_index10).Int())
		_94_v34 = (func() _dafny.Array {
			if _8_v2 {
				return _94_v34
			}
			return _94_v34
		})()
		var _97_v37 D0
		_ = _97_v37
		_97_v37 = Companion_D0_.Create_DC1_(Companion_Default___.Fm3(globalState))
		_97_v37 = p0
	} else if _source2.Is_DC4() {
		var _98___mcc_h17 _dafny.Sequence = _source2.Get_().(D1_DC4).Cf6
		_ = _98___mcc_h17
		var _99___mcc_h18 bool = _source2.Get_().(D1_DC4).Cf7
		_ = _99___mcc_h18
		var _100___mcc_h19 _dafny.Sequence = _source2.Get_().(D1_DC4).Cf8
		_ = _100___mcc_h19
		var _101___mcc_h20 _dafny.Sequence = _source2.Get_().(D1_DC4).Cf9
		_ = _101___mcc_h20
		var _102___mcc_h21 _dafny.CodePoint = _source2.Get_().(D1_DC4).Cf10
		_ = _102___mcc_h21
		var _103_cf10 _dafny.CodePoint = _102___mcc_h21
		_ = _103_cf10
		var _104_cf9 _dafny.Sequence = _101___mcc_h20
		_ = _104_cf9
		var _105_cf8 _dafny.Sequence = _100___mcc_h19
		_ = _105_cf8
		var _106_cf7 bool = _99___mcc_h18
		_ = _106_cf7
		var _107_cf6 _dafny.Sequence = _98___mcc_h17
		_ = _107_cf6
		var _108_v38 _dafny.Array
		_ = _108_v38
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_4
		var _nw11 _dafny.Array
		_ = _nw11
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw11 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) bool = (func(_109_v5 _dafny.Sequence, _110_v15 *C0, _111_v3 _dafny.Map, _112_p1 _dafny.Int) func(_dafny.Int) bool {
				return func(_113_i9 _dafny.Int) bool {
					return !(_111_v3).Contains((_109_v5).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_110_v15).F18(), (_110_v15).F18(), (_111_v3).Cardinality())).Cardinality(), _112_p1)).Cardinality(), _dafny.IntOfUint32((_109_v5).Cardinality()))).Uint32()).(bool))
				}
			})(_12_v5, _24_v15, _9_v3, p1)
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
		_108_v38 = _nw11
		var _rhs3 *C0 = _24_v15
		_ = _rhs3
		var _rhs4 _dafny.Array = _108_v38
		_ = _rhs4
		_24_v15 = _rhs3
		_108_v38 = _rhs4
		var _114_v39 _dafny.MultiSet
		_ = _114_v39
		_114_v39 = _dafny.MultiSetOf(_106_cf7, _106_cf7)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_43_v18), 0))
		_ = _index11
		(_43_v18).ArraySet1(((func() _dafny.MultiSet {
			if _8_v2 {
				return _114_v39
			}
			return Companion_Default___.Fm7(p1, _12_v5, _8_v2, globalState)
		})()).Cardinality(), (_index11).Int())
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(62), _dafny.ArrayLen((_43_v18), 0))
		_ = _index12
		(_43_v18).ArraySet1(p1, (_index12).Int())
		var _115_v40 _dafny.Map
		_ = _115_v40
		_115_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v15, (_24_v15).F18())
		_115_v40 = ((_115_v40).Merge(_115_v40)).Merge(_115_v40)
		_14_v7 = _14_v7
	} else if _source2.Is_DC2() {
		var _116___mcc_h22 _dafny.Map = _source2.Get_().(D1_DC2).Cf4
		_ = _116___mcc_h22
		var _117_cf4 _dafny.Map = _116___mcc_h22
		_ = _117_cf4
		var _118_v41 _dafny.Map
		_ = _118_v41
		_118_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v2, true)
		_118_v41 = (_118_v41).Merge(_118_v41)
		if _8_v2 {
			var _119_v42 _dafny.Array
			_ = _119_v42
			var _nw12 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw12
			_119_v42 = _nw12
			var _120_v43 _dafny.Sequence
			_ = _120_v43
			_120_v43 = _dafny.SeqOf(_119_v42)
			var _121_v44 _dafny.Map
			_ = _121_v44
			_121_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_24_v15, _120_v43)
			var _122_v45 _dafny.Map
			_ = _122_v45
			_122_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v2, _43_v18)
			var _123_v46 _dafny.Map
			_ = _123_v46
			_123_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v45).Cardinality(), _120_v43)
			var _124_v47 _dafny.Array
			_ = _124_v47
			var _nwElement0_4 _dafny.Sequence = _120_v43
			_ = _nwElement0_4
			var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
			_ = _nw13
			(_nw13).ArraySet1(_nwElement0_4, 0)
			(_nw13).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_120_v43, (func() _dafny.Sequence {
				if (_121_v44).Contains(_24_v15) {
					return (_121_v44).Get(_24_v15).(_dafny.Sequence)
				}
				return _120_v43
			})()), 1)
			(_nw13).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_120_v43, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_120_v43).Cardinality()))).Uint32(), _119_v42), _120_v43), 2)
			(_nw13).ArraySet1(_120_v43, 3)
			(_nw13).ArraySet1((func() _dafny.Sequence {
				if (_123_v46).Contains(p1) {
					return (_123_v46).Get(p1).(_dafny.Sequence)
				}
				return _120_v43
			})(), 4)
			(_nw13).ArraySet1(_120_v43, 5)
			(_nw13).ArraySet1(_120_v43, 6)
			(_nw13).ArraySet1(_120_v43, 7)
			(_nw13).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_120_v43, _120_v43), 8)
			_124_v47 = _nw13
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_124_v47), 0))
			_ = _index13
			(_124_v47).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_119_v42, _119_v42), _120_v43), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_43_v18), 0))
			_ = _index14
			(_43_v18).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_25_v16).Cardinality())), (_index14).Int())
			var _125_v48 D3
			_ = _125_v48
			_125_v48 = Companion_D3_.Create_DC10_(!(_8_v2) || (false), _43_v18, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-28))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_126_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_127_i10 _dafny.Int) _dafny.CodePoint {
					return _126_v7
				}
			})(_14_v7))), _dafny.UnicodeSeqOfUtf8Bytes("i")), (Companion_Default___.SafeIndex((_24_v15).F18(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-28))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_128_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_129_i10 _dafny.Int) _dafny.CodePoint {
					return _128_v7
				}
			})(_14_v7))), _dafny.UnicodeSeqOfUtf8Bytes("i"))).Cardinality()))).Uint32(), _14_v7))
			var _130_v49 _dafny.Map
			_ = _130_v49
			_130_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v2, _24_v15)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_124_v47), 0))
			_ = _index15
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_43_v18), 0))
			_ = _index16
			var _rhs5 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_120_v43, _120_v43)
			_ = _rhs5
			var _rhs6 _dafny.Sequence = _6_v0
			_ = _rhs6
			var _rhs7 _dafny.Int = ((_130_v49).Merge(_130_v49)).Cardinality()
			_ = _rhs7
			var _rhs8 D3 = _125_v48
			_ = _rhs8
			var _lhs3 _dafny.Array = _124_v47
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_124_v47), 0))
			_ = _lhs4
			var _lhs5 _dafny.Array = _43_v18
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_43_v18), 0))
			_ = _lhs6
			(_lhs3).ArraySet1(_rhs5, (_lhs4).Int())
			_6_v0 = _rhs6
			(_lhs5).ArraySet1(_rhs7, (_lhs6).Int())
			_125_v48 = _rhs8
			var _131_v50 _dafny.Map
			_ = _131_v50
			_131_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_43_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_43_v18), 0))).Int()).(_dafny.Int), (_24_v15).F18())
			_131_v50 = (_131_v50).Update((_43_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(368), _dafny.ArrayLen((_43_v18), 0))).Int()).(_dafny.Int), (_24_v15).F18())
			(globalState).F0 = p1
			_119_v42 = _119_v42
			var _132_v51 _dafny.MultiSet
			_ = _132_v51
			_132_v51 = _dafny.MultiSetOf(!(_8_v2), true)
			var _133_v52 *C0
			_ = _133_v52
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(((_132_v51).Intersection(_132_v51)).Cardinality())
			_133_v52 = _nw14
		} else {
			var _134_v53 _dafny.Array
			_ = _134_v53
			var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(20))
			_ = _nw15
			_134_v53 = _nw15
			_134_v53 = _134_v53
			var _135_v54 D1
			_ = _135_v54
			_135_v54 = Companion_D1_.Create_DC2_(_9_v3)
			var _136_v55 _dafny.Map
			_ = _136_v55
			_136_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_135_v54, (_24_v15).F18())
			_136_v55 = _136_v55
			var _137_v56 _dafny.Sequence
			_ = _137_v56
			_137_v56 = _dafny.SeqOf((_24_v15).F18(), (_24_v15).F18(), (_24_v15).F18())
			var _138_v57 _dafny.Map
			_ = _138_v57
			_138_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_8_v2, _dafny.Companion_Sequence_.Concatenate(_137_v56, _137_v56))
			r1 = (func() _dafny.Sequence {
				if (_138_v57).Contains(_8_v2) {
					return (_138_v57).Get(_8_v2).(_dafny.Sequence)
				}
				return (func() _dafny.Sequence {
					if _8_v2 {
						return _137_v56
					}
					return _dafny.SeqOf((_dafny.Zero).Minus((_117_cf4).Cardinality()))
				})()
			})()
			var _139_v58 _dafny.Array
			_ = _139_v58
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_5
			var _nw16 _dafny.Array
			_ = _nw16
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw16 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_140_v2 bool) func(_dafny.Int) bool {
					return func(_141_i11 _dafny.Int) bool {
						return _140_v2
					}
				})(_8_v2)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw16 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw16).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw16).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_139_v58 = _nw16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_139_v58), 0))
			_ = _index17
			(_139_v58).ArraySet1(true, (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.ArrayLen((_139_v58), 0))
			_ = _index18
			(_139_v58).ArraySet1(_8_v2, (_index18).Int())
			(globalState).F15 = (_24_v15).F18()
		}
		var _142_v59 _dafny.MultiSet
		_ = _142_v59
		_142_v59 = _dafny.MultiSetOf(false)
		r3 = ((_dafny.MultiSetOf(_8_v2)).Update(!(_8_v2), Companion_Default___.Abs(p1))).IsSubsetOf((_142_v59).Difference(_dafny.MultiSetOf(!(Companion_Default___.Fm1(_8_v2, globalState)), _8_v2)))
		var _143_v60 _dafny.Array
		_ = _143_v60
		var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
		_ = _nw17
		_143_v60 = _nw17
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
		_ = _nw18
		_143_v60 = _nw18
	} else {
		var _144___mcc_h23 D1 = _source2.Get_().(D1_DC5).Cf11
		_ = _144___mcc_h23
		var _145_cf11 D1 = _144___mcc_h23
		_ = _145_cf11
		var _146_v61 _dafny.Sequence
		_ = _146_v61
		_146_v61 = _dafny.SeqOf(_12_v5, _12_v5, _12_v5)
		var _147_v62 _dafny.Map
		_ = _147_v62
		_147_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-890), _14_v7)
		var _148_v63 _dafny.Sequence
		_ = _148_v63
		_148_v63 = _dafny.SeqOf((Companion_Default___.Fm8(globalState)).Cardinality(), (_24_v15).F18())
		var _149_v64 _dafny.Sequence
		_ = _149_v64
		_149_v64 = _dafny.SeqOf(_dafny.IntOfUint32((_148_v63).Cardinality()))
		var _150_v65 D3
		_ = _150_v65
		_150_v65 = Companion_D3_.Create_DC10_(_8_v2, _43_v18, _6_v0)
		var _151_v66 _dafny.Array
		_ = _151_v66
		var _nwElement0_5 bool = true
		_ = _nwElement0_5
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(22))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_5, 0)
		(_nw19).ArraySet1((_20_v13).Dtor_cf7(), 1)
		(_nw19).ArraySet1(_8_v2, 2)
		(_nw19).ArraySet1(Companion_Default___.Fm1(_8_v2, globalState), 3)
		(_nw19).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_146_v61, _dafny.SeqOf(_12_v5)), 4)
		(_nw19).ArraySet1(!(true), 5)
		(_nw19).ArraySet1(_8_v2, 6)
		(_nw19).ArraySet1(_8_v2, 7)
		(_nw19).ArraySet1(((_24_v15).F18()).Cmp((_24_v15).F18()) != 0, 8)
		(_nw19).ArraySet1(Companion_Default___.Fm1(_8_v2, globalState), 9)
		(_nw19).ArraySet1(_8_v2, 10)
		(_nw19).ArraySet1((_147_v62).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_149_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_6_v0).Cardinality()), _dafny.IntOfUint32((_149_v64).Cardinality()))).Uint32()).(_dafny.Int), _14_v7)), 11)
		(_nw19).ArraySet1(!(_8_v2), 12)
		(_nw19).ArraySet1(true, 13)
		(_nw19).ArraySet1((_12_v5).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_12_v5).Cardinality()))).Uint32()).(bool), 14)
		(_nw19).ArraySet1(_8_v2, 15)
		(_nw19).ArraySet1(((_24_v15).F18()).Cmp(p1) > 0, 16)
		(_nw19).ArraySet1(!(_8_v2), 17)
		(_nw19).ArraySet1((_150_v65).Dtor_cf16(), 18)
		(_nw19).ArraySet1(_8_v2, 19)
		(_nw19).ArraySet1(_8_v2, 20)
		(_nw19).ArraySet1(Companion_Default___.Fm1(!(false), globalState), 21)
		_151_v66 = _nw19
		_151_v66 = _151_v66
		var _152_v67 _dafny.Array
		_ = _152_v67
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
		_ = _nw20
		_152_v67 = _nw20
		_152_v67 = _152_v67
		r3 = _8_v2
		var _153_v68 _dafny.Map
		_ = _153_v68
		_153_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v66, _24_v15)
		_153_v68 = _153_v68
	}
	var _pat_let_tv7 = _24_v15
	_ = _pat_let_tv7
	r0 = (_dafny.Zero).Minus(func(_source3 D0) _dafny.Int {
		if _source3.Is_DC0() {
			var _154___mcc_h24 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
			_ = _154___mcc_h24
			var _155___mcc_h25 _dafny.Map = _source3.Get_().(D0_DC0).Cf1
			_ = _155___mcc_h25
			var _156___mcc_h26 _dafny.MultiSet = _source3.Get_().(D0_DC0).Cf2
			_ = _156___mcc_h26
			var _157_cf2 _dafny.MultiSet = _156___mcc_h26
			_ = _157_cf2
			var _158_cf1 _dafny.Map = _155___mcc_h25
			_ = _158_cf1
			var _159_cf0 _dafny.Int = _154___mcc_h24
			_ = _159_cf0
			return (_dafny.Zero).Minus((_159_cf0).Times(_dafny.IntOfInt64(483)))
		} else {
			var _160___mcc_h27 D0 = _source3.Get_().(D0_DC1).Cf3
			_ = _160___mcc_h27
			var _161_cf3 D0 = _160___mcc_h27
			_ = _161_cf3
			return (_pat_let_tv7).F18()
		}
	}(p0))
	var _162_v69 _dafny.Sequence
	_ = _162_v69
	_162_v69 = _dafny.SeqOf((_24_v15).F18(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jxu")).Cardinality()))
	r1 = _162_v69
	var _163_v70 _dafny.Set
	_ = _163_v70
	_163_v70 = _dafny.SetOf(p1, Companion_Default___.Fm0(_17_v10, globalState))
	r2 = (_163_v70).Union(_163_v70)
	r3 = _8_v2
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _164_v0 bool
	_ = _164_v0
	_164_v0 = true
	var _165_v1 _dafny.Set
	_ = _165_v1
	_165_v1 = _dafny.SetOf(_164_v0)
	var _166_v2 _dafny.Array
	_ = _166_v2
	var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
	_ = _nw21
	_166_v2 = _nw21
	var _167_v3 _dafny.Int
	_ = _167_v3
	_167_v3 = _dafny.IntOfInt64(639)
	var _168_v4 _dafny.MultiSet
	_ = _168_v4
	_168_v4 = _dafny.MultiSetOf(_167_v3, _dafny.IntOfInt64(-408))
	var _169_v5 _dafny.MultiSet
	_ = _169_v5
	_169_v5 = _dafny.MultiSetOf(_164_v0, _164_v0)
	var _170_v6 _dafny.Sequence
	_ = _170_v6
	_170_v6 = _dafny.SeqOf((_168_v4).Cardinality(), (_169_v5).Cardinality())
	var _171_v7 _dafny.Array
	_ = _171_v7
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_6
	var _nw22 _dafny.Array
	_ = _nw22
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw22 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = (func(_172_v0 bool) func(_dafny.Int) bool {
			return func(_173_i0 _dafny.Int) bool {
				return _172_v0
			}
		})(_164_v0)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw22 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw22).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw22).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_171_v7 = _nw22
	var _174_v8 _dafny.Sequence
	_ = _174_v8
	_174_v8 = _dafny.UnicodeSeqOfUtf8Bytes("mgx")
	var _175_globalState *GlobalState
	_ = _175_globalState
	var _nw23 *GlobalState = New_GlobalState_()
	_ = _nw23
	_nw23.Ctor__(_dafny.IntOfInt64(-206), (_165_v1).Intersection(_165_v1), false, false, _dafny.IntOfInt64(633), _dafny.IntOfInt64(769), true, _166_v2, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_170_v6, (Companion_Default___.SafeIndex(_167_v3, _dafny.IntOfUint32((_170_v6).Cardinality()))).Uint32(), _167_v3), _170_v6), _168_v4, _dafny.IntOfInt64(-738), _dafny.IntOfInt64(714), _dafny.CodePoint('r'), false, _dafny.IntOfInt64(760), _dafny.IntOfInt64(654), _171_v7, _174_v8)
	_175_globalState = _nw23
	var _176_v9 _dafny.Map
	_ = _176_v9
	_176_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v3, _167_v3)
	var _177_v10 _dafny.Map
	_ = _177_v10
	_177_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v0, (_176_v9).Cardinality())
	var _178_v11 _dafny.Map
	_ = _178_v11
	_178_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v8, (func() _dafny.Int {
		if (_177_v10).Contains(_164_v0) {
			return (_177_v10).Get(_164_v0).(_dafny.Int)
		}
		return _167_v3
	})())
	var _179_v12 D0
	_ = _179_v12
	_179_v12 = Companion_D0_.Create_DC0_(_167_v3, _178_v11, _168_v4)
	var _source4 D0 = _179_v12
	_ = _source4
	if _source4.Is_DC0() {
		var _180___mcc_h0 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
		_ = _180___mcc_h0
		var _181___mcc_h1 _dafny.Map = _source4.Get_().(D0_DC0).Cf1
		_ = _181___mcc_h1
		var _182___mcc_h2 _dafny.MultiSet = _source4.Get_().(D0_DC0).Cf2
		_ = _182___mcc_h2
		var _183_cf2 _dafny.MultiSet = _182___mcc_h2
		_ = _183_cf2
		var _184_cf1 _dafny.Map = _181___mcc_h1
		_ = _184_cf1
		var _185_cf0 _dafny.Int = _180___mcc_h0
		_ = _185_cf0
		var _186_v13 D0
		_ = _186_v13
		_186_v13 = Companion_D0_.Create_DC1_(Companion_D0_.Create_DC1_(Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_170_v6).Cardinality()), _178_v11, _183_cf2)))
		var _187_v14 D0
		_ = _187_v14
		_187_v14 = Companion_D0_.Create_DC1_(_186_v13)
		var _188_v15 _dafny.Int
		_ = _188_v15
		var _189_v16 _dafny.Sequence
		_ = _189_v16
		var _190_v17 _dafny.Set
		_ = _190_v17
		var _191_v18 bool
		_ = _191_v18
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Sequence
		_ = _out1
		var _out2 _dafny.Set
		_ = _out2
		var _out3 bool
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0(Companion_D0_.Create_DC1_(_187_v14), Companion_Default___.Fm0(_179_v12, _175_globalState), _175_globalState)
		_188_v15 = _out0
		_189_v16 = _out1
		_190_v17 = _out2
		_191_v18 = _out3
		_164_v0 = !((Companion_Default___.Fm0(_179_v12, _175_globalState)).Cmp(((_dafny.Zero).Minus(_167_v3)).Minus(_185_cf0)) == 0)
		if Companion_Default___.Fm1(_164_v0, _175_globalState) {
			var _192_v19 D0
			_ = _192_v19
			_192_v19 = Companion_D0_.Create_DC1_(_186_v13)
			var _193_v20 _dafny.Int
			_ = _193_v20
			var _194_v21 _dafny.Sequence
			_ = _194_v21
			var _195_v22 _dafny.Set
			_ = _195_v22
			var _196_v23 bool
			_ = _196_v23
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Sequence
			_ = _out5
			var _out6 _dafny.Set
			_ = _out6
			var _out7 bool
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0((func() D0 {
				if _164_v0 {
					return _192_v19
				}
				return _192_v19
			})(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_174_v8).Cardinality()), _185_cf0), _175_globalState)
			_193_v20 = _out4
			_194_v21 = _out5
			_195_v22 = _out6
			_196_v23 = _out7
			(_175_globalState).F13 = Companion_Default___.Fm1(_196_v23, _175_globalState)
			(_175_globalState).F13 = _191_v18
			_177_v10 = (_177_v10).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_174_v8, _174_v8), _167_v3)
			(_175_globalState).F0 = _167_v3
		} else {
			_179_v12 = _179_v12
			var _197_v24 D0
			_ = _197_v24
			_197_v24 = Companion_D0_.Create_DC1_(_186_v13)
			var _198_v25 _dafny.Int
			_ = _198_v25
			var _199_v26 _dafny.Sequence
			_ = _199_v26
			var _200_v27 _dafny.Set
			_ = _200_v27
			var _201_v28 bool
			_ = _201_v28
			var _out8 _dafny.Int
			_ = _out8
			var _out9 _dafny.Sequence
			_ = _out9
			var _out10 _dafny.Set
			_ = _out10
			var _out11 bool
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0(_197_v24, (_dafny.IntOfUint32((_174_v8).Cardinality())).Times(_167_v3), _175_globalState)
			_198_v25 = _out8
			_199_v26 = _out9
			_200_v27 = _out10
			_201_v28 = _out11
			(_175_globalState).F13 = true
			_167_v3 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(472), _dafny.IntOfUint32((_174_v8).Cardinality()))
			var _202_v29 *C0
			_ = _202_v29
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__(_198_v25)
			_202_v29 = _nw24
		}
		var _203_v30 _dafny.Array
		_ = _203_v30
		var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw25
		_203_v30 = _nw25
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_203_v30), 0))
		_ = _index19
		(_203_v30).ArraySet1((_167_v3).Minus(_185_cf0), (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_203_v30), 0))
		_ = _index20
		var _rhs9 _dafny.Int = _185_cf0
		_ = _rhs9
		var _rhs10 _dafny.Int = _167_v3
		_ = _rhs10
		var _rhs11 bool = !(((_185_cf0).Times(_188_v15)).Cmp((_176_v9).Cardinality()) > 0)
		_ = _rhs11
		var _lhs7 _dafny.Array = _203_v30
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(841), _dafny.ArrayLen((_203_v30), 0))
		_ = _lhs8
		(_lhs7).ArraySet1(_rhs9, (_lhs8).Int())
		_167_v3 = _rhs10
		_164_v0 = _rhs11
	} else {
		var _204___mcc_h3 D0 = _source4.Get_().(D0_DC1).Cf3
		_ = _204___mcc_h3
		var _205_cf3 D0 = _204___mcc_h3
		_ = _205_cf3
		var _source5 D0 = (func() D0 {
			if false {
				return _179_v12
			}
			return _179_v12
		})()
		_ = _source5
		if _source5.Is_DC0() {
			var _206___mcc_h4 _dafny.Int = _source5.Get_().(D0_DC0).Cf0
			_ = _206___mcc_h4
			var _207___mcc_h5 _dafny.Map = _source5.Get_().(D0_DC0).Cf1
			_ = _207___mcc_h5
			var _208___mcc_h6 _dafny.MultiSet = _source5.Get_().(D0_DC0).Cf2
			_ = _208___mcc_h6
			var _209_cf2 _dafny.MultiSet = _208___mcc_h6
			_ = _209_cf2
			var _210_cf1 _dafny.Map = _207___mcc_h5
			_ = _210_cf1
			var _211_cf0 _dafny.Int = _206___mcc_h4
			_ = _211_cf0
			var _212_v31 _dafny.Map
			_ = _212_v31
			_212_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_169_v5).Contains(_164_v0) {
					return (_169_v5).Multiplicity(_164_v0)
				}
				return _167_v3
			})(), _164_v0)
			_164_v0 = Companion_Default___.Fm1((func() bool {
				if (_212_v31).Contains(_211_cf0) {
					return (_212_v31).Get(_211_cf0).(bool)
				}
				return _164_v0
			})(), _175_globalState)
			(_175_globalState).F15 = (func() _dafny.Int {
				if (_165_v1).IsDisjointFrom(_165_v1) {
					return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v0, _164_v0)).Cardinality()).Plus(_dafny.IntOfUint32((_174_v8).Cardinality()))
				}
				return (_dafny.Zero).Minus(_167_v3)
			})()
			var _213_v33 _dafny.Map
			_ = _213_v33
			_213_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v8, _dafny.SeqOf(!(Companion_Default___.Fm1(false, _175_globalState))))
			var _214_v34 D0
			_ = _214_v34
			_214_v34 = Companion_D0_.Create_DC0_(_211_cf0, func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_213_v33).Keys().Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _215_v32 _dafny.Sequence
					_215_v32 = interface{}(_compr_4).(_dafny.Sequence)
					if (_213_v33).Contains(_215_v32) {
						_coll4.Add(_215_v32, (_179_v12).Dtor_cf0())
					}
				}
				return _coll4.ToMap()
			}(), (_168_v4).Update((func() _dafny.Int {
				if (_168_v4).Contains(_211_cf0) {
					return (_168_v4).Multiplicity(_211_cf0)
				}
				return _167_v3
			})(), Companion_Default___.Abs(_167_v3)))
			var _216_v35 D0
			_ = _216_v35
			_216_v35 = Companion_D0_.Create_DC1_(_214_v34)
			var _217_v36 D0
			_ = _217_v36
			_217_v36 = Companion_D0_.Create_DC1_(_216_v35)
			var _218_v37 D0
			_ = _218_v37
			_218_v37 = Companion_D0_.Create_DC1_(_217_v36)
			var _219_v38 _dafny.Set
			_ = _219_v38
			_219_v38 = _dafny.SetOf(_174_v8)
			var _220_v39 _dafny.Int
			_ = _220_v39
			var _221_v40 _dafny.Sequence
			_ = _221_v40
			var _222_v41 _dafny.Set
			_ = _222_v41
			var _223_v42 bool
			_ = _223_v42
			var _out12 _dafny.Int
			_ = _out12
			var _out13 _dafny.Sequence
			_ = _out13
			var _out14 _dafny.Set
			_ = _out14
			var _out15 bool
			_ = _out15
			_out12, _out13, _out14, _out15 = Companion_Default___.M0(_218_v37, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_211_cf0))), _167_v3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(830))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_224_v8 _dafny.Sequence, _225_v3 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_226_i1 _dafny.Int) _dafny.CodePoint {
					return (_224_v8).Select((Companion_Default___.SafeIndex(_225_v3, _dafny.IntOfUint32((_224_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_174_v8, _167_v3)))).Cardinality())), _167_v3, (_dafny.Zero).Minus((_219_v38).Cardinality()))).Cardinality()), _175_globalState)
			_220_v39 = _out12
			_221_v40 = _out13
			_222_v41 = _out14
			_223_v42 = _out15
			var _227_v43 *C0
			_ = _227_v43
			var _nw26 *C0 = New_C0_()
			_ = _nw26
			_nw26.Ctor__(_211_cf0)
			_227_v43 = _nw26
			_227_v43 = _227_v43
		} else {
			var _228___mcc_h7 D0 = _source5.Get_().(D0_DC1).Cf3
			_ = _228___mcc_h7
			var _229_cf3 D0 = _228___mcc_h7
			_ = _229_cf3
			var _230_v44 *C0
			_ = _230_v44
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__((_167_v3).Minus(_167_v3))
			_230_v44 = _nw27
			_230_v44 = _230_v44
			var _231_v45 D0
			_ = _231_v45
			_231_v45 = Companion_D0_.Create_DC0_(_167_v3, _178_v11, _168_v4)
			var _232_v46 _dafny.Int
			_ = _232_v46
			var _233_v47 _dafny.Sequence
			_ = _233_v47
			var _234_v48 _dafny.Set
			_ = _234_v48
			var _235_v49 bool
			_ = _235_v49
			var _out16 _dafny.Int
			_ = _out16
			var _out17 _dafny.Sequence
			_ = _out17
			var _out18 _dafny.Set
			_ = _out18
			var _out19 bool
			_ = _out19
			_out16, _out17, _out18, _out19 = Companion_Default___.M0(Companion_D0_.Create_DC1_(_231_v45), (Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_170_v6).Cardinality()), _178_v11, Companion_Default___.Fm2(_164_v0, _175_globalState))).Dtor_cf0(), _175_globalState)
			_232_v46 = _out16
			_233_v47 = _out17
			_234_v48 = _out18
			_235_v49 = _out19
			var _236_v50 D0
			_ = _236_v50
			_236_v50 = Companion_D0_.Create_DC1_(_231_v45)
			var _237_v51 _dafny.Int
			_ = _237_v51
			var _238_v52 _dafny.Sequence
			_ = _238_v52
			var _239_v53 _dafny.Set
			_ = _239_v53
			var _240_v54 bool
			_ = _240_v54
			var _out20 _dafny.Int
			_ = _out20
			var _out21 _dafny.Sequence
			_ = _out21
			var _out22 _dafny.Set
			_ = _out22
			var _out23 bool
			_ = _out23
			_out20, _out21, _out22, _out23 = Companion_Default___.M0(_236_v50, (_169_v5).Cardinality(), _175_globalState)
			_237_v51 = _out20
			_238_v52 = _out21
			_239_v53 = _out22
			_240_v54 = _out23
			var _pat_let_tv8 = _231_v45
			_ = _pat_let_tv8
			var _241_v55 _dafny.Int
			_ = _241_v55
			var _242_v56 _dafny.Sequence
			_ = _242_v56
			var _243_v57 _dafny.Set
			_ = _243_v57
			var _244_v58 bool
			_ = _244_v58
			var _out24 _dafny.Int
			_ = _out24
			var _out25 _dafny.Sequence
			_ = _out25
			var _out26 _dafny.Set
			_ = _out26
			var _out27 bool
			_ = _out27
			_out24, _out25, _out26, _out27 = Companion_Default___.M0(func(_pat_let3_0 D0) D0 {
				return func(_245_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let4_0 D0) D0 {
						return func(_246_dt__update_hcf3_h0 D0) D0 {
							return Companion_D0_.Create_DC1_(_246_dt__update_hcf3_h0)
						}(_pat_let4_0)
					}(_pat_let_tv8)
				}(_pat_let3_0)
			}(_236_v50), ((_dafny.SetOf((_230_v44).F18(), _232_v46)).Difference(_234_v48)).Cardinality(), _175_globalState)
			_241_v55 = _out24
			_242_v56 = _out25
			_243_v57 = _out26
			_244_v58 = _out27
		}
		var _247_v59 _dafny.MultiSet
		_ = _247_v59
		_247_v59 = _dafny.MultiSetOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(_164_v0), _169_v5, _169_v5)
		var _248_v60 _dafny.Sequence
		_ = _248_v60
		_248_v60 = _dafny.SeqOf(_169_v5)
		_247_v59 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_248_v60, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-680))).Uint32(), func(coer11 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}((func(_249_v0 bool) func(_dafny.Int) _dafny.MultiSet {
			return func(_250_i2 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetFromSeq(_dafny.SeqOf(_249_v0, false, _249_v0))
			}
		})(_164_v0)))))
		(_175_globalState).F0 = _dafny.IntOfInt64(12)
		(_175_globalState).F10 = _167_v3
	}
	var _251_v62 _dafny.Sequence
	_ = _251_v62
	_251_v62 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_170_v6), _168_v4, _168_v4, _168_v4)
	var _252_v63 *C0
	_ = _252_v63
	var _nw28 *C0 = New_C0_()
	_ = _nw28
	_nw28.Ctor__((_168_v4).Cardinality())
	_252_v63 = _nw28
	var _253_v64 _dafny.Map
	_ = _253_v64
	_253_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v0, _252_v63)
	var _254_v65 D1
	_ = _254_v65
	_254_v65 = Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm1(false, _175_globalState)), _167_v3))
	var _255_v66 _dafny.Array
	_ = _255_v66
	var _nwElement0_6 _dafny.Int = Companion_Default___.Fm0(Companion_Default___.Fm3(_175_globalState), _175_globalState)
	_ = _nwElement0_6
	var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(14))
	_ = _nw29
	(_nw29).ArraySet1(_nwElement0_6, 0)
	(_nw29).ArraySet1(_167_v3, 1)
	(_nw29).ArraySet1((_167_v3).Plus(_167_v3), 2)
	(_nw29).ArraySet1(_167_v3, 3)
	(_nw29).ArraySet1(_167_v3, 4)
	(_nw29).ArraySet1(_dafny.IntOfInt64(-437), 5)
	(_nw29).ArraySet1((_170_v6).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(Companion_D0_.Create_DC0_(_167_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_174_v8, _167_v3), _168_v4), _175_globalState), _dafny.IntOfUint32((_170_v6).Cardinality()))).Uint32()).(_dafny.Int), 6)
	(_nw29).ArraySet1((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_251_v62).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _256_v61 _dafny.MultiSet
			_256_v61 = interface{}(_compr_5).(_dafny.MultiSet)
			if _dafny.Companion_Sequence_.Contains(_251_v62, _256_v61) {
				_coll5.Add(_256_v61, _167_v3)
			}
		}
		return _coll5.ToMap()
	}()).Cardinality()), 7)
	(_nw29).ArraySet1(_167_v3, 8)
	(_nw29).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_174_v8, _174_v8)).Cardinality()), 9)
	(_nw29).ArraySet1((_253_v64).Cardinality(), 10)
	(_nw29).ArraySet1(((_254_v65).Dtor_cf4()).Cardinality(), 11)
	(_nw29).ArraySet1(((_177_v10).Cardinality()).Plus((_252_v63).F18()), 12)
	(_nw29).ArraySet1(_167_v3, 13)
	_255_v66 = _nw29
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
	_ = _index21
	(_255_v66).ArraySet1(((_252_v63).F18()).Times(_dafny.IntOfUint32((_174_v8).Cardinality())), (_index21).Int())
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_255_v66), 0))
	_ = _index22
	(_255_v66).ArraySet1(((_252_v63).F18()).Times(_167_v3), (_index22).Int())
	var _257_v67 _dafny.Set
	_ = _257_v67
	_257_v67 = _dafny.SetOf((_252_v63).F18(), (_252_v63).F18())
	var _258_v68 _dafny.Sequence
	_ = _258_v68
	_258_v68 = _dafny.SeqOf(_164_v0)
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
	_ = _index23
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_255_v66), 0))
	_ = _index24
	var _rhs12 _dafny.Int = (func() _dafny.Int {
		if _164_v0 {
			return _167_v3
		}
		return (_257_v67).Cardinality()
	})()
	_ = _rhs12
	var _rhs13 _dafny.Int = Companion_Default___.SafeModuloInt(((_252_v63).F18()).Minus(_dafny.IntOfUint32((_174_v8).Cardinality())), _dafny.IntOfUint32((_258_v68).Cardinality()))
	_ = _rhs13
	var _rhs14 _dafny.Int = (_252_v63).F18()
	_ = _rhs14
	var _lhs9 *GlobalState = _175_globalState
	_ = _lhs9
	var _lhs10 _dafny.Array = _255_v66
	_ = _lhs10
	var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
	_ = _lhs11
	var _lhs12 _dafny.Array = _255_v66
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_255_v66), 0))
	_ = _lhs13
	_lhs9.F10 = _rhs12
	(_lhs10).ArraySet1(_rhs13, (_lhs11).Int())
	(_lhs12).ArraySet1(_rhs14, (_lhs13).Int())
	var _259_v69 *C0
	_ = _259_v69
	var _nw30 *C0 = New_C0_()
	_ = _nw30
	_nw30.Ctor__(_167_v3)
	_259_v69 = _nw30
	_168_v4 = _168_v4
	var _260_v70 _dafny.Map
	_ = _260_v70
	_260_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_252_v63).F18(), !(_164_v0))
	if (func() bool {
		if (_260_v70).Contains(Companion_Default___.SafeModuloInt((_252_v63).F18(), _dafny.IntOfInt64(-905))) {
			return (_260_v70).Get(Companion_Default___.SafeModuloInt((_252_v63).F18(), _dafny.IntOfInt64(-905))).(bool)
		}
		return !(Companion_Default___.Fm1(_164_v0, _175_globalState))
	})() {
		var _261_v71 _dafny.Array
		_ = _261_v71
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_7
		var _nw31 _dafny.Array
		_ = _nw31
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw31 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Map = (func(_262_v0 bool, _263_v10 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_264_i3 _dafny.Int) _dafny.Map {
					return (func() _dafny.Map {
						if _262_v0 {
							return _263_v10
						}
						return _263_v10
					})()
				}
			})(_164_v0, _177_v10)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw31 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw31).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw31).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_261_v71 = _nw31
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_261_v71), 0))
		_ = _index25
		(_261_v71).ArraySet1(_177_v10, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_261_v71), 0))
		_ = _index26
		(_261_v71).ArraySet1(_177_v10, (_index26).Int())
		var _265_v72 _dafny.Array
		_ = _265_v72
		var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(13))
		_ = _nw32
		_265_v72 = _nw32
		var _266_v73 _dafny.Array
		_ = _266_v73
		var _nwElement0_7 D1 = _254_v65
		_ = _nwElement0_7
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(2))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_7, 0)
		(_nw33).ArraySet1(_254_v65, 1)
		_266_v73 = _nw33
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_265_v72), 0))
		_ = _index27
		(_265_v72).ArraySet1(_dafny.MultiSetOf(_266_v73, _266_v73, _266_v73, _266_v73, _266_v73), (_index27).Int())
		var _267_v74 _dafny.MultiSet
		_ = _267_v74
		_267_v74 = _dafny.MultiSetOf(_266_v73, _266_v73)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_265_v72), 0))
		_ = _index28
		(_265_v72).ArraySet1(((_267_v74).Difference(_dafny.MultiSetOf(_266_v73, _266_v73))).Intersection((_267_v74).Intersection(_267_v74)), (_index28).Int())
		var _268_v75 _dafny.Array
		_ = _268_v75
		var _nw34 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(25))
		_ = _nw34
		_268_v75 = _nw34
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_268_v75), 0))
		_ = _index29
		(_268_v75).ArraySet1(_252_v63, (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(299), _dafny.ArrayLen((_268_v75), 0))
		_ = _index30
		(_268_v75).ArraySet1(_252_v63, (_index30).Int())
		var _269_v76 _dafny.Array
		_ = _269_v76
		var _nw35 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
		_ = _nw35
		_269_v76 = _nw35
		_269_v76 = _269_v76
		var _270_v77 D1
		_ = _270_v77
		_270_v77 = Companion_D1_.Create_DC3_(_165_v1)
		var _271_v78 _dafny.Array
		_ = _271_v78
		var _nwElement0_8 D1 = Companion_D1_.Create_DC3_(_165_v1)
		_ = _nwElement0_8
		var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(16))
		_ = _nw36
		(_nw36).ArraySet1(_nwElement0_8, 0)
		(_nw36).ArraySet1(_270_v77, 1)
		(_nw36).ArraySet1(_270_v77, 2)
		(_nw36).ArraySet1(_270_v77, 3)
		(_nw36).ArraySet1(_270_v77, 4)
		(_nw36).ArraySet1(_270_v77, 5)
		(_nw36).ArraySet1(Companion_D1_.Create_DC3_(_165_v1), 6)
		(_nw36).ArraySet1(Companion_D1_.Create_DC3_(_165_v1), 7)
		(_nw36).ArraySet1(_270_v77, 8)
		(_nw36).ArraySet1(_270_v77, 9)
		(_nw36).ArraySet1((func() D1 {
			if true {
				return _270_v77
			}
			return _270_v77
		})(), 10)
		(_nw36).ArraySet1(_270_v77, 11)
		(_nw36).ArraySet1(_270_v77, 12)
		(_nw36).ArraySet1(Companion_D1_.Create_DC3_(Companion_Default___.Fm4((_252_v63).F18(), (_259_v69).F18(), (_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int), (_252_v63).F18(), _175_globalState)), 13)
		(_nw36).ArraySet1(Companion_D1_.Create_DC3_(_165_v1), 14)
		(_nw36).ArraySet1(_270_v77, 15)
		_271_v78 = _nw36
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_271_v78), 0))
		_ = _index31
		(_271_v78).ArraySet1(_270_v77, (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_271_v78), 0))
		_ = _index32
		(_271_v78).ArraySet1(_270_v77, (_index32).Int())
	} else {
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(826), _dafny.ArrayLen((_171_v7), 0))
		_ = _index33
		(_171_v7).ArraySet1(_164_v0, (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(826), _dafny.ArrayLen((_171_v7), 0))
		_ = _index34
		(_171_v7).ArraySet1(_164_v0, (_index34).Int())
		var _272_v79 D0
		_ = _272_v79
		_272_v79 = Companion_D0_.Create_DC0_((_252_v63).F18(), ((_178_v11).Update(_174_v8, (_dafny.MultiSetFromSeq(_258_v68)).Cardinality())).Update(_174_v8, (_252_v63).F18()), _168_v4)
		var _273_v80 D0
		_ = _273_v80
		_273_v80 = Companion_D0_.Create_DC1_(_272_v79)
		var _274_v81 D0
		_ = _274_v81
		_274_v81 = Companion_D0_.Create_DC1_(_272_v79)
		var _275_v82 _dafny.Int
		_ = _275_v82
		var _276_v83 _dafny.Sequence
		_ = _276_v83
		var _277_v84 _dafny.Set
		_ = _277_v84
		var _278_v85 bool
		_ = _278_v85
		var _out28 _dafny.Int
		_ = _out28
		var _out29 _dafny.Sequence
		_ = _out29
		var _out30 _dafny.Set
		_ = _out30
		var _out31 bool
		_ = _out31
		_out28, _out29, _out30, _out31 = Companion_Default___.M0(_274_v81, (func() _dafny.Int {
			if _164_v0 {
				return (_252_v63).F18()
			}
			return _dafny.IntOfInt64(-94)
		})(), _175_globalState)
		_275_v82 = _out28
		_276_v83 = _out29
		_277_v84 = _out30
		_278_v85 = _out31
		var _279_v86 *C0
		_ = _279_v86
		var _nw37 *C0 = New_C0_()
		_ = _nw37
		_nw37.Ctor__(_167_v3)
		_279_v86 = _nw37
		var _280_v87 _dafny.Map
		_ = _280_v87
		_280_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_171_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(826), _dafny.ArrayLen((_171_v7), 0))).Int()).(bool), _274_v81)
		_280_v87 = (_280_v87).Update(_278_v85, Companion_D0_.Create_DC1_(_272_v79))
		(_175_globalState).F13 = _164_v0
	}
	_176_v9 = (_176_v9).Update((_252_v63).F18(), (_252_v63).F18())
	var _281_v88 _dafny.CodePoint
	_ = _281_v88
	_281_v88 = _dafny.CodePoint('j')
	var _282_v89 _dafny.Map
	_ = _282_v89
	_282_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_259_v69).F18(), (_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int), _167_v3, _167_v3, _dafny.IntOfInt64(-242)), _281_v88)
	var _283_v90 _dafny.Map
	_ = _283_v90
	_283_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v0, _282_v89)
	_282_v89 = (func() _dafny.Map {
		if (_283_v90).Contains(true) {
			return (_283_v90).Get(true).(_dafny.Map)
		}
		return _282_v89
	})()
	(_175_globalState).F4 = (_dafny.Zero).Minus(_167_v3)
	_164_v0 = !(_164_v0)
	var _284_v91 _dafny.Array
	_ = _284_v91
	var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
	_ = _nw38
	_284_v91 = _nw38
	var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_284_v91), 0))
	_ = _index35
	(_284_v91).ArraySet1(_170_v6, (_index35).Int())
	var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_284_v91), 0))
	_ = _index36
	(_284_v91).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_170_v6, _170_v6), (_index36).Int())
	_171_v7 = _171_v7
	_164_v0 = _164_v0
	if ((_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_174_v8).Cardinality())) == 0 {
		var _285_v92 _dafny.Sequence
		_ = _285_v92
		_285_v92 = _dafny.SeqOf(_259_v69, _259_v69, _252_v63, _259_v69)
		var _286_v93 _dafny.Map
		_ = _286_v93
		_286_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v88, _dafny.Companion_Sequence_.IsProperPrefixOf(_285_v92, _dafny.SeqOf(_252_v63, _252_v63)))
		_286_v93 = (_286_v93).Update(_281_v88, _164_v0)
		if _164_v0 {
			var _287_v94 _dafny.Map
			_ = _287_v94
			_287_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(54), _179_v12)
			var _288_v96 _dafny.Sequence
			_ = _288_v96
			_288_v96 = _dafny.SeqOf(_257_v67, _257_v67, func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_287_v94).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _289_v95 _dafny.Int
					_289_v95 = interface{}(_compr_6).(_dafny.Int)
					if (_287_v94).Contains(_289_v95) {
						_coll6.Add((_289_v95).Plus(_dafny.IntOfInt64(514)))
					}
				}
				return _coll6.ToSet()
			}())
			_288_v96 = _dafny.Companion_Sequence_.Concatenate(_288_v96, _288_v96)
			var _290_v97 D1
			_ = _290_v97
			_290_v97 = Companion_D1_.Create_DC3_(_165_v1)
			var _291_v98 _dafny.Map
			_ = _291_v98
			_291_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_165_v1, _174_v8)
			var _pat_let_tv9 = _165_v1
			_ = _pat_let_tv9
			_164_v0 = !(_291_v98).Contains((func(_pat_let5_0 D1) D1 {
				return func(_292_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let6_0 _dafny.Set) D1 {
						return func(_293_dt__update_hcf5_h0 _dafny.Set) D1 {
							return Companion_D1_.Create_DC3_(_293_dt__update_hcf5_h0)
						}(_pat_let6_0)
					}(_pat_let_tv9)
				}(_pat_let5_0)
			}(_290_v97)).Dtor_cf5())
			var _294_v99 D0
			_ = _294_v99
			_294_v99 = Companion_D0_.Create_DC0_((_259_v69).F18(), _178_v11, _168_v4)
			var _295_v100 D0
			_ = _295_v100
			_295_v100 = Companion_D0_.Create_DC1_(_294_v99)
			var _296_v101 _dafny.Int
			_ = _296_v101
			var _297_v102 _dafny.Sequence
			_ = _297_v102
			var _298_v103 _dafny.Set
			_ = _298_v103
			var _299_v104 bool
			_ = _299_v104
			var _out32 _dafny.Int
			_ = _out32
			var _out33 _dafny.Sequence
			_ = _out33
			var _out34 _dafny.Set
			_ = _out34
			var _out35 bool
			_ = _out35
			_out32, _out33, _out34, _out35 = Companion_Default___.M0(_295_v100, Companion_Default___.Fm0(_179_v12, _175_globalState), _175_globalState)
			_296_v101 = _out32
			_297_v102 = _out33
			_298_v103 = _out34
			_299_v104 = _out35
			var _300_v105 _dafny.Map
			_ = _300_v105
			_300_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_296_v101, _285_v92)
			var _301_v106 _dafny.Map
			_ = _301_v106
			_301_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v0, _285_v92)
			var _302_v107 _dafny.MultiSet
			_ = _302_v107
			_302_v107 = _dafny.MultiSetOf((func() _dafny.Sequence {
				if (_300_v105).Contains(_167_v3) {
					return (_300_v105).Get(_167_v3).(_dafny.Sequence)
				}
				return (func() _dafny.Sequence {
					if (_301_v106).Contains(_164_v0) {
						return (_301_v106).Get(_164_v0).(_dafny.Sequence)
					}
					return _285_v92
				})()
			})(), _285_v92)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
			_ = _index37
			(_255_v66).ArraySet1((((_302_v107).Intersection(_302_v107)).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tprjqefl")).Cardinality())), (_index37).Int())
			_285_v92 = _285_v92
		} else {
			(_175_globalState).F13 = (_dafny.IntOfUint32((_285_v92).Cardinality())).Cmp((_259_v69).F18()) < 0
			var _303_v108 _dafny.Sequence
			_ = _303_v108
			_303_v108 = _dafny.SeqOf(_174_v8, _174_v8)
			var _304_v109 D1
			_ = _304_v109
			_304_v109 = Companion_D1_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("cs"), _164_v0, _303_v108, Companion_Default___.Fm5((_168_v4).Cardinality(), ((_284_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_284_v91), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_179_v12, _175_globalState), _dafny.IntOfUint32(((_284_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_284_v91), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int), (_259_v69).F18(), (_252_v63).F18(), _175_globalState), _281_v88)
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
			_ = _index38
			var _rhs15 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_174_v8, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(650))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_305_v88 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_306_i4 _dafny.Int) _dafny.CodePoint {
					return _305_v88
				}
			})(_281_v88))), _174_v8))).Cardinality())
			_ = _rhs15
			var _rhs16 *C0 = _259_v69
			_ = _rhs16
			var _rhs17 _dafny.Map = _260_v70
			_ = _rhs17
			var _rhs18 bool = (_304_v109).Dtor_cf7()
			_ = _rhs18
			var _rhs19 _dafny.Int = (_252_v63).F18()
			_ = _rhs19
			var _lhs14 *GlobalState = _175_globalState
			_ = _lhs14
			var _lhs15 *GlobalState = _175_globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _255_v66
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
			_ = _lhs17
			_lhs14.F15 = _rhs15
			_252_v63 = _rhs16
			_260_v70 = _rhs17
			_lhs15.F13 = _rhs18
			(_lhs16).ArraySet1(_rhs19, (_lhs17).Int())
			var _307_v110 D2
			_ = _307_v110
			_307_v110 = Companion_D2_.Create_DC6_(_259_v69)
			var _308_v111 _dafny.Array
			_ = _308_v111
			var _nwElement0_9 *C0 = _259_v69
			_ = _nwElement0_9
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(29))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_9, 0)
			(_nw39).ArraySet1((func() *C0 {
				if !(_164_v0) {
					return _259_v69
				}
				return (_307_v110).Dtor_cf12()
			})(), 1)
			(_nw39).ArraySet1(_259_v69, 2)
			(_nw39).ArraySet1(_252_v63, 3)
			(_nw39).ArraySet1(_252_v63, 4)
			(_nw39).ArraySet1(_252_v63, 5)
			(_nw39).ArraySet1(_252_v63, 6)
			(_nw39).ArraySet1(_252_v63, 7)
			(_nw39).ArraySet1(_252_v63, 8)
			(_nw39).ArraySet1(_252_v63, 9)
			(_nw39).ArraySet1(_259_v69, 10)
			(_nw39).ArraySet1((_285_v92).Select((Companion_Default___.SafeIndex((_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_285_v92).Cardinality()))).Uint32()).(*C0), 11)
			(_nw39).ArraySet1(_252_v63, 12)
			(_nw39).ArraySet1(_252_v63, 13)
			(_nw39).ArraySet1(_259_v69, 14)
			(_nw39).ArraySet1(_259_v69, 15)
			(_nw39).ArraySet1(_252_v63, 16)
			(_nw39).ArraySet1(_252_v63, 17)
			(_nw39).ArraySet1((_285_v92).Select((Companion_Default___.SafeIndex((_259_v69).F18(), _dafny.IntOfUint32((_285_v92).Cardinality()))).Uint32()).(*C0), 18)
			(_nw39).ArraySet1(_252_v63, 19)
			(_nw39).ArraySet1((_285_v92).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.IntOfUint32((_285_v92).Cardinality()))).Uint32()).(*C0), 20)
			(_nw39).ArraySet1(_259_v69, 21)
			(_nw39).ArraySet1(_259_v69, 22)
			(_nw39).ArraySet1(_252_v63, 23)
			(_nw39).ArraySet1(_259_v69, 24)
			(_nw39).ArraySet1(_252_v63, 25)
			(_nw39).ArraySet1(_252_v63, 26)
			(_nw39).ArraySet1(_259_v69, 27)
			(_nw39).ArraySet1(_259_v69, 28)
			_308_v111 = _nw39
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_308_v111), 0))
			_ = _index39
			(_308_v111).ArraySet1(_259_v69, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_308_v111), 0))
			_ = _index40
			(_308_v111).ArraySet1(_259_v69, (_index40).Int())
			_179_v12 = (func() D0 {
				if _164_v0 {
					return _179_v12
				}
				return _179_v12
			})()
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_171_v7), 0))
			_ = _index41
			(_171_v7).ArraySet1(!(_164_v0) || (_164_v0), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_171_v7), 0))
			_ = _index42
			(_171_v7).ArraySet1(_164_v0, (_index42).Int())
		}
		var _309_v112 D0
		_ = _309_v112
		_309_v112 = Companion_D0_.Create_DC0_((_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int), (_178_v11).Update(_174_v8, (_252_v63).F18()), _dafny.MultiSetOf((_dafny.Zero).Minus(_167_v3), (_259_v69).F18()))
		var _310_v113 D0
		_ = _310_v113
		_310_v113 = Companion_D0_.Create_DC1_(_309_v112)
		var _311_v114 _dafny.Int
		_ = _311_v114
		var _312_v115 _dafny.Sequence
		_ = _312_v115
		var _313_v116 _dafny.Set
		_ = _313_v116
		var _314_v117 bool
		_ = _314_v117
		var _out36 _dafny.Int
		_ = _out36
		var _out37 _dafny.Sequence
		_ = _out37
		var _out38 _dafny.Set
		_ = _out38
		var _out39 bool
		_ = _out39
		_out36, _out37, _out38, _out39 = Companion_Default___.M0(Companion_D0_.Create_DC1_(_310_v113), ((_259_v69).F18()).Minus((_259_v69).F18()), _175_globalState)
		_311_v114 = _out36
		_312_v115 = _out37
		_313_v116 = _out38
		_314_v117 = _out39
		_164_v0 = _164_v0
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_171_v7), 0))
		_ = _index43
		(_171_v7).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("eqmqr"), _174_v8), (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_171_v7), 0))
		_ = _index44
		(_171_v7).ArraySet1(_314_v117, (_index44).Int())
	} else {
		var _315_v118 *C0
		_ = _315_v118
		var _nw40 *C0 = New_C0_()
		_ = _nw40
		_nw40.Ctor__((_252_v63).F18())
		_315_v118 = _nw40
		(_175_globalState).F4 = (_252_v63).F18()
		var _rhs20 _dafny.Int = (_252_v63).F18()
		_ = _rhs20
		var _rhs21 *C0 = _259_v69
		_ = _rhs21
		var _rhs22 bool = !(Companion_Default___.Fm1(_164_v0, _175_globalState))
		_ = _rhs22
		var _lhs18 *GlobalState = _175_globalState
		_ = _lhs18
		var _lhs19 *GlobalState = _175_globalState
		_ = _lhs19
		_lhs18.F4 = _rhs20
		_259_v69 = _rhs21
		_lhs19.F13 = _rhs22
		var _316_v119 _dafny.Map
		_ = _316_v119
		_316_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_281_v88, Companion_Default___.Fm1(_164_v0, _175_globalState))
		var _rhs23 bool = !(_164_v0)
		_ = _rhs23
		var _rhs24 bool = _164_v0
		_ = _rhs24
		var _rhs25 _dafny.Map = _316_v119
		_ = _rhs25
		var _lhs20 *GlobalState = _175_globalState
		_ = _lhs20
		_164_v0 = _rhs23
		_lhs20.F13 = _rhs24
		_316_v119 = _rhs25
		var _317_v120 _dafny.Map
		_ = _317_v120
		_317_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_259_v69).F18(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-148))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_318_v88 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_319_i5 _dafny.Int) _dafny.CodePoint {
				return _318_v88
			}
		})(_281_v88))))
		var _320_v121 D3
		_ = _320_v121
		_320_v121 = Companion_D3_.Create_DC9_(_317_v120)
		var _321_v123 _dafny.Array
		_ = _321_v123
		var _nwElement0_10 _dafny.Map = _317_v120
		_ = _nwElement0_10
		var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(7))
		_ = _nw41
		(_nw41).ArraySet1(_nwElement0_10, 0)
		(_nw41).ArraySet1((func() _dafny.Map {
			if _164_v0 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(696), _174_v8)
			}
			return _317_v120
		})(), 1)
		(_nw41).ArraySet1((_320_v121).Dtor_cf15(), 2)
		(_nw41).ArraySet1(func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(905), _dafny.IntOfInt64(714))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _322_v122 _dafny.Int
				_322_v122 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(905)).Cmp(_322_v122) <= 0) && ((_322_v122).Cmp(_dafny.IntOfInt64(714)) < 0) {
					_coll7.Add(Companion_Default___.SafeDivisionInt(_322_v122, _dafny.IntOfUint32((_258_v68).Cardinality())), _174_v8)
				}
			}
			return _coll7.ToMap()
		}(), 3)
		(_nw41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int), _174_v8), 4)
		(_nw41).ArraySet1(_317_v120, 5)
		(_nw41).ArraySet1(_317_v120, 6)
		_321_v123 = _nw41
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_321_v123), 0))
		_ = _index45
		(_321_v123).ArraySet1(_317_v120, (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_321_v123), 0))
		_ = _index46
		(_321_v123).ArraySet1((func() _dafny.Map {
			if _164_v0 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_259_v69).F18(), _174_v8)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_174_v8).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-266))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_323_v88 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_324_i6 _dafny.Int) _dafny.CodePoint {
					return _323_v88
				}
			})(_281_v88))))
		})(), (_index46).Int())
	}
	_281_v88 = Companion_Default___.Fm6((_257_v67).IsProperSubsetOf(_257_v67), _175_globalState)
	_177_v10 = (_177_v10).Update((_164_v0) == (!(_164_v0)), (_252_v63).F18())
	if _164_v0 {
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
		_ = _index47
		(_255_v66).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if _164_v0 {
				return (_259_v69).F18()
			}
			return _dafny.IntOfInt64(706)
		})(), (_252_v63).F18()), (_index47).Int())
		var _rhs26 _dafny.Int = (_252_v63).F18()
		_ = _rhs26
		var _rhs27 _dafny.Int = Companion_Default___.Fm0(_179_v12, _175_globalState)
		_ = _rhs27
		var _lhs21 *GlobalState = _175_globalState
		_ = _lhs21
		var _lhs22 *GlobalState = _175_globalState
		_ = _lhs22
		_lhs21.F4 = _rhs26
		_lhs22.F4 = _rhs27
		var _325_v124 *C0
		_ = _325_v124
		var _nw42 *C0 = New_C0_()
		_ = _nw42
		_nw42.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_326_v68 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_327_i7 _dafny.Int) _dafny.Sequence {
				return _326_v68
			}
		})(_258_v68))), (Companion_Default___.SafeIndex((_255_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_328_v68 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_329_i7 _dafny.Int) _dafny.Sequence {
				return _328_v68
			}
		})(_258_v68)))).Cardinality()))).Uint32(), _258_v68)).Cardinality()))
		_325_v124 = _nw42
		_178_v11 = (_178_v11).Update(_174_v8, _dafny.IntOfInt64(759))
		(_175_globalState).F4 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_259_v69).F18(), (_170_v6).Select((Companion_Default___.SafeIndex((_325_v124).F18(), _dafny.IntOfUint32((_170_v6).Cardinality()))).Uint32()).(_dafny.Int)), Companion_Default___.Fm0(_179_v12, _175_globalState))
	} else {
		var _330_v125 _dafny.Map
		_ = _330_v125
		_330_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm7(_dafny.IntOfUint32((_dafny.SeqOf((func() bool {
			if (_260_v70).Contains((_259_v69).F18()) {
				return (_260_v70).Get((_259_v69).F18()).(bool)
			}
			return _164_v0
		})(), !(_164_v0), _164_v0)).Cardinality()), _258_v68, _164_v0, _175_globalState)).Cardinality(), _174_v8)
		var _331_v126 _dafny.Array
		_ = _331_v126
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(19))
		_ = _nw43
		_331_v126 = _nw43
		var _332_v127 _dafny.Map
		_ = _332_v127
		_332_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v125, _331_v126)
		_332_v127 = (_332_v127).Update(_330_v125, _331_v126)
		var _333_v128 *C0
		_ = _333_v128
		var _nw44 *C0 = New_C0_()
		_ = _nw44
		_nw44.Ctor__((_dafny.IntOfInt64(868)).Plus(_dafny.IntOfInt64(823)))
		_333_v128 = _nw44
		(_175_globalState).F13 = (func() bool {
			if true {
				return _164_v0
			}
			return !(_164_v0)
		})()
		var _334_v129 _dafny.Map
		_ = _334_v129
		_334_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v0, _255_v66)
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(580), _dafny.ArrayLen((_255_v66), 0))
		_ = _index48
		(_255_v66).ArraySet1((((_dafny.Zero).Minus(_167_v3)).Plus((_252_v63).F18())).Minus((_334_v129).Cardinality()), (_index48).Int())
		(_175_globalState).F0 = (_167_v3).Plus((_dafny.IntOfInt64(326)).Plus((_168_v4).Cardinality()))
	}
	_dafny.Print(_164_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v1).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_167_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v4).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(-408))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v5).Equals(_dafny.MultiSetOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_170_v6, _dafny.SeqOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v7).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_174_v8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F1()).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_175_globalState).F8(), _dafny.SeqOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(639), _dafny.IntOfInt64(2), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F9()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(-408))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_175_globalState).F16()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_175_globalState).F17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(639), _dafny.IntOfInt64(639)).UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_177_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(639)).UpdateUnsafe(false, _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mgx"), _dafny.IntOfInt64(759))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_179_v12).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_179_v12).Dtor_cf1()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mgx"), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_179_v12).Dtor_cf2()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(-408))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_251_v62, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(2)), _dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(-408)), _dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(-408)), _dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(-408)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_252_v63).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v64).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_254_v65).Dtor_cf4()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(639))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v66).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v67).Equals(_dafny.SetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_258_v68, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v69).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v70).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_281_v88)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v89).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(639), _dafny.IntOfInt64(639), _dafny.Zero, _dafny.IntOfInt64(-242)), _dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_283_v90).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(639), _dafny.IntOfInt64(639), _dafny.IntOfInt64(639), _dafny.Zero, _dafny.IntOfInt64(-242)), _dafny.CodePoint('j')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_284_v91).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(2), _dafny.IntOfInt64(2), _dafny.IntOfInt64(2), _dafny.IntOfInt64(2))))
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
	Cf1 _dafny.Map
	Cf2 _dafny.MultiSet
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 _dafny.Map, Cf2 _dafny.MultiSet) D0 {
	return D0{D0_DC0{Cf0, Cf1, Cf2}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf3 D0
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf3 D0) D0 {
	return D0{D0_DC1{Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, _dafny.EmptyMap, _dafny.EmptyMultiSet)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.MultiSet {
	return _this.Get_().(D0_DC0).Cf2
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1.Equals(data2.Cf1) && data1.Cf2.Equals(data2.Cf2)
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
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

type D1_DC3 struct {
	Cf5 _dafny.Set
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 _dafny.Set) D1 {
	return D1{D1_DC3{Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf6  _dafny.Sequence
	Cf7  bool
	Cf8  _dafny.Sequence
	Cf9  _dafny.Sequence
	Cf10 _dafny.CodePoint
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf6 _dafny.Sequence, Cf7 bool, Cf8 _dafny.Sequence, Cf9 _dafny.Sequence, Cf10 _dafny.CodePoint) D1 {
	return D1{D1_DC4{Cf6, Cf7, Cf8, Cf9, Cf10}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf4 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Map) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC5 struct {
	Cf11 D1
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf11 D1) D1 {
	return D1{D1_DC5{Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(_dafny.EmptySet)
}

func (_this D1) Dtor_cf5() _dafny.Set {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf4() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf11() D1 {
	return _this.Get_().(D1_DC5).Cf11
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
			return "D1.DC4" + "(" + data.Cf6.VerbatimString(true) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + data.Cf9.VerbatimString(true) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf11) + ")"
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
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf6.Equals(data2.Cf6) && data1.Cf7 == data2.Cf7 && data1.Cf8.Equals(data2.Cf8) && data1.Cf9.Equals(data2.Cf9) && data1.Cf10 == data2.Cf10
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
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

type D2_DC7 struct {
	Cf13 _dafny.Array
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 _dafny.Array) D2 {
	return D2{D2_DC7{Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf12 *C0
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf12 *C0) D2 {
	return D2{D2_DC6{Cf12}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC8 struct {
	Cf14 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf14 D2) D2 {
	return D2{D2_DC8{Cf14}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D2) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf12() *C0 {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf14() D2 {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf13 == data2.Cf13
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf12 == data2.Cf12
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf14.Equals(data2.Cf14)
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
	Cf16 bool
	Cf17 _dafny.Array
	Cf18 _dafny.Sequence
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 bool, Cf17 _dafny.Array, Cf18 _dafny.Sequence) D3 {
	return D3{D3_DC10{Cf16, Cf17, Cf18}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf15 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 _dafny.Map) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq)
}

func (_this D3) Dtor_cf16() bool {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf18
}

func (_this D3) Dtor_cf15() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + data.Cf18.VerbatimString(true) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17 == data2.Cf17 && data1.Cf18.Equals(data2.Cf18)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf15.Equals(data2.Cf15)
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

type D4_DC12 struct {
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_() D4 {
	return D4{D4_DC12{}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC11 struct {
	Cf19 _dafny.Map
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf19 _dafny.Map) D4 {
	return D4{D4_DC11{Cf19}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_()
}

func (_this D4) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D4_DC11).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC12:
		{
			_, ok := other.Get_().(D4_DC12)
			return ok
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf19.Equals(data2.Cf19)
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
	F0   _dafny.Int
	F4   _dafny.Int
	F10  _dafny.Int
	F13  bool
	F15  _dafny.Int
	_f1  _dafny.Set
	_f2  bool
	_f3  bool
	_f5  _dafny.Int
	_f6  bool
	_f7  _dafny.Array
	_f8  _dafny.Sequence
	_f9  _dafny.MultiSet
	_f11 _dafny.Int
	_f12 _dafny.CodePoint
	_f14 _dafny.Int
	_f16 _dafny.Array
	_f17 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F4 = _dafny.Zero
	_this.F10 = _dafny.Zero
	_this.F13 = false
	_this.F15 = _dafny.Zero
	_this._f1 = _dafny.EmptySet
	_this._f2 = false
	_this._f3 = false
	_this._f5 = _dafny.Zero
	_this._f6 = false
	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptySeq
	_this._f9 = _dafny.EmptyMultiSet
	_this._f11 = _dafny.Zero
	_this._f12 = _dafny.CodePoint('D')
	_this._f14 = _dafny.Zero
	_this._f16 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f17 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Set, f2 bool, f3 bool, f4 _dafny.Int, f5 _dafny.Int, f6 bool, f7 _dafny.Array, f8 _dafny.Sequence, f9 _dafny.MultiSet, f10 _dafny.Int, f11 _dafny.Int, f12 _dafny.CodePoint, f13 bool, f14 _dafny.Int, f15 _dafny.Int, f16 _dafny.Array, f17 _dafny.Sequence) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *GlobalState) F1() _dafny.Set {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() _dafny.Int {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Array {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.MultiSet {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.CodePoint {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F16() _dafny.Array {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f18 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f18 = _dafny.Zero
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

func (_this *C0) Ctor__(f18 _dafny.Int) {
	{
		(_this)._f18 = f18
	}
}
func (_this *C0) F18() _dafny.Int {
	{
		return _this._f18
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
