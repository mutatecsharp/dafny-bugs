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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, p1 bool, globalState *GlobalState) bool {
	if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(true, false, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(252))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i1 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf(false, false)).Cardinality()
	}))).Cardinality())))).IsSubsetOf(_dafny.MultiSetOf(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(294), _dafny.IntOfInt64(982))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _1_v0 _dafny.Int
			_1_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(294)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(982)) < 0) {
				_coll0.Add((_1_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg1 _dafny.Int) interface{} {
						return coer1(arg1)
					}
				}(func(_2_i0 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))).Cardinality()
				})))).Cardinality())).Cardinality()), (_dafny.SetOf(false)).Cardinality())
			}
		}
		return _coll0.ToMap()
	}())) {
		return true
	} else {
		return false
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Cardinality(), ((_dafny.MultiSetOf(_dafny.IntOfInt64(-146))).Cardinality()).Times(_dafny.IntOfInt64(-601)))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("t")
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true, true, (_dafny.MultiSetOf(_dafny.IntOfInt64(693), _dafny.IntOfInt64(689))).IsProperSubsetOf(_dafny.MultiSetOf((_dafny.MultiSetOf(false, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.CodePoint, p1 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('p')
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!((!(true)) == (false)), !(!(true)), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_dafny.SetOf((Companion_D8_.Create_DC20_(false, _dafny.IntOfInt64(962))).Dtor_cf26(), _dafny.IntOfInt64(-800))), _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(-845)))), false, (_dafny.IntOfInt64(782)).Cmp(_dafny.IntOfInt64(905)) == 0)
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(522))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_D3_.Create_DC9_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lihiauqtn")).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(654))).Cardinality()), _dafny.IntOfInt64(-799))).Cardinality()))).Dtor_cf13()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(631), false)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(690))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D3_.Create_DC9_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(70), _dafny.IntOfInt64(280)))).Dtor_cf13(), _dafny.IntOfInt64(299))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-53))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality()))
	}))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.SetOf(true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-674))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})), !(true))).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 _dafny.Sequence
			_5_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-674))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			})), !(true))).Contains(_5_v0) {
				_coll1.Add(_5_v0)
			}
		}
		return _coll1.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-122), Companion_D3_.Create_DC9_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(_dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('e'))).Cardinality())).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-32)))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(333))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hqgfpgnr")).Cardinality())
	}))
}
func (_static *CompanionStruct_Default___) Fm12(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(68), _dafny.IntOfInt64(47))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(79))).Cardinality())))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-498)), (Companion_D0_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Dtor_cf3())).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(310), _dafny.IntOfInt64(-947))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _7_v0 _dafny.Int
			_7_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(310)).Cmp(_7_v0) <= 0) && ((_7_v0).Cmp(_dafny.IntOfInt64(-947)) < 0) {
				_coll2.Add((_7_v0).Times((_dafny.MultiSetOf((_dafny.MultiSetOf(true, false)).Cardinality())).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cjm")).Cardinality())))
			}
		}
		return _coll2.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC4_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(Companion_D3_.Create_DC9_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqbk")).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, false, true, true, true)).Cardinality()))))).Cardinality()))).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-7), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dnpmcdw")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('y'))).Cardinality()).Minus(_dafny.IntOfInt64(96)), _dafny.CodePoint('o'))
}
func (_static *CompanionStruct_Default___) M0(p0 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, _dafny.Array, _dafny.CodePoint) {
	var r0 _dafny.Sequence = _dafny.EmptySeq
	_ = r0
	var r1 _dafny.Int = _dafny.Zero
	_ = r1
	var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r2
	var r3 _dafny.CodePoint = _dafny.CodePoint('D')
	_ = r3
	var _8_i0 _dafny.Int
	_ = _8_i0
	_8_i0 = _dafny.Zero
	{
		for p0 {
			{
				if (_8_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_8_i0 = (_8_i0).Plus(_dafny.One)
				var _9_v0 *C0
				_ = _9_v0
				var _nw0 *C0 = New_C0_()
				_ = _nw0
				_nw0.Ctor__()
				_9_v0 = _nw0
				(globalState).F3 = !(p0)
				var _10_v1 _dafny.Sequence
				_ = _10_v1
				_10_v1 = _dafny.SeqOf(p0)
				(globalState).F3 = (func() bool {
					if p0 {
						return p0
					}
					return (_dafny.MultiSetFromSeq(_10_v1)).IsProperSubsetOf(_dafny.MultiSetOf(p0))
				})()
				var _11_v2 _dafny.Set
				_ = _11_v2
				_11_v2 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jbq"), _dafny.UnicodeSeqOfUtf8Bytes("jjrtckkdd")))
				var _12_v3 _dafny.Int
				_ = _12_v3
				_12_v3 = _dafny.IntOfInt64(346)
				_11_v2 = Companion_Default___.Fm9(p0, _12_v3, globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _13_v4 _dafny.Int
	_ = _13_v4
	_13_v4 = _dafny.IntOfInt64(346)
	var _14_v5 D0
	_ = _14_v5
	_14_v5 = Companion_D0_.Create_DC2_(_13_v4, _13_v4)
	var _15_v6 _dafny.Sequence
	_ = _15_v6
	_15_v6 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dfufllfmk")).Cardinality()), _13_v4)
	(globalState).F22 = ((_14_v5).Dtor_cf2()).Minus(_dafny.IntOfUint32((_15_v6).Cardinality()))
	var _16_v7 D2
	_ = _16_v7
	_16_v7 = Companion_D2_.Create_DC6_()
	_16_v7 = _16_v7
	var _source0 D8 = Companion_D8_.Create_DC20_(p0, _13_v4)
	_ = _source0
	if _source0.Is_DC20() {
		var _17___mcc_h0 bool = _source0.Get_().(D8_DC20).Cf25
		_ = _17___mcc_h0
		var _18___mcc_h1 _dafny.Int = _source0.Get_().(D8_DC20).Cf26
		_ = _18___mcc_h1
		var _19_cf26 _dafny.Int = _18___mcc_h1
		_ = _19_cf26
		var _20_cf25 bool = _17___mcc_h0
		_ = _20_cf25
		var _21_v8 _dafny.Array
		_ = _21_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_0
		var _nw1 _dafny.Array
		_ = _nw1
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw1 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_22_p0 bool) func(_dafny.Int) bool {
				return func(_23_i1 _dafny.Int) bool {
					return _22_p0
				}
			})(p0)
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
		_21_v8 = _nw1
		var _24_v9 _dafny.Map
		_ = _24_v9
		_24_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v4, _13_v4)
		var _25_v10 _dafny.Set
		_ = _25_v10
		_25_v10 = _dafny.SetOf(_24_v9, (_24_v9).Update(_dafny.IntOfUint32((_dafny.SeqOf(_19_cf26)).Cardinality()), _19_cf26), _24_v9, _24_v9, _24_v9)
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_21_v8), 0))
		_ = _index0
		(_21_v8).ArraySet1((_25_v10).Equals(_25_v10), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_21_v8), 0))
		_ = _index1
		(_21_v8).ArraySet1(p0, (_index1).Int())
		var _26_v11 _dafny.CodePoint
		_ = _26_v11
		_26_v11 = _dafny.CodePoint('l')
		var _27_v12 _dafny.MultiSet
		_ = _27_v12
		_27_v12 = _dafny.MultiSetOf(_13_v4, _13_v4, _19_cf26, _19_cf26)
		var _28_v13 _dafny.Map
		_ = _28_v13
		_28_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_26_v11, (_19_cf26).Times((_27_v12).Cardinality()))
		_28_v13 = (_28_v13).Update(_26_v11, (func() _dafny.Int {
			if p0 {
				return _13_v4
			}
			return _13_v4
		})())
		var _29_v14 _dafny.Map
		_ = _29_v14
		_29_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_13_v4).Cmp(_dafny.IntOfInt64(584)) != 0, Companion_Default___.SafeModuloInt(_19_cf26, _13_v4))
		_29_v14 = (_29_v14).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_21_v8), 0))).Int()).(bool), _19_cf26))
		var _30_v15 _dafny.Map
		_ = _30_v15
		_30_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_21_v8), 0))).Int()).(bool))
		_30_v15 = (_30_v15).Update((_21_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(554), _dafny.ArrayLen((_21_v8), 0))).Int()).(bool), p0)
	} else if _source0.Is_DC19() {
		var _31___mcc_h2 _dafny.Map = _source0.Get_().(D8_DC19).Cf24
		_ = _31___mcc_h2
		var _32_cf24 _dafny.Map = _31___mcc_h2
		_ = _32_cf24
		(globalState).F14 = (func() bool {
			if false {
				return p0
			}
			return (func() bool {
				if p0 {
					return p0
				}
				return p0
			})()
		})()
		var _33_v16 _dafny.CodePoint
		_ = _33_v16
		_33_v16 = _dafny.CodePoint('a')
		var _34_v17 _dafny.Map
		_ = _34_v17
		_34_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v16, _dafny.CodePoint('k'))
		var _35_v18 _dafny.Set
		_ = _35_v18
		_35_v18 = _dafny.SetOf(p0)
		if (_13_v4).Cmp(Companion_Default___.SafeModuloInt((_34_v17).Cardinality(), (_35_v18).Cardinality())) <= 0 {
			var _36_v19 *C0
			_ = _36_v19
			var _nw2 *C0 = New_C0_()
			_ = _nw2
			_nw2.Ctor__()
			_36_v19 = _nw2
			var _37_v20 _dafny.Sequence
			_ = _37_v20
			_37_v20 = _dafny.UnicodeSeqOfUtf8Bytes("ntnbowsn")
			(globalState).F21 = _37_v20
			var _38_v21 *C0
			_ = _38_v21
			var _nw3 *C0 = New_C0_()
			_ = _nw3
			_nw3.Ctor__()
			_38_v21 = _nw3
			_33_v16 = _33_v16
			(globalState).F22 = _13_v4
		} else {
			(globalState).F14 = p0
			var _39_v22 _dafny.MultiSet
			_ = _39_v22
			_39_v22 = _dafny.MultiSetOf(p0)
			_39_v22 = _39_v22
			(globalState).F25 = (_13_v4).Times(_dafny.IntOfInt64(-892))
			var _40_v23 _dafny.Array
			_ = _40_v23
			var _nw4 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
			_ = _nw4
			_40_v23 = _nw4
			var _41_v24 *C0
			_ = _41_v24
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__()
			_41_v24 = _nw5
			var _42_v25 _dafny.Map
			_ = _42_v25
			_42_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_41_v24, _40_v23)
			var _43_v26 _dafny.Sequence
			_ = _43_v26
			_43_v26 = _dafny.SeqOf(_40_v23, (func() _dafny.Array {
				if (_42_v25).Contains(_41_v24) {
					return (_42_v25).Get(_41_v24).(_dafny.Array)
				}
				return _40_v23
			})())
			(globalState).F3 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_43_v26, _dafny.SeqOf(_40_v23))).Cardinality())).Cmp(_13_v4) <= 0
			var _44_v27 _dafny.Map
			_ = _44_v27
			_44_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _45_v28 D0
			_ = _45_v28
			_45_v28 = Companion_D0_.Create_DC1_(p0)
			var _46_v29 _dafny.Sequence
			_ = _46_v29
			_46_v29 = _dafny.SeqOf(_45_v28)
			var _47_v30 _dafny.Map
			_ = _47_v30
			_47_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_44_v27, _46_v29)
			var _48_v31 _dafny.Map
			_ = _48_v31
			_48_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D5_.Create_DC14_(_47_v30, _13_v4)).Dtor_cf20(), _dafny.IntOfInt64(309))
			var _49_v32 D3
			_ = _49_v32
			_49_v32 = Companion_D3_.Create_DC9_(p0, _48_v31)
			var _50_v33 _dafny.Map
			_ = _50_v33
			_50_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_51_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_52_i2 _dafny.Int) _dafny.CodePoint {
					return _51_v16
				}
			})(_33_v16)))).Cardinality()), _49_v32)
			var _53_v34 _dafny.Array
			_ = _53_v34
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw6
			_53_v34 = _nw6
			var _54_v35 D4
			_ = _54_v35
			_54_v35 = Companion_D4_.Create_DC11_(_50_v33, _53_v34)
			_54_v35 = _54_v35
		}
		var _55_v36 _dafny.Sequence
		_ = _55_v36
		_55_v36 = _dafny.UnicodeSeqOfUtf8Bytes("llmnukb")
		var _56_v37 _dafny.Map
		_ = _56_v37
		_56_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_v16, _55_v36)
		var _source1 D0 = Companion_D0_.Create_DC1_(Companion_Default___.Fm0(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_56_v37).Contains(_33_v16) {
				return (_56_v37).Get(_33_v16).(_dafny.Sequence)
			}
			return _55_v36
		})()).Cardinality()), p0, globalState))
		_ = _source1
		if _source1.Is_DC1() {
			var _57___mcc_h4 bool = _source1.Get_().(D0_DC1).Cf1
			_ = _57___mcc_h4
			var _58_cf1 bool = _57___mcc_h4
			_ = _58_cf1
			(globalState).F24 = true
			(globalState).F9 = ((func() _dafny.Int {
				if _58_cf1 {
					return _13_v4
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iwdi")).Cardinality())
			})()).Times(_13_v4)
			var _59_v38 _dafny.Sequence
			_ = _59_v38
			_59_v38 = _dafny.SeqOf(_58_cf1)
			var _60_v39 _dafny.Sequence
			_ = _60_v39
			_60_v39 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_59_v38, _dafny.SeqOf(p0, p0)))
			r0 = (_60_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bkf")).Cardinality()), _dafny.IntOfUint32((_60_v39).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _61_v40 _dafny.Map
			_ = _61_v40
			_61_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_13_v4, _59_v38)).Cardinality(), _dafny.IntOfInt64(177))
			var _62_v41 D3
			_ = _62_v41
			_62_v41 = Companion_D3_.Create_DC9_(Companion_Default___.Fm0(_13_v4, _58_cf1, globalState), (_61_v40).Merge(_61_v40))
			_62_v41 = _62_v41
		} else if _source1.Is_DC2() {
			var _63___mcc_h5 _dafny.Int = _source1.Get_().(D0_DC2).Cf2
			_ = _63___mcc_h5
			var _64___mcc_h6 _dafny.Int = _source1.Get_().(D0_DC2).Cf3
			_ = _64___mcc_h6
			var _65_cf3 _dafny.Int = _64___mcc_h6
			_ = _65_cf3
			var _66_cf2 _dafny.Int = _63___mcc_h5
			_ = _66_cf2
			var _67_v42 *C0
			_ = _67_v42
			var _nw7 *C0 = New_C0_()
			_ = _nw7
			_nw7.Ctor__()
			_67_v42 = _nw7
			var _68_v43 _dafny.Set
			_ = _68_v43
			_68_v43 = _dafny.SetOf(_33_v16, _33_v16, _33_v16)
			var _69_v44 _dafny.Array
			_ = _69_v44
			var _nwElement0_0 _dafny.Set = _68_v43
			_ = _nwElement0_0
			var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(7))
			_ = _nw8
			(_nw8).ArraySet1(_nwElement0_0, 0)
			(_nw8).ArraySet1(_68_v43, 1)
			(_nw8).ArraySet1(_dafny.SetOf(_33_v16, _33_v16, _33_v16, _33_v16), 2)
			(_nw8).ArraySet1(_68_v43, 3)
			(_nw8).ArraySet1((func() _dafny.Set {
				if !(p0) {
					return _68_v43
				}
				return _68_v43
			})(), 4)
			(_nw8).ArraySet1(_68_v43, 5)
			(_nw8).ArraySet1(_68_v43, 6)
			_69_v44 = _nw8
			var _70_v45 _dafny.Array
			_ = _70_v45
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_1
			var _nw9 _dafny.Array
			_ = _nw9
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw9 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_71_v36 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_72_i3 _dafny.Int) _dafny.Int {
						return (_72_i3).Times(_dafny.IntOfUint32((_71_v36).Cardinality()))
					}
				})(_55_v36)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw9 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw9).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw9).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_70_v45 = _nw9
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_70_v45), 0))
			_ = _index2
			(_70_v45).ArraySet1(_dafny.IntOfInt64(551), (_index2).Int())
			var _73_v46 _dafny.Array
			_ = _73_v46
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_2
			var _nw10 _dafny.Array
			_ = _nw10
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw10 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = (func(_74_p0 bool) func(_dafny.Int) bool {
					return func(_75_i4 _dafny.Int) bool {
						return _74_p0
					}
				})(p0)
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
			_73_v46 = _nw10
			var _76_v47 D4
			_ = _76_v47
			_76_v47 = Companion_D4_.Create_DC11_(Companion_Default___.Fm10(_33_v16, _dafny.IntOfUint32((_15_v6).Cardinality()), globalState), _73_v46)
			var _77_v48 _dafny.Map
			_ = _77_v48
			_77_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v47, p0)
			var _78_v49 _dafny.Array
			_ = _78_v49
			var _nwElement0_1 _dafny.Map = _77_v48
			_ = _nwElement0_1
			var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(9))
			_ = _nw11
			(_nw11).ArraySet1(_nwElement0_1, 0)
			(_nw11).ArraySet1(_77_v48, 1)
			(_nw11).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v47, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v47, p0)), 2)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v47, p0), 3)
			(_nw11).ArraySet1((_77_v48).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v47, p0)).Update(_76_v47, p0)), 4)
			(_nw11).ArraySet1(_77_v48, 5)
			(_nw11).ArraySet1(_77_v48, 6)
			(_nw11).ArraySet1((_77_v48).Merge(_77_v48), 7)
			(_nw11).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v47, p0), 8)
			_78_v49 = _nw11
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_78_v49), 0))
			_ = _index3
			(_78_v49).ArraySet1(_77_v48, (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_70_v45), 0))
			_ = _index4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_78_v49), 0))
			_ = _index5
			var _rhs0 _dafny.Array = _69_v44
			_ = _rhs0
			var _rhs1 _dafny.Int = _65_cf3
			_ = _rhs1
			var _rhs2 bool = !(!((func() bool {
				if p0 {
					return p0
				}
				return p0
			})()))
			_ = _rhs2
			var _rhs3 _dafny.Map = (_77_v48).Merge((_77_v48).Merge(_77_v48))
			_ = _rhs3
			var _rhs4 bool = p0
			_ = _rhs4
			var _lhs0 _dafny.Array = _70_v45
			_ = _lhs0
			var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_70_v45), 0))
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			var _lhs3 _dafny.Array = _78_v49
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_78_v49), 0))
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_69_v44 = _rhs0
			(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
			_lhs2.F24 = _rhs2
			(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
			_lhs5.F3 = _rhs4
			var _79_v50 _dafny.Set
			_ = _79_v50
			_79_v50 = _dafny.SetOf(_73_v46, _73_v46)
			var _rhs5 _dafny.Set = (_79_v50).Intersection((func() _dafny.Set {
				if p0 {
					return _79_v50
				}
				return _dafny.SetOf(_73_v46)
			})())
			_ = _rhs5
			var _rhs6 _dafny.CodePoint = _33_v16
			_ = _rhs6
			var _rhs7 _dafny.Int = (_70_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_70_v45), 0))).Int()).(_dafny.Int)
			_ = _rhs7
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			_79_v50 = _rhs5
			r3 = _rhs6
			_lhs6.F9 = _rhs7
			var _80_v51 _dafny.Map
			_ = _80_v51
			_80_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _66_cf2)
			var _81_v52 _dafny.Map
			_ = _81_v52
			_81_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v51, (func() bool {
				if p0 {
					return !(p0)
				}
				return !(p0)
			})())
			_81_v52 = (_81_v52).Update(_80_v51, p0)
		} else {
			var _82___mcc_h7 _dafny.Array = _source1.Get_().(D0_DC0).Cf0
			_ = _82___mcc_h7
			var _83_cf0 _dafny.Array = _82___mcc_h7
			_ = _83_cf0
			var _84_v53 *C0
			_ = _84_v53
			var _nw12 *C0 = New_C0_()
			_ = _nw12
			_nw12.Ctor__()
			_84_v53 = _nw12
			(globalState).F21 = _dafny.UnicodeSeqOfUtf8Bytes("ukla")
			var _85_v54 _dafny.Sequence
			_ = _85_v54
			_85_v54 = _dafny.SeqOf(Companion_Default___.Fm11(p0, globalState), Companion_Default___.Fm11(p0, globalState), _15_v6)
			r1 = (_dafny.IntOfUint32((Companion_Default___.Fm2((_84_v53).Fm5(globalState), _dafny.IntOfUint32((_15_v6).Cardinality()), p0, _13_v4, globalState)).Cardinality())).Times(_dafny.IntOfUint32((_85_v54).Cardinality()))
			(globalState).F9 = _13_v4
		}
		var _86_v55 *C0
		_ = _86_v55
		var _nw13 *C0 = New_C0_()
		_ = _nw13
		_nw13.Ctor__()
		_86_v55 = _nw13
		var _87_v56 _dafny.Array
		_ = _87_v56
		var _nwElement0_2 *C0 = _86_v55
		_ = _nwElement0_2
		var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.One)
		_ = _nw14
		(_nw14).ArraySet1(_nwElement0_2, 0)
		_87_v56 = _nw14
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_87_v56), 0))
		_ = _index6
		(_87_v56).ArraySet1(_86_v55, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(359), _dafny.ArrayLen((_87_v56), 0))
		_ = _index7
		(_87_v56).ArraySet1(_86_v55, (_index7).Int())
	} else {
		var _88___mcc_h3 D8 = _source0.Get_().(D8_DC21).Cf27
		_ = _88___mcc_h3
		var _89_cf27 D8 = _88___mcc_h3
		_ = _89_cf27
		var _90_v57 *C0
		_ = _90_v57
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__()
		_90_v57 = _nw15
		r1 = _dafny.IntOfInt64(788)
		var _91_v58 _dafny.Sequence
		_ = _91_v58
		_91_v58 = _dafny.SeqOf(!(p0))
		var _92_v59 _dafny.Map
		_ = _92_v59
		_92_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v58, p0)
		_92_v59 = _92_v59
		var _93_v60 _dafny.Array
		_ = _93_v60
		var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
		_ = _nw16
		_93_v60 = _nw16
		var _94_v61 _dafny.CodePoint
		_ = _94_v61
		_94_v61 = _dafny.CodePoint('x')
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_93_v60), 0))
		_ = _index8
		(_93_v60).ArraySet1CodePoint((func() _dafny.CodePoint {
			if p0 {
				return _94_v61
			}
			return _94_v61
		})(), (_index8).Int())
		var _95_v62 _dafny.Map
		_ = _95_v62
		_95_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _96_v63 D0
		_ = _96_v63
		_96_v63 = Companion_D0_.Create_DC1_(p0)
		var _97_v64 _dafny.Sequence
		_ = _97_v64
		_97_v64 = _dafny.SeqOf(_96_v63, _96_v63)
		var _98_v65 _dafny.Map
		_ = _98_v65
		_98_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v62, _97_v64)
		var _99_v66 D5
		_ = _99_v66
		_99_v66 = Companion_D5_.Create_DC14_((_98_v65).Update(_95_v62, _97_v64), _13_v4)
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_93_v60), 0))
		_ = _index9
		var _rhs8 _dafny.Int = (_99_v66).Dtor_cf20()
		_ = _rhs8
		var _rhs9 bool = !(!(p0))
		_ = _rhs9
		var _rhs10 _dafny.CodePoint = _dafny.CodePoint('c')
		_ = _rhs10
		var _rhs11 _dafny.Int = ((_dafny.Zero).Minus(_13_v4)).Plus(_13_v4)
		_ = _rhs11
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		var _lhs8 _dafny.Array = _93_v60
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_93_v60), 0))
		_ = _lhs9
		var _lhs10 *GlobalState = globalState
		_ = _lhs10
		r1 = _rhs8
		_lhs7.F24 = _rhs9
		(_lhs8).ArraySet1CodePoint(_rhs10, (_lhs9).Int())
		_lhs10.F25 = _rhs11
	}
	var _100_v67 *C0
	_ = _100_v67
	var _nw17 *C0 = New_C0_()
	_ = _nw17
	_nw17.Ctor__()
	_100_v67 = _nw17
	var _101_v68 _dafny.Map
	_ = _101_v68
	_101_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v67, (_100_v67).Fm5(globalState))
	var _102_v69 _dafny.Map
	_ = _102_v69
	_102_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_101_v68).Contains(_100_v67) {
			return (_101_v68).Get(_100_v67).(_dafny.Int)
		}
		return _dafny.IntOfInt64(86)
	})(), _13_v4)
	var _rhs12 _dafny.Int = _13_v4
	_ = _rhs12
	var _rhs13 _dafny.Map = (_102_v69).Merge((_102_v69).Merge(_102_v69))
	_ = _rhs13
	_13_v4 = _rhs12
	_102_v69 = _rhs13
	var _hi0 _dafny.Int = _13_v4
	_ = _hi0
	for _103_i5 := _13_v4; _103_i5.Cmp(_hi0) < 0; _103_i5 = _103_i5.Plus(_dafny.One) {
		var _104_v70 _dafny.Array
		_ = _104_v70
		var _nwElement0_3 bool = p0
		_ = _nwElement0_3
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(2))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_3, 0)
		(_nw18).ArraySet1(Companion_Default___.Fm0(_13_v4, true, globalState), 1)
		_104_v70 = _nw18
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_104_v70), 0))
		_ = _index10
		(_104_v70).ArraySet1(false, (_index10).Int())
		var _105_v71 _dafny.MultiSet
		_ = _105_v71
		_105_v71 = _dafny.MultiSetOf(p0)
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_104_v70), 0))
		_ = _index11
		(_104_v70).ArraySet1(Companion_Default___.Fm0(_103_i5, (_dafny.IntOfInt64(417)).Cmp((func() _dafny.Int {
			if (_105_v71).Contains(p0) {
				return (_105_v71).Multiplicity(p0)
			}
			return _103_i5
		})()) == 0, globalState), (_index11).Int())
		(globalState).F14 = (_104_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_104_v70), 0))).Int()).(bool)
		var _106_v72 _dafny.Sequence
		_ = _106_v72
		_106_v72 = _dafny.SeqOf(p0, !(p0), p0, p0)
		var _107_v73 _dafny.Sequence
		_ = _107_v73
		_107_v73 = _dafny.UnicodeSeqOfUtf8Bytes("fjmbsow")
		(globalState).F9 = (((_14_v5).Dtor_cf2()).Minus(_dafny.IntOfUint32((_106_v72).Cardinality()))).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_108_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), _107_v73)).Cardinality()))
		var _109_v74 D2
		_ = _109_v74
		_109_v74 = Companion_D2_.Create_DC4_(Companion_Default___.Fm12(_106_v72, (_104_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_104_v70), 0))).Int()).(bool), globalState))
		_109_v74 = Companion_Default___.Fm13(globalState)
	}
	var _110_v75 _dafny.Sequence
	_ = _110_v75
	_110_v75 = _dafny.SeqOf(p0, !(p0), false, false)
	var _111_v76 _dafny.Sequence
	_ = _111_v76
	_111_v76 = _dafny.SeqOf(_110_v75)
	var _112_v77 _dafny.MultiSet
	_ = _112_v77
	_112_v77 = _dafny.MultiSetOf(_13_v4, _13_v4)
	r0 = (_111_v76).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
		if (_112_v77).Contains((_dafny.Zero).Minus(_13_v4)) {
			return (_112_v77).Multiplicity((_dafny.Zero).Minus(_13_v4))
		}
		return (Companion_Default___.Fm14(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(798))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_113_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_114_i7 _dafny.Int) _dafny.Int {
				return _113_v4
			}
		})(_13_v4))), _13_v4, _13_v4, _112_v77, globalState)).Cardinality()
	})(), _dafny.IntOfUint32((_111_v76).Cardinality()))).Uint32()).(_dafny.Sequence)
	r1 = _13_v4
	var _115_v78 _dafny.Array
	_ = _115_v78
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_3
	var _nw19 _dafny.Array
	_ = _nw19
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw19 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Int = (func(_116_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_117_i8 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_117_i8, _116_v4)
			}
		})(_13_v4)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw19 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw19).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw19).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_115_v78 = _nw19
	var _118_v79 _dafny.Sequence
	_ = _118_v79
	_118_v79 = _dafny.SeqOf(_115_v78, _115_v78, _115_v78, _115_v78)
	r2 = (_118_v79).Select((Companion_Default___.SafeIndex(_13_v4, _dafny.IntOfUint32((_118_v79).Cardinality()))).Uint32()).(_dafny.Array)
	var _119_v80 _dafny.CodePoint
	_ = _119_v80
	_119_v80 = _dafny.CodePoint('f')
	r3 = _119_v80
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _120_v0 _dafny.Int
	_ = _120_v0
	_120_v0 = _dafny.IntOfInt64(691)
	var _121_v1 _dafny.Set
	_ = _121_v1
	_121_v1 = _dafny.SetOf(_120_v0)
	var _122_v2 _dafny.Map
	_ = _122_v2
	_122_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_121_v1, _120_v0)
	var _123_v3 _dafny.Sequence
	_ = _123_v3
	_123_v3 = _dafny.SeqOf(_120_v0, _120_v0)
	var _124_v4 _dafny.Map
	_ = _124_v4
	_124_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_120_v0)).Cardinality(), _123_v3)
	var _125_v5 _dafny.Array
	_ = _125_v5
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
	_ = _nw20
	_125_v5 = _nw20
	var _126_v6 D0
	_ = _126_v6
	_126_v6 = Companion_D0_.Create_DC0_(_125_v5)
	var _127_v7 _dafny.Array
	_ = _127_v7
	var _nwElement0_4 _dafny.Array = _125_v5
	_ = _nwElement0_4
	var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(4))
	_ = _nw21
	(_nw21).ArraySet1(_nwElement0_4, 0)
	(_nw21).ArraySet1(_125_v5, 1)
	(_nw21).ArraySet1(_125_v5, 2)
	(_nw21).ArraySet1((_126_v6).Dtor_cf0(), 3)
	_127_v7 = _nw21
	var _128_v8 _dafny.MultiSet
	_ = _128_v8
	_128_v8 = _dafny.MultiSetOf(_dafny.IntOfUint32((_123_v3).Cardinality()))
	var _129_v9 _dafny.Sequence
	_ = _129_v9
	_129_v9 = _dafny.UnicodeSeqOfUtf8Bytes("px")
	var _130_v10 _dafny.Map
	_ = _130_v10
	_130_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v7, _dafny.SetOf((_128_v8).Cardinality(), _120_v0, _dafny.IntOfUint32((_129_v9).Cardinality()), _120_v0))
	var _131_v11 _dafny.Array
	_ = _131_v11
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(23)
	_ = _len0_4
	var _nw22 _dafny.Array
	_ = _nw22
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw22 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.CodePoint = func(_132_i0 _dafny.Int) _dafny.CodePoint {
			return (_dafny.CodePoint('j'))
		}
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw22 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw22).ArraySet1CodePoint(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw22).ArraySet1CodePoint(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_131_v11 = _nw22
	var _133_v12 bool
	_ = _133_v12
	_133_v12 = true
	var _134_v13 _dafny.Map
	_ = _134_v13
	_134_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v12, _133_v12)
	var _135_v14 _dafny.Array
	_ = _135_v14
	var _nwElement0_5 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("eqobbdls")
	_ = _nwElement0_5
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(2))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_5, 0)
	(_nw23).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rgyeuu"), 1)
	_135_v14 = _nw23
	var _136_v15 _dafny.Set
	_ = _136_v15
	_136_v15 = _dafny.SetOf(_123_v3)
	var _137_globalState *GlobalState
	_ = _137_globalState
	var _nw24 *GlobalState = New_GlobalState_()
	_ = _nw24
	_nw24.Ctor__((_122_v2).Merge(_122_v2), _121_v1, _dafny.IntOfInt64(445), false, _dafny.IntOfInt64(988), _dafny.CodePoint('l'), _124_v4, _dafny.IntOfInt64(496), _dafny.IntOfInt64(-233), _dafny.IntOfInt64(938), (func() _dafny.Set {
		if (_130_v10).Contains(_127_v7) {
			return (_130_v10).Get(_127_v7).(_dafny.Set)
		}
		return _121_v1
	})(), _dafny.IntOfInt64(987), false, true, false, _dafny.IntOfInt64(99), false, _131_v11, _dafny.IntOfInt64(161), _134_v13, _135_v14, _dafny.UnicodeSeqOfUtf8Bytes("txxtamr"), _dafny.IntOfInt64(577), _136_v15, false, _dafny.IntOfInt64(181))
	_137_globalState = _nw24
	var _138_v16 _dafny.CodePoint
	_ = _138_v16
	_138_v16 = _dafny.CodePoint('j')
	var _139_v17 _dafny.Map
	_ = _139_v17
	_139_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v12, _120_v0), _138_v16)
	_139_v17 = _139_v17
	var _140_i1 _dafny.Int
	_ = _140_i1
	_140_i1 = _dafny.Zero
	{
		for _133_v12 {
			{
				if (_140_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_140_i1 = (_140_i1).Plus(_dafny.One)
				var _141_v18 _dafny.Sequence
				_ = _141_v18
				_141_v18 = _dafny.SeqOf(_133_v12)
				(_137_globalState).F14 = (_141_v18).Select((Companion_Default___.SafeIndex(_120_v0, _dafny.IntOfUint32((_141_v18).Cardinality()))).Uint32()).(bool)
				var _142_v19 D0
				_ = _142_v19
				_142_v19 = Companion_D0_.Create_DC1_(_133_v12)
				var _143_v20 _dafny.Map
				_ = _143_v20
				_143_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_142_v19, (_128_v8).Cardinality())
				var _144_v21 _dafny.Map
				_ = _144_v21
				_144_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v0, _133_v12)
				var _145_v22 _dafny.Set
				_ = _145_v22
				_145_v22 = _dafny.SetOf(_133_v12)
				var _146_v23 _dafny.Map
				_ = _146_v23
				_146_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_144_v21).Contains((_145_v22).Cardinality()) {
						return (_144_v21).Get((_145_v22).Cardinality()).(bool)
					}
					return _133_v12
				})(), _120_v0)
				_143_v20 = (_143_v20).Update(_142_v19, (_146_v23).Cardinality())
				var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_125_v5), 0))
				_ = _index12
				(_125_v5).ArraySet1(_dafny.IntOfUint32((_129_v9).Cardinality()), (_index12).Int())
				var _147_v24 _dafny.Map
				_ = _147_v24
				_147_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_141_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_141_v18).Cardinality()), _dafny.IntOfUint32((_141_v18).Cardinality()))).Uint32()).(bool), _138_v16)
				var _148_v25 _dafny.Map
				_ = _148_v25
				_148_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_147_v24).Cardinality(), _125_v5)
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(874), _dafny.ArrayLen((_125_v5), 0))
				_ = _index13
				(_125_v5).ArraySet1((_120_v0).Plus(((_148_v25).Merge(_148_v25)).Cardinality()), (_index13).Int())
				var _149_v26 _dafny.Map
				_ = _149_v26
				_149_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v16, _133_v12)
				var _150_v27 _dafny.CodePoint
				_ = _150_v27
				_150_v27 = _138_v16
				_149_v26 = (_149_v26).Update(_150_v27, _133_v12)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _151_v28 _dafny.Sequence
	_ = _151_v28
	var _152_v29 _dafny.Int
	_ = _152_v29
	var _153_v30 _dafny.Array
	_ = _153_v30
	var _154_v31 _dafny.CodePoint
	_ = _154_v31
	var _out0 _dafny.Sequence
	_ = _out0
	var _out1 _dafny.Int
	_ = _out1
	var _out2 _dafny.Array
	_ = _out2
	var _out3 _dafny.CodePoint
	_ = _out3
	_out0, _out1, _out2, _out3 = Companion_Default___.M0(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gb")).Cardinality()), _133_v12, _137_globalState), _137_globalState)
	_151_v28 = _out0
	_152_v29 = _out1
	_153_v30 = _out2
	_154_v31 = _out3
	(_137_globalState).F22 = (_dafny.Zero).Minus(_152_v29)
	if _133_v12 {
		(_137_globalState).F14 = _133_v12
		var _155_v32 D0
		_ = _155_v32
		_155_v32 = Companion_D0_.Create_DC1_(_133_v12)
		var _source2 D0 = _155_v32
		_ = _source2
		if _source2.Is_DC1() {
			var _156___mcc_h0 bool = _source2.Get_().(D0_DC1).Cf1
			_ = _156___mcc_h0
			var _157_cf1 bool = _156___mcc_h0
			_ = _157_cf1
			var _158_v33 _dafny.Array
			_ = _158_v33
			var _nwElement0_6 bool = (_152_v29).Cmp(_152_v29) != 0
			_ = _nwElement0_6
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(8))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_6, 0)
			(_nw25).ArraySet1(true, 1)
			(_nw25).ArraySet1(_133_v12, 2)
			(_nw25).ArraySet1(_133_v12, 3)
			(_nw25).ArraySet1((_121_v1).IsDisjointFrom(_121_v1), 4)
			(_nw25).ArraySet1(_157_cf1, 5)
			(_nw25).ArraySet1((_dafny.IntOfInt64(492)).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_120_v0))) == 0, 6)
			(_nw25).ArraySet1(_157_cf1, 7)
			_158_v33 = _nw25
			_158_v33 = _158_v33
			(_137_globalState).F14 = (!(_133_v12)) || (_157_cf1)
			(_137_globalState).F22 = _152_v29
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_125_v5), 0))
			_ = _index14
			(_125_v5).ArraySet1(_152_v29, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_125_v5), 0))
			_ = _index15
			(_125_v5).ArraySet1(_dafny.IntOfInt64(882), (_index15).Int())
		} else if _source2.Is_DC2() {
			var _159___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC2).Cf2
			_ = _159___mcc_h1
			var _160___mcc_h2 _dafny.Int = _source2.Get_().(D0_DC2).Cf3
			_ = _160___mcc_h2
			var _161_cf3 _dafny.Int = _160___mcc_h2
			_ = _161_cf3
			var _162_cf2 _dafny.Int = _159___mcc_h1
			_ = _162_cf2
			var _163_v34 _dafny.MultiSet
			_ = _163_v34
			_163_v34 = _dafny.MultiSetOf(_133_v12, !(_133_v12))
			var _164_v36 _dafny.Map
			_ = _164_v36
			_164_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_cf2, _120_v0)
			var _165_v37 _dafny.Map
			_ = _165_v37
			_165_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v12, Companion_Default___.Fm1(_dafny.IntOfInt64(716), (_151_v28).Select((Companion_Default___.SafeIndex(_162_cf2, _dafny.IntOfUint32((_151_v28).Cardinality()))).Uint32()).(bool), _129_v9, _133_v12, _137_globalState))
			var _166_v40 _dafny.Array
			_ = _166_v40
			var _nwElement0_7 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_163_v34).Cardinality(), _161_cf3)).Merge(func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(717), _dafny.IntOfInt64(364))); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _167_v35 _dafny.Int
					_167_v35 = interface{}(_compr_3).(_dafny.Int)
					if ((_dafny.IntOfInt64(717)).Cmp(_167_v35) <= 0) && ((_167_v35).Cmp(_dafny.IntOfInt64(364)) < 0) {
						_coll3.Add((_167_v35).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v12, _162_cf2), true)).Cardinality()), (_dafny.MultiSetOf(_162_cf2, (_dafny.Zero).Minus(_120_v0))).Cardinality())
					}
				}
				return _coll3.ToMap()
			}())
			_ = _nwElement0_7
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(19))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_7, 0)
			(_nw26).ArraySet1(_164_v36, 1)
			(_nw26).ArraySet1(_164_v36, 2)
			(_nw26).ArraySet1(_164_v36, 3)
			(_nw26).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_cf2, _152_v29)).Merge(_164_v36), 4)
			(_nw26).ArraySet1(_164_v36, 5)
			(_nw26).ArraySet1(_164_v36, 6)
			(_nw26).ArraySet1((_164_v36).Update((_165_v37).Cardinality(), _161_cf3), 7)
			(_nw26).ArraySet1(_164_v36, 8)
			(_nw26).ArraySet1((func() _dafny.Map {
				if _133_v12 {
					return _164_v36
				}
				return (Companion_D2_.Create_DC4_(_164_v36)).Dtor_cf5()
			})(), 9)
			(_nw26).ArraySet1(_164_v36, 10)
			(_nw26).ArraySet1(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-473), _dafny.IntOfInt64(31))); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _168_v38 _dafny.Int
					_168_v38 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(-473)).Cmp(_168_v38) <= 0) && ((_168_v38).Cmp(_dafny.IntOfInt64(31)) < 0) {
						_coll4.Add((_168_v38).Plus((_dafny.Zero).Minus(_152_v29)), (_dafny.Zero).Minus(_162_cf2))
					}
				}
				return _coll4.ToMap()
			}(), 11)
			(_nw26).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v29, _162_cf2)).Merge(_164_v36), 12)
			(_nw26).ArraySet1((_164_v36).Merge(_164_v36), 13)
			(_nw26).ArraySet1(_164_v36, 14)
			(_nw26).ArraySet1(_164_v36, 15)
			(_nw26).ArraySet1(_164_v36, 16)
			(_nw26).ArraySet1(_164_v36, 17)
			(_nw26).ArraySet1((func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(((_164_v36).Update(_dafny.IntOfUint32((_151_v28).Cardinality()), _161_cf3)).Keys().Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _169_v39 _dafny.Int
					_169_v39 = interface{}(_compr_5).(_dafny.Int)
					if ((_164_v36).Update(_dafny.IntOfUint32((_151_v28).Cardinality()), _161_cf3)).Contains(_169_v39) {
						_coll5.Add(Companion_Default___.SafeModuloInt(_169_v39, _dafny.IntOfInt64(990)), _dafny.IntOfUint32((_151_v28).Cardinality()))
					}
				}
				return _coll5.ToMap()
			}()).Update(_162_cf2, (_dafny.Zero).Minus(_dafny.IntOfInt64(-994))), 18)
			_166_v40 = _nw26
			var _170_v41 _dafny.Map
			_ = _170_v41
			_170_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v29, _155_v32)
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_153_v30), 0))
			_ = _index16
			(_153_v30).ArraySet1(Companion_Default___.Fm1((_170_v41).Cardinality(), _133_v12, _dafny.UnicodeSeqOfUtf8Bytes("vjvcgp"), _133_v12, _137_globalState), (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_153_v30), 0))
			_ = _index17
			var _rhs14 _dafny.Int = _162_cf2
			_ = _rhs14
			var _rhs15 _dafny.Int = _162_cf2
			_ = _rhs15
			var _rhs16 _dafny.Array = _166_v40
			_ = _rhs16
			var _rhs17 bool = !(_133_v12)
			_ = _rhs17
			var _rhs18 _dafny.Int = _120_v0
			_ = _rhs18
			var _lhs11 *GlobalState = _137_globalState
			_ = _lhs11
			var _lhs12 *GlobalState = _137_globalState
			_ = _lhs12
			var _lhs13 _dafny.Array = _153_v30
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_153_v30), 0))
			_ = _lhs14
			_152_v29 = _rhs14
			_lhs11.F22 = _rhs15
			_166_v40 = _rhs16
			_lhs12.F14 = _rhs17
			(_lhs13).ArraySet1(_rhs18, (_lhs14).Int())
			var _171_v42 _dafny.Set
			_ = _171_v42
			_171_v42 = _dafny.SetOf((_164_v36).Merge((_164_v36).Update(_dafny.IntOfInt64(12), (_153_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_153_v30), 0))).Int()).(_dafny.Int))))
			var _172_v43 _dafny.Map
			_ = _172_v43
			_172_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v36, _133_v12)
			var _pat_let_tv0 = _171_v42
			_ = _pat_let_tv0
			var _pat_let_tv1 = _171_v42
			_ = _pat_let_tv1
			_171_v42 = ((func(_pat_let0_0 D3) D3 {
				return func(_175_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let3_0 _dafny.Set) D3 {
						return func(_176_dt__update_hcf12_h1 _dafny.Set) D3 {
							return Companion_D3_.Create_DC8_(_176_dt__update_hcf12_h1)
						}(_pat_let3_0)
					}(_pat_let_tv1)
				}(_pat_let0_0)
			}(func(_pat_let1_0 D3) D3 {
				return func(_173_dt__update__tmp_h0 D3) D3 {
					return func(_pat_let2_0 _dafny.Set) D3 {
						return func(_174_dt__update_hcf12_h0 _dafny.Set) D3 {
							return Companion_D3_.Create_DC8_(_174_dt__update_hcf12_h0)
						}(_pat_let2_0)
					}(_pat_let_tv0)
				}(_pat_let1_0)
			}(Companion_D3_.Create_DC8_(_dafny.SetOf(_164_v36))))).Dtor_cf12()).Intersection((_171_v42).Intersection(func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_172_v43).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _177_v44 _dafny.Map
					_177_v44 = interface{}(_compr_6).(_dafny.Map)
					if (_172_v43).Contains(_177_v44) {
						_coll6.Add(_177_v44)
					}
				}
				return _coll6.ToSet()
			}()))
			(_137_globalState).F22 = Companion_Default___.Fm1(Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(true)).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_152_v29))), _133_v12, Companion_Default___.Fm2(_152_v29, _120_v0, (_151_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg9 _dafny.Int) interface{} {
					return coer9(arg9)
				}
			}((func(_178_v29 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_179_i2 _dafny.Int) _dafny.Int {
					return _178_v29
				}
			})(_152_v29)))).Cardinality()), _dafny.IntOfUint32((_151_v28).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus(_161_cf3), _137_globalState), _133_v12, _137_globalState)
			var _180_v45 _dafny.Set
			_ = _180_v45
			_180_v45 = _dafny.SetOf(true, _133_v12)
			_180_v45 = ((_180_v45).Difference(Companion_Default___.Fm3((_153_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_153_v30), 0))).Int()).(_dafny.Int), _133_v12, _137_globalState))).Union((_dafny.SetOf(_133_v12, _133_v12)).Difference(_180_v45))
		} else {
			var _181___mcc_h3 _dafny.Array = _source2.Get_().(D0_DC0).Cf0
			_ = _181___mcc_h3
			var _182_cf0 _dafny.Array = _181___mcc_h3
			_ = _182_cf0
			(_137_globalState).F25 = (_121_v1).Cardinality()
			(_137_globalState).F9 = _120_v0
			var _183_v46 _dafny.Sequence
			_ = _183_v46
			var _184_v47 _dafny.Int
			_ = _184_v47
			var _185_v48 _dafny.Array
			_ = _185_v48
			var _186_v49 _dafny.CodePoint
			_ = _186_v49
			var _out4 _dafny.Sequence
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			var _out6 _dafny.Array
			_ = _out6
			var _out7 _dafny.CodePoint
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0(!(_133_v12), _137_globalState)
			_183_v46 = _out4
			_184_v47 = _out5
			_185_v48 = _out6
			_186_v49 = _out7
			var _187_v50 _dafny.MultiSet
			_ = _187_v50
			_187_v50 = _dafny.MultiSetOf(_133_v12, _133_v12, _133_v12, _133_v12, _133_v12)
			var _188_v51 _dafny.CodePoint
			_ = _188_v51
			_188_v51 = _138_v16
			var _rhs19 _dafny.Int = ((_187_v50).Cardinality()).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_152_v29, _dafny.IntOfUint32((_123_v3).Cardinality())))))
			_ = _rhs19
			var _rhs20 _dafny.CodePoint = Companion_Default___.Fm4(_188_v51, _dafny.SeqOf(true), _137_globalState)
			_ = _rhs20
			_184_v47 = _rhs19
			_186_v49 = _rhs20
		}
		var _189_v52 _dafny.Map
		_ = _189_v52
		_189_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_120_v0), _120_v0)
		var _190_v54 _dafny.Sequence
		_ = _190_v54
		_190_v54 = _dafny.SeqOf(_123_v3, _dafny.Companion_Sequence_.Update(_123_v3, (Companion_Default___.SafeIndex(_120_v0, _dafny.IntOfUint32((_123_v3).Cardinality()))).Uint32(), _120_v0))
		_189_v52 = (_189_v52).Merge((func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate((_190_v54).Elements()); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _191_v53 _dafny.Sequence
				_191_v53 = interface{}(_compr_7).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_190_v54, _191_v53) {
					_coll7.Add(_191_v53, _120_v0)
				}
			}
			return _coll7.ToMap()
		}()).Merge(_189_v52))
		var _192_v55 _dafny.MultiSet
		_ = _192_v55
		_192_v55 = _dafny.MultiSetOf(_133_v12, _133_v12, true, _133_v12)
		var _193_v56 D0
		_ = _193_v56
		_193_v56 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(623), Companion_Default___.Fm1(_152_v29, _133_v12, _129_v9, _133_v12, _137_globalState))
		var _source3 D2 = Companion_D2_.Create_DC5_((_121_v1).IsProperSubsetOf(_121_v1), ((func() _dafny.Int {
			if (_192_v55).Contains(false) {
				return (_192_v55).Multiplicity(false)
			}
			return (_193_v56).Dtor_cf2()
		})()).Cmp(_152_v29) < 0, true, _152_v29, (Companion_Default___.Fm1(_152_v29, false, Companion_Default___.Fm2(_152_v29, _152_v29, _133_v12, _120_v0, _137_globalState), _133_v12, _137_globalState)).Cmp(_120_v0) < 0)
		_ = _source3
		if _source3.Is_DC5() {
			var _194___mcc_h4 bool = _source3.Get_().(D2_DC5).Cf6
			_ = _194___mcc_h4
			var _195___mcc_h5 bool = _source3.Get_().(D2_DC5).Cf7
			_ = _195___mcc_h5
			var _196___mcc_h6 bool = _source3.Get_().(D2_DC5).Cf8
			_ = _196___mcc_h6
			var _197___mcc_h7 _dafny.Int = _source3.Get_().(D2_DC5).Cf9
			_ = _197___mcc_h7
			var _198___mcc_h8 bool = _source3.Get_().(D2_DC5).Cf10
			_ = _198___mcc_h8
			var _199_cf10 bool = _198___mcc_h8
			_ = _199_cf10
			var _200_cf9 _dafny.Int = _197___mcc_h7
			_ = _200_cf9
			var _201_cf8 bool = _196___mcc_h6
			_ = _201_cf8
			var _202_cf7 bool = _195___mcc_h5
			_ = _202_cf7
			var _203_cf6 bool = _194___mcc_h4
			_ = _203_cf6
			var _204_v57 _dafny.Set
			_ = _204_v57
			_204_v57 = _dafny.SetOf(_121_v1)
			_202_cf7 = ((_204_v57).Intersection(_204_v57)).IsSubsetOf(_204_v57)
			var _rhs21 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_200_cf9).Times(_152_v29)))
			_ = _rhs21
			var _rhs22 _dafny.Array = (func() _dafny.Array {
				if _133_v12 {
					return _125_v5
				}
				return (func() _dafny.Array {
					if _201_cf8 {
						return _125_v5
					}
					return _125_v5
				})()
			})()
			_ = _rhs22
			var _rhs23 _dafny.Sequence = _129_v9
			_ = _rhs23
			var _rhs24 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_200_cf9, _120_v0, _202_cf7, _dafny.IntOfUint32((_151_v28).Cardinality()), _137_globalState), _129_v9)
			_ = _rhs24
			var _lhs15 *GlobalState = _137_globalState
			_ = _lhs15
			var _lhs16 *GlobalState = _137_globalState
			_ = _lhs16
			_152_v29 = _rhs21
			_125_v5 = _rhs22
			_lhs15.F21 = _rhs23
			_lhs16.F21 = _rhs24
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_125_v5), 0))
			_ = _index18
			(_125_v5).ArraySet1(_120_v0, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_125_v5), 0))
			_ = _index19
			(_125_v5).ArraySet1(_152_v29, (_index19).Int())
			var _205_v58 _dafny.Sequence
			_ = _205_v58
			var _206_v59 _dafny.Int
			_ = _206_v59
			var _207_v60 _dafny.Array
			_ = _207_v60
			var _208_v61 _dafny.CodePoint
			_ = _208_v61
			var _out8 _dafny.Sequence
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Array
			_ = _out10
			var _out11 _dafny.CodePoint
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0(!(_133_v12), _137_globalState)
			_205_v58 = _out8
			_206_v59 = _out9
			_207_v60 = _out10
			_208_v61 = _out11
		} else if _source3.Is_DC6() {
			var _209_v62 *C0
			_ = _209_v62
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__()
			_209_v62 = _nw27
			var _210_v63 _dafny.MultiSet
			_ = _210_v63
			_210_v63 = _dafny.MultiSetOf(_131_v11)
			var _211_v64 _dafny.Map
			_ = _211_v64
			_211_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v63, _209_v62)
			var _212_v65 _dafny.Sequence
			_ = _212_v65
			_212_v65 = _dafny.SeqOf(_209_v62)
			var _213_v66 _dafny.Map
			_ = _213_v66
			_213_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v12, _209_v62)
			var _214_v67 _dafny.Array
			_ = _214_v67
			var _nwElement0_8 *C0 = _209_v62
			_ = _nwElement0_8
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(26))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_8, 0)
			(_nw28).ArraySet1(_209_v62, 1)
			(_nw28).ArraySet1(_209_v62, 2)
			(_nw28).ArraySet1(_209_v62, 3)
			(_nw28).ArraySet1(_209_v62, 4)
			(_nw28).ArraySet1(_209_v62, 5)
			(_nw28).ArraySet1(_209_v62, 6)
			(_nw28).ArraySet1(_209_v62, 7)
			(_nw28).ArraySet1((func() *C0 {
				if Companion_Default___.Fm0(_120_v0, _133_v12, _137_globalState) {
					return _209_v62
				}
				return _209_v62
			})(), 8)
			(_nw28).ArraySet1((func() *C0 {
				if (_211_v64).Contains(_210_v63) {
					return (_211_v64).Get(_210_v63).(*C0)
				}
				return (_212_v65).Select((Companion_Default___.SafeIndex(_120_v0, _dafny.IntOfUint32((_212_v65).Cardinality()))).Uint32()).(*C0)
			})(), 9)
			(_nw28).ArraySet1(_209_v62, 10)
			(_nw28).ArraySet1(_209_v62, 11)
			(_nw28).ArraySet1(_209_v62, 12)
			(_nw28).ArraySet1(_209_v62, 13)
			(_nw28).ArraySet1(_209_v62, 14)
			(_nw28).ArraySet1(_209_v62, 15)
			(_nw28).ArraySet1(_209_v62, 16)
			(_nw28).ArraySet1(_209_v62, 17)
			(_nw28).ArraySet1(_209_v62, 18)
			(_nw28).ArraySet1(_209_v62, 19)
			(_nw28).ArraySet1((func() *C0 {
				if (_213_v66).Contains(Companion_Default___.Fm0(_120_v0, _133_v12, _137_globalState)) {
					return (_213_v66).Get(Companion_Default___.Fm0(_120_v0, _133_v12, _137_globalState)).(*C0)
				}
				return _209_v62
			})(), 20)
			(_nw28).ArraySet1((func() *C0 {
				if (_213_v66).Contains(_133_v12) {
					return (_213_v66).Get(_133_v12).(*C0)
				}
				return _209_v62
			})(), 21)
			(_nw28).ArraySet1(_209_v62, 22)
			(_nw28).ArraySet1(_209_v62, 23)
			(_nw28).ArraySet1(_209_v62, 24)
			(_nw28).ArraySet1(_209_v62, 25)
			_214_v67 = _nw28
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_214_v67), 0))
			_ = _index20
			(_214_v67).ArraySet1(_209_v62, (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(521), _dafny.ArrayLen((_214_v67), 0))
			_ = _index21
			(_214_v67).ArraySet1((func() *C0 {
				if _133_v12 {
					return _209_v62
				}
				return _209_v62
			})(), (_index21).Int())
			var _215_v68 *C0
			_ = _215_v68
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__()
			_215_v68 = _nw29
			var _216_v69 _dafny.Sequence
			_ = _216_v69
			var _217_v70 _dafny.Int
			_ = _217_v70
			var _218_v71 _dafny.Array
			_ = _218_v71
			var _219_v72 _dafny.CodePoint
			_ = _219_v72
			var _out12 _dafny.Sequence
			_ = _out12
			var _out13 _dafny.Int
			_ = _out13
			var _out14 _dafny.Array
			_ = _out14
			var _out15 _dafny.CodePoint
			_ = _out15
			_out12, _out13, _out14, _out15 = Companion_Default___.M0(_133_v12, _137_globalState)
			_216_v69 = _out12
			_217_v70 = _out13
			_218_v71 = _out14
			_219_v72 = _out15
		} else if _source3.Is_DC4() {
			var _220___mcc_h9 _dafny.Map = _source3.Get_().(D2_DC4).Cf5
			_ = _220___mcc_h9
			var _221_cf5 _dafny.Map = _220___mcc_h9
			_ = _221_cf5
			var _222_v73 *C0
			_ = _222_v73
			var _nw30 *C0 = New_C0_()
			_ = _nw30
			_nw30.Ctor__()
			_222_v73 = _nw30
			var _223_v74 _dafny.Map
			_ = _223_v74
			_223_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v73, _222_v73)
			var _224_v75 D4
			_ = _224_v75
			_224_v75 = Companion_D4_.Create_DC10_(_222_v73)
			var _rhs25 *C0 = (func() *C0 {
				if (_223_v74).Contains((_224_v75).Dtor_cf15()) {
					return (_223_v74).Get((_224_v75).Dtor_cf15()).(*C0)
				}
				return _222_v73
			})()
			_ = _rhs25
			var _rhs26 _dafny.Int = (func() _dafny.Int {
				if (_192_v55).Contains((func() bool {
					if (_134_v13).Contains(_133_v12) {
						return (_134_v13).Get(_133_v12).(bool)
					}
					return _133_v12
				})()) {
					return (_192_v55).Multiplicity((func() bool {
						if (_134_v13).Contains(_133_v12) {
							return (_134_v13).Get(_133_v12).(bool)
						}
						return _133_v12
					})())
				}
				return Companion_Default___.SafeDivisionInt(_120_v0, _152_v29)
			})()
			_ = _rhs26
			var _lhs17 *GlobalState = _137_globalState
			_ = _lhs17
			_222_v73 = _rhs25
			_lhs17.F22 = _rhs26
			var _225_v76 _dafny.Array
			_ = _225_v76
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
			_ = _nw31
			_225_v76 = _nw31
			var _rhs27 _dafny.Int = _152_v29
			_ = _rhs27
			var _rhs28 bool = ((_123_v3).Select((Companion_Default___.SafeIndex(_120_v0, _dafny.IntOfUint32((_123_v3).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(-295)) >= 0
			_ = _rhs28
			var _rhs29 _dafny.CodePoint = _154_v31
			_ = _rhs29
			var _rhs30 bool = !(_133_v12)
			_ = _rhs30
			var _rhs31 _dafny.Array = _225_v76
			_ = _rhs31
			var _lhs18 *GlobalState = _137_globalState
			_ = _lhs18
			var _lhs19 *GlobalState = _137_globalState
			_ = _lhs19
			var _lhs20 *GlobalState = _137_globalState
			_ = _lhs20
			_lhs18.F25 = _rhs27
			_lhs19.F14 = _rhs28
			_154_v31 = _rhs29
			_lhs20.F14 = _rhs30
			_225_v76 = _rhs31
			(_137_globalState).F21 = _129_v9
			var _226_v77 _dafny.Sequence
			_ = _226_v77
			var _227_v78 _dafny.Int
			_ = _227_v78
			var _228_v79 _dafny.Array
			_ = _228_v79
			var _229_v80 _dafny.CodePoint
			_ = _229_v80
			var _out16 _dafny.Sequence
			_ = _out16
			var _out17 _dafny.Int
			_ = _out17
			var _out18 _dafny.Array
			_ = _out18
			var _out19 _dafny.CodePoint
			_ = _out19
			_out16, _out17, _out18, _out19 = Companion_Default___.M0((_dafny.IntOfInt64(845)).Cmp(_152_v29) > 0, _137_globalState)
			_226_v77 = _out16
			_227_v78 = _out17
			_228_v79 = _out18
			_229_v80 = _out19
		} else {
			var _230___mcc_h10 D2 = _source3.Get_().(D2_DC7).Cf11
			_ = _230___mcc_h10
			var _231_cf11 D2 = _230___mcc_h10
			_ = _231_cf11
			var _232_v81 _dafny.Array
			_ = _232_v81
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw32
			_232_v81 = _nw32
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_232_v81), 0))
			_ = _index22
			(_232_v81).ArraySet1((_155_v32).Dtor_cf1(), (_index22).Int())
			var _233_v82 _dafny.Sequence
			_ = _233_v82
			_233_v82 = _dafny.SeqOf(_dafny.MultiSetFromSeq(Companion_Default___.Fm6(_120_v0, _137_globalState)))
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_232_v81), 0))
			_ = _index23
			(_232_v81).ArraySet1((((_233_v82).Select((Companion_Default___.SafeIndex(_152_v29, _dafny.IntOfUint32((_233_v82).Cardinality()))).Uint32()).(_dafny.MultiSet)).Intersection((_192_v55).Update(_133_v12, Companion_Default___.Abs((_dafny.Zero).Minus(_120_v0))))).IsSubsetOf((func() _dafny.MultiSet {
				if _133_v12 {
					return _192_v55
				}
				return _dafny.MultiSetOf(_133_v12, _133_v12)
			})()), (_index23).Int())
			(_137_globalState).F24 = (func() bool {
				if _133_v12 {
					return (_dafny.MultiSetOf(_152_v29)).IsProperSubsetOf(_128_v8)
				}
				return _133_v12
			})()
			var _234_v83 _dafny.Map
			_ = _234_v83
			_234_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_123_v3).Cardinality()), _152_v29)
			var _235_v84 _dafny.Set
			_ = _235_v84
			_235_v84 = _dafny.SetOf(_133_v12)
			var _236_v85 _dafny.Set
			_ = _236_v85
			_236_v85 = _dafny.SetOf(_234_v83, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_129_v9).Cardinality()), _152_v29), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_235_v84).Cardinality(), _152_v29), _234_v83, _234_v83)
			var _237_v86 D3
			_ = _237_v86
			_237_v86 = Companion_D3_.Create_DC8_(_236_v85)
			_237_v86 = _237_v86
			(_137_globalState).F14 = (_232_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_232_v81), 0))).Int()).(bool)
		}
		var _238_v87 _dafny.Sequence
		_ = _238_v87
		var _239_v88 _dafny.Int
		_ = _239_v88
		var _240_v89 _dafny.Array
		_ = _240_v89
		var _241_v90 _dafny.CodePoint
		_ = _241_v90
		var _out20 _dafny.Sequence
		_ = _out20
		var _out21 _dafny.Int
		_ = _out21
		var _out22 _dafny.Array
		_ = _out22
		var _out23 _dafny.CodePoint
		_ = _out23
		_out20, _out21, _out22, _out23 = Companion_Default___.M0(_133_v12, _137_globalState)
		_238_v87 = _out20
		_239_v88 = _out21
		_240_v89 = _out22
		_241_v90 = _out23
	} else {
		var _source4 _dafny.CodePoint = _154_v31
		_ = _source4
		var _242___mcc_h11 _dafny.CodePoint = _source4
		_ = _242___mcc_h11
		var _243_cf4 _dafny.CodePoint = _242___mcc_h11
		_ = _243_cf4
		var _244_v91 _dafny.Set
		_ = _244_v91
		_244_v91 = _dafny.SetOf(_133_v12)
		_244_v91 = _244_v91
		var _245_v92 _dafny.Array
		_ = _245_v92
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw33
		_245_v92 = _nw33
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_245_v92), 0))
		_ = _index24
		(_245_v92).ArraySet1(_133_v12, (_index24).Int())
		var _246_v93 _dafny.Array
		_ = _246_v93
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
		_ = _nw34
		_246_v93 = _nw34
		var _247_v94 _dafny.Sequence
		_ = _247_v94
		_247_v94 = _dafny.SeqOf(_125_v5, _153_v30)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_246_v93), 0))
		_ = _index25
		(_246_v93).ArraySet1(_247_v94, (_index25).Int())
		var _248_v95 _dafny.Map
		_ = _248_v95
		_248_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_133_v12), _dafny.IntOfUint32((_151_v28).Cardinality()))
		var _249_v96 _dafny.Map
		_ = _249_v96
		_249_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_248_v95).Cardinality(), _133_v12), _153_v30)
		var _250_v97 _dafny.Map
		_ = _250_v97
		_250_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v29, _133_v12)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_245_v92), 0))
		_ = _index26
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_246_v93), 0))
		_ = _index27
		var _rhs32 _dafny.Array = (func() _dafny.Array {
			if (_249_v96).Contains(_250_v97) {
				return (_249_v96).Get(_250_v97).(_dafny.Array)
			}
			return (func() _dafny.Array {
				if _133_v12 {
					return _125_v5
				}
				return _153_v30
			})()
		})()
		_ = _rhs32
		var _rhs33 bool = _133_v12
		_ = _rhs33
		var _rhs34 bool = !(_133_v12)
		_ = _rhs34
		var _rhs35 _dafny.Int = (_123_v3).Select((Companion_Default___.SafeIndex(_152_v29, _dafny.IntOfUint32((_123_v3).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs35
		var _rhs36 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_247_v94, _247_v94)
		_ = _rhs36
		var _lhs21 *GlobalState = _137_globalState
		_ = _lhs21
		var _lhs22 _dafny.Array = _245_v92
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_245_v92), 0))
		_ = _lhs23
		var _lhs24 *GlobalState = _137_globalState
		_ = _lhs24
		var _lhs25 _dafny.Array = _246_v93
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_246_v93), 0))
		_ = _lhs26
		_153_v30 = _rhs32
		_lhs21.F14 = _rhs33
		(_lhs22).ArraySet1(_rhs34, (_lhs23).Int())
		_lhs24.F22 = _rhs35
		(_lhs25).ArraySet1(_rhs36, (_lhs26).Int())
		var _251_v98 *C0
		_ = _251_v98
		var _nw35 *C0 = New_C0_()
		_ = _nw35
		_nw35.Ctor__()
		_251_v98 = _nw35
		var _252_v99 _dafny.Sequence
		_ = _252_v99
		var _253_v100 _dafny.Int
		_ = _253_v100
		var _254_v101 _dafny.Array
		_ = _254_v101
		var _255_v102 _dafny.CodePoint
		_ = _255_v102
		var _out24 _dafny.Sequence
		_ = _out24
		var _out25 _dafny.Int
		_ = _out25
		var _out26 _dafny.Array
		_ = _out26
		var _out27 _dafny.CodePoint
		_ = _out27
		_out24, _out25, _out26, _out27 = Companion_Default___.M0(Companion_Default___.Fm0(_120_v0, (_245_v92).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_245_v92), 0))).Int()).(bool), _137_globalState), _137_globalState)
		_252_v99 = _out24
		_253_v100 = _out25
		_254_v101 = _out26
		_255_v102 = _out27
		_129_v9 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aml"), _129_v9), _dafny.Companion_Sequence_.Concatenate(_129_v9, _129_v9))
		(_137_globalState).F14 = (func() bool {
			if (_152_v29).Cmp(_120_v0) != 0 {
				return _133_v12
			}
			return (func() bool {
				if _133_v12 {
					return _133_v12
				}
				return _133_v12
			})()
		})()
		(_137_globalState).F24 = _133_v12
		var _256_v103 _dafny.Map
		_ = _256_v103
		_256_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v29, _129_v9)
		_256_v103 = (_256_v103).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v0, !(false))).Cardinality(), _129_v9)
	}
	for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_125_v5), 0))); ; {
		_guard_loop_0, _ok8 := _iter8()
		if !_ok8 {
			break
		}
		var _257_i3 _dafny.Int
		_257_i3 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_257_i3).Sign() != -1) && ((_257_i3).Cmp(_dafny.ArrayLen((_125_v5), 0)) < 0)) {
			(_125_v5).ArraySet1((_257_i3).Plus(_dafny.IntOfInt64(110)), (_257_i3).Int())
		}
	}
	var _258_v104 _dafny.Map
	_ = _258_v104
	_258_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_152_v29), _154_v31)
	(_137_globalState).F24 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_129_v9, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_120_v0), _dafny.IntOfUint32((_129_v9).Cardinality()))).Uint32(), _154_v31), (func() _dafny.CodePoint {
		if (_258_v104).Contains(_120_v0) {
			return (_258_v104).Get(_120_v0).(_dafny.CodePoint)
		}
		return _dafny.CodePoint('g')
	})())
	var _259_i4 _dafny.Int
	_ = _259_i4
	_259_i4 = _dafny.Zero
	{
		for _133_v12 {
			{
				if (_259_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_259_i4 = (_259_i4).Plus(_dafny.One)
				var _260_v105 _dafny.Sequence
				_ = _260_v105
				_260_v105 = _dafny.SeqOf(_125_v5, _125_v5, _125_v5)
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_153_v30), 0))
				_ = _index28
				(_153_v30).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf((_151_v28).Select((Companion_Default___.SafeIndex(_152_v29, _dafny.IntOfUint32((_151_v28).Cardinality()))).Uint32()).(bool))).Cardinality())).Minus(_dafny.IntOfUint32((_260_v105).Cardinality())), (_index28).Int())
				var _261_v106 _dafny.Sequence
				_ = _261_v106
				_261_v106 = _dafny.SeqOf(_134_v13)
				var _262_v107 _dafny.Map
				_ = _262_v107
				_262_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_261_v106).Cardinality()), _133_v12)
				var _263_v108 _dafny.Map
				_ = _263_v108
				_263_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_262_v107, _120_v0)
				var _264_v111 _dafny.Sequence
				_ = _264_v111
				_264_v111 = _dafny.SeqOf(_262_v107, _262_v107, func() _dafny.Map {
					var _coll8 = _dafny.NewMapBuilder()
					_ = _coll8
					for _iter9 := _dafny.Iterate((_dafny.MultiSetOf(_152_v29, _120_v0)).Elements()); ; {
						_compr_8, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _265_v110 _dafny.Int
						_265_v110 = interface{}(_compr_8).(_dafny.Int)
						if (_dafny.MultiSetOf(_152_v29, _120_v0)).Contains(_265_v110) {
							_coll8.Add((_265_v110).Plus(_120_v0), _133_v12)
						}
					}
					return _coll8.ToMap()
				}())
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_153_v30), 0))
				_ = _index29
				var _rhs37 _dafny.Int = (((_263_v108).Merge(func() _dafny.Map {
					var _coll9 = _dafny.NewMapBuilder()
					_ = _coll9
					for _iter10 := _dafny.Iterate((_264_v111).Elements()); ; {
						_compr_9, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _266_v109 _dafny.Map
						_266_v109 = interface{}(_compr_9).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(_264_v111, _266_v109) {
							_coll9.Add(_266_v109, _152_v29)
						}
					}
					return _coll9.ToMap()
				}())).Merge(_263_v108)).Cardinality()
				_ = _rhs37
				var _rhs38 _dafny.Int = _152_v29
				_ = _rhs38
				var _lhs27 *GlobalState = _137_globalState
				_ = _lhs27
				var _lhs28 _dafny.Array = _153_v30
				_ = _lhs28
				var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_153_v30), 0))
				_ = _lhs29
				_lhs27.F9 = _rhs37
				(_lhs28).ArraySet1(_rhs38, (_lhs29).Int())
				var _267_v112 _dafny.Set
				_ = _267_v112
				_267_v112 = _dafny.SetOf(_133_v12, _133_v12)
				var _268_v113 _dafny.Map
				_ = _268_v113
				_268_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v29, _138_v16), true)
				(_137_globalState).F9 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.Fm1((_267_v112).Cardinality(), _133_v12, _129_v9, _133_v12, _137_globalState)), (_268_v113).Cardinality()), (_153_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_153_v30), 0))).Int()).(_dafny.Int))).Cardinality()
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_153_v30), 0))
				_ = _index30
				(_153_v30).ArraySet1(Companion_Default___.SafeModuloInt(_152_v29, (_262_v107).Cardinality()), (_index30).Int())
				(_137_globalState).F22 = (_dafny.Zero).Minus((_153_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_153_v30), 0))).Int()).(_dafny.Int))
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	(_137_globalState).F25 = _152_v29
	var _hi1 _dafny.Int = _dafny.IntOfInt64(63)
	_ = _hi1
	for _269_i5 := _dafny.IntOfUint32((_129_v9).Cardinality()); _269_i5.Cmp(_hi1) < 0; _269_i5 = _269_i5.Plus(_dafny.One) {
		var _270_v114 _dafny.Map
		_ = _270_v114
		_270_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_i5, _152_v29)
		var _271_v115 _dafny.Set
		_ = _271_v115
		_271_v115 = _dafny.SetOf(_270_v114)
		var _272_v116 D3
		_ = _272_v116
		_272_v116 = Companion_D3_.Create_DC8_(_271_v115)
		var _pat_let_tv2 = _270_v114
		_ = _pat_let_tv2
		var _pat_let_tv3 = _270_v114
		_ = _pat_let_tv3
		var _pat_let_tv4 = _271_v115
		_ = _pat_let_tv4
		var _source5 D3 = func(_pat_let4_0 D3) D3 {
			return func(_273_dt__update__tmp_h2 D3) D3 {
				return func(_pat_let5_0 _dafny.Set) D3 {
					return func(_274_dt__update_hcf12_h2 _dafny.Set) D3 {
						return Companion_D3_.Create_DC8_(_274_dt__update_hcf12_h2)
					}(_pat_let5_0)
				}((_dafny.SetOf(_pat_let_tv2, _pat_let_tv3)).Union(_pat_let_tv4))
			}(_pat_let4_0)
		}(_272_v116)
		_ = _source5
		if _source5.Is_DC9() {
			var _275___mcc_h12 bool = _source5.Get_().(D3_DC9).Cf13
			_ = _275___mcc_h12
			var _276___mcc_h13 _dafny.Map = _source5.Get_().(D3_DC9).Cf14
			_ = _276___mcc_h13
			var _277_cf14 _dafny.Map = _276___mcc_h13
			_ = _277_cf14
			var _278_cf13 bool = _275___mcc_h12
			_ = _278_cf13
			var _279_v117 _dafny.Map
			_ = _279_v117
			_279_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_151_v28, _151_v28), _152_v29)
			_279_v117 = (_279_v117).Update(_151_v28, (_dafny.Zero).Minus(_269_i5))
			var _280_v118 _dafny.Array
			_ = _280_v118
			var _nwElement0_9 bool = _278_cf13
			_ = _nwElement0_9
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(16))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_9, 0)
			(_nw36).ArraySet1(_133_v12, 1)
			(_nw36).ArraySet1(_133_v12, 2)
			(_nw36).ArraySet1(_278_cf13, 3)
			(_nw36).ArraySet1(_278_cf13, 4)
			(_nw36).ArraySet1(_278_cf13, 5)
			(_nw36).ArraySet1(_133_v12, 6)
			(_nw36).ArraySet1(_278_cf13, 7)
			(_nw36).ArraySet1(_278_cf13, 8)
			(_nw36).ArraySet1(_133_v12, 9)
			(_nw36).ArraySet1(_133_v12, 10)
			(_nw36).ArraySet1(_133_v12, 11)
			(_nw36).ArraySet1(_133_v12, 12)
			(_nw36).ArraySet1(_133_v12, 13)
			(_nw36).ArraySet1(_278_cf13, 14)
			(_nw36).ArraySet1(_278_cf13, 15)
			_280_v118 = _nw36
			var _281_v119 _dafny.Map
			_ = _281_v119
			_281_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_i5, _280_v118)
			(_137_globalState).F3 = !((func() _dafny.Map {
				if !(_278_cf13) {
					return _281_v119
				}
				return _281_v119
			})()).Equals(_281_v119)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_153_v30), 0))
			_ = _index31
			(_153_v30).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm1(_152_v29, _133_v12, _129_v9, _278_cf13, _137_globalState), (_dafny.Zero).Minus(_152_v29)), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_153_v30), 0))
			_ = _index32
			(_153_v30).ArraySet1(_120_v0, (_index32).Int())
			_151_v28 = Companion_Default___.Fm6(_dafny.IntOfUint32((_129_v9).Cardinality()), _137_globalState)
		} else {
			var _282___mcc_h14 _dafny.Set = _source5.Get_().(D3_DC8).Cf12
			_ = _282___mcc_h14
			var _283_cf12 _dafny.Set = _282___mcc_h14
			_ = _283_cf12
			var _284_v120 _dafny.Map
			_ = _284_v120
			_284_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_133_v12, _dafny.IntOfInt64(875))
			_284_v120 = (_284_v120).Update(true, (func() _dafny.Int {
				if _133_v12 {
					return _dafny.IntOfUint32((_dafny.SeqOf(_120_v0)).Cardinality())
				}
				return _152_v29
			})())
			(_137_globalState).F3 = _133_v12
			(_137_globalState).F9 = _152_v29
			var _285_v121 *C0
			_ = _285_v121
			var _nw37 *C0 = New_C0_()
			_ = _nw37
			_nw37.Ctor__()
			_285_v121 = _nw37
		}
		if false {
			var _286_v122 _dafny.Sequence
			_ = _286_v122
			var _287_v123 _dafny.Int
			_ = _287_v123
			var _288_v124 _dafny.Array
			_ = _288_v124
			var _289_v125 _dafny.CodePoint
			_ = _289_v125
			var _out28 _dafny.Sequence
			_ = _out28
			var _out29 _dafny.Int
			_ = _out29
			var _out30 _dafny.Array
			_ = _out30
			var _out31 _dafny.CodePoint
			_ = _out31
			_out28, _out29, _out30, _out31 = Companion_Default___.M0(!(_133_v12), _137_globalState)
			_286_v122 = _out28
			_287_v123 = _out29
			_288_v124 = _out30
			_289_v125 = _out31
			var _290_v126 D2
			_ = _290_v126
			_290_v126 = Companion_D2_.Create_DC4_(_270_v114)
			var _291_v127 D3
			_ = _291_v127
			_291_v127 = Companion_D3_.Create_DC9_(_133_v12, (_290_v126).Dtor_cf5())
			var _292_v128 _dafny.Sequence
			_ = _292_v128
			_292_v128 = _dafny.SeqOf((_270_v114).Update(_120_v0, _dafny.IntOfUint32((_129_v9).Cardinality())))
			var _293_v129 _dafny.Map
			_ = _293_v129
			_293_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v0, Companion_D3_.Create_DC9_((_291_v127).Dtor_cf13(), (_292_v128).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_120_v0, _133_v12, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(410))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}((func(_294_v125 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_295_i6 _dafny.Int) _dafny.CodePoint {
					return _294_v125
				}
			})(_289_v125))), _133_v12, _137_globalState), _dafny.IntOfUint32((_292_v128).Cardinality()))).Uint32()).(_dafny.Map)))
			var _296_v130 _dafny.Map
			_ = _296_v130
			_296_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_293_v129).Merge(_293_v129), _154_v31)
			_296_v130 = (_296_v130).Update(_293_v129, _138_v16)
			var _297_v131 _dafny.Array
			_ = _297_v131
			var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
			_ = _nw38
			_297_v131 = _nw38
			var _298_v132 _dafny.Sequence
			_ = _298_v132
			_298_v132 = _dafny.SeqOf(_123_v3, _123_v3, _dafny.SeqOf(_152_v29, (_dafny.Zero).Minus(_269_i5)))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_297_v131), 0))
			_ = _index33
			(_297_v131).ArraySet1(_dafny.MultiSetFromSeq(_298_v132), (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_297_v131), 0))
			_ = _index34
			(_297_v131).ArraySet1(_dafny.MultiSetOf(_dafny.SeqOf(_269_i5, _dafny.IntOfInt64(683), _dafny.IntOfUint32((_123_v3).Cardinality()))), (_index34).Int())
			(_137_globalState).F22 = (Companion_Default___.Fm1(_152_v29, _133_v12, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(807))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_299_v125 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_300_i7 _dafny.Int) _dafny.CodePoint {
					return _299_v125
				}
			})(_289_v125))), Companion_Default___.Fm0(_269_i5, _133_v12, _137_globalState), _137_globalState)).Plus(_dafny.IntOfInt64(228))
			var _301_v133 *C0
			_ = _301_v133
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__()
			_301_v133 = _nw39
		} else {
			var _302_v134 _dafny.Array
			_ = _302_v134
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw40
			_302_v134 = _nw40
			var _303_v135 _dafny.Array
			_ = _303_v135
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_5
			var _nw41 _dafny.Array
			_ = _nw41
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw41 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = func(_304_i8 _dafny.Int) bool {
					return true
				}
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw41 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw41).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw41).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_303_v135 = _nw41
			var _305_v136 _dafny.Array
			_ = _305_v136
			var _nwElement0_10 _dafny.Array = _303_v135
			_ = _nwElement0_10
			var _nw42 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(12))
			_ = _nw42
			(_nw42).ArraySet1(_nwElement0_10, 0)
			(_nw42).ArraySet1(_303_v135, 1)
			(_nw42).ArraySet1(_303_v135, 2)
			(_nw42).ArraySet1(_303_v135, 3)
			(_nw42).ArraySet1(_303_v135, 4)
			(_nw42).ArraySet1(_303_v135, 5)
			(_nw42).ArraySet1(_303_v135, 6)
			(_nw42).ArraySet1(_303_v135, 7)
			(_nw42).ArraySet1(_303_v135, 8)
			(_nw42).ArraySet1(_303_v135, 9)
			(_nw42).ArraySet1(_303_v135, 10)
			(_nw42).ArraySet1(_303_v135, 11)
			_305_v136 = _nw42
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_302_v134), 0))
			_ = _index35
			(_302_v134).ArraySet1(_305_v136, (_index35).Int())
			var _306_v137 _dafny.Map
			_ = _306_v137
			_306_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_125_v5, _125_v5)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_302_v134), 0))
			_ = _index36
			var _rhs39 bool = (func() bool {
				if !(true) {
					return (_133_v12) && (Companion_Default___.Fm0((_306_v137).Cardinality(), _133_v12, _137_globalState))
				}
				return _133_v12
			})()
			_ = _rhs39
			var _rhs40 _dafny.Array = _305_v136
			_ = _rhs40
			var _rhs41 _dafny.Int = _120_v0
			_ = _rhs41
			var _lhs30 _dafny.Array = _302_v134
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_302_v134), 0))
			_ = _lhs31
			var _lhs32 *GlobalState = _137_globalState
			_ = _lhs32
			_133_v12 = _rhs39
			(_lhs30).ArraySet1(_rhs40, (_lhs31).Int())
			_lhs32.F25 = _rhs41
			var _307_v138 *C0
			_ = _307_v138
			var _nw43 *C0 = New_C0_()
			_ = _nw43
			_nw43.Ctor__()
			_307_v138 = _nw43
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_125_v5), 0))
			_ = _index37
			(_125_v5).ArraySet1(_dafny.IntOfInt64(964), (_index37).Int())
			var _308_v139 _dafny.Map
			_ = _308_v139
			_308_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_307_v138, _307_v138)
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_125_v5), 0))
			_ = _index38
			var _rhs42 bool = Companion_Default___.Fm0(_152_v29, _133_v12, _137_globalState)
			_ = _rhs42
			var _rhs43 *C0 = (func() *C0 {
				if (_308_v139).Contains(_307_v138) {
					return (_308_v139).Get(_307_v138).(*C0)
				}
				return _307_v138
			})()
			_ = _rhs43
			var _rhs44 _dafny.Int = _152_v29
			_ = _rhs44
			var _rhs45 _dafny.Int = (_123_v3).Select((Companion_Default___.SafeIndex(_120_v0, _dafny.IntOfUint32((_123_v3).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs45
			var _rhs46 _dafny.Int = (_dafny.Zero).Minus(_120_v0)
			_ = _rhs46
			var _lhs33 *GlobalState = _137_globalState
			_ = _lhs33
			var _lhs34 *GlobalState = _137_globalState
			_ = _lhs34
			var _lhs35 *GlobalState = _137_globalState
			_ = _lhs35
			var _lhs36 _dafny.Array = _125_v5
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_125_v5), 0))
			_ = _lhs37
			_lhs33.F24 = _rhs42
			_307_v138 = _rhs43
			_lhs34.F22 = _rhs44
			_lhs35.F25 = _rhs45
			(_lhs36).ArraySet1(_rhs46, (_lhs37).Int())
			var _309_v140 D4
			_ = _309_v140
			_309_v140 = Companion_D4_.Create_DC10_(_307_v138)
			var _310_v141 _dafny.Map
			_ = _310_v141
			_310_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v140, (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))
			_310_v141 = (_310_v141).Update(_309_v140, (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))
			var _311_v142 *C0
			_ = _311_v142
			var _nw44 *C0 = New_C0_()
			_ = _nw44
			_nw44.Ctor__()
			_311_v142 = _nw44
		}
		var _312_v143 _dafny.Array
		_ = _312_v143
		var _nw45 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
		_ = _nw45
		_312_v143 = _nw45
		var _313_v144 _dafny.Map
		_ = _313_v144
		_313_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_133_v12), _152_v29)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_312_v143), 0))
		_ = _index39
		(_312_v143).ArraySet1(_313_v144, (_index39).Int())
		var _314_v145 D5
		_ = _314_v145
		_314_v145 = Companion_D5_.Create_DC12_(_313_v144)
		var _pat_let_tv5 = _313_v144
		_ = _pat_let_tv5
		var _pat_let_tv6 = _133_v12
		_ = _pat_let_tv6
		var _pat_let_tv7 = _313_v144
		_ = _pat_let_tv7
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_312_v143), 0))
		_ = _index40
		(_312_v143).ArraySet1((func(_pat_let6_0 D5) D5 {
			return func(_319_dt__update__tmp_h5 D5) D5 {
				return func(_pat_let11_0 _dafny.Map) D5 {
					return func(_320_dt__update_hcf18_h2 _dafny.Map) D5 {
						return Companion_D5_.Create_DC12_(_320_dt__update_hcf18_h2)
					}(_pat_let11_0)
				}(_pat_let_tv7)
			}(_pat_let6_0)
		}(func(_pat_let7_0 D5) D5 {
			return func(_317_dt__update__tmp_h4 D5) D5 {
				return func(_pat_let10_0 _dafny.Map) D5 {
					return func(_318_dt__update_hcf18_h1 _dafny.Map) D5 {
						return Companion_D5_.Create_DC12_(_318_dt__update_hcf18_h1)
					}(_pat_let10_0)
				}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv6, _269_i5))
			}(_pat_let7_0)
		}(func(_pat_let8_0 D5) D5 {
			return func(_315_dt__update__tmp_h3 D5) D5 {
				return func(_pat_let9_0 _dafny.Map) D5 {
					return func(_316_dt__update_hcf18_h0 _dafny.Map) D5 {
						return Companion_D5_.Create_DC12_(_316_dt__update_hcf18_h0)
					}(_pat_let9_0)
				}(_pat_let_tv5)
			}(_pat_let8_0)
		}(_314_v145)))).Dtor_cf18(), (_index40).Int())
		var _321_v146 _dafny.MultiSet
		_ = _321_v146
		_321_v146 = _dafny.MultiSetOf(_133_v12, _133_v12)
		var _322_v147 _dafny.Sequence
		_ = _322_v147
		_322_v147 = _dafny.SeqOf(_dafny.MultiSetOf(true), _321_v146, _321_v146, _321_v146)
		(_137_globalState).F14 = (func() bool {
			if (_128_v8).IsSubsetOf(_128_v8) {
				return _133_v12
			}
			return ((_322_v147).Select((Companion_Default___.SafeIndex(_152_v29, _dafny.IntOfUint32((_322_v147).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsDisjointFrom(_321_v146)
		})()
	}
	var _323_v148 *C0
	_ = _323_v148
	var _nw46 *C0 = New_C0_()
	_ = _nw46
	_nw46.Ctor__()
	_323_v148 = _nw46
	var _324_v149 _dafny.Sequence
	_ = _324_v149
	_324_v149 = _dafny.SeqOf(_323_v148, _323_v148)
	var _325_v150 _dafny.Map
	_ = _325_v150
	_325_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v0, _129_v9)
	_323_v148 = (_324_v149).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_152_v29, _133_v12, (func() _dafny.Sequence {
		if (_325_v150).Contains(_120_v0) {
			return (_325_v150).Get(_120_v0).(_dafny.Sequence)
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-209))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_326_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_327_i9 _dafny.Int) _dafny.CodePoint {
				return _326_v16
			}
		})(_138_v16)))
	})(), _133_v12, _137_globalState), _dafny.IntOfUint32((_324_v149).Cardinality()))).Uint32()).(*C0)
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))
	_ = _index41
	(_125_v5).ArraySet1((_dafny.Zero).Minus(_152_v29), (_index41).Int())
	var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))
	_ = _index42
	var _rhs47 _dafny.Int = _120_v0
	_ = _rhs47
	var _rhs48 _dafny.CodePoint = _154_v31
	_ = _rhs48
	var _rhs49 bool = (_120_v0).Cmp(_120_v0) <= 0
	_ = _rhs49
	var _rhs50 bool = _133_v12
	_ = _rhs50
	var _lhs38 _dafny.Array = _125_v5
	_ = _lhs38
	var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))
	_ = _lhs39
	var _lhs40 *GlobalState = _137_globalState
	_ = _lhs40
	var _lhs41 *GlobalState = _137_globalState
	_ = _lhs41
	(_lhs38).ArraySet1(_rhs47, (_lhs39).Int())
	_138_v16 = _rhs48
	_lhs40.F3 = _rhs49
	_lhs41.F24 = _rhs50
	var _328_v151 _dafny.Map
	_ = _328_v151
	_328_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _133_v12)
	var _329_v152 _dafny.MultiSet
	_ = _329_v152
	_329_v152 = _dafny.MultiSetOf((func() bool {
		if (_328_v151).Contains(_152_v29) {
			return (_328_v151).Get(_152_v29).(bool)
		}
		return !(_133_v12)
	})())
	if (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_151_v28, _151_v28))).IsProperSubsetOf(_329_v152) {
		(_137_globalState).F24 = true
		(_137_globalState).F22 = (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int)
		var _330_v153 _dafny.Map
		_ = _330_v153
		_330_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))
		var _331_v154 _dafny.Map
		_ = _331_v154
		_331_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v0, Companion_D3_.Create_DC9_(false, _330_v153))
		var _332_v155 _dafny.Array
		_ = _332_v155
		var _nw47 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw47
		_332_v155 = _nw47
		var _333_v156 D4
		_ = _333_v156
		_333_v156 = Companion_D4_.Create_DC11_(_331_v154, _332_v155)
		var _334_v157 _dafny.Map
		_ = _334_v157
		_334_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_133_v12), _333_v156)
		var _335_v158 _dafny.Map
		_ = _335_v158
		_335_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_120_v0, (_334_v157).Cardinality())
		var _source6 D2 = Companion_D2_.Create_DC4_((_335_v158).Merge(_330_v153))
		_ = _source6
		if _source6.Is_DC5() {
			var _336___mcc_h15 bool = _source6.Get_().(D2_DC5).Cf6
			_ = _336___mcc_h15
			var _337___mcc_h16 bool = _source6.Get_().(D2_DC5).Cf7
			_ = _337___mcc_h16
			var _338___mcc_h17 bool = _source6.Get_().(D2_DC5).Cf8
			_ = _338___mcc_h17
			var _339___mcc_h18 _dafny.Int = _source6.Get_().(D2_DC5).Cf9
			_ = _339___mcc_h18
			var _340___mcc_h19 bool = _source6.Get_().(D2_DC5).Cf10
			_ = _340___mcc_h19
			var _341_cf10 bool = _340___mcc_h19
			_ = _341_cf10
			var _342_cf9 _dafny.Int = _339___mcc_h18
			_ = _342_cf9
			var _343_cf8 bool = _338___mcc_h17
			_ = _343_cf8
			var _344_cf7 bool = _337___mcc_h16
			_ = _344_cf7
			var _345_cf6 bool = _336___mcc_h15
			_ = _345_cf6
			var _346_v159 _dafny.Sequence
			_ = _346_v159
			var _347_v160 _dafny.Int
			_ = _347_v160
			var _348_v161 _dafny.Array
			_ = _348_v161
			var _349_v162 _dafny.CodePoint
			_ = _349_v162
			var _out32 _dafny.Sequence
			_ = _out32
			var _out33 _dafny.Int
			_ = _out33
			var _out34 _dafny.Array
			_ = _out34
			var _out35 _dafny.CodePoint
			_ = _out35
			_out32, _out33, _out34, _out35 = Companion_Default___.M0(Companion_Default___.Fm0((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _345_cf6, _137_globalState), _137_globalState)
			_346_v159 = _out32
			_347_v160 = _out33
			_348_v161 = _out34
			_349_v162 = _out35
			var _350_v163 D6
			_ = _350_v163
			_350_v163 = Companion_D6_.Create_DC16_(_121_v1)
			var _351_v164 _dafny.Map
			_ = _351_v164
			_351_v164 = _134_v13
			var _pat_let_tv8 = _121_v1
			_ = _pat_let_tv8
			var _rhs51 _dafny.Set = (func(_pat_let12_0 D6) D6 {
				return func(_352_dt__update__tmp_h6 D6) D6 {
					return func(_pat_let13_0 _dafny.Set) D6 {
						return func(_353_dt__update_hcf22_h0 _dafny.Set) D6 {
							return Companion_D6_.Create_DC16_(_353_dt__update_hcf22_h0)
						}(_pat_let13_0)
					}(_pat_let_tv8)
				}(_pat_let12_0)
			}(_350_v163)).Dtor_cf22()
			_ = _rhs51
			var _rhs52 bool = _133_v12
			_ = _rhs52
			var _rhs53 _dafny.Sequence = _324_v149
			_ = _rhs53
			var _rhs54 _dafny.Int = (_351_v164).Cardinality()
			_ = _rhs54
			var _lhs42 *GlobalState = _137_globalState
			_ = _lhs42
			var _lhs43 *GlobalState = _137_globalState
			_ = _lhs43
			_121_v1 = _rhs51
			_lhs42.F24 = _rhs52
			_324_v149 = _rhs53
			_lhs43.F9 = _rhs54
			(_137_globalState).F14 = _133_v12
			var _354_v165 *C0
			_ = _354_v165
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__()
			_354_v165 = _nw48
		} else if _source6.Is_DC6() {
			var _355_v166 _dafny.Sequence
			_ = _355_v166
			var _356_v167 _dafny.Int
			_ = _356_v167
			var _357_v168 _dafny.Array
			_ = _357_v168
			var _358_v169 _dafny.CodePoint
			_ = _358_v169
			var _out36 _dafny.Sequence
			_ = _out36
			var _out37 _dafny.Int
			_ = _out37
			var _out38 _dafny.Array
			_ = _out38
			var _out39 _dafny.CodePoint
			_ = _out39
			_out36, _out37, _out38, _out39 = Companion_Default___.M0(_133_v12, _137_globalState)
			_355_v166 = _out36
			_356_v167 = _out37
			_357_v168 = _out38
			_358_v169 = _out39
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_332_v155), 0))
			_ = _index43
			(_332_v155).ArraySet1(_133_v12, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(756), _dafny.ArrayLen((_332_v155), 0))
			_ = _index44
			(_332_v155).ArraySet1((_152_v29).Cmp((_121_v1).Cardinality()) <= 0, (_index44).Int())
			_151_v28 = _355_v166
			(_137_globalState).F9 = _dafny.IntOfInt64(102)
		} else if _source6.Is_DC4() {
			var _359___mcc_h20 _dafny.Map = _source6.Get_().(D2_DC4).Cf5
			_ = _359___mcc_h20
			var _360_cf5 _dafny.Map = _359___mcc_h20
			_ = _360_cf5
			(_137_globalState).F21 = _129_v9
			(_137_globalState).F24 = (Companion_Default___.SafeDivisionInt(_152_v29, (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))).Cmp(_dafny.IntOfUint32((_151_v28).Cardinality())) > 0
			var _361_v170 _dafny.Array
			_ = _361_v170
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
			_ = _nw49
			_361_v170 = _nw49
			var _362_v171 _dafny.Map
			_ = _362_v171
			_362_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_323_v148, _258_v104)
			var _363_v172 D8
			_ = _363_v172
			_363_v172 = Companion_D8_.Create_DC19_((func() _dafny.Map {
				if (_362_v171).Contains(_323_v148) {
					return (_362_v171).Get(_323_v148).(_dafny.Map)
				}
				return _258_v104
			})())
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_361_v170), 0))
			_ = _index45
			(_361_v170).ArraySet1(((_363_v172).Dtor_cf24()).Merge(_258_v104), (_index45).Int())
			var _364_v173 _dafny.Sequence
			_ = _364_v173
			_364_v173 = _dafny.SeqOf((_258_v104).Update((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _154_v31), _258_v104, _258_v104)
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_361_v170), 0))
			_ = _index46
			(_361_v170).ArraySet1((_258_v104).Merge((_364_v173).Select((Companion_Default___.SafeIndex((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_364_v173).Cardinality()))).Uint32()).(_dafny.Map)), (_index46).Int())
			var _365_v174 _dafny.Map
			_ = _365_v174
			_365_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v28, _dafny.Companion_Sequence_.Concatenate(_123_v3, _123_v3))
			_365_v174 = (_365_v174).Update(_dafny.Companion_Sequence_.Concatenate(_151_v28, _151_v28), _dafny.Companion_Sequence_.Concatenate(_123_v3, _123_v3))
		} else {
			var _366___mcc_h21 D2 = _source6.Get_().(D2_DC7).Cf11
			_ = _366___mcc_h21
			var _367_cf11 D2 = _366___mcc_h21
			_ = _367_cf11
			var _368_v175 _dafny.Sequence
			_ = _368_v175
			_368_v175 = _dafny.SeqOf(Companion_Default___.Fm7(_133_v12, _137_globalState))
			(_137_globalState).F14 = _dafny.Companion_Sequence_.IsPrefixOf(_368_v175, _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm8(_137_globalState), _368_v175))
			var _369_v176 _dafny.Sequence
			_ = _369_v176
			var _370_v177 _dafny.Int
			_ = _370_v177
			var _371_v178 _dafny.Array
			_ = _371_v178
			var _372_v179 _dafny.CodePoint
			_ = _372_v179
			var _out40 _dafny.Sequence
			_ = _out40
			var _out41 _dafny.Int
			_ = _out41
			var _out42 _dafny.Array
			_ = _out42
			var _out43 _dafny.CodePoint
			_ = _out43
			_out40, _out41, _out42, _out43 = Companion_Default___.M0(_133_v12, _137_globalState)
			_369_v176 = _out40
			_370_v177 = _out41
			_371_v178 = _out42
			_372_v179 = _out43
			_154_v31 = Companion_Default___.Fm4(_372_v179, _151_v28, _137_globalState)
			var _rhs55 _dafny.Int = (_dafny.IntOfInt64(-121)).Times((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))
			_ = _rhs55
			var _rhs56 bool = _133_v12
			_ = _rhs56
			var _lhs44 *GlobalState = _137_globalState
			_ = _lhs44
			var _lhs45 *GlobalState = _137_globalState
			_ = _lhs45
			_lhs44.F25 = _rhs55
			_lhs45.F24 = _rhs56
		}
		(_137_globalState).F14 = !(_133_v12)
		if (_dafny.MultiSetFromSeq(_151_v28)).IsSubsetOf(_329_v152) {
			_154_v31 = _154_v31
			_154_v31 = (func() _dafny.CodePoint {
				if false {
					return _138_v16
				}
				return _154_v31
			})()
			_125_v5 = _153_v30
			var _373_v180 *C0
			_ = _373_v180
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__()
			_373_v180 = _nw50
			var _374_v181 _dafny.Set
			_ = _374_v181
			_374_v181 = _dafny.SetOf(true, _dafny.Companion_Sequence_.Contains(_123_v3, _152_v29), _133_v12, (_120_v0).Cmp(_dafny.IntOfUint32((_151_v28).Cardinality())) != 0, (_133_v12) == (_133_v12))
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))
			_ = _index47
			var _rhs57 _dafny.Int = (func() _dafny.Int {
				if (_335_v158).Contains(_152_v29) {
					return (_335_v158).Get(_152_v29).(_dafny.Int)
				}
				return (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int)
			})()
			_ = _rhs57
			var _rhs58 _dafny.Array = _125_v5
			_ = _rhs58
			var _rhs59 _dafny.Set = ((_374_v181).Union(_374_v181)).Union(_374_v181)
			_ = _rhs59
			var _rhs60 _dafny.Int = _120_v0
			_ = _rhs60
			var _rhs61 bool = _133_v12
			_ = _rhs61
			var _lhs46 _dafny.Array = _125_v5
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))
			_ = _lhs47
			var _lhs48 *GlobalState = _137_globalState
			_ = _lhs48
			var _lhs49 *GlobalState = _137_globalState
			_ = _lhs49
			(_lhs46).ArraySet1(_rhs57, (_lhs47).Int())
			_125_v5 = _rhs58
			_374_v181 = _rhs59
			_lhs48.F25 = _rhs60
			_lhs49.F24 = _rhs61
		} else {
			var _375_v182 *C0
			_ = _375_v182
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__()
			_375_v182 = _nw51
			(_137_globalState).F25 = (_dafny.Zero).Minus(_152_v29)
			(_137_globalState).F25 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_133_v12, _133_v12), _151_v28)).Cardinality())
			(_137_globalState).F25 = _152_v29
			var _376_v183 *C0
			_ = _376_v183
			var _nw52 *C0 = New_C0_()
			_ = _nw52
			_nw52.Ctor__()
			_376_v183 = _nw52
		}
	} else {
		var _377_v185 _dafny.Map
		_ = _377_v185
		_377_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_152_v29, (func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter11 := _dafny.Iterate((_123_v3).Elements()); ; {
				_compr_10, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _378_v184 _dafny.Int
				_378_v184 = interface{}(_compr_10).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_123_v3, _378_v184) {
					_coll10.Add((_378_v184).Plus(_120_v0), (_128_v8).Cardinality())
				}
			}
			return _coll10.ToMap()
		}()).Cardinality())
		var _rhs62 _dafny.Int = (_dafny.Zero).Minus((((_377_v185).Cardinality()).Plus((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))).Minus(Companion_Default___.SafeDivisionInt(_152_v29, _152_v29)))
		_ = _rhs62
		var _rhs63 bool = _133_v12
		_ = _rhs63
		var _rhs64 _dafny.Set = _121_v1
		_ = _rhs64
		var _lhs50 *GlobalState = _137_globalState
		_ = _lhs50
		var _lhs51 *GlobalState = _137_globalState
		_ = _lhs51
		_lhs50.F25 = _rhs62
		_lhs51.F3 = _rhs63
		_121_v1 = _rhs64
		var _379_v186 _dafny.Sequence
		_ = _379_v186
		var _380_v187 _dafny.Int
		_ = _380_v187
		var _381_v188 _dafny.Array
		_ = _381_v188
		var _382_v189 _dafny.CodePoint
		_ = _382_v189
		var _out44 _dafny.Sequence
		_ = _out44
		var _out45 _dafny.Int
		_ = _out45
		var _out46 _dafny.Array
		_ = _out46
		var _out47 _dafny.CodePoint
		_ = _out47
		_out44, _out45, _out46, _out47 = Companion_Default___.M0(_133_v12, _137_globalState)
		_379_v186 = _out44
		_380_v187 = _out45
		_381_v188 = _out46
		_382_v189 = _out47
		var _383_v190 _dafny.Sequence
		_ = _383_v190
		var _384_v191 _dafny.Int
		_ = _384_v191
		var _385_v192 _dafny.Array
		_ = _385_v192
		var _386_v193 _dafny.CodePoint
		_ = _386_v193
		var _out48 _dafny.Sequence
		_ = _out48
		var _out49 _dafny.Int
		_ = _out49
		var _out50 _dafny.Array
		_ = _out50
		var _out51 _dafny.CodePoint
		_ = _out51
		_out48, _out49, _out50, _out51 = Companion_Default___.M0(!(_133_v12) || (_133_v12), _137_globalState)
		_383_v190 = _out48
		_384_v191 = _out49
		_385_v192 = _out50
		_386_v193 = _out51
		var _387_v194 _dafny.Array
		_ = _387_v194
		var _nw53 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(10))
		_ = _nw53
		_387_v194 = _nw53
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_387_v194), 0))
		_ = _index48
		(_387_v194).ArraySet1(_323_v148, (_index48).Int())
		var _388_v195 _dafny.Array
		_ = _388_v195
		var _nwElement0_11 bool = _133_v12
		_ = _nwElement0_11
		var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(14))
		_ = _nw54
		(_nw54).ArraySet1(_nwElement0_11, 0)
		(_nw54).ArraySet1(_133_v12, 1)
		(_nw54).ArraySet1((func() bool {
			if _133_v12 {
				return _133_v12
			}
			return _133_v12
		})(), 2)
		(_nw54).ArraySet1((!(_133_v12)) && (_133_v12), 3)
		(_nw54).ArraySet1(true, 4)
		(_nw54).ArraySet1(!(_133_v12), 5)
		(_nw54).ArraySet1(_133_v12, 6)
		(_nw54).ArraySet1(_133_v12, 7)
		(_nw54).ArraySet1(_133_v12, 8)
		(_nw54).ArraySet1(_133_v12, 9)
		(_nw54).ArraySet1(_133_v12, 10)
		(_nw54).ArraySet1((func() bool {
			if _133_v12 {
				return _133_v12
			}
			return _133_v12
		})(), 11)
		(_nw54).ArraySet1((_152_v29).Cmp(_152_v29) != 0, 12)
		(_nw54).ArraySet1(_133_v12, 13)
		_388_v195 = _nw54
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_388_v195), 0))
		_ = _index49
		(_388_v195).ArraySet1(_133_v12, (_index49).Int())
		var _389_v196 _dafny.Set
		_ = _389_v196
		_389_v196 = _dafny.SetOf(_129_v9, _129_v9, _dafny.UnicodeSeqOfUtf8Bytes("bxdpjq"), Companion_Default___.Fm2((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _152_v29, false, (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _137_globalState), _129_v9)
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_387_v194), 0))
		_ = _index50
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_388_v195), 0))
		_ = _index51
		var _rhs65 *C0 = _323_v148
		_ = _rhs65
		var _rhs66 _dafny.Array = _131_v11
		_ = _rhs66
		var _rhs67 _dafny.Sequence = _123_v3
		_ = _rhs67
		var _rhs68 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_383_v190, _383_v190), !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("b"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(781))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}((func(_390_v193 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_391_i10 _dafny.Int) _dafny.CodePoint {
				return _390_v193
			}
		})(_386_v193)))))
		_ = _rhs68
		var _rhs69 bool = (_389_v196).IsProperSubsetOf(_389_v196)
		_ = _rhs69
		var _lhs52 _dafny.Array = _387_v194
		_ = _lhs52
		var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_387_v194), 0))
		_ = _lhs53
		var _lhs54 *GlobalState = _137_globalState
		_ = _lhs54
		var _lhs55 _dafny.Array = _388_v195
		_ = _lhs55
		var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_388_v195), 0))
		_ = _lhs56
		(_lhs52).ArraySet1(_rhs65, (_lhs53).Int())
		_131_v11 = _rhs66
		_123_v3 = _rhs67
		_lhs54.F3 = _rhs68
		(_lhs55).ArraySet1(_rhs69, (_lhs56).Int())
		if false {
			(_137_globalState).F9 = (_dafny.Zero).Minus(_380_v187)
			(_137_globalState).F25 = (_323_v148).Fm5(_137_globalState)
			var _392_v197 _dafny.Sequence
			_ = _392_v197
			var _393_v198 _dafny.Int
			_ = _393_v198
			var _394_v199 _dafny.Array
			_ = _394_v199
			var _395_v200 _dafny.CodePoint
			_ = _395_v200
			var _out52 _dafny.Sequence
			_ = _out52
			var _out53 _dafny.Int
			_ = _out53
			var _out54 _dafny.Array
			_ = _out54
			var _out55 _dafny.CodePoint
			_ = _out55
			_out52, _out53, _out54, _out55 = Companion_Default___.M0(_133_v12, _137_globalState)
			_392_v197 = _out52
			_393_v198 = _out53
			_394_v199 = _out54
			_395_v200 = _out55
			var _396_v201 *C0
			_ = _396_v201
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__()
			_396_v201 = _nw55
			var _397_v202 _dafny.Map
			_ = _397_v202
			_397_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_388_v195).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_388_v195), 0))).Int()).(bool), _120_v0)
			_397_v202 = _397_v202
		} else {
			(_137_globalState).F25 = _dafny.IntOfUint32((_383_v190).Cardinality())
			(_137_globalState).F24 = ((_384_v191).Plus((_dafny.Zero).Minus(_152_v29))).Cmp(_dafny.IntOfUint32((_123_v3).Cardinality())) >= 0
			var _398_v203 *C0
			_ = _398_v203
			var _nw56 *C0 = New_C0_()
			_ = _nw56
			_nw56.Ctor__()
			_398_v203 = _nw56
			(_137_globalState).F14 = (_388_v195).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_388_v195), 0))).Int()).(bool)
			(_137_globalState).F24 = (_388_v195).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_388_v195), 0))).Int()).(bool)
		}
	}
	var _399_v204 _dafny.Set
	_ = _399_v204
	_399_v204 = _dafny.SetOf((_151_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_129_v9).Cardinality()), _dafny.IntOfUint32((_151_v28).Cardinality()))).Uint32()).(bool), _133_v12, _133_v12)
	if ((_399_v204).Difference(_399_v204)).IsDisjointFrom(_399_v204) {
		var _400_v205 _dafny.Array
		_ = _400_v205
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
		_ = _nw57
		_400_v205 = _nw57
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_400_v205), 0))
		_ = _index52
		(_400_v205).ArraySet1(_133_v12, (_index52).Int())
		var _401_v206 _dafny.Map
		_ = _401_v206
		_401_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_328_v151).Cardinality(), _dafny.IntOfInt64(-212))
		var _402_v207 D2
		_ = _402_v207
		_402_v207 = Companion_D2_.Create_DC4_(_401_v206)
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_400_v205), 0))
		_ = _index53
		(_400_v205).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int), _133_v12, _152_v29, _137_globalState), _129_v9)).Cardinality()), !((_402_v207).Dtor_cf5()).Contains((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int)), _137_globalState), (_index53).Int())
		(_137_globalState).F22 = (_dafny.Zero).Minus((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int))
		var _403_v208 _dafny.Sequence
		_ = _403_v208
		var _404_v209 _dafny.Int
		_ = _404_v209
		var _405_v210 _dafny.Array
		_ = _405_v210
		var _406_v211 _dafny.CodePoint
		_ = _406_v211
		var _out56 _dafny.Sequence
		_ = _out56
		var _out57 _dafny.Int
		_ = _out57
		var _out58 _dafny.Array
		_ = _out58
		var _out59 _dafny.CodePoint
		_ = _out59
		_out56, _out57, _out58, _out59 = Companion_Default___.M0((_400_v205).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_400_v205), 0))).Int()).(bool), _137_globalState)
		_403_v208 = _out56
		_404_v209 = _out57
		_405_v210 = _out58
		_406_v211 = _out59
		_153_v30 = _405_v210
		(_137_globalState).F24 = (_400_v205).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_400_v205), 0))).Int()).(bool)
	} else {
		if !((_dafny.IntOfUint32((_151_v28).Cardinality())).Cmp((_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int)) != 0) || ((_151_v28).Select((Companion_Default___.SafeIndex(_152_v29, _dafny.IntOfUint32((_151_v28).Cardinality()))).Uint32()).(bool)) {
			var _407_v212 _dafny.Array
			_ = _407_v212
			var _nwElement0_12 _dafny.Array = _127_v7
			_ = _nwElement0_12
			var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(19))
			_ = _nw58
			(_nw58).ArraySet1(_nwElement0_12, 0)
			(_nw58).ArraySet1(_127_v7, 1)
			(_nw58).ArraySet1(_127_v7, 2)
			(_nw58).ArraySet1(_127_v7, 3)
			(_nw58).ArraySet1((func() _dafny.Array {
				if _133_v12 {
					return _127_v7
				}
				return _127_v7
			})(), 4)
			(_nw58).ArraySet1(_127_v7, 5)
			(_nw58).ArraySet1(_127_v7, 6)
			(_nw58).ArraySet1(_127_v7, 7)
			(_nw58).ArraySet1(_127_v7, 8)
			(_nw58).ArraySet1(_127_v7, 9)
			(_nw58).ArraySet1(_127_v7, 10)
			(_nw58).ArraySet1(_127_v7, 11)
			(_nw58).ArraySet1(_127_v7, 12)
			(_nw58).ArraySet1((func() _dafny.Array {
				if _133_v12 {
					return _127_v7
				}
				return _127_v7
			})(), 13)
			(_nw58).ArraySet1(_127_v7, 14)
			(_nw58).ArraySet1(_127_v7, 15)
			(_nw58).ArraySet1(_127_v7, 16)
			(_nw58).ArraySet1(_127_v7, 17)
			(_nw58).ArraySet1(_127_v7, 18)
			_407_v212 = _nw58
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_407_v212), 0))
			_ = _index54
			(_407_v212).ArraySet1(_127_v7, (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_407_v212), 0))
			_ = _index55
			(_407_v212).ArraySet1(_127_v7, (_index55).Int())
			(_137_globalState).F24 = (_133_v12) && (!(_133_v12) || (_133_v12))
			var _408_v213 D6
			_ = _408_v213
			_408_v213 = Companion_D6_.Create_DC17_()
			var _rhs70 _dafny.Sequence = _129_v9
			_ = _rhs70
			var _rhs71 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("smrlm")
			_ = _rhs71
			var _rhs72 D6 = _408_v213
			_ = _rhs72
			var _lhs57 *GlobalState = _137_globalState
			_ = _lhs57
			var _lhs58 *GlobalState = _137_globalState
			_ = _lhs58
			_lhs57.F21 = _rhs70
			_lhs58.F21 = _rhs71
			_408_v213 = _rhs72
			_153_v30 = _153_v30
			(_137_globalState).F24 = Companion_Default___.Fm0(_152_v29, Companion_Default___.Fm0(_120_v0, _133_v12, _137_globalState), _137_globalState)
		} else {
			var _409_v214 _dafny.Sequence
			_ = _409_v214
			var _410_v215 _dafny.Int
			_ = _410_v215
			var _411_v216 _dafny.Array
			_ = _411_v216
			var _412_v217 _dafny.CodePoint
			_ = _412_v217
			var _out60 _dafny.Sequence
			_ = _out60
			var _out61 _dafny.Int
			_ = _out61
			var _out62 _dafny.Array
			_ = _out62
			var _out63 _dafny.CodePoint
			_ = _out63
			_out60, _out61, _out62, _out63 = Companion_Default___.M0(_133_v12, _137_globalState)
			_409_v214 = _out60
			_410_v215 = _out61
			_411_v216 = _out62
			_412_v217 = _out63
			(_137_globalState).F14 = _133_v12
			var _413_v218 _dafny.Map
			_ = _413_v218
			_413_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v216, false)
			_413_v218 = (_413_v218).Merge(_413_v218)
			var _414_v219 _dafny.Sequence
			_ = _414_v219
			var _415_v220 _dafny.Int
			_ = _415_v220
			var _416_v221 _dafny.Array
			_ = _416_v221
			var _417_v222 _dafny.CodePoint
			_ = _417_v222
			var _out64 _dafny.Sequence
			_ = _out64
			var _out65 _dafny.Int
			_ = _out65
			var _out66 _dafny.Array
			_ = _out66
			var _out67 _dafny.CodePoint
			_ = _out67
			_out64, _out65, _out66, _out67 = Companion_Default___.M0(_133_v12, _137_globalState)
			_414_v219 = _out64
			_415_v220 = _out65
			_416_v221 = _out66
			_417_v222 = _out67
			(_137_globalState).F21 = _129_v9
		}
		var _418_v223 _dafny.Sequence
		_ = _418_v223
		var _419_v224 _dafny.Int
		_ = _419_v224
		var _420_v225 _dafny.Array
		_ = _420_v225
		var _421_v226 _dafny.CodePoint
		_ = _421_v226
		var _out68 _dafny.Sequence
		_ = _out68
		var _out69 _dafny.Int
		_ = _out69
		var _out70 _dafny.Array
		_ = _out70
		var _out71 _dafny.CodePoint
		_ = _out71
		_out68, _out69, _out70, _out71 = Companion_Default___.M0((_399_v204).IsDisjointFrom(_399_v204), _137_globalState)
		_418_v223 = _out68
		_419_v224 = _out69
		_420_v225 = _out70
		_421_v226 = _out71
		var _422_v227 _dafny.Sequence
		_ = _422_v227
		var _423_v228 _dafny.Int
		_ = _423_v228
		var _424_v229 _dafny.Array
		_ = _424_v229
		var _425_v230 _dafny.CodePoint
		_ = _425_v230
		var _out72 _dafny.Sequence
		_ = _out72
		var _out73 _dafny.Int
		_ = _out73
		var _out74 _dafny.Array
		_ = _out74
		var _out75 _dafny.CodePoint
		_ = _out75
		_out72, _out73, _out74, _out75 = Companion_Default___.M0(_133_v12, _137_globalState)
		_422_v227 = _out72
		_423_v228 = _out73
		_424_v229 = _out74
		_425_v230 = _out75
		(_137_globalState).F24 = _133_v12
		(_137_globalState).F14 = _133_v12
	}
	(_137_globalState).F9 = (_125_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(998), _dafny.ArrayLen((_125_v5), 0))).Int()).(_dafny.Int)
	_120_v0 = (_323_v148).Fm5(_137_globalState)
	_dafny.Print(_120_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v1).Equals(_dafny.SetOf(_dafny.IntOfInt64(691))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(691)), _dafny.IntOfInt64(691))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_123_v3, _dafny.SeqOf(_dafny.IntOfInt64(691), _dafny.IntOfInt64(691))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.SeqOf(_dafny.IntOfInt64(691), _dafny.IntOfInt64(691)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v5).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_126_v6).Dtor_cf0()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_127_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_128_v8).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_129_v9.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_131_v11).ArrayGet1CodePoint((_dafny.IntOfInt64(22)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_133_v12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v13).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v14).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v15).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(691), _dafny.IntOfInt64(691)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState.F0).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(691)), _dafny.IntOfInt64(691))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F1()).Equals(_dafny.SetOf(_dafny.IntOfInt64(691))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F6()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.SeqOf(_dafny.IntOfInt64(691), _dafny.IntOfInt64(691)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F10()).Equals(_dafny.SetOf(_dafny.One, _dafny.IntOfInt64(691), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F12())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(19)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(20)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(21)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F17()).ArrayGet1CodePoint((_dafny.IntOfInt64(22)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F19()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F20()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_137_globalState).F20()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F21.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_globalState.F23).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(691), _dafny.IntOfInt64(691)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_137_globalState.F25)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_138_v16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(691)), _dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_140_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_151_v28, _dafny.SeqOf(false, true, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_152_v29)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v30).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_154_v31)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v104).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-119716), _dafny.CodePoint('f'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_259_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_324_v149).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_325_v150).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(691), _dafny.UnicodeSeqOfUtf8Bytes("px"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_328_v151).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(691), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_329_v152).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_399_v204).Equals(_dafny.SetOf(false, true)))
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 _dafny.Int
	Cf3 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 _dafny.Int, Cf3 _dafny.Int) D0 {
	return D0{D0_DC2{Cf2, Cf3}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Array
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Array) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf0() _dafny.Array {
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
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
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
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
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
	Cf4 _dafny.CodePoint
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf4 _dafny.CodePoint) D1 {
	return D1{D1_DC3{Cf4}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D1) Dtor_cf4() _dafny.CodePoint {
	return _this.Get_().(D1_DC3).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf4 == data2.Cf4
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
	Cf6  bool
	Cf7  bool
	Cf8  bool
	Cf9  _dafny.Int
	Cf10 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf6 bool, Cf7 bool, Cf8 bool, Cf9 _dafny.Int, Cf10 bool) D2 {
	return D2{D2_DC5{Cf6, Cf7, Cf8, Cf9, Cf10}}
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

type D2_DC4 struct {
	Cf5 _dafny.Map
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 _dafny.Map) D2 {
	return D2{D2_DC4{Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC7 struct {
	Cf11 D2
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 D2) D2 {
	return D2{D2_DC7{Cf11}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false, false, false, _dafny.Zero, false)
}

func (_this D2) Dtor_cf6() bool {
	return _this.Get_().(D2_DC5).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC5).Cf7
}

func (_this D2) Dtor_cf8() bool {
	return _this.Get_().(D2_DC5).Cf8
}

func (_this D2) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC5).Cf10
}

func (_this D2) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf11() D2 {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7 && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10
		}
	case D2_DC6:
		{
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5.Equals(data2.Cf5)
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf13 bool
	Cf14 _dafny.Map
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf13 bool, Cf14 _dafny.Map) D3 {
	return D3{D3_DC9{Cf13, Cf14}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf12 _dafny.Set
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 _dafny.Set) D3 {
	return D3{D3_DC8{Cf12}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(false, _dafny.EmptyMap)
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC9).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Map {
	return _this.Get_().(D3_DC9).Cf14
}

func (_this D3) Dtor_cf12() _dafny.Set {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Equals(data2.Cf14)
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf12.Equals(data2.Cf12)
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
	Cf16 _dafny.Map
	Cf17 _dafny.Array
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf16 _dafny.Map, Cf17 _dafny.Array) D4 {
	return D4{D4_DC11{Cf16, Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC10 struct {
	Cf15 *C0
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf15 *C0) D4 {
	return D4{D4_DC10{Cf15}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC11_(_dafny.EmptyMap, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D4) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D4_DC11).Cf16
}

func (_this D4) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) Dtor_cf15() *C0 {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
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
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf16.Equals(data2.Cf16) && data1.Cf17 == data2.Cf17
		}
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
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
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_() D5 {
	return D5{D5_DC13{}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf19 _dafny.Map
	Cf20 _dafny.Int
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf19 _dafny.Map, Cf20 _dafny.Int) D5 {
	return D5{D5_DC14{Cf19, Cf20}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC12 struct {
	Cf18 _dafny.Map
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 _dafny.Map) D5 {
	return D5{D5_DC12{Cf18}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC15 struct {
	Cf21 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf21 D5) D5 {
	return D5{D5_DC15{Cf21}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC13_()
}

func (_this D5) Dtor_cf19() _dafny.Map {
	return _this.Get_().(D5_DC14).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf20
}

func (_this D5) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) Dtor_cf21() D5 {
	return _this.Get_().(D5_DC15).Cf21
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC13:
		{
			return "D5.DC13"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf21) + ")"
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
			_, ok := other.Get_().(D5_DC13)
			return ok
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf21.Equals(data2.Cf21)
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

type D6_DC17 struct {
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_() D6 {
	return D6{D6_DC17{}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf22 _dafny.Set
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf22 _dafny.Set) D6 {
	return D6{D6_DC16{Cf22}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_()
}

func (_this D6) Dtor_cf22() _dafny.Set {
	return _this.Get_().(D6_DC16).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf22) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC17:
		{
			_, ok := other.Get_().(D6_DC17)
			return ok
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf22.Equals(data2.Cf22)
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

// Definition of datatype D7
type D7 struct {
	Data_D7_
}

func (_this D7) Get_() Data_D7_ {
	return _this.Data_D7_
}

type Data_D7_ interface {
	isD7()
}

type CompanionStruct_D7_ struct {
}

var Companion_D7_ = CompanionStruct_D7_{}

type D7_DC18 struct {
	Cf23 _dafny.Map
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf23 _dafny.Map) D7 {
	return D7{D7_DC18{Cf23}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D7) Dtor_cf23() _dafny.Map {
	return _this.Get_().(D7_DC18).Cf23
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf23.Equals(data2.Cf23)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D7) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D7)
	return ok && _this.Equals(typed)
}

func Type_D7_() _dafny.TypeDescriptor {
	return type_D7_{}
}

type type_D7_ struct {
}

func (_this type_D7_) Default() interface{} {
	return Companion_D7_.Default()
}

func (_this type_D7_) String() string {
	return "main.D7"
}
func (_this D7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D7{}

// End of datatype D7

// Definition of datatype D8
type D8 struct {
	Data_D8_
}

func (_this D8) Get_() Data_D8_ {
	return _this.Data_D8_
}

type Data_D8_ interface {
	isD8()
}

type CompanionStruct_D8_ struct {
}

var Companion_D8_ = CompanionStruct_D8_{}

type D8_DC20 struct {
	Cf25 bool
	Cf26 _dafny.Int
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf25 bool, Cf26 _dafny.Int) D8 {
	return D8{D8_DC20{Cf25, Cf26}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC19 struct {
	Cf24 _dafny.Map
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf24 _dafny.Map) D8 {
	return D8{D8_DC19{Cf24}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC21 struct {
	Cf27 D8
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf27 D8) D8 {
	return D8{D8_DC21{Cf27}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC20_(false, _dafny.Zero)
}

func (_this D8) Dtor_cf25() bool {
	return _this.Get_().(D8_DC20).Cf25
}

func (_this D8) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D8_DC20).Cf26
}

func (_this D8) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D8_DC19).Cf24
}

func (_this D8) Dtor_cf27() D8 {
	return _this.Get_().(D8_DC21).Cf27
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26.Cmp(data2.Cf26) == 0
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf24.Equals(data2.Cf24)
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf27.Equals(data2.Cf27)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D8) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D8)
	return ok && _this.Equals(typed)
}

func Type_D8_() _dafny.TypeDescriptor {
	return type_D8_{}
}

type type_D8_ struct {
}

func (_this type_D8_) Default() interface{} {
	return Companion_D8_.Default()
}

func (_this type_D8_) String() string {
	return "main.D8"
}
func (_this D8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D8{}

// End of datatype D8

// Definition of class GlobalState
type GlobalState struct {
	F0   _dafny.Map
	F3   bool
	F9   _dafny.Int
	F14  bool
	F21  _dafny.Sequence
	F22  _dafny.Int
	F23  _dafny.Set
	F24  bool
	F25  _dafny.Int
	_f1  _dafny.Set
	_f2  _dafny.Int
	_f4  _dafny.Int
	_f5  _dafny.CodePoint
	_f6  _dafny.Map
	_f7  _dafny.Int
	_f8  _dafny.Int
	_f10 _dafny.Set
	_f11 _dafny.Int
	_f12 bool
	_f13 bool
	_f15 _dafny.Int
	_f16 bool
	_f17 _dafny.Array
	_f18 _dafny.Int
	_f19 _dafny.Map
	_f20 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.EmptyMap
	_this.F3 = false
	_this.F9 = _dafny.Zero
	_this.F14 = false
	_this.F21 = _dafny.EmptySeq
	_this.F22 = _dafny.Zero
	_this.F23 = _dafny.EmptySet
	_this.F24 = false
	_this.F25 = _dafny.Zero
	_this._f1 = _dafny.EmptySet
	_this._f2 = _dafny.Zero
	_this._f4 = _dafny.Zero
	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = _dafny.EmptyMap
	_this._f7 = _dafny.Zero
	_this._f8 = _dafny.Zero
	_this._f10 = _dafny.EmptySet
	_this._f11 = _dafny.Zero
	_this._f12 = false
	_this._f13 = false
	_this._f15 = _dafny.Zero
	_this._f16 = false
	_this._f17 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f18 = _dafny.Zero
	_this._f19 = _dafny.EmptyMap
	_this._f20 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 _dafny.Set, f2 _dafny.Int, f3 bool, f4 _dafny.Int, f5 _dafny.CodePoint, f6 _dafny.Map, f7 _dafny.Int, f8 _dafny.Int, f9 _dafny.Int, f10 _dafny.Set, f11 _dafny.Int, f12 bool, f13 bool, f14 bool, f15 _dafny.Int, f16 bool, f17 _dafny.Array, f18 _dafny.Int, f19 _dafny.Map, f20 _dafny.Array, f21 _dafny.Sequence, f22 _dafny.Int, f23 _dafny.Set, f24 bool, f25 _dafny.Int) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this).F14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this)._f19 = f19
		(_this)._f20 = f20
		(_this).F21 = f21
		(_this).F22 = f22
		(_this).F23 = f23
		(_this).F24 = f24
		(_this).F25 = f25
	}
}
func (_this *GlobalState) F1() _dafny.Set {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
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
func (_this *GlobalState) F6() _dafny.Map {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.Int {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F10() _dafny.Set {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Int {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() bool {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Array {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() _dafny.Int {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F19() _dafny.Map {
	{
		return _this._f19
	}
}
func (_this *GlobalState) F20() _dafny.Array {
	{
		return _this._f20
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
func (_this *C0) Fm5(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rpbp")).Cardinality())).Minus((_dafny.IntOfInt64(704)).Times(_dafny.IntOfInt64(515)))
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
