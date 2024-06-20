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
func (_static *CompanionStruct_Default___) Fm0(globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfUint32(((func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(62))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))
		}
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(584))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _dafny.UnicodeSeqOfUtf8Bytes("jkumqri"))
	})()).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	return !((((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-940)))).Plus(_dafny.IntOfInt64(910))).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rrv"), _dafny.UnicodeSeqOfUtf8Bytes("ffpcdlj"))).Cardinality())) == 0)
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-974))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(561)
	})))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(224), _dafny.IntOfInt64(957))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _3_v0 _dafny.Int
			_3_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(224)).Cmp(_3_v0) <= 0) && ((_3_v0).Cmp(_dafny.IntOfInt64(957)) < 0) {
				_coll0.Add(Companion_Default___.SafeModuloInt(_3_v0, _dafny.IntOfInt64(-930)))
			}
		}
		return _coll0.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	if (!(false)) || (true) {
		return _dafny.CodePoint('a')
	} else {
		return _dafny.CodePoint('s')
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true), _dafny.SeqOf(false)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((Companion_D2_.Create_DC7_(true, _dafny.IntOfInt64(-582), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(false), true), _dafny.IntOfInt64(370)))).Dtor_cf6(), false), _dafny.SeqOf(true, true)))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 D2, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(767))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	}))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	var _source0 D2 = Companion_D2_.Create_DC7_(false, _dafny.IntOfInt64(726), true, (Companion_D2_.Create_DC7_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(153), func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(101), _dafny.IntOfInt64(603))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _5_v0 _dafny.Int
				_5_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(101)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(603)) < 0) {
					_coll2.Add(Companion_Default___.SafeModuloInt(_5_v0, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality()), (Companion_D1_.Create_DC2_(false)).Dtor_cf2())
				}
			}
			return _coll2.ToMap()
		}()).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _6_v1 _dafny.Int
			_6_v1 = interface{}(_compr_1).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(101), _dafny.IntOfInt64(603))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _5_v0 _dafny.Int
					_5_v0 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(101)).Cmp(_5_v0) <= 0) && ((_5_v0).Cmp(_dafny.IntOfInt64(603)) < 0) {
						_coll3.Add(Companion_Default___.SafeModuloInt(_5_v0, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality()), (Companion_D1_.Create_DC2_(false)).Dtor_cf2())
					}
				}
				return _coll3.ToMap()
			}()).Contains(_6_v1) {
				_coll1.Add(Companion_Default___.SafeDivisionInt(_6_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-790))).Cardinality()), _dafny.One)).Cardinality()))).Cardinality())))
			}
		}
		return _coll1.ToSet()
	}())).Cardinality(), false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(false)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gcxukgk")).Cardinality())))).Dtor_cf9())
	_ = _source0
	if _source0.Is_DC5() {
		var _7___mcc_h0 _dafny.Int = _source0.Get_().(D2_DC5).Cf4
		_ = _7___mcc_h0
		var _8___mcc_h1 _dafny.Int = _source0.Get_().(D2_DC5).Cf5
		_ = _8___mcc_h1
		var _9_cf5 _dafny.Int = _8___mcc_h1
		_ = _9_cf5
		var _10_cf4 _dafny.Int = _7___mcc_h0
		_ = _10_cf4
		return _dafny.SeqOf(_dafny.IntOfInt64(859))
	} else if _source0.Is_DC6() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(989))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_11_i0 _dafny.Int) _dafny.Int {
			return (func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf(true))).Elements()); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _12_v2 _dafny.MultiSet
					_12_v2 = interface{}(_compr_4).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf(true)), _12_v2) {
						_coll4.Add(_12_v2)
					}
				}
				return _coll4.ToSet()
			}()).Cardinality()
		})), _dafny.SeqOf(_dafny.IntOfInt64(71)))
	} else if _source0.Is_DC7() {
		var _13___mcc_h2 bool = _source0.Get_().(D2_DC7).Cf6
		_ = _13___mcc_h2
		var _14___mcc_h3 _dafny.Int = _source0.Get_().(D2_DC7).Cf7
		_ = _14___mcc_h3
		var _15___mcc_h4 bool = _source0.Get_().(D2_DC7).Cf8
		_ = _15___mcc_h4
		var _16___mcc_h5 _dafny.Map = _source0.Get_().(D2_DC7).Cf9
		_ = _16___mcc_h5
		var _17_cf9 _dafny.Map = _16___mcc_h5
		_ = _17_cf9
		var _18_cf8 bool = _15___mcc_h4
		_ = _18_cf8
		var _19_cf7 _dafny.Int = _14___mcc_h3
		_ = _19_cf7
		var _20_cf6 bool = _13___mcc_h2
		_ = _20_cf6
		return _dafny.SeqOf(_19_cf7)
	} else {
		var _21___mcc_h6 _dafny.Int = _source0.Get_().(D2_DC4).Cf3
		_ = _21___mcc_h6
		var _22_cf3 _dafny.Int = _21___mcc_h6
		_ = _22_cf3
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(!(false)))).Cardinality(), _22_cf3, _22_cf3, _22_cf3)
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, true, !(true))).Difference((_dafny.SetOf(false, true)).Difference(_dafny.SetOf(!(false))))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Union((_dafny.MultiSetOf(false, false)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (_dafny.Array, _dafny.Int) {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	r1 = Companion_Default___.Fm0(globalState)
	var _23_v0 _dafny.Int
	_ = _23_v0
	_23_v0 = _dafny.IntOfInt64(727)
	var _24_v1 _dafny.Set
	_ = _24_v1
	_24_v1 = _dafny.SetOf(_23_v0)
	var _25_v2 bool
	_ = _25_v2
	_25_v2 = false
	var _26_v3 _dafny.MultiSet
	_ = _26_v3
	_26_v3 = _dafny.MultiSetOf(_25_v2)
	var _27_v4 _dafny.Sequence
	_ = _27_v4
	_27_v4 = _dafny.SeqOf(_26_v3, _26_v3)
	var _28_v5 _dafny.CodePoint
	_ = _28_v5
	_28_v5 = _dafny.CodePoint('k')
	var _29_v6 _dafny.Map
	_ = _29_v6
	_29_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_27_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_28_v5)).Cardinality()), _dafny.IntOfUint32((_27_v4).Cardinality()))).Uint32()).(_dafny.MultiSet), (_26_v3).Cardinality())
	var _30_v7 D2
	_ = _30_v7
	_30_v7 = Companion_D2_.Create_DC7_((_24_v1).IsProperSubsetOf(_24_v1), _23_v0, _25_v2, _29_v6)
	var _source1 D2 = _30_v7
	_ = _source1
	if _source1.Is_DC5() {
		var _31___mcc_h0 _dafny.Int = _source1.Get_().(D2_DC5).Cf4
		_ = _31___mcc_h0
		var _32___mcc_h1 _dafny.Int = _source1.Get_().(D2_DC5).Cf5
		_ = _32___mcc_h1
		var _33_cf5 _dafny.Int = _32___mcc_h1
		_ = _33_cf5
		var _34_cf4 _dafny.Int = _31___mcc_h0
		_ = _34_cf4
		var _35_v8 _dafny.Array
		_ = _35_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Set = func(_36_i0 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("sldfemqf"), _dafny.UnicodeSeqOfUtf8Bytes("e"))
			}
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw0).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw0).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_35_v8 = _nw0
		var _37_v9 _dafny.Set
		_ = _37_v9
		_37_v9 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("adcvwjlwm"))
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_35_v8), 0))
		_ = _index0
		(_35_v8).ArraySet1((_37_v9).Union(_37_v9), (_index0).Int())
		var _38_v10 _dafny.Sequence
		_ = _38_v10
		_38_v10 = _dafny.UnicodeSeqOfUtf8Bytes("rbff")
		var _39_v11 _dafny.Sequence
		_ = _39_v11
		_39_v11 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("oktgtfhvg"), _38_v10)
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_35_v8), 0))
		_ = _index1
		(_35_v8).ArraySet1(((_37_v9).Union(func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_39_v11).Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _40_v12 _dafny.Sequence
				_40_v12 = interface{}(_compr_5).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_39_v11, _40_v12) {
					_coll5.Add(_40_v12)
				}
			}
			return _coll5.ToSet()
		}())).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(589))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}((func(_41_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_42_i1 _dafny.Int) _dafny.CodePoint {
				return _41_v5
			}
		})(_28_v5))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}((func(_43_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_44_i2 _dafny.Int) _dafny.CodePoint {
				return _43_v5
			}
		})(_28_v5))))), (_index1).Int())
		var _45_v13 _dafny.Map
		_ = _45_v13
		_45_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
			if _25_v2 {
				return _28_v5
			}
			return _28_v5
		})(), !(_25_v2))
		(globalState).F1 = (func() bool {
			if (_45_v13).Contains(Companion_Default___.Fm4(_25_v2, globalState)) {
				return (_45_v13).Get(Companion_Default___.Fm4(_25_v2, globalState)).(bool)
			}
			return _25_v2
		})()
		var _46_v14 _dafny.Map
		_ = _46_v14
		_46_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_cf5, true)
		var _47_v15 _dafny.Map
		_ = _47_v15
		_47_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_33_cf5), (func() bool {
			if (_46_v14).Contains(_dafny.IntOfUint32((Companion_Default___.Fm5(_33_cf5, _34_cf4, globalState)).Cardinality())) {
				return (_46_v14).Get(_dafny.IntOfUint32((Companion_Default___.Fm5(_33_cf5, _34_cf4, globalState)).Cardinality())).(bool)
			}
			return _25_v2
		})())
		var _48_v16 _dafny.Sequence
		_ = _48_v16
		_48_v16 = _dafny.SeqOf(true, _25_v2, _25_v2)
		var _49_v17 _dafny.Sequence
		_ = _49_v17
		_49_v17 = _dafny.SeqOf(_48_v16, _48_v16, _dafny.SeqOf(_25_v2, true, _25_v2, _25_v2, _25_v2))
		var _50_v18 _dafny.Array
		_ = _50_v18
		var _nwElement0_0 _dafny.Sequence = _dafny.SeqOf((func() bool {
			if (_47_v15).Contains(_33_cf5) {
				return (_47_v15).Get(_33_cf5).(bool)
			}
			return !(true)
		})(), _25_v2)
		_ = _nwElement0_0
		var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(21))
		_ = _nw1
		(_nw1).ArraySet1(_nwElement0_0, 0)
		(_nw1).ArraySet1(_48_v16, 1)
		(_nw1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_48_v16, _48_v16), 2)
		(_nw1).ArraySet1(_dafny.SeqOf(Companion_Default___.Fm1(_23_v0, _25_v2, _25_v2, true, globalState)), 3)
		(_nw1).ArraySet1(_48_v16, 4)
		(_nw1).ArraySet1(_dafny.SeqOf(false, !(_25_v2)), 5)
		(_nw1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_48_v16, _48_v16), 6)
		(_nw1).ArraySet1(_48_v16, 7)
		(_nw1).ArraySet1(_48_v16, 8)
		(_nw1).ArraySet1(_48_v16, 9)
		(_nw1).ArraySet1((_49_v17).Select((Companion_Default___.SafeIndex(_23_v0, _dafny.IntOfUint32((_49_v17).Cardinality()))).Uint32()).(_dafny.Sequence), 10)
		(_nw1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_48_v16, _48_v16), 11)
		(_nw1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_48_v16, _48_v16), 12)
		(_nw1).ArraySet1(_48_v16, 13)
		(_nw1).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_48_v16, _48_v16), 14)
		(_nw1).ArraySet1(_48_v16, 15)
		(_nw1).ArraySet1(_48_v16, 16)
		(_nw1).ArraySet1(_48_v16, 17)
		(_nw1).ArraySet1(_48_v16, 18)
		(_nw1).ArraySet1(_48_v16, 19)
		(_nw1).ArraySet1(_48_v16, 20)
		_50_v18 = _nw1
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_50_v18), 0))
		_ = _index2
		(_50_v18).ArraySet1(_48_v16, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_50_v18), 0))
		_ = _index3
		(_50_v18).ArraySet1(_48_v16, (_index3).Int())
		var _rhs0 bool = (func() bool {
			if true {
				return _25_v2
			}
			return (_25_v2) || (_25_v2)
		})()
		_ = _rhs0
		var _rhs1 bool = !(_25_v2)
		_ = _rhs1
		var _rhs2 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wmqcbpj")).Cardinality())
		_ = _rhs2
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		_lhs0.F1 = _rhs0
		_25_v2 = _rhs1
		r1 = _rhs2
	} else if _source1.Is_DC6() {
		var _51_v19 _dafny.Set
		_ = _51_v19
		_51_v19 = _dafny.SetOf(_25_v2, _25_v2, _25_v2)
		var _52_v20 *C0
		_ = _52_v20
		var _nw2 *C0 = New_C0_()
		_ = _nw2
		_nw2.Ctor__((_dafny.MultiSetOf(_25_v2, _25_v2, _25_v2)).Cardinality())
		_52_v20 = _nw2
		var _53_v21 _dafny.Set
		_ = _53_v21
		_53_v21 = _dafny.SetOf(_52_v20, _52_v20, _52_v20, _52_v20, _52_v20)
		var _54_v22 _dafny.Array
		_ = _54_v22
		var _nwElement0_1 bool = _25_v2
		_ = _nwElement0_1
		var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(11))
		_ = _nw3
		(_nw3).ArraySet1(_nwElement0_1, 0)
		(_nw3).ArraySet1(!(_25_v2) || (!(!(false))), 1)
		(_nw3).ArraySet1(_25_v2, 2)
		(_nw3).ArraySet1(_25_v2, 3)
		(_nw3).ArraySet1(_25_v2, 4)
		(_nw3).ArraySet1((_dafny.SetOf(_25_v2, true, _25_v2, false, _25_v2)).IsProperSubsetOf(_51_v19), 5)
		(_nw3).ArraySet1(_25_v2, 6)
		(_nw3).ArraySet1((_23_v0).Cmp(_23_v0) != 0, 7)
		(_nw3).ArraySet1(_25_v2, 8)
		(_nw3).ArraySet1(_25_v2, 9)
		(_nw3).ArraySet1((_53_v21).IsProperSubsetOf(_53_v21), 10)
		_54_v22 = _nw3
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))
		_ = _index4
		(_54_v22).ArraySet1(_25_v2, (_index4).Int())
		var _55_v23 _dafny.Sequence
		_ = _55_v23
		_55_v23 = _dafny.UnicodeSeqOfUtf8Bytes("h")
		var _56_v24 D2
		_ = _56_v24
		_56_v24 = Companion_D2_.Create_DC4_(_52_v20.F5)
		var _pat_let_tv0 = _52_v20
		_ = _pat_let_tv0
		var _pat_let_tv1 = _52_v20
		_ = _pat_let_tv1
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))
		_ = _index5
		(_54_v22).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_55_v23, _55_v23), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm6(_23_v0, func(_pat_let0_0 D2) D2 {
			return func(_57_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let1_0 _dafny.Int) D2 {
					return func(_58_dt__update_hcf3_h0 _dafny.Int) D2 {
						return Companion_D2_.Create_DC4_(_58_dt__update_hcf3_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0.F5)
			}(_pat_let0_0)
		}(_56_v24), globalState), (Companion_Default___.SafeIndex((_dafny.SetOf(_25_v2, _25_v2, _25_v2, _25_v2, _25_v2)).Cardinality(), _dafny.IntOfUint32((Companion_Default___.Fm6(_23_v0, func(_pat_let2_0 D2) D2 {
			return func(_59_dt__update__tmp_h1 D2) D2 {
				return func(_pat_let3_0 _dafny.Int) D2 {
					return func(_60_dt__update_hcf3_h1 _dafny.Int) D2 {
						return Companion_D2_.Create_DC4_(_60_dt__update_hcf3_h1)
					}(_pat_let3_0)
				}(_pat_let_tv1.F5)
			}(_pat_let2_0)
		}(_56_v24), globalState)).Cardinality()))).Uint32(), _28_v5)), (_index5).Int())
		var _61_v25 _dafny.Array
		_ = _61_v25
		var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
		_ = _nw4
		_61_v25 = _nw4
		var _62_v26 _dafny.Array
		_ = _62_v26
		var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw5
		_62_v26 = _nw5
		var _63_v27 _dafny.Sequence
		_ = _63_v27
		_63_v27 = _dafny.SeqOf(_62_v26)
		var _64_v28 _dafny.Sequence
		_ = _64_v28
		_64_v28 = _63_v27
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_61_v25), 0))
		_ = _index6
		(_61_v25).ArraySet1(_dafny.SeqOf(_64_v28), (_index6).Int())
		var _65_v29 D3
		_ = _65_v29
		_65_v29 = Companion_D3_.Create_DC10_((_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool), (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool), _23_v0, _23_v0, _52_v20.F5)
		var _66_v30 _dafny.Sequence
		_ = _66_v30
		_66_v30 = _dafny.SeqOf(_63_v27, (func() _dafny.Sequence {
			if (_65_v29).Dtor_cf13() {
				return _64_v28
			}
			return _64_v28
		})(), _64_v28, _64_v28)
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.ArrayLen((_61_v25), 0))
		_ = _index7
		(_61_v25).ArraySet1(_66_v30, (_index7).Int())
		var _67_v31 _dafny.MultiSet
		_ = _67_v31
		_67_v31 = _dafny.MultiSetOf(_52_v20)
		var _68_v32 _dafny.Sequence
		_ = _68_v32
		_68_v32 = _dafny.SeqOf((_67_v31).Cardinality(), Companion_Default___.Fm0(globalState))
		var _69_v33 _dafny.MultiSet
		_ = _69_v33
		_69_v33 = _dafny.MultiSetOf(_68_v32)
		_69_v33 = (_69_v33).Intersection(_69_v33)
		if (Companion_D3_.Create_DC10_(_25_v2, (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool), (_68_v32).Select((Companion_Default___.SafeIndex(_23_v0, _dafny.IntOfUint32((_68_v32).Cardinality()))).Uint32()).(_dafny.Int), _52_v20.F5, _23_v0)).Dtor_cf12() {
			var _70_v34 _dafny.Map
			_ = _70_v34
			_70_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v5, _52_v20.F5)
			var _71_v35 _dafny.MultiSet
			_ = _71_v35
			_71_v35 = _dafny.MultiSetOf(_dafny.IntOfInt64(-939), (_70_v34).Cardinality())
			(globalState).F1 = ((_26_v3).IsProperSubsetOf(_dafny.MultiSetOf(Companion_Default___.Fm1((func() _dafny.Int {
				if (_71_v35).Contains(_23_v0) {
					return (_71_v35).Multiplicity(_23_v0)
				}
				return _23_v0
			})(), Companion_Default___.Fm1(_52_v20.F5, _25_v2, (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool), _25_v2, globalState), (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool), true, globalState)))) == (!(_25_v2))
			var _72_v36 _dafny.Array
			_ = _72_v36
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
			_ = _nw6
			_72_v36 = _nw6
			var _73_v37 _dafny.MultiSet
			_ = _73_v37
			_73_v37 = _dafny.MultiSetOf(_72_v36)
			var _74_v38 _dafny.Map
			_ = _74_v38
			_74_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_73_v37).Cardinality(), _dafny.IntOfInt64(-224)), (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool))
			_74_v38 = (_74_v38).Update(_52_v20.F5, (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool))
			var _75_v39 *C0
			_ = _75_v39
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__(Companion_Default___.SafeModuloInt(_52_v20.F5, _dafny.IntOfInt64(844)))
			_75_v39 = _nw7
			var _nw8 *C0 = New_C0_()
			_ = _nw8
			_nw8.Ctor__((_75_v39.F5).Times(_23_v0))
			_75_v39 = _nw8
			(_52_v20).F5 = Companion_Default___.Fm0(globalState)
		} else {
			_68_v32 = _68_v32
			(globalState).F1 = _25_v2
			var _76_v40 _dafny.Array
			_ = _76_v40
			var _nwElement0_2 _dafny.CodePoint = Companion_Default___.Fm4(false, globalState)
			_ = _nwElement0_2
			var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(27))
			_ = _nw9
			(_nw9).ArraySet1CodePoint(_nwElement0_2, 0)
			(_nw9).ArraySet1CodePoint(_28_v5, 1)
			(_nw9).ArraySet1CodePoint(_28_v5, 2)
			(_nw9).ArraySet1CodePoint(_28_v5, 3)
			(_nw9).ArraySet1CodePoint(_28_v5, 4)
			(_nw9).ArraySet1CodePoint(_dafny.CodePoint('a'), 5)
			(_nw9).ArraySet1CodePoint(_28_v5, 6)
			(_nw9).ArraySet1CodePoint(_28_v5, 7)
			(_nw9).ArraySet1CodePoint(_28_v5, 8)
			(_nw9).ArraySet1CodePoint(_28_v5, 9)
			(_nw9).ArraySet1CodePoint(_28_v5, 10)
			(_nw9).ArraySet1CodePoint(_28_v5, 11)
			(_nw9).ArraySet1CodePoint(_28_v5, 12)
			(_nw9).ArraySet1CodePoint(_28_v5, 13)
			(_nw9).ArraySet1CodePoint(_28_v5, 14)
			(_nw9).ArraySet1CodePoint(_28_v5, 15)
			(_nw9).ArraySet1CodePoint(_28_v5, 16)
			(_nw9).ArraySet1CodePoint(_dafny.CodePoint('n'), 17)
			(_nw9).ArraySet1CodePoint(_28_v5, 18)
			(_nw9).ArraySet1CodePoint(_28_v5, 19)
			(_nw9).ArraySet1CodePoint(_28_v5, 20)
			(_nw9).ArraySet1CodePoint((func() _dafny.CodePoint {
				if Companion_Default___.Fm1(_52_v20.F5, !((_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool)), _25_v2, (_54_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_54_v22), 0))).Int()).(bool), globalState) {
					return _28_v5
				}
				return _dafny.CodePoint('c')
			})(), 21)
			(_nw9).ArraySet1CodePoint(_28_v5, 22)
			(_nw9).ArraySet1CodePoint(_dafny.CodePoint('n'), 23)
			(_nw9).ArraySet1CodePoint(_28_v5, 24)
			(_nw9).ArraySet1CodePoint(_28_v5, 25)
			(_nw9).ArraySet1CodePoint(_28_v5, 26)
			_76_v40 = _nw9
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_76_v40), 0))
			_ = _index8
			(_76_v40).ArraySet1CodePoint(_28_v5, (_index8).Int())
			var _77_v41 _dafny.Map
			_ = _77_v41
			_77_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v23, _dafny.Companion_Sequence_.Concatenate(_68_v32, _68_v32))
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_76_v40), 0))
			_ = _index9
			var _rhs3 _dafny.Int = _52_v20.F5
			_ = _rhs3
			var _rhs4 _dafny.CodePoint = _dafny.CodePoint('v')
			_ = _rhs4
			var _rhs5 bool = ((_52_v20.F5).Times(_52_v20.F5)).Cmp(_23_v0) == 0
			_ = _rhs5
			var _rhs6 _dafny.Sequence = (func() _dafny.Sequence {
				if (_77_v41).Contains(_dafny.Companion_Sequence_.Concatenate(_55_v23, _55_v23)) {
					return (_77_v41).Get(_dafny.Companion_Sequence_.Concatenate(_55_v23, _55_v23)).(_dafny.Sequence)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg7 _dafny.Int) interface{} {
						return coer7(arg7)
					}
				}((func(_78_v20 *C0) func(_dafny.Int) _dafny.Int {
					return func(_79_i3 _dafny.Int) _dafny.Int {
						return _78_v20.F5
					}
				})(_52_v20)))
			})()
			_ = _rhs6
			var _lhs1 *C0 = _52_v20
			_ = _lhs1
			var _lhs2 _dafny.Array = _76_v40
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_76_v40), 0))
			_ = _lhs3
			var _lhs4 *GlobalState = globalState
			_ = _lhs4
			_lhs1.F5 = _rhs3
			(_lhs2).ArraySet1CodePoint(_rhs4, (_lhs3).Int())
			_lhs4.F1 = _rhs5
			_68_v32 = _rhs6
			var _80_v42 *C0
			_ = _80_v42
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__(_52_v20.F5)
			_80_v42 = _nw10
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw11
			_54_v22 = _nw11
		}
	} else if _source1.Is_DC7() {
		var _81___mcc_h2 bool = _source1.Get_().(D2_DC7).Cf6
		_ = _81___mcc_h2
		var _82___mcc_h3 _dafny.Int = _source1.Get_().(D2_DC7).Cf7
		_ = _82___mcc_h3
		var _83___mcc_h4 bool = _source1.Get_().(D2_DC7).Cf8
		_ = _83___mcc_h4
		var _84___mcc_h5 _dafny.Map = _source1.Get_().(D2_DC7).Cf9
		_ = _84___mcc_h5
		var _85_cf9 _dafny.Map = _84___mcc_h5
		_ = _85_cf9
		var _86_cf8 bool = _83___mcc_h4
		_ = _86_cf8
		var _87_cf7 _dafny.Int = _82___mcc_h3
		_ = _87_cf7
		var _88_cf6 bool = _81___mcc_h2
		_ = _88_cf6
		var _89_v43 D2
		_ = _89_v43
		_89_v43 = Companion_D2_.Create_DC4_(_23_v0)
		var _90_v44 _dafny.Sequence
		_ = _90_v44
		_90_v44 = _dafny.UnicodeSeqOfUtf8Bytes("xautmy")
		var _91_v45 D5
		_ = _91_v45
		_91_v45 = Companion_D5_.Create_DC12_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_89_v43, _89_v43, _89_v43, _89_v43, _89_v43), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_90_v44).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_89_v43, _89_v43, _89_v43, _89_v43, _89_v43)).Cardinality()))).Uint32(), Companion_D2_.Create_DC4_(_23_v0)))
		var _92_v46 _dafny.Sequence
		_ = _92_v46
		_92_v46 = _dafny.SeqOf(_89_v43)
		var _93_v47 _dafny.Sequence
		_ = _93_v47
		_93_v47 = _dafny.SeqOf((_dafny.Zero).Minus(_87_cf7))
		var _pat_let_tv2 = _89_v43
		_ = _pat_let_tv2
		var _94_v48 _dafny.Array
		_ = _94_v48
		var _nwElement0_3 bool = (func() bool {
			if _86_cf8 {
				return _86_cf8
			}
			return _25_v2
		})()
		_ = _nwElement0_3
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(28))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_3, 0)
		(_nw12).ArraySet1(_88_cf6, 1)
		(_nw12).ArraySet1(_25_v2, 2)
		(_nw12).ArraySet1(true, 3)
		(_nw12).ArraySet1(_88_cf6, 4)
		(_nw12).ArraySet1(!(_86_cf8), 5)
		(_nw12).ArraySet1(_25_v2, 6)
		(_nw12).ArraySet1(_25_v2, 7)
		(_nw12).ArraySet1((_87_cf7).Cmp((_dafny.SetOf(_87_cf7)).Cardinality()) == 0, 8)
		(_nw12).ArraySet1(_88_cf6, 9)
		(_nw12).ArraySet1((_dafny.MultiSetFromSeq(_92_v46)).IsProperSubsetOf(_dafny.MultiSetFromSeq((func(_pat_let4_0 D5) D5 {
			return func(_95_dt__update__tmp_h2 D5) D5 {
				return func(_pat_let5_0 _dafny.Sequence) D5 {
					return func(_96_dt__update_hcf18_h0 _dafny.Sequence) D5 {
						return Companion_D5_.Create_DC12_(_96_dt__update_hcf18_h0)
					}(_pat_let5_0)
				}(_dafny.SeqOf(_pat_let_tv2))
			}(_pat_let4_0)
		}(_91_v45)).Dtor_cf18())), 10)
		(_nw12).ArraySet1((func() bool {
			if !(_25_v2) {
				return _86_cf8
			}
			return _25_v2
		})(), 11)
		(_nw12).ArraySet1(_86_cf8, 12)
		(_nw12).ArraySet1(_25_v2, 13)
		(_nw12).ArraySet1(!(_88_cf6), 14)
		(_nw12).ArraySet1((_dafny.IntOfUint32((Companion_Default___.Fm7(globalState)).Cardinality())).Cmp(_dafny.IntOfUint32((_90_v44).Cardinality())) >= 0, 15)
		(_nw12).ArraySet1(_86_cf8, 16)
		(_nw12).ArraySet1(((_dafny.Zero).Minus(_23_v0)).Cmp((_dafny.Zero).Minus(_23_v0)) != 0, 17)
		(_nw12).ArraySet1(_25_v2, 18)
		(_nw12).ArraySet1(_25_v2, 19)
		(_nw12).ArraySet1(Companion_Default___.Fm1(_dafny.IntOfUint32((_93_v47).Cardinality()), true, false, _86_cf8, globalState), 20)
		(_nw12).ArraySet1((_26_v3).IsProperSubsetOf(_26_v3), 21)
		(_nw12).ArraySet1(_88_cf6, 22)
		(_nw12).ArraySet1(_86_cf8, 23)
		(_nw12).ArraySet1(!((_30_v7).Dtor_cf8()), 24)
		(_nw12).ArraySet1(_86_cf8, 25)
		(_nw12).ArraySet1(!_dafny.Companion_Sequence_.Equal(_90_v44, _90_v44), 26)
		(_nw12).ArraySet1(_25_v2, 27)
		_94_v48 = _nw12
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_94_v48), 0))
		_ = _index10
		(_94_v48).ArraySet1(_25_v2, (_index10).Int())
		var _97_v49 _dafny.Map
		_ = _97_v49
		_97_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_23_v0, Companion_Default___.Fm0(globalState))
		var _98_v50 _dafny.Array
		_ = _98_v50
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
		_ = _nw13
		_98_v50 = _nw13
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_94_v48), 0))
		_ = _index11
		var _rhs7 _dafny.Array = (func() _dafny.Array {
			if true {
				return _98_v50
			}
			return _98_v50
		})()
		_ = _rhs7
		var _rhs8 _dafny.Array = _98_v50
		_ = _rhs8
		var _rhs9 bool = (func() bool {
			if _86_cf8 {
				return _25_v2
			}
			return _88_cf6
		})()
		_ = _rhs9
		var _rhs10 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_cf7, (_93_v47).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_23_v0), _dafny.IntOfUint32((_93_v47).Cardinality()))).Uint32()).(_dafny.Int))
		_ = _rhs10
		var _rhs11 _dafny.Array = _94_v48
		_ = _rhs11
		var _lhs5 _dafny.Array = _94_v48
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_94_v48), 0))
		_ = _lhs6
		r0 = _rhs7
		r0 = _rhs8
		(_lhs5).ArraySet1(_rhs9, (_lhs6).Int())
		_97_v49 = _rhs10
		_94_v48 = _rhs11
		var _99_v51 *C0
		_ = _99_v51
		var _nw14 *C0 = New_C0_()
		_ = _nw14
		_nw14.Ctor__(_23_v0)
		_99_v51 = _nw14
		var _100_v52 _dafny.Array
		_ = _100_v52
		var _nwElement0_4 *C0 = _99_v51
		_ = _nwElement0_4
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(6))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_4, 0)
		(_nw15).ArraySet1(_99_v51, 1)
		(_nw15).ArraySet1(_99_v51, 2)
		(_nw15).ArraySet1((func() *C0 {
			if Companion_Default___.Fm1(_23_v0, false, false, _88_cf6, globalState) {
				return _99_v51
			}
			return _99_v51
		})(), 3)
		(_nw15).ArraySet1(_99_v51, 4)
		(_nw15).ArraySet1(_99_v51, 5)
		_100_v52 = _nw15
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_100_v52), 0))
		_ = _index12
		(_100_v52).ArraySet1(_99_v51, (_index12).Int())
		var _101_v53 _dafny.Map
		_ = _101_v53
		_101_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_cf7, (func() _dafny.Sequence {
			if (_94_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_94_v48), 0))).Int()).(bool) {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}((func(_102_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_103_i4 _dafny.Int) _dafny.CodePoint {
						return _102_v5
					}
				})(_28_v5)))
			}
			return _90_v44
		})())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_100_v52), 0))
		_ = _index13
		var _rhs12 *C0 = _99_v51
		_ = _rhs12
		var _rhs13 _dafny.Map = (_101_v53).Merge(_101_v53)
		_ = _rhs13
		var _lhs7 _dafny.Array = _100_v52
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_100_v52), 0))
		_ = _lhs8
		(_lhs7).ArraySet1(_rhs12, (_lhs8).Int())
		_101_v53 = _rhs13
		if _25_v2 {
			_97_v49 = (_97_v49).Update((_dafny.Zero).Minus(_99_v51.F5), Companion_Default___.Fm0(globalState))
			var _104_v54 _dafny.MultiSet
			_ = _104_v54
			_104_v54 = _dafny.MultiSetOf(_98_v50)
			_104_v54 = _104_v54
			var _105_v55 _dafny.Set
			_ = _105_v55
			_105_v55 = _dafny.SetOf(!(!(false)))
			var _106_v56 _dafny.Array
			_ = _106_v56
			var _nwElement0_5 _dafny.Array = _98_v50
			_ = _nwElement0_5
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(14))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_5, 0)
			(_nw16).ArraySet1(_98_v50, 1)
			(_nw16).ArraySet1(_98_v50, 2)
			(_nw16).ArraySet1(_98_v50, 3)
			(_nw16).ArraySet1(_98_v50, 4)
			(_nw16).ArraySet1(_98_v50, 5)
			(_nw16).ArraySet1(_98_v50, 6)
			(_nw16).ArraySet1(_98_v50, 7)
			(_nw16).ArraySet1(_98_v50, 8)
			(_nw16).ArraySet1(_98_v50, 9)
			(_nw16).ArraySet1(_98_v50, 10)
			(_nw16).ArraySet1(_98_v50, 11)
			(_nw16).ArraySet1(_98_v50, 12)
			(_nw16).ArraySet1(_98_v50, 13)
			_106_v56 = _nw16
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_106_v56), 0))
			_ = _index14
			(_106_v56).ArraySet1(_98_v50, (_index14).Int())
			var _107_v57 _dafny.Array
			_ = _107_v57
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_1
			var _nw17 _dafny.Array
			_ = _nw17
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw17 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Sequence = (func(_108_v47 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_109_i5 _dafny.Int) _dafny.Sequence {
						return _108_v47
					}
				})(_93_v47)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw17 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw17).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw17).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_107_v57 = _nw17
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_107_v57), 0))
			_ = _index15
			(_107_v57).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_90_v44).Cardinality()), _87_cf7), (_index15).Int())
			var _110_v58 _dafny.Sequence
			_ = _110_v58
			_110_v58 = _dafny.SeqOf((_94_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_94_v48), 0))).Int()).(bool))
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_106_v56), 0))
			_ = _index16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_107_v57), 0))
			_ = _index17
			var _rhs14 _dafny.Set = Companion_Default___.Fm8(_88_cf6, _86_cf8, _dafny.Companion_Sequence_.Concatenate(_110_v58, _110_v58), (_105_v55).Equals(_105_v55), globalState)
			_ = _rhs14
			var _rhs15 _dafny.Array = _98_v50
			_ = _rhs15
			var _rhs16 _dafny.Sequence = Companion_Default___.Fm7(globalState)
			_ = _rhs16
			var _lhs9 _dafny.Array = _106_v56
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_106_v56), 0))
			_ = _lhs10
			var _lhs11 _dafny.Array = _107_v57
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_107_v57), 0))
			_ = _lhs12
			_105_v55 = _rhs14
			(_lhs9).ArraySet1(_rhs15, (_lhs10).Int())
			(_lhs11).ArraySet1(_rhs16, (_lhs12).Int())
			var _111_v59 _dafny.MultiSet
			_ = _111_v59
			_111_v59 = _dafny.MultiSetOf(_26_v3, _26_v3)
			_111_v59 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(422))).Uint32(), func(coer9 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_112_v3 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_113_i6 _dafny.Int) _dafny.MultiSet {
					return _112_v3
				}
			})(_26_v3))), _dafny.SeqOf(_26_v3, _dafny.MultiSetFromSeq(_110_v58), Companion_Default___.Fm9(globalState), _26_v3)))).Intersection(_111_v59)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_98_v50), 0))
			_ = _index18
			(_98_v50).ArraySet1(_99_v51.F5, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_98_v50), 0))
			_ = _index19
			(_98_v50).ArraySet1(_87_cf7, (_index19).Int())
		} else {
			_94_v48 = _94_v48
			var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt(_23_v0, (_dafny.Zero).Minus(_87_cf7))
			_ = _rhs17
			var _rhs18 bool = _88_cf6
			_ = _rhs18
			var _lhs13 *C0 = _99_v51
			_ = _lhs13
			var _lhs14 *GlobalState = globalState
			_ = _lhs14
			_lhs13.F5 = _rhs17
			_lhs14.F1 = _rhs18
			_86_cf8 = _25_v2
			_28_v5 = _28_v5
			var _114_v60 _dafny.Array
			_ = _114_v60
			var _nw18 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw18
			_114_v60 = _nw18
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_114_v60), 0))
			_ = _index20
			(_114_v60).ArraySet1(_90_v44, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_114_v60), 0))
			_ = _index21
			var _rhs19 _dafny.Array = _114_v60
			_ = _rhs19
			var _rhs20 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_90_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(621))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_115_v44 _dafny.Sequence, _116_v51 *C0) func(_dafny.Int) _dafny.CodePoint {
				return func(_117_i7 _dafny.Int) _dafny.CodePoint {
					return (_115_v44).Select((Companion_Default___.SafeIndex(_116_v51.F5, _dafny.IntOfUint32((_115_v44).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_90_v44, _99_v51)))), _90_v44)
			_ = _rhs20
			var _lhs15 _dafny.Array = _114_v60
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_114_v60), 0))
			_ = _lhs16
			_114_v60 = _rhs19
			(_lhs15).ArraySet1(_rhs20, (_lhs16).Int())
		}
		var _118_v61 _dafny.Sequence
		_ = _118_v61
		_118_v61 = _dafny.SeqOf(_98_v50)
		var _119_v62 _dafny.MultiSet
		_ = _119_v62
		_119_v62 = _dafny.MultiSetOf(_dafny.IntOfInt64(-259))
		var _120_v63 _dafny.Sequence
		_ = _120_v63
		_120_v63 = _dafny.SeqOf(_98_v50)
		_118_v61 = (func() _dafny.Sequence {
			if (_119_v62).IsProperSubsetOf(_119_v62) {
				return _120_v63
			}
			return _120_v63
		})()
	} else {
		var _121___mcc_h6 _dafny.Int = _source1.Get_().(D2_DC4).Cf3
		_ = _121___mcc_h6
		var _122_cf3 _dafny.Int = _121___mcc_h6
		_ = _122_cf3
		var _123_v64 _dafny.Array
		_ = _123_v64
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw19
		_123_v64 = _nw19
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_123_v64), 0))
		_ = _index22
		(_123_v64).ArraySet1(_25_v2, (_index22).Int())
		var _124_v65 _dafny.Map
		_ = _124_v65
		_124_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_23_v0, _23_v0), true)
		var _125_v66 D2
		_ = _125_v66
		_125_v66 = Companion_D2_.Create_DC4_(_23_v0)
		var _126_v67 _dafny.Sequence
		_ = _126_v67
		_126_v67 = _dafny.SeqOf(_25_v2)
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_123_v64), 0))
		_ = _index23
		var _rhs21 bool = _25_v2
		_ = _rhs21
		var _rhs22 bool = Companion_Default___.Fm1(_122_cf3, (func() bool {
			if !(Companion_Default___.Fm1(_122_cf3, _25_v2, true, _25_v2, globalState)) {
				return _25_v2
			}
			return (_126_v67).Select((Companion_Default___.SafeIndex(_122_cf3, _dafny.IntOfUint32((_126_v67).Cardinality()))).Uint32()).(bool)
		})(), (_122_cf3).Cmp(_23_v0) == 0, !(_25_v2), globalState)
		_ = _rhs22
		var _rhs23 _dafny.Map = (_124_v65).Update(_23_v0, _25_v2)
		_ = _rhs23
		var _rhs24 D2 = Companion_D2_.Create_DC4_(_23_v0)
		_ = _rhs24
		var _rhs25 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_122_cf3, _122_cf3))
		_ = _rhs25
		var _lhs17 _dafny.Array = _123_v64
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_123_v64), 0))
		_ = _lhs18
		var _lhs19 *GlobalState = globalState
		_ = _lhs19
		(_lhs17).ArraySet1(_rhs21, (_lhs18).Int())
		_lhs19.F1 = _rhs22
		_124_v65 = _rhs23
		_125_v66 = _rhs24
		_122_cf3 = _rhs25
		(globalState).F1 = (_123_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_123_v64), 0))).Int()).(bool)
		_28_v5 = _28_v5
		var _127_v68 _dafny.Map
		_ = _127_v68
		_127_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_28_v5, _dafny.IntOfInt64(382))
		var _128_v69 *C0
		_ = _128_v69
		var _nw20 *C0 = New_C0_()
		_ = _nw20
		_nw20.Ctor__(_23_v0)
		_128_v69 = _nw20
		var _129_v70 _dafny.Map
		_ = _129_v70
		_129_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_123_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_123_v64), 0))).Int()).(bool), (_123_v64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_123_v64), 0))).Int()).(bool))
		var _130_v71 _dafny.MultiSet
		_ = _130_v71
		_130_v71 = _dafny.MultiSetOf(_28_v5)
		var _131_v72 _dafny.Array
		_ = _131_v72
		var _nwElement0_6 _dafny.Int = _122_cf3
		_ = _nwElement0_6
		var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw21
		(_nw21).ArraySet1(_nwElement0_6, 0)
		(_nw21).ArraySet1((_129_v70).Cardinality(), 1)
		(_nw21).ArraySet1((_130_v71).Cardinality(), 2)
		(_nw21).ArraySet1(_128_v69.F5, 3)
		(_nw21).ArraySet1(_dafny.IntOfInt64(679), 4)
		(_nw21).ArraySet1((_dafny.Zero).Minus(_122_cf3), 5)
		(_nw21).ArraySet1(_128_v69.F5, 6)
		(_nw21).ArraySet1(_23_v0, 7)
		(_nw21).ArraySet1(_dafny.IntOfInt64(821), 8)
		(_nw21).ArraySet1(_128_v69.F5, 9)
		(_nw21).ArraySet1(_23_v0, 10)
		(_nw21).ArraySet1(_23_v0, 11)
		(_nw21).ArraySet1(_128_v69.F5, 12)
		_131_v72 = _nw21
		var _132_v73 _dafny.Sequence
		_ = _132_v73
		_132_v73 = _dafny.SeqOf(_131_v72, _131_v72)
		var _133_v74 _dafny.Sequence
		_ = _133_v74
		_133_v74 = _132_v73
		var _134_v75 _dafny.Map
		_ = _134_v75
		_134_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_128_v69)).Cardinality()), _133_v74)
		_127_v68 = (_127_v68).Update(_28_v5, (_134_v75).Cardinality())
	}
	var _135_v76 *C0
	_ = _135_v76
	var _nw22 *C0 = New_C0_()
	_ = _nw22
	_nw22.Ctor__((_23_v0).Times(_23_v0))
	_135_v76 = _nw22
	var _136_v77 _dafny.MultiSet
	_ = _136_v77
	_136_v77 = _dafny.MultiSetOf((_135_v76.F5).Times(_135_v76.F5), _23_v0)
	_136_v77 = (_136_v77).Intersection(_136_v77)
	var _pat_let_tv3 = _29_v6
	_ = _pat_let_tv3
	var _pat_let_tv4 = _23_v0
	_ = _pat_let_tv4
	_30_v7 = func(_pat_let6_0 D2) D2 {
		return func(_137_dt__update__tmp_h4 D2) D2 {
			return func(_pat_let7_0 _dafny.Map) D2 {
				return func(_138_dt__update_hcf9_h0 _dafny.Map) D2 {
					return func(_pat_let8_0 _dafny.Int) D2 {
						return func(_139_dt__update_hcf7_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC7_((_137_dt__update__tmp_h4).Dtor_cf6(), _139_dt__update_hcf7_h0, (_137_dt__update__tmp_h4).Dtor_cf8(), _138_dt__update_hcf9_h0)
						}(_pat_let8_0)
					}(_pat_let_tv4)
				}(_pat_let7_0)
			}(_pat_let_tv3)
		}(_pat_let6_0)
	}(_30_v7)
	var _140_v78 D3
	_ = _140_v78
	_140_v78 = Companion_D3_.Create_DC10_(_25_v2, _25_v2, _23_v0, _135_v76.F5, _23_v0)
	var _141_v79 _dafny.Set
	_ = _141_v79
	_141_v79 = _dafny.SetOf(_25_v2)
	var _142_v80 _dafny.Sequence
	_ = _142_v80
	_142_v80 = _dafny.SeqOf(Companion_Default___.Fm1(_135_v76.F5, _25_v2, _25_v2, _25_v2, globalState), _25_v2, false)
	var _143_v81 D3
	_ = _143_v81
	_143_v81 = Companion_D3_.Create_DC9_(_136_v77)
	var _144_v82 _dafny.MultiSet
	_ = _144_v82
	_144_v82 = _dafny.MultiSetOf(_143_v81, _143_v81)
	var _145_v83 _dafny.Set
	_ = _145_v83
	_145_v83 = _dafny.SetOf(_24_v1)
	var _146_v84 _dafny.Array
	_ = _146_v84
	var _nwElement0_7 bool = (_140_v78).Dtor_cf12()
	_ = _nwElement0_7
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(20))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_7, 0)
	(_nw23).ArraySet1(_25_v2, 1)
	(_nw23).ArraySet1(_25_v2, 2)
	(_nw23).ArraySet1((func() bool {
		if !(!(_25_v2)) {
			return _25_v2
		}
		return _25_v2
	})(), 3)
	(_nw23).ArraySet1((_135_v76.F5).Cmp((_dafny.Zero).Minus((_141_v79).Cardinality())) >= 0, 4)
	(_nw23).ArraySet1((_142_v80).Select((Companion_Default___.SafeIndex(_135_v76.F5, _dafny.IntOfUint32((_142_v80).Cardinality()))).Uint32()).(bool), 5)
	(_nw23).ArraySet1(_25_v2, 6)
	(_nw23).ArraySet1(_25_v2, 7)
	(_nw23).ArraySet1(_25_v2, 8)
	(_nw23).ArraySet1(_25_v2, 9)
	(_nw23).ArraySet1(_25_v2, 10)
	(_nw23).ArraySet1(_25_v2, 11)
	(_nw23).ArraySet1(_25_v2, 12)
	(_nw23).ArraySet1(_25_v2, 13)
	(_nw23).ArraySet1(true, 14)
	(_nw23).ArraySet1(_25_v2, 15)
	(_nw23).ArraySet1(_25_v2, 16)
	(_nw23).ArraySet1(_25_v2, 17)
	(_nw23).ArraySet1(!(_144_v82).Equals(_dafny.MultiSetOf(_143_v81)), 18)
	(_nw23).ArraySet1((_145_v83).IsSubsetOf(_145_v83), 19)
	_146_v84 = _nw23
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_146_v84), 0))
	_ = _index24
	(_146_v84).ArraySet1(_25_v2, (_index24).Int())
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_146_v84), 0))
	_ = _index25
	(_146_v84).ArraySet1((_140_v78).Dtor_cf12(), (_index25).Int())
	var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
	_ = _nw24
	r0 = _nw24
	var _147_v85 _dafny.Sequence
	_ = _147_v85
	_147_v85 = _dafny.UnicodeSeqOfUtf8Bytes("matixu")
	var _148_v86 _dafny.Map
	_ = _148_v86
	_148_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_147_v85, _dafny.SeqOf(_28_v5))
	var _149_v87 _dafny.Sequence
	_ = _149_v87
	_149_v87 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
		if (_148_v86).Contains(_147_v85) {
			return (_148_v86).Get(_147_v85).(_dafny.Sequence)
		}
		return _147_v85
	})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(814))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}((func(_150_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_151_i8 _dafny.Int) _dafny.CodePoint {
			return _150_v5
		}
	})(_28_v5))))).Cardinality()), _135_v76.F5, (_23_v0).Times(_135_v76.F5))
	r1 = (_149_v87).Select((Companion_Default___.SafeIndex(_23_v0, _dafny.IntOfUint32((_149_v87).Cardinality()))).Uint32()).(_dafny.Int)
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _152_v0 _dafny.Int
	_ = _152_v0
	_152_v0 = _dafny.IntOfInt64(-560)
	var _153_v1 _dafny.Sequence
	_ = _153_v1
	_153_v1 = _dafny.SeqOf(_152_v0, _152_v0)
	var _154_v2 _dafny.Array
	_ = _154_v2
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_2
	var _nw25 _dafny.Array
	_ = _nw25
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw25 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.MultiSet = func(_155_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))
		}
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
	_154_v2 = _nw25
	var _156_globalState *GlobalState
	_ = _156_globalState
	var _nw26 *GlobalState = New_GlobalState_()
	_ = _nw26
	_nw26.Ctor__(_dafny.IntOfInt64(791), false, _153_v1, _dafny.UnicodeSeqOfUtf8Bytes("kpno"), _154_v2)
	_156_globalState = _nw26
	var _157_v3 bool
	_ = _157_v3
	_157_v3 = true
	var _158_i1 _dafny.Int
	_ = _158_i1
	_158_i1 = _dafny.Zero
	{
		for _157_v3 {
			{
				if (_158_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_158_i1 = (_158_i1).Plus(_dafny.One)
				var _159_v4 _dafny.Array
				_ = _159_v4
				var _nw27 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw27
				_159_v4 = _nw27
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))
				_ = _index26
				(_159_v4).ArraySet1(false, (_index26).Int())
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))
				_ = _index27
				(_159_v4).ArraySet1(!(_157_v3) || (_157_v3), (_index27).Int())
				if !((_159_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))).Int()).(bool)) {
					var _160_v5 _dafny.Array
					_ = _160_v5
					var _161_v6 _dafny.Int
					_ = _161_v6
					var _out0 _dafny.Array
					_ = _out0
					var _out1 _dafny.Int
					_ = _out1
					_out0, _out1 = Companion_Default___.M0(_156_globalState)
					_160_v5 = _out0
					_161_v6 = _out1
					(_156_globalState).F1 = !(_157_v3) || (_157_v3)
					_153_v1 = _dafny.Companion_Sequence_.Update((Companion_D0_.Create_DC0_(_153_v1)).Dtor_cf0(), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_156_globalState), _dafny.IntOfUint32(((Companion_D0_.Create_DC0_(_153_v1)).Dtor_cf0()).Cardinality()))).Uint32(), _152_v0)
					(_156_globalState).F1 = (_159_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))).Int()).(bool)
					_152_v0 = Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_156_globalState), Companion_Default___.SafeDivisionInt(_161_v6, _152_v0))
				} else {
					var _162_v7 _dafny.Array
					_ = _162_v7
					var _163_v8 _dafny.Int
					_ = _163_v8
					var _out2 _dafny.Array
					_ = _out2
					var _out3 _dafny.Int
					_ = _out3
					_out2, _out3 = Companion_Default___.M0(_156_globalState)
					_162_v7 = _out2
					_163_v8 = _out3
					_157_v3 = (_159_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))).Int()).(bool)
					var _164_v9 D1
					_ = _164_v9
					_164_v9 = Companion_D1_.Create_DC2_((_159_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))).Int()).(bool))
					var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))
					_ = _index28
					(_159_v4).ArraySet1((_164_v9).Dtor_cf2(), (_index28).Int())
					var _165_v10 *C0
					_ = _165_v10
					var _nw28 *C0 = New_C0_()
					_ = _nw28
					_nw28.Ctor__(_152_v0)
					_165_v10 = _nw28
					var _rhs26 _dafny.Int = _163_v8
					_ = _rhs26
					var _rhs27 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(90), _152_v0)
					_ = _rhs27
					var _lhs20 *C0 = _165_v10
					_ = _lhs20
					_152_v0 = _rhs26
					_lhs20.F5 = _rhs27
				}
				(_156_globalState).F1 = true
				var _166_v11 _dafny.MultiSet
				_ = _166_v11
				_166_v11 = _dafny.MultiSetOf(_157_v3, (_159_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(258), _dafny.ArrayLen((_159_v4), 0))).Int()).(bool))
				var _167_v12 D0
				_ = _167_v12
				_167_v12 = Companion_D0_.Create_DC0_((func() _dafny.Sequence {
					if _157_v3 {
						return _153_v1
					}
					return _dafny.SeqOf((_166_v11).Cardinality())
				})())
				var _rhs28 _dafny.Int = (func() _dafny.Int {
					if _157_v3 {
						return _dafny.IntOfInt64(612)
					}
					return _152_v0
				})()
				_ = _rhs28
				var _rhs29 D0 = _167_v12
				_ = _rhs29
				_152_v0 = _rhs28
				_167_v12 = _rhs29
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _168_v13 _dafny.Array
	_ = _168_v13
	var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
	_ = _nw29
	_168_v13 = _nw29
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_168_v13), 0))); ; {
		_guard_loop_0, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _169_i2 _dafny.Int
		_169_i2 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_169_i2).Sign() != -1) && ((_169_i2).Cmp(_dafny.ArrayLen((_168_v13), 0)) < 0)) {
			(_168_v13).ArraySet1(Companion_Default___.SafeModuloInt(_169_i2, (Companion_D2_.Create_DC4_(_152_v0)).Dtor_cf3()), (_169_i2).Int())
		}
	}
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_168_v13), 0))); ; {
		_guard_loop_1, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _170_i3 _dafny.Int
		_170_i3 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_170_i3).Sign() != -1) && ((_170_i3).Cmp(_dafny.ArrayLen((_168_v13), 0)) < 0)) {
			(_168_v13).ArraySet1((_170_i3).Times(_152_v0), (_170_i3).Int())
		}
	}
	(_156_globalState).F3 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-708))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_171_i4 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))
	var _172_v14 *C0
	_ = _172_v14
	var _nw30 *C0 = New_C0_()
	_ = _nw30
	_nw30.Ctor__(_dafny.IntOfInt64(190))
	_172_v14 = _nw30
	var _173_v15 _dafny.Sequence
	_ = _173_v15
	_173_v15 = _dafny.UnicodeSeqOfUtf8Bytes("xqlvtuvhq")
	var _174_i5 _dafny.Int
	_ = _174_i5
	_174_i5 = _dafny.Zero
	{
		for (func() bool {
			if (_152_v0).Cmp(_152_v0) >= 0 {
				return _157_v3
			}
			return !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_178_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), _173_v15)
		})() {
			{
				if (_174_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_174_i5 = (_174_i5).Plus(_dafny.One)
				(_172_v14).F5 = _dafny.IntOfUint32((_173_v15).Cardinality())
				var _175_v16 _dafny.Map
				_ = _175_v16
				_175_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v3, _172_v14.F5)
				var _176_v17 _dafny.Map
				_ = _176_v17
				_176_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v0, _175_v16)
				var _177_v18 _dafny.Map
				_ = _177_v18
				_177_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v17, _172_v14.F5)
				_177_v18 = (_177_v18).Update(_176_v17, _172_v14.F5)
				_153_v1 = _153_v1
				_152_v0 = _172_v14.F5
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _179_v19 _dafny.CodePoint
	_ = _179_v19
	_179_v19 = _dafny.CodePoint('q')
	_179_v19 = (_173_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.IntOfUint32((_173_v15).Cardinality()))).Uint32()).(_dafny.CodePoint)
	var _180_v20 _dafny.Array
	_ = _180_v20
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(17)
	_ = _len0_3
	var _nw31 _dafny.Array
	_ = _nw31
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw31 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Map = (func(_181_v0 _dafny.Int, _182_v14 *C0) func(_dafny.Int) _dafny.Map {
			return func(_183_i7 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v0, _182_v14.F5)
			}
		})(_152_v0, _172_v14)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw31 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw31).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw31).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_180_v20 = _nw31
	var _184_v21 _dafny.Map
	_ = _184_v21
	_184_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v3, _152_v0)
	var _185_v22 _dafny.Map
	_ = _185_v22
	_185_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v14.F5, (_184_v21).Cardinality())
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_180_v20), 0))
	_ = _index29
	(_180_v20).ArraySet1(_185_v22, (_index29).Int())
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_180_v20), 0))
	_ = _index30
	(_180_v20).ArraySet1(((_185_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_172_v14.F5), _172_v14.F5))).Update(_152_v0, _172_v14.F5), (_index30).Int())
	var _hi0 _dafny.Int = _152_v0
	_ = _hi0
	for _186_i8 := _dafny.IntOfInt64(232); _186_i8.Cmp(_hi0) < 0; _186_i8 = _186_i8.Plus(_dafny.One) {
		var _187_v23 *C0
		_ = _187_v23
		var _nw32 *C0 = New_C0_()
		_ = _nw32
		_nw32.Ctor__(_152_v0)
		_187_v23 = _nw32
		var _188_v24 _dafny.Array
		_ = _188_v24
		var _189_v25 _dafny.Int
		_ = _189_v25
		var _out4 _dafny.Array
		_ = _out4
		var _out5 _dafny.Int
		_ = _out5
		_out4, _out5 = Companion_Default___.M0(_156_globalState)
		_188_v24 = _out4
		_189_v25 = _out5
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_188_v24), 0))
		_ = _index31
		(_188_v24).ArraySet1(_186_i8, (_index31).Int())
		var _190_v26 _dafny.Array
		_ = _190_v26
		var _nwElement0_8 bool = !(_157_v3)
		_ = _nwElement0_8
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(13))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_8, 0)
		(_nw33).ArraySet1(false, 1)
		(_nw33).ArraySet1(!(_157_v3), 2)
		(_nw33).ArraySet1(_157_v3, 3)
		(_nw33).ArraySet1(_157_v3, 4)
		(_nw33).ArraySet1(_157_v3, 5)
		(_nw33).ArraySet1(_157_v3, 6)
		(_nw33).ArraySet1(_157_v3, 7)
		(_nw33).ArraySet1(_157_v3, 8)
		(_nw33).ArraySet1(_157_v3, 9)
		(_nw33).ArraySet1(_157_v3, 10)
		(_nw33).ArraySet1(_157_v3, 11)
		(_nw33).ArraySet1(_157_v3, 12)
		_190_v26 = _nw33
		var _191_v27 _dafny.Map
		_ = _191_v27
		_191_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_190_v26, Companion_Default___.Fm1(_189_v25, false, _157_v3, _157_v3, _156_globalState))
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_188_v24), 0))
		_ = _index32
		var _rhs30 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm0(_156_globalState))
		_ = _rhs30
		var _rhs31 _dafny.Map = _191_v27
		_ = _rhs31
		var _rhs32 _dafny.Int = (_dafny.Zero).Minus(_172_v14.F5)
		_ = _rhs32
		var _lhs21 _dafny.Array = _188_v24
		_ = _lhs21
		var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_188_v24), 0))
		_ = _lhs22
		(_lhs21).ArraySet1(_rhs30, (_lhs22).Int())
		_191_v27 = _rhs31
		_189_v25 = _rhs32
		var _source2 D0 = Companion_Default___.Fm2(_157_v3, _dafny.IntOfInt64(321), _156_globalState)
		_ = _source2
		if _source2.Is_DC1() {
			var _192___mcc_h0 _dafny.Array = _source2.Get_().(D0_DC1).Cf1
			_ = _192___mcc_h0
			var _193_cf1 _dafny.Array = _192___mcc_h0
			_ = _193_cf1
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_188_v24), 0))
			_ = _index33
			(_188_v24).ArraySet1((_172_v14.F5).Minus((_172_v14.F5).Times(Companion_Default___.Fm0(_156_globalState))), (_index33).Int())
			var _194_v28 *C0
			_ = _194_v28
			var _nw34 *C0 = New_C0_()
			_ = _nw34
			_nw34.Ctor__(_152_v0)
			_194_v28 = _nw34
			var _195_v29 D2
			_ = _195_v29
			_195_v29 = Companion_D2_.Create_DC5_(_172_v14.F5, _186_i8)
			var _pat_let_tv5 = _152_v0
			_ = _pat_let_tv5
			var _196_v30 _dafny.Map
			_ = _196_v30
			_196_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_173_v15, func(_pat_let9_0 D2) D2 {
				return func(_197_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let10_0 _dafny.Int) D2 {
						return func(_198_dt__update_hcf5_h0 _dafny.Int) D2 {
							return Companion_D2_.Create_DC5_((_197_dt__update__tmp_h0).Dtor_cf4(), _198_dt__update_hcf5_h0)
						}(_pat_let10_0)
					}(_pat_let_tv5)
				}(_pat_let9_0)
			}(_195_v29))
			_196_v30 = _196_v30
			var _199_v31 _dafny.Sequence
			_ = _199_v31
			_199_v31 = _dafny.SeqOf(_190_v26)
			var _200_v32 _dafny.Sequence
			_ = _200_v32
			_200_v32 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_157_v3), _dafny.SeqOf(_157_v3)))
			var _201_v33 _dafny.Array
			_ = _201_v33
			var _nwElement0_9 _dafny.Array = _190_v26
			_ = _nwElement0_9
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(25))
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_9, 0)
			(_nw35).ArraySet1(_190_v26, 1)
			(_nw35).ArraySet1(_190_v26, 2)
			(_nw35).ArraySet1(_190_v26, 3)
			(_nw35).ArraySet1(_190_v26, 4)
			(_nw35).ArraySet1(_190_v26, 5)
			(_nw35).ArraySet1(_190_v26, 6)
			(_nw35).ArraySet1(_190_v26, 7)
			(_nw35).ArraySet1(_190_v26, 8)
			(_nw35).ArraySet1(_190_v26, 9)
			(_nw35).ArraySet1(_190_v26, 10)
			(_nw35).ArraySet1(_190_v26, 11)
			(_nw35).ArraySet1(_190_v26, 12)
			(_nw35).ArraySet1((_199_v31).Select((Companion_Default___.SafeIndex(((_200_v32).Select((Companion_Default___.SafeIndex(_187_v23.F5, _dafny.IntOfUint32((_200_v32).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), _dafny.IntOfUint32((_199_v31).Cardinality()))).Uint32()).(_dafny.Array), 13)
			(_nw35).ArraySet1(_190_v26, 14)
			(_nw35).ArraySet1(_190_v26, 15)
			(_nw35).ArraySet1(_190_v26, 16)
			(_nw35).ArraySet1(_190_v26, 17)
			(_nw35).ArraySet1(_190_v26, 18)
			(_nw35).ArraySet1(_190_v26, 19)
			(_nw35).ArraySet1(_190_v26, 20)
			(_nw35).ArraySet1(_190_v26, 21)
			(_nw35).ArraySet1(_190_v26, 22)
			(_nw35).ArraySet1(_190_v26, 23)
			(_nw35).ArraySet1(_190_v26, 24)
			_201_v33 = _nw35
			var _202_v34 _dafny.Map
			_ = _202_v34
			_202_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v14, _190_v26)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_201_v33), 0))
			_ = _index34
			(_201_v33).ArraySet1((func() _dafny.Array {
				if (_202_v34).Contains(_187_v23) {
					return (_202_v34).Get(_187_v23).(_dafny.Array)
				}
				return _190_v26
			})(), (_index34).Int())
			var _203_v35 _dafny.Map
			_ = _203_v35
			_203_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_180_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_180_v20), 0))).Int()).(_dafny.Map)).Cardinality(), _157_v3)
			var _204_v36 _dafny.MultiSet
			_ = _204_v36
			_204_v36 = _dafny.MultiSetOf(_203_v35)
			var _205_v37 _dafny.MultiSet
			_ = _205_v37
			_205_v37 = _dafny.MultiSetOf((_204_v36).Cardinality(), _186_i8, (_dafny.Zero).Minus((_188_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_188_v24), 0))).Int()).(_dafny.Int)), _194_v28.F5, _189_v25)
			var _206_v38 _dafny.Sequence
			_ = _206_v38
			_206_v38 = _dafny.SeqOf(_194_v28)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_201_v33), 0))
			_ = _index35
			var _rhs33 bool = ((_205_v37).Update(_dafny.IntOfUint32((_206_v38).Cardinality()), Companion_Default___.Abs(_dafny.IntOfUint32((_173_v15).Cardinality())))).IsDisjointFrom(_205_v37)
			_ = _rhs33
			var _rhs34 _dafny.Array = _190_v26
			_ = _rhs34
			var _lhs23 *GlobalState = _156_globalState
			_ = _lhs23
			var _lhs24 _dafny.Array = _201_v33
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_201_v33), 0))
			_ = _lhs25
			_lhs23.F1 = _rhs33
			(_lhs24).ArraySet1(_rhs34, (_lhs25).Int())
		} else {
			var _207___mcc_h1 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf0
			_ = _207___mcc_h1
			var _208_cf0 _dafny.Sequence = _207___mcc_h1
			_ = _208_cf0
			var _209_v39 *C0
			_ = _209_v39
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__((_188_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_188_v24), 0))).Int()).(_dafny.Int))
			_209_v39 = _nw36
			var _210_v40 _dafny.MultiSet
			_ = _210_v40
			_210_v40 = _dafny.MultiSetOf((_172_v14.F5).Cmp(_187_v23.F5) >= 0)
			var _211_v41 _dafny.Sequence
			_ = _211_v41
			_211_v41 = _dafny.SeqOf(_157_v3, _157_v3)
			_210_v40 = _dafny.MultiSetOf((func() bool {
				if _157_v3 {
					return _157_v3
				}
				return false
			})(), (_211_v41).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_156_globalState), _dafny.IntOfUint32((_211_v41).Cardinality()))).Uint32()).(bool), (_210_v40).IsProperSubsetOf(_210_v40))
			var _212_v42 _dafny.Array
			_ = _212_v42
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw37
			_212_v42 = _nw37
			var _213_v43 D0
			_ = _213_v43
			_213_v43 = Companion_D0_.Create_DC1_(_212_v42)
			var _214_v44 _dafny.Map
			_ = _214_v44
			_214_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v43, false)
			var _215_v45 D3
			_ = _215_v45
			_215_v45 = Companion_D3_.Create_DC8_(_214_v44)
			var _216_v46 _dafny.Sequence
			_ = _216_v46
			_216_v46 = _dafny.SeqOf((_214_v44).Merge((_215_v45).Dtor_cf10()), _214_v44, _214_v44, (_214_v44).Update(Companion_D0_.Create_DC1_(_212_v42), _157_v3))
			_189_v25 = _dafny.IntOfUint32((_216_v46).Cardinality())
			(_209_v39).F5 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_189_v25))
		}
	}
	if _157_v3 {
		var _217_v47 _dafny.Array
		_ = _217_v47
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_4
		var _nw38 _dafny.Array
		_ = _nw38
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw38 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) bool = func(_218_i9 _dafny.Int) bool {
				return false
			}
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw38 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw38).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw38).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_217_v47 = _nw38
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_217_v47), 0))
		_ = _index36
		(_217_v47).ArraySet1(_157_v3, (_index36).Int())
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_217_v47), 0))
		_ = _index37
		(_217_v47).ArraySet1(!(_157_v3), (_index37).Int())
		(_172_v14).F5 = _152_v0
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_168_v13), 0))
		_ = _index38
		(_168_v13).ArraySet1((func() _dafny.Int {
			if ((_180_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_180_v20), 0))).Int()).(_dafny.Map)).Contains((_dafny.Zero).Minus(_172_v14.F5)) {
				return ((_180_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_180_v20), 0))).Int()).(_dafny.Map)).Get((_dafny.Zero).Minus(_172_v14.F5)).(_dafny.Int)
			}
			return _152_v0
		})(), (_index38).Int())
		var _219_v48 _dafny.Sequence
		_ = _219_v48
		_219_v48 = _dafny.SeqOf(_168_v13, _168_v13, _168_v13, _168_v13)
		var _220_v50 _dafny.Set
		_ = _220_v50
		_220_v50 = _dafny.SetOf(_dafny.IntOfUint32((_153_v1).Cardinality()), _152_v0)
		var _221_v51 _dafny.Map
		_ = _221_v51
		_221_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v14, _168_v13)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_217_v47), 0))
		_ = _index39
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_168_v13), 0))
		_ = _index40
		var _rhs35 bool = ((_220_v50).Union(Companion_Default___.Fm3((_217_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_217_v47), 0))).Int()).(bool), !(_157_v3), _156_globalState))).IsSubsetOf(func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(803), _dafny.IntOfInt64(238))); ; {
				_compr_6, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _222_v49 _dafny.Int
				_222_v49 = interface{}(_compr_6).(_dafny.Int)
				if ((_dafny.IntOfInt64(803)).Cmp(_222_v49) <= 0) && ((_222_v49).Cmp(_dafny.IntOfInt64(238)) < 0) {
					_coll6.Add(Companion_Default___.SafeDivisionInt(_222_v49, _172_v14.F5))
				}
			}
			return _coll6.ToSet()
		}())
		_ = _rhs35
		var _rhs36 _dafny.Int = Companion_Default___.Fm0(_156_globalState)
		_ = _rhs36
		var _rhs37 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_219_v48, (_dafny.SeqOf((func() _dafny.Array {
			if (_221_v51).Contains(_172_v14) {
				return (_221_v51).Get(_172_v14).(_dafny.Array)
			}
			return _168_v13
		})()))), _dafny.SeqOf(_168_v13))
		_ = _rhs37
		var _rhs38 *C0 = _172_v14
		_ = _rhs38
		var _lhs26 _dafny.Array = _217_v47
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_217_v47), 0))
		_ = _lhs27
		var _lhs28 _dafny.Array = _168_v13
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_168_v13), 0))
		_ = _lhs29
		(_lhs26).ArraySet1(_rhs35, (_lhs27).Int())
		(_lhs28).ArraySet1(_rhs36, (_lhs29).Int())
		_219_v48 = _rhs37
		_172_v14 = _rhs38
		(_156_globalState).F3 = _173_v15
		var _223_v52 _dafny.Array
		_ = _223_v52
		var _224_v53 _dafny.Int
		_ = _224_v53
		var _out6 _dafny.Array
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out6, _out7 = Companion_Default___.M0(_156_globalState)
		_223_v52 = _out6
		_224_v53 = _out7
	} else {
		_152_v0 = (_dafny.Zero).Minus(_172_v14.F5)
		var _225_v54 _dafny.MultiSet
		_ = _225_v54
		_225_v54 = _dafny.MultiSetOf(_152_v0)
		if Companion_Default___.Fm1(_172_v14.F5, (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(930))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_226_v14 *C0) func(_dafny.Int) _dafny.Int {
			return func(_227_i10 _dafny.Int) _dafny.Int {
				return _226_v14.F5
			}
		})(_172_v14)))).Cardinality())).Cmp(_152_v0) != 0, ((_225_v54).Cardinality()).Cmp(_152_v0) < 0, _157_v3, _156_globalState) {
			var _228_v55 _dafny.Array
			_ = _228_v55
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw39
			_228_v55 = _nw39
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))
			_ = _index41
			(_228_v55).ArraySet1(!((_172_v14.F5).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_173_v15).Cardinality()))) >= 0), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))
			_ = _index42
			(_228_v55).ArraySet1(_157_v3, (_index42).Int())
			(_156_globalState).F1 = _157_v3
			var _229_v56 _dafny.Sequence
			_ = _229_v56
			_229_v56 = _dafny.SeqOf((_228_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))).Int()).(bool), (_228_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))).Int()).(bool))
			var _230_v57 _dafny.Sequence
			_ = _230_v57
			_230_v57 = _dafny.SeqOf(_229_v56, _229_v56)
			var _231_v58 _dafny.Set
			_ = _231_v58
			_231_v58 = _dafny.SetOf(_152_v0, Companion_Default___.SafeModuloInt(_152_v0, _172_v14.F5), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_173_v15, (Companion_Default___.SafeIndex((_dafny.MultiSetOf(Companion_Default___.Fm1(_dafny.IntOfUint32(((_230_v57).Select((Companion_Default___.SafeIndex(_172_v14.F5, _dafny.IntOfUint32((_230_v57).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_228_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))).Int()).(bool), true, _157_v3, _156_globalState), (_228_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))).Int()).(bool))).Cardinality(), _dafny.IntOfUint32((_173_v15).Cardinality()))).Uint32(), Companion_Default___.Fm4((_228_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))).Int()).(bool), _156_globalState))).Cardinality()), _dafny.IntOfInt64(321))
			_231_v58 = ((_231_v58).Union(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter9 := _dafny.Iterate((_153_v1).Elements()); ; {
					_compr_7, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _232_v59 _dafny.Int
					_232_v59 = interface{}(_compr_7).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_153_v1, _232_v59) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_232_v59, _dafny.IntOfInt64(-437)))
					}
				}
				return _coll7.ToSet()
			}())).Intersection(_dafny.SetOf(_152_v0))
			_179_v19 = (func() _dafny.CodePoint {
				if true {
					return _179_v19
				}
				return _179_v19
			})()
			var _233_v60 _dafny.Array
			_ = _233_v60
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_5
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_234_v19 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_235_i11 _dafny.Int) _dafny.CodePoint {
						return _234_v19
					}
				})(_179_v19)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw40 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw40).ArraySet1CodePoint(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw40).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_233_v60 = _nw40
			var _236_v61 _dafny.Map
			_ = _236_v61
			_236_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v0, _233_v60)
			_233_v60 = (func() _dafny.Array {
				if (_236_v61).Contains(_152_v0) {
					return (_236_v61).Get(_152_v0).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if (_228_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_228_v55), 0))).Int()).(bool) {
						return _233_v60
					}
					return _233_v60
				})()
			})()
		} else {
			var _237_v62 D2
			_ = _237_v62
			_237_v62 = Companion_D2_.Create_DC5_(_172_v14.F5, _172_v14.F5)
			_237_v62 = _237_v62
			var _238_v63 _dafny.MultiSet
			_ = _238_v63
			_238_v63 = _dafny.MultiSetOf(_157_v3)
			var _239_v64 _dafny.Map
			_ = _239_v64
			_239_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_172_v14, _172_v14)
			var _240_v65 _dafny.Sequence
			_ = _240_v65
			_240_v65 = _dafny.SeqOf(_157_v3)
			var _nwElement0_10 _dafny.Int = _172_v14.F5
			_ = _nwElement0_10
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(8))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_10, 0)
			(_nw41).ArraySet1(_172_v14.F5, 1)
			(_nw41).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_238_v63).Cardinality()), _dafny.IntOfUint32((_173_v15).Cardinality())), 2)
			(_nw41).ArraySet1(Companion_Default___.SafeDivisionInt(_172_v14.F5, _172_v14.F5), 3)
			(_nw41).ArraySet1((_239_v64).Cardinality(), 4)
			(_nw41).ArraySet1(Companion_Default___.Fm0(_156_globalState), 5)
			(_nw41).ArraySet1(_dafny.IntOfUint32((_240_v65).Cardinality()), 6)
			(_nw41).ArraySet1(_172_v14.F5, 7)
			_168_v13 = _nw41
			_152_v0 = (_172_v14.F5).Minus(_172_v14.F5)
			(_172_v14).F5 = _152_v0
			var _241_v66 _dafny.Array
			_ = _241_v66
			var _242_v67 _dafny.Int
			_ = _242_v67
			var _out8 _dafny.Array
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			_out8, _out9 = Companion_Default___.M0(_156_globalState)
			_241_v66 = _out8
			_242_v67 = _out9
		}
		(_156_globalState).F1 = _157_v3
		var _243_v68 D0
		_ = _243_v68
		_243_v68 = Companion_D0_.Create_DC0_(_153_v1)
		var _pat_let_tv6 = _153_v1
		_ = _pat_let_tv6
		_243_v68 = func(_pat_let11_0 D0) D0 {
			return func(_244_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let12_0 _dafny.Sequence) D0 {
					return func(_245_dt__update_hcf0_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_245_dt__update_hcf0_h0)
					}(_pat_let12_0)
				}(_pat_let_tv6)
			}(_pat_let11_0)
		}(_243_v68)
		var _246_v69 _dafny.Set
		_ = _246_v69
		_246_v69 = _dafny.SetOf(_157_v3)
		var _247_v70 _dafny.Set
		_ = _247_v70
		_247_v70 = _dafny.SetOf(_246_v69, _246_v69)
		_247_v70 = (_247_v70).Union(_247_v70)
	}
	var _248_v71 _dafny.Array
	_ = _248_v71
	var _len0_6 _dafny.Int = _dafny.One
	_ = _len0_6
	var _nw42 _dafny.Array
	_ = _nw42
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw42 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = (func(_249_v3 bool) func(_dafny.Int) bool {
			return func(_250_i12 _dafny.Int) bool {
				return _249_v3
			}
		})(_157_v3)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw42 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw42).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw42).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_248_v71 = _nw42
	_248_v71 = _248_v71
	var _251_v72 _dafny.Map
	_ = _251_v72
	_251_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v3, _248_v71)
	_251_v72 = (_251_v72).Update(false, _248_v71)
	(_172_v14).F5 = Companion_Default___.SafeModuloInt(_172_v14.F5, Companion_Default___.SafeModuloInt((_153_v1).Select((Companion_Default___.SafeIndex(_172_v14.F5, _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32()).(_dafny.Int), Companion_Default___.Fm0(_156_globalState)))
	var _252_v73 _dafny.Array
	_ = _252_v73
	var _253_v74 _dafny.Int
	_ = _253_v74
	var _out10 _dafny.Array
	_ = _out10
	var _out11 _dafny.Int
	_ = _out11
	_out10, _out11 = Companion_Default___.M0(_156_globalState)
	_252_v73 = _out10
	_253_v74 = _out11
	(_156_globalState).F1 = _157_v3
	var _hi1 _dafny.Int = _172_v14.F5
	_ = _hi1
	for _254_i13 := _172_v14.F5; _254_i13.Cmp(_hi1) < 0; _254_i13 = _254_i13.Plus(_dafny.One) {
		var _255_v75 _dafny.Array
		_ = _255_v75
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
		_ = _nw43
		_255_v75 = _nw43
		var _256_v76 _dafny.Array
		_ = _256_v76
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
		_ = _nw44
		_256_v76 = _nw44
		var _257_v77 _dafny.Map
		_ = _257_v77
		_257_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_256_v76, _172_v14.F5)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_255_v75), 0))
		_ = _index43
		(_255_v75).ArraySet1((func() _dafny.Map {
			if _157_v3 {
				return _257_v77
			}
			return _257_v77
		})(), (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(843), _dafny.ArrayLen((_255_v75), 0))
		_ = _index44
		(_255_v75).ArraySet1(_257_v77, (_index44).Int())
		_157_v3 = ((_153_v1).Select((Companion_Default___.SafeIndex(_152_v0, _dafny.IntOfUint32((_153_v1).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_172_v14.F5) < 0
		_152_v0 = (_dafny.IntOfUint32((_173_v15).Cardinality())).Plus(_172_v14.F5)
		(_172_v14).F5 = (_152_v0).Times(_254_i13)
	}
	_dafny.Print(_152_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_153_v1, _dafny.SeqOf(_dafny.IntOfInt64(-560), _dafny.IntOfInt64(-560))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v2).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_156_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_156_globalState).F2(), _dafny.SeqOf(_dafny.IntOfInt64(-560), _dafny.IntOfInt64(-560))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_156_globalState.F3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_156_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.CodePoint('r'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_157_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_158_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v13).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v13).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_168_v13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_172_v14.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_173_v15.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_174_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_179_v19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9)).UpdateUnsafe(_dafny.IntOfInt64(-9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_180_v20).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v21).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_v22).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_248_v71).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_251_v72).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_253_v74)
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1}}
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
	return Companion_D0_.Create_DC1_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D0) Dtor_cf1() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf1
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ")"
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
			return ok && data1.Cf1 == data2.Cf1
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
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_() D1 {
	return D1{D1_DC3{}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf2 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 bool) D1 {
	return D1{D1_DC2{Cf2}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_()
}

func (_this D1) Dtor_cf2() bool {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ")"
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
			_, ok := other.Get_().(D1_DC3)
			return ok
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2 == data2.Cf2
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
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf4 _dafny.Int, Cf5 _dafny.Int) D2 {
	return D2{D2_DC5{Cf4, Cf5}}
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

type D2_DC7 struct {
	Cf6 bool
	Cf7 _dafny.Int
	Cf8 bool
	Cf9 _dafny.Map
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf6 bool, Cf7 _dafny.Int, Cf8 bool, Cf9 _dafny.Map) D2 {
	return D2{D2_DC7{Cf6, Cf7, Cf8, Cf9}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC4 struct {
	Cf3 _dafny.Int
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf3 _dafny.Int) D2 {
	return D2{D2_DC4{Cf3}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf4
}

func (_this D2) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC7).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf3
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D2_DC6:
		{
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8 == data2.Cf8 && data1.Cf9.Equals(data2.Cf9)
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0
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

type D3_DC9 struct {
	Cf11 _dafny.MultiSet
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf11 _dafny.MultiSet) D3 {
	return D3{D3_DC9{Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC10 struct {
	Cf12 bool
	Cf13 bool
	Cf14 _dafny.Int
	Cf15 _dafny.Int
	Cf16 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf12 bool, Cf13 bool, Cf14 _dafny.Int, Cf15 _dafny.Int, Cf16 _dafny.Int) D3 {
	return D3{D3_DC10{Cf12, Cf13, Cf14, Cf15, Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC8 struct {
	Cf10 _dafny.Map
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf10 _dafny.Map) D3 {
	return D3{D3_DC8{Cf10}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.EmptyMultiSet)
}

func (_this D3) Dtor_cf11() _dafny.MultiSet {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf12() bool {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC10).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) Dtor_cf10() _dafny.Map {
	return _this.Get_().(D3_DC8).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
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

type D4_DC11 struct {
	Cf17 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf17() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	Cf19 D0
	Cf20 _dafny.Int
	Cf21 _dafny.Int
	Cf22 _dafny.Int
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf19 D0, Cf20 _dafny.Int, Cf21 _dafny.Int, Cf22 _dafny.Int) D5 {
	return D5{D5_DC13{Cf19, Cf20, Cf21, Cf22}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC12 struct {
	Cf18 _dafny.Sequence
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 _dafny.Sequence) D5 {
	return D5{D5_DC12{Cf18}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(Companion_D0_.Default(), _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D5) Dtor_cf19() D0 {
	return _this.Get_().(D5_DC13).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf20
}

func (_this D5) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf22
}

func (_this D5) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

// Definition of class GlobalState
type GlobalState struct {
	F1  bool
	F3  _dafny.Sequence
	_f0 _dafny.Int
	_f2 _dafny.Sequence
	_f4 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = false
	_this.F3 = _dafny.EmptySeq
	_this._f0 = _dafny.Zero
	_this._f2 = _dafny.EmptySeq
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Sequence, f3 _dafny.Sequence, f4 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Sequence {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F5 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F5 = _dafny.Zero
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

func (_this *C0) Ctor__(f5 _dafny.Int) {
	{
		(_this).F5 = f5
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
