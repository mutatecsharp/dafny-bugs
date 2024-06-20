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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfUint32((_dafny.SeqOf(!(!(!(false))), true, true, false, false)).Cardinality())).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tialul"), _dafny.UnicodeSeqOfUtf8Bytes("olu"))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Sequence, p1 _dafny.Set, p2 bool, globalState *GlobalState) bool {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true, true)).Cardinality(), true)).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(552), _dafny.IntOfInt64(799))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(552)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(799)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_0_v0, (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("vi"))).Cardinality()), (Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("nqxdlqnf"), (Companion_D0_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg0 _dafny.Int) interface{} {
						return coer0(arg0)
					}
				}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				})), true, true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Dtor_cf3(), false, _dafny.IntOfInt64(326))).Dtor_cf3())
			}
		}
		return _coll0.ToMap()
	}())).Contains(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(859), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(246), _dafny.IntOfInt64(396))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('b')
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D0 = Companion_D0_.Create_DC2_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(118))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), !(false), _dafny.SeqOf(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(948), true)).Cardinality(), _dafny.CodePoint('v'))
	_ = _source0
	if _source0.Is_DC1() {
		var _3___mcc_h0 _dafny.Sequence = _source0.Get_().(D0_DC1).Cf1
		_ = _3___mcc_h0
		var _4___mcc_h1 bool = _source0.Get_().(D0_DC1).Cf2
		_ = _4___mcc_h1
		var _5___mcc_h2 bool = _source0.Get_().(D0_DC1).Cf3
		_ = _5___mcc_h2
		var _6___mcc_h3 _dafny.Int = _source0.Get_().(D0_DC1).Cf4
		_ = _6___mcc_h3
		var _7_cf4 _dafny.Int = _6___mcc_h3
		_ = _7_cf4
		var _8_cf3 bool = _5___mcc_h2
		_ = _8_cf3
		var _9_cf2 bool = _4___mcc_h1
		_ = _9_cf2
		var _10_cf1 _dafny.Sequence = _3___mcc_h0
		_ = _10_cf1
		if _8_cf3 {
			return _dafny.SeqOf(_7_cf4)
		} else {
			return _dafny.SeqOf(_7_cf4, _7_cf4)
		}
	} else if _source0.Is_DC2() {
		var _11___mcc_h4 _dafny.Sequence = _source0.Get_().(D0_DC2).Cf5
		_ = _11___mcc_h4
		var _12___mcc_h5 bool = _source0.Get_().(D0_DC2).Cf6
		_ = _12___mcc_h5
		var _13___mcc_h6 _dafny.Sequence = _source0.Get_().(D0_DC2).Cf7
		_ = _13___mcc_h6
		var _14___mcc_h7 _dafny.Int = _source0.Get_().(D0_DC2).Cf8
		_ = _14___mcc_h7
		var _15___mcc_h8 _dafny.CodePoint = _source0.Get_().(D0_DC2).Cf9
		_ = _15___mcc_h8
		var _16_cf9 _dafny.CodePoint = _15___mcc_h8
		_ = _16_cf9
		var _17_cf8 _dafny.Int = _14___mcc_h7
		_ = _17_cf8
		var _18_cf7 _dafny.Sequence = _13___mcc_h6
		_ = _18_cf7
		var _19_cf6 bool = _12___mcc_h5
		_ = _19_cf6
		var _20_cf5 _dafny.Sequence = _11___mcc_h4
		_ = _20_cf5
		return _dafny.SeqOf(_17_cf8)
	} else if _source0.Is_DC3() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(855), !(false))).Cardinality(), _dafny.IntOfInt64(794), _dafny.IntOfInt64(211), _dafny.IntOfInt64(165), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hrwsevras")).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqrxt")).Cardinality())))
	} else if _source0.Is_DC0() {
		var _21___mcc_h9 _dafny.Int = _source0.Get_().(D0_DC0).Cf0
		_ = _21___mcc_h9
		var _22_cf0 _dafny.Int = _21___mcc_h9
		_ = _22_cf0
		return _dafny.SeqOf(_dafny.IntOfInt64(734))
	} else {
		var _23___mcc_h10 D0 = _source0.Get_().(D0_DC4).Cf10
		_ = _23___mcc_h10
		var _24_cf10 D0 = _23___mcc_h10
		_ = _24_cf10
		return _dafny.SeqOf(_dafny.IntOfInt64(911), _dafny.IntOfInt64(62))
	}
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_25_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), !(false) || (true), !((_dafny.MultiSetOf(true)).IsSubsetOf(_dafny.MultiSetOf(true, false))), _dafny.IntOfInt64(895))
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(438))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_26_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality()))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-401))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(138), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(889))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_27_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality()))).Cardinality())).Cardinality(), _dafny.IntOfInt64(556))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_28_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality()))).Union(_dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(-292))).Cardinality())))).Difference(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("my")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if false {
			return func() _dafny.Set {
				var _coll1 = _dafny.NewBuilder()
				_ = _coll1
				for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("s"))).Elements()); ; {
					_compr_1, _ok1 := _iter1()
					if !_ok1 {
						break
					}
					var _29_v0 _dafny.Sequence
					_29_v0 = interface{}(_compr_1).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("s")), _29_v0) {
						_coll1.Add(_29_v0)
					}
				}
				return _coll1.ToSet()
			}()
		}
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ncxyqxudm"))
	})()).Union((func() _dafny.Set {
		if false {
			return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("pyfiuigns"))
		}
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("wruwanx"), _dafny.UnicodeSeqOfUtf8Bytes("mpqo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(137))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_30_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})))
	})())
}
func (_static *CompanionStruct_Default___) M0(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) {
	var _31_v0 _dafny.Map
	_ = _31_v0
	_31_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
	var _32_v1 _dafny.MultiSet
	_ = _32_v1
	_32_v1 = _dafny.MultiSetOf(!(_31_v0).Equals(_31_v0), !(p0))
	(globalState).F4 = (func() _dafny.Int {
		if (_32_v1).Contains(p0) {
			return (_32_v1).Multiplicity(p0)
		}
		return Companion_Default___.Fm0((_dafny.Zero).Minus(p3), p3, p0, globalState)
	})()
	var _33_v2 _dafny.Array
	_ = _33_v2
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
	_ = _nw0
	_33_v2 = _nw0
	var _34_v3 _dafny.CodePoint
	_ = _34_v3
	_34_v3 = _dafny.CodePoint('s')
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_33_v2), 0))
	_ = _index0
	(_33_v2).ArraySet1CodePoint(_34_v3, (_index0).Int())
	var _35_v4 _dafny.MultiSet
	_ = _35_v4
	_35_v4 = _dafny.MultiSetOf(_dafny.IntOfInt64(163), p1)
	var _36_v5 *C0
	_ = _36_v5
	var _nw1 *C0 = New_C0_()
	_ = _nw1
	_nw1.Ctor__()
	_36_v5 = _nw1
	var _37_v6 D4
	_ = _37_v6
	_37_v6 = Companion_D4_.Create_DC12_(_34_v3, _35_v4, p0, p1, _36_v5)
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_33_v2), 0))
	_ = _index1
	(_33_v2).ArraySet1CodePoint((_37_v6).Dtor_cf16(), (_index1).Int())
	var _38_v7 _dafny.Array
	_ = _38_v7
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_0
	var _nw2 _dafny.Array
	_ = _nw2
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw2 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = func(_39_i0 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeDivisionInt(_39_i0, _dafny.IntOfInt64(824))
		}
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
	_38_v7 = _nw2
	_38_v7 = _38_v7
	var _40_v8 _dafny.Map
	_ = _40_v8
	_40_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
	var _41_v9 _dafny.Sequence
	_ = _41_v9
	_41_v9 = _dafny.SeqOf(p0)
	var _42_v10 _dafny.Map
	_ = _42_v10
	_42_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _41_v9)
	var _43_v11 _dafny.Map
	_ = _43_v11
	_43_v11 = _42_v10
	var _pat_let_tv0 = _40_v8
	_ = _pat_let_tv0
	_40_v8 = func(_source1 _dafny.Map) _dafny.Map {
		var _44___mcc_h0 _dafny.Map = _source1
		_ = _44___mcc_h0
		var _45_cf13 _dafny.Map = _44___mcc_h0
		_ = _45_cf13
		return _pat_let_tv0
	}(_43_v11)
	var _46_v12 _dafny.Sequence
	_ = _46_v12
	_46_v12 = _dafny.SeqOf((_33_v2).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_33_v2), 0))).Int()))
	var _47_v13 _dafny.Sequence
	_ = _47_v13
	_47_v13 = _dafny.SeqOf(_46_v12, _46_v12)
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_47_v13).Cardinality())
	_ = _hi0
	for _48_i1 := _dafny.IntOfUint32((_41_v9).Cardinality()); _48_i1.Cmp(_hi0) < 0; _48_i1 = _48_i1.Plus(_dafny.One) {
		(globalState).F4 = (Companion_Default___.Fm0(p3, _dafny.IntOfUint32((_41_v9).Cardinality()), !(false), globalState)).Times(_48_i1)
		var _49_v14 _dafny.Map
		_ = _49_v14
		_49_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_46_v12, p0)
		var _50_v15 _dafny.Set
		_ = _50_v15
		_50_v15 = _dafny.SetOf(p1, p3)
		(globalState).F1 = (func() bool {
			if (_49_v14).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(624))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_51_v2 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
				return func(_52_i2 _dafny.Int) _dafny.CodePoint {
					return (_51_v2).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_51_v2), 0))).Int())
				}
			})(_33_v2))), _46_v12)) {
				return (_49_v14).Get(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(624))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg8 _dafny.Int) interface{} {
						return coer8(arg8)
					}
				}((func(_53_v2 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_54_i2 _dafny.Int) _dafny.CodePoint {
						return (_53_v2).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_53_v2), 0))).Int())
					}
				})(_33_v2))), _46_v12)).(bool)
			}
			return !(!(_50_v15).Contains(_48_i1))
		})()
		(globalState).F4 = Companion_Default___.Fm0((_dafny.IntOfInt64(501)).Times(p3), _dafny.IntOfInt64(-702), p0, globalState)
		(globalState).F5 = !(p0)
	}
	var _hi1 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hnssgkhut")).Cardinality())
	_ = _hi1
	for _55_i3 := p3; _55_i3.Cmp(_hi1) < 0; _55_i3 = _55_i3.Plus(_dafny.One) {
		(globalState).F4 = p3
		if p0 {
			(globalState).F5 = p0
			var _56_v16 D1
			_ = _56_v16
			_56_v16 = Companion_D1_.Create_DC6_(p1)
			_56_v16 = _56_v16
			var _57_v17 D0
			_ = _57_v17
			_57_v17 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(843))
			var _58_v18 _dafny.Sequence
			_ = _58_v18
			_58_v18 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(545))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_59_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_60_i4 _dafny.Int) _dafny.Int {
					return _59_p3
				}
			})(p3))))
			var _61_v19 _dafny.Sequence
			_ = _61_v19
			_61_v19 = _dafny.SeqOf(_57_v17, Companion_Default___.Fm6(_dafny.IntOfUint32(((_58_v18).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_58_v18).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _40_v8, p0, p0, globalState))
			var _62_v20 _dafny.Array
			_ = _62_v20
			var _nwElement0_0 _dafny.Sequence = _61_v19
			_ = _nwElement0_0
			var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(11))
			_ = _nw3
			(_nw3).ArraySet1(_nwElement0_0, 0)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_61_v19, _61_v19), 1)
			(_nw3).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(709))).Uint32(), func(coer10 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_63_v17 D0) func(_dafny.Int) D0 {
				return func(_64_i5 _dafny.Int) D0 {
					return _63_v17
				}
			})(_57_v17))), 2)
			(_nw3).ArraySet1(_61_v19, 3)
			(_nw3).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_61_v19, _61_v19), 4)
			(_nw3).ArraySet1(_61_v19, 5)
			(_nw3).ArraySet1(_61_v19, 6)
			(_nw3).ArraySet1(_61_v19, 7)
			(_nw3).ArraySet1(_61_v19, 8)
			(_nw3).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(965))).Uint32(), func(coer11 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_65_v17 D0) func(_dafny.Int) D0 {
				return func(_66_i6 _dafny.Int) D0 {
					return _65_v17
				}
			})(_57_v17))), 9)
			(_nw3).ArraySet1(_61_v19, 10)
			_62_v20 = _nw3
			var _67_v21 _dafny.Array
			_ = _67_v21
			var _nw4 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw4
			_67_v21 = _nw4
			var _68_v22 _dafny.Array
			_ = _68_v22
			_68_v22 = _67_v21
			var _rhs0 _dafny.Array = _62_v20
			_ = _rhs0
			var _rhs1 _dafny.Array = (_68_v22)
			_ = _rhs1
			var _rhs2 _dafny.Array = _38_v7
			_ = _rhs2
			_62_v20 = _rhs0
			_67_v21 = _rhs1
			_38_v7 = _rhs2
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_38_v7), 0))
			_ = _index2
			(_38_v7).ArraySet1((_dafny.Zero).Minus(p3), (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_38_v7), 0))
			_ = _index3
			(_38_v7).ArraySet1(Companion_Default___.SafeDivisionInt(p1, _55_i3), (_index3).Int())
			var _69_v23 _dafny.Map
			_ = _69_v23
			_69_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
			var _70_v24 _dafny.Map
			_ = _70_v24
			_70_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(218), p3)
			var _71_v25 _dafny.Set
			_ = _71_v25
			_71_v25 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(625))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_72_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_73_i7 _dafny.Int) _dafny.Int {
					return _72_p3
				}
			})(p3)))).Cardinality()))
			var _74_v26 _dafny.Sequence
			_ = _74_v26
			_74_v26 = _dafny.SeqOf(_36_v5, _36_v5)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_38_v7), 0))
			_ = _index4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_38_v7), 0))
			_ = _index5
			var _rhs3 _dafny.Int = (func() _dafny.Int {
				if (_69_v23).Contains((Companion_Default___.Fm0((func() _dafny.Int {
					if (_70_v24).Contains(p3) {
						return (_70_v24).Get(p3).(_dafny.Int)
					}
					return _55_i3
				})(), p3, p0, globalState)).Cmp(p3) != 0) {
					return (_69_v23).Get((Companion_Default___.Fm0((func() _dafny.Int {
						if (_70_v24).Contains(p3) {
							return (_70_v24).Get(p3).(_dafny.Int)
						}
						return _55_i3
					})(), p3, p0, globalState)).Cmp(p3) != 0).(_dafny.Int)
				}
				return p1
			})()
			_ = _rhs3
			var _rhs4 _dafny.Int = Companion_Default___.SafeModuloInt(p3, Companion_Default___.Fm0((p2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(_dafny.Int), _55_i3, (_37_v6).Dtor_cf18(), globalState))
			_ = _rhs4
			var _rhs5 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_41_v9, _41_v9), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_55_i3), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_41_v9, _41_v9)).Cardinality()))).Uint32(), Companion_Default___.Fm1(p2, _71_v25, p0, globalState)), (_32_v1).IsProperSubsetOf(_32_v1))
			_ = _rhs5
			var _rhs6 _dafny.Int = p3
			_ = _rhs6
			var _rhs7 _dafny.Int = (_dafny.Zero).Minus((_55_i3).Plus(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((_74_v26).Cardinality()))))
			_ = _rhs7
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			var _lhs1 _dafny.Array = _38_v7
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_38_v7), 0))
			_ = _lhs2
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			var _lhs4 _dafny.Array = _38_v7
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_38_v7), 0))
			_ = _lhs5
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			_lhs0.F4 = _rhs3
			(_lhs1).ArraySet1(_rhs4, (_lhs2).Int())
			_lhs3.F1 = _rhs5
			(_lhs4).ArraySet1(_rhs6, (_lhs5).Int())
			_lhs6.F4 = _rhs7
			(globalState).F4 = (Companion_D4_.Create_DC12_((_33_v2).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_33_v2), 0))).Int()), _35_v4, !(false), (_40_v8).Cardinality(), _36_v5)).Dtor_cf19()
		} else {
			(globalState).F5 = (_41_v9).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_41_v9).Cardinality()))).Uint32()).(bool)
			var _rhs8 bool = p0
			_ = _rhs8
			var _rhs9 bool = p0
			_ = _rhs9
			var _rhs10 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(939), (p1).Minus(_dafny.IntOfUint32((_46_v12).Cardinality())))
			_ = _rhs10
			var _lhs7 *GlobalState = globalState
			_ = _lhs7
			var _lhs8 *GlobalState = globalState
			_ = _lhs8
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			_lhs7.F1 = _rhs8
			_lhs8.F5 = _rhs9
			_lhs9.F4 = _rhs10
			var _75_v27 _dafny.Array
			_ = _75_v27
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_1
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Map = (func(_76_v9 _dafny.Sequence, _77_p3 _dafny.Int, _78_v2 _dafny.Array, _79_p0 bool) func(_dafny.Int) _dafny.Map {
					return func(_80_i8 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_76_v9).Select((Companion_Default___.SafeIndex(_77_p3, _dafny.IntOfUint32((_76_v9).Cardinality()))).Uint32()).(bool)), (_78_v2).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(204), _dafny.ArrayLen((_78_v2), 0))).Int()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_p0, _dafny.CodePoint('e')))
					}
				})(_41_v9, p3, _33_v2, p0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw5 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw5).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw5).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_75_v27 = _nw5
			var _81_v28 _dafny.Map
			_ = _81_v28
			_81_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.CodePoint('c'))
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_75_v27), 0))
			_ = _index6
			(_75_v27).ArraySet1(_81_v28, (_index6).Int())
			var _82_v29 _dafny.Map
			_ = _82_v29
			_82_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_40_v8).Merge(_40_v8), _81_v28)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_75_v27), 0))
			_ = _index7
			var _rhs11 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.Fm0(_55_i3, (func() _dafny.Int {
				if (_35_v4).Contains(p3) {
					return (_35_v4).Multiplicity(p3)
				}
				return _dafny.IntOfUint32((_41_v9).Cardinality())
			})(), p0, globalState)))
			_ = _rhs11
			var _rhs12 _dafny.Map = (func() _dafny.Map {
				if (_82_v29).Contains(_40_v8) {
					return (_82_v29).Get(_40_v8).(_dafny.Map)
				}
				return (_81_v28).Update(p0, _34_v3)
			})()
			_ = _rhs12
			var _lhs10 *GlobalState = globalState
			_ = _lhs10
			var _lhs11 _dafny.Array = _75_v27
			_ = _lhs11
			var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_75_v27), 0))
			_ = _lhs12
			_lhs10.F4 = _rhs11
			(_lhs11).ArraySet1(_rhs12, (_lhs12).Int())
			_36_v5 = _36_v5
			(globalState).F5 = p0
		}
		(globalState).F5 = p0
		var _83_v30 _dafny.Array
		_ = _83_v30
		var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw6
		_83_v30 = _nw6
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_83_v30), 0))
		_ = _index8
		(_83_v30).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ouixcdf"), (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_83_v30), 0))
		_ = _index9
		var _rhs13 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_46_v12, _46_v12)
		_ = _rhs13
		var _rhs14 bool = ((p1).Times(_55_i3)).Cmp(_55_i3) == 0
		_ = _rhs14
		var _lhs13 _dafny.Array = _83_v30
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(754), _dafny.ArrayLen((_83_v30), 0))
		_ = _lhs14
		var _lhs15 *GlobalState = globalState
		_ = _lhs15
		(_lhs13).ArraySet1(_rhs13, (_lhs14).Int())
		_lhs15.F5 = _rhs14
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _84_v0 _dafny.Int
	_ = _84_v0
	_84_v0 = _dafny.IntOfInt64(624)
	var _85_v1 _dafny.Sequence
	_ = _85_v1
	_85_v1 = _dafny.UnicodeSeqOfUtf8Bytes("vo")
	var _86_v2 _dafny.Array
	_ = _86_v2
	var _nwElement0_1 _dafny.Int = (_dafny.Zero).Minus(_84_v0)
	_ = _nwElement0_1
	var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(12))
	_ = _nw7
	(_nw7).ArraySet1(_nwElement0_1, 0)
	(_nw7).ArraySet1(_84_v0, 1)
	(_nw7).ArraySet1(_84_v0, 2)
	(_nw7).ArraySet1(_84_v0, 3)
	(_nw7).ArraySet1(_84_v0, 4)
	(_nw7).ArraySet1(_84_v0, 5)
	(_nw7).ArraySet1(_84_v0, 6)
	(_nw7).ArraySet1(_84_v0, 7)
	(_nw7).ArraySet1(_84_v0, 8)
	(_nw7).ArraySet1(_84_v0, 9)
	(_nw7).ArraySet1(_84_v0, 10)
	(_nw7).ArraySet1(_dafny.IntOfUint32((_85_v1).Cardinality()), 11)
	_86_v2 = _nw7
	var _87_globalState *GlobalState
	_ = _87_globalState
	var _nw8 *GlobalState = New_GlobalState_()
	_ = _nw8
	_nw8.Ctor__(true, false, _86_v2, _dafny.IntOfInt64(965), _dafny.IntOfInt64(-64), false)
	_87_globalState = _nw8
	(_87_globalState).F4 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_84_v0, _84_v0))
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
	_ = _index10
	(_86_v2).ArraySet1(_84_v0, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
	_ = _index11
	(_86_v2).ArraySet1(_84_v0, (_index11).Int())
	var _hi2 _dafny.Int = _dafny.IntOfUint32((_85_v1).Cardinality())
	_ = _hi2
	for _88_i0 := _dafny.IntOfUint32((_85_v1).Cardinality()); _88_i0.Cmp(_hi2) < 0; _88_i0 = _88_i0.Plus(_dafny.One) {
		var _89_v3 bool
		_ = _89_v3
		_89_v3 = false
		var _90_v4 _dafny.Array
		_ = _90_v4
		var _nwElement0_2 bool = _89_v3
		_ = _nwElement0_2
		var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(17))
		_ = _nw9
		(_nw9).ArraySet1(_nwElement0_2, 0)
		(_nw9).ArraySet1(_89_v3, 1)
		(_nw9).ArraySet1(_89_v3, 2)
		(_nw9).ArraySet1(_89_v3, 3)
		(_nw9).ArraySet1(_89_v3, 4)
		(_nw9).ArraySet1(_89_v3, 5)
		(_nw9).ArraySet1(_89_v3, 6)
		(_nw9).ArraySet1(false, 7)
		(_nw9).ArraySet1(_89_v3, 8)
		(_nw9).ArraySet1(_89_v3, 9)
		(_nw9).ArraySet1(_89_v3, 10)
		(_nw9).ArraySet1(_89_v3, 11)
		(_nw9).ArraySet1(_89_v3, 12)
		(_nw9).ArraySet1(_89_v3, 13)
		(_nw9).ArraySet1(_89_v3, 14)
		(_nw9).ArraySet1(_89_v3, 15)
		(_nw9).ArraySet1(_89_v3, 16)
		_90_v4 = _nw9
		var _91_v5 _dafny.Map
		_ = _91_v5
		_91_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_i0, _90_v4)
		var _92_v6 _dafny.Map
		_ = _92_v6
		_92_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_88_i0, Companion_Default___.Fm0(_88_i0, _84_v0, _89_v3, _87_globalState))
		var _93_v7 _dafny.Sequence
		_ = _93_v7
		_93_v7 = _dafny.SeqOf((_92_v6).Cardinality())
		Companion_Default___.M0(_89_v3, ((_91_v5).Cardinality()).Minus(_84_v0), _93_v7, (func() _dafny.Int {
			if _89_v3 {
				return (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)
			}
			return _88_i0
		})(), _87_globalState)
		var _94_v8 _dafny.Set
		_ = _94_v8
		_94_v8 = _dafny.SetOf((_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)))
		var _95_v9 _dafny.Sequence
		_ = _95_v9
		_95_v9 = _dafny.SeqOf(_89_v3, Companion_Default___.Fm1(_93_v7, _94_v8, _89_v3, _87_globalState), _89_v3, _89_v3)
		var _96_v10 _dafny.Map
		_ = _96_v10
		_96_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_93_v7, _94_v8, _89_v3, _87_globalState), _dafny.Companion_Sequence_.Concatenate(_95_v9, _95_v9))
		_96_v10 = (_96_v10).Update(_89_v3, _95_v9)
		var _97_v11 _dafny.Map
		_ = _97_v11
		_97_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_89_v3, _89_v3)
		var _98_v12 _dafny.MultiSet
		_ = _98_v12
		_98_v12 = _dafny.MultiSetOf(_88_i0, _88_i0, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt(_88_i0, (_97_v11).Cardinality()))
		(_87_globalState).F4 = (func() _dafny.Int {
			if (_98_v12).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_86_v2)).Cardinality())) {
				return (_98_v12).Multiplicity(_dafny.IntOfUint32((_dafny.SeqOf(_86_v2)).Cardinality()))
			}
			return _dafny.IntOfInt64(568)
		})()
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index12
		(_86_v2).ArraySet1(_dafny.IntOfInt64(955), (_index12).Int())
	}
	var _99_v13 *C0
	_ = _99_v13
	var _nw10 *C0 = New_C0_()
	_ = _nw10
	_nw10.Ctor__()
	_99_v13 = _nw10
	var _100_v14 _dafny.Array
	_ = _100_v14
	var _nw11 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
	_ = _nw11
	_100_v14 = _nw11
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))
	_ = _index13
	(_100_v14).ArraySet1(_99_v13, (_index13).Int())
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))
	_ = _index14
	(_100_v14).ArraySet1(_99_v13, (_index14).Int())
	var _101_v15 D0
	_ = _101_v15
	_101_v15 = Companion_D0_.Create_DC1_(_85_v1, !(false), false, _84_v0)
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
	_ = _index15
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
	_ = _index16
	var _rhs15 _dafny.Array = _86_v2
	_ = _rhs15
	var _rhs16 _dafny.Int = (_101_v15).Dtor_cf4()
	_ = _rhs16
	var _rhs17 _dafny.Int = (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)
	_ = _rhs17
	var _rhs18 _dafny.Int = _84_v0
	_ = _rhs18
	var _lhs16 _dafny.Array = _86_v2
	_ = _lhs16
	var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
	_ = _lhs17
	var _lhs18 _dafny.Array = _86_v2
	_ = _lhs18
	var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
	_ = _lhs19
	_86_v2 = _rhs15
	_84_v0 = _rhs16
	(_lhs16).ArraySet1(_rhs17, (_lhs17).Int())
	(_lhs18).ArraySet1(_rhs18, (_lhs19).Int())
	var _102_i1 _dafny.Int
	_ = _102_i1
	_102_i1 = _dafny.Zero
	{
		for ((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).Cmp(((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).Plus(_84_v0)) < 0 {
			{
				if (_102_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_102_i1 = (_102_i1).Plus(_dafny.One)
				var _103_v16 D0
				_ = _103_v16
				_103_v16 = Companion_D0_.Create_DC0_((_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)))
				_103_v16 = _103_v16
				var _104_v17 *C0
				_ = _104_v17
				var _nw12 *C0 = New_C0_()
				_ = _nw12
				_nw12.Ctor__()
				_104_v17 = _nw12
				(_87_globalState).F4 = (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)
				var _105_v18 _dafny.Map
				_ = _105_v18
				_105_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(775), _84_v0)
				var _106_v19 _dafny.Set
				_ = _106_v19
				_106_v19 = _dafny.SetOf((_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)))
				(_87_globalState).F5 = (_106_v19).Contains((func() _dafny.Int {
					if (_105_v18).Contains(_84_v0) {
						return (_105_v18).Get(_84_v0).(_dafny.Int)
					}
					return _84_v0
				})())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _107_v20 _dafny.Map
	_ = _107_v20
	_107_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_108_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})), false)
	var _109_v21 _dafny.MultiSet
	_ = _109_v21
	_109_v21 = _dafny.MultiSetOf(_84_v0)
	var _110_v22 _dafny.Set
	_ = _110_v22
	_110_v22 = _dafny.SetOf(_dafny.IntOfInt64(-744), _dafny.IntOfUint32((_85_v1).Cardinality()), (_109_v21).Cardinality(), (_dafny.Zero).Minus((func() _dafny.Int {
		if (_109_v21).Contains((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)) {
			return (_109_v21).Multiplicity((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
		}
		return (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)
	})()), _84_v0)
	var _111_v23 bool
	_ = _111_v23
	_111_v23 = false
	if Companion_Default___.Fm1(_dafny.SeqOf(_84_v0, _84_v0, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), (_107_v20).Cardinality(), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)), _110_v22, _111_v23, _87_globalState) {
		var _112_v24 _dafny.Array
		_ = _112_v24
		var _nw13 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw13
		_112_v24 = _nw13
		_112_v24 = _112_v24
		var _113_v25 _dafny.Sequence
		_ = _113_v25
		_113_v25 = _dafny.SeqOf(_dafny.IntOfInt64(-239))
		if (_111_v23) || (Companion_Default___.Fm1(_113_v25, _dafny.SetOf(_dafny.IntOfInt64(582)), _111_v23, _87_globalState)) {
			var _114_v26 _dafny.Sequence
			_ = _114_v26
			_114_v26 = _dafny.SeqOf(_111_v23)
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
			_ = _index17
			(_86_v2).ArraySet1(_dafny.IntOfUint32((_114_v26).Cardinality()), (_index17).Int())
			var _115_v27 _dafny.Map
			_ = _115_v27
			_115_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v23, (_110_v22).IsProperSubsetOf(_110_v22))
			_115_v27 = (_115_v27).Update(_111_v23, !(_111_v23))
			var _116_v28 _dafny.Array
			_ = _116_v28
			var _nwElement0_3 _dafny.Sequence = _85_v1
			_ = _nwElement0_3
			var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(3))
			_ = _nw14
			(_nw14).ArraySet1(_nwElement0_3, 0)
			(_nw14).ArraySet1(_85_v1, 1)
			(_nw14).ArraySet1(_85_v1, 2)
			_116_v28 = _nw14
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_116_v28), 0))
			_ = _index18
			(_116_v28).ArraySet1(_85_v1, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_116_v28), 0))
			_ = _index19
			(_116_v28).ArraySet1(_85_v1, (_index19).Int())
			(_87_globalState).F4 = _84_v0
			var _117_v29 _dafny.CodePoint
			_ = _117_v29
			_117_v29 = _dafny.CodePoint('g')
			var _118_v30 _dafny.Map
			_ = _118_v30
			_118_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v1, _117_v29)
			_118_v30 = (_118_v30).Update(_dafny.Companion_Sequence_.Concatenate((_116_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_116_v28), 0))).Int()).(_dafny.Sequence), _85_v1), (func() _dafny.CodePoint {
				if _111_v23 {
					return Companion_Default___.Fm4(!(_111_v23), _84_v0, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _87_globalState)
				}
				return _dafny.CodePoint('m')
			})())
		} else {
			var _119_v31 _dafny.Sequence
			_ = _119_v31
			_119_v31 = _dafny.SeqOf((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0), _99_v13)
			(_87_globalState).F4 = (_dafny.Zero).Minus((_113_v25).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_119_v31).Cardinality()), _dafny.IntOfUint32((_113_v25).Cardinality()))).Uint32()).(_dafny.Int))
			_111_v23 = (func() bool {
				if _111_v23 {
					return (_84_v0).Cmp((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)) > 0
				}
				return Companion_Default___.Fm1(_113_v25, _dafny.SetOf(_84_v0), false, _87_globalState)
			})()
			var _120_v32 _dafny.Map
			_ = _120_v32
			_120_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_110_v22).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v15, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))).Cardinality()), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
			_120_v32 = _120_v32
			var _121_v33 _dafny.Sequence
			_ = _121_v33
			_121_v33 = _dafny.SeqOf(_111_v23)
			var _122_v34 _dafny.CodePoint
			_ = _122_v34
			_122_v34 = _dafny.CodePoint('b')
			var _123_v35 D0
			_ = _123_v35
			_123_v35 = Companion_D0_.Create_DC2_(_dafny.UnicodeSeqOfUtf8Bytes("cdwg"), _111_v23, _121_v33, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _122_v34)
			(_87_globalState).F4 = (_123_v35).Dtor_cf8()
			(_87_globalState).F1 = (_111_v23) == (_111_v23)
		}
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_112_v24), 0))
		_ = _index20
		(_112_v24).ArraySet1(true, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_112_v24), 0))
		_ = _index21
		(_112_v24).ArraySet1(!(_111_v23), (_index21).Int())
		var _124_v36 D0
		_ = _124_v36
		_124_v36 = Companion_D0_.Create_DC0_(_84_v0)
		var _125_v37 _dafny.Map
		_ = _125_v37
		_125_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v23, _111_v23)
		var _126_v38 _dafny.Map
		_ = _126_v38
		_126_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_113_v25).Cardinality())).Plus(_84_v0), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_112_v24), 0))
		_ = _index22
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_112_v24), 0))
		_ = _index23
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index24
		var _rhs19 bool = _dafny.Companion_Sequence_.IsPrefixOf(_113_v25, Companion_Default___.Fm5((_124_v36).Dtor_cf0(), _111_v23, (func() bool {
			if (_125_v37).Contains(_111_v23) {
				return (_125_v37).Get(_111_v23).(bool)
			}
			return _111_v23
		})(), _87_globalState))
		_ = _rhs19
		var _rhs20 bool = _111_v23
		_ = _rhs20
		var _rhs21 _dafny.Int = (_84_v0).Plus((_84_v0).Plus(_dafny.IntOfInt64(634)))
		_ = _rhs21
		var _rhs22 _dafny.Int = (func() _dafny.Int {
			if (_126_v38).Contains((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)) {
				return (_126_v38).Get((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).(_dafny.Int)
			}
			return _84_v0
		})()
		_ = _rhs22
		var _rhs23 bool = true
		_ = _rhs23
		var _lhs20 _dafny.Array = _112_v24
		_ = _lhs20
		var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_112_v24), 0))
		_ = _lhs21
		var _lhs22 _dafny.Array = _112_v24
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_112_v24), 0))
		_ = _lhs23
		var _lhs24 _dafny.Array = _86_v2
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _lhs25
		(_lhs20).ArraySet1(_rhs19, (_lhs21).Int())
		(_lhs22).ArraySet1(_rhs20, (_lhs23).Int())
		_84_v0 = _rhs21
		(_lhs24).ArraySet1(_rhs22, (_lhs25).Int())
		_111_v23 = _rhs23
		var _127_v39 _dafny.MultiSet
		_ = _127_v39
		_127_v39 = _dafny.MultiSetOf((_112_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_112_v24), 0))).Int()).(bool), _111_v23, _111_v23)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index25
		(_86_v2).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(111), (func() _dafny.Int {
			if (_127_v39).Contains(_111_v23) {
				return (_127_v39).Multiplicity(_111_v23)
			}
			return (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)
		})()), (_index25).Int())
		var _128_v40 _dafny.MultiSet
		_ = _128_v40
		_128_v40 = _dafny.MultiSetOf(_112_v24, _112_v24)
		var _129_v41 _dafny.Map
		_ = _129_v41
		_129_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), false)
		var _130_v42 _dafny.Map
		_ = _130_v42
		_130_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0), _111_v23)
		var _131_v43 D0
		_ = _131_v43
		_131_v43 = Companion_D0_.Create_DC4_(Companion_Default___.Fm6((_128_v40).Cardinality(), _129_v41, (func() bool {
			if (_130_v42).Contains((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0)) {
				return (_130_v42).Get((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0)).(bool)
			}
			return (_112_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_112_v24), 0))).Int()).(bool)
		})(), false, _87_globalState))
		var _132_v44 _dafny.Sequence
		_ = _132_v44
		_132_v44 = _dafny.SeqOf(!(true), (_112_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_112_v24), 0))).Int()).(bool))
		var _133_v45 _dafny.CodePoint
		_ = _133_v45
		_133_v45 = _dafny.CodePoint('a')
		var _134_v46 D0
		_ = _134_v46
		_134_v46 = Companion_D0_.Create_DC2_(_85_v1, _111_v23, _132_v44, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qra")).Cardinality()), _133_v45)
		var _pat_let_tv1 = _134_v46
		_ = _pat_let_tv1
		_131_v43 = func(_pat_let0_0 D0) D0 {
			return func(_135_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 D0) D0 {
					return func(_136_dt__update_hcf10_h0 D0) D0 {
						return Companion_D0_.Create_DC4_(_136_dt__update_hcf10_h0)
					}(_pat_let1_0)
				}(_pat_let_tv1)
			}(_pat_let0_0)
		}(_131_v43)
	} else {
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))
		_ = _index26
		(_100_v14).ArraySet1(_99_v13, (_index26).Int())
		var _137_v47 *C0
		_ = _137_v47
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__()
		_137_v47 = _nw15
		var _138_v48 _dafny.Sequence
		_ = _138_v48
		_138_v48 = _dafny.SeqOf(_84_v0, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_85_v1).Cardinality()))
		var _139_v50 _dafny.Sequence
		_ = _139_v50
		_139_v50 = _dafny.SeqOf(_111_v23)
		(_87_globalState).F1 = Companion_Default___.Fm1(_138_v48, (_110_v22).Difference(func() _dafny.Set {
			var _coll2 = _dafny.NewBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate((_dafny.MultiSetFromSeq(_138_v48)).Elements()); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _140_v49 _dafny.Int
				_140_v49 = interface{}(_compr_2).(_dafny.Int)
				if (_dafny.MultiSetFromSeq(_138_v48)).Contains(_140_v49) {
					_coll2.Add((_140_v49).Plus(_dafny.IntOfInt64(426)))
				}
			}
			return _coll2.ToSet()
		}()), (_111_v23) || ((_139_v50).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_139_v50).Cardinality()))).Uint32()).(bool)), _87_globalState)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index27
		(_86_v2).ArraySet1(_84_v0, (_index27).Int())
		var _141_v51 _dafny.Map
		_ = _141_v51
		_141_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, _111_v23)
		var _142_v52 _dafny.CodePoint
		_ = _142_v52
		_142_v52 = _dafny.CodePoint('s')
		var _143_v53 _dafny.Map
		_ = _143_v53
		_143_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(281), _142_v52)
		var _144_v54 _dafny.Array
		_ = _144_v54
		var _nwElement0_4 bool = !(false)
		_ = _nwElement0_4
		var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
		_ = _nw16
		(_nw16).ArraySet1(_nwElement0_4, 0)
		(_nw16).ArraySet1(!(Companion_Default___.Fm1(_dafny.SeqOf(_84_v0), _110_v22, _111_v23, _87_globalState)) || (_111_v23), 1)
		(_nw16).ArraySet1(!((_139_v50).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_139_v50).Cardinality()))).Uint32()).(bool)) || (!(_111_v23)), 2)
		(_nw16).ArraySet1(_111_v23, 3)
		(_nw16).ArraySet1((func() bool {
			if (func() bool {
				if (_141_v51).Contains((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)) {
					return (_141_v51).Get((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).(bool)
				}
				return _111_v23
			})() {
				return _111_v23
			}
			return !(_111_v23)
		})(), 4)
		(_nw16).ArraySet1(true, 5)
		(_nw16).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_138_v48, _138_v48), 6)
		(_nw16).ArraySet1(!(_143_v53).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, _142_v52)), 7)
		(_nw16).ArraySet1(_111_v23, 8)
		_144_v54 = _nw16
		_144_v54 = _144_v54
	}
	var _145_v55 _dafny.MultiSet
	_ = _145_v55
	_145_v55 = _dafny.MultiSetOf(_111_v23, _111_v23)
	var _146_v56 _dafny.Sequence
	_ = _146_v56
	_146_v56 = _dafny.SeqOf(_145_v55)
	(_87_globalState).F4 = (_dafny.Zero).Minus(((_146_v56).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_146_v56).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality())
	var _147_v58 _dafny.CodePoint
	_ = _147_v58
	_147_v58 = _dafny.CodePoint('i')
	var _148_v59 _dafny.Map
	_ = _148_v59
	_148_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SetOf(_147_v58)).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _149_v57 _dafny.CodePoint
			_149_v57 = interface{}(_compr_3).(_dafny.CodePoint)
			if (_dafny.SetOf(_147_v58)).Contains(_149_v57) {
				_coll3.Add(_149_v57, _111_v23)
			}
		}
		return _coll3.ToMap()
	}())
	var _hi3 _dafny.Int = _84_v0
	_ = _hi3
	for _150_i3 := (_148_v59).Cardinality(); _150_i3.Cmp(_hi3) < 0; _150_i3 = _150_i3.Plus(_dafny.One) {
		var _151_v60 *C0
		_ = _151_v60
		var _nw17 *C0 = New_C0_()
		_ = _nw17
		_nw17.Ctor__()
		_151_v60 = _nw17
		var _152_v61 _dafny.MultiSet
		_ = _152_v61
		_152_v61 = _dafny.MultiSetOf(_100_v14)
		_152_v61 = (_152_v61).Union((_152_v61).Update(_100_v14, Companion_Default___.Abs((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))))
		var _153_v62 D0
		_ = _153_v62
		_153_v62 = Companion_D0_.Create_DC4_(Companion_D0_.Create_DC1_(_85_v1, _111_v23, _111_v23, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(437))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_154_v0 _dafny.Int, _155_v23 bool) func(_dafny.Int) _dafny.Map {
			return func(_156_i4 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v0, _155_v23)
			}
		})(_84_v0, _111_v23)))).Cardinality())))
		var _157_v63 D0
		_ = _157_v63
		_157_v63 = Companion_D0_.Create_DC4_(_153_v62)
		var _158_v64 _dafny.MultiSet
		_ = _158_v64
		_158_v64 = _dafny.MultiSetOf(Companion_D0_.Create_DC4_(_153_v62))
		_84_v0 = (_dafny.Zero).Minus((((_158_v64).Difference(_158_v64)).Cardinality()).Times((_145_v55).Cardinality()))
		(_87_globalState).F1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_85_v1, _85_v1)
	}
	var _159_i5 _dafny.Int
	_ = _159_i5
	_159_i5 = _dafny.Zero
	{
		for _111_v23 {
			{
				if (_159_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_159_i5 = (_159_i5).Plus(_dafny.One)
				(_87_globalState).F4 = _84_v0
				(_87_globalState).F5 = !(Companion_Default___.Fm7(_147_v58, _87_globalState)).Contains(_111_v23)
				(_87_globalState).F4 = ((_dafny.Zero).Minus(_84_v0)).Minus(_dafny.IntOfInt64(215))
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
				_ = _index28
				(_86_v2).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(698))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_160_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_161_i6 _dafny.Int) _dafny.CodePoint {
						return _160_v58
					}
				})(_147_v58))), _dafny.UnicodeSeqOfUtf8Bytes("ydsnaxu"))).Cardinality()), (_index28).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	if _111_v23 {
		var _162_v65 _dafny.Sequence
		_ = _162_v65
		_162_v65 = _dafny.SeqOf((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0))
		var _163_v66 D0
		_ = _163_v66
		_163_v66 = Companion_D0_.Create_DC0_((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))
		_ = _index29
		(_100_v14).ArraySet1((_162_v65).Select((Companion_Default___.SafeIndex((_163_v66).Dtor_cf0(), _dafny.IntOfUint32((_162_v65).Cardinality()))).Uint32()).(*C0), (_index29).Int())
		(_87_globalState).F4 = (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)
		var _164_v67 _dafny.Sequence
		_ = _164_v67
		_164_v67 = _dafny.SeqOf(_84_v0)
		var _165_v68 _dafny.Array
		_ = _165_v68
		var _nwElement0_5 bool = false
		_ = _nwElement0_5
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(26))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_5, 0)
		(_nw18).ArraySet1(_dafny.Companion_Sequence_.Contains(_85_v1, _147_v58), 1)
		(_nw18).ArraySet1(_111_v23, 2)
		(_nw18).ArraySet1(_111_v23, 3)
		(_nw18).ArraySet1(_111_v23, 4)
		(_nw18).ArraySet1(_111_v23, 5)
		(_nw18).ArraySet1(_111_v23, 6)
		(_nw18).ArraySet1((_111_v23) || (_111_v23), 7)
		(_nw18).ArraySet1(_111_v23, 8)
		(_nw18).ArraySet1(!(_111_v23), 9)
		(_nw18).ArraySet1(!(_111_v23), 10)
		(_nw18).ArraySet1(((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).Cmp(_84_v0) <= 0, 11)
		(_nw18).ArraySet1(_111_v23, 12)
		(_nw18).ArraySet1(false, 13)
		(_nw18).ArraySet1(_111_v23, 14)
		(_nw18).ArraySet1(_111_v23, 15)
		(_nw18).ArraySet1(Companion_Default___.Fm1(_164_v67, _110_v22, false, _87_globalState), 16)
		(_nw18).ArraySet1(_111_v23, 17)
		(_nw18).ArraySet1(_111_v23, 18)
		(_nw18).ArraySet1(!(_111_v23) || (_111_v23), 19)
		(_nw18).ArraySet1(_111_v23, 20)
		(_nw18).ArraySet1(_111_v23, 21)
		(_nw18).ArraySet1(_111_v23, 22)
		(_nw18).ArraySet1(true, 23)
		(_nw18).ArraySet1(_111_v23, 24)
		(_nw18).ArraySet1(false, 25)
		_165_v68 = _nw18
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_165_v68), 0))
		_ = _index30
		(_165_v68).ArraySet1((_111_v23) && (_111_v23), (_index30).Int())
		var _166_v69 _dafny.Set
		_ = _166_v69
		_166_v69 = _dafny.SetOf(_111_v23)
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_165_v68), 0))
		_ = _index31
		(_165_v68).ArraySet1(!((_dafny.SetOf(_111_v23, _111_v23, true)).IsSubsetOf(_166_v69)), (_index31).Int())
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index32
		(_86_v2).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_85_v1, _85_v1), _dafny.UnicodeSeqOfUtf8Bytes("xvxf")), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-810), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_85_v1, _85_v1), _dafny.UnicodeSeqOfUtf8Bytes("xvxf"))).Cardinality()))).Uint32(), _147_v58)).Cardinality()), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))
		_ = _index33
		(_100_v14).ArraySet1((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0), (_index33).Int())
	} else {
		(_87_globalState).F1 = _111_v23
		var _167_v70 _dafny.Sequence
		_ = _167_v70
		_167_v70 = _dafny.SeqOf((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index34
		var _rhs24 bool = (func() bool {
			if _111_v23 {
				return _111_v23
			}
			return (_145_v55).IsSubsetOf((_145_v55).Update(_111_v23, Companion_Default___.Abs((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v2, _84_v0)).Cardinality())))
		})()
		_ = _rhs24
		var _rhs25 bool = (_111_v23) && (Companion_Default___.Fm1(_167_v70, _dafny.SetOf((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)), _111_v23, _87_globalState))
		_ = _rhs25
		var _rhs26 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(70), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
		_ = _rhs26
		var _lhs26 *GlobalState = _87_globalState
		_ = _lhs26
		var _lhs27 _dafny.Array = _86_v2
		_ = _lhs27
		var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _lhs28
		_lhs26.F1 = _rhs24
		_111_v23 = _rhs25
		(_lhs27).ArraySet1(_rhs26, (_lhs28).Int())
		Companion_Default___.M0((!(_111_v23)) == (Companion_Default___.Fm1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(627))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}((func(_168_v2 _dafny.Array) func(_dafny.Int) _dafny.Int {
			return func(_169_i7 _dafny.Int) _dafny.Int {
				return (_168_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_168_v2), 0))).Int()).(_dafny.Int)
			}
		})(_86_v2))), _110_v22, Companion_Default___.Fm1(_dafny.SeqOf(_dafny.IntOfInt64(758)), _110_v22, _111_v23, _87_globalState), _87_globalState)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qjjinm")).Cardinality()), _167_v70, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _87_globalState)
		(_87_globalState).F4 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_84_v0, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), false, _87_globalState), _84_v0)
		var _170_v71 _dafny.Array
		_ = _170_v71
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
		_ = _nw19
		_170_v71 = _nw19
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_170_v71), 0))
		_ = _index35
		(_170_v71).ArraySet1((func() _dafny.Sequence {
			if Companion_Default___.Fm1(_167_v70, _110_v22, _111_v23, _87_globalState) {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-850))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}((func(_171_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_172_i8 _dafny.Int) _dafny.CodePoint {
						return _171_v58
					}
				})(_147_v58)))
			}
			return (_99_v13).Fm2(!(_111_v23), _84_v0, _87_globalState)
		})(), (_index35).Int())
		var _173_v72 _dafny.Array
		_ = _173_v72
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw20
		_173_v72 = _nw20
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_173_v72), 0))
		_ = _index36
		(_173_v72).ArraySet1(false, (_index36).Int())
		var _174_v73 _dafny.Map
		_ = _174_v73
		_174_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _85_v1)
		var _175_v74 _dafny.Sequence
		_ = _175_v74
		_175_v74 = _dafny.SeqOf(_111_v23, _111_v23, !(!((func() bool {
			if _111_v23 {
				return _111_v23
			}
			return _111_v23
		})())), Companion_Default___.Fm1(_167_v70, _110_v22, _111_v23, _87_globalState))
		var _176_v75 _dafny.MultiSet
		_ = _176_v75
		_176_v75 = _dafny.MultiSetOf(_86_v2)
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_170_v71), 0))
		_ = _index37
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _index38
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_173_v72), 0))
		_ = _index39
		var _rhs27 _dafny.Sequence = _85_v1
		_ = _rhs27
		var _rhs28 _dafny.Int = ((Companion_D1_.Create_DC5_(_174_v73)).Dtor_cf11()).Cardinality()
		_ = _rhs28
		var _rhs29 _dafny.Int = (_84_v0).Plus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
		_ = _rhs29
		var _rhs30 bool = !((_175_v74).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_175_v74).Cardinality()))).Uint32()).(bool))
		_ = _rhs30
		var _rhs31 bool = (_176_v75).IsProperSubsetOf(_dafny.MultiSetOf(_86_v2))
		_ = _rhs31
		var _lhs29 _dafny.Array = _170_v71
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_170_v71), 0))
		_ = _lhs30
		var _lhs31 _dafny.Array = _86_v2
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))
		_ = _lhs32
		var _lhs33 *GlobalState = _87_globalState
		_ = _lhs33
		var _lhs34 _dafny.Array = _173_v72
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_173_v72), 0))
		_ = _lhs35
		(_lhs29).ArraySet1(_rhs27, (_lhs30).Int())
		(_lhs31).ArraySet1(_rhs28, (_lhs32).Int())
		_84_v0 = _rhs29
		_lhs33.F1 = _rhs30
		(_lhs34).ArraySet1(_rhs31, (_lhs35).Int())
	}
	var _177_v76 _dafny.Sequence
	_ = _177_v76
	_177_v76 = _dafny.SeqOf(_111_v23, _111_v23, false)
	var _178_v77 _dafny.Map
	_ = _178_v77
	_178_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v23, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_111_v23), _177_v76))
	var _179_v78 D0
	_ = _179_v78
	_179_v78 = Companion_D0_.Create_DC3_()
	var _180_v79 D0
	_ = _180_v79
	_180_v79 = Companion_D0_.Create_DC4_(_179_v78)
	var _181_v80 _dafny.Sequence
	_ = _181_v80
	_181_v80 = _dafny.SeqOf(_84_v0, _dafny.IntOfUint32((_177_v76).Cardinality()))
	var _pat_let_tv2 = _178_v77
	_ = _pat_let_tv2
	var _pat_let_tv3 = _178_v77
	_ = _pat_let_tv3
	var _pat_let_tv4 = _177_v76
	_ = _pat_let_tv4
	var _pat_let_tv5 = _178_v77
	_ = _pat_let_tv5
	var _pat_let_tv6 = _178_v77
	_ = _pat_let_tv6
	var _pat_let_tv7 = _178_v77
	_ = _pat_let_tv7
	var _rhs32 _dafny.Map = func(_source2 D0) _dafny.Map {
		if _source2.Is_DC1() {
			var _182___mcc_h0 _dafny.Sequence = _source2.Get_().(D0_DC1).Cf1
			_ = _182___mcc_h0
			var _183___mcc_h1 bool = _source2.Get_().(D0_DC1).Cf2
			_ = _183___mcc_h1
			var _184___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf3
			_ = _184___mcc_h2
			var _185___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC1).Cf4
			_ = _185___mcc_h3
			var _186_cf4 _dafny.Int = _185___mcc_h3
			_ = _186_cf4
			var _187_cf3 bool = _184___mcc_h2
			_ = _187_cf3
			var _188_cf2 bool = _183___mcc_h1
			_ = _188_cf2
			var _189_cf1 _dafny.Sequence = _182___mcc_h0
			_ = _189_cf1
			return (_pat_let_tv2).Merge(_pat_let_tv3)
		} else if _source2.Is_DC2() {
			var _190___mcc_h4 _dafny.Sequence = _source2.Get_().(D0_DC2).Cf5
			_ = _190___mcc_h4
			var _191___mcc_h5 bool = _source2.Get_().(D0_DC2).Cf6
			_ = _191___mcc_h5
			var _192___mcc_h6 _dafny.Sequence = _source2.Get_().(D0_DC2).Cf7
			_ = _192___mcc_h6
			var _193___mcc_h7 _dafny.Int = _source2.Get_().(D0_DC2).Cf8
			_ = _193___mcc_h7
			var _194___mcc_h8 _dafny.CodePoint = _source2.Get_().(D0_DC2).Cf9
			_ = _194___mcc_h8
			var _195_cf9 _dafny.CodePoint = _194___mcc_h8
			_ = _195_cf9
			var _196_cf8 _dafny.Int = _193___mcc_h7
			_ = _196_cf8
			var _197_cf7 _dafny.Sequence = _192___mcc_h6
			_ = _197_cf7
			var _198_cf6 bool = _191___mcc_h5
			_ = _198_cf6
			var _199_cf5 _dafny.Sequence = _190___mcc_h4
			_ = _199_cf5
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _pat_let_tv4)
		} else if _source2.Is_DC3() {
			return _pat_let_tv5
		} else if _source2.Is_DC0() {
			var _200___mcc_h9 _dafny.Int = _source2.Get_().(D0_DC0).Cf0
			_ = _200___mcc_h9
			var _201_cf0 _dafny.Int = _200___mcc_h9
			_ = _201_cf0
			return _pat_let_tv6
		} else {
			var _202___mcc_h10 D0 = _source2.Get_().(D0_DC4).Cf10
			_ = _202___mcc_h10
			var _203_cf10 D0 = _202___mcc_h10
			_ = _203_cf10
			return (_pat_let_tv7)
		}
	}(_180_v79)
	_ = _rhs32
	var _rhs33 D0 = _101_v15
	_ = _rhs33
	var _rhs34 _dafny.Int = _84_v0
	_ = _rhs34
	var _rhs35 _dafny.Sequence = (_99_v13).Fm2(!(Companion_Default___.Fm1(_181_v80, Companion_Default___.Fm8(_111_v23, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(652), _87_globalState), false, _87_globalState)), _84_v0, _87_globalState)
	_ = _rhs35
	var _rhs36 _dafny.Int = ((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).Times(_84_v0)
	_ = _rhs36
	var _lhs36 *GlobalState = _87_globalState
	_ = _lhs36
	_178_v77 = _rhs32
	_101_v15 = _rhs33
	_lhs36.F4 = _rhs34
	_85_v1 = _rhs35
	_84_v0 = _rhs36
	var _204_i9 _dafny.Int
	_ = _204_i9
	_204_i9 = _dafny.Zero
	{
		for !(_111_v23) {
			{
				if (_204_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_204_i9 = (_204_i9).Plus(_dafny.One)
				var _205_v81 _dafny.Map
				_ = _205_v81
				_205_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_111_v23), _84_v0)
				Companion_Default___.M0(((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
					if (_205_v81).Contains(_111_v23) {
						return (_205_v81).Get(_111_v23).(_dafny.Int)
					}
					return _dafny.IntOfInt64(637)
				})()) != 0, ((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)).Minus(_dafny.IntOfInt64(820)), _181_v80, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(525))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_206_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_207_i10 _dafny.Int) _dafny.Int {
						return _206_v0
					}
				})(_84_v0)))).Cardinality()), _87_globalState)
				if _111_v23 {
					Companion_Default___.M0((_84_v0).Cmp((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)) < 0, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), (func() _dafny.Sequence {
						if _111_v23 {
							return _181_v80
						}
						return _181_v80
					})(), _84_v0, _87_globalState)
					(_87_globalState).F1 = _111_v23
					(_87_globalState).F5 = (_111_v23) == ((_dafny.SetOf(_84_v0)).IsDisjointFrom(_110_v22))
					var _208_v82 _dafny.Set
					_ = _208_v82
					_208_v82 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("oiyf"))
					(_87_globalState).F5 = ((_dafny.SetOf(_85_v1)).Union(_208_v82)).IsSubsetOf(Companion_Default___.Fm9(_111_v23, false, _84_v0, _87_globalState))
					var _209_v83 _dafny.Sequence
					_ = _209_v83
					_209_v83 = _dafny.SeqOf(_85_v1)
					var _210_v84 _dafny.Sequence
					_ = _210_v84
					_210_v84 = _dafny.SeqOf((_209_v83).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_209_v83).Cardinality()))).Uint32()).(_dafny.Sequence))
					var _rhs37 _dafny.MultiSet = (func() _dafny.MultiSet {
						if (_208_v82).IsDisjointFrom(func() _dafny.Set {
							var _coll4 = _dafny.NewBuilder()
							_ = _coll4
							for _iter4 := _dafny.Iterate((func() _dafny.Set {
								var _coll5 = _dafny.NewBuilder()
								_ = _coll5
								for _iter5 := _dafny.Iterate((_210_v84).Elements()); ; {
									_compr_5, _ok5 := _iter5()
									if !_ok5 {
										break
									}
									var _211_v85 _dafny.Sequence
									_211_v85 = interface{}(_compr_5).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_210_v84, _211_v85) {
										_coll5.Add(_211_v85)
									}
								}
								return _coll5.ToSet()
							}()).Elements()); ; {
								_compr_4, _ok4 := _iter4()
								if !_ok4 {
									break
								}
								var _212_v86 _dafny.Sequence
								_212_v86 = interface{}(_compr_4).(_dafny.Sequence)
								if (func() _dafny.Set {
									var _coll6 = _dafny.NewBuilder()
									_ = _coll6
									for _iter6 := _dafny.Iterate((_210_v84).Elements()); ; {
										_compr_6, _ok6 := _iter6()
										if !_ok6 {
											break
										}
										var _213_v85 _dafny.Sequence
										_213_v85 = interface{}(_compr_6).(_dafny.Sequence)
										if _dafny.Companion_Sequence_.Contains(_210_v84, _213_v85) {
											_coll6.Add(_213_v85)
										}
									}
									return _coll6.ToSet()
								}()).Contains(_212_v86) {
									_coll4.Add(_212_v86)
								}
							}
							return _coll4.ToSet()
						}()) {
							return _145_v55
						}
						return _145_v55
					})()
					_ = _rhs37
					var _rhs38 _dafny.Map = (_205_v81).Merge((_205_v81).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v23, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))))
					_ = _rhs38
					_145_v55 = _rhs37
					_205_v81 = _rhs38
				} else {
					Companion_Default___.M0(!(_111_v23), (func() _dafny.Int {
						if _111_v23 {
							return _84_v0
						}
						return (_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
					})(), _181_v80, _dafny.IntOfInt64(391), _87_globalState)
					var _214_v87 *C0
					_ = _214_v87
					var _nw21 *C0 = New_C0_()
					_ = _nw21
					_nw21.Ctor__()
					_214_v87 = _nw21
					var _215_v88 _dafny.Sequence
					_ = _215_v88
					_215_v88 = _dafny.SeqOf((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0), (_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0), _214_v87)
					Companion_Default___.M0(true, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_215_v88).Cardinality()), _84_v0), (_181_v80), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _87_globalState)
					var _216_v90 _dafny.Map
					_ = _216_v90
					_216_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_85_v1).Cardinality()), _181_v80)
					var _217_v91 _dafny.Map
					_ = _217_v91
					_217_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0), (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))
					var _218_v92 _dafny.Sequence
					_ = _218_v92
					_218_v92 = _dafny.SeqOf(_217_v91, _217_v91)
					var _219_v93 _dafny.Sequence
					_ = _219_v93
					_219_v93 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_181_v80, (Companion_Default___.SafeIndex((Companion_Default___.Fm7(_147_v58, _87_globalState)).Cardinality(), _dafny.IntOfUint32((_181_v80).Cardinality()))).Uint32(), (func() _dafny.Map {
						var _coll7 = _dafny.NewMapBuilder()
						_ = _coll7
						for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(958), _dafny.IntOfInt64(-17))); ; {
							_compr_7, _ok7 := _iter7()
							if !_ok7 {
								break
							}
							var _220_v89 _dafny.Int
							_220_v89 = interface{}(_compr_7).(_dafny.Int)
							if ((_dafny.IntOfInt64(958)).Cmp(_220_v89) <= 0) && ((_220_v89).Cmp(_dafny.IntOfInt64(-17)) < 0) {
								_coll7.Add((_220_v89).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)), _111_v23)
							}
						}
						return _coll7.ToMap()
					}()).Cardinality()), _181_v80, (func() _dafny.Sequence {
						if (_216_v90).Contains(_dafny.IntOfUint32((_218_v92).Cardinality())) {
							return (_216_v90).Get(_dafny.IntOfUint32((_218_v92).Cardinality())).(_dafny.Sequence)
						}
						return _181_v80
					})(), _181_v80)
					(_87_globalState).F1 = ((_dafny.Zero).Minus((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfUint32(((_219_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.IntOfUint32((_219_v93).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) <= 0
					Companion_Default___.M0(_111_v23, (_84_v0).Times(_84_v0), _181_v80, (_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _87_globalState)
				}
				var _221_v94 _dafny.Array
				_ = _221_v94
				var _nw22 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw22
				_221_v94 = _nw22
				var _222_v95 _dafny.Map
				_ = _222_v95
				_222_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int), _111_v23)
				var _223_v96 _dafny.Map
				_ = _223_v96
				_223_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(544), _222_v95)
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_221_v94), 0))
				_ = _index40
				(_221_v94).ArraySet1(_223_v96, (_index40).Int())
				var _224_v98 D4
				_ = _224_v98
				_224_v98 = Companion_D4_.Create_DC10_(func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-558), _dafny.IntOfInt64(248))); ; {
						_compr_8, _ok8 := _iter8()
						if !_ok8 {
							break
						}
						var _225_v97 _dafny.Int
						_225_v97 = interface{}(_compr_8).(_dafny.Int)
						if ((_dafny.IntOfInt64(-558)).Cmp(_225_v97) <= 0) && ((_225_v97).Cmp(_dafny.IntOfInt64(248)) < 0) {
							_coll8.Add((_225_v97).Times(_84_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v0, _111_v23))
						}
					}
					return _coll8.ToMap()
				}())
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_221_v94), 0))
				_ = _index41
				(_221_v94).ArraySet1(((_224_v98).Dtor_cf15()).Merge(_223_v96), (_index41).Int())
				var _226_v99 _dafny.Set
				_ = _226_v99
				_226_v99 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v23, (_100_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_100_v14), 0))).Int()).(*C0)))
				var _227_v100 _dafny.Map
				_ = _227_v100
				_227_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_85_v1).Cardinality()), _226_v99)
				var _228_v101 _dafny.Sequence
				_ = _228_v101
				_228_v101 = _dafny.SeqOf(_226_v99, _226_v99, (func() _dafny.Set {
					if (_227_v100).Contains(_dafny.IntOfInt64(921)) {
						return (_227_v100).Get(_dafny.IntOfInt64(921)).(_dafny.Set)
					}
					return _226_v99
				})())
				var _229_v102 _dafny.Array
				_ = _229_v102
				var _nwElement0_6 bool = (_145_v55).Contains(_111_v23)
				_ = _nwElement0_6
				var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(24))
				_ = _nw23
				(_nw23).ArraySet1(_nwElement0_6, 0)
				(_nw23).ArraySet1(true, 1)
				(_nw23).ArraySet1(_111_v23, 2)
				(_nw23).ArraySet1(_111_v23, 3)
				(_nw23).ArraySet1(_111_v23, 4)
				(_nw23).ArraySet1(_111_v23, 5)
				(_nw23).ArraySet1(Companion_Default___.Fm1(_181_v80, _110_v22, _111_v23, _87_globalState), 6)
				(_nw23).ArraySet1((_dafny.IntOfInt64(407)).Cmp((_86_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_86_v2), 0))).Int()).(_dafny.Int)) < 0, 7)
				(_nw23).ArraySet1(_111_v23, 8)
				(_nw23).ArraySet1(_111_v23, 9)
				(_nw23).ArraySet1(_111_v23, 10)
				(_nw23).ArraySet1(_111_v23, 11)
				(_nw23).ArraySet1(_111_v23, 12)
				(_nw23).ArraySet1(!_dafny.Companion_Sequence_.Contains(_181_v80, _84_v0), 13)
				(_nw23).ArraySet1(!(_111_v23) || (_111_v23), 14)
				(_nw23).ArraySet1((_226_v99).IsProperSubsetOf((_228_v101).Select((Companion_Default___.SafeIndex(_84_v0, _dafny.IntOfUint32((_228_v101).Cardinality()))).Uint32()).(_dafny.Set)), 15)
				(_nw23).ArraySet1(_111_v23, 16)
				(_nw23).ArraySet1(_111_v23, 17)
				(_nw23).ArraySet1(_111_v23, 18)
				(_nw23).ArraySet1(_111_v23, 19)
				(_nw23).ArraySet1(_111_v23, 20)
				(_nw23).ArraySet1(!(_111_v23), 21)
				(_nw23).ArraySet1(_111_v23, 22)
				(_nw23).ArraySet1(_111_v23, 23)
				_229_v102 = _nw23
				_229_v102 = _229_v102
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _230_v103 _dafny.Array
	_ = _230_v103
	var _nw24 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
	_ = _nw24
	_230_v103 = _nw24
	for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_230_v103), 0))); ; {
		_guard_loop_0, _ok9 := _iter9()
		if !_ok9 {
			break
		}
		var _231_i11 _dafny.Int
		_231_i11 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_231_i11).Sign() != -1) && ((_231_i11).Cmp(_dafny.ArrayLen((_230_v103), 0)) < 0)) {
			(_230_v103).ArraySet1(_111_v23, (_231_i11).Int())
		}
	}
	var _232_v104 D4
	_ = _232_v104
	_232_v104 = Companion_D4_.Create_DC12_(_147_v58, _109_v21, _111_v23, _dafny.IntOfUint32((_85_v1).Cardinality()), _99_v13)
	var _233_v105 _dafny.MultiSet
	_ = _233_v105
	_233_v105 = _dafny.MultiSetOf((_232_v104).Dtor_cf20())
	_233_v105 = _233_v105
	_dafny.Print(_84_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v2).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_87_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_87_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_87_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_87_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v15).Dtor_cf1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v15).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v15).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v15).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_102_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_107_v20).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t')), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v21).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(624))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_110_v22).Equals(_dafny.SetOf(_dafny.IntOfInt64(-744), _dafny.IntOfInt64(2), _dafny.One, _dafny.IntOfInt64(-1), _dafny.IntOfInt64(624))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_v23)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_v55).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_146_v56, _dafny.SeqOf(_dafny.MultiSetOf(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_147_v58)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_v59).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(624), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_159_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_177_v76, _dafny.SeqOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v77).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false, false, false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_181_v80, _dafny.SeqOf(_dafny.IntOfInt64(70), _dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_204_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v103).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v104).Dtor_cf16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_232_v104).Dtor_cf17()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(624))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v104).Dtor_cf18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v104).Dtor_cf19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v105).Cardinality())
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
	Cf2 bool
	Cf3 bool
	Cf4 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Sequence, Cf2 bool, Cf3 bool, Cf4 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 _dafny.Sequence
	Cf6 bool
	Cf7 _dafny.Sequence
	Cf8 _dafny.Int
	Cf9 _dafny.CodePoint
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 _dafny.Sequence, Cf6 bool, Cf7 _dafny.Sequence, Cf8 _dafny.Int, Cf9 _dafny.CodePoint) D0 {
	return D0{D0_DC2{Cf5, Cf6, Cf7, Cf8, Cf9}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_() D0 {
	return D0{D0_DC3{}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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

type D0_DC4 struct {
	Cf10 D0
}

func (D0_DC4) isD0() {}

func (CompanionStruct_D0_) Create_DC4_(Cf10 D0) D0 {
	return D0{D0_DC4{Cf10}}
}

func (_this D0) Is_DC4() bool {
	_, ok := _this.Get_().(D0_DC4)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.EmptySeq, false, false, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf8
}

func (_this D0) Dtor_cf9() _dafny.CodePoint {
	return _this.Get_().(D0_DC2).Cf9
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf10() D0 {
	return _this.Get_().(D0_DC4).Cf10
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + data.Cf1.VerbatimString(true) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + data.Cf5.VerbatimString(true) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC4:
		{
			return "D0.DC4" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf1.Equals(data2.Cf1) && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3 && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6 && data1.Cf7.Equals(data2.Cf7) && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9
		}
	case D0_DC3:
		{
			_, ok := other.Get_().(D0_DC3)
			return ok
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
		}
	case D0_DC4:
		{
			data2, ok := other.Get_().(D0_DC4)
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

type D1_DC6 struct {
	Cf12 _dafny.Int
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf12 _dafny.Int) D1 {
	return D1{D1_DC6{Cf12}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC7 struct {
}

func (D1_DC7) isD1() {}

func (CompanionStruct_D1_) Create_DC7_() D1 {
	return D1{D1_DC7{}}
}

func (_this D1) Is_DC7() bool {
	_, ok := _this.Get_().(D1_DC7)
	return ok
}

type D1_DC5 struct {
	Cf11 _dafny.Map
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf11 _dafny.Map) D1 {
	return D1{D1_DC5{Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC6_(_dafny.Zero)
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf12
}

func (_this D1) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC7:
		{
			return "D1.DC7"
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
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D1_DC7:
		{
			_, ok := other.Get_().(D1_DC7)
			return ok
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

type D2_DC8 struct {
	Cf13 _dafny.Map
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Map) D2 {
	return D2{D2_DC8{Cf13}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D2) Dtor_cf13() _dafny.Map {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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

type D3_DC9 struct {
	Cf14 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf14 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf14}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D3) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf14) + ")"
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
			return ok && data1.Cf14.Equals(data2.Cf14)
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
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_() D4 {
	return D4{D4_DC11{}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC12 struct {
	Cf16 _dafny.CodePoint
	Cf17 _dafny.MultiSet
	Cf18 bool
	Cf19 _dafny.Int
	Cf20 *C0
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf16 _dafny.CodePoint, Cf17 _dafny.MultiSet, Cf18 bool, Cf19 _dafny.Int, Cf20 *C0) D4 {
	return D4{D4_DC12{Cf16, Cf17, Cf18, Cf19, Cf20}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC10 struct {
	Cf15 _dafny.Map
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf15 _dafny.Map) D4 {
	return D4{D4_DC10{Cf15}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_()
}

func (_this D4) Dtor_cf16() _dafny.CodePoint {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() _dafny.MultiSet {
	return _this.Get_().(D4_DC12).Cf17
}

func (_this D4) Dtor_cf18() bool {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf20() *C0 {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf15() _dafny.Map {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf15) + ")"
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
			_, ok := other.Get_().(D4_DC11)
			return ok
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf16 == data2.Cf16 && data1.Cf17.Equals(data2.Cf17) && data1.Cf18 == data2.Cf18 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
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

type D5_DC13 struct {
	Cf21 _dafny.Array
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf21 _dafny.Array) D5 {
	return D5{D5_DC13{Cf21}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf21 == data2.Cf21
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
	F4  _dafny.Int
	F5  bool
	_f0 bool
	_f2 _dafny.Array
	_f3 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = false
	_this.F4 = _dafny.Zero
	_this.F5 = false
	_this._f0 = false
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 _dafny.Array, f3 _dafny.Int, f4 _dafny.Int, f5 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
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
func (_this *C0) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if true {
				return _dafny.UnicodeSeqOfUtf8Bytes("ytnl")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("wfgpisr")
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qsbhxkh"), _dafny.UnicodeSeqOfUtf8Bytes("p")))
	}
}
func (_this *C0) Fm3(globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.IntOfInt64(376)).Times(_dafny.IntOfInt64(-205)))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
