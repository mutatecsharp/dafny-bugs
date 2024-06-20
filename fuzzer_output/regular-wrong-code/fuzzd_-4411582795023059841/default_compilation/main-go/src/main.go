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
	return (_dafny.IntOfInt64(-427)).Times(_dafny.IntOfInt64(806))
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	return ((_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)))).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, !(true), true)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.IntOfInt64(366)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-655))).Cardinality()) < 0, (_dafny.IntOfInt64(252)).Cmp(_dafny.IntOfInt64(907)) != 0, !((_dafny.SetOf(_dafny.IntOfInt64(424))).IsSubsetOf(func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg0 _dafny.Int) interface{} {
					return coer0(arg0)
				}
			}(func(_0_i0 _dafny.Int) _dafny.Int {
				return (_dafny.IntOfInt64(-901))
			}))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _1_v0 _dafny.Int
				_1_v0 = interface{}(_compr_1).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg1 _dafny.Int) interface{} {
						return coer1(arg1)
					}
				}(func(_0_i0 _dafny.Int) _dafny.Int {
					return (_dafny.IntOfInt64(-901))
				})), _1_v0) {
					_coll1.Add(Companion_Default___.SafeModuloInt(_1_v0, _dafny.IntOfInt64(949)), true)
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
				for _iter2 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_0_i0 _dafny.Int) _dafny.Int {
					return (_dafny.IntOfInt64(-901))
				}))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _1_v0 _dafny.Int
					_1_v0 = interface{}(_compr_2).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg3 _dafny.Int) interface{} {
							return coer3(arg3)
						}
					}(func(_0_i0 _dafny.Int) _dafny.Int {
						return (_dafny.IntOfInt64(-901))
					})), _1_v0) {
						_coll2.Add(Companion_Default___.SafeModuloInt(_1_v0, _dafny.IntOfInt64(949)), true)
					}
				}
				return _coll2.ToMap()
			}()).Contains(_2_v1) {
				_coll0.Add((_2_v1).Minus(_dafny.IntOfUint32(((Companion_D1_.Create_DC3_(_dafny.SeqOf(_dafny.IntOfInt64(57), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality()), true)).Cardinality())), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality()))).Dtor_cf4()).Cardinality())))
			}
		}
		return _coll0.ToSet()
	}())), !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('k')), _dafny.CodePoint('y')), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(403), _dafny.IntOfInt64(754))).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wcdmhndul"), _dafny.UnicodeSeqOfUtf8Bytes("cf"), _dafny.UnicodeSeqOfUtf8Bytes("vdotcjy"))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC1_(!(true))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-626)), _dafny.SeqOf(_dafny.IntOfInt64(-882), (_dafny.SetOf(false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("alqfd")).Cardinality()), !(!(!(true)))))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdcr")).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-427))))), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vc")).Cardinality())).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ikfvgwhey")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("isjymm")).Cardinality())).Plus(_dafny.IntOfInt64(687)), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()))).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfInt64(381), _dafny.IntOfInt64(900), (_dafny.SetOf(!(false), false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), false)).Merge((Companion_D5_.Create_DC12_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Dtor_cf14())
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(679), _dafny.IntOfInt64(756), _dafny.IntOfInt64(9))).Difference(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("apyq")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('t'))).Cardinality()), _dafny.IntOfInt64(272))).Cardinality(), (func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('k')))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _4_v0 _dafny.Sequence
			_4_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('k'))), _4_v0) {
				_coll3.Add(_4_v0, true)
			}
		}
		return _coll3.ToMap()
	}()).Cardinality()))).Intersection(func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("st")).Cardinality()))).Cardinality())), true)).Keys().Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _5_v1 _dafny.Int
			_5_v1 = interface{}(_compr_4).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("st")).Cardinality()))).Cardinality())), true)).Contains(_5_v1) {
				_coll4.Add(Companion_Default___.SafeModuloInt(_5_v1, (_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())))
			}
		}
		return _coll4.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('t')
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(364), (_dafny.SetOf(true, !(true))).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-803))).Cardinality()), _dafny.IntOfInt64(-364)), _dafny.SeqOf(_dafny.IntOfInt64(290), _dafny.IntOfInt64(-495))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.SetOf(_dafny.CodePoint('f'))).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-600), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(_dafny.IntOfInt64(929)))).Cardinality())).Cardinality()), _dafny.IntOfInt64(616))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var _hi0 _dafny.Int = _dafny.IntOfInt64(-648)
	_ = _hi0
	for _6_i0 := p1; _6_i0.Cmp(_hi0) < 0; _6_i0 = _6_i0.Plus(_dafny.One) {
		var _7_v0 _dafny.Sequence
		_ = _7_v0
		_7_v0 = _dafny.SeqOf(p1, Companion_Default___.Fm0(globalState), _6_i0)
		r2 = (_7_v0).Select((Companion_Default___.SafeIndex(_6_i0, _dafny.IntOfUint32((_7_v0).Cardinality()))).Uint32()).(_dafny.Int)
		if (_6_i0).Cmp(p1) != 0 {
			var _8_v1 _dafny.CodePoint
			_ = _8_v1
			_8_v1 = _dafny.CodePoint('h')
			_8_v1 = _8_v1
			var _9_v2 _dafny.Array
			_ = _9_v2
			var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
			_ = _nw0
			_9_v2 = _nw0
			var _10_v3 _dafny.Array
			_ = _10_v3
			var _nwElement0_0 _dafny.Array = _9_v2
			_ = _nwElement0_0
			var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(13))
			_ = _nw1
			(_nw1).ArraySet1(_nwElement0_0, 0)
			(_nw1).ArraySet1(_9_v2, 1)
			(_nw1).ArraySet1(_9_v2, 2)
			(_nw1).ArraySet1(_9_v2, 3)
			(_nw1).ArraySet1(_9_v2, 4)
			(_nw1).ArraySet1(_9_v2, 5)
			(_nw1).ArraySet1(_9_v2, 6)
			(_nw1).ArraySet1(_9_v2, 7)
			(_nw1).ArraySet1(_9_v2, 8)
			(_nw1).ArraySet1(_9_v2, 9)
			(_nw1).ArraySet1(_9_v2, 10)
			(_nw1).ArraySet1(_9_v2, 11)
			(_nw1).ArraySet1(_9_v2, 12)
			_10_v3 = _nw1
			var _11_v4 bool
			_ = _11_v4
			_11_v4 = false
			var _12_v5 _dafny.Map
			_ = _12_v5
			_12_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_10_v3, _11_v4)
			_12_v5 = (_12_v5).Update(_10_v3, _11_v4)
			r0 = _11_v4
			var _13_v6 _dafny.Array
			_ = _13_v6
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_0
			var _nw2 _dafny.Array
			_ = _nw2
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw2 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = (func(_14_v4 bool) func(_dafny.Int) _dafny.Int {
					return func(_15_i1 _dafny.Int) _dafny.Int {
						return (_15_i1).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_14_v4, _14_v4)).Cardinality()))
					}
				})(_11_v4)
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
			_13_v6 = _nw2
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_13_v6), 0))
			_ = _index0
			(_13_v6).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-266)), (_index0).Int())
			var _16_v7 _dafny.Sequence
			_ = _16_v7
			_16_v7 = _dafny.SeqOf(_11_v4)
			var _17_v8 _dafny.MultiSet
			_ = _17_v8
			_17_v8 = _dafny.MultiSetOf((_dafny.Zero).Minus(p1))
			var _18_v9 _dafny.Map
			_ = _18_v9
			_18_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v6, _17_v8)
			var _19_v10 _dafny.Sequence
			_ = _19_v10
			_19_v10 = _dafny.SeqOf(_11_v4, (_11_v4) == (_11_v4), (_dafny.MultiSetOf(_dafny.IntOfUint32((_16_v7).Cardinality()))).IsDisjointFrom((func() _dafny.MultiSet {
				if (_18_v9).Contains(_13_v6) {
					return (_18_v9).Get(_13_v6).(_dafny.MultiSet)
				}
				return (_17_v8).Update(_6_i0, Companion_Default___.Abs(_6_i0))
			})()), _11_v4)
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_13_v6), 0))
			_ = _index1
			(_13_v6).ArraySet1(_dafny.IntOfUint32((_19_v10).Cardinality()), (_index1).Int())
			var _20_v11 _dafny.Map
			_ = _20_v11
			_20_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_11_v4, _17_v8)
			_20_v11 = (_20_v11).Update(_11_v4, _17_v8)
		} else {
			var _21_v12 _dafny.MultiSet
			_ = _21_v12
			_21_v12 = _dafny.MultiSetOf(p1, _6_i0, _6_i0)
			var _22_v13 _dafny.Map
			_ = _22_v13
			_22_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _21_v12)
			var _23_v14 _dafny.CodePoint
			_ = _23_v14
			_23_v14 = _dafny.CodePoint('f')
			var _24_v15 _dafny.MultiSet
			_ = _24_v15
			_24_v15 = _dafny.MultiSetOf(_23_v14, _dafny.CodePoint('a'), _dafny.CodePoint('x'))
			_22_v13 = (_22_v13).Update(p1, _dafny.MultiSetOf(p1, (func() _dafny.Int {
				if (_24_v15).Contains(_dafny.CodePoint('b')) {
					return (_24_v15).Multiplicity(_dafny.CodePoint('b'))
				}
				return _6_i0
			})()))
			var _25_v16 _dafny.Array
			_ = _25_v16
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw3
			_25_v16 = _nw3
			var _26_v17 _dafny.MultiSet
			_ = _26_v17
			_26_v17 = _dafny.MultiSetOf(_25_v16)
			_26_v17 = (_26_v17).Union(_dafny.MultiSetOf(_25_v16, _25_v16))
			var _27_v18 *C0
			_ = _27_v18
			var _nw4 *C0 = New_C0_()
			_ = _nw4
			_nw4.Ctor__(_dafny.CodePoint('t'))
			_27_v18 = _nw4
			_27_v18 = _27_v18
			var _28_v19 _dafny.Array
			_ = _28_v19
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw5
			_28_v19 = _nw5
			var _29_v20 bool
			_ = _29_v20
			_29_v20 = false
			var _30_v21 _dafny.Map
			_ = _30_v21
			_30_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_27_v18, !(_29_v20))
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_28_v19), 0))
			_ = _index2
			(_28_v19).ArraySet1(!(_30_v21).Contains(_27_v18), (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_28_v19), 0))
			_ = _index3
			(_28_v19).ArraySet1(_29_v20, (_index3).Int())
			var _31_v22 _dafny.Array
			_ = _31_v22
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_1
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Set = (func(_32_v20 bool, _33_v19 _dafny.Array) func(_dafny.Int) _dafny.Set {
					return func(_34_i2 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_32_v20)).Union(_dafny.SetOf((_33_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_33_v19), 0))).Int()).(bool), _32_v20, (_33_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_33_v19), 0))).Int()).(bool)))
					}
				})(_29_v20, _28_v19)
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
			_31_v22 = _nw6
			_31_v22 = _31_v22
		}
		var _35_v23 bool
		_ = _35_v23
		_35_v23 = false
		var _36_v24 _dafny.Map
		_ = _36_v24
		_36_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v23, p1)
		var _37_v25 _dafny.Sequence
		_ = _37_v25
		_37_v25 = _dafny.SeqOf(_35_v23)
		var _38_v26 _dafny.Map
		_ = _38_v26
		_38_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_6_i0, _dafny.IntOfInt64(-755))
		var _39_v27 _dafny.Sequence
		_ = _39_v27
		_39_v27 = _dafny.SeqOf(_38_v26, _38_v26)
		var _40_v28 _dafny.Map
		_ = _40_v28
		_40_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_39_v27).Cardinality()), _7_v0)
		var _41_v29 _dafny.Array
		_ = _41_v29
		var _nwElement0_1 _dafny.Int = p1
		_ = _nwElement0_1
		var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(27))
		_ = _nw7
		(_nw7).ArraySet1(_nwElement0_1, 0)
		(_nw7).ArraySet1(_6_i0, 1)
		(_nw7).ArraySet1(_6_i0, 2)
		(_nw7).ArraySet1(p1, 3)
		(_nw7).ArraySet1(_6_i0, 4)
		(_nw7).ArraySet1((func() _dafny.Int {
			if (_36_v24).Contains(_35_v23) {
				return (_36_v24).Get(_35_v23).(_dafny.Int)
			}
			return _6_i0
		})(), 5)
		(_nw7).ArraySet1(((_dafny.SetOf(p1)).Cardinality()).Plus(_6_i0), 6)
		(_nw7).ArraySet1((p1).Minus(_dafny.IntOfInt64(206)), 7)
		(_nw7).ArraySet1(p1, 8)
		(_nw7).ArraySet1((p1).Minus(p1), 9)
		(_nw7).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_37_v25, _37_v25)).Cardinality()), 10)
		(_nw7).ArraySet1(_dafny.IntOfInt64(-159), 11)
		(_nw7).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p1, p1)), 12)
		(_nw7).ArraySet1(_6_i0, 13)
		(_nw7).ArraySet1(((_40_v28).Update(p1, _dafny.Companion_Sequence_.Update(_7_v0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_7_v0).Cardinality()))).Uint32(), _6_i0))).Cardinality(), 14)
		(_nw7).ArraySet1((_dafny.IntOfInt64(750)).Minus(_6_i0), 15)
		(_nw7).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p1, p1, (_dafny.Zero).Minus((_dafny.Zero).Minus(p1)))).Cardinality()), 16)
		(_nw7).ArraySet1(_6_i0, 17)
		(_nw7).ArraySet1((_dafny.Zero).Minus((Companion_Default___.Fm9(_35_v23, _35_v23, _35_v23, globalState)).Cardinality()), 18)
		(_nw7).ArraySet1(p1, 19)
		(_nw7).ArraySet1(p1, 20)
		(_nw7).ArraySet1((_6_i0).Minus(_6_i0), 21)
		(_nw7).ArraySet1((_dafny.Zero).Minus(_6_i0), 22)
		(_nw7).ArraySet1((_6_i0).Minus(_dafny.IntOfInt64(-539)), 23)
		(_nw7).ArraySet1(_6_i0, 24)
		(_nw7).ArraySet1(_6_i0, 25)
		(_nw7).ArraySet1(p1, 26)
		_41_v29 = _nw7
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_41_v29), 0))
		_ = _index4
		(_41_v29).ArraySet1(_6_i0, (_index4).Int())
		var _42_v30 _dafny.Map
		_ = _42_v30
		_42_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_7_v0, _35_v23)
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_41_v29), 0))
		_ = _index5
		var _rhs0 _dafny.Int = _6_i0
		_ = _rhs0
		var _rhs1 _dafny.Map = _42_v30
		_ = _rhs1
		var _lhs0 _dafny.Array = _41_v29
		_ = _lhs0
		var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_41_v29), 0))
		_ = _lhs1
		(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
		_42_v30 = _rhs1
		var _43_v31 _dafny.CodePoint
		_ = _43_v31
		_43_v31 = _dafny.CodePoint('t')
		var _44_v32 T0
		_ = _44_v32
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__(_43_v31)
		_44_v32 = _nw8
		var _45_v33 _dafny.MultiSet
		_ = _45_v33
		_45_v33 = _dafny.MultiSetOf(_44_v32)
		var _46_v34 _dafny.MultiSet
		_ = _46_v34
		_46_v34 = _dafny.MultiSetOf((_45_v33).Cardinality(), _dafny.IntOfUint32((_7_v0).Cardinality()))
		var _47_v35 _dafny.Map
		_ = _47_v35
		_47_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_35_v23, !((_37_v25).Select((Companion_Default___.SafeIndex((_46_v34).Cardinality(), _dafny.IntOfUint32((_37_v25).Cardinality()))).Uint32()).(bool)))
		_47_v35 = (_47_v35).Update(_35_v23, _35_v23)
	}
	var _48_v36 _dafny.CodePoint
	_ = _48_v36
	_48_v36 = _dafny.CodePoint('w')
	_48_v36 = _48_v36
	var _49_v37 bool
	_ = _49_v37
	_49_v37 = true
	var _50_v38 _dafny.Sequence
	_ = _50_v38
	_50_v38 = _dafny.UnicodeSeqOfUtf8Bytes("ejbib")
	var _51_v39 _dafny.MultiSet
	_ = _51_v39
	_51_v39 = _dafny.MultiSetOf(p1, _dafny.IntOfUint32((_50_v38).Cardinality()))
	var _52_v40 _dafny.Sequence
	_ = _52_v40
	_52_v40 = _dafny.SeqOf(_49_v37)
	var _53_v41 _dafny.Map
	_ = _53_v41
	_53_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_52_v40).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_52_v40).Cardinality()))).Uint32()).(bool))
	var _54_v42 _dafny.Sequence
	_ = _54_v42
	_54_v42 = _dafny.SeqOf(p1, (_dafny.SetOf(_49_v37, _49_v37, Companion_Default___.Fm1(_dafny.SeqOf(_51_v39), _49_v37, globalState))).Cardinality(), (p0).Cardinality(), p1, (_53_v41).Cardinality())
	var _55_v43 _dafny.Array
	_ = _55_v43
	var _nwElement0_2 _dafny.Int = p1
	_ = _nwElement0_2
	var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(13))
	_ = _nw9
	(_nw9).ArraySet1(_nwElement0_2, 0)
	(_nw9).ArraySet1(p1, 1)
	(_nw9).ArraySet1((p1).Times(p1), 2)
	(_nw9).ArraySet1(p1, 3)
	(_nw9).ArraySet1((Companion_Default___.Fm10(globalState)).Cardinality(), 4)
	(_nw9).ArraySet1(p1, 5)
	(_nw9).ArraySet1(_dafny.IntOfInt64(-831), 6)
	(_nw9).ArraySet1((p1).Plus(p1), 7)
	(_nw9).ArraySet1(p1, 8)
	(_nw9).ArraySet1((_54_v42).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_54_v42).Cardinality()))).Uint32()).(_dafny.Int), 9)
	(_nw9).ArraySet1(Companion_Default___.SafeModuloInt(p1, (_dafny.Zero).Minus((_53_v41).Cardinality())), 10)
	(_nw9).ArraySet1(p1, 11)
	(_nw9).ArraySet1(p1, 12)
	_55_v43 = _nw9
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))
	_ = _index6
	(_55_v43).ArraySet1((p1).Times(p1), (_index6).Int())
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))
	_ = _index7
	(_55_v43).ArraySet1(p1, (_index7).Int())
	r2 = (_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)
	var _56_v44 _dafny.Array
	_ = _56_v44
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(27)
	_ = _len0_2
	var _nw10 _dafny.Array
	_ = _nw10
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw10 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) bool = (func(_57_v37 bool) func(_dafny.Int) bool {
			return func(_58_i3 _dafny.Int) bool {
				return _57_v37
			}
		})(_49_v37)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw10 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw10).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw10).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_56_v44 = _nw10
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
	_ = _index8
	(_56_v44).ArraySet1(_49_v37, (_index8).Int())
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))
	_ = _index9
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
	_ = _index10
	var _rhs2 bool = Companion_Default___.Fm1(_dafny.SeqOf(_51_v39), _49_v37, globalState)
	_ = _rhs2
	var _rhs3 _dafny.Sequence = _52_v40
	_ = _rhs3
	var _rhs4 _dafny.Int = p1
	_ = _rhs4
	var _rhs5 _dafny.Int = p1
	_ = _rhs5
	var _rhs6 bool = _49_v37
	_ = _rhs6
	var _lhs2 *GlobalState = globalState
	_ = _lhs2
	var _lhs3 _dafny.Array = _55_v43
	_ = _lhs3
	var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))
	_ = _lhs4
	var _lhs5 _dafny.Array = _56_v44
	_ = _lhs5
	var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
	_ = _lhs6
	_lhs2.F0 = _rhs2
	_52_v40 = _rhs3
	r2 = _rhs4
	(_lhs3).ArraySet1(_rhs5, (_lhs4).Int())
	(_lhs5).ArraySet1(_rhs6, (_lhs6).Int())
	var _59_v45 D1
	_ = _59_v45
	_59_v45 = Companion_D1_.Create_DC4_()
	var _pat_let_tv0 = _55_v43
	_ = _pat_let_tv0
	var _pat_let_tv1 = _55_v43
	_ = _pat_let_tv1
	var _pat_let_tv2 = _51_v39
	_ = _pat_let_tv2
	var _pat_let_tv3 = _51_v39
	_ = _pat_let_tv3
	var _pat_let_tv4 = _55_v43
	_ = _pat_let_tv4
	var _pat_let_tv5 = _55_v43
	_ = _pat_let_tv5
	var _pat_let_tv6 = _49_v37
	_ = _pat_let_tv6
	var _pat_let_tv7 = _49_v37
	_ = _pat_let_tv7
	var _pat_let_tv8 = _49_v37
	_ = _pat_let_tv8
	if func(_source0 D1) bool {
		if _source0.Is_DC2() {
			var _60___mcc_h0 _dafny.Int = _source0.Get_().(D1_DC2).Cf2
			_ = _60___mcc_h0
			var _61___mcc_h1 bool = _source0.Get_().(D1_DC2).Cf3
			_ = _61___mcc_h1
			var _62_cf3 bool = _61___mcc_h1
			_ = _62_cf3
			var _63_cf2 _dafny.Int = _60___mcc_h0
			_ = _63_cf2
			return _62_cf3
		} else if _source0.Is_DC3() {
			var _64___mcc_h2 _dafny.Sequence = _source0.Get_().(D1_DC3).Cf4
			_ = _64___mcc_h2
			var _65___mcc_h3 _dafny.Int = _source0.Get_().(D1_DC3).Cf5
			_ = _65___mcc_h3
			var _66_cf5 _dafny.Int = _65___mcc_h3
			_ = _66_cf5
			var _67_cf4 _dafny.Sequence = _64___mcc_h2
			_ = _67_cf4
			return (_dafny.SetOf((_pat_let_tv1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_pat_let_tv0), 0))).Int()).(_dafny.Int))).IsDisjointFrom(_dafny.SetOf((_pat_let_tv2).Cardinality(), (_pat_let_tv3).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sp")).Cardinality()), _dafny.IntOfUint32((_67_cf4).Cardinality()), (_pat_let_tv5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_pat_let_tv4), 0))).Int()).(_dafny.Int)))
		} else if _source0.Is_DC4() {
			if true {
				return (Companion_D1_.Create_DC1_(_pat_let_tv6)).Dtor_cf1()
			} else {
				return _pat_let_tv7
			}
		} else {
			var _68___mcc_h4 bool = _source0.Get_().(D1_DC1).Cf1
			_ = _68___mcc_h4
			var _69_cf1 bool = _68___mcc_h4
			_ = _69_cf1
			return _pat_let_tv8
		}
	}(_59_v45) {
		_56_v44 = _56_v44
		_48_v36 = Companion_Default___.Fm11(_dafny.Companion_Sequence_.Concatenate(_54_v42, _54_v42), p1, _dafny.IntOfInt64(766), (func() bool {
			if (_56_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))).Int()).(bool) {
				return _49_v37
			}
			return _49_v37
		})(), globalState)
		var _70_v46 *C0
		_ = _70_v46
		var _nw11 *C0 = New_C0_()
		_ = _nw11
		_nw11.Ctor__(_48_v36)
		_70_v46 = _nw11
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
		_ = _index11
		(_56_v44).ArraySet1((func() bool {
			if (_70_v46).Fm3(globalState) {
				return _49_v37
			}
			return _49_v37
		})(), (_index11).Int())
		r1 = (p1).Minus(_dafny.IntOfInt64(440))
	} else {
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
		_ = _index12
		(_56_v44).ArraySet1(!((_56_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))).Int()).(bool)), (_index12).Int())
		var _71_v47 *C0
		_ = _71_v47
		var _nw12 *C0 = New_C0_()
		_ = _nw12
		_nw12.Ctor__(_48_v36)
		_71_v47 = _nw12
		var _72_v48 _dafny.Sequence
		_ = _72_v48
		_72_v48 = _dafny.SeqOf(_71_v47)
		var _73_v49 _dafny.Map
		_ = _73_v49
		_73_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_72_v48, _dafny.SeqOf(_71_v47, _71_v47)), (_56_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))).Int()).(bool))
		_73_v49 = (_73_v49).Update(_72_v48, false)
		_49_v37 = (p1).Cmp(((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)).Times(p1)) <= 0
		if _49_v37 {
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
			_ = _index13
			(_56_v44).ArraySet1((_56_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))).Int()).(bool), (_index13).Int())
			var _74_v50 _dafny.Array
			_ = _74_v50
			var _nw13 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
			_ = _nw13
			_74_v50 = _nw13
			var _75_v51 T0
			_ = _75_v51
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_dafny.CodePoint('q'))
			_75_v51 = _nw14
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_74_v50), 0))
			_ = _index14
			(_74_v50).ArraySet1(_75_v51, (_index14).Int())
			var _76_v52 D2
			_ = _76_v52
			_76_v52 = Companion_D2_.Create_DC6_(_75_v51)
			var _77_v53 _dafny.Sequence
			_ = _77_v53
			_77_v53 = _dafny.SeqOf(_50_v38)
			var _78_v54 D4
			_ = _78_v54
			_78_v54 = Companion_D4_.Create_DC10_((_77_v53).Select((Companion_Default___.SafeIndex((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_77_v53).Cardinality()))).Uint32()).(_dafny.Sequence))
			var _79_v55 D4
			_ = _79_v55
			_79_v55 = Companion_D4_.Create_DC10_((_78_v54).Dtor_cf11())
			var _80_v56 _dafny.Map
			_ = _80_v56
			_80_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v37, (_78_v54).Dtor_cf11())
			var _81_v57 D1
			_ = _81_v57
			_81_v57 = Companion_D1_.Create_DC1_((_56_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))).Int()).(bool))
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_74_v50), 0))
			_ = _index15
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
			_ = _index16
			var _rhs7 T0 = (_76_v52).Dtor_cf7()
			_ = _rhs7
			var _rhs8 bool = _dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
				if (_80_v56).Contains(_49_v37) {
					return (_80_v56).Get(_49_v37).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("pcctiawd")
			})(), Companion_Default___.Fm8(_dafny.SeqOf(p1, (_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)), globalState))
			_ = _rhs8
			var _rhs9 bool = !((_81_v57).Dtor_cf1())
			_ = _rhs9
			var _lhs7 _dafny.Array = _74_v50
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(611), _dafny.ArrayLen((_74_v50), 0))
			_ = _lhs8
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			var _lhs10 _dafny.Array = _56_v44
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
			_ = _lhs11
			(_lhs7).ArraySet1(_rhs7, (_lhs8).Int())
			_lhs9.F0 = _rhs8
			(_lhs10).ArraySet1(_rhs9, (_lhs11).Int())
			var _82_v58 _dafny.Map
			_ = _82_v58
			_82_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v45, (_dafny.Zero).Minus((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)))
			var _83_v59 _dafny.Array
			_ = _83_v59
			var _nwElement0_3 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_50_v38, Companion_Default___.Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(79))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}((func(_84_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_85_i4 _dafny.Int) _dafny.Int {
					return _84_p1
				}
			})(p1))), globalState))
			_ = _nwElement0_3
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(26))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_3, 0)
			(_nw15).ArraySet1(_50_v38, 1)
			(_nw15).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(684))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_86_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_87_i5 _dafny.Int) _dafny.CodePoint {
					return _86_v36
				}
			})(_48_v36))), 2)
			(_nw15).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(21))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_88_v51 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_89_i6 _dafny.Int) _dafny.CodePoint {
					return (_88_v51).F3()
				}
			})(_75_v51))), (Companion_Default___.SafeIndex((_82_v58).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(21))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_90_v51 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_91_i6 _dafny.Int) _dafny.CodePoint {
					return (_90_v51).F3()
				}
			})(_75_v51)))).Cardinality()))).Uint32(), (_75_v51).F3()), 3)
			(_nw15).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(103))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_92_v51 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_93_i7 _dafny.Int) _dafny.CodePoint {
					return (_92_v51).F3()
				}
			})(_75_v51))), _50_v38), 4)
			(_nw15).ArraySet1(_50_v38, 5)
			(_nw15).ArraySet1(_50_v38, 6)
			(_nw15).ArraySet1(_50_v38, 7)
			(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("w"), 8)
			(_nw15).ArraySet1(_dafny.Companion_Sequence_.Update(_50_v38, (Companion_Default___.SafeIndex((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_50_v38).Cardinality()))).Uint32(), _48_v36), 9)
			(_nw15).ArraySet1(_50_v38, 10)
			(_nw15).ArraySet1(_50_v38, 11)
			(_nw15).ArraySet1(_dafny.Companion_Sequence_.Update(_50_v38, (Companion_Default___.SafeIndex((_54_v42).Select((Companion_Default___.SafeIndex((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_54_v42).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_50_v38).Cardinality()))).Uint32(), (_75_v51).F3()), 12)
			(_nw15).ArraySet1(Companion_Default___.Fm8(_54_v42, globalState), 13)
			(_nw15).ArraySet1((func() _dafny.Sequence {
				if _49_v37 {
					return _50_v38
				}
				return Companion_Default___.Fm8(Companion_Default___.Fm12(Companion_Default___.Fm0(globalState), p1, globalState), globalState)
			})(), 14)
			(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jumanw"), 15)
			(_nw15).ArraySet1(_50_v38, 16)
			(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("f"), 17)
			(_nw15).ArraySet1(_50_v38, 18)
			(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kyeeyllc"), 19)
			(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("k"), 20)
			(_nw15).ArraySet1(_50_v38, 21)
			(_nw15).ArraySet1(_50_v38, 22)
			(_nw15).ArraySet1(_50_v38, 23)
			(_nw15).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hqetjr"), 24)
			(_nw15).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(80))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_94_v51 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_95_i8 _dafny.Int) _dafny.CodePoint {
					return (_94_v51).F3()
				}
			})(_75_v51))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(392))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_96_v51 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_97_i9 _dafny.Int) _dafny.CodePoint {
					return (_96_v51).F3()
				}
			})(_75_v51)))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(80))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}((func(_98_v51 T0) func(_dafny.Int) _dafny.CodePoint {
				return func(_99_i8 _dafny.Int) _dafny.CodePoint {
					return (_98_v51).F3()
				}
			})(_75_v51)))).Cardinality()))).Uint32(), _48_v36), 25)
			_83_v59 = _nw15
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_83_v59), 0))
			_ = _index17
			(_83_v59).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mqdbtkjg"), (_index17).Int())
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_83_v59), 0))
			_ = _index18
			(_83_v59).ArraySet1(_50_v38, (_index18).Int())
			r1 = (_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))
			_ = _index19
			(_56_v44).ArraySet1(true, (_index19).Int())
		} else {
			r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_50_v38, _50_v38), _50_v38)).Cardinality())
			var _100_v60 _dafny.Array
			_ = _100_v60
			var _nwElement0_4 *C0 = (func() *C0 {
				if (_56_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_56_v44), 0))).Int()).(bool) {
					return _71_v47
				}
				return _71_v47
			})()
			_ = _nwElement0_4
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(16))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_4, 0)
			(_nw16).ArraySet1(_71_v47, 1)
			(_nw16).ArraySet1(_71_v47, 2)
			(_nw16).ArraySet1(_71_v47, 3)
			(_nw16).ArraySet1(_71_v47, 4)
			(_nw16).ArraySet1(_71_v47, 5)
			(_nw16).ArraySet1((func() *C0 {
				if !(_49_v37) {
					return _71_v47
				}
				return _71_v47
			})(), 6)
			(_nw16).ArraySet1(_71_v47, 7)
			(_nw16).ArraySet1(_71_v47, 8)
			(_nw16).ArraySet1(_71_v47, 9)
			(_nw16).ArraySet1(_71_v47, 10)
			(_nw16).ArraySet1(_71_v47, 11)
			(_nw16).ArraySet1(_71_v47, 12)
			(_nw16).ArraySet1(_71_v47, 13)
			(_nw16).ArraySet1(_71_v47, 14)
			(_nw16).ArraySet1(_71_v47, 15)
			_100_v60 = _nw16
			_100_v60 = _100_v60
			_50_v38 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(128))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}((func(_101_v36 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_102_i10 _dafny.Int) _dafny.CodePoint {
					return _101_v36
				}
			})(_48_v36)))
			var _103_v61 _dafny.Map
			_ = _103_v61
			_103_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v37, (_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int))
			var _104_v62 _dafny.Map
			_ = _104_v62
			_104_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _55_v43)
			r2 = ((_103_v61).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v37, p1)).Update(_49_v37, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v44, (_104_v62).Cardinality())).Cardinality())))).Cardinality()
			r2 = (_dafny.Zero).Minus(((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)).Minus((_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)))
		}
	}
	r0 = _49_v37
	r1 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(455), (_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int))
	r2 = (_55_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_55_v43), 0))).Int()).(_dafny.Int)
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _105_globalState *GlobalState
	_ = _105_globalState
	var _nw17 *GlobalState = New_GlobalState_()
	_ = _nw17
	_nw17.Ctor__(true, true, false)
	_105_globalState = _nw17
	var _106_v0 _dafny.Sequence
	_ = _106_v0
	_106_v0 = _dafny.UnicodeSeqOfUtf8Bytes("jract")
	_106_v0 = _dafny.Companion_Sequence_.Concatenate(_106_v0, _106_v0)
	var _107_v1 _dafny.Int
	_ = _107_v1
	_107_v1 = _dafny.IntOfInt64(-27)
	var _108_v2 _dafny.Map
	_ = _108_v2
	_108_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v1, (_dafny.Zero).Minus(_107_v1))
	_108_v2 = (_108_v2).Update(_107_v1, _107_v1)
	_107_v1 = _107_v1
	var _109_v3 bool
	_ = _109_v3
	_109_v3 = true
	(_105_globalState).F0 = _109_v3
	(_105_globalState).F0 = true
	if _109_v3 {
		var _110_v4 _dafny.Array
		_ = _110_v4
		var _len0_3 _dafny.Int = _dafny.One
		_ = _len0_3
		var _nw18 _dafny.Array
		_ = _nw18
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw18 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = func(_111_i0 _dafny.Int) bool {
				return true
			}
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw18 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw18).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw18).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_110_v4 = _nw18
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))
		_ = _index20
		(_110_v4).ArraySet1(_109_v3, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))
		_ = _index21
		(_110_v4).ArraySet1((_107_v1).Cmp((Companion_Default___.Fm0(_105_globalState)).Minus(_107_v1)) >= 0, (_index21).Int())
		var _112_v5 _dafny.Array
		_ = _112_v5
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_4
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Map = (func(_113_v3 bool) func(_dafny.Int) _dafny.Map {
				return func(_114_i1 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v3, !(_113_v3))
				}
			})(_109_v3)
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
		_112_v5 = _nw19
		var _115_v6 _dafny.MultiSet
		_ = _115_v6
		_115_v6 = _dafny.MultiSetOf(_107_v1)
		var _116_v7 _dafny.Sequence
		_ = _116_v7
		_116_v7 = _dafny.SeqOf(_115_v6)
		var _117_v8 _dafny.MultiSet
		_ = _117_v8
		_117_v8 = _dafny.MultiSetOf((_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool))
		var _118_v9 _dafny.Map
		_ = _118_v9
		_118_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_117_v8).Cardinality(), (_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool))
		var _119_v10 _dafny.Map
		_ = _119_v10
		_119_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v5, (func() _dafny.Map {
			if Companion_Default___.Fm1(_116_v7, _109_v3, _105_globalState) {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_118_v9).Contains(_dafny.IntOfInt64(531)) {
						return (_118_v9).Get(_dafny.IntOfInt64(531)).(bool)
					}
					return _109_v3
				})(), _117_v8)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool), Companion_Default___.Fm2(_107_v1, _105_globalState))
		})())
		var _120_v11 _dafny.Map
		_ = _120_v11
		_120_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _117_v8)
		_119_v10 = (_119_v10).Update(_112_v5, _120_v11)
		var _121_v12 _dafny.Sequence
		_ = _121_v12
		_121_v12 = _dafny.SeqOf(_107_v1)
		var _122_v13 _dafny.Map
		_ = _122_v13
		_122_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v12, _107_v1)
		var _123_v14 _dafny.Set
		_ = _123_v14
		_123_v14 = _dafny.SetOf(_107_v1, (_dafny.Zero).Minus(_107_v1), _107_v1)
		var _124_v15 _dafny.Int
		_ = _124_v15
		_124_v15 = (_dafny.Zero).Minus((_123_v14).Cardinality())
		var _125_v16 _dafny.Map
		_ = _125_v16
		_125_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v1, _118_v9)
		var _126_v17 _dafny.Map
		_ = _126_v17
		_126_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool))
		var _127_v18 _dafny.Array
		_ = _127_v18
		var _nwElement0_5 _dafny.Int = (_dafny.IntOfUint32((_106_v0).Cardinality())).Times(_dafny.IntOfInt64(-941))
		_ = _nwElement0_5
		var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(17))
		_ = _nw20
		(_nw20).ArraySet1(_nwElement0_5, 0)
		(_nw20).ArraySet1(_107_v1, 1)
		(_nw20).ArraySet1(_107_v1, 2)
		(_nw20).ArraySet1((_107_v1).Times(_107_v1), 3)
		(_nw20).ArraySet1(Companion_Default___.SafeModuloInt((_122_v13).Cardinality(), _107_v1), 4)
		(_nw20).ArraySet1(Companion_Default___.SafeDivisionInt(_107_v1, _107_v1), 5)
		(_nw20).ArraySet1((_dafny.Zero).Minus(_107_v1), 6)
		(_nw20).ArraySet1(_dafny.IntOfInt64(784), 7)
		(_nw20).ArraySet1(_107_v1, 8)
		(_nw20).ArraySet1(_107_v1, 9)
		(_nw20).ArraySet1((_124_v15), 10)
		(_nw20).ArraySet1(_dafny.IntOfInt64(-619), 11)
		(_nw20).ArraySet1(_107_v1, 12)
		(_nw20).ArraySet1((_107_v1).Plus((_125_v16).Cardinality()), 13)
		(_nw20).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if (_115_v6).Contains(_107_v1) {
				return (_115_v6).Multiplicity(_107_v1)
			}
			return _107_v1
		})()), 14)
		(_nw20).ArraySet1(Companion_Default___.SafeModuloInt(_107_v1, (_126_v17).Cardinality()), 15)
		(_nw20).ArraySet1(_107_v1, 16)
		_127_v18 = _nw20
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))
		_ = _index22
		(_127_v18).ArraySet1((_107_v1).Minus(_dafny.IntOfInt64(888)), (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))
		_ = _index23
		var _rhs10 _dafny.Int = ((func() _dafny.Int {
			if (_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool) {
				return _107_v1
			}
			return (_115_v6).Cardinality()
		})()).Times(_107_v1)
		_ = _rhs10
		var _rhs11 _dafny.Map = (func() _dafny.Map {
			if ((_dafny.Zero).Minus(_107_v1)).Cmp(_107_v1) <= 0 {
				return _126_v17
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool)), (_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool))).Update((_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool), (_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool))
		})()
		_ = _rhs11
		var _rhs12 bool = (_dafny.IntOfInt64(90)).Cmp((_124_v15)) <= 0
		_ = _rhs12
		var _lhs12 _dafny.Array = _127_v18
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))
		_ = _lhs13
		var _lhs14 *GlobalState = _105_globalState
		_ = _lhs14
		(_lhs12).ArraySet1(_rhs10, (_lhs13).Int())
		_126_v17 = _rhs11
		_lhs14.F0 = _rhs12
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))
		_ = _index24
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))
		_ = _index25
		var _rhs13 _dafny.Int = Companion_Default___.Fm0(_105_globalState)
		_ = _rhs13
		var _rhs14 bool = (_110_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))).Int()).(bool)
		_ = _rhs14
		var _rhs15 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_127_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))).Int()).(_dafny.Int), (_127_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))).Int()).(_dafny.Int)), _dafny.IntOfInt64(772))
		_ = _rhs15
		var _lhs15 _dafny.Array = _127_v18
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(995), _dafny.ArrayLen((_127_v18), 0))
		_ = _lhs16
		var _lhs17 _dafny.Array = _110_v4
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_110_v4), 0))
		_ = _lhs18
		(_lhs15).ArraySet1(_rhs13, (_lhs16).Int())
		(_lhs17).ArraySet1(_rhs14, (_lhs18).Int())
		_107_v1 = _rhs15
		var _128_v19 _dafny.CodePoint
		_ = _128_v19
		_128_v19 = _dafny.CodePoint('c')
		var _129_v20 *C0
		_ = _129_v20
		var _nw21 *C0 = New_C0_()
		_ = _nw21
		_nw21.Ctor__(_128_v19)
		_129_v20 = _nw21
	} else {
		var _130_v21 _dafny.Array
		_ = _130_v21
		var _nw22 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw22
		_130_v21 = _nw22
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))
		_ = _index26
		(_130_v21).ArraySet1(_109_v3, (_index26).Int())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))
		_ = _index27
		var _rhs16 bool = _109_v3
		_ = _rhs16
		var _rhs17 _dafny.Int = Companion_Default___.Fm0(_105_globalState)
		_ = _rhs17
		var _lhs19 _dafny.Array = _130_v21
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))
		_ = _lhs20
		(_lhs19).ArraySet1(_rhs16, (_lhs20).Int())
		_107_v1 = _rhs17
		var _131_v22 _dafny.Sequence
		_ = _131_v22
		_131_v22 = _dafny.SeqOf(_109_v3)
		var _source1 D1 = Companion_Default___.Fm5(_dafny.IntOfUint32((_131_v22).Cardinality()), _105_globalState)
		_ = _source1
		if _source1.Is_DC2() {
			var _132___mcc_h0 _dafny.Int = _source1.Get_().(D1_DC2).Cf2
			_ = _132___mcc_h0
			var _133___mcc_h1 bool = _source1.Get_().(D1_DC2).Cf3
			_ = _133___mcc_h1
			var _134_cf3 bool = _133___mcc_h1
			_ = _134_cf3
			var _135_cf2 _dafny.Int = _132___mcc_h0
			_ = _135_cf2
			var _136_v23 _dafny.Set
			_ = _136_v23
			_136_v23 = _dafny.SetOf(_134_cf3, !(false), _134_cf3)
			var _137_v24 _dafny.Set
			_ = _137_v24
			_137_v24 = _dafny.SetOf(_130_v21, _130_v21)
			var _138_v25 _dafny.MultiSet
			_ = _138_v25
			_138_v25 = _dafny.MultiSetOf((_136_v23).Cardinality(), (_137_v24).Cardinality())
			_134_cf3 = (func() bool {
				if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_107_v1, _107_v1))).IsSubsetOf(_138_v25) {
					return (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool)
				}
				return _134_cf3
			})()
			_135_cf2 = (_107_v1).Times(_135_cf2)
			var _139_v26 _dafny.Array
			_ = _139_v26
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw23
			_139_v26 = _nw23
			var _140_v27 D1
			_ = _140_v27
			_140_v27 = Companion_D1_.Create_DC1_(_134_cf3)
			var _141_v28 _dafny.CodePoint
			_ = _141_v28
			_141_v28 = _dafny.CodePoint('o')
			var _142_v29 *C0
			_ = _142_v29
			var _nw24 *C0 = New_C0_()
			_ = _nw24
			_nw24.Ctor__(_141_v28)
			_142_v29 = _nw24
			var _143_v30 _dafny.Map
			_ = _143_v30
			_143_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v27, _142_v29)
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_139_v26), 0))
			_ = _index28
			(_139_v26).ArraySet1((_143_v30).Cardinality(), (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_139_v26), 0))
			_ = _index29
			(_139_v26).ArraySet1(Companion_Default___.SafeDivisionInt(_107_v1, _135_cf2), (_index29).Int())
			(_105_globalState).F0 = (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool)
		} else if _source1.Is_DC3() {
			var _144___mcc_h2 _dafny.Sequence = _source1.Get_().(D1_DC3).Cf4
			_ = _144___mcc_h2
			var _145___mcc_h3 _dafny.Int = _source1.Get_().(D1_DC3).Cf5
			_ = _145___mcc_h3
			var _146_cf5 _dafny.Int = _145___mcc_h3
			_ = _146_cf5
			var _147_cf4 _dafny.Sequence = _144___mcc_h2
			_ = _147_cf4
			(_105_globalState).F0 = _109_v3
			_109_v3 = !(_109_v3)
			_107_v1 = _dafny.IntOfInt64(988)
			_146_cf5 = _dafny.IntOfInt64(-855)
		} else if _source1.Is_DC4() {
			var _148_v31 _dafny.MultiSet
			_ = _148_v31
			_148_v31 = _dafny.MultiSetOf(_109_v3)
			var _149_v32 bool
			_ = _149_v32
			var _150_v33 _dafny.Int
			_ = _150_v33
			var _151_v34 _dafny.Int
			_ = _151_v34
			var _out0 bool
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			_out0, _out1, _out2 = Companion_Default___.M0(_148_v31, _107_v1, _105_globalState)
			_149_v32 = _out0
			_150_v33 = _out1
			_151_v34 = _out2
			_106_v0 = _dafny.Companion_Sequence_.Concatenate(_106_v0, _106_v0)
			var _152_v35 _dafny.CodePoint
			_ = _152_v35
			_152_v35 = _dafny.CodePoint('v')
			var _153_v36 *C0
			_ = _153_v36
			var _nw25 *C0 = New_C0_()
			_ = _nw25
			_nw25.Ctor__(_152_v35)
			_153_v36 = _nw25
			var _154_v37 _dafny.Set
			_ = _154_v37
			_154_v37 = _dafny.SetOf(_151_v34, _dafny.IntOfInt64(308))
			var _155_v38 _dafny.MultiSet
			_ = _155_v38
			_155_v38 = _dafny.MultiSetOf((_154_v37).Cardinality())
			var _156_v39 _dafny.Sequence
			_ = _156_v39
			_156_v39 = _dafny.SeqOf(_155_v38)
			var _157_v40 _dafny.Set
			_ = _157_v40
			_157_v40 = _dafny.SetOf((_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool), Companion_Default___.Fm1(_156_v39, (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool), _105_globalState))
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))
			_ = _index30
			var _rhs18 *C0 = _153_v36
			_ = _rhs18
			var _rhs19 bool = ((_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool)) == (!(_157_v40).Contains(!(_109_v3)))
			_ = _rhs19
			var _lhs21 _dafny.Array = _130_v21
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))
			_ = _lhs22
			_153_v36 = _rhs18
			(_lhs21).ArraySet1(_rhs19, (_lhs22).Int())
			var _158_v41 _dafny.Array
			_ = _158_v41
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw26
			_158_v41 = _nw26
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_158_v41), 0))
			_ = _index31
			(_158_v41).ArraySet1(_151_v34, (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_158_v41), 0))
			_ = _index32
			var _rhs20 _dafny.CodePoint = _152_v35
			_ = _rhs20
			var _rhs21 _dafny.Int = _150_v33
			_ = _rhs21
			var _lhs23 _dafny.Array = _158_v41
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_158_v41), 0))
			_ = _lhs24
			_152_v35 = _rhs20
			(_lhs23).ArraySet1(_rhs21, (_lhs24).Int())
		} else {
			var _159___mcc_h4 bool = _source1.Get_().(D1_DC1).Cf1
			_ = _159___mcc_h4
			var _160_cf1 bool = _159___mcc_h4
			_ = _160_cf1
			var _161_v42 _dafny.Sequence
			_ = _161_v42
			_161_v42 = _dafny.SeqOf(_107_v1, _107_v1, _107_v1, _107_v1, _107_v1)
			var _162_v43 _dafny.Set
			_ = _162_v43
			_162_v43 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_161_v42, (Companion_Default___.SafeIndex(_107_v1, _dafny.IntOfUint32((_161_v42).Cardinality()))).Uint32(), _107_v1), _161_v42)
			(_105_globalState).F0 = !(_162_v43).Equals(_162_v43)
			var _163_v44 _dafny.Array
			_ = _163_v44
			var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
			_ = _nw27
			_163_v44 = _nw27
			_163_v44 = (func() _dafny.Array {
				if (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool) {
					return _163_v44
				}
				return _163_v44
			})()
			var _164_v45 _dafny.MultiSet
			_ = _164_v45
			_164_v45 = _dafny.MultiSetOf(_109_v3, _160_cf1)
			var _165_v46 D2
			_ = _165_v46
			_165_v46 = Companion_D2_.Create_DC5_(_164_v45)
			var _166_v47 bool
			_ = _166_v47
			var _167_v48 _dafny.Int
			_ = _167_v48
			var _168_v49 _dafny.Int
			_ = _168_v49
			var _out3 bool
			_ = _out3
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			_out3, _out4, _out5 = Companion_Default___.M0(((_165_v46).Dtor_cf6()).Difference(_164_v45), _107_v1, _105_globalState)
			_166_v47 = _out3
			_167_v48 = _out4
			_168_v49 = _out5
			var _169_v50 D1
			_ = _169_v50
			_169_v50 = Companion_D1_.Create_DC2_(_107_v1, _160_cf1)
			var _rhs22 _dafny.Array = _130_v21
			_ = _rhs22
			var _rhs23 D1 = _169_v50
			_ = _rhs23
			_130_v21 = _rhs22
			_169_v50 = _rhs23
		}
		_107_v1 = (_dafny.Zero).Minus(_107_v1)
		var _170_v51 D1
		_ = _170_v51
		_170_v51 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_106_v0).Cardinality()), (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool))
		var _171_v52 _dafny.Map
		_ = _171_v52
		_171_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v1, _dafny.SeqOf(_dafny.IntOfInt64(-327), (_170_v51).Dtor_cf2()))
		var _172_v54 _dafny.Map
		_ = _172_v54
		_172_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), _107_v1)
		var _173_v55 _dafny.Sequence
		_ = _173_v55
		_173_v55 = _dafny.SeqOf(_131_v22)
		var _174_v56 _dafny.Sequence
		_ = _174_v56
		_174_v56 = _dafny.SeqOf(_dafny.IntOfUint32(((_173_v55).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_106_v0).Cardinality()), _dafny.IntOfUint32((_173_v55).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _107_v1)
		_171_v52 = (_171_v52).Update((func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((_172_v54).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _175_v53 _dafny.CodePoint
				_175_v53 = interface{}(_compr_5).(_dafny.CodePoint)
				if (_172_v54).Contains(_175_v53) {
					_coll5.Add(_175_v53, _109_v3)
				}
			}
			return _coll5.ToMap()
		}()).Cardinality(), _174_v56)
		var _176_v57 _dafny.Map
		_ = _176_v57
		_176_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v1, (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool))
		var _177_v58 _dafny.MultiSet
		_ = _177_v58
		_177_v58 = _dafny.MultiSetOf(_109_v3, (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool), (_130_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(869), _dafny.ArrayLen((_130_v21), 0))).Int()).(bool), (func() bool {
			if (_176_v57).Contains((Companion_Default___.Fm6(_105_globalState)).Dtor_cf5()) {
				return (_176_v57).Get((Companion_Default___.Fm6(_105_globalState)).Dtor_cf5()).(bool)
			}
			return _109_v3
		})())
		var _178_v59 bool
		_ = _178_v59
		var _179_v60 _dafny.Int
		_ = _179_v60
		var _180_v61 _dafny.Int
		_ = _180_v61
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		var _out8 _dafny.Int
		_ = _out8
		_out6, _out7, _out8 = Companion_Default___.M0((_177_v58).Union((_177_v58).Update(_109_v3, Companion_Default___.Abs(_107_v1))), Companion_Default___.Fm0(_105_globalState), _105_globalState)
		_178_v59 = _out6
		_179_v60 = _out7
		_180_v61 = _out8
	}
	var _hi1 _dafny.Int = _107_v1
	_ = _hi1
	for _181_i2 := _107_v1; _181_i2.Cmp(_hi1) < 0; _181_i2 = _181_i2.Plus(_dafny.One) {
		var _182_v62 _dafny.MultiSet
		_ = _182_v62
		_182_v62 = _dafny.MultiSetOf(_109_v3)
		var _183_v63 bool
		_ = _183_v63
		var _184_v64 _dafny.Int
		_ = _184_v64
		var _185_v65 _dafny.Int
		_ = _185_v65
		var _out9 bool
		_ = _out9
		var _out10 _dafny.Int
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out9, _out10, _out11 = Companion_Default___.M0(_182_v62, _107_v1, _105_globalState)
		_183_v63 = _out9
		_184_v64 = _out10
		_185_v65 = _out11
		var _186_v66 _dafny.CodePoint
		_ = _186_v66
		_186_v66 = _dafny.CodePoint('i')
		var _187_v67 T0
		_ = _187_v67
		var _nw28 *C0 = New_C0_()
		_ = _nw28
		_nw28.Ctor__(_186_v66)
		_187_v67 = _nw28
		var _188_v68 _dafny.Map
		_ = _188_v68
		_188_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC6_(_187_v67), Companion_Default___.Fm7(_105_globalState))
		var _189_v70 _dafny.Sequence
		_ = _189_v70
		_189_v70 = _dafny.SeqOf(_183_v63)
		var _190_v71 _dafny.Array
		_ = _190_v71
		var _nwElement0_6 _dafny.MultiSet = (_182_v62).Difference(_dafny.MultiSetOf(_183_v63))
		_ = _nwElement0_6
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(21))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_6, 0)
		(_nw29).ArraySet1(Companion_Default___.Fm2((func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.SetOf(_185_v65)).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _191_v69 _dafny.Int
				_191_v69 = interface{}(_compr_6).(_dafny.Int)
				if (_dafny.SetOf(_185_v65)).Contains(_191_v69) {
					_coll6.Add((_191_v69).Times(_dafny.IntOfUint32((_106_v0).Cardinality())), _183_v63)
				}
			}
			return _coll6.ToMap()
		}()).Cardinality(), _105_globalState), 1)
		(_nw29).ArraySet1(_dafny.MultiSetFromSeq(_189_v70), 2)
		(_nw29).ArraySet1((_182_v62).Update(_109_v3, Companion_Default___.Abs(_184_v64)), 3)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_183_v63, _109_v3, _109_v3, _109_v3, _183_v63), 4)
		(_nw29).ArraySet1(_182_v62, 5)
		(_nw29).ArraySet1((Companion_Default___.Fm2(_107_v1, _105_globalState)).Intersection(_182_v62), 6)
		(_nw29).ArraySet1(_182_v62, 7)
		(_nw29).ArraySet1((_182_v62).Difference(_182_v62), 8)
		(_nw29).ArraySet1(_182_v62, 9)
		(_nw29).ArraySet1(Companion_Default___.Fm2((_dafny.Zero).Minus(_dafny.IntOfUint32((_106_v0).Cardinality())), _105_globalState), 10)
		(_nw29).ArraySet1(_dafny.MultiSetFromSeq(_189_v70), 11)
		(_nw29).ArraySet1(_182_v62, 12)
		(_nw29).ArraySet1(_182_v62, 13)
		(_nw29).ArraySet1((_182_v62).Difference(_182_v62), 14)
		(_nw29).ArraySet1(_182_v62, 15)
		(_nw29).ArraySet1(_182_v62, 16)
		(_nw29).ArraySet1(_182_v62, 17)
		(_nw29).ArraySet1(_dafny.MultiSetOf(_109_v3), 18)
		(_nw29).ArraySet1(_182_v62, 19)
		(_nw29).ArraySet1(_182_v62, 20)
		_190_v71 = _nw29
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_190_v71), 0))
		_ = _index33
		(_190_v71).ArraySet1(_182_v62, (_index33).Int())
		var _192_v72 *C0
		_ = _192_v72
		var _nw30 *C0 = New_C0_()
		_ = _nw30
		_nw30.Ctor__(_186_v66)
		_192_v72 = _nw30
		var _193_v73 D2
		_ = _193_v73
		_193_v73 = Companion_D2_.Create_DC6_(_187_v67)
		var _194_v74 _dafny.Map
		_ = _194_v74
		_194_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_184_v64), _192_v72)
		var _195_v75 D1
		_ = _195_v75
		_195_v75 = Companion_D1_.Create_DC2_((_194_v74).Cardinality(), true)
		var _196_v76 _dafny.Set
		_ = _196_v76
		_196_v76 = _dafny.SetOf(_184_v64, _107_v1)
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_190_v71), 0))
		_ = _index34
		var _rhs24 _dafny.Int = (_184_v64).Minus(_184_v64)
		_ = _rhs24
		var _rhs25 _dafny.Map = (_188_v68).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v73, _195_v75))
		_ = _rhs25
		var _rhs26 _dafny.MultiSet = Companion_Default___.Fm2(_184_v64, _105_globalState)
		_ = _rhs26
		var _rhs27 bool = (_185_v65).Cmp((_dafny.Zero).Minus((_185_v65).Plus((_196_v76).Cardinality()))) > 0
		_ = _rhs27
		var _rhs28 *C0 = _192_v72
		_ = _rhs28
		var _lhs25 _dafny.Array = _190_v71
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_190_v71), 0))
		_ = _lhs26
		var _lhs27 *GlobalState = _105_globalState
		_ = _lhs27
		_107_v1 = _rhs24
		_188_v68 = _rhs25
		(_lhs25).ArraySet1(_rhs26, (_lhs26).Int())
		_lhs27.F0 = _rhs27
		_192_v72 = _rhs28
		var _197_v77 _dafny.Array
		_ = _197_v77
		var _nw31 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw31
		_197_v77 = _nw31
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_197_v77), 0))
		_ = _index35
		(_197_v77).ArraySet1(!(_109_v3), (_index35).Int())
		var _198_v78 _dafny.Sequence
		_ = _198_v78
		_198_v78 = _dafny.SeqOf(_107_v1)
		var _199_v79 D1
		_ = _199_v79
		_199_v79 = Companion_D1_.Create_DC3_(_198_v78, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_192_v72)).Cardinality())))
		var _200_v80 _dafny.Map
		_ = _200_v80
		_200_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dqkshokld")).Cardinality()))
		var _201_v81 _dafny.Map
		_ = _201_v81
		_201_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v67, _109_v3)
		var _202_v82 _dafny.Map
		_ = _202_v82
		_202_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_187_v67, _107_v1)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_197_v77), 0))
		_ = _index36
		var _rhs29 _dafny.Int = (_199_v79).Dtor_cf5()
		_ = _rhs29
		var _rhs30 bool = _183_v63
		_ = _rhs30
		var _rhs31 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_189_v70).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_200_v80).Contains(_183_v63) {
				return (_200_v80).Get(_183_v63).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_106_v0).Cardinality())
		})(), _dafny.IntOfUint32((_189_v70).Cardinality()))).Uint32()).(bool), _183_v63), (Companion_Default___.SafeIndex((func() _dafny.Int {
			if _183_v63 {
				return _181_i2
			}
			return (_201_v81).Cardinality()
		})(), _dafny.IntOfUint32((_dafny.SeqOf((_189_v70).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_200_v80).Contains(_183_v63) {
				return (_200_v80).Get(_183_v63).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_106_v0).Cardinality())
		})(), _dafny.IntOfUint32((_189_v70).Cardinality()))).Uint32()).(bool), _183_v63)).Cardinality()))).Uint32(), !(!((_202_v82).Contains(_187_v67))))
		_ = _rhs31
		var _rhs32 bool = (_184_v64).Cmp(_dafny.IntOfInt64(282)) != 0
		_ = _rhs32
		var _lhs28 _dafny.Array = _197_v77
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_197_v77), 0))
		_ = _lhs29
		var _lhs30 *GlobalState = _105_globalState
		_ = _lhs30
		_185_v65 = _rhs29
		(_lhs28).ArraySet1(_rhs30, (_lhs29).Int())
		_189_v70 = _rhs31
		_lhs30.F0 = _rhs32
		var _203_v83 *C0
		_ = _203_v83
		var _nw32 *C0 = New_C0_()
		_ = _nw32
		_nw32.Ctor__(_186_v66)
		_203_v83 = _nw32
	}
	var _204_v84 _dafny.CodePoint
	_ = _204_v84
	_204_v84 = _dafny.CodePoint('m')
	var _205_v85 _dafny.Array
	_ = _205_v85
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(6)
	_ = _len0_5
	var _nw33 _dafny.Array
	_ = _nw33
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw33 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) bool = (func(_206_v3 bool) func(_dafny.Int) bool {
			return func(_207_i3 _dafny.Int) bool {
				return !(_206_v3)
			}
		})(_109_v3)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw33 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw33).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw33).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_205_v85 = _nw33
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))
	_ = _index37
	(_205_v85).ArraySet1(false, (_index37).Int())
	var _208_v86 _dafny.MultiSet
	_ = _208_v86
	_208_v86 = _dafny.MultiSetOf(_109_v3)
	var _209_v87 _dafny.Set
	_ = _209_v87
	_209_v87 = _dafny.SetOf((func() _dafny.Int {
		if (_208_v86).Contains(_109_v3) {
			return (_208_v86).Multiplicity(_109_v3)
		}
		return _107_v1
	})())
	var _210_v88 _dafny.Sequence
	_ = _210_v88
	_210_v88 = _dafny.SeqOf(_209_v87, _209_v87, _209_v87)
	var _211_v89 _dafny.Set
	_ = _211_v89
	_211_v89 = _dafny.SetOf(_109_v3)
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))
	_ = _index38
	var _rhs33 _dafny.Sequence = _106_v0
	_ = _rhs33
	var _rhs34 bool = !(_109_v3)
	_ = _rhs34
	var _rhs35 bool = (func() bool {
		if _109_v3 {
			return (_209_v87).IsSubsetOf((_210_v88).Select((Companion_Default___.SafeIndex(_107_v1, _dafny.IntOfUint32((_210_v88).Cardinality()))).Uint32()).(_dafny.Set))
		}
		return (_211_v89).IsProperSubsetOf(_211_v89)
	})()
	_ = _rhs35
	var _rhs36 _dafny.CodePoint = _204_v84
	_ = _rhs36
	var _rhs37 bool = !(!(_109_v3) || (_109_v3))
	_ = _rhs37
	var _lhs31 *GlobalState = _105_globalState
	_ = _lhs31
	var _lhs32 _dafny.Array = _205_v85
	_ = _lhs32
	var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))
	_ = _lhs33
	_106_v0 = _rhs33
	_109_v3 = _rhs34
	_lhs31.F0 = _rhs35
	_204_v84 = _rhs36
	(_lhs32).ArraySet1(_rhs37, (_lhs33).Int())
	var _212_v90 D1
	_ = _212_v90
	_212_v90 = Companion_D1_.Create_DC2_(_107_v1, true)
	var _213_v91 bool
	_ = _213_v91
	var _214_v92 _dafny.Int
	_ = _214_v92
	var _215_v93 _dafny.Int
	_ = _215_v93
	var _out12 bool
	_ = _out12
	var _out13 _dafny.Int
	_ = _out13
	var _out14 _dafny.Int
	_ = _out14
	_out12, _out13, _out14 = Companion_Default___.M0((Companion_Default___.Fm2(_107_v1, _105_globalState)).Update((_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool), Companion_Default___.Abs(_107_v1)), (_212_v90).Dtor_cf2(), _105_globalState)
	_213_v91 = _out12
	_214_v92 = _out13
	_215_v93 = _out14
	if _109_v3 {
		var _216_v94 _dafny.Map
		_ = _216_v94
		_216_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_214_v92, (_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool))
		var _217_v95 _dafny.Array
		_ = _217_v95
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_6
		var _nw34 _dafny.Array
		_ = _nw34
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw34 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = (func(_218_v93 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_219_i4 _dafny.Int) _dafny.Int {
					return (_219_i4).Plus(_218_v93)
				}
			})(_215_v93)
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
		_217_v95 = _nw34
		var _220_v96 _dafny.Map
		_ = _220_v96
		_220_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v95, _205_v85)
		var _221_v97 _dafny.Sequence
		_ = _221_v97
		_221_v97 = _dafny.SeqOf(_213_v91, !(_213_v91), (func() bool {
			if (_216_v94).Contains(_107_v1) {
				return (_216_v94).Get(_107_v1).(bool)
			}
			return (_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool)
		})(), (_220_v96).Contains(_217_v95), _213_v91)
		_213_v91 = !((_221_v97).Select((Companion_Default___.SafeIndex(_215_v93, _dafny.IntOfUint32((_221_v97).Cardinality()))).Uint32()).(bool))
		var _222_v98 T0
		_ = _222_v98
		var _nw35 *C0 = New_C0_()
		_ = _nw35
		_nw35.Ctor__(_204_v84)
		_222_v98 = _nw35
		_222_v98 = _222_v98
		var _223_v99 _dafny.Map
		_ = _223_v99
		_223_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(925))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_224_v98 T0) func(_dafny.Int) _dafny.CodePoint {
			return func(_225_i5 _dafny.Int) _dafny.CodePoint {
				return (_224_v98).F3()
			}
		})(_222_v98)))).Cardinality()), _204_v84)
		var _226_v100 _dafny.Map
		_ = _226_v100
		_226_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v3, (_223_v99).Cardinality())
		var _227_v101 _dafny.Map
		_ = _227_v101
		_227_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_221_v97).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(513), _dafny.IntOfUint32((_221_v97).Cardinality()))).Uint32()).(bool), _109_v3)
		var _228_v102 _dafny.Map
		_ = _228_v102
		_228_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v93, _226_v100)
		var _229_v103 _dafny.Sequence
		_ = _229_v103
		_229_v103 = _dafny.SeqOf(_215_v93, (_226_v100).Cardinality())
		var _230_v104 _dafny.Sequence
		_ = _230_v104
		_230_v104 = _dafny.SeqOf(_226_v100, _226_v100)
		var _rhs38 bool = !((func() bool {
			if (_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool) {
				return (_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool)
			}
			return !((func() bool {
				if (_227_v101).Contains((_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool)) {
					return (_227_v101).Get((_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool)).(bool)
				}
				return !((Companion_D1_.Create_DC2_(_107_v1, _213_v91)).Dtor_cf3())
			})())
		})())
		_ = _rhs38
		var _rhs39 _dafny.CodePoint = _204_v84
		_ = _rhs39
		var _rhs40 _dafny.Map = (((func() _dafny.Map {
			if (_228_v102).Contains((_229_v103).Select((Companion_Default___.SafeIndex(_107_v1, _dafny.IntOfUint32((_229_v103).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_228_v102).Get((_229_v103).Select((Companion_Default___.SafeIndex(_107_v1, _dafny.IntOfUint32((_229_v103).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Map)
			}
			return (_230_v104).Select((Companion_Default___.SafeIndex(_215_v93, _dafny.IntOfUint32((_230_v104).Cardinality()))).Uint32()).(_dafny.Map)
		})()).Update(_213_v91, _214_v92)).Merge(_226_v100)
		_ = _rhs40
		var _lhs34 *GlobalState = _105_globalState
		_ = _lhs34
		_lhs34.F0 = _rhs38
		_204_v84 = _rhs39
		_226_v100 = _rhs40
		_213_v91 = _109_v3
		_227_v101 = (_227_v101).Update(false, !((_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool)))
	} else {
		var _231_v105 bool
		_ = _231_v105
		var _232_v106 _dafny.Int
		_ = _232_v106
		var _233_v107 _dafny.Int
		_ = _233_v107
		var _out15 bool
		_ = _out15
		var _out16 _dafny.Int
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		_out15, _out16, _out17 = Companion_Default___.M0(_208_v86, _215_v93, _105_globalState)
		_231_v105 = _out15
		_232_v106 = _out16
		_233_v107 = _out17
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))
		_ = _index39
		(_205_v85).ArraySet1((_209_v87).IsDisjointFrom(func() _dafny.Set {
			var _coll7 = _dafny.NewBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(105), _dafny.IntOfInt64(36))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _234_v108 _dafny.Int
				_234_v108 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(105)).Cmp(_234_v108) <= 0) && ((_234_v108).Cmp(_dafny.IntOfInt64(36)) < 0) {
					_coll7.Add(Companion_Default___.SafeModuloInt(_234_v108, (func() _dafny.Int {
						if (_208_v86).Contains(_231_v105) {
							return (_208_v86).Multiplicity(_231_v105)
						}
						return _107_v1
					})()))
				}
			}
			return _coll7.ToSet()
		}()), (_index39).Int())
		var _235_v109 _dafny.Sequence
		_ = _235_v109
		_235_v109 = _dafny.SeqOf((_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool))
		_235_v109 = _235_v109
		var _236_v110 _dafny.Array
		_ = _236_v110
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_7
		var _nw36 _dafny.Array
		_ = _nw36
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw36 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.MultiSet = (func(_237_v2 _dafny.Map, _238_v107 _dafny.Int, _239_v105 bool, _240_v106 _dafny.Int, _241_v85 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
				return func(_242_i6 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if (_237_v2).Contains(_dafny.IntOfInt64(354)) {
							return (_237_v2).Get(_dafny.IntOfInt64(354)).(_dafny.Int)
						}
						return _238_v107
					})(), _239_v105), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v106, (Companion_D1_.Create_DC1_((_241_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_241_v85), 0))).Int()).(bool))).Dtor_cf1()))
				}
			})(_108_v2, _233_v107, _231_v105, _232_v106, _205_v85)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw36 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw36).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw36).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_236_v110 = _nw36
		var _243_v111 _dafny.Map
		_ = _243_v111
		_243_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_233_v107, true)
		var _244_v112 _dafny.MultiSet
		_ = _244_v112
		_244_v112 = _dafny.MultiSetOf(_243_v111, _243_v111)
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_236_v110), 0))
		_ = _index40
		(_236_v110).ArraySet1((_244_v112).Union(_244_v112), (_index40).Int())
		var _245_v113 _dafny.Array
		_ = _245_v113
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw37
		_245_v113 = _nw37
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
		_ = _index41
		(_245_v113).ArraySet1(_215_v93, (_index41).Int())
		var _246_v114 _dafny.Sequence
		_ = _246_v114
		_246_v114 = _dafny.SeqOf(_215_v93)
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_245_v113), 0))
		_ = _index42
		(_245_v113).ArraySet1((func() _dafny.Int {
			if _109_v3 {
				return _233_v107
			}
			return _dafny.IntOfUint32((_246_v114).Cardinality())
		})(), (_index42).Int())
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_236_v110), 0))
		_ = _index43
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
		_ = _index44
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_245_v113), 0))
		_ = _index45
		var _rhs41 _dafny.MultiSet = (_244_v112).Intersection(_dafny.MultiSetOf(func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_243_v111).Keys().Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _247_v115 _dafny.Int
				_247_v115 = interface{}(_compr_8).(_dafny.Int)
				if (_243_v111).Contains(_247_v115) {
					_coll8.Add((_247_v115).Times(_dafny.IntOfInt64(-496)), true)
				}
			}
			return _coll8.ToMap()
		}(), _243_v111))
		_ = _rhs41
		var _rhs42 bool = !((_208_v86).IsDisjointFrom((_208_v86).Update(_109_v3, Companion_Default___.Abs(_107_v1))))
		_ = _rhs42
		var _rhs43 _dafny.Sequence = (func() _dafny.Sequence {
			if (_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool) {
				return Companion_Default___.Fm8(_246_v114, _105_globalState)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("l")
		})()
		_ = _rhs43
		var _rhs44 _dafny.Int = Companion_Default___.SafeDivisionInt(_107_v1, Companion_Default___.SafeDivisionInt(_233_v107, (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm8(_246_v114, _105_globalState)).Cardinality()))))
		_ = _rhs44
		var _rhs45 _dafny.Int = (_214_v92).Times(_233_v107)
		_ = _rhs45
		var _lhs35 _dafny.Array = _236_v110
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_236_v110), 0))
		_ = _lhs36
		var _lhs37 *GlobalState = _105_globalState
		_ = _lhs37
		var _lhs38 _dafny.Array = _245_v113
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
		_ = _lhs39
		var _lhs40 _dafny.Array = _245_v113
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_245_v113), 0))
		_ = _lhs41
		(_lhs35).ArraySet1(_rhs41, (_lhs36).Int())
		_lhs37.F0 = _rhs42
		_106_v0 = _rhs43
		(_lhs38).ArraySet1(_rhs44, (_lhs39).Int())
		(_lhs40).ArraySet1(_rhs45, (_lhs41).Int())
		if !(!(_109_v3) || (_231_v105)) {
			var _248_v116 _dafny.Map
			_ = _248_v116
			_248_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v111, _208_v86)
			var _249_v117 D1
			_ = _249_v117
			_249_v117 = Companion_D1_.Create_DC1_((_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool))
			var _250_v118 bool
			_ = _250_v118
			var _251_v119 _dafny.Int
			_ = _251_v119
			var _252_v120 _dafny.Int
			_ = _252_v120
			var _out18 bool
			_ = _out18
			var _out19 _dafny.Int
			_ = _out19
			var _out20 _dafny.Int
			_ = _out20
			_out18, _out19, _out20 = Companion_Default___.M0((_208_v86).Intersection(((func() _dafny.MultiSet {
				if (_248_v116).Contains(_243_v111) {
					return (_248_v116).Get(_243_v111).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_231_v105)
			})()).Update((_249_v117).Dtor_cf1(), Companion_Default___.Abs((_243_v111).Cardinality()))), _233_v107, _105_globalState)
			_250_v118 = _out18
			_251_v119 = _out19
			_252_v120 = _out20
			var _253_v121 D1
			_ = _253_v121
			_253_v121 = Companion_D1_.Create_DC4_()
			_253_v121 = Companion_D1_.Create_DC4_()
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))
			_ = _index46
			(_205_v85).ArraySet1(((func() _dafny.Int {
				if _213_v91 {
					return (_212_v90).Dtor_cf2()
				}
				return _215_v93
			})()).Cmp((_dafny.Zero).Minus(((_dafny.SetOf((_208_v86).Cardinality(), _215_v93, _215_v93, _dafny.IntOfUint32((_235_v109).Cardinality()), (_245_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))).Int()).(_dafny.Int))).Cardinality()).Minus(_107_v1))) != 0, (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
			_ = _index47
			var _rhs46 _dafny.Int = _252_v120
			_ = _rhs46
			var _rhs47 _dafny.Int = Companion_Default___.SafeDivisionInt(_232_v106, (_214_v92).Plus(_215_v93))
			_ = _rhs47
			var _lhs42 _dafny.Array = _245_v113
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
			_ = _lhs43
			(_lhs42).ArraySet1(_rhs46, (_lhs43).Int())
			_252_v120 = _rhs47
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
			_ = _index48
			(_245_v113).ArraySet1(_251_v119, (_index48).Int())
		} else {
			var _254_v122 *C0
			_ = _254_v122
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__(_204_v84)
			_254_v122 = _nw38
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))
			_ = _index49
			(_205_v85).ArraySet1(!(_213_v91), (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
			_ = _index50
			(_245_v113).ArraySet1(_dafny.IntOfInt64(514), (_index50).Int())
			var _255_v123 *C0
			_ = _255_v123
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__(_204_v84)
			_255_v123 = _nw39
			var _256_v124 _dafny.Array
			_ = _256_v124
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_8
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.Sequence = (func(_257_v91 bool, _258_v0 _dafny.Sequence, _259_v84 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_260_i7 _dafny.Int) _dafny.Sequence {
						return (func() _dafny.Sequence {
							if _257_v91 {
								return _258_v0
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(495))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg15 _dafny.Int) interface{} {
									return coer15(arg15)
								}
							}((func(_261_v84 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_262_i8 _dafny.Int) _dafny.CodePoint {
									return _261_v84
								}
							})(_259_v84)))
						})()
					}
				})(_213_v91, _106_v0, _204_v84)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw40 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw40).ArraySet1(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw40).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_256_v124 = _nw40
			var _263_v125 D3
			_ = _263_v125
			_263_v125 = Companion_D3_.Create_DC8_(_256_v124)
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
			_ = _index51
			var _rhs48 _dafny.Int = (_dafny.IntOfInt64(633)).Times(((_246_v114).Select((Companion_Default___.SafeIndex(_215_v93, _dafny.IntOfUint32((_246_v114).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_dafny.IntOfInt64(-719)))
			_ = _rhs48
			var _rhs49 _dafny.Array = (_263_v125).Dtor_cf9()
			_ = _rhs49
			var _rhs50 _dafny.Int = _215_v93
			_ = _rhs50
			var _rhs51 bool = ((func() _dafny.Int {
				if (_208_v86).Contains(_213_v91) {
					return (_208_v86).Multiplicity(_213_v91)
				}
				return Companion_Default___.Fm0(_105_globalState)
			})()).Cmp(_233_v107) == 0
			_ = _rhs51
			var _lhs44 _dafny.Array = _245_v113
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_245_v113), 0))
			_ = _lhs45
			var _lhs46 *GlobalState = _105_globalState
			_ = _lhs46
			(_lhs44).ArraySet1(_rhs48, (_lhs45).Int())
			_256_v124 = _rhs49
			_215_v93 = _rhs50
			_lhs46.F0 = _rhs51
		}
	}
	var _264_v126 _dafny.Sequence
	_ = _264_v126
	_264_v126 = _dafny.SeqOf(_214_v92)
	var _265_v127 *C0
	_ = _265_v127
	var _nw41 *C0 = New_C0_()
	_ = _nw41
	_nw41.Ctor__(_204_v84)
	_265_v127 = _nw41
	var _266_v128 _dafny.Map
	_ = _266_v128
	_266_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_214_v92, _265_v127)
	var _267_v129 _dafny.Sequence
	_ = _267_v129
	_267_v129 = _dafny.SeqOf((func() *C0 {
		if (_266_v128).Contains(_215_v93) {
			return (_266_v128).Get(_215_v93).(*C0)
		}
		return _265_v127
	})(), _265_v127)
	_107_v1 = (_107_v1).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_264_v126, _dafny.SeqOf(_dafny.IntOfInt64(638), _214_v92, _dafny.IntOfUint32((_267_v129).Cardinality())))).Cardinality()))
	_107_v1 = _215_v93
	_109_v3 = !(_213_v91)
	var _268_i9 _dafny.Int
	_ = _268_i9
	_268_i9 = _dafny.Zero
	{
		for _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_264_v126, (Companion_Default___.SafeIndex(_215_v93, _dafny.IntOfUint32((_264_v126).Cardinality()))).Uint32(), _214_v92), _dafny.Companion_Sequence_.Concatenate(_264_v126, _264_v126)) {
			{
				if (_268_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_268_i9 = (_268_i9).Plus(_dafny.One)
				var _269_v130 _dafny.Sequence
				_ = _269_v130
				_269_v130 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v93, _213_v91), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_264_v126).Cardinality()), (Companion_D1_.Create_DC2_(_dafny.IntOfInt64(58), (_205_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_205_v85), 0))).Int()).(bool))).Dtor_cf3()))
				_269_v130 = _269_v130
				_215_v93 = (_264_v126).Select((Companion_Default___.SafeIndex(_107_v1, _dafny.IntOfUint32((_264_v126).Cardinality()))).Uint32()).(_dafny.Int)
				_214_v92 = _107_v1
				var _270_v131 T0
				_ = _270_v131
				var _nw42 *C0 = New_C0_()
				_ = _nw42
				_nw42.Ctor__(_dafny.CodePoint('g'))
				_270_v131 = _nw42
				_270_v131 = _270_v131
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	(_105_globalState).F0 = (_109_v3) && (_213_v91)
	_213_v91 = _109_v3
	_dafny.Print(_105_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_106_v0, _dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'), _dafny.CodePoint('b'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_107_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_108_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-27), _dafny.IntOfInt64(-27))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_109_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_204_v84)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v85).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v85).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v85).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v85).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v85).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_205_v85).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_208_v86).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_209_v87).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_210_v88, _dafny.SeqOf(_dafny.SetOf(_dafny.One), _dafny.SetOf(_dafny.One), _dafny.SetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_211_v89).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_212_v90).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_213_v91)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_214_v92)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_215_v93)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_264_v126, _dafny.SeqOf(_dafny.IntOfInt64(455))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_266_v128).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_267_v129).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_268_i9)
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

type D1_DC2 struct {
	Cf2 _dafny.Int
	Cf3 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 bool) D1 {
	return D1{D1_DC2{Cf2, Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC3 struct {
	Cf4 _dafny.Sequence
	Cf5 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.Sequence, Cf5 _dafny.Int) D1 {
	return D1{D1_DC3{Cf4, Cf5}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_() D1 {
	return D1{D1_DC4{}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC1 struct {
	Cf1 bool
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 bool) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, false)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() bool {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf1() bool {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4"
		}
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
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3 == data2.Cf3
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf4.Equals(data2.Cf4) && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC4:
		{
			_, ok := other.Get_().(D1_DC4)
			return ok
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1 == data2.Cf1
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

type D2_DC6 struct {
	Cf7 T0
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf7 T0) D2 {
	return D2{D2_DC6{Cf7}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf6 _dafny.MultiSet
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 _dafny.MultiSet) D2 {
	return D2{D2_DC5{Cf6}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC7 struct {
	Cf8 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf8 D2) D2 {
	return D2{D2_DC7{Cf8}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_((T0)(nil))
}

func (_this D2) Dtor_cf7() T0 {
	return _this.Get_().(D2_DC6).Cf7
}

func (_this D2) Dtor_cf6() _dafny.MultiSet {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf8() D2 {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf7) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf8) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && _dafny.AreEqual(data1.Cf7, data2.Cf7)
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf10 _dafny.CodePoint
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf10 _dafny.CodePoint) D3 {
	return D3{D3_DC9{Cf10}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf9 _dafny.Array
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf9 _dafny.Array) D3 {
	return D3{D3_DC8{Cf9}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.CodePoint('D'))
}

func (_this D3) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D3_DC8).Cf9
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf10) + ")"
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
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf10 == data2.Cf10
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf9 == data2.Cf9
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
	Cf12 D1
	Cf13 _dafny.MultiSet
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf12 D1, Cf13 _dafny.MultiSet) D4 {
	return D4{D4_DC11{Cf12, Cf13}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf11 _dafny.Sequence
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf11 _dafny.Sequence) D4 {
	return D4{D4_DC10{Cf11}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(Companion_D1_.Default(), _dafny.EmptyMultiSet)
}

func (_this D4) Dtor_cf12() D1 {
	return _this.Get_().(D4_DC11).Cf12
}

func (_this D4) Dtor_cf13() _dafny.MultiSet {
	return _this.Get_().(D4_DC11).Cf13
}

func (_this D4) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D4_DC10).Cf11
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D4_DC10:
		{
			return "D4.DC10" + "(" + data.Cf11.VerbatimString(true) + ")"
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
			return ok && data1.Cf12.Equals(data2.Cf12) && data1.Cf13.Equals(data2.Cf13)
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf15 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf15 bool) D5 {
	return D5{D5_DC13{Cf15}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC12 struct {
	Cf14 _dafny.Map
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf14 _dafny.Map) D5 {
	return D5{D5_DC12{Cf14}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC14 struct {
	Cf16 D5
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf16 D5) D5 {
	return D5{D5_DC14{Cf16}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_(false)
}

func (_this D5) Dtor_cf15() bool {
	return _this.Get_().(D5_DC13).Cf15
}

func (_this D5) Dtor_cf14() _dafny.Map {
	return _this.Get_().(D5_DC12).Cf14
}

func (_this D5) Dtor_cf16() D5 {
	return _this.Get_().(D5_DC14).Cf16
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf15 == data2.Cf15
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf14.Equals(data2.Cf14)
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm3(globalState *GlobalState) bool
	F3() _dafny.CodePoint
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
	F0  bool
	_f1 bool
	_f2 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = false
	_this._f1 = false
	_this._f2 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 bool, f2 bool) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f3 _dafny.CodePoint
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f3 = _dafny.CodePoint('D')
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

func (_this *C0) F3() _dafny.CodePoint {
	return _this._f3
}
func (_this *C0) Ctor__(f3 _dafny.CodePoint) {
	{
		(_this)._f3 = f3
	}
}
func (_this *C0) Fm3(globalState *GlobalState) bool {
	{
		var _source2 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tqw")).Cardinality())
		_ = _source2
		var _271___mcc_h0 _dafny.Int = _source2
		_ = _271___mcc_h0
		var _272_cf0 _dafny.Int = _271___mcc_h0
		_ = _272_cf0
		return !(false)
	}
}
func (_this *C0) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return (Companion_D1_.Create_DC1_(true)).Dtor_cf1()
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
