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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.CodePoint, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(719))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfInt64(658)).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(788), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(161))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(!(!(false))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(971))).Contains(false))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D0 = Companion_D0_.Create_DC1_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(435)), false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(192))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality())), _dafny.IntOfInt64(848))
	_ = _source0
	if _source0.Is_DC1() {
		var _2___mcc_h0 bool = _source0.Get_().(D0_DC1).Cf1
		_ = _2___mcc_h0
		var _3___mcc_h1 _dafny.Map = _source0.Get_().(D0_DC1).Cf2
		_ = _3___mcc_h1
		var _4___mcc_h2 bool = _source0.Get_().(D0_DC1).Cf3
		_ = _4___mcc_h2
		var _5___mcc_h3 _dafny.Int = _source0.Get_().(D0_DC1).Cf4
		_ = _5___mcc_h3
		var _6___mcc_h4 _dafny.Int = _source0.Get_().(D0_DC1).Cf5
		_ = _6___mcc_h4
		var _7_cf5 _dafny.Int = _6___mcc_h4
		_ = _7_cf5
		var _8_cf4 _dafny.Int = _5___mcc_h3
		_ = _8_cf4
		var _9_cf3 bool = _4___mcc_h2
		_ = _9_cf3
		var _10_cf2 _dafny.Map = _3___mcc_h1
		_ = _10_cf2
		var _11_cf1 bool = _2___mcc_h0
		_ = _11_cf1
		return _dafny.CodePoint('p')
	} else if _source0.Is_DC2() {
		var _12___mcc_h5 _dafny.Int = _source0.Get_().(D0_DC2).Cf6
		_ = _12___mcc_h5
		var _13___mcc_h6 bool = _source0.Get_().(D0_DC2).Cf7
		_ = _13___mcc_h6
		var _14___mcc_h7 bool = _source0.Get_().(D0_DC2).Cf8
		_ = _14___mcc_h7
		var _15___mcc_h8 _dafny.Int = _source0.Get_().(D0_DC2).Cf9
		_ = _15___mcc_h8
		var _16_cf9 _dafny.Int = _15___mcc_h8
		_ = _16_cf9
		var _17_cf8 bool = _14___mcc_h7
		_ = _17_cf8
		var _18_cf7 bool = _13___mcc_h6
		_ = _18_cf7
		var _19_cf6 _dafny.Int = _12___mcc_h5
		_ = _19_cf6
		return _dafny.CodePoint('e')
	} else if _source0.Is_DC0() {
		var _20___mcc_h9 bool = _source0.Get_().(D0_DC0).Cf0
		_ = _20___mcc_h9
		var _21_cf0 bool = _20___mcc_h9
		_ = _21_cf0
		return _dafny.CodePoint('u')
	} else {
		var _22___mcc_h10 D0 = _source0.Get_().(D0_DC3).Cf10
		_ = _22___mcc_h10
		var _23_cf10 D0 = _22___mcc_h10
		_ = _23_cf10
		return _dafny.CodePoint('i')
	}
}
func (_static *CompanionStruct_Default___) M0(p0 bool, globalState *GlobalState) {
	var _24_v0 _dafny.Sequence
	_ = _24_v0
	_24_v0 = _dafny.UnicodeSeqOfUtf8Bytes("i")
	_24_v0 = _24_v0
	var _25_v2 _dafny.Int
	_ = _25_v2
	_25_v2 = _dafny.IntOfInt64(795)
	var _26_v3 _dafny.Map
	_ = _26_v3
	_26_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v2, _25_v2)
	var _27_v4 _dafny.Map
	_ = _27_v4
	_27_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v3, true)
	var _28_v5 _dafny.Set
	_ = _28_v5
	_28_v5 = _dafny.SetOf(_25_v2, (_27_v4).Cardinality())
	var _29_v6 _dafny.MultiSet
	_ = _29_v6
	_29_v6 = _dafny.MultiSetOf(_25_v2, _25_v2, Companion_Default___.Fm3(_25_v2, p0, p0, globalState), _dafny.IntOfUint32((_24_v0).Cardinality()))
	_24_v0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_28_v5).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _30_v1 _dafny.Int
			_30_v1 = interface{}(_compr_0).(_dafny.Int)
			if (_28_v5).Contains(_30_v1) {
				_coll0.Add((_30_v1).Plus(_25_v2), p0)
			}
		}
		return _coll0.ToMap()
	}()).Cardinality(), _29_v6, globalState), _24_v0)
	var _31_v7 _dafny.Array
	_ = _31_v7
	var _nwElement0_0 bool = p0
	_ = _nwElement0_0
	var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(11))
	_ = _nw0
	(_nw0).ArraySet1(_nwElement0_0, 0)
	(_nw0).ArraySet1(p0, 1)
	(_nw0).ArraySet1(!(p0), 2)
	(_nw0).ArraySet1(p0, 3)
	(_nw0).ArraySet1(p0, 4)
	(_nw0).ArraySet1(p0, 5)
	(_nw0).ArraySet1(p0, 6)
	(_nw0).ArraySet1((p0) == (p0), 7)
	(_nw0).ArraySet1(p0, 8)
	(_nw0).ArraySet1(p0, 9)
	(_nw0).ArraySet1(p0, 10)
	_31_v7 = _nw0
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))
	_ = _index0
	(_31_v7).ArraySet1(p0, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))
	_ = _index1
	(_31_v7).ArraySet1(!(((_25_v2).Cmp(_25_v2) < 0) || (_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(p0), Companion_Default___.Fm4(_25_v2, globalState)))), (_index1).Int())
	var _32_v8 _dafny.Array
	_ = _32_v8
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
	_ = _nw1
	_32_v8 = _nw1
	var _33_v9 _dafny.Map
	_ = _33_v9
	_33_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v2, (_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool))
	var _34_v10 _dafny.Sequence
	_ = _34_v10
	_34_v10 = _dafny.SeqOf((_33_v9).Cardinality())
	var _35_v11 _dafny.Sequence
	_ = _35_v11
	_35_v11 = _dafny.SeqOf((_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool))
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))
	_ = _index2
	(_32_v8).ArraySet1((_34_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_35_v11).Cardinality()), _dafny.IntOfUint32((_34_v10).Cardinality()))).Uint32()).(_dafny.Int), (_index2).Int())
	var _36_v12 D0
	_ = _36_v12
	_36_v12 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_24_v0).Cardinality()), (_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool), p0, (_33_v9).Cardinality())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))
	_ = _index3
	(_32_v8).ArraySet1((_36_v12).Dtor_cf6(), (_index3).Int())
	var _37_i0 _dafny.Int
	_ = _37_i0
	_37_i0 = _dafny.Zero
	{
		for p0 {
			{
				if (_37_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_37_i0 = (_37_i0).Plus(_dafny.One)
				var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))
				_ = _index4
				(_32_v8).ArraySet1((_dafny.Zero).Minus((((_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_24_v0).Cardinality()))).Plus(_25_v2)), (_index4).Int())
				var _38_v13 _dafny.CodePoint
				_ = _38_v13
				_38_v13 = _dafny.CodePoint('b')
				var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))
				_ = _index5
				(_31_v7).ArraySet1(!(Companion_Default___.Fm0(p0, Companion_Default___.Fm5((_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int), globalState), _38_v13, _25_v2, globalState)), (_index5).Int())
				(globalState).F5 = _25_v2
				var _39_v14 _dafny.Array
				_ = _39_v14
				var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
				_ = _nw2
				_39_v14 = _nw2
				var _40_v15 D1
				_ = _40_v15
				_40_v15 = Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(845))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_41_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				})))
				var _42_v16 D1
				_ = _42_v16
				_42_v16 = Companion_D1_.Create_DC6_(_40_v15)
				var _43_v17 D1
				_ = _43_v17
				_43_v17 = Companion_D1_.Create_DC6_(_42_v16)
				var _44_v18 D1
				_ = _44_v18
				_44_v18 = Companion_D1_.Create_DC6_(_42_v16)
				var _45_v19 D1
				_ = _45_v19
				_45_v19 = Companion_D1_.Create_DC6_(_42_v16)
				var _46_v20 _dafny.Map
				_ = _46_v20
				_46_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_45_v19, (_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool))
				var _47_v21 _dafny.Map
				_ = _47_v21
				_47_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v20, Companion_Default___.Fm3(_dafny.IntOfInt64(172), false, p0, globalState))
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_39_v14), 0))
				_ = _index6
				(_39_v14).ArraySet1((_47_v21).Merge(_47_v21), (_index6).Int())
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_39_v14), 0))
				_ = _index7
				(_39_v14).ArraySet1(_47_v21, (_index7).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _48_v22 _dafny.CodePoint
	_ = _48_v22
	_48_v22 = _dafny.CodePoint('b')
	var _49_i2 _dafny.Int
	_ = _49_i2
	_49_i2 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.IsProperPrefixOf(_24_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_24_v0, _24_v0), (Companion_Default___.SafeIndex((_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_24_v0, _24_v0)).Cardinality()))).Uint32(), _48_v22)) {
			{
				if (_49_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_49_i2 = (_49_i2).Plus(_dafny.One)
				var _50_v24 _dafny.Array
				_ = _50_v24
				var _nwElement0_1 _dafny.Map = _33_v9
				_ = _nwElement0_1
				var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(16))
				_ = _nw3
				(_nw3).ArraySet1(_nwElement0_1, 0)
				(_nw3).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int), (_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool)), 1)
				(_nw3).ArraySet1(_33_v9, 2)
				(_nw3).ArraySet1(_33_v9, 3)
				(_nw3).ArraySet1(_33_v9, 4)
				(_nw3).ArraySet1(_33_v9, 5)
				(_nw3).ArraySet1(func() _dafny.Map {
					var _coll1 = _dafny.NewMapBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate((_26_v3).Keys().Elements()); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _51_v23 _dafny.Int
						_51_v23 = interface{}(_compr_1).(_dafny.Int)
						if (_26_v3).Contains(_51_v23) {
							_coll1.Add((_51_v23).Minus((_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int)), true)
						}
					}
					return _coll1.ToMap()
				}(), 6)
				(_nw3).ArraySet1(_33_v9, 7)
				(_nw3).ArraySet1(_33_v9, 8)
				(_nw3).ArraySet1(_33_v9, 9)
				(_nw3).ArraySet1(_33_v9, 10)
				(_nw3).ArraySet1(_33_v9, 11)
				(_nw3).ArraySet1(_33_v9, 12)
				(_nw3).ArraySet1(_33_v9, 13)
				(_nw3).ArraySet1(_33_v9, 14)
				(_nw3).ArraySet1(_33_v9, 15)
				_50_v24 = _nw3
				var _52_v25 _dafny.Map
				_ = _52_v25
				_52_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v24, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, p0), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_34_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_34_v10).Cardinality()), _dafny.IntOfUint32((_34_v10).Cardinality()))).Uint32(), (_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality()))).Uint32(), (_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool)))
				_52_v25 = (_52_v25).Update(_50_v24, _35_v11)
				var _53_v26 _dafny.Map
				_ = _53_v26
				_53_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_32_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(533), _dafny.ArrayLen((_32_v8), 0))).Int()).(_dafny.Int), _24_v0)
				var _54_v27 _dafny.Map
				_ = _54_v27
				_54_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_53_v26).Cardinality())
				_54_v27 = _54_v27
				(globalState).F0 = !(p0)
				var _55_v28 *C0
				_ = _55_v28
				var _nw4 *C0 = New_C0_()
				_ = _nw4
				_nw4.Ctor__(_25_v2, _dafny.UnicodeSeqOfUtf8Bytes("bxfg"))
				_55_v28 = _nw4
				var _56_v29 D1
				_ = _56_v29
				_56_v29 = Companion_D1_.Create_DC5_((_31_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(469), _dafny.ArrayLen((_31_v7), 0))).Int()).(bool), _55_v28, p0)
				(globalState).F7 = (_56_v29).Dtor_cf14()
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _57_v0 _dafny.Sequence
	_ = _57_v0
	_57_v0 = _dafny.UnicodeSeqOfUtf8Bytes("sraobfhjd")
	var _58_v1 _dafny.Int
	_ = _58_v1
	_58_v1 = _dafny.IntOfInt64(178)
	var _59_v2 _dafny.CodePoint
	_ = _59_v2
	_59_v2 = _dafny.CodePoint('q')
	var _60_v3 _dafny.Array
	_ = _60_v3
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_0
	var _nw5 _dafny.Array
	_ = _nw5
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw5 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Map = (func(_61_v1 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_62_i0 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC2_(_61_v1, true, true, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Dtor_cf7(), _61_v1)
			}
		})(_58_v1)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw5 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw5).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw5).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_60_v3 = _nw5
	var _63_v4 _dafny.Array
	_ = _63_v4
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
	_ = _len0_1
	var _nw6 _dafny.Array
	_ = _nw6
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw6 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Set = (func(_64_v1 _dafny.Int) func(_dafny.Int) _dafny.Set {
			return func(_65_i1 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_64_v1)
			}
		})(_58_v1)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw6 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw6).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw6).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_63_v4 = _nw6
	var _66_v5 _dafny.Array
	_ = _66_v5
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_2
	var _nw7 _dafny.Array
	_ = _nw7
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw7 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Int = func(_67_i2 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeModuloInt(_67_i2, _dafny.IntOfInt64(855))
		}
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw7 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw7).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw7).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_66_v5 = _nw7
	var _68_v6 _dafny.Array
	_ = _68_v6
	var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
	_ = _nw8
	_68_v6 = _nw8
	var _69_v7 bool
	_ = _69_v7
	_69_v7 = false
	var _70_v8 _dafny.Set
	_ = _70_v8
	_70_v8 = _dafny.SetOf(_59_v2, _59_v2, _dafny.CodePoint('u'), _59_v2)
	var _71_globalState *GlobalState
	_ = _71_globalState
	var _nw9 *GlobalState = New_GlobalState_()
	_ = _nw9
	_nw9.Ctor__(false, _dafny.IntOfInt64(237), _dafny.Companion_Sequence_.Update(_57_v0, (Companion_Default___.SafeIndex(_58_v1, _dafny.IntOfUint32((_57_v0).Cardinality()))).Uint32(), _59_v2), _60_v3, _dafny.CodePoint('w'), _dafny.IntOfInt64(-46), _63_v4, false, _66_v5, _68_v6, _66_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v7, (_dafny.Zero).Minus(_58_v1)), false, true, _70_v8, true)
	_71_globalState = _nw9
	_59_v2 = _59_v2
	var _72_v9 _dafny.MultiSet
	_ = _72_v9
	_72_v9 = _dafny.MultiSetOf(_69_v7)
	var _hi0 _dafny.Int = (func() _dafny.Int {
		if (_72_v9).Contains(_69_v7) {
			return (_72_v9).Multiplicity(_69_v7)
		}
		return _58_v1
	})()
	_ = _hi0
	for _73_i3 := _58_v1; _73_i3.Cmp(_hi0) < 0; _73_i3 = _73_i3.Plus(_dafny.One) {
		var _74_v10 D0
		_ = _74_v10
		_74_v10 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(998), _69_v7, _69_v7, _dafny.IntOfUint32((_57_v0).Cardinality()))
		var _pat_let_tv0 = _57_v0
		_ = _pat_let_tv0
		if (func() bool {
			if _69_v7 {
				return (_73_i3).Cmp(_58_v1) >= 0
			}
			return (func(_pat_let0_0 D0) D0 {
				return func(_75_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 _dafny.Int) D0 {
						return func(_76_dt__update_hcf9_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC2_((_75_dt__update__tmp_h0).Dtor_cf6(), (_75_dt__update__tmp_h0).Dtor_cf7(), (_75_dt__update__tmp_h0).Dtor_cf8(), _76_dt__update_hcf9_h0)
						}(_pat_let1_0)
					}((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_i3, _dafny.IntOfUint32((_pat_let_tv0).Cardinality()))).Cardinality())
				}(_pat_let0_0)
			}(_74_v10)).Dtor_cf8()
		})() {
			var _77_v11 _dafny.Array
			_ = _77_v11
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_3
			var _nw10 _dafny.Array
			_ = _nw10
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw10 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = func(_78_i4 _dafny.Int) bool {
					return !(!(!(true)))
				}
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw10 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw10).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw10).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_77_v11 = _nw10
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_77_v11), 0))
			_ = _index8
			(_77_v11).ArraySet1(((_72_v9).Cardinality()).Cmp(_73_i3) > 0, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_77_v11), 0))
			_ = _index9
			(_77_v11).ArraySet1(!(_69_v7) || (_69_v7), (_index9).Int())
			var _79_v12 _dafny.Sequence
			_ = _79_v12
			_79_v12 = _dafny.SeqOf(_66_v5)
			var _80_v13 _dafny.Map
			_ = _80_v13
			_80_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_77_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_77_v11), 0))).Int()).(bool), _69_v7)
			var _81_v14 _dafny.Array
			_ = _81_v14
			var _nwElement0_2 _dafny.Array = _66_v5
			_ = _nwElement0_2
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(5))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_2, 0)
			(_nw11).ArraySet1(_66_v5, 1)
			(_nw11).ArraySet1(_66_v5, 2)
			(_nw11).ArraySet1((_79_v12).Select((Companion_Default___.SafeIndex((_80_v13).Cardinality(), _dafny.IntOfUint32((_79_v12).Cardinality()))).Uint32()).(_dafny.Array), 3)
			(_nw11).ArraySet1(_66_v5, 4)
			_81_v14 = _nw11
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_81_v14), 0))
			_ = _index10
			(_81_v14).ArraySet1(_66_v5, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_81_v14), 0))
			_ = _index11
			(_81_v14).ArraySet1(_66_v5, (_index11).Int())
			var _82_v15 _dafny.Map
			_ = _82_v15
			_82_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_66_v5, !(true) || (_69_v7))
			var _83_v16 _dafny.Map
			_ = _83_v16
			_83_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v7, _66_v5)
			_82_v15 = (_82_v15).Update((func() _dafny.Array {
				if (_83_v16).Contains((_77_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_77_v11), 0))).Int()).(bool)) {
					return (_83_v16).Get((_77_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_77_v11), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _66_v5
			})(), (_73_i3).Cmp(_73_i3) > 0)
			(_71_globalState).F5 = (func() _dafny.Int {
				if (_77_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_77_v11), 0))).Int()).(bool) {
					return _58_v1
				}
				return _dafny.IntOfInt64(823)
			})()
			(_71_globalState).F5 = (_58_v1).Minus(_73_i3)
		} else {
			Companion_Default___.M0(_69_v7, _71_globalState)
			(_71_globalState).F12 = Companion_Default___.Fm0((_69_v7) || (_69_v7), _59_v2, _59_v2, _58_v1, _71_globalState)
			_59_v2 = _59_v2
			_58_v1 = _dafny.IntOfUint32((_57_v0).Cardinality())
			(_71_globalState).F0 = true
		}
		var _84_v17 _dafny.Map
		_ = _84_v17
		_84_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _69_v7)
		var _85_v18 _dafny.Map
		_ = _85_v18
		_85_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v17, _57_v0)
		(_71_globalState).F5 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_85_v18).Contains(_84_v17) {
				return (_85_v18).Get(_84_v17).(_dafny.Sequence)
			}
			return _57_v0
		})(), (func() _dafny.Sequence {
			if _69_v7 {
				return _57_v0
			}
			return _57_v0
		})())).Cardinality())
		var _86_v19 *C0
		_ = _86_v19
		var _nw12 *C0 = New_C0_()
		_ = _nw12
		_nw12.Ctor__(_73_i3, _57_v0)
		_86_v19 = _nw12
		(_71_globalState).F0 = (func() bool {
			if _69_v7 {
				return _69_v7
			}
			return _69_v7
		})()
	}
	Companion_Default___.M0(_69_v7, _71_globalState)
	var _hi1 _dafny.Int = _58_v1
	_ = _hi1
	for _87_i5 := (_dafny.Zero).Minus((_58_v1).Plus(_58_v1)); _87_i5.Cmp(_hi1) < 0; _87_i5 = _87_i5.Plus(_dafny.One) {
		var _rhs0 _dafny.Array = _66_v5
		_ = _rhs0
		var _rhs1 _dafny.CodePoint = _59_v2
		_ = _rhs1
		var _rhs2 bool = !(_69_v7) || (_69_v7)
		_ = _rhs2
		var _rhs3 bool = _69_v7
		_ = _rhs3
		var _lhs0 *GlobalState = _71_globalState
		_ = _lhs0
		var _lhs1 *GlobalState = _71_globalState
		_ = _lhs1
		var _lhs2 *GlobalState = _71_globalState
		_ = _lhs2
		_lhs0.F10 = _rhs0
		_59_v2 = _rhs1
		_lhs1.F12 = _rhs2
		_lhs2.F12 = _rhs3
		var _88_v20 *C0
		_ = _88_v20
		var _nw13 *C0 = New_C0_()
		_ = _nw13
		_nw13.Ctor__(_87_i5, _dafny.Companion_Sequence_.Concatenate(_57_v0, _57_v0))
		_88_v20 = _nw13
		var _89_v21 _dafny.Sequence
		_ = _89_v21
		_89_v21 = _dafny.SeqOf((_88_v20).F16())
		_89_v21 = (func() _dafny.Sequence {
			if _69_v7 {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(346))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}((func(_90_v20 *C0) func(_dafny.Int) _dafny.Int {
					return func(_91_i6 _dafny.Int) _dafny.Int {
						return (_90_v20).F16()
					}
				})(_88_v20)))
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(801))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}((func(_92_i5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_93_i7 _dafny.Int) _dafny.Int {
					return _92_i5
				}
			})(_87_i5)))
		})()
		var _94_v22 D0
		_ = _94_v22
		_94_v22 = Companion_D0_.Create_DC0_(_69_v7)
		var _95_v23 _dafny.Map
		_ = _95_v23
		_95_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v22, (_88_v20).F17())
		var _96_v24 D1
		_ = _96_v24
		_96_v24 = Companion_D1_.Create_DC4_((_88_v20).F17())
		_95_v23 = (_95_v23).Update(Companion_D0_.Create_DC0_(!(_69_v7)), _dafny.Companion_Sequence_.Concatenate((_96_v24).Dtor_cf11(), _57_v0))
	}
	var _97_v25 _dafny.MultiSet
	_ = _97_v25
	_97_v25 = _dafny.MultiSetOf(_58_v1)
	(_71_globalState).F7 = _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm2(_58_v1, _97_v25, _71_globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-49))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}((func(_98_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_99_i8 _dafny.Int) _dafny.CodePoint {
			return _98_v2
		}
	})(_59_v2))))
	var _100_v26 _dafny.Set
	_ = _100_v26
	_100_v26 = _dafny.SetOf(_69_v7, _69_v7, _69_v7)
	var _101_v27 _dafny.Map
	_ = _101_v27
	_101_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v26, Companion_Default___.Fm0(_69_v7, _59_v2, _59_v2, (_dafny.Zero).Minus(_58_v1), _71_globalState))
	var _hi2 _dafny.Int = _dafny.IntOfInt64(103)
	_ = _hi2
	for _102_i9 := (_101_v27).Cardinality(); _102_i9.Cmp(_hi2) < 0; _102_i9 = _102_i9.Plus(_dafny.One) {
		var _103_v28 _dafny.Array
		_ = _103_v28
		var _nw14 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
		_ = _nw14
		_103_v28 = _nw14
		var _104_v29 *C0
		_ = _104_v29
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__(_102_i9, _dafny.UnicodeSeqOfUtf8Bytes("hrcnf"))
		_104_v29 = _nw15
		var _105_v30 D1
		_ = _105_v30
		_105_v30 = Companion_D1_.Create_DC5_(_69_v7, _104_v29, _69_v7)
		var _pat_let_tv1 = _69_v7
		_ = _pat_let_tv1
		var _pat_let_tv2 = _69_v7
		_ = _pat_let_tv2
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_103_v28), 0))
		_ = _index12
		(_103_v28).ArraySet1(func(_pat_let2_0 D1) D1 {
			return func(_106_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let3_0 bool) D1 {
					return func(_107_dt__update_hcf14_h0 bool) D1 {
						return func(_pat_let4_0 bool) D1 {
							return func(_108_dt__update_hcf12_h0 bool) D1 {
								return Companion_D1_.Create_DC5_(_108_dt__update_hcf12_h0, (_106_dt__update__tmp_h1).Dtor_cf13(), _107_dt__update_hcf14_h0)
							}(_pat_let4_0)
						}(_pat_let_tv2)
					}(_pat_let3_0)
				}(_pat_let_tv1)
			}(_pat_let2_0)
		}(_105_v30), (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_103_v28), 0))
		_ = _index13
		(_103_v28).ArraySet1(_105_v30, (_index13).Int())
		var _109_v31 _dafny.Set
		_ = _109_v31
		_109_v31 = _dafny.SetOf(_104_v29)
		var _rhs4 _dafny.Int = (_dafny.Zero).Minus((_104_v29).F16())
		_ = _rhs4
		var _rhs5 _dafny.Int = ((_109_v31).Union(_109_v31)).Cardinality()
		_ = _rhs5
		var _lhs3 *GlobalState = _71_globalState
		_ = _lhs3
		var _lhs4 *GlobalState = _71_globalState
		_ = _lhs4
		_lhs3.F5 = _rhs4
		_lhs4.F5 = _rhs5
		var _110_v32 D1
		_ = _110_v32
		_110_v32 = Companion_D1_.Create_DC4_((_104_v29).F17())
		var _111_v33 D1
		_ = _111_v33
		_111_v33 = Companion_D1_.Create_DC6_(_110_v32)
		_111_v33 = Companion_D1_.Create_DC6_(_110_v32)
		Companion_Default___.M0(_69_v7, _71_globalState)
	}
	var _112_v34 _dafny.Map
	_ = _112_v34
	_112_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _69_v7)
	_112_v34 = (_112_v34).Update(false, _69_v7)
	var _113_v35 *C0
	_ = _113_v35
	var _nw16 *C0 = New_C0_()
	_ = _nw16
	_nw16.Ctor__(_58_v1, _57_v0)
	_113_v35 = _nw16
	(_71_globalState).F5 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v1, _113_v35)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_58_v1, _113_v35))).Cardinality()
	var _114_v36 _dafny.Map
	_ = _114_v36
	_114_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_113_v35).F16(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}((func(_115_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_116_i10 _dafny.Int) _dafny.CodePoint {
			return _115_v2
		}
	})(_59_v2))), _dafny.UnicodeSeqOfUtf8Bytes("fal")))
	_114_v36 = (_114_v36).Update(_dafny.IntOfInt64(-282), (_113_v35).F17())
	var _117_v37 *C0
	_ = _117_v37
	var _nw17 *C0 = New_C0_()
	_ = _nw17
	_nw17.Ctor__(Companion_Default___.SafeModuloInt((_113_v35).F16(), (_113_v35).F16()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}((func(_118_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_119_i11 _dafny.Int) _dafny.CodePoint {
			return _118_v2
		}
	})(_59_v2))), (_113_v35).F17()))
	_117_v37 = _nw17
	Companion_Default___.M0(((_117_v37).F16()).Cmp((_113_v35).F16()) != 0, _71_globalState)
	var _120_v38 _dafny.Map
	_ = _120_v38
	_120_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v2, _dafny.IntOfUint32((_57_v0).Cardinality()))
	_120_v38 = (func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(_59_v2)).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _121_v39 _dafny.CodePoint
			_121_v39 = interface{}(_compr_2).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_59_v2), _121_v39) {
				_coll2.Add(_121_v39, (_dafny.SetOf((_117_v37).F16(), (_113_v35).F16(), (_117_v37).F16(), (_112_v34).Cardinality(), _58_v1)).Cardinality())
			}
		}
		return _coll2.ToMap()
	}()).Merge(_120_v38)
	_69_v7 = _69_v7
	(_71_globalState).F5 = (_117_v37).F16()
	Companion_Default___.M0(_69_v7, _71_globalState)
	var _122_v40 _dafny.Map
	_ = _122_v40
	_122_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_v7, _113_v35)
	(_71_globalState).F5 = (((_117_v37).F16()).Times(((Companion_D2_.Create_DC7_(_122_v40)).Dtor_cf16()).Cardinality())).Plus(_dafny.IntOfInt64(255))
	_dafny.Print(_57_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_58_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_59_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_60_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_63_v4).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_66_v5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_69_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_70_v8).Equals(_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('u'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F2().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_71_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.One).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState.F6).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Set)).Equals(_dafny.SetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F8()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F10).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_71_globalState).F11()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_71_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState.F14).Equals(_dafny.SetOf(_dafny.CodePoint('q'), _dafny.CodePoint('u'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_71_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_72_v9).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v25).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(178))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v26).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v27).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_112_v34).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v35).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v35).F17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v36).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(178), _dafny.UnicodeSeqOfUtf8Bytes("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqfal")).UpdateUnsafe(_dafny.IntOfInt64(-282), _dafny.UnicodeSeqOfUtf8Bytes("sraobfhjd"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_117_v37).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_117_v37).F17().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v38).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v40).Cardinality())
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
	Cf1 bool
	Cf2 _dafny.Map
	Cf3 bool
	Cf4 _dafny.Int
	Cf5 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Map, Cf3 bool, Cf4 _dafny.Int, Cf5 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf6 _dafny.Int
	Cf7 bool
	Cf8 bool
	Cf9 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf6 _dafny.Int, Cf7 bool, Cf8 bool, Cf9 _dafny.Int) D0 {
	return D0{D0_DC2{Cf6, Cf7, Cf8, Cf9}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
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

type D0_DC3 struct {
	Cf10 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf10 D0) D0 {
	return D0{D0_DC3{Cf10}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.EmptyMap, false, _dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Map {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() bool {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf8() bool {
	return _this.Get_().(D0_DC2).Cf8
}

func (_this D0) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf9
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf10() D0 {
	return _this.Get_().(D0_DC3).Cf10
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2.Equals(data2.Cf2) && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D1_DC5 struct {
	Cf12 bool
	Cf13 *C0
	Cf14 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf12 bool, Cf13 *C0, Cf14 bool) D1 {
	return D1{D1_DC5{Cf12, Cf13, Cf14}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf11 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf11 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC6 struct {
	Cf15 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf15 D1) D1 {
	return D1{D1_DC6{Cf15}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false, (*C0)(nil), false)
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf13() *C0 {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf14() bool {
	return _this.Get_().(D1_DC5).Cf14
}

func (_this D1) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf15() D1 {
	return _this.Get_().(D1_DC6).Cf15
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + data.Cf11.VerbatimString(true) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf15.Equals(data2.Cf15)
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

type D2_DC8 struct {
	Cf17 _dafny.Int
	Cf18 _dafny.Int
	Cf19 *C0
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf17 _dafny.Int, Cf18 _dafny.Int, Cf19 *C0) D2 {
	return D2{D2_DC8{Cf17, Cf18, Cf19}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf16 _dafny.Map
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf16 _dafny.Map) D2 {
	return D2{D2_DC7{Cf16}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.Zero, _dafny.Zero, (*C0)(nil))
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf18
}

func (_this D2) Dtor_cf19() *C0 {
	return _this.Get_().(D2_DC8).Cf19
}

func (_this D2) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf16
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

// Definition of class GlobalState
type GlobalState struct {
	F0   bool
	F5   _dafny.Int
	F6   _dafny.Array
	F7   bool
	F10  _dafny.Array
	F12  bool
	F14  _dafny.Set
	_f1  _dafny.Int
	_f2  _dafny.Sequence
	_f3  _dafny.Array
	_f4  _dafny.CodePoint
	_f8  _dafny.Array
	_f9  _dafny.Array
	_f11 _dafny.Map
	_f13 bool
	_f15 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this.F5 = _dafny.Zero
	_this.F6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F7 = false
	_this.F10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F12 = false
	_this.F14 = _dafny.EmptySet
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.EmptySeq
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f4 = _dafny.CodePoint('D')
	_this._f8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f9 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = _dafny.EmptyMap
	_this._f13 = false
	_this._f15 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Sequence, f3 _dafny.Array, f4 _dafny.CodePoint, f5 _dafny.Int, f6 _dafny.Array, f7 bool, f8 _dafny.Array, f9 _dafny.Array, f10 _dafny.Array, f11 _dafny.Map, f12 bool, f13 bool, f14 _dafny.Set, f15 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Sequence {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.CodePoint {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F8() _dafny.Array {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Array {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F11() _dafny.Map {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f16 _dafny.Int
	_f17 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f16 = _dafny.Zero
	_this._f17 = _dafny.EmptySeq
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

func (_this *C0) Ctor__(f16 _dafny.Int, f17 _dafny.Sequence) {
	{
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *C0) Fm1(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32(((func() _dafny.Sequence {
			if !((_dafny.MultiSetOf((Companion_D0_.Create_DC1_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F16()), true, (_this).F16(), (_this).F16())).Dtor_cf1())).IsSubsetOf(_dafny.MultiSetOf(false, true))) {
				return (_this).F17()
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(552))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_123_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			}))
		})()).Cardinality())
	}
}
func (_this *C0) F16() _dafny.Int {
	{
		return _this._f16
	}
}
func (_this *C0) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
